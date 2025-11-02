package committee_verifier

import (
	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/gobindings/generated/latest/committee_verifier"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/gobindings/generated/latest/proxy"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

var ContractType cldf_deployment.ContractType = "CommitteeVerifier"

var ProxyType cldf_deployment.ContractType = "CommitteeVerifierProxy"

type DynamicConfig = committee_verifier.CommitteeVerifierDynamicConfig

type ConstructorArgs struct {
	DynamicConfig   DynamicConfig
	StorageLocation string
}

type ProxyConstructorArgs struct {
	RampAddress common.Address
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

var Deploy = contract.NewDeploy(contract.DeployParams[ConstructorArgs]{
	Name:             "committee-verifier:deploy",
	Version:          semver.MustParse("1.7.0"),
	Description:      "Deploys the CommitteeVerifier contract",
	ContractMetadata: committee_verifier.CommitteeVerifierMetaData,
	BytecodeByTypeAndVersion: map[string]contract.Bytecode{
		cldf_deployment.NewTypeAndVersion(ContractType, *semver.MustParse("1.7.0")).String(): {
			EVM: common.FromHex(committee_verifier.CommitteeVerifierBin),
		},
	},
	Validate: func(ConstructorArgs) error { return nil },
})

var DeployProxy = contract.NewDeploy(contract.DeployParams[ProxyConstructorArgs]{
	Name:             "committee-verifier-proxy:deploy",
	Version:          semver.MustParse("1.7.0"),
	Description:      "Deploys the CommitteeVerifierProxy contract",
	ContractMetadata: proxy.ProxyMetaData,
	BytecodeByTypeAndVersion: map[string]contract.Bytecode{
		cldf_deployment.NewTypeAndVersion(ProxyType, *semver.MustParse("1.7.0")).String(): {
			EVM: common.FromHex(proxy.ProxyBin),
		},
	},
	Validate: func(ProxyConstructorArgs) error { return nil },
})

var SetDynamicConfig = contract.NewWrite(contract.WriteParams[SetDynamicConfigArgs, *committee_verifier.CommitteeVerifier]{
	Name:            "committee-verifier:set-dynamic-config",
	Version:         semver.MustParse("1.7.0"),
	Description:     "Sets the dynamic configuration on the CommitteeVerifier",
	ContractType:    ContractType,
	ContractABI:     committee_verifier.CommitteeVerifierABI,
	NewContract:     committee_verifier.NewCommitteeVerifier,
	IsAllowedCaller: contract.OnlyOwner[*committee_verifier.CommitteeVerifier, SetDynamicConfigArgs],
	Validate:        func(SetDynamicConfigArgs) error { return nil },
	CallContract: func(committeeVerifier *committee_verifier.CommitteeVerifier, opts *bind.TransactOpts, args SetDynamicConfigArgs) (*types.Transaction, error) {
		return committeeVerifier.SetDynamicConfig(opts, args.DynamicConfig)
	},
})

var ApplyDestChainConfigUpdates = contract.NewWrite(contract.WriteParams[[]DestChainConfigArgs, *committee_verifier.CommitteeVerifier]{
	Name:            "committee-verifier:apply-dest-chain-config-updates",
	Version:         semver.MustParse("1.7.0"),
	Description:     "Applies updates to destination chain configurations on the CommitteeVerifier",
	ContractType:    ContractType,
	ContractABI:     committee_verifier.CommitteeVerifierABI,
	NewContract:     committee_verifier.NewCommitteeVerifier,
	IsAllowedCaller: contract.OnlyOwner[*committee_verifier.CommitteeVerifier, []DestChainConfigArgs],
	Validate:        func([]DestChainConfigArgs) error { return nil },
	CallContract: func(committeeVerifier *committee_verifier.CommitteeVerifier, opts *bind.TransactOpts, args []DestChainConfigArgs) (*types.Transaction, error) {
		return committeeVerifier.ApplyDestChainConfigUpdates(opts, args)
	},
})

var ApplyAllowlistUpdates = contract.NewWrite(contract.WriteParams[[]AllowlistConfigArgs, *committee_verifier.CommitteeVerifier]{
	Name:            "committee-verifier:apply-allowlist-updates",
	Version:         semver.MustParse("1.7.0"),
	Description:     "Applies updates to the allowlist (those authorized to send messages) on the CommitteeVerifier",
	ContractType:    ContractType,
	ContractABI:     committee_verifier.CommitteeVerifierABI,
	NewContract:     committee_verifier.NewCommitteeVerifier,
	IsAllowedCaller: contract.OnlyOwner[*committee_verifier.CommitteeVerifier, []AllowlistConfigArgs],
	Validate:        func([]AllowlistConfigArgs) error { return nil },
	CallContract: func(committeeVerifier *committee_verifier.CommitteeVerifier, opts *bind.TransactOpts, args []AllowlistConfigArgs) (*types.Transaction, error) {
		return committeeVerifier.ApplyAllowlistUpdates(opts, args)
	},
})

var WithdrawFeeTokens = contract.NewWrite(contract.WriteParams[WithdrawFeeTokensArgs, *committee_verifier.CommitteeVerifier]{
	Name:            "committee-verifier:withdraw-fee-tokens",
	Version:         semver.MustParse("1.7.0"),
	Description:     "Withdraws fee tokens from the CommitteeVerifier",
	ContractType:    ContractType,
	ContractABI:     committee_verifier.CommitteeVerifierABI,
	NewContract:     committee_verifier.NewCommitteeVerifier,
	IsAllowedCaller: contract.OnlyOwner[*committee_verifier.CommitteeVerifier, WithdrawFeeTokensArgs],
	Validate:        func(WithdrawFeeTokensArgs) error { return nil },
	CallContract: func(committeeVerifier *committee_verifier.CommitteeVerifier, opts *bind.TransactOpts, args WithdrawFeeTokensArgs) (*types.Transaction, error) {
		return committeeVerifier.WithdrawFeeTokens(opts, args.FeeTokens)
	},
})

var SetSignatureConfigs = contract.NewWrite(contract.WriteParams[SetSignatureConfigArgs, *committee_verifier.CommitteeVerifier]{
	Name:            "committee-verifier:set-signature-config",
	Version:         semver.MustParse("1.7.0"),
	Description:     "Sets the signature configuration on the CommitteeVerifier",
	ContractType:    ContractType,
	ContractABI:     committee_verifier.CommitteeVerifierABI,
	NewContract:     committee_verifier.NewCommitteeVerifier,
	IsAllowedCaller: contract.OnlyOwner[*committee_verifier.CommitteeVerifier, SetSignatureConfigArgs],
	Validate:        func(SetSignatureConfigArgs) error { return nil },
	CallContract: func(committeeVerifier *committee_verifier.CommitteeVerifier, opts *bind.TransactOpts, args SetSignatureConfigArgs) (*types.Transaction, error) {
		return committeeVerifier.SetSignatureConfig(opts, args.Signers, args.Threshold)
	},
})

var GetDestChainConfig = contract.NewRead(contract.ReadParams[uint64, DestChainConfig, *committee_verifier.CommitteeVerifier]{
	Name:         "committee-verifier:get-dest-chain-config",
	Version:      semver.MustParse("1.7.0"),
	Description:  "Gets the destination chain configuration for a given destination chain selector",
	ContractType: ContractType,
	NewContract:  committee_verifier.NewCommitteeVerifier,
	CallContract: func(committeeVerifier *committee_verifier.CommitteeVerifier, opts *bind.CallOpts, args uint64) (DestChainConfig, error) {
		return committeeVerifier.GetDestChainConfig(opts, args)
	},
})
