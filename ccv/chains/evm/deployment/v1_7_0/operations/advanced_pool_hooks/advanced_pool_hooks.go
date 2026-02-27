package advanced_pool_hooks

import (
	"math/big"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/gobindings/generated/latest/advanced_pool_hooks"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

var ContractType cldf_deployment.ContractType = "AdvancedPoolHooks"

var Version = semver.MustParse("1.7.0")

type ConstructorArgs struct {
	Allowlist                        []common.Address
	ThresholdAmountForAdditionalCCVs *big.Int
	PolicyEngine                     common.Address
	AuthorizedCallers                []common.Address
}

type CCVConfigArg = advanced_pool_hooks.AdvancedPoolHooksCCVConfigArg

type AllowlistUpdatesArgs struct {
	Adds    []common.Address
	Removes []common.Address
}

type GetRequiredCCVsArgs struct {
	RemoteChainSelector uint64
	Amount              *big.Int
	Direction           uint8 // IPoolV2.MessageDirection: Outbound = 0, Inbound = 1
}

var Deploy = contract.NewDeploy(contract.DeployParams[ConstructorArgs]{
	Name:             "advanced-pool-hooks:deploy",
	Version:          Version,
	Description:      "Deploys the AdvancedPoolHooks contract",
	ContractMetadata: advanced_pool_hooks.AdvancedPoolHooksMetaData,
	BytecodeByTypeAndVersion: map[string]contract.Bytecode{
		cldf_deployment.NewTypeAndVersion(ContractType, *Version).String(): {
			EVM: common.FromHex(advanced_pool_hooks.AdvancedPoolHooksBin),
		},
	},
	Validate: func(ConstructorArgs) error { return nil },
})

var ApplyCCVConfigUpdates = contract.NewWrite(contract.WriteParams[[]CCVConfigArg, *advanced_pool_hooks.AdvancedPoolHooks]{
	Name:            "advanced-pool-hooks:apply-ccv-config-updates",
	Version:         Version,
	Description:     "Applies CCV config updates to the AdvancedPoolHooks contract",
	ContractType:    ContractType,
	ContractABI:     advanced_pool_hooks.AdvancedPoolHooksABI,
	NewContract:     advanced_pool_hooks.NewAdvancedPoolHooks,
	IsAllowedCaller: contract.OnlyOwner[*advanced_pool_hooks.AdvancedPoolHooks, []CCVConfigArg],
	Validate:        func([]CCVConfigArg) error { return nil },
	CallContract: func(advancedPoolHooks *advanced_pool_hooks.AdvancedPoolHooks, opts *bind.TransactOpts, args []CCVConfigArg) (*types.Transaction, error) {
		return advancedPoolHooks.ApplyCCVConfigUpdates(opts, args)
	},
})

var ApplyAllowlistUpdates = contract.NewWrite(contract.WriteParams[AllowlistUpdatesArgs, *advanced_pool_hooks.AdvancedPoolHooks]{
	Name:            "advanced-pool-hooks:apply-allowlist-updates",
	Version:         Version,
	Description:     "Applies allowlist updates to the AdvancedPoolHooks contract",
	ContractType:    ContractType,
	ContractABI:     advanced_pool_hooks.AdvancedPoolHooksABI,
	NewContract:     advanced_pool_hooks.NewAdvancedPoolHooks,
	IsAllowedCaller: contract.OnlyOwner[*advanced_pool_hooks.AdvancedPoolHooks, AllowlistUpdatesArgs],
	Validate:        func(AllowlistUpdatesArgs) error { return nil },
	CallContract: func(advancedPoolHooks *advanced_pool_hooks.AdvancedPoolHooks, opts *bind.TransactOpts, args AllowlistUpdatesArgs) (*types.Transaction, error) {
		return advancedPoolHooks.ApplyAllowListUpdates(opts, args.Removes, args.Adds)
	},
})

var SetThresholdAmount = contract.NewWrite(contract.WriteParams[*big.Int, *advanced_pool_hooks.AdvancedPoolHooks]{
	Name:            "advanced-pool-hooks:set-threshold-amount",
	Version:         Version,
	Description:     "Sets the threshold amount above which additional CCVs are required",
	ContractType:    ContractType,
	ContractABI:     advanced_pool_hooks.AdvancedPoolHooksABI,
	NewContract:     advanced_pool_hooks.NewAdvancedPoolHooks,
	IsAllowedCaller: contract.OnlyOwner[*advanced_pool_hooks.AdvancedPoolHooks, *big.Int],
	Validate:        func(*big.Int) error { return nil },
	CallContract: func(advancedPoolHooks *advanced_pool_hooks.AdvancedPoolHooks, opts *bind.TransactOpts, args *big.Int) (*types.Transaction, error) {
		return advancedPoolHooks.SetThresholdAmount(opts, args)
	},
})

var GetAllowListEnabled = contract.NewRead(contract.ReadParams[any, bool, *advanced_pool_hooks.AdvancedPoolHooks]{
	Name:         "advanced-pool-hooks:get-allowlist-enabled",
	Version:      Version,
	Description:  "Gets whether the allowlist is enabled on the AdvancedPoolHooks contract",
	ContractType: ContractType,
	NewContract:  advanced_pool_hooks.NewAdvancedPoolHooks,
	CallContract: func(advancedPoolHooks *advanced_pool_hooks.AdvancedPoolHooks, opts *bind.CallOpts, args any) (bool, error) {
		return advancedPoolHooks.GetAllowListEnabled(opts)
	},
})

var GetAllowList = contract.NewRead(contract.ReadParams[any, []common.Address, *advanced_pool_hooks.AdvancedPoolHooks]{
	Name:         "advanced-pool-hooks:get-allowlist",
	Version:      Version,
	Description:  "Gets the allowlist on the AdvancedPoolHooks contract",
	ContractType: ContractType,
	NewContract:  advanced_pool_hooks.NewAdvancedPoolHooks,
	CallContract: func(advancedPoolHooks *advanced_pool_hooks.AdvancedPoolHooks, opts *bind.CallOpts, args any) ([]common.Address, error) {
		return advancedPoolHooks.GetAllowList(opts)
	},
})

var GetThresholdAmount = contract.NewRead(contract.ReadParams[any, *big.Int, *advanced_pool_hooks.AdvancedPoolHooks]{
	Name:         "advanced-pool-hooks:get-threshold-amount",
	Version:      Version,
	Description:  "Gets the threshold amount above which additional CCVs are required",
	ContractType: ContractType,
	NewContract:  advanced_pool_hooks.NewAdvancedPoolHooks,
	CallContract: func(advancedPoolHooks *advanced_pool_hooks.AdvancedPoolHooks, opts *bind.CallOpts, args any) (*big.Int, error) {
		return advancedPoolHooks.GetThresholdAmount(opts)
	},
})

var SetPolicyEngine = contract.NewWrite(contract.WriteParams[common.Address, *advanced_pool_hooks.AdvancedPoolHooks]{
	Name:            "advanced-pool-hooks:set-policy-engine",
	Version:         Version,
	Description:     "Sets a new policy engine on the AdvancedPoolHooks contract",
	ContractType:    ContractType,
	ContractABI:     advanced_pool_hooks.AdvancedPoolHooksABI,
	NewContract:     advanced_pool_hooks.NewAdvancedPoolHooks,
	IsAllowedCaller: contract.OnlyOwner[*advanced_pool_hooks.AdvancedPoolHooks, common.Address],
	Validate:        func(common.Address) error { return nil },
	CallContract: func(advancedPoolHooks *advanced_pool_hooks.AdvancedPoolHooks, opts *bind.TransactOpts, newPolicyEngine common.Address) (*types.Transaction, error) {
		return advancedPoolHooks.SetPolicyEngine(opts, newPolicyEngine)
	},
})

var SetPolicyEngineAllowFailedDetach = contract.NewWrite(contract.WriteParams[common.Address, *advanced_pool_hooks.AdvancedPoolHooks]{
	Name:            "advanced-pool-hooks:set-policy-engine-allow-failed-detach",
	Version:         Version,
	Description:     "Sets a new policy engine while tolerating a pre-existing policy engine's detach reverting",
	ContractType:    ContractType,
	ContractABI:     advanced_pool_hooks.AdvancedPoolHooksABI,
	NewContract:     advanced_pool_hooks.NewAdvancedPoolHooks,
	IsAllowedCaller: contract.OnlyOwner[*advanced_pool_hooks.AdvancedPoolHooks, common.Address],
	Validate:        func(common.Address) error { return nil },
	CallContract: func(advancedPoolHooks *advanced_pool_hooks.AdvancedPoolHooks, opts *bind.TransactOpts, newPolicyEngine common.Address) (*types.Transaction, error) {
		return advancedPoolHooks.SetPolicyEngineAllowFailedDetach(opts, newPolicyEngine)
	},
})

var GetPolicyEngine = contract.NewRead(contract.ReadParams[any, common.Address, *advanced_pool_hooks.AdvancedPoolHooks]{
	Name:         "advanced-pool-hooks:get-policy-engine",
	Version:      Version,
	Description:  "Gets the current policy engine address on the AdvancedPoolHooks contract",
	ContractType: ContractType,
	NewContract:  advanced_pool_hooks.NewAdvancedPoolHooks,
	CallContract: func(advancedPoolHooks *advanced_pool_hooks.AdvancedPoolHooks, opts *bind.CallOpts, args any) (common.Address, error) {
		return advancedPoolHooks.GetPolicyEngine(opts)
	},
})

type AuthorizedCallerArgs = advanced_pool_hooks.AuthorizedCallersAuthorizedCallerArgs

var ApplyAuthorizedCallerUpdates = contract.NewWrite(contract.WriteParams[AuthorizedCallerArgs, *advanced_pool_hooks.AdvancedPoolHooks]{
	Name:            "advanced-pool-hooks:apply-authorized-caller-updates",
	Version:         Version,
	Description:     "Applies authorized caller updates to the AdvancedPoolHooks contract",
	ContractType:    ContractType,
	ContractABI:     advanced_pool_hooks.AdvancedPoolHooksABI,
	NewContract:     advanced_pool_hooks.NewAdvancedPoolHooks,
	IsAllowedCaller: contract.OnlyOwner[*advanced_pool_hooks.AdvancedPoolHooks, AuthorizedCallerArgs],
	Validate:        func(AuthorizedCallerArgs) error { return nil },
	CallContract: func(advancedPoolHooks *advanced_pool_hooks.AdvancedPoolHooks, opts *bind.TransactOpts, args AuthorizedCallerArgs) (*types.Transaction, error) {
		return advancedPoolHooks.ApplyAuthorizedCallerUpdates(opts, args)
	},
})

var GetAllAuthorizedCallers = contract.NewRead(contract.ReadParams[any, []common.Address, *advanced_pool_hooks.AdvancedPoolHooks]{
	Name:         "advanced-pool-hooks:get-all-authorized-callers",
	Version:      Version,
	Description:  "Gets all authorized callers on the AdvancedPoolHooks contract",
	ContractType: ContractType,
	NewContract:  advanced_pool_hooks.NewAdvancedPoolHooks,
	CallContract: func(advancedPoolHooks *advanced_pool_hooks.AdvancedPoolHooks, opts *bind.CallOpts, args any) ([]common.Address, error) {
		return advancedPoolHooks.GetAllAuthorizedCallers(opts)
	},
})

var GetRequiredCCVs = contract.NewRead(contract.ReadParams[GetRequiredCCVsArgs, []common.Address, *advanced_pool_hooks.AdvancedPoolHooks]{
	Name:         "advanced-pool-hooks:get-required-ccvs",
	Version:      Version,
	Description:  "Gets the required CCVs for a given remote chain selector, amount, and direction",
	ContractType: ContractType,
	NewContract:  advanced_pool_hooks.NewAdvancedPoolHooks,
	CallContract: func(advancedPoolHooks *advanced_pool_hooks.AdvancedPoolHooks, opts *bind.CallOpts, args GetRequiredCCVsArgs) ([]common.Address, error) {
		return advancedPoolHooks.GetRequiredCCVs(opts, common.Address{}, args.RemoteChainSelector, args.Amount, 0, []byte{}, args.Direction)
	},
})
