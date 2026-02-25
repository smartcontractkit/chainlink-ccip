package token_pool

import (
	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/gobindings/generated/latest/token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

var ContractType cldf_deployment.ContractType = "TokenPool"

var Version = semver.MustParse("1.7.0")

type ChainUpdate struct {
	RemoteChainSelector       uint64
	RemotePoolAddresses       [][]byte
	RemoteTokenAddress        []byte
	OutboundRateLimiterConfig tokens.RateLimiterConfig
	InboundRateLimiterConfig  tokens.RateLimiterConfig
}

type CCVConfigArg = token_pool.TokenPool

type ApplyChainUpdatesArgs struct {
	RemoteChainSelectorsToRemove []uint64
	ChainsToAdd                  []ChainUpdate
}

type RemotePoolArgs struct {
	RemoteChainSelector uint64
	RemotePoolAddress   []byte
}

type DynamicConfig = token_pool.GetDynamicConfig

type RateLimiterStates = token_pool.GetCurrentRateLimiterState

type SetRateLimitConfigArg struct {
	RemoteChainSelector       uint64
	CustomBlockConfirmation   bool
	OutboundRateLimiterConfig tokens.RateLimiterConfig
	InboundRateLimiterConfig  tokens.RateLimiterConfig
}

type CustomBlockConfirmationRateLimiterStates = token_pool.GetCurrentRateLimiterState

type SetChainRateLimiterConfigArg struct {
	RemoteChainSelector       uint64
	InboundRateLimiterConfig  tokens.RateLimiterConfig
	OutboundRateLimiterConfig tokens.RateLimiterConfig
}

type ApplyAllowListUpdatesArgs struct {
	Adds    []common.Address
	Removes []common.Address
}

type DynamicConfigArgs struct {
	Router         common.Address
	RateLimitAdmin common.Address
	FeeAdmin       common.Address
}

type CustomBlockConfirmationRateLimitConfigArg struct {
	RemoteChainSelector       uint64
	OutboundRateLimiterConfig tokens.RateLimiterConfig
	InboundRateLimiterConfig  tokens.RateLimiterConfig
}

type ApplyCustomBlockConfirmationConfigArgs struct {
	MinBlockConfirmation uint16
	RateLimitConfigArgs  []CustomBlockConfirmationRateLimitConfigArg
}

type GetCurrentRateLimiterStateArgs struct {
	RemoteChainSelector     uint64
	CustomBlockConfirmation bool
}

type WithdrawFeeTokensArgs struct {
	FeeTokens []common.Address
	Recipient common.Address
}

type TokenTransferFeeConfigUpdate struct {
	DestChainSelector                      uint64
	DestGasOverhead                        uint32
	DestBytesOverhead                      uint32
	DefaultBlockConfirmationFeeUSDCents    uint32
	CustomBlockConfirmationFeeUSDCents     uint32
	DefaultBlockConfirmationTransferFeeBps uint16
	CustomBlockConfirmationTransferFeeBps  uint16
	IsEnabled                              bool
}

type TokenTransferFeeConfigArgs struct {
	DestChainSelectorsToRemove    []uint64
	TokenTransferFeeConfigUpdates []TokenTransferFeeConfigUpdate
}

var SetMinBlockConfirmation = contract.NewWrite(contract.WriteParams[uint16, *token_pool.TokenPool]{
	Name:            "token-pool:set-min-block-confirmation",
	Version:         Version,
	Description:     "Sets the minimum block confirmation required for a TokenPool",
	ContractType:    ContractType,
	ContractABI:     token_pool.TokenPoolABI,
	NewContract:     token_pool.NewTokenPool,
	IsAllowedCaller: contract.OnlyOwner[*token_pool.TokenPool, uint16],
	Validate:        func(uint16) error { return nil },
	CallContract: func(tokenPool *token_pool.TokenPool, opts *bind.TransactOpts, args uint16) (*types.Transaction, error) {
		return tokenPool.SetMinBlockConfirmations(opts, args)
	},
})

var ApplyChainUpdates = contract.NewWrite(contract.WriteParams[ApplyChainUpdatesArgs, *token_pool.TokenPool]{
	Name:            "token-pool:apply-chain-updates",
	Version:         Version,
	Description:     "Applies chain updates to a TokenPool, enabling / disabling remote chains and setting rate limits",
	ContractType:    ContractType,
	ContractABI:     token_pool.TokenPoolABI,
	NewContract:     token_pool.NewTokenPool,
	IsAllowedCaller: contract.OnlyOwner[*token_pool.TokenPool, ApplyChainUpdatesArgs],
	Validate:        func(ApplyChainUpdatesArgs) error { return nil },
	CallContract: func(tokenPool *token_pool.TokenPool, opts *bind.TransactOpts, args ApplyChainUpdatesArgs) (*types.Transaction, error) {
		chainsToAdd := make([]token_pool.TokenPoolChainUpdate, len(args.ChainsToAdd))
		for i, chain := range args.ChainsToAdd {
			chainsToAdd[i] = token_pool.TokenPoolChainUpdate{
				RemoteChainSelector:       chain.RemoteChainSelector,
				RemotePoolAddresses:       chain.RemotePoolAddresses,
				RemoteTokenAddress:        chain.RemoteTokenAddress,
				OutboundRateLimiterConfig: token_pool.RateLimiterConfig(chain.OutboundRateLimiterConfig),
				InboundRateLimiterConfig:  token_pool.RateLimiterConfig(chain.InboundRateLimiterConfig),
			}
		}
		return tokenPool.ApplyChainUpdates(opts, args.RemoteChainSelectorsToRemove, chainsToAdd)
	},
})

var AddRemotePool = contract.NewWrite(contract.WriteParams[RemotePoolArgs, *token_pool.TokenPool]{
	Name:            "token-pool:add-remote-pool",
	Version:         Version,
	Description:     "Adds a remote pool for a given chain selector to a TokenPool",
	ContractType:    ContractType,
	ContractABI:     token_pool.TokenPoolABI,
	NewContract:     token_pool.NewTokenPool,
	IsAllowedCaller: contract.OnlyOwner[*token_pool.TokenPool, RemotePoolArgs],
	Validate:        func(RemotePoolArgs) error { return nil },
	CallContract: func(tokenPool *token_pool.TokenPool, opts *bind.TransactOpts, args RemotePoolArgs) (*types.Transaction, error) {
		return tokenPool.AddRemotePool(opts, args.RemoteChainSelector, args.RemotePoolAddress)
	},
})

var RemoveRemotePool = contract.NewWrite(contract.WriteParams[RemotePoolArgs, *token_pool.TokenPool]{
	Name:            "token-pool:remove-remote-pool",
	Version:         Version,
	Description:     "Removes a remote pool for a given chain selector from a TokenPool",
	ContractType:    ContractType,
	ContractABI:     token_pool.TokenPoolABI,
	NewContract:     token_pool.NewTokenPool,
	IsAllowedCaller: contract.OnlyOwner[*token_pool.TokenPool, RemotePoolArgs],
	Validate:        func(RemotePoolArgs) error { return nil },
	CallContract: func(tokenPool *token_pool.TokenPool, opts *bind.TransactOpts, args RemotePoolArgs) (*types.Transaction, error) {
		return tokenPool.RemoveRemotePool(opts, args.RemoteChainSelector, args.RemotePoolAddress)
	},
})

var SetRateLimitConfig = contract.NewWrite(contract.WriteParams[[]SetRateLimitConfigArg, *token_pool.TokenPool]{
	Name:            "token-pool:set-rate-limit-config",
	Version:         Version,
	Description:     "Sets the rate limit configs for existing remote chains on a TokenPool",
	ContractType:    ContractType,
	ContractABI:     token_pool.TokenPoolABI,
	NewContract:     token_pool.NewTokenPool,
	IsAllowedCaller: contract.OnlyOwner[*token_pool.TokenPool, []SetRateLimitConfigArg],
	Validate:        func([]SetRateLimitConfigArg) error { return nil },
	CallContract: func(tokenPool *token_pool.TokenPool, opts *bind.TransactOpts, args []SetRateLimitConfigArg) (*types.Transaction, error) {
		rateLimitConfigArgs := make([]token_pool.TokenPoolRateLimitConfigArgs, 0, len(args))
		for _, arg := range args {
			rateLimitConfigArgs = append(rateLimitConfigArgs, token_pool.TokenPoolRateLimitConfigArgs{
				RemoteChainSelector:       arg.RemoteChainSelector,
				CustomBlockConfirmations:  arg.CustomBlockConfirmation,
				OutboundRateLimiterConfig: token_pool.RateLimiterConfig(arg.OutboundRateLimiterConfig),
				InboundRateLimiterConfig:  token_pool.RateLimiterConfig(arg.InboundRateLimiterConfig),
			})
		}
		return tokenPool.SetRateLimitConfig(opts, rateLimitConfigArgs)
	},
})

var SetDynamicConfig = contract.NewWrite(contract.WriteParams[DynamicConfigArgs, *token_pool.TokenPool]{
	Name:            "token-pool:set-dynamic-config",
	Version:         Version,
	Description:     "Sets the router, rate limit admin and fee aggregator for a TokenPool",
	ContractType:    ContractType,
	ContractABI:     token_pool.TokenPoolABI,
	NewContract:     token_pool.NewTokenPool,
	IsAllowedCaller: contract.OnlyOwner[*token_pool.TokenPool, DynamicConfigArgs],
	Validate:        func(DynamicConfigArgs) error { return nil },
	CallContract: func(tokenPool *token_pool.TokenPool, opts *bind.TransactOpts, args DynamicConfigArgs) (*types.Transaction, error) {
		return tokenPool.SetDynamicConfig(opts, args.Router, args.RateLimitAdmin, args.FeeAdmin)
	},
})

var WithdrawFeeTokens = contract.NewWrite(contract.WriteParams[WithdrawFeeTokensArgs, *token_pool.TokenPool]{
	Name:            "token-pool:withdraw-fee-tokens",
	Version:         Version,
	Description:     "Withdraws fee tokens to a recipient from a TokenPool",
	ContractType:    ContractType,
	ContractABI:     token_pool.TokenPoolABI,
	NewContract:     token_pool.NewTokenPool,
	IsAllowedCaller: contract.AllCallersAllowed[*token_pool.TokenPool, WithdrawFeeTokensArgs],
	Validate:        func(WithdrawFeeTokensArgs) error { return nil },
	CallContract: func(tokenPool *token_pool.TokenPool, opts *bind.TransactOpts, args WithdrawFeeTokensArgs) (*types.Transaction, error) {
		return tokenPool.WithdrawFeeTokens(opts, args.FeeTokens, args.Recipient)
	},
})

var UpdateAdvancedPoolHooks = contract.NewWrite(contract.WriteParams[common.Address, *token_pool.TokenPool]{
	Name:            "token-pool:update-advanced-pool-hooks",
	Version:         Version,
	Description:     "Updates the advanced pool hooks address on a TokenPool",
	ContractType:    ContractType,
	ContractABI:     token_pool.TokenPoolABI,
	NewContract:     token_pool.NewTokenPool,
	IsAllowedCaller: contract.OnlyOwner[*token_pool.TokenPool, common.Address],
	Validate:        func(common.Address) error { return nil },
	CallContract: func(tokenPool *token_pool.TokenPool, opts *bind.TransactOpts, args common.Address) (*types.Transaction, error) {
		return tokenPool.UpdateAdvancedPoolHooks(opts, args)
	},
})

var ApplyTokenTransferFeeConfigUpdates = contract.NewWrite(contract.WriteParams[TokenTransferFeeConfigArgs, *token_pool.TokenPool]{
	Name:            "token-pool:apply-token-transfer-fee-config-updates",
	Version:         Version,
	Description:     "Applies token transfer fee config updates to a TokenPool",
	ContractType:    ContractType,
	ContractABI:     token_pool.TokenPoolABI,
	NewContract:     token_pool.NewTokenPool,
	IsAllowedCaller: contract.OnlyOwner[*token_pool.TokenPool, TokenTransferFeeConfigArgs],
	Validate:        func(TokenTransferFeeConfigArgs) error { return nil },
	CallContract: func(tokenPool *token_pool.TokenPool, opts *bind.TransactOpts, args TokenTransferFeeConfigArgs) (*types.Transaction, error) {
		tokenTransferFeeConfigUpdates := make([]token_pool.TokenPoolTokenTransferFeeConfigArgs, 0, len(args.TokenTransferFeeConfigUpdates))
		for _, arg := range args.TokenTransferFeeConfigUpdates {
			tokenTransferFeeConfigUpdates = append(tokenTransferFeeConfigUpdates, token_pool.TokenPoolTokenTransferFeeConfigArgs{
				DestChainSelector: arg.DestChainSelector,
				TokenTransferFeeConfig: token_pool.IPoolV2TokenTransferFeeConfig{
					DestGasOverhead:                         arg.DestGasOverhead,
					DestBytesOverhead:                       arg.DestBytesOverhead,
					DefaultBlockConfirmationsFeeUSDCents:    arg.DefaultBlockConfirmationFeeUSDCents,
					CustomBlockConfirmationsFeeUSDCents:     arg.CustomBlockConfirmationFeeUSDCents,
					DefaultBlockConfirmationsTransferFeeBps: arg.DefaultBlockConfirmationTransferFeeBps,
					CustomBlockConfirmationsTransferFeeBps:  arg.CustomBlockConfirmationTransferFeeBps,
					IsEnabled:                               arg.IsEnabled,
				},
			})
		}
		return tokenPool.ApplyTokenTransferFeeConfigUpdates(opts, tokenTransferFeeConfigUpdates, args.DestChainSelectorsToRemove)
	},
})

var GetDynamicConfig = contract.NewRead(contract.ReadParams[any, DynamicConfig, *token_pool.TokenPool]{
	Name:         "token-pool:get-dynamic-config",
	Version:      Version,
	Description:  "Gets the router and rate limit admin configuration on a TokenPool",
	ContractType: ContractType,
	NewContract:  token_pool.NewTokenPool,
	CallContract: func(tokenPool *token_pool.TokenPool, opts *bind.CallOpts, args any) (DynamicConfig, error) {
		return tokenPool.GetDynamicConfig(opts)
	},
})

var GetRMNProxy = contract.NewRead(contract.ReadParams[any, common.Address, *token_pool.TokenPool]{
	Name:         "token-pool:get-rmn-proxy",
	Version:      Version,
	Description:  "Gets the RMN proxy address on a TokenPool",
	ContractType: ContractType,
	NewContract:  token_pool.NewTokenPool,
	CallContract: func(tokenPool *token_pool.TokenPool, opts *bind.CallOpts, args any) (common.Address, error) {
		return tokenPool.GetRmnProxy(opts)
	},
})

var GetCurrentRateLimiterState = contract.NewRead(contract.ReadParams[GetCurrentRateLimiterStateArgs, RateLimiterStates, *token_pool.TokenPool]{
	Name:         "token-pool:get-current-rate-limiter-state",
	Version:      Version,
	Description:  "Gets both outbound and inbound rate limiter states for a given remote chain selector",
	ContractType: ContractType,
	NewContract:  token_pool.NewTokenPool,
	CallContract: func(tokenPool *token_pool.TokenPool, opts *bind.CallOpts, args GetCurrentRateLimiterStateArgs) (RateLimiterStates, error) {
		return tokenPool.GetCurrentRateLimiterState(opts, args.RemoteChainSelector, args.CustomBlockConfirmation)
	},
})

var GetMinBlockConfirmation = contract.NewRead(contract.ReadParams[any, uint16, *token_pool.TokenPool]{
	Name:         "token-pool:get-configured-min-block-confirmation",
	Version:      Version,
	Description:  "Gets the globally configured minimum block confirmations for custom block confirmation transfers",
	ContractType: ContractType,
	NewContract:  token_pool.NewTokenPool,
	CallContract: func(tokenPool *token_pool.TokenPool, opts *bind.CallOpts, args any) (uint16, error) {
		return tokenPool.GetMinBlockConfirmations(opts)
	},
})

var GetCurrentRateLimiterStateByRemoteChainSelector = contract.NewRead(contract.ReadParams[uint64, RateLimiterStates, *token_pool.TokenPool]{
	Name:         "token-pool:get-current-rate-limiter-state-by-remote-chain-selector",
	Version:      Version,
	Description:  "Gets the outbound and inbound rate limiter states for a given remote chain selector",
	ContractType: ContractType,
	NewContract:  token_pool.NewTokenPool,
	CallContract: func(tokenPool *token_pool.TokenPool, opts *bind.CallOpts, args uint64) (RateLimiterStates, error) {
		return tokenPool.GetCurrentRateLimiterState(opts, args, false)
	},
})

var GetSupportedChains = contract.NewRead(contract.ReadParams[any, []uint64, *token_pool.TokenPool]{
	Name:         "token-pool:supported-chains",
	Version:      Version,
	Description:  "Gets the list of supported remote chain selectors on a TokenPool",
	ContractType: ContractType,
	NewContract:  token_pool.NewTokenPool,
	CallContract: func(tokenPool *token_pool.TokenPool, opts *bind.CallOpts, args any) ([]uint64, error) {
		return tokenPool.GetSupportedChains(opts)
	},
})

var GetRemotePools = contract.NewRead(contract.ReadParams[uint64, [][]byte, *token_pool.TokenPool]{
	Name:         "token-pool:get-remote-pools",
	Version:      Version,
	Description:  "Gets the remote pool addresses for a given remote chain selector on a TokenPool",
	ContractType: ContractType,
	NewContract:  token_pool.NewTokenPool,
	CallContract: func(tokenPool *token_pool.TokenPool, opts *bind.CallOpts, args uint64) ([][]byte, error) {
		return tokenPool.GetRemotePools(opts, args)
	},
})

var GetRemoteToken = contract.NewRead(contract.ReadParams[uint64, []byte, *token_pool.TokenPool]{
	Name:         "token-pool:get-remote-token",
	Version:      Version,
	Description:  "Gets the remote pool address for a given remote chain selector on a TokenPool",
	ContractType: ContractType,
	NewContract:  token_pool.NewTokenPool,
	CallContract: func(tokenPool *token_pool.TokenPool, opts *bind.CallOpts, args uint64) ([]byte, error) {
		return tokenPool.GetRemoteToken(opts, args)
	},
})

var GetToken = contract.NewRead(contract.ReadParams[any, common.Address, *token_pool.TokenPool]{
	Name:         "token-pool:get-token",
	Version:      Version,
	Description:  "Gets the local token address for a TokenPool",
	ContractType: ContractType,
	NewContract:  token_pool.NewTokenPool,
	CallContract: func(tokenPool *token_pool.TokenPool, opts *bind.CallOpts, args any) (common.Address, error) {
		return tokenPool.GetToken(opts)
	},
})

var GetTokenDecimals = contract.NewRead(contract.ReadParams[any, uint8, *token_pool.TokenPool]{
	Name:         "token-pool:get-token-decimals",
	Version:      Version,
	Description:  "Gets the token decimals for a TokenPool",
	ContractType: ContractType,
	NewContract:  token_pool.NewTokenPool,
	CallContract: func(tokenPool *token_pool.TokenPool, opts *bind.CallOpts, args any) (uint8, error) {
		return tokenPool.GetTokenDecimals(opts)
	},
})

var IsSupportedToken = contract.NewRead(contract.ReadParams[common.Address, bool, *token_pool.TokenPool]{
	Name:         "token-pool:is-supported-token",
	Version:      Version,
	Description:  "Checks whether a token is supported by a TokenPool",
	ContractType: ContractType,
	NewContract:  token_pool.NewTokenPool,
	CallContract: func(tokenPool *token_pool.TokenPool, opts *bind.CallOpts, token common.Address) (bool, error) {
		return tokenPool.IsSupportedToken(opts, token)
	},
})

var GetAdvancedPoolHooks = contract.NewRead(contract.ReadParams[any, common.Address, *token_pool.TokenPool]{
	Name:         "token-pool:get-advanced-pool-hooks",
	Version:      Version,
	Description:  "Gets the advanced pool hooks address on a TokenPool",
	ContractType: ContractType,
	NewContract:  token_pool.NewTokenPool,
	CallContract: func(tokenPool *token_pool.TokenPool, opts *bind.CallOpts, args any) (common.Address, error) {
		return tokenPool.GetAdvancedPoolHooks(opts)
	},
})
