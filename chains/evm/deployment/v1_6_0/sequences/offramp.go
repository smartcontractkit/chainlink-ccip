package sequences

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_0/offramp"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"

	offrampops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/offramp"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	mcms_types "github.com/smartcontractkit/mcms/types"
)

type OffRampApplySourceChainConfigUpdatesSequenceInput struct {
	Address        common.Address
	UpdatesByChain map[uint64][]offramp.OffRampSourceChainConfigArgs
}

var (
	OffRampApplySourceChainConfigUpdatesSequence = operations.NewSequence(
		"OffRampApplySourceChainConfigUpdatesSequence",
		semver.MustParse("1.6.0"),
		"Applies updates to source chain configurations stored on OffRamp contracts on multiple EVM chains",
		func(b operations.Bundle, chains cldf_chain.BlockChains, input OffRampApplySourceChainConfigUpdatesSequenceInput) (sequences.OnChainOutput, error) {
			writes := make([]contract.WriteOutput, 0)
			for chainSel, update := range input.UpdatesByChain {
				chain, ok := chains.EVMChains()[chainSel]
				if !ok {
					return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not defined", chainSel)
				}
				report, err := operations.ExecuteOperation(b, offrampops.OffRampApplySourceChainConfigUpdates, chain, contract.FunctionInput[[]offramp.OffRampSourceChainConfigArgs]{
					ChainSelector: chain.Selector,
					Address:       input.Address,
					Args:          update,
				})
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to execute OffRampApplySourceChainConfigUpdatesOp on %s: %w", chain, err)
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
