package deployment

import (
	"math/big"
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/stretchr/testify/require"

	evm_datastore_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/datastore"
	bnmERC20ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/burn_mint_erc20"
	bnmOpsV2_0_0 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/burn_mint_token_pool"
	evmtokensseq "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/sequences/tokens"
	tarbindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_5_0/token_admin_registry"
	tokenpoolV2_0_0 "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v2_0_0/token_pool"
	"github.com/smartcontractkit/chainlink-ccip/deployment/fees"
	"github.com/smartcontractkit/chainlink-ccip/deployment/finality"
	"github.com/smartcontractkit/chainlink-ccip/deployment/lanes"
	tokensapi "github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	cciputils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"

	evmseqV1_6_0 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/sequences"
	testsetupV2_0_0 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/testsetup"

	_ "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_1/adapters"
	_ "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_1/adapters"
	_ "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/adapters"
)

// Chain A is 18 decimals, chain B is 6 decimals. The decimal mismatch exercises the pre-1.6.1 inbound
// rate-limit decimal rebasing in the auto-migrate import path.
const (
	migDecimalsA = 18
	migDecimalsB = 6
)

// legacyBnMPair is the result of deploying and connecting a legacy BurnMint pool pair across two chains,
// used as the starting point for v2.0 upgrade tests.
type legacyBnMPair struct {
	env          *deployment.Environment
	selA, selB   uint64
	tokAddrA     common.Address
	tokAddrB     common.Address
	oldPoolAddrA common.Address
	oldPoolAddrB common.Address
	tarAddrA     common.Address
}

// autoMigrateUpgradeOpts configures optional YAML overrides for runAutoMigrateUpgrade.
type autoMigrateUpgradeOpts struct {
	feeOverride    *tokensapi.PartialTokenTransferFeeConfig
	explicitRemote bool // list chain B with explicit remoteToken/remotePool from the legacy pair
}

// TestTokenExpansionMigration_AutoMigrate exercises the AutoMigrateRemoteChains upgrade path: it deploys
// a legacy BurnMint pool pair, seeds legacy lane fees, then upgrades chain A's pool to v2.0 with an EMPTY
// RemoteChains map. The new pool must inherit chain B (token, remote pool, rate limits, and legacy fees
// carried forward from the active pool / FeeQuoter), and the TokenAdminRegistry must switch to the new pool.
//
// Both legacy versions are covered: v1.5.1 (< 1.6.1) stores inbound rate limits in remote/source decimals and
// is rebased on import; v1.6.1 (>= 1.6.1) stores them in local decimals and is imported as-is. With chain A at
// 18 decimals and chain B at 6 decimals, both paths must converge (within float ULP) on the same on-chain
// values for the new pool, which validates the inbound decimal rebasing end-to-end.
func TestTokenExpansionMigration_AutoMigrate(t *testing.T) {
	cases := []struct {
		name           string
		oldPoolVersion *semver.Version
	}{
		{"v1_5_1_to_v2_0_0", cciputils.Version_1_5_1},
		{"v1_6_1_to_v2_0_0", cciputils.Version_1_6_1},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			runAutoMigrateUpgrade(t, tc.oldPoolVersion, nil)
		})
	}
}

// TestTokenExpansionMigration_AutoMigrate_PartialYAMLMerge verifies that explicit YAML fee overrides merge
// with legacy lane fees imported by AutoMigrateRemoteChains: set fields win, unset fields are imported from
// the legacy FeeQuoter. Remote token/pool are backfilled from the active pool; chain B must still be listed
// as a counterpart when A lists B in RemoteChains. Both legacy pool versions are covered (v1.5.1 rebases
// inbound rate limits on import; v1.6.1 stores them in local decimals).
func TestTokenExpansionMigration_AutoMigrate_PartialYAMLMerge(t *testing.T) {
	cases := []struct {
		name           string
		oldPoolVersion *semver.Version
	}{
		{"v1_5_1_to_v2_0_0", cciputils.Version_1_5_1},
		{"v1_6_1_to_v2_0_0", cciputils.Version_1_6_1},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			runAutoMigrateUpgrade(t, tc.oldPoolVersion, &autoMigrateUpgradeOpts{
				feeOverride: &tokensapi.PartialTokenTransferFeeConfig{
					DefaultFinalityFeeUSDCents: cciputils.NewOptional(uint32(99)),
				},
			})
		})
	}
}

// TestTokenExpansionMigration_AutoMigrate_PreservesExplicitRemoteRefs verifies that when remoteToken and
// remotePool are set in YAML, auto-migrate uses them (backfill is skipped) and the upgrade still succeeds.
func TestTokenExpansionMigration_AutoMigrate_PreservesExplicitRemoteRefs(t *testing.T) {
	runAutoMigrateUpgrade(t, cciputils.Version_1_6_1, &autoMigrateUpgradeOpts{explicitRemote: true})
}

// TestTokenExpansionMigration_RequiresAllChainsWithoutAutoMigrate is the negative case: when
// AutoMigrateRemoteChains is false (the default), configuring a v2.0 pool over an existing active pool must
// still error if RemoteChains does not include every chain the active pool supports. This preserves the
// pre-existing upgrade-safety check. We invoke the ConfigureTokenPoolForRemoteChains sequence directly with
// an empty RemoteChains so the active pool's supported chain (B) is missing.
func TestTokenExpansionMigration_RequiresAllChainsWithoutAutoMigrate(t *testing.T) {
	s := setupLegacyConnectedBnMPair(t, cciputils.Version_1_5_1)

	// RegistryAddress + TokenAddress are set, so the active-pool supported-chains check runs. The active
	// pool (old pool A) supports chain B, which is absent from RemoteChains and the flag is off → error.
	chainA := s.env.BlockChains.EVMChains()[s.selA]
	_, err := cldf_ops.ExecuteSequence(
		s.env.OperationsBundle,
		evmtokensseq.ConfigureTokenPoolForRemoteChains,
		chainA,
		evmtokensseq.ConfigureTokenPoolForRemoteChainsInput{
			ChainSelector:    s.selA,
			TokenPoolAddress: s.oldPoolAddrA, // unused before the error returns
			RegistryAddress:  s.tarAddrA,
			TokenAddress:     s.tokAddrA,
			RemoteChains:     nil,
		},
	)
	require.Error(t, err)
	require.Contains(t, err.Error(), "remoteChains must include all active pool supported chains")
}

// setupLegacyConnectedBnMPair deploys v2.0 core contracts on two chains, then deploys a legacy
// (oldPoolVersion) BurnMint token+pool pair, connects them bidirectionally, and registers each in its
// TokenAdminRegistry. It returns the resolved addresses for use in upgrade tests.
func setupLegacyConnectedBnMPair(t *testing.T, oldPoolVersion *semver.Version) legacyBnMPair {
	t.Helper()

	const oldPoolQualA = "MIG_OLD_POOL_A"
	const oldPoolQualB = "MIG_OLD_POOL_B"
	const tokenSymbolA = "MIG_TOK_A"
	const tokenSymbolB = "MIG_TOK_B"

	selA := chainsel.TEST_90000001.Selector
	selB := chainsel.TEST_90000002.Selector
	e, err := environment.New(t.Context(), environment.WithEVMSimulated(t, []uint64{selA, selB}))
	require.NoError(t, err)

	// Deploy v2.0 core contracts (deployer-owned, so changesets run with the deployer key and no MCMS).
	cumulative := datastore.NewMemoryDataStore()
	DeployChainContractsV2_0_0(t, e, cumulative, selA)
	DeployChainContractsV2_0_0(t, e, cumulative, selB)
	e.DataStore = cumulative.Seal()

	// Wire CCIP lanes between the test chains. Required for legacy fee import during auto-migrate
	// (ResolveFeeAdapter / FeeQuoter reads on A → B).
	// DeployChainContracts caches proxy GetTarget reads from before SetTarget; refresh the bundle
	// so lane configuration observes the updated on-chain executor target.
	e.OperationsBundle = testsetupV2_0_0.BundleWithFreshReporter(e.OperationsBundle)
	connectOut, err := lanes.ConnectChains(lanes.GetLaneAdapterRegistry(), changesets.GetRegistry()).Apply(*e, lanes.ConnectChainsConfig{
		MCMS: mcms.Input{},
		Lanes: []lanes.LaneConfig{
			{
				Version: cciputils.Version_2_0_0,
				ChainA:  NewLaneChainDefinitionForV2(selA, selB),
				ChainB:  NewLaneChainDefinitionForV2(selB, selA),
			},
		},
	})
	require.NoError(t, err)
	MergeAddresses(t, e, connectOut.DataStore)

	// Deploy a legacy BurnMint pool pair (token + pool on each chain), connect them, and register in TAR.
	tokenPoolRL := tokensapi.RateLimiterConfigFloatInput{IsEnabled: true, Capacity: 100, Rate: 10}
	bnmPoolType := cciputils.BurnMintTokenPool
	oldOut, err := tokensapi.TokenExpansion().Apply(*e, tokensapi.TokenExpansionInput{
		ChainAdapterVersion: cciputils.Version_1_6_0,
		MCMS:                mcms.Input{},
		TokenExpansionInputPerChain: map[uint64]tokensapi.TokenExpansionInputPerChain{
			selA: {
				SkipOwnershipTransfer: true,
				TokenPoolVersion:      oldPoolVersion,
				DeployTokenInput: &tokensapi.DeployTokenInput{
					Name: "Migration Token A", Symbol: tokenSymbolA, Decimals: migDecimalsA,
					Type: bnmERC20ops.ContractType, Supply: nil,
				},
				DeployTokenPoolInput: &tokensapi.DeployTokenPoolInput{
					TokenPoolQualifier: oldPoolQualA,
					PoolType:           bnmPoolType.String(),
				},
				TokenTransferConfig: &tokensapi.TokenTransferConfig{
					RemoteChains: map[uint64]tokensapi.RemoteChainConfig[*datastore.AddressRef, datastore.AddressRef]{
						selB: {OutboundRateLimiterConfig: &tokenPoolRL},
					},
				},
			},
			selB: {
				SkipOwnershipTransfer: true,
				TokenPoolVersion:      oldPoolVersion,
				DeployTokenInput: &tokensapi.DeployTokenInput{
					Name: "Migration Token B", Symbol: tokenSymbolB, Decimals: migDecimalsB,
					Type: bnmERC20ops.ContractType, Supply: nil,
				},
				DeployTokenPoolInput: &tokensapi.DeployTokenPoolInput{
					TokenPoolQualifier: oldPoolQualB,
					PoolType:           bnmPoolType.String(),
				},
				TokenTransferConfig: &tokensapi.TokenTransferConfig{
					RemoteChains: map[uint64]tokensapi.RemoteChainConfig[*datastore.AddressRef, datastore.AddressRef]{
						selA: {OutboundRateLimiterConfig: &tokenPoolRL},
					},
				},
			},
		},
	})
	require.NoError(t, err)
	MergeAddresses(t, e, oldOut.DataStore)

	// Fetch the pool and token addresses from the datastore.
	evmAdapter := evmseqV1_6_0.EVMAdapter{}
	oldPoolAddrA, err := evmAdapter.FindLatestAddressRef(e.DataStore, datastore.AddressRef{ChainSelector: selA, Qualifier: oldPoolQualA, Type: datastore.ContractType(bnmPoolType)})
	require.NoError(t, err)
	oldPoolAddrB, err := evmAdapter.FindLatestAddressRef(e.DataStore, datastore.AddressRef{ChainSelector: selB, Qualifier: oldPoolQualB, Type: datastore.ContractType(bnmPoolType)})
	require.NoError(t, err)
	tokAddrA, err := evmAdapter.FindOneTokenAddress(e.DataStore, selA, &datastore.AddressRef{Qualifier: tokenSymbolA})
	require.NoError(t, err)
	tokAddrB, err := evmAdapter.FindOneTokenAddress(e.DataStore, selB, &datastore.AddressRef{Qualifier: tokenSymbolB})
	require.NoError(t, err)

	// Sanity-check: old pool A is connected to B and is the active pool in the TAR.
	// getSupportedChains is ABI-identical across pool versions, so the v2.0 binding reads any version.
	chainA := e.BlockChains.EVMChains()[selA]
	oldPoolA, err := tokenpoolV2_0_0.NewTokenPool(oldPoolAddrA, chainA.Client)
	require.NoError(t, err)
	oldSupported, err := oldPoolA.GetSupportedChains(&bind.CallOpts{Context: t.Context()})
	require.NoError(t, err)
	require.Contains(t, oldSupported, selB, "old pool A should support chain B before upgrade")

	// Ensure the TokenAdminRegistry points at the old pool before upgrade.
	tarAddrA, err := evmAdapter.GetTokenAdminRegistryAddress(e.DataStore, selA)
	require.NoError(t, err)
	tarA, err := tarbindings.NewTokenAdminRegistry(tarAddrA, chainA.Client)
	require.NoError(t, err)
	cfgBefore, err := tarA.GetTokenConfig(&bind.CallOpts{Context: t.Context()}, tokAddrA)
	require.NoError(t, err)
	require.Equal(t, oldPoolAddrA, cfgBefore.TokenPool, "active pool before upgrade should be the legacy pool")

	return legacyBnMPair{
		env:          e,
		selA:         selA,
		selB:         selB,
		tokAddrA:     tokAddrA,
		tokAddrB:     tokAddrB,
		oldPoolAddrA: oldPoolAddrA,
		oldPoolAddrB: oldPoolAddrB,
		tarAddrA:     tarAddrA,
	}
}

func runAutoMigrateUpgrade(t *testing.T, oldPoolVersion *semver.Version, opts *autoMigrateUpgradeOpts) {
	t.Helper()

	const newPoolQualA = "MIG_NEW_POOL_A"

	s := setupLegacyConnectedBnMPair(t, oldPoolVersion)
	e, selA, selB := s.env, s.selA, s.selB
	chainA := e.BlockChains.EVMChains()[selA]

	// ConnectChains + token expansion may have cached pre-lane router reads; refresh before fee I/O.
	e.OperationsBundle = testsetupV2_0_0.BundleWithFreshReporter(e.OperationsBundle)

	// Get fee adapter for chain A -> B.
	feeAdapter, fqRef, err := fees.ResolveFeeAdapter(e.OperationsBundle, e.BlockChains, e.DataStore, selA, selB)
	require.NoError(t, err)
	partialFee := fees.UnresolvedTokenTransferFeeArgs{
		MinFeeUSDCents:    cciputils.NewOptional(uint32(17)),
		DestGasOverhead:   cciputils.NewOptional(uint32(50_000)),
		DestBytesOverhead: cciputils.NewOptional(uint32(150_000)),
		IsEnabled:         cciputils.NewOptional(true),
	}

	// Seed legacy lane fees for chain A -> B (direct sequence, no MCMS — deployer-owned contracts).
	resolvedFee := partialFee.Resolve(feeAdapter.GetDefaultTokenTransferFeeConfig(selA, selB))
	_, err = cldf_ops.ExecuteSequence(
		e.OperationsBundle,
		feeAdapter.SetTokenTransferFee(e.DataStore, fqRef),
		e.BlockChains,
		fees.SetTokenTransferFeeSequenceInput{
			Selector: selA,
			Settings: map[uint64]map[string]*fees.TokenTransferFeeArgs{
				selB: {s.tokAddrA.Hex(): resolvedFee},
			},
		},
	)
	require.NoError(t, err)

	// FeeQuoter should have been seeded with the values above for chain A -> B.
	legacyFee, err := feeAdapter.GetOnchainTokenTransferFeeConfig(e.OperationsBundle, e.BlockChains, fqRef, selA, selB, s.tokAddrA.Hex())
	require.NoError(t, err)
	require.Equal(t, resolvedFee.MinFeeUSDCents, legacyFee.MinFeeUSDCents)
	require.Equal(t, resolvedFee.DestGasOverhead, legacyFee.DestGasOverhead)
	require.Equal(t, resolvedFee.DestBytesOverhead, legacyFee.DestBytesOverhead)
	require.Equal(t, resolvedFee.IsEnabled, legacyFee.IsEnabled)
	require.True(t, legacyFee.IsEnabled, "seeded legacy fee config should be enabled")
	perChain := map[uint64]tokensapi.TokenExpansionInputPerChain{
		selA: {
			SkipOwnershipTransfer: true,
			TokenPoolVersion:      cciputils.Version_2_0_0,
			DeployTokenPoolInput: &tokensapi.DeployTokenPoolInput{
				TokenPoolQualifier: newPoolQualA,
				PoolType:           bnmOpsV2_0_0.ContractType.String(),
				TokenRef:           &datastore.AddressRef{Address: s.tokAddrA.Hex()},
			},
			TokenTransferConfig: &tokensapi.TokenTransferConfig{
				AutoMigrateRemoteChains: true,
				RemoteChains:            nil,
			},
		},
	}

	// Allows us to test the code path for explicit fee configs and/or explicit remote chains
	if opts != nil && (opts.feeOverride != nil || opts.explicitRemote) {
		rc := tokensapi.RemoteChainConfig[*datastore.AddressRef, datastore.AddressRef]{}
		if opts.feeOverride != nil {
			rc.TokenTransferFeeConfig = opts.feeOverride
		}
		if opts.explicitRemote {
			rc.RemoteToken = &datastore.AddressRef{Address: s.tokAddrB.Hex()}
			rc.RemotePool = &datastore.AddressRef{Address: s.oldPoolAddrB.Hex()}
		}
		// ConfigureTokensForTransfers requires a counterpart entry for chain B when A lists B in RemoteChains.
		perChain[selA].TokenTransferConfig.RemoteChains = map[uint64]tokensapi.RemoteChainConfig[*datastore.AddressRef, datastore.AddressRef]{
			selB: rc,
		}
		perChain[selB] = tokensapi.TokenExpansionInputPerChain{
			SkipOwnershipTransfer: true,
			TokenTransferConfig: &tokensapi.TokenTransferConfig{
				TokenPoolRef: datastore.AddressRef{Address: s.oldPoolAddrB.Hex()},
				TokenRef:     datastore.AddressRef{Address: s.tokAddrB.Hex()},
				RemoteChains: map[uint64]tokensapi.RemoteChainConfig[*datastore.AddressRef, datastore.AddressRef]{
					selA: {},
				},
			},
		}
	}

	// Migrate chain A's pool to v2.0 using AutoMigrateRemoteChains (remotes, rate limits, and legacy fees
	// imported automatically). Refresh the bundle so fee discovery sees post-lane-connect router state.
	e.OperationsBundle = testsetupV2_0_0.BundleWithFreshReporter(e.OperationsBundle)
	upgradeOut, err := tokensapi.TokenExpansion().Apply(*e, tokensapi.TokenExpansionInput{
		ChainAdapterVersion:         cciputils.Version_2_0_0,
		MCMS:                        mcms.Input{},
		TokenExpansionInputPerChain: perChain,
	})
	require.NoError(t, err)
	MergeAddresses(t, e, upgradeOut.DataStore)

	// Ensure the new v2.0 pool was added to the datastore.
	newPoolRef := datastore.AddressRef{ChainSelector: selA, Type: datastore.ContractType(bnmOpsV2_0_0.ContractType), Version: bnmOpsV2_0_0.Version, Qualifier: newPoolQualA}
	newPoolAddrA, err := datastore_utils.FindAndFormatRef(e.DataStore, newPoolRef, selA, evm_datastore_utils.ToEVMAddress)
	require.NoError(t, err)
	require.NotEqual(t, s.oldPoolAddrA, newPoolAddrA, "new pool must be a distinct contract from the old pool")
	newPoolA, err := tokenpoolV2_0_0.NewTokenPool(newPoolAddrA, chainA.Client)
	require.NoError(t, err)

	var yamlPartial tokensapi.PartialTokenTransferFeeConfig
	if opts != nil && opts.feeOverride != nil {
		yamlPartial = *opts.feeOverride
	}
	legacyTpCfg := tokensapi.TokenTransferFeeConfig{
		DestGasOverhead:               legacyFee.DestGasOverhead,
		DestBytesOverhead:             legacyFee.DestBytesOverhead,
		DefaultFinalityFeeUSDCents:    legacyFee.MinFeeUSDCents,
		CustomFinalityFeeUSDCents:     0,
		DefaultFinalityTransferFeeBps: legacyFee.DeciBps,
		CustomFinalityTransferFeeBps:  0,
		IsEnabled:                     legacyFee.IsEnabled,
	}

	gotFee, err := newPoolA.GetTokenTransferFeeConfig(&bind.CallOpts{Context: t.Context()}, common.Address{}, selB, finality.RawWaitForFinality, []byte{})
	require.NoError(t, err)
	expectedFee := yamlPartial.MergeWith(legacyTpCfg)
	require.Equal(t, expectedFee.DefaultFinalityFeeUSDCents, gotFee.FinalityFeeUSDCents, "finality fee USD cents")
	require.Equal(t, expectedFee.DestGasOverhead, gotFee.DestGasOverhead, "dest gas overhead")
	require.Equal(t, expectedFee.DestBytesOverhead, gotFee.DestBytesOverhead, "dest bytes overhead")
	require.Equal(t, expectedFee.IsEnabled, gotFee.IsEnabled, "fee config enabled")
	require.Equal(t, expectedFee.DefaultFinalityTransferFeeBps, gotFee.FinalityTransferFeeBps, "finality transfer fee bps")
	require.Equal(t, expectedFee.CustomFinalityTransferFeeBps, gotFee.FastFinalityTransferFeeBps, "fast finality transfer fee bps")
	require.Equal(t, expectedFee.CustomFinalityFeeUSDCents, gotFee.FastFinalityFeeUSDCents, "fast finality fee USD cents")
	if opts != nil && opts.feeOverride != nil {
		legacyOnlyFee := tokensapi.PartialTokenTransferFeeConfig{}.MergeWith(legacyTpCfg)
		require.NotEqual(t, legacyOnlyFee, expectedFee, "YAML fee override should change at least one resolved field vs legacy-only import")
	}

	// The new pool supports chain B (discovered from the active pool, or backfilled when listed for a fee override).
	newSupported, err := newPoolA.GetSupportedChains(&bind.CallOpts{Context: t.Context()})
	require.NoError(t, err)
	require.Contains(t, newSupported, selB, "auto-migrated new pool A should support chain B")

	// Remote token carried forward from the active pool.
	gotRemoteToken, err := newPoolA.GetRemoteToken(&bind.CallOpts{Context: t.Context()}, selB)
	require.NoError(t, err)
	require.Equal(t, common.LeftPadBytes(s.tokAddrB.Bytes(), 32), gotRemoteToken, "remote token for chain B should be carried forward")

	// Remote pool carried forward — the old remote pool (pool B) must be among the registered remotes.
	gotRemotePools, err := newPoolA.GetRemotePools(&bind.CallOpts{Context: t.Context()}, selB)
	require.NoError(t, err)
	require.Contains(t, gotRemotePools, common.LeftPadBytes(s.oldPoolAddrB.Bytes(), 32), "remote pool B should be carried forward")

	// Rate limits carried forward and decimal-correct (all expressed in chain A's 18 decimals).
	// Outbound is scaled by local decimals with no premium, so it is exact for both versions. Inbound is
	// the counterpart's outbound + 10%; the x1.1 introduces a float64 rounding artifact in GenerateTPRLConfigs
	// whose magnitude differs between the rebased (v1.5.1) and native (v1.6.1) paths, so we assert inbound
	// within a tolerance far below 1 token yet far above any decimal-magnitude error (which would be off by
	// factors of 10^6 or more).
	e18 := new(big.Int).Exp(big.NewInt(10), big.NewInt(migDecimalsA), nil)
	const rlTolerance = int64(1e9) // ~1e-9 tokens at 18 decimals; absorbs float ULP, catches decimal bugs
	rl, err := newPoolA.GetCurrentRateLimiterState(&bind.CallOpts{Context: t.Context()}, selB, false)
	require.NoError(t, err)
	require.True(t, rl.OutboundRateLimiterState.IsEnabled, "outbound rate limit should be enabled")
	RequireBigIntsEqual(t, new(big.Int).Mul(big.NewInt(100), e18), rl.OutboundRateLimiterState.Capacity, "outbound capacity")
	RequireBigIntsEqual(t, new(big.Int).Mul(big.NewInt(10), e18), rl.OutboundRateLimiterState.Rate, "outbound rate")
	require.True(t, rl.InboundRateLimiterState.IsEnabled, "inbound rate limit should be enabled")
	RequireBigIntsApprox(t, new(big.Int).Mul(big.NewInt(110), e18), rl.InboundRateLimiterState.Capacity, rlTolerance, "inbound capacity")
	RequireBigIntsApprox(t, new(big.Int).Mul(big.NewInt(11), e18), rl.InboundRateLimiterState.Rate, rlTolerance, "inbound rate")

	// The TokenAdminRegistry now points at the new pool.
	tarA, err := tarbindings.NewTokenAdminRegistry(s.tarAddrA, chainA.Client)
	require.NoError(t, err)
	cfgAfter, err := tarA.GetTokenConfig(&bind.CallOpts{Context: t.Context()}, s.tokAddrA)
	require.NoError(t, err)
	require.Equal(t, newPoolAddrA, cfgAfter.TokenPool, "TAR should be switched to the new v2.0 pool")
}
