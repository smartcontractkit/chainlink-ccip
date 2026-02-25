package executor_config

import (
	"fmt"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/Masterminds/semver/v3"

	"github.com/smartcontractkit/chainlink-deployments-framework/operations"

	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/offchain"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/operations/shared"
)

// ExecutorPoolInput defines the configuration for an executor pool.
type ExecutorPoolInput struct {
	// NOPAliases is the list of NOP aliases that are members of this executor pool.
	NOPAliases []shared.NOPAlias
	// ExecutionInterval is the interval between execution cycles.
	ExecutionInterval time.Duration
	// NtpServer is the NTP server address for time synchronization (optional).
	NtpServer string
	// IndexerQueryLimit is the maximum number of records to fetch from the indexer per query.
	IndexerQueryLimit uint64
	// BackoffDuration is the duration to wait before retrying after a failure.
	BackoffDuration time.Duration
	// LookbackWindow is the time window for looking back at historical data.
	LookbackWindow time.Duration
	// ReaderCacheExpiry is the TTL for cached chain reader data.
	ReaderCacheExpiry time.Duration
	// MaxRetryDuration is the maximum duration to retry failed operations.
	MaxRetryDuration time.Duration
	// WorkerCount is the number of concurrent workers for processing executions.
	WorkerCount int
}

type BuildJobSpecsInput struct {
	GeneratedConfig   *ExecutorGeneratedConfig
	ExecutorQualifier string
	// TargetNOPs limits which NOPs will have their job specs updated. Defaults to all NOPs in the executor pool when empty.
	TargetNOPs     []shared.NOPAlias
	ExecutorPool   ExecutorPoolInput
	IndexerAddress []string
	PyroscopeURL   string
	Monitoring     shared.MonitoringInput
}

type BuildJobSpecsOutput struct {
	JobSpecs      shared.NOPJobSpecs
	AffectedScope shared.ExecutorJobScope
}

var BuildJobSpecs = operations.NewOperation(
	"build-executor-job-specs",
	semver.MustParse("1.0.0"),
	"Builds executor job specs from generated config and explicit input",
	func(b operations.Bundle, deps struct{}, input BuildJobSpecsInput) (BuildJobSpecsOutput, error) {
		jobSpecs := make(shared.NOPJobSpecs)
		scope := shared.ExecutorJobScope{
			ExecutorQualifier: input.ExecutorQualifier,
		}

		nopAliases := input.TargetNOPs
		if len(nopAliases) == 0 {
			nopAliases = input.ExecutorPool.NOPAliases
		}

		for _, nopAlias := range nopAliases {
			chainConfigs := make(map[string]offchain.ExecutorChainConfiguration)
			for chainSelectorStr, genCfg := range input.GeneratedConfig.ChainConfigs {
				chainConfigs[chainSelectorStr] = offchain.ExecutorChainConfiguration{
					OffRampAddress:         genCfg.OffRampAddress,
					RmnAddress:             genCfg.RmnAddress,
					DefaultExecutorAddress: genCfg.ExecutorProxyAddress,
					ExecutorPool:           shared.ConvertNopAliasToString(input.ExecutorPool.NOPAliases),
					ExecutionInterval:      input.ExecutorPool.ExecutionInterval,
				}
			}

			jobSpecID := shared.NewExecutorJobID(nopAlias, scope)

			executorCfg := offchain.ExecutorConfiguration{
				IndexerAddress:     input.IndexerAddress,
				ExecutorID:         jobSpecID.GetExecutorID(),
				PyroscopeURL:       input.PyroscopeURL,
				NtpServer:          input.ExecutorPool.NtpServer,
				IndexerQueryLimit:  input.ExecutorPool.IndexerQueryLimit,
				BackoffDuration:    input.ExecutorPool.BackoffDuration,
				LookbackWindow:     input.ExecutorPool.LookbackWindow,
				ReaderCacheExpiry:  input.ExecutorPool.ReaderCacheExpiry,
				MaxRetryDuration:   input.ExecutorPool.MaxRetryDuration,
				WorkerCount:        input.ExecutorPool.WorkerCount,
				Monitoring:         convertMonitoringInput(input.Monitoring),
				ChainConfiguration: chainConfigs,
			}

			configBytes, err := toml.Marshal(executorCfg)
			if err != nil {
				return BuildJobSpecsOutput{}, fmt.Errorf("failed to marshal executor config to TOML for NOP %q: %w", nopAlias, err)
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

		return BuildJobSpecsOutput{
			JobSpecs:      jobSpecs,
			AffectedScope: scope,
		}, nil
	},
)

func convertMonitoringInput(cfg shared.MonitoringInput) offchain.MonitoringConfig {
	return offchain.MonitoringConfig{
		Enabled: cfg.Enabled,
		Type:    cfg.Type,
		Beholder: offchain.BeholderConfig{
			InsecureConnection:       cfg.Beholder.InsecureConnection,
			CACertFile:               cfg.Beholder.CACertFile,
			OtelExporterGRPCEndpoint: cfg.Beholder.OtelExporterGRPCEndpoint,
			OtelExporterHTTPEndpoint: cfg.Beholder.OtelExporterHTTPEndpoint,
			LogStreamingEnabled:      cfg.Beholder.LogStreamingEnabled,
			MetricReaderInterval:     cfg.Beholder.MetricReaderInterval,
			TraceSampleRatio:         cfg.Beholder.TraceSampleRatio,
			TraceBatchTimeout:        cfg.Beholder.TraceBatchTimeout,
		},
	}
}
