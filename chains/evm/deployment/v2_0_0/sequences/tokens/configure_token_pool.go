package tokens

import (
	evmops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations2/contract"
	aph_bindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v2_0_0/advanced_pool_hooks"
	tp_bindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v2_0_0/token_pool"
	"fmt"
	"math/big"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/advanced_pool_hooks"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/token_pool"
)

// ConfigureTokenPoolInput is the input for the ConfigureTokenPool sequence.
type ConfigureTokenPoolInput struct {
	// ChainSelector is the chain selector for the chain being configured.
	ChainSelector uint64
	// TokenPoolAddress is the address of the token pool.
	TokenPoolAddress common.Address
	// AdvancedPoolHooks is the address of the AdvancedPoolHooks contract.
	AdvancedPoolHooks common.Address
	// RouterAddress is the address of the Router contract on this chain.
	// If left empty, setRouter will not be attempted.
	RouterAddress common.Address
	// ThresholdAmountForAdditionalCCVs is the transfer threshold where additional CCVs are required.
	// If nil, the existing threshold will be retained.
	ThresholdAmountForAdditionalCCVs *big.Int
	// RateLimitAdmin is an additional address allowed to set rate limiters.
	// If left empty, setRateLimitAdmin will not be attempted.
	RateLimitAdmin common.Address
	// FeeAggregator is the address that will receive fee tokens when WithdrawFeeTokens is called.
	FeeAggregator common.Address
}

var ConfigureTokenPool = cldf_ops.NewSequence(
	"configure-token-pool",
	semver.MustParse("2.0.0"),
	"Configures a token pool on an EVM chain",
	func(b cldf_ops.Bundle, chain evm.Chain, input ConfigureTokenPoolInput) (output sequences.OnChainOutput, err error) {
		writes := make([]contract.WriteOutput, 0)

		// Set threshold amount for additional CCVs (if necessary)
		if input.ThresholdAmountForAdditionalCCVs != nil {
			currentThresholdAmountReport, err := evmops.ExecuteRead(b, chain, input.AdvancedPoolHooks, evmops.BindAs[aph_bindings.AdvancedPoolHooksInterface](aph_bindings.NewAdvancedPoolHooks), advanced_pool_hooks.NewReadGetThresholdAmount, struct{}{})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to get current threshold amount for additional CCVs on advanced pool hooks with address %s on %s: %w", input.AdvancedPoolHooks, chain, err)
			}
			if currentThresholdAmountReport.Output.Cmp(input.ThresholdAmountForAdditionalCCVs) != 0 {
				setThresholdAmountReport, err := evmops.ExecuteWrite(b, chain, input.AdvancedPoolHooks, evmops.BindAs[aph_bindings.AdvancedPoolHooksInterface](aph_bindings.NewAdvancedPoolHooks), advanced_pool_hooks.NewWriteSetThresholdAmount, input.ThresholdAmountForAdditionalCCVs)
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to set threshold amount for additional CCVs on advanced pool hooks with address %s on %s: %w", input.AdvancedPoolHooks, chain, err)
				}
				writes = append(writes, setThresholdAmountReport.Output)
			}
		}

		// Set dynamic config (if necessary)
		if input.RouterAddress != (common.Address{}) || input.RateLimitAdmin != (common.Address{}) || input.FeeAggregator != (common.Address{}) {
			currentDynamicConfigReport, err := evmops.ExecuteRead(b, chain, input.TokenPoolAddress, evmops.BindAs[tp_bindings.TokenPoolInterface](tp_bindings.NewTokenPool), token_pool.NewReadGetDynamicConfig, struct{}{})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to get current dynamic config from token pool with address %s on %s: %w", input.TokenPoolAddress, chain, err)
			}

			currentDynamicConfig := currentDynamicConfigReport.Output

			desiredRouter := currentDynamicConfig.Router
			if input.RouterAddress != (common.Address{}) {
				desiredRouter = input.RouterAddress
			}

			desiredRateLimitAdmin := currentDynamicConfig.RateLimitAdmin
			if input.RateLimitAdmin != (common.Address{}) {
				desiredRateLimitAdmin = input.RateLimitAdmin
			}

			desiredFeeAdmin := currentDynamicConfig.FeeAdmin
			if input.FeeAggregator != (common.Address{}) {
				desiredFeeAdmin = input.FeeAggregator
			}

			if desiredRouter != currentDynamicConfig.Router || desiredRateLimitAdmin != currentDynamicConfig.RateLimitAdmin || desiredFeeAdmin != currentDynamicConfig.FeeAdmin {
				setDynamicConfigReport, err := evmops.ExecuteWrite(b, chain, input.TokenPoolAddress, evmops.BindAs[tp_bindings.TokenPoolInterface](tp_bindings.NewTokenPool), token_pool.NewWriteSetDynamicConfig, token_pool.SetDynamicConfigArgs{
					Router:         desiredRouter,
					RateLimitAdmin: desiredRateLimitAdmin,
					FeeAdmin:       desiredFeeAdmin,
				})
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to set dynamic config on token pool with address %s on %s: %w", input.TokenPoolAddress, chain, err)
				}
				writes = append(writes, setDynamicConfigReport.Output)
			}
		}

		batchOp, err := contract.NewBatchOperationFromWrites(writes)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to create batch operation from writes: %w", err)
		}

		return sequences.OnChainOutput{
			BatchOps: []mcms_types.BatchOperation{batchOp},
		}, nil
	},
)

// makeAllowListUpdates compares the current and desired allow-lists and returns the addresses to add and remove.
func makeAllowListUpdates(current, desired []common.Address) (adds, removes []common.Address) {
	currentSet := make(map[common.Address]struct{}, len(current))
	for _, addr := range current {
		currentSet[addr] = struct{}{}
	}
	desiredSet := make(map[common.Address]struct{}, len(desired))
	for _, addr := range desired {
		desiredSet[addr] = struct{}{}
	}

	for addr := range desiredSet {
		if _, exists := currentSet[addr]; !exists {
			adds = append(adds, addr)
		}
	}
	for addr := range currentSet {
		if _, exists := desiredSet[addr]; !exists {
			removes = append(removes, addr)
		}
	}
	return adds, removes
}
