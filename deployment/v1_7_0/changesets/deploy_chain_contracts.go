package changesets

import (
	"fmt"
	"slices"
	"sort"
	"strconv"

	mcmstypes "github.com/smartcontractkit/mcms/types"

	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/adapters"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/offchain"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
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
	// MCMS configures deployment of CLLCCIP and RMNMCMS instances.
	// When non-nil, both instances are deployed and ownership of product contracts
	// is transferred to the CLLCCIP timelock.
	MCMS *adapters.MCMSDeployParams
}

type DeployChainContractsCfg struct {
	Topology       *offchain.EnvironmentTopology
	ChainSelectors []uint64
	DefaultCfg     DeployChainContractsPerChainCfg
	ChainCfgs      map[uint64]DeployChainContractsPerChainCfg
}

func (c DeployChainContractsCfg) resolveChainCfg(sel uint64) DeployChainContractsPerChainCfg {
	if override, ok := c.ChainCfgs[sel]; ok {
		return override
	}
	return c.DefaultCfg
}

func DeployChainContracts(
	registry *adapters.DeployChainContractsRegistry,
	mcmsReaderRegistry *changesets.MCMSReaderRegistry,
) deployment.ChangeSetV2[changesets.WithMCMS[DeployChainContractsCfg]] {
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
				MCMS:              perChain.MCMS,
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
