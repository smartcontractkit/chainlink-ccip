package tokenprice

import (
	"fmt"

	mapset "github.com/deckarep/golang-set/v2"

	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

func (p *processor) ValidateObservation(
	prevOutcome Outcome,
	query Query,
	ao plugincommon.AttributedObservation[Observation],
) error {
	obs := ao.Observation

	if err := validateFChain(obs.FChain); err != nil {
		return fmt.Errorf("failed to validate FChain: %w", err)
	}

	if err := validateObservedTokenPrices(obs.FeedTokenPrices); err != nil {
		return fmt.Errorf("failed to validate observed token prices: %w", err)
	}

	return nil
}

func validateFChain(fChain map[cciptypes.ChainSelector]int) error {
	for chainSelector, f := range fChain {
		if f < 0 {
			return fmt.Errorf("fChain for chain %d is not positive: %d", chainSelector, f)
		}
	}

	return nil
}

func validateObservedTokenPrices(tokenPrices []cciptypes.TokenPrice) error {
	tokensWithPrice := mapset.NewSet[cciptypes.UnknownEncodedAddress]()
	for _, t := range tokenPrices {
		if tokensWithPrice.Contains(t.TokenID) {
			return fmt.Errorf("duplicate token price for token: %s", t.TokenID)
		}
		tokensWithPrice.Add(t.TokenID)

		if t.Price.IsEmpty() {
			return fmt.Errorf("token price of token %v must not be empty", t.TokenID)
		}
	}
	return nil
}
