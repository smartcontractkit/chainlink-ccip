package contractreader

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/smartcontractkit/chainlink-common/pkg/types"
)

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
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			contractName := "testContract"

			var extendedBindings []ExtendedBoundContract
			for _, binding := range tt.bindings {
				extendedBindings = append(extendedBindings, ExtendedBoundContract{
					Binding: binding,
				})
			}
			extendedReader := &extendedContractReader{
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
