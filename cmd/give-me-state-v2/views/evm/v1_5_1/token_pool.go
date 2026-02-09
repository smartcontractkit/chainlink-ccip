package v1_5_1

import (
	"give-me-state-v2/views"
	"give-me-state-v2/views/evm/common"
	"sync"
)

// Function selectors for token pool methods
var (
	// getToken() returns (address)
	selectorGetToken = common.HexToSelector("21df0da7")
	// getSupportedChains() returns (uint64[])
	selectorGetSupportedChains = common.HexToSelector("c4bffe2b")
	// getRebalancer() returns (address) - only for LockReleaseTokenPool
	selectorGetRebalancer = common.HexToSelector("432a6ba3")
	// getAllowList() returns (address[])
	selectorGetAllowList = common.HexToSelector("a7cd63b7")
	// getAllowListEnabled() returns (bool)
	selectorGetAllowListEnabled = common.HexToSelector("e0351e13")
	// getRemotePools(uint64) returns (bytes[])
	selectorGetRemotePools = common.HexToSelector("a42a7b8b")
	// getRemoteToken(uint64) returns (bytes)
	selectorGetRemoteToken = common.HexToSelector("b7946580")
	// getCurrentInboundRateLimiterState(uint64) returns (RateLimiter.TokenBucket)
	selectorGetInboundRateLimiter = common.HexToSelector("af58d59f")
	// getCurrentOutboundRateLimiterState(uint64) returns (RateLimiter.TokenBucket)
	selectorGetOutboundRateLimiter = common.HexToSelector("c75eea9c")
)

// getTokenPoolToken fetches the token address.
func getTokenPoolToken(ctx *views.ViewContext) (string, error) {
	data, err := common.ExecuteCall(ctx, selectorGetToken)
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

	// Dynamic array: offset (32) + length (32) + addresses
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
		addr, _ := common.DecodeAddress(data[offset : offset+32])
		addresses = append(addresses, addr)
	}

	return addresses, nil
}

// getTokenPoolAllowListEnabled fetches whether allow list is enabled.
func getTokenPoolAllowListEnabled(ctx *views.ViewContext) (bool, error) {
	data, err := common.ExecuteCall(ctx, selectorGetAllowListEnabled)
	if err != nil {
		return false, err
	}
	return common.DecodeBool(data)
}

// getRemoteChainConfigs fetches remote chain configurations concurrently.
func getRemoteChainConfigs(ctx *views.ViewContext, supportedChains []uint64) map[string]any {
	if len(supportedChains) == 0 {
		return map[string]any{}
	}

	result := make(map[string]any)
	var mu sync.Mutex
	var wg sync.WaitGroup

	for _, chainSel := range supportedChains {
		wg.Add(1)
		go func(cs uint64) {
			defer wg.Done()

			config := make(map[string]any)

			// Get remote token
			remoteToken, err := getRemoteTokenForChain(ctx, cs)
			if err == nil && remoteToken != "" {
				config["remoteTokenAddress"] = remoteToken
			}

			// Get remote pool(s)
			remotePools, err := getRemotePoolsForChain(ctx, cs)
			if err == nil && len(remotePools) > 0 {
				config["remotePoolAddresses"] = remotePools
			}

			// Get inbound rate limiter
			inbound, err := getInboundRateLimiter(ctx, cs)
			if err == nil {
				config["inboundRateLimiterConfig"] = inbound
			}

			// Get outbound rate limiter
			outbound, err := getOutboundRateLimiter(ctx, cs)
			if err == nil {
				config["outboundRateLimiterConfig"] = outbound
			}

			// Only add if we have any data
			if len(config) > 0 {
				mu.Lock()
				result[views.Uint64ToString(cs)] = config
				mu.Unlock()
			}
		}(chainSel)
	}

	wg.Wait()
	return result
}

// getRemoteTokenForChain fetches the remote token address for a chain.
func getRemoteTokenForChain(ctx *views.ViewContext, chainSel uint64) (string, error) {
	data, err := common.ExecuteCall(ctx, selectorGetRemoteToken, common.EncodeUint64(chainSel))
	if err != nil {
		return "", err
	}

	// Returns bytes - decode based on length
	if len(data) < 64 {
		return "", nil
	}

	// Dynamic bytes: offset (32) + length (32) + data
	bytesLen := common.DecodeUint64FromBytes(data[32:64])
	if bytesLen == 0 {
		return "", nil
	}

	if bytesLen == 20 {
		// EVM address
		addr, _ := common.DecodeAddress(data[64:96])
		return addr, nil
	} else if bytesLen == 32 {
		// Solana or other 32-byte address
		return views.BytesToHex(data[64:96]), nil
	}

	// Unknown format, return as hex
	return views.BytesToHex(data[64 : 64+bytesLen]), nil
}

// getRemotePoolsForChain fetches the remote pool addresses for a chain.
func getRemotePoolsForChain(ctx *views.ViewContext, chainSel uint64) ([]string, error) {
	data, err := common.ExecuteCall(ctx, selectorGetRemotePools, common.EncodeUint64(chainSel))
	if err != nil {
		return nil, err
	}

	// Returns bytes[] - dynamic array of dynamic bytes
	if len(data) < 64 {
		return []string{}, nil
	}

	// Read array length
	arrayLen := common.DecodeUint64FromBytes(data[32:64])
	if arrayLen == 0 {
		return []string{}, nil
	}

	pools := make([]string, 0, arrayLen)
	baseOffset := uint64(64) // After the initial offset and length

	for i := uint64(0); i < arrayLen; i++ {
		if baseOffset+i*32+32 > uint64(len(data)) {
			break
		}
		elementOffset := common.DecodeUint64FromBytes(data[baseOffset+i*32 : baseOffset+i*32+32])
		absOffset := 32 + elementOffset

		if absOffset+32 > uint64(len(data)) {
			continue
		}
		bytesLen := common.DecodeUint64FromBytes(data[absOffset : absOffset+32])
		if bytesLen == 0 {
			continue
		}
		if absOffset+32+bytesLen > uint64(len(data)) {
			continue
		}

		bytesData := data[absOffset+32 : absOffset+32+bytesLen]
		if bytesLen == 20 {
			addr, _ := common.DecodeAddress(append(make([]byte, 12), bytesData...))
			pools = append(pools, addr)
		} else {
			pools = append(pools, views.BytesToHex(bytesData))
		}
	}

	return pools, nil
}

// getInboundRateLimiter fetches the inbound rate limiter state.
func getInboundRateLimiter(ctx *views.ViewContext, chainSel uint64) (map[string]any, error) {
	data, err := common.ExecuteCall(ctx, selectorGetInboundRateLimiter, common.EncodeUint64(chainSel))
	if err != nil {
		return nil, err
	}

	return decodeRateLimiter(data), nil
}

// getOutboundRateLimiter fetches the outbound rate limiter state.
func getOutboundRateLimiter(ctx *views.ViewContext, chainSel uint64) (map[string]any, error) {
	data, err := common.ExecuteCall(ctx, selectorGetOutboundRateLimiter, common.EncodeUint64(chainSel))
	if err != nil {
		return nil, err
	}

	return decodeRateLimiter(data), nil
}

// decodeRateLimiter decodes a RateLimiter.TokenBucket struct.
// TokenBucket: uint128 tokens, uint32 lastUpdated, bool isEnabled, uint128 capacity, uint128 rate
func decodeRateLimiter(data []byte) map[string]any {
	result := make(map[string]any)

	if len(data) < 32 {
		return result
	}

	offset := 0

	// tokens (uint128 stored in 32 bytes)
	if len(data) >= offset+32 {
		tokens := common.DecodeUint64FromBytes(data[offset : offset+32])
		result["tokens"] = tokens
		offset += 32
	}

	// lastUpdated (uint32)
	if len(data) >= offset+32 {
		lastUpdated := common.DecodeUint64FromBytes(data[offset : offset+32])
		result["lastUpdated"] = lastUpdated
		offset += 32
	}

	// isEnabled (bool)
	if len(data) >= offset+32 {
		isEnabled, _ := common.DecodeBool(data[offset : offset+32])
		result["isEnabled"] = isEnabled
		offset += 32
	}

	// capacity (uint128)
	if len(data) >= offset+32 {
		capacity := common.DecodeUint64FromBytes(data[offset : offset+32])
		result["capacity"] = capacity
		offset += 32
	}

	// rate (uint128)
	if len(data) >= offset+32 {
		rate := common.DecodeUint64FromBytes(data[offset : offset+32])
		result["rate"] = rate
	}

	return result
}

// ViewBurnMintTokenPool generates a view of the BurnMintTokenPool contract (v1.5.1).
func ViewBurnMintTokenPool(ctx *views.ViewContext) (map[string]any, error) {
	result := make(map[string]any)

	result["address"] = ctx.AddressHex
	result["chainSelector"] = ctx.ChainSelector
	result["version"] = "1.5.1"

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
		// Fetch token symbol
		if symbol, err := common.GetERC20Symbol(ctx, token); err == nil {
			result["symbol"] = symbol
		} else {
			result["symbol_error"] = err.Error()
		}
	}

	supportedChains, err := getTokenPoolSupportedChains(ctx)
	if err != nil {
		result["supportedChains_error"] = err.Error()
	} else {
		result["supportedChains"] = supportedChains

		// Get remote chain configs for each supported chain
		remoteConfigs := getRemoteChainConfigs(ctx, supportedChains)
		if len(remoteConfigs) > 0 {
			result["remoteChainConfigs"] = remoteConfigs
		}
	}

	// Allow list
	allowList, err := getTokenPoolAllowList(ctx)
	if err == nil {
		result["allowList"] = allowList
	}

	allowListEnabled, err := getTokenPoolAllowListEnabled(ctx)
	if err == nil {
		result["allowListEnabled"] = allowListEnabled
	}

	return result, nil
}

// ViewLockReleaseTokenPool generates a view of the LockReleaseTokenPool contract (v1.5.1).
func ViewLockReleaseTokenPool(ctx *views.ViewContext) (map[string]any, error) {
	result := make(map[string]any)

	result["address"] = ctx.AddressHex
	result["chainSelector"] = ctx.ChainSelector
	result["version"] = "1.5.1"

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
		// Fetch token symbol
		if symbol, err := common.GetERC20Symbol(ctx, token); err == nil {
			result["symbol"] = symbol
		} else {
			result["symbol_error"] = err.Error()
		}
	}

	supportedChains, err := getTokenPoolSupportedChains(ctx)
	if err != nil {
		result["supportedChains_error"] = err.Error()
	} else {
		result["supportedChains"] = supportedChains

		// Get remote chain configs for each supported chain
		remoteConfigs := getRemoteChainConfigs(ctx, supportedChains)
		if len(remoteConfigs) > 0 {
			result["remoteChainConfigs"] = remoteConfigs
		}
	}

	rebalancer, err := getTokenPoolRebalancer(ctx)
	if err == nil {
		result["rebalancer"] = rebalancer
	}

	// Allow list
	allowList, err := getTokenPoolAllowList(ctx)
	if err == nil {
		result["allowList"] = allowList
	}

	allowListEnabled, err := getTokenPoolAllowListEnabled(ctx)
	if err == nil {
		result["allowListEnabled"] = allowListEnabled
	}

	return result, nil
}

// ViewBurnFromMintTokenPool generates a view of the BurnFromMintTokenPool contract (v1.5.1).
// Same interface as BurnMintTokenPool.
func ViewBurnFromMintTokenPool(ctx *views.ViewContext) (map[string]any, error) {
	result := make(map[string]any)

	result["address"] = ctx.AddressHex
	result["chainSelector"] = ctx.ChainSelector
	result["version"] = "1.5.1"

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
		// Fetch token symbol
		if symbol, err := common.GetERC20Symbol(ctx, token); err == nil {
			result["symbol"] = symbol
		} else {
			result["symbol_error"] = err.Error()
		}
	}

	supportedChains, err := getTokenPoolSupportedChains(ctx)
	if err != nil {
		result["supportedChains_error"] = err.Error()
	} else {
		result["supportedChains"] = supportedChains

		// Get remote chain configs for each supported chain
		remoteConfigs := getRemoteChainConfigs(ctx, supportedChains)
		if len(remoteConfigs) > 0 {
			result["remoteChainConfigs"] = remoteConfigs
		}
	}

	// Allow list
	allowList, err := getTokenPoolAllowList(ctx)
	if err == nil {
		result["allowList"] = allowList
	}

	allowListEnabled, err := getTokenPoolAllowListEnabled(ctx)
	if err == nil {
		result["allowListEnabled"] = allowListEnabled
	}

	return result, nil
}

// ViewBurnWithFromMintTokenPool generates a view of the BurnWithFromMintTokenPool contract (v1.5.1).
// Same interface as BurnMintTokenPool.
func ViewBurnWithFromMintTokenPool(ctx *views.ViewContext) (map[string]any, error) {
	result := make(map[string]any)

	result["address"] = ctx.AddressHex
	result["chainSelector"] = ctx.ChainSelector
	result["version"] = "1.5.1"

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
		// Fetch token symbol
		if symbol, err := common.GetERC20Symbol(ctx, token); err == nil {
			result["symbol"] = symbol
		} else {
			result["symbol_error"] = err.Error()
		}
	}

	supportedChains, err := getTokenPoolSupportedChains(ctx)
	if err != nil {
		result["supportedChains_error"] = err.Error()
	} else {
		result["supportedChains"] = supportedChains

		// Get remote chain configs for each supported chain
		remoteConfigs := getRemoteChainConfigs(ctx, supportedChains)
		if len(remoteConfigs) > 0 {
			result["remoteChainConfigs"] = remoteConfigs
		}
	}

	// Allow list
	allowList, err := getTokenPoolAllowList(ctx)
	if err == nil {
		result["allowList"] = allowList
	}

	allowListEnabled, err := getTokenPoolAllowListEnabled(ctx)
	if err == nil {
		result["allowListEnabled"] = allowListEnabled
	}

	return result, nil
}
