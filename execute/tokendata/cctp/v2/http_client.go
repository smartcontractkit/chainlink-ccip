// Package v2 provides an HTTP client wrapper for Circle's CCTP v2 attestation API.
// This package handles the HTTP communication layer for fetching CCTP v2 messages
// and attestations from Circle's API endpoints.
// The CCTPv2 "get messages" API is documented here:
// https://developers.circle.com/api-reference/cctp/all/get-messages-v-2
package v2

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	httputil "github.com/smartcontractkit/chainlink-ccip/execute/tokendata/http"
	"github.com/smartcontractkit/chainlink-ccip/pluginconfig"
)

const (
	apiVersionV2 = "v2"
	messagesPath = "messages"
)

type CCTPv2HTTPClient interface {
	GetMessages(
		ctx context.Context, sourceChain cciptypes.ChainSelector, sourceDomainID uint32, transactionHash string,
	) (CCTPv2Messages, error)
}

// MetricsReporter provides metrics reporting for attestation API calls
type MetricsReporter interface {
	TrackAttestationAPILatency(
		sourceChain cciptypes.ChainSelector, sourceDomain uint32, success bool, httpStatus string, latency time.Duration)
}

// CCTPv2HTTPClientImpl implements CCTPv2AttestationClient using HTTP calls to Circle's attestation API
type CCTPv2HTTPClientImpl struct {
	lggr            logger.Logger
	client          httputil.HTTPClient
	metricsReporter MetricsReporter
}

// NewCCTPv2Client creates a new HTTP-based CCTP v2 attestation client
func NewCCTPv2Client(
	lggr logger.Logger,
	config pluginconfig.USDCCCTPObserverConfig,
	metricsReporter MetricsReporter,
) (*CCTPv2HTTPClientImpl, error) {
	if lggr == nil {
		return nil, fmt.Errorf("logger cannot be nil")
	}
	if metricsReporter == nil {
		return nil, fmt.Errorf("metricsReporter cannot be nil")
	}

	client, err := httputil.GetHTTPClient(
		lggr,
		config.AttestationAPI,
		config.AttestationAPIInterval.Duration(),
		config.AttestationAPITimeout.Duration(),
		config.AttestationAPICooldown.Duration(),
	)
	if err != nil {
		return nil, fmt.Errorf("create HTTP client: %w", err)
	}
	return &CCTPv2HTTPClientImpl{
		lggr:            lggr,
		client:          client,
		metricsReporter: metricsReporter,
	}, nil
}

// GetMessages fetches CCTP v2 messages and attestations for the given transaction.
func (c *CCTPv2HTTPClientImpl) GetMessages(
	ctx context.Context,
	sourceChain cciptypes.ChainSelector,
	sourceDomainID uint32,
	transactionHash string,
) (CCTPv2Messages, error) {
	startTime := time.Now()
	success := false
	httpStatus := ""

	defer func() {
		latency := time.Since(startTime)
		c.metricsReporter.TrackAttestationAPILatency(sourceChain, sourceDomainID, success, httpStatus, latency)
	}()

	// Validate transaction hash
	if transactionHash == "" {
		return CCTPv2Messages{}, fmt.Errorf("transaction hash cannot be empty")
	}
	if !strings.HasPrefix(transactionHash, "0x") || len(transactionHash) != 66 {
		return CCTPv2Messages{}, fmt.Errorf("invalid transaction hash format: %s", transactionHash)
	}

	path := fmt.Sprintf("%s/%s/%d?transactionHash=%s",
		apiVersionV2, messagesPath, sourceDomainID, url.QueryEscape(transactionHash))
	body, status, err := c.client.Get(ctx, path)
	httpStatus = strconv.Itoa(int(status))

	if err != nil {
		return CCTPv2Messages{},
			fmt.Errorf("http call failed to get CCTPv2 messages for sourceDomainID %d and transactionHash %s, error: %w",
				sourceDomainID, transactionHash, err)
	}

	if status != http.StatusOK {
		c.lggr.Warnw(
			"Non-200 status from Circle API",
			"status", status,
			"path", path,
			"sourceDomainID", sourceDomainID,
			"transactionHash", transactionHash,
			"responseBody", string(body),
		)
		return CCTPv2Messages{}, fmt.Errorf(
			"circle API returned status %d for path %s", status, path)
	}

	result, err := parseResponseBody(body)
	if err != nil {
		return CCTPv2Messages{}, err
	}

	success = true
	return result, nil
}

// parseResponseBody parses the JSON response from Circle's attestation API
// and returns a CCTPv2Messages struct containing the decoded CCTP v2 messages.
func parseResponseBody(body cciptypes.Bytes) (CCTPv2Messages, error) {
	var messages CCTPv2Messages
	if err := json.Unmarshal(body, &messages); err != nil {
		return CCTPv2Messages{}, fmt.Errorf("failed to decode json: %w", err)
	}
	return messages, nil
}

// CCTPv2Messages represents the response structure from Circle's attestation API,
// containing a list of CCTP v2 messages with their attestations.
// This API response type is documented here:
// https://developers.circle.com/api-reference/cctp/all/get-messages-v-2
type CCTPv2Messages struct {
	Messages []CCTPv2Message `json:"messages"`
}

// CCTPv2Message represents a single CCTP v2 message from Circle's attestation API.
// It contains the message data, attestation signature, and decoded message details
// needed for cross-chain USDC transfers.
type CCTPv2Message struct {
	Message        string               `json:"message"`
	EventNonce     string               `json:"eventNonce"`
	Attestation    string               `json:"attestation"`
	DecodedMessage CCTPv2DecodedMessage `json:"decodedMessage"`
	CCTPVersion    int                  `json:"cctpVersion"`
	Status         string               `json:"status"`
}

// CCTPv2DecodedMessage represents the 'decodedMessage' object within a CCTPv2Message.
type CCTPv2DecodedMessage struct {
	SourceDomain       string                   `json:"sourceDomain"`
	DestinationDomain  string                   `json:"destinationDomain"`
	Nonce              string                   `json:"nonce"`
	Sender             string                   `json:"sender"`
	Recipient          string                   `json:"recipient"`
	DestinationCaller  string                   `json:"destinationCaller"`
	MessageBody        string                   `json:"messageBody"`
	DecodedMessageBody CCTPv2DecodedMessageBody `json:"decodedMessageBody"`
	// The following fields are optional.
	MinFinalityThreshold      string `json:"minFinalityThreshold,omitempty"`
	FinalityThresholdExecuted string `json:"finalityThresholdExecuted,omitempty"`
}

// CCTPv2DecodedMessageBody represents the 'decodedMessageBody' object within a CCTPv2DecodedMessage.
type CCTPv2DecodedMessageBody struct {
	BurnToken     string `json:"burnToken"`
	MintRecipient string `json:"mintRecipient"`
	Amount        string `json:"amount"`
	MessageSender string `json:"messageSender"`
	// The following fields are optional.
	MaxFee          string `json:"maxFee,omitempty"`
	FeeExecuted     string `json:"feeExecuted,omitempty"`
	ExpirationBlock string `json:"expirationBlock,omitempty"`
	HookData        string `json:"hookData,omitempty"`
}
