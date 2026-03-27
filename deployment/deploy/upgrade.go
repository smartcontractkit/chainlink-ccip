package deploy

import (
	"fmt"
	"sync"

	"github.com/Masterminds/semver/v3"
	chain_selectors "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"

	"github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
)

// Upgrader is a separate interface from Deployer for chains that support
// in-place program/contract upgrades. Not all chain families support this
// (e.g. EVM uses proxy patterns instead), so this is kept separate to avoid
// forcing all Deployer implementations to stub upgrade methods.
type Upgrader interface {
	UpgradeChainContracts() *cldf_ops.Sequence[ContractUpgradeConfigWithAddress, sequences.OnChainOutput, cldf_chain.BlockChains]
}

// ContractUpgradeConfig targets a single chain. Upgrades are chain-specific
// operations — unlike deploys, they won't be batched across chains.
type ContractUpgradeConfig struct {
	ChainSelector uint64
	// Version selects which upgrader adapter to use from the registry.
	Version *semver.Version
	// Contracts lists the contract types to upgrade. The upgrade authority
	// and existing addresses are read from on-chain state by the adapter.
	Contracts []cldf.ContractType
	// ChainSpecific holds chain-family-specific upgrade configuration.
	// Solana adapters expect *SolanaBuildConfig here for artifact preparation.
	ChainSpecific any
	MCMS          mcms.Input
}

// ContractUpgradeConfigWithAddress is the input passed to the upgrade sequence,
// enriched with existing on-chain addresses by the changeset.
type ContractUpgradeConfigWithAddress struct {
	ContractUpgradeConfig
	ExistingAddresses []datastore.AddressRef
}

// UpgraderRegistry is a registry for chain-family-specific upgrade adapters.
type UpgraderRegistry struct {
	mu        sync.Mutex
	upgraders map[string]Upgrader
}

var (
	singletonUpgraderRegistry *UpgraderRegistry
	upgraderRegistryOnce      sync.Once
)

func newUpgraderRegistry() *UpgraderRegistry {
	return &UpgraderRegistry{
		upgraders: make(map[string]Upgrader),
	}
}

func GetUpgraderRegistry() *UpgraderRegistry {
	upgraderRegistryOnce.Do(func() {
		singletonUpgraderRegistry = newUpgraderRegistry()
	})
	return singletonUpgraderRegistry
}

func (r *UpgraderRegistry) RegisterUpgrader(chainFamily string, version *semver.Version, upgrader Upgrader) {
	r.mu.Lock()
	defer r.mu.Unlock()
	id := utils.NewRegistererID(chainFamily, version)
	if _, exists := r.upgraders[id]; !exists {
		r.upgraders[id] = upgrader
	}
}

func (r *UpgraderRegistry) GetUpgrader(chainFamily string, version *semver.Version) (Upgrader, bool) {
	r.mu.Lock()
	defer r.mu.Unlock()
	id := utils.NewRegistererID(chainFamily, version)
	upgrader, ok := r.upgraders[id]
	return upgrader, ok
}

func UpgradeContracts(upgraderReg *UpgraderRegistry, deployerReg *DeployerRegistry) cldf.ChangeSetV2[ContractUpgradeConfig] {
	return cldf.CreateChangeSet(upgradeContractsApply(upgraderReg, deployerReg), upgradeContractsVerify)
}

func upgradeContractsVerify(_ cldf.Environment, cfg ContractUpgradeConfig) error {
	_, err := chain_selectors.GetSelectorFamily(cfg.ChainSelector)
	if err != nil {
		return fmt.Errorf("no selector %d found: %w", cfg.ChainSelector, err)
	}
	if cfg.Version == nil {
		return fmt.Errorf("no version specified for chain with selector %d", cfg.ChainSelector)
	}
	if len(cfg.Contracts) == 0 {
		return fmt.Errorf("no contracts specified for upgrade on chain with selector %d", cfg.ChainSelector)
	}
	return nil
}

func upgradeContractsApply(u *UpgraderRegistry, d *DeployerRegistry) func(cldf.Environment, ContractUpgradeConfig) (cldf.ChangesetOutput, error) {
	return func(e cldf.Environment, cfg ContractUpgradeConfig) (cldf.ChangesetOutput, error) {
		family, err := chain_selectors.GetSelectorFamily(cfg.ChainSelector)
		if err != nil {
			return cldf.ChangesetOutput{}, err
		}
		upgrader, exists := u.GetUpgrader(family, cfg.Version)
		if !exists {
			return cldf.ChangesetOutput{}, fmt.Errorf("no upgrader registered for chain family %s and version %s", family, cfg.Version.String())
		}
		// If the deployer also implements ArtifactPreparer, run it for upgrades too.
		// Upgrade builds typically need key replacement to match existing program IDs.
		deployer, deployerExists := d.GetDeployer(family, cfg.Version)
		if deployerExists {
			if preparer, ok := deployer.(ArtifactPreparer); ok {
				deployCfg := ContractDeploymentConfigPerChain{
					Version:       cfg.Version,
					ChainSpecific: cfg.ChainSpecific,
				}
				if err := preparer.PrepareArtifacts(e, cfg.ChainSelector, deployCfg); err != nil {
					return cldf.ChangesetOutput{}, fmt.Errorf("failed to prepare artifacts for upgrade on chain %d: %w", cfg.ChainSelector, err)
				}
			}
		}
		existingAddrs := d.ExistingAddressesForChain(e, cfg.ChainSelector)
		seqCfg := ContractUpgradeConfigWithAddress{
			ContractUpgradeConfig: cfg,
			ExistingAddresses:     existingAddrs,
		}
		upgradeReport, err := cldf_ops.ExecuteSequence(e.OperationsBundle, upgrader.UpgradeChainContracts(), e.BlockChains, seqCfg)
		if err != nil {
			return cldf.ChangesetOutput{}, fmt.Errorf("failed to upgrade contracts on chain with selector %d: %w", cfg.ChainSelector, err)
		}
		ds := datastore.NewMemoryDataStore()
		for _, r := range upgradeReport.Output.Addresses {
			if err := ds.Addresses().Add(r); err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("failed to add %s %s with address %s on chain %d to datastore: %w", r.Type, r.Version, r.Address, r.ChainSelector, err)
			}
		}

		return changesets.NewOutputBuilder(e, nil).
			WithReports(upgradeReport.ExecutionReports).
			WithDataStore(ds).
			Build(cfg.MCMS)
	}
}
