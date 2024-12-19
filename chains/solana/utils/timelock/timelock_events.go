package timelock

import (
	"github.com/gagliardetto/solana-go"
)

// Events - temporary event struct to decode
// anchor-go does not support events
// https://github.com/fragmetric-labs/solana-anchor-go does but requires upgrade to anchor >= v0.30.0

// CallScheduled represents an event emitted when a call is scheduled
type CallScheduled struct {
	ID          [32]byte         // id
	Index       uint64           // index
	Target      solana.PublicKey // target
	Predecessor [32]byte         // predecessor
	Salt        [32]byte         // salt
	Delay       uint64           // delay
	Data        []byte           // data: Vec<u8>
}

// CallExecuted represents an event emitted when a call is performed
type CallExecuted struct {
	ID     [32]byte         // id
	Index  uint64           // index
	Target solana.PublicKey // target
	Data   []byte           // data: Vec<u8>
}

// BypasserCallExecuted represents an event emitted when a call is performed via bypasser
type BypasserCallExecuted struct {
	Index  uint64           // index
	Target solana.PublicKey // target
	Data   []byte           // data: Vec<u8>
}

// Cancelled represents an event emitted when an operation is cancelled
type Cancelled struct {
	ID [32]byte // id
}

// MinDelayChange represents an event emitted when the minimum delay is modified
type MinDelayChange struct {
	OldDuration uint64 // old_duration
	NewDuration uint64 // new_duration
}

// FunctionSelectorBlocked represents an event emitted when a function selector is blocked
type FunctionSelectorBlocked struct {
	Selector [8]byte // selector
}

// FunctionSelectorUnblocked represents an event emitted when a function selector is unblocked
type FunctionSelectorUnblocked struct {
	Selector [8]byte // selector
}
