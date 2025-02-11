package pluginconfig

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	commonconfig "github.com/smartcontractkit/chainlink-common/pkg/config"
)

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
	// Deprecated: this is replaced by MaxReportMessages and MaxSingleChainReports.
	BatchingStrategyID uint32 `json:"batchingStrategyID"`

	// TokenDataObservers registers different strategies for processing token data.
	TokenDataObservers []TokenDataObserverConfig `json:"tokenDataObservers"`

	// TransmissionDelayMultiplier is used to calculate the transmission delay for each oracle.
	TransmissionDelayMultiplier time.Duration `json:"transmissionDelayMultiplier"`

	// MaxReportMessages is the maximum number of messages that can be included in a report.
	// When set to 0, this setting is ignored.
	MaxReportMessages uint64 `json:"maxReportMessages"`

	// MaxSingleChainReports is the maximum number of single chain reports that can be included in a report.
	// When set to 0, this setting is ignored.
	MaxSingleChainReports uint64 `json:"maxSingleChainReports"`

	// MaxTxmStatusChecks is the maximum number of times to check the transaction manager status for a fatal
	// transaction result. When set to 0 Txm status checking is disabled.
	MaxTxmStatusChecks uint64 `json:"maxTxmStatusChecks"`
}

func (e *ExecuteOffchainConfig) ApplyDefaultsAndValidate() error {
	e.applyDefaults()
	return e.Validate()
}

func (e *ExecuteOffchainConfig) applyDefaults() {
	if e.TransmissionDelayMultiplier == 0 {
		e.TransmissionDelayMultiplier = defaultTransmissionDelayMultiplier
	}
	if e.RelativeBoostPerWaitHour == 0 {
		e.RelativeBoostPerWaitHour = defaultRelativeBoostPerWaitHour
	}
}

func (e *ExecuteOffchainConfig) Validate() error {
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

	if e.RelativeBoostPerWaitHour > 1 || e.RelativeBoostPerWaitHour < 0 {
		return fmt.Errorf("RelativeBoostPerWaitHour must be <= 1 and >= 0, got: %f", e.RelativeBoostPerWaitHour)
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

func (e *ExecuteOffchainConfig) IsUSDCEnabled() bool {
	for _, ob := range e.TokenDataObservers {
		if ob.WellFormed() != nil {
			continue
		}
		if ob.IsUSDC() {
			return true
		}
	}
	return false

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
