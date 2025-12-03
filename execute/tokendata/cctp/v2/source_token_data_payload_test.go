package v2

import (
	"encoding/hex"
	"strings"
	"testing"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSourceTokenDataPayload_Decode(t *testing.T) {
	tests := []struct {
		name        string
		extraData   cciptypes.Bytes
		want        *SourceTokenDataPayloadV2
		wantErr     bool
		errContains string
	}{
		{
			name: "valid V2 payload with standard tag",
			extraData: createValidExtraData(
				CCTPVersion2Tag, 1, "1234567890123456789012345678901234567890123456789012345678901234",
			),
			want: &SourceTokenDataPayloadV2{
				SourceDomain: 1,
				DepositHash:  mustHexToBytes32("1234567890123456789012345678901234567890123456789012345678901234"),
			},
			wantErr: false,
		},
		{
			name: "valid V2 payload with CCV tag",
			extraData: createValidExtraData(
				CCTPVersion2CCVTag, 2, "abcdef0123456789abcdef0123456789abcdef0123456789abcdef0123456789",
			),
			want: &SourceTokenDataPayloadV2{
				SourceDomain: 2,
				DepositHash:  mustHexToBytes32("abcdef0123456789abcdef0123456789abcdef0123456789abcdef0123456789"),
			},
			wantErr: false,
		},
		{
			name:        "invalid length - too short",
			extraData:   cciptypes.Bytes(make([]byte, 39)),
			wantErr:     true,
			errContains: "invalid V2 source pool data length: expected 40 bytes, got 39",
		},
		{
			name:        "invalid length - too long",
			extraData:   cciptypes.Bytes(make([]byte, 41)),
			wantErr:     true,
			errContains: "invalid V2 source pool data length: expected 40 bytes, got 41",
		},
		{
			name:        "invalid length - empty",
			extraData:   cciptypes.Bytes{},
			wantErr:     true,
			errContains: "invalid V2 source pool data length: expected 40 bytes, got 0",
		},
		{
			name:        "invalid version tag - v1 tag",
			extraData:   createValidExtraData(0x12345678, 1, "1234567890123456789012345678901234567890123456789012345678901234"),
			wantErr:     true,
			errContains: "invalid CCTPv2 version tag: expected 0xb148ea5f or 0x3047587c, got 0x12345678",
		},
		{
			name:        "invalid version tag - zero",
			extraData:   createValidExtraData(0x00000000, 1, "1234567890123456789012345678901234567890123456789012345678901234"),
			wantErr:     true,
			errContains: "invalid CCTPv2 version tag: expected 0xb148ea5f or 0x3047587c, got 0x0",
		},
		{
			name: "valid with max source domain",
			extraData: createValidExtraData(
				CCTPVersion2Tag, 0xFFFFFFFF, "1234567890123456789012345678901234567890123456789012345678901234",
			),
			want: &SourceTokenDataPayloadV2{
				SourceDomain: 0xFFFFFFFF,
				DepositHash:  mustHexToBytes32("1234567890123456789012345678901234567890123456789012345678901234"),
			},
			wantErr: false,
		},
		{
			name: "valid with zero source domain",
			extraData: createValidExtraData(
				CCTPVersion2Tag, 0, "0000000000000000000000000000000000000000000000000000000000000000",
			),
			want: &SourceTokenDataPayloadV2{
				SourceDomain: 0,
				DepositHash:  [32]byte{},
			},
			wantErr: false,
		},
		{
			name: "valid with all bytes set to 0xFF in depositHash",
			extraData: createValidExtraData(
				CCTPVersion2Tag, 123, "ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff",
			),
			want: &SourceTokenDataPayloadV2{
				SourceDomain: 123,
				DepositHash:  mustHexToBytes32("ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff"),
			},
			wantErr: false,
		},
		{
			name: "valid payload from testnet data",
			extraData: createValidExtraData(
				CCTPVersion2Tag, 6, "d4e39c1e2b2b3f0c8e5a9d7f1234567890abcdef1234567890abcdef12345678",
			),
			want: &SourceTokenDataPayloadV2{
				SourceDomain: 6,
				DepositHash:  mustHexToBytes32("d4e39c1e2b2b3f0c8e5a9d7f1234567890abcdef1234567890abcdef12345678"),
			},
			wantErr: false,
		},
		{
			name:        "invalid payload with V1 structure",
			extraData:   cciptypes.Bytes(append([]byte{0x00, 0x00, 0x00, 0x01}, make([]byte, 36)...)),
			wantErr:     true,
			errContains: "invalid CCTPv2 version tag",
		},
		{
			name: "boundary test - source domain 0",
			extraData: createValidExtraData(
				CCTPVersion2Tag, 0, "a1b2c3d4e5f6789012345678901234567890123456789012345678901234567f",
			),
			want: &SourceTokenDataPayloadV2{
				SourceDomain: 0,
				DepositHash:  mustHexToBytes32("a1b2c3d4e5f6789012345678901234567890123456789012345678901234567f"),
			},
			wantErr: false,
		},
		{
			name: "real world example - Ethereum to Polygon",
			extraData: createValidExtraData(
				CCTPVersion2CCVTag, 0, "8f3a2e1b4c5d6789abcdef0123456789fedcba9876543210123456789abcdef0",
			),
			want: &SourceTokenDataPayloadV2{
				SourceDomain: 0, // Ethereum domain
				DepositHash:  mustHexToBytes32("8f3a2e1b4c5d6789abcdef0123456789fedcba9876543210123456789abcdef0"),
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DecodeSourceTokenDataPayloadV2(tt.extraData)
			if tt.wantErr {
				require.Error(t, err)
				if tt.errContains != "" {
					assert.Contains(t, err.Error(), tt.errContains)
				}
				assert.Nil(t, got)
			} else {
				require.NoError(t, err)
				require.NotNil(t, got)
				assert.Equal(t, tt.want.SourceDomain, got.SourceDomain)
				assert.Equal(t, tt.want.DepositHash, got.DepositHash)
			}
		})
	}
}

// Helper functions

func createValidExtraData(versionTag uint32, sourceDomain uint32, depositHashHex string) cciptypes.Bytes {
	data := make([]byte, 40)

	// Version tag (bytes 0-3)
	data[0] = byte(versionTag >> 24)
	data[1] = byte(versionTag >> 16)
	data[2] = byte(versionTag >> 8)
	data[3] = byte(versionTag)

	// Source domain (bytes 4-7)
	data[4] = byte(sourceDomain >> 24)
	data[5] = byte(sourceDomain >> 16)
	data[6] = byte(sourceDomain >> 8)
	data[7] = byte(sourceDomain)

	// DepositHash (bytes 8-39)
	hashBytes, _ := hex.DecodeString(depositHashHex)
	copy(data[8:40], hashBytes)

	return cciptypes.Bytes(data)
}

func mustHexToBytes32(hexStr string) [32]byte {
	hexStr = strings.TrimPrefix(hexStr, "0x")
	decoded, err := hex.DecodeString(hexStr)
	if err != nil {
		panic(err)
	}

	var result [32]byte
	copy(result[32-len(decoded):], decoded)
	return result
}
