// Package lock_release_token_pool exposes the v1.4.0 LockReleaseTokenPool.
//
// It shares the whole TokenPool surface with BurnMintTokenPool (see the
// token_pool package) and adds the liquidity and rebalancer functions below.
// Like BurnMintTokenPool it is deployed behind a separate Proxy contract.
package lock_release_token_pool

import (
	"math/big"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_4_0/lock_release_token_pool"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations/contract"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

var ContractType cldf_deployment.ContractType = "LockReleaseTokenPool"
var Version *semver.Version = semver.MustParse("1.4.0")
var TypeAndVersion = cldf_deployment.NewTypeAndVersion(ContractType, *Version)

var TypeAndVersionRead = contract.NewRead(contract.ReadParams[struct{}, string, *lock_release_token_pool.LockReleaseTokenPool]{
	Name:         "lock-release-token-pool:type-and-version",
	Version:      Version,
	Description:  "Gets the type and version of a v1.4.0 LockReleaseTokenPool",
	ContractType: ContractType,
	NewContract:  lock_release_token_pool.NewLockReleaseTokenPool,
	CallContract: func(pool *lock_release_token_pool.LockReleaseTokenPool, opts *bind.CallOpts, _ struct{}) (string, error) {
		return pool.TypeAndVersion(opts)
	},
})

var GetRebalancer = contract.NewRead(contract.ReadParams[struct{}, common.Address, *lock_release_token_pool.LockReleaseTokenPool]{
	Name:         "lock-release-token-pool:get-rebalancer",
	Version:      Version,
	Description:  "Gets the rebalancer of a v1.4.0 LockReleaseTokenPool",
	ContractType: ContractType,
	NewContract:  lock_release_token_pool.NewLockReleaseTokenPool,
	CallContract: func(pool *lock_release_token_pool.LockReleaseTokenPool, opts *bind.CallOpts, _ struct{}) (common.Address, error) {
		return pool.GetRebalancer(opts)
	},
})

var GetRateLimitAdmin = contract.NewRead(contract.ReadParams[struct{}, common.Address, *lock_release_token_pool.LockReleaseTokenPool]{
	Name:         "lock-release-token-pool:get-rate-limit-admin",
	Version:      Version,
	Description:  "Gets the rate limit admin of a v1.4.0 LockReleaseTokenPool",
	ContractType: ContractType,
	NewContract:  lock_release_token_pool.NewLockReleaseTokenPool,
	CallContract: func(pool *lock_release_token_pool.LockReleaseTokenPool, opts *bind.CallOpts, _ struct{}) (common.Address, error) {
		return pool.GetRateLimitAdmin(opts)
	},
})

var CanAcceptLiquidity = contract.NewRead(contract.ReadParams[struct{}, bool, *lock_release_token_pool.LockReleaseTokenPool]{
	Name:         "lock-release-token-pool:can-accept-liquidity",
	Version:      Version,
	Description:  "Reports whether a v1.4.0 LockReleaseTokenPool accepts liquidity",
	ContractType: ContractType,
	NewContract:  lock_release_token_pool.NewLockReleaseTokenPool,
	CallContract: func(pool *lock_release_token_pool.LockReleaseTokenPool, opts *bind.CallOpts, _ struct{}) (bool, error) {
		return pool.CanAcceptLiquidity(opts)
	},
})

// SetRebalancerArgs is the input for SetRebalancer.
type SetRebalancerArgs struct {
	Rebalancer common.Address
}

var SetRebalancer = contract.NewWrite(contract.WriteParams[SetRebalancerArgs, *lock_release_token_pool.LockReleaseTokenPool]{
	Name:            "lock-release-token-pool:set-rebalancer",
	Version:         Version,
	Description:     "Sets the rebalancer on a v1.4.0 LockReleaseTokenPool",
	ContractType:    ContractType,
	ContractABI:     lock_release_token_pool.LockReleaseTokenPoolABI,
	NewContract:     lock_release_token_pool.NewLockReleaseTokenPool,
	IsAllowedCaller: contract.OnlyOwner[*lock_release_token_pool.LockReleaseTokenPool, SetRebalancerArgs],
	Validate:        func(_ SetRebalancerArgs) error { return nil },
	CallContract: func(pool *lock_release_token_pool.LockReleaseTokenPool, opts *bind.TransactOpts, args SetRebalancerArgs) (*types.Transaction, error) {
		return pool.SetRebalancer(opts, args.Rebalancer)
	},
})

// SetRateLimitAdminArgs is the input for SetRateLimitAdmin.
type SetRateLimitAdminArgs struct {
	RateLimitAdmin common.Address
}

var SetRateLimitAdmin = contract.NewWrite(contract.WriteParams[SetRateLimitAdminArgs, *lock_release_token_pool.LockReleaseTokenPool]{
	Name:            "lock-release-token-pool:set-rate-limit-admin",
	Version:         Version,
	Description:     "Sets the rate limit admin on a v1.4.0 LockReleaseTokenPool",
	ContractType:    ContractType,
	ContractABI:     lock_release_token_pool.LockReleaseTokenPoolABI,
	NewContract:     lock_release_token_pool.NewLockReleaseTokenPool,
	IsAllowedCaller: contract.OnlyOwner[*lock_release_token_pool.LockReleaseTokenPool, SetRateLimitAdminArgs],
	Validate:        func(_ SetRateLimitAdminArgs) error { return nil },
	CallContract: func(pool *lock_release_token_pool.LockReleaseTokenPool, opts *bind.TransactOpts, args SetRateLimitAdminArgs) (*types.Transaction, error) {
		return pool.SetRateLimitAdmin(opts, args.RateLimitAdmin)
	},
})

// ProvideLiquidityArgs is the input for ProvideLiquidity.
type ProvideLiquidityArgs struct {
	Amount *big.Int
}

// ProvideLiquidity is gated on the rebalancer, not the owner, on v1.4.0.
var ProvideLiquidity = contract.NewWrite(contract.WriteParams[ProvideLiquidityArgs, *lock_release_token_pool.LockReleaseTokenPool]{
	Name:            "lock-release-token-pool:provide-liquidity",
	Version:         Version,
	Description:     "Provides liquidity to a v1.4.0 LockReleaseTokenPool",
	ContractType:    ContractType,
	ContractABI:     lock_release_token_pool.LockReleaseTokenPoolABI,
	NewContract:     lock_release_token_pool.NewLockReleaseTokenPool,
	IsAllowedCaller: contract.AllCallersAllowed[*lock_release_token_pool.LockReleaseTokenPool, ProvideLiquidityArgs],
	Validate:        func(_ ProvideLiquidityArgs) error { return nil },
	CallContract: func(pool *lock_release_token_pool.LockReleaseTokenPool, opts *bind.TransactOpts, args ProvideLiquidityArgs) (*types.Transaction, error) {
		return pool.ProvideLiquidity(opts, args.Amount)
	},
})

// WithdrawLiquidityArgs is the input for WithdrawLiquidity.
type WithdrawLiquidityArgs struct {
	Amount *big.Int
}

// WithdrawLiquidity is gated on the rebalancer, not the owner, on v1.4.0.
var WithdrawLiquidity = contract.NewWrite(contract.WriteParams[WithdrawLiquidityArgs, *lock_release_token_pool.LockReleaseTokenPool]{
	Name:            "lock-release-token-pool:withdraw-liquidity",
	Version:         Version,
	Description:     "Withdraws liquidity from a v1.4.0 LockReleaseTokenPool",
	ContractType:    ContractType,
	ContractABI:     lock_release_token_pool.LockReleaseTokenPoolABI,
	NewContract:     lock_release_token_pool.NewLockReleaseTokenPool,
	IsAllowedCaller: contract.AllCallersAllowed[*lock_release_token_pool.LockReleaseTokenPool, WithdrawLiquidityArgs],
	Validate:        func(_ WithdrawLiquidityArgs) error { return nil },
	CallContract: func(pool *lock_release_token_pool.LockReleaseTokenPool, opts *bind.TransactOpts, args WithdrawLiquidityArgs) (*types.Transaction, error) {
		return pool.WithdrawLiquidity(opts, args.Amount)
	},
})
