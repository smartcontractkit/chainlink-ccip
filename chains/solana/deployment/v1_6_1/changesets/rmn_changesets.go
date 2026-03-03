package changesets

import (
	"fmt"

	"github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_1/sequences"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

// DeployTokenGovernor returns a changeset that deploys a TokenGovernor contract
func RMNRemoteSetEventAuthoritiesSequenceChangeset() cldf.ChangeSetV2[sequences.RMNRemoteSetEventAuthoritiesSequenceInput] {
	return cldf.CreateChangeSet(rmnSetEventAuthoritiesApply(), rmnSetEventAuthoritiesVerify())
}

func rmnSetEventAuthoritiesVerify() func(cldf.Environment, sequences.RMNRemoteSetEventAuthoritiesSequenceInput) error {
	return func(env cldf.Environment, input sequences.RMNRemoteSetEventAuthoritiesSequenceInput) error {

		if err := deployment.IsValidChainSelector(input.ChainSelector); err != nil {
			return fmt.Errorf("invalid chain selector: %d - %w", input.ChainSelector, err)
		}
		if !env.BlockChains.Exists(input.ChainSelector) {
			return fmt.Errorf("chain with selector %d does not exist", input.ChainSelector)
		}

		// Validate UpgradeConfig
		// if err := input.UpgradeConfig.MCMS.ValidateSolana(env, input.ChainSelector); err != nil {
		// 	return fmt.Errorf("invalid MCMS configuration: %w", err)
		// }

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

		return cldf.ChangesetOutput{
			DataStore: ds,
			Reports:   reports,
		}, nil
	}
}
