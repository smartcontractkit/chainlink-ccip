package adapters

import (
	"testing"

	"github.com/stretchr/testify/require"

	cctpseq "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/sequences/cctp"
)

func TestShouldUpdateUSDCLockOrBurnMechanismToCCTPV2WithCCV(t *testing.T) {
	tests := []struct {
		name          string
		current       uint8
		wantUpdate    bool
		wantErrSubstr string
	}{
		{
			name:       "skips invalid mechanism",
			current:    cctpseq.LockOrBurnMechanismInvalid,
			wantUpdate: false,
		},
		{
			name:       "updates cctp v1",
			current:    cctpseq.LockOrBurnMechanismCCTPV1,
			wantUpdate: true,
		},
		{
			name:       "updates cctp v2",
			current:    cctpseq.LockOrBurnMechanismCCTPV2,
			wantUpdate: true,
		},
		{
			name:       "skips lock release",
			current:    cctpseq.LockOrBurnMechanismLockRelease,
			wantUpdate: false,
		},
		{
			name:       "skips ccv",
			current:    cctpseq.LockOrBurnMechanismCCV,
			wantUpdate: false,
		},
		{
			name:          "errors on unexpected mechanism",
			current:       5,
			wantErrSubstr: "unexpected mechanism 5",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := shouldUpdateUSDCLockOrBurnMechanismToCCTPV2WithCCV(test.current)
			if test.wantErrSubstr != "" {
				require.ErrorContains(t, err, test.wantErrSubstr)
				return
			}

			require.NoError(t, err)
			require.Equal(t, test.wantUpdate, got)
		})
	}
}
