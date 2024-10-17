package tokenprice

import (
	"context"
	"time"

	"github.com/smartcontractkit/chainlink-ccip/internal/plugintypes"

	"github.com/smartcontractkit/libocr/offchainreporting2plus/types"

	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

type Query struct {
}

type Outcome struct {
	TokenPrices []cciptypes.TokenPrice `json:"tokenPrices"`
}

type Observation struct {
	FeedTokenPrices       []cciptypes.TokenPrice                       `json:"feedTokenPrices"`
	FeeQuoterTokenUpdates map[types.Account]plugintypes.TimestampedBig `json:"feeQuoterTokenUpdates"`
	FChain                map[cciptypes.ChainSelector]int              `json:"fChain"`
	Timestamp             time.Time                                    `json:"timestamp"`
}

// AggregateObservation is the aggregation of a list of observations
type AggregateObservation struct {
	FeedTokenPrices       map[types.Account][]cciptypes.TokenPrice
	FeeQuoterTokenUpdates map[types.Account][]plugintypes.TimestampedBig
	FChain                map[cciptypes.ChainSelector][]int `json:"fChain"`
	Timestamps            []time.Time
}

// ConsensusObservation holds the consensus values for all observations in a round
type ConsensusObservation struct {
	FeedTokenPrices       map[types.Account]cciptypes.TokenPrice
	FeeQuoterTokenUpdates map[types.Account]plugintypes.TimestampedBig
	FChain                map[cciptypes.ChainSelector]int `json:"fChain"`
	Timestamp             time.Time
}

type Observer interface {
	// ObserveFeedTokenPrices returns the latest token prices from the feed chain
	ObserveFeedTokenPrices(ctx context.Context) []cciptypes.TokenPrice

	// ObserveFeeQuoterTokenUpdates returns the latest token prices from the FeeQuoter on the dest chain
	ObserveFeeQuoterTokenUpdates(ctx context.Context) map[types.Account]plugintypes.TimestampedBig

	ObserveFChain() map[cciptypes.ChainSelector]int
}
