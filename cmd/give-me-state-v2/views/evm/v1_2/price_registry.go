package v1_2

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"

	"give-me-state-v2/views"
	"give-me-state-v2/views/evm/common"

	gethCommon "github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_2_0/price_registry"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

// PriceRegistryABI is parsed once at startup.
var PriceRegistryABI abi.ABI

func init() {
	parsed, err := price_registry.PriceRegistryMetaData.GetAbi()
	if err != nil {
		panic(fmt.Sprintf("failed to parse PriceRegistry ABI: %v", err))
	}
	PriceRegistryABI = *parsed
}

// executePriceRegistryCall packs a call, executes it, and returns raw response bytes.
func executePriceRegistryCall(ctx *views.ViewContext, method string, args ...interface{}) ([]byte, error) {
	calldata, err := PriceRegistryABI.Pack(method, args...)
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

// getPriceRegistryFeeTokens fetches the list of fee token addresses using ABI bindings.
func getPriceRegistryFeeTokens(ctx *views.ViewContext) ([]string, error) {
	data, err := executePriceRegistryCall(ctx, "getFeeTokens")
	if err != nil {
		return nil, err
	}
	results, err := PriceRegistryABI.Unpack("getFeeTokens", data)
	if err != nil {
		return nil, fmt.Errorf("failed to unpack getFeeTokens: %w", err)
	}
	if len(results) == 0 {
		return []string{}, nil
	}
	addrs, ok := results[0].([]gethCommon.Address)
	if !ok {
		return nil, fmt.Errorf("unexpected type for fee tokens: %T", results[0])
	}
	tokens := make([]string, len(addrs))
	for i, a := range addrs {
		tokens[i] = a.Hex()
	}
	return tokens, nil
}

// getPriceRegistryERC20Decimals fetches decimals from an ERC20 token at tokenAddr.
func getPriceRegistryERC20Decimals(ctx *views.ViewContext, tokenAddrHex string) (uint8, error) {
	tokenAddrHex = strings.TrimPrefix(tokenAddrHex, "0x")
	tokenAddr, err := hex.DecodeString(tokenAddrHex)
	if err != nil {
		return 0, err
	}
	if len(tokenAddr) < 20 {
		padded := make([]byte, 20)
		copy(padded[20-len(tokenAddr):], tokenAddr)
		tokenAddr = padded
	}
	selectorERC20Decimals := common.HexToSelector("313ce567")
	calldata := views.ABIEncodeCall(selectorERC20Decimals)
	call := views.Call{ChainID: ctx.ChainSelector, Target: tokenAddr, Data: calldata}
	result := ctx.TypedOrchestrator.Execute(call)
	if result.Error != nil {
		return 0, result.Error
	}
	if len(result.Data) < 32 {
		return 0, nil
	}
	return uint8(result.Data[31]), nil
}

// getPriceRegistryFeeTokensEnriched returns fee tokens with name, symbol, and decimals per token.
func getPriceRegistryFeeTokensEnriched(ctx *views.ViewContext) ([]map[string]any, error) {
	addresses, err := getPriceRegistryFeeTokens(ctx)
	if err != nil {
		return nil, err
	}
	out := make([]map[string]any, 0, len(addresses))
	for _, addr := range addresses {
		entry := map[string]any{"address": addr}

		if name, err := common.GetERC20Name(ctx, addr); err != nil {
			entry["name_error"] = err.Error()
		} else {
			entry["name"] = name
		}
		if symbol, err := common.GetERC20Symbol(ctx, addr); err != nil {
			entry["symbol_error"] = err.Error()
		} else {
			entry["symbol"] = symbol
		}
		if decimals, err := getPriceRegistryERC20Decimals(ctx, addr); err != nil {
			entry["decimals_error"] = err.Error()
		} else {
			entry["decimals"] = decimals
		}

		out = append(out, entry)
	}
	return out, nil
}

// getPriceRegistryStalenessThreshold fetches the staleness threshold using ABI bindings.
func getPriceRegistryStalenessThreshold(ctx *views.ViewContext) (string, error) {
	data, err := executePriceRegistryCall(ctx, "getStalenessThreshold")
	if err != nil {
		return "", err
	}
	results, err := PriceRegistryABI.Unpack("getStalenessThreshold", data)
	if err != nil {
		return "", fmt.Errorf("failed to unpack getStalenessThreshold: %w", err)
	}
	if len(results) == 0 {
		return "0", nil
	}
	val, ok := results[0].(*big.Int)
	if !ok {
		return "", fmt.Errorf("unexpected type for staleness threshold: %T", results[0])
	}
	return val.String(), nil
}

// getPriceRegistryUpdaters fetches the list of price updaters using ABI bindings.
func getPriceRegistryUpdaters(ctx *views.ViewContext) ([]string, error) {
	data, err := executePriceRegistryCall(ctx, "getPriceUpdaters")
	if err != nil {
		return nil, err
	}
	results, err := PriceRegistryABI.Unpack("getPriceUpdaters", data)
	if err != nil {
		return nil, fmt.Errorf("failed to unpack getPriceUpdaters: %w", err)
	}
	if len(results) == 0 {
		return []string{}, nil
	}
	addrs, ok := results[0].([]gethCommon.Address)
	if !ok {
		return nil, fmt.Errorf("unexpected type for price updaters: %T", results[0])
	}
	updaters := make([]string, len(addrs))
	for i, a := range addrs {
		updaters[i] = a.Hex()
	}
	return updaters, nil
}

// ViewPriceRegistry generates a view of the PriceRegistry contract (v1.2.0).
// Uses ABI bindings for proper decoding.
func ViewPriceRegistry(ctx *views.ViewContext) (map[string]any, error) {
	result := make(map[string]any)

	result["address"] = ctx.AddressHex
	result["chainSelector"] = ctx.ChainSelector
	result["version"] = "1.2.0"

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

	feeTokens, err := getPriceRegistryFeeTokensEnriched(ctx)
	if err != nil {
		result["feeTokens_error"] = err.Error()
	} else {
		result["feeTokens"] = feeTokens
	}

	stalenessThreshold, err := getPriceRegistryStalenessThreshold(ctx)
	if err != nil {
		result["stalenessThreshold_error"] = err.Error()
	} else {
		result["stalenessThreshold"] = stalenessThreshold
	}

	updaters, err := getPriceRegistryUpdaters(ctx)
	if err != nil {
		result["updaters_error"] = err.Error()
	} else {
		result["updaters"] = updaters
	}

	return result, nil
}
