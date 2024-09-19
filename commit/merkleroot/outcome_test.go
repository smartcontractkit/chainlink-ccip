package merkleroot

import (
	"testing"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/stretchr/testify/require"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-ccip/plugintypes"
)

func Test_buildReport(t *testing.T) {
	t.Run("determinism check", func(t *testing.T) {
		const rounds = 50

		obs := ConsensusObservation{
			MerkleRoots: map[cciptypes.ChainSelector]cciptypes.MerkleRootChain{
				cciptypes.ChainSelector(1): {
					ChainSel:     1,
					SeqNumsRange: cciptypes.NewSeqNumRange(10, 20),
					MerkleRoot:   cciptypes.Bytes32{1},
				},
				cciptypes.ChainSelector(2): {
					ChainSel:     2,
					SeqNumsRange: cciptypes.NewSeqNumRange(20, 30),
					MerkleRoot:   cciptypes.Bytes32{2},
				},
			},
		}

		lggr := logger.Test(t)

		for i := 0; i < rounds; i++ {
			report1 := buildReport(Query{}, lggr, obs, Outcome{})
			report2 := buildReport(Query{}, lggr, obs, Outcome{})
			require.Equal(t, report1, report2)
		}
	})
}

func Test_reportRangesOutcome(t *testing.T) {
	lggr := logger.Test(t)

	testCases := []struct {
		name                 string
		consensusObservation ConsensusObservation
		merkleTreeSizeLimit  uint64
		expectedOutcome      Outcome
	}{
		{
			name: "base empty outcome",
			expectedOutcome: Outcome{
				OutcomeType:             ReportIntervalsSelected,
				RangesSelectedForReport: []plugintypes.ChainRange{},
				OffRampNextSeqNums:      []plugintypes.SeqNumChain{},
			},
		},
		{
			name: "simple scenario with one chain",
			consensusObservation: ConsensusObservation{
				OnRampMaxSeqNums: map[cciptypes.ChainSelector]cciptypes.SeqNum{
					1: 20,
				},
				OffRampNextSeqNums: map[cciptypes.ChainSelector]cciptypes.SeqNum{
					1: 18, // off ramp next is 18, on ramp max is 20 so new msgs are: [18, 19, 20]
				},
			},
			merkleTreeSizeLimit: 256, // default limit should be used
			expectedOutcome: Outcome{
				OutcomeType: ReportIntervalsSelected,
				RangesSelectedForReport: []plugintypes.ChainRange{
					{ChainSel: 1, SeqNumRange: cciptypes.NewSeqNumRange(18, 20)},
				},
				OffRampNextSeqNums: []plugintypes.SeqNumChain{
					{ChainSel: 1, SeqNum: 18},
				},
			},
		},
		{
			name: "simple scenario with one chain",
			consensusObservation: ConsensusObservation{
				OnRampMaxSeqNums: map[cciptypes.ChainSelector]cciptypes.SeqNum{
					1: 20,
					2: 1000,
					3: 10000,
				},
				OffRampNextSeqNums: map[cciptypes.ChainSelector]cciptypes.SeqNum{
					1: 18,  // off ramp next is 18, on ramp max is 20 so new msgs are: [18, 19, 20]
					2: 995, // off ramp next is 995, on ramp max is 1000 so new msgs are: [995, 996, 997, 998, 999, 1000]
					3: 500, // off ramp next is 500, we have new messages up to 10000 (default limit applied)
				},
			},
			merkleTreeSizeLimit: 5,
			expectedOutcome: Outcome{
				OutcomeType: ReportIntervalsSelected,
				RangesSelectedForReport: []plugintypes.ChainRange{
					{ChainSel: 1, SeqNumRange: cciptypes.NewSeqNumRange(18, 20)},
					{ChainSel: 2, SeqNumRange: cciptypes.NewSeqNumRange(995, 999)},
					{ChainSel: 3, SeqNumRange: cciptypes.NewSeqNumRange(500, 504)},
				},
				OffRampNextSeqNums: []plugintypes.SeqNumChain{
					{ChainSel: 1, SeqNum: 18},
					{ChainSel: 2, SeqNum: 995},
					{ChainSel: 3, SeqNum: 500},
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			outc := reportRangesOutcome(Query{}, lggr, tc.consensusObservation, tc.merkleTreeSizeLimit)
			require.Equal(t, tc.expectedOutcome, outc)
		})
	}
}
