package plugintypes

import (
	"math/big"
	"time"

	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

type TimestampedBig struct {
	Timestamp time.Time        `json:"timestamp"`
	Value     cciptypes.BigInt `json:"value"`
}

// TimestampedUnixBig Maps to on-chain struct
// https://github.com/smartcontractkit/chainlink/blob/37f3132362ec90b0b1c12fb1b69b9c16c46b399d/contracts/src/v0.8/ccip/libraries/Internal.sol#L43-L47
//
//nolint:lll //url
type TimestampedUnixBig struct {
	Value     *big.Int `json:"value"`
	Timestamp uint32   `json:"timestamp"`
}

func NewTimestampedBig(value int64, timestamp time.Time) TimestampedBig {
	return TimestampedBig{
		Value:     cciptypes.BigInt{Int: big.NewInt(value)},
		Timestamp: timestamp,
	}
}

func TimeStampedBigFromUnix(input TimestampedUnixBig) TimestampedBig {
	return TimestampedBig{
		Value:     cciptypes.NewBigInt(input.Value),
		Timestamp: time.Unix(int64(input.Timestamp), 0),
	}
}
