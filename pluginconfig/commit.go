package pluginconfig

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/smartcontractkit/chainlink-common/pkg/merklemulti"

	commonconfig "github.com/smartcontractkit/chainlink-common/pkg/config"

	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

// We use this default value when the config is not set for a specific chain.
const (
	defaultRMNSignaturesTimeout               = 5 * time.Second
	defaultNewMsgScanBatchSize                = merklemulti.MaxNumberTreeLeaves
	defaultEvmDefaultMaxMerkleTreeSize        = merklemulti.MaxNumberTreeLeaves
	defaultMaxReportTransmissionCheckAttempts = 5
	defaultRMNEnabled                         = false
	defaultRemoteGasPriceBatchWriteFrequency  = 1 * time.Minute
	defaultSignObservationPrefix              = "chainlink ccip 1.6 rmn observation"
	defaultTransmissionDelayMultiplier        = 30 * time.Second
	defaultInflightPriceCheckRetries          = 5
	defaultRelativeBoostPerWaitHour           = 0.2 // 20 percent
)

type FeeInfo struct {
	ExecDeviationPPB             cciptypes.BigInt `json:"execDeviationPPB"`
	DataAvailabilityDeviationPPB cciptypes.BigInt `json:"dataAvailabilityDeviationPPB"`
}

type TokenInfo struct {
	// AggregatorAddress is the address of the price feed TOKEN/USD aggregator on the feed chain.
	AggregatorAddress cciptypes.UnknownEncodedAddress `json:"aggregatorAddress"`

	// DeviationPPB is the deviation in parts per billion that the price feed is allowed to deviate
	// from the last written price on-chain before we write a new price.
	DeviationPPB cciptypes.BigInt `json:"deviationPPB"`

	// Decimals is the number of decimals for the token (NOT the feed).
	Decimals uint8 `json:"decimals"`
}

func (a TokenInfo) Validate() error {
	if a.AggregatorAddress == "" {
		return errors.New("aggregatorAddress not set")
	}

	// aggregator must be an ethereum address
	decoded, err := hex.DecodeString(strings.ToLower(strings.TrimPrefix(string(a.AggregatorAddress), "0x")))
	if err != nil {
		return fmt.Errorf("aggregatorAddress must be a valid ethereum address (i.e hex encoded 20 bytes): %w", err)
	}
	if len(decoded) != 20 {
		return fmt.Errorf("aggregatorAddress must be a valid ethereum address, got %d bytes expected 20", len(decoded))
	}

	if a.DeviationPPB.Int.Cmp(big.NewInt(0)) <= 0 {
		return errors.New("deviationPPB not set or negative, must be positive")
	}

	if a.Decimals == 0 {
		return fmt.Errorf("tokenDecimals can't be zero")
	}

	return nil
}

// CommitOffchainConfig is the OCR offchainConfig for the commit plugin.
// This is posted onchain as part of the OCR configuration process of the commit plugin.
// Every plugin is provided this configuration in its encoded form in the NewReportingPlugin
// method on the ReportingPluginFactory interface.
// WARN: The JSON encoding of this struct is a hard dependency for RMN.
type CommitOffchainConfig struct {
	// RemoteGasPriceBatchWriteFrequency is the frequency at which the commit plugin
	// should write gas prices to the remote chain.
	//TODO: Rename to something with ChainFee
	RemoteGasPriceBatchWriteFrequency commonconfig.Duration `json:"remoteGasPriceBatchWriteFrequency"`

	FeeInfo map[cciptypes.ChainSelector]FeeInfo `json:"feeInfo"`

	// TokenPriceBatchWriteFrequency is the frequency at which the commit plugin should
	// write token prices to the remote chain.
	// If set to zero, no prices will be written (i.e keystone feeds would be active).
	TokenPriceBatchWriteFrequency commonconfig.Duration `json:"tokenPriceBatchWriteFrequency"`

	// TokenInfo is a map of Arbitrum price sources for each token.
	// Note that the token address is that on the remote chain.
	TokenInfo map[cciptypes.UnknownEncodedAddress]TokenInfo `json:"tokenInfo"`

	// PriceFeedChainSelector is the chain selector for the chain on which
	// the token prices are read from.
	// This will typically be an arbitrum testnet/mainnet chain depending on
	// the deployment.
	PriceFeedChainSelector cciptypes.ChainSelector `json:"tokenPriceChainSelector"`

	// NewMsgScanBatchSize is the number of max new messages to scan, typically set to 256.
	NewMsgScanBatchSize int `json:"newMsgScanBatchSize"`

	// The maximum number of times to check if the previous report has been transmitted
	MaxReportTransmissionCheckAttempts uint `json:"maxReportTransmissionCheckAttempts"`

	// RMNSignaturesTimeout is the timeout for RMN signature verification.
	// Typically set to `MaxQueryDuration - e`, where e some small duration.
	RMNSignaturesTimeout time.Duration `json:"rmnSignaturesTimeout"`

	// RMNEnabled is a flag to enable/disable RMN signature verification.
	// WARN: This is a hard dependency for RMN including the json encoding of CommitOffchainConfig.
	RMNEnabled bool `json:"rmnEnabled"`

	// MaxMerkleTreeSize is the maximum size of a merkle tree to create prior to calculating the merkle root.
	// If for example in the next round we have 1000 pending messages and a max tree size of 256, only 256 seq nums
	// will be in the report. If a value is not set we fallback to EvmDefaultMaxMerkleTreeSize.
	MaxMerkleTreeSize uint64 `json:"maxTreeSize"`

	// SignObservationPrefix is the prefix used by the RMN node to sign observations.
	SignObservationPrefix string `json:"signObservationPrefix"`

	// transmissionDelayMultiplier is used to calculate the transmission delay for each oracle.
	TransmissionDelayMultiplier time.Duration `json:"transmissionDelayMultiplier"`

	// InflightPriceCheckRetries is the number of rounds we wait for a price report to get recorded on the blockchain.
	InflightPriceCheckRetries int `json:"inflightPriceCheckRetries"`
}

func (c *CommitOffchainConfig) applyDefaults() {
	if c.RMNEnabled && c.RMNSignaturesTimeout == 0 {
		c.RMNSignaturesTimeout = defaultRMNSignaturesTimeout
	}

	if c.NewMsgScanBatchSize == 0 {
		c.NewMsgScanBatchSize = defaultNewMsgScanBatchSize
	}

	if c.MaxReportTransmissionCheckAttempts == 0 {
		c.MaxReportTransmissionCheckAttempts = defaultMaxReportTransmissionCheckAttempts
	}

	if c.MaxMerkleTreeSize == 0 {
		c.MaxMerkleTreeSize = defaultEvmDefaultMaxMerkleTreeSize
	}

	if c.RemoteGasPriceBatchWriteFrequency.Duration() == 0 {
		c.RemoteGasPriceBatchWriteFrequency = *commonconfig.MustNewDuration(defaultRemoteGasPriceBatchWriteFrequency)
	}

	if c.SignObservationPrefix == "" {
		c.SignObservationPrefix = defaultSignObservationPrefix
	}

	if c.TransmissionDelayMultiplier == 0 {
		c.TransmissionDelayMultiplier = defaultTransmissionDelayMultiplier
	}

	if c.InflightPriceCheckRetries == 0 {
		c.InflightPriceCheckRetries = defaultInflightPriceCheckRetries
	}
}

func (c *CommitOffchainConfig) Validate() error {
	if c.RemoteGasPriceBatchWriteFrequency.Duration() == 0 {
		return errors.New("remoteGasPriceBatchWriteFrequency not set")
	}

	// Note that commit may not have to submit prices if keystone feeds
	// are enabled for the chain.
	// If price sources are provided the batch write frequency and token price chain selector
	// config fields MUST be provided.
	if len(c.TokenInfo) > 0 &&
		(c.TokenPriceBatchWriteFrequency.Duration() == 0 || c.PriceFeedChainSelector == 0) {
		return fmt.Errorf("tokenPriceBatchWriteFrequency (%s) or tokenPriceChainSelector (%d) not set",
			c.TokenPriceBatchWriteFrequency, c.PriceFeedChainSelector)
	}

	for token, tokenInfo := range c.TokenInfo {
		if err := tokenInfo.Validate(); err != nil {
			return fmt.Errorf("invalid token info for token %s: %w", token, err)
		}
	}

	if c.NewMsgScanBatchSize == 0 {
		return fmt.Errorf("newMsgScanBatchSize not set")
	}

	if c.RMNEnabled && c.RMNSignaturesTimeout == 0 {
		return fmt.Errorf("rmnSignaturesTimeout not set")
	}

	if c.MaxReportTransmissionCheckAttempts == 0 {
		return fmt.Errorf("maxReportTransmissionCheckAttempts not set")
	}

	if c.MaxMerkleTreeSize == 0 {
		return fmt.Errorf("maxMerkleTreeSize not set")
	}

	if c.SignObservationPrefix == "" {
		return fmt.Errorf("signObservationPrefix not set")
	}

	return nil
}

func (c *CommitOffchainConfig) ApplyDefaultsAndValidate() error {
	c.applyDefaults()
	return c.Validate()
}

// EncodeCommitOffchainConfig encodes a CommitOffchainConfig into bytes using JSON.
func EncodeCommitOffchainConfig(c CommitOffchainConfig) ([]byte, error) {
	return json.Marshal(c)
}

// DecodeCommitOffchainConfig JSON decodes a CommitOffchainConfig from bytes.
func DecodeCommitOffchainConfig(encodedCommitOffchainConfig []byte) (CommitOffchainConfig, error) {
	var c CommitOffchainConfig
	if err := json.Unmarshal(encodedCommitOffchainConfig, &c); err != nil {
		return c, err
	}
	return c, nil
}
