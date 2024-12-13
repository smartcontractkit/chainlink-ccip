package usdc

import (
	"context"
	"strconv"
	"time"

	"github.com/patrickmn/go-cache"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"

	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

const (
	cacheExpiration = 70 * time.Minute
	cacheCleanup    = 10 * time.Minute
)

var (
	promAttestationDurations = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "ccip_attestation_request_duration",
			Help: "The amount of time elapsed during the attestation request",
			Buckets: []float64{
				float64(50 * time.Millisecond),
				float64(100 * time.Millisecond),
				float64(200 * time.Millisecond),
				float64(500 * time.Millisecond),
				float64(700 * time.Millisecond),
				float64(time.Second),
				float64(2 * time.Second),
				float64(5 * time.Second),
			},
		},
		[]string{"statusCode"},
	)
	promTimeToAttestation = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "ccip_time_to_attestation",
			Help: "The time taken to attest a message from seeing it for the first time till " +
				"getting a successful attestation",
			Buckets: []float64{
				float64(30 * time.Second),
				float64(time.Minute),
				float64(2 * time.Minute),
				float64(5 * time.Minute),
				float64(10 * time.Minute),
				float64(30 * time.Minute),
				float64(60 * time.Minute),
			},
		},
		[]string{},
	)
)

type observedHTTPClient struct {
	HTTPClient
	lggr                   logger.Logger
	requestDurations       *prometheus.HistogramVec
	timeToAttestation      *prometheus.HistogramVec
	timeToAttestationCache *cache.Cache
}

func newObservedHTTPClient(
	httpClient HTTPClient,
	lggr logger.Logger,
) *observedHTTPClient {
	return &observedHTTPClient{
		HTTPClient:             httpClient,
		lggr:                   lggr,
		requestDurations:       promAttestationDurations,
		timeToAttestationCache: cache.New(cacheExpiration, cacheCleanup),
		timeToAttestation:      promTimeToAttestation,
	}
}

func (o *observedHTTPClient) Get(
	ctx context.Context,
	messageHash cciptypes.Bytes32,
) (cciptypes.Bytes, HTTPStatus, error) {
	start := time.Now()
	result, status, err := o.HTTPClient.Get(ctx, messageHash)
	duration := time.Since(start)

	o.trackTimeToAttestation(start, messageHash.String(), status)

	o.requestDurations.
		WithLabelValues(strconv.Itoa(int(status))).
		Observe(float64(duration))

	return result, status, err
}

func (o *observedHTTPClient) trackTimeToAttestation(
	start time.Time,
	hash string,
	status HTTPStatus,
) {
	// If status is different from 200, we mark the message as seen by putting it into the cache with
	// the current timestamp
	if status != HTTPStatus(200) {
		_, seen := o.timeToAttestationCache.Get(hash)
		if !seen {
			o.timeToAttestationCache.Set(hash, start, cache.DefaultExpiration)
		}
	} else {
		// If status is 200, we calculate the time taken to attest the message
		messageSeenFirst, seen := o.timeToAttestationCache.Get(hash)
		// If seen is false it means that attestation was resolved immediately in the first OCR3 round it was spotted,
		// and we don't need to track that
		if seen {
			duration := time.Since(messageSeenFirst.(time.Time))
			o.timeToAttestation.WithLabelValues().Observe(float64(duration))
			o.lggr.Infow("Observed time to attestation for USDC message",
				"hash", hash,
				"duration", duration.String(),
			)
			o.timeToAttestationCache.Delete(hash)
		}
	}
}
