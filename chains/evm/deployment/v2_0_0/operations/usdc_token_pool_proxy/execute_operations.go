package usdc_token_pool_proxy

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	gobindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v2_0_0/usdc_token_pool_proxy"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
)

// PoolAddresses groups pool addresses for deploy/updatePoolAddresses call sites.
type PoolAddresses = gobindings.USDCTokenPoolProxyPoolAddresses

var UpdateLockOrBurnMechanisms = contract.NewWrite(contract.WriteParams[UpdateLockOrBurnMechanismsArgs, *gobindings.USDCTokenPoolProxy]{
	Name:            "usdc-token-pool-proxy:update-lock-or-burn-mechanisms",
	Version:         Version,
	Description:     "Calls updateLockOrBurnMechanisms on the contract",
	ContractType:    ContractType,
	ContractABI:     gobindings.USDCTokenPoolProxyMetaData.ABI,
	NewContract:     gobindings.NewUSDCTokenPoolProxy,
	IsAllowedCaller: contract.OnlyOwner[*gobindings.USDCTokenPoolProxy, UpdateLockOrBurnMechanismsArgs],
	Validate:        func(UpdateLockOrBurnMechanismsArgs) error { return nil },
	CallContract: func(c *gobindings.USDCTokenPoolProxy, opts *bind.TransactOpts, args UpdateLockOrBurnMechanismsArgs) (*types.Transaction, error) {
		return c.UpdateLockOrBurnMechanisms(opts, args.RemoteChainSelectors, args.Mechanisms)
	},
})

var SetFeeAggregator = contract.NewWrite(contract.WriteParams[common.Address, *gobindings.USDCTokenPoolProxy]{
	Name:            "usdc-token-pool-proxy:set-fee-aggregator",
	Version:         Version,
	Description:     "Calls setFeeAggregator on the contract",
	ContractType:    ContractType,
	ContractABI:     gobindings.USDCTokenPoolProxyMetaData.ABI,
	NewContract:     gobindings.NewUSDCTokenPoolProxy,
	IsAllowedCaller: contract.OnlyOwner[*gobindings.USDCTokenPoolProxy, common.Address],
	Validate:        func(common.Address) error { return nil },
	CallContract: func(c *gobindings.USDCTokenPoolProxy, opts *bind.TransactOpts, args common.Address) (*types.Transaction, error) {
		return c.SetFeeAggregator(opts, args)
	},
})

var GetPools = contract.NewRead(contract.ReadParams[struct{}, gobindings.USDCTokenPoolProxyPoolAddresses, *gobindings.USDCTokenPoolProxy]{
	Name:         "usdc-token-pool-proxy:get-pools",
	Version:      Version,
	Description:  "Calls getPools on the contract",
	ContractType: ContractType,
	NewContract:  gobindings.NewUSDCTokenPoolProxy,
	CallContract: func(c *gobindings.USDCTokenPoolProxy, opts *bind.CallOpts, args struct{}) (gobindings.USDCTokenPoolProxyPoolAddresses, error) {
		return c.GetPools(opts)
	},
})

var UpdatePoolAddresses = contract.NewWrite(contract.WriteParams[gobindings.USDCTokenPoolProxyPoolAddresses, *gobindings.USDCTokenPoolProxy]{
	Name:            "usdc-token-pool-proxy:update-pool-addresses",
	Version:         Version,
	Description:     "Calls updatePoolAddresses on the contract",
	ContractType:    ContractType,
	ContractABI:     gobindings.USDCTokenPoolProxyMetaData.ABI,
	NewContract:     gobindings.NewUSDCTokenPoolProxy,
	IsAllowedCaller: contract.OnlyOwner[*gobindings.USDCTokenPoolProxy, gobindings.USDCTokenPoolProxyPoolAddresses],
	Validate:        func(gobindings.USDCTokenPoolProxyPoolAddresses) error { return nil },
	CallContract: func(c *gobindings.USDCTokenPoolProxy, opts *bind.TransactOpts, args gobindings.USDCTokenPoolProxyPoolAddresses) (*types.Transaction, error) {
		return c.UpdatePoolAddresses(opts, args)
	},
})
