package changesets

import (
	"fmt"
	"slices"
	"strconv"

	"github.com/BurntSushi/toml"

	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"

	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/adapters"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/offchain"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/offchain/operations/fetch_node_chain_support"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/offchain/sequences"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/offchain/shared"
)

type ApplyExecutorConfigInput struct {
	Topology          *offchain.EnvironmentTopology
	ExecutorQualifier string
	TargetNOPs        []shared.NOPAlias
	// RevokeOrphanedJobs when true revokes and cleans up orphaned jobs; default false.
	RevokeOrphanedJobs bool
}

func ApplyExecutorConfig(registry *adapters.ExecutorConfigRegistry) deployment.ChangeSetV2[ApplyExecutorConfigInput] {
	validate := func(e deployment.Environment, cfg ApplyExecutorConfigInput) error {
		if cfg.Topology == nil {
			return fmt.Errorf("topology is required")
		}

		if cfg.ExecutorQualifier == "" {
			return fmt.Errorf("executor qualifier is required")
		}

		if cfg.Topology.NOPTopology == nil || len(cfg.Topology.NOPTopology.NOPs) == 0 {
			return fmt.Errorf("NOP topology with at least one NOP is required")
		}

		if len(cfg.Topology.IndexerAddress) == 0 {
			return fmt.Errorf("indexer address is required in topology")
		}

		pool, ok := cfg.Topology.ExecutorPools[cfg.ExecutorQualifier]
		if !ok {
			return fmt.Errorf("executor pool %q not found in topology", cfg.ExecutorQualifier)
		}

		if len(pool.ChainConfigs) == 0 {
			return fmt.Errorf("executor pool %q requires non-empty chain_configs", cfg.ExecutorQualifier)
		}

		poolNOPs := getExecutorPoolNOPAliases(pool)
		if len(poolNOPs) == 0 {
			return fmt.Errorf("executor pool %q has no NOPs", cfg.ExecutorQualifier)
		}
		for _, alias := range cfg.TargetNOPs {
			if !slices.Contains(poolNOPs, alias) {
				return fmt.Errorf("NOP alias %q not found in executor pool %q", alias, cfg.ExecutorQualifier)
			}
		}

		if shared.IsProductionEnvironment(e.Name) {
			if cfg.Topology.PyroscopeURL != "" {
				return fmt.Errorf("pyroscope URL is not supported for production environments")
			}
		}
		return nil
	}

	apply := func(e deployment.Environment, cfg ApplyExecutorConfigInput) (deployment.ChangesetOutput, error) {
		selectors := registry.AllDeployedChains(e.DataStore, cfg.ExecutorQualifier)
		pool := cfg.Topology.ExecutorPools[cfg.ExecutorQualifier]

		if len(selectors) == 0 {
			if !cfg.RevokeOrphanedJobs {
				e.Logger.Infow("No deployed chains found for executor pool, nothing to do",
					"qualifier", cfg.ExecutorQualifier)
				ds := datastore.NewMemoryDataStore()
				if e.DataStore != nil {
					if err := ds.Merge(e.DataStore); err != nil {
						return deployment.ChangesetOutput{}, fmt.Errorf("failed to merge datastore: %w", err)
					}
				}
				return deployment.ChangesetOutput{DataStore: ds}, nil
			}
			e.Logger.Infow("No deployed chains for executor pool, running orphan cleanup only",
				"qualifier", cfg.ExecutorQualifier)
			nopModes := buildNOPModes(cfg.Topology.NOPTopology.NOPs)
			scope := shared.ExecutorJobScope{ExecutorQualifier: cfg.ExecutorQualifier}
			manageReport, err := operations.ExecuteSequence(
				e.OperationsBundle,
				sequences.ManageJobProposals,
				sequences.ManageJobProposalsDeps{Env: e},
				sequences.ManageJobProposalsInput{
					JobSpecs:           nil,
					AffectedScope:      scope,
					Labels:             map[string]string{"job_type": "executor", "executor": cfg.ExecutorQualifier},
					NOPs:               sequences.NOPContext{Modes: nopModes, TargetNOPs: cfg.TargetNOPs, AllNOPs: getAllNOPAliases(cfg.Topology.NOPTopology.NOPs)},
					RevokeOrphanedJobs: true,
				},
			)
			if err != nil {
				return deployment.ChangesetOutput{Reports: manageReport.ExecutionReports}, fmt.Errorf("failed to manage job proposals (orphan cleanup): %w", err)
			}
			return deployment.ChangesetOutput{Reports: manageReport.ExecutionReports, DataStore: manageReport.Output.DataStore}, nil
		}

		nopsToValidate := cfg.TargetNOPs
		if len(nopsToValidate) == 0 {
			nopsToValidate = getExecutorPoolNOPAliases(pool)
		}

		clNOPs := filterCLModeNOPs(nopsToValidate, cfg.Topology.NOPTopology.NOPs)
		if err := validateExecutorChainSupport(e, pool, clNOPs, selectors); err != nil {
			return deployment.ChangesetOutput{}, err
		}

		chainConfigs, err := buildChainConfigs(registry, e.DataStore, selectors, cfg.ExecutorQualifier)
		if err != nil {
			return deployment.ChangesetOutput{}, err
		}

		monitoring := convertTopologyMonitoring(&cfg.Topology.Monitoring)
		nopModes := buildNOPModes(cfg.Topology.NOPTopology.NOPs)

		jobSpecs, scope, err := buildExecutorJobSpecs(
			chainConfigs,
			cfg.ExecutorQualifier,
			cfg.TargetNOPs,
			pool,
			cfg.Topology.IndexerAddress,
			cfg.Topology.PyroscopeURL,
			monitoring,
		)
		if err != nil {
			return deployment.ChangesetOutput{}, err
		}

		manageReport, err := operations.ExecuteSequence(
			e.OperationsBundle,
			sequences.ManageJobProposals,
			sequences.ManageJobProposalsDeps{Env: e},
			sequences.ManageJobProposalsInput{
				JobSpecs:      jobSpecs,
				AffectedScope: scope,
				Labels: map[string]string{
					"job_type": "executor",
					"executor": cfg.ExecutorQualifier,
				},
				NOPs: sequences.NOPContext{
					Modes:      nopModes,
					TargetNOPs: cfg.TargetNOPs,
					AllNOPs:    getAllNOPAliases(cfg.Topology.NOPTopology.NOPs),
				},
				RevokeOrphanedJobs: cfg.RevokeOrphanedJobs,
			},
		)
		if err != nil {
			return deployment.ChangesetOutput{
				Reports: manageReport.ExecutionReports,
			}, fmt.Errorf("failed to manage job proposals: %w", err)
		}

		e.Logger.Infow("Executor config applied",
			"jobsCount", len(manageReport.Output.Jobs),
			"revokedCount", len(manageReport.Output.RevokedJobs))

		return deployment.ChangesetOutput{
			Reports:   manageReport.ExecutionReports,
			DataStore: manageReport.Output.DataStore,
		}, nil
	}

	return deployment.CreateChangeSet(apply, validate)
}

func buildChainConfigs(
	registry *adapters.ExecutorConfigRegistry,
	ds datastore.DataStore,
	selectors []uint64,
	qualifier string,
) (map[string]adapters.ExecutorChainConfig, error) {
	chainConfigs := make(map[string]adapters.ExecutorChainConfig, len(selectors))
	for _, sel := range selectors {
		adapter, err := registry.GetByChain(sel)
		if err != nil {
			return nil, fmt.Errorf("no adapter for chain %d: %w", sel, err)
		}
		cfg, err := adapter.BuildChainConfig(ds, sel, qualifier)
		if err != nil {
			return nil, fmt.Errorf("failed to build config for chain %d: %w", sel, err)
		}
		chainConfigs[strconv.FormatUint(sel, 10)] = cfg
	}
	return chainConfigs, nil
}

func getExecutorPoolNOPAliases(pool offchain.ExecutorPoolConfig) []shared.NOPAlias {
	aliasSet := make(map[string]struct{})
	for _, chainCfg := range pool.ChainConfigs {
		for _, alias := range chainCfg.NOPAliases {
			aliasSet[alias] = struct{}{}
		}
	}
	aliases := make([]string, 0, len(aliasSet))
	for a := range aliasSet {
		aliases = append(aliases, a)
	}
	slices.Sort(aliases)
	return shared.ConvertStringToNopAliases(aliases)
}

func validateExecutorChainSupport(
	e deployment.Environment,
	pool offchain.ExecutorPoolConfig,
	nopsToValidate []shared.NOPAlias,
	deployedChains []uint64,
) error {
	if e.Offchain == nil {
		e.Logger.Debugw("Offchain client not available, skipping chain support validation")
		return nil
	}

	nopAliasStrings := shared.ConvertNopAliasToString(nopsToValidate)

	supportedChains, err := fetchNodeChainSupport(e, nopAliasStrings)
	if err != nil {
		return fmt.Errorf("failed to fetch node chain support: %w", err)
	}
	if supportedChains == nil {
		return nil
	}

	var validationResults []shared.ChainValidationResult
	for _, nopAlias := range nopsToValidate {
		requiredChains, err := getRequiredChainsForExecutorNOP(string(nopAlias), pool, deployedChains)
		if err != nil {
			return err
		}
		result := shared.ValidateNOPChainSupport(
			string(nopAlias),
			requiredChains,
			supportedChains[string(nopAlias)],
		)
		if result != nil {
			validationResults = append(validationResults, *result)
		}
	}

	return shared.FormatChainValidationError(validationResults)
}

func getRequiredChainsForExecutorNOP(nopAlias string, pool offchain.ExecutorPoolConfig, deployedChains []uint64) ([]uint64, error) {
	var requiredChains []uint64
	for chainSelectorStr, chainCfg := range pool.ChainConfigs {
		if !slices.Contains(chainCfg.NOPAliases, nopAlias) {
			continue
		}
		sel, err := strconv.ParseUint(chainSelectorStr, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("executor pool chain_configs key %q is not a valid chain selector: %w", chainSelectorStr, err)
		}
		requiredChains = append(requiredChains, sel)
	}
	return requiredChains, nil
}

func fetchNodeChainSupport(e deployment.Environment, nopAliases []string) (shared.ChainSupportByNOP, error) {
	if len(nopAliases) == 0 {
		return nil, nil
	}

	report, err := operations.ExecuteOperation(
		e.OperationsBundle,
		fetch_node_chain_support.FetchNodeChainSupport,
		fetch_node_chain_support.FetchNodeChainSupportDeps{
			JDClient: e.Offchain,
			Logger:   e.Logger,
			NodeIDs:  e.NodeIDs,
		},
		fetch_node_chain_support.FetchNodeChainSupportInput{
			NOPAliases: nopAliases,
		},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch node chain support from JD: %w", err)
	}

	return report.Output.SupportedChains, nil
}

func buildExecutorJobSpecs(
	chainConfigs map[string]adapters.ExecutorChainConfig,
	executorQualifier string,
	targetNOPs []shared.NOPAlias,
	pool offchain.ExecutorPoolConfig,
	indexerAddress []string,
	pyroscopeURL string,
	monitoring shared.MonitoringInput,
) (shared.NOPJobSpecs, shared.ExecutorJobScope, error) {
	scope := shared.ExecutorJobScope{
		ExecutorQualifier: executorQualifier,
	}

	poolNOPs := getExecutorPoolNOPAliases(pool)
	nopAliases := targetNOPs
	if len(nopAliases) == 0 {
		nopAliases = poolNOPs
	}

	jobSpecs := make(shared.NOPJobSpecs)

	for _, nopAlias := range nopAliases {
		chainCfgs := make(map[string]offchain.ExecutorChainCfg)
		for chainSelectorStr, genCfg := range chainConfigs {
			chainCfg, ok := pool.ChainConfigs[chainSelectorStr]
			if !ok {
				continue
			}
			chainCfgs[chainSelectorStr] = offchain.ExecutorChainCfg{
				OffRampAddress:         genCfg.OffRampAddress,
				RmnAddress:             genCfg.RmnAddress,
				DefaultExecutorAddress: genCfg.ExecutorProxyAddress,
				ExecutorPool:           chainCfg.NOPAliases,
				ExecutionInterval:      chainCfg.ExecutionInterval,
			}
		}

		jobSpecID := shared.NewExecutorJobID(nopAlias, scope)

		executorCfg := offchain.ExecutorConfiguration{
			IndexerAddress:    indexerAddress,
			ExecutorID:        jobSpecID.GetExecutorID(),
			PyroscopeURL:      pyroscopeURL,
			NtpServer:         pool.NtpServer,
			IndexerQueryLimit: pool.IndexerQueryLimit,
			BackoffDuration:   pool.BackoffDuration,
			LookbackWindow:    pool.LookbackWindow,
			ReaderCacheExpiry: pool.ReaderCacheExpiry,
			MaxRetryDuration:  pool.MaxRetryDuration,
			WorkerCount:       pool.WorkerCount,
			Monitoring: offchain.ExecutorMonitoringConfig{
				Enabled: monitoring.Enabled,
				Type:    monitoring.Type,
				Beholder: offchain.ExecutorBeholderConfig{
					InsecureConnection:       monitoring.Beholder.InsecureConnection,
					CACertFile:               monitoring.Beholder.CACertFile,
					OtelExporterGRPCEndpoint: monitoring.Beholder.OtelExporterGRPCEndpoint,
					OtelExporterHTTPEndpoint: monitoring.Beholder.OtelExporterHTTPEndpoint,
					LogStreamingEnabled:      monitoring.Beholder.LogStreamingEnabled,
					MetricReaderInterval:     monitoring.Beholder.MetricReaderInterval,
					TraceSampleRatio:         monitoring.Beholder.TraceSampleRatio,
					TraceBatchTimeout:        monitoring.Beholder.TraceBatchTimeout,
				},
			},
			ChainConfiguration: chainCfgs,
		}

		configBytes, err := toml.Marshal(executorCfg)
		if err != nil {
			return nil, scope, fmt.Errorf("failed to marshal executor config for NOP %q: %w", nopAlias, err)
		}

		jobID := jobSpecID.ToJobID()
		jobSpec := fmt.Sprintf(`schemaVersion = 1
type = "ccvexecutor"
name = "%s"
externalJobID = "%s"
executorConfig = """
%s"""
`, string(jobID), jobID.ToExternalJobID(), string(configBytes))

		if jobSpecs[nopAlias] == nil {
			jobSpecs[nopAlias] = make(map[shared.JobID]string)
		}
		jobSpecs[nopAlias][jobID] = jobSpec
	}

	return jobSpecs, scope, nil
}

func convertTopologyMonitoring(m *offchain.MonitoringConfig) shared.MonitoringInput {
	if m == nil {
		return shared.MonitoringInput{}
	}
	return shared.MonitoringInput{
		Enabled: m.Enabled,
		Type:    m.Type,
		Beholder: shared.BeholderInput{
			InsecureConnection:       m.Beholder.InsecureConnection,
			CACertFile:               m.Beholder.CACertFile,
			OtelExporterGRPCEndpoint: m.Beholder.OtelExporterGRPCEndpoint,
			OtelExporterHTTPEndpoint: m.Beholder.OtelExporterHTTPEndpoint,
			LogStreamingEnabled:      m.Beholder.LogStreamingEnabled,
			MetricReaderInterval:     m.Beholder.MetricReaderInterval,
			TraceSampleRatio:         m.Beholder.TraceSampleRatio,
			TraceBatchTimeout:        m.Beholder.TraceBatchTimeout,
		},
	}
}

func buildNOPModes(nops []offchain.NOPConfig) map[shared.NOPAlias]shared.NOPMode {
	nopModes := make(map[shared.NOPAlias]shared.NOPMode)
	for _, nop := range nops {
		mode := nop.GetMode()
		nopModes[shared.NOPAlias(nop.Alias)] = mode
	}
	return nopModes
}

func filterCLModeNOPs(aliases []shared.NOPAlias, nops []offchain.NOPConfig) []shared.NOPAlias {
	modeByAlias := buildNOPModes(nops)
	filtered := make([]shared.NOPAlias, 0, len(aliases))
	for _, alias := range aliases {
		if mode, ok := modeByAlias[alias]; ok && mode == shared.NOPModeCL {
			filtered = append(filtered, alias)
		}
	}
	return filtered
}

func getAllNOPAliases(nops []offchain.NOPConfig) []shared.NOPAlias {
	aliases := make([]shared.NOPAlias, len(nops))
	for i, nop := range nops {
		aliases[i] = shared.NOPAlias(nop.Alias)
	}
	return aliases
}
