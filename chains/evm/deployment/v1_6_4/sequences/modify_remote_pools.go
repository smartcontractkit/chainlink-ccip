package sequences

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	usdc_pool_ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_4/operations/usdc_token_pool_cctp_v2"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
)

const (
	AddRemotePoolOperation    = "ADD_REMOTE_POOL"
	RemoveRemotePoolOperation = "REMOVE_REMOTE_POOL"
)

// RemotePoolModification combines the operation type with the pool parameters
type RemotePoolModification struct {
	Operation           string
	RemoteChainSelector uint64
	RemotePoolAddress   []byte
}

// The input is structured to use the same format for both adding and removing remote pools. It also allows for
// performing operations on multiple addresses for a given chain selector.
type ModifyRemotePoolsSequenceInput struct {
	AddressesByChain     map[uint64][]common.Address
	ModificationsByChain map[uint64]map[common.Address]RemotePoolModification
}

var ModifyRemotePoolsSequence = operations.NewSequence(
	"ModifyRemotePoolsSequence",
	usdc_pool_ops.Version,
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

				// Execute the operation based on the operation type
				var report operations.Report[contract.FunctionInput[usdc_pool_ops.AddRemotePoolArgs], contract.WriteOutput]
				var err error
				switch modification.Operation {
				case AddRemotePoolOperation:
					report, err = operations.ExecuteOperation(b, usdc_pool_ops.AddRemotePool, chain, contract.FunctionInput[usdc_pool_ops.AddRemotePoolArgs]{
						ChainSelector: chain.Selector,
						Address:       address,
						Args: usdc_pool_ops.AddRemotePoolArgs{
							RemoteChainSelector: modification.RemoteChainSelector,
							RemotePoolAddress:   modification.RemotePoolAddress,
						},
					})
				case RemoveRemotePoolOperation:
					removeReport, removeErr := operations.ExecuteOperation(b, usdc_pool_ops.RemoveRemotePool, chain, contract.FunctionInput[usdc_pool_ops.RemoveRemotePoolArgs]{
						ChainSelector: chain.Selector,
						Address:       address,
						Args: usdc_pool_ops.RemoveRemotePoolArgs{
							RemoteChainSelector: modification.RemoteChainSelector,
							RemotePoolAddress:   modification.RemotePoolAddress,
						},
					})
					report.Output = removeReport.Output
					err = removeErr
				default:
					return sequences.OnChainOutput{}, fmt.Errorf("invalid operation: %s", modification.Operation)
				}
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
