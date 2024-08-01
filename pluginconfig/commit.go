package pluginconfig

import (
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"strings"
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

// ArbitrumPriceSource is the source of the TOKEN/USD price data of a particular token
// on Arbitrum.
type ArbitrumPriceSource struct {
	// AggregatorAddress is the address of the price feed TOKEN/USD aggregator on arbitrum.
	AggregatorAddress string `json:"aggregatorAddress"`

	// DeviationPPB is the deviation in parts per billion that the price feed is allowed to deviate
	// from the last written price on-chain before we write a new price.
	DeviationPPB cciptypes.BigInt `json:"deviationPPB"`
}

func (a ArbitrumPriceSource) Validate() error {
	if a.AggregatorAddress == "" {
		return errors.New("aggregatorAddress not set")
	}

	// aggregator must be an ethereum address
	decoded, err := hex.DecodeString(strings.ToLower(strings.TrimPrefix(a.AggregatorAddress, "0x")))
	if err != nil {
		return fmt.Errorf("aggregatorAddress must be a valid ethereum address (i.e hex encoded 20 bytes): %w", err)
	}
	if len(decoded) != 20 {
		return fmt.Errorf("aggregatorAddress must be a valid ethereum address, got %d bytes expected 20", len(decoded))
	}

	if a.DeviationPPB.Int.Cmp(big.NewInt(0)) <= 0 {
		return errors.New("deviationPPB not set or negative, must be positive")
	}

	return nil
}

// CommitOffchainConfig is the OCR offchainConfig for the commit plugin.
type CommitOffchainConfig struct {
	// RemoteGasPriceBatchWriteFrequency is the frequency at which the commit plugin
	// should write gas prices to the remote chain.
	RemoteGasPriceBatchWriteFrequency commonconfig.Duration `json:"remoteGasPriceBatchWriteFrequency"`

	// TokenPriceBatchWriteFrequency is the frequency at which the commit plugin should
	// write token prices to the remote chain.
	// If set to zero, no prices will be written (i.e keystone feeds would be active).
	TokenPriceBatchWriteFrequency commonconfig.Duration `json:"tokenPriceBatchWriteFrequency"`

	// PriceSources is a map of Arbitrum price sources for each token.
	// Note that the token address is that on the remote chain.
	PriceSources map[types.Account]ArbitrumPriceSource `json:"priceSources"`
}

func (c CommitOffchainConfig) Validate() error {
	if c.RemoteGasPriceBatchWriteFrequency.Duration() == 0 {
		return errors.New("remoteGasPriceBatchWriteFrequency not set")
	}

	// Note that commit may not have to submit prices if keystone feeds
	// are enabled for the chain.
	// If neither frequency nor price sources are set, then the node will
	// not submit token prices.
	if c.TokenPriceBatchWriteFrequency.Duration() != 0 {
		if len(c.PriceSources) == 0 {
			return errors.New("tokenPriceBatchWriteFrequency set but no price sources provided")
		}
		for _, priceSource := range c.PriceSources {
			if err := priceSource.Validate(); err != nil {
				return fmt.Errorf("price source validation failed on %+v: %w", priceSource, err)
			}
		}
	} else if len(c.PriceSources) != 0 {
		return errors.New("price sources provided but tokenPriceBatchWriteFrequency not set")
	}

	return nil
}
