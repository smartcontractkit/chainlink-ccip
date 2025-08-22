package merkleroot

import (
	"context"
	"fmt"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-ccip/internal/libs/slicelib"
	"github.com/smartcontractkit/chainlink-ccip/pkg/reader"
)

func ValidateRootBlessings(
	ctx context.Context, ccipReader reader.CCIPReader, blessedRoots, unblessedRoots []cciptypes.MerkleRootChain) error {
	allSourceChains := make([]cciptypes.ChainSelector, 0, len(blessedRoots)+len(unblessedRoots))
	for _, r := range append(blessedRoots, unblessedRoots...) {
		allSourceChains = append(allSourceChains, r.ChainSel)
	}
	if slicelib.CountUnique(allSourceChains) != len(allSourceChains) {
		return fmt.Errorf("duplicate chain in blessed and unblessed roots")
	}

	sourceChainsConfig, err := ccipReader.GetOffRampSourceChainsConfig(ctx, allSourceChains)
	if err != nil {
		return fmt.Errorf("get offRamp source chains config: %w", err)
	}

	for _, r := range blessedRoots {
		sourceChainCfg, ok := sourceChainsConfig[r.ChainSel]
		if !ok {
			return fmt.Errorf("chain %d is not in the offRampSourceChainsConfig", r.ChainSel)
		}
		if sourceChainCfg.IsRMNVerificationDisabled {
			return fmt.Errorf("chain %d is RMN-disabled but root is blessed", r.ChainSel)
		}
		if !sourceChainCfg.IsEnabled {
			return fmt.Errorf("chain %d is disabled but root is blessed", r.ChainSel)
		}
	}

	for _, r := range unblessedRoots {
		sourceChainCfg, ok := sourceChainsConfig[r.ChainSel]
		if !ok {
			return fmt.Errorf("chain %d is not in the offRampSourceChainsConfig", r.ChainSel)
		}
		if !sourceChainCfg.IsRMNVerificationDisabled {
			return fmt.Errorf("chain %d is RMN-enabled but root is unblessed", r.ChainSel)
		}
		if !sourceChainCfg.IsEnabled {
			return fmt.Errorf("chain %d is disabled but root is reported", r.ChainSel)
		}
	}

	return nil
}
