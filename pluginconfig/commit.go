package pluginconfig

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/smartcontractkit/chainlink-common/pkg/merklemulti"

	commonconfig "github.com/smartcontractkit/chainlink-common/pkg/config"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
)

// We use this default value when the config is not set for a specific chain.
const (
	defaultRMNSignaturesTimeout               = 5 * time.Second
	defaultNewMsgScanBatchSize                = merklemulti.MaxNumberTreeLeaves
	defaultEvmDefaultMaxMerkleTreeSize        = merklemulti.MaxNumberTreeLeaves
	defaultMaxReportTransmissionCheckAttempts = 5
	defaultRemoteGasPriceBatchWriteFrequency  = 1 * time.Minute
	defaultSignObservationPrefix              = "chainlink ccip 1.6 rmn observation"
	defaultTransmissionDelayMultiplier        = 30 * time.Second
	defaultInflightPriceCheckRetries          = 5
	defaultAsyncObserverSyncFreq              = 5 * time.Second
	defaultAsyncObserverSyncTimeout           = 10 * time.Second
)

// Deprecated: use cciptypes.TokenInfo instead.
type TokenInfo = cciptypes.TokenInfo

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

	// TransmissionDelayMultiplier is used to calculate the transmission delay for each oracle.
	TransmissionDelayMultiplier time.Duration `json:"transmissionDelayMultiplier"`

	// InflightPriceCheckRetries is the number of rounds we wait for a price report to get recorded on the blockchain.
	InflightPriceCheckRetries int `json:"inflightPriceCheckRetries"`

	// MerkleRootAsyncObserverDisabled defines whether the async observer should be disabled. Default it is enabled.
	MerkleRootAsyncObserverDisabled bool `json:"merkleRootAsyncObserverDisabled"`

	// MerkleRootAsyncObserverSyncFreq defines how frequently the async merkle roots observer should sync.
	MerkleRootAsyncObserverSyncFreq time.Duration `json:"merkleRootAsyncObserverSyncFreq"`

	// MerkleRootAsyncObserverSyncTimeout defines the timeout for a single sync operation (e.g. fetch seqNums).
	MerkleRootAsyncObserverSyncTimeout time.Duration `json:"merkleRootAsyncObserverSyncTimeout"`

	// ChainFeeAsyncObserverDisabled defines whether the async observer should be disabled. Default it is enabled.
	ChainFeeAsyncObserverDisabled bool `json:"chainFeeAsyncObserverDisabled"`

	// ChainFeeAsyncObserverSyncFreq defines how frequently the async chain fee observer should sync.
	ChainFeeAsyncObserverSyncFreq time.Duration `json:"chainFeeAsyncObserverSyncFreq"`

	// ChainFeeAsyncObserverSyncTimeout defines the timeout for a single
	// chain fee observation operation (e.g. fetch token prices).
	// NOTE: It is also used when the async observer is disabled, while making the sync calls.
	ChainFeeAsyncObserverSyncTimeout time.Duration `json:"chainFeeAsyncObserverSyncTimeout"`

	// TokenPriceAsyncObserverDisabled defines whether the async observer should be disabled. Default it is enabled.
	TokenPriceAsyncObserverDisabled bool `json:"tokenPriceAsyncObserverDisabled"`

	// TokenPriceAsyncObserverSyncFreq defines how frequently the async token price observer should sync.
	TokenPriceAsyncObserverSyncFreq commonconfig.Duration `json:"tokenPriceAsyncObserverSyncFreq"`

	// TokenPriceAsyncObserverSyncTimeout defines the timeout for a single
	// token price observation operation (e.g. fetch token prices).
	// NOTE: It is also used when the async observer is disabled, while making the sync calls.
	TokenPriceAsyncObserverSyncTimeout commonconfig.Duration `json:"tokenPriceAsyncObserverSyncTimeout"`

	// DonBreakingChangesVersion is a generic feature flag for releases that contain breaking changes for the DON.
	// Set/Increment the value iff every oracle has upgraded and is ready to use that release.
	// Example Usage:
	//    - You are adding a new Observed field, oracles running old version cannot parse it.
	//    - The new logic is used only if DonBreakingChangesVersion==1.
	//    - Release with DonBreakingChangesVersion=0 and wait until every oracle upgrades.
	//    - Now you can set it to `1` and enable your new feature.
	//    - In the next release the deprecated code can be removed.
	DonBreakingChangesVersion int `json:"donBreakingChangesVersion"`

	// MaxRootsPerReport is the maximum number of roots to include in a single report.
	// Set this to 1 for destination chains that cannot process more than one commit root per report (e.g, Solana)
	// Disable by setting to 0.
	// NOTE:
	//  * this can only be used if RMNEnabled == false.
	//  * if MaxMerkleRootsPerReport is non-zero, MultipleReportsEnabled should be set to true.
	MaxMerkleRootsPerReport uint64 `json:"maxRootsPerReport"`

	// MaxPricesPerReport is the maximum number of token and/or gas prices that may be included in a single report.
	// Price data will not be included with MerkleRoots when this value is set.
	// Disable by setting to 0.
	// NOTE:
	//  * this can only be used if RMNEnabled == false.
	//  * if MaxPricesPerReport is non-zero, MultipleReportsEnabled should be set to true.
	MaxPricesPerReport uint64 `json:"maxPricesPerReport"`

	// MultipleReportsEnabled is a flag to enable/disable multiple reports per round.
	// This is typically set to true on chains that use 'MaxMerkleRootsPerReport'
	// in order to avoid delays when there are reports from multiple sources.
	// NOTE: this can only be used if RMNEnabled == false.
	MultipleReportsEnabled bool `json:"multipleReports"`
}

const (
	// DonBreakingChangesVersion1RoleDonSupport is a release that changes the logic that oracles run to
	// generate Observation and Outcome.
	DonBreakingChangesVersion1RoleDonSupport = 1
)

//nolint:gocyclo // it is considered ok since we don't have complicated logic here
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

	// We want to apply defaults only if the async feature is enabled.
	if !c.MerkleRootAsyncObserverDisabled {
		if c.MerkleRootAsyncObserverSyncFreq == 0 {
			c.MerkleRootAsyncObserverSyncFreq = defaultAsyncObserverSyncFreq
		}
		if c.MerkleRootAsyncObserverSyncTimeout == 0 {
			c.MerkleRootAsyncObserverSyncTimeout = defaultAsyncObserverSyncTimeout
		}
	}

	if !c.ChainFeeAsyncObserverDisabled {
		if c.ChainFeeAsyncObserverSyncFreq == 0 {
			c.ChainFeeAsyncObserverSyncFreq = defaultAsyncObserverSyncFreq
		}
	}
	if c.ChainFeeAsyncObserverSyncTimeout == 0 {
		c.ChainFeeAsyncObserverSyncTimeout = defaultAsyncObserverSyncTimeout
	}

	if !c.TokenPriceAsyncObserverDisabled {
		if c.TokenPriceAsyncObserverSyncFreq.Duration() == 0 {
			c.TokenPriceAsyncObserverSyncFreq = *commonconfig.MustNewDuration(defaultAsyncObserverSyncFreq)
		}
	}
	if c.TokenPriceAsyncObserverSyncTimeout.Duration() == 0 {
		c.TokenPriceAsyncObserverSyncTimeout = *commonconfig.MustNewDuration(defaultAsyncObserverSyncTimeout)
	}
}

//nolint:gocyclo // it is considered ok since we don't have complicated logic here
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

	if !c.MerkleRootAsyncObserverDisabled &&
		(c.MerkleRootAsyncObserverSyncFreq == 0 || c.MerkleRootAsyncObserverSyncTimeout == 0) {
		return fmt.Errorf("merkle root async observer sync freq (%s) or sync timeout (%s) not set",
			c.MerkleRootAsyncObserverSyncFreq, c.MerkleRootAsyncObserverSyncTimeout)
	}

	if !c.ChainFeeAsyncObserverDisabled &&
		(c.ChainFeeAsyncObserverSyncFreq == 0 || c.ChainFeeAsyncObserverSyncTimeout == 0) {
		return fmt.Errorf("chain fee async observer sync freq (%s) or sync timeout (%s) not set",
			c.ChainFeeAsyncObserverSyncFreq, c.ChainFeeAsyncObserverSyncTimeout)
	}

	// Options for multiple reports. These settings were added so that Solana can be configured
	// to split merkle roots across multiple reports. The functions do not support RMN, so it is
	// an error to use them unless RMNEnabled == false.
	var errs []error
	if c.RMNEnabled {
		if c.MultipleReportsEnabled {
			errs = append(errs, fmt.Errorf("multipleReports do not support RMN, RMNEnabled cannot be true"))
		}
		if c.MaxMerkleRootsPerReport != 0 {
			errs = append(errs, fmt.Errorf("maxMerkleRootsPerReport does not support RMN, RMNEnabled cannot be true"))
		}
	}
	if c.MaxMerkleRootsPerReport != 0 && !c.MultipleReportsEnabled {
		errs = append(errs, fmt.Errorf("maxMerkleRootsPerReport cannot be used without MultipleReportsEnabled"))
	}
	if c.MaxPricesPerReport != 0 && !c.MultipleReportsEnabled {
		errs = append(errs, fmt.Errorf("maxPricesPerReport cannot be used without MultipleReportsEnabled"))
	}

	if len(errs) > 0 {
		return errors.Join(errs...)
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
