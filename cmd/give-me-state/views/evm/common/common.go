package common

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"

	"call-orchestrator-demo/views"
)

// Common EVM function selectors (first 4 bytes of keccak256 of signature)
var (
	// Common getters
	SelectorOwner            = HexToSelector("8da5cb5b") // owner()
	SelectorTypeAndVersion   = HexToSelector("181f5a77") // typeAndVersion()
	SelectorGetStaticConfig  = HexToSelector("06285c69") // getStaticConfig()
	SelectorGetDynamicConfig = HexToSelector("7437ff9f") // getDynamicConfig() - Note: actual selector may vary

	// Router specific
	SelectorGetOffRamps = HexToSelector("f4e36afd") // getOffRamps()
	SelectorGetOnRamps  = HexToSelector("58aa62b5") // getOnRamps() - Note: may vary by version

	// FeeQuoter specific
	SelectorGetFeeTokens       = HexToSelector("94c50a7a") // getFeeTokens()
	SelectorGetDestChainConfig = HexToSelector("56a0d4d6") // getDestChainConfig(uint64)

	// TokenAdminRegistry specific
	SelectorGetAllConfiguredTokens = HexToSelector("c9b22e4e") // getAllConfiguredTokens(uint64,uint64)

	// OnRamp specific
	SelectorGetExpectedNextSequenceNumber = HexToSelector("d5b89f28") // getExpectedNextSequenceNumber()
	SelectorGetSupportedTokens            = HexToSelector("a0e3c0ab") // getSupportedTokens(uint64) - varies by version

	// OffRamp specific
	SelectorGetSourceChainConfigs = HexToSelector("87b8d879") // getSourceChainConfigs(uint64) - varies
)

// HexToSelector converts a hex string to a 4-byte selector.
func HexToSelector(h string) []byte {
	h = strings.TrimPrefix(h, "0x")
	b, _ := hex.DecodeString(h)
	if len(b) > 4 {
		return b[:4]
	}
	return b
}

// MakeCall creates a Call for an EVM contract.
func MakeCall(ctx *views.ViewContext, selector []byte, args ...[]byte) views.Call {
	calldata := views.ABIEncodeCall(selector, args...)
	return views.Call{
		ChainID: ctx.ChainSelector,
		Target:  ctx.Address,
		Data:    calldata,
	}
}

// ExecuteCall executes a call and returns the result.
func ExecuteCall(ctx *views.ViewContext, selector []byte, args ...[]byte) ([]byte, error) {
	call := MakeCall(ctx, selector, args...)
	result := ctx.CallManager.Execute(call)
	if result.Error != nil {
		return nil, result.Error
	}
	return result.Data, nil
}

// DecodeString decodes an ABI-encoded string from call result.
func DecodeString(data []byte) (string, error) {
	if len(data) < 64 {
		return "", fmt.Errorf("data too short for string: %d bytes", len(data))
	}

	// First 32 bytes: offset to string data
	// Next 32 bytes at offset: length of string
	// Following bytes: actual string data

	// Read offset (should be 0x20 = 32 for a single return value)
	offset := new(big.Int).SetBytes(data[0:32]).Uint64()
	if offset+32 > uint64(len(data)) {
		return "", fmt.Errorf("invalid string offset")
	}

	// Read length
	length := new(big.Int).SetBytes(data[offset : offset+32]).Uint64()
	if offset+32+length > uint64(len(data)) {
		return "", fmt.Errorf("invalid string length")
	}

	// Read string data
	return string(data[offset+32 : offset+32+length]), nil
}

// DecodeAddress decodes a single ABI-encoded address.
func DecodeAddress(data []byte) (string, error) {
	if len(data) < 32 {
		return "", fmt.Errorf("data too short for address: %d bytes", len(data))
	}
	addr := data[12:32]
	return "0x" + hex.EncodeToString(addr), nil
}

// DecodeUint64 decodes a single ABI-encoded uint64.
func DecodeUint64(data []byte) (uint64, error) {
	if len(data) < 32 {
		return 0, fmt.Errorf("data too short for uint64: %d bytes", len(data))
	}
	return new(big.Int).SetBytes(data[:32]).Uint64(), nil
}

// DecodeUint64FromBytes decodes a uint64 from a 32-byte slice (no error return).
func DecodeUint64FromBytes(data []byte) uint64 {
	if len(data) < 32 {
		return 0
	}
	return new(big.Int).SetBytes(data[:32]).Uint64()
}

// DecodeUint256 decodes a single ABI-encoded uint256 as a hex string.
func DecodeUint256(data []byte) (string, error) {
	if len(data) < 32 {
		return "", fmt.Errorf("data too short for uint256: %d bytes", len(data))
	}
	n := new(big.Int).SetBytes(data[:32])
	return "0x" + n.Text(16), nil
}

// DecodeBool decodes a single ABI-encoded boolean.
func DecodeBool(data []byte) (bool, error) {
	if len(data) < 32 {
		return false, fmt.Errorf("data too short for bool: %d bytes", len(data))
	}
	return data[31] != 0, nil
}

// EncodeUint64 encodes a uint64 as a 32-byte ABI value.
func EncodeUint64(v uint64) []byte {
	b := new(big.Int).SetUint64(v).Bytes()
	padded := make([]byte, 32)
	copy(padded[32-len(b):], b)
	return padded
}

// EncodeAddress encodes an address as a 32-byte ABI value.
func EncodeAddress(addr []byte) []byte {
	padded := make([]byte, 32)
	if len(addr) == 20 {
		copy(padded[12:], addr)
	} else {
		copy(padded[32-len(addr):], addr)
	}
	return padded
}

// GetOwner fetches the owner address of a contract.
func GetOwner(ctx *views.ViewContext) (string, error) {
	data, err := ExecuteCall(ctx, SelectorOwner)
	if err != nil {
		return "", fmt.Errorf("owner() failed: %w", err)
	}
	return DecodeAddress(data)
}

// GetTypeAndVersion fetches the typeAndVersion string of a contract.
func GetTypeAndVersion(ctx *views.ViewContext) (string, error) {
	data, err := ExecuteCall(ctx, SelectorTypeAndVersion)
	if err != nil {
		return "", fmt.Errorf("typeAndVersion() failed: %w", err)
	}
	return DecodeString(data)
}
