package tokens

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations/contract"
	ops2contract "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations2/contract"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/token_pool"
	tpbinding "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v2_0_0/token_pool"
	"github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
)

var SetTokenTransferFeeConfigForTokenPools = operations.NewSequence(
	"set-token-transfer-fee-config-for-token-pools",
	utils.Version_2_0_0,
	"Sets token transfer fee configs for token pools. Takes a map of pool address to a map of dest chain selector to fee config (or nil to disable the fee config for that dest).",
	func(b operations.Bundle, chains cldf_chain.BlockChains, input tokens.SetTokenTransferFeeSequenceInput) (sequences.OnChainOutput, error) {
		chain, ok := chains.EVMChains()[input.Selector]
		if !ok {
			return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not defined", input.Selector)
		}

		writes := make([]contract.WriteOutput, 0)
		for pool, cfg := range input.Settings {
			src := chain.Selector
			if !common.IsHexAddress(pool) {
				return sequences.OnChainOutput{}, fmt.Errorf("invalid pool address for src %d: %s", src, pool)
			}

			addr := common.HexToAddress(pool)
			if addr == (common.Address{}) {
				return sequences.OnChainOutput{}, fmt.Errorf("pool address cannot be the zero address for src %d", src)
			}

			tp, err := bindTokenPool(addr, chain)
			if err != nil {
				return sequences.OnChainOutput{}, err
			}

			args := token_pool.ApplyTokenTransferFeeConfigUpdatesArgs{
				DisableTokenTransferFeeConfigs: []uint64{},
				TokenTransferFeeConfigArgs:     []tpbinding.TokenPoolTokenTransferFeeConfigArgs{},
			}

			for dst, fee := range cfg {
				if fee == nil {
					args.DisableTokenTransferFeeConfigs = append(args.DisableTokenTransferFeeConfigs, dst)
				} else {
					args.TokenTransferFeeConfigArgs = append(
						args.TokenTransferFeeConfigArgs,
						tpbinding.TokenPoolTokenTransferFeeConfigArgs{
							DestChainSelector: dst,
							TokenTransferFeeConfig: tpbinding.IPoolV2TokenTransferFeeConfig{
								FinalityTransferFeeBps:     fee.DefaultFinalityTransferFeeBps,
								FastFinalityTransferFeeBps: fee.CustomFinalityTransferFeeBps,
								FinalityFeeUSDCents:        fee.DefaultFinalityFeeUSDCents,
								FastFinalityFeeUSDCents:    fee.CustomFinalityFeeUSDCents,
								DestBytesOverhead:          fee.DestBytesOverhead,
								DestGasOverhead:            fee.DestGasOverhead,
								IsEnabled:                  fee.IsEnabled,
							},
						},
					)
				}
			}

			if len(args.DisableTokenTransferFeeConfigs) > 0 || len(args.TokenTransferFeeConfigArgs) > 0 {
				report, err := operations.ExecuteOperation(
					b, token_pool.NewWriteApplyTokenTransferFeeConfigUpdates(tp), chain,
					ops2contract.FunctionInput[token_pool.ApplyTokenTransferFeeConfigUpdatesArgs]{
						Args: args,
					},
				)
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to execute token_pool.ApplyTokenTransferFeeConfigUpdates on %s: %w", chain.String(), err)
				}

				writes = append(writes, writeOutputOps2ToLegacy(report.Output))
			}
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
