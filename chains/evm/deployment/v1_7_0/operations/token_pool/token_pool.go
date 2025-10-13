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

var SetDynamicConfig = contract.NewWrite(contract.WriteParams[DynamicConfigArgs, *token_pool.TokenPool]{
	Name:            "token-pool:set-dynamic-config",
	Version:         semver.MustParse("1.7.0"),
	Description:     "Sets the dynamic configuration (router, threshold) for a TokenPool",
	ContractType:    ContractType,
	ContractABI:     token_pool.TokenPoolABI,
	NewContract:     token_pool.NewTokenPool,
	IsAllowedCaller: contract.OnlyOwner[*token_pool.TokenPool, DynamicConfigArgs],
	Validate:        func(DynamicConfigArgs) error { return nil },
	CallContract: func(tokenPool *token_pool.TokenPool, opts *bind.TransactOpts, args DynamicConfigArgs) (*types.Transaction, error) {
		return tokenPool.SetDynamicConfig(opts, args.Router, args.ThresholdAmountForAdditionalCCVs)
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

var GetDynamicConfig = contract.NewRead(contract.ReadParams[any, token_pool.GetDynamicConfig, *token_pool.TokenPool]{
	Name:         "token-pool:get-dynamic-config",
	Version:      semver.MustParse("1.7.0"),
	Description:  "Gets the router and threshold dynamic config on a TokenPool",
	ContractType: ContractType,
	NewContract:  token_pool.NewTokenPool,
	CallContract: func(tokenPool *token_pool.TokenPool, opts *bind.CallOpts, args any) (token_pool.GetDynamicConfig, error) {
		return tokenPool.GetDynamicConfig(opts)
	},
})

var GetRmnProxy = contract.NewRead(contract.ReadParams[any, common.Address, *token_pool.TokenPool]{
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
