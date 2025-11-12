package v2

import (
	"testing"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/execute/tokendata"
)

// Helper function to create UnknownAddress from hex string
func mustUnknownAddressFromHex(hex string) cciptypes.UnknownAddress {
	addr, err := cciptypes.NewUnknownAddressFromHex(hex)
	if err != nil {
		panic(err)
	}
	return addr
}

func Test_getV2TokenPayloads(t *testing.T) {
	const (
		sourceChain = cciptypes.ChainSelector(1)
		poolAddr    = "0x1234567890123456789012345678901234567890"
		otherPool   = "0xabcdefabcdefabcdefabcdefabcdefabcdefabcd"
	)

	validDepositHash := "1234567890123456789012345678901234567890123456789012345678901234"

	tests := []struct {
		name                     string
		chainSelector            cciptypes.ChainSelector
		messages                 map[cciptypes.SeqNum]cciptypes.Message
		supportedPoolsBySelector map[cciptypes.ChainSelector]string
		expectedPayloads         map[cciptypes.SeqNum]map[int]SourceTokenDataPayloadV2
	}{
		{
			name:          "empty messages",
			chainSelector: sourceChain,
			messages:      map[cciptypes.SeqNum]cciptypes.Message{},
			supportedPoolsBySelector: map[cciptypes.ChainSelector]string{
				sourceChain: poolAddr,
			},
			expectedPayloads: map[cciptypes.SeqNum]map[int]SourceTokenDataPayloadV2{},
		},
		{
			name:          "single message with single CCTPv2 token",
			chainSelector: sourceChain,
			messages: map[cciptypes.SeqNum]cciptypes.Message{
				1: {
					TokenAmounts: []cciptypes.RampTokenAmount{
						{
							SourcePoolAddress: mustUnknownAddressFromHex(poolAddr),
							ExtraData:         createValidExtraData(CCTPVersion2Tag, 1, validDepositHash),
						},
					},
				},
			},
			supportedPoolsBySelector: map[cciptypes.ChainSelector]string{
				sourceChain: poolAddr,
			},
			expectedPayloads: map[cciptypes.SeqNum]map[int]SourceTokenDataPayloadV2{
				1: {
					0: {
						SourceDomain: 1,
						DepositHash:  mustHexToBytes32(validDepositHash),
					},
				},
			},
		},
		{
			name:          "single message with multiple CCTPv2 tokens",
			chainSelector: sourceChain,
			messages: map[cciptypes.SeqNum]cciptypes.Message{
				1: {
					TokenAmounts: []cciptypes.RampTokenAmount{
						{
							SourcePoolAddress: mustUnknownAddressFromHex(poolAddr),
							ExtraData:         createValidExtraData(CCTPVersion2Tag, 1, validDepositHash),
						},
						{
							SourcePoolAddress: mustUnknownAddressFromHex(poolAddr),
							ExtraData: createValidExtraData(
								CCTPVersion2CCVTag, 2, "abcdef0123456789abcdef0123456789abcdef0123456789abcdef0123456789",
							),
						},
					},
				},
			},
			supportedPoolsBySelector: map[cciptypes.ChainSelector]string{
				sourceChain: poolAddr,
			},
			expectedPayloads: map[cciptypes.SeqNum]map[int]SourceTokenDataPayloadV2{
				1: {
					0: {
						SourceDomain: 1,
						DepositHash:  mustHexToBytes32(validDepositHash),
					},
					1: {
						SourceDomain: 2,
						DepositHash: mustHexToBytes32(
							"abcdef0123456789abcdef0123456789abcdef0123456789abcdef0123456789",
						),
					},
				},
			},
		},
		{
			name:          "multiple messages with CCTPv2 tokens",
			chainSelector: sourceChain,
			messages: map[cciptypes.SeqNum]cciptypes.Message{
				1: {
					TokenAmounts: []cciptypes.RampTokenAmount{
						{
							SourcePoolAddress: mustUnknownAddressFromHex(poolAddr),
							ExtraData:         createValidExtraData(CCTPVersion2Tag, 1, validDepositHash),
						},
					},
				},
				2: {
					TokenAmounts: []cciptypes.RampTokenAmount{
						{
							SourcePoolAddress: mustUnknownAddressFromHex(poolAddr),
							ExtraData: createValidExtraData(
								CCTPVersion2CCVTag, 2, "abcdef0123456789abcdef0123456789abcdef0123456789abcdef0123456789",
							),
						},
					},
				},
			},
			supportedPoolsBySelector: map[cciptypes.ChainSelector]string{
				sourceChain: poolAddr,
			},
			expectedPayloads: map[cciptypes.SeqNum]map[int]SourceTokenDataPayloadV2{
				1: {
					0: {
						SourceDomain: 1,
						DepositHash:  mustHexToBytes32(validDepositHash),
					},
				},
				2: {
					0: {
						SourceDomain: 2,
						DepositHash: mustHexToBytes32(
							"abcdef0123456789abcdef0123456789abcdef0123456789abcdef0123456789",
						),
					},
				},
			},
		},
		{
			name:          "message with unsupported pool address",
			chainSelector: sourceChain,
			messages: map[cciptypes.SeqNum]cciptypes.Message{
				1: {
					TokenAmounts: []cciptypes.RampTokenAmount{
						{
							SourcePoolAddress: mustUnknownAddressFromHex(otherPool), // Different pool
							ExtraData:         createValidExtraData(CCTPVersion2Tag, 1, validDepositHash),
						},
					},
				},
			},
			supportedPoolsBySelector: map[cciptypes.ChainSelector]string{
				sourceChain: poolAddr,
			},
			expectedPayloads: map[cciptypes.SeqNum]map[int]SourceTokenDataPayloadV2{},
		},
		{
			name:          "message with invalid ExtraData",
			chainSelector: sourceChain,
			messages: map[cciptypes.SeqNum]cciptypes.Message{
				1: {
					TokenAmounts: []cciptypes.RampTokenAmount{
						{
							SourcePoolAddress: mustUnknownAddressFromHex(poolAddr),
							ExtraData:         cciptypes.Bytes{0x01, 0x02, 0x03}, // Invalid data
						},
					},
				},
			},
			supportedPoolsBySelector: map[cciptypes.ChainSelector]string{
				sourceChain: poolAddr,
			},
			expectedPayloads: map[cciptypes.SeqNum]map[int]SourceTokenDataPayloadV2{},
		},
		{
			name:          "message with mixed valid and invalid tokens",
			chainSelector: sourceChain,
			messages: map[cciptypes.SeqNum]cciptypes.Message{
				1: {
					TokenAmounts: []cciptypes.RampTokenAmount{
						{
							SourcePoolAddress: mustUnknownAddressFromHex(poolAddr),
							ExtraData:         createValidExtraData(CCTPVersion2Tag, 1, validDepositHash),
						},
						{
							SourcePoolAddress: mustUnknownAddressFromHex(otherPool), // Unsupported pool
							ExtraData: createValidExtraData(
								CCTPVersion2Tag, 2, "abcdef0123456789abcdef0123456789abcdef0123456789abcdef0123456789",
							),
						},
						{
							SourcePoolAddress: mustUnknownAddressFromHex(poolAddr),
							ExtraData:         cciptypes.Bytes{0x01, 0x02}, // Invalid data
						},
						{
							SourcePoolAddress: mustUnknownAddressFromHex(poolAddr),
							ExtraData: createValidExtraData(
								CCTPVersion2CCVTag, 3, "ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff",
							),
						},
					},
				},
			},
			supportedPoolsBySelector: map[cciptypes.ChainSelector]string{
				sourceChain: poolAddr,
			},
			expectedPayloads: map[cciptypes.SeqNum]map[int]SourceTokenDataPayloadV2{
				1: {
					0: {
						SourceDomain: 1,
						DepositHash:  mustHexToBytes32(validDepositHash),
					},
					3: {
						SourceDomain: 3,
						DepositHash: mustHexToBytes32(
							"ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff",
						),
					},
				},
			},
		},
		{
			name:          "message with no tokens",
			chainSelector: sourceChain,
			messages: map[cciptypes.SeqNum]cciptypes.Message{
				1: {
					TokenAmounts: []cciptypes.RampTokenAmount{},
				},
			},
			supportedPoolsBySelector: map[cciptypes.ChainSelector]string{
				sourceChain: poolAddr,
			},
			expectedPayloads: map[cciptypes.SeqNum]map[int]SourceTokenDataPayloadV2{},
		},
		{
			name:          "message with CCTPv2 CCV tag",
			chainSelector: sourceChain,
			messages: map[cciptypes.SeqNum]cciptypes.Message{
				1: {
					TokenAmounts: []cciptypes.RampTokenAmount{
						{
							SourcePoolAddress: mustUnknownAddressFromHex(poolAddr),
							ExtraData: createValidExtraData(
								CCTPVersion2CCVTag, 5, "deadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeef",
							),
						},
					},
				},
			},
			supportedPoolsBySelector: map[cciptypes.ChainSelector]string{
				sourceChain: poolAddr,
			},
			expectedPayloads: map[cciptypes.SeqNum]map[int]SourceTokenDataPayloadV2{
				1: {
					0: {
						SourceDomain: 5,
						DepositHash: mustHexToBytes32(
							"deadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeef",
						),
					},
				},
			},
		},
		{
			name:          "no supported pools configured",
			chainSelector: sourceChain,
			messages: map[cciptypes.SeqNum]cciptypes.Message{
				1: {
					TokenAmounts: []cciptypes.RampTokenAmount{
						{
							SourcePoolAddress: mustUnknownAddressFromHex(poolAddr),
							ExtraData:         createValidExtraData(CCTPVersion2Tag, 1, validDepositHash),
						},
					},
				},
			},
			supportedPoolsBySelector: map[cciptypes.ChainSelector]string{},
			expectedPayloads:         map[cciptypes.SeqNum]map[int]SourceTokenDataPayloadV2{},
		},
		{
			name:          "pool address case insensitive matching",
			chainSelector: sourceChain,
			messages: map[cciptypes.SeqNum]cciptypes.Message{
				1: {
					TokenAmounts: []cciptypes.RampTokenAmount{
						{
							SourcePoolAddress: mustUnknownAddressFromHex("0x1234567890abcdef1234567890abcdef12345678"),
							ExtraData:         createValidExtraData(CCTPVersion2Tag, 1, validDepositHash),
						},
						{
							// Same address, different case
							SourcePoolAddress: mustUnknownAddressFromHex("0x1234567890ABCDEF1234567890ABCDEF12345678"),
							ExtraData: createValidExtraData(
								CCTPVersion2Tag, 2, "abcdef0123456789abcdef0123456789abcdef0123456789abcdef0123456789",
							),
						},
					},
				},
			},
			supportedPoolsBySelector: map[cciptypes.ChainSelector]string{
				sourceChain: "0x1234567890abcdef1234567890abcdef12345678",
			},
			expectedPayloads: map[cciptypes.SeqNum]map[int]SourceTokenDataPayloadV2{
				1: {
					0: {
						SourceDomain: 1,
						DepositHash:  mustHexToBytes32(validDepositHash),
					},
					1: {
						SourceDomain: 2,
						DepositHash: mustHexToBytes32(
							"abcdef0123456789abcdef0123456789abcdef0123456789abcdef0123456789",
						),
					},
				},
			},
		},
		{
			name:          "multiple messages with only some containing CCTPv2 tokens",
			chainSelector: sourceChain,
			messages: map[cciptypes.SeqNum]cciptypes.Message{
				1: {
					TokenAmounts: []cciptypes.RampTokenAmount{
						{
							SourcePoolAddress: mustUnknownAddressFromHex(poolAddr),
							ExtraData:         createValidExtraData(CCTPVersion2Tag, 1, validDepositHash),
						},
					},
				},
				2: {
					TokenAmounts: []cciptypes.RampTokenAmount{
						{
							SourcePoolAddress: mustUnknownAddressFromHex(otherPool), // Wrong pool
							ExtraData: createValidExtraData(
								CCTPVersion2Tag, 2, "abcdef0123456789abcdef0123456789abcdef0123456789abcdef0123456789",
							),
						},
					},
				},
				3: {
					TokenAmounts: []cciptypes.RampTokenAmount{
						{
							SourcePoolAddress: mustUnknownAddressFromHex(poolAddr),
							ExtraData: createValidExtraData(
								CCTPVersion2CCVTag, 3, "ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff",
							),
						},
					},
				},
			},
			supportedPoolsBySelector: map[cciptypes.ChainSelector]string{
				sourceChain: poolAddr,
			},
			expectedPayloads: map[cciptypes.SeqNum]map[int]SourceTokenDataPayloadV2{
				1: {
					0: {
						SourceDomain: 1,
						DepositHash:  mustHexToBytes32(validDepositHash),
					},
				},
				3: {
					0: {
						SourceDomain: 3,
						DepositHash: mustHexToBytes32(
							"ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff",
						),
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			observer := NewCCTPv2TokenDataObserver(
				logger.Test(t),
				cciptypes.ChainSelector(999), // destChain
				tt.supportedPoolsBySelector,
				nil, // attestationEncoder not needed for this test
				nil, // httpClient not needed for this test
			)

			result := observer.getV2TokenPayloads(tt.chainSelector, tt.messages)

			require.Equal(t, len(tt.expectedPayloads), len(result), "unexpected number of messages with payloads")

			for seqNum, expectedTokenPayloads := range tt.expectedPayloads {
				actualTokenPayloads, ok := result[seqNum]
				require.True(t, ok, "expected sequence number %d not found in result", seqNum)
				require.Equal(
					t, len(expectedTokenPayloads), len(actualTokenPayloads),
					"unexpected number of token payloads for seqNum %d", seqNum,
				)

				for tokenIdx, expectedPayload := range expectedTokenPayloads {
					actualPayload, ok := actualTokenPayloads[tokenIdx]
					require.True(t, ok, "expected token index %d not found for seqNum %d", tokenIdx, seqNum)
					assert.Equal(
						t, expectedPayload.SourceDomain, actualPayload.SourceDomain,
						"source domain mismatch for seqNum %d, tokenIdx %d", seqNum, tokenIdx,
					)
					assert.Equal(
						t, expectedPayload.DepositHash, actualPayload.DepositHash,
						"deposit hash mismatch for seqNum %d, tokenIdx %d", seqNum, tokenIdx,
					)
				}
			}
		})
	}
}

func Test_getSourceDomainID(t *testing.T) {
	tests := []struct {
		name             string
		v2TokenPayloads  map[cciptypes.SeqNum]map[int]SourceTokenDataPayloadV2
		expectedDomainID uint32
		wantErr          bool
		errContains      string
	}{
		{
			name: "single message with single token",
			v2TokenPayloads: map[cciptypes.SeqNum]map[int]SourceTokenDataPayloadV2{
				1: {
					0: {
						SourceDomain: 5,
						DepositHash:  mustHexToBytes32("1234567890123456789012345678901234567890123456789012345678901234"),
					},
				},
			},
			expectedDomainID: 5,
			wantErr:          false,
		},
		{
			name: "single message with multiple tokens - same domain",
			v2TokenPayloads: map[cciptypes.SeqNum]map[int]SourceTokenDataPayloadV2{
				1: {
					0: {
						SourceDomain: 3,
						DepositHash:  mustHexToBytes32("1234567890123456789012345678901234567890123456789012345678901234"),
					},
					1: {
						SourceDomain: 3,
						DepositHash:  mustHexToBytes32("abcdef0123456789abcdef0123456789abcdef0123456789abcdef0123456789"),
					},
					2: {
						SourceDomain: 3,
						DepositHash:  mustHexToBytes32("fedcba9876543210fedcba9876543210fedcba9876543210fedcba9876543210"),
					},
				},
			},
			expectedDomainID: 3,
			wantErr:          false,
		},
		{
			name: "multiple messages - same domain",
			v2TokenPayloads: map[cciptypes.SeqNum]map[int]SourceTokenDataPayloadV2{
				1: {
					0: {
						SourceDomain: 7,
						DepositHash:  mustHexToBytes32("1234567890123456789012345678901234567890123456789012345678901234"),
					},
				},
				2: {
					0: {
						SourceDomain: 7,
						DepositHash:  mustHexToBytes32("abcdef0123456789abcdef0123456789abcdef0123456789abcdef0123456789"),
					},
				},
				3: {
					0: {
						SourceDomain: 7,
						DepositHash:  mustHexToBytes32("fedcba9876543210fedcba9876543210fedcba9876543210fedcba9876543210"),
					},
				},
			},
			expectedDomainID: 7,
			wantErr:          false,
		},
		{
			name: "multiple messages with multiple tokens - same domain",
			v2TokenPayloads: map[cciptypes.SeqNum]map[int]SourceTokenDataPayloadV2{
				1: {
					0: {
						SourceDomain: 2,
						DepositHash:  mustHexToBytes32("1234567890123456789012345678901234567890123456789012345678901234"),
					},
					1: {
						SourceDomain: 2,
						DepositHash:  mustHexToBytes32("abcdef0123456789abcdef0123456789abcdef0123456789abcdef0123456789"),
					},
				},
				2: {
					0: {
						SourceDomain: 2,
						DepositHash:  mustHexToBytes32("fedcba9876543210fedcba9876543210fedcba9876543210fedcba9876543210"),
					},
				},
				5: {
					0: {
						SourceDomain: 2,
						DepositHash:  mustHexToBytes32("0000000000000000000000000000000000000000000000000000000000000000"),
					},
					2: {
						SourceDomain: 2,
						DepositHash:  mustHexToBytes32("ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff"),
					},
				},
			},
			expectedDomainID: 2,
			wantErr:          false,
		},
		{
			name:            "empty payload map",
			v2TokenPayloads: map[cciptypes.SeqNum]map[int]SourceTokenDataPayloadV2{},
			wantErr:         true,
			errContains:     "no CCTPv2 token payloads found",
		},
		{
			name:            "nil payload map",
			v2TokenPayloads: nil,
			wantErr:         true,
			errContains:     "no CCTPv2 token payloads found",
		},
		{
			name: "inconsistent domain IDs - different messages",
			v2TokenPayloads: map[cciptypes.SeqNum]map[int]SourceTokenDataPayloadV2{
				1: {
					0: {
						SourceDomain: 5,
						DepositHash:  mustHexToBytes32("1234567890123456789012345678901234567890123456789012345678901234"),
					},
				},
				2: {
					0: {
						SourceDomain: 7, // Different domain!
						DepositHash:  mustHexToBytes32("abcdef0123456789abcdef0123456789abcdef0123456789abcdef0123456789"),
					},
				},
			},
			wantErr:     true,
			errContains: "inconsistent source domain IDs found",
		},
		{
			name: "inconsistent domain IDs - same message",
			v2TokenPayloads: map[cciptypes.SeqNum]map[int]SourceTokenDataPayloadV2{
				1: {
					0: {
						SourceDomain: 3,
						DepositHash:  mustHexToBytes32("1234567890123456789012345678901234567890123456789012345678901234"),
					},
					1: {
						SourceDomain: 4, // Different domain!
						DepositHash:  mustHexToBytes32("abcdef0123456789abcdef0123456789abcdef0123456789abcdef0123456789"),
					},
				},
			},
			wantErr:     true,
			errContains: "inconsistent source domain IDs found",
		},
		{
			name: "inconsistent domain IDs - multiple messages, later conflict",
			v2TokenPayloads: map[cciptypes.SeqNum]map[int]SourceTokenDataPayloadV2{
				1: {
					0: {
						SourceDomain: 1,
						DepositHash:  mustHexToBytes32("1234567890123456789012345678901234567890123456789012345678901234"),
					},
				},
				2: {
					0: {
						SourceDomain: 1,
						DepositHash:  mustHexToBytes32("abcdef0123456789abcdef0123456789abcdef0123456789abcdef0123456789"),
					},
				},
				3: {
					0: {
						SourceDomain: 2, // Conflict on third message
						DepositHash:  mustHexToBytes32("fedcba9876543210fedcba9876543210fedcba9876543210fedcba9876543210"),
					},
				},
			},
			wantErr:     true,
			errContains: "inconsistent source domain IDs found",
		},
		{
			name: "domain ID zero is valid",
			v2TokenPayloads: map[cciptypes.SeqNum]map[int]SourceTokenDataPayloadV2{
				1: {
					0: {
						SourceDomain: 0,
						DepositHash:  mustHexToBytes32("1234567890123456789012345678901234567890123456789012345678901234"),
					},
				},
			},
			expectedDomainID: 0,
			wantErr:          false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getSourceDomainID(tt.v2TokenPayloads)

			if tt.wantErr {
				require.Error(t, err)
				if tt.errContains != "" {
					assert.Contains(t, err.Error(), tt.errContains)
				}
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.expectedDomainID, got)
			}
		})
	}
}

func Test_getTxHashes(t *testing.T) {
	tests := []struct {
		name           string
		messages       map[cciptypes.SeqNum]cciptypes.Message
		expectedResult map[string][]cciptypes.SeqNum
	}{
		{
			name:           "empty messages",
			messages:       map[cciptypes.SeqNum]cciptypes.Message{},
			expectedResult: map[string][]cciptypes.SeqNum{},
		},
		{
			name: "single message",
			messages: map[cciptypes.SeqNum]cciptypes.Message{
				1: {
					Header: cciptypes.RampMessageHeader{
						TxHash: "0x1234567890123456789012345678901234567890123456789012345678901234",
					},
				},
			},
			expectedResult: map[string][]cciptypes.SeqNum{
				"0x1234567890123456789012345678901234567890123456789012345678901234": {1},
			},
		},
		{
			name: "multiple messages with different tx hashes",
			messages: map[cciptypes.SeqNum]cciptypes.Message{
				1: {
					Header: cciptypes.RampMessageHeader{
						TxHash: "0x1234567890123456789012345678901234567890123456789012345678901234",
					},
				},
				2: {
					Header: cciptypes.RampMessageHeader{
						TxHash: "0xabcdef0123456789abcdef0123456789abcdef0123456789abcdef0123456789",
					},
				},
				3: {
					Header: cciptypes.RampMessageHeader{
						TxHash: "0xfedcba9876543210fedcba9876543210fedcba9876543210fedcba9876543210",
					},
				},
			},
			expectedResult: map[string][]cciptypes.SeqNum{
				"0x1234567890123456789012345678901234567890123456789012345678901234": {1},
				"0xabcdef0123456789abcdef0123456789abcdef0123456789abcdef0123456789": {2},
				"0xfedcba9876543210fedcba9876543210fedcba9876543210fedcba9876543210": {3},
			},
		},
		{
			name: "multiple messages sharing the same tx hash",
			messages: map[cciptypes.SeqNum]cciptypes.Message{
				1: {
					Header: cciptypes.RampMessageHeader{
						TxHash: "0x1234567890123456789012345678901234567890123456789012345678901234",
					},
				},
				2: {
					Header: cciptypes.RampMessageHeader{
						TxHash: "0x1234567890123456789012345678901234567890123456789012345678901234",
					},
				},
				3: {
					Header: cciptypes.RampMessageHeader{
						TxHash: "0x1234567890123456789012345678901234567890123456789012345678901234",
					},
				},
			},
			expectedResult: map[string][]cciptypes.SeqNum{
				"0x1234567890123456789012345678901234567890123456789012345678901234": {1, 2, 3},
			},
		},
		{
			name: "mixed - some messages share tx hash, others don't",
			messages: map[cciptypes.SeqNum]cciptypes.Message{
				1: {
					Header: cciptypes.RampMessageHeader{
						TxHash: "0x1111111111111111111111111111111111111111111111111111111111111111",
					},
				},
				2: {
					Header: cciptypes.RampMessageHeader{
						TxHash: "0x2222222222222222222222222222222222222222222222222222222222222222",
					},
				},
				3: {
					Header: cciptypes.RampMessageHeader{
						TxHash: "0x1111111111111111111111111111111111111111111111111111111111111111",
					},
				},
				4: {
					Header: cciptypes.RampMessageHeader{
						TxHash: "0x3333333333333333333333333333333333333333333333333333333333333333",
					},
				},
				5: {
					Header: cciptypes.RampMessageHeader{
						TxHash: "0x2222222222222222222222222222222222222222222222222222222222222222",
					},
				},
			},
			expectedResult: map[string][]cciptypes.SeqNum{
				"0x1111111111111111111111111111111111111111111111111111111111111111": {1, 3},
				"0x2222222222222222222222222222222222222222222222222222222222222222": {2, 5},
				"0x3333333333333333333333333333333333333333333333333333333333333333": {4},
			},
		},
		{
			name: "empty tx hash string",
			messages: map[cciptypes.SeqNum]cciptypes.Message{
				1: {
					Header: cciptypes.RampMessageHeader{
						TxHash: "",
					},
				},
				2: {
					Header: cciptypes.RampMessageHeader{
						TxHash: "",
					},
				},
			},
			expectedResult: map[string][]cciptypes.SeqNum{
				"": {1, 2},
			},
		},
		{
			name: "non-sequential sequence numbers",
			messages: map[cciptypes.SeqNum]cciptypes.Message{
				10: {
					Header: cciptypes.RampMessageHeader{
						TxHash: "0xaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
					},
				},
				100: {
					Header: cciptypes.RampMessageHeader{
						TxHash: "0xaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
					},
				},
				1000: {
					Header: cciptypes.RampMessageHeader{
						TxHash: "0xbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb",
					},
				},
			},
			expectedResult: map[string][]cciptypes.SeqNum{
				"0xaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa": {10, 100},
				"0xbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb": {1000},
			},
		},
		{
			name: "case sensitivity - different case same hash value",
			messages: map[cciptypes.SeqNum]cciptypes.Message{
				1: {
					Header: cciptypes.RampMessageHeader{
						TxHash: "0xabcdef0123456789abcdef0123456789abcdef0123456789abcdef0123456789",
					},
				},
				2: {
					Header: cciptypes.RampMessageHeader{
						TxHash: "0xABCDEF0123456789ABCDEF0123456789ABCDEF0123456789ABCDEF0123456789",
					},
				},
			},
			expectedResult: map[string][]cciptypes.SeqNum{
				"0xabcdef0123456789abcdef0123456789abcdef0123456789abcdef0123456789": {1},
				"0xABCDEF0123456789ABCDEF0123456789ABCDEF0123456789ABCDEF0123456789": {2},
			},
		},
		{
			name: "large number of messages with same tx hash",
			messages: map[cciptypes.SeqNum]cciptypes.Message{
				1: {
					Header: cciptypes.RampMessageHeader{
						TxHash: "0x9999999999999999999999999999999999999999999999999999999999999999",
					},
				},
				2: {
					Header: cciptypes.RampMessageHeader{
						TxHash: "0x9999999999999999999999999999999999999999999999999999999999999999",
					},
				},
				3: {
					Header: cciptypes.RampMessageHeader{
						TxHash: "0x9999999999999999999999999999999999999999999999999999999999999999",
					},
				},
				4: {
					Header: cciptypes.RampMessageHeader{
						TxHash: "0x9999999999999999999999999999999999999999999999999999999999999999",
					},
				},
				5: {
					Header: cciptypes.RampMessageHeader{
						TxHash: "0x9999999999999999999999999999999999999999999999999999999999999999",
					},
				},
			},
			expectedResult: map[string][]cciptypes.SeqNum{
				"0x9999999999999999999999999999999999999999999999999999999999999999": {1, 2, 3, 4, 5},
			},
		},
		{
			name: "tx hash without 0x prefix",
			messages: map[cciptypes.SeqNum]cciptypes.Message{
				1: {
					Header: cciptypes.RampMessageHeader{
						TxHash: "1234567890123456789012345678901234567890123456789012345678901234",
					},
				},
			},
			expectedResult: map[string][]cciptypes.SeqNum{
				"1234567890123456789012345678901234567890123456789012345678901234": {1},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := getTxHashes(tt.messages)

			require.Equal(t, len(tt.expectedResult), len(result), "unexpected number of unique tx hashes")

			for txHash, expectedSeqNums := range tt.expectedResult {
				actualSeqNums, ok := result[txHash]
				require.True(t, ok, "expected tx hash %s not found in result", txHash)

				// Sort both slices for comparison since map iteration order is non-deterministic
				require.ElementsMatch(t, expectedSeqNums, actualSeqNums,
					"sequence numbers mismatch for tx hash %s", txHash)
			}
		})
	}
}

func TestIsTokenSupported(t *testing.T) {
	const (
		sourceChain = cciptypes.ChainSelector(1)
		poolAddr    = "0x1234567890123456789012345678901234567890"
		otherPool   = "0xabcdefabcdefabcdefabcdefabcdefabcdefabcd"
	)

	validDepositHash := "1234567890123456789012345678901234567890123456789012345678901234"

	tests := []struct {
		name                     string
		sourceChain              cciptypes.ChainSelector
		msgToken                 cciptypes.RampTokenAmount
		supportedPoolsBySelector map[cciptypes.ChainSelector]string
		want                     bool
	}{
		{
			name:        "supported token with valid V2 payload",
			sourceChain: sourceChain,
			msgToken: cciptypes.RampTokenAmount{
				SourcePoolAddress: mustUnknownAddressFromHex(poolAddr),
				ExtraData:         createValidExtraData(CCTPVersion2Tag, 1, validDepositHash),
			},
			supportedPoolsBySelector: map[cciptypes.ChainSelector]string{
				sourceChain: poolAddr,
			},
			want: true,
		},
		{
			name:        "supported token with V2 CCV tag",
			sourceChain: sourceChain,
			msgToken: cciptypes.RampTokenAmount{
				SourcePoolAddress: mustUnknownAddressFromHex(poolAddr),
				ExtraData:         createValidExtraData(CCTPVersion2CCVTag, 1, validDepositHash),
			},
			supportedPoolsBySelector: map[cciptypes.ChainSelector]string{
				sourceChain: poolAddr,
			},
			want: true,
		},
		{
			name:        "unsupported pool address",
			sourceChain: sourceChain,
			msgToken: cciptypes.RampTokenAmount{
				SourcePoolAddress: mustUnknownAddressFromHex(otherPool),
				ExtraData:         createValidExtraData(CCTPVersion2Tag, 1, validDepositHash),
			},
			supportedPoolsBySelector: map[cciptypes.ChainSelector]string{
				sourceChain: poolAddr,
			},
			want: false,
		},
		{
			name:        "supported pool but invalid ExtraData",
			sourceChain: sourceChain,
			msgToken: cciptypes.RampTokenAmount{
				SourcePoolAddress: mustUnknownAddressFromHex(poolAddr),
				ExtraData:         cciptypes.Bytes{0x01, 0x02, 0x03},
			},
			supportedPoolsBySelector: map[cciptypes.ChainSelector]string{
				sourceChain: poolAddr,
			},
			want: false,
		},
		{
			name:        "supported pool but empty ExtraData",
			sourceChain: sourceChain,
			msgToken: cciptypes.RampTokenAmount{
				SourcePoolAddress: mustUnknownAddressFromHex(poolAddr),
				ExtraData:         cciptypes.Bytes{},
			},
			supportedPoolsBySelector: map[cciptypes.ChainSelector]string{
				sourceChain: poolAddr,
			},
			want: false,
		},
		{
			name:        "no pool configured for chain",
			sourceChain: sourceChain,
			msgToken: cciptypes.RampTokenAmount{
				SourcePoolAddress: mustUnknownAddressFromHex(poolAddr),
				ExtraData:         createValidExtraData(CCTPVersion2Tag, 1, validDepositHash),
			},
			supportedPoolsBySelector: map[cciptypes.ChainSelector]string{},
			want:                     false,
		},
		{
			name:        "case insensitive pool address matching - different hex case",
			sourceChain: sourceChain,
			msgToken: cciptypes.RampTokenAmount{
				SourcePoolAddress: mustUnknownAddressFromHex("0x1234567890ABCDEF1234567890ABCDEF12345678"),
				ExtraData:         createValidExtraData(CCTPVersion2Tag, 1, validDepositHash),
			},
			supportedPoolsBySelector: map[cciptypes.ChainSelector]string{
				sourceChain: "0x1234567890abcdef1234567890abcdef12345678",
			},
			want: true,
		},
		{
			name:        "case insensitive pool address matching - mixed case",
			sourceChain: sourceChain,
			msgToken: cciptypes.RampTokenAmount{
				SourcePoolAddress: mustUnknownAddressFromHex("0x1234567890ABCDEF1234567890ABCDEF12345678"),
				ExtraData:         createValidExtraData(CCTPVersion2Tag, 1, validDepositHash),
			},
			supportedPoolsBySelector: map[cciptypes.ChainSelector]string{
				sourceChain: "0x1234567890abcdef1234567890abcdef12345678",
			},
			want: true,
		},
		{
			name:        "invalid version tag",
			sourceChain: sourceChain,
			msgToken: cciptypes.RampTokenAmount{
				SourcePoolAddress: mustUnknownAddressFromHex(poolAddr),
				ExtraData:         createValidExtraData(0x12345678, 1, validDepositHash),
			},
			supportedPoolsBySelector: map[cciptypes.ChainSelector]string{
				sourceChain: poolAddr,
			},
			want: false,
		},
		{
			name:        "wrong chain selector",
			sourceChain: cciptypes.ChainSelector(999),
			msgToken: cciptypes.RampTokenAmount{
				SourcePoolAddress: mustUnknownAddressFromHex(poolAddr),
				ExtraData:         createValidExtraData(CCTPVersion2Tag, 1, validDepositHash),
			},
			supportedPoolsBySelector: map[cciptypes.ChainSelector]string{
				sourceChain: poolAddr,
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			observer := NewCCTPv2TokenDataObserver(
				logger.Test(t),
				cciptypes.ChainSelector(999),
				tt.supportedPoolsBySelector,
				nil,
				nil,
			)

			got := observer.IsTokenSupported(tt.sourceChain, tt.msgToken)
			assert.Equal(t, tt.want, got)
		})
	}
}

// Helper function to create a success attestation status
func createSuccessAttestation(id, body, attestation string) tokendata.AttestationStatus {
	return tokendata.SuccessAttestationStatus(
		cciptypes.Bytes(id),
		cciptypes.Bytes(body),
		cciptypes.Bytes(attestation),
	)
}

func Test_assignAttestationForV2TokenPayload(t *testing.T) {
	validDepositHash1 := mustHexToBytes32("1234567890123456789012345678901234567890123456789012345678901234")
	validDepositHash2 := mustHexToBytes32("abcdef0123456789abcdef0123456789abcdef0123456789abcdef0123456789")
	validDepositHash3 := mustHexToBytes32("fedcba9876543210fedcba9876543210fedcba9876543210fedcba9876543210")

	attestation1 := createSuccessAttestation("id1", "body1", "attestation1")
	attestation2 := createSuccessAttestation("id2", "body2", "attestation2")
	attestation3 := createSuccessAttestation("id3", "body3", "attestation3")
	v2TokenPayload1 := SourceTokenDataPayloadV2{
		SourceDomain: 1,
		DepositHash:  validDepositHash1,
	}

	tests := []struct {
		name                   string
		attestations           map[[32]byte][]tokendata.AttestationStatus
		v2TokenPayload         SourceTokenDataPayloadV2
		expectedAttestation    tokendata.AttestationStatus
		expectedRemainingCount int
		checkRemainingInMap    bool
		expectedError          error
	}{
		{
			name: "single attestation for deposit hash - success",
			attestations: map[[32]byte][]tokendata.AttestationStatus{
				validDepositHash1: {attestation1},
			},
			v2TokenPayload:         v2TokenPayload1,
			expectedAttestation:    attestation1,
			expectedRemainingCount: 0,
			checkRemainingInMap:    true,
		},
		{
			name: "multiple attestations - returns first and removes it",
			attestations: map[[32]byte][]tokendata.AttestationStatus{
				validDepositHash1: {attestation1, attestation2, attestation3},
			},
			v2TokenPayload:         v2TokenPayload1,
			expectedAttestation:    attestation1,
			expectedRemainingCount: 2,
			checkRemainingInMap:    true,
		},
		{
			name: "deposit hash not found in map",
			attestations: map[[32]byte][]tokendata.AttestationStatus{
				validDepositHash1: {attestation1},
			},
			v2TokenPayload: SourceTokenDataPayloadV2{
				SourceDomain: 2,
				DepositHash:  validDepositHash2, // Different hash not in map
			},
			expectedAttestation: tokendata.ErrorAttestationStatus(tokendata.ErrDataMissing),
			expectedError:       tokendata.ErrDataMissing,
		},
		{
			name: "empty attestation slice for deposit hash",
			attestations: map[[32]byte][]tokendata.AttestationStatus{
				validDepositHash1: {}, // Empty slice
			},
			v2TokenPayload:      v2TokenPayload1,
			expectedAttestation: tokendata.ErrorAttestationStatus(tokendata.ErrDataMissing),
			expectedError:       tokendata.ErrDataMissing,
		},
		{
			name:                "empty attestations map",
			attestations:        map[[32]byte][]tokendata.AttestationStatus{},
			v2TokenPayload:      v2TokenPayload1,
			expectedAttestation: tokendata.ErrorAttestationStatus(tokendata.ErrDataMissing),
			expectedError:       tokendata.ErrDataMissing,
		},
		{
			name: "attestation with error status",
			attestations: map[[32]byte][]tokendata.AttestationStatus{
				validDepositHash1: {
					tokendata.ErrorAttestationStatus(tokendata.ErrNotReady),
				},
			},
			v2TokenPayload:      v2TokenPayload1,
			expectedAttestation: tokendata.ErrorAttestationStatus(tokendata.ErrNotReady),
			expectedError:       tokendata.ErrNotReady,
		},
		{
			name: "multiple hashes in map - correct one selected",
			attestations: map[[32]byte][]tokendata.AttestationStatus{
				validDepositHash1: {attestation1},
				validDepositHash2: {attestation2},
				validDepositHash3: {attestation3},
			},
			v2TokenPayload: SourceTokenDataPayloadV2{
				SourceDomain: 2,
				DepositHash:  validDepositHash2,
			},
			expectedAttestation:    attestation2,
			expectedRemainingCount: 0,
			checkRemainingInMap:    true,
		},
		{
			name: "two attestations - first call gets first, map updated",
			attestations: map[[32]byte][]tokendata.AttestationStatus{
				validDepositHash1: {attestation1, attestation2},
			},
			v2TokenPayload:         v2TokenPayload1,
			expectedAttestation:    attestation1,
			expectedRemainingCount: 1,
			checkRemainingInMap:    true,
		},
		{
			name: "attestation with various error types - ErrRateLimit",
			attestations: map[[32]byte][]tokendata.AttestationStatus{
				validDepositHash1: {
					tokendata.ErrorAttestationStatus(tokendata.ErrRateLimit),
				},
			},
			v2TokenPayload:      v2TokenPayload1,
			expectedAttestation: tokendata.ErrorAttestationStatus(tokendata.ErrRateLimit),
			expectedError:       tokendata.ErrRateLimit,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Make a copy of the attestations map to verify mutation
			attestationsCopy := make(map[[32]byte][]tokendata.AttestationStatus)
			for k, v := range tt.attestations {
				statusesCopy := make([]tokendata.AttestationStatus, len(v))
				copy(statusesCopy, v)
				attestationsCopy[k] = statusesCopy
			}

			// Call the function
			result := assignAttestationForV2TokenPayload(attestationsCopy, tt.v2TokenPayload)

			// Verify the returned attestation
			if tt.expectedError != nil {
				// Check error case
				require.Error(t, result.Error)
				assert.ErrorIs(t, result.Error, tt.expectedError)
				assert.Nil(t, result.ID)
				assert.Nil(t, result.MessageBody)
				assert.Nil(t, result.Attestation)
			} else {
				// Check success case
				require.NoError(t, result.Error)
				assert.Equal(t, tt.expectedAttestation.ID, result.ID)
				assert.Equal(t, tt.expectedAttestation.MessageBody, result.MessageBody)
				assert.Equal(t, tt.expectedAttestation.Attestation, result.Attestation)
			}

			// Verify map mutation if applicable
			if tt.checkRemainingInMap {
				remainingStatuses, ok := attestationsCopy[tt.v2TokenPayload.DepositHash]
				require.True(t, ok, "deposit hash should still exist in map")
				assert.Equal(
					t, tt.expectedRemainingCount, len(remainingStatuses),
					"unexpected number of remaining attestations in map",
				)
			}
		})
	}
}

func Test_assignAttestationForV2TokenPayload_Mutation(t *testing.T) {
	// This test specifically verifies that the function mutates the attestations map
	// by removing the first element and that subsequent calls work correctly
	depositHash := mustHexToBytes32("1234567890123456789012345678901234567890123456789012345678901234")

	attestation1 := createSuccessAttestation("id1", "body1", "attestation1")
	attestation2 := createSuccessAttestation("id2", "body2", "attestation2")
	attestation3 := createSuccessAttestation("id3", "body3", "attestation3")

	attestations := map[[32]byte][]tokendata.AttestationStatus{
		depositHash: {attestation1, attestation2, attestation3},
	}

	payload := SourceTokenDataPayloadV2{
		SourceDomain: 1,
		DepositHash:  depositHash,
	}

	// First call - should get first attestation
	result1 := assignAttestationForV2TokenPayload(attestations, payload)
	require.NoError(t, result1.Error)
	assert.Equal(t, cciptypes.Bytes("id1"), result1.ID)
	assert.Equal(t, 2, len(attestations[depositHash]), "should have 2 attestations remaining")

	// Second call - should get second attestation (which was originally at index 1)
	result2 := assignAttestationForV2TokenPayload(attestations, payload)
	require.NoError(t, result2.Error)
	assert.Equal(t, cciptypes.Bytes("id2"), result2.ID)
	assert.Equal(t, 1, len(attestations[depositHash]), "should have 1 attestation remaining")

	// Third call - should get third attestation (which was originally at index 2)
	result3 := assignAttestationForV2TokenPayload(attestations, payload)
	require.NoError(t, result3.Error)
	assert.Equal(t, cciptypes.Bytes("id3"), result3.ID)
	assert.Equal(t, 0, len(attestations[depositHash]), "should have 0 attestations remaining")

	// Fourth call - should return error since no attestations left
	result4 := assignAttestationForV2TokenPayload(attestations, payload)
	require.Error(t, result4.Error)
	assert.ErrorIs(t, result4.Error, tokendata.ErrDataMissing)
}
