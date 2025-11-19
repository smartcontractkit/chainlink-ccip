package v2

import (
	"context"
	"encoding/binary"
	"encoding/hex"
	"strings"
	"testing"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
	"github.com/stretchr/testify/assert"

	"github.com/smartcontractkit/chainlink-ccip/execute/exectypes"
)

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
			observer := NewCCTPv2TokenDataObserver(
				logger.Test(t),
				cciptypes.ChainSelector(999),
				map[cciptypes.ChainSelector]string{},
				nil,
				nil,
			)

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

			observer := NewCCTPv2TokenDataObserver(
				logger.Test(t),
				cciptypes.ChainSelector(999),
				map[cciptypes.ChainSelector]string{},
				nil,
				mock,
			)

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
