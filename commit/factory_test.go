package commit

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/commit/chainfee"
	"github.com/smartcontractkit/chainlink-ccip/commit/merkleroot"
	"github.com/smartcontractkit/chainlink-ccip/commit/merkleroot/rmn"
	"github.com/smartcontractkit/chainlink-ccip/commit/merkleroot/rmn/rmnpb"
	"github.com/smartcontractkit/chainlink-ccip/commit/tokenprice"
)

func Test_maxQueryLength(t *testing.T) {
	// This test will verify that the maxQueryLength constant is set to a proper value.

	// Estimate the maximum number of source chains we are going to ever have.
	// This value should be tweaked after we are close to supporting that many chains.
	const estimateMaxNumberOfSourceChains = 1000

	// Estimate the maximum number of RMN report signers we are going to ever have.
	// This value is defined in RMNRemote contract as `f`.
	// This value should be tweaked if necessary in order to define new limits.
	const estimatedMaxRmnReportSigners = 200

	sigs := make([]*rmnpb.EcdsaSignature, estimatedMaxRmnReportSigners)
	for i := range sigs {
		sigs[i] = &rmnpb.EcdsaSignature{R: make([]byte, 32), S: make([]byte, 32)}
	}

	laneUpdates := make([]*rmnpb.FixedDestLaneUpdate, estimateMaxNumberOfSourceChains)
	for i := range laneUpdates {
		laneUpdates[i] = &rmnpb.FixedDestLaneUpdate{
			LaneSource: &rmnpb.LaneSource{
				SourceChainSelector: math.MaxUint64,
				OnrampAddress:       make([]byte, 40),
			},
			ClosedInterval: &rmnpb.ClosedInterval{
				MinMsgNr: math.MaxUint64,
				MaxMsgNr: math.MaxUint64,
			},
			Root: make([]byte, 32),
		}
	}

	q := Query{
		MerkleRootQuery: merkleroot.Query{
			RetryRMNSignatures: true,
			RMNSignatures: &rmn.ReportSignatures{
				Signatures:  sigs,
				LaneUpdates: laneUpdates,
			},
		},
		TokenPriceQuery: tokenprice.Query{},
		ChainFeeQuery:   chainfee.Query{},
	}
	b, err := q.Encode()
	require.NoError(t, err)

	// We set twice the size, for extra safety while making breaking changes between oracle versions.
	assert.Equal(t, 2*len(b), maxQueryLength)
}
