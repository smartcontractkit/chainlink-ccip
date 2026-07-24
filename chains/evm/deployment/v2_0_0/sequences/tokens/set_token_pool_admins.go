package tokens

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations/contract"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/token_pool"
	"github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
)

var SetTokenPoolAdmins = operations.NewSequence(
	"set-token-pool-admins",
	utils.Version_2_0_0,
	"Updates the rate limit admin and/or fee admin on a v2 token pool; no-op when values already match",
	func(b operations.Bundle, chains cldf_chain.BlockChains, input tokens.SetTokenPoolAdminsSequenceInput) (sequences.OnChainOutput, error) {
		chain, ok := chains.EVMChains()[input.Selector]
		if !ok {
			return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not defined", input.Selector)
		}
		if !common.IsHexAddress(input.PoolAddress) {
			return sequences.OnChainOutput{}, fmt.Errorf("invalid pool address for chain %d: %s", input.Selector, input.PoolAddress)
		}
		poolAddr := common.HexToAddress(input.PoolAddress)

		// Force-execute so a repeated apply on the same bundle sees fresh on-chain state
		// instead of a cached pre-write read (the no-op guarantee depends on this).
		currentReport, err := operations.ExecuteOperation(b, token_pool.GetDynamicConfig, chain, contract.FunctionInput[struct{}]{
			ChainSelector: input.Selector,
			Address:       poolAddr,
		}, operations.WithForceExecute[contract.FunctionInput[struct{}], evm.Chain]())
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to get dynamic config from token pool %s on chain %d: %w", poolAddr.Hex(), input.Selector, err)
		}
		current := currentReport.Output

		desiredRateLimitAdmin := current.RateLimitAdmin
		if input.RateLimitAdmin != nil {
			if !common.IsHexAddress(*input.RateLimitAdmin) {
				return sequences.OnChainOutput{}, fmt.Errorf("invalid rate limit admin address for chain %d: %s", input.Selector, *input.RateLimitAdmin)
			}
			desiredRateLimitAdmin = common.HexToAddress(*input.RateLimitAdmin)
		}
		desiredFeeAdmin := current.FeeAdmin
		if input.FeeAdmin != nil {
			if !common.IsHexAddress(*input.FeeAdmin) {
				return sequences.OnChainOutput{}, fmt.Errorf("invalid fee admin address for chain %d: %s", input.Selector, *input.FeeAdmin)
			}
			desiredFeeAdmin = common.HexToAddress(*input.FeeAdmin)
		}

		if desiredRateLimitAdmin == current.RateLimitAdmin && desiredFeeAdmin == current.FeeAdmin {
			b.Logger.Infof("Token pool admins already match desired values for pool %s on chain %d; skipping", poolAddr.Hex(), input.Selector)
			return sequences.OnChainOutput{}, nil
		}

		write, err := operations.ExecuteOperation(b, token_pool.SetDynamicConfig, chain, contract.FunctionInput[token_pool.SetDynamicConfigArgs]{
			ChainSelector: input.Selector,
			Address:       poolAddr,
			Args: token_pool.SetDynamicConfigArgs{
				Router:         current.Router,
				RateLimitAdmin: desiredRateLimitAdmin,
				FeeAdmin:       desiredFeeAdmin,
			},
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to set dynamic config on token pool %s on chain %d: %w", poolAddr.Hex(), input.Selector, err)
		}

		batch, err := contract.NewBatchOperationFromWrites([]contract.WriteOutput{write.Output})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to create batch operation from writes: %w", err)
		}
		return sequences.OnChainOutput{BatchOps: []mcms_types.BatchOperation{batch}}, nil
	},
)
