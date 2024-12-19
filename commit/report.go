package commit

import (
	"bytes"
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"time"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"golang.org/x/exp/maps"

	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"

	"github.com/smartcontractkit/chainlink-ccip/commit/committypes"
	"github.com/smartcontractkit/chainlink-ccip/commit/merkleroot"
	"github.com/smartcontractkit/chainlink-ccip/internal/libs/slicelib"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon/consensus"
	"github.com/smartcontractkit/chainlink-ccip/pkg/consts"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

// ReportInfo is the info data that will be sent with the along with the report
// It will be used to determine if the report should be accepted or not
type ReportInfo struct {
	// RemoteF Max number of faulty RMN nodes; f+1 signers are required to verify a report.
	RemoteF uint64 `json:"remoteF"`
}

func (ri ReportInfo) Encode() ([]byte, error) {
	return json.Marshal(ri)
}

// Decode should be used to decode the report info
func (ri *ReportInfo) Decode(encodedReportInfo []byte) error {
	return json.Unmarshal(encodedReportInfo, ri)
}

func (p *Plugin) Reports(
	ctx context.Context, seqNr uint64, outcomeBytes ocr3types.Outcome,
) ([]ocr3types.ReportPlus[[]byte], error) {
	outcome, err := committypes.DecodeOutcome(outcomeBytes)
	if err != nil {
		p.lggr.Errorw("failed to decode Outcome", "outcome", string(outcomeBytes), "err", err)
		return nil, fmt.Errorf("decode outcome: %w", err)
	}

	p.lggr.Infow("generating report",
		"roots", outcome.MerkleRootOutcome.RootsToReport,
		"tokenPriceUpdates", outcome.TokenPriceOutcome.TokenPrices,
		"gasPriceUpdates", outcome.ChainFeeOutcome.GasPrices,
		"rmnSignatures", outcome.MerkleRootOutcome.RMNReportSignatures,
	)

	var (
		rep     cciptypes.CommitPluginReport
		repInfo ReportInfo
	)

	// MerkleRoots and RMNSignatures will be empty arrays if there is nothing to report
	rep = cciptypes.CommitPluginReport{
		MerkleRoots: outcome.MerkleRootOutcome.RootsToReport,
		PriceUpdates: cciptypes.PriceUpdates{
			TokenPriceUpdates: outcome.TokenPriceOutcome.TokenPrices,
			GasPriceUpdates:   outcome.ChainFeeOutcome.GasPrices,
		},
		RMNSignatures: outcome.MerkleRootOutcome.RMNReportSignatures,
	}

	if outcome.MerkleRootOutcome.OutcomeType == merkleroot.ReportEmpty {
		rep.MerkleRoots = []cciptypes.MerkleRootChain{}
		rep.RMNSignatures = []cciptypes.RMNECDSASignature{}
	}

	if outcome.MerkleRootOutcome.OutcomeType == merkleroot.ReportGenerated {
		repInfo = ReportInfo{RemoteF: outcome.MerkleRootOutcome.RMNRemoteCfg.FSign}
	}

	if rep.IsEmpty() {
		p.lggr.Infow("empty report", "report", rep)
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
	p.lggr.Debugw("transmission schedule override",
		"transmissionSchedule", transmissionSchedule, "oracleIDToP2PID", p.oracleIDToP2PID)

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
	seqNr uint64,
	r ocr3types.ReportWithInfo[[]byte],
) (bool, cciptypes.CommitPluginReport, error) {
	lggr := logger.With(p.lggr, "seqNr", seqNr)

	if r.Report == nil {
		lggr.Warn("nil report")
		return false, cciptypes.CommitPluginReport{}, nil
	}

	decodedReport, err := p.decodeReport(ctx, r.Report)
	if err != nil {
		return false, cciptypes.CommitPluginReport{}, fmt.Errorf("decode report: %w, report: %x", err, r.Report)
	}

	if decodedReport.IsEmpty() {
		lggr.Warnw("empty report after decoding", "decodedReport", decodedReport)
		return false, cciptypes.CommitPluginReport{}, nil
	}

	var reportInfo ReportInfo
	if err := reportInfo.Decode(r.Info); err != nil {
		return false, cciptypes.CommitPluginReport{}, fmt.Errorf("decode report info: %w", err)
	}

	if p.offchainCfg.RMNEnabled &&
		len(decodedReport.MerkleRoots) > 0 &&
		consensus.LtFPlusOne(int(reportInfo.RemoteF), len(decodedReport.RMNSignatures)) {
		lggr.Infof("report with insufficient RMN signatures %d < %d+1",
			len(decodedReport.RMNSignatures), reportInfo.RemoteF)
		return false, cciptypes.CommitPluginReport{}, nil
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

	latestSeqNr, err := p.ccipReader.GetLatestPriceSeqNr(ctx)
	if err != nil {
		return false, cciptypes.CommitPluginReport{}, fmt.Errorf("get latest price seq nr: %w", err)
	}

	if p.isStaleReport(seqNr, latestSeqNr, decodedReport) {
		return false, cciptypes.CommitPluginReport{}, nil
	}

	err = merkleroot.ValidateMerkleRootsState(ctx, decodedReport.MerkleRoots, p.ccipReader)
	if err != nil {
		lggr.Infow("report reached transmission protocol but not transmitted, invalid merkle roots state",
			"err", err, "merkleRoots", decodedReport.MerkleRoots)
		return false, cciptypes.CommitPluginReport{}, nil
	}

	return true, decodedReport, nil
}

func (p *Plugin) ShouldAcceptAttestedReport(
	ctx context.Context, seqNr uint64, r ocr3types.ReportWithInfo[[]byte],
) (bool, error) {
	valid, decodedReport, err := p.validateReport(ctx, seqNr, r)
	if err != nil {
		return false, fmt.Errorf("validating report: %w", err)
	}

	if !valid {
		p.lggr.Infow("report is not accepted", "seqNr", seqNr)
		return false, nil
	}

	// TODO: consider doing this in validateReport,
	// will end up doing it in both ShouldAccept and ShouldTransmit.
	if isCursed, err := p.checkReportCursed(ctx, decodedReport); err != nil || isCursed {
		return false, err
	}

	p.lggr.Infow("ShouldAcceptedAttestedReport passed checks",
		"seqNr", seqNr,
		"timestamp", time.Now().UTC(),
		"rootsLen", len(decodedReport.MerkleRoots),
		"tokenPriceUpdatesLen", len(decodedReport.PriceUpdates.TokenPriceUpdates),
		"gasPriceUpdatesLen", len(decodedReport.PriceUpdates.GasPriceUpdates),
	)
	return true, nil
}

func (p *Plugin) decodeReport(ctx context.Context, report []byte) (cciptypes.CommitPluginReport, error) {
	decodedReport, err := p.reportCodec.Decode(ctx, report)
	if err != nil {
		return cciptypes.CommitPluginReport{}, fmt.Errorf("decode commit plugin report: %w", err)
	}
	if decodedReport.IsEmpty() {
		p.lggr.Infow("empty report")
	}
	return decodedReport, nil
}

func (p *Plugin) isStaleReport(seqNr, latestSeqNr uint64, decodedReport cciptypes.CommitPluginReport) bool {
	if seqNr <= latestSeqNr && len(decodedReport.MerkleRoots) == 0 {
		p.lggr.Infow("skipping stale report", "seqNr", seqNr, "latestSeqNr", latestSeqNr)
		return true
	}
	return false
}

func (p *Plugin) checkReportCursed(ctx context.Context, decodedReport cciptypes.CommitPluginReport) (bool, error) {
	sourceChains := slicelib.Map(decodedReport.MerkleRoots,
		func(r cciptypes.MerkleRootChain) cciptypes.ChainSelector {
			return r.ChainSel
		})
	isCursed, err := plugincommon.IsReportCursed(ctx, p.lggr, p.ccipReader, p.chainSupport.DestChain(), sourceChains)
	if err != nil {
		p.lggr.Errorw("report not accepted due to curse checking error", "err", err)
		return false, err
	}
	return isCursed, nil
}

func (p *Plugin) ShouldTransmitAcceptedReport(
	ctx context.Context, seqNr uint64, r ocr3types.ReportWithInfo[[]byte],
) (bool, error) {
	valid, decodedReport, err := p.validateReport(ctx, seqNr, r)
	if err != nil {
		return false, fmt.Errorf("validating report: %w", err)
	}

	if !valid {
		p.lggr.Infow("report not valid, not transmitting", "seqNr", seqNr)
		return false, nil
	}

	p.lggr.Infow("ShouldTransmitAcceptedReport passed checks",
		"seqNr", seqNr,
		"timestamp", time.Now().UTC(),
		"rootsLen", len(decodedReport.MerkleRoots),
		"tokenPriceUpdatesLen", len(decodedReport.PriceUpdates.TokenPriceUpdates),
		"gasPriceUpdatesLen", len(decodedReport.PriceUpdates.GasPriceUpdates),
	)
	return true, nil
}
