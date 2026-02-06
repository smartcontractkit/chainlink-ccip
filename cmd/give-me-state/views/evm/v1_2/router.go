package v1_2

import (
	"fmt"

	"call-orchestrator-demo/views"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"

	router "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/gobindings/generated/latest/router"
)

var (
	RouterABI abi.ABI
)

func init() {
	// Parse the Router ABI once at startup (uses ccv local bindings)
	parsedRouter, err := router.RouterMetaData.GetAbi()
	if err != nil {
		panic(fmt.Sprintf("failed to parse Router ABI: %v", err))
	}
	RouterABI = *parsedRouter
}

// packRouterCall packs a method call using the Router ABI and returns the calldata bytes.
func packRouterCall(method string, args ...interface{}) ([]byte, error) {
	return RouterABI.Pack(method, args...)
}

// executeRouterCall packs a call, executes it via the orchestrator, and returns raw response bytes.
func executeRouterCall(ctx *views.ViewContext, method string, args ...interface{}) ([]byte, error) {
	calldata, err := packRouterCall(method, args...)
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

// getRouterOwner fetches the owner address using the Router bindings.
func getRouterOwner(ctx *views.ViewContext) (string, error) {
	data, err := executeRouterCall(ctx, "owner")
	if err != nil {
		return "", err
	}

	results, err := RouterABI.Unpack("owner", data)
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

// getRouterTypeAndVersion fetches the typeAndVersion string using the Router bindings.
func getRouterTypeAndVersion(ctx *views.ViewContext) (string, error) {
	data, err := executeRouterCall(ctx, "typeAndVersion")
	if err != nil {
		return "", err
	}

	results, err := RouterABI.Unpack("typeAndVersion", data)
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

// getRouterOffRamps fetches all offRamps using the Router bindings.
func getRouterOffRamps(ctx *views.ViewContext) ([]router.RouterOffRamp, error) {
	data, err := executeRouterCall(ctx, "getOffRamps")
	if err != nil {
		return nil, err
	}

	results, err := RouterABI.Unpack("getOffRamps", data)
	if err != nil {
		return nil, fmt.Errorf("failed to unpack getOffRamps response: %w", err)
	}

	if len(results) == 0 {
		return nil, fmt.Errorf("no results from getOffRamps call")
	}

	offRamps, ok := results[0].([]struct {
		SourceChainSelector uint64         `json:"sourceChainSelector"`
		OffRamp             common.Address `json:"offRamp"`
	})
	if !ok {
		return nil, fmt.Errorf("unexpected type for getOffRamps: %T", results[0])
	}

	// Convert to binding type
	result := make([]router.RouterOffRamp, len(offRamps))
	for i, or := range offRamps {
		result[i] = router.RouterOffRamp{
			SourceChainSelector: or.SourceChainSelector,
			OffRamp:             or.OffRamp,
		}
	}

	return result, nil
}

// getRouterOnRamp fetches the onRamp for a specific destination chain selector.
func getRouterOnRamp(ctx *views.ViewContext, destChainSelector uint64) (common.Address, error) {
	data, err := executeRouterCall(ctx, "getOnRamp", destChainSelector)
	if err != nil {
		return common.Address{}, err
	}

	results, err := RouterABI.Unpack("getOnRamp", data)
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to unpack getOnRamp response: %w", err)
	}

	if len(results) == 0 {
		return common.Address{}, fmt.Errorf("no results from getOnRamp call")
	}

	onRamp, ok := results[0].(common.Address)
	if !ok {
		return common.Address{}, fmt.Errorf("unexpected type for getOnRamp: %T", results[0])
	}

	return onRamp, nil
}

// getRouterWrappedNative fetches the wrapped native token address.
func getRouterWrappedNative(ctx *views.ViewContext) (string, error) {
	data, err := executeRouterCall(ctx, "getWrappedNative")
	if err != nil {
		return "", err
	}

	results, err := RouterABI.Unpack("getWrappedNative", data)
	if err != nil {
		return "", fmt.Errorf("failed to unpack getWrappedNative response: %w", err)
	}

	if len(results) == 0 {
		return "", fmt.Errorf("no results from getWrappedNative call")
	}

	wrappedNative, ok := results[0].(common.Address)
	if !ok {
		return "", fmt.Errorf("unexpected type for getWrappedNative: %T", results[0])
	}

	return wrappedNative.Hex(), nil
}

// getRouterArmProxy fetches the ARM proxy address.
func getRouterArmProxy(ctx *views.ViewContext) (string, error) {
	data, err := executeRouterCall(ctx, "getArmProxy")
	if err != nil {
		return "", err
	}

	results, err := RouterABI.Unpack("getArmProxy", data)
	if err != nil {
		return "", fmt.Errorf("failed to unpack getArmProxy response: %w", err)
	}

	if len(results) == 0 {
		return "", fmt.Errorf("no results from getArmProxy call")
	}

	armProxy, ok := results[0].(common.Address)
	if !ok {
		return "", fmt.Errorf("unexpected type for getArmProxy: %T", results[0])
	}

	return armProxy.Hex(), nil
}

// collectRouterMetadata collects all router metadata including offRamps and onRamps.
func collectRouterMetadata(ctx *views.ViewContext) (map[string]any, error) {
	// Get all offRamps
	offRamps, err := getRouterOffRamps(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get offRamps: %w", err)
	}

	// Group offRamps by sourceChainSelector
	offRampsByChain := make(map[uint64][]string)
	uniqueChainSelectors := make(map[uint64]struct{})
	for _, offRamp := range offRamps {
		sourceChainSelector := offRamp.SourceChainSelector
		offRampsByChain[sourceChainSelector] = append(offRampsByChain[sourceChainSelector], offRamp.OffRamp.Hex())
		uniqueChainSelectors[sourceChainSelector] = struct{}{}
	}

	// Get onRamp for each unique sourceChainSelector
	onRampsByChain := make(map[uint64]string)
	for chainSelector := range uniqueChainSelectors {
		onRamp, err := getRouterOnRamp(ctx, chainSelector)
		if err != nil {
			// Log error but continue
			onRampsByChain[chainSelector] = fmt.Sprintf("error: %s", err.Error())
			continue
		}
		onRampsByChain[chainSelector] = onRamp.Hex()
	}

	// Convert to string keys for JSON serialization
	onRampsMap := make(map[string]string)
	for selector, onRampAddr := range onRampsByChain {
		onRampsMap[fmt.Sprintf("%d", selector)] = onRampAddr
	}

	offRampsMap := make(map[string][]string)
	for selector, offRampAddrs := range offRampsByChain {
		offRampsMap[fmt.Sprintf("%d", selector)] = offRampAddrs
	}

	return map[string]any{
		"onRamps":  onRampsMap,
		"offRamps": offRampsMap,
	}, nil
}

// ViewRouter generates a view of the Router contract (v1.2.0).
// Uses the generated bindings to pack/unpack calls.
func ViewRouter(ctx *views.ViewContext) (map[string]any, error) {
	result := make(map[string]any)

	// Basic info
	result["address"] = ctx.AddressHex
	result["chainSelector"] = ctx.ChainSelector
	result["version"] = "1.2.0"

	// Get owner using bindings
	owner, err := getRouterOwner(ctx)
	if err != nil {
		result["owner_error"] = err.Error()
	} else {
		result["owner"] = owner
	}

	// Get typeAndVersion using bindings
	typeAndVersion, err := getRouterTypeAndVersion(ctx)
	if err != nil {
		result["typeAndVersion_error"] = err.Error()
	} else {
		result["typeAndVersion"] = typeAndVersion
	}

	// Get wrappedNative
	wrappedNative, err := getRouterWrappedNative(ctx)
	if err != nil {
		result["wrappedNative_error"] = err.Error()
	} else {
		result["wrappedNative"] = wrappedNative
	}

	// Get armProxy
	armProxy, err := getRouterArmProxy(ctx)
	if err != nil {
		result["armProxy_error"] = err.Error()
	} else {
		result["armProxy"] = armProxy
	}

	// Collect router metadata (onRamps and offRamps)
	routerMetadata, err := collectRouterMetadata(ctx)
	if err != nil {
		result["routerMetadata_error"] = err.Error()
	} else {
		result["onRamps"] = routerMetadata["onRamps"]
		result["offRamps"] = routerMetadata["offRamps"]
	}

	return result, nil
}
