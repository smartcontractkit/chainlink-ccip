package v2

import (
	"context"
	"encoding/binary"
	"errors"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-ccip/execute/exectypes"
	"github.com/smartcontractkit/chainlink-ccip/execute/tokendata"
	"github.com/smartcontractkit/chainlink-ccip/internal"
	"github.com/smartcontractkit/chainlink-ccip/pkg/reader"
)

// MockCCTPv2HTTPClient is a mock implementation of CCTPv2HTTPClient for testing
type MockCCTPv2HTTPClient struct {
	mock.Mock
}

func (m *MockCCTPv2HTTPClient) GetMessages(
	ctx context.Context,
	sourceChain cciptypes.ChainSelector,
	sourceDomainID uint32,
	transactionHash string,
) (CCTPv2Messages, error) {
	args := m.Called(ctx, sourceChain, sourceDomainID, transactionHash)
	return args.Get(0).(CCTPv2Messages), args.Error(1)
}

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

// Helper function to create a valid CCTPv2Message with all required fields for depositHash calculation
func createValidCCTPv2Message(message, attestation string) CCTPv2Message {
	return CCTPv2Message{
		Message:     message,
		Attestation: attestation,
		DecodedMessage: CCTPv2DecodedMessage{
			SourceDomain:      "0",
			DestinationDomain: "1",
			Recipient:         "0x1234567890123456789012345678901234567890",
			DestinationCaller: "0x0000000000000000000000000000000000000000",
			MessageBody:       "0x",
			DecodedMessageBody: CCTPv2DecodedMessageBody{
				BurnToken:     "0x0000000000000000000000001111111111111111",
				MintRecipient: "0x0000000000000000000000002222222222222222",
				Amount:        "1000000",
				MessageSender: "0x3333333333333333333333333333333333333333",
				MaxFee:        "0",
			},
			MinFinalityThreshold: "2000",
		},
	}
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
		name          string
		seqNum        cciptypes.SeqNum
		tokenIndex    int
		attestations  map[reader.MessageTokenID]tokendata.AttestationStatus
		encoder       AttestationEncoder
		expectedReady bool
		expectedError error
		expectedData  cciptypes.Bytes
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
			name:          "Error - missing attestation",
			seqNum:        100,
			tokenIndex:    0,
			attestations:  map[reader.MessageTokenID]tokendata.AttestationStatus{},
			encoder:       successEncoder,
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

// Helper functions for test data
func newReadyTokenData(data []byte) exectypes.TokenData {
	return exectypes.TokenData{
		Ready:     true,
		Error:     nil,
		Data:      data,
		Supported: true,
	}
}

func newErrorTokenData(err error) exectypes.TokenData {
	return exectypes.TokenData{
		Ready:     false,
		Error:     err,
		Data:      nil,
		Supported: true,
	}
}

func Test_extractTokenData(t *testing.T) {
	// Setup supported pools for testing
	ethereumChain := cciptypes.ChainSelector(1)
	avalancheChain := cciptypes.ChainSelector(2)
	ethereumUSDCPool := "0x1111111111111111111111111111111111111111"
	avalancheUSDCPool := "0x2222222222222222222222222222222222222222"
	unsupportedPool1 := "0x3333333333333333333333333333333333333333"
	unsupportedPool2 := "0x4444444444444444444444444444444444444444"

	supportedPoolsBySelector := map[cciptypes.ChainSelector]string{
		ethereumChain:  ethereumUSDCPool,
		avalancheChain: avalancheUSDCPool,
	}

	tests := []struct {
		name              string
		messages          exectypes.MessageObservations
		attestations      map[cciptypes.ChainSelector]map[reader.MessageTokenID]tokendata.AttestationStatus
		encoder           AttestationEncoder
		expectedTokenData exectypes.TokenDataObservations
	}{
		{
			name:     "no messages",
			messages: exectypes.MessageObservations{},
			attestations: map[cciptypes.ChainSelector]map[reader.MessageTokenID]tokendata.AttestationStatus{
				ethereumChain: {},
			},
			encoder:           successEncoder,
			expectedTokenData: exectypes.TokenDataObservations{},
		},
		{
			name: "no supported tokens",
			messages: exectypes.MessageObservations{
				ethereumChain: {
					cciptypes.SeqNum(10): internal.MessageWithTokens(t, unsupportedPool1, unsupportedPool2),
				},
			},
			attestations: nil,
			encoder:      successEncoder,
			expectedTokenData: exectypes.TokenDataObservations{
				ethereumChain: {
					cciptypes.SeqNum(10): exectypes.NewMessageTokenData(
						exectypes.NotSupportedTokenData(),
						exectypes.NotSupportedTokenData(),
					),
				},
			},
		},
		{
			name: "empty token amounts",
			messages: exectypes.MessageObservations{
				ethereumChain: {
					cciptypes.SeqNum(10): internal.MessageWithTokens(t),
				},
			},
			attestations: nil,
			encoder:      successEncoder,
			expectedTokenData: exectypes.TokenDataObservations{
				ethereumChain: {
					cciptypes.SeqNum(10): exectypes.NewMessageTokenData(),
				},
			},
		},
		{
			name: "single supported token with attestation",
			messages: exectypes.MessageObservations{
				ethereumChain: {
					cciptypes.SeqNum(10): internal.MessageWithTokens(t, ethereumUSDCPool),
				},
			},
			attestations: map[cciptypes.ChainSelector]map[reader.MessageTokenID]tokendata.AttestationStatus{
				ethereumChain: {
					reader.NewMessageTokenID(10, 0): {
						MessageBody: []byte{0x01, 0x02},
						Attestation: []byte{0xAB, 0xCD},
						Error:       nil,
					},
				},
			},
			encoder: successEncoder,
			expectedTokenData: exectypes.TokenDataObservations{
				ethereumChain: {
					cciptypes.SeqNum(10): exectypes.NewMessageTokenData(
						newReadyTokenData([]byte{0xAB, 0xCD}),
					),
				},
			},
		},
		{
			name: "multiple chains with supported tokens",
			messages: exectypes.MessageObservations{
				ethereumChain: {
					cciptypes.SeqNum(10): internal.MessageWithTokens(t, ethereumUSDCPool),
				},
				avalancheChain: {
					cciptypes.SeqNum(20): internal.MessageWithTokens(t, avalancheUSDCPool),
				},
			},
			attestations: map[cciptypes.ChainSelector]map[reader.MessageTokenID]tokendata.AttestationStatus{
				ethereumChain: {
					reader.NewMessageTokenID(10, 0): {
						MessageBody: []byte{0x01},
						Attestation: []byte{0xAA},
						Error:       nil,
					},
				},
				avalancheChain: {
					reader.NewMessageTokenID(20, 0): {
						MessageBody: []byte{0x02},
						Attestation: []byte{0xBB},
						Error:       nil,
					},
				},
			},
			encoder: successEncoder,
			expectedTokenData: exectypes.TokenDataObservations{
				ethereumChain: {
					cciptypes.SeqNum(10): exectypes.NewMessageTokenData(
						newReadyTokenData([]byte{0xAA}),
					),
				},
				avalancheChain: {
					cciptypes.SeqNum(20): exectypes.NewMessageTokenData(
						newReadyTokenData([]byte{0xBB}),
					),
				},
			},
		},
		{
			name: "mixed supported and unsupported tokens in single message",
			messages: exectypes.MessageObservations{
				ethereumChain: {
					cciptypes.SeqNum(10): internal.MessageWithTokens(t, unsupportedPool1, ethereumUSDCPool, unsupportedPool2),
				},
			},
			attestations: map[cciptypes.ChainSelector]map[reader.MessageTokenID]tokendata.AttestationStatus{
				ethereumChain: {
					reader.NewMessageTokenID(10, 1): {
						MessageBody: []byte{0x01, 0x02},
						Attestation: []byte{0xCC, 0xDD},
						Error:       nil,
					},
				},
			},
			encoder: successEncoder,
			expectedTokenData: exectypes.TokenDataObservations{
				ethereumChain: {
					cciptypes.SeqNum(10): exectypes.NewMessageTokenData(
						exectypes.NotSupportedTokenData(),
						newReadyTokenData([]byte{0xCC, 0xDD}),
						exectypes.NotSupportedTokenData(),
					),
				},
			},
		},
		{
			name: "multiple supported tokens in single message",
			messages: exectypes.MessageObservations{
				ethereumChain: {
					cciptypes.SeqNum(10): internal.MessageWithTokens(t, ethereumUSDCPool, ethereumUSDCPool),
				},
			},
			attestations: map[cciptypes.ChainSelector]map[reader.MessageTokenID]tokendata.AttestationStatus{
				ethereumChain: {
					reader.NewMessageTokenID(10, 0): {
						MessageBody: []byte{0x01},
						Attestation: []byte{0xAA},
						Error:       nil,
					},
					reader.NewMessageTokenID(10, 1): {
						MessageBody: []byte{0x02},
						Attestation: []byte{0xBB},
						Error:       nil,
					},
				},
			},
			encoder: successEncoder,
			expectedTokenData: exectypes.TokenDataObservations{
				ethereumChain: {
					cciptypes.SeqNum(10): exectypes.NewMessageTokenData(
						newReadyTokenData([]byte{0xAA}),
						newReadyTokenData([]byte{0xBB}),
					),
				},
			},
		},
		{
			name: "multiple messages per chain",
			messages: exectypes.MessageObservations{
				ethereumChain: {
					cciptypes.SeqNum(10): internal.MessageWithTokens(t, ethereumUSDCPool),
					cciptypes.SeqNum(11): internal.MessageWithTokens(t, ethereumUSDCPool),
					cciptypes.SeqNum(12): internal.MessageWithTokens(t, unsupportedPool1),
				},
			},
			attestations: map[cciptypes.ChainSelector]map[reader.MessageTokenID]tokendata.AttestationStatus{
				ethereumChain: {
					reader.NewMessageTokenID(10, 0): {
						MessageBody: []byte{0x01},
						Attestation: []byte{0xAA},
						Error:       nil,
					},
					reader.NewMessageTokenID(11, 0): {
						MessageBody: []byte{0x02},
						Attestation: []byte{0xBB},
						Error:       nil,
					},
				},
			},
			encoder: successEncoder,
			expectedTokenData: exectypes.TokenDataObservations{
				ethereumChain: {
					cciptypes.SeqNum(10): exectypes.NewMessageTokenData(
						newReadyTokenData([]byte{0xAA}),
					),
					cciptypes.SeqNum(11): exectypes.NewMessageTokenData(
						newReadyTokenData([]byte{0xBB}),
					),
					cciptypes.SeqNum(12): exectypes.NewMessageTokenData(
						exectypes.NotSupportedTokenData(),
					),
				},
			},
		},
		{
			name: "nil attestations map",
			messages: exectypes.MessageObservations{
				ethereumChain: {
					cciptypes.SeqNum(10): internal.MessageWithTokens(t, ethereumUSDCPool),
				},
			},
			attestations: nil,
			encoder:      successEncoder,
			expectedTokenData: exectypes.TokenDataObservations{
				ethereumChain: {
					cciptypes.SeqNum(10): exectypes.NewMessageTokenData(
						newErrorTokenData(tokendata.ErrDataMissing),
					),
				},
			},
		},
		{
			name: "missing attestation for supported token",
			messages: exectypes.MessageObservations{
				ethereumChain: {
					cciptypes.SeqNum(10): internal.MessageWithTokens(t, ethereumUSDCPool),
				},
			},
			attestations: map[cciptypes.ChainSelector]map[reader.MessageTokenID]tokendata.AttestationStatus{
				ethereumChain: {},
			},
			encoder: successEncoder,
			expectedTokenData: exectypes.TokenDataObservations{
				ethereumChain: {
					cciptypes.SeqNum(10): exectypes.NewMessageTokenData(
						newErrorTokenData(tokendata.ErrDataMissing),
					),
				},
			},
		},
		{
			name: "attestation with ErrNotReady",
			messages: exectypes.MessageObservations{
				ethereumChain: {
					cciptypes.SeqNum(10): internal.MessageWithTokens(t, ethereumUSDCPool),
				},
			},
			attestations: map[cciptypes.ChainSelector]map[reader.MessageTokenID]tokendata.AttestationStatus{
				ethereumChain: {
					reader.NewMessageTokenID(10, 0): {
						MessageBody: []byte{0x01, 0x02},
						Attestation: []byte{0xAB, 0xCD},
						Error:       tokendata.ErrNotReady,
					},
				},
			},
			encoder: successEncoder,
			expectedTokenData: exectypes.TokenDataObservations{
				ethereumChain: {
					cciptypes.SeqNum(10): exectypes.NewMessageTokenData(
						newErrorTokenData(tokendata.ErrNotReady),
					),
				},
			},
		},
		{
			name: "attestation with custom error",
			messages: exectypes.MessageObservations{
				ethereumChain: {
					cciptypes.SeqNum(10): internal.MessageWithTokens(t, ethereumUSDCPool),
				},
			},
			attestations: map[cciptypes.ChainSelector]map[reader.MessageTokenID]tokendata.AttestationStatus{
				ethereumChain: {
					reader.NewMessageTokenID(10, 0): {
						MessageBody: []byte{0x01, 0x02},
						Attestation: []byte{0xAB, 0xCD},
						Error:       errors.New("API rate limited"),
					},
				},
			},
			encoder: successEncoder,
			expectedTokenData: exectypes.TokenDataObservations{
				ethereumChain: {
					cciptypes.SeqNum(10): exectypes.NewMessageTokenData(
						newErrorTokenData(errors.New("API rate limited")),
					),
				},
			},
		},
		{
			name: "encoder failure",
			messages: exectypes.MessageObservations{
				ethereumChain: {
					cciptypes.SeqNum(10): internal.MessageWithTokens(t, ethereumUSDCPool),
				},
			},
			attestations: map[cciptypes.ChainSelector]map[reader.MessageTokenID]tokendata.AttestationStatus{
				ethereumChain: {
					reader.NewMessageTokenID(10, 0): {
						MessageBody: []byte{0x01, 0x02},
						Attestation: []byte{0xAB, 0xCD},
						Error:       nil,
					},
				},
			},
			encoder: errorEncoder,
			expectedTokenData: exectypes.TokenDataObservations{
				ethereumChain: {
					cciptypes.SeqNum(10): exectypes.NewMessageTokenData(
						newErrorTokenData(errors.New("unable to encode attestation: encoding failed")),
					),
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			observer := &CCTPv2TokenDataObserver{
				lggr:                     logger.Test(t),
				destChainSelector:        ethereumChain,
				supportedPoolsBySelector: supportedPoolsBySelector,
				attestationEncoder:       tc.encoder,
			}

			result, err := observer.extractTokenData(
				context.Background(),
				logger.Test(t),
				tc.messages,
				tc.attestations,
			)

			require.NoError(t, err)
			assert.Equal(t, len(tc.expectedTokenData), len(result), "Number of chains mismatch")

			for chainSelector, expectedChainData := range tc.expectedTokenData {
				actualChainData, ok := result[chainSelector]
				require.True(t, ok, "Chain %d not found in result", chainSelector)
				assert.Equal(t, len(expectedChainData), len(actualChainData), "Number of messages mismatch for chain %d", chainSelector)

				for seqNum, expectedMsgData := range expectedChainData {
					actualMsgData, ok := actualChainData[seqNum]
					require.True(t, ok, "SeqNum %d not found in result for chain %d", seqNum, chainSelector)
					assert.Equal(t, len(expectedMsgData.TokenData), len(actualMsgData.TokenData), "Number of tokens mismatch for seqNum %d", seqNum)

					for i, expectedToken := range expectedMsgData.TokenData {
						actualToken := actualMsgData.TokenData[i]
						assert.Equal(t, expectedToken.Ready, actualToken.Ready, "Ready mismatch for token %d", i)
						assert.Equal(t, expectedToken.Supported, actualToken.Supported, "Supported mismatch for token %d", i)
						assert.Equal(t, expectedToken.Data, actualToken.Data, "Data mismatch for token %d", i)

						if expectedToken.Error != nil {
							require.Error(t, actualToken.Error, "Expected error for token %d", i)
							assert.Equal(t, expectedToken.Error.Error(), actualToken.Error.Error(), "Error message mismatch for token %d", i)
						} else {
							assert.NoError(t, actualToken.Error, "Unexpected error for token %d", i)
						}
					}
				}
			}
		})
	}
}

func Test_buildAttestationStatuses(t *testing.T) {
	lggr := logger.Test(t)
	observer := NewCCTPv2TokenDataObserver(
		lggr,
		1, // destChainSelector
		map[cciptypes.ChainSelector]string{
			1: "0x1111111111111111111111111111111111111111",
		},
		successEncoder,
		nil,
	)

	t.Run("Success - valid depositHash match", func(t *testing.T) {
		// Create valid CCTP message with all required fields
		cctpMsg := CCTPv2Message{
			Message:     "0xaabbccdd",
			Attestation: "0x11223344",
			DecodedMessage: CCTPv2DecodedMessage{
				SourceDomain:      "0",
				DestinationDomain: "1",
				Recipient:         "0x1234567890123456789012345678901234567890",
				DestinationCaller: "0x0000000000000000000000000000000000000000",
				MessageBody:       "0x",
				DecodedMessageBody: CCTPv2DecodedMessageBody{
					BurnToken:     "0x0000000000000000000000001111111111111111",
					MintRecipient: "0x0000000000000000000000002222222222222222",
					Amount:        "1000000",
					MessageSender: "0x3333333333333333333333333333333333333333",
					MaxFee:        "0",
				},
				MinFinalityThreshold: "2000",
			},
		}

		// Calculate expected depositHash
		expectedHash, err := calculateDepositHash(cctpMsg.DecodedMessage)
		require.NoError(t, err)

		v2Msg := &SourceTokenDataPayloadV2{
			SourceDomain: 0,
			DepositHash:  expectedHash,
		}

		msgTokenID := reader.NewMessageTokenID(1, 0)
		cctpMessages := map[cciptypes.ChainSelector]map[reader.MessageTokenID]CCTPv2Message{
			1: {msgTokenID: cctpMsg},
		}
		v2Messages := map[cciptypes.ChainSelector]map[reader.MessageTokenID]*SourceTokenDataPayloadV2{
			1: {msgTokenID: v2Msg},
		}

		result := observer.buildAttestationStatuses(cctpMessages, v2Messages)

		require.Contains(t, result, cciptypes.ChainSelector(1))
		require.Contains(t, result[1], msgTokenID)
		status := result[1][msgTokenID]
		assert.NoError(t, status.Error)
		assert.Equal(t, expectedHash[:], []byte(status.ID))
		assert.Equal(t, []byte{0xaa, 0xbb, 0xcc, 0xdd}, []byte(status.MessageBody))
		assert.Equal(t, []byte{0x11, 0x22, 0x33, 0x44}, []byte(status.Attestation))
	})

	t.Run("Error - missing CCTP message", func(t *testing.T) {
		expectedHash := [32]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16,
			17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32}
		v2Msg := &SourceTokenDataPayloadV2{
			SourceDomain: 0,
			DepositHash:  expectedHash,
		}

		msgTokenID := reader.NewMessageTokenID(1, 0)
		cctpMessages := map[cciptypes.ChainSelector]map[reader.MessageTokenID]CCTPv2Message{
			1: {}, // Empty - no CCTP message
		}
		v2Messages := map[cciptypes.ChainSelector]map[reader.MessageTokenID]*SourceTokenDataPayloadV2{
			1: {msgTokenID: v2Msg},
		}

		result := observer.buildAttestationStatuses(cctpMessages, v2Messages)

		require.Contains(t, result, cciptypes.ChainSelector(1))
		require.Contains(t, result[1], msgTokenID)
		status := result[1][msgTokenID]
		require.Error(t, status.Error)
		assert.ErrorIs(t, status.Error, tokendata.ErrDataMissing)
	})

	t.Run("Error - invalid hex in message", func(t *testing.T) {
		cctpMsg := CCTPv2Message{
			Message:     "0xINVALIDHEX", // Invalid hex
			Attestation: "0x11223344",
			DecodedMessage: CCTPv2DecodedMessage{
				SourceDomain:      "0",
				DestinationDomain: "1",
				Recipient:         "0x1234567890123456789012345678901234567890",
				DestinationCaller: "0x0000000000000000000000000000000000000000",
				DecodedMessageBody: CCTPv2DecodedMessageBody{
					BurnToken:     "0x0000000000000000000000001111111111111111",
					MintRecipient: "0x0000000000000000000000002222222222222222",
					Amount:        "1000000",
					MaxFee:        "0",
				},
				MinFinalityThreshold: "2000",
			},
		}

		// Calculate expected depositHash
		expectedHash, err := calculateDepositHash(cctpMsg.DecodedMessage)
		require.NoError(t, err)

		v2Msg := &SourceTokenDataPayloadV2{
			SourceDomain: 0,
			DepositHash:  expectedHash,
		}

		msgTokenID := reader.NewMessageTokenID(1, 0)
		cctpMessages := map[cciptypes.ChainSelector]map[reader.MessageTokenID]CCTPv2Message{
			1: {msgTokenID: cctpMsg},
		}
		v2Messages := map[cciptypes.ChainSelector]map[reader.MessageTokenID]*SourceTokenDataPayloadV2{
			1: {msgTokenID: v2Msg},
		}

		result := observer.buildAttestationStatuses(cctpMessages, v2Messages)

		require.Contains(t, result, cciptypes.ChainSelector(1))
		require.Contains(t, result[1], msgTokenID)
		status := result[1][msgTokenID]
		require.Error(t, status.Error)
		assert.Contains(t, status.Error.Error(), "decode message hex")
	})

	t.Run("Error - invalid hex in attestation", func(t *testing.T) {
		cctpMsg := CCTPv2Message{
			Message:     "0xaabbccdd",
			Attestation: "0xINVALIDHEX", // Invalid hex
			DecodedMessage: CCTPv2DecodedMessage{
				SourceDomain:      "0",
				DestinationDomain: "1",
				Recipient:         "0x1234567890123456789012345678901234567890",
				DestinationCaller: "0x0000000000000000000000000000000000000000",
				DecodedMessageBody: CCTPv2DecodedMessageBody{
					BurnToken:     "0x0000000000000000000000001111111111111111",
					MintRecipient: "0x0000000000000000000000002222222222222222",
					Amount:        "1000000",
					MaxFee:        "0",
				},
				MinFinalityThreshold: "2000",
			},
		}

		// Calculate expected depositHash
		expectedHash, err := calculateDepositHash(cctpMsg.DecodedMessage)
		require.NoError(t, err)

		v2Msg := &SourceTokenDataPayloadV2{
			SourceDomain: 0,
			DepositHash:  expectedHash,
		}

		msgTokenID := reader.NewMessageTokenID(1, 0)
		cctpMessages := map[cciptypes.ChainSelector]map[reader.MessageTokenID]CCTPv2Message{
			1: {msgTokenID: cctpMsg},
		}
		v2Messages := map[cciptypes.ChainSelector]map[reader.MessageTokenID]*SourceTokenDataPayloadV2{
			1: {msgTokenID: v2Msg},
		}

		result := observer.buildAttestationStatuses(cctpMessages, v2Messages)

		require.Contains(t, result, cciptypes.ChainSelector(1))
		require.Contains(t, result[1], msgTokenID)
		status := result[1][msgTokenID]
		require.Error(t, status.Error)
		assert.Contains(t, status.Error.Error(), "decode attestation hex")
	})

	t.Run("Multiple chains and messages", func(t *testing.T) {
		// Chain 1, Message 1
		cctpMsg1 := CCTPv2Message{
			Message:     "0xaa",
			Attestation: "0xbb",
			DecodedMessage: CCTPv2DecodedMessage{
				SourceDomain:      "0",
				DestinationDomain: "1",
				Recipient:         "0x1111111111111111111111111111111111111111",
				DestinationCaller: "0x0000000000000000000000000000000000000000",
				DecodedMessageBody: CCTPv2DecodedMessageBody{
					BurnToken:     "0x0000000000000000000000001111111111111111",
					MintRecipient: "0x0000000000000000000000002222222222222222",
					Amount:        "1000",
					MaxFee:        "0",
				},
				MinFinalityThreshold: "2000",
			},
		}
		hash1, err := calculateDepositHash(cctpMsg1.DecodedMessage)
		require.NoError(t, err)

		// Chain 2, Message 2
		cctpMsg2 := CCTPv2Message{
			Message:     "0xcc",
			Attestation: "0xdd",
			DecodedMessage: CCTPv2DecodedMessage{
				SourceDomain:      "1",
				DestinationDomain: "2",
				Recipient:         "0x2222222222222222222222222222222222222222",
				DestinationCaller: "0x0000000000000000000000000000000000000000",
				DecodedMessageBody: CCTPv2DecodedMessageBody{
					BurnToken:     "0x0000000000000000000000003333333333333333",
					MintRecipient: "0x0000000000000000000000004444444444444444",
					Amount:        "2000",
					MaxFee:        "0",
				},
				MinFinalityThreshold: "2000",
			},
		}
		hash2, err := calculateDepositHash(cctpMsg2.DecodedMessage)
		require.NoError(t, err)

		msgTokenID1 := reader.NewMessageTokenID(1, 0)
		msgTokenID2 := reader.NewMessageTokenID(2, 0)

		cctpMessages := map[cciptypes.ChainSelector]map[reader.MessageTokenID]CCTPv2Message{
			1: {msgTokenID1: cctpMsg1},
			2: {msgTokenID2: cctpMsg2},
		}
		v2Messages := map[cciptypes.ChainSelector]map[reader.MessageTokenID]*SourceTokenDataPayloadV2{
			1: {msgTokenID1: &SourceTokenDataPayloadV2{SourceDomain: 0, DepositHash: hash1}},
			2: {msgTokenID2: &SourceTokenDataPayloadV2{SourceDomain: 1, DepositHash: hash2}},
		}

		result := observer.buildAttestationStatuses(cctpMessages, v2Messages)

		// Verify both chains are in result
		require.Contains(t, result, cciptypes.ChainSelector(1))
		require.Contains(t, result, cciptypes.ChainSelector(2))

		// Verify chain 1 message
		require.Contains(t, result[1], msgTokenID1)
		status1 := result[1][msgTokenID1]
		assert.NoError(t, status1.Error)
		assert.Equal(t, hash1[:], []byte(status1.ID))

		// Verify chain 2 message
		require.Contains(t, result[2], msgTokenID2)
		status2 := result[2][msgTokenID2]
		assert.NoError(t, status2.Error)
		assert.Equal(t, hash2[:], []byte(status2.ID))
	})
}

func Test_fetchCCTPv2Attestations(t *testing.T) {
	ctx := context.Background()
	lggr := logger.Test(t)

	t.Run("Success - single transaction with single message", func(t *testing.T) {
		mockClient := new(MockCCTPv2HTTPClient)
		observer := NewCCTPv2TokenDataObserver(
			lggr,
			1,
			map[cciptypes.ChainSelector]string{1: "0x1111"},
			successEncoder,
			mockClient,
		)

		txHash := "0x" + strings.Repeat("a", 64)
		msgTokenID := reader.NewMessageTokenID(10, 0)

		// Create valid CCTP message with all required fields
		expectedMsg := createValidCCTPv2Message("0xaabbccdd", "0x11223344")

		// Calculate depositHash for this message
		expectedHash, err := calculateDepositHash(expectedMsg.DecodedMessage)
		require.NoError(t, err)

		// Create v2Messages with the calculated depositHash
		v2Messages := map[cciptypes.ChainSelector]map[reader.MessageTokenID]*SourceTokenDataPayloadV2{
			1: {
				msgTokenID: &SourceTokenDataPayloadV2{
					SourceDomain: 0,
					DepositHash:  expectedHash,
				},
			},
		}

		mockClient.On("GetMessages", ctx, cciptypes.ChainSelector(1), uint32(0), txHash).
			Return(CCTPv2Messages{Messages: []CCTPv2Message{expectedMsg}}, nil)

		txGroups := map[cciptypes.ChainSelector]map[TxKey][]reader.MessageTokenID{
			1: {
				{SourceDomain: 0, TxHash: txHash}: {msgTokenID},
			},
		}

		result, err := observer.fetchCCTPv2Attestations(ctx, txGroups, v2Messages)

		require.NoError(t, err)
		require.Contains(t, result, cciptypes.ChainSelector(1))
		require.Contains(t, result[1], msgTokenID)
		assert.Equal(t, expectedMsg, result[1][msgTokenID])
		mockClient.AssertExpectations(t)
	})

	t.Run("Success - single transaction with multiple messages", func(t *testing.T) {
		mockClient := new(MockCCTPv2HTTPClient)
		observer := NewCCTPv2TokenDataObserver(
			lggr,
			1,
			map[cciptypes.ChainSelector]string{1: "0x1111"},
			successEncoder,
			mockClient,
		)

		txHash := "0x" + strings.Repeat("b", 64)
		msgTokenID1 := reader.NewMessageTokenID(10, 0)
		msgTokenID2 := reader.NewMessageTokenID(10, 1)

		// Create valid CCTP messages with all required fields
		msg1 := createValidCCTPv2Message("0x01", "0x01")
		msg2 := createValidCCTPv2Message("0x02", "0x02")
		// Make msg2 different by changing amount
		msg2.DecodedMessage.DecodedMessageBody.Amount = "2000000"

		// Calculate depositHashes for each message
		hash1, err := calculateDepositHash(msg1.DecodedMessage)
		require.NoError(t, err)
		hash2, err := calculateDepositHash(msg2.DecodedMessage)
		require.NoError(t, err)

		// Create v2Messages with calculated depositHashes
		v2Messages := map[cciptypes.ChainSelector]map[reader.MessageTokenID]*SourceTokenDataPayloadV2{
			1: {
				msgTokenID1: &SourceTokenDataPayloadV2{
					SourceDomain: 0,
					DepositHash:  hash1,
				},
				msgTokenID2: &SourceTokenDataPayloadV2{
					SourceDomain: 0,
					DepositHash:  hash2,
				},
			},
		}

		mockClient.On("GetMessages", ctx, cciptypes.ChainSelector(1), uint32(0), txHash).
			Return(CCTPv2Messages{Messages: []CCTPv2Message{msg1, msg2}}, nil)

		txGroups := map[cciptypes.ChainSelector]map[TxKey][]reader.MessageTokenID{
			1: {
				{SourceDomain: 0, TxHash: txHash}: {msgTokenID1, msgTokenID2},
			},
		}

		result, err := observer.fetchCCTPv2Attestations(ctx, txGroups, v2Messages)

		require.NoError(t, err)
		require.Contains(t, result[1], msgTokenID1)
		require.Contains(t, result[1], msgTokenID2)
		assert.Equal(t, msg1, result[1][msgTokenID1])
		assert.Equal(t, msg2, result[1][msgTokenID2])
		mockClient.AssertExpectations(t)
	})

	t.Run("Success - multiple transactions across multiple chains", func(t *testing.T) {
		mockClient := new(MockCCTPv2HTTPClient)
		observer := NewCCTPv2TokenDataObserver(
			lggr,
			1,
			map[cciptypes.ChainSelector]string{1: "0x1111", 2: "0x2222"},
			successEncoder,
			mockClient,
		)

		txHash1 := "0x" + strings.Repeat("c", 64)
		txHash2 := "0x" + strings.Repeat("d", 64)
		msgTokenID1 := reader.NewMessageTokenID(10, 0)
		msgTokenID2 := reader.NewMessageTokenID(20, 0)

		// Create valid CCTP messages with all required fields
		msg1 := createValidCCTPv2Message("0xaa", "0xaa")
		msg2 := createValidCCTPv2Message("0xbb", "0xbb")
		// Make msg2 different by changing amount and destination domain
		msg2.DecodedMessage.DecodedMessageBody.Amount = "3000000"
		msg2.DecodedMessage.DestinationDomain = "2"

		// Calculate depositHashes for each message
		hash1, err := calculateDepositHash(msg1.DecodedMessage)
		require.NoError(t, err)
		hash2, err := calculateDepositHash(msg2.DecodedMessage)
		require.NoError(t, err)

		// Create v2Messages with calculated depositHashes
		v2Messages := map[cciptypes.ChainSelector]map[reader.MessageTokenID]*SourceTokenDataPayloadV2{
			1: {
				msgTokenID1: &SourceTokenDataPayloadV2{
					SourceDomain: 0,
					DepositHash:  hash1,
				},
			},
			2: {
				msgTokenID2: &SourceTokenDataPayloadV2{
					SourceDomain: 1,
					DepositHash:  hash2,
				},
			},
		}

		mockClient.On("GetMessages", ctx, cciptypes.ChainSelector(1), uint32(0), txHash1).
			Return(CCTPv2Messages{Messages: []CCTPv2Message{msg1}}, nil)
		mockClient.On("GetMessages", ctx, cciptypes.ChainSelector(2), uint32(1), txHash2).
			Return(CCTPv2Messages{Messages: []CCTPv2Message{msg2}}, nil)

		txGroups := map[cciptypes.ChainSelector]map[TxKey][]reader.MessageTokenID{
			1: {{SourceDomain: 0, TxHash: txHash1}: {msgTokenID1}},
			2: {{SourceDomain: 1, TxHash: txHash2}: {msgTokenID2}},
		}

		result, err := observer.fetchCCTPv2Attestations(ctx, txGroups, v2Messages)

		require.NoError(t, err)
		require.Contains(t, result, cciptypes.ChainSelector(1))
		require.Contains(t, result, cciptypes.ChainSelector(2))
		assert.Equal(t, msg1, result[1][msgTokenID1])
		assert.Equal(t, msg2, result[2][msgTokenID2])
		mockClient.AssertExpectations(t)
	})

	t.Run("Error - HTTP client returns error", func(t *testing.T) {
		mockClient := new(MockCCTPv2HTTPClient)
		observer := NewCCTPv2TokenDataObserver(
			lggr,
			1,
			map[cciptypes.ChainSelector]string{1: "0x1111"},
			successEncoder,
			mockClient,
		)

		txHash := "0x" + strings.Repeat("e", 64)
		msgTokenID := reader.NewMessageTokenID(10, 0)

		// Create v2Messages (won't be matched due to API error)
		v2Messages := map[cciptypes.ChainSelector]map[reader.MessageTokenID]*SourceTokenDataPayloadV2{
			1: {
				msgTokenID: &SourceTokenDataPayloadV2{
					SourceDomain: 0,
					DepositHash:  [32]byte{1, 2, 3},
				},
			},
		}

		// Mock API error
		mockClient.On("GetMessages", ctx, cciptypes.ChainSelector(1), uint32(0), txHash).
			Return(CCTPv2Messages{}, errors.New("API error"))

		txGroups := map[cciptypes.ChainSelector]map[TxKey][]reader.MessageTokenID{
			1: {{SourceDomain: 0, TxHash: txHash}: {msgTokenID}},
		}

		result, err := observer.fetchCCTPv2Attestations(ctx, txGroups, v2Messages)

		require.NoError(t, err) // Function doesn't fail on API errors
		require.Contains(t, result, cciptypes.ChainSelector(1))
		// Message not in result - will be marked as ErrDataMissing in buildAttestationStatuses
		assert.NotContains(t, result[1], msgTokenID)
		mockClient.AssertExpectations(t)
	})

	t.Run("Edge case - fewer messages than MessageTokenIDs", func(t *testing.T) {
		mockClient := new(MockCCTPv2HTTPClient)
		observer := NewCCTPv2TokenDataObserver(
			lggr,
			1,
			map[cciptypes.ChainSelector]string{1: "0x1111"},
			successEncoder,
			mockClient,
		)

		txHash := "0x" + strings.Repeat("f", 64)
		msgTokenID1 := reader.NewMessageTokenID(10, 0)
		msgTokenID2 := reader.NewMessageTokenID(10, 1)
		msgTokenID3 := reader.NewMessageTokenID(10, 2)

		// Create valid CCTP messages with all required fields
		msg1 := createValidCCTPv2Message("0x01", "0x01")
		msg2 := createValidCCTPv2Message("0x02", "0x02")
		msg2.DecodedMessage.DecodedMessageBody.Amount = "4000000"

		// Calculate depositHashes for the messages we have
		hash1, err := calculateDepositHash(msg1.DecodedMessage)
		require.NoError(t, err)
		hash2, err := calculateDepositHash(msg2.DecodedMessage)
		require.NoError(t, err)

		// Create v2Messages for all 3 MessageTokenIDs, but only 2 will match
		v2Messages := map[cciptypes.ChainSelector]map[reader.MessageTokenID]*SourceTokenDataPayloadV2{
			1: {
				msgTokenID1: &SourceTokenDataPayloadV2{
					SourceDomain: 0,
					DepositHash:  hash1,
				},
				msgTokenID2: &SourceTokenDataPayloadV2{
					SourceDomain: 0,
					DepositHash:  hash2,
				},
				msgTokenID3: &SourceTokenDataPayloadV2{
					SourceDomain: 0,
					DepositHash:  [32]byte{99}, // Won't match any message
				},
			},
		}

		// Mock API response with only 2 messages for 3 MessageTokenIDs
		mockClient.On("GetMessages", ctx, cciptypes.ChainSelector(1), uint32(0), txHash).
			Return(CCTPv2Messages{Messages: []CCTPv2Message{msg1, msg2}}, nil)

		txGroups := map[cciptypes.ChainSelector]map[TxKey][]reader.MessageTokenID{
			1: {{SourceDomain: 0, TxHash: txHash}: {msgTokenID1, msgTokenID2, msgTokenID3}},
		}

		result, err := observer.fetchCCTPv2Attestations(ctx, txGroups, v2Messages)

		require.NoError(t, err)
		require.Contains(t, result[1], msgTokenID1)
		require.Contains(t, result[1], msgTokenID2)
		// msgTokenID3 not in result - will be marked as ErrDataMissing
		assert.NotContains(t, result[1], msgTokenID3)
		mockClient.AssertExpectations(t)
	})

	t.Run("Edge case - more messages than MessageTokenIDs", func(t *testing.T) {
		mockClient := new(MockCCTPv2HTTPClient)
		observer := NewCCTPv2TokenDataObserver(
			lggr,
			1,
			map[cciptypes.ChainSelector]string{1: "0x1111"},
			successEncoder,
			mockClient,
		)

		txHash := "0x" + strings.Repeat("1", 64)
		msgTokenID1 := reader.NewMessageTokenID(10, 0)

		// Create 3 different CCTP messages
		msg1 := createValidCCTPv2Message("0x01", "0x01")
		msg2 := createValidCCTPv2Message("0x02", "0x02")
		msg2.DecodedMessage.DecodedMessageBody.Amount = "5000000"
		msg3 := createValidCCTPv2Message("0x03", "0x03")
		msg3.DecodedMessage.DecodedMessageBody.Amount = "6000000"

		// Calculate hash for msg1 only (this is what we expect)
		hash1, err := calculateDepositHash(msg1.DecodedMessage)
		require.NoError(t, err)

		// Create v2Messages with only msg1's hash
		v2Messages := map[cciptypes.ChainSelector]map[reader.MessageTokenID]*SourceTokenDataPayloadV2{
			1: {
				msgTokenID1: &SourceTokenDataPayloadV2{
					SourceDomain: 0,
					DepositHash:  hash1,
				},
			},
		}

		// Mock API response with 3 messages for 1 MessageTokenID
		mockClient.On("GetMessages", ctx, cciptypes.ChainSelector(1), uint32(0), txHash).
			Return(CCTPv2Messages{Messages: []CCTPv2Message{msg1, msg2, msg3}}, nil)

		txGroups := map[cciptypes.ChainSelector]map[TxKey][]reader.MessageTokenID{
			1: {{SourceDomain: 0, TxHash: txHash}: {msgTokenID1}},
		}

		result, err := observer.fetchCCTPv2Attestations(ctx, txGroups, v2Messages)

		require.NoError(t, err)
		require.Contains(t, result[1], msgTokenID1)
		// Only first message is assigned
		assert.Equal(t, msg1, result[1][msgTokenID1])
		mockClient.AssertExpectations(t)
	})

	t.Run("Edge case - empty txGroups", func(t *testing.T) {
		mockClient := new(MockCCTPv2HTTPClient)
		observer := NewCCTPv2TokenDataObserver(
			lggr,
			1,
			map[cciptypes.ChainSelector]string{1: "0x1111"},
			successEncoder,
			mockClient,
		)

		txGroups := map[cciptypes.ChainSelector]map[TxKey][]reader.MessageTokenID{}
		v2Messages := map[cciptypes.ChainSelector]map[reader.MessageTokenID]*SourceTokenDataPayloadV2{}

		result, err := observer.fetchCCTPv2Attestations(ctx, txGroups, v2Messages)

		require.NoError(t, err)
		assert.Empty(t, result)
		mockClient.AssertExpectations(t)
	})

	t.Run("Edge case - empty messages from API", func(t *testing.T) {
		mockClient := new(MockCCTPv2HTTPClient)
		observer := NewCCTPv2TokenDataObserver(
			lggr,
			1,
			map[cciptypes.ChainSelector]string{1: "0x1111"},
			successEncoder,
			mockClient,
		)

		txHash := "0x" + strings.Repeat("2", 64)
		msgTokenID := reader.NewMessageTokenID(10, 0)

		// Create v2Messages (won't be matched due to empty API response)
		v2Messages := map[cciptypes.ChainSelector]map[reader.MessageTokenID]*SourceTokenDataPayloadV2{
			1: {
				msgTokenID: &SourceTokenDataPayloadV2{
					SourceDomain: 0,
					DepositHash:  [32]byte{5, 6, 7},
				},
			},
		}

		// Mock API response with empty messages array
		mockClient.On("GetMessages", ctx, cciptypes.ChainSelector(1), uint32(0), txHash).
			Return(CCTPv2Messages{Messages: []CCTPv2Message{}}, nil)

		txGroups := map[cciptypes.ChainSelector]map[TxKey][]reader.MessageTokenID{
			1: {{SourceDomain: 0, TxHash: txHash}: {msgTokenID}},
		}

		result, err := observer.fetchCCTPv2Attestations(ctx, txGroups, v2Messages)

		require.NoError(t, err)
		require.Contains(t, result, cciptypes.ChainSelector(1))
		// No messages assigned - will be marked as ErrDataMissing
		assert.NotContains(t, result[1], msgTokenID)
		mockClient.AssertExpectations(t)
	})

	t.Run("Edge case - out-of-order messages from Circle API", func(t *testing.T) {
		mockClient := new(MockCCTPv2HTTPClient)
		observer := NewCCTPv2TokenDataObserver(
			lggr,
			1,
			map[cciptypes.ChainSelector]string{1: "0x1111"},
			successEncoder,
			mockClient,
		)

		txHash := "0x" + strings.Repeat("3", 64)
		msgTokenID1 := reader.NewMessageTokenID(10, 0)
		msgTokenID2 := reader.NewMessageTokenID(10, 1)
		msgTokenID3 := reader.NewMessageTokenID(10, 2)

		// Create 3 different CCTP messages with distinct amounts
		msg1 := createValidCCTPv2Message("0x01", "0x01")
		msg1.DecodedMessage.DecodedMessageBody.Amount = "1000000"
		msg2 := createValidCCTPv2Message("0x02", "0x02")
		msg2.DecodedMessage.DecodedMessageBody.Amount = "2000000"
		msg3 := createValidCCTPv2Message("0x03", "0x03")
		msg3.DecodedMessage.DecodedMessageBody.Amount = "3000000"

		// Calculate depositHashes for each message
		hash1, err := calculateDepositHash(msg1.DecodedMessage)
		require.NoError(t, err)
		hash2, err := calculateDepositHash(msg2.DecodedMessage)
		require.NoError(t, err)
		hash3, err := calculateDepositHash(msg3.DecodedMessage)
		require.NoError(t, err)

		// Create v2Messages in order: msgTokenID1 -> hash1, msgTokenID2 -> hash2, msgTokenID3 -> hash3
		v2Messages := map[cciptypes.ChainSelector]map[reader.MessageTokenID]*SourceTokenDataPayloadV2{
			1: {
				msgTokenID1: &SourceTokenDataPayloadV2{
					SourceDomain: 0,
					DepositHash:  hash1,
				},
				msgTokenID2: &SourceTokenDataPayloadV2{
					SourceDomain: 0,
					DepositHash:  hash2,
				},
				msgTokenID3: &SourceTokenDataPayloadV2{
					SourceDomain: 0,
					DepositHash:  hash3,
				},
			},
		}

		// Circle API returns messages in REVERSE order: msg3, msg2, msg1
		// This tests that depositHash matching works regardless of order
		mockClient.On("GetMessages", ctx, cciptypes.ChainSelector(1), uint32(0), txHash).
			Return(CCTPv2Messages{Messages: []CCTPv2Message{msg3, msg2, msg1}}, nil)

		txGroups := map[cciptypes.ChainSelector]map[TxKey][]reader.MessageTokenID{
			1: {{SourceDomain: 0, TxHash: txHash}: {msgTokenID1, msgTokenID2, msgTokenID3}},
		}

		result, err := observer.fetchCCTPv2Attestations(ctx, txGroups, v2Messages)

		require.NoError(t, err)
		require.Contains(t, result[1], msgTokenID1)
		require.Contains(t, result[1], msgTokenID2)
		require.Contains(t, result[1], msgTokenID3)
		// Verify correct assignments despite reverse order
		assert.Equal(t, msg1, result[1][msgTokenID1])
		assert.Equal(t, msg2, result[1][msgTokenID2])
		assert.Equal(t, msg3, result[1][msgTokenID3])
		mockClient.AssertExpectations(t)
	})
}
