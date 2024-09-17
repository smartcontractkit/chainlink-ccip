package types

import (
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
)

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
