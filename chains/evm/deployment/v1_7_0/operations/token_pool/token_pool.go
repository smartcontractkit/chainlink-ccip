package token_pool

import (
	"math/big"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/operations/burn_mint_token_pool"
	bnm_bindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/latest/burn_mint_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/latest/token_pool"
	"github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

var ContractType cldf_deployment.ContractType = "TokenPool"

// ConstructorArgs are the arguments for the TokenPool constructor.
// They apply to the basic token pool types: BurnMint variants and LockRelease pools.
type ConstructorArgs struct {
	Token              common.Address
	LocalTokenDecimals uint8
	Allowlist          []common.Address
	RMNProxy           common.Address
	Router             common.Address
}

type ChainUpdate struct {
	RemoteChainSelector       uint64
	RemotePoolAddresses       [][]byte
	RemoteTokenAddress        []byte
	OutboundRateLimiterConfig tokens.RateLimiterConfig
	InboundRateLimiterConfig  tokens.RateLimiterConfig
}

type CCVConfigArg = token_pool.TokenPoolCCVConfigArg

type ApplyChainUpdatesArgs struct {
	RemoteChainSelectorsToRemove []uint64
	ChainsToAdd                  []ChainUpdate
}

type RemotePoolArgs struct {
	RemoteChainSelector uint64
	RemotePoolAddress   []byte
}

type RateLimiterBucket = token_pool.RateLimiterTokenBucket
type DynamicConfig = token_pool.GetDynamicConfig

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
	Router                           common.Address
	ThresholdAmountForAdditionalCCVs *big.Int
}

var Deploy = contract.NewDeploy(contract.DeployParams[ConstructorArgs]{
	Name:             "token-pool:deploy",
	Version:          semver.MustParse("1.7.0"),
	Description:      "Deploys various TokenPool contracts (i.e. BurnMint, LockRelease)",
	ContractMetadata: bnm_bindings.BurnMintTokenPoolMetaData, // Just to get the expected constructor args
	BytecodeByTypeAndVersion: map[string]contract.Bytecode{
		cldf_deployment.NewTypeAndVersion(burn_mint_token_pool.ContractType, *semver.MustParse("1.7.0")).String(): {
			EVM: common.FromHex(bnm_bindings.BurnMintTokenPoolBin),
		},
	},
	Validate: func(ConstructorArgs) error { return nil },
})

var ApplyChainUpdates = contract.NewWrite(contract.WriteParams[ApplyChainUpdatesArgs, *token_pool.TokenPool]{
	Name:            "token-pool:apply-chain-updates",
	Version:         semver.MustParse("1.7.0"),
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

var ApplyCCVConfigUpdates = contract.NewWrite(contract.WriteParams[[]CCVConfigArg, *token_pool.TokenPool]{
	Name:            "token-pool:apply-ccv-config-updates",
	Version:         semver.MustParse("1.7.0"),
	Description:     "Applies CCV config updates for remote chains to a TokenPool",
	ContractType:    ContractType,
	ContractABI:     token_pool.TokenPoolABI,
	NewContract:     token_pool.NewTokenPool,
	IsAllowedCaller: contract.OnlyOwner[*token_pool.TokenPool, []CCVConfigArg],
	Validate:        func([]CCVConfigArg) error { return nil },
	CallContract: func(tokenPool *token_pool.TokenPool, opts *bind.TransactOpts, args []CCVConfigArg) (*types.Transaction, error) {
		return tokenPool.ApplyCCVConfigUpdates(opts, args)
	},
})

var AddRemotePool = contract.NewWrite(contract.WriteParams[RemotePoolArgs, *token_pool.TokenPool]{
	Name:            "token-pool:add-remote-pool",
	Version:         semver.MustParse("1.7.0"),
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
	Version:         semver.MustParse("1.7.0"),
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

var SetRateLimitAdmin = contract.NewWrite(contract.WriteParams[common.Address, *token_pool.TokenPool]{
	Name:            "token-pool:set-rate-limit-admin",
	Version:         semver.MustParse("1.7.0"),
	Description:     "Sets the rate limit admin (additional address allowed to set rate limits) for a TokenPool",
	ContractType:    ContractType,
	ContractABI:     token_pool.TokenPoolABI,
	NewContract:     token_pool.NewTokenPool,
	IsAllowedCaller: contract.OnlyOwner[*token_pool.TokenPool, common.Address],
	Validate:        func(common.Address) error { return nil },
	CallContract: func(tokenPool *token_pool.TokenPool, opts *bind.TransactOpts, args common.Address) (*types.Transaction, error) {
		return tokenPool.SetRateLimitAdmin(opts, args)
	},
})

var SetChainRateLimiterConfigs = contract.NewWrite(contract.WriteParams[[]SetChainRateLimiterConfigArg, *token_pool.TokenPool]{
	Name:         "token-pool:set-chain-rate-limiter-configs",
	Version:      semver.MustParse("1.7.0"),
	Description:  "Sets the rate limiter configs for existing remote chains on a TokenPool",
	ContractType: ContractType,
	ContractABI:  token_pool.TokenPoolABI,
	NewContract:  token_pool.NewTokenPool,
	IsAllowedCaller: func(contract *token_pool.TokenPool, opts *bind.CallOpts, caller common.Address, args []SetChainRateLimiterConfigArg) (bool, error) {
		owner, err := contract.Owner(opts)
		if err != nil {
			return false, err
		}
		rateLimitAdmin, err := contract.GetRateLimitAdmin(opts)
		if err != nil {
			return false, err
		}
		return caller == owner || caller == rateLimitAdmin, nil
	},
	Validate: func([]SetChainRateLimiterConfigArg) error { return nil },
	CallContract: func(tokenPool *token_pool.TokenPool, opts *bind.TransactOpts, args []SetChainRateLimiterConfigArg) (*types.Transaction, error) {
		var remoteChainSelectors []uint64
		var inboundConfigs []token_pool.RateLimiterConfig
		var outboundConfigs []token_pool.RateLimiterConfig
		for _, arg := range args {
			remoteChainSelectors = append(remoteChainSelectors, arg.RemoteChainSelector)
			inboundConfigs = append(inboundConfigs, token_pool.RateLimiterConfig(arg.InboundRateLimiterConfig))
			outboundConfigs = append(outboundConfigs, token_pool.RateLimiterConfig(arg.OutboundRateLimiterConfig))
		}
		return tokenPool.SetChainRateLimiterConfigs(opts, remoteChainSelectors, outboundConfigs, inboundConfigs)
	},
})

var SetDynamicConfig = contract.NewWrite(contract.WriteParams[DynamicConfigArgs, *token_pool.TokenPool]{
	Name:            "token-pool:set-dynamic-config",
	Version:         semver.MustParse("1.7.0"),
	Description:     "Sets the router and CCV threshold configuration for a TokenPool",
	ContractType:    ContractType,
	ContractABI:     token_pool.TokenPoolABI,
	NewContract:     token_pool.NewTokenPool,
	IsAllowedCaller: contract.OnlyOwner[*token_pool.TokenPool, DynamicConfigArgs],
	Validate:        func(DynamicConfigArgs) error { return nil },
	CallContract: func(tokenPool *token_pool.TokenPool, opts *bind.TransactOpts, args DynamicConfigArgs) (*types.Transaction, error) {
		threshold := args.ThresholdAmountForAdditionalCCVs
		if threshold == nil {
			currentConfig, err := tokenPool.GetDynamicConfig(&bind.CallOpts{
				Context: opts.Context,
				From:    opts.From,
			})
			if err != nil {
				return nil, err
			}
			threshold = currentConfig.ThresholdAmountForAdditionalCCVs
			if threshold == nil {
				threshold = big.NewInt(0)
			}
		}
		return tokenPool.SetDynamicConfig(opts, args.Router, threshold)
	},
})

var ApplyAllowListUpdates = contract.NewWrite(contract.WriteParams[ApplyAllowListUpdatesArgs, *token_pool.TokenPool]{
	Name:            "token-pool:apply-allowlist-updates",
	Version:         semver.MustParse("1.7.0"),
	Description:     "Applies allowlist updates to a TokenPool",
	ContractType:    ContractType,
	ContractABI:     token_pool.TokenPoolABI,
	NewContract:     token_pool.NewTokenPool,
	IsAllowedCaller: contract.OnlyOwner[*token_pool.TokenPool, ApplyAllowListUpdatesArgs],
	Validate:        func(ApplyAllowListUpdatesArgs) error { return nil },
	CallContract: func(tokenPool *token_pool.TokenPool, opts *bind.TransactOpts, args ApplyAllowListUpdatesArgs) (*types.Transaction, error) {
		return tokenPool.ApplyAllowListUpdates(opts, args.Removes, args.Adds)
	},
})

var GetAllowListEnabled = contract.NewRead(contract.ReadParams[any, bool, *token_pool.TokenPool]{
	Name:         "token-pool:get-allowlist-enabled",
	Version:      semver.MustParse("1.7.0"),
	Description:  "Gets whether the allowlist is enabled on a TokenPool",
	ContractType: ContractType,
	NewContract:  token_pool.NewTokenPool,
	CallContract: func(tokenPool *token_pool.TokenPool, opts *bind.CallOpts, args any) (bool, error) {
		return tokenPool.GetAllowListEnabled(opts)
	},
})

var GetAllowList = contract.NewRead(contract.ReadParams[any, []common.Address, *token_pool.TokenPool]{
	Name:         "token-pool:get-allowlist",
	Version:      semver.MustParse("1.7.0"),
	Description:  "Gets the allowlist on a TokenPool",
	ContractType: ContractType,
	NewContract:  token_pool.NewTokenPool,
	CallContract: func(tokenPool *token_pool.TokenPool, opts *bind.CallOpts, args any) ([]common.Address, error) {
		return tokenPool.GetAllowList(opts)
	},
})

var GetDynamicConfig = contract.NewRead(contract.ReadParams[any, DynamicConfig, *token_pool.TokenPool]{
	Name:         "token-pool:get-dynamic-config",
	Version:      semver.MustParse("1.7.0"),
	Description:  "Gets the router and CCV threshold configuration on a TokenPool",
	ContractType: ContractType,
	NewContract:  token_pool.NewTokenPool,
	CallContract: func(tokenPool *token_pool.TokenPool, opts *bind.CallOpts, args any) (DynamicConfig, error) {
		return tokenPool.GetDynamicConfig(opts)
	},
})

var GetRMNProxy = contract.NewRead(contract.ReadParams[any, common.Address, *token_pool.TokenPool]{
	Name:         "token-pool:get-rmn-proxy",
	Version:      semver.MustParse("1.7.0"),
	Description:  "Gets the RMN proxy address on a TokenPool",
	ContractType: ContractType,
	NewContract:  token_pool.NewTokenPool,
	CallContract: func(tokenPool *token_pool.TokenPool, opts *bind.CallOpts, args any) (common.Address, error) {
		return tokenPool.GetRmnProxy(opts)
	},
})

var GetRateLimitAdmin = contract.NewRead(contract.ReadParams[any, common.Address, *token_pool.TokenPool]{
	Name:         "token-pool:get-rate-limit-admin",
	Version:      semver.MustParse("1.7.0"),
	Description:  "Gets the rate limit admin address on a TokenPool",
	ContractType: ContractType,
	NewContract:  token_pool.NewTokenPool,
	CallContract: func(tokenPool *token_pool.TokenPool, opts *bind.CallOpts, args any) (common.Address, error) {
		return tokenPool.GetRateLimitAdmin(opts)
	},
})

var GetCurrentInboundRateLimiterState = contract.NewRead(contract.ReadParams[uint64, RateLimiterBucket, *token_pool.TokenPool]{
	Name:         "token-pool:get-current-inbound-rate-limiter-state",
	Version:      semver.MustParse("1.7.0"),
	Description:  "Gets the current inbound rate limiter state for a given remote chain selector on a TokenPool",
	ContractType: ContractType,
	NewContract:  token_pool.NewTokenPool,
	CallContract: func(tokenPool *token_pool.TokenPool, opts *bind.CallOpts, args uint64) (RateLimiterBucket, error) {
		return tokenPool.GetCurrentInboundRateLimiterState(opts, args)
	},
})

var GetCurrentOutboundRateLimiterState = contract.NewRead(contract.ReadParams[uint64, RateLimiterBucket, *token_pool.TokenPool]{
	Name:         "token-pool:get-current-outbound-rate-limiter-state",
	Version:      semver.MustParse("1.7.0"),
	Description:  "Gets the current outbound rate limiter state for a given remote chain selector on a TokenPool",
	ContractType: ContractType,
	NewContract:  token_pool.NewTokenPool,
	CallContract: func(tokenPool *token_pool.TokenPool, opts *bind.CallOpts, args uint64) (RateLimiterBucket, error) {
		return tokenPool.GetCurrentOutboundRateLimiterState(opts, args)
	},
})

var GetSupportedChains = contract.NewRead(contract.ReadParams[any, []uint64, *token_pool.TokenPool]{
	Name:         "token-pool:supported-chains",
	Version:      semver.MustParse("1.7.0"),
	Description:  "Gets the list of supported remote chain selectors on a TokenPool",
	ContractType: ContractType,
	NewContract:  token_pool.NewTokenPool,
	CallContract: func(tokenPool *token_pool.TokenPool, opts *bind.CallOpts, args any) ([]uint64, error) {
		return tokenPool.GetSupportedChains(opts)
	},
})

var GetRemotePools = contract.NewRead(contract.ReadParams[uint64, [][]byte, *token_pool.TokenPool]{
	Name:         "token-pool:get-remote-pools",
	Version:      semver.MustParse("1.7.0"),
	Description:  "Gets the remote pool addresses for a given remote chain selector on a TokenPool",
	ContractType: ContractType,
	NewContract:  token_pool.NewTokenPool,
	CallContract: func(tokenPool *token_pool.TokenPool, opts *bind.CallOpts, args uint64) ([][]byte, error) {
		return tokenPool.GetRemotePools(opts, args)
	},
})

var GetRemoteToken = contract.NewRead(contract.ReadParams[uint64, []byte, *token_pool.TokenPool]{
	Name:         "token-pool:get-remote-token",
	Version:      semver.MustParse("1.7.0"),
	Description:  "Gets the remote pool address for a given remote chain selector on a TokenPool",
	ContractType: ContractType,
	NewContract:  token_pool.NewTokenPool,
	CallContract: func(tokenPool *token_pool.TokenPool, opts *bind.CallOpts, args uint64) ([]byte, error) {
		return tokenPool.GetRemoteToken(opts, args)
	},
})

var GetToken = contract.NewRead(contract.ReadParams[any, common.Address, *token_pool.TokenPool]{
	Name:         "token-pool:get-token",
	Version:      semver.MustParse("1.7.0"),
	Description:  "Gets the local token address for a TokenPool",
	ContractType: ContractType,
	NewContract:  token_pool.NewTokenPool,
	CallContract: func(tokenPool *token_pool.TokenPool, opts *bind.CallOpts, args any) (common.Address, error) {
		return tokenPool.GetToken(opts)
	},
})
