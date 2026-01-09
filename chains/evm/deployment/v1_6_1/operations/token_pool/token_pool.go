package token_pool

import (
	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_1/burn_mint_token_pool"
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

type CCVConfigArg = burn_mint_token_pool.BurnMintTokenPool

type RateLimiterConfig = burn_mint_token_pool.RateLimiterConfig

type RateLimiterTokenBucket = burn_mint_token_pool.RateLimiterTokenBucket

type ApplyChainUpdatesArgs struct {
	RemoteChainSelectorsToRemove []uint64
	ChainsToAdd                  []ChainUpdate
}

type RemotePoolArgs struct {
	RemoteChainSelector uint64
	RemotePoolAddress   []byte
}

type SetChainRateLimiterConfigArg struct {
	RemoteChainSelector       uint64
	InboundRateLimiterConfig  tokens.RateLimiterConfig
	OutboundRateLimiterConfig tokens.RateLimiterConfig
}

type ApplyAllowListUpdatesArgs struct {
	Adds    []common.Address
	Removes []common.Address
}

type WithdrawFeeTokensArgs struct {
	FeeTokens []common.Address
	Recipient common.Address
}

var ApplyChainUpdates = contract.NewWrite(contract.WriteParams[ApplyChainUpdatesArgs, *burn_mint_token_pool.BurnMintTokenPool]{
	Name:            "token-pool:apply-chain-updates",
	Version:         Version,
	Description:     "Applies chain updates to a TokenPool, enabling / disabling remote chains and setting rate limits",
	ContractType:    ContractType,
	ContractABI:     burn_mint_token_pool.BurnMintTokenPoolABI,
	NewContract:     burn_mint_token_pool.NewBurnMintTokenPool,
	IsAllowedCaller: contract.OnlyOwner[*burn_mint_token_pool.BurnMintTokenPool, ApplyChainUpdatesArgs],
	Validate:        func(ApplyChainUpdatesArgs) error { return nil },
	CallContract: func(tokenPool *burn_mint_token_pool.BurnMintTokenPool, opts *bind.TransactOpts, args ApplyChainUpdatesArgs) (*types.Transaction, error) {
		chainsToAdd := make([]burn_mint_token_pool.TokenPoolChainUpdate, len(args.ChainsToAdd))
		for i, chain := range args.ChainsToAdd {
			chainsToAdd[i] = burn_mint_token_pool.TokenPoolChainUpdate{
				RemoteChainSelector:       chain.RemoteChainSelector,
				RemotePoolAddresses:       chain.RemotePoolAddresses,
				RemoteTokenAddress:        chain.RemoteTokenAddress,
				OutboundRateLimiterConfig: burn_mint_token_pool.RateLimiterConfig(chain.OutboundRateLimiterConfig),
				InboundRateLimiterConfig:  burn_mint_token_pool.RateLimiterConfig(chain.InboundRateLimiterConfig),
			}
		}
		return tokenPool.ApplyChainUpdates(opts, args.RemoteChainSelectorsToRemove, chainsToAdd)
	},
})

var AddRemotePool = contract.NewWrite(contract.WriteParams[RemotePoolArgs, *burn_mint_token_pool.BurnMintTokenPool]{
	Name:            "token-pool:add-remote-pool",
	Version:         Version,
	Description:     "Adds a remote pool for a given chain selector to a TokenPool",
	ContractType:    ContractType,
	ContractABI:     burn_mint_token_pool.BurnMintTokenPoolABI,
	NewContract:     burn_mint_token_pool.NewBurnMintTokenPool,
	IsAllowedCaller: contract.OnlyOwner[*burn_mint_token_pool.BurnMintTokenPool, RemotePoolArgs],
	Validate:        func(RemotePoolArgs) error { return nil },
	CallContract: func(tokenPool *burn_mint_token_pool.BurnMintTokenPool, opts *bind.TransactOpts, args RemotePoolArgs) (*types.Transaction, error) {
		return tokenPool.AddRemotePool(opts, args.RemoteChainSelector, args.RemotePoolAddress)
	},
})

var RemoveRemotePool = contract.NewWrite(contract.WriteParams[RemotePoolArgs, *burn_mint_token_pool.BurnMintTokenPool]{
	Name:            "token-pool:remove-remote-pool",
	Version:         Version,
	Description:     "Removes a remote pool for a given chain selector from a TokenPool",
	ContractType:    ContractType,
	ContractABI:     burn_mint_token_pool.BurnMintTokenPoolABI,
	NewContract:     burn_mint_token_pool.NewBurnMintTokenPool,
	IsAllowedCaller: contract.OnlyOwner[*burn_mint_token_pool.BurnMintTokenPool, RemotePoolArgs],
	Validate:        func(RemotePoolArgs) error { return nil },
	CallContract: func(tokenPool *burn_mint_token_pool.BurnMintTokenPool, opts *bind.TransactOpts, args RemotePoolArgs) (*types.Transaction, error) {
		return tokenPool.RemoveRemotePool(opts, args.RemoteChainSelector, args.RemotePoolAddress)
	},
})

var SetRateLimitConfig = contract.NewWrite(contract.WriteParams[[]SetChainRateLimiterConfigArg, *burn_mint_token_pool.BurnMintTokenPool]{
	Name:            "token-pool:set-rate-limit-config",
	Version:         Version,
	Description:     "Sets the rate limit configs for existing remote chains on a TokenPool",
	ContractType:    ContractType,
	ContractABI:     burn_mint_token_pool.BurnMintTokenPoolABI,
	NewContract:     burn_mint_token_pool.NewBurnMintTokenPool,
	IsAllowedCaller: contract.OnlyOwner[*burn_mint_token_pool.BurnMintTokenPool, []SetChainRateLimiterConfigArg],
	Validate:        func([]SetChainRateLimiterConfigArg) error { return nil },
	CallContract: func(tokenPool *burn_mint_token_pool.BurnMintTokenPool, opts *bind.TransactOpts, args []SetChainRateLimiterConfigArg) (*types.Transaction, error) {
		inboundRateLimitConfigs := make([]burn_mint_token_pool.RateLimiterConfig, 0, len(args))
		outboundRateLimitConfigs := make([]burn_mint_token_pool.RateLimiterConfig, 0, len(args))
		remoteChainSelectors := make([]uint64, 0, len(args))
		for _, arg := range args {
			remoteChainSelectors = append(remoteChainSelectors, arg.RemoteChainSelector)
			outboundRateLimitConfigs = append(outboundRateLimitConfigs, burn_mint_token_pool.RateLimiterConfig{
				Rate:      arg.OutboundRateLimiterConfig.Rate,
				Capacity:  arg.OutboundRateLimiterConfig.Capacity,
				IsEnabled: arg.OutboundRateLimiterConfig.IsEnabled,
			})
			inboundRateLimitConfigs = append(inboundRateLimitConfigs, burn_mint_token_pool.RateLimiterConfig{
				Rate:      arg.InboundRateLimiterConfig.Rate,
				Capacity:  arg.InboundRateLimiterConfig.Capacity,
				IsEnabled: arg.InboundRateLimiterConfig.IsEnabled,
			})
		}
		return tokenPool.SetChainRateLimiterConfigs(opts, remoteChainSelectors, outboundRateLimitConfigs, inboundRateLimitConfigs)
	},
})

var SetRouter = contract.NewWrite(contract.WriteParams[common.Address, *burn_mint_token_pool.BurnMintTokenPool]{
	Name:            "token-pool:set-router-and-rate-limit-admin",
	Version:         Version,
	Description:     "Sets the router and rate limit admin for a TokenPool",
	ContractType:    ContractType,
	ContractABI:     burn_mint_token_pool.BurnMintTokenPoolABI,
	NewContract:     burn_mint_token_pool.NewBurnMintTokenPool,
	IsAllowedCaller: contract.OnlyOwner[*burn_mint_token_pool.BurnMintTokenPool, common.Address],
	Validate:        func(common.Address) error { return nil },
	CallContract: func(tokenPool *burn_mint_token_pool.BurnMintTokenPool, opts *bind.TransactOpts, args common.Address) (*types.Transaction, error) {
		return tokenPool.SetRouter(opts, args)
	},
})

var SetRateLimitAdmin = contract.NewWrite(contract.WriteParams[common.Address, *burn_mint_token_pool.BurnMintTokenPool]{
	Name:            "token-pool:set-rate-limit-admin",
	Version:         Version,
	Description:     "Sets the rate limit admin on a TokenPool",
	ContractType:    ContractType,
	ContractABI:     burn_mint_token_pool.BurnMintTokenPoolABI,
	NewContract:     burn_mint_token_pool.NewBurnMintTokenPool,
	IsAllowedCaller: contract.OnlyOwner[*burn_mint_token_pool.BurnMintTokenPool, common.Address],
	Validate:        func(common.Address) error { return nil },
	CallContract: func(tokenPool *burn_mint_token_pool.BurnMintTokenPool, opts *bind.TransactOpts, args common.Address) (*types.Transaction, error) {
		return tokenPool.SetRateLimitAdmin(opts, args)
	},
})

var ApplyAllowlistUpdates = contract.NewWrite(contract.WriteParams[ApplyAllowListUpdatesArgs, *burn_mint_token_pool.BurnMintTokenPool]{
	Name:            "token-pool:apply-allowlist-updates",
	Version:         Version,
	Description:     "Applies allow list updates to a TokenPool",
	ContractType:    ContractType,
	ContractABI:     burn_mint_token_pool.BurnMintTokenPoolABI,
	NewContract:     burn_mint_token_pool.NewBurnMintTokenPool,
	IsAllowedCaller: contract.OnlyOwner[*burn_mint_token_pool.BurnMintTokenPool, ApplyAllowListUpdatesArgs],
	Validate:        func(ApplyAllowListUpdatesArgs) error { return nil },
	CallContract: func(tokenPool *burn_mint_token_pool.BurnMintTokenPool, opts *bind.TransactOpts, args ApplyAllowListUpdatesArgs) (*types.Transaction, error) {
		return tokenPool.ApplyAllowListUpdates(opts, args.Removes, args.Adds)
	},
})

var GetRMNProxy = contract.NewRead(contract.ReadParams[any, common.Address, *burn_mint_token_pool.BurnMintTokenPool]{
	Name:         "token-pool:get-rmn-proxy",
	Version:      Version,
	Description:  "Gets the RMN proxy address on a TokenPool",
	ContractType: ContractType,
	NewContract:  burn_mint_token_pool.NewBurnMintTokenPool,
	CallContract: func(tokenPool *burn_mint_token_pool.BurnMintTokenPool, opts *bind.CallOpts, args any) (common.Address, error) {
		return tokenPool.GetRmnProxy(opts)
	},
})

var GetCurrentInboundRateLimiterState = contract.NewRead(contract.ReadParams[uint64, RateLimiterTokenBucket, *burn_mint_token_pool.BurnMintTokenPool]{
	Name:         "token-pool:get-current-rate-limiter-state-by-remote-chain-selector",
	Version:      Version,
	Description:  "Gets the outbound and inbound rate limiter states for a given remote chain selector",
	ContractType: ContractType,
	NewContract:  burn_mint_token_pool.NewBurnMintTokenPool,
	CallContract: func(tokenPool *burn_mint_token_pool.BurnMintTokenPool, opts *bind.CallOpts, args uint64) (RateLimiterTokenBucket, error) {
		return tokenPool.GetCurrentInboundRateLimiterState(opts, args)
	},
})

var GetCurrentOutboundRateLimiterState = contract.NewRead(contract.ReadParams[uint64, RateLimiterTokenBucket, *burn_mint_token_pool.BurnMintTokenPool]{
	Name:         "token-pool:get-current-rate-limiter-state-by-remote-chain-selector",
	Version:      Version,
	Description:  "Gets the outbound and inbound rate limiter states for a given remote chain selector",
	ContractType: ContractType,
	NewContract:  burn_mint_token_pool.NewBurnMintTokenPool,
	CallContract: func(tokenPool *burn_mint_token_pool.BurnMintTokenPool, opts *bind.CallOpts, args uint64) (RateLimiterTokenBucket, error) {
		return tokenPool.GetCurrentOutboundRateLimiterState(opts, args)
	},
})

var GetSupportedChains = contract.NewRead(contract.ReadParams[any, []uint64, *burn_mint_token_pool.BurnMintTokenPool]{
	Name:         "token-pool:supported-chains",
	Version:      Version,
	Description:  "Gets the list of supported remote chain selectors on a TokenPool",
	ContractType: ContractType,
	NewContract:  burn_mint_token_pool.NewBurnMintTokenPool,
	CallContract: func(tokenPool *burn_mint_token_pool.BurnMintTokenPool, opts *bind.CallOpts, args any) ([]uint64, error) {
		return tokenPool.GetSupportedChains(opts)
	},
})

var GetRemotePools = contract.NewRead(contract.ReadParams[uint64, [][]byte, *burn_mint_token_pool.BurnMintTokenPool]{
	Name:         "token-pool:get-remote-pools",
	Version:      Version,
	Description:  "Gets the remote pool addresses for a given remote chain selector on a TokenPool",
	ContractType: ContractType,
	NewContract:  burn_mint_token_pool.NewBurnMintTokenPool,
	CallContract: func(tokenPool *burn_mint_token_pool.BurnMintTokenPool, opts *bind.CallOpts, args uint64) ([][]byte, error) {
		return tokenPool.GetRemotePools(opts, args)
	},
})

var GetRemoteToken = contract.NewRead(contract.ReadParams[uint64, []byte, *burn_mint_token_pool.BurnMintTokenPool]{
	Name:         "token-pool:get-remote-token",
	Version:      Version,
	Description:  "Gets the remote pool address for a given remote chain selector on a TokenPool",
	ContractType: ContractType,
	NewContract:  burn_mint_token_pool.NewBurnMintTokenPool,
	CallContract: func(tokenPool *burn_mint_token_pool.BurnMintTokenPool, opts *bind.CallOpts, args uint64) ([]byte, error) {
		return tokenPool.GetRemoteToken(opts, args)
	},
})

var GetToken = contract.NewRead(contract.ReadParams[any, common.Address, *burn_mint_token_pool.BurnMintTokenPool]{
	Name:         "token-pool:get-token",
	Version:      Version,
	Description:  "Gets the local token address for a TokenPool",
	ContractType: ContractType,
	NewContract:  burn_mint_token_pool.NewBurnMintTokenPool,
	CallContract: func(tokenPool *burn_mint_token_pool.BurnMintTokenPool, opts *bind.CallOpts, args any) (common.Address, error) {
		return tokenPool.GetToken(opts)
	},
})

var GetAllowListEnabled = contract.NewRead(contract.ReadParams[any, bool, *burn_mint_token_pool.BurnMintTokenPool]{
	Name:         "token-pool:get-allow-list-enabled",
	Version:      Version,
	Description:  "Gets the allow list enabled status on a TokenPool",
	ContractType: ContractType,
	NewContract:  burn_mint_token_pool.NewBurnMintTokenPool,
	CallContract: func(tokenPool *burn_mint_token_pool.BurnMintTokenPool, opts *bind.CallOpts, args any) (bool, error) {
		return tokenPool.GetAllowListEnabled(opts)
	},
})

var GetAllowList = contract.NewRead(contract.ReadParams[any, []common.Address, *burn_mint_token_pool.BurnMintTokenPool]{
	Name:         "token-pool:get-allow-list",
	Version:      Version,
	Description:  "Gets the allow list on a TokenPool",
	ContractType: ContractType,
	NewContract:  burn_mint_token_pool.NewBurnMintTokenPool,
	CallContract: func(tokenPool *burn_mint_token_pool.BurnMintTokenPool, opts *bind.CallOpts, args any) ([]common.Address, error) {
		return tokenPool.GetAllowList(opts)
	},
})

var GetRouter = contract.NewRead(contract.ReadParams[any, common.Address, *burn_mint_token_pool.BurnMintTokenPool]{
	Name:         "token-pool:get-router",
	Version:      Version,
	Description:  "Gets the router address on a TokenPool",
	ContractType: ContractType,
	NewContract:  burn_mint_token_pool.NewBurnMintTokenPool,
	CallContract: func(tokenPool *burn_mint_token_pool.BurnMintTokenPool, opts *bind.CallOpts, args any) (common.Address, error) {
		return tokenPool.GetRouter(opts)
	},
})

var GetRateLimitAdmin = contract.NewRead(contract.ReadParams[any, common.Address, *burn_mint_token_pool.BurnMintTokenPool]{
	Name:         "token-pool:get-rate-limit-admin",
	Version:      Version,
	Description:  "Gets the rate limit admin address on a TokenPool",
	ContractType: ContractType,
	NewContract:  burn_mint_token_pool.NewBurnMintTokenPool,
	CallContract: func(tokenPool *burn_mint_token_pool.BurnMintTokenPool, opts *bind.CallOpts, args any) (common.Address, error) {
		return tokenPool.GetRateLimitAdmin(opts)
	},
})
