package token_pool

import (
	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_5_1/token_pool"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

var (
	ContractType cldf_deployment.ContractType = "TokenPool"
	Version      *semver.Version              = semver.MustParse("1.5.1")
)

type ApplyChainUpdatesArgs struct {
	RemoteChainSelectorsToRemove []uint64
	ChainsToAdd                  []token_pool.TokenPoolChainUpdate
}

var ApplyChainUpdates = contract.NewWrite(contract.WriteParams[ApplyChainUpdatesArgs, *token_pool.TokenPool]{
	Name:            "token_pool:apply-chain-updates",
	Version:         Version,
	Description:     "Applies chain updates to the TokenPool 1.5.1 contract",
	ContractType:    ContractType,
	ContractABI:     token_pool.TokenPoolABI,
	NewContract:     token_pool.NewTokenPool,
	IsAllowedCaller: contract.OnlyOwner[*token_pool.TokenPool, ApplyChainUpdatesArgs],
	Validate:        func(args ApplyChainUpdatesArgs) error { return nil },
	CallContract: func(tp *token_pool.TokenPool, opts *bind.TransactOpts, args ApplyChainUpdatesArgs) (*types.Transaction, error) {
		return tp.ApplyChainUpdates(opts, args.RemoteChainSelectorsToRemove, args.ChainsToAdd)
	},
})
