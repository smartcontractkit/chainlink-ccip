package shared

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
)

// Test_fChainConsensus is a legacy test that was replaced by GetConsensusMap, now it's a regression test.
func Test_fChainConsensus(t *testing.T) {
	lggr := logger.Test(t)
	f := 3
	fChainValues := map[cciptypes.ChainSelector][]int{
		cciptypes.ChainSelector(1): {5, 5, 5, 5, 5},
		cciptypes.ChainSelector(2): {5, 5, 5, 3, 5},
		cciptypes.ChainSelector(3): {5, 5},             // not enough observations, must be observed at least f times.
		cciptypes.ChainSelector(4): {5, 3, 5, 3, 5, 3}, // both values appear at least f times, no consensus
	}

	minObs := map[cciptypes.ChainSelector]int{
		cciptypes.ChainSelector(1): f,
		cciptypes.ChainSelector(2): f,
		cciptypes.ChainSelector(3): f,
		cciptypes.ChainSelector(4): f,
	}
	fChainFinal := GetConsensusMap(lggr, "fChain", fChainValues, minObs)

	assert.Equal(t, map[cciptypes.ChainSelector]int{
		cciptypes.ChainSelector(1): 5,
		cciptypes.ChainSelector(2): 5,
	}, fChainFinal)
}
