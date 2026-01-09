package erc20

import (
	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-evm/gethwrappers/shared/generated/initial/erc20"
)

var ContractType cldf_deployment.ContractType = "ERC20Token"

type ConstructorArgs struct {
	Name   string
	Symbol string
}

var Deploy = contract.NewDeploy(contract.DeployParams[ConstructorArgs]{
	Name:             "erc20:deploy",
	Version:          utils.Version_1_0_0,
	Description:      "Deploys the ERC20 Token contract",
	ContractMetadata: erc20.ERC20MetaData,
	BytecodeByTypeAndVersion: map[string]contract.Bytecode{
		cldf_deployment.NewTypeAndVersion(ContractType, *utils.Version_1_0_0).String(): {
			EVM: common.FromHex(erc20.ERC20Bin),
		},
	},
	Validate: func(args ConstructorArgs) error { return nil },
})
