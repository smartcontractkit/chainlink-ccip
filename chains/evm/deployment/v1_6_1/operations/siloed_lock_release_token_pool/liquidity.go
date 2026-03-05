package siloed_lock_release_token_pool

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
)

type WithdrawSiloedLiquidityInput struct {
	RemoteChainSelector uint64
	Amount              *big.Int
}

// WithdrawLiquidity requires custom access control (rebalancer-only) that the
// operations generator does not support, so it is defined here manually.
var WithdrawLiquidity = contract.NewWrite(contract.WriteParams[*big.Int, *SiloedLockReleaseTokenPoolContract]{
	Name:         "siloed-lock-release-token-pool:withdraw-liquidity",
	Version:      Version,
	Description:  "Withdraws unsiloed liquidity from the SiloedLockReleaseTokenPool to the rebalancer",
	ContractType: ContractType,
	ContractABI:  SiloedLockReleaseTokenPoolABI,
	NewContract:  NewSiloedLockReleaseTokenPoolContract,
	IsAllowedCaller: func(c *SiloedLockReleaseTokenPoolContract, opts *bind.CallOpts, caller common.Address, _ *big.Int) (bool, error) {
		rebalancer, err := c.GetRebalancer(opts)
		if err != nil {
			return false, err
		}
		return rebalancer == caller, nil
	},
	Validate: func(amount *big.Int) error { return nil },
	CallContract: func(c *SiloedLockReleaseTokenPoolContract, opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
		return c.contract.Transact(opts, "withdrawLiquidity", amount)
	},
})

// WithdrawSiloedLiquidity requires custom access control (chain-rebalancer-only) that the
// operations generator does not support, so it is defined here manually.
var WithdrawSiloedLiquidity = contract.NewWrite(contract.WriteParams[WithdrawSiloedLiquidityInput, *SiloedLockReleaseTokenPoolContract]{
	Name:         "siloed-lock-release-token-pool:withdraw-siloed-liquidity",
	Version:      Version,
	Description:  "Withdraws siloed liquidity for a specific remote chain from the SiloedLockReleaseTokenPool",
	ContractType: ContractType,
	ContractABI:  SiloedLockReleaseTokenPoolABI,
	NewContract:  NewSiloedLockReleaseTokenPoolContract,
	IsAllowedCaller: func(c *SiloedLockReleaseTokenPoolContract, opts *bind.CallOpts, caller common.Address, input WithdrawSiloedLiquidityInput) (bool, error) {
		rebalancer, err := c.GetChainRebalancer(opts, input.RemoteChainSelector)
		if err != nil {
			return false, err
		}
		return rebalancer == caller, nil
	},
	Validate: func(WithdrawSiloedLiquidityInput) error { return nil },
	CallContract: func(c *SiloedLockReleaseTokenPoolContract, opts *bind.TransactOpts, input WithdrawSiloedLiquidityInput) (*types.Transaction, error) {
		return c.contract.Transact(opts, "withdrawSiloedLiquidity", input.RemoteChainSelector, input.Amount)
	},
})
