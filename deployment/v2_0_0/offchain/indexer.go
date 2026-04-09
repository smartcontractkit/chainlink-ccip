package offchain

type IndexerVerifierConfig struct {
	Name            string   `json:"name"`
	IssuerAddresses []string `json:"issuerAddresses"`
}

type IndexerGeneratedConfig struct {
	Verifiers []IndexerVerifierConfig `json:"verifiers"`
}
