// TODO(remove-blessing): delete once chainlink drops RMN blessing imports.
package types

import (
	"crypto/ed25519"

	mapset "github.com/deckarep/golang-set/v2"

	ragep2ptypes "github.com/smartcontractkit/libocr/ragep2p/types"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
)

type NodeID uint32

// HomeNodeInfo is retained for chainlink compile compatibility only.
type HomeNodeInfo struct {
	ID                    NodeID
	PeerID                ragep2ptypes.PeerID
	SupportedSourceChains mapset.Set[cciptypes.ChainSelector]
	OffchainPublicKey     *ed25519.PublicKey
	StreamNamePrefix      string
}
