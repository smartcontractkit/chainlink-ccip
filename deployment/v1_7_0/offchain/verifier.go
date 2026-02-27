package offchain

// VerifierJobConfig mirrors chainlink-ccv/verifier/commit.Config for TOML job spec serialization.
type VerifierJobConfig struct {
	VerifierID                     string            `toml:"verifier_id"`
	AggregatorAddress              string            `toml:"aggregator_address"`
	InsecureAggregatorConnection   bool              `toml:"insecure_aggregator_connection"`
	SignerAddress                  string            `toml:"signer_address"`
	PyroscopeURL                   string            `toml:"pyroscope_url"`
	CommitteeVerifierAddresses     map[string]string `toml:"committee_verifier_addresses"`
	OnRampAddresses                map[string]string `toml:"on_ramp_addresses"`
	DefaultExecutorOnRampAddresses map[string]string `toml:"default_executor_on_ramp_addresses"`
	RMNRemoteAddresses             map[string]string `toml:"rmn_remote_addresses"`
	DisableFinalityCheckers        []string          `toml:"disable_finality_checkers"`
	Monitoring                     VerifierMonitoringConfig `toml:"monitoring"`
}

// VerifierMonitoringConfig mirrors chainlink-ccv/verifier.MonitoringConfig for TOML serialization.
type VerifierMonitoringConfig struct {
	Enabled  bool                   `toml:"Enabled"`
	Type     string                 `toml:"Type"`
	Beholder VerifierBeholderConfig `toml:"Beholder"`
}

// VerifierBeholderConfig mirrors chainlink-ccv/verifier.BeholderConfig for TOML serialization.
type VerifierBeholderConfig struct {
	InsecureConnection       bool    `toml:"InsecureConnection"`
	CACertFile               string  `toml:"CACertFile"`
	OtelExporterGRPCEndpoint string  `toml:"OtelExporterGRPCEndpoint"`
	OtelExporterHTTPEndpoint string  `toml:"OtelExporterHTTPEndpoint"`
	LogStreamingEnabled      bool    `toml:"LogStreamingEnabled"`
	MetricReaderInterval     int64   `toml:"MetricReaderInterval"`
	TraceSampleRatio         float64 `toml:"TraceSampleRatio"`
	TraceBatchTimeout        int64   `toml:"TraceBatchTimeout"`
}
