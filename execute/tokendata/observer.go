package tokendata

import (
	"context"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-ccip/execute/exectypes"
)

type TokenDataObserver interface {
	Observe(
		ctx context.Context,
		observations exectypes.MessageObservations,
	) (exectypes.TokenDataObservations, error)
}

// CompositeTokenDataObserver is a TokenDataObserver that combines multiple TokenDataObserver behind the same interface.
// Goal of that is to support multiple token observers supporting different tokens (e.g. CCTP, MyFancyToken etc)
type CompositeTokenDataObserver struct {
	Observers []TokenDataObserver
}

func (t *CompositeTokenDataObserver) Observe(
	ctx context.Context,
	messages exectypes.MessageObservations,
) (exectypes.TokenDataObservations, error) {
	var tokenObservation exectypes.TokenDataObservations
	for _, ob := range t.Observers {
		tokenData, err := ob.Observe(ctx, messages)
		if err != nil {
			return nil, err
		}
		tokenObservation = merge(tokenObservation, tokenData)
	}
	return tokenObservation, nil
}

func merge(_ exectypes.TokenDataObservations, _ exectypes.TokenDataObservations) exectypes.TokenDataObservations {
	return exectypes.TokenDataObservations{}
}

type NoopTokenDataObserver struct{}

func (n *NoopTokenDataObserver) Observe(
	_ context.Context,
	observations exectypes.MessageObservations,
) (exectypes.TokenDataObservations, error) {
	tokenObservations := make(exectypes.TokenDataObservations)

	for selector, obs := range observations {
		tokenObservations[selector] = make(map[cciptypes.SeqNum]exectypes.MessageTokenData)

		for seq, message := range obs {
			tokenData := make([]exectypes.TokenData, len(message.TokenAmounts))
			for i := range message.TokenAmounts {
				tokenData[i] = exectypes.NewEmptyTokenData()
			}
			tokenObservations[selector][seq] = exectypes.MessageTokenData{TokenData: tokenData}
		}
	}
	return tokenObservations, nil
}
