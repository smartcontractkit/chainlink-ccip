package usdc

import (
	"errors"
	nethttp "net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"

	"github.com/smartcontractkit/chainlink-ccip/execute/tokendata"
	"github.com/smartcontractkit/chainlink-ccip/execute/tokendata/http"
)

var PromUSDCAttestationFetchCounter = promauto.NewCounterVec(
	prometheus.CounterOpts{
		Name: "ccip_exec_tokendata_usdc_attestation_fetch",
		Help: "Count of USDC attestation fetches by outcome",
	},
	[]string{"outcome"},
)

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
