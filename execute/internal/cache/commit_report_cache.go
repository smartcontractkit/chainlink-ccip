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

	"github.com/smartcontractkit/chainlink-ccip/pkg/reader"
	"github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

const (
	// DefaultMaxCommitReportsToFetch defines the default limit for fetching commit reports in one go.
	DefaultMaxCommitReportsToFetch = 1000
)

// CommitReportCache stores CommitPluginReportWithMeta objects.
// It's designed to handle potential LogPoller delays by using a lookback grace period.
type CommitReportCache interface {
	// RefreshCache fetches new reports using the configured reader and adds them to the cache.
	// It filters out reports without Merkle roots.
	// It returns an error if fetching reports fails.
	RefreshCache(ctx context.Context) error

	// GetReportsToQueryFromTimestamp determines the timestamp from which new reports should be queried.
	// This considers the messageVisibilityInterval, the latest report timestamp in the cache,
	// and a lookbackGracePeriod to account for LogPoller delays.
	GetReportsToQueryFromTimestamp() time.Time

	// GetCachedReports retrieves reports from the cache that have a timestamp greater than or equal to fromTimestamp.
	// Reports are sorted by their timestamp.
	GetCachedReports(fromTimestamp time.Time) []ccipocr3.CommitPluginReportWithMeta
}

type commitReportCache struct {
	lggr   logger.Logger
	cfg    CommitReportCacheConfig
	reader reader.CCIPReader

	cacheMu               sync.RWMutex
	reportsCache          *cache.Cache
	latestReportTimestamp time.Time
	timeProvider          TimeProvider
}

// CommitReportCacheConfig holds configuration for the CommitReportCache.
type CommitReportCacheConfig struct {
	MessageVisibilityInterval time.Duration // How long a report is considered "visible" or actionable.
	EvictionGracePeriod       time.Duration // Additional time before an item is evicted after its natural expiry
	CleanupInterval           time.Duration // How often the cache is scanned for expired items.
	LookbackGracePeriod       time.Duration // How far back to look from the latest known report to catch delayed logs.
}

// NewCommitReportCache creates a new CommitReportCache.
func NewCommitReportCache(
	lggr logger.Logger,
	cfg CommitReportCacheConfig,
	timeProvider TimeProvider,
	reader reader.CCIPReader,
) CommitReportCache {
	if cfg.LookbackGracePeriod == 0 {
		// Default lookback grace period if not set, e.g., messageVisibilityInterval / 8
		// This should be configured based on expected LogPoller delays.
		// For now, let's ensure it's at least a fraction of the cleanup interval or a fixed minimum.
		// A zero lookback might defeat the purpose of the cache for LogPoller delays.
		// Setting a sensible default or requiring it to be explicitly set is important.
		defaultLookback := cfg.MessageVisibilityInterval / 8
		if defaultLookback < 1*time.Minute && cfg.CleanupInterval > 0 {
			defaultLookback = cfg.CleanupInterval / 4
		}
		if defaultLookback < 1*time.Minute {
			// absolute minimum fallback
			defaultLookback = 1 * time.Minute
		}
		lggr.Warnw("LookbackGracePeriod is zero, setting a default.", "defaultLookback", defaultLookback)
		cfg.LookbackGracePeriod = defaultLookback
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
		lggr:                  lggr,
		cfg:                   cfg,
		reportsCache:          reportsCacheInstance,
		timeProvider:          timeProvider,
		reader:                reader,
		latestReportTimestamp: time.Time{}, // Initialize to zero, will be updated as reports are added
	}
}

// generateKey creates a unique key for a commit report
func generateKey(report ccipocr3.CommitPluginReportWithMeta) string {
	// We'll create a unique key for each report by combining relevant fields
	// For each blessed merkle root
	// NOTE: This keying strategy means if a report has multiple blessed/unblessed roots,
	// only the *first* one encountered will define its key. This might be problematic if reports
	// can have multiple roots that need to be treated as distinct entries or if the order is not guaranteed.
	// For now, implementing as requested.
	for _, mrc := range report.Report.BlessedMerkleRoots {
		return fmt.Sprintf("%s_%s", mrc.ChainSel.String(), mrc.MerkleRoot.String())
	}

	// If no blessed roots (unlikely but possible), check unblessed roots
	for _, mrc := range report.Report.UnblessedMerkleRoots {
		return fmt.Sprintf("%s_%s", mrc.ChainSel.String(), mrc.MerkleRoot.String())
	}

	// Fallback if no merkle roots (this should ideally not be reached if reports are pre-filtered
	// before attempting to add to a cache that expects roots, or if generateKey is only called for reports with roots)
	return fmt.Sprintf("%d_%d", report.Timestamp.Unix(), report.BlockNum)
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

	c.cacheMu.Lock()
	defer c.cacheMu.Unlock()

	addedCount := 0
	for _, report := range reports {
		// Filter: only store reports that have Merkle roots.
		if len(report.Report.BlessedMerkleRoots) == 0 && len(report.Report.UnblessedMerkleRoots) == 0 {
			c.lggr.Debugw("RefreshCache: skipping report with no Merkle roots",
				"timestamp", report.Timestamp,
				"blockNum", report.BlockNum)
			continue
		}

		key := generateKey(report)
		// Add to cache with default expiration (MessageVisibilityInterval + EvictionGracePeriod)
		c.reportsCache.SetDefault(key, report)
		addedCount++

		if report.Timestamp.After(c.latestReportTimestamp) {
			c.latestReportTimestamp = report.Timestamp.UTC()
		}
	}

	if addedCount > 0 {
		c.lggr.Infow("Refreshed CommitReportCache",
			"addedCount", addedCount,
			"fetchedCount", len(reports),
			"totalItemsInCache", c.reportsCache.ItemCount(),
			"latestReportTimestamp", c.latestReportTimestamp,
			"queryTimestampUsed", queryTs)
	}
	return nil
}

func (c *commitReportCache) GetReportsToQueryFromTimestamp() time.Time {
	c.cacheMu.RLock()
	defer c.cacheMu.RUnlock()

	now := c.timeProvider.Now()
	// Lower bound for query is dictated by the message visibility window.
	// We should not query for reports older than this, as they are no longer actionable for execution.
	visibilityWindowStart := now.Add(-c.cfg.MessageVisibilityInterval)

	// Default query start time is the beginning of the visibility window.
	queryFrom := visibilityWindowStart

	// If we have a latest report timestamp from the cache, we can potentially optimize.
	// We need to look back from this latest timestamp by the LookbackGracePeriod
	// to account for LogPoller delays.
	if !c.latestReportTimestamp.IsZero() {
		potentialQueryStart := c.latestReportTimestamp.Add(-c.cfg.LookbackGracePeriod)
		// We take the later of the visibilityWindowStart or (latestReportTimestamp - lookback).
		// This ensures we don't query too far back if latestReportTimestamp is very old,
		// but also ensures we apply the lookback if latestReportTimestamp is recent.
		if potentialQueryStart.After(queryFrom) {
			queryFrom = potentialQueryStart
		}
	}

	c.lggr.Debugw("Determined timestamp to query reports from",
		"now", now,
		"messageVisibilityInterval", c.cfg.MessageVisibilityInterval,
		"visibilityWindowStart", visibilityWindowStart,
		"latestReportTimestampInCache", c.latestReportTimestamp,
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
		// So, they should generally be recent enough.
		if report.Timestamp.After(fromTimestamp) || report.Timestamp.Equal(fromTimestamp) {
			result = append(result, report)
		}
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

// Ensure commitReportCache implements CommitReportCache.
var _ CommitReportCache = (*commitReportCache)(nil)
