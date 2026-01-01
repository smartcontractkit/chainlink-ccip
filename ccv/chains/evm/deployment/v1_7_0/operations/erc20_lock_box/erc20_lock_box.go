package erc20_lock_box

import (
	"errors"
	"fmt"
	"slices"

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
	Validate: func(erc20LockBox *erc20_lock_box.ERC20LockBox, backend bind.ContractBackend, opts *bind.CallOpts, args AuthorizedCallerArgs) error {
		for _, caller := range args.AddedCallers {
			if caller == (common.Address{}) {
				return errors.New("caller cannot be the zero address")
			}
		}
		return nil
	},
	IsNoop: func(erc20LockBox *erc20_lock_box.ERC20LockBox, opts *bind.CallOpts, args AuthorizedCallerArgs) (bool, error) {
		allowedCallers, err := erc20LockBox.GetAllAuthorizedCallers(opts)
		if err != nil {
			return false, fmt.Errorf("failed to get all authorized callers: %w", err)
		}
		for _, caller := range args.AddedCallers {
			if !slices.Contains(allowedCallers, caller) {
				return false, nil
			}
		}
		for _, caller := range args.RemovedCallers {
			if slices.Contains(allowedCallers, caller) {
				return false, nil
			}
		}
		return true, nil
	},
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
