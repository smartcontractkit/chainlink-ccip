package changesets

import (
	"errors"
	"fmt"
	"sync"

	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	"github.com/smartcontractkit/mcms"
	mcms_types "github.com/smartcontractkit/mcms/types"

	chain_selectors "github.com/smartcontractkit/chain-selectors"

	mcms_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
)

// MCMSReader is an interface for reading MCMS state from a chain type.
type MCMSReader interface {
	// GetChainMetadata returns the chain metadata for a given MCMS input.
	// Each chain family defines its own implementation of this method.
	GetChainMetadata(e deployment.Environment, chainSelector uint64, input mcms_utils.Input) (mcms_types.ChainMetadata, error)
	// GetTimelockRef returns the timelock contract address reference for a given MCMS input.
	GetTimelockRef(e deployment.Environment, chainSelector uint64, input mcms_utils.Input) (datastore.AddressRef, error)
	// GetMCMSRef returns the MCMS contract address reference for a given MCMS input.
	GetMCMSRef(e deployment.Environment, chainSelector uint64, input mcms_utils.Input) (datastore.AddressRef, error)
}

// MCMSReaderRegistry maintains a registry of MCMS readers.
type MCMSReaderRegistry struct {
	mu sync.Mutex
	m  map[string]MCMSReader
}

func newMCMSReaderRegistry() *MCMSReaderRegistry {
	return &MCMSReaderRegistry{
		m: make(map[string]MCMSReader),
	}
}

var (
	singletonRegistry *MCMSReaderRegistry
	once              sync.Once
)

// GetRegistry returns the global singleton instance.
// The first call creates the registry; subsequent calls return the same pointer.
func GetRegistry() *MCMSReaderRegistry {
	once.Do(func() {
		singletonRegistry = newMCMSReaderRegistry()
	})
	return singletonRegistry
}

// RegisterMCMSReader registers an MCMSReader for a specific chain family.
func (r *MCMSReaderRegistry) RegisterMCMSReader(chainFamily string, reader MCMSReader) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.m == nil {
		r.m = make(map[string]MCMSReader)
	}
	if _, exists := r.m[chainFamily]; !exists {
		r.m[chainFamily] = reader
	}
}

// GetMCMSReader retrieves an MCMSReader for a specific chain family.
func (r *MCMSReaderRegistry) GetMCMSReader(chainFamily string) (MCMSReader, bool) {
	r.mu.Lock()
	defer r.mu.Unlock()
	reader, ok := r.m[chainFamily]
	return reader, ok
}

// OutputBuilder helps construct a ChangesetOutput, including building an MCMS proposal if there are batch operations.
type OutputBuilder struct {
	registry        *MCMSReaderRegistry
	environment     deployment.Environment
	batchOps        []mcms_types.BatchOperation
	changesetOutput deployment.ChangesetOutput
}

// NewOutputBuilder creates a new OutputBuilder.
func NewOutputBuilder(e deployment.Environment, registry *MCMSReaderRegistry) *OutputBuilder {
	return &OutputBuilder{
		registry:        registry,
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

	timelockAddresses, err := b.getTimelockAddresses(input, b.batchOps)
	if err != nil {
		return deployment.ChangesetOutput{}, fmt.Errorf("failed to get timelock addresses: %w", err)
	}
	chainMetadata, err := b.getChainMetadata(input, b.batchOps)
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

// getTimelockAddresses resolves the timelock contract addresses for each chain selector in the list of batch operations.
func (b *OutputBuilder) getTimelockAddresses(
	input mcms_utils.Input,
	ops []mcms_types.BatchOperation,
) (map[mcms_types.ChainSelector]string, error) {
	timelocks := make(map[mcms_types.ChainSelector]string)
	for _, op := range ops {
		if _, exists := timelocks[op.ChainSelector]; exists {
			continue // Already resolved timelock for this chain selector
		}
		family, err := chain_selectors.GetSelectorFamily(uint64(op.ChainSelector))
		if err != nil {
			return nil, fmt.Errorf("failed to get chain family for chain selector %d: %w", op.ChainSelector, err)
		}
		reader, ok := b.registry.GetMCMSReader(family)
		if !ok {
			return nil, fmt.Errorf("no MCMS reader registered for chain family '%s'", family)
		}
		timelockRef, err := reader.GetTimelockRef(b.environment, uint64(op.ChainSelector), input)
		if err != nil {
			return nil, fmt.Errorf("failed to get timelock ref for chain with selector %d: %w", op.ChainSelector, err)
		}
		if err != nil {
			return nil, fmt.Errorf("failed to resolve timelock ref on chain with selector %d: %w", op.ChainSelector, err)
		}
		timelocks[op.ChainSelector] = timelockRef.Address
	}

	return timelocks, nil
}

// getChainMetadata fetches the current chain metadata (e.g. starting op count, mcm address) for each chain selector in the list of batch operations.
func (b *OutputBuilder) getChainMetadata(
	input mcms_utils.Input,
	ops []mcms_types.BatchOperation,
) (map[mcms_types.ChainSelector]mcms_types.ChainMetadata, error) {
	metadata := make(map[mcms_types.ChainSelector]mcms_types.ChainMetadata)
	for _, op := range ops {
		if _, ok := metadata[op.ChainSelector]; ok {
			continue // Already fetched metadata for this chain selector
		}
		family, err := chain_selectors.GetSelectorFamily(uint64(op.ChainSelector))
		if err != nil {
			return nil, fmt.Errorf("failed to get chain family for chain selector %d: %w", op.ChainSelector, err)
		}
		reader, ok := b.registry.GetMCMSReader(family)
		if !ok {
			return nil, fmt.Errorf("no MCMS reader registered for chain family '%s'", family)
		}
		chainMetadata, err := reader.GetChainMetadata(b.environment, uint64(op.ChainSelector), input)
		if err != nil {
			return nil, fmt.Errorf("failed to get MCMS chain metadata for chain with selector %d: %w", op.ChainSelector, err)
		}
		metadata[op.ChainSelector] = chainMetadata
	}

	return metadata, nil
}
