package usdc

import (
	"context"
	"errors"
	"fmt"

	"github.com/smartcontractkit/chainlink-ccip/pkg/reader"
	"github.com/smartcontractkit/chainlink-common/pkg/hashutil"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-ccip/pluginconfig"
)

var (
	ErrDataMissing     = errors.New("token data missing")
	ErrNotReady        = errors.New("token data not ready")
	ErrRateLimit       = errors.New("token data API is being rate limited")
	ErrTimeout         = errors.New("token data API timed out")
	ErrUnknownResponse = errors.New("unexpected response from attestation API")
)

// AttestationStatus is a struct holding all the necessary information to build payload to
// mint USDC on the destination chain. Valid AttestationStatus always contains MessageHash and Attestation.
// In case of failure, Error is populated with more details.
type AttestationStatus struct {
	// MessageHash is the hash of the message that the attestation was fetched for. It's going to be MessageSent event hash
	MessageHash cciptypes.Bytes
	// Attestation is the attestation data fetched from the API, encoded in bytes
	Attestation cciptypes.Bytes
	// Error is the error that occurred during fetching the attestation data
	Error error
}

func SuccessAttestationStatus(messageHash cciptypes.Bytes, attestation cciptypes.Bytes) AttestationStatus {
	return AttestationStatus{MessageHash: messageHash, Attestation: attestation}
}

func ErrorAttestationStatus(err error) AttestationStatus {
	return AttestationStatus{Error: err}
}

type AttestationEncoder func(context.Context, cciptypes.Bytes, cciptypes.Bytes) (cciptypes.Bytes, error)

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
		msgs map[cciptypes.ChainSelector]map[reader.MessageTokenID]cciptypes.Bytes,
	) (map[cciptypes.ChainSelector]map[reader.MessageTokenID]AttestationStatus, error)
}

type sequentialAttestationClient struct {
	lggr   logger.Logger
	client HTTPClient
	hasher hashutil.Hasher[[32]byte]
}

func NewSequentialAttestationClient(
	lggr logger.Logger,
	config pluginconfig.USDCCCTPObserverConfig,
) (AttestationClient, error) {
	client, err := NewHTTPClient(
		lggr,
		config.AttestationAPI,
		config.AttestationAPIInterval.Duration(),
		config.AttestationAPITimeout.Duration(),
	)
	if err != nil {
		return nil, fmt.Errorf("create HTTP client: %w", err)
	}
	return &sequentialAttestationClient{
		lggr:   lggr,
		client: client,
		hasher: hashutil.NewKeccak(),
	}, nil
}

func (s *sequentialAttestationClient) Attestations(
	ctx context.Context,
	msgs map[cciptypes.ChainSelector]map[reader.MessageTokenID]cciptypes.Bytes,
) (map[cciptypes.ChainSelector]map[reader.MessageTokenID]AttestationStatus, error) {
	outcome := make(map[cciptypes.ChainSelector]map[reader.MessageTokenID]AttestationStatus)

	for chainSelector, hashes := range msgs {
		outcome[chainSelector] = make(map[reader.MessageTokenID]AttestationStatus)

		for tokenID, messageHash := range hashes {
			s.lggr.Debugw(
				"Fetching attestation from the API",
				"chainSelector", chainSelector,
				"messageHash", messageHash,
				"messageTokenID", tokenID,
			)
			// TODO sequential processing
			outcome[chainSelector][tokenID] = s.fetchSingleMessage(ctx, messageHash)
		}
	}
	return outcome, nil
}

func (s *sequentialAttestationClient) fetchSingleMessage(
	ctx context.Context,
	messageHash cciptypes.Bytes,
) AttestationStatus {
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
	msgs map[cciptypes.ChainSelector]map[reader.MessageTokenID]cciptypes.Bytes,
) (map[cciptypes.ChainSelector]map[reader.MessageTokenID]AttestationStatus, error) {
	outcome := make(map[cciptypes.ChainSelector]map[reader.MessageTokenID]AttestationStatus)

	for chainSelector, hashes := range msgs {
		outcome[chainSelector] = make(map[reader.MessageTokenID]AttestationStatus)

		for tokenID, messageHash := range hashes {
			outcome[chainSelector][tokenID] = f.Data[string(messageHash)]
		}
	}
	return outcome, nil
}
