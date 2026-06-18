package changesets

import (
	"fmt"
	"slices"
	"sort"
	"strconv"

	mcmstypes "github.com/smartcontractkit/mcms/types"

	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"

	"github.com/smartcontractkit/chainlink-ccip/deployment/deploy"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v2_0_0/adapters"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v2_0_0/offchain"
)

// DeployChainContractsPerChainCfg holds optional per-chain deploy settings.
// ContractParams is the only field whose unset pointer sub-fields are merged with
// adapter defaults via BuildDeployContractParams; other fields coalesce against
// changeset defaults (resolved deployer address, false for both deploy flags).
type DeployChainContractsPerChainCfg struct {
	DeployerContract *string                                 `json:"deployerContract,omitempty" yaml:"deployerContract,omitempty"`
	DeployTestRouter bool                                    `json:"deployTestRouter,omitempty" yaml:"deployTestRouter,omitempty"`
	DeployerKeyOwned bool                                    `json:"deployerKeyOwned,omitempty" yaml:"deployerKeyOwned,omitempty"`
	ContractParams   *adapters.DeployContractParamsOverrides `json:"contractParams,omitempty" yaml:"contractParams,omitempty"`
}

// DeployChainContractsCfg is the durable-pipeline payload for deploying CCIP 2.0 chain contracts.
// Contract versions, fee-quoter static config, CREATE2 factory resolution, and
// executor/mock-receiver wiring are resolved via the chain-family DeployChainContractsAdapter;
// topology supplies committee verifiers.
type DeployChainContractsCfg struct {
	Topology       *offchain.EnvironmentTopology              `json:"topology" yaml:"topology"`
	ChainSelectors []uint64                                   `json:"chainSelectors" yaml:"chainSelectors"`
	ChainOverrides map[uint64]DeployChainContractsPerChainCfg `json:"chainOverrides,omitempty" yaml:"chainOverrides,omitempty"`
}

func (c DeployChainContractsCfg) resolvePerChainCfg(sel uint64) DeployChainContractsPerChainCfg {
	if c.ChainOverrides == nil {
		return DeployChainContractsPerChainCfg{}
	}
	return c.ChainOverrides[sel]
}

func DeployChainContracts(registry *adapters.DeployChainContractsRegistry, chainFamilyRegistry *adapters.ChainFamilyRegistry) deployment.ChangeSetV2[changesets.WithMCMS[DeployChainContractsCfg]] {
	mcmsReaderRegistry := changesets.GetRegistry()
	validate := func(e deployment.Environment, cfg changesets.WithMCMS[DeployChainContractsCfg]) error {
		if cfg.Cfg.Topology == nil {
			return fmt.Errorf("topology is required")
		}
		if err := cfg.Cfg.Topology.ValidateForEnvironment(e.Name, chainFamilyRegistry); err != nil {
			return fmt.Errorf("topology validation failed: %w", err)
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
				return fmt.Errorf("duplicate chain selector %d in chainSelectors", sel)
			}
			seen[sel] = true
			if !slices.Contains(envSelectors, sel) {
				return fmt.Errorf("chain selector %d is not available in environment", sel)
			}
		}

		for sel := range cfg.Cfg.ChainOverrides {
			if !slices.Contains(cfg.Cfg.ChainSelectors, sel) {
				return fmt.Errorf("chainOverrides contains selector %d which is not in chainSelectors", sel)
			}
			if cfg.Cfg.ChainOverrides[sel].DeployerContract != nil && *cfg.Cfg.ChainOverrides[sel].DeployerContract == "" {
				return fmt.Errorf("chain %d: deployerContract override cannot be empty string", sel)
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

			perChainCfg := cfg.Cfg.resolvePerChainCfg(sel)

			resolved, resolveErr := adapter.ResolveDeployAddresses(e, sel)
			if resolveErr != nil {
				return deployment.ChangesetOutput{}, fmt.Errorf("failed to resolve deploy addresses on chain %d: %w", sel, resolveErr)
			}

			contractParams, err := adapter.BuildDeployContractParams(adapters.BuildDeployContractParamsInput{
				ChainSelector:      sel,
				CommitteeVerifiers: committeeVerifiers,
				Overrides:          perChainCfg.ContractParams,
				Defaults:           adapter.GetDefaultDeployContractParams(sel),
			})
			if err != nil {
				return deployment.ChangesetOutput{}, fmt.Errorf("failed to build contract params for chain %d: %w", sel, err)
			}

			existingAddresses := e.DataStore.Addresses().Filter(
				datastore.AddressRefByChainSelector(sel),
			)
			existingAddresses = append(existingAddresses, resolved.NewAddressRefs...)

			deployerContract := utils.Coalesce(perChainCfg.DeployerContract, resolved.DeployerContract)

			input := adapters.DeployChainContractsInput{
				ChainSelector:     sel,
				DeployerContract:  deployerContract,
				DeployTestRouter:  perChainCfg.DeployTestRouter,
				ExistingAddresses: existingAddresses,
				DeployerKeyOwned:  perChainCfg.DeployerKeyOwned,
				ContractParams:    contractParams,
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

			if !input.DeployerKeyOwned {
				ownershipBatchOps, ownershipReports, err := deploy.TransferToTimelock(sel, e, cfg.MCMS, report.Output.RefsToTransferOwnership)
				if err != nil {
					return deployment.ChangesetOutput{Reports: allReports},
						fmt.Errorf("failed to transfer to timelock for chain %d: %w", sel, err)
				}
				allReports = append(allReports, ownershipReports...)
				allBatchOps = append(allBatchOps, ownershipBatchOps...)
			}

			for _, ref := range resolved.NewAddressRefs {
				if addErr := ds.Addresses().Add(ref); addErr != nil {
					return deployment.ChangesetOutput{Reports: allReports},
						fmt.Errorf("failed to add resolved %s %s at %s on chain %d to datastore: %w",
							ref.Type, ref.Version, ref.Address, ref.ChainSelector, addErr)
				}
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
