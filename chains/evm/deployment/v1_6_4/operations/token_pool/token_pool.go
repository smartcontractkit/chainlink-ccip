package token_pool

import (
	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_5_1/token_pool"
)

var Version = semver.MustParse("1.6.4")

type ChainUpdate = token_pool.TokenPoolChainUpdate

type ApplyChainUpdatesArgs struct {
	RemoteChainSelectorsToRemove []uint64
	ChainsToAdd                  []ChainUpdate
}

type RemotePoolModification struct {
	Operation           string
	RemoteChainSelector uint64
	RemotePoolAddress   []byte
}

// Note: No "Deploy" Operation for Token Pool is needed since the TokenPool contract itself is an abstract
// contract, and will only exist as the parent of another contract to be deployed.

// Operation to call the applyChainUpdates() function on a token pool.
var ApplyChainUpdates = contract.NewWrite(contract.WriteParams[ApplyChainUpdatesArgs, *token_pool.TokenPool]{
	Name:            "token-pool:apply-chain-updates",
	Version:         Version,
	Description:     "Applies chain updates to the TokenPool contract",
	ContractType:    "TokenPool",
	ContractABI:     token_pool.TokenPoolABI,
	NewContract:     token_pool.NewTokenPool,
	IsAllowedCaller: contract.OnlyOwner[*token_pool.TokenPool, ApplyChainUpdatesArgs],
	Validate:        func(args ApplyChainUpdatesArgs) error { return nil },
	CallContract: func(tokenPool *token_pool.TokenPool, opts *bind.TransactOpts, args ApplyChainUpdatesArgs) (*types.Transaction, error) {
		return tokenPool.ApplyChainUpdates(opts, args.RemoteChainSelectorsToRemove, args.ChainsToAdd)
	},
})

var AddRemotePool = contract.NewWrite(contract.WriteParams[RemotePoolModification, *token_pool.TokenPool]{
	Name:            "token-pool:add-remote-pool",
	Version:         Version,
	Description:     "Adds a remote pool to the TokenPool remote chain config",
	ContractType:    "TokenPool",
	ContractABI:     token_pool.TokenPoolABI,
	NewContract:     token_pool.NewTokenPool,
	IsAllowedCaller: contract.OnlyOwner[*token_pool.TokenPool, RemotePoolModification],
	Validate:        func(args RemotePoolModification) error { return nil },
	CallContract: func(tokenPool *token_pool.TokenPool, opts *bind.TransactOpts, args RemotePoolModification) (*types.Transaction, error) {
		return tokenPool.AddRemotePool(opts, args.RemoteChainSelector, args.RemotePoolAddress)
	},
})

var RemoveRemotePool = contract.NewWrite(contract.WriteParams[RemotePoolModification, *token_pool.TokenPool]{
	Name:            "token-pool:remove-remote-pool",
	Version:         Version,
	Description:     "Removes a remote pool from the TokenPool remote chain config",
	ContractType:    "TokenPool",
	ContractABI:     token_pool.TokenPoolABI,
	NewContract:     token_pool.NewTokenPool,
	IsAllowedCaller: contract.OnlyOwner[*token_pool.TokenPool, RemotePoolModification],
	Validate:        func(args RemotePoolModification) error { return nil },
	CallContract: func(tokenPool *token_pool.TokenPool, opts *bind.TransactOpts, args RemotePoolModification) (*types.Transaction, error) {
		return tokenPool.RemoveRemotePool(opts, args.RemoteChainSelector, args.RemotePoolAddress)
	},
})
