package libs

import (
	"testing"

	"github.com/stretchr/testify/assert"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
)

func TestGetChainInfoFromSelector(t *testing.T) {
	tests := []struct {
		selector       cciptypes.ChainSelector
		expectedFamily string
		expectedID     string
		expectValid    bool
	}{
		{
			selector:       5009297550715157269,
			expectedFamily: "evm",
			expectedID:     "1",
			expectValid:    true,
		},
		{
			selector:       16015286601757825753,
			expectedFamily: "evm",
			expectedID:     "11155111",
			expectValid:    true,
		},
		{
			selector:       124615329519749607,
			expectedFamily: "solana",
			expectedID:     "5eykt4UsFv8P8NJdTREpY1vzqKqZKvdpKuc147dw2N9d",
			expectValid:    true,
		},
		{
			selector:       4741433654826277614,
			expectedFamily: "aptos",
			expectedID:     "1",
			expectValid:    true,
		},
		{
			selector:       3,
			expectedFamily: "unknown",
			expectedID:     "unknown",
			expectValid:    false,
		},
	}

	for _, test := range tests {
		t.Run(test.expectedID, func(t *testing.T) {
			family, id, valid := GetChainInfoFromSelector(test.selector)
			assert.Equal(t, test.expectedFamily, family)
			assert.Equal(t, test.expectedID, id)
			assert.Equal(t, test.expectValid, valid)
		})
	}
}
