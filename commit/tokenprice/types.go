package tokenprice

import (
	"context"
	"time"

	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

const (
	processorsLabel            = "tokenprice"
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
	FeedTokenPrices       cciptypes.TokenPriceMap                                      `json:"feedTokenPrices"`
	FeeQuoterTokenUpdates map[cciptypes.UnknownEncodedAddress]cciptypes.TimestampedBig `json:"feeQuoterTokenUpdates"`
	FChain                map[cciptypes.ChainSelector]int                              `json:"fChain"`
	Timestamp             time.Time                                                    `json:"timestamp"`
}

func (obs Observation) IsEmpty() bool {
	return len(obs.FeedTokenPrices) == 0 && len(obs.FeeQuoterTokenUpdates) == 0 &&
		len(obs.FChain) == 0 && obs.Timestamp.IsZero()
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
	FeeQuoterTokenUpdates map[cciptypes.UnknownEncodedAddress][]cciptypes.TimestampedBig
	FChain                map[cciptypes.ChainSelector][]int `json:"fChain"`
	Timestamps            []time.Time
}

// ConsensusObservation holds the consensus values for all observations in a round
type ConsensusObservation struct {
	FeedTokenPrices       map[cciptypes.UnknownEncodedAddress]cciptypes.TokenPrice
	FeeQuoterTokenUpdates map[cciptypes.UnknownEncodedAddress]cciptypes.TimestampedBig
	FChain                map[cciptypes.ChainSelector]int `json:"fChain"`
	Timestamp             time.Time
}

type Observer interface {
	// ObserveFeedTokenPrices returns the latest token prices from the feed chain
	ObserveFeedTokenPrices(ctx context.Context) []cciptypes.TokenPrice

	// ObserveFeeQuoterTokenUpdates returns the latest token prices from the FeeQuoter on the dest chain
	ObserveFeeQuoterTokenUpdates(ctx context.Context) map[cciptypes.UnknownEncodedAddress]cciptypes.TimestampedBig

	ObserveFChain() map[cciptypes.ChainSelector]int
}
