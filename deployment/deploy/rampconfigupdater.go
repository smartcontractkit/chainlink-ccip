package deploy

import (
	"fmt"
	"sync"

	"github.com/Masterminds/semver/v3"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"

	"github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
)

var (
	singletonRampConfigUpdaterRegistry *RampConfigUpdaterRegistry
	rampConfigUpdaterOnce              sync.Once
)

// SetRampConfigInput is the input for the RampConfigApplier sequence. It carries the contract
// metadata from ConfigImporter.SequenceImportConfig() so that static config, dynamic config and
// dest/source chain config can be applied to onRamp and offRamp contracts.
type SetRampConfigInput struct {
	ChainSelector uint64
	ImportOutput  []datastore.ContractMetadata
}

// RampConfigApplier applies imported config (static, dynamic, dest chain config) to onRamp and offRamp contracts.
type RampConfigApplier interface {
	DeriveConfigImporterVersions() []*semver.Version
	SequenceSetRampConfig() *cldf_ops.Sequence[SetRampConfigInput, sequences.OnChainOutput, chain.BlockChains]
}

// RampConfigUpdaterRegistry holds ConfigImporter and RampConfigApplier per chain family+version for the UpdateRampConfig changeset.
type RampConfigUpdaterRegistry struct {
	ConfigImporter    map[string]ConfigImporter
	RampConfigApplier map[string]RampConfigApplier
	mu                sync.Mutex
}

func (r *RampConfigUpdaterRegistry) RegisterConfigImporter(family string, version *semver.Version, importer ConfigImporter) {
	r.mu.Lock()
	defer r.mu.Unlock()
	id := utils.NewRegistererID(family, version)
	if _, exists := r.ConfigImporter[id]; !exists {
		r.ConfigImporter[id] = importer
	}
}

func (r *RampConfigUpdaterRegistry) GetConfigImporter(chainsel uint64, version *semver.Version) (ConfigImporter, bool) {
	r.mu.Lock()
	defer r.mu.Unlock()
	id := utils.NewIDFromSelector(chainsel, version)
	importer, ok := r.ConfigImporter[id]
	return importer, ok
}

func (r *RampConfigUpdaterRegistry) RegisterRampConfigApplier(family string, version *semver.Version, applier RampConfigApplier) {
	r.mu.Lock()
	defer r.mu.Unlock()
	id := utils.NewRegistererID(family, version)
	if _, exists := r.RampConfigApplier[id]; !exists {
		r.RampConfigApplier[id] = applier
	}
}

func (r *RampConfigUpdaterRegistry) GetRampConfigApplier(chainsel uint64, version *semver.Version) (RampConfigApplier, bool) {
	r.mu.Lock()
	defer r.mu.Unlock()
	id := utils.NewIDFromSelector(chainsel, version)
	applier, ok := r.RampConfigApplier[id]
	return applier, ok
}

func newRampConfigUpdaterRegistry() *RampConfigUpdaterRegistry {
	return &RampConfigUpdaterRegistry{
		ConfigImporter:    make(map[string]ConfigImporter),
		RampConfigApplier: make(map[string]RampConfigApplier),
	}
}

// GetRampConfigUpdaterRegistry returns the global singleton RampConfigUpdaterRegistry.
func GetRampConfigUpdaterRegistry() *RampConfigUpdaterRegistry {
	rampConfigUpdaterOnce.Do(func() {
		singletonRampConfigUpdaterRegistry = newRampConfigUpdaterRegistry()
	})
	return singletonRampConfigUpdaterRegistry
}

// UpdateRampConfigInput is the input for the UpdateRampConfig changeset.
// It imports config from chain (via ConfigImporter) and sets onRamp and offRamp's
// static config, dynamic config and dest chain config (onRamp) / source chain config (offRamp).
type UpdateRampConfigInput struct {
	Chains map[uint64]UpdateRampConfigInputPerChain
	MCMS   mcms.Input
}

// UpdateRampConfigInputPerChain is the per-chain input for UpdateRampConfig.
type UpdateRampConfigInputPerChain struct {
	// ImportRampConfigFromVersions specifies adapter versions to use for importing config from chain.
	// Config is read from existing onRamp/offRamp contracts and then re-applied (e.g. after migration).
	ImportRampConfigFromVersions []*semver.Version
	// RampsVersion is the version of the RampConfigApplier to use for setting config on ramps.
	RampsVersion *semver.Version
}

// UpdateRampConfigChangeset creates a changeset that imports config from chain and sets onRamp and offRamp's
// static config, dynamic config and dest chain config (onRamp) / source chain config (offRamp).
// It follows the same pattern as UpdateFeeQuoterChangeset: optionally populate config from existing contracts,
// then apply that config to the ramp contracts.
func UpdateRampConfigChangeset(registry *RampConfigUpdaterRegistry, mcmsRegistry *changesets.MCMSReaderRegistry) cldf.ChangeSetV2[UpdateRampConfigInput] {
	return cldf.CreateChangeSet(updateRampConfigApply(registry, mcmsRegistry), updateRampConfigVerify())
}

func updateRampConfigVerify() func(cldf.Environment, UpdateRampConfigInput) error {
	return func(e cldf.Environment, input UpdateRampConfigInput) error {
		for chainSel, perChainInput := range input.Chains {
			if !e.BlockChains.Exists(chainSel) {
				return fmt.Errorf("chain with selector %d not found in environment", chainSel)
			}
			if perChainInput.RampsVersion == nil {
				return fmt.Errorf("ramps version is required for chain selector %d", chainSel)
			}
			if len(perChainInput.ImportRampConfigFromVersions) == 0 {
				return fmt.Errorf("at least one import config version is required for chain selector %d", chainSel)
			}
		}
		return nil
	}
}

func updateRampConfigApply(registry *RampConfigUpdaterRegistry, mcmsRegistry *changesets.MCMSReaderRegistry) func(cldf.Environment, UpdateRampConfigInput) (cldf.ChangesetOutput, error) {
	return func(e cldf.Environment, input UpdateRampConfigInput) (cldf.ChangesetOutput, error) {
		batchOps := make([]mcms_types.BatchOperation, 0)
		reports := make([]cldf_ops.Report[any, any], 0)

		for chainSel, perChainInput := range input.Chains {
			rampConfigApplier, ok := registry.GetRampConfigApplier(chainSel, perChainInput.RampsVersion)
			if !ok {
				return cldf.ChangesetOutput{}, utils.ErrNoAdapterRegistered("RampConfigApplier", perChainInput.RampsVersion)
			}

			// Loop through all ImportRampConfigFromVersions and consolidate contract metadata
			contractMeta := make([]datastore.ContractMetadata, 0)
			for _, version := range perChainInput.ImportRampConfigFromVersions {
				configImporter, ok := registry.GetConfigImporter(chainSel, version)
				if !ok {
					return cldf.ChangesetOutput{}, utils.ErrNoAdapterRegistered("ConfigImporter", version)
				}
				err := configImporter.InitializeAdapter(e, chainSel)
				if err != nil {
					return cldf.ChangesetOutput{}, fmt.Errorf("failed to initialize config importer for chain %d (version %s): %w", chainSel, version, err)
				}
				supportedTokensPerRemoteChain, err := configImporter.SupportedTokensPerRemoteChain(e, chainSel)
				if err != nil {
					return cldf.ChangesetOutput{}, fmt.Errorf("failed to get supported tokens per remote chain for chain %d: %w", chainSel, err)
				}
				connectedChains, err := configImporter.ConnectedChains(e, chainSel)
				if err != nil {
					return cldf.ChangesetOutput{}, fmt.Errorf("failed to get connected chains for chain %d: %w", chainSel, err)
				}
				importReport, err := cldf_ops.ExecuteSequence(e.OperationsBundle, configImporter.SequenceImportConfig(), e.BlockChains, ImportConfigPerChainInput{
					ChainSelector:        chainSel,
					RemoteChains:         connectedChains,
					TokensPerRemoteChain: supportedTokensPerRemoteChain,
				})
				if err != nil {
					return cldf.ChangesetOutput{}, fmt.Errorf("failed to import config for ramps on chain %d (version %s): %w", chainSel, version, err)
				}
				if len(importReport.Output.Metadata.Contracts) == 0 {
					return cldf.ChangesetOutput{}, fmt.Errorf("no contract metadata returned from import config for chain %d (version %s)", chainSel, version)
				}
				contractMeta = append(contractMeta, importReport.Output.Metadata.Contracts...)
				reports = append(reports, importReport.ExecutionReports...)
			}

			// Run SequenceSetRampConfig once with all consolidated contract metadata
			setReport, err := cldf_ops.ExecuteSequence(e.OperationsBundle, rampConfigApplier.SequenceSetRampConfig(), e.BlockChains, SetRampConfigInput{
				ChainSelector: chainSel,
				ImportOutput:  contractMeta,
			})
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("failed to set ramp config on chain %d: %w", chainSel, err)
			}

			batchOps = append(batchOps, setReport.Output.BatchOps...)
			reports = append(reports, setReport.ExecutionReports...)
		}

		return changesets.NewOutputBuilder(e, mcmsRegistry).
			WithReports(reports).
			WithBatchOps(batchOps).
			Build(input.MCMS)
	}
}
