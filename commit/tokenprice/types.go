package tokenprice

import (
	"context"
	"time"

	"github.com/smartcontractkit/chainlink-ccip/internal/plugintypes"

	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

const (
	tokenPricesLabel           = "tokenPrices"
	feedTokenPricesLabel       = "feedTokenPrices"
	feeQuoterTokenUpdatesLabel = "feeQuoterTokenUpdates"
)

type Query struct {
}

type Outcome struct {
	TokenPrices cciptypes.TokenPriceMap `json:"tokenPrices"`
}

func (out Outcome) Stats() map[string]int {
	return map[string]int{
		tokenPricesLabel: len(out.TokenPrices),
	}
}

type Observation struct {
	FeedTokenPrices       cciptypes.TokenPriceMap                                        `json:"feedTokenPrices"`
	FeeQuoterTokenUpdates map[cciptypes.UnknownEncodedAddress]plugintypes.TimestampedBig `json:"feeQuoterTokenUpdates"`
	FChain                map[cciptypes.ChainSelector]int                                `json:"fChain"`
	Timestamp             time.Time                                                      `json:"timestamp"`
}

func (obs Observation) IsEmpty() bool {
	return len(obs.FeedTokenPrices) == 0 && len(obs.FeeQuoterTokenUpdates) == 0 && len(obs.FChain) == 0
}

func (obs Observation) Stats() map[string]int {
	return map[string]int{
		feedTokenPricesLabel:       len(obs.FeedTokenPrices),
		feeQuoterTokenUpdatesLabel: len(obs.FeeQuoterTokenUpdates),
	}
}

// AggregateObservation is the aggregation of a list of observations
type AggregateObservation struct {
	FeedTokenPrices       map[cciptypes.UnknownEncodedAddress][]cciptypes.TokenPrice
	FeeQuoterTokenUpdates map[cciptypes.UnknownEncodedAddress][]plugintypes.TimestampedBig
	FChain                map[cciptypes.ChainSelector][]int `json:"fChain"`
	Timestamps            []time.Time
}

// ConsensusObservation holds the consensus values for all observations in a round
type ConsensusObservation struct {
	FeedTokenPrices       map[cciptypes.UnknownEncodedAddress]cciptypes.TokenPrice
	FeeQuoterTokenUpdates map[cciptypes.UnknownEncodedAddress]plugintypes.TimestampedBig
	FChain                map[cciptypes.ChainSelector]int `json:"fChain"`
	Timestamp             time.Time
}

type Observer interface {
	// ObserveFeedTokenPrices returns the latest token prices from the feed chain
	ObserveFeedTokenPrices(ctx context.Context) []cciptypes.TokenPrice

	// ObserveFeeQuoterTokenUpdates returns the latest token prices from the FeeQuoter on the dest chain
	ObserveFeeQuoterTokenUpdates(ctx context.Context) map[cciptypes.UnknownEncodedAddress]plugintypes.TimestampedBig

	ObserveFChain() map[cciptypes.ChainSelector]int
}

// MetricsReporter exposes only relevant methods for reporting token prices from metrics.Reporter
type MetricsReporter interface {
	TrackProcessorLatency(processor string, method string, latency time.Duration)
	TrackProcessorObservation(processor string, obs plugintypes.Trackable, err error)
	TrackProcessorOutcome(processor string, out plugintypes.Trackable, err error)
}

type NoopMetrics struct{}

func (n NoopMetrics) TrackTokenPricesObservation(Observation) {}

func (n NoopMetrics) TrackTokenPricesOutcome(Outcome) {}

func (n NoopMetrics) TrackProcessorLatency(string, string, time.Duration) {}

func (n NoopMetrics) TrackProcessorObservation(string, plugintypes.Trackable, error) {}

func (n NoopMetrics) TrackProcessorOutcome(string, plugintypes.Trackable, error) {}
