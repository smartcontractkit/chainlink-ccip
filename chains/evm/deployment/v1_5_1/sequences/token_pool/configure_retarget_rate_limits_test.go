package token_pool

import (
	"math/big"
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	chain_selectors "github.com/smartcontractkit/chain-selectors"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	bnm_bindings "github.com/smartcontractkit/chainlink-evm/gethwrappers/shared/generated/initial/burn_mint_erc20"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/burn_mint_erc20"
	rmnproxyops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/rmn_proxy"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/router"
	bnmPoolOps "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_1/operations/burn_mint_token_pool"
	seqV1_5_1 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_1/sequences"
	poolBindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_5_1/token_pool"
	tokenapi "github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils"
)

var (
	retargetLocalSelector  = chain_selectors.ETHEREUM_MAINNET.Selector
	retargetRemoteSelector = chain_selectors.AVALANCHE_MAINNET.Selector

	// Two remote tokens differing only in the decimals declared for them, which is what makes a
	// carried-forward inbound bucket wrong.
	remoteTokenSixDecimals      = common.HexToAddress("0x00000000000000000000000000000000000000a1")
	remoteTokenEighteenDecimals = common.HexToAddress("0x00000000000000000000000000000000000000a2")
)

// TestConfigureTokenPoolForRemoteChains_RejectsRemoteTokenRetargetWithEnabledLimits is the v1.5.1
// counterpart of the v1.5.0 test of the same name.
//
// v1.5.1 pools also denominate the inbound bucket in the REMOTE token's decimals
// (DoesPoolUseLocalDecimals reports false below 1.6.1), so reusing the stored bucket across a
// change of remote token rescales the limit by 10^(oldRemoteDecimals-newRemoteDecimals) - here
// 10^12. See tokens.ErrInboundRateLimitNotPortable.
func TestConfigureTokenPoolForRemoteChains_RejectsRemoteTokenRetargetWithEnabledLimits(t *testing.T) {
	t.Parallel()

	e, poolAddr := deployRetargetTestPool(t)
	pool := mustRetargetPool(t, e, poolAddr)

	require.NoError(t, configureRetargetLane(t, e, poolAddr, remoteTokenSixDecimals, 6,
		&tokenapi.RateLimiterConfigFloatInput{IsEnabled: true, Capacity: 100, Rate: 10}))

	before, err := pool.GetCurrentInboundRateLimiterState(&bind.CallOpts{}, retargetRemoteSelector)
	require.NoError(t, err)
	require.Equal(t, big.NewInt(110_000_000).String(), before.Capacity.String(),
		"inbound bucket should be denominated in the remote token's 6 decimals")

	err = configureRetargetLane(t, e, poolAddr, remoteTokenEighteenDecimals, 18, nil)
	require.ErrorContains(t, err, "cannot be reinterpreted for the new one")

	// The failed run must not have written anything.
	gotRemoteToken, err := pool.GetRemoteToken(&bind.CallOpts{}, retargetRemoteSelector)
	require.NoError(t, err)
	require.Equal(t, remoteTokenSixDecimals.Bytes(), gotRemoteToken)
	after, err := pool.GetCurrentInboundRateLimiterState(&bind.CallOpts{}, retargetRemoteSelector)
	require.NoError(t, err)
	require.Equal(t, before.Capacity.String(), after.Capacity.String())
}

// TestConfigureTokenPoolForRemoteChains_AllowsRemoteTokenRetargetWithDisabledLimits asserts the
// guard is scoped to enabled buckets: an all-zero bucket carries no decimal denomination.
func TestConfigureTokenPoolForRemoteChains_AllowsRemoteTokenRetargetWithDisabledLimits(t *testing.T) {
	t.Parallel()

	e, poolAddr := deployRetargetTestPool(t)
	pool := mustRetargetPool(t, e, poolAddr)

	require.NoError(t, configureRetargetLane(t, e, poolAddr, remoteTokenSixDecimals, 6,
		&tokenapi.RateLimiterConfigFloatInput{IsEnabled: false, Capacity: 0, Rate: 0}))
	require.NoError(t, configureRetargetLane(t, e, poolAddr, remoteTokenEighteenDecimals, 18, nil))

	gotRemoteToken, err := pool.GetRemoteToken(&bind.CallOpts{}, retargetRemoteSelector)
	require.NoError(t, err)
	require.Equal(t, remoteTokenEighteenDecimals.Bytes(), gotRemoteToken, "lane should have been retargeted")
}

// TestConfigureTokenPoolForRemoteChains_RetargetWithExplicitLimitsRescales is the escape hatch the
// error points operators at: supplying the limits lets the retarget proceed, scaled for the NEW
// remote token rather than carried over.
func TestConfigureTokenPoolForRemoteChains_RetargetWithExplicitLimitsRescales(t *testing.T) {
	t.Parallel()

	e, poolAddr := deployRetargetTestPool(t)
	pool := mustRetargetPool(t, e, poolAddr)

	rl := &tokenapi.RateLimiterConfigFloatInput{IsEnabled: true, Capacity: 100, Rate: 10}
	require.NoError(t, configureRetargetLane(t, e, poolAddr, remoteTokenSixDecimals, 6, rl))
	require.NoError(t, configureRetargetLane(t, e, poolAddr, remoteTokenEighteenDecimals, 18, rl))

	inbound, err := pool.GetCurrentInboundRateLimiterState(&bind.CallOpts{}, retargetRemoteSelector)
	require.NoError(t, err)
	require.True(t, inbound.IsEnabled)

	// 110 tokens at the NEW remote token's 18 decimals, not the 110e6 the old bucket held. The x1.1
	// inbound premium is computed in float64, so allow a tolerance well below one token.
	want := new(big.Int).Mul(big.NewInt(110), new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil))
	diff := new(big.Int).Abs(new(big.Int).Sub(want, inbound.Capacity))
	require.Negative(t, diff.Cmp(big.NewInt(1e12)),
		"inbound should be rescaled to the new remote token's decimals: want ~%s, got %s", want, inbound.Capacity)
}

func configureRetargetLane(
	t *testing.T,
	e *cldf.Environment,
	poolAddr, remoteToken common.Address,
	remoteDecimals uint8,
	rl *tokenapi.RateLimiterConfigFloatInput,
) error {
	t.Helper()

	chain := e.BlockChains.EVMChains()[retargetLocalSelector]
	cfg := tokenapi.RemoteChainConfig[[]byte, string]{
		RemoteToken:    remoteToken.Bytes(),
		RemotePool:     common.HexToAddress("0x00000000000000000000000000000000000000bb").Bytes(),
		RemoteDecimals: remoteDecimals,
	}
	if rl != nil {
		cfg.OutboundRateLimiterConfig = rl
		cfg.InboundRateLimiterConfig = rl
	}

	_, err := cldf_ops.ExecuteSequence(e.OperationsBundle, ConfigureTokenPoolForRemoteChains, chain,
		ConfigureTokenPoolForRemoteChainsInput{
			TokenPoolAddress: poolAddr,
			TokenPoolVersion: utils.Version_1_5_1,
			RemoteChains:     map[uint64]tokenapi.RemoteChainConfig[[]byte, string]{retargetRemoteSelector: cfg},
		})

	return err
}

// deployRetargetTestPool stands up a simulated chain with an 18-decimal token and a v1.5.1
// BurnMintTokenPool over it.
func deployRetargetTestPool(t *testing.T) (*cldf.Environment, common.Address) {
	t.Helper()

	e, err := environment.New(t.Context(), environment.WithEVMSimulated(t, []uint64{retargetLocalSelector}))
	require.NoError(t, err)

	chain := e.BlockChains.EVMChains()[retargetLocalSelector]
	ds := datastore.NewMemoryDataStore()

	tokenAddr := deployRetargetTestToken(t, chain)
	tokenRef := datastore.AddressRef{
		Type:          datastore.ContractType(burn_mint_erc20.ContractType),
		Version:       semver.MustParse("1.0.0"),
		Address:       tokenAddr.Hex(),
		ChainSelector: retargetLocalSelector,
		Qualifier:     "RETARGET",
	}
	require.NoError(t, ds.Addresses().Add(tokenRef))
	require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
		Type:          datastore.ContractType(router.ContractType),
		Version:       router.Version,
		Address:       common.HexToAddress("0x1234567890abcdef1234567890abcdef12345678").Hex(),
		ChainSelector: retargetLocalSelector,
	}))
	require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
		Type:          datastore.ContractType(rmnproxyops.ContractType),
		Version:       semver.MustParse("1.0.0"),
		Address:       common.HexToAddress("0xabcdef1234567890abcdef1234567890abcdef12").Hex(),
		ChainSelector: retargetLocalSelector,
	}))
	e.DataStore = ds.Seal()

	report, err := cldf_ops.ExecuteSequence(e.OperationsBundle, seqV1_5_1.DeployTokenPool, e.BlockChains, tokenapi.DeployTokenPoolInput{
		TokenRef:          &tokenRef,
		PoolType:          string(bnmPoolOps.ContractType),
		TokenPoolVersion:  utils.Version_1_5_1,
		ChainSelector:     retargetLocalSelector,
		ExistingDataStore: e.DataStore,
	})
	require.NoError(t, err)
	require.Len(t, report.Output.Addresses, 1)

	return e, common.HexToAddress(report.Output.Addresses[0].Address)
}

func deployRetargetTestToken(t *testing.T, chain evm.Chain) common.Address {
	t.Helper()

	tokenAddr, tx, _, err := bnm_bindings.DeployBurnMintERC20(
		chain.DeployerKey, chain.Client, "Retarget Token", "RETARGET", 18,
		new(big.Int).Mul(big.NewInt(1e9), big.NewInt(1e18)), big.NewInt(0),
	)
	require.NoError(t, err)
	_, err = cldf.ConfirmIfNoError(chain, tx, err)
	require.NoError(t, err)

	return tokenAddr
}

func mustRetargetPool(t *testing.T, e *cldf.Environment, poolAddr common.Address) *poolBindings.TokenPool {
	t.Helper()

	chain := e.BlockChains.EVMChains()[retargetLocalSelector]
	pool, err := poolBindings.NewTokenPool(poolAddr, chain.Client)
	require.NoError(t, err)

	return pool
}
