package executor_onramp

import (
	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/call"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/deployment"
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

var Deploy = deployment.New(
	"executor-onramp:deploy",
	semver.MustParse("1.7.0"),
	"Deploys the ExecutorOnRamp contract",
	ContractType,
	executor_onramp.ExecutorOnRampABI,
	func(ConstructorArgs) error { return nil },
	deployment.VMDeployers[ConstructorArgs]{
		DeployEVM: func(opts *bind.TransactOpts, backend bind.ContractBackend, args ConstructorArgs) (common.Address, *types.Transaction, error) {
			address, tx, _, err := executor_onramp.DeployExecutorOnRamp(opts, backend, args.MaxCCVsPerMsg)
			return address, tx, err
		},
		// DeployZksyncVM: func(opts *accounts.TransactOpts, client *clients.Client, wallet *accounts.Wallet, backend bind.ContractBackend, args ConstructorArgs) (common.Address, error)
	},
)

var SetMaxCCVsPerMsg = call.NewWrite(
	"executor-onramp:set-max-ccvs-per-msg",
	semver.MustParse("1.7.0"),
	"Sets the maximum number of CCVs per message on the ExecutorOnRamp",
	ContractType,
	executor_onramp.ExecutorOnRampABI,
	executor_onramp.NewExecutorOnRamp,
	call.OnlyOwner,
	func(uint8) error { return nil },
	func(executorOnRamp *executor_onramp.ExecutorOnRamp, opts *bind.TransactOpts, args uint8) (*types.Transaction, error) {
		return executorOnRamp.SetMaxCCVsPerMsg(opts, args)
	},
)

var ApplyDestChainUpdates = call.NewWrite(
	"executor-onramp:apply-dest-chain-updates",
	semver.MustParse("1.7.0"),
	"Applies updates to supported destination chains on the ExecutorOnRamp",
	ContractType,
	executor_onramp.ExecutorOnRampABI,
	executor_onramp.NewExecutorOnRamp,
	call.OnlyOwner,
	func(ApplyDestChainUpdatesArgs) error { return nil },
	func(executorOnRamp *executor_onramp.ExecutorOnRamp, opts *bind.TransactOpts, args ApplyDestChainUpdatesArgs) (*types.Transaction, error) {
		return executorOnRamp.ApplyDestChainUpdates(opts, args.DestChainSelectorsToAdd, args.DestChainSelectorsToRemove)
	},
)

var ApplyAllowedCCVUpdates = call.NewWrite(
	"executor-onramp:apply-allowed-ccv-updates",
	semver.MustParse("1.7.0"),
	"Applies updates to the CCV allowlist on the ExecutorOnRamp",
	ContractType,
	executor_onramp.ExecutorOnRampABI,
	executor_onramp.NewExecutorOnRamp,
	call.OnlyOwner,
	func(ApplyAllowedCCVUpdatesArgs) error { return nil },
	func(executorOnRamp *executor_onramp.ExecutorOnRamp, opts *bind.TransactOpts, args ApplyAllowedCCVUpdatesArgs) (*types.Transaction, error) {
		return executorOnRamp.ApplyAllowedCCVUpdates(opts, args.CCVsToAdd, args.CCVsToRemove, args.AllowlistEnabled)
	},
)
