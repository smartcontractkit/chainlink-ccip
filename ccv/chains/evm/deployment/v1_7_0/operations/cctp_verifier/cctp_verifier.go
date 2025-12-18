package cctp_verifier

import (
	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/gobindings/generated/latest/cctp_verifier"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

var ContractType cldf_deployment.ContractType = "CCTPVerifier"

var ResolverType cldf_deployment.ContractType = "CCTPVerifierResolver"

var Version *semver.Version = semver.MustParse("1.7.0")

type ConstructorArgs struct {
	TokenMessenger          common.Address
	MessageTransmitterProxy common.Address
	USDCToken               common.Address
	StorageLocations        []string
	DynamicConfig           DynamicConfig
	RMN                     common.Address
}

type ResolverConstructorArgs struct{}

type DynamicConfig = cctp_verifier.CCTPVerifierDynamicConfig

type RemoteChainConfigArgs = cctp_verifier.BaseVerifierRemoteChainConfigArgs

type SetDomainArgs = cctp_verifier.CCTPVerifierSetDomainArgs

var Deploy = contract.NewDeploy(contract.DeployParams[ConstructorArgs]{
	Name:             "cctp-verifier:deploy",
	Version:          semver.MustParse("1.7.0"),
	Description:      "Deploys the CCTPVerifier contract",
	ContractMetadata: cctp_verifier.CCTPVerifierMetaData,
	BytecodeByTypeAndVersion: map[string]contract.Bytecode{
		cldf_deployment.NewTypeAndVersion(ContractType, *semver.MustParse("1.7.0")).String(): {
			EVM: common.FromHex(cctp_verifier.CCTPVerifierBin),
		},
	},
	Validate: func(ConstructorArgs) error { return nil },
})

var ApplyRemoteChainConfigUpdates = contract.NewWrite(contract.WriteParams[[]RemoteChainConfigArgs, *cctp_verifier.CCTPVerifier]{
	Name:            "cctp-verifier:apply-remote-chain-config-updates",
	Version:         semver.MustParse("1.7.0"),
	Description:     "Applies updates to remote chain configurations on the CCTPVerifier",
	ContractType:    ContractType,
	ContractABI:     cctp_verifier.CCTPVerifierABI,
	NewContract:     cctp_verifier.NewCCTPVerifier,
	IsAllowedCaller: contract.OnlyOwner[*cctp_verifier.CCTPVerifier, []RemoteChainConfigArgs],
	Validate:        func([]RemoteChainConfigArgs) error { return nil },
	CallContract: func(cctpVerifier *cctp_verifier.CCTPVerifier, opts *bind.TransactOpts, args []RemoteChainConfigArgs) (*types.Transaction, error) {
		return cctpVerifier.ApplyRemoteChainConfigUpdates(opts, args)
	},
})

var SetDomains = contract.NewWrite(contract.WriteParams[[]SetDomainArgs, *cctp_verifier.CCTPVerifier]{
	Name:            "cctp-verifier:set-domains",
	Version:         semver.MustParse("1.7.0"),
	Description:     "Sets domain configurations on the CCTPVerifier",
	ContractType:    ContractType,
	ContractABI:     cctp_verifier.CCTPVerifierABI,
	NewContract:     cctp_verifier.NewCCTPVerifier,
	IsAllowedCaller: contract.OnlyOwner[*cctp_verifier.CCTPVerifier, []SetDomainArgs],
	Validate:        func([]SetDomainArgs) error { return nil },
	CallContract: func(cctpVerifier *cctp_verifier.CCTPVerifier, opts *bind.TransactOpts, args []SetDomainArgs) (*types.Transaction, error) {
		return cctpVerifier.SetDomains(opts, args)
	},
})

var GetVersionTag = contract.NewRead(contract.ReadParams[any, [4]byte, *cctp_verifier.CCTPVerifier]{
	Name:         "cctp-verifier:get-version-tag",
	Version:      semver.MustParse("1.7.0"),
	Description:  "Gets the version tag of the CCTPVerifier contract",
	ContractType: ContractType,
	NewContract:  cctp_verifier.NewCCTPVerifier,
	CallContract: func(cctpVerifier *cctp_verifier.CCTPVerifier, opts *bind.CallOpts, args any) ([4]byte, error) {
		return cctpVerifier.VersionTag(opts)
	},
})
