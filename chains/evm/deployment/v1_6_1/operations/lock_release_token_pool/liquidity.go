package lock_release_token_pool

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
)

// WithdrawLiquidity requires custom access control (rebalancer-only) that the
// operations generator does not support, so it is defined here manually.
var WithdrawLiquidity = contract.NewWrite(contract.WriteParams[*big.Int, *LockReleaseTokenPoolContract]{
	Name:         "lock-release-token-pool:withdraw-liquidity",
	Version:      Version,
	Description:  "Withdraws liquidity from the LockReleaseTokenPool to the rebalancer",
	ContractType: ContractType,
	ContractABI:  LockReleaseTokenPoolABI,
	NewContract:  NewLockReleaseTokenPoolContract,
	IsAllowedCaller: func(c *LockReleaseTokenPoolContract, opts *bind.CallOpts, caller common.Address, _ *big.Int) (bool, error) {
		rebalancer, err := c.GetRebalancer(opts)
		if err != nil {
			return false, err
		}
		return rebalancer == caller, nil
	},
	Validate: func(amount *big.Int) error { return nil },
	CallContract: func(c *LockReleaseTokenPoolContract, opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
		return c.contract.Transact(opts, "withdrawLiquidity", amount)
	},
})
