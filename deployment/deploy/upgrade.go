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
	UpgradeChainContracts() *cldf_ops.Sequence[ContractUpgradeConfigPerChainWithAddress, sequences.OnChainOutput, cldf_chain.BlockChains]
}

type ContractUpgradeConfig struct {
	Chains map[uint64]ContractUpgradeConfigPerChain
	MCMS   mcms.Input
}

type ContractUpgradeConfigPerChain struct {
	Version *semver.Version
	// Upgrades maps contract types to their target upgrade versions.
	Upgrades map[cldf.ContractType]*semver.Version
	// UpgradeAuthority is the current authority that owns the deployed programs.
	// On Solana this is typically the timelock signer PDA or the deployer key.
	UpgradeAuthority string
	// ChainSpecific holds chain-family-specific upgrade configuration.
	// Solana adapters expect *SolanaUpgradeExtensions here.
	ChainSpecific any
}

type ContractUpgradeConfigPerChainWithAddress struct {
	ContractUpgradeConfigPerChain
	ChainSelector     uint64
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
	for selector, config := range cfg.Chains {
		_, err := chain_selectors.GetSelectorFamily(selector)
		if err != nil {
			return fmt.Errorf("no selector %d found in environment: %w", selector, err)
		}
		if config.Version == nil {
			return fmt.Errorf("no version specified for chain with selector %d", selector)
		}
		if len(config.Upgrades) == 0 {
			return fmt.Errorf("no upgrades specified for chain with selector %d", selector)
		}
	}
	return nil
}

func upgradeContractsApply(u *UpgraderRegistry, d *DeployerRegistry) func(cldf.Environment, ContractUpgradeConfig) (cldf.ChangesetOutput, error) {
	return func(e cldf.Environment, cfg ContractUpgradeConfig) (cldf.ChangesetOutput, error) {
		reports := make([]cldf_ops.Report[any, any], 0)
		ds := datastore.NewMemoryDataStore()
		for selector, upgradeCfg := range cfg.Chains {
			family, err := chain_selectors.GetSelectorFamily(selector)
			if err != nil {
				return cldf.ChangesetOutput{}, err
			}
			upgrader, exists := u.GetUpgrader(family, upgradeCfg.Version)
			if !exists {
				return cldf.ChangesetOutput{}, fmt.Errorf("no upgrader registered for chain family %s and version %s", family, upgradeCfg.Version.String())
			}
			// If the deployer also implements ArtifactPreparer, run it for upgrades too.
			// Upgrade builds typically need key replacement to match existing program IDs.
			deployer, deployerExists := d.GetDeployer(family, upgradeCfg.Version)
			if deployerExists {
				if preparer, ok := deployer.(ArtifactPreparer); ok {
					deployCfg := ContractDeploymentConfigPerChain{
						Version:       upgradeCfg.Version,
						ChainSpecific: upgradeCfg.ChainSpecific,
					}
					if err := preparer.PrepareArtifacts(e, selector, deployCfg); err != nil {
						return cldf.ChangesetOutput{}, fmt.Errorf("failed to prepare artifacts for upgrade on chain %d: %w", selector, err)
					}
				}
			}
			existingAddrs := d.ExistingAddressesForChain(e, selector)
			seqCfg := ContractUpgradeConfigPerChainWithAddress{
				ContractUpgradeConfigPerChain: upgradeCfg,
				ExistingAddresses:             existingAddrs,
				ChainSelector:                 selector,
			}
			upgradeReport, err := cldf_ops.ExecuteSequence(e.OperationsBundle, upgrader.UpgradeChainContracts(), e.BlockChains, seqCfg)
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("failed to upgrade contracts on chain with selector %d: %w", selector, err)
			}
			for _, r := range upgradeReport.Output.Addresses {
				if err := ds.Addresses().Add(r); err != nil {
					return cldf.ChangesetOutput{}, fmt.Errorf("failed to add %s %s with address %s on chain %d to datastore: %w", r.Type, r.Version, r.Address, r.ChainSelector, err)
				}
			}
			reports = append(reports, upgradeReport.ExecutionReports...)
		}

		return changesets.NewOutputBuilder(e, nil).
			WithReports(reports).
			WithDataStore(ds).
			Build(cfg.MCMS)
	}
}
