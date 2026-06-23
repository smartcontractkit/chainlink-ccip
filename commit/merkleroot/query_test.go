package merkleroot

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-ccip/internal/plugintypes"
)

func TestProcessor_Query(t *testing.T) {
	ctx := t.Context()

	testCases := []struct {
		name        string
		prevOutcome Outcome
	}{
		{
			name:        "empty previous outcome",
			prevOutcome: Outcome{},
		},
		{
			name: "building report state",
			prevOutcome: Outcome{
				OutcomeType: ReportIntervalsSelected,
				RangesSelectedForReport: []plugintypes.ChainRange{
					{ChainSel: ccipocr3.ChainSelector(1), SeqNumRange: ccipocr3.NewSeqNumRange(10, 20)},
				},
			},
		},
		{
			name: "waiting for transmission",
			prevOutcome: Outcome{
				OutcomeType: ReportInFlight,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			p := Processor{
				lggr: logger.Test(t),
			}

			q, err := p.Query(ctx, tc.prevOutcome)
			require.NoError(t, err)
			require.Equal(t, Query{}, q)
		})
	}
}
