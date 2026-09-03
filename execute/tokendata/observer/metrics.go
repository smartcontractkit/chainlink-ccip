package observer

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	PromObserverErrorCounter = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "ccip_exec_tokendata_observer_error",
			Help: "Count of errors swallowed by the composite token data observer",
		},
		[]string{"observer", "stage"},
	)

	PromBackgroundQueueDepthGauge = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "ccip_exec_tokendata_background_queue_depth",
			Help: "Number of messages waiting in the background observer queue",
		},
		[]string{"observer"},
	)

	PromBackgroundCacheSizeGauge = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "ccip_exec_tokendata_background_cache_size",
			Help: "Number of token data entries cached by the background observer",
		},
		[]string{"observer"},
	)

	PromBackgroundObserveFailureCounter = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "ccip_exec_tokendata_background_observe_failure",
			Help: "Count of background worker observe failures by reason",
		},
		[]string{"observer", "reason"},
	)
)
