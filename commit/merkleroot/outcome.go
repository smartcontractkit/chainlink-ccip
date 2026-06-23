package merkleroot

import (
	"context"
	"fmt"
	"maps"
	"slices"
	"sort"
	"time"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon/consensus"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugintypes"
	"github.com/smartcontractkit/chainlink-ccip/pkg/logutil"
)

const SendingOutcome = "Sending Outcome"

// Outcome depending on the current state, either:
// - chooses the seq num ranges for the next round
// - builds a report
// - checks for the transmission of a previous report
func (p *Processor) Outcome(
	ctx context.Context,
	prevOutcome Outcome,
	query Query,
	aos []plugincommon.AttributedObservation[Observation],
) (Outcome, error) {
	lggr := logutil.WithContextValues(ctx, p.lggr)

	tStart := time.Now()

	outcome, nextState, err := p.getOutcome(lggr, prevOutcome, query, aos)
	if err != nil {
		lggr.Errorw("outcome failed with error", "err", err)
		return Outcome{}, err
	}

	lggr.Infow(SendingOutcome,
		"outcome", outcome, "nextState", nextState, "outcomeDuration", time.Since(tStart))
	return outcome, nil
}

func (p *Processor) getOutcome(
	lggr logger.Logger,
	previousOutcome Outcome,
	_ Query,
	aos []plugincommon.AttributedObservation[Observation],
) (Outcome, processorState, error) {
	nextState := previousOutcome.nextState()

	consObservation, err := getConsensusObservation(lggr, p.reportingCfg.F, p.destChain, aos)
	if err != nil {
		lggr.Warnw("Get consensus observation failed, empty outcome", "err", err)
		return Outcome{}, nextState, nil
	}

	switch nextState {
	case selectingRangesForReport:
		return reportRangesOutcome(lggr, consObservation, p.offchainCfg.MaxMerkleTreeSize),
			nextState,
			nil
	case buildingReport:
		merkleRootsOutcome := buildMerkleRootsOutcome(lggr, consObservation, previousOutcome)
		return merkleRootsOutcome, nextState, nil
	case waitingForReportTransmission:
		attempts := p.offchainCfg.MaxReportTransmissionCheckAttempts
		multipleReports := p.offchainCfg.MultipleReportsEnabled
		outcome := checkForReportTransmission(lggr, attempts, multipleReports, previousOutcome, consObservation)
		return outcome, nextState, nil
	default:
		return Outcome{}, nextState, fmt.Errorf("unexpected next state in Outcome: %v", nextState)
	}
}

// reportRangesOutcome determines the sequence number ranges for each chain to build a report from in the next round
func reportRangesOutcome(
	lggr logger.Logger,
	consObservation consensusObservation,
	maxMerkleTreeSize uint64,
) Outcome {
	rangesToReport := make([]plugintypes.ChainRange, 0)

	observedOnRampMaxSeqNumsMap := consObservation.OnRampMaxSeqNums
	observedOffRampNextSeqNumsMap := consObservation.OffRampNextSeqNums

	offRampNextSeqNums := make([]plugintypes.SeqNumChain, 0)

	for chainSel, offRampNextSeqNum := range observedOffRampNextSeqNumsMap {
		offRampNextSeqNums = append(offRampNextSeqNums, plugintypes.SeqNumChain{
			ChainSel: chainSel,
			SeqNum:   offRampNextSeqNum,
		})

		onRampMaxSeqNum, exists := observedOnRampMaxSeqNumsMap[chainSel]
		if !exists {
			continue
		}

		if onRampMaxSeqNum < offRampNextSeqNum-1 {
			if onRampMaxSeqNum == 0 {
				lggr.Infow("OnRamp max sequence numbers consensus = 0. This might not indicate an issue" +
					" but if it persists without progress on the commit plugin, investigate why oracles observe 0")
			} else {
				lggr.Errorw("sequence numbers between offRamp and onRamp reached an impossible state, "+
					"offRamp latest executed sequence number is greater than onRamp latest executed sequence number",
					"chain", chainSel, "onRampMaxSeqNum", onRampMaxSeqNum, "offRampNextSeqNum", offRampNextSeqNum)
			}
		}

		newMsgsExist := offRampNextSeqNum <= onRampMaxSeqNum
		if newMsgsExist {
			rng := cciptypes.NewSeqNumRange(offRampNextSeqNum, onRampMaxSeqNum)

			chainRange := plugintypes.ChainRange{
				ChainSel:    chainSel,
				SeqNumRange: rng.Limit(maxMerkleTreeSize),
			}
			rangesToReport = append(rangesToReport, chainRange)

			if rng.End() != chainRange.SeqNumRange.End() { // Check if the range was truncated.
				lggr.Debugf("Range for chain %d: %s (before truncate: %v)", chainSel, chainRange.SeqNumRange, rng)
			} else {
				lggr.Debugf("Range for chain %d: %s", chainSel, chainRange.SeqNumRange)
			}
		}
	}

	// deterministic outcome
	sort.Slice(rangesToReport, func(i, j int) bool { return rangesToReport[i].ChainSel < rangesToReport[j].ChainSel })
	sort.Slice(offRampNextSeqNums, func(i, j int) bool {
		return offRampNextSeqNums[i].ChainSel < offRampNextSeqNums[j].ChainSel
	})

	if len(rangesToReport) == 0 {
		lggr.Debug("No ranges to report, outcomeType is ReportEmpty")
		return Outcome{OutcomeType: ReportEmpty}
	}

	return Outcome{
		OutcomeType:             ReportIntervalsSelected,
		RangesSelectedForReport: rangesToReport,
		OffRampNextSeqNums:      offRampNextSeqNums,
	}
}

// buildMerkleRootsOutcome constructs a merkle roots outcome from agreed observations.
func buildMerkleRootsOutcome(
	lggr logger.Logger,
	consensusObservation consensusObservation,
	prevOutcome Outcome,
) Outcome {
	roots := slices.Collect(maps.Values(consensusObservation.MerkleRoots))

	outcomeType := ReportGenerated
	if len(roots) == 0 {
		outcomeType = ReportEmpty
	}

	lggr.Debugw("building merkle roots outcome", "roots", roots)

	sort.Slice(roots, func(i, j int) bool { return roots[i].ChainSel < roots[j].ChainSel })

	return Outcome{
		OutcomeType:        outcomeType,
		RootsToReport:      roots,
		OffRampNextSeqNums: prevOutcome.OffRampNextSeqNums,
	}
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
	multipleReports bool,
	previousOutcome Outcome,
	consensusObservation consensusObservation,
) Outcome {
	// Check that all sources have been updates using a set initialized from the previous outcome.
	// Check that all have been updated in case there were multiple reports generated in the previous round.
	pendingSources := make(map[cciptypes.ChainSelector]struct{})
	for _, root := range previousOutcome.RootsToReport {
		pendingSources[root.ChainSel] = struct{}{}
	}

	for _, previousSeqNumChain := range previousOutcome.OffRampNextSeqNums {
		if currentSeqNum, exists := consensusObservation.OffRampNextSeqNums[previousSeqNumChain.ChainSel]; exists {
			if previousSeqNumChain.SeqNum < currentSeqNum {
				// if there is only one report, any single update means the report has been transmitted.
				if !multipleReports {
					return Outcome{
						OutcomeType: ReportTransmitted,
					}
				}

				delete(pendingSources, previousSeqNumChain.ChainSel)
			}

			if previousSeqNumChain.SeqNum > currentSeqNum {
				lggr.Errorw("OffRampNextSeqNums reached an impossible state, "+
					"previous offRampNextSeqNum is greater than current offRampNextSeqNum",
					"chain", previousSeqNumChain.ChainSel,
					"previousSeqNum", previousSeqNumChain.SeqNum,
					"currentSeqNum", currentSeqNum,
				)
			}
		}
	}

	// All pending sources have been updated, we can move to the next state.
	if len(pendingSources) == 0 {
		return Outcome{
			OutcomeType: ReportTransmitted,
		}
	}

	if previousOutcome.ReportTransmissionCheckAttempts+1 >= maxReportTransmissionCheckAttempts {
		lggr.Warnw("report not transmitted, max check attempts reached, moving to next state")
		return Outcome{
			OutcomeType: ReportTransmissionFailed,
		}
	}

	return Outcome{
		OutcomeType:                     ReportInFlight,
		OffRampNextSeqNums:              previousOutcome.OffRampNextSeqNums,
		ReportTransmissionCheckAttempts: previousOutcome.ReportTransmissionCheckAttempts + 1,
		// Carry over the previous roots since they're still in-flight.
		// We won't re-report since outcome type is ReportInFlight.
		RootsToReport: previousOutcome.RootsToReport,
	}
}

// getConsensusObservation Combine the list of observations into a single consensus observation
func getConsensusObservation(
	lggr logger.Logger,
	fRoleDON int,
	destChain cciptypes.ChainSelector,
	aos []plugincommon.AttributedObservation[Observation],
) (consensusObservation, error) {
	aggObs := aggregateObservations(aos)

	// consensus on the fChain map uses the role DON F value
	// because all nodes can observe the home chain.
	donThresh := consensus.MakeConstantThreshold[cciptypes.ChainSelector](consensus.TwoFPlus1(fRoleDON))
	fChains := consensus.GetConsensusMap(lggr, "fChain", aggObs.FChain, donThresh)

	_, exists := fChains[destChain]
	if !exists {
		return consensusObservation{},
			fmt.Errorf("no consensus value for fDestChain, destChain: %d, fChainObs: %+v, fChainsConsensus: %+v",
				destChain, aggObs.FChain, fChains,
			)
	}

	// convert aggObs.RMNRemoteConfigs to a map of RMNRemoteConfigs
	rmnRemoteConfigs := map[cciptypes.ChainSelector][]cciptypes.RemoteConfig{destChain: aggObs.RMNRemoteConfigs}

	// Get consensus using strict 2fChain+1 threshold.
	twoFChainPlus1 := consensus.MakeMultiThreshold(fChains, consensus.TwoFPlus1)
	fChain := consensus.MakeMultiThreshold(fChains, consensus.F)

	fDestChain, ok := fChain.Get(destChain)
	if !ok {
		return consensusObservation{}, fmt.Errorf("no consensus value for fDestChain(%d): %v", fDestChain, fChain)
	}

	consensusObs := consensusObservation{
		MerkleRoots:      consensus.GetConsensusMap(lggr, "Merkle Root", aggObs.MerkleRoots, twoFChainPlus1),
		RMNEnabledChains: consensus.GetConsensusMap(lggr, "RMNEnabledChains", aggObs.RMNEnabledChains, twoFChainPlus1),
		OnRampMaxSeqNums: consensus.GetOrderedConsensus(
			lggr,
			"OnRamp Max Seq Nums",
			aggObs.OnRampMaxSeqNums,
			fChain),
		OffRampNextSeqNums: getOffRampNextSequenceNumbersConsensus(lggr, uint(fDestChain), aggObs.OffRampNextSeqNums),
		RMNRemoteConfig:    consensus.GetConsensusMap(lggr, "RMNRemote cfg", rmnRemoteConfigs, twoFChainPlus1),
		FChain:             fChains,
	}

	return consensusObs, nil
}

// getOffRampNextSequenceNumbersConsensus accepts a list of offramp sequence number observations per chain
// and computes the consensus value for each chain.
//
// Similar to consensus.GetOrderedConsensus but uses fDestChain, since this values are observed
// from the destination chain, instead of fChain for each source chain.
func getOffRampNextSequenceNumbersConsensus(
	lggr logger.Logger,
	fDestChain uint,
	observationsPerChain map[cciptypes.ChainSelector][]cciptypes.SeqNum,
) map[cciptypes.ChainSelector]cciptypes.SeqNum {
	lggr = logger.With(lggr, "fDestChain", fDestChain)

	offRampNextSeqNumsConsensus := make(map[cciptypes.ChainSelector]cciptypes.SeqNum)
	for sourceChain, observedNextSeqNums := range observationsPerChain {
		if uint(len(observedNextSeqNums)) < 2*fDestChain+1 {
			lggr.Warnw("not enough observations for OffRampNextSeqNums consensus on chain",
				"sourceChain", sourceChain, "observedNextSeqNums", observedNextSeqNums,
			)
			continue
		}

		slices.Sort(observedNextSeqNums)
		offRampNextSeqNumsConsensus[sourceChain] = observedNextSeqNums[fDestChain]
	}

	lggr.Debugw("computed offRampNextSeqNumsConsensus",
		"offRampNextSeqNumsConsensus", offRampNextSeqNumsConsensus, "observations", observationsPerChain)
	return offRampNextSeqNumsConsensus
}
