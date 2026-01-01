package fee_quoter

import (
	"errors"
	"fmt"
	"slices"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/gobindings/generated/latest/fee_quoter"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/adapters"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

const minLockOrBurnRetBytes = 32

var ContractType cldf_deployment.ContractType = "FeeQuoter"

var Version = semver.MustParse("1.7.0")

type StaticConfig = fee_quoter.FeeQuoterStaticConfig

type DestChainConfigArgs struct {
	DestChainSelector uint64
	DestChainConfig   adapters.FeeQuoterDestChainConfig
}

type TokenTransferFeeConfigArgs = fee_quoter.FeeQuoterTokenTransferFeeConfigArgs

type TokenTransferFeeConfigRemoveArgs = fee_quoter.FeeQuoterTokenTransferFeeConfigRemoveArgs

type ConstructorArgs struct {
	StaticConfig               StaticConfig
	PriceUpdaters              []common.Address
	TokenTransferFeeConfigArgs []TokenTransferFeeConfigArgs
	DestChainConfigArgs        []DestChainConfigArgs
}

type ApplyFeeTokensUpdatesArgs struct {
	FeeTokensToAdd    []common.Address
	FeeTokensToRemove []common.Address
}

type AuthorizedCallerArgs = fee_quoter.AuthorizedCallersAuthorizedCallerArgs

type ApplyTokenTransferFeeConfigUpdatesArgs struct {
	TokenTransferFeeConfigArgs   []TokenTransferFeeConfigArgs
	TokensToUseDefaultFeeConfigs []TokenTransferFeeConfigRemoveArgs
}

type TokenPriceUpdate = fee_quoter.InternalTokenPriceUpdate

type GasPriceUpdate = fee_quoter.InternalGasPriceUpdate

type PriceUpdates = fee_quoter.InternalPriceUpdates

var Deploy = contract.NewDeploy(contract.DeployParams[ConstructorArgs]{
	Name:             "fee-quoter-v2:deploy",
	Version:          Version,
	Description:      "Deploys the FeeQuoter contract",
	ContractMetadata: fee_quoter.FeeQuoterMetaData,
	BytecodeByTypeAndVersion: map[string]contract.Bytecode{
		cldf_deployment.NewTypeAndVersion(ContractType, *Version).String(): {
			EVM: common.FromHex(fee_quoter.FeeQuoterBin),
		},
	},
	Validate: func(ConstructorArgs) error { return nil },
})

var ApplyAuthorizedCallerUpdates = contract.NewWrite(contract.WriteParams[AuthorizedCallerArgs, *fee_quoter.FeeQuoter]{
	Name:            "fee-quoter-v2:apply-authorized-caller-updates",
	Version:         Version,
	Description:     "Updates authorized price updaters on the FeeQuoter contract",
	ContractType:    ContractType,
	ContractABI:     fee_quoter.FeeQuoterABI,
	NewContract:     fee_quoter.NewFeeQuoter,
	IsAllowedCaller: contract.OnlyOwner[*fee_quoter.FeeQuoter, AuthorizedCallerArgs],
	Validate: func(feeQuoter *fee_quoter.FeeQuoter, backend bind.ContractBackend, opts *bind.CallOpts, args AuthorizedCallerArgs) error {
		for _, caller := range args.AddedCallers {
			if caller == (common.Address{}) {
				return errors.New("caller cannot be the zero address")
			}
		}
		return nil
	},
	IsNoop: func(feeQuoter *fee_quoter.FeeQuoter, opts *bind.CallOpts, args AuthorizedCallerArgs) (bool, error) {
		allowedCallers, err := feeQuoter.GetAllAuthorizedCallers(opts)
		if err != nil {
			return false, fmt.Errorf("failed to get all authorized callers: %w", err)
		}
		for _, caller := range args.AddedCallers {
			if !slices.Contains(allowedCallers, caller) {
				return false, nil
			}
		}
		return true, nil
	},
	CallContract: func(feeQuoter *fee_quoter.FeeQuoter, opts *bind.TransactOpts, args AuthorizedCallerArgs) (*types.Transaction, error) {
		return feeQuoter.ApplyAuthorizedCallerUpdates(opts, args)
	},
})

var ApplyDestChainConfigUpdates = contract.NewWrite(contract.WriteParams[[]DestChainConfigArgs, *fee_quoter.FeeQuoter]{
	Name:            "fee-quoter-v2:apply-dest-chain-config-updates",
	Version:         Version,
	Description:     "Applies updates to destination chain configurations on the FeeQuoter",
	ContractType:    ContractType,
	ContractABI:     fee_quoter.FeeQuoterABI,
	NewContract:     fee_quoter.NewFeeQuoter,
	IsAllowedCaller: contract.OnlyOwner[*fee_quoter.FeeQuoter, []DestChainConfigArgs],
	Validate: func(feeQuoter *fee_quoter.FeeQuoter, backend bind.ContractBackend, opts *bind.CallOpts, args []DestChainConfigArgs) error {
		for _, arg := range args {
			if arg.DestChainSelector == 0 {
				return errors.New("dest chain selector cannot be 0")
			}
			if arg.DestChainConfig.DefaultTxGasLimit == 0 {
				return errors.New("default tx gas limit cannot be 0")
			}
			if arg.DestChainConfig.DefaultTxGasLimit > arg.DestChainConfig.MaxPerMsgGasLimit {
				return errors.New("default tx gas limit cannot be greater than max per msg gas limit")
			}
		}
		return nil
	},
	IsNoop: func(feeQuoter *fee_quoter.FeeQuoter, opts *bind.CallOpts, args []DestChainConfigArgs) (bool, error) {
		argsTransformed := transformDestChainConfigArgs(args)
		for _, arg := range argsTransformed {
			actualDestChainConfig, err := feeQuoter.GetDestChainConfig(opts, arg.DestChainSelector)
			if err != nil {
				return false, fmt.Errorf("failed to get dest chain config for dest chain selector %d: %w", arg.DestChainSelector, err)
			}
			if actualDestChainConfig != arg.DestChainConfig {
				return false, nil
			}
		}

		return true, nil
	},
	CallContract: func(feeQuoter *fee_quoter.FeeQuoter, opts *bind.TransactOpts, args []DestChainConfigArgs) (*types.Transaction, error) {
		return feeQuoter.ApplyDestChainConfigUpdates(opts, transformDestChainConfigArgs(args))
	},
})

var ApplyTokenTransferFeeConfigUpdates = contract.NewWrite(contract.WriteParams[ApplyTokenTransferFeeConfigUpdatesArgs, *fee_quoter.FeeQuoter]{
	Name:            "fee-quoter-v2:apply-token-transfer-fee-config-updates",
	Version:         Version,
	Description:     "Applies updates to the token transfer fee configurations on the FeeQuoter",
	ContractType:    ContractType,
	ContractABI:     fee_quoter.FeeQuoterABI,
	NewContract:     fee_quoter.NewFeeQuoter,
	IsAllowedCaller: contract.OnlyOwner[*fee_quoter.FeeQuoter, ApplyTokenTransferFeeConfigUpdatesArgs],
	Validate: func(feeQuoter *fee_quoter.FeeQuoter, backend bind.ContractBackend, opts *bind.CallOpts, args ApplyTokenTransferFeeConfigUpdatesArgs) error {
		for _, arg := range args.TokenTransferFeeConfigArgs {
			for _, tokenTransferFeeConfig := range arg.TokenTransferFeeConfigs {
				if tokenTransferFeeConfig.TokenTransferFeeConfig.DestBytesOverhead < minLockOrBurnRetBytes {
					return fmt.Errorf("dest bytes overhead cannot be less than %d", minLockOrBurnRetBytes)
				}
			}
		}
		return nil
	},
	IsNoop: func(feeQuoter *fee_quoter.FeeQuoter, opts *bind.CallOpts, args ApplyTokenTransferFeeConfigUpdatesArgs) (bool, error) {
		for _, arg := range args.TokenTransferFeeConfigArgs {
			for _, tokenTransferFeeConfig := range arg.TokenTransferFeeConfigs {
				actualTokenTransferFeeConfig, err := feeQuoter.GetTokenTransferFee(opts, arg.DestChainSelector, tokenTransferFeeConfig.Token)
				if err != nil {
					return false, fmt.Errorf("failed to get token transfer fee for dest chain selector %d and token %s: %w", arg.DestChainSelector, tokenTransferFeeConfig.Token, err)
				}
				if actualTokenTransferFeeConfig.FeeUSDCents != tokenTransferFeeConfig.TokenTransferFeeConfig.FeeUSDCents ||
					actualTokenTransferFeeConfig.DestGasOverhead != tokenTransferFeeConfig.TokenTransferFeeConfig.DestGasOverhead ||
					actualTokenTransferFeeConfig.DestBytesOverhead != tokenTransferFeeConfig.TokenTransferFeeConfig.DestBytesOverhead {
					return false, nil
				}
			}
		}
		for _, arg := range args.TokensToUseDefaultFeeConfigs {
			actualTokenTransferFeeConfig, err := feeQuoter.GetTokenTransferFee(opts, arg.DestChainSelector, arg.Token)
			if err != nil {
				return false, fmt.Errorf("failed to get token transfer fee for dest chain selector %d and token %s: %w", arg.DestChainSelector, arg.Token, err)
			}
			if actualTokenTransferFeeConfig.FeeUSDCents != 0 ||
				actualTokenTransferFeeConfig.DestGasOverhead != 0 ||
				actualTokenTransferFeeConfig.DestBytesOverhead != 0 {
				return false, nil
			}
		}

		return true, nil
	},
	CallContract: func(feeQuoter *fee_quoter.FeeQuoter, opts *bind.TransactOpts, args ApplyTokenTransferFeeConfigUpdatesArgs) (*types.Transaction, error) {
		return feeQuoter.ApplyTokenTransferFeeConfigUpdates(opts, args.TokenTransferFeeConfigArgs, args.TokensToUseDefaultFeeConfigs)
	},
})

var UpdatePrices = contract.NewWrite(contract.WriteParams[PriceUpdates, *fee_quoter.FeeQuoter]{
	Name:         "fee-quoter-v2:update-prices",
	Version:      Version,
	Description:  "Updates token prices on the FeeQuoter",
	ContractType: ContractType,
	ContractABI:  fee_quoter.FeeQuoterABI,
	NewContract:  fee_quoter.NewFeeQuoter,
	IsAllowedCaller: func(feeQuoter *fee_quoter.FeeQuoter, opts *bind.CallOpts, caller common.Address, args PriceUpdates) (bool, error) {
		priceUpdaters, err := feeQuoter.GetAllAuthorizedCallers(opts)
		if err != nil {
			return false, fmt.Errorf("failed to get authorized callers from FeeQuoter (%s): %w", feeQuoter.Address(), err)
		}
		if slices.Contains(priceUpdaters, caller) {
			return true, nil
		}
		return false, nil
	},
	Validate: func(feeQuoter *fee_quoter.FeeQuoter, backend bind.ContractBackend, opts *bind.CallOpts, args PriceUpdates) error {
		return nil
	},
	IsNoop: func(feeQuoter *fee_quoter.FeeQuoter, opts *bind.CallOpts, args PriceUpdates) (bool, error) {
		for _, arg := range args.TokenPriceUpdates {
			actualTokenPrice, err := feeQuoter.GetTokenPrice(opts, arg.SourceToken)
			if err != nil {
				return false, fmt.Errorf("failed to get token price for source token %s: %w", arg.SourceToken, err)
			}
			if actualTokenPrice.Value.Cmp(arg.UsdPerToken) != 0 {
				return false, nil
			}
		}
		for _, arg := range args.GasPriceUpdates {
			actualGasPrice, err := feeQuoter.GetDestinationChainGasPrice(opts, arg.DestChainSelector)
			if err != nil {
				return false, fmt.Errorf("failed to get gas price for dest chain selector %d: %w", arg.DestChainSelector, err)
			}
			if actualGasPrice.Value.Cmp(arg.UsdPerUnitGas) != 0 {
				return false, nil
			}
		}

		return true, nil
	},
	CallContract: func(feeQuoter *fee_quoter.FeeQuoter, opts *bind.TransactOpts, args PriceUpdates) (*types.Transaction, error) {
		return feeQuoter.UpdatePrices(opts, args)
	},
})

func transformDestChainConfigArgs(args []DestChainConfigArgs) []fee_quoter.FeeQuoterDestChainConfigArgs {
	// The reason we avoid exposing fee_quoter.FeeQuoterDestChainConfigArgs directly
	// is so that consumers do not need to import the gobindings package.
	argsTransformed := make([]fee_quoter.FeeQuoterDestChainConfigArgs, 0, len(args))
	for _, arg := range args {
		argsTransformed = append(argsTransformed, fee_quoter.FeeQuoterDestChainConfigArgs{
			DestChainSelector: arg.DestChainSelector,
			DestChainConfig: fee_quoter.FeeQuoterDestChainConfig{
				IsEnabled:                   arg.DestChainConfig.IsEnabled,
				MaxDataBytes:                arg.DestChainConfig.MaxDataBytes,
				MaxPerMsgGasLimit:           arg.DestChainConfig.MaxPerMsgGasLimit,
				DestGasOverhead:             arg.DestChainConfig.DestGasOverhead,
				DestGasPerPayloadByteBase:   arg.DestChainConfig.DestGasPerPayloadByteBase,
				ChainFamilySelector:         arg.DestChainConfig.ChainFamilySelector,
				DefaultTokenFeeUSDCents:     arg.DestChainConfig.DefaultTokenFeeUSDCents,
				DefaultTokenDestGasOverhead: arg.DestChainConfig.DefaultTokenDestGasOverhead,
				DefaultTxGasLimit:           arg.DestChainConfig.DefaultTxGasLimit,
				NetworkFeeUSDCents:          arg.DestChainConfig.NetworkFeeUSDCents,
				LinkFeeMultiplierPercent:    arg.DestChainConfig.LinkFeeMultiplierPercent,
			},
		})
	}

	return argsTransformed
}
