package tokendata

import (
	"context"
	"errors"
	"time"

	"github.com/smartcontractkit/chainlink-ccip/pkg/reader"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

const (
	// cacheExpiration sets how long we keep the seen attestations in the cache for computing duration.
	// We can't keep them forever, so we don't track messages which are waiting for longer than 70 minutes
	cacheExpiration = 70 * time.Minute
	cacheCleanup    = 10 * time.Minute
)

var (
	ErrDataMissing     = errors.New("token data missing")
	ErrNotReady        = errors.New("token data not ready")
	ErrRateLimit       = errors.New("token data API is being rate limited")
	ErrTimeout         = errors.New("token data API timed out")
	ErrUnknownResponse = errors.New("unexpected response from attestation API")
)

type AttestationStatus struct {
	// ID is the identifier of the message that the attestation was fetched for
	ID cciptypes.Bytes
	// MessageBody is the raw message content
	MessageBody cciptypes.Bytes
	// Attestation is the attestation data fetched from the API, encoded in bytes
	Attestation cciptypes.Bytes
	// Error is the error that occurred during fetching the attestation data
	Error error
}

func SuccessAttestationStatus(id cciptypes.Bytes, body cciptypes.Bytes, attestation cciptypes.Bytes) AttestationStatus {
	return AttestationStatus{ID: id, MessageBody: body, Attestation: attestation}
}

func ErrorAttestationStatus(err error) AttestationStatus {
	return AttestationStatus{Error: err}
}

type AttestationClient interface {
	Attestations(
		ctx context.Context,
		msgs map[cciptypes.ChainSelector]map[reader.MessageTokenID]cciptypes.Bytes,
	) (map[cciptypes.ChainSelector]map[reader.MessageTokenID]AttestationStatus, error)

	Token() string
}

type FakeAttestationClient struct {
	Data map[string]AttestationStatus
}

func (f *FakeAttestationClient) Attestations(
	_ context.Context,
	messagesByChain map[cciptypes.ChainSelector]map[reader.MessageTokenID]cciptypes.Bytes,
) (map[cciptypes.ChainSelector]map[reader.MessageTokenID]AttestationStatus, error) {
	outcome := make(map[cciptypes.ChainSelector]map[reader.MessageTokenID]AttestationStatus)
	for chainSelector, messagesByTokenID := range messagesByChain {
		outcome[chainSelector] = make(map[reader.MessageTokenID]AttestationStatus)
		for tokenID, message := range messagesByTokenID {
			outcome[chainSelector][tokenID] = f.Data[string(message)]
		}
	}
	return outcome, nil
}

func (f *FakeAttestationClient) Token() string {
	return ""
}
