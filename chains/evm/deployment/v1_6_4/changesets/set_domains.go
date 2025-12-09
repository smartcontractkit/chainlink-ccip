package changesets

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_4/operations/usdc_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_4/sequences"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
)

type SetDomainsInput struct {
	ChainInputs []SetDomainsPerChainInput
	MCMS        mcms.Input
}

type SetDomainsPerChainInput struct {
	ChainSelector uint64
	Address       common.Address
	Domains       []usdc_token_pool.DomainUpdate
}

// This changeset is used to set the domains for a given token in the USDCTokenPool contract.
func SetDomainsChangeset() cldf.ChangeSetV2[SetDomainsInput] {
	return cldf.CreateChangeSet(setDomainsApply(), setDomainsVerify())
}

func setDomainsApply() func(cldf.Environment, SetDomainsInput) (cldf.ChangesetOutput, error) {
	return func(e cldf.Environment, input SetDomainsInput) (cldf.ChangesetOutput, error) {
		batchOps := make([]mcms_types.BatchOperation, 0)
		reports := make([]cldf_ops.Report[any, any], 0)

		// By using a mapping from chain selector to a slice of addresses and another mapping from chain selector to a mapping
		// from address to a slice of DomainUpdate structs, we can allow for multiple contract calls on the same chain selector
		// but to different addresses.

		// On chains where there are multiple USDC Token pools of different versions, such
		// as on Ethereum mainnet, where V1 and V2 pools will coexist, this allows for simpler invocations of this changeset
		// across all contracts. Otherwise, this changeset would only be capable of setting domains for a single contract
		// at a time per chain, which would require multiple invocations of this changeset to achieve the same result.
		addressesByChain := make(map[uint64][]common.Address)
		domainsByChain := make(map[uint64]map[common.Address][]usdc_token_pool.DomainUpdate)

		for _, perChainInput := range input.ChainInputs {
			// For each chain input, add the address to the addressesByChain map.
			addressesByChain[perChainInput.ChainSelector] = append(addressesByChain[perChainInput.ChainSelector], perChainInput.Address)

			// Initialize the map for the chain selector if it doesn't exist yet to prevent a nil pointer dereference.
			if _, ok := domainsByChain[perChainInput.ChainSelector]; !ok {
				domainsByChain[perChainInput.ChainSelector] = make(map[common.Address][]usdc_token_pool.DomainUpdate)
			}

			// Map the provided DomainUpdate to the given address and chain selector.
			// By using append instead of assigning the slice directly, we can add multiple DomainUpdate structs to the same address and chain selector and allow them to all be executed at once by the sequence.
			domainsByChain[perChainInput.ChainSelector][perChainInput.Address] = append(domainsByChain[perChainInput.ChainSelector][perChainInput.Address], perChainInput.Domains...)
		}

		sequenceInput := sequences.SetDomainsSequenceInput{
			AddressesByChain: addressesByChain,
			DomainsByChain:   domainsByChain,
		}

		// Execute the sequence with the combined input
		report, err := cldf_ops.ExecuteSequence(e.OperationsBundle, sequences.USDCTokenPoolSetDomainsSequence, e.BlockChains, sequenceInput)
		if err != nil {
			return cldf.ChangesetOutput{}, err
		}

		batchOps = append(batchOps, report.Output.BatchOps...)
		reports = append(reports, report.ExecutionReports...)

		return changesets.NewOutputBuilder(e, nil).
			WithReports(reports).
			WithBatchOps(batchOps).
			Build(input.MCMS)
	}
}

func setDomainsVerify() func(cldf.Environment, SetDomainsInput) error {
	return func(e cldf.Environment, input SetDomainsInput) error {
		for _, perChainInput := range input.ChainInputs {
			if perChainInput.ChainSelector == 0 {
				return fmt.Errorf("chain selector must be provided for each chain input")
			}

			if exists := e.BlockChains.Exists(perChainInput.ChainSelector); !exists {
				return fmt.Errorf("chain with selector %d does not exist", perChainInput.ChainSelector)
			}
		}
		return nil
	}
}
