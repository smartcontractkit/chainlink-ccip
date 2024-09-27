package merkleroot

import (
	"context"
	"fmt"

	mapset "github.com/deckarep/golang-set/v2"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
	"github.com/smartcontractkit/libocr/commontypes"

	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugintypes"
	"github.com/smartcontractkit/chainlink-ccip/pkg/reader"
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
		return fmt.Errorf("failed to validate FChain: %w", err)
	}
	observerSupportedChains, err := w.chainSupport.SupportedChains(ao.OracleID)
	if err != nil {
		return fmt.Errorf("failed to get supported chains: %w", err)
	}

	supportsDestChain, err := w.chainSupport.SupportsDestChain(ao.OracleID)
	if err != nil {
		return fmt.Errorf("call to supportsDestChain failed: %w", err)
	}

	if err := validateObservedMerkleRoots(obs.MerkleRoots, ao.OracleID, observerSupportedChains); err != nil {
		return fmt.Errorf("failed to validate MerkleRoots: %w", err)
	}

	if err := validateObservedOnRampMaxSeqNums(obs.OnRampMaxSeqNums, ao.OracleID, observerSupportedChains); err != nil {
		return fmt.Errorf("failed to validate OnRampMaxSeqNums: %w", err)
	}

	if err := validateObservedOffRampMaxSeqNums(obs.OffRampNextSeqNums, ao.OracleID, supportsDestChain); err != nil {
		return fmt.Errorf("failed to validate OffRampNextSeqNums: %w", err)
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

func validateFChain(fChain map[cciptypes.ChainSelector]int) error {
	for _, f := range fChain {
		if f < 0 {
			return fmt.Errorf("fChain %d is negative", f)
		}
	}

	return nil
}

// ValidateMerkleRootsState merkle roots seq nums validation by comparing with on-chain state.
func ValidateMerkleRootsState(
	ctx context.Context,
	lggr logger.Logger,
	report cciptypes.CommitPluginReport,
	reader reader.CCIPReader,
) (bool, error) {
	reportChains := make([]cciptypes.ChainSelector, 0)
	reportMinSeqNums := make(map[cciptypes.ChainSelector]cciptypes.SeqNum)
	for _, mr := range report.MerkleRoots {
		reportChains = append(reportChains, mr.ChainSel)
		reportMinSeqNums[mr.ChainSel] = mr.SeqNumsRange.Start()
	}

	if len(reportChains) == 0 {
		return true, nil
	}

	onchainNextSeqNums, err := reader.NextSeqNum(ctx, reportChains)
	if err != nil {
		return false, fmt.Errorf("get next sequence numbers: %w", err)
	}
	if len(onchainNextSeqNums) != len(reportChains) {
		return false, fmt.Errorf("critical error: onchainSeqNums length mismatch")
	}

	for i, nextSeqNum := range onchainNextSeqNums {
		chain := reportChains[i]
		reportMinSeqNum, ok := reportMinSeqNums[chain]
		if !ok {
			return false, fmt.Errorf("critical error: reportSeqNum not found for chain %d", chain)
		}

		if reportMinSeqNum != nextSeqNum {
			lggr.Warnw("report is not valid due to seq num mismatch",
				"chain", chain, "reportMinSeqNum", reportMinSeqNum, "onchainNextSeqNum", nextSeqNum)
			return false, nil
		}
	}

	return true, nil
}
