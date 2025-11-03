package v2

import (
	"context"
	"encoding/binary"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-ccip/execute/tokendata"
	"github.com/smartcontractkit/chainlink-ccip/pkg/reader"
)

func TestDecodeSourceTokenDataPayloadV2(t *testing.T) {
	t.Run("Valid decode with CCTP_VERSION_2_TAG", func(t *testing.T) {
		// Build valid payload: tag(4) + domain(4) + hash(32)
		payload := make([]byte, 40)

		// Version tag (bytes 0-3)
		binary.BigEndian.PutUint32(payload[0:4], CCTP_VERSION_2_TAG)

		// Source domain (bytes 4-7)
		expectedDomain := uint32(0)
		binary.BigEndian.PutUint32(payload[4:8], expectedDomain)

		// Deposit hash (bytes 8-39)
		expectedHash := [32]byte{}
		for i := range expectedHash {
			expectedHash[i] = byte(i)
		}
		copy(payload[8:40], expectedHash[:])

		// Decode
		result, err := DecodeSourceTokenDataPayloadV2(cciptypes.Bytes(payload))
		require.NoError(t, err)
		require.NotNil(t, result)

		// Verify
		assert.Equal(t, expectedDomain, result.SourceDomain)
		assert.Equal(t, expectedHash, result.DepositHash)
	})

	t.Run("Valid decode with CCTP_VERSION_2_CCV_TAG", func(t *testing.T) {
		// Build valid payload with CCV tag
		payload := make([]byte, 40)

		// Version tag (bytes 0-3)
		binary.BigEndian.PutUint32(payload[0:4], CCTP_VERSION_2_CCV_TAG)

		// Source domain (bytes 4-7)
		expectedDomain := uint32(1) // Avalanche
		binary.BigEndian.PutUint32(payload[4:8], expectedDomain)

		// Deposit hash (bytes 8-39)
		expectedHash := [32]byte{}
		for i := range expectedHash {
			expectedHash[i] = byte(255 - i) // Different pattern
		}
		copy(payload[8:40], expectedHash[:])

		// Decode
		result, err := DecodeSourceTokenDataPayloadV2(cciptypes.Bytes(payload))
		require.NoError(t, err)
		require.NotNil(t, result)

		// Verify
		assert.Equal(t, expectedDomain, result.SourceDomain)
		assert.Equal(t, expectedHash, result.DepositHash)
	})

	t.Run("Valid decode with different domain IDs", func(t *testing.T) {
		testCases := []struct {
			name   string
			domain uint32
		}{
			{"Ethereum", 0},
			{"Avalanche", 1},
			{"Optimism", 2},
			{"Arbitrum", 3},
			{"Base", 6},
			{"Polygon", 7},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				payload := make([]byte, 40)
				binary.BigEndian.PutUint32(payload[0:4], CCTP_VERSION_2_TAG)
				binary.BigEndian.PutUint32(payload[4:8], tc.domain)
				// depositHash all zeros is fine

				result, err := DecodeSourceTokenDataPayloadV2(cciptypes.Bytes(payload))
				require.NoError(t, err)
				assert.Equal(t, tc.domain, result.SourceDomain)
			})
		}
	})

	t.Run("Error: Invalid length (too short)", func(t *testing.T) {
		payload := make([]byte, 39) // One byte short
		result, err := DecodeSourceTokenDataPayloadV2(cciptypes.Bytes(payload))
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), "invalid V2 source pool data length")
		assert.Contains(t, err.Error(), "expected 40 bytes, got 39")
	})

	t.Run("Error: Invalid length (too long)", func(t *testing.T) {
		payload := make([]byte, 41) // One byte too long
		result, err := DecodeSourceTokenDataPayloadV2(cciptypes.Bytes(payload))
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), "invalid V2 source pool data length")
		assert.Contains(t, err.Error(), "expected 40 bytes, got 41")
	})

	t.Run("Error: Invalid length (empty)", func(t *testing.T) {
		payload := make([]byte, 0)
		result, err := DecodeSourceTokenDataPayloadV2(cciptypes.Bytes(payload))
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), "invalid V2 source pool data length")
	})

	t.Run("Error: Invalid version tag (V1 tag)", func(t *testing.T) {
		payload := make([]byte, 40)
		// Use CCTP V1 tag instead of V2
		v1Tag := uint32(0xf3567d18) // keccak256("CCTP_V1")
		binary.BigEndian.PutUint32(payload[0:4], v1Tag)

		result, err := DecodeSourceTokenDataPayloadV2(cciptypes.Bytes(payload))
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), "invalid CCTPv2 version tag")
		assert.Contains(t, err.Error(), "0xf3567d18") // Should include actual tag
	})

	t.Run("Error: Invalid version tag (random value)", func(t *testing.T) {
		payload := make([]byte, 40)
		// Use completely invalid tag
		invalidTag := uint32(0xdeadbeef)
		binary.BigEndian.PutUint32(payload[0:4], invalidTag)

		result, err := DecodeSourceTokenDataPayloadV2(cciptypes.Bytes(payload))
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), "invalid CCTPv2 version tag")
		assert.Contains(t, err.Error(), "0xdeadbeef")
	})

	t.Run("Error: Invalid version tag (all zeros)", func(t *testing.T) {
		payload := make([]byte, 40)
		// All zeros, including version tag
		result, err := DecodeSourceTokenDataPayloadV2(cciptypes.Bytes(payload))
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), "invalid CCTPv2 version tag")
	})

	t.Run("DepositHash edge cases", func(t *testing.T) {
		t.Run("All zeros", func(t *testing.T) {
			payload := make([]byte, 40)
			binary.BigEndian.PutUint32(payload[0:4], CCTP_VERSION_2_TAG)
			binary.BigEndian.PutUint32(payload[4:8], 0)
			// depositHash already all zeros

			result, err := DecodeSourceTokenDataPayloadV2(cciptypes.Bytes(payload))
			require.NoError(t, err)
			assert.Equal(t, [32]byte{}, result.DepositHash)
		})

		t.Run("All 0xFF", func(t *testing.T) {
			payload := make([]byte, 40)
			binary.BigEndian.PutUint32(payload[0:4], CCTP_VERSION_2_TAG)
			binary.BigEndian.PutUint32(payload[4:8], 0)
			for i := 8; i < 40; i++ {
				payload[i] = 0xFF
			}

			result, err := DecodeSourceTokenDataPayloadV2(cciptypes.Bytes(payload))
			require.NoError(t, err)

			expectedHash := [32]byte{}
			for i := range expectedHash {
				expectedHash[i] = 0xFF
			}
			assert.Equal(t, expectedHash, result.DepositHash)
		})
	})

	t.Run("Round trip: both tags produce same structure", func(t *testing.T) {
		// Create identical payloads except for version tag
		sourceDomain := uint32(2)
		depositHash := [32]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
			11, 12, 13, 14, 15, 16, 17, 18, 19, 20,
			21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32}

		// Payload with regular V2 tag
		payload1 := make([]byte, 40)
		binary.BigEndian.PutUint32(payload1[0:4], CCTP_VERSION_2_TAG)
		binary.BigEndian.PutUint32(payload1[4:8], sourceDomain)
		copy(payload1[8:40], depositHash[:])

		// Payload with CCV tag
		payload2 := make([]byte, 40)
		binary.BigEndian.PutUint32(payload2[0:4], CCTP_VERSION_2_CCV_TAG)
		binary.BigEndian.PutUint32(payload2[4:8], sourceDomain)
		copy(payload2[8:40], depositHash[:])

		// Decode both
		result1, err1 := DecodeSourceTokenDataPayloadV2(cciptypes.Bytes(payload1))
		require.NoError(t, err1)

		result2, err2 := DecodeSourceTokenDataPayloadV2(cciptypes.Bytes(payload2))
		require.NoError(t, err2)

		// Both should produce identical structures (except we don't store the tag)
		assert.Equal(t, result1.SourceDomain, result2.SourceDomain)
		assert.Equal(t, result1.DepositHash, result2.DepositHash)
	})
}

// Mock encoders for testing
func successEncoder(ctx context.Context, msgBody, attestation cciptypes.Bytes) (cciptypes.Bytes, error) {
	// Simple pass-through encoder that returns the attestation
	return attestation, nil
}

func errorEncoder(ctx context.Context, msgBody, attestation cciptypes.Bytes) (cciptypes.Bytes, error) {
	return nil, errors.New("encoding failed")
}

func transformEncoder(ctx context.Context, msgBody, attestation cciptypes.Bytes) (cciptypes.Bytes, error) {
	// Combines message body and attestation for verification
	result := make(cciptypes.Bytes, 0, len(msgBody)+len(attestation))
	result = append(result, msgBody...)
	result = append(result, attestation...)
	return result, nil
}

func Test_attestationToTokenData(t *testing.T) {
	tests := []struct {
		name           string
		seqNum         cciptypes.SeqNum
		tokenIndex     int
		attestations   map[reader.MessageTokenID]tokendata.AttestationStatus
		encoder        AttestationEncoder
		expectedReady  bool
		expectedError  error
		expectedData   cciptypes.Bytes
	}{
		{
			name:       "Success - valid attestation",
			seqNum:     100,
			tokenIndex: 0,
			attestations: map[reader.MessageTokenID]tokendata.AttestationStatus{
				reader.NewMessageTokenID(100, 0): {
					MessageBody: []byte{0x01, 0x02},
					Attestation: []byte{0xAB, 0xCD},
					Error:       nil,
				},
			},
			encoder:       successEncoder,
			expectedReady: true,
			expectedError: nil,
			expectedData:  []byte{0xAB, 0xCD},
		},
		{
			name:       "Success - transform encoder combines data",
			seqNum:     200,
			tokenIndex: 1,
			attestations: map[reader.MessageTokenID]tokendata.AttestationStatus{
				reader.NewMessageTokenID(200, 1): {
					MessageBody: []byte{0x11, 0x22},
					Attestation: []byte{0x33, 0x44},
					Error:       nil,
				},
			},
			encoder:       transformEncoder,
			expectedReady: true,
			expectedError: nil,
			expectedData:  []byte{0x11, 0x22, 0x33, 0x44},
		},
		{
			name:         "Error - missing attestation",
			seqNum:       100,
			tokenIndex:   0,
			attestations: map[reader.MessageTokenID]tokendata.AttestationStatus{},
			encoder:      successEncoder,
			expectedReady: false,
			expectedError: tokendata.ErrDataMissing,
			expectedData:  nil,
		},
		{
			name:       "Error - attestation has ErrNotReady",
			seqNum:     100,
			tokenIndex: 0,
			attestations: map[reader.MessageTokenID]tokendata.AttestationStatus{
				reader.NewMessageTokenID(100, 0): {
					MessageBody: []byte{0x01, 0x02},
					Attestation: []byte{0xAB, 0xCD},
					Error:       tokendata.ErrNotReady,
				},
			},
			encoder:       successEncoder,
			expectedReady: false,
			expectedError: tokendata.ErrNotReady,
			expectedData:  nil,
		},
		{
			name:       "Error - attestation has custom error",
			seqNum:     100,
			tokenIndex: 0,
			attestations: map[reader.MessageTokenID]tokendata.AttestationStatus{
				reader.NewMessageTokenID(100, 0): {
					MessageBody: []byte{0x01, 0x02},
					Attestation: []byte{0xAB, 0xCD},
					Error:       errors.New("custom API error"),
				},
			},
			encoder:       successEncoder,
			expectedReady: false,
			expectedError: errors.New("custom API error"),
			expectedData:  nil,
		},
		{
			name:       "Error - encoder failure",
			seqNum:     100,
			tokenIndex: 0,
			attestations: map[reader.MessageTokenID]tokendata.AttestationStatus{
				reader.NewMessageTokenID(100, 0): {
					MessageBody: []byte{0x01, 0x02},
					Attestation: []byte{0xAB, 0xCD},
					Error:       nil,
				},
			},
			encoder:       errorEncoder,
			expectedReady: false,
			expectedError: errors.New("unable to encode attestation: encoding failed"),
			expectedData:  nil,
		},
		{
			name:       "Edge case - empty data with success encoder",
			seqNum:     100,
			tokenIndex: 0,
			attestations: map[reader.MessageTokenID]tokendata.AttestationStatus{
				reader.NewMessageTokenID(100, 0): {
					MessageBody: []byte{},
					Attestation: []byte{},
					Error:       nil,
				},
			},
			encoder:       successEncoder,
			expectedReady: true,
			expectedError: nil,
			expectedData:  []byte{},
		},
		{
			name:       "Edge case - different seqNum/tokenIndex combination",
			seqNum:     999,
			tokenIndex: 5,
			attestations: map[reader.MessageTokenID]tokendata.AttestationStatus{
				reader.NewMessageTokenID(999, 5): {
					MessageBody: []byte{0xFF},
					Attestation: []byte{0xEE},
					Error:       nil,
				},
			},
			encoder:       successEncoder,
			expectedReady: true,
			expectedError: nil,
			expectedData:  []byte{0xEE},
		},
		{
			name:       "Error - wrong MessageTokenID in map",
			seqNum:     100,
			tokenIndex: 0,
			attestations: map[reader.MessageTokenID]tokendata.AttestationStatus{
				reader.NewMessageTokenID(100, 1): { // Different tokenIndex
					MessageBody: []byte{0x01, 0x02},
					Attestation: []byte{0xAB, 0xCD},
					Error:       nil,
				},
			},
			encoder:       successEncoder,
			expectedReady: false,
			expectedError: tokendata.ErrDataMissing,
			expectedData:  nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			observer := &CCTPv2TokenDataObserver{
				attestationEncoder: tc.encoder,
			}

			result := observer.attestationToTokenData(
				context.Background(),
				tc.seqNum,
				tc.tokenIndex,
				tc.attestations,
			)

			// Verify Ready flag
			assert.Equal(t, tc.expectedReady, result.Ready, "Ready flag mismatch")

			// Verify Supported flag (should always be true for this function)
			assert.True(t, result.Supported, "Supported should always be true")

			// Verify Data
			assert.Equal(t, tc.expectedData, result.Data, "Data mismatch")

			// Verify Error
			if tc.expectedError != nil {
				require.Error(t, result.Error, "Expected error but got nil")
				assert.Equal(t, tc.expectedError.Error(), result.Error.Error(), "Error message mismatch")
			} else {
				assert.NoError(t, result.Error, "Expected no error but got one")
			}
		})
	}
}
