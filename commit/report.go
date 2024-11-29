package commit

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"
	"golang.org/x/exp/maps"

	"github.com/smartcontractkit/chainlink-ccip/commit/merkleroot"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon/consensus"
	"github.com/smartcontractkit/chainlink-ccip/pkg/consts"
	ccipreader "github.com/smartcontractkit/chainlink-ccip/pkg/reader"
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
	latestOnchainSeqNr, err := p.ccipReader.GetLatestPriceSeqNr(ctx)
	if err != nil {
		return false, fmt.Errorf("get latest price seq nr: %w", err)
	}

	decodedReport, err := p.reportCodec.Decode(ctx, r.Report)
	if err != nil {
		return false, fmt.Errorf("decode commit plugin report: %w", err)
	}

	isEmpty := decodedReport.IsEmpty()
	if isEmpty {
		p.lggr.Infow("skipping empty report")
		return false, nil
	}

	p.lggr.Infow("ShouldAcceptAttestedReport",
		"seqNr", seqNr,
		"latestOnchainSeqNr", latestOnchainSeqNr,
		"decodedReport", decodedReport,
	)
	if seqNr < latestOnchainSeqNr && len(decodedReport.MerkleRoots) == 0 {
		p.lggr.Infow("skipping stale report", "seqNr", seqNr, "latestSeqNr", latestOnchainSeqNr)
		return false, nil
	}

	if err := cursingValidation(ctx, p.ccipReader, decodedReport.MerkleRoots); err != nil {
		p.lggr.Errorw("report not accepted due to cursing", "err", err)
		return false, nil
	}

	var reportInfo ReportInfo
	if err := reportInfo.Decode(r.Info); err != nil {
		return false, fmt.Errorf("decode report info: %w", err)
	}

	if p.offchainCfg.RMNEnabled &&
		len(decodedReport.MerkleRoots) > 0 &&
		consensus.LtFPlusOne(int(reportInfo.RemoteF), len(decodedReport.RMNSignatures)) {
		p.lggr.Infow("skipping report with insufficient RMN signatures %d < %d+1",
			len(decodedReport.RMNSignatures), reportInfo.RemoteF)
		return false, nil
	}

	return true, nil
}

func (p *Plugin) ShouldTransmitAcceptedReport(
	ctx context.Context, seqNr uint64, r ocr3types.ReportWithInfo[[]byte],
) (bool, error) {
	// we only transmit reports if we are the "active" instance.
	// we can check this by reading the OCR configs from the home chain.
	isCandidate, err := p.isCandidateInstance(ctx)
	if err != nil {
		return false, fmt.Errorf("isCandidateInstance: %w", err)
	}

	if isCandidate {
		p.lggr.Infow("not the active instance, skipping report transmission")
		return false, nil
	}

	latestOnchainSeqNr, err := p.ccipReader.GetLatestPriceSeqNr(ctx)
	if err != nil {
		return false, fmt.Errorf("get latest price seq nr: %w", err)
	}

	decodedReport, err := p.reportCodec.Decode(ctx, r.Report)
	if err != nil {
		return false, fmt.Errorf("decode commit plugin report: %w", err)
	}

	isEmpty := decodedReport.IsEmpty()
	if isEmpty {
		p.lggr.Infow("skipping empty report")
		return false, nil
	}

	p.lggr.Infow("ShouldTransmitAcceptedReport",
		"seqNr", seqNr,
		"latestOnchainSeqNr", latestOnchainSeqNr,
		"decodedReport", decodedReport,
	)

	if seqNr < latestOnchainSeqNr && len(decodedReport.MerkleRoots) == 0 {
		p.lggr.Infow("skipping stale report", "seqNr", seqNr, "latestOnchainSeqNr", latestOnchainSeqNr)
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

// cursingValidation will make one contract call to get rmn curse info.
// If destination is cursed or some source chain is cursed it returns error.
func cursingValidation(
	ctx context.Context,
	ccipReader ccipreader.CCIPReader,
	merkleRoots []cciptypes.MerkleRootChain,
) error {
	// If merkleRoots are empty we still want to transmit chain fee and token prices.
	// So the report is considered valid.
	if len(merkleRoots) == 0 {
		return nil
	}

	sourceChains := make([]cciptypes.ChainSelector, 0, len(merkleRoots))
	for _, mr := range merkleRoots {
		sourceChains = append(sourceChains, mr.ChainSel)
	}

	curseInfo, err := ccipReader.GetRmnCurseInfo(ctx, sourceChains)
	if err != nil {
		return fmt.Errorf("get rmn curse info sourceChains=%v: %w", sourceChains, err)
	}

	if curseInfo.CursedDestination || curseInfo.GlobalCurse {
		return fmt.Errorf("destination chain is cursed: %v", curseInfo)
	}

	for sourceChain, isCursed := range curseInfo.CursedSourceChains {
		if isCursed {
			return fmt.Errorf("source chain %d is cursed", sourceChain)
		}
	}

	return nil
}

func (p *Plugin) isCandidateInstance(ctx context.Context) (bool, error) {
	ocrConfigs, err := p.homeChain.GetOCRConfigs(ctx, p.donID, consts.PluginTypeCommit)
	if err != nil {
		return false, fmt.Errorf("failed to get ocr configs from home chain: %w", err)
	}

	return ocrConfigs.CandidateConfig.ConfigDigest == p.reportingCfg.ConfigDigest, nil
}
