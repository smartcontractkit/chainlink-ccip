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
	lggr               logger.Logger
	reports            *cache.Cache // key string -> CommitPluginReportWithMeta
	lastQueryTimestamp time.Time    // timestamp of the most recent query
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

// generateKey creates a unique key for a commit report
func generateKey(report ccipocr3.CommitPluginReportWithMeta) string {
	// We'll create a unique key for each report by combining relevant fields
	// For each blessed merkle root
	for _, mrc := range report.Report.BlessedMerkleRoots {
		return fmt.Sprintf("%s_%s", mrc.ChainSel.String(), mrc.MerkleRoot.String())
	}

	// If no blessed roots (unlikely but possible), check unblessed roots
	for _, mrc := range report.Report.UnblessedMerkleRoots {
		return fmt.Sprintf("%s_%s", mrc.ChainSel.String(), mrc.MerkleRoot.String())
	}

	// Fallback if no merkle roots (should never happen)
	return fmt.Sprintf("%d_%d", report.Timestamp.Unix(), report.BlockNum)
}

// GetCachedAndNewReports combines cached reports with newly fetched ones
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

	// Even if there were no new finalized reports, we want to avoid repeatedly querying the
	// same range. Advance the lastQueryTimestamp optimistically to queryFrom so that on the
	// next invocation we start from this point onward.
	if len(newFinalizedReports) == 0 {
		c.lastQueryTimestamp = queryFrom
	}

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
	if len(reports) == 0 {
		return
	}

	// Add reports to cache
	var mostRecentTimestamp time.Time
	for _, report := range reports {
		key := generateKey(report)
		c.reports.SetDefault(key, report)

		// Keep track of the most recent report timestamp
		if mostRecentTimestamp.Before(report.Timestamp) {
			mostRecentTimestamp = report.Timestamp
		}
	}

	// Update last query timestamp to the most recent report timestamp
	// Only update if we found a newer timestamp
	if !mostRecentTimestamp.IsZero() && mostRecentTimestamp.After(c.lastQueryTimestamp) {
		c.lastQueryTimestamp = mostRecentTimestamp
	} else {
		// Fallback: if no newer timestamp found, use the query timestamp
		c.lastQueryTimestamp = queryTimestamp
	}
}

// mergeCachedAndNewReports combines cached and new reports, removing duplicates
func (c *CommitReportsCache) mergeCachedAndNewReports(
	cachedReports []ccipocr3.CommitPluginReportWithMeta,
	newReports []ccipocr3.CommitPluginReportWithMeta,
	minTimestamp time.Time,
) []ccipocr3.CommitPluginReportWithMeta {
	// Create a map to deduplicate reports
	reportMap := make(map[string]ccipocr3.CommitPluginReportWithMeta)

	// Add cached reports
	for _, report := range cachedReports {
		if !report.Timestamp.Before(minTimestamp) {
			key := generateKey(report)
			reportMap[key] = report
		}
	}

	// Add new reports (will override cached ones if duplicates)
	for _, report := range newReports {
		key := generateKey(report)
		reportMap[key] = report
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
