package sequences

import (
	"fmt"

	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"

	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
)

// WithChainSelector is implemented by configs that specify a target chain selector.
type WithChainSelector interface {
	ChainSelector() uint64
}

// Metadata defines any metadata pertaining to the sequence execution.
type Metadata struct {
	// Contracts defines any metadata pertaining to contracts that the sequence deployed.
	Contracts []datastore.ContractMetadata
	// Chain defines any metadata pertaining to the chain that the sequence operated against.
	Chain *datastore.ChainMetadata
	// Env defines any metadata pertaining to the environment that the sequence deployed to.
	Env *datastore.EnvMetadata
}

// OnChainOutput is a standard output type for sequences that deploy contracts on-chain and perform write operations.
type OnChainOutput struct {
	// Addresses are the contract addresses that the sequence deployed.
	Addresses []datastore.AddressRef
	// Metadata defines any metadata pertaining to the sequence execution.
	// This metadata gets pushed to the datastore alongside the addresses.
	// NOTE: This metadata gets upserted into the datastore, meaning that keys will be entirely overridden.
	// - Contract metadata key = address + chain selector
	// - Chain metadata key = chain selector
	// - 1 Env metadata record per environment
	// So, if writing to a particular key, ensure that you are writing all required information.
	Metadata Metadata
	// BatchOps are operations that must be executed via MCMS.
	// Order is important and should be preserved during construction of the proposal.
	// Each batch operation is executed atomically.
	BatchOps []mcms_types.BatchOperation
}

func RunAndMergeSequence[IN any](
	b operations.Bundle,
	chains cldf_chain.BlockChains,
	seq *operations.Sequence[IN, OnChainOutput, cldf_chain.BlockChains],
	input IN,
	agg OnChainOutput,
) (OnChainOutput, error) {
	report, err := operations.ExecuteSequence(b, seq, chains, input)
	if err != nil {
		return agg, fmt.Errorf("failed to execute %s: %w", seq.ID(), err)
	}
	agg.BatchOps = append(agg.BatchOps, report.Output.BatchOps...)
	agg.Addresses = append(agg.Addresses, report.Output.Addresses...)
	agg.Metadata.Contracts = append(agg.Metadata.Contracts, report.Output.Metadata.Contracts...)
	if report.Output.Metadata.Chain != nil {
		agg.Metadata.Chain = report.Output.Metadata.Chain
	}
	if report.Output.Metadata.Env != nil {
		agg.Metadata.Env = report.Output.Metadata.Env
	}
	return agg, nil
}

func WriteMetadataToDatastore(ds datastore.MutableDataStore, metadata Metadata) error {
	for _, contract := range metadata.Contracts {
		if err := ds.ContractMetadata().Upsert(contract); err != nil {
			return fmt.Errorf("failed to add contract metadata: %w", err)
		}
	}
	if metadata.Chain != nil {
		if err := ds.ChainMetadata().Upsert(*metadata.Chain); err != nil {
			return fmt.Errorf("failed to add chain metadata: %w", err)
		}
	}
	if metadata.Env != nil {
		if err := ds.EnvMetadata().Set(*metadata.Env); err != nil {
			return fmt.Errorf("failed to add env metadata: %w", err)
		}
	}
	return nil
}
