package commit_ramp

import (
	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/latest/commit_ramp"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/latest/ownable_ramp_proxy"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

var ContractType cldf_deployment.ContractType = "CommitRamp"
var ProxyType cldf_deployment.ContractType = "CommitRampProxy"

type DynamicConfig = commit_ramp.CommitRampDynamicConfig

type ConstructorArgs struct {
	DynamicConfig DynamicConfig
}

type SetDynamicConfigArgs struct {
	DynamicConfig DynamicConfig
}

type DestChainConfigArgs = commit_ramp.BaseOnRampDestChainConfigArgs

type AllowlistConfigArgs = commit_ramp.BaseOnRampAllowlistConfigArgs

type WithdrawFeeTokensArgs struct {
	FeeTokens []common.Address
}

type DestChainConfig = commit_ramp.GetDestChainConfig

type SetSignatureConfigArgs struct {
	Signers   []common.Address
	Threshold uint8
}

var Deploy = contract.NewDeploy(
	"commit-ramp:deploy",
	semver.MustParse("1.7.0"),
	"Deploys the CommitRamp contract",
	ContractType,
	commit_ramp.CommitRampABI,
	func(ConstructorArgs) error { return nil },
	contract.VMDeployers[ConstructorArgs]{
		DeployEVM: func(opts *bind.TransactOpts, backend bind.ContractBackend, args ConstructorArgs) (common.Address, *types.Transaction, error) {
			address, tx, _, err := commit_ramp.DeployCommitRamp(opts, backend, args.DynamicConfig)
			return address, tx, err
		},
		// DeployZksyncVM: func(opts *accounts.TransactOpts, client *clients.Client, wallet *accounts.Wallet, backend bind.ContractBackend, args ConstructorArgs) (common.Address, error)
	},
)

var DeployProxy = contract.NewDeploy(
	"commit-ramp-proxy:deploy",
	semver.MustParse("1.7.0"),
	"Deploys the CommitRampProxy contract",
	ProxyType,
	ownable_ramp_proxy.OwnableRampProxyABI,
	func(common.Address) error { return nil },
	contract.VMDeployers[common.Address]{
		DeployEVM: func(opts *bind.TransactOpts, backend bind.ContractBackend, rampAddress common.Address) (common.Address, *types.Transaction, error) {
			address, tx, _, err := ownable_ramp_proxy.DeployOwnableRampProxy(opts, backend, rampAddress)
			return address, tx, err
		},
		// DeployZksyncVM: func(opts *accounts.TransactOpts, client *clients.Client, wallet *accounts.Wallet, backend bind.ContractBackend, rampAddress common.Address) (common.Address, error)
	},
)

var SetDynamicConfig = contract.NewWrite(
	"commit-ramp:set-dynamic-config",
	semver.MustParse("1.7.0"),
	"Sets the dynamic configuration on the CommitRamp",
	ContractType,
	commit_ramp.CommitRampABI,
	commit_ramp.NewCommitRamp,
	contract.OnlyOwner,
	func(SetDynamicConfigArgs) error { return nil },
	func(commitRamp *commit_ramp.CommitRamp, opts *bind.TransactOpts, args SetDynamicConfigArgs) (*types.Transaction, error) {
		return commitRamp.SetDynamicConfig(opts, args.DynamicConfig)
	},
)

var ApplyDestChainConfigUpdates = contract.NewWrite(
	"commit-ramp:apply-dest-chain-config-updates",
	semver.MustParse("1.7.0"),
	"Applies updates to destination chain configurations on the CommitRamp",
	ContractType,
	commit_ramp.CommitRampABI,
	commit_ramp.NewCommitRamp,
	contract.OnlyOwner,
	func([]DestChainConfigArgs) error { return nil },
	func(commitRamp *commit_ramp.CommitRamp, opts *bind.TransactOpts, args []DestChainConfigArgs) (*types.Transaction, error) {
		return commitRamp.ApplyDestChainConfigUpdates(opts, args)
	},
)

var ApplyAllowlistUpdates = contract.NewWrite(
	"commit-ramp:apply-allowlist-updates",
	semver.MustParse("1.7.0"),
	"Applies updates to the allowlist (those authorized to send messages) on the CommitRamp",
	ContractType,
	commit_ramp.CommitRampABI,
	commit_ramp.NewCommitRamp,
	contract.OnlyOwner,
	func([]AllowlistConfigArgs) error { return nil },
	func(commitRamp *commit_ramp.CommitRamp, opts *bind.TransactOpts, args []AllowlistConfigArgs) (*types.Transaction, error) {
		return commitRamp.ApplyAllowlistUpdates(opts, args)
	},
)

var WithdrawFeeTokens = contract.NewWrite(
	"commit-ramp:withdraw-fee-tokens",
	semver.MustParse("1.7.0"),
	"Withdraws fee tokens from the CommitRamp",
	ContractType,
	commit_ramp.CommitRampABI,
	commit_ramp.NewCommitRamp,
	contract.OnlyOwner,
	func(WithdrawFeeTokensArgs) error { return nil },
	func(commitRamp *commit_ramp.CommitRamp, opts *bind.TransactOpts, args WithdrawFeeTokensArgs) (*types.Transaction, error) {
		return commitRamp.WithdrawFeeTokens(opts, args.FeeTokens)
	},
)

var GetDestChainConfig = contract.NewRead(
	"commit-ramp:get-dest-chain-config",
	semver.MustParse("1.7.0"),
	"Gets the destination chain configuration for a given destination chain selector",
	ContractType,
	commit_ramp.NewCommitRamp,
	func(commitRamp *commit_ramp.CommitRamp, opts *bind.CallOpts, destChainSelector uint64) (DestChainConfig, error) {
		return commitRamp.GetDestChainConfig(opts, destChainSelector)
	},
)

var SetSignatureConfigs = contract.NewWrite(
	"commit-ramp:set-signature-config",
	semver.MustParse("1.7.0"),
	"Sets the signature configuration on the CommitRamp",
	ContractType,
	commit_ramp.CommitRampABI,
	commit_ramp.NewCommitRamp,
	contract.OnlyOwner,
	func(SetSignatureConfigArgs) error { return nil },
	func(commitRamp *commit_ramp.CommitRamp, opts *bind.TransactOpts, args SetSignatureConfigArgs) (*types.Transaction, error) {
		return commitRamp.SetSignatureConfig(opts, args.Signers, args.Threshold)
	},
)
