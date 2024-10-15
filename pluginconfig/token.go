package pluginconfig

import (
	"errors"
	"time"

	commonconfig "github.com/smartcontractkit/chainlink-common/pkg/config"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
)

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
	if t.IsUSDC() {
		if t.USDCCCTPObserverConfig == nil {
			return errors.New("USDCCCTPObserverConfig is empty")
		}
		return nil
	}
	return errors.New("unknown token data observer type")
}

// Validate checks that the observer's config is semantically correct - fields are set correctly
// depending on the config's type
func (t TokenDataObserverConfig) Validate() error {
	if t.IsUSDC() {
		return t.USDCCCTPObserverConfig.Validate()
	}
	return errors.New("unknown token data observer type " + t.Type)
}

func (t TokenDataObserverConfig) IsUSDC() bool {
	return t.Type == USDCCCTPHandlerType
}

const USDCCCTPHandlerType = "usdc-cctp"

type USDCCCTPObserverConfig struct {
	Tokens         map[cciptypes.ChainSelector]USDCCCTPTokenConfig `json:"tokens"`
	AttestationAPI string                                          `json:"attestationAPI"`
	// AttestationAPITimeout defines the timeout for the attestation API.
	AttestationAPITimeout *commonconfig.Duration `json:"attestationAPITimeout"`
	// AttestationAPIInterval defines the rate in requests per second that the attestation API can be called.
	// Default set according to the APIs documentated 10 requests per second rate limit.
	AttestationAPIInterval *commonconfig.Duration `json:"attestationAPIInterval"`
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

	// Default to 100 millis if AttestationAPIInterval is not set this is set according to the APIs documented
	// 10 requests per second rate limit.
	if p.AttestationAPIInterval == nil {
		p.AttestationAPIInterval = commonconfig.MustNewDuration(100 * time.Millisecond)
	}
}

//nolint:lll // CCTP link
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
