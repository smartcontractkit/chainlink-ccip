package sequences

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"

	usdc_token_pool_proxy_ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_4/operations/usdc_token_pool_proxy"
)

type UpdatePoolAddressesSequenceInput struct {
	Address              map[uint64]common.Address
	PoolAddressesByChain map[uint64]usdc_token_pool_proxy_ops.PoolAddresses
}

var (
	USDCTokenPoolProxyUpdatePoolAddressesSequence = operations.NewSequence(
		"USDCTokenPoolProxyUpdatePoolAddressesSequence",
		semver.MustParse("1.6.4"),
		"Updates pool addresses on a sequence of USDCTokenPoolProxy contracts on multiple chains",
		func(b operations.Bundle, chains cldf_chain.BlockChains, input UpdatePoolAddressesSequenceInput) (sequences.OnChainOutput, error) {
			writes := make([]contract.WriteOutput, 0)

			// Iterate over each chain selector in the input
			for chainSel, poolAddresses := range input.PoolAddressesByChain {

				// Get the chain object based on the chain selector
				chain, ok := chains.EVMChains()[chainSel]
				if !ok {
					return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not defined", chainSel)
				}

				// Execute the operation USDCTokenPoolProxyUpdatePoolAddresses, on "chain", with the input being
				// PoolAddresses for the given chain selector
				report, err := operations.ExecuteOperation(b, usdc_token_pool_proxy_ops.USDCTokenPoolProxyUpdatePoolAddresses, chain, contract.FunctionInput[usdc_token_pool_proxy_ops.PoolAddresses]{
					ChainSelector: chain.Selector,
					Address:       input.Address[chainSel],
					Args:          poolAddresses,
				})
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to execute USDCTokenPoolProxyUpdatePoolAddressesOp on %s: %w", chain, err)
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
		})
)
