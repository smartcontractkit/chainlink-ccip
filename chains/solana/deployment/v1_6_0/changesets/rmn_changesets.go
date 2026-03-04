package changesets

import (
	"fmt"

	"github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_0/sequences"
	_ "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_0/sequences"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
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
	if err := cldf.IsValidChainSelector(input.ChainSelector); err != nil {
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

	return changesets.
		NewOutputBuilder(e, changesets.GetRegistry()).
		WithReports(reports).
		WithDataStore(ds).
		WithBatchOps(report.Output.BatchOps).
		Build(input.MCMS)
}
