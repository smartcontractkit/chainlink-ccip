package mocks

import (
	"context"

	"github.com/stretchr/testify/mock"

	"github.com/smartcontractkit/chainlink-common/pkg/types"
	"github.com/smartcontractkit/chainlink-common/pkg/types/query"
	"github.com/smartcontractkit/chainlink-common/pkg/types/query/primitives"
)

var _ types.ContractReader = (*ContractReaderMock)(nil)

type ContractReaderMock struct {
	*mock.Mock
}

func NewContractReaderMock() *ContractReaderMock {
	return &ContractReaderMock{
		Mock: &mock.Mock{},
	}
}

func (cr *ContractReaderMock) GetLatestValue(ctx context.Context, readIdentifier string,
	confidenceLevel primitives.ConfidenceLevel, params, returnVal any) error {
	args := cr.Called(ctx, readIdentifier, confidenceLevel, params, returnVal)
	return args.Error(0)
}

func (cr *ContractReaderMock) BatchGetLatestValues(ctx context.Context,
	request types.BatchGetLatestValuesRequest) (types.BatchGetLatestValuesResult, error) {
	args := cr.Called(ctx, request)
	return args.Get(0).(types.BatchGetLatestValuesResult), args.Error(1)
}

func (cr *ContractReaderMock) Bind(ctx context.Context, bindings []types.BoundContract) error {
	args := cr.Called(ctx, bindings)
	return args.Error(0)
}

func (cr *ContractReaderMock) QueryKey(
	ctx context.Context,
	contract types.BoundContract,
	filter query.KeyFilter,
	limitAndSort query.LimitAndSort,
	sequenceDataType any,
) ([]types.Sequence, error) {
	args := cr.Called(ctx, contract, filter, limitAndSort, sequenceDataType)
	return args.Get(0).([]types.Sequence), args.Error(1)
}

func (cr *ContractReaderMock) Start(ctx context.Context) error {
	args := cr.Called(ctx)
	return args.Error(0)
}

func (cr *ContractReaderMock) Close() error {
	args := cr.Called()
	return args.Error(0)
}

func (cr *ContractReaderMock) Ready() error {
	panic("unimplemented")
}

func (cr *ContractReaderMock) HealthReport() map[string]error {
	args := cr.Called()
	return args.Get(0).(map[string]error)
}

func (cr *ContractReaderMock) Name() string {
	args := cr.Called()
	return args.String(0)
}

// Unbind implements types.ContractReader.
func (cr *ContractReaderMock) Unbind(ctx context.Context, bindings []types.BoundContract) error {
	args := cr.Called(ctx, bindings)
	return args.Error(0)
}
