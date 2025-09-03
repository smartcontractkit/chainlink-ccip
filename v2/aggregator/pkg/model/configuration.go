package model

// Signer represents a participant in the commit verification process.
type Signer struct {
	ParticipantID string `toml:"participantID"`
	Addresses     []byte `toml:"addresses"`
}

// Committee represents a group of signers participating in the commit verification process.
type Committee struct {
	// QuorumConfigs stores a QuorumConfig for each chain selector
	// there is a commit verifier for.
	// The aggregator uses this to verify signatures from each chain's
	// commit verifier set.
	QuorumConfigs map[uint64]QuorumConfig `toml:"quorumConfigs"`
}

// QuorumConfig represents the configuration for a quorum of signers.
type QuorumConfig struct {
	Signers []Signer `toml:"signers"`
	F       uint8    `toml:"f"`
}

// StorageConfig represents the configuration for the storage backend.
type StorageConfig struct {
	StorageType string `toml:"type,default=memory"`
}

// AggregationConfig represents the configuration for the aggregation process.
type AggregationConfig struct {
	AggregationStrategy string `toml:"strategy,default=stub"`
}

// ServerConfig represents the configuration for the server.
type ServerConfig struct {
	Address string `toml:"address,default=:50051"`
}

// AggregatorConfig is the root configuration for the aggregator.
type AggregatorConfig struct {
	Server      ServerConfig         `toml:"server"`
	Storage     StorageConfig        `toml:"storage"`
	Aggregation AggregationConfig    `toml:"aggregation"`
	Committees  map[string]Committee `toml:"committees"`
}

// Validate validates the aggregator configuration for integrity and correctness.
func (c *AggregatorConfig) Validate() error {
	// TODO: Add Validate() method to AggregatorConfig to ensure configuration integrity
	// Should validate:
	// - No duplicate signers within the same QuorumConfig
	// - StorageType is supported (memory, etc.)
	// - AggregationStrategy is supported (stub, etc.)
	// - F value follows N = 3F + 1 rule, so F = (N-1) // 3
	// - Committee names are valid
	// - QuorumConfig chain selectors are valid
	// - Server address format is correct
	return nil
}
