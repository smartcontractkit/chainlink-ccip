package rmn

import (
	"context"
	"errors"
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
