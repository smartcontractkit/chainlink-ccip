package tokens_test

import (
	"math/big"
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/rmn_proxy"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/router"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/burn_mint_erc20_with_drip"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/create2_factory"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/erc20_lock_box"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/siloed_lock_release_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/sequences"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/sequences/tokens"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/testsetup"
)

// Arbitrary remote chain selectors; the pool only stores them as map keys.
const (
	siloedRemoteChainA = uint64(1111)
	siloedRemoteChainB = uint64(2222)
	siloedRemoteChainC = uint64(3333)
)

// setupSiloedPoolDeps stands up a simulated chain with the contracts the siloed pool constructor
// needs, and returns the base input the tests then vary.
func setupSiloedPoolDeps(t *testing.T, chainSel uint64) (*deployment.Environment, tokens.DeployTokenPoolInput) {
	t.Helper()

	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{chainSel}),
	)
	require.NoError(t, err, "Failed to create environment")
	require.NotNil(t, e, "Environment should be created")

	create2FactoryRef, err := contract.MaybeDeployContract(e.OperationsBundle, create2_factory.Deploy, e.BlockChains.EVMChains()[chainSel], contract.DeployInput[create2_factory.ConstructorArgs]{
		TypeAndVersion: deployment.NewTypeAndVersion(create2_factory.ContractType, *semver.MustParse("2.0.0")),
		ChainSelector:  chainSel,
		Args: create2_factory.ConstructorArgs{
			AllowList: []common.Address{e.BlockChains.EVMChains()[chainSel].DeployerKey.From},
		},
	}, nil)
	require.NoError(t, err, "Failed to deploy CREATE2Factory")

	chainReport, err := operations.ExecuteSequence(
		e.OperationsBundle,
		sequences.DeployChainContracts,
		e.BlockChains.EVMChains()[chainSel],
		sequences.DeployChainContractsInput{
			ChainSelector:    chainSel,
			CREATE2Factory:   common.HexToAddress(create2FactoryRef.Address),
			ContractParams:   testsetup.CreateBasicContractParams(),
			DeployerKeyOwned: true,
		},
	)
	require.NoError(t, err, "ExecuteSequence should not error")

	tokenReport, err := operations.ExecuteOperation(
		e.OperationsBundle,
		burn_mint_erc20_with_drip.Deploy,
		e.BlockChains.EVMChains()[chainSel],
		contract.DeployInput[burn_mint_erc20_with_drip.ConstructorArgs]{
			ChainSelector:  chainSel,
			TypeAndVersion: deployment.NewTypeAndVersion(burn_mint_erc20_with_drip.ContractType, *burn_mint_erc20_with_drip.Version),
			Args: burn_mint_erc20_with_drip.ConstructorArgs{
				Name:   "Test Token",
				Symbol: "TEST",
			},
		},
	)
	require.NoError(t, err, "ExecuteOperation should not error")

	var rmnProxyAddress, routerAddress common.Address
	for _, addr := range chainReport.Output.Addresses {
		if addr.Type == datastore.ContractType(rmn_proxy.ContractType) {
			rmnProxyAddress = common.HexToAddress(addr.Address)
		}
		if addr.Type == datastore.ContractType(router.ContractType) {
			routerAddress = common.HexToAddress(addr.Address)
		}
	}

	return e, tokens.DeployTokenPoolInput{
		ChainSel:                         chainSel,
		TokenPoolType:                    datastore.ContractType(siloed_lock_release_token_pool.ContractType),
		TokenPoolVersion:                 siloed_lock_release_token_pool.Version,
		TokenSymbol:                      tokenReport.Input.Args.Symbol,
		RateLimitAdmin:                   common.HexToAddress("0x01"),
		ThresholdAmountForAdditionalCCVs: big.NewInt(1e18),
		ConstructorArgs: tokens.ConstructorArgs{
			Token:    common.HexToAddress(tokenReport.Output.Address),
			Decimals: 18,
			RMNProxy: rmnProxyAddress,
			Router:   routerAddress,
		},
	}
}

func TestDeploySiloedLockReleaseTokenPool(t *testing.T) {
	chainSel := uint64(5009297550715157269)

	t.Run("happy path - chains A and B share a lock box, C is siloed", func(t *testing.T) {
		e, input := setupSiloedPoolDeps(t, chainSel)
		input.LockBoxGroups = [][]uint64{
			{siloedRemoteChainA, siloedRemoteChainB},
			{siloedRemoteChainC},
		}

		poolReport, err := operations.ExecuteSequence(
			e.OperationsBundle,
			tokens.DeploySiloedLockReleaseTokenPool,
			e.BlockChains.EVMChains()[chainSel],
			input,
		)
		require.NoError(t, err, "ExecuteSequence should not error")
		require.Len(t, poolReport.Output.Addresses, 4, "Expected 4 addresses in output (pool, hooks, 2 lock boxes)")

		poolAddress := common.HexToAddress(poolReport.Output.Addresses[0].Address)
		sharedLockBox := common.HexToAddress(poolReport.Output.Addresses[2].Address)
		siloedLockBox := common.HexToAddress(poolReport.Output.Addresses[3].Address)
		require.NotEqual(t, sharedLockBox, siloedLockBox, "Expected a distinct lock box per silo group")

		// The pool holds no liquidity of its own, so the token/router wiring plus the lock box
		// mapping is what makes a lane usable.
		getTokenReport, err := operations.ExecuteOperation(
			testsetup.BundleWithFreshReporter(e.OperationsBundle),
			token_pool.GetToken,
			e.BlockChains.EVMChains()[chainSel],
			contract.FunctionInput[struct{}]{ChainSelector: chainSel, Address: poolAddress},
		)
		require.NoError(t, err, "ExecuteOperation should not error")
		require.Equal(t, input.ConstructorArgs.Token, getTokenReport.Output, "Expected the pool's token to be the deployed token")

		// Every chain in a group must resolve to that group's lock box.
		for chain, expected := range map[uint64]common.Address{
			siloedRemoteChainA: sharedLockBox,
			siloedRemoteChainB: sharedLockBox,
			siloedRemoteChainC: siloedLockBox,
		} {
			getLockBoxReport, err := operations.ExecuteOperation(
				testsetup.BundleWithFreshReporter(e.OperationsBundle),
				siloed_lock_release_token_pool.GetLockBox,
				e.BlockChains.EVMChains()[chainSel],
				contract.FunctionInput[uint64]{ChainSelector: chainSel, Address: poolAddress, Args: chain},
			)
			require.NoError(t, err, "ExecuteOperation should not error for chain %d", chain)
			require.Equal(t, expected, getLockBoxReport.Output, "Unexpected lock box for remote chain %d", chain)
		}

		getAllConfigsReport, err := operations.ExecuteOperation(
			testsetup.BundleWithFreshReporter(e.OperationsBundle),
			siloed_lock_release_token_pool.GetAllLockBoxConfigs,
			e.BlockChains.EVMChains()[chainSel],
			contract.FunctionInput[struct{}]{ChainSelector: chainSel, Address: poolAddress},
		)
		require.NoError(t, err, "ExecuteOperation should not error")
		require.Len(t, getAllConfigsReport.Output, 3, "Expected one lock box config per remote chain")

		// The pool must be able to deposit into and withdraw from each of its lock boxes.
		for _, lockBox := range []common.Address{sharedLockBox, siloedLockBox} {
			getCallersReport, err := operations.ExecuteOperation(
				testsetup.BundleWithFreshReporter(e.OperationsBundle),
				erc20_lock_box.GetAllAuthorizedCallers,
				e.BlockChains.EVMChains()[chainSel],
				contract.FunctionInput[struct{}]{ChainSelector: chainSel, Address: lockBox},
			)
			require.NoError(t, err, "ExecuteOperation should not error")
			require.Equal(t, []common.Address{poolAddress}, getCallersReport.Output, "Expected only the pool to be authorized on lock box %s", lockBox)
		}
	})

	t.Run("lock box groups not defined", func(t *testing.T) {
		e, input := setupSiloedPoolDeps(t, chainSel)

		_, err := operations.ExecuteSequence(
			e.OperationsBundle,
			tokens.DeploySiloedLockReleaseTokenPool,
			e.BlockChains.EVMChains()[chainSel],
			input,
		)
		require.Error(t, err, "ExecuteSequence should error")
		require.Contains(t, err.Error(), "lock box groups must be defined")
	})

	t.Run("remote chain in more than one lock box group", func(t *testing.T) {
		e, input := setupSiloedPoolDeps(t, chainSel)
		input.LockBoxGroups = [][]uint64{
			{siloedRemoteChainA, siloedRemoteChainB},
			{siloedRemoteChainB},
		}

		_, err := operations.ExecuteSequence(
			e.OperationsBundle,
			tokens.DeploySiloedLockReleaseTokenPool,
			e.BlockChains.EVMChains()[chainSel],
			input,
		)
		require.Error(t, err, "ExecuteSequence should error")
		require.Contains(t, err.Error(), "appears in lock box groups 0 and 1")
	})

	t.Run("empty lock box group", func(t *testing.T) {
		e, input := setupSiloedPoolDeps(t, chainSel)
		input.LockBoxGroups = [][]uint64{{siloedRemoteChainA}, {}}

		_, err := operations.ExecuteSequence(
			e.OperationsBundle,
			tokens.DeploySiloedLockReleaseTokenPool,
			e.BlockChains.EVMChains()[chainSel],
			input,
		)
		require.Error(t, err, "ExecuteSequence should error")
		require.Contains(t, err.Error(), "lock box group 1 is empty")
	})
}
