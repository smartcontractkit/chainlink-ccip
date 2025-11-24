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

type ConfigureAllowedCallersSequenceInput struct {
	AddressesByChain               map[uint64]common.Address
	ConfigureAllowedCallersByChain map[uint64][]erc20_lock_box_ops.AllowedCallerConfigArgs
}

var (
	ERC20LockboxConfigureAllowedCallersSequence = operations.NewSequence(
		"ERC20LockboxConfigureAllowedCallersSequence",
		semver.MustParse("1.6.4"),
		"Configures allowed callers on a sequence of ERC20Lockbox contracts on multiple chains",
		func(b operations.Bundle, chains cldf_chain.BlockChains, input ConfigureAllowedCallersSequenceInput) (sequences.OnChainOutput, error) {
			writes := make([]contract.WriteOutput, 0)

			// Iterate over each chain selector in the input
			for chainSel, allowedCallers := range input.ConfigureAllowedCallersByChain {
				address, ok := input.AddressesByChain[chainSel]
				if !ok {
					return sequences.OnChainOutput{}, fmt.Errorf("address not found for chain selector %d", chainSel)
				}

				// Get the chain object based on the chain selector
				chain, ok := chains.EVMChains()[chainSel]
				if !ok {
					return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not defined", chainSel)
				}

				// Execute the operation ERC20LockboxConfigureAllowedCallers, on "chain", with the input being an array of
				// AllowedCallerConfigArgs structs, with the first and only item being the allowed callers for the given chain selector
				report, err := operations.ExecuteOperation(b, erc20_lock_box_ops.ERC20LockboxConfigureAllowedCallers, chain, contract.FunctionInput[[]erc20_lock_box_ops.AllowedCallerConfigArgs]{
					ChainSelector: chain.Selector,
					Address:       address,
					Args:          allowedCallers,
				})
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to execute ERC20LockboxConfigureAllowedCallersOp on %s: %w", chain, err)
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
