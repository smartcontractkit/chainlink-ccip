package commit

import (
	"testing"

	"github.com/stretchr/testify/assert"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
)

func Test_validateFChain(t *testing.T) {
	testCases := []struct {
		name   string
		fChain map[cciptypes.ChainSelector]int
		expErr bool
	}{
		{
			name: "FChain contains negative values",
			fChain: map[cciptypes.ChainSelector]int{
				1: 11,
				2: -4,
			},
			expErr: true,
		},
		{
			name: "FChain valid",
			fChain: map[cciptypes.ChainSelector]int{
				12: 6,
				7:  9,
			},
			expErr: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := validateFChain(tc.fChain)

			if tc.expErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
		})
	}
}
