package token_pools

import (
	"github.com/ethereum/go-ethereum/common"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_4/operations/usdc_token_pool_ops"
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
	Domains       []usdc_token_pool_ops.DomainUpdate
}

func SetDomainsChangeset(mcmsRegistry *changesets.MCMSReaderRegistry) cldf.ChangeSetV2[SetDomainsInput] {
	return cldf.CreateChangeSet(setDomainsApply(mcmsRegistry), setDomainsVerify(mcmsRegistry))
}

func setDomainsApply(mcmsRegistry *changesets.MCMSReaderRegistry) func(cldf.Environment, SetDomainsInput) (cldf.ChangesetOutput, error) {
	return func(e cldf.Environment, input SetDomainsInput) (cldf.ChangesetOutput, error) {
		batchOps := make([]mcms_types.BatchOperation, 0)
		reports := make([]cldf_ops.Report[any, any], 0)

		// Build the DomainsByChain map from per-chain inputs
		domainsByChain := make(map[uint64][]usdc_token_pool_ops.DomainUpdate)
		var address common.Address
		for _, perChainInput := range input.ChainInputs {
			if address == (common.Address{}) {
				address = perChainInput.Address
			}
			domainsByChain[perChainInput.ChainSelector] = perChainInput.Domains
		}

		// Execute the sequence with the combined input
		sequenceInput := sequences.SetDomainsSequenceInput{
			Address:        address,
			DomainsByChain: domainsByChain,
		}

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
