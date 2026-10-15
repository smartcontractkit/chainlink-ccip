package token_pool

import (
	"fmt"
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
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/type_and_version"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/router"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/burn_mint_erc20_with_drip"
	bmtp "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/burn_mint_token_pool"
	bmtpap "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/burn_mint_token_pool_and_proxy"
	lrtp "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/lock_release_token_pool"
	lrtpap "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/lock_release_token_pool_and_proxy"
	tpap "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/token_pool_and_proxy"
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
	pool, err := tpap.NewTokenPoolAndProxyContract(common.HexToAddress(poolRef.Address), chain.Client)
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

// TestDeployPlainBurnMintTokenPool covers the plain (non-proxy) v1.5.0 burn-mint pool. It is a
// separate contract from BurnMintTokenPoolAndProxy - same constructor, same base surface, but no
// legacy-proxy getPreviousPool/setPreviousPool - and many v1.5.0 tokens were deployed on it, so it
// needs the same deploy (and, through the adapter, upgrade) path as the proxy pool.
func TestDeployPlainBurnMintTokenPool(t *testing.T) {
	t.Parallel()

	e, tokenRef, routerAddress, _ := setupDeployEnv(t)

	report, err := cldf_ops.ExecuteSequence(e.OperationsBundle, DeployTokenPool, e.BlockChains, tokenapi.DeployTokenPoolInput{
		TokenRef:          &tokenRef,
		PoolType:          string(bmtp.ContractType),
		TokenPoolVersion:  utils.Version_1_5_0,
		ChainSelector:     testChainSelector,
		ExistingDataStore: e.DataStore,
	})
	require.NoError(t, err, "Failed to execute DeployTokenPool sequence")
	require.Len(t, report.Output.Addresses, 1, "Should have deployed exactly one pool")

	poolRef := report.Output.Addresses[0]
	require.Equal(t, string(bmtp.ContractType), string(poolRef.Type))
	require.Equal(t, utils.Version_1_5_0.String(), poolRef.Version.String())
	require.Equal(t, tokenRef.Address, poolRef.Qualifier, "Pool qualifier should default to the token address")

	chain := e.BlockChains.EVMChains()[testChainSelector]
	poolAddr := common.HexToAddress(poolRef.Address)
	requireTypeAndVersion(t, chain, poolAddr, "BurnMintTokenPool 1.5.0")

	// The shared base surface drives the plain pool too.
	base, err := tpap.NewTokenPoolAndProxyContract(poolAddr, chain.Client)
	require.NoError(t, err)
	onChainToken, err := base.GetToken(&bind.CallOpts{})
	require.NoError(t, err)
	require.Equal(t, common.HexToAddress(tokenRef.Address), onChainToken, "Token address mismatch")
	onChainRouter, err := base.GetRouter(&bind.CallOpts{})
	require.NoError(t, err)
	require.Equal(t, routerAddress, onChainRouter, "Router address mismatch")

	// The one base-ops function the plain pool does NOT have. Pinned so nobody starts relying on
	// it in the pool-agnostic adapter path.
	_, err = base.GetPreviousPool(&bind.CallOpts{})
	require.Error(t, err, "plain v1.5.0 pools have no getPreviousPool")
}

// TestDeployLockReleaseTokenPool covers both v1.5.0 lock-release variants (the *AndProxy pool and
// the plain pool), whose constructor differs from the burn-mint one by the immutable
// acceptLiquidity flag.
func TestDeployLockReleaseTokenPool(t *testing.T) {
	t.Parallel()

	poolTypes := []struct {
		contractType   deployment.ContractType
		typeAndVersion string
	}{
		{contractType: lrtpap.ContractType, typeAndVersion: "LockReleaseTokenPoolAndProxy 1.5.0"},
		{contractType: lrtp.ContractType, typeAndVersion: "LockReleaseTokenPool 1.5.0"},
	}

	for _, pt := range poolTypes {
		for _, acceptLiquidity := range []bool{true, false} {
			t.Run(fmt.Sprintf("%s/acceptLiquidity=%v", pt.contractType, acceptLiquidity), func(t *testing.T) {
				t.Parallel()

				e, tokenRef, routerAddress, _ := setupDeployEnv(t)

				report, err := cldf_ops.ExecuteSequence(e.OperationsBundle, DeployTokenPool, e.BlockChains, tokenapi.DeployTokenPoolInput{
					TokenRef:          &tokenRef,
					PoolType:          string(pt.contractType),
					TokenPoolVersion:  utils.Version_1_5_0,
					ChainSelector:     testChainSelector,
					ExistingDataStore: e.DataStore,
					AcceptLiquidity:   &acceptLiquidity,
				})
				require.NoError(t, err, "Failed to execute DeployTokenPool sequence")
				require.Len(t, report.Output.Addresses, 1, "Should have deployed exactly one pool")

				poolRef := report.Output.Addresses[0]
				require.Equal(t, string(pt.contractType), string(poolRef.Type))
				require.Equal(t, utils.Version_1_5_0.String(), poolRef.Version.String())

				chain := e.BlockChains.EVMChains()[testChainSelector]
				poolAddr := common.HexToAddress(poolRef.Address)
				requireTypeAndVersion(t, chain, poolAddr, pt.typeAndVersion)

				// The shared base surface must work against a lock-release pool too - that is the
				// whole premise of routing every v1.5.0 pool type through TokenPoolAndProxy.
				base, err := tpap.NewTokenPoolAndProxyContract(poolAddr, chain.Client)
				require.NoError(t, err)

				onChainToken, err := base.GetToken(&bind.CallOpts{})
				require.NoError(t, err)
				require.Equal(t, common.HexToAddress(tokenRef.Address), onChainToken, "Token address mismatch")

				onChainRouter, err := base.GetRouter(&bind.CallOpts{})
				require.NoError(t, err)
				require.Equal(t, routerAddress, onChainRouter, "Router address mismatch")

				// Both lock-release types share the canAcceptLiquidity selector, so either
				// binding reads either pool.
				lr, err := lrtp.NewLockReleaseTokenPoolContract(poolAddr, chain.Client)
				require.NoError(t, err)

				canAccept, err := lr.CanAcceptLiquidity(&bind.CallOpts{})
				require.NoError(t, err)
				require.Equal(t, acceptLiquidity, canAccept, "acceptLiquidity constructor arg not honoured")
			})
		}
	}
}

// TestDeployLockReleaseTokenPool_RequiresAcceptLiquidity asserts the flag is demanded rather than
// defaulted for both v1.5.0 lock-release types: it is immutable on-chain, so guessing wrong is
// unrecoverable.
func TestDeployLockReleaseTokenPool_RequiresAcceptLiquidity(t *testing.T) {
	t.Parallel()

	for _, poolType := range []deployment.ContractType{lrtpap.ContractType, lrtp.ContractType} {
		t.Run(string(poolType), func(t *testing.T) {
			t.Parallel()

			e, tokenRef, _, _ := setupDeployEnv(t)

			_, err := cldf_ops.ExecuteSequence(e.OperationsBundle, DeployTokenPool, e.BlockChains, tokenapi.DeployTokenPoolInput{
				TokenRef:          &tokenRef,
				PoolType:          string(poolType),
				TokenPoolVersion:  utils.Version_1_5_0,
				ChainSelector:     testChainSelector,
				ExistingDataStore: e.DataStore,
			})
			require.ErrorContains(t, err, fmt.Sprintf("AcceptLiquidity is required when deploying %s v1.5.0", poolType))
		})
	}
}

// TestDeployTokenPool_RejectsUnsupportedTypes locks in the scope: v1.5.0 support covers
// BurnMintTokenPoolAndProxy, LockReleaseTokenPoolAndProxy, and the plain BurnMintTokenPool and
// LockReleaseTokenPool. The rebasing *AndProxy variant has bindings but no adapter support, so
// deploying it would produce a pool no changeset could subsequently configure.
func TestDeployTokenPool_RejectsUnsupportedTypes(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name        string
		poolType    string
		poolVersion *semver.Version
	}{
		{name: "BurnWithFromMintTokenPoolAndProxy", poolType: "BurnWithFromMintTokenPoolAndProxy", poolVersion: utils.Version_1_5_0},
		{name: "UnknownType", poolType: "NotARealTokenPool", poolVersion: utils.Version_1_5_0},
		// Right type, wrong version: the switch keys on the full "Type Version" string.
		{name: "RightTypeWrongVersion", poolType: string(bmtpap.ContractType), poolVersion: utils.Version_1_5_1},
		{name: "RightLockReleaseTypeWrongVersion", poolType: string(lrtpap.ContractType), poolVersion: utils.Version_1_5_1},
		// The plain types exist at v1.5.1 too, but that is the v1.5.1 sequence's job, not this one's.
		{name: "RightPlainBurnMintTypeWrongVersion", poolType: string(bmtp.ContractType), poolVersion: utils.Version_1_5_1},
		{name: "RightPlainLockReleaseTypeWrongVersion", poolType: string(lrtp.ContractType), poolVersion: utils.Version_1_5_1},
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

// requireTypeAndVersion asserts the contract at addr reports the given typeAndVersion() string, so a
// test can prove which of the byte-identical-surface v1.5.0 pool contracts actually got deployed.
func requireTypeAndVersion(t *testing.T, chain evm.Chain, addr common.Address, want string) {
	t.Helper()

	tv, err := type_and_version.NewTypeAndVersionContract(addr, chain.Client)
	require.NoError(t, err)
	got, err := tv.TypeAndVersion(&bind.CallOpts{})
	require.NoError(t, err)
	require.Equal(t, want, got)
}
