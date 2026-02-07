package sequences_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	sequence1_7 "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/sequences"
)

func TestFeeQuoterUpdate_IsEmpty(t *testing.T) {
	empty := sequence1_7.FeeQuoterUpdate{}
	isEmpty, err := empty.IsEmpty()
	require.NoError(t, err)
	require.True(t, isEmpty, "Empty FeeQuoterUpdate should return true")

	nonEmpty := sequence1_7.FeeQuoterUpdate{
		ChainSelector: 1,
	}
	isEmpty, err = nonEmpty.IsEmpty()
	require.NoError(t, err)
	require.False(t, isEmpty, "Non-empty FeeQuoterUpdate should return false")
}
