// peerclient.go contains functionality for low-level communication with RMN peers.

package rmn

import (
	"context"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"strings"
	"sync"

	"github.com/smartcontractkit/libocr/commontypes"
	"github.com/smartcontractkit/libocr/networking"
	ocr2types "github.com/smartcontractkit/libocr/offchainreporting2plus/types"
	"github.com/smartcontractkit/libocr/ragep2p"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"

	rmntypes "github.com/smartcontractkit/chainlink-ccip/commit/merkleroot/rmn/types"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
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
		peerIDs []string, // union of oraclePeerIDs and rmnNodePeerIDs (oracles required for peer discovery)
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
	peerGroupFactory            PeerGroupFactory
	respChan                    chan PeerResponse
	peerGroup                   PeerGroup // nil initially, until InitConnection is called
	genericEndpointConfigDigest cciptypes.Bytes32
	rageP2PStreams              map[rmntypes.NodeID]Stream
	bootstrappers               []commontypes.BootstrapperLocator
	mu                          *sync.RWMutex
}

func NewPeerClient(
	lggr logger.Logger,
	peerGroupFactory PeerGroupFactory,
	bootstrappers []commontypes.BootstrapperLocator,
) PeerClient {
	return &peerClient{
		lggr:                        lggr,
		peerGroupFactory:            peerGroupFactory,
		respChan:                    make(chan PeerResponse),
		peerGroup:                   nil,
		rageP2PStreams:              make(map[rmntypes.NodeID]Stream),
		genericEndpointConfigDigest: cciptypes.Bytes32{},
		bootstrappers:               bootstrappers,
		mu:                          &sync.RWMutex{},
	}
}

func (r *peerClient) InitConnection(
	_ context.Context,
	commitConfigDigest, rmnHomeConfigDigest cciptypes.Bytes32,
	peerIDs []string,
) error {
	if err := r.Close(); err != nil {
		return fmt.Errorf("close existing peer group: %w", err)
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	h := sha256.Sum256(append(commitConfigDigest[:], rmnHomeConfigDigest[:]...))
	r.genericEndpointConfigDigest = writePrefix(ocr2types.ConfigDigestPrefixCCIPMultiRoleRMNCombo, h)
	r.lggr.Infow("Creating new peer group",
		"genericEndpointConfigDigest", r.genericEndpointConfigDigest.String())

	peerGroup, err := r.peerGroupFactory.NewPeerGroup(
		[32]byte(r.genericEndpointConfigDigest),
		peerIDs,
		r.bootstrappers,
	)

	if err != nil {
		return fmt.Errorf("new peer group: %w", err)
	}

	r.peerGroup = peerGroup
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

	return nil
}

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

	stream.SendMessage(request)

	return nil
}

func (r *peerClient) getOrCreateRageP2PStream(rmnNode rmntypes.HomeNodeInfo) (Stream, error) {
	stream, ok := r.rageP2PStreams[rmnNode.ID]
	if ok {
		return stream, nil
	}

	// todo: versioning for stream names e.g. for 'v1_7'
	streamName := fmt.Sprintf("ccip-rmn/v1_6/%s",
		strings.TrimPrefix(r.genericEndpointConfigDigest.String(), "0x"))
	r.lggr.Infow("Creating new stream", "streamName", streamName)

	var err error
	stream, err = r.peerGroup.NewStream(
		rmnNode.PeerID.String(),
		networking.NewStreamArgs1{
			StreamName:         streamName,
			OutgoingBufferSize: 1,
			IncomingBufferSize: 1,
			MaxMessageLength:   2_097_152, // 2MB
			MessagesLimit: ragep2p.TokenBucketParams{
				Rate:     20,
				Capacity: 100,
			},
			BytesLimit: ragep2p.TokenBucketParams{
				Rate:     20_971_520,  // 20MB
				Capacity: 104_857_600, // 100MB
			},
		},
	)
	if err != nil {
		return nil, fmt.Errorf("new stream %s: %w", streamName, err)
	}

	r.rageP2PStreams[rmnNode.ID] = stream
	go r.listenToStream(rmnNode.ID, stream)
	return stream, nil
}

func (r *peerClient) listenToStream(rmnNodeID rmntypes.NodeID, stream Stream) {
	for msg := range stream.ReceiveMessages() {
		r.respChan <- PeerResponse{
			RMNNodeID: rmnNodeID,
			Body:      msg,
		}
	}
}

func (r *peerClient) Recv() <-chan PeerResponse {
	return r.respChan
}

// writePrefix writes the prefix to the first 2 bytes of the hash.
func writePrefix(prefix ocr2types.ConfigDigestPrefix, hash cciptypes.Bytes32) cciptypes.Bytes32 {
	var prefixBytes [2]byte
	binary.BigEndian.PutUint16(prefixBytes[:], uint16(prefix))

	hCopy := hash
	hCopy[0] = prefixBytes[0]
	hCopy[1] = prefixBytes[1]

	return hCopy
}

// Redeclare interfaces for mocking purposes.

type PeerGroupFactory interface {
	networking.PeerGroupFactory
}

type PeerGroup interface {
	networking.PeerGroup
}

type Stream interface {
	networking.Stream
}
