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

	usdc_token_pool_ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_4/operations/usdc_token_pool"
)

type SetDomainsSequenceInput struct {
	Address        common.Address
	DomainsByChain map[uint64][]usdc_token_pool_ops.DomainUpdate
}

var (
	USDCTokenPoolSetDomainsSequence = operations.NewSequence(
		"USDCTokenPoolSetDomainsSequence",
		semver.MustParse("1.6.4"),
		"Sets domains on a sequence of USDCTokenPool contracts on multiple chains",
		func(b operations.Bundle, chains cldf_chain.BlockChains, input SetDomainsSequenceInput) (sequences.OnChainOutput, error) {

			writes := make([]contract.WriteOutput, 0)

			// Iterate over each chain selector in the input
			for chainSel, domains := range input.DomainsByChain {

				// Get the chain object based on the chain selector
				chain, ok := chains.EVMChains()[chainSel]
				if !ok {
					return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not defined", chainSel)
				}

				// Execute the operation USDCTokenPoolSetDomains, on "chain", with the input being an array of
				// DomainUpdate structs, with the first and only item being the domains for the given chain selector
				report, err := operations.ExecuteOperation(b, usdc_token_pool_ops.USDCTokenPoolSetDomains, chain, contract.FunctionInput[[]usdc_token_pool_ops.DomainUpdate]{
					ChainSelector: chain.Selector,
					Address:       input.Address,
					Args:          domains,
				})
				fmt.Println("Report output in sequence: ", report.Output)
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to execute USDCTokenPoolSetDomainsOp on %s: %w", chain, err)
				}
				writes = append(writes, report.Output)
			}
			batch, err := contract.NewBatchOperationFromWrites(writes)
			fmt.Println("Batch: ", batch)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to create batch operation from writes: %w", err)
			}
			return sequences.OnChainOutput{
				BatchOps: []mcms_types.BatchOperation{batch},
			}, nil
		})
)
