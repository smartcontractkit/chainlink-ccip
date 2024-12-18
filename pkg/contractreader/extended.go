package contractreader

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/smartcontractkit/chainlink-common/pkg/services"
	"github.com/smartcontractkit/chainlink-common/pkg/types"
	clcommontypes "github.com/smartcontractkit/chainlink-common/pkg/types"
	"github.com/smartcontractkit/chainlink-common/pkg/types/query"
	"github.com/smartcontractkit/chainlink-common/pkg/types/query/primitives"

	"github.com/smartcontractkit/chainlink-ccip/internal/libs/slicelib"
)

var (
	ErrFinalityViolated = errors.New("finality violated")
	ErrTooManyBindings  = errors.New("contract binding not found")
	ErrNoBindings       = errors.New("no bindings found")
)

// Extended version of a ContractReader.
type Extended interface {
	Bind(ctx context.Context, bindings []types.BoundContract) error

	GetBindings(contractName string) []ExtendedBoundContract

	// ExtendedQueryKey performs automatic binding from contractName to the first bound contract.
	// An error is generated if there are more than one bound contract for the contractName.
	ExtendedQueryKey(
		ctx context.Context,
		contractName string,
		filter query.KeyFilter,
		limitAndSort query.LimitAndSort,
		sequenceDataType any,
	) ([]types.Sequence, error)

	// ExtendedGetLatestValue performs automatic binding from contractName to the first bound contract, and
	// constructs a read identifier for a given method name. An error is generated if there are more than one
	// bound contract for the contractName.
	ExtendedGetLatestValue(
		ctx context.Context,
		contractName, methodName string,
		confidenceLevel primitives.ConfidenceLevel,
		params, returnVal any,
	) error

	// ExtendedBatchGetLatestValues performs automatic binding from contractNames to bound contracts, and
	// contructs a BatchGetLatestValuesRequest with the resolved bindings.
	ExtendedBatchGetLatestValues(
		ctx context.Context,
		request ExtendedBatchGetLatestValuesRequest,
	) (types.BatchGetLatestValuesResult, error)
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
	mu                     *sync.RWMutex
}

func NewExtendedContractReader(baseContractReader ContractReaderFacade) Extended {
	return &extendedContractReader{
		reader:                 baseContractReader,
		contractBindingsByName: make(map[string][]ExtendedBoundContract),
		mu:                     &sync.RWMutex{},
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

func (e *extendedContractReader) ExtendedQueryKey(
	ctx context.Context,
	contractName string,
	filter query.KeyFilter,
	limitAndSort query.LimitAndSort,
	sequenceDataType any,
) ([]types.Sequence, error) {
	if e.hasFinalityViolation() {
		return nil, ErrFinalityViolated
	}

	binding, err := e.getOneBinding(contractName)
	if err != nil {
		return nil, fmt.Errorf("ExtendedQueryKey: %w", err)
	}

	return e.reader.QueryKey(
		ctx,
		binding.Binding,
		filter,
		limitAndSort,
		sequenceDataType,
	)
}

func (e *extendedContractReader) ExtendedGetLatestValue(
	ctx context.Context,
	contractName, methodName string,
	confidenceLevel primitives.ConfidenceLevel,
	params, returnVal any,
) error {
	if e.hasFinalityViolation() {
		return ErrFinalityViolated
	}

	binding, err := e.getOneBinding(contractName)
	if err != nil {
		return fmt.Errorf("ExtendedGetLatestValue: %w", err)
	}
	readIdentifier := binding.Binding.ReadIdentifier(methodName)

	return e.reader.GetLatestValue(
		ctx,
		readIdentifier,
		confidenceLevel,
		params,
		returnVal,
	)
}

func (e *extendedContractReader) ExtendedBatchGetLatestValues(
	ctx context.Context,
	request ExtendedBatchGetLatestValuesRequest,
) (types.BatchGetLatestValuesResult, error) {
	if e.hasFinalityViolation() {
		return nil, ErrFinalityViolated
	}

	// Convert the request from contract names to BoundContracts
	convertedRequest := make(types.BatchGetLatestValuesRequest)

	for contractName, batch := range request {
		// Get the binding for this contract name
		binding, err := e.getOneBinding(contractName)
		if err != nil {
			return nil, fmt.Errorf("BatchGetLatestValues: failed to get binding for contract %s: %w", contractName, err)
		}

		// Use the resolved binding for the request
		convertedRequest[binding.Binding] = batch
	}

	// Call the underlying BatchGetLatestValues with the converted request
	return e.reader.BatchGetLatestValues(ctx, convertedRequest)
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
		e.contractBindingsByName[binding.Name] = append(e.contractBindingsByName[binding.Name], ExtendedBoundContract{
			BoundAt: time.Now(),
			Binding: binding,
		})
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

func (e *extendedContractReader) hasFinalityViolation() bool {
	return services.ContainsError(
		e.reader.HealthReport(),
		clcommontypes.ErrFinalityViolated)
}

// Interface compliance check
var _ Extended = (*extendedContractReader)(nil)
