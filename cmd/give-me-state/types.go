package main

import (
	"crypto/sha256"
	"encoding/binary"
)

// Call represents a chain-agnostic call to be executed.
// Works for EVM, Solana, Aptos, or any other chain type.
type Call struct {
	ChainID uint64 // Chain selector/identifier
	Target  []byte // Address (20 bytes EVM, 32 bytes Solana, etc.)
	Data    []byte // Calldata (ABI-encoded for EVM, Borsh for Solana, etc.)
}

// CacheKey generates a unique key for caching based on chain, target, and data.
// Returns a SHA256 hash that uniquely identifies this call.
func (c Call) CacheKey() [32]byte {
	h := sha256.New()
	_ = binary.Write(h, binary.BigEndian, c.ChainID)
	h.Write(c.Target)
	h.Write(c.Data)
	var key [32]byte
	copy(key[:], h.Sum(nil))
	return key
}

// CallResult contains the result of executing a Call.
type CallResult struct {
	Data    []byte // Raw result bytes (nil if error)
	Error   error  // nil if success
	Cached  bool   // Was this a cache hit?
	Retries int    // How many retries were needed
}

// ChainExecutor is the interface that chain-specific executors must implement.
// Each chain type (EVM, Solana, Aptos) provides its own implementation.
type ChainExecutor interface {
	// Execute performs the actual RPC call to the chain.
	// target: the contract/program address
	// data: the encoded call data
	// Returns the raw response bytes or an error.
	Execute(target, data []byte) ([]byte, error)
}

// Stats holds statistics about the CallManager's operation.
type Stats struct {
	TotalCalls   int64
	CacheHits    int64
	DedupedCalls int64
	TotalRetries int64
	Errors       int64
}
