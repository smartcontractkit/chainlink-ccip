package token_pool

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	gobindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_1/token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
)

// Config is the rate limiter configuration tuple used on token pools.
type Config = gobindings.RateLimiterConfig

// TokenBucket is the on-chain rate limiter token bucket type.
type TokenBucket = gobindings.RateLimiterTokenBucket

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

var SetChainRateLimiterConfig = contract.NewWrite(contract.WriteParams[SetChainRateLimiterConfigArgs, *gobindings.TokenPool]{
	Name:            "token-pool:set-chain-rate-limiter-config",
	Version:         Version,
	Description:     "Calls setChainRateLimiterConfig on the contract",
	ContractType:    ContractType,
	ContractABI:     gobindings.TokenPoolMetaData.ABI,
	NewContract:     gobindings.NewTokenPool,
	IsAllowedCaller: contract.OnlyOwner[*gobindings.TokenPool, SetChainRateLimiterConfigArgs],
	Validate:        func(SetChainRateLimiterConfigArgs) error { return nil },
	CallContract: func(c *gobindings.TokenPool, opts *bind.TransactOpts, args SetChainRateLimiterConfigArgs) (*types.Transaction, error) {
		return c.SetChainRateLimiterConfig(opts, args.RemoteChainSelector, args.OutboundConfig, args.InboundConfig)
	},
})

var SetRateLimitAdmin = contract.NewWrite(contract.WriteParams[common.Address, *gobindings.TokenPool]{
	Name:            "token-pool:set-rate-limit-admin",
	Version:         Version,
	Description:     "Calls setRateLimitAdmin on the contract",
	ContractType:    ContractType,
	ContractABI:     gobindings.TokenPoolMetaData.ABI,
	NewContract:     gobindings.NewTokenPool,
	IsAllowedCaller: contract.OnlyOwner[*gobindings.TokenPool, common.Address],
	Validate:        func(common.Address) error { return nil },
	CallContract: func(c *gobindings.TokenPool, opts *bind.TransactOpts, args common.Address) (*types.Transaction, error) {
		return c.SetRateLimitAdmin(opts, args)
	},
})

var GetCurrentInboundRateLimiterState = contract.NewRead(contract.ReadParams[uint64, gobindings.RateLimiterTokenBucket, *gobindings.TokenPool]{
	Name:         "token-pool:get-current-inbound-rate-limiter-state",
	Version:      Version,
	Description:  "Calls getCurrentInboundRateLimiterState on the contract",
	ContractType: ContractType,
	NewContract:  gobindings.NewTokenPool,
	CallContract: func(c *gobindings.TokenPool, opts *bind.CallOpts, args uint64) (gobindings.RateLimiterTokenBucket, error) {
		return c.GetCurrentInboundRateLimiterState(opts, args)
	},
})

var GetCurrentOutboundRateLimiterState = contract.NewRead(contract.ReadParams[uint64, gobindings.RateLimiterTokenBucket, *gobindings.TokenPool]{
	Name:         "token-pool:get-current-outbound-rate-limiter-state",
	Version:      Version,
	Description:  "Calls getCurrentOutboundRateLimiterState on the contract",
	ContractType: ContractType,
	NewContract:  gobindings.NewTokenPool,
	CallContract: func(c *gobindings.TokenPool, opts *bind.CallOpts, args uint64) (gobindings.RateLimiterTokenBucket, error) {
		return c.GetCurrentOutboundRateLimiterState(opts, args)
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
