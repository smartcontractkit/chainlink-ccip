package chainfee

import (
	"math/big"
	"time"

	"github.com/smartcontractkit/chainlink-common/pkg/types"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon/consensus"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugintypes"
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

func FeeUpdatesFromTimestampedBig(
	updates map[cciptypes.ChainSelector]plugintypes.TimestampedBig,
) map[cciptypes.ChainSelector]Update {
	chainFeeUpdates := make(map[cciptypes.ChainSelector]Update, len(updates))
	for chain, u := range updates {
		chainFeeUpdates[chain] = Update{
			ChainFee:  FromPackedFee(u.Value.Int),
			Timestamp: u.Timestamp,
		}
	}
	return chainFeeUpdates
}

// FromPackedFee creates a new Update
// @param packedFee: Is the fee components packed into a single big.Int
// packedFee = (dataAvFeeUSD << 112) | executionFeeUSD
func FromPackedFee(packedFee *big.Int) ComponentsUSDPrices {
	ones112 := big.NewInt(0)
	for i := 0; i < 112; i++ {
		ones112 = ones112.SetBit(ones112, i, 1)
	}

	execFee := new(big.Int).And(packedFee, ones112)
	daFee := new(big.Int).Rsh(packedFee, 112)
	return ComponentsUSDPrices{
		ExecutionFeePriceUSD: execFee,
		DataAvFeePriceUSD:    daFee,
	}
}

// ToPackedFee PackedFee is a Bitwise operation:
// (dataAvFeeUSD << 112) | executionFeeUSD
//
// https://github.com/smartcontractkit/chainlink/blob/60e8b1181dd74b66903cf5b9a8427557b85357ec/contracts/src/v0.8/ccip/FeeQuoter.sol#L498
//
//nolint:lll
func (c ComponentsUSDPrices) ToPackedFee() *big.Int {
	daShifted := new(big.Int).Lsh(c.DataAvFeePriceUSD, 112)
	return new(big.Int).Or(daShifted, c.ExecutionFeePriceUSD)
}

// ChainFeeUpdateAggregator aggregates a slice of ChainFeeUpdates into a single Update
// by taking the median of each price component and the timestamps
func ChainFeeUpdateAggregator(updates []Update) Update {
	execFeeUSDs := make([]*big.Int, len(updates))
	dataAvFeeUSDs := make([]*big.Int, len(updates))
	timestamps := make([]time.Time, len(updates))
	for i := range updates {
		execFeeUSDs[i] = updates[i].ChainFee.ExecutionFeePriceUSD
		dataAvFeeUSDs[i] = updates[i].ChainFee.DataAvFeePriceUSD
		timestamps[i] = updates[i].Timestamp
	}
	medianExecFeeUSD := consensus.Median(execFeeUSDs, func(a, b *big.Int) bool {
		return a.Cmp(b) == -1
	})
	medianDataAvFeeUSD := consensus.Median(dataAvFeeUSDs, func(a, b *big.Int) bool {
		return a.Cmp(b) == -1
	})

	return Update{
		ChainFee: ComponentsUSDPrices{
			ExecutionFeePriceUSD: medianExecFeeUSD,
			DataAvFeePriceUSD:    medianDataAvFeeUSD,
		},
		Timestamp: consensus.TimestampsMedian(timestamps),
	}
}
