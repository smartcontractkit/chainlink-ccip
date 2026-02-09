package views

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"
)

// ParseAddress parses a hex address string to bytes.
func ParseAddress(addr string) ([]byte, error) {
	addr = strings.TrimPrefix(addr, "0x")
	return hex.DecodeString(addr)
}

// FormatAddress formats address bytes as a hex string with 0x prefix.
func FormatAddress(addr []byte) string {
	return "0x" + hex.EncodeToString(addr)
}

// ABIEncodeCall creates calldata from a function selector and arguments.
func ABIEncodeCall(selector []byte, args ...[]byte) []byte {
	result := make([]byte, len(selector))
	copy(result, selector)
	for _, arg := range args {
		padded := make([]byte, 32)
		copy(padded[32-len(arg):], arg)
		result = append(result, padded...)
	}
	return result
}

// ABIDecodeAddress decodes a 32-byte ABI-encoded address to 20 bytes.
func ABIDecodeAddress(data []byte) []byte {
	if len(data) < 32 {
		return nil
	}
	return data[12:32]
}

// ABIDecodeUint256 decodes a 32-byte big-endian uint256.
func ABIDecodeUint256(data []byte) []byte {
	if len(data) < 32 {
		return nil
	}
	return data[:32]
}

// ABIDecodeBool decodes a 32-byte ABI-encoded boolean.
func ABIDecodeBool(data []byte) bool {
	if len(data) < 32 {
		return false
	}
	return data[31] != 0
}

// HexToBytes converts a hex string (with or without 0x prefix) to bytes.
func HexToBytes(h string) []byte {
	h = strings.TrimPrefix(h, "0x")
	b, _ := hex.DecodeString(h)
	return b
}

// BytesToHex converts bytes to a hex string with 0x prefix.
func BytesToHex(b []byte) string {
	return "0x" + hex.EncodeToString(b)
}

// Uint64ToString converts a uint64 to a string.
func Uint64ToString(v uint64) string {
	return fmt.Sprintf("%d", v)
}

// AptosCallType indicates the type of Aptos API call
type AptosCallType byte

const (
	AptosCallResources AptosCallType = 0
	AptosCallResource  AptosCallType = 1
	AptosCallView      AptosCallType = 2
)

// AptosViewCall represents a view function call
type AptosViewCall struct {
	Function      string   `json:"function"`
	TypeArguments []string `json:"type_arguments"`
	Arguments     []any    `json:"arguments"`
}

// EncodeAptosResourcesCall creates calldata to fetch all resources for an account
func EncodeAptosResourcesCall() []byte {
	return []byte{byte(AptosCallResources)}
}

// EncodeAptosResourceCall creates calldata to fetch a specific resource
func EncodeAptosResourceCall(resourceType string) []byte {
	data := make([]byte, 1+len(resourceType))
	data[0] = byte(AptosCallResource)
	copy(data[1:], resourceType)
	return data
}

// EncodeAptosViewCall creates calldata to execute a view function
func EncodeAptosViewCall(moduleAddr, moduleName, funcName string, typeArgs []string, args []any) ([]byte, error) {
	call := AptosViewCall{
		Function:      fmt.Sprintf("%s::%s::%s", moduleAddr, moduleName, funcName),
		TypeArguments: typeArgs,
		Arguments:     args,
	}
	if call.TypeArguments == nil {
		call.TypeArguments = []string{}
	}
	if call.Arguments == nil {
		call.Arguments = []any{}
	}
	jsonData, err := json.Marshal(call)
	if err != nil {
		return nil, err
	}
	data := make([]byte, 1+len(jsonData))
	data[0] = byte(AptosCallView)
	copy(data[1:], jsonData)
	return data, nil
}

// ExecuteAptosView constructs and executes an Aptos view call via TypedOrchestrator.
func ExecuteAptosView(ctx *ViewContext, moduleAddr, moduleName, funcName string, typeArgs []string, args []any) ([]byte, error) {
	callData, err := EncodeAptosViewCall(moduleAddr, moduleName, funcName, typeArgs, args)
	if err != nil {
		return nil, fmt.Errorf("failed to encode view call: %w", err)
	}
	result := ctx.TypedOrchestrator.Execute(Call{
		ChainID: ctx.ChainSelector,
		Target:  ctx.Address,
		Data:    callData,
	})
	if result.Error != nil {
		return nil, result.Error
	}
	return result.Data, nil
}

// ExecuteAptosViewOnAddress executes a view call on a specific address (not ctx.Address)
func ExecuteAptosViewOnAddress(ctx *ViewContext, targetAddr, moduleAddr, moduleName, funcName string, typeArgs []string, args []any) ([]byte, error) {
	callData, err := EncodeAptosViewCall(moduleAddr, moduleName, funcName, typeArgs, args)
	if err != nil {
		return nil, fmt.Errorf("failed to encode view call: %w", err)
	}
	result := ctx.TypedOrchestrator.Execute(Call{
		ChainID: ctx.ChainSelector,
		Target:  []byte(targetAddr),
		Data:    callData,
	})
	if result.Error != nil {
		return nil, result.Error
	}
	return result.Data, nil
}
