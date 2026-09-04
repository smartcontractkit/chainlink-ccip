package lbtc

import (
	"context"

	"go.opentelemetry.io/otel/attribute"

	"github.com/smartcontractkit/chainlink-ccip/execute/tokendata/metricsutil"
)

var batchFetchCounter = metricsutil.NewLazyCounter("ccip_exec_tokendata_lbtc_batch_fetch")

// trackBatchFetch counts one LBTC attestation batch fetch by outcome.
func trackBatchFetch(outcome string) {
	batchFetchCounter.Add(context.Background(), 1, attribute.String("outcome", outcome))
}

const (
	lbtcOutcomeOK        = "ok"
	lbtcOutcomeHTTPError = "http_error"
	lbtcOutcomeAPIError  = "api_error"
	lbtcOutcomeParseErr  = "parse_error"
)
