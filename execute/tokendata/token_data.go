package tokendata

import (
	"context"
	"encoding/hex"
	"errors"
	"time"

	"github.com/patrickmn/go-cache"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"

	"github.com/smartcontractkit/chainlink-ccip/pkg/reader"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

const (
	// cacheExpiration sets how long we keep the seen attestations in the cache for computing duration.
	// We can't keep them forever, so we don't track messages which are waiting for longer than 70 minutes
	cacheExpiration = 70 * time.Minute
	cacheCleanup    = 10 * time.Minute
)

var (
	ErrDataMissing     = errors.New("token data missing")
	ErrNotReady        = errors.New("token data not ready")
	ErrRateLimit       = errors.New("token data API is being rate limited")
	ErrTimeout         = errors.New("token data API timed out")
	ErrUnknownResponse = errors.New("unexpected response from attestation API")

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

type AttestationStatus struct {
	// ID is the identifier of the message that the attestation was fetched for
	ID cciptypes.Bytes
	// Attestation is the attestation data fetched from the API, encoded in bytes
	Attestation cciptypes.Bytes
	// Error is the error that occurred during fetching the attestation data
	Error error
}

func SuccessAttestationStatus(id cciptypes.Bytes, attestation cciptypes.Bytes) AttestationStatus {
	return AttestationStatus{ID: id, Attestation: attestation}
}

func ErrorAttestationStatus(err error) AttestationStatus {
	return AttestationStatus{Error: err}
}

type AttestationClient interface {
	Attestations(
		ctx context.Context,
		msgs map[cciptypes.ChainSelector]map[reader.MessageTokenID]cciptypes.Bytes,
	) (map[cciptypes.ChainSelector]map[reader.MessageTokenID]AttestationStatus, error)

	Token() string
}

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

type FakeAttestationClient struct {
	Data map[string]AttestationStatus
}

func (f *FakeAttestationClient) Attestations(
	_ context.Context,
	messagesByChain map[cciptypes.ChainSelector]map[reader.MessageTokenID]cciptypes.Bytes,
) (map[cciptypes.ChainSelector]map[reader.MessageTokenID]AttestationStatus, error) {
	outcome := make(map[cciptypes.ChainSelector]map[reader.MessageTokenID]AttestationStatus)
	for chainSelector, messagesByTokenID := range messagesByChain {
		outcome[chainSelector] = make(map[reader.MessageTokenID]AttestationStatus)
		for tokenID, message := range messagesByTokenID {
			outcome[chainSelector][tokenID] = f.Data[string(message)]
		}
	}
	return outcome, nil
}

func (f *FakeAttestationClient) Token() string {
	return ""
}
