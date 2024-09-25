package usdc

import (
	"context"
	"errors"

	"github.com/smartcontractkit/chainlink-common/pkg/hashutil"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-ccip/execute/exectypes"
	"github.com/smartcontractkit/chainlink-ccip/pkg/reader"
	"github.com/smartcontractkit/chainlink-ccip/pluginconfig"
)

var (
	ErrDataMissing     = errors.New("token data missing")
	ErrNotReady        = errors.New("token data not ready")
	ErrRateLimit       = errors.New("token data API is being rate limited")
	ErrTimeout         = errors.New("token data API timed out")
	ErrUnknownResponse = errors.New("unexpected response from attestation API")
)

type AttestationStatus struct {
	MessageHash []byte
	Attestation []byte
	Error       error
}

func SuccessAttestationStatus(messageHash []byte, attestation []byte) AttestationStatus {
	return AttestationStatus{MessageHash: messageHash, Attestation: attestation}
}

func ErrorAttestationStatus(err error) AttestationStatus {
	return AttestationStatus{Error: err}
}

// AttestationClient is an interface for fetching attestation data from the Circle API.
// It returns a data grouped by chainSelector, sequenceNumber and tokenIndex
//
// Example: if we have two USDC tokens transferred (slot 0 and slot 2) within a single message with sequence number 12
// on Ethereum chain with, it's going to look like this:
// Ethereum ->
//
//	12 ->
//	  0 -> AttestationStatus{Error: nil, Attestation: "ABCDEF", MessageHash: bytes}
//	  2 -> AttestationStatus{Error: ErrNotRead, Attestation: nil, MessageHash: nil}
type AttestationClient interface {
	Attestations(
		ctx context.Context,
		msgs map[cciptypes.ChainSelector]map[exectypes.MessageTokenID]reader.MessageHash,
	) (map[cciptypes.ChainSelector]map[exectypes.MessageTokenID]AttestationStatus, error)
}

type sequentialAttestationClient struct {
	client HTTPClient
	hasher hashutil.Hasher[[32]byte]
}

//nolint:revive
func NewAttestationClient(config pluginconfig.USDCCCTPObserverConfig) *sequentialAttestationClient {
	return &sequentialAttestationClient{
		client: NewHTTPClient(
			config.AttestationAPI,
			config.AttestationAPIInterval.Duration(),
			config.AttestationAPITimeout.Duration(),
		),
		hasher: hashutil.NewKeccak(),
	}
}

func (s *sequentialAttestationClient) Attestations(
	ctx context.Context,
	msgs map[cciptypes.ChainSelector]map[exectypes.MessageTokenID]reader.MessageHash,
) (map[cciptypes.ChainSelector]map[exectypes.MessageTokenID]AttestationStatus, error) {
	outcome := make(map[cciptypes.ChainSelector]map[exectypes.MessageTokenID]AttestationStatus)

	for chainSelector, hashes := range msgs {
		outcome[chainSelector] = make(map[exectypes.MessageTokenID]AttestationStatus)

		for tokenID, messageHash := range hashes {
			// TODO sequential processing
			outcome[chainSelector][tokenID] = s.fetchSingleMessage(ctx, messageHash)
		}
	}
	return outcome, nil
}

func (s *sequentialAttestationClient) fetchSingleMessage(ctx context.Context, messageHash []byte) AttestationStatus {
	response, _, err := s.client.Get(ctx, s.hasher.Hash(messageHash))
	if err != nil {
		return ErrorAttestationStatus(err)
	}

	return SuccessAttestationStatus(messageHash, response)
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
