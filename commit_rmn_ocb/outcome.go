package commitrmnocb

import (
	"fmt"
	"sort"

	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"

	"github.com/smartcontractkit/libocr/offchainreporting2plus/types"

	"github.com/smartcontractkit/chainlink-ccip/internal/libs/slicelib"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
)

// Outcome TODO: doc
func (p *Plugin) Outcome(
	outCtx ocr3types.OutcomeContext, query types.Query, aos []types.AttributedObservation,
) (ocr3types.Outcome, error) {
	previousOutcome, nextState := p.decodeOutcome(outCtx.PreviousOutcome)
	commitQuery, err := DecodeCommitPluginQuery(query)
	if err != nil {
		return ocr3types.Outcome{}, err
	}

	consensusObservation, err := p.getConsensusObservation(aos)
	if err != nil {
		return ocr3types.Outcome{}, err
	}

	switch nextState {
	case SelectingRangesForReport:
		return p.ReportRangesOutcome(commitQuery, consensusObservation)

	case BuildingReport:
		return p.buildReport(commitQuery, consensusObservation)

	case WaitingForReportTransmission:
		return p.checkForReportTransmission(previousOutcome, consensusObservation)

	default:
		return nil, fmt.Errorf("outcome unexpected state: %d", nextState)
	}
}

// ReportRangesOutcome TODO: doc
func (p *Plugin) ReportRangesOutcome(
	query CommitQuery,
	consensusObservation ConsensusObservation,
) (ocr3types.Outcome, error) {
	rangesToReport := make([]ChainRange, 0)

	rmnOnRampMaxSeqNumsMap := make(map[cciptypes.ChainSelector]cciptypes.SeqNum)
	for _, seqNumChain := range query.RmnOnRampMaxSeqNums {
		rmnOnRampMaxSeqNumsMap[seqNumChain.ChainSel] = seqNumChain.SeqNum
	}

	observedOnRampMaxSeqNumsMap := consensusObservation.OnRampMaxSeqNums
	observedOffRampMaxSeqNumsMap := consensusObservation.OffRampMaxSeqNums

	for chainSel, offRampMaxSeqNum := range observedOffRampMaxSeqNumsMap {
		onRampMaxSeqNum, exists := observedOnRampMaxSeqNumsMap[chainSel]
		if !exists {
			continue
		}

		if rmnOnRampMaxSeqNum, exists := rmnOnRampMaxSeqNumsMap[chainSel]; exists {
			onRampMaxSeqNum = min(onRampMaxSeqNum, rmnOnRampMaxSeqNum)
		}

		if offRampMaxSeqNum > onRampMaxSeqNum {
			// TODO: metrics
			p.log.Warnw("Found an offRampMaxSeqNum greater than an onRampMaxSeqNum",
				"offRampMaxSeqNum", offRampMaxSeqNum,
				"onRampMaxSeqNum", onRampMaxSeqNum,
				"chainSelector", chainSel)
		}

		if offRampMaxSeqNum < onRampMaxSeqNum {
			chainRange := ChainRange{
				ChainSel:    chainSel,
				SeqNumRange: [2]cciptypes.SeqNum{offRampMaxSeqNum, onRampMaxSeqNum},
			}
			rangesToReport = append(rangesToReport, chainRange)
		}
	}

	// TODO: explain this
	sort.Slice(rangesToReport, func(i, j int) bool { return rangesToReport[i].ChainSel < rangesToReport[j].ChainSel })

	outcome := CommitPluginOutcome{
		OutcomeType:             ReportIntervalsSelected,
		RangesSelectedForReport: rangesToReport,
	}

	return outcome.Encode()
}

// TODO: doc, break apart
func (p *Plugin) buildReport(
	query CommitQuery,
	consensusObservation ConsensusObservation,
) (ocr3types.Outcome, error) {
	verifiedSignedRoots := p.verifySignedRoots(query.SignedMerkleRoots, consensusObservation.MerkleRoots)

	chainsToExclude := make([]cciptypes.ChainSelector, 0)
	for chain := range verifiedSignedRoots {
		if _, exists := consensusObservation.GasPrices[chain]; !exists {
			// TODO: metrics
			p.log.Warnw(
				"did not find a consensus gas price for chain %d, excluding it from the report", chain)
			chainsToExclude = append(chainsToExclude, chain)
		}
	}

	for _, chainSelector := range chainsToExclude {
		delete(verifiedSignedRoots, chainSelector)
	}

	// TODO: token prices validation
	// exclude merkle roots if expected token prices don't exist

	verifiedSignedRootsArray := make([]SignedMerkleRoot, 0, len(verifiedSignedRoots))
	for _, signedRoot := range verifiedSignedRoots {
		verifiedSignedRootsArray = append(verifiedSignedRootsArray, signedRoot)
	}

	// TODO: explain this
	sort.Slice(verifiedSignedRootsArray, func(i, j int) bool {
		return verifiedSignedRootsArray[i].chain() < verifiedSignedRootsArray[j].chain()
	})

	return CommitPluginOutcome{
		OutcomeType:         ReportGenerated,
		SignedRootsToReport: verifiedSignedRootsArray,
		GasPrices:           consensusObservation.GasPricesSortedArray(),
		TokenPrices:         consensusObservation.TokenPricesSortedArray(),
	}.Encode()
}

// TODO: doc, break apart
func (p *Plugin) verifySignedRoots(
	rmnSignedRoots []SignedMerkleRoot,
	observedRoots map[cciptypes.ChainSelector]MerkleRoot,
) map[cciptypes.ChainSelector]SignedMerkleRoot {
	verifiedSignedRoots := make(map[cciptypes.ChainSelector]SignedMerkleRoot)
	for _, signedRoot := range rmnSignedRoots {
		if err := p.rmn.VerifySignedMerkleRoot(signedRoot); err != nil {
			// TODO: metrics
			p.log.Warnw("failed to verify signed merkle root",
				"err", err,
				"signedRoot", signedRoot)
			continue
		}

		if observedMerkleRoot, exists := observedRoots[signedRoot.chain()]; exists {
			if observedMerkleRoot == signedRoot.MerkleRoot {
				verifiedSignedRoots[signedRoot.chain()] = signedRoot
			} else {
				// TODO: metrics
				p.log.Warnw("observed merkle root does not match merkle root received from RMN",
					"rmnSignedRoot", signedRoot,
					"observedMerkleRoot", observedMerkleRoot)
			}
		} else {
			// TODO: metrics
			p.log.Warnw(
				"received a signed merkle root from RMN for a chain, but did not observe a merkle root for "+
					"this chain",
				"rmnSignedRoot", signedRoot)
		}
	}

	// TODO: explain this
	for chain, observedMerkleRoot := range observedRoots {
		if _, exists := verifiedSignedRoots[chain]; !exists {
			if p.rmn.ChainThreshold(chain) == 0 {
				verifiedSignedRoots[chain] = SignedMerkleRoot{
					MerkleRoot: observedMerkleRoot,
					RmnSigs:    []RmnSig{},
				}
			} else {
				// TODO: metrics
				p.log.Warnw(
					"did not receive RMN signatures for chain %d that requires %d RMN signatures, "+
						"MerkleRoot: %v", chain, p.rmn.ChainThreshold(chain), observedMerkleRoot)
			}
		}
	}

	return verifiedSignedRoots
}

// TODO: doc
func (p *Plugin) checkForReportTransmission(
	previousOutcome CommitPluginOutcome,
	consensusObservation ConsensusObservation,
) (ocr3types.Outcome, error) {

	offRampUpdated := false
	for _, previousSeqNumChain := range previousOutcome.OffRampMaxSeqNums {
		if currentSeqNum, exists := consensusObservation.OffRampMaxSeqNums[previousSeqNumChain.ChainSel]; exists {
			if previousSeqNumChain.SeqNum != currentSeqNum {
				offRampUpdated = true
				break
			}
		}
	}

	if offRampUpdated {
		return CommitPluginOutcome{
			OutcomeType: ReportGenerated,
		}.Encode()
	}

	if previousOutcome.ReportTransmissionCheckAttempts+1 >= p.cfg.MaxReportTransmissionCheckAttempts {
		return CommitPluginOutcome{
			OutcomeType: ReportNotTransmitted,
		}.Encode()
	}

	return CommitPluginOutcome{
		OutcomeType:                     ReportNotYetTransmitted,
		OffRampMaxSeqNums:               previousOutcome.OffRampMaxSeqNums,
		ReportTransmissionCheckAttempts: previousOutcome.ReportTransmissionCheckAttempts + 1,
	}.Encode()
}

// getConsensusObservation TODO: doc
func (p *Plugin) getConsensusObservation(aos []types.AttributedObservation) (ConsensusObservation, error) {
	aggObs := aggregateObservations(aos)
	fChains := p.fChainConsensus(aggObs.FChain)

	fDestChain, exists := fChains[p.cfg.DestChain]
	if !exists {
		return ConsensusObservation{}, fmt.Errorf("no consensus value for fDestChain, DestChain: %d", p.cfg.DestChain)
	}

	fTokenChain, exists := fChains[p.cfg.TokenPriceChain]
	if !exists {
		return ConsensusObservation{}, fmt.Errorf("no consensus value for fTokenChain, TokenPriceChain: %d", p.cfg.TokenPriceChain)
	}

	consensusObs := ConsensusObservation{
		MerkleRoots:       p.merkleRootConsensus(aggObs.MerkleRoots, fChains),
		GasPrices:         p.gasPriceConsensus(aggObs.GasPrices, fChains),
		TokenPrices:       p.tokenPriceConsensus(aggObs.TokenPrices, fTokenChain),
		OnRampMaxSeqNums:  p.onRampMaxSeqNumsConsensus(aggObs.OnRampMaxSeqNums, fChains),
		OffRampMaxSeqNums: p.offRampMaxSeqNumsConsensus(aggObs.OffRampMaxSeqNums, fDestChain),
		FChain:            fChains,
	}

	return consensusObs, nil
}

func (p *Plugin) merkleRootConsensus(
	rootsByChain map[cciptypes.ChainSelector][]MerkleRoot,
	fChains map[cciptypes.ChainSelector]int,
) map[cciptypes.ChainSelector]MerkleRoot {
	consensus := make(map[cciptypes.ChainSelector]MerkleRoot)

	for chain, roots := range rootsByChain {
		if f, exists := fChains[chain]; exists {
			root, count := mostFrequentElem(roots)

			if count <= f {
				// TODO: metrics
				p.log.Warnf("failed to reach consensus on a merkle root for chain %d "+
					"because no single merkle root was observed more than the expected %d times, found merkle root %d "+
					"observed by only %d oracles, all observed merkle roots: %v",
					chain, f, root, count, roots)
			}

			consensus[chain] = root
		} else {
			// TODO: metrics
			p.log.Warnf("merkleRootConsensus: fChain not found for chain %d", chain)
		}
	}

	return consensus
}

func (p *Plugin) gasPriceConsensus(
	pricesByChain map[cciptypes.ChainSelector][]cciptypes.BigInt,
	fChains map[cciptypes.ChainSelector]int,
) map[cciptypes.ChainSelector]cciptypes.BigInt {
	consensus := make(map[cciptypes.ChainSelector]cciptypes.BigInt)

	for chain, prices := range pricesByChain {
		if f, exists := fChains[chain]; exists {
			if len(prices) < 2*f+1 {
				// TODO: metrics
				p.log.Warnf("could not reach consensus on gas prices for chain %d "+
					"because we did not receive more than 2f+1 observed prices, 2f+1: %d, len(prices): %d, prices: %v",
					chain, 2*f+1, len(prices), prices)
			}

			consensus[chain] = slicelib.BigIntSortedMiddle(prices)
		} else {
			// TODO: metrics
			p.log.Warnf("could not reach consensus on gas prices for chain %d because "+
				"there was no consensus f value for this chain", chain)
		}
	}

	return consensus
}

func (p *Plugin) tokenPriceConsensus(
	pricesByToken map[types.Account][]cciptypes.BigInt,
	fTokenChain int,
) map[types.Account]cciptypes.BigInt {
	consensus := make(map[types.Account]cciptypes.BigInt)

	for tokenID, prices := range pricesByToken {
		if len(prices) < 2*fTokenChain+1 {
			// TODO: metrics
			p.log.Warnf("could not reach consensus on token prices for token %s because "+
				"we did not receive more than 2f+1 observed prices, 2f+1: %d, len(prices): %d, prices: %v",
				tokenID, 2*fTokenChain+1, len(prices), prices)
		}

		consensus[tokenID] = slicelib.BigIntSortedMiddle(prices)
	}

	return consensus
}

func (p *Plugin) onRampMaxSeqNumsConsensus(
	onRampMaxSeqNumsByChain map[cciptypes.ChainSelector][]cciptypes.SeqNum,
	fChains map[cciptypes.ChainSelector]int,
) map[cciptypes.ChainSelector]cciptypes.SeqNum {
	consensus := make(map[cciptypes.ChainSelector]cciptypes.SeqNum)

	for chain, onRampMaxSeqNums := range onRampMaxSeqNumsByChain {
		if f, exists := fChains[chain]; exists {
			if len(onRampMaxSeqNums) < 2*f+1 {
				// TODO: metrics
				p.log.Warnf("could not reach consensus on onRampMaxSeqNums for chain %d "+
					"because we did not receive more than 2f+1 observed sequence numbers, 2f+1: %d, "+
					"len(onRampMaxSeqNums): %d, onRampMaxSeqNums: %v",
					chain, 2*f+1, len(onRampMaxSeqNums), onRampMaxSeqNums)
			} else {
				sort.Slice(onRampMaxSeqNums, func(i, j int) bool { return onRampMaxSeqNums[i] < onRampMaxSeqNums[j] })
				consensus[chain] = onRampMaxSeqNums[f]
			}
		} else {
			// TODO: metrics
			p.log.Warnf("could not reach consensus on onRampMaxSeqNums for chain %d "+
				"because there was no consensus f value for this chain", chain)
		}
	}

	return consensus
}

func (p *Plugin) offRampMaxSeqNumsConsensus(
	offRampMaxSeqNumsByChain map[cciptypes.ChainSelector][]cciptypes.SeqNum,
	fDestChain int,
) map[cciptypes.ChainSelector]cciptypes.SeqNum {
	consensus := make(map[cciptypes.ChainSelector]cciptypes.SeqNum)

	for chain, offRampMaxSeqNums := range offRampMaxSeqNumsByChain {
		seqNum, count := mostFrequentElem(offRampMaxSeqNums)
		if count <= fDestChain {
			// TODO: metrics
			p.log.Warnf("could not reach consensus on offRampMaxSeqNums for chain %d "+
				"because we did not receive a sequence number that was observed by at least f (%d) oracles, "+
				"offRampMaxSeqNums: %v", chain, fDestChain, offRampMaxSeqNums)
		} else {
			consensus[chain] = seqNum
		}
	}

	return consensus
}

// fChainConsensus TODO: doc
func (p *Plugin) fChainConsensus(fChainValues map[cciptypes.ChainSelector][]int) map[cciptypes.ChainSelector]int {
	consensus := make(map[cciptypes.ChainSelector]int)

	for chain, fValues := range fChainValues {
		f, count := mostFrequentElem(fValues)
		if count < p.reportingCfg.F {
			// TODO: metrics
			p.log.Warnf("failed to reach consensus on fChain values for chain %d because no single f "+
				"value was observed more than the expected %d times, found f value %d observed by only %d oracles, "+
				"f values: %v",
				chain, p.reportingCfg.F, f, count, fValues)
		}

		consensus[chain] = f
	}

	return consensus
}

// mostFrequentElem TODO: doc
func mostFrequentElem[T comparable](elems []T) (T, int) {
	var mostFrequentElem T

	counts := counts(elems)
	maxCount := 0

	for elem, count := range counts {
		if count > maxCount {
			mostFrequentElem = elem
			maxCount = count
		}
	}

	return mostFrequentElem, maxCount
}

// counts TODO: doc
func counts[T comparable](elems []T) map[T]int {
	m := make(map[T]int)
	for _, elem := range elems {
		m[elem]++
	}

	return m
}
