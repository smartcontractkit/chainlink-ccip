package adapters

import (
	"testing"

	"github.com/stretchr/testify/require"
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
			current:    usdcLockOrBurnMechanismInvalid,
			wantUpdate: false,
		},
		{
			name:       "updates cctp v1",
			current:    usdcLockOrBurnMechanismCCTPV1,
			wantUpdate: true,
		},
		{
			name:       "updates cctp v2",
			current:    usdcLockOrBurnMechanismCCTPV2,
			wantUpdate: true,
		},
		{
			name:       "skips lock release",
			current:    usdcLockOrBurnMechanismLockRelease,
			wantUpdate: false,
		},
		{
			name:       "skips ccv",
			current:    usdcLockOrBurnMechanismCCV,
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
