package sequences

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_4/operations/token_pool"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"

	token_pool_ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_4/operations/token_pool"
)

type ApplyChainUpdatesSequenceInput struct {
	AddressesByChain map[uint64][]common.Address
	UpdatesByChain   map[uint64]map[common.Address]token_pool_ops.ApplyChainUpdatesArgs
}

var (
	TokenPoolApplyChainUpdatesSequence = operations.NewSequence(
		"TokenPoolApplyChainUpdatesSequence",
		token_pool.Version,
		"Applies chain updates to a sequence of TokenPool contracts on multiple chains",
		func(b operations.Bundle, chains cldf_chain.BlockChains, input ApplyChainUpdatesSequenceInput) (sequences.OnChainOutput, error) {
			writes := make([]contract.WriteOutput, 0)
			// For each chain selector in the address list
			for chainSel, addresses := range input.AddressesByChain {
				chain, ok := chains.EVMChains()[chainSel]
				if !ok {
					return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not defined", chainSel)
				}

				// For each address for a given chain selector
				for _, address := range addresses {
					// Get the chain update for the given address and chain selector
					chainUpdate := input.UpdatesByChain[chainSel][address]

					// Execute the operation ApplyChainUpdates
					report, err := operations.ExecuteOperation(b, token_pool_ops.ApplyChainUpdates, chain, contract.FunctionInput[token_pool_ops.ApplyChainUpdatesArgs]{
						ChainSelector: chain.Selector,
						Address:       address,
						Args:          chainUpdate,
					})
					if err != nil {
						return sequences.OnChainOutput{}, fmt.Errorf("failed to execute TokenPoolApplyChainUpdatesOp on %s: %w", chain, err)
					}
					writes = append(writes, report.Output)
				}
			}
			batch, err := contract.NewBatchOperationFromWrites(writes)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to create batch operation from writes: %w", err)
			}
			return sequences.OnChainOutput{
				BatchOps: []mcms_types.BatchOperation{batch},
			}, nil
		})
)
