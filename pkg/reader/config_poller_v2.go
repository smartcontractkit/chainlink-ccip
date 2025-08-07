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

	// Track known source chains for the destination chain. NOTE: config_poller.go's version of knownSourceChains is
	// a map[cciptypes.ChainSelector]map[cciptypes.ChainSelector]struct{} where the first key was the destination chain.
	// However, we never track any other dest chain besides the destChain configured in the CCIPReader in ccip.go, so we
	// can simplify this to a map[cciptypes.ChainSelector]struct{} since it's only used for a single destination chain.
	knownSourceChains map[cciptypes.ChainSelector]struct{}

	// Background polling control
	stopChan               chan struct{}
	wg                     sync.WaitGroup
	consecutiveFailedPolls atomic.Uint32
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
		knownSourceChains: make(map[cciptypes.ChainSelector]struct{}),
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
		c.consecutiveFailedPolls.Store(0)
		return nil
	})
}

func (c *configPollerV2) Name() string {
	return c.lggr.Name()
}

func (c *configPollerV2) HealthReport() map[string]error {
	// Check if consecutive failed polls exceeds the maximum
	failCount := c.consecutiveFailedPolls.Load()
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
) (cciptypes.ChainConfigSnapshot, error) {
	// Confirm we have an accessor for this chain
	_, err := getChainAccessor(c.chainAccessors, chainSel)
	if err != nil {
		c.lggr.Errorw("No chain accessor for chain", "chain", chainSel, "error", err)
		return cciptypes.ChainConfigSnapshot{}, fmt.Errorf("no chain accessor for %s: %w", chainSel, err)
	}

	cache := c.getOrCreateChainCache(chainSel)
	if cache == nil {
		return cciptypes.ChainConfigSnapshot{},
			fmt.Errorf("failed to get or create chain cache for chain %s", chainSel)
	}

	// Check if we have any data in cache
	cache.chainConfigMu.RLock()
	if !cache.chainConfigRefresh.IsZero() {
		defer cache.chainConfigMu.RUnlock()
		c.lggr.Debugw("Returning cached chain config",
			"chain", chainSel,
			"cacheAge", time.Since(cache.chainConfigRefresh))
		return cache.chainConfigData, nil
	}
	cache.chainConfigMu.RUnlock()

	// Cache miss: batch fetch all configs for this chain. Don't hold the lock while fetching.
	// TODO: alternatively, if we want to prevent multiple goroutines from fetching the same chain config (especially
	//     during node startup), we could block on this fetch if the cache is empty.
	if err := c.batchRefreshChainAndSourceConfigs(ctx, chainSel); err != nil {
		return cciptypes.ChainConfigSnapshot{}, err
	}

	// Re-acquire read lock to return the data
	cache.chainConfigMu.RLock()
	defer cache.chainConfigMu.RUnlock()
	return cache.chainConfigData, nil
}

func (c *configPollerV2) GetOfframpSourceChainConfigs(
	ctx context.Context,
	destChain cciptypes.ChainSelector,
	sourceChains []cciptypes.ChainSelector,
) (map[cciptypes.ChainSelector]StaticSourceChainConfig, error) {
	if destChain != c.destChainSelector {
		// TODO: Remove destChain arg from this interface.
		//  Based on current usage and callers of the existing configPoller, this function should never be called
		// 	with a different destChain than the configured destChainSelector stored in c.destChainSelector. See:
		//  ccip.go getOfframpSourceChainConfigs().
		return nil,
			fmt.Errorf("the destChain passed in should never be different from the configured destChainSelector: %s != %s",
				destChain, c.destChainSelector)
	}

	// Ensure we're not trying to fetch source chain configs for the destination chain itself
	filteredSourceChains := filterOutChainSelector(sourceChains, c.destChainSelector)
	if len(filteredSourceChains) == 0 {
		return make(map[cciptypes.ChainSelector]StaticSourceChainConfig), nil
	}

	// Add any new source chains to list of tracked source chains for background refreshing
	for _, chain := range filteredSourceChains {
		if !c.trackSourceChainForDest(chain) {
			c.lggr.Warnw("Could not track source chain for background refreshing",
				"destChain", c.destChainSelector,
				"sourceChain", chain)
		}
	}

	destChainCache := c.getOrCreateChainCache(c.destChainSelector)
	if destChainCache == nil {
		return nil, fmt.Errorf("failed to get chain cache for destination chain %s", c.destChainSelector)
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
			"destChain", c.destChainSelector,
			"sourceChains", filteredSourceChains)
		return cachedSourceConfigs, nil
	}

	// Release lock before issuing batch refresh
	destChainCache.sourceChainMu.RUnlock()

	if err := c.batchRefreshChainAndSourceConfigs(ctx, c.destChainSelector); err != nil {
		return nil, err
	}

	// Re-acquire the lock to return only the cached configs that were requested
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

// refreshAllKnownChains iterates through all known chains and issues a batch refresh for each.
func (c *configPollerV2) refreshAllKnownChains() {
	chainsToRefresh := c.getChainsToRefresh()

	refreshFailed := false
	for _, chain := range chainsToRefresh {
		ctx, cancel := context.WithTimeout(context.Background(), bgRefreshTimeout)
		c.lggr.Debugw("Issuing background refresh for known chain",
			"chain", chain, "destChain", c.destChainSelector)
		if err := c.batchRefreshChainAndSourceConfigs(ctx, chain); err != nil {
			refreshFailed = true
			c.lggr.Warnw("Failed to batch refresh configs", "chain", chain, "error", err)
		}
		cancel()
	}
	if refreshFailed {
		c.consecutiveFailedPolls.Add(1)
		failCount := c.consecutiveFailedPolls.Load()
		c.lggr.Warnw("Chain config refresh failed", "consecutiveFailures", failCount, "maxAllowed", MaxFailedPolls)
	} else if len(chainsToRefresh) > 0 {
		c.consecutiveFailedPolls.Store(0)
	}
}

// batchRefreshChainAndSourceConfigs fetches both ChainConfigSnapshot and StaticSourceChainConfigs for a specific
// chain using the chain's chainAccessor. It updates the cache with the results for both ChainConfigSnapshot and
// StaticSourceChainConfigs.
func (c *configPollerV2) batchRefreshChainAndSourceConfigs(
	ctx context.Context,
	chainSel cciptypes.ChainSelector,
) error {
	start := time.Now()
	fetchingForDestChain := chainSel == c.destChainSelector

	sourceChainSelectors := make([]cciptypes.ChainSelector, 0)
	if fetchingForDestChain {
		// Acquires read lock on 'c'
		sourceChainSelectors = c.getKnownSourceChainsForDestChain()
	}

	// Use chainAccessor to fetch ChainConfigSnapshot (and SourceChainConfigs if destChain)
	accessor, err := getChainAccessor(c.chainAccessors, chainSel)
	if err != nil {
		c.lggr.Errorw("Failed to get chain accessor", "chain", chainSel, "error", err)
		return fmt.Errorf("failed to get chain accessor for %s: %w", chainSel, err)
	}

	// NO LOCKING DURING IO
	chainConfigSnapshot, sourceChainConfigs, err := accessor.GetAllConfigsLegacy(
		ctx,
		c.destChainSelector,
		sourceChainSelectors,
	)
	if err != nil {
		c.lggr.Errorw("Failed batch fetch via chainAccessor", "chain", chainSel, "error", err)
		return err
	}

	cache := c.getOrCreateChainCache(chainSel)
	if cache == nil {
		return fmt.Errorf("failed to get chain cache for chain %s", chainSel)
	}

	// Acquire ChainConfigSnapshot lock and update
	cache.chainConfigMu.Lock()
	cache.chainConfigData = chainConfigSnapshot
	cache.chainConfigRefresh = time.Now()
	cache.chainConfigMu.Unlock()

	// Acquire StaticSourceChainConfigs lock and update
	if fetchingForDestChain && len(sourceChainConfigs) > 0 {
		cache.sourceChainMu.Lock()
		for chain, cfg := range sourceChainConfigs {
			cache.staticSourceChainConfigs[chain] = staticSourceChainConfigFromSourceChainConfig(cfg)
		}
		cache.sourceChainRefresh = time.Now()
		cache.sourceChainMu.Unlock()
	} else if !fetchingForDestChain && len(sourceChainConfigs) > 0 {
		c.lggr.Errorw("OffRamp SourceChainConfigs were returned when fetching configs from a source chain, "+
			"this is not expected",
			"destChainSelector", c.destChainSelector,
			"chainSel", chainSel,
			"sourceChainSelectors", sourceChainSelectors,
		)
	}
	c.lggr.Debugw("Batch refreshed configs via chainAccessor", "chain", chainSel, "latency", time.Since(start))
	return nil
}

// getChainsToRefresh returns all chains in the cache in addition to all of their associated source chains.
// This method acquires a read lock for the duration of its execution
func (c *configPollerV2) getChainsToRefresh() []cciptypes.ChainSelector {
	c.RLock()
	defer c.RUnlock()

	// Init empty Set to get all chains present in cache, including known source chains for the destination chain
	allChainsInCache := make(map[cciptypes.ChainSelector]struct{})
	for chainSel := range c.chainCaches {
		allChainsInCache[chainSel] = struct{}{}
	}

	// Add all known source chains for the destination chain
	if len(c.knownSourceChains) > 0 {
		for sourceChain := range c.knownSourceChains {
			allChainsInCache[sourceChain] = struct{}{}
		}
	}

	allChainsToTrack := make([]cciptypes.ChainSelector, 0, len(allChainsInCache))
	for v := range allChainsInCache {
		allChainsToTrack = append(allChainsToTrack, v)
	}

	return allChainsToTrack
}

// trackSourceChain adds a source chain to the known source chains map for a destination chain
func (c *configPollerV2) trackSourceChainForDest(sourceChain cciptypes.ChainSelector) bool {
	if c.destChainSelector == sourceChain {
		c.lggr.Debugw("Skipping tracking source chain - destination chain is the same as source chain",
			"destChain", c.destChainSelector,
			"sourceChain", sourceChain)
		return false
	}

	// Check if we have a chain accessor for the dest chain. We always should.
	if _, exists := c.chainAccessors[c.destChainSelector]; !exists {
		c.lggr.Errorw("Cannot track source chain - no chain accessor for dest chain",
			"destChain", c.destChainSelector,
			"sourceChain", sourceChain)
		return false
	}

	c.Lock()
	defer c.Unlock()

	// Add the source chain to the knownSourceChains map for the destination chain
	c.knownSourceChains[sourceChain] = struct{}{}
	return true
}

// getKnownSourceChainsForDestChain is a helper function that retrieves all known source chains for the destination
// chain and return them as a slice of ChainSelectors.
func (c *configPollerV2) getKnownSourceChainsForDestChain() []cciptypes.ChainSelector {
	c.RLock()
	defer c.RUnlock()
	sourceChains := make([]cciptypes.ChainSelector, 0)
	for sourceChain := range c.knownSourceChains {
		sourceChains = append(sourceChains, sourceChain)
	}
	return sourceChains
}
