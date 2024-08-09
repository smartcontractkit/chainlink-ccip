package commitrmnocb

import (
	"fmt"
	"sort"

	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"
	"golang.org/x/exp/maps"

	"github.com/smartcontractkit/libocr/offchainreporting2plus/types"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
)

// Outcome depending on the current state, either:
// - chooses the seq num ranges for the next round
// - builds a report
// - checks for the transmission of a previous report
func (p *Plugin) Outcome(
	outCtx ocr3types.OutcomeContext, query types.Query, aos []types.AttributedObservation,
) (ocr3types.Outcome, error) {
	previousOutcome, nextState := p.decodeOutcome(outCtx.PreviousOutcome)
	commitQuery := Query{}

	consensusObservation, err := p.getConsensusObservation(aos)
	if err != nil {
		return ocr3types.Outcome{}, err
	}

	outcome := Outcome{}

	switch nextState {
	case SelectingRangesForReport:
		outcome = p.ReportRangesOutcome(commitQuery, consensusObservation)

	case BuildingReport:
		outcome = p.buildReport(commitQuery, consensusObservation)

	case WaitingForReportTransmission:
		outcome = p.checkForReportTransmission(previousOutcome, consensusObservation)

	default:
		p.lggr.Warnw("Unexpected state in Outcome", "state", nextState)
		return outcome.Encode()
	}

	p.lggr.Infow("Commit Plugin Outcome", "outcome", outcome, "oid", p.nodeID)
	return outcome.Encode()
}

// ReportRangesOutcome determines the sequence number ranges for each chain to build a report from in the next round
func (p *Plugin) ReportRangesOutcome(
	query Query,
	consensusObservation ConsensusObservation,
) Outcome {
	rangesToReport := make([]ChainRange, 0)

	rmnOnRampMaxSeqNumsMap := make(map[cciptypes.ChainSelector]cciptypes.SeqNum)
	for _, seqNumChain := range query.RmnOnRampMaxSeqNums {
		rmnOnRampMaxSeqNumsMap[seqNumChain.ChainSel] = seqNumChain.SeqNum
	}

	observedOnRampMaxSeqNumsMap := consensusObservation.OnRampMaxSeqNums
	observedOffRampNextSeqNumsMap := consensusObservation.OffRampNextSeqNums

	for chainSel, offRampNextSeqNum := range observedOffRampNextSeqNumsMap {
		onRampMaxSeqNum, exists := observedOnRampMaxSeqNumsMap[chainSel]
		if !exists {
			continue
		}

		if rmnOnRampMaxSeqNum, exists := rmnOnRampMaxSeqNumsMap[chainSel]; exists {
			onRampMaxSeqNum = min(onRampMaxSeqNum, rmnOnRampMaxSeqNum)
		}

		if offRampNextSeqNum <= onRampMaxSeqNum {
			chainRange := ChainRange{
				ChainSel:    chainSel,
				SeqNumRange: [2]cciptypes.SeqNum{offRampNextSeqNum, onRampMaxSeqNum},
			}
			rangesToReport = append(rangesToReport, chainRange)
		}
	}

	// We sort here so that Outcome serializes deterministically
	sort.Slice(rangesToReport, func(i, j int) bool { return rangesToReport[i].ChainSel < rangesToReport[j].ChainSel })

	outcome := Outcome{
		OutcomeType:             ReportIntervalsSelected,
		RangesSelectedForReport: rangesToReport,
	}

	return outcome
}

// Given a set of observed merkle roots, gas prices and token prices, and roots from RMN, construct a report
// to transmit on-chain
func (p *Plugin) buildReport(
	_ Query,
	consensusObservation ConsensusObservation,
) Outcome {
	roots := maps.Values(consensusObservation.MerkleRoots)

	// We sort here so that Outcome serializes deterministically
	sort.Slice(roots, func(i, j int) bool {
		return roots[i].ChainSel < roots[j].ChainSel
	})

	outcomeType := ReportGenerated
	if len(roots) == 0 {
		outcomeType = ReportEmpty
	}

	outcome := Outcome{
		OutcomeType:   outcomeType,
		RootsToReport: roots,
		GasPrices:     consensusObservation.GasPricesSortedArray(),
		TokenPrices:   consensusObservation.TokenPricesSortedArray(),
	}

	return outcome
}

// checkForReportTransmission checks if the OffRamp has an updated set of max seq nums compared to the seq nums that
// were observed when the most recent report was generated. If an update to these max seq sums is detected, it means
// that the previous report has been transmitted, and we output ReportTransmitted to dictate that a new report
// generation phase should begin. If no update is detected, and we've exhausted our check attempts, output
// ReportTransmissionFailed to signify we stop checking for updates and start a new report generation phase. If no
// update is detected, and we haven't exhausted our check attempts, output ReportInFlight to signify that we should
// check again next round.
func (p *Plugin) checkForReportTransmission(
	previousOutcome Outcome,
	consensusObservation ConsensusObservation,
) Outcome {

	offRampUpdated := false
	for _, previousSeqNumChain := range previousOutcome.OffRampNextSeqNums {
		if currentSeqNum, exists := consensusObservation.OffRampNextSeqNums[previousSeqNumChain.ChainSel]; exists {
			if previousSeqNumChain.SeqNum != currentSeqNum {
				offRampUpdated = true
				break
			}
		}
	}

	if offRampUpdated {
		return Outcome{
			OutcomeType: ReportTransmitted,
		}
	}

	if previousOutcome.ReportTransmissionCheckAttempts+1 >= p.cfg.MaxReportTransmissionCheckAttempts {
		p.lggr.Warnw("Failed to detect report transmission")
		return Outcome{
			OutcomeType: ReportTransmissionFailed,
		}
	}

	return Outcome{
		OutcomeType:                     ReportInFlight,
		OffRampNextSeqNums:              previousOutcome.OffRampNextSeqNums,
		ReportTransmissionCheckAttempts: previousOutcome.ReportTransmissionCheckAttempts + 1,
	}
}

// getConsensusObservation Combine the list of observations into a single consensus observation
func (p *Plugin) getConsensusObservation(aos []types.AttributedObservation) (ConsensusObservation, error) {
	aggObs := aggregateObservations(aos)
	fChains := p.fChainConsensus(aggObs.FChain)

	fDestChain, exists := fChains[p.cfg.DestChain]
	if !exists {
		return ConsensusObservation{},
			fmt.Errorf("no consensus value for fDestChain, DestChain: %d", p.cfg.DestChain)
	}

	consensusObs := ConsensusObservation{
		MerkleRoots: p.merkleRootConsensus(aggObs.MerkleRoots, fChains),
		// TODO: use consensus of observed gas prices
		GasPrices: make(map[cciptypes.ChainSelector]cciptypes.BigInt),
		// TODO: use consensus of observed token prices
		TokenPrices:        make(map[types.Account]cciptypes.BigInt),
		OnRampMaxSeqNums:   p.onRampMaxSeqNumsConsensus(aggObs.OnRampMaxSeqNums, fChains),
		OffRampNextSeqNums: p.offRampMaxSeqNumsConsensus(aggObs.OffRampNextSeqNums, fDestChain),
		FChain:             fChains,
	}

	return consensusObs, nil
}

// Given a mapping from chains to a list of merkle roots, return a mapping from chains to a single consensus merkle
// root. The consensus merkle root for a given chain is the merkle root with the most observations that was observed at
// least fChain times.
func (p *Plugin) merkleRootConsensus(
	rootsByChain map[cciptypes.ChainSelector][]cciptypes.MerkleRootChain,
	fChains map[cciptypes.ChainSelector]int,
) map[cciptypes.ChainSelector]cciptypes.MerkleRootChain {
	consensus := make(map[cciptypes.ChainSelector]cciptypes.MerkleRootChain)

	for chain, roots := range rootsByChain {
		if f, exists := fChains[chain]; exists {
			root, count := mostFrequentElem(roots)

			if count <= f {
				// TODO: metrics
				p.lggr.Warnf("failed to reach consensus on a merkle root for chain %d "+
					"because no single merkle root was observed more than the expected %d times, found merkle root %d "+
					"observed by only %d oracles, all observed merkle roots: %v",
					chain, f, root, count, roots)
			}

			consensus[chain] = root
		} else {
			// TODO: metrics
			p.lggr.Warnf("merkleRootConsensus: fChain not found for chain %d", chain)
		}
	}

	return consensus
}

// Given a mapping from chains to a list of max seq nums on their corresponding OnRamp, return a mapping from chains
// to a single max seq num. The consensus max seq num for a given chain is the f'th lowest max seq num if the number
// of max seq num observations is greater or equal than 2f+1, where f is the FChain of the corresponding source chain.
func (p *Plugin) onRampMaxSeqNumsConsensus(
	onRampMaxSeqNumsByChain map[cciptypes.ChainSelector][]cciptypes.SeqNum,
	fChains map[cciptypes.ChainSelector]int,
) map[cciptypes.ChainSelector]cciptypes.SeqNum {
	consensus := make(map[cciptypes.ChainSelector]cciptypes.SeqNum)

	for chain, onRampMaxSeqNums := range onRampMaxSeqNumsByChain {
		if f, exists := fChains[chain]; exists {
			if len(onRampMaxSeqNums) < 2*f+1 {
				// TODO: metrics
				p.lggr.Warnf("could not reach consensus on onRampMaxSeqNums for chain %d "+
					"because we did not receive more than 2f+1 observed sequence numbers, 2f+1: %d, "+
					"len(onRampMaxSeqNums): %d, onRampMaxSeqNums: %v",
					chain, 2*f+1, len(onRampMaxSeqNums), onRampMaxSeqNums)
			} else {
				sort.Slice(onRampMaxSeqNums, func(i, j int) bool { return onRampMaxSeqNums[i] < onRampMaxSeqNums[j] })
				consensus[chain] = onRampMaxSeqNums[f]
			}
		} else {
			// TODO: metrics
			p.lggr.Warnf("could not reach consensus on onRampMaxSeqNums for chain %d "+
				"because there was no consensus f value for this chain", chain)
		}
	}

	return consensus
}

// Given a mapping from chains to a list of max seq nums on the OffRamp, return a mapping from chains
// to a single max seq num. The consensus max seq num for a given chain is the max seq num with the most observations
// that was observed at least f times, where f is the FChain of the dest chain.
func (p *Plugin) offRampMaxSeqNumsConsensus(
	offRampMaxSeqNumsByChain map[cciptypes.ChainSelector][]cciptypes.SeqNum,
	fDestChain int,
) map[cciptypes.ChainSelector]cciptypes.SeqNum {
	consensus := make(map[cciptypes.ChainSelector]cciptypes.SeqNum)

	for chain, offRampMaxSeqNums := range offRampMaxSeqNumsByChain {
		seqNum, count := mostFrequentElem(offRampMaxSeqNums)
		if count <= fDestChain {
			// TODO: metrics
			p.lggr.Warnf("could not reach consensus on offRampMaxSeqNums for chain %d "+
				"because we did not receive a sequence number that was observed by at least f (%d) oracles, "+
				"offRampMaxSeqNums: %v", chain, fDestChain, offRampMaxSeqNums)
		} else {
			consensus[chain] = seqNum
		}
	}

	return consensus
}

// Given a mapping from chains to a list of FChain values for each chain, return a mapping from chains
// to a single FChain. The consensus FChain for a given chain is the FChain with the most observations
// that was observed at least f times, where f is the F of the DON (p.reportingCfg.F).
func (p *Plugin) fChainConsensus(fChainValues map[cciptypes.ChainSelector][]int) map[cciptypes.ChainSelector]int {
	consensus := make(map[cciptypes.ChainSelector]int)

	for chain, fValues := range fChainValues {
		f, _ := mostFrequentElem(fValues)
		// TODO: uncomment when p.reportingCfg is added back
		//if count < p.reportingCfg.F {
		//	// TODO: metrics
		//	p.lggr.Warnf("failed to reach consensus on fChain values for chain %d because no single f "+
		//		"value was observed more than the expected %d times, found f value %d observed by only %d oracles, "+
		//		"f values: %v",
		//		chain, p.reportingCfg.F, f, count, fValues)
		//}

		consensus[chain] = f
	}

	return consensus
}

// Given a list of elems, return the elem that occurs most frequently and how often it occurs
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

// Given a list of elems, return a map from elems to how often they occur in the given list
func counts[T comparable](elems []T) map[T]int {
	m := make(map[T]int)
	for _, elem := range elems {
		m[elem]++
	}

	return m
}
