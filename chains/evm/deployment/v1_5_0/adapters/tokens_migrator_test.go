package adapters

import (
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	chain_selectors "github.com/smartcontractkit/chain-selectors"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	drip_bindings "github.com/smartcontractkit/chainlink-evm/gethwrappers/shared/generated/1_5_0/burn_mint_erc20_with_drip"

	rmnproxyops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/rmn_proxy"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/router"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/burn_mint_erc20_with_drip"
	bmtpap "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/burn_mint_token_pool_and_proxy"
	tpSeq "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/sequences/token_pool"
	tokensapi "github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils"
)

var (
	localSelector  = chain_selectors.ETHEREUM_MAINNET.Selector
	remoteSelector = chain_selectors.AVALANCHE_MAINNET.Selector
)

// TestTokenAdapter_MigratorReads exercises the TokenPoolMigrator surface against a real pool.
//
// GetRemotePools is the only one with logic of its own: v1.5.0 has no getRemotePools(), so the
// singular getRemotePool is adapted to the plural contract. A v1.5.0 pool holds at most one remote
// pool per lane, so the result is always empty or single-element - which is what makes retargeting
// a hard cutover rather than the append that v1.5.1+ can do.
func TestTokenAdapter_MigratorReads(t *testing.T) {
	t.Parallel()

	e, poolAddr := setupConfiguredPool(t)
	adapter := NewTokenAdapter()
	poolBytes := poolAddr.Bytes()

	supported, err := adapter.GetSupportedChains(*e, localSelector, poolBytes, testRemoteToken.Bytes())
	require.NoError(t, err)
	require.Equal(t, []uint64{remoteSelector}, supported)

	remoteToken, err := adapter.GetRemoteToken(*e, localSelector, poolBytes, testRemoteToken.Bytes(), remoteSelector)
	require.NoError(t, err)
	require.Equal(t, testRemoteToken.Bytes(), remoteToken)

	remotePools, err := adapter.GetRemotePools(*e, localSelector, poolBytes, testRemoteToken.Bytes(), remoteSelector)
	require.NoError(t, err)
	require.Len(t, remotePools, 1, "a v1.5.0 pool holds exactly one remote pool per lane")
	require.Equal(t, common.LeftPadBytes(testRemotePool.Bytes(), 32), remotePools[0])
}

// TestTokenAdapter_GetRemotePools_UnconfiguredLane asserts the shim reports "none" rather than a
// slice holding an empty entry, which would read as a configured pool at the zero address.
func TestTokenAdapter_GetRemotePools_UnconfiguredLane(t *testing.T) {
	t.Parallel()

	e, poolAddr := setupConfiguredPool(t)
	adapter := NewTokenAdapter()

	unconfigured := chain_selectors.POLYGON_MAINNET.Selector
	remotePools, err := adapter.GetRemotePools(*e, localSelector, poolAddr.Bytes(), testRemoteToken.Bytes(), unconfigured)
	require.NoError(t, err)
	require.Empty(t, remotePools)
}

func TestTokenAdapter_DeriveTokenDecimals_ReadsFromToken(t *testing.T) {
	t.Parallel()

	e, poolAddr := setupConfiguredPool(t)
	adapter := NewTokenAdapter()

	poolRef := datastore.AddressRef{ChainSelector: localSelector, Address: poolAddr.Hex()}
	tokenAddr, err := adapter.DeriveTokenAddress(*e, localSelector, poolRef)
	require.NoError(t, err)

	// v1.5.0 pools have no getTokenDecimals(); the adapter must fall through to the token's own
	// ERC20 decimals(). BurnMintERC20WithDrip is fixed at 18.
	decimals, err := adapter.DeriveTokenDecimals(*e, localSelector, poolRef, common.HexToAddress(tokenAddr).Bytes())
	require.NoError(t, err)
	require.Equal(t, uint8(18), decimals)
}

var (
	testRemoteToken = common.HexToAddress("0x00000000000000000000000000000000000000aa")
	testRemotePool  = common.HexToAddress("0x00000000000000000000000000000000000000bb")
)

// setupConfiguredPool deploys a drip token and a BurnMintTokenPoolAndProxy, then wires one remote
// chain onto it.
func setupConfiguredPool(t *testing.T) (*cldf.Environment, common.Address) {
	t.Helper()

	e, err := environment.New(t.Context(), environment.WithEVMSimulated(t, []uint64{localSelector}))
	require.NoError(t, err)

	chain := e.BlockChains.EVMChains()[localSelector]
	ds := datastore.NewMemoryDataStore()

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
		Address:       common.HexToAddress("0x1234567890abcdef1234567890abcdef12345678").Hex(),
		ChainSelector: localSelector,
	}))
	require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
		Type:          datastore.ContractType(rmnproxyops.ContractType),
		Version:       semver.MustParse("1.0.0"),
		Address:       common.HexToAddress("0xabcdef1234567890abcdef1234567890abcdef12").Hex(),
		ChainSelector: localSelector,
	}))
	e.DataStore = ds.Seal()

	deployReport, err := cldf_ops.ExecuteSequence(e.OperationsBundle, tpSeq.DeployTokenPool, e.BlockChains, tokensapi.DeployTokenPoolInput{
		TokenRef:          &tokenRef,
		PoolType:          string(bmtpap.ContractType),
		TokenPoolVersion:  utils.Version_1_5_0,
		ChainSelector:     localSelector,
		ExistingDataStore: e.DataStore,
	})
	require.NoError(t, err)
	require.Len(t, deployReport.Output.Addresses, 1)
	poolAddr := common.HexToAddress(deployReport.Output.Addresses[0].Address)

	rl := &tokensapi.RateLimiterConfigFloatInput{IsEnabled: true, Capacity: 100, Rate: 10}
	_, err = cldf_ops.ExecuteSequence(e.OperationsBundle, tpSeq.ConfigureTokenPoolForRemoteChains, chain,
		tpSeq.ConfigureTokenPoolForRemoteChainsInput{
			TokenPoolAddress: poolAddr,
			TokenPoolVersion: utils.Version_1_5_0,
			RemoteChains: map[uint64]tokensapi.RemoteChainConfig[[]byte, string]{
				remoteSelector: {
					RemoteToken:               testRemoteToken.Bytes(),
					RemotePool:                testRemotePool.Bytes(),
					RemoteDecimals:            18,
					OutboundRateLimiterConfig: rl,
					InboundRateLimiterConfig:  rl,
				},
			},
		})
	require.NoError(t, err)

	return e, poolAddr
}
