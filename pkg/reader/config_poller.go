package reader

import (
	"context"
	"fmt"
	"sync"
	"time"

	commonconfig "github.com/smartcontractkit/chainlink-common/pkg/config"

	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
)

// ConfigPoller defines the interface for caching chain configuration data
type ConfigPoller interface {
	// GetChainConfig retrieves the cached configuration for a chain
	GetChainConfig(ctx context.Context, chainSel cciptypes.ChainSelector) (ChainConfigSnapshot, error)
	// RefreshChainConfig forces a refresh of the chain configuration
	RefreshChainConfig(ctx context.Context, chainSel cciptypes.ChainSelector) (ChainConfigSnapshot, error)
	// GetOfframpSourceChainConfigs retrieves cached source chain configurations.
	// This method returns StaticSourceChainConfig objects which deliberately exclude MinSeqNr.
	// For obtaining sequence numbers, use ccipChainReader.GetLatestMinSeqNrs instead.
	GetOfframpSourceChainConfigs(
		ctx context.Context,
		destChain cciptypes.ChainSelector,
		sourceChains []cciptypes.ChainSelector) (map[cciptypes.ChainSelector]StaticSourceChainConfig, error)
	// RefreshSourceChainConfigs forces a refresh of source chain configurations
	RefreshSourceChainConfigs(
		ctx context.Context,
		destChain cciptypes.ChainSelector,
		sourceChains []cciptypes.ChainSelector) (map[cciptypes.ChainSelector]StaticSourceChainConfig, error)
}

// configPoller handles caching of chain configuration data for multiple chains.
// It is used by the ccipChainReader to store and retrieve configuration data,
// avoiding unnecessary contract calls and improving performance.
// configPoller handles caching of chain configuration data for multiple chains
type configPoller struct {
	sync.RWMutex
	chainCaches   map[cciptypes.ChainSelector]*chainCache
	refreshPeriod commonconfig.Duration
	reader        *ccipChainReader // Reference to the reader for fetching configs
	lggr          logger.Logger
}

// chainCache represents the cache for a single chain.
// It stores the configuration data for a specific chain and manages
// the last refresh time to determine when the data needs to be updated.
type chainCache struct {
	// Chain config specific lock and data
	chainConfigMu      sync.RWMutex
	chainConfigData    ChainConfigSnapshot
	chainConfigRefresh time.Time

	// Source chain config specific lock and data
	sourceChainMu            sync.RWMutex
	staticSourceChainConfigs map[cciptypes.ChainSelector]StaticSourceChainConfig
	sourceChainRefresh       time.Time // Single timestamp for all source chain configs
}

// newConfigPoller creates a new config cache instance
func newConfigPoller(
	lggr logger.Logger,
	reader *ccipChainReader,
	refreshPeriod commonconfig.Duration,
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

	cache := &chainCache{
		staticSourceChainConfigs: make(map[cciptypes.ChainSelector]StaticSourceChainConfig),
	}
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

	chainCache.chainConfigMu.RLock()
	timeSinceLastRefresh := time.Since(chainCache.chainConfigRefresh)
	if timeSinceLastRefresh < c.refreshPeriod.Duration() {
		defer chainCache.chainConfigMu.RUnlock()
		c.lggr.Debugw("Cache hit",
			"chain", chainSel,
			"timeSinceLastRefresh", timeSinceLastRefresh,
			"refreshPeriod", c.refreshPeriod)
		return chainCache.chainConfigData, nil
	}
	chainCache.chainConfigMu.RUnlock()

	return c.RefreshChainConfig(ctx, chainSel)
}

// GetOfframpSourceChainConfigs retrieves cached source chain configurations
func (c *configPoller) GetOfframpSourceChainConfigs(
	ctx context.Context,
	destChain cciptypes.ChainSelector,
	sourceChains []cciptypes.ChainSelector,
) (map[cciptypes.ChainSelector]StaticSourceChainConfig, error) {
	// Verify we have a reader for the destination chain
	if _, exists := c.reader.contractReaders[destChain]; !exists {
		c.lggr.Errorw("No contract reader for destination chain", "chain", destChain)
		return nil, fmt.Errorf("no contract reader for destination chain %d", destChain)
	}

	// Filter out destination chain from source chains
	filteredSourceChains := filterOutChainSelector(sourceChains, destChain)
	if len(filteredSourceChains) == 0 {
		return make(map[cciptypes.ChainSelector]StaticSourceChainConfig), nil
	}

	// Get or create cache for the destination chain
	chainCache := c.getOrCreateChainCache(destChain)
	if chainCache == nil {
		return nil, fmt.Errorf("failed to create cache for chain %d", destChain)
	}

	chainCache.sourceChainMu.RLock()

	// Initialize results map and track which chains need to be fetched
	cachedSourceConfigs := make(map[cciptypes.ChainSelector]StaticSourceChainConfig)
	var chainsToFetch []cciptypes.ChainSelector

	// Check if the global refresh time has expired
	needsGlobalRefresh := chainCache.sourceChainRefresh.IsZero() ||
		time.Since(chainCache.sourceChainRefresh) >= c.refreshPeriod.Duration()

	c.lggr.Debugw("Checking if refresh needed",
		"sourceChainRefresh", chainCache.sourceChainRefresh,
		"timeSince", time.Since(chainCache.sourceChainRefresh),
		"refreshPeriod", c.refreshPeriod,
		"needsRefresh", needsGlobalRefresh)

	// Determine which chains need to be fetched
	for _, chain := range filteredSourceChains {
		config, exists := chainCache.staticSourceChainConfigs[chain]
		if !exists || needsGlobalRefresh {
			chainsToFetch = append(chainsToFetch, chain)
		} else {
			// Use cached version
			cachedSourceConfigs[chain] = config
		}
	}

	// If all chains are in cache and fresh, return them
	if len(chainsToFetch) == 0 {
		chainCache.sourceChainMu.RUnlock()
		c.lggr.Debugw("All source chain configs found in cache",
			"destChain", destChain,
			"sourceChains", sourceChains)
		return cachedSourceConfigs, nil
	}

	chainCache.sourceChainMu.RUnlock()

	// Need to fetch some chains from the contract
	c.lggr.Debugw("Fetching source chain configs",
		"destChain", destChain,
		"chainsToFetch", chainsToFetch)

	// Get the missing configs
	newCachedConfigs, err := c.RefreshSourceChainConfigs(ctx, destChain, chainsToFetch)
	if err != nil {
		return nil, err
	}

	// Merge the new configs with existing cached results
	for chain, config := range newCachedConfigs {
		cachedSourceConfigs[chain] = config
	}

	return cachedSourceConfigs, nil
}

// RefreshSourceChainConfigs forces a refresh of source chain configurations
func (c *configPoller) RefreshSourceChainConfigs(
	ctx context.Context,
	destChain cciptypes.ChainSelector,
	chainsToFetch []cciptypes.ChainSelector,
) (map[cciptypes.ChainSelector]StaticSourceChainConfig, error) {
	if len(chainsToFetch) == 0 {
		return make(map[cciptypes.ChainSelector]StaticSourceChainConfig), nil
	}

	// Get the chain cache
	chainCache := c.getOrCreateChainCache(destChain)
	if chainCache == nil {
		return nil, fmt.Errorf("failed to get cache for chain %d", destChain)
	}

	// Fetch configs from the contract
	startTime := time.Now()
	sourceChainConfigs, err := c.reader.fetchFreshSourceChainConfigs(ctx, destChain, chainsToFetch)
	fetchConfigLatency := time.Since(startTime)

	if err != nil {
		c.lggr.Errorw("Failed to fetch source chain configs",
			"destChain", destChain,
			"chainsToFetch", chainsToFetch,
			"error", err,
			"fetchConfigLatency", fetchConfigLatency)
		return nil, fmt.Errorf("fetch source chain configs: %w", err)
	}

	// Update the cache with new configs
	chainCache.sourceChainMu.Lock()

	// Initialize the map if needed
	if chainCache.staticSourceChainConfigs == nil {
		chainCache.staticSourceChainConfigs = make(map[cciptypes.ChainSelector]StaticSourceChainConfig)
	}

	// Update configs in the map
	for chain, config := range sourceChainConfigs {
		cachedConfig := staticSourceChainConfigFromSourceChainConfig(config)
		chainCache.staticSourceChainConfigs[chain] = cachedConfig
	}

	// Update the refresh timestamp
	chainCache.sourceChainRefresh = time.Now()

	chainCache.sourceChainMu.Unlock()

	c.lggr.Debugw("Successfully refreshed source chain configs",
		"destChain", destChain,
		"chainsCount", len(sourceChainConfigs),
		"fetchConfigLatency", fetchConfigLatency)

	result := make(map[cciptypes.ChainSelector]StaticSourceChainConfig, len(sourceChainConfigs))
	for chain, config := range sourceChainConfigs {
		result[chain] = staticSourceChainConfigFromSourceChainConfig(config)
	}

	return result, nil
}

// RefreshChainConfig forces a refresh of the chain configuration
func (c *configPoller) RefreshChainConfig(
	ctx context.Context,
	chainSel cciptypes.ChainSelector,
) (ChainConfigSnapshot, error) {
	chainCache := c.getOrCreateChainCache(chainSel)

	chainCache.chainConfigMu.Lock()
	defer chainCache.chainConfigMu.Unlock()

	// Double check if another goroutine has already refreshed
	timeSinceLastRefresh := time.Since(chainCache.chainConfigRefresh)
	if timeSinceLastRefresh < c.refreshPeriod.Duration() {
		c.lggr.Debugw("Cache was refreshed by another goroutine",
			"chain", chainSel,
			"timeSinceLastRefresh", timeSinceLastRefresh)
		return chainCache.chainConfigData, nil
	}

	startTime := time.Now()
	newData, err := c.fetchChainConfig(ctx, chainSel)
	fetchConfigLatency := time.Since(startTime)

	if err != nil {
		if !chainCache.chainConfigRefresh.IsZero() {
			c.lggr.Warnw("Failed to refresh cache, using old data",
				"chain", chainSel,
				"error", err,
				"lastRefresh", chainCache.chainConfigRefresh,
				"fetchConfigLatency", fetchConfigLatency)
			return chainCache.chainConfigData, nil
		}
		c.lggr.Errorw("Failed to refresh cache, no old data available",
			"chain", chainSel,
			"error", err,
			"fetchConfigLatency", fetchConfigLatency)
		return ChainConfigSnapshot{}, fmt.Errorf("failed to refresh cache for chain %d: %w", chainSel, err)
	}

	chainCache.chainConfigData = newData
	chainCache.chainConfigRefresh = time.Now()

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

// filterOutChainSelector removes a specified chain selector from a slice of chain selectors
func filterOutChainSelector(
	chains []cciptypes.ChainSelector,
	chainToFilter cciptypes.ChainSelector) []cciptypes.ChainSelector {
	if len(chains) == 0 {
		return nil
	}

	filtered := make([]cciptypes.ChainSelector, 0, len(chains))
	for _, chain := range chains {
		if chain != chainToFilter {
			filtered = append(filtered, chain)
		}
	}
	return filtered
}

// StaticSourceChainConfigFromSourceChainConfig creates a StaticSourceChainConfig from a SourceChainConfig,
// omitting the MinSeqNr field.
func staticSourceChainConfigFromSourceChainConfig(sc SourceChainConfig) StaticSourceChainConfig {
	return StaticSourceChainConfig{
		Router:                    sc.Router,
		IsEnabled:                 sc.IsEnabled,
		IsRMNVerificationDisabled: sc.IsRMNVerificationDisabled,
		OnRamp:                    sc.OnRamp,
	}
}

// resultProcessor defines a function type for processing individual results
type resultProcessor func(interface{}) error

// Ensure configCache implements ConfigPoller
var _ ConfigPoller = (*configPoller)(nil)
