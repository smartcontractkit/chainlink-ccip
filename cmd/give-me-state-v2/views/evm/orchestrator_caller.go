package evm

import (
	"context"
	"math/big"

	"give-me-state-v2/orchestrator"
	"give-me-state-v2/views"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
)

// OrchestratorCaller implements bind.ContractCaller using the TypedOrchestrator.
type OrchestratorCaller struct {
	TypedOrchestrator orchestrator.TypedOrchestratorInterface
	ChainSelector     uint64
}

// NewOrchestratorCaller creates a new OrchestratorCaller for the given chain.
func NewOrchestratorCaller(orc orchestrator.TypedOrchestratorInterface, chainSelector uint64) *OrchestratorCaller {
	return &OrchestratorCaller{TypedOrchestrator: orc, ChainSelector: chainSelector}
}

// CodeAt returns the code at the given contract address.
func (o *OrchestratorCaller) CodeAt(ctx context.Context, contract common.Address, blockNumber *big.Int) ([]byte, error) {
	return nil, nil
}

// CallContract executes an eth_call through the TypedOrchestrator.
func (o *OrchestratorCaller) CallContract(ctx context.Context, call ethereum.CallMsg, blockNumber *big.Int) ([]byte, error) {
	var targetBytes []byte
	if call.To != nil {
		targetBytes = call.To.Bytes()
	} else {
		targetBytes = make([]byte, 20)
	}
	result := o.TypedOrchestrator.Execute(orchestrator.Call{
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
func OrchestratorCallerFromContext(ctx *views.ViewContext) *OrchestratorCaller {
	return &OrchestratorCaller{
		TypedOrchestrator: ctx.TypedOrchestrator,
		ChainSelector:     ctx.ChainSelector,
	}
}
