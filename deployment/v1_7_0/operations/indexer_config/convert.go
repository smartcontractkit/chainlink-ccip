package indexer_config

import (
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/offchain"
)

// GeneratedVerifiersToGeneratedConfig converts a slice of GeneratedVerifier
// into offchain.IndexerGeneratedConfig.
func GeneratedVerifiersToGeneratedConfig(verifiers []GeneratedVerifier) *offchain.IndexerGeneratedConfig {
	verifierSlice := make([]offchain.IndexerGeneratedVerifierConfig, 0, len(verifiers))

	for _, v := range verifiers {
		verifierSlice = append(verifierSlice, offchain.IndexerGeneratedVerifierConfig{
			Name:            v.Name,
			IssuerAddresses: v.IssuerAddresses,
		})
	}

	return &offchain.IndexerGeneratedConfig{
		Verifier: verifierSlice,
	}
}
