package v2

import (
	"fmt"
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

func TestMatchCCTPv2MessagesToSourceTokenDataPayloads(t *testing.T) {
	// Helper to create a test Message with a given nonce
	createMessage := func(nonce string, sourceDomain string, destDomain string, amount string) Message {
		return Message{
			EventNonce:  nonce,
			CCTPVersion: 2,
			DecodedMessage: DecodedMessage{
				SourceDomain:      sourceDomain,
				DestinationDomain: destDomain,
				Nonce:             nonce,
				DecodedMessageBody: DecodedMessageBody{
					Amount:        amount,
					BurnToken:     "0x1111",
					MintRecipient: "0x2222",
				},
			},
		}
	}

	// Helper to create a test SourceTokenDataPayload
	createPayload := func(sourceDomain uint32, destDomain uint32, amount int64) SourceTokenDataPayload {
		return SourceTokenDataPayload{
			SourceDomain:      sourceDomain,
			DestinationDomain: destDomain,
			CCTPVersion:       reader.CctpVersion2,
			Amount:            cciptypes.NewBigIntFromInt64(amount),
			BurnToken:         mustBytes32("0x1111"),
			MintRecipient:     mustBytes32("0x2222"),
		}
	}

	// Always match function for testing
	alwaysMatch := func(SourceTokenDataPayload, Message) bool { return true }

	// Never match function for testing
	neverMatch := func(SourceTokenDataPayload, Message) bool { return false }

	// Real match function for testing actual matching
	realMatch := func(payload SourceTokenDataPayload, msg Message) bool {
		return matchesCctpMessage(payload, msg)
	}

	tests := []struct {
		name                      string
		cctpV2Messages            map[string]Message
		sourceTokenDataPayloads   map[cciptypes.SeqNum]map[int]SourceTokenDataPayload
		isMatch                   func(SourceTokenDataPayload, Message) bool
		expectedResult            map[cciptypes.SeqNum]map[int]CCTPv2MessageOrError
		expectedCCTPMessagesAfter map[string]Message // expected state of cctpV2Messages after function call
	}{
		{
			name:                      "empty inputs",
			cctpV2Messages:            map[string]Message{},
			sourceTokenDataPayloads:   map[cciptypes.SeqNum]map[int]SourceTokenDataPayload{},
			isMatch:                   alwaysMatch,
			expectedResult:            map[cciptypes.SeqNum]map[int]CCTPv2MessageOrError{},
			expectedCCTPMessagesAfter: map[string]Message{},
		},
		{
			name: "empty source token data payloads",
			cctpV2Messages: map[string]Message{
				"123": createMessage("123", "1", "2", "1000"),
			},
			sourceTokenDataPayloads: map[cciptypes.SeqNum]map[int]SourceTokenDataPayload{},
			isMatch:                 alwaysMatch,
			expectedResult:          map[cciptypes.SeqNum]map[int]CCTPv2MessageOrError{},
			expectedCCTPMessagesAfter: map[string]Message{
				"123": createMessage("123", "1", "2", "1000"),
			},
		},
		{
			name:           "empty cctp messages",
			cctpV2Messages: map[string]Message{},
			sourceTokenDataPayloads: map[cciptypes.SeqNum]map[int]SourceTokenDataPayload{
				1: {
					0: createPayload(1, 2, 1000),
				},
			},
			isMatch: alwaysMatch,
			expectedResult: map[cciptypes.SeqNum]map[int]CCTPv2MessageOrError{
				1: {
					0: CCTPv2MessageOrError{
						err: fmt.Errorf("no CCTPv2 message found for source token data payload, seqNum: %d, tokenIndex: %d", 1, 0),
					},
				},
			},
			expectedCCTPMessagesAfter: map[string]Message{},
		},
		{
			name: "single match - perfect scenario",
			cctpV2Messages: map[string]Message{
				"123": createMessage("123", "1", "2", "1000"),
			},
			sourceTokenDataPayloads: map[cciptypes.SeqNum]map[int]SourceTokenDataPayload{
				1: {
					0: createPayload(1, 2, 1000),
				},
			},
			isMatch: realMatch,
			expectedResult: map[cciptypes.SeqNum]map[int]CCTPv2MessageOrError{
				1: {
					0: CCTPv2MessageOrError{
						message: createMessage("123", "1", "2", "1000"),
					},
				},
			},
			expectedCCTPMessagesAfter: map[string]Message{}, // message should be removed after matching
		},
		{
			name: "multiple matches in same sequence",
			cctpV2Messages: map[string]Message{
				"123": createMessage("123", "1", "2", "1000"),
				"456": createMessage("456", "1", "2", "2000"),
			},
			sourceTokenDataPayloads: map[cciptypes.SeqNum]map[int]SourceTokenDataPayload{
				1: {
					0: createPayload(1, 2, 1000),
					1: createPayload(1, 2, 2000),
				},
			},
			isMatch: realMatch,
			expectedResult: map[cciptypes.SeqNum]map[int]CCTPv2MessageOrError{
				1: {
					0: CCTPv2MessageOrError{
						message: createMessage("123", "1", "2", "1000"),
					},
					1: CCTPv2MessageOrError{
						message: createMessage("456", "1", "2", "2000"),
					},
				},
			},
			expectedCCTPMessagesAfter: map[string]Message{}, // all messages should be removed
		},
		{
			name: "multiple sequences with matches",
			cctpV2Messages: map[string]Message{
				"123": createMessage("123", "1", "2", "1000"),
				"456": createMessage("456", "3", "4", "2000"),
			},
			sourceTokenDataPayloads: map[cciptypes.SeqNum]map[int]SourceTokenDataPayload{
				1: {
					0: createPayload(1, 2, 1000),
				},
				2: {
					0: createPayload(3, 4, 2000),
				},
			},
			isMatch: realMatch,
			expectedResult: map[cciptypes.SeqNum]map[int]CCTPv2MessageOrError{
				1: {
					0: CCTPv2MessageOrError{
						message: createMessage("123", "1", "2", "1000"),
					},
				},
				2: {
					0: CCTPv2MessageOrError{
						message: createMessage("456", "3", "4", "2000"),
					},
				},
			},
			expectedCCTPMessagesAfter: map[string]Message{}, // all messages should be removed
		},
		{
			name: "partial matches - some payloads have no matching messages",
			cctpV2Messages: map[string]Message{
				"123": createMessage("123", "1", "2", "1000"),
			},
			sourceTokenDataPayloads: map[cciptypes.SeqNum]map[int]SourceTokenDataPayload{
				1: {
					0: createPayload(1, 2, 1000), // will match
					1: createPayload(3, 4, 2000), // no matching message
				},
			},
			isMatch: realMatch,
			expectedResult: map[cciptypes.SeqNum]map[int]CCTPv2MessageOrError{
				1: {
					0: CCTPv2MessageOrError{
						message: createMessage("123", "1", "2", "1000"),
					},
					1: CCTPv2MessageOrError{
						err: fmt.Errorf("no CCTPv2 message found for source token data payload, seqNum: %d, tokenIndex: %d", 1, 1),
					},
				},
			},
			expectedCCTPMessagesAfter: map[string]Message{}, // matched message should be removed
		},
		{
			name: "no matches found - isMatch always returns false",
			cctpV2Messages: map[string]Message{
				"123": createMessage("123", "1", "2", "1000"),
				"456": createMessage("456", "3", "4", "2000"),
			},
			sourceTokenDataPayloads: map[cciptypes.SeqNum]map[int]SourceTokenDataPayload{
				1: {
					0: createPayload(1, 2, 1000),
				},
				2: {
					0: createPayload(3, 4, 2000),
				},
			},
			isMatch: neverMatch,
			expectedResult: map[cciptypes.SeqNum]map[int]CCTPv2MessageOrError{
				1: {
					0: CCTPv2MessageOrError{
						err: fmt.Errorf("no CCTPv2 message found for source token data payload, seqNum: %d, tokenIndex: %d", 1, 0),
					},
				},
				2: {
					0: CCTPv2MessageOrError{
						err: fmt.Errorf("no CCTPv2 message found for source token data payload, seqNum: %d, tokenIndex: %d", 2, 0),
					},
				},
			},
			expectedCCTPMessagesAfter: map[string]Message{
				"123": createMessage("123", "1", "2", "1000"),
				"456": createMessage("456", "3", "4", "2000"),
			},
		},
		{
			name: "leftover messages after matching",
			cctpV2Messages: map[string]Message{
				"123": createMessage("123", "1", "2", "1000"),
				"456": createMessage("456", "3", "4", "2000"),
				"789": createMessage("789", "5", "6", "3000"), // no payload for this
			},
			sourceTokenDataPayloads: map[cciptypes.SeqNum]map[int]SourceTokenDataPayload{
				1: {
					0: createPayload(1, 2, 1000),
				},
			},
			isMatch: realMatch,
			expectedResult: map[cciptypes.SeqNum]map[int]CCTPv2MessageOrError{
				1: {
					0: CCTPv2MessageOrError{
						message: createMessage("123", "1", "2", "1000"),
					},
				},
			},
			expectedCCTPMessagesAfter: map[string]Message{
				"456": createMessage("456", "3", "4", "2000"),
				"789": createMessage("789", "5", "6", "3000"),
			},
		},
		{
			name: "empty token payload maps in sequence",
			cctpV2Messages: map[string]Message{
				"123": createMessage("123", "1", "2", "1000"),
			},
			sourceTokenDataPayloads: map[cciptypes.SeqNum]map[int]SourceTokenDataPayload{
				1: {}, // empty map
				2: {
					0: createPayload(1, 2, 1000),
				},
			},
			isMatch: realMatch,
			expectedResult: map[cciptypes.SeqNum]map[int]CCTPv2MessageOrError{
				2: {
					0: CCTPv2MessageOrError{
						message: createMessage("123", "1", "2", "1000"),
					},
				},
			},
			expectedCCTPMessagesAfter: map[string]Message{}, // message should be consumed
		},
		{
			name: "first match wins - message removed after first match",
			cctpV2Messages: map[string]Message{
				"123": createMessage("123", "1", "2", "1000"),
			},
			sourceTokenDataPayloads: map[cciptypes.SeqNum]map[int]SourceTokenDataPayload{
				1: {
					0: createPayload(1, 2, 1000), // will match
				},
				2: {
					0: createPayload(1, 2, 1000), // same payload, but message already consumed
				},
			},
			isMatch: realMatch,
			expectedResult: map[cciptypes.SeqNum]map[int]CCTPv2MessageOrError{
				1: {
					0: CCTPv2MessageOrError{
						message: createMessage("123", "1", "2", "1000"),
					},
				},
				2: {
					0: CCTPv2MessageOrError{
						err: fmt.Errorf("no CCTPv2 message found for source token data payload, seqNum: %d, tokenIndex: %d", 2, 0),
					},
				},
			},
			expectedCCTPMessagesAfter: map[string]Message{}, // message consumed by first match
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Make a copy of cctpV2Messages to test mutation
			cctpV2MessagesCopy := make(map[string]Message)
			for k, v := range tt.cctpV2Messages {
				cctpV2MessagesCopy[k] = v
			}

			result := matchCCTPv2MessagesToSourceTokenDataPayloads(
				cctpV2MessagesCopy,
				tt.sourceTokenDataPayloads,
				tt.isMatch,
			)

			// Check result structure and content
			require.Equal(t, len(tt.expectedResult), len(result), "Result map length mismatch")

			for expectedSeqNum, expectedTokenMap := range tt.expectedResult {
				actualTokenMap, exists := result[expectedSeqNum]
				require.True(t, exists, "Expected sequence number %d not found in result", expectedSeqNum)
				require.Equal(t, len(expectedTokenMap), len(actualTokenMap), "Token map length mismatch for seqNum %d", expectedSeqNum)

				for expectedTokenIndex, expectedEntry := range expectedTokenMap {
					actualEntry, exists := actualTokenMap[expectedTokenIndex]
					require.True(t, exists, "Expected token index %d not found for seqNum %d", expectedTokenIndex, expectedSeqNum)

					if expectedEntry.err != nil {
						require.Error(t, actualEntry.err, "Expected error for seqNum %d, tokenIndex %d", expectedSeqNum, expectedTokenIndex)
						require.Contains(t, actualEntry.err.Error(), "no CCTPv2 message found", "Error message should indicate no message found")
						require.Equal(t, Message{}, actualEntry.message, "Message should be empty when error is present")
					} else {
						require.NoError(t, actualEntry.err, "Expected no error for seqNum %d, tokenIndex %d", expectedSeqNum, expectedTokenIndex)
						require.Equal(t, expectedEntry.message, actualEntry.message, "Message mismatch for seqNum %d, tokenIndex %d", expectedSeqNum, expectedTokenIndex)
					}
				}
			}

			// Check that cctpV2Messages was properly mutated (messages removed after matching)
			require.Equal(t, len(tt.expectedCCTPMessagesAfter), len(cctpV2MessagesCopy), "CCTP messages map length mismatch after function call")
			for expectedNonce, expectedMessage := range tt.expectedCCTPMessagesAfter {
				actualMessage, exists := cctpV2MessagesCopy[expectedNonce]
				require.True(t, exists, "Expected CCTP message with nonce %s not found after function call", expectedNonce)
				require.Equal(t, expectedMessage, actualMessage, "CCTP message mismatch for nonce %s", expectedNonce)
			}
		})
	}
}

func TestGetTxHashes(t *testing.T) {
	// Helper to create a test CCIP message with a given sequence number and tx hash
	createCCIPMessage := func(seqNum cciptypes.SeqNum, txHash string) cciptypes.Message {
		return cciptypes.Message{
			Header: cciptypes.RampMessageHeader{
				MessageID:         [32]byte{byte(seqNum)},
				SequenceNumber:    seqNum,
				TxHash:            txHash,
			},
			TokenAmounts: []cciptypes.RampTokenAmount{},
		}
	}

	// Helper to create a test source token data payload
	createPayload := func(sourceDomain uint32, destDomain uint32, amount int64) SourceTokenDataPayload {
		return SourceTokenDataPayload{
			SourceDomain:      sourceDomain,
			DestinationDomain: destDomain,
			CCTPVersion:       reader.CctpVersion2,
			Amount:            cciptypes.NewBigIntFromInt64(amount),
			BurnToken:         mustBytes32("0x1111"),
			MintRecipient:     mustBytes32("0x2222"),
		}
	}

	tests := []struct {
		name                    string
		sourceTokenDataPayloads map[cciptypes.SeqNum]map[int]SourceTokenDataPayload
		ccipMessages            map[cciptypes.SeqNum]cciptypes.Message
		expectedTxHashes        []string
	}{
		{
			name:                    "empty inputs",
			sourceTokenDataPayloads: map[cciptypes.SeqNum]map[int]SourceTokenDataPayload{},
			ccipMessages:            map[cciptypes.SeqNum]cciptypes.Message{},
			expectedTxHashes:        []string{},
		},
		{
			name:                    "empty source token data payloads",
			sourceTokenDataPayloads: map[cciptypes.SeqNum]map[int]SourceTokenDataPayload{},
			ccipMessages: map[cciptypes.SeqNum]cciptypes.Message{
				1: createCCIPMessage(1, "0xabc123"),
				2: createCCIPMessage(2, "0xdef456"),
			},
			expectedTxHashes: []string{},
		},
		{
			name: "empty ccip messages",
			sourceTokenDataPayloads: map[cciptypes.SeqNum]map[int]SourceTokenDataPayload{
				1: {
					0: createPayload(1, 2, 1000),
				},
				2: {
					0: createPayload(3, 4, 2000),
				},
			},
			ccipMessages:     map[cciptypes.SeqNum]cciptypes.Message{},
			expectedTxHashes: []string{},
		},
		{
			name: "single sequence with single token",
			sourceTokenDataPayloads: map[cciptypes.SeqNum]map[int]SourceTokenDataPayload{
				1: {
					0: createPayload(1, 2, 1000),
				},
			},
			ccipMessages: map[cciptypes.SeqNum]cciptypes.Message{
				1: createCCIPMessage(1, "0xabc123"),
			},
			expectedTxHashes: []string{"0xabc123"},
		},
		{
			name: "single sequence with multiple tokens",
			sourceTokenDataPayloads: map[cciptypes.SeqNum]map[int]SourceTokenDataPayload{
				1: {
					0: createPayload(1, 2, 1000),
					1: createPayload(1, 2, 2000),
					2: createPayload(1, 2, 3000),
				},
			},
			ccipMessages: map[cciptypes.SeqNum]cciptypes.Message{
				1: createCCIPMessage(1, "0xabc123"),
			},
			expectedTxHashes: []string{"0xabc123"}, // Only one hash even with multiple tokens
		},
		{
			name: "multiple sequences with single tokens each",
			sourceTokenDataPayloads: map[cciptypes.SeqNum]map[int]SourceTokenDataPayload{
				1: {
					0: createPayload(1, 2, 1000),
				},
				2: {
					0: createPayload(3, 4, 2000),
				},
				3: {
					0: createPayload(5, 6, 3000),
				},
			},
			ccipMessages: map[cciptypes.SeqNum]cciptypes.Message{
				1: createCCIPMessage(1, "0xabc123"),
				2: createCCIPMessage(2, "0xdef456"),
				3: createCCIPMessage(3, "0x789xyz"),
			},
			expectedTxHashes: []string{"0xabc123", "0xdef456", "0x789xyz"},
		},
		{
			name: "multiple sequences with multiple tokens each",
			sourceTokenDataPayloads: map[cciptypes.SeqNum]map[int]SourceTokenDataPayload{
				1: {
					0: createPayload(1, 2, 1000),
					1: createPayload(1, 2, 2000),
				},
				2: {
					0: createPayload(3, 4, 3000),
					1: createPayload(3, 4, 4000),
					2: createPayload(3, 4, 5000),
				},
			},
			ccipMessages: map[cciptypes.SeqNum]cciptypes.Message{
				1: createCCIPMessage(1, "0xabc123"),
				2: createCCIPMessage(2, "0xdef456"),
			},
			expectedTxHashes: []string{"0xabc123", "0xdef456"},
		},
		{
			name: "partial matches - some sequences have no corresponding messages",
			sourceTokenDataPayloads: map[cciptypes.SeqNum]map[int]SourceTokenDataPayload{
				1: {
					0: createPayload(1, 2, 1000),
				},
				2: {
					0: createPayload(3, 4, 2000),
				},
				3: {
					0: createPayload(5, 6, 3000),
				},
			},
			ccipMessages: map[cciptypes.SeqNum]cciptypes.Message{
				1: createCCIPMessage(1, "0xabc123"),
				3: createCCIPMessage(3, "0x789xyz"),
				// Missing seqNum 2
			},
			expectedTxHashes: []string{"0xabc123", "0x789xyz"},
		},
		{
			name: "empty token payload maps should be ignored",
			sourceTokenDataPayloads: map[cciptypes.SeqNum]map[int]SourceTokenDataPayload{
				1: {}, // empty map
				2: {
					0: createPayload(3, 4, 2000),
				},
				3: {}, // empty map
			},
			ccipMessages: map[cciptypes.SeqNum]cciptypes.Message{
				1: createCCIPMessage(1, "0xabc123"),
				2: createCCIPMessage(2, "0xdef456"),
				3: createCCIPMessage(3, "0x789xyz"),
			},
			expectedTxHashes: []string{"0xdef456"}, // Only seqNum 2 has non-empty payloads
		},
		{
			name: "duplicate transaction hashes should be deduplicated",
			sourceTokenDataPayloads: map[cciptypes.SeqNum]map[int]SourceTokenDataPayload{
				1: {
					0: createPayload(1, 2, 1000),
				},
				2: {
					0: createPayload(3, 4, 2000),
				},
				3: {
					0: createPayload(5, 6, 3000),
				},
			},
			ccipMessages: map[cciptypes.SeqNum]cciptypes.Message{
				1: createCCIPMessage(1, "0xsametxhash"),
				2: createCCIPMessage(2, "0xsametxhash"), // Same hash as seqNum 1
				3: createCCIPMessage(3, "0xdifferenthash"),
			},
			expectedTxHashes: []string{"0xsametxhash", "0xdifferenthash"}, // Deduplicated
		},
		{
			name: "extra ccip messages without corresponding payloads should be ignored",
			sourceTokenDataPayloads: map[cciptypes.SeqNum]map[int]SourceTokenDataPayload{
				1: {
					0: createPayload(1, 2, 1000),
				},
			},
			ccipMessages: map[cciptypes.SeqNum]cciptypes.Message{
				1: createCCIPMessage(1, "0xabc123"),
				2: createCCIPMessage(2, "0xdef456"), // No payloads for seqNum 2
				3: createCCIPMessage(3, "0x789xyz"), // No payloads for seqNum 3
			},
			expectedTxHashes: []string{"0xabc123"},
		},
		{
			name: "nil source token data payloads",
			sourceTokenDataPayloads: nil,
			ccipMessages: map[cciptypes.SeqNum]cciptypes.Message{
				1: createCCIPMessage(1, "0xabc123"),
			},
			expectedTxHashes: []string{},
		},
		{
			name: "nil ccip messages",
			sourceTokenDataPayloads: map[cciptypes.SeqNum]map[int]SourceTokenDataPayload{
				1: {
					0: createPayload(1, 2, 1000),
				},
			},
			ccipMessages:     nil,
			expectedTxHashes: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := getTxHashes(tt.sourceTokenDataPayloads, tt.ccipMessages)

			// Convert result set to slice for easier comparison
			actualTxHashes := make([]string, 0, result.Cardinality())
			for hash := range result.Iter() {
				actualTxHashes = append(actualTxHashes, hash)
			}

			// Sort both slices for consistent comparison (sets are unordered)
			require.ElementsMatch(t, tt.expectedTxHashes, actualTxHashes, "TX hashes mismatch")
			require.Equal(t, len(tt.expectedTxHashes), result.Cardinality(), "Set cardinality mismatch")
		})
	}
}
