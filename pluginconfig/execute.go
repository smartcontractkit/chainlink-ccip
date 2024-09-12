package pluginconfig

import (
	"encoding/json"
	"errors"
	"time"

	commonconfig "github.com/smartcontractkit/chainlink-common/pkg/config"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
)

// ExecutePluginConfig is configuration for the execute plugin
// which includes the offchain configuration as well as other parameters
// fetched from the OCR configuration.
type ExecutePluginConfig struct {
	// DestChain is the ccip destination chain configured for the execute DON.
	DestChain cciptypes.ChainSelector `json:"destChain"`

	// SyncTimeout is the timeout for syncing the exec plugin ccip reader.
	SyncTimeout time.Duration `json:"syncTimeout"`

	// SyncFrequency is the frequency at which the exec plugin ccip reader should sync.
	SyncFrequency time.Duration `json:"syncFrequency"`

	// OffchainConfig is the offchain config set for the exec DON.
	OffchainConfig ExecuteOffchainConfig `json:"offchainConfig"`
}

// ExecuteOffchainConfig is the OCR offchainConfig for the exec plugin.
// This is posted onchain as part of the OCR configuration process of the exec plugin.
// Every plugin is provided this configuration in its encoded form in the NewReportingPlugin
// method on the ReportingPluginFactory interface.
type ExecuteOffchainConfig struct {
	// BatchGasLimit is the maximum sum of user callback gas we permit in one execution report.
	// EVM only.
	BatchGasLimit uint64 `json:"batchGasLimit"`

	// RelativeBoostPerWaitHour indicates how much to increase (artificially) the fee paid on the source chain per hour
	// of wait time, such that eventually the fee paid is greater than the execution cost, and we’ll execute it.
	// For example: if set to 0.5, that means the fee paid is increased by 50% every hour the message has been waiting.
	RelativeBoostPerWaitHour float64 `json:"relativeBoostPerWaitHour"`

	// InflightCacheExpiry indicates how long we keep a report in the plugin cache before we expire it.
	// The caching prevents us from issuing another report while one is already in flight.
	InflightCacheExpiry commonconfig.Duration `json:"inflightCacheExpiry"`

	// RootSnoozeTime is the interval at which we check roots for executable messages.
	RootSnoozeTime commonconfig.Duration `json:"rootSnoozeTime"`

	// MessageVisibilityInterval is the time interval for which the messages are visible by the plugin.
	MessageVisibilityInterval commonconfig.Duration `json:"messageVisibilityInterval"`

	// BatchingStrategyID is the strategy to use for batching messages.
	BatchingStrategyID uint32 `json:"batchingStrategyID"`

	// RawTokenDataProcessors registers different strategies for processing token data.
	RawTokenDataProcessors []json.RawMessage `json:"tokenDataProcessors"`

	// TokenDataProcessors is the parsed version of RawTokenDataProcessors.
	TokenDataProcessors []interface{}
}

type TokenDataProcessor struct {
	Type    string `json:"type"`
	Version string `json:"version"`
}

const UsdcProcessorType = "usdc-cctp"

type UsdcCctpTokenDataProcessor struct {
	TokenDataProcessor
	Tokens                       []UsdcCctpToken        `json:"tokens"`
	AttestationAPI               string                 `json:"attestationAPI"`
	AttestationAPITimeout        *commonconfig.Duration `json:"attestationAPITimeout"`
	AttestationAPIIntervalMillis *commonconfig.Duration `json:"attestationAPIIntervalMilliseconds"`
}

func (p UsdcCctpTokenDataProcessor) Validate() error {
	if p.AttestationAPI == "" {
		return errors.New("AttestationAPI not set")
	}
	if len(p.Tokens) == 0 {
		return errors.New("Tokens not set")
	}
	for _, token := range p.Tokens {
		if err := token.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UsdcCctpToken struct {
	SourceChainSelector          int    `json:"sourceChainSelector"`
	SourceTokenAddress           string `json:"sourceTokenAddress"`
	SourceMessageTransmitterAddr string `json:"sourceMessageTransmitterAddress"`
}

func (t UsdcCctpToken) Validate() error {
	if t.SourceChainSelector == 0 {
		return errors.New("SourceChainSelector not set")
	}
	if t.SourceTokenAddress == "" {
		return errors.New("SourceTokenAddress not set")
	}
	if t.SourceMessageTransmitterAddr == "" {
		return errors.New("SourceMessageTransmitterAddress not set")
	}
	return nil
}

func (e ExecuteOffchainConfig) Validate() error {
	// TODO: this doesn't really make much sense for non-EVM chains.
	// Maybe we need to have a field in the config that is not JSON-encoded
	// that indicates chain family?
	if e.BatchGasLimit == 0 {
		return errors.New("BatchGasLimit not set")
	}

	if e.RelativeBoostPerWaitHour == 0 {
		return errors.New("RelativeBoostPerWaitHour not set")
	}

	if e.InflightCacheExpiry.Duration() == 0 {
		return errors.New("InflightCacheExpiry not set")
	}

	if e.RootSnoozeTime.Duration() == 0 {
		return errors.New("RootSnoozeTime not set")
	}

	if e.MessageVisibilityInterval.Duration() == 0 {
		return errors.New("MessageVisibilityInterval not set")
	}

	for _, processor := range e.TokenDataProcessors {
		switch casted := processor.(type) {
		case UsdcCctpTokenDataProcessor:
			if err := casted.Validate(); err != nil {
				return err
			}
		default:
			return errors.New("unknown token data processor type")
		}
	}
	return nil
}

// EncodeExecuteOffchainConfig encodes a ExecuteOffchainConfig into bytes using JSON.
func EncodeExecuteOffchainConfig(e ExecuteOffchainConfig) ([]byte, error) {
	return json.Marshal(e)
}

// DecodeExecuteOffchainConfig JSON decodes a ExecuteOffchainConfig from bytes.
func DecodeExecuteOffchainConfig(encodedExecuteOffchainConfig []byte) (ExecuteOffchainConfig, error) {
	var e ExecuteOffchainConfig
	if err := json.Unmarshal(encodedExecuteOffchainConfig, &e); err != nil {
		return e, err
	}

	if len(e.RawTokenDataProcessors) == 0 {
		return e, nil
	}

	tokenDataProcessors := make([]interface{}, len(e.RawTokenDataProcessors))
	for i, processor := range e.RawTokenDataProcessors {
		var baseProcessor TokenDataProcessor
		if err := json.Unmarshal(processor, &baseProcessor); err != nil {
			return e, err
		}
		switch baseProcessor.Type {
		case UsdcProcessorType:
			var usdcProcessor UsdcCctpTokenDataProcessor
			if err := json.Unmarshal(processor, &usdcProcessor); err != nil {
				return e, err
			}
			tokenDataProcessors[i] = usdcProcessor
		default:
			return e, errors.New("unknown token data processor type")
		}
	}
	e.TokenDataProcessors = tokenDataProcessors

	return e, nil
}
