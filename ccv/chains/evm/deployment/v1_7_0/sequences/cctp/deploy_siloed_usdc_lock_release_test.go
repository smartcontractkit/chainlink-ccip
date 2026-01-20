package cctp

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/usdc_token_pool_proxy"
	erc20_lock_box_bindings "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/gobindings/generated/latest/erc20_lock_box"
	siloed_usdc_token_pool_bindings "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/gobindings/generated/latest/siloed_usdc_token_pool"
	usdc_token_pool_proxy_bindings "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/gobindings/generated/latest/usdc_token_pool_proxy"
	contract_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	"github.com/stretchr/testify/require"
)

func TestDeploySiloedUSDCLockRelease(t *testing.T) {
	chainSelector := uint64(5009297550715157269)
	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{chainSelector}),
	)
	require.NoError(t, err, "Failed to create environment")
	require.NotNil(t, e, "Environment should be created")

	setup := setupCCTPTestEnvironment(t, e, chainSelector)
	chain := e.BlockChains.EVMChains()[chainSelector]

	proxyReport, err := operations.ExecuteOperation(
		e.OperationsBundle,
		usdc_token_pool_proxy.Deploy,
		chain,
		contract_utils.DeployInput[usdc_token_pool_proxy.ConstructorArgs]{
			TypeAndVersion: deployment.NewTypeAndVersion(usdc_token_pool_proxy.ContractType, *usdc_token_pool_proxy.Version),
			ChainSelector:  chain.Selector,
			Args: usdc_token_pool_proxy.ConstructorArgs{
				Token:        setup.USDCToken,
				Pools:        usdc_token_pool_proxy.USDCTokenPoolProxyPoolAddresses{},
				Router:       setup.Router,
				CCTPVerifier: chain.DeployerKey.From,
			},
		},
	)
	require.NoError(t, err, "Failed to deploy USDCTokenPoolProxy")

	lockReleaseSelectors := []uint64{4949039107694359620, 6433500567565415381}
	report, err := operations.ExecuteSequence(
		e.OperationsBundle,
		DeploySiloedUSDCLockRelease,
		e.BlockChains,
		DeploySiloedUSDCLockReleaseInput{
			ChainSelector:             chainSelector,
			USDCToken:                 setup.USDCToken.Hex(),
			Router:                    setup.Router.Hex(),
			RMN:                       setup.RMN.Hex(),
			USDCTokenPoolProxy:        proxyReport.Output.Address,
			SiloedUSDCTokenPool:       "",
			LockReleaseChainSelectors: lockReleaseSelectors,
		},
	)
	require.NoError(t, err, "ExecuteSequence should not error")
	require.NotEmpty(t, report.Output.SiloedPoolAddress, "SiloedUSDCTokenPool address should be set")
	require.Len(t, report.Output.LockBoxes, len(lockReleaseSelectors), "Expected lockboxes for each lock-release chain")

	siloedPoolAddr := common.HexToAddress(report.Output.SiloedPoolAddress)
	pool, err := siloed_usdc_token_pool_bindings.NewSiloedUSDCTokenPool(siloedPoolAddr, chain.Client)
	require.NoError(t, err, "Failed to instantiate SiloedUSDCTokenPool contract")

	poolCallers, err := pool.GetAllAuthorizedCallers(nil)
	require.NoError(t, err, "Failed to get authorized callers from SiloedUSDCTokenPool")
	require.Contains(t, poolCallers, common.HexToAddress(proxyReport.Output.Address), "USDCTokenPoolProxy should be an authorized caller")

	for _, sel := range lockReleaseSelectors {
		lockBoxAddr, ok := report.Output.LockBoxes[sel]
		require.True(t, ok, "Lockbox should be recorded for chain %d", sel)
		lockBoxFromPool, err := pool.GetLockBox(nil, sel)
		require.NoError(t, err, "Failed to get lockbox from SiloedUSDCTokenPool")
		require.Equal(t, common.HexToAddress(lockBoxAddr), lockBoxFromPool, "Lockbox address should match pool config")

		lockBox, err := erc20_lock_box_bindings.NewERC20LockBox(common.HexToAddress(lockBoxAddr), chain.Client)
		require.NoError(t, err, "Failed to instantiate ERC20LockBox contract")
		lockBoxCallers, err := lockBox.GetAllAuthorizedCallers(nil)
		require.NoError(t, err, "Failed to get authorized callers from ERC20LockBox")
		require.Contains(t, lockBoxCallers, siloedPoolAddr, "SiloedUSDCTokenPool should be an authorized caller on lockbox")
	}

	proxy, err := usdc_token_pool_proxy_bindings.NewUSDCTokenPoolProxy(common.HexToAddress(proxyReport.Output.Address), chain.Client)
	require.NoError(t, err, "Failed to instantiate USDCTokenPoolProxy contract")
	expectedMechanism, err := convertMechanismToUint8("LOCK_RELEASE")
	require.NoError(t, err, "Failed to convert mechanism to uint8")
	for _, sel := range lockReleaseSelectors {
		mechanism, err := proxy.GetLockOrBurnMechanism(nil, sel)
		require.NoError(t, err, "Failed to get lock or burn mechanism from USDCTokenPoolProxy")
		require.Equal(t, expectedMechanism, uint8(mechanism), "Lock or burn mechanism should match for lock release chain")
	}

	pools, err := proxy.GetPools(nil)
	require.NoError(t, err, "Failed to get pools from USDCTokenPoolProxy")
	require.Equal(t, siloedPoolAddr, pools.SiloedLockReleasePool, "Siloed LockRelease pool should match")
}

func TestDeploySiloedUSDCLockReleaseWithoutProxy(t *testing.T) {
	chainSelector := uint64(5009297550715157269)
	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{chainSelector}),
	)
	require.NoError(t, err, "Failed to create environment")
	require.NotNil(t, e, "Environment should be created")

	setup := setupCCTPTestEnvironment(t, e, chainSelector)
	chain := e.BlockChains.EVMChains()[chainSelector]

	lockReleaseSelectors := []uint64{4949039107694359620, 6433500567565415381}
	report, err := operations.ExecuteSequence(
		e.OperationsBundle,
		DeploySiloedUSDCLockRelease,
		e.BlockChains,
		DeploySiloedUSDCLockReleaseInput{
			ChainSelector:             chainSelector,
			USDCToken:                 setup.USDCToken.Hex(),
			Router:                    setup.Router.Hex(),
			RMN:                       setup.RMN.Hex(),
			USDCTokenPoolProxy:        "",
			SiloedUSDCTokenPool:       "",
			LockReleaseChainSelectors: lockReleaseSelectors,
		},
	)
	require.NoError(t, err, "ExecuteSequence should not error")
	require.NotEmpty(t, report.Output.SiloedPoolAddress, "SiloedUSDCTokenPool address should be set")
	require.Len(t, report.Output.LockBoxes, len(lockReleaseSelectors), "Expected lockboxes for each lock-release chain")

	siloedPoolAddr := common.HexToAddress(report.Output.SiloedPoolAddress)
	pool, err := siloed_usdc_token_pool_bindings.NewSiloedUSDCTokenPool(siloedPoolAddr, chain.Client)
	require.NoError(t, err, "Failed to instantiate SiloedUSDCTokenPool contract")

	poolCallers, err := pool.GetAllAuthorizedCallers(nil)
	require.NoError(t, err, "Failed to get authorized callers from SiloedUSDCTokenPool")
	require.Empty(t, poolCallers, "SiloedUSDCTokenPool should not have authorized callers before proxy wiring")

	for _, sel := range lockReleaseSelectors {
		lockBoxAddr, ok := report.Output.LockBoxes[sel]
		require.True(t, ok, "Lockbox should be recorded for chain %d", sel)
		lockBoxFromPool, err := pool.GetLockBox(nil, sel)
		require.NoError(t, err, "Failed to get lockbox from SiloedUSDCTokenPool")
		require.Equal(t, common.HexToAddress(lockBoxAddr), lockBoxFromPool, "Lockbox address should match pool config")

		lockBox, err := erc20_lock_box_bindings.NewERC20LockBox(common.HexToAddress(lockBoxAddr), chain.Client)
		require.NoError(t, err, "Failed to instantiate ERC20LockBox contract")
		lockBoxCallers, err := lockBox.GetAllAuthorizedCallers(nil)
		require.NoError(t, err, "Failed to get authorized callers from ERC20LockBox")
		require.Contains(t, lockBoxCallers, siloedPoolAddr, "SiloedUSDCTokenPool should be an authorized caller on lockbox")
	}
}
