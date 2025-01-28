package tokenprice

import (
	"context"
	"time"

	"github.com/smartcontractkit/chainlink-ccip/internal/plugintypes"

	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

type Query struct {
}

type Outcome struct {
	TokenPrices cciptypes.TokenPriceMap `json:"tokenPrices"`
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
	TrackTokenPricesObservation(obs Observation)
	TrackTokenPricesOutcome(outcome Outcome)
}

type NoopMetrics struct{}

func (n NoopMetrics) TrackTokenPricesObservation(Observation) {}

func (n NoopMetrics) TrackTokenPricesOutcome(Outcome) {}
