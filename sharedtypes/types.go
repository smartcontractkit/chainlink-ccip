package sharedtypes

import (
	"time"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
)

type NumericalUpdate struct {
	Timestamp time.Time
	Value     cciptypes.BigInt
}
