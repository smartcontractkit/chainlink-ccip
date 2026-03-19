package tokens_test

import (
	"math/big"
	"strings"
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"

	evm_contract "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/rmn_proxy"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/router"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/burn_mint_erc20_with_drip"
	tar "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/token_admin_registry"
	old_lrtp "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_1/operations/lock_release_token_pool"
	old_siloed "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_1/operations/siloed_lock_release_token_pool"

	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/latest/operations/erc20"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/create2_factory"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/sequences"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/sequences/tokens"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/testsetup"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v2_0_0/operations/erc20_lock_box"
	new_lrtp "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v2_0_0/operations/lock_release_token_pool"
	latest_siloed "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/gobindings/generated/latest/siloed_lock_release_token_pool"
	tokens_core "github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

func TestMigrateLockReleasePoolLiquidity_Validation(t *testing.T) {
	chainSel := uint64(5009297550715157269)
	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{chainSel}),
	)
	require.NoError(t, err)

	tests := []struct {
		name        string
		input       tokens_core.MigrateLockReleasePoolLiquidityInput
		expectedErr string
	}{
		{
			name: "both Amount and BasisPoints provided",
			input: tokens_core.MigrateLockReleasePoolLiquidityInput{
				ChainSelector:   chainSel,
				OldPoolAddress:  "0x0000000000000000000000000000000000000001",
				NewPoolAddress:  "0x0000000000000000000000000000000000000002",
				TimelockAddress: "0x0000000000000000000000000000000000000003",
				Amount:          big.NewInt(100),
				BasisPoints:     uint16Ptr(5000),
			},
			expectedErr: "Amount and BasisPoints are mutually exclusive",
		},
		{
			name: "neither Amount nor BasisPoints provided",
			input: tokens_core.MigrateLockReleasePoolLiquidityInput{
				ChainSelector:   chainSel,
				OldPoolAddress:  "0x0000000000000000000000000000000000000001",
				NewPoolAddress:  "0x0000000000000000000000000000000000000002",
				TimelockAddress: "0x0000000000000000000000000000000000000003",
			},
			expectedErr: "one of Amount or BasisPoints must be provided",
		},
		{
			name: "BasisPoints zero",
			input: tokens_core.MigrateLockReleasePoolLiquidityInput{
				ChainSelector:   chainSel,
				OldPoolAddress:  "0x0000000000000000000000000000000000000001",
				NewPoolAddress:  "0x0000000000000000000000000000000000000002",
				TimelockAddress: "0x0000000000000000000000000000000000000003",
				BasisPoints:     uint16Ptr(0),
			},
			expectedErr: "BasisPoints must be between 1 and 10000",
		},
		{
			name: "BasisPoints exceeds 10000",
			input: tokens_core.MigrateLockReleasePoolLiquidityInput{
				ChainSelector:   chainSel,
				OldPoolAddress:  "0x0000000000000000000000000000000000000001",
				NewPoolAddress:  "0x0000000000000000000000000000000000000002",
				TimelockAddress: "0x0000000000000000000000000000000000000003",
				BasisPoints:     uint16Ptr(10001),
			},
			expectedErr: "BasisPoints must be between 1 and 10000",
		},
		{
			name: "Amount is zero",
			input: tokens_core.MigrateLockReleasePoolLiquidityInput{
				ChainSelector:   chainSel,
				OldPoolAddress:  "0x0000000000000000000000000000000000000001",
				NewPoolAddress:  "0x0000000000000000000000000000000000000002",
				TimelockAddress: "0x0000000000000000000000000000000000000003",
				Amount:          big.NewInt(0),
			},
			expectedErr: "Amount must be positive",
		},
		{
			name: "Amount is negative",
			input: tokens_core.MigrateLockReleasePoolLiquidityInput{
				ChainSelector:   chainSel,
				OldPoolAddress:  "0x0000000000000000000000000000000000000001",
				NewPoolAddress:  "0x0000000000000000000000000000000000000002",
				TimelockAddress: "0x0000000000000000000000000000000000000003",
				Amount:          big.NewInt(-1),
			},
			expectedErr: "Amount must be positive",
		},
		{
			name: "missing OldPoolAddress",
			input: tokens_core.MigrateLockReleasePoolLiquidityInput{
				ChainSelector:   chainSel,
				NewPoolAddress:  "0x0000000000000000000000000000000000000002",
				TimelockAddress: "0x0000000000000000000000000000000000000003",
				BasisPoints:     uint16Ptr(10000),
			},
			expectedErr: "OldPoolAddress and NewPoolAddress must be provided",
		},
		{
			name: "missing TimelockAddress",
			input: tokens_core.MigrateLockReleasePoolLiquidityInput{
				ChainSelector:  chainSel,
				OldPoolAddress: "0x0000000000000000000000000000000000000001",
				NewPoolAddress: "0x0000000000000000000000000000000000000002",
				BasisPoints:    uint16Ptr(10000),
			},
			expectedErr: "TimelockAddress must be provided",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			_, err := operations.ExecuteSequence(
				e.OperationsBundle,
				tokens.MigrateLockReleasePoolLiquidity,
				e.BlockChains,
				tc.input,
			)
			require.Error(t, err)
			require.Contains(t, err.Error(), tc.expectedErr)
		})
	}
}

type migrationTestSetup struct {
	env         *deployment.Environment
	chainSel    uint64
	deployer    common.Address
	tokenAddr   common.Address
	oldPoolAddr common.Address
	newPoolAddr common.Address
	lockBoxAddr common.Address
}

// setupMigrationTest deploys all contracts needed for a migration test:
// ERC20 token, old v1.6.1 LockReleaseTokenPool, and new v2.0 LockReleaseTokenPool
// with its lockbox. Mints liquidityAmount tokens directly into the old pool.
func setupMigrationTest(t *testing.T, chainSel uint64, liquidityAmount *big.Int) migrationTestSetup {
	t.Helper()

	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{chainSel}),
	)
	require.NoError(t, err)

	chain := e.BlockChains.EVMChains()[chainSel]
	deployer := chain.DeployerKey.From

	create2FactoryRef, err := evm_contract.MaybeDeployContract(e.OperationsBundle, create2_factory.Deploy, chain, evm_contract.DeployInput[create2_factory.ConstructorArgs]{
		TypeAndVersion: deployment.NewTypeAndVersion(create2_factory.ContractType, *semver.MustParse("2.0.0")),
		ChainSelector:  chainSel,
		Args: create2_factory.ConstructorArgs{
			AllowList: []common.Address{deployer},
		},
	}, nil)
	require.NoError(t, err)

	chainReport, err := operations.ExecuteSequence(
		e.OperationsBundle,
		sequences.DeployChainContracts,
		chain,
		sequences.DeployChainContractsInput{
			ChainSelector:    chainSel,
			ContractParams:   testsetup.CreateBasicContractParams(),
			CREATE2Factory:   common.HexToAddress(create2FactoryRef.Address),
			DeployerKeyOwned: true,
		},
	)
	require.NoError(t, err)

	var rmnProxyAddr, routerAddr common.Address
	for _, addr := range chainReport.Output.Addresses {
		switch addr.Type {
		case datastore.ContractType(rmn_proxy.ContractType):
			rmnProxyAddr = common.HexToAddress(addr.Address)
		case datastore.ContractType(router.ContractType):
			routerAddr = common.HexToAddress(addr.Address)
		}
	}

	tokenReport, err := operations.ExecuteOperation(
		e.OperationsBundle,
		burn_mint_erc20_with_drip.Deploy,
		chain,
		evm_contract.DeployInput[burn_mint_erc20_with_drip.ConstructorArgs]{
			ChainSelector:  chainSel,
			TypeAndVersion: deployment.NewTypeAndVersion(burn_mint_erc20_with_drip.ContractType, *burn_mint_erc20_with_drip.Version),
			Args: burn_mint_erc20_with_drip.ConstructorArgs{
				Name:   "Test Token",
				Symbol: "TEST",
			},
		},
	)
	require.NoError(t, err)
	tokenAddr := common.HexToAddress(tokenReport.Output.Address)

	oldPoolReport, err := operations.ExecuteOperation(
		e.OperationsBundle,
		old_lrtp.Deploy,
		chain,
		evm_contract.DeployInput[old_lrtp.ConstructorArgs]{
			ChainSelector:  chainSel,
			TypeAndVersion: deployment.NewTypeAndVersion(old_lrtp.ContractType, *old_lrtp.Version),
			Args: old_lrtp.ConstructorArgs{
				Token:              tokenAddr,
				LocalTokenDecimals: 18,
				Allowlist:          []common.Address{},
				RmnProxy:           rmnProxyAddr,
				Router:             routerAddr,
			},
		},
	)
	require.NoError(t, err)
	oldPoolAddr := common.HexToAddress(oldPoolReport.Output.Address)

	newPoolReport, err := operations.ExecuteSequence(
		e.OperationsBundle,
		tokens.DeployLockReleaseTokenPool,
		chain,
		tokens.DeployTokenPoolInput{
			ChainSel:                         chainSel,
			TokenPoolType:                    datastore.ContractType(new_lrtp.ContractType),
			TokenPoolVersion:                 new_lrtp.Version,
			TokenSymbol:                      "TEST",
			ThresholdAmountForAdditionalCCVs: big.NewInt(1e18),
			ConstructorArgs: tokens.ConstructorArgs{
				Token:    tokenAddr,
				Decimals: 18,
				RMNProxy: rmnProxyAddr,
				Router:   routerAddr,
			},
		},
	)
	require.NoError(t, err)
	newPoolAddr := common.HexToAddress(newPoolReport.Output.Addresses[0].Address)
	lockBoxAddr := common.HexToAddress(newPoolReport.Output.Addresses[2].Address)

	// Grant mint role to deployer and mint tokens into the old pool
	_, err = operations.ExecuteOperation(
		e.OperationsBundle,
		burn_mint_erc20_with_drip.GrantMintAndBurnRoles,
		chain,
		evm_contract.FunctionInput[common.Address]{
			ChainSelector: chainSel,
			Address:       tokenAddr,
			Args:          deployer,
		},
	)
	require.NoError(t, err)

	_, err = operations.ExecuteOperation(
		e.OperationsBundle,
		burn_mint_erc20_with_drip.Mint,
		chain,
		evm_contract.FunctionInput[burn_mint_erc20_with_drip.MintArgs]{
			ChainSelector: chainSel,
			Address:       tokenAddr,
			Args: burn_mint_erc20_with_drip.MintArgs{
				Account: oldPoolAddr,
				Amount:  liquidityAmount,
			},
		},
	)
	require.NoError(t, err)

	// Verify old pool holds the minted tokens
	balReport, err := operations.ExecuteOperation(
		testsetup.BundleWithFreshReporter(e.OperationsBundle),
		erc20.BalanceOf,
		chain,
		evm_contract.FunctionInput[common.Address]{
			ChainSelector: chainSel,
			Address:       tokenAddr,
			Args:          oldPoolAddr,
		},
	)
	require.NoError(t, err)
	require.Equal(t, 0, liquidityAmount.Cmp(balReport.Output), "Old pool should hold the minted tokens")

	return migrationTestSetup{
		env:         e,
		chainSel:    chainSel,
		deployer:    deployer,
		tokenAddr:   tokenAddr,
		oldPoolAddr: oldPoolAddr,
		newPoolAddr: newPoolAddr,
		lockBoxAddr: lockBoxAddr,
	}
}

func TestMigrateLockReleasePoolLiquidity_UnsiloedPartialBasisPoints(t *testing.T) {
	chainSel := uint64(5009297550715157269)
	totalLiquidity := big.NewInt(10000)
	s := setupMigrationTest(t, chainSel, totalLiquidity)
	chain := s.env.BlockChains.EVMChains()[chainSel]

	basisPoints := uint16(8000) // 80%
	_, err := operations.ExecuteSequence(
		s.env.OperationsBundle,
		tokens.MigrateLockReleasePoolLiquidity,
		s.env.BlockChains,
		tokens_core.MigrateLockReleasePoolLiquidityInput{
			ChainSelector:   chainSel,
			OldPoolAddress:  s.oldPoolAddr.Hex(),
			NewPoolAddress:  s.newPoolAddr.Hex(),
			TimelockAddress: s.deployer.Hex(),
			BasisPoints:     &basisPoints,
		},
	)
	require.NoError(t, err)

	expectedMigrated := new(big.Int).Div(
		new(big.Int).Mul(totalLiquidity, big.NewInt(8000)),
		big.NewInt(10000),
	)
	expectedRemaining := new(big.Int).Sub(totalLiquidity, expectedMigrated)

	oldPoolBal, err := operations.ExecuteOperation(
		testsetup.BundleWithFreshReporter(s.env.OperationsBundle),
		erc20.BalanceOf,
		chain,
		evm_contract.FunctionInput[common.Address]{
			ChainSelector: chainSel,
			Address:       s.tokenAddr,
			Args:          s.oldPoolAddr,
		},
	)
	require.NoError(t, err)
	require.Equal(t, 0, expectedRemaining.Cmp(oldPoolBal.Output), "Old pool should retain 20%% of liquidity")

	lockboxBal, err := operations.ExecuteOperation(
		testsetup.BundleWithFreshReporter(s.env.OperationsBundle),
		erc20.BalanceOf,
		chain,
		evm_contract.FunctionInput[common.Address]{
			ChainSelector: chainSel,
			Address:       s.tokenAddr,
			Args:          s.lockBoxAddr,
		},
	)
	require.NoError(t, err)
	require.Equal(t, 0, expectedMigrated.Cmp(lockboxBal.Output), "Lockbox should hold 80%% of liquidity")

	// Verify rebalancer was restored to original (zero address, since we didn't set one)
	rebalancerReport, err := operations.ExecuteOperation(
		testsetup.BundleWithFreshReporter(s.env.OperationsBundle),
		old_lrtp.GetRebalancer,
		chain,
		evm_contract.FunctionInput[struct{}]{
			ChainSelector: chainSel,
			Address:       s.oldPoolAddr,
		},
	)
	require.NoError(t, err)
	require.Equal(t, common.Address{}, rebalancerReport.Output,
		"Rebalancer should be restored to original value (zero address)")

	// Verify timelock was removed from lockbox authorized callers
	authCallersReport, err := operations.ExecuteOperation(
		testsetup.BundleWithFreshReporter(s.env.OperationsBundle),
		erc20_lock_box.GetAllAuthorizedCallers,
		chain,
		evm_contract.FunctionInput[struct{}]{
			ChainSelector: chainSel,
			Address:       s.lockBoxAddr,
		},
	)
	require.NoError(t, err)
	require.NotContains(t, authCallersReport.Output, s.deployer,
		"Timelock should be removed from lockbox authorized callers after migration")
}

func TestMigrateLockReleasePoolLiquidity_UnsiloedFullBasisPoints(t *testing.T) {
	chainSel := uint64(5009297550715157269)
	totalLiquidity := big.NewInt(5000)
	s := setupMigrationTest(t, chainSel, totalLiquidity)
	chain := s.env.BlockChains.EVMChains()[chainSel]

	basisPoints := uint16(10000) // 100%
	_, err := operations.ExecuteSequence(
		s.env.OperationsBundle,
		tokens.MigrateLockReleasePoolLiquidity,
		s.env.BlockChains,
		tokens_core.MigrateLockReleasePoolLiquidityInput{
			ChainSelector:   chainSel,
			OldPoolAddress:  s.oldPoolAddr.Hex(),
			NewPoolAddress:  s.newPoolAddr.Hex(),
			TimelockAddress: s.deployer.Hex(),
			BasisPoints:     &basisPoints,
		},
	)
	require.NoError(t, err)

	oldPoolBal, err := operations.ExecuteOperation(
		testsetup.BundleWithFreshReporter(s.env.OperationsBundle),
		erc20.BalanceOf,
		chain,
		evm_contract.FunctionInput[common.Address]{
			ChainSelector: chainSel,
			Address:       s.tokenAddr,
			Args:          s.oldPoolAddr,
		},
	)
	require.NoError(t, err)
	require.Equal(t, 0, big.NewInt(0).Cmp(oldPoolBal.Output), "Old pool should be fully drained")

	lockboxBal, err := operations.ExecuteOperation(
		testsetup.BundleWithFreshReporter(s.env.OperationsBundle),
		erc20.BalanceOf,
		chain,
		evm_contract.FunctionInput[common.Address]{
			ChainSelector: chainSel,
			Address:       s.tokenAddr,
			Args:          s.lockBoxAddr,
		},
	)
	require.NoError(t, err)
	require.Equal(t, 0, totalLiquidity.Cmp(lockboxBal.Output), "Lockbox should hold all liquidity")
}

func TestMigrateLockReleasePoolLiquidity_ExactAmount(t *testing.T) {
	chainSel := uint64(5009297550715157269)
	totalLiquidity := big.NewInt(10000)
	s := setupMigrationTest(t, chainSel, totalLiquidity)
	chain := s.env.BlockChains.EVMChains()[chainSel]

	exactAmount := big.NewInt(2500)
	_, err := operations.ExecuteSequence(
		s.env.OperationsBundle,
		tokens.MigrateLockReleasePoolLiquidity,
		s.env.BlockChains,
		tokens_core.MigrateLockReleasePoolLiquidityInput{
			ChainSelector:   chainSel,
			OldPoolAddress:  s.oldPoolAddr.Hex(),
			NewPoolAddress:  s.newPoolAddr.Hex(),
			TimelockAddress: s.deployer.Hex(),
			Amount:          exactAmount,
		},
	)
	require.NoError(t, err)

	expectedRemaining := new(big.Int).Sub(totalLiquidity, exactAmount)

	oldPoolBal, err := operations.ExecuteOperation(
		testsetup.BundleWithFreshReporter(s.env.OperationsBundle),
		erc20.BalanceOf,
		chain,
		evm_contract.FunctionInput[common.Address]{
			ChainSelector: chainSel,
			Address:       s.tokenAddr,
			Args:          s.oldPoolAddr,
		},
	)
	require.NoError(t, err)
	require.Equal(t, 0, expectedRemaining.Cmp(oldPoolBal.Output), "Old pool should retain remaining tokens")

	lockboxBal, err := operations.ExecuteOperation(
		testsetup.BundleWithFreshReporter(s.env.OperationsBundle),
		erc20.BalanceOf,
		chain,
		evm_contract.FunctionInput[common.Address]{
			ChainSelector: chainSel,
			Address:       s.tokenAddr,
			Args:          s.lockBoxAddr,
		},
	)
	require.NoError(t, err)
	require.Equal(t, 0, exactAmount.Cmp(lockboxBal.Output), "Lockbox should hold the exact migrated amount")
}

func TestMigrateLockReleasePoolLiquidity_RebalancerRestore(t *testing.T) {
	chainSel := uint64(5009297550715157269)
	totalLiquidity := big.NewInt(1000)
	s := setupMigrationTest(t, chainSel, totalLiquidity)
	chain := s.env.BlockChains.EVMChains()[chainSel]

	// Set a non-zero original rebalancer before migration
	originalRebalancer := common.HexToAddress("0x1111111111111111111111111111111111111111")
	_, err := operations.ExecuteOperation(
		s.env.OperationsBundle,
		old_lrtp.SetRebalancer,
		chain,
		evm_contract.FunctionInput[common.Address]{
			ChainSelector: chainSel,
			Address:       s.oldPoolAddr,
			Args:          originalRebalancer,
		},
	)
	require.NoError(t, err)

	basisPoints := uint16(10000)
	// Use a fresh reporter so the test setup's SetRebalancer report doesn't
	// collide with the migration's restore-SetRebalancer call (same def+input hash).
	freshBundle := testsetup.BundleWithFreshReporter(s.env.OperationsBundle)
	_, err = operations.ExecuteSequence(
		freshBundle,
		tokens.MigrateLockReleasePoolLiquidity,
		s.env.BlockChains,
		tokens_core.MigrateLockReleasePoolLiquidityInput{
			ChainSelector:   chainSel,
			OldPoolAddress:  s.oldPoolAddr.Hex(),
			NewPoolAddress:  s.newPoolAddr.Hex(),
			TimelockAddress: s.deployer.Hex(),
			BasisPoints:     &basisPoints,
		},
	)
	require.NoError(t, err)

	rebalancerReport, err := operations.ExecuteOperation(
		testsetup.BundleWithFreshReporter(s.env.OperationsBundle),
		old_lrtp.GetRebalancer,
		chain,
		evm_contract.FunctionInput[struct{}]{
			ChainSelector: chainSel,
			Address:       s.oldPoolAddr,
		},
	)
	require.NoError(t, err)
	require.Equal(t, originalRebalancer, rebalancerReport.Output,
		"Rebalancer should be restored to the original non-zero address")
}

func TestMigrateLockReleasePoolLiquidity_AmountExceedsBalance(t *testing.T) {
	chainSel := uint64(5009297550715157269)
	totalLiquidity := big.NewInt(1000)
	s := setupMigrationTest(t, chainSel, totalLiquidity)

	tooMuch := new(big.Int).Add(totalLiquidity, big.NewInt(1))
	_, err := operations.ExecuteSequence(
		s.env.OperationsBundle,
		tokens.MigrateLockReleasePoolLiquidity,
		s.env.BlockChains,
		tokens_core.MigrateLockReleasePoolLiquidityInput{
			ChainSelector:   chainSel,
			OldPoolAddress:  s.oldPoolAddr.Hex(),
			NewPoolAddress:  s.newPoolAddr.Hex(),
			TimelockAddress: s.deployer.Hex(),
			Amount:          tooMuch,
		},
	)
	require.Error(t, err)
	require.Contains(t, err.Error(), "exceeds old pool balance")
}

// ABI struct types used by bound contract calls for siloed pool test setup.
type testRateLimiterConfig struct {
	IsEnabled bool
	Capacity  *big.Int
	Rate      *big.Int
}

type testChainUpdate struct {
	RemoteChainSelector       uint64
	RemotePoolAddresses       [][]byte
	RemoteTokenAddress        []byte
	OutboundRateLimiterConfig testRateLimiterConfig
	InboundRateLimiterConfig  testRateLimiterConfig
}

func TestMigrateLockReleasePoolLiquidity_SiloedPool(t *testing.T) {
	chainSel := uint64(5009297550715157269)
	remoteChain1 := uint64(3379446385462418246)
	remoteChain2 := uint64(4949039107694359620)
	silo1Amount := big.NewInt(3000)
	silo2Amount := big.NewInt(2000)
	unsiloedAmount := big.NewInt(1000)
	totalMint := new(big.Int).Add(new(big.Int).Add(silo1Amount, silo2Amount), unsiloedAmount)

	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{chainSel}),
	)
	require.NoError(t, err)

	chain := e.BlockChains.EVMChains()[chainSel]
	deployer := chain.DeployerKey.From

	create2FactoryRef, err := evm_contract.MaybeDeployContract(e.OperationsBundle, create2_factory.Deploy, chain, evm_contract.DeployInput[create2_factory.ConstructorArgs]{
		TypeAndVersion: deployment.NewTypeAndVersion(create2_factory.ContractType, *semver.MustParse("2.0.0")),
		ChainSelector:  chainSel,
		Args:           create2_factory.ConstructorArgs{AllowList: []common.Address{deployer}},
	}, nil)
	require.NoError(t, err)

	chainReport, err := operations.ExecuteSequence(
		e.OperationsBundle, sequences.DeployChainContracts, chain,
		sequences.DeployChainContractsInput{
			ChainSelector:    chainSel,
			ContractParams:   testsetup.CreateBasicContractParams(),
			CREATE2Factory:   common.HexToAddress(create2FactoryRef.Address),
			DeployerKeyOwned: true,
		},
	)
	require.NoError(t, err)

	var rmnProxyAddr, routerAddr common.Address
	for _, addr := range chainReport.Output.Addresses {
		switch addr.Type {
		case datastore.ContractType(rmn_proxy.ContractType):
			rmnProxyAddr = common.HexToAddress(addr.Address)
		case datastore.ContractType(router.ContractType):
			routerAddr = common.HexToAddress(addr.Address)
		}
	}

	tokenReport, err := operations.ExecuteOperation(
		e.OperationsBundle, burn_mint_erc20_with_drip.Deploy, chain,
		evm_contract.DeployInput[burn_mint_erc20_with_drip.ConstructorArgs]{
			ChainSelector:  chainSel,
			TypeAndVersion: deployment.NewTypeAndVersion(burn_mint_erc20_with_drip.ContractType, *burn_mint_erc20_with_drip.Version),
			Args:           burn_mint_erc20_with_drip.ConstructorArgs{Name: "Test Token", Symbol: "TEST"},
		},
	)
	require.NoError(t, err)
	tokenAddr := common.HexToAddress(tokenReport.Output.Address)

	// Deploy old v1.6.1 siloed pool
	oldPoolReport, err := operations.ExecuteOperation(
		e.OperationsBundle, old_siloed.Deploy, chain,
		evm_contract.DeployInput[old_siloed.ConstructorArgs]{
			ChainSelector:  chainSel,
			TypeAndVersion: deployment.NewTypeAndVersion(old_siloed.ContractType, *old_siloed.Version),
			Args: old_siloed.ConstructorArgs{
				Token: tokenAddr, LocalTokenDecimals: 18,
				Allowlist: []common.Address{}, RmnProxy: rmnProxyAddr, Router: routerAddr,
			},
		},
	)
	require.NoError(t, err)
	oldPoolAddr := common.HexToAddress(oldPoolReport.Output.Address)

	// Configure chains on old pool via bound contract
	parsed, err := abi.JSON(strings.NewReader(old_siloed.SiloedLockReleaseTokenPoolABI))
	require.NoError(t, err)
	oldPoolBound := bind.NewBoundContract(oldPoolAddr, parsed, chain.Client, chain.Client, chain.Client)

	disabledLimiter := testRateLimiterConfig{IsEnabled: false, Capacity: big.NewInt(0), Rate: big.NewInt(0)}
	dummyRemotePool := common.HexToAddress("0xdeadbeefdeadbeefdeadbeefdeadbeefdeadbeef").Bytes()
	dummyRemoteToken := common.HexToAddress("0xcafecafecafecafecafecafecafecafecafecafe").Bytes()

	tx, err := oldPoolBound.Transact(chain.DeployerKey, "applyChainUpdates",
		[]uint64{},
		[]testChainUpdate{
			{
				RemoteChainSelector: remoteChain1, RemotePoolAddresses: [][]byte{dummyRemotePool},
				RemoteTokenAddress: dummyRemoteToken, OutboundRateLimiterConfig: disabledLimiter, InboundRateLimiterConfig: disabledLimiter,
			},
			{
				RemoteChainSelector: remoteChain2, RemotePoolAddresses: [][]byte{dummyRemotePool},
				RemoteTokenAddress: dummyRemoteToken, OutboundRateLimiterConfig: disabledLimiter, InboundRateLimiterConfig: disabledLimiter,
			},
		},
	)
	require.NoError(t, err)
	_, err = chain.Confirm(tx)
	require.NoError(t, err)

	// Mark chains as siloed and set silo rebalancers via updateSiloDesignations
	type siloConfigUpdate struct {
		RemoteChainSelector uint64
		Rebalancer          common.Address
	}
	tx, err = oldPoolBound.Transact(chain.DeployerKey, "updateSiloDesignations",
		[]uint64{},
		[]siloConfigUpdate{
			{RemoteChainSelector: remoteChain1, Rebalancer: deployer},
			{RemoteChainSelector: remoteChain2, Rebalancer: deployer},
		},
	)
	require.NoError(t, err)
	_, err = chain.Confirm(tx)
	require.NoError(t, err)

	// Set deployer as the unsiloed rebalancer (for provideLiquidity)
	_, err = operations.ExecuteOperation(e.OperationsBundle, old_siloed.SetRebalancer, chain,
		evm_contract.FunctionInput[common.Address]{ChainSelector: chainSel, Address: oldPoolAddr, Args: deployer})
	require.NoError(t, err)

	// Mint tokens to deployer, approve old pool, then provide siloed + unsiloed liquidity
	_, err = operations.ExecuteOperation(e.OperationsBundle, burn_mint_erc20_with_drip.GrantMintAndBurnRoles, chain,
		evm_contract.FunctionInput[common.Address]{ChainSelector: chainSel, Address: tokenAddr, Args: deployer})
	require.NoError(t, err)

	_, err = operations.ExecuteOperation(e.OperationsBundle, burn_mint_erc20_with_drip.Mint, chain,
		evm_contract.FunctionInput[burn_mint_erc20_with_drip.MintArgs]{
			ChainSelector: chainSel, Address: tokenAddr,
			Args: burn_mint_erc20_with_drip.MintArgs{Account: deployer, Amount: totalMint},
		})
	require.NoError(t, err)

	_, err = operations.ExecuteOperation(e.OperationsBundle, erc20.Approve, chain,
		evm_contract.FunctionInput[erc20.ApproveArgs]{
			ChainSelector: chainSel, Address: tokenAddr,
			Args: erc20.ApproveArgs{Spender: oldPoolAddr, Amount: totalMint},
		})
	require.NoError(t, err)

	tx, err = oldPoolBound.Transact(chain.DeployerKey, "provideSiloedLiquidity", remoteChain1, silo1Amount)
	require.NoError(t, err)
	_, err = chain.Confirm(tx)
	require.NoError(t, err)

	tx, err = oldPoolBound.Transact(chain.DeployerKey, "provideSiloedLiquidity", remoteChain2, silo2Amount)
	require.NoError(t, err)
	_, err = chain.Confirm(tx)
	require.NoError(t, err)

	tx, err = oldPoolBound.Transact(chain.DeployerKey, "provideLiquidity", unsiloedAmount)
	require.NoError(t, err)
	_, err = chain.Confirm(tx)
	require.NoError(t, err)

	// Deploy new v2.0 siloed pool via gobindings
	auth := *chain.DeployerKey
	auth.Nonce = nil
	newPoolAddr, tx, newPoolContract, err := latest_siloed.DeploySiloedLockReleaseTokenPool(
		&auth, chain.Client, tokenAddr, 18, common.Address{}, rmnProxyAddr, routerAddr,
	)
	require.NoError(t, err)
	_, err = chain.Confirm(tx)
	require.NoError(t, err)

	// Add chains on new pool
	auth.Nonce = nil
	tx, err = newPoolContract.ApplyChainUpdates(&auth, []uint64{}, []latest_siloed.TokenPoolChainUpdate{
		{
			RemoteChainSelector: remoteChain1, RemotePoolAddresses: [][]byte{dummyRemotePool},
			RemoteTokenAddress:        dummyRemoteToken,
			OutboundRateLimiterConfig: latest_siloed.RateLimiterConfig{IsEnabled: false, Capacity: big.NewInt(0), Rate: big.NewInt(0)},
			InboundRateLimiterConfig:  latest_siloed.RateLimiterConfig{IsEnabled: false, Capacity: big.NewInt(0), Rate: big.NewInt(0)},
		},
		{
			RemoteChainSelector: remoteChain2, RemotePoolAddresses: [][]byte{dummyRemotePool},
			RemoteTokenAddress:        dummyRemoteToken,
			OutboundRateLimiterConfig: latest_siloed.RateLimiterConfig{IsEnabled: false, Capacity: big.NewInt(0), Rate: big.NewInt(0)},
			InboundRateLimiterConfig:  latest_siloed.RateLimiterConfig{IsEnabled: false, Capacity: big.NewInt(0), Rate: big.NewInt(0)},
		},
	})
	require.NoError(t, err)
	_, err = chain.Confirm(tx)
	require.NoError(t, err)

	// Deploy per-chain lockboxes
	lockbox1Report, err := operations.ExecuteOperation(e.OperationsBundle, erc20_lock_box.Deploy, chain,
		evm_contract.DeployInput[erc20_lock_box.ConstructorArgs]{
			ChainSelector:  chainSel,
			TypeAndVersion: deployment.NewTypeAndVersion(erc20_lock_box.ContractType, *erc20_lock_box.Version),
			Args:           erc20_lock_box.ConstructorArgs{Token: tokenAddr},
			Qualifier:      strPtr("chain1"),
		})
	require.NoError(t, err)
	lockbox1Addr := common.HexToAddress(lockbox1Report.Output.Address)

	lockbox2Report, err := operations.ExecuteOperation(e.OperationsBundle, erc20_lock_box.Deploy, chain,
		evm_contract.DeployInput[erc20_lock_box.ConstructorArgs]{
			ChainSelector:  chainSel,
			TypeAndVersion: deployment.NewTypeAndVersion(erc20_lock_box.ContractType, *erc20_lock_box.Version),
			Args:           erc20_lock_box.ConstructorArgs{Token: tokenAddr},
			Qualifier:      strPtr("chain2"),
		})
	require.NoError(t, err)
	lockbox2Addr := common.HexToAddress(lockbox2Report.Output.Address)

	// Configure lockboxes on new pool
	auth.Nonce = nil
	tx, err = newPoolContract.ConfigureLockBoxes(&auth, []latest_siloed.SiloedLockReleaseTokenPoolLockBoxConfig{
		{RemoteChainSelector: remoteChain1, LockBox: lockbox1Addr},
		{RemoteChainSelector: remoteChain2, LockBox: lockbox2Addr},
	})
	require.NoError(t, err)
	_, err = chain.Confirm(tx)
	require.NoError(t, err)

	// Run migration (100% of all liquidity)
	basisPoints := uint16(10000)
	_, err = operations.ExecuteSequence(
		testsetup.BundleWithFreshReporter(e.OperationsBundle),
		tokens.MigrateLockReleasePoolLiquidity,
		e.BlockChains,
		tokens_core.MigrateLockReleasePoolLiquidityInput{
			ChainSelector:   chainSel,
			OldPoolAddress:  oldPoolAddr.Hex(),
			NewPoolAddress:  newPoolAddr.Hex(),
			TimelockAddress: deployer.Hex(),
			BasisPoints:     &basisPoints,
		},
	)
	require.NoError(t, err)

	// Verify old pool is drained
	oldPoolBal, err := operations.ExecuteOperation(
		testsetup.BundleWithFreshReporter(e.OperationsBundle), erc20.BalanceOf, chain,
		evm_contract.FunctionInput[common.Address]{ChainSelector: chainSel, Address: tokenAddr, Args: oldPoolAddr})
	require.NoError(t, err)
	require.Equal(t, 0, big.NewInt(0).Cmp(oldPoolBal.Output), "Old siloed pool should be fully drained")

	// Verify lockbox1 received silo1 amount
	lb1Bal, err := operations.ExecuteOperation(
		testsetup.BundleWithFreshReporter(e.OperationsBundle), erc20.BalanceOf, chain,
		evm_contract.FunctionInput[common.Address]{ChainSelector: chainSel, Address: tokenAddr, Args: lockbox1Addr})
	require.NoError(t, err)
	require.True(t, lb1Bal.Output.Sign() > 0, "Lockbox 1 should have received tokens")

	// Verify lockbox2 received silo2 amount
	lb2Bal, err := operations.ExecuteOperation(
		testsetup.BundleWithFreshReporter(e.OperationsBundle), erc20.BalanceOf, chain,
		evm_contract.FunctionInput[common.Address]{ChainSelector: chainSel, Address: tokenAddr, Args: lockbox2Addr})
	require.NoError(t, err)
	require.True(t, lb2Bal.Output.Sign() > 0, "Lockbox 2 should have received tokens")

	// Total across both lockboxes should equal total minted
	totalInLockboxes := new(big.Int).Add(lb1Bal.Output, lb2Bal.Output)
	require.Equal(t, 0, totalMint.Cmp(totalInLockboxes),
		"Total lockbox balances should equal total original liquidity")
}

func TestMigrateLockReleasePoolLiquidity_WithSetPoolConfig(t *testing.T) {
	chainSel := uint64(5009297550715157269)
	totalLiquidity := big.NewInt(5000)
	s := setupMigrationTest(t, chainSel, totalLiquidity)
	chain := s.env.BlockChains.EVMChains()[chainSel]

	// Deploy TokenAdminRegistry
	tarReport, err := operations.ExecuteOperation(
		s.env.OperationsBundle, tar.Deploy, chain,
		evm_contract.DeployInput[tar.ConstructorArgs]{
			ChainSelector:  chainSel,
			TypeAndVersion: deployment.NewTypeAndVersion(tar.ContractType, *tar.Version),
			Args:           tar.ConstructorArgs{},
		})
	require.NoError(t, err)
	tarAddr := common.HexToAddress(tarReport.Output.Address)

	// Register deployer as administrator for the token
	_, err = operations.ExecuteOperation(
		testsetup.BundleWithFreshReporter(s.env.OperationsBundle), tar.ProposeAdministrator, chain,
		evm_contract.FunctionInput[tar.ProposeAdministratorArgs]{
			ChainSelector: chainSel, Address: tarAddr,
			Args: tar.ProposeAdministratorArgs{TokenAddress: s.tokenAddr, Administrator: s.deployer},
		})
	require.NoError(t, err)

	_, err = operations.ExecuteOperation(
		testsetup.BundleWithFreshReporter(s.env.OperationsBundle), tar.AcceptAdminRole, chain,
		evm_contract.FunctionInput[tar.AcceptAdminRoleArgs]{
			ChainSelector: chainSel, Address: tarAddr,
			Args: tar.AcceptAdminRoleArgs{TokenAddress: s.tokenAddr},
		})
	require.NoError(t, err)

	// Run migration with SetPoolConfig
	basisPoints := uint16(10000)
	_, err = operations.ExecuteSequence(
		testsetup.BundleWithFreshReporter(s.env.OperationsBundle),
		tokens.MigrateLockReleasePoolLiquidity,
		s.env.BlockChains,
		tokens_core.MigrateLockReleasePoolLiquidityInput{
			ChainSelector:   chainSel,
			OldPoolAddress:  s.oldPoolAddr.Hex(),
			NewPoolAddress:  s.newPoolAddr.Hex(),
			TimelockAddress: s.deployer.Hex(),
			BasisPoints:     &basisPoints,
			SetPoolConfig: &tokens_core.MigrationSetPoolConfig{
				RegistryAddress: tarAddr.Hex(),
				TokenAddress:    s.tokenAddr.Hex(),
			},
		},
	)
	require.NoError(t, err)

	// Verify the pool was set on the registry
	tokenConfig, err := operations.ExecuteOperation(
		testsetup.BundleWithFreshReporter(s.env.OperationsBundle), tar.GetTokenConfig, chain,
		evm_contract.FunctionInput[common.Address]{
			ChainSelector: chainSel, Address: tarAddr, Args: s.tokenAddr,
		})
	require.NoError(t, err)
	require.Equal(t, s.newPoolAddr, tokenConfig.Output.TokenPool,
		"TokenAdminRegistry should point to the new pool after migration")

	// Verify liquidity was also migrated
	lockboxBal, err := operations.ExecuteOperation(
		testsetup.BundleWithFreshReporter(s.env.OperationsBundle), erc20.BalanceOf, chain,
		evm_contract.FunctionInput[common.Address]{ChainSelector: chainSel, Address: s.tokenAddr, Args: s.lockBoxAddr})
	require.NoError(t, err)
	require.Equal(t, 0, totalLiquidity.Cmp(lockboxBal.Output), "Lockbox should hold all liquidity")
}

func TestMigrateLockReleasePoolLiquidity_AuthorizedCallerCleanup(t *testing.T) {
	chainSel := uint64(5009297550715157269)
	totalLiquidity := big.NewInt(5000)
	s := setupMigrationTest(t, chainSel, totalLiquidity)
	chain := s.env.BlockChains.EVMChains()[chainSel]

	// Add a pre-existing authorized caller to the lockbox before migration
	preExistingCaller := common.HexToAddress("0x1234567890123456789012345678901234567890")
	_, err := operations.ExecuteOperation(
		s.env.OperationsBundle, erc20_lock_box.ApplyAuthorizedCallerUpdates, chain,
		evm_contract.FunctionInput[erc20_lock_box.AuthorizedCallerArgs]{
			ChainSelector: chainSel, Address: s.lockBoxAddr,
			Args: erc20_lock_box.AuthorizedCallerArgs{
				AddedCallers:   []common.Address{preExistingCaller},
				RemovedCallers: []common.Address{},
			},
		})
	require.NoError(t, err)

	// Verify pre-existing caller is present before migration
	preCallersReport, err := operations.ExecuteOperation(
		testsetup.BundleWithFreshReporter(s.env.OperationsBundle), erc20_lock_box.GetAllAuthorizedCallers, chain,
		evm_contract.FunctionInput[struct{}]{ChainSelector: chainSel, Address: s.lockBoxAddr})
	require.NoError(t, err)
	require.Contains(t, preCallersReport.Output, preExistingCaller,
		"Pre-existing caller should be present before migration")

	// Run migration
	basisPoints := uint16(10000)
	_, err = operations.ExecuteSequence(
		testsetup.BundleWithFreshReporter(s.env.OperationsBundle),
		tokens.MigrateLockReleasePoolLiquidity,
		s.env.BlockChains,
		tokens_core.MigrateLockReleasePoolLiquidityInput{
			ChainSelector:   chainSel,
			OldPoolAddress:  s.oldPoolAddr.Hex(),
			NewPoolAddress:  s.newPoolAddr.Hex(),
			TimelockAddress: s.deployer.Hex(),
			BasisPoints:     &basisPoints,
		},
	)
	require.NoError(t, err)

	// Verify pre-existing caller is still present after migration
	postCallersReport, err := operations.ExecuteOperation(
		testsetup.BundleWithFreshReporter(s.env.OperationsBundle), erc20_lock_box.GetAllAuthorizedCallers, chain,
		evm_contract.FunctionInput[struct{}]{ChainSelector: chainSel, Address: s.lockBoxAddr})
	require.NoError(t, err)
	require.Contains(t, postCallersReport.Output, preExistingCaller,
		"Pre-existing authorized caller should be preserved after migration")

	// Verify timelock (deployer) was removed from authorized callers
	require.NotContains(t, postCallersReport.Output, s.deployer,
		"Timelock should be removed from lockbox authorized callers after migration")
}

func TestMigrateLockReleasePoolLiquidity_MultiplePartialMigrations(t *testing.T) {
	chainSel := uint64(5009297550715157269)
	totalLiquidity := big.NewInt(10000)
	s := setupMigrationTest(t, chainSel, totalLiquidity)
	chain := s.env.BlockChains.EVMChains()[chainSel]

	// Step 1: Migrate 50%
	basisPoints := uint16(5000)
	_, err := operations.ExecuteSequence(
		testsetup.BundleWithFreshReporter(s.env.OperationsBundle),
		tokens.MigrateLockReleasePoolLiquidity,
		s.env.BlockChains,
		tokens_core.MigrateLockReleasePoolLiquidityInput{
			ChainSelector:   chainSel,
			OldPoolAddress:  s.oldPoolAddr.Hex(),
			NewPoolAddress:  s.newPoolAddr.Hex(),
			TimelockAddress: s.deployer.Hex(),
			BasisPoints:     &basisPoints,
		},
	)
	require.NoError(t, err)

	// Verify 50% migrated
	oldPoolBal1, err := operations.ExecuteOperation(
		testsetup.BundleWithFreshReporter(s.env.OperationsBundle), erc20.BalanceOf, chain,
		evm_contract.FunctionInput[common.Address]{ChainSelector: chainSel, Address: s.tokenAddr, Args: s.oldPoolAddr})
	require.NoError(t, err)
	require.Equal(t, 0, big.NewInt(5000).Cmp(oldPoolBal1.Output), "Old pool should retain 50%% after first migration")

	lockboxBal1, err := operations.ExecuteOperation(
		testsetup.BundleWithFreshReporter(s.env.OperationsBundle), erc20.BalanceOf, chain,
		evm_contract.FunctionInput[common.Address]{ChainSelector: chainSel, Address: s.tokenAddr, Args: s.lockBoxAddr})
	require.NoError(t, err)
	require.Equal(t, 0, big.NewInt(5000).Cmp(lockboxBal1.Output), "Lockbox should hold 50%% after first migration")

	// Step 2: Migrate 100% of remaining
	basisPoints = uint16(10000)
	_, err = operations.ExecuteSequence(
		testsetup.BundleWithFreshReporter(s.env.OperationsBundle),
		tokens.MigrateLockReleasePoolLiquidity,
		s.env.BlockChains,
		tokens_core.MigrateLockReleasePoolLiquidityInput{
			ChainSelector:   chainSel,
			OldPoolAddress:  s.oldPoolAddr.Hex(),
			NewPoolAddress:  s.newPoolAddr.Hex(),
			TimelockAddress: s.deployer.Hex(),
			BasisPoints:     &basisPoints,
		},
	)
	require.NoError(t, err)

	// Verify all liquidity migrated
	oldPoolBal2, err := operations.ExecuteOperation(
		testsetup.BundleWithFreshReporter(s.env.OperationsBundle), erc20.BalanceOf, chain,
		evm_contract.FunctionInput[common.Address]{ChainSelector: chainSel, Address: s.tokenAddr, Args: s.oldPoolAddr})
	require.NoError(t, err)
	require.Equal(t, 0, big.NewInt(0).Cmp(oldPoolBal2.Output), "Old pool should be fully drained after second migration")

	lockboxBal2, err := operations.ExecuteOperation(
		testsetup.BundleWithFreshReporter(s.env.OperationsBundle), erc20.BalanceOf, chain,
		evm_contract.FunctionInput[common.Address]{ChainSelector: chainSel, Address: s.tokenAddr, Args: s.lockBoxAddr})
	require.NoError(t, err)
	require.Equal(t, 0, totalLiquidity.Cmp(lockboxBal2.Output), "Lockbox should hold all liquidity after second migration")

	// Verify rebalancer is restored (should be zero address since none was set)
	rebalancerReport, err := operations.ExecuteOperation(
		testsetup.BundleWithFreshReporter(s.env.OperationsBundle), old_lrtp.GetRebalancer, chain,
		evm_contract.FunctionInput[struct{}]{ChainSelector: chainSel, Address: s.oldPoolAddr})
	require.NoError(t, err)
	require.Equal(t, common.Address{}, rebalancerReport.Output,
		"Rebalancer should be restored after both migrations")
}

func strPtr(s string) *string {
	return &s
}

func uint16Ptr(v uint16) *uint16 {
	return &v
}
