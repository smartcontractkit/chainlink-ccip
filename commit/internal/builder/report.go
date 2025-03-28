package builder

import (
	"context"
	"fmt"
	"sort"

	"github.com/smartcontractkit/chainlink-ccip/commit/committypes"
	"github.com/smartcontractkit/chainlink-ccip/commit/merkleroot"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
	"github.com/smartcontractkit/chainlink-ccip/pluginconfig"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
)

type Report struct {
	Report     cciptypes.CommitPluginReport
	ReportInfo cciptypes.CommitReportInfo
}

// ReportBuilderFunc is used to inject different algorithms for building commit reports.
type ReportBuilderFunc func(
	ctx context.Context,
	lggr logger.Logger,
	outcome committypes.Outcome,
	config pluginconfig.CommitOffchainConfig,
) ([]Report, error)

// NewReportBuilder returns a ReportBuilderFunc based on the provided config.
func NewReportBuilder(RMNEnabled bool, MaxMerkleRootsPerReport, MaxPricesPerReport uint64) (ReportBuilderFunc, error) {
	// These options were added to allow for more flexibility around report building. For example Solana
	// only supports a single merkle root per report.

	if RMNEnabled {
		if MaxPricesPerReport > 0 || MaxMerkleRootsPerReport > 0 {
			return nil, fmt.Errorf("RMNEnabled is not supported with MaxPricesPerReport or MaxMerkleRootsPerReport set")
		}
		return buildStandardReport, nil
	}

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
// algorithms can reassemble reports by selecting which blessed and unblessed merkle
// roots to include in the report, and which price updates and rmn signatures to use.
func buildOneReport(
	ctx context.Context,
	lggr logger.Logger,
	merkleOutcomeType merkleroot.OutcomeType,
	blessedMerkleRoots []cciptypes.MerkleRootChain,
	unblessedMerkleRoots []cciptypes.MerkleRootChain,
	rmnSignatures []cciptypes.RMNECDSASignature,
	rmnRemoteFSign uint64,
	priceUpdates cciptypes.PriceUpdates,
) (Report, error) {
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

	// ReportEmpty and ReportInFlight means there's no roots to report.
	// However, there still may be price updates to report.
	if merkleOutcomeType == merkleroot.ReportEmpty || merkleOutcomeType == merkleroot.ReportInFlight {
		rep.BlessedMerkleRoots = []cciptypes.MerkleRootChain{}
		rep.UnblessedMerkleRoots = []cciptypes.MerkleRootChain{}
		rep.RMNSignatures = []cciptypes.RMNECDSASignature{}
	}

	if merkleOutcomeType == merkleroot.ReportGenerated {
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
		lggr.Errorw("buildOneReport: generated an empty report",
			"blessedMerkleRoots", blessedMerkleRoots,
			"unblessedMerkleRoots", unblessedMerkleRoots,
			"priceUpdates", priceUpdates,
		)

		return Report{}, nil
	}

	return Report{
		Report:     rep,
		ReportInfo: repInfo,
	}, nil
}

// buildStandardReport builds a one report with all the merkle roots and price updates.
func buildStandardReport(
	ctx context.Context,
	lggr logger.Logger,
	outcome committypes.Outcome,
	_ pluginconfig.CommitOffchainConfig,
) ([]Report, error) {
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
		outcome.MerkleRootOutcome.OutcomeType,
		blessedMerkleRoots,
		unblessedMerkleRoots,
		outcome.MerkleRootOutcome.RMNReportSignatures,
		outcome.MerkleRootOutcome.RMNRemoteCfg.FSign,
		priceUpdates,
	)
	if err != nil {
		return nil, fmt.Errorf("buildStandardReport err: %w", err)
	}
	// Do not return an empty report.
	if report.Report.IsEmpty() {
		return nil, nil
	}
	return []Report{report}, nil
}

// buildMultiplePriceReports builds many reports of with at most maxMerkleRootsPerReport roots.
func buildMultiplePriceAndMerkleRootReports(
	ctx context.Context,
	lggr logger.Logger,
	outcome committypes.Outcome,
	config pluginconfig.CommitOffchainConfig,
) ([]Report, error) {
	// 1. Build price reports.
	maxPrices := config.MaxPricesPerReport
	reports, err := buildMultiplePriceReports(ctx, lggr, outcome, maxPrices)
	if err != nil {
		return nil, fmt.Errorf("problem building price reports: %w", err)
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

	rootReports, err := rootReportBuilder(ctx, lggr, outcome, config)
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
	outcome committypes.Outcome,
	config pluginconfig.CommitOffchainConfig,
) ([]Report, error) {
	var reports []Report

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
				outcome.MerkleRootOutcome.OutcomeType,
				blessedMerkleRoots,
				unblessedMerkleRoots,
				nil, // no RMN for partial reports.
				0,   // no RMN for partial reports.
				priceUpdates,
			)
			if err != nil {
				return nil, fmt.Errorf("buildMultipleMerkleRootReports err: %w", err)
			}
			reports = append(reports, report)

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
			outcome.MerkleRootOutcome.OutcomeType,
			blessedMerkleRoots,
			unblessedMerkleRoots,
			nil, // no RMN for partial reports.
			0,   // no RMN for partial reports.
			priceUpdates,
		)
		if err != nil {
			return nil, fmt.Errorf("buildMultipleMerkleRootReports err: %w", err)
		}

		reports = append(reports, report)
	}

	return reports, nil
}

// buildMultiplePriceReports is a helper to split price data into multiple reports.
// Helper for buildMultiplePriceAndMerkleRootReports.
// Merkle root data is ignored.
func buildMultiplePriceReports(
	ctx context.Context,
	lggr logger.Logger,
	outcome committypes.Outcome,
	maxPricesPerReport uint64, // passed in directly to avoid implementing ReportBuilderFunc
) ([]Report, error) {
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
			report, err := buildOneReport(
				ctx,
				lggr,
				outcome.MerkleRootOutcome.OutcomeType,
				nil,
				nil,
				nil, // no RMN for partial reports.
				0,   // no RMN for partial reports.
				priceUpdates,
			)
			if err != nil {
				return nil, fmt.Errorf("buildMultiplePriceReports err: %w", err)
			}
			reports = append(reports, report)

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
			outcome.MerkleRootOutcome.OutcomeType,
			nil,
			nil,
			nil, // no RMN for partial reports.
			0,   // no RMN for partial reports.
			priceUpdates,
		)
		if err != nil {
			return nil, fmt.Errorf("buildMultiplePriceReports err: %w", err)
		}
		reports = append(reports, report)
	}

	return reports, nil
}
