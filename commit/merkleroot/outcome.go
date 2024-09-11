package merkleroot

import (
	"fmt"
	"sort"
	"time"

	ct "github.com/smartcontractkit/chainlink-ccip/commit/committypes"
	"golang.org/x/exp/maps"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-ccip/plugintypes"
	"github.com/smartcontractkit/chainlink-ccip/shared"
)

// Outcome ct.MerkleRootOutcome depending on the current state, either:
// - chooses the seq num ranges for the next round
// - builds a report
// - checks for the transmission of a previous report
func (w *Processor) Outcome(
	prevOutcome ct.Outcome,
	query ct.Query,
	aos []shared.AttributedObservation[ct.Observation],
) (ct.Outcome, error) {
	tStart := time.Now()
	merkleAos := mapAttributedObs(aos)
	outcome, nextState := w.getOutcome(prevOutcome.MerkleRootOutcome, query.MerkleRootQuery, merkleAos)
	w.lggr.Infow("Sending ct.MerkleRootOutcome",
		"outcome", outcome, "oid", w.oracleID, "nextState", nextState, "outcomeDuration", time.Since(tStart))
	return ct.Outcome{MerkleRootOutcome: outcome}, nil
}

func (w *Processor) getOutcome(
	previousOutcome ct.MerkleRootOutcome,
	q ct.MerkleRootQuery,
	aos []shared.AttributedObservation[ct.MerkleRootObservation],
) (ct.MerkleRootOutcome, ct.MerkleRootState) {
	nextState := previousOutcome.NextState()

	consensusObservation, err := getConsensusObservation(w.lggr, w.reportingCfg.F, w.cfg.DestChain, aos)
	if err != nil {
		w.lggr.Warnw("Get consensus observation failed, empty outcome", "error", err)
		return ct.MerkleRootOutcome{}, nextState
	}

	switch nextState {
	case ct.SelectingRangesForReport:
		return reportRangesOutcome(q, consensusObservation), nextState
	case ct.BuildingReport:
		if q.RetryRMNSignatures {
			// We want to retry getting the RMN signatures on the exact same outcome we had before.
			// The current observations should all be empty.
			return previousOutcome, ct.BuildingReport
		}
		return buildReport(q, consensusObservation, previousOutcome), nextState
	case ct.WaitingForReportTransmission:
		return checkForReportTransmission(
			w.lggr, w.cfg.MaxReportTransmissionCheckAttempts, previousOutcome, consensusObservation), nextState
	default:
		w.lggr.Warnw("Unexpected next state in MerkleRootOutcome", "state", nextState)
		return ct.MerkleRootOutcome{}, nextState
	}
}

// reportRangesOutcome determines the sequence number ranges for each chain to build a report from in the next round
// TODO: ensure each range is below a limit
func reportRangesOutcome(
	_ ct.MerkleRootQuery,
	consensusObservation ConsensusObservation,
) ct.MerkleRootOutcome {
	rangesToReport := make([]plugintypes.ChainRange, 0)

	observedOnRampMaxSeqNumsMap := consensusObservation.OnRampMaxSeqNums
	observedOffRampNextSeqNumsMap := consensusObservation.OffRampNextSeqNums
	offRampNextSeqNums := make([]plugintypes.SeqNumChain, 0)

	for chainSel, offRampNextSeqNum := range observedOffRampNextSeqNumsMap {
		onRampMaxSeqNum, exists := observedOnRampMaxSeqNumsMap[chainSel]
		if !exists {
			continue
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

	outcome := ct.MerkleRootOutcome{
		OutcomeType:             ct.ReportIntervalsSelected,
		RangesSelectedForReport: rangesToReport,
		OffRampNextSeqNums:      offRampNextSeqNums,
	}

	return outcome
}

// Given a set of observed merkle roots, gas prices and token prices, and roots from RMN, construct a report
// to transmit on-chain
func buildReport(
	_ ct.MerkleRootQuery,
	consensusObservation ConsensusObservation,
	prevOutcome ct.MerkleRootOutcome,
) ct.MerkleRootOutcome {
	roots := maps.Values(consensusObservation.MerkleRoots)

	outcomeType := ct.ReportGenerated
	if len(roots) == 0 {
		outcomeType = ct.ReportEmpty
	}

	sort.Slice(roots, func(i, j int) bool { return roots[i].ChainSel < roots[j].ChainSel })

	// TODO: use q.RMNSignatures in the generated outcome and eventually report. - Blocked by onchain work.

	outcome := ct.MerkleRootOutcome{
		OutcomeType:        outcomeType,
		RootsToReport:      roots,
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
	previousOutcome ct.MerkleRootOutcome,
	consensusObservation ConsensusObservation,
) ct.MerkleRootOutcome {

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
		return ct.MerkleRootOutcome{
			OutcomeType: ct.ReportTransmitted,
		}
	}

	if previousOutcome.ReportTransmissionCheckAttempts+1 >= maxReportTransmissionCheckAttempts {
		lggr.Warnw("Failed to detect report transmission")
		return ct.MerkleRootOutcome{
			OutcomeType: ct.ReportTransmissionFailed,
		}
	}

	return ct.MerkleRootOutcome{
		OutcomeType:                     ct.ReportInFlight,
		OffRampNextSeqNums:              previousOutcome.OffRampNextSeqNums,
		ReportTransmissionCheckAttempts: previousOutcome.ReportTransmissionCheckAttempts + 1,
	}
}

// getConsensusObservation Combine the list of observations into a single consensus observation
func getConsensusObservation(
	lggr logger.Logger,
	F int,
	destChain cciptypes.ChainSelector,
	aos []shared.AttributedObservation[ct.MerkleRootObservation],
) (ConsensusObservation, error) {
	aggObs := aggregateObservations(aos)

	fMin := make(map[cciptypes.ChainSelector]int)
	for chain := range aggObs.FChain {
		fMin[chain] = F
	}
	// consensus on the fChain map uses the role DON F value
	// because all nodes can observe the home chain.
	fChains := shared.GetConsensusMap(lggr, "fChain", aggObs.FChain, fMin)

	_, exists := fChains[destChain]
	if !exists {
		return ConsensusObservation{},
			fmt.Errorf("no consensus value for fDestChain, destChain: %d", destChain)
	}

	twoFPlus1 := make(map[cciptypes.ChainSelector]int)
	for chain, f := range fChains {
		twoFPlus1[chain] = 2*f + 1
	}

	consensusObs := ConsensusObservation{
		MerkleRoots:        shared.GetConsensusMap(lggr, "Merkle Root", aggObs.MerkleRoots, fChains),
		OnRampMaxSeqNums:   shared.GetConsensusMap(lggr, "OnRamp Max Seq Nums", aggObs.OnRampMaxSeqNums, twoFPlus1),
		OffRampNextSeqNums: shared.GetConsensusMap(lggr, "OffRamp Next Seq Nums", aggObs.OffRampNextSeqNums, fChains),
		FChain:             fChains,
	}

	return consensusObs, nil
}
