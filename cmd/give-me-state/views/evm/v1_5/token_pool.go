package v1_5

import (
	"call-orchestrator-demo/views"
	"call-orchestrator-demo/views/evm/common"
	"encoding/hex"
	"math/big"
)

// Function selectors for TokenPool v1.5
var (
	// getToken() returns (address)
	selectorGetToken = common.HexToSelector("21df0da7")
	// getSupportedChains() returns (uint64[])
	selectorGetSupportedChains = common.HexToSelector("ed231d68")
	// getRebalancer() returns (address) - only for LockReleaseTokenPool
	selectorGetRebalancer = common.HexToSelector("ebcd4ac5")
	// getAllowList() returns (address[])
	selectorGetAllowList = common.HexToSelector("e0e03cae")
	// getAllowListEnabled() returns (bool)
	selectorGetAllowListEnabled = common.HexToSelector("72b0d90c")
	// getRemotePool(uint64) returns (bytes)
	selectorGetRemotePool = common.HexToSelector("f2e3d80c")
	// getRemoteToken(uint64) returns (bytes)
	selectorGetRemoteToken = common.HexToSelector("62af5f86")
	// getCurrentInboundRateLimiterState(uint64) returns (RateLimiter.TokenBucket)
	selectorGetCurrentInboundRateLimiterState = common.HexToSelector("7ee2fb5f")
	// getCurrentOutboundRateLimiterState(uint64) returns (RateLimiter.TokenBucket)
	selectorGetCurrentOutboundRateLimiterState = common.HexToSelector("e5d77e29")
	// getRmnProxy() returns (address)
	selectorGetRmnProxy = common.HexToSelector("dc0bd971")
	// getRouter() returns (address)
	selectorGetRouter = common.HexToSelector("b0f479a1")
)

// RateLimiterConfig represents the rate limiter configuration
type RateLimiterConfig struct {
	IsEnabled bool   `json:"isEnabled"`
	Capacity  string `json:"capacity"`
	Rate      string `json:"rate"`
}

// RemoteChainConfig represents the configuration for a remote chain
type RemoteChainConfig struct {
	RemoteTokenAddress        string            `json:"remoteTokenAddress"`
	RemotePoolAddresses       []string          `json:"remotePoolAddresses"`
	InboundRateLimiterConfig  RateLimiterConfig `json:"inboundRateLimiterConfig"`
	OutboundRateLimiterConfig RateLimiterConfig `json:"outboundRateLimiterConfig"`
}

// ViewBurnMintTokenPool generates a view of a BurnMintTokenPool contract (v1.5.0).
func ViewBurnMintTokenPool(ctx *views.ViewContext) (map[string]any, error) {
	return viewTokenPoolCommon(ctx, "1.5.0", false)
}

// ViewLockReleaseTokenPool generates a view of a LockReleaseTokenPool contract (v1.5.0).
func ViewLockReleaseTokenPool(ctx *views.ViewContext) (map[string]any, error) {
	return viewTokenPoolCommon(ctx, "1.5.0", true)
}

// ViewBurnMintTokenPoolAndProxy generates a view of a BurnMintTokenPoolAndProxy contract (v1.5.0).
func ViewBurnMintTokenPoolAndProxy(ctx *views.ViewContext) (map[string]any, error) {
	return viewTokenPoolCommon(ctx, "1.5.0", false)
}

// ViewLockReleaseTokenPoolAndProxy generates a view of a LockReleaseTokenPoolAndProxy contract (v1.5.0).
func ViewLockReleaseTokenPoolAndProxy(ctx *views.ViewContext) (map[string]any, error) {
	return viewTokenPoolCommon(ctx, "1.5.0", true)
}

// viewTokenPoolCommon is the common implementation for all token pool views.
func viewTokenPoolCommon(ctx *views.ViewContext, version string, hasRebalancer bool) (map[string]any, error) {
	result := make(map[string]any)

	result["address"] = ctx.AddressHex
	result["chainSelector"] = ctx.ChainSelector
	result["version"] = version

	owner, err := common.GetOwner(ctx)
	if err != nil {
		result["owner_error"] = err.Error()
	} else {
		result["owner"] = owner
	}

	typeAndVersion, err := common.GetTypeAndVersion(ctx)
	if err != nil {
		result["typeAndVersion_error"] = err.Error()
	} else {
		result["typeAndVersion"] = typeAndVersion
	}

	token, err := getTokenPoolToken(ctx)
	if err != nil {
		result["token_error"] = err.Error()
	} else {
		result["token"] = token
	}

	// Get RMN Proxy
	rmnProxy, err := getRmnProxyAddress(ctx)
	if err != nil {
		result["rmnProxy_error"] = err.Error()
	} else {
		result["rmnProxy"] = rmnProxy
	}

	// Get Router
	router, err := getTokenPoolRouter(ctx)
	if err != nil {
		result["router_error"] = err.Error()
	} else {
		result["router"] = router
	}

	// Get allow list
	allowList, err := getTokenPoolAllowList(ctx)
	if err != nil {
		result["allowList_error"] = err.Error()
	} else {
		result["allowList"] = allowList
	}

	// Get allow list enabled
	allowListEnabled, err := getTokenPoolAllowListEnabled(ctx)
	if err != nil {
		result["allowListEnabled_error"] = err.Error()
	} else {
		result["allowListEnabled"] = allowListEnabled
	}

	// Get supported chains
	supportedChains, err := getTokenPoolSupportedChains(ctx)
	if err != nil {
		result["supportedChains_error"] = err.Error()
	} else {
		result["supportedChains"] = supportedChains

		// For each supported chain, get the remote chain config
		remoteChainConfigs := make(map[uint64]RemoteChainConfig)
		for _, chainSelector := range supportedChains {
			config, err := getRemoteChainConfig(ctx, chainSelector)
			if err != nil {
				// Store error but continue with other chains
				result["remoteChainConfig_"+formatChainSelector(chainSelector)+"_error"] = err.Error()
			} else {
				remoteChainConfigs[chainSelector] = config
			}
		}
		result["remoteChainConfigs"] = remoteChainConfigs
	}

	// Get rebalancer if this is a LockRelease pool
	if hasRebalancer {
		rebalancer, err := getTokenPoolRebalancer(ctx)
		if err != nil {
			result["rebalancer_error"] = err.Error()
		} else {
			result["rebalancer"] = rebalancer
		}
	}

	return result, nil
}

// formatChainSelector formats a chain selector for use in error keys.
func formatChainSelector(selector uint64) string {
	return new(big.Int).SetUint64(selector).String()
}

// getTokenPoolToken fetches the token address.
func getTokenPoolToken(ctx *views.ViewContext) (string, error) {
	data, err := common.ExecuteCall(ctx, selectorGetToken)
	if err != nil {
		return "", err
	}
	return common.DecodeAddress(data)
}

// getRmnProxyAddress fetches the RMN proxy address.
func getRmnProxyAddress(ctx *views.ViewContext) (string, error) {
	data, err := common.ExecuteCall(ctx, selectorGetRmnProxy)
	if err != nil {
		return "", err
	}
	return common.DecodeAddress(data)
}

// getTokenPoolRouter fetches the router address.
func getTokenPoolRouter(ctx *views.ViewContext) (string, error) {
	data, err := common.ExecuteCall(ctx, selectorGetRouter)
	if err != nil {
		return "", err
	}
	return common.DecodeAddress(data)
}

// getTokenPoolSupportedChains fetches the supported chains.
func getTokenPoolSupportedChains(ctx *views.ViewContext) ([]uint64, error) {
	data, err := common.ExecuteCall(ctx, selectorGetSupportedChains)
	if err != nil {
		return nil, err
	}
	if len(data) < 64 {
		return []uint64{}, nil
	}
	length := common.DecodeUint64FromBytes(data[32:64])
	if length == 0 {
		return []uint64{}, nil
	}
	chains := make([]uint64, length)
	for i := uint64(0); i < length; i++ {
		offset := 64 + i*32
		if offset+32 > uint64(len(data)) {
			break
		}
		chains[i] = common.DecodeUint64FromBytes(data[offset : offset+32])
	}
	return chains, nil
}

// getTokenPoolRebalancer fetches the rebalancer address.
func getTokenPoolRebalancer(ctx *views.ViewContext) (string, error) {
	data, err := common.ExecuteCall(ctx, selectorGetRebalancer)
	if err != nil {
		return "", err
	}
	return common.DecodeAddress(data)
}

// getTokenPoolAllowList fetches the allow list addresses.
func getTokenPoolAllowList(ctx *views.ViewContext) ([]string, error) {
	data, err := common.ExecuteCall(ctx, selectorGetAllowList)
	if err != nil {
		return nil, err
	}
	if len(data) < 64 {
		return []string{}, nil
	}
	length := common.DecodeUint64FromBytes(data[32:64])
	if length == 0 {
		return []string{}, nil
	}
	addresses := make([]string, 0, length)
	for i := uint64(0); i < length; i++ {
		offset := 64 + i*32
		if offset+32 > uint64(len(data)) {
			break
		}
		addr := "0x" + hex.EncodeToString(data[offset+12:offset+32])
		addresses = append(addresses, addr)
	}
	return addresses, nil
}

// getTokenPoolAllowListEnabled fetches whether the allow list is enabled.
func getTokenPoolAllowListEnabled(ctx *views.ViewContext) (bool, error) {
	data, err := common.ExecuteCall(ctx, selectorGetAllowListEnabled)
	if err != nil {
		return false, err
	}
	return common.DecodeBool(data)
}

// getRemoteChainConfig fetches the configuration for a remote chain.
func getRemoteChainConfig(ctx *views.ViewContext, remoteChainSelector uint64) (RemoteChainConfig, error) {
	config := RemoteChainConfig{
		RemotePoolAddresses: []string{},
	}

	// Encode the chain selector as argument
	chainSelectorArg := common.EncodeUint64(remoteChainSelector)

	// Get remote pool
	remotePoolData, err := common.ExecuteCall(ctx, selectorGetRemotePool, chainSelectorArg)
	if err == nil {
		remotePool := decodeRemoteAddress(remotePoolData, remoteChainSelector)
		if remotePool != "" {
			config.RemotePoolAddresses = []string{remotePool}
		}
	}

	// Get remote token
	remoteTokenData, err := common.ExecuteCall(ctx, selectorGetRemoteToken, chainSelectorArg)
	if err == nil {
		config.RemoteTokenAddress = decodeRemoteAddress(remoteTokenData, remoteChainSelector)
	}

	// Get inbound rate limiter state
	inboundData, err := common.ExecuteCall(ctx, selectorGetCurrentInboundRateLimiterState, chainSelectorArg)
	if err == nil {
		config.InboundRateLimiterConfig = decodeRateLimiterConfig(inboundData)
	}

	// Get outbound rate limiter state
	outboundData, err := common.ExecuteCall(ctx, selectorGetCurrentOutboundRateLimiterState, chainSelectorArg)
	if err == nil {
		config.OutboundRateLimiterConfig = decodeRateLimiterConfig(outboundData)
	}

	return config, nil
}

// decodeRemoteAddress decodes a remote address from bytes data.
// The data is ABI-encoded as bytes, which contains the actual address.
func decodeRemoteAddress(data []byte, chainSelector uint64) string {
	// The return is bytes, so:
	// - First 32 bytes: offset to bytes data
	// - At offset: length of bytes
	// - Following: actual bytes data

	if len(data) < 64 {
		return ""
	}

	// Get offset
	offset := common.DecodeUint64FromBytes(data[0:32])
	if offset+32 > uint64(len(data)) {
		return ""
	}

	// Get length
	length := common.DecodeUint64FromBytes(data[offset : offset+32])
	if length == 0 {
		return ""
	}

	// Get the bytes data
	bytesStart := offset + 32
	if bytesStart+length > uint64(len(data)) {
		return ""
	}

	addressBytes := data[bytesStart : bytesStart+length]

	// For EVM chains (20-byte addresses), decode as hex
	if length == 20 {
		return "0x" + hex.EncodeToString(addressBytes)
	}

	// For other chain types, just return as hex
	return "0x" + hex.EncodeToString(addressBytes)
}

// decodeRateLimiterConfig decodes a RateLimiter.TokenBucket from call result.
// TokenBucket struct: (uint128 tokens, uint32 lastUpdated, bool isEnabled, uint128 capacity, uint128 rate)
func decodeRateLimiterConfig(data []byte) RateLimiterConfig {
	if len(data) < 160 {
		return RateLimiterConfig{}
	}

	// The struct has:
	// - uint128 tokens (32 bytes, but only lower 16 used)
	// - uint32 lastUpdated (32 bytes)
	// - bool isEnabled (32 bytes)
	// - uint128 capacity (32 bytes)
	// - uint128 rate (32 bytes)

	isEnabled := data[64+31] != 0
	capacity := new(big.Int).SetBytes(data[96:128])
	rate := new(big.Int).SetBytes(data[128:160])

	return RateLimiterConfig{
		IsEnabled: isEnabled,
		Capacity:  capacity.String(),
		Rate:      rate.String(),
	}
}
