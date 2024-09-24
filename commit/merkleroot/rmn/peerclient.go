// peerclient.go contains functionality for low-level communication with RMN peers.

package rmn

import (
	"context"
	"encoding/binary"
	"fmt"
	"sync"

	"github.com/smartcontractkit/libocr/commontypes"
	ocrnetworking "github.com/smartcontractkit/libocr/networking"
	ocr2types "github.com/smartcontractkit/libocr/offchainreporting2plus/types"
	"github.com/smartcontractkit/libocr/ragep2p"
)

// PeerClient performs low-level communication with RMN peers.
type PeerClient interface {
	// InitConnection initializes the connection to the generic peer endpoint and must be called before
	// further PeerClient interaction.
	InitConnection(ctx context.Context /* and some other params */) error

	// Close closes all the streams and connections.
	Close(ctx context.Context) error

	// Send sends a request to the RMN node with the given NodeID.
	Send(rmnNodeID NodeID, request []byte) error

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
	netEndpointFactory                ocrnetworking.GenericNetworkEndpointFactory
	respChan                          chan PeerResponse
	genericEndpoint                   ocrnetworking.GenericNetworkEndpoint // might be nil initially
	genericEndpointConfigDigestPrefix ocr2types.ConfigDigestPrefix
	rageP2PStreams                    map[NodeID]*ragep2p.Stream
	mu                                *sync.RWMutex
}

func NewPeerClient(netEndpointFactory ocrnetworking.GenericNetworkEndpointFactory) PeerClient {
	return &peerClient{
		netEndpointFactory: netEndpointFactory,
		respChan:           make(chan PeerResponse),

		genericEndpoint:                   nil,
		rageP2PStreams:                    make(map[NodeID]*ragep2p.Stream),
		genericEndpointConfigDigestPrefix: 0,
		mu:                                &sync.RWMutex{},
	}
}

func (r *peerClient) InitConnection(ctx context.Context) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	// configDigest = prefix(ConfigDigestPrefixCCIPMultiRoleRMNCombo, sha256(commitConfigDigest | rmnHomeConfigDigest))
	r.genericEndpointConfigDigestPrefix = 0

	/*
		prefix(
			ConfigDigestPrefixCCIPMultiRoleRMNCombo,
			sha256(
				commit OCR config digest | RMN home config digest
			)
		)
	*/

	// peerIDs := union(oraclePeerIDs, rmnNodePeerIDs)
	var peerIDs []string

	// ?????
	var bootstrappers []commontypes.BootstrapperLocator

	genericEndpoint, err := r.netEndpointFactory.NewGenericEndpoint(
		[32]byte{},
		peerIDs,
		bootstrappers,
	)
	fmt.Println(genericEndpoint)

	if err != nil {
		return fmt.Errorf("new generic endpoint: %w", err)
	}
	return nil
}

func (r *peerClient) Close(ctx context.Context) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, stream := range r.rageP2PStreams {
		if err := stream.Close(); err != nil {
			// todo: lggr
			continue
		}
	}
	return r.genericEndpoint.Close()
}

func (r *peerClient) Send(rmnNodeID NodeID, request []byte) error {
	stream, err := r.getOrCreateRageP2PStream(rmnNodeID)
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

func (r *peerClient) getOrCreateRageP2PStream(rmnNodeID NodeID) (*ragep2p.Stream, error) {
	r.mu.RLock()
	stream, ok := r.rageP2PStreams[rmnNodeID]
	r.mu.RUnlock()
	if ok {
		return stream, nil
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	var ragePeerID [32]byte
	uint32Bytes := make([]byte, 4)
	binary.BigEndian.PutUint32(uint32Bytes, uint32(rmnNodeID)) // todo: what's the rage peer id

	// todo: params configurable and param tuning
	stream, err := r.genericEndpoint.NewStream(
		ragePeerID,
		fmt.Sprintf("ccip-rmn/v1_6/%d", r.genericEndpointConfigDigestPrefix), // todo: version
		100,
		100,
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
		return nil, fmt.Errorf("new stream: %w", err)
	}

	r.rageP2PStreams[rmnNodeID] = stream
	go r.listenToStream(rmnNodeID, stream)
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
