package erc20

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	gobindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v2_0_0/cross_chain_token"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
)

var BalanceOf = contract.NewRead(contract.ReadParams[common.Address, *big.Int, *gobindings.CrossChainToken]{
	Name:         "erc20:balance-of",
	Version:      Version,
	Description:  "Calls balanceOf on the contract",
	ContractType: ContractType,
	NewContract:  gobindings.NewCrossChainToken,
	CallContract: func(c *gobindings.CrossChainToken, opts *bind.CallOpts, args common.Address) (*big.Int, error) {
		return c.BalanceOf(opts, args)
	},
})

var Approve = contract.NewWrite(contract.WriteParams[ApproveArgs, *gobindings.CrossChainToken]{
	Name:            "erc20:approve",
	Version:         Version,
	Description:     "Calls approve on the contract",
	ContractType:    ContractType,
	ContractABI:     gobindings.CrossChainTokenMetaData.ABI,
	NewContract:     gobindings.NewCrossChainToken,
	IsAllowedCaller: contract.AllCallersAllowed[*gobindings.CrossChainToken, ApproveArgs],
	Validate:        func(ApproveArgs) error { return nil },
	CallContract: func(c *gobindings.CrossChainToken, opts *bind.TransactOpts, args ApproveArgs) (*types.Transaction, error) {
		return c.Approve(opts, args.Spender, args.Value)
	},
})
