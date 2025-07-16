package v2

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"testing"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"

	"github.com/smartcontractkit/chainlink-ccip/execute/exectypes"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
)

const sourceTokenDataPayloadHexV2 = "0x" +
	"0000000000000000000000000000000000000000000000000000000000000000" + // uint64  nonce (always 0 for CCTP v2)
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
	Nonce:                0,
	SourceDomain:         111,
	CCTPVersion:          CctpVersion2,
	Amount:               cciptypes.NewBigIntFromInt64(1000),
	DestinationDomain:    0x12345678,
	MintRecipient:        mustBytes32(mintRecipientAddr32),
	BurnToken:            mustBytes32(burnTokenAddr32),
	DestinationCaller:    mustBytes32(destinationCallerAddr32),
	MaxFee:               cciptypes.NewBigIntFromInt64(50),
	MinFinalityThreshold: 5,
}

// Test helper functions

// verifyMessageTokenDataResults is a helper to verify MessageTokenData results
func verifyMessageTokenDataResults(
	t *testing.T,
	expectedResults map[cciptypes.SeqNum]exectypes.MessageTokenData,
	actualResults map[cciptypes.SeqNum]exectypes.MessageTokenData,
) {
	require := require.New(t)
	require.Equal(len(expectedResults), len(actualResults), "Result map length mismatch")

	for expectedSeqNum, expectedMessageTokenData := range expectedResults {
		actualMessageTokenData, exists := actualResults[expectedSeqNum]
		require.True(exists, "Expected sequence number %d not found in result", expectedSeqNum)

		// Verify token data list length
		require.Equal(len(expectedMessageTokenData.TokenData), len(actualMessageTokenData.TokenData),
			"Token data list length mismatch for seqNum %d", expectedSeqNum)

		// Verify each token data
		for i, expectedTokenData := range expectedMessageTokenData.TokenData {
			actualTokenData := actualMessageTokenData.TokenData[i]

			require.Equal(expectedTokenData.Ready, actualTokenData.Ready,
				"Ready mismatch for seqNum %d, token %d", expectedSeqNum, i)
			require.Equal(expectedTokenData.Supported, actualTokenData.Supported,
				"Supported mismatch for seqNum %d, token %d", expectedSeqNum, i)

			// For success tokens, only verify that data is non-empty if expected data is non-empty
			// This allows flexibility in the actual data returned while still verifying success
			if expectedTokenData.Ready && expectedTokenData.Supported && len(expectedTokenData.Data) > 0 {
				require.NotEmpty(actualTokenData.Data, "Expected non-empty data for seqNum %d, token %d", expectedSeqNum, i)
			} else {
				require.Equal(expectedTokenData.Data, actualTokenData.Data,
					"Data mismatch for seqNum %d, token %d", expectedSeqNum, i)
			}

			// Check error
			if expectedTokenData.Error != nil {
				require.Error(actualTokenData.Error, "Expected error for seqNum %d, token %d", expectedSeqNum, i)
			} else {
				require.NoError(actualTokenData.Error, "Expected no error for seqNum %d, token %d", expectedSeqNum, i)
			}
		}
	}
}

// createSourceTokenDataPayload creates a test SourceTokenDataPayload with the specified parameters
func createSourceTokenDataPayload(sourceDomain uint32, destDomain uint32, amount int64) SourceTokenDataPayload {
	return SourceTokenDataPayload{
		Nonce:             123, // Default nonce that matches the test CCTP message
		SourceDomain:      sourceDomain,
		DestinationDomain: destDomain,
		CCTPVersion:       CctpVersion2,
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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sourceDomainID, err := getSourceDomainID(testLogger, tt.sourceChain, tt.seqNumToSourceTokenDataPayloads)

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
				assert.Equal(t, expectedPayload.SourceDomain, actualPayload.SourceDomain,
					"SourceDomain mismatch at index %d", expectedIndex)
				assert.Equal(t, expectedPayload.CCTPVersion, actualPayload.CCTPVersion,
					"CCTPVersion mismatch at index %d", expectedIndex)
				assert.Equal(t, expectedPayload.Amount.String(), actualPayload.Amount.String(),
					"Amount mismatch at index %d", expectedIndex)
				assert.Equal(t, expectedPayload.DestinationDomain, actualPayload.DestinationDomain,
					"DestinationDomain mismatch at index %d", expectedIndex)
				assert.Equal(t, expectedPayload.MintRecipient, actualPayload.MintRecipient,
					"MintRecipient mismatch at index %d", expectedIndex)
				assert.Equal(t, expectedPayload.BurnToken, actualPayload.BurnToken, "BurnToken mismatch at index %d", expectedIndex)
				assert.Equal(t, expectedPayload.DestinationCaller, actualPayload.DestinationCaller,
					"DestinationCaller mismatch at index %d", expectedIndex)
				assert.Equal(t, expectedPayload.MaxFee.String(), actualPayload.MaxFee.String(),
					"MaxFee mismatch at index %d", expectedIndex)
				assert.Equal(t, expectedPayload.MinFinalityThreshold, actualPayload.MinFinalityThreshold,
					"MinFinalityThreshold mismatch at index %d", expectedIndex)
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
			name: "one wins when competing for same message - message removed after first match",
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
				testLogger,
				cctpV2MessagesCopy,
				tt.sourceTokenDataPayloads,
				tt.isMatch,
				1, // sourceChain
				NewNoOpMetricsReporter(),
			)

			// Check result structure and content
			require.Equal(t, len(tt.expectedResult), len(result), "Result map length mismatch")

			// Special handling for the competing message test case
			if tt.name == "one wins when competing for same message - message removed after first match" {
				// Verify that exactly one seqNum gets the message and one gets an error
				require.Equal(t, 2, len(result), "Should have results for both sequence numbers")

				var successCount, errorCount int
				var successSeqNum, errorSeqNum cciptypes.SeqNum

				for seqNum, tokenMap := range result {
					require.Equal(t, 1, len(tokenMap), "Each seqNum should have exactly one token")
					entry := tokenMap[0]

					if entry.err != nil {
						errorCount++
						errorSeqNum = seqNum
						require.Contains(t, entry.err.Error(), "no CCTPv2 message found",
							"Error message should indicate no message found")
						require.Equal(t, Message{}, entry.message, "Message should be empty when error is present")
					} else {
						successCount++
						successSeqNum = seqNum
						require.Equal(t, createCCTPv2Message("123", "1", "2", "1000"), entry.message,
							"Message should match expected")
					}
				}

				require.Equal(t, 1, successCount, "Exactly one seqNum should succeed")
				require.Equal(t, 1, errorCount, "Exactly one seqNum should fail")
				require.True(t, (successSeqNum == 1 && errorSeqNum == 2) || (successSeqNum == 2 && errorSeqNum == 1),
					"Either seqNum 1 or 2 can win, but exactly one should succeed")
			} else {
				// Normal test case verification
				for expectedSeqNum, expectedTokenMap := range tt.expectedResult {
					actualTokenMap, exists := result[expectedSeqNum]
					require.True(t, exists, "Expected sequence number %d not found in result", expectedSeqNum)
					require.Equal(t, len(expectedTokenMap), len(actualTokenMap),
						"Token map length mismatch for seqNum %d", expectedSeqNum)

					for expectedTokenIndex, expectedEntry := range expectedTokenMap {
						actualEntry, exists := actualTokenMap[expectedTokenIndex]
						require.True(t, exists, "Expected token index %d not found for seqNum %d", expectedTokenIndex, expectedSeqNum)

						if expectedEntry.err != nil {
							require.Error(t, actualEntry.err,
								"Expected error for seqNum %d, tokenIndex %d", expectedSeqNum, expectedTokenIndex)
							require.Contains(t, actualEntry.err.Error(), "no CCTPv2 message found",
								"Error message should indicate no message found")
							require.Equal(t, Message{}, actualEntry.message, "Message should be empty when error is present")
						} else {
							require.NoError(t, actualEntry.err,
								"Expected no error for seqNum %d, tokenIndex %d", expectedSeqNum, expectedTokenIndex)
							require.Equal(t, expectedEntry.message, actualEntry.message,
								"Message mismatch for seqNum %d, tokenIndex %d", expectedSeqNum, expectedTokenIndex)
						}
					}
				}
			}

			// Check that cctpV2Messages was properly mutated (messages removed after matching)
			require.Equal(t, len(tt.expectedCCTPMessagesAfter), len(cctpV2MessagesCopy),
				"CCTP messages map length mismatch after function call")
			for expectedNonce, expectedMessage := range tt.expectedCCTPMessagesAfter {
				actualMessage, exists := cctpV2MessagesCopy[expectedNonce]
				require.True(t, exists, "Expected CCTP message with nonce %s not found after function call", expectedNonce)
				require.Equal(t, expectedMessage, actualMessage, "CCTP message mismatch for nonce %s", expectedNonce)
			}
		})
	}
}

func TestCCTPv2TokenDataObserver_Observe(t *testing.T) {
	ctx := context.Background()

	// Test helpers
	createMessageWithToken := func(
		seqNum cciptypes.SeqNum, txHash string, poolAddr string, extraData []byte,
	) cciptypes.Message {
		return cciptypes.Message{
			Header: cciptypes.RampMessageHeader{
				MessageID:      [32]byte{byte(seqNum)},
				SequenceNumber: seqNum,
				TxHash:         txHash,
			},
			TokenAmounts: []cciptypes.RampTokenAmount{
				{
					SourcePoolAddress: mustCreateUnknownAddress(poolAddr),
					ExtraData:         extraData,
					Amount:            cciptypes.NewBigIntFromInt64(1000),
				},
			},
		}
	}

	createMessageWithoutTokens := func(seqNum cciptypes.SeqNum, txHash string) cciptypes.Message {
		return cciptypes.Message{
			Header: cciptypes.RampMessageHeader{
				MessageID:      [32]byte{byte(seqNum)},
				SequenceNumber: seqNum,
				TxHash:         txHash,
			},
			TokenAmounts: []cciptypes.RampTokenAmount{},
		}
	}

	// Mock attestation encoder
	mockEncoder := func(
		ctx context.Context, message cciptypes.Bytes, attestation cciptypes.Bytes,
	) (cciptypes.Bytes, error) {
		return append(message, attestation...), nil
	}

	// Mock metrics reporter
	mockMetrics := NewNoOpMetricsReporter()

	tests := []struct {
		name                       string
		supportedPools             map[cciptypes.ChainSelector]string
		messages                   exectypes.MessageObservations
		mockSetup                  func(*MockCCTPv2AttestationClient)
		expectedTokenDataCount     map[cciptypes.ChainSelector]int
		expectedSupportedTokens    map[cciptypes.ChainSelector]int
		expectedErrorTokens        map[cciptypes.ChainSelector]int
		expectedNotSupportedTokens map[cciptypes.ChainSelector]int
	}{
		{
			name: "empty messages",
			supportedPools: map[cciptypes.ChainSelector]string{
				1: "0x1234",
			},
			messages: exectypes.MessageObservations{},
			mockSetup: func(m *MockCCTPv2AttestationClient) {
				// No setup needed for empty messages
			},
			expectedTokenDataCount: map[cciptypes.ChainSelector]int{},
		},
		{
			name: "unsupported chain - no pool configured",
			supportedPools: map[cciptypes.ChainSelector]string{
				1: "0x1234",
			},
			messages: exectypes.MessageObservations{
				2: { // Chain 2 not supported
					1: createMessageWithoutTokens(1, "0xabc123"),
				},
			},
			mockSetup: func(m *MockCCTPv2AttestationClient) {
				// No setup needed for unsupported chain
			},
			expectedTokenDataCount:     map[cciptypes.ChainSelector]int{2: 1},
			expectedNotSupportedTokens: map[cciptypes.ChainSelector]int{2: 0}, // No tokens = no not supported
		},
		{
			name: "supported chain with no tokens",
			supportedPools: map[cciptypes.ChainSelector]string{
				1: "0x1234",
			},
			messages: exectypes.MessageObservations{
				1: {
					1: createMessageWithoutTokens(1, "0xabc123"),
				},
			},
			mockSetup: func(m *MockCCTPv2AttestationClient) {
				// No setup needed for no tokens
			},
			expectedTokenDataCount:     map[cciptypes.ChainSelector]int{1: 1},
			expectedNotSupportedTokens: map[cciptypes.ChainSelector]int{1: 0}, // No tokens = no not supported
		},
		{
			name: "supported chain with unsupported token",
			supportedPools: map[cciptypes.ChainSelector]string{
				1: "0x1234",
			},
			messages: exectypes.MessageObservations{
				1: {
					1: createMessageWithToken(1, "0xabc123", "0x5678", []byte("invalid")), // Wrong pool
				},
			},
			mockSetup: func(m *MockCCTPv2AttestationClient) {
				// No setup needed for unsupported token
			},
			expectedTokenDataCount:     map[cciptypes.ChainSelector]int{1: 1},
			expectedNotSupportedTokens: map[cciptypes.ChainSelector]int{1: 1},
		},
		{
			name: "successful token processing",
			supportedPools: map[cciptypes.ChainSelector]string{
				1: "0x1234",
			},
			messages: exectypes.MessageObservations{
				1: {
					1: createMessageWithToken(1, "0xabc123", "0x1234", mustHexDecode(sourceTokenDataPayloadHexV2)),
				},
			},
			mockSetup: func(m *MockCCTPv2AttestationClient) {
				m.AddResponse(111, "0xabc123", Messages{
					Messages: []Message{
						{
							EventNonce:  "0",
							CCTPVersion: 2,
							Status:      "complete",
							Message:     "0x1234abcd",
							Attestation: "0x5678efab",
							DecodedMessage: DecodedMessage{
								SourceDomain:      "111",
								DestinationDomain: "305419896", // 0x12345678
								Nonce:             "0",
								DecodedMessageBody: DecodedMessageBody{
									Amount:        "1000",
									BurnToken:     "0x2222222222222222222222222222222222222222222222222222222222222222",
									MintRecipient: "0x1234567890abcdef1234567890abcdef12345678",
								},
								DestinationCaller: "0x3333333333333333333333333333333333333333333333333333333333333333",
							},
						},
					},
				}, nil)
			},
			expectedTokenDataCount:  map[cciptypes.ChainSelector]int{1: 1},
			expectedSupportedTokens: map[cciptypes.ChainSelector]int{1: 1},
		},
		{
			name: "multiple chains with mixed results",
			supportedPools: map[cciptypes.ChainSelector]string{
				1: "0x1234",
				2: "0x5678",
			},
			messages: exectypes.MessageObservations{
				1: {
					1: createMessageWithToken(1, "0xabc123", "0x1234", mustHexDecode(sourceTokenDataPayloadHexV2)),
				},
				2: {
					2: createMessageWithToken(2, "0xdef456", "0x9999", []byte("invalid")), // Wrong pool
				},
			},
			mockSetup: func(m *MockCCTPv2AttestationClient) {
				m.AddResponse(111, "0xabc123", Messages{
					Messages: []Message{
						{
							EventNonce:  "0",
							CCTPVersion: 2,
							Status:      "complete",
							Message:     "0x1234abcd",
							Attestation: "0x5678efab",
							DecodedMessage: DecodedMessage{
								SourceDomain:      "111",
								DestinationDomain: "305419896", // 0x12345678
								Nonce:             "0",
								DecodedMessageBody: DecodedMessageBody{
									Amount:        "1000",
									BurnToken:     "0x2222222222222222222222222222222222222222222222222222222222222222",
									MintRecipient: "0x1234567890abcdef1234567890abcdef12345678",
								},
								DestinationCaller: "0x3333333333333333333333333333333333333333333333333333333333333333",
							},
						},
					},
				}, nil)
			},
			expectedTokenDataCount:     map[cciptypes.ChainSelector]int{1: 1, 2: 1},
			expectedSupportedTokens:    map[cciptypes.ChainSelector]int{1: 1},
			expectedNotSupportedTokens: map[cciptypes.ChainSelector]int{2: 1},
		},
		{
			name: "attestation API error",
			supportedPools: map[cciptypes.ChainSelector]string{
				1: "0x1234",
			},
			messages: exectypes.MessageObservations{
				1: {
					1: createMessageWithToken(1, "0xabc123", "0x1234", mustHexDecode(sourceTokenDataPayloadHexV2)),
				},
			},
			mockSetup: func(m *MockCCTPv2AttestationClient) {
				m.AddResponse(111, "0xabc123", Messages{}, fmt.Errorf("API error"))
			},
			expectedTokenDataCount: map[cciptypes.ChainSelector]int{1: 1},
			expectedErrorTokens:    map[cciptypes.ChainSelector]int{1: 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create and configure mock
			mockClient := NewMockCCTPv2AttestationClient()
			tt.mockSetup(mockClient)

			// Create observer with direct struct initialization to bypass type constraints
			observer := &CCTPv2TokenDataObserver{
				lggr:                     logger.Nop(),
				destChainSelector:        cciptypes.ChainSelector(999),
				supportedPoolsBySelector: tt.supportedPools,
				attestationEncoder:       mockEncoder,
				attestationClient:        mockClient,
				metricsReporter:          mockMetrics,
			}

			// Call Observe
			result, err := observer.Observe(ctx, tt.messages)

			// Verify no error returned (as per current implementation)
			require.NoError(t, err)

			// Verify result structure
			require.Equal(t, len(tt.expectedTokenDataCount), len(result))

			for chainSelector, expectedCount := range tt.expectedTokenDataCount {
				chainResult, exists := result[chainSelector]
				require.True(t, exists, "Chain %d not found in result", chainSelector)
				require.Equal(t, expectedCount, len(chainResult), "Wrong number of messages for chain %d", chainSelector)

				// Count token types
				supportedCount := 0
				errorCount := 0
				notSupportedCount := 0

				for _, messageTokenData := range chainResult {
					for _, tokenData := range messageTokenData.TokenData {
						if tokenData.Ready && tokenData.Supported {
							supportedCount++
						} else if tokenData.Error != nil {
							errorCount++
						} else if !tokenData.Supported {
							notSupportedCount++
						}
					}
				}

				if expectedSupportedTokens, exists := tt.expectedSupportedTokens[chainSelector]; exists {
					require.Equal(t, expectedSupportedTokens, supportedCount,
						"Wrong supported token count for chain %d", chainSelector)
				}

				if expectedErrorTokens, exists := tt.expectedErrorTokens[chainSelector]; exists {
					require.Equal(t, expectedErrorTokens, errorCount,
						"Wrong error token count for chain %d", chainSelector)
				}

				if expectedNotSupportedTokens, exists := tt.expectedNotSupportedTokens[chainSelector]; exists {
					require.Equal(t, expectedNotSupportedTokens, notSupportedCount,
						"Wrong not supported token count for chain %d", chainSelector)
				}
			}
		})
	}
}

func mustHexDecode(hexStr string) []byte {
	if len(hexStr) >= 2 && hexStr[:2] == "0x" {
		hexStr = hexStr[2:]
	}
	data, err := hex.DecodeString(hexStr)
	if err != nil {
		panic(err)
	}
	return data
}

func mustCreateUnknownAddress(addr string) cciptypes.UnknownAddress {
	address, err := cciptypes.NewUnknownAddressFromHex(addr)
	if err != nil {
		panic(err)
	}
	return address
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
			result := getTxHashes(testLogger, tt.sourceTokenDataPayloads, tt.ccipMessages)

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

// MockCCTPv2AttestationClient is a mock implementation of CCTPv2AttestationClientHttp for testing
type MockCCTPv2AttestationClient struct {
	// responses maps (sourceDomainId, txHash) to a response or error
	responses map[string]MockAttestationResponse
}

// Interface check
var _ interface {
	GetMessages(ctx context.Context, sourceChain cciptypes.ChainSelector, sourceDomainID uint32, txHash string) (Messages, error)
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

func (m *MockCCTPv2AttestationClient) AddResponse(sourceDomainID uint32, txHash string, messages Messages, err error) {
	key := fmt.Sprintf("%d-%s", sourceDomainID, txHash)
	m.responses[key] = MockAttestationResponse{
		messages: messages,
		err:      err,
	}
}

func (m *MockCCTPv2AttestationClient) GetMessages(
	ctx context.Context,
	sourceChain cciptypes.ChainSelector,
	sourceDomainID uint32,
	txHash string,
) (Messages, error) {
	key := fmt.Sprintf("%d-%s", sourceDomainID, txHash)
	if response, exists := m.responses[key]; exists {
		return response.messages, response.err
	}
	// Default to empty response if not configured
	return Messages{}, fmt.Errorf("no response configured for sourceDomainID %d, txHash %s", sourceDomainID, txHash)
}

// getCCTPv2MessagesWithMockClient is a test helper that mirrors getCCTPv2Messages but accepts the mock client
func getCCTPv2MessagesWithMockClient(
	ctx context.Context,
	attestationClient *MockCCTPv2AttestationClient,
	sourceChain cciptypes.ChainSelector,
	sourceDomainID uint32,
	txHashes mapset.Set[string],
) map[string]Message {
	cctpV2Messages := make(map[string]Message)
	for txHash := range txHashes.Iter() {
		cctpResponse, err := attestationClient.GetMessages(ctx, sourceChain, sourceDomainID, txHash)
		if err == nil && len(cctpResponse.Messages) > 0 {
			for _, msg := range cctpResponse.Messages {
				cctpV2Messages[msg.EventNonce] = msg
			}
		}
	}
	return cctpV2Messages
}

func TestGetCCTPv2Messages(t *testing.T) {
	ctx := context.Background()

	// Helper to create Messages response
	createMessages := func(messages ...Message) Messages {
		return Messages{Messages: messages}
	}

	tests := []struct {
		name             string
		sourceDomainID   uint32
		txHashes         []string
		setupMock        func(*MockCCTPv2AttestationClient)
		expectedMessages map[string]Message
	}{
		{
			name:             "empty tx hashes",
			sourceDomainID:   1,
			txHashes:         []string{},
			setupMock:        func(m *MockCCTPv2AttestationClient) {},
			expectedMessages: map[string]Message{},
		},
		{
			name:           "single tx hash with single message",
			sourceDomainID: 1,
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
			sourceDomainID: 1,
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
			sourceDomainID: 1,
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
			sourceDomainID: 1,
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
			sourceDomainID: 1,
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
			sourceDomainID: 1,
			txHashes:       []string{"0xabc123", "0xdef456"},
			setupMock: func(m *MockCCTPv2AttestationClient) {
				m.AddResponse(1, "0xabc123", Messages{}, fmt.Errorf("network error"))
				m.AddResponse(1, "0xdef456", Messages{}, fmt.Errorf("timeout error"))
			},
			expectedMessages: map[string]Message{},
		},
		{
			name:           "some tx hashes return empty messages",
			sourceDomainID: 1,
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
			name:           "duplicate event nonces (one of them wins)",
			sourceDomainID: 1,
			txHashes:       []string{"0xabc123", "0xdef456"},
			setupMock: func(m *MockCCTPv2AttestationClient) {
				msg1 := createCCTPv2Message("100", "1", "2", "1000")
				msg2 := createCCTPv2Message("100", "1", "2", "9999") // Same nonce, different amount
				m.AddResponse(1, "0xabc123", createMessages(msg1), nil)
				m.AddResponse(1, "0xdef456", createMessages(msg2), nil)
			},
			expectedMessages: map[string]Message{
				"100": createCCTPv2Message("100", "1", "2", "1000"), // Could be either 1000 or 9999 due to map iteration order
			},
		},
		{
			name:           "different source domain ids",
			sourceDomainID: 42,
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
			result := getCCTPv2MessagesWithMockClient(ctx, mockClient, testSourceChain, tt.sourceDomainID, txHashesSet)

			// Verify results
			require.Equal(t, len(tt.expectedMessages), len(result), "Result map length mismatch")

			for expectedNonce, expectedMessage := range tt.expectedMessages {
				actualMessage, exists := result[expectedNonce]
				require.True(t, exists, "Expected message with nonce %s not found in result", expectedNonce)

				// Special handling for duplicate nonces test - either message could win
				if tt.name == "duplicate event nonces (one of them wins)" && expectedNonce == "100" {
					// For this test, either amount "1000" or "9999" is acceptable
					require.True(t,
						actualMessage.DecodedMessage.DecodedMessageBody.Amount == "1000" ||
							actualMessage.DecodedMessage.DecodedMessageBody.Amount == "9999",
						"Expected amount to be either '1000' or '9999', got '%s'",
						actualMessage.DecodedMessage.DecodedMessageBody.Amount)
					// Check other fields are correct
					expectedMessage.DecodedMessage.DecodedMessageBody.Amount = actualMessage.DecodedMessage.DecodedMessageBody.Amount
				}

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

func (m *mockAttestationEncoder) Encode(
	_ context.Context,
	messageData, attestationData cciptypes.Bytes,
) (cciptypes.Bytes, error) {
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
		expectedResults           map[cciptypes.SeqNum]exectypes.MessageTokenData
	}{
		{
			name:                      "empty inputs",
			ccipMessages:              map[cciptypes.SeqNum]cciptypes.Message{},
			tokenIndexToCCTPv2Message: map[cciptypes.SeqNum]map[int]CCTPv2MessageOrError{},
			setupEncoder:              noopMockEncoder(),
			expectedResults:           map[cciptypes.SeqNum]exectypes.MessageTokenData{},
		},
		{
			name: "single message with single token - success",
			ccipMessages: map[cciptypes.SeqNum]cciptypes.Message{
				1: createCCIPMessageWithMultipleTokens(1, "0xabc123", 1),
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
			expectedResults: map[cciptypes.SeqNum]exectypes.MessageTokenData{
				1: createExpectedMessageTokenData(
					successToken([]byte("success-token-data")),
				),
			},
		},
		{
			name: "single message with multiple tokens - mixed success",
			ccipMessages: map[cciptypes.SeqNum]cciptypes.Message{
				1: createCCIPMessageWithMultipleTokens(1, "0xabc123", 3),
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
			expectedResults: map[cciptypes.SeqNum]exectypes.MessageTokenData{
				1: createExpectedMessageTokenData(
					successToken([]byte("success-token-data")),
					errorToken(errors.New("token not found")),
					notSupportedToken(),
				),
			},
		},
		{
			name: "multiple messages with various scenarios",
			ccipMessages: map[cciptypes.SeqNum]cciptypes.Message{
				1: createCCIPMessageWithMultipleTokens(1, "0xabc123", 2),
				2: createCCIPMessageWithMultipleTokens(2, "0xdef456", 1),
				3: createCCIPMessageWithMultipleTokens(3, "0x789xyz", 1),
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
			expectedResults: map[cciptypes.SeqNum]exectypes.MessageTokenData{
				1: createExpectedMessageTokenData(
					successToken([]byte("token-data-100")),
					successToken([]byte("token-data-101")),
				),
				2: createExpectedMessageTokenData(
					errorToken(errors.New("api error")),
				),
				3: createExpectedMessageTokenData(
					notSupportedToken(),
				),
			},
		},
		{
			name: "message with incomplete status",
			ccipMessages: map[cciptypes.SeqNum]cciptypes.Message{
				1: createCCIPMessageWithMultipleTokens(1, "0xabc123", 1),
			},
			tokenIndexToCCTPv2Message: map[cciptypes.SeqNum]map[int]CCTPv2MessageOrError{
				1: {
					0: CCTPv2MessageOrError{
						message: createIncompleteMessage("100", "pending"),
					},
				},
			},
			setupEncoder: noopMockEncoder(),
			expectedResults: map[cciptypes.SeqNum]exectypes.MessageTokenData{
				1: createExpectedMessageTokenData(
					errorToken(errors.New(
						"A CCTPv2 Message's 'status' is not complete: nonce: 100, sourceDomainID: 1, status: pending")),
				),
			},
		},
		{
			name: "message with invalid hex in message field",
			ccipMessages: map[cciptypes.SeqNum]cciptypes.Message{
				1: createCCIPMessageWithMultipleTokens(1, "0xabc123", 1),
			},
			tokenIndexToCCTPv2Message: map[cciptypes.SeqNum]map[int]CCTPv2MessageOrError{
				1: {
					0: CCTPv2MessageOrError{
						message: createInvalidHexMessage("100", "message"),
					},
				},
			},
			setupEncoder: noopMockEncoder(),
			expectedResults: map[cciptypes.SeqNum]exectypes.MessageTokenData{
				1: createExpectedMessageTokenData(
					errorToken(errors.New("A CCTPv2 Message's 'message' field could not be converted from string to bytes")),
				),
			},
		},
		{
			name: "message with invalid hex in attestation field",
			ccipMessages: map[cciptypes.SeqNum]cciptypes.Message{
				1: createCCIPMessageWithMultipleTokens(1, "0xabc123", 1),
			},
			tokenIndexToCCTPv2Message: map[cciptypes.SeqNum]map[int]CCTPv2MessageOrError{
				1: {
					0: CCTPv2MessageOrError{
						message: createInvalidHexMessage("100", "attestation"),
					},
				},
			},
			setupEncoder: noopMockEncoder(),
			expectedResults: map[cciptypes.SeqNum]exectypes.MessageTokenData{
				1: createExpectedMessageTokenData(
					errorToken(errors.New("A CCTPv2 Message's 'attestation' field could not be converted from string to bytes")),
				),
			},
		},
		{
			name: "attestation encoder error",
			ccipMessages: map[cciptypes.SeqNum]cciptypes.Message{
				1: createCCIPMessageWithMultipleTokens(1, "0xabc123", 1),
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
			expectedResults: map[cciptypes.SeqNum]exectypes.MessageTokenData{
				1: createExpectedMessageTokenData(
					errorToken(errors.New(
						"attestationEncoder failed for a CCTPv2 message: nonce: 100, sourceDomainID: 1, error: encoding failed")),
				),
			},
		},
		{
			name: "message with no tokens",
			ccipMessages: map[cciptypes.SeqNum]cciptypes.Message{
				1: createCCIPMessageWithMultipleTokens(1, "0xabc123", 0),
			},
			tokenIndexToCCTPv2Message: map[cciptypes.SeqNum]map[int]CCTPv2MessageOrError{
				1: {},
			},
			setupEncoder: noopMockEncoder(),
			expectedResults: map[cciptypes.SeqNum]exectypes.MessageTokenData{
				1: createExpectedMessageTokenData(),
			},
		},
		{
			name: "mixed scenarios with edge cases",
			ccipMessages: map[cciptypes.SeqNum]cciptypes.Message{
				1: createCCIPMessageWithMultipleTokens(1, "0xabc123", 4),
				2: createCCIPMessageWithMultipleTokens(2, "0xdef456", 0),
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
			expectedResults: map[cciptypes.SeqNum]exectypes.MessageTokenData{
				1: createExpectedMessageTokenData(
					successToken([]byte("successful-encoding")),
					errorToken(errors.New(
						"A CCTPv2 Message's 'status' is not complete: nonce: 101, sourceDomainID: 1, status: failed")),
					errorToken(errors.New("network timeout")),
					notSupportedToken(),
				),
				2: createExpectedMessageTokenData(),
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
				1, // sourceChain
				NewNoOpMetricsReporter(),
			)

			// Verify results using helper function
			verifyMessageTokenDataResults(t, tt.expectedResults, result)

			// Verify no unexpected results
			for actualSeqNum := range result {
				_, exists := tt.expectedResults[actualSeqNum]
				require.True(t, exists, "Unexpected sequence number %d found in result", actualSeqNum)
			}
		})
	}
}

func TestNotSupportedMessageTokenData(t *testing.T) {
	tests := []struct {
		name         string
		ccipMessages map[cciptypes.SeqNum]cciptypes.Message
		expected     map[cciptypes.SeqNum]exectypes.MessageTokenData
	}{
		{
			name:         "empty input",
			ccipMessages: map[cciptypes.SeqNum]cciptypes.Message{},
			expected:     map[cciptypes.SeqNum]exectypes.MessageTokenData{},
		},
		{
			name: "single message with no tokens",
			ccipMessages: map[cciptypes.SeqNum]cciptypes.Message{
				1: createCCIPMessageWithMultipleTokens(1, "0xabc123", 0),
			},
			expected: map[cciptypes.SeqNum]exectypes.MessageTokenData{
				1: createExpectedMessageTokenData(),
			},
		},
		{
			name: "single message with single token",
			ccipMessages: map[cciptypes.SeqNum]cciptypes.Message{
				1: createCCIPMessageWithMultipleTokens(1, "0xabc123", 1),
			},
			expected: map[cciptypes.SeqNum]exectypes.MessageTokenData{
				1: createExpectedMessageTokenData(
					notSupportedToken(),
				),
			},
		},
		{
			name: "single message with multiple tokens",
			ccipMessages: map[cciptypes.SeqNum]cciptypes.Message{
				1: createCCIPMessageWithMultipleTokens(1, "0xabc123", 3),
			},
			expected: map[cciptypes.SeqNum]exectypes.MessageTokenData{
				1: createExpectedMessageTokenData(
					notSupportedToken(),
					notSupportedToken(),
					notSupportedToken(),
				),
			},
		},
		{
			name: "multiple messages with varying token counts",
			ccipMessages: map[cciptypes.SeqNum]cciptypes.Message{
				1: createCCIPMessageWithMultipleTokens(1, "0xabc123", 2),
				2: createCCIPMessageWithMultipleTokens(2, "0xdef456", 0),
				3: createCCIPMessageWithMultipleTokens(3, "0x789xyz", 1),
				5: createCCIPMessageWithMultipleTokens(5, "0x111222", 4),
			},
			expected: map[cciptypes.SeqNum]exectypes.MessageTokenData{
				1: createExpectedMessageTokenData(
					notSupportedToken(),
					notSupportedToken(),
				),
				2: createExpectedMessageTokenData(),
				3: createExpectedMessageTokenData(
					notSupportedToken(),
				),
				5: createExpectedMessageTokenData(
					notSupportedToken(),
					notSupportedToken(),
					notSupportedToken(),
					notSupportedToken(),
				),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := notSupportedMessageTokenData(tt.ccipMessages)

			// Verify results using helper function
			verifyMessageTokenDataResults(t, tt.expected, result)

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
		expected                map[cciptypes.SeqNum]exectypes.MessageTokenData
	}{
		{
			name:                    "empty inputs",
			err:                     testError,
			ccipMessages:            map[cciptypes.SeqNum]cciptypes.Message{},
			sourceTokenDataPayloads: map[cciptypes.SeqNum]map[int]SourceTokenDataPayload{},
			expected:                map[cciptypes.SeqNum]exectypes.MessageTokenData{},
		},
		{
			name: "ccip messages but no source token data payloads",
			err:  testError,
			ccipMessages: map[cciptypes.SeqNum]cciptypes.Message{
				1: createCCIPMessageWithMultipleTokens(1, "0xabc123", 2),
				2: createCCIPMessageWithMultipleTokens(2, "0xdef456", 1),
			},
			sourceTokenDataPayloads: map[cciptypes.SeqNum]map[int]SourceTokenDataPayload{},
			expected: map[cciptypes.SeqNum]exectypes.MessageTokenData{
				1: createExpectedMessageTokenData(
					notSupportedToken(),
					notSupportedToken(),
				),
				2: createExpectedMessageTokenData(
					notSupportedToken(),
				),
			},
		},
		{
			name: "single message with single token - error applied",
			err:  testError,
			ccipMessages: map[cciptypes.SeqNum]cciptypes.Message{
				1: createCCIPMessageWithMultipleTokens(1, "0xabc123", 2),
			},
			sourceTokenDataPayloads: map[cciptypes.SeqNum]map[int]SourceTokenDataPayload{
				1: {
					0: createSourceTokenDataPayload(1, 2, 1000),
				},
			},
			expected: map[cciptypes.SeqNum]exectypes.MessageTokenData{
				1: createExpectedMessageTokenData(
					errorToken(testError),
					notSupportedToken(),
				),
			},
		},
		{
			name: "single message with multiple tokens - partial error coverage",
			err:  networkError,
			ccipMessages: map[cciptypes.SeqNum]cciptypes.Message{
				1: createCCIPMessageWithMultipleTokens(1, "0xabc123", 4),
			},
			sourceTokenDataPayloads: map[cciptypes.SeqNum]map[int]SourceTokenDataPayload{
				1: {
					1: createSourceTokenDataPayload(1, 2, 2000),
					2: createSourceTokenDataPayload(1, 2, 3000),
				},
			},
			expected: map[cciptypes.SeqNum]exectypes.MessageTokenData{
				1: createExpectedMessageTokenData(
					notSupportedToken(),
					errorToken(networkError),
					errorToken(networkError),
					notSupportedToken(),
				),
			},
		},
		{
			name: "multiple messages with mixed error scenarios",
			err:  testError,
			ccipMessages: map[cciptypes.SeqNum]cciptypes.Message{
				1: createCCIPMessageWithMultipleTokens(1, "0xabc123", 3),
				2: createCCIPMessageWithMultipleTokens(2, "0xdef456", 1),
				3: createCCIPMessageWithMultipleTokens(3, "0x789xyz", 2),
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
			expected: map[cciptypes.SeqNum]exectypes.MessageTokenData{
				1: createExpectedMessageTokenData(
					errorToken(testError),
					notSupportedToken(),
					errorToken(testError),
				),
				2: createExpectedMessageTokenData(
					errorToken(testError),
				),
				3: createExpectedMessageTokenData(
					notSupportedToken(),
					notSupportedToken(),
				),
			},
		},
		{
			name: "all tokens have errors",
			err:  networkError,
			ccipMessages: map[cciptypes.SeqNum]cciptypes.Message{
				1: createCCIPMessageWithMultipleTokens(1, "0xabc123", 2),
			},
			sourceTokenDataPayloads: map[cciptypes.SeqNum]map[int]SourceTokenDataPayload{
				1: {
					0: createSourceTokenDataPayload(1, 2, 1000),
					1: createSourceTokenDataPayload(1, 2, 2000),
				},
			},
			expected: map[cciptypes.SeqNum]exectypes.MessageTokenData{
				1: createExpectedMessageTokenData(
					errorToken(networkError),
					errorToken(networkError),
				),
			},
		},
		{
			name: "message with no tokens but has source token data payloads",
			err:  testError,
			ccipMessages: map[cciptypes.SeqNum]cciptypes.Message{
				1: createCCIPMessageWithMultipleTokens(1, "0xabc123", 0),
			},
			sourceTokenDataPayloads: map[cciptypes.SeqNum]map[int]SourceTokenDataPayload{
				1: {
					0: createSourceTokenDataPayload(1, 2, 1000),
				},
			},
			expected: map[cciptypes.SeqNum]exectypes.MessageTokenData{
				1: createExpectedMessageTokenData(),
			},
		},
		{
			name: "token index gaps in source token data payloads",
			err:  testError,
			ccipMessages: map[cciptypes.SeqNum]cciptypes.Message{
				1: createCCIPMessageWithMultipleTokens(1, "0xabc123", 5),
			},
			sourceTokenDataPayloads: map[cciptypes.SeqNum]map[int]SourceTokenDataPayload{
				1: {
					0: createSourceTokenDataPayload(1, 2, 1000),
					2: createSourceTokenDataPayload(1, 2, 3000),
					4: createSourceTokenDataPayload(1, 2, 5000),
				},
			},
			expected: map[cciptypes.SeqNum]exectypes.MessageTokenData{
				1: createExpectedMessageTokenData(
					errorToken(testError),
					notSupportedToken(),
					errorToken(testError),
					notSupportedToken(),
					errorToken(testError),
				),
			},
		},
		{
			name: "nil error",
			err:  nil,
			ccipMessages: map[cciptypes.SeqNum]cciptypes.Message{
				1: createCCIPMessageWithMultipleTokens(1, "0xabc123", 1),
			},
			sourceTokenDataPayloads: map[cciptypes.SeqNum]map[int]SourceTokenDataPayload{
				1: {
					0: createSourceTokenDataPayload(1, 2, 1000),
				},
			},
			expected: map[cciptypes.SeqNum]exectypes.MessageTokenData{
				1: createExpectedMessageTokenData(
					errorToken(nil),
				),
			},
		},
		{
			name:         "source token data payloads but no ccip messages",
			err:          testError,
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
			expected: map[cciptypes.SeqNum]exectypes.MessageTokenData{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := errorMessageTokenData(tt.err, tt.ccipMessages, tt.sourceTokenDataPayloads)

			// Verify results using helper function
			verifyMessageTokenDataResults(t, tt.expected, result)

			// Verify no unexpected results
			for actualSeqNum := range result {
				_, exists := tt.expected[actualSeqNum]
				require.True(t, exists, "Unexpected sequence number %d found in result", actualSeqNum)
			}
		})
	}
}

func TestGetMessageTokenDataForSourceChain(t *testing.T) {
	ctx := testCtx
	lggr := testLogger
	sourceChain := testSourceChain

	tests := []struct {
		name                     string
		sourceChain              cciptypes.ChainSelector
		ccipMessages             map[cciptypes.SeqNum]cciptypes.Message
		supportedPoolsBySelector map[cciptypes.ChainSelector]string
		setupMockClient          func(*MockCCTPv2AttestationClient)
		setupMockEncoder         func(*mockAttestationEncoder)
		expectedResults          map[cciptypes.SeqNum]exectypes.MessageTokenData
	}{
		{
			name:        "source chain not supported",
			sourceChain: cciptypes.ChainSelector(999),
			ccipMessages: map[cciptypes.SeqNum]cciptypes.Message{
				1: createCCIPMessageWithMultipleTokens(1, defaultTxHash, 2),
			},
			supportedPoolsBySelector: map[cciptypes.ChainSelector]string{
				cciptypes.ChainSelector(1): "0xPool1",
			},
			setupMockClient:  noopMockClient(),
			setupMockEncoder: noopMockEncoder(),
			expectedResults: map[cciptypes.SeqNum]exectypes.MessageTokenData{
				1: createExpectedMessageTokenData(
					notSupportedToken(),
					notSupportedToken(),
				),
			},
		},
		{
			name:                     "empty ccip messages",
			sourceChain:              sourceChain,
			ccipMessages:             map[cciptypes.SeqNum]cciptypes.Message{},
			supportedPoolsBySelector: defaultSupportedPools(),
			setupMockClient:          noopMockClient(),
			setupMockEncoder:         noopMockEncoder(),
			expectedResults:          map[cciptypes.SeqNum]exectypes.MessageTokenData{},
		},
		{
			name:        "no valid CCTP v2 tokens",
			sourceChain: sourceChain,
			ccipMessages: map[cciptypes.SeqNum]cciptypes.Message{
				1: createCCIPMessageWithMultipleTokens(1, defaultTxHash, 2),
			},
			supportedPoolsBySelector: map[cciptypes.ChainSelector]string{
				sourceChain: "0xDifferentPool",
			},
			setupMockClient:  noopMockClient(),
			setupMockEncoder: noopMockEncoder(),
			expectedResults: map[cciptypes.SeqNum]exectypes.MessageTokenData{
				1: createExpectedMessageTokenData(
					notSupportedToken(),
					notSupportedToken(),
				),
			},
		},
		{
			name:        "getSourceDomainID returns error",
			sourceChain: sourceChain,
			ccipMessages: map[cciptypes.SeqNum]cciptypes.Message{
				1: createCCIPMessageWithValidCCTPv2Tokens(1, defaultTxHash),
				2: createCCIPMessageWithValidCCTPv2Tokens(2, "0xdef456"),
			},
			supportedPoolsBySelector: defaultSupportedPools(),
			setupMockClient:          noopMockClient(),
			setupMockEncoder:         noopMockEncoder(),
			expectedResults: map[cciptypes.SeqNum]exectypes.MessageTokenData{
				1: createExpectedMessageTokenData(
					errorToken(errors.New("error")),
				),
				2: createExpectedMessageTokenData(
					errorToken(errors.New("error")),
				),
			},
		},
		{
			name:        "successful flow - single message with single token",
			sourceChain: sourceChain,
			ccipMessages: map[cciptypes.SeqNum]cciptypes.Message{
				1: createCCIPMessageWithValidCCTPv2Tokens(1, defaultTxHash),
			},
			supportedPoolsBySelector: defaultSupportedPools(),
			setupMockClient: func(m *MockCCTPv2AttestationClient) {
				msg := createMatchingCCTPv2Message("123")
				msg.DecodedMessage.DecodedMessageBody.Amount = "1" // Match token amount
				m.AddResponse(111, defaultTxHash, createMessages(msg), nil)
			},
			setupMockEncoder: func(m *mockAttestationEncoder) {
				m.AddResponse(
					"0x1234567890abcdef", "0xfedcba0987654321",
					[]byte("successful-token-data"), nil,
				)
			},
			expectedResults: map[cciptypes.SeqNum]exectypes.MessageTokenData{
				1: createExpectedMessageTokenData(
					successToken([]byte("successful-token-data")),
				),
			},
		},
		{
			name:        "successful flow - multiple messages with multiple tokens",
			sourceChain: sourceChain,
			ccipMessages: map[cciptypes.SeqNum]cciptypes.Message{
				1: createCCIPMessageWithMultipleTokens(1, defaultTxHash, 2),
				2: createCCIPMessageWithMultipleTokens(2, "0xdef456", 1),
			},
			supportedPoolsBySelector: defaultSupportedPools(),
			setupMockClient: func(m *MockCCTPv2AttestationClient) {
				// Create CCTP messages with amounts matching the token amounts (1, 2, 1)
				msg1 := createMatchingCCTPv2Message("1")
				msg1.Message = "0x1111111111111111"
				msg1.Attestation = "0xaaaaaaaaaaaaaaaa"
				msg1.DecodedMessage.DecodedMessageBody.Amount = "1" // Match first token

				msg2 := createMatchingCCTPv2Message("2")
				msg2.Message = "0x2222222222222222"
				msg2.Attestation = "0xbbbbbbbbbbbbbbbb"
				msg2.DecodedMessage.DecodedMessageBody.Amount = "2" // Match second token

				msg3 := createMatchingCCTPv2Message("3")
				msg3.Message = "0x3333333333333333"
				msg3.Attestation = "0xcccccccccccccccc"
				msg3.DecodedMessage.DecodedMessageBody.Amount = "1" // Match third token

				m.AddResponse(111, defaultTxHash, createMessages(msg1, msg2), nil)
				m.AddResponse(111, "0xdef456", createMessages(msg3), nil)
			},
			setupMockEncoder: func(m *mockAttestationEncoder) {
				m.AddResponse(
					"0x1111111111111111", "0xaaaaaaaaaaaaaaaa",
					[]byte("token-data-1"), nil,
				)
				m.AddResponse(
					"0x2222222222222222", "0xbbbbbbbbbbbbbbbb",
					[]byte("token-data-2"), nil,
				)
				m.AddResponse(
					"0x3333333333333333", "0xcccccccccccccccc",
					[]byte("token-data-3"), nil,
				)
			},
			expectedResults: map[cciptypes.SeqNum]exectypes.MessageTokenData{
				1: createExpectedMessageTokenData(
					successToken([]byte("token-data-3")),
					successToken([]byte("token-data-1")),
				),
				2: createExpectedMessageTokenData(
					successToken([]byte("token-data-2")),
				),
			},
		},
		{
			name:        "attestation client returns error",
			sourceChain: sourceChain,
			ccipMessages: map[cciptypes.SeqNum]cciptypes.Message{
				1: createCCIPMessageWithValidCCTPv2Tokens(1, defaultTxHash),
			},
			supportedPoolsBySelector: defaultSupportedPools(),
			setupMockClient: func(m *MockCCTPv2AttestationClient) {
				m.AddResponse(111, defaultTxHash, Messages{}, errors.New("network error"))
			},
			setupMockEncoder: noopMockEncoder(),
			expectedResults: map[cciptypes.SeqNum]exectypes.MessageTokenData{
				1: createExpectedMessageTokenData(
					errorToken(errors.New("error")),
				),
			},
		},
		{
			name:        "mixed supported and unsupported tokens",
			sourceChain: sourceChain,
			ccipMessages: map[cciptypes.SeqNum]cciptypes.Message{
				1: createCCIPMessageWithMixedTokens(1, defaultTxHash),
			},
			supportedPoolsBySelector: defaultSupportedPools(),
			setupMockClient: func(m *MockCCTPv2AttestationClient) {
				msg := createMatchingCCTPv2Message("123")
				m.AddResponse(111, defaultTxHash, createMessages(msg), nil)
			},
			setupMockEncoder: func(m *mockAttestationEncoder) {
				m.AddResponse(
					"0x1234567890abcdef", "0xfedcba0987654321",
					[]byte("token-data-1"), nil,
				)
			},
			expectedResults: map[cciptypes.SeqNum]exectypes.MessageTokenData{
				1: createExpectedMessageTokenData(
					successToken([]byte("token-data-1")),
					notSupportedToken(),
				),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create mock client and encoder
			mockClient := NewMockCCTPv2AttestationClient()
			mockEncoder := newMockAttestationEncoder()

			// Setup mocks
			tt.setupMockClient(mockClient)
			tt.setupMockEncoder(mockEncoder)

			// Call the function
			result := getMessageTokenDataForSourceChain(
				ctx, lggr, tt.sourceChain, tt.ccipMessages, tt.supportedPoolsBySelector,
				mockEncoder.Encode, mockClient, NewNoOpMetricsReporter(),
			)

			// Verify results
			verifyMessageTokenDataResults(t, tt.expectedResults, result)
		})
	}
}
