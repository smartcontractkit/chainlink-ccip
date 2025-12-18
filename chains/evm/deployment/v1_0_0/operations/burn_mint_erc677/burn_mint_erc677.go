package burn_mint_erc677

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-evm/gethwrappers/shared/generated/initial/burn_mint_erc677"
)

var ContractType cldf_deployment.ContractType = "BurnMintERC677Token"

type ConstructorArgs struct {
	Name      string
	Symbol    string
	Decimals  uint8
	MaxSupply *big.Int
}

var Deploy = contract.NewDeploy(contract.DeployParams[ConstructorArgs]{
	Name:             "burn_mint_erc677:deploy",
	Version:          utils.Version_1_0_0,
	Description:      "Deploys the BurnMintERC677 Token contract",
	ContractMetadata: burn_mint_erc677.BurnMintERC677MetaData,
	BytecodeByTypeAndVersion: map[string]contract.Bytecode{
		cldf_deployment.NewTypeAndVersion(ContractType, *utils.Version_1_0_0).String(): {
			EVM: common.FromHex(burn_mint_erc677.BurnMintERC677Bin),
		},
	},
	Validate: func(args ConstructorArgs) error { return nil },
})
