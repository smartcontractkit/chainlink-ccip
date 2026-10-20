package lock_release_token_pool

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_4_0/lock_release_token_pool"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations/contract"
)

// WithdrawLiquidityProposalOnly forces the rebalancer-gated withdraw into an MCMS proposal. The
// on-chain gate is the rebalancer, which none of operations-gen's access aliases describe, so it
// is never executed directly with the deployer key.
var WithdrawLiquidityProposalOnly = contract.NewWrite(contract.WriteParams[*big.Int, *lock_release_token_pool.LockReleaseTokenPool]{
	Name:            "lock_release_token_pool:withdraw-liquidity-proposal-only",
	Version:         Version,
	Description:     "Calls withdrawLiquidity on the LockReleaseTokenPool 1.4.0 contract (proposal-only, never executed directly)",
	ContractType:    ContractType,
	ContractABI:     lock_release_token_pool.LockReleaseTokenPoolABI,
	NewContract:     lock_release_token_pool.NewLockReleaseTokenPool,
	IsAllowedCaller: contract.NoCallersAllowed[*lock_release_token_pool.LockReleaseTokenPool, *big.Int],
	Validate:        func(*big.Int) error { return nil },
	CallContract: func(c *lock_release_token_pool.LockReleaseTokenPool, opts *bind.TransactOpts, args *big.Int) (*types.Transaction, error) {
		return c.WithdrawLiquidity(opts, args)
	},
})
