package deploy

import (
	"fmt"
	"sync"

	"github.com/Masterminds/semver/v3"
	chain_selectors "github.com/smartcontractkit/chain-selectors"
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
	singletonFQAndRampUpdaterRegistry *FQAndRampUpdaterRegistry
	fqupdaterOnce                     sync.Once
)

type UpdateFeeQuoterInput struct {
	Chains map[uint64]UpdateFeeQuoterInputPerChain
	MCMS   mcms.Input
}

type UpdateFeeQuoterInputPerChain struct {
	ImportFeeQuoterConfigFromVersions []*semver.Version
	FeeQuoterVersion                  *semver.Version
	RampsVersion                      *semver.Version
	SourceChainAddsToOffRamp          []uint64
}

type FeeQuoterUpdateInput struct {
	ChainSelector     uint64
	ExistingAddresses []datastore.AddressRef
	ContractMeta      []datastore.ContractMetadata
}

type SourceChainConfig struct {
	Router datastore.AddressRef
	OnRamp datastore.AddressRef
}

type UpdateRampsInput struct {
	ChainSelector     uint64
	FeeQuoterAddress  datastore.AddressRef
	OnRampAddressRef  datastore.AddressRef
	OffRampAddressRef datastore.AddressRef
	SourceChains      map[uint64]SourceChainConfig
}

// FeeQuoterUpdater provides methods to update FeeQuoter contract on a chain.
type FeeQuoterUpdater[FeeQUpdateArgs any] interface {
	SequenceFeeQuoterInputCreation() *cldf_ops.Sequence[FeeQuoterUpdateInput, FeeQUpdateArgs, chain.BlockChains]
	SequenceDeployOrUpdateFeeQuoter() *cldf_ops.Sequence[FeeQUpdateArgs, sequences.OnChainOutput, chain.BlockChains]
}

type RampUpdater interface {
	ResolveRampsInput(e cldf.Environment, input UpdateRampsInput) (UpdateRampsInput, error)
	SequenceUpdateRampsWithFeeQuoter() *cldf_ops.Sequence[UpdateRampsInput, sequences.OnChainOutput, chain.BlockChains]
}

type FQAndRampUpdaterRegistry struct {
	FeeQuoterUpdater map[string]FeeQuoterUpdater[any]
	RampUpdater      map[string]RampUpdater
	ConfigImporter   map[string]ConfigImporter
	mu               sync.Mutex
}

func (r *FQAndRampUpdaterRegistry) RegisterFeeQuoterUpdater(family string, version *semver.Version, updater FeeQuoterUpdater[any]) {
	r.mu.Lock()
	defer r.mu.Unlock()
	id := utils.NewRegistererID(family, version)
	if _, exists := r.FeeQuoterUpdater[id]; !exists {
		r.FeeQuoterUpdater[id] = updater
	}
}

func (r *FQAndRampUpdaterRegistry) RegisterRampUpdater(family string, version *semver.Version, updater RampUpdater) {
	r.mu.Lock()
	defer r.mu.Unlock()
	id := utils.NewRegistererID(family, version)
	if _, exists := r.RampUpdater[id]; !exists {
		r.RampUpdater[id] = updater
	}
}

func (r *FQAndRampUpdaterRegistry) RegisterConfigImporter(family string, version *semver.Version, importer ConfigImporter) {
	r.mu.Lock()
	defer r.mu.Unlock()
	id := utils.NewRegistererID(family, version)
	if _, exists := r.ConfigImporter[id]; !exists {
		r.ConfigImporter[id] = importer
	}
}

func (r *FQAndRampUpdaterRegistry) GetConfigImporter(chainsel uint64, version *semver.Version) (ConfigImporter, bool) {
	// Get the chain family from the chain selector
	family, err := chain_selectors.GetSelectorFamily(chainsel)
	if err != nil {
		return nil, false
	}
	r.mu.Lock()
	defer r.mu.Unlock()
	id := utils.NewRegistererID(family, version)
	importer, ok := r.ConfigImporter[id]
	return importer, ok
}

func (r *FQAndRampUpdaterRegistry) GetFeeQuoterUpdater(chainsel uint64, version *semver.Version) (FeeQuoterUpdater[any], bool) {
	// Get the chain family from the chain selector
	family, err := chain_selectors.GetSelectorFamily(chainsel)
	if err != nil {
		return nil, false
	}
	r.mu.Lock()
	defer r.mu.Unlock()
	id := utils.NewRegistererID(family, version)
	updater, ok := r.FeeQuoterUpdater[id]
	return updater, ok
}

func (r *FQAndRampUpdaterRegistry) GetRampUpdater(chainsel uint64, version *semver.Version) (RampUpdater, bool) {
	// Get the chain family from the chain selector
	family, err := chain_selectors.GetSelectorFamily(chainsel)
	if err != nil {
		return nil, false
	}
	r.mu.Lock()
	defer r.mu.Unlock()
	id := utils.NewRegistererID(family, version)
	updater, ok := r.RampUpdater[id]
	return updater, ok
}

func newFQUpdaterRegistry() *FQAndRampUpdaterRegistry {
	return &FQAndRampUpdaterRegistry{
		FeeQuoterUpdater: make(map[string]FeeQuoterUpdater[any]),
		RampUpdater:      make(map[string]RampUpdater),
		ConfigImporter:   make(map[string]ConfigImporter),
	}
}

func GetFQAndRampUpdaterRegistry() *FQAndRampUpdaterRegistry {
	fqupdaterOnce.Do(func() {
		singletonFQAndRampUpdaterRegistry = newFQUpdaterRegistry()
	})
	return singletonFQAndRampUpdaterRegistry
}

// UpdateFeeQuoterChangeset creates a changeset that updates FeeQuoter contracts on specified chains.
// It first optionally populates configuration values, then creates FeeQuoterUpdateInput,
// deploys or updates the FeeQuoter contract, and finally updates the Ramps contracts to use the new FeeQuoter address.
// If needed, it also updates OffRamp source chain configs ( specifically used when during updating feequoter only specific source chain needs to be added to offramp).
func UpdateFeeQuoterChangeset(fquRegistry *FQAndRampUpdaterRegistry, mcmsRegistry *changesets.MCMSReaderRegistry) cldf.ChangeSetV2[UpdateFeeQuoterInput] {
	return cldf.CreateChangeSet(updateFeeQuoterApply(fquRegistry, mcmsRegistry), updateFeeQuoterVerify(fquRegistry, mcmsRegistry))
}

func updateFeeQuoterVerify(fquRegistry *FQAndRampUpdaterRegistry, mcmsRegistry *changesets.MCMSReaderRegistry) func(cldf.Environment, UpdateFeeQuoterInput) error {
	return func(e cldf.Environment, input UpdateFeeQuoterInput) error {
		for chainSel, perChainInput := range input.Chains {
			if !e.BlockChains.Exists(chainSel) {
				return fmt.Errorf("chain with selector %d not found in environment", chainSel)
			}
			if perChainInput.FeeQuoterVersion == nil {
				return fmt.Errorf("fee quoter version is required for chain selector %d", chainSel)
			}
			if perChainInput.RampsVersion == nil {
				return fmt.Errorf("ramps version is required for chain selector %d", chainSel)
			}
		}
		return nil
	}
}

func updateFeeQuoterApply(fquRegistry *FQAndRampUpdaterRegistry, mcmsRegistry *changesets.MCMSReaderRegistry) func(cldf.Environment, UpdateFeeQuoterInput) (cldf.ChangesetOutput, error) {
	return func(e cldf.Environment, input UpdateFeeQuoterInput) (cldf.ChangesetOutput, error) {
		batchOps := make([]mcms_types.BatchOperation, 0)
		reports := make([]cldf_ops.Report[any, any], 0)
		addressRefs := make([]datastore.AddressRef, 0)
		contractMetadata := make([]datastore.ContractMetadata, 0)
		for chainSel, perChainInput := range input.Chains {
			fquUpdater, ok := fquRegistry.GetFeeQuoterUpdater(chainSel, perChainInput.FeeQuoterVersion)
			if !ok {
				return cldf.ChangesetOutput{}, utils.ErrNoAdapterRegistered("FeeQuoterUpdater", perChainInput.FeeQuoterVersion)
			}
			rampUpdater, ok := fquRegistry.GetRampUpdater(chainSel, perChainInput.RampsVersion)
			if !ok {
				return cldf.ChangesetOutput{}, utils.ErrNoAdapterRegistered("RampUpdater", perChainInput.RampsVersion)
			}
			contractMeta := make([]datastore.ContractMetadata, 0)
			for _, version := range perChainInput.ImportFeeQuoterConfigFromVersions {
				configImporter, ok := fquRegistry.GetConfigImporter(chainSel, version)
				if !ok {
					return cldf.ChangesetOutput{}, utils.ErrNoAdapterRegistered("ConfigImporter", perChainInput.FeeQuoterVersion)
				}
				err := configImporter.InitializeAdapter(e, chainSel)
				if err != nil {
					return cldf.ChangesetOutput{}, fmt.Errorf("failed to initialize config importer for chain %d: %w", chainSel, err)
				}
				supportedTokensPerRemoteChain, err := configImporter.SupportedTokensPerRemoteChain(e, chainSel)
				if err != nil {
					return cldf.ChangesetOutput{}, fmt.Errorf("failed to get supported tokens per remote chain for chain %d: %w", chainSel, err)
				}
				connectedChains, err := configImporter.ConnectedChains(e, chainSel)
				if err != nil {
					return cldf.ChangesetOutput{}, fmt.Errorf("failed to get connected chains for chain %d: %w", chainSel, err)
				}
				populateConfigReport, err := cldf_ops.ExecuteSequence(e.OperationsBundle, configImporter.SequenceImportConfig(), e.BlockChains, ImportConfigPerChainInput{
					ChainSelector:        chainSel,
					RemoteChains:         connectedChains,
					TokensPerRemoteChain: supportedTokensPerRemoteChain,
				})
				if err != nil {
					return cldf.ChangesetOutput{}, fmt.Errorf("failed to populate config for FeeQuoter on chain %d: %w", chainSel, err)
				}
				if len(populateConfigReport.Output.Metadata.Contracts) == 0 {
					return cldf.ChangesetOutput{}, fmt.Errorf("no contract metadata returned from populate config for FeeQuoter on chain %d", chainSel)
				}

				contractMeta = append(contractMeta, populateConfigReport.Output.Metadata.Contracts...)
				contractMetadata = append(contractMetadata, populateConfigReport.Output.Metadata.Contracts...)
			}
			// Create FeeQuoterUpdateInput
			reportFQInputCreation, err := cldf_ops.ExecuteSequence(e.OperationsBundle, fquUpdater.SequenceFeeQuoterInputCreation(), e.BlockChains, FeeQuoterUpdateInput{
				ChainSelector:     chainSel,
				ExistingAddresses: e.DataStore.Addresses().Filter(datastore.AddressRefByChainSelector(chainSel)),
				ContractMeta:      contractMeta,
			})
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("failed to create FeeQuoterUpdateInput for chain %d: %w", chainSel, err)
			}
			// Deploy or update FeeQuoter
			reportFQUpdate, err := cldf_ops.ExecuteSequence(e.OperationsBundle, fquUpdater.SequenceDeployOrUpdateFeeQuoter(), e.BlockChains, reportFQInputCreation.Output)
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("failed to deploy or update FeeQuoter for chain %d: %w", chainSel, err)
			}
			batchOps = append(batchOps, reportFQUpdate.Output.BatchOps...)
			addressRefs = append(addressRefs, reportFQUpdate.Output.Addresses...)
			reports = append(reports, reportFQUpdate.ExecutionReports...)
			if len(reportFQUpdate.Output.Addresses) == 0 {
				return cldf.ChangesetOutput{}, fmt.Errorf("no FeeQuoter address returned for chain %d", chainSel)
			}
			// Update Ramps with new FeeQuoter address
			// fetch the address refs
			feeQuoterAddrRef := reportFQUpdate.Output.Addresses[len(reportFQUpdate.Output.Addresses)-1]
			if perChainInput.RampsVersion != nil {
				rampsInput := UpdateRampsInput{
					ChainSelector:    chainSel,
					FeeQuoterAddress: feeQuoterAddrRef,
					SourceChains:     make(map[uint64]SourceChainConfig),
				}
				for _, srcChainSel := range perChainInput.SourceChainAddsToOffRamp {
					rampsInput.SourceChains[srcChainSel] = SourceChainConfig{}
				}
				// Resolve Ramps input
				resolvedRampsInput, err := rampUpdater.ResolveRampsInput(e, rampsInput)
				if err != nil {
					return cldf.ChangesetOutput{}, fmt.Errorf("failed to resolve ramps input for chain %d: %w", chainSel, err)
				}
				// Execute Ramps update sequence
				reportRampsUpdate, err := cldf_ops.ExecuteSequence(e.OperationsBundle, rampUpdater.SequenceUpdateRampsWithFeeQuoter(), e.BlockChains, resolvedRampsInput)
				if err != nil {
					return cldf.ChangesetOutput{}, fmt.Errorf("failed to update ramps with FeeQuoter for chain %d: %w", chainSel, err)
				}
				batchOps = append(batchOps, reportRampsUpdate.Output.BatchOps...)
				addressRefs = append(addressRefs, reportRampsUpdate.Output.Addresses...)
				reports = append(reports, reportRampsUpdate.ExecutionReports...)
			}
		}
		// Prepare datastore with all address refs
		ds := datastore.NewMemoryDataStore()
		for _, addrRef := range addressRefs {
			if err := ds.Addresses().Add(addrRef); err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("failed to add %s %s with address %s on chain with selector %d to datastore: %w", addrRef.Type, addrRef.Version, addrRef.Address, addrRef.ChainSelector, err)
			}
		}
		if err := sequences.WriteMetadataToDatastore(ds, sequences.Metadata{
			Contracts: contractMetadata,
		}); err != nil {
			return cldf.ChangesetOutput{Reports: reports}, fmt.Errorf("failed to write metadata to datastore: %w", err)
		}
		return changesets.NewOutputBuilder(e, mcmsRegistry).
			WithReports(reports).
			WithDataStore(ds).
			WithBatchOps(batchOps).
			Build(input.MCMS)
	}
}
