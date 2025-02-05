package reader

import (
	"context"
	"fmt"
	"sync"
	"time"

	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

// configCache handles caching of chain configuration data for multiple chains.
// It is used by the ccipChainReader to store and retrieve configuration data,
// avoiding unnecessary contract calls and improving performance.
type configCache struct {
	sync.RWMutex
	chainCaches   map[cciptypes.ChainSelector]*chainCache
	refreshPeriod time.Duration
}

// chainCache represents the cache for a single chain.
// It stores the configuration data for a specific chain and manages
// the last refresh time to determine when the data needs to be updated.
type chainCache struct {
	sync.RWMutex
	data        ChainConfigSnapshot
	lastRefresh time.Time
}

// newConfigCache creates a new multi-chain config cache with the specified refresh period.
// The refresh period determines how often the cached data is considered stale and needs to be updated.
func newConfigCache(refreshPeriod time.Duration) *configCache {
	return &configCache{
		chainCaches:   make(map[cciptypes.ChainSelector]*chainCache),
		refreshPeriod: refreshPeriod,
	}
}

// getOrCreateChainCache safely retrieves or creates a cache for a specific chain.
// It ensures thread safety by using locks when accessing the cache map.
func (c *configCache) getOrCreateChainCache(chainSel cciptypes.ChainSelector) *chainCache {
	c.Lock()
	defer c.Unlock()

	if cache, exists := c.chainCaches[chainSel]; exists {
		return cache
	}

	cache := &chainCache{}
	c.chainCaches[chainSel] = cache
	return cache
}

// getChainConfig returns the cached chain configuration for a specific chain
func (r *ccipChainReader) getChainConfig(
	ctx context.Context,
	chainSel cciptypes.ChainSelector) (ChainConfigSnapshot, error) {
	chainCache := r.cache.getOrCreateChainCache(chainSel)

	chainCache.RLock()
	timeSinceLastRefresh := time.Since(chainCache.lastRefresh)
	if timeSinceLastRefresh < r.cache.refreshPeriod {
		defer chainCache.RUnlock()
		r.lggr.Infow("Cache hit",
			"chain", chainSel,
			"timeSinceLastRefresh", timeSinceLastRefresh,
			"refreshPeriod", r.cache.refreshPeriod)
		return chainCache.data, nil
	}
	chainCache.RUnlock()

	return r.refreshChainCache(ctx, chainSel)
}

func (r *ccipChainReader) refreshChainCache(
	ctx context.Context,
	chainSel cciptypes.ChainSelector) (ChainConfigSnapshot, error) {
	chainCache := r.cache.getOrCreateChainCache(chainSel)

	chainCache.Lock()
	defer chainCache.Unlock()

	timeSinceLastRefresh := time.Since(chainCache.lastRefresh)
	if timeSinceLastRefresh < r.cache.refreshPeriod {
		r.lggr.Infow("Cache was refreshed by another goroutine",
			"chain", chainSel,
			"timeSinceLastRefresh", timeSinceLastRefresh)
		return chainCache.data, nil
	}

	startTime := time.Now()
	newData, err := r.fetchChainConfig(ctx, chainSel)
	refreshDuration := time.Since(startTime)

	if err != nil {
		if !chainCache.lastRefresh.IsZero() {
			r.lggr.Warnw("Failed to refresh cache, using old data",
				"chain", chainSel,
				"error", err,
				"lastRefresh", chainCache.lastRefresh,
				"refreshDuration", refreshDuration)
			return chainCache.data, nil
		}
		r.lggr.Errorw("Failed to refresh cache, no old data available",
			"chain", chainSel,
			"error", err,
			"refreshDuration", refreshDuration)
		return ChainConfigSnapshot{}, fmt.Errorf("failed to refresh cache for chain %d: %w", chainSel, err)
	}

	chainCache.data = newData
	chainCache.lastRefresh = time.Now()

	return newData, nil
}
