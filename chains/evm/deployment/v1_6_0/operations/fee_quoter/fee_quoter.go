package fee_quoter

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/call"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/deployment"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_0/fee_quoter"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

var ContractType cldf_deployment.ContractType = "FeeQuoter"

type StaticConfig = fee_quoter.FeeQuoterStaticConfig

type DestChainConfig = fee_quoter.FeeQuoterDestChainConfig

type DestChainConfigArgs struct {
	DestChainSelector uint64
	DestChainConfig   DestChainConfig
}

type PremiumMultiplierWeiPerEthArgs = fee_quoter.FeeQuoterPremiumMultiplierWeiPerEthArgs

type TokenTransferFeeConfigArgs = fee_quoter.FeeQuoterTokenTransferFeeConfigArgs

type TokenTransferFeeConfigRemoveArgs = fee_quoter.FeeQuoterTokenTransferFeeConfigRemoveArgs

type TokenPriceFeedUpdate = fee_quoter.FeeQuoterTokenPriceFeedUpdate

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

type AuthorizedCallerArgs = fee_quoter.AuthorizedCallersAuthorizedCallerArgs

type ApplyTokenTransferFeeConfigUpdatesArgs struct {
	TokenTransferFeeConfigArgs   []TokenTransferFeeConfigArgs
	TokensToUseDefaultFeeConfigs []TokenTransferFeeConfigRemoveArgs
}

type InternalPriceUpdates = fee_quoter.InternalPriceUpdates

var Deploy = deployment.New(
	"fee-quoter:deploy",
	semver.MustParse("1.7.0"),
	"Deploys the FeeQuoter contract",
	ContractType,
	fee_quoter.FeeQuoterABI,
	func(ConstructorArgs) error { return nil },
	deployment.VMDeployers[ConstructorArgs]{
		DeployEVM: func(opts *bind.TransactOpts, backend bind.ContractBackend, args ConstructorArgs) (common.Address, *types.Transaction, error) {
			address, tx, _, err := fee_quoter.DeployFeeQuoter(
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

var ApplyAuthorizedCallerUpdates = call.NewWrite(
	"fee-quoter:apply-authorized-caller-updates",
	semver.MustParse("1.7.0"),
	"Updates authorized price updaters on the FeeQuoter contract",
	ContractType,
	fee_quoter.FeeQuoterABI,
	fee_quoter.NewFeeQuoter,
	call.OnlyOwner,
	func(AuthorizedCallerArgs) error { return nil },
	func(feeQuoter *fee_quoter.FeeQuoter, opts *bind.TransactOpts, args AuthorizedCallerArgs) (*types.Transaction, error) {
		return feeQuoter.ApplyAuthorizedCallerUpdates(opts, args)
	},
)

var ApplyDestChainConfigUpdates = call.NewWrite(
	"fee-quoter:apply-dest-chain-config-updates",
	semver.MustParse("1.7.0"),
	"Applies updates to destination chain configurations on the FeeQuoter",
	ContractType,
	fee_quoter.FeeQuoterABI,
	fee_quoter.NewFeeQuoter,
	call.OnlyOwner,
	func([]DestChainConfigArgs) error { return nil },
	func(feeQuoter *fee_quoter.FeeQuoter, opts *bind.TransactOpts, args []DestChainConfigArgs) (*types.Transaction, error) {
		return feeQuoter.ApplyDestChainConfigUpdates(opts, transformDestChainConfigArgs(args))
	},
)

var ApplyFeeTokensUpdates = call.NewWrite(
	"fee-quoter:apply-fee-tokens-updates",
	semver.MustParse("1.7.0"),
	"Applies updates to the fee tokens supported by the FeeQuoter",
	ContractType,
	fee_quoter.FeeQuoterABI,
	fee_quoter.NewFeeQuoter,
	call.OnlyOwner,
	func(ApplyFeeTokensUpdatesArgs) error { return nil },
	func(feeQuoter *fee_quoter.FeeQuoter, opts *bind.TransactOpts, args ApplyFeeTokensUpdatesArgs) (*types.Transaction, error) {
		return feeQuoter.ApplyFeeTokensUpdates(opts, args.FeeTokensToRemove, args.FeeTokensToAdd)
	},
)

var ApplyPremiumMultiplierWeiPerEthUpdates = call.NewWrite(
	"fee-quoter:apply-premium-multiplier-wei-per-eth-updates",
	semver.MustParse("1.7.0"),
	"Applies updates to the premium multiplier (in wei per ETH) for various tokens on the FeeQuoter",
	ContractType,
	fee_quoter.FeeQuoterABI,
	fee_quoter.NewFeeQuoter,
	call.OnlyOwner,
	func([]PremiumMultiplierWeiPerEthArgs) error { return nil },
	func(feeQuoter *fee_quoter.FeeQuoter, opts *bind.TransactOpts, args []PremiumMultiplierWeiPerEthArgs) (*types.Transaction, error) {
		return feeQuoter.ApplyPremiumMultiplierWeiPerEthUpdates(opts, args)
	},
)

var ApplyTokenTransferFeeConfigUpdates = call.NewWrite(
	"fee-quoter:apply-token-transfer-fee-config-updates",
	semver.MustParse("1.7.0"),
	"Applies updates to the token transfer fee configurations on the FeeQuoter",
	ContractType,
	fee_quoter.FeeQuoterABI,
	fee_quoter.NewFeeQuoter,
	call.OnlyOwner,
	func(ApplyTokenTransferFeeConfigUpdatesArgs) error { return nil },
	func(feeQuoter *fee_quoter.FeeQuoter, opts *bind.TransactOpts, args ApplyTokenTransferFeeConfigUpdatesArgs) (*types.Transaction, error) {
		return feeQuoter.ApplyTokenTransferFeeConfigUpdates(opts, args.TokenTransferFeeConfigArgs, args.TokensToUseDefaultFeeConfigs)
	},
)

var UpdatePrices = call.NewWrite(
	"fee-quoter:update-prices",
	semver.MustParse("1.7.0"),
	"Updates token prices on the FeeQuoter",
	ContractType,
	fee_quoter.FeeQuoterABI,
	fee_quoter.NewFeeQuoter,
	func(feeQuoter *fee_quoter.FeeQuoter, opts *bind.CallOpts) ([]common.Address, error) {
		priceUpdaters, err := feeQuoter.GetAllAuthorizedCallers(opts)
		if err != nil {
			return nil, fmt.Errorf("failed to get authorized callers from FeeQuoter (%s): %w", feeQuoter.Address(), err)
		}
		return priceUpdaters, nil
	},
	func(InternalPriceUpdates) error { return nil },
	func(feeQuoter *fee_quoter.FeeQuoter, opts *bind.TransactOpts, args InternalPriceUpdates) (*types.Transaction, error) {
		return feeQuoter.UpdatePrices(opts, args)
	},
)

var UpdateTokenPriceFeeds = call.NewWrite(
	"fee-quoter:update-token-price-feeds",
	semver.MustParse("1.7.0"),
	"Updates the token price feeds on the FeeQuoter",
	ContractType,
	fee_quoter.FeeQuoterABI,
	fee_quoter.NewFeeQuoter,
	call.OnlyOwner,
	func([]TokenPriceFeedUpdate) error { return nil },
	func(feeQuoter *fee_quoter.FeeQuoter, opts *bind.TransactOpts, args []TokenPriceFeedUpdate) (*types.Transaction, error) {
		return feeQuoter.UpdateTokenPriceFeeds(opts, args)
	},
)

func transformDestChainConfigArgs(args []DestChainConfigArgs) []fee_quoter.FeeQuoterDestChainConfigArgs {
	// The reason we avoid exposing fee_quoter.FeeQuoterDestChainConfigArgs directly
	// is so that consumers do not need to import the gobindings package.
	argsTransformed := make([]fee_quoter.FeeQuoterDestChainConfigArgs, 0, len(args))
	for _, arg := range args {
		argsTransformed = append(argsTransformed, fee_quoter.FeeQuoterDestChainConfigArgs{
			DestChainSelector: arg.DestChainSelector,
			DestChainConfig:   arg.DestChainConfig,
		})
	}

	return argsTransformed
}
