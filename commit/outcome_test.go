package commit

import (
	"testing"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
			GasPrices: map[cciptypes.ChainSelector]cciptypes.BigInt{
				cciptypes.ChainSelector(1): cciptypes.NewBigIntFromInt64(1000),
				cciptypes.ChainSelector(2): cciptypes.NewBigIntFromInt64(2000),
			},
			TokenPrices: map[types.Account]cciptypes.BigInt{
				types.Account("1"): cciptypes.NewBigIntFromInt64(1000),
				types.Account("2"): cciptypes.NewBigIntFromInt64(2000),
			},
		}

		for i := 0; i < rounds; i++ {
			report1 := buildReport(Query{}, obs, Outcome{})
			report2 := buildReport(Query{}, obs, Outcome{})
			require.Equal(t, report1, report2)
		}
	})
}

func Test_fChainConsensus(t *testing.T) {
	lggr := logger.Test(t)
	f := 3
	fChainValues := map[cciptypes.ChainSelector][]int{
		cciptypes.ChainSelector(1): {5, 5, 5, 5, 5},
		cciptypes.ChainSelector(2): {5, 5, 5, 3, 5},
		cciptypes.ChainSelector(3): {5, 5},             // not enough observations, must be observed at least f times.
		cciptypes.ChainSelector(4): {5, 3, 5, 3, 5, 3}, // both values appear at least f times
	}
	fChainFinal := fChainConsensus(lggr, f, fChainValues)

	assert.Equal(t, map[cciptypes.ChainSelector]int{
		cciptypes.ChainSelector(1): 5,
		cciptypes.ChainSelector(2): 5,
		cciptypes.ChainSelector(4): 3,
	}, fChainFinal)
}

func Test_mostFrequentElem(t *testing.T) {
	testCases := []struct {
		name            string
		input           []int
		expectedElem    int
		expectedCnt     int
		overrideOnEqual func(int, int) bool
	}{
		{
			name:         "empty",
			input:        []int{},
			expectedElem: 0,
			expectedCnt:  0,
		},
		{
			name:            "empty",
			input:           []int{},
			expectedElem:    0,
			expectedCnt:     0,
			overrideOnEqual: func(a, b int) bool { return a > b },
		},
		{
			name:            "happy path with override 1",
			input:           []int{1, 1, 1, 2, 2, 2, 3, 4, 3},
			expectedElem:    1,
			expectedCnt:     3,
			overrideOnEqual: func(curr, new int) bool { return new < curr },
		},
		{
			name:            "happy path with override 2",
			input:           []int{1, 1, 1, 2, 2, 2, 3, 4, 3},
			expectedElem:    2,
			expectedCnt:     3,
			overrideOnEqual: func(curr, new int) bool { return new > curr },
		},
		{
			name:            "happy path no overrides",
			input:           []int{1, 2, 2, 2, 3, 4, 3},
			expectedElem:    2,
			expectedCnt:     3,
			overrideOnEqual: func(curr, new int) bool { return new > curr },
		},
		{
			name:            "happy path no overrides",
			input:           []int{1, 2, 2, 2, 3, 4, 3},
			expectedElem:    2,
			expectedCnt:     3,
			overrideOnEqual: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, cnt := mostFrequentElem(tc.input, tc.overrideOnEqual)
			assert.Equal(t, tc.expectedElem, actual)
			assert.Equal(t, tc.expectedCnt, cnt)
		})
	}
}
