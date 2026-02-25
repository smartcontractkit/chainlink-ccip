package offchain

import (
	"fmt"
)

// MonitoringConfig provides monitoring configuration shared across executor and verifier job specs.
type MonitoringConfig struct {
	Enabled  bool           `toml:"Enabled"`
	Type     string         `toml:"Type"`
	Beholder BeholderConfig `toml:"Beholder"`
}

// BeholderConfig wraps OpenTelemetry configuration for the beholder client.
type BeholderConfig struct {
	InsecureConnection       bool    `toml:"InsecureConnection"`
	CACertFile               string  `toml:"CACertFile"`
	OtelExporterGRPCEndpoint string  `toml:"OtelExporterGRPCEndpoint"`
	OtelExporterHTTPEndpoint string  `toml:"OtelExporterHTTPEndpoint"`
	LogStreamingEnabled      bool    `toml:"LogStreamingEnabled"`
	MetricReaderInterval     int64   `toml:"MetricReaderInterval"`
	TraceSampleRatio         float64 `toml:"TraceSampleRatio"`
	TraceBatchTimeout        int64   `toml:"TraceBatchTimeout"`
}

// Validate performs validation on the monitoring configuration.
func (m *MonitoringConfig) Validate() error {
	if m.Enabled && m.Type == "" {
		return fmt.Errorf("monitoring type is required when monitoring is enabled")
	}

	if m.Enabled && m.Type == "beholder" {
		if err := m.Beholder.Validate(); err != nil {
			return fmt.Errorf("beholder config validation failed: %w", err)
		}
	}

	return nil
}

// Validate performs validation on the beholder configuration.
func (b *BeholderConfig) Validate() error {
	if b.MetricReaderInterval <= 0 {
		return fmt.Errorf("metric_reader_interval must be positive, got %d", b.MetricReaderInterval)
	}

	if b.TraceSampleRatio < 0 || b.TraceSampleRatio > 1 {
		return fmt.Errorf("trace_sample_ratio must be between 0 and 1, got %f", b.TraceSampleRatio)
	}

	if b.TraceBatchTimeout <= 0 {
		return fmt.Errorf("trace_batch_timeout must be positive, got %d", b.TraceBatchTimeout)
	}

	return nil
}
