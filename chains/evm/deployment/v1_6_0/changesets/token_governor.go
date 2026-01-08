package changesets

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	mcms_types "github.com/smartcontractkit/mcms/types"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/sequences"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

// DeployTokenGovernor returns a changeset that deploys a TokenGovernor contract
func DeployTokenGovernorChangeset() cldf.ChangeSetV2[sequences.DeployTokenGovernorInput] {
	return cldf.CreateChangeSet(deployTokenGovernorApply(), deployTokenGovernorVerify())
}

func deployTokenGovernorApply() func(cldf.Environment, sequences.DeployTokenGovernorInput) (cldf.ChangesetOutput, error) {
	return func(e cldf.Environment, input sequences.DeployTokenGovernorInput) (cldf.ChangesetOutput, error) {
		reports := make([]cldf_ops.Report[any, any], 0)

		// Prepare sequence input
		input.ExistingDataStore = e.DataStore

		report, err := cldf_ops.ExecuteSequence(e.OperationsBundle, sequences.DeployTokenGovernor, e.BlockChains, input)
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

func deployTokenGovernorVerify() func(cldf.Environment, sequences.DeployTokenGovernorInput) error {
	return func(e cldf.Environment, input sequences.DeployTokenGovernorInput) error {
		if input.ChainSelector == 0 {
			return fmt.Errorf("chain selector must be provided")
		}
		if !e.BlockChains.Exists(input.ChainSelector) {
			return fmt.Errorf("chain with selector %d does not exist", input.ChainSelector)
		}
		if input.Token == "" {
			return fmt.Errorf("token address must be provided")
		}
		if !common.IsHexAddress(input.Token) {
			return fmt.Errorf("token address is not a valid hex address: %s", input.Token)
		}
		if input.InitialDefaultAdmin != "" && !common.IsHexAddress(input.InitialDefaultAdmin) {
			return fmt.Errorf("initial default admin is not a valid hex address: %s", input.InitialDefaultAdmin)
		}
		return nil
	}
}

// GrantRoleChangeset returns a changeset that grants a role to an account on the TokenGovernor contract
func GrantRoleChangeset() cldf.ChangeSetV2[sequences.TokenGovernorRoleInput] {
	return cldf.CreateChangeSet(grantRoleApply(), grantRoleVerify())
}

func grantRoleApply() func(cldf.Environment, sequences.TokenGovernorRoleInput) (cldf.ChangesetOutput, error) {
	return func(e cldf.Environment, input sequences.TokenGovernorRoleInput) (cldf.ChangesetOutput, error) {
		batchOps := make([]mcms_types.BatchOperation, 0)
		reports := make([]cldf_ops.Report[any, any], 0)

		// Prepare sequence input
		input.ExistingDataStore = e.DataStore

		report, err := cldf_ops.ExecuteSequence(e.OperationsBundle, sequences.GrantRole, e.BlockChains, input)
		if err != nil {
			return cldf.ChangesetOutput{}, fmt.Errorf("failed to grant role: %w", err)
		}

		batchOps = append(batchOps, report.Output.BatchOps...)
		reports = append(reports, report.ExecutionReports...)

		mcmsInput := mcms.Input{}
		if input.MCMS != nil {
			mcmsInput = *input.MCMS
		}

		return changesets.NewOutputBuilder(e, nil).
			WithReports(reports).
			WithBatchOps(batchOps).
			Build(mcmsInput)
	}
}

func grantRoleVerify() func(cldf.Environment, sequences.TokenGovernorRoleInput) error {
	return func(e cldf.Environment, input sequences.TokenGovernorRoleInput) error {
		if len(input.Tokens) == 0 {
			return fmt.Errorf("tokens map must be provided")
		}
		for chainSelector, tokenMap := range input.Tokens {
			if !e.BlockChains.Exists(chainSelector) {
				return fmt.Errorf("chain with selector %d does not exist", chainSelector)
			}
			for tokenSymbol, roleConfig := range tokenMap {
				if tokenSymbol == "" {
					return fmt.Errorf("token symbol must be provided for chain %d", chainSelector)
				}
				if roleConfig.Role == "" {
					return fmt.Errorf("role must be provided for token %s on chain %d", tokenSymbol, chainSelector)
				}
				// Validate the role string using the Role type
				if !sequences.Role(roleConfig.Role).IsValid() {
					return fmt.Errorf("invalid role '%s' for token %s on chain %d", roleConfig.Role, tokenSymbol, chainSelector)
				}
				if roleConfig.Account == (common.Address{}) {
					return fmt.Errorf("account address must be provided for token %s on chain %d", tokenSymbol, chainSelector)
				}
			}
		}
		return nil
	}
}

// RevokeRoleChangeset returns a changeset that revokes a role from an account on the TokenGovernor contract
func RevokeRoleChangeset() cldf.ChangeSetV2[sequences.TokenGovernorRoleInput] {
	return cldf.CreateChangeSet(revokeRoleApply(), revokeRoleVerify())
}

func revokeRoleApply() func(cldf.Environment, sequences.TokenGovernorRoleInput) (cldf.ChangesetOutput, error) {
	return func(e cldf.Environment, input sequences.TokenGovernorRoleInput) (cldf.ChangesetOutput, error) {
		batchOps := make([]mcms_types.BatchOperation, 0)
		reports := make([]cldf_ops.Report[any, any], 0)

		// Prepare sequence input
		input.ExistingDataStore = e.DataStore

		report, err := cldf_ops.ExecuteSequence(e.OperationsBundle, sequences.RevokeRole, e.BlockChains, input)
		if err != nil {
			return cldf.ChangesetOutput{}, fmt.Errorf("failed to revoke role: %w", err)
		}

		batchOps = append(batchOps, report.Output.BatchOps...)
		reports = append(reports, report.ExecutionReports...)

		mcmsInput := mcms.Input{}
		if input.MCMS != nil {
			mcmsInput = *input.MCMS
		}

		return changesets.NewOutputBuilder(e, nil).
			WithReports(reports).
			WithBatchOps(batchOps).
			Build(mcmsInput)
	}
}

func revokeRoleVerify() func(cldf.Environment, sequences.TokenGovernorRoleInput) error {
	// perform same verification as grant role
	return grantRoleVerify()
}

// RenounceRoleChangeset returns a changeset that renounces a role from an account on the TokenGovernor contract
func RenounceRoleChangeset() cldf.ChangeSetV2[sequences.TokenGovernorRoleInput] {
	return cldf.CreateChangeSet(renounceRoleApply(), renounceRoleVerify())
}

func renounceRoleApply() func(cldf.Environment, sequences.TokenGovernorRoleInput) (cldf.ChangesetOutput, error) {
	return func(e cldf.Environment, input sequences.TokenGovernorRoleInput) (cldf.ChangesetOutput, error) {
		batchOps := make([]mcms_types.BatchOperation, 0)
		reports := make([]cldf_ops.Report[any, any], 0)

		// Prepare sequence input
		input.ExistingDataStore = e.DataStore

		report, err := cldf_ops.ExecuteSequence(e.OperationsBundle, sequences.RenounceRole, e.BlockChains, input)
		if err != nil {
			return cldf.ChangesetOutput{}, fmt.Errorf("failed to renounce role: %w", err)
		}

		batchOps = append(batchOps, report.Output.BatchOps...)
		reports = append(reports, report.ExecutionReports...)

		mcmsInput := mcms.Input{}
		if input.MCMS != nil {
			mcmsInput = *input.MCMS
		}

		return changesets.NewOutputBuilder(e, nil).
			WithReports(reports).
			WithBatchOps(batchOps).
			Build(mcmsInput)
	}
}

func renounceRoleVerify() func(cldf.Environment, sequences.TokenGovernorRoleInput) error {
	return grantRoleVerify()
}

// TransferOwnershipChangeset returns a changeset that transfers ownership of the TokenGovernor contract
func TransferOwnershipChangeset() cldf.ChangeSetV2[sequences.TokenGovernorOwnershipInput] {
	return cldf.CreateChangeSet(transferOwnershipApply(), transferOwnershipVerify())
}

// AcceptOwnershipChangeset returns a changeset that accepts ownership of the TokenGovernor contract
func AcceptOwnershipChangeset() cldf.ChangeSetV2[sequences.TokenGovernorOwnershipInput] {
	return cldf.CreateChangeSet(acceptOwnershipApply(), acceptOwnershipVerify())
}

func transferOwnershipApply() func(cldf.Environment, sequences.TokenGovernorOwnershipInput) (cldf.ChangesetOutput, error) {
	return func(e cldf.Environment, input sequences.TokenGovernorOwnershipInput) (cldf.ChangesetOutput, error) {
		batchOps := make([]mcms_types.BatchOperation, 0)
		reports := make([]cldf_ops.Report[any, any], 0)

		// Prepare sequence input
		input.ExistingDataStore = e.DataStore

		report, err := cldf_ops.ExecuteSequence(e.OperationsBundle, sequences.TransferOwnership, e.BlockChains, input)
		if err != nil {
			return cldf.ChangesetOutput{}, fmt.Errorf("failed to transfer ownership: %w", err)
		}

		batchOps = append(batchOps, report.Output.BatchOps...)
		reports = append(reports, report.ExecutionReports...)

		mcmsInput := mcms.Input{}
		if input.MCMS != nil {
			mcmsInput = *input.MCMS
		}

		return changesets.NewOutputBuilder(e, nil).
			WithReports(reports).
			WithBatchOps(batchOps).
			Build(mcmsInput)
	}
}

func transferOwnershipVerify() func(cldf.Environment, sequences.TokenGovernorOwnershipInput) error {
	return func(e cldf.Environment, input sequences.TokenGovernorOwnershipInput) error {
		if len(input.Tokens) == 0 {
			return fmt.Errorf("tokens map must be provided")
		}
		for chainSelector, tokenMap := range input.Tokens {
			if !e.BlockChains.Exists(chainSelector) {
				return fmt.Errorf("chain with selector %d does not exist", chainSelector)
			}
			for tokenSymbol, newOwner := range tokenMap {
				if tokenSymbol == "" {
					return fmt.Errorf("token symbol must be provided for chain %d", chainSelector)
				}
				if newOwner == (common.Address{}) {
					return fmt.Errorf("new owner address must be provided for token %s on chain %d", tokenSymbol, chainSelector)
				}
			}
		}
		return nil
	}
}

func acceptOwnershipApply() func(cldf.Environment, sequences.TokenGovernorOwnershipInput) (cldf.ChangesetOutput, error) {
	return func(e cldf.Environment, input sequences.TokenGovernorOwnershipInput) (cldf.ChangesetOutput, error) {
		batchOps := make([]mcms_types.BatchOperation, 0)
		reports := make([]cldf_ops.Report[any, any], 0)

		// Prepare sequence input
		input.ExistingDataStore = e.DataStore

		report, err := cldf_ops.ExecuteSequence(e.OperationsBundle, sequences.AcceptOwnership, e.BlockChains, input)
		if err != nil {
			return cldf.ChangesetOutput{}, fmt.Errorf("failed to accept ownership: %w", err)
		}

		batchOps = append(batchOps, report.Output.BatchOps...)
		reports = append(reports, report.ExecutionReports...)

		mcmsInput := mcms.Input{}
		if input.MCMS != nil {
			mcmsInput = *input.MCMS
		}

		return changesets.NewOutputBuilder(e, nil).
			WithReports(reports).
			WithBatchOps(batchOps).
			Build(mcmsInput)
	}
}

func acceptOwnershipVerify() func(cldf.Environment, sequences.TokenGovernorOwnershipInput) error {
	return transferOwnershipVerify()
}

// BeginDefaultAdminTransferChangeset returns a changeset that begins the transfer of default admin role
func BeginDefaultAdminTransferChangeset() cldf.ChangeSetV2[sequences.TokenGovernorOwnershipInput] {
	return cldf.CreateChangeSet(beginDefaultAdminTransferApply(), beginDefaultAdminTransferVerify())
}

func beginDefaultAdminTransferApply() func(cldf.Environment, sequences.TokenGovernorOwnershipInput) (cldf.ChangesetOutput, error) {
	return func(e cldf.Environment, input sequences.TokenGovernorOwnershipInput) (cldf.ChangesetOutput, error) {
		batchOps := make([]mcms_types.BatchOperation, 0)
		reports := make([]cldf_ops.Report[any, any], 0)

		// Prepare sequence input
		input.ExistingDataStore = e.DataStore

		report, err := cldf_ops.ExecuteSequence(e.OperationsBundle, sequences.BeginDefaultAdminTransfer, e.BlockChains, input)
		if err != nil {
			return cldf.ChangesetOutput{}, fmt.Errorf("failed to begin default admin transfer: %w", err)
		}

		batchOps = append(batchOps, report.Output.BatchOps...)
		reports = append(reports, report.ExecutionReports...)

		mcmsInput := mcms.Input{}
		if input.MCMS != nil {
			mcmsInput = *input.MCMS
		}

		return changesets.NewOutputBuilder(e, nil).
			WithReports(reports).
			WithBatchOps(batchOps).
			Build(mcmsInput)
	}
}

func beginDefaultAdminTransferVerify() func(cldf.Environment, sequences.TokenGovernorOwnershipInput) error {
	return func(e cldf.Environment, input sequences.TokenGovernorOwnershipInput) error {
		if len(input.Tokens) == 0 {
			return fmt.Errorf("tokens map must be provided")
		}
		for chainSelector, tokenMap := range input.Tokens {
			if !e.BlockChains.Exists(chainSelector) {
				return fmt.Errorf("chain with selector %d does not exist", chainSelector)
			}
			for tokenSymbol, admin := range tokenMap {
				if tokenSymbol == "" {
					return fmt.Errorf("token symbol must be provided for chain %d", chainSelector)
				}
				if admin == (common.Address{}) {
					return fmt.Errorf("new admin owner address must be provided for token %s on chain %d", tokenSymbol, chainSelector)
				}
			}
		}
		return nil
	}
}

// AcceptDefaultAdminTransferChangeset returns a changeset that accepts the pending default admin role transfer
func AcceptDefaultAdminTransferChangeset() cldf.ChangeSetV2[sequences.TokenGovernorOwnershipInput] {
	return cldf.CreateChangeSet(acceptDefaultAdminTransferApply(), acceptDefaultAdminTransferVerify())
}

func acceptDefaultAdminTransferApply() func(cldf.Environment, sequences.TokenGovernorOwnershipInput) (cldf.ChangesetOutput, error) {
	return func(e cldf.Environment, input sequences.TokenGovernorOwnershipInput) (cldf.ChangesetOutput, error) {
		batchOps := make([]mcms_types.BatchOperation, 0)
		reports := make([]cldf_ops.Report[any, any], 0)

		// Prepare sequence input
		input.ExistingDataStore = e.DataStore

		report, err := cldf_ops.ExecuteSequence(e.OperationsBundle, sequences.AcceptDefaultAdminTransfer, e.BlockChains, input)
		if err != nil {
			return cldf.ChangesetOutput{}, fmt.Errorf("failed to accept default admin transfer: %w", err)
		}

		batchOps = append(batchOps, report.Output.BatchOps...)
		reports = append(reports, report.ExecutionReports...)

		mcmsInput := mcms.Input{}
		if input.MCMS != nil {
			mcmsInput = *input.MCMS
		}

		return changesets.NewOutputBuilder(e, nil).
			WithReports(reports).
			WithBatchOps(batchOps).
			Build(mcmsInput)
	}
}

func acceptDefaultAdminTransferVerify() func(cldf.Environment, sequences.TokenGovernorOwnershipInput) error {
	return beginDefaultAdminTransferVerify()
}
