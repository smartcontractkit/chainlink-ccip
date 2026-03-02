package token_pool

import (
	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"

	// We use the proxy bindings because we need to access the GetPreviousPool function.
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_5_0/burn_mint_token_pool_and_proxy"
)

var ContractType cldf_deployment.ContractType = "TokenPool"

var Version *semver.Version = semver.MustParse("1.5.0")

var GetRemotePool = contract.NewRead(contract.ReadParams[uint64, []byte, *burn_mint_token_pool_and_proxy.BurnMintTokenPoolAndProxy]{
	Name:         "token-pool:get-remote-pool",
	Version:      Version,
	Description:  "Gets the remote pool address for a TokenPool",
	ContractType: ContractType,
	NewContract:  burn_mint_token_pool_and_proxy.NewBurnMintTokenPoolAndProxy,
	CallContract: func(tokenPool *burn_mint_token_pool_and_proxy.BurnMintTokenPoolAndProxy, opts *bind.CallOpts, remoteChainSelector uint64) ([]byte, error) {
		return tokenPool.GetRemotePool(opts, remoteChainSelector)
	},
})

var GetInboundRateLimiterState = contract.NewRead(contract.ReadParams[uint64, burn_mint_token_pool_and_proxy.RateLimiterTokenBucket, *burn_mint_token_pool_and_proxy.BurnMintTokenPoolAndProxy]{
	Name:         "token-pool:get-inbound-rate-limiter-state",
	Version:      Version,
	Description:  "Gets the inbound rate limiter state for a TokenPool",
	ContractType: ContractType,
	NewContract:  burn_mint_token_pool_and_proxy.NewBurnMintTokenPoolAndProxy,
	CallContract: func(tokenPool *burn_mint_token_pool_and_proxy.BurnMintTokenPoolAndProxy, opts *bind.CallOpts, remoteChainSelector uint64) (burn_mint_token_pool_and_proxy.RateLimiterTokenBucket, error) {
		return tokenPool.GetCurrentInboundRateLimiterState(opts, remoteChainSelector)
	},
})

var GetOutboundRateLimiterState = contract.NewRead(contract.ReadParams[uint64, burn_mint_token_pool_and_proxy.RateLimiterTokenBucket, *burn_mint_token_pool_and_proxy.BurnMintTokenPoolAndProxy]{
	Name:         "token-pool:get-outbound-rate-limiter-state",
	Version:      Version,
	Description:  "Gets the outbound rate limiter state for a TokenPool",
	ContractType: ContractType,
	NewContract:  burn_mint_token_pool_and_proxy.NewBurnMintTokenPoolAndProxy,
	CallContract: func(tokenPool *burn_mint_token_pool_and_proxy.BurnMintTokenPoolAndProxy, opts *bind.CallOpts, remoteChainSelector uint64) (burn_mint_token_pool_and_proxy.RateLimiterTokenBucket, error) {
		return tokenPool.GetCurrentOutboundRateLimiterState(opts, remoteChainSelector)
	},
})

var GetPreviousPool = contract.NewRead(contract.ReadParams[struct{}, common.Address, *burn_mint_token_pool_and_proxy.BurnMintTokenPoolAndProxy]{
	Name:         "token-pool:get-previous-pool",
	Version:      Version,
	Description:  "Gets the previous pool address for a TokenPool",
	ContractType: ContractType,
	NewContract:  burn_mint_token_pool_and_proxy.NewBurnMintTokenPoolAndProxy,
	CallContract: func(tokenPool *burn_mint_token_pool_and_proxy.BurnMintTokenPoolAndProxy, opts *bind.CallOpts, args struct{}) (common.Address, error) {
		return tokenPool.GetPreviousPool(opts)
	},
})
