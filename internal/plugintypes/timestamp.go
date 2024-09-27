package plugintypes

import (
	"math/big"
	"time"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
)

type TimestampedBig struct {
	Timestamp time.Time        `json:"timestamp"`
	Value     cciptypes.BigInt `json:"value"`
}

func NewTimestampedBig(value int64, timestamp time.Time) TimestampedBig {
	return TimestampedBig{
		Value:     cciptypes.BigInt{Int: big.NewInt(value)},
		Timestamp: timestamp,
	}
}

// NewTimestampedBigNow NewTimestampedBig Returns an update with timestamp now as UTC
func NewTimestampedBigNow(value int64) TimestampedBig {
	return TimestampedBig{
		Value:     cciptypes.BigInt{Int: big.NewInt(value)},
		Timestamp: time.Now().UTC(),
	}
}
