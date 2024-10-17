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
	Nodes                   []HomeNodeInfo
	SourceChainMinObservers map[cciptypes.ChainSelector]int
	ConfigDigest            cciptypes.Bytes32
	OffchainConfig          cciptypes.Bytes // The raw offchain config
}

// HomeNodeInfo contains information about a node from the RMNHome contract.
type HomeNodeInfo struct {
	ID                    NodeID                              // ID is the index of this node in the RMN config
	PeerID                ragep2ptypes.PeerID                 // The peer ID of the node
	SupportedSourceChains mapset.Set[cciptypes.ChainSelector] // Set of supported source chains by the node
	OffchainPublicKey     *ed25519.PublicKey                  // The public key is used to verify observations
}

// RemoteConfig contains the configuration fetched from the RMNRemote contract.
type RemoteConfig struct {
	ContractAddress  cciptypes.Bytes    `json:"contractAddress"`
	ConfigDigest     cciptypes.Bytes32  `json:"configDigest"`
	Signers          []RemoteSignerInfo `json:"signers"`
	MinSigners       uint64             `json:"minSigners"`
	ConfigVersion    uint32             `json:"configVersion"`
	RmnReportVersion cciptypes.Bytes32  `json:"rmnReportVersion"` // e.g., keccak256("RMN_V1_6_ANY2EVM_REPORT")
}

func (r RemoteConfig) IsEmpty() bool {
	return len(r.ContractAddress) == 0 &&
		r.ConfigDigest == (cciptypes.Bytes32{}) &&
		len(r.Signers) == 0 &&
		r.MinSigners == 0 &&
		r.ConfigVersion == 0 &&
		r.RmnReportVersion == (cciptypes.Bytes32{})
}

// RemoteSignerInfo contains information about a signer from the RMNRemote contract.
type RemoteSignerInfo struct {
	// The signer's onchain address, used to verify report signature
	OnchainPublicKey cciptypes.Bytes `json:"onchainPublicKey"`
	// The index of the node in the RMN config
	NodeIndex uint64 `json:"nodeIndex"`
}
