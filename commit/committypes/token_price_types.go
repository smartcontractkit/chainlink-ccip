package committypes

import (
	"time"

	"github.com/smartcontractkit/chainlink-ccip/shared"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/libocr/offchainreporting2plus/types"
)

type TokenPriceOutcome struct {
	TokenPrices []cciptypes.TokenPrice `json:"tokenPrices"`
}

type TokenPriceObservation struct {
	FeedTokenPrices       []cciptypes.TokenPrice                  `json:"feedTokenPrices"`
	FeeQuoterTokenUpdates map[types.Account]shared.TimestampedBig `json:"feeQuoterTokenUpdates"`
	FChain                map[cciptypes.ChainSelector]int         `json:"fChain"`
	Timestamp             time.Time                               `json:"timestamp"`
}
