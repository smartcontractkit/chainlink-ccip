package views

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"
	"sync"
)

// ViewFunc is the interface that all contract view generators must implement.
// It fetches the current state of a contract and returns it as JSON-serializable data.
type ViewFunc func(ctx *ViewContext) (map[string]any, error)

// ViewContext provides all the context needed to generate a view.
type ViewContext struct {
	// Contract information
	Address       []byte // 20 bytes for EVM, 32 bytes for Solana, etc.
	AddressHex    string // Hex string with 0x prefix
	ChainSelector uint64
	Qualifier     string // e.g., "default", "TEST", "MIXTEST"

	// Call orchestrator for making RPC calls
	CallManager CallManagerInterface

	// Chain family for type-specific logic
	ChainFamily string // "evm", "svm", "aptos", "ton"

	// All chain selectors in the deployment (for discovering remote chain configs)
	AllChainSelectors []uint64

	// Optional chain-specific clients (to avoid import cycles, use any and type assert)
	// For Aptos: *aptos.Client from github.com/aptos-labs/aptos-go-sdk
	AptosClient any
}

// CallManagerInterface is the interface the views package needs from the orchestrator.
// This avoids circular imports.
type CallManagerInterface interface {
	Execute(call Call) CallResult
}

// Call represents a chain-agnostic call (matches main package)
type Call struct {
	ChainID uint64
	Target  []byte
	Data    []byte
}

// CallResult represents the result of a call (matches main package)
type CallResult struct {
	Data    []byte
	Error   error
	Cached  bool
	Retries int
}

// =====================================================
// Registry for View Functions
// =====================================================

// ViewKey uniquely identifies a view function by type and version.
type ViewKey struct {
	ChainFamily string // "evm", "svm", "aptos", "ton"
	Type        string // "FeeQuoter", "Router", "OnRamp", etc.
	Version     string // "1.6.0", "1.7.0", etc.
}

func (k ViewKey) String() string {
	return fmt.Sprintf("%s/%s@%s", k.ChainFamily, k.Type, k.Version)
}

// Registry holds all registered view functions.
type Registry struct {
	views map[ViewKey]ViewFunc
	mu    sync.RWMutex
}

// Global registry instance
var globalRegistry = &Registry{
	views: make(map[ViewKey]ViewFunc),
}

// Register adds a view function to the global registry.
func Register(chainFamily, contractType, version string, fn ViewFunc) {
	key := ViewKey{
		ChainFamily: chainFamily,
		Type:        contractType,
		Version:     version,
	}
	globalRegistry.mu.Lock()
	defer globalRegistry.mu.Unlock()
	globalRegistry.views[key] = fn
}

// Get retrieves a view function from the global registry.
func Get(chainFamily, contractType, version string) (ViewFunc, bool) {
	key := ViewKey{
		ChainFamily: chainFamily,
		Type:        contractType,
		Version:     version,
	}
	globalRegistry.mu.RLock()
	defer globalRegistry.mu.RUnlock()
	fn, ok := globalRegistry.views[key]
	return fn, ok
}

// GetWithFallback tries exact version, then falls back to compatible versions.
// e.g., for version "1.7.0", it might try "1.7.0", then "1.7", then "1.6.0"
func GetWithFallback(chainFamily, contractType, version string) (ViewFunc, string, bool) {
	// Try exact match first
	if fn, ok := Get(chainFamily, contractType, version); ok {
		return fn, version, true
	}

	// Try minor version wildcard (e.g., "1.7.0" -> "1.7")
	parts := strings.Split(version, ".")
	if len(parts) >= 2 {
		minorVersion := parts[0] + "." + parts[1]
		if fn, ok := Get(chainFamily, contractType, minorVersion); ok {
			return fn, minorVersion, true
		}
	}

	// Try finding any version for this type (for fallback warning purposes)
	globalRegistry.mu.RLock()
	defer globalRegistry.mu.RUnlock()
	for key := range globalRegistry.views {
		if key.ChainFamily == chainFamily && key.Type == contractType {
			return globalRegistry.views[key], key.Version, true
		}
	}

	return nil, "", false
}

// ListRegistered returns all registered view keys.
func ListRegistered() []ViewKey {
	globalRegistry.mu.RLock()
	defer globalRegistry.mu.RUnlock()

	keys := make([]ViewKey, 0, len(globalRegistry.views))
	for k := range globalRegistry.views {
		keys = append(keys, k)
	}
	return keys
}

// IsSupported checks if a specific chain family, type, and version is supported.
func IsSupported(chainFamily, contractType, version string) bool {
	_, ok := Get(chainFamily, contractType, version)
	if ok {
		return true
	}
	// Try fallback
	_, _, ok = GetWithFallback(chainFamily, contractType, version)
	return ok
}

// =====================================================
// Helper Functions for View Implementations
// =====================================================

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
// For now, only supports calls with no arguments or simple types.
func ABIEncodeCall(selector []byte, args ...[]byte) []byte {
	result := make([]byte, len(selector))
	copy(result, selector)
	for _, arg := range args {
		// Pad to 32 bytes
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
	return data[12:32] // Skip 12 zero bytes
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

// =====================================================
// Aptos Call Helpers
// =====================================================

// AptosCallType indicates the type of Aptos API call
type AptosCallType byte

const (
	// AptosCallResources fetches all resources for an account
	AptosCallResources AptosCallType = 0
	// AptosCallResource fetches a specific resource
	AptosCallResource AptosCallType = 1
	// AptosCallView executes a view function
	AptosCallView AptosCallType = 2
)

// AptosViewCall represents a view function call
type AptosViewCall struct {
	Function      string   `json:"function"`       // e.g., "0x1::module::function"
	TypeArguments []string `json:"type_arguments"` // Generic type args
	Arguments     []any    `json:"arguments"`      // Function arguments
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

// ExecuteAptosView is a helper that constructs and executes an Aptos view call
func ExecuteAptosView(ctx *ViewContext, moduleAddr, moduleName, funcName string, typeArgs []string, args []any) ([]byte, error) {
	callData, err := EncodeAptosViewCall(moduleAddr, moduleName, funcName, typeArgs, args)
	if err != nil {
		return nil, fmt.Errorf("failed to encode view call: %w", err)
	}

	result := ctx.CallManager.Execute(Call{
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

	result := ctx.CallManager.Execute(Call{
		ChainID: ctx.ChainSelector,
		Target:  []byte(targetAddr),
		Data:    callData,
	})

	if result.Error != nil {
		return nil, result.Error
	}
	return result.Data, nil
}
