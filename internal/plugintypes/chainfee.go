package plugintypes

import (
	"math/big"
	"time"
)

type ChainFeePrices struct {
	ExecutionFeePriceUSD *big.Int `json:"execFee"`
	DataAvFeePriceUSD    *big.Int `json:"daFee"`
}

type ChainFeeUpdate struct {
	ChainFee  ChainFeePrices `json:"chainFee"`
	Timestamp time.Time      `json:"timestamp"`
}

// FromPackedFee creates a new ChainFeeUpdate
// @param packedFee: Is the fee components packed into a single big.Int
// packedFee = (dataAvFeeUSD << 112) | executionFeeUSD
func FromPackedFee(packedFee *big.Int) ChainFeePrices {
	ones112 := big.NewInt(0)
	for i := 0; i < 112; i++ {
		ones112 = ones112.SetBit(ones112, i, 1)
	}

	execFee := new(big.Int).And(packedFee, ones112)
	daFee := new(big.Int).Rsh(packedFee, 112)
	return ChainFeePrices{
		ExecutionFeePriceUSD: execFee,
		DataAvFeePriceUSD:    daFee,
	}
}

// ToPackedFee PackedFee is a Bitwise operation:
// (dataAvFeeUSD << 112) | executionFeeUSD
// nolint:lll
// https://github.com/smartcontractkit/chainlink/blob/60e8b1181dd74b66903cf5b9a8427557b85357ec/contracts/src/v0.8/ccip/FeeQuoter.sol#L498
func (c ChainFeePrices) ToPackedFee() *big.Int {
	daShifted := new(big.Int).Lsh(c.DataAvFeePriceUSD, 112)
	return new(big.Int).Or(daShifted, c.ExecutionFeePriceUSD)
}
