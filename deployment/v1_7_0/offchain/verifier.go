package offchain

import (
	"fmt"
)

// VerifierCommitConfig is the committee verifier job spec configuration.
type VerifierCommitConfig struct {
	VerifierID        string `toml:"verifier_id"`
	AggregatorAddress string `toml:"aggregator_address"`
	InsecureAggregatorConnection bool `toml:"insecure_aggregator_connection"`

	SignerAddress string `toml:"signer_address"`

	PyroscopeURL string `toml:"pyroscope_url"`
	CommitteeVerifierAddresses map[string]string `toml:"committee_verifier_addresses"`
	OnRampAddresses            map[string]string `toml:"on_ramp_addresses"`
	DefaultExecutorOnRampAddresses map[string]string `toml:"default_executor_on_ramp_addresses"`
	RMNRemoteAddresses         map[string]string `toml:"rmn_remote_addresses"`
	DisableFinalityCheckers    []string          `toml:"disable_finality_checkers"`
	Monitoring                 MonitoringConfig  `toml:"monitoring"`
}

// Validate validates the verifier commit configuration.
func (c *VerifierCommitConfig) Validate() error {
	if len(c.OnRampAddresses) != len(c.CommitteeVerifierAddresses) ||
		len(c.OnRampAddresses) != len(c.RMNRemoteAddresses) {
		return fmt.Errorf(
			"invalid verifier configuration, mismatched lengths for onramp (%d), committee verifier (%d), and RMN Remote addresses (%d)",
			len(c.OnRampAddresses),
			len(c.CommitteeVerifierAddresses),
			len(c.RMNRemoteAddresses),
		)
	}

	for k := range c.OnRampAddresses {
		if _, ok := c.CommitteeVerifierAddresses[k]; !ok {
			return fmt.Errorf("invalid verifier configuration, chain selector in onramp (%s) not in committee verifier addresses", k)
		}
		if _, ok := c.RMNRemoteAddresses[k]; !ok {
			return fmt.Errorf("invalid verifier configuration, chain selector in onramp (%s) not in RMN Remote addresses", k)
		}
	}

	return nil
}
