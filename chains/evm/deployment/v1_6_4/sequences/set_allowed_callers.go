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

	erc20_lock_box_ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_4/operations/erc20_lock_box"
)

type SetAllowedCallersSequenceInput struct {
	Address               map[uint64]common.Address
	AllowedCallersByChain map[uint64][]erc20_lock_box_ops.AllowedCallerConfigArgs
}

var (
	ERC20LockboxSetAllowedCallersSequence = operations.NewSequence(
		"ERC20LockboxSetAllowedCallersSequence",
		semver.MustParse("1.6.4"),
		"Sets allowed callers on a sequence of ERC20Lockbox contracts on multiple chains",
		func(b operations.Bundle, chains cldf_chain.BlockChains, input SetAllowedCallersSequenceInput) (sequences.OnChainOutput, error) {
			writes := make([]contract.WriteOutput, 0)

			// Iterate over each chain selector in the input
			for chainSel, allowedCallers := range input.AllowedCallersByChain {

				// Get the chain object based on the chain selector
				chain, ok := chains.EVMChains()[chainSel]
				if !ok {
					return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not defined", chainSel)
				}

				// Execute the operation ERC20LockboxSetAllowedCallers, on "chain", with the input being an array of
				// AllowedCallerConfigArgs structs, with the first and only item being the allowed callers for the given chain selector
				report, err := operations.ExecuteOperation(b, erc20_lock_box_ops.ERC20LockboxSetAllowedCallers, chain, contract.FunctionInput[[]erc20_lock_box_ops.AllowedCallerConfigArgs]{
					ChainSelector: chain.Selector,
					Address:       input.Address[chainSel],
					Args:          allowedCallers,
				})
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to execute ERC20LockboxSetAllowedCallersOp on %s: %w", chain, err)
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
