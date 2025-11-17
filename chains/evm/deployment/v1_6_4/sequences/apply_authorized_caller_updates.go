package sequences

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"

	authorized_caller_ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_4/operations/authorized_caller"
)

type ApplyAuthorizedCallerUpdatesSequenceInput struct {
	Address                        map[uint64]common.Address
	AuthorizedCallerUpdatesByChain map[uint64]authorized_caller_ops.AuthorizedCallerUpdateArgs
}

var (
	ApplyAuthorizedCallerUpdatesSequence = operations.NewSequence(
		"ApplyAuthorizedCallerUpdatesSequence",
		semver.MustParse("1.6.4"),
		"Applies authorized caller updates to a sequence of contracts implementing the AuthorizedCallers interface on multiple chains",
		func(b operations.Bundle, chains cldf_chain.BlockChains, input ApplyAuthorizedCallerUpdatesSequenceInput) (sequences.OnChainOutput, error) {
			writes := make([]contract.WriteOutput, 0)

			// Iterate over each chain selector in the input
			for chainSel, authorizedCallerUpdate := range input.AuthorizedCallerUpdatesByChain {

				// Get the chain object based on the chain selector
				chain, ok := chains.EVMChains()[chainSel]
				if !ok {
					return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not defined", chainSel)
				}

				// Execute the operation ApplyAuthorizedCallerUpdates, on "chain", with the input being
				// AuthorizedCallerUpdateArgs for the given chain selector
				report, err := operations.ExecuteOperation(b, authorized_caller_ops.ApplyAuthorizedCallerUpdates, chain, contract.FunctionInput[authorized_caller_ops.AuthorizedCallerUpdateArgs]{
					ChainSelector: chain.Selector,
					Address:       input.Address[chainSel],
					Args:          authorizedCallerUpdate,
				})
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to execute ApplyAuthorizedCallerUpdatesOp on %s: %w", chain, err)
				}
				writes = append(writes, report.Output)
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
