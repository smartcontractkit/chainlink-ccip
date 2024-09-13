package tokendata

import (
	"context"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-ccip/execute/exectypes"
)

type TokenExecutionProcessor interface {
	ProcessTokenData(ctx context.Context, observations exectypes.MessageObservations) (exectypes.TokenDataObservations, error)
}

type TokenExecutionComposite struct {
	Processors []TokenExecutionProcessor
}

func (t *TokenExecutionComposite) ProcessTokenData(ctx context.Context, messages exectypes.MessageObservations) (exectypes.TokenDataObservations, error) {
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

func merge(observation exectypes.TokenDataObservations, data exectypes.TokenDataObservations) exectypes.TokenDataObservations {
	return exectypes.TokenDataObservations{}
}

type NoopTokenProcessor struct{}

func (n *NoopTokenProcessor) ProcessTokenData(_ context.Context, observations exectypes.MessageObservations) (exectypes.TokenDataObservations, error) {
	tokenObservations := make(exectypes.TokenDataObservations)

	for selector, obs := range observations {
		tokenObservations[selector] = make(map[cciptypes.SeqNum]exectypes.MessageTokensData)

		for seq, message := range obs {
			tokenData := make([]exectypes.TokenData, len(message.TokenAmounts))
			for i, _ := range message.TokenAmounts {
				tokenData[i] = exectypes.TokenData{
					Ready: true,
					Data:  []byte{},
				}
			}
			tokenObservations[selector][seq] = exectypes.MessageTokensData{TokenData: tokenData}
		}
	}
	return tokenObservations, nil
}
