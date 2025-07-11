package cctpv2

import (
	"fmt"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"

	"github.com/smartcontractkit/chainlink-ccip/execute/tokendata/http"
)

var (
	PromHTTPRequest = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "ccip_cctp_v2_request",
			Help: "Tracks success and duration of CCTPv2 attestation HTTP requests",
			Buckets: []float64{
				float64(50 * time.Millisecond),
				float64(100 * time.Millisecond),
				float64(250 * time.Millisecond),
				float64(500 * time.Millisecond),
				float64(time.Second),
				float64(2 * time.Second),
				float64(5 * time.Second),
			},
		},
		[]string{"sourceDomainID", "status", "error"},
	)
)

func TrackHTTPRequest(sourceDomainID uint32, status http.HTTPStatus, isErr bool, duration time.Duration) {
	sourceDomainIDStr := fmt.Sprintf("%d", sourceDomainID)
	statusStr := fmt.Sprintf("%d", status)
	errorStr := fmt.Sprintf("%t", isErr)

	PromHTTPRequest.
		WithLabelValues(sourceDomainIDStr, statusStr, errorStr).
		Observe(float64(duration.Milliseconds()))
}
