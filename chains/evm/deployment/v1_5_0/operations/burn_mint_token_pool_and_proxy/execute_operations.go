package burn_mint_token_pool_and_proxy

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	gobindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_5_0/burn_mint_token_pool_and_proxy"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
)

// TokenBucket is the on-chain rate limiter token bucket type (v1.5.0 proxy pool).
type TokenBucket = gobindings.RateLimiterTokenBucket

var GetPreviousPool = contract.NewRead(contract.ReadParams[struct{}, common.Address, *gobindings.BurnMintTokenPoolAndProxy]{
	Name:         "burn-mint-token-pool-and-proxy:get-previous-pool",
	Version:      Version,
	Description:  "Calls getPreviousPool on the contract",
	ContractType: ContractType,
	NewContract:  gobindings.NewBurnMintTokenPoolAndProxy,
	CallContract: func(c *gobindings.BurnMintTokenPoolAndProxy, opts *bind.CallOpts, args struct{}) (common.Address, error) {
		return c.GetPreviousPool(opts)
	},
})

var GetCurrentInboundRateLimiterState = contract.NewRead(contract.ReadParams[uint64, gobindings.RateLimiterTokenBucket, *gobindings.BurnMintTokenPoolAndProxy]{
	Name:         "burn-mint-token-pool-and-proxy:get-current-inbound-rate-limiter-state",
	Version:      Version,
	Description:  "Calls getCurrentInboundRateLimiterState on the contract",
	ContractType: ContractType,
	NewContract:  gobindings.NewBurnMintTokenPoolAndProxy,
	CallContract: func(c *gobindings.BurnMintTokenPoolAndProxy, opts *bind.CallOpts, args uint64) (gobindings.RateLimiterTokenBucket, error) {
		return c.GetCurrentInboundRateLimiterState(opts, args)
	},
})

var GetCurrentOutboundRateLimiterState = contract.NewRead(contract.ReadParams[uint64, gobindings.RateLimiterTokenBucket, *gobindings.BurnMintTokenPoolAndProxy]{
	Name:         "burn-mint-token-pool-and-proxy:get-current-outbound-rate-limiter-state",
	Version:      Version,
	Description:  "Calls getCurrentOutboundRateLimiterState on the contract",
	ContractType: ContractType,
	NewContract:  gobindings.NewBurnMintTokenPoolAndProxy,
	CallContract: func(c *gobindings.BurnMintTokenPoolAndProxy, opts *bind.CallOpts, args uint64) (gobindings.RateLimiterTokenBucket, error) {
		return c.GetCurrentOutboundRateLimiterState(opts, args)
	},
})

var GetRemotePool = contract.NewRead(contract.ReadParams[uint64, []byte, *gobindings.BurnMintTokenPoolAndProxy]{
	Name:         "burn-mint-token-pool-and-proxy:get-remote-pool",
	Version:      Version,
	Description:  "Calls getRemotePool on the contract",
	ContractType: ContractType,
	NewContract:  gobindings.NewBurnMintTokenPoolAndProxy,
	CallContract: func(c *gobindings.BurnMintTokenPoolAndProxy, opts *bind.CallOpts, args uint64) ([]byte, error) {
		return c.GetRemotePool(opts, args)
	},
})
