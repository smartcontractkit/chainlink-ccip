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

		// Build the DomainsByChain map from per-chain inputs
		addressesByChain := make(map[uint64]common.Address)
		domainsByChain := make(map[uint64][]usdc_token_pool.DomainUpdate)
		for _, perChainInput := range input.ChainInputs {
			// Since the Token Pool may be of type CCTP V1 or CCTP V2, the address must be specified for each chain input.
			addressesByChain[perChainInput.ChainSelector] = perChainInput.Address
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
