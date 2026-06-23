package builder

import (
	"fmt"
	"sort"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-ccip/commit/committypes"
	"github.com/smartcontractkit/chainlink-ccip/commit/merkleroot"
	"github.com/smartcontractkit/chainlink-ccip/pluginconfig"
)

type Report struct {
	Report     cciptypes.CommitPluginReport
	ReportInfo cciptypes.CommitReportInfo
}

// ReportBuilderFunc is used to inject different algorithms for building commit reports.
//
// This function should only return valid reports. An empty report is not considered valid.
type ReportBuilderFunc func(
	lggr logger.Logger,
	outcome committypes.Outcome,
	config pluginconfig.CommitOffchainConfig,
) ([]Report, error)

// NewReportBuilder returns a ReportBuilderFunc based on the provided config.
func NewReportBuilder(MaxMerkleRootsPerReport, MaxPricesPerReport uint64) (ReportBuilderFunc, error) {
	// These options were added to allow for more flexibility around report building. For example Solana
	// only supports a single merkle root per report.

	// MaxPricesPerReport is a superset of MaxMerkleRootsPerReport, so check it first.
	if MaxPricesPerReport > 0 {
		return buildMultiplePriceAndMerkleRootReports, nil
	}

	if MaxMerkleRootsPerReport > 0 {
		return buildMultipleMerkleRootReports, nil
	}

	// Default to the standard report.
	return buildStandardReport, nil
}

// buildOneReport is the common logic for building a report. Different report building
// algorithms can reassemble reports by selecting which merkle roots to include in the report,
// and which price updates to use.
//
// This function may return empty reports, for example roots are ignored if the
// merkleOutcomeType is not "ReportGenerated".
func buildOneReport(
	lggr logger.Logger,
	merkleOutcomeType merkleroot.OutcomeType,
	unblessedMerkleRoots []cciptypes.MerkleRootChain,
	priceUpdates cciptypes.PriceUpdates,
) Report {
	var (
		rep     cciptypes.CommitPluginReport
		repInfo cciptypes.CommitReportInfo
	)

	// Merkle root data is only included when the outcomeType is "ReportGenerated".
	if merkleOutcomeType == merkleroot.ReportGenerated {
		// MerkleRoots will be empty arrays if there is nothing to report.
		// BlessedMerkleRoots is left empty for onchain interface compatibility.
		rep = cciptypes.CommitPluginReport{
			UnblessedMerkleRoots: unblessedMerkleRoots,
		}

		allRoots := append([]cciptypes.MerkleRootChain(nil), unblessedMerkleRoots...)
		sort.Slice(allRoots, func(i, j int) bool { return allRoots[i].ChainSel < allRoots[j].ChainSel })
		repInfo = cciptypes.CommitReportInfo{
			MerkleRoots: allRoots,
		}
	}

	// Price updates are always allowed.
	rep.PriceUpdates = priceUpdates
	repInfo.PriceUpdates = rep.PriceUpdates

	if rep.IsEmpty() {
		lggr.Warnw("buildOneReport: generated an empty report",
			"merkleOutcomeType", merkleOutcomeType,
			"blessedMerkleRoots", rep.BlessedMerkleRoots,
			"unblessedMerkleRoots", unblessedMerkleRoots,
			"priceUpdates", priceUpdates,
		)

		return Report{}
	}

	return Report{
		Report:     rep,
		ReportInfo: repInfo,
	}
}

// buildStandardReport builds a one report with all the merkle roots and price updates.
func buildStandardReport(
	lggr logger.Logger,
	outcome committypes.Outcome,
	_ pluginconfig.CommitOffchainConfig,
) ([]Report, error) {
	unblessedMerkleRoots := make([]cciptypes.MerkleRootChain, 0, len(outcome.MerkleRootOutcome.RootsToReport))
	unblessedMerkleRoots = append(unblessedMerkleRoots, outcome.MerkleRootOutcome.RootsToReport...)

	priceUpdates := cciptypes.PriceUpdates{
		TokenPriceUpdates: outcome.TokenPriceOutcome.TokenPrices.ToSortedSlice(),
		GasPriceUpdates:   outcome.ChainFeeOutcome.GasPrices,
	}

	report := buildOneReport(
		lggr,
		outcome.MerkleRootOutcome.OutcomeType,
		unblessedMerkleRoots,
		priceUpdates,
	)
	// Do not include empty reports, which may sometimes happen for merkle root reports.
	if report.Report.IsEmpty() {
		return nil, nil
	}
	return []Report{report}, nil
}

// buildMultiplePriceReports builds many reports of with at most maxMerkleRootsPerReport roots.
func buildMultiplePriceAndMerkleRootReports(
	lggr logger.Logger,
	outcome committypes.Outcome,
	config pluginconfig.CommitOffchainConfig,
) ([]Report, error) {
	// 1. Build price reports.
	maxPrices := config.MaxPricesPerReport
	reports := buildMultiplePriceReports(lggr, outcome, maxPrices)
	for _, report := range reports {
		if report.Report.IsEmpty() {
			return nil, fmt.Errorf("err in buildMultiplePriceReports(): price report should not be empty")
		}
	}

	// remove prices from the outcome so that they won't be included in the merkle root reports.
	outcome.TokenPriceOutcome.TokenPrices = make(cciptypes.TokenPriceMap)
	outcome.ChainFeeOutcome.GasPrices = nil

	// 2. Select which algorithm to use for building merkle root reports.
	var rootReportBuilder ReportBuilderFunc
	if config.MaxMerkleRootsPerReport == 0 {
		rootReportBuilder = buildStandardReport
	} else {
		rootReportBuilder = buildMultipleMerkleRootReports
	}

	rootReports, err := rootReportBuilder(lggr, outcome, config)
	if err != nil {
		return nil, fmt.Errorf("err in buildMultiplePriceReports(): problem building merkle root reports: %w", err)
	}
	// Add merkle root reports to price reports.
	for _, rootReport := range rootReports {
		// Ignore empty reports, which may sometimes happen for merkle root reports.
		if !rootReport.Report.IsEmpty() {
			reports = append(reports, rootReport)
		}
	}

	return reports, nil
}

// buildMultipleMerkleRootReports builds many reports of with at most maxMerkleRootsPerReport roots.
// Any price reports in the outcome are included in the first merkle root.
func buildMultipleMerkleRootReports(
	lggr logger.Logger,
	outcome committypes.Outcome,
	config pluginconfig.CommitOffchainConfig,
) ([]Report, error) {
	var reports []Report

	numRoots := uint64(0)
	unblessedMerkleRoots := make([]cciptypes.MerkleRootChain, 0)

	priceUpdates := cciptypes.PriceUpdates{
		TokenPriceUpdates: outcome.TokenPriceOutcome.TokenPrices.ToSortedSlice(),
		GasPriceUpdates:   outcome.ChainFeeOutcome.GasPrices,
	}

	for _, r := range outcome.MerkleRootOutcome.RootsToReport {
		unblessedMerkleRoots = append(unblessedMerkleRoots, r)
		numRoots++

		if numRoots == config.MaxMerkleRootsPerReport {
			report := buildOneReport(
				lggr,
				outcome.MerkleRootOutcome.OutcomeType,
				unblessedMerkleRoots,
				priceUpdates,
			)

			// Do not include empty reports, which may sometimes happen for merkle root reports.
			if !report.Report.IsEmpty() {
				reports = append(reports, report)
			}

			// reset accumulators for next report.
			numRoots = 0
			unblessedMerkleRoots = make([]cciptypes.MerkleRootChain, 0)

			// price updates are only included in the first report.
			priceUpdates = cciptypes.PriceUpdates{}
		}
	}

	// check for final partial report, or a price-only outcome (no roots).
	if numRoots > 0 || len(outcome.MerkleRootOutcome.RootsToReport) == 0 {
		report := buildOneReport(
			lggr,
			outcome.MerkleRootOutcome.OutcomeType,
			unblessedMerkleRoots,
			priceUpdates,
		)

		// Do not include empty reports, which may sometimes happen for merkle root reports.
		if !report.Report.IsEmpty() {
			reports = append(reports, report)
		}
	}

	return reports, nil
}

// buildMultiplePriceReports is a helper to split price data into multiple reports.
// Helper for buildMultiplePriceAndMerkleRootReports.
// Merkle root data is ignored.
func buildMultiplePriceReports(
	lggr logger.Logger,
	outcome committypes.Outcome,
	maxPricesPerReport uint64, // passed in directly to avoid implementing ReportBuilderFunc
) []Report {
	// update is a union of the different types of price updates. This is done so that one loop can
	// create all the reports.
	type update struct {
		cciptypes.TokenPrice
		cciptypes.GasPriceChain
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

	// Build reports
	var reports []Report
	numUpdates := uint64(0)
	priceUpdates := cciptypes.PriceUpdates{}
	for _, u := range updates {
		numUpdates++
		// Get the specific update type and add it to the priceUpdates object.
		if (u.TokenPrice != cciptypes.TokenPrice{}) {
			priceUpdates.TokenPriceUpdates = append(priceUpdates.TokenPriceUpdates, u.TokenPrice)
		}
		if (u.GasPriceChain != cciptypes.GasPriceChain{}) {
			priceUpdates.GasPriceUpdates = append(priceUpdates.GasPriceUpdates, u.GasPriceChain)
		}

		// Build a report when we have enough
		if numUpdates == maxPricesPerReport {
			report := buildOneReport(
				lggr,
				outcome.MerkleRootOutcome.OutcomeType,
				nil,
				priceUpdates,
			)
			reports = append(reports, report)

			// reset accumulators for next report.
			numUpdates = 0

			// price updates are only included in the first report.
			priceUpdates = cciptypes.PriceUpdates{}
		}
	}

	// check for final partial report.
	if numUpdates > 0 {
		report := buildOneReport(
			lggr,
			outcome.MerkleRootOutcome.OutcomeType,
			nil,
			priceUpdates,
		)
		reports = append(reports, report)
	}

	return reports
}
