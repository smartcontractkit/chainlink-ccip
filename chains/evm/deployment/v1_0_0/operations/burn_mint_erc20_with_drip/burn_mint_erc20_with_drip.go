package burn_mint_erc20_with_drip

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-evm/gethwrappers/shared/generated/initial/burn_mint_erc20_with_drip"
)

var ContractType cldf_deployment.ContractType = "BurnMintERC20WithDripToken"

type ConstructorArgs struct {
	Name      string
	Symbol    string
	Decimals  uint8
	MaxSupply *big.Int
	PreMint   *big.Int
}

var Deploy = contract.NewDeploy(contract.DeployParams[ConstructorArgs]{
	Name:             "burn_mint_erc20_with_drip:deploy",
	Version:          utils.Version_1_0_0,
	Description:      "Deploys the BurnMintERC20 Token contract with drip function (for testing)",
	ContractMetadata: burn_mint_erc20_with_drip.BurnMintERC20MetaData,
	BytecodeByTypeAndVersion: map[string]contract.Bytecode{
		cldf_deployment.NewTypeAndVersion(ContractType, *utils.Version_1_0_0).String(): {
			EVM: common.FromHex(burn_mint_erc20_with_drip.BurnMintERC20Bin),
		},
	},
	Validate: func(args ConstructorArgs) error { return nil },
})
