package tokendata

import (
	"errors"

	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
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
	// Attestation is the attestation data fetched from the API, encoded in bytes
	Attestation cciptypes.Bytes
	// Error is the error that occurred during fetching the attestation data
	Error error
}

func SuccessAttestationStatus(id cciptypes.Bytes, attestation cciptypes.Bytes) AttestationStatus {
	return AttestationStatus{ID: id, Attestation: attestation}
}

func ErrorAttestationStatus(err error) AttestationStatus {
	return AttestationStatus{Error: err}
}
