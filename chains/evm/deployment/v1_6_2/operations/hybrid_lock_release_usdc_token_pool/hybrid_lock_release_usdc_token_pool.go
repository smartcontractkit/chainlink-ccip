package hybrid_lock_release_usdc_token_pool

import (
	"fmt"
	"math/big"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	cldf_evm "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations2/contract"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cld_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"

	gobindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_2/hybrid_lock_release_usdc_token_pool"
)

var ContractType cldf_deployment.ContractType = "HybridLockReleaseUSDCTokenPool"

var Version = semver.MustParse("1.6.2")

type WithdrawLiquidityArgs struct {
	RemoteChainSelector uint64
	Amount              *big.Int
}

func NewReadGetLockedTokensForChain(c gobindings.HybridLockReleaseUSDCTokenPoolInterface) *cld_ops.Operation[contract.FunctionInput[uint64], *big.Int, cldf_evm.Chain] {
	return contract.NewRead(contract.ReadParams[uint64, *big.Int, gobindings.HybridLockReleaseUSDCTokenPoolInterface]{
		Name:         "hybrid-lock-release-usdc-token-pool:get-locked-tokens-for-chain",
		Version:      Version,
		Description:  "Gets locked token balance for a remote chain selector",
		ContractType: ContractType,
		Contract:     c,
		CallContract: func(pool gobindings.HybridLockReleaseUSDCTokenPoolInterface, opts *bind.CallOpts, remoteChainSelector uint64) (*big.Int, error) {
			return pool.GetLockedTokensForChain(opts, remoteChainSelector)
		},
	})
}

func NewReadGetLiquidityProvider(c gobindings.HybridLockReleaseUSDCTokenPoolInterface) *cld_ops.Operation[contract.FunctionInput[uint64], common.Address, cldf_evm.Chain] {
	return contract.NewRead(contract.ReadParams[uint64, common.Address, gobindings.HybridLockReleaseUSDCTokenPoolInterface]{
		Name:         "hybrid-lock-release-usdc-token-pool:get-liquidity-provider",
		Version:      Version,
		Description:  "Gets the liquidity provider for a remote chain selector",
		ContractType: ContractType,
		Contract:     c,
		CallContract: func(pool gobindings.HybridLockReleaseUSDCTokenPoolInterface, opts *bind.CallOpts, remoteChainSelector uint64) (common.Address, error) {
			return pool.GetLiquidityProvider(opts, remoteChainSelector)
		},
	})
}

func NewReadShouldUseLockRelease(c gobindings.HybridLockReleaseUSDCTokenPoolInterface) *cld_ops.Operation[contract.FunctionInput[uint64], bool, cldf_evm.Chain] {
	return contract.NewRead(contract.ReadParams[uint64, bool, gobindings.HybridLockReleaseUSDCTokenPoolInterface]{
		Name:         "hybrid-lock-release-usdc-token-pool:should-use-lock-release",
		Version:      Version,
		Description:  "Returns whether a remote chain selector should use lock-release",
		ContractType: ContractType,
		Contract:     c,
		CallContract: func(pool gobindings.HybridLockReleaseUSDCTokenPoolInterface, opts *bind.CallOpts, remoteChainSelector uint64) (bool, error) {
			return pool.ShouldUseLockRelease(opts, remoteChainSelector)
		},
	})
}

func NewWriteWithdrawLiquidity(c gobindings.HybridLockReleaseUSDCTokenPoolInterface) *cld_ops.Operation[contract.FunctionInput[WithdrawLiquidityArgs], contract.WriteOutput, cldf_evm.Chain] {
	return contract.NewWrite(contract.WriteParams[WithdrawLiquidityArgs, gobindings.HybridLockReleaseUSDCTokenPoolInterface]{
		Name:            "hybrid-lock-release-usdc-token-pool:withdraw-liquidity",
		Version:         Version,
		Description:     "Withdraws liquidity for a remote chain selector",
		ContractType:    ContractType,
		ContractABI:     gobindings.HybridLockReleaseUSDCTokenPoolABI,
		Contract:        c,
		IsAllowedCaller: contract.OnlyOwner[gobindings.HybridLockReleaseUSDCTokenPoolInterface, WithdrawLiquidityArgs],
		Validate: func(args WithdrawLiquidityArgs) error {
			if args.Amount == nil || args.Amount.Sign() <= 0 {
				return fmt.Errorf("amount must be greater than zero")
			}
			return nil
		},
		CallContract: func(pool gobindings.HybridLockReleaseUSDCTokenPoolInterface, opts *bind.TransactOpts, args WithdrawLiquidityArgs) (*types.Transaction, error) {
			return pool.WithdrawLiquidity(opts, args.RemoteChainSelector, args.Amount)
		},
	})
}
