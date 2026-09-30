package adapters

import (
	"math/big"
	"strings"
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	drip_bindings "github.com/smartcontractkit/chainlink-evm/gethwrappers/shared/generated/1_5_0/burn_mint_erc20_with_drip"

	v1_4_0_token_pool "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_4_0/token_pool"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_4_0/burn_mint_token_pool"

	rmnproxyops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/rmn_proxy"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/router"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/burn_mint_erc20_with_drip"
	bmtpap "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/burn_mint_token_pool_and_proxy"
	tpap "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/token_pool_and_proxy"
	tpSeq "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/sequences/token_pool"
	tokensapi "github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils"
)

// TestEffectiveMigrationRateLimits_V14Previous mirrors the live wxUSD case from the spec: a
// v1.5.0 *AndProxy pool with its own rate limiter disabled, fronting a v1.4 previous pool with an
// enabled limiter. The effective limit must be the previous pool's, not disabled.
func TestEffectiveMigrationRateLimits_V14Previous(t *testing.T) {
	t.Parallel()

	e, err := environment.New(t.Context(), environment.WithEVMSimulated(t, []uint64{localSelector}))
	require.NoError(t, err)
	chain := e.BlockChains.EVMChains()[localSelector]
	ds := datastore.NewMemoryDataStore()

	routerAddr := common.HexToAddress("0x1234567890abcdef1234567890abcdef12345678")
	rmnProxyAddr := common.HexToAddress("0xabcdef1234567890abcdef1234567890abcdef12")

	tokenAddr, tx, _, err := drip_bindings.DeployBurnMintERC20WithDrip(chain.DeployerKey, chain.Client, "Test Drip", "TEST")
	require.NoError(t, err)
	_, err = cldf.ConfirmIfNoError(chain, tx, err)
	require.NoError(t, err)

	tokenRef := datastore.AddressRef{
		Type:          datastore.ContractType(burn_mint_erc20_with_drip.ContractType),
		Version:       burn_mint_erc20_with_drip.Version,
		Address:       tokenAddr.Hex(),
		ChainSelector: localSelector,
		Qualifier:     "TEST",
	}
	require.NoError(t, ds.Addresses().Add(tokenRef))
	require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
		Type:          datastore.ContractType(router.ContractType),
		Version:       router.Version,
		Address:       routerAddr.Hex(),
		ChainSelector: localSelector,
	}))
	require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
		Type:          datastore.ContractType(rmnproxyops.ContractType),
		Version:       semver.MustParse("1.0.0"),
		Address:       rmnProxyAddr.Hex(),
		ChainSelector: localSelector,
	}))
	e.DataStore = ds.Seal()

	// Deploy the v1.5.0 *AndProxy pool.
	acceptLiquidity := true
	deployReport, err := cldf_ops.ExecuteSequence(e.OperationsBundle, tpSeq.DeployTokenPool, e.BlockChains, tokensapi.DeployTokenPoolInput{
		TokenRef:          &tokenRef,
		PoolType:          string(bmtpap.ContractType),
		TokenPoolVersion:  utils.Version_1_5_0,
		ChainSelector:     localSelector,
		ExistingDataStore: e.DataStore,
		AcceptLiquidity:   &acceptLiquidity,
	})
	require.NoError(t, err)
	require.Len(t, deployReport.Output.Addresses, 1)
	proxyAddr := common.HexToAddress(deployReport.Output.Addresses[0].Address)

	// Deploy the v1.4 previous pool the proxy fronts.
	previousAddr, tx, _, err := burn_mint_token_pool.DeployBurnMintTokenPool(
		chain.DeployerKey, chain.Client, tokenAddr, nil, rmnProxyAddr, routerAddr,
	)
	require.NoError(t, err)
	_, err = cldf.ConfirmIfNoError(chain, tx, err)
	require.NoError(t, err)

	// Wire proxy.setPreviousPool(previous). No generated wrapper exists for this setter (only the
	// ABI constant), so build the call from the raw ABI, mirroring the pattern the generated
	// TokenPoolAndProxyContract wrapper itself uses.
	parsedABI, err := abi.JSON(strings.NewReader(tpap.TokenPoolAndProxyABI))
	require.NoError(t, err)
	bound := bind.NewBoundContract(proxyAddr, parsedABI, chain.Client, chain.Client, chain.Client)
	tx, err = bound.Transact(chain.DeployerKey, "setPreviousPool", previousAddr)
	require.NoError(t, err)
	_, err = cldf.ConfirmIfNoError(chain, tx, err)
	require.NoError(t, err)

	// Configure the previous pool's rate limits for the lane: enabled, mirroring the live wxUSD ->
	// previous-pool figures from the spec (cap 490000000000, rate 136111000).
	wantCapacity := big.NewInt(490000000000)
	wantRate := big.NewInt(136111000)
	previousTransactor, err := v1_4_0_token_pool.NewTokenPoolTransactor(previousAddr, chain.Client)
	require.NoError(t, err)
	tx, err = previousTransactor.ApplyChainUpdates(chain.DeployerKey, []v1_4_0_token_pool.TokenPoolChainUpdate{{
		RemoteChainSelector:       remoteSelector,
		Allowed:                   true,
		OutboundRateLimiterConfig: v1_4_0_token_pool.RateLimiterConfig{IsEnabled: true, Capacity: wantCapacity, Rate: wantRate},
		InboundRateLimiterConfig:  v1_4_0_token_pool.RateLimiterConfig{IsEnabled: true, Capacity: wantCapacity, Rate: wantRate},
	}})
	require.NoError(t, err)
	_, err = cldf.ConfirmIfNoError(chain, tx, err)
	require.NoError(t, err)

	// Configure the proxy pool's own limits for the lane as disabled, matching the "Affected" wxUSD
	// row in the spec (proxy disabled, previous enabled).
	disabledRL := &tokensapi.RateLimiterConfigFloatInput{IsEnabled: false}
	_, err = cldf_ops.ExecuteSequence(e.OperationsBundle, tpSeq.ConfigureTokenPoolForRemoteChains, chain,
		tpSeq.ConfigureTokenPoolForRemoteChainsInput{
			ChainSelector:    localSelector,
			TokenPoolAddress: proxyAddr,
			TokenPoolVersion: utils.Version_1_5_0,
			RemoteChains: map[uint64]tokensapi.RemoteChainConfig[[]byte, string]{
				remoteSelector: {
					RemoteToken:               testRemoteToken.Bytes(),
					RemotePool:                testRemotePool.Bytes(),
					RemoteDecimals:            18,
					OutboundRateLimiterConfig: disabledRL,
					InboundRateLimiterConfig:  disabledRL,
				},
			},
		})
	require.NoError(t, err)

	// Read what the migration reader would see on the proxy alone: both disabled.
	proxyLimits, err := (&poolOpsV150{}).GetCurrentRateLimits(e.OperationsBundle, chain, proxyAddr, remoteSelector, false)
	require.NoError(t, err)
	require.False(t, proxyLimits.Outbound.IsEnabled)
	require.False(t, proxyLimits.Inbound.IsEnabled)

	// The effective limit must fall through to the enabled previous-pool limits.
	effective, err := EffectiveMigrationRateLimits(
		e.OperationsBundle, chain, proxyAddr, remoteSelector,
		proxyLimits.Outbound, proxyLimits.Inbound,
		18, 18,
	)
	require.NoError(t, err)

	require.True(t, effective.Outbound.IsEnabled)
	require.Equal(t, 0, wantCapacity.Cmp(effective.Outbound.Capacity))
	require.Equal(t, 0, wantRate.Cmp(effective.Outbound.Rate))

	require.True(t, effective.Inbound.IsEnabled)
	require.Equal(t, 0, wantCapacity.Cmp(effective.Inbound.Capacity))
	require.Equal(t, 0, wantRate.Cmp(effective.Inbound.Rate))
}
