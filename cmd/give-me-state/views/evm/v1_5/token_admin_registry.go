package v1_5

import (
	"fmt"
	"sync"

	"call-orchestrator-demo/views"

	"github.com/ethereum/go-ethereum/common"

	token_admin_registry "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/gobindings/generated/latest/token_admin_registry"
)

// =====================================================
// TokenAdminRegistry Helpers
// =====================================================

// executeTokenAdminRegistryCall packs a call, executes it via the orchestrator, and returns raw response bytes.
func executeTokenAdminRegistryCall(ctx *views.ViewContext, method string, args ...interface{}) ([]byte, error) {
	calldata, err := TokenAdminRegistryABI.Pack(method, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to pack %s call: %w", method, err)
	}

	call := views.Call{
		ChainID: ctx.ChainSelector,
		Target:  ctx.Address,
		Data:    calldata,
	}

	result := ctx.CallManager.Execute(call)
	if result.Error != nil {
		return nil, fmt.Errorf("%s call failed: %w", method, result.Error)
	}

	return result.Data, nil
}

// getTokenAdminRegistryOwner fetches the owner address.
func getTokenAdminRegistryOwner(ctx *views.ViewContext) (string, error) {
	data, err := executeTokenAdminRegistryCall(ctx, "owner")
	if err != nil {
		return "", err
	}

	results, err := TokenAdminRegistryABI.Unpack("owner", data)
	if err != nil {
		return "", fmt.Errorf("failed to unpack owner response: %w", err)
	}

	if len(results) == 0 {
		return "", fmt.Errorf("no results from owner call")
	}

	owner, ok := results[0].(common.Address)
	if !ok {
		return "", fmt.Errorf("unexpected type for owner: %T", results[0])
	}

	return owner.Hex(), nil
}

// getTokenAdminRegistryTypeAndVersion fetches the typeAndVersion string.
func getTokenAdminRegistryTypeAndVersion(ctx *views.ViewContext) (string, error) {
	data, err := executeTokenAdminRegistryCall(ctx, "typeAndVersion")
	if err != nil {
		return "", err
	}

	results, err := TokenAdminRegistryABI.Unpack("typeAndVersion", data)
	if err != nil {
		return "", fmt.Errorf("failed to unpack typeAndVersion response: %w", err)
	}

	if len(results) == 0 {
		return "", fmt.Errorf("no results from typeAndVersion call")
	}

	typeAndVersion, ok := results[0].(string)
	if !ok {
		return "", fmt.Errorf("unexpected type for typeAndVersion: %T", results[0])
	}

	return typeAndVersion, nil
}

// getAllConfiguredTokens fetches all configured tokens from the TokenAdminRegistry.
func getAllConfiguredTokens(ctx *views.ViewContext) ([]common.Address, error) {
	// Use max uint64 for maxCount to get all tokens
	maxUint64 := uint64(18446744073709551615)

	data, err := executeTokenAdminRegistryCall(ctx, "getAllConfiguredTokens", uint64(0), maxUint64)
	if err != nil {
		return nil, err
	}

	results, err := TokenAdminRegistryABI.Unpack("getAllConfiguredTokens", data)
	if err != nil {
		return nil, fmt.Errorf("failed to unpack getAllConfiguredTokens response: %w", err)
	}

	if len(results) == 0 {
		return nil, fmt.Errorf("no results from getAllConfiguredTokens call")
	}

	tokens, ok := results[0].([]common.Address)
	if !ok {
		return nil, fmt.Errorf("unexpected type for getAllConfiguredTokens: %T", results[0])
	}

	return tokens, nil
}

// getTokenConfig fetches the token config for a specific token.
func getTokenConfig(ctx *views.ViewContext, tokenAddr common.Address) (token_admin_registry.TokenAdminRegistryTokenConfig, error) {
	data, err := executeTokenAdminRegistryCall(ctx, "getTokenConfig", tokenAddr)
	if err != nil {
		return token_admin_registry.TokenAdminRegistryTokenConfig{}, err
	}

	results, err := TokenAdminRegistryABI.Unpack("getTokenConfig", data)
	if err != nil {
		return token_admin_registry.TokenAdminRegistryTokenConfig{}, fmt.Errorf("failed to unpack getTokenConfig response: %w", err)
	}

	if len(results) == 0 {
		return token_admin_registry.TokenAdminRegistryTokenConfig{}, fmt.Errorf("no results from getTokenConfig call")
	}

	config, ok := results[0].(struct {
		Administrator        common.Address `json:"administrator"`
		PendingAdministrator common.Address `json:"pendingAdministrator"`
		TokenPool            common.Address `json:"tokenPool"`
	})
	if !ok {
		return token_admin_registry.TokenAdminRegistryTokenConfig{}, fmt.Errorf("unexpected type for getTokenConfig: %T", results[0])
	}

	return token_admin_registry.TokenAdminRegistryTokenConfig{
		Administrator:        config.Administrator,
		PendingAdministrator: config.PendingAdministrator,
		TokenPool:            config.TokenPool,
	}, nil
}

// =====================================================
// TokenPool Helpers
// =====================================================

// executeTokenPoolCall packs a call using the TokenPool ABI and executes it.
func executeTokenPoolCall(ctx *views.ViewContext, poolAddr common.Address, method string, args ...interface{}) ([]byte, error) {
	calldata, err := TokenPoolABI.Pack(method, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to pack TokenPool %s call: %w", method, err)
	}

	call := views.Call{
		ChainID: ctx.ChainSelector,
		Target:  poolAddr.Bytes(),
		Data:    calldata,
	}

	result := ctx.CallManager.Execute(call)
	if result.Error != nil {
		return nil, fmt.Errorf("TokenPool %s call failed: %w", method, result.Error)
	}

	return result.Data, nil
}

// getTokenPoolRmnProxy fetches the RMN proxy address from a token pool.
func getTokenPoolRmnProxy(ctx *views.ViewContext, poolAddr common.Address) (common.Address, error) {
	data, err := executeTokenPoolCall(ctx, poolAddr, "getRmnProxy")
	if err != nil {
		return common.Address{}, err
	}

	results, err := TokenPoolABI.Unpack("getRmnProxy", data)
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to unpack getRmnProxy response: %w", err)
	}

	if len(results) == 0 {
		return common.Address{}, fmt.Errorf("no results from getRmnProxy call")
	}

	rmnProxy, ok := results[0].(common.Address)
	if !ok {
		return common.Address{}, fmt.Errorf("unexpected type for getRmnProxy: %T", results[0])
	}

	return rmnProxy, nil
}

// =====================================================
// RMNProxy Helpers
// =====================================================

// executeRmnProxyCall packs a call using the RMNProxy ABI and executes it.
func executeRmnProxyCall(ctx *views.ViewContext, proxyAddr common.Address, method string, args ...interface{}) ([]byte, error) {
	calldata, err := RMNProxyABI.Pack(method, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to pack RMNProxy %s call: %w", method, err)
	}

	call := views.Call{
		ChainID: ctx.ChainSelector,
		Target:  proxyAddr.Bytes(),
		Data:    calldata,
	}

	result := ctx.CallManager.Execute(call)
	if result.Error != nil {
		return nil, fmt.Errorf("RMNProxy %s call failed: %w", method, result.Error)
	}

	return result.Data, nil
}

// getRmnProxyARM fetches the ARM address from an RMN proxy.
func getRmnProxyARM(ctx *views.ViewContext, proxyAddr common.Address) (common.Address, error) {
	data, err := executeRmnProxyCall(ctx, proxyAddr, "getARM")
	if err != nil {
		return common.Address{}, err
	}

	results, err := RMNProxyABI.Unpack("getARM", data)
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to unpack getARM response: %w", err)
	}

	if len(results) == 0 {
		return common.Address{}, fmt.Errorf("no results from getARM call")
	}

	arm, ok := results[0].(common.Address)
	if !ok {
		return common.Address{}, fmt.Errorf("unexpected type for getARM: %T", results[0])
	}

	return arm, nil
}

// getRmnProxyOwner fetches the owner address from an RMN proxy.
func getRmnProxyOwner(ctx *views.ViewContext, proxyAddr common.Address) (common.Address, error) {
	data, err := executeRmnProxyCall(ctx, proxyAddr, "owner")
	if err != nil {
		return common.Address{}, err
	}

	results, err := RMNProxyABI.Unpack("owner", data)
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to unpack owner response: %w", err)
	}

	if len(results) == 0 {
		return common.Address{}, fmt.Errorf("no results from owner call")
	}

	owner, ok := results[0].(common.Address)
	if !ok {
		return common.Address{}, fmt.Errorf("unexpected type for owner: %T", results[0])
	}

	return owner, nil
}

// =====================================================
// ERC20 Helpers
// =====================================================

// executeERC20Call packs a call using the ERC20 ABI, executes it via the orchestrator, and returns raw response bytes.
func executeERC20Call(ctx *views.ViewContext, tokenAddr common.Address, method string, args ...interface{}) ([]byte, error) {
	calldata, err := ERC20ABI.Pack(method, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to pack ERC20 %s call: %w", method, err)
	}

	call := views.Call{
		ChainID: ctx.ChainSelector,
		Target:  tokenAddr.Bytes(),
		Data:    calldata,
	}

	result := ctx.CallManager.Execute(call)
	if result.Error != nil {
		return nil, fmt.Errorf("ERC20 %s call failed: %w", method, result.Error)
	}

	return result.Data, nil
}

// getERC20Name fetches the name of an ERC20 token.
func getERC20Name(ctx *views.ViewContext, tokenAddr common.Address) (string, error) {
	data, err := executeERC20Call(ctx, tokenAddr, "name")
	if err != nil {
		return "", err
	}

	results, err := ERC20ABI.Unpack("name", data)
	if err != nil {
		return "", fmt.Errorf("failed to unpack name response: %w", err)
	}

	if len(results) == 0 {
		return "", fmt.Errorf("no results from name call")
	}

	name, ok := results[0].(string)
	if !ok {
		return "", fmt.Errorf("unexpected type for name: %T", results[0])
	}

	return name, nil
}

// getERC20Symbol fetches the symbol of an ERC20 token.
func getERC20Symbol(ctx *views.ViewContext, tokenAddr common.Address) (string, error) {
	data, err := executeERC20Call(ctx, tokenAddr, "symbol")
	if err != nil {
		return "", err
	}

	results, err := ERC20ABI.Unpack("symbol", data)
	if err != nil {
		return "", fmt.Errorf("failed to unpack symbol response: %w", err)
	}

	if len(results) == 0 {
		return "", fmt.Errorf("no results from symbol call")
	}

	symbol, ok := results[0].(string)
	if !ok {
		return "", fmt.Errorf("unexpected type for symbol: %T", results[0])
	}

	return symbol, nil
}

// getERC20Decimals fetches the decimals of an ERC20 token.
func getERC20Decimals(ctx *views.ViewContext, tokenAddr common.Address) (uint8, error) {
	data, err := executeERC20Call(ctx, tokenAddr, "decimals")
	if err != nil {
		return 0, err
	}

	results, err := ERC20ABI.Unpack("decimals", data)
	if err != nil {
		return 0, fmt.Errorf("failed to unpack decimals response: %w", err)
	}

	if len(results) == 0 {
		return 0, fmt.Errorf("no results from decimals call")
	}

	decimals, ok := results[0].(uint8)
	if !ok {
		return 0, fmt.Errorf("unexpected type for decimals: %T", results[0])
	}

	return decimals, nil
}

// =====================================================
// Token Metadata Collection
// =====================================================

// tokenMetadataResult holds the result of collecting metadata for a single token.
type tokenMetadataResult struct {
	index    int
	metadata map[string]any
}

// collectTokenMetadata collects all token metadata from the TokenAdminRegistry.
// Uses concurrent calls to speed up data collection.
func collectTokenMetadata(ctx *views.ViewContext) ([]map[string]any, error) {
	// Get all configured tokens
	tokens, err := getAllConfiguredTokens(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get all configured tokens: %w", err)
	}

	if len(tokens) == 0 {
		return []map[string]any{}, nil
	}

	// Create results slice and channel for concurrent collection
	results := make([]map[string]any, len(tokens))
	resultChan := make(chan tokenMetadataResult, len(tokens))

	// Launch goroutines to collect metadata for each token concurrently
	var wg sync.WaitGroup
	for i, tokenAddr := range tokens {
		wg.Add(1)
		go func(idx int, addr common.Address) {
			defer wg.Done()
			metadata := collectSingleTokenMetadata(ctx, addr)
			resultChan <- tokenMetadataResult{index: idx, metadata: metadata}
		}(i, tokenAddr)
	}

	// Close channel when all goroutines complete
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	// Collect results
	for result := range resultChan {
		results[result.index] = result.metadata
	}

	return results, nil
}

// collectSingleTokenMetadata collects metadata for a single token concurrently.
func collectSingleTokenMetadata(ctx *views.ViewContext, tokenAddr common.Address) map[string]any {
	tokenMetadata := map[string]any{
		"address": tokenAddr.Hex(),
	}

	// Use a WaitGroup to make all calls for this token concurrently
	var wg sync.WaitGroup
	var mu sync.Mutex

	// Get token config from TokenAdminRegistry
	var tokenConfig token_admin_registry.TokenAdminRegistryTokenConfig
	var tokenConfigErr error
	wg.Add(1)
	go func() {
		defer wg.Done()
		tokenConfig, tokenConfigErr = getTokenConfig(ctx, tokenAddr)
		mu.Lock()
		if tokenConfigErr != nil {
			tokenMetadata["tokenConfig_error"] = tokenConfigErr.Error()
		} else {
			tokenMetadata["admin"] = tokenConfig.Administrator.Hex()
			tokenMetadata["pendingAdministrator"] = tokenConfig.PendingAdministrator.Hex()
		}
		mu.Unlock()
	}()

	// Get ERC20 name
	wg.Add(1)
	go func() {
		defer wg.Done()
		name, err := getERC20Name(ctx, tokenAddr)
		mu.Lock()
		if err != nil {
			tokenMetadata["name_error"] = err.Error()
		} else {
			tokenMetadata["name"] = name
		}
		mu.Unlock()
	}()

	// Get ERC20 symbol
	wg.Add(1)
	go func() {
		defer wg.Done()
		symbol, err := getERC20Symbol(ctx, tokenAddr)
		mu.Lock()
		if err != nil {
			tokenMetadata["symbol_error"] = err.Error()
		} else {
			tokenMetadata["symbol"] = symbol
		}
		mu.Unlock()
	}()

	// Get ERC20 decimals
	wg.Add(1)
	go func() {
		defer wg.Done()
		decimals, err := getERC20Decimals(ctx, tokenAddr)
		mu.Lock()
		if err != nil {
			tokenMetadata["decimals_error"] = err.Error()
		} else {
			tokenMetadata["decimals"] = decimals
		}
		mu.Unlock()
	}()

	// Wait for all initial calls to complete
	wg.Wait()

	// Now get pool metadata if token has a pool (needs tokenConfig result)
	if tokenConfigErr == nil && tokenConfig.TokenPool != (common.Address{}) {
		poolMetadata := collectTokenPoolMetadataConcurrent(ctx, tokenConfig.TokenPool)
		tokenMetadata["tokenPool"] = poolMetadata
	} else if tokenConfigErr == nil {
		tokenMetadata["tokenPool"] = nil
	}

	return tokenMetadata
}

// collectTokenPoolMetadataConcurrent collects metadata from a token pool concurrently.
func collectTokenPoolMetadataConcurrent(ctx *views.ViewContext, poolAddr common.Address) map[string]any {
	poolMetadata := map[string]any{
		"address": poolAddr.Hex(),
	}

	// Get RMN proxy from TokenPool
	rmnProxyAddr, err := getTokenPoolRmnProxy(ctx, poolAddr)
	if err != nil {
		poolMetadata["rmnProxy_error"] = err.Error()
		return poolMetadata
	}

	// Collect RMN proxy metadata concurrently
	rmnProxyMetadata := map[string]any{
		"address": rmnProxyAddr.Hex(),
	}

	var wg sync.WaitGroup
	var mu sync.Mutex

	// Get ARM from RMNProxy
	wg.Add(1)
	go func() {
		defer wg.Done()
		arm, err := getRmnProxyARM(ctx, rmnProxyAddr)
		mu.Lock()
		if err != nil {
			rmnProxyMetadata["arm_error"] = err.Error()
		} else {
			rmnProxyMetadata["arm"] = arm.Hex()
		}
		mu.Unlock()
	}()

	// Get owner from RMNProxy
	wg.Add(1)
	go func() {
		defer wg.Done()
		owner, err := getRmnProxyOwner(ctx, rmnProxyAddr)
		mu.Lock()
		if err != nil {
			rmnProxyMetadata["owner_error"] = err.Error()
		} else {
			rmnProxyMetadata["owner"] = owner.Hex()
		}
		mu.Unlock()
	}()

	wg.Wait()
	poolMetadata["rmnProxy"] = rmnProxyMetadata

	return poolMetadata
}

// =====================================================
// View Function
// =====================================================

// ViewTokenAdminRegistry generates a view of the TokenAdminRegistry contract (v1.5.0).
// Uses the generated bindings to pack/unpack calls.
func ViewTokenAdminRegistry(ctx *views.ViewContext) (map[string]any, error) {
	result := make(map[string]any)

	// Basic info
	result["address"] = ctx.AddressHex
	result["chainSelector"] = ctx.ChainSelector
	result["version"] = "1.5.0"

	// Get owner using bindings
	owner, err := getTokenAdminRegistryOwner(ctx)
	if err != nil {
		result["owner_error"] = err.Error()
	} else {
		result["owner"] = owner
	}

	// Get typeAndVersion using bindings
	typeAndVersion, err := getTokenAdminRegistryTypeAndVersion(ctx)
	if err != nil {
		result["typeAndVersion_error"] = err.Error()
	} else {
		result["typeAndVersion"] = typeAndVersion
	}

	// Collect all token metadata
	tokens, err := collectTokenMetadata(ctx)
	if err != nil {
		result["tokens_error"] = err.Error()
	} else {
		result["tokens"] = tokens
	}

	return result, nil
}
