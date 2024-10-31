// peerclient.go contains functionality for low-level communication with RMN peers.

package rmn

import (
	"context"
	"fmt"
	"strings"
	"sync"

	"github.com/smartcontractkit/libocr/commontypes"
	"github.com/smartcontractkit/libocr/networking"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"

	rmntypes "github.com/smartcontractkit/chainlink-ccip/commit/merkleroot/rmn/types"
	"github.com/smartcontractkit/chainlink-ccip/pkg/peergroup"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"

	ragep2ptypes "github.com/smartcontractkit/libocr/ragep2p/types"
)

var ErrNoConn = fmt.Errorf("no connection, please call InitConnection before further interaction")

// PeerClient performs low-level communication with RMN peers.
type PeerClient interface {
	// InitConnection initializes the connection to the peer group endpoint and must be called before
	// further PeerClient interaction.
	InitConnection(
		ctx context.Context,
		commitConfigDigest cciptypes.Bytes32,
		rmnHomeConfigDigest cciptypes.Bytes32,
		oraclePeerIDs []ragep2ptypes.PeerID,
		rmnNodes []rmntypes.HomeNodeInfo,
	) error

	Close() error

	// Send will send a message to the target RMN node.
	// If Send is called before InitConnection, it returns an ErrNoConn.
	Send(rmnNode rmntypes.HomeNodeInfo, request []byte) error

	// Recv returns a channel which can be used to listen on for
	// responses by all RMN nodes. This is expected to be monitored
	// by the plugin in order to get RMN responses.
	Recv() <-chan PeerResponse
}

// PeerResponse represents a low-level response from an RMN node.
type PeerResponse struct {
	RMNNodeID rmntypes.NodeID
	Body      []byte // pb
}

type peerClient struct {
	lggr             logger.Logger
	peerGroupCreator *peergroup.Creator
	respChan         chan PeerResponse
	currentGroup     peergroup.PeerGroup
	configDigest     cciptypes.Bytes32
	rageP2PStreams   map[rmntypes.NodeID]Stream
	mu               *sync.RWMutex
}

func NewPeerClient(
	lggr logger.Logger,
	peerGroupFactory peergroup.PeerGroupFactory,
	bootstrappers []commontypes.BootstrapperLocator,
) PeerClient {
	return &peerClient{
		lggr:             lggr,
		peerGroupCreator: peergroup.NewCreator(lggr, peerGroupFactory, bootstrappers),
		respChan:         make(chan PeerResponse),
		rageP2PStreams:   make(map[rmntypes.NodeID]Stream),
		mu:               &sync.RWMutex{},
	}
}

func (r *peerClient) InitConnection(
	_ context.Context,
	commitConfigDigest, rmnHomeConfigDigest cciptypes.Bytes32,
	oraclePeerIDs []ragep2ptypes.PeerID,
	rmnNodes []rmntypes.HomeNodeInfo,

) error {
	if err := r.Close(); err != nil {
		return fmt.Errorf("close existing peer group: %w", err)
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	result, err := r.peerGroupCreator.Create(peergroup.CreateOpts{
		CommitConfigDigest:  commitConfigDigest,
		RMNHomeConfigDigest: rmnHomeConfigDigest,
		// Note: For RMN peer client, we receive the combined peer IDs directly
		// and don't need to separate oracle/RMN peers
		OraclePeerIDs: oraclePeerIDs,
		RMNNodes:      rmnNodes,
	})
	if err != nil {
		return fmt.Errorf("create peer group: %w", err)
	}

	r.currentGroup = result.PeerGroup
	r.configDigest = result.ConfigDigest
	return nil
}

func (r *peerClient) Close() error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.currentGroup == nil {
		return nil
	}

	// individual streams are closed by the peer group
	if err := r.currentGroup.Close(); err != nil {
		return fmt.Errorf("close peer group: %w", err)
	}

	r.currentGroup = nil
	r.rageP2PStreams = make(map[rmntypes.NodeID]Stream)
	return nil
}

func (r *peerClient) Send(rmnNode rmntypes.HomeNodeInfo, request []byte) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.currentGroup == nil {
		return ErrNoConn
	}

	stream, err := r.getOrCreateRageP2PStream(rmnNode)
	if err != nil {
		return fmt.Errorf("get or create rage p2p stream: %w", err)
	}

	r.lggr.Infow("sending message to RMN node", "rmnNodeID", rmnNode.ID, "requestSize", len(request))
	stream.SendMessage(request)

	return nil
}

func (r *peerClient) getOrCreateRageP2PStream(rmnNode rmntypes.HomeNodeInfo) (Stream, error) {
	stream, ok := r.rageP2PStreams[rmnNode.ID]
	if ok {
		return stream, nil
	}

	rmnPeerID := rmnNode.PeerID.String()

	// todo: versioning for stream names e.g. for 'v1_7'
	streamName := fmt.Sprintf("ccip-rmn/v1_6/%s",
		strings.TrimPrefix(r.configDigest.String(), "0x"))

	r.lggr.Infow("creating new stream",
		"streamName", streamName,
		"rmnPeerID", rmnPeerID,
		"rmnNodeIdx", rmnNode.ID,
		"rmnNodeSupportedSourceChains", rmnNode.SupportedSourceChains.String(),
	)

	pg := r.currentGroup.(peergroup.PeerGroup) // safe cast since we control creation
	stream, err := pg.NewStream(rmnPeerID, newStreamConfig(r.lggr, streamName))
	if err != nil {
		return nil, fmt.Errorf("new stream %s: %w", streamName, err)
	}

	r.rageP2PStreams[rmnNode.ID] = stream
	go r.listenToStream(rmnNode.ID, stream)
	return stream, nil
}

func (r *peerClient) listenToStream(rmnNodeID rmntypes.NodeID, stream Stream) {
	for msg := range stream.ReceiveMessages() {
		r.lggr.Infow("received message from RMN node", "rmnNodeID", rmnNodeID, "msgSize", len(msg))
		r.respChan <- PeerResponse{
			RMNNodeID: rmnNodeID,
			Body:      msg,
		}
	}
}

func (r *peerClient) Recv() <-chan PeerResponse {
	return r.respChan
}

// Redeclare interfaces for mocking purposes.
type Stream interface {
	networking.Stream
}
