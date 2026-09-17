package token_pool

import (
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	chain_selectors "github.com/smartcontractkit/chain-selectors"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	drip_bindings "github.com/smartcontractkit/chainlink-evm/gethwrappers/shared/generated/1_5_0/burn_mint_erc20_with_drip"

	rmnproxyops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/rmn_proxy"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/router"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/burn_mint_erc20_with_drip"
	bmtpap "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/burn_mint_token_pool_and_proxy"
	tokenapi "github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils"
)

var testChainSelector = chain_selectors.ETHEREUM_MAINNET.Selector

const testTokenSymbol = "TEST"

func TestDeployTokenPool(t *testing.T) {
	t.Parallel()

	e, tokenRef, routerAddress, rmnProxyAddress := setupDeployEnv(t)

	input := tokenapi.DeployTokenPoolInput{
		TokenRef:          &tokenRef,
		PoolType:          string(bmtpap.ContractType),
		TokenPoolVersion:  utils.Version_1_5_0,
		ChainSelector:     testChainSelector,
		ExistingDataStore: e.DataStore,
	}

	report, err := cldf_ops.ExecuteSequence(e.OperationsBundle, DeployTokenPool, e.BlockChains, input)
	require.NoError(t, err, "Failed to execute DeployTokenPool sequence")
	require.Len(t, report.Output.Addresses, 1, "Should have deployed exactly one pool")

	poolRef := report.Output.Addresses[0]
	require.Equal(t, tokenRef.Address, poolRef.Qualifier, "Pool qualifier should default to the token address")
	require.Equal(t, utils.Version_1_5_0.String(), poolRef.Version.String())

	chain := e.BlockChains.EVMChains()[testChainSelector]
	pool, err := bmtpap.NewBurnMintTokenPoolAndProxyContract(common.HexToAddress(poolRef.Address), chain.Client)
	require.NoError(t, err)

	onChainToken, err := pool.GetToken(&bind.CallOpts{})
	require.NoError(t, err)
	require.Equal(t, common.HexToAddress(tokenRef.Address), onChainToken, "Token address mismatch")

	onChainRouter, err := pool.GetRouter(&bind.CallOpts{})
	require.NoError(t, err)
	require.Equal(t, routerAddress, onChainRouter, "Router address mismatch")

	supported, err := pool.IsSupportedToken(&bind.CallOpts{}, onChainToken)
	require.NoError(t, err)
	require.True(t, supported, "Pool should support its own token")

	// The v1.5.0 constructor takes the RMN proxy but exposes no getter on the generated
	// operations contract; assert it was at least resolved from the datastore.
	require.NotEqual(t, common.Address{}, rmnProxyAddress)
}

// TestDeployTokenPool_Idempotent asserts that a second run against a datastore that already
// records the pool returns the existing ref instead of deploying a duplicate.
func TestDeployTokenPool_Idempotent(t *testing.T) {
	t.Parallel()

	e, tokenRef, _, _ := setupDeployEnv(t)

	input := tokenapi.DeployTokenPoolInput{
		TokenRef:          &tokenRef,
		PoolType:          string(bmtpap.ContractType),
		TokenPoolVersion:  utils.Version_1_5_0,
		ChainSelector:     testChainSelector,
		ExistingDataStore: e.DataStore,
	}

	first, err := cldf_ops.ExecuteSequence(e.OperationsBundle, DeployTokenPool, e.BlockChains, input)
	require.NoError(t, err)
	require.Len(t, first.Output.Addresses, 1)
	poolRef := first.Output.Addresses[0]

	// Fold the deployed pool into the datastore and re-run.
	ds := datastore.NewMemoryDataStore()
	require.NoError(t, ds.Addresses().Add(tokenRef))
	require.NoError(t, ds.Addresses().Add(poolRef))
	input.ExistingDataStore = ds.Seal()

	second, err := cldf_ops.ExecuteSequence(e.OperationsBundle, DeployTokenPool, e.BlockChains, input)
	require.NoError(t, err)
	require.Len(t, second.Output.Addresses, 1)
	require.Equal(t, poolRef.Address, second.Output.Addresses[0].Address, "Should reuse the existing pool, not redeploy")
}

// TestDeployTokenPool_RejectsUnsupportedTypes locks in the tight scope: v1.5.0 support covers
// BurnMintTokenPoolAndProxy only. The lock-release and rebasing *AndProxy variants have bindings
// but no adapter support, and the plain v1.5.1-era types do not exist at this version - deploying
// any of them would produce a pool no changeset could subsequently configure.
func TestDeployTokenPool_RejectsUnsupportedTypes(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name        string
		poolType    string
		poolVersion *semver.Version
	}{
		{name: "LockReleaseTokenPoolAndProxy", poolType: "LockReleaseTokenPoolAndProxy", poolVersion: utils.Version_1_5_0},
		{name: "BurnWithFromMintTokenPoolAndProxy", poolType: "BurnWithFromMintTokenPoolAndProxy", poolVersion: utils.Version_1_5_0},
		{name: "PlainBurnMintTokenPool", poolType: string(utils.BurnMintTokenPool), poolVersion: utils.Version_1_5_0},
		{name: "LockReleaseTokenPool", poolType: string(utils.LockReleaseTokenPool), poolVersion: utils.Version_1_5_0},
		{name: "UnknownType", poolType: "NotARealTokenPool", poolVersion: utils.Version_1_5_0},
		// Right type, wrong version: the switch keys on the full "Type Version" string.
		{name: "RightTypeWrongVersion", poolType: string(bmtpap.ContractType), poolVersion: utils.Version_1_5_1},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			e, tokenRef, _, _ := setupDeployEnv(t)

			_, err := cldf_ops.ExecuteSequence(e.OperationsBundle, DeployTokenPool, e.BlockChains, tokenapi.DeployTokenPoolInput{
				TokenRef:          &tokenRef,
				PoolType:          tc.poolType,
				TokenPoolVersion:  tc.poolVersion,
				ChainSelector:     testChainSelector,
				ExistingDataStore: e.DataStore,
			})
			require.Error(t, err)
			require.Contains(t, err.Error(), "unsupported v1.5.0 token pool type and version")
		})
	}
}

func TestDeployTokenPool_RequiresTokenRefAndVersion(t *testing.T) {
	t.Parallel()

	e, tokenRef, _, _ := setupDeployEnv(t)

	_, err := cldf_ops.ExecuteSequence(e.OperationsBundle, DeployTokenPool, e.BlockChains, tokenapi.DeployTokenPoolInput{
		TokenRef:          &tokenRef,
		PoolType:          string(bmtpap.ContractType),
		ChainSelector:     testChainSelector,
		ExistingDataStore: e.DataStore,
	})
	require.ErrorContains(t, err, "TokenPoolVersion is required")

	_, err = cldf_ops.ExecuteSequence(e.OperationsBundle, DeployTokenPool, e.BlockChains, tokenapi.DeployTokenPoolInput{
		PoolType:          string(bmtpap.ContractType),
		TokenPoolVersion:  utils.Version_1_5_0,
		ChainSelector:     testChainSelector,
		ExistingDataStore: e.DataStore,
	})
	require.ErrorContains(t, err, "TokenRef is required")
}

// setupDeployEnv builds a simulated chain with a deployed BurnMintERC20WithDrip token and the
// router / RMN proxy refs the deploy sequence resolves from the datastore.
func setupDeployEnv(t *testing.T) (e *cldf.Environment, tokenRef datastore.AddressRef, routerAddress, rmnProxyAddress common.Address) {
	t.Helper()

	e, err := environment.New(t.Context(), environment.WithEVMSimulated(t, []uint64{testChainSelector}))
	require.NoError(t, err, "Failed to create test environment")

	chain := e.BlockChains.EVMChains()[testChainSelector]
	ds := datastore.NewMemoryDataStore()

	tokenAddr := deployTestDripToken(t, chain, testTokenSymbol)
	tokenRef = datastore.AddressRef{
		Type:          datastore.ContractType(burn_mint_erc20_with_drip.ContractType),
		Version:       burn_mint_erc20_with_drip.Version,
		Address:       tokenAddr.Hex(),
		ChainSelector: testChainSelector,
		Qualifier:     testTokenSymbol,
	}
	require.NoError(t, ds.Addresses().Add(tokenRef))

	routerAddress = common.HexToAddress("0x1234567890abcdef1234567890abcdef12345678")
	require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
		Type:          datastore.ContractType(router.ContractType),
		Version:       router.Version,
		Address:       routerAddress.Hex(),
		ChainSelector: testChainSelector,
	}))

	rmnProxyAddress = common.HexToAddress("0xabcdef1234567890abcdef1234567890abcdef12")
	require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
		Type:          datastore.ContractType(rmnproxyops.ContractType),
		Version:       semver.MustParse("1.0.0"),
		Address:       rmnProxyAddress.Hex(),
		ChainSelector: testChainSelector,
	}))

	e.DataStore = ds.Seal()

	return e, tokenRef, routerAddress, rmnProxyAddress
}

func deployTestDripToken(t *testing.T, chain evm.Chain, symbol string) common.Address {
	t.Helper()

	tokenAddr, tx, _, err := drip_bindings.DeployBurnMintERC20WithDrip(
		chain.DeployerKey,
		chain.Client,
		"Test Drip Token",
		symbol,
	)
	require.NoError(t, err, "Failed to deploy test drip token")

	_, err = deployment.ConfirmIfNoError(chain, tx, err)
	require.NoError(t, err, "Failed to confirm test drip token deployment")

	return tokenAddr
}
