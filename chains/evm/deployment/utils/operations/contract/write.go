package contract

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	cldf_evm "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	upstream "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations/contract"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

// ExecInfo contains information about an executed transaction.
//
// Deprecated: Import github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations/contract directly.
type ExecInfo = upstream.ExecInfo

// WriteOutput is the output of a write operation.
//
// Deprecated: Import github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations/contract directly.
type WriteOutput = upstream.WriteOutput

// WriteParams contains the parameters to create a write operation.
//
// Deprecated: Import github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations/contract directly.
type WriteParams[ARGS any, C any] = upstream.WriteParams[ARGS, C]

// NewWrite creates a new write operation.
//
// Deprecated: Import github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations/contract directly.
func NewWrite[ARGS any, C any](params WriteParams[ARGS, C]) *operations.Operation[FunctionInput[ARGS], WriteOutput, cldf_evm.Chain] {
	return upstream.NewWrite(params)
}

// OnlyOwner checks if the caller is the owner of the contract.
//
// Deprecated: Import github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations/contract directly.
func OnlyOwner[C interface {
	Address() common.Address
	Owner(opts *bind.CallOpts) (common.Address, error)
}, ARGS any](contract C, opts *bind.CallOpts, caller common.Address, args ARGS) (bool, error) {
	return upstream.OnlyOwner(contract, opts, caller, args)
}

// AllCallersAllowed always returns true, allowing any caller to execute directly.
//
// Deprecated: Import github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations/contract directly.
func AllCallersAllowed[C any, ARGS any](contract C, opts *bind.CallOpts, caller common.Address, args ARGS) (bool, error) {
	return upstream.AllCallersAllowed(contract, opts, caller, args)
}

// NoCallersAllowed always returns false, forcing the write into an MCMS proposal.
//
// Deprecated: Import github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations/contract directly.
func NoCallersAllowed[C any, ARGS any](contract C, opts *bind.CallOpts, caller common.Address, args ARGS) (bool, error) {
	return upstream.NoCallersAllowed(contract, opts, caller, args)
}

// NewBatchOperationFromWrites constructs an MCMS BatchOperation from a slice of WriteOutputs.
//
// Deprecated: Import github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations/contract directly.
var NewBatchOperationFromWrites = upstream.NewBatchOperationFromWrites
