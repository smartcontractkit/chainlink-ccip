package configcache

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/smartcontractkit/chainlink-common/pkg/types"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"

	"github.com/smartcontractkit/chainlink-ccip/pkg/consts"
	"github.com/smartcontractkit/chainlink-ccip/pkg/contractreader"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

const (
	configCacheRefreshInterval = 30 * time.Second
)

// configCacher defines the interface for accessing cached config values
type ConfigCacher interface {
	// OCR Config related methods
	GetOffRampConfigDigest(ctx context.Context, pluginType uint8) ([32]byte, error)

	// Token related methods
	GetNativeTokenAddress(ctx context.Context) (cciptypes.Bytes, error)

	// OnRamp related methods
	GetOnRampDynamicConfig(ctx context.Context) (cciptypes.GetOnRampDynamicConfigResponse, error)

	// OffRamp related methods
	GetOffRampStaticConfig(ctx context.Context) (cciptypes.OffRampStaticChainConfig, error)
	GetOffRampDynamicConfig(ctx context.Context) (cciptypes.OffRampDynamicChainConfig, error)
	GetOffRampAllChains(ctx context.Context) (cciptypes.SelectorsAndConfigs, error)

	// RMN related methods
	GetRMNDigestHeader(ctx context.Context) (cciptypes.RMNDigestHeader, error)
	GetRMNVersionedConfig(ctx context.Context) (cciptypes.VersionedConfigRemote, error)
	GetRMNRemoteAddress(ctx context.Context) (cciptypes.Bytes, error)

	// FeeQuoter related methods
	GetFeeQuoterConfig(ctx context.Context) (cciptypes.FeeQuoterStaticConfig, error)
}

// configCache handles caching of contract configurations with automatic refresh
type configCache struct {
	reader       contractreader.Extended
	lggr         logger.Logger
	cacheMu      sync.RWMutex
	lastUpdateAt time.Time

	// Internal state
	nativeTokenAddress    cciptypes.Bytes
	commitLatestOCRConfig cciptypes.OCRConfigResponse
	execLatestOCRConfig   cciptypes.OCRConfigResponse
	offrampStaticConfig   cciptypes.OffRampStaticChainConfig
	offrampDynamicConfig  cciptypes.OffRampDynamicChainConfig
	offrampAllChains      cciptypes.SelectorsAndConfigs
	onrampDynamicConfig   cciptypes.GetOnRampDynamicConfigResponse
	rmnDigestHeader       cciptypes.RMNDigestHeader
	rmnVersionedConfig    cciptypes.VersionedConfigRemote
	rmnRemoteAddress      cciptypes.Bytes
	feeQuoterConfig       cciptypes.FeeQuoterStaticConfig
}

// NewConfigCache creates a new instance of the configuration cache
func NewConfigCache(reader contractreader.Extended, lggr logger.Logger) ConfigCacher {
	return &configCache{
		reader: reader,
		lggr:   lggr,
	}
}

// refreshIfNeeded refreshes the cache if the refresh interval has elapsed
func (c *configCache) refreshIfNeeded(ctx context.Context) error {
	c.cacheMu.Lock()
	defer c.cacheMu.Unlock()

	if time.Since(c.lastUpdateAt) < configCacheRefreshInterval {
		return nil
	}

	if err := c.refresh(ctx); err != nil {
		return fmt.Errorf("refresh cache: %w", err)
	}

	c.lastUpdateAt = time.Now()
	return nil
}

// refresh fetches all configurations and updates the cache
func (c *configCache) refresh(ctx context.Context) error {
	requests := c.prepareBatchRequests()

	batchResult, err := c.reader.ExtendedBatchGetLatestValuesGraceful(ctx, requests)
	if err != nil {
		return fmt.Errorf("batch get configs: %w", err)
	}

	// Log skipped contracts if any for debugging
	// Clear skipped contract values from cache
	if len(batchResult.SkippedNoBinds) > 0 {
		c.lggr.Infow("some contracts were skipped due to no bindings: %v", batchResult.SkippedNoBinds)
		c.clearSkippedContractValues(batchResult.SkippedNoBinds)
	}

	if err := c.updateFromResults(batchResult.Results); err != nil {
		return fmt.Errorf("update cache from results: %w", err)
	}

	return nil
}

// prepareBatchRequests creates the batch request for all configurations
func (c *configCache) prepareBatchRequests() contractreader.ExtendedBatchGetLatestValuesRequest {
	var (
		nativeTokenAddress    cciptypes.Bytes
		onrampDynamicConfig   cciptypes.GetOnRampDynamicConfigResponse
		commitLatestOCRConfig cciptypes.OCRConfigResponse
		execLatestOCRConfig   cciptypes.OCRConfigResponse
		staticConfig          cciptypes.OffRampStaticChainConfig
		dynamicConfig         cciptypes.OffRampDynamicChainConfig
		selectorsAndConf      cciptypes.SelectorsAndConfigs
		rmnDigestHeader       cciptypes.RMNDigestHeader
		rmnVersionConfig      cciptypes.VersionedConfigRemote
		rmnRemoteAddress      []byte
		feeQuoterConfig       cciptypes.FeeQuoterStaticConfig
	)

	return contractreader.ExtendedBatchGetLatestValuesRequest{
		consts.ContractNameRouter: {{
			ReadName:  consts.MethodNameRouterGetWrappedNative,
			Params:    map[string]any{},
			ReturnVal: &nativeTokenAddress,
		}},
		consts.ContractNameOnRamp: {{
			ReadName:  consts.MethodNameOnRampGetDynamicConfig,
			Params:    map[string]any{},
			ReturnVal: &onrampDynamicConfig,
		}},
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
		consts.ContractNameRMNProxy: {{
			ReadName:  consts.MethodNameGetARM,
			Params:    map[string]any{},
			ReturnVal: &rmnRemoteAddress,
		}},
		consts.ContractNameFeeQuoter: {{
			ReadName:  consts.MethodNameFeeQuoterGetStaticConfig,
			Params:    map[string]any{},
			ReturnVal: &feeQuoterConfig,
		}},
	}
}

// updateFromResults updates the cache with results from the batch request
func (c *configCache) updateFromResults(batchResult types.BatchGetLatestValuesResult) error {
	for contract, results := range batchResult {
		if err := c.handleContractResults(contract, results); err != nil {
			return err
		}
	}
	return nil
}

// handleContractResults processes results for a specific contract
func (c *configCache) handleContractResults(contract types.BoundContract, results []types.BatchReadResult) error {
	switch contract.Name {
	case consts.ContractNameRouter:
		return c.handleRouterResults(results)
	case consts.ContractNameOnRamp:
		return c.handleOnRampResults(results)
	case consts.ContractNameOffRamp:
		return c.handleOffRampResults(results)
	case consts.ContractNameRMNRemote:
		return c.handleRMNRemoteResults(results)
	case consts.ContractNameRMNProxy:
		return c.handleRMNProxyResults(results)
	case consts.ContractNameFeeQuoter:
		return c.handleFeeQuoterResults(results)
	}
	return nil
}

// handleRouterResults processes router-specific results
func (c *configCache) handleRouterResults(results []types.BatchReadResult) error {
	if len(results) > 0 {
		val, err := results[0].GetResult()
		if err != nil {
			return fmt.Errorf("get router result: %w", err)
		}
		if typed, ok := val.(*cciptypes.Bytes); ok {
			c.nativeTokenAddress = *typed
		}
	}
	return nil
}

// handleOnRampResults processes onramp-specific results
func (c *configCache) handleOnRampResults(results []types.BatchReadResult) error {
	if len(results) > 0 {
		val, err := results[0].GetResult()
		if err != nil {
			return fmt.Errorf("get onramp result: %w", err)
		}
		if typed, ok := val.(*cciptypes.GetOnRampDynamicConfigResponse); ok {
			c.onrampDynamicConfig = *typed
		}
	}
	return nil
}

// handleOffRampResults processes offramp-specific results
func (c *configCache) handleOffRampResults(results []types.BatchReadResult) error {
	for i, result := range results {
		val, err := result.GetResult()
		if err != nil {
			return fmt.Errorf("get offramp result %d: %w", i, err)
		}
		switch i {
		case 0:
			if typed, ok := val.(*cciptypes.OCRConfigResponse); ok {
				c.commitLatestOCRConfig = *typed
			}
		case 1:
			if typed, ok := val.(*cciptypes.OCRConfigResponse); ok {
				c.execLatestOCRConfig = *typed
			}
		case 2:
			if typed, ok := val.(*cciptypes.OffRampStaticChainConfig); ok {
				c.offrampStaticConfig = *typed
			}
		case 3:
			if typed, ok := val.(*cciptypes.OffRampDynamicChainConfig); ok {
				c.offrampDynamicConfig = *typed
			}
		case 4:
			if typed, ok := val.(*cciptypes.SelectorsAndConfigs); ok {
				c.offrampAllChains = *typed
			}
		}
	}
	return nil
}

// handleRMNRemoteResults processes RMN remote-specific results
func (c *configCache) handleRMNRemoteResults(results []types.BatchReadResult) error {
	for i, result := range results {
		val, err := result.GetResult()
		if err != nil {
			return fmt.Errorf("get rmn remote result %d: %w", i, err)
		}
		switch i {
		case 0:
			if typed, ok := val.(*cciptypes.RMNDigestHeader); ok {
				c.rmnDigestHeader = *typed
			}
		case 1:
			if typed, ok := val.(*cciptypes.VersionedConfigRemote); ok {
				c.rmnVersionedConfig = *typed
			}
		}
	}
	return nil
}

// handleRMNProxyResults processes RMN proxy-specific results
func (c *configCache) handleRMNProxyResults(results []types.BatchReadResult) error {
	if len(results) > 0 {
		val, err := results[0].GetResult()
		if err != nil {
			return fmt.Errorf("get rmn proxy result: %w", err)
		}
		if typed, ok := val.(*cciptypes.Bytes); ok {
			c.rmnRemoteAddress = *typed
		}
	}
	return nil
}

// handleFeeQuoterResults processes fee quoter-specific results
func (c *configCache) handleFeeQuoterResults(results []types.BatchReadResult) error {
	if len(results) > 0 {
		val, err := results[0].GetResult()
		if err != nil {
			return fmt.Errorf("get fee quoter result: %w", err)
		}
		if typed, ok := val.(*cciptypes.FeeQuoterStaticConfig); ok {
			c.feeQuoterConfig = *typed
		}
	}
	return nil
}

// clearSkippedContractValues resets cache values for contracts that had no bindings
func (c *configCache) clearSkippedContractValues(skippedContracts []string) {
	for _, contractName := range skippedContracts {
		switch contractName {
		case consts.ContractNameRouter:
			c.nativeTokenAddress = cciptypes.Bytes{}
		case consts.ContractNameOnRamp:
			c.onrampDynamicConfig = cciptypes.GetOnRampDynamicConfigResponse{}
		case consts.ContractNameOffRamp:
			c.commitLatestOCRConfig = cciptypes.OCRConfigResponse{}
			c.execLatestOCRConfig = cciptypes.OCRConfigResponse{}
			c.offrampStaticConfig = cciptypes.OffRampStaticChainConfig{}
			c.offrampDynamicConfig = cciptypes.OffRampDynamicChainConfig{}
			c.offrampAllChains = cciptypes.SelectorsAndConfigs{}
		case consts.ContractNameRMNRemote:
			c.rmnDigestHeader = cciptypes.RMNDigestHeader{}
			c.rmnVersionedConfig = cciptypes.VersionedConfigRemote{}
		case consts.ContractNameRMNProxy:
			c.rmnRemoteAddress = cciptypes.Bytes{}
		case consts.ContractNameFeeQuoter:
			c.feeQuoterConfig = cciptypes.FeeQuoterStaticConfig{}
		}
	}
}

func (c *configCache) GetOffRampConfigDigest(ctx context.Context, pluginType uint8) ([32]byte, error) {
	if err := c.refreshIfNeeded(ctx); err != nil {
		return [32]byte{}, fmt.Errorf("refresh cache: %w", err)
	}

	c.cacheMu.RLock()
	defer c.cacheMu.RUnlock()

	if pluginType == consts.PluginTypeCommit {
		return c.commitLatestOCRConfig.OCRConfig.ConfigInfo.ConfigDigest, nil
	}
	return c.execLatestOCRConfig.OCRConfig.ConfigInfo.ConfigDigest, nil
}

// GetNativeTokenAddress returns the cached native token address
func (c *configCache) GetNativeTokenAddress(ctx context.Context) (cciptypes.Bytes, error) {
	if err := c.refreshIfNeeded(ctx); err != nil {
		return cciptypes.Bytes{}, fmt.Errorf("refresh cache: %w", err)
	}

	c.cacheMu.RLock()
	defer c.cacheMu.RUnlock()
	return c.nativeTokenAddress, nil
}

// GetOnRampDynamicConfig returns the cached onramp dynamic config
func (c *configCache) GetOnRampDynamicConfig(ctx context.Context) (cciptypes.GetOnRampDynamicConfigResponse, error) {
	if err := c.refreshIfNeeded(ctx); err != nil {
		return cciptypes.GetOnRampDynamicConfigResponse{}, fmt.Errorf("refresh cache: %w", err)
	}

	c.cacheMu.RLock()
	defer c.cacheMu.RUnlock()
	return c.onrampDynamicConfig, nil
}

// GetOffRampStaticConfig returns the cached offramp static config
func (c *configCache) GetOffRampStaticConfig(ctx context.Context) (cciptypes.OffRampStaticChainConfig, error) {
	if err := c.refreshIfNeeded(ctx); err != nil {
		return cciptypes.OffRampStaticChainConfig{}, fmt.Errorf("refresh cache: %w", err)
	}

	c.cacheMu.RLock()
	defer c.cacheMu.RUnlock()
	return c.offrampStaticConfig, nil
}

// GetOffRampDynamicConfig returns the cached offramp dynamic config
func (c *configCache) GetOffRampDynamicConfig(ctx context.Context) (cciptypes.OffRampDynamicChainConfig, error) {
	if err := c.refreshIfNeeded(ctx); err != nil {
		return cciptypes.OffRampDynamicChainConfig{}, fmt.Errorf("refresh cache: %w", err)
	}

	c.cacheMu.RLock()
	defer c.cacheMu.RUnlock()
	return c.offrampDynamicConfig, nil
}

// GetOffRampAllChains returns the cached offramp all chains config
func (c *configCache) GetOffRampAllChains(ctx context.Context) (cciptypes.SelectorsAndConfigs, error) {
	if err := c.refreshIfNeeded(ctx); err != nil {
		return cciptypes.SelectorsAndConfigs{}, fmt.Errorf("refresh cache: %w", err)
	}

	c.cacheMu.RLock()
	defer c.cacheMu.RUnlock()
	return c.offrampAllChains, nil
}

// GetRMNDigestHeader returns the cached RMN digest header
func (c *configCache) GetRMNDigestHeader(ctx context.Context) (cciptypes.RMNDigestHeader, error) {
	if err := c.refreshIfNeeded(ctx); err != nil {
		return cciptypes.RMNDigestHeader{}, fmt.Errorf("refresh cache: %w", err)
	}

	c.cacheMu.RLock()
	defer c.cacheMu.RUnlock()
	return c.rmnDigestHeader, nil
}

// GetRMNVersionedConfig returns the cached RMN versioned config
func (c *configCache) GetRMNVersionedConfig(ctx context.Context) (cciptypes.VersionedConfigRemote, error) {
	if err := c.refreshIfNeeded(ctx); err != nil {
		return cciptypes.VersionedConfigRemote{}, fmt.Errorf("refresh cache: %w", err)
	}

	c.cacheMu.RLock()
	defer c.cacheMu.RUnlock()
	return c.rmnVersionedConfig, nil
}

// GetFeeQuoterConfig returns the cached fee quoter config
func (c *configCache) GetFeeQuoterConfig(ctx context.Context) (cciptypes.FeeQuoterStaticConfig, error) {
	if err := c.refreshIfNeeded(ctx); err != nil {
		return cciptypes.FeeQuoterStaticConfig{}, fmt.Errorf("refresh cache: %w", err)
	}

	c.cacheMu.RLock()
	defer c.cacheMu.RUnlock()
	return c.feeQuoterConfig, nil
}

// GetRMNRemoteAddress returns the cached RMN remote address
func (c *configCache) GetRMNRemoteAddress(ctx context.Context) (cciptypes.Bytes, error) {
	if err := c.refreshIfNeeded(ctx); err != nil {
		return cciptypes.Bytes{}, fmt.Errorf("refresh cache: %w", err)
	}

	c.cacheMu.RLock()
	defer c.cacheMu.RUnlock()
	return c.rmnRemoteAddress, nil
}

var _ ConfigCacher = (*configCache)(nil)
