package tokens

import (
	"math/big"
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/stretchr/testify/require"
)

func TestMigrateLockReleasePoolLiquidity_VerifyPreconditions_ExactAmounts(t *testing.T) {
	baseMigration := func() LockReleasePoolMigration {
		return LockReleasePoolMigration{
			ChainSelector: 1,
			OldPoolRef: datastore.AddressRef{
				ChainSelector: 1,
				Address:       "0x0000000000000000000000000000000000000001",
				Version:       semver.MustParse("1.6.1"),
			},
			NewPoolRef: datastore.AddressRef{
				ChainSelector: 1,
				Address:       "0x0000000000000000000000000000000000000002",
				Version:       semver.MustParse("2.0.0"),
			},
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

func TestMigrateLockReleasePoolLiquidity_VerifyPreconditions_PoolRefs(t *testing.T) {
	oldRef := func() datastore.AddressRef {
		return datastore.AddressRef{
			ChainSelector: 1,
			Address:       "0x0000000000000000000000000000000000000001",
			Version:       semver.MustParse("1.6.1"),
		}
	}
	newRef := func() datastore.AddressRef {
		return datastore.AddressRef{
			ChainSelector: 1,
			Address:       "0x0000000000000000000000000000000000000002",
			Version:       semver.MustParse("2.0.0"),
		}
	}

	tests := []struct {
		name        string
		mutate      func(m *LockReleasePoolMigration)
		expectedErr string
	}{
		{
			name: "Success - valid legacy to v2 pair",
			mutate: func(m *LockReleasePoolMigration) {
			},
		},
		{
			name: "Failure - empty OldPoolRef",
			mutate: func(m *LockReleasePoolMigration) {
				m.OldPoolRef = datastore.AddressRef{}
			},
			expectedErr: "OldPoolRef is required",
		},
		{
			name: "Failure - empty NewPoolRef",
			mutate: func(m *LockReleasePoolMigration) {
				m.NewPoolRef = datastore.AddressRef{}
			},
			expectedErr: "NewPoolRef is required",
		},
		{
			name: "Failure - OldPoolRef chain selector mismatch",
			mutate: func(m *LockReleasePoolMigration) {
				ref := oldRef()
				ref.ChainSelector = 2
				m.OldPoolRef = ref
			},
			expectedErr: "OldPoolRef.ChainSelector 2 does not match the migration's ChainSelector 1",
		},
		{
			name: "Failure - NewPoolRef chain selector mismatch",
			mutate: func(m *LockReleasePoolMigration) {
				ref := newRef()
				ref.ChainSelector = 2
				m.NewPoolRef = ref
			},
			expectedErr: "NewPoolRef.ChainSelector 2 does not match the migration's ChainSelector 1",
		},
		{
			name: "Failure - reversed versions",
			mutate: func(m *LockReleasePoolMigration) {
				m.OldPoolRef = newRef()
				m.NewPoolRef = oldRef()
			},
			expectedErr: "OldPoolRef version 2.0.0 must be strictly older than NewPoolRef version 1.6.1",
		},
		{
			name: "Failure - equal versions",
			mutate: func(m *LockReleasePoolMigration) {
				m.NewPoolRef = oldRef()
			},
			expectedErr: "OldPoolRef version 1.6.1 must be strictly older than NewPoolRef version 1.6.1",
		},
		{
			name: "Success - version comparison skipped when a ref has no version",
			mutate: func(m *LockReleasePoolMigration) {
				ref := newRef()
				ref.Version = nil
				m.NewPoolRef = ref
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			migration := LockReleasePoolMigration{
				ChainSelector: 1,
				OldPoolRef:    oldRef(),
				NewPoolRef:    newRef(),
				LiquidityMigrationAmount: &LockReleasePoolLiquidityMigrationAmount{
					Format: LiquidityMigrationAmountFormatBPS,
					Value:  "5000",
				},
			}
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
						ChainSelector: 1,
						OldPoolRef: datastore.AddressRef{
							ChainSelector: 1,
							Address:       "0x0000000000000000000000000000000000000001",
							Version:       semver.MustParse("1.6.1"),
						},
						NewPoolRef: datastore.AddressRef{
							ChainSelector: 1,
							Address:       "0x0000000000000000000000000000000000000002",
							Version:       semver.MustParse("2.0.0"),
						},
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
