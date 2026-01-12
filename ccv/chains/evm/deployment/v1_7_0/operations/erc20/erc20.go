package erc20

import (
	"math/big"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/gobindings/generated/latest/factory_burn_mint_erc20"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

var ContractType cldf_deployment.ContractType = "ERC20"

var Version = semver.MustParse("1.0.0")

// ERC20ABI is the standard ERC20 ABI for name, symbol, and decimals functions
const ERC20ABI = `[
	{"constant": true, "inputs": [], "name": "name", "outputs": [{"name": "", "type": "string"}], "type": "function"},
	{"constant": true, "inputs": [], "name": "symbol", "outputs": [{"name": "", "type": "string"}], "type": "function"},
	{"constant": true, "inputs": [], "name": "decimals", "outputs": [{"name": "", "type": "uint8"}], "type": "function"}
]`

type ConstructorArgs struct {
	Name      string
	Symbol    string
	Decimals  uint8
	MaxSupply *big.Int
	PreMint   *big.Int
	NewOwner  common.Address
}

var Deploy = contract.NewDeploy(contract.DeployParams[ConstructorArgs]{
	Name:             "erc20:deploy",
	Version:          Version,
	Description:      "Deploys the ERC20 contract",
	ContractMetadata: factory_burn_mint_erc20.FactoryBurnMintERC20MetaData,
	BytecodeByTypeAndVersion: map[string]contract.Bytecode{
		cldf_deployment.NewTypeAndVersion(ContractType, *Version).String(): {
			EVM: common.FromHex(factory_burn_mint_erc20.FactoryBurnMintERC20Bin),
		},
	},
	Validate: func(ConstructorArgs) error { return nil },
})

var Name = contract.NewRead(contract.ReadParams[any, string, *factory_burn_mint_erc20.FactoryBurnMintERC20]{
	Name:         "erc20:name",
	Version:      Version,
	Description:  "Gets the name of an ERC20 token",
	ContractType: ContractType,
	NewContract:  factory_burn_mint_erc20.NewFactoryBurnMintERC20,
	CallContract: func(token *factory_burn_mint_erc20.FactoryBurnMintERC20, opts *bind.CallOpts, _ any) (string, error) {
		return token.Name(opts)
	},
})

var Symbol = contract.NewRead(contract.ReadParams[any, string, *factory_burn_mint_erc20.FactoryBurnMintERC20]{
	Name:         "erc20:symbol",
	Version:      Version,
	Description:  "Gets the symbol of an ERC20 token",
	ContractType: ContractType,
	NewContract:  factory_burn_mint_erc20.NewFactoryBurnMintERC20,
	CallContract: func(token *factory_burn_mint_erc20.FactoryBurnMintERC20, opts *bind.CallOpts, _ any) (string, error) {
		return token.Symbol(opts)
	},
})

var Decimals = contract.NewRead(contract.ReadParams[any, uint8, *factory_burn_mint_erc20.FactoryBurnMintERC20]{
	Name:         "erc20:decimals",
	Version:      Version,
	Description:  "Gets the decimals of an ERC20 token",
	ContractType: ContractType,
	NewContract:  factory_burn_mint_erc20.NewFactoryBurnMintERC20,
	CallContract: func(token *factory_burn_mint_erc20.FactoryBurnMintERC20, opts *bind.CallOpts, _ any) (uint8, error) {
		return token.Decimals(opts)
	},
})
