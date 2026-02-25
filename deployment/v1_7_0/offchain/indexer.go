package offchain

// IndexerGeneratedVerifierConfig contains auto-generated verifier configuration
// for the indexer service.
type IndexerGeneratedVerifierConfig struct {
	Name            string   `toml:"Name"`
	IssuerAddresses []string `toml:"IssuerAddresses"`
}

// IndexerGeneratedConfig contains auto-generated deployment configuration
// for the indexer service.
type IndexerGeneratedConfig struct {
	Verifier []IndexerGeneratedVerifierConfig `toml:"Verifier"`
}
