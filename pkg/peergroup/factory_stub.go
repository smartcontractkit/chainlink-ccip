// Package peergroup is a temporary chainlink compatibility shim.
// TODO(remove-blessing): delete once chainlink drops peergroup imports.
package peergroup

import (
	"fmt"

	"github.com/smartcontractkit/libocr/commontypes"
	"github.com/smartcontractkit/libocr/networking"
	ragep2ptypes "github.com/smartcontractkit/libocr/ragep2p/types"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	rmntypes "github.com/smartcontractkit/chainlink-ccip/commit/merkleroot/rmn/types"
)

// Creator handles peer group creation.
type Creator struct {
	lggr          logger.Logger
	factory       networking.PeerGroupFactory
	bootstrappers []commontypes.BootstrapperLocator
}

func NewCreator(
	lggr logger.Logger,
	factory networking.PeerGroupFactory,
	bootstrappers []commontypes.BootstrapperLocator,
) *Creator {
	return &Creator{
		lggr:          lggr,
		factory:       factory,
		bootstrappers: bootstrappers,
	}
}

// CreateOpts defines options for creating a peer group.
type CreateOpts struct {
	CommitConfigDigest  cciptypes.Bytes32
	RMNHomeConfigDigest cciptypes.Bytes32
	OraclePeerIDs       []ragep2ptypes.PeerID
	RMNNodes            []rmntypes.HomeNodeInfo
}

// Result contains the created peer group and its config digest.
type Result struct {
	PeerGroup    networking.PeerGroup
	ConfigDigest cciptypes.Bytes32
}

// Create is a no-op stub retained for chainlink bootstrap compile compatibility.
func (c *Creator) Create(_ CreateOpts) (Result, error) {
	return Result{}, fmt.Errorf("RMN peer groups are disabled")
}
