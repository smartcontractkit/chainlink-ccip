package v2

import (
	"context"
	"encoding/json"
	"errors"
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

// MockHTTPClient is a mock implementation of httputil.HTTPClient
type MockHTTPClient struct {
	mock.Mock
}

func (m *MockHTTPClient) Get(ctx context.Context, path string) (cciptypes.Bytes, httputil.HTTPStatus, error) {
	args := m.Called(ctx, path)
	return args.Get(0).(cciptypes.Bytes), args.Get(1).(httputil.HTTPStatus), args.Error(2)
}

func (m *MockHTTPClient) Post(ctx context.Context, path string, requestData cciptypes.Bytes) (cciptypes.Bytes, httputil.HTTPStatus, error) {
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
	status string,
	latency time.Duration,
) {
	m.Called(sourceChain, sourceDomain, status, latency)
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
	txHash := "0x1234567890123456789012345678901234567890123456789012345678901234"

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
		sourceChain, sourceDomainID, "success", mock.AnythingOfType("time.Duration")).Return()

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
			// Setup mock to expect error status tracking
			mockMetrics.On("TrackAttestationAPILatency",
				sourceChain, sourceDomainID, "error", mock.AnythingOfType("time.Duration")).Return().Once()

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
	txHash := "0x1234567890123456789012345678901234567890123456789012345678901234"

	// Setup mocks
	httpErr := errors.New("connection timeout")
	mockHTTPClient.On("Get", ctx, mock.Anything).Return(
		cciptypes.Bytes{},
		httputil.HTTPStatus(0),
		httpErr,
	)
	mockMetrics.On("TrackAttestationAPILatency",
		sourceChain, sourceDomainID, "error", mock.AnythingOfType("time.Duration")).Return()

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
	txHash := "0x1234567890123456789012345678901234567890123456789012345678901234"

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
			mockMetrics.On("TrackAttestationAPILatency",
				sourceChain, sourceDomainID, "error", mock.AnythingOfType("time.Duration")).Return().Once()

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
	txHash := "0x1234567890123456789012345678901234567890123456789012345678901234"

	// Setup mocks with invalid JSON
	invalidJSON := cciptypes.Bytes(`{invalid json}`)
	mockHTTPClient.On("Get", ctx, mock.Anything).Return(
		invalidJSON,
		httputil.HTTPStatus(200),
		nil,
	)
	mockMetrics.On("TrackAttestationAPILatency",
		sourceChain, sourceDomainID, "error", mock.AnythingOfType("time.Duration")).Return()

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
	txHash := "0x1234567890123456789012345678901234567890123456789012345678901234"

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
		sourceChain, sourceDomainID, "success", mock.AnythingOfType("time.Duration")).Return()

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
	txHash := "0x1234567890123456789012345678901234567890123456789012345678901234"

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

		// Verify metrics called with "success"
		mockMetrics.On("TrackAttestationAPILatency",
			sourceChain, sourceDomainID, "success", mock.AnythingOfType("time.Duration")).Return().Once()

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

		// Verify metrics called with "error"
		mockMetrics.On("TrackAttestationAPILatency",
			sourceChain, sourceDomainID, "error", mock.AnythingOfType("time.Duration")).Return().Once()

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
			sourceChain, sourceDomainID, "success", mock.AnythingOfType("time.Duration")).Return().Once()

		_, err := client.GetMessages(ctx, sourceChain, sourceDomainID, txHash)
		require.NoError(t, err)

		mockHTTPClient.AssertExpectations(t)
		mockMetrics.AssertExpectations(t)
		mockMetrics.AssertNumberOfCalls(t, "TrackAttestationAPILatency", 1)
	})
}
