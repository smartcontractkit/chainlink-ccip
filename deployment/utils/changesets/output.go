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
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/operations/contract"
)

// MCMSReader is an interface for reading MCMS state from a chain type.
type MCMSReader interface {
	OpCount(e deployment.Environment, chainSelector uint64, mcmAddress string) (uint64, error)
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

// OutputBuilder helps construct a ChangesetOutput, including building an MCMS proposal if there are write operations.
// Should be kept chain-family agnostic in case we want to move it out of evm-specific package later.
// Even call.WriteOutput is not EVM-specific, could potentially extract as a standard.
type OutputBuilder struct {
	environment     deployment.Environment
	writeOutputs    []contract.WriteOutput
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

// WithWriteOutputs sets the write outputs on the OutputBuilder.
func (b *OutputBuilder) WithWriteOutputs(outs []contract.WriteOutput) *OutputBuilder {
	b.writeOutputs = outs
	return b
}

// Build constructs the final ChangesetOutput, including building an MCMS proposal if there are write operations that have not been executed.
func (b *OutputBuilder) Build(input mcms_utils.Input) (deployment.ChangesetOutput, error) {
	ops := b.convertWriteOutputsToBatchOperations()
	if ops == nil || len(ops) == 0 {
		// No write operations to include in MCMS proposal
		return b.changesetOutput, nil
	}

	timelockAddresses, err := b.getTimelockAddresses(input.TimelockAddressRef, ops)
	if err != nil {
		return deployment.ChangesetOutput{}, fmt.Errorf("failed to get timelock addresses: %w", err)
	}
	chainMetadata, err := b.getChainMetadata(input.MCMSAddressRef, ops)
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
		SetOperations(ops).
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

// TODO: Incorporate batch size?
func (b *OutputBuilder) convertWriteOutputsToBatchOperations() []mcms_types.BatchOperation {
	batchOps := make(map[uint64]mcms_types.BatchOperation)
	for _, outs := range b.writeOutputs {
		if outs.Executed() {
			continue // Skip executed transactions, should not be included in MCMS proposal
		}
		batchOp, exists := batchOps[outs.ChainSelector]
		if !exists {
			batchOps[outs.ChainSelector] = mcms_types.BatchOperation{
				ChainSelector: mcms_types.ChainSelector(outs.ChainSelector),
				Transactions:  []mcms_types.Transaction{outs.Tx},
			}
		} else {
			batchOp.Transactions = append(batchOp.Transactions, outs.Tx)
			batchOps[outs.ChainSelector] = batchOp
		}
	}
	var batchOpsSlice []mcms_types.BatchOperation
	for _, batchOps := range batchOps {
		batchOpsSlice = append(batchOpsSlice, batchOps)
	}
	return batchOpsSlice
}

func (b *OutputBuilder) getTimelockAddresses(
	timelockRef datastore.AddressRef,
	ops []mcms_types.BatchOperation,
) (map[mcms_types.ChainSelector]string, error) {
	timelocks := make(map[mcms_types.ChainSelector]string)
	for _, op := range ops {
		timelockRef.ChainSelector = uint64(op.ChainSelector)
		refs, err := datastore_utils.FindAndFormatEachRef(b.environment.DataStore, []datastore.AddressRef{timelockRef}, datastore_utils.FullRef)
		if err != nil {
			return nil, fmt.Errorf("failed to resolve timelock ref on chain with selector %d: %w", op.ChainSelector, err)
		}
		timelocks[op.ChainSelector] = refs[0].Address
	}

	return timelocks, nil
}

func (b *OutputBuilder) getChainMetadata(
	mcmRef datastore.AddressRef,
	ops []mcms_types.BatchOperation,
) (map[mcms_types.ChainSelector]mcms_types.ChainMetadata, error) {
	metadata := make(map[mcms_types.ChainSelector]mcms_types.ChainMetadata)
	for _, op := range ops {
		mcmRef.ChainSelector = uint64(op.ChainSelector)
		refs, err := datastore_utils.FindAndFormatEachRef(b.environment.DataStore, []datastore.AddressRef{mcmRef}, datastore_utils.FullRef)
		if err != nil {
			return nil, fmt.Errorf("failed to resolve mcm ref on chain with selector %d: %w", op.ChainSelector, err)
		}
		mcmAddress := refs[0].Address

		family, err := chain_selectors.GetSelectorFamily(uint64(op.ChainSelector))
		if err != nil {
			return nil, fmt.Errorf("failed to get chain family for chain selector %d: %w", op.ChainSelector, err)
		}
		reader, ok := registeredMCMSReaders[family]
		if !ok {
			return nil, fmt.Errorf("no MCMS reader registered for chain family '%s'", family)
		}

		opCount, err := reader.OpCount(b.environment, uint64(op.ChainSelector), mcmAddress)
		if err != nil {
			return nil, fmt.Errorf("failed to get current op count from MCMS at address %s on chain with selector %d: %w", mcmAddress, op.ChainSelector, err)
		}
		metadata[op.ChainSelector] = mcms_types.ChainMetadata{
			MCMAddress:      mcmAddress,
			StartingOpCount: opCount,
		}
	}

	return metadata, nil
}
