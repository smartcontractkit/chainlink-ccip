package sequences

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"

	executor_bindings "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/gobindings/generated/latest/executor"

	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v2_0_0/operations/executor"
	contract_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/deployment/lanes"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

var ExecutorProxyType cldf_deployment.ContractType = "ExecutorProxy"

type ExecutorRemoteChainConfigArgs struct {
	DestChainSelector uint64
	Config            lanes.ExecutorDestChainConfig
}

type ExecutorApplyDestChainUpdatesArgs struct {
	DestChainSelectorsToAdd    []ExecutorRemoteChainConfigArgs
	DestChainSelectorsToRemove []uint64
}

var ExecutorApplyDestChainUpdates = contract_utils.NewWrite(contract_utils.WriteParams[ExecutorApplyDestChainUpdatesArgs, *executor_bindings.Executor]{
	Name:            "executor:apply-dest-chain-updates",
	Version:         executor.Version,
	Description:     "Applies updates to supported destination chains on the Executor",
	ContractType:    executor.ContractType,
	ContractABI:     executor.ExecutorABI,
	NewContract:     executor_bindings.NewExecutor,
	IsAllowedCaller: contract_utils.OnlyOwner[*executor_bindings.Executor, ExecutorApplyDestChainUpdatesArgs],
	Validate:        func(ExecutorApplyDestChainUpdatesArgs) error { return nil },
	CallContract: func(e *executor_bindings.Executor, opts *bind.TransactOpts, args ExecutorApplyDestChainUpdatesArgs) (*types.Transaction, error) {
		destChainSelectorsToAdd := make([]executor_bindings.ExecutorRemoteChainConfigArgs, 0, len(args.DestChainSelectorsToAdd))
		for _, destChainSelectorToAdd := range args.DestChainSelectorsToAdd {
			destChainSelectorsToAdd = append(destChainSelectorsToAdd, executor_bindings.ExecutorRemoteChainConfigArgs{
				DestChainSelector: destChainSelectorToAdd.DestChainSelector,
				Config: executor_bindings.ExecutorRemoteChainConfig{
					UsdCentsFee: destChainSelectorToAdd.Config.USDCentsFee,
					Enabled:     destChainSelectorToAdd.Config.Enabled,
				},
			})
		}
		return e.ApplyDestChainUpdates(opts, args.DestChainSelectorsToRemove, destChainSelectorsToAdd)
	},
})
