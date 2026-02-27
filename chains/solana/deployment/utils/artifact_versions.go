package utils

// Mapping between Semver Tags and commit SHA from chainlink-ccip repository for Solana Contracts.
// Source tags (GitHub releases):
// - solana-v0.1.2
// - solana-v0.1.1-cctp
// - solana-v0.1.1
// - solana-v0.1.0

// Public version constants.
const (
	// Versions for Chainlink CCIP Solana contracts
	VersionSolanaV0_1_2           = "solana-v0.1.2"
	VersionSolanaV0_1_1           = "solana-v0.1.1-cctp"
	VersionSolanaV0_1_1TokenPools = "solana-v0.1.1"
	VersionSolanaV0_1_0           = "solana-v0.1.0"
	VersionSolanaV1_6_0           = "solana-v1.6.0"
	VersionSolanaV1_6_1           = "solana-v1.6.1"
)

// VersionToShortCommitSHA maps a version tag to its corresponding short commit SHA.
var VersionToShortCommitSHA = map[string]string{
	VersionSolanaV0_1_2:           "b96a80a69ad2",
	VersionSolanaV0_1_1:           "7f8a0f403c3a",
	VersionSolanaV0_1_1TokenPools: "ee587a6c0562",
	VersionSolanaV0_1_0:           "be8d09930aaa",
	VersionSolanaV1_6_0:           "d0d81df31957",
	VersionSolanaV1_6_1:           "cb23ec38649f",
}

var VersionToFullCommitSHA = map[string]string{
	VersionSolanaV0_1_2:           "b96a80a69ad2696c48d645d0cf7807fd02a212c8",
	VersionSolanaV0_1_1:           "7f8a0f403c3acbf740fa6d50d71bfb80a8b12ab8",
	VersionSolanaV0_1_1TokenPools: "ee587a6c056204009310019b790ed6d474825316",
	VersionSolanaV0_1_0:           "be8d09930aaaae31b574ef316ca73021fe272b08",
	VersionSolanaV1_6_0:           "d0d81df3195728091cad1b0569a2980201a92e97",
	VersionSolanaV1_6_1:           "cb23ec38649f9d23aabd0350e30d3d649ebc2174",
}
