package erc20_lock_box

import (
	"fmt"
	"math/big"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/gobindings/generated/latest/erc20_lock_box"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

var ContractType cldf_deployment.ContractType = "ERC20LockBox"

var Version = semver.MustParse("1.7.0")

type ConstructorArgs struct {
	Token common.Address
}

type AuthorizedCallerArgs = erc20_lock_box.AuthorizedCallersAuthorizedCallerArgs

type DepositArgs struct {
	Token               common.Address
	RemoteChainSelector uint64
	Amount              *big.Int
}

var Deploy = contract.NewDeploy(contract.DeployParams[ConstructorArgs]{
	Name:             "erc20-lock-box:deploy",
	Version:          Version,
	Description:      "Deploys the ERC20LockBox contract",
	ContractMetadata: erc20_lock_box.ERC20LockBoxMetaData,
	BytecodeByTypeAndVersion: map[string]contract.Bytecode{
		cldf_deployment.NewTypeAndVersion(ContractType, *Version).String(): {
			EVM: common.FromHex(erc20_lock_box.ERC20LockBoxBin),
		},
	},
})

var ApplyAuthorizedCallerUpdates = contract.NewWrite(contract.WriteParams[AuthorizedCallerArgs, *erc20_lock_box.ERC20LockBox]{
	Name:            "erc20-lock-box:apply-authorized-caller-updates",
	Version:         Version,
	Description:     "Applies the authorized caller updates to the ERC20LockBox contract",
	ContractType:    ContractType,
	ContractABI:     erc20_lock_box.ERC20LockBoxABI,
	NewContract:     erc20_lock_box.NewERC20LockBox,
	IsAllowedCaller: contract.OnlyOwner[*erc20_lock_box.ERC20LockBox, AuthorizedCallerArgs],
	Validate:        func(AuthorizedCallerArgs) error { return nil },
	CallContract: func(erc20LockBox *erc20_lock_box.ERC20LockBox, opts *bind.TransactOpts, args AuthorizedCallerArgs) (*types.Transaction, error) {
		return erc20LockBox.ApplyAuthorizedCallerUpdates(opts, args)
	},
})

var GetAllAuthorizedCallers = contract.NewRead(contract.ReadParams[any, []common.Address, *erc20_lock_box.ERC20LockBox]{
	Name:         "erc20-lock-box:get-all-authorized-callers",
	Version:      Version,
	Description:  "Gets all authorized callers on the ERC20LockBox",
	ContractType: ContractType,
	NewContract:  erc20_lock_box.NewERC20LockBox,
	CallContract: func(erc20LockBox *erc20_lock_box.ERC20LockBox, opts *bind.CallOpts, args any) ([]common.Address, error) {
		return erc20LockBox.GetAllAuthorizedCallers(opts)
	},
})

var Deposit = contract.NewWrite(contract.WriteParams[DepositArgs, *erc20_lock_box.ERC20LockBox]{
	Name:         "erc20-lock-box:deposit",
	Version:      Version,
	Description:  "Deposits tokens into the ERC20LockBox",
	ContractType: ContractType,
	ContractABI:  erc20_lock_box.ERC20LockBoxABI,
	NewContract:  erc20_lock_box.NewERC20LockBox,
	IsAllowedCaller: func(erc20LockBox *erc20_lock_box.ERC20LockBox, opts *bind.CallOpts, caller common.Address, args DepositArgs) (bool, error) {
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
	CallContract: func(erc20LockBox *erc20_lock_box.ERC20LockBox, opts *bind.TransactOpts, args DepositArgs) (*types.Transaction, error) {
		return erc20LockBox.Deposit(opts, args.Token, args.RemoteChainSelector, args.Amount)
	},
})
