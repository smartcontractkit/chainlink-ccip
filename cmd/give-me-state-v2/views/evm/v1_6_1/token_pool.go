package v1_6_1

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
	// getRebalancer() returns (address)
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

func getTokenPoolToken(ctx *views.ViewContext) (string, error) {
	data, err := common.ExecuteCall(ctx, selectorGetToken)
	if err != nil {
		return "", err
	}
	return common.DecodeAddress(data)
}

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

func getTokenPoolRebalancer(ctx *views.ViewContext) (string, error) {
	data, err := common.ExecuteCall(ctx, selectorGetRebalancer)
	if err != nil {
		return "", err
	}
	return common.DecodeAddress(data)
}

func getAllowList(ctx *views.ViewContext) ([]string, error) {
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
		addr, _ := common.DecodeAddress(data[offset : offset+32])
		addresses = append(addresses, addr)
	}
	return addresses, nil
}

func getAllowListEnabled(ctx *views.ViewContext) (bool, error) {
	data, err := common.ExecuteCall(ctx, selectorGetAllowListEnabled)
	if err != nil {
		return false, err
	}
	return common.DecodeBool(data)
}

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

			if remoteToken, err := getRemoteToken(ctx, cs); err == nil && remoteToken != "" {
				config["remoteTokenAddress"] = remoteToken
			}
			if remotePools, err := getRemotePools(ctx, cs); err == nil && len(remotePools) > 0 {
				config["remotePoolAddresses"] = remotePools
			}
			if inbound, err := getInboundRateLimiter(ctx, cs); err == nil {
				config["inboundRateLimiterConfig"] = inbound
			}
			if outbound, err := getOutboundRateLimiter(ctx, cs); err == nil {
				config["outboundRateLimiterConfig"] = outbound
			}

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

func getRemoteToken(ctx *views.ViewContext, chainSel uint64) (string, error) {
	data, err := common.ExecuteCall(ctx, selectorGetRemoteToken, common.EncodeUint64(chainSel))
	if err != nil {
		return "", err
	}
	if len(data) < 64 {
		return "", nil
	}
	bytesLen := common.DecodeUint64FromBytes(data[32:64])
	if bytesLen == 0 {
		return "", nil
	}
	if bytesLen == 20 {
		addr, _ := common.DecodeAddress(data[64:96])
		return addr, nil
	}
	return views.BytesToHex(data[64 : 64+bytesLen]), nil
}

func getRemotePools(ctx *views.ViewContext, chainSel uint64) ([]string, error) {
	data, err := common.ExecuteCall(ctx, selectorGetRemotePools, common.EncodeUint64(chainSel))
	if err != nil {
		return nil, err
	}

	// Returns bytes[] - dynamic array of dynamic bytes
	if len(data) < 64 {
		return []string{}, nil
	}

	arrayLen := common.DecodeUint64FromBytes(data[32:64])
	if arrayLen == 0 {
		return []string{}, nil
	}

	pools := make([]string, 0, arrayLen)
	baseOffset := uint64(64)

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

func getInboundRateLimiter(ctx *views.ViewContext, chainSel uint64) (map[string]any, error) {
	data, err := common.ExecuteCall(ctx, selectorGetInboundRateLimiter, common.EncodeUint64(chainSel))
	if err != nil {
		return nil, err
	}
	return decodeRateLimiter(data), nil
}

func getOutboundRateLimiter(ctx *views.ViewContext, chainSel uint64) (map[string]any, error) {
	data, err := common.ExecuteCall(ctx, selectorGetOutboundRateLimiter, common.EncodeUint64(chainSel))
	if err != nil {
		return nil, err
	}
	return decodeRateLimiter(data), nil
}

func decodeRateLimiter(data []byte) map[string]any {
	result := make(map[string]any)
	if len(data) < 32 {
		return result
	}
	offset := 0
	if len(data) >= offset+32 {
		result["tokens"] = common.DecodeUint64FromBytes(data[offset : offset+32])
		offset += 32
	}
	if len(data) >= offset+32 {
		result["lastUpdated"] = common.DecodeUint64FromBytes(data[offset : offset+32])
		offset += 32
	}
	if len(data) >= offset+32 {
		isEnabled, _ := common.DecodeBool(data[offset : offset+32])
		result["isEnabled"] = isEnabled
		offset += 32
	}
	if len(data) >= offset+32 {
		result["capacity"] = common.DecodeUint64FromBytes(data[offset : offset+32])
		offset += 32
	}
	if len(data) >= offset+32 {
		result["rate"] = common.DecodeUint64FromBytes(data[offset : offset+32])
	}
	return result
}

// ViewBurnMintTokenPool generates a view of the BurnMintTokenPool contract (v1.6.1).
func ViewBurnMintTokenPool(ctx *views.ViewContext) (map[string]any, error) {
	result := make(map[string]any)

	result["address"] = ctx.AddressHex
	result["chainSelector"] = ctx.ChainSelector
	result["version"] = "1.6.1"

	if owner, err := common.GetOwner(ctx); err == nil {
		result["owner"] = owner
	}
	if typeAndVersion, err := common.GetTypeAndVersion(ctx); err == nil {
		result["typeAndVersion"] = typeAndVersion
	}
	if token, err := getTokenPoolToken(ctx); err == nil {
		result["token"] = token
		// Fetch token symbol
		if symbol, err := common.GetERC20Symbol(ctx, token); err == nil {
			result["symbol"] = symbol
		} else {
			result["symbol_error"] = err.Error()
		}
	}

	if supportedChains, err := getTokenPoolSupportedChains(ctx); err == nil {
		result["supportedChains"] = supportedChains
		if remoteConfigs := getRemoteChainConfigs(ctx, supportedChains); len(remoteConfigs) > 0 {
			result["remoteChainConfigs"] = remoteConfigs
		}
	}

	if allowList, err := getAllowList(ctx); err == nil {
		result["allowList"] = allowList
	}
	if allowListEnabled, err := getAllowListEnabled(ctx); err == nil {
		result["allowListEnabled"] = allowListEnabled
	}

	return result, nil
}

// ViewLockReleaseTokenPool generates a view of the LockReleaseTokenPool contract (v1.6.1).
func ViewLockReleaseTokenPool(ctx *views.ViewContext) (map[string]any, error) {
	result := make(map[string]any)

	result["address"] = ctx.AddressHex
	result["chainSelector"] = ctx.ChainSelector
	result["version"] = "1.6.1"

	if owner, err := common.GetOwner(ctx); err == nil {
		result["owner"] = owner
	}
	if typeAndVersion, err := common.GetTypeAndVersion(ctx); err == nil {
		result["typeAndVersion"] = typeAndVersion
	}
	if token, err := getTokenPoolToken(ctx); err == nil {
		result["token"] = token
		// Fetch token symbol
		if symbol, err := common.GetERC20Symbol(ctx, token); err == nil {
			result["symbol"] = symbol
		} else {
			result["symbol_error"] = err.Error()
		}
	}

	if supportedChains, err := getTokenPoolSupportedChains(ctx); err == nil {
		result["supportedChains"] = supportedChains
		if remoteConfigs := getRemoteChainConfigs(ctx, supportedChains); len(remoteConfigs) > 0 {
			result["remoteChainConfigs"] = remoteConfigs
		}
	}

	if rebalancer, err := getTokenPoolRebalancer(ctx); err == nil {
		result["rebalancer"] = rebalancer
	}
	if allowList, err := getAllowList(ctx); err == nil {
		result["allowList"] = allowList
	}
	if allowListEnabled, err := getAllowListEnabled(ctx); err == nil {
		result["allowListEnabled"] = allowListEnabled
	}

	return result, nil
}
