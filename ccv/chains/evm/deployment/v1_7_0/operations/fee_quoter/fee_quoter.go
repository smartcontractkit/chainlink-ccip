package fee_quoter

import (
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

var ContractType cldf_deployment.ContractType = "FeeQuoter"

type StaticConfig = fee_quoter.FeeQuoterStaticConfig

type DestChainConfigArgs struct {
	DestChainSelector uint64
	DestChainConfig   adapters.FeeQuoterDestChainConfig
}

type FeeTokenArgs = fee_quoter.FeeQuoterFeeTokenArgs

type TokenTransferFeeConfigArgs = fee_quoter.FeeQuoterTokenTransferFeeConfigArgs

type TokenTransferFeeConfigRemoveArgs = fee_quoter.FeeQuoterTokenTransferFeeConfigRemoveArgs

type ConstructorArgs struct {
	StaticConfig               StaticConfig
	PriceUpdaters              []common.Address
	FeeTokens                  []FeeTokenArgs
	TokenTransferFeeConfigArgs []TokenTransferFeeConfigArgs
	DestChainConfigArgs        []DestChainConfigArgs
}

type ApplyFeeTokensUpdatesArgs struct {
	FeeTokensToAdd    []FeeTokenArgs
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
	Version:          semver.MustParse("1.7.0"),
	Description:      "Deploys the FeeQuoter contract",
	ContractMetadata: fee_quoter.FeeQuoterMetaData,
	BytecodeByTypeAndVersion: map[string]contract.Bytecode{
		cldf_deployment.NewTypeAndVersion(ContractType, *semver.MustParse("1.7.0")).String(): {
			EVM: common.FromHex(fee_quoter.FeeQuoterBin),
		},
	},
	Validate: func(ConstructorArgs) error { return nil },
})

var ApplyAuthorizedCallerUpdates = contract.NewWrite(contract.WriteParams[AuthorizedCallerArgs, *fee_quoter.FeeQuoter]{
	Name:            "fee-quoter-v2:apply-authorized-caller-updates",
	Version:         semver.MustParse("1.7.0"),
	Description:     "Updates authorized price updaters on the FeeQuoter contract",
	ContractType:    ContractType,
	ContractABI:     fee_quoter.FeeQuoterABI,
	NewContract:     fee_quoter.NewFeeQuoter,
	IsAllowedCaller: contract.OnlyOwner[*fee_quoter.FeeQuoter, AuthorizedCallerArgs],
	Validate:        func(AuthorizedCallerArgs) error { return nil },
	CallContract: func(FeeQuoter *fee_quoter.FeeQuoter, opts *bind.TransactOpts, args AuthorizedCallerArgs) (*types.Transaction, error) {
		return FeeQuoter.ApplyAuthorizedCallerUpdates(opts, args)
	},
})

var ApplyDestChainConfigUpdates = contract.NewWrite(contract.WriteParams[[]DestChainConfigArgs, *fee_quoter.FeeQuoter]{
	Name:            "fee-quoter-v2:apply-dest-chain-config-updates",
	Version:         semver.MustParse("1.7.0"),
	Description:     "Applies updates to destination chain configurations on the FeeQuoter",
	ContractType:    ContractType,
	ContractABI:     fee_quoter.FeeQuoterABI,
	NewContract:     fee_quoter.NewFeeQuoter,
	IsAllowedCaller: contract.OnlyOwner[*fee_quoter.FeeQuoter, []DestChainConfigArgs],
	Validate:        func([]DestChainConfigArgs) error { return nil },
	CallContract: func(FeeQuoter *fee_quoter.FeeQuoter, opts *bind.TransactOpts, args []DestChainConfigArgs) (*types.Transaction, error) {
		return FeeQuoter.ApplyDestChainConfigUpdates(opts, transformDestChainConfigArgs(args))
	},
})

var ApplyFeeTokensUpdates = contract.NewWrite(contract.WriteParams[ApplyFeeTokensUpdatesArgs, *fee_quoter.FeeQuoter]{
	Name:            "fee-quoter-v2:apply-fee-tokens-updates",
	Version:         semver.MustParse("1.7.0"),
	Description:     "Applies updates to the fee tokens supported by the FeeQuoter",
	ContractType:    ContractType,
	ContractABI:     fee_quoter.FeeQuoterABI,
	NewContract:     fee_quoter.NewFeeQuoter,
	IsAllowedCaller: contract.OnlyOwner[*fee_quoter.FeeQuoter, ApplyFeeTokensUpdatesArgs],
	Validate:        func(ApplyFeeTokensUpdatesArgs) error { return nil },
	CallContract: func(FeeQuoter *fee_quoter.FeeQuoter, opts *bind.TransactOpts, args ApplyFeeTokensUpdatesArgs) (*types.Transaction, error) {
		return FeeQuoter.ApplyFeeTokensUpdates(opts, args.FeeTokensToRemove, args.FeeTokensToAdd)
	},
})

var ApplyTokenTransferFeeConfigUpdates = contract.NewWrite(contract.WriteParams[ApplyTokenTransferFeeConfigUpdatesArgs, *fee_quoter.FeeQuoter]{
	Name:            "fee-quoter-v2:apply-token-transfer-fee-config-updates",
	Version:         semver.MustParse("1.7.0"),
	Description:     "Applies updates to the token transfer fee configurations on the FeeQuoter",
	ContractType:    ContractType,
	ContractABI:     fee_quoter.FeeQuoterABI,
	NewContract:     fee_quoter.NewFeeQuoter,
	IsAllowedCaller: contract.OnlyOwner[*fee_quoter.FeeQuoter, ApplyTokenTransferFeeConfigUpdatesArgs],
	Validate:        func(ApplyTokenTransferFeeConfigUpdatesArgs) error { return nil },
	CallContract: func(FeeQuoter *fee_quoter.FeeQuoter, opts *bind.TransactOpts, args ApplyTokenTransferFeeConfigUpdatesArgs) (*types.Transaction, error) {
		return FeeQuoter.ApplyTokenTransferFeeConfigUpdates(opts, args.TokenTransferFeeConfigArgs, args.TokensToUseDefaultFeeConfigs)
	},
})

var UpdatePrices = contract.NewWrite(contract.WriteParams[PriceUpdates, *fee_quoter.FeeQuoter]{
	Name:         "fee-quoter-v2:update-prices",
	Version:      semver.MustParse("1.7.0"),
	Description:  "Updates token prices on the FeeQuoter",
	ContractType: ContractType,
	ContractABI:  fee_quoter.FeeQuoterABI,
	NewContract:  fee_quoter.NewFeeQuoter,
	IsAllowedCaller: func(FeeQuoter *fee_quoter.FeeQuoter, opts *bind.CallOpts, caller common.Address, args PriceUpdates) (bool, error) {
		priceUpdaters, err := FeeQuoter.GetAllAuthorizedCallers(opts)
		if err != nil {
			return false, fmt.Errorf("failed to get authorized callers from FeeQuoter (%s): %w", FeeQuoter.Address(), err)
		}
		if slices.Contains(priceUpdaters, caller) {
			return true, nil
		}
		return false, nil
	},
	Validate: func(PriceUpdates) error { return nil },
	CallContract: func(FeeQuoter *fee_quoter.FeeQuoter, opts *bind.TransactOpts, args PriceUpdates) (*types.Transaction, error) {
		return FeeQuoter.UpdatePrices(opts, args)
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
			},
		})
	}

	return argsTransformed
}
