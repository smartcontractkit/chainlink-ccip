package tokendata

import (
	"context"

	"github.com/smartcontractkit/chainlink-ccip/execute/exectypes"
)

type TokenExecutionProcessor interface {
	ProcessTokenData(ctx context.Context, messages exectypes.MessageObservations) (exectypes.TokenDataObservations, error)
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
