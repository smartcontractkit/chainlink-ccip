package usdc

import (
	"context"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"

	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

var (
	buckets = []float64{
		float64(50 * time.Millisecond),
		float64(100 * time.Millisecond),
		float64(200 * time.Millisecond),
		float64(500 * time.Millisecond),
		float64(700 * time.Millisecond),
		float64(time.Second),
		float64(2 * time.Second),
		float64(5 * time.Second),
	}

	promAttestationDurations = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "ccip_attestation_request_duration",
			Help:    "The amount of time elapsed during the attestation request",
			Buckets: buckets,
		},
		[]string{"statusCode"},
	)
)

type observedHTTPClient struct {
	HTTPClient
	requestDurations *prometheus.HistogramVec
}

func newObservedHTTPClient(httpClient HTTPClient) *observedHTTPClient {
	return &observedHTTPClient{
		HTTPClient:       httpClient,
		requestDurations: promAttestationDurations,
	}
}

func (o *observedHTTPClient) Get(
	ctx context.Context,
	messageHash cciptypes.Bytes32,
) (cciptypes.Bytes, HTTPStatus, error) {
	start := time.Now()
	result, status, err := o.HTTPClient.Get(ctx, messageHash)
	duration := time.Since(start)

	o.requestDurations.
		WithLabelValues(strconv.Itoa(int(status))).
		Observe(float64(duration))

	return result, status, err
}
