package tokenprice

import (
	"fmt"
	"time"

	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
	"github.com/smartcontractkit/chainlink-ccip/pluginconfig"
)

func (p *processor) ValidateObservation(
	prevOutcome Outcome,
	query Query,
	ao plugincommon.AttributedObservation[Observation],
) error {
	obs := ao.Observation

	if obs.IsEmpty() {
		return nil
	}

	if err := plugincommon.ValidateFChain(obs.FChain); err != nil {
		return fmt.Errorf("failed to validate FChain: %w", err)
	}

	supportedChains, err := p.chainSupport.SupportedChains(ao.OracleID)
	if err != nil {
		return fmt.Errorf("failed to get supported chains: %w", err)
	}

	if len(obs.FeedTokenPrices) > 0 && !supportedChains.Contains(p.offChainCfg.PriceFeedChainSelector) {
		return fmt.Errorf("feed chain must be supported to read feed token prices, oreacleID: %d feedChain: %d",
			ao.OracleID, p.offChainCfg.PriceFeedChainSelector)

	}

	if err = validateObservedTokenPrices(obs.FeedTokenPrices, p.offChainCfg.TokenInfo); err != nil {
		return fmt.Errorf("failed to validate observed token prices: %w", err)
	}

	if len(obs.FeeQuoterTokenUpdates) > 0 && !supportedChains.Contains(p.destChain) {
		return fmt.Errorf("dest chain must be supported to read fee quoter token updates "+
			"oracleID: %d destChain: %d", ao.OracleID, p.destChain)
	}

	if err = validateObservedTokenUpdates(obs.FeeQuoterTokenUpdates, p.offChainCfg.TokenInfo); err != nil {
		return fmt.Errorf("failed to validate observed fee quoter token updates: %w", err)
	}
	if obs.Timestamp.IsZero() || obs.Timestamp.After(time.Now().UTC()) {
		return fmt.Errorf("invalid timestamp value %s", obs.Timestamp.String())
	}

	return nil
}

func validateObservedTokenPrices(
	tokenPrices cciptypes.TokenPriceMap,
	tokensToQuery map[cciptypes.UnknownEncodedAddress]pluginconfig.TokenInfo) error {
	for tokenID, price := range tokenPrices {
		if _, ok := tokensToQuery[tokenID]; !ok {
			return fmt.Errorf("observed token %v is not in the list of tokens to query", tokenID)
		}
		if !price.IsPositive() {
			return fmt.Errorf("non positive value for token price of token %v", tokenID)
		}
	}
	return nil
}

func validateObservedTokenUpdates(
	tokenUpdates map[cciptypes.UnknownEncodedAddress]cciptypes.TimestampedBig,
	tokensToQuery map[cciptypes.UnknownEncodedAddress]pluginconfig.TokenInfo) error {
	for tokenID, update := range tokenUpdates {
		if _, ok := tokensToQuery[tokenID]; !ok {
			return fmt.Errorf("observed token %v is not in the list of tokens to query", tokenID)
		}
		if !update.Value.IsPositive() {
			return fmt.Errorf("non positive value for token update of token %v", tokenID)
		}
		if update.Timestamp.IsZero() || update.Timestamp.After(time.Now().UTC()) {
			return fmt.Errorf("token update of token %v must have "+
				"a timestamp in the past that's not 0. timestamp observed: %s", tokenID, update.Timestamp.String())
		}
	}

	return nil
}
