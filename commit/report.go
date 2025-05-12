package commit

import (
	"bytes"
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"time"

	"golang.org/x/exp/maps"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"

	"github.com/smartcontractkit/chainlink-ccip/commit/internal/builder"
	"github.com/smartcontractkit/chainlink-ccip/commit/merkleroot"
	"github.com/smartcontractkit/chainlink-ccip/internal/libs/slicelib"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon/consensus"
	"github.com/smartcontractkit/chainlink-ccip/pkg/consts"
	"github.com/smartcontractkit/chainlink-ccip/pkg/logutil"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
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

		lggr.Infow("encoding report and report info",
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

	lggr.Infow("generating report",
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
		lggr.Infof("report with insufficient RMN signatures %d < %d+1",
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
		lggr.Warnw("dest chain not supported, can't run report acceptance procedures")
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

	latestPriceSeqNr, err := p.ccipReader.GetLatestPriceSeqNr(ctx)
	if err != nil {
		return cciptypes.CommitPluginReport{},
			plugincommon.NewErrValidatingReport(fmt.Errorf("get latest price seq nr: %w", err))
	}

	if p.isStaleReport(lggr, seqNr, latestPriceSeqNr, decodedReport) {
		return cciptypes.CommitPluginReport{}, plugincommon.NewErrInvalidReport("stale report")
	}

	err = merkleroot.ValidateMerkleRootsState(
		ctx,
		decodedReport.BlessedMerkleRoots,
		decodedReport.UnblessedMerkleRoots,
		p.ccipReader,
	)
	if err != nil {
		lggr.Infow("report reached transmission protocol but not transmitted, invalid merkle roots state",
			"err", err,
			"blessedMerkleRoots", decodedReport.BlessedMerkleRoots,
			"unblessedMerkleRoots", decodedReport.UnblessedMerkleRoots)
		err = plugincommon.NewErrInvalidReport(fmt.Sprintf("invalid merkle roots state %v", err))
		return cciptypes.CommitPluginReport{}, err
	}

	return decodedReport, nil
}

func (p *Plugin) ShouldAcceptAttestedReport(
	ctx context.Context, seqNr uint64, r ocr3types.ReportWithInfo[[]byte],
) (bool, error) {
	ctx, lggr := logutil.WithOCRInfo(ctx, p.lggr, seqNr, logutil.PhaseShouldAccept)

	decodedReport, err := p.validateReport(ctx, lggr, seqNr, r)
	if errors.Is(err, plugincommon.ErrInvalidReport) {
		lggr.Infow("report not valid, not accepting", "err", err)
		return false, nil
	}
	if err != nil {
		lggr.Infow("validation error", "err", err)
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
	lggr logger.Logger,
	seqNr,
	latestPriceSeqNr uint64,
	decodedReport cciptypes.CommitPluginReport,
) bool {
	if seqNr <= latestPriceSeqNr &&
		len(decodedReport.BlessedMerkleRoots) == 0 &&
		len(decodedReport.UnblessedMerkleRoots) == 0 {
		lggr.Infow(
			"skipping stale report due to stale price seq nr and no merkle roots",
			"latestPriceSeqNr", latestPriceSeqNr)
		return true
	}
	return false
}

func (p *Plugin) checkReportCursed(
	ctx context.Context,
	lggr logger.Logger,
	decodedReport cciptypes.CommitPluginReport,
) (bool, error) {
	allRoots := append(decodedReport.BlessedMerkleRoots, decodedReport.UnblessedMerkleRoots...)

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
	if errors.Is(err, plugincommon.ErrInvalidReport) {
		lggr.Infow("report not valid, not transmitting", "err", err)
		return false, nil
	}
	if err != nil {
		lggr.Infow("validation error", "err", err)
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
