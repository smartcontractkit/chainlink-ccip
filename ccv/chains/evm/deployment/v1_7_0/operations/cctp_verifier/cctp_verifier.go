package cctp_verifier

import (
	"errors"
	"fmt"

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
	Version:          Version,
	Description:      "Deploys the CCTPVerifier contract",
	ContractMetadata: cctp_verifier.CCTPVerifierMetaData,
	BytecodeByTypeAndVersion: map[string]contract.Bytecode{
		cldf_deployment.NewTypeAndVersion(ContractType, *Version).String(): {
			EVM: common.FromHex(cctp_verifier.CCTPVerifierBin),
		},
	},
	Validate: func(ConstructorArgs) error { return nil },
})

var ApplyRemoteChainConfigUpdates = contract.NewWrite(contract.WriteParams[[]RemoteChainConfigArgs, *cctp_verifier.CCTPVerifier]{
	Name:            "cctp-verifier:apply-remote-chain-config-updates",
	Version:         Version,
	Description:     "Applies updates to remote chain configurations on the CCTPVerifier",
	ContractType:    ContractType,
	ContractABI:     cctp_verifier.CCTPVerifierABI,
	NewContract:     cctp_verifier.NewCCTPVerifier,
	IsAllowedCaller: contract.OnlyOwner[*cctp_verifier.CCTPVerifier, []RemoteChainConfigArgs],
	Validate: func(cctpVerifier *cctp_verifier.CCTPVerifier, backend bind.ContractBackend, opts *bind.CallOpts, args []RemoteChainConfigArgs) error {
		for _, cfg := range args {
			if cfg.RemoteChainSelector == 0 {
				return errors.New("remote chain selector cannot be 0")
			}
			if cfg.GasForVerification == 0 {
				return errors.New("gas for verification cannot be 0")
			}
		}

		return nil
	},
	IsNoop: func(cctpVerifier *cctp_verifier.CCTPVerifier, opts *bind.CallOpts, args []RemoteChainConfigArgs) (bool, error) {
		for _, cfg := range args {
			remoteChainConfig, err := cctpVerifier.GetRemoteChainConfig(opts, cfg.RemoteChainSelector)
			if err != nil {
				return false, fmt.Errorf("failed to get remote chain config for remote chain selector %d: %w", cfg.RemoteChainSelector, err)
			}
			if remoteChainConfig.Router == (common.Address{}) {
				return false, nil
			}
			feeConfig, err := cctpVerifier.GetFee(opts, cfg.RemoteChainSelector, cctp_verifier.ClientEVM2AnyMessage{}, []byte{}, 0)
			if err != nil {
				return false, fmt.Errorf("failed to get fee for remote chain selector %d: %w", cfg.RemoteChainSelector, err)
			}
			if feeConfig.GasForVerification != cfg.GasForVerification ||
				feeConfig.PayloadSizeBytes != cfg.PayloadSizeBytes ||
				feeConfig.FeeUSDCents != cfg.FeeUSDCents ||
				remoteChainConfig.Router != cfg.Router ||
				remoteChainConfig.AllowlistEnabled != cfg.AllowlistEnabled {
				return false, nil
			}
		}

		return true, nil
	},
	CallContract: func(cctpVerifier *cctp_verifier.CCTPVerifier, opts *bind.TransactOpts, args []RemoteChainConfigArgs) (*types.Transaction, error) {
		return cctpVerifier.ApplyRemoteChainConfigUpdates(opts, args)
	},
})

var SetDomains = contract.NewWrite(contract.WriteParams[[]SetDomainArgs, *cctp_verifier.CCTPVerifier]{
	Name:            "cctp-verifier:set-domains",
	Version:         Version,
	Description:     "Sets domain configurations on the CCTPVerifier",
	ContractType:    ContractType,
	ContractABI:     cctp_verifier.CCTPVerifierABI,
	NewContract:     cctp_verifier.NewCCTPVerifier,
	IsAllowedCaller: contract.OnlyOwner[*cctp_verifier.CCTPVerifier, []SetDomainArgs],
	Validate: func(cctpVerifier *cctp_verifier.CCTPVerifier, backend bind.ContractBackend, opts *bind.CallOpts, domains []SetDomainArgs) error {
		for _, domain := range domains {
			if domain.ChainSelector == 0 {
				return errors.New("chain selector cannot be 0")
			}
			if domain.AllowedCallerOnDest == ([32]byte{}) {
				return errors.New("allowed caller on destination cannot be the zero address")
			}
			if domain.AllowedCallerOnSource == ([32]byte{}) {
				return errors.New("allowed caller on source cannot be the zero address")
			}
		}

		return nil
	},
	IsNoop: func(cctpVerifier *cctp_verifier.CCTPVerifier, opts *bind.CallOpts, domains []SetDomainArgs) (bool, error) {
		for _, domainArgs := range domains {
			domain, err := cctpVerifier.GetDomain(opts, domainArgs.ChainSelector)
			if err != nil {
				return false, fmt.Errorf("failed to get domain for chain selector %d: %w", domainArgs.ChainSelector, err)
			}
			if domain.AllowedCallerOnDest != domainArgs.AllowedCallerOnDest ||
				domain.AllowedCallerOnSource != domainArgs.AllowedCallerOnSource ||
				domain.MintRecipientOnDest != domainArgs.MintRecipientOnDest ||
				domain.DomainIdentifier != domainArgs.DomainIdentifier ||
				domain.Enabled != domainArgs.Enabled {
				return false, nil
			}
		}

		return true, nil
	},
	CallContract: func(cctpVerifier *cctp_verifier.CCTPVerifier, opts *bind.TransactOpts, args []SetDomainArgs) (*types.Transaction, error) {
		return cctpVerifier.SetDomains(opts, args)
	},
})

var GetVersionTag = contract.NewRead(contract.ReadParams[any, [4]byte, *cctp_verifier.CCTPVerifier]{
	Name:         "cctp-verifier:get-version-tag",
	Version:      Version,
	Description:  "Gets the version tag of the CCTPVerifier contract",
	ContractType: ContractType,
	NewContract:  cctp_verifier.NewCCTPVerifier,
	CallContract: func(cctpVerifier *cctp_verifier.CCTPVerifier, opts *bind.CallOpts, args any) ([4]byte, error) {
		return cctpVerifier.VersionTag(opts)
	},
})
