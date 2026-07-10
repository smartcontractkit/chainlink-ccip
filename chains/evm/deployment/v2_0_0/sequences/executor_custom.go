package sequences

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"

	gobindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v2_0_0/executor"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/executor"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v2_0_0/adapters"
	cldf_evm "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations2/contract"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cld_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

var ExecutorProxyType cldf_deployment.ContractType = "ExecutorProxy"

type ExecutorRemoteChainConfigArgs struct {
	DestChainSelector uint64
	Config            adapters.ExecutorDestChainConfig
}

type ExecutorApplyDestChainUpdatesArgs struct {
	DestChainSelectorsToAdd    []ExecutorRemoteChainConfigArgs
	DestChainSelectorsToRemove []uint64
}

func NewWriteExecutorApplyDestChainUpdates(c gobindings.ExecutorInterface) *cld_ops.Operation[contract.FunctionInput[ExecutorApplyDestChainUpdatesArgs], contract.WriteOutput, cldf_evm.Chain] {
	return contract.NewWrite(contract.WriteParams[ExecutorApplyDestChainUpdatesArgs, gobindings.ExecutorInterface]{
		Name:            "executor:apply-dest-chain-updates",
		Version:         executor.Version,
		Description:     "Applies updates to supported destination chains on the Executor",
		ContractType:    executor.ContractType,
		ContractABI:     gobindings.ExecutorABI,
		Contract:        c,
		IsAllowedCaller: contract.OnlyOwner[gobindings.ExecutorInterface, ExecutorApplyDestChainUpdatesArgs],
		Validate:        func(ExecutorApplyDestChainUpdatesArgs) error { return nil },
		CallContract: func(e gobindings.ExecutorInterface, opts *bind.TransactOpts, args ExecutorApplyDestChainUpdatesArgs) (*types.Transaction, error) {
			destChainSelectorsToAdd := make([]gobindings.ExecutorRemoteChainConfigArgs, 0, len(args.DestChainSelectorsToAdd))
			for _, destChainSelectorToAdd := range args.DestChainSelectorsToAdd {
				destChainSelectorsToAdd = append(destChainSelectorsToAdd, gobindings.ExecutorRemoteChainConfigArgs{
					DestChainSelector: destChainSelectorToAdd.DestChainSelector,
					Config: gobindings.ExecutorRemoteChainConfig{
						UsdCentsFee: destChainSelectorToAdd.Config.USDCentsFee,
						Enabled:     destChainSelectorToAdd.Config.Enabled,
					},
				})
			}
			return e.ApplyDestChainUpdates(opts, args.DestChainSelectorsToRemove, destChainSelectorsToAdd)
		},
	})
}
