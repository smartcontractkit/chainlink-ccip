package tokenprice

import (
	"fmt"

	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

func (p *processor) ValidateObservation(
	prevOutcome Outcome,
	query Query,
	ao plugincommon.AttributedObservation[Observation],
) error {
	obs := ao.Observation

	if err := plugincommon.ValidateFChain(obs.FChain); err != nil {
		return fmt.Errorf("failed to validate FChain: %w", err)
	}

	if err := validateObservedTokenPrices(obs.FeedTokenPrices); err != nil {
		return fmt.Errorf("failed to validate observed token prices: %w", err)
	}

	//TODO: validate observed fee quoter token updates
	//TODO: validate observed timestamp

	return nil
}

func validateObservedTokenPrices(tokenPrices cciptypes.TokenPriceMap) error {
	// TODO: cross check that only tokens from tokensToQuery are present
	for tokenID, price := range tokenPrices {
		if price.IsEmpty() {
			return fmt.Errorf("token price of token %v must not be empty", tokenID)
		}
	}
	return nil
}
