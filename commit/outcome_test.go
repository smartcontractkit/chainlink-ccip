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
		cciptypes.ChainSelector(4): {5, 3, 5, 3, 5, 3}, // both values appear at least f times, no consensus
	}
	fChainFinal := fChainConsensus(lggr, f, fChainValues)

	assert.Equal(t, map[cciptypes.ChainSelector]int{
		cciptypes.ChainSelector(1): 5,
		cciptypes.ChainSelector(2): 5,
	}, fChainFinal)
}

func Test_mostFrequentElement(t *testing.T) {
	testCases := []struct {
		name         string
		input        []int
		expectedElem int
		expectedCnt  int
		expErr       bool
	}{
		{
			name:         "empty",
			input:        []int{},
			expectedElem: 0,
			expectedCnt:  0,
			expErr:       true,
		},
		{
			name:         "empty",
			input:        []int{},
			expectedElem: 0,
			expectedCnt:  0,
			expErr:       true,
		},
		{
			name:         "base",
			input:        []int{33},
			expectedElem: 33,
			expectedCnt:  1,
		},
		{
			name:         "no major elem, 1 and 2 appear same number of times",
			input:        []int{1, 1, 1, 2, 2, 2, 3, 4, 3},
			expectedElem: 0,
			expectedCnt:  0,
			expErr:       true,
		},
		{
			name:         "happy path no overrides",
			input:        []int{1, 2, 2, 2, 3, 4, 3},
			expectedElem: 2,
			expectedCnt:  3,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, cnt, err := mostFrequentElement(tc.input)
			if tc.expErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tc.expectedElem, actual)
			assert.Equal(t, tc.expectedCnt, cnt)

		})
	}
}
