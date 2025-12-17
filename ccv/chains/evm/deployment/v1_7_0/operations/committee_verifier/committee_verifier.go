package committee_verifier

import (
	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/gobindings/generated/latest/committee_verifier"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/gobindings/generated/latest/versioned_verifier_resolver"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

var ContractType cldf_deployment.ContractType = "CommitteeVerifier"

var ResolverType cldf_deployment.ContractType = "CommitteeVerifierResolver"

type DynamicConfig = committee_verifier.CommitteeVerifierDynamicConfig

type ConstructorArgs struct {
	DynamicConfig    DynamicConfig
	StorageLocations []string
	RMN              common.Address
}

type ResolverConstructorArgs struct{}

type SetDynamicConfigArgs struct {
	DynamicConfig DynamicConfig
}

type RemoteChainConfigArgs = committee_verifier.BaseVerifierRemoteChainConfigArgs

type AllowlistConfigArgs = committee_verifier.BaseVerifierAllowlistConfigArgs

type WithdrawFeeTokensArgs struct {
	FeeTokens []common.Address
}

type RemoteChainConfig = committee_verifier.GetRemoteChainConfig

type SignatureConfig = committee_verifier.SignatureQuorumValidatorSignatureConfig

type SignatureConfigArgs struct {
	SourceChainSelectorsToRemove []uint64
	SignatureConfigUpdates       []SignatureConfig
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

var DeployResolver = contract.NewDeploy(contract.DeployParams[ResolverConstructorArgs]{
	Name:             "committee-verifier-resolver:deploy",
	Version:          semver.MustParse("1.7.0"),
	Description:      "Deploys the CommitteeVerifierResolver contract",
	ContractMetadata: versioned_verifier_resolver.VersionedVerifierResolverMetaData,
	BytecodeByTypeAndVersion: map[string]contract.Bytecode{
		cldf_deployment.NewTypeAndVersion(ResolverType, *semver.MustParse("1.7.0")).String(): {
			EVM: common.FromHex(versioned_verifier_resolver.VersionedVerifierResolverBin),
		},
	},
	Validate: func(ResolverConstructorArgs) error { return nil },
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

var ApplyRemoteChainConfigUpdates = contract.NewWrite(contract.WriteParams[[]RemoteChainConfigArgs, *committee_verifier.CommitteeVerifier]{
	Name:            "committee-verifier:apply-remote-chain-config-updates",
	Version:         semver.MustParse("1.7.0"),
	Description:     "Applies updates to remote chain configurations on the CommitteeVerifier",
	ContractType:    ContractType,
	ContractABI:     committee_verifier.CommitteeVerifierABI,
	NewContract:     committee_verifier.NewCommitteeVerifier,
	IsAllowedCaller: contract.OnlyOwner[*committee_verifier.CommitteeVerifier, []RemoteChainConfigArgs],
	Validate:        func([]RemoteChainConfigArgs) error { return nil },
	CallContract: func(committeeVerifier *committee_verifier.CommitteeVerifier, opts *bind.TransactOpts, args []RemoteChainConfigArgs) (*types.Transaction, error) {
		return committeeVerifier.ApplyRemoteChainConfigUpdates(opts, args)
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

var ApplySignatureConfigs = contract.NewWrite(contract.WriteParams[SignatureConfigArgs, *committee_verifier.CommitteeVerifier]{
	Name:            "committee-verifier:apply-signature-configs",
	Version:         semver.MustParse("1.7.0"),
	Description:     "Applies the signature configurations on the CommitteeVerifier",
	ContractType:    ContractType,
	ContractABI:     committee_verifier.CommitteeVerifierABI,
	NewContract:     committee_verifier.NewCommitteeVerifier,
	IsAllowedCaller: contract.OnlyOwner[*committee_verifier.CommitteeVerifier, SignatureConfigArgs],
	Validate:        func(SignatureConfigArgs) error { return nil },
	CallContract: func(committeeVerifier *committee_verifier.CommitteeVerifier, opts *bind.TransactOpts, args SignatureConfigArgs) (*types.Transaction, error) {
		return committeeVerifier.ApplySignatureConfigs(opts, args.SourceChainSelectorsToRemove, args.SignatureConfigUpdates)
	},
})

var GetRemoteChainConfig = contract.NewRead(contract.ReadParams[uint64, RemoteChainConfig, *committee_verifier.CommitteeVerifier]{
	Name:         "committee-verifier:get-remote-chain-config",
	Version:      semver.MustParse("1.7.0"),
	Description:  "Gets the remote chain configuration for a given remote chain selector",
	ContractType: ContractType,
	NewContract:  committee_verifier.NewCommitteeVerifier,
	CallContract: func(committeeVerifier *committee_verifier.CommitteeVerifier, opts *bind.CallOpts, args uint64) (RemoteChainConfig, error) {
		return committeeVerifier.GetRemoteChainConfig(opts, args)
	},
})

var GetVersionTag = contract.NewRead(contract.ReadParams[any, [4]byte, *committee_verifier.CommitteeVerifier]{
	Name:         "committee-verifier:get-version-tag",
	Version:      semver.MustParse("1.7.0"),
	Description:  "Gets the version tag of the CommitteeVerifier contract",
	ContractType: ContractType,
	NewContract:  committee_verifier.NewCommitteeVerifier,
	CallContract: func(committeeVerifier *committee_verifier.CommitteeVerifier, opts *bind.CallOpts, args any) ([4]byte, error) {
		return committeeVerifier.VersionTag(opts)
	},
})
