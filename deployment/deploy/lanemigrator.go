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
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
)

var (
	singletonLaneMigratorRegistry *LaneMigratorRegistry
	laneMigratorOnce              sync.Once
)

type LaneMigratorConfig struct {
	Input map[uint64]LaneMigratorConfigPerChain
	MCMS  mcms.Input
}

type LaneMigratorConfigPerChain struct {
	RemoteChains  []uint64
	RouterVersion *semver.Version
	RampVersion   *semver.Version
}

type RouterUpdaterConfig struct {
	ChainSelector        uint64
	RemoteChainSelectors []uint64
	OnRamp               datastore.AddressRef
	OffRamp              datastore.AddressRef
	ExistingAddresses    []datastore.AddressRef
}

type RampUpdateInRouter interface {
	UpdateRouter() *cldf_ops.Sequence[RouterUpdaterConfig, sequences.OnChainOutput, chain.BlockChains]
}

type RampUpdaterConfig struct {
	ChainSelector        uint64
	RemoteChainSelectors []uint64
	RouterAddr           datastore.AddressRef
	ExistingAddresses    []datastore.AddressRef
}

type RouterUpdateInRamp interface {
	UpdateVersionWithRouter() *cldf_ops.Sequence[RampUpdaterConfig, sequences.OnChainOutput, chain.BlockChains]
}

type LaneMigratorRegistry struct {
	RouterUpdater map[string]RampUpdateInRouter
	RampUpdater   map[string]RouterUpdateInRamp
}

func newLaneMigratorRegistry() *LaneMigratorRegistry {
	return &LaneMigratorRegistry{
		RouterUpdater: make(map[string]RampUpdateInRouter),
		RampUpdater:   make(map[string]RouterUpdateInRamp),
	}
}

func GetLaneMigratorRegistry() *LaneMigratorRegistry {
	laneMigratorOnce.Do(func() {
		singletonLaneMigratorRegistry = newLaneMigratorRegistry()
	})
	return singletonLaneMigratorRegistry
}

func (r *LaneMigratorRegistry) RegisterRouterUpdater(chainfamily string, version *semver.Version, updater RampUpdateInRouter) {
	id := utils.NewRegistererID(chainfamily, version)
	if _, exists := r.RouterUpdater[id]; !exists {
		r.RouterUpdater[id] = updater
	}
}

func (r *LaneMigratorRegistry) RegisterRampUpdater(chainfamily string, version *semver.Version, updater RouterUpdateInRamp) {
	id := utils.NewRegistererID(chainfamily, version)
	if _, exists := r.RampUpdater[id]; !exists {
		r.RampUpdater[id] = updater
	}
}

func (r *LaneMigratorRegistry) GetRouterUpdater(chainsel uint64, version *semver.Version) (RampUpdateInRouter, error) {
	id := utils.NewIDFromSelector(chainsel, version)
	updater, exists := r.RouterUpdater[id]
	if !exists {
		return nil, utils.ErrNoAdapterRegistered(id, version)
	}
	return updater, nil
}

func (r *LaneMigratorRegistry) GetRampUpdater(chainsel uint64, version *semver.Version) (RouterUpdateInRamp, error) {
	id := utils.NewIDFromSelector(chainsel, version)
	updater, exists := r.RampUpdater[id]
	if !exists {
		return nil, utils.ErrNoAdapterRegistered(id, version)
	}
	return updater, nil
}

func LaneMigrateToNewVersionChangeset(migratorReg *LaneMigratorRegistry, mcmsRegistry *changesets.MCMSReaderRegistry) cldf.ChangeSetV2[LaneMigratorConfig] {
	return cldf.CreateChangeSet(lanemigrateApply(migratorReg, mcmsRegistry), lanemigrateVerify(migratorReg))
}

func lanemigrateApply(migratorReg *LaneMigratorRegistry, mcmsRegistry *changesets.MCMSReaderRegistry) func(cldf.Environment, LaneMigratorConfig) (cldf.ChangesetOutput, error) {
	return func(e cldf.Environment, input LaneMigratorConfig) (cldf.ChangesetOutput, error) {
		batchOps := make([]mcms_types.BatchOperation, 0)
		reports := make([]cldf_ops.Report[any, any], 0)
		for chainSel, perChainConfig := range input.Input {
			existingAddresses := e.DataStore.Addresses().Filter(datastore.AddressRefByChainSelector(chainSel))
			routerUpdater, err := migratorReg.GetRouterUpdater(chainSel, perChainConfig.RouterVersion)
			if err != nil {
				return cldf.ChangesetOutput{}, err
			}
			rampUpdater, err := migratorReg.GetRampUpdater(chainSel, perChainConfig.RampVersion)
			if err != nil {
				return cldf.ChangesetOutput{}, err
			}
			// get the ramp address refs
			onRampRef, err := datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{
				ChainSelector: chainSel,
				Type:          "OnRamp", // does not work with 1.5 ramps
				Version:       perChainConfig.RampVersion,
			}, chainSel, datastore_utils.FullRef)
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("error finding onRamp address ref for chain selector %d: %w", chainSel, err)
			}
			offRampRef, err := datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{
				ChainSelector: chainSel,
				Type:          "OffRamp", // does not work with 1.5 ramps
				Version:       perChainConfig.RampVersion,
			}, chainSel, datastore_utils.FullRef)
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("error finding offRamp address ref for chain selector %d: %w", chainSel, err)
			}
			routerRef, err := datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{
				ChainSelector: chainSel,
				Type:          "Router",
				Version:       perChainConfig.RouterVersion,
			}, chainSel, datastore_utils.FullRef)
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("error finding router address ref for chain selector %d: %w", chainSel, err)
			}
			routerUpdateConfig := RouterUpdaterConfig{
				ChainSelector:        chainSel,
				RemoteChainSelectors: perChainConfig.RemoteChains,
				OnRamp:               onRampRef,
				OffRamp:              offRampRef,
				ExistingAddresses:    existingAddresses,
			}
			report, err := cldf_ops.ExecuteSequence(e.OperationsBundle, routerUpdater.UpdateRouter(), e.BlockChains, routerUpdateConfig)
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("error executing router update sequence for chain selector %d: %w", chainSel, err)
			}
			batchOps = append(batchOps, report.Output.BatchOps...)
			reports = append(reports, report.ExecutionReports...)
			rampUpdateConfig := RampUpdaterConfig{
				ChainSelector:        chainSel,
				RemoteChainSelectors: perChainConfig.RemoteChains,
				RouterAddr:           routerRef,
				ExistingAddresses:    existingAddresses,
			}
			updateRampRep, err := cldf_ops.ExecuteSequence(e.OperationsBundle, rampUpdater.UpdateVersionWithRouter(), e.BlockChains, rampUpdateConfig)
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("error executing ramp update sequence for chain selector %d: %w", chainSel, err)
			}
			batchOps = append(batchOps, updateRampRep.Output.BatchOps...)
			reports = append(reports, updateRampRep.ExecutionReports...)
		}
		return changesets.NewOutputBuilder(e, mcmsRegistry).
			WithReports(reports).
			WithBatchOps(batchOps).
			Build(input.MCMS)
	}
}

func lanemigrateVerify(migratorReg *LaneMigratorRegistry) func(cldf.Environment, LaneMigratorConfig) error {
	return func(e cldf.Environment, input LaneMigratorConfig) error {
		for chainSel, perChainConfig := range input.Input {
			_, err := migratorReg.GetRouterUpdater(chainSel, perChainConfig.RouterVersion)
			if err != nil {
				return fmt.Errorf("error verifying existence of router updater for chain selector %d: %w", chainSel, err)
			}
			_, err = migratorReg.GetRampUpdater(chainSel, perChainConfig.RampVersion)
			if err != nil {
				return fmt.Errorf("error verifying existence of ramp updater for chain selector %d: %w", chainSel, err)
			}
			if !e.BlockChains.Exists(chainSel) {
				return fmt.Errorf("error verifying existence of blockchain with selector %d in environment: blockchain not found", chainSel)
			}
			for _, remoteChainSel := range perChainConfig.RemoteChains {
				if !e.BlockChains.Exists(remoteChainSel) {
					return fmt.Errorf("error verifying existence of remote blockchain with selector %d in environment: blockchain not found", remoteChainSel)
				}
			}
			// verify that the existing addresses for the chain selector are present in the environment datastore
			existingAddresses := e.DataStore.Addresses().Filter(datastore.AddressRefByChainSelector(chainSel))
			if len(existingAddresses) == 0 {
				return fmt.Errorf("error verifying existence of existing addresses for chain selector %d in environment datastore: no addresses found", chainSel)
			}
		}
		return nil
	}
}
