package cache

import (
	"context"
	"fmt"
	"sort"
	"sync"
	"time"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"

	readerpkg "github.com/smartcontractkit/chainlink-ccip/pkg/reader"
	plugintypes2 "github.com/smartcontractkit/chainlink-ccip/plugintypes"
)

// CommitReportsCache optimizes calls to CommitReportsGTETimestamp by caching reports
// and only fetching new reports since the last successful query.
type CommitReportsCache struct {
	lggr               logger.Logger
	reportsMu          sync.RWMutex
	reports            map[string]plugintypes2.CommitPluginReportWithMeta // key is source_chain + merkle_root
	lastQueryTimestamp time.Time                                          // timestamp of the most recent query
}

// NewCommitReportsCache creates a new commit reports cache
func NewCommitReportsCache(lggr logger.Logger) *CommitReportsCache {
	return &CommitReportsCache{
		lggr:    lggr,
		reports: make(map[string]plugintypes2.CommitPluginReportWithMeta),
	}
}

// generateKey creates a unique key for a commit report
func generateKey(report plugintypes2.CommitPluginReportWithMeta) string {
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
	ccipReader readerpkg.CCIPReader,
	fetchFrom time.Time,
	limit int,
) ([]plugintypes2.CommitPluginReportWithMeta, error) {
	// Start with cached reports that are newer than or equal to fetchFrom
	cachedReports := c.getCachedReports(fetchFrom)

	// Determine if we need to fetch new reports
	var queryFrom time.Time
	c.reportsMu.RLock()
	if c.lastQueryTimestamp.IsZero() {
		// First query, use the original fetchFrom
		queryFrom = fetchFrom
	} else {
		// For subsequent queries, start from the last query timestamp
		// This is the key optimization: we only query for new reports since our last query
		queryFrom = c.lastQueryTimestamp
	}
	c.reportsMu.RUnlock()

	// Fetch new reports since the last query
	c.lggr.Debugw("Fetching new commit reports",
		"originalFetchFrom", fetchFrom,
		"optimizedQueryFrom", queryFrom,
		"cachedReportsCount", len(cachedReports))

	newReports, err := ccipReader.CommitReportsGTETimestamp(ctx, queryFrom, limit)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch commit reports: %w", err)
	}

	c.lggr.Debugw("Fetched new commit reports",
		"count", len(newReports),
		"queryFrom", queryFrom)

	// Update the cache with new reports
	c.addReportsToCache(newReports, queryFrom)

	// Merge cached and new reports, ensuring we don't exceed the limit
	allReports := c.mergeCachedAndNewReports(cachedReports, newReports, fetchFrom, limit)

	return allReports, nil
}

// getCachedReports returns cached reports newer than or equal to the given timestamp
func (c *CommitReportsCache) getCachedReports(minTimestamp time.Time) []plugintypes2.CommitPluginReportWithMeta {
	c.reportsMu.RLock()
	defer c.reportsMu.RUnlock()

	reports := make([]plugintypes2.CommitPluginReportWithMeta, 0, len(c.reports))
	for _, report := range c.reports {
		if report.Timestamp.Equal(minTimestamp) || report.Timestamp.After(minTimestamp) {
			reports = append(reports, report)
		}
	}

	return reports
}

// addReportsToCache adds reports to the cache and updates last query timestamp
func (c *CommitReportsCache) addReportsToCache(
	reports []plugintypes2.CommitPluginReportWithMeta,
	queryTimestamp time.Time) {
	if len(reports) == 0 {
		return
	}

	c.reportsMu.Lock()
	defer c.reportsMu.Unlock()

	// Add reports to cache
	var mostRecentTimestamp time.Time
	for _, report := range reports {
		key := generateKey(report)
		c.reports[key] = report

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
	cachedReports []plugintypes2.CommitPluginReportWithMeta,
	newReports []plugintypes2.CommitPluginReportWithMeta,
	minTimestamp time.Time,
	limit int,
) []plugintypes2.CommitPluginReportWithMeta {
	// Create a map to deduplicate reports
	reportMap := make(map[string]plugintypes2.CommitPluginReportWithMeta)

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
	allReports := make([]plugintypes2.CommitPluginReportWithMeta, 0, len(reportMap))
	for _, report := range reportMap {
		allReports = append(allReports, report)
	}

	// Sort reports by timestamp (newest first)
	sort.Slice(allReports, func(i, j int) bool {
		return allReports[i].Timestamp.After(allReports[j].Timestamp)
	})

	// Limit the number of reports
	if len(allReports) > limit {
		allReports = allReports[:limit]
	}

	return allReports
}

// RemoveReport removes a report from the cache
func (c *CommitReportsCache) RemoveReport(report plugintypes2.CommitPluginReportWithMeta) {
	c.reportsMu.Lock()
	defer c.reportsMu.Unlock()

	key := generateKey(report)
	delete(c.reports, key)
}

// Clear empties the cache
func (c *CommitReportsCache) Clear() {
	c.reportsMu.Lock()
	defer c.reportsMu.Unlock()

	c.reports = make(map[string]plugintypes2.CommitPluginReportWithMeta)
	c.lastQueryTimestamp = time.Time{}
}
