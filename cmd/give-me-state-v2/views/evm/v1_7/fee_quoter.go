package v1_7

import (
	"fmt"

	"give-me-state-v2/views"

	"github.com/ethereum/go-ethereum/common"
)

// packFeeQuoterCall packs a method call using the FeeQuoter ABI and returns the calldata bytes.
func packFeeQuoterCall(method string, args ...interface{}) ([]byte, error) {
	return FeeQuoterABI.Pack(method, args...)
}

// executeFeeQuoterCall packs a call, executes it via the orchestrator, and returns raw response bytes.
func executeFeeQuoterCall(ctx *views.ViewContext, method string, args ...interface{}) ([]byte, error) {
	calldata, err := packFeeQuoterCall(method, args...)
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

// getFeeQuoterOwner fetches the owner address using the FeeQuoter bindings.
func getFeeQuoterOwner(ctx *views.ViewContext) (string, error) {
	data, err := executeFeeQuoterCall(ctx, "owner")
	if err != nil {
		return "", err
	}

	results, err := FeeQuoterABI.Unpack("owner", data)
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

// getFeeQuoterTypeAndVersion fetches the typeAndVersion string using the FeeQuoter bindings.
func getFeeQuoterTypeAndVersion(ctx *views.ViewContext) (string, error) {
	data, err := executeFeeQuoterCall(ctx, "typeAndVersion")
	if err != nil {
		return "", err
	}

	results, err := FeeQuoterABI.Unpack("typeAndVersion", data)
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

// getFeeQuoterFeeTokens fetches the list of fee tokens using the FeeQuoter bindings.
func getFeeQuoterFeeTokens(ctx *views.ViewContext) ([]common.Address, error) {
	data, err := executeFeeQuoterCall(ctx, "getFeeTokens")
	if err != nil {
		return nil, err
	}

	results, err := FeeQuoterABI.Unpack("getFeeTokens", data)
	if err != nil {
		return nil, fmt.Errorf("failed to unpack getFeeTokens response: %w", err)
	}

	if len(results) == 0 {
		return nil, fmt.Errorf("no results from getFeeTokens call")
	}

	tokens, ok := results[0].([]common.Address)
	if !ok {
		return nil, fmt.Errorf("unexpected type for getFeeTokens: %T", results[0])
	}

	return tokens, nil
}

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

	result := ctx.TypedOrchestrator.Execute(call)
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

// collectFeeTokens fetches fee tokens and their ERC20 metadata.
func collectFeeTokens(ctx *views.ViewContext) ([]map[string]any, error) {
	tokenAddrs, err := getFeeQuoterFeeTokens(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get fee tokens: %w", err)
	}

	feeTokens := make([]map[string]any, 0, len(tokenAddrs))

	for _, tokenAddr := range tokenAddrs {
		tokenInfo := map[string]any{
			"address":       tokenAddr.Hex(),
			"chainSelector": ctx.ChainSelector,
		}

		name, err := getERC20Name(ctx, tokenAddr)
		if err != nil {
			tokenInfo["name_error"] = err.Error()
		} else {
			tokenInfo["name"] = name
		}

		symbol, err := getERC20Symbol(ctx, tokenAddr)
		if err != nil {
			tokenInfo["symbol_error"] = err.Error()
		} else {
			tokenInfo["symbol"] = symbol
		}

		decimals, err := getERC20Decimals(ctx, tokenAddr)
		if err != nil {
			tokenInfo["decimals_error"] = err.Error()
		} else {
			tokenInfo["decimals"] = decimals
		}

		feeTokens = append(feeTokens, tokenInfo)
	}

	return feeTokens, nil
}

// ViewFeeQuoter generates a view of the FeeQuoter contract (v1.7.0).
// Uses the generated bindings to pack/unpack calls.
func ViewFeeQuoter(ctx *views.ViewContext) (map[string]any, error) {
	result := make(map[string]any)

	result["address"] = ctx.AddressHex
	result["chainSelector"] = ctx.ChainSelector
	result["version"] = "1.7.0"

	owner, err := getFeeQuoterOwner(ctx)
	if err != nil {
		result["owner_error"] = err.Error()
	} else {
		result["owner"] = owner
	}

	typeAndVersion, err := getFeeQuoterTypeAndVersion(ctx)
	if err != nil {
		result["typeAndVersion_error"] = err.Error()
	} else {
		result["typeAndVersion"] = typeAndVersion
	}

	feeTokens, err := collectFeeTokens(ctx)
	if err != nil {
		result["feeTokens_error"] = err.Error()
	} else {
		result["feeTokens"] = feeTokens
	}

	return result, nil
}
