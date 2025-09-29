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

var Deploy = contract.NewDeploy(
	"fee-quoter-v2:deploy",
	semver.MustParse("1.7.0"),
	"Deploys the FeeQuoterV2 contract",
	ContractType,
	fee_quoter_v2.FeeQuoterV2ABI,
	func(ConstructorArgs) error { return nil },
	contract.VMDeployers[ConstructorArgs]{
		DeployEVM: func(opts *bind.TransactOpts, backend bind.ContractBackend, args ConstructorArgs) (common.Address, *types.Transaction, error) {
			address, tx, _, err := fee_quoter_v2.DeployFeeQuoterV2(
				opts,
				backend,
				args.StaticConfig,
				args.PriceUpdaters,
				args.FeeTokens,
				args.TokenPriceFeeds,
				args.TokenTransferFeeConfigArgs,
				args.PremiumMultiplierWeiPerEthArgs,
				transformDestChainConfigArgs(args.DestChainConfigArgs),
			)
			return address, tx, err
		},
		// DeployZksyncVM: func(opts *accounts.TransactOpts, client *clients.Client, wallet *accounts.Wallet, backend bind.ContractBackend, args ConstructorArgs) (common.Address, error)
	},
)

var ApplyAuthorizedCallerUpdates = contract.NewWrite(
	"fee-quoter-v2:apply-authorized-caller-updates",
	semver.MustParse("1.7.0"),
	"Updates authorized price updaters on the FeeQuoterV2 contract",
	ContractType,
	fee_quoter_v2.FeeQuoterV2ABI,
	fee_quoter_v2.NewFeeQuoterV2,
	contract.OnlyOwner,
	func(AuthorizedCallerArgs) error { return nil },
	func(feeQuoterV2 *fee_quoter_v2.FeeQuoterV2, opts *bind.TransactOpts, args AuthorizedCallerArgs) (*types.Transaction, error) {
		return feeQuoterV2.ApplyAuthorizedCallerUpdates(opts, args)
	},
)

var ApplyDestChainConfigUpdates = contract.NewWrite(
	"fee-quoter-v2:apply-dest-chain-config-updates",
	semver.MustParse("1.7.0"),
	"Applies updates to destination chain configurations on the FeeQuoterV2",
	ContractType,
	fee_quoter_v2.FeeQuoterV2ABI,
	fee_quoter_v2.NewFeeQuoterV2,
	contract.OnlyOwner,
	func([]DestChainConfigArgs) error { return nil },
	func(feeQuoterV2 *fee_quoter_v2.FeeQuoterV2, opts *bind.TransactOpts, args []DestChainConfigArgs) (*types.Transaction, error) {
		return feeQuoterV2.ApplyDestChainConfigUpdates(opts, transformDestChainConfigArgs(args))
	},
)

var ApplyFeeTokensUpdates = contract.NewWrite(
	"fee-quoter-v2:apply-fee-tokens-updates",
	semver.MustParse("1.7.0"),
	"Applies updates to the fee tokens supported by the FeeQuoterV2",
	ContractType,
	fee_quoter_v2.FeeQuoterV2ABI,
	fee_quoter_v2.NewFeeQuoterV2,
	contract.OnlyOwner,
	func(ApplyFeeTokensUpdatesArgs) error { return nil },
	func(feeQuoterV2 *fee_quoter_v2.FeeQuoterV2, opts *bind.TransactOpts, args ApplyFeeTokensUpdatesArgs) (*types.Transaction, error) {
		return feeQuoterV2.ApplyFeeTokensUpdates(opts, args.FeeTokensToRemove, args.FeeTokensToAdd)
	},
)

var ApplyPremiumMultiplierWeiPerEthUpdates = contract.NewWrite(
	"fee-quoter-v2:apply-premium-multiplier-wei-per-eth-updates",
	semver.MustParse("1.7.0"),
	"Applies updates to the premium multiplier (in wei per ETH) for various tokens on the FeeQuoterV2",
	ContractType,
	fee_quoter_v2.FeeQuoterV2ABI,
	fee_quoter_v2.NewFeeQuoterV2,
	contract.OnlyOwner,
	func([]PremiumMultiplierWeiPerEthArgs) error { return nil },
	func(feeQuoterV2 *fee_quoter_v2.FeeQuoterV2, opts *bind.TransactOpts, args []PremiumMultiplierWeiPerEthArgs) (*types.Transaction, error) {
		return feeQuoterV2.ApplyPremiumMultiplierWeiPerEthUpdates(opts, args)
	},
)

var ApplyTokenTransferFeeConfigUpdates = contract.NewWrite(
	"fee-quoter-v2:apply-token-transfer-fee-config-updates",
	semver.MustParse("1.7.0"),
	"Applies updates to the token transfer fee configurations on the FeeQuoterV2",
	ContractType,
	fee_quoter_v2.FeeQuoterV2ABI,
	fee_quoter_v2.NewFeeQuoterV2,
	contract.OnlyOwner,
	func(ApplyTokenTransferFeeConfigUpdatesArgs) error { return nil },
	func(feeQuoterV2 *fee_quoter_v2.FeeQuoterV2, opts *bind.TransactOpts, args ApplyTokenTransferFeeConfigUpdatesArgs) (*types.Transaction, error) {
		return feeQuoterV2.ApplyTokenTransferFeeConfigUpdates(opts, args.TokenTransferFeeConfigArgs, args.TokensToUseDefaultFeeConfigs)
	},
)

var UpdatePrices = contract.NewWrite(
	"fee-quoter-v2:update-prices",
	semver.MustParse("1.7.0"),
	"Updates token prices on the FeeQuoterV2",
	ContractType,
	fee_quoter_v2.FeeQuoterV2ABI,
	fee_quoter_v2.NewFeeQuoterV2,
	func(feeQuoterV2 *fee_quoter_v2.FeeQuoterV2, opts *bind.CallOpts, caller common.Address) (bool, error) {
		priceUpdaters, err := feeQuoterV2.GetAllAuthorizedCallers(opts)
		if err != nil {
			return false, fmt.Errorf("failed to get authorized callers from FeeQuoterV2 (%s): %w", feeQuoterV2.Address(), err)
		}
		if slices.Contains(priceUpdaters, caller) {
			return true, nil
		}
		return false, nil
	},
	func(PriceUpdates) error { return nil },
	func(feeQuoterV2 *fee_quoter_v2.FeeQuoterV2, opts *bind.TransactOpts, args PriceUpdates) (*types.Transaction, error) {
		return feeQuoterV2.UpdatePrices(opts, args)
	},
)

var UpdateTokenPriceFeeds = contract.NewWrite(
	"fee-quoter-v2:update-token-price-feeds",
	semver.MustParse("1.7.0"),
	"Updates the token price feeds on the FeeQuoterV2",
	ContractType,
	fee_quoter_v2.FeeQuoterV2ABI,
	fee_quoter_v2.NewFeeQuoterV2,
	contract.OnlyOwner,
	func([]TokenPriceFeedUpdate) error { return nil },
	func(feeQuoterV2 *fee_quoter_v2.FeeQuoterV2, opts *bind.TransactOpts, args []TokenPriceFeedUpdate) (*types.Transaction, error) {
		return feeQuoterV2.UpdateTokenPriceFeeds(opts, args)
	},
)

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
