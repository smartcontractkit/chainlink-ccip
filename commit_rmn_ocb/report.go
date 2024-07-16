package commitrmnocb

import (
	"context"
	"fmt"

	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"
)

func (p *Plugin) Reports(seqNr uint64, outcomeBytes ocr3types.Outcome) ([]ocr3types.ReportWithInfo[[]byte], error) {
	outcome, err := DecodeCommitPluginOutcome(outcomeBytes)
	if err != nil {
		// TODO: metrics
		p.log.Errorw("failed to decode CommitPluginOutcome", "outcomeBytes", outcomeBytes, "err", err)
		return nil, fmt.Errorf("failed to decode CommitPluginOutcome: %w", err)
	}

	report := CommitPluginReport{
		SignedRoots: outcome.SignedRootsToReport,
		GasPrices:   outcome.GasPrices,
		TokenPrices: outcome.TokenPrices,
	}

	p.log.Infof("Generated report: %v", report)

	// TODO: metrics

	encodedReport, err := report.Encode()
	if err != nil {
		return nil, fmt.Errorf("encode commit plugin report: %w", err)
	}

	return []ocr3types.ReportWithInfo[[]byte]{{Report: encodedReport, Info: nil}}, nil
}

func (p *Plugin) ShouldAcceptAttestedReport(
	_ context.Context, _ uint64, r ocr3types.ReportWithInfo[[]byte],
) (bool, error) {
	decodedReport, err := DecodeCommitPluginReport(r.Report)
	if err != nil {
		// TODO: metrics
		p.log.Errorw("failed to decode CommitPluginOutcome", "outcomeBytes", r.Report, "err", err)
		return false, err
	}

	if decodedReport.IsEmpty() {
		// TODO: metrics
		p.log.Warnf("found an empty report")
		return false, nil
	}

	return true, nil
}

func (p *Plugin) ShouldTransmitAcceptedReport(
	_ context.Context, _ uint64, r ocr3types.ReportWithInfo[[]byte],
) (bool, error) {
	destChainSupported, err := p.supportsDestChain(p.nodeID)
	if err != nil {
		return false, fmt.Errorf("call to supportsDestChain failed: %w", err)
	}
	if !destChainSupported {
		p.log.Debugw("oracle does not support dest chain, skipping report transmission")
		return false, nil
	}

	decodedReport, err := DecodeCommitPluginReport(r.Report)
	if err != nil {
		return false, fmt.Errorf("decode commit plugin report: %w", err)
	}

	// TODO: metrics
	p.log.Debugw("transmitting report",
		"signedRoots", len(decodedReport.SignedRoots),
		"tokenPrices", len(decodedReport.TokenPrices),
		"gasPriceUpdates", len(decodedReport.GasPrices),
	)
	return true, nil
}
