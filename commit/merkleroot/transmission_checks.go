package merkleroot

import (
	"context"
	"fmt"
	"sort"

	mapset "github.com/deckarep/golang-set/v2"

	"github.com/smartcontractkit/chainlink-ccip/pkg/reader"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

// ValidateMerkleRootsState validates the proposed merkle roots against the current on-chain state.
// This function is not-pure as it reads from the chain by making one network/reader call.
func ValidateMerkleRootsState(
	ctx context.Context,
	proposedMerkleRoots []cciptypes.MerkleRootChain,
	reader reader.CCIPReader,
) error {
	if len(proposedMerkleRoots) == 0 {
		return nil
	}

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
			return fmt.Errorf("unexpected seq nums offRampNext=%d newOnRampNext=%d",
				offRampExpNextSeqNum, newNextOnRampSeqNum)
		}
	}

	return nil
}
