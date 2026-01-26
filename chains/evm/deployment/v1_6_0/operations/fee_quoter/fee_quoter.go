package fee_quoter

import (
	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_0/fee_quoter"
)

var ContractType cldf_deployment.ContractType = "FeeQuoter"
var Version = semver.MustParse("1.6.3")

type ConstructorArgs struct {
	StaticConfig                   fee_quoter.FeeQuoterStaticConfig
	PriceUpdaters                  []common.Address
	FeeTokens                      []common.Address
	TokenPriceFeeds                []fee_quoter.FeeQuoterTokenPriceFeedUpdate
	TokenTransferFeeConfigArgs     []fee_quoter.FeeQuoterTokenTransferFeeConfigArgs
	PremiumMultiplierWeiPerEthArgs []fee_quoter.FeeQuoterPremiumMultiplierWeiPerEthArgs
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

var ApplyDestChainConfigUpdates = contract.NewWrite(contract.WriteParams[[]fee_quoter.FeeQuoterDestChainConfigArgs, *fee_quoter.FeeQuoter]{
	Name:            "fee-quoter:apply-dest-chain-config-updates",
	Version:         Version,
	Description:     "Calls applyDestChainConfigUpdates on the contract",
	ContractType:    ContractType,
	ContractABI:     fee_quoter.FeeQuoterABI,
	NewContract:     fee_quoter.NewFeeQuoter,
	IsAllowedCaller: contract.OnlyOwner[*fee_quoter.FeeQuoter, []fee_quoter.FeeQuoterDestChainConfigArgs],
	Validate:        func([]fee_quoter.FeeQuoterDestChainConfigArgs) error { return nil },
	CallContract: func(feeQuoter *fee_quoter.FeeQuoter, opts *bind.TransactOpts, args []fee_quoter.FeeQuoterDestChainConfigArgs) (*types.Transaction, error) {
		return feeQuoter.ApplyDestChainConfigUpdates(opts, args)
	},
})

var UpdatePrices = contract.NewWrite(contract.WriteParams[fee_quoter.InternalPriceUpdates, *fee_quoter.FeeQuoter]{
	Name:            "fee-quoter:update-prices",
	Version:         Version,
	Description:     "Calls updatePrices on the contract",
	ContractType:    ContractType,
	ContractABI:     fee_quoter.FeeQuoterABI,
	NewContract:     fee_quoter.NewFeeQuoter,
	IsAllowedCaller: contract.AllCallersAllowed[*fee_quoter.FeeQuoter, fee_quoter.InternalPriceUpdates],
	Validate:        func(fee_quoter.InternalPriceUpdates) error { return nil },
	CallContract: func(feeQuoter *fee_quoter.FeeQuoter, opts *bind.TransactOpts, args fee_quoter.InternalPriceUpdates) (*types.Transaction, error) {
		return feeQuoter.UpdatePrices(opts, args)
	},
})

var ApplyAuthorizedCallerUpdates = contract.NewWrite(contract.WriteParams[fee_quoter.AuthorizedCallersAuthorizedCallerArgs, *fee_quoter.FeeQuoter]{
	Name:            "fee-quoter:apply-authorized-caller-updates",
	Version:         Version,
	Description:     "Calls applyAuthorizedCallerUpdates on the contract",
	ContractType:    ContractType,
	ContractABI:     fee_quoter.FeeQuoterABI,
	NewContract:     fee_quoter.NewFeeQuoter,
	IsAllowedCaller: contract.AllCallersAllowed[*fee_quoter.FeeQuoter, fee_quoter.AuthorizedCallersAuthorizedCallerArgs],
	Validate:        func(fee_quoter.AuthorizedCallersAuthorizedCallerArgs) error { return nil },
	CallContract: func(feeQuoter *fee_quoter.FeeQuoter, opts *bind.TransactOpts, args fee_quoter.AuthorizedCallersAuthorizedCallerArgs) (*types.Transaction, error) {
		return feeQuoter.ApplyAuthorizedCallerUpdates(opts, args)
	},
})

type ApplyTokenTransferFeeConfigUpdatesArgs struct {
	TokenTransferFeeConfigArgs   []fee_quoter.FeeQuoterTokenTransferFeeConfigArgs
	TokensToUseDefaultFeeConfigs []fee_quoter.FeeQuoterTokenTransferFeeConfigRemoveArgs
}

var ApplyTokenTransferFeeConfigUpdates = contract.NewWrite(contract.WriteParams[ApplyTokenTransferFeeConfigUpdatesArgs, *fee_quoter.FeeQuoter]{
	Name:            "fee-quoter:apply-token-transfer-fee-config-updates",
	Version:         Version,
	Description:     "Calls applyTokenTransferFeeConfigUpdates on the contract",
	ContractType:    ContractType,
	ContractABI:     fee_quoter.FeeQuoterABI,
	NewContract:     fee_quoter.NewFeeQuoter,
	IsAllowedCaller: contract.OnlyOwner[*fee_quoter.FeeQuoter, ApplyTokenTransferFeeConfigUpdatesArgs],
	Validate:        func(ApplyTokenTransferFeeConfigUpdatesArgs) error { return nil },
	CallContract: func(feeQuoter *fee_quoter.FeeQuoter, opts *bind.TransactOpts, args ApplyTokenTransferFeeConfigUpdatesArgs) (*types.Transaction, error) {
		return feeQuoter.ApplyTokenTransferFeeConfigUpdates(opts, args.TokenTransferFeeConfigArgs, args.TokensToUseDefaultFeeConfigs)
	},
})
