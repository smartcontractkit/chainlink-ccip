package committee_verifier

import (
	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/latest/committee_verifier"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/latest/verifier_proxy"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

var ContractType cldf_deployment.ContractType = "CommitteeVerifier"

var ProxyType cldf_deployment.ContractType = "CommitteeVerifierProxy"

type DynamicConfig = committee_verifier.CommitteeVerifierDynamicConfig

type ConstructorArgs struct {
	DynamicConfig   DynamicConfig
	StorageLocation string
}

type SetDynamicConfigArgs struct {
	DynamicConfig DynamicConfig
}

type DestChainConfigArgs = committee_verifier.BaseVerifierDestChainConfigArgs

type AllowlistConfigArgs = committee_verifier.BaseVerifierAllowlistConfigArgs

type WithdrawFeeTokensArgs struct {
	FeeTokens []common.Address
}

type DestChainConfig = committee_verifier.GetDestChainConfig

type SetSignatureConfigArgs struct {
	Signers   []common.Address
	Threshold uint8
}

var Deploy = contract.NewDeploy(
	"committee-verifier:deploy",
	semver.MustParse("1.7.0"),
	"Deploys the CommitteeVerifier contract",
	ContractType,
	committee_verifier.CommitteeVerifierABI,
	func(ConstructorArgs) error { return nil },
	contract.VMDeployers[ConstructorArgs]{
		DeployEVM: func(opts *bind.TransactOpts, backend bind.ContractBackend, args ConstructorArgs) (common.Address, *types.Transaction, error) {
			address, tx, _, err := committee_verifier.DeployCommitteeVerifier(opts, backend, args.DynamicConfig, args.StorageLocation)
			return address, tx, err
		},
		// DeployZksyncVM: func(opts *accounts.TransactOpts, client *clients.Client, wallet *accounts.Wallet, backend bind.ContractBackend, args ConstructorArgs) (common.Address, error)
	},
)

var DeployProxy = contract.NewDeploy(
	"committee-verifier-proxy:deploy",
	semver.MustParse("1.7.0"),
	"Deploys the CommitteeVerifierProxy contract",
	ProxyType,
	verifier_proxy.VerifierProxyABI,
	func(common.Address) error { return nil },
	contract.VMDeployers[common.Address]{
		DeployEVM: func(opts *bind.TransactOpts, backend bind.ContractBackend, rampAddress common.Address) (common.Address, *types.Transaction, error) {
			address, tx, _, err := verifier_proxy.DeployVerifierProxy(opts, backend, rampAddress)
			return address, tx, err
		},
		// DeployZksyncVM: func(opts *accounts.TransactOpts, client *clients.Client, wallet *accounts.Wallet, backend bind.ContractBackend, rampAddress common.Address) (common.Address, error)
	},
)

var SetDynamicConfig = contract.NewWrite(
	"committee-verifier:set-dynamic-config",
	semver.MustParse("1.7.0"),
	"Sets the dynamic configuration on the CommitteeVerifier",
	ContractType,
	committee_verifier.CommitteeVerifierABI,
	committee_verifier.NewCommitteeVerifier,
	contract.OnlyOwner,
	func(SetDynamicConfigArgs) error { return nil },
	func(committeeVerifier *committee_verifier.CommitteeVerifier, opts *bind.TransactOpts, args SetDynamicConfigArgs) (*types.Transaction, error) {
		return committeeVerifier.SetDynamicConfig(opts, args.DynamicConfig)
	},
)

var ApplyDestChainConfigUpdates = contract.NewWrite(
	"committee-verifier:apply-dest-chain-config-updates",
	semver.MustParse("1.7.0"),
	"Applies updates to destination chain configurations on the CommitteeVerifier",
	ContractType,
	committee_verifier.CommitteeVerifierABI,
	committee_verifier.NewCommitteeVerifier,
	contract.OnlyOwner,
	func([]DestChainConfigArgs) error { return nil },
	func(committeeVerifier *committee_verifier.CommitteeVerifier, opts *bind.TransactOpts, args []DestChainConfigArgs) (*types.Transaction, error) {
		return committeeVerifier.ApplyDestChainConfigUpdates(opts, args)
	},
)

var ApplyAllowlistUpdates = contract.NewWrite(
	"committee-verifier:apply-allowlist-updates",
	semver.MustParse("1.7.0"),
	"Applies updates to the allowlist (those authorized to send messages) on the CommitteeVerifier",
	ContractType,
	committee_verifier.CommitteeVerifierABI,
	committee_verifier.NewCommitteeVerifier,
	contract.OnlyOwner,
	func([]AllowlistConfigArgs) error { return nil },
	func(committeeVerifier *committee_verifier.CommitteeVerifier, opts *bind.TransactOpts, args []AllowlistConfigArgs) (*types.Transaction, error) {
		return committeeVerifier.ApplyAllowlistUpdates(opts, args)
	},
)

var WithdrawFeeTokens = contract.NewWrite(
	"committee-verifier:withdraw-fee-tokens",
	semver.MustParse("1.7.0"),
	"Withdraws fee tokens from the CommitteeVerifier",
	ContractType,
	committee_verifier.CommitteeVerifierABI,
	committee_verifier.NewCommitteeVerifier,
	contract.OnlyOwner,
	func(WithdrawFeeTokensArgs) error { return nil },
	func(committeeVerifier *committee_verifier.CommitteeVerifier, opts *bind.TransactOpts, args WithdrawFeeTokensArgs) (*types.Transaction, error) {
		return committeeVerifier.WithdrawFeeTokens(opts, args.FeeTokens)
	},
)

var GetDestChainConfig = contract.NewRead(
	"committee-verifier:get-dest-chain-config",
	semver.MustParse("1.7.0"),
	"Gets the destination chain configuration for a given destination chain selector",
	ContractType,
	committee_verifier.NewCommitteeVerifier,
	func(committeeVerifier *committee_verifier.CommitteeVerifier, opts *bind.CallOpts, destChainSelector uint64) (DestChainConfig, error) {
		return committeeVerifier.GetDestChainConfig(opts, destChainSelector)
	},
)

var SetSignatureConfigs = contract.NewWrite(
	"committee-verifier:set-signature-config",
	semver.MustParse("1.7.0"),
	"Sets the signature configuration on the CommitteeVerifier",
	ContractType,
	committee_verifier.CommitteeVerifierABI,
	committee_verifier.NewCommitteeVerifier,
	contract.OnlyOwner,
	func(SetSignatureConfigArgs) error { return nil },
	func(committeeVerifier *committee_verifier.CommitteeVerifier, opts *bind.TransactOpts, args SetSignatureConfigArgs) (*types.Transaction, error) {
		return committeeVerifier.SetSignatureConfig(opts, args.Signers, args.Threshold)
	},
)
