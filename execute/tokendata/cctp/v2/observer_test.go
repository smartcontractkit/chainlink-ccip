package v2

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/pkg/reader"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

const sourceTokenDataPayloadHexV2 = "0x" +
	"000000000000000000000000000000000000000000000000000000000000007b" + // uint64  nonce
	"000000000000000000000000000000000000000000000000000000000000006f" + // uint32  sourceDomain
	"0000000000000000000000000000000000000000000000000000000000000002" + // uint8   cctpVersion (2 for CctpVersion2)
	"00000000000000000000000000000000000000000000000000000000000003e8" + // uint256 amount
	"0000000000000000000000000000000000000000000000000000000012345678" + // uint32  destinationDomain
	"0000000000000000000000001234567890abcdef1234567890abcdef12345678" + // bytes32 mintRecipient
	"2222222222222222222222222222222222222222222222222222222222222222" + // bytes32 burnToken
	"3333333333333333333333333333333333333333333333333333333333333333" + // bytes32 destinationCaller
	"0000000000000000000000000000000000000000000000000000000000000032" + // uint256 maxFee
	"0000000000000000000000000000000000000000000000000000000000000005" // uint32  minFinalityThreshold

var sourceTokenDataPayload = SourceTokenDataPayload{
	Nonce:                123,
	SourceDomain:         111,
	CCTPVersion:          reader.CctpVersion2,
	Amount:               cciptypes.NewBigIntFromInt64(1000),
	DestinationDomain:    0x12345678,
	MintRecipient:        mustBytes32("0x0000000000000000000000001234567890abcdef1234567890abcdef12345678"),
	BurnToken:            mustBytes32("0x2222222222222222222222222222222222222222222222222222222222222222"),
	DestinationCaller:    mustBytes32("0x3333333333333333333333333333333333333333333333333333333333333333"),
	MaxFee:               cciptypes.NewBigIntFromInt64(50),
	MinFinalityThreshold: 5,
}

func TestGetSourceDomainID(t *testing.T) {
	tests := []struct {
		name                            string
		sourceChain                     cciptypes.ChainSelector
		seqNumToSourceTokenDataPayloads map[cciptypes.SeqNum]map[int]SourceTokenDataPayload
		expectedSourceDomainID          uint32
		expectError                     bool
	}{
		{
			name:        "single sequence number with single token - valid",
			sourceChain: cciptypes.ChainSelector(1),
			seqNumToSourceTokenDataPayloads: map[cciptypes.SeqNum]map[int]SourceTokenDataPayload{
				1: {
					0: {SourceDomain: 123},
				},
			},
			expectedSourceDomainID: 123,
			expectError:            false,
		},
		{
			name:        "single sequence number with multiple tokens - same domain",
			sourceChain: cciptypes.ChainSelector(1),
			seqNumToSourceTokenDataPayloads: map[cciptypes.SeqNum]map[int]SourceTokenDataPayload{
				1: {
					0: {SourceDomain: 123},
					1: {SourceDomain: 123},
					2: {SourceDomain: 123},
				},
			},
			expectedSourceDomainID: 123,
			expectError:            false,
		},
		{
			name:        "multiple sequence numbers with single tokens - same domain",
			sourceChain: cciptypes.ChainSelector(1),
			seqNumToSourceTokenDataPayloads: map[cciptypes.SeqNum]map[int]SourceTokenDataPayload{
				1: {
					0: {SourceDomain: 456},
				},
				2: {
					0: {SourceDomain: 456},
				},
				3: {
					0: {SourceDomain: 456},
				},
			},
			expectedSourceDomainID: 456,
			expectError:            false,
		},
		{
			name:        "multiple sequence numbers with multiple tokens - same domain",
			sourceChain: cciptypes.ChainSelector(1),
			seqNumToSourceTokenDataPayloads: map[cciptypes.SeqNum]map[int]SourceTokenDataPayload{
				1: {
					0: {SourceDomain: 789},
					1: {SourceDomain: 789},
				},
				2: {
					0: {SourceDomain: 789},
					1: {SourceDomain: 789},
					2: {SourceDomain: 789},
				},
				3: {
					0: {SourceDomain: 789},
				},
			},
			expectedSourceDomainID: 789,
			expectError:            false,
		},
		{
			name:        "single sequence number with multiple tokens - different domains",
			sourceChain: cciptypes.ChainSelector(1),
			seqNumToSourceTokenDataPayloads: map[cciptypes.SeqNum]map[int]SourceTokenDataPayload{
				1: {
					0: {SourceDomain: 123},
					1: {SourceDomain: 456},
				},
			},
			expectedSourceDomainID: 0,
			expectError:            true,
		},
		{
			name:        "multiple sequence numbers - different domains",
			sourceChain: cciptypes.ChainSelector(1),
			seqNumToSourceTokenDataPayloads: map[cciptypes.SeqNum]map[int]SourceTokenDataPayload{
				1: {
					0: {SourceDomain: 123},
				},
				2: {
					0: {SourceDomain: 456},
				},
			},
			expectedSourceDomainID: 0,
			expectError:            true,
		},
		{
			name:        "mixed tokens with different domains in different sequences",
			sourceChain: cciptypes.ChainSelector(1),
			seqNumToSourceTokenDataPayloads: map[cciptypes.SeqNum]map[int]SourceTokenDataPayload{
				1: {
					0: {SourceDomain: 123},
					1: {SourceDomain: 123},
				},
				2: {
					0: {SourceDomain: 456},
				},
			},
			expectedSourceDomainID: 0,
			expectError:            true,
		},
		{
			name:                            "empty source token data payloads",
			sourceChain:                     cciptypes.ChainSelector(1),
			seqNumToSourceTokenDataPayloads: map[cciptypes.SeqNum]map[int]SourceTokenDataPayload{},
			expectedSourceDomainID:          0,
			expectError:                     true,
		},
		{
			name:                            "nil source token data payloads",
			sourceChain:                     cciptypes.ChainSelector(1),
			seqNumToSourceTokenDataPayloads: nil,
			expectedSourceDomainID:          0,
			expectError:                     true,
		},
		{
			name:        "sequence numbers with empty token maps",
			sourceChain: cciptypes.ChainSelector(1),
			seqNumToSourceTokenDataPayloads: map[cciptypes.SeqNum]map[int]SourceTokenDataPayload{
				1: {},
				2: {},
			},
			expectedSourceDomainID: 0,
			expectError:            true,
		},
		{
			name:        "mixed empty and non-empty token maps",
			sourceChain: cciptypes.ChainSelector(1),
			seqNumToSourceTokenDataPayloads: map[cciptypes.SeqNum]map[int]SourceTokenDataPayload{
				1: {},
				2: {
					0: {SourceDomain: 999},
				},
			},
			expectedSourceDomainID: 999,
			expectError:            false,
		},
		{
			name:        "zero domain ID should be valid",
			sourceChain: cciptypes.ChainSelector(1),
			seqNumToSourceTokenDataPayloads: map[cciptypes.SeqNum]map[int]SourceTokenDataPayload{
				1: {
					0: {SourceDomain: 0},
				},
			},
			expectedSourceDomainID: 0,
			expectError:            false,
		},
		{
			name:        "error occurs at first conflicting sequence number",
			sourceChain: cciptypes.ChainSelector(1),
			seqNumToSourceTokenDataPayloads: map[cciptypes.SeqNum]map[int]SourceTokenDataPayload{
				5: {
					0: {SourceDomain: 100},
				},
				3: {
					0: {SourceDomain: 200}, // This will cause error when processed
				},
				1: {
					0: {SourceDomain: 100},
				},
			},
			expectedSourceDomainID: 0,
			expectError:            true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sourceDomainID, err := getSourceDomainID(tt.sourceChain, tt.seqNumToSourceTokenDataPayloads)

			if tt.expectError {
				require.Error(t, err)
				assert.Equal(t, uint32(0), sourceDomainID)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.expectedSourceDomainID, sourceDomainID)
			}
		})
	}
}

func TestGetSourceTokenDataPayloads(t *testing.T) {
	cctpV2SourcePoolAddress, _ := cciptypes.NewUnknownAddressFromHex("0x1234567890abcdef1234567890abcdef12345678")
	unknownSourcePoolAddress1, _ := cciptypes.NewUnknownAddressFromHex("0x9999999999999999999999999999999999999999")
	unknownSourcePoolAddress2, _ := cciptypes.NewUnknownAddressFromHex("0x8888888888888888888888888888888888888888")

	tests := []struct {
		name                          string
		ccipMessage                   cciptypes.Message
		cctpV2EnabledTokenPoolAddress string
		expectedResults               map[int]SourceTokenDataPayload
		expectedMapSize               int
	}{
		{
			name: "single valid CCTP v2 token",
			ccipMessage: cciptypes.Message{
				TokenAmounts: []cciptypes.RampTokenAmount{
					{
						SourcePoolAddress: cctpV2SourcePoolAddress,
						ExtraData:         mustBytes(sourceTokenDataPayloadHexV2),
					},
				},
			},
			cctpV2EnabledTokenPoolAddress: "0x1234567890abcdef1234567890abcdef12345678",
			expectedResults: map[int]SourceTokenDataPayload{
				0: sourceTokenDataPayload,
			},
			expectedMapSize: 1,
		},
		{
			name: "multiple valid CCTP v2 tokens",
			ccipMessage: cciptypes.Message{
				TokenAmounts: []cciptypes.RampTokenAmount{
					{
						SourcePoolAddress: cctpV2SourcePoolAddress,
						ExtraData:         mustBytes(sourceTokenDataPayloadHexV2),
					},
					{
						SourcePoolAddress: cctpV2SourcePoolAddress,
						ExtraData:         mustBytes(sourceTokenDataPayloadHexV2),
					},
				},
			},
			cctpV2EnabledTokenPoolAddress: "0x1234567890abcdef1234567890abcdef12345678",
			expectedResults: map[int]SourceTokenDataPayload{
				0: sourceTokenDataPayload,
				1: sourceTokenDataPayload,
			},
			expectedMapSize: 2,
		},
		{
			name: "mix of valid and invalid tokens - wrong pool address",
			ccipMessage: cciptypes.Message{
				TokenAmounts: []cciptypes.RampTokenAmount{
					{
						SourcePoolAddress: cctpV2SourcePoolAddress,
						ExtraData:         mustBytes(sourceTokenDataPayloadHexV2),
					},
					{
						SourcePoolAddress: unknownSourcePoolAddress1,
						ExtraData:         mustBytes(sourceTokenDataPayloadHexV2),
					},
					{
						SourcePoolAddress: cctpV2SourcePoolAddress,
						ExtraData:         mustBytes(sourceTokenDataPayloadHexV2),
					},
				},
			},
			cctpV2EnabledTokenPoolAddress: "0x1234567890abcdef1234567890abcdef12345678",
			expectedResults: map[int]SourceTokenDataPayload{
				0: sourceTokenDataPayload,
				2: sourceTokenDataPayload,
			},
			expectedMapSize: 2,
		},
		{
			name: "mix of valid and invalid tokens - invalid ExtraData",
			ccipMessage: cciptypes.Message{
				TokenAmounts: []cciptypes.RampTokenAmount{
					{
						SourcePoolAddress: cctpV2SourcePoolAddress,
						ExtraData:         mustBytes(sourceTokenDataPayloadHexV2),
					},
					{
						SourcePoolAddress: cctpV2SourcePoolAddress,
						ExtraData:         []byte("invalid data"), // Invalid ExtraData
					},
				},
			},
			cctpV2EnabledTokenPoolAddress: "0x1234567890abcdef1234567890abcdef12345678",
			expectedResults: map[int]SourceTokenDataPayload{
				0: sourceTokenDataPayload,
			},
			expectedMapSize: 1, // Only the first token should be included
		},
		{
			name: "empty token amounts",
			ccipMessage: cciptypes.Message{
				TokenAmounts: []cciptypes.RampTokenAmount{},
			},
			cctpV2EnabledTokenPoolAddress: "0x1234567890abcdef1234567890abcdef12345678",
			expectedResults:               map[int]SourceTokenDataPayload{},
			expectedMapSize:               0,
		},
		{
			name: "no valid CCTP v2 tokens - all wrong pool addresses",
			ccipMessage: cciptypes.Message{
				TokenAmounts: []cciptypes.RampTokenAmount{
					{
						SourcePoolAddress: unknownSourcePoolAddress1,
						ExtraData:         mustBytes(sourceTokenDataPayloadHexV2),
					},
					{
						SourcePoolAddress: unknownSourcePoolAddress2,
						ExtraData:         mustBytes(sourceTokenDataPayloadHexV2),
					},
				},
			},
			cctpV2EnabledTokenPoolAddress: "0x1234567890abcdef1234567890abcdef12345678",
			expectedResults:               map[int]SourceTokenDataPayload{},
			expectedMapSize:               0,
		},
		{
			name: "case insensitive pool address matching",
			ccipMessage: cciptypes.Message{
				TokenAmounts: []cciptypes.RampTokenAmount{
					{
						SourcePoolAddress: cctpV2SourcePoolAddress,
						ExtraData:         mustBytes(sourceTokenDataPayloadHexV2),
					},
				},
			},
			cctpV2EnabledTokenPoolAddress: "0x1234567890ABCDEF1234567890ABCDEF12345678", // Uppercase
			expectedResults: map[int]SourceTokenDataPayload{
				0: sourceTokenDataPayload,
			},
			expectedMapSize: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := getSourceTokenDataPayloads(tt.ccipMessage, tt.cctpV2EnabledTokenPoolAddress)

			// Check map size
			assert.Equal(t, tt.expectedMapSize, len(result), "Map size mismatch")

			// Check each expected result
			for expectedIndex, expectedPayload := range tt.expectedResults {
				actualPayload, exists := result[expectedIndex]
				require.True(t, exists, "Expected payload at index %d not found", expectedIndex)

				// Compare all fields
				assert.Equal(t, expectedPayload.Nonce, actualPayload.Nonce, "Nonce mismatch at index %d", expectedIndex)
				assert.Equal(t, expectedPayload.SourceDomain, actualPayload.SourceDomain, "SourceDomain mismatch at index %d", expectedIndex)
				assert.Equal(t, expectedPayload.CCTPVersion, actualPayload.CCTPVersion, "CCTPVersion mismatch at index %d", expectedIndex)
				assert.Equal(t, expectedPayload.Amount.String(), actualPayload.Amount.String(), "Amount mismatch at index %d", expectedIndex)
				assert.Equal(t, expectedPayload.DestinationDomain, actualPayload.DestinationDomain, "DestinationDomain mismatch at index %d", expectedIndex)
				assert.Equal(t, expectedPayload.MintRecipient, actualPayload.MintRecipient, "MintRecipient mismatch at index %d", expectedIndex)
				assert.Equal(t, expectedPayload.BurnToken, actualPayload.BurnToken, "BurnToken mismatch at index %d", expectedIndex)
				assert.Equal(t, expectedPayload.DestinationCaller, actualPayload.DestinationCaller, "DestinationCaller mismatch at index %d", expectedIndex)
				assert.Equal(t, expectedPayload.MaxFee.String(), actualPayload.MaxFee.String(), "MaxFee mismatch at index %d", expectedIndex)
				assert.Equal(t, expectedPayload.MinFinalityThreshold, actualPayload.MinFinalityThreshold, "MinFinalityThreshold mismatch at index %d", expectedIndex)
			}
		})
	}
}
