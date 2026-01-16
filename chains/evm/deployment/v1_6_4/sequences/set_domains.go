package sequences

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	usdc_token_pool_ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_4/operations/usdc_token_pool_cctp_v2"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
)

type SetDomainsSequenceInput struct {
	AddressesByChain map[uint64][]common.Address
	DomainsByChain   map[uint64]map[common.Address][]usdc_token_pool_ops.DomainUpdate
}

var (
	USDCTokenPoolSetDomainsSequence = operations.NewSequence(
		"USDCTokenPoolSetDomainsSequence",
		usdc_token_pool_ops.Version,
		"Sets domains on a sequence of USDCTokenPool contracts on multiple chains",
		func(b operations.Bundle, chains cldf_chain.BlockChains, input SetDomainsSequenceInput) (sequences.OnChainOutput, error) {

			writes := make([]contract.WriteOutput, 0)

			// Iterate over each chain selector in AddressesByChain
			for chainSelector, addresses := range input.AddressesByChain {
				// Get the chain object based on the chain selector
				chain, ok := chains.EVMChains()[chainSelector]
				if !ok {
					return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not defined", chainSelector)
				}

				// Iterate over each address for this chain selector
				for _, address := range addresses {
					// Get the domains for this specific address
					domainUpdates := input.DomainsByChain[chainSelector][address]

					// Execute the operation USDCTokenPoolSetDomains for this address
					report, err := operations.ExecuteOperation(b, usdc_token_pool_ops.USDCTokenPoolSetDomains, chain, contract.FunctionInput[[]usdc_token_pool_ops.DomainUpdate]{
						ChainSelector: chain.Selector,
						Address:       address,
						Args:          domainUpdates,
					})
					if err != nil {
						return sequences.OnChainOutput{}, fmt.Errorf("failed to execute USDCTokenPoolSetDomainsOp on %s: %w", chain, err)
					}
					writes = append(writes, report.Output)
				}
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
