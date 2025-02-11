package reader

import (
	"context"
	"fmt"
	"sync"
	"time"

	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
)

// ConfigPoller defines the interface for caching chain configuration data
type ConfigPoller interface {
	// GetChainConfig retrieves the cached configuration for a chain
	GetChainConfig(ctx context.Context, chainSel cciptypes.ChainSelector) (ChainConfigSnapshot, error)
	// RefreshChainConfig forces a refresh of the chain configuration
	RefreshChainConfig(ctx context.Context, chainSel cciptypes.ChainSelector) (ChainConfigSnapshot, error)
}

// configPoller handles caching of chain configuration data for multiple chains.
// It is used by the ccipChainReader to store and retrieve configuration data,
// avoiding unnecessary contract calls and improving performance.
// configPoller handles caching of chain configuration data for multiple chains
type configPoller struct {
	sync.RWMutex
	chainCaches   map[cciptypes.ChainSelector]*chainCache
	refreshPeriod time.Duration
	reader        *ccipChainReader // Reference to the reader for fetching configs
	lggr          logger.Logger
}

// chainCache represents the cache for a single chain.
// It stores the configuration data for a specific chain and manages
// the last refresh time to determine when the data needs to be updated.
type chainCache struct {
	sync.RWMutex
	data        ChainConfigSnapshot
	lastRefresh time.Time
}

// newConfigPoller creates a new config cache instance
func newConfigPoller(
	lggr logger.Logger,
	reader *ccipChainReader,
	refreshPeriod time.Duration,
) *configPoller {
	return &configPoller{
		chainCaches:   make(map[cciptypes.ChainSelector]*chainCache),
		refreshPeriod: refreshPeriod,
		reader:        reader,
		lggr:          lggr,
	}
}

// getOrCreateChainCache safely retrieves or creates a cache for a specific chain
func (c *configPoller) getOrCreateChainCache(chainSel cciptypes.ChainSelector) *chainCache {
	c.Lock()
	defer c.Unlock()

	if cache, exists := c.chainCaches[chainSel]; exists {
		return cache
	}

	// verify we have the reader for this chain
	if _, exists := c.reader.contractReaders[chainSel]; !exists {
		c.lggr.Errorw("No contract reader for chain", "chain", chainSel)
		return nil
	}

	cache := &chainCache{}
	c.chainCaches[chainSel] = cache
	return cache
}

// GetChainConfig retrieves the cached configuration for a chain
func (c *configPoller) GetChainConfig(
	ctx context.Context,
	chainSel cciptypes.ChainSelector,
) (ChainConfigSnapshot, error) {
	// Check if we have a reader for this chain
	reader, exists := c.reader.contractReaders[chainSel]
	if !exists || reader == nil {
		c.lggr.Errorw("No contract reader for chain", "chain", chainSel)
		return ChainConfigSnapshot{}, fmt.Errorf("no contract reader for chain %d", chainSel)
	}

	chainCache := c.getOrCreateChainCache(chainSel)

	chainCache.RLock()
	timeSinceLastRefresh := time.Since(chainCache.lastRefresh)
	if timeSinceLastRefresh < c.refreshPeriod {
		defer chainCache.RUnlock()
		c.lggr.Debugw("Cache hit",
			"chain", chainSel,
			"timeSinceLastRefresh", timeSinceLastRefresh,
			"refreshPeriod", c.refreshPeriod)
		return chainCache.data, nil
	}
	chainCache.RUnlock()

	return c.RefreshChainConfig(ctx, chainSel)
}

// RefreshChainConfig forces a refresh of the chain configuration
func (c *configPoller) RefreshChainConfig(
	ctx context.Context,
	chainSel cciptypes.ChainSelector,
) (ChainConfigSnapshot, error) {
	chainCache := c.getOrCreateChainCache(chainSel)

	chainCache.Lock()
	defer chainCache.Unlock()

	// Double check if another goroutine has already refreshed
	timeSinceLastRefresh := time.Since(chainCache.lastRefresh)
	if timeSinceLastRefresh < c.refreshPeriod {
		c.lggr.Debugw("Cache was refreshed by another goroutine",
			"chain", chainSel,
			"timeSinceLastRefresh", timeSinceLastRefresh)
		return chainCache.data, nil
	}

	startTime := time.Now()
	newData, err := c.fetchChainConfig(ctx, chainSel)
	fetchConfigLatency := time.Since(startTime)

	if err != nil {
		if !chainCache.lastRefresh.IsZero() {
			c.lggr.Warnw("Failed to refresh cache, using old data",
				"chain", chainSel,
				"error", err,
				"lastRefresh", chainCache.lastRefresh,
				"fetchConfigLatency", fetchConfigLatency)
			return chainCache.data, nil
		}
		c.lggr.Errorw("Failed to refresh cache, no old data available",
			"chain", chainSel,
			"error", err,
			"fetchConfigLatency", fetchConfigLatency)
		return ChainConfigSnapshot{}, fmt.Errorf("failed to refresh cache for chain %d: %w", chainSel, err)
	}

	chainCache.data = newData
	chainCache.lastRefresh = time.Now()

	c.lggr.Debugw("Successfully refreshed cache",
		"chain", chainSel,
		"fetchConfigLatency", fetchConfigLatency)

	return newData, nil
}

func (c *configPoller) fetchChainConfig(
	ctx context.Context,
	chainSel cciptypes.ChainSelector) (ChainConfigSnapshot, error) {

	reader, exists := c.reader.contractReaders[chainSel]
	if !exists {
		return ChainConfigSnapshot{}, fmt.Errorf("no contract reader for chain %d", chainSel)
	}

	requests := c.reader.prepareBatchConfigRequests(chainSel)
	batchResult, skipped, err := reader.ExtendedBatchGetLatestValues(ctx, requests, true)
	if err != nil {
		return ChainConfigSnapshot{}, fmt.Errorf("batch get latest values for chain %d: %w", chainSel, err)
	}

	if len(skipped) > 0 {
		c.lggr.Infow("some contracts were skipped due to no bindings",
			"chain", chainSel,
			"contracts", skipped)
	}

	return c.reader.processConfigResults(chainSel, batchResult)
}

// resultProcessor defines a function type for processing individual results
type resultProcessor func(interface{}) error

// Ensure configCache implements ConfigCache
var _ ConfigPoller = (*configPoller)(nil)
