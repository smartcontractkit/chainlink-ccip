package rmn

import (
	"context"
	"errors"

	"github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
)

// ErrTimeout is returned when the signature computation times out.
var ErrTimeout = errors.New("signature computation timeout")

// Client contains the methods required by the plugin to interact with the RMN nodes.
type Client interface {
	// ComputeSignatures computes and returns the signatures for the provided lane updates.
	ComputeSignatures(
		ctx context.Context,
		destChain DestChainInfo,
		requestedUpdates []FixedDestLaneUpdateRequest,
	) (*NodeSignatures, error)
}

type NodeSignatures struct {
	// Signatures are the ECDSA signatures for the lane updates for each node.
	// NOTE: A signature[i] corresponds to the whole LaneUpdates slice and NOT LaneUpdates[i].
	Signatures  []ECDSASignature
	LaneUpdates []FixedDestLaneUpdate
}

type FixedDestLaneUpdateRequest struct {
	SourceChainInfo SourceChainInfo
	Interval        ClosedInterval
}

type FixedDestLaneUpdate struct {
	SourceChain SourceChainInfo
	Interval    ClosedInterval
	MerkleRoot  []byte
}

type ClosedInterval struct {
	Min ccipocr3.SeqNum
	Max ccipocr3.SeqNum
}

type SourceChainInfo struct {
	Chain         ccipocr3.ChainSelector
	OnRampAddress []byte
}

type DestChainInfo struct {
	Chain          ccipocr3.ChainSelector
	OffRampAddress []byte
}

type ECDSASignature struct {
	R []byte
	S []byte
}
