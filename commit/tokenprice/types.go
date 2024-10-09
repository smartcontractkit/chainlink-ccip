package tokenprice

import (
	"context"
	"time"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-ccip/internal/plugintypes"
)

type Query struct {
}

type Outcome struct {
	TokenPrices []cciptypes.TokenPrice `json:"tokenPrices"`
}

type Observation struct {
	FeedTokenPrices       []cciptypes.TokenPrice                                         `json:"feedTokenPrices"`
	FeeQuoterTokenUpdates map[cciptypes.UnknownEncodedAddress]plugintypes.TimestampedBig `json:"feeQuoterTokenUpdates"`
	FChain                map[cciptypes.ChainSelector]int                                `json:"fChain"`
	Timestamp             time.Time                                                      `json:"timestamp"`
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
