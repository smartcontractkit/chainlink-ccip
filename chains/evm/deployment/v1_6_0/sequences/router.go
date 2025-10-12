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

	routerops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/router"
)

type RouterApplyRampUpdatesSequenceInput struct {
	Address        common.Address
	UpdatesByChain map[uint64]routerops.ApplyRampsUpdatesArgs
}

var (
	RouterApplyRampUpdatesSequence = operations.NewSequence(
		"RouterApplyRampUpdatesSequence",
		semver.MustParse("1.6.0"),
		"Updates OnRamps and OffRamps on Router contracts across multiple EVM chains",
		func(b operations.Bundle, chains cldf_chain.BlockChains, input RouterApplyRampUpdatesSequenceInput) (sequences.OnChainOutput, error) {
			writes := make([]contract.WriteOutput, 0)
			for chainSel, update := range input.UpdatesByChain {
				chain, ok := chains.EVMChains()[chainSel]
				if !ok {
					return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not defined", chainSel)
				}
				report, err := operations.ExecuteOperation(b, routerops.ApplyRampUpdates, chain, contract.FunctionInput[routerops.ApplyRampsUpdatesArgs]{
					ChainSelector: chain.Selector,
					Address:       input.Address,
					Args:          update,
				})
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to execute RouterApplyRampUpdatesOp on %s: %w", chain, err)
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
