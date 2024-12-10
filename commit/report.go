package commit

import (
	"bytes"
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"time"

	"golang.org/x/exp/maps"

	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"

	"github.com/smartcontractkit/chainlink-ccip/commit/merkleroot"
	"github.com/smartcontractkit/chainlink-ccip/internal/libs/slicelib"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon/consensus"
	"github.com/smartcontractkit/chainlink-ccip/pkg/consts"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

const (
	// transmissionDelayMultiplier is used to calculate the transmission delay for each oracle.
	transmissionDelayMultiplier = 3 * time.Second
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
	outcome, err := decodeOutcome(outcomeBytes)
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
	rep = cciptypes.CommitPluginReport{
		MerkleRoots: outcome.MerkleRootOutcome.RootsToReport,
		PriceUpdates: cciptypes.PriceUpdates{
			TokenPriceUpdates: outcome.TokenPriceOutcome.TokenPrices,
			GasPriceUpdates:   outcome.ChainFeeOutcome.GasPrices,
		},
		RMNSignatures: outcome.MerkleRootOutcome.RMNReportSignatures,
	}

	if outcome.MerkleRootOutcome.OutcomeType == merkleroot.ReportGenerated {
		repInfo = ReportInfo{RemoteF: outcome.MerkleRootOutcome.RMNRemoteCfg.F}
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
		transmissionDelayMultiplier,
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

func (p *Plugin) ShouldAcceptAttestedReport(
	ctx context.Context, seqNr uint64, r ocr3types.ReportWithInfo[[]byte],
) (bool, error) {
	// check if we support the dest, if not we can't do the checks needed.
	supports, err := p.chainSupport.SupportsDestChain(p.oracleID)
	if err != nil {
		return false, fmt.Errorf("supports dest chain: %w", err)
	}

	if !supports {
		p.lggr.Warnw("dest chain not supported, can't run report acceptance procedures")

		// TODO: return false here or a non-nil error?
		return false, nil
	}

	offRampConfigDigest, err := p.ccipReader.GetOffRampConfigDigest(ctx, consts.PluginTypeCommit)
	if err != nil {
		return false, fmt.Errorf("get offramp config digest: %w", err)
	}

	if !bytes.Equal(offRampConfigDigest[:], p.reportingCfg.ConfigDigest[:]) {
		p.lggr.Warnw("my config digest doesn't match offramp's config digest, not accepting report",
			"myConfigDigest", p.reportingCfg.ConfigDigest,
			"offRampConfigDigest", hex.EncodeToString(offRampConfigDigest[:]),
		)
		return false, nil
	}

	latestSeqNr, err := p.ccipReader.GetLatestPriceSeqNr(ctx)
	if err != nil {
		return false, fmt.Errorf("get latest price seq nr: %w", err)
	}

	decodedReport, err := p.decodeReport(ctx, r.Report)
	if err != nil || decodedReport.IsEmpty() {
		return false, err
	}

	if p.isStaleReport(seqNr, latestSeqNr, decodedReport) {
		return false, nil
	}

	if isCursed, err := p.checkReportCursed(ctx, decodedReport); err != nil || isCursed {
		return false, err
	}

	return p.validateReportInfo(r.Info, decodedReport)
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
	if seqNr < latestSeqNr && len(decodedReport.MerkleRoots) == 0 {
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

func (p *Plugin) validateReportInfo(info []byte, decodedReport cciptypes.CommitPluginReport) (bool, error) {
	var reportInfo ReportInfo
	if err := reportInfo.Decode(info); err != nil {
		return false, fmt.Errorf("decode report info: %w", err)
	}

	if p.offchainCfg.RMNEnabled &&
		len(decodedReport.MerkleRoots) > 0 &&
		consensus.LtFPlusOne(int(reportInfo.RemoteF), len(decodedReport.RMNSignatures)) {
		p.lggr.Infow("report with insufficient RMN signatures %d < %d+1",
			len(decodedReport.RMNSignatures), reportInfo.RemoteF)
		return false, nil
	}

	return true, nil
}

func (p *Plugin) ShouldTransmitAcceptedReport(
	ctx context.Context, seqNr uint64, r ocr3types.ReportWithInfo[[]byte],
) (bool, error) {
	// check if we support the dest, if not we can't do the checks needed.
	supports, err := p.chainSupport.SupportsDestChain(p.oracleID)
	if err != nil {
		return false, fmt.Errorf("supports dest chain: %w", err)
	}

	if !supports {
		p.lggr.Warnw("dest chain not supported, can't run report acceptance procedures")

		// TODO: return false here or a non-nil error?
		return false, nil
	}

	offRampConfigDigest, err := p.ccipReader.GetOffRampConfigDigest(ctx, consts.PluginTypeCommit)
	if err != nil {
		return false, fmt.Errorf("get offramp config digest: %w", err)
	}

	if !bytes.Equal(offRampConfigDigest[:], p.reportingCfg.ConfigDigest[:]) {
		p.lggr.Warnw("my config digest doesn't match offramp's config digest, not accepting report",
			"myConfigDigest", p.reportingCfg.ConfigDigest,
			"offRampConfigDigest", hex.EncodeToString(offRampConfigDigest[:]),
		)
		return false, nil
	}

	latestSeqNr, err := p.ccipReader.GetLatestPriceSeqNr(ctx)
	if err != nil {
		return false, fmt.Errorf("get latest price seq nr: %w", err)
	}

	decodedReport, err := p.decodeReport(ctx, r.Report)
	if err != nil || decodedReport.IsEmpty() {
		return false, err
	}

	if p.isStaleReport(seqNr, latestSeqNr, decodedReport) {
		return false, nil
	}

	err = merkleroot.ValidateMerkleRootsState(ctx, decodedReport.MerkleRoots, p.ccipReader)
	if err != nil {
		p.lggr.Warnw("report reached transmission protocol but not transmitted, invalid merkle roots state",
			"err", err, "merkleRoots", decodedReport.MerkleRoots)
		return false, nil
	}

	p.lggr.Infow("transmitting report",
		"roots", len(decodedReport.MerkleRoots),
		"tokenPriceUpdates", len(decodedReport.PriceUpdates.TokenPriceUpdates),
		"gasPriceUpdates", len(decodedReport.PriceUpdates.GasPriceUpdates),
	)
	return true, nil
}
