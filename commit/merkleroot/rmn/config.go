package rmn

import (
	"crypto/ed25519"

	mapset "github.com/deckarep/golang-set/v2"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
)

// Config contains the RMN configuration required by the plugin and the RMN client in a single struct.
type Config struct {
	Home   RMNHomeConfig
	Remote RMNRemoteConfig
}

type RMNHomeConfig struct {
	RmnNodes         []RMNNodeInfo
	ConfigDigest     cciptypes.Bytes32
	RmnReportVersion string // e.g. "RMN_V1_6_ANY2EVM_REPORT"
}

type RMNRemoteConfig struct {
	ContractAddress cciptypes.Bytes
	MinObservers    int
	MinSigners      int
}

// RMNNodeInfo contains the information about an RMN node.
type RMNNodeInfo struct {
	// ID is the index of this node in the RMN config
	ID                        NodeID
	SupportedSourceChains     mapset.Set[cciptypes.ChainSelector]
	IsSigner                  bool
	SignReportsAddress        cciptypes.Bytes
	SignObservationsPublicKey *ed25519.PublicKey // offChainPublicKey
	// TODO: clarify this field
	SignObservationPrefix string // e.g. "chainlink ccip 1.6 rmn observation"
}

type NodeID uint32
