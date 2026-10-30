package usdc

import (
	"context"
	"errors"
	nethttp "net/http"

	"go.opentelemetry.io/otel/attribute"

	"github.com/smartcontractkit/chainlink-ccip/execute/tokendata"
	"github.com/smartcontractkit/chainlink-ccip/execute/tokendata/http"
	"github.com/smartcontractkit/chainlink-ccip/execute/tokendata/metricsutil"
)

var attestationFetchCounter = metricsutil.NewLazyCounter("ccip_exec_tokendata_usdc_attestation_fetch")

// trackAttestationFetch counts one USDC attestation fetch by outcome.
func trackAttestationFetch(outcome string) {
	attestationFetchCounter.Add(context.Background(), 1, attribute.String("outcome", outcome))
}

const (
	usdcOutcomeOK          = "ok"
	usdcOutcomeNotReady    = "not_ready"
	usdcOutcomeRateLimited = "rate_limited"
	usdcOutcomeTimeout     = "timeout"
	usdcOutcomeParseError  = "parse_error"
	usdcOutcomeHTTPError   = "http_error"
	usdcOutcomeDataMissing = "data_missing"
)

func attestationFetchOutcome(err error, status http.HTTPStatus) string {
	switch {
	case errors.Is(err, tokendata.ErrNotReady), status == nethttp.StatusNotFound:
		return usdcOutcomeNotReady
	case errors.Is(err, tokendata.ErrRateLimit):
		return usdcOutcomeRateLimited
	case errors.Is(err, tokendata.ErrTimeout):
		return usdcOutcomeTimeout
	default:
		return usdcOutcomeHTTPError
	}
}
