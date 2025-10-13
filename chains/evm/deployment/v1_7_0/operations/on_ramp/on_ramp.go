package on_ramp

import (
	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/latest/on_ramp"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

var ContractType cldf_deployment.ContractType = "OnRamp"

type StaticConfig = on_ramp.OnRampStaticConfig

type DynamicConfig = on_ramp.OnRampDynamicConfig

type ConstructorArgs struct {
	StaticConfig  StaticConfig
	DynamicConfig DynamicConfig
}

type SetDynamicConfigArgs struct {
	DynamicConfig DynamicConfig
}

type DestChainConfigArgs = on_ramp.OnRampDestChainConfigArgs

type WithdrawFeeTokensArgs struct {
	FeeTokens []common.Address
}

type DestChainConfig = on_ramp.OnRampDestChainConfig

var Deploy = contract.NewDeploy(contract.DeployParams[ConstructorArgs]{
	Name:             "on-ramp:deploy",
	Version:          semver.MustParse("1.7.0"),
	Description:      "Deploys the OnRamp contract",
	ContractMetadata: on_ramp.OnRampMetaData,
	BytecodeByTypeAndVersion: map[string]contract.Bytecode{
		cldf_deployment.NewTypeAndVersion(ContractType, *semver.MustParse("1.7.0")).String(): {
			EVM: common.FromHex(on_ramp.OnRampBin),
		},
	},
	Validate: func(ConstructorArgs) error { return nil },
})

var SetDynamicConfig = contract.NewWrite(contract.WriteParams[SetDynamicConfigArgs, *on_ramp.OnRamp]{
	Name:            "on-ramp:set-dynamic-config",
	Version:         semver.MustParse("1.7.0"),
	Description:     "Sets the dynamic configuration on the OnRamp",
	ContractType:    ContractType,
	ContractABI:     on_ramp.OnRampABI,
	NewContract:     on_ramp.NewOnRamp,
	IsAllowedCaller: contract.OnlyOwner[*on_ramp.OnRamp, SetDynamicConfigArgs],
	Validate:        func(SetDynamicConfigArgs) error { return nil },
	CallContract: func(onRamp *on_ramp.OnRamp, opts *bind.TransactOpts, args SetDynamicConfigArgs) (*types.Transaction, error) {
		return onRamp.SetDynamicConfig(opts, args.DynamicConfig)
	},
})

var ApplyDestChainConfigUpdates = contract.NewWrite(contract.WriteParams[[]DestChainConfigArgs, *on_ramp.OnRamp]{
	Name:            "on-ramp:apply-dest-chain-config-updates",
	Version:         semver.MustParse("1.7.0"),
	Description:     "Applies updates to destination chain configuration on the OnRamp",
	ContractType:    ContractType,
	ContractABI:     on_ramp.OnRampABI,
	NewContract:     on_ramp.NewOnRamp,
	IsAllowedCaller: contract.OnlyOwner[*on_ramp.OnRamp, []DestChainConfigArgs],
	Validate:        func([]DestChainConfigArgs) error { return nil },
	CallContract: func(onRamp *on_ramp.OnRamp, opts *bind.TransactOpts, args []DestChainConfigArgs) (*types.Transaction, error) {
		return onRamp.ApplyDestChainConfigUpdates(opts, args)
	},
})

var WithdrawFeeTokens = contract.NewWrite(contract.WriteParams[WithdrawFeeTokensArgs, *on_ramp.OnRamp]{
	Name:            "on-ramp:withdraw-fee-tokens",
	Version:         semver.MustParse("1.7.0"),
	Description:     "Withdraws fee tokens from the OnRamp",
	ContractType:    ContractType,
	ContractABI:     on_ramp.OnRampABI,
	NewContract:     on_ramp.NewOnRamp,
	IsAllowedCaller: contract.OnlyOwner[*on_ramp.OnRamp, WithdrawFeeTokensArgs],
	Validate:        func(WithdrawFeeTokensArgs) error { return nil },
	CallContract: func(onRamp *on_ramp.OnRamp, opts *bind.TransactOpts, args WithdrawFeeTokensArgs) (*types.Transaction, error) {
		return onRamp.WithdrawFeeTokens(opts, args.FeeTokens)
	},
})

var GetDestChainConfig = contract.NewRead(contract.ReadParams[uint64, DestChainConfig, *on_ramp.OnRamp]{
	Name:         "on-ramp:get-dest-chain-config",
	Version:      semver.MustParse("1.7.0"),
	Description:  "Gets the destination chain configuration for a given destination chain selector",
	ContractType: ContractType,
	NewContract:  on_ramp.NewOnRamp,
	CallContract: func(onRamp *on_ramp.OnRamp, opts *bind.CallOpts, args uint64) (DestChainConfig, error) {
		return onRamp.GetDestChainConfig(opts, args)
	},
})
