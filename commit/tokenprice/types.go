package tokenprice

import (
	"context"
	"time"

	"github.com/smartcontractkit/chainlink-ccip/shared"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/libocr/offchainreporting2plus/types"
)

type Query struct {
}

type Outcome struct {
	TokenPrices []cciptypes.TokenPrice `json:"tokenPrices"`
}

type Observation struct {
	FeedTokenPrices       []cciptypes.TokenPrice                  `json:"feedTokenPrices"`
	FeeQuoterTokenUpdates map[types.Account]shared.TimestampedBig `json:"feeQuoterTokenUpdates"`
	FDestChain            int                                     `json:"fDestChain"`
	Timestamp             time.Time                               `json:"timestamp"`
}

// AggregateObservation is the aggregation of a list of observations
type AggregateObservation struct {
	FeedTokenPrices       map[types.Account][]cciptypes.TokenPrice
	FeeQuoterTokenUpdates map[types.Account][]shared.TimestampedBig
	FDestChain            []int
	Timestamps            []time.Time
}

// ConsensusObservation holds the consensus values for all observations in a round
type ConsensusObservation struct {
	FeedTokenPrices       map[types.Account]cciptypes.TokenPrice
	FeeQuoterTokenUpdates map[types.Account]shared.TimestampedBig
	FDestChain            int
	Timestamp             time.Time
}

type Observer interface {
	// ObserveFeedTokenPrices returns the latest token prices from the feed chain
	ObserveFeedTokenPrices(ctx context.Context) []cciptypes.TokenPrice

	// ObserveFeeQuoterTokenUpdates returns the latest token prices from the FeeQuoter on the dest chain
	ObserveFeeQuoterTokenUpdates(ctx context.Context) map[types.Account]shared.TimestampedBig

	ObserveFDestChain() (int, error)
}
