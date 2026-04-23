package offchain

import "time"

// ExecutorConfiguration mirrors chainlink-ccv/executor.Configuration for TOML job spec serialization.
type ExecutorConfiguration struct {
	IndexerAddress     []string                          `toml:"indexer_address"`
	BackoffDuration    time.Duration                     `toml:"source_backoff_duration"`
	LookbackWindow     time.Duration                     `toml:"startup_lookback_window"`
	IndexerQueryLimit  uint64                            `toml:"indexer_query_limit"`
	PyroscopeURL       string                            `toml:"pyroscope_url"`
	ExecutorID         string                            `toml:"executor_id"`
	Monitoring         ExecutorMonitoringConfig          `toml:"Monitoring"`
	ReaderCacheExpiry  time.Duration                     `toml:"reader_cache_expiry"`
	MaxRetryDuration   time.Duration                     `toml:"max_retry_duration"`
	NtpServer          string                            `toml:"ntp_server"`
	ChainConfiguration map[string]ExecutorChainCfg       `toml:"chain_configuration"`
	WorkerCount        int                               `toml:"worker_count"`
}

// ExecutorChainCfg mirrors chainlink-ccv/executor.ChainConfiguration for TOML serialization.
type ExecutorChainCfg struct {
	RmnAddress             string        `toml:"rmn_address"`
	OffRampAddress         string        `toml:"off_ramp_address"`
	ExecutorPool           []string      `toml:"executor_pool"`
	ExecutionInterval      time.Duration `toml:"execution_interval"`
	DefaultExecutorAddress string        `toml:"default_executor_address"`
}

type ExecutorMonitoringConfig struct {
	Enabled  bool                    `toml:"Enabled"`
	Type     string                  `toml:"Type"`
	Beholder ExecutorBeholderConfig  `toml:"Beholder"`
}

type ExecutorBeholderConfig struct {
	InsecureConnection       bool    `toml:"InsecureConnection"`
	CACertFile               string  `toml:"CACertFile"`
	OtelExporterGRPCEndpoint string  `toml:"OtelExporterGRPCEndpoint"`
	OtelExporterHTTPEndpoint string  `toml:"OtelExporterHTTPEndpoint"`
	LogStreamingEnabled      bool    `toml:"LogStreamingEnabled"`
	MetricReaderInterval     int64   `toml:"MetricReaderInterval"`
	TraceSampleRatio         float64 `toml:"TraceSampleRatio"`
	TraceBatchTimeout        int64   `toml:"TraceBatchTimeout"`
}
