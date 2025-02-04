package contractreader

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/smartcontractkit/chainlink-ccip/pkg/consts"

	"github.com/smartcontractkit/chainlink-common/pkg/services"
	"github.com/smartcontractkit/chainlink-common/pkg/types"
	clcommontypes "github.com/smartcontractkit/chainlink-common/pkg/types"
	"github.com/smartcontractkit/chainlink-common/pkg/types/query"
	"github.com/smartcontractkit/chainlink-common/pkg/types/query/primitives"

	"github.com/smartcontractkit/chainlink-ccip/internal/libs/slicelib"
)

var (
	ErrFinalityViolated = errors.New("finality violated")
	ErrTooManyBindings  = errors.New("too many bindings")
	ErrNoBindings       = errors.New("no bindings found")
)

// Extended version of a ContractReader.
type Extended interface {
	// Unbind is included for compatibility with ContractReader
	Unbind(ctx context.Context, bindings []types.BoundContract) error
	// HealthReport is included for compatibility with ContractReader
	HealthReport() map[string]error

	Bind(ctx context.Context, bindings []types.BoundContract) error

	GetBindings(contractName string) []ExtendedBoundContract

	// QueryKey is from the base contract reader interface.
	QueryKey(
		ctx context.Context,
		contract types.BoundContract,
		filter query.KeyFilter,
		limitAndSort query.LimitAndSort,
		sequenceDataType any,
	) ([]types.Sequence, error)

	// ExtendedQueryKey performs automatic binding from contractName to the first bound contract.
	// An error is generated if there are more than one bound contract for the contractName.
	ExtendedQueryKey(
		ctx context.Context,
		contractName string,
		filter query.KeyFilter,
		limitAndSort query.LimitAndSort,
		sequenceDataType any,
	) ([]types.Sequence, error)

	// GetLatestValue is from the base contract reader interface.
	GetLatestValue(
		ctx context.Context,
		readIdentifier string,
		confidenceLevel primitives.ConfidenceLevel,
		params, returnVal any,
	) error

	// ExtendedGetLatestValue performs automatic binding from contractName to the first bound contract, and
	// constructs a read identifier for a given method name. An error is generated if there are more than one
	// bound contract for the contractName.
	ExtendedGetLatestValue(
		ctx context.Context,
		contractName, methodName string,
		confidenceLevel primitives.ConfidenceLevel,
		params, returnVal any,
	) error

	// BatchGetLatestValues is from the base contract reader interface.
	BatchGetLatestValues(
		ctx context.Context,
		request types.BatchGetLatestValuesRequest,
	) (types.BatchGetLatestValuesResult, error)

	// ExtendedBatchGetLatestValues performs automatic binding from contractNames to bound contracts,
	// and constructs a BatchGetLatestValuesRequest with the resolved bindings.
	// If graceful is true, contracts with no bindings will be skipped and returned in the skipped slice
	// rather than returning an error.
	ExtendedBatchGetLatestValues(
		ctx context.Context,
		request ExtendedBatchGetLatestValuesRequest,
		graceful bool,
	) (types.BatchGetLatestValuesResult, []string, error)
}

type ExtendedBatchGetLatestValuesRequest map[string]types.ContractBatch

type ExtendedBoundContract struct {
	BoundAt time.Time
	Binding types.BoundContract
}

// extendedContractReader is an extended version of the contract reader.
type extendedContractReader struct {
	reader                 ContractReaderFacade
	contractBindingsByName map[string][]ExtendedBoundContract
	// contract names that allow multiple bindings
	multiBindAllowed map[string]bool
	mu               *sync.RWMutex
}

func NewExtendedContractReader(baseContractReader ContractReaderFacade) Extended {
	// avoid double wrapping
	if ecr, ok := baseContractReader.(Extended); ok {
		return ecr
	}
	return &extendedContractReader{
		reader:                 baseContractReader,
		contractBindingsByName: make(map[string][]ExtendedBoundContract),
		// so far this is the only contract that allows multiple bindings
		// if more contracts are added, this should be moved to a config
		multiBindAllowed: map[string]bool{consts.ContractNamePriceAggregator: true},
		mu:               &sync.RWMutex{},
	}
}

func (e *extendedContractReader) getOneBinding(contractName string) (ExtendedBoundContract, error) {
	extendedBindings := e.GetBindings(contractName)
	switch numBindings := len(extendedBindings); numBindings {
	case 1:
		return extendedBindings[0], nil
	case 0:
		return ExtendedBoundContract{}, fmt.Errorf(
			"getOneBinding: no binding found for %s contract: %w", contractName, ErrNoBindings)
	default:
		return ExtendedBoundContract{}, fmt.Errorf(
			"getOneBinding: expected one binding got %d: %w", numBindings, ErrTooManyBindings)
	}
}

func (e *extendedContractReader) QueryKey(
	ctx context.Context,
	contract types.BoundContract,
	filter query.KeyFilter,
	limitAndSort query.LimitAndSort,
	sequenceDataType any,
) ([]types.Sequence, error) {
	result, err := e.reader.QueryKey(
		ctx,
		contract,
		filter,
		limitAndSort,
		sequenceDataType,
	)

	// reads may update the reader health, so check for violations after every read.
	if e.hasFinalityViolation() {
		return nil, ErrFinalityViolated
	}

	return result, err
}

func (e *extendedContractReader) ExtendedQueryKey(
	ctx context.Context,
	contractName string,
	filter query.KeyFilter,
	limitAndSort query.LimitAndSort,
	sequenceDataType any,
) ([]types.Sequence, error) {
	binding, err := e.getOneBinding(contractName)
	if err != nil {
		return nil, fmt.Errorf("ExtendedQueryKey: %w", err)
	}

	return e.QueryKey(
		ctx,
		binding.Binding,
		filter,
		limitAndSort,
		sequenceDataType,
	)
}

func (e *extendedContractReader) GetLatestValue(
	ctx context.Context,
	readIdentifier string,
	confidenceLevel primitives.ConfidenceLevel,
	params, returnVal any,
) error {
	err := e.reader.GetLatestValue(
		ctx,
		readIdentifier,
		confidenceLevel,
		params,
		returnVal,
	)

	// reads may update the reader health, so check for violations after every read.
	if e.hasFinalityViolation() {
		return ErrFinalityViolated
	}

	return err
}

func (e *extendedContractReader) ExtendedGetLatestValue(
	ctx context.Context,
	contractName, methodName string,
	confidenceLevel primitives.ConfidenceLevel,
	params, returnVal any,
) error {
	binding, err := e.getOneBinding(contractName)
	if err != nil {
		return fmt.Errorf("ExtendedGetLatestValue: %w", err)
	}
	readIdentifier := binding.Binding.ReadIdentifier(methodName)

	return e.GetLatestValue(
		ctx,
		readIdentifier,
		confidenceLevel,
		params,
		returnVal,
	)
}

func (e *extendedContractReader) BatchGetLatestValues(
	ctx context.Context,
	request types.BatchGetLatestValuesRequest,
) (types.BatchGetLatestValuesResult, error) {
	result, err := e.reader.BatchGetLatestValues(ctx, request)

	// reads may update the reader health, so check for violations after every read.
	if e.hasFinalityViolation() {
		return nil, ErrFinalityViolated
	}

	return result, err
}

func (e *extendedContractReader) ExtendedBatchGetLatestValues(
	ctx context.Context,
	request ExtendedBatchGetLatestValuesRequest,
	graceful bool,
) (types.BatchGetLatestValuesResult, []string, error) {
	// Convert the request from contract names to BoundContracts
	convertedRequest := make(types.BatchGetLatestValuesRequest)
	var skippedContracts []string

	for contractName, batch := range request {
		// Get the binding for this contract name
		binding, err := e.getOneBinding(contractName)
		if err != nil {
			if graceful && errors.Is(err, ErrNoBindings) {
				// Track skipped contracts but continue processing
				skippedContracts = append(skippedContracts, contractName)
				continue
			}
			return nil, nil, fmt.Errorf("BatchGetLatestValues: failed to get binding for contract %s: %w", contractName, err)
		}

		// Use the resolved binding for the request
		convertedRequest[binding.Binding] = batch
	}

	// If we have no valid bindings after filtering in graceful mode, return empty result
	if graceful && len(convertedRequest) == 0 {
		return make(types.BatchGetLatestValuesResult), skippedContracts, nil
	}

	// Call the underlying BatchGetLatestValues with the converted request
	results, err := e.BatchGetLatestValues(ctx, convertedRequest)
	if err != nil {
		return nil, nil, err
	}

	return results, skippedContracts, nil
}

func (e *extendedContractReader) Bind(ctx context.Context, allBindings []types.BoundContract) error {
	validBindings := slicelib.Filter(allBindings, func(b types.BoundContract) bool { return !e.bindingExists(b) })
	if len(validBindings) == 0 {
		return nil
	}

	err := e.reader.Bind(ctx, validBindings)
	if err != nil {
		return fmt.Errorf("failed to call ContractReader.Bind: %w", err)
	}

	e.mu.Lock()
	defer e.mu.Unlock()
	for _, binding := range validBindings {
		if e.multiBindAllowed[binding.Name] {
			e.contractBindingsByName[binding.Name] = append(e.contractBindingsByName[binding.Name], ExtendedBoundContract{
				BoundAt: time.Now(),
				Binding: binding,
			})
		} else {
			if len(e.contractBindingsByName[binding.Name]) > 0 {
				// Unbind the previous binding
				err := e.reader.Unbind(ctx, []types.BoundContract{e.contractBindingsByName[binding.Name][0].Binding})
				if err != nil {
					return fmt.Errorf("failed to unbind previous binding: %w", err)
				}
			}
			// Override the previous binding
			e.contractBindingsByName[binding.Name] = []ExtendedBoundContract{{
				BoundAt: time.Now(),
				Binding: binding,
			}}
		}
	}

	return nil
}

func (e *extendedContractReader) GetBindings(contractName string) []ExtendedBoundContract {
	e.mu.RLock()
	defer e.mu.RUnlock()

	bindings, exists := e.contractBindingsByName[contractName]
	if !exists {
		return []ExtendedBoundContract{}
	}
	return bindings
}

func (e *extendedContractReader) bindingExists(b types.BoundContract) bool {
	e.mu.RLock()
	defer e.mu.RUnlock()

	for _, boundContracts := range e.contractBindingsByName {
		for _, boundContract := range boundContracts {
			if boundContract.Binding.String() == b.String() {
				return true
			}
		}
	}
	return false
}

// hasFinalityViolation checks the reader's HealthReport for a finality violated error.
// The report is based on the current known state, it does not proactively check for new errors.
// The state is typically updated as the LogPoller reads events from an rpc.
func (e *extendedContractReader) hasFinalityViolation() bool {
	report := e.reader.HealthReport()
	return services.ContainsError(
		report,
		clcommontypes.ErrFinalityViolated)
}

func (e *extendedContractReader) Unbind(ctx context.Context, bindings []types.BoundContract) error {
	return e.reader.Unbind(ctx, bindings)
}

func (e *extendedContractReader) HealthReport() map[string]error {
	return e.reader.HealthReport()
}

// Interface compliance check
var _ Extended = (*extendedContractReader)(nil)
