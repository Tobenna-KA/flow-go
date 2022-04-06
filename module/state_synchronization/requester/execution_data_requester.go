package requester

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/hashicorp/go-multierror"
	"github.com/ipfs/go-blockservice"
	"github.com/ipfs/go-datastore"
	blockstore "github.com/ipfs/go-ipfs-blockstore"
	"github.com/rs/zerolog"
	"github.com/sethvargo/go-retry"

	"github.com/onflow/flow-go/consensus/hotstuff/model"
	"github.com/onflow/flow-go/engine"
	"github.com/onflow/flow-go/model/encoding/cbor"
	"github.com/onflow/flow-go/model/flow"
	"github.com/onflow/flow-go/module"
	"github.com/onflow/flow-go/module/component"
	"github.com/onflow/flow-go/module/irrecoverable"
	"github.com/onflow/flow-go/module/jobqueue"
	"github.com/onflow/flow-go/module/metrics"
	"github.com/onflow/flow-go/module/state_synchronization"
	"github.com/onflow/flow-go/module/state_synchronization/requester/status"
	"github.com/onflow/flow-go/network"
	"github.com/onflow/flow-go/network/compressor"
	"github.com/onflow/flow-go/network/p2p"
	"github.com/onflow/flow-go/state/protocol"
	"github.com/onflow/flow-go/storage"
)

// The ExecutionDataRequester downloads ExecutionData for sealed blocks from other participants
// in the network running the ExecutionDataService. The requester is made up of 4 subcomponents:
//
// * OnBlockFinalized:        receives block finalized events from the finalization distributor and
//                            pushes them to the finalizedBlocks channel.
//
// * bootstrap: is a single worker that consumes from the finalizedBlocks channel,
//                            and generates ExecutionData fetch requests for the all blocks with
//                            results sealed in the finalized block.
//
// * fetchRequestProcessor:   is a pool of workers that consume fetch requests and download the
//                            ExecutionData from the network.
//
// * notificationProcessor:   is a single worker that sends notifications to registered consumers
//                            in sequential block height order.
//
//                                                               +-----------------------+
//                                                            +--| fetchRequestProcessor |--+
//                                                            |  +-----------------------+  |
//    +------------------+      +-------------------------+   |  +-----------------------+  |   +-----------------------+
// -->| OnBlockFinalized |<---->| bootstrap |<--+->| fetchRequestProcessor |<-+-->| notificationProcessor |
//    +------------------+      +-------------------------+   |  +-----------------------+  |   +-----------------------+
//                                                            |  +-----------------------+  |
//                                                            +--| fetchRequestProcessor |--+
//                                                               +-----------------------+
//
// The requester has 2 main priorities:
//   1. ensure execution state is as widely distributed as possible among the network participants
//   2. make the execution state available to local subscribers
// #1 is the top priority, and this component is optimized to download and seed ExecutionData for
// as many blocks as possible, making them available to other nodes in the network. This ensures
// execution state is available for all participants, and reduces network load on the execution
// nodes that source the data.

const (
	// DefaultFetchTimeout is the default timeout for fetching ExecutionData from the db/network
	DefaultFetchTimeout = 5 * time.Minute

	// DefaultRetryDelay is the default initial delay used in the exponential backoff for failed
	// ExecutionData download retries
	DefaultRetryDelay = 10 * time.Second

	// DefaultMaxRetryDelay is the default maximum delay used in the exponential backoff for failed
	// ExecutionData download retries
	DefaultMaxRetryDelay = 5 * time.Minute

	// DefaultMaxCachedEntries is the default the number of ExecutionData objects to keep when
	// waiting to send notifications.
	DefaultMaxCachedEntries = 50

	// DefaultMaxSearchAhead is the default max number of unsent notifications to allow before
	// pausing new fetches.
	DefaultMaxSearchAhead = 5000

	// Number of goroutines to use for downloading new ExecutionData from the network.
	fetchWorkers = 4
)

// ErrRequesterHalted is returned when an invalid ExectutionData is encountered
var ErrRequesterHalted = errors.New("requester was halted due to invalid data")

// ExecutionDataReceivedCallback is a callback that is called ExecutionData is received for a new block
type ExecutionDataReceivedCallback func(*state_synchronization.ExecutionData)

type ExecutionDataConfig struct {
	// The first block height for which to request ExecutionData
	StartBlockHeight uint64

	// Max number of ExecutionData objects to keep when waiting to send notifications.
	// Dropped data is refetched from disk.
	MaxCachedEntries uint64

	// Max number of unsent notifications to allow before pausing new fetches. After exceeding this
	// limit, the requester will stop processing new finalized block notifications. This prevents
	// unbounded memory use by the requester if it gets stuck fetching a specific height.
	MaxSearchAhead uint64

	// The timeout for fetching ExecutionData from the db/network
	FetchTimeout time.Duration

	// Exponential backoff settings for download retries
	RetryDelay    time.Duration
	MaxRetryDelay time.Duration

	// Whether or not to run datastore check on startup
	CheckEnabled bool
}

// ExecutionDataRequester downloads ExecutionData for newly sealed blocks from the network using the
// ExecutionDataService.
type ExecutionDataRequester interface {
	component.Component
	OnBlockFinalized(*model.Block)
	AddOnExecutionDataFetchedConsumer(fn ExecutionDataReceivedCallback)
}

type executionDataRequesterImpl struct {
	component.Component
	cm      *component.ComponentManager
	ds      datastore.Batching
	bs      network.BlobService
	eds     state_synchronization.ExecutionDataService
	metrics module.ExecutionDataRequesterMetrics
	log     zerolog.Logger

	// Local db objects
	blocks  storage.Blocks
	results storage.ExecutionResults

	status *status.Status
	config ExecutionDataConfig

	// Notifiers for queue consumers
	blockNotifier engine.Notifier

	// Job queues
	blockConsumer        *jobqueue.WrappedConsumer
	notificationConsumer *jobqueue.WrappedConsumer

	// List of callbacks to call when ExecutionData is successfully fetched for a block
	consumers []ExecutionDataReceivedCallback

	consumerMu sync.RWMutex
}

// New creates a new execution data requester component
func New(
	log zerolog.Logger,
	edrMetrics module.ExecutionDataRequesterMetrics,
	datastore datastore.Batching,
	blobservice network.BlobService,
	eds state_synchronization.ExecutionDataService,
	processedHeight storage.ConsumerProgress,
	processedNotifications storage.ConsumerProgress,
	state protocol.State,
	blocks storage.Blocks,
	results storage.ExecutionResults,
	cfg ExecutionDataConfig,
) (ExecutionDataRequester, error) {
	e := &executionDataRequesterImpl{
		log:           log.With().Str("component", "execution_data_requester").Logger(),
		ds:            datastore,
		bs:            blobservice,
		eds:           eds,
		metrics:       edrMetrics,
		blocks:        blocks,
		results:       results,
		config:        cfg,
		blockNotifier: engine.NewNotifier(),
	}

	executionDataNotifier := engine.NewNotifier()
	rootHeight := e.config.StartBlockHeight - 1

	var err error
	e.blockConsumer, err = jobqueue.NewWrappedConsumer(
		log.With().Str("module", "block_consumer").Logger(),
		processedHeight,
		status.NewSealedBlockReader(state, blocks),
		e.processBlockJob,
		e.blockNotifier.Channel(),
		rootHeight,
		jobqueue.WithMaxProcessing(fetchWorkers),
		// notifies notificationConsumer when new ExecutionData blobs are available
		jobqueue.WitNotifier(func(module.JobID) { executionDataNotifier.Notify() }),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create block consumer: %w", err)
	}

	e.status = status.New(
		log.With().Str("module", "requester_status").Logger(),
		cfg.MaxCachedEntries,
		cfg.MaxSearchAhead,
		e.blockConsumer,
		processedNotifications,
	)

	e.notificationConsumer, err = jobqueue.NewWrappedConsumer(
		log.With().Str("module", "notification_consumer").Logger(),
		processedNotifications,
		e.status,
		e.processNotificationJob,
		executionDataNotifier.Channel(),
		rootHeight,
		// kick notifier to make sure we scan until the last available notification
		jobqueue.WitNotifier(func(module.JobID) { executionDataNotifier.Notify() }),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create notification consumer: %w", err)
	}

	builder := component.NewComponentManagerBuilder().
		AddWorker(e.bootstrap).
		AddWorker(func(ctx irrecoverable.SignalerContext, ready component.ReadyFunc) {
			ready()

			// Don't start processing new blocks until everything else is ready
			<-e.notificationConsumer.Ready()
			<-e.Ready()

			e.blockConsumer.Start(ctx)
			<-e.blockConsumer.Ready()

			<-e.blockConsumer.Done()
		}).
		AddWorker(func(ctx irrecoverable.SignalerContext, ready component.ReadyFunc) {
			e.notificationConsumer.Start(ctx)
			<-e.notificationConsumer.Ready()
			ready()

			<-e.notificationConsumer.Done()
		})

	e.cm = builder.Build()
	e.Component = e.cm

	return e, nil
}

func (e *executionDataRequesterImpl) OnBlockFinalized(*model.Block) {
	e.blockNotifier.Notify()
}

// AddOnExecutionDataFetchedConsumer adds a callback to be called when a new ExecutionData is received
// Callback Implementations must:
//   * be concurrency safe
//   * be non-blocking
//   * handle repetition of the same events (with some processing overhead).
func (e *executionDataRequesterImpl) AddOnExecutionDataFetchedConsumer(fn ExecutionDataReceivedCallback) {
	e.consumerMu.Lock()
	defer e.consumerMu.Unlock()

	e.consumers = append(e.consumers, fn)
}

// bootstrap runs the main process that processes finalized block notifications and
// requests ExecutionData for each block seal
func (e *executionDataRequesterImpl) bootstrap(ctx irrecoverable.SignalerContext, ready component.ReadyFunc) {
	// needs state from notificationConsumer
	<-e.notificationConsumer.Ready()

	// Load previous requester state from db if it exists
	err := e.status.Load()
	if err != nil {
		e.log.Error().Err(err).Msg("failed to load notification state. using defaults")
	}

	// Requester came up halted.
	if e.status.Halted() {
		e.log.Error().Msg("HALTING REQUESTER: requester was halted on a previous run due to invalid data")
		// By returning before ready, none of the other workers will start, effectively
		// disabling the component.
		return
	}

	e.log.Debug().Msgf("starting with LastNotified: %d", e.status.LastNotified())

	<-e.eds.Ready()

	// only run datastore check if enabled
	if e.config.CheckEnabled {
		err = e.checkDatastore(ctx, e.status.LastNotified())

		// Any error should crash the component
		if err != nil {
			ctx.Throw(err)
		}
	}

	ready()
}

// Fetch Worker Methods

// processBlockJob consumes jobs from the blockConsumer and attempts to download an ExecutionData
// for the given block height.
func (e *executionDataRequesterImpl) processBlockJob(ctx irrecoverable.SignalerContext, job module.Job, complete func()) {
	// convert job into a block entry
	block, err := status.JobToBlock(job)
	if err != nil {
		ctx.Throw(fmt.Errorf("failed to convert job to block: %w", err))
	}

	request := &status.BlockEntry{
		BlockID: block.ID(),
		Height:  block.Header.Height,
	}

	err = e.processSealedHeight(ctx, request)
	if err == nil {
		complete()
	}

	// all errors are thrown as irrecoverable errors except context cancellation
}

// processSealedHeight downloads ExecutionData for the given block height.
// If the download fails, it will retry using forever, using exponential backoff.
func (e *executionDataRequesterImpl) processSealedHeight(ctx irrecoverable.SignalerContext, request *status.BlockEntry) error {
	// download execution data for the block

	backoff, err := retry.NewExponential(e.config.RetryDelay)
	if err != nil {
		ctx.Throw(fmt.Errorf("failed to create retry mechanism: %w", err))
	}
	backoff = retry.WithCappedDuration(e.config.MaxRetryDelay, backoff)
	backoff = retry.WithJitterPercent(15, backoff)

	// the only error returned is ctx.Err()
	attempt := 0
	return retry.Do(ctx, backoff, func(context.Context) error {
		err := e.processFetchRequest(ctx, request)

		if attempt > 0 {
			e.metrics.FetchRetried()
		}
		attempt++

		return retry.RetryableError(err)
	})
}

func (e *executionDataRequesterImpl) processFetchRequest(ctx irrecoverable.SignalerContext, request *status.BlockEntry) error {
	logger := e.log.With().Str("block_id", request.BlockID.String()).Logger()
	logger.Debug().Msgf("processing fetch request for block %d", request.Height)

	result, err := e.results.ByBlockID(request.BlockID)

	// By the time the block is sealed, the ExecutionResult must be in the db
	if err != nil {
		ctx.Throw(fmt.Errorf("failed to lookup execution result for block: %w", err))
	}

	executionData, err := e.fetchExecutionData(ctx, result.ExecutionDataID, request.Height)

	if isInvalidBlobError(err) {
		// This means an execution result was sealed with an invalid execution data id (invalid data).
		// Eventually, verification nodes will verify that the execution data is valid, and not sign the receipt
		logger.Error().Err(err).
			Str("execution_data_id", result.ExecutionDataID.String()).
			Msg("HALTING REQUESTER: invalid execution data found")

		e.status.Halt()
		e.metrics.Halted()
		ctx.Throw(ErrRequesterHalted)
	}

	// Some or all of the blob was missing or corrupt. retry
	if isBlobNotFoundError(err) || errors.Is(err, context.DeadlineExceeded) || errors.Is(err, context.Canceled) {
		logger.Error().Err(err).Msg("failed to get execution data for block")

		return err
	}

	// Any other error is unexpected
	if err != nil {
		logger.Error().Err(err).
			Str("execution_data_id", result.ExecutionDataID.String()).
			Msg("unexpected error fetching execution data")

		ctx.Throw(err)
	}

	logger.Debug().Msgf("Fetched execution data for block %d", request.Height)

	request.ExecutionData = executionData

	e.status.Fetched(request)

	return nil
}

// fetchExecutionData fetches the ExecutionData by its ID, and times out if fetchTimeout is exceeded
func (e *executionDataRequesterImpl) fetchExecutionData(signalerCtx irrecoverable.SignalerContext, executionDataID flow.Identifier, height uint64) (executionData *state_synchronization.ExecutionData, err error) {
	ctx, cancel := context.WithTimeout(signalerCtx, e.config.FetchTimeout)
	defer cancel()

	start := time.Now()
	e.metrics.ExecutionDataFetchStarted()
	defer func() {
		// needs to be run inside a closure so the variables are resolved when the defer is executed
		e.metrics.ExecutionDataFetchFinished(time.Since(start), err == nil, height)
	}()

	// Get the data from the network
	// this is a blocking call, won't be unblocked until either hitting error (including timeout) or 
	// the data is received
	executionData, err = e.eds.Get(ctx, executionDataID)

	if err != nil {
		return nil, err
	}

	// Write it to the local blobstore
	_, _, err = e.eds.Add(signalerCtx, executionData)

	if err != nil {
		return nil, fmt.Errorf("failed to write execution data to blobstore: %w", err)
	}

	return executionData, nil
}

// Notification Worker Methods

func (e *executionDataRequesterImpl) processNotificationJob(ctx irrecoverable.SignalerContext, job module.Job, complete func()) {
	// convert job into a block entry
	entry, err := status.JobToBlockEntry(job)
	if err != nil {
		ctx.Throw(fmt.Errorf("failed to convert job to entry: %w", err))
	}

	e.processNotification(ctx, entry)
	complete()
}

func (e *executionDataRequesterImpl) processNotification(ctx irrecoverable.SignalerContext, entry *status.BlockEntry) {
	e.log.Debug().Msgf("notifying for block %d", entry.Height)

	// ExecutionData may have been purged, in which case, look it up again
	if entry.ExecutionData == nil {
		e.log.Debug().Msgf("execution data not in cache for block %d", entry.Height)

		// get it from db
		result, err := e.results.ByBlockID(entry.BlockID)
		if err != nil {
			ctx.Throw(fmt.Errorf("failed to lookup execution result for block: %w", err))
		}

		entry.ExecutionData, err = e.eds.Get(ctx, result.ExecutionDataID)

		if err != nil {
			// At this point the data has been downloaded and validated, so it should be available
			ctx.Throw(fmt.Errorf("failed to get execution data for block: %w", err))
		}
	}

	// send notifications
	e.notifyConsumers(entry.ExecutionData)
	e.metrics.NotificationSent(entry.Height)
}

func (e *executionDataRequesterImpl) notifyConsumers(executionData *state_synchronization.ExecutionData) {

	e.consumerMu.RLock()
	defer e.consumerMu.RUnlock()

	for _, fn := range e.consumers {
		fn(executionData)
	}
}

// Bootstrap Methods

func (e *executionDataRequesterImpl) checkDatastore(ctx irrecoverable.SignalerContext, lastHeight uint64) error {
	// we're only interested in inspecting blobs that exist in our local db, so create an
	// ExecutionDataService that only uses the local datastore
	localEDS, err := e.localExecutionDataService(ctx)

	if errors.Is(err, context.Canceled) {
		return nil
	}

	if err != nil {
		return fmt.Errorf("failed to create local ExecutionDataService: %w", err)
	}

	// Search from genesis to the lastHeight, and confirm data is still available for all heights
	// Update the notification state based on the data in the db. All data should be present, otherwise
	// data was deleted or corrupted on disk.
	for height := e.config.StartBlockHeight; height <= lastHeight; height++ {
		if ctx.Err() != nil {
			return nil
		}

		block, err := e.blocks.ByHeight(height)
		if err != nil {
			return fmt.Errorf("failed to get block for height %d: %w", height, err)
		}

		result, err := e.results.ByBlockID(block.ID())
		if err != nil {
			return fmt.Errorf("failed to lookup execution result for block %d: %w", block.ID(), err)
		}

		exists, err := e.checkExecutionData(ctx, localEDS, result.ExecutionDataID)

		if errors.Is(err, ErrRequesterHalted) {
			e.log.Error().Err(err).
				Str("block_id", block.ID().String()).
				Str("execution_data_id", result.ExecutionDataID.String()).
				Msg("HALTING REQUESTER: invalid execution data found")

			e.status.Halt()
			e.metrics.Halted()
			ctx.Throw(ErrRequesterHalted)
		}

		if err != nil {
			return err
		}

		if !exists {
			// blocking download of any missing block that should exist in the datastore
			e.processSealedHeight(ctx, &status.BlockEntry{
				BlockID: block.ID(),
				Height:  height,
			})
		}
	}

	return nil
}

func (e *executionDataRequesterImpl) checkExecutionData(ctx irrecoverable.SignalerContext, localEDS state_synchronization.ExecutionDataService, rootID flow.Identifier) (bool, error) {
	invalidCIDs, ok := localEDS.Check(ctx, rootID)

	// Check returns a list of CIDs with the corresponding errors encountered while retrieving their
	// data from the local datastore.

	if ok {
		return true, nil
	}

	// Track if this blob should be refetched
	missing := false

	var errs *multierror.Error
	for _, invalidCID := range invalidCIDs {
		cid := invalidCID.Cid
		err := invalidCID.Err

		// Not Found, just report and continue
		if errors.Is(err, blockservice.ErrNotFound) {
			missing = true
			continue
		}

		// The blob's hash didn't match. This is a special case where the data was corrupted on
		// disk. Delete and refetch
		if errors.Is(err, blockstore.ErrHashMismatch) {
			if err := e.bs.DeleteBlob(ctx, cid); err != nil {
				return false, fmt.Errorf("failed to delete corrupted CID %s from rootCID %s: %w", cid, rootID, err)
			}

			missing = true
			continue
		}

		// It should not be possible to encounter one of these errors since they are checked when
		// the block is originally received
		if isInvalidBlobError(err) {
			return false, ErrRequesterHalted
		}

		// Record any other errors to return
		errs = multierror.Append(errs, err)
	}

	return missing, errs.ErrorOrNil()
}

// localExecutionDataService returns an ExecutionDataService that's configured to use only the local
// datastore, and to rehash blobs on read. This is used to check the validity of blobs that exist in
// the local db.
func (e *executionDataRequesterImpl) localExecutionDataService(ctx irrecoverable.SignalerContext) (state_synchronization.ExecutionDataService, error) {
	blobService := p2p.NewBlobService(e.ds, p2p.WithHashOnRead(true))
	blobService.Start(ctx)

	eds := state_synchronization.NewExecutionDataService(
		new(cbor.Codec),
		compressor.NewLz4Compressor(),
		blobService,
		metrics.NewNoopCollector(),
		e.log,
	)

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case <-eds.Ready():
	}

	return eds, nil
}

func isInvalidBlobError(err error) bool {
	var malformedDataError *state_synchronization.MalformedDataError
	var blobSizeLimitExceededError *state_synchronization.BlobSizeLimitExceededError
	return errors.As(err, &malformedDataError) ||
		errors.As(err, &blobSizeLimitExceededError) ||
		errors.Is(err, state_synchronization.ErrBlobTreeDepthExceeded)
}

func isBlobNotFoundError(err error) bool {
	var blobNotFoundError *state_synchronization.BlobNotFoundError
	return errors.As(err, &blobNotFoundError)
}
