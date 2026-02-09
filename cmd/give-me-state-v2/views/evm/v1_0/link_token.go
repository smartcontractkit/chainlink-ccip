package v1_0

import (
	"fmt"
	"math/big"
	"strings"

	"give-me-state-v2/views"
	"give-me-state-v2/views/evm/common"

	gethCommon "github.com/ethereum/go-ethereum/common"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

const linkTokenABIJson = `[{"inputs":[],"name":"name","outputs":[{"internalType":"string","name":"","type":"string"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"symbol","outputs":[{"internalType":"string","name":"","type":"string"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"decimals","outputs":[{"internalType":"uint8","name":"","type":"uint8"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"totalSupply","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"getMinters","outputs":[{"internalType":"address[]","name":"","type":"address[]"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"getBurners","outputs":[{"internalType":"address[]","name":"","type":"address[]"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"owner","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"typeAndVersion","outputs":[{"internalType":"string","name":"","type":"string"}],"stateMutability":"view","type":"function"}]`

var linkTokenABI abi.ABI

func init() {
	var err error
	linkTokenABI, err = abi.JSON(strings.NewReader(linkTokenABIJson))
	if err != nil {
		panic("Failed to parse LinkToken ABI: " + err.Error())
	}
}

// executeLinkTokenCall packs a call, executes it, and returns raw response bytes.
func executeLinkTokenCall(ctx *views.ViewContext, method string, args ...interface{}) ([]byte, error) {
	calldata, err := linkTokenABI.Pack(method, args...)
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

// getName fetches the token name using ABI bindings.
func getName(ctx *views.ViewContext) (string, error) {
	data, err := executeLinkTokenCall(ctx, "name")
	if err != nil {
		return "", err
	}
	results, err := linkTokenABI.Unpack("name", data)
	if err != nil {
		return "", fmt.Errorf("failed to unpack name: %w", err)
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

// getSymbol fetches the token symbol using ABI bindings.
func getSymbol(ctx *views.ViewContext) (string, error) {
	data, err := executeLinkTokenCall(ctx, "symbol")
	if err != nil {
		return "", err
	}
	results, err := linkTokenABI.Unpack("symbol", data)
	if err != nil {
		return "", fmt.Errorf("failed to unpack symbol: %w", err)
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

// getDecimals fetches the token decimals using ABI bindings.
func getDecimals(ctx *views.ViewContext) (uint8, error) {
	data, err := executeLinkTokenCall(ctx, "decimals")
	if err != nil {
		return 0, err
	}
	results, err := linkTokenABI.Unpack("decimals", data)
	if err != nil {
		return 0, fmt.Errorf("failed to unpack decimals: %w", err)
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

// getTotalSupply fetches the total supply using ABI bindings.
func getTotalSupply(ctx *views.ViewContext) (string, error) {
	data, err := executeLinkTokenCall(ctx, "totalSupply")
	if err != nil {
		return "", err
	}
	results, err := linkTokenABI.Unpack("totalSupply", data)
	if err != nil {
		return "", fmt.Errorf("failed to unpack totalSupply: %w", err)
	}
	if len(results) == 0 {
		return "0", nil
	}
	supply, ok := results[0].(*big.Int)
	if !ok {
		return "", fmt.Errorf("unexpected type for totalSupply: %T", results[0])
	}
	return supply.String(), nil
}

// getMinters fetches the list of minter addresses using ABI bindings.
func getMinters(ctx *views.ViewContext) ([]string, error) {
	data, err := executeLinkTokenCall(ctx, "getMinters")
	if err != nil {
		return nil, err
	}
	results, err := linkTokenABI.Unpack("getMinters", data)
	if err != nil {
		return nil, fmt.Errorf("failed to unpack getMinters: %w", err)
	}
	if len(results) == 0 {
		return []string{}, nil
	}
	addrs, ok := results[0].([]gethCommon.Address)
	if !ok {
		return nil, fmt.Errorf("unexpected type for minters: %T", results[0])
	}
	minters := make([]string, len(addrs))
	for i, a := range addrs {
		minters[i] = a.Hex()
	}
	return minters, nil
}

// getBurners fetches the list of burner addresses using ABI bindings.
func getBurners(ctx *views.ViewContext) ([]string, error) {
	data, err := executeLinkTokenCall(ctx, "getBurners")
	if err != nil {
		return nil, err
	}
	results, err := linkTokenABI.Unpack("getBurners", data)
	if err != nil {
		return nil, fmt.Errorf("failed to unpack getBurners: %w", err)
	}
	if len(results) == 0 {
		return []string{}, nil
	}
	addrs, ok := results[0].([]gethCommon.Address)
	if !ok {
		return nil, fmt.Errorf("unexpected type for burners: %T", results[0])
	}
	burners := make([]string, len(addrs))
	for i, a := range addrs {
		burners[i] = a.Hex()
	}
	return burners, nil
}

// ViewLinkToken generates a view of the LinkToken contract (v1.0.0).
// Uses bespoke ABI JSON for proper decoding (no Go bindings available).
func ViewLinkToken(ctx *views.ViewContext) (map[string]any, error) {
	result := make(map[string]any)

	result["address"] = ctx.AddressHex
	result["chainSelector"] = ctx.ChainSelector
	result["version"] = "1.0.0"

	owner, err := common.GetOwner(ctx)
	if err != nil {
		result["owner_error"] = err.Error()
	} else {
		result["owner"] = owner
	}

	typeAndVersion, err := common.GetTypeAndVersion(ctx)
	if err != nil {
		result["typeAndVersion"] = "LinkToken 1.0.0"
	} else {
		result["typeAndVersion"] = typeAndVersion
	}

	name, err := getName(ctx)
	if err != nil {
		result["name_error"] = err.Error()
	} else {
		result["name"] = name
	}

	symbol, err := getSymbol(ctx)
	if err != nil {
		result["symbol_error"] = err.Error()
	} else {
		result["symbol"] = symbol
	}

	decimals, err := getDecimals(ctx)
	if err != nil {
		result["decimals_error"] = err.Error()
	} else {
		result["decimals"] = decimals
	}

	supply, err := getTotalSupply(ctx)
	if err != nil {
		result["supply_error"] = err.Error()
	} else {
		result["supply"] = supply
	}

	minters, err := getMinters(ctx)
	if err != nil {
		result["minters_error"] = err.Error()
	} else {
		result["minters"] = minters
	}

	burners, err := getBurners(ctx)
	if err != nil {
		result["burners_error"] = err.Error()
	} else {
		result["burners"] = burners
	}

	return result, nil
}

// ViewStaticLinkToken generates a view of the StaticLinkToken contract (v1.0.0).
// This is a simpler version of LinkToken used for wrapped/static LINK tokens.
func ViewStaticLinkToken(ctx *views.ViewContext) (map[string]any, error) {
	result := make(map[string]any)

	result["address"] = ctx.AddressHex
	result["chainSelector"] = ctx.ChainSelector

	typeAndVersion, err := common.GetTypeAndVersion(ctx)
	if err != nil {
		result["typeAndVersion"] = "StaticLinkToken 1.0.0"
	} else {
		result["typeAndVersion"] = typeAndVersion
	}

	owner, err := common.GetOwner(ctx)
	if err != nil {
		result["owner"] = "0x0000000000000000000000000000000000000000"
	} else {
		result["owner"] = owner
	}

	decimals, err := getDecimals(ctx)
	if err != nil {
		result["decimals"] = 18
	} else {
		result["decimals"] = decimals
	}

	supply, err := getTotalSupply(ctx)
	if err != nil {
		result["supply"] = nil
	} else {
		result["supply"] = supply
	}

	return result, nil
}
