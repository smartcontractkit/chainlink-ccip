package hybrid_lock_release_usdc_token_pool

import (
	"fmt"
	"math/big"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_2/hybrid_lock_release_usdc_token_pool"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

var ContractType cldf_deployment.ContractType = "HybridLockReleaseUSDCTokenPool"

var Version = semver.MustParse("1.6.2")

type WithdrawLiquidityArgs struct {
	RemoteChainSelector uint64
	Amount              *big.Int
}

var GetLockedTokensForChain = contract.NewRead(contract.ReadParams[uint64, *big.Int, *hybrid_lock_release_usdc_token_pool.HybridLockReleaseUSDCTokenPool]{
	Name:         "hybrid-lock-release-usdc-token-pool:get-locked-tokens-for-chain",
	Version:      Version,
	Description:  "Gets locked token balance for a remote chain selector",
	ContractType: ContractType,
	NewContract:  hybrid_lock_release_usdc_token_pool.NewHybridLockReleaseUSDCTokenPool,
	CallContract: func(pool *hybrid_lock_release_usdc_token_pool.HybridLockReleaseUSDCTokenPool, opts *bind.CallOpts, remoteChainSelector uint64) (*big.Int, error) {
		return pool.GetLockedTokensForChain(opts, remoteChainSelector)
	},
})

var ShouldUseLockRelease = contract.NewRead(contract.ReadParams[uint64, bool, *hybrid_lock_release_usdc_token_pool.HybridLockReleaseUSDCTokenPool]{
	Name:         "hybrid-lock-release-usdc-token-pool:should-use-lock-release",
	Version:      Version,
	Description:  "Returns whether a remote chain selector should use lock-release",
	ContractType: ContractType,
	NewContract:  hybrid_lock_release_usdc_token_pool.NewHybridLockReleaseUSDCTokenPool,
	CallContract: func(pool *hybrid_lock_release_usdc_token_pool.HybridLockReleaseUSDCTokenPool, opts *bind.CallOpts, remoteChainSelector uint64) (bool, error) {
		return pool.ShouldUseLockRelease(opts, remoteChainSelector)
	},
})

var WithdrawLiquidity = contract.NewWrite(contract.WriteParams[WithdrawLiquidityArgs, *hybrid_lock_release_usdc_token_pool.HybridLockReleaseUSDCTokenPool]{
	Name:            "hybrid-lock-release-usdc-token-pool:withdraw-liquidity",
	Version:         Version,
	Description:     "Withdraws liquidity for a remote chain selector",
	ContractType:    ContractType,
	ContractABI:     hybrid_lock_release_usdc_token_pool.HybridLockReleaseUSDCTokenPoolABI,
	NewContract:     hybrid_lock_release_usdc_token_pool.NewHybridLockReleaseUSDCTokenPool,
	IsAllowedCaller: contract.OnlyOwner[*hybrid_lock_release_usdc_token_pool.HybridLockReleaseUSDCTokenPool, WithdrawLiquidityArgs],
	Validate: func(args WithdrawLiquidityArgs) error {
		if args.Amount == nil || args.Amount.Sign() <= 0 {
			return fmt.Errorf("amount must be greater than zero")
		}
		return nil
	},
	CallContract: func(pool *hybrid_lock_release_usdc_token_pool.HybridLockReleaseUSDCTokenPool, opts *bind.TransactOpts, args WithdrawLiquidityArgs) (*types.Transaction, error) {
		return pool.WithdrawLiquidity(opts, args.RemoteChainSelector, args.Amount)
	},
})
