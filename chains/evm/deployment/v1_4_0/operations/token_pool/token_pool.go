// Package token_pool exposes the shared v1.4.0 TokenPool surface.
//
// v1.4.0 pools are the pre-proxy generation: BurnMintTokenPool and
// LockReleaseTokenPool are deployed behind a separate Proxy contract, and the
// pool itself stores only remote chain selectors and rate limits. It has no
// remote token or remote pool storage - those live on the OnRamp/OffRamp - and
// no TokenAdminRegistry exists at this version.
//
// Every read and write below is signature-identical between BurnMintTokenPool
// and LockReleaseTokenPool, so both pool types are driven through this one
// package. LockReleaseTokenPool adds a small liquidity surface on top, which
// lives in the lock_release_token_pool package.
package token_pool

import (
	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_4_0/token_pool"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations/contract"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

var ContractType cldf_deployment.ContractType = "TokenPool"
var Version *semver.Version = semver.MustParse("1.4.0")

// ChainUpdate mirrors TokenPool.ChainUpdate. Unlike v1.5.0 it carries no
// remotePoolAddress or remoteTokenAddress: a v1.4.0 pool knows nothing about
// the remote side of a lane.
type ChainUpdate = token_pool.TokenPoolChainUpdate

// RateLimiterConfig mirrors RateLimiter.Config.
type RateLimiterConfig = token_pool.RateLimiterConfig

// TokenBucket mirrors RateLimiter.TokenBucket.
type TokenBucket = token_pool.RateLimiterTokenBucket

var GetToken = contract.NewRead(contract.ReadParams[struct{}, common.Address, *token_pool.TokenPool]{
	Name:         "token-pool:get-token",
	Version:      Version,
	Description:  "Gets the local token address for a v1.4.0 TokenPool",
	ContractType: ContractType,
	NewContract:  token_pool.NewTokenPool,
	CallContract: func(tokenPool *token_pool.TokenPool, opts *bind.CallOpts, _ struct{}) (common.Address, error) {
		return tokenPool.GetToken(opts)
	},
})

var GetRouter = contract.NewRead(contract.ReadParams[struct{}, common.Address, *token_pool.TokenPool]{
	Name:         "token-pool:get-router",
	Version:      Version,
	Description:  "Gets the router address for a v1.4.0 TokenPool",
	ContractType: ContractType,
	NewContract:  token_pool.NewTokenPool,
	CallContract: func(tokenPool *token_pool.TokenPool, opts *bind.CallOpts, _ struct{}) (common.Address, error) {
		return tokenPool.GetRouter(opts)
	},
})

var GetArmProxy = contract.NewRead(contract.ReadParams[struct{}, common.Address, *token_pool.TokenPool]{
	Name:         "token-pool:get-arm-proxy",
	Version:      Version,
	Description:  "Gets the ARM proxy address for a v1.4.0 TokenPool",
	ContractType: ContractType,
	NewContract:  token_pool.NewTokenPool,
	CallContract: func(tokenPool *token_pool.TokenPool, opts *bind.CallOpts, _ struct{}) (common.Address, error) {
		return tokenPool.GetArmProxy(opts)
	},
})

var GetSupportedChains = contract.NewRead(contract.ReadParams[struct{}, []uint64, *token_pool.TokenPool]{
	Name:         "token-pool:get-supported-chains",
	Version:      Version,
	Description:  "Gets the remote chain selectors a v1.4.0 TokenPool is configured for",
	ContractType: ContractType,
	NewContract:  token_pool.NewTokenPool,
	CallContract: func(tokenPool *token_pool.TokenPool, opts *bind.CallOpts, _ struct{}) ([]uint64, error) {
		return tokenPool.GetSupportedChains(opts)
	},
})

var IsSupportedChain = contract.NewRead(contract.ReadParams[uint64, bool, *token_pool.TokenPool]{
	Name:         "token-pool:is-supported-chain",
	Version:      Version,
	Description:  "Reports whether a remote chain selector is configured on a v1.4.0 TokenPool",
	ContractType: ContractType,
	NewContract:  token_pool.NewTokenPool,
	CallContract: func(tokenPool *token_pool.TokenPool, opts *bind.CallOpts, remoteChainSelector uint64) (bool, error) {
		return tokenPool.IsSupportedChain(opts, remoteChainSelector)
	},
})

var GetCurrentInboundRateLimiterState = contract.NewRead(contract.ReadParams[uint64, TokenBucket, *token_pool.TokenPool]{
	Name:         "token-pool:get-current-inbound-rate-limiter-state",
	Version:      Version,
	Description:  "Gets the current inbound rate limiter state for a remote chain on a v1.4.0 TokenPool",
	ContractType: ContractType,
	NewContract:  token_pool.NewTokenPool,
	CallContract: func(tokenPool *token_pool.TokenPool, opts *bind.CallOpts, remoteChainSelector uint64) (TokenBucket, error) {
		return tokenPool.GetCurrentInboundRateLimiterState(opts, remoteChainSelector)
	},
})

var GetCurrentOutboundRateLimiterState = contract.NewRead(contract.ReadParams[uint64, TokenBucket, *token_pool.TokenPool]{
	Name:         "token-pool:get-current-outbound-rate-limiter-state",
	Version:      Version,
	Description:  "Gets the current outbound rate limiter state for a remote chain on a v1.4.0 TokenPool",
	ContractType: ContractType,
	NewContract:  token_pool.NewTokenPool,
	CallContract: func(tokenPool *token_pool.TokenPool, opts *bind.CallOpts, remoteChainSelector uint64) (TokenBucket, error) {
		return tokenPool.GetCurrentOutboundRateLimiterState(opts, remoteChainSelector)
	},
})

var Owner = contract.NewRead(contract.ReadParams[struct{}, common.Address, *token_pool.TokenPool]{
	Name:         "token-pool:owner",
	Version:      Version,
	Description:  "Gets the owner of a v1.4.0 TokenPool",
	ContractType: ContractType,
	NewContract:  token_pool.NewTokenPool,
	CallContract: func(tokenPool *token_pool.TokenPool, opts *bind.CallOpts, _ struct{}) (common.Address, error) {
		return tokenPool.Owner(opts)
	},
})

// ApplyChainUpdatesArgs is the input for ApplyChainUpdates.
type ApplyChainUpdatesArgs struct {
	Chains []ChainUpdate
}

// ApplyChainUpdates adds or removes remote chains on a v1.4.0 pool. Removals are
// expressed as entries with Allowed=false, which deletes the whole remote chain
// config (rate limits included) rather than a single remote pool.
var ApplyChainUpdates = contract.NewWrite(contract.WriteParams[ApplyChainUpdatesArgs, *token_pool.TokenPool]{
	Name:            "token-pool:apply-chain-updates",
	Version:         Version,
	Description:     "Applies remote chain updates to a v1.4.0 TokenPool",
	ContractType:    ContractType,
	ContractABI:     token_pool.TokenPoolABI,
	NewContract:     token_pool.NewTokenPool,
	IsAllowedCaller: contract.OnlyOwner[*token_pool.TokenPool, ApplyChainUpdatesArgs],
	Validate:        func(_ ApplyChainUpdatesArgs) error { return nil },
	CallContract: func(tokenPool *token_pool.TokenPool, opts *bind.TransactOpts, args ApplyChainUpdatesArgs) (*types.Transaction, error) {
		return tokenPool.ApplyChainUpdates(opts, args.Chains)
	},
})

// SetChainRateLimiterConfigArgs is the input for SetChainRateLimiterConfig.
type SetChainRateLimiterConfigArgs struct {
	RemoteChainSelector uint64
	OutboundConfig      RateLimiterConfig
	InboundConfig       RateLimiterConfig
}

// SetChainRateLimiterConfig sets the rate limits for a remote chain on a v1.4.0
// pool. On the base TokenPool this is onlyOwner; LockReleaseTokenPool widens it
// to the rate limit admin as well.
var SetChainRateLimiterConfig = contract.NewWrite(contract.WriteParams[SetChainRateLimiterConfigArgs, *token_pool.TokenPool]{
	Name:            "token-pool:set-chain-rate-limiter-config",
	Version:         Version,
	Description:     "Sets the rate limiter config for a remote chain on a v1.4.0 TokenPool",
	ContractType:    ContractType,
	ContractABI:     token_pool.TokenPoolABI,
	NewContract:     token_pool.NewTokenPool,
	IsAllowedCaller: contract.OnlyOwner[*token_pool.TokenPool, SetChainRateLimiterConfigArgs],
	Validate:        func(_ SetChainRateLimiterConfigArgs) error { return nil },
	CallContract: func(tokenPool *token_pool.TokenPool, opts *bind.TransactOpts, args SetChainRateLimiterConfigArgs) (*types.Transaction, error) {
		return tokenPool.SetChainRateLimiterConfig(opts, args.RemoteChainSelector, args.OutboundConfig, args.InboundConfig)
	},
})

// SetRouterArgs is the input for SetRouter.
type SetRouterArgs struct {
	NewRouter common.Address
}

var SetRouter = contract.NewWrite(contract.WriteParams[SetRouterArgs, *token_pool.TokenPool]{
	Name:            "token-pool:set-router",
	Version:         Version,
	Description:     "Sets the router on a v1.4.0 TokenPool",
	ContractType:    ContractType,
	ContractABI:     token_pool.TokenPoolABI,
	NewContract:     token_pool.NewTokenPool,
	IsAllowedCaller: contract.OnlyOwner[*token_pool.TokenPool, SetRouterArgs],
	Validate:        func(_ SetRouterArgs) error { return nil },
	CallContract: func(tokenPool *token_pool.TokenPool, opts *bind.TransactOpts, args SetRouterArgs) (*types.Transaction, error) {
		return tokenPool.SetRouter(opts, args.NewRouter)
	},
})
