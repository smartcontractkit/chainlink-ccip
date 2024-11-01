package peergroup

import (
	"crypto/sha256"
	"encoding/binary"
	"fmt"

	"github.com/smartcontractkit/libocr/commontypes"
	"github.com/smartcontractkit/libocr/networking"
	ocr2types "github.com/smartcontractkit/libocr/offchainreporting2plus/types"
	ragep2ptypes "github.com/smartcontractkit/libocr/ragep2p/types"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"

	rmntypes "github.com/smartcontractkit/chainlink-ccip/commit/merkleroot/rmn/types"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

// Creator handles peer group creation
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

// CreateOpts defines options for creating a peer group
type CreateOpts struct {
	CommitConfigDigest  cciptypes.Bytes32
	RMNHomeConfigDigest cciptypes.Bytes32
	OraclePeerIDs       []ragep2ptypes.PeerID
	RMNNodes            []rmntypes.HomeNodeInfo
}

// Result contains the created peer group and its config digest
type Result struct {
	PeerGroup    networking.PeerGroup
	ConfigDigest cciptypes.Bytes32
}

// Create handles peer group creation with the given options
func (c *Creator) Create(opts CreateOpts) (Result, error) {
	// Calculate generic endpoint config digest
	h := sha256.Sum256(append(opts.CommitConfigDigest[:], opts.RMNHomeConfigDigest[:]...))
	genericEndpointConfigDigest := writePrefix(ocr2types.ConfigDigestPrefixCCIPMultiRoleRMNCombo, h)

	// Combine peer IDs
	peerIDs := make([]string, 0, len(opts.OraclePeerIDs)+len(opts.RMNNodes))
	for _, p := range opts.OraclePeerIDs {
		peerIDs = append(peerIDs, p.String())
	}
	for _, n := range opts.RMNNodes {
		peerIDs = append(peerIDs, n.PeerID.String())
	}

	c.lggr.Infow("Creating new peer group",
		"genericEndpointConfigDigest", genericEndpointConfigDigest.String(),
		"peerIDs", peerIDs,
		"bootstrappers", c.bootstrappers,
	)

	peerGroup, err := c.factory.NewPeerGroup(
		[32]byte(genericEndpointConfigDigest),
		peerIDs,
		c.bootstrappers,
	)
	if err != nil {
		return Result{}, fmt.Errorf("new peer group: %w", err)
	}

	c.lggr.Infow("Created new peer group successfully")

	return Result{
		PeerGroup:    peerGroup,
		ConfigDigest: genericEndpointConfigDigest,
	}, nil
}

// writePrefix is kept in the same package for direct use by the Creator
func writePrefix(prefix ocr2types.ConfigDigestPrefix, hash cciptypes.Bytes32) cciptypes.Bytes32 {
	var prefixBytes [2]byte
	binary.BigEndian.PutUint16(prefixBytes[:], uint16(prefix))

	hCopy := hash
	hCopy[0] = prefixBytes[0]
	hCopy[1] = prefixBytes[1]

	return hCopy
}
