package merkleroot

import (
	"context"
	"fmt"
	"sort"

	mapset "github.com/deckarep/golang-set/v2"

	"github.com/smartcontractkit/libocr/commontypes"

	"github.com/smartcontractkit/chainlink-ccip/commit/merkleroot/rmn/types"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugintypes"
	"github.com/smartcontractkit/chainlink-ccip/pkg/reader"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

func (w *Processor) ValidateObservation(
	_ Outcome,
	q Query,
	ao plugincommon.AttributedObservation[Observation]) error {

	if q.RetryRMNSignatures && !ao.Observation.IsEmpty() {
		return fmt.Errorf("observation should be empty when retrying getting rmn signature")
	}

	obs := ao.Observation
	if err := validateFChain(obs.FChain); err != nil {
		return fmt.Errorf("validate FChain: %w", err)
	}
	observerSupportedChains, err := w.chainSupport.SupportedChains(ao.OracleID)
	if err != nil {
		return fmt.Errorf("get supported chains: %w", err)
	}

	supportsDestChain, err := w.chainSupport.SupportsDestChain(ao.OracleID)
	if err != nil {
		return fmt.Errorf("call to supportsDestChain failed: %w", err)
	}

	if err := validateObservedMerkleRoots(obs.MerkleRoots, ao.OracleID, observerSupportedChains); err != nil {
		return fmt.Errorf("validate MerkleRoots: %w", err)
	}

	if err := validateObservedOnRampMaxSeqNums(obs.OnRampMaxSeqNums, ao.OracleID, observerSupportedChains); err != nil {
		return fmt.Errorf("validate OnRampMaxSeqNums: %w", err)
	}

	if err := validateObservedOffRampMaxSeqNums(obs.OffRampNextSeqNums, ao.OracleID, supportsDestChain); err != nil {
		return fmt.Errorf("validate OffRampNextSeqNums: %w", err)
	}

	if err := validateRMNRemoteConfig(ao.OracleID, supportsDestChain, obs.RMNRemoteConfig); err != nil {
		w.lggr.Errorw("validate RMNRemoteConfig failed", "err", err, "cfg", obs.RMNRemoteConfig)
		return fmt.Errorf("validate RMNRemoteConfig: %w", err)
	}

	return nil
}

func validateObservedMerkleRoots(
	merkleRoots []cciptypes.MerkleRootChain,
	observer commontypes.OracleID,
	observerSupportedChains mapset.Set[cciptypes.ChainSelector],
) error {
	if len(merkleRoots) == 0 {
		return nil
	}

	seenChains := mapset.NewSet[cciptypes.ChainSelector]()
	for _, root := range merkleRoots {
		if !observerSupportedChains.Contains(root.ChainSel) {
			return fmt.Errorf("found merkle root for chain %d, but this chain is not supported by Observer %d",
				root.ChainSel, observer)
		}

		if seenChains.Contains(root.ChainSel) {
			return fmt.Errorf("duplicate merkle root for chain %d", root.ChainSel)
		}
		seenChains.Add(root.ChainSel)
	}

	return nil
}

func validateObservedOnRampMaxSeqNums(
	onRampMaxSeqNums []plugintypes.SeqNumChain,
	observer commontypes.OracleID,
	observerSupportedChains mapset.Set[cciptypes.ChainSelector],
) error {
	if len(onRampMaxSeqNums) == 0 {
		return nil
	}

	seenChains := mapset.NewSet[cciptypes.ChainSelector]()
	for _, seqNumChain := range onRampMaxSeqNums {
		if !observerSupportedChains.Contains(seqNumChain.ChainSel) {
			return fmt.Errorf("found onRampMaxSeqNum for chain %d, but this chain is not supported by Observer %d, "+
				"observerSupportedChains: %v, onRampMaxSeqNums: %v",
				seqNumChain.ChainSel, observer, observerSupportedChains, onRampMaxSeqNums)
		}

		if seenChains.Contains(seqNumChain.ChainSel) {
			return fmt.Errorf("duplicate onRampMaxSeqNum for chain %d", seqNumChain.ChainSel)
		}
		seenChains.Add(seqNumChain.ChainSel)
	}

	return nil
}

func validateObservedOffRampMaxSeqNums(
	offRampMaxSeqNums []plugintypes.SeqNumChain,
	observer commontypes.OracleID,
	supportsDestChain bool,
) error {
	if len(offRampMaxSeqNums) == 0 {
		return nil
	}

	if !supportsDestChain {
		return fmt.Errorf("observer %d does not support dest chain, but has observed %d offRampMaxSeqNums",
			observer, len(offRampMaxSeqNums))
	}

	seenChains := mapset.NewSet[cciptypes.ChainSelector]()
	for _, seqNumChain := range offRampMaxSeqNums {
		if seenChains.Contains(seqNumChain.ChainSel) {
			return fmt.Errorf("duplicate offRampMaxSeqNum for chain %d", seqNumChain.ChainSel)
		}
		seenChains.Add(seqNumChain.ChainSel)
	}

	return nil
}

func validateRMNRemoteConfig(
	observer commontypes.OracleID,
	supportsDestChain bool,
	rmnRemoteConfig types.RemoteConfig,
) error {
	if rmnRemoteConfig.IsEmpty() {
		return nil
	}

	if !supportsDestChain {
		return fmt.Errorf("oracle %d does not support dest chain, but has observed an RMNRemoteConfig", observer)
	}

	if rmnRemoteConfig.ConfigDigest.IsEmpty() {
		return fmt.Errorf("empty ConfigDigest")
	}

	if rmnRemoteConfig.RmnReportVersion.IsEmpty() {
		return fmt.Errorf("empty RmnReportVersion")
	}

	if uint64(len(rmnRemoteConfig.Signers)) < rmnRemoteConfig.F+1 {
		return fmt.Errorf("not enough signers to cover F+1 threshold")
	}

	if len(rmnRemoteConfig.ContractAddress) == 0 {
		return fmt.Errorf("empty ContractAddress")
	}

	seenNodeIndexes := mapset.NewSet[uint64]()
	for _, signer := range rmnRemoteConfig.Signers {
		if len(signer.OnchainPublicKey) == 0 {
			return fmt.Errorf("empty signer OnchainPublicKey")
		}

		if seenNodeIndexes.Contains(signer.NodeIndex) {
			return fmt.Errorf("duplicate NodeIndex %d", signer.NodeIndex)
		}
		seenNodeIndexes.Add(signer.NodeIndex)
	}

	return nil
}

func validateFChain(fChain map[cciptypes.ChainSelector]int) error {
	for chainSelector, f := range fChain {
		if f <= 0 {
			return fmt.Errorf("fChain for chain %d is not positive: %d", chainSelector, f)
		}
	}

	return nil
}

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

	if len(offRampExpNextSeqNums) != len(chainSlice) {
		return fmt.Errorf("critical reader error: seq nums length mismatch")
	}

	for i, offRampExpNextSeqNum := range offRampExpNextSeqNums {
		chain := chainSlice[i]

		newNextOnRampSeqNum, ok := newNextOnRampSeqNums[chain]
		if !ok {
			return fmt.Errorf("critical unexpected error: newOnRampSeqNum not found")
		}

		if newNextOnRampSeqNum != offRampExpNextSeqNum {
			return fmt.Errorf("unexpected seq nums offRampNext=%d newOnRampNext=%d",
				offRampExpNextSeqNum, newNextOnRampSeqNum)
		}
	}

	return nil
}
