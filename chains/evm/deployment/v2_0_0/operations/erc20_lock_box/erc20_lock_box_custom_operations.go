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

// DepositArgs is the input for the custom deposit operation.
type DepositArgs struct {
	Token               common.Address `json:"token"`
	RemoteChainSelector uint64         `json:"remoteChainSelector"`
	Amount              *big.Int       `json:"amount"`
}

// NewReadOwner returns a read operation for the lock box owner.
func NewReadOwner(c gobindings.ERC20LockBoxInterface) *cld_ops.Operation[contract.FunctionInput[struct{}], common.Address, cldf_evm.Chain] {
	return contract.NewRead(contract.ReadParams[struct{}, common.Address, gobindings.ERC20LockBoxInterface]{
		Name:         "erc20-lock-box:owner",
		Version:      Version,
		Description:  "Calls owner on the contract",
		ContractType: ContractType,
		Contract:     c,
		CallContract: func(c gobindings.ERC20LockBoxInterface, opts *bind.CallOpts, args struct{}) (common.Address, error) {
			return c.Owner(opts)
		},
	})
}

// NewWriteDeposit returns a write operation that deposits into the lock box for an authorized caller.
func NewWriteDeposit(c gobindings.ERC20LockBoxInterface) *cld_ops.Operation[contract.FunctionInput[DepositArgs], contract.WriteOutput, cldf_evm.Chain] {
	return contract.NewWrite(contract.WriteParams[DepositArgs, gobindings.ERC20LockBoxInterface]{
		Name:         "erc20-lock-box:deposit",
		Version:      Version,
		Description:  "Deposits tokens into the ERC20LockBox",
		ContractType: ContractType,
		ContractABI:  gobindings.ERC20LockBoxMetaData.ABI,
		Contract:     c,
		IsAllowedCaller: func(c gobindings.ERC20LockBoxInterface, opts *bind.CallOpts, caller common.Address, args DepositArgs) (bool, error) {
			callers, err := c.GetAllAuthorizedCallers(opts)
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
		CallContract: func(
			c gobindings.ERC20LockBoxInterface,
			opts *bind.TransactOpts,
			args DepositArgs,
		) (*types.Transaction, error) {
			return c.Deposit(opts, args.Token, args.RemoteChainSelector, args.Amount)
		},
	})
}
