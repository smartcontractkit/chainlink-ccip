package siloed_lock_release_token_pool

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	gobindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_1/siloed_lock_release_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
)

var GetRebalancer = contract.NewRead(contract.ReadParams[struct{}, common.Address, *gobindings.SiloedLockReleaseTokenPool]{
	Name:         "siloed-lock-release-token-pool:get-rebalancer",
	Version:      Version,
	Description:  "Calls getRebalancer on the contract",
	ContractType: ContractType,
	NewContract:  gobindings.NewSiloedLockReleaseTokenPool,
	CallContract: func(c *gobindings.SiloedLockReleaseTokenPool, opts *bind.CallOpts, args struct{}) (common.Address, error) {
		return c.GetRebalancer(opts)
	},
})

var GetChainRebalancer = contract.NewRead(contract.ReadParams[uint64, common.Address, *gobindings.SiloedLockReleaseTokenPool]{
	Name:         "siloed-lock-release-token-pool:get-chain-rebalancer",
	Version:      Version,
	Description:  "Calls getChainRebalancer on the contract",
	ContractType: ContractType,
	NewContract:  gobindings.NewSiloedLockReleaseTokenPool,
	CallContract: func(c *gobindings.SiloedLockReleaseTokenPool, opts *bind.CallOpts, args uint64) (common.Address, error) {
		return c.GetChainRebalancer(opts, args)
	},
})

var GetSupportedChains = contract.NewRead(contract.ReadParams[struct{}, []uint64, *gobindings.SiloedLockReleaseTokenPool]{
	Name:         "siloed-lock-release-token-pool:get-supported-chains",
	Version:      Version,
	Description:  "Calls getSupportedChains on the contract",
	ContractType: ContractType,
	NewContract:  gobindings.NewSiloedLockReleaseTokenPool,
	CallContract: func(c *gobindings.SiloedLockReleaseTokenPool, opts *bind.CallOpts, args struct{}) ([]uint64, error) {
		return c.GetSupportedChains(opts)
	},
})

var IsSiloed = contract.NewRead(contract.ReadParams[uint64, bool, *gobindings.SiloedLockReleaseTokenPool]{
	Name:         "siloed-lock-release-token-pool:is-siloed",
	Version:      Version,
	Description:  "Calls isSiloed on the contract",
	ContractType: ContractType,
	NewContract:  gobindings.NewSiloedLockReleaseTokenPool,
	CallContract: func(c *gobindings.SiloedLockReleaseTokenPool, opts *bind.CallOpts, args uint64) (bool, error) {
		return c.IsSiloed(opts, args)
	},
})

var GetAvailableTokens = contract.NewRead(contract.ReadParams[uint64, *big.Int, *gobindings.SiloedLockReleaseTokenPool]{
	Name:         "siloed-lock-release-token-pool:get-available-tokens",
	Version:      Version,
	Description:  "Calls getAvailableTokens on the contract",
	ContractType: ContractType,
	NewContract:  gobindings.NewSiloedLockReleaseTokenPool,
	CallContract: func(c *gobindings.SiloedLockReleaseTokenPool, opts *bind.CallOpts, args uint64) (*big.Int, error) {
		return c.GetAvailableTokens(opts, args)
	},
})

var GetUnsiloedLiquidity = contract.NewRead(contract.ReadParams[struct{}, *big.Int, *gobindings.SiloedLockReleaseTokenPool]{
	Name:         "siloed-lock-release-token-pool:get-unsiloed-liquidity",
	Version:      Version,
	Description:  "Calls getUnsiloedLiquidity on the contract",
	ContractType: ContractType,
	NewContract:  gobindings.NewSiloedLockReleaseTokenPool,
	CallContract: func(c *gobindings.SiloedLockReleaseTokenPool, opts *bind.CallOpts, args struct{}) (*big.Int, error) {
		return c.GetUnsiloedLiquidity(opts)
	},
})

var SetRebalancer = contract.NewWrite(contract.WriteParams[common.Address, *gobindings.SiloedLockReleaseTokenPool]{
	Name:            "siloed-lock-release-token-pool:set-rebalancer",
	Version:         Version,
	Description:     "Calls setRebalancer on the contract",
	ContractType:    ContractType,
	ContractABI:     gobindings.SiloedLockReleaseTokenPoolMetaData.ABI,
	NewContract:     gobindings.NewSiloedLockReleaseTokenPool,
	IsAllowedCaller: contract.OnlyOwner[*gobindings.SiloedLockReleaseTokenPool, common.Address],
	Validate:        func(common.Address) error { return nil },
	CallContract: func(c *gobindings.SiloedLockReleaseTokenPool, opts *bind.TransactOpts, args common.Address) (*types.Transaction, error) {
		return c.SetRebalancer(opts, args)
	},
})

var SetSiloRebalancer = contract.NewWrite(contract.WriteParams[SetSiloRebalancerArgs, *gobindings.SiloedLockReleaseTokenPool]{
	Name:            "siloed-lock-release-token-pool:set-silo-rebalancer",
	Version:         Version,
	Description:     "Calls setSiloRebalancer on the contract",
	ContractType:    ContractType,
	ContractABI:     gobindings.SiloedLockReleaseTokenPoolMetaData.ABI,
	NewContract:     gobindings.NewSiloedLockReleaseTokenPool,
	IsAllowedCaller: contract.OnlyOwner[*gobindings.SiloedLockReleaseTokenPool, SetSiloRebalancerArgs],
	Validate:        func(SetSiloRebalancerArgs) error { return nil },
	CallContract: func(c *gobindings.SiloedLockReleaseTokenPool, opts *bind.TransactOpts, args SetSiloRebalancerArgs) (*types.Transaction, error) {
		return c.SetSiloRebalancer(opts, args.RemoteChainSelector, args.NewRebalancer)
	},
})

var WithdrawLiquidity = contract.NewWrite(contract.WriteParams[*big.Int, *gobindings.SiloedLockReleaseTokenPool]{
	Name:            "siloed-lock-release-token-pool:withdraw-liquidity",
	Version:         Version,
	Description:     "Calls withdrawLiquidity on the contract",
	ContractType:    ContractType,
	ContractABI:     gobindings.SiloedLockReleaseTokenPoolMetaData.ABI,
	NewContract:     gobindings.NewSiloedLockReleaseTokenPool,
	IsAllowedCaller: contract.OnlyOwner[*gobindings.SiloedLockReleaseTokenPool, *big.Int],
	Validate:        func(*big.Int) error { return nil },
	CallContract: func(c *gobindings.SiloedLockReleaseTokenPool, opts *bind.TransactOpts, args *big.Int) (*types.Transaction, error) {
		return c.WithdrawLiquidity(opts, args)
	},
})

var WithdrawSiloedLiquidity = contract.NewWrite(contract.WriteParams[WithdrawSiloedLiquidityArgs, *gobindings.SiloedLockReleaseTokenPool]{
	Name:            "siloed-lock-release-token-pool:withdraw-siloed-liquidity",
	Version:         Version,
	Description:     "Calls withdrawSiloedLiquidity on the contract",
	ContractType:    ContractType,
	ContractABI:     gobindings.SiloedLockReleaseTokenPoolMetaData.ABI,
	NewContract:     gobindings.NewSiloedLockReleaseTokenPool,
	IsAllowedCaller: contract.OnlyOwner[*gobindings.SiloedLockReleaseTokenPool, WithdrawSiloedLiquidityArgs],
	Validate:        func(WithdrawSiloedLiquidityArgs) error { return nil },
	CallContract: func(c *gobindings.SiloedLockReleaseTokenPool, opts *bind.TransactOpts, args WithdrawSiloedLiquidityArgs) (*types.Transaction, error) {
		return c.WithdrawSiloedLiquidity(opts, args.RemoteChainSelector, args.Amount)
	},
})
