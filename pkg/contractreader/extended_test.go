package contractreader_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-common/pkg/types"

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
