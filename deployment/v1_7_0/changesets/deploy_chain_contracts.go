package changesets

import (
	"fmt"
	"slices"
	"sort"
	"strconv"

	"github.com/ethereum/go-ethereum/common"

	mcmstypes "github.com/smartcontractkit/mcms/types"

	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/sequences"
	changesetscore "github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	cldfsequences "github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/topology"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

type DeployChainContractsFromTopologyCfgPerChain struct {
	CREATE2Factory   common.Address
	DeployTestRouter bool
	RMNRemote        sequences.RMNRemoteParams
	OffRamp          sequences.OffRampParams
	OnRamp           sequences.OnRampParams
	FeeQuoter        sequences.FeeQuoterParams
	Executors        []sequences.ExecutorParams
	MockReceivers    []sequences.MockReceiverParams
}

type DeployChainContractsFromTopologyCfg struct {
	Topology       *topology.EnvironmentTopology
	ChainSelectors []uint64
	DefaultCfg     DeployChainContractsFromTopologyCfgPerChain
	ChainCfgs      map[uint64]DeployChainContractsFromTopologyCfgPerChain
}

func (c DeployChainContractsFromTopologyCfg) resolveChainCfg(sel uint64) DeployChainContractsFromTopologyCfgPerChain {
	if override, ok := c.ChainCfgs[sel]; ok {
		return override
	}
	return c.DefaultCfg
}

func DeployChainContractsFromTopology(
	mcmsReaderRegistry *changesetscore.MCMSReaderRegistry,
) deployment.ChangeSetV2[changesetscore.WithMCMS[DeployChainContractsFromTopologyCfg]] {
	validate := func(e deployment.Environment, cfg changesetscore.WithMCMS[DeployChainContractsFromTopologyCfg]) error {
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
		evmChains := e.BlockChains.EVMChains()
		for _, sel := range cfg.Cfg.ChainSelectors {
			if seen[sel] {
				return fmt.Errorf("duplicate chain selector %d in ChainSelectors", sel)
			}
			seen[sel] = true
			if !slices.Contains(envSelectors, sel) {
				return fmt.Errorf("chain selector %d is not available in environment", sel)
			}
			if _, ok := evmChains[sel]; !ok {
				return fmt.Errorf("chain selector %d is not an EVM chain", sel)
			}

			perChain := cfg.Cfg.resolveChainCfg(sel)
			if perChain.CREATE2Factory == (common.Address{}) {
				return fmt.Errorf("CREATE2Factory address is required for chain %d", sel)
			}
		}

		for sel := range cfg.Cfg.ChainCfgs {
			if !slices.Contains(cfg.Cfg.ChainSelectors, sel) {
				return fmt.Errorf("ChainCfgs contains selector %d which is not in ChainSelectors", sel)
			}
		}

		return nil
	}

	apply := func(e deployment.Environment, cfg changesetscore.WithMCMS[DeployChainContractsFromTopologyCfg]) (deployment.ChangesetOutput, error) {
		evmChains := e.BlockChains.EVMChains()
		ds := datastore.NewMemoryDataStore()
		var allReports []operations.Report[any, any]
		var allBatchOps []mcmstypes.BatchOperation

		for _, sel := range cfg.Cfg.ChainSelectors {
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

			chain, ok := evmChains[sel]
			if !ok {
				return deployment.ChangesetOutput{}, fmt.Errorf("EVM chain not found for selector %d", sel)
			}

			input := sequences.DeployChainContractsInput{
				ChainSelector:     sel,
				CREATE2Factory:    perChain.CREATE2Factory,
				ExistingAddresses: existingAddresses,
				ContractParams: sequences.ContractParams{
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

			report, err := operations.ExecuteSequence(e.OperationsBundle, sequences.DeployChainContracts, chain, input)
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

			if writeErr := cldfsequences.WriteMetadataToDatastore(ds, report.Output.Metadata); writeErr != nil {
				return deployment.ChangesetOutput{Reports: allReports},
					fmt.Errorf("failed to write metadata to datastore for chain %d: %w", sel, writeErr)
			}

			allReports = append(allReports, report.ExecutionReports...)
			allBatchOps = append(allBatchOps, report.Output.BatchOps...)
		}

		return changesetscore.NewOutputBuilder(e, mcmsReaderRegistry).
			WithReports(allReports).
			WithDataStore(ds).
			WithBatchOps(allBatchOps).
			Build(cfg.MCMS)
	}

	return deployment.CreateChangeSet(apply, validate)
}

// BuildCommitteeVerifierParams builds CommitteeVerifierParams from the topology
// for a specific chain. FeeAggregator and AllowlistAdmin are read from each
// committee's per-chain config (ChainCommitteeConfig).
// Committees that don't have a chain_config entry for this selector are skipped.
func BuildCommitteeVerifierParams(
	topology *topology.EnvironmentTopology,
	chainSelector uint64,
) ([]sequences.CommitteeVerifierParams, error) {
	if topology.NOPTopology == nil {
		return nil, fmt.Errorf("NOPTopology is nil")
	}

	chainKey := strconv.FormatUint(chainSelector, 10)

	qualifiers := make([]string, 0, len(topology.NOPTopology.Committees))
	for q := range topology.NOPTopology.Committees {
		qualifiers = append(qualifiers, q)
	}
	sort.Strings(qualifiers)

	params := make([]sequences.CommitteeVerifierParams, 0, len(qualifiers))
	for _, qualifier := range qualifiers {
		committee := topology.NOPTopology.Committees[qualifier]

		chainCfg, ok := committee.ChainConfigs[chainKey]
		if !ok {
			continue
		}

		if committee.VerifierVersion == nil {
			return nil, fmt.Errorf("committee %q has nil VerifierVersion", qualifier)
		}

		feeAggregator, err := parseRequiredAddress(chainCfg.FeeAggregator, "FeeAggregator", qualifier)
		if err != nil {
			return nil, err
		}

		allowlistAdmin := common.Address{}
		if chainCfg.AllowlistAdmin != "" {
			if !common.IsHexAddress(chainCfg.AllowlistAdmin) {
				return nil, fmt.Errorf("committee %q chain %d: AllowlistAdmin %q is not a valid hex address",
					qualifier, chainSelector, chainCfg.AllowlistAdmin)
			}
			allowlistAdmin = common.HexToAddress(chainCfg.AllowlistAdmin)
		}

		params = append(params, sequences.CommitteeVerifierParams{
			Version:          committee.VerifierVersion,
			FeeAggregator:    feeAggregator,
			AllowlistAdmin:   allowlistAdmin,
			StorageLocations: committee.StorageLocations,
			Qualifier:        qualifier,
		})
	}

	return params, nil
}

func parseRequiredAddress(hex, field, qualifier string) (common.Address, error) {
	if hex == "" {
		return common.Address{}, fmt.Errorf("committee %q: %s is required", qualifier, field)
	}
	if !common.IsHexAddress(hex) {
		return common.Address{}, fmt.Errorf("committee %q: %s %q is not a valid hex address", qualifier, field, hex)
	}
	addr := common.HexToAddress(hex)
	if addr == (common.Address{}) {
		return common.Address{}, fmt.Errorf("committee %q: %s cannot be zero address", qualifier, field)
	}
	return addr, nil
}
