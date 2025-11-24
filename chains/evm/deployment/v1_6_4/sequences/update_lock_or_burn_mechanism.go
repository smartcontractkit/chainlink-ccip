package sequences

import (
	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"

	"fmt"

	contract_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	usdc_token_pool_proxy_ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_4/operations/usdc_token_pool_proxy"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"
)

type UpdateLockOrBurnMechanismSequenceInput struct {
	Address    map[uint64]common.Address
	Mechanisms map[uint64]usdc_token_pool_proxy_ops.UpdateLockOrBurnMechanismsArgs
}

var USDCTokenPoolProxyUpdateLockOrBurnMechanismSequence = operations.NewSequence(
	"USDCTokenPoolProxyUpdateLockOrBurnMechanismSequence",
	semver.MustParse("1.6.4"),
	"Updates the lock or burn mechanisms on a sequence of USDCTokenPoolProxy contracts on multiple chains",
	func(b operations.Bundle, chains cldf_chain.BlockChains, input UpdateLockOrBurnMechanismSequenceInput) (sequences.OnChainOutput, error) {
		writes := make([]contract_utils.WriteOutput, 0)

		for chainSel, mechanisms := range input.Mechanisms {
			chain, ok := chains.EVMChains()[chainSel]
			if !ok {
				return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not defined", chainSel)
			}
			address, ok := input.Address[chainSel]
			if !ok {
				return sequences.OnChainOutput{}, fmt.Errorf("address not found for chain selector %d", chainSel)
			}
			report, err := operations.ExecuteOperation(b, usdc_token_pool_proxy_ops.USDCTokenPoolProxyUpdateLockOrBurnMechanisms, chain, contract_utils.FunctionInput[usdc_token_pool_proxy_ops.UpdateLockOrBurnMechanismsArgs]{
				ChainSelector: chain.Selector,
				Address:       address,
				Args:          mechanisms,
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to execute USDCTokenPoolProxyUpdateLockOrBurnMechanismsOp on %s: %w", chain, err)
			}
			writes = append(writes, report.Output)
		}
		batch, err := contract_utils.NewBatchOperationFromWrites(writes)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to create batch operation from writes: %w", err)
		}
		return sequences.OnChainOutput{
			BatchOps: []mcms_types.BatchOperation{batch},
		}, nil
	})
