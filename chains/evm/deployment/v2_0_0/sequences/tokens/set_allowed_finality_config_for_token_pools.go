package tokens

import (
	"bytes"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations/contract"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/token_pool"
	"github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
)

var SetAllowedFinalityConfigForTokenPools = operations.NewSequence(
	"set-finality-config-for-token-pools",
	utils.Version_2_0_0,
	"Sets the finality config for token pools. Takes a map of pool address to finality configs.",
	func(b operations.Bundle, chains cldf_chain.BlockChains, input tokens.SetAllowedFinalityConfigSequenceInput) (sequences.OnChainOutput, error) {
		chain, ok := chains.EVMChains()[input.Selector]
		if !ok {
			return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not defined", input.Selector)
		}

		writes := make([]contract.WriteOutput, 0, len(input.Settings))
		for pool, finalityConfig := range input.Settings {
			if err := finalityConfig.Validate(); err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("invalid finality config for pool %s: %w", pool, err)
			}

			selector := chain.Selector
			if !common.IsHexAddress(pool) {
				return sequences.OnChainOutput{}, fmt.Errorf("invalid pool address for src %d: %s", selector, pool)
			}

			addr := common.HexToAddress(pool)
			if addr == (common.Address{}) {
				return sequences.OnChainOutput{}, fmt.Errorf("pool address cannot be the zero address for src %d", selector)
			}

			currentFinalityConfigReport, err := operations.ExecuteOperation(b, token_pool.GetAllowedFinalityConfig, chain, contract.FunctionInput[struct{}]{
				ChainSelector: selector,
				Address:       addr,
				Args:          struct{}{},
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to get current finality config for token pool at address %s on chain %d: %w", addr.Hex(), selector, err)
			}

			requestedFinalityConfig := finalityConfig.Raw()
			if !bytes.Equal(currentFinalityConfigReport.Output[:], requestedFinalityConfig[:]) {
				write, err := operations.ExecuteOperation(b, token_pool.SetAllowedFinalityConfig, chain, contract.FunctionInput[[4]byte]{
					ChainSelector: selector,
					Address:       addr,
					Args:          requestedFinalityConfig,
				})
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to set finality config for token pool at address %s on chain %d: %w", addr.Hex(), selector, err)
				}
				writes = append(writes, write.Output)
			}
		}

		if len(writes) == 0 {
			return sequences.OnChainOutput{}, nil
		}

		batch, err := contract.NewBatchOperationFromWrites(writes)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to create batch operation from writes: %w", err)
		}

		return sequences.OnChainOutput{
			BatchOps: []mcms_types.BatchOperation{batch},
		}, nil
	},
)
