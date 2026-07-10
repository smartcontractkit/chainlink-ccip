package erc20_lock_box

import (
	"fmt"
	"math/big"
	"slices"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations/contract"
)

type DepositArgs struct {
	Token               common.Address
	RemoteChainSelector uint64
	Amount              *big.Int
}

func (c *ERC20LockBoxContract) Deposit(opts *bind.TransactOpts, token common.Address, arg1 uint64, amount *big.Int) (*types.Transaction, error) {
	return c.contract.Transact(opts, "deposit", token, arg1, amount)
}

var Owner = contract.NewRead(contract.ReadParams[struct{}, common.Address, *ERC20LockBoxContract]{
	Name:         "erc20-lock-box:owner",
	Version:      Version,
	Description:  "Calls owner on the contract",
	ContractType: ContractType,
	NewContract:  NewERC20LockBoxContract,
	CallContract: func(c *ERC20LockBoxContract, opts *bind.CallOpts, args struct{}) (common.Address, error) {
		return c.Owner(opts)
	},
})

var Deposit = contract.NewWrite(contract.WriteParams[DepositArgs, *ERC20LockBoxContract]{
	Name:         "erc20-lock-box:deposit",
	Version:      Version,
	Description:  "Deposits tokens into the ERC20LockBox",
	ContractType: ContractType,
	ContractABI:  ERC20LockBoxABI,
	NewContract:  NewERC20LockBoxContract,
	IsAllowedCaller: func(erc20LockBox *ERC20LockBoxContract, opts *bind.CallOpts, caller common.Address, args DepositArgs) (bool, error) {
		callers, err := erc20LockBox.GetAllAuthorizedCallers(opts)
		if err != nil {
			return false, fmt.Errorf("failed to get authorized callers: %w", err)
		}
		if slices.Contains(callers, caller) {
			return true, nil
		}
		return false, nil
	},
	Validate: func(args DepositArgs) error {
		if args.Amount == nil || args.Amount.Sign() <= 0 {
			return fmt.Errorf("amount must be greater than zero")
		}
		if args.Token == (common.Address{}) {
			return fmt.Errorf("token address must be set")
		}
		return nil
	},
	CallContract: func(erc20LockBox *ERC20LockBoxContract, opts *bind.TransactOpts, args DepositArgs) (*types.Transaction, error) {
		return erc20LockBox.Deposit(opts, args.Token, args.RemoteChainSelector, args.Amount)
	},
})

// ApplyAuthorizedCallerUpdatesProposalOnly is identical to ApplyAuthorizedCallerUpdates but
// forces the operation into an MCMS proposal. Use when lockbox ownership may still be with
// the deployer during proposal construction (e.g. liquidity migration batched with UpdateAuthorities).
var ApplyAuthorizedCallerUpdatesProposalOnly = contract.NewWrite(contract.WriteParams[AuthorizedCallerArgs, *ERC20LockBoxContract]{
	Name:            "erc20-lock-box:apply-authorized-caller-updates-proposal-only",
	Version:         Version,
	Description:     "Calls applyAuthorizedCallerUpdates on the contract (proposal-only, never executed directly)",
	ContractType:    ContractType,
	ContractABI:     ERC20LockBoxABI,
	NewContract:     NewERC20LockBoxContract,
	IsAllowedCaller: contract.NoCallersAllowed[*ERC20LockBoxContract, AuthorizedCallerArgs],
	Validate:        func(AuthorizedCallerArgs) error { return nil },
	CallContract: func(c *ERC20LockBoxContract, opts *bind.TransactOpts, args AuthorizedCallerArgs) (*types.Transaction, error) {
		return c.ApplyAuthorizedCallerUpdates(opts, args)
	},
})
