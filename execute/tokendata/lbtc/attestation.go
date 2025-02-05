package lbtc

import (
	"context"
	"encoding/json"
	"fmt"
	"maps"
	"slices"
	"time"

	"github.com/smartcontractkit/chainlink-ccip/execute/tokendata"
	"github.com/smartcontractkit/chainlink-ccip/execute/tokendata/http"
	"github.com/smartcontractkit/chainlink-ccip/pkg/reader"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
	"github.com/smartcontractkit/chainlink-ccip/pluginconfig"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
)

type attestationStatus string

const (
	apiVersion              = "v1"
	attestationPath         = "deposits/getByHash"
	defaultCoolDownDuration = 30 * time.Second

	attestationStatusUnspecified attestationStatus = "NOTARIZATION_STATUS_UNSPECIFIED"
	attestationStatusFailed      attestationStatus = "NOTARIZATION_STATUS_FAILED"
	attestationStatusPending     attestationStatus = "NOTARIZATION_STATUS_PENDING"
	attestationStatusSubmitted   attestationStatus = "NOTARIZATION_STATUS_SUBMITTED"
	_                            attestationStatus = "NOTARIZATION_STATUS_SESSION_APPROVED"
)

type LBTCAttestationClient struct {
	config     pluginconfig.LBTCObserverConfig
	httpClient http.HTTPClient
}

func NewLBTCAttestationClient(lggr logger.Logger, config pluginconfig.LBTCObserverConfig) (*LBTCAttestationClient, error) {
	httpClient, err := http.GetHTTPClient(
		lggr,
		http.LBTC,
		fmt.Sprintf("%s/bridge/%s/%s", config.AttestationAPI, apiVersion, attestationPath),
		config.AttestationAPIInterval.Duration(),
		config.AttestationAPITimeout.Duration(),
		defaultCoolDownDuration,
	)
	if err != nil {
		return nil, fmt.Errorf("get http client: %w", err)
	}
	return &LBTCAttestationClient{
		config:     config,
		httpClient: httpClient,
	}, nil
}

type messageAttestationResponse struct {
	MessageHash string            `json:"message_hash"`
	Status      attestationStatus `json:"status"`
	Attestation string            `json:"attestation,omitempty"` // Attestation represented by abi.encode(payload, proof)
}

type attestationRequest struct {
	PayloadHashes []string `json:"messageHash"`
}

type attestationResponse struct {
	Attestations []messageAttestationResponse `json:"attestations"`
}

func (c *LBTCAttestationClient) fetchBatch(ctx context.Context, batch map[string]reader.MessageTokenID) (map[reader.MessageTokenID]tokendata.AttestationStatus, error) {
	request := attestationRequest{PayloadHashes: slices.Collect(maps.Keys(batch))}
	encodedRequest, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal attestation request: %w", err)
	}
	respRaw, _, err := c.httpClient.Post(ctx, encodedRequest)
	if err != nil {
		return nil, err
	}
	var attestationResp attestationResponse
	err = json.Unmarshal(respRaw, &attestationResp)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal attestation response: %w", err)
	}
	attestations := make(map[reader.MessageTokenID]tokendata.AttestationStatus)
	for _, attestation := range attestationResp.Attestations {
		messageID := batch[attestation.MessageHash]
		attestations[messageID] = attestationToTokenData(attestation)
	}
	return attestations, nil
}

func attestationToTokenData(attestation messageAttestationResponse) tokendata.AttestationStatus {
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
	attestationBytes, err := cciptypes.NewBytesFromString(attestation.Attestation)
	if err != nil {
		return tokendata.ErrorAttestationStatus(fmt.Errorf("failed to decode attestation: %w", err))
	}
	return tokendata.SuccessAttestationStatus(payloadHashBytes, attestationBytes)
}
