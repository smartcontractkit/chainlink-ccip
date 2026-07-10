package sequences

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_2_0/router"
	evmops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations2/contract"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"

	routerops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/router"
)

type RouterApplyRampUpdatesSequenceInput struct {
	Address        common.Address
	ChainSelector  uint64
	UpdatesByChain routerops.ApplyRampsUpdatesArgs
}

var (
	RouterApplyRampUpdatesSequence = operations.NewSequence(
		"RouterApplyRampUpdatesSequence",
		semver.MustParse("1.6.0"),
		"Updates OnRamps and OffRamps on Router contracts across multiple EVM chains",
		func(b operations.Bundle, chains cldf_chain.BlockChains, input RouterApplyRampUpdatesSequenceInput) (sequences.OnChainOutput, error) {
			writes := make([]contract.WriteOutput, 0)
			chain, ok := chains.EVMChains()[input.ChainSelector]
			if !ok {
				return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not defined", input.ChainSelector)
			}
			report, err := evmops.ExecuteWrite(b, chain, input.Address, router.NewRouter, routerops.NewWriteApplyRampUpdates, input.UpdatesByChain)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to execute RouterApplyRampUpdatesOp on %s: %w", chain, err)
			}
			writes = append(writes, report.Output)
			batch, err := contract.NewBatchOperationFromWrites(writes)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to create batch operation from writes: %w", err)
			}
			return sequences.OnChainOutput{
				BatchOps: []mcms_types.BatchOperation{batch},
			}, nil
		})
)
