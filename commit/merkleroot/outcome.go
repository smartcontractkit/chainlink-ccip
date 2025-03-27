package merkleroot

import (
	"context"
	"fmt"
	"sort"
	"time"

	"golang.org/x/exp/maps"

	mapset "github.com/deckarep/golang-set/v2"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"

	rmnpb "github.com/smartcontractkit/chainlink-protos/rmn/v1.6/go/serialization"

	"github.com/smartcontractkit/chainlink-ccip/commit/merkleroot/rmn"
	rmntypes "github.com/smartcontractkit/chainlink-ccip/commit/merkleroot/rmn/types"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon/consensus"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugintypes"
	"github.com/smartcontractkit/chainlink-ccip/pkg/logutil"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

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

	lggr.Infow("Sending Outcome",
		"outcome", outcome, "nextState", nextState, "outcomeDuration", time.Since(tStart))
	return outcome, nil
}

func (p *Processor) getOutcome(
	lggr logger.Logger,
	previousOutcome Outcome,
	q Query,
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
		return reportRangesOutcome(q, lggr, consObservation, p.offchainCfg.MaxMerkleTreeSize, p.destChain),
			nextState,
			nil
	case buildingReport:
		if q.RetryRMNSignatures {
			// We want to retry getting the RMN signatures on the exact same outcome we had before.
			// The current observations should all be empty.
			return previousOutcome, buildingReport, nil
		}

		merkleRootsOutcome, err := buildMerkleRootsOutcome(
			q, p.offchainCfg.RMNEnabled, lggr, consObservation, previousOutcome, p.addressCodec)

		return merkleRootsOutcome, nextState, err
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
	_ Query,
	lggr logger.Logger,
	consObservation consensusObservation,
	maxMerkleTreeSize uint64,
	dstChain cciptypes.ChainSelector,
) Outcome {
	rangesToReport := make([]plugintypes.ChainRange, 0)

	observedOnRampMaxSeqNumsMap := consObservation.OnRampMaxSeqNums
	observedOffRampNextSeqNumsMap := consObservation.OffRampNextSeqNums
	observedRMNRemoteConfig := consObservation.RMNRemoteConfig

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
			lggr.Errorw("sequence numbers between offRamp and onRamp reached an impossible state, "+
				"offRamp latest executed sequence number is greater than onRamp latest executed sequence number",
				"chain", chainSel, "onRampMaxSeqNum", onRampMaxSeqNum, "offRampNextSeqNum", offRampNextSeqNum)
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
				lggr.Infof("Range for chain %d: %s (before truncate: %v)", chainSel, chainRange.SeqNumRange, rng)
			} else {
				lggr.Infof("Range for chain %d: %s", chainSel, chainRange.SeqNumRange)
			}
		}
	}

	// deterministic outcome
	sort.Slice(rangesToReport, func(i, j int) bool { return rangesToReport[i].ChainSel < rangesToReport[j].ChainSel })
	sort.Slice(offRampNextSeqNums, func(i, j int) bool {
		return offRampNextSeqNums[i].ChainSel < offRampNextSeqNums[j].ChainSel
	})

	var rmnRemoteConfig rmntypes.RemoteConfig
	if observedRMNRemoteConfig[dstChain].IsEmpty() {
		lggr.Warn("RMNRemoteConfig is empty")
	} else {
		rmnRemoteConfig = observedRMNRemoteConfig[dstChain]
	}

	if len(rangesToReport) == 0 {
		lggr.Info("No ranges to report, outcomeType is ReportEmpty")
		return Outcome{OutcomeType: ReportEmpty}
	}

	outcome := Outcome{
		OutcomeType:             ReportIntervalsSelected,
		RangesSelectedForReport: rangesToReport,
		OffRampNextSeqNums:      offRampNextSeqNums,
		RMNRemoteCfg:            rmnRemoteConfig,
	}

	return outcome
}

// buildMerkleRootsOutcome is given a set of agreed observed merkle roots and RMN signatures
// and construct a merkleRoots outcome.
func buildMerkleRootsOutcome(
	q Query,
	rmnEnabled bool,
	lggr logger.Logger,
	consensusObservation consensusObservation,
	prevOutcome Outcome,
	addressCodec cciptypes.AddressCodec,
) (Outcome, error) {
	roots := maps.Values(consensusObservation.MerkleRoots)

	outcomeType := ReportGenerated
	if len(roots) == 0 {
		outcomeType = ReportEmpty
	}

	lggr.Debugw("building merkle roots outcome",
		"rmnEnabled", rmnEnabled,
		"rmnEnabledChains", consensusObservation.RMNEnabledChains,
		"roots", roots,
		"rmnSignatures", q.RMNSignatures)

	sort.Slice(roots, func(i, j int) bool { return roots[i].ChainSel < roots[j].ChainSel })

	sigs := make([]cciptypes.RMNECDSASignature, 0)
	var err error

	if len(roots) > 0 && rmnEnabled {
		sigs, err = rmn.NewECDSASigsFromPB(q.RMNSignatures.Signatures)
		if err != nil {
			return Outcome{}, fmt.Errorf("failed to parse RMN signatures: %w", err)
		}

		roots = filterRootsBasedOnRmnSigs(
			lggr, q.RMNSignatures.LaneUpdates, roots, consensusObservation.RMNEnabledChains, addressCodec)
	}

	outcome := Outcome{
		OutcomeType:         outcomeType,
		RootsToReport:       roots,
		RMNEnabledChains:    consensusObservation.RMNEnabledChains,
		OffRampNextSeqNums:  prevOutcome.OffRampNextSeqNums,
		RMNReportSignatures: sigs,
		RMNRemoteCfg:        prevOutcome.RMNRemoteCfg,
	}

	return outcome, nil
}

// filterRootsBasedOnRmnSigs filters the roots to only include the ones that are either:
// 1) RMN-enabled and have RMN signatures
// 2) RMN-disabled and do not have RMN signatures
func filterRootsBasedOnRmnSigs(
	lggr logger.Logger,
	signedLaneUpdates []*rmnpb.FixedDestLaneUpdate,
	roots []cciptypes.MerkleRootChain,
	rmnEnabledChains map[cciptypes.ChainSelector]bool,
	addressCodec cciptypes.AddressCodec,
) []cciptypes.MerkleRootChain {
	// Create a set of signed roots for quick lookup.
	signedRoots := mapset.NewSet[rootKey]()
	for _, laneUpdate := range signedLaneUpdates {
		addrStr, err := addressCodec.AddressBytesToString(
			laneUpdate.LaneSource.OnrampAddress,
			cciptypes.ChainSelector(laneUpdate.LaneSource.SourceChainSelector))
		if err != nil {
			lggr.Errorw("can't convert Onramp address to string",
				"err", err,
				"sourceChainSelector", laneUpdate.LaneSource.SourceChainSelector,
				"onrampAddress", laneUpdate.LaneSource.OnrampAddress)
			continue
		}
		rk := rootKey{
			ChainSel: cciptypes.ChainSelector(laneUpdate.LaneSource.SourceChainSelector),
			SeqNumsRange: cciptypes.NewSeqNumRange(
				cciptypes.SeqNum(laneUpdate.ClosedInterval.MinMsgNr),
				cciptypes.SeqNum(laneUpdate.ClosedInterval.MaxMsgNr),
			),
			MerkleRoot: cciptypes.Bytes32(laneUpdate.Root),
			// NOTE: convert address into a comparable value for mapset.
			OnRampAddress: addrStr,
		}
		lggr.Infow("Found signed root", "root", rk)
		signedRoots.Add(rk)
	}

	validRoots := make([]cciptypes.MerkleRootChain, 0)
	for _, root := range roots {
		addrStr, err := addressCodec.AddressBytesToString(root.OnRampAddress, root.ChainSel)
		if err != nil {
			lggr.Errorw("can't convert Onramp address to string",
				"err", err, "sourceChainSelector", root.ChainSel, "onrampAddress", root.OnRampAddress)
			continue
		}
		rk := rootKey{
			ChainSel:      root.ChainSel,
			SeqNumsRange:  root.SeqNumsRange,
			MerkleRoot:    root.MerkleRoot,
			OnRampAddress: addrStr,
		}

		rootIsSignedAndRmnEnabled := signedRoots.Contains(rk) &&
			rmnEnabledChains[root.ChainSel]

		rootNotSignedButRmnDisabled := !signedRoots.Contains(rk) &&
			!rmnEnabledChains[root.ChainSel]

		if rootIsSignedAndRmnEnabled || rootNotSignedButRmnDisabled {
			lggr.Infow("Adding root to the report",
				"root", rk,
				"rootIsSignedAndRmnEnabled", rootIsSignedAndRmnEnabled,
				"rootNotSignedButRmnDisabled", rootNotSignedButRmnDisabled)
			validRoots = append(validRoots, root)
		} else {
			lggr.Infow("Root invalid, skipping from the report",
				"root", rk,
				"rootIsSignedAndRmnEnabled", rootIsSignedAndRmnEnabled,
				"rootNotSignedButRmnDisabled", rootNotSignedButRmnDisabled,
			)
		}
	}

	return validRoots
}

type rootKey struct {
	ChainSel      cciptypes.ChainSelector
	SeqNumsRange  cciptypes.SeqNumRange
	MerkleRoot    cciptypes.Bytes32
	OnRampAddress string
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
	rmnRemoteConfigs := map[cciptypes.ChainSelector][]rmntypes.RemoteConfig{destChain: aggObs.RMNRemoteConfigs}

	// Get consensus using strict 2fChain+1 threshold.
	twoFChainPlus1 := consensus.MakeMultiThreshold(fChains, consensus.TwoFPlus1)
	fChain := consensus.MakeMultiThreshold(fChains, consensus.F)

	consensusObs := consensusObservation{
		MerkleRoots:      consensus.GetConsensusMap(lggr, "Merkle Root", aggObs.MerkleRoots, twoFChainPlus1),
		RMNEnabledChains: consensus.GetConsensusMap(lggr, "RMNEnabledChains", aggObs.RMNEnabledChains, twoFChainPlus1),
		OnRampMaxSeqNums: consensus.GetOrderedConsensus(
			lggr,
			"OnRamp Max Seq Nums",
			aggObs.OnRampMaxSeqNums,
			fChain),
		OffRampNextSeqNums: consensus.GetOrderedConsensus(
			lggr,
			"OffRamp Next Seq Nums",
			aggObs.OffRampNextSeqNums,
			fChain),
		RMNRemoteConfig: consensus.GetConsensusMap(lggr, "RMNRemote cfg", rmnRemoteConfigs, twoFChainPlus1),
		FChain:          fChains,
	}

	return consensusObs, nil
}
