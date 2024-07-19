package pluginconfig

import (
	"fmt"
	"time"

	"github.com/smartcontractkit/libocr/offchainreporting2plus/types"

	commonconfig "github.com/smartcontractkit/chainlink-common/pkg/config"
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

// ArbitrumPriceSource is the source of the TOKEN/USD price data of a particular chain
// on Arbitrum.
type ArbitrumPriceSource struct {
	// AggregatorAddress is the address of the price feed TOKEN/USD aggregator on arbitrum.
	AggregatorAddress types.Account `json:"aggregatorAddress"`

	// DeviationPPB is the deviation in parts per billion that the price feed is allowed to deviate
	// from the last written price on-chain before we write a new price.
	DeviationPPB cciptypes.BigInt `json:"deviationPPB"`
}

// CommitOffchainConfig is the OCR offchainConfig for the commit plugin.
type CommitOffchainConfig struct {
	// RemoteGasPriceBatchWriteFrequency is the frequency at which the commit plugin should write gas prices to the remote chain.
	RemoteGasPriceBatchWriteFrequency commonconfig.Duration `json:"remoteGasPriceBatchWriteFrequency"`

	// TokenPriceBatchWriteFrequency is the frequency at which the commit plugin should write token prices to the remote chain.
	// If set to zero, no prices will be written (i.e keystone feeds would be active).
	TokenPriceBatchWriteFrequency commonconfig.Duration `json:"tokenPriceBatchWriteFrequency"`

	// PriceSources is a map of Arbitrum price sources for each token.
	// Note that the token address is that on the remote chain.
	PriceSources map[types.Account]ArbitrumPriceSource `json:"priceSources"`
}
