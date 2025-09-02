package commit

import (
	"bytes"
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"sort"
	"time"

	mapset "github.com/deckarep/golang-set/v2"

	"golang.org/x/exp/maps"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-ccip/commit/internal/builder"
	"github.com/smartcontractkit/chainlink-ccip/commit/merkleroot"
	"github.com/smartcontractkit/chainlink-ccip/internal/libs/slicelib"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon/consensus"
	"github.com/smartcontractkit/chainlink-ccip/pkg/consts"
	"github.com/smartcontractkit/chainlink-ccip/pkg/logutil"
)

func encodeReports(
	ctx context.Context,
	lggr logger.Logger,
	reports []builder.Report,
	transmissionSchedule *ocr3types.TransmissionSchedule,
	reportCodec cciptypes.CommitPluginCodec,
) ([]ocr3types.ReportPlus[[]byte], error) {
	var encodedReports []ocr3types.ReportPlus[[]byte]
	// Encode the reports and report info
	for _, report := range reports {
		// the report builder should not include empty reports.
		if report.Report.IsEmpty() {
			return nil, fmt.Errorf("found empty report")
		}

		lggr.Debugw("encoding report and report info",
			"report", report.Report,
			"reportInfo", report.ReportInfo)

		encodedReport, err := reportCodec.Encode(ctx, report.Report)
		if err != nil {
			return nil, fmt.Errorf("encode commit plugin report: %w", err)
		}

		encodedInfo, err := report.ReportInfo.Encode()
		if err != nil {
			return nil, fmt.Errorf("encode commit plugin report info: %w", err)
		}

		encodedReports = append(encodedReports, ocr3types.ReportPlus[[]byte]{
			ReportWithInfo: ocr3types.ReportWithInfo[[]byte]{
				Report: encodedReport,
				Info:   encodedInfo,
			},
			TransmissionScheduleOverride: transmissionSchedule,
		})
	}
	return encodedReports, nil
}

func (p *Plugin) Reports(
	ctx context.Context, seqNr uint64, outcomeBytes ocr3types.Outcome,
) ([]ocr3types.ReportPlus[[]byte], error) {
	ctx, lggr := logutil.WithOCRInfo(ctx, p.lggr, seqNr, logutil.PhaseReports)

	outcome, err := p.ocrTypeCodec.DecodeOutcome(outcomeBytes)
	if err != nil {
		lggr.Errorw("failed to decode Outcome", "outcome", string(outcomeBytes), "err", err)
		return nil, fmt.Errorf("decode outcome: %w", err)
	}

	lggr.Debugw("generating report",
		"roots", outcome.MerkleRootOutcome.RootsToReport,
		"tokenPriceUpdates", outcome.TokenPriceOutcome.TokenPrices,
		"gasPriceUpdates", outcome.ChainFeeOutcome.GasPrices,
		"rmnSignatures", outcome.MerkleRootOutcome.RMNReportSignatures,
	)

	transmissionSchedule, err := plugincommon.GetTransmissionSchedule(
		p.chainSupport,
		maps.Keys(p.oracleIDToP2PID),
		p.offchainCfg.TransmissionDelayMultiplier,
	)
	if err != nil {
		return nil, fmt.Errorf("get transmission schedule: %w", err)
	}
	lggr.Debugw("transmission schedule override",
		"transmissionSchedule", transmissionSchedule, "oracleIDToP2PID", p.oracleIDToP2PID)

	// Build reports for outcome
	reports, err := p.reportBuilder(lggr, outcome, p.offchainCfg)
	if err != nil {
		lggr.Errorw("failed to build reports",
			"outcome", outcome,
			"err", err)
		return nil, fmt.Errorf("err in Reports(): %w", err)
	}

	encodedReports, err := encodeReports(ctx, lggr, reports, transmissionSchedule, p.reportCodec)
	if err != nil {
		lggr.Errorw("unable to encode reports",
			"reports", reports,
			"outcome", outcome,
			"err", err)
	}

	lggr.Infow(fmt.Sprintf("Report building complete: built %d reports", len(reports)),
		"numReport", len(reports),
	)

	return encodedReports, nil
}

// validateReport validates various aspects of the report.
// Pure checks are placed earlier in the function on purpose to avoid
// unnecessary network or DB I/O.
// If you're added more checks make sure to follow this pattern.
//
//nolint:gocyclo
func (p *Plugin) validateReport(
	ctx context.Context,
	lggr logger.Logger,
	seqNr uint64,
	r ocr3types.ReportWithInfo[[]byte],
) (cciptypes.CommitPluginReport, error) {
	if r.Report == nil {
		lggr.Warn("nil report")
		return cciptypes.CommitPluginReport{}, nil
	}

	decodedReport, err := p.decodeReport(ctx, r.Report)
	if err != nil {
		return cciptypes.CommitPluginReport{},
			plugincommon.NewErrValidatingReport(fmt.Errorf("decode report: %w, report: %x", err, r.Report))
	}

	var reportInfo cciptypes.CommitReportInfo
	if reportInfo, err = cciptypes.DecodeCommitReportInfo(r.Info); err != nil {
		return cciptypes.CommitPluginReport{},
			plugincommon.NewErrValidatingReport(fmt.Errorf("decode report info: %w", err))
	}

	for _, root := range decodedReport.BlessedMerkleRoots {
		if root.MerkleRoot == (cciptypes.Bytes32{}) {
			lggr.Warnw("empty blessed merkle root", "root", root)
			return cciptypes.CommitPluginReport{}, plugincommon.NewErrInvalidReport("empty blessed merkle root")

		}
		if root.SeqNumsRange.Start() > root.SeqNumsRange.End() {
			lggr.Warnw("invalid seqNumsRange", "blessed root", root)
			return cciptypes.CommitPluginReport{}, plugincommon.NewErrInvalidReport("invalid seqNumsRange")
		}
	}

	for _, root := range decodedReport.UnblessedMerkleRoots {
		if root.MerkleRoot == (cciptypes.Bytes32{}) {
			lggr.Warnw("empty unblessed merkle root", "root", root)
			return cciptypes.CommitPluginReport{}, plugincommon.NewErrInvalidReport("empty unblessed merkle root")
		}
		if root.SeqNumsRange.Start() > root.SeqNumsRange.End() {
			lggr.Warnw("invalid seqNumsRange", "unblessed root", root)
			return cciptypes.CommitPluginReport{}, plugincommon.NewErrInvalidReport("invalid seqNumsRange")
		}
	}

	seen := make(map[cciptypes.RMNECDSASignature]struct{})
	for _, sig := range decodedReport.RMNSignatures {

		if _, ok := seen[sig]; ok {
			lggr.Warnw("duplicate RMN signature", "sig", sig)
			return cciptypes.CommitPluginReport{}, plugincommon.NewErrInvalidReport("duplicate RMN signature")
		}
		seen[sig] = struct{}{}
	}

	if p.offchainCfg.RMNEnabled &&
		len(decodedReport.BlessedMerkleRoots) > 0 &&
		consensus.LtFPlusOne(int(reportInfo.RemoteF), len(decodedReport.RMNSignatures)) {
		lggr.Debugf("report with insufficient RMN signatures %d < %d+1",
			len(decodedReport.RMNSignatures), reportInfo.RemoteF)
		return cciptypes.CommitPluginReport{}, plugincommon.NewErrInvalidReport("insufficient RMN signatures")
	}

	if isCursed, err := p.checkReportCursed(ctx, lggr, decodedReport); err != nil || isCursed {
		if err != nil {
			return cciptypes.CommitPluginReport{},
				plugincommon.NewErrValidatingReport(fmt.Errorf("check report cursed: %w", err))
		}
		return cciptypes.CommitPluginReport{}, plugincommon.NewErrInvalidReport("report cursed")
	}

	// check if we support the dest, if not we can't do the checks needed.
	supports, err := p.chainSupport.SupportsDestChain(p.oracleID)
	if err != nil {
		return cciptypes.CommitPluginReport{},
			plugincommon.NewErrValidatingReport(fmt.Errorf("supports dest chain: %w", err))
	}

	if !supports {
		lggr.Errorw("dest chain not supported by this oracle, can't run report acceptance procedures, " +
			"transmission schedule is wrong - check CCIPHome chainConfigs and ensure that the right oracles are " +
			"assigned as readers of the destination chain, or if " +
			"this oracle should support the destination chain but isn't!")
		return cciptypes.CommitPluginReport{}, plugincommon.NewErrInvalidReport("dest chain not supported")
	}

	offRampConfigDigest, err := p.ccipReader.GetOffRampConfigDigest(ctx, consts.PluginTypeCommit)
	if err != nil {
		err = plugincommon.NewErrValidatingReport(fmt.Errorf("get offramp config digest: %w", err))
		return cciptypes.CommitPluginReport{}, plugincommon.NewErrValidatingReport(err)
	}

	if !bytes.Equal(offRampConfigDigest[:], p.reportingCfg.ConfigDigest[:]) {
		lggr.Warnw("my config digest doesn't match offramp's config digest, not accepting report",
			"myConfigDigest", p.reportingCfg.ConfigDigest,
			"offRampConfigDigest", hex.EncodeToString(offRampConfigDigest[:]),
		)
		return cciptypes.CommitPluginReport{}, plugincommon.NewErrInvalidReport("config digest mismatch")
	}

	err = p.isStaleReport(ctx, seqNr, decodedReport)
	if err != nil {
		return cciptypes.CommitPluginReport{}, plugincommon.NewErrStaleReport(fmt.Sprintf("%v", err))
	}

	err = merkleroot.ValidateRootBlessings(
		ctx,
		p.ccipReader,
		decodedReport.BlessedMerkleRoots,
		decodedReport.UnblessedMerkleRoots,
	)
	if err != nil {
		lggr.Errorw("report not accepted due to root blessings validation error", "err", err)
		err = plugincommon.NewErrInvalidReport(fmt.Sprintf("root blessings validation: %v", err))
		return cciptypes.CommitPluginReport{}, err
	}

	return decodedReport, nil
}

func (p *Plugin) ShouldAcceptAttestedReport(
	ctx context.Context, seqNr uint64, r ocr3types.ReportWithInfo[[]byte],
) (bool, error) {
	ctx, lggr := logutil.WithOCRInfo(ctx, p.lggr, seqNr, logutil.PhaseShouldAccept)

	decodedReport, err := p.validateReport(ctx, lggr, seqNr, r)
	if errors.Is(err, plugincommon.ErrStaleReport) {
		lggr.Infow("stale report, not accepting", "err", err)
		return false, nil
	}
	if err != nil {
		lggr.Errorw("validation error", "err", err)
		return false, fmt.Errorf("validating report: %w", err)
	}

	lggr.Infow("ShouldAcceptedAttestedReport passed checks",
		"timestamp", time.Now().UTC(),
		"blessedRootsLen", len(decodedReport.BlessedMerkleRoots),
		"unblessedRootsLen", len(decodedReport.UnblessedMerkleRoots),
		"tokenPriceUpdatesLen", len(decodedReport.PriceUpdates.TokenPriceUpdates),
		"gasPriceUpdatesLen", len(decodedReport.PriceUpdates.GasPriceUpdates),
	)
	return true, nil
}

func (p *Plugin) decodeReport(
	ctx context.Context,
	report []byte,
) (cciptypes.CommitPluginReport, error) {
	decodedReport, err := p.reportCodec.Decode(ctx, report)
	if err != nil {
		return cciptypes.CommitPluginReport{},
			fmt.Errorf("decode commit plugin report: %w", err)
	}
	if decodedReport.IsEmpty() {
		return cciptypes.CommitPluginReport{},
			fmt.Errorf("empty report after decoding")
	}
	return decodedReport, nil
}

func (p *Plugin) isStaleReport(
	ctx context.Context,
	ocrSeqNr uint64,
	decodedReport cciptypes.CommitPluginReport,
) error {
	latestPriceSeqNr, err := p.ccipReader.GetLatestPriceSeqNr(ctx)
	if err != nil {
		return fmt.Errorf("get latest price seq nr: %w", err)
	}

	if ocrSeqNr <= latestPriceSeqNr &&
		len(decodedReport.BlessedMerkleRoots) == 0 &&
		len(decodedReport.UnblessedMerkleRoots) == 0 {
		return fmt.Errorf(
			"stale report: ocrSeqNr %d <= latestPriceSeqNr %d and no merkle roots", ocrSeqNr, latestPriceSeqNr,
		)
	}
	proposedBlessedMerkleRoots := decodedReport.BlessedMerkleRoots
	proposedUnblessedMerkleRoots := decodedReport.UnblessedMerkleRoots

	if len(proposedBlessedMerkleRoots) == 0 && len(proposedUnblessedMerkleRoots) == 0 {
		return nil
	}

	proposedMerkleRoots := append(proposedBlessedMerkleRoots, proposedUnblessedMerkleRoots...)

	chainSet := mapset.NewSet[cciptypes.ChainSelector]()
	newNextOnRampSeqNums := make(map[cciptypes.ChainSelector]cciptypes.SeqNum)

	for _, r := range proposedMerkleRoots {
		if chainSet.Contains(r.ChainSel) {
			return fmt.Errorf("duplicate chain %d", r.ChainSel)
		}
		chainSet.Add(r.ChainSel)
		newNextOnRampSeqNums[r.ChainSel] = r.SeqNumsRange.Start()
	}

	chainSlice := chainSet.ToSlice()
	sort.Slice(chainSlice, func(i, j int) bool { return chainSlice[i] < chainSlice[j] })

	offRampExpNextSeqNums, err := p.ccipReader.NextSeqNum(ctx, chainSlice)
	if err != nil {
		return fmt.Errorf("get next sequence numbers: %w", err)
	}

	for chain, newNextOnRampSeqNum := range newNextOnRampSeqNums {
		offRampExpNextSeqNum, ok := offRampExpNextSeqNums[chain]
		if !ok {
			// Due to some chain being disabled while the sequence numbers were already observed.
			// Report should not be considered valid in that case.
			return fmt.Errorf("offRamp expected next sequence number for chain %d was not found", chain)
		}

		if newNextOnRampSeqNum != offRampExpNextSeqNum {
			return fmt.Errorf("the merkle root that we are about to propose is stale, some previous report "+
				"made it on-chain, consider waiting more time for the reports to make it on chain. "+
				"offramp expects %d but we are proposing a root with min seq num %d for chain %d. "+
				"BlessedRoots: %v \nUnblessedRoots: %v",
				offRampExpNextSeqNum, newNextOnRampSeqNum, chain,
				proposedBlessedMerkleRoots, proposedUnblessedMerkleRoots,
			)
		}
	}
	return nil
}

func (p *Plugin) checkReportCursed(
	ctx context.Context,
	lggr logger.Logger,
	decodedReport cciptypes.CommitPluginReport,
) (bool, error) {
	allRoots := append(decodedReport.BlessedMerkleRoots, decodedReport.UnblessedMerkleRoots...)

	supportsDest, err := p.chainSupport.SupportsDestChain(p.oracleID)
	if err != nil {
		lggr.Errorw("error checking if destination chain is supported", "err", err)
		return false, fmt.Errorf("checking if destination chain is supported: %w", err)
	}
	if !supportsDest {
		lggr.Errorw("dest chain not supported by this oracle, can't run report acceptance procedures, " +
			"transmission schedule is wrong - check CCIPHome chainConfigs and ensure that the right oracles are " +
			"assigned as readers of the destination chain, or if " +
			"this oracle should support the destination chain but isn't!")
		return false, plugincommon.NewErrInvalidReport("destination chain not supported")
	}

	sourceChains := slicelib.Map(allRoots,
		func(r cciptypes.MerkleRootChain) cciptypes.ChainSelector { return r.ChainSel })

	isCursed, err := plugincommon.IsReportCursed(ctx, lggr, p.ccipReader, sourceChains)
	if err != nil {
		lggr.Errorw("report not accepted due to curse checking error", "err", err)
		return false, err
	}
	return isCursed, nil
}

func (p *Plugin) ShouldTransmitAcceptedReport(
	ctx context.Context, seqNr uint64, r ocr3types.ReportWithInfo[[]byte],
) (bool, error) {
	ctx, lggr := logutil.WithOCRInfo(ctx, p.lggr, seqNr, logutil.PhaseShouldTransmit)

	decodedReport, err := p.validateReport(ctx, lggr, seqNr, r)
	if errors.Is(err, plugincommon.ErrStaleReport) {
		lggr.Infow("stale report, not accepting", "err", err)
		return false, nil
	}
	if err != nil {
		lggr.Errorw("validation error", "err", err)
		return false, fmt.Errorf("validating report: %w", err)
	}

	lggr.Infow("ShouldTransmitAcceptedReport passed checks",
		"seqNr", seqNr,
		"timestamp", time.Now().UTC(),
		"blessedRootsLen", len(decodedReport.BlessedMerkleRoots),
		"unblessedRootsLen", len(decodedReport.UnblessedMerkleRoots),
		"tokenPriceUpdatesLen", len(decodedReport.PriceUpdates.TokenPriceUpdates),
		"gasPriceUpdatesLen", len(decodedReport.PriceUpdates.GasPriceUpdates),
	)
	return true, nil
}
