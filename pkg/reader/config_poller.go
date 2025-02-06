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
}

// configPoller handles caching of chain configuration data for multiple chains.
// It is used by the ccipChainReader to store and retrieve configuration data,
// avoiding unnecessary contract calls and improving performance.
// configPoller handles caching of chain configuration data for multiple chains
type configPoller struct {
	sync.RWMutex
	chainCaches   map[cciptypes.ChainSelector]*chainCache
	refreshPeriod time.Duration
	readers       map[cciptypes.ChainSelector]contractreader.Extended
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
	readers map[cciptypes.ChainSelector]contractreader.Extended,
	refreshPeriod time.Duration,
) *configPoller {
	return &configPoller{
		chainCaches:   make(map[cciptypes.ChainSelector]*chainCache),
		refreshPeriod: refreshPeriod,
		readers:       readers,
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
	if _, exists := c.readers[chainSel]; !exists {
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
	reader, exists := c.readers[chainSel]
	if !exists || reader == nil {
		c.lggr.Errorw("No contract reader for chain", "chain", chainSel)
		return ChainConfigSnapshot{}, fmt.Errorf("no contract reader for chain %d", chainSel)
	}

	chainCache := c.getOrCreateChainCache(chainSel)

	chainCache.RLock()
	timeSinceLastRefresh := time.Since(chainCache.lastRefresh)
	if timeSinceLastRefresh < c.refreshPeriod {
		defer chainCache.RUnlock()
		c.lggr.Infow("Cache hit",
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
		c.lggr.Infow("Cache was refreshed by another goroutine",
			"chain", chainSel,
			"timeSinceLastRefresh", timeSinceLastRefresh)
		return chainCache.data, nil
	}

	startTime := time.Now()
	newData, err := c.fetchChainConfig(ctx, chainSel)
	refreshDuration := time.Since(startTime)

	if err != nil {
		if !chainCache.lastRefresh.IsZero() {
			c.lggr.Warnw("Failed to refresh cache, using old data",
				"chain", chainSel,
				"error", err,
				"lastRefresh", chainCache.lastRefresh,
				"refreshDuration", refreshDuration)
			return chainCache.data, nil
		}
		c.lggr.Errorw("Failed to refresh cache, no old data available",
			"chain", chainSel,
			"error", err,
			"refreshDuration", refreshDuration)
		return ChainConfigSnapshot{}, fmt.Errorf("failed to refresh cache for chain %d: %w", chainSel, err)
	}

	chainCache.data = newData
	chainCache.lastRefresh = time.Now()

	c.lggr.Infow("Successfully refreshed cache",
		"chain", chainSel,
		"refreshDuration", refreshDuration)

	return newData, nil
}

// prepareBatchRequests creates the batch request for all configurations
func (c *configPoller) prepareBatchRequests() contractreader.ExtendedBatchGetLatestValuesRequest {
	var (
		commitLatestOCRConfig OCRConfigResponse
		execLatestOCRConfig   OCRConfigResponse
		staticConfig          offRampStaticChainConfig
		dynamicConfig         offRampDynamicChainConfig
		selectorsAndConf      selectorsAndConfigs
		rmnRemoteAddress      []byte
		rmnDigestHeader       rmnDigestHeader
		rmnVersionConfig      versionedConfig
		feeQuoterConfig       feeQuoterStaticConfig
	)

	return contractreader.ExtendedBatchGetLatestValuesRequest{
		consts.ContractNameOffRamp: {
			{
				ReadName: consts.MethodNameOffRampLatestConfigDetails,
				Params: map[string]any{
					"ocrPluginType": consts.PluginTypeCommit,
				},
				ReturnVal: &commitLatestOCRConfig,
			},
			{
				ReadName: consts.MethodNameOffRampLatestConfigDetails,
				Params: map[string]any{
					"ocrPluginType": consts.PluginTypeExecute,
				},
				ReturnVal: &execLatestOCRConfig,
			},
			{
				ReadName:  consts.MethodNameOffRampGetStaticConfig,
				Params:    map[string]any{},
				ReturnVal: &staticConfig,
			},
			{
				ReadName:  consts.MethodNameOffRampGetDynamicConfig,
				Params:    map[string]any{},
				ReturnVal: &dynamicConfig,
			},
			{
				ReadName:  consts.MethodNameOffRampGetAllSourceChainConfigs,
				Params:    map[string]any{},
				ReturnVal: &selectorsAndConf,
			},
		},
		consts.ContractNameRMNProxy: {{
			ReadName:  consts.MethodNameGetARM,
			Params:    map[string]any{},
			ReturnVal: &rmnRemoteAddress,
		}},
		consts.ContractNameRMNRemote: {
			{
				ReadName:  consts.MethodNameGetReportDigestHeader,
				Params:    map[string]any{},
				ReturnVal: &rmnDigestHeader,
			},
			{
				ReadName:  consts.MethodNameGetVersionedConfig,
				Params:    map[string]any{},
				ReturnVal: &rmnVersionConfig,
			},
		},
		consts.ContractNameFeeQuoter: {{
			ReadName:  consts.MethodNameFeeQuoterGetStaticConfig,
			Params:    map[string]any{},
			ReturnVal: &feeQuoterConfig,
		}},
	}
}

// fetchChainConfig fetches the latest configuration for a specific chain
func (c *configPoller) fetchChainConfig(
	ctx context.Context,
	chainSel cciptypes.ChainSelector) (ChainConfigSnapshot, error) {
	reader, exists := c.readers[chainSel]
	if !exists {
		return ChainConfigSnapshot{}, fmt.Errorf("no contract reader for chain %d", chainSel)
	}

	requests := c.prepareBatchRequests()
	batchResult, skipped, err := reader.ExtendedBatchGetLatestValues(ctx, requests, true)
	if err != nil {
		return ChainConfigSnapshot{}, fmt.Errorf("batch get latest values for chain %d: %w", chainSel, err)
	}

	if len(skipped) > 0 {
		c.lggr.Infow("some contracts were skipped due to no bindings",
			"chain", chainSel,
			"contracts", skipped)
	}

	return c.updateFromResults(batchResult)
}

func (c *configPoller) updateFromResults(batchResult types.BatchGetLatestValuesResult) (ChainConfigSnapshot, error) {
	config := ChainConfigSnapshot{}

	for contract, results := range batchResult {
		var err error
		switch contract.Name {
		case consts.ContractNameOffRamp:
			config.Offramp, err = c.processOfframpResults(results)
			if err != nil {
				return ChainConfigSnapshot{}, fmt.Errorf("process offramp results: %w", err)
			}

		case consts.ContractNameRMNProxy:
			config.RMNProxy, err = c.processRMNProxyResults(results)
			if err != nil {
				return ChainConfigSnapshot{}, fmt.Errorf("process RMN proxy results: %w", err)
			}

		case consts.ContractNameRMNRemote:
			config.RMNRemote, err = c.processRMNRemoteResults(results)
			if err != nil {
				return ChainConfigSnapshot{}, fmt.Errorf("process RMN remote results: %w", err)
			}

		case consts.ContractNameFeeQuoter:
			config.FeeQuoter, err = c.processFeeQuoterResults(results)
			if err != nil {
				return ChainConfigSnapshot{}, fmt.Errorf("process fee quoter results: %w", err)
			}

		default:
			c.lggr.Warnw("Unhandled contract in batch results", "contract", contract.Name)
		}
	}

	return config, nil
}

// resultProcessor defines a function type for processing individual results
type resultProcessor func(interface{}) error

func (c *configPoller) processOfframpResults(
	results []types.BatchReadResult) (OfframpConfig, error) {

	if len(results) != 5 {
		return OfframpConfig{}, fmt.Errorf("expected 5 offramp results, got %d", len(results))
	}

	config := OfframpConfig{}

	// Define processors for each expected result
	processors := []resultProcessor{
		// CommitLatestOCRConfig
		func(val interface{}) error {
			typed, ok := val.(*OCRConfigResponse)
			if !ok {
				return fmt.Errorf("invalid type for CommitLatestOCRConfig: %T", val)
			}
			config.CommitLatestOCRConfig = *typed
			return nil
		},
		// ExecLatestOCRConfig
		func(val interface{}) error {
			typed, ok := val.(*OCRConfigResponse)
			if !ok {
				return fmt.Errorf("invalid type for ExecLatestOCRConfig: %T", val)
			}
			config.ExecLatestOCRConfig = *typed
			return nil
		},
		// StaticConfig
		func(val interface{}) error {
			typed, ok := val.(*offRampStaticChainConfig)
			if !ok {
				return fmt.Errorf("invalid type for StaticConfig: %T", val)
			}
			config.StaticConfig = *typed
			return nil
		},
		// DynamicConfig
		func(val interface{}) error {
			typed, ok := val.(*offRampDynamicChainConfig)
			if !ok {
				return fmt.Errorf("invalid type for DynamicConfig: %T", val)
			}
			config.DynamicConfig = *typed
			return nil
		},
		// SelectorsAndConf
		func(val interface{}) error {
			typed, ok := val.(*selectorsAndConfigs)
			if !ok {
				return fmt.Errorf("invalid type for SelectorsAndConf: %T", val)
			}
			config.SelectorsAndConf = *typed
			return nil
		},
	}

	// Process each result with its corresponding processor
	for i, result := range results {
		val, err := result.GetResult()
		if err != nil {
			return OfframpConfig{}, fmt.Errorf("get offramp result %d: %w", i, err)
		}

		if err := processors[i](val); err != nil {
			return OfframpConfig{}, fmt.Errorf("process result %d: %w", i, err)
		}
	}

	return config, nil
}

func (c *configPoller) processRMNProxyResults(results []types.BatchReadResult) (RMNProxyConfig, error) {
	if len(results) != 1 {
		return RMNProxyConfig{}, fmt.Errorf("expected 1 RMN proxy result, got %d", len(results))
	}

	val, err := results[0].GetResult()
	if err != nil {
		return RMNProxyConfig{}, fmt.Errorf("get RMN proxy result: %w", err)
	}

	if bytes, ok := val.(*[]byte); ok {
		return RMNProxyConfig{
			RemoteAddress: *bytes,
		}, nil
	}

	return RMNProxyConfig{}, fmt.Errorf("invalid type for RMN proxy remote address: %T", val)
}

func (c *configPoller) processRMNRemoteResults(results []types.BatchReadResult) (RMNRemoteConfig, error) {
	config := RMNRemoteConfig{}

	if len(results) != 2 {
		return RMNRemoteConfig{}, fmt.Errorf("expected 2 RMN remote results, got %d", len(results))
	}

	// Process DigestHeader
	val, err := results[0].GetResult()
	if err != nil {
		return RMNRemoteConfig{}, fmt.Errorf("get RMN remote digest header result: %w", err)
	}

	typed, ok := val.(*rmnDigestHeader)
	if !ok {
		return RMNRemoteConfig{}, fmt.Errorf("invalid type for RMN remote digest header: %T", val)
	}
	config.DigestHeader = *typed

	// Process VersionedConfig
	val, err = results[1].GetResult()
	if err != nil {
		return RMNRemoteConfig{}, fmt.Errorf("get RMN remote versioned config result: %w", err)
	}

	vconf, ok := val.(*versionedConfig)
	if !ok {
		return RMNRemoteConfig{}, fmt.Errorf("invalid type for RMN remote versioned config: %T", val)
	}
	config.VersionedConfig = *vconf

	return config, nil
}

func (c *configPoller) processFeeQuoterResults(results []types.BatchReadResult) (FeeQuoterConfig, error) {
	if len(results) != 1 {
		return FeeQuoterConfig{}, fmt.Errorf("expected 1 fee quoter result, got %d", len(results))
	}

	val, err := results[0].GetResult()
	if err != nil {
		return FeeQuoterConfig{}, fmt.Errorf("get fee quoter result: %w", err)
	}

	if typed, ok := val.(*feeQuoterStaticConfig); ok {
		return FeeQuoterConfig{
			StaticConfig: *typed,
		}, nil
	}

	return FeeQuoterConfig{}, fmt.Errorf("invalid type for fee quoter static config: %T", val)
}

// Ensure configCache implements ConfigCache
var _ ConfigPoller = (*configPoller)(nil)
