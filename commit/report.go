package commit

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"

	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"

	"github.com/smartcontractkit/chainlink-ccip/commit/merkleroot"
	"github.com/smartcontractkit/chainlink-ccip/pkg/consts"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
)

// ReportInfo is the info data that will be sent with the along with the report
// It will be used to determine if the report should be accepted or not
type ReportInfo struct {
	// MinSigners is the minimum number of RMN signatures required for the report to be accepted
	MinSigners uint64 `json:"minSigners"`
}

func (ri ReportInfo) Encode() ([]byte, error) {
	return json.Marshal(ri)
}

// decode should be used to decode the report info
func (ri *ReportInfo) Decode(encodedReportInfo []byte) error {
	return json.Unmarshal(encodedReportInfo, ri)
}

func (p *Plugin) Reports(seqNr uint64, outcomeBytes ocr3types.Outcome) ([]ocr3types.ReportWithInfo[[]byte], error) {
	outcome, err := DecodeOutcome(outcomeBytes)
	if err != nil {
		// TODO: metrics
		p.lggr.Errorw("failed to decode Outcome", "outcomeBytes", outcomeBytes, "err", err)
		return nil, fmt.Errorf("failed to decode Outcome (%s): %w", hex.EncodeToString(outcomeBytes), err)
	}

	// Until we start adding tokens and gas to the report, we don't need to report anything
	if outcome.MerkleRootOutcome.OutcomeType != merkleroot.ReportGenerated {
		return []ocr3types.ReportWithInfo[[]byte]{}, nil
	}
	p.lggr.Infow("generating report",
		"roots", outcome.MerkleRootOutcome.RootsToReport,
		"tokenPriceUpdates", outcome.TokenPriceOutcome.TokenPrices,
		"gasPriceUpdates", outcome.ChainFeeOutcome.GasPrices,
		"rmnSignatures", outcome.MerkleRootOutcome.RMNReportSignatures,
		"rmnRawVs", outcome.MerkleRootOutcome.RMNRawVs,
	)

	rep := cciptypes.CommitPluginReport{
		MerkleRoots: outcome.MerkleRootOutcome.RootsToReport,
		PriceUpdates: cciptypes.PriceUpdates{
			TokenPriceUpdates: outcome.TokenPriceOutcome.TokenPrices,
			GasPriceUpdates:   outcome.ChainFeeOutcome.GasPrices,
		},
		RMNSignatures: outcome.MerkleRootOutcome.RMNReportSignatures,
		RMNRawVs:      outcome.MerkleRootOutcome.RMNRawVs,
	}

	encodedReport, err := p.reportCodec.Encode(context.Background(), rep)
	if err != nil {
		return nil, fmt.Errorf("encode commit plugin report: %w", err)
	}

	// Prepare the info data
	reportInfo := ReportInfo{
		MinSigners: outcome.MerkleRootOutcome.RMNRemoteCfg.MinSigners,
	}

	// Serialize reportInfo to []byte
	infoBytes, err := reportInfo.Encode()
	if err != nil {
		return nil, fmt.Errorf("encode report info: %w", err)
	}

	return []ocr3types.ReportWithInfo[[]byte]{{Report: encodedReport, Info: infoBytes}}, nil
}

func (p *Plugin) ShouldAcceptAttestedReport(
	ctx context.Context, u uint64, r ocr3types.ReportWithInfo[[]byte],
) (bool, error) {
	decodedReport, err := p.reportCodec.Decode(ctx, r.Report)
	if err != nil {
		return false, fmt.Errorf("decode commit plugin report: %w", err)
	}

	isEmpty := decodedReport.IsEmpty()
	if isEmpty {
		p.lggr.Infow("skipping empty report")
		return false, nil
	}

	var reportInfo ReportInfo
	if err := reportInfo.Decode(r.Info); err != nil {
		return false, fmt.Errorf("decode report info: %w", err)
	}

	if p.offchainCfg.RMNEnabled &&
		len(decodedReport.MerkleRoots) > 0 &&
		len(decodedReport.RMNSignatures) < int(reportInfo.MinSigners) {
		p.lggr.Infow("skipping report with insufficient RMN signatures %d < %d",
			len(decodedReport.RMNSignatures), reportInfo.MinSigners)
		return false, nil
	}

	return true, nil
}

func (p *Plugin) ShouldTransmitAcceptedReport(
	ctx context.Context, u uint64, r ocr3types.ReportWithInfo[[]byte],
) (bool, error) {
	isWriter, err := p.chainSupport.SupportsDestChain(p.oracleID)
	if err != nil {
		return false, fmt.Errorf("can't know if it's a writer: %w", err)
	}
	if !isWriter {
		p.lggr.Infow("not a writer, skipping report transmission")
		return false, nil
	}

	// we only transmit reports if we are the "active" instance.
	// we can check this by reading the OCR conigs home chain.
	isCandidate, err := p.isCandidateInstance(ctx)
	if err != nil {
		return false, fmt.Errorf("isCandidateInstance: %w", err)
	}

	if isCandidate {
		p.lggr.Infow("not the active instance, skipping report transmission")
		return false, nil
	}

	decodedReport, err := p.reportCodec.Decode(ctx, r.Report)
	if err != nil {
		return false, fmt.Errorf("decode commit plugin report: %w", err)
	}

	isValid, err := merkleroot.ValidateMerkleRootsState(ctx, p.lggr, decodedReport, p.ccipReader)
	if !isValid {
		return false, nil
	}
	if err != nil {
		return false, fmt.Errorf("validate merkle roots state: %w", err)
	}

	p.lggr.Infow("transmitting report",
		"roots", len(decodedReport.MerkleRoots),
		"tokenPriceUpdates", len(decodedReport.PriceUpdates.TokenPriceUpdates),
		"gasPriceUpdates", len(decodedReport.PriceUpdates.GasPriceUpdates),
	)
	return true, nil
}

func (p *Plugin) isCandidateInstance(ctx context.Context) (bool, error) {
	ocrConfigs, err := p.homeChain.GetOCRConfigs(ctx, p.donID, consts.PluginTypeCommit)
	if err != nil {
		return false, fmt.Errorf("failed to get ocr configs from home chain: %w", err)
	}

	return len(ocrConfigs) == 2 && ocrConfigs[1].ConfigDigest == p.reportingCfg.ConfigDigest, nil
}
