package tokendata

import (
	"context"
	"encoding/hex"
	"time"

	"github.com/patrickmn/go-cache"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"

	"github.com/smartcontractkit/chainlink-ccip/pkg/reader"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
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
		[]string{"token"},
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
		[]string{"token"},
	)
)

type ObservedAttestationClient struct {
	lggr                   logger.Logger
	delegate               AttestationClient
	timeToAttestationCache *cache.Cache
}

func NewObservedAttestationClient(lggr logger.Logger, delegate AttestationClient) *ObservedAttestationClient {
	return &ObservedAttestationClient{
		lggr:                   lggr,
		delegate:               delegate,
		timeToAttestationCache: cache.New(cacheExpiration, cacheCleanup),
	}
}

func (o *ObservedAttestationClient) Attestations(
	ctx context.Context,
	msgs map[cciptypes.ChainSelector]map[reader.MessageTokenID]cciptypes.Bytes,
) (map[cciptypes.ChainSelector]map[reader.MessageTokenID]AttestationStatus, error) {
	start := time.Now()
	attestations, err := o.delegate.Attestations(ctx, msgs)
	if err != nil {
		return nil, err
	}
	duration := time.Since(start)

	for _, msgIDToAttestation := range attestations {
		for _, attestation := range msgIDToAttestation {
			o.trackTimeToAttestation(start, hex.EncodeToString(attestation.ID), attestation.Error)
			promAttestationDurations.
				WithLabelValues(o.Token()).
				Observe(float64(duration))
		}
	}
	return attestations, err
}

func (o *ObservedAttestationClient) Token() string {
	return o.delegate.Token()
}

func (o *ObservedAttestationClient) trackTimeToAttestation(start time.Time, id string, err error) {
	// If attestation is not successful, we mark the message as seen by putting it into the cache
	// with the current timestamp
	if err != nil {
		_, seen := o.timeToAttestationCache.Get(id)
		if !seen {
			o.timeToAttestationCache.Set(id, start, cache.DefaultExpiration)
		}
	} else {
		// If attestation is successful, we calculate the time taken to attest the message
		messageSeenFirst, seen := o.timeToAttestationCache.Get(id)
		// If seen is false it means that attestation was resolved immediately in the first OCR3 round it was spotted,
		// and we don't need to track that
		if seen {
			duration := time.Since(messageSeenFirst.(time.Time))
			promTimeToAttestation.WithLabelValues(o.Token()).Observe(float64(duration))
			o.lggr.Infow("Observed time to attestation for a message",
				"token", o.Token(),
				"hash", id,
				"duration", duration.String(),
			)
			o.timeToAttestationCache.Delete(id)
		}
	}
}
