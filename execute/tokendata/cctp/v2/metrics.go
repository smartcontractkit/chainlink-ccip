package v2

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// CCTP v2 specific metrics following the repository patterns
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
		[]string{"destChainFamily", "destChainID", "sourceChain", "method"},
	)

	// Attestation API call latency
	PromCCTPv2AttestationAPILatencyHistogram = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "ccip_exec_tokendata_cctpv2_attestation_api_latency",
			Help: "Latency of CCTP v2 attestation API calls",
			Buckets: []float64{
				float64(100 * time.Millisecond),
				float64(200 * time.Millisecond),
				float64(500 * time.Millisecond),
				float64(time.Second),
				float64(2 * time.Second),
				float64(5 * time.Second),
				float64(10 * time.Second),
				float64(20 * time.Second),
				float64(30 * time.Second),
				float64(60 * time.Second),
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
		[]string{"destChainFamily", "destChainID", "sourceChain", "status"},
	)

	// Message matching counters
	PromCCTPv2MessageMatchingCounter = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "ccip_exec_tokendata_cctpv2_message_matching_total",
			Help: "Total number of CCTP v2 message matching attempts by result",
		},
		[]string{"destChainFamily", "destChainID", "sourceChain", "result"},
	)
)
