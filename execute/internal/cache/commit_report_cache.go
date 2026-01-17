package cache

import (
	"context"
	"fmt"
	"sort"
	"sync"
	"time"

	"github.com/patrickmn/go-cache"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/types/query/primitives"

	"github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-ccip/pkg/reader"
)

const (
	// DefaultMaxCommitReportsToFetch defines the default limit for fetching commit reports in one go.
	DefaultMaxCommitReportsToFetch = 1000
	MinimumLookbackGracePeriod     = 1 * time.Minute  // Minimum sensible lookback (potential LogPoller delays)
	DefaultLookbackGracePeriod     = 30 * time.Minute // Default lookback grace period if not set
)

// TimeProvider interface to make time testable
type TimeProvider interface {
	Now() time.Time
}

// RealTimeProvider provides actual current time
type RealTimeProvider struct{}

func (r *RealTimeProvider) Now() time.Time {
	return time.Now().UTC()
}

// CommitReportCache stores CommitPluginReportWithMeta objects.
// It's designed to handle potential LogPoller delays by using a lookback grace period.
type CommitReportCache interface {
	// RefreshCache fetches new reports using the configured reader and adds them to the cache.
	// It filters out reports without Merkle roots before storing them.
	// It returns an error if fetching reports fails.
	RefreshCache(ctx context.Context) error

	// GetReportsToQueryFromTimestamp determines the timestamp from which new reports should be queried.
	// This considers the messageVisibilityInterval, the timestamp of the latest report seen by the
	// reader during the last refresh (latestFinalizedReportTimestamp), and a lookbackGracePeriod
	// to account for LogPoller delays or other potential inconsistencies.
	GetReportsToQueryFromTimestamp() time.Time

	// GetCachedReports retrieves reports from the cache that have a timestamp greater than or equal to fromTimestamp.
	// Reports are sorted by their timestamp.
	GetCachedReports(fromTimestamp time.Time) []ccipocr3.CommitPluginReportWithMeta
}

type commitReportCache struct {
	lggr   logger.Logger
	cfg    CommitReportCacheConfig
	reader reader.CCIPReader

	cacheMu                        sync.RWMutex
	reportsCache                   *cache.Cache
	latestFinalizedReportTimestamp time.Time // Considering also empty merkle root reports
	timeProvider                   TimeProvider
}

// CommitReportCacheConfig holds configuration for the CommitReportCache.
type CommitReportCacheConfig struct {
	MessageVisibilityInterval time.Duration // How long a report is considered "visible" or actionable.
	EvictionGracePeriod       time.Duration // Additional time before an item is evicted after its natural expiry
	CleanupInterval           time.Duration // How often the cache is scanned for expired items.
	LookbackGracePeriod       time.Duration // How far back to look from the latest known report to catch delayed logs.
}

// NewCommitReportCache creates a new CommitReportCache.
// Should only be used for oracles supporting the remote/destination chain.
func NewCommitReportCache(
	lggr logger.Logger,
	cfg CommitReportCacheConfig,
	timeProvider TimeProvider,
	reader reader.CCIPReader,
) CommitReportCache {
	if cfg.LookbackGracePeriod == 0 {
		// Use a fixed default of 30 minutes instead of calculating based on messageVisibilityInterval
		lggr.Warnf("LookbackGracePeriod is not set, using default of %v", DefaultLookbackGracePeriod)
		cfg.LookbackGracePeriod = DefaultLookbackGracePeriod
	}

	lggr.Infow(
		"Creating CommitReportCache",
		"messageVisibilityInterval", cfg.MessageVisibilityInterval,
		"evictionGracePeriod", cfg.EvictionGracePeriod,
		"cleanupInterval", cfg.CleanupInterval,
		"lookbackGracePeriod", cfg.LookbackGracePeriod,
	)

	reportsCacheInstance := cache.New(cfg.MessageVisibilityInterval+cfg.EvictionGracePeriod, cfg.CleanupInterval)

	return &commitReportCache{
		lggr:                           lggr,
		cfg:                            cfg,
		reportsCache:                   reportsCacheInstance,
		timeProvider:                   timeProvider,
		reader:                         reader,
		latestFinalizedReportTimestamp: time.Time{}, // Initialize to zero, will be updated as reports are added
	}
}

// generateKey creates a unique key for a commit report.
// The key is based on the ChainSelector and MerkleRoot of the first Merkle root found,
// prioritizing BlessedMerkleRoots over UnblessedMerkleRoots.
// It is sufficient to use only the first Merkle root (e.g., BlessedMerkleRoots[0] or UnblessedMerkleRoots[0])
// because this cache stores finalized commit reports. For finalized reports, the set of Merkle roots
// and their order within the report are stable. We do not expect re-organizations to cause these.
// Returns an error if the report has no Merkle roots.
func generateKey(report ccipocr3.CommitPluginReportWithMeta) (string, error) {
	// Check for blessed roots first
	if len(report.Report.BlessedMerkleRoots) > 0 {
		root := report.Report.BlessedMerkleRoots[0]
		return fmt.Sprintf("%d_%s", root.ChainSel, root.MerkleRoot), nil
	}

	// Then check for unblessed roots
	if len(report.Report.UnblessedMerkleRoots) > 0 {
		root := report.Report.UnblessedMerkleRoots[0]
		return fmt.Sprintf("%d_%s", root.ChainSel, root.MerkleRoot), nil
	}

	// Return error if no roots
	return "", fmt.Errorf("report has no Merkle roots")
}

func (c *commitReportCache) RefreshCache(ctx context.Context) error {
	queryTs := c.GetReportsToQueryFromTimestamp()
	c.lggr.Debugw("RefreshCache: attempting to fetch reports", "queryTimestamp", queryTs)

	reports, err := c.reader.CommitReportsGTETimestamp(
		ctx,
		queryTs,
		primitives.Finalized,
		DefaultMaxCommitReportsToFetch,
	)
	if err != nil {
		c.lggr.Errorw("RefreshCache: failed to get reports from reader", "err", err, "queryTimestamp", queryTs)
		return fmt.Errorf("failed to get reports from reader: %w", err)
	}

	c.lggr.Debugw("RefreshCache: received reports from reader", "count", len(reports), "queryTimestamp", queryTs)

	addedCount := 0

	// Update latestFinalizedReportTimestamp based on all fetched reports before filtering
	if len(reports) > 0 {
		maxTs := reports[0].Timestamp
		for _, r := range reports[1:] {
			if r.Timestamp.After(maxTs) {
				maxTs = r.Timestamp
			}
		}

		c.cacheMu.Lock()
		if maxTs.After(c.latestFinalizedReportTimestamp) {
			c.latestFinalizedReportTimestamp = maxTs.UTC()
			c.lggr.Debugw("RefreshCache: updated latestFinalizedReportTimestamp from fetched batch",
				"newLatest", c.latestFinalizedReportTimestamp,
				"previousLatest", queryTs)
		}
		c.cacheMu.Unlock()
	}

	// The for loop structure itself is maintained as per the selection boundaries.
	// Filtering and key generation are done outside the lock.
	for _, report := range reports {
		// Filter: only store reports that have Merkle roots.
		if len(report.Report.BlessedMerkleRoots) == 0 && len(report.Report.UnblessedMerkleRoots) == 0 {
			c.lggr.Debugw("RefreshCache: skipping report with no Merkle roots",
				"timestamp", report.Timestamp,
				"blockNum", report.BlockNum)
			continue
		}

		key, err := generateKey(report)
		if err != nil {
			c.lggr.Errorw("RefreshCache: failed to generate key for report",
				"err", err,
				"timestamp", report.Timestamp,
				"blockNum", report.BlockNum)
			continue
		}
		c.cacheMu.Lock()
		// Add to cache with default expiration (MessageVisibilityInterval + EvictionGracePeriod)
		c.reportsCache.SetDefault(key, report)
		c.cacheMu.Unlock()

		addedCount++ // Increment for each report successfully processed and added.
	}

	if addedCount > 0 {
		c.lggr.Infow("Refreshed CommitReportCache",
			"addedCount", addedCount,
			"fetchedCount", len(reports),
			"totalItemsInCache", c.reportsCache.ItemCount(),
			"latestReportTimestamp", c.latestFinalizedReportTimestamp,
			"queryTimestampUsed", queryTs)
	}
	return nil
}

func (c *commitReportCache) GetReportsToQueryFromTimestamp() time.Time {
	c.cacheMu.RLock()
	defer c.cacheMu.RUnlock()

	now := c.timeProvider.Now()
	// Default query start time is the beginning of the visibility window (messageVisibilityInterval ago).
	// This also serves as the lower bound for the query (reports older non actionable).
	queryFrom := now.Add(-c.cfg.MessageVisibilityInterval)

	// visibilityWindowStart is defined here to capture the initial queryFrom value,
	// as this specific variable name is used in logging outside this selection.
	visibilityWindowStart := queryFrom

	// If we have a latest report timestamp from the cache, we can potentially optimize.
	// We need to look back from this latest timestamp by the LookbackGracePeriod
	// to account for LogPoller delays.
	if !c.latestFinalizedReportTimestamp.IsZero() {
		potentialQueryStart := c.latestFinalizedReportTimestamp.Add(-c.cfg.LookbackGracePeriod)
		// We take the later of the initial queryFrom value (i.e., visibilityWindowStart)
		// or (latestReportTimestamp - lookbackGracePeriod).
		// This ensures we don't query too far back if latestFinalizedReportTimestamp is very old,
		// but also ensures we apply the lookback if latestFinalizedReportTimestamp is recent.
		if potentialQueryStart.After(queryFrom) { // queryFrom here is effectively visibilityWindowStart
			queryFrom = potentialQueryStart
		}
	}

	c.lggr.Debugw("Determined timestamp to query reports from",
		"now", now,
		"messageVisibilityInterval", c.cfg.MessageVisibilityInterval,
		"visibilityWindowStart", visibilityWindowStart,
		"latestReportTimestampInCache", c.latestFinalizedReportTimestamp,
		"lookbackGracePeriod", c.cfg.LookbackGracePeriod,
		"calculatedQueryFrom", queryFrom,
	)
	return queryFrom
}

func (c *commitReportCache) GetCachedReports(fromTimestamp time.Time) []ccipocr3.CommitPluginReportWithMeta {
	c.cacheMu.RLock()
	defer c.cacheMu.RUnlock()

	var result []ccipocr3.CommitPluginReportWithMeta
	items := c.reportsCache.Items() // Gets a copy of all items: map[string]cache.Item

	for key, item := range items {
		if item.Expired() {
			continue
		}
		report, ok := item.Object.(ccipocr3.CommitPluginReportWithMeta)
		if !ok {
			c.lggr.Errorw("Failed to assert type for cached report item", "key", key) // Use map key here
			continue
		}

		// Filter by fromTimestamp. We want reports >= fromTimestamp.
		// The cache stores items that are within (MessageVisibilityInterval + EvictionGracePeriod)
		if report.Timestamp.Before(fromTimestamp) {
			continue
		}
		result = append(result, report)
	}

	// Sort reports by timestamp (ascending)
	sort.Slice(result, func(i, j int) bool {
		if result[i].Timestamp.Equal(result[j].Timestamp) {
			// If timestamps are equal, sort by block number as a secondary criterion
			return result[i].BlockNum < result[j].BlockNum
		}
		return result[i].Timestamp.Before(result[j].Timestamp)
	})

	c.lggr.Debugw("Retrieved cached reports",
		"fromTimestamp", fromTimestamp,
		"count", len(result),
		"totalItemsInCache", len(items))
	return result
}

// DeduplicateReports deduplicates reports by their Merkle root.
// It uses the generateKey function to identify unique reports.
// Reports for which a key cannot be generated (e.g., no Merkle roots)
// are logged and skipped.
func DeduplicateReports(
	lggr logger.Logger,
	reports []ccipocr3.CommitPluginReportWithMeta,
) []ccipocr3.CommitPluginReportWithMeta {
	seen := make(map[string]bool)
	deduplicated := make([]ccipocr3.CommitPluginReportWithMeta, 0, len(reports))

	for _, report := range reports {
		if len(report.Report.BlessedMerkleRoots) == 0 && len(report.Report.UnblessedMerkleRoots) == 0 {
			lggr.Debugw("DeduplicateReports: skipping report with no Merkle roots",
				"timestamp", report.Timestamp,
				"blockNum", report.BlockNum)
			continue
		}

		key, err := generateKey(report) // generateKey is already in the package
		if err != nil {
			lggr.Errorw("DeduplicateReports: Failed to generate key for report, skipping",
				"err", err,
				"timestamp", report.Timestamp,
				"blockNum", report.BlockNum)
			continue
		}

		if !seen[key] {
			deduplicated = append(deduplicated, report)
			seen[key] = true
		}
	}
	return deduplicated
}

// Ensure commitReportCache implements CommitReportCache.
var _ CommitReportCache = (*commitReportCache)(nil)
