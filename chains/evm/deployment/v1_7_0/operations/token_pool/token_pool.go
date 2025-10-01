package token_pool

import (
	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/operations/burn_mint_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/latest/burn_mint_token_pool_v2"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/latest/token_pool_v2"
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

type ChainUpdate = token_pool_v2.TokenPoolChainUpdate

type CCVConfigArg = token_pool_v2.TokenPoolV2CCVConfigArg

type ApplyChainUpdatesArgs struct {
	RemoteChainSelectorsToRemove []uint64
	ChainsToAdd                  []ChainUpdate
}

type RemotePoolArgs struct {
	RemoteChainSelector uint64
	RemotePoolAddress   []byte
}

type RateLimiterBucket = token_pool_v2.RateLimiterTokenBucket

type RateLimiterConfig = token_pool_v2.RateLimiterConfig

type SetChainRateLimiterConfigArg struct {
	RemoteChainSelector       uint64
	InboundRateLimiterConfig  RateLimiterConfig
	OutboundRateLimiterConfig RateLimiterConfig
}

type ApplyAllowListUpdatesArgs struct {
	Adds    []common.Address
	Removes []common.Address
}

var Deploy = contract.NewDeploy(contract.DeployParams[ConstructorArgs]{
	Name:             "token-pool:deploy",
	Version:          semver.MustParse("1.7.0"),
	Description:      "Deploys various TokenPool contracts (i.e. BurnMint, LockRelease)",
	ContractMetadata: token_pool_v2.TokenPoolV2MetaData,
	BytecodeByTypeAndVersion: map[string]contract.Bytecode{
		cldf_deployment.NewTypeAndVersion(burn_mint_token_pool.ContractType, *semver.MustParse("1.7.0")).String(): {
			EVM: common.FromHex(burn_mint_token_pool_v2.BurnMintTokenPoolV2Bin),
		},
	},
	Validate: func(ConstructorArgs) error { return nil },
})

var ApplyChainUpdates = contract.NewWrite(contract.WriteParams[ApplyChainUpdatesArgs, *token_pool_v2.TokenPoolV2]{
	Name:            "token-pool:apply-chain-updates",
	Version:         semver.MustParse("1.7.0"),
	Description:     "Applies chain updates to a TokenPool, enabling / disabling remote chains and setting rate limits",
	ContractType:    ContractType,
	ContractABI:     token_pool_v2.TokenPoolV2ABI,
	NewContract:     token_pool_v2.NewTokenPoolV2,
	IsAllowedCaller: contract.OnlyOwner[*token_pool_v2.TokenPoolV2, ApplyChainUpdatesArgs],
	Validate:        func(ApplyChainUpdatesArgs) error { return nil },
	CallContract: func(tokenPoolV2 *token_pool_v2.TokenPoolV2, opts *bind.TransactOpts, args ApplyChainUpdatesArgs) (*types.Transaction, error) {
		return tokenPoolV2.ApplyChainUpdates(opts, args.RemoteChainSelectorsToRemove, args.ChainsToAdd)
	},
})

var ApplyCCVConfigUpdates = contract.NewWrite(contract.WriteParams[[]CCVConfigArg, *token_pool_v2.TokenPoolV2]{
	Name:            "token-pool:apply-ccv-config-updates",
	Version:         semver.MustParse("1.7.0"),
	Description:     "Applies CCV config updates for remote chains to a TokenPool",
	ContractType:    ContractType,
	ContractABI:     token_pool_v2.TokenPoolV2ABI,
	NewContract:     token_pool_v2.NewTokenPoolV2,
	IsAllowedCaller: contract.OnlyOwner[*token_pool_v2.TokenPoolV2, []CCVConfigArg],
	Validate:        func([]CCVConfigArg) error { return nil },
	CallContract: func(tokenPoolV2 *token_pool_v2.TokenPoolV2, opts *bind.TransactOpts, args []CCVConfigArg) (*types.Transaction, error) {
		return tokenPoolV2.ApplyCCVConfigUpdates(opts, args)
	},
})

var AddRemotePool = contract.NewWrite(contract.WriteParams[RemotePoolArgs, *token_pool_v2.TokenPoolV2]{
	Name:            "token-pool:add-remote-pool",
	Version:         semver.MustParse("1.7.0"),
	Description:     "Adds a remote pool for a given chain selector to a TokenPool",
	ContractType:    ContractType,
	ContractABI:     token_pool_v2.TokenPoolV2ABI,
	NewContract:     token_pool_v2.NewTokenPoolV2,
	IsAllowedCaller: contract.OnlyOwner[*token_pool_v2.TokenPoolV2, RemotePoolArgs],
	Validate:        func(RemotePoolArgs) error { return nil },
	CallContract: func(tokenPoolV2 *token_pool_v2.TokenPoolV2, opts *bind.TransactOpts, args RemotePoolArgs) (*types.Transaction, error) {
		return tokenPoolV2.AddRemotePool(opts, args.RemoteChainSelector, args.RemotePoolAddress)
	},
})

var RemoveRemotePool = contract.NewWrite(contract.WriteParams[RemotePoolArgs, *token_pool_v2.TokenPoolV2]{
	Name:            "token-pool:remove-remote-pool",
	Version:         semver.MustParse("1.7.0"),
	Description:     "Removes a remote pool for a given chain selector from a TokenPool",
	ContractType:    ContractType,
	ContractABI:     token_pool_v2.TokenPoolV2ABI,
	NewContract:     token_pool_v2.NewTokenPoolV2,
	IsAllowedCaller: contract.OnlyOwner[*token_pool_v2.TokenPoolV2, RemotePoolArgs],
	Validate:        func(RemotePoolArgs) error { return nil },
	CallContract: func(tokenPoolV2 *token_pool_v2.TokenPoolV2, opts *bind.TransactOpts, args RemotePoolArgs) (*types.Transaction, error) {
		return tokenPoolV2.RemoveRemotePool(opts, args.RemoteChainSelector, args.RemotePoolAddress)
	},
})

var SetRateLimitAdmin = contract.NewWrite(contract.WriteParams[common.Address, *token_pool_v2.TokenPoolV2]{
	Name:            "token-pool:set-rate-limit-admin",
	Version:         semver.MustParse("1.7.0"),
	Description:     "Sets the rate limit admin (additional address allowed to set rate limits) for a TokenPool",
	ContractType:    ContractType,
	ContractABI:     token_pool_v2.TokenPoolV2ABI,
	NewContract:     token_pool_v2.NewTokenPoolV2,
	IsAllowedCaller: contract.OnlyOwner[*token_pool_v2.TokenPoolV2, common.Address],
	Validate:        func(common.Address) error { return nil },
	CallContract: func(tokenPoolV2 *token_pool_v2.TokenPoolV2, opts *bind.TransactOpts, args common.Address) (*types.Transaction, error) {
		return tokenPoolV2.SetRateLimitAdmin(opts, args)
	},
})

var SetChainRateLimiterConfigs = contract.NewWrite(contract.WriteParams[[]SetChainRateLimiterConfigArg, *token_pool_v2.TokenPoolV2]{
	Name:         "token-pool:set-chain-rate-limiter-configs",
	Version:      semver.MustParse("1.7.0"),
	Description:  "Sets the rate limiter configs for existing remote chains on a TokenPool",
	ContractType: ContractType,
	ContractABI:  token_pool_v2.TokenPoolV2ABI,
	NewContract:  token_pool_v2.NewTokenPoolV2,
	IsAllowedCaller: func(contract *token_pool_v2.TokenPoolV2, opts *bind.CallOpts, caller common.Address, args []SetChainRateLimiterConfigArg) (bool, error) {
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
	CallContract: func(tokenPoolV2 *token_pool_v2.TokenPoolV2, opts *bind.TransactOpts, args []SetChainRateLimiterConfigArg) (*types.Transaction, error) {
		var remoteChainSelectors []uint64
		var inboundConfigs []RateLimiterConfig
		var outboundConfigs []RateLimiterConfig
		for _, arg := range args {
			remoteChainSelectors = append(remoteChainSelectors, arg.RemoteChainSelector)
			inboundConfigs = append(inboundConfigs, arg.InboundRateLimiterConfig)
			outboundConfigs = append(outboundConfigs, arg.OutboundRateLimiterConfig)
		}
		return tokenPoolV2.SetChainRateLimiterConfigs(opts, remoteChainSelectors, inboundConfigs, outboundConfigs)
	},
})

var SetRouter = contract.NewWrite(contract.WriteParams[common.Address, *token_pool_v2.TokenPoolV2]{
	Name:            "token-pool:set-router",
	Version:         semver.MustParse("1.7.0"),
	Description:     "Sets the router address for a TokenPool",
	ContractType:    ContractType,
	ContractABI:     token_pool_v2.TokenPoolV2ABI,
	NewContract:     token_pool_v2.NewTokenPoolV2,
	IsAllowedCaller: contract.OnlyOwner[*token_pool_v2.TokenPoolV2, common.Address],
	Validate:        func(common.Address) error { return nil },
	CallContract: func(tokenPoolV2 *token_pool_v2.TokenPoolV2, opts *bind.TransactOpts, args common.Address) (*types.Transaction, error) {
		return tokenPoolV2.SetRouter(opts, args)
	},
})

var ApplyAllowListUpdates = contract.NewWrite(contract.WriteParams[ApplyAllowListUpdatesArgs, *token_pool_v2.TokenPoolV2]{
	Name:            "token-pool:apply-allowlist-updates",
	Version:         semver.MustParse("1.7.0"),
	Description:     "Applies allowlist updates to a TokenPool",
	ContractType:    ContractType,
	ContractABI:     token_pool_v2.TokenPoolV2ABI,
	NewContract:     token_pool_v2.NewTokenPoolV2,
	IsAllowedCaller: contract.OnlyOwner[*token_pool_v2.TokenPoolV2, ApplyAllowListUpdatesArgs],
	Validate:        func(ApplyAllowListUpdatesArgs) error { return nil },
	CallContract: func(tokenPoolV2 *token_pool_v2.TokenPoolV2, opts *bind.TransactOpts, args ApplyAllowListUpdatesArgs) (*types.Transaction, error) {
		return tokenPoolV2.ApplyAllowListUpdates(opts, args.Removes, args.Adds)
	},
})

var GetAllowListEnabled = contract.NewRead(contract.ReadParams[any, bool, *token_pool_v2.TokenPoolV2]{
	Name:         "token-pool:get-allowlist-enabled",
	Version:      semver.MustParse("1.7.0"),
	Description:  "Gets whether the allowlist is enabled on a TokenPool",
	ContractType: ContractType,
	NewContract:  token_pool_v2.NewTokenPoolV2,
	CallContract: func(tokenPoolV2 *token_pool_v2.TokenPoolV2, opts *bind.CallOpts, args any) (bool, error) {
		return tokenPoolV2.GetAllowListEnabled(opts)
	},
})

var GetAllowList = contract.NewRead(contract.ReadParams[any, []common.Address, *token_pool_v2.TokenPoolV2]{
	Name:         "token-pool:get-allowlist",
	Version:      semver.MustParse("1.7.0"),
	Description:  "Gets the allowlist on a TokenPool",
	ContractType: ContractType,
	NewContract:  token_pool_v2.NewTokenPoolV2,
	CallContract: func(tokenPoolV2 *token_pool_v2.TokenPoolV2, opts *bind.CallOpts, args any) ([]common.Address, error) {
		return tokenPoolV2.GetAllowList(opts)
	},
})

var GetRouter = contract.NewRead(contract.ReadParams[any, common.Address, *token_pool_v2.TokenPoolV2]{
	Name:         "token-pool:get-router",
	Version:      semver.MustParse("1.7.0"),
	Description:  "Gets the router address on a TokenPool",
	ContractType: ContractType,
	NewContract:  token_pool_v2.NewTokenPoolV2,
	CallContract: func(tokenPoolV2 *token_pool_v2.TokenPoolV2, opts *bind.CallOpts, args any) (common.Address, error) {
		return tokenPoolV2.GetRouter(opts)
	},
})

var GetRateLimitAdmin = contract.NewRead(contract.ReadParams[any, common.Address, *token_pool_v2.TokenPoolV2]{
	Name:         "token-pool:get-rate-limit-admin",
	Version:      semver.MustParse("1.7.0"),
	Description:  "Gets the rate limit admin address on a TokenPool",
	ContractType: ContractType,
	NewContract:  token_pool_v2.NewTokenPoolV2,
	CallContract: func(tokenPoolV2 *token_pool_v2.TokenPoolV2, opts *bind.CallOpts, args any) (common.Address, error) {
		return tokenPoolV2.GetRateLimitAdmin(opts)
	},
})

var GetCurrentInboundRateLimiterState = contract.NewRead(contract.ReadParams[uint64, RateLimiterBucket, *token_pool_v2.TokenPoolV2]{
	Name:         "token-pool:get-current-inbound-rate-limiter-state",
	Version:      semver.MustParse("1.7.0"),
	Description:  "Gets the current inbound rate limiter state for a given remote chain selector on a TokenPool",
	ContractType: ContractType,
	NewContract:  token_pool_v2.NewTokenPoolV2,
	CallContract: func(tokenPoolV2 *token_pool_v2.TokenPoolV2, opts *bind.CallOpts, args uint64) (RateLimiterBucket, error) {
		return tokenPoolV2.GetCurrentInboundRateLimiterState(opts, args)
	},
})

var GetCurrentOutboundRateLimiterState = contract.NewRead(contract.ReadParams[uint64, RateLimiterBucket, *token_pool_v2.TokenPoolV2]{
	Name:         "token-pool:get-current-outbound-rate-limiter-state",
	Version:      semver.MustParse("1.7.0"),
	Description:  "Gets the current outbound rate limiter state for a given remote chain selector on a TokenPool",
	ContractType: ContractType,
	NewContract:  token_pool_v2.NewTokenPoolV2,
	CallContract: func(tokenPoolV2 *token_pool_v2.TokenPoolV2, opts *bind.CallOpts, args uint64) (RateLimiterBucket, error) {
		return tokenPoolV2.GetCurrentOutboundRateLimiterState(opts, args)
	},
})

var GetSupportedChains = contract.NewRead(contract.ReadParams[any, []uint64, *token_pool_v2.TokenPoolV2]{
	Name:         "token-pool:supported-chains",
	Version:      semver.MustParse("1.7.0"),
	Description:  "Gets the list of supported remote chain selectors on a TokenPool",
	ContractType: ContractType,
	NewContract:  token_pool_v2.NewTokenPoolV2,
	CallContract: func(tokenPoolV2 *token_pool_v2.TokenPoolV2, opts *bind.CallOpts, args any) ([]uint64, error) {
		return tokenPoolV2.GetSupportedChains(opts)
	},
})

var GetRemotePools = contract.NewRead(contract.ReadParams[uint64, [][]byte, *token_pool_v2.TokenPoolV2]{
	Name:         "token-pool:get-remote-pools",
	Version:      semver.MustParse("1.7.0"),
	Description:  "Gets the remote pool address for a given remote chain selector on a TokenPool",
	ContractType: ContractType,
	NewContract:  token_pool_v2.NewTokenPoolV2,
	CallContract: func(tokenPoolV2 *token_pool_v2.TokenPoolV2, opts *bind.CallOpts, args uint64) ([][]byte, error) {
		return tokenPoolV2.GetRemotePools(opts, args)
	},
})

var GetRemoteToken = contract.NewRead(contract.ReadParams[uint64, []byte, *token_pool_v2.TokenPoolV2]{
	Name:         "token-pool:get-remote-token",
	Version:      semver.MustParse("1.7.0"),
	Description:  "Gets the remote token address for a given remote chain selector on a TokenPool",
	ContractType: ContractType,
	NewContract:  token_pool_v2.NewTokenPoolV2,
	CallContract: func(tokenPoolV2 *token_pool_v2.TokenPoolV2, opts *bind.CallOpts, args uint64) ([]byte, error) {
		return tokenPoolV2.GetRemoteToken(opts, args)
	},
})
