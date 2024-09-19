package types

import (
	"crypto/ed25519"

	mapset "github.com/deckarep/golang-set/v2"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
)

type NodeID uint32

// RMNHomeConfig contains the configuration fetched from the RMNHome contract.
type RMNHomeConfig struct {
	Nodes                   []RMNHomeNodeInfo
	SourceChainMinObservers map[cciptypes.ChainSelector]uint64
	ConfigDigest            cciptypes.Bytes32
	OffchainConfig          cciptypes.Bytes // The raw offchain config
}

// RMNHomeNodeInfo contains information about a node from the RMNHome contract.
type RMNHomeNodeInfo struct {
	ID                    NodeID                              // ID is the index of this node in the RMN config
	PeerID                cciptypes.Bytes32                   // The peer ID of the node
	SupportedSourceChains mapset.Set[cciptypes.ChainSelector] // Set of supported source chains by the node
	OffchainPublicKey     *ed25519.PublicKey                  // The private key is used to verify observations
}

// RMNRemoteConfig contains the configuration fetched from the RMNRemote contract.
type RMNRemoteConfig struct {
	ContractAddress  cciptypes.Bytes
	ConfigDigest     cciptypes.Bytes32
	Signers          []RMNRemoteSignerInfo
	MinSigners       uint64
	ConfigVersion    uint32
	RmnReportVersion string // e.g., "RMN_V1_6_ANY2EVM_REPORT"
}

// RMNRemoteSignerInfo contains information about a signer from the RMNRemote contract.
type RMNRemoteSignerInfo struct {
	SignerOnchainAddress  cciptypes.Bytes // The signer's onchain address, used to verify report signature
	NodeIndex             uint64          // The index of the node in the RMN config
	SignObservationPrefix string          // The prefix of the observation to sign
}

// VersionedConfigWithDigest mirrors RMNHome.sol's VersionedConfigWithDigest struct
type VersionedConfigWithDigest struct {
	// nolint:lll // don't split up the long url
	// https://github.com/smartcontractkit/ccip/blob/e6e26ad31eef625faf68806a2b4f0549bc89b15c/contracts/src/v0.8/ccip/RMNRemote.sol#L34
	ConfigDigest    cciptypes.Bytes32 `json:"configDigest"`
	VersionedConfig VersionedConfig   `json:"versionedConfig"`
}

// VersionedConfig mirrors RMNHome.sol's VersionedConfig struct
type VersionedConfig struct {
	Version uint32 `json:"version"`
	Config  Config `json:"config"`
}

// Config mirrors RMNHome.sol's Config struct
type Config struct {
	Nodes          []Node          `json:"nodes"`
	SourceChains   []SourceChain   `json:"sourceChains"`
	OffchainConfig cciptypes.Bytes `json:"offchainConfig"`
}

// Node mirrors RMNHome.sol's Node struct
type Node struct {
	PeerID            cciptypes.Bytes32 `json:"peerId"`
	OffchainPublicKey cciptypes.Bytes32 `json:"offchainPublicKey"`
}

// SourceChain mirrors RMNHome.sol's SourceChain struct
type SourceChain struct {
	ChainSelector       cciptypes.ChainSelector `json:"chainSelector"`
	MinObservers        uint64                  `json:"minObservers"`
	ObserverNodesBitmap cciptypes.BigInt        `json:"observerNodesBitmap"`
}
