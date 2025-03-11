package commit

import (
	"bytes"
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"sort"
	"time"

	"golang.org/x/exp/maps"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"

	"github.com/smartcontractkit/chainlink-ccip/commit/committypes"
	"github.com/smartcontractkit/chainlink-ccip/commit/merkleroot"
	"github.com/smartcontractkit/chainlink-ccip/internal/libs/slicelib"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon/consensus"
	"github.com/smartcontractkit/chainlink-ccip/pkg/consts"
	"github.com/smartcontractkit/chainlink-ccip/pkg/logutil"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

// reportBuilderFunc is used to inject different algorithms for building commit reports.
type reportBuilderFunc func(
	ctx context.Context,
	lggr logger.Logger,
	reportCodec cciptypes.CommitPluginCodec,
	transmissionSchedule *ocr3types.TransmissionSchedule,
	outcome committypes.Outcome,
	maxMerkleRootsPerReport uint64,
) ([]ocr3types.ReportPlus[[]byte], error)

// buildOneReport is the common logic for building a report. Different report building
// algorithms can reassemble reports by selecting which blessed and unblessed merkle
// roots to include in the report, and which price updates and rmn signatures to use.
func buildOneReport(
	ctx context.Context,
	lggr logger.Logger,
	reportCodec cciptypes.CommitPluginCodec,
	transmissionSchedule *ocr3types.TransmissionSchedule,
	outcomeType merkleroot.OutcomeType,
	blessedMerkleRoots []cciptypes.MerkleRootChain,
	unblessedMerkleRoots []cciptypes.MerkleRootChain,
	rmnSignatures []cciptypes.RMNECDSASignature,
	rmnRemoteFSign uint64,
	priceUpdates cciptypes.PriceUpdates,
) (*ocr3types.ReportPlus[[]byte], error) {
	var (
		rep     cciptypes.CommitPluginReport
		repInfo cciptypes.CommitReportInfo
	)

	// MerkleRoots and RMNSignatures will be empty arrays if there is nothing to report
	rep = cciptypes.CommitPluginReport{
		BlessedMerkleRoots:   blessedMerkleRoots,
		UnblessedMerkleRoots: unblessedMerkleRoots,
		PriceUpdates:         priceUpdates,
		RMNSignatures:        rmnSignatures,
	}

	if outcomeType == merkleroot.ReportEmpty {
		rep.BlessedMerkleRoots = []cciptypes.MerkleRootChain{}
		rep.UnblessedMerkleRoots = []cciptypes.MerkleRootChain{}
		rep.RMNSignatures = []cciptypes.RMNECDSASignature{}
	}

	if outcomeType == merkleroot.ReportGenerated {
		allRoots := append(blessedMerkleRoots, unblessedMerkleRoots...)
		sort.Slice(allRoots, func(i, j int) bool { return allRoots[i].ChainSel < allRoots[j].ChainSel })
		repInfo = cciptypes.CommitReportInfo{
			RemoteF:     rmnRemoteFSign,
			MerkleRoots: allRoots,
		}
	}
	// in case of a price-only update, add prices regardless of outcome type.
	repInfo.PriceUpdates = rep.PriceUpdates

	if rep.IsEmpty() {
		lggr.Infow("empty report", "report", rep)
		return nil, nil
	}

	encodedReport, err := reportCodec.Encode(ctx, rep)
	if err != nil {
		return nil, fmt.Errorf("encode commit plugin report: %w", err)
	}

	encodedInfo, err := repInfo.Encode()
	if err != nil {
		return nil, fmt.Errorf("encode commit plugin report info: %w", err)
	}

	lggr.Infow("commit plugin generated reports", "report", rep, "reportInfo", repInfo)

	return &ocr3types.ReportPlus[[]byte]{
		ReportWithInfo: ocr3types.ReportWithInfo[[]byte]{
			Report: encodedReport,
			Info:   encodedInfo,
		},
		TransmissionScheduleOverride: transmissionSchedule,
	}, nil
}

func buildStandardReport(
	ctx context.Context,
	lggr logger.Logger,
	reportCodec cciptypes.CommitPluginCodec,
	transmission *ocr3types.TransmissionSchedule,
	outcome committypes.Outcome,
	_ uint64,
) ([]ocr3types.ReportPlus[[]byte], error) {
	blessedMerkleRoots := make([]cciptypes.MerkleRootChain, 0)
	unblessedMerkleRoots := make([]cciptypes.MerkleRootChain, 0)

	for _, r := range outcome.MerkleRootOutcome.RootsToReport {
		if outcome.MerkleRootOutcome.RMNEnabledChains[r.ChainSel] {
			blessedMerkleRoots = append(blessedMerkleRoots, r)
		} else {
			unblessedMerkleRoots = append(unblessedMerkleRoots, r)
		}
	}

	priceUpdates := cciptypes.PriceUpdates{
		TokenPriceUpdates: outcome.TokenPriceOutcome.TokenPrices.ToSortedSlice(),
		GasPriceUpdates:   outcome.ChainFeeOutcome.GasPrices,
	}

	report, err := buildOneReport(
		ctx,
		lggr,
		reportCodec,
		transmission,
		outcome.MerkleRootOutcome.OutcomeType,
		blessedMerkleRoots,
		unblessedMerkleRoots,
		outcome.MerkleRootOutcome.RMNReportSignatures,
		outcome.MerkleRootOutcome.RMNRemoteCfg.FSign,
		priceUpdates,
	)
	if err != nil {
		return nil, err
	}
	if report != nil {
		return []ocr3types.ReportPlus[[]byte]{*report}, nil
	}
	return nil, nil
}

// buildMultipleReports builds many reports of with at most maxMerkleRootsPerReport roots.
func buildMultipleReports(
	ctx context.Context,
	lggr logger.Logger,
	reportCodec cciptypes.CommitPluginCodec,
	transmissionSchedule *ocr3types.TransmissionSchedule,
	outcome committypes.Outcome,
	maxMerkleRootsPerReport uint64,
) ([]ocr3types.ReportPlus[[]byte], error) {
	var reports []ocr3types.ReportPlus[[]byte]

	numRoots := uint64(0)
	blessedMerkleRoots := make([]cciptypes.MerkleRootChain, 0)
	unblessedMerkleRoots := make([]cciptypes.MerkleRootChain, 0)

	priceUpdates := cciptypes.PriceUpdates{
		TokenPriceUpdates: outcome.TokenPriceOutcome.TokenPrices.ToSortedSlice(),
		GasPriceUpdates:   outcome.ChainFeeOutcome.GasPrices,
	}

	for _, r := range outcome.MerkleRootOutcome.RootsToReport {

		// TODO: Support RMN.
		/*
			if outcome.MerkleRootOutcome.RMNEnabledChains[r.ChainSel] {
				blessedMerkleRoots = append(blessedMerkleRoots, r)
			} else {
				unblessedMerkleRoots = append(unblessedMerkleRoots, r)
			}
		*/
		unblessedMerkleRoots = append(unblessedMerkleRoots, r)
		numRoots++

		if numRoots == maxMerkleRootsPerReport {
			report, err := buildOneReport(
				ctx,
				lggr,
				reportCodec,
				transmissionSchedule,
				outcome.MerkleRootOutcome.OutcomeType,
				blessedMerkleRoots,
				unblessedMerkleRoots,
				nil, // no RMN for partial reports.
				0,   // no RMN for partial reports.
				priceUpdates,
			)
			if err != nil {
				return nil, err
			}
			if report != nil {
				reports = append(reports, *report)
			}

			// reset accumulators for next report.
			numRoots = 0
			blessedMerkleRoots = make([]cciptypes.MerkleRootChain, 0)
			unblessedMerkleRoots = make([]cciptypes.MerkleRootChain, 0)

			// price updates are only included in the first report.
			priceUpdates = cciptypes.PriceUpdates{}
		}
	}

	// check for final partial report.
	if numRoots > 0 {
		report, err := buildOneReport(
			ctx,
			lggr,
			reportCodec,
			transmissionSchedule,
			outcome.MerkleRootOutcome.OutcomeType,
			blessedMerkleRoots,
			unblessedMerkleRoots,
			nil, // no RMN for partial reports.
			0,   // no RMN for partial reports.
			priceUpdates,
		)
		if err != nil {
			return nil, err
		}
		if report != nil {
			reports = append(reports, *report)
		}
	}

	return reports, nil
}

// buildTruncatedReport returns a single truncated report.
func buildTruncatedReport(
	ctx context.Context,
	lggr logger.Logger,
	reportCodec cciptypes.CommitPluginCodec,
	transmissionSchedule *ocr3types.TransmissionSchedule,
	outcome committypes.Outcome,
	maxMerkleRootsPerReport uint64,
) ([]ocr3types.ReportPlus[[]byte], error) {
	reports, err := buildMultipleReports(ctx, lggr, reportCodec, transmissionSchedule, outcome, maxMerkleRootsPerReport)
	if err != nil {
		return nil, err
	}
	if len(reports) > 1 {
		lggr.Warnf("buildTruncatedReport: truncating report %d -> 1", len(reports))
		return reports[:1], nil
	}

	return reports, nil
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

	maxRoots := p.offchainCfg.MaxMerkleRootsPerReport
	reports, err := p.reportBuilder(ctx, lggr, p.reportCodec, transmissionSchedule, outcome, maxRoots)

	if err != nil {
		return nil, fmt.Errorf("error while building reports: %w", err)
	}
	if len(reports) != 0 {
		lggr.Infof("built %d reports", len(reports))
	}

	return reports, nil
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
		lggr.Infow("report not valid, transmitting", "err", err)
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
