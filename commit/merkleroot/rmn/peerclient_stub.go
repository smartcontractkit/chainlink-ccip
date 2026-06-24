// Package rmn provides a no-op PeerClient for chainlink compatibility.
// TODO(remove-blessing): delete once chainlink drops RMN blessing imports.
package rmn

import (
	"context"
	"fmt"
	"time"

	"github.com/smartcontractkit/libocr/commontypes"
	"github.com/smartcontractkit/libocr/networking"
	ragep2ptypes "github.com/smartcontractkit/libocr/ragep2p/types"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	rmntypes "github.com/smartcontractkit/chainlink-ccip/commit/merkleroot/rmn/types"
)

// PeerClient performs low-level communication with RMN peers.
type PeerClient interface {
	InitConnection(
		ctx context.Context,
		commitConfigDigest cciptypes.Bytes32,
		rmnHomeConfigDigest cciptypes.Bytes32,
		oraclePeerIDs []ragep2ptypes.PeerID,
		rmnNodes []rmntypes.HomeNodeInfo,
	) error
	Close() error
	Send(rmnNode rmntypes.HomeNodeInfo, request []byte) error
	Recv() <-chan PeerResponse
}

// PeerResponse represents a low-level response from an RMN node.
type PeerResponse struct {
	RMNNodeID rmntypes.NodeID
	Body      []byte
}

type noopPeerClient struct {
	lggr     logger.Logger
	respChan chan PeerResponse
}

// NewPeerClient returns a no-op PeerClient.
func NewPeerClient(
	lggr logger.Logger,
	_ networking.PeerGroupFactory,
	_ []commontypes.BootstrapperLocator,
	_ time.Duration,
) PeerClient {
	return &noopPeerClient{
		lggr:     lggr,
		respChan: make(chan PeerResponse),
	}
}

func (c *noopPeerClient) InitConnection(
	_ context.Context,
	_ cciptypes.Bytes32,
	_ cciptypes.Bytes32,
	_ []ragep2ptypes.PeerID,
	_ []rmntypes.HomeNodeInfo,
) error {
	c.lggr.Debugw("noop RMN peer client InitConnection")
	return nil
}

func (c *noopPeerClient) Close() error {
	return nil
}

func (c *noopPeerClient) Send(_ rmntypes.HomeNodeInfo, _ []byte) error {
	return fmt.Errorf("RMN peer client is disabled")
}

func (c *noopPeerClient) Recv() <-chan PeerResponse {
	return c.respChan
}
