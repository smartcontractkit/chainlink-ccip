package common

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"

	"give-me-state-v2/views"
)

// Common EVM function selectors (first 4 bytes of keccak256 of signature)
var (
	SelectorOwner            = HexToSelector("8da5cb5b")
	SelectorTypeAndVersion   = HexToSelector("181f5a77")
	SelectorGetStaticConfig  = HexToSelector("06285c69")
	SelectorGetDynamicConfig = HexToSelector("7437ff9f")
	SelectorGetOffRamps      = HexToSelector("f4e36afd")
	SelectorGetOnRamps       = HexToSelector("58aa62b5")
	SelectorGetFeeTokens     = HexToSelector("94c50a7a")
	SelectorGetDestChainConfig = HexToSelector("56a0d4d6")
	SelectorGetAllConfiguredTokens = HexToSelector("c9b22e4e")
	SelectorGetExpectedNextSequenceNumber = HexToSelector("d5b89f28")
	SelectorGetSupportedTokens            = HexToSelector("a0e3c0ab")
	SelectorGetSourceChainConfigs = HexToSelector("87b8d879")
)

func HexToSelector(h string) []byte {
	h = strings.TrimPrefix(h, "0x")
	b, _ := hex.DecodeString(h)
	if len(b) > 4 {
		return b[:4]
	}
	return b
}

func MakeCall(ctx *views.ViewContext, selector []byte, args ...[]byte) views.Call {
	calldata := views.ABIEncodeCall(selector, args...)
	return views.Call{
		ChainID: ctx.ChainSelector,
		Target:  ctx.Address,
		Data:    calldata,
	}
}

func ExecuteCall(ctx *views.ViewContext, selector []byte, args ...[]byte) ([]byte, error) {
	call := MakeCall(ctx, selector, args...)
	result := ctx.TypedOrchestrator.Execute(call)
	if result.Error != nil {
		return nil, result.Error
	}
	return result.Data, nil
}

func DecodeString(data []byte) (string, error) {
	if len(data) < 64 {
		return "", fmt.Errorf("data too short for string: %d bytes", len(data))
	}
	offset := new(big.Int).SetBytes(data[0:32]).Uint64()
	if offset+32 > uint64(len(data)) {
		return "", fmt.Errorf("invalid string offset")
	}
	length := new(big.Int).SetBytes(data[offset : offset+32]).Uint64()
	if offset+32+length > uint64(len(data)) {
		return "", fmt.Errorf("invalid string length")
	}
	return string(data[offset+32 : offset+32+length]), nil
}

func DecodeAddress(data []byte) (string, error) {
	if len(data) < 32 {
		return "", fmt.Errorf("data too short for address: %d bytes", len(data))
	}
	addr := data[12:32]
	return "0x" + hex.EncodeToString(addr), nil
}

func DecodeUint64(data []byte) (uint64, error) {
	if len(data) < 32 {
		return 0, fmt.Errorf("data too short for uint64: %d bytes", len(data))
	}
	return new(big.Int).SetBytes(data[:32]).Uint64(), nil
}

func DecodeUint64FromBytes(data []byte) uint64 {
	if len(data) < 32 {
		return 0
	}
	return new(big.Int).SetBytes(data[:32]).Uint64()
}

func DecodeUint256(data []byte) (string, error) {
	if len(data) < 32 {
		return "", fmt.Errorf("data too short for uint256: %d bytes", len(data))
	}
	n := new(big.Int).SetBytes(data[:32])
	return "0x" + n.Text(16), nil
}

func DecodeBool(data []byte) (bool, error) {
	if len(data) < 32 {
		return false, fmt.Errorf("data too short for bool: %d bytes", len(data))
	}
	return data[31] != 0, nil
}

func EncodeUint64(v uint64) []byte {
	b := new(big.Int).SetUint64(v).Bytes()
	padded := make([]byte, 32)
	copy(padded[32-len(b):], b)
	return padded
}

func EncodeAddress(addr []byte) []byte {
	padded := make([]byte, 32)
	if len(addr) == 20 {
		copy(padded[12:], addr)
	} else {
		copy(padded[32-len(addr):], addr)
	}
	return padded
}

func GetOwner(ctx *views.ViewContext) (string, error) {
	data, err := ExecuteCall(ctx, SelectorOwner)
	if err != nil {
		return "", fmt.Errorf("owner() failed: %w", err)
	}
	return DecodeAddress(data)
}

func GetTypeAndVersion(ctx *views.ViewContext) (string, error) {
	data, err := ExecuteCall(ctx, SelectorTypeAndVersion)
	if err != nil {
		return "", fmt.Errorf("typeAndVersion() failed: %w", err)
	}
	return DecodeString(data)
}

var (
	SelectorERC20Symbol = HexToSelector("95d89b41")
	SelectorERC20Name   = HexToSelector("06fdde03")
)

func GetERC20Symbol(ctx *views.ViewContext, tokenAddrHex string) (string, error) {
	tokenAddrHex = strings.TrimPrefix(tokenAddrHex, "0x")
	tokenAddr, err := hex.DecodeString(tokenAddrHex)
	if err != nil {
		return "", fmt.Errorf("invalid token address: %w", err)
	}
	if len(tokenAddr) < 20 {
		padded := make([]byte, 20)
		copy(padded[20-len(tokenAddr):], tokenAddr)
		tokenAddr = padded
	}
	calldata := views.ABIEncodeCall(SelectorERC20Symbol)
	call := views.Call{ChainID: ctx.ChainSelector, Target: tokenAddr, Data: calldata}
	result := ctx.TypedOrchestrator.Execute(call)
	if result.Error != nil {
		return "", fmt.Errorf("symbol() call failed: %w", result.Error)
	}
	return DecodeString(result.Data)
}

func GetERC20Name(ctx *views.ViewContext, tokenAddrHex string) (string, error) {
	tokenAddrHex = strings.TrimPrefix(tokenAddrHex, "0x")
	tokenAddr, err := hex.DecodeString(tokenAddrHex)
	if err != nil {
		return "", fmt.Errorf("invalid token address: %w", err)
	}
	if len(tokenAddr) < 20 {
		padded := make([]byte, 20)
		copy(padded[20-len(tokenAddr):], tokenAddr)
		tokenAddr = padded
	}
	calldata := views.ABIEncodeCall(SelectorERC20Name)
	call := views.Call{ChainID: ctx.ChainSelector, Target: tokenAddr, Data: calldata}
	result := ctx.TypedOrchestrator.Execute(call)
	if result.Error != nil {
		return "", fmt.Errorf("name() call failed: %w", result.Error)
	}
	return DecodeString(result.Data)
}
