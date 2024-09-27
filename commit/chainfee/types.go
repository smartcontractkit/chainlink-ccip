package chainfee

import (
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
)

type Query struct {
}

type Outcome struct {
	GasPrices []cciptypes.GasPriceChain `json:"gasPrices"`
}

type Observation struct {
	GasPrices []cciptypes.GasPriceChain `json:"gasPrices"`
}
