package reader

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/services"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
)

var _ ConfigPoller = (*configPollerV2)(nil)
var _ services.Service = (*configPollerV2)(nil)

// configPollerV2 provides the same API and cache structure as config_poller.go but always
// batch fetches both ChainConfigSnapshots and StaticSourceChainConfigs on any cache miss
// and uses chainAccessor for batch refreshes.
type configPollerV2 struct {
	services.StateMachine
	sync.RWMutex
	chainCaches       map[cciptypes.ChainSelector]*chainCache
	refreshPeriod     time.Duration
	lggr              logger.Logger
	chainAccessors    map[cciptypes.ChainSelector]cciptypes.ChainAccessor
	destChainSelector cciptypes.ChainSelector

	// Track known source chains for each destination chain
	knownSourceChains map[cciptypes.ChainSelector]map[cciptypes.ChainSelector]bool

	// Background polling control
	stopChan    chan struct{}
	wg          sync.WaitGroup
	failedPolls atomic.Uint32
}

func newConfigPollerV2(
	lggr logger.Logger,
	accessors map[cciptypes.ChainSelector]cciptypes.ChainAccessor,
	destChainSelector cciptypes.ChainSelector,
	refreshPeriod time.Duration,
) *configPollerV2 {
	return &configPollerV2{
		chainCaches:       make(map[cciptypes.ChainSelector]*chainCache),
		refreshPeriod:     refreshPeriod,
		chainAccessors:    accessors,
		destChainSelector: destChainSelector,
		lggr:              lggr,
		knownSourceChains: make(map[cciptypes.ChainSelector]map[cciptypes.ChainSelector]bool),
		stopChan:          make(chan struct{}),
	}
}

func (c *configPollerV2) Start(ctx context.Context) error {
	return c.StartOnce("configPollerV2", func() error {
		c.startBackgroundPolling()
		c.lggr.Info("Background poller started (v2)")
		return nil
	})
}

func (c *configPollerV2) Close() error {
	return c.StopOnce("configPollerV2", func() error {
		close(c.stopChan)
		c.wg.Wait()
		c.failedPolls.Store(0)
		return nil
	})
}

func (c *configPollerV2) Name() string {
	return c.lggr.Name()
}

func (c *configPollerV2) HealthReport() map[string]error {
	// Check if consecutive failed polls exceeds the maximum
	failCount := c.failedPolls.Load()
	if failCount >= MaxFailedPolls {
		c.SvcErrBuffer.Append(fmt.Errorf("polling failed %d times in a row", MaxFailedPolls))
	}

	return map[string]error{c.Name(): c.Healthy()}
}

func (c *configPollerV2) Ready() error {
	return c.StateMachine.Ready()
}

func (c *configPollerV2) startBackgroundPolling() {
	c.wg.Add(1)
	go func() {
		defer c.wg.Done()
		ticker := time.NewTicker(c.refreshPeriod)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				c.refreshAllKnownChains()
			case <-c.stopChan:
				return
			}
		}
	}()
}

// GetChainConfig retrieves the ChainConfigSnapshot for a specific chain. If the config is not in cache,
// it will issue a batch refresh to fetch all configs.
func (c *configPollerV2) GetChainConfig(
	ctx context.Context,
	chainSel cciptypes.ChainSelector,
) (ChainConfigSnapshot, error) {
	// Confirm we have an accessor for this chain
	_, err := getChainAccessor(c.chainAccessors, chainSel)
	if err != nil {
		c.lggr.Errorw("No chain accessor for chain", "chain", chainSel, "error", err)
		return ChainConfigSnapshot{}, fmt.Errorf("no chain accessor for %s: %w", chainSel, err)
	}

	chainCache := c.getOrCreateChainCache(chainSel)

	// Check if we have any data in cache
	chainCache.chainConfigMu.RLock()
	if !chainCache.chainConfigRefresh.IsZero() {
		defer chainCache.chainConfigMu.RUnlock()
		c.lggr.Debugw("Returning cached chain config",
			"chain", chainSel,
			"cacheAge", time.Since(chainCache.chainConfigRefresh))
		return chainCache.chainConfigData, nil
	}
	chainCache.chainConfigMu.RUnlock()

	// Cache miss: batch fetch all configs for this chain. Don't hold the lock while fetching.
	if err := c.batchRefreshChainAndSourceConfigs(ctx, chainSel); err != nil {
		return ChainConfigSnapshot{}, err
	}

	// Re-acquire read lock to return the data
	chainCache.chainConfigMu.RLock()
	defer chainCache.chainConfigMu.RUnlock()
	return chainCache.chainConfigData, nil
}

func (c *configPollerV2) GetOfframpSourceChainConfigs(
	ctx context.Context,
	destChain cciptypes.ChainSelector,
	sourceChains []cciptypes.ChainSelector,
) (map[cciptypes.ChainSelector]StaticSourceChainConfig, error) {
	filteredSourceChains := filterOutChainSelector(sourceChains, destChain)
	if len(filteredSourceChains) == 0 {
		return make(map[cciptypes.ChainSelector]StaticSourceChainConfig), nil
	}

	// Track all requested source chains for background refreshing
	for _, chain := range filteredSourceChains {
		if !c.trackSourceChain(destChain, chain) {
			c.lggr.Warnw("Could not track source chain for background refreshing",
				"destChain", destChain,
				"sourceChain", chain)
		}
	}

	destChainCache := c.getOrCreateChainCache(destChain)
	if destChainCache == nil {
		return nil, fmt.Errorf("failed to get chain cache for destination chain %s", destChain)
	}

	destChainCache.sourceChainMu.RLock()

	// Initialize results map
	cachedSourceConfigs := make(map[cciptypes.ChainSelector]StaticSourceChainConfig)
	var missingChains []cciptypes.ChainSelector

	// Check which chains exist in cache
	for _, chain := range filteredSourceChains {
		staticSourceChainConfig, exists := destChainCache.staticSourceChainConfigs[chain]
		if exists {
			cachedSourceConfigs[chain] = staticSourceChainConfig
		} else {
			// This chain isn't in cache yet
			missingChains = append(missingChains, chain)
		}
	}

	// If all chains are in cache, return them immediately
	if len(missingChains) == 0 {
		destChainCache.sourceChainMu.RUnlock()
		c.lggr.Debugw("All source chain configs found in cache",
			"destChain", destChain,
			"sourceChains", filteredSourceChains)
		return cachedSourceConfigs, nil
	}

	// Release lock before issuing batch refresh
	destChainCache.sourceChainMu.RUnlock()

	if err := c.batchRefreshChainAndSourceConfigs(ctx, destChain); err != nil {
		return nil, err
	}

	destChainCache.sourceChainMu.RLock()
	defer destChainCache.sourceChainMu.RUnlock()
	result := make(map[cciptypes.ChainSelector]StaticSourceChainConfig)
	for _, chain := range filteredSourceChains {
		if cfg, exists := destChainCache.staticSourceChainConfigs[chain]; exists {
			result[chain] = cfg
		}
	}
	return result, nil
}

// getOrCreateChainCache safely retrieves or creates a cache for a specific chain
func (c *configPollerV2) getOrCreateChainCache(chainSel cciptypes.ChainSelector) *chainCache {
	c.Lock()
	defer c.Unlock()
	if cache, exists := c.chainCaches[chainSel]; exists {
		return cache
	}

	// Verify we have an accessor for this chain
	_, err := getChainAccessor(c.chainAccessors, chainSel)
	if err != nil {
		c.lggr.Errorw("No chain accessor for chain", "chain", chainSel, "error", err)
		return nil
	}

	cache := &chainCache{
		staticSourceChainConfigs: make(map[cciptypes.ChainSelector]StaticSourceChainConfig),
	}
	c.chainCaches[chainSel] = cache
	return cache
}

// batchRefreshChainAndSourceConfigs fetches both ChainConfigSnapshot and StaticSourceChainConfigs for a specific
// chain using the chain's chainAccessor. It updates the cache with the results for both ChainConfigSnapshot and
// StaticSourceChainConfigs.
func (c *configPollerV2) batchRefreshChainAndSourceConfigs(ctx context.Context, chainSel cciptypes.ChainSelector) error {
	start := time.Now()

	// Use chainAccessor to fetch both ChainConfigSnapshot and SourceChainConfig
	accessor, err := getChainAccessor(c.chainAccessors, chainSel)
	if err != nil {
		c.lggr.Errorw("Failed to get chain accessor", "chain", chainSel, "error", err)
		return fmt.Errorf("failed to get chain accessor for %s: %w", chainSel, err)
	}
	chainConfigSnapshot, sourceChainConfigs, err := accessor.GetAllConfigsForChain(
		ctx,
		c.destChainSelector,
		c.knownSourceChains[chainSel],
	)
	if err != nil {
		c.lggr.Errorw("Failed batch fetch via chainAccessor", "chain", chainSel, "error", err)
		return err
	}

	cache := c.getOrCreateChainCache(chainSel)

	// Acquire ChainConfigSnapshot lock and update
	if chainConfigSnapshot != nil {
		cache.chainConfigMu.Lock()
		cache.chainConfigData = chainConfigSnapshot
		cache.chainConfigRefresh = time.Now()
		cache.chainConfigMu.Unlock()
	}

	// Acquire StaticSourceChainConfigs lock and update
	if len(sourceChainConfigs) > 0 {
		cache.sourceChainMu.Lock()
		for chain, cfg := range sourceChainConfigs {
			cache.staticSourceChainConfigs[chain] = cfg
		}
		cache.sourceChainRefresh = time.Now()
		cache.sourceChainMu.Unlock()
	}
	c.lggr.Debugw("Batch refreshed configs via chainAccessor", "chain", chainSel, "latency", time.Since(start))
	return nil
}

func (c *configPollerV2) refreshAllKnownChains() {
	chainsToRefresh := c.getChainsToRefresh()

	refreshFailed := false
	for _, chain := range chainsToRefresh {
		ctx, cancel := context.WithTimeout(context.Background(), bgRefreshTimeout)
		if err := c.batchRefreshChainAndSourceConfigs(ctx, chain); err != nil {
			refreshFailed = true
			c.lggr.Warnw("Failed to batch refresh configs", "chain", chain, "error", err)
		}
		cancel()
	}
	if refreshFailed {
		c.failedPolls.Add(1)
		failCount := c.failedPolls.Load()
		c.lggr.Warnw("Chain config refresh failed", "consecutiveFailures", failCount, "maxAllowed", MaxFailedPolls)
	} else if len(chainsToRefresh) > 0 {
		c.failedPolls.Store(0)
	}
}

// getChainsToRefresh returns all chains in the cache and their associated source chains
// This method acquires a read lock for the duration of its execution
func (c *configPollerV2) getChainsToRefresh() []cciptypes.ChainSelector {
	c.RLock()
	defer c.RUnlock()

	// Get all chains present in cache, including known source chains for the destination chain
	allChainsInCache := make(map[cciptypes.ChainSelector]struct{})
	for chainSel := range c.chainCaches {
		allChainsInCache[chainSel] = struct{}{}

		// If this chain is the destination chain and has source chains, gather them
		if chainSel != c.destChainSelector {
			continue
		}
		if sourceChains, exists := c.knownSourceChains[chainSel]; exists && len(sourceChains) > 0 {
			for sourceChain := range sourceChains {
				allChainsInCache[sourceChain] = struct{}{}
			}
		}
	}

	allChainsToTrack := make([]cciptypes.ChainSelector, 0, len(allChainsInCache))
	for v := range allChainsInCache {
		allChainsToTrack = append(allChainsToTrack, v)
	}

	return allChainsToTrack
}

// trackSourceChain adds a source chain to the known source chains map for a destination chain
func (c *configPollerV2) trackSourceChain(destChain, sourceChain cciptypes.ChainSelector) bool {
	if destChain == sourceChain {
		return false // Don't track destination chain as its own source
	}

	// First check if we have a chain accessor for this destination chain
	if _, exists := c.chainAccessors[destChain]; !exists {
		c.lggr.Debugw("Cannot track source chain - no chain accessor for dest chain",
			"destChain", destChain,
			"sourceChain", sourceChain)
		return false
	}

	c.Lock()
	defer c.Unlock()

	if _, exists := c.knownSourceChains[destChain]; !exists {
		c.knownSourceChains[destChain] = make(map[cciptypes.ChainSelector]bool)
	}

	c.knownSourceChains[destChain][sourceChain] = true
	return true
}
