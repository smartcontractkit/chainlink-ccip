package contractreader

import (
	"context"
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
		tt := tt
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
		name           string
		bindings       map[string][]ExtendedBoundContract
		request        ExtendedBatchGetLatestValuesRequest
		mockResponse   types.BatchGetLatestValuesResult
		expectedError  error
		expectedResult types.BatchGetLatestValuesResult
	}{
		{
			name: "single contract single read success",
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
			expectedError: nil,
			expectedResult: types.BatchGetLatestValuesResult{
				types.BoundContract{Name: "contract1", Address: "0x123"}: {
					{ReadName: "read1"},
				},
			},
		},
		{
			name: "contract not found",
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
			mockResponse:   nil,
			expectedError:  ErrNoBindings,
			expectedResult: nil,
		},
		{
			name: "multiple bindings for contract",
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
			mockResponse:   nil,
			expectedError:  ErrTooManyBindings,
			expectedResult: nil,
		},
		{
			name: "multiple contracts success",
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
			expectedError: nil,
			expectedResult: types.BatchGetLatestValuesResult{
				types.BoundContract{Name: "contract1", Address: "0x123"}: {
					{ReadName: "read1"},
				},
				types.BoundContract{Name: "contract2", Address: "0x456"}: {
					{ReadName: "read2"},
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockReader := &mockContractReader{
				BatchGetLatestValuesResponse: tt.mockResponse,
			}

			// Create extended reader with mock
			extendedReader := &extendedContractReader{
				ContractReaderFacade:   mockReader,
				contractBindingsByName: tt.bindings,
				mu:                     &sync.RWMutex{},
			}

			// Execute test
			result, err := extendedReader.ExtendedBatchGetLatestValues(context.Background(), tt.request)

			// Verify results
			if tt.expectedError != nil {
				assert.ErrorIs(t, err, tt.expectedError)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedResult, result)
			}
		})
	}
}

// mockContractReader implements ContractReaderFacade for testing
type mockContractReader struct {
	ContractReaderFacade
	BatchGetLatestValuesResponse types.BatchGetLatestValuesResult
}

func (m *mockContractReader) BatchGetLatestValues(_ context.Context, _ types.BatchGetLatestValuesRequest) (types.BatchGetLatestValuesResult, error) {
	return m.BatchGetLatestValuesResponse, nil
}
