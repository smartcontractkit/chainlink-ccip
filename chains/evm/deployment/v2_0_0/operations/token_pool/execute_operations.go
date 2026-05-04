package token_pool

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	gobindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v2_0_0/token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
)

// TokenPoolABI is the JSON ABI of the TokenPool contract.
var TokenPoolABI = gobindings.TokenPoolMetaData.ABI

// Type aliases for sequence and helper call sites.
type (
	TokenBucket                = gobindings.RateLimiterTokenBucket
	Config                     = gobindings.RateLimiterConfig
	RateLimitConfigArgs        = gobindings.TokenPoolRateLimitConfigArgs
	TokenTransferFeeConfigArgs = gobindings.TokenPoolTokenTransferFeeConfigArgs
	TokenTransferFeeConfig     = gobindings.IPoolV2TokenTransferFeeConfig
	ChainUpdate                = gobindings.TokenPoolChainUpdate
)

var SetAllowedFinalityConfig = contract.NewWrite(contract.WriteParams[[4]byte, *gobindings.TokenPool]{
	Name:            "token-pool:set-allowed-finality-config",
	Version:         Version,
	Description:     "Calls setAllowedFinalityConfig on the contract",
	ContractType:    ContractType,
	ContractABI:     gobindings.TokenPoolMetaData.ABI,
	NewContract:     gobindings.NewTokenPool,
	IsAllowedCaller: contract.OnlyOwner[*gobindings.TokenPool, [4]byte],
	Validate:        func([4]byte) error { return nil },
	CallContract: func(c *gobindings.TokenPool, opts *bind.TransactOpts, args [4]byte) (*types.Transaction, error) {
		return c.SetAllowedFinalityConfig(opts, args)
	},
})

var AddRemotePool = contract.NewWrite(contract.WriteParams[AddRemotePoolArgs, *gobindings.TokenPool]{
	Name:            "token-pool:add-remote-pool",
	Version:         Version,
	Description:     "Calls addRemotePool on the contract",
	ContractType:    ContractType,
	ContractABI:     gobindings.TokenPoolMetaData.ABI,
	NewContract:     gobindings.NewTokenPool,
	IsAllowedCaller: contract.OnlyOwner[*gobindings.TokenPool, AddRemotePoolArgs],
	Validate:        func(AddRemotePoolArgs) error { return nil },
	CallContract: func(c *gobindings.TokenPool, opts *bind.TransactOpts, args AddRemotePoolArgs) (*types.Transaction, error) {
		return c.AddRemotePool(opts, args.RemoteChainSelector, args.RemotePoolAddress)
	},
})

var SetDynamicConfig = contract.NewWrite(contract.WriteParams[SetDynamicConfigArgs, *gobindings.TokenPool]{
	Name:            "token-pool:set-dynamic-config",
	Version:         Version,
	Description:     "Calls setDynamicConfig on the contract",
	ContractType:    ContractType,
	ContractABI:     gobindings.TokenPoolMetaData.ABI,
	NewContract:     gobindings.NewTokenPool,
	IsAllowedCaller: contract.OnlyOwner[*gobindings.TokenPool, SetDynamicConfigArgs],
	Validate:        func(SetDynamicConfigArgs) error { return nil },
	CallContract: func(c *gobindings.TokenPool, opts *bind.TransactOpts, args SetDynamicConfigArgs) (*types.Transaction, error) {
		return c.SetDynamicConfig(opts, args.Router, args.RateLimitAdmin, args.FeeAdmin)
	},
})

var ApplyChainUpdates = contract.NewWrite(contract.WriteParams[ApplyChainUpdatesArgs, *gobindings.TokenPool]{
	Name:            "token-pool:apply-chain-updates",
	Version:         Version,
	Description:     "Calls applyChainUpdates on the contract",
	ContractType:    ContractType,
	ContractABI:     gobindings.TokenPoolMetaData.ABI,
	NewContract:     gobindings.NewTokenPool,
	IsAllowedCaller: contract.OnlyOwner[*gobindings.TokenPool, ApplyChainUpdatesArgs],
	Validate:        func(ApplyChainUpdatesArgs) error { return nil },
	CallContract: func(c *gobindings.TokenPool, opts *bind.TransactOpts, args ApplyChainUpdatesArgs) (*types.Transaction, error) {
		return c.ApplyChainUpdates(opts, args.RemoteChainSelectorsToRemove, args.ChainsToAdd)
	},
})

var SetRateLimitConfig = contract.NewWrite(contract.WriteParams[[]gobindings.TokenPoolRateLimitConfigArgs, *gobindings.TokenPool]{
	Name:            "token-pool:set-rate-limit-config",
	Version:         Version,
	Description:     "Calls setRateLimitConfig on the contract",
	ContractType:    ContractType,
	ContractABI:     gobindings.TokenPoolMetaData.ABI,
	NewContract:     gobindings.NewTokenPool,
	IsAllowedCaller: contract.OnlyOwner[*gobindings.TokenPool, []gobindings.TokenPoolRateLimitConfigArgs],
	Validate:        func([]gobindings.TokenPoolRateLimitConfigArgs) error { return nil },
	CallContract: func(c *gobindings.TokenPool, opts *bind.TransactOpts, args []gobindings.TokenPoolRateLimitConfigArgs) (*types.Transaction, error) {
		return c.SetRateLimitConfig(opts, args)
	},
})

var ApplyTokenTransferFeeConfigUpdates = contract.NewWrite(contract.WriteParams[ApplyTokenTransferFeeConfigUpdatesArgs, *gobindings.TokenPool]{
	Name:            "token-pool:apply-token-transfer-fee-config-updates",
	Version:         Version,
	Description:     "Calls applyTokenTransferFeeConfigUpdates on the contract",
	ContractType:    ContractType,
	ContractABI:     gobindings.TokenPoolMetaData.ABI,
	NewContract:     gobindings.NewTokenPool,
	IsAllowedCaller: contract.OnlyOwner[*gobindings.TokenPool, ApplyTokenTransferFeeConfigUpdatesArgs],
	Validate:        func(ApplyTokenTransferFeeConfigUpdatesArgs) error { return nil },
	CallContract: func(c *gobindings.TokenPool, opts *bind.TransactOpts, args ApplyTokenTransferFeeConfigUpdatesArgs) (*types.Transaction, error) {
		return c.ApplyTokenTransferFeeConfigUpdates(opts, args.TokenTransferFeeConfigArgs, args.DisableTokenTransferFeeConfigs)
	},
})

var GetDynamicConfig = contract.NewRead(contract.ReadParams[struct{}, GetDynamicConfigResult, *gobindings.TokenPool]{
	Name:         "token-pool:get-dynamic-config",
	Version:      Version,
	Description:  "Calls getDynamicConfig on the contract",
	ContractType: ContractType,
	NewContract:  gobindings.NewTokenPool,
	CallContract: func(c *gobindings.TokenPool, opts *bind.CallOpts, args struct{}) (GetDynamicConfigResult, error) {
		res, err := c.GetDynamicConfig(opts)
		if err != nil {
			return GetDynamicConfigResult{}, err
		}
		return GetDynamicConfigResult{Router: res.Router, RateLimitAdmin: res.RateLimitAdmin, FeeAdmin: res.FeeAdmin}, nil
	},
})

var GetRmnProxy = contract.NewRead(contract.ReadParams[struct{}, common.Address, *gobindings.TokenPool]{
	Name:         "token-pool:get-rmn-proxy",
	Version:      Version,
	Description:  "Calls getRmnProxy on the contract",
	ContractType: ContractType,
	NewContract:  gobindings.NewTokenPool,
	CallContract: func(c *gobindings.TokenPool, opts *bind.CallOpts, args struct{}) (common.Address, error) {
		return c.GetRmnProxy(opts)
	},
})

var GetCurrentRateLimiterState = contract.NewRead(contract.ReadParams[GetCurrentRateLimiterStateArgs, GetCurrentRateLimiterStateResult, *gobindings.TokenPool]{
	Name:         "token-pool:get-current-rate-limiter-state",
	Version:      Version,
	Description:  "Calls getCurrentRateLimiterState on the contract",
	ContractType: ContractType,
	NewContract:  gobindings.NewTokenPool,
	CallContract: func(c *gobindings.TokenPool, opts *bind.CallOpts, args GetCurrentRateLimiterStateArgs) (GetCurrentRateLimiterStateResult, error) {
		res, err := c.GetCurrentRateLimiterState(opts, args.RemoteChainSelector, args.FastFinality)
		if err != nil {
			return GetCurrentRateLimiterStateResult{}, err
		}
		return GetCurrentRateLimiterStateResult{
			OutboundRateLimiterState: res.OutboundRateLimiterState,
			InboundRateLimiterState:  res.InboundRateLimiterState,
		}, nil
	},
})

var GetAllowedFinalityConfig = contract.NewRead(contract.ReadParams[struct{}, [4]byte, *gobindings.TokenPool]{
	Name:         "token-pool:get-allowed-finality-config",
	Version:      Version,
	Description:  "Calls getAllowedFinalityConfig on the contract",
	ContractType: ContractType,
	NewContract:  gobindings.NewTokenPool,
	CallContract: func(c *gobindings.TokenPool, opts *bind.CallOpts, args struct{}) ([4]byte, error) {
		return c.GetAllowedFinalityConfig(opts)
	},
})

var GetSupportedChains = contract.NewRead(contract.ReadParams[struct{}, []uint64, *gobindings.TokenPool]{
	Name:         "token-pool:get-supported-chains",
	Version:      Version,
	Description:  "Calls getSupportedChains on the contract",
	ContractType: ContractType,
	NewContract:  gobindings.NewTokenPool,
	CallContract: func(c *gobindings.TokenPool, opts *bind.CallOpts, args struct{}) ([]uint64, error) {
		return c.GetSupportedChains(opts)
	},
})

var GetRemotePools = contract.NewRead(contract.ReadParams[uint64, [][]byte, *gobindings.TokenPool]{
	Name:         "token-pool:get-remote-pools",
	Version:      Version,
	Description:  "Calls getRemotePools on the contract",
	ContractType: ContractType,
	NewContract:  gobindings.NewTokenPool,
	CallContract: func(c *gobindings.TokenPool, opts *bind.CallOpts, args uint64) ([][]byte, error) {
		return c.GetRemotePools(opts, args)
	},
})

var GetRemoteToken = contract.NewRead(contract.ReadParams[uint64, []byte, *gobindings.TokenPool]{
	Name:         "token-pool:get-remote-token",
	Version:      Version,
	Description:  "Calls getRemoteToken on the contract",
	ContractType: ContractType,
	NewContract:  gobindings.NewTokenPool,
	CallContract: func(c *gobindings.TokenPool, opts *bind.CallOpts, args uint64) ([]byte, error) {
		return c.GetRemoteToken(opts, args)
	},
})

var GetToken = contract.NewRead(contract.ReadParams[struct{}, common.Address, *gobindings.TokenPool]{
	Name:         "token-pool:get-token",
	Version:      Version,
	Description:  "Calls getToken on the contract",
	ContractType: ContractType,
	NewContract:  gobindings.NewTokenPool,
	CallContract: func(c *gobindings.TokenPool, opts *bind.CallOpts, args struct{}) (common.Address, error) {
		return c.GetToken(opts)
	},
})

var GetTokenDecimals = contract.NewRead(contract.ReadParams[struct{}, uint8, *gobindings.TokenPool]{
	Name:         "token-pool:get-token-decimals",
	Version:      Version,
	Description:  "Calls getTokenDecimals on the contract",
	ContractType: ContractType,
	NewContract:  gobindings.NewTokenPool,
	CallContract: func(c *gobindings.TokenPool, opts *bind.CallOpts, args struct{}) (uint8, error) {
		return c.GetTokenDecimals(opts)
	},
})

var IsSupportedToken = contract.NewRead(contract.ReadParams[common.Address, bool, *gobindings.TokenPool]{
	Name:         "token-pool:is-supported-token",
	Version:      Version,
	Description:  "Calls isSupportedToken on the contract",
	ContractType: ContractType,
	NewContract:  gobindings.NewTokenPool,
	CallContract: func(c *gobindings.TokenPool, opts *bind.CallOpts, args common.Address) (bool, error) {
		return c.IsSupportedToken(opts, args)
	},
})

var GetAdvancedPoolHooks = contract.NewRead(contract.ReadParams[struct{}, common.Address, *gobindings.TokenPool]{
	Name:         "token-pool:get-advanced-pool-hooks",
	Version:      Version,
	Description:  "Calls getAdvancedPoolHooks on the contract",
	ContractType: ContractType,
	NewContract:  gobindings.NewTokenPool,
	CallContract: func(c *gobindings.TokenPool, opts *bind.CallOpts, args struct{}) (common.Address, error) {
		return c.GetAdvancedPoolHooks(opts)
	},
})

var GetTokenTransferFeeConfig = contract.NewRead(contract.ReadParams[GetTokenTransferFeeConfigArgs, gobindings.IPoolV2TokenTransferFeeConfig, *gobindings.TokenPool]{
	Name:         "token-pool:get-token-transfer-fee-config",
	Version:      Version,
	Description:  "Calls getTokenTransferFeeConfig on the contract",
	ContractType: ContractType,
	NewContract:  gobindings.NewTokenPool,
	CallContract: func(c *gobindings.TokenPool, opts *bind.CallOpts, args GetTokenTransferFeeConfigArgs) (gobindings.IPoolV2TokenTransferFeeConfig, error) {
		return c.GetTokenTransferFeeConfig(opts, args.Arg0, args.DestChainSelector, args.Arg2, args.Arg3)
	},
})
