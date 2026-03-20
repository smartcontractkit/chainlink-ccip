package deploy

import (
	"fmt"
	"math/big"
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
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
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
	FeeQuoterVersion *semver.Version
	FeeQuoterConfig  *AdditionalFeeQuoterConfig
	RampsVersion     *semver.Version
	// RemoteChainSelectors is used to determine which remote chains to pull config for when populating config for the FeeQuoter
	// if RemoteChainSelectors is empty, it will pull all remote chain configs using 1.5.0 and 1.6.0 config importer
	RemoteChainSelectors []uint64
}

type AdditionalFeeQuoterConfig struct {
	GasPricesPerRemoteChain map[uint64]string // uses string values (parsed as base-10 big.Int).
}

type FeeQuoterUpdateInput struct {
	ChainSelector     uint64
	ExistingAddresses []datastore.AddressRef
	// PreviousVersions lists the supported config-importer / lane versions that
	// should be consulted when deriving the FeeQuoter configuration for this chain.
	// It does NOT refer to previous FeeQuoter contract deployment versions.
	PreviousVersions     []*semver.Version
	RemoteChainSelectors []uint64
	AdditionalConfig     *AdditionalFeeQuoterConfig
	ContractMeta         []datastore.ContractMetadata
	// TimelockAddress is the address of the CCIP timelock contract to be added as a price updater on the fee quoter.
	TimelockAddress string
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
	FeeQuoterUpdater            map[string]FeeQuoterUpdater[any]
	RampUpdater                 map[string]RampUpdater
	ConfigImporter              map[string]ConfigImporter
	ImportconfigVersionResolver map[string]LaneVersionResolver
	mu                          sync.Mutex
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

func (r *FQAndRampUpdaterRegistry) RegisterConfigImporterVersionResolver(family string, resolver LaneVersionResolver) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, exists := r.ImportconfigVersionResolver[family]; !exists {
		r.ImportconfigVersionResolver[family] = resolver
	}
}

func (r *FQAndRampUpdaterRegistry) GetConfigImporter(chainsel uint64, version *semver.Version) (ConfigImporter, bool) {
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

func (r *FQAndRampUpdaterRegistry) GetConfigImporterVersionResolver(chainsel uint64) (LaneVersionResolver, bool) {
	family, err := chain_selectors.GetSelectorFamily(chainsel)
	if err != nil {
		return nil, false
	}
	r.mu.Lock()
	defer r.mu.Unlock()
	resolver, ok := r.ImportconfigVersionResolver[family]
	return resolver, ok
}

func newFQUpdaterRegistry() *FQAndRampUpdaterRegistry {
	return &FQAndRampUpdaterRegistry{
		FeeQuoterUpdater:            make(map[string]FeeQuoterUpdater[any]),
		RampUpdater:                 make(map[string]RampUpdater),
		ConfigImporter:              make(map[string]ConfigImporter),
		ImportconfigVersionResolver: make(map[string]LaneVersionResolver),
	}
}

func GetFQAndRampUpdaterRegistry() *FQAndRampUpdaterRegistry {
	fqupdaterOnce.Do(func() {
		singletonFQAndRampUpdaterRegistry = newFQUpdaterRegistry()
	})
	return singletonFQAndRampUpdaterRegistry
}

// UpdateFeeQuoterChangeset creates a changeset that updates FeeQuoter contracts on specified chains.
// This can support either upgrading to a 2.0 or higher version of the FeeQuoter, which allows re-configuration of the existing contract, or deploying a new FeeQuoter and updating Ramps to point to it.
// It first optionally populates configuration values, then creates FeeQuoterUpdateInput,
// deploys or updates the FeeQuoter contract, and finally updates the Ramps contracts to use the new FeeQuoter address.
// This also supports downgrading the FQ contract to a prior version (i.e. a rollback) where only the ramps
// are updated and the existing FQ is not touched. This would be triggered when specifying a FQ version < 2.0.0, which do not support re-configuration, and an existing FQ address is found in the datastore for the chain.
func UpdateFeeQuoterChangeset() cldf.ChangeSetV2[UpdateFeeQuoterInput] {
	return cldf.CreateChangeSet(updateFeeQuoterApply(), updateFeeQuoterVerify())
}

func updateFeeQuoterVerify() func(cldf.Environment, UpdateFeeQuoterInput) error {
	return func(e cldf.Environment, input UpdateFeeQuoterInput) error {
		for chainSel, perChainInput := range input.Chains {
			if !e.BlockChains.Exists(chainSel) {
				return fmt.Errorf("chain with selector %d not found in environment", chainSel)
			}
			if perChainInput.FeeQuoterVersion == nil {
				return fmt.Errorf("fee quoter version is required for chain selector %d", chainSel)
			}
			if perChainInput.FeeQuoterConfig != nil {
				for remoteChainSel := range perChainInput.FeeQuoterConfig.GasPricesPerRemoteChain {
					_, ok := new(big.Int).SetString(perChainInput.FeeQuoterConfig.GasPricesPerRemoteChain[remoteChainSel], 10)
					if !ok {
						return fmt.Errorf("invalid gas price %s for remote chain selector %d in fee quoter config for chain selector %d", perChainInput.FeeQuoterConfig.GasPricesPerRemoteChain[remoteChainSel], remoteChainSel, chainSel)
					}
				}
			}
			_, err := datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{
				ChainSelector: chainSel,
				Type:          datastore.ContractType(utils.FeeQuoter),
				Version:       perChainInput.FeeQuoterVersion,
			}, chainSel, datastore_utils.FullRef)
			if err != nil {
				// errors are alright if we don't expect to find the ref
				// but we only support deploying/updating fee quoters with versions >= 2.0.0
				supportedVersion := semver.MustParse("2.0.0")
				if perChainInput.FeeQuoterVersion.LessThan(supportedVersion) {
					return fmt.Errorf("fee quoter address not found for chain selector %d and version %s: %w", chainSel, perChainInput.FeeQuoterVersion.String(), err)
				}
			}
		}
		return nil
	}
}

func updateFeeQuoterApply() func(cldf.Environment, UpdateFeeQuoterInput) (cldf.ChangesetOutput, error) {
	return func(e cldf.Environment, input UpdateFeeQuoterInput) (cldf.ChangesetOutput, error) {
		batchOps := make([]mcms_types.BatchOperation, 0)
		reports := make([]cldf_ops.Report[any, any], 0)
		addressRefs := make([]datastore.AddressRef, 0)
		contractMetadata := make([]datastore.ContractMetadata, 0)
		fquRegistry := GetFQAndRampUpdaterRegistry()
		mcmsRegistry := changesets.GetRegistry()
		for chainSel, perChainInput := range input.Chains {
			var feeQuoterAddrRef datastore.AddressRef
			feeQuoterAddrRefs := e.DataStore.Addresses().Filter(
				datastore.AddressRefByChainSelector(chainSel),
				datastore.AddressRefByType(datastore.ContractType(utils.FeeQuoter)),
				datastore.AddressRefByVersion(perChainInput.FeeQuoterVersion),
			)
			if len(feeQuoterAddrRefs) > 0 {
				feeQuoterAddrRef = feeQuoterAddrRefs[0]
				e.Logger.Infof("Found existing FeeQuoter address %s for chain selector %d and version %s",
					feeQuoterAddrRef.Address, chainSel, perChainInput.FeeQuoterVersion.String())
			}
			isNewFeeQuoterDeployment := len(feeQuoterAddrRefs) == 0
			if perChainInput.FeeQuoterVersion.GreaterThanEqual(semver.MustParse("2.0.0")) {
				e.Logger.Infof("No existing FeeQuoter address found for chain selector %d and version %s, proceeding with deployment and upgrade", chainSel, perChainInput.FeeQuoterVersion.String())
				fquUpdater, ok := fquRegistry.GetFeeQuoterUpdater(chainSel, perChainInput.FeeQuoterVersion)
				if !ok {
					return cldf.ChangesetOutput{}, utils.ErrNoAdapterForSelectorRegistered("FeeQuoterUpdater", chainSel, perChainInput.FeeQuoterVersion)
				}

				versionResolver, ok := fquRegistry.GetConfigImporterVersionResolver(chainSel)
				if !ok {
					return cldf.ChangesetOutput{}, utils.ErrNoAdapterForSelectorRegistered("ConfigImporterVersionResolver", chainSel, nil)
				}
				_, configImporterVersions, err := versionResolver.DeriveLaneVersionsForChain(e, chainSel)
				if err != nil {
					return cldf.ChangesetOutput{}, fmt.Errorf("failed to resolve config importer version for chain %d: %w", chainSel, err)
				}
				contractMeta := make([]datastore.ContractMetadata, 0)
				for _, version := range configImporterVersions {
					configImporter, ok := fquRegistry.GetConfigImporter(chainSel, version)
					if !ok {
						return cldf.ChangesetOutput{}, utils.ErrNoAdapterForSelectorRegistered("ConfigImporter", chainSel, version)
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
				timelockAddr := ""
				family, err := chain_selectors.GetSelectorFamily(chainSel)
				if err != nil {
					return cldf.ChangesetOutput{}, fmt.Errorf("failed to get chain family for selector %d: %w", chainSel, err)
				}
				mcmsReader, ok := mcmsRegistry.GetMCMSReader(family)
				if !ok {
					return cldf.ChangesetOutput{}, fmt.Errorf("no MCMS reader registered for chain family '%s'", family)
				}
				timelockRef, err := mcmsReader.GetTimelockRef(e, chainSel, input.MCMS)
				if err != nil {
					e.Logger.Warnf("Could not resolve timelock ref for chain %d, skipping timelock as price updater: %v", chainSel, err)
				} else {
					timelockAddr = timelockRef.Address
				}

				reportFQInputCreation, err := cldf_ops.ExecuteSequence(e.OperationsBundle, fquUpdater.SequenceFeeQuoterInputCreation(), e.BlockChains, FeeQuoterUpdateInput{
					ChainSelector:        chainSel,
					ExistingAddresses:    e.DataStore.Addresses().Filter(datastore.AddressRefByChainSelector(chainSel)),
					ContractMeta:         contractMeta,
					RemoteChainSelectors: perChainInput.RemoteChainSelectors,
					TimelockAddress:      timelockAddr,
					AdditionalConfig:     perChainInput.FeeQuoterConfig,
					PreviousVersions:     configImporterVersions,
				})
				if err != nil {
					return cldf.ChangesetOutput{}, fmt.Errorf("failed to create FeeQuoterUpdateInput for chain %d: %w", chainSel, err)
				}
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
				feeQuoterAddrRef = reportFQUpdate.Output.Addresses[len(reportFQUpdate.Output.Addresses)-1]

				if isNewFeeQuoterDeployment && timelockAddr != "" {
					fqTransferBatches, fqTransferReports, err := TransferToTimelock(chainSel, &e, input.MCMS, []datastore.AddressRef{feeQuoterAddrRef})
					if err != nil {
						return cldf.ChangesetOutput{}, fmt.Errorf("failed to transfer ownership to timelock for chain %d: %w", chainSel, err)
					}
					batchOps = append(batchOps, fqTransferBatches...)
					reports = append(reports, fqTransferReports...)
				}
			}
			if perChainInput.RampsVersion != nil {
				if feeQuoterAddrRef.Address == "" {
					return cldf.ChangesetOutput{}, fmt.Errorf("fee quoter address ref is required to update ramps for chain %d", chainSel)
				}
				rampUpdater, ok := fquRegistry.GetRampUpdater(chainSel, perChainInput.RampsVersion)
				if !ok {
					return cldf.ChangesetOutput{}, utils.ErrNoAdapterForSelectorRegistered("RampUpdater", chainSel, perChainInput.RampsVersion)
				}
				rampsInput := UpdateRampsInput{
					ChainSelector:    chainSel,
					FeeQuoterAddress: feeQuoterAddrRef,
				}
				resolvedRampsInput, err := rampUpdater.ResolveRampsInput(e, rampsInput)
				if err != nil {
					return cldf.ChangesetOutput{}, fmt.Errorf("failed to resolve ramps input for chain %d: %w", chainSel, err)
				}
				reportRampsUpdate, err := cldf_ops.ExecuteSequence(e.OperationsBundle, rampUpdater.SequenceUpdateRampsWithFeeQuoter(), e.BlockChains, resolvedRampsInput)
				if err != nil {
					return cldf.ChangesetOutput{}, fmt.Errorf("failed to update ramps with FeeQuoter for chain %d: %w", chainSel, err)
				}
				batchOps = append(batchOps, reportRampsUpdate.Output.BatchOps...)
				addressRefs = append(addressRefs, reportRampsUpdate.Output.Addresses...)
				reports = append(reports, reportRampsUpdate.ExecutionReports...)
			}
		}
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

func PopulateMetaDataFromConfigImporter(e cldf.Environment, configImporter ConfigImporter, chainSel uint64) (sequences.OnChainOutput, error) {
	err := configImporter.InitializeAdapter(e, chainSel)
	if err != nil {
		return sequences.OnChainOutput{}, fmt.Errorf("failed to initialize config importer for chain %d: %w", chainSel, err)
	}
	supportedTokensPerRemoteChain, err := configImporter.SupportedTokensPerRemoteChain(e, chainSel)
	if err != nil {
		return sequences.OnChainOutput{}, fmt.Errorf("failed to get supported tokens per remote chain for chain %d: %w", chainSel, err)
	}
	connectedChains, err := configImporter.ConnectedChains(e, chainSel)
	if err != nil {
		return sequences.OnChainOutput{}, fmt.Errorf("failed to get connected chains for chain %d: %w", chainSel, err)
	}
	populateConfigReport, err := cldf_ops.ExecuteSequence(e.OperationsBundle, configImporter.SequenceImportConfig(), e.BlockChains, ImportConfigPerChainInput{
		ChainSelector:        chainSel,
		RemoteChains:         connectedChains,
		TokensPerRemoteChain: supportedTokensPerRemoteChain,
	})
	if err != nil {
		return sequences.OnChainOutput{}, fmt.Errorf("failed to populate config on chain %d: %w", chainSel, err)
	}

	return populateConfigReport.Output, nil
}
