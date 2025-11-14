package v2

import (
	"encoding/hex"
	"math/big"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDepositHash_CalculateDepositHash(t *testing.T) {
	tests := []struct {
		name        string
		msg         CCTPv2DecodedMessage
		wantHash    string
		wantErr     bool
		errContains string
	}{
		{
			name: "valid message with all fields",
			msg: CCTPv2DecodedMessage{
				SourceDomain:         "1",
				DestinationDomain:    "2",
				MinFinalityThreshold: "3",
				DestinationCaller:    "0x1234567890123456789012345678901234567890",
				DecodedMessageBody: CCTPv2DecodedMessageBody{
					Amount:        "1000000",
					MintRecipient: "0xabcdef0123456789abcdef0123456789abcdef01",
					BurnToken:     "0xfedcba9876543210fedcba9876543210fedcba98",
					MaxFee:        "5000",
				},
			},
			wantHash: "c4725a4a49f7c0ac609229956d379d1fe4c65cc2c999e2b4eb615eeac81acf4b",
			wantErr:  false,
		},
		{
			name: "valid message with zero maxFee",
			msg: CCTPv2DecodedMessage{
				SourceDomain:         "0",
				DestinationDomain:    "6",
				MinFinalityThreshold: "1",
				DestinationCaller:    "0x0000000000000000000000000000000000000000",
				DecodedMessageBody: CCTPv2DecodedMessageBody{
					Amount:        "100",
					MintRecipient: "0x1111111111111111111111111111111111111111",
					BurnToken:     "0x2222222222222222222222222222222222222222",
					MaxFee:        "",
				},
			},
			wantHash: "c2f4f2be83c80fba28e0376776129c2340ab878a3af6863f45dc1dd83b32b467",
			wantErr:  false,
		},
		{
			name: "valid message with Solana address (32 bytes)",
			msg: CCTPv2DecodedMessage{
				SourceDomain:         "5",
				DestinationDomain:    "1",
				MinFinalityThreshold: "2",
				DestinationCaller:    "1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef",
				DecodedMessageBody: CCTPv2DecodedMessageBody{
					Amount:        "999999999",
					MintRecipient: "abcdef0123456789abcdef0123456789abcdef0123456789abcdef0123456789",
					BurnToken:     "fedcba9876543210fedcba9876543210fedcba9876543210fedcba9876543210",
					MaxFee:        "100000",
				},
			},
			wantHash: "5fe7eece7497635d1406a4d0d8b865cd080cf032182aefe47952427506e1b295",
			wantErr:  false,
		},
		{
			name: "invalid sourceDomain",
			msg: CCTPv2DecodedMessage{
				SourceDomain: "not_a_number",
			},
			wantErr:     true,
			errContains: "failed to parse source domain",
		},
		{
			name: "invalid destinationDomain",
			msg: CCTPv2DecodedMessage{
				SourceDomain:      "1",
				DestinationDomain: "invalid",
			},
			wantErr:     true,
			errContains: "failed to parse destination domain",
		},
		{
			name: "invalid amount",
			msg: CCTPv2DecodedMessage{
				SourceDomain:      "1",
				DestinationDomain: "2",
				DecodedMessageBody: CCTPv2DecodedMessageBody{
					Amount: "not_a_number",
				},
			},
			wantErr:     true,
			errContains: "parse amount",
		},
		{
			name: "invalid maxFee",
			msg: CCTPv2DecodedMessage{
				SourceDomain:      "1",
				DestinationDomain: "2",
				DecodedMessageBody: CCTPv2DecodedMessageBody{
					Amount: "1000",
					MaxFee: "invalid_fee",
				},
			},
			wantErr:     true,
			errContains: "failed to parse max fee",
		},
		{
			name: "invalid minFinalityThreshold",
			msg: CCTPv2DecodedMessage{
				SourceDomain:         "1",
				DestinationDomain:    "2",
				MinFinalityThreshold: "not_a_number",
				DecodedMessageBody: CCTPv2DecodedMessageBody{
					Amount: "1000",
				},
			},
			wantErr:     true,
			errContains: "failed to parse min finality threshold",
		},
		{
			name: "invalid mintRecipient - not hex",
			msg: CCTPv2DecodedMessage{
				SourceDomain:         "1",
				DestinationDomain:    "2",
				MinFinalityThreshold: "1",
				DecodedMessageBody: CCTPv2DecodedMessageBody{
					Amount:        "1000",
					MintRecipient: "not_hex",
				},
			},
			wantErr:     true,
			errContains: "failed to parse mint recipient",
		},
		{
			name: "invalid burnToken - too long",
			msg: CCTPv2DecodedMessage{
				SourceDomain:         "1",
				DestinationDomain:    "2",
				MinFinalityThreshold: "1",
				DecodedMessageBody: CCTPv2DecodedMessageBody{
					Amount:        "1000",
					MintRecipient: "0x1234567890123456789012345678901234567890",
					BurnToken:     "0x" + strings.Repeat("a", 66), // 33 bytes
				},
			},
			wantErr:     true,
			errContains: "failed to parse burn token",
		},
		{
			name: "invalid destinationCaller - odd hex length",
			msg: CCTPv2DecodedMessage{
				SourceDomain:         "1",
				DestinationDomain:    "2",
				MinFinalityThreshold: "1",
				DestinationCaller:    "0x12345", // 5 chars after 0x, which is odd length
				DecodedMessageBody: CCTPv2DecodedMessageBody{
					Amount:        "1000",
					MintRecipient: "0x1234567890123456789012345678901234567890",
					BurnToken:     "0xfedcba9876543210fedcba9876543210fedcba98",
				},
			},
			wantErr:     true,
			errContains: "failed to parse destination caller",
		},
		{
			name: "very large amount",
			msg: CCTPv2DecodedMessage{
				SourceDomain:         "1",
				DestinationDomain:    "2",
				MinFinalityThreshold: "1",
				DestinationCaller:    "0x0000000000000000000000000000000000000000",
				DecodedMessageBody: CCTPv2DecodedMessageBody{
					Amount:        "115792089237316195423570985008687907853269984665640564039457584007913129639935", // max uint256
					MintRecipient: "0x1234567890123456789012345678901234567890",
					BurnToken:     "0xfedcba9876543210fedcba9876543210fedcba98",
					MaxFee:        "0",
				},
			},
			wantHash: "61fd67d439c34bf4b24232540b37f82aca74874c1bc3a6ddeaae0563db0a9e88",
			wantErr:  false,
		},
		{
			name: "sourceDomain overflow",
			msg: CCTPv2DecodedMessage{
				SourceDomain: "4294967296", // 2^32
			},
			wantErr:     true,
			errContains: "failed to parse source domain",
		},
		{
			name: "negative maxFee",
			msg: CCTPv2DecodedMessage{
				SourceDomain:         "1",
				DestinationDomain:    "2",
				MinFinalityThreshold: "1",
				DestinationCaller:    "0x0000000000000000000000000000000000000000",
				DecodedMessageBody: CCTPv2DecodedMessageBody{
					Amount:        "1000",
					MintRecipient: "0x1234567890123456789012345678901234567890",
					BurnToken:     "0xfedcba9876543210fedcba9876543210fedcba98",
					MaxFee:        "-100",
				},
			},
			wantErr:     true,
			errContains: "failed to parse max fee: negative fee not allowed",
		},
		{
			name: "negative amount",
			msg: CCTPv2DecodedMessage{
				SourceDomain:         "1",
				DestinationDomain:    "2",
				MinFinalityThreshold: "1",
				DestinationCaller:    "0x0000000000000000000000000000000000000000",
				DecodedMessageBody: CCTPv2DecodedMessageBody{
					Amount:        "-1000",
					MintRecipient: "0x1234567890123456789012345678901234567890",
					BurnToken:     "0xfedcba9876543210fedcba9876543210fedcba98",
				},
			},
			wantErr:     true,
			errContains: "failed to parse amount: negative amount not allowed",
		},
		{
			name: "all maximum values",
			msg: CCTPv2DecodedMessage{
				SourceDomain:         "4294967295", // max uint32
				DestinationDomain:    "4294967295",
				MinFinalityThreshold: "4294967295",
				DestinationCaller:    "0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff",
				DecodedMessageBody: CCTPv2DecodedMessageBody{
					Amount:        "115792089237316195423570985008687907853269984665640564039457584007913129639935", // max uint256
					MintRecipient: "0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff",
					BurnToken:     "0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff",
					MaxFee:        "115792089237316195423570985008687907853269984665640564039457584007913129639935",
				},
			},
			wantHash: "f5d6bde1f508b7f18864e809a28dcce4275d87dc9886829f9c752e2789c982fd",
			wantErr:  false,
		},
		{
			name:        "empty message",
			msg:         CCTPv2DecodedMessage{},
			wantErr:     true,
			errContains: "failed to parse source domain",
		},
		{
			name: "addresses without 0x prefix",
			msg: CCTPv2DecodedMessage{
				SourceDomain:         "1",
				DestinationDomain:    "2",
				MinFinalityThreshold: "1",
				DestinationCaller:    "1234567890123456789012345678901234567890",
				DecodedMessageBody: CCTPv2DecodedMessageBody{
					Amount:        "1000",
					MintRecipient: "abcdef0123456789abcdef0123456789abcdef01",
					BurnToken:     "fedcba9876543210fedcba9876543210fedcba98",
					MaxFee:        "100",
				},
			},
			wantHash: "41beb119bf267455c2fc99e2af43f78bf87c990da647138969e8849a3f159a92",
			wantErr:  false,
		},
		{
			name: "real world testnet example",
			msg: CCTPv2DecodedMessage{
				SourceDomain:         "0", // Ethereum
				DestinationDomain:    "6", // Avalanche
				MinFinalityThreshold: "12",
				DestinationCaller:    "0x0000000000000000000000000000000000000000",
				DecodedMessageBody: CCTPv2DecodedMessageBody{
					Amount:        "1000000", // 1 USDC
					MintRecipient: "0x5FC8d32690cc91D4c39d9d3abcBD16989F875707",
					BurnToken:     "0x07865c6e87b9f70255377e024ace6630c1eaa37f", // Goerli USDC
					MaxFee:        "0",
				},
			},
			wantHash: "e503377195659620e33836dc4b6321ad6bec38908b54e901edf5302d3d083212",
			wantErr:  false,
		},
		{
			name: "cross-chain arbitrage scenario",
			msg: CCTPv2DecodedMessage{
				SourceDomain:         "1", // Ethereum
				DestinationDomain:    "7", // Polygon
				MinFinalityThreshold: "100",
				DestinationCaller:    "0xDEADBEEF00000000000000000000000000000000",
				DecodedMessageBody: CCTPv2DecodedMessageBody{
					Amount:        "50000000000", // 50k USDC
					MintRecipient: "0xCAFEBABE00000000000000000000000000000000",
					BurnToken:     "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48", // Mainnet USDC
					MaxFee:        "1000000",                                    // 1 USDC fee
				},
			},
			wantHash: "ab11c698c8bb206a261ab5312f3c57e61caa7b5f2085855c5827b5a841484c6e",
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotHash, err := CalculateDepositHash(tt.msg)
			if tt.wantErr {
				require.Error(t, err)
				if tt.errContains != "" {
					assert.Contains(t, err.Error(), tt.errContains)
				}
				assert.Equal(t, [32]byte{}, gotHash)
			} else {
				require.NoError(t, err)
				wantBytes, err := hex.DecodeString(tt.wantHash)
				require.NoError(t, err)
				var want [32]byte
				copy(want[:], wantBytes)
				assert.Equal(t, want, gotHash, "hash mismatch for test: %s\nExpected: %x\nGot: %x", tt.name, want, gotHash)
			}
		})
	}
}

func TestDepositHash_HexToBytes32(t *testing.T) {
	tests := []struct {
		name        string
		hexStr      string
		want        [32]byte
		wantErr     bool
		errContains string
	}{
		{
			name:    "valid 20-byte address with 0x prefix",
			hexStr:  "0x1234567890123456789012345678901234567890",
			want:    mustHexToBytes32("0000000000000000000000001234567890123456789012345678901234567890"),
			wantErr: false,
		},
		{
			name:    "valid 20-byte address without 0x prefix",
			hexStr:  "1234567890123456789012345678901234567890",
			want:    mustHexToBytes32("0000000000000000000000001234567890123456789012345678901234567890"),
			wantErr: false,
		},
		{
			name:    "valid 32-byte Solana address",
			hexStr:  "1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef",
			want:    mustHexToBytes32("1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef"),
			wantErr: false,
		},
		{
			name:        "invalid hex characters",
			hexStr:      "0xGGGG567890123456789012345678901234567890",
			wantErr:     true,
			errContains: "decode hex",
		},
		{
			name:        "odd length hex",
			hexStr:      "0x12345", // 5 chars after 0x, which is odd
			wantErr:     true,
			errContains: "decode hex",
		},
		{
			name:        "too long (33 bytes)",
			hexStr:      "0x" + strings.Repeat("aa", 33),
			wantErr:     true,
			errContains: "hex string too long (33 bytes)",
		},
		{
			name:    "empty string",
			hexStr:  "",
			want:    [32]byte{},
			wantErr: false,
		},
		{
			name:    "just 0x prefix",
			hexStr:  "0x",
			want:    [32]byte{},
			wantErr: false,
		},
		{
			name:    "1 byte input",
			hexStr:  "0xFF",
			want:    mustHexToBytes32("00000000000000000000000000000000000000000000000000000000000000FF"),
			wantErr: false,
		},
		{
			name:    "31 bytes input",
			hexStr:  strings.Repeat("ab", 31),
			want:    mustHexToBytes32("00" + strings.Repeat("ab", 31)),
			wantErr: false,
		},
		{
			name:    "all zeros",
			hexStr:  "0x0000000000000000000000000000000000000000",
			want:    [32]byte{},
			wantErr: false,
		},
		{
			name:    "all ones",
			hexStr:  "0xffffffffffffffffffffffffffffffffffffffff",
			want:    mustHexToBytes32("000000000000000000000000ffffffffffffffffffffffffffffffffffffffff"),
			wantErr: false,
		},
		{
			name:    "mixed case hex",
			hexStr:  "0xAbCdEf0123456789aBcDeF0123456789AbCdEf01",
			want:    mustHexToBytes32("000000000000000000000000AbCdEf0123456789aBcDeF0123456789AbCdEf01"),
			wantErr: false,
		},
		{
			name:        "special characters in hex",
			hexStr:      "0x12-34-56-78",
			wantErr:     true,
			errContains: "decode hex",
		},
		{
			name:    "exactly 32 bytes with 0x",
			hexStr:  "0x" + strings.Repeat("cc", 32),
			want:    mustHexToBytes32(strings.Repeat("cc", 32)),
			wantErr: false,
		},
		{
			name:        "spaces in hex string",
			hexStr:      "0x12 34 56",
			wantErr:     true,
			errContains: "decode hex",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := hexToBytes32(tt.hexStr)
			if tt.wantErr {
				require.Error(t, err)
				if tt.errContains != "" {
					assert.Contains(t, err.Error(), tt.errContains)
				}
				assert.Equal(t, [32]byte{}, got)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestDepositHash_CreateDepositHashABIArguments(t *testing.T) {
	// Test that the cached ABI argument structure is correct
	args := depositHashABIArguments

	require.Len(t, args, 8, "should have 8 arguments")

	// Check types in order
	assert.Equal(t, "uint32", args[0].Type.String(), "sourceDomain should be uint32")
	assert.Equal(t, "uint256", args[1].Type.String(), "amount should be uint256")
	assert.Equal(t, "uint32", args[2].Type.String(), "destinationDomain should be uint32")
	assert.Equal(t, "bytes32", args[3].Type.String(), "mintRecipient should be bytes32")
	assert.Equal(t, "bytes32", args[4].Type.String(), "burnToken should be bytes32")
	assert.Equal(t, "bytes32", args[5].Type.String(), "destinationCaller should be bytes32")
	assert.Equal(t, "uint256", args[6].Type.String(), "maxFee should be uint256")
	assert.Equal(t, "uint32", args[7].Type.String(), "minFinalityThreshold should be uint32")

	// Test that the arguments can be used for packing
	testData, err := args.Pack(
		uint32(1),
		big.NewInt(1000),
		uint32(2),
		[32]byte{1, 2, 3},
		[32]byte{4, 5, 6},
		[32]byte{7, 8, 9},
		big.NewInt(100),
		uint32(3),
	)
	require.NoError(t, err, "should be able to pack test data")
	require.NotEmpty(t, testData, "packed data should not be empty")

	// Verify the packed data length is correct
	// uint32 + uint256 + uint32 + bytes32 + bytes32 + bytes32 + uint256 + uint32
	// = 32 + 32 + 32 + 32 + 32 + 32 + 32 + 32 = 256 bytes
	assert.Equal(t, 256, len(testData), "packed data should be 256 bytes")
}

func TestDepositHash_ParseDepositHashParams(t *testing.T) {
	tests := []struct {
		name        string
		msg         CCTPv2DecodedMessage
		checkResult func(t *testing.T, params depositHashParams, err error)
	}{
		{
			name: "parse all fields correctly",
			msg: CCTPv2DecodedMessage{
				SourceDomain:         "123",
				DestinationDomain:    "456",
				MinFinalityThreshold: "789",
				DestinationCaller:    "0xDEADBEEF",
				DecodedMessageBody: CCTPv2DecodedMessageBody{
					Amount:        "999999",
					MintRecipient: "0xCAFEBABE",
					BurnToken:     "0xFEEDFACE",
					MaxFee:        "111",
				},
			},
			checkResult: func(t *testing.T, params depositHashParams, err error) {
				require.NoError(t, err)
				assert.Equal(t, uint32(123), params.sourceDomain)
				assert.Equal(t, uint32(456), params.destinationDomain)
				assert.Equal(t, uint32(789), params.minFinalityThreshold)
				assert.Equal(t, big.NewInt(999999), params.amount)
				assert.Equal(t, big.NewInt(111), params.maxFee)

				// Check that addresses are properly padded
				expectedMint := mustHexToBytes32("000000000000000000000000000000000000000000000000000000CAFEBABE")
				assert.Equal(t, expectedMint, params.mintRecipient)

				expectedBurn := mustHexToBytes32("000000000000000000000000000000000000000000000000000000FEEDFACE")
				assert.Equal(t, expectedBurn, params.burnToken)

				expectedCaller := mustHexToBytes32("000000000000000000000000000000000000000000000000000000DEADBEEF")
				assert.Equal(t, expectedCaller, params.destinationCaller)
			},
		},
		{
			name: "handle empty maxFee",
			msg: CCTPv2DecodedMessage{
				SourceDomain:         "1",
				DestinationDomain:    "2",
				MinFinalityThreshold: "3",
				DestinationCaller:    "0x00",
				DecodedMessageBody: CCTPv2DecodedMessageBody{
					Amount:        "100",
					MintRecipient: "0x01",
					BurnToken:     "0x02",
					MaxFee:        "",
				},
			},
			checkResult: func(t *testing.T, params depositHashParams, err error) {
				require.NoError(t, err)
				assert.Equal(t, big.NewInt(0), params.maxFee)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			params, err := parseDepositHashParams(tt.msg)
			tt.checkResult(t, params, err)
		})
	}
}

// TestDepositHash_SolidityCompatibility verifies that the Go implementation of CalculateDepositHash
// produces identical results to the Solidity implementation in USDCSourcePoolDataCodec._calculateDepositHash.
//
// This test is critical for ensuring cross-chain compatibility in CCTP V2, where deposit hashes
// must match exactly between on-chain (Solidity) and off-chain (Go) components to properly
// match CCIP messages with Circle's attestations.
//
// The test includes:
//  1. A basic test vector with simple values
//  2. A comprehensive test case that matches the exact parameters used in the Solidity test file:
//     contracts/test/libraries/USDCSourcePoolDataCodec/USDCSourcePoolDataCodec._calculateDepositHash.t.sol
//
// If this test fails, it indicates a breaking change in either the Go or Solidity implementation
// that would cause CCTP V2 transfers to fail in production.
func TestDepositHash_SolidityCompatibility(t *testing.T) {
	tests := []struct {
		name         string
		msg          CCTPv2DecodedMessage
		expectedHash string // From Solidity _calculateDepositHash
	}{
		{
			name: "basic test vector",
			msg: CCTPv2DecodedMessage{
				SourceDomain:         "0",
				DestinationDomain:    "1",
				MinFinalityThreshold: "0",
				DestinationCaller:    "0x0000000000000000000000000000000000000000",
				DecodedMessageBody: CCTPv2DecodedMessageBody{
					Amount:        "1000000",
					MintRecipient: "0x5FC8d32690cc91D4c39d9d3abcBD16989F875707",
					BurnToken:     "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48",
					MaxFee:        "0",
				},
			},
			expectedHash: "4f6e1f3125db0b6602c089812be5e711d3e4afe9f71eb9e841ec09b2502ee6bb",
		},
		{
			name: "Solidity test compatibility - USDCSourcePoolDataCodec._calculateDepositHash.t.sol",
			msg: CCTPv2DecodedMessage{
				SourceDomain:         "1553252", // Test value from Solidity test
				DestinationDomain:    "1",
				MinFinalityThreshold: "2000",
				// These addresses are the actual output from Forge's makeAddr() function
				// in the Solidity test. They were obtained by emitting events in the test.
				DestinationCaller: "0x1B6d84C0A6b7e5eC5Cb5a679e08d200B87c5D41E", // makeAddr("destinationCaller")
				DecodedMessageBody: CCTPv2DecodedMessageBody{
					Amount:        "1000000",                                    // 1e6 in Solidity
					MintRecipient: "0x838d977f3B6b620D1841E12d60329B712F01ED51", // makeAddr("mintRecipient")
					BurnToken:     "0xA635A9970Eee7D14113c6e889689e604660b82e4", // makeAddr("burnToken")
					MaxFee:        "0",
				},
			},
			// This exact hash was emitted by the Solidity test and verifies that our Go
			// implementation produces identical results to the on-chain implementation.
			expectedHash: "d13c921ca6bc2af27d28e734fcd0f8c236d6e45f052823c1787b43d6c22616b8",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotHash, err := CalculateDepositHash(tt.msg)
			require.NoError(t, err)

			wantBytes, err := hex.DecodeString(tt.expectedHash)
			require.NoError(t, err)

			var want [32]byte
			copy(want[:], wantBytes)

			assert.Equal(t, want, gotHash,
				"Solidity compatibility failed\nExpected: %x\nGot: %x",
				want, gotHash)
		})
	}
}
