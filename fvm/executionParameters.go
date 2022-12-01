package fvm

import (
	"context"
	"fmt"
	"math"

	"github.com/onflow/cadence"
	"github.com/onflow/cadence/runtime"
	"github.com/onflow/cadence/runtime/common"

	"github.com/onflow/flow-go/fvm/blueprints"
	"github.com/onflow/flow-go/fvm/environment"
	"github.com/onflow/flow-go/fvm/errors"
	"github.com/onflow/flow-go/fvm/meter"
	"github.com/onflow/flow-go/fvm/programs"
	"github.com/onflow/flow-go/fvm/state"
	"github.com/onflow/flow-go/fvm/utils"
)

// getBasicMeterParameters returns the set of meter parameters used for
// general procedure execution.  Subparts of the procedure execution may
// specify custom meter parameters via nested transactions.
func getBasicMeterParameters(
	ctx Context,
	proc Procedure,
) meter.MeterParameters {
	params := meter.DefaultParameters().
		WithComputationLimit(uint(proc.ComputationLimit(ctx))).
		WithMemoryLimit(proc.MemoryLimit(ctx)).
		WithEventEmitByteLimit(ctx.EventCollectionByteSizeLimit).
		WithStorageInteractionLimit(ctx.MaxStateInteractionSize)

	// NOTE: The memory limit (and interaction limit) may be overridden by the
	// environment.  We need to ignore the override in that case.
	if proc.ShouldDisableMemoryAndInteractionLimits(ctx) {
		params = params.WithMemoryLimit(math.MaxUint64).
			WithStorageInteractionLimit(math.MaxUint64)
	}

	return params
}

// getBodyMeterParameters returns the set of meter parameters used for
// transaction/script body execution.
func getBodyMeterParameters(
	ctx Context,
	proc Procedure,
	txnState *state.TransactionState,
	derivedTxnData *programs.DerivedTransactionData,
) (
	meter.MeterParameters,
	error,
) {
	procParams := getBasicMeterParameters(ctx, proc)

	overrides, err := derivedTxnData.GetMeterParamOverrides(
		txnState,
		NewMeterParamOverridesComputer(ctx, derivedTxnData))
	if err != nil {
		return procParams, err
	}

	if overrides.ComputationWeights != nil {
		procParams = procParams.WithComputationWeights(
			overrides.ComputationWeights)
	}

	if overrides.MemoryWeights != nil {
		procParams = procParams.WithMemoryWeights(overrides.MemoryWeights)
	}

	if overrides.MemoryLimit != nil {
		procParams = procParams.WithMemoryLimit(*overrides.MemoryLimit)
	}

	// NOTE: The memory limit (and interaction limit) may be overridden by the
	// environment.  We need to ignore the override in that case.
	if proc.ShouldDisableMemoryAndInteractionLimits(ctx) {
		procParams = procParams.WithMemoryLimit(math.MaxUint64).
			WithStorageInteractionLimit(math.MaxUint64)
	}

	return procParams, nil
}

type MeterParamOverridesComputer struct {
	ctx            Context
	derivedTxnData *programs.DerivedTransactionData
}

func NewMeterParamOverridesComputer(
	ctx Context,
	derivedTxnData *programs.DerivedTransactionData,
) MeterParamOverridesComputer {
	return MeterParamOverridesComputer{ctx, derivedTxnData}
}

func (computer MeterParamOverridesComputer) Compute(
	txnState *state.TransactionState,
	_ struct{},
) (
	programs.MeterParamOverrides,
	error,
) {
	var overrides programs.MeterParamOverrides
	var err error
	txnState.RunWithAllLimitsDisabled(func() {
		overrides, err = computer.getMeterParamOverrides(txnState)
	})

	if err != nil {
		return overrides, fmt.Errorf(
			"error getting environment meter parameter overrides: %w",
			err)
	}

	return overrides, nil
}

func (computer MeterParamOverridesComputer) getMeterParamOverrides(
	txnState *state.TransactionState,
) (
	programs.MeterParamOverrides,
	error,
) {
	// Check that the service account exists because all the settings are
	// stored in it
	serviceAddress := computer.ctx.Chain.ServiceAddress()
	service := runtime.Address(serviceAddress)

	env := environment.NewScriptEnvironment(
		context.Background(),
		computer.ctx.EnvironmentParams,
		txnState,
		computer.derivedTxnData)

	overrides := programs.MeterParamOverrides{}

	// set the property if no error, but if the error is a fatal error then
	// return it
	setIfOk := func(prop string, err error, setter func()) (fatal error) {
		err, fatal = errors.SplitErrorTypes(err)
		if fatal != nil {
			// this is a fatal error. return it
			computer.ctx.Logger.
				Error().
				Err(fatal).
				Msgf("error getting %s", prop)
			return fatal
		}
		if err != nil {
			// this is a general error.
			// could be that no setting was present in the state,
			// or that the setting was not parseable,
			// or some other deterministic thing.
			computer.ctx.Logger.
				Debug().
				Err(err).
				Msgf("could not set %s. Using defaults", prop)
			return nil
		}
		// everything is ok. do the setting
		setter()
		return nil
	}

	computationWeights, err := GetExecutionEffortWeights(env, service)
	err = setIfOk(
		"execution effort weights",
		err,
		func() { overrides.ComputationWeights = computationWeights })
	if err != nil {
		return overrides, err
	}

	memoryWeights, err := GetExecutionMemoryWeights(env, service)
	err = setIfOk(
		"execution memory weights",
		err,
		func() { overrides.MemoryWeights = memoryWeights })
	if err != nil {
		return overrides, err
	}

	memoryLimit, err := GetExecutionMemoryLimit(env, service)
	err = setIfOk(
		"execution memory limit",
		err,
		func() { overrides.MemoryLimit = &memoryLimit })
	if err != nil {
		return overrides, err
	}

	return overrides, nil
}

func getExecutionWeights[KindType common.ComputationKind | common.MemoryKind](
	env environment.Environment,
	service runtime.Address,
	path cadence.Path,
	defaultWeights map[KindType]uint64,
) (
	map[KindType]uint64,
	error,
) {
	runtime := env.BorrowCadenceRuntime()
	defer env.ReturnCadenceRuntime(runtime)

	value, err := runtime.ReadStored(service, path)

	if err != nil {
		// this might be fatal, return as is
		return nil, err
	}

	weightsRaw, ok := utils.CadenceValueToWeights(value)
	if !ok {
		// this is a non-fatal error. It is expected if the weights are not set up on the network yet.
		return nil, errors.NewCouldNotGetExecutionParameterFromStateError(
			service.Hex(),
			path.String())
	}

	// Merge the default weights with the weights from the state.
	// This allows for weights that are not set in the state, to be set by default.
	// In case the network is stuck because of a transaction using an FVM feature that has 0 weight
	// (or is not metered at all), the defaults can be changed and the network restarted
	// instead of trying to change the weights with a transaction.
	weights := make(map[KindType]uint64, len(defaultWeights))
	for k, v := range defaultWeights {
		weights[k] = v
	}
	for k, v := range weightsRaw {
		weights[KindType(k)] = v
	}

	return weights, nil
}

// GetExecutionEffortWeights reads stored execution effort weights from the service account
func GetExecutionEffortWeights(
	env environment.Environment,
	service runtime.Address,
) (
	computationWeights meter.ExecutionEffortWeights,
	err error,
) {
	return getExecutionWeights(
		env,
		service,
		blueprints.TransactionFeesExecutionEffortWeightsPath,
		meter.DefaultComputationWeights)
}

// GetExecutionMemoryWeights reads stored execution memory weights from the service account
func GetExecutionMemoryWeights(
	env environment.Environment,
	service runtime.Address,
) (
	memoryWeights meter.ExecutionMemoryWeights,
	err error,
) {
	return getExecutionWeights(
		env,
		service,
		blueprints.TransactionFeesExecutionMemoryWeightsPath,
		meter.DefaultMemoryWeights)
}

// GetExecutionMemoryLimit reads the stored execution memory limit from the service account
func GetExecutionMemoryLimit(
	env environment.Environment,
	service runtime.Address,
) (
	memoryLimit uint64,
	err error,
) {
	runtime := env.BorrowCadenceRuntime()
	defer env.ReturnCadenceRuntime(runtime)

	value, err := runtime.ReadStored(
		service,
		blueprints.TransactionFeesExecutionMemoryLimitPath)
	if err != nil {
		// this might be fatal, return as is
		return 0, err
	}

	memoryLimitRaw, ok := value.(cadence.UInt64)
	if value == nil || !ok {
		// this is a non-fatal error. It is expected if the weights are not set up on the network yet.
		return 0, errors.NewCouldNotGetExecutionParameterFromStateError(
			service.Hex(),
			blueprints.TransactionFeesExecutionMemoryLimitPath.String())
	}

	return memoryLimitRaw.ToGoValue().(uint64), nil
}
