package changesets

import (
	"errors"
	"fmt"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	"github.com/smartcontractkit/mcms"
	mcms_types "github.com/smartcontractkit/mcms/types"
)

// OutputBuilder helps construct a ChangesetOutput, including building an MCMS proposal if there are write operations.
// Should be kept chain-family agnostic in case we want to move it out of evm-specific package later.
// Even call.WriteOutput is not EVM-specific, could potentially extract as a standard.
type OutputBuilder struct {
	writeOutputs    []contract.WriteOutput
	changesetOutput deployment.ChangesetOutput
}

// MCMSParams holds configuration for building an MCMS proposal.
type MCMSParams struct {
	// Description is a human-readable description of the proposal.
	Description string
	// OverridePreviousRoot indicates whether to override the root of the MCMS contract.
	OverridePreviousRoot bool
	// ValidUntil is a unix timestamp indicating when the proposal expires.
	// Root can't be set or executed after this time.
	ValidUntil uint32
	// TimelockDelay is the amount of time each operation in the proposal must wait before it can be executed.
	TimelockDelay mcms_types.Duration
	// TimelockAction is the action to perform on the timelock contract (schedule, bypass, or cancel).
	TimelockAction mcms_types.TimelockAction
	// TimelockAddresses is a map of chain selectors to timelock contract addresses.
	TimelockAddresses map[mcms_types.ChainSelector]string
	// ChainMetadata is optional metadata to include for each chain in the proposal.
	// Includes MCM address & starting op count.
	ChainMetadata map[mcms_types.ChainSelector]mcms_types.ChainMetadata
}

// NewOutputBuilder creates a new OutputBuilder.
func NewOutputBuilder() *OutputBuilder {
	return &OutputBuilder{
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
func (b *OutputBuilder) Build(params MCMSParams) (deployment.ChangesetOutput, error) {
	ops := b.convertWriteOutputsToBatchOperations()
	if ops == nil || len(ops) == 0 {
		// No write operations to include in MCMS proposal
		return b.changesetOutput, nil
	}
	proposal, err := mcms.NewTimelockProposalBuilder().
		SetVersion("v1").
		SetDescription(params.Description).
		SetOverridePreviousRoot(params.OverridePreviousRoot).
		SetValidUntil(params.ValidUntil).
		SetDelay(params.TimelockDelay).
		SetAction(params.TimelockAction).
		SetOperations(ops).
		SetTimelockAddresses(params.TimelockAddresses).
		SetChainMetadata(params.ChainMetadata).
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
		if outs.Executed {
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
