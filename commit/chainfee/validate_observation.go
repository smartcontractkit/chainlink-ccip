package chainfee

import (
	"fmt"
	"math/big"

	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"golang.org/x/exp/maps"
)

// nolint:revive,govet
func (p *processor) ValidateObservation(
	prevOutcome Outcome,
	query Query,
	ao plugincommon.AttributedObservation[Observation],
) error {
	obs := ao.Observation
	zero := big.NewInt(0)

	if err := validateFChain(obs.FChain); err != nil {
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

func validateFChain(fChain map[cciptypes.ChainSelector]int) error {
	for _, f := range fChain {
		if f < 0 {
			return fmt.Errorf("fChain %d is negative", f)
		}
	}
	return nil
}
