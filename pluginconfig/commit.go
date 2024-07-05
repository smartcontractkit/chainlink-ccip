package pluginconfig

import (
	"fmt"
	"time"

	"github.com/smartcontractkit/libocr/offchainreporting2plus/types"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
)

type CommitPluginConfig struct {
	// DestChain is the ccip destination chain configured for the commit plugin DON.
	DestChain cciptypes.ChainSelector `json:"destChain"`

	// PricedTokens is a list of tokens that we want to submit price updates for.
	PricedTokens []types.Account `json:"pricedTokens"`

	// TokenPricesObserver indicates that the node can observe token prices.
	TokenPricesObserver bool `json:"tokenPricesObserver"`

	// NewMsgScanBatchSize is the number of max new messages to scan, typically set to 256.
	NewMsgScanBatchSize int `json:"newMsgScanBatchSize"`

	// SyncTimeout is the timeout for syncing the commit plugin reader.
	SyncTimeout time.Duration `json:"syncTimeout"`

	// SyncFrequency is the frequency at which the commit plugin reader should sync.
	SyncFrequency time.Duration `json:"syncFrequency"`
}

func (c CommitPluginConfig) Validate() error {
	if c.DestChain == cciptypes.ChainSelector(0) {
		return fmt.Errorf("destChain not set")
	}

	if len(c.PricedTokens) == 0 {
		return fmt.Errorf("priced tokens not set, at least one priced token is required")
	}

	if c.NewMsgScanBatchSize == 0 {
		return fmt.Errorf("newMsgScanBatchSize not set")
	}

	return nil
}
