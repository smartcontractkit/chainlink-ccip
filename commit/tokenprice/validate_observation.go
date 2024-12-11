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

	if err := plugincommon.ValidateFChain(obs.FChain); err != nil {
		return fmt.Errorf("failed to validate FChain: %w", err)
	}

	observerSupportedChains, err := p.chainSupport.SupportedChains(ao.OracleID)
	if err != nil {
		return fmt.Errorf("failed to get supported chains: %w", err)
	}

	for chain := range obs.FChain {
		if !observerSupportedChains.Contains(chain) {
			return fmt.Errorf("chain %d is not supported by observer", chain)
		}
	}

	if err := validateObservedTokenPrices(obs.FeedTokenPrices); err != nil {
		return fmt.Errorf("failed to validate observed token prices: %w", err)
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
