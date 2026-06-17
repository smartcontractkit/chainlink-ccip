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
	tokensapi "github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	cciputils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"

	evmseqV1_6_0 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/sequences"

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

// setupLegacyConnectedBnMPair deploys v2.0 core contracts on two chains, then deploys a legacy
// (oldPoolVersion) BurnMint token+pool pair, connects them bidirectionally, and registers each in its
// TokenAdminRegistry. It returns the resolved addresses for use in upgrade tests.
func setupLegacyConnectedBnMPair(t *testing.T, oldPoolVersion *semver.Version) legacyBnMPair {
	t.Helper()

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

	defaultRL := tokensapi.RateLimiterConfigFloatInput{IsEnabled: true, Capacity: 100, Rate: 10}
	maxSupply := uint64(1e6)
	oldPoolQualA := "MIG_OLD_POOL_A"
	oldPoolQualB := "MIG_OLD_POOL_B"
	bmPoolType := cciputils.BurnMintTokenPool

	// Deploy a legacy BurnMint pool pair (token + pool on each chain), connect them, and register in TAR.
	oldOut, err := tokensapi.TokenExpansion().Apply(*e, tokensapi.TokenExpansionInput{
		ChainAdapterVersion: cciputils.Version_1_6_0,
		MCMS:                mcms.Input{},
		TokenExpansionInputPerChain: map[uint64]tokensapi.TokenExpansionInputPerChain{
			selA: {
				SkipOwnershipTransfer: true,
				TokenPoolVersion:      oldPoolVersion,
				DeployTokenInput: &tokensapi.DeployTokenInput{
					Name: "Migration Token A", Symbol: tokenSymbolA, Decimals: migDecimalsA,
					Type: bnmERC20ops.ContractType, Supply: &maxSupply,
				},
				DeployTokenPoolInput: &tokensapi.DeployTokenPoolInput{
					TokenPoolQualifier: oldPoolQualA,
					PoolType:           bmPoolType.String(),
				},
				TokenTransferConfig: &tokensapi.TokenTransferConfig{
					RemoteChains: map[uint64]tokensapi.RemoteChainConfig[*datastore.AddressRef, datastore.AddressRef]{
						selB: {OutboundRateLimiterConfig: &defaultRL},
					},
				},
			},
			selB: {
				SkipOwnershipTransfer: true,
				TokenPoolVersion:      oldPoolVersion,
				DeployTokenInput: &tokensapi.DeployTokenInput{
					Name: "Migration Token B", Symbol: tokenSymbolB, Decimals: migDecimalsB,
					Type: bnmERC20ops.ContractType, Supply: &maxSupply,
				},
				DeployTokenPoolInput: &tokensapi.DeployTokenPoolInput{
					TokenPoolQualifier: oldPoolQualB,
					PoolType:           bmPoolType.String(),
				},
				TokenTransferConfig: &tokensapi.TokenTransferConfig{
					RemoteChains: map[uint64]tokensapi.RemoteChainConfig[*datastore.AddressRef, datastore.AddressRef]{
						selA: {OutboundRateLimiterConfig: &defaultRL},
					},
				},
			},
		},
	})
	require.NoError(t, err)
	MergeAddresses(t, e, oldOut.DataStore)

	evmAdapter := evmseqV1_6_0.EVMAdapter{}
	chainA := e.BlockChains.EVMChains()[selA]

	oldPoolAddrA, err := evmAdapter.FindLatestAddressRef(e.DataStore, datastore.AddressRef{ChainSelector: selA, Qualifier: oldPoolQualA, Type: datastore.ContractType(bmPoolType)})
	require.NoError(t, err)
	oldPoolAddrB, err := evmAdapter.FindLatestAddressRef(e.DataStore, datastore.AddressRef{ChainSelector: selB, Qualifier: oldPoolQualB, Type: datastore.ContractType(bmPoolType)})
	require.NoError(t, err)
	tokAddrA, err := evmAdapter.FindOneTokenAddress(e.DataStore, selA, &datastore.AddressRef{Qualifier: tokenSymbolA})
	require.NoError(t, err)
	tokAddrB, err := evmAdapter.FindOneTokenAddress(e.DataStore, selB, &datastore.AddressRef{Qualifier: tokenSymbolB})
	require.NoError(t, err)

	// Sanity-check: old pool A is connected to B and is the active pool in the TAR.
	// getSupportedChains is ABI-identical across pool versions, so the v2.0 binding reads any version.
	oldPoolA, err := tokenpoolV2_0_0.NewTokenPool(oldPoolAddrA, chainA.Client)
	require.NoError(t, err)
	oldSupported, err := oldPoolA.GetSupportedChains(&bind.CallOpts{Context: t.Context()})
	require.NoError(t, err)
	require.Contains(t, oldSupported, selB, "old pool A should support chain B before upgrade")

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

// TestTokenExpansionMigration_AutoMigrateRemoteChains exercises the AutoMigrateRemoteChains upgrade path:
// it deploys a legacy BurnMint pool pair, connects them, then upgrades chain A's pool to v2.0 using
// AutoMigrateRemoteChains with an EMPTY RemoteChains map. The new pool must inherit chain B as a supported
// remote (token + remote pool + rate limits carried forward from the active pool), and the TokenAdminRegistry
// must be switched to point at the new pool — all without the operator listing any remote chains.
//
// Both legacy versions are covered: v1.5.1 (< 1.6.1) stores inbound rate limits in remote/source decimals and
// is rebased on import; v1.6.1 (>= 1.6.1) stores them in local decimals and is imported as-is. With chain A at
// 18 decimals and chain B at 6 decimals, both paths must converge (within float ULP) on the same on-chain
// values for the new pool, which validates the inbound decimal rebasing end-to-end.
func TestTokenExpansionMigration_AutoMigrateRemoteChains(t *testing.T) {
	cases := []struct {
		name           string
		oldPoolVersion *semver.Version
	}{
		{"v1_5_1_to_v2_0_0", cciputils.Version_1_5_1},
		{"v1_6_1_to_v2_0_0", cciputils.Version_1_6_1},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			runAutoMigrateRemoteChainsUpgrade(t, tc.oldPoolVersion)
		})
	}
}

func runAutoMigrateRemoteChainsUpgrade(t *testing.T, oldPoolVersion *semver.Version) {
	s := setupLegacyConnectedBnMPair(t, oldPoolVersion)
	e, selA, selB := s.env, s.selA, s.selB
	chainA := e.BlockChains.EVMChains()[selA]

	// Upgrade chain A's pool to v2.0 using AutoMigrateRemoteChains with an EMPTY RemoteChains map.
	// Chain B must be discovered from the active pool.
	newPoolQualA := "MIG_NEW_POOL_A"
	upgradeOut, err := tokensapi.TokenExpansion().Apply(*e, tokensapi.TokenExpansionInput{
		ChainAdapterVersion: cciputils.Version_2_0_0,
		MCMS:                mcms.Input{},
		TokenExpansionInputPerChain: map[uint64]tokensapi.TokenExpansionInputPerChain{
			selA: {
				SkipOwnershipTransfer: true,
				TokenPoolVersion:      bnmOpsV2_0_0.Version,
				DeployTokenPoolInput: &tokensapi.DeployTokenPoolInput{
					TokenPoolQualifier: newPoolQualA,
					PoolType:           bnmOpsV2_0_0.ContractType.String(),
					TokenRef:           &datastore.AddressRef{Address: s.tokAddrA.Hex()},
				},
				TokenTransferConfig: &tokensapi.TokenTransferConfig{
					// No remote chains listed — they must be discovered from the active pool.
					AutoMigrateRemoteChains: true,
				},
			},
		},
	})
	require.NoError(t, err)
	MergeAddresses(t, e, upgradeOut.DataStore)

	// Assert the new v2.0 pool inherited chain B and the TAR switched.
	newPoolRef := datastore.AddressRef{ChainSelector: selA, Type: datastore.ContractType(bnmOpsV2_0_0.ContractType), Version: bnmOpsV2_0_0.Version, Qualifier: newPoolQualA}
	newPoolAddrA, err := datastore_utils.FindAndFormatRef(e.DataStore, newPoolRef, selA, evm_datastore_utils.ToEVMAddress)
	require.NoError(t, err)
	require.NotEqual(t, s.oldPoolAddrA, newPoolAddrA, "new pool must be a distinct contract from the old pool")
	newPoolA, err := tokenpoolV2_0_0.NewTokenPool(newPoolAddrA, chainA.Client)
	require.NoError(t, err)

	// The new pool supports chain B even though the operator never listed it.
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
			TokenPoolAddress:   s.oldPoolAddrA, // unused before the error returns
			RegistryAddress:    s.tarAddrA,
			TokenAddress:       s.tokAddrA,
			RemoteChains:       nil,
		},
	)
	require.Error(t, err)
	require.Contains(t, err.Error(), "remoteChains must include all active pool supported chains")
}
