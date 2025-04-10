package types

import (
	"crypto/ed25519"

	mapset "github.com/deckarep/golang-set/v2"

	ragep2ptypes "github.com/smartcontractkit/libocr/ragep2p/types"

	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

type NodeID uint32

// HomeConfig contains the configuration fetched from the RMNHome contract.
type HomeConfig struct {
	Nodes []HomeNodeInfo
	// SourceChainF contains the "fObserve" for RMN interactions for each source chain.
	// If a chain does not appear in this map, it is assumed that it is not RMN-enabled and signatures are not required.
	SourceChainF   map[cciptypes.ChainSelector]int
	ConfigDigest   cciptypes.Bytes32
	OffchainConfig cciptypes.Bytes // The raw offchain config
}

// HomeNodeInfo contains information about a node from the RMNHome contract.
type HomeNodeInfo struct {
	ID                    NodeID                              // ID is the index of this node in the RMN config
	PeerID                ragep2ptypes.PeerID                 // The peer ID of the node
	SupportedSourceChains mapset.Set[cciptypes.ChainSelector] // Set of supported source chains by the node
	OffchainPublicKey     *ed25519.PublicKey                  // The public key is used to verify observations
	StreamNamePrefix      string                              // RageP2P stream name prefix e.g. "ccip-rmn/v1_6/"
}

// Deprecated: use cciptypes.RemoteConfig
type RemoteConfig = cciptypes.RemoteConfig

// Deprecated: use cciptypes.RemoteSignerInfo
type RemoteSignerInfo = cciptypes.RemoteSignerInfo
