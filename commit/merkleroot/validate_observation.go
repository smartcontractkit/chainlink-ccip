package merkleroot

import (
	"fmt"

	mapset "github.com/deckarep/golang-set/v2"

	"github.com/smartcontractkit/libocr/commontypes"

	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugintypes"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

func (p *Processor) ValidateObservation(
	_ Outcome,
	q Query,
	ao plugincommon.AttributedObservation[Observation]) error {

	if q.RetryRMNSignatures && !ao.Observation.IsEmpty() {
		return fmt.Errorf("observation should be empty when retrying getting rmn signature")
	}

	obs := ao.Observation
	if err := plugincommon.ValidateFChain(obs.FChain); err != nil {
		return fmt.Errorf("validate FChain: %w", err)
	}
	observerSupportedChains, err := p.chainSupport.SupportedChains(ao.OracleID)
	if err != nil {
		return fmt.Errorf("get supported chains: %w", err)
	}

	supportsDestChain, err := p.chainSupport.SupportsDestChain(ao.OracleID)
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
		p.lggr.Errorw("validate RMNRemoteConfig failed", "err", err, "cfg", obs.RMNRemoteConfig)
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
			return fmt.Errorf("%s invalid: observed chain not supported by Observer %d",
				root, observer)
		}

		if seenChains.Contains(root.ChainSel) {
			return fmt.Errorf("%s invalid: chain already appears in another observed root", root)
		}

		if root.OnRampAddress.IsZeroOrEmpty() {
			return fmt.Errorf("%s invalid: empty OnRampAddress", root)
		}

		if root.MerkleRoot.IsEmpty() {
			return fmt.Errorf("%s invalid: empty MerkleRoot", root)
		}

		if root.SeqNumsRange.Start() == cciptypes.SeqNum(0) {
			return fmt.Errorf("%s invalid: start seq num is 0", root)
		}

		if root.SeqNumsRange.End() == cciptypes.SeqNum(0) {
			return fmt.Errorf("%s invalid: end seq num is 0", root)
		}

		if root.SeqNumsRange.End() < root.SeqNumsRange.Start() {
			return fmt.Errorf("%s invalid: end seq num is less than start seq num", root)
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

		if seqNumChain.SeqNum == 0 {
			return fmt.Errorf("offRampMaxSeqNum for chain %d is 0", seqNumChain.ChainSel)
		}

		if seqNumChain.ChainSel == 0 {
			return fmt.Errorf("offRampMaxSeqNum for chain %d has chain selector 0", seqNumChain.ChainSel)
		}

		seenChains.Add(seqNumChain.ChainSel)
	}

	return nil
}

func validateRMNRemoteConfig(
	observer commontypes.OracleID,
	supportsDestChain bool,
	rmnRemoteConfig cciptypes.RemoteConfig,
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

	if uint64(len(rmnRemoteConfig.Signers)) < rmnRemoteConfig.FSign+1 {
		return fmt.Errorf("not enough signers to cover F+1 threshold")
	}

	if rmnRemoteConfig.ContractAddress.IsZeroOrEmpty() {
		return fmt.Errorf("empty ContractAddress: %s", rmnRemoteConfig.ContractAddress)
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
