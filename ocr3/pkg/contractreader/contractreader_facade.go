package contractreader

import (
	"context"

	"github.com/smartcontractkit/chainlink-common/pkg/types"
	"github.com/smartcontractkit/chainlink-common/pkg/types/query"
	"github.com/smartcontractkit/chainlink-common/pkg/types/query/primitives"
)

// ContractReaderFacade wraps the public functions of ContractReader in chainlink-common so that we can mock it.
// See types.ContractReader in chainlink-common/pkg/types/contract_reader.go for details.
//
//nolint:lll // don't read this interface.
type ContractReaderFacade interface {
	GetLatestValue(ctx context.Context, readIdentifier string, confidenceLevel primitives.ConfidenceLevel, params, returnVal any) error
	BatchGetLatestValues(ctx context.Context, request types.BatchGetLatestValuesRequest) (types.BatchGetLatestValuesResult, error)
	Bind(ctx context.Context, bindings []types.BoundContract) error
	Unbind(ctx context.Context, bindings []types.BoundContract) error
	QueryKey(ctx context.Context, contract types.BoundContract, filter query.KeyFilter, limitAndSort query.LimitAndSort, sequenceDataType any) ([]types.Sequence, error)

	// HealthReport returns a full health report of the callee including its dependencies.
	// Keys are based on Name(), with nil values when healthy or errors otherwise.
	// Use CopyHealth to collect reports from sub-services.
	// This should run very fast, so avoid doing computation and instead prefer reporting pre-calculated state.
	// On finality violation report must contain at least one ErrFinalityViolation.
	HealthReport() map[string]error
}
