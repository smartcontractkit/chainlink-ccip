package v2

import (
	"fmt"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-ccip/internal/libs"
)

var (
	// Observation latency histogram
	PromCCTPv2ObservationLatencyHistogram = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "ccip_exec_tokendata_cctpv2_observation_latency",
			Help: "Latency of CCTP v2 token data observation operations",
			Buckets: []float64{
				float64(50 * time.Millisecond),
				float64(100 * time.Millisecond),
				float64(200 * time.Millisecond),
				float64(500 * time.Millisecond),
				float64(700 * time.Millisecond),
				float64(time.Second),
				float64(2 * time.Second),
				float64(5 * time.Second),
				float64(7 * time.Second),
				float64(10 * time.Second),
				float64(20 * time.Second),
			},
		},
		[]string{"destChainFamily", "destChainID", "sourceChainFamily", "sourceChainID", "method"},
	)

	// Attestation API call latency
	PromCCTPv2AttestationAPILatencyHistogram = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "ccip_exec_tokendata_cctpv2_attestation_api_latency",
			Help: "Latency of CCTP v2 attestation API calls",
			Buckets: []float64{
				float64(50 * time.Millisecond),
				float64(100 * time.Millisecond),
				float64(200 * time.Millisecond),
				float64(500 * time.Millisecond),
				float64(time.Second),
				float64(2 * time.Second),
				float64(5 * time.Second),
				float64(10 * time.Second),
			},
		},
		[]string{"destChainFamily", "destChainID", "sourceChainFamily", "sourceChainID", "sourceDomain", "status"},
	)

	// Token processing counters
	PromCCTPv2TokenProcessingCounter = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "ccip_exec_tokendata_cctpv2_tokens_processed_total",
			Help: "Total number of CCTP v2 tokens processed by status",
		},
		[]string{"destChainFamily", "destChainID", "sourceChainFamily", "sourceChainID", "status"},
	)

	// Message matching counters
	PromCCTPv2MessageMatchingCounter = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "ccip_exec_tokendata_cctpv2_message_matching_total",
			Help: "Total number of CCTP v2 message matching attempts by result",
		},
		[]string{"destChainFamily", "destChainID", "sourceChainFamily", "sourceChainID", "result"},
	)
)

// MetricsReporter provides metrics reporting functionality for CCTP v2 observer
type MetricsReporter interface {
	TrackObservationLatency(sourceChain cciptypes.ChainSelector, method string, latency time.Duration)
	TrackAttestationAPILatency(
		sourceChain cciptypes.ChainSelector, sourceDomain uint32, status string, latency time.Duration)
	TrackTokenProcessed(sourceChain cciptypes.ChainSelector, status string, count int)
	TrackMessageMatching(sourceChain cciptypes.ChainSelector, result string, count int)
}

// metricsReporter implements MetricsReporter with actual Prometheus metrics
type metricsReporter struct {
	lggr            logger.Logger
	destChainFamily string
	destChainID     string
}

// noOpMetricsReporter implements MetricsReporter with no-op methods
type noOpMetricsReporter struct{}

func (n *noOpMetricsReporter) TrackObservationLatency(cciptypes.ChainSelector, string, time.Duration) {
}
func (n *noOpMetricsReporter) TrackAttestationAPILatency(cciptypes.ChainSelector, uint32, string, time.Duration) {
}
func (n *noOpMetricsReporter) TrackTokenProcessed(cciptypes.ChainSelector, string, int)  {}
func (n *noOpMetricsReporter) TrackMessageMatching(cciptypes.ChainSelector, string, int) {}

// NewMetricsReporter creates a new metrics reporter for CCTP v2
func NewMetricsReporter(lggr logger.Logger, destChainSelector cciptypes.ChainSelector) (MetricsReporter, error) {
	chainFamily, chainID, ok := libs.GetChainInfoFromSelector(destChainSelector)
	if !ok {
		return &noOpMetricsReporter{}, fmt.Errorf("chainFamily and chainID not found for selector %d",
			destChainSelector)
	}

	return &metricsReporter{
		lggr:            lggr,
		destChainFamily: chainFamily,
		destChainID:     chainID,
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
	sourceChainFamily, sourceChainID, _ := libs.GetChainInfoFromSelector(sourceChain)
	PromCCTPv2ObservationLatencyHistogram.
		WithLabelValues(r.destChainFamily, r.destChainID, sourceChainFamily, sourceChainID, method).
		Observe(latency.Seconds())
}

// TrackAttestationAPILatency tracks the latency of attestation API calls
func (r *metricsReporter) TrackAttestationAPILatency(
	sourceChain cciptypes.ChainSelector, sourceDomain uint32, status string, latency time.Duration,
) {
	sourceChainFamily, sourceChainID, _ := libs.GetChainInfoFromSelector(sourceChain)
	PromCCTPv2AttestationAPILatencyHistogram.
		WithLabelValues(r.destChainFamily, r.destChainID, sourceChainFamily, sourceChainID,
			strconv.FormatUint(uint64(sourceDomain), 10), status).
		Observe(latency.Seconds())
}

// TrackTokenProcessed tracks tokens processed by status
func (r *metricsReporter) TrackTokenProcessed(sourceChain cciptypes.ChainSelector, status string, count int) {
	sourceChainFamily, sourceChainID, _ := libs.GetChainInfoFromSelector(sourceChain)
	PromCCTPv2TokenProcessingCounter.
		WithLabelValues(r.destChainFamily, r.destChainID, sourceChainFamily, sourceChainID, status).
		Add(float64(count))
}

// TrackMessageMatching tracks message matching attempts
func (r *metricsReporter) TrackMessageMatching(sourceChain cciptypes.ChainSelector, result string, count int) {
	sourceChainFamily, sourceChainID, _ := libs.GetChainInfoFromSelector(sourceChain)
	PromCCTPv2MessageMatchingCounter.
		WithLabelValues(r.destChainFamily, r.destChainID, sourceChainFamily, sourceChainID, result).
		Add(float64(count))
}
