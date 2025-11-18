package v2

import (
	"context"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"strings"
	"testing"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
	"github.com/stretchr/testify/assert"

	"github.com/smartcontractkit/chainlink-ccip/execute/exectypes"
	"github.com/smartcontractkit/chainlink-ccip/execute/tokendata"
)

// Test helper functions

// defaultTestEncoder returns a simple encoder that concatenates message and attestation for testing
func defaultTestEncoder() AttestationEncoder {
	return func(ctx context.Context, msg cciptypes.Bytes, att cciptypes.Bytes) (cciptypes.Bytes, error) {
		result := make([]byte, len(msg)+len(att))
		copy(result, msg)
		copy(result[len(msg):], att)
		return result, nil
	}
}

// newTestCCTPv2Observer creates a test observer with sensible defaults.
// Pass nil for httpClient to use default (no client).
// Pass 0 for chainSelector to use default (999).
// Note: This creates the observer directly using a struct literal, bypassing the constructor.
// This is appropriate for unit tests where we want to inject mock dependencies.
// Tests should use the setup callback to populate supportedPoolsBySelector as needed.
func newTestCCTPv2Observer(
	t *testing.T,
	httpClient CCTPv2HTTPClient,
	chainSelector cciptypes.ChainSelector,
) *CCTPv2TokenDataObserver {
	// Apply defaults
	if chainSelector == 0 {
		chainSelector = cciptypes.ChainSelector(999)
	}

	// Create observer directly using struct literal (bypassing constructor for test flexibility)
	observer := &CCTPv2TokenDataObserver{
		lggr:                     logger.Test(t),
		destChainSelector:        chainSelector,
		supportedPoolsBySelector: make(map[cciptypes.ChainSelector]string), // Tests populate this via setup callback
		attestationEncoder:       defaultTestEncoder(),
		httpClient:               httpClient,
		calculateDepositHashFn:   CalculateDepositHash,
		messageToTokenDataFn:     CCTPv2MessageToTokenData,
		metricsReporter:          NewNoOpMetricsReporter(),
	}

	return observer
}

func TestCCTPv2TokenDataObserver_GetCCTPv2RequestParams(t *testing.T) {
	const (
		testChain1   = cciptypes.ChainSelector(1)
		testChain2   = cciptypes.ChainSelector(2)
		testPoolAddr = "0x1234567890123456789012345678901234567890"
		testTxHash1  = "0xabcdef0123456789abcdef0123456789abcdef0123456789abcdef0123456789"
		testTxHash2  = "0xfedcba9876543210fedcba9876543210fedcba9876543210fedcba9876543210"
		depositHex1  = "1111111111111111111111111111111111111111111111111111111111111111"
		depositHex2  = "2222222222222222222222222222222222222222222222222222222222222222"
	)

	tests := []struct {
		name     string
		messages exectypes.MessageObservations
		setup    func(*CCTPv2TokenDataObserver)
		validate func(*testing.T, CCTPv2TokenDataObserver, exectypes.MessageObservations)
	}{
		{
			name:     "empty messages returns empty set",
			messages: exectypes.MessageObservations{},
			validate: func(t *testing.T, observer CCTPv2TokenDataObserver, _ exectypes.MessageObservations) {
				result := observer.getCCTPv2RequestParams(exectypes.MessageObservations{})
				assert.Equal(t, 0, result.Cardinality())
			},
		},
		{
			name: "single supported token returns correct params",
			messages: exectypes.MessageObservations{
				testChain1: {
					1: createTestMessage(testTxHash1, []cciptypes.RampTokenAmount{
						createCCTPv2Token(testPoolAddr, 100, depositHex1),
					}),
				},
			},
			setup: func(observer *CCTPv2TokenDataObserver) {
				observer.supportedPoolsBySelector = map[cciptypes.ChainSelector]string{
					testChain1: testPoolAddr,
				}
			},
			validate: func(t *testing.T, observer CCTPv2TokenDataObserver, msgs exectypes.MessageObservations) {
				result := observer.getCCTPv2RequestParams(msgs)
				assert.Equal(t, 1, result.Cardinality())

				expectedParam := CCTPv2RequestParams{
					chainSelector: testChain1,
					sourceDomain:  100,
					txHash:        testTxHash1,
				}
				assert.True(t, result.Contains(expectedParam))
			},
		},
		{
			name: "multiple supported tokens in same message with same params returns one param",
			messages: exectypes.MessageObservations{
				testChain1: {
					1: createTestMessage(testTxHash1, []cciptypes.RampTokenAmount{
						createCCTPv2Token(testPoolAddr, 100, depositHex1),
						createCCTPv2Token(testPoolAddr, 100, depositHex2),
					}),
				},
			},
			setup: func(observer *CCTPv2TokenDataObserver) {
				observer.supportedPoolsBySelector = map[cciptypes.ChainSelector]string{
					testChain1: testPoolAddr,
				}
			},
			validate: func(t *testing.T, observer CCTPv2TokenDataObserver, msgs exectypes.MessageObservations) {
				result := observer.getCCTPv2RequestParams(msgs)
				// Both tokens have same (chainSelector, sourceDomain, txHash) → deduped to 1
				assert.Equal(t, 1, result.Cardinality())

				expectedParam := CCTPv2RequestParams{
					chainSelector: testChain1,
					sourceDomain:  100,
					txHash:        testTxHash1,
				}
				assert.True(t, result.Contains(expectedParam))
			},
		},
		{
			name: "multiple messages with different txHashes returns multiple params",
			messages: exectypes.MessageObservations{
				testChain1: {
					1: createTestMessage(testTxHash1, []cciptypes.RampTokenAmount{
						createCCTPv2Token(testPoolAddr, 100, depositHex1),
					}),
					2: createTestMessage(testTxHash2, []cciptypes.RampTokenAmount{
						createCCTPv2Token(testPoolAddr, 100, depositHex2),
					}),
				},
			},
			setup: func(observer *CCTPv2TokenDataObserver) {
				observer.supportedPoolsBySelector = map[cciptypes.ChainSelector]string{
					testChain1: testPoolAddr,
				}
			},
			validate: func(t *testing.T, observer CCTPv2TokenDataObserver, msgs exectypes.MessageObservations) {
				result := observer.getCCTPv2RequestParams(msgs)
				assert.Equal(t, 2, result.Cardinality())

				assert.True(t, result.Contains(CCTPv2RequestParams{
					chainSelector: testChain1,
					sourceDomain:  100,
					txHash:        testTxHash1,
				}))
				assert.True(t, result.Contains(CCTPv2RequestParams{
					chainSelector: testChain1,
					sourceDomain:  100,
					txHash:        testTxHash2,
				}))
			},
		},
		{
			name: "unsupported token with wrong pool address is filtered out",
			messages: exectypes.MessageObservations{
				testChain1: {
					1: createTestMessage(testTxHash1, []cciptypes.RampTokenAmount{
						createCCTPv2Token(
							"0x9999999999999999999999999999999999999999",
							100,
							depositHex1,
						),
					}),
				},
			},
			setup: func(observer *CCTPv2TokenDataObserver) {
				observer.supportedPoolsBySelector = map[cciptypes.ChainSelector]string{
					testChain1: testPoolAddr,
				}
			},
			validate: func(t *testing.T, observer CCTPv2TokenDataObserver, msgs exectypes.MessageObservations) {
				result := observer.getCCTPv2RequestParams(msgs)
				assert.Equal(t, 0, result.Cardinality())
			},
		},
		{
			name: "token with invalid ExtraData is filtered out",
			messages: exectypes.MessageObservations{
				testChain1: {
					1: createTestMessage(testTxHash1, []cciptypes.RampTokenAmount{
						{
							SourcePoolAddress: mustDecodeAddress(testPoolAddr),
							ExtraData:         cciptypes.Bytes{0x01, 0x02, 0x03}, // Invalid length
						},
					}),
				},
			},
			setup: func(observer *CCTPv2TokenDataObserver) {
				observer.supportedPoolsBySelector = map[cciptypes.ChainSelector]string{
					testChain1: testPoolAddr,
				}
			},
			validate: func(t *testing.T, observer CCTPv2TokenDataObserver, msgs exectypes.MessageObservations) {
				result := observer.getCCTPv2RequestParams(msgs)
				assert.Equal(t, 0, result.Cardinality())
			},
		},
		{
			name: "message without txHash is skipped",
			messages: exectypes.MessageObservations{
				testChain1: {
					1: createTestMessage("", []cciptypes.RampTokenAmount{
						createCCTPv2Token(testPoolAddr, 100, depositHex1),
					}),
				},
			},
			setup: func(observer *CCTPv2TokenDataObserver) {
				observer.supportedPoolsBySelector = map[cciptypes.ChainSelector]string{
					testChain1: testPoolAddr,
				}
			},
			validate: func(t *testing.T, observer CCTPv2TokenDataObserver, msgs exectypes.MessageObservations) {
				result := observer.getCCTPv2RequestParams(msgs)
				assert.Equal(t, 0, result.Cardinality())
			},
		},
		{
			name: "multiple chains returns params for each chain",
			messages: exectypes.MessageObservations{
				testChain1: {
					1: createTestMessage(testTxHash1, []cciptypes.RampTokenAmount{
						createCCTPv2Token(testPoolAddr, 100, depositHex1),
					}),
				},
				testChain2: {
					1: createTestMessage(testTxHash2, []cciptypes.RampTokenAmount{
						createCCTPv2Token(testPoolAddr, 200, depositHex2),
					}),
				},
			},
			setup: func(observer *CCTPv2TokenDataObserver) {
				observer.supportedPoolsBySelector = map[cciptypes.ChainSelector]string{
					testChain1: testPoolAddr,
					testChain2: testPoolAddr,
				}
			},
			validate: func(t *testing.T, observer CCTPv2TokenDataObserver, msgs exectypes.MessageObservations) {
				result := observer.getCCTPv2RequestParams(msgs)
				assert.Equal(t, 2, result.Cardinality())

				assert.True(t, result.Contains(CCTPv2RequestParams{
					chainSelector: testChain1,
					sourceDomain:  100,
					txHash:        testTxHash1,
				}))
				assert.True(t, result.Contains(CCTPv2RequestParams{
					chainSelector: testChain2,
					sourceDomain:  200,
					txHash:        testTxHash2,
				}))
			},
		},
		{
			name: "deduplication works across multiple messages",
			messages: exectypes.MessageObservations{
				testChain1: {
					1: createTestMessage(testTxHash1, []cciptypes.RampTokenAmount{
						createCCTPv2Token(testPoolAddr, 100, depositHex1),
					}),
					2: createTestMessage(testTxHash1, []cciptypes.RampTokenAmount{
						createCCTPv2Token(testPoolAddr, 100, depositHex2),
					}),
				},
			},
			setup: func(observer *CCTPv2TokenDataObserver) {
				observer.supportedPoolsBySelector = map[cciptypes.ChainSelector]string{
					testChain1: testPoolAddr,
				}
			},
			validate: func(t *testing.T, observer CCTPv2TokenDataObserver, msgs exectypes.MessageObservations) {
				result := observer.getCCTPv2RequestParams(msgs)
				// Same txHash and sourceDomain → deduped to 1
				assert.Equal(t, 1, result.Cardinality())
			},
		},
		{
			name: "different source domains create different params",
			messages: exectypes.MessageObservations{
				testChain1: {
					1: createTestMessage(testTxHash1, []cciptypes.RampTokenAmount{
						createCCTPv2Token(testPoolAddr, 100, depositHex1),
						createCCTPv2Token(testPoolAddr, 200, depositHex2),
					}),
				},
			},
			setup: func(observer *CCTPv2TokenDataObserver) {
				observer.supportedPoolsBySelector = map[cciptypes.ChainSelector]string{
					testChain1: testPoolAddr,
				}
			},
			validate: func(t *testing.T, observer CCTPv2TokenDataObserver, msgs exectypes.MessageObservations) {
				result := observer.getCCTPv2RequestParams(msgs)
				// Different sourceDomains → 2 different params
				assert.Equal(t, 2, result.Cardinality())
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			observer := newTestCCTPv2Observer(t, nil, 0)

			if tt.setup != nil {
				tt.setup(observer)
			}

			tt.validate(t, *observer, tt.messages)
		})
	}
}

// createTestMessage creates a test message with the given txHash and tokens
func createTestMessage(txHash string, tokens []cciptypes.RampTokenAmount) cciptypes.Message {
	return cciptypes.Message{
		Header: cciptypes.RampMessageHeader{
			TxHash: txHash,
		},
		TokenAmounts: tokens,
	}
}

// createCCTPv2Token creates a valid CCTPv2 token with the given pool address and source domain
func createCCTPv2Token(poolAddr string, sourceDomain uint32, depositHashHex string) cciptypes.RampTokenAmount {
	// Create valid CCTPv2 ExtraData: versionTag (4) + sourceDomain (4) + depositHash (32)
	extraData := make([]byte, 40)

	// Version tag (CCTPVersion2Tag)
	binary.BigEndian.PutUint32(extraData[0:4], CCTPVersion2Tag)

	// Source domain
	binary.BigEndian.PutUint32(extraData[4:8], sourceDomain)

	// Deposit hash (use a simple hash based on the hex string)
	depositHash := mustHexToBytes32(depositHashHex)
	copy(extraData[8:40], depositHash[:])

	return cciptypes.RampTokenAmount{
		SourcePoolAddress: mustDecodeAddress(poolAddr),
		ExtraData:         cciptypes.Bytes(extraData),
	}
}

// mustDecodeAddress decodes a hex address string to UnknownAddress
func mustDecodeAddress(hexAddr string) cciptypes.UnknownAddress {
	hexAddr = strings.TrimPrefix(hexAddr, "0x")
	decoded, err := hex.DecodeString(hexAddr)
	if err != nil {
		panic(err)
	}
	return cciptypes.UnknownAddress(decoded)
}

func TestCCTPv2TokenDataObserver_MakeCCTPv2Requests(t *testing.T) {
	const (
		testChain1  = cciptypes.ChainSelector(1)
		testChain2  = cciptypes.ChainSelector(2)
		testTxHash1 = "0xabcdef0123456789abcdef0123456789abcdef0123456789abcdef0123456789"
		testTxHash2 = "0xfedcba9876543210fedcba9876543210fedcba9876543210fedcba9876543210"
	)

	tests := []struct {
		name     string
		params   []CCTPv2RequestParams
		setup    func(*mockCCTPv2HTTPClient)
		validate func(*testing.T, map[CCTPv2RequestParams]CCTPv2Messages, *mockCCTPv2HTTPClient)
	}{
		{
			name:   "empty request params returns empty map",
			params: []CCTPv2RequestParams{},
			validate: func(t *testing.T, result map[CCTPv2RequestParams]CCTPv2Messages, mock *mockCCTPv2HTTPClient) {
				assert.Equal(t, 0, len(result))
				assert.Equal(t, 0, len(mock.calls))
			},
		},
		{
			name: "single successful request returns correct mapping",
			params: []CCTPv2RequestParams{
				{chainSelector: testChain1, sourceDomain: 100, txHash: testTxHash1},
			},
			setup: func(mock *mockCCTPv2HTTPClient) {
				params := CCTPv2RequestParams{chainSelector: testChain1, sourceDomain: 100, txHash: testTxHash1}
				mock.responses[params] = CCTPv2Messages{
					Messages: []CCTPv2Message{{Status: "complete"}},
				}
			},
			validate: func(t *testing.T, result map[CCTPv2RequestParams]CCTPv2Messages, mock *mockCCTPv2HTTPClient) {
				assert.Equal(t, 1, len(result))
				params := CCTPv2RequestParams{chainSelector: testChain1, sourceDomain: 100, txHash: testTxHash1}
				assert.Contains(t, result, params)
				assert.Equal(t, "complete", result[params].Messages[0].Status)
				assert.Equal(t, 1, len(mock.calls))
			},
		},
		{
			name: "multiple successful requests return multiple mappings",
			params: []CCTPv2RequestParams{
				{chainSelector: testChain1, sourceDomain: 100, txHash: testTxHash1},
				{chainSelector: testChain2, sourceDomain: 200, txHash: testTxHash2},
			},
			setup: func(mock *mockCCTPv2HTTPClient) {
				params1 := CCTPv2RequestParams{chainSelector: testChain1, sourceDomain: 100, txHash: testTxHash1}
				params2 := CCTPv2RequestParams{chainSelector: testChain2, sourceDomain: 200, txHash: testTxHash2}
				mock.responses[params1] = CCTPv2Messages{Messages: []CCTPv2Message{{Status: "complete1"}}}
				mock.responses[params2] = CCTPv2Messages{Messages: []CCTPv2Message{{Status: "complete2"}}}
			},
			validate: func(t *testing.T, result map[CCTPv2RequestParams]CCTPv2Messages, mock *mockCCTPv2HTTPClient) {
				assert.Equal(t, 2, len(result))
				assert.Equal(t, 2, len(mock.calls))
			},
		},
		{
			name: "request fails with error - logs and continues",
			params: []CCTPv2RequestParams{
				{chainSelector: testChain1, sourceDomain: 100, txHash: testTxHash1},
				{chainSelector: testChain2, sourceDomain: 200, txHash: testTxHash2},
			},
			setup: func(mock *mockCCTPv2HTTPClient) {
				params1 := CCTPv2RequestParams{chainSelector: testChain1, sourceDomain: 100, txHash: testTxHash1}
				params2 := CCTPv2RequestParams{chainSelector: testChain2, sourceDomain: 200, txHash: testTxHash2}
				mock.errors[params1] = assert.AnError
				mock.responses[params2] = CCTPv2Messages{Messages: []CCTPv2Message{{Status: "complete"}}}
			},
			validate: func(t *testing.T, result map[CCTPv2RequestParams]CCTPv2Messages, mock *mockCCTPv2HTTPClient) {
				assert.Equal(t, 1, len(result))
				params2 := CCTPv2RequestParams{chainSelector: testChain2, sourceDomain: 200, txHash: testTxHash2}
				assert.Contains(t, result, params2)
				assert.Equal(t, 2, len(mock.calls))
			},
		},
		{
			name: "all requests fail returns empty map",
			params: []CCTPv2RequestParams{
				{chainSelector: testChain1, sourceDomain: 100, txHash: testTxHash1},
				{chainSelector: testChain2, sourceDomain: 200, txHash: testTxHash2},
			},
			setup: func(mock *mockCCTPv2HTTPClient) {
				params1 := CCTPv2RequestParams{chainSelector: testChain1, sourceDomain: 100, txHash: testTxHash1}
				params2 := CCTPv2RequestParams{chainSelector: testChain2, sourceDomain: 200, txHash: testTxHash2}
				mock.errors[params1] = assert.AnError
				mock.errors[params2] = assert.AnError
			},
			validate: func(t *testing.T, result map[CCTPv2RequestParams]CCTPv2Messages, mock *mockCCTPv2HTTPClient) {
				assert.Equal(t, 0, len(result))
				assert.Equal(t, 2, len(mock.calls))
			},
		},
		{
			name: "preserves request param mapping correctly",
			params: []CCTPv2RequestParams{
				{chainSelector: testChain1, sourceDomain: 100, txHash: testTxHash1},
				{chainSelector: testChain1, sourceDomain: 200, txHash: testTxHash1},
			},
			setup: func(mock *mockCCTPv2HTTPClient) {
				params1 := CCTPv2RequestParams{chainSelector: testChain1, sourceDomain: 100, txHash: testTxHash1}
				params2 := CCTPv2RequestParams{chainSelector: testChain1, sourceDomain: 200, txHash: testTxHash1}
				mock.responses[params1] = CCTPv2Messages{Messages: []CCTPv2Message{{EventNonce: "1"}}}
				mock.responses[params2] = CCTPv2Messages{Messages: []CCTPv2Message{{EventNonce: "2"}}}
			},
			validate: func(t *testing.T, result map[CCTPv2RequestParams]CCTPv2Messages, mock *mockCCTPv2HTTPClient) {
				assert.Equal(t, 2, len(result))
				params1 := CCTPv2RequestParams{chainSelector: testChain1, sourceDomain: 100, txHash: testTxHash1}
				params2 := CCTPv2RequestParams{chainSelector: testChain1, sourceDomain: 200, txHash: testTxHash1}
				assert.Equal(t, "1", result[params1].Messages[0].EventNonce)
				assert.Equal(t, "2", result[params2].Messages[0].EventNonce)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := newMockCCTPv2HTTPClient()
			if tt.setup != nil {
				tt.setup(mock)
			}

			observer := newTestCCTPv2Observer(t, mock, 0)

			// Build set from params
			paramsSet := mapset.NewSet[CCTPv2RequestParams]()
			for _, p := range tt.params {
				paramsSet.Add(p)
			}

			result := observer.makeCCTPv2Requests(context.Background(), paramsSet)

			tt.validate(t, result, mock)
		})
	}
}

// mockCCTPv2HTTPClient is a mock implementation of CCTPv2HTTPClient for testing
type mockCCTPv2HTTPClient struct {
	responses map[CCTPv2RequestParams]CCTPv2Messages
	errors    map[CCTPv2RequestParams]error
	calls     []CCTPv2RequestParams
}

func newMockCCTPv2HTTPClient() *mockCCTPv2HTTPClient {
	return &mockCCTPv2HTTPClient{
		responses: make(map[CCTPv2RequestParams]CCTPv2Messages),
		errors:    make(map[CCTPv2RequestParams]error),
		calls:     []CCTPv2RequestParams{},
	}
}

func (m *mockCCTPv2HTTPClient) GetMessages(
	ctx context.Context,
	sourceChain cciptypes.ChainSelector,
	sourceDomainID uint32,
	transactionHash string,
) (CCTPv2Messages, error) {
	params := CCTPv2RequestParams{
		chainSelector: sourceChain,
		sourceDomain:  sourceDomainID,
		txHash:        transactionHash,
	}
	m.calls = append(m.calls, params)

	if err, hasError := m.errors[params]; hasError {
		return CCTPv2Messages{}, err
	}

	if response, hasResponse := m.responses[params]; hasResponse {
		return response, nil
	}

	return CCTPv2Messages{}, nil
}

func (m *mockCCTPv2HTTPClient) addResponse(
	chain cciptypes.ChainSelector,
	domain uint32,
	txHash string,
	messages CCTPv2Messages,
) {
	params := CCTPv2RequestParams{
		chainSelector: chain,
		sourceDomain:  domain,
		txHash:        txHash,
	}
	m.responses[params] = messages
}

func (m *mockCCTPv2HTTPClient) addError(
	chain cciptypes.ChainSelector,
	domain uint32,
	txHash string,
	err error,
) {
	params := CCTPv2RequestParams{
		chainSelector: chain,
		sourceDomain:  domain,
		txHash:        txHash,
	}
	m.errors[params] = err
}

func (m *mockCCTPv2HTTPClient) getCallCount() int {
	return len(m.calls)
}

func (m *mockCCTPv2HTTPClient) wasCalledWith(
	chain cciptypes.ChainSelector,
	domain uint32,
	txHash string,
) bool {
	params := CCTPv2RequestParams{
		chainSelector: chain,
		sourceDomain:  domain,
		txHash:        txHash,
	}
	for _, call := range m.calls {
		if call == params {
			return true
		}
	}
	return false
}

func TestCCTPv2MessageToTokenData(t *testing.T) {
	validMessage := "1234567890abcdef"
	validAttestation := "fedcba0987654321"
	validMessageWithPrefix := "0x" + validMessage
	validAttestationWithPrefix := "0x" + validAttestation

	tests := []struct {
		name           string
		msg            CCTPv2Message
		encoderSetup   func() AttestationEncoder
		expectedReady  bool
		expectedError  error
		validateResult func(*testing.T, exectypes.TokenData)
	}{
		{
			name: "happy path - complete status with valid hex returns success",
			msg: CCTPv2Message{
				Status:      "complete",
				Message:     validMessage,
				Attestation: validAttestation,
			},
			encoderSetup: func() AttestationEncoder {
				return func(ctx context.Context, msg, att cciptypes.Bytes) (cciptypes.Bytes, error) {
					// Verify encoder receives decoded bytes
					expectedMsg, _ := hex.DecodeString(validMessage)
					expectedAtt, _ := hex.DecodeString(validAttestation)
					assert.Equal(t, expectedMsg, []byte(msg))
					assert.Equal(t, expectedAtt, []byte(att))
					return cciptypes.Bytes("encoded-data"), nil
				}
			},
			expectedReady: true,
			expectedError: nil,
			validateResult: func(t *testing.T, result exectypes.TokenData) {
				assert.True(t, result.Ready)
				assert.Equal(t, cciptypes.Bytes("encoded-data"), result.Data)
				assert.NoError(t, result.Error)
			},
		},
		{
			name: "happy path - message and attestation with 0x prefix",
			msg: CCTPv2Message{
				Status:      "complete",
				Message:     validMessageWithPrefix,
				Attestation: validAttestationWithPrefix,
			},
			encoderSetup: func() AttestationEncoder {
				return func(ctx context.Context, msg, att cciptypes.Bytes) (cciptypes.Bytes, error) {
					return cciptypes.Bytes("encoded-with-prefix"), nil
				}
			},
			expectedReady: true,
			expectedError: nil,
			validateResult: func(t *testing.T, result exectypes.TokenData) {
				assert.True(t, result.Ready)
				assert.Equal(t, cciptypes.Bytes("encoded-with-prefix"), result.Data)
			},
		},
		{
			name: "status not complete - pending",
			msg: CCTPv2Message{
				Status:      "pending",
				Message:     validMessage,
				Attestation: validAttestation,
			},
			encoderSetup: func() AttestationEncoder {
				return func(ctx context.Context, msg, att cciptypes.Bytes) (cciptypes.Bytes, error) {
					t.Fatal("encoder should not be called when status is not complete")
					return nil, nil
				}
			},
			expectedReady: false,
			expectedError: tokendata.ErrNotReady,
			validateResult: func(t *testing.T, result exectypes.TokenData) {
				assert.False(t, result.Ready)
				assert.Equal(t, tokendata.ErrNotReady, result.Error)
			},
		},
		{
			name: "status empty string",
			msg: CCTPv2Message{
				Status:      "",
				Message:     validMessage,
				Attestation: validAttestation,
			},
			encoderSetup: func() AttestationEncoder {
				return func(ctx context.Context, msg, att cciptypes.Bytes) (cciptypes.Bytes, error) {
					t.Fatal("encoder should not be called when status is empty")
					return nil, nil
				}
			},
			expectedReady: false,
			expectedError: tokendata.ErrNotReady,
			validateResult: func(t *testing.T, result exectypes.TokenData) {
				assert.False(t, result.Ready)
				assert.Equal(t, tokendata.ErrNotReady, result.Error)
			},
		},
		{
			name: "invalid message hex - non-hex characters",
			msg: CCTPv2Message{
				Status:      "complete",
				Message:     "0xINVALIDHEX",
				Attestation: validAttestation,
			},
			encoderSetup: func() AttestationEncoder {
				return func(ctx context.Context, msg, att cciptypes.Bytes) (cciptypes.Bytes, error) {
					t.Fatal("encoder should not be called when message is invalid")
					return nil, nil
				}
			},
			expectedReady: false,
			validateResult: func(t *testing.T, result exectypes.TokenData) {
				assert.False(t, result.Ready)
				assert.Error(t, result.Error)
				assert.Contains(t, result.Error.Error(), "invalid byte")
			},
		},
		{
			name: "invalid message hex - odd length",
			msg: CCTPv2Message{
				Status:      "complete",
				Message:     "0x123",
				Attestation: validAttestation,
			},
			encoderSetup: func() AttestationEncoder {
				return func(ctx context.Context, msg, att cciptypes.Bytes) (cciptypes.Bytes, error) {
					t.Fatal("encoder should not be called when message is invalid")
					return nil, nil
				}
			},
			expectedReady: false,
			validateResult: func(t *testing.T, result exectypes.TokenData) {
				assert.False(t, result.Ready)
				assert.Error(t, result.Error)
			},
		},
		{
			name: "empty message string",
			msg: CCTPv2Message{
				Status:      "complete",
				Message:     "",
				Attestation: validAttestation,
			},
			encoderSetup: func() AttestationEncoder {
				return func(ctx context.Context, msg, att cciptypes.Bytes) (cciptypes.Bytes, error) {
					t.Fatal("encoder should not be called when message is empty")
					return nil, nil
				}
			},
			expectedReady: false,
			expectedError: tokendata.ErrDataMissing,
			validateResult: func(t *testing.T, result exectypes.TokenData) {
				assert.False(t, result.Ready)
				assert.ErrorIs(t, result.Error, tokendata.ErrDataMissing)
			},
		},
		{
			name: "message is just 0x prefix",
			msg: CCTPv2Message{
				Status:      "complete",
				Message:     "0x",
				Attestation: validAttestation,
			},
			encoderSetup: func() AttestationEncoder {
				return func(ctx context.Context, msg, att cciptypes.Bytes) (cciptypes.Bytes, error) {
					t.Fatal("encoder should not be called when message is empty")
					return nil, nil
				}
			},
			expectedReady: false,
			expectedError: tokendata.ErrDataMissing,
			validateResult: func(t *testing.T, result exectypes.TokenData) {
				assert.False(t, result.Ready)
				assert.ErrorIs(t, result.Error, tokendata.ErrDataMissing)
			},
		},
		{
			name: "invalid attestation hex - non-hex characters",
			msg: CCTPv2Message{
				Status:      "complete",
				Message:     validMessage,
				Attestation: "0xGHIJKL",
			},
			encoderSetup: func() AttestationEncoder {
				return func(ctx context.Context, msg, att cciptypes.Bytes) (cciptypes.Bytes, error) {
					t.Fatal("encoder should not be called when attestation is invalid")
					return nil, nil
				}
			},
			expectedReady: false,
			validateResult: func(t *testing.T, result exectypes.TokenData) {
				assert.False(t, result.Ready)
				assert.Error(t, result.Error)
				assert.Contains(t, result.Error.Error(), "invalid byte")
			},
		},
		{
			name: "invalid attestation hex - odd length",
			msg: CCTPv2Message{
				Status:      "complete",
				Message:     validMessage,
				Attestation: "abc",
			},
			encoderSetup: func() AttestationEncoder {
				return func(ctx context.Context, msg, att cciptypes.Bytes) (cciptypes.Bytes, error) {
					t.Fatal("encoder should not be called when attestation is invalid")
					return nil, nil
				}
			},
			expectedReady: false,
			validateResult: func(t *testing.T, result exectypes.TokenData) {
				assert.False(t, result.Ready)
				assert.Error(t, result.Error)
			},
		},
		{
			name: "empty attestation string",
			msg: CCTPv2Message{
				Status:      "complete",
				Message:     validMessage,
				Attestation: "",
			},
			encoderSetup: func() AttestationEncoder {
				return func(ctx context.Context, msg, att cciptypes.Bytes) (cciptypes.Bytes, error) {
					t.Fatal("encoder should not be called when attestation is empty")
					return nil, nil
				}
			},
			expectedReady: false,
			expectedError: tokendata.ErrDataMissing,
			validateResult: func(t *testing.T, result exectypes.TokenData) {
				assert.False(t, result.Ready)
				assert.ErrorIs(t, result.Error, tokendata.ErrDataMissing)
			},
		},
		{
			name: "encoder returns error",
			msg: CCTPv2Message{
				Status:      "complete",
				Message:     validMessage,
				Attestation: validAttestation,
			},
			encoderSetup: func() AttestationEncoder {
				return func(ctx context.Context, msg, att cciptypes.Bytes) (cciptypes.Bytes, error) {
					return nil, assert.AnError
				}
			},
			expectedReady: false,
			validateResult: func(t *testing.T, result exectypes.TokenData) {
				assert.False(t, result.Ready)
				assert.Error(t, result.Error)
				assert.Contains(t, result.Error.Error(), "unable to encode attestation")
			},
		},
		{
			name: "encoder returns wrapped error",
			msg: CCTPv2Message{
				Status:      "complete",
				Message:     validMessage,
				Attestation: validAttestation,
			},
			encoderSetup: func() AttestationEncoder {
				return func(ctx context.Context, msg, att cciptypes.Bytes) (cciptypes.Bytes, error) {
					return nil, fmt.Errorf("custom encoding error: %w", assert.AnError)
				}
			},
			expectedReady: false,
			validateResult: func(t *testing.T, result exectypes.TokenData) {
				assert.False(t, result.Ready)
				assert.Error(t, result.Error)
				assert.Contains(t, result.Error.Error(), "unable to encode attestation")
				assert.Contains(t, result.Error.Error(), "custom encoding error")
			},
		},
		{
			name: "mixed case hex strings are valid",
			msg: CCTPv2Message{
				Status:      "complete",
				Message:     "0xAbCdEf",
				Attestation: "0xFeDcBa",
			},
			encoderSetup: func() AttestationEncoder {
				return func(ctx context.Context, msg, att cciptypes.Bytes) (cciptypes.Bytes, error) {
					return cciptypes.Bytes("encoded-mixed-case"), nil
				}
			},
			expectedReady: true,
			validateResult: func(t *testing.T, result exectypes.TokenData) {
				assert.True(t, result.Ready)
				assert.Equal(t, cciptypes.Bytes("encoded-mixed-case"), result.Data)
			},
		},
		{
			name: "very long hex strings work correctly",
			msg: CCTPv2Message{
				Status:      "complete",
				Message:     "0x" + strings.Repeat("ab", 500),
				Attestation: "0x" + strings.Repeat("cd", 500),
			},
			encoderSetup: func() AttestationEncoder {
				return func(ctx context.Context, msg, att cciptypes.Bytes) (cciptypes.Bytes, error) {
					assert.Equal(t, 500, len(msg))
					assert.Equal(t, 500, len(att))
					return cciptypes.Bytes("encoded-long"), nil
				}
			},
			expectedReady: true,
			validateResult: func(t *testing.T, result exectypes.TokenData) {
				assert.True(t, result.Ready)
			},
		},
		{
			name: "context is passed to encoder",
			msg: CCTPv2Message{
				Status:      "complete",
				Message:     validMessage,
				Attestation: validAttestation,
			},
			encoderSetup: func() AttestationEncoder {
				return func(ctx context.Context, msg, att cciptypes.Bytes) (cciptypes.Bytes, error) {
					assert.NotNil(t, ctx)
					return cciptypes.Bytes("encoded-with-context"), nil
				}
			},
			expectedReady: true,
			validateResult: func(t *testing.T, result exectypes.TokenData) {
				assert.True(t, result.Ready)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			encoder := tt.encoderSetup()
			result := CCTPv2MessageToTokenData(context.Background(), logger.Test(t), tt.msg, encoder)

			assert.Equal(t, tt.expectedReady, result.Ready)
			if tt.expectedError != nil {
				assert.ErrorIs(t, result.Error, tt.expectedError)
			}
			if tt.validateResult != nil {
				tt.validateResult(t, result)
			}
		})
	}
}

func TestCCTPv2TokenDataObserver_ConvertCCTPv2MessagesToTokenData(t *testing.T) {
	const (
		testChain1  = cciptypes.ChainSelector(1)
		testChain2  = cciptypes.ChainSelector(2)
		testTxHash1 = "0xabcdef0123456789abcdef0123456789abcdef0123456789abcdef0123456789"
		testTxHash2 = "0xfedcba9876543210fedcba9876543210fedcba9876543210fedcba9876543210"
	)

	// Shared test data
	validDecodedMessage := CCTPv2DecodedMessage{
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
	}

	validDecodedMessage2 := CCTPv2DecodedMessage{
		SourceDomain:         "1",
		DestinationDomain:    "3",
		MinFinalityThreshold: "3",
		DestinationCaller:    "0x1234567890123456789012345678901234567890",
		DecodedMessageBody: CCTPv2DecodedMessageBody{
			Amount:        "2000000",
			MintRecipient: "0xabcdef0123456789abcdef0123456789abcdef01",
			BurnToken:     "0xfedcba9876543210fedcba9876543210fedcba98",
			MaxFee:        "5000",
		},
	}

	hash1 := mustHexToBytes32("1111111111111111111111111111111111111111111111111111111111111111")
	hash2 := mustHexToBytes32("2222222222222222222222222222222222222222222222222222222222222222")

	type (
		messageToTokenDataFn func(context.Context, logger.Logger, CCTPv2Message, AttestationEncoder) exectypes.TokenData
		tokenDataMap         map[CCTPv2RequestParams]map[DepositHash][]exectypes.TokenData
	)

	tests := []struct {
		name                  string
		input                 map[CCTPv2RequestParams]CCTPv2Messages
		setupCalculateHashFn  func() func(CCTPv2DecodedMessage) ([32]byte, error)
		setupMessageToTokenFn func() messageToTokenDataFn
		validate              func(*testing.T, tokenDataMap)
	}{
		{
			name:  "empty input map returns empty result",
			input: map[CCTPv2RequestParams]CCTPv2Messages{},
			setupCalculateHashFn: func() func(CCTPv2DecodedMessage) ([32]byte, error) {
				return func(msg CCTPv2DecodedMessage) ([32]byte, error) {
					t.Fatal("should not be called for empty input")
					return [32]byte{}, nil
				}
			},
			setupMessageToTokenFn: func() messageToTokenDataFn {
				return func(_ context.Context, _ logger.Logger, msg CCTPv2Message, _ AttestationEncoder) exectypes.TokenData {
					t.Fatal("should not be called for empty input")
					return exectypes.TokenData{}
				}
			},
			validate: func(t *testing.T, result tokenDataMap) {
				assert.Equal(t, 0, len(result))
			},
		},
		{
			name: "single request with single valid message",
			input: map[CCTPv2RequestParams]CCTPv2Messages{
				{chainSelector: testChain1, sourceDomain: 100, txHash: testTxHash1}: {
					Messages: []CCTPv2Message{
						{
							Status:         "complete",
							Message:        "1234567890abcdef",
							Attestation:    "fedcba0987654321",
							DecodedMessage: validDecodedMessage,
						},
					},
				},
			},
			setupCalculateHashFn: func() func(CCTPv2DecodedMessage) ([32]byte, error) {
				return func(msg CCTPv2DecodedMessage) ([32]byte, error) {
					return hash1, nil
				}
			},
			setupMessageToTokenFn: func() messageToTokenDataFn {
				return func(_ context.Context, _ logger.Logger, msg CCTPv2Message, _ AttestationEncoder) exectypes.TokenData {
					return exectypes.NewSuccessTokenData(cciptypes.Bytes("test-data-1"))
				}
			},
			validate: func(t *testing.T, result tokenDataMap) {
				assert.Equal(t, 1, len(result))
				params := CCTPv2RequestParams{chainSelector: testChain1, sourceDomain: 100, txHash: testTxHash1}
				assert.Contains(t, result, params)
				assert.Equal(t, 1, len(result[params]))
				assert.Contains(t, result[params], hash1)
				assert.Equal(t, 1, len(result[params][hash1]))
				assert.True(t, result[params][hash1][0].Ready)
				assert.Equal(t, cciptypes.Bytes("test-data-1"), result[params][hash1][0].Data)
			},
		},
		{
			name: "single request with multiple messages different deposit hashes",
			input: map[CCTPv2RequestParams]CCTPv2Messages{
				{chainSelector: testChain1, sourceDomain: 100, txHash: testTxHash1}: {
					Messages: []CCTPv2Message{
						{
							Status:         "complete",
							Message:        "1234567890abcdef",
							Attestation:    "fedcba0987654321",
							DecodedMessage: validDecodedMessage,
						},
						{
							Status:         "complete",
							Message:        "abcdef1234567890",
							Attestation:    "098765fedcba4321",
							DecodedMessage: validDecodedMessage2,
						},
					},
				},
			},
			setupCalculateHashFn: func() func(CCTPv2DecodedMessage) ([32]byte, error) {
				callCount := 0
				return func(msg CCTPv2DecodedMessage) ([32]byte, error) {
					callCount++
					if callCount == 1 {
						return hash1, nil
					}
					return hash2, nil
				}
			},
			setupMessageToTokenFn: func() messageToTokenDataFn {
				callCount := 0
				return func(_ context.Context, _ logger.Logger, msg CCTPv2Message, _ AttestationEncoder) exectypes.TokenData {
					callCount++
					return exectypes.NewSuccessTokenData(cciptypes.Bytes(fmt.Sprintf("test-data-%d", callCount)))
				}
			},
			validate: func(t *testing.T, result tokenDataMap) {
				assert.Equal(t, 1, len(result))
				params := CCTPv2RequestParams{chainSelector: testChain1, sourceDomain: 100, txHash: testTxHash1}
				assert.Equal(t, 2, len(result[params]))
				assert.Contains(t, result[params], hash1)
				assert.Contains(t, result[params], hash2)
				assert.Equal(t, 1, len(result[params][hash1]))
				assert.Equal(t, 1, len(result[params][hash2]))
			},
		},
		{
			name: "multiple messages with same deposit hash appends to slice",
			input: map[CCTPv2RequestParams]CCTPv2Messages{
				{chainSelector: testChain1, sourceDomain: 100, txHash: testTxHash1}: {
					Messages: []CCTPv2Message{
						{
							Status:         "complete",
							Message:        "1234567890abcdef",
							Attestation:    "fedcba0987654321",
							DecodedMessage: validDecodedMessage,
						},
						{
							Status:         "complete",
							Message:        "aabbccddeeff1122",
							Attestation:    "112233445566",
							DecodedMessage: validDecodedMessage,
						},
					},
				},
			},
			setupCalculateHashFn: func() func(CCTPv2DecodedMessage) ([32]byte, error) {
				return func(msg CCTPv2DecodedMessage) ([32]byte, error) {
					return hash1, nil
				}
			},
			setupMessageToTokenFn: func() messageToTokenDataFn {
				callCount := 0
				return func(_ context.Context, _ logger.Logger, msg CCTPv2Message, _ AttestationEncoder) exectypes.TokenData {
					callCount++
					return exectypes.NewSuccessTokenData(cciptypes.Bytes(fmt.Sprintf("data-%d", callCount)))
				}
			},
			validate: func(t *testing.T, result tokenDataMap) {
				params := CCTPv2RequestParams{chainSelector: testChain1, sourceDomain: 100, txHash: testTxHash1}
				assert.Equal(t, 1, len(result[params]))
				assert.Equal(t, 2, len(result[params][hash1]))
				assert.Equal(t, cciptypes.Bytes("data-1"), result[params][hash1][0].Data)
				assert.Equal(t, cciptypes.Bytes("data-2"), result[params][hash1][1].Data)
			},
		},
		{
			name: "multiple requests processed independently",
			input: map[CCTPv2RequestParams]CCTPv2Messages{
				{chainSelector: testChain1, sourceDomain: 100, txHash: testTxHash1}: {
					Messages: []CCTPv2Message{
						{
							Status:         "complete",
							Message:        "1234567890abcdef",
							Attestation:    "fedcba0987654321",
							DecodedMessage: validDecodedMessage,
						},
					},
				},
				{chainSelector: testChain2, sourceDomain: 200, txHash: testTxHash2}: {
					Messages: []CCTPv2Message{
						{
							Status:         "complete",
							Message:        "abcdef1234567890",
							Attestation:    "098765fedcba4321",
							DecodedMessage: validDecodedMessage2,
						},
					},
				},
			},
			setupCalculateHashFn: func() func(CCTPv2DecodedMessage) ([32]byte, error) {
				return func(msg CCTPv2DecodedMessage) ([32]byte, error) {
					if msg.DestinationDomain == "2" {
						return hash1, nil
					}
					return hash2, nil
				}
			},
			setupMessageToTokenFn: func() messageToTokenDataFn {
				return func(_ context.Context, _ logger.Logger, msg CCTPv2Message, _ AttestationEncoder) exectypes.TokenData {
					return exectypes.NewSuccessTokenData(cciptypes.Bytes("data"))
				}
			},
			validate: func(t *testing.T, result tokenDataMap) {
				assert.Equal(t, 2, len(result))
				params1 := CCTPv2RequestParams{chainSelector: testChain1, sourceDomain: 100, txHash: testTxHash1}
				params2 := CCTPv2RequestParams{chainSelector: testChain2, sourceDomain: 200, txHash: testTxHash2}
				assert.Contains(t, result, params1)
				assert.Contains(t, result, params2)
			},
		},
		{
			name: "calculateDepositHash error skips message and continues",
			input: map[CCTPv2RequestParams]CCTPv2Messages{
				{chainSelector: testChain1, sourceDomain: 100, txHash: testTxHash1}: {
					Messages: []CCTPv2Message{
						{
							Status:         "complete",
							Message:        "1234567890abcdef",
							Attestation:    "fedcba0987654321",
							DecodedMessage: validDecodedMessage,
						},
						{
							Status:         "complete",
							Message:        "abcdef1234567890",
							Attestation:    "098765fedcba4321",
							DecodedMessage: validDecodedMessage2,
						},
					},
				},
			},
			setupCalculateHashFn: func() func(CCTPv2DecodedMessage) ([32]byte, error) {
				callCount := 0
				return func(msg CCTPv2DecodedMessage) ([32]byte, error) {
					callCount++
					if callCount == 1 {
						return [32]byte{}, fmt.Errorf("hash calculation failed")
					}
					return hash2, nil
				}
			},
			setupMessageToTokenFn: func() messageToTokenDataFn {
				callCount := 0
				return func(_ context.Context, _ logger.Logger, msg CCTPv2Message, _ AttestationEncoder) exectypes.TokenData {
					callCount++
					if callCount > 1 {
						t.Fatal("should only be called once, second message has hash error")
					}
					return exectypes.NewSuccessTokenData(cciptypes.Bytes("data"))
				}
			},
			validate: func(t *testing.T, result tokenDataMap) {
				params := CCTPv2RequestParams{chainSelector: testChain1, sourceDomain: 100, txHash: testTxHash1}
				assert.Equal(t, 1, len(result[params]))
				assert.Contains(t, result[params], hash2)
				assert.Equal(t, 1, len(result[params][hash2]))
			},
		},
		{
			name: "all messages fail calculateDepositHash returns empty inner map",
			input: map[CCTPv2RequestParams]CCTPv2Messages{
				{chainSelector: testChain1, sourceDomain: 100, txHash: testTxHash1}: {
					Messages: []CCTPv2Message{
						{
							Status:         "complete",
							Message:        "1234567890abcdef",
							Attestation:    "fedcba0987654321",
							DecodedMessage: validDecodedMessage,
						},
					},
				},
			},
			setupCalculateHashFn: func() func(CCTPv2DecodedMessage) ([32]byte, error) {
				return func(msg CCTPv2DecodedMessage) ([32]byte, error) {
					return [32]byte{}, fmt.Errorf("all hash calculations fail")
				}
			},
			setupMessageToTokenFn: func() messageToTokenDataFn {
				return func(_ context.Context, _ logger.Logger, msg CCTPv2Message, _ AttestationEncoder) exectypes.TokenData {
					t.Fatal("should not be called when hash calculation fails")
					return exectypes.TokenData{}
				}
			},
			validate: func(t *testing.T, result tokenDataMap) {
				params := CCTPv2RequestParams{chainSelector: testChain1, sourceDomain: 100, txHash: testTxHash1}
				assert.Contains(t, result, params)
				assert.Equal(t, 0, len(result[params]))
			},
		},
		{
			name: "request with empty Messages slice creates empty inner map",
			input: map[CCTPv2RequestParams]CCTPv2Messages{
				{chainSelector: testChain1, sourceDomain: 100, txHash: testTxHash1}: {
					Messages: []CCTPv2Message{},
				},
			},
			setupCalculateHashFn: func() func(CCTPv2DecodedMessage) ([32]byte, error) {
				return func(msg CCTPv2DecodedMessage) ([32]byte, error) {
					t.Fatal("should not be called for empty messages")
					return [32]byte{}, nil
				}
			},
			setupMessageToTokenFn: func() messageToTokenDataFn {
				return func(_ context.Context, _ logger.Logger, msg CCTPv2Message, _ AttestationEncoder) exectypes.TokenData {
					t.Fatal("should not be called for empty messages")
					return exectypes.TokenData{}
				}
			},
			validate: func(t *testing.T, result tokenDataMap) {
				params := CCTPv2RequestParams{chainSelector: testChain1, sourceDomain: 100, txHash: testTxHash1}
				assert.Contains(t, result, params)
				assert.Equal(t, 0, len(result[params]))
			},
		},
		{
			name: "messageToTokenData returns error TokenData",
			input: map[CCTPv2RequestParams]CCTPv2Messages{
				{chainSelector: testChain1, sourceDomain: 100, txHash: testTxHash1}: {
					Messages: []CCTPv2Message{
						{
							Status:         "pending",
							Message:        "1234567890abcdef",
							Attestation:    "fedcba0987654321",
							DecodedMessage: validDecodedMessage,
						},
					},
				},
			},
			setupCalculateHashFn: func() func(CCTPv2DecodedMessage) ([32]byte, error) {
				return func(msg CCTPv2DecodedMessage) ([32]byte, error) {
					return hash1, nil
				}
			},
			setupMessageToTokenFn: func() messageToTokenDataFn {
				return func(_ context.Context, _ logger.Logger, msg CCTPv2Message, _ AttestationEncoder) exectypes.TokenData {
					return exectypes.NewErrorTokenData(tokendata.ErrNotReady)
				}
			},
			validate: func(t *testing.T, result tokenDataMap) {
				params := CCTPv2RequestParams{chainSelector: testChain1, sourceDomain: 100, txHash: testTxHash1}
				assert.Equal(t, 1, len(result[params][hash1]))
				assert.False(t, result[params][hash1][0].Ready)
				assert.ErrorIs(t, result[params][hash1][0].Error, tokendata.ErrNotReady)
			},
		},
		{
			name: "context passed to messageToTokenData",
			input: map[CCTPv2RequestParams]CCTPv2Messages{
				{chainSelector: testChain1, sourceDomain: 100, txHash: testTxHash1}: {
					Messages: []CCTPv2Message{
						{
							Status:         "complete",
							Message:        "1234567890abcdef",
							Attestation:    "fedcba0987654321",
							DecodedMessage: validDecodedMessage,
						},
					},
				},
			},
			setupCalculateHashFn: func() func(CCTPv2DecodedMessage) ([32]byte, error) {
				return func(msg CCTPv2DecodedMessage) ([32]byte, error) {
					return hash1, nil
				}
			},
			setupMessageToTokenFn: func() messageToTokenDataFn {
				return func(ctx context.Context, _ logger.Logger, _ CCTPv2Message, _ AttestationEncoder) exectypes.TokenData {
					assert.NotNil(t, ctx)
					return exectypes.NewSuccessTokenData(cciptypes.Bytes("data"))
				}
			},
			validate: func(t *testing.T, result tokenDataMap) {
				params := CCTPv2RequestParams{chainSelector: testChain1, sourceDomain: 100, txHash: testTxHash1}
				assert.Equal(t, 1, len(result[params][hash1]))
				assert.True(t, result[params][hash1][0].Ready)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			observer := &CCTPv2TokenDataObserver{
				lggr:                   logger.Test(t),
				calculateDepositHashFn: tt.setupCalculateHashFn(),
				messageToTokenDataFn:   tt.setupMessageToTokenFn(),
				attestationEncoder: func(ctx context.Context, msg, att cciptypes.Bytes) (cciptypes.Bytes, error) {
					return cciptypes.Bytes("encoded"), nil
				},
			}

			result := observer.convertCCTPv2MessagesToTokenData(context.Background(), tt.input)

			tt.validate(t, result)
		})
	}
}

func TestCCTPv2TokenDataObserver_AssignSingleTokenData(t *testing.T) {
	const (
		testChain1   = cciptypes.ChainSelector(1)
		testPoolAddr = "0x1234567890123456789012345678901234567890"
		testTxHash1  = "0xabcdef0123456789abcdef0123456789abcdef0123456789abcdef0123456789"
		depositHex1  = "1111111111111111111111111111111111111111111111111111111111111111"
		depositHex2  = "2222222222222222222222222222222222222222222222222222222222222222"
	)

	hash1 := mustHexToBytes32(depositHex1)
	hash2 := mustHexToBytes32(depositHex2)

	type tokenDataMap = map[CCTPv2RequestParams]map[DepositHash][]exectypes.TokenData

	tests := []struct {
		name          string
		chainSelector cciptypes.ChainSelector
		txHash        TxHash
		tokenAmount   cciptypes.RampTokenAmount
		tokenData     tokenDataMap
		poolConfig    map[cciptypes.ChainSelector]string
		validate      func(*testing.T, exectypes.TokenData, tokenDataMap)
	}{
		{
			name:          "unsupported token - wrong pool address",
			chainSelector: testChain1,
			txHash:        testTxHash1,
			tokenAmount:   createCCTPv2Token("0x9999999999999999999999999999999999999999", 100, depositHex1),
			tokenData:     map[CCTPv2RequestParams]map[DepositHash][]exectypes.TokenData{},
			poolConfig: map[cciptypes.ChainSelector]string{
				testChain1: testPoolAddr,
			},
			validate: func(t *testing.T, result exectypes.TokenData, tokenData tokenDataMap) {
				assert.False(t, result.Ready)
				assert.False(t, result.Supported)
				assert.Nil(t, result.Error)
				// tokenData not modified for unsupported tokens
			},
		},
		{
			name:          "unsupported token - invalid ExtraData",
			chainSelector: testChain1,
			txHash:        testTxHash1,
			tokenAmount: cciptypes.RampTokenAmount{
				SourcePoolAddress: mustDecodeAddress(testPoolAddr),
				ExtraData:         cciptypes.Bytes{0x01, 0x02, 0x03}, // Invalid length
			},
			tokenData:  map[CCTPv2RequestParams]map[DepositHash][]exectypes.TokenData{},
			poolConfig: map[cciptypes.ChainSelector]string{testChain1: testPoolAddr},
			validate: func(t *testing.T, result exectypes.TokenData, tokenData tokenDataMap) {
				assert.False(t, result.Ready)
				assert.False(t, result.Supported)
				assert.Nil(t, result.Error)
			},
		},
		{
			name:          "supported token with data found",
			chainSelector: testChain1,
			txHash:        testTxHash1,
			tokenAmount:   createCCTPv2Token(testPoolAddr, 100, depositHex1),
			tokenData: map[CCTPv2RequestParams]map[DepositHash][]exectypes.TokenData{
				{chainSelector: testChain1, sourceDomain: 100, txHash: testTxHash1}: {
					hash1: {
						exectypes.NewSuccessTokenData(cciptypes.Bytes("attestation-1")),
					},
				},
			},
			poolConfig: map[cciptypes.ChainSelector]string{testChain1: testPoolAddr},
			validate: func(t *testing.T, result exectypes.TokenData, tokenData tokenDataMap) {
				assert.True(t, result.Ready)
				assert.True(t, result.Supported)
				assert.Equal(t, cciptypes.Bytes("attestation-1"), result.Data)
				assert.NoError(t, result.Error)
				// Verify item was removed from tokenData
				params := CCTPv2RequestParams{chainSelector: testChain1, sourceDomain: 100, txHash: testTxHash1}
				assert.Len(t, tokenData[params][hash1], 0)
			},
		},
		{
			name:          "request params not found in tokenData map",
			chainSelector: testChain1,
			txHash:        testTxHash1,
			tokenAmount:   createCCTPv2Token(testPoolAddr, 100, depositHex1),
			tokenData:     map[CCTPv2RequestParams]map[DepositHash][]exectypes.TokenData{},
			poolConfig:    map[cciptypes.ChainSelector]string{testChain1: testPoolAddr},
			validate: func(t *testing.T, result exectypes.TokenData, tokenData tokenDataMap) {
				assert.False(t, result.Ready)
				assert.True(t, result.Supported)
				assert.ErrorIs(t, result.Error, tokendata.ErrDataMissing)
			},
		},
		{
			name:          "deposit hash not found in tokenData map",
			chainSelector: testChain1,
			txHash:        testTxHash1,
			tokenAmount:   createCCTPv2Token(testPoolAddr, 100, depositHex1),
			tokenData: map[CCTPv2RequestParams]map[DepositHash][]exectypes.TokenData{
				{chainSelector: testChain1, sourceDomain: 100, txHash: testTxHash1}: {
					hash2: { // Different hash
						exectypes.NewSuccessTokenData(cciptypes.Bytes("attestation")),
					},
				},
			},
			poolConfig: map[cciptypes.ChainSelector]string{testChain1: testPoolAddr},
			validate: func(t *testing.T, result exectypes.TokenData, tokenData tokenDataMap) {
				assert.False(t, result.Ready)
				assert.True(t, result.Supported)
				assert.ErrorIs(t, result.Error, tokendata.ErrDataMissing)
			},
		},
		{
			name:          "all token data consumed (empty slice)",
			chainSelector: testChain1,
			txHash:        testTxHash1,
			tokenAmount:   createCCTPv2Token(testPoolAddr, 100, depositHex1),
			tokenData: map[CCTPv2RequestParams]map[DepositHash][]exectypes.TokenData{
				{chainSelector: testChain1, sourceDomain: 100, txHash: testTxHash1}: {
					hash1: {}, // Empty slice
				},
			},
			poolConfig: map[cciptypes.ChainSelector]string{testChain1: testPoolAddr},
			validate: func(t *testing.T, result exectypes.TokenData, tokenData tokenDataMap) {
				assert.False(t, result.Ready)
				assert.True(t, result.Supported)
				assert.ErrorIs(t, result.Error, tokendata.ErrDataMissing)
			},
		},
		{
			name:          "consumes first item and leaves second",
			chainSelector: testChain1,
			txHash:        testTxHash1,
			tokenAmount:   createCCTPv2Token(testPoolAddr, 100, depositHex1),
			tokenData: map[CCTPv2RequestParams]map[DepositHash][]exectypes.TokenData{
				{chainSelector: testChain1, sourceDomain: 100, txHash: testTxHash1}: {
					hash1: {
						exectypes.NewSuccessTokenData(cciptypes.Bytes("attestation-1")),
						exectypes.NewSuccessTokenData(cciptypes.Bytes("attestation-2")),
					},
				},
			},
			poolConfig: map[cciptypes.ChainSelector]string{testChain1: testPoolAddr},
			validate: func(t *testing.T, result exectypes.TokenData, tokenData tokenDataMap) {
				assert.True(t, result.Ready)
				assert.Equal(t, cciptypes.Bytes("attestation-1"), result.Data) // Gets first item
				// Verify second item remains
				params := CCTPv2RequestParams{chainSelector: testChain1, sourceDomain: 100, txHash: testTxHash1}
				assert.Len(t, tokenData[params][hash1], 1)
				assert.Equal(t, cciptypes.Bytes("attestation-2"), tokenData[params][hash1][0].Data)
			},
		},
		{
			name:          "multiple calls consume items in order",
			chainSelector: testChain1,
			txHash:        testTxHash1,
			tokenAmount:   createCCTPv2Token(testPoolAddr, 100, depositHex1),
			tokenData: map[CCTPv2RequestParams]map[DepositHash][]exectypes.TokenData{
				{chainSelector: testChain1, sourceDomain: 100, txHash: testTxHash1}: {
					hash1: {
						exectypes.NewSuccessTokenData(cciptypes.Bytes("first")),
						exectypes.NewSuccessTokenData(cciptypes.Bytes("second")),
						exectypes.NewSuccessTokenData(cciptypes.Bytes("third")),
					},
				},
			},
			poolConfig: map[cciptypes.ChainSelector]string{testChain1: testPoolAddr},
			validate: func(t *testing.T, result exectypes.TokenData, tokenData tokenDataMap) {
				// First call gets "first", second call gets "second", third gets "third"
				assert.Equal(t, cciptypes.Bytes("first"), result.Data)
				params := CCTPv2RequestParams{chainSelector: testChain1, sourceDomain: 100, txHash: testTxHash1}
				assert.Len(t, tokenData[params][hash1], 2) // Two remaining
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			observer := &CCTPv2TokenDataObserver{
				lggr:                     logger.Test(t),
				supportedPoolsBySelector: tt.poolConfig,
			}

			result := observer.assignSingleTokenData(
				tt.chainSelector,
				tt.txHash,
				tt.tokenAmount,
				tt.tokenData,
			)

			tt.validate(t, result, tt.tokenData)
		})
	}
}

func TestCCTPv2TokenDataObserver_AssignTokenData(t *testing.T) {
	const (
		testChain1   = cciptypes.ChainSelector(1)
		testChain2   = cciptypes.ChainSelector(2)
		testPoolAddr = "0x1234567890123456789012345678901234567890"
		testTxHash1  = "0xabcd"
		testTxHash2  = "0xef01"
		depositHex1  = "1111111111111111111111111111111111111111111111111111111111111111"
		depositHex2  = "2222222222222222222222222222222222222222222222222222222222222222"
	)

	hash1 := mustHexToBytes32(depositHex1)
	hash2 := mustHexToBytes32(depositHex2)

	tests := []struct {
		name      string
		messages  exectypes.MessageObservations
		tokenData map[CCTPv2RequestParams]map[DepositHash][]exectypes.TokenData
		validate  func(t *testing.T, result exectypes.TokenDataObservations)
	}{
		{
			name:      "empty messages",
			messages:  exectypes.MessageObservations{},
			tokenData: make(map[CCTPv2RequestParams]map[DepositHash][]exectypes.TokenData),
			validate: func(t *testing.T, result exectypes.TokenDataObservations) {
				assert.Empty(t, result)
			},
		},
		{
			name: "single message with single token - data found",
			messages: exectypes.MessageObservations{
				testChain1: {
					10: createTestMessage(testTxHash1, []cciptypes.RampTokenAmount{
						createCCTPv2Token(testPoolAddr, 100, depositHex1),
					}),
				},
			},
			tokenData: map[CCTPv2RequestParams]map[DepositHash][]exectypes.TokenData{
				{chainSelector: testChain1, sourceDomain: 100, txHash: testTxHash1}: {
					hash1: {exectypes.NewSuccessTokenData([]byte("success"))},
				},
			},
			validate: func(t *testing.T, result exectypes.TokenDataObservations) {
				assert.Contains(t, result, testChain1)
				assert.Contains(t, result[testChain1], cciptypes.SeqNum(10))
				tokenData := result[testChain1][10].TokenData
				assert.Len(t, tokenData, 1)
				assert.True(t, tokenData[0].Ready)
				assert.True(t, tokenData[0].Supported)
				assert.Equal(t, []byte("success"), []byte(tokenData[0].Data))
			},
		},
		{
			name: "single message with multiple tokens",
			messages: exectypes.MessageObservations{
				testChain1: {
					10: createTestMessage(testTxHash1, []cciptypes.RampTokenAmount{
						createCCTPv2Token(testPoolAddr, 100, depositHex1),
						createCCTPv2Token(testPoolAddr, 100, depositHex1),
					}),
				},
			},
			tokenData: map[CCTPv2RequestParams]map[DepositHash][]exectypes.TokenData{
				{chainSelector: testChain1, sourceDomain: 100, txHash: testTxHash1}: {
					hash1: {
						exectypes.NewSuccessTokenData([]byte("first")),
						exectypes.NewSuccessTokenData([]byte("second")),
					},
				},
			},
			validate: func(t *testing.T, result exectypes.TokenDataObservations) {
				assert.Contains(t, result, testChain1)
				assert.Contains(t, result[testChain1], cciptypes.SeqNum(10))
				tokenData := result[testChain1][10].TokenData
				assert.Len(t, tokenData, 2)
				assert.Equal(t, []byte("first"), []byte(tokenData[0].Data))
				assert.Equal(t, []byte("second"), []byte(tokenData[1].Data))
			},
		},
		{
			name: "multiple messages in single chain",
			messages: exectypes.MessageObservations{
				testChain1: {
					10: createTestMessage(testTxHash1, []cciptypes.RampTokenAmount{
						createCCTPv2Token(testPoolAddr, 100, depositHex1),
					}),
					11: createTestMessage(testTxHash2, []cciptypes.RampTokenAmount{
						createCCTPv2Token(testPoolAddr, 100, depositHex2),
					}),
				},
			},
			tokenData: map[CCTPv2RequestParams]map[DepositHash][]exectypes.TokenData{
				{chainSelector: testChain1, sourceDomain: 100, txHash: testTxHash1}: {
					hash1: {exectypes.NewSuccessTokenData([]byte("msg10"))},
				},
				{chainSelector: testChain1, sourceDomain: 100, txHash: testTxHash2}: {
					hash2: {exectypes.NewSuccessTokenData([]byte("msg11"))},
				},
			},
			validate: func(t *testing.T, result exectypes.TokenDataObservations) {
				assert.Contains(t, result, testChain1)
				assert.Contains(t, result[testChain1], cciptypes.SeqNum(10))
				assert.Contains(t, result[testChain1], cciptypes.SeqNum(11))
				assert.Equal(t, []byte("msg10"), []byte(result[testChain1][10].TokenData[0].Data))
				assert.Equal(t, []byte("msg11"), []byte(result[testChain1][11].TokenData[0].Data))
			},
		},
		{
			name: "multiple chains processed independently",
			messages: exectypes.MessageObservations{
				testChain1: {
					10: createTestMessage(testTxHash1, []cciptypes.RampTokenAmount{
						createCCTPv2Token(testPoolAddr, 100, depositHex1),
					}),
				},
				testChain2: {
					20: createTestMessage(testTxHash2, []cciptypes.RampTokenAmount{
						createCCTPv2Token(testPoolAddr, 100, depositHex2),
					}),
				},
			},
			tokenData: map[CCTPv2RequestParams]map[DepositHash][]exectypes.TokenData{
				{chainSelector: testChain1, sourceDomain: 100, txHash: testTxHash1}: {
					hash1: {exectypes.NewSuccessTokenData([]byte("chain1"))},
				},
				{chainSelector: testChain2, sourceDomain: 100, txHash: testTxHash2}: {
					hash2: {exectypes.NewSuccessTokenData([]byte("chain2"))},
				},
			},
			validate: func(t *testing.T, result exectypes.TokenDataObservations) {
				assert.Contains(t, result, testChain1)
				assert.Contains(t, result, testChain2)
				assert.Equal(t, []byte("chain1"), []byte(result[testChain1][10].TokenData[0].Data))
				assert.Equal(t, []byte("chain2"), []byte(result[testChain2][20].TokenData[0].Data))
			},
		},
		{
			name: "mixed supported and unsupported tokens",
			messages: exectypes.MessageObservations{
				testChain1: {
					10: createTestMessage(testTxHash1, []cciptypes.RampTokenAmount{
						createCCTPv2Token(testPoolAddr, 100, depositHex1),
						createCCTPv2Token("0x9999999999999999999999999999999999999999", 100, depositHex2),
					}),
				},
			},
			tokenData: map[CCTPv2RequestParams]map[DepositHash][]exectypes.TokenData{
				{chainSelector: testChain1, sourceDomain: 100, txHash: testTxHash1}: {
					hash1: {exectypes.NewSuccessTokenData([]byte("supported"))},
				},
			},
			validate: func(t *testing.T, result exectypes.TokenDataObservations) {
				assert.Contains(t, result, testChain1)
				assert.Contains(t, result[testChain1], cciptypes.SeqNum(10))
				tokenData := result[testChain1][10].TokenData
				assert.Len(t, tokenData, 2)
				assert.True(t, tokenData[0].Ready)
				assert.True(t, tokenData[0].Supported)
				assert.Equal(t, []byte("supported"), []byte(tokenData[0].Data))
				assert.False(t, tokenData[1].Supported)
			},
		},
		{
			name: "preserves structure for missing data",
			messages: exectypes.MessageObservations{
				testChain1: {
					10: createTestMessage(testTxHash1, []cciptypes.RampTokenAmount{
						createCCTPv2Token(testPoolAddr, 100, depositHex1),
					}),
				},
			},
			tokenData: make(map[CCTPv2RequestParams]map[DepositHash][]exectypes.TokenData),
			validate: func(t *testing.T, result exectypes.TokenDataObservations) {
				assert.Contains(t, result, testChain1)
				assert.Contains(t, result[testChain1], cciptypes.SeqNum(10))
				tokenData := result[testChain1][10].TokenData
				assert.Len(t, tokenData, 1)
				assert.False(t, tokenData[0].Ready)
				assert.True(t, tokenData[0].Supported)
				assert.ErrorIs(t, tokenData[0].Error, tokendata.ErrDataMissing)
			},
		},
		{
			name: "FIFO consumption - multiple tokens with same deposit hash",
			messages: exectypes.MessageObservations{
				testChain1: {
					10: createTestMessage(testTxHash1, []cciptypes.RampTokenAmount{
						createCCTPv2Token(testPoolAddr, 100, depositHex1),
						createCCTPv2Token(testPoolAddr, 100, depositHex1),
						createCCTPv2Token(testPoolAddr, 100, depositHex1),
					}),
				},
			},
			tokenData: map[CCTPv2RequestParams]map[DepositHash][]exectypes.TokenData{
				{chainSelector: testChain1, sourceDomain: 100, txHash: testTxHash1}: {
					hash1: {
						exectypes.NewSuccessTokenData([]byte("first")),
						exectypes.NewSuccessTokenData([]byte("second")),
						exectypes.NewSuccessTokenData([]byte("third")),
					},
				},
			},
			validate: func(t *testing.T, result exectypes.TokenDataObservations) {
				assert.Contains(t, result, testChain1)
				assert.Contains(t, result[testChain1], cciptypes.SeqNum(10))
				// Check that all three tokens consumed the data in FIFO order
				tokenData := result[testChain1][10].TokenData
				assert.Len(t, tokenData, 3)
				assert.Equal(t, []byte("first"), []byte(tokenData[0].Data))
				assert.Equal(t, []byte("second"), []byte(tokenData[1].Data))
				assert.Equal(t, []byte("third"), []byte(tokenData[2].Data))
			},
		},
		{
			name: "preserves sequence numbers correctly",
			messages: exectypes.MessageObservations{
				testChain1: {
					100: createTestMessage(testTxHash1, []cciptypes.RampTokenAmount{
						createCCTPv2Token(testPoolAddr, 100, depositHex1),
					}),
					999: createTestMessage(testTxHash2, []cciptypes.RampTokenAmount{
						createCCTPv2Token(testPoolAddr, 100, depositHex2),
					}),
				},
			},
			tokenData: map[CCTPv2RequestParams]map[DepositHash][]exectypes.TokenData{
				{chainSelector: testChain1, sourceDomain: 100, txHash: testTxHash1}: {
					hash1: {exectypes.NewSuccessTokenData([]byte("seq100"))},
				},
				{chainSelector: testChain1, sourceDomain: 100, txHash: testTxHash2}: {
					hash2: {exectypes.NewSuccessTokenData([]byte("seq999"))},
				},
			},
			validate: func(t *testing.T, result exectypes.TokenDataObservations) {
				assert.Contains(t, result, testChain1)
				assert.Contains(t, result[testChain1], cciptypes.SeqNum(100))
				assert.Contains(t, result[testChain1], cciptypes.SeqNum(999))
				assert.Equal(t, []byte("seq100"), []byte(result[testChain1][100].TokenData[0].Data))
				assert.Equal(t, []byte("seq999"), []byte(result[testChain1][999].TokenData[0].Data))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			observer := &CCTPv2TokenDataObserver{
				lggr: logger.Test(t),
				supportedPoolsBySelector: map[cciptypes.ChainSelector]string{
					testChain1: testPoolAddr,
					testChain2: testPoolAddr,
				},
			}

			result := observer.assignTokenData(tt.messages, tt.tokenData)
			tt.validate(t, result)
		})
	}
}
