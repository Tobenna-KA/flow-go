package connection

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/sony/gobreaker"

	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/status"

	"github.com/onflow/flow-go/module"
)

// DefaultClientTimeout is used when making a GRPC request to a collection node or an execution node.
const DefaultClientTimeout = 3 * time.Second

// clientType is an enumeration type used to differentiate between different types of gRPC clients.
type clientType int

const (
	AccessClient clientType = iota
	ExecutionClient
)

type noopCloser struct{}

func (c *noopCloser) Close() error {
	return nil
}

// Manager provides methods for getting and managing gRPC client connections.
type Manager struct {
	cache                *Cache
	logger               zerolog.Logger
	metrics              module.AccessMetrics
	maxMsgSize           uint
	circuitBreakerConfig CircuitBreakerConfig
}

// CircuitBreakerConfig is a configuration struct for the circuit breaker.
type CircuitBreakerConfig struct {
	// Enabled specifies whether the circuit breaker is enabled for collection and execution API clients.
	Enabled bool
	// RestoreTimeout specifies the duration after which the circuit breaker will restore the connection to the client
	// after closing it due to failures.
	RestoreTimeout time.Duration
	// MaxFailures specifies the maximum number of failed calls to the client that will cause the circuit breaker
	// to close the connection.
	MaxFailures uint32
	// MaxRequests specifies the maximum number of requests to check if connection restored after timeout.
	MaxRequests uint32
}

// NewManager creates a new Manager with the specified parameters.
func NewManager(
	cache *Cache,
	logger zerolog.Logger,
	metrics module.AccessMetrics,
	maxMsgSize uint,
	circuitBreakerConfig CircuitBreakerConfig,
) Manager {
	return Manager{
		cache:                cache,
		logger:               logger,
		metrics:              metrics,
		maxMsgSize:           maxMsgSize,
		circuitBreakerConfig: circuitBreakerConfig,
	}
}

// GetConnection returns a gRPC client connection for the given grpcAddress and timeout.
// If a cache is used, it retrieves a cached connection, otherwise creates a new connection.
// It returns the client connection and an io.Closer to close the connection when done.
func (m *Manager) GetConnection(grpcAddress string, timeout time.Duration, clientType clientType) (*grpc.ClientConn, io.Closer, error) {
	if m.cache != nil {
		conn, err := m.retrieveConnection(grpcAddress, timeout, clientType)
		if err != nil {
			return nil, nil, err
		}
		return conn, &noopCloser{}, err
	}

	conn, err := m.createConnection(grpcAddress, timeout, nil, clientType)
	if err != nil {
		return nil, nil, err
	}

	return conn, io.Closer(conn), nil
}

// Remove removes the gRPC client connection associated with the given grpcAddress from the cache.
// It returns true if the connection was removed successfully, false otherwise.
func (m *Manager) Remove(grpcAddress string) bool {
	if m.cache == nil {
		return false
	}

	res, ok := m.cache.Get(grpcAddress)
	if !ok {
		return false
	}

	if !m.cache.Remove(grpcAddress) {
		return false
	}

	// Obtain the lock here to ensure that ClientConn was initialized, avoiding a situation with a nil ClientConn.
	res.mu.Lock()
	defer res.mu.Unlock()

	// Close the connection only if it is successfully removed from the cache
	res.Close()
	return true
}

// HasCache returns true if the Manager has a cache, false otherwise.
func (m *Manager) HasCache() bool {
	return m.cache != nil
}

// retrieveConnection retrieves the CachedClient for the given grpcAddress from the cache or adds a new one if not present.
// If the connection is already cached, it waits for the lock and returns the connection from the cache.
// Otherwise, it creates a new connection and caches it.
func (m *Manager) retrieveConnection(grpcAddress string, timeout time.Duration, clientType clientType) (*grpc.ClientConn, error) {
	client, ok := m.cache.GetOrAdd(grpcAddress, timeout)
	if ok {
		// The client was retrieved from the cache, wait for the lock
		client.mu.Lock()
		if m.metrics != nil {
			m.metrics.ConnectionFromPoolReused()
		}
	} else {
		// The client is new, lock is already held
		if m.metrics != nil {
			m.metrics.ConnectionAddedToPool()
		}
	}
	defer client.mu.Unlock()

	if client.ClientConn != nil && client.ClientConn.GetState() != connectivity.Shutdown {
		// Return the client connection from the cache
		return client.ClientConn, nil
	}

	// The connection is not cached or is closed, create a new connection and cache it
	conn, err := m.createConnection(grpcAddress, timeout, client, clientType)
	if err != nil {
		return nil, err
	}

	client.ClientConn = conn
	if m.metrics != nil {
		m.metrics.NewConnectionEstablished()
		m.metrics.TotalConnectionsInPool(uint(m.cache.Len()), uint(m.cache.MaxSize()))
	}

	return client.ClientConn, nil
}

// createConnection creates a new gRPC connection to the remote node at the given address with the specified timeout.
// If the cachedClient is not nil, it means a new entry in the cache is being created, so it's locked to give priority
// to the caller working with the new client, allowing it to create the underlying connection.
func (m *Manager) createConnection(address string, timeout time.Duration, cachedClient *CachedClient, clientType clientType) (*grpc.ClientConn, error) {
	if timeout == 0 {
		timeout = DefaultClientTimeout
	}

	keepaliveParams := keepalive.ClientParameters{
		Time:    10 * time.Second, // How long the client will wait before sending a keepalive to the server if there is no activity.
		Timeout: timeout,          // How long the client will wait for a response from the keepalive before closing.
	}

	// The order in which interceptors are added to the `connInterceptors` slice is important since they will be called
	// in the opposite order during gRPC requests. See documentation for more info:
	// https://grpc.io/blog/grpc-web-interceptor/#binding-interceptors
	var connInterceptors []grpc.UnaryClientInterceptor

	if !m.circuitBreakerConfig.Enabled {
		connInterceptors = append(connInterceptors, m.createClientInvalidationInterceptor(address, clientType))
	}

	connInterceptors = append(connInterceptors, createClientTimeoutInterceptor(timeout))

	// This interceptor monitors ongoing requests before passing control to subsequent interceptors.
	if cachedClient != nil {
		connInterceptors = append(connInterceptors, createRequestWatcherInterceptor(cachedClient))
	}

	if m.circuitBreakerConfig.Enabled {
		// If the circuit breaker interceptor is enabled, it should always be called first before passing control to
		// subsequent interceptors.
		connInterceptors = append(connInterceptors, m.createCircuitBreakerInterceptor())
	}

	// ClientConn's default KeepAlive on connections is indefinite, assuming the timeout isn't reached
	// The connections should be safe to be persisted and reused.
	// https://pkg.go.dev/google.golang.org/grpc#WithKeepaliveParams
	// https://grpc.io/blog/grpc-on-http2/#keeping-connections-alive
	conn, err := grpc.Dial(
		address,
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(int(m.maxMsgSize))),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithKeepaliveParams(keepaliveParams),
		grpc.WithChainUnaryInterceptor(connInterceptors...),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to address %s: %w", address, err)
	}
	return conn, nil
}

// createRequestWatcherInterceptor creates a request watcher interceptor to wait for unfinished requests before closing.
func createRequestWatcherInterceptor(cachedClient *CachedClient) grpc.UnaryClientInterceptor {
	requestWatcherInterceptor := func(
		ctx context.Context,
		method string,
		req interface{},
		reply interface{},
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error {
		// Prevent new requests from being sent if the connection is marked for closure.
		if cachedClient.closeRequested.Load() {
			return status.Errorf(codes.Unavailable, "the connection to %s was closed", cachedClient.Address)
		}

		// Increment the request counter to track ongoing requests, then decrement the request counter before returning.
		cachedClient.wg.Add(1)
		defer cachedClient.wg.Done()

		// Invoke the actual RPC method.
		return invoker(ctx, method, req, reply, cc, opts...)
	}

	return requestWatcherInterceptor
}

// WithClientTimeoutOption is a helper function to create a GRPC dial option
// with the specified client timeout interceptor.
func WithClientTimeoutOption(timeout time.Duration) grpc.DialOption {
	return grpc.WithUnaryInterceptor(createClientTimeoutInterceptor(timeout))
}

// createClientTimeoutInterceptor creates a client interceptor with a context that expires after the timeout.
func createClientTimeoutInterceptor(timeout time.Duration) grpc.UnaryClientInterceptor {
	clientTimeoutInterceptor := func(
		ctx context.Context,
		method string,
		req interface{},
		reply interface{},
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error {
		// Create a context that expires after the specified timeout.
		ctxWithTimeout, cancel := context.WithTimeout(ctx, timeout)
		defer cancel()

		// Call the remote GRPC using the short context.
		err := invoker(ctxWithTimeout, method, req, reply, cc, opts...)

		return err
	}

	return clientTimeoutInterceptor
}

// createClientInvalidationInterceptor creates a client interceptor for client invalidation. It should only be created
// if the circuit breaker is disabled. If the response from the server indicates an unavailable status, it invalidates
// the corresponding client.
func (m *Manager) createClientInvalidationInterceptor(
	address string,
	clientType clientType,
) grpc.UnaryClientInterceptor {
	if !m.circuitBreakerConfig.Enabled {
		clientInvalidationInterceptor := func(
			ctx context.Context,
			method string,
			req interface{},
			reply interface{},
			cc *grpc.ClientConn,
			invoker grpc.UnaryInvoker,
			opts ...grpc.CallOption,
		) error {
			err := invoker(ctx, method, req, reply, cc, opts...)
			if status.Code(err) == codes.Unavailable {
				switch clientType {
				case AccessClient:
					if m.Remove(address) {
						m.logger.Debug().Str("cached_access_client_invalidated", address).Msg("invalidating cached access client")
						if m.metrics != nil {
							m.metrics.ConnectionFromPoolInvalidated()
						}
					}
				case ExecutionClient:
					if m.Remove(address) {
						m.logger.Debug().Str("cached_execution_client_invalidated", address).Msg("invalidating cached execution client")
						if m.metrics != nil {
							m.metrics.ConnectionFromPoolInvalidated()
						}
					}
				default:
					m.logger.Info().Str("client_invalidation_interceptor", address).Msg(fmt.Sprintf("unexpected client type: %d", clientType))
				}
			}

			return err
		}

		return clientInvalidationInterceptor
	}

	return nil
}

// The simplified representation and description of circuit breaker pattern, that used to handle node connectivity:
//
// Circuit Open --> Circuit Half-Open --> Circuit Closed
//      ^                                      |
//      |                                      |
//      +--------------------------------------+
//
// The "Circuit Open" state represents the circuit being open, indicating that the node is not available.
// This state is entered when the number of consecutive failures exceeds the maximum allowed failures.
//
// The "Circuit Half-Open" state represents the circuit transitioning from the open state to the half-open
// state after a configured restore timeout. In this state, the circuit allows a limited number of requests
// to test if the node has recovered.
//
// The "Circuit Closed" state represents the circuit being closed, indicating that the node is available.
// This state is initial or entered when the test requests in the half-open state succeed.

// createCircuitBreakerInterceptor creates a client interceptor for circuit breaker functionality. It should only be
// created if the circuit breaker is enabled. All invocations will go through the circuit breaker to be tracked for
// success or failure of the call.
func (m *Manager) createCircuitBreakerInterceptor() grpc.UnaryClientInterceptor {
	if m.circuitBreakerConfig.Enabled {
		circuitBreaker := gobreaker.NewCircuitBreaker(gobreaker.Settings{
			// Timeout defines how long the circuit breaker will remain open before transitioning to the HalfClose state.
			Timeout: m.circuitBreakerConfig.RestoreTimeout,
			// ReadyToTrip returns true when the circuit breaker should trip and transition to the Open state
			ReadyToTrip: func(counts gobreaker.Counts) bool {
				// The number of maximum failures is checked before the circuit breaker goes to the Open state.
				return counts.ConsecutiveFailures >= m.circuitBreakerConfig.MaxFailures
			},
			// MaxRequests defines the max number of concurrent requests while the circuit breaker is in the HalfClosed
			// state.
			MaxRequests: m.circuitBreakerConfig.MaxRequests,
		})

		circuitBreakerInterceptor := func(
			ctx context.Context,
			method string,
			req interface{},
			reply interface{},
			cc *grpc.ClientConn,
			invoker grpc.UnaryInvoker,
			opts ...grpc.CallOption,
		) error {
			// The circuit breaker integration occurs here, where all invoked calls to the node pass through the
			// CircuitBreaker.Execute method. This method counts successful and failed invocations, and switches to the
			// "StateOpen" when the maximum failure threshold is reached. When the circuit breaker is in the "StateOpen"
			// it immediately rejects connections and returns without waiting for the call timeout. After the
			// "RestoreTimeout" period elapses, the circuit breaker transitions to the "StateHalfOpen" and attempts the
			// invocation again. If the invocation fails, it returns to the "StateOpen"; otherwise, it transitions to
			// the "StateClosed" and handles invocations as usual.
			_, err := circuitBreaker.Execute(func() (interface{}, error) {
				err := invoker(ctx, method, req, reply, cc, opts...)
				return nil, err
			})
			return err
		}

		return circuitBreakerInterceptor
	}

	return nil
}
