package advanced_pool_hooks

import (
	"errors"
	"fmt"
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
}

type CCVConfigArg = advanced_pool_hooks.AdvancedPoolHooksCCVConfigArg

type AllowlistUpdatesArgs struct {
	Adds    []common.Address
	Removes []common.Address
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
	Validate: func(advancedPoolHooks *advanced_pool_hooks.AdvancedPoolHooks, backend bind.ContractBackend, opts *bind.CallOpts, configs []CCVConfigArg) error {
		for _, cfg := range configs {
			// Ensure that OutboundCCVs has no duplicates.
			if hasDuplicates(cfg.OutboundCCVs) {
				return errors.New("outbound CCVs must not contain duplicates")
			}
			// Ensure that InboundCCVs has no duplicates.
			if hasDuplicates(cfg.InboundCCVs) {
				return errors.New("inbound CCVs must not contain duplicates")
			}

			if len(cfg.ThresholdOutboundCCVs) > 0 {
				if len(cfg.OutboundCCVs) == 0 {
					return errors.New("threshold outbound CCVs must be specified when outbound CCVs are specified")
				}
				// Ensure that ThresholdOutboundCCVs has no duplicates.
				if hasDuplicates(cfg.ThresholdOutboundCCVs) {
					return errors.New("threshold outbound CCVs must not contain duplicates")
				}
				// Ensure that ThresholdOutboundCCVs and OutboundCCVs do not overlap.
				if hasOverlap(cfg.OutboundCCVs, cfg.ThresholdOutboundCCVs) {
					return errors.New("threshold outbound CCVs must not overlap with outbound CCVs")
				}
			}

			if len(cfg.ThresholdInboundCCVs) > 0 {
				if len(cfg.InboundCCVs) == 0 {
					return errors.New("threshold inbound CCVs must be specified when inbound CCVs are specified")
				}
				// Ensure that ThresholdInboundCCVs has no duplicates.
				if hasDuplicates(cfg.ThresholdInboundCCVs) {
					return errors.New("threshold inbound CCVs must not contain duplicates")
				}
				// Ensure that ThresholdInboundCCVs and InboundCCVs do not overlap.
				if hasOverlap(cfg.InboundCCVs, cfg.ThresholdInboundCCVs) {
					return errors.New("threshold inbound CCVs must not overlap with inbound CCVs")
				}
			}
		}

		return nil
	},
	IsNoop: func(advancedPoolHooks *advanced_pool_hooks.AdvancedPoolHooks, opts *bind.CallOpts, args []CCVConfigArg) (bool, error) {
		// GetRequiredCCVs is not a reliable way to determine if the operation is a noop because it doesn't give a full picture of state.
		// It just provides a list of required CCVs for given parameters. Therefore, we just check if the args are empty to determine if the operation is a noop.
		return len(args) == 0, nil
	},
	CallContract: func(advancedPoolHooks *advanced_pool_hooks.AdvancedPoolHooks, opts *bind.TransactOpts, args []CCVConfigArg) (*types.Transaction, error) {
		return advancedPoolHooks.ApplyCCVConfigUpdates(opts, args)
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
	Validate: func(advancedPoolHooks *advanced_pool_hooks.AdvancedPoolHooks, backend bind.ContractBackend, opts *bind.CallOpts, args *big.Int) error {
		return nil
	},
	IsNoop: func(advancedPoolHooks *advanced_pool_hooks.AdvancedPoolHooks, opts *bind.CallOpts, args *big.Int) (bool, error) {
		if args == nil {
			return true, nil
		}

		currentThresholdAmount, err := advancedPoolHooks.GetThresholdAmount(opts)
		if err != nil {
			return false, fmt.Errorf("failed to get current threshold amount on advanced pool hooks with address %s: %w", advancedPoolHooks.Address(), err)
		}

		return currentThresholdAmount.Cmp(args) == 0, nil
	},
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

func hasDuplicates(list []common.Address) bool {
	seen := make(map[common.Address]bool)
	for _, addr := range list {
		if seen[addr] {
			return true
		}
		seen[addr] = true
	}
	return false
}

func hasOverlap(listA, listB []common.Address) bool {
	for _, addrA := range listA {
		for _, addrB := range listB {
			if addrA == addrB {
				return true
			}
		}
	}
	return false
}
