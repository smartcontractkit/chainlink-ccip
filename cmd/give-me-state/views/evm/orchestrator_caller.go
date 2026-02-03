package evm

import (
	"context"
	"math/big"

	"call-orchestrator-demo/views"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
)

// OrchestratorCaller implements bind.ContractCaller using the CallManager.
// This allows go-ethereum generated bindings to use the orchestrator for RPC calls,
// benefiting from caching, deduplication, and retry logic.
type OrchestratorCaller struct {
	CallManager   views.CallManagerInterface
	ChainSelector uint64
}

// NewOrchestratorCaller creates a new OrchestratorCaller for the given chain.
func NewOrchestratorCaller(mgr views.CallManagerInterface, chainSelector uint64) *OrchestratorCaller {
	return &OrchestratorCaller{
		CallManager:   mgr,
		ChainSelector: chainSelector,
	}
}

// CodeAt returns the code at the given contract address.
// This is required by bind.ContractCaller but not typically needed for read-only view calls.
// We return nil which works for most contract binding use cases.
func (o *OrchestratorCaller) CodeAt(ctx context.Context, contract common.Address, blockNumber *big.Int) ([]byte, error) {
	// Not implemented - not needed for read-only contract calls
	// Returning nil tells the binding library the contract exists
	return nil, nil
}

// CallContract executes an eth_call through the CallManager.
// This routes all contract calls through the orchestrator's caching/retry infrastructure.
func (o *OrchestratorCaller) CallContract(ctx context.Context, call ethereum.CallMsg, blockNumber *big.Int) ([]byte, error) {
	// Build target address bytes (20 bytes for EVM)
	var targetBytes []byte
	if call.To != nil {
		targetBytes = call.To.Bytes()
	} else {
		// No target address - invalid for a call
		targetBytes = make([]byte, 20)
	}

	// Execute through the orchestrator
	result := o.CallManager.Execute(views.Call{
		ChainID: o.ChainSelector,
		Target:  targetBytes,
		Data:    call.Data,
	})

	if result.Error != nil {
		return nil, result.Error
	}

	return result.Data, nil
}

// OrchestratorCallerFromContext creates an OrchestratorCaller from a ViewContext.
// This is a convenience function for use within view functions.
func OrchestratorCallerFromContext(ctx *views.ViewContext) *OrchestratorCaller {
	return &OrchestratorCaller{
		CallManager:   ctx.CallManager,
		ChainSelector: ctx.ChainSelector,
	}
}
