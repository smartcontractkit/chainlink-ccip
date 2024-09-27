package types

import (
	"crypto/ed25519"

	mapset "github.com/deckarep/golang-set/v2"

	ragep2ptypes "github.com/smartcontractkit/libocr/ragep2p/types"

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
	PeerID                ragep2ptypes.PeerID                 // The peer ID of the node
	SupportedSourceChains mapset.Set[cciptypes.ChainSelector] // Set of supported source chains by the node
	OffchainPublicKey     *ed25519.PublicKey                  // The public key is used to verify observations
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
	OnchainPublicKey      cciptypes.Bytes // The signer's onchain address, used to verify report signature
	NodeIndex             uint64          // The index of the node in the RMN config
	SignObservationPrefix string          // The prefix of the observation to sign
}
