package deploy

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	chain_selectors "github.com/smartcontractkit/chain-selectors"
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
			// if partial refs are provided, resolve to full refs
			for i, contractRef := range perChainInputs.ContractRef {
				fullRef, err := datastore_utils.FindAndFormatRef(e.DataStore, contractRef, perChainInputs.ChainSelector, datastore_utils.FullRef)
				if err != nil {
					return cldf.ChangesetOutput{}, err
				}
				perChainInputs.ContractRef[i] = fullRef
			}
			chainBatchOps, chainReports, err := transferAndAcceptOwnership(e, adapter, perChainInputs, input.MCMS)
			if err != nil {
				return cldf.ChangesetOutput{}, err
			}
			batchOps = append(batchOps, chainBatchOps...)
			reports = append(reports, chainReports...)
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

func TransferToTimelock(chainSel uint64, e *cldf.Environment, mcmsInput mcms.Input, addressRefs []datastore.AddressRef, ownerOverride map[datastore.ContractType]string) ([]mcms_types.BatchOperation, []cldf_ops.Report[any, any], error) {
	mcmsRegistry := changesets.GetRegistry()
	transferOwnershipReg := GetTransferOwnershipRegistry()
	family, err := chain_selectors.GetSelectorFamily(chainSel)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get chain family for selector %d: %w", chainSel, err)
	}
	mcmsReader, ok := mcmsRegistry.GetMCMSReader(family)
	if !ok {
		return nil, nil, fmt.Errorf("no MCMS reader registered for chain family '%s'", family)
	}
	timelockRef, err := mcmsReader.GetTimelockRef(*e, chainSel, mcmsInput)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get timelock ref for chain %d: %w", chainSel, err)
	}

	adapter, err := transferOwnershipReg.GetAdapterByChainSelector(chainSel, MCMSVersion)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get transfer ownership adapter for chain %d: %w", chainSel, err)
	}
	proposedOwner := timelockRef.Address

	ownerShipInputByProposedOwner := make(map[string]TransferOwnershipPerChainInput)
	for _, contractRef := range addressRefs {
		if override, exists := ownerOverride[contractRef.Type]; exists {
			proposedOwner = override
		}
		input, exists := ownerShipInputByProposedOwner[proposedOwner]
		if !exists {
			input = TransferOwnershipPerChainInput{
				ChainSelector: chainSel,
				ContractRef:   []datastore.AddressRef{contractRef},
				ProposedOwner: proposedOwner,
			}
		} else {
			input.ContractRef = append(input.ContractRef, contractRef)
		}
		ownerShipInputByProposedOwner[proposedOwner] = input
	}
	var batchOps []mcms_types.BatchOperation
	var reports []cldf_ops.Report[any, any]
	for _, ownershipInput := range ownerShipInputByProposedOwner {
		ops, rep, err := transferAndAcceptOwnership(*e, adapter, ownershipInput, mcmsInput)
		if err != nil {
			return nil, nil, err
		}
		batchOps = append(batchOps, ops...)
		reports = append(reports, rep...)
	}

	return batchOps, reports, nil
}

// transferAndAcceptOwnership executes the transfer ownership sequence via MCMS and,
// if needed, the accept ownership sequence for the given contracts on a single chain.
func transferAndAcceptOwnership(
	e cldf.Environment,
	adapter TransferOwnershipAdapter,
	input TransferOwnershipPerChainInput,
	mcmsInput mcms.Input,
) ([]mcms_types.BatchOperation, []cldf_ops.Report[any, any], error) {
	batchOps := make([]mcms_types.BatchOperation, 0)
	reports := make([]cldf_ops.Report[any, any], 0)

	if err := adapter.InitializeTimelockAddress(e, mcmsInput); err != nil {
		return nil, nil, fmt.Errorf("failed to initialize timelock address for chain %d: %w", input.ChainSelector, err)
	}

	transferReport, err := cldf_ops.ExecuteSequence(e.OperationsBundle, adapter.SequenceTransferOwnershipViaMCMS(), e.BlockChains, input)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to transfer ownership on chain %d: %w", input.ChainSelector, err)
	}
	batchOps = append(batchOps, transferReport.Output.BatchOps...)
	reports = append(reports, transferReport.ExecutionReports...)

	needAccept, err := adapter.ShouldAcceptOwnershipWithTransferOwnership(e, input)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to check accept ownership on chain %d: %w", input.ChainSelector, err)
	}
	if needAccept {
		acceptReport, err := cldf_ops.ExecuteSequence(e.OperationsBundle, adapter.SequenceAcceptOwnership(), e.BlockChains, input)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to accept ownership on chain %d: %w", input.ChainSelector, err)
		}
		batchOps = append(batchOps, acceptReport.Output.BatchOps...)
		reports = append(reports, acceptReport.ExecutionReports...)
	}
	return batchOps, reports, nil
}
