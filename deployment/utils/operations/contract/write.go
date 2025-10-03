package contract

import (
	mcms_types "github.com/smartcontractkit/mcms/types"
)

// ExecInfo contains information about an executed transaction.
// Defined as a struct in case we want to add more fields in the future without breaking existing usage.
type ExecInfo struct {
	// Hash is the transaction hash.
	Hash string
}

// WriteOutput is the output of a write operation.
type WriteOutput struct {
	// ChainSelector is the selector of the target chain.
	ChainSelector uint64 `json:"chainSelector"`
	// Tx is the prepared transaction (in MCMS format).
	Tx mcms_types.Transaction `json:"tx"`
	// ExecInfo is populated if the write was executed, contains info about the executed transaction.
	ExecInfo *ExecInfo `json:"execInfo,omitempty"`
}

func (o WriteOutput) Executed() bool {
	return o.ExecInfo != nil
}
