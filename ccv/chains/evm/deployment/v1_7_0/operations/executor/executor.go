package executor

import (
	"errors"
	"fmt"
	"slices"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/gobindings/generated/latest/executor"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/gobindings/generated/latest/proxy"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/adapters"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

var ContractType cldf_deployment.ContractType = "Executor"

var ProxyType cldf_deployment.ContractType = "ExecutorProxy"

var Version = semver.MustParse("1.7.0")

type ConstructorArgs struct {
	MaxCCVsPerMsg uint8
	DynamicConfig executor.ExecutorDynamicConfig
}

type ProxyConstructorArgs struct {
	ExecutorAddress common.Address
}

type ApplyDestChainUpdatesArgs struct {
	DestChainSelectorsToAdd    []RemoteChainConfigArgs
	DestChainSelectorsToRemove []uint64
}

type RemoteChainConfigArgs struct {
	DestChainSelector uint64
	Config            adapters.ExecutorDestChainConfig
}

type ApplyAllowedCCVUpdatesArgs struct {
	CCVsToAdd        []common.Address
	CCVsToRemove     []common.Address
	AllowlistEnabled bool
}

type SetDynamicConfigArgs = executor.ExecutorDynamicConfig

var Deploy = contract.NewDeploy(contract.DeployParams[ConstructorArgs]{
	Name:             "executor:deploy",
	Version:          Version,
	Description:      "Deploys the Executor contract",
	ContractMetadata: executor.ExecutorMetaData,
	BytecodeByTypeAndVersion: map[string]contract.Bytecode{
		cldf_deployment.NewTypeAndVersion(ContractType, *Version).String(): {
			EVM: common.FromHex(executor.ExecutorBin),
		},
	},
	Validate: func(ConstructorArgs) error { return nil },
})

var DeployProxy = contract.NewDeploy(contract.DeployParams[ProxyConstructorArgs]{
	Name:             "executor-proxy:deploy",
	Version:          Version,
	Description:      "Deploys the ExecutorProxy contract",
	ContractMetadata: proxy.ProxyMetaData,
	BytecodeByTypeAndVersion: map[string]contract.Bytecode{
		cldf_deployment.NewTypeAndVersion(ProxyType, *Version).String(): {
			EVM: common.FromHex(proxy.ProxyBin),
		},
	},
	Validate: func(ProxyConstructorArgs) error { return nil },
})

var ApplyDestChainUpdates = contract.NewWrite(contract.WriteParams[ApplyDestChainUpdatesArgs, *executor.Executor]{
	Name:            "executor:apply-dest-chain-updates",
	Version:         Version,
	Description:     "Applies updates to supported destination chains on the Executor",
	ContractType:    ContractType,
	ContractABI:     executor.ExecutorABI,
	NewContract:     executor.NewExecutor,
	IsAllowedCaller: contract.OnlyOwner[*executor.Executor, ApplyDestChainUpdatesArgs],
	Validate: func(executor *executor.Executor, backend bind.ContractBackend, opts *bind.CallOpts, args ApplyDestChainUpdatesArgs) error {
		for _, destChainSelectorToAdd := range args.DestChainSelectorsToAdd {
			if destChainSelectorToAdd.DestChainSelector == 0 {
				return errors.New("dest chain selector cannot be 0")
			}
		}

		return nil
	},
	IsNoop: func(executor *executor.Executor, opts *bind.CallOpts, args ApplyDestChainUpdatesArgs) (bool, error) {
		activeDestChains, err := executor.GetDestChains(opts)
		if err != nil {
			return false, fmt.Errorf("failed to get dest chains: %w", err)
		}
		for _, activeDestChain := range activeDestChains {
			found := false
			for _, destChainToAdd := range args.DestChainSelectorsToAdd {
				if activeDestChain.DestChainSelector == destChainToAdd.DestChainSelector {
					found = true
					break
				}
			}
			if !found {
				return false, nil
			}
			for _, destChainToRemove := range args.DestChainSelectorsToRemove {
				if activeDestChain.DestChainSelector == destChainToRemove {
					return false, nil
				}
			}
		}

		return true, nil
	},
	CallContract: func(Executor *executor.Executor, opts *bind.TransactOpts, args ApplyDestChainUpdatesArgs) (*types.Transaction, error) {
		destChainSelectorsToAdd := make([]executor.ExecutorRemoteChainConfigArgs, 0, len(args.DestChainSelectorsToAdd))
		for _, destChainSelectorToAdd := range args.DestChainSelectorsToAdd {
			destChainSelectorsToAdd = append(destChainSelectorsToAdd, executor.ExecutorRemoteChainConfigArgs{
				DestChainSelector: destChainSelectorToAdd.DestChainSelector,
				Config: executor.ExecutorRemoteChainConfig{
					UsdCentsFee: destChainSelectorToAdd.Config.USDCentsFee,
					Enabled:     destChainSelectorToAdd.Config.Enabled,
				},
			})
		}
		return Executor.ApplyDestChainUpdates(opts, args.DestChainSelectorsToRemove, destChainSelectorsToAdd)
	},
})

var ApplyAllowedCCVUpdates = contract.NewWrite(contract.WriteParams[ApplyAllowedCCVUpdatesArgs, *executor.Executor]{
	Name:            "executor:apply-allowed-ccv-updates",
	Version:         Version,
	Description:     "Applies updates to the CCV allowlist on the Executor",
	ContractType:    ContractType,
	ContractABI:     executor.ExecutorABI,
	NewContract:     executor.NewExecutor,
	IsAllowedCaller: contract.OnlyOwner[*executor.Executor, ApplyAllowedCCVUpdatesArgs],
	Validate: func(executor *executor.Executor, backend bind.ContractBackend, opts *bind.CallOpts, args ApplyAllowedCCVUpdatesArgs) error {
		for _, ccv := range args.CCVsToAdd {
			if ccv == (common.Address{}) {
				return errors.New("CCV cannot be the zero address")
			}
		}

		return nil
	},
	IsNoop: func(executor *executor.Executor, opts *bind.CallOpts, args ApplyAllowedCCVUpdatesArgs) (bool, error) {
		allowedCCVs, err := executor.GetAllowedCCVs(opts)
		if err != nil {
			return false, fmt.Errorf("failed to get allowed CCVs: %w", err)
		}
		for _, ccv := range args.CCVsToAdd {
			if !slices.Contains(allowedCCVs, ccv) {
				return false, nil
			}
		}
		for _, ccv := range args.CCVsToRemove {
			if slices.Contains(allowedCCVs, ccv) {
				return false, nil
			}
		}
		return true, nil
	},
	CallContract: func(Executor *executor.Executor, opts *bind.TransactOpts, args ApplyAllowedCCVUpdatesArgs) (*types.Transaction, error) {
		return Executor.ApplyAllowedCCVUpdates(opts, args.CCVsToRemove, args.CCVsToAdd, args.AllowlistEnabled)
	},
})

var SetDynamicConfig = contract.NewWrite(contract.WriteParams[SetDynamicConfigArgs, *executor.Executor]{
	Name:            "executor:set-min-block-confirmations",
	Version:         Version,
	Description:     "Sets the minimum block confirmations on the Executor",
	ContractType:    ContractType,
	ContractABI:     executor.ExecutorABI,
	NewContract:     executor.NewExecutor,
	IsAllowedCaller: contract.OnlyOwner[*executor.Executor, SetDynamicConfigArgs],
	Validate: func(executor *executor.Executor, backend bind.ContractBackend, opts *bind.CallOpts, args SetDynamicConfigArgs) error {
		if args.FeeAggregator == (common.Address{}) {
			return errors.New("fee aggregator cannot be the zero address")
		}
		return nil
	},
	IsNoop: func(executor *executor.Executor, opts *bind.CallOpts, args SetDynamicConfigArgs) (bool, error) {
		currentDynamicConfig, err := executor.GetDynamicConfig(opts)
		if err != nil {
			return false, fmt.Errorf("failed to get dynamic configuration: %w", err)
		}
		return currentDynamicConfig == args, nil
	},
	CallContract: func(Executor *executor.Executor, opts *bind.TransactOpts, args SetDynamicConfigArgs) (*types.Transaction, error) {
		return Executor.SetDynamicConfig(opts, args)
	},
})

var GetDestChains = contract.NewRead(contract.ReadParams[any, []executor.ExecutorRemoteChainConfigArgs, *executor.Executor]{
	Name:         "executor:get-dest-chains",
	Version:      Version,
	Description:  "Gets the supported destination chains on the Executor",
	ContractType: ContractType,
	NewContract:  executor.NewExecutor,
	CallContract: func(Executor *executor.Executor, opts *bind.CallOpts, args any) ([]executor.ExecutorRemoteChainConfigArgs, error) {
		return Executor.GetDestChains(opts)
	},
})
