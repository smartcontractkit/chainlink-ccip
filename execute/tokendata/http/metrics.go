package http

import (
	"context"

	"go.opentelemetry.io/otel/attribute"

	"github.com/smartcontractkit/chainlink-ccip/execute/tokendata/metricsutil"
)

var (
	httpRequestCounter  = metricsutil.NewLazyCounter("ccip_exec_tokendata_http_request")
	cooldownActiveGauge = metricsutil.NewLazyGauge("ccip_exec_tokendata_http_cooldown_active")
)

// trackRequest counts one attestation-API request outcome for the given API.
func trackRequest(ctx context.Context, api, outcome string) {
	httpRequestCounter.Add(ctx, 1,
		attribute.String("api", api),
		attribute.String("outcome", outcome))
}

// trackCooldownActive reports whether the client is in a rate-limit cooldown period.
func trackCooldownActive(ctx context.Context, api string, active bool) {
	var val int64
	if active {
		val = 1
	}
	cooldownActiveGauge.Record(ctx, val, attribute.String("api", api))
}

const (
	outcomeOK          = "ok"
	outcomeNotReady404 = "not_ready_404"
	outcomeRateLimited = "rate_limited"
	outcomeTimeout     = "timeout"
	outcomeHTTPError   = "http_error"
	outcomeUnknown     = "unknown_status"
	outcomeParseError  = "parse_error"
)
