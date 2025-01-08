package contractreader_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-common/pkg/types"
	"github.com/smartcontractkit/chainlink-common/pkg/types/query"
	"github.com/smartcontractkit/chainlink-common/pkg/types/query/primitives"
	"github.com/smartcontractkit/chainlink-common/pkg/utils/tests"

	chainreadermocks "github.com/smartcontractkit/chainlink-ccip/mocks/pkg/contractreader"
	"github.com/smartcontractkit/chainlink-ccip/pkg/contractreader"
)

func TestExtendedContractReader(t *testing.T) {
	const contractName = "testContract"
	cr := chainreadermocks.NewMockContractReaderFacade(t)
	extCr := contractreader.NewExtendedContractReader(cr)

	bindings := extCr.GetBindings(contractName)
	assert.Len(t, bindings, 0)

	cr.On("Bind", context.Background(),
		[]types.BoundContract{{Name: contractName, Address: "0x123"}}).Return(nil)
	cr.On("Bind", context.Background(),
		[]types.BoundContract{{Name: contractName, Address: "0x124"}}).Return(nil)
	cr.On("Bind", context.Background(),
		[]types.BoundContract{{Name: contractName, Address: "0x125"}}).Return(fmt.Errorf("some err"))

	err := extCr.Bind(context.Background(), []types.BoundContract{{Name: contractName, Address: "0x123"}})
	assert.NoError(t, err)

	// ignored since 0x123 already exists
	err = extCr.Bind(context.Background(), []types.BoundContract{{Name: contractName, Address: "0x123"}})
	assert.NoError(t, err)

	err = extCr.Bind(context.Background(), []types.BoundContract{{Name: contractName, Address: "0x124"}})
	assert.NoError(t, err)

	// Bind fails
	err = extCr.Bind(context.Background(), []types.BoundContract{{Name: contractName, Address: "0x125"}})
	assert.Error(t, err)

	bindings = extCr.GetBindings(contractName)
	assert.Len(t, bindings, 2)
	assert.Equal(t, "0x123", bindings[0].Binding.Address)
	assert.Equal(t, "0x124", bindings[1].Binding.Address)
}

func TestDoubleWrap(t *testing.T) {
	var cr contractreader.ContractReaderFacade

	wrapped := contractreader.NewExtendedContractReader(cr)
	require.NotEqual(t, &cr, &wrapped)

	doubleWrapped := contractreader.NewExtendedContractReader(cr)
	require.Equal(t, wrapped, doubleWrapped)
}

func TestFinalityViolation(t *testing.T) {
	cr := chainreadermocks.NewMockContractReaderFacade(t)
	cr.EXPECT().HealthReport().Return(map[string]error{"farmerwolfcabbagegoat": types.ErrFinalityViolated}).Times(3)
	cr.EXPECT().QueryKey(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil, nil)
	cr.EXPECT().GetLatestValue(mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)
	cr.EXPECT().BatchGetLatestValues(mock.Anything, mock.Anything).Return(nil, nil)

	wrapped := contractreader.NewExtendedContractReader(cr)

	_, err := wrapped.QueryKey(
		tests.Context(t),
		types.BoundContract{},
		query.KeyFilter{},
		query.LimitAndSort{},
		nil)
	require.ErrorIs(t, err, contractreader.ErrFinalityViolated)

	err = wrapped.GetLatestValue(
		tests.Context(t),
		"",
		primitives.Finalized,
		nil,
		nil)
	require.ErrorIs(t, err, contractreader.ErrFinalityViolated)

	_, err = wrapped.BatchGetLatestValues(
		tests.Context(t),
		types.BatchGetLatestValuesRequest{})
	require.ErrorIs(t, err, contractreader.ErrFinalityViolated)
}
