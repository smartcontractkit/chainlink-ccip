package v2

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	commonconfig "github.com/smartcontractkit/chainlink-common/pkg/config"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	httputil "github.com/smartcontractkit/chainlink-ccip/execute/tokendata/http"
	"github.com/smartcontractkit/chainlink-ccip/pluginconfig"
)

const (
	testTxHash = "0x1234567890123456789012345678901234567890123456789012345678901234"
)

// MockHTTPClient is a mock implementation of httputil.HTTPClient
type MockHTTPClient struct {
	mock.Mock
}

func (m *MockHTTPClient) Get(ctx context.Context, path string) (cciptypes.Bytes, httputil.HTTPStatus, error) {
	args := m.Called(ctx, path)
	return args.Get(0).(cciptypes.Bytes), args.Get(1).(httputil.HTTPStatus), args.Error(2)
}

func (m *MockHTTPClient) Post(
	ctx context.Context, path string, requestData cciptypes.Bytes,
) (cciptypes.Bytes, httputil.HTTPStatus, error) {
	args := m.Called(ctx, path, requestData)
	return args.Get(0).(cciptypes.Bytes), args.Get(1).(httputil.HTTPStatus), args.Error(2)
}

// MockMetricsReporter is a mock implementation of MetricsReporter
type MockMetricsReporter struct {
	mock.Mock
}

func (m *MockMetricsReporter) TrackAttestationAPILatency(
	sourceChain cciptypes.ChainSelector,
	sourceDomain uint32,
	success bool,
	httpStatus string,
	latency time.Duration,
) {
	m.Called(sourceChain, sourceDomain, success, httpStatus, latency)
}

func TestNewCCTPv2Client(t *testing.T) {
	lggr := logger.Test(t)
	config := pluginconfig.USDCCCTPObserverConfig{
		AttestationConfig: pluginconfig.AttestationConfig{
			AttestationAPI:         "https://iris-api-sandbox.circle.com/v2",
			AttestationAPITimeout:  commonconfig.MustNewDuration(time.Second),
			AttestationAPIInterval: commonconfig.MustNewDuration(100 * time.Millisecond),
		},
		AttestationAPICooldown: commonconfig.MustNewDuration(5 * time.Minute),
	}
	metricsReporter := &MockMetricsReporter{}

	t.Run("Valid inputs create client successfully", func(t *testing.T) {
		client, err := NewCCTPv2Client(lggr, config, metricsReporter)
		require.NoError(t, err)
		assert.NotNil(t, client)
		assert.NotNil(t, client.lggr)
		assert.NotNil(t, client.client)
		assert.NotNil(t, client.metricsReporter)
	})

	t.Run("Nil logger returns error", func(t *testing.T) {
		client, err := NewCCTPv2Client(nil, config, metricsReporter)
		assert.Error(t, err)
		assert.Nil(t, client)
		assert.Contains(t, err.Error(), "logger cannot be nil")
	})

	t.Run("Nil metricsReporter returns error", func(t *testing.T) {
		client, err := NewCCTPv2Client(lggr, config, nil)
		assert.Error(t, err)
		assert.Nil(t, client)
		assert.Contains(t, err.Error(), "metricsReporter cannot be nil")
	})
}

func TestGetMessages_Success(t *testing.T) {
	lggr := logger.Test(t)
	mockHTTPClient := &MockHTTPClient{}
	mockMetrics := &MockMetricsReporter{}

	client := &CCTPv2HTTPClientImpl{
		lggr:            lggr,
		client:          mockHTTPClient,
		metricsReporter: mockMetrics,
	}

	ctx := context.Background()
	sourceChain := cciptypes.ChainSelector(1)
	sourceDomainID := uint32(0)
	txHash := testTxHash

	// Create response
	responseData := CCTPv2Messages{
		Messages: []CCTPv2Message{
			{
				Message:     "0xabcd",
				EventNonce:  "123",
				Attestation: "0xdef",
				Status:      "complete",
			},
		},
	}
	responseJSON, _ := json.Marshal(responseData)

	// Setup mocks
	mockHTTPClient.On("Get", ctx, mock.Anything).Return(
		cciptypes.Bytes(responseJSON),
		httputil.HTTPStatus(200),
		nil,
	)
	mockMetrics.On("TrackAttestationAPILatency",
		sourceChain, sourceDomainID, true, "200", mock.AnythingOfType("time.Duration")).Return()

	// Execute
	result, err := client.GetMessages(ctx, sourceChain, sourceDomainID, txHash)

	// Verify
	require.NoError(t, err)
	assert.Len(t, result.Messages, 1)
	assert.Equal(t, "0xabcd", result.Messages[0].Message)
	assert.Equal(t, "complete", result.Messages[0].Status)

	mockHTTPClient.AssertExpectations(t)
	mockMetrics.AssertExpectations(t)
}

func TestGetMessages_InvalidTransactionHash(t *testing.T) {
	lggr := logger.Test(t)
	mockHTTPClient := &MockHTTPClient{}
	mockMetrics := &MockMetricsReporter{}

	client := &CCTPv2HTTPClientImpl{
		lggr:            lggr,
		client:          mockHTTPClient,
		metricsReporter: mockMetrics,
	}

	ctx := context.Background()
	sourceChain := cciptypes.ChainSelector(1)
	sourceDomainID := uint32(0)

	testCases := []struct {
		name   string
		txHash string
		errMsg string
	}{
		{
			name:   "Empty transaction hash",
			txHash: "",
			errMsg: "transaction hash cannot be empty",
		},
		{
			name:   "Missing 0x prefix",
			txHash: "1234567890123456789012345678901234567890123456789012345678901234",
			errMsg: "invalid transaction hash format",
		},
		{
			name:   "Wrong length",
			txHash: "0x123456",
			errMsg: "invalid transaction hash format",
		},
		{
			name:   "Too long",
			txHash: "0x12345678901234567890123456789012345678901234567890123456789012345678",
			errMsg: "invalid transaction hash format",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Setup mock to expect error status tracking (validation errors have empty httpStatus)
			mockMetrics.On("TrackAttestationAPILatency",
				sourceChain, sourceDomainID, false, "", mock.AnythingOfType("time.Duration")).Return().Once()

			// Execute
			result, err := client.GetMessages(ctx, sourceChain, sourceDomainID, tc.txHash)

			// Verify
			assert.Error(t, err)
			assert.Contains(t, err.Error(), tc.errMsg)
			assert.Empty(t, result.Messages)

			mockMetrics.AssertExpectations(t)
		})
	}
}

func TestGetMessages_HTTPError(t *testing.T) {
	lggr := logger.Test(t)
	mockHTTPClient := &MockHTTPClient{}
	mockMetrics := &MockMetricsReporter{}

	client := &CCTPv2HTTPClientImpl{
		lggr:            lggr,
		client:          mockHTTPClient,
		metricsReporter: mockMetrics,
	}

	ctx := context.Background()
	sourceChain := cciptypes.ChainSelector(1)
	sourceDomainID := uint32(0)
	txHash := testTxHash

	// Setup mocks
	httpErr := errors.New("connection timeout")
	mockHTTPClient.On("Get", ctx, mock.Anything).Return(
		cciptypes.Bytes{},
		httputil.HTTPStatus(0),
		httpErr,
	)
	mockMetrics.On("TrackAttestationAPILatency",
		sourceChain, sourceDomainID, false, "0", mock.AnythingOfType("time.Duration")).Return()

	// Execute
	result, err := client.GetMessages(ctx, sourceChain, sourceDomainID, txHash)

	// Verify
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "http call failed")
	assert.Contains(t, err.Error(), "connection timeout")
	assert.Empty(t, result.Messages)

	mockHTTPClient.AssertExpectations(t)
	mockMetrics.AssertExpectations(t)
}

func TestGetMessages_NonOKStatus(t *testing.T) {
	lggr := logger.Test(t)
	mockHTTPClient := &MockHTTPClient{}
	mockMetrics := &MockMetricsReporter{}

	client := &CCTPv2HTTPClientImpl{
		lggr:            lggr,
		client:          mockHTTPClient,
		metricsReporter: mockMetrics,
	}

	ctx := context.Background()
	sourceChain := cciptypes.ChainSelector(1)
	sourceDomainID := uint32(0)
	txHash := testTxHash

	testCases := []struct {
		name       string
		statusCode httputil.HTTPStatus
		response   string
	}{
		{"Not Found", 404, `{"error":"not found"}`},
		{"Too Many Requests", 429, `{"error":"rate limited"}`},
		{"Internal Server Error", 500, `{"error":"server error"}`},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Setup mocks
			mockHTTPClient.On("Get", ctx, mock.Anything).Return(
				cciptypes.Bytes(tc.response),
				tc.statusCode,
				nil,
			).Once()
			expectedStatus := fmt.Sprintf("%d", tc.statusCode)
			mockMetrics.On("TrackAttestationAPILatency",
				sourceChain, sourceDomainID, false, expectedStatus, mock.AnythingOfType("time.Duration")).Return().Once()

			// Execute
			result, err := client.GetMessages(ctx, sourceChain, sourceDomainID, txHash)

			// Verify
			assert.Error(t, err)
			assert.Contains(t, err.Error(), "circle API returned status")
			assert.Empty(t, result.Messages)

			mockHTTPClient.AssertExpectations(t)
			mockMetrics.AssertExpectations(t)
		})
	}
}

func TestGetMessages_ParseError(t *testing.T) {
	lggr := logger.Test(t)
	mockHTTPClient := &MockHTTPClient{}
	mockMetrics := &MockMetricsReporter{}

	client := &CCTPv2HTTPClientImpl{
		lggr:            lggr,
		client:          mockHTTPClient,
		metricsReporter: mockMetrics,
	}

	ctx := context.Background()
	sourceChain := cciptypes.ChainSelector(1)
	sourceDomainID := uint32(0)
	txHash := testTxHash

	// Setup mocks with invalid JSON
	invalidJSON := cciptypes.Bytes(`{invalid json}`)
	mockHTTPClient.On("Get", ctx, mock.Anything).Return(
		invalidJSON,
		httputil.HTTPStatus(200),
		nil,
	)
	mockMetrics.On("TrackAttestationAPILatency",
		sourceChain, sourceDomainID, false, "200", mock.AnythingOfType("time.Duration")).Return()

	// Execute
	result, err := client.GetMessages(ctx, sourceChain, sourceDomainID, txHash)

	// Verify
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to decode json")
	assert.Empty(t, result.Messages)

	mockHTTPClient.AssertExpectations(t)
	mockMetrics.AssertExpectations(t)
}

func TestGetMessages_URLEncoding(t *testing.T) {
	lggr := logger.Test(t)
	mockHTTPClient := &MockHTTPClient{}
	mockMetrics := &MockMetricsReporter{}

	client := &CCTPv2HTTPClientImpl{
		lggr:            lggr,
		client:          mockHTTPClient,
		metricsReporter: mockMetrics,
	}

	ctx := context.Background()
	sourceChain := cciptypes.ChainSelector(1)
	sourceDomainID := uint32(0)
	txHash := testTxHash

	responseData := CCTPv2Messages{Messages: []CCTPv2Message{}}
	responseJSON, _ := json.Marshal(responseData)

	// Setup mock to capture the actual path
	var capturedPath string
	mockHTTPClient.On("Get", ctx, mock.MatchedBy(func(path string) bool {
		capturedPath = path
		return true
	})).Return(
		cciptypes.Bytes(responseJSON),
		httputil.HTTPStatus(200),
		nil,
	)
	mockMetrics.On("TrackAttestationAPILatency",
		sourceChain, sourceDomainID, true, "200", mock.AnythingOfType("time.Duration")).Return()

	// Execute
	_, err := client.GetMessages(ctx, sourceChain, sourceDomainID, txHash)

	// Verify
	require.NoError(t, err)
	assert.Contains(t, capturedPath, "v2/messages/0")
	assert.Contains(t, capturedPath, "transactionHash=0x")
	assert.Contains(t, capturedPath, txHash)

	mockHTTPClient.AssertExpectations(t)
	mockMetrics.AssertExpectations(t)
}

func TestGetMessages_MetricsTracking(t *testing.T) {
	lggr := logger.Test(t)
	ctx := context.Background()
	sourceChain := cciptypes.ChainSelector(1)
	sourceDomainID := uint32(0)
	txHash := testTxHash

	t.Run("Success path tracks success metrics", func(t *testing.T) {
		mockHTTPClient := &MockHTTPClient{}
		mockMetrics := &MockMetricsReporter{}

		client := &CCTPv2HTTPClientImpl{
			lggr:            lggr,
			client:          mockHTTPClient,
			metricsReporter: mockMetrics,
		}

		responseData := CCTPv2Messages{Messages: []CCTPv2Message{}}
		responseJSON, _ := json.Marshal(responseData)

		mockHTTPClient.On("Get", ctx, mock.Anything).Return(
			cciptypes.Bytes(responseJSON),
			httputil.HTTPStatus(200),
			nil,
		).Once()

		// Verify metrics called with success
		mockMetrics.On("TrackAttestationAPILatency",
			sourceChain, sourceDomainID, true, "200", mock.AnythingOfType("time.Duration")).Return().Once()

		_, err := client.GetMessages(ctx, sourceChain, sourceDomainID, txHash)
		require.NoError(t, err)

		mockHTTPClient.AssertExpectations(t)
		mockMetrics.AssertExpectations(t)
	})

	t.Run("Error path tracks error metrics", func(t *testing.T) {
		mockHTTPClient := &MockHTTPClient{}
		mockMetrics := &MockMetricsReporter{}

		client := &CCTPv2HTTPClientImpl{
			lggr:            lggr,
			client:          mockHTTPClient,
			metricsReporter: mockMetrics,
		}

		mockHTTPClient.On("Get", ctx, mock.Anything).Return(
			cciptypes.Bytes{},
			httputil.HTTPStatus(0),
			errors.New("http error"),
		).Once()

		// Verify metrics called with error
		mockMetrics.On("TrackAttestationAPILatency",
			sourceChain, sourceDomainID, false, "0", mock.AnythingOfType("time.Duration")).Return().Once()

		_, err := client.GetMessages(ctx, sourceChain, sourceDomainID, txHash)
		assert.Error(t, err)

		mockHTTPClient.AssertExpectations(t)
		mockMetrics.AssertExpectations(t)
	})

	t.Run("Metrics called exactly once per request", func(t *testing.T) {
		mockHTTPClient := &MockHTTPClient{}
		mockMetrics := &MockMetricsReporter{}

		client := &CCTPv2HTTPClientImpl{
			lggr:            lggr,
			client:          mockHTTPClient,
			metricsReporter: mockMetrics,
		}

		responseData := CCTPv2Messages{Messages: []CCTPv2Message{}}
		responseJSON, _ := json.Marshal(responseData)

		mockHTTPClient.On("Get", ctx, mock.Anything).Return(
			cciptypes.Bytes(responseJSON),
			httputil.HTTPStatus(200),
			nil,
		).Once()

		// Should be called exactly once
		mockMetrics.On("TrackAttestationAPILatency",
			sourceChain, sourceDomainID, true, "200", mock.AnythingOfType("time.Duration")).Return().Once()

		_, err := client.GetMessages(ctx, sourceChain, sourceDomainID, txHash)
		require.NoError(t, err)

		mockHTTPClient.AssertExpectations(t)
		mockMetrics.AssertExpectations(t)
		mockMetrics.AssertNumberOfCalls(t, "TrackAttestationAPILatency", 1)
	})
}

func TestParseResponseBody(t *testing.T) {
	t.Run("Valid: Empty messages array", func(t *testing.T) {
		jsonData := `{"messages": []}`
		result, err := parseResponseBody(cciptypes.Bytes(jsonData))
		require.NoError(t, err)
		assert.Empty(t, result.Messages)
	})

	t.Run("Valid: Single complete message", func(t *testing.T) {
		jsonData := `{
			"messages": [{
				"message": "0xabcdef1234567890",
				"eventNonce": "12345",
				"attestation": "0x9876543210fedcba",
				"cctpVersion": 2,
				"status": "complete",
				"decodedMessage": {
					"sourceDomain": "0",
					"destinationDomain": "1",
					"nonce": "100",
					"sender": "0x1111111111111111111111111111111111111111",
					"recipient": "0x2222222222222222222222222222222222222222",
					"destinationCaller": "0x3333333333333333333333333333333333333333",
					"messageBody": "0xdeadbeef",
					"decodedMessageBody": {
						"burnToken": "0x4444444444444444444444444444444444444444",
						"mintRecipient": "0x5555555555555555555555555555555555555555",
						"amount": "1000000",
						"messageSender": "0x6666666666666666666666666666666666666666"
					}
				}
			}]
		}`
		result, err := parseResponseBody(cciptypes.Bytes(jsonData))
		require.NoError(t, err)
		require.Len(t, result.Messages, 1)

		msg := result.Messages[0]
		assert.Equal(t, "0xabcdef1234567890", msg.Message)
		assert.Equal(t, "12345", msg.EventNonce)
		assert.Equal(t, "0x9876543210fedcba", msg.Attestation)
		assert.Equal(t, 2, msg.CCTPVersion)
		assert.Equal(t, "complete", msg.Status)
		assert.Equal(t, "0", msg.DecodedMessage.SourceDomain)
		assert.Equal(t, "1000000", msg.DecodedMessage.DecodedMessageBody.Amount)
	})

	t.Run("Valid: Message with optional fields", func(t *testing.T) {
		jsonData := `{
			"messages": [{
				"message": "0xtest",
				"eventNonce": "100",
				"attestation": "0xabc",
				"status": "complete",
				"cctpVersion": 2,
				"decodedMessage": {
					"sourceDomain": "0",
					"destinationDomain": "1",
					"nonce": "50",
					"sender": "0x1111111111111111111111111111111111111111",
					"recipient": "0x2222222222222222222222222222222222222222",
					"destinationCaller": "0x3333333333333333333333333333333333333333",
					"messageBody": "0xbody",
					"minFinalityThreshold": "65",
					"finalityThresholdExecuted": "128",
					"decodedMessageBody": {
						"burnToken": "0x4444444444444444444444444444444444444444",
						"mintRecipient": "0x5555555555555555555555555555555555555555",
						"amount": "2000000",
						"messageSender": "0x6666666666666666666666666666666666666666",
						"maxFee": "10000",
						"feeExecuted": "5000",
						"expirationBlock": "1000000",
						"hookData": "0xhookdata123"
					}
				}
			}]
		}`
		result, err := parseResponseBody(cciptypes.Bytes(jsonData))
		require.NoError(t, err)
		require.Len(t, result.Messages, 1)

		body := result.Messages[0].DecodedMessage.DecodedMessageBody
		assert.Equal(t, "10000", body.MaxFee)
		assert.Equal(t, "5000", body.FeeExecuted)
		assert.Equal(t, "1000000", body.ExpirationBlock)
		assert.Equal(t, "0xhookdata123", body.HookData)
	})

	t.Run("Edge: Empty strings for fields", func(t *testing.T) {
		jsonData := `{
			"messages": [{
				"message": "",
				"eventNonce": "",
				"attestation": "",
				"status": "",
				"cctpVersion": 0
			}]
		}`
		result, err := parseResponseBody(cciptypes.Bytes(jsonData))
		require.NoError(t, err)
		require.Len(t, result.Messages, 1)

		msg := result.Messages[0]
		assert.Equal(t, "", msg.Message)
		assert.Equal(t, "", msg.EventNonce)
		assert.Equal(t, "", msg.Attestation)
		assert.Equal(t, "", msg.Status)
	})

	t.Run("Edge: Very long hex string", func(t *testing.T) {
		// Create a 10KB+ hex string
		longHex := "0x" + strings.Repeat("ab", 5000)
		jsonData := fmt.Sprintf(`{
			"messages": [{
				"message": "%s",
				"eventNonce": "1",
				"attestation": "0xshort",
				"status": "complete",
				"cctpVersion": 2
			}]
		}`, longHex)
		result, err := parseResponseBody(cciptypes.Bytes(jsonData))
		require.NoError(t, err)
		require.Len(t, result.Messages, 1)
		assert.Equal(t, longHex, result.Messages[0].Message)
		assert.Len(t, result.Messages[0].Message, 10002) // "0x" + 10000 chars
	})

	t.Run("Edge: Extra unknown fields ignored", func(t *testing.T) {
		jsonData := `{
			"messages": [{
				"message": "0xtest",
				"eventNonce": "1",
				"attestation": "0xabc",
				"status": "complete",
				"cctpVersion": 2,
				"unknownField1": "should be ignored",
				"futureFeature": 999
			}],
			"extraTopLevel": "also ignored"
		}`
		result, err := parseResponseBody(cciptypes.Bytes(jsonData))
		require.NoError(t, err)
		require.Len(t, result.Messages, 1)
		assert.Equal(t, "0xtest", result.Messages[0].Message)
	})

	t.Run("Invalid: Empty JSON object", func(t *testing.T) {
		jsonData := `{}`
		result, err := parseResponseBody(cciptypes.Bytes(jsonData))
		require.NoError(t, err)
		assert.Nil(t, result.Messages) // Messages field not populated, so nil slice
	})

	t.Run("Invalid: Null messages field", func(t *testing.T) {
		jsonData := `{"messages": null}`
		result, err := parseResponseBody(cciptypes.Bytes(jsonData))
		require.NoError(t, err)
		assert.Nil(t, result.Messages)
	})

	t.Run("Invalid: Wrong type for messages (string)", func(t *testing.T) {
		jsonData := `{"messages": "not an array"}`
		result, err := parseResponseBody(cciptypes.Bytes(jsonData))
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed to decode json")
		assert.Empty(t, result.Messages)
	})

	t.Run("Invalid: Unclosed JSON", func(t *testing.T) {
		jsonData := `{"messages": [`
		result, err := parseResponseBody(cciptypes.Bytes(jsonData))
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed to decode json")
		assert.Empty(t, result.Messages)
	})

	t.Run("Invalid: Missing messages field", func(t *testing.T) {
		jsonData := `{"other": [], "data": "test"}`
		result, err := parseResponseBody(cciptypes.Bytes(jsonData))
		require.NoError(t, err)
		assert.Nil(t, result.Messages) // Field not present, so nil
	})
}
