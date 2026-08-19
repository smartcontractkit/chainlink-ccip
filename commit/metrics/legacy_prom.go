package metrics

import (
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// TODO: consider removing these metrics entirely in a follow-up.
var (
	rmnLatencyBucketsMilliseconds = []float64{
		5, 10, 25, 50, 100, 250, 500, 1000, 2500, 5000, 10000,
	}
	promProcessorOutputCounter = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "ccip_commit_processor_output_sizes",
			Help: "This metric tracks the number of different items in the commit processor",
		},
		[]string{"chainFamily", "chainID", "processor", "method", "type"},
	)
	promProcessorLatencyHistogram = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "ccip_commit_processor_latency",
			Help: "This metric tracks the client-observed latency of a single processor method",
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
		[]string{"chainFamily", "chainID", "processor", "method"},
	)
	promProcessorErrors = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "ccip_commit_processor_errors",
			Help: "This metric tracks the number of errors in the commit processor observation",
		},
		[]string{"chainFamily", "chainID", "processor", "method"},
	)
	promSequenceNumbers = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "ccip_commit_max_sequence_number",
			Help: "This metric tracks the max sequence number observed by the commit processor",
		},
		[]string{"chainFamily", "chainID", "sourceChainFamily", "sourceChain",
			"method", "source_network_name", "dest_network_name"},
	)
	promMerkleProcessorRmnReportLatency = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "ccip_commit_merkle_processor_rmn_report_latency_ms",
			Help:    "This metric tracks the client-observed latency of building an full RMN report with signatures",
			Buckets: rmnLatencyBucketsMilliseconds,
		},
		[]string{"chainID", "success"},
	)
	promRmnControllerRmnRequestLatency = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "ccip_commit_rmn_controller_rmn_request_latency_ms",
			Help:    "This metric tracks the client-observed latency of a single RMN request",
			Buckets: rmnLatencyBucketsMilliseconds,
		},
		[]string{"method", "nodeID", "error"},
	)
	promCommitLatestRoundID = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "ccip_commit_latest_round_id",
		Help: "The latest round ID observed by the commit plugin",
	}, []string{"source_network_name", "dest_network_name", "contract_address", "plugin"})
	promLooppCCIPProviderSupported = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "ccip_commit_loopp_ccip_provider_supported",
		Help: "Tracks whether LOOPP CCIP provider is supported for each chain family (1 = supported, 0 = not supported)",
	}, []string{"chain_family"})
	promCommitConfigDigestMismatch = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "ccip_commit_config_digest_mismatch",
		Help: "Reports whether the home chain config digest differs from the offramp config digest (1 = mismatch, 0 = match)",
	}, []string{"chain_family", "chain_id"})
)

// Deprecated - RMN is perma-disabled, should be removed in a follow-up.
func (p *PromReporter) TrackRmnReport(latency float64, success bool) {
	successStr := strconv.FormatBool(success)
	p.merkleProcessorRmnReportHistogram.WithLabelValues(p.chainID, successStr).Observe(latency)
}

// Deprecated - RMN is perma-disabled, should be removed in a follow-up.
func (p *PromReporter) TrackRmnRequest(method string, latency float64, nodeID uint64, err string) {
	nodeIDStr := strconv.FormatUint(nodeID, 10)
	p.rmnControllerRmnRequestHistogram.WithLabelValues(method, nodeIDStr, err).Observe(latency)
}
