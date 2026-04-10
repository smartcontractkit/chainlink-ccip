package tokens

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
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

		writes := make([]contract.WriteOutput, 0)
		for pool, minBlockConfirmations := range input.Settings {
			src := chain.Selector
			if !common.IsHexAddress(pool) {
				return sequences.OnChainOutput{}, fmt.Errorf("invalid pool address for src %d: %s", src, pool)
			}

			addr := common.HexToAddress(pool)
			if addr == (common.Address{}) {
				return sequences.OnChainOutput{}, fmt.Errorf("pool address cannot be the zero address for src %d", src)
			}

			report, err := operations.ExecuteOperation(
				b, token_pool.SetAllowedFinalityConfig, chain,
				contract.FunctionInput[[4]byte]{
					ChainSelector: src,
					Address:       addr,
					Args:          minBlockConfirmations,
				},
			)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to execute token_pool.SetAllowedFinalityConfig for pool %s on src %d: %w", pool, src, err)
			}

			writes = append(writes, report.Output)
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
