package erc20_lock_box

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	elboxg "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v2_0_0/erc20_lock_box"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
)

type DepositArgs struct {
	Token               common.Address
	RemoteChainSelector uint64
	Amount              *big.Int
}

var Owner = contract.NewRead(contract.ReadParams[struct{}, common.Address, *elboxg.ERC20LockBox]{
	Name:         "erc20-lock-box:owner",
	Version:      Version,
	Description:  "Calls owner on the contract",
	ContractType: ContractType,
	NewContract:  elboxg.NewERC20LockBox,
	CallContract: func(c *elboxg.ERC20LockBox, opts *bind.CallOpts, args struct{}) (common.Address, error) {
		return c.Owner(opts)
	},
})

var Deposit = contract.NewWrite(contract.WriteParams[DepositArgs, *elboxg.ERC20LockBox]{
	Name:         "erc20-lock-box:deposit",
	Version:      Version,
	Description:  "Deposits tokens into the ERC20LockBox",
	ContractType: ContractType,
	ContractABI:  elboxg.ERC20LockBoxABI,
	NewContract:  elboxg.NewERC20LockBox,
	IsAllowedCaller: func(erc20LockBox *elboxg.ERC20LockBox, opts *bind.CallOpts, caller common.Address, args DepositArgs) (bool, error) {
		callers, err := erc20LockBox.GetAllAuthorizedCallers(opts)
		if err != nil {
			return false, err
		}
		for _, authorized := range callers {
			if authorized == caller {
				return true, nil
			}
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
	CallContract: func(erc20LockBox *elboxg.ERC20LockBox, opts *bind.TransactOpts, args DepositArgs) (*types.Transaction, error) {
		return erc20LockBox.Deposit(opts, args.Token, args.RemoteChainSelector, args.Amount)
	},
})

// AuthorizedCallerArgs matches applyAuthorizedCallerUpdates input.
type AuthorizedCallerArgs = elboxg.AuthorizedCallersAuthorizedCallerArgs

var ApplyAuthorizedCallerUpdates = contract.NewWrite(contract.WriteParams[elboxg.AuthorizedCallersAuthorizedCallerArgs, *elboxg.ERC20LockBox]{
	Name:            "erc20-lock-box:apply-authorized-caller-updates",
	Version:         Version,
	Description:     "Calls applyAuthorizedCallerUpdates on the contract",
	ContractType:    ContractType,
	ContractABI:     elboxg.ERC20LockBoxMetaData.ABI,
	NewContract:     elboxg.NewERC20LockBox,
	IsAllowedCaller: contract.OnlyOwner[*elboxg.ERC20LockBox, elboxg.AuthorizedCallersAuthorizedCallerArgs],
	Validate:        func(elboxg.AuthorizedCallersAuthorizedCallerArgs) error { return nil },
	CallContract: func(c *elboxg.ERC20LockBox, opts *bind.TransactOpts, args elboxg.AuthorizedCallersAuthorizedCallerArgs) (*types.Transaction, error) {
		return c.ApplyAuthorizedCallerUpdates(opts, args)
	},
})

var GetAllAuthorizedCallers = contract.NewRead(contract.ReadParams[struct{}, []common.Address, *elboxg.ERC20LockBox]{
	Name:         "erc20-lock-box:get-all-authorized-callers",
	Version:      Version,
	Description:  "Calls getAllAuthorizedCallers on the contract",
	ContractType: ContractType,
	NewContract:  elboxg.NewERC20LockBox,
	CallContract: func(c *elboxg.ERC20LockBox, opts *bind.CallOpts, args struct{}) ([]common.Address, error) {
		return c.GetAllAuthorizedCallers(opts)
	},
})

var TransferOwnership = contract.NewWrite(contract.WriteParams[common.Address, *elboxg.ERC20LockBox]{
	Name:            "erc20-lock-box:transfer-ownership",
	Version:         Version,
	Description:     "Calls transferOwnership on the contract",
	ContractType:    ContractType,
	ContractABI:     elboxg.ERC20LockBoxMetaData.ABI,
	NewContract:     elboxg.NewERC20LockBox,
	IsAllowedCaller: contract.OnlyOwner[*elboxg.ERC20LockBox, common.Address],
	Validate:        func(common.Address) error { return nil },
	CallContract: func(c *elboxg.ERC20LockBox, opts *bind.TransactOpts, args common.Address) (*types.Transaction, error) {
		return c.TransferOwnership(opts, args)
	},
})
