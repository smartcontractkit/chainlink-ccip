package changesets

import (
	"fmt"

	"github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_1/sequences"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

// DeployTokenGovernor returns a changeset that deploys a TokenGovernor contract
func RMNRemoteSetEventAuthoritiesSequenceChangeset() cldf.ChangeSetV2[sequences.RMNRemoteSetEventAuthoritiesSequenceInput] {
	return cldf.CreateChangeSet(rmnSetEventAuthoritiesApply(), rmnSetEventAuthoritiesVerify())
}

func rmnSetEventAuthoritiesVerify() func(cldf.Environment, sequences.RMNRemoteSetEventAuthoritiesSequenceInput) error {
	return func(e cldf.Environment, input sequences.RMNRemoteSetEventAuthoritiesSequenceInput) error {
		// TODO: Implement verification logic to check that the event authorities were set correctly on the RMNRemote contract. This may involve querying the contract state and comparing it to the expected values.
		return nil
	}
}

func rmnSetEventAuthoritiesApply() func(cldf.Environment, sequences.RMNRemoteSetEventAuthoritiesSequenceInput) (cldf.ChangesetOutput, error) {
	return func(e cldf.Environment, input sequences.RMNRemoteSetEventAuthoritiesSequenceInput) (cldf.ChangesetOutput, error) {
		reports := make([]cldf_ops.Report[any, any], 0)

		// Prepare sequence input
		input.DataStore = e.DataStore

		report, err := cldf_ops.ExecuteSequence(e.OperationsBundle, sequences.SetRMNRemoteEventAuthorities, e.BlockChains, input)
		if err != nil {
			return cldf.ChangesetOutput{}, fmt.Errorf("failed to deploy TokenGovernor: %w", err)
		}

		reports = append(reports, report.ExecutionReports...)

		// Create the datastore with the addresses from the report
		ds := datastore.NewMemoryDataStore()
		for _, addr := range report.Output.Addresses {
			if err := ds.Addresses().Add(addr); err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("failed to add address to datastore: %w", err)
			}
		}

		return cldf.ChangesetOutput{
			DataStore: ds,
			Reports:   reports,
		}, nil
	}
}
