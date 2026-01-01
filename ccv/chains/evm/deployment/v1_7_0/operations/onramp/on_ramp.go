package onramp

import (
	"bytes"
	"fmt"
	"slices"

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
	Validate: func(onRamp *onramp.OnRamp, backend bind.ContractBackend, opts *bind.CallOpts, args SetDynamicConfigArgs) error {
		return nil
	},
	IsNoop: func(onRamp *onramp.OnRamp, opts *bind.CallOpts, args SetDynamicConfigArgs) (bool, error) {
		actualDynamicConfig, err := onRamp.GetDynamicConfig(opts)
		if err != nil {
			return false, fmt.Errorf("failed to get dynamic config: %w", err)
		}
		return actualDynamicConfig == args.DynamicConfig, nil
	},
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
	Validate: func(onRamp *onramp.OnRamp, backend bind.ContractBackend, opts *bind.CallOpts, args []DestChainConfigArgs) error {
		return nil
	},
	IsNoop: func(onRamp *onramp.OnRamp, opts *bind.CallOpts, args []DestChainConfigArgs) (bool, error) {
		for _, arg := range args {
			actualDestChainConfig, err := onRamp.GetDestChainConfig(opts, arg.DestChainSelector)
			if err != nil {
				return false, fmt.Errorf("failed to get dest chain config: %w", err)
			}
			if actualDestChainConfig.AddressBytesLength != arg.AddressBytesLength ||
				actualDestChainConfig.TokenReceiverAllowed != arg.TokenReceiverAllowed ||
				actualDestChainConfig.MessageNetworkFeeUSDCents != arg.MessageNetworkFeeUSDCents ||
				actualDestChainConfig.TokenNetworkFeeUSDCents != arg.TokenNetworkFeeUSDCents ||
				actualDestChainConfig.BaseExecutionGasCost != arg.BaseExecutionGasCost ||
				actualDestChainConfig.DefaultExecutor != arg.DefaultExecutor ||
				actualDestChainConfig.Router != arg.Router ||
				!bytes.Equal(actualDestChainConfig.OffRamp, arg.OffRamp) {
				return false, nil
			}
			slices.SortFunc(actualDestChainConfig.DefaultCCVs, func(a, b common.Address) int {
				return bytes.Compare(a[:], b[:])
			})
			slices.SortFunc(actualDestChainConfig.LaneMandatedCCVs, func(a, b common.Address) int {
				return bytes.Compare(a[:], b[:])
			})
			slices.SortFunc(arg.DefaultCCVs, func(a, b common.Address) int {
				return bytes.Compare(a[:], b[:])
			})
			slices.SortFunc(arg.LaneMandatedCCVs, func(a, b common.Address) int {
				return bytes.Compare(a[:], b[:])
			})
			if !slices.Equal(actualDestChainConfig.DefaultCCVs, arg.DefaultCCVs) ||
				!slices.Equal(actualDestChainConfig.LaneMandatedCCVs, arg.LaneMandatedCCVs) {
				return false, nil
			}
		}
		return true, nil
	},
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
	Validate: func(onRamp *onramp.OnRamp, backend bind.ContractBackend, opts *bind.CallOpts, args WithdrawFeeTokensArgs) error {
		return nil
	},
	IsNoop: func(onRamp *onramp.OnRamp, opts *bind.CallOpts, args WithdrawFeeTokensArgs) (bool, error) {
		return false, nil
	},
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
