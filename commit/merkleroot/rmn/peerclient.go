// peerclient.go contains functionality for low-level communication with RMN peers.

package rmn

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

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
	lggr                        logger.Logger
	peerGroupCreator            *peergroup.Creator
	respChan                    chan PeerResponse
	peerGroup                   networking.PeerGroup
	genericEndpointConfigDigest cciptypes.Bytes32
	rageP2PStreams              map[rmntypes.NodeID]Stream
	bootstrappers               []commontypes.BootstrapperLocator
	// ocrRoundInterval is the estimated interval between OCR rounds.
	ocrRoundInterval time.Duration
	mu               *sync.RWMutex
}

func NewPeerClient(
	lggr logger.Logger,
	peerGroupFactory networking.PeerGroupFactory,
	bootstrappers []commontypes.BootstrapperLocator,
	ocrRoundInterval time.Duration,
) PeerClient {
	return &peerClient{
		lggr:                        lggr,
		peerGroupCreator:            peergroup.NewCreator(lggr, peerGroupFactory, bootstrappers),
		respChan:                    make(chan PeerResponse),
		peerGroup:                   nil,
		rageP2PStreams:              make(map[rmntypes.NodeID]Stream),
		genericEndpointConfigDigest: cciptypes.Bytes32{},
		bootstrappers:               bootstrappers,
		ocrRoundInterval:            ocrRoundInterval,
		mu:                          &sync.RWMutex{},
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

	r.peerGroup = result.PeerGroup
	r.genericEndpointConfigDigest = result.ConfigDigest
	return nil
}

func (r *peerClient) Close() error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.peerGroup == nil {
		return nil
	}

	// individual streams are closed by the peer group
	if err := r.peerGroup.Close(); err != nil {
		return fmt.Errorf("close peer group: %w", err)
	}

	r.peerGroup = nil
	r.rageP2PStreams = make(map[rmntypes.NodeID]Stream)
	return nil
}

// TODO: pass in ctx to get OCR seqNr. Then can include seqNr in the log.
func (r *peerClient) Send(rmnNode rmntypes.HomeNodeInfo, request []byte) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.peerGroup == nil {
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
	streamName := rmnNode.StreamNamePrefix + strings.TrimPrefix(r.genericEndpointConfigDigest.String(), "0x")

	r.lggr.Infow("creating new stream",
		"streamName", streamName,
		"rmnPeerID", rmnPeerID,
		"rmnNodeIdx", rmnNode.ID,
		"rmnNodeSupportedSourceChains", rmnNode.SupportedSourceChains.String(),
	)

	var err error
	stream, err = r.peerGroup.NewStream(rmnPeerID, newStreamConfig(r.lggr, streamName, r.ocrRoundInterval))
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
