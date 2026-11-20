package adapters

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"

	type_and_version "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/type_and_version"
	tokensapi "github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	cciputils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations/optest"
)

// baseSourceInput returns a valid exact-mode input that passes Validate. Callers mutate it to build
// the invalid combinations under test.
func baseSourceInput() tokensapi.MigrateLockReleasePoolLiquidityInput {
	return tokensapi.MigrateLockReleasePoolLiquidityInput{
		ChainSelector:   1,
		OldPoolAddress:  common.HexToAddress("0x0000000000000000000000000000000000000001").Hex(),
		NewPoolAddress:  common.HexToAddress("0x0000000000000000000000000000000000000002").Hex(),
		TimelockAddress: common.HexToAddress("0x0000000000000000000000000000000000000003").Hex(),
		SiloExactAmounts: []tokensapi.SiloExactAmount{
			{ChainSelector: 1, Amount: big.NewInt(1)},
		},
	}
}

// executeSourceSequence runs the adapter's public source sequence. The sequence validates the input
// before any chain access, so invalid inputs return before the zero BlockChains is touched.
func executeSourceSequence(t *testing.T, input tokensapi.MigrateLockReleasePoolLiquidityInput) error {
	t.Helper()

	seq := NewTokenAdapter().MigrateLockReleasePoolLiquiditySequence()
	_, err := cldf_ops.ExecuteSequence(optest.NewBundle(t), seq, cldf_chain.BlockChains{}, input)
	return err
}

// TestMigrateSiloedLockReleaseSource_RejectsSharedCollateralWithoutSkip pins the guard that stops
// the custom XERC20 group pool's unsiloed balance (shared lockbox collateral backing the XERC20
// token) from being migrated. The guard runs before any chain call, so it is testable without
// custom bytecode.
func TestMigrateSiloedLockReleaseSource_RejectsSharedCollateralWithoutSkip(t *testing.T) {
	t.Parallel()

	adapter := NewTokenAdapter()
	oldPool := common.HexToAddress("0x0000000000000000000000000000000000000001")
	newPool := common.HexToAddress("0x0000000000000000000000000000000000000002")
	timelock := common.HexToAddress("0x0000000000000000000000000000000000000003")

	oldTV := type_and_version.TypeAndVersion{
		Type:    cciputils.SiloedWithUnsiloedXERC20GroupTokenPool,
		Version: cciputils.Version_1_6_0,
	}

	_, err := adapter.migrateSiloedLockReleaseSource(
		cldf_ops.Bundle{}, evm.Chain{}, baseSourceInput(), oldTV, oldPool, newPool, timelock,
	)
	require.Error(t, err)
	require.Contains(t, err.Error(), "shared XERC20LockBox collateral")
	require.Contains(t, err.Error(), "SkipUnsiloed")
}

// TestMigrateSiloedLockReleaseSource_EnforcesInputContract pins that the public source sequence
// enforces MigrateLockReleasePoolLiquidityInput.Validate for direct callers, not just the composer.
// Every case here is an input the composer would reject; the sequence must reject it too, before any
// chain access, so a direct caller cannot drive a contradictory or unsafe migration.
func TestMigrateSiloedLockReleaseSource_EnforcesInputContract(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		mutate      func(*tokensapi.MigrateLockReleasePoolLiquidityInput)
		expectedErr string
	}{
		{
			// The guard only rejected legacy mode when SiloExactAmounts was empty; with exact mode
			// present the legacy amount used to be silently ignored.
			name: "SkipUnsiloed with SiloExactAmounts and Amount",
			mutate: func(in *tokensapi.MigrateLockReleasePoolLiquidityInput) {
				in.Amount = big.NewInt(5)
				in.SkipUnsiloed = true
			},
			expectedErr: "mutually exclusive with Amount/BasisPoints",
		},
		{
			name: "SkipUnsiloed with SiloExactAmounts and BasisPoints",
			mutate: func(in *tokensapi.MigrateLockReleasePoolLiquidityInput) {
				in.BasisPoints = new(uint16(5000))
				in.SkipUnsiloed = true
			},
			expectedErr: "mutually exclusive with Amount/BasisPoints",
		},
		{
			name: "SkipUnsiloed outside exact mode",
			mutate: func(in *tokensapi.MigrateLockReleasePoolLiquidityInput) {
				in.SiloExactAmounts = nil
				in.Amount = big.NewInt(5)
				in.SkipUnsiloed = true
			},
			expectedErr: "SkipUnsiloed requires SiloExactAmounts",
		},
		{
			name: "SkipUnsiloed with UnsiloedExactAmount",
			mutate: func(in *tokensapi.MigrateLockReleasePoolLiquidityInput) {
				in.UnsiloedExactAmount = big.NewInt(5)
				in.SkipUnsiloed = true
			},
			expectedErr: "SkipUnsiloed and UnsiloedExactAmount are mutually exclusive",
		},
		{
			// Amount and BasisPoints are mutually exclusive even without exact mode.
			name: "Amount and BasisPoints",
			mutate: func(in *tokensapi.MigrateLockReleasePoolLiquidityInput) {
				in.SiloExactAmounts = nil
				in.Amount = big.NewInt(5)
				in.BasisPoints = new(uint16(5000))
			},
			expectedErr: "Amount and BasisPoints are mutually exclusive",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			input := baseSourceInput()
			tc.mutate(&input)

			err := executeSourceSequence(t, input)
			require.Error(t, err)
			require.Contains(t, err.Error(), tc.expectedErr)
		})
	}
}
