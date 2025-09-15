package commit_onramp

import (
	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/latest/commit_onramp"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/latest/ownable_ccv_ramp_proxy"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

var ContractType cldf_deployment.ContractType = "CommitOnRamp"
var ProxyType cldf_deployment.ContractType = "CommitOnRampProxy"

type DynamicConfig = commit_onramp.CommitOnRampDynamicConfig

type ConstructorArgs struct {
	RMNRemote     common.Address
	NonceManager  common.Address
	DynamicConfig DynamicConfig
}

type SetDynamicConfigArgs struct {
	DynamicConfig DynamicConfig
}

type DestChainConfigArgs = commit_onramp.BaseOnRampDestChainConfigArgs

type AllowlistConfigArgs = commit_onramp.BaseOnRampAllowlistConfigArgs

type WithdrawFeeTokensArgs struct {
	FeeTokens []common.Address
}

type DestChainConfig = commit_onramp.GetDestChainConfig

var Deploy = contract.NewDeploy(
	"commit-onramp:deploy",
	semver.MustParse("1.7.0"),
	"Deploys the CommitOnRamp contract",
	ContractType,
	commit_onramp.CommitOnRampABI,
	func(ConstructorArgs) error { return nil },
	contract.VMDeployers[ConstructorArgs]{
		DeployEVM: func(opts *bind.TransactOpts, backend bind.ContractBackend, args ConstructorArgs) (common.Address, *types.Transaction, error) {
			address, tx, _, err := commit_onramp.DeployCommitOnRamp(opts, backend, args.RMNRemote, args.NonceManager, args.DynamicConfig)
			return address, tx, err
		},
		// DeployZksyncVM: func(opts *accounts.TransactOpts, client *clients.Client, wallet *accounts.Wallet, backend bind.ContractBackend, args ConstructorArgs) (common.Address, error)
	},
)

var DeployProxy = contract.NewDeploy(
	"commit-off-ramp-proxy:deploy",
	semver.MustParse("1.7.0"),
	"Deploys the CommitOffRampProxy contract",
	ProxyType,
	ownable_ccv_ramp_proxy.OwnableRampProxyABI,
	func(any) error { return nil },
	contract.VMDeployers[any]{
		DeployEVM: func(opts *bind.TransactOpts, backend bind.ContractBackend, args any) (common.Address, *types.Transaction, error) {
			address, tx, _, err := ownable_ccv_ramp_proxy.DeployOwnableRampProxy(opts, backend)
			return address, tx, err
		},
		// DeployZksyncVM: func(opts *accounts.TransactOpts, client *clients.Client, wallet *accounts.Wallet, backend bind.ContractBackend, args any) (common.Address, error)
	},
)

var SetDynamicConfig = contract.NewWrite(
	"commit-onramp:set-dynamic-config",
	semver.MustParse("1.7.0"),
	"Sets the dynamic configuration on the CommitOnRamp",
	ContractType,
	commit_onramp.CommitOnRampABI,
	commit_onramp.NewCommitOnRamp,
	contract.OnlyOwner,
	func(SetDynamicConfigArgs) error { return nil },
	func(commitOnRamp *commit_onramp.CommitOnRamp, opts *bind.TransactOpts, args SetDynamicConfigArgs) (*types.Transaction, error) {
		return commitOnRamp.SetDynamicConfig(opts, args.DynamicConfig)
	},
)

var ApplyDestChainConfigUpdates = contract.NewWrite(
	"commit-onramp:apply-dest-chain-config-updates",
	semver.MustParse("1.7.0"),
	"Applies updates to destination chain configurations on the CommitOnRamp",
	ContractType,
	commit_onramp.CommitOnRampABI,
	commit_onramp.NewCommitOnRamp,
	contract.OnlyOwner,
	func([]DestChainConfigArgs) error { return nil },
	func(commitOnRamp *commit_onramp.CommitOnRamp, opts *bind.TransactOpts, args []DestChainConfigArgs) (*types.Transaction, error) {
		return commitOnRamp.ApplyDestChainConfigUpdates(opts, args)
	},
)

var ApplyAllowlistUpdates = contract.NewWrite(
	"commit-onramp:apply-allowlist-updates",
	semver.MustParse("1.7.0"),
	"Applies updates to the allowlist (those authorized to send messages) on the CommitOnRamp",
	ContractType,
	commit_onramp.CommitOnRampABI,
	commit_onramp.NewCommitOnRamp,
	contract.OnlyOwner,
	func([]AllowlistConfigArgs) error { return nil },
	func(commitOnRamp *commit_onramp.CommitOnRamp, opts *bind.TransactOpts, args []AllowlistConfigArgs) (*types.Transaction, error) {
		return commitOnRamp.ApplyAllowlistUpdates(opts, args)
	},
)

var WithdrawFeeTokens = contract.NewWrite(
	"commit-onramp:withdraw-fee-tokens",
	semver.MustParse("1.7.0"),
	"Withdraws fee tokens from the CommitOnRamp",
	ContractType,
	commit_onramp.CommitOnRampABI,
	commit_onramp.NewCommitOnRamp,
	contract.OnlyOwner,
	func(WithdrawFeeTokensArgs) error { return nil },
	func(commitOnRamp *commit_onramp.CommitOnRamp, opts *bind.TransactOpts, args WithdrawFeeTokensArgs) (*types.Transaction, error) {
		return commitOnRamp.WithdrawFeeTokens(opts, args.FeeTokens)
	},
)

var GetDestChainConfig = contract.NewRead(
	"commit-onramp:get-dest-chain-config",
	semver.MustParse("1.7.0"),
	"Gets the destination chain configuration for a given destination chain selector",
	ContractType,
	commit_onramp.NewCommitOnRamp,
	func(commitOnRamp *commit_onramp.CommitOnRamp, opts *bind.CallOpts, destChainSelector uint64) (DestChainConfig, error) {
		return commitOnRamp.GetDestChainConfig(opts, destChainSelector)
	},
)
