package v1_6_1

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"sync"

	"give-me-state-v2/views"
	"give-me-state-v2/views/evm/common"

	gethCommon "github.com/ethereum/go-ethereum/common"
)

// executeTokenPoolCall packs a call using TokenPoolABI, executes it, and returns raw response bytes.
func executeTokenPoolCall(ctx *views.ViewContext, method string, args ...interface{}) ([]byte, error) {
	calldata, err := TokenPoolABI.Pack(method, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to pack %s call: %w", method, err)
	}

	call := views.Call{
		ChainID: ctx.ChainSelector,
		Target:  ctx.Address,
		Data:    calldata,
	}

	result := ctx.TypedOrchestrator.Execute(call)
	if result.Error != nil {
		return nil, fmt.Errorf("%s call failed: %w", method, result.Error)
	}

	return result.Data, nil
}

// getTokenPoolToken fetches the token address.
func getTokenPoolToken(ctx *views.ViewContext) (string, error) {
	data, err := executeTokenPoolCall(ctx, "getToken")
	if err != nil {
		return "", err
	}
	results, err := TokenPoolABI.Unpack("getToken", data)
	if err != nil {
		return "", fmt.Errorf("failed to unpack getToken: %w", err)
	}
	if len(results) == 0 {
		return "", fmt.Errorf("no results from getToken call")
	}
	addr, ok := results[0].(gethCommon.Address)
	if !ok {
		return "", fmt.Errorf("unexpected type for token: %T", results[0])
	}
	return addr.Hex(), nil
}

// getTokenPoolSupportedChains fetches the supported chains.
func getTokenPoolSupportedChains(ctx *views.ViewContext) ([]uint64, error) {
	data, err := executeTokenPoolCall(ctx, "getSupportedChains")
	if err != nil {
		return nil, err
	}
	results, err := TokenPoolABI.Unpack("getSupportedChains", data)
	if err != nil {
		return nil, fmt.Errorf("failed to unpack getSupportedChains: %w", err)
	}
	if len(results) == 0 {
		return []uint64{}, nil
	}
	chains, ok := results[0].([]uint64)
	if !ok {
		return nil, fmt.Errorf("unexpected type for supported chains: %T", results[0])
	}
	return chains, nil
}

// getTokenPoolRebalancer fetches the rebalancer address.
func getTokenPoolRebalancer(ctx *views.ViewContext) (string, error) {
	data, err := executeTokenPoolCall(ctx, "getRebalancer")
	if err != nil {
		return "", err
	}
	results, err := TokenPoolABI.Unpack("getRebalancer", data)
	if err != nil {
		return "", fmt.Errorf("failed to unpack getRebalancer: %w", err)
	}
	if len(results) == 0 {
		return "", fmt.Errorf("no results from getRebalancer call")
	}
	addr, ok := results[0].(gethCommon.Address)
	if !ok {
		return "", fmt.Errorf("unexpected type for rebalancer: %T", results[0])
	}
	return addr.Hex(), nil
}

// getAllowList fetches the allow list addresses.
func getAllowList(ctx *views.ViewContext) ([]string, error) {
	data, err := executeTokenPoolCall(ctx, "getAllowList")
	if err != nil {
		return nil, err
	}
	results, err := TokenPoolABI.Unpack("getAllowList", data)
	if err != nil {
		return nil, fmt.Errorf("failed to unpack getAllowList: %w", err)
	}
	if len(results) == 0 {
		return []string{}, nil
	}
	addrs, ok := results[0].([]gethCommon.Address)
	if !ok {
		return nil, fmt.Errorf("unexpected type for allow list: %T", results[0])
	}
	addresses := make([]string, len(addrs))
	for i, a := range addrs {
		addresses[i] = a.Hex()
	}
	return addresses, nil
}

// getAllowListEnabled fetches whether allow list is enabled.
func getAllowListEnabled(ctx *views.ViewContext) (bool, error) {
	data, err := executeTokenPoolCall(ctx, "getAllowListEnabled")
	if err != nil {
		return false, err
	}
	results, err := TokenPoolABI.Unpack("getAllowListEnabled", data)
	if err != nil {
		return false, fmt.Errorf("failed to unpack getAllowListEnabled: %w", err)
	}
	if len(results) == 0 {
		return false, fmt.Errorf("no results from getAllowListEnabled call")
	}
	enabled, ok := results[0].(bool)
	if !ok {
		return false, fmt.Errorf("unexpected type for allowListEnabled: %T", results[0])
	}
	return enabled, nil
}

// getRemoteToken fetches the remote token address for a chain.
func getRemoteToken(ctx *views.ViewContext, chainSel uint64) (string, error) {
	data, err := executeTokenPoolCall(ctx, "getRemoteToken", chainSel)
	if err != nil {
		return "", err
	}
	results, err := TokenPoolABI.Unpack("getRemoteToken", data)
	if err != nil {
		return "", fmt.Errorf("failed to unpack getRemoteToken: %w", err)
	}
	if len(results) == 0 {
		return "", nil
	}
	bytesVal, ok := results[0].([]byte)
	if !ok {
		return "", fmt.Errorf("unexpected type for remote token: %T", results[0])
	}
	if len(bytesVal) == 0 {
		return "", nil
	}
	if len(bytesVal) == 20 {
		return gethCommon.BytesToAddress(bytesVal).Hex(), nil
	}
	return "0x" + hex.EncodeToString(bytesVal), nil
}

// getRemotePools fetches the remote pool addresses for a chain.
func getRemotePools(ctx *views.ViewContext, chainSel uint64) ([]string, error) {
	data, err := executeTokenPoolCall(ctx, "getRemotePools", chainSel)
	if err != nil {
		return nil, err
	}
	results, err := TokenPoolABI.Unpack("getRemotePools", data)
	if err != nil {
		return nil, fmt.Errorf("failed to unpack getRemotePools: %w", err)
	}
	if len(results) == 0 {
		return []string{}, nil
	}
	bytesArr, ok := results[0].([][]byte)
	if !ok {
		return nil, fmt.Errorf("unexpected type for remote pools: %T", results[0])
	}
	pools := make([]string, 0, len(bytesArr))
	for _, b := range bytesArr {
		if len(b) == 20 {
			pools = append(pools, gethCommon.BytesToAddress(b).Hex())
		} else if len(b) > 0 {
			pools = append(pools, "0x"+hex.EncodeToString(b))
		}
	}
	return pools, nil
}

// getRateLimiter fetches a rate limiter state for a chain.
func getRateLimiter(ctx *views.ViewContext, method string, chainSel uint64) (map[string]any, error) {
	data, err := executeTokenPoolCall(ctx, method, chainSel)
	if err != nil {
		return nil, err
	}
	results, err := TokenPoolABI.Unpack(method, data)
	if err != nil {
		return nil, fmt.Errorf("failed to unpack %s: %w", method, err)
	}
	if len(results) == 0 {
		return nil, fmt.Errorf("no results from %s call", method)
	}

	bucket, ok := results[0].(struct {
		Tokens      *big.Int `json:"tokens"`
		LastUpdated uint32   `json:"lastUpdated"`
		IsEnabled   bool     `json:"isEnabled"`
		Capacity    *big.Int `json:"capacity"`
		Rate        *big.Int `json:"rate"`
	})
	if !ok {
		return nil, fmt.Errorf("unexpected type for rate limiter: %T", results[0])
	}

	return map[string]any{
		"tokens":      bucket.Tokens.String(),
		"lastUpdated": bucket.LastUpdated,
		"isEnabled":   bucket.IsEnabled,
		"capacity":    bucket.Capacity.String(),
		"rate":        bucket.Rate.String(),
	}, nil
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

			if remoteToken, err := getRemoteToken(ctx, cs); err == nil && remoteToken != "" {
				config["remoteTokenAddress"] = remoteToken
			}
			if remotePools, err := getRemotePools(ctx, cs); err == nil && len(remotePools) > 0 {
				config["remotePoolAddresses"] = remotePools
			}
			if inbound, err := getRateLimiter(ctx, "getCurrentInboundRateLimiterState", cs); err == nil {
				config["inboundRateLimiterConfig"] = inbound
			}
			if outbound, err := getRateLimiter(ctx, "getCurrentOutboundRateLimiterState", cs); err == nil {
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

// ViewBurnMintTokenPool generates a view of the BurnMintTokenPool contract (v1.6.1).
// Uses ABI bindings for proper decoding.
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
// Uses ABI bindings for proper decoding.
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
