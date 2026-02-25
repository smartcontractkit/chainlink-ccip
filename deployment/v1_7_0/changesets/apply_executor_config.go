package changesets

import (
	"fmt"
	"slices"

	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"

	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/operations/shared"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/topology"
	executorconfig "github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/operations/executor_config"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/operations/fetch_node_chain_support"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/sequences"
)

type ApplyExecutorConfigCfg struct {
	Topology *topology.EnvironmentTopology
	// ExecutorQualifier identifies which executor pool from topology to use
	ExecutorQualifier string
	// ChainSelectors limits which chains to configure. Defaults to all.
	ChainSelectors []uint64
	// TargetNOPs limits which NOPs to update. Defaults to all in pool.
	TargetNOPs []shared.NOPAlias
}

type ExecutorApplyDeps struct {
	Env      deployment.Environment
	JDClient shared.JDClient
	NodeIDs  []string
}

func makeExecutorApply(
	applyFn func(ExecutorApplyDeps, ApplyExecutorConfigCfg) (deployment.ChangesetOutput, error),
) func(deployment.Environment, ApplyExecutorConfigCfg) (deployment.ChangesetOutput, error) {
	return func(e deployment.Environment, cfg ApplyExecutorConfigCfg) (deployment.ChangesetOutput, error) {
		return applyFn(ExecutorApplyDeps{
			Env:      e,
			JDClient: e.Offchain,
			NodeIDs:  e.NodeIDs,
		}, cfg)
	}
}

func ApplyExecutorConfig() deployment.ChangeSetV2[ApplyExecutorConfigCfg] {
	validate := func(e deployment.Environment, cfg ApplyExecutorConfigCfg) error {
		if cfg.Topology == nil {
			return fmt.Errorf("topology is required")
		}

		if cfg.ExecutorQualifier == "" {
			return fmt.Errorf("executor qualifier is required")
		}

		if len(cfg.Topology.IndexerAddress) == 0 {
			return fmt.Errorf("indexer address is required in topology")
		}

		pool, ok := cfg.Topology.ExecutorPools[cfg.ExecutorQualifier]
		if !ok {
			return fmt.Errorf("executor pool %q not found in topology", cfg.ExecutorQualifier)
		}

		if len(pool.NOPAliases) == 0 {
			return fmt.Errorf("executor pool %q has no NOPs", cfg.ExecutorQualifier)
		}

		poolNOPs := shared.ConvertStringToNopAliases(pool.NOPAliases)
		for _, alias := range cfg.TargetNOPs {
			if !slices.Contains(poolNOPs, alias) {
				return fmt.Errorf("NOP alias %q not found in executor pool %q", alias, cfg.ExecutorQualifier)
			}
		}

		envSelectors := e.BlockChains.ListChainSelectors()
		for _, s := range cfg.ChainSelectors {
			if !slices.Contains(envSelectors, s) {
				return fmt.Errorf("selector %d is not available in environment", s)
			}
		}

		if shared.IsProductionEnvironment(e.Name) {
			if cfg.Topology.PyroscopeURL != "" {
				return fmt.Errorf("pyroscope URL is not supported for production environments")
			}
		}
		return nil
	}

	return deployment.CreateChangeSet(makeExecutorApply(ApplyExecutorConfigWithDeps), validate)
}

func ApplyExecutorConfigWithDeps(deps ExecutorApplyDeps, cfg ApplyExecutorConfigCfg) (deployment.ChangesetOutput, error) {
	deployedChains := getExecutorDeployedChains(deps.Env.DataStore, cfg.ExecutorQualifier)

	selectors := cfg.ChainSelectors
	if len(selectors) == 0 {
		selectors = deployedChains
	} else {
		selectors = filterChains(selectors, deployedChains)
	}

	pool := cfg.Topology.ExecutorPools[cfg.ExecutorQualifier]

	nopsToValidate := cfg.TargetNOPs
	if len(nopsToValidate) == 0 {
		nopsToValidate = shared.ConvertStringToNopAliases(pool.NOPAliases)
	}

	if err := validateExecutorChainSupport(deps, nopsToValidate, selectors); err != nil {
		return deployment.ChangesetOutput{}, err
	}

	executorPool := convertTopologyExecutorPool(pool)
	monitoring := convertTopologyMonitoring(&cfg.Topology.Monitoring)
	nopModes := buildNOPModes(cfg.Topology.NOPTopology.NOPs)

	input := sequences.GenerateExecutorConfigInput{
		ExecutorQualifier: cfg.ExecutorQualifier,
		ChainSelectors:    selectors,
		TargetNOPs:        cfg.TargetNOPs,
		ExecutorPool:      executorPool,
		IndexerAddress:    cfg.Topology.IndexerAddress,
		PyroscopeURL:      cfg.Topology.PyroscopeURL,
		Monitoring:        monitoring,
		NOPModes:          nopModes,
	}

	report, err := operations.ExecuteSequence(deps.Env.OperationsBundle, sequences.GenerateExecutorConfig, sequences.GenerateExecutorConfigDeps{Env: deps.Env}, input)
	if err != nil {
		return deployment.ChangesetOutput{
			Reports: report.ExecutionReports,
		}, fmt.Errorf("failed to generate executor config: %w", err)
	}

	manageReport, err := operations.ExecuteSequence(
		deps.Env.OperationsBundle,
		sequences.ManageJobProposals,
		sequences.ManageJobProposalsDeps{Env: deps.Env},
		sequences.ManageJobProposalsInput{
			JobSpecs:      report.Output.JobSpecs,
			AffectedScope: report.Output.AffectedScope,
			Labels: map[string]string{
				"job_type": "executor",
				"executor": cfg.ExecutorQualifier,
			},
			NOPs: sequences.NOPContext{
				Modes:      nopModes,
				TargetNOPs: cfg.TargetNOPs,
				AllNOPs:    getAllNOPAliases(cfg.Topology.NOPTopology.NOPs),
			},
		},
	)
	if err != nil {
		return deployment.ChangesetOutput{
			Reports: report.ExecutionReports,
		}, fmt.Errorf("failed to manage job proposals: %w", err)
	}

	deps.Env.Logger.Infow("Executor config applied",
		"jobsCount", len(manageReport.Output.Jobs),
		"revokedCount", len(manageReport.Output.RevokedJobs))

	return deployment.ChangesetOutput{
		Reports:   report.ExecutionReports,
		DataStore: manageReport.Output.DataStore,
	}, nil
}

func convertTopologyExecutorPool(pool topology.ExecutorPoolConfig) executorconfig.ExecutorPoolInput {
	return executorconfig.ExecutorPoolInput{
		NOPAliases:        shared.ConvertStringToNopAliases(pool.NOPAliases),
		ExecutionInterval: pool.ExecutionInterval,
		NtpServer:         pool.NtpServer,
		IndexerQueryLimit: pool.IndexerQueryLimit,
		BackoffDuration:   pool.BackoffDuration,
		LookbackWindow:    pool.LookbackWindow,
		ReaderCacheExpiry: pool.ReaderCacheExpiry,
		MaxRetryDuration:  pool.MaxRetryDuration,
		WorkerCount:       pool.WorkerCount,
	}
}

func validateExecutorChainSupport(
	deps ExecutorApplyDeps,
	nopsToValidate []shared.NOPAlias,
	selectors []uint64,
) error {
	if deps.JDClient == nil {
		deps.Env.Logger.Debugw("Offchain client not available, skipping chain support validation")
		return nil
	}

	nopAliasStrings := shared.ConvertNopAliasToString(nopsToValidate)
	supportedChains := fetchExecutorNodeChainSupport(deps, nopAliasStrings)
	if supportedChains == nil {
		return nil
	}

	var validationResults []shared.ChainValidationResult
	for _, nopAlias := range nopsToValidate {
		result := shared.ValidateNOPChainSupport(
			string(nopAlias),
			selectors,
			supportedChains[string(nopAlias)],
		)
		if result != nil {
			validationResults = append(validationResults, *result)
		}
	}

	return shared.FormatChainValidationError(validationResults)
}

func fetchExecutorNodeChainSupport(deps ExecutorApplyDeps, nopAliases []string) shared.ChainSupportByNOP {
	if deps.JDClient == nil {
		return nil
	}

	if len(nopAliases) == 0 {
		return nil
	}

	report, err := operations.ExecuteOperation(
		deps.Env.OperationsBundle,
		fetch_node_chain_support.FetchNodeChainSupport,
		fetch_node_chain_support.FetchNodeChainSupportDeps{
			JDClient: deps.JDClient,
			Logger:   deps.Env.Logger,
			NodeIDs:  deps.NodeIDs,
		},
		fetch_node_chain_support.FetchNodeChainSupportInput{
			NOPAliases: nopAliases,
		},
	)
	if err != nil {
		deps.Env.Logger.Warnw("Failed to fetch node chain support from JD", "error", err)
		return nil
	}

	return report.Output.SupportedChains
}
