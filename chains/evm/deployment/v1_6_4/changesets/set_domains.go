package changesets

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_4/operations/usdc_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_4/sequences"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"

	evm_datastore_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/datastore"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
)

type SetDomainsInput struct {
	ChainInputs []SetDomainsPerChainInput
	MCMS        mcms.Input
}

type SetDomainsPerChainInput struct {
	ChainSelector uint64
	Domains       []usdc_token_pool.DomainUpdate
}

// This changeset is used to set the domains for a given token in the USDCTokenPool contract.
func SetDomainsChangeset(mcmsRegistry *changesets.MCMSReaderRegistry) cldf.ChangeSetV2[SetDomainsInput] {
	return cldf.CreateChangeSet(setDomainsApply(mcmsRegistry), setDomainsVerify(mcmsRegistry))
}

func setDomainsApply(mcmsRegistry *changesets.MCMSReaderRegistry) func(cldf.Environment, SetDomainsInput) (cldf.ChangesetOutput, error) {
	return func(e cldf.Environment, input SetDomainsInput) (cldf.ChangesetOutput, error) {
		batchOps := make([]mcms_types.BatchOperation, 0)
		reports := make([]cldf_ops.Report[any, any], 0)

		// Build the DomainsByChain map from per-chain inputs
		addressesByChain := make(map[uint64]common.Address)
		domainsByChain := make(map[uint64][]usdc_token_pool.DomainUpdate)
		for _, perChainInput := range input.ChainInputs {
			// For Each chain input, find the USDCTokenPool address from the datastore for the given chain selector
			// using
			usdc_token_pool_address, err := datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{
				Type:          datastore.ContractType(usdc_token_pool.ContractType),
				Version:       usdc_token_pool.Version,
				ChainSelector: perChainInput.ChainSelector,
			}, perChainInput.ChainSelector, evm_datastore_utils.ToEVMAddress)
			if err != nil {
				return cldf.ChangesetOutput{}, err
			}
			addressesByChain[perChainInput.ChainSelector] = usdc_token_pool_address
			domainsByChain[perChainInput.ChainSelector] = perChainInput.Domains
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

		return changesets.NewOutputBuilder(e, mcmsRegistry).
			WithReports(reports).
			WithBatchOps(batchOps).
			Build(input.MCMS)
	}
}

func setDomainsVerify(mcmsRegistry *changesets.MCMSReaderRegistry) func(cldf.Environment, SetDomainsInput) error {
	return func(e cldf.Environment, input SetDomainsInput) error {
		if err := input.MCMS.Validate(); err != nil {
			return err
		}
		return nil
	}
}
