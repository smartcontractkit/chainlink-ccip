package committee_verifier

import (
	"bytes"
	"errors"
	"fmt"
	"slices"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/gobindings/generated/latest/committee_verifier"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

var ContractType cldf_deployment.ContractType = "CommitteeVerifier"

var ResolverType cldf_deployment.ContractType = "CommitteeVerifierResolver"

var Version = semver.MustParse("1.7.0")

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
	Version:          Version,
	Description:      "Deploys the CommitteeVerifier contract",
	ContractMetadata: committee_verifier.CommitteeVerifierMetaData,
	BytecodeByTypeAndVersion: map[string]contract.Bytecode{
		cldf_deployment.NewTypeAndVersion(ContractType, *Version).String(): {
			EVM: common.FromHex(committee_verifier.CommitteeVerifierBin),
		},
	},
	Validate: func(ConstructorArgs) error { return nil },
})

var SetDynamicConfig = contract.NewWrite(contract.WriteParams[SetDynamicConfigArgs, *committee_verifier.CommitteeVerifier]{
	Name:            "committee-verifier:set-dynamic-config",
	Version:         Version,
	Description:     "Sets the dynamic configuration on the CommitteeVerifier",
	ContractType:    ContractType,
	ContractABI:     committee_verifier.CommitteeVerifierABI,
	NewContract:     committee_verifier.NewCommitteeVerifier,
	IsAllowedCaller: contract.OnlyOwner[*committee_verifier.CommitteeVerifier, SetDynamicConfigArgs],
	Validate: func(committeeVerifier *committee_verifier.CommitteeVerifier, backend bind.ContractBackend, opts *bind.CallOpts, args SetDynamicConfigArgs) error {
		if args.DynamicConfig.FeeAggregator == (common.Address{}) {
			return errors.New("fee aggregator cannot be the zero address")
		}

		return nil
	},
	IsNoop: func(committeeVerifier *committee_verifier.CommitteeVerifier, opts *bind.CallOpts, args SetDynamicConfigArgs) (bool, error) {
		currentDynamicConfig, err := committeeVerifier.GetDynamicConfig(opts)
		if err != nil {
			return false, fmt.Errorf("failed to get dynamic configuration: %w", err)
		}

		return currentDynamicConfig == args.DynamicConfig, nil
	},
	CallContract: func(committeeVerifier *committee_verifier.CommitteeVerifier, opts *bind.TransactOpts, args SetDynamicConfigArgs) (*types.Transaction, error) {
		return committeeVerifier.SetDynamicConfig(opts, args.DynamicConfig)
	},
})

var ApplyRemoteChainConfigUpdates = contract.NewWrite(contract.WriteParams[[]RemoteChainConfigArgs, *committee_verifier.CommitteeVerifier]{
	Name:            "committee-verifier:apply-remote-chain-config-updates",
	Version:         Version,
	Description:     "Applies updates to remote chain configurations on the CommitteeVerifier",
	ContractType:    ContractType,
	ContractABI:     committee_verifier.CommitteeVerifierABI,
	NewContract:     committee_verifier.NewCommitteeVerifier,
	IsAllowedCaller: contract.OnlyOwner[*committee_verifier.CommitteeVerifier, []RemoteChainConfigArgs],
	Validate: func(committeeVerifier *committee_verifier.CommitteeVerifier, backend bind.ContractBackend, opts *bind.CallOpts, args []RemoteChainConfigArgs) error {
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
	IsNoop: func(committeeVerifier *committee_verifier.CommitteeVerifier, opts *bind.CallOpts, args []RemoteChainConfigArgs) (bool, error) {
		for _, cfg := range args {
			remoteChainConfig, err := committeeVerifier.GetRemoteChainConfig(opts, cfg.RemoteChainSelector)
			if err != nil {
				return false, fmt.Errorf("failed to get remote chain config for remote chain selector %d: %w", cfg.RemoteChainSelector, err)
			}
			if remoteChainConfig.Router == (common.Address{}) {
				return false, nil
			}
			feeConfig, err := committeeVerifier.GetFee(opts, cfg.RemoteChainSelector, committee_verifier.ClientEVM2AnyMessage{}, []byte{}, 0)
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
	CallContract: func(committeeVerifier *committee_verifier.CommitteeVerifier, opts *bind.TransactOpts, args []RemoteChainConfigArgs) (*types.Transaction, error) {
		return committeeVerifier.ApplyRemoteChainConfigUpdates(opts, args)
	},
})

var WithdrawFeeTokens = contract.NewWrite(contract.WriteParams[WithdrawFeeTokensArgs, *committee_verifier.CommitteeVerifier]{
	Name:         "committee-verifier:withdraw-fee-tokens",
	Version:      Version,
	Description:  "Withdraws fee tokens from the CommitteeVerifier",
	ContractType: ContractType,
	ContractABI:  committee_verifier.CommitteeVerifierABI,
	NewContract:  committee_verifier.NewCommitteeVerifier,
	IsAllowedCaller: func(committeeVerifier *committee_verifier.CommitteeVerifier, opts *bind.CallOpts, caller common.Address, args WithdrawFeeTokensArgs) (bool, error) {
		return true, nil
	},
	Validate: func(committeeVerifier *committee_verifier.CommitteeVerifier, backend bind.ContractBackend, opts *bind.CallOpts, args WithdrawFeeTokensArgs) error {
		return nil
	},
	IsNoop: func(committeeVerifier *committee_verifier.CommitteeVerifier, opts *bind.CallOpts, args WithdrawFeeTokensArgs) (bool, error) {
		return false, nil
	},
	CallContract: func(committeeVerifier *committee_verifier.CommitteeVerifier, opts *bind.TransactOpts, args WithdrawFeeTokensArgs) (*types.Transaction, error) {
		return committeeVerifier.WithdrawFeeTokens(opts, args.FeeTokens)
	},
})

var ApplySignatureConfigs = contract.NewWrite(contract.WriteParams[SignatureConfigArgs, *committee_verifier.CommitteeVerifier]{
	Name:            "committee-verifier:apply-signature-configs",
	Version:         Version,
	Description:     "Applies the signature configurations on the CommitteeVerifier",
	ContractType:    ContractType,
	ContractABI:     committee_verifier.CommitteeVerifierABI,
	NewContract:     committee_verifier.NewCommitteeVerifier,
	IsAllowedCaller: contract.OnlyOwner[*committee_verifier.CommitteeVerifier, SignatureConfigArgs],
	Validate: func(committeeVerifier *committee_verifier.CommitteeVerifier, backend bind.ContractBackend, opts *bind.CallOpts, args SignatureConfigArgs) error {
		for _, cfg := range args.SignatureConfigUpdates {
			if cfg.Threshold == 0 {
				return errors.New("threshold cannot be 0")
			}
			if cfg.Threshold > uint8(len(cfg.Signers)) {
				return errors.New("threshold cannot be greater than the number of signers")
			}
		}

		return nil
	},
	IsNoop: func(committeeVerifier *committee_verifier.CommitteeVerifier, opts *bind.CallOpts, args SignatureConfigArgs) (bool, error) {
		for _, cfg := range args.SignatureConfigUpdates {
			signers, threshold, err := committeeVerifier.GetSignatureConfig(opts, cfg.SourceChainSelector)
			if err != nil {
				return false, fmt.Errorf("failed to get signature configuration for source chain selector %d: %w", cfg.SourceChainSelector, err)
			}
			if threshold != cfg.Threshold {
				return false, nil
			}
			if len(signers) != len(cfg.Signers) {
				return false, nil
			}
			slices.SortFunc(signers, func(a, b common.Address) int {
				return bytes.Compare(a[:], b[:])
			})
			slices.SortFunc(cfg.Signers, func(a, b common.Address) int {
				return bytes.Compare(a[:], b[:])
			})
			if !slices.Equal(signers, cfg.Signers) {
				return false, nil
			}
		}

		for _, sourceChainSelector := range args.SourceChainSelectorsToRemove {
			signers, threshold, err := committeeVerifier.GetSignatureConfig(opts, sourceChainSelector)
			if err != nil {
				return false, fmt.Errorf("failed to get signature configuration for source chain selector %d: %w", sourceChainSelector, err)
			}
			if len(signers) != 0 || threshold != 0 {
				return false, nil
			}
		}

		return true, nil
	},
	CallContract: func(committeeVerifier *committee_verifier.CommitteeVerifier, opts *bind.TransactOpts, args SignatureConfigArgs) (*types.Transaction, error) {
		return committeeVerifier.ApplySignatureConfigs(opts, args.SourceChainSelectorsToRemove, args.SignatureConfigUpdates)
	},
})

var UpdateStorageLocations = contract.NewWrite(contract.WriteParams[[]string, *committee_verifier.CommitteeVerifier]{
	Name:            "committee-verifier:update-storage-locations",
	Version:         Version,
	Description:     "Updates the storage locations on the CommitteeVerifier",
	ContractType:    ContractType,
	ContractABI:     committee_verifier.CommitteeVerifierABI,
	NewContract:     committee_verifier.NewCommitteeVerifier,
	IsAllowedCaller: contract.OnlyOwner[*committee_verifier.CommitteeVerifier, []string],
	Validate: func(committeeVerifier *committee_verifier.CommitteeVerifier, backend bind.ContractBackend, opts *bind.CallOpts, args []string) error {
		return nil
	},
	IsNoop: func(committeeVerifier *committee_verifier.CommitteeVerifier, opts *bind.CallOpts, args []string) (bool, error) {
		storageLocations, err := committeeVerifier.GetStorageLocations(opts)
		if err != nil {
			return false, fmt.Errorf("failed to get storage locations: %w", err)
		}
		if len(storageLocations) != len(args) {
			return false, nil
		}
		slices.Sort(storageLocations)
		slices.Sort(args)
		if !slices.Equal(storageLocations, args) {
			return false, nil
		}
		return true, nil
	},
	CallContract: func(committeeVerifier *committee_verifier.CommitteeVerifier, opts *bind.TransactOpts, args []string) (*types.Transaction, error) {
		return committeeVerifier.UpdateStorageLocations(opts, args)
	},
})

var GetRemoteChainConfig = contract.NewRead(contract.ReadParams[uint64, RemoteChainConfig, *committee_verifier.CommitteeVerifier]{
	Name:         "committee-verifier:get-remote-chain-config",
	Version:      Version,
	Description:  "Gets the remote chain configuration for a given remote chain selector",
	ContractType: ContractType,
	NewContract:  committee_verifier.NewCommitteeVerifier,
	CallContract: func(committeeVerifier *committee_verifier.CommitteeVerifier, opts *bind.CallOpts, args uint64) (RemoteChainConfig, error) {
		return committeeVerifier.GetRemoteChainConfig(opts, args)
	},
})

var GetVersionTag = contract.NewRead(contract.ReadParams[any, [4]byte, *committee_verifier.CommitteeVerifier]{
	Name:         "committee-verifier:get-version-tag",
	Version:      Version,
	Description:  "Gets the version tag of the CommitteeVerifier contract",
	ContractType: ContractType,
	NewContract:  committee_verifier.NewCommitteeVerifier,
	CallContract: func(committeeVerifier *committee_verifier.CommitteeVerifier, opts *bind.CallOpts, args any) ([4]byte, error) {
		return committeeVerifier.VersionTag(opts)
	},
})

var GetSignatureConfig = contract.NewRead(contract.ReadParams[uint64, SignatureConfig, *committee_verifier.CommitteeVerifier]{
	Name:         "committee-verifier:get-signature-config",
	Version:      Version,
	Description:  "Gets the signature configuration for a given source chain selector",
	ContractType: ContractType,
	NewContract:  committee_verifier.NewCommitteeVerifier,
	CallContract: func(committeeVerifier *committee_verifier.CommitteeVerifier, opts *bind.CallOpts, args uint64) (SignatureConfig, error) {
		signers, threshold, err := committeeVerifier.GetSignatureConfig(opts, args)
		if err != nil {
			return SignatureConfig{}, fmt.Errorf("failed to get signature configuration for source chain selector %d: %w", args, err)
		}
		return SignatureConfig{
			Signers:             signers,
			Threshold:           threshold,
			SourceChainSelector: args,
		}, nil
	},
})
