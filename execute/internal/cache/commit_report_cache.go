package cache

import (
	"context"
	"fmt"
	"sort"
	"time"

	"github.com/patrickmn/go-cache"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/types/query/primitives"

	readerpkg "github.com/smartcontractkit/chainlink-ccip/pkg/reader"
	"github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

// CommitReportsCache optimizes calls to CommitReportsGTETimestamp by caching reports
// and only fetching new reports since the last successful query.
type CommitReportsCache struct {
	lggr    logger.Logger
	reports *cache.Cache // key string -> CommitPluginReportWithMeta
	// lastQueryTimestamp is the starting point for the next finalized reports query to the CCIP reader.
	// It is updated to the timestamp of the newest cached finalized report (with roots) or the reader query time,
	// ensuring the query window always moves forward to optimize fetching and avoid redundant calls.
	lastQueryTimestamp time.Time
	ccipReader         readerpkg.CCIPReader
	queryLimit         int
	returnLimit        int
}

// NewCommitReportsCache creates a new commit reports cache.
// Cached items live for messageVisibilityInterval + EvictionGracePeriod, mirroring
// the lifetime used in commitRootsCache for executed roots.
func NewCommitReportsCache(
	lggr logger.Logger,
	messageVisibilityInterval time.Duration,
	ccipReader readerpkg.CCIPReader,
	queryLimit int,
	returnLimit int,
) *CommitReportsCache {
	ttl := messageVisibilityInterval + EvictionGracePeriod
	return &CommitReportsCache{
		lggr:        lggr,
		reports:     cache.New(ttl, CleanupInterval),
		ccipReader:  ccipReader,
		queryLimit:  queryLimit,
		returnLimit: returnLimit,
	}
}

// generateKey creates a unique key for a commit report.
// It expects reports to have Merkle roots, as rootless reports should be filtered out by callers.
// If a report without roots is passed (which indicates a logic error in filtering), it returns an error.
func generateKey(report ccipocr3.CommitPluginReportWithMeta) (string, error) {
	// We'll create a unique key for each report by combining relevant fields
	// For each blessed merkle root
	for _, mrc := range report.Report.BlessedMerkleRoots {
		return fmt.Sprintf("%d_%x", mrc.ChainSel, mrc.MerkleRoot), nil
	}

	// If no blessed roots (solana or other non-rmn enabled chains), check unblessed roots
	for _, mrc := range report.Report.UnblessedMerkleRoots {
		return fmt.Sprintf("%d_%x", mrc.ChainSel, mrc.MerkleRoot), nil
	}

	// Fallback: This should ideally not be reached if callers correctly filter reports using HasNoRoots().
	// If reached, it indicates a logic error where a rootless report was passed for key generation.
	return "", fmt.Errorf(
		"generateKey: report at timestamp %d, block %d has no roots but was not filtered before key generation",
		report.Timestamp.Unix(), report.BlockNum)
}

// GetCachedAndNewReports fetches and combines cached and new commit reports.
// It aims to return up to `c.returnLimit` of the newest reports >= `fetchFrom`.
// Key steps:
//  1. Determine `queryFrom` timestamp: Uses `fetchFrom` for the first call, or `c.lastQueryTimestamp`
//     for subsequent calls to fetch only newer reports.
//  2. Fetch and cache finalized reports: Retrieves finalized reports since `queryFrom`,
//     adds them to the cache, and updates `c.lastQueryTimestamp`.
//  3. Early exit: If cached reports (including newly finalized ones) are sufficient to meet
//     `c.returnLimit`, returns them directly, sorted by newest first.
//  4. Fetch unconfirmed reports: If early exit isn't taken, fetches unconfirmed reports since `queryFrom`.
//  5. Merge, sort, and limit: Combines initial cached reports (from before finalized fetch) with
//     new unconfirmed reports, de-duplicates, filters by `fetchFrom`, sorts by newest, and limits to `c.returnLimit`.
func (c *CommitReportsCache) GetCachedAndNewReports(
	ctx context.Context,
	fetchFrom time.Time,
) ([]ccipocr3.CommitPluginReportWithMeta, error) {
	// Start with cached reports that are newer than or equal to fetchFrom
	// These are used if we don't have enough after adding finalized reports and need to merge with unconfirmed.
	initialCachedReports := c.getCachedReports(fetchFrom)

	// Determine if we need to fetch new reports
	var queryFrom time.Time
	// No lock required for reports (go-cache is thread-safe), but we still protect lastQueryTimestamp.
	if c.lastQueryTimestamp.IsZero() {
		// First query, use the original fetchFrom
		queryFrom = fetchFrom
	} else {
		// For subsequent queries, start from the last query timestamp
		// This is the key optimization: we only query for new reports since our last query
		queryFrom = c.lastQueryTimestamp
	}

	// Fetch finalized reports since last query and update the cache with them.
	newFinalizedReports, err := c.ccipReader.CommitReportsGTETimestamp(ctx, queryFrom, primitives.Finalized, c.queryLimit)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch finalized commit reports: %w", err)
	}
	c.addReportsToCache(newFinalizedReports, queryFrom)

	// If the reader returned no new finalized reports for the query window [queryFrom, limit):
	// - If queryFrom itself was more recent than c.lastQueryTimestamp, then c.lastQueryTimestamp
	//   should advance to queryFrom to avoid re-querying an empty range from an older LQT.
	// - If queryFrom was not more recent (i.e., c.lastQueryTimestamp was already >= queryFrom),
	//   c.lastQueryTimestamp should not change (it shouldn't regress).
	if len(newFinalizedReports) == 0 {
		if queryFrom.After(c.lastQueryTimestamp) {
			c.lastQueryTimestamp = queryFrom
		}
	}
	// If newFinalizedReports was not empty, addReportsToCache has already handled updating c.lastQueryTimestamp.

	// Optimization: Check if the cache (after adding finalized reports) now has enough reports.
	currentAllCachedReports := c.getCachedReports(fetchFrom) // Get all cached reports >= fetchFrom
	sort.Slice(currentAllCachedReports, func(i, j int) bool {
		return currentAllCachedReports[i].Timestamp.After(currentAllCachedReports[j].Timestamp)
	})

	if len(currentAllCachedReports) >= c.returnLimit {
		c.lggr.Debugw("Returning early with sufficient reports from cache after finalized update",
			"count", len(currentAllCachedReports),
			"returnLimit", c.returnLimit,
			"fetchFrom", fetchFrom)
		if len(currentAllCachedReports) > c.returnLimit { // Ensure we don't exceed the returnLimit
			return currentAllCachedReports[:c.returnLimit], nil
		}
		return currentAllCachedReports, nil
	}

	// 2. Fetch unconfirmed (finalized + unfinalized) reports for returning to the caller.
	newUnconfirmedReports, err := c.ccipReader.CommitReportsGTETimestamp(
		ctx, queryFrom, primitives.Unconfirmed, c.queryLimit,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch unconfirmed commit reports: %w", err)
	}

	c.lggr.Debugw("Fetched new commit reports",
		"finalized", len(newFinalizedReports),
		"unconfirmed", len(newUnconfirmedReports),
		"queryFrom", queryFrom)

	// Merge cached finalized reports with fresh unconfirmed ones (which include finalized+unfinalized).
	allReports := c.mergeCachedAndNewReports(initialCachedReports, newUnconfirmedReports, fetchFrom)

	return allReports, nil
}

// getCachedReports returns cached reports newer than or equal to the given timestamp
func (c *CommitReportsCache) getCachedReports(minTimestamp time.Time) []ccipocr3.CommitPluginReportWithMeta {
	var reports []ccipocr3.CommitPluginReportWithMeta
	for _, item := range c.reports.Items() {
		rpt := item.Object.(ccipocr3.CommitPluginReportWithMeta)
		if rpt.Timestamp.Equal(minTimestamp) || rpt.Timestamp.After(minTimestamp) {
			reports = append(reports, rpt)
		}
	}
	return reports
}

// addReportsToCache adds reports to the cache and updates last query timestamp
func (c *CommitReportsCache) addReportsToCache(
	reports []ccipocr3.CommitPluginReportWithMeta,
	queryTimestamp time.Time) {
	// Add reports to cache
	var mostRecentCachedTimestamp time.Time
	addedToCacheCount := 0
	for _, report := range reports {
		if report.Report.HasNoRoots() {
			c.lggr.Debugw(
				"Skipping report with no Merkle roots in addReportsToCache",
				"timestamp", report.Timestamp,
				"blockNum", report.BlockNum)
			continue
		}
		key, err := generateKey(report)
		if err != nil {
			c.lggr.Errorw("Failed to generate key for report in addReportsToCache",
				"error", err,
				"timestamp", report.Timestamp,
				"blockNum", report.BlockNum)
			continue // Skip this report
		}
		c.reports.SetDefault(key, report)
		addedToCacheCount++

		// Keep track of the most recent report timestamp that was actually cached
		if mostRecentCachedTimestamp.Before(report.Timestamp) {
			mostRecentCachedTimestamp = report.Timestamp
		}
	}

	if addedToCacheCount > 0 {
		c.lggr.Debugw(
			"Added reports to cache",
			"count", addedToCacheCount,
			"mostRecentCachedTimestamp", mostRecentCachedTimestamp)
	}

	// Update last query timestamp.
	// If we cached any new reports with timestamps newer than the current lastQueryTimestamp, update to that.
	// Otherwise, update to the queryTimestamp to ensure we advance the window even if reports had no roots or were older.
	if !mostRecentCachedTimestamp.IsZero() && mostRecentCachedTimestamp.After(c.lastQueryTimestamp) {
		c.lastQueryTimestamp = mostRecentCachedTimestamp
	} else if len(reports) > 0 {
		// Only proceed if there were reports in this batch
		// This block handles cases where:
		// - mostRecentCachedTimestamp was zero (no rooted reports cached from this batch)
		// - OR mostRecentCachedTimestamp was not after c.lastQueryTimestamp (cached reports were not newer)
		// We need to ensure LQT advances if possible, especially past a batch of rootless reports.
		maxTimestampInBatch := time.Time{} // Zero time
		for _, r := range reports {
			if r.Timestamp.After(maxTimestampInBatch) {
				maxTimestampInBatch = r.Timestamp
			}
		}

		// If the latest timestamp from this batch is newer than current LQT, advance LQT.
		// This handles advancing LQT past a batch of (potentially all) rootless reports.
		if !maxTimestampInBatch.IsZero() && maxTimestampInBatch.After(c.lastQueryTimestamp) {
			c.lastQueryTimestamp = maxTimestampInBatch
		} else if queryTimestamp.After(c.lastQueryTimestamp) {
			// Fallback: If batch didn't provide a newer timestamp (e.g., all reports in batch were older than or same as LQT),
			// but the query itself started from a more recent point than LQT (e.g. initial fetchFrom was newer).
			c.lastQueryTimestamp = queryTimestamp
		}
		// If none of these conditions are met, LQT remains unchanged by this function.
		// This is appropriate if the current batch of reports and the queryTimestamp were all
		// not newer than the existing c.lastQueryTimestamp.
	}
	// If len(reports) == 0 (i.e., newFinalizedReports was empty), LQT is not changed by this function.
	// That case is handled by the caller (GetCachedAndNewReports).
}

// mergeCachedAndNewReports combines cached and new reports, removing duplicates
func (c *CommitReportsCache) mergeCachedAndNewReports(
	cachedReports []ccipocr3.CommitPluginReportWithMeta,
	newReports []ccipocr3.CommitPluginReportWithMeta,
	minTimestamp time.Time,
) []ccipocr3.CommitPluginReportWithMeta {
	// Create a map to deduplicate reports
	reportMap := make(map[string]ccipocr3.CommitPluginReportWithMeta)

	// Add cached reports (already filtered for roots and >= minTimestamp from getCachedReports)
	for _, report := range cachedReports {
		// Double check minTimestamp, though getCachedReports should handle this.
		// No need to check HasNoRoots, as they wouldn't be in cachedReports if they didn't have them.
		if !report.Timestamp.Before(minTimestamp) {
			key, err := generateKey(report)
			if err != nil {
				c.lggr.Errorw("Failed to generate key for cachedReport in mergeCachedAndNewReports",
					"error", err,
					"timestamp", report.Timestamp,
					"blockNum", report.BlockNum)
				continue // Skip this report
			}
			reportMap[key] = report
		}
	}

	// Add new reports (will override cached ones if duplicates)
	// Filter for roots here as newReports come directly from the reader.
	for _, report := range newReports {
		if report.Report.HasNoRoots() {
			c.lggr.Debugw("Skipping new report with no Merkle roots in mergeCachedAndNewReports",
				"timestamp", report.Timestamp, "blockNum", report.BlockNum)
			continue
		}
		// Only add if newer than minTimestamp. Note: the primary time filtering happens
		// via queryFrom and lastQueryTimestamp for fetching, and getCachedReports for initial cached set.
		// This minTimestamp check here is mostly for reports in newReports that might be older
		// than an explicit fetchFrom provided to GetCachedAndNewReports, though typically queryFrom
		// would align or be newer than fetchFrom.
		if !report.Timestamp.Before(minTimestamp) {
			key, err := generateKey(report)
			if err != nil {
				c.lggr.Errorw("Failed to generate key for newReport in mergeCachedAndNewReports",
					"error", err,
					"timestamp", report.Timestamp,
					"blockNum", report.BlockNum)
				continue // Skip this report
			}
			reportMap[key] = report
		}
	}

	// Convert map to slice
	allReports := make([]ccipocr3.CommitPluginReportWithMeta, 0, len(reportMap))
	for _, report := range reportMap {
		allReports = append(allReports, report)
	}

	// Sort reports by timestamp (newest first)
	sort.Slice(allReports, func(i, j int) bool {
		return allReports[i].Timestamp.After(allReports[j].Timestamp)
	})

	// Limit the number of reports
	if len(allReports) > c.returnLimit {
		allReports = allReports[:c.returnLimit]
	}

	return allReports
}
