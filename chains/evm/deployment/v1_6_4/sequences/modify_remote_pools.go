package sequences

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_4/operations/token_pool"
	token_pool_ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_4/operations/token_pool"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
)

const (
	AddRemotePoolOperation    = "ADD_REMOTE_POOL"
	RemoveRemotePoolOperation = "REMOVE_REMOTE_POOL"
)

// The input is structured to use the same format for both adding and removing remote pools. It also allows for
// performing operations on multiple addresses for a given chain selector.
type ModifyRemotePoolsSequenceInput struct {
	AddressesByChain     map[uint64][]common.Address
	ModificationsByChain map[uint64]map[common.Address]token_pool.RemotePoolModification
}

var ModifyRemotePoolsSequence = operations.NewSequence(
	"ModifyRemotePoolsSequence",
	token_pool.Version,
	"Modifies remote pools in the remote chain config for a sequence of TokenPool contracts on multiple chains",
	func(b operations.Bundle, chains cldf_chain.BlockChains, input ModifyRemotePoolsSequenceInput) (sequences.OnChainOutput, error) {
		writes := make([]contract.WriteOutput, 0)
		for chainSel, addresses := range input.AddressesByChain {
			chain, ok := chains.EVMChains()[chainSel]
			if !ok {
				return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not defined", chainSel)
			}
			for _, address := range addresses {
				modification := input.ModificationsByChain[chainSel][address]

				// Get the operation based on the operation type
				var operation *operations.Operation[contract.FunctionInput[token_pool_ops.RemotePoolModification], contract.WriteOutput, evm.Chain] = nil
				switch modification.Operation {
				case AddRemotePoolOperation:
					operation = token_pool_ops.AddRemotePool
				case RemoveRemotePoolOperation:
					operation = token_pool_ops.RemoveRemotePool
				default:
					return sequences.OnChainOutput{}, fmt.Errorf("invalid operation: %s", modification.Operation)
				}

				// Execute the operation on the chain with the input being the remote pool modification
				// The function signatures have the same input parameters, so the only difference is the operation type.
				// Therefore it is safe to perform this operation as "Args" does not need to be changed.
				report, err := operations.ExecuteOperation(b, operation, chain, contract.FunctionInput[token_pool_ops.RemotePoolModification]{
					ChainSelector: chain.Selector,
					Address:       address,
					Args:          modification,
				})
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to execute AddRemotePoolOp on %s: %w", chain, err)
				}
				writes = append(writes, report.Output)
			}
		}
		batch, err := contract.NewBatchOperationFromWrites(writes)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to create batch operation from writes: %w", err)
		}
		return sequences.OnChainOutput{BatchOps: []mcms_types.BatchOperation{batch}}, nil
	})
