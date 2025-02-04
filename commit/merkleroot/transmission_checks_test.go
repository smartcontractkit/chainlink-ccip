package merkleroot

import (
	"fmt"
	"testing"

	"github.com/smartcontractkit/chainlink-common/pkg/utils/tests"
	"github.com/stretchr/testify/assert"

	"github.com/smartcontractkit/chainlink-ccip/internal/plugintypes"
	readermock "github.com/smartcontractkit/chainlink-ccip/mocks/pkg/reader"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

func Test_validateMerkleRootsState(t *testing.T) {
	testCases := []struct {
		name                 string
		onRampNextSeqNum     []plugintypes.SeqNumChain
		offRampExpNextSeqNum map[cciptypes.ChainSelector]cciptypes.SeqNum
		readerErr            error
		expErr               bool
	}{
		{
			name: "happy path",
			onRampNextSeqNum: []plugintypes.SeqNumChain{
				plugintypes.NewSeqNumChain(10, 100),
				plugintypes.NewSeqNumChain(20, 200),
			},
			offRampExpNextSeqNum: map[cciptypes.ChainSelector]cciptypes.SeqNum{10: 100, 20: 200},
			expErr:               false,
		},
		{
			name: "one root is stale",
			onRampNextSeqNum: []plugintypes.SeqNumChain{
				plugintypes.NewSeqNumChain(10, 100),
				plugintypes.NewSeqNumChain(20, 200),
			},
			// <- 200 is already on chain
			offRampExpNextSeqNum: map[cciptypes.ChainSelector]cciptypes.SeqNum{10: 100, 20: 201},
			expErr:               true,
		},
		{
			name: "one root has gap",
			onRampNextSeqNum: []plugintypes.SeqNumChain{
				plugintypes.NewSeqNumChain(10, 101), // <- onchain 99 but we submit 101 instead of 100
				plugintypes.NewSeqNumChain(20, 200),
			},
			offRampExpNextSeqNum: map[cciptypes.ChainSelector]cciptypes.SeqNum{10: 100, 20: 200},
			expErr:               true,
		},
		{
			name: "reader returned wrong number of seq nums, should be ok",
			onRampNextSeqNum: []plugintypes.SeqNumChain{
				plugintypes.NewSeqNumChain(10, 100),
				plugintypes.NewSeqNumChain(20, 200),
			},
			offRampExpNextSeqNum: map[cciptypes.ChainSelector]cciptypes.SeqNum{10: 100, 20: 200, 30: 300},
			expErr:               false,
		},
		{
			name: "reader error",
			onRampNextSeqNum: []plugintypes.SeqNumChain{
				plugintypes.NewSeqNumChain(10, 100),
				plugintypes.NewSeqNumChain(20, 200),
			},
			offRampExpNextSeqNum: map[cciptypes.ChainSelector]cciptypes.SeqNum{10: 100, 20: 200},
			readerErr:            fmt.Errorf("reader error"),
			expErr:               true,
		},
	}

	ctx := tests.Context(t)
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			reader := readermock.NewMockCCIPReader(t)
			rep := cciptypes.CommitPluginReport{}
			chains := make([]cciptypes.ChainSelector, 0, len(tc.onRampNextSeqNum))
			for _, snc := range tc.onRampNextSeqNum {
				rep.BlessedMerkleRoots = append(rep.BlessedMerkleRoots, cciptypes.MerkleRootChain{
					ChainSel:     snc.ChainSel,
					SeqNumsRange: cciptypes.NewSeqNumRange(snc.SeqNum, snc.SeqNum+10),
				})
				chains = append(chains, snc.ChainSel)
			}
			reader.EXPECT().NextSeqNum(ctx, chains).Return(tc.offRampExpNextSeqNum, tc.readerErr)

			err := ValidateMerkleRootsState(ctx, rep.BlessedMerkleRoots, rep.UnblessedMerkleRoots, reader)
			if tc.expErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
		})
	}
}
