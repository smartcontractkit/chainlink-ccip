package executor

import (
	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/latest/executor"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

var ContractType cldf_deployment.ContractType = "Executor"

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

var SetMaxCCVsPerMsg = contract.NewWrite(contract.WriteParams[uint8, *executor.Executor]{
	Name:            "executor:set-max-ccvs-per-msg",
	Version:         semver.MustParse("1.7.0"),
	Description:     "Sets the maximum number of CCVs per message on the Executor",
	ContractType:    ContractType,
	ContractABI:     executor.ExecutorABI,
	NewContract:     executor.NewExecutor,
	IsAllowedCaller: contract.OnlyOwner[*executor.Executor, uint8],
	Validate:        func(uint8) error { return nil },
	CallContract: func(Executor *executor.Executor, opts *bind.TransactOpts, args uint8) (*types.Transaction, error) {
		return Executor.SetMaxCCVsPerMsg(opts, args)
	},
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
		return Executor.ApplyDestChainUpdates(opts, args.DestChainSelectorsToRemove, args.DestChainSelectorsToAdd)
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

var GetDestChains = contract.NewRead(contract.ReadParams[any, []uint64, *executor.Executor]{
	Name:         "executor:get-dest-chains",
	Version:      semver.MustParse("1.7.0"),
	Description:  "Gets the supported destination chains on the Executor",
	ContractType: ContractType,
	NewContract:  executor.NewExecutor,
	CallContract: func(Executor *executor.Executor, opts *bind.CallOpts, args any) ([]uint64, error) {
		return Executor.GetDestChains(opts)
	},
})
