package fee_quoter

import (
	"errors"
	"math/big"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_3/fee_quoter"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
)

var (
	ContractType cldf_deployment.ContractType = "FeeQuoter"
	Version      *semver.Version              = semver.MustParse("1.6.3")
)

type ConstructorArgs struct {
	StaticConfig                   fee_quoter.FeeQuoterStaticConfig
	PriceUpdaters                  []common.Address
	FeeTokens                      []common.Address
	TokenPriceFeedUpdates          []fee_quoter.FeeQuoterTokenPriceFeedUpdate
	TokenTransferFeeConfigArgs     []fee_quoter.FeeQuoterTokenTransferFeeConfigArgs
	MorePremiumMultiplierWeiPerEth []fee_quoter.FeeQuoterPremiumMultiplierWeiPerEthArgs
	DestChainConfigArgs            []fee_quoter.FeeQuoterDestChainConfigArgs
}

var Deploy = contract.NewDeploy(contract.DeployParams[ConstructorArgs]{
	Name:             "fee-quoter:deploy",
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

var FeeQuoterApplyDestChainConfigUpdates = contract.NewWrite(contract.WriteParams[[]fee_quoter.FeeQuoterDestChainConfigArgs, *fee_quoter.FeeQuoter]{
	Name:            "fee-quoter:apply-dest-chain-config-updates",
	Version:         Version,
	Description:     "Applies updates to destination chain configs on the FeeQuoter 1.6.0 contract",
	ContractType:    ContractType,
	ContractABI:     fee_quoter.FeeQuoterABI,
	NewContract:     fee_quoter.NewFeeQuoter,
	IsAllowedCaller: contract.OnlyOwner[*fee_quoter.FeeQuoter, []fee_quoter.FeeQuoterDestChainConfigArgs],
	Validate:        func([]fee_quoter.FeeQuoterDestChainConfigArgs) error { return nil },
	CallContract: func(feeQuoter *fee_quoter.FeeQuoter, opts *bind.TransactOpts, args []fee_quoter.FeeQuoterDestChainConfigArgs) (*types.Transaction, error) {
		return feeQuoter.ApplyDestChainConfigUpdates(opts, args)
	},
})

var FeeQuoterUpdatePrices = contract.NewWrite(contract.WriteParams[fee_quoter.InternalPriceUpdates, *fee_quoter.FeeQuoter]{
	Name:            "fee-quoter:update-prices",
	Version:         Version,
	Description:     "Updates prices on the FeeQuoter 1.6.0 contract",
	ContractType:    ContractType,
	ContractABI:     fee_quoter.FeeQuoterABI,
	NewContract:     fee_quoter.NewFeeQuoter,
	IsAllowedCaller: contract.OnlyOwner[*fee_quoter.FeeQuoter, fee_quoter.InternalPriceUpdates],
	Validate:        func(fee_quoter.InternalPriceUpdates) error { return nil },
	CallContract: func(feeQuoter *fee_quoter.FeeQuoter, opts *bind.TransactOpts, args fee_quoter.InternalPriceUpdates) (*types.Transaction, error) {
		return feeQuoter.UpdatePrices(opts, args)
	},
})

type AuthorizedCallerArgs = fee_quoter.AuthorizedCallersAuthorizedCallerArgs

var ApplyAuthorizedCallerUpdates = contract.NewWrite(contract.WriteParams[AuthorizedCallerArgs, *fee_quoter.FeeQuoter]{
	Name:            "fee-quoter:apply-authorized-caller-updates",
	Version:         semver.MustParse("1.6.0"),
	Description:     "Applies updates to the list of authorized callers on the FeeQuoter 1.6.0 contract",
	ContractType:    ContractType,
	ContractABI:     fee_quoter.FeeQuoterABI,
	NewContract:     fee_quoter.NewFeeQuoter,
	IsAllowedCaller: contract.OnlyOwner[*fee_quoter.FeeQuoter, AuthorizedCallerArgs],
	Validate:        func(AuthorizedCallerArgs) error { return nil },
	CallContract: func(feeQuoter *fee_quoter.FeeQuoter, opts *bind.TransactOpts, args AuthorizedCallerArgs) (*types.Transaction, error) {
		return feeQuoter.ApplyAuthorizedCallerUpdates(opts, args)
	},
})

type FeeQuoterParams struct {
	MaxFeeJuelsPerMsg              *big.Int
	TokenPriceStalenessThreshold   uint32
	LinkPremiumMultiplierWeiPerEth uint64
	WethPremiumMultiplierWeiPerEth uint64
	MorePremiumMultiplierWeiPerEth []fee_quoter.FeeQuoterPremiumMultiplierWeiPerEthArgs
	TokenPriceFeedUpdates          []fee_quoter.FeeQuoterTokenPriceFeedUpdate
	TokenTransferFeeConfigArgs     []fee_quoter.FeeQuoterTokenTransferFeeConfigArgs
	DestChainConfigArgs            []fee_quoter.FeeQuoterDestChainConfigArgs
}

func (c FeeQuoterParams) Validate() error {
	if c.MaxFeeJuelsPerMsg == nil {
		return errors.New("MaxFeeJuelsPerMsg is nil")
	}
	if c.MaxFeeJuelsPerMsg.Cmp(big.NewInt(0)) <= 0 {
		return errors.New("MaxFeeJuelsPerMsg must be positive")
	}
	if c.TokenPriceStalenessThreshold == 0 {
		return errors.New("TokenPriceStalenessThreshold can't be 0")
	}
	return nil
}

type ApplyTokenTransferFeeConfigUpdatesInput struct {
	TokenTransferFeeConfigArgs   []fee_quoter.FeeQuoterTokenTransferFeeConfigArgs
	TokensToUseDefaultFeeConfigs []fee_quoter.FeeQuoterTokenTransferFeeConfigRemoveArgs
}

// FeeQuoterApplyTokenTransferFeeConfigUpdates applies updates to token transfer fee configs on the FeeQuoter contract.
// https://etherscan.io/address/0x40858070814a57FdF33a613ae84fE0a8b4a874f7#code#F1#L836
var FeeQuoterApplyTokenTransferFeeConfigUpdates = contract.NewWrite(contract.WriteParams[ApplyTokenTransferFeeConfigUpdatesInput, *fee_quoter.FeeQuoter]{
	Name:            "fee-quoter:apply-token-transfer-fee-config-updates",
	Version:         Version,
	Description:     "Applies updates to token transfer fee configs on the FeeQuoter 1.6.0 contract",
	ContractType:    ContractType,
	ContractABI:     fee_quoter.FeeQuoterABI,
	NewContract:     fee_quoter.NewFeeQuoter,
	IsAllowedCaller: contract.OnlyOwner[*fee_quoter.FeeQuoter, ApplyTokenTransferFeeConfigUpdatesInput],
	Validate:        func(args ApplyTokenTransferFeeConfigUpdatesInput) error { return nil },
	CallContract: func(feeQuoter *fee_quoter.FeeQuoter, opts *bind.TransactOpts, args ApplyTokenTransferFeeConfigUpdatesInput) (*types.Transaction, error) {
		return feeQuoter.ApplyTokenTransferFeeConfigUpdates(opts, args.TokenTransferFeeConfigArgs, args.TokensToUseDefaultFeeConfigs)
	},
})

var GetDestChainConfig = contract.NewRead(contract.ReadParams[uint64, fee_quoter.FeeQuoterDestChainConfig, *fee_quoter.FeeQuoter]{
	Name:         "fee-quoter:dest-chain-config",
	Version:      Version,
	Description:  "Reads the destination chain config from the FeeQuoter 1.6.0 contract",
	ContractType: ContractType,
	NewContract:  fee_quoter.NewFeeQuoter,
	CallContract: func(feeQuoter *fee_quoter.FeeQuoter, opts *bind.CallOpts, chainSelector uint64) (fee_quoter.FeeQuoterDestChainConfig, error) {
		return feeQuoter.GetDestChainConfig(opts, chainSelector)
	},
})

type GetTokenTransferFeeConfigInput struct {
	Token             common.Address
	DestChainSelector uint64
}

var GetTokenTransferFeeConfig = contract.NewRead(contract.ReadParams[GetTokenTransferFeeConfigInput, fee_quoter.FeeQuoterTokenTransferFeeConfig, *fee_quoter.FeeQuoter]{
	Name:         "fee-quoter:token-transfer-fee-config",
	Version:      Version,
	Description:  "Reads the token transfer fee config from the FeeQuoter 1.6.0 contract",
	ContractType: ContractType,
	NewContract:  fee_quoter.NewFeeQuoter,
	CallContract: func(feeQuoter *fee_quoter.FeeQuoter, opts *bind.CallOpts, in GetTokenTransferFeeConfigInput) (fee_quoter.FeeQuoterTokenTransferFeeConfig, error) {
		return feeQuoter.GetTokenTransferFeeConfig(opts, in.DestChainSelector, in.Token)
	},
})

var GetStaticConfig = contract.NewRead(contract.ReadParams[any, fee_quoter.FeeQuoterStaticConfig, *fee_quoter.FeeQuoter]{
	Name:         "fee-quoter:get-static-config",
	Version:      Version,
	Description:  "Gets the static config from the FeeQuoter 1.6.0 contract",
	ContractType: ContractType,
	NewContract:  fee_quoter.NewFeeQuoter,
	CallContract: func(feeQuoter *fee_quoter.FeeQuoter, opts *bind.CallOpts, args any) (fee_quoter.FeeQuoterStaticConfig, error) {
		return feeQuoter.GetStaticConfig(opts)
	},
})

var GetAllAuthorizedCallers = contract.NewRead(contract.ReadParams[any, []common.Address, *fee_quoter.FeeQuoter]{
	Name:         "fee-quoter:get-all-authorized-callers",
	Version:      Version,
	Description:  "Gets all authorized callers (price updaters) from the FeeQuoter 1.6.0 contract",
	ContractType: ContractType,
	NewContract:  fee_quoter.NewFeeQuoter,
	CallContract: func(feeQuoter *fee_quoter.FeeQuoter, opts *bind.CallOpts, args any) ([]common.Address, error) {
		return feeQuoter.GetAllAuthorizedCallers(opts)
	},
})
