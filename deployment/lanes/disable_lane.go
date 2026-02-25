package lanes

import (
	"fmt"
	"sync"

	"github.com/Masterminds/semver/v3"
	chain_selectors "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"
)

// DisableRemoteChainInput provides the local and remote chain info
// needed to disable a lane on the local chain.
type DisableRemoteChainInput struct {
	LocalChainSelector  uint64
	RemoteChainSelector uint64
	OnRamp              []byte
	OffRamp             []byte
	Router              []byte
	FeeQuoter           []byte
}

// DisableLanePair identifies two chains whose bidirectional lane should be disabled.
type DisableLanePair struct {
	ChainA  uint64
	ChainB  uint64
	Version *semver.Version
}

// DisableLaneConfig is the input for the DisableLane changeset.
type DisableLaneConfig struct {
	Lanes []DisableLanePair
	MCMS  mcms.Input
}

// DisableLaneAdapter defines chain-family-specific logic for disabling a lane.
type DisableLaneAdapter interface {
	DisableRemoteChain() *cldf_ops.Sequence[DisableRemoteChainInput, sequences.OnChainOutput, cldf_chain.BlockChains]

	GetOnRampAddress(ds datastore.DataStore, chainSelector uint64) ([]byte, error)
	GetOffRampAddress(ds datastore.DataStore, chainSelector uint64) ([]byte, error)
	GetRouterAddress(ds datastore.DataStore, chainSelector uint64) ([]byte, error)
	GetFQAddress(ds datastore.DataStore, chainSelector uint64) ([]byte, error)
}

// DisableLaneAdapterRegistry maintains a registry of DisableLaneAdapters.
type DisableLaneAdapterRegistry struct {
	mu sync.Mutex
	m  map[disableLaneAdapterID]DisableLaneAdapter
}

type disableLaneAdapterID string

func newDisableLaneAdapterRegistry() *DisableLaneAdapterRegistry {
	return &DisableLaneAdapterRegistry{
		m: make(map[disableLaneAdapterID]DisableLaneAdapter),
	}
}

func (r *DisableLaneAdapterRegistry) RegisterDisableLaneAdapter(chainFamily string, version *semver.Version, adapter DisableLaneAdapter) {
	id := newDisableLaneAdapterID(chainFamily, version)

	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.m[id]; exists {
		panic(fmt.Errorf("DisableLaneAdapter '%s %s' already registered", chainFamily, version))
	}
	r.m[id] = adapter
}

func (r *DisableLaneAdapterRegistry) GetDisableLaneAdapter(chainFamily string, version *semver.Version) (DisableLaneAdapter, bool) {
	id := newDisableLaneAdapterID(chainFamily, version)

	r.mu.Lock()
	defer r.mu.Unlock()

	adapter, ok := r.m[id]
	return adapter, ok
}

var (
	singletonDisableLaneRegistry *DisableLaneAdapterRegistry
	onceDisableLane              sync.Once
)

func GetDisableLaneAdapterRegistry() *DisableLaneAdapterRegistry {
	onceDisableLane.Do(func() {
		singletonDisableLaneRegistry = newDisableLaneAdapterRegistry()
	})
	return singletonDisableLaneRegistry
}

func newDisableLaneAdapterID(chainFamily string, version *semver.Version) disableLaneAdapterID {
	return disableLaneAdapterID(fmt.Sprintf("%s-%s", chainFamily, version.String()))
}

// DisableLane returns a changeset that disables bidirectional CCIP lanes.
func DisableLane(
	disableRegistry *DisableLaneAdapterRegistry,
	mcmsRegistry *changesets.MCMSReaderRegistry,
) cldf.ChangeSetV2[DisableLaneConfig] {
	return cldf.CreateChangeSet(
		makeDisableApply(disableRegistry, mcmsRegistry),
		makeDisableVerify(),
	)
}

func makeDisableVerify() func(cldf.Environment, DisableLaneConfig) error {
	return func(_ cldf.Environment, _ DisableLaneConfig) error {
		return nil
	}
}

func makeDisableApply(
	disableRegistry *DisableLaneAdapterRegistry,
	mcmsRegistry *changesets.MCMSReaderRegistry,
) func(cldf.Environment, DisableLaneConfig) (cldf.ChangesetOutput, error) {
	return func(e cldf.Environment, cfg DisableLaneConfig) (cldf.ChangesetOutput, error) {
		batchOps := make([]mcms_types.BatchOperation, 0)
		reports := make([]cldf_ops.Report[any, any], 0)

		for _, lane := range cfg.Lanes {
			chainAFamily, err := chain_selectors.GetSelectorFamily(lane.ChainA)
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("failed to get chain family for selector %d: %w", lane.ChainA, err)
			}
			chainBFamily, err := chain_selectors.GetSelectorFamily(lane.ChainB)
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("failed to get chain family for selector %d: %w", lane.ChainB, err)
			}

			chainAAdapter, exists := disableRegistry.GetDisableLaneAdapter(chainAFamily, lane.Version)
			if !exists {
				return cldf.ChangesetOutput{}, fmt.Errorf("no DisableLaneAdapter registered for chain family '%s' version %s", chainAFamily, lane.Version)
			}
			chainBAdapter, exists := disableRegistry.GetDisableLaneAdapter(chainBFamily, lane.Version)
			if !exists {
				return cldf.ChangesetOutput{}, fmt.Errorf("no DisableLaneAdapter registered for chain family '%s' version %s", chainBFamily, lane.Version)
			}

			chainADef := &ChainDefinition{Selector: lane.ChainA}
			chainBDef := &ChainDefinition{Selector: lane.ChainB}

			err = populateDisableAddresses(e.DataStore, chainADef, chainAAdapter)
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("error fetching addresses for chain %d: %w", lane.ChainA, err)
			}
			err = populateDisableAddresses(e.DataStore, chainBDef, chainBAdapter)
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("error fetching addresses for chain %d: %w", lane.ChainB, err)
			}

			type disablePair struct {
				local   *ChainDefinition
				remote  *ChainDefinition
				adapter DisableLaneAdapter
			}
			for _, pair := range []disablePair{
				{local: chainADef, remote: chainBDef, adapter: chainAAdapter},
				{local: chainBDef, remote: chainADef, adapter: chainBAdapter},
			} {
				report, err := cldf_ops.ExecuteSequence(
					e.OperationsBundle,
					pair.adapter.DisableRemoteChain(),
					e.BlockChains,
					DisableRemoteChainInput{
						LocalChainSelector:  pair.local.Selector,
						RemoteChainSelector: pair.remote.Selector,
						OnRamp:              pair.local.OnRamp,
						OffRamp:             pair.local.OffRamp,
						Router:              pair.local.Router,
						FeeQuoter:           pair.local.FeeQuoter,
					},
				)
				if err != nil {
					return cldf.ChangesetOutput{}, fmt.Errorf("failed to disable remote chain %d on chain %d: %w", pair.remote.Selector, pair.local.Selector, err)
				}
				batchOps = append(batchOps, report.Output.BatchOps...)
				reports = append(reports, report.ExecutionReports...)
			}
		}

		return changesets.NewOutputBuilder(e, mcmsRegistry).
			WithReports(reports).
			WithBatchOps(batchOps).
			Build(cfg.MCMS)
	}
}

func populateDisableAddresses(ds datastore.DataStore, chainDef *ChainDefinition, adapter DisableLaneAdapter) error {
	var err error
	chainDef.OnRamp, err = adapter.GetOnRampAddress(ds, chainDef.Selector)
	if err != nil {
		return fmt.Errorf("error fetching onramp address for chain %d: %w", chainDef.Selector, err)
	}
	chainDef.OffRamp, err = adapter.GetOffRampAddress(ds, chainDef.Selector)
	if err != nil {
		return fmt.Errorf("error fetching offramp address for chain %d: %w", chainDef.Selector, err)
	}
	chainDef.FeeQuoter, err = adapter.GetFQAddress(ds, chainDef.Selector)
	if err != nil {
		return fmt.Errorf("error fetching fee quoter address for chain %d: %w", chainDef.Selector, err)
	}
	chainDef.Router, err = adapter.GetRouterAddress(ds, chainDef.Selector)
	if err != nil {
		return fmt.Errorf("error fetching router address for chain %d: %w", chainDef.Selector, err)
	}
	return nil
}
