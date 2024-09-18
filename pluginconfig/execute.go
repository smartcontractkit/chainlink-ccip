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

	// TokenDataObservers registers different strategies for processing token data.
	// TokenDataObservers registers different strategies for processing token data.
	TokenDataObservers []TokenDataObserverConfig `json:"tokenDataObservers"`
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

	set := make(map[string]struct{})
	for _, ob := range e.TokenDataObservers {
		if err := ob.Validate(); err != nil {
			return err
		}

		key := ob.Type + ob.Version
		if _, exists := set[key]; exists {
			return errors.New("duplicate token data observer type and version")
		}
		set[key] = struct{}{}
	}
	return nil
}

// TokenDataObserverConfig is the base struct for token data observers. Every token data observer
// has to define its type and version. The type and version is used to determine which observer's
// implementation to use. Whenever you want to add a new observer type, you need to add a new struct and embed that
// in the TokenDataObserverConfig similarly to how USDCCCTPObserverConfig is embedded in the TokenDataObserverConfig.
// There are two additional checks for the TokenDataObserverConfig to enforce that it's semantically (Validate)
// and syntactically correct (WellFormed).
type TokenDataObserverConfig struct {
	// Type is the type of the token data observer. You can think of different token data observers as different
	// strategies for processing token data. For example, you can have a token data observer for USDC tokens using CCTP
	// and different one for processing LINK token.
	Type string `json:"type"`
	// Version is the version of the TokenObserverConfig and the matching Observer implementation for that config.
	// This is used to determine which version of the observer to use. Right now, we only have one version
	// of the observer, but in the future, we might have multiple versions.
	// This is a precautionary measure to ensure that we can upgrade the observer without breaking the existing ones.
	// Example would be CCTPv1 using AttestationAPI and CCTPv2 using a different API or completely
	// different strategy which requires different configuration and implementation during Observation phase.
	// [
	//  {
	//    "type": "usdc-cctp",
	//    "version": "1.0",
	//    "attestationAPI": "http://circle.com/attestation",
	//    "attestationAPITimeout": "1s",
	//    "attestationAPIIntervalMilliseconds": "500ms"
	//  },
	//  {
	//    "type": "usdc-cctp",
	//    "version": "2.0",
	//    "customCirlceAPI": "http://cirle.com/gohere",
	//    "yetAnotherAPI": "http://cirle.com/anotherone",
	//    "customCircleAPITimeout": "1s",
	//    "yetAnotherAPITimeout": "500ms"
	//  }
	//]
	// Having version in that JSON isn't expensive, but it could reduce the risk of breaking the observers in the future.
	Version string `json:"version"`

	*USDCCCTPObserverConfig
}

// WellFormed checks that the observer's config is syntactically correct - proper struct is initialized based on type
func (t TokenDataObserverConfig) WellFormed() error {
	switch t.Type {
	case USDCCCTPHandlerType:
		if t.USDCCCTPObserverConfig == nil {
			return errors.New("USDCCCTPObserverConfig is empty")
		}
	default:
		return errors.New("unknown token data observer type")
	}
	return nil
}

// Validate checks that the observer's config is semantically correct - fields are set correctly
// depending on the config's type
func (t TokenDataObserverConfig) Validate() error {
	switch t.Type {
	case USDCCCTPHandlerType:
		if err := t.USDCCCTPObserverConfig.Validate(); err != nil {
			return err
		}
	default:
		return errors.New("unknown token data observer type " + t.Type)
	}
	return nil
}

const USDCCCTPHandlerType = "usdc-cctp"

type USDCCCTPObserverConfig struct {
	Tokens                 map[cciptypes.ChainSelector]USDCCCTPTokenConfig `json:"tokens"`
	AttestationAPI         string                                          `json:"attestationAPI"`
	AttestationAPITimeout  *commonconfig.Duration                          `json:"attestationAPITimeout"`
	AttestationAPIInterval *commonconfig.Duration                          `json:"attestationAPIInterval"`
}

func (p *USDCCCTPObserverConfig) Validate() error {
	p.setDefaults()

	if p.AttestationAPI == "" {
		return errors.New("AttestationAPI not set")
	}
	if len(p.Tokens) == 0 {
		return errors.New("Tokens not set")
	}
	if p.AttestationAPIInterval == nil || p.AttestationAPIInterval.Duration() == 0 {
		return errors.New("AttestationAPIInterval not set")
	}
	if p.AttestationAPITimeout == nil || p.AttestationAPITimeout.Duration() == 0 {
		return errors.New("AttestationAPITimeout not set")
	}
	for _, token := range p.Tokens {
		if err := token.Validate(); err != nil {
			return err
		}
	}
	return nil
}

func (p *USDCCCTPObserverConfig) setDefaults() {
	// Default to 1 second if AttestationAPITimeout is not set
	if p.AttestationAPITimeout == nil {
		p.AttestationAPITimeout = commonconfig.MustNewDuration(5 * time.Second)
	}

	// Default to 1 second if AttestationAPIInterval is not set
	if p.AttestationAPIInterval == nil {
		p.AttestationAPIInterval = commonconfig.MustNewDuration(100 * time.Millisecond)
	}
}

// nolint:lll // CCTP link
type USDCCCTPTokenConfig struct {
	// SourcePoolAddress is the address of the USDC token pool on the source chain that support USDC token transfers
	SourcePoolAddress string `json:"sourceTokenAddress"`
	// SourceMessageTransmitterAddr is the address of the CCTP MessageTransmitter address on the source chain
	// https://github.com/circlefin/evm-cctp-contracts/blob/adb2a382b09ea574f4d18d8af5b6706e8ed9b8f2/src/MessageTransmitter.sol
	SourceMessageTransmitterAddr string `json:"sourceMessageTransmitterAddress"`
}

func (t USDCCCTPTokenConfig) Validate() error {
	if t.SourcePoolAddress == "" {
		return errors.New("SourcePoolAddress not set")
	}
	if t.SourceMessageTransmitterAddr == "" {
		return errors.New("SourceMessageTransmitterAddress not set")
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

	for _, ob := range e.TokenDataObservers {
		if err := ob.WellFormed(); err != nil {
			return e, err
		}
	}

	return e, nil
}
