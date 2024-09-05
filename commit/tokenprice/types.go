package tokenprice

import (
	"time"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
)

type Query struct {
}

type Outcome struct {
	TokenPrices []cciptypes.TokenPrice `json:"tokenPrices"`
}

type Observation struct {
	// FeedTokenPrices for tokens from the feeds on the feed chain
	FeedTokenPrices []cciptypes.TokenPrice `json:"feedTokenPrices"`
	// PriceRegistryTokenUpdates for tokens from the PriceRegistry on the dest chain
	PriceRegistryTokenUpdates []cciptypes.TokenPrice `json:"priceRegistryTokenPrices"`
	// Observation time
	Timestamp time.Time
}
