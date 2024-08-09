package commitrmnocb

import (
	"context"
	"fmt"

	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-ccip/commit"
)

func (p *Plugin) Reports(seqNr uint64, outcomeBytes ocr3types.Outcome) ([]ocr3types.ReportWithInfo[[]byte], error) {
	outcome, err := DecodeCommitPluginOutcome(outcomeBytes)
	if err != nil {
		// TODO: metrics
		p.lggr.Errorw("failed to decode Outcome", "outcomeBytes", outcomeBytes, "err", err)
		return nil, fmt.Errorf("failed to decode Outcome: %w", err)
	}

	if outcome.OutcomeType != ReportGenerated {
		return []ocr3types.ReportWithInfo[[]byte]{}, nil
	}

	rep := cciptypes.NewCommitPluginReport(outcome.RootsToReport, outcome.TokenPrices, outcome.GasPrices)

	encodedReport, err := p.reportCodec.Encode(context.Background(), rep)
	if err != nil {
		return nil, fmt.Errorf("encode commit plugin report: %w", err)
	}

	return []ocr3types.ReportWithInfo[[]byte]{{Report: encodedReport, Info: nil}}, nil
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

	return true, nil
}

func (p *Plugin) ShouldTransmitAcceptedReport(
	ctx context.Context, u uint64, r ocr3types.ReportWithInfo[[]byte],
) (bool, error) {
	isWriter, err := p.supportsDestChain(p.nodeID)
	if err != nil {
		return false, fmt.Errorf("can't know if it's a writer: %w", err)
	}
	if !isWriter {
		p.lggr.Infow("not a writer, skipping report transmission")
		return false, nil
	}

	decodedReport, err := p.reportCodec.Decode(ctx, r.Report)
	if err != nil {
		return false, fmt.Errorf("decode commit plugin report: %w", err)
	}

	isValid, err := commit.ValidateMerkleRootsState(ctx, p.lggr, decodedReport, p.ccipReader)
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
