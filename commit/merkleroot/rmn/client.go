package rmn

import (
	"context"
	"errors"
)

// ErrTimeout is returned when the signature computation times out.
var ErrTimeout = errors.New("signature computation timeout")

// Client contains the methods required by the plugin to interact with the RMN nodes.
type Client interface {
	// ComputeReportSignatures computes and returns the signatures for the provided lane updates.
	//
	// This method abstracts away the RMN specific requests (ObservationRequest, ReportSignaturesRequest) and all the
	// necessary steps to compute the signatures, like retrying and caching which are handled by the implementation.
	ComputeReportSignatures(
		ctx context.Context,
		destChain DestChainInfo,
		requestedUpdates []FixedDestLaneUpdateRequest,
	) (*ReportSignatures, error)
}
