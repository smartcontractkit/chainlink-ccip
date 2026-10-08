package deployment

import (
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/stretchr/testify/require"

	bmp120bindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_2_0/burn_mint_token_pool"
	bmpap150bindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_5_0/burn_mint_token_pool_and_proxy"
	tarbindings150 "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_5_0/token_admin_registry"
	tokenpoolV2 "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v2_0_0/token_pool"

	bnmOpsV2 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/burn_mint_token_pool"
	testsetupV2 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/testsetup"
	tokensapi "github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	cciputils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"

	evm_datastore_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"

	_ "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/adapters"
)

// attachLegacyV1_2_0Pool reproduces the production legacy topology for the
// OLDEST surviving pool generation, on top of an already-deployed v1.5.0
// *TokenPoolAndProxy:
//
//	TokenAdminRegistry -> BurnMintTokenPoolAndProxy 1.5.0
//	                        └─ getPreviousPool() -> BurnMintTokenPool 1.2.0
//
// The wiring differs from v1.4.0 in the way that matters. A v1.4.0 pool is
// attached by repointing its ROUTER at the proxy; a v1.2.0 pool has no router
// at all, so the proxy is instead registered on it as both an onRamp and an
// offRamp via applyRampUpdates. That is what lets the proxy pass the legacy
// pool's onlyOnRamp/onlyOffRamp gates when it delegates, and it is also why the
// legacy pool's rate limits are keyed by the proxy address rather than by a
// remote chain selector.
func attachLegacyV1_2_0Pool(
	t *testing.T,
	e *deployment.Environment,
	selector uint64,
	token common.Address,
	proxyAddr common.Address,
) common.Address {
	t.Helper()

	chain := e.BlockChains.EVMChains()[selector]

	legacyAddr, tx, legacyPool, err := bmp120bindings.DeployBurnMintTokenPool(
		chain.DeployerKey, chain.Client,
		token, []common.Address{}, legacyArmProxy,
	)
	require.NoError(t, err, "failed to deploy v1.2.0 BurnMintTokenPool on chain %d", selector)
	_, err = chain.Confirm(tx)
	require.NoError(t, err, "failed to confirm v1.2.0 pool deployment on chain %d", selector)

	rampUpdate := []bmp120bindings.TokenPoolRampUpdate{{
		Ramp:    proxyAddr,
		Allowed: true,
		RateLimiterConfig: bmp120bindings.RateLimiterConfig{
			IsEnabled: false,
			Capacity:  common.Big0,
			Rate:      common.Big0,
		},
	}}
	tx, err = legacyPool.ApplyRampUpdates(chain.DeployerKey, rampUpdate, rampUpdate)
	require.NoError(t, err, "failed to register the proxy as a ramp on the v1.2.0 pool")
	_, err = chain.Confirm(tx)
	require.NoError(t, err, "failed to confirm applyRampUpdates on chain %d", selector)

	proxy, err := bmpap150bindings.NewBurnMintTokenPoolAndProxy(proxyAddr, chain.Client)
	require.NoError(t, err)
	tx, err = proxy.SetPreviousPool(chain.DeployerKey, legacyAddr)
	require.NoError(t, err, "failed to set previous pool on proxy %s", proxyAddr.Hex())
	_, err = chain.Confirm(tx)
	require.NoError(t, err, "failed to confirm setPreviousPool on chain %d", selector)

	gotPrev, err := proxy.GetPreviousPool(&bind.CallOpts{Context: t.Context()})
	require.NoError(t, err)
	require.Equal(t, legacyAddr, gotPrev, "proxy on chain %d should front the v1.2.0 pool", selector)

	isOn, err := legacyPool.IsOnRamp(&bind.CallOpts{Context: t.Context()}, proxyAddr)
	require.NoError(t, err)
	require.True(t, isOn, "proxy must be a registered onRamp on the v1.2.0 pool")
	isOff, err := legacyPool.IsOffRamp(&bind.CallOpts{Context: t.Context()}, proxyAddr)
	require.NoError(t, err)
	require.True(t, isOff, "proxy must be a registered offRamp on the v1.2.0 pool")

	gotToken, err := legacyPool.GetToken(&bind.CallOpts{Context: t.Context()})
	require.NoError(t, err)
	require.Equal(t, token, gotToken, "v1.2.0 pool on chain %d should hold the same token as its proxy", selector)

	return legacyAddr
}

func TestV1_2_0Adapter_ResolvesLaneStateThroughFrontingProxy(t *testing.T) {
	s := setupLegacyConnectedBnMPair(t, cciputils.Version_1_5_0)
	e, selA, selB := s.env, s.selA, s.selB

	legacyA := attachLegacyV1_2_0Pool(t, e, selA, s.tokAddrA, s.oldPoolAddrA)
	attachLegacyV1_2_0Pool(t, e, selB, s.tokAddrB, s.oldPoolAddrB)

	adapter, ok := tokensapi.GetTokenAdapterRegistry().GetTokenAdapter(chainsel.FamilyEVM, cciputils.Version_1_2_0)
	require.True(t, ok, "the v1.2.0 adapter should be registered")
	migrator, ok := adapter.(tokensapi.TokenPoolMigrator)
	require.True(t, ok, "the v1.2.0 adapter should implement TokenPoolMigrator")

	e.OperationsBundle = testsetupV2.BundleWithFreshReporter(e.OperationsBundle)

	supported, err := migrator.GetSupportedChains(*e, selA, legacyA.Bytes(), s.tokAddrA.Bytes())
	require.NoError(t, err)
	require.Contains(t, supported, selB, "supported chains must be resolved from the fronting proxy")

	remoteToken, err := migrator.GetRemoteToken(*e, selA, legacyA.Bytes(), s.tokAddrA.Bytes(), selB)
	require.NoError(t, err)
	require.Equal(t, common.LeftPadBytes(s.tokAddrB.Bytes(), 32), remoteToken)

	remotePools, err := migrator.GetRemotePools(*e, selA, legacyA.Bytes(), s.tokAddrA.Bytes(), selB)
	require.NoError(t, err)
	require.Contains(t, remotePools, common.LeftPadBytes(s.oldPoolAddrB.Bytes(), 32),
		"the remote pool is the remote PROXY, not the legacy pool behind it")

	supportedNoToken, err := migrator.GetSupportedChains(*e, selA, legacyA.Bytes(), nil)
	require.NoError(t, err, "adapter should resolve the proxy without being handed the token")
	require.Equal(t, supported, supportedNoToken)

	reader, ok := adapter.(tokensapi.RateLimitReaderAdapter)
	require.True(t, ok)
	limits, err := reader.GetOnchainRateLimits(
		e.OperationsBundle, e.BlockChains, e.DataStore, selA,
		datastore.AddressRef{Address: legacyA.Hex()},
		datastore.AddressRef{Address: s.tokAddrA.Hex()},
		selB, false,
	)
	require.NoError(t, err)
	require.True(t, limits.Outbound.IsEnabled, "outbound limit for the lane should be read from the proxy")
	require.Positive(t, limits.Outbound.Capacity.Sign())

	_, err = reader.GetOnchainRateLimits(
		e.OperationsBundle, e.BlockChains, e.DataStore, selA,
		datastore.AddressRef{Address: legacyA.Hex()},
		datastore.AddressRef{Address: s.tokAddrA.Hex()},
		selB, true,
	)
	require.ErrorContains(t, err, "fast finality buckets are not supported on v1.2.0 token pools")
}

func TestV1_2_0Adapter_RejectsPoolNotBehindAProxy(t *testing.T) {
	s := setupLegacyConnectedBnMPair(t, cciputils.Version_1_5_0)
	e, selA := s.env, s.selA
	chainA := e.BlockChains.EVMChains()[selA]

	orphan, tx, _, err := bmp120bindings.DeployBurnMintTokenPool(
		chainA.DeployerKey, chainA.Client,
		s.tokAddrA, []common.Address{}, legacyArmProxy,
	)
	require.NoError(t, err)
	_, err = chainA.Confirm(tx)
	require.NoError(t, err)

	adapter, ok := tokensapi.GetTokenAdapterRegistry().GetTokenAdapter(chainsel.FamilyEVM, cciputils.Version_1_2_0)
	require.True(t, ok)
	migrator := adapter.(tokensapi.TokenPoolMigrator)

	e.OperationsBundle = testsetupV2.BundleWithFreshReporter(e.OperationsBundle)
	_, err = migrator.GetSupportedChains(*e, selA, orphan.Bytes(), s.tokAddrA.Bytes())
	require.ErrorContains(t, err, "does not front legacy pool")
}

func TestTokenExpansionMigration_LegacyV1_2_0BehindProxy_ToV2_0_0(t *testing.T) {
	const newPoolQualA = "MIG120_NEW_POOL_A"

	s := setupLegacyConnectedBnMPair(t, cciputils.Version_1_5_0)
	e, selA, selB := s.env, s.selA, s.selB
	chainA := e.BlockChains.EVMChains()[selA]

	legacyA := attachLegacyV1_2_0Pool(t, e, selA, s.tokAddrA, s.oldPoolAddrA)
	legacyB := attachLegacyV1_2_0Pool(t, e, selB, s.tokAddrB, s.oldPoolAddrB)

	legacyPoolA, err := bmp120bindings.NewBurnMintTokenPool(legacyA, chainA.Client)
	require.NoError(t, err)

	onRamps, err := legacyPoolA.GetOnRamps(&bind.CallOpts{Context: t.Context()})
	require.NoError(t, err)
	require.Contains(t, onRamps, s.oldPoolAddrA,
		"the proxy must be a registered ramp; ramps are all the pool knows, it has no per-lane view")

	e.OperationsBundle = testsetupV2.BundleWithFreshReporter(e.OperationsBundle)
	upgradeOut, err := tokensapi.TokenExpansion().Apply(*e, tokensapi.TokenExpansionInput{
		ChainAdapterVersion: cciputils.Version_2_0_0,
		MCMS:                mcms.Input{},
		TokenExpansionInputPerChain: map[uint64]tokensapi.TokenExpansionInputPerChain{
			selA: {
				SkipOwnershipTransfer: true,
				TokenPoolVersion:      cciputils.Version_2_0_0,
				DeployTokenPoolInput: &tokensapi.DeployTokenPoolInput{
					TokenPoolQualifier: newPoolQualA,
					PoolType:           bnmOpsV2.ContractType.String(),
					TokenRef:           &datastore.AddressRef{Address: s.tokAddrA.Hex()},
				},
				TokenTransferConfig: &tokensapi.TokenTransferConfig{
					AutoMigrateRemoteChains: true,
				},
			},
		},
	})
	require.NoError(t, err, "migration from a proxy with an attached v1.2.0 previous pool should succeed")
	MergeAddresses(t, e, upgradeOut.DataStore)

	newPoolAddrA, err := datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{
		ChainSelector: selA,
		Type:          datastore.ContractType(bnmOpsV2.ContractType),
		Version:       bnmOpsV2.Version,
		Qualifier:     newPoolQualA,
	}, selA, evm_datastore_utils.ToEVMAddress)
	require.NoError(t, err)
	require.NotEqual(t, s.oldPoolAddrA, newPoolAddrA)
	require.NotEqual(t, legacyA, newPoolAddrA)

	newPoolA, err := tokenpoolV2.NewTokenPool(newPoolAddrA, chainA.Client)
	require.NoError(t, err)

	newSupported, err := newPoolA.GetSupportedChains(&bind.CallOpts{Context: t.Context()})
	require.NoError(t, err)
	require.Contains(t, newSupported, selB, "new pool A should support chain B after migration")

	gotRemoteToken, err := newPoolA.GetRemoteToken(&bind.CallOpts{Context: t.Context()}, selB)
	require.NoError(t, err)
	require.Equal(t, common.LeftPadBytes(s.tokAddrB.Bytes(), 32), gotRemoteToken,
		"remote token must come from the proxy; the v1.2.0 pool stores none")

	gotRemotePools, err := newPoolA.GetRemotePools(&bind.CallOpts{Context: t.Context()}, selB)
	require.NoError(t, err)
	require.Contains(t, gotRemotePools, common.LeftPadBytes(s.oldPoolAddrB.Bytes(), 32),
		"remote pool B should be the proxy")
	require.NotContains(t, gotRemotePools, common.LeftPadBytes(legacyB.Bytes(), 32),
		"the v1.2.0 pool behind proxy B must not be registered as a remote pool")

	tarA, err := tarbindings150.NewTokenAdminRegistry(s.tarAddrA, chainA.Client)
	require.NoError(t, err)
	cfgAfter, err := tarA.GetTokenConfig(&bind.CallOpts{Context: t.Context()}, s.tokAddrA)
	require.NoError(t, err)
	require.Equal(t, newPoolAddrA, cfgAfter.TokenPool, "registry should point at the new v2.0 pool after migration")

	stillOnRamps, err := legacyPoolA.GetOnRamps(&bind.CallOpts{Context: t.Context()})
	require.NoError(t, err)
	require.Equal(t, onRamps, stillOnRamps, "migration must not change the legacy pool's ramps")

	proxyA, err := bmpap150bindings.NewBurnMintTokenPoolAndProxy(s.oldPoolAddrA, chainA.Client)
	require.NoError(t, err)
	stillPrev, err := proxyA.GetPreviousPool(&bind.CallOpts{Context: t.Context()})
	require.NoError(t, err)
	require.Equal(t, legacyA, stillPrev, "migration must not clear the proxy's previous pool")
}
