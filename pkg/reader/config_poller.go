package reader

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/smartcontractkit/chainlink-ccip/pkg/consts"
	"github.com/smartcontractkit/chainlink-ccip/pkg/contractreader"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/services"
	"github.com/smartcontractkit/chainlink-common/pkg/types"
)

// refreshAllKnownChains refreshes all known chains in background using batched requests where possible
// bgRefreshTimeout defines the timeout for background refresh operations
const (
	bgRefreshTimeout = 30 * time.Second
	MaxFailedPolls   = 10
)

// ConfigPoller defines the interface for caching chain configuration data
type ConfigPoller interface {
	// GetChainConfig retrieves the cached configuration for a chain
	GetChainConfig(ctx context.Context, chainSel cciptypes.ChainSelector) (ChainConfigSnapshot, error)
	// GetOfframpSourceChainConfigs retrieves cached source chain configurations
	GetOfframpSourceChainConfigs(
		ctx context.Context,
		destChain cciptypes.ChainSelector,
		sourceChains []cciptypes.ChainSelector) (map[cciptypes.ChainSelector]StaticSourceChainConfig, error)
	services.Service
}

// configPoller handles caching of chain configuration data for multiple chains
type configPoller struct {
	services.StateMachine // Embeds the StateMachine for lifecycle management

	sync.RWMutex
	chainCaches   map[cciptypes.ChainSelector]*chainCache
	refreshPeriod time.Duration
	reader        ccipReaderInternal
	lggr          logger.Logger

	// Track known source chains for each destination chain
	knownSourceChains map[cciptypes.ChainSelector]map[cciptypes.ChainSelector]bool

	// Background polling control
	stopChan chan struct{}
	wg       sync.WaitGroup

	// Track consecutive failed polls
	failedPolls atomic.Uint32
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
	reader ccipReaderInternal,
	refreshPeriod time.Duration,
) *configPoller {
	return &configPoller{
		chainCaches:       make(map[cciptypes.ChainSelector]*chainCache),
		refreshPeriod:     refreshPeriod,
		reader:            reader,
		lggr:              lggr,
		knownSourceChains: make(map[cciptypes.ChainSelector]map[cciptypes.ChainSelector]bool),
		stopChan:          make(chan struct{}),
	}
}

// startBackgroundPolling starts a goroutine that periodically refreshes the cache
func (c *configPoller) startBackgroundPolling() {
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

func (c *configPoller) Start(ctx context.Context) error {
	return c.StartOnce("ConfigPoller", func() error {
		c.startBackgroundPolling()
		c.lggr.Info("Background poller started")
		return nil
	})
}

func (c *configPoller) Close() error {
	return c.StopOnce("ConfigPoller", func() error {
		close(c.stopChan)
		c.wg.Wait()
		// Reset failed polls counter on shutdown
		c.failedPolls.Store(0)
		return nil
	})
}

func (c *configPoller) Name() string {
	return c.lggr.Name()
}

func (c *configPoller) HealthReport() map[string]error {
	// Check if consecutive failed polls exceeds the maximum
	failCount := c.failedPolls.Load()
	if failCount >= MaxFailedPolls {
		c.SvcErrBuffer.Append(fmt.Errorf("polling failed %d times in a row", MaxFailedPolls))
	}

	return map[string]error{c.Name(): c.Healthy()}
}

func (c *configPoller) Ready() error {
	return c.StateMachine.Ready()
}

// getChainsToRefresh returns all chains in the cache and their associated source chains
// This method acquires a read lock for the duration of its execution
func (c *configPoller) getChainsToRefresh() (
	[]cciptypes.ChainSelector,
	map[cciptypes.ChainSelector][]cciptypes.ChainSelector) {
	c.RLock()
	defer c.RUnlock()

	// Get all chains to refresh
	chainSelectors := make([]cciptypes.ChainSelector, 0, len(c.chainCaches))
	sourceChainsMap := make(map[cciptypes.ChainSelector][]cciptypes.ChainSelector)

	for chainSel := range c.chainCaches {
		chainSelectors = append(chainSelectors, chainSel)

		// If this chain is the destination chain and has source chains, gather them
		if chainSel != c.reader.getDestChain() {
			continue
		}
		if sourceChains, exists := c.knownSourceChains[chainSel]; exists && len(sourceChains) > 0 {
			sourceList := make([]cciptypes.ChainSelector, 0, len(sourceChains))
			for sourceChain := range sourceChains {
				sourceList = append(sourceList, sourceChain)
			}
			sourceChainsMap[chainSel] = sourceList
		}
	}

	return chainSelectors, sourceChainsMap
}

// refreshAllKnownChains refreshes all known chains in background using batched requests where possible
func (c *configPoller) refreshAllKnownChains() {
	// Get all chains to refresh using the helper method
	chainSelectors, sourceChainsMap := c.getChainsToRefresh()

	// Log the starting of refresh cycle
	c.lggr.Debugw("Starting refresh cycle for known chains",
		"chainsCount", len(chainSelectors),
		"refreshPeriod", c.refreshPeriod)

	// Track success of this refresh cycle
	refreshFailed := false

	// Refresh each chain (and its source chains if applicable)
	for _, chainSel := range chainSelectors {
		ctx, cancel := context.WithTimeout(context.Background(), bgRefreshTimeout)

		// If this is the destination chain and has source chains, use batch refresh
		if chainSel == c.reader.getDestChain() && len(sourceChainsMap[chainSel]) > 0 {
			sourceChains := sourceChainsMap[chainSel]
			err := c.batchRefreshChainAndSourceConfigs(ctx, chainSel, sourceChains)
			if err != nil {
				refreshFailed = true
				c.lggr.Warnw("Failed to batch refresh chain and source configs",
					"destChain", chainSel,
					"sourceChains", sourceChains,
					"error", err)
				// Fall back to individual refreshes if batch fails
				if _, err := c.refreshChainConfig(ctx, chainSel); err != nil {
					c.lggr.Warnw("Failed to refresh chain config in background",
						"chain", chainSel, "error", err)
				}

				_, err = c.refreshSourceChainConfigs(ctx, chainSel, sourceChains)
				if err != nil {
					c.lggr.Warnw("Failed to refresh source chain configs in background",
						"destChain", chainSel, "sourceChains", sourceChains, "error", err)
				}
			}
		} else {
			// This is a source chain or a chain without source chains, just refresh its config
			if _, err := c.refreshChainConfig(ctx, chainSel); err != nil {
				c.lggr.Warnw("Failed to refresh chain config in background",
					"chain", chainSel, "error", err)
			}
		}

		cancel()
	}

	// Update the failed polls counter
	if refreshFailed {
		c.failedPolls.Add(1)
		failCount := c.failedPolls.Load()
		c.lggr.Warnw("Chain config refresh failed",
			"consecutiveFailures", failCount,
			"maxAllowed", MaxFailedPolls)
	} else if len(chainSelectors) > 0 {
		// Only reset on successful refresh with actual chains
		c.failedPolls.Store(0) // Reset counter on success
	}

}

// extractStandardChainConfigResults creates a copy of the batch results with only the standard
// chain config results (limiting OffRamp results to the standard count)
func (c *configPoller) extractStandardChainConfigResults(
	batchResult types.BatchGetLatestValuesResult,
	standardOffRampRequestCount int,
) types.BatchGetLatestValuesResult {
	chainConfigResultsCopy := make(types.BatchGetLatestValuesResult)

	for contract, results := range batchResult {
		if contract.Name == consts.ContractNameOffRamp && len(results) > standardOffRampRequestCount {
			// Only include the standard results (first N results)
			chainConfigResultsCopy[contract] = results[:standardOffRampRequestCount]
		} else {
			// Copy as-is
			chainConfigResultsCopy[contract] = results
		}
	}

	return chainConfigResultsCopy
}

// processSourceChainResults extracts and processes source chain config results from the batch
func (c *configPoller) processSourceChainResults(
	batchResult types.BatchGetLatestValuesResult,
	standardOffRampRequestCount int,
	filteredSourceChains []cciptypes.ChainSelector,
) map[cciptypes.ChainSelector]SourceChainConfig {
	sourceConfigs := make(map[cciptypes.ChainSelector]SourceChainConfig)

	// Find the OffRamp results
	for contract, results := range batchResult {
		if contract.Name == consts.ContractNameOffRamp && len(results) > standardOffRampRequestCount {
			// Extract just the source chain results (everything after standard results)
			sourceChainResults := results[standardOffRampRequestCount:]

			if len(sourceChainResults) != len(filteredSourceChains) {
				c.lggr.Warnw("Source chain result count mismatch",
					"expected", len(filteredSourceChains),
					"got", len(sourceChainResults))
			} else {
				// Process each source chain result
				for i, chain := range filteredSourceChains {
					if i >= len(sourceChainResults) {
						continue
					}

					v, err := sourceChainResults[i].GetResult()
					if err != nil {
						c.lggr.Errorw("Failed to get source chain config",
							"chain", chain,
							"error", err)
						continue
					}

					cfg, ok := v.(*SourceChainConfig)
					if !ok {
						c.lggr.Errorw("Invalid result type from GetSourceChainConfig",
							"chain", chain,
							"type", fmt.Sprintf("%T", v))
						continue
					}

					// Store the config
					sourceConfigs[chain] = *cfg
				}
			}

			break // Found and processed the OffRamp results
		}
	}

	return sourceConfigs
}

// Helper to prepare source chain queries
func (c *configPoller) prepareSourceChainQueries(sourceChains []cciptypes.ChainSelector) []types.BatchRead {
	sourceConfigQueries := make([]types.BatchRead, 0, len(sourceChains))
	for _, chain := range sourceChains {
		sourceConfigQueries = append(sourceConfigQueries, types.BatchRead{
			ReadName: consts.MethodNameGetSourceChainConfig,
			Params: map[string]any{
				"sourceChainSelector": chain,
			},
			ReturnVal: new(SourceChainConfig),
		})
	}
	return sourceConfigQueries
}

// Helper to find and append source chain queries to existing requests
func (c *configPoller) appendSourceQueriesToRequests(
	chainConfigRequests contractreader.ExtendedBatchGetLatestValuesRequest,
	sourceQueries []types.BatchRead) (int, bool) {

	var standardOffRampRequestCount int
	var found bool

	// Find if OffRamp already exists and get standard request count
	for contractName, requests := range chainConfigRequests {
		if contractName == consts.ContractNameOffRamp {
			standardOffRampRequestCount = len(requests)

			// Append source queries if we have any
			if len(sourceQueries) > 0 {
				chainConfigRequests[contractName] = append(requests, sourceQueries...)
			}
			found = true
			break
		}
	}

	return standardOffRampRequestCount, found
}

// Main batch refresh function using the helper functions
func (c *configPoller) batchRefreshChainAndSourceConfigs(
	ctx context.Context,
	destChain cciptypes.ChainSelector,
	sourceChains []cciptypes.ChainSelector,
) error {
	startTime := time.Now()

	// 1. Prepare the standard chain config request
	chainConfigRequests := c.reader.prepareBatchConfigRequests(destChain)

	// 2. Filter source chains and prepare queries
	filteredSourceChains := filterOutChainSelector(sourceChains, destChain)
	sourceQueries := c.prepareSourceChainQueries(filteredSourceChains)

	// 3. Append source queries to existing requests and get standard count
	standardOffRampRequestCount, _ := c.appendSourceQueriesToRequests(chainConfigRequests, sourceQueries)

	// 4. Get the contract reader for this chain
	reader, exists := c.reader.getContractReader(destChain)
	if !exists {
		return fmt.Errorf("no contract reader for chain %d", destChain)
	}

	// 5. Execute the combined batch request
	batchResult, skipped, err := reader.ExtendedBatchGetLatestValues(ctx, chainConfigRequests, true)
	if err != nil {
		return fmt.Errorf("batch get latest values for chain %d: %w", destChain, err)
	}

	if len(skipped) > 0 {
		c.lggr.Infow("some contracts were skipped due to no bindings",
			"chain", destChain,
			"contracts", skipped)
	}

	// 6. Extract and process standard chain config results
	standardResultsCopy := c.extractStandardChainConfigResults(batchResult, standardOffRampRequestCount)

	chainConfig, err := c.reader.processConfigResults(destChain, standardResultsCopy)
	if err != nil {
		return fmt.Errorf("process config results: %w", err)
	}

	// 7. Update chain config cache
	chainCache := c.getOrCreateChainCache(destChain)

	chainCache.chainConfigMu.Lock()
	chainCache.chainConfigData = chainConfig
	chainCache.chainConfigRefresh = time.Now()
	chainCache.chainConfigMu.Unlock()

	// 8. Process source chain configs if we requested any
	if len(filteredSourceChains) > 0 {
		sourceConfigs := c.processSourceChainResults(
			batchResult,
			standardOffRampRequestCount,
			filteredSourceChains,
		)

		// Update source chain config cache with any successful results
		if len(sourceConfigs) > 0 {
			chainCache.sourceChainMu.Lock()

			// Update configs in the map
			for chain, config := range sourceConfigs {
				cachedConfig := staticSourceChainConfigFromSourceChainConfig(config)
				chainCache.staticSourceChainConfigs[chain] = cachedConfig
			}

			// Update the refresh timestamp
			chainCache.sourceChainRefresh = time.Now()

			chainCache.sourceChainMu.Unlock()
		}
	}

	fetchConfigLatency := time.Since(startTime)
	c.lggr.Debugw("Successfully refreshed chain and source configs in single batch",
		"destChain", destChain,
		"standardOffRampRequests", standardOffRampRequestCount,
		"sourceChains", len(filteredSourceChains),
		"fetchConfigLatency", fetchConfigLatency)

	return nil
}

// trackSourceChain adds a source chain to the known source chains map for a destination chain
func (c *configPoller) trackSourceChain(destChain, sourceChain cciptypes.ChainSelector) bool {
	if destChain == sourceChain {
		return false // Don't track destination chain as its own source
	}

	// First check if we have a contract reader for this destination chain
	if _, exists := c.reader.getContractReader(destChain); !exists {
		c.lggr.Debugw("Cannot track source chain - no contract reader for dest chain",
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

// getOrCreateChainCache safely retrieves or creates a cache for a specific chain
func (c *configPoller) getOrCreateChainCache(chainSel cciptypes.ChainSelector) *chainCache {
	c.Lock()
	defer c.Unlock()

	if cache, exists := c.chainCaches[chainSel]; exists {
		return cache
	}

	// verify we have the reader for this chain
	if _, exists := c.reader.getContractReader(chainSel); !exists {
		c.lggr.Errorw("No contract reader for chain", "chain", chainSel)
		return nil
	}

	cache := &chainCache{
		staticSourceChainConfigs: make(map[cciptypes.ChainSelector]StaticSourceChainConfig),
	}
	c.chainCaches[chainSel] = cache
	return cache
}

// Modified GetChainConfig to be non-blocking, returning whatever is in cache
func (c *configPoller) GetChainConfig(
	ctx context.Context,
	chainSel cciptypes.ChainSelector,
) (ChainConfigSnapshot, error) {
	// Check if we have a reader for this chain
	reader, exists := c.reader.getContractReader(chainSel)
	if !exists || reader == nil {
		c.lggr.Errorw("No contract reader for chain", "chain", chainSel)
		return ChainConfigSnapshot{}, fmt.Errorf("no contract reader for chain %d", chainSel)
	}

	chainCache := c.getOrCreateChainCache(chainSel)

	chainCache.chainConfigMu.RLock()
	// Check if we have any data in cache
	if !chainCache.chainConfigRefresh.IsZero() {
		defer chainCache.chainConfigMu.RUnlock()
		c.lggr.Debugw("Returning cached chain config",
			"chain", chainSel,
			"cacheAge", time.Since(chainCache.chainConfigRefresh))
		return chainCache.chainConfigData, nil
	}
	chainCache.chainConfigMu.RUnlock()

	// No cached data yet, must block for initial load
	c.lggr.Debugw("No cached data available, performing initial fetch",
		"chain", chainSel)
	return c.refreshChainConfig(ctx, chainSel)
}

// Modified GetOfframpSourceChainConfigs to track chains and never check for staleness
func (c *configPoller) GetOfframpSourceChainConfigs(
	ctx context.Context,
	destChain cciptypes.ChainSelector,
	sourceChains []cciptypes.ChainSelector,
) (map[cciptypes.ChainSelector]StaticSourceChainConfig, error) {
	// Verify we have a reader for the destination chain
	if _, exists := c.reader.getContractReader(destChain); !exists {
		c.lggr.Errorw("No contract reader for destination chain", "chain", destChain)
		return nil, fmt.Errorf("no contract reader for destination chain %d", destChain)
	}

	// Filter out destination chain from source chains
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

	// Get or create cache for the destination chain
	chainCache := c.getOrCreateChainCache(destChain)
	if chainCache == nil {
		return nil, fmt.Errorf("failed to create cache for chain %d", destChain)
	}

	chainCache.sourceChainMu.RLock()

	// Initialize results map
	cachedSourceConfigs := make(map[cciptypes.ChainSelector]StaticSourceChainConfig)
	var missingChains []cciptypes.ChainSelector

	// Check which chains exist in cache
	for _, chain := range filteredSourceChains {
		config, exists := chainCache.staticSourceChainConfigs[chain]
		if exists {
			cachedSourceConfigs[chain] = config
		} else {
			// This chain isn't in cache yet
			missingChains = append(missingChains, chain)
		}
	}

	// If all chains are in cache, return them immediately
	if len(missingChains) == 0 {
		chainCache.sourceChainMu.RUnlock()
		c.lggr.Debugw("All source chain configs found in cache",
			"destChain", destChain,
			"sourceChains", filteredSourceChains)
		return cachedSourceConfigs, nil
	}

	// First-time fetch for some chains, must block for initial load
	chainCache.sourceChainMu.RUnlock()

	c.lggr.Debugw("Some chains missing from cache, fetching initial data",
		"destChain", destChain,
		"missingChains", missingChains)

	// Get the missing configs
	newCachedConfigs, err := c.refreshSourceChainConfigs(ctx, destChain, missingChains)
	if err != nil {
		return nil, err
	}

	// Merge the new configs with existing cached results
	for chain, config := range newCachedConfigs {
		cachedSourceConfigs[chain] = config
	}

	return cachedSourceConfigs, nil
}

// refreshChainConfig forces a refresh of the chain configuration
func (c *configPoller) refreshChainConfig(
	ctx context.Context,
	chainSel cciptypes.ChainSelector,
) (ChainConfigSnapshot, error) {
	chainCache := c.getOrCreateChainCache(chainSel)

	// Check if context is done and we have cached data (short read lock)
	chainCache.chainConfigMu.RLock()
	hasCachedData := !chainCache.chainConfigRefresh.IsZero()
	cachedData := chainCache.chainConfigData
	chainCache.chainConfigMu.RUnlock()

	if hasCachedData && ctx.Err() != nil {
		c.lggr.Debugw("Context done but returning cached data", "chain", chainSel)
		return cachedData, nil
	}

	// Perform I/O operation WITHOUT holding the lock
	startTime := time.Now()
	newData, err := c.fetchChainConfig(ctx, chainSel)
	fetchConfigLatency := time.Since(startTime)

	if err != nil {
		// Handle error, possibly returning cached data
		chainCache.chainConfigMu.RLock()
		hasCachedData = !chainCache.chainConfigRefresh.IsZero()
		cachedData = chainCache.chainConfigData
		chainCache.chainConfigMu.RUnlock()

		if hasCachedData {
			c.lggr.Warnw("Failed to refresh cache, using old data",
				"chain", chainSel,
				"error", err,
				"lastRefresh", chainCache.chainConfigRefresh,
				"fetchConfigLatency", fetchConfigLatency)
			return cachedData, nil
		}

		c.lggr.Errorw("Failed to refresh cache, no old data available",
			"chain", chainSel,
			"error", err,
			"fetchConfigLatency", fetchConfigLatency)
		return ChainConfigSnapshot{}, fmt.Errorf("failed to refresh cache for chain %d: %w", chainSel, err)
	}

	// Acquire write lock only for updating the cache
	chainCache.chainConfigMu.Lock()
	chainCache.chainConfigData = newData
	chainCache.chainConfigRefresh = time.Now()
	chainCache.chainConfigMu.Unlock()

	c.lggr.Debugw("Successfully refreshed cache",
		"chain", chainSel,
		"fetchConfigLatency", fetchConfigLatency)

	return newData, nil
}

// refreshSourceChainConfigs forces a refresh of source chain configurations
func (c *configPoller) refreshSourceChainConfigs(
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

func (c *configPoller) fetchChainConfig(
	ctx context.Context,
	chainSel cciptypes.ChainSelector) (ChainConfigSnapshot, error) {

	reader, exists := c.reader.getContractReader(chainSel)
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
var _ services.Service = (*configPoller)(nil)
