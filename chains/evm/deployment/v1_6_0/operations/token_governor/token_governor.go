package token_governor

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/smartcontractkit/ccip-contract-examples/chains/evm/gobindings/generated/latest/token_governor"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

var ContractType cldf_deployment.ContractType = "TokenGovernor"
var TypeAndVersion = cldf_deployment.NewTypeAndVersion(ContractType, *Version)

var Version = utils.Version_1_6_0

type ConstructorArgs struct {
	Token               common.Address
	InitialDelay        *big.Int
	InitialDefaultAdmin common.Address
}

type RoleAssignment struct {
	Role [32]byte
	To   common.Address
}

var Deploy = contract.NewDeploy(contract.DeployParams[ConstructorArgs]{
	Name:             "token_governor:deploy",
	Version:          Version,
	Description:      "Deploys the TokenGovernor contract",
	ContractMetadata: token_governor.TokenGovernorMetaData,
	BytecodeByTypeAndVersion: map[string]contract.Bytecode{
		cldf_deployment.NewTypeAndVersion(ContractType, *Version).String(): {
			EVM: common.FromHex(token_governor.TokenGovernorBin),
		},
	},
	Validate: func(args ConstructorArgs) error { return nil },
})

var GrantRole = contract.NewWrite(contract.WriteParams[RoleAssignment, *token_governor.TokenGovernor]{
	Name:         "token_governor:grant_roles",
	Version:      utils.Version_1_0_0,
	Description:  "Grant access to given roles on TokenGovernor contract",
	ContractType: ContractType,
	ContractABI:  token_governor.TokenGovernorABI,
	NewContract:  token_governor.NewTokenGovernor,
	IsAllowedCaller: func(tg *token_governor.TokenGovernor, opts *bind.CallOpts, caller common.Address, input RoleAssignment) (bool, error) {
		defaultAdmin, err := tg.DefaultAdmin(opts)
		if err != nil {
			return false, err
		}
		return defaultAdmin == caller, nil
	},
	Validate: func(RoleAssignment) error {
		return nil
	},
	CallContract: func(tg *token_governor.TokenGovernor, opts *bind.TransactOpts, input RoleAssignment) (*types.Transaction, error) {
		if input.Role == [32]byte{} {
			return nil, errors.New(fmt.Sprintf("invalid role input %v for grant access", input.Role))
		}
		return tg.GrantRole(opts, input.Role, input.To)
	},
})

var RenounceRole = contract.NewWrite(contract.WriteParams[RoleAssignment, *token_governor.TokenGovernor]{
	Name:         "token_governor:renounce_roles",
	Version:      utils.Version_1_0_0,
	Description:  "Renounce access to given roles on TokenGovernor contract",
	ContractType: ContractType,
	ContractABI:  token_governor.TokenGovernorABI,
	NewContract:  token_governor.NewTokenGovernor,
	IsAllowedCaller: func(tg *token_governor.TokenGovernor, opts *bind.CallOpts, caller common.Address, input RoleAssignment) (bool, error) {
		defaultAdmin, err := tg.DefaultAdmin(opts)
		if err != nil {
			return false, err
		}
		return defaultAdmin == caller, nil
	},
	Validate: func(RoleAssignment) error {
		return nil
	},
	CallContract: func(tg *token_governor.TokenGovernor, opts *bind.TransactOpts, input RoleAssignment) (*types.Transaction, error) {
		if input.Role == [32]byte{} {
			return nil, errors.New(fmt.Sprintf("invalid role input %v for renounce role access", input.Role))
		}
		return tg.RenounceRole(opts, input.Role, input.To)
	},
})

var RevokeRole = contract.NewWrite(contract.WriteParams[RoleAssignment, *token_governor.TokenGovernor]{
	Name:         "token_governor:revoke_roles",
	Version:      utils.Version_1_0_0,
	Description:  "Revoke access to given roles on TokenGovernor contract",
	ContractType: ContractType,
	ContractABI:  token_governor.TokenGovernorABI,
	NewContract:  token_governor.NewTokenGovernor,
	IsAllowedCaller: func(tg *token_governor.TokenGovernor, opts *bind.CallOpts, caller common.Address, input RoleAssignment) (bool, error) {
		defaultAdmin, err := tg.DefaultAdmin(opts)
		if err != nil {
			return false, err
		}
		return defaultAdmin == caller, nil
	},
	Validate: func(RoleAssignment) error {
		return nil
	},
	CallContract: func(tg *token_governor.TokenGovernor, opts *bind.TransactOpts, input RoleAssignment) (*types.Transaction, error) {
		if input.Role == [32]byte{} {
			return nil, errors.New(fmt.Sprintf("invalid role input %v for revoke access", input.Role))
		}
		return tg.RevokeRole(opts, input.Role, input.To)
	},
})

var TransferOwnership = contract.NewWrite(contract.WriteParams[common.Address, *token_governor.TokenGovernor]{
	Name:         "token_governor:transfer_ownership",
	Version:      utils.Version_1_0_0,
	Description:  "TransferOwnership to given account on TokenGovernor contract",
	ContractType: ContractType,
	ContractABI:  token_governor.TokenGovernorABI,
	NewContract:  token_governor.NewTokenGovernor,
	IsAllowedCaller: func(tg *token_governor.TokenGovernor, opts *bind.CallOpts, caller common.Address, input common.Address) (bool, error) {
		owner, err := tg.Owner(opts)
		if err != nil {
			return false, err
		}
		return owner == caller, nil
	},
	Validate: func(common.Address) error {
		return nil
	},
	CallContract: func(tg *token_governor.TokenGovernor, opts *bind.TransactOpts, input common.Address) (*types.Transaction, error) {
		return tg.TransferOwnership(opts, input)
	},
})

var AcceptOwnership = contract.NewWrite(contract.WriteParams[common.Address, *token_governor.TokenGovernor]{
	Name:         "token_governor:accept_ownership",
	Version:      utils.Version_1_0_0,
	Description:  "AcceptOwnership on TokenGovernor contract",
	ContractType: ContractType,
	ContractABI:  token_governor.TokenGovernorABI,
	NewContract:  token_governor.NewTokenGovernor,
	IsAllowedCaller: func(tg *token_governor.TokenGovernor, opts *bind.CallOpts, caller common.Address, input common.Address) (bool, error) {
		owner, err := tg.Owner(opts)
		if err != nil {
			return false, err
		}
		return owner == caller, nil
	},
	Validate: func(common.Address) error {
		return nil
	},
	CallContract: func(tg *token_governor.TokenGovernor, opts *bind.TransactOpts, input common.Address) (*types.Transaction, error) {
		return tg.AcceptOwnership(opts)
	},
})

var BeginDefaultAdminTransfer = contract.NewWrite(contract.WriteParams[common.Address, *token_governor.TokenGovernor]{
	Name:         "token_governor:begin_default_admin_transfer",
	Version:      utils.Version_1_0_0,
	Description:  "BeginDefaultAdminTransfer on TokenGovernor contract",
	ContractType: ContractType,
	ContractABI:  token_governor.TokenGovernorABI,
	NewContract:  token_governor.NewTokenGovernor,
	IsAllowedCaller: func(tg *token_governor.TokenGovernor, opts *bind.CallOpts, caller common.Address, input common.Address) (bool, error) {
		owner, err := tg.Owner(opts)
		if err != nil {
			return false, err
		}
		return owner == caller, nil
	},
	Validate: func(common.Address) error {
		return nil
	},
	CallContract: func(tg *token_governor.TokenGovernor, opts *bind.TransactOpts, input common.Address) (*types.Transaction, error) {
		return tg.BeginDefaultAdminTransfer(opts, input)
	},
})

var AcceptDefaultAdminTransfer = contract.NewWrite(contract.WriteParams[common.Address, *token_governor.TokenGovernor]{
	Name:         "token_governor:accept_default_admin",
	Version:      utils.Version_1_0_0,
	Description:  "AcceptDefaultAdminTransfer on TokenGovernor contract",
	ContractType: ContractType,
	ContractABI:  token_governor.TokenGovernorABI,
	NewContract:  token_governor.NewTokenGovernor,
	IsAllowedCaller: func(tg *token_governor.TokenGovernor, opts *bind.CallOpts, caller common.Address, input common.Address) (bool, error) {
		owner, err := tg.Owner(opts)
		if err != nil {
			return false, err
		}
		return owner == caller, nil
	},
	Validate: func(common.Address) error {
		return nil
	},
	CallContract: func(tg *token_governor.TokenGovernor, opts *bind.TransactOpts, input common.Address) (*types.Transaction, error) {
		return tg.AcceptDefaultAdminTransfer(opts)
	},
})
