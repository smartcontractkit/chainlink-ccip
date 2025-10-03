package contract

import (
	"errors"

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

// NewBatchOperation constructs an MCMS BatchOperation from a slice of WriteOutputs.
// It filters out any WriteOutputs that have already been executed.
// Returns an error if the WriteOutputs target multiple chains.
// If all WriteOutputs are executed, it returns an empty BatchOperation and no error.
func NewBatchOperationFromWrites(outs []WriteOutput) (mcms_types.BatchOperation, error) {
	if len(outs) == 0 {
		return mcms_types.BatchOperation{}, nil
	}

	batchOps := make(map[uint64]mcms_types.BatchOperation)
	var chainSelector uint64
	for i, out := range outs {
		if out.Executed() {
			continue // Skip executed transactions, they should not be included.
		}
		if batchOp, exists := batchOps[out.ChainSelector]; !exists {
			if i != 0 {
				return mcms_types.BatchOperation{}, errors.New("failed to make batch operation: writes target multiple chains")
			}
			batchOps[out.ChainSelector] = mcms_types.BatchOperation{
				ChainSelector: mcms_types.ChainSelector(out.ChainSelector),
				Transactions:  []mcms_types.Transaction{out.Tx},
			}
			chainSelector = out.ChainSelector
		} else {
			batchOp.Transactions = append(batchOp.Transactions, out.Tx)
			batchOps[out.ChainSelector] = batchOp
		}
	}

	// If there are no unexecuted writes, return an empty BatchOperation.
	if len(batchOps) == 0 {
		return mcms_types.BatchOperation{}, nil
	}

	return batchOps[chainSelector], nil
}
