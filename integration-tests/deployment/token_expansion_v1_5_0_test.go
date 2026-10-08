package deployment

import (
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/type_and_version"
	bnmERC20DripOps "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/burn_mint_erc20_with_drip"
	lrtpOps "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/lock_release_token_pool"
	tpapOps "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/token_pool_and_proxy"
	bmtpapBindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_5_0/burn_mint_token_pool_and_proxy"
	lrtpapBindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_5_0/lock_release_token_pool_and_proxy"
	tarbindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_5_0/token_admin_registry"
	tokensapi "github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	cciputils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	bnmERC20DripBindings "github.com/smartcontractkit/chainlink-evm/gethwrappers/shared/generated/1_5_0/burn_mint_erc20_with_drip"
)

// TestTokenExpansion_V1_5_0_ProxyPool covers the v1.5.0 expansion step itself - deploy a
// BurnMintERC20WithDrip token, deploy a BurnMintTokenPoolAndProxy over it, wire the lane, and
// register in the TokenAdminRegistry - independently of the later v2.0.0 upgrade.
//
// The two assertions that are specific to v1.5.0, rather than shared with v1.5.1+:
//
//   - the pool is a BurnMintTokenPoolAndProxy, one contract that is its own proxy, so every read
//     targets a single address rather than an impl/proxy pair;
//   - the pool holds MINTER_ROLE and BURNER_ROLE on the token. That grant runs only when
//     utils.IsBurnMintPoolType reports true for the configured pool type, and
//     BurnMintTokenPoolAndProxy had to be added to that predicate. Without it the grant is
//     silently skipped: the pipeline still reports success and transfers fail only at execution
//     time, so this is the assertion that pins the fix.
func TestTokenExpansion_V1_5_0_ProxyPool(t *testing.T) {
	s := setupLegacyConnectedBnMPair(t, cciputils.Version_1_5_0)

	chainA := s.env.BlockChains.EVMChains()[s.selA]
	opts := &bind.CallOpts{Context: t.Context()}

	pool, err := bmtpapBindings.NewBurnMintTokenPoolAndProxy(s.oldPoolAddrA, chainA.Client)
	require.NoError(t, err)
	token, err := bnmERC20DripBindings.NewBurnMintERC20WithDrip(s.tokAddrA, chainA.Client)
	require.NoError(t, err)

	// The pool is the proxy-combined v1.5.0 type, deployed at the version the adapter claims.
	tv, err := pool.TypeAndVersion(opts)
	require.NoError(t, err)
	require.Equal(t, "BurnMintTokenPoolAndProxy 1.5.0", tv)

	// Pool and token are wired to each other.
	gotToken, err := pool.GetToken(opts)
	require.NoError(t, err)
	require.Equal(t, s.tokAddrA, gotToken, "pool should point at the deployed drip token")
	supportsToken, err := pool.IsSupportedToken(opts, s.tokAddrA)
	require.NoError(t, err)
	require.True(t, supportsToken)

	// The token is a real, initialized BurnMintERC20WithDrip. Its decimals are fixed at 18 by the
	// constructor, which is why legacyPairSpecFor cannot give the v1.5.0 case a decimal mismatch.
	symbol, err := token.Symbol(opts)
	require.NoError(t, err)
	require.Equal(t, "MIG_TOK_A", symbol)
	decimals, err := token.Decimals(opts)
	require.NoError(t, err)
	require.Equal(t, uint8(18), decimals)

	// Mint/burn roles were granted to the pool. See the note above: this is what would regress if
	// BurnMintTokenPoolAndProxy were dropped from utils.IsBurnMintPoolType.
	hasMinter, err := token.HasRole(opts, bnmERC20DripOps.MintRole, s.oldPoolAddrA)
	require.NoError(t, err)
	require.True(t, hasMinter, "pool must hold MINTER_ROLE on the token or transfers fail at execution time")
	hasBurner, err := token.HasRole(opts, bnmERC20DripOps.BurnRole, s.oldPoolAddrA)
	require.NoError(t, err)
	require.True(t, hasBurner, "pool must hold BURNER_ROLE on the token or transfers fail at execution time")

	// No external admin was requested, so the CCIP admin stays the deploying key and the deployer
	// keeps DEFAULT_ADMIN_ROLE - TidyTokenRoles only hands it to a timelock when the CLL timelock
	// is present in the datastore, which it is not in this environment.
	ccipAdmin, err := token.GetCCIPAdmin(opts)
	require.NoError(t, err)
	require.Equal(t, chainA.DeployerKey.From, ccipAdmin)
	deployerIsAdmin, err := token.HasRole(opts, bnmERC20DripOps.DefaultAdminRole, chainA.DeployerKey.From)
	require.NoError(t, err)
	require.True(t, deployerIsAdmin)

	// Lane wiring, read through the v1.5.0 ABI: getRemotePool is singular here, unlike v1.5.1+.
	supported, err := pool.GetSupportedChains(opts)
	require.NoError(t, err)
	require.Equal(t, []uint64{s.selB}, supported)
	remoteToken, err := pool.GetRemoteToken(opts, s.selB)
	require.NoError(t, err)
	require.Equal(t, common.LeftPadBytes(s.tokAddrB.Bytes(), 32), remoteToken,
		"remote token is stored 32-byte left-padded, same as the remote pool")
	remotePool, err := pool.GetRemotePool(opts, s.selB)
	require.NoError(t, err)
	require.Equal(t, common.LeftPadBytes(s.oldPoolAddrB.Bytes(), 32), remotePool,
		"remote pool is stored 32-byte left-padded to match what the protocol sends")

	// Rate limits from the expansion input (100 capacity / 10 rate, scaled by 18 decimals).
	outbound, err := pool.GetCurrentOutboundRateLimiterState(opts, s.selB)
	require.NoError(t, err)
	require.True(t, outbound.IsEnabled)
	require.Positive(t, outbound.Capacity.Sign())

	// The TokenAdminRegistry was registered and points at this pool.
	tar, err := tarbindings.NewTokenAdminRegistry(s.tarAddrA, chainA.Client)
	require.NoError(t, err)
	cfg, err := tar.GetTokenConfig(opts, s.tokAddrA)
	require.NoError(t, err)
	require.Equal(t, s.oldPoolAddrA, cfg.TokenPool)
	require.False(t, cfg.Administrator == (common.Address{}), "token should have an administrator set")
}

// TestTokenExpansion_V1_5_0_LockReleaseProxyPool is the lock-release counterpart of the test above:
// deploy a LockReleaseTokenPoolAndProxy over a v1.5.0 token, wire the lane, register in the TAR.
//
// What it adds over the burn-mint case:
//
//   - the immutable acceptLiquidity constructor arg survives the expansion path. It is the one
//     v1.5.0 deploy input with no v1.5.1 analogue, and it cannot be corrected after deployment;
//   - the pool is driven entirely through the shared TokenPoolAndProxy base surface. Lane wiring
//     is read back here through the LOCK-RELEASE ABI to confirm the two contracts really do agree
//     on that surface, which is the premise the single configure sequence rests on;
//   - the pool must NOT hold MINTER_ROLE/BURNER_ROLE. That is the mirror image of the burn-mint
//     assertion above and pins the predicate asymmetry in deployment/utils: a lock-release pool
//     must not be caught by utils.IsBurnMintPoolType, or every lock-release deployment would
//     start handing out mint authority it has no use for.
func TestTokenExpansion_V1_5_0_LockReleaseProxyPool(t *testing.T) {
	acceptLiquidity := true
	s := setupLegacyConnectedPair(t, cciputils.Version_1_5_0, legacyPairSpec{
		poolType:        cciputils.LockReleaseTokenPoolAndProxy,
		tokenType:       bnmERC20DripOps.ContractType,
		decimalsA:       18,
		decimalsB:       18,
		acceptLiquidity: &acceptLiquidity,
		// Same single-remote-pool-per-lane ABI as the burn-mint proxy pool.
		singlePool: true,
	})

	chainA := s.env.BlockChains.EVMChains()[s.selA]
	opts := &bind.CallOpts{Context: t.Context()}

	pool, err := lrtpapBindings.NewLockReleaseTokenPoolAndProxy(s.oldPoolAddrA, chainA.Client)
	require.NoError(t, err)
	token, err := bnmERC20DripBindings.NewBurnMintERC20WithDrip(s.tokAddrA, chainA.Client)
	require.NoError(t, err)

	tv, err := pool.TypeAndVersion(opts)
	require.NoError(t, err)
	require.Equal(t, "LockReleaseTokenPoolAndProxy 1.5.0", tv)

	// The immutable constructor flag made it through TokenExpansion intact.
	canAccept, err := pool.CanAcceptLiquidity(opts)
	require.NoError(t, err)
	require.True(t, canAccept, "acceptLiquidity must survive the expansion path; it cannot be fixed after deploy")

	// Nothing in the expansion path touches lock-release liquidity state.
	rebalancer, err := pool.GetRebalancer(opts)
	require.NoError(t, err)
	require.Equal(t, common.Address{}, rebalancer, "expansion must not set a rebalancer")

	// Pool and token are wired to each other.
	gotToken, err := pool.GetToken(opts)
	require.NoError(t, err)
	require.Equal(t, s.tokAddrA, gotToken)
	supportsToken, err := pool.IsSupportedToken(opts, s.tokAddrA)
	require.NoError(t, err)
	require.True(t, supportsToken)

	// A lock-release pool locks and releases; it must never have been granted mint/burn authority.
	hasMinter, err := token.HasRole(opts, bnmERC20DripOps.MintRole, s.oldPoolAddrA)
	require.NoError(t, err)
	require.False(t, hasMinter, "lock-release pool must not hold MINTER_ROLE")
	hasBurner, err := token.HasRole(opts, bnmERC20DripOps.BurnRole, s.oldPoolAddrA)
	require.NoError(t, err)
	require.False(t, hasBurner, "lock-release pool must not hold BURNER_ROLE")

	// Lane wiring, read back through the lock-release ABI.
	supported, err := pool.GetSupportedChains(opts)
	require.NoError(t, err)
	require.Equal(t, []uint64{s.selB}, supported)
	remoteToken, err := pool.GetRemoteToken(opts, s.selB)
	require.NoError(t, err)
	require.Equal(t, common.LeftPadBytes(s.tokAddrB.Bytes(), 32), remoteToken)
	remotePool, err := pool.GetRemotePool(opts, s.selB)
	require.NoError(t, err)
	require.Equal(t, common.LeftPadBytes(s.oldPoolAddrB.Bytes(), 32), remotePool,
		"v1.5.0 stores the single remote pool 32-byte left-padded")

	outbound, err := pool.GetCurrentOutboundRateLimiterState(opts, s.selB)
	require.NoError(t, err)
	require.True(t, outbound.IsEnabled)
	require.Positive(t, outbound.Capacity.Sign())

	// The TokenAdminRegistry was registered and points at this pool.
	tar, err := tarbindings.NewTokenAdminRegistry(s.tarAddrA, chainA.Client)
	require.NoError(t, err)
	cfg, err := tar.GetTokenConfig(opts, s.tokAddrA)
	require.NoError(t, err)
	require.Equal(t, s.oldPoolAddrA, cfg.TokenPool)
	require.False(t, cfg.Administrator == (common.Address{}), "token should have an administrator set")
}

// TestTokenExpansion_V1_5_0_PlainPools covers the expansion step for the plain (non-proxy) v1.5.0
// pools, BurnMintTokenPool and LockReleaseTokenPool. They are separate contracts from the *AndProxy
// pools above but share the v1.5.0 TokenPool base, so the same adapter, deploy sequence, and
// configure sequence must drive them. The v2.0.0 upgrade off the plain burn-mint pool is covered by
// the v1_5_0_plain cases in TestTokenExpansionMigration_AutoMigrate.
//
// The lane is read back through the shared TokenPoolAndProxy ops contract rather than each pool's
// own binding, because that is the binding the adapter uses against them - reading through it here
// is what proves the shared surface really covers the plain pools.
func TestTokenExpansion_V1_5_0_PlainPools(t *testing.T) {
	acceptLiquidity := true
	cases := []struct {
		name           string
		spec           legacyPairSpec
		typeAndVersion string
		isBurnMint     bool
	}{
		{
			name: "BurnMintTokenPool",
			spec: legacyPairSpec{
				poolType:   cciputils.BurnMintTokenPool,
				tokenType:  bnmERC20DripOps.ContractType,
				decimalsA:  18,
				decimalsB:  18,
				singlePool: true,
			},
			typeAndVersion: "BurnMintTokenPool 1.5.0",
			isBurnMint:     true,
		},
		{
			name: "LockReleaseTokenPool",
			spec: legacyPairSpec{
				poolType:        cciputils.LockReleaseTokenPool,
				tokenType:       bnmERC20DripOps.ContractType,
				decimalsA:       18,
				decimalsB:       18,
				acceptLiquidity: &acceptLiquidity,
				singlePool:      true,
			},
			typeAndVersion: "LockReleaseTokenPool 1.5.0",
			isBurnMint:     false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			s := setupLegacyConnectedPair(t, cciputils.Version_1_5_0, tc.spec)

			chainA := s.env.BlockChains.EVMChains()[s.selA]
			opts := &bind.CallOpts{Context: t.Context()}

			// The plain v1.5.0 contract was deployed, not the *AndProxy one.
			tvContract, err := type_and_version.NewTypeAndVersionContract(s.oldPoolAddrA, chainA.Client)
			require.NoError(t, err)
			tv, err := tvContract.TypeAndVersion(opts)
			require.NoError(t, err)
			require.Equal(t, tc.typeAndVersion, tv)

			pool, err := tpapOps.NewTokenPoolAndProxyContract(s.oldPoolAddrA, chainA.Client)
			require.NoError(t, err)
			token, err := bnmERC20DripBindings.NewBurnMintERC20WithDrip(s.tokAddrA, chainA.Client)
			require.NoError(t, err)

			// Pool and token are wired to each other.
			gotToken, err := pool.GetToken(opts)
			require.NoError(t, err)
			require.Equal(t, s.tokAddrA, gotToken)
			supportsToken, err := pool.IsSupportedToken(opts, s.tokAddrA)
			require.NoError(t, err)
			require.True(t, supportsToken)

			// Mint/burn authority follows the pool type: granted to the burn-mint pool (the plain
			// type was always in utils.IsBurnMintPoolType), never to the lock-release one.
			hasMinter, err := token.HasRole(opts, bnmERC20DripOps.MintRole, s.oldPoolAddrA)
			require.NoError(t, err)
			require.Equal(t, tc.isBurnMint, hasMinter, "MINTER_ROLE must be held iff the pool is burn-mint")
			hasBurner, err := token.HasRole(opts, bnmERC20DripOps.BurnRole, s.oldPoolAddrA)
			require.NoError(t, err)
			require.Equal(t, tc.isBurnMint, hasBurner, "BURNER_ROLE must be held iff the pool is burn-mint")

			if !tc.isBurnMint {
				lr, err := lrtpOps.NewLockReleaseTokenPoolContract(s.oldPoolAddrA, chainA.Client)
				require.NoError(t, err)
				canAccept, err := lr.CanAcceptLiquidity(opts)
				require.NoError(t, err)
				require.True(t, canAccept, "acceptLiquidity must survive the expansion path; it cannot be fixed after deploy")
				rebalancer, err := lr.GetRebalancer(opts)
				require.NoError(t, err)
				require.Equal(t, common.Address{}, rebalancer, "expansion must not set a rebalancer")
			}

			// Lane wiring: same single-remote-pool-per-lane ABI as the *AndProxy pools.
			supported, err := pool.GetSupportedChains(opts)
			require.NoError(t, err)
			require.Equal(t, []uint64{s.selB}, supported)
			remoteToken, err := pool.GetRemoteToken(opts, s.selB)
			require.NoError(t, err)
			require.Equal(t, common.LeftPadBytes(s.tokAddrB.Bytes(), 32), remoteToken)
			remotePool, err := pool.GetRemotePool(opts, s.selB)
			require.NoError(t, err)
			require.Equal(t, common.LeftPadBytes(s.oldPoolAddrB.Bytes(), 32), remotePool)

			outbound, err := pool.GetCurrentOutboundRateLimiterState(opts, s.selB)
			require.NoError(t, err)
			require.True(t, outbound.IsEnabled)
			require.Positive(t, outbound.Capacity.Sign())

			// The TokenAdminRegistry was registered and points at this pool.
			tar, err := tarbindings.NewTokenAdminRegistry(s.tarAddrA, chainA.Client)
			require.NoError(t, err)
			cfg, err := tar.GetTokenConfig(opts, s.tokAddrA)
			require.NoError(t, err)
			require.Equal(t, s.oldPoolAddrA, cfg.TokenPool)
			require.False(t, cfg.Administrator == (common.Address{}), "token should have an administrator set")
		})
	}
}

// TestRemoveRemotePools_V1_5_0_Deactivate deactivates a v1.5.0 pool whose peer is also v1.5.0.
// v1.5.0 stores a single remote pool per chain and has no removeRemotePool, so both the forward
// removal (on A) and the reverse removal (on peer B) clear the slot with setRemotePool(chain, empty
// bytes). The slot must read back empty (not a zero-address pool) while the chain stays supported,
// and A is unregistered from its TAR because it is the active pool.
func TestRemoveRemotePools_V1_5_0_Deactivate(t *testing.T) {
	s := setupLegacyConnectedBnMPair(t, cciputils.Version_1_5_0)
	opts := &bind.CallOpts{Context: t.Context()}
	chainA := s.env.BlockChains.EVMChains()[s.selA]
	poolA, err := bmtpapBindings.NewBurnMintTokenPoolAndProxy(s.oldPoolAddrA, chainA.Client)
	require.NoError(t, err)
	poolB, err := bmtpapBindings.NewBurnMintTokenPoolAndProxy(s.oldPoolAddrB, s.env.BlockChains.EVMChains()[s.selB].Client)
	require.NoError(t, err)

	// Pre-state: each pool's single remote-pool slot points at the other pool.
	remoteOfA, err := poolA.GetRemotePool(opts, s.selB)
	require.NoError(t, err)
	require.Equal(t, s.oldPoolAddrB, common.BytesToAddress(remoteOfA), "A should list B's pool")
	remoteOfB, err := poolB.GetRemotePool(opts, s.selA)
	require.NoError(t, err)
	require.Equal(t, s.oldPoolAddrA, common.BytesToAddress(remoteOfB), "B should list A's pool")

	input := tokensapi.RemoveRemotePoolsInput{
		Pools: []tokensapi.RemoveRemotePoolsPerPool{{
			ChainSelector: s.selA,
			Pool:          datastore.AddressRef{Address: s.oldPoolAddrA.Hex()},
			Deactivate:    true,
		}},
	}
	require.NoError(t, applyRemoveRemotePools(t, s.env, input))

	// Both slots are cleared (empty, not a zero-address pool); the chain configs are kept.
	remoteOfA, err = poolA.GetRemotePool(opts, s.selB)
	require.NoError(t, err)
	require.Empty(t, remoteOfA, "A's remote pool for B should be cleared")
	remoteOfB, err = poolB.GetRemotePool(opts, s.selA)
	require.NoError(t, err)
	require.Empty(t, remoteOfB, "B's remote pool for A should be cleared")
	supported, err := poolA.IsSupportedChain(opts, s.selB)
	require.NoError(t, err)
	require.True(t, supported, "A should still support chain B (only the remote pool is cleared)")

	// A was the active pool, so it is unregistered from A's TAR.
	tar, err := tarbindings.NewTokenAdminRegistry(s.tarAddrA, chainA.Client)
	require.NoError(t, err)
	cfg, err := tar.GetTokenConfig(opts, s.tokAddrA)
	require.NoError(t, err)
	require.Equal(t, common.Address{}, cfg.TokenPool, "A should be unregistered from A's TAR")

	// Re-running is a no-op: both slots are already empty and the TAR entry is already cleared.
	require.NoError(t, applyRemoveRemotePools(t, s.env, input))
}

// TestRemoveRemotePools_V1_5_0_DeactivateRawEncodedSlots covers v1.5.0 slots holding the raw 20-byte
// form of the remote pool instead of the usual 32-byte padded form: deactivate clears both slots.
func TestRemoveRemotePools_V1_5_0_DeactivateRawEncodedSlots(t *testing.T) {
	s := setupLegacyConnectedBnMPair(t, cciputils.Version_1_5_0)
	opts := &bind.CallOpts{Context: t.Context()}
	chainA := s.env.BlockChains.EVMChains()[s.selA]
	chainB := s.env.BlockChains.EVMChains()[s.selB]
	poolA, err := bmtpapBindings.NewBurnMintTokenPoolAndProxy(s.oldPoolAddrA, chainA.Client)
	require.NoError(t, err)
	poolB, err := bmtpapBindings.NewBurnMintTokenPoolAndProxy(s.oldPoolAddrB, chainB.Client)
	require.NoError(t, err)

	// Rewrite both slots in the raw 20-byte form.
	tx, err := poolA.SetRemotePool(chainA.DeployerKey, s.selB, s.oldPoolAddrB.Bytes())
	require.NoError(t, err)
	_, err = chainA.Confirm(tx)
	require.NoError(t, err)
	tx, err = poolB.SetRemotePool(chainB.DeployerKey, s.selA, s.oldPoolAddrA.Bytes())
	require.NoError(t, err)
	_, err = chainB.Confirm(tx)
	require.NoError(t, err)
	remoteOfA, err := poolA.GetRemotePool(opts, s.selB)
	require.NoError(t, err)
	require.Equal(t, s.oldPoolAddrB.Bytes(), remoteOfA, "A's slot should hold B's pool as raw 20 bytes")

	require.NoError(t, applyRemoveRemotePools(t, s.env, tokensapi.RemoveRemotePoolsInput{
		Pools: []tokensapi.RemoveRemotePoolsPerPool{{
			ChainSelector: s.selA,
			Pool:          datastore.AddressRef{Address: s.oldPoolAddrA.Hex()},
			Deactivate:    true,
		}},
	}))

	remoteOfA, err = poolA.GetRemotePool(opts, s.selB)
	require.NoError(t, err)
	require.Empty(t, remoteOfA, "A's remote pool for B should be cleared")
	remoteOfB, err := poolB.GetRemotePool(opts, s.selA)
	require.NoError(t, err)
	require.Empty(t, remoteOfB, "B's remote pool for A should be cleared")
}
