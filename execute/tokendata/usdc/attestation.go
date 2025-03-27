package usdc

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/smartcontractkit/chainlink-common/pkg/hashutil"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"

	"github.com/smartcontractkit/chainlink-ccip/execute/tokendata"
	"github.com/smartcontractkit/chainlink-ccip/execute/tokendata/http"

	"github.com/smartcontractkit/chainlink-ccip/pkg/logutil"
	"github.com/smartcontractkit/chainlink-ccip/pkg/reader"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
	"github.com/smartcontractkit/chainlink-ccip/pluginconfig"
)

type attestationStatus string

const (
	apiVersion      = "v1"
	attestationPath = "attestations"

	attestationStatusSuccess attestationStatus = "complete"
	attestationStatusPending attestationStatus = "pending_confirmations"
)

type httpResponse struct {
	Status      attestationStatus `json:"status"`
	Attestation string            `json:"attestation"`
	Error       string            `json:"error"`
}

func (r httpResponse) validate() error {
	if r.Error != "" {
		return fmt.Errorf("attestation API error: %s", r.Error)
	}
	if r.Status == "" {
		return fmt.Errorf("invalid attestation response")
	}
	if r.Status != attestationStatusSuccess {
		return tokendata.ErrNotReady
	}
	return nil
}

func (r httpResponse) attestationToBytes() (cciptypes.Bytes, error) {
	attestationBytes, err := cciptypes.NewBytesFromString(r.Attestation)
	if err != nil {
		return nil, fmt.Errorf("failed to decode attestation hex: %w", err)
	}
	return attestationBytes, nil
}

type AttestationEncoder func(context.Context, cciptypes.Bytes, cciptypes.Bytes) (cciptypes.Bytes, error)

// USDCAttestationClient is an client for fetching attestation data from the Circle API.
// It returns a data grouped by chainSelector, sequenceNumber and tokenIndex
//
// Example: if we have two USDC tokens transferred (slot 0 and slot 2) within a single message with sequence number 12
// on Ethereum chain with, it's going to look like this:
// Ethereum ->
//
//	12 ->
//	  0 -> AttestationStatus{Error: nil, Attestation: "ABCDEF", MessageHash: bytes}
//	  2 -> AttestationStatus{Error: ErrNotRead, Attestation: nil, MessageHash: nil}
type USDCAttestationClient struct {
	lggr   logger.Logger
	client http.HTTPClient
	hasher hashutil.Hasher[[32]byte]
}

func NewSequentialAttestationClient(
	lggr logger.Logger,
	config pluginconfig.USDCCCTPObserverConfig,
) (tokendata.AttestationClient, error) {
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
	return &USDCAttestationClient{
		lggr:   lggr,
		client: client,
		hasher: hashutil.NewKeccak(),
	}, nil
}

func (s *USDCAttestationClient) Attestations(
	ctx context.Context,
	messagesByChain map[cciptypes.ChainSelector]map[reader.MessageTokenID]cciptypes.Bytes,
) (map[cciptypes.ChainSelector]map[reader.MessageTokenID]tokendata.AttestationStatus, error) {
	lggr := logutil.WithContextValues(ctx, s.lggr)
	outcome := make(map[cciptypes.ChainSelector]map[reader.MessageTokenID]tokendata.AttestationStatus)

	for chainSelector, messagesByTokenID := range messagesByChain {
		outcome[chainSelector] = make(map[reader.MessageTokenID]tokendata.AttestationStatus)

		for tokenID, message := range messagesByTokenID {
			lggr.Debugw(
				"Fetching attestation from the API",
				"chainSelector", chainSelector,
				"message", message,
				"messageTokenID", tokenID,
			)
			outcome[chainSelector][tokenID] = s.fetchSingleMessage(ctx, message)
		}
	}
	return outcome, nil
}

func (s *USDCAttestationClient) Token() string {
	return USDCToken
}

func (s *USDCAttestationClient) fetchSingleMessage(
	ctx context.Context,
	message cciptypes.Bytes,
) tokendata.AttestationStatus {
	messageHash := cciptypes.Bytes32(s.hasher.Hash(message))
	body, _, err := s.client.Get(ctx, fmt.Sprintf("%s/%s/%s", apiVersion, attestationPath, messageHash.String()))
	if err != nil {
		return tokendata.ErrorAttestationStatus(err)
	}
	response, err := attestationFromResponse(body)
	if err != nil {
		return tokendata.ErrorAttestationStatus(err)
	}
	return tokendata.SuccessAttestationStatus(messageHash[:], message, response)
}

func attestationFromResponse(body cciptypes.Bytes) (cciptypes.Bytes, error) {
	var response httpResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("failed to decode json: %w", err)
	}

	if err := response.validate(); err != nil {
		return nil, err
	}

	return response.attestationToBytes()
}
