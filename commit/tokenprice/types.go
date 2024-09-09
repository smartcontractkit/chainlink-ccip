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

type Observer interface {
	// ObserveFeedTokenPrices returns the latest token prices from the feed chain
	ObserveFeedTokenPrices(ctx context.Context) []cciptypes.TokenPrice

	// ObserveFeeQuoterTokenUpdates returns the latest token prices from the FeeQuoter on the dest chain
	ObserveFeeQuoterTokenUpdates(ctx context.Context) map[types.Account]shared.TimestampedBig

	ObserveFDestChain() (*int, error)
}

//type TimestampedBig struct {
//	Timestamp time.Time        `json:"timestamp"`
//	Value     cciptypes.BigInt `json:"value"`
//}
//
//func NewTimestampedBig(value int64, timestamp time.Time) TimestampedBig {
//	return TimestampedBig{
//		Value:     cciptypes.BigInt{Int: big.NewInt(value)},
//		Timestamp: timestamp,
//	}
//}
//
//// NewTimestampedBigNow NewTimestampedBig Returns an update with timestamp now as UTC
//func NewTimestampedBigNow(value int64) TimestampedBig {
//	return TimestampedBig{
//		Value:     cciptypes.BigInt{Int: big.NewInt(value)},
//		Timestamp: time.Now().UTC(),
//	}
//}
