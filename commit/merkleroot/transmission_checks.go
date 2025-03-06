package merkleroot

import (
	"context"
	"fmt"
	"sort"

	mapset "github.com/deckarep/golang-set/v2"

	"github.com/smartcontractkit/chainlink-ccip/internal/libs/slicelib"
	"github.com/smartcontractkit/chainlink-ccip/pkg/reader"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

// ValidateMerkleRootsState validates the proposed merkle roots against the current on-chain state.
// This function is not-pure as it reads from the chain by making one network/reader call.
func ValidateMerkleRootsState(
	ctx context.Context,
	proposedBlessedMerkleRoots []cciptypes.MerkleRootChain,
	proposedUnblessedMerkleRoots []cciptypes.MerkleRootChain,
	reader reader.CCIPReader,
) error {
	if len(proposedBlessedMerkleRoots) == 0 && len(proposedUnblessedMerkleRoots) == 0 {
		return nil
	}

	proposedMerkleRoots := append(proposedBlessedMerkleRoots, proposedUnblessedMerkleRoots...)

	chainSet := mapset.NewSet[cciptypes.ChainSelector]()
	newNextOnRampSeqNums := make(map[cciptypes.ChainSelector]cciptypes.SeqNum)

	for _, r := range proposedMerkleRoots {
		if chainSet.Contains(r.ChainSel) {
			return fmt.Errorf("duplicate chain %d", r.ChainSel)
		}
		chainSet.Add(r.ChainSel)
		newNextOnRampSeqNums[r.ChainSel] = r.SeqNumsRange.Start()
	}

	chainSlice := chainSet.ToSlice()
	sort.Slice(chainSlice, func(i, j int) bool { return chainSlice[i] < chainSlice[j] })

	offRampExpNextSeqNums, err := reader.NextSeqNum(ctx, chainSlice)
	if err != nil {
		return fmt.Errorf("get next sequence numbers: %w", err)
	}

	for chain, newNextOnRampSeqNum := range newNextOnRampSeqNums {
		offRampExpNextSeqNum, ok := offRampExpNextSeqNums[chain]
		if !ok {
			// Due to some chain being disabled while the sequence numbers were already observed.
			// Report should not be considered valid in that case.
			return fmt.Errorf("offRamp expected next sequence number for chain %d was not found", chain)
		}

		if newNextOnRampSeqNum != offRampExpNextSeqNum {
			return fmt.Errorf("the merkle root that we are about to propose is stale, some previous report "+
				"made it on-chain, consider waiting more time for the reports to make it on chain. "+
				"offramp expects %d but we are proposing a root with min seq num %d for chain %d",
				offRampExpNextSeqNum, newNextOnRampSeqNum, chain,
			)
		}
	}

	return validateRootBlessings(ctx, reader, proposedBlessedMerkleRoots, proposedUnblessedMerkleRoots)
}

func validateRootBlessings(
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
