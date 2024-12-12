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

	observerSupportedChains, err := p.chainSupport.SupportedChains(ao.OracleID)
	if err != nil {
		return fmt.Errorf("failed to get supported chains: %w", err)
	}

	observedChains := append(maps.Keys(obs.FeeComponents), maps.Keys(obs.NativeTokenPrices)...)

	for _, chain := range observedChains {
		if !observerSupportedChains.Contains(chain) {
			return fmt.Errorf("chain %d is not supported by observer", chain)
		}
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

	return nil
}
