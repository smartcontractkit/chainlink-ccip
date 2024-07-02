package commitrmnocb

import (
	"fmt"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/smartcontractkit/libocr/commontypes"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/types"

	"github.com/smartcontractkit/chainlink-ccip/commit"
	"github.com/smartcontractkit/chainlink-ccip/plugintypes"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
)

// ValidateObservation TODO: doc
func (p *Plugin) ValidateObservation(_ ocr3types.OutcomeContext, _ types.Query, ao types.AttributedObservation) error {
	obs, err := DecodeCommitPluginObservation(ao.Observation)
	if err != nil {
		return fmt.Errorf("failed to decode commit plugin observation: %w", err)
	}

	if err := p.validateFChain(obs.FChain); err != nil {
		return fmt.Errorf("failed to validate FChain: %w", err)
	}

	if err := p.validateObservedMerkleRoots(obs.MerkleRoots, ao.Observer); err != nil {
		return fmt.Errorf("failed to validate MerkleRoots: %w", err)
	}

	if err := p.validateObservedOnRampMaxSeqNums(obs.OnRampMaxSeqNums, ao.Observer); err != nil {
		return fmt.Errorf("failed to validate OnRampMaxSeqNums: %w", err)
	}

	if err := p.validateObservedOffRampMaxSeqNums(obs.OffRampMaxSeqNums, ao.Observer); err != nil {
		return fmt.Errorf("failed to validate OffRampMaxSeqNums: %w", err)
	}

	if err := commit.ValidateObservedTokenPrices(obs.TokenPrices); err != nil {
		return fmt.Errorf("failed to validate token prices: %w", err)
	}

	if err := commit.ValidateObservedGasPrices(obs.GasPrices); err != nil {
		return fmt.Errorf("failed to validate gas prices: %w", err)
	}

	return nil
}

// validateMerkleRoots TODO: doc
// No duplicate chains, only contains chainSelector that the owner can read
func (p *Plugin) validateObservedMerkleRoots(merkleRoots []MerkleRootAndChain, observer commontypes.OracleID) error {
	if merkleRoots == nil || len(merkleRoots) == 0 {
		return nil
	}

	observerSupportedChains, err := p.supportedChains(observer)

	if err != nil {
		return fmt.Errorf("call to supportedChains failed: %w", err)
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

// validateObservedOnRampMaxSeqNums TODO: doc
func (p *Plugin) validateObservedOnRampMaxSeqNums(
	onRampMaxSeqNums []plugintypes.SeqNumChain,
	observer commontypes.OracleID,
) error {
	if onRampMaxSeqNums == nil || len(onRampMaxSeqNums) == 0 {
		return nil
	}

	observerSupportedChains, err := p.supportedChains(observer)

	if err != nil {
		return fmt.Errorf("call to supportedChains failed: %w", err)
	}

	seenChains := mapset.NewSet[cciptypes.ChainSelector]()
	for _, seqNumChain := range onRampMaxSeqNums {
		if !observerSupportedChains.Contains(seqNumChain.ChainSel) {
			return fmt.Errorf("found onRampMaxSeqNum for chain %d, but this chain is not supported by Observer %d",
				seqNumChain.ChainSel, observer)
		}

		if seenChains.Contains(seqNumChain.ChainSel) {
			return fmt.Errorf("duplicate onRampMaxSeqNum for chain %d", seqNumChain.ChainSel)
		}
		seenChains.Add(seqNumChain.ChainSel)
	}

	return nil
}

// validateObservedOffRampMaxSeqNums TODO: doc
func (p *Plugin) validateObservedOffRampMaxSeqNums(
	offRampMaxSeqNums []plugintypes.SeqNumChain,
	observer commontypes.OracleID,
) error {
	if offRampMaxSeqNums == nil || len(offRampMaxSeqNums) == 0 {
		return nil
	}

	supportsDestChain, err := p.supportsDestChain(observer)
	if err != nil {
		return fmt.Errorf("call to supportsDestChain failed: %w", err)
	}

	if !supportsDestChain {
		return fmt.Errorf("observer %d does not support dest chain, but has observed offRampMaxSeqNums", observer)
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

// validateFChains TODO: doc
// FChain must not be empty
func (p *Plugin) validateFChain(fchain map[cciptypes.ChainSelector]int) error {
	if fchain == nil || len(fchain) == 0 {
		return fmt.Errorf("fchain map is empty")
	}

	seenChains := mapset.NewSet[cciptypes.ChainSelector]()
	for chainSel, f := range fchain {
		if seenChains.Contains(chainSel) {
			return fmt.Errorf("duplicate fChain for chain %d", chainSel)
		}

		if f < 0 {
			return fmt.Errorf("fChain %d is negative", f)
		}

		seenChains.Add(chainSel)
	}

	return nil
}
