package changesets

import (
	"errors"
	"fmt"

	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	"github.com/smartcontractkit/mcms"
	mcms_types "github.com/smartcontractkit/mcms/types"

	chain_selectors "github.com/smartcontractkit/chain-selectors"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	mcms_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
)

// MCMSReader is an interface for reading MCMS state from a chain type.
type MCMSReader interface {
	// GetChainMetadata returns the chain metadata for a given MCM contract reference.
	// Each chain family defines its own implementation of this method.
	GetChainMetadata(e deployment.Environment, mcmRef datastore.AddressRef) (mcms_types.ChainMetadata, error)
}

var registeredMCMSReaders map[string]MCMSReader

// RegisterMCMSReader registers an MCMSReader for a specific chain family.
func RegisterMCMSReader(chainFamily string, reader MCMSReader) {
	if registeredMCMSReaders == nil {
		registeredMCMSReaders = make(map[string]MCMSReader)
	}
	if _, exists := registeredMCMSReaders[chainFamily]; exists {
		panic(fmt.Sprintf("MCMS reader already registered for chain family: %s", chainFamily))
	}
	registeredMCMSReaders[chainFamily] = reader
}

// OutputBuilder helps construct a ChangesetOutput, including building an MCMS proposal if there are batch operations.
type OutputBuilder struct {
	environment     deployment.Environment
	batchOps        []mcms_types.BatchOperation
	changesetOutput deployment.ChangesetOutput
}

// NewOutputBuilder creates a new OutputBuilder.
func NewOutputBuilder(e deployment.Environment) *OutputBuilder {
	return &OutputBuilder{
		environment:     e,
		changesetOutput: deployment.ChangesetOutput{},
	}
}

// WithReports sets the reports on the ChangesetOutput.
func (b *OutputBuilder) WithReports(reports []operations.Report[any, any]) *OutputBuilder {
	b.changesetOutput.Reports = reports
	return b
}

// WithDataStore sets the datastore on the ChangesetOutput.
func (b *OutputBuilder) WithDataStore(ds datastore.MutableDataStore) *OutputBuilder {
	b.changesetOutput.DataStore = ds
	return b
}

// WithBatchOps sets the batch operations on the OutputBuilder.
func (b *OutputBuilder) WithBatchOps(ops []mcms_types.BatchOperation) *OutputBuilder {
	// Filter out any batch operations that have no transactions.
	filteredOps := make([]mcms_types.BatchOperation, 0, len(ops))
	for _, op := range ops {
		if len(op.Transactions) > 0 {
			filteredOps = append(filteredOps, op)
		}
	}

	b.batchOps = filteredOps
	return b
}

// Build constructs the final ChangesetOutput, including building an MCMS proposal if there are write operations that have not been executed.
func (b *OutputBuilder) Build(input mcms_utils.Input) (deployment.ChangesetOutput, error) {
	if len(b.batchOps) == 0 {
		// No write operations to include in MCMS proposal
		return b.changesetOutput, nil
	}

	timelockAddresses, err := b.getTimelockAddresses(input.TimelockAddressRef, b.batchOps)
	if err != nil {
		return deployment.ChangesetOutput{}, fmt.Errorf("failed to get timelock addresses: %w", err)
	}
	chainMetadata, err := b.getChainMetadata(input.MCMSAddressRef, b.batchOps)
	if err != nil {
		return deployment.ChangesetOutput{}, fmt.Errorf("failed to get chain metadata: %w", err)
	}

	proposal, err := mcms.NewTimelockProposalBuilder().
		SetVersion("v1").
		SetDescription(input.Description).
		SetOverridePreviousRoot(input.OverridePreviousRoot).
		SetValidUntil(input.ValidUntil).
		SetDelay(input.TimelockDelay).
		SetAction(input.TimelockAction).
		SetOperations(b.batchOps).
		SetTimelockAddresses(timelockAddresses).
		SetChainMetadata(chainMetadata).
		Build()
	if err != nil {
		return deployment.ChangesetOutput{}, fmt.Errorf("failed to build MCMS proposal: %w", err)
	}
	if proposal == nil {
		return deployment.ChangesetOutput{}, errors.New("unexpected nil MCMS proposal")
	}
	b.changesetOutput.MCMSTimelockProposals = []mcms.TimelockProposal{*proposal}

	return b.changesetOutput, nil
}

func (b *OutputBuilder) getTimelockAddresses(
	timelockRef datastore.AddressRef,
	ops []mcms_types.BatchOperation,
) (map[mcms_types.ChainSelector]string, error) {
	timelocks := make(map[mcms_types.ChainSelector]string)
	for _, op := range ops {
		fullTimelockRef, err := datastore_utils.FindAndFormatRef(
			b.environment.DataStore,
			timelockRef,
			uint64(op.ChainSelector),
			datastore_utils.FullRef,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to resolve timelock ref on chain with selector %d: %w", op.ChainSelector, err)
		}
		timelocks[op.ChainSelector] = fullTimelockRef.Address
	}

	return timelocks, nil
}

func (b *OutputBuilder) getChainMetadata(
	mcmRef datastore.AddressRef,
	ops []mcms_types.BatchOperation,
) (map[mcms_types.ChainSelector]mcms_types.ChainMetadata, error) {
	metadata := make(map[mcms_types.ChainSelector]mcms_types.ChainMetadata)
	for _, op := range ops {
		fullMCMRef, err := datastore_utils.FindAndFormatRef(
			b.environment.DataStore,
			mcmRef,
			uint64(op.ChainSelector),
			datastore_utils.FullRef,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to resolve mcm ref on chain with selector %d: %w", op.ChainSelector, err)
		}

		family, err := chain_selectors.GetSelectorFamily(uint64(op.ChainSelector))
		if err != nil {
			return nil, fmt.Errorf("failed to get chain family for chain selector %d: %w", op.ChainSelector, err)
		}
		reader, ok := registeredMCMSReaders[family]
		if !ok {
			return nil, fmt.Errorf("no MCMS reader registered for chain family '%s'", family)
		}

		chainMetadata, err := reader.GetChainMetadata(b.environment, fullMCMRef)
		if err != nil {
			return nil, fmt.Errorf("failed to get current op count from MCMS at address %s on chain with selector %d: %w", fullMCMRef.Address, op.ChainSelector, err)
		}
		metadata[op.ChainSelector] = chainMetadata
	}

	return metadata, nil
}
