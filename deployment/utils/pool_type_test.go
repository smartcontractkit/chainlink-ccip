package utils

import (
	"testing"

	"github.com/stretchr/testify/require"

	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

// TestPoolTypePredicates_V1_5_0AndProxyTypes pins the deliberate asymmetry between the two
// predicates for the v1.5.0 *AndProxy pool types.
//
// BurnMintTokenPoolAndProxy IS a burn-mint type because EVMPoolAdapter.TidyTokenPoolRoles gates
// the mint/burn role grant on IsBurnMintPoolType, and a v1.5.0 burn-mint pool needs that grant.
//
// LockReleaseTokenPoolAndProxy is NOT a lock-release type because there is no equivalent
// motivating call site: IsLockReleasePoolType is consulted in exactly one place, the v2.0.0
// DeployTokenPool dispatch, which has no v1.5.0 contract to deploy. Including it there would
// route the request into DeployLockReleaseTokenPool, which deploys an ERC20LockBox and an
// AdvancedPoolHooks before failing on the missing v1.5.0 bytecode - leaving both orphaned
// on-chain instead of failing cleanly on the switch's default case.
//
// v1.5.0 pools are deployed by the v1.5.0 DeployTokenPool sequence, which keys on the full
// "Type Version" string and never consults either predicate.
func TestPoolTypePredicates_V1_5_0AndProxyTypes(t *testing.T) {
	t.Parallel()

	require.True(t, IsBurnMintPoolType(BurnMintTokenPoolAndProxy.String()),
		"BurnMintTokenPoolAndProxy must stay a burn-mint type or the v1.5.0 pool silently loses its mint/burn role grant")

	require.False(t, IsLockReleasePoolType(LockReleaseTokenPoolAndProxy.String()),
		"LockReleaseTokenPoolAndProxy must NOT be a lock-release type: the only consumer is the v2.0.0 "+
			"deploy dispatch, which would deploy a lockbox and hooks before failing on the missing bytecode")
}

// TestPoolTypePredicates_StandardTypes guards the members that the v2.0.0 deploy dispatch relies
// on, so the fix above cannot be "simplified" into dropping real types.
func TestPoolTypePredicates_StandardTypes(t *testing.T) {
	t.Parallel()

	for _, pt := range []cldf.ContractType{LockReleaseTokenPool, SiloedLockReleaseTokenPool} {
		require.True(t, IsLockReleasePoolType(pt.String()), "%s should be a lock-release type", pt)
		require.False(t, IsBurnMintPoolType(pt.String()), "%s should not be a burn-mint type", pt)
	}

	for _, pt := range []cldf.ContractType{
		BurnMintTokenPool, BurnFromMintTokenPool, BurnWithFromMintTokenPool, BurnToAddressMintTokenPool,
	} {
		require.True(t, IsBurnMintPoolType(pt.String()), "%s should be a burn-mint type", pt)
		require.False(t, IsLockReleasePoolType(pt.String()), "%s should not be a lock-release type", pt)
	}

	// Intentionally excluded from both: see the predicate doc comments.
	require.False(t, IsLockReleasePoolType(HybridLockReleaseUSDCTokenPool.String()))
	require.False(t, IsLockReleasePoolType(BurnMintWithLockReleaseFlag.String()))
}

// TestPoolTypePredicates_SiloedLockReleaseVariants pins the custom legacy XERC20 group pool's
// classification. It implements the v1.6.1 siloed interface and must take the siloed migration
// source path, but must NOT be a deployable lock-release type: the v2.0.0 deploy dispatch has no
// contract for it and would deploy an orphaned lockbox and hooks before failing.
func TestPoolTypePredicates_SiloedLockReleaseVariants(t *testing.T) {
	t.Parallel()

	require.True(t, IsSiloedLockReleasePoolType(SiloedLockReleaseTokenPool.String()),
		"the standard siloed pool is a siloed migration source")
	require.True(t, IsSiloedLockReleasePoolType(SiloedWithUnsiloedXERC20GroupTokenPool.String()),
		"the custom XERC20 group pool implements the siloed interface and is a siloed migration source")

	// Not siloed migration sources: the plain pool is not siloed, and the v1.5.0 *AndProxy pool has
	// no siloed variant.
	require.False(t, IsSiloedLockReleasePoolType(LockReleaseTokenPool.String()))
	require.False(t, IsSiloedLockReleasePoolType(LockReleaseTokenPoolAndProxy.String()))

	// The custom type must stay out of the deploy dispatch predicate.
	require.False(t, IsLockReleasePoolType(SiloedWithUnsiloedXERC20GroupTokenPool.String()),
		"SiloedWithUnsiloedXERC20GroupTokenPool must NOT be a deployable lock-release type: the v2.0.0 "+
			"deploy dispatch would deploy a lockbox and hooks before failing on the missing bytecode")

	// Its unsiloed balance is shared XERC20LockBox collateral, not pool-owned liquidity.
	require.True(t, IsSharedCollateralUnsiloedPool(SiloedWithUnsiloedXERC20GroupTokenPool.String()),
		"the custom pool's unsiloed balance backs the XERC20 token and must be gated")
	require.False(t, IsSharedCollateralUnsiloedPool(SiloedLockReleaseTokenPool.String()),
		"the standard siloed pool holds its own unsiloed liquidity")
	require.False(t, IsSharedCollateralUnsiloedPool(LockReleaseTokenPool.String()))
}
