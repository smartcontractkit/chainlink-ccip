package chainfee

import (
	"fmt"
	"math/big"

	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
)

func (p *processor) ValidateObservation(
	prevOutcome Outcome,
	query Query,
	ao plugincommon.AttributedObservation[Observation],
) error {
	//obs := ao.Observation
	//
	//if err := validateFChain(obs.FChain); err != nil {
	//	return fmt.Errorf("failed to validate FChain: %w", err)
	//}
	//
	//observerSupportedChains, err := p.chainSupport.SupportedChains(ao.OracleID)
	//if err != nil {
	//	return fmt.Errorf("failed to get supported chains: %w", err)
	//}
	//
	//observedChains := append(maps.Keys(obs.FeeComponents), maps.Keys(obs.NativeTokenPrices)...)
	//
	//for _, chain := range observedChains {
	//	if !observerSupportedChains.Contains(chain) {
	//		return fmt.Errorf("chain %d is not supported by observer", chain)
	//	}
	//}
	//
	//for _, feeComponent := range obs.FeeComponents {
	//	err := validateBigInt(feeComponent.ExecutionFee, "execution fee")
	//	if err != nil {
	//		return err
	//	}
	//
	//	err = validateBigInt(feeComponent.DataAvailabilityFee, "data availability fee")
	//	if err != nil {
	//		return err
	//	}
	//
	//}
	//
	//for _, token := range obs.NativeTokenPrices {
	//	err := validateBigInt(token.Int, "native token price")
	//	if err != nil {
	//		return err
	//	}
	//}
	return nil
}

func validateBigInt(b *big.Int, name string) error {
	zero := big.NewInt(0)
	if b == nil {
		return fmt.Errorf("nil %s", name)
	}
	if b.Cmp(zero) < 0 || b.Cmp(zero) == 0 {
		return fmt.Errorf("%s must be positive", name)
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
