package lombard_token_pool

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"

	gobindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v2_0_0/lombard_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
)

var SetPath = contract.NewWrite(contract.WriteParams[SetPathArgs, *gobindings.LombardTokenPool]{
	Name:            "lombard-token-pool:set-path",
	Version:         Version,
	Description:     "Calls setPath on the contract",
	ContractType:    ContractType,
	ContractABI:     gobindings.LombardTokenPoolMetaData.ABI,
	NewContract:     gobindings.NewLombardTokenPool,
	IsAllowedCaller: contract.OnlyOwner[*gobindings.LombardTokenPool, SetPathArgs],
	Validate:        func(SetPathArgs) error { return nil },
	CallContract: func(c *gobindings.LombardTokenPool, opts *bind.TransactOpts, args SetPathArgs) (*types.Transaction, error) {
		return c.SetPath(opts, args.RemoteChainSelector, args.LChainId, args.AllowedCaller, args.RemoteAdapter)
	},
})

var GetPath = contract.NewRead(contract.ReadParams[uint64, gobindings.LombardTokenPoolPath, *gobindings.LombardTokenPool]{
	Name:         "lombard-token-pool:get-path",
	Version:      Version,
	Description:  "Calls getPath on the contract",
	ContractType: ContractType,
	NewContract:  gobindings.NewLombardTokenPool,
	CallContract: func(c *gobindings.LombardTokenPool, opts *bind.CallOpts, args uint64) (gobindings.LombardTokenPoolPath, error) {
		return c.GetPath(opts, args)
	},
})
