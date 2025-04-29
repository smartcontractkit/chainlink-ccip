package chainfee

import (
	"fmt"
	"math/big"
	"time"

	mapset "github.com/deckarep/golang-set/v2"

	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon"
	"github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"

	"golang.org/x/exp/maps"
)

func (p *processor) ValidateObservation(
	_ Outcome,
	_ Query,
	ao plugincommon.AttributedObservation[Observation],
) error {
	obs := ao.Observation

	if obs.IsEmpty() {
		return nil
	}

	if err := plugincommon.ValidateFChain(obs.FChain); err != nil {
		return fmt.Errorf("failed to validate FChain: %w", err)
	}
	observerSupportedChains, err := p.chainSupport.SupportedChains(ao.OracleID)
	if err != nil {
		return fmt.Errorf("failed to get supported chains: %w", err)
	}
	if err := validateObservedChains(ao, observerSupportedChains); err != nil {
		return fmt.Errorf("failed to validate observed chains: %w", err)
	}

	if err := validateFeeComponents(ao); err != nil {
		return fmt.Errorf("failed to validate fee components: %w", err)
	}

	if err := validateChainFeeUpdates(ao); err != nil {
		return fmt.Errorf("failed to validate chain fee updates: %w", err)
	}

	for _, token := range obs.NativeTokenPrices {
		if token.Int == nil {
			return fmt.Errorf("nil token price: %+v", token)
		}
		if token.Int.Cmp(big.NewInt(0)) <= 0 {
			return fmt.Errorf("non-positive token price: %+v", token)
		}
	}

	if obs.TimestampNow.IsZero() {
		return fmt.Errorf("timestamp now cannot be zero")
	}
	if obs.TimestampNow.After(time.Now().UTC()) {
		return fmt.Errorf("timestamp now %s cannot be in the future", obs.TimestampNow.String())
	}
	return nil
}

func validateChainFeeUpdates(
	ao plugincommon.AttributedObservation[Observation],
) error {
	for chainSelector, update := range ao.Observation.ChainFeeUpdates {
		if update.ChainFee.ExecutionFeePriceUSD == nil || update.ChainFee.ExecutionFeePriceUSD.Cmp(big.NewInt(0)) < 0 {
			return fmt.Errorf("nil or negative execution fee: %+v, chain: %d",
				update.ChainFee.ExecutionFeePriceUSD, chainSelector)
		}

		if update.ChainFee.DataAvFeePriceUSD == nil || update.ChainFee.DataAvFeePriceUSD.Cmp(big.NewInt(0)) < 0 {
			return fmt.Errorf("nil or negative data availability fee: %+v, chain: %d",
				update.ChainFee.DataAvFeePriceUSD, chainSelector)
		}
		if update.Timestamp.IsZero() {
			return fmt.Errorf("timestamp cannot be zero, chain: %d", chainSelector)
		}
		if update.Timestamp.After(time.Now().UTC()) {
			return fmt.Errorf("timestamp %s cannot be in the future, chain: %d", update.Timestamp.String(), chainSelector)
		}
	}
	return nil
}

func validateFeeComponents(
	ao plugincommon.AttributedObservation[Observation],
) error {
	for chainSelector, feeComponent := range ao.Observation.FeeComponents {
		if feeComponent.ExecutionFee == nil || feeComponent.ExecutionFee.Cmp(big.NewInt(0)) < 0 {
			return fmt.Errorf("nil or negative execution fee: %+v, chain: %d",
				feeComponent.ExecutionFee, chainSelector)
		}

		if feeComponent.DataAvailabilityFee == nil || feeComponent.DataAvailabilityFee.Cmp(big.NewInt(0)) < 0 {
			return fmt.Errorf("nil or negative data availability fee: %+v, chain: %d",
				feeComponent.DataAvailabilityFee, chainSelector)
		}
	}
	return nil
}

func validateObservedChains(
	ao plugincommon.AttributedObservation[Observation],
	observerSupportedChains mapset.Set[ccipocr3.ChainSelector],
) error {
	obs := ao.Observation
	if !areMapKeysEqual(obs.FeeComponents, obs.NativeTokenPrices) {
		return fmt.Errorf("fee components and native token prices have different observed chains")
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

func areMapKeysEqual[T, T1 comparable](map1 map[ccipocr3.ChainSelector]T, map2 map[ccipocr3.ChainSelector]T1) bool {
	if len(map1) != len(map2) {
		return false
	}

	for key := range map1 {
		if _, exists := map2[key]; !exists {
			return false
		}
	}

	return true
}
