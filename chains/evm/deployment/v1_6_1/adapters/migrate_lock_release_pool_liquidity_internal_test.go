package adapters

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"

	type_and_version "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/type_and_version"
	tokensapi "github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	cciputils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

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

	input := tokensapi.MigrateLockReleasePoolLiquidityInput{
		ChainSelector:   1,
		OldPoolAddress:  oldPool.Hex(),
		NewPoolAddress:  newPool.Hex(),
		TimelockAddress: timelock.Hex(),
		SiloExactAmounts: []tokensapi.SiloExactAmount{
			{ChainSelector: 1, Amount: big.NewInt(1)},
		},
	}
	oldTV := type_and_version.TypeAndVersion{
		Type:    cciputils.SiloedWithUnsiloedXERC20GroupTokenPool,
		Version: cciputils.Version_1_6_0,
	}

	_, err := adapter.migrateSiloedLockReleaseSource(
		cldf_ops.Bundle{}, evm.Chain{}, input, oldTV, oldPool, newPool, timelock,
	)
	require.Error(t, err)
	require.Contains(t, err.Error(), "shared XERC20LockBox collateral")
	require.Contains(t, err.Error(), "SkipUnsiloed")
}

// TestMigrateSiloedLockReleaseSource_RejectsSkipUnsiloedWithLegacyAmount pins the source-level
// invariant: SkipUnsiloed only takes effect in exact mode, so a legacy Amount/BasisPoints would
// otherwise fall through to the proportional branch and withdraw the shared bucket. The guard runs
// before any chain call, so it is testable without a chain.
func TestMigrateSiloedLockReleaseSource_RejectsSkipUnsiloedWithLegacyAmount(t *testing.T) {
	t.Parallel()

	adapter := NewTokenAdapter()
	oldPool := common.HexToAddress("0x0000000000000000000000000000000000000001")
	newPool := common.HexToAddress("0x0000000000000000000000000000000000000002")
	timelock := common.HexToAddress("0x0000000000000000000000000000000000000003")

	input := tokensapi.MigrateLockReleasePoolLiquidityInput{
		ChainSelector:   1,
		OldPoolAddress:  oldPool.Hex(),
		NewPoolAddress:  newPool.Hex(),
		TimelockAddress: timelock.Hex(),
		Amount:          big.NewInt(1),
		SkipUnsiloed:    true,
	}
	oldTV := type_and_version.TypeAndVersion{
		Type:    cciputils.SiloedWithUnsiloedXERC20GroupTokenPool,
		Version: cciputils.Version_1_6_0,
	}

	_, err := adapter.migrateSiloedLockReleaseSource(
		cldf_ops.Bundle{}, evm.Chain{}, input, oldTV, oldPool, newPool, timelock,
	)
	require.Error(t, err)
	require.Contains(t, err.Error(), "SkipUnsiloed requires SiloExactAmounts")
}
