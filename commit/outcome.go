package commit

import (
	"fmt"
	"sort"

	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"
	"golang.org/x/exp/maps"

	"github.com/smartcontractkit/libocr/offchainreporting2plus/types"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-ccip/plugintypes"
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

	consensusObservation, err := getConsensusObservation(p.lggr, p.reportingCfg.F, p.cfg.DestChain, aos)
	if err != nil {
		return ocr3types.Outcome{}, err
	}

	outcome := Outcome{}

	switch nextState {
	case SelectingRangesForReport:
		outcome = ReportRangesOutcome(commitQuery, consensusObservation)

	case BuildingReport:
		outcome = buildReport(commitQuery, consensusObservation, previousOutcome)

	case WaitingForReportTransmission:
		outcome = checkForReportTransmission(
			p.lggr, p.cfg.MaxReportTransmissionCheckAttempts, previousOutcome, consensusObservation)

	default:
		p.lggr.Warnw("Unexpected state in Outcome", "state", nextState)
		return outcome.Encode()
	}

	p.lggr.Infow("Commit Plugin Outcome", "outcome", outcome, "oid", p.nodeID)
	return outcome.Encode()
}

// ReportRangesOutcome determines the sequence number ranges for each chain to build a report from in the next round
// TODO: ensure each range is below a limit
func ReportRangesOutcome(
	query Query,
	consensusObservation ConsensusObservation,
) Outcome {
	rangesToReport := make([]plugintypes.ChainRange, 0)

	rmnOnRampMaxSeqNumsMap := make(map[cciptypes.ChainSelector]cciptypes.SeqNum)
	for _, seqNumChain := range query.RmnOnRampMaxSeqNums {
		rmnOnRampMaxSeqNumsMap[seqNumChain.ChainSel] = seqNumChain.SeqNum
	}

	observedOnRampMaxSeqNumsMap := consensusObservation.OnRampMaxSeqNums
	observedOffRampNextSeqNumsMap := consensusObservation.OffRampNextSeqNums
	offRampNextSeqNums := make([]plugintypes.SeqNumChain, 0)

	for chainSel, offRampNextSeqNum := range observedOffRampNextSeqNumsMap {
		onRampMaxSeqNum, exists := observedOnRampMaxSeqNumsMap[chainSel]
		if !exists {
			continue
		}

		if rmnOnRampMaxSeqNum, exists := rmnOnRampMaxSeqNumsMap[chainSel]; exists {
			onRampMaxSeqNum = min(onRampMaxSeqNum, rmnOnRampMaxSeqNum)
		}

		if offRampNextSeqNum <= onRampMaxSeqNum {
			chainRange := plugintypes.ChainRange{
				ChainSel:    chainSel,
				SeqNumRange: [2]cciptypes.SeqNum{offRampNextSeqNum, onRampMaxSeqNum},
			}
			rangesToReport = append(rangesToReport, chainRange)
		}

		offRampNextSeqNums = append(offRampNextSeqNums, plugintypes.SeqNumChain{
			ChainSel: chainSel,
			SeqNum:   offRampNextSeqNum,
		})
	}

	// deterministic outcome
	sort.Slice(rangesToReport, func(i, j int) bool { return rangesToReport[i].ChainSel < rangesToReport[j].ChainSel })
	sort.Slice(offRampNextSeqNums, func(i, j int) bool {
		return offRampNextSeqNums[i].ChainSel < offRampNextSeqNums[j].ChainSel
	})

	outcome := Outcome{
		OutcomeType:             ReportIntervalsSelected,
		RangesSelectedForReport: rangesToReport,
		OffRampNextSeqNums:      offRampNextSeqNums,
	}

	return outcome
}

// Given a set of observed merkle roots, gas prices and token prices, and roots from RMN, construct a report
// to transmit on-chain
func buildReport(
	_ Query,
	consensusObservation ConsensusObservation,
	prevOutcome Outcome,
) Outcome {
	roots := maps.Values(consensusObservation.MerkleRoots)

	outcomeType := ReportGenerated
	if len(roots) == 0 {
		outcomeType = ReportEmpty
	}

	sort.Slice(roots, func(i, j int) bool { return roots[i].ChainSel < roots[j].ChainSel })

	outcome := Outcome{
		OutcomeType:        outcomeType,
		RootsToReport:      roots,
		GasPrices:          consensusObservation.GasPricesArray(),
		TokenPrices:        consensusObservation.TokenPricesArray(),
		OffRampNextSeqNums: prevOutcome.OffRampNextSeqNums,
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
func checkForReportTransmission(
	lggr logger.Logger,
	maxReportTransmissionCheckAttempts uint,
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

	if previousOutcome.ReportTransmissionCheckAttempts+1 >= maxReportTransmissionCheckAttempts {
		lggr.Warnw("Failed to detect report transmission")
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
func getConsensusObservation(
	lggr logger.Logger,
	F int,
	destChain cciptypes.ChainSelector,
	aos []types.AttributedObservation,
) (ConsensusObservation, error) {
	aggObs := aggregateObservations(aos)
	fChains := fChainConsensus(lggr, F, aggObs.FChain)

	fDestChain, exists := fChains[destChain]
	if !exists {
		return ConsensusObservation{},
			fmt.Errorf("no consensus value for fDestChain, destChain: %d", destChain)
	}

	consensusObs := ConsensusObservation{
		MerkleRoots: merkleRootConsensus(lggr, aggObs.MerkleRoots, fChains),
		// TODO: use consensus of observed gas prices
		GasPrices: make(map[cciptypes.ChainSelector]cciptypes.BigInt),
		// TODO: use consensus of observed token prices
		TokenPrices:        make(map[types.Account]cciptypes.BigInt),
		OnRampMaxSeqNums:   onRampMaxSeqNumsConsensus(lggr, aggObs.OnRampMaxSeqNums, fChains),
		OffRampNextSeqNums: offRampMaxSeqNumsConsensus(lggr, aggObs.OffRampNextSeqNums, fDestChain),
		FChain:             fChains,
	}

	return consensusObs, nil
}

// Given a mapping from chains to a list of merkle roots, return a mapping from chains to a single consensus merkle
// root. The consensus merkle root for a given chain is the merkle root with the most observations that was observed at
// least fChain times.
func merkleRootConsensus(
	lggr logger.Logger,
	rootsByChain map[cciptypes.ChainSelector][]cciptypes.MerkleRootChain,
	fChains map[cciptypes.ChainSelector]int,
) map[cciptypes.ChainSelector]cciptypes.MerkleRootChain {
	consensus := make(map[cciptypes.ChainSelector]cciptypes.MerkleRootChain)

	for chain, roots := range rootsByChain {
		if fChain, exists := fChains[chain]; exists {
			root, count := majorElem(roots)

			if count <= fChain {
				// TODO: metrics
				lggr.Warnf("failed to reach consensus on a merkle root for chain %d "+
					"because no single merkle root was observed more than the expected fChain (%d) times, found "+
					"merkle root %d observed by only %d oracles, all observed merkle roots: %v",
					chain, fChain, root, count, roots)
			} else {
				consensus[chain] = root
			}
		} else {
			// TODO: metrics
			lggr.Warnf("merkleRootConsensus: fChain not found for chain %d", chain)
		}
	}

	return consensus
}

// Given a mapping from chains to a list of max seq nums on their corresponding OnRamp, return a mapping from chains
// to a single max seq num. The consensus max seq num for a given chain is the f'th lowest max seq num if the number
// of max seq num observations is greater or equal than 2f+1, where f is the FChain of the corresponding source chain.
func onRampMaxSeqNumsConsensus(
	lggr logger.Logger,
	onRampMaxSeqNumsByChain map[cciptypes.ChainSelector][]cciptypes.SeqNum,
	fChains map[cciptypes.ChainSelector]int,
) map[cciptypes.ChainSelector]cciptypes.SeqNum {
	consensus := make(map[cciptypes.ChainSelector]cciptypes.SeqNum)

	for chain, onRampMaxSeqNums := range onRampMaxSeqNumsByChain {
		if fChain, exists := fChains[chain]; exists {
			if len(onRampMaxSeqNums) < 2*fChain+1 {
				// TODO: metrics
				lggr.Warnf("could not reach consensus on onRampMaxSeqNums for chain %d "+
					"because we did not receive more than 2fChain+1 observed sequence numbers, 2fChain+1: %d, "+
					"len(onRampMaxSeqNums): %d, onRampMaxSeqNums: %v",
					chain, 2*fChain+1, len(onRampMaxSeqNums), onRampMaxSeqNums)
			} else {
				sort.Slice(onRampMaxSeqNums, func(i, j int) bool { return onRampMaxSeqNums[i] < onRampMaxSeqNums[j] })
				consensus[chain] = onRampMaxSeqNums[fChain]
			}
		} else {
			// TODO: metrics
			lggr.Warnf("could not reach consensus on onRampMaxSeqNums for chain %d "+
				"because there was no consensus fChain value for this chain", chain)
		}
	}

	return consensus
}

// Given a mapping from chains to a list of max seq nums on the OffRamp, return a mapping from chains
// to a single max seq num. The consensus max seq num for a given chain is the max seq num with the most observations
// that was observed at least f times, where f is the FChain of the dest chain.
func offRampMaxSeqNumsConsensus(
	lggr logger.Logger,
	offRampMaxSeqNumsByChain map[cciptypes.ChainSelector][]cciptypes.SeqNum,
	fDestChain int,
) map[cciptypes.ChainSelector]cciptypes.SeqNum {
	consensus := make(map[cciptypes.ChainSelector]cciptypes.SeqNum)

	for chain, offRampMaxSeqNums := range offRampMaxSeqNumsByChain {
		seqNum, count := majorElem(offRampMaxSeqNums)
		if count <= fDestChain {
			// TODO: metrics
			lggr.Warnf("could not reach consensus on offRampMaxSeqNums for chain %d "+
				"because we did not receive a sequence number that was observed by at least fChain (%d) oracles, "+
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
func fChainConsensus(
	lggr logger.Logger,
	F int,
	fChainValues map[cciptypes.ChainSelector][]int,
) map[cciptypes.ChainSelector]int {
	consensus := make(map[cciptypes.ChainSelector]int)

	for chain, fValues := range fChainValues {
		fChain, count := majorElem(fValues)
		if count < F {
			// TODO: metrics
			lggr.Warnf("failed to reach consensus on fChain values for chain %d because no single fChain "+
				"value was observed more than the expected %d times, found fChain value %d observed by only %d oracles, "+
				"fChain values: %v",
				chain, F, fChain, count, fValues)
			continue
		}

		consensus[chain] = fChain
	}

	return consensus
}

// Given a list of elems, return the elem that occurs most frequently and how often it occurs.
// If there is no major element (e.g. elements appear the same number of times), an empty value and zero is returned.
func majorElem[T comparable](elems []T) (T, int) {
	var mostFrequentElem T

	counts := counts(elems)
	maxCount := 0
	isMajor := false

	for elem, count := range counts {
		if count == maxCount {
			isMajor = false
		}
		if count > maxCount {
			mostFrequentElem = elem
			maxCount = count
			isMajor = true
		}
	}

	if !isMajor {
		var empty T
		return empty, 0
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
