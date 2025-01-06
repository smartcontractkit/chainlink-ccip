package chainfee

import (
	"fmt"
	"math/big"

	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon"

	"golang.org/x/exp/maps"
)

func (p *processor) ValidateObservation(
	_ Outcome,
	_ Query,
	ao plugincommon.AttributedObservation[Observation],
) error {
	obs := ao.Observation
	zero := big.NewInt(0)

	if err := plugincommon.ValidateFChain(obs.FChain); err != nil {
		return fmt.Errorf("failed to validate FChain: %w", err)
	}

	if err := p.ValidateObservedChains(ao); err != nil {
		return fmt.Errorf("failed to validate observed chains: %w", err)
	}

	if err := validateFeeComponents(ao); err != nil {
		return fmt.Errorf("failed to validate fee components: %w", err)
	}

	if err := validateChainFeeUpdates(ao); err != nil {
		return fmt.Errorf("failed to validate chain fee updates: %w", err)
	}

	for _, token := range obs.NativeTokenPrices {
		if token.Int == nil || token.Int.Cmp(zero) <= 0 {
			return fmt.Errorf("nil or non-positive %s", "execution fee")
		}
	}

	return nil
}

func validateChainFeeUpdates(
	ao plugincommon.AttributedObservation[Observation],
) error {
	for _, update := range ao.Observation.ChainFeeUpdates {
		if update.ChainFee.ExecutionFeePriceUSD == nil || update.ChainFee.ExecutionFeePriceUSD.Cmp(big.NewInt(0)) <= 0 {
			return fmt.Errorf("nil or non-positive %s", "execution fee price")
		}

		if update.ChainFee.DataAvFeePriceUSD == nil || update.ChainFee.DataAvFeePriceUSD.Cmp(big.NewInt(0)) < 0 {
			return fmt.Errorf("nil or negative %s", "data availability fee price")
		}
		if update.Timestamp.IsZero() {
			return fmt.Errorf("zero timestamp")
		}
	}
	return nil
}

func validateFeeComponents(
	ao plugincommon.AttributedObservation[Observation],
) error {
	for _, feeComponent := range ao.Observation.FeeComponents {
		if feeComponent.ExecutionFee == nil || feeComponent.ExecutionFee.Cmp(big.NewInt(0)) <= 0 {
			return fmt.Errorf("nil or non-positive %s", "execution fee")
		}

		if feeComponent.DataAvailabilityFee == nil || feeComponent.DataAvailabilityFee.Cmp(big.NewInt(0)) < 0 {
			return fmt.Errorf("nil or negative %s", "data availability fee")
		}
	}
	return nil
}

func (p *processor) ValidateObservedChains(
	ao plugincommon.AttributedObservation[Observation],
) error {
	obs := ao.Observation
	observerSupportedChains, err := p.chainSupport.SupportedChains(ao.OracleID)
	if err != nil {
		return fmt.Errorf("failed to get supported chains: %w", err)
	}

	observedChains := append(maps.Keys(obs.FeeComponents), maps.Keys(obs.NativeTokenPrices)...)
	observedChains = append(observedChains, maps.Keys(obs.ChainFeeUpdates)...)
	for _, chain := range observedChains {
		if !observerSupportedChains.Contains(chain) {
			return fmt.Errorf("chain %d is not supported by observer", chain)
		}
	}

	return nil
}
