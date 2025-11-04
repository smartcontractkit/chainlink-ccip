package executor

import (
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
	Version:          semver.MustParse("1.7.0"),
	Description:      "Deploys the Executor contract",
	ContractMetadata: executor.ExecutorMetaData,
	BytecodeByTypeAndVersion: map[string]contract.Bytecode{
		cldf_deployment.NewTypeAndVersion(ContractType, *semver.MustParse("1.7.0")).String(): {
			EVM: common.FromHex(executor.ExecutorBin),
		},
	},
	Validate: func(ConstructorArgs) error { return nil },
})

var DeployProxy = contract.NewDeploy(contract.DeployParams[ProxyConstructorArgs]{
	Name:             "executor-proxy:deploy",
	Version:          semver.MustParse("1.7.0"),
	Description:      "Deploys the ExecutorProxy contract",
	ContractMetadata: proxy.ProxyMetaData,
	BytecodeByTypeAndVersion: map[string]contract.Bytecode{
		cldf_deployment.NewTypeAndVersion(ProxyType, *semver.MustParse("1.7.0")).String(): {
			EVM: common.FromHex(proxy.ProxyBin),
		},
	},
	Validate: func(ProxyConstructorArgs) error { return nil },
})

var ApplyDestChainUpdates = contract.NewWrite(contract.WriteParams[ApplyDestChainUpdatesArgs, *executor.Executor]{
	Name:            "executor:apply-dest-chain-updates",
	Version:         semver.MustParse("1.7.0"),
	Description:     "Applies updates to supported destination chains on the Executor",
	ContractType:    ContractType,
	ContractABI:     executor.ExecutorABI,
	NewContract:     executor.NewExecutor,
	IsAllowedCaller: contract.OnlyOwner[*executor.Executor, ApplyDestChainUpdatesArgs],
	Validate:        func(ApplyDestChainUpdatesArgs) error { return nil },
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
	Version:         semver.MustParse("1.7.0"),
	Description:     "Applies updates to the CCV allowlist on the Executor",
	ContractType:    ContractType,
	ContractABI:     executor.ExecutorABI,
	NewContract:     executor.NewExecutor,
	IsAllowedCaller: contract.OnlyOwner[*executor.Executor, ApplyAllowedCCVUpdatesArgs],
	Validate:        func(ApplyAllowedCCVUpdatesArgs) error { return nil },
	CallContract: func(Executor *executor.Executor, opts *bind.TransactOpts, args ApplyAllowedCCVUpdatesArgs) (*types.Transaction, error) {
		return Executor.ApplyAllowedCCVUpdates(opts, args.CCVsToRemove, args.CCVsToAdd, args.AllowlistEnabled)
	},
})

var SetDynamicConfig = contract.NewWrite(contract.WriteParams[SetDynamicConfigArgs, *executor.Executor]{
	Name:            "executor:set-min-block-confirmations",
	Version:         semver.MustParse("1.7.0"),
	Description:     "Sets the minimum block confirmations on the Executor",
	ContractType:    ContractType,
	ContractABI:     executor.ExecutorABI,
	NewContract:     executor.NewExecutor,
	IsAllowedCaller: contract.OnlyOwner[*executor.Executor, SetDynamicConfigArgs],
	Validate:        func(SetDynamicConfigArgs) error { return nil },
	CallContract: func(Executor *executor.Executor, opts *bind.TransactOpts, args SetDynamicConfigArgs) (*types.Transaction, error) {
		return Executor.SetDynamicConfig(opts, args)
	},
})

var GetDestChains = contract.NewRead(contract.ReadParams[any, []executor.ExecutorRemoteChainConfigArgs, *executor.Executor]{
	Name:         "executor:get-dest-chains",
	Version:      semver.MustParse("1.7.0"),
	Description:  "Gets the supported destination chains on the Executor",
	ContractType: ContractType,
	NewContract:  executor.NewExecutor,
	CallContract: func(Executor *executor.Executor, opts *bind.CallOpts, args any) ([]executor.ExecutorRemoteChainConfigArgs, error) {
		return Executor.GetDestChains(opts)
	},
})
