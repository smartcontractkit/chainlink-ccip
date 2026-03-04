package changesets

import (
	"fmt"

	chain_selectors "github.com/smartcontractkit/chain-selectors"
	solseq "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_0/sequences" // TODO import from 1.6.1 when it is fully implemented
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_1/sequences"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

type RMNRemoteSetEventAuthoritiesChangesetInput struct {
	ChainSelector uint64     `json:"chainSelector"`
	MCMS          mcms.Input `json:"mcms"`
}

func RMNRemoteSetEventAuthoritiesChangeset() cldf.ChangeSetV2[RMNRemoteSetEventAuthoritiesChangesetInput] {
	return cldf.CreateChangeSet(rmnSetEventAuthoritiesApply, rmnSetEventAuthoritiesVerify)
}

func rmnSetEventAuthoritiesVerify(env cldf.Environment, input RMNRemoteSetEventAuthoritiesChangesetInput) error {
	if err := deployment.IsValidChainSelector(input.ChainSelector); err != nil {
		return fmt.Errorf("invalid chain selector: %d - %w", input.ChainSelector, err)
	}
	if !env.BlockChains.Exists(input.ChainSelector) {
		return fmt.Errorf("chain with selector %d does not exist", input.ChainSelector)
	}

	if err := input.MCMS.Validate(); err != nil {
		return fmt.Errorf("invalid MCMS configuration: %w", err)
	}

	return nil
}

func rmnSetEventAuthoritiesApply(e cldf.Environment, input RMNRemoteSetEventAuthoritiesChangesetInput) (cldf.ChangesetOutput, error) {
	reports := make([]cldf_ops.Report[any, any], 0)

	// Prepare sequence input
	seqInput := sequences.RMNRemoteSetEventAuthoritiesSequenceInput{
		DataStore: e.DataStore,
		Selector:  input.ChainSelector,
	}

	report, err := cldf_ops.ExecuteSequence(e.OperationsBundle, sequences.SetRMNRemoteEventAuthorities, e.BlockChains, seqInput)
	if err != nil {
		return cldf.ChangesetOutput{}, fmt.Errorf("failed to set rmn event authorities: %w", err)
	}

	reports = append(reports, report.ExecutionReports...)

	// Create the datastore with the addresses from the report
	ds := datastore.NewMemoryDataStore()
	for _, addr := range report.Output.Addresses {
		if err := ds.Addresses().Add(addr); err != nil {
			return cldf.ChangesetOutput{}, fmt.Errorf("failed to add address to datastore: %w", err)
		}
	}

	mcmsRegistry := changesets.GetRegistry()
	mcmsRegistry.RegisterMCMSReader(chain_selectors.FamilySolana, &solseq.SolanaAdapter{})

	return changesets.
		NewOutputBuilder(e, mcmsRegistry).
		WithReports(reports).
		WithDataStore(ds).
		WithBatchOps(report.Output.BatchOps).
		Build(input.MCMS)
}
