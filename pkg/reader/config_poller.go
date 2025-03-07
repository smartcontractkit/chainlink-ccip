package reader

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/smartcontractkit/chainlink-ccip/pkg/consts"
	"github.com/smartcontractkit/chainlink-ccip/pkg/contractreader"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/types"
)

// ConfigPoller defines the interface for caching chain configuration data
type ConfigPoller interface {
	// GetChainConfig retrieves the cached configuration for a chain
	GetChainConfig(ctx context.Context, chainSel cciptypes.ChainSelector) (ChainConfigSnapshot, error)
	// RefreshChainConfig forces a refresh of the chain configuration
	RefreshChainConfig(ctx context.Context, chainSel cciptypes.ChainSelector) (ChainConfigSnapshot, error)
	// GetSourceChainConfigs retrieves cached source chain configurations
	GetSourceChainConfigs(
		ctx context.Context,
		destChain cciptypes.ChainSelector,
		sourceChains []cciptypes.ChainSelector) (map[cciptypes.ChainSelector]SourceChainConfig, error)
	// RefreshSourceChainConfigs forces a refresh of source chain configurations
	RefreshSourceChainConfigs(
		ctx context.Context,
		destChain cciptypes.ChainSelector,
		sourceChains []cciptypes.ChainSelector) (map[cciptypes.ChainSelector]SourceChainConfig, error)
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
	data               ChainConfigSnapshot
	lastRefresh        time.Time
	sourceChainConfigs map[cciptypes.ChainSelector]*sourceChainCacheEntry
}

// sourceChainCacheEntry holds cached data for a specific source chain
type sourceChainCacheEntry struct {
	config      SourceChainConfig
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

// GetSourceChainConfigs retrieves cached source chain configurations
func (c *configPoller) GetSourceChainConfigs(
	ctx context.Context,
	destChain cciptypes.ChainSelector,
	sourceChains []cciptypes.ChainSelector,
) (map[cciptypes.ChainSelector]SourceChainConfig, error) {
	// Verify we have a reader for the destination chain
	if _, exists := c.reader.contractReaders[destChain]; !exists {
		c.lggr.Errorw("No contract reader for destination chain", "chain", destChain)
		return nil, fmt.Errorf("no contract reader for destination chain %d", destChain)
	}

	// Get or create cache for the destination chain
	chainCache := c.getOrCreateChainCache(destChain)
	if chainCache == nil {
		return nil, fmt.Errorf("failed to create cache for chain %d", destChain)
	}

	chainCache.RLock()

	// Initialize results map and track which chains need to be fetched
	result := make(map[cciptypes.ChainSelector]SourceChainConfig)
	var chainsToFetch []cciptypes.ChainSelector

	// Check if sourceChainConfigs is initialized
	if chainCache.sourceChainConfigs == nil {
		// Need to fetch all chains
		chainsToFetch = sourceChains
	} else {
		// Check each source chain to see if it needs refreshing
		for _, chain := range sourceChains {
			if chain == destChain {
				continue // Skip if it's the destination chain
			}

			entry, exists := chainCache.sourceChainConfigs[chain]
			if !exists || time.Since(entry.lastRefresh) >= c.refreshPeriod {
				chainsToFetch = append(chainsToFetch, chain)
			} else {
				// Use cached version
				result[chain] = entry.config
			}
		}
	}

	// If all chains are in cache and fresh, return them
	if len(chainsToFetch) == 0 {
		chainCache.RUnlock()
		c.lggr.Debugw("All source chain configs found in cache",
			"destChain", destChain,
			"sourceChains", sourceChains)
		return result, nil
	}

	chainCache.RUnlock()

	// Need to fetch some chains from the contract
	c.lggr.Debugw("Fetching source chain configs",
		"destChain", destChain,
		"chainsToFetch", chainsToFetch)

	// Get the missing configs
	newConfigs, err := c.RefreshSourceChainConfigs(ctx, destChain, chainsToFetch)
	if err != nil {
		return nil, err
	}

	// Merge the new configs with existing cached results
	for chain, config := range newConfigs {
		result[chain] = config
	}

	return result, nil
}

// RefreshSourceChainConfigs forces a refresh of source chain configurations
func (c *configPoller) RefreshSourceChainConfigs(
	ctx context.Context,
	destChain cciptypes.ChainSelector,
	chainsToFetch []cciptypes.ChainSelector,
) (map[cciptypes.ChainSelector]SourceChainConfig, error) {
	if len(chainsToFetch) == 0 {
		return make(map[cciptypes.ChainSelector]SourceChainConfig), nil
	}

	// Get the chain cache
	chainCache := c.getOrCreateChainCache(destChain)
	if chainCache == nil {
		return nil, fmt.Errorf("failed to get cache for chain %d", destChain)
	}

	// Fetch configs from the contract
	startTime := time.Now()
	newConfigs, err := c.fetchSourceChainConfigs(ctx, destChain, chainsToFetch)
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
	chainCache.Lock()

	// Initialize the map if needed
	if chainCache.sourceChainConfigs == nil {
		chainCache.sourceChainConfigs = make(map[cciptypes.ChainSelector]*sourceChainCacheEntry)
	}

	now := time.Now()
	for chain, config := range newConfigs {
		chainCache.sourceChainConfigs[chain] = &sourceChainCacheEntry{
			config:      config,
			lastRefresh: now,
		}
	}

	chainCache.Unlock()

	c.lggr.Debugw("Successfully refreshed source chain configs",
		"destChain", destChain,
		"chainsCount", len(newConfigs),
		"fetchConfigLatency", fetchConfigLatency)

	return newConfigs, nil
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

// fetchSourceChainConfigs fetches source chain configs directly from contracts
func (c *configPoller) fetchSourceChainConfigs(
	ctx context.Context,
	destChain cciptypes.ChainSelector,
	sourceChains []cciptypes.ChainSelector,
) (map[cciptypes.ChainSelector]SourceChainConfig, error) {
	reader, exists := c.reader.contractReaders[destChain]
	if !exists {
		return nil, fmt.Errorf("no contract reader for chain %d", destChain)
	}

	// Prepare batch requests for the sourceChains
	contractBatch := make([]types.BatchRead, 0, len(sourceChains))
	validSourceChains := make([]cciptypes.ChainSelector, 0, len(sourceChains))

	for _, chain := range sourceChains {
		if chain == destChain {
			continue
		}

		validSourceChains = append(validSourceChains, chain)
		contractBatch = append(contractBatch, types.BatchRead{
			ReadName: consts.MethodNameGetSourceChainConfig,
			Params: map[string]any{
				"sourceChainSelector": chain,
			},
			ReturnVal: new(SourceChainConfig),
		})
	}

	// Execute batch request
	results, _, err := reader.ExtendedBatchGetLatestValues(
		ctx,
		contractreader.ExtendedBatchGetLatestValuesRequest{
			consts.ContractNameOffRamp: contractBatch,
		},
		false,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to get source chain configs: %w", err)
	}

	// Process results
	configs := make(map[cciptypes.ChainSelector]SourceChainConfig)

	for _, readResult := range results {
		if len(readResult) != len(validSourceChains) {
			return nil, fmt.Errorf("selectors and source chain configs length mismatch: sourceChains=%v, results=%v",
				validSourceChains, results)
		}

		for i, chain := range validSourceChains {
			v, err := readResult[i].GetResult()
			if err != nil {
				c.lggr.Errorw("Failed to get source chain config",
					"chain", chain,
					"error", err)
				return nil, fmt.Errorf("GetSourceChainConfig for chainSelector=%d failed: %w", chain, err)
			}

			cfg, ok := v.(*SourceChainConfig)
			if !ok {
				c.lggr.Errorw("Invalid result type from GetSourceChainConfig",
					"chain", chain,
					"type", fmt.Sprintf("%T", v))
				return nil, fmt.Errorf("invalid result type from GetSourceChainConfig for chainSelector=%d", chain)
			}

			// Store the config - we don't filter here as that's done at the reader level
			configs[chain] = *cfg
		}
	}

	return configs, nil
}

// resultProcessor defines a function type for processing individual results
type resultProcessor func(interface{}) error

// Ensure configCache implements ConfigPoller
var _ ConfigPoller = (*configPoller)(nil)
