package chainfee

import (
	"time"

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
	FeeComponents map[cciptypes.ChainSelector]types.ChainFeeComponents `json:"feeComponents"`
	FChain        map[cciptypes.ChainSelector]int                      `json:"fChain"`
	Timestamp     time.Time
}
