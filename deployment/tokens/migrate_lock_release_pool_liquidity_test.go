package tokens

import (
	"math/big"
	"testing"

	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/stretchr/testify/require"
)

func TestMigrateLockReleasePoolLiquidity_VerifyPreconditions_ExactAmounts(t *testing.T) {
	baseMigration := func() LockReleasePoolMigration {
		return LockReleasePoolMigration{
			ChainSelector: 1,
			OldPoolRef:    datastore.AddressRef{},
			NewPoolRef:    datastore.AddressRef{},
		}
	}

	tests := []struct {
		name        string
		mutate      func(m *LockReleasePoolMigration)
		expectedErr string
	}{
		{
			name: "SiloExactAmounts and LiquidityMigrationAmount mutually exclusive",
			mutate: func(m *LockReleasePoolMigration) {
				m.LiquidityMigrationAmount = &LockReleasePoolLiquidityMigrationAmount{Format: LiquidityMigrationAmountFormatBPS, Value: "5000"}
				m.SiloExactAmounts = []SiloExactAmount{{ChainSelector: 2, Amount: big.NewInt(100)}}
			},
			expectedErr: "SiloExactAmounts/UnsiloedExactAmount are mutually exclusive with LiquidityMigrationAmount",
		},
		{
			name: "UnsiloedExactAmount and LiquidityMigrationAmount mutually exclusive",
			mutate: func(m *LockReleasePoolMigration) {
				m.LiquidityMigrationAmount = &LockReleasePoolLiquidityMigrationAmount{Format: LiquidityMigrationAmountFormatRAW, Value: "100"}
				m.UnsiloedExactAmount = big.NewInt(50)
			},
			expectedErr: "SiloExactAmounts/UnsiloedExactAmount are mutually exclusive with LiquidityMigrationAmount",
		},
		{
			name: "no migration mode provided",
			mutate: func(m *LockReleasePoolMigration) {
			},
			expectedErr: "one of LiquidityMigrationAmount or SiloExactAmounts/UnsiloedExactAmount must be provided",
		},
		{
			name: "duplicate ChainSelector in SiloExactAmounts",
			mutate: func(m *LockReleasePoolMigration) {
				m.SiloExactAmounts = []SiloExactAmount{
					{ChainSelector: 2, Amount: big.NewInt(100)},
					{ChainSelector: 2, Amount: big.NewInt(200)},
				}
			},
			expectedErr: "duplicate ChainSelector 2 in SiloExactAmounts",
		},
		{
			name: "nil SiloExactAmount.Amount",
			mutate: func(m *LockReleasePoolMigration) {
				m.SiloExactAmounts = []SiloExactAmount{{ChainSelector: 2, Amount: nil}}
			},
			expectedErr: "SiloExactAmounts[0].Amount must be positive",
		},
		{
			name: "negative UnsiloedExactAmount",
			mutate: func(m *LockReleasePoolMigration) {
				m.SiloExactAmounts = []SiloExactAmount{{ChainSelector: 2, Amount: big.NewInt(100)}}
				m.UnsiloedExactAmount = big.NewInt(-1)
			},
			expectedErr: "UnsiloedExactAmount must be positive",
		},
		{
			name: "UnsiloedExactAmount without SiloExactAmounts",
			mutate: func(m *LockReleasePoolMigration) {
				m.UnsiloedExactAmount = big.NewInt(50)
			},
			expectedErr: "UnsiloedExactAmount requires SiloExactAmounts to also be set",
		},
		{
			name: "zero SiloExactAmount.Amount is allowed",
			mutate: func(m *LockReleasePoolMigration) {
				m.SiloExactAmounts = []SiloExactAmount{{ChainSelector: 2, Amount: big.NewInt(0)}}
				m.UnsiloedExactAmount = big.NewInt(50)
			},
			expectedErr: "",
		},
		{
			name: "valid SiloExactAmounts and UnsiloedExactAmount",
			mutate: func(m *LockReleasePoolMigration) {
				m.SiloExactAmounts = []SiloExactAmount{{ChainSelector: 2, Amount: big.NewInt(100)}}
				m.UnsiloedExactAmount = big.NewInt(50)
			},
			expectedErr: "",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			migration := baseMigration()
			tc.mutate(&migration)

			cs := MigrateLockReleasePoolLiquidity(nil, nil)
			err := cs.VerifyPreconditions(cldf.Environment{}, MigrateLockReleasePoolLiquidityConfig{
				Migrations: []LockReleasePoolMigration{migration},
			})

			if tc.expectedErr == "" {
				require.NoError(t, err)
				return
			}
			require.Error(t, err)
			require.Contains(t, err.Error(), tc.expectedErr)
		})
	}
}

func TestMigrateLockReleasePoolLiquidity_VerifyPreconditions_AmountFormat(t *testing.T) {
	tests := []struct {
		name        string
		amount      *LockReleasePoolLiquidityMigrationAmount
		expectedErr string
	}{
		{
			name:   "valid raw amount",
			amount: &LockReleasePoolLiquidityMigrationAmount{Format: LiquidityMigrationAmountFormatRAW, Value: "100"},
		},
		{
			name:   "valid basis points",
			amount: &LockReleasePoolLiquidityMigrationAmount{Format: LiquidityMigrationAmountFormatBPS, Value: "5000"},
		},
		{
			name:        "out of range basis points",
			amount:      &LockReleasePoolLiquidityMigrationAmount{Format: LiquidityMigrationAmountFormatBPS, Value: "10001"},
			expectedErr: "must be between 1 and 10000",
		},
		{
			name:        "unknown format",
			amount:      &LockReleasePoolLiquidityMigrationAmount{Format: LiquidityMigrationAmountFormat("bogus"), Value: "100"},
			expectedErr: "invalid format",
		},
		{
			name:        "empty value",
			amount:      &LockReleasePoolLiquidityMigrationAmount{Format: LiquidityMigrationAmountFormatBPS, Value: ""},
			expectedErr: "value must be provided",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			cs := MigrateLockReleasePoolLiquidity(nil, nil)
			err := cs.VerifyPreconditions(cldf.Environment{}, MigrateLockReleasePoolLiquidityConfig{
				Migrations: []LockReleasePoolMigration{
					{
						ChainSelector:            1,
						OldPoolRef:               datastore.AddressRef{},
						NewPoolRef:               datastore.AddressRef{},
						LiquidityMigrationAmount: tc.amount,
					},
				},
			})

			if tc.expectedErr == "" {
				require.NoError(t, err)
				return
			}
			require.Error(t, err)
			require.Contains(t, err.Error(), tc.expectedErr)
		})
	}
}
