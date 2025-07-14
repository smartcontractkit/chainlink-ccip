package v2

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"

	"github.com/smartcontractkit/chainlink-ccip/execute/exectypes"
	"github.com/smartcontractkit/chainlink-ccip/execute/tokendata/http"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
	"github.com/smartcontractkit/chainlink-ccip/pluginconfig"
)

const (
	apiVersionV2                     = "v2"
	messagesPath                     = "messages"
	attestationStatusComplete string = "complete"
)

type CCTPv2AttestationClient interface {
	GetMessages(ctx context.Context, sourceDomainID uint32, transactionHash string) (Messages, error)
}

type CCTPv2AttestationClientHTTP struct {
	lggr   logger.Logger
	client http.HTTPClient
}

func NewCCTPv2AttestationClientHTTP(
	lggr logger.Logger,
	config pluginconfig.USDCCCTPObserverConfig,
) (*CCTPv2AttestationClientHTTP, error) {
	client, err := http.GetHTTPClient(
		lggr,
		config.AttestationAPI,
		config.AttestationAPIInterval.Duration(),
		config.AttestationAPITimeout.Duration(),
		config.AttestationAPICooldown.Duration(),
	)
	if err != nil {
		return nil, fmt.Errorf("create HTTP client: %w", err)
	}
	return &CCTPv2AttestationClientHTTP{
		lggr:   lggr,
		client: client,
	}, nil
}

// GetMessages fetches CCTP v2 messages and their attestations from Circle's attestation API.
// It queries the API using the source domain ID and transaction hash to retrieve all
// CCTP v2 messages associated with the given transaction.
func (c *CCTPv2AttestationClientHTTP) GetMessages(
	ctx context.Context,
	sourceDomainID uint32,
	transactionHash string,
) (Messages, error) {
	path := fmt.Sprintf("%s/%s/%d?transactionHash=%s", apiVersionV2, messagesPath, sourceDomainID, transactionHash)
	body, status, err := c.client.Get(ctx, path)
	if err != nil {
		return Messages{},
			fmt.Errorf("http call failed to get CCTPv2 messages for sourceDomainID %d and transactionHash %s, error: %w",
				sourceDomainID, transactionHash, err)
	}

	if status != 200 {
		return Messages{}, fmt.Errorf(
			"http call failed to get CCTPv2 messages returned non-200 status: http status %d", status)
	}

	return parseResponseBody(body)
}

// parseResponseBody parses the JSON response from Circle's attestation API
// and returns a Messages struct containing the decoded CCTP v2 messages.
func parseResponseBody(body cciptypes.Bytes) (Messages, error) {
	var messages Messages
	if err := json.Unmarshal(body, &messages); err != nil {
		return Messages{}, fmt.Errorf("failed to decode json: %w", err)
	}
	return messages, nil
}

// Messages represents the response structure from Circle's attestation API,
// containing a list of CCTP v2 messages with their attestations.
type Messages struct {
	Messages []Message `json:"messages"`
}

// Message represents a single CCTP v2 message from Circle's attestation API.
// It contains the message data, attestation signature, and decoded message details
// needed for cross-chain USDC transfers.
type Message struct {
	Message        string         `json:"message"`
	EventNonce     string         `json:"eventNonce"`
	Attestation    string         `json:"attestation"`
	DecodedMessage DecodedMessage `json:"decodedMessage"`
	CCTPVersion    int            `json:"cctpVersion"`
	Status         string         `json:"status"`
}

// DecodedMessage represents the 'decodedMessage' object within a Message.
type DecodedMessage struct {
	SourceDomain       string             `json:"sourceDomain"`
	DestinationDomain  string             `json:"destinationDomain"`
	Nonce              string             `json:"nonce"`
	Sender             string             `json:"sender"`
	Recipient          string             `json:"recipient"`
	DestinationCaller  string             `json:"destinationCaller"`
	MessageBody        string             `json:"messageBody"`
	DecodedMessageBody DecodedMessageBody `json:"decodedMessageBody"`
	// The following fields are optional.
	MinFinalityThreshold      string `json:"minFinalityThreshold,omitempty"`
	FinalityThresholdExecuted string `json:"finalityThresholdExecuted,omitempty"`
}

// DecodedMessageBody represents the 'decodedMessageBody' object within a DecodedMessage.
type DecodedMessageBody struct {
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

// TokenData converts a CCTP v2 Message into TokenData for use in CCIP execution.
// It encodes the message bytes and attestation together using the provided encoder.
// The nonce and sourceDomainID can be used to uniquely identify and fetch this
// message from Circle's CCTP v2 API.
func (m *Message) TokenData(
	ctx context.Context,
	attestationEncoder AttestationEncoder,
) exectypes.TokenData {
	if m.Status != attestationStatusComplete {
		return exectypes.NewErrorTokenData(
			fmt.Errorf("A CCTPv2 Message's 'status' is not %s: "+
				"nonce: %s, sourceDomainID: %s, status: %s",
				attestationStatusComplete, m.EventNonce, m.DecodedMessage.SourceDomain, m.Status),
		)
	}

	messageBytes, err := cciptypes.NewBytesFromString(m.Message)
	if err != nil {
		return exectypes.NewErrorTokenData(
			fmt.Errorf("A CCTPv2 Message's 'message' field could not be converted from string to bytes: "+
				"nonce: %s, sourceDomainID: %s, error: %w",
				m.EventNonce, m.DecodedMessage.SourceDomain, err),
		)
	}

	attestationBytes, err := cciptypes.NewBytesFromString(m.Attestation)
	if err != nil {
		return exectypes.NewErrorTokenData(
			fmt.Errorf("A CCTPv2 Message's 'attestation' field could not be converted from string to bytes: "+
				"nonce: %s, sourceDomainID: %s, error: %w",
				m.EventNonce, m.DecodedMessage.SourceDomain, err),
		)
	}

	tokenDataBytes, err := attestationEncoder(ctx, messageBytes, attestationBytes)
	if err != nil {
		return exectypes.NewErrorTokenData(
			fmt.Errorf("attestationEncoder failed for a CCTPv2 message: "+
				"nonce: %s, sourceDomainID: %s, error: %w",
				m.EventNonce, m.DecodedMessage.SourceDomain, err),
		)
	}

	return exectypes.NewSuccessTokenData(tokenDataBytes)
}
