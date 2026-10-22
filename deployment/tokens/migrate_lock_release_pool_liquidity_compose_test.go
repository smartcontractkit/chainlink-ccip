package tokens

import (
	"context"
	"fmt"
	"math/big"
	"testing"

	"github.com/Masterminds/semver/v3"
	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"

	"github.com/smartcontractkit/chainlink-ccip/deployment/deploy"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
)

// migrationComposeTestAdapter is a TokenAdapter whose only meaningful method is the liquidity
// migration sequence; everything else comes from reverseTestAdapter.
type migrationComposeTestAdapter struct {
	reverseTestAdapter
	seq *cldf_ops.Sequence[MigrateLockReleasePoolLiquidityInput, sequences.OnChainOutput, cldf_chain.BlockChains]
}

func (a *migrationComposeTestAdapter) MigrateLockReleasePoolLiquiditySequence() *cldf_ops.Sequence[MigrateLockReleasePoolLiquidityInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return a.seq
}

// testSourceSequence records that it ran, resolves amount into input.Resolved, and emits one
// transaction addressed to `to`.
func testSourceSequence(id string, amount int64, to string, ran *bool) *cldf_ops.Sequence[MigrateLockReleasePoolLiquidityInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return cldf_ops.NewSequence(id, semver.MustParse("1.0.0"), "test source",
		func(_ cldf_ops.Bundle, _ cldf_chain.BlockChains, in MigrateLockReleasePoolLiquidityInput) (sequences.OnChainOutput, error) {
			*ran = true
			if in.Resolved != nil {
				in.Resolved.Amount = big.NewInt(amount)
			}
			return sequences.OnChainOutput{BatchOps: []mcms_types.BatchOperation{{
				ChainSelector: mcms_types.ChainSelector(in.ChainSelector),
				Transactions:  []mcms_types.Transaction{{To: to}},
			}}}, nil
		})
}

// testSinkSequence records the input it received and emits one transaction addressed to "SINK".
func testSinkSequence(received *MigrateLockReleasePoolLiquidityInput) *cldf_ops.Sequence[MigrateLockReleasePoolLiquidityInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return cldf_ops.NewSequence("test:migration-sink", semver.MustParse("2.0.0"), "test sink",
		func(_ cldf_ops.Bundle, _ cldf_chain.BlockChains, in MigrateLockReleasePoolLiquidityInput) (sequences.OnChainOutput, error) {
			*received = in
			return sequences.OnChainOutput{BatchOps: []mcms_types.BatchOperation{{
				ChainSelector: mcms_types.ChainSelector(in.ChainSelector),
				Transactions:  []mcms_types.Transaction{{To: "SINK"}},
			}}}, nil
		})
}

// newMigrationComposeTestEnv returns an environment whose datastore records OLD_POOL at oldPoolVersion,
// so the composer resolves the old pool's version from the datastore.
func newMigrationComposeTestEnv(t *testing.T, oldPoolVersion *semver.Version) cldf.Environment {
	t.Helper()
	family, err := chainsel.GetSelectorFamily(reverseTestLocalSel)
	require.NoError(t, err)
	deploy.GetAddressNormalizerRegistry().RegisterAddressNormalizer(family, reverseTestIdentityNormalizer{})

	ds := datastore.NewMemoryDataStore()
	require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
		ChainSelector: reverseTestLocalSel,
		Address:       "OLD_POOL",
		Type:          datastore.ContractType("LockReleaseTokenPool"),
		Version:       oldPoolVersion,
	}))

	lggr := logger.Test(t)
	return cldf.Environment{
		Logger:           lggr,
		OperationsBundle: cldf_ops.NewBundle(func() context.Context { return t.Context() }, lggr, cldf_ops.NewMemoryReporter()),
		DataStore:        ds.Seal(),
		BlockChains:      cldf_chain.NewBlockChains(map[uint64]cldf_chain.BlockChain{}),
	}
}

func migrationComposeTestInput() MigrateLockReleasePoolLiquidityInput {
	bps := uint16(5000)
	return MigrateLockReleasePoolLiquidityInput{
		ChainSelector:   reverseTestLocalSel,
		OldPoolAddress:  "OLD_POOL",
		NewPoolAddress:  "NEW_POOL",
		TimelockAddress: "TIMELOCK",
		BasisPoints:     &bps,
	}
}

func TestRunLockReleaseLiquidityMigration_ComposesSourceAndSink(t *testing.T) {
	family, err := chainsel.GetSelectorFamily(reverseTestLocalSel)
	require.NoError(t, err)
	v161 := semver.MustParse("1.6.1")
	v151 := semver.MustParse("1.5.1")

	var ran161, ran151 bool
	var sinkInput MigrateLockReleasePoolLiquidityInput
	reg := newTokenAdapterRegistry()
	reg.RegisterTokenAdapter(family, v161, &migrationComposeTestAdapter{seq: testSourceSequence("test:source-1.6.1", 42, "SOURCE_1_6_1", &ran161)})
	reg.RegisterTokenAdapter(family, v151, &migrationComposeTestAdapter{seq: testSourceSequence("test:source-1.5.1", 7, "SOURCE_1_5_1", &ran151)})
	reg.RegisterTokenAdapter(family, reverseTestV2_0_0, &migrationComposeTestAdapter{seq: testSinkSequence(&sinkInput)})

	out, _, err := runLockReleaseLiquidityMigration(newMigrationComposeTestEnv(t, v151), reg, reverseTestLocalSel, migrationComposeTestInput())
	require.NoError(t, err)

	// The source is chosen by the OLD pool's version, never the new pool's.
	require.True(t, ran151, "v1.5.1 source should run for a v1.5.1 old pool")
	require.False(t, ran161, "v1.6.1 source must not run for a v1.5.1 old pool")

	// The sink receives exactly the amount the source resolved, with basis points cleared.
	require.Equal(t, big.NewInt(7), sinkInput.Amount)
	require.Nil(t, sinkInput.BasisPoints)
	require.Nil(t, sinkInput.Resolved)
	require.Equal(t, "NEW_POOL", sinkInput.NewPoolAddress)

	// Source and sink land in one atomic batch, source first.
	require.Len(t, out.BatchOps, 1)
	require.Equal(t, []mcms_types.Transaction{{To: "SOURCE_1_5_1"}, {To: "SINK"}}, out.BatchOps[0].Transactions)
}

func TestRunLockReleaseLiquidityMigration_UnsupportedOldPool(t *testing.T) {
	family, err := chainsel.GetSelectorFamily(reverseTestLocalSel)
	require.NoError(t, err)
	v150 := semver.MustParse("1.5.0")

	tests := []struct {
		name           string
		oldPoolVersion *semver.Version
		expectedErr    string
	}{
		{
			name:           "adapter without a source",
			oldPoolVersion: v150,
			expectedErr:    fmt.Sprintf("old pool OLD_POOL (chain %d) does not support liquidity migration yet", reverseTestLocalSel),
		},
		{
			name:           "no adapter registered at the old pool's version",
			oldPoolVersion: semver.MustParse("1.6.0"), // not registered in this test's registry
			expectedErr:    "does not support liquidity migration yet",
		},
		{
			name:           "old pool is not a legacy pool",
			oldPoolVersion: reverseTestV2_0_0,
			expectedErr:    "the migration source must be older than v2.0.0",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var sinkInput MigrateLockReleasePoolLiquidityInput
			reg := newTokenAdapterRegistry()
			reg.RegisterTokenAdapter(family, v150, &migrationComposeTestAdapter{}) // nil sequence
			reg.RegisterTokenAdapter(family, reverseTestV2_0_0, &migrationComposeTestAdapter{seq: testSinkSequence(&sinkInput)})

			_, _, err := runLockReleaseLiquidityMigration(newMigrationComposeTestEnv(t, tc.oldPoolVersion), reg, reverseTestLocalSel, migrationComposeTestInput())
			require.ErrorContains(t, err, tc.expectedErr)
			require.Empty(t, sinkInput.NewPoolAddress, "sink must not run when the source is unsupported")
		})
	}
}

func TestRunLockReleaseLiquidityMigration_SourceResolvedNothing(t *testing.T) {
	family, err := chainsel.GetSelectorFamily(reverseTestLocalSel)
	require.NoError(t, err)
	v161 := semver.MustParse("1.6.1")

	// A source that never writes Resolved stands in for a report replayed from the reporter cache.
	silentSource := cldf_ops.NewSequence("test:silent-source", semver.MustParse("1.0.0"), "test source",
		func(_ cldf_ops.Bundle, _ cldf_chain.BlockChains, _ MigrateLockReleasePoolLiquidityInput) (sequences.OnChainOutput, error) {
			return sequences.OnChainOutput{}, nil
		})
	var sinkInput MigrateLockReleasePoolLiquidityInput
	reg := newTokenAdapterRegistry()
	reg.RegisterTokenAdapter(family, v161, &migrationComposeTestAdapter{seq: silentSource})
	reg.RegisterTokenAdapter(family, reverseTestV2_0_0, &migrationComposeTestAdapter{seq: testSinkSequence(&sinkInput)})

	_, _, err = runLockReleaseLiquidityMigration(newMigrationComposeTestEnv(t, v161), reg, reverseTestLocalSel, migrationComposeTestInput())
	require.ErrorContains(t, err, "resolved no amounts")
	require.Empty(t, sinkInput.NewPoolAddress, "sink must not run when the source resolved nothing")
}

func TestRunLockReleaseLiquidityMigration_InvalidInput(t *testing.T) {
	input := migrationComposeTestInput()
	input.Amount = big.NewInt(1) // mutually exclusive with BasisPoints

	_, _, err := runLockReleaseLiquidityMigration(cldf.Environment{}, newTokenAdapterRegistry(), reverseTestLocalSel, input)
	require.ErrorContains(t, err, "Amount and BasisPoints are mutually exclusive")
}

func TestMergeBatchOpsByChain(t *testing.T) {
	merged := mergeBatchOpsByChain([]mcms_types.BatchOperation{
		{ChainSelector: 1, Transactions: []mcms_types.Transaction{{To: "a"}}},
		{ChainSelector: 2, Transactions: []mcms_types.Transaction{{To: "b"}}},
		{}, // every write executed directly
		{ChainSelector: 1, Transactions: []mcms_types.Transaction{{To: "c"}}},
	})
	require.Equal(t, []mcms_types.BatchOperation{
		{ChainSelector: 1, Transactions: []mcms_types.Transaction{{To: "a"}, {To: "c"}}},
		{ChainSelector: 2, Transactions: []mcms_types.Transaction{{To: "b"}}},
	}, merged)
}
