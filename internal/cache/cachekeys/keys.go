package cachekeys

import (
	"fmt"

	"github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

// TokenDecimals creates a cache key for token decimals
func TokenDecimals(token ccipocr3.UnknownEncodedAddress, address string) string {
	return fmt.Sprintf("token-decimals:%s:%s", token, address)
}

// FeeQuoterTokenUpdate creates a cache key for fee quoter token updates
func FeeQuoterTokenUpdate(token ccipocr3.UnknownEncodedAddress, chain ccipocr3.ChainSelector) string {
	return fmt.Sprintf("fee-quoter-update:%d:%s", chain, token)
}

// TokenPrice creates a cache key for token USD prices
func FeedPricesUSD(token ccipocr3.UnknownEncodedAddress) string {
	return fmt.Sprintf("token-price:%s", token)
}
