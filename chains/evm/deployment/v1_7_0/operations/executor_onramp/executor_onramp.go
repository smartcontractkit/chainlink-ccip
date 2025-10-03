package executor_onramp

import (
	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/latest/executor_onramp"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

var ContractType cldf_deployment.ContractType = "ExecutorOnRamp"

type ConstructorArgs struct {
	MaxCCVsPerMsg uint8
}

type ApplyDestChainUpdatesArgs struct {
	DestChainSelectorsToAdd    []uint64
	DestChainSelectorsToRemove []uint64
}

type ApplyAllowedCCVUpdatesArgs struct {
	CCVsToAdd        []common.Address
	CCVsToRemove     []common.Address
	AllowlistEnabled bool
}

var Deploy = contract.NewDeploy(contract.DeployParams[ConstructorArgs]{
	Name:             "executor-onramp:deploy",
	Version:          semver.MustParse("1.7.0"),
	Description:      "Deploys the ExecutorOnRamp contract",
	ContractMetadata: executor_onramp.ExecutorOnRampMetaData,
	BytecodeByTypeAndVersion: map[string]contract.Bytecode{
		cldf_deployment.NewTypeAndVersion(ContractType, *semver.MustParse("1.7.0")).String(): {
			EVM: common.FromHex(executor_onramp.ExecutorOnRampBin),
		},
	},
	Validate: func(ConstructorArgs) error { return nil },
})

var SetMaxCCVsPerMsg = contract.NewWrite(contract.WriteParams[uint8, *executor_onramp.ExecutorOnRamp]{
	Name:            "executor-onramp:set-max-ccvs-per-msg",
	Version:         semver.MustParse("1.7.0"),
	Description:     "Sets the maximum number of CCVs per message on the ExecutorOnRamp",
	ContractType:    ContractType,
	ContractABI:     executor_onramp.ExecutorOnRampABI,
	NewContract:     executor_onramp.NewExecutorOnRamp,
	IsAllowedCaller: contract.OnlyOwner[*executor_onramp.ExecutorOnRamp, uint8],
	Validate:        func(uint8) error { return nil },
	CallContract: func(executorOnRamp *executor_onramp.ExecutorOnRamp, opts *bind.TransactOpts, args uint8) (*types.Transaction, error) {
		return executorOnRamp.SetMaxCCVsPerMsg(opts, args)
	},
})

var ApplyDestChainUpdates = contract.NewWrite(contract.WriteParams[ApplyDestChainUpdatesArgs, *executor_onramp.ExecutorOnRamp]{
	Name:            "executor-onramp:apply-dest-chain-updates",
	Version:         semver.MustParse("1.7.0"),
	Description:     "Applies updates to supported destination chains on the ExecutorOnRamp",
	ContractType:    ContractType,
	ContractABI:     executor_onramp.ExecutorOnRampABI,
	NewContract:     executor_onramp.NewExecutorOnRamp,
	IsAllowedCaller: contract.OnlyOwner[*executor_onramp.ExecutorOnRamp, ApplyDestChainUpdatesArgs],
	Validate:        func(ApplyDestChainUpdatesArgs) error { return nil },
	CallContract: func(executorOnRamp *executor_onramp.ExecutorOnRamp, opts *bind.TransactOpts, args ApplyDestChainUpdatesArgs) (*types.Transaction, error) {
		return executorOnRamp.ApplyDestChainUpdates(opts, args.DestChainSelectorsToRemove, args.DestChainSelectorsToAdd)
	},
})

var ApplyAllowedCCVUpdates = contract.NewWrite(contract.WriteParams[ApplyAllowedCCVUpdatesArgs, *executor_onramp.ExecutorOnRamp]{
	Name:            "executor-onramp:apply-allowed-ccv-updates",
	Version:         semver.MustParse("1.7.0"),
	Description:     "Applies updates to the CCV allowlist on the ExecutorOnRamp",
	ContractType:    ContractType,
	ContractABI:     executor_onramp.ExecutorOnRampABI,
	NewContract:     executor_onramp.NewExecutorOnRamp,
	IsAllowedCaller: contract.OnlyOwner[*executor_onramp.ExecutorOnRamp, ApplyAllowedCCVUpdatesArgs],
	Validate:        func(ApplyAllowedCCVUpdatesArgs) error { return nil },
	CallContract: func(executorOnRamp *executor_onramp.ExecutorOnRamp, opts *bind.TransactOpts, args ApplyAllowedCCVUpdatesArgs) (*types.Transaction, error) {
		return executorOnRamp.ApplyAllowedCCVUpdates(opts, args.CCVsToRemove, args.CCVsToAdd, args.AllowlistEnabled)
	},
})

var GetDestChains = contract.NewRead(contract.ReadParams[any, []uint64, *executor_onramp.ExecutorOnRamp]{
	Name:         "executor-onramp:get-dest-chains",
	Version:      semver.MustParse("1.7.0"),
	Description:  "Gets the supported destination chains on the ExecutorOnRamp",
	ContractType: ContractType,
	NewContract:  executor_onramp.NewExecutorOnRamp,
	CallContract: func(executorOnRamp *executor_onramp.ExecutorOnRamp, opts *bind.CallOpts, args any) ([]uint64, error) {
		return executorOnRamp.GetDestChains(opts)
	},
})
