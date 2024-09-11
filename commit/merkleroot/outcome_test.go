package merkleroot

import (
	"testing"

	ct "github.com/smartcontractkit/chainlink-ccip/commit/committypes"
	"github.com/stretchr/testify/require"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
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

		for i := 0; i < rounds; i++ {
			report1 := buildReport(ct.MerkleRootQuery{}, obs, ct.MerkleRootOutcome{})
			report2 := buildReport(ct.MerkleRootQuery{}, obs, ct.MerkleRootOutcome{})
			require.Equal(t, report1, report2)
		}
	})
}
