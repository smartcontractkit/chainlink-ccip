package changesets

import (
	"fmt"
	"slices"
	"sort"
	"strconv"

	"github.com/Masterminds/semver/v3"
	mcmstypes "github.com/smartcontractkit/mcms/types"

	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"

	"github.com/smartcontractkit/chainlink-ccip/deployment/deploy"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/adapters"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/offchain"
)

type DeployChainContractsPerChainCfg struct {
	DeployerContract string
	DeployTestRouter bool
	RMNRemote        adapters.RMNRemoteDeployParams
	OffRamp          adapters.OffRampDeployParams
	OnRamp           adapters.OnRampDeployParams
	FeeQuoter        adapters.FeeQuoterDeployParams
	Executors        []adapters.ExecutorDeployParams
	MockReceivers    []adapters.MockReceiverDeployParams
	// DeployerKeyOwned, when true, skips the transfer-ownership step so that
	// contracts remain owned by the deployer key. By default (false) the
	// sequence transfers ownership of product contracts to the CLLCCIP
	// RBACTimelock, failing fast if the required MCMS instances are not found.
	DeployerKeyOwned bool
}

type DeployChainContractsCfg struct {
	Topology                                *offchain.EnvironmentTopology
	ChainSelectors                          []uint64
	DefaultCfg                              DeployChainContractsPerChainCfg
	ChainCfgs                               map[uint64]DeployChainContractsPerChainCfg
	IgnoreImportedConfigFromPreviousVersion bool // if true, the changeset will not import any config from previous version,
	// and will only rely on the config provided in DeployChainContractsCfg, else it will import config from previous version
	// and will give precedence to the imported config over the config provided in DeployChainContractsCfg.
}

func (c DeployChainContractsCfg) resolveChainCfg(sel uint64) DeployChainContractsPerChainCfg {
	if override, ok := c.ChainCfgs[sel]; ok {
		return override
	}
	return c.DefaultCfg
}

func DeployChainContracts(registry *adapters.DeployChainContractsRegistry) deployment.ChangeSetV2[changesets.WithMCMS[DeployChainContractsCfg]] {
	mcmsReaderRegistry := changesets.GetRegistry()
	validate := func(e deployment.Environment, cfg changesets.WithMCMS[DeployChainContractsCfg]) error {
		if cfg.Cfg.Topology == nil {
			return fmt.Errorf("topology is required")
		}

		if cfg.Cfg.Topology.NOPTopology == nil || len(cfg.Cfg.Topology.NOPTopology.Committees) == 0 {
			return fmt.Errorf("no committees defined in topology")
		}

		if len(cfg.Cfg.ChainSelectors) == 0 {
			return fmt.Errorf("at least one chain selector is required")
		}

		seen := make(map[uint64]bool, len(cfg.Cfg.ChainSelectors))
		envSelectors := e.BlockChains.ListChainSelectors()
		for _, sel := range cfg.Cfg.ChainSelectors {
			if seen[sel] {
				return fmt.Errorf("duplicate chain selector %d in ChainSelectors", sel)
			}
			seen[sel] = true
			if !slices.Contains(envSelectors, sel) {
				return fmt.Errorf("chain selector %d is not available in environment", sel)
			}

			perChain := cfg.Cfg.resolveChainCfg(sel)
			if perChain.DeployerContract == "" {
				return fmt.Errorf("DeployerContract is required for chain %d", sel)
			}
		}

		for sel := range cfg.Cfg.ChainCfgs {
			if !slices.Contains(cfg.Cfg.ChainSelectors, sel) {
				return fmt.Errorf("ChainCfgs contains selector %d which is not in ChainSelectors", sel)
			}
		}

		return nil
	}

	apply := func(e deployment.Environment, cfg changesets.WithMCMS[DeployChainContractsCfg]) (deployment.ChangesetOutput, error) {
		ds := datastore.NewMemoryDataStore()
		var allReports []operations.Report[any, any]
		var allBatchOps []mcmstypes.BatchOperation

		for _, sel := range cfg.Cfg.ChainSelectors {
			adapter, err := registry.GetByChain(sel)
			if err != nil {
				return deployment.ChangesetOutput{}, fmt.Errorf("failed to get deploy adapter for chain %d: %w", sel, err)
			}

			committeeVerifiers, err := BuildCommitteeVerifierParams(cfg.Cfg.Topology, sel)
			if err != nil {
				return deployment.ChangesetOutput{}, fmt.Errorf("failed to build committee verifier params for chain %d: %w", sel, err)
			}
			if len(committeeVerifiers) == 0 {
				return deployment.ChangesetOutput{}, fmt.Errorf("chain %d: no committees have chain_config for this selector", sel)
			}

			perChain := cfg.Cfg.resolveChainCfg(sel)

			existingAddresses := e.DataStore.Addresses().Filter(
				datastore.AddressRefByChainSelector(sel),
			)

			input := adapters.DeployChainContractsInput{
				ChainSelector:     sel,
				DeployerContract:  perChain.DeployerContract,
				DeployTestRouter:  perChain.DeployTestRouter,
				ExistingAddresses: existingAddresses,
				DeployerKeyOwned:  perChain.DeployerKeyOwned,
				ContractParams: adapters.DeployContractParams{
					RMNRemote:          perChain.RMNRemote,
					OffRamp:            perChain.OffRamp,
					CommitteeVerifiers: committeeVerifiers,
					OnRamp:             perChain.OnRamp,
					FeeQuoter:          perChain.FeeQuoter,
					Executors:          perChain.Executors,
					MockReceivers:      perChain.MockReceivers,
				},
			}
			if !cfg.Cfg.IgnoreImportedConfigFromPreviousVersion {
				importNeeded, err := shouldImportConfigFromPreviousVersion(e, registry, sel)
				if err != nil {
					return deployment.ChangesetOutput{}, fmt.Errorf("failed to check if config import from previous version is needed for chain %d: %w", sel, err)
				}
				if importNeeded {
					// so far we do not need any config from 1.5.0 , so only importing config from 1.6.0
					// if in the future we need to import some config from 1.5.0,
					// we should leverage lane version resolver to get the right adapter to import config
					version := semver.MustParse("1.6.0")
					configImporter, ok := registry.GetConfigImporter(sel, version)
					if !ok {
						return deployment.ChangesetOutput{}, utils.ErrNoAdapterForSelectorRegistered("ConfigImporter", sel, version)
					}
					populateCfgOutput, err := deploy.PopulateMetaDataFromConfigImporter(e, configImporter, sel)
					if err != nil {
						return deployment.ChangesetOutput{}, fmt.Errorf("failed to populate metadata from config importer for chain %d: %w", sel, err)
					}
					// if the importer found config from previous version,
					// import it and merge it with the input config, giving precedence to the imported config
					if len(populateCfgOutput.Metadata.Contracts) > 0 {
						importReport, err := operations.ExecuteSequence(
							e.OperationsBundle,
							adapter.SetContractParamsFromImportedConfig(),
							e.BlockChains, adapters.DeployChainConfigCreatorInput{
								ChainSelector:      sel,
								ContractMeta:       populateCfgOutput.Metadata.Contracts,
								ExistingAddresses:  existingAddresses,
								UserProvidedConfig: input.ContractParams,
							})
						if err != nil {
							return deployment.ChangesetOutput{Reports: allReports},
								fmt.Errorf("failed to execute config importer sequence for chain %d: %w", sel, err)
						}
						allReports = append(allReports, importReport.ExecutionReports...)
						input.ContractParams, err = input.ContractParams.MergeWithOverrideIfNotEmpty(importReport.Output)
						if err != nil {
							return deployment.ChangesetOutput{}, fmt.Errorf("failed to merge imported config with input config for chain %d: %w", sel, err)
						}
					}
				}
			}

			e.Logger.Infow(
				"Deploying chain contracts with topology-derived committee verifiers",
				"chain", sel,
				"committees", len(committeeVerifiers),
			)

			report, err := operations.ExecuteSequence(e.OperationsBundle, adapter.DeployChainContracts(), e.BlockChains, input)
			if err != nil {
				return deployment.ChangesetOutput{Reports: allReports},
					fmt.Errorf("failed to deploy chain contracts on chain %d: %w", sel, err)
			}

			for _, ref := range report.Output.Addresses {
				if addErr := ds.Addresses().Add(ref); addErr != nil {
					return deployment.ChangesetOutput{Reports: allReports},
						fmt.Errorf("failed to add %s %s at %s on chain %d to datastore: %w",
							ref.Type, ref.Version, ref.Address, ref.ChainSelector, addErr)
				}
			}

			if writeErr := sequences.WriteMetadataToDatastore(ds, report.Output.Metadata); writeErr != nil {
				return deployment.ChangesetOutput{Reports: allReports},
					fmt.Errorf("failed to write metadata to datastore for chain %d: %w", sel, writeErr)
			}

			allReports = append(allReports, report.ExecutionReports...)
			allBatchOps = append(allBatchOps, report.Output.BatchOps...)
		}

		return changesets.NewOutputBuilder(e, mcmsReaderRegistry).
			WithReports(allReports).
			WithDataStore(ds).
			WithBatchOps(allBatchOps).
			Build(cfg.MCMS)
	}

	return deployment.CreateChangeSet(apply, validate)
}

// BuildCommitteeVerifierParams extracts committee verifier deployment parameters
// from the topology for a specific chain. Committees without a chain_config
// entry for the given selector are skipped.
// Address format validation is deferred to the chain-family adapter.
func BuildCommitteeVerifierParams(
	topology *offchain.EnvironmentTopology,
	chainSelector uint64,
) ([]adapters.CommitteeVerifierDeployParams, error) {
	if topology == nil {
		return nil, fmt.Errorf("topology is nil")
	}
	if topology.NOPTopology == nil {
		return nil, fmt.Errorf("NOPTopology is nil")
	}

	chainKey := strconv.FormatUint(chainSelector, 10)

	qualifiers := make([]string, 0, len(topology.NOPTopology.Committees))
	for q := range topology.NOPTopology.Committees {
		qualifiers = append(qualifiers, q)
	}
	sort.Strings(qualifiers)

	params := make([]adapters.CommitteeVerifierDeployParams, 0, len(qualifiers))
	for _, qualifier := range qualifiers {
		committee := topology.NOPTopology.Committees[qualifier]

		chainCfg, ok := committee.ChainConfigs[chainKey]
		if !ok {
			continue
		}

		if committee.VerifierVersion == nil {
			return nil, fmt.Errorf("committee %q has nil VerifierVersion", qualifier)
		}

		if chainCfg.FeeAggregator == "" {
			return nil, fmt.Errorf("committee %q: FeeAggregator is required", qualifier)
		}

		params = append(params, adapters.CommitteeVerifierDeployParams{
			Version:          committee.VerifierVersion,
			FeeAggregator:    chainCfg.FeeAggregator,
			AllowlistAdmin:   chainCfg.AllowlistAdmin,
			StorageLocations: committee.StorageLocations,
			Qualifier:        qualifier,
		})
	}

	return params, nil
}

// shouldImportConfigFromPreviousVersion checks if any of the lanes connected to the given chain selector are using version 1.6.0,
// if so, it returns true, indicating that we should import config from previous version, otherwise it returns false
func shouldImportConfigFromPreviousVersion(e deployment.Environment, registry *adapters.DeployChainContractsRegistry, sel uint64) (bool, error) {
	res, exists := registry.GetLaneVersionResolver(sel)
	if !exists {
		return false, fmt.Errorf("no lane version resolver registered for chain %d", sel)
	}
	if !res.IsSupportedChain(e, sel) {
		return false, nil
	}
	_, versions, err := res.DeriveLaneVersionsForChain(e, sel)
	if err != nil {
		return false, fmt.Errorf("failed to derive lane versions for chain %d: %w", sel, err)
	}
	reqVersion := semver.MustParse("1.6.0")
	for _, v := range versions {
		if v.Equal(reqVersion) {
			return true, nil
		}
	}
	return false, nil
}
