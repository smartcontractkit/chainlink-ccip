package usdc

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"golang.org/x/exp/maps"

	"github.com/smartcontractkit/chainlink-ccip/execute/exectypes"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	//"github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-ccip/execute/tokendata"
	"github.com/smartcontractkit/chainlink-ccip/execute/tokendata/http"

	"github.com/smartcontractkit/chainlink-ccip/pkg/logutil"
	"github.com/smartcontractkit/chainlink-ccip/pkg/reader"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
	"github.com/smartcontractkit/chainlink-ccip/pluginconfig"
)

const (
	apiVersionV2 = "v2"
	messagesPath = "messages"
)

//// CCTPVersion represents the Solidity enum.
//// In Go, it's typically an int or uint, and you handle the mapping.
//type CCTPVersion uint8
//
//const (
//	CttpUnknownVersion CCTPVersion = iota
//	CttpVersion1
//	CttpVersion2
//)
//
//// SourceTokenDataPayload mirrors the Solidity struct.
//type SourceTokenDataPayload struct {
//	Nonce        uint64      `json:"nonce"`
//	SourceDomain uint32      `json:"sourceDomain"`
//	CCTPVersion  CCTPVersion `json:"cctpVersion"`
//}
//
//// TODO: doc
//func NewSourceTokenDataPayloadFromBytes(extraData cciptypes.Bytes) (*SourceTokenDataPayload, error) {
//	if len(extraData) < 96 {
//		return nil, fmt.Errorf("extraData is too short, expected at least 64 bytes")
//	}
//
//	// Extract the nonce (first 8 bytes), padded to 32 bytes
//	nonce := binary.BigEndian.Uint64(extraData[24:32])
//	// Extract the sourceDomain (next 4 bytes), padded to 32 bytes
//	sourceDomain := binary.BigEndian.Uint32(extraData[60:64])
//	// Extract the CCTP version (next 1 byte), padded to 32 bytes
//	cctpVersion := binary.BigEndian.Uint32(extraData[95:96])
//
//	return &SourceTokenDataPayload{
//		Nonce:        nonce,
//		SourceDomain: sourceDomain,
//		CCTPVersion:  CCTPVersion(cctpVersion),
//	}, nil
//}

// AttestationV2Client is a client for fetching attestation data from the V2 Circle API.
// TODO: doc
type AttestationV2Client struct {
	lggr                     logger.Logger
	client                   http.HTTPClient
	supportedPoolsBySelector map[cciptypes.ChainSelector]string
}

func NewAttestationV2Client(
	lggr logger.Logger,
	config pluginconfig.USDCCCTPObserverConfig,
	supportedPoolsBySelector map[cciptypes.ChainSelector]string,
) (*AttestationV2Client, error) {
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
	return &AttestationV2Client{
		lggr:                     lggr,
		client:                   client,
		supportedPoolsBySelector: supportedPoolsBySelector,
	}, nil
}

func (c *AttestationV2Client) Attestations(
	ctx context.Context,
	messages exectypes.MessageObservations,
) (map[cciptypes.ChainSelector]map[reader.MessageTokenID]tokendata.AttestationStatus, error) {
	lggr := logutil.WithContextValues(ctx, c.lggr)
	outcome := make(map[cciptypes.ChainSelector]map[reader.MessageTokenID]tokendata.AttestationStatus)

	for chainSelector, seqNum2Messages := range messages {
		outcome[chainSelector] = make(map[reader.MessageTokenID]tokendata.AttestationStatus)
		for seqNum, message := range seqNum2Messages {
			supportedPools := c.supportedPoolsBySelector[chainSelector]
			outcome[chainSelector] = c.getV2AttestationsForMessage(ctx, lggr, seqNum, message, supportedPools)
		}
	}

	return outcome, nil
}

// Needed?
//func (s *AttestationV2Client) Type() string {
//	return pluginconfig.USDCCCTPHandlerType
//}

// ResponseV2 represents the top-level structure of the JSON.
type ResponseV2 struct {
	Messages []MessageV2 `json:"messages"`
}

// MessageV2 represents an individual message entry.
type MessageV2 struct {
	Message        string         `json:"message"`
	EventNonce     string         `json:"eventNonce"`
	Attestation    string         `json:"attestation"`
	DecodedMessage DecodedMessage `json:"decodedMessage"`
	CCTPVersion    int            `json:"cctpVersion"`
	Status         string         `json:"status"`
}

// DecodedMessage represents the 'decodedMessage' object within a Message.
type DecodedMessage struct {
	SourceDomain       int                `json:"sourceDomain"`
	DestinationDomain  int                `json:"destinationDomain"`
	Nonce              string             `json:"nonce"`
	Sender             string             `json:"sender"`
	Recipient          string             `json:"recipient"`
	DestinationCaller  string             `json:"destinationCaller"`
	MessageBody        string             `json:"messageBody"`
	DecodedMessageBody DecodedMessageBody `json:"decodedMessageBody"`
	// The following fields are optional. We use pointers to indicate they might be null or absent in the JSON.
	MinFinalityThreshold      *int `json:"minFinalityThreshold,omitempty"`
	FinalityThresholdExecuted *int `json:"finalityThresholdExecuted,omitempty"`
}

// DecodedMessageBody represents the 'decodedMessageBody' object within a DecodedMessage.
type DecodedMessageBody struct {
	BurnToken     string `json:"burnToken"`
	MintRecipient string `json:"mintRecipient"`
	Amount        int    `json:"amount"`
	MessageSender string `json:"messageSender"`
	// The following fields are optional. We use pointers to indicate they might be null or absent in the JSON.
	MaxFee          *int    `json:"maxFee,omitempty"`
	FeeExecuted     *int    `json:"feeExecuted,omitempty"`
	ExpirationBlock *int    `json:"expirationBlock,omitempty"`
	HookData        *string `json:"hookData,omitempty"`
}

// TODO: doc
func (c *AttestationV2Client) getMessages(ctx context.Context, txHash string, sourceDomainId uint32) (*ResponseV2, error) {
	body, status, err := c.client.Get(ctx, fmt.Sprintf("%s/%s/%d?transactionHash=%s", apiVersionV2, messagesPath, sourceDomainId, txHash))
	if err != nil {
		return nil, fmt.Errorf("failed to get messages: %w", err)
	}

	if status != 200 {
		return nil, fmt.Errorf("failed to get messages: http status %d", status)
	}

	return parseResponseV2(body)
}

// TODO: doc
func parseResponseV2(body cciptypes.Bytes) (*ResponseV2, error) {
	var response ResponseV2
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("failed to decode json: %w", err)
	}
	return &response, nil
}

// TODO: doc
func (c *AttestationV2Client) associateAttestations(
	ctx context.Context,
	lggr logger.Logger,
	response *ResponseV2,
	message cciptypes.Message,
	tokenIdxs []int,
) map[reader.MessageTokenID]tokendata.AttestationStatus {
	return nil
}

// TODO: rename
func (c *AttestationV2Client) getV2AttestationsForMessage(
	ctx context.Context,
	lggr logger.Logger,
	seqNr cciptypes.SeqNum,
	message cciptypes.Message,
	supportedPools string,
) map[reader.MessageTokenID]tokendata.AttestationStatus {
	result := make(map[reader.MessageTokenID]tokendata.AttestationStatus)

	// Maps source domain IDs to a list of token indices.
	sourceDomainIds := make(map[uint32][]int)

	for idx, ta := range message.TokenAmounts {
		// TODO: doc
		if !strings.EqualFold(supportedPools, ta.SourcePoolAddress.String()) {
			// debug log
			continue
		}

		sourceTokenDataPayload, err := reader.NewSourceTokenDataPayloadFromBytes(ta.ExtraData)
		if err != nil {
			// debug log
			continue
		}

		if sourceTokenDataPayload.CCTPVersion != reader.CttpVersion2 {
			// debug log
			continue
		}

		sourceDomainIds[sourceTokenDataPayload.SourceDomain] =
			append(sourceDomainIds[sourceTokenDataPayload.SourceDomain], idx)
	}

	for sourceDomainId, tokenIndices := range sourceDomainIds {
		response, err := c.getMessages(ctx, message.Header.TxHash, sourceDomainId)
		if err != nil {
			for _, idx := range tokenIndices {
				result[reader.NewMessageTokenID(seqNr, idx)] = tokendata.ErrorAttestationStatus(err)
			}
		}

		attestations := c.associateAttestations(ctx, lggr, response, message, tokenIndices)
		maps.Copy(result, attestations)
	}

	return result
}
