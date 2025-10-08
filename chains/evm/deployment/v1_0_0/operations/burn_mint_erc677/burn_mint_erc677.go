package burn_mint_erc677

import (
	"math/big"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-evm/gethwrappers/shared/generated/initial/burn_mint_erc677"
)

var ContractType cldf_deployment.ContractType = "BurnMintERC677"

type ConstructorArgs struct {
	Name      string
	Symbol    string
	Decimals  uint8
	MaxSupply *big.Int
}

type MintArgs struct {
	Account common.Address
	Amount  *big.Int
}

var Deploy = contract.NewDeploy(contract.DeployParams[ConstructorArgs]{
	Name:             "burn-mint-erc677:deploy",
	Version:          semver.MustParse("1.0.0"),
	Description:      "Deploys the BurnMintERC677 contract",
	ContractMetadata: burn_mint_erc677.BurnMintERC677MetaData,
	BytecodeByTypeAndVersion: map[string]contract.Bytecode{
		cldf_deployment.NewTypeAndVersion(ContractType, *semver.MustParse("1.0.0")).String(): {
			EVM: common.FromHex(burn_mint_erc677.BurnMintERC677Bin),
		},
	},
	Validate: func(ConstructorArgs) error { return nil },
})

var GrantMintAndBurnRoles = contract.NewWrite(contract.WriteParams[common.Address, *burn_mint_erc677.BurnMintERC677]{
	Name:            "burn-mint-erc677:grant-mint-and-burn-roles",
	Version:         semver.MustParse("1.0.0"),
	Description:     "Grants mint and burn roles on the token to an account",
	ContractType:    ContractType,
	ContractABI:     burn_mint_erc677.BurnMintERC677ABI,
	NewContract:     burn_mint_erc677.NewBurnMintERC677,
	IsAllowedCaller: contract.OnlyOwner[*burn_mint_erc677.BurnMintERC677, common.Address],
	Validate:        func(common.Address) error { return nil },
	CallContract: func(token *burn_mint_erc677.BurnMintERC677, opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
		return token.GrantMintAndBurnRoles(opts, account)
	},
})

var GrantMintRole = contract.NewWrite(contract.WriteParams[common.Address, *burn_mint_erc677.BurnMintERC677]{
	Name:            "burn-mint-erc677:grant-mint-role",
	Version:         semver.MustParse("1.0.0"),
	Description:     "Grants the mint role on the token to an account",
	ContractType:    ContractType,
	ContractABI:     burn_mint_erc677.BurnMintERC677ABI,
	NewContract:     burn_mint_erc677.NewBurnMintERC677,
	IsAllowedCaller: contract.OnlyOwner[*burn_mint_erc677.BurnMintERC677, common.Address],
	Validate:        func(common.Address) error { return nil },
	CallContract: func(token *burn_mint_erc677.BurnMintERC677, opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
		return token.GrantMintRole(opts, account)
	},
})

var RevokeBurnRole = contract.NewWrite(contract.WriteParams[common.Address, *burn_mint_erc677.BurnMintERC677]{
	Name:            "burn-mint-erc677:revoke-burn-role",
	Version:         semver.MustParse("1.0.0"),
	Description:     "Revokes the burn role on the token from an account",
	ContractType:    ContractType,
	ContractABI:     burn_mint_erc677.BurnMintERC677ABI,
	NewContract:     burn_mint_erc677.NewBurnMintERC677,
	IsAllowedCaller: contract.OnlyOwner[*burn_mint_erc677.BurnMintERC677, common.Address],
	Validate:        func(common.Address) error { return nil },
	CallContract: func(token *burn_mint_erc677.BurnMintERC677, opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
		return token.RevokeBurnRole(opts, account)
	},
})

var RevokeMintRole = contract.NewWrite(contract.WriteParams[common.Address, *burn_mint_erc677.BurnMintERC677]{
	Name:            "burn-mint-erc677:revoke-mint-role",
	Version:         semver.MustParse("1.0.0"),
	Description:     "Revokes the mint role on the token from an account",
	ContractType:    ContractType,
	ContractABI:     burn_mint_erc677.BurnMintERC677ABI,
	NewContract:     burn_mint_erc677.NewBurnMintERC677,
	IsAllowedCaller: contract.OnlyOwner[*burn_mint_erc677.BurnMintERC677, common.Address],
	Validate:        func(common.Address) error { return nil },
	CallContract: func(token *burn_mint_erc677.BurnMintERC677, opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
		return token.RevokeMintRole(opts, account)
	},
})

var Mint = contract.NewWrite(contract.WriteParams[MintArgs, *burn_mint_erc677.BurnMintERC677]{
	Name:            "burn-mint-erc677:mint",
	Version:         semver.MustParse("1.0.0"),
	Description:     "Mint tokens to an account",
	ContractType:    ContractType,
	ContractABI:     burn_mint_erc677.BurnMintERC677ABI,
	NewContract:     burn_mint_erc677.NewBurnMintERC677,
	IsAllowedCaller: contract.OnlyOwner[*burn_mint_erc677.BurnMintERC677, MintArgs],
	Validate:        func(MintArgs) error { return nil },
	CallContract: func(token *burn_mint_erc677.BurnMintERC677, opts *bind.TransactOpts, args MintArgs) (*types.Transaction, error) {
		return token.Mint(opts, args.Account, args.Amount)
	},
})
