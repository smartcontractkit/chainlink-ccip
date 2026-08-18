package tokens_test

import (
	"fmt"
	"math/big"
	"strings"
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/rmn_proxy"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/router"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/burn_mint_erc20_with_drip"
	tar "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/token_admin_registry"
	old_lrtp "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_1/operations/lock_release_token_pool"
	old_siloed "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_1/operations/siloed_lock_release_token_pool"

	"github.com/smartcontractkit/chainlink-deployments-framework/chain"
	evm_contract "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations/contract"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/erc20"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/create2_factory"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/erc20_lock_box"
	new_lrtp "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/lock_release_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/sequences"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/sequences/tokens"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/testsetup"
	latest_siloed "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v2_0_0/siloed_lock_release_token_pool"
	tokens_core "github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	seqtypes "github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
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
				BasisPoints:     new(uint16(5000)),
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
				BasisPoints:     new(uint16(0)),
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
				BasisPoints:     new(uint16(10001)),
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
				BasisPoints:     new(uint16(10000)),
			},
			expectedErr: "OldPoolAddress and NewPoolAddress must be provided",
		},
		{
			name: "missing TimelockAddress",
			input: tokens_core.MigrateLockReleasePoolLiquidityInput{
				ChainSelector:  chainSel,
				OldPoolAddress: "0x0000000000000000000000000000000000000001",
				NewPoolAddress: "0x0000000000000000000000000000000000000002",
				BasisPoints:    new(uint16(10000)),
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
			ChainSelector:     chainSel,
			ContractParams:    testsetup.CreateBasicContractParams(),
			CREATE2Factory:    common.HexToAddress(create2FactoryRef.Address),
			DeployerKeyOwned:  true,
			ExistingAddresses: testsetup.UltraFastCurseMCMSRefs(chainSel),
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

// executeMigrationSequence runs the liquidity migration sequence and sends any prepared
// batch transactions using the chain deployer key (standing in for the timelock in tests).
func executeMigrationSequence(
	t *testing.T,
	bundle operations.Bundle,
	blockChains chain.BlockChains,
	input tokens_core.MigrateLockReleasePoolLiquidityInput,
) seqtypes.OnChainOutput {
	t.Helper()

	report, err := operations.ExecuteSequence(
		bundle,
		tokens.MigrateLockReleasePoolLiquidity,
		blockChains,
		input,
	)
	require.NoError(t, err)

	evmChain, ok := blockChains.EVMChains()[input.ChainSelector]
	require.True(t, ok, "chain %d not found", input.ChainSelector)

	for _, batch := range report.Output.BatchOps {
		for _, mcmsTx := range batch.Transactions {
			to := common.HexToAddress(mcmsTx.To)
			nonce, err := evmChain.Client.PendingNonceAt(t.Context(), evmChain.DeployerKey.From)
			require.NoError(t, err)

			gasLimit, err := evmChain.Client.EstimateGas(t.Context(), ethereum.CallMsg{
				From: evmChain.DeployerKey.From,
				To:   &to,
				Data: mcmsTx.Data,
			})
			require.NoError(t, err)

			gasPrice, err := evmChain.Client.SuggestGasPrice(t.Context())
			require.NoError(t, err)

			tx := types.NewTransaction(nonce, to, big.NewInt(0), gasLimit, gasPrice, mcmsTx.Data)
			signedTx, err := evmChain.DeployerKey.Signer(evmChain.DeployerKey.From, tx)
			require.NoError(t, err)

			err = evmChain.Client.SendTransaction(t.Context(), signedTx)
			require.NoError(t, err)

			_, err = evmChain.Confirm(signedTx)
			require.NoError(t, err)
		}
	}

	return report.Output
}

func TestMigrateLockReleasePoolLiquidity_UnsiloedPartialBasisPoints(t *testing.T) {
	chainSel := uint64(5009297550715157269)
	totalLiquidity := big.NewInt(10000)
	s := setupMigrationTest(t, chainSel, totalLiquidity)
	chain := s.env.BlockChains.EVMChains()[chainSel]

	basisPoints := uint16(8000) // 80%
	executeMigrationSequence(t,
		s.env.OperationsBundle,
		s.env.BlockChains,
		tokens_core.MigrateLockReleasePoolLiquidityInput{
			ChainSelector:   chainSel,
			OldPoolAddress:  s.oldPoolAddr.Hex(),
			NewPoolAddress:  s.newPoolAddr.Hex(),
			TimelockAddress: s.deployer.Hex(),
			BasisPoints:     &basisPoints,
		},
	)

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

	// Verify the timelock was added to the lockbox authorized callers (deposit flow leaves it in place)
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
	require.Contains(t, authCallersReport.Output, s.deployer,
		"Timelock should be in lockbox authorized callers (deposit flow)")
}

func TestMigrateLockReleasePoolLiquidity_UnsiloedFullBasisPoints(t *testing.T) {
	chainSel := uint64(5009297550715157269)
	totalLiquidity := big.NewInt(5000)
	s := setupMigrationTest(t, chainSel, totalLiquidity)
	chain := s.env.BlockChains.EVMChains()[chainSel]

	basisPoints := uint16(10000) // 100%
	executeMigrationSequence(t,
		s.env.OperationsBundle,
		s.env.BlockChains,
		tokens_core.MigrateLockReleasePoolLiquidityInput{
			ChainSelector:   chainSel,
			OldPoolAddress:  s.oldPoolAddr.Hex(),
			NewPoolAddress:  s.newPoolAddr.Hex(),
			TimelockAddress: s.deployer.Hex(),
			BasisPoints:     &basisPoints,
		},
	)

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
	executeMigrationSequence(t,
		s.env.OperationsBundle,
		s.env.BlockChains,
		tokens_core.MigrateLockReleasePoolLiquidityInput{
			ChainSelector:   chainSel,
			OldPoolAddress:  s.oldPoolAddr.Hex(),
			NewPoolAddress:  s.newPoolAddr.Hex(),
			TimelockAddress: s.deployer.Hex(),
			Amount:          exactAmount,
		},
	)

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
	executeMigrationSequence(t,
		freshBundle,
		s.env.BlockChains,
		tokens_core.MigrateLockReleasePoolLiquidityInput{
			ChainSelector:   chainSel,
			OldPoolAddress:  s.oldPoolAddr.Hex(),
			NewPoolAddress:  s.newPoolAddr.Hex(),
			TimelockAddress: s.deployer.Hex(),
			BasisPoints:     &basisPoints,
		},
	)

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

// siloedChainSpec is one siloed remote chain and the liquidity reserved for it.
type siloedChainSpec struct {
	selector uint64
	amount   *big.Int
}

// siloedTopology describes the liquidity layout of a legacy siloed pool and the lockbox layout of
// the v2 pool it migrates into. Shared chains model the unsiloed case: they are not siloed on the
// old pool, and on the new pool they all map to one lockbox that no siloed chain uses.
type siloedTopology struct {
	siloedChains   []siloedChainSpec
	sharedChains   []uint64
	unsiloedAmount *big.Int
}

type siloedMigrationSetup struct {
	env           *deployment.Environment
	chainSel      uint64
	deployer      common.Address
	tokenAddr     common.Address
	oldPoolAddr   common.Address
	newPoolAddr   common.Address
	siloLockBoxes map[uint64]common.Address
	sharedLockBox common.Address // zero when the topology declares no shared chains
	totalMint     *big.Int
}

// balanceOf reads a token balance through a fresh reporter, matching how the assertions in these
// tests read state after a migration has been executed.
func (s siloedMigrationSetup) balanceOf(t *testing.T, holder common.Address) *big.Int {
	t.Helper()

	chain := s.env.BlockChains.EVMChains()[s.chainSel]
	report, err := operations.ExecuteOperation(
		testsetup.BundleWithFreshReporter(s.env.OperationsBundle), erc20.BalanceOf, chain,
		evm_contract.FunctionInput[common.Address]{ChainSelector: s.chainSel, Address: s.tokenAddr, Args: holder})
	require.NoError(t, err)

	return report.Output
}

// setupSiloedMigrationTest builds a legacy v1.6.1 siloed pool holding the topology's liquidity and
// a v2.0 siloed pool with one lockbox per siloed chain plus, when the topology declares shared
// chains, one further lockbox mapped to all of them.
func setupSiloedMigrationTest(t *testing.T, chainSel uint64, topology siloedTopology) siloedMigrationSetup {
	t.Helper()

	totalMint := new(big.Int).Set(topology.unsiloedAmount)
	for _, silo := range topology.siloedChains {
		totalMint.Add(totalMint, silo.amount)
	}

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
			ChainSelector:     chainSel,
			ContractParams:    testsetup.CreateBasicContractParams(),
			CREATE2Factory:    common.HexToAddress(create2FactoryRef.Address),
			DeployerKeyOwned:  true,
			ExistingAddresses: testsetup.UltraFastCurseMCMSRefs(chainSel),
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

	allChains := make([]uint64, 0, len(topology.siloedChains)+len(topology.sharedChains))
	for _, silo := range topology.siloedChains {
		allChains = append(allChains, silo.selector)
	}
	allChains = append(allChains, topology.sharedChains...)

	parsed, err := abi.JSON(strings.NewReader(old_siloed.SiloedLockReleaseTokenPoolABI))
	require.NoError(t, err)
	oldPoolBound := bind.NewBoundContract(oldPoolAddr, parsed, chain.Client, chain.Client, chain.Client)

	disabledLimiter := testRateLimiterConfig{IsEnabled: false, Capacity: big.NewInt(0), Rate: big.NewInt(0)}
	dummyRemotePool := common.HexToAddress("0xdeadbeefdeadbeefdeadbeefdeadbeefdeadbeef").Bytes()
	dummyRemoteToken := common.HexToAddress("0xcafecafecafecafecafecafecafecafecafecafe").Bytes()

	oldChainUpdates := make([]testChainUpdate, 0, len(allChains))
	for _, sel := range allChains {
		oldChainUpdates = append(oldChainUpdates, testChainUpdate{
			RemoteChainSelector: sel, RemotePoolAddresses: [][]byte{dummyRemotePool},
			RemoteTokenAddress: dummyRemoteToken, OutboundRateLimiterConfig: disabledLimiter, InboundRateLimiterConfig: disabledLimiter,
		})
	}
	tx, err := oldPoolBound.Transact(chain.DeployerKey, "applyChainUpdates", []uint64{}, oldChainUpdates)
	require.NoError(t, err)
	_, err = chain.Confirm(tx)
	require.NoError(t, err)

	// Only the siloed chains get a silo designation; the shared chains stay on the pool's unsiloed
	// balance, which is the case the migration has to be told where to send.
	siloAdds := make([]old_siloed.SiloConfigUpdate, 0, len(topology.siloedChains))
	for _, silo := range topology.siloedChains {
		siloAdds = append(siloAdds, old_siloed.SiloConfigUpdate{RemoteChainSelector: silo.selector, Rebalancer: deployer})
	}
	_, err = operations.ExecuteOperation(e.OperationsBundle, old_siloed.UpdateSiloDesignations, chain,
		evm_contract.FunctionInput[old_siloed.UpdateSiloDesignationsArgs]{
			ChainSelector: chainSel, Address: oldPoolAddr,
			Args: old_siloed.UpdateSiloDesignationsArgs{Removes: []uint64{}, Adds: siloAdds},
		})
	require.NoError(t, err)

	_, err = operations.ExecuteOperation(e.OperationsBundle, old_siloed.SetRebalancer, chain,
		evm_contract.FunctionInput[common.Address]{ChainSelector: chainSel, Address: oldPoolAddr, Args: deployer})
	require.NoError(t, err)

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
			Args: erc20.ApproveArgs{Spender: oldPoolAddr, Value: totalMint},
		})
	require.NoError(t, err)

	for _, silo := range topology.siloedChains {
		_, err = operations.ExecuteOperation(e.OperationsBundle, old_siloed.ProvideSiloedLiquidity, chain,
			evm_contract.FunctionInput[old_siloed.ProvideSiloedLiquidityArgs]{
				ChainSelector: chainSel, Address: oldPoolAddr,
				Args: old_siloed.ProvideSiloedLiquidityArgs{RemoteChainSelector: silo.selector, Amount: silo.amount},
			})
		require.NoError(t, err)
	}

	if topology.unsiloedAmount.Sign() > 0 {
		_, err = operations.ExecuteOperation(e.OperationsBundle, old_siloed.ProvideLiquidity, chain,
			evm_contract.FunctionInput[*big.Int]{ChainSelector: chainSel, Address: oldPoolAddr, Args: topology.unsiloedAmount})
		require.NoError(t, err)
	}

	auth := *chain.DeployerKey
	auth.Nonce = nil
	newPoolAddr, tx, newPoolContract, err := latest_siloed.DeploySiloedLockReleaseTokenPool(
		&auth, chain.Client, tokenAddr, 18, common.Address{}, rmnProxyAddr, routerAddr,
	)
	require.NoError(t, err)
	_, err = chain.Confirm(tx)
	require.NoError(t, err)

	newChainUpdates := make([]latest_siloed.TokenPoolChainUpdate, 0, len(allChains))
	for _, sel := range allChains {
		newChainUpdates = append(newChainUpdates, latest_siloed.TokenPoolChainUpdate{
			RemoteChainSelector: sel, RemotePoolAddresses: [][]byte{dummyRemotePool},
			RemoteTokenAddress:        dummyRemoteToken,
			OutboundRateLimiterConfig: latest_siloed.RateLimiterConfig{IsEnabled: false, Capacity: big.NewInt(0), Rate: big.NewInt(0)},
			InboundRateLimiterConfig:  latest_siloed.RateLimiterConfig{IsEnabled: false, Capacity: big.NewInt(0), Rate: big.NewInt(0)},
		})
	}
	auth.Nonce = nil
	tx, err = newPoolContract.ApplyChainUpdates(&auth, []uint64{}, newChainUpdates)
	require.NoError(t, err)
	_, err = chain.Confirm(tx)
	require.NoError(t, err)

	deployLockBox := func(qualifier string) common.Address {
		q := qualifier
		report, deployErr := operations.ExecuteOperation(e.OperationsBundle, erc20_lock_box.Deploy, chain,
			evm_contract.DeployInput[erc20_lock_box.ConstructorArgs]{
				ChainSelector:  chainSel,
				TypeAndVersion: deployment.NewTypeAndVersion(erc20_lock_box.ContractType, *erc20_lock_box.Version),
				Args:           erc20_lock_box.ConstructorArgs{Token: tokenAddr},
				Qualifier:      &q,
			})
		require.NoError(t, deployErr)

		return common.HexToAddress(report.Output.Address)
	}

	lockBoxConfigs := make([]latest_siloed.SiloedLockReleaseTokenPoolLockBoxConfig, 0, len(allChains))
	siloLockBoxes := make(map[uint64]common.Address, len(topology.siloedChains))
	for i, silo := range topology.siloedChains {
		lockBox := deployLockBox(fmt.Sprintf("silo-%d", i))
		siloLockBoxes[silo.selector] = lockBox
		lockBoxConfigs = append(lockBoxConfigs, latest_siloed.SiloedLockReleaseTokenPoolLockBoxConfig{
			RemoteChainSelector: silo.selector, LockBox: lockBox,
		})
	}

	var sharedLockBox common.Address
	if len(topology.sharedChains) > 0 {
		sharedLockBox = deployLockBox("shared")
		for _, sel := range topology.sharedChains {
			lockBoxConfigs = append(lockBoxConfigs, latest_siloed.SiloedLockReleaseTokenPoolLockBoxConfig{
				RemoteChainSelector: sel, LockBox: sharedLockBox,
			})
		}
	}

	auth.Nonce = nil
	tx, err = newPoolContract.ConfigureLockBoxes(&auth, lockBoxConfigs)
	require.NoError(t, err)
	_, err = chain.Confirm(tx)
	require.NoError(t, err)

	return siloedMigrationSetup{
		env:           e,
		chainSel:      chainSel,
		deployer:      deployer,
		tokenAddr:     tokenAddr,
		oldPoolAddr:   oldPoolAddr,
		newPoolAddr:   newPoolAddr,
		siloLockBoxes: siloLockBoxes,
		sharedLockBox: sharedLockBox,
		totalMint:     totalMint,
	}
}

// TestMigrateLockReleasePoolLiquidity_SiloedSharedLockBoxIsDistinct covers the topology the
// unsiloed destination exists for: the shared chains map to a lockbox that no siloed chain uses.
// Before the destination was explicit, the shared balance went to whichever siloed chain came
// first, leaving the shared chains unable to release.
func TestMigrateLockReleasePoolLiquidity_SiloedSharedLockBoxIsDistinct(t *testing.T) {
	chainSel := uint64(5009297550715157269)
	siloed1 := uint64(3379446385462418246)
	siloed2 := uint64(4949039107694359620)
	shared1 := uint64(15971525489660198786)
	shared2 := uint64(3734403246176062136)

	silo1Amount := big.NewInt(3000)
	silo2Amount := big.NewInt(2000)
	unsiloedAmount := big.NewInt(1000)

	s := setupSiloedMigrationTest(t, chainSel, siloedTopology{
		siloedChains: []siloedChainSpec{
			{selector: siloed1, amount: silo1Amount},
			{selector: siloed2, amount: silo2Amount},
		},
		sharedChains:   []uint64{shared1, shared2},
		unsiloedAmount: unsiloedAmount,
	})

	basisPoints := uint16(10000)
	executeMigrationSequence(t,
		testsetup.BundleWithFreshReporter(s.env.OperationsBundle),
		s.env.BlockChains,
		tokens_core.MigrateLockReleasePoolLiquidityInput{
			ChainSelector:          chainSel,
			OldPoolAddress:         s.oldPoolAddr.Hex(),
			NewPoolAddress:         s.newPoolAddr.Hex(),
			TimelockAddress:        s.deployer.Hex(),
			BasisPoints:            &basisPoints,
			UnsiloedLockBoxAddress: s.sharedLockBox.Hex(),
		},
	)

	require.Equal(t, 0, big.NewInt(0).Cmp(s.balanceOf(t, s.oldPoolAddr)), "old pool should be fully drained")

	// Each silo keeps exactly its own reserved balance - the shared balance must not land here.
	require.Equal(t, 0, silo1Amount.Cmp(s.balanceOf(t, s.siloLockBoxes[siloed1])),
		"siloed chain 1 lockbox should hold only its own reserved balance")
	require.Equal(t, 0, silo2Amount.Cmp(s.balanceOf(t, s.siloLockBoxes[siloed2])),
		"siloed chain 2 lockbox should hold only its own reserved balance")

	require.Equal(t, 0, unsiloedAmount.Cmp(s.balanceOf(t, s.sharedLockBox)),
		"shared lockbox should hold the whole unsiloed balance")
}

// TestMigrateLockReleasePoolLiquidity_SiloedUnsiloedDestinationRequired asserts the migration
// refuses to guess: unsiloed liquidity with no destination is an error, not a fallback.
func TestMigrateLockReleasePoolLiquidity_SiloedUnsiloedDestinationRequired(t *testing.T) {
	chainSel := uint64(5009297550715157269)
	s := setupSiloedMigrationTest(t, chainSel, siloedTopology{
		siloedChains:   []siloedChainSpec{{selector: uint64(3379446385462418246), amount: big.NewInt(3000)}},
		sharedChains:   []uint64{uint64(15971525489660198786)},
		unsiloedAmount: big.NewInt(1000),
	})

	// A distinct timelock so the rebalancer assertion below is meaningful - if the handover had
	// been emitted, the old pool's rebalancer would no longer be the deployer.
	timelock := common.HexToAddress("0x00000000000000000000000000000000000000A1")

	basisPoints := uint16(10000)
	_, err := operations.ExecuteSequence(
		testsetup.BundleWithFreshReporter(s.env.OperationsBundle),
		tokens.MigrateLockReleasePoolLiquidity,
		s.env.BlockChains,
		tokens_core.MigrateLockReleasePoolLiquidityInput{
			ChainSelector:   chainSel,
			OldPoolAddress:  s.oldPoolAddr.Hex(),
			NewPoolAddress:  s.newPoolAddr.Hex(),
			TimelockAddress: timelock.Hex(),
			BasisPoints:     &basisPoints,
		},
	)
	require.Error(t, err)
	require.Contains(t, err.Error(), "UnsiloedLockBoxAddress was not set")

	// The failure must land before anything is touched, so the operator can fix the input and
	// re-run rather than unwinding a half-migrated pool.
	require.Equal(t, 0, s.totalMint.Cmp(s.balanceOf(t, s.oldPoolAddr)),
		"old pool must still hold its full balance after a rejected migration")

	chain := s.env.BlockChains.EVMChains()[chainSel]
	rebalancer, err := operations.ExecuteOperation(
		testsetup.BundleWithFreshReporter(s.env.OperationsBundle), old_siloed.GetRebalancer, chain,
		evm_contract.FunctionInput[struct{}]{ChainSelector: chainSel, Address: s.oldPoolAddr})
	require.NoError(t, err)
	require.Equal(t, s.deployer, rebalancer.Output,
		"old pool's rebalancer must not have been repointed at the timelock")
}

// TestMigrateLockReleasePoolLiquidity_SiloedUnsiloedDestinationNotConfigured guards against a
// typo or a lockbox belonging to a different pool being accepted.
func TestMigrateLockReleasePoolLiquidity_SiloedUnsiloedDestinationNotConfigured(t *testing.T) {
	chainSel := uint64(5009297550715157269)
	s := setupSiloedMigrationTest(t, chainSel, siloedTopology{
		siloedChains:   []siloedChainSpec{{selector: uint64(3379446385462418246), amount: big.NewInt(3000)}},
		sharedChains:   []uint64{uint64(15971525489660198786)},
		unsiloedAmount: big.NewInt(1000),
	})

	basisPoints := uint16(10000)
	_, err := operations.ExecuteSequence(
		testsetup.BundleWithFreshReporter(s.env.OperationsBundle),
		tokens.MigrateLockReleasePoolLiquidity,
		s.env.BlockChains,
		tokens_core.MigrateLockReleasePoolLiquidityInput{
			ChainSelector:          chainSel,
			OldPoolAddress:         s.oldPoolAddr.Hex(),
			NewPoolAddress:         s.newPoolAddr.Hex(),
			TimelockAddress:        s.deployer.Hex(),
			BasisPoints:            &basisPoints,
			UnsiloedLockBoxAddress: "0x000000000000000000000000000000000000dEaD",
		},
	)
	require.Error(t, err)
	require.Contains(t, err.Error(), "is not one of the lockboxes configured on new pool")
}

// TestMigrateLockReleasePoolLiquidity_SiloedOutputIsDeterministic runs the same topology in two
// independent environments and requires the emitted transfers to land in the same order and on the
// same roles. The destination used to be picked by Go map iteration, so repeated runs of identical
// input could route the shared balance differently and reviewers could not reproduce a proposal.
//
// Addresses differ between environments, so each transfer receiver is mapped back to its role
// before comparison.
func TestMigrateLockReleasePoolLiquidity_SiloedOutputIsDeterministic(t *testing.T) {
	chainSel := uint64(5009297550715157269)
	siloed1 := uint64(3379446385462418246)
	siloed2 := uint64(4949039107694359620)
	topology := siloedTopology{
		siloedChains: []siloedChainSpec{
			{selector: siloed1, amount: big.NewInt(3000)},
			{selector: siloed2, amount: big.NewInt(2000)},
		},
		sharedChains:   []uint64{uint64(15971525489660198786), uint64(3734403246176062136)},
		unsiloedAmount: big.NewInt(1000),
	}

	// transferRoutes returns, in emission order, "<role>:<amount>" for every ERC20 transfer the
	// migration proposes.
	transferRoutes := func() []string {
		s := setupSiloedMigrationTest(t, chainSel, topology)

		basisPoints := uint16(10000)
		report, err := operations.ExecuteSequence(
			testsetup.BundleWithFreshReporter(s.env.OperationsBundle),
			tokens.MigrateLockReleasePoolLiquidity,
			s.env.BlockChains,
			tokens_core.MigrateLockReleasePoolLiquidityInput{
				ChainSelector:          chainSel,
				OldPoolAddress:         s.oldPoolAddr.Hex(),
				NewPoolAddress:         s.newPoolAddr.Hex(),
				TimelockAddress:        s.deployer.Hex(),
				BasisPoints:            &basisPoints,
				UnsiloedLockBoxAddress: s.sharedLockBox.Hex(),
			},
		)
		require.NoError(t, err)

		roleOf := map[common.Address]string{
			s.siloLockBoxes[siloed1]: "silo1",
			s.siloLockBoxes[siloed2]: "silo2",
			s.sharedLockBox:          "shared",
		}

		var routes []string
		for _, batch := range report.Output.BatchOps {
			for _, mcmsTx := range batch.Transactions {
				// ERC20 transfer(address,uint256): selector, then two 32-byte words.
				if len(mcmsTx.Data) != 4+32+32 {
					continue
				}
				receiver := common.BytesToAddress(mcmsTx.Data[4+12 : 4+32])
				role, ok := roleOf[receiver]
				require.True(t, ok, "transfer to unexpected receiver %s", receiver)
				routes = append(routes, fmt.Sprintf("%s:%s", role, new(big.Int).SetBytes(mcmsTx.Data[4+32:])))
			}
		}
		require.NotEmpty(t, routes)

		return routes
	}

	first := transferRoutes()
	require.Equal(t, []string{"silo1:3000", "silo2:2000", "shared:1000"}, first,
		"each silo keeps its own balance and the shared balance goes to the named lockbox")
	require.Equal(t, first, transferRoutes(), "identical input must produce identical routing")
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
			ChainSelector:     chainSel,
			ContractParams:    testsetup.CreateBasicContractParams(),
			CREATE2Factory:    common.HexToAddress(create2FactoryRef.Address),
			DeployerKeyOwned:  true,
			ExistingAddresses: testsetup.UltraFastCurseMCMSRefs(chainSel),
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
	_, err = operations.ExecuteOperation(e.OperationsBundle, old_siloed.UpdateSiloDesignations, chain,
		evm_contract.FunctionInput[old_siloed.UpdateSiloDesignationsArgs]{
			ChainSelector: chainSel, Address: oldPoolAddr,
			Args: old_siloed.UpdateSiloDesignationsArgs{
				Removes: []uint64{},
				Adds: []old_siloed.SiloConfigUpdate{
					{RemoteChainSelector: remoteChain1, Rebalancer: deployer},
					{RemoteChainSelector: remoteChain2, Rebalancer: deployer},
				},
			},
		})
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
			Args: erc20.ApproveArgs{Spender: oldPoolAddr, Value: totalMint},
		})
	require.NoError(t, err)

	for _, silo := range []struct {
		remoteChainSelector uint64
		amount              *big.Int
	}{
		{remoteChain1, silo1Amount},
		{remoteChain2, silo2Amount},
	} {
		_, err = operations.ExecuteOperation(e.OperationsBundle, old_siloed.ProvideSiloedLiquidity, chain,
			evm_contract.FunctionInput[old_siloed.ProvideSiloedLiquidityArgs]{
				ChainSelector: chainSel, Address: oldPoolAddr,
				Args: old_siloed.ProvideSiloedLiquidityArgs{
					RemoteChainSelector: silo.remoteChainSelector,
					Amount:              silo.amount,
				},
			})
		require.NoError(t, err)
	}

	_, err = operations.ExecuteOperation(e.OperationsBundle, old_siloed.ProvideLiquidity, chain,
		evm_contract.FunctionInput[*big.Int]{ChainSelector: chainSel, Address: oldPoolAddr, Args: unsiloedAmount})
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
			Qualifier:      new("chain1"),
		})
	require.NoError(t, err)
	lockbox1Addr := common.HexToAddress(lockbox1Report.Output.Address)

	lockbox2Report, err := operations.ExecuteOperation(e.OperationsBundle, erc20_lock_box.Deploy, chain,
		evm_contract.DeployInput[erc20_lock_box.ConstructorArgs]{
			ChainSelector:  chainSel,
			TypeAndVersion: deployment.NewTypeAndVersion(erc20_lock_box.ContractType, *erc20_lock_box.Version),
			Args:           erc20_lock_box.ConstructorArgs{Token: tokenAddr},
			Qualifier:      new("chain2"),
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
	executeMigrationSequence(t,
		testsetup.BundleWithFreshReporter(e.OperationsBundle),
		e.BlockChains,
		tokens_core.MigrateLockReleasePoolLiquidityInput{
			ChainSelector:   chainSel,
			OldPoolAddress:  oldPoolAddr.Hex(),
			NewPoolAddress:  newPoolAddr.Hex(),
			TimelockAddress: deployer.Hex(),
			BasisPoints:     &basisPoints,
			// Every chain on this pool is siloed, so there is no separate shared lockbox; the
			// unsiloed balance is sent to chain 1's lockbox explicitly.
			UnsiloedLockBoxAddress: lockbox1Addr.Hex(),
		},
	)

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
	executeMigrationSequence(t,
		testsetup.BundleWithFreshReporter(s.env.OperationsBundle),
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

func TestMigrateLockReleasePoolLiquidity_DoesNotTamperWithAuthorizedCallers(t *testing.T) {
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

	// Run migration (transfer-based, never touches authorized callers)
	basisPoints := uint16(10000)
	output := executeMigrationSequence(t,
		testsetup.BundleWithFreshReporter(s.env.OperationsBundle),
		s.env.BlockChains,
		tokens_core.MigrateLockReleasePoolLiquidityInput{
			ChainSelector:    chainSel,
			OldPoolAddress:   s.oldPoolAddr.Hex(),
			NewPoolAddress:   s.newPoolAddr.Hex(),
			TimelockAddress:  s.deployer.Hex(),
			BasisPoints:      &basisPoints,
			UsePlainTransfer: true,
		},
	)

	// Verify no applyAuthorizedCallerUpdates calls appear in any batch transaction
	parsedABI, err := abi.JSON(strings.NewReader(erc20_lock_box.ERC20LockBoxABI))
	require.NoError(t, err)
	authSelector := parsedABI.Methods["applyAuthorizedCallerUpdates"].ID
	for _, batch := range output.BatchOps {
		for _, tx := range batch.Transactions {
			if len(tx.Data) >= 4 {
				require.NotEqual(t, authSelector[:], tx.Data[:4],
					"migration batch should not contain applyAuthorizedCallerUpdates calls")
			}
		}
	}

	// Verify pre-existing caller is still present after migration
	postCallersReport, err := operations.ExecuteOperation(
		testsetup.BundleWithFreshReporter(s.env.OperationsBundle), erc20_lock_box.GetAllAuthorizedCallers, chain,
		evm_contract.FunctionInput[struct{}]{ChainSelector: chainSel, Address: s.lockBoxAddr})
	require.NoError(t, err)
	require.Contains(t, postCallersReport.Output, preExistingCaller,
		"Pre-existing authorized caller should be preserved after migration (transfer does not touch authorized callers)")

	// Verify timelock was never added to authorized callers
	require.NotContains(t, postCallersReport.Output, s.deployer,
		"Timelock should not be in lockbox authorized callers (transfer-only flow)")
}

func TestMigrateLockReleasePoolLiquidity_MultiplePartialMigrations(t *testing.T) {
	chainSel := uint64(5009297550715157269)
	totalLiquidity := big.NewInt(10000)
	s := setupMigrationTest(t, chainSel, totalLiquidity)
	chain := s.env.BlockChains.EVMChains()[chainSel]

	// Step 1: Migrate 50%
	basisPoints := uint16(5000)
	executeMigrationSequence(t,
		testsetup.BundleWithFreshReporter(s.env.OperationsBundle),
		s.env.BlockChains,
		tokens_core.MigrateLockReleasePoolLiquidityInput{
			ChainSelector:   chainSel,
			OldPoolAddress:  s.oldPoolAddr.Hex(),
			NewPoolAddress:  s.newPoolAddr.Hex(),
			TimelockAddress: s.deployer.Hex(),
			BasisPoints:     &basisPoints,
		},
	)

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
	executeMigrationSequence(t,
		testsetup.BundleWithFreshReporter(s.env.OperationsBundle),
		s.env.BlockChains,
		tokens_core.MigrateLockReleasePoolLiquidityInput{
			ChainSelector:   chainSel,
			OldPoolAddress:  s.oldPoolAddr.Hex(),
			NewPoolAddress:  s.newPoolAddr.Hex(),
			TimelockAddress: s.deployer.Hex(),
			BasisPoints:     &basisPoints,
		},
	)

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

func TestMigrateSiloedPool_ZeroUnsiloedLiquidity_OmittedDestination(t *testing.T) {
	chainSel := uint64(5009297550715157269)
	siloed1 := uint64(3379446385462418246)
	siloAmount := big.NewInt(3000)

	s := setupSiloedMigrationTest(t, chainSel, siloedTopology{
		siloedChains:   []siloedChainSpec{{selector: siloed1, amount: siloAmount}},
		sharedChains:   []uint64{},
		unsiloedAmount: big.NewInt(0), // Zero shared liquidity
	})

	basisPoints := uint16(10000)
	executeMigrationSequence(t,
		testsetup.BundleWithFreshReporter(s.env.OperationsBundle),
		s.env.BlockChains,
		tokens_core.MigrateLockReleasePoolLiquidityInput{
			ChainSelector:          chainSel,
			OldPoolAddress:         s.oldPoolAddr.Hex(),
			NewPoolAddress:         s.newPoolAddr.Hex(),
			TimelockAddress:        s.deployer.Hex(),
			BasisPoints:            &basisPoints,
			UnsiloedLockBoxAddress: "", // Empty destination. This is fine because shared liquidity is 0.
		},
	)

	require.Equal(t, 0, big.NewInt(0).Cmp(s.balanceOf(t, s.oldPoolAddr)), "Old pool balance should be zero")
	require.Equal(t, 0, siloAmount.Cmp(s.balanceOf(t, s.siloLockBoxes[siloed1])), "Silo lockbox should have its correct balance")
}

func TestMigrateSiloedPool_PartialBasisPoints(t *testing.T) {
	chainSel := uint64(5009297550715157269)
	siloed1 := uint64(3379446385462418246)
	siloed2 := uint64(4949039107694359620)
	shared1 := uint64(15971525489660198786)

	silo1Amount := big.NewInt(4000)
	silo2Amount := big.NewInt(2000)
	unsiloedAmount := big.NewInt(1000)

	s := setupSiloedMigrationTest(t, chainSel, siloedTopology{
		siloedChains: []siloedChainSpec{
			{selector: siloed1, amount: silo1Amount},
			{selector: siloed2, amount: silo2Amount},
		},
		sharedChains:   []uint64{shared1},
		unsiloedAmount: unsiloedAmount,
	})

	basisPoints := uint16(5000) // 50% migration
	executeMigrationSequence(t,
		testsetup.BundleWithFreshReporter(s.env.OperationsBundle),
		s.env.BlockChains,
		tokens_core.MigrateLockReleasePoolLiquidityInput{
			ChainSelector:          chainSel,
			OldPoolAddress:         s.oldPoolAddr.Hex(),
			NewPoolAddress:         s.newPoolAddr.Hex(),
			TimelockAddress:        s.deployer.Hex(),
			BasisPoints:            &basisPoints,
			UnsiloedLockBoxAddress: s.sharedLockBox.Hex(),
		},
	)

	// Expected to migrate 50%
	expectedSilo1Migrated := new(big.Int).Div(new(big.Int).Mul(silo1Amount, big.NewInt(5000)), big.NewInt(10000))
	expectedSilo2Migrated := new(big.Int).Div(new(big.Int).Mul(silo2Amount, big.NewInt(5000)), big.NewInt(10000))
	expectedUnsiloedMigrated := new(big.Int).Div(new(big.Int).Mul(unsiloedAmount, big.NewInt(5000)), big.NewInt(10000))
	expectedTotalMigrated := new(big.Int).Add(expectedSilo1Migrated, new(big.Int).Add(expectedSilo2Migrated, expectedUnsiloedMigrated))

	expectedOldPoolBalance := new(big.Int).Sub(s.totalMint, expectedTotalMigrated)

	require.Equal(t, 0, expectedOldPoolBalance.Cmp(s.balanceOf(t, s.oldPoolAddr)), "Old pool should have exactly 50%% of its original balance")

	require.Equal(t, 0, expectedSilo1Migrated.Cmp(s.balanceOf(t, s.siloLockBoxes[siloed1])), "Silo 1 lockbox should have exactly 50%% of its balance")
	require.Equal(t, 0, expectedSilo2Migrated.Cmp(s.balanceOf(t, s.siloLockBoxes[siloed2])), "Silo 2 lockbox should have exactly 50%% of its balance")
	require.Equal(t, 0, expectedUnsiloedMigrated.Cmp(s.balanceOf(t, s.sharedLockBox)), "Shared lockbox should have exactly 50%% of its balance")
}

func TestMigrateSiloedPool_SiloRebalancerRestore(t *testing.T) {
	chainSel := uint64(5009297550715157269)
	siloed1 := uint64(3379446385462418246)
	siloAmount := big.NewInt(3000)

	s := setupSiloedMigrationTest(t, chainSel, siloedTopology{
		siloedChains:   []siloedChainSpec{{selector: siloed1, amount: siloAmount}},
		sharedChains:   []uint64{},
		unsiloedAmount: big.NewInt(0),
	})

	chain := s.env.BlockChains.EVMChains()[chainSel]

	// Set a custom silo rebalancer address before running migration
	originalRebalancer := common.HexToAddress("0x2222222222222222222222222222222222222222")
	_, err := operations.ExecuteOperation(
		s.env.OperationsBundle,
		old_siloed.SetSiloRebalancer,
		chain,
		evm_contract.FunctionInput[old_siloed.SetSiloRebalancerArgs]{
			ChainSelector: chainSel,
			Address:       s.oldPoolAddr,
			Args: old_siloed.SetSiloRebalancerArgs{
				RemoteChainSelector: siloed1,
				NewRebalancer:       originalRebalancer,
			},
		},
	)
	require.NoError(t, err)

	basisPoints := uint16(10000)
	freshBundle := testsetup.BundleWithFreshReporter(s.env.OperationsBundle)
	executeMigrationSequence(t,
		freshBundle,
		s.env.BlockChains,
		tokens_core.MigrateLockReleasePoolLiquidityInput{
			ChainSelector:   chainSel,
			OldPoolAddress:  s.oldPoolAddr.Hex(),
			NewPoolAddress:  s.newPoolAddr.Hex(),
			TimelockAddress: s.deployer.Hex(),
			BasisPoints:     &basisPoints,
		},
	)

	rebalancerReport, err := operations.ExecuteOperation(
		testsetup.BundleWithFreshReporter(s.env.OperationsBundle),
		old_siloed.GetChainRebalancer,
		chain,
		evm_contract.FunctionInput[uint64]{
			ChainSelector: chainSel,
			Address:       s.oldPoolAddr,
			Args:          siloed1,
		},
	)
	require.NoError(t, err)
	require.Equal(t, originalRebalancer, rebalancerReport.Output, "Silo rebalancer should be restored to the custom address set before migration")
}

func TestMigrateSiloedPool_SharedLockBox_MultiplePartialMigrations(t *testing.T) {
	chainSel := uint64(5009297550715157269)
	siloed1 := uint64(3379446385462418246)
	shared1 := uint64(15971525489660198786)

	siloAmount := big.NewInt(4000)
	unsiloedAmount := big.NewInt(2000)

	s := setupSiloedMigrationTest(t, chainSel, siloedTopology{
		siloedChains:   []siloedChainSpec{{selector: siloed1, amount: siloAmount}},
		sharedChains:   []uint64{shared1},
		unsiloedAmount: unsiloedAmount,
	})

	// Step 1: Migrate 50%
	basisPoints := uint16(5000)
	executeMigrationSequence(t,
		testsetup.BundleWithFreshReporter(s.env.OperationsBundle),
		s.env.BlockChains,
		tokens_core.MigrateLockReleasePoolLiquidityInput{
			ChainSelector:          chainSel,
			OldPoolAddress:         s.oldPoolAddr.Hex(),
			NewPoolAddress:         s.newPoolAddr.Hex(),
			TimelockAddress:        s.deployer.Hex(),
			BasisPoints:            &basisPoints,
			UnsiloedLockBoxAddress: s.sharedLockBox.Hex(),
		},
	)

	expectedSilo1Migrated := new(big.Int).Div(new(big.Int).Mul(siloAmount, big.NewInt(5000)), big.NewInt(10000))
	expectedUnsiloedMigrated := new(big.Int).Div(new(big.Int).Mul(unsiloedAmount, big.NewInt(5000)), big.NewInt(10000))

	require.Equal(t, 0, expectedSilo1Migrated.Cmp(s.balanceOf(t, s.siloLockBoxes[siloed1])), "Silo 1 lockbox should have exactly 50%% after the first migration")
	require.Equal(t, 0, expectedUnsiloedMigrated.Cmp(s.balanceOf(t, s.sharedLockBox)), "Shared lockbox should have exactly 50%% after the first migration")

	// Step 2: Migrate the remaining balance (100% of what is left)
	basisPoints = uint16(10000)
	executeMigrationSequence(t,
		testsetup.BundleWithFreshReporter(s.env.OperationsBundle),
		s.env.BlockChains,
		tokens_core.MigrateLockReleasePoolLiquidityInput{
			ChainSelector:          chainSel,
			OldPoolAddress:         s.oldPoolAddr.Hex(),
			NewPoolAddress:         s.newPoolAddr.Hex(),
			TimelockAddress:        s.deployer.Hex(),
			BasisPoints:            &basisPoints,
			UnsiloedLockBoxAddress: s.sharedLockBox.Hex(),
		},
	)

	require.Equal(t, 0, big.NewInt(0).Cmp(s.balanceOf(t, s.oldPoolAddr)), "Old pool balance should be zero after the second migration")
	require.Equal(t, 0, siloAmount.Cmp(s.balanceOf(t, s.siloLockBoxes[siloed1])), "Silo 1 lockbox should have all its tokens after the second migration")
	require.Equal(t, 0, unsiloedAmount.Cmp(s.balanceOf(t, s.sharedLockBox)), "Shared lockbox should have all its tokens after the second migration")
}
