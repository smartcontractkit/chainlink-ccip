package pluginconfig

import (
	"errors"
	"time"

	commonconfig "github.com/smartcontractkit/chainlink-common/pkg/config"

	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

const (
	USDCCCTPHandlerType = "usdc-cctp"
	LBTCHandlerType     = "lbtc"
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
	*LBTCObserverConfig
}

// WellFormed checks that the observer's config is syntactically correct - proper struct is initialized based on type
func (t TokenDataObserverConfig) WellFormed() error {
	if t.IsUSDC() {
		if t.USDCCCTPObserverConfig == nil {
			return errors.New("USDCCCTPObserverConfig is empty")
		}
		return nil
	}
	if t.IsLBTC() {
		if t.LBTCObserverConfig == nil {
			return errors.New("LBTCObserverConfig is empty")
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
	if t.IsLBTC() {
		return t.LBTCObserverConfig.Validate()
	}
	return errors.New("unknown token data observer type " + t.Type)
}

func (t TokenDataObserverConfig) IsUSDC() bool {
	return t.Type == USDCCCTPHandlerType
}

func (t TokenDataObserverConfig) IsLBTC() bool {
	return t.Type == LBTCHandlerType
}

type USDCCCTPObserverConfig struct {
	Tokens map[cciptypes.ChainSelector]USDCCCTPTokenConfig `json:"tokens"`

	AttestationAPI string `json:"attestationAPI"`
	// AttestationAPITimeout defines the timeout for the attestation API.
	AttestationAPITimeout *commonconfig.Duration `json:"attestationAPITimeout"`
	// AttestationAPIInterval defines the rate in requests per second that the attestation API can be called.
	// Default set according to the APIs documentated 10 requests per second rate limit.
	AttestationAPIInterval *commonconfig.Duration `json:"attestationAPIInterval"`
	// NumWorkers is the number of concurrent workers.
	NumWorkers int `json:"numWorkers"`
	// CacheExpirationInterval is the interval after which the cached token data will expire.
	CacheExpirationInterval *commonconfig.Duration `json:"cacheExpirationInterval"`
	// CacheCleanupInterval is the interval after which the cache expired data will be cleaned up.
	CacheCleanupInterval *commonconfig.Duration `json:"cacheCleanupInterval"`
	// ObserveTimeout is the timeout for the actual synchronous Observe calls.
	ObserveTimeout *commonconfig.Duration `json:"observeTimeout"`
}

func (c *USDCCCTPObserverConfig) setDefaults() {
	// Default to 10 minutes if CacheExpirationInterval is not set
	if c.CacheExpirationInterval == nil {
		c.CacheExpirationInterval = commonconfig.MustNewDuration(10 * time.Minute)
	}
	// Default to 15 minutes if CacheCleanupInterval is not set
	if c.CacheCleanupInterval == nil {
		c.CacheCleanupInterval = commonconfig.MustNewDuration(15 * time.Minute)
	}
	// Default to 5 seconds if ObserveTimeout is not set
	if c.ObserveTimeout == nil {
		c.ObserveTimeout = commonconfig.MustNewDuration(5 * time.Second)
	}
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

func (c *USDCCCTPObserverConfig) Validate() error {
	if len(p.Tokens) == 0 {
		return errors.New("tokens not set")
	}
	for _, token := range p.Tokens {
		if err := token.Validate(); err != nil {
			return err
		}
	}
	if p.AttestationAPI == "" {
		return errors.New("AttestationAPI not set")
	}
	if p.AttestationAPIInterval == nil || p.AttestationAPIInterval.Duration() == 0 {
		return errors.New("AttestationAPIInterval not set")
	}
	if p.AttestationAPITimeout == nil || p.AttestationAPITimeout.Duration() == 0 {
		return errors.New("AttestationAPITimeout not set")
	}
	if err != nil {
		return err
	}
	if c.CacheExpirationInterval == nil || c.CacheExpirationInterval.Duration() == 0 {
		return errors.New("CacheExpirationInterval not set")
	}
	if c.CacheCleanupInterval == nil || c.CacheCleanupInterval.Duration() == 0 {
		return errors.New("CacheCleanupInterval not set")
	}
	if c.ObserveTimeout == nil || c.ObserveTimeout.Duration() == 0 {
		return errors.New("ObserveTimeout not set")
	}
	return nil
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

type LBTCObserverConfig struct {
	SourcePoolAddressByChain map[cciptypes.ChainSelector]string `json:"sourcePoolAddressByChain"`

	AttestationAPI string `json:"attestationAPI"`
	// AttestationAPITimeout defines the timeout for the attestation API.
	AttestationAPITimeout *commonconfig.Duration `json:"attestationAPITimeout"`
	// AttestationAPIInterval defines the rate in requests per second that the attestation API can be called.
	// Default set according to the APIs documentated 10 requests per second rate limit.
	AttestationAPIInterval *commonconfig.Duration `json:"attestationAPIInterval"`
	// AttestationAPITimeout defines size of the batch that can be made in one API request
	AttestationAPIBatchSize int `json:"attestationAPIBatchSize"`
	// NumWorkers is the number of concurrent workers.
	NumWorkers int `json:"numWorkers"`
	// CacheExpirationInterval is the interval after which the cached token data will expire.
	CacheExpirationInterval *commonconfig.Duration `json:"cacheExpirationInterval"`
	// CacheCleanupInterval is the interval after which the cache expired data will be cleaned up.
	CacheCleanupInterval *commonconfig.Duration `json:"cacheCleanupInterval"`
	// ObserveTimeout is the timeout for the actual synchronous Observe calls.
	ObserveTimeout *commonconfig.Duration `json:"observeTimeout"`
}

func (c *LBTCObserverConfig) setDefaults() {
	// Default to 10 minutes if CacheExpirationInterval is not set
	if c.CacheExpirationInterval == nil {
		c.CacheExpirationInterval = commonconfig.MustNewDuration(10 * time.Minute)
	}
	// Default to 15 minutes if CacheCleanupInterval is not set
	if c.CacheCleanupInterval == nil {
		c.CacheCleanupInterval = commonconfig.MustNewDuration(15 * time.Minute)
	}
	// Default to 5 seconds if ObserveTimeout is not set
	if c.ObserveTimeout == nil {
		c.ObserveTimeout = commonconfig.MustNewDuration(5 * time.Second)
	}
	// Default to 1 second if AttestationAPITimeout is not set
	if p.AttestationAPITimeout == nil {
		p.AttestationAPITimeout = commonconfig.MustNewDuration(5 * time.Second)
	}
	// Default to 100 millis if AttestationAPIInterval is not set this is set according to the APIs documented
	// 10 requests per second rate limit.
	if p.AttestationAPIInterval == nil {
		p.AttestationAPIInterval = commonconfig.MustNewDuration(100 * time.Millisecond)
	}
	if c.AttestationAPIBatchSize == 0 {
		c.AttestationAPIBatchSize = 50
	}
}

func (c *LBTCObserverConfig) Validate() error {
	c.setDefaults()
	if p.AttestationAPI == "" {
		return errors.New("AttestationAPI not set")
	}
	if p.AttestationAPIInterval == nil || p.AttestationAPIInterval.Duration() == 0 {
		return errors.New("AttestationAPIInterval not set")
	}
	if p.AttestationAPITimeout == nil || p.AttestationAPITimeout.Duration() == 0 {
		return errors.New("AttestationAPITimeout not set")
	}
	if c.CacheExpirationInterval == nil || c.CacheExpirationInterval.Duration() == 0 {
		return errors.New("CacheExpirationInterval not set")
	}
	if c.CacheCleanupInterval == nil || c.CacheCleanupInterval.Duration() == 0 {
		return errors.New("CacheCleanupInterval not set")
	}
	if c.ObserveTimeout == nil || c.ObserveTimeout.Duration() == 0 {
		return errors.New("ObserveTimeout not set")
	}
	if c.AttestationAPIBatchSize == 0 {
		return errors.New("AttestationAPIBatchSize is not set")
	}
	if len(c.SourcePoolAddressByChain) == 0 {
		return errors.New("SourcePoolAddressByChain is not set")
	}
	for _, sourcePoolAddress := range c.SourcePoolAddressByChain {
		if sourcePoolAddress == "" {
			return errors.New("SourcePoolAddressByChain is empty")
		}
	}
	return nil
}
