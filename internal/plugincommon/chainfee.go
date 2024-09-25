package plugincommon

import (
	"math/big"
	"time"

	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon/consensus"
)

type ChainFeeUSDPrices struct {
	ExecutionFeePriceUSD *big.Int `json:"execFee"`
	DataAvFeePriceUSD    *big.Int `json:"daFee"`
}

type ChainFeeUpdate struct {
	ChainFee  ChainFeeUSDPrices `json:"chainFee"`
	Timestamp time.Time         `json:"timestamp"`
}

// FromPackedFee creates a new ChainFeeUpdate
// @param packedFee: Is the fee components packed into a single big.Int
// packedFee = (dataAvFeeUSD << 112) | executionFeeUSD
func FromPackedFee(packedFee *big.Int) ChainFeeUSDPrices {
	ones112 := big.NewInt(0)
	for i := 0; i < 112; i++ {
		ones112 = ones112.SetBit(ones112, i, 1)
	}

	execFee := new(big.Int).And(packedFee, ones112)
	daFee := new(big.Int).Rsh(packedFee, 112)
	return ChainFeeUSDPrices{
		ExecutionFeePriceUSD: execFee,
		DataAvFeePriceUSD:    daFee,
	}
}

// ToPackedFee PackedFee is a Bitwise operation:
// (dataAvFeeUSD << 112) | executionFeeUSD
// nolint:lll
// https://github.com/smartcontractkit/chainlink/blob/60e8b1181dd74b66903cf5b9a8427557b85357ec/contracts/src/v0.8/ccip/FeeQuoter.sol#L498
func (c ChainFeeUSDPrices) ToPackedFee() *big.Int {
	daShifted := new(big.Int).Lsh(c.DataAvFeePriceUSD, 112)
	return new(big.Int).Or(daShifted, c.ExecutionFeePriceUSD)
}

// ChainFeeUpdateAggregator aggregates a slice of ChainFeeUpdates into a single ChainFeeUpdate
// by taking the median of each price component and the timestamps
func ChainFeeUpdateAggregator(updates []ChainFeeUpdate) ChainFeeUpdate {
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

	return ChainFeeUpdate{
		ChainFee: ChainFeeUSDPrices{
			ExecutionFeePriceUSD: medianExecFeeUSD,
			DataAvFeePriceUSD:    medianDataAvFeeUSD,
		},
		Timestamp: consensus.TimestampsMedian(timestamps),
	}
}
