package usdc

import (
	"context"
	"errors"
	"time"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
)

var (
	ErrDataMissing = errors.New("token data missing")
	ErrNotReady    = errors.New("token data not ready")
	ErrRateLimit   = errors.New("token data API is being rate limited")
	ErrTimeout     = errors.New("token data API timed out")
)

type AttestationStatus struct {
	Error error
	Data  [32]byte
}

type AttestationClient interface {
	Attestations(
		ctx context.Context,
		msgs map[cciptypes.ChainSelector]map[cciptypes.SeqNum]map[int][]byte,
	) (map[cciptypes.ChainSelector]map[cciptypes.SeqNum]map[int]AttestationStatus, error)
}

type attestationStatus string

const (
	attestationStatusSuccess attestationStatus = "complete"
	attestationStatusPending attestationStatus = "pending_confirmations"
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
	msgs map[cciptypes.ChainSelector]map[cciptypes.SeqNum]map[int][]byte,
) (map[cciptypes.ChainSelector]map[cciptypes.SeqNum]map[int]AttestationStatus, error) {
	for _, chainMessages := range msgs {
		for _, messages := range chainMessages {
			for _, message := range messages {
				_, err := s.callAttestationAPI(ctx, message)
				if err != nil {
					return nil, err
				}
			}
		}
	}
	return nil, nil
}

func (s *SequentialAttestationClient) callAttestationAPI(
	_ context.Context,
	usdcMessageHash []byte,
) (attestationResponse, error) {
	if len(usdcMessageHash) == 0 {
		return attestationResponse{Status: attestationStatusPending}, ErrNotReady
	}
	return attestationResponse{Status: attestationStatusSuccess}, nil
}

type FakeAttestationClient struct {
	Data map[string]AttestationStatus
}

func (f FakeAttestationClient) Attestations(
	_ context.Context,
	msgs map[cciptypes.ChainSelector]map[cciptypes.SeqNum]map[int][]byte,
) (map[cciptypes.ChainSelector]map[cciptypes.SeqNum]map[int]AttestationStatus, error) {
	outcome := make(map[cciptypes.ChainSelector]map[cciptypes.SeqNum]map[int]AttestationStatus)

	for chainSelector, chainMessages := range msgs {
		outcome[chainSelector] = make(map[cciptypes.SeqNum]map[int]AttestationStatus)

		for seqNum, messages := range chainMessages {
			outcome[chainSelector][seqNum] = make(map[int]AttestationStatus)

			for index := range messages {
				outcome[chainSelector][seqNum][index] = f.Data[string(messages[index])]
			}
		}
	}
	return outcome, nil
}
