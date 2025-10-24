package operations

import (
	"math/big"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/mcms/sdk/evm/bindings"

	evmutils "github.com/smartcontractkit/chainlink-evm/pkg/utils"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils"
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
