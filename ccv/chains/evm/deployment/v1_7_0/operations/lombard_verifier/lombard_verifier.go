package lombard_verifier

import (
	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/gobindings/generated/latest/lombard_verifier"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"

	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

var ContractType cldf_deployment.ContractType = "LombardVerifier"

var ResolverType cldf_deployment.ContractType = "LombardVerifierResolver"

var Version *semver.Version = semver.MustParse("1.7.0")

type DynamicConfig = lombard_verifier.LombardVerifierDynamicConfig

type RemoteChainConfigArgs = lombard_verifier.BaseVerifierRemoteChainConfigArgs

type SupportedTokensArgs = lombard_verifier.LombardVerifierSupportedTokenArgs

type RemotePathArgs struct {
	RemoteChainSelector uint64
	AllowedCaller       [32]byte
	LChainId            [32]byte
}

type ConstructorArgs struct {
	DynamicConfig    DynamicConfig
	Bridge           common.Address
	StorageLocations []string
	RMN              common.Address
}

type SupportedTokenArgs struct {
	TokensToRemove []common.Address
	TokensToSet    []SupportedTokensArgs
}

var Deploy = contract.NewDeploy(contract.DeployParams[ConstructorArgs]{
	Name:             "lombard-verifier:deploy",
	Version:          Version,
	Description:      "Deploys the LombardVerifier contract",
	ContractMetadata: lombard_verifier.LombardVerifierMetaData,
	BytecodeByTypeAndVersion: map[string]contract.Bytecode{
		cldf_deployment.NewTypeAndVersion(ContractType, *Version).String(): {
			EVM: common.FromHex(lombard_verifier.LombardVerifierBin),
		},
	},
	Validate: func(ConstructorArgs) error { return nil },
})

var ApplyRemoteChainConfigUpdates = contract.NewWrite(contract.WriteParams[[]RemoteChainConfigArgs, *lombard_verifier.LombardVerifier]{
	Name:            "lombard-verifier:apply-remote-chain-config-updates",
	Version:         Version,
	Description:     "Applies updates to remote chain configurations on the LombardVerifier",
	ContractType:    ContractType,
	ContractABI:     lombard_verifier.LombardVerifierABI,
	NewContract:     lombard_verifier.NewLombardVerifier,
	IsAllowedCaller: contract.OnlyOwner[*lombard_verifier.LombardVerifier, []RemoteChainConfigArgs],
	Validate:        func([]RemoteChainConfigArgs) error { return nil },
	CallContract: func(lombardVerifier *lombard_verifier.LombardVerifier, opts *bind.TransactOpts, args []RemoteChainConfigArgs) (*types.Transaction, error) {
		return lombardVerifier.ApplyRemoteChainConfigUpdates(opts, args)
	},
})

var SetRemotePath = contract.NewWrite(contract.WriteParams[RemotePathArgs, *lombard_verifier.LombardVerifier]{
	Name:            "lombard-verifier:set-remote-path",
	Version:         Version,
	Description:     "Sets remote path on the LombardVerifier",
	ContractType:    ContractType,
	ContractABI:     lombard_verifier.LombardVerifierABI,
	NewContract:     lombard_verifier.NewLombardVerifier,
	IsAllowedCaller: contract.OnlyOwner[*lombard_verifier.LombardVerifier, RemotePathArgs],
	Validate:        func(RemotePathArgs) error { return nil },
	CallContract: func(lombardVerifier *lombard_verifier.LombardVerifier, opts *bind.TransactOpts, args RemotePathArgs) (*types.Transaction, error) {
		return lombardVerifier.SetPath(opts, args.RemoteChainSelector, args.LChainId, args.AllowedCaller)
	},
})

var UpdateSupportedTokens = contract.NewWrite(contract.WriteParams[SupportedTokenArgs, *lombard_verifier.LombardVerifier]{
	Name:            "lombard-verifier:update-supported-tokens",
	Version:         Version,
	Description:     "Updates supported tokens on the LombardVerifier",
	ContractType:    ContractType,
	ContractABI:     lombard_verifier.LombardVerifierABI,
	NewContract:     lombard_verifier.NewLombardVerifier,
	IsAllowedCaller: contract.OnlyOwner[*lombard_verifier.LombardVerifier, SupportedTokenArgs],
	Validate:        func(SupportedTokenArgs) error { return nil },
	CallContract: func(lombardVerifier *lombard_verifier.LombardVerifier, opts *bind.TransactOpts, args SupportedTokenArgs) (*types.Transaction, error) {
		return lombardVerifier.UpdateSupportedTokens(opts, args.TokensToRemove, args.TokensToSet)
	},
})
