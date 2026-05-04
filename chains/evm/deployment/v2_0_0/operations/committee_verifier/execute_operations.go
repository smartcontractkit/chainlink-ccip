package committee_verifier

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"

	gobindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v2_0_0/committee_verifier"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
)

// Type aliases for call sites (match historical names).
type (
	RemoteChainConfigArgs  = gobindings.BaseVerifierRemoteChainConfigArgs
	AllowlistConfigArgs    = gobindings.BaseVerifierAllowlistConfigArgs
	SignatureConfig        = gobindings.SignatureQuorumValidatorSignatureConfig
	DynamicConfig          = gobindings.CommitteeVerifierDynamicConfig
)

var GetDynamicConfig = contract.NewRead(contract.ReadParams[struct{}, gobindings.CommitteeVerifierDynamicConfig, *gobindings.CommitteeVerifier]{
	Name:         "committee-verifier:get-dynamic-config",
	Version:      Version,
	Description:  "Calls getDynamicConfig on the contract",
	ContractType: ContractType,
	NewContract:  gobindings.NewCommitteeVerifier,
	CallContract: func(c *gobindings.CommitteeVerifier, opts *bind.CallOpts, args struct{}) (gobindings.CommitteeVerifierDynamicConfig, error) {
		return c.GetDynamicConfig(opts)
	},
})

var SetDynamicConfig = contract.NewWrite(contract.WriteParams[gobindings.CommitteeVerifierDynamicConfig, *gobindings.CommitteeVerifier]{
	Name:            "committee-verifier:set-dynamic-config",
	Version:         Version,
	Description:     "Calls setDynamicConfig on the contract",
	ContractType:    ContractType,
	ContractABI:     gobindings.CommitteeVerifierMetaData.ABI,
	NewContract:     gobindings.NewCommitteeVerifier,
	IsAllowedCaller: contract.OnlyOwner[*gobindings.CommitteeVerifier, gobindings.CommitteeVerifierDynamicConfig],
	Validate:        func(gobindings.CommitteeVerifierDynamicConfig) error { return nil },
	CallContract: func(c *gobindings.CommitteeVerifier, opts *bind.TransactOpts, args gobindings.CommitteeVerifierDynamicConfig) (*types.Transaction, error) {
		return c.SetDynamicConfig(opts, args)
	},
})

var ApplyRemoteChainConfigUpdates = contract.NewWrite(contract.WriteParams[[]gobindings.BaseVerifierRemoteChainConfigArgs, *gobindings.CommitteeVerifier]{
	Name:            "committee-verifier:apply-remote-chain-config-updates",
	Version:         Version,
	Description:     "Calls applyRemoteChainConfigUpdates on the contract",
	ContractType:    ContractType,
	ContractABI:     gobindings.CommitteeVerifierMetaData.ABI,
	NewContract:     gobindings.NewCommitteeVerifier,
	IsAllowedCaller: contract.OnlyOwner[*gobindings.CommitteeVerifier, []gobindings.BaseVerifierRemoteChainConfigArgs],
	Validate:        func([]gobindings.BaseVerifierRemoteChainConfigArgs) error { return nil },
	CallContract: func(c *gobindings.CommitteeVerifier, opts *bind.TransactOpts, args []gobindings.BaseVerifierRemoteChainConfigArgs) (*types.Transaction, error) {
		return c.ApplyRemoteChainConfigUpdates(opts, args)
	},
})

var ApplyAllowlistUpdates = contract.NewWrite(contract.WriteParams[[]gobindings.BaseVerifierAllowlistConfigArgs, *gobindings.CommitteeVerifier]{
	Name:            "committee-verifier:apply-allowlist-updates",
	Version:         Version,
	Description:     "Calls applyAllowlistUpdates on the contract",
	ContractType:    ContractType,
	ContractABI:     gobindings.CommitteeVerifierMetaData.ABI,
	NewContract:     gobindings.NewCommitteeVerifier,
	IsAllowedCaller: contract.OnlyOwner[*gobindings.CommitteeVerifier, []gobindings.BaseVerifierAllowlistConfigArgs],
	Validate:        func([]gobindings.BaseVerifierAllowlistConfigArgs) error { return nil },
	CallContract: func(c *gobindings.CommitteeVerifier, opts *bind.TransactOpts, args []gobindings.BaseVerifierAllowlistConfigArgs) (*types.Transaction, error) {
		return c.ApplyAllowlistUpdates(opts, args)
	},
})

var SetAllowedFinalityConfig = contract.NewWrite(contract.WriteParams[[4]byte, *gobindings.CommitteeVerifier]{
	Name:            "committee-verifier:set-allowed-finality-config",
	Version:         Version,
	Description:     "Calls setAllowedFinalityConfig on the contract",
	ContractType:    ContractType,
	ContractABI:     gobindings.CommitteeVerifierMetaData.ABI,
	NewContract:     gobindings.NewCommitteeVerifier,
	IsAllowedCaller: contract.OnlyOwner[*gobindings.CommitteeVerifier, [4]byte],
	Validate:        func([4]byte) error { return nil },
	CallContract: func(c *gobindings.CommitteeVerifier, opts *bind.TransactOpts, args [4]byte) (*types.Transaction, error) {
		return c.SetAllowedFinalityConfig(opts, args)
	},
})

var ApplySignatureConfigs = contract.NewWrite(contract.WriteParams[ApplySignatureConfigsArgs, *gobindings.CommitteeVerifier]{
	Name:            "committee-verifier:apply-signature-configs",
	Version:         Version,
	Description:     "Calls applySignatureConfigs on the contract",
	ContractType:    ContractType,
	ContractABI:     gobindings.CommitteeVerifierMetaData.ABI,
	NewContract:     gobindings.NewCommitteeVerifier,
	IsAllowedCaller: contract.OnlyOwner[*gobindings.CommitteeVerifier, ApplySignatureConfigsArgs],
	Validate:        func(ApplySignatureConfigsArgs) error { return nil },
	CallContract: func(c *gobindings.CommitteeVerifier, opts *bind.TransactOpts, args ApplySignatureConfigsArgs) (*types.Transaction, error) {
		return c.ApplySignatureConfigs(opts, args.SourceChainSelectorsToRemove, args.SignatureConfigs)
	},
})

var GetRemoteChainConfig = contract.NewRead(contract.ReadParams[uint64, GetRemoteChainConfigResult, *gobindings.CommitteeVerifier]{
	Name:         "committee-verifier:get-remote-chain-config",
	Version:      Version,
	Description:  "Calls getRemoteChainConfig on the contract",
	ContractType: ContractType,
	NewContract:  gobindings.NewCommitteeVerifier,
	CallContract: func(c *gobindings.CommitteeVerifier, opts *bind.CallOpts, args uint64) (GetRemoteChainConfigResult, error) {
		res, err := c.GetRemoteChainConfig(opts, args)
		if err != nil {
			return GetRemoteChainConfigResult{}, err
		}
		return GetRemoteChainConfigResult{RemoteChainConfig: res.RemoteChainConfig, AllowedSendersList: res.AllowedSendersList}, nil
	},
})

var GetFee = contract.NewRead(contract.ReadParams[GetFeeArgs, GetFeeResult, *gobindings.CommitteeVerifier]{
	Name:         "committee-verifier:get-fee",
	Version:      Version,
	Description:  "Calls getFee on the contract",
	ContractType: ContractType,
	NewContract:  gobindings.NewCommitteeVerifier,
	CallContract: func(c *gobindings.CommitteeVerifier, opts *bind.CallOpts, args GetFeeArgs) (GetFeeResult, error) {
		res, err := c.GetFee(opts, args.DestChainSelector, args.Arg1, args.Arg2, args.RequestedFinality)
		if err != nil {
			return GetFeeResult{}, err
		}
		return GetFeeResult{FeeUSDCents: res.FeeUSDCents, GasForVerification: res.GasForVerification, PayloadSizeBytes: res.PayloadSizeBytes}, nil
	},
})

var GetAllowedFinalityConfig = contract.NewRead(contract.ReadParams[struct{}, [4]byte, *gobindings.CommitteeVerifier]{
	Name:         "committee-verifier:get-allowed-finality-config",
	Version:      Version,
	Description:  "Calls getAllowedFinalityConfig on the contract",
	ContractType: ContractType,
	NewContract:  gobindings.NewCommitteeVerifier,
	CallContract: func(c *gobindings.CommitteeVerifier, opts *bind.CallOpts, args struct{}) ([4]byte, error) {
		return c.GetAllowedFinalityConfig(opts)
	},
})

var GetSignatureConfig = contract.NewRead(contract.ReadParams[uint64, GetSignatureConfigResult, *gobindings.CommitteeVerifier]{
	Name:         "committee-verifier:get-signature-config",
	Version:      Version,
	Description:  "Calls getSignatureConfig on the contract",
	ContractType: ContractType,
	NewContract:  gobindings.NewCommitteeVerifier,
	CallContract: func(c *gobindings.CommitteeVerifier, opts *bind.CallOpts, args uint64) (GetSignatureConfigResult, error) {
		res, err := c.GetSignatureConfig(opts, args)
		if err != nil {
			return GetSignatureConfigResult{}, err
		}
		return GetSignatureConfigResult{Signers: res.Signers, Threshold: res.Threshold}, nil
	},
})

var VersionTag = contract.NewRead(contract.ReadParams[struct{}, [4]byte, *gobindings.CommitteeVerifier]{
	Name:         "committee-verifier:version-tag",
	Version:      Version,
	Description:  "Calls versionTag on the contract",
	ContractType: ContractType,
	NewContract:  gobindings.NewCommitteeVerifier,
	CallContract: func(c *gobindings.CommitteeVerifier, opts *bind.CallOpts, args struct{}) ([4]byte, error) {
		return c.VersionTag(opts)
	},
})
