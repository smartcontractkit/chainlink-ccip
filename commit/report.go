package commit

import (
	"bytes"
	"context"
	"encoding/hex"
	"fmt"
	"sort"
	"time"

	"golang.org/x/exp/maps"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"

	"github.com/smartcontractkit/chainlink-ccip/commit/merkleroot"
	"github.com/smartcontractkit/chainlink-ccip/internal/libs/slicelib"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon/consensus"
	"github.com/smartcontractkit/chainlink-ccip/pkg/consts"
	"github.com/smartcontractkit/chainlink-ccip/pkg/logutil"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

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

	var (
		rep     cciptypes.CommitPluginReport
		repInfo cciptypes.CommitReportInfo
	)

	blessedMerkleRoots := make([]cciptypes.MerkleRootChain, 0)
	unblessedMerkleRoots := make([]cciptypes.MerkleRootChain, 0)

	for _, r := range outcome.MerkleRootOutcome.RootsToReport {
		if outcome.MerkleRootOutcome.RMNEnabledChains[r.ChainSel] {
			blessedMerkleRoots = append(blessedMerkleRoots, r)
		} else {
			unblessedMerkleRoots = append(unblessedMerkleRoots, r)
		}
	}

	// MerkleRoots and RMNSignatures will be empty arrays if there is nothing to report
	rep = cciptypes.CommitPluginReport{
		BlessedMerkleRoots:   blessedMerkleRoots,
		UnblessedMerkleRoots: unblessedMerkleRoots,
		PriceUpdates: cciptypes.PriceUpdates{
			TokenPriceUpdates: outcome.TokenPriceOutcome.TokenPrices.ToSortedSlice(),
			GasPriceUpdates:   outcome.ChainFeeOutcome.GasPrices,
		},
		RMNSignatures: outcome.MerkleRootOutcome.RMNReportSignatures,
	}

	if outcome.MerkleRootOutcome.OutcomeType == merkleroot.ReportEmpty {
		rep.BlessedMerkleRoots = []cciptypes.MerkleRootChain{}
		rep.UnblessedMerkleRoots = []cciptypes.MerkleRootChain{}
		rep.RMNSignatures = []cciptypes.RMNECDSASignature{}
	}

	if outcome.MerkleRootOutcome.OutcomeType == merkleroot.ReportGenerated {
		allRoots := append(blessedMerkleRoots, unblessedMerkleRoots...)
		sort.Slice(allRoots, func(i, j int) bool { return allRoots[i].ChainSel < allRoots[j].ChainSel })
		repInfo = cciptypes.CommitReportInfo{
			RemoteF:     outcome.MerkleRootOutcome.RMNRemoteCfg.FSign,
			MerkleRoots: allRoots,
			TokenPrices: rep.PriceUpdates.TokenPriceUpdates,
		}
	}

	if rep.IsEmpty() {
		lggr.Infow("empty report", "report", rep)
		return []ocr3types.ReportPlus[[]byte]{}, nil
	}

	encodedReport, err := p.reportCodec.Encode(ctx, rep)
	if err != nil {
		return nil, fmt.Errorf("encode commit plugin report: %w", err)
	}

	encodedInfo, err := repInfo.Encode()
	if err != nil {
		return nil, fmt.Errorf("encode commit plugin report info: %w", err)
	}

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

	lggr.Infow("commit plugin generated reports", "report", rep, "reportInfo", repInfo)

	return []ocr3types.ReportPlus[[]byte]{
		{
			ReportWithInfo: ocr3types.ReportWithInfo[[]byte]{
				Report: encodedReport,
				Info:   encodedInfo,
			},
			TransmissionScheduleOverride: transmissionSchedule,
		},
	}, nil
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
) (bool, cciptypes.CommitPluginReport, error) {
	if r.Report == nil {
		lggr.Warn("nil report")
		return false, cciptypes.CommitPluginReport{}, nil
	}

	decodedReport, err := p.decodeReport(ctx, r.Report)
	if err != nil {
		return false, cciptypes.CommitPluginReport{}, fmt.Errorf("decode report: %w, report: %x", err, r.Report)
	}

	var reportInfo cciptypes.CommitReportInfo
	if reportInfo, err = cciptypes.DecodeCommitReportInfo(r.Info); err != nil {
		return false, cciptypes.CommitPluginReport{}, fmt.Errorf("decode report info: %w", err)
	}

	for _, root := range decodedReport.BlessedMerkleRoots {
		if root.MerkleRoot == (cciptypes.Bytes32{}) {
			lggr.Warnw("empty blessed merkle root", "root", root)
			return false, cciptypes.CommitPluginReport{}, nil
		}
		if root.SeqNumsRange.Start() > root.SeqNumsRange.End() {
			lggr.Warnw("invalid seqNumsRange", "blessed root", root)
			return false, cciptypes.CommitPluginReport{}, nil
		}
	}

	for _, root := range decodedReport.UnblessedMerkleRoots {
		if root.MerkleRoot == (cciptypes.Bytes32{}) {
			lggr.Warnw("empty unblessed merkle root", "root", root)
			return false, cciptypes.CommitPluginReport{}, nil
		}
		if root.SeqNumsRange.Start() > root.SeqNumsRange.End() {
			lggr.Warnw("invalid seqNumsRange", "unblessed root", root)
			return false, cciptypes.CommitPluginReport{}, nil
		}
	}

	seen := make(map[cciptypes.RMNECDSASignature]struct{})
	for _, sig := range decodedReport.RMNSignatures {

		if _, ok := seen[sig]; ok {
			lggr.Warnw("duplicate RMN signature", "sig", sig)
			return false, cciptypes.CommitPluginReport{}, nil
		}
		seen[sig] = struct{}{}
	}

	if p.offchainCfg.RMNEnabled &&
		len(decodedReport.BlessedMerkleRoots) > 0 &&
		consensus.LtFPlusOne(int(reportInfo.RemoteF), len(decodedReport.RMNSignatures)) {
		lggr.Infof("report with insufficient RMN signatures %d < %d+1",
			len(decodedReport.RMNSignatures), reportInfo.RemoteF)
		return false, cciptypes.CommitPluginReport{}, nil
	}

	if isCursed, err := p.checkReportCursed(ctx, lggr, decodedReport); err != nil || isCursed {
		return false, cciptypes.CommitPluginReport{}, err
	}
	// check if we support the dest, if not we can't do the checks needed.
	supports, err := p.chainSupport.SupportsDestChain(p.oracleID)
	if err != nil {
		return false, cciptypes.CommitPluginReport{}, fmt.Errorf("supports dest chain: %w", err)
	}

	if !supports {
		lggr.Warnw("dest chain not supported, can't run report acceptance procedures")
		return false, cciptypes.CommitPluginReport{}, nil
	}

	offRampConfigDigest, err := p.ccipReader.GetOffRampConfigDigest(ctx, consts.PluginTypeCommit)
	if err != nil {
		return false, cciptypes.CommitPluginReport{}, fmt.Errorf("get offramp config digest: %w", err)
	}

	if !bytes.Equal(offRampConfigDigest[:], p.reportingCfg.ConfigDigest[:]) {
		lggr.Warnw("my config digest doesn't match offramp's config digest, not accepting report",
			"myConfigDigest", p.reportingCfg.ConfigDigest,
			"offRampConfigDigest", hex.EncodeToString(offRampConfigDigest[:]),
		)
		return false, cciptypes.CommitPluginReport{}, nil
	}

	latestPriceSeqNr, err := p.ccipReader.GetLatestPriceSeqNr(ctx)
	if err != nil {
		return false, cciptypes.CommitPluginReport{}, fmt.Errorf("get latest price seq nr: %w", err)
	}

	if p.isStaleReport(lggr, seqNr, latestPriceSeqNr, decodedReport) {
		return false, cciptypes.CommitPluginReport{}, nil
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
		return false, cciptypes.CommitPluginReport{}, nil
	}

	return true, decodedReport, nil
}

func (p *Plugin) ShouldAcceptAttestedReport(
	ctx context.Context, seqNr uint64, r ocr3types.ReportWithInfo[[]byte],
) (bool, error) {
	ctx, lggr := logutil.WithOCRInfo(ctx, p.lggr, seqNr, logutil.PhaseShouldAccept)

	valid, decodedReport, err := p.validateReport(ctx, lggr, seqNr, r)
	if err != nil {
		return false, fmt.Errorf("validating report: %w", err)
	}

	if !valid {
		lggr.Infow("report is not accepted")
		return false, nil
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

	valid, decodedReport, err := p.validateReport(ctx, lggr, seqNr, r)
	if err != nil {
		return false, fmt.Errorf("validating report: %w", err)
	}

	if !valid {
		lggr.Infow("report not valid, not transmitting")
		return false, nil
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
