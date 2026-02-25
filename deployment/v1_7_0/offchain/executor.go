package offchain

import (
	"fmt"
	"slices"
	"time"
)

const (
	BackoffDurationDefault   = 15 * time.Second
	LookbackWindowDefault    = 1 * time.Hour
	ReaderCacheExpiryDefault = 5 * time.Minute
	MaxRetryDurationDefault  = 8 * time.Hour
	ExecutionIntervalDefault = 1 * time.Minute
	NtpServerDefault         = "time.google.com"
	WorkerCountDefault       = 100
	IndexerQueryLimitDefault = 100
)

// ExecutorConfiguration is the complete set of information an executor needs to operate normally.
// Used for job spec TOML; BurntSushi toml parses duration from strings.
type ExecutorConfiguration struct {
	IndexerAddress     []string                        `toml:"indexer_address"`
	BackoffDuration    time.Duration                   `toml:"source_backoff_duration"`
	LookbackWindow     time.Duration                   `toml:"startup_lookback_window"`
	IndexerQueryLimit  uint64                          `toml:"indexer_query_limit"`
	PyroscopeURL       string                          `toml:"pyroscope_url"`
	ExecutorID         string                          `toml:"executor_id"`
	Monitoring         MonitoringConfig                `toml:"Monitoring"`
	ReaderCacheExpiry  time.Duration                   `toml:"reader_cache_expiry"`
	MaxRetryDuration   time.Duration                   `toml:"max_retry_duration"`
	NtpServer          string                          `toml:"ntp_server"`
	ChainConfiguration map[string]ExecutorChainConfiguration `toml:"chain_configuration"`
	WorkerCount        int                             `toml:"worker_count"`
}

// ExecutorChainConfiguration is all the configuration an executor needs for a specific chain.
type ExecutorChainConfiguration struct {
	RmnAddress             string        `toml:"rmn_address"`
	OffRampAddress         string        `toml:"off_ramp_address"`
	ExecutorPool           []string      `toml:"executor_pool"`
	ExecutionInterval      time.Duration `toml:"execution_interval"`
	DefaultExecutorAddress string        `toml:"default_executor_address"`
}

// Validate validates the executor configuration.
func (c *ExecutorConfiguration) Validate() error {
	if c.ExecutorID == "" {
		return fmt.Errorf("this_executor_id must be configured")
	}

	if len(c.ChainConfiguration) == 0 {
		return fmt.Errorf("at least one chain must be configured")
	}

	if len(c.IndexerAddress) < 1 {
		return fmt.Errorf("at least one indexer address must be configured")
	}

	for chainSel, chainConfig := range c.ChainConfiguration {
		if len(chainConfig.ExecutorPool) == 0 {
			return fmt.Errorf("executor_pool must be configured for chain %s", chainSel)
		}
		if !slices.Contains(chainConfig.ExecutorPool, c.ExecutorID) {
			return fmt.Errorf("this_executor_id '%s' not found in executor_pool for chain %s", c.ExecutorID, chainSel)
		}
	}

	return nil
}

// GetNormalizedConfig validates the configuration and applies defaults.
func (c *ExecutorConfiguration) GetNormalizedConfig() (*ExecutorConfiguration, error) {
	normalized := *c
	normalized.ChainConfiguration = make(map[string]ExecutorChainConfiguration, len(c.ChainConfiguration))
	for k, v := range c.ChainConfiguration {
		normalized.ChainConfiguration[k] = v
	}

	if err := normalized.Validate(); err != nil {
		return nil, err
	}

	if c.NtpServer == "" {
		normalized.NtpServer = NtpServerDefault
	}

	parseOrDefault := func(raw, defaultVal time.Duration) time.Duration {
		if raw == 0 {
			return defaultVal
		}
		return raw
	}

	normalized.BackoffDuration = parseOrDefault(c.BackoffDuration, BackoffDurationDefault)
	normalized.LookbackWindow = parseOrDefault(c.LookbackWindow, LookbackWindowDefault)
	normalized.ReaderCacheExpiry = parseOrDefault(c.ReaderCacheExpiry, ReaderCacheExpiryDefault)
	normalized.MaxRetryDuration = parseOrDefault(c.MaxRetryDuration, MaxRetryDurationDefault)
	if c.IndexerQueryLimit == 0 {
		normalized.IndexerQueryLimit = IndexerQueryLimitDefault
	}
	if c.WorkerCount == 0 {
		normalized.WorkerCount = WorkerCountDefault
	}

	for chainSel, updatedChain := range normalized.ChainConfiguration {
		updatedChain.ExecutionInterval = parseOrDefault(updatedChain.ExecutionInterval, ExecutionIntervalDefault)
		normalized.ChainConfiguration[chainSel] = updatedChain
	}

	return &normalized, nil
}
