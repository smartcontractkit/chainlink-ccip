package contractreader

import (
	"context"
	"sync"
	"testing"

	"github.com/smartcontractkit/chainlink-common/pkg/types"
	"github.com/stretchr/testify/assert"
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

func TestExtendedBatchGetLatestValues(t *testing.T) {
	tests := []struct {
		name            string
		bindings        map[string][]ExtendedBoundContract
		request         ExtendedBatchGetLatestValuesRequest
		mockResponse    types.BatchGetLatestValuesResult
		graceful        bool
		expectedError   error
		expectedResult  types.BatchGetLatestValuesResult
		expectedSkipped []string
	}{
		{
			name: "single contract single read success - non-graceful",
			bindings: map[string][]ExtendedBoundContract{
				"contract1": {
					{Binding: types.BoundContract{Name: "contract1", Address: "0x123"}},
				},
			},
			request: ExtendedBatchGetLatestValuesRequest{
				"contract1": {
					{ReadName: "read1", Params: "params1", ReturnVal: "return1"},
				},
			},
			mockResponse: types.BatchGetLatestValuesResult{
				types.BoundContract{Name: "contract1", Address: "0x123"}: {
					{ReadName: "read1"},
				},
			},
			graceful:      false,
			expectedError: nil,
			expectedResult: types.BatchGetLatestValuesResult{
				types.BoundContract{Name: "contract1", Address: "0x123"}: {
					{ReadName: "read1"},
				},
			},
			expectedSkipped: nil,
		},
		{
			name: "single contract single read success - graceful",
			bindings: map[string][]ExtendedBoundContract{
				"contract1": {
					{Binding: types.BoundContract{Name: "contract1", Address: "0x123"}},
				},
			},
			request: ExtendedBatchGetLatestValuesRequest{
				"contract1": {
					{ReadName: "read1", Params: "params1", ReturnVal: "return1"},
				},
			},
			mockResponse: types.BatchGetLatestValuesResult{
				types.BoundContract{Name: "contract1", Address: "0x123"}: {
					{ReadName: "read1"},
				},
			},
			graceful:      true,
			expectedError: nil,
			expectedResult: types.BatchGetLatestValuesResult{
				types.BoundContract{Name: "contract1", Address: "0x123"}: {
					{ReadName: "read1"},
				},
			},
			expectedSkipped: nil,
		},
		{
			name: "contract not found - non-graceful",
			bindings: map[string][]ExtendedBoundContract{
				"contract1": {
					{Binding: types.BoundContract{Name: "contract1", Address: "0x123"}},
				},
			},
			request: ExtendedBatchGetLatestValuesRequest{
				"nonexistent": {
					{ReadName: "read1", Params: "params1", ReturnVal: "return1"},
				},
			},
			mockResponse:    nil,
			graceful:        false,
			expectedError:   ErrNoBindings,
			expectedResult:  nil,
			expectedSkipped: nil,
		},
		{
			name: "contract not found - graceful",
			bindings: map[string][]ExtendedBoundContract{
				"contract1": {
					{Binding: types.BoundContract{Name: "contract1", Address: "0x123"}},
				},
			},
			request: ExtendedBatchGetLatestValuesRequest{
				"nonexistent": {
					{ReadName: "read1", Params: "params1", ReturnVal: "return1"},
				},
			},
			mockResponse:    nil,
			graceful:        true,
			expectedError:   nil,
			expectedResult:  types.BatchGetLatestValuesResult{},
			expectedSkipped: []string{"nonexistent"},
		},
		{
			name: "multiple bindings error - both modes",
			bindings: map[string][]ExtendedBoundContract{
				"contract1": {
					{Binding: types.BoundContract{Name: "contract1", Address: "0x123"}},
					{Binding: types.BoundContract{Name: "contract1", Address: "0x124"}},
				},
			},
			request: ExtendedBatchGetLatestValuesRequest{
				"contract1": {
					{ReadName: "read1", Params: "params1", ReturnVal: "return1"},
				},
			},
			mockResponse:    nil,
			graceful:        true, // Should fail even in graceful mode
			expectedError:   ErrTooManyBindings,
			expectedResult:  nil,
			expectedSkipped: nil,
		},
		{
			name: "multiple contracts with mixed results - graceful",
			bindings: map[string][]ExtendedBoundContract{
				"contract1": {
					{Binding: types.BoundContract{Name: "contract1", Address: "0x123"}},
				},
				"contract2": {
					{Binding: types.BoundContract{Name: "contract2", Address: "0x456"}},
				},
			},
			request: ExtendedBatchGetLatestValuesRequest{
				"contract1": {
					{ReadName: "read1", Params: "params1", ReturnVal: "return1"},
				},
				"contract2": {
					{ReadName: "read2", Params: "params2", ReturnVal: "return2"},
				},
				"nonexistent": {
					{ReadName: "read3", Params: "params3", ReturnVal: "return3"},
				},
			},
			mockResponse: types.BatchGetLatestValuesResult{
				types.BoundContract{Name: "contract1", Address: "0x123"}: {
					{ReadName: "read1"},
				},
				types.BoundContract{Name: "contract2", Address: "0x456"}: {
					{ReadName: "read2"},
				},
			},
			graceful:      true,
			expectedError: nil,
			expectedResult: types.BatchGetLatestValuesResult{
				types.BoundContract{Name: "contract1", Address: "0x123"}: {
					{ReadName: "read1"},
				},
				types.BoundContract{Name: "contract2", Address: "0x456"}: {
					{ReadName: "read2"},
				},
			},
			expectedSkipped: []string{"nonexistent"},
		},
		{
			name: "multiple contracts - non-graceful",
			bindings: map[string][]ExtendedBoundContract{
				"contract1": {
					{Binding: types.BoundContract{Name: "contract1", Address: "0x123"}},
				},
				"contract2": {
					{Binding: types.BoundContract{Name: "contract2", Address: "0x456"}},
				},
			},
			request: ExtendedBatchGetLatestValuesRequest{
				"contract1": {
					{ReadName: "read1", Params: "params1", ReturnVal: "return1"},
				},
				"contract2": {
					{ReadName: "read2", Params: "params2", ReturnVal: "return2"},
				},
			},
			mockResponse: types.BatchGetLatestValuesResult{
				types.BoundContract{Name: "contract1", Address: "0x123"}: {
					{ReadName: "read1"},
				},
				types.BoundContract{Name: "contract2", Address: "0x456"}: {
					{ReadName: "read2"},
				},
			},
			graceful:      false,
			expectedError: nil,
			expectedResult: types.BatchGetLatestValuesResult{
				types.BoundContract{Name: "contract1", Address: "0x123"}: {
					{ReadName: "read1"},
				},
				types.BoundContract{Name: "contract2", Address: "0x456"}: {
					{ReadName: "read2"},
				},
			},
			expectedSkipped: nil,
		},
		{
			name: "all contracts skipped - graceful",
			bindings: map[string][]ExtendedBoundContract{
				"contract1": {
					{Binding: types.BoundContract{Name: "contract1", Address: "0x123"}},
				},
			},
			request: ExtendedBatchGetLatestValuesRequest{
				"nonexistent1": {
					{ReadName: "read1", Params: "params1", ReturnVal: "return1"},
				},
				"nonexistent2": {
					{ReadName: "read2", Params: "params2", ReturnVal: "return2"},
				},
			},
			mockResponse:    nil,
			graceful:        true,
			expectedError:   nil,
			expectedResult:  types.BatchGetLatestValuesResult{},
			expectedSkipped: []string{"nonexistent1", "nonexistent2"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockReader := &mockContractReader{
				BatchGetLatestValuesResponse: tt.mockResponse,
			}

			extendedReader := &extendedContractReader{
				reader:                 mockReader,
				contractBindingsByName: tt.bindings,
				mu:                     &sync.RWMutex{},
			}

			result, skipped, err := extendedReader.ExtendedBatchGetLatestValues(context.Background(), tt.request, tt.graceful)

			if tt.expectedError != nil {
				assert.ErrorIs(t, err, tt.expectedError)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedResult, result)
				assert.ElementsMatch(t, tt.expectedSkipped, skipped)
			}
		})
	}
}

// mockContractReader implements ContractReaderFacade for testing
type mockContractReader struct {
	ContractReaderFacade
	BatchGetLatestValuesResponse types.BatchGetLatestValuesResult
}

func (m *mockContractReader) BatchGetLatestValues(
	_ context.Context,
	_ types.BatchGetLatestValuesRequest,
) (types.BatchGetLatestValuesResult, error) {
	return m.BatchGetLatestValuesResponse, nil
}

func (m *mockContractReader) HealthReport() map[string]error {
	return nil
}
