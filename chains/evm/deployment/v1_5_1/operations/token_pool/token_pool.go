package token_pool

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
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

type SetChainRateLimiterConfigArgs struct {
	RemoteChainSelector     uint64
	OutboundRateLimitConfig token_pool.RateLimiterConfig
	InboundRateLimitConfig  token_pool.RateLimiterConfig
}

type AddRemotePoolArgs struct {
	RemoteChainSelector uint64
	RemotePoolAddress   []byte
}

type SetRateLimitAdminArgs struct {
	NewAdmin common.Address
}

var GetToken = contract.NewRead(contract.ReadParams[struct{}, common.Address, *token_pool.TokenPool]{
	Name:         "token-pool:get-token",
	Version:      Version,
	Description:  "Gets the token address managed by the TokenPool 1.5.1 contract",
	ContractType: ContractType,
	NewContract:  token_pool.NewTokenPool,
	CallContract: func(tp *token_pool.TokenPool, opts *bind.CallOpts, args struct{}) (common.Address, error) {
		return tp.GetToken(opts)
	},
})

var ApplyChainUpdates = contract.NewWrite(contract.WriteParams[ApplyChainUpdatesArgs, *token_pool.TokenPool]{
	Name:            "token-pool:apply-chain-updates",
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

var SetChainRateLimiterConfig = contract.NewWrite(contract.WriteParams[SetChainRateLimiterConfigArgs, *token_pool.TokenPool]{
	Name:         "token-pool:set-chain-rate-limiter-config",
	Version:      Version,
	Description:  "Sets the rate limiter configuration for a remote chain on the TokenPool 1.5.1 contract",
	ContractType: ContractType,
	ContractABI:  token_pool.TokenPoolABI,
	NewContract:  token_pool.NewTokenPool,
	IsAllowedCaller: func(tp *token_pool.TokenPool, opts *bind.CallOpts, caller common.Address, input SetChainRateLimiterConfigArgs) (bool, error) {
		admin, err := tp.GetRateLimitAdmin(opts)
		if err != nil {
			return false, fmt.Errorf("failed to get rate limit admin for pool at address %q: %w", tp.Address().Hex(), err)
		}

		owner, err := tp.Owner(opts)
		if err != nil {
			return false, fmt.Errorf("failed to get owner for pool at address %q: %w", tp.Address().Hex(), err)
		}
		fmt.Println("Caller address:", caller.Hex())
		fmt.Println("Admin address:", admin.Hex())
		fmt.Println("Owner address:", owner.Hex())

		// Rate limit config can be set by either the rate limit admin or the owner
		return caller.Cmp(admin) == 0 || caller.Cmp(owner) == 0, nil
	},
	Validate: func(args SetChainRateLimiterConfigArgs) error { return nil },
	CallContract: func(tp *token_pool.TokenPool, opts *bind.TransactOpts, args SetChainRateLimiterConfigArgs) (*types.Transaction, error) {
		fmt.Println("Setting rate limits on token pool at address:", tp.Address().Hex())
		fmt.Println("Data for setting rate limits:", args)
		return tp.SetChainRateLimiterConfig(opts, args.RemoteChainSelector, args.OutboundRateLimitConfig, args.InboundRateLimitConfig)
	},
})

var AddRemotePool = contract.NewWrite(contract.WriteParams[AddRemotePoolArgs, *token_pool.TokenPool]{
	Name:            "token-pool:add-remote-pool",
	Version:         Version,
	Description:     "Adds a remote pool for a given chain selector on the TokenPool 1.5.1 contract",
	ContractType:    ContractType,
	ContractABI:     token_pool.TokenPoolABI,
	NewContract:     token_pool.NewTokenPool,
	IsAllowedCaller: contract.OnlyOwner[*token_pool.TokenPool, AddRemotePoolArgs],
	Validate:        func(args AddRemotePoolArgs) error { return nil },
	CallContract: func(tp *token_pool.TokenPool, opts *bind.TransactOpts, args AddRemotePoolArgs) (*types.Transaction, error) {
		return tp.AddRemotePool(opts, args.RemoteChainSelector, args.RemotePoolAddress)
	},
})

var SetRateLimitAdmin = contract.NewWrite(contract.WriteParams[SetRateLimitAdminArgs, *token_pool.TokenPool]{
	Name:            "token-pool:set-rate-limit-admin",
	Version:         Version,
	Description:     "Sets the rate limit admin for the TokenPool 1.5.1 contract",
	ContractType:    ContractType,
	ContractABI:     token_pool.TokenPoolABI,
	NewContract:     token_pool.NewTokenPool,
	IsAllowedCaller: contract.OnlyOwner[*token_pool.TokenPool, SetRateLimitAdminArgs],
	Validate:        func(args SetRateLimitAdminArgs) error { return nil },
	CallContract: func(tp *token_pool.TokenPool, opts *bind.TransactOpts, args SetRateLimitAdminArgs) (*types.Transaction, error) {
		return tp.SetRateLimitAdmin(opts, args.NewAdmin)
	},
})
