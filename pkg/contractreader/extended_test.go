package contractreader

import (
	"context"
	"fmt"
	"sync"
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

func TestGetOneBinding(t *testing.T) {
	tests := []struct {
		name          string
		bindings      []types.BoundContract
		expectedError error
	}{
		{
			name:          "no bindings",
			bindings:      []types.BoundContract{},
			expectedError: ErrNoBindings,
		},
		{
			name: "one binding",
			bindings: []types.BoundContract{
				{Name: "testContract", Address: "0x123"},
			},
			expectedError: nil,
		},
		{
			name: "multiple bindings",
			bindings: []types.BoundContract{
				{Name: "testContract", Address: "0x123"},
				{Name: "testContract", Address: "0x124"},
			},
			expectedError: ErrTooManyBindings,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			cr := chainreadermocks.NewMockContractReaderFacade(t)
			contractName := "testContract"

			var extendedBindings []ExtendedBoundContract
			for _, binding := range tt.bindings {
				extendedBindings = append(extendedBindings, ExtendedBoundContract{
					Binding: binding,
				})
			}
			extendedReader := &extendedContractReader{
				ContractReaderFacade: cr,
				contractBindingsByName: map[string][]ExtendedBoundContract{
					contractName: extendedBindings,
				},
				mu: &sync.RWMutex{},
			}

			_, err := extendedReader.getOneBinding(contractName)
			if tt.expectedError != nil {
				assert.ErrorIs(t, err, tt.expectedError)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
