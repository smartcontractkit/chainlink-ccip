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
	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon/consensus"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugintypes"
	"github.com/smartcontractkit/chainlink-ccip/pkg/logutil"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
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

		rmnSigs := rmn.ReportSignatures{}
		if q.RMNSignatures != nil {
			rmnSigs = *q.RMNSignatures
		}

		merkleRootsOutcome, err := buildMerkleRootsOutcome(
			rmnSigs, p.offchainCfg.RMNEnabled, lggr, consObservation, previousOutcome, p.addressCodec)

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

	var rmnRemoteConfig cciptypes.RemoteConfig
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
	rmnSigs rmn.ReportSignatures,
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
		"rmnSignatures", rmnSigs)

	sort.Slice(roots, func(i, j int) bool { return roots[i].ChainSel < roots[j].ChainSel })

	sigs := make([]cciptypes.RMNECDSASignature, 0)
	var err error

	if len(roots) > 0 && rmnEnabled {
		sigs, err = rmn.NewECDSASigsFromPB(rmnSigs.Signatures)
		if err != nil {
			return Outcome{}, fmt.Errorf("failed to parse RMN signatures: %w", err)
		}

		roots = filterRootsBasedOnRmnSigs(
			lggr, rmnSigs.LaneUpdates, roots, consensusObservation.RMNEnabledChains, addressCodec)
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
// 1) RMN-enabled and roots exist for all the roots that are included in the RMN signature.
// 2) RMN-disabled and we do not have RMN signatures
func filterRootsBasedOnRmnSigs(
	lggr logger.Logger,
	signedLaneUpdates []*rmnpb.FixedDestLaneUpdate,
	proposedRoots []cciptypes.MerkleRootChain,
	rmnEnabledChains map[cciptypes.ChainSelector]bool,
	addressCodec cciptypes.AddressCodec,
) []cciptypes.MerkleRootChain {
	signedRoots, err := computeSignedRootsSet(lggr, signedLaneUpdates, addressCodec)
	if err != nil {
		lggr.Errorw("failed to compute signed roots set, skipping RMN-enabled roots", "err", err)
		// we don't return since we still want to make progress with the non-RMN related roots (RMN disabled chains)
	}
	lggr.Debugw("computed signed roots set", "signedRoots", signedRoots.ToSlice())

	// If at least ONE root that is signed is not proposed
	// then we cannot make progress with ANY of the existing signed roots
	// since the signature will be invalid
	proposedRootsSet := mapset.NewSet[rootKey]()
	for _, root := range proposedRoots {
		addrStr, err := addressCodec.AddressBytesToString(root.OnRampAddress, root.ChainSel)
		if err != nil {
			lggr.Errorw("convert proposed root OnRamp address to string to build a set", "root", root, "err", err)
			continue
		}
		proposedRootsSet.Add(rootKey{
			ChainSel:      root.ChainSel,
			SeqNumsRange:  root.SeqNumsRange,
			MerkleRoot:    root.MerkleRoot,
			OnRampAddress: addrStr,
		})
	}
	if !signedRoots.IsSubset(proposedRootsSet) {
		lggr.Errorw("signed roots are not a subset of proposed roots, skipping RMN-enabled roots",
			"proposedRoots", proposedRootsSet.ToSlice(), "signedRoots", signedRoots.ToSlice())
		// clear signed roots, so we can make progress with the non-signed roots below
		signedRoots = mapset.NewSet[rootKey]()
	}

	validRoots := filterValidRoots(lggr, proposedRoots, signedRoots, addressCodec, rmnEnabledChains)
	return validRoots
}

// computeSignedRootsSet generates a set of signed roots based on the provided signed lane updates.
func computeSignedRootsSet(
	lggr logger.Logger,
	signedLaneUpdates []*rmnpb.FixedDestLaneUpdate,
	addressCodec cciptypes.AddressCodec,
) (mapset.Set[rootKey], error) {
	signedRoots := mapset.NewSet[rootKey]()
	for _, laneUpdate := range signedLaneUpdates {
		addrStr, err := addressCodec.AddressBytesToString(
			laneUpdate.LaneSource.OnrampAddress,
			cciptypes.ChainSelector(laneUpdate.LaneSource.SourceChainSelector),
		)
		if err != nil {
			return mapset.NewSet[rootKey](), fmt.Errorf("convert address to string %v : %w", laneUpdate.LaneSource, err)
		}

		rk := rootKey{
			ChainSel: cciptypes.ChainSelector(laneUpdate.LaneSource.SourceChainSelector),
			SeqNumsRange: cciptypes.NewSeqNumRange(
				cciptypes.SeqNum(laneUpdate.ClosedInterval.MinMsgNr),
				cciptypes.SeqNum(laneUpdate.ClosedInterval.MaxMsgNr),
			),
			MerkleRoot:    cciptypes.Bytes32(laneUpdate.Root),
			OnRampAddress: addrStr,
		}

		lggr.Debugw("found signed root", "root", rk)
		signedRoots.Add(rk)
	}

	return signedRoots, nil
}

// filterValidRoots filters the roots based on the RMN-enabled chains and signed roots and returns the ones that
// are valid to proceed with. Valid roots are either:
// 1) RMN-enabled and signed
// 2) RMN-disabled and not signed
func filterValidRoots(
	lggr logger.Logger,
	proposedRoots []cciptypes.MerkleRootChain,
	signedRoots mapset.Set[rootKey],
	addressCodec cciptypes.AddressCodec,
	rmnEnabledChains map[cciptypes.ChainSelector]bool,
) []cciptypes.MerkleRootChain {
	validRoots := make([]cciptypes.MerkleRootChain, 0)
	for _, root := range proposedRoots {
		addrStr, err := addressCodec.AddressBytesToString(root.OnRampAddress, root.ChainSel)
		if err != nil {
			lggr.Errorw("convert proposed root OnRamp address to string to check root", "root", root, "err", err)
			continue
		}
		rk := rootKey{
			ChainSel:      root.ChainSel,
			SeqNumsRange:  root.SeqNumsRange,
			MerkleRoot:    root.MerkleRoot,
			OnRampAddress: addrStr,
		}

		rootIsSignedAndRmnEnabled := signedRoots.Contains(rk) && rmnEnabledChains[root.ChainSel]
		rootNotSignedAndRmnDisabled := !signedRoots.Contains(rk) && !rmnEnabledChains[root.ChainSel]
		rootIsValid := rootIsSignedAndRmnEnabled || rootNotSignedAndRmnDisabled
		lggr2 := logger.With(lggr,
			"root", rk, "isSigned", signedRoots.Contains(rk), "rmnEnabled", rmnEnabledChains[root.ChainSel])

		if rootIsValid {
			lggr2.Infow("root valid, added to the results")
			validRoots = append(validRoots, root)
		} else {
			lggr2.Infow("root invalid, skipping")
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
	rmnRemoteConfigs := map[cciptypes.ChainSelector][]cciptypes.RemoteConfig{destChain: aggObs.RMNRemoteConfigs}

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
