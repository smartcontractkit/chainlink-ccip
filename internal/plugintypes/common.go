package plugintypes

import (
	"math/big"

	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

type SeqNumChain struct {
	ChainSel cciptypes.ChainSelector `json:"chainSel"`
	SeqNum   cciptypes.SeqNum        `json:"seqNum"`
}

func NewSeqNumChain(chainSel cciptypes.ChainSelector, seqNum cciptypes.SeqNum) SeqNumChain {
	return SeqNumChain{
		ChainSel: chainSel,
		SeqNum:   seqNum,
	}
}

type ChainRange struct {
	ChainSel    cciptypes.ChainSelector `json:"chain"`
	SeqNumRange cciptypes.SeqNumRange   `json:"seqNumRange"`
}

type DonID = uint32

// USD18 is a small unit of USD, where 1 USD18 is 1e-18 USD, meaning it is 18 decimal places smaller than 1 USD.
//
// 1 USD18 = 1e-18 USD   = 0.000000000000000001 USD
// 1 USD   = 1e18  USD18 = 1,000,000,000,000,000,000 USD18
//
// Token prices stored in many contracts (e.g. FeeQuoter) are denominated in USD18
type USD18 = *big.Int

func NewUSD18(value int64) USD18 {
	return big.NewInt(value)
}

// Trackable is an interface for types that can be tracked using Prometheus metrics.
// That way we are moving responsibility of presenting proper stats and metrics to
// the underlying types. It's meant to be implemented by Observation and Outcome in
// the Commit Plugin's processors.
type Trackable interface {
	Stats() map[string]int
}

// Deprecated: use cciptypes.TimestampedBig
type TimestampedBig = cciptypes.TimestampedBig
