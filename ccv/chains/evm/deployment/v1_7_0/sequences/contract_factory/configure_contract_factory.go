package contract_factory

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/contract_factory"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"
)

type ConfigureContractFactoryInput struct {
	ChainSelector          uint64
	ContractFactoryAddress common.Address
	AllowListAdds          []common.Address
	AllowListRemoves       []common.Address
}

var ConfigureContractFactory = cldf_ops.NewSequence(
	"configure-contract-factory",
	semver.MustParse("1.7.0"),
	"Configures the ContractFactory contract",
	func(b operations.Bundle, chain evm.Chain, input ConfigureContractFactoryInput) (output sequences.OnChainOutput, err error) {
		writes := make([]contract.WriteOutput, 0)
		configureReport, err := cldf_ops.ExecuteOperation(b, contract_factory.ApplyAllowListUpdates, chain, contract.FunctionInput[contract_factory.ApplyAllowListUpdatesArgs]{
			ChainSelector: chain.Selector,
			Address:       input.ContractFactoryAddress,
			Args: contract_factory.ApplyAllowListUpdatesArgs{
				Adds:    input.AllowListAdds,
				Removes: input.AllowListRemoves,
			},
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to configure ContractFactory: %w", err)
		}
		writes = append(writes, configureReport.Output)

		batch, err := contract.NewBatchOperationFromWrites(writes)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to create batch operation from writes: %w", err)
		}
		return sequences.OnChainOutput{
			BatchOps: []mcms_types.BatchOperation{batch},
		}, nil
	},
)
