package chainfee

import (
	"time"

	"github.com/smartcontractkit/chainlink-ccip/internal/plugintypes"
	"github.com/smartcontractkit/chainlink-common/pkg/types"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
)

type Query struct {
}

type Outcome struct {
	// Each Gas Price is the combination of Execution and DataAvailability Fees using bitwise operations
	GasPrices []cciptypes.GasPriceChain `json:"gasPrices"`
}

type Observation struct {
	FeeComponents        map[cciptypes.ChainSelector]types.ChainFeeComponents   `json:"feeComponents"`
	NativeTokenPrice     map[cciptypes.ChainSelector]cciptypes.BigInt           `json:"nativeTokenPrice"`
	FChain               map[cciptypes.ChainSelector]int                        `json:"fChain"`
	ChainFeePriceUpdates map[cciptypes.ChainSelector]plugintypes.TimestampedBig `json:"chainFeePriceUpdates"`
	Timestamp            time.Time                                              `json:"timestamp"`
}

// AggregateObservation is the aggregation of a list of observations
type AggregateObservation struct {
	FeeComponents        map[cciptypes.ChainSelector][]types.ChainFeeComponents   `json:"feeComponents"`
	NativeTokenPrices    map[cciptypes.ChainSelector][]cciptypes.BigInt           `json:"nativeTokenPrice"`
	FChain               map[cciptypes.ChainSelector][]int                        `json:"fChain"`
	ChainFeePriceUpdates map[cciptypes.ChainSelector][]plugintypes.TimestampedBig `json:"chainFeePriceUpdate"`
	Timestamps           []time.Time                                              `json:"timestamps"`
}

// ConsensusObservation holds the consensus values for all observations in a round
type ConsensusObservation struct {
	FeeComponents        map[cciptypes.ChainSelector]types.ChainFeeComponents   `json:"feeComponents"`
	NativeTokenPrices    map[cciptypes.ChainSelector]cciptypes.BigInt           `json:"nativeTokenPrice"`
	FChain               map[cciptypes.ChainSelector]int                        `json:"fChain"`
	ChainFeePriceUpdates map[cciptypes.ChainSelector]plugintypes.TimestampedBig `json:"chainFeePriceUpdates"`
	Timestamp            time.Time                                              `json:"timestamp"`
	ShouldUpdate         bool                                                   `json:"shouldUpdate"`
}

var EmptyConsensusObservation = ConsensusObservation{}
