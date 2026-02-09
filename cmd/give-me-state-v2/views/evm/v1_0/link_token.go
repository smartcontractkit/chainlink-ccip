package v1_0

import (
	"give-me-state-v2/views"
	"give-me-state-v2/views/evm/common"
	"encoding/hex"
	"math/big"
)

// Function selectors for LinkToken
var (
	// name() returns (string)
	selectorName = common.HexToSelector("06fdde03")
	// symbol() returns (string)
	selectorSymbol = common.HexToSelector("95d89b41")
	// decimals() returns (uint8)
	selectorDecimals = common.HexToSelector("313ce567")
	// totalSupply() returns (uint256)
	selectorTotalSupply = common.HexToSelector("18160ddd")
	// getMinters() returns (address[])
	selectorGetMinters = common.HexToSelector("6b32810b")
	// getBurners() returns (address[])
	selectorGetBurners = common.HexToSelector("86fe8b43")
)

// ViewLinkToken generates a view of the LinkToken contract (v1.0.0).
func ViewLinkToken(ctx *views.ViewContext) (map[string]any, error) {
	result := make(map[string]any)

	result["address"] = ctx.AddressHex
	result["chainSelector"] = ctx.ChainSelector
	result["version"] = "1.0.0"

	// Get owner
	owner, err := common.GetOwner(ctx)
	if err != nil {
		result["owner_error"] = err.Error()
	} else {
		result["owner"] = owner
	}

	// Get typeAndVersion (may not exist on older tokens)
	typeAndVersion, err := common.GetTypeAndVersion(ctx)
	if err != nil {
		// Some LinkTokens don't have typeAndVersion, that's okay
		result["typeAndVersion"] = "LinkToken 1.0.0"
	} else {
		result["typeAndVersion"] = typeAndVersion
	}

	// Get name
	name, err := getName(ctx)
	if err != nil {
		result["name_error"] = err.Error()
	} else {
		result["name"] = name
	}

	// Get symbol
	symbol, err := getSymbol(ctx)
	if err != nil {
		result["symbol_error"] = err.Error()
	} else {
		result["symbol"] = symbol
	}

	// Get decimals
	decimals, err := getDecimals(ctx)
	if err != nil {
		result["decimals_error"] = err.Error()
	} else {
		result["decimals"] = decimals
	}

	// Get totalSupply
	supply, err := getTotalSupply(ctx)
	if err != nil {
		result["supply_error"] = err.Error()
	} else {
		result["supply"] = supply
	}

	// Get minters
	minters, err := getMinters(ctx)
	if err != nil {
		result["minters_error"] = err.Error()
	} else {
		result["minters"] = minters
	}

	// Get burners
	burners, err := getBurners(ctx)
	if err != nil {
		result["burners_error"] = err.Error()
	} else {
		result["burners"] = burners
	}

	return result, nil
}

// getName fetches the token name.
func getName(ctx *views.ViewContext) (string, error) {
	data, err := common.ExecuteCall(ctx, selectorName)
	if err != nil {
		return "", err
	}
	return common.DecodeString(data)
}

// getSymbol fetches the token symbol.
func getSymbol(ctx *views.ViewContext) (string, error) {
	data, err := common.ExecuteCall(ctx, selectorSymbol)
	if err != nil {
		return "", err
	}
	return common.DecodeString(data)
}

// getDecimals fetches the token decimals.
func getDecimals(ctx *views.ViewContext) (uint8, error) {
	data, err := common.ExecuteCall(ctx, selectorDecimals)
	if err != nil {
		return 0, err
	}
	if len(data) < 32 {
		return 0, nil
	}
	return uint8(data[31]), nil
}

// getTotalSupply fetches the total supply.
func getTotalSupply(ctx *views.ViewContext) (string, error) {
	data, err := common.ExecuteCall(ctx, selectorTotalSupply)
	if err != nil {
		return "", err
	}
	if len(data) < 32 {
		return "0", nil
	}
	supply := new(big.Int).SetBytes(data[:32])
	return supply.String(), nil
}

// getMinters fetches the list of minter addresses.
func getMinters(ctx *views.ViewContext) ([]string, error) {
	data, err := common.ExecuteCall(ctx, selectorGetMinters)
	if err != nil {
		return nil, err
	}
	return decodeAddressArray(data)
}

// getBurners fetches the list of burner addresses.
func getBurners(ctx *views.ViewContext) ([]string, error) {
	data, err := common.ExecuteCall(ctx, selectorGetBurners)
	if err != nil {
		return nil, err
	}
	return decodeAddressArray(data)
}

// decodeAddressArray decodes an ABI-encoded dynamic array of addresses.
func decodeAddressArray(data []byte) ([]string, error) {
	if len(data) < 64 {
		return []string{}, nil
	}
	// First 32 bytes: offset to array (usually 0x20)
	// At offset: length
	// Then addresses
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

// ViewStaticLinkToken generates a view of the StaticLinkToken contract (v1.0.0).
// This is a simpler version of LinkToken used for wrapped/static LINK tokens.
func ViewStaticLinkToken(ctx *views.ViewContext) (map[string]any, error) {
	result := make(map[string]any)

	result["address"] = ctx.AddressHex
	result["chainSelector"] = ctx.ChainSelector

	// Get typeAndVersion
	typeAndVersion, err := common.GetTypeAndVersion(ctx)
	if err != nil {
		result["typeAndVersion"] = "StaticLinkToken 1.0.0"
	} else {
		result["typeAndVersion"] = typeAndVersion
	}

	// Get owner (may not exist, return zero address)
	owner, err := common.GetOwner(ctx)
	if err != nil {
		result["owner"] = "0x0000000000000000000000000000000000000000"
	} else {
		result["owner"] = owner
	}

	// Get decimals
	decimals, err := getDecimals(ctx)
	if err != nil {
		result["decimals"] = 18 // Default for LINK
	} else {
		result["decimals"] = decimals
	}

	// Get totalSupply
	supply, err := getTotalSupply(ctx)
	if err != nil {
		result["supply"] = nil
	} else {
		result["supply"] = supply
	}

	return result, nil
}
