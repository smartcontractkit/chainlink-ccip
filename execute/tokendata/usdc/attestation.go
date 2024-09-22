package usdc

import (
	"context"
	"errors"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-ccip/execute/exectypes"
	"github.com/smartcontractkit/chainlink-ccip/pkg/reader"
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

// AttestationClient is an interface for fetching attestation data from the Circle API.
// It returns a data grouped by chainSelector, sequenceNumber and tokenIndex
//
// Example: if we have two USDC tokens transferred (slot 0 and slot 2) within a single message with sequence number 12
// on Ethereum chain with, it's going to look like this:
// Ethereum ->
//
//	12 ->
//	  0 -> AttestationStatus{Error: nil, Data: [32]byte{attestation_data}}
//	  2 -> AttestationStatus{Error: ErrNotRead, Data: nil}
type AttestationClient interface {
	Attestations(
		ctx context.Context,
		msgs map[cciptypes.ChainSelector]map[exectypes.MessageTokenID]reader.MessageHash,
	) (map[cciptypes.ChainSelector]map[exectypes.MessageTokenID]AttestationStatus, error)
}

type FakeAttestationClient struct {
	Data map[string]AttestationStatus
}

func (f FakeAttestationClient) Attestations(
	_ context.Context,
	msgs map[cciptypes.ChainSelector]map[exectypes.MessageTokenID]reader.MessageHash,
) (map[cciptypes.ChainSelector]map[exectypes.MessageTokenID]AttestationStatus, error) {
	outcome := make(map[cciptypes.ChainSelector]map[exectypes.MessageTokenID]AttestationStatus)

	for chainSelector, hashes := range msgs {
		outcome[chainSelector] = make(map[exectypes.MessageTokenID]AttestationStatus)

		for tokenID, messageHash := range hashes {
			outcome[chainSelector][tokenID] = f.Data[string(messageHash)]
		}
	}
	return outcome, nil
}
