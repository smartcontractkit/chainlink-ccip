package tokens

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLockReleasePoolLiquidityMigrationAmount(t *testing.T) {
	tests := []struct {
		name        string
		amount      LockReleasePoolLiquidityMigrationAmount
		wantRaw     string
		wantBP      *uint16
		expectedErr string
	}{
		{
			name:    "valid raw amount",
			amount:  LockReleasePoolLiquidityMigrationAmount{Format: LiquidityMigrationAmountFormatRAW, Value: "100"},
			wantRaw: "100",
		},
		{
			name:   "valid maximum basis points",
			amount: LockReleasePoolLiquidityMigrationAmount{Format: LiquidityMigrationAmountFormatBPS, Value: "10000"},
			wantBP: new(uint16(10000)),
		},
		{
			name:        "raw zero rejected",
			amount:      LockReleasePoolLiquidityMigrationAmount{Format: LiquidityMigrationAmountFormatRAW, Value: "0"},
			expectedErr: "must be positive",
		},
		{
			name:        "raw negative rejected",
			amount:      LockReleasePoolLiquidityMigrationAmount{Format: LiquidityMigrationAmountFormatRAW, Value: "-1"},
			expectedErr: "must be positive",
		},
		{
			name:        "raw non-numeric rejected",
			amount:      LockReleasePoolLiquidityMigrationAmount{Format: LiquidityMigrationAmountFormatRAW, Value: "1.5"},
			expectedErr: "invalid value for RawAmount",
		},
		{
			name:        "basis points zero rejected",
			amount:      LockReleasePoolLiquidityMigrationAmount{Format: LiquidityMigrationAmountFormatBPS, Value: "0"},
			expectedErr: "between 1 and 10000",
		},
		{
			name:        "basis points above 10000 rejected",
			amount:      LockReleasePoolLiquidityMigrationAmount{Format: LiquidityMigrationAmountFormatBPS, Value: "10001"},
			expectedErr: "between 1 and 10000",
		},
		{
			name:        "basis points non-numeric rejected",
			amount:      LockReleasePoolLiquidityMigrationAmount{Format: LiquidityMigrationAmountFormatBPS, Value: "abc"},
			expectedErr: "invalid value for BasisPoints",
		},
		{
			name:        "unknown format rejected",
			amount:      LockReleasePoolLiquidityMigrationAmount{Format: LiquidityMigrationAmountFormat("bogus"), Value: "100"},
			expectedErr: "invalid format",
		},
		{
			name:        "empty value rejected",
			amount:      LockReleasePoolLiquidityMigrationAmount{Format: LiquidityMigrationAmountFormatBPS, Value: ""},
			expectedErr: "value must be provided",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.amount.Validate()
			if err == nil {
				_, err = tc.amount.BasisPoints()
			}
			if err == nil {
				_, err = tc.amount.RawAmount()
			}

			if tc.expectedErr == "" {
				require.NoError(t, err)
				if tc.wantRaw != "" {
					raw, rawErr := tc.amount.RawAmount()
					require.NoError(t, rawErr)
					require.NotNil(t, raw)
					require.Equal(t, tc.wantRaw, raw.String())
				}
				if tc.wantBP != nil {
					bp, bpErr := tc.amount.BasisPoints()
					require.NoError(t, bpErr)
					require.NotNil(t, bp)
					require.Equal(t, *tc.wantBP, *bp)
				}
				return
			}
			require.Error(t, err)
			require.Contains(t, err.Error(), tc.expectedErr)
		})
	}
}
