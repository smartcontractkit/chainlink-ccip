package offchain

import "time"

type TokenVerifierGeneratedConfig struct {
	PyroscopeURL       string                        `json:"pyroscope_url"`
	OnRampAddresses    map[string]string             `json:"on_ramp_addresses"`
	RMNRemoteAddresses map[string]string             `json:"rmn_remote_addresses"`
	TokenVerifiers     []TokenVerifierEntry          `json:"token_verifiers"`
	Monitoring         TokenVerifierMonitoringConfig `json:"monitoring"`
}

type TokenVerifierEntry struct {
	VerifierID string                  `json:"verifier_id"`
	Type       string                  `json:"type"`
	Version    string                  `json:"version"`
	CCTP       *CCTPVerifierConfig     `json:"cctp,omitempty"`
	Lombard    *LombardVerifierConfig  `json:"lombard,omitempty"`
}

type CCTPVerifierConfig struct {
	AttestationAPI         string            `json:"attestation_api"`
	AttestationAPITimeout  time.Duration     `json:"attestation_api_timeout"`
	AttestationAPIInterval time.Duration     `json:"attestation_api_interval"`
	AttestationAPICooldown time.Duration     `json:"attestation_api_cooldown"`
	VerifierVersion        []byte            `json:"verifier_version"`
	Verifiers              map[string]string `json:"verifiers"`
	VerifierResolvers      map[string]string `json:"verifier_resolvers"`
}

type LombardVerifierConfig struct {
	AttestationAPI          string            `json:"attestation_api"`
	AttestationAPITimeout   time.Duration     `json:"attestation_api_timeout"`
	AttestationAPIInterval  time.Duration     `json:"attestation_api_interval"`
	AttestationAPIBatchSize int               `json:"attestation_api_batch_size"`
	VerifierVersion         []byte            `json:"verifier_version"`
	VerifierResolvers       map[string]string `json:"verifier_resolvers"`
}

type TokenVerifierMonitoringConfig struct {
	Enabled  bool                          `json:"enabled"`
	Type     string                        `json:"type"`
	Beholder TokenVerifierBeholderConfig   `json:"beholder"`
}

type TokenVerifierBeholderConfig struct {
	InsecureConnection       bool    `json:"insecure_connection"`
	CACertFile               string  `json:"ca_cert_file"`
	OtelExporterGRPCEndpoint string  `json:"otel_exporter_grpc_endpoint"`
	OtelExporterHTTPEndpoint string  `json:"otel_exporter_http_endpoint"`
	LogStreamingEnabled      bool    `json:"log_streaming_enabled"`
	MetricReaderInterval     int64   `json:"metric_reader_interval"`
	TraceSampleRatio         float64 `json:"trace_sample_ratio"`
	TraceBatchTimeout        int64   `json:"trace_batch_timeout"`
}
