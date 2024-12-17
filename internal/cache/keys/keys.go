package cachekeys

import (
	"fmt"

	"github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

// TokenDecimals creates a cache key for token decimals
func TokenDecimals(token ccipocr3.UnknownEncodedAddress, address string) string {
	return fmt.Sprintf("token-decimals:%s:%s", token, address)
}
