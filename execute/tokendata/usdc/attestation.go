package usdc

import (
	"context"
	"errors"

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
