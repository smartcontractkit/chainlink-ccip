package erc20

import (
	"math/big"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/gobindings/generated/latest/factory_burn_mint_erc20"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

var ContractType cldf_deployment.ContractType = "ERC20"

var Version = semver.MustParse("1.0.0")

type ApproveInput struct {
	Spender common.Address
	Amount  *big.Int
}

var Approve = contract.NewWrite(contract.WriteParams[ApproveInput, *factory_burn_mint_erc20.FactoryBurnMintERC20]{
	Name:         "erc20:approve",
	Version:      Version,
	Description:  "Approves a spender to transfer tokens on behalf of the caller",
	ContractType: ContractType,
	ContractABI:  factory_burn_mint_erc20.FactoryBurnMintERC20ABI,
	NewContract:  factory_burn_mint_erc20.NewFactoryBurnMintERC20,
	IsAllowedCaller: func(_ *factory_burn_mint_erc20.FactoryBurnMintERC20, _ *bind.CallOpts, _ common.Address, _ ApproveInput) (bool, error) {
		return true, nil
	},
	Validate: func(ApproveInput) error { return nil },
	CallContract: func(token *factory_burn_mint_erc20.FactoryBurnMintERC20, opts *bind.TransactOpts, input ApproveInput) (*types.Transaction, error) {
		return token.Approve(opts, input.Spender, input.Amount)
	},
})

var BalanceOf = contract.NewRead(contract.ReadParams[common.Address, *big.Int, *factory_burn_mint_erc20.FactoryBurnMintERC20]{
	Name:         "erc20:balance-of",
	Version:      Version,
	Description:  "Gets the token balance of an account",
	ContractType: ContractType,
	NewContract:  factory_burn_mint_erc20.NewFactoryBurnMintERC20,
	CallContract: func(token *factory_burn_mint_erc20.FactoryBurnMintERC20, opts *bind.CallOpts, account common.Address) (*big.Int, error) {
		return token.BalanceOf(opts, account)
	},
})
