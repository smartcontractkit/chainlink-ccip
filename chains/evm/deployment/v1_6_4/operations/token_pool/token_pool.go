package token_pool

import (
	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_5_1/token_pool"
)

// This version is being used since the Token Pool gobinding was included in version 1.5.1, but the actual token
// pool this operation is being performed on may be different.
var Version *semver.Version = semver.MustParse("1.5.1")

type ConstructorArgs struct {
	Token             common.Address
	Decimals          uint8
	RateLimiterConfig token_pool.RateLimiterConfig
}

type ChainUpdate = token_pool.TokenPoolChainUpdate

type ApplyChainUpdatesArgs struct {
	RemoteChainSelectorsToRemove []uint64
	ChainsToAdd                  []ChainUpdate
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
