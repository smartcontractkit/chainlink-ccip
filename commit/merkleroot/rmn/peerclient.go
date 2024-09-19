// peerclient.go contains functionality for low-level communication with RMN peers.

package rmn

import (
	"errors"

	mapset "github.com/deckarep/golang-set/v2"
	ragep2ptypes "github.com/smartcontractkit/libocr/ragep2p/types"
)

var (
	ErrRMNNodeNotFound = errors.New("rmn node not found")
)

// PeerClient performs low-level communication with RMN peers.
type PeerClient interface {
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
	ragep2ptypes.Address
	peer       ragep2ptypes.PeerInfo
	rmnPeerIDs mapset.Set[NodeID]
	respChan   chan PeerResponse
}

func NewPeerClient(peer ragep2ptypes.PeerInfo, rmnNodeIDs []NodeID) PeerClient {
	return &peerClient{
		peer:       peer,
		rmnPeerIDs: mapset.NewSet(rmnNodeIDs...),
		respChan:   make(chan PeerResponse),
	}
}

func (r *peerClient) Send(rmnNodeID NodeID, request []byte) error {
	if !r.rmnPeerIDs.Contains(rmnNodeID) {
		return ErrRMNNodeNotFound
	}
	panic("implement actual send")
}

func (r *peerClient) Recv() <-chan PeerResponse {
	return r.respChan
}
