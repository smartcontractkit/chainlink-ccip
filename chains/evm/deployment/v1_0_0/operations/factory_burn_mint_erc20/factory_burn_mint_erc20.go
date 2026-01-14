package factory_burn_mint_erc20

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/latest/factory_burn_mint_erc20"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

var ContractType cldf_deployment.ContractType = "FactoryBurnMintERC20Token"

type ConstructorArgs struct {
	Name      string
	Symbol    string
	Decimals  uint8
	MaxSupply *big.Int
	PreMint   *big.Int
	NewOwner  common.Address
}

var Deploy = contract.NewDeploy(contract.DeployParams[ConstructorArgs]{
	Name:             "factory_burn_mint_erc20:deploy",
	Version:          utils.Version_1_0_0,
	Description:      "Deploys the FactoryBurnMintERC20 token contract",
	ContractMetadata: factory_burn_mint_erc20.FactoryBurnMintERC20MetaData,
	BytecodeByTypeAndVersion: map[string]contract.Bytecode{
		cldf_deployment.NewTypeAndVersion(ContractType, *utils.Version_1_0_0).String(): {
			EVM: common.FromHex(factory_burn_mint_erc20.FactoryBurnMintERC20Bin),
		},
	},
	Validate: func(args ConstructorArgs) error { return nil },
})
