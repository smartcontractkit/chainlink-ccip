package lock_release_token_pool

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	gobindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_1/lock_release_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
)

var GetRebalancer = contract.NewRead(contract.ReadParams[struct{}, common.Address, *gobindings.LockReleaseTokenPool]{
	Name:         "lock-release-token-pool:get-rebalancer",
	Version:      Version,
	Description:  "Calls getRebalancer on the contract",
	ContractType: ContractType,
	NewContract:  gobindings.NewLockReleaseTokenPool,
	CallContract: func(c *gobindings.LockReleaseTokenPool, opts *bind.CallOpts, args struct{}) (common.Address, error) {
		return c.GetRebalancer(opts)
	},
})

var SetRebalancer = contract.NewWrite(contract.WriteParams[common.Address, *gobindings.LockReleaseTokenPool]{
	Name:            "lock-release-token-pool:set-rebalancer",
	Version:         Version,
	Description:     "Calls setRebalancer on the contract",
	ContractType:    ContractType,
	ContractABI:     gobindings.LockReleaseTokenPoolMetaData.ABI,
	NewContract:     gobindings.NewLockReleaseTokenPool,
	IsAllowedCaller: contract.OnlyOwner[*gobindings.LockReleaseTokenPool, common.Address],
	Validate:        func(common.Address) error { return nil },
	CallContract: func(c *gobindings.LockReleaseTokenPool, opts *bind.TransactOpts, args common.Address) (*types.Transaction, error) {
		return c.SetRebalancer(opts, args)
	},
})

var WithdrawLiquidity = contract.NewWrite(contract.WriteParams[*big.Int, *gobindings.LockReleaseTokenPool]{
	Name:            "lock-release-token-pool:withdraw-liquidity",
	Version:         Version,
	Description:     "Calls withdrawLiquidity on the contract",
	ContractType:    ContractType,
	ContractABI:     gobindings.LockReleaseTokenPoolMetaData.ABI,
	NewContract:     gobindings.NewLockReleaseTokenPool,
	IsAllowedCaller: contract.OnlyOwner[*gobindings.LockReleaseTokenPool, *big.Int],
	Validate:        func(*big.Int) error { return nil },
	CallContract: func(c *gobindings.LockReleaseTokenPool, opts *bind.TransactOpts, args *big.Int) (*types.Transaction, error) {
		return c.WithdrawLiquidity(opts, args)
	},
})
