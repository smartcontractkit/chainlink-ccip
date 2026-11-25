package tokens

import (
	"errors"
	"fmt"
	"maps"
	"slices"
	"strings"

	"github.com/Masterminds/semver/v3"
	"github.com/google/uuid"
	chain_selectors "github.com/smartcontractkit/chain-selectors"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
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
	// timelock is used. When no timelock can be resolved the changeset fails, so set Admin
	// explicitly to use any other account, including the deployer.
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
			admin, selectedTimelockAddress, err := resolveMultiTokenPoolAdmin(e, mcmsRegistry, selector, family, chainCfg.Admin, mcmsInput)
			if err != nil {
				return err
			}
			if err := adapter.ValidateDeployMultiTokenPool(e, DeployMultiTokenPoolInput{
				ChainSelector:           selector,
				PoolType:                chainCfg.PoolType,
				PoolVersion:             chainCfg.PoolVersion,
				PoolQualifier:           chainCfg.PoolQualifier,
				Admin:                   admin,
				AdminExplicit:           chainCfg.Admin != "",
				SelectedTimelockAddress: selectedTimelockAddress,
				ExistingDataStore:       e.DataStore,
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
		// fail returns the work done so far with the error. The durable runner discards the output
		// when the changeset fails, so the error text also lists every contract deployed in this
		// run for the operator to record before retrying.
		fail := func(err error) (cldf.ChangesetOutput, error) {
			if deployed := deployedRefsDescription(ds, reports); deployed != "" {
				err = fmt.Errorf("%w; contracts already deployed in this run, add them to the datastore before retrying: %s", err, deployed)
			}
			return cldf.ChangesetOutput{DataStore: ds, Reports: reports}, err
		}

		for _, selector := range slices.Sorted(maps.Keys(cfg.Chains)) {
			chainCfg := cfg.Chains[selector]
			adapter, family, err := resolveMultiTokenPoolAdapter(registry, selector, chainCfg.PoolVersion)
			if err != nil {
				return fail(err)
			}

			admin, selectedTimelockAddress, err := resolveMultiTokenPoolAdmin(e, mcmsRegistry, selector, family, chainCfg.Admin, cfg.MCMS)
			if err != nil {
				return fail(err)
			}

			input := DeployMultiTokenPoolInput{
				ChainSelector:           selector,
				PoolType:                chainCfg.PoolType,
				PoolVersion:             chainCfg.PoolVersion,
				PoolQualifier:           chainCfg.PoolQualifier,
				Admin:                   admin,
				AdminExplicit:           chainCfg.Admin != "",
				SelectedTimelockAddress: selectedTimelockAddress,
				ExistingDataStore:       e.DataStore,
			}
			deployReport, err := cldf_ops.ExecuteSequence(e.OperationsBundle, adapter.DeployMultiTokenPool(), e.BlockChains, input)
			reports = append(reports, deployReport.ExecutionReports...)
			if err != nil {
				return fail(fmt.Errorf("failed to deploy multi-token pool on chain selector %d: %w", selector, err))
			}
			batchOps = append(batchOps, deployReport.Output.BatchOps...)
			for _, ref := range deployReport.Output.Addresses {
				// Skip refs already in the datastore; re-emitting them would collide on merge.
				existing, err := e.DataStore.Addresses().Get(ref.Key())
				switch {
				case err == nil:
					if existing.Address != ref.Address {
						return fail(fmt.Errorf("%s on chain selector %d conflicts with the existing datastore address %s", datastore_utils.SprintRef(ref), selector, existing.Address))
					}
					continue
				case !errors.Is(err, datastore.ErrAddressRefNotFound):
					return fail(fmt.Errorf("failed to look up %s on chain selector %d in the existing datastore: %w", datastore_utils.SprintRef(ref), selector, err))
				}
				if err := ds.Addresses().Add(ref); err != nil {
					return fail(fmt.Errorf("failed to add %s on chain selector %d to the output datastore: %w", datastore_utils.SprintRef(ref), selector, err))
				}
			}

			authorizationDataStore := datastore.NewMemoryDataStore()
			if err := authorizationDataStore.Merge(e.DataStore); err != nil {
				return fail(fmt.Errorf("failed to merge the existing datastore for multi-token pool authorization on chain selector %d: %w", selector, err))
			}
			if err := authorizationDataStore.Merge(ds.Seal()); err != nil {
				return fail(fmt.Errorf("failed to merge new deployment refs for multi-token pool authorization on chain selector %d: %w", selector, err))
			}
			authorizationInput := input
			authorizationInput.ExistingDataStore = authorizationDataStore.Seal()
			authorizationReport, err := cldf_ops.ExecuteSequence(
				e.OperationsBundle,
				adapter.AuthorizeMultiTokenPool(),
				e.BlockChains,
				authorizationInput,
				cldf_ops.WithSequenceIdempotencyKey[DeployMultiTokenPoolInput, cldf_chain.BlockChains](uuid.NewString()),
			)
			reports = append(reports, authorizationReport.ExecutionReports...)
			if err != nil {
				return fail(fmt.Errorf("failed to authorize multi-token pool on chain selector %d: %w", selector, err))
			}
			batchOps = append(batchOps, authorizationReport.Output.BatchOps...)
		}

		out, err := changesets.NewOutputBuilder(e, mcmsRegistry).
			WithReports(reports).
			WithSingleBatchOpPerChain(batchOps).
			WithDataStore(ds).
			Build(cfg.MCMS)
		if err != nil {
			return fail(err)
		}
		return out, nil
	}
}

// resolveMultiTokenPoolAdmin returns the PoolConfigManager admin and the MCMS timelock selected for
// the chain, which is empty when none resolves. An empty admin resolves to the timelock; there is
// no deployer default, so deployer administration requires setting the admin explicitly.
func resolveMultiTokenPoolAdmin(e cldf.Environment, mcmsRegistry *changesets.MCMSReaderRegistry, selector uint64, family, admin string, mcmsInput mcms.Input) (string, string, error) {
	noTimelock := func() (string, string, error) {
		if admin != "" {
			return admin, "", nil
		}
		return "", "", fmt.Errorf("no MCMS timelock found for chain selector %d (MCMS qualifier %q); refusing to fall back to the deployer as PoolConfigManager admin: set admin explicitly to choose one", selector, mcmsInput.Qualifier)
	}
	mcmsReader, ok := mcmsRegistry.GetMCMSReader(family)
	if !ok {
		return noTimelock()
	}
	timelockRef, err := mcmsReader.GetTimelockRef(e, selector, mcmsInput)
	if err != nil {
		return "", "", fmt.Errorf("failed to resolve timelock for chain selector %d: %w", selector, err)
	}
	if datastore_utils.IsAddressRefEmpty(timelockRef) {
		return noTimelock()
	}
	if timelockRef.Address == "" {
		return "", "", fmt.Errorf("resolved timelock %s for chain selector %d has an empty address", datastore_utils.SprintRef(timelockRef), selector)
	}
	if admin == "" {
		admin = timelockRef.Address
	}
	return admin, timelockRef.Address, nil
}

// deployedRefsDescription lists the contracts deployed so far, taken from the output datastore
// and from the execution reports of deployments that did not reach it.
func deployedRefsDescription(ds *datastore.MemoryDataStore, reports []cldf_ops.Report[any, any]) string {
	seen := make(map[string]struct{})
	lines := make([]string, 0)
	add := func(ref datastore.AddressRef) {
		line := fmt.Sprintf("%s at %s", datastore_utils.SprintRef(ref), ref.Address)
		if _, ok := seen[line]; ok {
			return
		}
		seen[line] = struct{}{}
		lines = append(lines, line)
	}
	if refs, err := ds.Addresses().Fetch(); err == nil {
		for _, ref := range refs {
			add(ref)
		}
	}
	for _, report := range reports {
		switch ref := report.Output.(type) {
		case datastore.AddressRef:
			add(ref)
		case *datastore.AddressRef:
			if ref != nil {
				add(*ref)
			}
		}
	}
	return strings.Join(lines, "; ")
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
