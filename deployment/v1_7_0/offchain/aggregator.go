package offchain

// SourceSelector represents a source chain selector as a string.
type SourceSelector = string

// DestinationSelector represents a destination chain selector as a string.
type DestinationSelector = string

// Signer represents a participant in the commit verification process.
type Signer struct {
	Address string `toml:"address"`
}

// QuorumConfig represents the configuration for a quorum of signers.
type QuorumConfig struct {
	SourceVerifierAddress string   `toml:"sourceVerifierAddress"`
	Signers               []Signer `toml:"signers"`
	Threshold             uint8    `toml:"threshold"`
}

// Committee represents the serializable portion of an aggregator committee config.
// This is a deployment-only copy of the runtime model.Committee type, containing
// only the fields needed for config generation and datastore persistence.
type Committee struct {
	QuorumConfigs        map[SourceSelector]*QuorumConfig `toml:"quorumConfigs"`
	DestinationVerifiers map[DestinationSelector]string   `toml:"destinationVerifiers"`
}
