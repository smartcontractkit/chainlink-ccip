package types

import (
	"crypto/ed25519"

	mapset "github.com/deckarep/golang-set/v2"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
)

type NodeID uint32

// RMNConfig contains the RMN configuration required by the plugin and the RMN client in a single struct.
type RMNConfig struct {
	Home   RMNHomeConfig
	Remote RMNRemoteConfig
}

// RMNHomeConfig contains the configuration fetched from the RMNHome contract.
type RMNHomeConfig struct {
	Nodes          []RMNHomeNodeInfo
	MinObservers   map[cciptypes.ChainSelector]uint64
	ConfigDigest   cciptypes.Bytes32
	OffchainConfig cciptypes.Bytes // Raw offchain configuration bytes
}

// RMNHomeNodeInfo contains information about a node from the RMNHome contract.
type RMNHomeNodeInfo struct {
	ID                        NodeID // ID is the index of this node in the RMN config
	PeerID                    cciptypes.Bytes32
	SupportedSourceChains     mapset.Set[cciptypes.ChainSelector]
	SignObservationsPublicKey *ed25519.PublicKey // offchainPublicKey
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
	SignReportsAddress    cciptypes.Bytes // for signing reports
	NodeIndex             uint64          // maps to nodes in RMNHome
	SignObservationPrefix string          // for signing observations
}
