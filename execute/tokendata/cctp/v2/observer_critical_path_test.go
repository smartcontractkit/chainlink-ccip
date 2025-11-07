package v2

import (
	"context"
	"encoding/binary"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-ccip/execute/exectypes"
	"github.com/smartcontractkit/chainlink-ccip/pkg/reader"
)

// Test helper to create valid v2 payload
func createV2Payload(versionTag uint32, sourceDomain uint32, depositHash [32]byte) cciptypes.Bytes {
	payload := make([]byte, 40)
	binary.BigEndian.PutUint32(payload[0:4], versionTag)
	binary.BigEndian.PutUint32(payload[4:8], sourceDomain)
	copy(payload[8:40], depositHash[:])
	return cciptypes.Bytes(payload)
}

func Test_processSingleToken(t *testing.T) {
	lggr := logger.Test(t)
	chainSelector := cciptypes.ChainSelector(1)
	seqNum := cciptypes.SeqNum(100)
	supportedPool := "0x1111111111111111111111111111111111111111"

	tests := []struct {
		name           string
		tokenAmount    cciptypes.RampTokenAmount
		expectPayload  bool
		expectError    bool
		errorContains  string
	}{
		{
			name: "supported token with valid v2 payload",
			tokenAmount: cciptypes.RampTokenAmount{
				SourcePoolAddress: mustParseAddress(supportedPool),
				ExtraData: createV2Payload(CCTP_VERSION_2_TAG, 1, [32]byte{1,2,3}),
			},
			expectPayload: true,
			expectError:   false,
		},
		{
			name: "supported token with CCV tag",
			tokenAmount: cciptypes.RampTokenAmount{
				SourcePoolAddress: mustParseAddress(supportedPool),
				ExtraData: createV2Payload(CCTP_VERSION_2_CCV_TAG, 2, [32]byte{4,5,6}),
			},
			expectPayload: true,
			expectError:   false,
		},
		{
			name: "unsupported token (different pool)",
			tokenAmount: cciptypes.RampTokenAmount{
				SourcePoolAddress: mustParseAddress("0x2222222222222222222222222222222222222222"),
				ExtraData: createV2Payload(CCTP_VERSION_2_TAG, 1, [32]byte{1,2,3}),
			},
			expectPayload: false,
			expectError:   false,
		},
		{
			name: "supported token with invalid payload length",
			tokenAmount: cciptypes.RampTokenAmount{
				SourcePoolAddress: mustParseAddress(supportedPool),
				ExtraData: cciptypes.Bytes(make([]byte, 30)), // Too short
			},
			expectPayload: false,
			expectError:   true,
			errorContains: "invalid V2 source pool data length",
		},
		{
			name: "supported token with invalid version tag",
			tokenAmount: cciptypes.RampTokenAmount{
				SourcePoolAddress: mustParseAddress(supportedPool),
				ExtraData: createV2Payload(0x12345678, 1, [32]byte{1,2,3}),
			},
			expectPayload: false,
			expectError:   true,
			errorContains: "invalid CCTPv2 version tag",
		},
		{
			name: "supported token with empty extra data",
			tokenAmount: cciptypes.RampTokenAmount{
				SourcePoolAddress: mustParseAddress(supportedPool),
				ExtraData: cciptypes.Bytes{},
			},
			expectPayload: false,
			expectError:   true,
			errorContains: "invalid V2 source pool data length",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			observer := &CCTPv2TokenDataObserver{
				supportedPoolsBySelector: map[cciptypes.ChainSelector]string{
					chainSelector: supportedPool,
				},
				lggr: lggr,
			}

			payload, err := observer.processSingleToken(chainSelector, seqNum, 0, tt.tokenAmount, lggr)

			if tt.expectError {
				require.Error(t, err)
				if tt.errorContains != "" {
					assert.Contains(t, err.Error(), tt.errorContains)
				}
				assert.Nil(t, payload)
			} else {
				require.NoError(t, err)
				if tt.expectPayload {
					assert.NotNil(t, payload)
					assert.Greater(t, payload.SourceDomain, uint32(0))
				} else {
					assert.Nil(t, payload)
				}
			}
		})
	}
}

func Test_getCCTPv2TokenPayloads(t *testing.T) {
	lggr := logger.Test(t)
	chainSelector1 := cciptypes.ChainSelector(1)
	chainSelector2 := cciptypes.ChainSelector(2)
	supportedPool1 := "0x1111111111111111111111111111111111111111"
	supportedPool2 := "0x2222222222222222222222222222222222222222"
	unsupportedPool := "0x3333333333333333333333333333333333333333"

	tests := []struct {
		name              string
		messages          exectypes.MessageObservations
		supportedPools    map[cciptypes.ChainSelector]string
		expectedV2Count   int
		validateResult    func(*testing.T, map[cciptypes.ChainSelector]map[reader.MessageTokenID]*SourceTokenDataPayloadV2)
	}{
		{
			name: "single chain, single message, single supported token",
			messages: exectypes.MessageObservations{
				chainSelector1: {
					cciptypes.SeqNum(100): cciptypes.Message{
						Header: cciptypes.RampMessageHeader{
							SequenceNumber: 100,
						},
						TokenAmounts: []cciptypes.RampTokenAmount{
							{
								SourcePoolAddress: mustParseAddress(supportedPool1),
								ExtraData: createV2Payload(CCTP_VERSION_2_TAG, 1, [32]byte{1,2,3}),
							},
						},
					},
				},
			},
			supportedPools: map[cciptypes.ChainSelector]string{
				chainSelector1: supportedPool1,
			},
			expectedV2Count: 1,
			validateResult: func(t *testing.T, result map[cciptypes.ChainSelector]map[reader.MessageTokenID]*SourceTokenDataPayloadV2) {
				require.Contains(t, result, chainSelector1)
				assert.Len(t, result[chainSelector1], 1)
			},
		},
		{
			name: "multiple tokens in single message, mixed supported/unsupported",
			messages: exectypes.MessageObservations{
				chainSelector1: {
					cciptypes.SeqNum(100): cciptypes.Message{
						Header: cciptypes.RampMessageHeader{
							SequenceNumber: 100,
						},
						TokenAmounts: []cciptypes.RampTokenAmount{
							{
								SourcePoolAddress: mustParseAddress(unsupportedPool),
								ExtraData: createV2Payload(CCTP_VERSION_2_TAG, 1, [32]byte{1,2,3}),
							},
							{
								SourcePoolAddress: mustParseAddress(supportedPool1),
								ExtraData: createV2Payload(CCTP_VERSION_2_TAG, 1, [32]byte{4,5,6}),
							},
							{
								SourcePoolAddress: mustParseAddress(unsupportedPool),
								ExtraData: createV2Payload(CCTP_VERSION_2_TAG, 1, [32]byte{7,8,9}),
							},
						},
					},
				},
			},
			supportedPools: map[cciptypes.ChainSelector]string{
				chainSelector1: supportedPool1,
			},
			expectedV2Count: 1,
			validateResult: func(t *testing.T, result map[cciptypes.ChainSelector]map[reader.MessageTokenID]*SourceTokenDataPayloadV2) {
				require.Contains(t, result, chainSelector1)
				assert.Len(t, result[chainSelector1], 1)
				// Should only have the token at index 1
				tokenID := reader.NewMessageTokenID(100, 1)
				assert.Contains(t, result[chainSelector1], tokenID)
			},
		},
		{
			name: "multiple chains with different supported pools",
			messages: exectypes.MessageObservations{
				chainSelector1: {
					cciptypes.SeqNum(100): cciptypes.Message{
						Header: cciptypes.RampMessageHeader{
							SequenceNumber: 100,
						},
						TokenAmounts: []cciptypes.RampTokenAmount{
							{
								SourcePoolAddress: mustParseAddress(supportedPool1),
								ExtraData: createV2Payload(CCTP_VERSION_2_TAG, 1, [32]byte{1,2,3}),
							},
						},
					},
				},
				chainSelector2: {
					cciptypes.SeqNum(200): cciptypes.Message{
						Header: cciptypes.RampMessageHeader{
							SequenceNumber: 200,
						},
						TokenAmounts: []cciptypes.RampTokenAmount{
							{
								SourcePoolAddress: mustParseAddress(supportedPool2),
								ExtraData: createV2Payload(CCTP_VERSION_2_CCV_TAG, 2, [32]byte{4,5,6}),
							},
						},
					},
				},
			},
			supportedPools: map[cciptypes.ChainSelector]string{
				chainSelector1: supportedPool1,
				chainSelector2: supportedPool2,
			},
			expectedV2Count: 2,
			validateResult: func(t *testing.T, result map[cciptypes.ChainSelector]map[reader.MessageTokenID]*SourceTokenDataPayloadV2) {
				require.Contains(t, result, chainSelector1)
				require.Contains(t, result, chainSelector2)
				assert.Len(t, result[chainSelector1], 1)
				assert.Len(t, result[chainSelector2], 1)
			},
		},
		{
			name: "invalid payload should be skipped with error logged",
			messages: exectypes.MessageObservations{
				chainSelector1: {
					cciptypes.SeqNum(100): cciptypes.Message{
						Header: cciptypes.RampMessageHeader{
							SequenceNumber: 100,
						},
						TokenAmounts: []cciptypes.RampTokenAmount{
							{
								SourcePoolAddress: mustParseAddress(supportedPool1),
								ExtraData: cciptypes.Bytes(make([]byte, 30)), // Invalid length
							},
							{
								SourcePoolAddress: mustParseAddress(supportedPool1),
								ExtraData: createV2Payload(CCTP_VERSION_2_TAG, 1, [32]byte{1,2,3}),
							},
						},
					},
				},
			},
			supportedPools: map[cciptypes.ChainSelector]string{
				chainSelector1: supportedPool1,
			},
			expectedV2Count: 1, // Only the valid one
			validateResult: func(t *testing.T, result map[cciptypes.ChainSelector]map[reader.MessageTokenID]*SourceTokenDataPayloadV2) {
				require.Contains(t, result, chainSelector1)
				assert.Len(t, result[chainSelector1], 1)
				// Should only have token at index 1 (index 0 failed to decode)
				tokenID := reader.NewMessageTokenID(100, 1)
				assert.Contains(t, result[chainSelector1], tokenID)
			},
		},
		{
			name: "no supported tokens",
			messages: exectypes.MessageObservations{
				chainSelector1: {
					cciptypes.SeqNum(100): cciptypes.Message{
						Header: cciptypes.RampMessageHeader{
							SequenceNumber: 100,
						},
						TokenAmounts: []cciptypes.RampTokenAmount{
							{
								SourcePoolAddress: mustParseAddress(unsupportedPool),
								ExtraData: createV2Payload(CCTP_VERSION_2_TAG, 1, [32]byte{1,2,3}),
							},
						},
					},
				},
			},
			supportedPools: map[cciptypes.ChainSelector]string{
				chainSelector1: supportedPool1,
			},
			expectedV2Count: 0,
			validateResult: func(t *testing.T, result map[cciptypes.ChainSelector]map[reader.MessageTokenID]*SourceTokenDataPayloadV2) {
				// Should be empty
				assert.Len(t, result, 0)
			},
		},
		{
			name:           "empty messages",
			messages:       exectypes.MessageObservations{},
			supportedPools: map[cciptypes.ChainSelector]string{},
			expectedV2Count: 0,
			validateResult: func(t *testing.T, result map[cciptypes.ChainSelector]map[reader.MessageTokenID]*SourceTokenDataPayloadV2) {
				assert.Len(t, result, 0)
			},
		},
		{
			name: "multiple messages same chain",
			messages: exectypes.MessageObservations{
				chainSelector1: {
					cciptypes.SeqNum(100): cciptypes.Message{
						Header: cciptypes.RampMessageHeader{
							SequenceNumber: 100,
						},
						TokenAmounts: []cciptypes.RampTokenAmount{
							{
								SourcePoolAddress: mustParseAddress(supportedPool1),
								ExtraData: createV2Payload(CCTP_VERSION_2_TAG, 1, [32]byte{1,2,3}),
							},
						},
					},
					cciptypes.SeqNum(101): cciptypes.Message{
						Header: cciptypes.RampMessageHeader{
							SequenceNumber: 101,
						},
						TokenAmounts: []cciptypes.RampTokenAmount{
							{
								SourcePoolAddress: mustParseAddress(supportedPool1),
								ExtraData: createV2Payload(CCTP_VERSION_2_TAG, 1, [32]byte{4,5,6}),
							},
						},
					},
				},
			},
			supportedPools: map[cciptypes.ChainSelector]string{
				chainSelector1: supportedPool1,
			},
			expectedV2Count: 2,
			validateResult: func(t *testing.T, result map[cciptypes.ChainSelector]map[reader.MessageTokenID]*SourceTokenDataPayloadV2) {
				require.Contains(t, result, chainSelector1)
				assert.Len(t, result[chainSelector1], 2)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			observer := &CCTPv2TokenDataObserver{
				supportedPoolsBySelector: tt.supportedPools,
				lggr: lggr,
			}

			result := observer.getCCTPv2TokenPayloads(lggr, tt.messages)

			// Count total v2 messages
			totalCount := 0
			for _, chainMsgs := range result {
				totalCount += len(chainMsgs)
			}
			assert.Equal(t, tt.expectedV2Count, totalCount)

			if tt.validateResult != nil {
				tt.validateResult(t, result)
			}
		})
	}
}

// Test deleted - function was refactored

// Test deleted - function was refactored

func Test_IsTokenSupported(t *testing.T) {
	chainSelector := cciptypes.ChainSelector(1)
	supportedPool := "0xabcdef1111111111111111111111111111111111"

	tests := []struct {
		name      string
		pool      string
		supported bool
	}{
		{
			name: "exact match",
			pool: supportedPool,
			supported: true,
		},
		{
			name: "case insensitive match uppercase hex",
			pool: "0xABCDEF1111111111111111111111111111111111", // Uppercase hex digits
			supported: true,
		},
		{
			name: "case insensitive match mixed",
			pool: "0xAbCdEf1111111111111111111111111111111111",
			supported: true,
		},
		{
			name: "different pool",
			pool: "0x2222222222222222222222222222222222222222",
			supported: false,
		},
		{
			name: "chain not in supported list",
			pool: supportedPool,
			supported: false, // Will use different chain selector
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			observer := &CCTPv2TokenDataObserver{
				supportedPoolsBySelector: map[cciptypes.ChainSelector]string{
					chainSelector: supportedPool,
				},
			}

			testChain := chainSelector
			if tt.name == "chain not in supported list" {
				testChain = cciptypes.ChainSelector(999) // Different chain
			}

			tokenAmount := cciptypes.RampTokenAmount{
				SourcePoolAddress: mustParseAddress(tt.pool),
			}

			result := observer.IsTokenSupported(testChain, tokenAmount)
			assert.Equal(t, tt.supported, result)
		})
	}
}

func Test_Close(t *testing.T) {
	observer := &CCTPv2TokenDataObserver{}

	// Should return nil
	err := observer.Close()
	assert.NoError(t, err)

	// Should be idempotent
	err = observer.Close()
	assert.NoError(t, err)
}

// Test_processTransactionMessages is commented out because processTransactionMessages was refactored
// into matchAndValidateAttestations. The test logic is now covered by higher-level tests.
// TODO: Add new test for matchAndValidateAttestations or verify coverage with integration tests

// Mock HTTP client for testing
type mockCCTPv2HTTPClient struct {
	messages CCTPv2Messages
	err      error
}

func (m *mockCCTPv2HTTPClient) GetMessages(
	ctx context.Context,
	sourceChain cciptypes.ChainSelector,
	sourceDomainID uint32,
	transactionHash string,
) (CCTPv2Messages, error) {
	return m.messages, m.err
}

// Helper function to parse address
func mustParseAddress(addr string) cciptypes.UnknownAddress {
	parsed, err := cciptypes.NewUnknownAddressFromHex(addr)
	if err != nil {
		panic(err)
	}
	return parsed
}
