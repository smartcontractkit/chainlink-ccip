package token_pool

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	gobindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_5_1/token_pool"
	cldf_evm "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations2/contract"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cld_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

var (
	ContractType cldf_deployment.ContractType = "TokenPool"
	Version      *semver.Version              = semver.MustParse("1.5.1")
)

type ApplyChainUpdatesArgs struct {
	RemoteChainSelectorsToRemove []uint64
	ChainsToAdd                  []gobindings.TokenPoolChainUpdate
}

type SetChainRateLimiterConfigArgs struct {
	RemoteChainSelector     uint64
	OutboundRateLimitConfig gobindings.RateLimiterConfig
	InboundRateLimitConfig  gobindings.RateLimiterConfig
}

type AddRemotePoolArgs struct {
	RemoteChainSelector uint64
	RemotePoolAddress   []byte
}

type SetRateLimitAdminArgs struct {
	NewAdmin common.Address
}

func NewReadGetTokenDecimals(c *gobindings.TokenPool) *cld_ops.Operation[contract.FunctionInput[struct{}], uint8, cldf_evm.Chain] {
	return contract.NewRead(contract.ReadParams[struct{}, uint8, *gobindings.TokenPool]{
		Name:         "token-pool:get-token-decimals",
		Version:      Version,
		Description:  "Gets the decimals of the token managed by the TokenPool 1.5.1 contract",
		ContractType: ContractType,
		Contract:     c,
		CallContract: func(tp *gobindings.TokenPool, opts *bind.CallOpts, args struct{}) (uint8, error) {
			return tp.GetTokenDecimals(opts)
		},
	})
}

func NewReadGetToken(c *gobindings.TokenPool) *cld_ops.Operation[contract.FunctionInput[struct{}], common.Address, cldf_evm.Chain] {
	return contract.NewRead(contract.ReadParams[struct{}, common.Address, *gobindings.TokenPool]{
		Name:         "token-pool:get-token",
		Version:      Version,
		Description:  "Gets the token address managed by the TokenPool 1.5.1 contract",
		ContractType: ContractType,
		Contract:     c,
		CallContract: func(tp *gobindings.TokenPool, opts *bind.CallOpts, args struct{}) (common.Address, error) {
			return tp.GetToken(opts)
		},
	})
}

func NewWriteApplyChainUpdates(c *gobindings.TokenPool) *cld_ops.Operation[contract.FunctionInput[ApplyChainUpdatesArgs], contract.WriteOutput, cldf_evm.Chain] {
	return contract.NewWrite(contract.WriteParams[ApplyChainUpdatesArgs, *gobindings.TokenPool]{
		Name:            "token-pool:apply-chain-updates",
		Version:         Version,
		Description:     "Applies chain updates to the TokenPool 1.5.1 contract",
		ContractType:    ContractType,
		ContractABI:     gobindings.TokenPoolABI,
		Contract:        c,
		IsAllowedCaller: contract.OnlyOwner[*gobindings.TokenPool, ApplyChainUpdatesArgs],
		Validate:        func(args ApplyChainUpdatesArgs) error { return nil },
		CallContract: func(tp *gobindings.TokenPool, opts *bind.TransactOpts, args ApplyChainUpdatesArgs) (*types.Transaction, error) {
			return tp.ApplyChainUpdates(opts, args.RemoteChainSelectorsToRemove, args.ChainsToAdd)
		},
	})
}

func NewWriteSetChainRateLimiterConfig(c *gobindings.TokenPool) *cld_ops.Operation[contract.FunctionInput[SetChainRateLimiterConfigArgs], contract.WriteOutput, cldf_evm.Chain] {
	return contract.NewWrite(contract.WriteParams[SetChainRateLimiterConfigArgs, *gobindings.TokenPool]{
		Name:         "token-pool:set-chain-rate-limiter-config",
		Version:      Version,
		Description:  "Sets the rate limiter configuration for a remote chain on the TokenPool 1.5.1 contract",
		ContractType: ContractType,
		ContractABI:  gobindings.TokenPoolABI,
		Contract:     c,
		IsAllowedCaller: func(tp *gobindings.TokenPool, opts *bind.CallOpts, caller common.Address, input SetChainRateLimiterConfigArgs) (bool, error) {
			admin, err := tp.GetRateLimitAdmin(opts)
			if err != nil {
				return false, fmt.Errorf("failed to get rate limit admin for pool at address %q: %w", tp.Address().Hex(), err)
			}

			owner, err := tp.Owner(opts)
			if err != nil {
				return false, fmt.Errorf("failed to get owner for pool at address %q: %w", tp.Address().Hex(), err)
			}

			return caller.Cmp(admin) == 0 || caller.Cmp(owner) == 0, nil
		},
		Validate: func(args SetChainRateLimiterConfigArgs) error { return nil },
		CallContract: func(tp *gobindings.TokenPool, opts *bind.TransactOpts, args SetChainRateLimiterConfigArgs) (*types.Transaction, error) {
			return tp.SetChainRateLimiterConfig(opts, args.RemoteChainSelector, args.OutboundRateLimitConfig, args.InboundRateLimitConfig)
		},
	})
}

func NewWriteAddRemotePool(c *gobindings.TokenPool) *cld_ops.Operation[contract.FunctionInput[AddRemotePoolArgs], contract.WriteOutput, cldf_evm.Chain] {
	return contract.NewWrite(contract.WriteParams[AddRemotePoolArgs, *gobindings.TokenPool]{
		Name:            "token-pool:add-remote-pool",
		Version:         Version,
		Description:     "Adds a remote pool for a given chain selector on the TokenPool 1.5.1 contract",
		ContractType:    ContractType,
		ContractABI:     gobindings.TokenPoolABI,
		Contract:        c,
		IsAllowedCaller: contract.OnlyOwner[*gobindings.TokenPool, AddRemotePoolArgs],
		Validate:        func(args AddRemotePoolArgs) error { return nil },
		CallContract: func(tp *gobindings.TokenPool, opts *bind.TransactOpts, args AddRemotePoolArgs) (*types.Transaction, error) {
			return tp.AddRemotePool(opts, args.RemoteChainSelector, args.RemotePoolAddress)
		},
	})
}

func NewWriteSetRateLimitAdmin(c *gobindings.TokenPool) *cld_ops.Operation[contract.FunctionInput[SetRateLimitAdminArgs], contract.WriteOutput, cldf_evm.Chain] {
	return contract.NewWrite(contract.WriteParams[SetRateLimitAdminArgs, *gobindings.TokenPool]{
		Name:            "token-pool:set-rate-limit-admin",
		Version:         Version,
		Description:     "Sets the rate limit admin for the TokenPool 1.5.1 contract",
		ContractType:    ContractType,
		ContractABI:     gobindings.TokenPoolABI,
		Contract:        c,
		IsAllowedCaller: contract.OnlyOwner[*gobindings.TokenPool, SetRateLimitAdminArgs],
		Validate:        func(args SetRateLimitAdminArgs) error { return nil },
		CallContract: func(tp *gobindings.TokenPool, opts *bind.TransactOpts, args SetRateLimitAdminArgs) (*types.Transaction, error) {
			return tp.SetRateLimitAdmin(opts, args.NewAdmin)
		},
	})
}

func NewReadGetCurrentInboundRateLimiterState(c *gobindings.TokenPool) *cld_ops.Operation[contract.FunctionInput[uint64], gobindings.RateLimiterTokenBucket, cldf_evm.Chain] {
	return contract.NewRead(contract.ReadParams[uint64, gobindings.RateLimiterTokenBucket, *gobindings.TokenPool]{
		Name:         "token-pool:get-current-inbound-rate-limiter-state",
		Version:      Version,
		Description:  "Calls getCurrentInboundRateLimiterState on the contract",
		ContractType: ContractType,
		Contract:     c,
		CallContract: func(tp *gobindings.TokenPool, opts *bind.CallOpts, args uint64) (gobindings.RateLimiterTokenBucket, error) {
			return tp.GetCurrentInboundRateLimiterState(opts, args)
		},
	})
}

func NewReadGetCurrentOutboundRateLimiterState(c *gobindings.TokenPool) *cld_ops.Operation[contract.FunctionInput[uint64], gobindings.RateLimiterTokenBucket, cldf_evm.Chain] {
	return contract.NewRead(contract.ReadParams[uint64, gobindings.RateLimiterTokenBucket, *gobindings.TokenPool]{
		Name:         "token-pool:get-current-outbound-rate-limiter-state",
		Version:      Version,
		Description:  "Calls getCurrentOutboundRateLimiterState on the contract",
		ContractType: ContractType,
		Contract:     c,
		CallContract: func(tp *gobindings.TokenPool, opts *bind.CallOpts, args uint64) (gobindings.RateLimiterTokenBucket, error) {
			return tp.GetCurrentOutboundRateLimiterState(opts, args)
		},
	})
}

func NewReadGetSupportedChains(c *gobindings.TokenPool) *cld_ops.Operation[contract.FunctionInput[struct{}], []uint64, cldf_evm.Chain] {
	return contract.NewRead(contract.ReadParams[struct{}, []uint64, *gobindings.TokenPool]{
		Name:         "token-pool:get-supported-chains",
		Version:      Version,
		Description:  "Calls getSupportedChains on the contract",
		ContractType: ContractType,
		Contract:     c,
		CallContract: func(tp *gobindings.TokenPool, opts *bind.CallOpts, args struct{}) ([]uint64, error) {
			return tp.GetSupportedChains(opts)
		},
	})
}

func NewReadGetRemoteToken(c *gobindings.TokenPool) *cld_ops.Operation[contract.FunctionInput[uint64], []byte, cldf_evm.Chain] {
	return contract.NewRead(contract.ReadParams[uint64, []byte, *gobindings.TokenPool]{
		Name:         "token-pool:get-remote-token",
		Version:      Version,
		Description:  "Calls getRemoteToken on the contract",
		ContractType: ContractType,
		Contract:     c,
		CallContract: func(tp *gobindings.TokenPool, opts *bind.CallOpts, args uint64) ([]byte, error) {
			return tp.GetRemoteToken(opts, args)
		},
	})
}

func NewReadGetRemotePools(c *gobindings.TokenPool) *cld_ops.Operation[contract.FunctionInput[uint64], [][]byte, cldf_evm.Chain] {
	return contract.NewRead(contract.ReadParams[uint64, [][]byte, *gobindings.TokenPool]{
		Name:         "token-pool:get-remote-pools",
		Version:      Version,
		Description:  "Calls getRemotePools on the contract",
		ContractType: ContractType,
		Contract:     c,
		CallContract: func(tp *gobindings.TokenPool, opts *bind.CallOpts, args uint64) ([][]byte, error) {
			return tp.GetRemotePools(opts, args)
		},
	})
}
