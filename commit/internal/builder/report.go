package builder

import (
	"context"
	"fmt"
	"sort"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"

	"github.com/smartcontractkit/chainlink-ccip/commit/committypes"
	"github.com/smartcontractkit/chainlink-ccip/commit/merkleroot"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
	"github.com/smartcontractkit/chainlink-ccip/pluginconfig"
)

// ReportBuilderFunc is used to inject different algorithms for building commit reports.
type ReportBuilderFunc func(
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

// NewReportBuilder returns a ReportBuilderFunc based on the provided config.
func NewReportBuilder(config pluginconfig.CommitOffchainConfig) (ReportBuilderFunc, error) {
	if err := config.Validate(); err != nil {
		return nil, fmt.Errorf("invalid commit offchain config: %w", err)
	}

	// These options were added to allow for more flexibility around report building. For example Solana
	// only supports a single merkle root per report.

	// Currently RMN is only supported for standard reports. Because the config is validated,
	// we can assume that if RMN is enabled, we are building a standard report.
	if config.RMNEnabled {
		return buildStandardReport, nil
	}

	//if config.MaxPricesPerReport != 0 {
	//	return buildMultiplePriceReports, nil
	//}
	if config.MultipleReportsEnabled {
		return buildMultipleReports, nil
	}

	// Default to the standard report.
	return buildStandardReport, nil
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

/*
// buildMultiplePricesReports builds many reports of with at most maxMerkleRootsPerReport roots.
func buildMultiplePricesReports(
	ctx context.Context,
	lggr logger.Logger,
	reportCodec cciptypes.CommitPluginCodec,
	transmissionSchedule *ocr3types.TransmissionSchedule,
	outcome committypes.Outcome,
	maxMerkleRootsPerReport uint64,
) ([]ocr3types.ReportPlus[[]byte], error) {
	return nil, nil
}
*/

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
