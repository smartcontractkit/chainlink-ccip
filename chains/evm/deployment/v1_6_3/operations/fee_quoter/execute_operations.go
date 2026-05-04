package fee_quoter

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	gobindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_3/fee_quoter"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
)

// TokenTransferFeeConfig is the FeeQuoter GetTokenTransferFeeConfig return type (alias for call-site ergonomics).
type TokenTransferFeeConfig = gobindings.FeeQuoterTokenTransferFeeConfig

var ApplyDestChainConfigUpdates = contract.NewWrite(contract.WriteParams[[]gobindings.FeeQuoterDestChainConfigArgs, *gobindings.FeeQuoter]{
	Name:            "fee-quoter:apply-dest-chain-config-updates",
	Version:         Version,
	Description:     "Calls applyDestChainConfigUpdates on the contract",
	ContractType:    ContractType,
	ContractABI:     gobindings.FeeQuoterMetaData.ABI,
	NewContract:     gobindings.NewFeeQuoter,
	IsAllowedCaller: contract.OnlyOwner[*gobindings.FeeQuoter, []gobindings.FeeQuoterDestChainConfigArgs],
	Validate:        func([]gobindings.FeeQuoterDestChainConfigArgs) error { return nil },
	CallContract: func(c *gobindings.FeeQuoter, opts *bind.TransactOpts, args []gobindings.FeeQuoterDestChainConfigArgs) (*types.Transaction, error) {
		return c.ApplyDestChainConfigUpdates(opts, args)
	},
})

var UpdatePrices = contract.NewWrite(contract.WriteParams[gobindings.InternalPriceUpdates, *gobindings.FeeQuoter]{
	Name:            "fee-quoter:update-prices",
	Version:         Version,
	Description:     "Calls updatePrices on the contract",
	ContractType:    ContractType,
	ContractABI:     gobindings.FeeQuoterMetaData.ABI,
	NewContract:     gobindings.NewFeeQuoter,
	IsAllowedCaller: contract.OnlyOwner[*gobindings.FeeQuoter, gobindings.InternalPriceUpdates],
	Validate:        func(gobindings.InternalPriceUpdates) error { return nil },
	CallContract: func(c *gobindings.FeeQuoter, opts *bind.TransactOpts, args gobindings.InternalPriceUpdates) (*types.Transaction, error) {
		return c.UpdatePrices(opts, args)
	},
})

var ApplyAuthorizedCallerUpdates = contract.NewWrite(contract.WriteParams[gobindings.AuthorizedCallersAuthorizedCallerArgs, *gobindings.FeeQuoter]{
	Name:            "fee-quoter:apply-authorized-caller-updates",
	Version:         Version,
	Description:     "Calls applyAuthorizedCallerUpdates on the contract",
	ContractType:    ContractType,
	ContractABI:     gobindings.FeeQuoterMetaData.ABI,
	NewContract:     gobindings.NewFeeQuoter,
	IsAllowedCaller: contract.OnlyOwner[*gobindings.FeeQuoter, gobindings.AuthorizedCallersAuthorizedCallerArgs],
	Validate:        func(gobindings.AuthorizedCallersAuthorizedCallerArgs) error { return nil },
	CallContract: func(c *gobindings.FeeQuoter, opts *bind.TransactOpts, args gobindings.AuthorizedCallersAuthorizedCallerArgs) (*types.Transaction, error) {
		return c.ApplyAuthorizedCallerUpdates(opts, args)
	},
})

var ApplyTokenTransferFeeConfigUpdates = contract.NewWrite(contract.WriteParams[ApplyTokenTransferFeeConfigUpdatesArgs, *gobindings.FeeQuoter]{
	Name:            "fee-quoter:apply-token-transfer-fee-config-updates",
	Version:         Version,
	Description:     "Calls applyTokenTransferFeeConfigUpdates on the contract",
	ContractType:    ContractType,
	ContractABI:     gobindings.FeeQuoterMetaData.ABI,
	NewContract:     gobindings.NewFeeQuoter,
	IsAllowedCaller: contract.OnlyOwner[*gobindings.FeeQuoter, ApplyTokenTransferFeeConfigUpdatesArgs],
	Validate:        func(ApplyTokenTransferFeeConfigUpdatesArgs) error { return nil },
	CallContract: func(c *gobindings.FeeQuoter, opts *bind.TransactOpts, args ApplyTokenTransferFeeConfigUpdatesArgs) (*types.Transaction, error) {
		return c.ApplyTokenTransferFeeConfigUpdates(opts, args.TokenTransferFeeConfigArgs, args.TokensToUseDefaultFeeConfigs)
	},
})

var GetDestChainConfig = contract.NewRead(contract.ReadParams[uint64, gobindings.FeeQuoterDestChainConfig, *gobindings.FeeQuoter]{
	Name:         "fee-quoter:get-dest-chain-config",
	Version:      Version,
	Description:  "Calls getDestChainConfig on the contract",
	ContractType: ContractType,
	NewContract:  gobindings.NewFeeQuoter,
	CallContract: func(c *gobindings.FeeQuoter, opts *bind.CallOpts, args uint64) (gobindings.FeeQuoterDestChainConfig, error) {
		return c.GetDestChainConfig(opts, args)
	},
})

var GetDestinationChainGasPrice = contract.NewRead(contract.ReadParams[uint64, gobindings.InternalTimestampedPackedUint224, *gobindings.FeeQuoter]{
	Name:         "fee-quoter:get-destination-chain-gas-price",
	Version:      Version,
	Description:  "Calls getDestinationChainGasPrice on the contract",
	ContractType: ContractType,
	NewContract:  gobindings.NewFeeQuoter,
	CallContract: func(c *gobindings.FeeQuoter, opts *bind.CallOpts, args uint64) (gobindings.InternalTimestampedPackedUint224, error) {
		return c.GetDestinationChainGasPrice(opts, args)
	},
})

var GetTokenTransferFeeConfig = contract.NewRead(contract.ReadParams[GetTokenTransferFeeConfigArgs, gobindings.FeeQuoterTokenTransferFeeConfig, *gobindings.FeeQuoter]{
	Name:         "fee-quoter:get-token-transfer-fee-config",
	Version:      Version,
	Description:  "Calls getTokenTransferFeeConfig on the contract",
	ContractType: ContractType,
	NewContract:  gobindings.NewFeeQuoter,
	CallContract: func(c *gobindings.FeeQuoter, opts *bind.CallOpts, args GetTokenTransferFeeConfigArgs) (gobindings.FeeQuoterTokenTransferFeeConfig, error) {
		return c.GetTokenTransferFeeConfig(opts, args.DestChainSelector, args.Token)
	},
})

var GetStaticConfig = contract.NewRead(contract.ReadParams[struct{}, gobindings.FeeQuoterStaticConfig, *gobindings.FeeQuoter]{
	Name:         "fee-quoter:get-static-config",
	Version:      Version,
	Description:  "Calls getStaticConfig on the contract",
	ContractType: ContractType,
	NewContract:  gobindings.NewFeeQuoter,
	CallContract: func(c *gobindings.FeeQuoter, opts *bind.CallOpts, args struct{}) (gobindings.FeeQuoterStaticConfig, error) {
		return c.GetStaticConfig(opts)
	},
})

var GetAllAuthorizedCallers = contract.NewRead(contract.ReadParams[struct{}, []common.Address, *gobindings.FeeQuoter]{
	Name:         "fee-quoter:get-all-authorized-callers",
	Version:      Version,
	Description:  "Calls getAllAuthorizedCallers on the contract",
	ContractType: ContractType,
	NewContract:  gobindings.NewFeeQuoter,
	CallContract: func(c *gobindings.FeeQuoter, opts *bind.CallOpts, args struct{}) ([]common.Address, error) {
		return c.GetAllAuthorizedCallers(opts)
	},
})

var GetTokenPrices = contract.NewRead(contract.ReadParams[[]common.Address, []gobindings.InternalTimestampedPackedUint224, *gobindings.FeeQuoter]{
	Name:         "fee-quoter:get-token-prices",
	Version:      Version,
	Description:  "Calls getTokenPrices on the contract",
	ContractType: ContractType,
	NewContract:  gobindings.NewFeeQuoter,
	CallContract: func(c *gobindings.FeeQuoter, opts *bind.CallOpts, args []common.Address) ([]gobindings.InternalTimestampedPackedUint224, error) {
		return c.GetTokenPrices(opts, args)
	},
})

var GetFeeTokens = contract.NewRead(contract.ReadParams[struct{}, []common.Address, *gobindings.FeeQuoter]{
	Name:         "fee-quoter:get-fee-tokens",
	Version:      Version,
	Description:  "Calls getFeeTokens on the contract",
	ContractType: ContractType,
	NewContract:  gobindings.NewFeeQuoter,
	CallContract: func(c *gobindings.FeeQuoter, opts *bind.CallOpts, args struct{}) ([]common.Address, error) {
		return c.GetFeeTokens(opts)
	},
})
