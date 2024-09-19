package contractreader

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/smartcontractkit/chainlink-common/pkg/types"

	chainreadermocks "github.com/smartcontractkit/chainlink-ccip/mocks/internal_/reader/contractreader"
)

func TestExtendedContractReader(t *testing.T) {
	const contractName = "testContract"
	cr := chainreadermocks.NewMockContractReaderFacade(t)
	extCr := NewExtendedContractReader(cr)

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
