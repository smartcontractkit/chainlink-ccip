package lbtc

import (
	"context"
	"encoding/json"
	"fmt"
	"maps"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-ccip/execute/tokendata"
	"github.com/smartcontractkit/chainlink-ccip/execute/tokendata/http"
	"github.com/smartcontractkit/chainlink-ccip/pkg/reader"
	"github.com/smartcontractkit/chainlink-ccip/pluginconfig"
)

type AttestationStatus string

const (
	apiVersion      = "v1"
	attestationPath = "deposits/getByHash"

	attestationStatusUnspecified AttestationStatus = "NOTARIZATION_STATUS_UNSPECIFIED"
	attestationStatusFailed      AttestationStatus = "NOTARIZATION_STATUS_FAILED"
	attestationStatusPending     AttestationStatus = "NOTARIZATION_STATUS_PENDING"
	attestationStatusSubmitted   AttestationStatus = "NOTARIZATION_STATUS_SUBMITTED"
	_                            AttestationStatus = "NOTARIZATION_STATUS_SESSION_APPROVED"
)

type attestationRequest struct {
	PayloadHashes []string `json:"messageHash"`
}

type AttestationResponse struct {
	Attestations []Attestation `json:"attestations"`
	// fields in case of error
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

type Attestation struct {
	MessageHash string            `json:"message_hash"`
	Status      AttestationStatus `json:"status"`
	Data        string            `json:"attestation,omitempty"` // Data is represented by abi.encode(payload, proof)
}

type LBTCAttestationClient struct {
	lggr       logger.Logger
	config     pluginconfig.LBTCObserverConfig
	httpClient http.HTTPClient
}

func NewLBTCAttestationClient(
	lggr logger.Logger,
	config pluginconfig.LBTCObserverConfig,
) (tokendata.AttestationClient, error) {
	httpClient, err := http.GetHTTPClient(
		lggr,
		config.AttestationAPI,
		config.AttestationAPIInterval.Duration(),
		config.AttestationAPITimeout.Duration(),
		0, // no LBTC API cooldown
	)
	if err != nil {
		return nil, fmt.Errorf("get http client: %w", err)
	}
	return tokendata.NewObservedAttestationClient(
		lggr, &LBTCAttestationClient{
			lggr:       lggr,
			config:     config,
			httpClient: httpClient,
		},
	), nil
}

// Attestations is an AttestationClient method that accepts dict of messages and returns attestations under same keys.
// As values in input it accepts ExtraData bytes from incoming TokenData
func (c *LBTCAttestationClient) Attestations(
	ctx context.Context,
	messages map[cciptypes.ChainSelector]map[reader.MessageTokenID]cciptypes.Bytes,
) (map[cciptypes.ChainSelector]map[reader.MessageTokenID]tokendata.AttestationStatus, error) {
	attestations := make(map[string]tokendata.AttestationStatus)
	batch := make([]string, 0, c.config.AttestationAPIBatchSize)
	for _, tokenDatas := range messages {
		for _, tokenData := range tokenDatas {
			batch = append(batch, tokenData.String())
			if len(batch) == c.config.AttestationAPIBatchSize {
				batchAttestations, err := c.fetchBatch(ctx, batch)
				if err != nil {
					return nil, err
				}
				maps.Copy(attestations, batchAttestations)
				batch = make([]string, 0, c.config.AttestationAPIBatchSize)
			}
		}
	}
	if len(batch) > 0 {
		batchAttestations, err := c.fetchBatch(ctx, batch)
		if err != nil {
			return nil, err
		}
		maps.Copy(attestations, batchAttestations)
	}
	res := make(map[cciptypes.ChainSelector]map[reader.MessageTokenID]tokendata.AttestationStatus)
	for chainSelector, tokenDatas := range messages {
		res[chainSelector] = make(map[reader.MessageTokenID]tokendata.AttestationStatus)
		for messageTokenID, tokenData := range tokenDatas {
			if attestation, ok := attestations[tokenData.String()]; ok {
				res[chainSelector][messageTokenID] = attestation
			}
		}
	}
	return res, nil
}

func (c *LBTCAttestationClient) Type() string {
	return pluginconfig.LBTCHandlerType
}

// fetchBatch makes an HTTP Post request to Lombard Attestation API, requesting all messageHashes at once.
// This method doesn't check for input batch size.
func (c *LBTCAttestationClient) fetchBatch(
	ctx context.Context,
	messageHashes []string,
) (map[string]tokendata.AttestationStatus, error) {
	request := attestationRequest{PayloadHashes: messageHashes}
	encodedRequest, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal attestation request: %w", err)
	}
	attestations := make(map[string]tokendata.AttestationStatus)
	respRaw, _, err := c.httpClient.Post(ctx, fmt.Sprintf("bridge/%s/%s", apiVersion, attestationPath), encodedRequest)
	if err != nil {
		for _, inputMessageHash := range messageHashes {
			attestations[inputMessageHash] = tokendata.ErrorAttestationStatus(err)
		}
		// absorb api error to each token data status
		return attestations, nil
	}
	var attestationResp AttestationResponse
	err = json.Unmarshal(respRaw, &attestationResp)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal attestation response: %w", err)
	}
	if attestationResp.Code != 0 {
		for _, inputMessageHash := range messageHashes {
			attestations[inputMessageHash] = tokendata.ErrorAttestationStatus(
				fmt.Errorf("attestation request failed: %s", attestationResp.Message),
			)
		}
	}
	for _, attestation := range attestationResp.Attestations {
		attestations[attestation.MessageHash] = attestationToAttestationStatus(attestation)
	}
	for _, inputMessageHash := range messageHashes {
		if _, ok := attestations[inputMessageHash]; !ok {
			c.lggr.Warnw(
				"Requested messageHash is missing in the response. Considering tokendata.ErrDataMissing",
				"messageHash", inputMessageHash,
			)
		}
	}
	return attestations, nil
}

func attestationToAttestationStatus(attestation Attestation) tokendata.AttestationStatus {
	if attestation.Status == attestationStatusSubmitted || attestation.Status == attestationStatusPending {
		return tokendata.ErrorAttestationStatus(tokendata.ErrNotReady)
	}
	if attestation.Status == attestationStatusFailed || attestation.Status == attestationStatusUnspecified {
		return tokendata.ErrorAttestationStatus(tokendata.ErrUnknownResponse)
	}
	payloadHashBytes, err := cciptypes.NewBytesFromString(attestation.MessageHash)
	if err != nil {
		return tokendata.ErrorAttestationStatus(fmt.Errorf("failed to decode message hash in attestation: %w", err))
	}
	attestationBytes, err := cciptypes.NewBytesFromString(attestation.Data)
	if err != nil {
		return tokendata.ErrorAttestationStatus(fmt.Errorf("failed to decode attestation: %w", err))
	}
	return tokendata.SuccessAttestationStatus(payloadHashBytes, []byte{}, attestationBytes)
}
