package commit

import (
	"testing"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/types"
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
