package tokendata

import (
	"context"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-ccip/execute/exectypes"
)

type TokenDataProcessor interface {
	ProcessTokenData(
		ctx context.Context,
		observations exectypes.MessageObservations,
	) (exectypes.TokenDataObservations, error)
}

// CompositeTokenDataProcessor is a TokenDataProcessor that combines multiple TokenDataProcessors.
// Goal of that is to support multiple token processors supporting different tokens (e.g. CCTP, MyFancyToken etc)
type CompositeTokenDataProcessor struct {
	Processors []TokenDataProcessor
}

func (t *CompositeTokenDataProcessor) ProcessTokenData(
	ctx context.Context,
	messages exectypes.MessageObservations,
) (exectypes.TokenDataObservations, error) {
	// Dummy implementation, don't focus on that
	var tokenObservation exectypes.TokenDataObservations
	for _, processor := range t.Processors {
		tokenData, err := processor.ProcessTokenData(ctx, messages)
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

type NoopTokenProcessor struct{}

func (n *NoopTokenProcessor) ProcessTokenData(
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
