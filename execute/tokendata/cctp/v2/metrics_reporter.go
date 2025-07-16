package v2

import (
	"fmt"
	"strconv"
	"time"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"

	"github.com/smartcontractkit/chainlink-ccip/internal/libs"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
)

// MetricsReporter provides metrics reporting functionality for CCTP v2 observer
type MetricsReporter interface {
	TrackObservationLatency(sourceChain cciptypes.ChainSelector, method string, latency time.Duration)
	TrackAttestationAPILatency(sourceDomain uint32, status string, latency time.Duration)
	TrackTokenProcessed(sourceChain cciptypes.ChainSelector, status string, count int)
	TrackMessageMatching(sourceChain cciptypes.ChainSelector, result string, count int)
}

// metricsReporter implements MetricsReporter with actual Prometheus metrics
type metricsReporter struct {
	lggr        logger.Logger
	chainFamily string
	chainID     string
}

// noOpMetricsReporter implements MetricsReporter with no-op methods
type noOpMetricsReporter struct{}

func (n *noOpMetricsReporter) TrackObservationLatency(cciptypes.ChainSelector, string, time.Duration) {
}
func (n *noOpMetricsReporter) TrackAttestationAPILatency(uint32, string, time.Duration)  {}
func (n *noOpMetricsReporter) TrackTokenProcessed(cciptypes.ChainSelector, string, int)  {}
func (n *noOpMetricsReporter) TrackMessageMatching(cciptypes.ChainSelector, string, int) {}

// NewMetricsReporter creates a new metrics reporter for CCTP v2
func NewMetricsReporter(lggr logger.Logger, destChainSelector cciptypes.ChainSelector) (MetricsReporter, error) {
	chainFamily, chainID, ok := libs.GetChainInfoFromSelector(destChainSelector)
	if !ok {
		return &noOpMetricsReporter{}, fmt.Errorf("chainFamily and chainID not found for selector %d", destChainSelector)
	}

	return &metricsReporter{
		lggr:        lggr,
		chainFamily: chainFamily,
		chainID:     chainID,
	}, nil
}

// NewNoOpMetricsReporter creates a no-op metrics reporter
func NewNoOpMetricsReporter() MetricsReporter {
	return &noOpMetricsReporter{}
}

// TrackObservationLatency tracks the latency of observation operations
func (r *metricsReporter) TrackObservationLatency(
	sourceChain cciptypes.ChainSelector, method string, latency time.Duration,
) {
	PromCCTPv2ObservationLatencyHistogram.
		WithLabelValues(r.chainFamily, r.chainID, strconv.FormatUint(uint64(sourceChain), 10), method).
		Observe(latency.Seconds())
}

// TrackAttestationAPILatency tracks the latency of attestation API calls
func (r *metricsReporter) TrackAttestationAPILatency(sourceDomain uint32, status string, latency time.Duration) {
	PromCCTPv2AttestationAPILatencyHistogram.
		WithLabelValues(r.chainFamily, r.chainID, strconv.FormatUint(uint64(sourceDomain), 10), status).
		Observe(latency.Seconds())
}

// TrackTokenProcessed tracks tokens processed by status
func (r *metricsReporter) TrackTokenProcessed(sourceChain cciptypes.ChainSelector, status string, count int) {
	PromCCTPv2TokenProcessingCounter.
		WithLabelValues(r.chainFamily, r.chainID, strconv.FormatUint(uint64(sourceChain), 10), status).
		Add(float64(count))
}

// TrackMessageMatching tracks message matching attempts
func (r *metricsReporter) TrackMessageMatching(sourceChain cciptypes.ChainSelector, result string, count int) {
	PromCCTPv2MessageMatchingCounter.
		WithLabelValues(r.chainFamily, r.chainID, strconv.FormatUint(uint64(sourceChain), 10), result).
		Add(float64(count))
}
