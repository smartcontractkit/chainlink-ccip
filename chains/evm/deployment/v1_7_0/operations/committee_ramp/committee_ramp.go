package committee_ramp

import (
	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/latest/committee_ramp"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/latest/ramp_proxy"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

var ContractType cldf_deployment.ContractType = "CommitteeRamp"

var ProxyType cldf_deployment.ContractType = "CommitteeRampProxy"

type DynamicConfig = committee_ramp.CommitteeRampDynamicConfig

type ConstructorArgs struct {
	DynamicConfig DynamicConfig
}

type SetDynamicConfigArgs struct {
	DynamicConfig DynamicConfig
}

type DestChainConfigArgs = committee_ramp.BaseOnRampDestChainConfigArgs

type AllowlistConfigArgs = committee_ramp.BaseOnRampAllowlistConfigArgs

type WithdrawFeeTokensArgs struct {
	FeeTokens []common.Address
}

type DestChainConfig = committee_ramp.GetDestChainConfig

type SetSignatureConfigArgs struct {
	Signers   []common.Address
	Threshold uint8
}

var Deploy = contract.NewDeploy(
	"committee-ramp:deploy",
	semver.MustParse("1.7.0"),
	"Deploys the CommitteeRamp contract",
	ContractType,
	committee_ramp.CommitteeRampABI,
	func(ConstructorArgs) error { return nil },
	contract.VMDeployers[ConstructorArgs]{
		DeployEVM: func(opts *bind.TransactOpts, backend bind.ContractBackend, args ConstructorArgs) (common.Address, *types.Transaction, error) {
			address, tx, _, err := committee_ramp.DeployCommitteeRamp(opts, backend, args.DynamicConfig)
			return address, tx, err
		},
		// DeployZksyncVM: func(opts *accounts.TransactOpts, client *clients.Client, wallet *accounts.Wallet, backend bind.ContractBackend, args ConstructorArgs) (common.Address, error)
	},
)

var DeployProxy = contract.NewDeploy(
	"committee-ramp-proxy:deploy",
	semver.MustParse("1.7.0"),
	"Deploys the CommitteeRampProxy contract",
	ProxyType,
	ramp_proxy.RampProxyABI,
	func(common.Address) error { return nil },
	contract.VMDeployers[common.Address]{
		DeployEVM: func(opts *bind.TransactOpts, backend bind.ContractBackend, rampAddress common.Address) (common.Address, *types.Transaction, error) {
			address, tx, _, err := ramp_proxy.DeployRampProxy(opts, backend, rampAddress)
			return address, tx, err
		},
		// DeployZksyncVM: func(opts *accounts.TransactOpts, client *clients.Client, wallet *accounts.Wallet, backend bind.ContractBackend, rampAddress common.Address) (common.Address, error)
	},
)

var SetDynamicConfig = contract.NewWrite(
	"committee-ramp:set-dynamic-config",
	semver.MustParse("1.7.0"),
	"Sets the dynamic configuration on the CommitteeRamp",
	ContractType,
	committee_ramp.CommitteeRampABI,
	committee_ramp.NewCommitteeRamp,
	contract.OnlyOwner,
	func(SetDynamicConfigArgs) error { return nil },
	func(committeeRamp *committee_ramp.CommitteeRamp, opts *bind.TransactOpts, args SetDynamicConfigArgs) (*types.Transaction, error) {
		return committeeRamp.SetDynamicConfig(opts, args.DynamicConfig)
	},
)

var ApplyDestChainConfigUpdates = contract.NewWrite(
	"committee-ramp:apply-dest-chain-config-updates",
	semver.MustParse("1.7.0"),
	"Applies updates to destination chain configurations on the CommitteeRamp",
	ContractType,
	committee_ramp.CommitteeRampABI,
	committee_ramp.NewCommitteeRamp,
	contract.OnlyOwner,
	func([]DestChainConfigArgs) error { return nil },
	func(committeeRamp *committee_ramp.CommitteeRamp, opts *bind.TransactOpts, args []DestChainConfigArgs) (*types.Transaction, error) {
		return committeeRamp.ApplyDestChainConfigUpdates(opts, args)
	},
)

var ApplyAllowlistUpdates = contract.NewWrite(
	"committee-ramp:apply-allowlist-updates",
	semver.MustParse("1.7.0"),
	"Applies updates to the allowlist (those authorized to send messages) on the CommitteeRamp",
	ContractType,
	committee_ramp.CommitteeRampABI,
	committee_ramp.NewCommitteeRamp,
	contract.OnlyOwner,
	func([]AllowlistConfigArgs) error { return nil },
	func(committeeRamp *committee_ramp.CommitteeRamp, opts *bind.TransactOpts, args []AllowlistConfigArgs) (*types.Transaction, error) {
		return committeeRamp.ApplyAllowlistUpdates(opts, args)
	},
)

var WithdrawFeeTokens = contract.NewWrite(
	"committee-ramp:withdraw-fee-tokens",
	semver.MustParse("1.7.0"),
	"Withdraws fee tokens from the CommitteeRamp",
	ContractType,
	committee_ramp.CommitteeRampABI,
	committee_ramp.NewCommitteeRamp,
	contract.OnlyOwner,
	func(WithdrawFeeTokensArgs) error { return nil },
	func(committeeRamp *committee_ramp.CommitteeRamp, opts *bind.TransactOpts, args WithdrawFeeTokensArgs) (*types.Transaction, error) {
		return committeeRamp.WithdrawFeeTokens(opts, args.FeeTokens)
	},
)

var GetDestChainConfig = contract.NewRead(
	"committee-ramp:get-dest-chain-config",
	semver.MustParse("1.7.0"),
	"Gets the destination chain configuration for a given destination chain selector",
	ContractType,
	committee_ramp.NewCommitteeRamp,
	func(committeeRamp *committee_ramp.CommitteeRamp, opts *bind.CallOpts, destChainSelector uint64) (DestChainConfig, error) {
		return committeeRamp.GetDestChainConfig(opts, destChainSelector)
	},
)

var SetSignatureConfigs = contract.NewWrite(
	"committee-ramp:set-signature-config",
	semver.MustParse("1.7.0"),
	"Sets the signature configuration on the CommitteeRamp",
	ContractType,
	committee_ramp.CommitteeRampABI,
	committee_ramp.NewCommitteeRamp,
	contract.OnlyOwner,
	func(SetSignatureConfigArgs) error { return nil },
	func(committeeRamp *committee_ramp.CommitteeRamp, opts *bind.TransactOpts, args SetSignatureConfigArgs) (*types.Transaction, error) {
		return committeeRamp.SetSignatureConfig(opts, args.Signers, args.Threshold)
	},
)
