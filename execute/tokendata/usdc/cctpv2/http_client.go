// Exposes an HTTP client to call the Circle Cross-Chain Transfer Protocol (CCTP) V2 API.
//
// The goal is to fetch attestations for USDC transfers made via CCTPv2 by CCIP. The /v2/messages/ endpoint
// takes a source domain ID and transaction hash and returns a list of CCTPv2 messages executed on that source and
// transaction. Each message contains an attestation and a decoded message body, which are needed to pass to
// destination chain contracts where the attestation will be verified and the USDC transfer will be executed (e.g. the
// USDC will be minted and transferred on the destination chain).
//
// API docs are here: https://developers.circle.com/api-reference/stablecoins/common/get-messages-v-2
//
// An example API calls is:
//nolint:lll
// https://iris-api-sandbox.circle.com/v2/messages/0?transactionHash=0x8f1d58a2161b0bd662609b62537d57b76cfa6bb22a9afe8ec49b09f3810d44c2

package cctpv2

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/smartcontractkit/chainlink-ccip/execute/exectypes"
	"github.com/smartcontractkit/chainlink-ccip/execute/tokendata/http"
	"github.com/smartcontractkit/chainlink-ccip/execute/tokendata/usdc"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
	"github.com/smartcontractkit/chainlink-ccip/pluginconfig"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
)

const (
	apiVersionV2              = "v2"
	messagesPath              = "messages"
	attestationStatusComplete = "complete"
)

type CTTPv2AttestationClient struct {
	lggr   logger.Logger
	client http.HTTPClient
}

func NewCTTPv2AttestationClient(
	lggr logger.Logger,
	config pluginconfig.USDCCCTPObserverConfig,
) (*CTTPv2AttestationClient, error) {
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
	return &CTTPv2AttestationClient{
		lggr:   lggr,
		client: client,
	}, nil
}

// GetMessages Calls the CCTPv2 API to get messages for a given source domain ID and transaction hash.
func (c *CTTPv2AttestationClient) GetMessages(
	ctx context.Context,
	sourceDomainID uint32,
	transactionHash string,
) (Messages, error) {
	path := fmt.Sprintf("%s/%s/%d?transactionHash=%s", apiVersionV2, messagesPath, sourceDomainID, transactionHash)
	start := time.Now()
	body, status, err := c.client.Get(ctx, path)
	duration := time.Since(start)
	TrackHTTPRequest(sourceDomainID, status, err != nil, duration)
	if err != nil {
		return Messages{},
			fmt.Errorf(
				"http call failed to get CCTPv2 messages for sourceDomainID %d and transactionHash %s, error: %w",
				sourceDomainID, transactionHash, err,
			)
	}

	if status != 200 {
		return Messages{},
			fmt.Errorf(
				"http call for CCTPv2 messages returned non-200 status %d for sourceDomainID %d and transactionHash %s",
				status, sourceDomainID, transactionHash,
			)
	}

	return parseResponseBody(body)
}

// parses the response body from the CCTPv2 API into a Messages struct
func parseResponseBody(body []byte) (Messages, error) {
	var messages Messages
	if err := json.Unmarshal(body, &messages); err != nil {
		return Messages{}, fmt.Errorf("failed to decode json: %w", err)
	}
	return messages, nil
}

// Messages The struct returned by the CCTPv2 API when fetching messages.
type Messages struct {
	Messages []Message `json:"messages"`
}

// Message contains all the info for a single CCTPv2 transfer
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

// TokenData converts a CCTPv2 Message into TokenData
func (m *Message) TokenData(
	ctx context.Context,
	attestationEncoder usdc.AttestationEncoder,
) exectypes.TokenData {
	if m.Status != attestationStatusComplete {
		return exectypes.NewErrorTokenData(
			fmt.Errorf("expected CCTPv2 Message's 'status' to be %s but got %s: "+
				"nonce: %s, sourceDomainId: %s",
				attestationStatusComplete, m.Status, m.EventNonce, m.DecodedMessage.SourceDomain),
		)
	}

	messageBytes, err := cciptypes.NewBytesFromString(m.Message)
	if err != nil {
		return exectypes.NewErrorTokenData(
			fmt.Errorf("A CCTPv2 Message's 'message' field could not be converted from string to bytes: "+
				"nonce: %s, sourceDomainId: %s, error: %w",
				m.EventNonce, m.DecodedMessage.SourceDomain, err),
		)
	}

	attestationBytes, err := cciptypes.NewBytesFromString(m.Attestation)
	if err != nil {
		return exectypes.NewErrorTokenData(
			fmt.Errorf("A CCTPv2 Message's 'attestation' field could not be converted from string to bytes: "+
				"nonce: %s, sourceDomainId: %s, error: %w",
				m.EventNonce, m.DecodedMessage.SourceDomain, err),
		)
	}

	tokenDataBytes, err := attestationEncoder(ctx, messageBytes, attestationBytes)
	if err != nil {
		return exectypes.NewErrorTokenData(
			fmt.Errorf("attestationEncoder failed for a CCTPv2 message: "+
				"nonce: %s, sourceDomainId: %s, error: %w",
				m.EventNonce, m.DecodedMessage.SourceDomain, err),
		)
	}

	return exectypes.NewSuccessTokenData(tokenDataBytes)
}
