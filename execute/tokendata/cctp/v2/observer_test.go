package v2

import (
	"testing"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
