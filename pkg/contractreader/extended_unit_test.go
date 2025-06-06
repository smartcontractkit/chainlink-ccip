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

func TestExtractTxHash(t *testing.T) {
	testCases := []struct {
		name                 string // Name of the test case for t.Run
		cursor               string // Input string for ExtractTxHash
		expectedHash         string // Expected transaction hash if successful
		expectError          bool   // True if an error is expected
		expectedFullErrorMsg string // The exact error message string expected, if expectError is true
	}{
		// --- Valid Cases ---
		{
			name:         "Standard valid hash",
			cursor:       "12345-1-0xabcdef1234567890",
			expectedHash: "0xabcdef1234567890",
			expectError:  false,
		},
		{
			name:         "TxHash with internal hyphens",
			cursor:       "123-0-my-tx-hash-with-hyphens",
			expectedHash: "my-tx-hash-with-hyphens",
			expectError:  false,
		},
		{
			name:         "Conceptual empty BlockNumber part (first element of split)",
			cursor:       "-0-0xabcdef", // strings.SplitN yields ["", "0", "0xabcdef"]
			expectedHash: "0xabcdef",
			expectError:  false,
		},
		{
			name:         "Conceptual empty LogIndex part (second element of split)",
			cursor:       "123--0xabcdef", // strings.SplitN yields ["123", "", "0xabcdef"]
			expectedHash: "0xabcdef",
			expectError:  false,
		},
		{
			name:         "Zero values for BlockNumber and LogIndex parts",
			cursor:       "0-0-validhash",
			expectedHash: "validhash",
			expectError:  false,
		},

		// --- Invalid Format Cases (triggered by len(parts) < 3) ---
		{
			name:                 "Too few parts - one hyphen",
			cursor:               "12345-1",
			expectError:          true,
			expectedFullErrorMsg: "invalid cursor format: '12345-1'. Expected format 'BlockNumber-LogIndex-TxHash'",
		},
		{
			name:                 "Too few parts - no hyphens",
			cursor:               "1234510xabcdef",
			expectError:          true,
			expectedFullErrorMsg: "invalid cursor format: '1234510xabcdef'. Expected format 'BlockNumber-LogIndex-TxHash'",
		},
		{
			name:                 "Empty string input",
			cursor:               "",
			expectError:          true,
			expectedFullErrorMsg: "invalid cursor format: ''. Expected format 'BlockNumber-LogIndex-TxHash'",
		},
		{
			name:                 "Only one part and a trailing hyphen",
			cursor:               "123-", // strings.SplitN yields ["123", ""], len(parts) is 2
			expectError:          true,
			expectedFullErrorMsg: "invalid cursor format: '123-'. Expected format 'BlockNumber-LogIndex-TxHash'",
		},
		{
			name:                 "Only a single hyphen",
			cursor:               "-", // strings.SplitN yields ["", ""], len(parts) is 2
			expectError:          true,
			expectedFullErrorMsg: "invalid cursor format: '-'. Expected format 'BlockNumber-LogIndex-TxHash'",
		},

		// --- Empty Transaction Hash Case (triggered by len(txHash) == 0) ---
		{
			name:                 "Valid structure but empty TxHash part",
			cursor:               "123-0-", // strings.SplitN yields ["123", "0", ""], txHash is ""
			expectError:          true,
			expectedFullErrorMsg: "transaction hash is empty in cursor: '123-0-'",
		},
		{
			name:                 "Two hyphens, no content (results in empty TxHash part)",
			cursor:               "--", // strings.SplitN yields ["", "", ""], txHash is ""
			expectError:          true,
			expectedFullErrorMsg: "transaction hash is empty in cursor: '--'",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			hash, err := ExtractTxHash(tc.cursor)

			if tc.expectError {
				// We expect an error
				if err == nil {
					t.Errorf("ExtractTxHash(%q) expected an error, but got nil", tc.cursor)
				}
				// Check if the error message is exactly as expected
				if err.Error() != tc.expectedFullErrorMsg {
					t.Errorf("ExtractTxHash(%q)\n   error: %q\nwant error: %q", tc.cursor, err.Error(), tc.expectedFullErrorMsg)
				}
			} else {
				// We expect no error
				if err != nil {
					t.Errorf("ExtractTxHash(%q) expected no error, but got: %v", tc.cursor, err)
				}
				// Check if the returned hash matches the expected hash
				if hash != tc.expectedHash {
					t.Errorf("ExtractTxHash(%q) = %q, want %q", tc.cursor, hash, tc.expectedHash)
				}
			}
		})
	}
}
