package observer

import (
	"context"

	"go.opentelemetry.io/otel/attribute"

	"github.com/smartcontractkit/chainlink-ccip/execute/tokendata/metricsutil"
)

var (
	observerErrorCounter = metricsutil.NewLazyCounter("ccip_exec_tokendata_observer_error")
	backgroundQueueDepth = metricsutil.NewLazyGauge("ccip_exec_tokendata_background_queue_depth")
	backgroundCacheSize  = metricsutil.NewLazyGauge("ccip_exec_tokendata_background_cache_size")
	backgroundFailure    = metricsutil.NewLazyCounter("ccip_exec_tokendata_background_observe_failure")
)

// trackObserverError counts an error swallowed by the composite token data observer.
func trackObserverError(observer, stage string) {
	observerErrorCounter.Add(context.Background(), 1,
		attribute.String("observer", observer),
		attribute.String("stage", stage))
}

// trackBackgroundQueueDepth reports the number of messages waiting in the background
// observer queue.
func trackBackgroundQueueDepth(observer string, depth int) {
	backgroundQueueDepth.Record(context.Background(), int64(depth), attribute.String("observer", observer))
}

// trackBackgroundCacheSize reports the number of token data entries cached by the
// background observer.
func trackBackgroundCacheSize(observer string, size int) {
	backgroundCacheSize.Record(context.Background(), int64(size), attribute.String("observer", observer))
}

// trackBackgroundObserveFailure counts a background worker observe failure by reason.
func trackBackgroundObserveFailure(observer, reason string) {
	backgroundFailure.Add(context.Background(), 1,
		attribute.String("observer", observer),
		attribute.String("reason", reason))
}
