package proxy

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	gobindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v2_0_0/proxy"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
)

// ProxyABI is the JSON ABI of the Proxy contract.
var ProxyABI = gobindings.ProxyMetaData.ABI

// ProxyBin is the EVM runtime bytecode of the Proxy contract.
var ProxyBin = gobindings.ProxyMetaData.Bin

var GetTarget = contract.NewRead(contract.ReadParams[struct{}, common.Address, *gobindings.Proxy]{
	Name:         "proxy:get-target",
	Version:      Version,
	Description:  "Calls getTarget on the contract",
	ContractType: ContractType,
	NewContract:  gobindings.NewProxy,
	CallContract: func(c *gobindings.Proxy, opts *bind.CallOpts, args struct{}) (common.Address, error) {
		return c.GetTarget(opts)
	},
})

var SetTarget = contract.NewWrite(contract.WriteParams[common.Address, *gobindings.Proxy]{
	Name:            "proxy:set-target",
	Version:         Version,
	Description:     "Calls setTarget on the contract",
	ContractType:    ContractType,
	ContractABI:     gobindings.ProxyMetaData.ABI,
	NewContract:     gobindings.NewProxy,
	IsAllowedCaller: contract.OnlyOwner[*gobindings.Proxy, common.Address],
	Validate:        func(common.Address) error { return nil },
	CallContract: func(c *gobindings.Proxy, opts *bind.TransactOpts, args common.Address) (*types.Transaction, error) {
		return c.SetTarget(opts, args)
	},
})

var GetFeeAggregator = contract.NewRead(contract.ReadParams[struct{}, common.Address, *gobindings.Proxy]{
	Name:         "proxy:get-fee-aggregator",
	Version:      Version,
	Description:  "Calls getFeeAggregator on the contract",
	ContractType: ContractType,
	NewContract:  gobindings.NewProxy,
	CallContract: func(c *gobindings.Proxy, opts *bind.CallOpts, args struct{}) (common.Address, error) {
		return c.GetFeeAggregator(opts)
	},
})

var SetFeeAggregator = contract.NewWrite(contract.WriteParams[common.Address, *gobindings.Proxy]{
	Name:            "proxy:set-fee-aggregator",
	Version:         Version,
	Description:     "Calls setFeeAggregator on the contract",
	ContractType:    ContractType,
	ContractABI:     gobindings.ProxyMetaData.ABI,
	NewContract:     gobindings.NewProxy,
	IsAllowedCaller: contract.OnlyOwner[*gobindings.Proxy, common.Address],
	Validate:        func(common.Address) error { return nil },
	CallContract: func(c *gobindings.Proxy, opts *bind.TransactOpts, args common.Address) (*types.Transaction, error) {
		return c.SetFeeAggregator(opts, args)
	},
})
