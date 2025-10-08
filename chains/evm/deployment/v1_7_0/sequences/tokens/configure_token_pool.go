package tokens

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	evm_contract "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/operations/token_pool"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"
)

// ConfigureTokenPoolInput is the input for the ConfigureTokenPool sequence.
type ConfigureTokenPoolInput struct {
	// ChainSelector is the chain selector for the chain being configured.
	ChainSelector uint64
	// TokenPoolAddress is the address of the token pool.
	TokenPoolAddress common.Address
	// AllowList is the list of addresses allowed to transfer tokens.
	// If empty upon deployment, an allow-list can never be set.
	// Likewise, if populated upon deployment, the allow-list can never be disabled.
	AllowList []common.Address
	// RouterAddress is the address of the Router contract on this chain.
	// If left empty, setRouter will not be attempted.
	RouterAddress common.Address
	// RateLimitAdmin is an additional address allowed to set rate limiters.
	// If left empty, setRateLimitAdmin will not be attempted.
	RateLimitAdmin common.Address
}

var ConfigureTokenPool = cldf_ops.NewSequence(
	"configure-token-pool",
	semver.MustParse("1.7.0"),
	"Configures a token pool on an EVM chain",
	func(b operations.Bundle, chain evm.Chain, input ConfigureTokenPoolInput) (output sequences.OnChainOutput, err error) {
		writes := make([]contract.WriteOutput, 0)

		// First, check if the allow-list is enabled
		if len(input.AllowList) != 0 {
			allowListEnabledReport, err := cldf_ops.ExecuteOperation(b, token_pool.GetAllowListEnabled, chain, evm_contract.FunctionInput[any]{
				ChainSelector: input.ChainSelector,
				Address:       input.TokenPoolAddress,
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to get allow-list status from token pool with address %s on %s: %w", input.TokenPoolAddress, chain, err)
			}
			if allowListEnabledReport.Output {
				// Allow-list is enabled, so we first check the current allow-list
				currentAllowListReport, err := cldf_ops.ExecuteOperation(b, token_pool.GetAllowList, chain, evm_contract.FunctionInput[any]{
					ChainSelector: input.ChainSelector,
					Address:       input.TokenPoolAddress,
				})
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to get current allow-list from token pool with address %s on %s: %w", input.TokenPoolAddress, chain, err)
				}
				adds, removes := makeAllowListUpdates(currentAllowListReport.Output, input.AllowList)

				// Apply any updates to the allow-list if they exist
				if len(adds) != 0 || len(removes) != 0 {
					applyAllowListUpdatesReport, err := cldf_ops.ExecuteOperation(b, token_pool.ApplyAllowListUpdates, chain, evm_contract.FunctionInput[token_pool.ApplyAllowListUpdatesArgs]{
						ChainSelector: input.ChainSelector,
						Address:       input.TokenPoolAddress,
						Args: token_pool.ApplyAllowListUpdatesArgs{
							Adds:    adds,
							Removes: removes,
						},
					})
					if err != nil {
						return sequences.OnChainOutput{}, fmt.Errorf("failed to apply allow-list updates to token pool with address %s on %s: %w", input.TokenPoolAddress, chain, err)
					}
					writes = append(writes, applyAllowListUpdatesReport.Output)
				}
			}
		}

		// Set router (if necessary)
		// Check the router currently set on the token pool
		if input.RouterAddress != (common.Address{}) {
			currentRouterReport, err := cldf_ops.ExecuteOperation(b, token_pool.GetRouter, chain, evm_contract.FunctionInput[any]{
				ChainSelector: input.ChainSelector,
				Address:       input.TokenPoolAddress,
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to get current router from token pool with address %s on %s: %w", input.TokenPoolAddress, chain, err)
			}
			if currentRouterReport.Output != input.RouterAddress {
				// Router is not set to desired, so update it
				setRouterReport, err := cldf_ops.ExecuteOperation(b, token_pool.SetRouter, chain, evm_contract.FunctionInput[common.Address]{
					ChainSelector: input.ChainSelector,
					Address:       input.TokenPoolAddress,
					Args:          input.RouterAddress,
				})
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to set router on token pool with address %s on %s: %w", input.TokenPoolAddress, chain, err)
				}
				writes = append(writes, setRouterReport.Output)
			}
		}

		// Set rate limit admin (if necessary)
		// Check the rate limit admin currently set on the token pool
		if input.RateLimitAdmin != (common.Address{}) {
			currentRateLimitAdminReport, err := cldf_ops.ExecuteOperation(b, token_pool.GetRateLimitAdmin, chain, evm_contract.FunctionInput[any]{
				ChainSelector: input.ChainSelector,
				Address:       input.TokenPoolAddress,
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to get current rate limit admin from token pool with address %s on %s: %w", input.TokenPoolAddress, chain, err)
			}
			if currentRateLimitAdminReport.Output != input.RateLimitAdmin {
				// Rate limit admin is not set to desired, so update it
				setRateLimitAdminReport, err := cldf_ops.ExecuteOperation(b, token_pool.SetRateLimitAdmin, chain, evm_contract.FunctionInput[common.Address]{
					ChainSelector: input.ChainSelector,
					Address:       input.TokenPoolAddress,
					Args:          input.RateLimitAdmin,
				})
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to set rate limit admin on token pool with address %s on %s: %w", input.TokenPoolAddress, chain, err)
				}
				writes = append(writes, setRateLimitAdminReport.Output)
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
