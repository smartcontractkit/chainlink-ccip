package deploy

import (
	"github.com/Masterminds/semver/v3"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"

	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
)

type TransferOwnershipInput struct {
	ChainInputs    []TransferOwnershipPerChainInput
	AdapterVersion *semver.Version
	MCMS           mcms.Input
}

type TransferOwnershipPerChainInput struct {
	ChainSelector uint64
	ContractRef   []datastore.AddressRef
	CurrentOwner  string
	ProposedOwner string
}

func AcceptOwnershipChangeset(cr *TransferOwnershipAdapterRegistry, mcmsRegistry *changesets.MCMSReaderRegistry) cldf.ChangeSetV2[TransferOwnershipInput] {
	return cldf.CreateChangeSet(acceptOwnershipApply(cr, mcmsRegistry), acceptOwnershipVerify(cr, mcmsRegistry))
}

func acceptOwnershipApply(cr *TransferOwnershipAdapterRegistry, mcmsRegistry *changesets.MCMSReaderRegistry) func(cldf.Environment, TransferOwnershipInput) (cldf.ChangesetOutput, error) {
	return func(e cldf.Environment, input TransferOwnershipInput) (cldf.ChangesetOutput, error) {
		batchOps := make([]mcms_types.BatchOperation, 0)
		reports := make([]cldf_ops.Report[any, any], 0)
		for _, perChainInputs := range input.ChainInputs {
			adapter, err := cr.GetAdapterByChainSelector(perChainInputs.ChainSelector, input.AdapterVersion)
			if err != nil {
				return cldf.ChangesetOutput{}, err
			}
			err = adapter.InitializeTimelockAddress(e, input.MCMS)
			if err != nil {
				return cldf.ChangesetOutput{}, err
			}
			// if partial refs are provided, resolve to full refs
			for i, contractRef := range perChainInputs.ContractRef {
				fullRef, err := datastore_utils.FindAndFormatRef(e.DataStore, contractRef, perChainInputs.ChainSelector, datastore_utils.FullRef)
				if err != nil {
					return cldf.ChangesetOutput{}, err
				}
				perChainInputs.ContractRef[i] = fullRef
			}
			report, err := cldf_ops.ExecuteSequence(e.OperationsBundle, adapter.SequenceAcceptOwnership(), e.BlockChains, perChainInputs)
			if err != nil {
				return cldf.ChangesetOutput{}, err
			}
			batchOps = append(batchOps, report.Output.BatchOps...)
			reports = append(reports, report.ExecutionReports...)
		}
		return changesets.NewOutputBuilder(e, mcmsRegistry).
			WithReports(reports).
			WithBatchOps(batchOps).
			Build(input.MCMS)
	}
}

func acceptOwnershipVerify(cr *TransferOwnershipAdapterRegistry, mcmsRegistry *changesets.MCMSReaderRegistry) func(cldf.Environment, TransferOwnershipInput) error {
	return func(e cldf.Environment, input TransferOwnershipInput) error {
		if err := input.MCMS.Validate(); err != nil {
			return err
		}
		return nil
	}
}

func TransferOwnershipChangeset(cr *TransferOwnershipAdapterRegistry, mcmsRegistry *changesets.MCMSReaderRegistry) cldf.ChangeSetV2[TransferOwnershipInput] {
	return cldf.CreateChangeSet(transferOwnershipApply(cr, mcmsRegistry), transferOwnershipVerify(cr, mcmsRegistry))
}

func transferOwnershipApply(cr *TransferOwnershipAdapterRegistry, mcmsRegistry *changesets.MCMSReaderRegistry) func(cldf.Environment, TransferOwnershipInput) (cldf.ChangesetOutput, error) {
	return func(e cldf.Environment, input TransferOwnershipInput) (cldf.ChangesetOutput, error) {
		batchOps := make([]mcms_types.BatchOperation, 0)
		reports := make([]cldf_ops.Report[any, any], 0)
		for _, perChainInputs := range input.ChainInputs {
			adapter, err := cr.GetAdapterByChainSelector(perChainInputs.ChainSelector, input.AdapterVersion)
			if err != nil {
				return cldf.ChangesetOutput{}, err
			}
			err = adapter.InitializeTimelockAddress(e, input.MCMS)
			if err != nil {
				return cldf.ChangesetOutput{}, err
			}
			// if partial refs are provided, resolve to full refs
			for i, contractRef := range perChainInputs.ContractRef {
				fullRef, err := datastore_utils.FindAndFormatRef(e.DataStore, contractRef, perChainInputs.ChainSelector, datastore_utils.FullRef)
				if err != nil {
					return cldf.ChangesetOutput{}, err
				}
				perChainInputs.ContractRef[i] = fullRef
			}
			report, err := cldf_ops.ExecuteSequence(e.OperationsBundle, adapter.SequenceTransferOwnershipViaMCMS(), e.BlockChains, perChainInputs)
			if err != nil {
				return cldf.ChangesetOutput{}, err
			}
			batchOps = append(batchOps, report.Output.BatchOps...)
			reports = append(reports, report.ExecutionReports...)
			needAcceptOwnership, err := adapter.ShouldAcceptOwnershipWithTransferOwnership(e, perChainInputs)
			if err != nil {
				return cldf.ChangesetOutput{}, err
			}
			if needAcceptOwnership {
				report, err = cldf_ops.ExecuteSequence(e.OperationsBundle, adapter.SequenceAcceptOwnership(), e.BlockChains, perChainInputs)
				if err != nil {
					return cldf.ChangesetOutput{}, err
				}
				batchOps = append(batchOps, report.Output.BatchOps...)
				reports = append(reports, report.ExecutionReports...)
			}
		}
		return changesets.NewOutputBuilder(e, mcmsRegistry).
			WithReports(reports).
			WithBatchOps(batchOps).
			Build(input.MCMS)
	}
}

func transferOwnershipVerify(cr *TransferOwnershipAdapterRegistry, mcmsRegistry *changesets.MCMSReaderRegistry) func(cldf.Environment, TransferOwnershipInput) error {
	return func(e cldf.Environment, input TransferOwnershipInput) error {
		if err := input.MCMS.Validate(); err != nil {
			return err
		}
		return nil
	}
}
