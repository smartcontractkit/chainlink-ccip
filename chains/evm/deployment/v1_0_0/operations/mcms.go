package operations

import (
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	cldf_evm "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	"github.com/smartcontractkit/mcms/sdk/evm/bindings"
	mcms_types "github.com/smartcontractkit/mcms/types"

	evmutils "github.com/smartcontractkit/chainlink-evm/pkg/utils"

	"github.com/smartcontractkit/chainlink-ccip/deployment/utils"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
)

var (
	EXECUTOR_ROLE_STR = "EXECUTOR_ROLE"
	EXECUTOR_ROLE     = Role{
		ID:   evmutils.MustHash(EXECUTOR_ROLE_STR),
		Name: EXECUTOR_ROLE_STR,
	}
)

type Role struct {
	ID   common.Hash
	Name string
}

type OpDeployTimelockInput struct {
	TimelockMinDelay *big.Int         `json:"timelockMinDelay"`
	Admin            common.Address   `json:"admin"`
	Proposers        []common.Address `json:"proposers"`
	Executors        []common.Address `json:"executors"`
	Cancellers       []common.Address `json:"cancellers"`
	Bypassers        []common.Address `json:"bypassers"`
}

type OpSetConfigMCMInput struct {
	SignerAddresses []common.Address `json:"signerAddresses"`
	SignerGroups    []uint8          `json:"signerGroups"` // Signer 1 is int group 0 (root group) with quorum 1.
	GroupQuorums    [32]uint8        `json:"groupQuorums"`
	GroupParents    [32]uint8        `json:"groupParents"`
}

type OpDeployCallProxyInput struct {
	TimelockAddress common.Address `json:"timelockAddress"`
}

type OpGrantRoleTimelockInput struct {
	Account common.Address `json:"account"`
	RoleID  [32]byte       `json:"roleID"`
}

type OpTransferOwnershipInput struct {
	ChainSelector   uint64
	TimelockAddress common.Address
	Address         common.Address
	ProposedOwner   common.Address
	ContractType    cldf_deployment.ContractType
}

type OpEVMOwnershipDeps struct {
	Chain    cldf_evm.Chain
	OwnableC OwnershipTranferable
}

type OwnershipTranferable interface {
	Owner(opts *bind.CallOpts) (common.Address, error)
	TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error)
	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)
	Address() common.Address
}

var OpDeployTimelock = contract.NewDeploy(contract.DeployParams[OpDeployTimelockInput]{
	Name:             "evm-timelock:deploy",
	Version:          semver.MustParse("1.0.0"),
	Description:      "Deploys Timelock contract on the specified EVM chain",
	ContractMetadata: bindings.RBACTimelockMetaData,
	BytecodeByTypeAndVersion: map[string]contract.Bytecode{
		cldf_deployment.NewTypeAndVersion(utils.RBACTimelock, *semver.MustParse("1.0.0")).String(): {
			EVM: common.FromHex(bindings.RBACTimelockBin),
		},
	},
	Validate: func(input OpDeployTimelockInput) error { return nil },
})

var OpDeployBypasserMCM = contract.NewDeploy(contract.DeployParams[struct{}]{
	Name:             "evm-bypasser-mcm:deploy",
	Version:          semver.MustParse("1.0.0"),
	Description:      "Deploys Bypasser MCM contract",
	ContractMetadata: bindings.ManyChainMultiSigMetaData,
	BytecodeByTypeAndVersion: map[string]contract.Bytecode{
		cldf_deployment.NewTypeAndVersion(utils.BypasserManyChainMultisig, *semver.MustParse("1.0.0")).String(): {
			EVM: common.FromHex(bindings.ManyChainMultiSigBin),
		},
	},
	Validate: func(input struct{}) error { return nil },
})

var OpDeployCancellerMCM = contract.NewDeploy(contract.DeployParams[struct{}]{
	Name:             "evm-canceller-mcm:deploy",
	Version:          semver.MustParse("1.0.0"),
	Description:      "Deploys Canceller MCM contract",
	ContractMetadata: bindings.ManyChainMultiSigMetaData,
	BytecodeByTypeAndVersion: map[string]contract.Bytecode{
		cldf_deployment.NewTypeAndVersion(utils.CancellerManyChainMultisig, *semver.MustParse("1.0.0")).String(): {
			EVM: common.FromHex(bindings.ManyChainMultiSigBin),
		},
	},
	Validate: func(input struct{}) error { return nil },
})

var OpDeployProposerMCM = contract.NewDeploy(contract.DeployParams[struct{}]{
	Name:             "evm-proposer-mcm:deploy",
	Version:          semver.MustParse("1.0.0"),
	Description:      "Deploys Proposer MCM contract",
	ContractMetadata: bindings.ManyChainMultiSigMetaData,
	BytecodeByTypeAndVersion: map[string]contract.Bytecode{
		cldf_deployment.NewTypeAndVersion(utils.ProposerManyChainMultisig, *semver.MustParse("1.0.0")).String(): {
			EVM: common.FromHex(bindings.ManyChainMultiSigBin),
		},
	},
	Validate: func(input struct{}) error { return nil },
})

var OpDeployCallProxy = contract.NewDeploy(contract.DeployParams[OpDeployCallProxyInput]{
	Name:             "evm-call-proxy:deploy",
	Version:          semver.MustParse("1.0.0"),
	Description:      "Deploys CallProxy contract on the specified EVM chain",
	ContractMetadata: bindings.CallProxyMetaData,
	BytecodeByTypeAndVersion: map[string]contract.Bytecode{
		cldf_deployment.NewTypeAndVersion(utils.CallProxy, *semver.MustParse("1.0.0")).String(): {
			EVM: common.FromHex(bindings.CallProxyBin),
		},
	},
	Validate: func(input OpDeployCallProxyInput) error { return nil },
})

var OpEVMSetConfigMCM = contract.NewWrite(contract.WriteParams[OpSetConfigMCMInput, *bindings.ManyChainMultiSig]{
	Name:            "evm-mcm-set-config",
	Version:         semver.MustParse("1.0.0"),
	Description:     "Sets Config on the deployed MCM contract",
	ContractABI:     bindings.ManyChainMultiSigABI,
	ContractType:    "ManyChainMultiSig",
	NewContract:     bindings.NewManyChainMultiSig,
	IsAllowedCaller: contract.OnlyOwner[*bindings.ManyChainMultiSig, OpSetConfigMCMInput],
	Validate:        func(input OpSetConfigMCMInput) error { return nil },
	CallContract: func(mcm *bindings.ManyChainMultiSig, opts *bind.TransactOpts, input OpSetConfigMCMInput) (*types.Transaction, error) {
		return mcm.SetConfig(
			opts,
			input.SignerAddresses,
			input.SignerGroups,
			input.GroupQuorums,
			input.GroupParents,
			false,
		)
	},
})

var OpGrantRoleTimelock = contract.NewWrite(contract.WriteParams[OpGrantRoleTimelockInput, *bindings.RBACTimelock]{
	Name:         "evm-timelock-grant-role",
	Version:      semver.MustParse("1.0.0"),
	Description:  "Grants role on the deployed Timelock contract",
	ContractABI:  bindings.RBACTimelockABI,
	ContractType: "RBACTimelock",
	NewContract:  bindings.NewRBACTimelock,
	IsAllowedCaller: func(contract *bindings.RBACTimelock, opts *bind.CallOpts, caller common.Address, input OpGrantRoleTimelockInput) (bool, error) {
		roleAdmin, err := contract.GetRoleAdmin(opts, input.RoleID)
		if err != nil {
			return false, err
		}
		// Check if caller has admin role of the role being granted
		return contract.HasRole(opts, roleAdmin, caller)
	},
	Validate: func(input OpGrantRoleTimelockInput) error {
		if input.Account == (common.Address{}) {
			return utils.ErrZeroAddress
		}
		return nil
	},
	CallContract: func(timelock *bindings.RBACTimelock, opts *bind.TransactOpts, input OpGrantRoleTimelockInput) (*types.Transaction, error) {
		return timelock.GrantRole(opts, input.RoleID, input.Account)
	},
})

var OpTransferOwnership = operations.NewOperation(
	"evm-transfer-ownership",
	semver.MustParse("1.0.0"),
	"Transfer ownership of an ownable contract to the specified address",
	func(b operations.Bundle, deps OpEVMOwnershipDeps, in OpTransferOwnershipInput) (contract.WriteOutput, error) {
		currentOwner, err := deps.OwnableC.Owner(&bind.CallOpts{
			Context: b.GetContext(),
		})
		if err != nil {
			return contract.WriteOutput{}, fmt.Errorf(
				"failed to get current owner of contract %T: %w",
				in.Address.Hex(),
				err,
			)
		}
		var opts *bind.TransactOpts
		allowed := false
		// if current owner is deployer, we can send the tx directly
		if currentOwner == deps.Chain.DeployerKey.From {
			opts = deps.Chain.DeployerKey
			allowed = true
		} else if currentOwner == in.TimelockAddress {
			allowed = false
			opts = cldf_deployment.SimTransactOpts()
		} else {
			return contract.WriteOutput{}, fmt.Errorf(
				"current owner %s is neither deployer %s nor timelock %s for contract %T",
				currentOwner.Hex(),
				deps.Chain.DeployerKey.From.Hex(),
				in.TimelockAddress.Hex(),
				in.Address.Hex(),
			)
		}
		// if current owner is timelock, we return the mcms transaction
		// if not, we execute the transfer directly through deployer
		tx, err := deps.OwnableC.TransferOwnership(opts, common.HexToAddress(in.ProposedOwner.Hex()))
		if allowed {
			_, err = cldf_deployment.ConfirmIfNoError(deps.Chain, tx, err)
			if err != nil {
				return contract.WriteOutput{}, fmt.Errorf(
					"failed to transfer ownership of contract %T: %w",
					in.Address.Hex(),
					err,
				)
			}
			b.Logger.Infof("Transferred ownership of contract %T to %s", in.Address.Hex(), in.ProposedOwner.Hex())
			return contract.WriteOutput{
				ChainSelector: in.ChainSelector,
				ExecInfo: &contract.ExecInfo{
					Hash: tx.Hash().String(),
				},
			}, nil
		} else {
			if err != nil {
				return contract.WriteOutput{}, fmt.Errorf(
					"failed to generate tx data for transfer ownership of contract %T to %s via timelock %s: %w",
					in.Address.Hex(),
					in.ProposedOwner.Hex(),
					in.TimelockAddress.Hex(),
					err,
				)
			}
			b.Logger.Infof("Generated transfer ownership tx data for contract %T to %s via timelock %s",
				in.Address.Hex(), in.ProposedOwner.Hex(), in.TimelockAddress.Hex())
			return contract.WriteOutput{
				ChainSelector: in.ChainSelector,
				Tx: mcms_types.Transaction{
					OperationMetadata: mcms_types.OperationMetadata{
						ContractType: string(in.ContractType),
					},
					To:               in.Address.Hex(),
					Data:             tx.Data(),
					AdditionalFields: json.RawMessage(`{"value": 0}`),
				},
			}, nil
		}
	})

var OpAcceptOwnership = operations.NewOperation(
	"evm-accept-ownership",
	semver.MustParse("1.0.0"),
	"Accepts ownership of an ownable contract Via the Timelock contract",
	func(b operations.Bundle, deps OpEVMOwnershipDeps, in OpTransferOwnershipInput) (contract.WriteOutput, error) {
		var opts *bind.TransactOpts
		var allowed bool
		if in.ProposedOwner == deps.Chain.DeployerKey.From {
			opts = deps.Chain.DeployerKey
			allowed = true
		} else if in.ProposedOwner == in.TimelockAddress {
			allowed = false
			opts = cldf_deployment.SimTransactOpts()
		} else {
			return contract.WriteOutput{}, fmt.Errorf(
				"proposed owner %s is neither deployer %s nor timelock %s for contract %T",
				in.ProposedOwner.Hex(),
				deps.Chain.DeployerKey.From.Hex(),
				in.TimelockAddress.Hex(),
				in.Address.Hex(),
			)
		}
		tx, err := deps.OwnableC.AcceptOwnership(opts)
		if allowed {
			_, err = cldf_deployment.ConfirmIfNoError(deps.Chain, tx, err)
			if err != nil {
				return contract.WriteOutput{}, fmt.Errorf(
					"failed to accept ownership of contract %T: %w",
					in.Address.Hex(),
					err,
				)
			}
			return contract.WriteOutput{
				ChainSelector: in.ChainSelector,
				ExecInfo: &contract.ExecInfo{
					Hash: tx.Hash().String(),
				},
			}, nil
		} else {
			if err != nil {
				return contract.WriteOutput{}, fmt.Errorf(
					"failed to generate tx data to Accept ownership of contract %T via timelock %s: %w",
					in.Address.Hex(),
					in.TimelockAddress.Hex(),
					err,
				)
			}
			return contract.WriteOutput{
				ChainSelector: in.ChainSelector,
				Tx: mcms_types.Transaction{
					OperationMetadata: mcms_types.OperationMetadata{
						ContractType: string(in.ContractType),
					},
					To:               in.Address.Hex(),
					Data:             tx.Data(),
					AdditionalFields: json.RawMessage(`{"value": 0}`),
				},
			}, nil
		}
	})
