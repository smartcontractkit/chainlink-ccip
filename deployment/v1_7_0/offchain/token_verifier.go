package offchain

import (
	"encoding/hex"
	"fmt"
	"strings"
	"time"
)

// ByteSlice is a wrapper around []byte that marshals/unmarshals to/from hex
// instead of base64. Copied from chainlink-ccv/protocol.ByteSlice.
type ByteSlice []byte

func NewByteSliceFromHex(s string) (ByteSlice, error) {
	if s == "" {
		return ByteSlice{}, nil
	}

	s = strings.TrimPrefix(s, "0x")

	b, err := hex.DecodeString(s)
	if err != nil {
		return nil, fmt.Errorf("invalid hex string: %w", err)
	}

	return ByteSlice(b), nil
}

func (h ByteSlice) MarshalJSON() ([]byte, error) {
	if h == nil {
		return []byte("null"), nil
	}
	if len(h) == 0 {
		return []byte(`"0x"`), nil
	}
	return fmt.Appendf(nil, `"0x%s"`, hex.EncodeToString(h)), nil
}

func (h *ByteSlice) UnmarshalJSON(data []byte) error {
	v := string(data)

	if v == "null" {
		*h = nil
		return nil
	}

	if len(v) < 2 {
		return fmt.Errorf("invalid ByteSlice: %s", v)
	}

	if v[0] != '"' || v[len(v)-1] != '"' {
		return fmt.Errorf("invalid JSON string format for ByteSlice: %s", v)
	}

	v = v[1 : len(v)-1]

	if v == "" || v == "0x" {
		*h = ByteSlice{}
		return nil
	}

	v = strings.TrimPrefix(v, "0x")

	b, err := hex.DecodeString(v)
	if err != nil {
		return fmt.Errorf("failed to decode hex: %w", err)
	}

	*h = ByteSlice(b)
	return nil
}

func (h ByteSlice) String() string {
	if len(h) == 0 {
		return "0x"
	}
	return "0x" + hex.EncodeToString(h)
}

func mustDecodeHex(s string) ByteSlice {
	b, err := NewByteSliceFromHex(s)
	if err != nil {
		panic(err)
	}
	return b
}

// CCTPDefaultVerifierVersionHex is bytes4(keccak256("CCTPVerifier 1.7.0")).
const CCTPDefaultVerifierVersionHex = "0x8e1d1a9d"

// CCTPDefaultVerifierVersion is the default version of the 1.7 CCTPVerifier contract.
var CCTPDefaultVerifierVersion = mustDecodeHex(CCTPDefaultVerifierVersionHex)

// LombardDefaultVerifierVersionHex is bytes4(keccak256("LombardVerifier 1.7.0")).
const LombardDefaultVerifierVersionHex = "0xf0f3a135"

// LombardDefaultVerifierVersion is the default version of the 1.7 LombardVerifier contract.
var LombardDefaultVerifierVersion = mustDecodeHex(LombardDefaultVerifierVersionHex)

// CCTPConfig holds CCTP token verifier configuration.
// Deployment-only copy; runtime-parsed fields are omitted.
type CCTPConfig struct {
	AttestationAPI         string        `json:"attestation_api"          toml:"attestation_api"`
	AttestationAPITimeout  time.Duration `json:"attestation_api_timeout"  toml:"attestation_api_timeout"`
	AttestationAPIInterval time.Duration `json:"attestation_api_interval" toml:"attestation_api_interval"`
	AttestationAPICooldown time.Duration `json:"attestation_api_cooldown" toml:"attestation_api_cooldown"`
	VerifierVersion        ByteSlice     `json:"verifier_version"         toml:"verifier_version"`
	Verifiers              map[string]any `json:"verifier_addresses"          toml:"verifier_addresses"`
	VerifierResolvers      map[string]any `json:"verifier_resolver_addresses" toml:"verifier_resolver_addresses"`
}

// LombardConfig holds Lombard token verifier configuration.
// Deployment-only copy; runtime-parsed fields are omitted.
type LombardConfig struct {
	AttestationAPI          string        `json:"attestation_api"            toml:"attestation_api"`
	AttestationAPITimeout   time.Duration `json:"attestation_api_timeout"    toml:"attestation_api_timeout"`
	AttestationAPIInterval  time.Duration `json:"attestation_api_interval"   toml:"attestation_api_interval"`
	AttestationAPIBatchSize int           `json:"attestation_api_batch_size" toml:"attestation_api_batch_size"`
	VerifierVersion         ByteSlice     `json:"verifier_version"           toml:"verifier_version"`
	VerifierResolvers       map[string]any `json:"verifier_resolver_addresses" toml:"verifier_resolver_addresses"`
}

// TokenVerifierPluginConfig describes a single token verifier plugin instance.
type TokenVerifierPluginConfig struct {
	VerifierID string `json:"verifier_id" toml:"verifier_id"`
	Type       string `json:"type"        toml:"type"`
	Version    string `json:"version"     toml:"version"`

	*CCTPConfig    `json:"cctp_config,omitempty"`
	*LombardConfig `json:"lombard_config,omitempty"`
}

// TokenVerifierConfig is the root configuration for the token verifier service.
type TokenVerifierConfig struct {
	PyroscopeURL       string                      `json:"pyroscope_url"       toml:"pyroscope_url"`
	OnRampAddresses    map[string]string            `json:"on_ramp_addresses"   toml:"on_ramp_addresses"`
	RMNRemoteAddresses map[string]string            `json:"rmn_remote_addresses" toml:"rmn_remote_addresses"`
	TokenVerifiers     []TokenVerifierPluginConfig  `json:"token_verifiers"     toml:"token_verifiers"`
	Monitoring         MonitoringConfig             `json:"monitoring"          toml:"monitoring"`
}
