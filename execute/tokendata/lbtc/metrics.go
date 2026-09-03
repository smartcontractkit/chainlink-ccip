package lbtc

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var PromLBTCBatchFetchCounter = promauto.NewCounterVec(
	prometheus.CounterOpts{
		Name: "ccip_exec_tokendata_lbtc_batch_fetch",
		Help: "Count of LBTC attestation batch fetches by outcome",
	},
	[]string{"outcome"},
)

const (
	lbtcOutcomeOK         = "ok"
	lbtcOutcomeHTTPError  = "http_error"
	lbtcOutcomeAPIError   = "api_error"
	lbtcOutcomeParseError = "parse_error"
)
