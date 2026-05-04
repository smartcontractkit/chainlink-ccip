package rmn_remote

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"

	gobindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_0/rmn_remote"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
)

var Curse0 = contract.NewWrite(contract.WriteParams[[][16]byte, *gobindings.RMNRemote]{
	Name:            "rmn-remote:curse0",
	Version:         Version,
	Description:     "Calls curse0 on the contract",
	ContractType:    ContractType,
	ContractABI:     gobindings.RMNRemoteMetaData.ABI,
	NewContract:     gobindings.NewRMNRemote,
	IsAllowedCaller: contract.OnlyOwner[*gobindings.RMNRemote, [][16]byte],
	Validate:        func([][16]byte) error { return nil },
	CallContract: func(c *gobindings.RMNRemote, opts *bind.TransactOpts, args [][16]byte) (*types.Transaction, error) {
		return c.Curse0(opts, args)
	},
})

var Uncurse0 = contract.NewWrite(contract.WriteParams[[][16]byte, *gobindings.RMNRemote]{
	Name:            "rmn-remote:uncurse0",
	Version:         Version,
	Description:     "Calls uncurse0 on the contract",
	ContractType:    ContractType,
	ContractABI:     gobindings.RMNRemoteMetaData.ABI,
	NewContract:     gobindings.NewRMNRemote,
	IsAllowedCaller: contract.OnlyOwner[*gobindings.RMNRemote, [][16]byte],
	Validate:        func([][16]byte) error { return nil },
	CallContract: func(c *gobindings.RMNRemote, opts *bind.TransactOpts, args [][16]byte) (*types.Transaction, error) {
		return c.Uncurse0(opts, args)
	},
})
