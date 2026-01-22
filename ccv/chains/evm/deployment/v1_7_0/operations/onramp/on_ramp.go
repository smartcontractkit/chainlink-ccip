package onramp

import (
	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/gobindings/generated/latest/onramp"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

var ContractType cldf_deployment.ContractType = "OnRamp"

var Version = semver.MustParse("1.7.0")

type StaticConfig = onramp.OnRampStaticConfig

type DynamicConfig = onramp.OnRampDynamicConfig

type ConstructorArgs struct {
	StaticConfig  StaticConfig
	DynamicConfig DynamicConfig
}

type SetDynamicConfigArgs struct {
	DynamicConfig DynamicConfig
}

type DestChainConfigArgs = onramp.OnRampDestChainConfigArgs

type WithdrawFeeTokensArgs struct {
	FeeTokens []common.Address
}

type DestChainConfig = onramp.OnRampDestChainConfig

// GetAllDestChainConfigsResult is the return type for GetAllDestChainConfigs
type GetAllDestChainConfigsResult struct {
	DestChainSelectors []uint64
	DestChainConfigs   []DestChainConfig
}

var Deploy = contract.NewDeploy(contract.DeployParams[ConstructorArgs]{
	Name:             "on-ramp:deploy",
	Version:          Version,
	Description:      "Deploys the OnRamp contract",
	ContractMetadata: onramp.OnRampMetaData,
	BytecodeByTypeAndVersion: map[string]contract.Bytecode{
		cldf_deployment.NewTypeAndVersion(ContractType, *Version).String(): {
			EVM: common.FromHex(onramp.OnRampBin),
		},
	},
	Validate: func(ConstructorArgs) error { return nil },
})

var SetDynamicConfig = contract.NewWrite(contract.WriteParams[SetDynamicConfigArgs, *onramp.OnRamp]{
	Name:            "on-ramp:set-dynamic-config",
	Version:         Version,
	Description:     "Sets the dynamic configuration on the OnRamp",
	ContractType:    ContractType,
	ContractABI:     onramp.OnRampABI,
	NewContract:     onramp.NewOnRamp,
	IsAllowedCaller: contract.OnlyOwner[*onramp.OnRamp, SetDynamicConfigArgs],
	Validate:        func(SetDynamicConfigArgs) error { return nil },
	CallContract: func(onRamp *onramp.OnRamp, opts *bind.TransactOpts, args SetDynamicConfigArgs) (*types.Transaction, error) {
		return onRamp.SetDynamicConfig(opts, args.DynamicConfig)
	},
})

var ApplyDestChainConfigUpdates = contract.NewWrite(contract.WriteParams[[]DestChainConfigArgs, *onramp.OnRamp]{
	Name:            "on-ramp:apply-dest-chain-config-updates",
	Version:         Version,
	Description:     "Applies updates to destination chain configuration on the OnRamp",
	ContractType:    ContractType,
	ContractABI:     onramp.OnRampABI,
	NewContract:     onramp.NewOnRamp,
	IsAllowedCaller: contract.OnlyOwner[*onramp.OnRamp, []DestChainConfigArgs],
	Validate:        func([]DestChainConfigArgs) error { return nil },
	CallContract: func(onRamp *onramp.OnRamp, opts *bind.TransactOpts, args []DestChainConfigArgs) (*types.Transaction, error) {
		return onRamp.ApplyDestChainConfigUpdates(opts, args)
	},
})

var WithdrawFeeTokens = contract.NewWrite(contract.WriteParams[WithdrawFeeTokensArgs, *onramp.OnRamp]{
	Name:            "on-ramp:withdraw-fee-tokens",
	Version:         Version,
	Description:     "Withdraws fee tokens from the OnRamp",
	ContractType:    ContractType,
	ContractABI:     onramp.OnRampABI,
	NewContract:     onramp.NewOnRamp,
	IsAllowedCaller: contract.OnlyOwner[*onramp.OnRamp, WithdrawFeeTokensArgs],
	Validate:        func(WithdrawFeeTokensArgs) error { return nil },
	CallContract: func(onRamp *onramp.OnRamp, opts *bind.TransactOpts, args WithdrawFeeTokensArgs) (*types.Transaction, error) {
		return onRamp.WithdrawFeeTokens(opts, args.FeeTokens)
	},
})

var GetDestChainConfig = contract.NewRead(contract.ReadParams[uint64, DestChainConfig, *onramp.OnRamp]{
	Name:         "on-ramp:get-dest-chain-config",
	Version:      Version,
	Description:  "Gets the destination chain configuration for a given destination chain selector",
	ContractType: ContractType,
	NewContract:  onramp.NewOnRamp,
	CallContract: func(onRamp *onramp.OnRamp, opts *bind.CallOpts, args uint64) (DestChainConfig, error) {
		return onRamp.GetDestChainConfig(opts, args)
	},
})

// GetAllDestChainConfigs reads all destination chain configurations from the OnRamp.
// NOTE: This operation is defined but not yet functional until the gobindings are regenerated
// after the contract's getAllDestChainConfigs() function is merged.
var GetAllDestChainConfigs = contract.NewRead(contract.ReadParams[any, GetAllDestChainConfigsResult, *onramp.OnRamp]{
	Name:         "on-ramp:get-all-dest-chain-configs",
	Version:      Version,
	Description:  "Gets all destination chain configurations from the OnRamp",
	ContractType: ContractType,
	NewContract:  onramp.NewOnRamp,
	CallContract: func(onRamp *onramp.OnRamp, opts *bind.CallOpts, _ any) (GetAllDestChainConfigsResult, error) {
		selectors, configs, err := onRamp.GetAllDestChainConfigs(opts)
		if err != nil {
			return GetAllDestChainConfigsResult{}, err
		}
		return GetAllDestChainConfigsResult{
			DestChainSelectors: selectors,
			DestChainConfigs:   configs,
		}, nil
	},
})
