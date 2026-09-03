package http

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	PromTokendataHTTPRequestCounter = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "ccip_exec_tokendata_http_request",
			Help: "Count of tokendata attestation HTTP requests by API and outcome",
		},
		[]string{"api", "outcome"},
	)

	PromTokendataHTTPCooldownActiveGauge = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "ccip_exec_tokendata_http_cooldown_active",
			Help: "Whether the tokendata attestation HTTP client is in a rate-limit cooldown period",
		},
		[]string{"api"},
	)
)

const (
	outcomeOK          = "ok"
	outcomeNotReady404 = "not_ready_404"
	outcomeRateLimited = "rate_limited"
	outcomeTimeout     = "timeout"
	outcomeHTTPError   = "http_error"
	outcomeUnknown     = "unknown_status"
	outcomeParseError  = "parse_error"
)
