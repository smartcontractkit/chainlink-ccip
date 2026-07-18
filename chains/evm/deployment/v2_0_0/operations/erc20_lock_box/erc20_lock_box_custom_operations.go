package erc20_lock_box

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	gobindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v2_0_0/erc20_lock_box"
	cldf_evm "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations2/contract"
	cld_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

type DepositArgs struct {
	Token               common.Address
	RemoteChainSelector uint64
	Amount              *big.Int
}

func NewReadOwner(c gobindings.ERC20LockBoxInterface) *cld_ops.Operation[contract.FunctionInput[struct{}], common.Address, cldf_evm.Chain] {
	return contract.NewRead(contract.ReadParams[struct{}, common.Address, gobindings.ERC20LockBoxInterface]{
		Name:         "erc20-lock-box:owner",
		Version:      Version,
		Description:  "Calls owner on the contract",
		ContractType: ContractType,
		Contract:     c,
		CallContract: func(c gobindings.ERC20LockBoxInterface, opts *bind.CallOpts, _ struct{}) (common.Address, error) {
			return c.Owner(opts)
		},
	})
}

func NewWriteDeposit(c gobindings.ERC20LockBoxInterface) *cld_ops.Operation[contract.FunctionInput[DepositArgs], contract.WriteOutput, cldf_evm.Chain] {
	return contract.NewWrite(contract.WriteParams[DepositArgs, gobindings.ERC20LockBoxInterface]{
		Name:         "erc20-lock-box:deposit",
		Version:      Version,
		Description:  "Deposits tokens into the ERC20LockBox",
		ContractType: ContractType,
		ContractABI:  gobindings.ERC20LockBoxMetaData.ABI,
		Contract:     c,
		IsAllowedCaller: func(c gobindings.ERC20LockBoxInterface, opts *bind.CallOpts, caller common.Address, _ DepositArgs) (bool, error) {
			return contract.IsAuthorizedCaller(c, opts, caller)
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
		CallContract: func(c gobindings.ERC20LockBoxInterface, opts *bind.TransactOpts, args DepositArgs) (*types.Transaction, error) {
			return c.Deposit(opts, args.Token, args.RemoteChainSelector, args.Amount)
		},
	})
}

// NewWriteApplyAuthorizedCallerUpdatesProposalOnly is identical to NewWriteApplyAuthorizedCallerUpdates but
// forces the operation into an MCMS proposal. Use when lockbox ownership may still be with
// the deployer during proposal construction (e.g. liquidity migration batched with UpdateAuthorities).
func NewWriteApplyAuthorizedCallerUpdatesProposalOnly(c gobindings.ERC20LockBoxInterface) *cld_ops.Operation[contract.FunctionInput[gobindings.AuthorizedCallersAuthorizedCallerArgs], contract.WriteOutput, cldf_evm.Chain] {
	return contract.NewWrite(contract.WriteParams[gobindings.AuthorizedCallersAuthorizedCallerArgs, gobindings.ERC20LockBoxInterface]{
		Name:            "erc20-lock-box:apply-authorized-caller-updates-proposal-only",
		Version:         Version,
		Description:     "Calls applyAuthorizedCallerUpdates on the contract (proposal-only, never executed directly)",
		ContractType:    ContractType,
		ContractABI:     gobindings.ERC20LockBoxMetaData.ABI,
		Contract:        c,
		IsAllowedCaller: contract.NoCallersAllowed[gobindings.ERC20LockBoxInterface, gobindings.AuthorizedCallersAuthorizedCallerArgs],
		CallContract: func(
			c gobindings.ERC20LockBoxInterface,
			opts *bind.TransactOpts,
			args gobindings.AuthorizedCallersAuthorizedCallerArgs,
		) (*types.Transaction, error) {
			return c.ApplyAuthorizedCallerUpdates(opts, args)
		},
	})
}
