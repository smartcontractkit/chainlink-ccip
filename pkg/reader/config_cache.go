package reader

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/smartcontractkit/chainlink-ccip/pkg/consts"
	"github.com/smartcontractkit/chainlink-ccip/pkg/contractreader"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
	"github.com/smartcontractkit/chainlink-common/pkg/types"
)

const (
	configCacheRefreshInterval = 30 * time.Second
)

// configCache handles caching of contract configurations with automatic refresh
type configCache struct {
	reader       *ccipChainReader
	cacheMu      sync.RWMutex
	lastUpdateAt time.Time

	// Internal state
	nativeTokenAddress    cciptypes.Bytes
	commitLatestOCRConfig cciptypes.OCRConfigResponse
	execLatestOCRConfig   cciptypes.OCRConfigResponse
	offrampStaticConfig   offRampStaticChainConfig
	offrampDynamicConfig  offRampDynamicChainConfig
	offrampAllChains      selectorsAndConfigs
	onrampDynamicConfig   getOnRampDynamicConfigResponse
	rmnDigestHeader       cciptypes.RMNDigestHeader
	rmnVersionedConfig    versionedConfig
	rmnRemoteAddress      cciptypes.Bytes
	feeQuoterConfig       feeQuoterStaticConfig
}

// newConfigCache creates a new instance of the configuration cache
func newConfigCache(reader *ccipChainReader) *configCache {
	return &configCache{
		reader: reader,
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

	batchResult, err := c.reader.contractReaders[c.reader.destChain].ExtendedBatchGetLatestValues(ctx, requests)
	if err != nil {
		return fmt.Errorf("batch get configs: %w", err)
	}

	if err := c.updateFromResults(batchResult); err != nil {
		return fmt.Errorf("update cache from results: %w", err)
	}

	return nil
}

// prepareBatchRequests creates the batch request for all configurations
func (c *configCache) prepareBatchRequests() contractreader.ExtendedBatchGetLatestValuesRequest {
	var (
		nativeTokenAddress    cciptypes.Bytes
		onrampDynamicConfig   getOnRampDynamicConfigResponse
		commitLatestOCRConfig cciptypes.OCRConfigResponse
		execLatestOCRConfig   cciptypes.OCRConfigResponse
		staticConfig          offRampStaticChainConfig
		dynamicConfig         offRampDynamicChainConfig
		selectorsAndConf      selectorsAndConfigs
		rmnDigestHeader       cciptypes.RMNDigestHeader
		rmnVersionConfig      versionedConfig
		rmnRemoteAddress      []byte
		feeQuoterConfig       feeQuoterStaticConfig
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
		switch contract.Name {
		case consts.ContractNameRouter:
			if len(results) > 0 {
				val, err := results[0].GetResult()
				if err != nil {
					return fmt.Errorf("get router result: %w", err)
				}
				if typed, ok := val.(*cciptypes.Bytes); ok {
					c.nativeTokenAddress = *typed
				}
			}
		case consts.ContractNameOnRamp:
			if len(results) > 0 {
				val, err := results[0].GetResult()
				if err != nil {
					return fmt.Errorf("get onramp result: %w", err)
				}
				if typed, ok := val.(*getOnRampDynamicConfigResponse); ok {
					c.onrampDynamicConfig = *typed
				}
			}
		case consts.ContractNameOffRamp:
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
					if typed, ok := val.(*offRampStaticChainConfig); ok {
						c.offrampStaticConfig = *typed
					}
				case 3:
					if typed, ok := val.(*offRampDynamicChainConfig); ok {
						c.offrampDynamicConfig = *typed
					}
				case 4:
					if typed, ok := val.(*selectorsAndConfigs); ok {
						c.offrampAllChains = *typed
					}
				}
			}
		case consts.ContractNameRMNRemote:
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
					if typed, ok := val.(*versionedConfig); ok {
						c.rmnVersionedConfig = *typed
					}
				}
			}
		case consts.ContractNameRMNProxy:
			if len(results) > 0 {
				val, err := results[0].GetResult()
				if err != nil {
					return fmt.Errorf("get rmn proxy result: %w", err)
				}
				if typed, ok := val.(*cciptypes.Bytes); ok {
					c.rmnRemoteAddress = *typed
				}
			}
		case consts.ContractNameFeeQuoter:
			if len(results) > 0 {
				val, err := results[0].GetResult()
				if err != nil {
					return fmt.Errorf("get fee quoter result: %w", err)
				}
				if typed, ok := val.(*feeQuoterStaticConfig); ok {
					c.feeQuoterConfig = *typed
				}
			}
		}
	}
	return nil
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
func (c *configCache) GetOnRampDynamicConfig(ctx context.Context) (getOnRampDynamicConfigResponse, error) {
	if err := c.refreshIfNeeded(ctx); err != nil {
		return getOnRampDynamicConfigResponse{}, fmt.Errorf("refresh cache: %w", err)
	}

	c.cacheMu.RLock()
	defer c.cacheMu.RUnlock()
	return c.onrampDynamicConfig, nil
}

// GetOffRampStaticConfig returns the cached offramp static config
func (c *configCache) GetOffRampStaticConfig(ctx context.Context) (offRampStaticChainConfig, error) {
	if err := c.refreshIfNeeded(ctx); err != nil {
		return offRampStaticChainConfig{}, fmt.Errorf("refresh cache: %w", err)
	}

	c.cacheMu.RLock()
	defer c.cacheMu.RUnlock()
	return c.offrampStaticConfig, nil
}

// GetOffRampDynamicConfig returns the cached offramp dynamic config
func (c *configCache) GetOffRampDynamicConfig(ctx context.Context) (offRampDynamicChainConfig, error) {
	if err := c.refreshIfNeeded(ctx); err != nil {
		return offRampDynamicChainConfig{}, fmt.Errorf("refresh cache: %w", err)
	}

	c.cacheMu.RLock()
	defer c.cacheMu.RUnlock()
	return c.offrampDynamicConfig, nil
}

// GetOffRampAllChains returns the cached offramp all chains config
func (c *configCache) GetOffRampAllChains(ctx context.Context) (selectorsAndConfigs, error) {
	if err := c.refreshIfNeeded(ctx); err != nil {
		return selectorsAndConfigs{}, fmt.Errorf("refresh cache: %w", err)
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
func (c *configCache) GetRMNVersionedConfig(ctx context.Context) (versionedConfig, error) {
	if err := c.refreshIfNeeded(ctx); err != nil {
		return versionedConfig{}, fmt.Errorf("refresh cache: %w", err)
	}

	c.cacheMu.RLock()
	defer c.cacheMu.RUnlock()
	return c.rmnVersionedConfig, nil
}

// GetFeeQuoterConfig returns the cached fee quoter config
func (c *configCache) GetFeeQuoterConfig(ctx context.Context) (feeQuoterStaticConfig, error) {
	if err := c.refreshIfNeeded(ctx); err != nil {
		return feeQuoterStaticConfig{}, fmt.Errorf("refresh cache: %w", err)
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
