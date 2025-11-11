package contract_factory

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/contract_factory"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

type DeployContractFactoryInput struct {
	ChainSelector uint64
	AllowList     []common.Address
}

var DeployContractFactory = cldf_ops.NewSequence(
	"deploy-contract-factory",
	semver.MustParse("1.7.0"),
	"Deploys the ContractFactory contract",
	func(b operations.Bundle, chain evm.Chain, input DeployContractFactoryInput) (output sequences.OnChainOutput, err error) {
		deployReport, err := cldf_ops.ExecuteOperation(b, contract_factory.Deploy, chain, contract.DeployInput[contract_factory.ConstructorArgs]{
			ChainSelector: input.ChainSelector,
			Args: contract_factory.ConstructorArgs{
				AllowList: input.AllowList,
			},
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy ContractFactory: %w", err)
		}
		return sequences.OnChainOutput{
			Addresses: []datastore.AddressRef{deployReport.Output},
		}, nil
	},
)
