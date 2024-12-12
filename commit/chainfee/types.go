package chainfee

import (
	"math/big"
	"time"

	"github.com/smartcontractkit/chainlink-common/pkg/types"

	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

type Query struct {
}

type Outcome struct {
	// Each Gas Price is the combination of Execution and DataAvailability Fees using bitwise operations
	GasPrices []cciptypes.GasPriceChain `json:"gasPrices"`
}

type Observation struct {
	// FeeComponents: from the source chains, via chain writer
	FeeComponents map[cciptypes.ChainSelector]types.ChainFeeComponents `json:"feeComponents"`
	//NativeTokenPrices: from the source chains, via fee quoter (after getting the native token address from Router)
	NativeTokenPrices map[cciptypes.ChainSelector]cciptypes.BigInt `json:"nativeTokenPrice"`
	//ChainFeeUpdates: from the dest chain, via fee quoter
	ChainFeeUpdates map[cciptypes.ChainSelector]Update `json:"chainFeeUpdates"`
	FChain          map[cciptypes.ChainSelector]int    `json:"fChain"`
	TimestampNow    time.Time                          `json:"timestamp"`
}

// AggregateObservation is the aggregation of a list of observations
type AggregateObservation struct {
	FeeComponents     map[cciptypes.ChainSelector][]types.ChainFeeComponents `json:"feeComponents"`
	NativeTokenPrices map[cciptypes.ChainSelector][]cciptypes.BigInt         `json:"nativeTokenPrice"`
	FChain            map[cciptypes.ChainSelector][]int                      `json:"fChain"`
	ChainFeeUpdates   map[cciptypes.ChainSelector][]Update                   `json:"chainFeeUpdates"`
	Timestamps        []time.Time                                            `json:"timestamps"`
}

type ComponentsUSDPrices struct {
	ExecutionFeePriceUSD *big.Int `json:"execFee"`
	DataAvFeePriceUSD    *big.Int `json:"daFee"`
}

type Update struct {
	ChainFee  ComponentsUSDPrices `json:"chainFee"`
	Timestamp time.Time           `json:"timestamp"`
}
