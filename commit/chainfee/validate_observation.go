package chainfee

import (
	"fmt"
	"github.com/pkg/errors"
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
		return errors.Wrap(err, "failed to validate observed chains")
	}

	for _, feeComponent := range obs.FeeComponents {
		if feeComponent.ExecutionFee == nil || feeComponent.ExecutionFee.Cmp(zero) <= 0 {
			return fmt.Errorf("nil or non-positive %s", "execution fee")
		}

		if feeComponent.DataAvailabilityFee == nil || feeComponent.DataAvailabilityFee.Cmp(zero) < 0 {
			return fmt.Errorf("nil or negative %s", "data availability fee")
		}
	}

	for _, token := range obs.NativeTokenPrices {
		if token.Int == nil || token.Int.Cmp(zero) <= 0 {
			return fmt.Errorf("nil or non-positive %s", "execution fee")
		}
	}

	for _, update := range obs.ChainFeeUpdates {
		if update.ChainFee.ExecutionFeePriceUSD == nil || update.ChainFee.ExecutionFeePriceUSD.Cmp(zero) <= 0 {
			return fmt.Errorf("nil or non-positive %s", "execution fee price")
		}

		if update.ChainFee.DataAvFeePriceUSD == nil || update.ChainFee.DataAvFeePriceUSD.Cmp(zero) < 0 {
			return fmt.Errorf("nil or negative %s", "data availability fee price")
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
		return errors.Wrap(err, "failed to get supported chains")
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
