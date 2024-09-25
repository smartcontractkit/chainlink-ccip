// peerclient.go contains functionality for low-level communication with RMN peers.

package rmn

import (
	"context"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"sync"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
	"github.com/smartcontractkit/libocr/commontypes"
	ocrnetworking "github.com/smartcontractkit/libocr/networking"
	ocr2types "github.com/smartcontractkit/libocr/offchainreporting2plus/types"
	"github.com/smartcontractkit/libocr/ragep2p"
)

// PeerClient performs low-level communication with RMN peers.
type PeerClient interface {
	// InitConnection initializes the connection to the generic peer endpoint and must be called before
	// further PeerClient interaction.
	InitConnection(
		ctx context.Context,
		commitConfigDigest, rmnHomeConfigDigest cciptypes.Bytes32,
		peerIDs []string, // union of oraclePeerIDs and rmnNodePeerIDs
		bootstrappers []commontypes.BootstrapperLocator,
	) error

	// Close closes all the streams and connections.
	Close(ctx context.Context) error

	// Send sends a request to the RMN node with the given NodeID.
	Send(rmnNodeID RMNNodeInfo, request []byte) error

	// Recv returns a channel which can be used to listen on for
	// responses by all RMN nodes. This is expected to be monitored
	// by the plugin in order to get RMN responses.
	Recv() <-chan PeerResponse
}

// PeerResponse represents a low-level response from an RMN node.
type PeerResponse struct {
	RMNNodeID NodeID
	Body      []byte // pb
}

type peerClient struct {
	lggr                        logger.Logger
	netEndpointFactory          ocrnetworking.GenericNetworkEndpointFactory
	respChan                    chan PeerResponse
	genericEndpoint             ocrnetworking.GenericNetworkEndpoint // might be nil initially
	genericEndpointConfigDigest cciptypes.Bytes32
	rageP2PStreams              map[NodeID]*ragep2p.Stream
	mu                          *sync.RWMutex
}

func NewPeerClient(lggr logger.Logger, netEndpointFactory ocrnetworking.GenericNetworkEndpointFactory) PeerClient {
	return &peerClient{
		netEndpointFactory: netEndpointFactory,
		respChan:           make(chan PeerResponse),

		genericEndpoint:             nil,
		rageP2PStreams:              make(map[NodeID]*ragep2p.Stream),
		genericEndpointConfigDigest: cciptypes.Bytes32{},
		mu:                          &sync.RWMutex{},
	}
}

func (r *peerClient) InitConnection(
	ctx context.Context,
	commitConfigDigest, rmnHomeConfigDigest cciptypes.Bytes32,
	peerIDs []string,
	bootstrappers []commontypes.BootstrapperLocator,
) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	h := sha256.Sum256(append(commitConfigDigest[:], rmnHomeConfigDigest[:]...))
	r.genericEndpointConfigDigest = writePrefix(ocr2types.ConfigDigestPrefixCCIPMultiRoleRMNCombo, h)

	genericEndpoint, err := r.netEndpointFactory.NewGenericEndpoint(
		[32]byte(r.genericEndpointConfigDigest),
		peerIDs,
		bootstrappers,
	)

	if err != nil {
		return fmt.Errorf("new generic endpoint: %w", err)
	}

	r.genericEndpoint = genericEndpoint
	return nil
}

func (r *peerClient) Close(ctx context.Context) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	// individual streams are closed by the generic endpoint
	if err := r.genericEndpoint.Close(); err != nil {
		return fmt.Errorf("close generic endpoint: %w", err)
	}

	return nil
}

func (r *peerClient) Send(rmnNode RMNNodeInfo, request []byte) error {
	stream, err := r.getOrCreateRageP2PStream(rmnNode)
	if err != nil {
		return fmt.Errorf("get or create rage p2p stream: %w", err)
	}

	r.mu.Lock()
	stream.SendMessage(request)
	r.mu.Unlock()

	return nil
}

func (r *peerClient) Recv() <-chan PeerResponse {
	return r.respChan
}

func (r *peerClient) getOrCreateRageP2PStream(rmnNode RMNNodeInfo) (*ragep2p.Stream, error) {
	r.mu.RLock()
	stream, ok := r.rageP2PStreams[rmnNode.ID]
	r.mu.RUnlock()
	if ok {
		return stream, nil
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	// libocr accepts peer IDs either as string or [32]byte. Essentially it's a 32 byte hex string.
	ragePeerID, err := cciptypes.NewBytes32FromString(rmnNode.PeerID)
	if err != nil {
		return nil, fmt.Errorf("decode peer ID: %w", err)
	}

	// todo: versioning for stream names e.g. for 'v1_7'
	streamName := fmt.Sprintf("ccip-rmn/v1_6/%x", r.genericEndpointConfigDigest) // no '0x' prefix
	r.lggr.Info("Creating new stream", "streamName", streamName)

	// todo: params configurable and param tuning after consulting with research team
	stream, err = r.genericEndpoint.NewStream(
		[32]byte(ragePeerID),
		streamName,
		1,
		1,
		2_097_152, // 2MB
		ragep2p.TokenBucketParams{ /* messages rate limit*/
			Rate:     20,
			Capacity: 100,
		},
		ragep2p.TokenBucketParams{ /* bytes rate limit */
			Rate:     20_971_520,  // 20MB
			Capacity: 104_857_600, // 100MB
		},
	)
	if err != nil {
		return nil, fmt.Errorf("new stream %s: %w", streamName, err)
	}

	r.rageP2PStreams[rmnNode.ID] = stream
	go r.listenToStream(rmnNode.ID, stream)
	return stream, nil
}

func (r *peerClient) listenToStream(rmnNodeID NodeID, stream *ragep2p.Stream) {
	for msg := range stream.ReceiveMessages() {
		r.respChan <- PeerResponse{
			RMNNodeID: rmnNodeID,
			Body:      msg,
		}
	}
}

// writePrefix writes the prefix to the rightmost 2 bytes of the hash.
//
// Example:
//
//	prefix = [2]byte{10, 20}
//	h = [32]byte{1, 2, 3.....}
//	result = [32]byte{10, 20, 3, 4, ...}
//
// NOT TESTED!
func writePrefix(prefix ocr2types.ConfigDigestPrefix, hash cciptypes.Bytes32) cciptypes.Bytes32 {
	var prefixBytes [2]byte
	binary.BigEndian.PutUint16(prefixBytes[:], uint16(prefix))

	hCopy := hash
	hCopy[0] = prefixBytes[0]
	hCopy[1] = prefixBytes[1]

	return hCopy
}
