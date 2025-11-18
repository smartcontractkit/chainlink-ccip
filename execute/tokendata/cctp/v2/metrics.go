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
		[]string{"destChainFamily", "destChainID", "sourceChainFamily", "sourceChainID", "sourceDomain", "success", "httpStatus"},
	)
)

// MetricsReporter provides metrics reporting functionality for CCTP v2 observer
type MetricsReporter interface {
	TrackAttestationAPILatency(
		sourceChain cciptypes.ChainSelector, sourceDomain uint32, success bool, httpStatus string, latency time.Duration)
}

// MetricsReporterImpl implements MetricsReporter with actual Prometheus metrics
type MetricsReporterImpl struct {
	lggr            logger.Logger
	destChainFamily string
	destChainID     string
}

// noOpMetricsReporter implements MetricsReporter with no-op methods
type noOpMetricsReporter struct{}

func (n *noOpMetricsReporter) TrackAttestationAPILatency(cciptypes.ChainSelector, uint32, bool, string, time.Duration) {
}

// NewMetricsReporter creates a new metrics reporter for CCTP v2
func NewMetricsReporter(lggr logger.Logger, destChainSelector cciptypes.ChainSelector) (MetricsReporter, error) {
	chainFamily, chainID, ok := libs.GetChainInfoFromSelector(destChainSelector)
	if !ok {
		return &noOpMetricsReporter{}, fmt.Errorf("chainFamily and chainID not found for selector %d",
			destChainSelector)
	}

	return &MetricsReporterImpl{
		lggr:            lggr,
		destChainFamily: chainFamily,
		destChainID:     chainID,
	}, nil
}

// NewNoOpMetricsReporter creates a no-op metrics reporter
func NewNoOpMetricsReporter() MetricsReporter {
	return &noOpMetricsReporter{}
}

// TrackAttestationAPILatency tracks the latency of attestation API calls
func (r *MetricsReporterImpl) TrackAttestationAPILatency(
	sourceChain cciptypes.ChainSelector, sourceDomain uint32, success bool, httpStatus string, latency time.Duration,
) {
	sourceChainFamily, sourceChainID, _ := libs.GetChainInfoFromSelector(sourceChain)
	PromCCTPv2AttestationAPILatencyHistogram.
		WithLabelValues(r.destChainFamily, r.destChainID, sourceChainFamily, sourceChainID,
			strconv.FormatUint(uint64(sourceDomain), 10), strconv.FormatBool(success), httpStatus).
		Observe(latency.Seconds())
}
