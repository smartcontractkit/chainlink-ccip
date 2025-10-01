package fee_quoter_v2

import (
	"fmt"
	"slices"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/latest/fee_quoter_v2"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

var ContractType cldf_deployment.ContractType = "FeeQuoterV2"

type StaticConfig = fee_quoter_v2.FeeQuoterStaticConfig

type DestChainConfig = fee_quoter_v2.FeeQuoterDestChainConfig

type DestChainConfigArgs struct {
	DestChainSelector uint64
	DestChainConfig   DestChainConfig
}

type PremiumMultiplierWeiPerEthArgs = fee_quoter_v2.FeeQuoterPremiumMultiplierWeiPerEthArgs

type TokenTransferFeeConfigArgs = fee_quoter_v2.FeeQuoterTokenTransferFeeConfigArgs

type TokenTransferFeeConfigRemoveArgs = fee_quoter_v2.FeeQuoterTokenTransferFeeConfigRemoveArgs

type TokenPriceFeedUpdate = fee_quoter_v2.FeeQuoterTokenPriceFeedUpdate

type ConstructorArgs struct {
	StaticConfig                   StaticConfig
	PriceUpdaters                  []common.Address
	FeeTokens                      []common.Address
	TokenPriceFeeds                []TokenPriceFeedUpdate
	TokenTransferFeeConfigArgs     []TokenTransferFeeConfigArgs
	PremiumMultiplierWeiPerEthArgs []PremiumMultiplierWeiPerEthArgs
	DestChainConfigArgs            []DestChainConfigArgs
}

type ApplyFeeTokensUpdatesArgs struct {
	FeeTokensToAdd    []common.Address
	FeeTokensToRemove []common.Address
}

type AuthorizedCallerArgs = fee_quoter_v2.AuthorizedCallersAuthorizedCallerArgs

type ApplyTokenTransferFeeConfigUpdatesArgs struct {
	TokenTransferFeeConfigArgs   []TokenTransferFeeConfigArgs
	TokensToUseDefaultFeeConfigs []TokenTransferFeeConfigRemoveArgs
}

type TokenPriceUpdate = fee_quoter_v2.InternalTokenPriceUpdate

type GasPriceUpdate = fee_quoter_v2.InternalGasPriceUpdate

type PriceUpdates = fee_quoter_v2.InternalPriceUpdates

var Deploy = contract.NewDeploy(contract.DeployParams[ConstructorArgs]{
	Name:             "fee-quoter-v2:deploy",
	Version:          semver.MustParse("1.7.0"),
	Description:      "Deploys the FeeQuoterV2 contract",
	ContractType:     ContractType,
	ContractMetadata: fee_quoter_v2.FeeQuoterV2MetaData,
	BytecodeByVersion: map[string]contract.Bytecode{
		semver.MustParse("1.7.0").String(): {EVM: common.FromHex(fee_quoter_v2.FeeQuoterV2Bin)},
	},
	Validate: func(ConstructorArgs) error { return nil },
})

var ApplyAuthorizedCallerUpdates = contract.NewWrite(contract.WriteParams[AuthorizedCallerArgs, *fee_quoter_v2.FeeQuoterV2]{
	Name:            "fee-quoter-v2:apply-authorized-caller-updates",
	Version:         semver.MustParse("1.7.0"),
	Description:     "Updates authorized price updaters on the FeeQuoterV2 contract",
	ContractType:    ContractType,
	ContractABI:     fee_quoter_v2.FeeQuoterV2ABI,
	NewContract:     fee_quoter_v2.NewFeeQuoterV2,
	IsAllowedCaller: contract.OnlyOwner[*fee_quoter_v2.FeeQuoterV2],
	Validate:        func(AuthorizedCallerArgs) error { return nil },
	CallContract: func(feeQuoterV2 *fee_quoter_v2.FeeQuoterV2, opts *bind.TransactOpts, args AuthorizedCallerArgs) (*types.Transaction, error) {
		return feeQuoterV2.ApplyAuthorizedCallerUpdates(opts, args)
	},
})

var ApplyDestChainConfigUpdates = contract.NewWrite(contract.WriteParams[[]DestChainConfigArgs, *fee_quoter_v2.FeeQuoterV2]{
	Name:            "fee-quoter-v2:apply-dest-chain-config-updates",
	Version:         semver.MustParse("1.7.0"),
	Description:     "Applies updates to destination chain configurations on the FeeQuoterV2",
	ContractType:    ContractType,
	ContractABI:     fee_quoter_v2.FeeQuoterV2ABI,
	NewContract:     fee_quoter_v2.NewFeeQuoterV2,
	IsAllowedCaller: contract.OnlyOwner[*fee_quoter_v2.FeeQuoterV2],
	Validate:        func([]DestChainConfigArgs) error { return nil },
	CallContract: func(feeQuoterV2 *fee_quoter_v2.FeeQuoterV2, opts *bind.TransactOpts, args []DestChainConfigArgs) (*types.Transaction, error) {
		return feeQuoterV2.ApplyDestChainConfigUpdates(opts, transformDestChainConfigArgs(args))
	},
})

var ApplyFeeTokensUpdates = contract.NewWrite(contract.WriteParams[ApplyFeeTokensUpdatesArgs, *fee_quoter_v2.FeeQuoterV2]{
	Name:            "fee-quoter-v2:apply-fee-tokens-updates",
	Version:         semver.MustParse("1.7.0"),
	Description:     "Applies updates to the fee tokens supported by the FeeQuoterV2",
	ContractType:    ContractType,
	ContractABI:     fee_quoter_v2.FeeQuoterV2ABI,
	NewContract:     fee_quoter_v2.NewFeeQuoterV2,
	IsAllowedCaller: contract.OnlyOwner[*fee_quoter_v2.FeeQuoterV2],
	Validate:        func(ApplyFeeTokensUpdatesArgs) error { return nil },
	CallContract: func(feeQuoterV2 *fee_quoter_v2.FeeQuoterV2, opts *bind.TransactOpts, args ApplyFeeTokensUpdatesArgs) (*types.Transaction, error) {
		return feeQuoterV2.ApplyFeeTokensUpdates(opts, args.FeeTokensToRemove, args.FeeTokensToAdd)
	},
})

var ApplyPremiumMultiplierWeiPerEthUpdates = contract.NewWrite(contract.WriteParams[[]PremiumMultiplierWeiPerEthArgs, *fee_quoter_v2.FeeQuoterV2]{
	Name:            "fee-quoter-v2:apply-premium-multiplier-wei-per-eth-updates",
	Version:         semver.MustParse("1.7.0"),
	Description:     "Applies updates to the premium multiplier (in wei per ETH) for various tokens on the FeeQuoterV2",
	ContractType:    ContractType,
	ContractABI:     fee_quoter_v2.FeeQuoterV2ABI,
	NewContract:     fee_quoter_v2.NewFeeQuoterV2,
	IsAllowedCaller: contract.OnlyOwner[*fee_quoter_v2.FeeQuoterV2],
	Validate:        func([]PremiumMultiplierWeiPerEthArgs) error { return nil },
	CallContract: func(feeQuoterV2 *fee_quoter_v2.FeeQuoterV2, opts *bind.TransactOpts, args []PremiumMultiplierWeiPerEthArgs) (*types.Transaction, error) {
		return feeQuoterV2.ApplyPremiumMultiplierWeiPerEthUpdates(opts, args)
	},
})

var ApplyTokenTransferFeeConfigUpdates = contract.NewWrite(contract.WriteParams[ApplyTokenTransferFeeConfigUpdatesArgs, *fee_quoter_v2.FeeQuoterV2]{
	Name:            "fee-quoter-v2:apply-token-transfer-fee-config-updates",
	Version:         semver.MustParse("1.7.0"),
	Description:     "Applies updates to the token transfer fee configurations on the FeeQuoterV2",
	ContractType:    ContractType,
	ContractABI:     fee_quoter_v2.FeeQuoterV2ABI,
	NewContract:     fee_quoter_v2.NewFeeQuoterV2,
	IsAllowedCaller: contract.OnlyOwner[*fee_quoter_v2.FeeQuoterV2],
	Validate:        func(ApplyTokenTransferFeeConfigUpdatesArgs) error { return nil },
	CallContract: func(feeQuoterV2 *fee_quoter_v2.FeeQuoterV2, opts *bind.TransactOpts, args ApplyTokenTransferFeeConfigUpdatesArgs) (*types.Transaction, error) {
		return feeQuoterV2.ApplyTokenTransferFeeConfigUpdates(opts, args.TokenTransferFeeConfigArgs, args.TokensToUseDefaultFeeConfigs)
	},
})

var UpdatePrices = contract.NewWrite(contract.WriteParams[PriceUpdates, *fee_quoter_v2.FeeQuoterV2]{
	Name:         "fee-quoter-v2:update-prices",
	Version:      semver.MustParse("1.7.0"),
	Description:  "Updates token prices on the FeeQuoterV2",
	ContractType: ContractType,
	ContractABI:  fee_quoter_v2.FeeQuoterV2ABI,
	NewContract:  fee_quoter_v2.NewFeeQuoterV2,
	IsAllowedCaller: func(feeQuoterV2 *fee_quoter_v2.FeeQuoterV2, opts *bind.CallOpts, caller common.Address) (bool, error) {
		priceUpdaters, err := feeQuoterV2.GetAllAuthorizedCallers(opts)
		if err != nil {
			return false, fmt.Errorf("failed to get authorized callers from FeeQuoterV2 (%s): %w", feeQuoterV2.Address(), err)
		}
		if slices.Contains(priceUpdaters, caller) {
			return true, nil
		}
		return false, nil
	},
	Validate: func(PriceUpdates) error { return nil },
	CallContract: func(feeQuoterV2 *fee_quoter_v2.FeeQuoterV2, opts *bind.TransactOpts, args PriceUpdates) (*types.Transaction, error) {
		return feeQuoterV2.UpdatePrices(opts, args)
	},
})

var UpdateTokenPriceFeeds = contract.NewWrite(contract.WriteParams[[]TokenPriceFeedUpdate, *fee_quoter_v2.FeeQuoterV2]{
	Name:            "fee-quoter-v2:update-token-price-feeds",
	Version:         semver.MustParse("1.7.0"),
	Description:     "Updates the token price feeds on the FeeQuoterV2",
	ContractType:    ContractType,
	ContractABI:     fee_quoter_v2.FeeQuoterV2ABI,
	NewContract:     fee_quoter_v2.NewFeeQuoterV2,
	IsAllowedCaller: contract.OnlyOwner[*fee_quoter_v2.FeeQuoterV2],
	Validate:        func([]TokenPriceFeedUpdate) error { return nil },
	CallContract: func(feeQuoterV2 *fee_quoter_v2.FeeQuoterV2, opts *bind.TransactOpts, args []TokenPriceFeedUpdate) (*types.Transaction, error) {
		return feeQuoterV2.UpdateTokenPriceFeeds(opts, args)
	},
})

func transformDestChainConfigArgs(args []DestChainConfigArgs) []fee_quoter_v2.FeeQuoterDestChainConfigArgs {
	// The reason we avoid exposing fee_quoter_v2.FeeQuoterDestChainConfigArgs directly
	// is so that consumers do not need to import the gobindings package.
	argsTransformed := make([]fee_quoter_v2.FeeQuoterDestChainConfigArgs, 0, len(args))
	for _, arg := range args {
		argsTransformed = append(argsTransformed, fee_quoter_v2.FeeQuoterDestChainConfigArgs{
			DestChainSelector: arg.DestChainSelector,
			DestChainConfig:   arg.DestChainConfig,
		})
	}

	return argsTransformed
}
