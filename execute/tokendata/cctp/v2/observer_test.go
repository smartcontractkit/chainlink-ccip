package v2

import (
	"context"
	"errors"
	"fmt"
	"testing"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/internal/mocks"
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

// Test helper functions

// createSourceTokenDataPayload creates a test SourceTokenDataPayload with the specified parameters
func createSourceTokenDataPayload(sourceDomain uint32, destDomain uint32, amount int64) SourceTokenDataPayload {
	return SourceTokenDataPayload{
		SourceDomain:      sourceDomain,
		DestinationDomain: destDomain,
		CCTPVersion:       reader.CctpVersion2,
		Amount:            cciptypes.NewBigIntFromInt64(amount),
		BurnToken:         mustBytes32("0x1111"),
		MintRecipient:     mustBytes32("0x2222"),
	}
}

// createCCTPv2Message creates a test CCTP v2 Message with the specified parameters
func createCCTPv2Message(nonce string, sourceDomain string, destDomain string, amount string) Message {
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

// createCCIPMessage creates a test CCIP Message with the specified sequence number and transaction hash
func createCCIPMessage(seqNum cciptypes.SeqNum, txHash string) cciptypes.Message {
	return cciptypes.Message{
		Header: cciptypes.RampMessageHeader{
			MessageID:      [32]byte{byte(seqNum)},
			SequenceNumber: seqNum,
			TxHash:         txHash,
		},
		TokenAmounts: []cciptypes.RampTokenAmount{},
	}
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
				"123": createCCTPv2Message("123", "1", "2", "1000"),
			},
			sourceTokenDataPayloads: map[cciptypes.SeqNum]map[int]SourceTokenDataPayload{},
			isMatch:                 alwaysMatch,
			expectedResult:          map[cciptypes.SeqNum]map[int]CCTPv2MessageOrError{},
			expectedCCTPMessagesAfter: map[string]Message{
				"123": createCCTPv2Message("123", "1", "2", "1000"),
			},
		},
		{
			name:           "empty cctp messages",
			cctpV2Messages: map[string]Message{},
			sourceTokenDataPayloads: map[cciptypes.SeqNum]map[int]SourceTokenDataPayload{
				1: {
					0: createSourceTokenDataPayload(1, 2, 1000),
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
				"123": createCCTPv2Message("123", "1", "2", "1000"),
			},
			sourceTokenDataPayloads: map[cciptypes.SeqNum]map[int]SourceTokenDataPayload{
				1: {
					0: createSourceTokenDataPayload(1, 2, 1000),
				},
			},
			isMatch: realMatch,
			expectedResult: map[cciptypes.SeqNum]map[int]CCTPv2MessageOrError{
				1: {
					0: CCTPv2MessageOrError{
						message: createCCTPv2Message("123", "1", "2", "1000"),
					},
				},
			},
			expectedCCTPMessagesAfter: map[string]Message{}, // message should be removed after matching
		},
		{
			name: "multiple matches in same sequence",
			cctpV2Messages: map[string]Message{
				"123": createCCTPv2Message("123", "1", "2", "1000"),
				"456": createCCTPv2Message("456", "1", "2", "2000"),
			},
			sourceTokenDataPayloads: map[cciptypes.SeqNum]map[int]SourceTokenDataPayload{
				1: {
					0: createSourceTokenDataPayload(1, 2, 1000),
					1: createSourceTokenDataPayload(1, 2, 2000),
				},
			},
			isMatch: realMatch,
			expectedResult: map[cciptypes.SeqNum]map[int]CCTPv2MessageOrError{
				1: {
					0: CCTPv2MessageOrError{
						message: createCCTPv2Message("123", "1", "2", "1000"),
					},
					1: CCTPv2MessageOrError{
						message: createCCTPv2Message("456", "1", "2", "2000"),
					},
				},
			},
			expectedCCTPMessagesAfter: map[string]Message{}, // all messages should be removed
		},
		{
			name: "multiple sequences with matches",
			cctpV2Messages: map[string]Message{
				"123": createCCTPv2Message("123", "1", "2", "1000"),
				"456": createCCTPv2Message("456", "3", "4", "2000"),
			},
			sourceTokenDataPayloads: map[cciptypes.SeqNum]map[int]SourceTokenDataPayload{
				1: {
					0: createSourceTokenDataPayload(1, 2, 1000),
				},
				2: {
					0: createSourceTokenDataPayload(3, 4, 2000),
				},
			},
			isMatch: realMatch,
			expectedResult: map[cciptypes.SeqNum]map[int]CCTPv2MessageOrError{
				1: {
					0: CCTPv2MessageOrError{
						message: createCCTPv2Message("123", "1", "2", "1000"),
					},
				},
				2: {
					0: CCTPv2MessageOrError{
						message: createCCTPv2Message("456", "3", "4", "2000"),
					},
				},
			},
			expectedCCTPMessagesAfter: map[string]Message{}, // all messages should be removed
		},
		{
			name: "partial matches - some payloads have no matching messages",
			cctpV2Messages: map[string]Message{
				"123": createCCTPv2Message("123", "1", "2", "1000"),
			},
			sourceTokenDataPayloads: map[cciptypes.SeqNum]map[int]SourceTokenDataPayload{
				1: {
					0: createSourceTokenDataPayload(1, 2, 1000), // will match
					1: createSourceTokenDataPayload(3, 4, 2000), // no matching message
				},
			},
			isMatch: realMatch,
			expectedResult: map[cciptypes.SeqNum]map[int]CCTPv2MessageOrError{
				1: {
					0: CCTPv2MessageOrError{
						message: createCCTPv2Message("123", "1", "2", "1000"),
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
				"123": createCCTPv2Message("123", "1", "2", "1000"),
				"456": createCCTPv2Message("456", "3", "4", "2000"),
			},
			sourceTokenDataPayloads: map[cciptypes.SeqNum]map[int]SourceTokenDataPayload{
				1: {
					0: createSourceTokenDataPayload(1, 2, 1000),
				},
				2: {
					0: createSourceTokenDataPayload(3, 4, 2000),
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
				"123": createCCTPv2Message("123", "1", "2", "1000"),
				"456": createCCTPv2Message("456", "3", "4", "2000"),
			},
		},
		{
			name: "leftover messages after matching",
			cctpV2Messages: map[string]Message{
				"123": createCCTPv2Message("123", "1", "2", "1000"),
				"456": createCCTPv2Message("456", "3", "4", "2000"),
				"789": createCCTPv2Message("789", "5", "6", "3000"), // no payload for this
			},
			sourceTokenDataPayloads: map[cciptypes.SeqNum]map[int]SourceTokenDataPayload{
				1: {
					0: createSourceTokenDataPayload(1, 2, 1000),
				},
			},
			isMatch: realMatch,
			expectedResult: map[cciptypes.SeqNum]map[int]CCTPv2MessageOrError{
				1: {
					0: CCTPv2MessageOrError{
						message: createCCTPv2Message("123", "1", "2", "1000"),
					},
				},
			},
			expectedCCTPMessagesAfter: map[string]Message{
				"456": createCCTPv2Message("456", "3", "4", "2000"),
				"789": createCCTPv2Message("789", "5", "6", "3000"),
			},
		},
		{
			name: "empty token payload maps in sequence",
			cctpV2Messages: map[string]Message{
				"123": createCCTPv2Message("123", "1", "2", "1000"),
			},
			sourceTokenDataPayloads: map[cciptypes.SeqNum]map[int]SourceTokenDataPayload{
				1: {}, // empty map
				2: {
					0: createSourceTokenDataPayload(1, 2, 1000),
				},
			},
			isMatch: realMatch,
			expectedResult: map[cciptypes.SeqNum]map[int]CCTPv2MessageOrError{
				2: {
					0: CCTPv2MessageOrError{
						message: createCCTPv2Message("123", "1", "2", "1000"),
					},
				},
			},
			expectedCCTPMessagesAfter: map[string]Message{}, // message should be consumed
		},
		{
			name: "first match wins - message removed after first match",
			cctpV2Messages: map[string]Message{
				"123": createCCTPv2Message("123", "1", "2", "1000"),
			},
			sourceTokenDataPayloads: map[cciptypes.SeqNum]map[int]SourceTokenDataPayload{
				1: {
					0: createSourceTokenDataPayload(1, 2, 1000), // will match
				},
				2: {
					0: createSourceTokenDataPayload(1, 2, 1000), // same payload, but message already consumed
				},
			},
			isMatch: realMatch,
			expectedResult: map[cciptypes.SeqNum]map[int]CCTPv2MessageOrError{
				1: {
					0: CCTPv2MessageOrError{
						message: createCCTPv2Message("123", "1", "2", "1000"),
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
					0: createSourceTokenDataPayload(1, 2, 1000),
				},
				2: {
					0: createSourceTokenDataPayload(3, 4, 2000),
				},
			},
			ccipMessages:     map[cciptypes.SeqNum]cciptypes.Message{},
			expectedTxHashes: []string{},
		},
		{
			name: "single sequence with single token",
			sourceTokenDataPayloads: map[cciptypes.SeqNum]map[int]SourceTokenDataPayload{
				1: {
					0: createSourceTokenDataPayload(1, 2, 1000),
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
					0: createSourceTokenDataPayload(1, 2, 1000),
					1: createSourceTokenDataPayload(1, 2, 2000),
					2: createSourceTokenDataPayload(1, 2, 3000),
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
					0: createSourceTokenDataPayload(1, 2, 1000),
				},
				2: {
					0: createSourceTokenDataPayload(3, 4, 2000),
				},
				3: {
					0: createSourceTokenDataPayload(5, 6, 3000),
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
					0: createSourceTokenDataPayload(1, 2, 1000),
					1: createSourceTokenDataPayload(1, 2, 2000),
				},
				2: {
					0: createSourceTokenDataPayload(3, 4, 3000),
					1: createSourceTokenDataPayload(3, 4, 4000),
					2: createSourceTokenDataPayload(3, 4, 5000),
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
					0: createSourceTokenDataPayload(1, 2, 1000),
				},
				2: {
					0: createSourceTokenDataPayload(3, 4, 2000),
				},
				3: {
					0: createSourceTokenDataPayload(5, 6, 3000),
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
					0: createSourceTokenDataPayload(3, 4, 2000),
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
					0: createSourceTokenDataPayload(1, 2, 1000),
				},
				2: {
					0: createSourceTokenDataPayload(3, 4, 2000),
				},
				3: {
					0: createSourceTokenDataPayload(5, 6, 3000),
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
					0: createSourceTokenDataPayload(1, 2, 1000),
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
			name:                    "nil source token data payloads",
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
					0: createSourceTokenDataPayload(1, 2, 1000),
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

// MockCCTPv2AttestationClient is a mock implementation of CCTPv2AttestationClient for testing
type MockCCTPv2AttestationClient struct {
	// responses maps (sourceDomainId, txHash) to a response or error
	responses map[string]MockAttestationResponse
}

// Interface check
var _ interface {
	GetMessages(ctx context.Context, sourceDomainId uint32, txHash string) (Messages, error)
} = (*MockCCTPv2AttestationClient)(nil)

type MockAttestationResponse struct {
	messages Messages
	err      error
}

func NewMockCCTPv2AttestationClient() *MockCCTPv2AttestationClient {
	return &MockCCTPv2AttestationClient{
		responses: make(map[string]MockAttestationResponse),
	}
}

func (m *MockCCTPv2AttestationClient) AddResponse(sourceDomainId uint32, txHash string, messages Messages, err error) {
	key := fmt.Sprintf("%d-%s", sourceDomainId, txHash)
	m.responses[key] = MockAttestationResponse{
		messages: messages,
		err:      err,
	}
}

func (m *MockCCTPv2AttestationClient) GetMessages(ctx context.Context, sourceDomainId uint32, txHash string) (Messages, error) {
	key := fmt.Sprintf("%d-%s", sourceDomainId, txHash)
	if response, exists := m.responses[key]; exists {
		return response.messages, response.err
	}
	// Default to empty response if not configured
	return Messages{}, fmt.Errorf("no response configured for sourceDomainId %d, txHash %s", sourceDomainId, txHash)
}

// getCCTPv2MessagesWithMockClient is a test helper that mirrors getCCTPv2Messages but accepts the mock client
func getCCTPv2MessagesWithMockClient(
	ctx context.Context,
	lggr interface{},
	attestationClient *MockCCTPv2AttestationClient,
	sourceDomainId uint32,
	txHashes mapset.Set[string],
) map[string]Message {
	cctpV2Messages := make(map[string]Message)
	for txHash := range txHashes.Iter() {
		cctpResponse, err := attestationClient.GetMessages(ctx, sourceDomainId, txHash)
		if err != nil {
			// Log error like the real function does, but we're using a null logger
		} else {
			for _, msg := range cctpResponse.Messages {
				cctpV2Messages[msg.EventNonce] = msg
			}
		}
	}
	return cctpV2Messages
}

func TestGetCCTPv2Messages(t *testing.T) {
	ctx := context.Background()
	lggr := mocks.NullLogger

	// Helper to create Messages response
	createMessages := func(messages ...Message) Messages {
		return Messages{Messages: messages}
	}

	tests := []struct {
		name             string
		sourceDomainId   uint32
		txHashes         []string
		setupMock        func(*MockCCTPv2AttestationClient)
		expectedMessages map[string]Message
	}{
		{
			name:             "empty tx hashes",
			sourceDomainId:   1,
			txHashes:         []string{},
			setupMock:        func(m *MockCCTPv2AttestationClient) {},
			expectedMessages: map[string]Message{},
		},
		{
			name:           "single tx hash with single message",
			sourceDomainId: 1,
			txHashes:       []string{"0xabc123"},
			setupMock: func(m *MockCCTPv2AttestationClient) {
				msg := createCCTPv2Message("100", "1", "2", "1000")
				m.AddResponse(1, "0xabc123", createMessages(msg), nil)
			},
			expectedMessages: map[string]Message{
				"100": createCCTPv2Message("100", "1", "2", "1000"),
			},
		},
		{
			name:           "single tx hash with multiple messages",
			sourceDomainId: 1,
			txHashes:       []string{"0xabc123"},
			setupMock: func(m *MockCCTPv2AttestationClient) {
				msg1 := createCCTPv2Message("100", "1", "2", "1000")
				msg2 := createCCTPv2Message("101", "1", "2", "2000")
				m.AddResponse(1, "0xabc123", createMessages(msg1, msg2), nil)
			},
			expectedMessages: map[string]Message{
				"100": createCCTPv2Message("100", "1", "2", "1000"),
				"101": createCCTPv2Message("101", "1", "2", "2000"),
			},
		},
		{
			name:           "multiple tx hashes with single messages each",
			sourceDomainId: 1,
			txHashes:       []string{"0xabc123", "0xdef456"},
			setupMock: func(m *MockCCTPv2AttestationClient) {
				msg1 := createCCTPv2Message("100", "1", "2", "1000")
				msg2 := createCCTPv2Message("200", "1", "2", "2000")
				m.AddResponse(1, "0xabc123", createMessages(msg1), nil)
				m.AddResponse(1, "0xdef456", createMessages(msg2), nil)
			},
			expectedMessages: map[string]Message{
				"100": createCCTPv2Message("100", "1", "2", "1000"),
				"200": createCCTPv2Message("200", "1", "2", "2000"),
			},
		},
		{
			name:           "multiple tx hashes with multiple messages each",
			sourceDomainId: 1,
			txHashes:       []string{"0xabc123", "0xdef456"},
			setupMock: func(m *MockCCTPv2AttestationClient) {
				msg1 := createCCTPv2Message("100", "1", "2", "1000")
				msg2 := createCCTPv2Message("101", "1", "2", "1500")
				msg3 := createCCTPv2Message("200", "1", "2", "2000")
				msg4 := createCCTPv2Message("201", "1", "2", "2500")
				m.AddResponse(1, "0xabc123", createMessages(msg1, msg2), nil)
				m.AddResponse(1, "0xdef456", createMessages(msg3, msg4), nil)
			},
			expectedMessages: map[string]Message{
				"100": createCCTPv2Message("100", "1", "2", "1000"),
				"101": createCCTPv2Message("101", "1", "2", "1500"),
				"200": createCCTPv2Message("200", "1", "2", "2000"),
				"201": createCCTPv2Message("201", "1", "2", "2500"),
			},
		},
		{
			name:           "some tx hashes return errors",
			sourceDomainId: 1,
			txHashes:       []string{"0xabc123", "0xdef456", "0x789xyz"},
			setupMock: func(m *MockCCTPv2AttestationClient) {
				msg1 := createCCTPv2Message("100", "1", "2", "1000")
				m.AddResponse(1, "0xabc123", createMessages(msg1), nil)
				m.AddResponse(1, "0xdef456", Messages{}, fmt.Errorf("network error"))
				msg3 := createCCTPv2Message("300", "1", "2", "3000")
				m.AddResponse(1, "0x789xyz", createMessages(msg3), nil)
			},
			expectedMessages: map[string]Message{
				"100": createCCTPv2Message("100", "1", "2", "1000"),
				"300": createCCTPv2Message("300", "1", "2", "3000"),
			},
		},
		{
			name:           "all tx hashes return errors",
			sourceDomainId: 1,
			txHashes:       []string{"0xabc123", "0xdef456"},
			setupMock: func(m *MockCCTPv2AttestationClient) {
				m.AddResponse(1, "0xabc123", Messages{}, fmt.Errorf("network error"))
				m.AddResponse(1, "0xdef456", Messages{}, fmt.Errorf("timeout error"))
			},
			expectedMessages: map[string]Message{},
		},
		{
			name:           "some tx hashes return empty messages",
			sourceDomainId: 1,
			txHashes:       []string{"0xabc123", "0xdef456", "0x789xyz"},
			setupMock: func(m *MockCCTPv2AttestationClient) {
				msg1 := createCCTPv2Message("100", "1", "2", "1000")
				m.AddResponse(1, "0xabc123", createMessages(msg1), nil)
				m.AddResponse(1, "0xdef456", Messages{Messages: []Message{}}, nil) // Empty messages
				msg3 := createCCTPv2Message("300", "1", "2", "3000")
				m.AddResponse(1, "0x789xyz", createMessages(msg3), nil)
			},
			expectedMessages: map[string]Message{
				"100": createCCTPv2Message("100", "1", "2", "1000"),
				"300": createCCTPv2Message("300", "1", "2", "3000"),
			},
		},
		{
			name:           "duplicate event nonces (last one wins)",
			sourceDomainId: 1,
			txHashes:       []string{"0xabc123", "0xdef456"},
			setupMock: func(m *MockCCTPv2AttestationClient) {
				msg1 := createCCTPv2Message("100", "1", "2", "1000")
				msg2 := createCCTPv2Message("100", "1", "2", "9999") // Same nonce, different amount
				m.AddResponse(1, "0xabc123", createMessages(msg1), nil)
				m.AddResponse(1, "0xdef456", createMessages(msg2), nil)
			},
			expectedMessages: map[string]Message{
				"100": createCCTPv2Message("100", "1", "2", "9999"), // Last one wins
			},
		},
		{
			name:           "different source domain ids",
			sourceDomainId: 42,
			txHashes:       []string{"0xabc123"},
			setupMock: func(m *MockCCTPv2AttestationClient) {
				msg := createCCTPv2Message("100", "42", "2", "1000")
				m.AddResponse(42, "0xabc123", createMessages(msg), nil)
			},
			expectedMessages: map[string]Message{
				"100": createCCTPv2Message("100", "42", "2", "1000"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create mock client and setup responses
			mockClient := NewMockCCTPv2AttestationClient()
			tt.setupMock(mockClient)

			// Create tx hashes set
			txHashesSet := mapset.NewSet[string]()
			for _, txHash := range tt.txHashes {
				txHashesSet.Add(txHash)
			}

			// Call the function with our mock client
			result := getCCTPv2MessagesWithMockClient(ctx, lggr, mockClient, tt.sourceDomainId, txHashesSet)

			// Verify results
			require.Equal(t, len(tt.expectedMessages), len(result), "Result map length mismatch")

			for expectedNonce, expectedMessage := range tt.expectedMessages {
				actualMessage, exists := result[expectedNonce]
				require.True(t, exists, "Expected message with nonce %s not found in result", expectedNonce)
				require.Equal(t, expectedMessage, actualMessage, "Message mismatch for nonce %s", expectedNonce)
			}

			// Verify no unexpected messages
			for actualNonce := range result {
				_, exists := tt.expectedMessages[actualNonce]
				require.True(t, exists, "Unexpected message with nonce %s found in result", actualNonce)
			}
		})
	}
}

// Mock attestation encoder for testing
type mockAttestationEncoder struct {
	responses map[string]mockEncoderResponse
}

type mockEncoderResponse struct {
	data []byte
	err  error
}

func newMockAttestationEncoder() *mockAttestationEncoder {
	return &mockAttestationEncoder{
		responses: make(map[string]mockEncoderResponse),
	}
}

func (m *mockAttestationEncoder) AddResponse(messageHex, attestationHex string, responseData []byte, err error) {
	// Convert hex strings to bytes and back to get the actual key format
	messageBytes, _ := cciptypes.NewBytesFromString(messageHex)
	attestationBytes, _ := cciptypes.NewBytesFromString(attestationHex)
	key := fmt.Sprintf("0x%x-0x%x", messageBytes, attestationBytes)
	m.responses[key] = mockEncoderResponse{
		data: responseData,
		err:  err,
	}
}

func (m *mockAttestationEncoder) Encode(_ context.Context, messageData, attestationData cciptypes.Bytes) (cciptypes.Bytes, error) {
	// Convert bytes back to hex strings for key matching
	messageHex := fmt.Sprintf("0x%x", messageData)
	attestationHex := fmt.Sprintf("0x%x", attestationData)
	key := fmt.Sprintf("%s-%s", messageHex, attestationHex)
	if response, exists := m.responses[key]; exists {
		return response.data, response.err
	}
	// Default to success with concatenated data
	return append(messageData, attestationData...), nil
}

// Helper to create CCIP message with specific number of tokens
func createCCIPMessageWithTokens(seqNum cciptypes.SeqNum, txHash string, tokenCount int) cciptypes.Message {
	tokenAmounts := make([]cciptypes.RampTokenAmount, tokenCount)
	sourcePoolAddress, _ := cciptypes.NewUnknownAddressFromHex("0x1111111111111111111111111111111111111111")
	destTokenAddress, _ := cciptypes.NewUnknownAddressFromHex("0x2222222222222222222222222222222222222222")
	for i := 0; i < tokenCount; i++ {
		tokenAmounts[i] = cciptypes.RampTokenAmount{
			SourcePoolAddress: sourcePoolAddress,
			DestTokenAddress:  destTokenAddress,
			Amount:            cciptypes.NewBigIntFromInt64(1000),
			ExtraData:         []byte("extra"),
		}
	}
	return cciptypes.Message{
		Header: cciptypes.RampMessageHeader{
			MessageID:      [32]byte{byte(seqNum)},
			SequenceNumber: seqNum,
			TxHash:         txHash,
		},
		TokenAmounts: tokenAmounts,
	}
}

func TestConvertCCTPv2MessagesToMessageTokenData(t *testing.T) {
	ctx := context.Background()

	// Helper to create a complete message with valid status
	createCompleteMessage := func(nonce, sourceDomain, destDomain, amount string) Message {
		return Message{
			EventNonce:  nonce,
			CCTPVersion: 2,
			Status:      "complete",
			Message:     "0x1234567890abcdef",
			Attestation: "0xfedcba0987654321",
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

	// Helper to create incomplete message
	createIncompleteMessage := func(nonce, status string) Message {
		msg := createCompleteMessage(nonce, "1", "2", "1000")
		msg.Status = status
		return msg
	}

	// Helper to create message with invalid hex
	createInvalidHexMessage := func(nonce, invalidField string) Message {
		msg := createCompleteMessage(nonce, "1", "2", "1000")
		if invalidField == "message" {
			msg.Message = "0xinvalid-hex"
		} else if invalidField == "attestation" {
			msg.Attestation = "0xinvalid-hex"
		}
		return msg
	}

	tests := []struct {
		name                      string
		ccipMessages              map[cciptypes.SeqNum]cciptypes.Message
		tokenIndexToCCTPv2Message map[cciptypes.SeqNum]map[int]CCTPv2MessageOrError
		setupEncoder              func(*mockAttestationEncoder)
		expectedResults           map[cciptypes.SeqNum]expectedMessageTokenData
	}{
		{
			name:                      "empty inputs",
			ccipMessages:              map[cciptypes.SeqNum]cciptypes.Message{},
			tokenIndexToCCTPv2Message: map[cciptypes.SeqNum]map[int]CCTPv2MessageOrError{},
			setupEncoder:              func(m *mockAttestationEncoder) {},
			expectedResults:           map[cciptypes.SeqNum]expectedMessageTokenData{},
		},
		{
			name: "single message with single token - success",
			ccipMessages: map[cciptypes.SeqNum]cciptypes.Message{
				1: createCCIPMessageWithTokens(1, "0xabc123", 1),
			},
			tokenIndexToCCTPv2Message: map[cciptypes.SeqNum]map[int]CCTPv2MessageOrError{
				1: {
					0: CCTPv2MessageOrError{
						message: createCompleteMessage("100", "1", "2", "1000"),
					},
				},
			},
			setupEncoder: func(m *mockAttestationEncoder) {
				m.AddResponse(
					"0x1234567890abcdef", "0xfedcba0987654321",
					[]byte("success-token-data"), nil,
				)
			},
			expectedResults: map[cciptypes.SeqNum]expectedMessageTokenData{
				1: {
					tokenDataList: []expectedTokenData{
						{
							ready:     true,
							supported: true,
							data:      []byte("success-token-data"),
							err:       nil,
						},
					},
				},
			},
		},
		{
			name: "single message with multiple tokens - mixed success",
			ccipMessages: map[cciptypes.SeqNum]cciptypes.Message{
				1: createCCIPMessageWithTokens(1, "0xabc123", 3),
			},
			tokenIndexToCCTPv2Message: map[cciptypes.SeqNum]map[int]CCTPv2MessageOrError{
				1: {
					0: CCTPv2MessageOrError{
						message: createCompleteMessage("100", "1", "2", "1000"),
					},
					1: CCTPv2MessageOrError{
						err: errors.New("token not found"),
					},
					// token index 2 missing (should be not supported)
				},
			},
			setupEncoder: func(m *mockAttestationEncoder) {
				m.AddResponse(
					"0x1234567890abcdef", "0xfedcba0987654321",
					[]byte("success-token-data"), nil,
				)
			},
			expectedResults: map[cciptypes.SeqNum]expectedMessageTokenData{
				1: {
					tokenDataList: []expectedTokenData{
						{
							ready:     true,
							supported: true,
							data:      []byte("success-token-data"),
							err:       nil,
						},
						{
							ready:     false,
							supported: true,
							data:      nil,
							err:       errors.New("token not found"),
						},
						{
							ready:     false,
							supported: false,
							data:      nil,
							err:       nil,
						},
					},
				},
			},
		},
		{
			name: "multiple messages with various scenarios",
			ccipMessages: map[cciptypes.SeqNum]cciptypes.Message{
				1: createCCIPMessageWithTokens(1, "0xabc123", 2),
				2: createCCIPMessageWithTokens(2, "0xdef456", 1),
				3: createCCIPMessageWithTokens(3, "0x789xyz", 1),
			},
			tokenIndexToCCTPv2Message: map[cciptypes.SeqNum]map[int]CCTPv2MessageOrError{
				1: {
					0: CCTPv2MessageOrError{
						message: createCompleteMessage("100", "1", "2", "1000"),
					},
					1: CCTPv2MessageOrError{
						message: Message{
							EventNonce:  "101",
							CCTPVersion: 2,
							Status:      "complete",
							Message:     "0xabcdef1234567890",
							Attestation: "0x123456789abcdef0",
							DecodedMessage: DecodedMessage{
								SourceDomain:      "1",
								DestinationDomain: "2",
								Nonce:             "101",
								DecodedMessageBody: DecodedMessageBody{
									Amount:        "2000",
									BurnToken:     "0x1111",
									MintRecipient: "0x2222",
								},
							},
						},
					},
				},
				2: {
					0: CCTPv2MessageOrError{
						err: errors.New("api error"),
					},
				},
				// seqNum 3 has no CCTP messages (should be not supported)
			},
			setupEncoder: func(m *mockAttestationEncoder) {
				m.AddResponse(
					"0x1234567890abcdef", "0xfedcba0987654321",
					[]byte("token-data-100"), nil,
				)
				m.AddResponse(
					"0xabcdef1234567890", "0x123456789abcdef0",
					[]byte("token-data-101"), nil,
				)
			},
			expectedResults: map[cciptypes.SeqNum]expectedMessageTokenData{
				1: {
					tokenDataList: []expectedTokenData{
						{
							ready:     true,
							supported: true,
							data:      []byte("token-data-100"),
							err:       nil,
						},
						{
							ready:     true,
							supported: true,
							data:      []byte("token-data-101"),
							err:       nil,
						},
					},
				},
				2: {
					tokenDataList: []expectedTokenData{
						{
							ready:     false,
							supported: true,
							data:      nil,
							err:       errors.New("api error"),
						},
					},
				},
				3: {
					tokenDataList: []expectedTokenData{
						{
							ready:     false,
							supported: false,
							data:      nil,
							err:       nil,
						},
					},
				},
			},
		},
		{
			name: "message with incomplete status",
			ccipMessages: map[cciptypes.SeqNum]cciptypes.Message{
				1: createCCIPMessageWithTokens(1, "0xabc123", 1),
			},
			tokenIndexToCCTPv2Message: map[cciptypes.SeqNum]map[int]CCTPv2MessageOrError{
				1: {
					0: CCTPv2MessageOrError{
						message: createIncompleteMessage("100", "pending"),
					},
				},
			},
			setupEncoder: func(m *mockAttestationEncoder) {},
			expectedResults: map[cciptypes.SeqNum]expectedMessageTokenData{
				1: {
					tokenDataList: []expectedTokenData{
						{
							ready:     false,
							supported: true,
							data:      nil,
							err:       errors.New("A CCTPv2 Message's 'status' is not complete: nonce: 100, sourceDomainId: 1, status: pending"),
						},
					},
				},
			},
		},
		{
			name: "message with invalid hex in message field",
			ccipMessages: map[cciptypes.SeqNum]cciptypes.Message{
				1: createCCIPMessageWithTokens(1, "0xabc123", 1),
			},
			tokenIndexToCCTPv2Message: map[cciptypes.SeqNum]map[int]CCTPv2MessageOrError{
				1: {
					0: CCTPv2MessageOrError{
						message: createInvalidHexMessage("100", "message"),
					},
				},
			},
			setupEncoder: func(m *mockAttestationEncoder) {},
			expectedResults: map[cciptypes.SeqNum]expectedMessageTokenData{
				1: {
					tokenDataList: []expectedTokenData{
						{
							ready:     false,
							supported: true,
							data:      nil,
							err:       errors.New("A CCTPv2 Message's 'message' field could not be converted from string to bytes"),
						},
					},
				},
			},
		},
		{
			name: "message with invalid hex in attestation field",
			ccipMessages: map[cciptypes.SeqNum]cciptypes.Message{
				1: createCCIPMessageWithTokens(1, "0xabc123", 1),
			},
			tokenIndexToCCTPv2Message: map[cciptypes.SeqNum]map[int]CCTPv2MessageOrError{
				1: {
					0: CCTPv2MessageOrError{
						message: createInvalidHexMessage("100", "attestation"),
					},
				},
			},
			setupEncoder: func(m *mockAttestationEncoder) {},
			expectedResults: map[cciptypes.SeqNum]expectedMessageTokenData{
				1: {
					tokenDataList: []expectedTokenData{
						{
							ready:     false,
							supported: true,
							data:      nil,
							err:       errors.New("A CCTPv2 Message's 'attestation' field could not be converted from string to bytes"),
						},
					},
				},
			},
		},
		{
			name: "attestation encoder error",
			ccipMessages: map[cciptypes.SeqNum]cciptypes.Message{
				1: createCCIPMessageWithTokens(1, "0xabc123", 1),
			},
			tokenIndexToCCTPv2Message: map[cciptypes.SeqNum]map[int]CCTPv2MessageOrError{
				1: {
					0: CCTPv2MessageOrError{
						message: createCompleteMessage("100", "1", "2", "1000"),
					},
				},
			},
			setupEncoder: func(m *mockAttestationEncoder) {
				m.AddResponse(
					"0x1234567890abcdef", "0xfedcba0987654321",
					nil, errors.New("encoding failed"),
				)
			},
			expectedResults: map[cciptypes.SeqNum]expectedMessageTokenData{
				1: {
					tokenDataList: []expectedTokenData{
						{
							ready:     false,
							supported: true,
							data:      nil,
							err:       errors.New("attestationEncoder failed for a CCTPv2 message: nonce: 100, sourceDomainId: 1, error: encoding failed"),
						},
					},
				},
			},
		},
		{
			name: "message with no tokens",
			ccipMessages: map[cciptypes.SeqNum]cciptypes.Message{
				1: createCCIPMessageWithTokens(1, "0xabc123", 0),
			},
			tokenIndexToCCTPv2Message: map[cciptypes.SeqNum]map[int]CCTPv2MessageOrError{
				1: {},
			},
			setupEncoder: func(m *mockAttestationEncoder) {},
			expectedResults: map[cciptypes.SeqNum]expectedMessageTokenData{
				1: {
					tokenDataList: []expectedTokenData{},
				},
			},
		},
		{
			name: "mixed scenarios with edge cases",
			ccipMessages: map[cciptypes.SeqNum]cciptypes.Message{
				1: createCCIPMessageWithTokens(1, "0xabc123", 4),
				2: createCCIPMessageWithTokens(2, "0xdef456", 0),
			},
			tokenIndexToCCTPv2Message: map[cciptypes.SeqNum]map[int]CCTPv2MessageOrError{
				1: {
					0: CCTPv2MessageOrError{
						message: createCompleteMessage("100", "1", "2", "1000"),
					},
					1: CCTPv2MessageOrError{
						message: createIncompleteMessage("101", "failed"),
					},
					2: CCTPv2MessageOrError{
						err: errors.New("network timeout"),
					},
					// index 3 missing - should be not supported
				},
				2: {}, // no tokens, so empty map is fine
			},
			setupEncoder: func(m *mockAttestationEncoder) {
				m.AddResponse(
					"0x1234567890abcdef", "0xfedcba0987654321",
					[]byte("successful-encoding"), nil,
				)
			},
			expectedResults: map[cciptypes.SeqNum]expectedMessageTokenData{
				1: {
					tokenDataList: []expectedTokenData{
						{
							ready:     true,
							supported: true,
							data:      []byte("successful-encoding"),
							err:       nil,
						},
						{
							ready:     false,
							supported: true,
							data:      nil,
							err:       errors.New("A CCTPv2 Message's 'status' is not complete: nonce: 101, sourceDomainId: 1, status: failed"),
						},
						{
							ready:     false,
							supported: true,
							data:      nil,
							err:       errors.New("network timeout"),
						},
						{
							ready:     false,
							supported: false,
							data:      nil,
							err:       nil,
						},
					},
				},
				2: {
					tokenDataList: []expectedTokenData{},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create mock encoder
			mockEncoder := newMockAttestationEncoder()
			tt.setupEncoder(mockEncoder)

			// Call the function
			result := convertCCTPv2MessagesToMessageTokenData(
				ctx,
				tt.ccipMessages,
				tt.tokenIndexToCCTPv2Message,
				mockEncoder.Encode,
			)

			// Verify results
			require.Equal(t, len(tt.expectedResults), len(result), "Result map length mismatch")

			for expectedSeqNum, expectedMessageTokenData := range tt.expectedResults {
				actualMessageTokenData, exists := result[expectedSeqNum]
				require.True(t, exists, "Expected sequence number %d not found in result", expectedSeqNum)

				// Verify token data list length
				require.Equal(t, len(expectedMessageTokenData.tokenDataList), len(actualMessageTokenData.TokenData),
					"Token data list length mismatch for seqNum %d", expectedSeqNum)

				// Verify each token data
				for i, expectedTokenData := range expectedMessageTokenData.tokenDataList {
					actualTokenData := actualMessageTokenData.TokenData[i]

					require.Equal(t, expectedTokenData.ready, actualTokenData.Ready,
						"Ready mismatch for seqNum %d, token %d", expectedSeqNum, i)
					require.Equal(t, expectedTokenData.supported, actualTokenData.Supported,
						"Supported mismatch for seqNum %d, token %d", expectedSeqNum, i)
					require.Equal(t, expectedTokenData.data, []byte(actualTokenData.Data),
						"Data mismatch for seqNum %d, token %d", expectedSeqNum, i)

					// Check error - either both nil or both have error with matching message
					if expectedTokenData.err != nil {
						require.Error(t, actualTokenData.Error, "Expected error for seqNum %d, token %d", expectedSeqNum, i)
						require.Contains(t, actualTokenData.Error.Error(), expectedTokenData.err.Error(),
							"Error message mismatch for seqNum %d, token %d", expectedSeqNum, i)
					} else {
						require.NoError(t, actualTokenData.Error, "Expected no error for seqNum %d, token %d", expectedSeqNum, i)
					}
				}
			}

			// Verify no unexpected results
			for actualSeqNum := range result {
				_, exists := tt.expectedResults[actualSeqNum]
				require.True(t, exists, "Unexpected sequence number %d found in result", actualSeqNum)
			}
		})
	}
}

// Helper structs for test expectations
type expectedMessageTokenData struct {
	tokenDataList []expectedTokenData
}

type expectedTokenData struct {
	ready     bool
	supported bool
	data      []byte
	err       error
}

func TestNotSupportedMessageTokenData(t *testing.T) {
	tests := []struct {
		name         string
		ccipMessages map[cciptypes.SeqNum]cciptypes.Message
		expected     map[cciptypes.SeqNum]expectedMessageTokenData
	}{
		{
			name:         "empty input",
			ccipMessages: map[cciptypes.SeqNum]cciptypes.Message{},
			expected:     map[cciptypes.SeqNum]expectedMessageTokenData{},
		},
		{
			name: "single message with no tokens",
			ccipMessages: map[cciptypes.SeqNum]cciptypes.Message{
				1: createCCIPMessageWithTokens(1, "0xabc123", 0),
			},
			expected: map[cciptypes.SeqNum]expectedMessageTokenData{
				1: {
					tokenDataList: []expectedTokenData{},
				},
			},
		},
		{
			name: "single message with single token",
			ccipMessages: map[cciptypes.SeqNum]cciptypes.Message{
				1: createCCIPMessageWithTokens(1, "0xabc123", 1),
			},
			expected: map[cciptypes.SeqNum]expectedMessageTokenData{
				1: {
					tokenDataList: []expectedTokenData{
						{
							ready:     false,
							supported: false,
							data:      nil,
							err:       nil,
						},
					},
				},
			},
		},
		{
			name: "single message with multiple tokens",
			ccipMessages: map[cciptypes.SeqNum]cciptypes.Message{
				1: createCCIPMessageWithTokens(1, "0xabc123", 3),
			},
			expected: map[cciptypes.SeqNum]expectedMessageTokenData{
				1: {
					tokenDataList: make([]expectedTokenData, 3),
				},
			},
		},
		{
			name: "multiple messages with varying token counts",
			ccipMessages: map[cciptypes.SeqNum]cciptypes.Message{
				1: createCCIPMessageWithTokens(1, "0xabc123", 2),
				2: createCCIPMessageWithTokens(2, "0xdef456", 0),
				3: createCCIPMessageWithTokens(3, "0x789xyz", 1),
				5: createCCIPMessageWithTokens(5, "0x111222", 4),
			},
			expected: map[cciptypes.SeqNum]expectedMessageTokenData{
				1: {
					tokenDataList: make([]expectedTokenData, 2),
				},
				2: {
					tokenDataList: []expectedTokenData{},
				},
				3: {
					tokenDataList: make([]expectedTokenData, 1),
				},
				5: {
					tokenDataList: make([]expectedTokenData, 4),
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := notSupportedMessageTokenData(tt.ccipMessages)

			// Verify result map length
			require.Equal(t, len(tt.expected), len(result), "Result map length mismatch")

			for expectedSeqNum, expectedMessageTokenData := range tt.expected {
				actualMessageTokenData, exists := result[expectedSeqNum]
				require.True(t, exists, "Expected sequence number %d not found in result", expectedSeqNum)

				// Verify token data list length
				require.Equal(t, len(expectedMessageTokenData.tokenDataList), len(actualMessageTokenData.TokenData),
					"Token data list length mismatch for seqNum %d", expectedSeqNum)

				// Verify each token data
				for i, expectedTokenData := range expectedMessageTokenData.tokenDataList {
					actualTokenData := actualMessageTokenData.TokenData[i]

					require.Equal(t, expectedTokenData.ready, actualTokenData.Ready,
						"Ready mismatch for seqNum %d, token %d", expectedSeqNum, i)
					require.Equal(t, expectedTokenData.supported, actualTokenData.Supported,
						"Supported mismatch for seqNum %d, token %d", expectedSeqNum, i)
					require.Equal(t, expectedTokenData.data, []byte(actualTokenData.Data),
						"Data mismatch for seqNum %d, token %d", expectedSeqNum, i)
					require.NoError(t, actualTokenData.Error, "Expected no error for seqNum %d, token %d", expectedSeqNum, i)
				}
			}

			// Verify no unexpected results
			for actualSeqNum := range result {
				_, exists := tt.expected[actualSeqNum]
				require.True(t, exists, "Unexpected sequence number %d found in result", actualSeqNum)
			}
		})
	}
}

func TestErrorMessageTokenData(t *testing.T) {
	testError := errors.New("test error message")
	networkError := errors.New("network connection failed")

	tests := []struct {
		name                    string
		err                     error
		ccipMessages            map[cciptypes.SeqNum]cciptypes.Message
		sourceTokenDataPayloads map[cciptypes.SeqNum]map[int]SourceTokenDataPayload
		expected                map[cciptypes.SeqNum]expectedMessageTokenData
	}{
		{
			name:                    "empty inputs",
			err:                     testError,
			ccipMessages:            map[cciptypes.SeqNum]cciptypes.Message{},
			sourceTokenDataPayloads: map[cciptypes.SeqNum]map[int]SourceTokenDataPayload{},
			expected:                map[cciptypes.SeqNum]expectedMessageTokenData{},
		},
		{
			name: "ccip messages but no source token data payloads",
			err:  testError,
			ccipMessages: map[cciptypes.SeqNum]cciptypes.Message{
				1: createCCIPMessageWithTokens(1, "0xabc123", 2),
				2: createCCIPMessageWithTokens(2, "0xdef456", 1),
			},
			sourceTokenDataPayloads: map[cciptypes.SeqNum]map[int]SourceTokenDataPayload{},
			expected: map[cciptypes.SeqNum]expectedMessageTokenData{
				1: {
					tokenDataList: []expectedTokenData{
						{
							ready:     false,
							supported: false,
							data:      nil,
							err:       nil,
						},
						{
							ready:     false,
							supported: false,
							data:      nil,
							err:       nil,
						},
					},
				},
				2: {
					tokenDataList: []expectedTokenData{
						{
							ready:     false,
							supported: false,
							data:      nil,
							err:       nil,
						},
					},
				},
			},
		},
		{
			name: "single message with single token - error applied",
			err:  testError,
			ccipMessages: map[cciptypes.SeqNum]cciptypes.Message{
				1: createCCIPMessageWithTokens(1, "0xabc123", 2),
			},
			sourceTokenDataPayloads: map[cciptypes.SeqNum]map[int]SourceTokenDataPayload{
				1: {
					0: createSourceTokenDataPayload(1, 2, 1000),
				},
			},
			expected: map[cciptypes.SeqNum]expectedMessageTokenData{
				1: {
					tokenDataList: []expectedTokenData{
						{
							ready:     false,
							supported: true,
							data:      nil,
							err:       testError,
						},
						{
							ready:     false,
							supported: false,
							data:      nil,
							err:       nil,
						},
					},
				},
			},
		},
		{
			name: "single message with multiple tokens - partial error coverage",
			err:  networkError,
			ccipMessages: map[cciptypes.SeqNum]cciptypes.Message{
				1: createCCIPMessageWithTokens(1, "0xabc123", 4),
			},
			sourceTokenDataPayloads: map[cciptypes.SeqNum]map[int]SourceTokenDataPayload{
				1: {
					1: createSourceTokenDataPayload(1, 2, 2000),
					2: createSourceTokenDataPayload(1, 2, 3000),
				},
			},
			expected: map[cciptypes.SeqNum]expectedMessageTokenData{
				1: {
					tokenDataList: []expectedTokenData{
						{
							ready:     false,
							supported: false,
							data:      nil,
							err:       nil,
						},
						{
							ready:     false,
							supported: true,
							data:      nil,
							err:       networkError,
						},
						{
							ready:     false,
							supported: true,
							data:      nil,
							err:       networkError,
						},
						{
							ready:     false,
							supported: false,
							data:      nil,
							err:       nil,
						},
					},
				},
			},
		},
		{
			name: "multiple messages with mixed error scenarios",
			err:  testError,
			ccipMessages: map[cciptypes.SeqNum]cciptypes.Message{
				1: createCCIPMessageWithTokens(1, "0xabc123", 3),
				2: createCCIPMessageWithTokens(2, "0xdef456", 1),
				3: createCCIPMessageWithTokens(3, "0x789xyz", 2),
			},
			sourceTokenDataPayloads: map[cciptypes.SeqNum]map[int]SourceTokenDataPayload{
				1: {
					0: createSourceTokenDataPayload(1, 2, 1000),
					2: createSourceTokenDataPayload(1, 2, 3000),
				},
				2: {
					0: createSourceTokenDataPayload(3, 4, 2000),
				},
				// No payloads for seqNum 3
			},
			expected: map[cciptypes.SeqNum]expectedMessageTokenData{
				1: {
					tokenDataList: []expectedTokenData{
						{
							ready:     false,
							supported: true,
							data:      nil,
							err:       testError,
						},
						{
							ready:     false,
							supported: false,
							data:      nil,
							err:       nil,
						},
						{
							ready:     false,
							supported: true,
							data:      nil,
							err:       testError,
						},
					},
				},
				2: {
					tokenDataList: []expectedTokenData{
						{
							ready:     false,
							supported: true,
							data:      nil,
							err:       testError,
						},
					},
				},
				3: {
					tokenDataList: []expectedTokenData{
						{
							ready:     false,
							supported: false,
							data:      nil,
							err:       nil,
						},
						{
							ready:     false,
							supported: false,
							data:      nil,
							err:       nil,
						},
					},
				},
			},
		},
		{
			name: "all tokens have errors",
			err:  networkError,
			ccipMessages: map[cciptypes.SeqNum]cciptypes.Message{
				1: createCCIPMessageWithTokens(1, "0xabc123", 2),
			},
			sourceTokenDataPayloads: map[cciptypes.SeqNum]map[int]SourceTokenDataPayload{
				1: {
					0: createSourceTokenDataPayload(1, 2, 1000),
					1: createSourceTokenDataPayload(1, 2, 2000),
				},
			},
			expected: map[cciptypes.SeqNum]expectedMessageTokenData{
				1: {
					tokenDataList: []expectedTokenData{
						{
							ready:     false,
							supported: true,
							data:      nil,
							err:       networkError,
						},
						{
							ready:     false,
							supported: true,
							data:      nil,
							err:       networkError,
						},
					},
				},
			},
		},
		{
			name: "message with no tokens but has source token data payloads",
			err:  testError,
			ccipMessages: map[cciptypes.SeqNum]cciptypes.Message{
				1: createCCIPMessageWithTokens(1, "0xabc123", 0),
			},
			sourceTokenDataPayloads: map[cciptypes.SeqNum]map[int]SourceTokenDataPayload{
				1: {
					0: createSourceTokenDataPayload(1, 2, 1000),
				},
			},
			expected: map[cciptypes.SeqNum]expectedMessageTokenData{
				1: {
					tokenDataList: []expectedTokenData{},
				},
			},
		},
		{
			name: "token index gaps in source token data payloads",
			err:  testError,
			ccipMessages: map[cciptypes.SeqNum]cciptypes.Message{
				1: createCCIPMessageWithTokens(1, "0xabc123", 5),
			},
			sourceTokenDataPayloads: map[cciptypes.SeqNum]map[int]SourceTokenDataPayload{
				1: {
					0: createSourceTokenDataPayload(1, 2, 1000),
					2: createSourceTokenDataPayload(1, 2, 3000),
					4: createSourceTokenDataPayload(1, 2, 5000),
				},
			},
			expected: map[cciptypes.SeqNum]expectedMessageTokenData{
				1: {
					tokenDataList: []expectedTokenData{
						{
							ready:     false,
							supported: true,
							data:      nil,
							err:       testError,
						},
						{
							ready:     false,
							supported: false,
							data:      nil,
							err:       nil,
						},
						{
							ready:     false,
							supported: true,
							data:      nil,
							err:       testError,
						},
						{
							ready:     false,
							supported: false,
							data:      nil,
							err:       nil,
						},
						{
							ready:     false,
							supported: true,
							data:      nil,
							err:       testError,
						},
					},
				},
			},
		},
		{
			name: "nil error",
			err:  nil,
			ccipMessages: map[cciptypes.SeqNum]cciptypes.Message{
				1: createCCIPMessageWithTokens(1, "0xabc123", 1),
			},
			sourceTokenDataPayloads: map[cciptypes.SeqNum]map[int]SourceTokenDataPayload{
				1: {
					0: createSourceTokenDataPayload(1, 2, 1000),
				},
			},
			expected: map[cciptypes.SeqNum]expectedMessageTokenData{
				1: {
					tokenDataList: []expectedTokenData{
						{
							ready:     false,
							supported: true,
							data:      nil,
							err:       nil,
						},
					},
				},
			},
		},
		{
			name: "source token data payloads but no ccip messages",
			err:  testError,
			ccipMessages: map[cciptypes.SeqNum]cciptypes.Message{},
			sourceTokenDataPayloads: map[cciptypes.SeqNum]map[int]SourceTokenDataPayload{
				1: {
					0: createSourceTokenDataPayload(1, 2, 1000),
					1: createSourceTokenDataPayload(1, 2, 2000),
				},
				2: {
					0: createSourceTokenDataPayload(2, 3, 3000),
				},
			},
			expected: map[cciptypes.SeqNum]expectedMessageTokenData{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := errorMessageTokenData(tt.err, tt.ccipMessages, tt.sourceTokenDataPayloads)

			// Verify result map length
			require.Equal(t, len(tt.expected), len(result), "Result map length mismatch")

			for expectedSeqNum, expectedMessageTokenData := range tt.expected {
				actualMessageTokenData, exists := result[expectedSeqNum]
				require.True(t, exists, "Expected sequence number %d not found in result", expectedSeqNum)

				// Verify token data list length
				require.Equal(t, len(expectedMessageTokenData.tokenDataList), len(actualMessageTokenData.TokenData),
					"Token data list length mismatch for seqNum %d", expectedSeqNum)

				// Verify each token data
				for i, expectedTokenData := range expectedMessageTokenData.tokenDataList {
					actualTokenData := actualMessageTokenData.TokenData[i]

					require.Equal(t, expectedTokenData.ready, actualTokenData.Ready,
						"Ready mismatch for seqNum %d, token %d", expectedSeqNum, i)
					require.Equal(t, expectedTokenData.supported, actualTokenData.Supported,
						"Supported mismatch for seqNum %d, token %d", expectedSeqNum, i)
					require.Equal(t, expectedTokenData.data, []byte(actualTokenData.Data),
						"Data mismatch for seqNum %d, token %d", expectedSeqNum, i)

					// Check error
					if expectedTokenData.err != nil {
						require.Error(t, actualTokenData.Error, "Expected error for seqNum %d, token %d", expectedSeqNum, i)
						require.Equal(t, expectedTokenData.err.Error(), actualTokenData.Error.Error(),
							"Error message mismatch for seqNum %d, token %d", expectedSeqNum, i)
					} else {
						require.NoError(t, actualTokenData.Error, "Expected no error for seqNum %d, token %d", expectedSeqNum, i)
					}
				}
			}

			// Verify no unexpected results
			for actualSeqNum := range result {
				_, exists := tt.expected[actualSeqNum]
				require.True(t, exists, "Unexpected sequence number %d found in result", actualSeqNum)
			}
		})
	}
}
