package tokens

import (
	"errors"
	"fmt"
	"maps"
	"slices"

	"github.com/Masterminds/semver/v3"
	chain_selectors "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"

	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
)

// DeployMultiTokenPoolChainConfig configures one multi-token pool stack on a chain.
type DeployMultiTokenPoolChainConfig struct {
	PoolType      string          `json:"poolType" yaml:"poolType"`
	PoolVersion   *semver.Version `json:"poolVersion" yaml:"poolVersion"`
	PoolQualifier string          `json:"poolQualifier" yaml:"poolQualifier"`
	// Admin is the initial PoolConfigManager default admin. If empty, the selected MCMS
	// timelock is used when available, with a chain-specific deployer fallback otherwise.
	Admin string `json:"admin,omitempty" yaml:"admin,omitempty"`
}

// DeployMultiTokenPoolsConfig configures the DeployMultiTokenPools changeset.
type DeployMultiTokenPoolsConfig struct {
	Chains map[uint64]DeployMultiTokenPoolChainConfig `json:"chains" yaml:"chains"`
	MCMS   mcms.Input                                 `json:"mcms" yaml:"mcms,omitempty"`
}

// DeployMultiTokenPools deploys the three policy stores, PoolConfigManager, and multi-token
// pool on each requested chain. It also authorizes each PoolConfigManager on its stores.
// Token registration, policy configuration, role setup, and ownership transfers are handled
// by separate changesets. Existing refs are reused as-is; constructor-only wiring and the initial
// admin are not reconciled on reruns.
func DeployMultiTokenPools(registry *MultiTokenPoolAdapterRegistry, mcmsRegistry *changesets.MCMSReaderRegistry) cldf.ChangeSetV2[DeployMultiTokenPoolsConfig] {
	return cldf.CreateChangeSet(
		makeApplyDeployMultiTokenPools(registry, mcmsRegistry),
		makeVerifyDeployMultiTokenPools(registry, mcmsRegistry),
	)
}

func makeVerifyDeployMultiTokenPools(registry *MultiTokenPoolAdapterRegistry, mcmsRegistry *changesets.MCMSReaderRegistry) func(cldf.Environment, DeployMultiTokenPoolsConfig) error {
	return func(e cldf.Environment, cfg DeployMultiTokenPoolsConfig) error {
		if len(cfg.Chains) == 0 {
			return errors.New("at least one chain is required")
		}
		mcmsInput := cfg.MCMS
		if err := mcmsInput.PopulateDefaults(); err != nil {
			return fmt.Errorf("failed to populate MCMS defaults: %w", err)
		}
		if err := mcmsInput.Validate(); err != nil {
			return fmt.Errorf("invalid MCMS input: %w", err)
		}
		for selector, chainCfg := range cfg.Chains {
			if chainCfg.PoolType == "" {
				return fmt.Errorf("poolType is required for chain selector %d", selector)
			}
			if chainCfg.PoolVersion == nil {
				return fmt.Errorf("poolVersion is required for chain selector %d", selector)
			}
			if chainCfg.PoolQualifier == "" {
				return fmt.Errorf("poolQualifier is required for chain selector %d", selector)
			}
			adapter, family, err := resolveMultiTokenPoolAdapter(registry, selector, chainCfg.PoolVersion)
			if err != nil {
				return err
			}
			admin, err := resolveMultiTokenPoolAdmin(e, mcmsRegistry, selector, family, chainCfg.Admin, mcmsInput)
			if err != nil {
				return err
			}
			if err := adapter.ValidateDeployMultiTokenPool(e, DeployMultiTokenPoolInput{
				ChainSelector: selector,
				PoolType:      chainCfg.PoolType,
				PoolVersion:   chainCfg.PoolVersion,
				PoolQualifier: chainCfg.PoolQualifier,
				Admin:         admin,
			}); err != nil {
				return fmt.Errorf("invalid multi-token pool config for chain selector %d: %w", selector, err)
			}
		}
		return nil
	}
}

func makeApplyDeployMultiTokenPools(registry *MultiTokenPoolAdapterRegistry, mcmsRegistry *changesets.MCMSReaderRegistry) func(cldf.Environment, DeployMultiTokenPoolsConfig) (cldf.ChangesetOutput, error) {
	return func(e cldf.Environment, cfg DeployMultiTokenPoolsConfig) (cldf.ChangesetOutput, error) {
		ds := datastore.NewMemoryDataStore()
		batchOps := make([]mcms_types.BatchOperation, 0)
		reports := make([]cldf_ops.Report[any, any], 0)

		for _, selector := range slices.Sorted(maps.Keys(cfg.Chains)) {
			chainCfg := cfg.Chains[selector]
			adapter, family, err := resolveMultiTokenPoolAdapter(registry, selector, chainCfg.PoolVersion)
			if err != nil {
				return cldf.ChangesetOutput{}, err
			}

			admin, err := resolveMultiTokenPoolAdmin(e, mcmsRegistry, selector, family, chainCfg.Admin, cfg.MCMS)
			if err != nil {
				return cldf.ChangesetOutput{}, err
			}

			if admin == "" {
				e.Logger.Warnf("no timelock found for chain selector %d; PoolConfigManager admin will fall back to the deployer", selector)
			}

			report, err := cldf_ops.ExecuteSequence(e.OperationsBundle, adapter.DeployMultiTokenPool(), e.BlockChains, DeployMultiTokenPoolInput{
				ChainSelector:     selector,
				PoolType:          chainCfg.PoolType,
				PoolVersion:       chainCfg.PoolVersion,
				PoolQualifier:     chainCfg.PoolQualifier,
				Admin:             admin,
				ExistingDataStore: e.DataStore,
			})
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("failed to deploy multi-token pool on chain selector %d: %w", selector, err)
			}
			for _, ref := range report.Output.Addresses {
				// Skip refs already in the datastore; re-emitting them would collide on merge.
				existing, err := e.DataStore.Addresses().Get(ref.Key())
				switch {
				case err == nil:
					if existing.Address != ref.Address {
						return cldf.ChangesetOutput{}, fmt.Errorf("%s on chain selector %d conflicts with the existing datastore address %s", datastore_utils.SprintRef(ref), selector, existing.Address)
					}
					continue
				case !errors.Is(err, datastore.ErrAddressRefNotFound):
					return cldf.ChangesetOutput{}, fmt.Errorf("failed to look up %s on chain selector %d in the existing datastore: %w", datastore_utils.SprintRef(ref), selector, err)
				}
				if err := ds.Addresses().Add(ref); err != nil {
					return cldf.ChangesetOutput{}, fmt.Errorf("failed to add %s on chain selector %d to the output datastore: %w", datastore_utils.SprintRef(ref), selector, err)
				}
			}
			batchOps = append(batchOps, report.Output.BatchOps...)
			reports = append(reports, report.ExecutionReports...)
		}

		return changesets.NewOutputBuilder(e, mcmsRegistry).
			WithReports(reports).
			WithSingleBatchOpPerChain(batchOps).
			WithDataStore(ds).
			Build(cfg.MCMS)
	}
}

func resolveMultiTokenPoolAdmin(e cldf.Environment, mcmsRegistry *changesets.MCMSReaderRegistry, selector uint64, family, admin string, mcmsInput mcms.Input) (string, error) {
	if admin != "" {
		return admin, nil
	}
	mcmsReader, ok := mcmsRegistry.GetMCMSReader(family)
	if !ok {
		return "", fmt.Errorf("no MCMS reader registered for chain family %q", family)
	}
	timelockRef, err := mcmsReader.GetTimelockRef(e, selector, mcmsInput)
	if err != nil {
		return "", fmt.Errorf("failed to resolve timelock for chain selector %d: %w", selector, err)
	}
	if datastore_utils.IsAddressRefEmpty(timelockRef) {
		if mcmsInput.Qualifier != "" {
			return "", fmt.Errorf("MCMS qualifier %q was provided but no timelock was found for chain selector %d; refusing to fall back to the deployer as PoolConfigManager admin", mcmsInput.Qualifier, selector)
		}
		return "", nil
	}
	return timelockRef.Address, nil
}

func resolveMultiTokenPoolAdapter(registry *MultiTokenPoolAdapterRegistry, selector uint64, version *semver.Version) (MultiTokenPoolAdapter, string, error) {
	family, err := chain_selectors.GetSelectorFamily(selector)
	if err != nil {
		return nil, "", fmt.Errorf("invalid chain selector %d: %w", selector, err)
	}
	adapter, ok := registry.GetMultiTokenPoolAdapter(family, version)
	if !ok {
		return nil, "", fmt.Errorf("no multi-token pool adapter registered for chain family %q and version %s", family, version)
	}
	return adapter, family, nil
}
