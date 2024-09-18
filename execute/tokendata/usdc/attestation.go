package usdc

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-ccip/execute/exectypes"
)

var (
	ErrNotReady        = errors.New("token data not ready")
	ErrRateLimit       = errors.New("token data API is being rate limited")
	ErrTimeout         = errors.New("token data API timed out")
	ErrRequestsBlocked = errors.New("requests are currently blocked")
)

type AttestationClient interface {
	Attestations(
		ctx context.Context,
		msgs map[cciptypes.ChainSelector]map[cciptypes.SeqNum][]usdcTokenData,
	) (map[cciptypes.ChainSelector]map[cciptypes.SeqNum]exectypes.TokenData, error)
}

type attestationStatus string

//const (
//	attestationStatusSuccess attestationStatus = "complete"
//	attestationStatusPending attestationStatus = "pending_confirmations"
//)

const (
	apiVersion      = "v1"
	attestationPath = "attestations"
	//defaultAttestationTimeout = 5 * time.Second

	// defaultCoolDownDurationSec defines the default time to wait after getting rate limited.
	// this value is only used if the 429 response does not contain the Retry-After header
	defaultCoolDownDuration = 5 * time.Minute

	// maxCoolDownDuration defines the maximum duration we can wait till firing the next request
	maxCoolDownDuration = 10 * time.Minute

	// defaultRequestInterval defines the rate in requests per second that the attestation API can be called.
	// this is set according to the APIs documentated 10 requests per second rate limit.
	//defaultRequestInterval = 100 * time.Millisecond

	// APIIntervalRateLimitDisabled is a special value to disable the rate limiting.
	//APIIntervalRateLimitDisabled = -1
	// APIIntervalRateLimitDefault is a special value to select the default rate limit interval.
	//APIIntervalRateLimitDefault = 0
)

type attestationResponse struct {
	Status      attestationStatus `json:"status"`
	Attestation string            `json:"attestation"`
	Error       string            `json:"error"`
}

type SequentialAttestationClient struct {
	httpClient            HTTPClient
	attestationAPI        string
	attestationAPITimeout time.Duration
	coolDownPeriod        time.Duration
}

func NewSequentialAttestationClient(
	attestationAPI string,
	attestationAPITimeout time.Duration,
) *SequentialAttestationClient {
	return &SequentialAttestationClient{
		httpClient:            HTTPClient{},
		attestationAPI:        attestationAPI,
		attestationAPITimeout: attestationAPITimeout,
		coolDownPeriod:        0,
	}
}

func (s *SequentialAttestationClient) Attestations(
	ctx context.Context,
	msgs map[cciptypes.ChainSelector]map[cciptypes.SeqNum][]usdcTokenData,
) (map[cciptypes.ChainSelector]map[cciptypes.SeqNum]exectypes.TokenData, error) {
	for _, chainMessages := range msgs {
		for _, messages := range chainMessages {
			for _, message := range messages {
				_, err := s.callAttestationAPI(ctx, message.MessageHash)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	return nil, nil
}

// callAttestationAPI calls the USDC attestation API with the given USDC message hash.
// The attestation service rate limit is 10 requests per second. If you exceed 10 requests
// per second, the service blocks all API requests for the next 5 minutes and returns an
// HTTP 429 response.
//
// Documentation:
//
//	https://developers.circle.com/stablecoins/reference/getattestation
//	https://developers.circle.com/stablecoins/docs/transfer-usdc-on-testnet-from-ethereum-to-avalanche
func (s *SequentialAttestationClient) callAttestationAPI(
	ctx context.Context,
	usdcMessageHash [32]byte,
) (attestationResponse, error) {
	body, _, headers, err := s.httpClient.Get(
		ctx,
		fmt.Sprintf("%s/%s/%s/0x%x", s.attestationAPI, apiVersion, attestationPath, usdcMessageHash),
		s.attestationAPITimeout,
	)
	switch {
	case errors.Is(err, ErrRateLimit):
		coolDownDuration := defaultCoolDownDuration
		if retryAfterHeader, exists := headers["Retry-After"]; exists && len(retryAfterHeader) > 0 {
			if retryAfterSec, errParseInt := strconv.ParseInt(retryAfterHeader[0], 10, 64); errParseInt == nil {
				coolDownDuration = time.Duration(retryAfterSec) * time.Second
			}
		}
		s.setCoolDownPeriod(coolDownDuration)
		// Explicitly signal if the API is being rate limited
		return attestationResponse{}, ErrRateLimit
	case err != nil:
		return attestationResponse{}, fmt.Errorf("request error: %w", err)
	}

	var response attestationResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return attestationResponse{}, err
	}
	if response.Error != "" {
		return attestationResponse{}, fmt.Errorf("attestation API error: %s", response.Error)
	}
	if response.Status == "" {
		return attestationResponse{}, fmt.Errorf("invalid attestation response: %s", string(body))
	}
	return response, nil
}

func (s *SequentialAttestationClient) setCoolDownPeriod(duration time.Duration) {
	if duration > maxCoolDownDuration {
		duration = maxCoolDownDuration
	}
	s.coolDownPeriod = duration
}
