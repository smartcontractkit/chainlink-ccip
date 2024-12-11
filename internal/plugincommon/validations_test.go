package plugincommon

import (
	"testing"

	"github.com/stretchr/testify/require"

	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

func TestValidateFChain(t *testing.T) {
	tests := []struct {
		name       string
		fChain     map[cciptypes.ChainSelector]int
		expectErr  bool
		errMessage string
	}{
		{
			name: "Valid fChain with all positive values",
			fChain: map[cciptypes.ChainSelector]int{
				1: 10,
				2: 20,
				3: 30,
			},
			expectErr: false,
		},
		{
			name: "Invalid fChain with a zero value",
			fChain: map[cciptypes.ChainSelector]int{
				1: 0,
				2: 20,
			},
			expectErr:  true,
			errMessage: "fChain for chain 1 is not positive: 0",
		},
		{
			name: "Invalid fChain with a negative value",
			fChain: map[cciptypes.ChainSelector]int{
				1: -10,
				2: 20,
			},
			expectErr:  true,
			errMessage: "fChain for chain 1 is not positive: -10",
		},
		{
			name:      "Empty fChain map",
			fChain:    map[cciptypes.ChainSelector]int{},
			expectErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateFChain(tt.fChain)

			if tt.expectErr {
				require.Error(t, err)
				require.Contains(t, err.Error(), tt.errMessage)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
