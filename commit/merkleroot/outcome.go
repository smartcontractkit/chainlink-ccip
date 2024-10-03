package merkleroot

import (
	"fmt"
	"sort"
	"time"

	mapset "github.com/deckarep/golang-set/v2"
	"golang.org/x/exp/maps"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-ccip/commit/merkleroot/rmn"
	rmntypes "github.com/smartcontractkit/chainlink-ccip/commit/merkleroot/rmn/types"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon/consensus"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugintypes"
)

// Outcome depending on the current state, either:
// - chooses the seq num ranges for the next round
// - builds a report
// - checks for the transmission of a previous report
func (w *Processor) Outcome(
	prevOutcome Outcome,
	query Query,
	aos []plugincommon.AttributedObservation[Observation],
) (Outcome, error) {
	tStart := time.Now()
	outcome, nextState := w.getOutcome(prevOutcome, query, aos)
	w.lggr.Infow("Sending Outcome",
		"outcome", outcome, "nextState", nextState, "outcomeDuration", time.Since(tStart))
	return outcome, nil
}

func (w *Processor) getOutcome(
	previousOutcome Outcome,
	q Query,
	aos []plugincommon.AttributedObservation[Observation],
) (Outcome, State) {
	nextState := previousOutcome.NextState()

	consensusObservation, err := getConsensusObservation(w.lggr, w.reportingCfg.F, w.destChain, aos)
	if err != nil {
		w.lggr.Warnw("Get consensus observation failed, empty outcome", "err", err)
		return Outcome{}, nextState
	}

	switch nextState {
	case SelectingRangesForReport:
		return reportRangesOutcome(q, w.lggr, consensusObservation, w.offchainCfg.MaxMerkleTreeSize, w.destChain), nextState
	case BuildingReport:
		if q.RetryRMNSignatures {
			// We want to retry getting the RMN signatures on the exact same outcome we had before.
			// The current observations should all be empty.
			return previousOutcome, BuildingReport
		}
		return buildReport(q, w.lggr, consensusObservation, previousOutcome), nextState
	case WaitingForReportTransmission:
		return checkForReportTransmission(
			w.lggr, w.offchainCfg.MaxReportTransmissionCheckAttempts, previousOutcome, consensusObservation), nextState
	default:
		w.lggr.Warnw("Unexpected next state in Outcome", "state", nextState)
		return Outcome{}, nextState
	}
}

// reportRangesOutcome determines the sequence number ranges for each chain to build a report from in the next round
func reportRangesOutcome(
	_ Query,
	lggr logger.Logger,
	consensusObservation ConsensusObservation,
	maxMerkleTreeSize uint64,
	dstChain cciptypes.ChainSelector,
) Outcome {
	rangesToReport := make([]plugintypes.ChainRange, 0)

	observedOnRampMaxSeqNumsMap := consensusObservation.OnRampMaxSeqNums
	observedOffRampNextSeqNumsMap := consensusObservation.OffRampNextSeqNums
	observedRMNRemoteConfig := consensusObservation.RMNRemoteConfig

	offRampNextSeqNums := make([]plugintypes.SeqNumChain, 0)

	for chainSel, offRampNextSeqNum := range observedOffRampNextSeqNumsMap {
		onRampMaxSeqNum, exists := observedOnRampMaxSeqNumsMap[chainSel]
		if !exists {
			continue
		}

		if offRampNextSeqNum <= onRampMaxSeqNum {
			rng := cciptypes.NewSeqNumRange(offRampNextSeqNum, onRampMaxSeqNum)

			chainRange := plugintypes.ChainRange{
				ChainSel:    chainSel,
				SeqNumRange: rng.Limit(maxMerkleTreeSize),
			}
			rangesToReport = append(rangesToReport, chainRange)

			if rng.End() != chainRange.SeqNumRange.End() { // Check if the range was truncated.
				lggr.Infof("Range for chain %d: %s (before truncate: %v)", chainSel, chainRange.SeqNumRange, rng)
			} else {
				lggr.Infof("Range for chain %d: %s", chainSel, chainRange.SeqNumRange)
			}
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

	var rmnRemoteConfig rmntypes.RemoteConfig
	if observedRMNRemoteConfig[dstChain].IsEmpty() {
		lggr.Warn("RMNRemoteConfig is nil")
	} else {
		rmnRemoteConfig = observedRMNRemoteConfig[dstChain]
	}

	outcome := Outcome{
		OutcomeType:             ReportIntervalsSelected,
		RangesSelectedForReport: rangesToReport,
		OffRampNextSeqNums:      offRampNextSeqNums,
		RMNRemoteCfg:            rmnRemoteConfig,
	}

	return outcome
}

// Given a set of observed merkle roots, gas prices and token prices, and roots from RMN, construct a report
// to transmit on-chain
func buildReport(
	q Query,
	lggr logger.Logger,
	consensusObservation ConsensusObservation,
	prevOutcome Outcome,
) Outcome {
	roots := maps.Values(consensusObservation.MerkleRoots)

	outcomeType := ReportGenerated
	if len(roots) == 0 {
		outcomeType = ReportEmpty
	}

	sort.Slice(roots, func(i, j int) bool { return roots[i].ChainSel < roots[j].ChainSel })

	sigs := make([]cciptypes.RMNECDSASignature, 0)
	if q.RMNSignatures != nil { // TODO: should never be nil, error after e2e RMN integration.
		parsedSigs, err := rmn.NewECDSASigsFromPB(q.RMNSignatures.Signatures)
		if err != nil {
			lggr.Errorw("Failed to parse RMN signatures returning an empty outcome", "err", err)
			return Outcome{}
		}
		sigs = parsedSigs

		signedRoots := mapset.NewSet[cciptypes.MerkleRootChain]()
		for _, laneUpdate := range q.RMNSignatures.LaneUpdates {
			signedRoots.Add(cciptypes.MerkleRootChain{
				ChainSel: cciptypes.ChainSelector(laneUpdate.LaneSource.SourceChainSelector),
				SeqNumsRange: cciptypes.NewSeqNumRange(
					cciptypes.SeqNum(laneUpdate.ClosedInterval.MinMsgNr),
					cciptypes.SeqNum(laneUpdate.ClosedInterval.MaxMsgNr),
				),
				MerkleRoot: cciptypes.Bytes32(laneUpdate.Root),
			})
		}

		// Only report roots that are present in RMN signatures.
		rootsToReport := make([]cciptypes.MerkleRootChain, 0)
		for _, root := range roots {
			if signedRoots.Contains(root) {
				rootsToReport = append(rootsToReport, root)
			} else {
				lggr.Warnw("skipping merkle root not signed by RMN", "root", root)
			}
		}
		roots = rootsToReport
	}

	outcome := Outcome{
		OutcomeType:         outcomeType,
		RootsToReport:       roots,
		OffRampNextSeqNums:  prevOutcome.OffRampNextSeqNums,
		RMNReportSignatures: sigs,
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
	fRoleDON int,
	destChain cciptypes.ChainSelector,
	aos []plugincommon.AttributedObservation[Observation],
) (ConsensusObservation, error) {
	aggObs := aggregateObservations(aos)

	// consensus on the fChain map uses the role DON F value
	// because all nodes can observe the home chain.
	donThresh := consensus.MakeConstantThreshold[cciptypes.ChainSelector](consensus.TwoFPlus1(fRoleDON))
	fChains := consensus.GetConsensusMap(lggr, "fChain", aggObs.FChain, donThresh)

	_, exists := fChains[destChain]
	if !exists {
		return ConsensusObservation{},
			fmt.Errorf("no consensus value for fDestChain, destChain: %d", destChain)
	}

	// convert aggObs.RMNRemoteConfigs to a map of RMNRemoteConfigs
	rmnRemoteConfigs := map[cciptypes.ChainSelector][]rmntypes.RemoteConfig{destChain: aggObs.RMNRemoteConfigs}

	// Get consensus using strict 2fChain+1 threshold.
	twoFChainPlus1 := consensus.MakeMultiThreshold(fChains, consensus.TwoFPlus1)
	consensusObs := ConsensusObservation{
		MerkleRoots:      consensus.GetConsensusMap(lggr, "Merkle Root", aggObs.MerkleRoots, twoFChainPlus1),
		OnRampMaxSeqNums: consensus.GetConsensusMap(lggr, "OnRamp Max Seq Nums", aggObs.OnRampMaxSeqNums, twoFChainPlus1),
		OffRampNextSeqNums: consensus.GetConsensusMap(
			lggr,
			"OffRamp Next Seq Nums",
			aggObs.OffRampNextSeqNums,
			twoFChainPlus1),
		RMNRemoteConfig: consensus.GetConsensusMap(lggr, "RMNRemote cfg", rmnRemoteConfigs, twoFChainPlus1),
		FChain:          fChains,
	}

	return consensusObs, nil
}
