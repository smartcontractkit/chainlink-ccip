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
	config pluginconfig.CommitOffchainConfig,
) ([]ocr3types.ReportPlus[[]byte], error)

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

	// MaxPricesPerReport includes MaxMerkleRootsPerReport, so check it first.
	if config.MaxPricesPerReport != 0 {
		return buildMultiplePriceAndMerkleRootReports, nil
	}

	if config.MaxMerkleRootsPerReport != 0 {
		return buildMultipleMerkleRootReports, nil
	}

	// Default to the standard report.
	return buildStandardReport, nil
}

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
	_ pluginconfig.CommitOffchainConfig,
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

// buildMultiplePriceReports builds many reports of with at most maxMerkleRootsPerReport roots.
func buildMultiplePriceAndMerkleRootReports(
	ctx context.Context,
	lggr logger.Logger,
	reportCodec cciptypes.CommitPluginCodec,
	transmissionSchedule *ocr3types.TransmissionSchedule,
	outcome committypes.Outcome,
	config pluginconfig.CommitOffchainConfig,
) ([]ocr3types.ReportPlus[[]byte], error) {
	var reports []ocr3types.ReportPlus[[]byte]

	// 1. Build price reports and remove prices from outcome.
	maxPrices := config.MaxPricesPerReport
	reports, err := buildMultiplePriceReports(ctx, lggr, reportCodec, transmissionSchedule, outcome, maxPrices)
	if err != nil {
		return nil, fmt.Errorf("problem building price reports: %w", err)
	}
	outcome.TokenPriceOutcome.TokenPrices = make(cciptypes.TokenPriceMap)
	outcome.ChainFeeOutcome.GasPrices = nil

	// 2. Select algorithm for merkle root reports, and build merkle root report(s).
	var rootReportBuilder ReportBuilderFunc
	if config.MaxMerkleRootsPerReport == 0 {
		rootReportBuilder = buildStandardReport
	} else {
		rootReportBuilder = buildMultipleMerkleRootReports
	}
	rootReports, err := rootReportBuilder(ctx, lggr, reportCodec, transmissionSchedule, outcome, config)
	if err != nil {
		return nil, fmt.Errorf("problem building merkle root reports: %w", err)
	}
	// Add merkle root reports to price reports.
	reports = append(reports, rootReports...)

	return reports, nil
}

// buildMultipleMerkleRootReports builds many reports of with at most maxMerkleRootsPerReport roots.
// Any price reports in the outcome are included in the first merkle root.
func buildMultipleMerkleRootReports(
	ctx context.Context,
	lggr logger.Logger,
	reportCodec cciptypes.CommitPluginCodec,
	transmissionSchedule *ocr3types.TransmissionSchedule,
	outcome committypes.Outcome,
	config pluginconfig.CommitOffchainConfig,
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

		if numRoots == config.MaxMerkleRootsPerReport {
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

// buildMultiplePriceReports is a helper to splits up price data into multiple reports.
// Helper for buildMultiplePriceAndMerkleRootReports.
// Merkle root data is ignored.
func buildMultiplePriceReports(
	ctx context.Context,
	lggr logger.Logger,
	reportCodec cciptypes.CommitPluginCodec,
	transmissionSchedule *ocr3types.TransmissionSchedule,
	outcome committypes.Outcome,
	maxPricesPerReport uint64, // passed in directly to avoid implementing ReportBuilderFunc
) ([]ocr3types.ReportPlus[[]byte], error) {
	// update joins together the different types of price updates so that they can be
	// selected in one loop.
	type update struct {
		cciptypes.TokenPrice
		cciptypes.GasPriceChain
	}
	// Helper to add an update to the PriceUpdates type.
	add := func(updates cciptypes.PriceUpdates, update update) cciptypes.PriceUpdates {
		if (update.TokenPrice != cciptypes.TokenPrice{}) {
			updates.TokenPriceUpdates = append(updates.TokenPriceUpdates, update.TokenPrice)
		}
		if (update.GasPriceChain != cciptypes.GasPriceChain{}) {
			updates.GasPriceUpdates = append(updates.GasPriceUpdates, update.GasPriceChain)
		}
		return updates
	}

	var updates []update
	for _, tokenPriceUpdate := range outcome.TokenPriceOutcome.TokenPrices.ToSortedSlice() {
		updates = append(updates, update{
			TokenPrice: tokenPriceUpdate,
		})
	}
	for _, gasPriceUpdate := range outcome.ChainFeeOutcome.GasPrices {
		updates = append(updates, update{
			GasPriceChain: gasPriceUpdate,
		})
	}

	var reports []ocr3types.ReportPlus[[]byte]
	numUpdates := uint64(0)
	priceUpdates := cciptypes.PriceUpdates{}

	for _, u := range updates {
		numUpdates++
		priceUpdates = add(priceUpdates, u)

		if numUpdates == maxPricesPerReport {
			report, err := buildOneReport(
				ctx,
				lggr,
				reportCodec,
				transmissionSchedule,
				outcome.MerkleRootOutcome.OutcomeType,
				nil,
				nil,
				nil, // no RMN for partial reports.
				0,   // no RMN for partial reports.
				priceUpdates,
			)
			if err != nil {
				return nil, fmt.Errorf("buildingMultiplePriceReports: priceUpdates(%+v): %w", priceUpdates, err)
			}
			if report == nil {
				return nil,
					fmt.Errorf("buildingMultiplePriceReports: unexpected empty report for updates(%+v): %w",
						priceUpdates, err)
			}

			reports = append(reports, *report)

			// reset accumulators for next report.
			numUpdates = 0

			// price updates are only included in the first report.
			priceUpdates = cciptypes.PriceUpdates{}
		}
	}

	// check for final partial report.
	if numUpdates > 0 {
		report, err := buildOneReport(
			ctx,
			lggr,
			reportCodec,
			transmissionSchedule,
			outcome.MerkleRootOutcome.OutcomeType,
			nil,
			nil,
			nil, // no RMN for partial reports.
			0,   // no RMN for partial reports.
			priceUpdates,
		)
		if err != nil {
			return nil, fmt.Errorf("buildingMultiplePriceReports: priceUpdates(%+v): %w", priceUpdates, err)
		}
		if report == nil {
			return nil,
				fmt.Errorf("buildingMultiplePriceReports: unexpected empty report for updates(%+v): %w",
					priceUpdates, err)
		}

		reports = append(reports, *report)
	}

	return reports, nil
}
