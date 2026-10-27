package tokens_test

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/require"

	mcms_types "github.com/smartcontractkit/mcms/types"

	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	evm_contract "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations/contract"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/erc20"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/burn_mint_erc20_with_drip"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/testsetup"
	lock_release_token_pool_v120 "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_2_0/lock_release_token_pool"
	lock_release_token_pool_v140 "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_4_0/lock_release_token_pool"
	burn_mint_token_pool_and_proxy_v150 "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_5_0/burn_mint_token_pool_and_proxy"
	lock_release_token_pool_and_proxy_v150 "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_5_0/lock_release_token_pool_and_proxy"
	tokens_core "github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	seqtypes "github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
)

// legacyRebalancer is the rebalancer the v1.4.0 previous pool and the proxy start with. The
// migration hands the rebalancer role to the timelock and must restore this one.
var legacyRebalancer = common.HexToAddress("0x000000000000000000000000000000000000BEEF")

// Sequence IDs of the liquidity migration sources, to tell which one ran.
const (
	v150SourceID = "evm-v1.5.0-adapter:migrate-lock-release-pool-liquidity-source"
	v140SourceID = "evm-v1.4.0-adapter:migrate-lock-release-pool-liquidity-source"
	v120SourceID = "evm-v1.2.0-adapter:migrate-lock-release-pool-liquidity-source"
)

// ownable is the ownership surface every legacy pool binding shares.
type ownable interface {
	TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error)
	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)
}

// proxyFixture is a v1.5.0 LockReleaseTokenPoolAndProxy whose getPreviousPool() points at a legacy
// pool, the way production proxies front v1.2.0 and v1.4.0 pools. The embedded setup's oldPoolAddr
// is the PREVIOUS pool; the migration input names the proxy.
type proxyFixture struct {
	migrationTestSetup
	chain       evm.Chain
	proxy       *lock_release_token_pool_and_proxy_v150.LockReleaseTokenPoolAndProxy
	proxyAddr   common.Address
	timelockKey *bind.TransactOpts
	timelock    common.Address
}

// setupProxyFixture deploys a previous pool of previousKind holding previousLiquidity and a proxy in
// front of it holding proxyLiquidity. The proxy's rebalancer, and a v1.4.0 previous pool's, is set to
// legacyRebalancer; for a v1.2.0 previous pool the proxy is registered as its on- and offRamp with
// an enabled limiter, as a production v1.2.0 pool behind a proxy is.
func setupProxyFixture(t *testing.T, previousKind oldPoolKind, previousLiquidity, proxyLiquidity *big.Int) proxyFixture {
	t.Helper()

	chainSel := uint64(5009297550715157269)
	s := setupMigrationTestWithOldPool(t, chainSel, previousLiquidity, previousKind)
	chain := s.env.BlockChains.EVMChains()[chainSel]

	proxyAddr, tx, proxy, err := lock_release_token_pool_and_proxy_v150.DeployLockReleaseTokenPoolAndProxy(
		chain.DeployerKey, chain.Client, s.tokenAddr, []common.Address{}, s.rmnProxyAddr, true, s.routerAddr)
	require.NoError(t, err)
	confirm(t, chain, tx)
	sendTx(t, chain)(proxy.SetPreviousPool(chain.DeployerKey, s.oldPoolAddr))
	sendTx(t, chain)(proxy.SetRebalancer(chain.DeployerKey, legacyRebalancer))

	switch previousKind {
	case oldPoolV1_4_0:
		previous, err := lock_release_token_pool_v140.NewLockReleaseTokenPool(s.oldPoolAddr, chain.Client)
		require.NoError(t, err)
		sendTx(t, chain)(previous.SetRebalancer(chain.DeployerKey, legacyRebalancer))
	case oldPoolV1_2_0, oldPoolV1_2_0AcceptingLiquidity:
		previous, err := lock_release_token_pool_v120.NewLockReleaseTokenPool(s.oldPoolAddr, chain.Client)
		require.NoError(t, err)
		limiter := lock_release_token_pool_v120.RateLimiterConfig{IsEnabled: true, Capacity: big.NewInt(1_000_000), Rate: big.NewInt(1)}
		update := []lock_release_token_pool_v120.TokenPoolRampUpdate{{Ramp: proxyAddr, Allowed: true, RateLimiterConfig: limiter}}
		sendTx(t, chain)(previous.ApplyRampUpdates(chain.DeployerKey, update, update))
	}

	if proxyLiquidity.Sign() > 0 {
		_, err = operations.ExecuteOperation(s.env.OperationsBundle, burn_mint_erc20_with_drip.Mint, chain,
			evm_contract.FunctionInput[burn_mint_erc20_with_drip.MintArgs]{
				ChainSelector: chainSel,
				Address:       s.tokenAddr,
				Args:          burn_mint_erc20_with_drip.MintArgs{Account: proxyAddr, Amount: proxyLiquidity},
			})
		require.NoError(t, err)
	}

	timelockKey := newFundedKey(t, chain)
	return proxyFixture{
		migrationTestSetup: s,
		chain:              chain,
		proxy:              proxy,
		proxyAddr:          proxyAddr,
		timelockKey:        timelockKey,
		timelock:           timelockKey.From,
	}
}

// transferToTimelock hands ownership of a pool to the fixture's timelock.
func (f proxyFixture) transferToTimelock(t *testing.T, pool ownable) {
	t.Helper()
	sendTx(t, f.chain)(pool.TransferOwnership(f.chain.DeployerKey, f.timelock))
	sendTx(t, f.chain)(pool.AcceptOwnership(f.timelockKey))
}

// previousOwnable binds the previous pool's ownership surface for its kind.
func (f proxyFixture) previousOwnable(t *testing.T, kind oldPoolKind) ownable {
	t.Helper()
	switch kind {
	case oldPoolV1_4_0:
		pool, err := lock_release_token_pool_v140.NewLockReleaseTokenPool(f.oldPoolAddr, f.chain.Client)
		require.NoError(t, err)
		return pool
	case oldPoolV1_2_0, oldPoolV1_2_0AcceptingLiquidity:
		pool, err := lock_release_token_pool_v120.NewLockReleaseTokenPool(f.oldPoolAddr, f.chain.Client)
		require.NoError(t, err)
		return pool
	case oldPoolV1_5_0AndProxy:
		pool, err := lock_release_token_pool_and_proxy_v150.NewLockReleaseTokenPoolAndProxy(f.oldPoolAddr, f.chain.Client)
		require.NoError(t, err)
		return pool
	}
	t.Fatalf("unsupported previous pool kind %d", kind)
	return nil
}

func (f proxyFixture) balanceOf(t *testing.T, holder common.Address) *big.Int {
	t.Helper()
	report, err := operations.ExecuteOperation(testsetup.BundleWithFreshReporter(f.env.OperationsBundle), erc20.BalanceOf, f.chain,
		evm_contract.FunctionInput[common.Address]{ChainSelector: f.chainSel, Address: f.tokenAddr, Args: holder})
	require.NoError(t, err)
	return report.Output
}

// input names the PROXY as the old pool, as the TokenAdminRegistry and operator config do.
func (f proxyFixture) input(input tokens_core.MigrateLockReleasePoolLiquidityInput) tokens_core.MigrateLockReleasePoolLiquidityInput {
	input.ChainSelector = f.chainSel
	input.OldPoolAddress = f.proxyAddr.Hex()
	input.NewPoolAddress = f.newPoolAddr.Hex()
	input.TimelockAddress = f.timelock.Hex()
	return input
}

// migrate runs the migration on a fresh reporter and returns the IDs of every operation and sequence
// that ran, so a test can tell which source extracted the liquidity even when two sources would emit
// identical calldata (the v1.4.0 and v1.5.0 rebalancer selectors are the same).
func (f proxyFixture) migrate(t *testing.T, input tokens_core.MigrateLockReleasePoolLiquidityInput) (seqtypes.OnChainOutput, []string, error) {
	t.Helper()
	reporter := operations.NewMemoryReporter()
	bundle := operations.NewBundle(f.env.OperationsBundle.GetContext, f.env.OperationsBundle.Logger, reporter)
	output, err := migrateLockReleasePoolLiquidity(bundle, f.env.BlockChains, f.input(input))
	reports, reportsErr := reporter.GetReports()
	require.NoError(t, reportsErr)
	ids := make([]string, 0, len(reports))
	for _, report := range reports {
		ids = append(ids, report.Def.ID)
	}
	return output, ids, err
}

// requireProxyUnchanged checks the proxy still fronts the previous pool with its original rebalancer.
func (f proxyFixture) requireProxyUnchanged(t *testing.T) {
	t.Helper()
	callOpts := &bind.CallOpts{Context: t.Context()}
	previous, err := f.proxy.GetPreviousPool(callOpts)
	require.NoError(t, err)
	require.Equal(t, f.oldPoolAddr, previous, "the proxy must keep fronting its previous pool")
	rebalancer, err := f.proxy.GetRebalancer(callOpts)
	require.NoError(t, err)
	require.Equal(t, legacyRebalancer, rebalancer, "the proxy's rebalancer must be untouched")
}

// requireSourceThenSink checks the merged batch: the source writes first, each targeting target with
// exactly the expected calldata, then only the sink's funding of the lockbox.
func (f proxyFixture) requireSourceThenSink(t *testing.T, txs []mcms_types.Transaction, target common.Address, expectedSource [][]byte) {
	t.Helper()
	require.Greater(t, len(txs), len(expectedSource), "the sink's funding must follow the source writes")
	for i, want := range expectedSource {
		require.Equal(t, target, common.HexToAddress(txs[i].To), "source tx %d target", i)
		require.Equal(t, want, txs[i].Data, "source tx %d calldata", i)
	}
	for i, sinkTx := range txs[len(expectedSource):] {
		to := common.HexToAddress(sinkTx.To)
		require.True(t, to == f.tokenAddr || to == f.lockBoxAddr,
			"sink tx %d must target the token or the lockbox, got %s", i, to)
	}
	require.Equal(t, f.lockBoxAddr, common.HexToAddress(txs[len(txs)-1].To), "the batch must end with the lockbox deposit")
}

// TestMigrateLockReleasePoolLiquidity_ProxyDelegatesToPreviousPool migrates a v1.5.0 proxy whose
// liquidity sits in its previous pool. The input names the proxy; the v1.5.0 source must delegate to
// the previous pool's own source (v1.4.0 rebalancer withdraw or v1.2.0 offRamp rescue), so every
// source write targets the previous pool and none targets the proxy.
func TestMigrateLockReleasePoolLiquidity_ProxyDelegatesToPreviousPool(t *testing.T) {
	bps5000 := uint16(5000)
	bps10000 := uint16(10000)

	v140ABI, err := lock_release_token_pool_v140.LockReleaseTokenPoolMetaData.GetAbi()
	require.NoError(t, err)
	v120ABI, err := lock_release_token_pool_v120.LockReleaseTokenPoolMetaData.GetAbi()
	require.NoError(t, err)

	cases := []struct {
		name           string
		previousKind   oldPoolKind
		input          tokens_core.MigrateLockReleasePoolLiquidityInput
		expectedAmount *big.Int
		// expectedSource is the sequence ID of the previous pool's own source.
		expectedSource string
	}{
		{name: "V1_4_0/BasisPoints", previousKind: oldPoolV1_4_0, input: tokens_core.MigrateLockReleasePoolLiquidityInput{BasisPoints: &bps5000}, expectedAmount: big.NewInt(5000), expectedSource: v140SourceID},
		{name: "V1_4_0/ExactAmount", previousKind: oldPoolV1_4_0, input: tokens_core.MigrateLockReleasePoolLiquidityInput{Amount: big.NewInt(3000)}, expectedAmount: big.NewInt(3000), expectedSource: v140SourceID},
		{name: "V1_2_0/BasisPoints", previousKind: oldPoolV1_2_0, input: tokens_core.MigrateLockReleasePoolLiquidityInput{BasisPoints: &bps5000}, expectedAmount: big.NewInt(5000), expectedSource: v120SourceID},
		{name: "V1_2_0/FullBalance", previousKind: oldPoolV1_2_0, input: tokens_core.MigrateLockReleasePoolLiquidityInput{BasisPoints: &bps10000}, expectedAmount: big.NewInt(10000), expectedSource: v120SourceID},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			totalLiquidity := big.NewInt(10000)
			f := setupProxyFixture(t, tc.previousKind, totalLiquidity, big.NewInt(0))
			f.transferToTimelock(t, f.proxy)
			f.transferToTimelock(t, f.previousOwnable(t, tc.previousKind))

			report, ran, err := f.migrate(t, tc.input)
			require.NoError(t, err)
			require.Contains(t, ran, v150SourceID, "the proxy's version resolves the v1.5.0 source")
			require.Contains(t, ran, tc.expectedSource, "the v1.5.0 source must delegate to the previous pool's own source")
			require.Equal(t, 0, totalLiquidity.Cmp(f.balanceOf(t, f.oldPoolAddr)), "previous pool must keep its liquidity before the proposal executes")

			var expectedSource [][]byte
			pack := func(abiPack func(string, ...any) ([]byte, error), method string, args ...any) []byte {
				data, err := abiPack(method, args...)
				require.NoError(t, err)
				return data
			}
			switch tc.previousKind {
			case oldPoolV1_4_0:
				expectedSource = [][]byte{
					pack(v140ABI.Pack, "setRebalancer", f.timelock),
					pack(v140ABI.Pack, "withdrawLiquidity", tc.expectedAmount),
					pack(v140ABI.Pack, "setRebalancer", legacyRebalancer),
				}
			case oldPoolV1_2_0:
				disabled := lock_release_token_pool_v120.RateLimiterConfig{IsEnabled: false, Capacity: big.NewInt(0), Rate: big.NewInt(0)}
				rampUpdate := func(allowed bool) []byte {
					return pack(v120ABI.Pack, "applyRampUpdates", []lock_release_token_pool_v120.TokenPoolRampUpdate{},
						[]lock_release_token_pool_v120.TokenPoolRampUpdate{{Ramp: f.timelock, Allowed: allowed, RateLimiterConfig: disabled}})
				}
				expectedSource = [][]byte{
					rampUpdate(true),
					pack(v120ABI.Pack, "releaseOrMint", []byte{}, f.timelock, tc.expectedAmount, uint64(0), []byte{}),
					rampUpdate(false),
				}
			}
			require.Len(t, report.BatchOps, 1)
			f.requireSourceThenSink(t, report.BatchOps[0].Transactions, f.oldPoolAddr, expectedSource)

			executeBatchOpsAs(t, f.chain, f.timelockKey, report.BatchOps)

			remaining := new(big.Int).Sub(totalLiquidity, tc.expectedAmount)
			require.Equal(t, 0, remaining.Cmp(f.balanceOf(t, f.oldPoolAddr)), "previous pool should keep the rest of its liquidity")
			require.Equal(t, 0, tc.expectedAmount.Cmp(f.balanceOf(t, f.lockBoxAddr)), "lockbox should hold the migrated liquidity")
			require.Equal(t, 0, big.NewInt(0).Cmp(f.balanceOf(t, f.proxyAddr)), "the proxy must hold nothing")
			require.Equal(t, 0, big.NewInt(0).Cmp(f.balanceOf(t, f.timelock)), "no liquidity may be left in the timelock")
			f.requireProxyUnchanged(t)

			callOpts := &bind.CallOpts{Context: t.Context()}
			switch tc.previousKind {
			case oldPoolV1_4_0:
				previous, err := lock_release_token_pool_v140.NewLockReleaseTokenPool(f.oldPoolAddr, f.chain.Client)
				require.NoError(t, err)
				rebalancer, err := previous.GetRebalancer(callOpts)
				require.NoError(t, err)
				require.Equal(t, legacyRebalancer, rebalancer, "previous pool rebalancer must be restored")
			case oldPoolV1_2_0:
				previous, err := lock_release_token_pool_v120.NewLockReleaseTokenPool(f.oldPoolAddr, f.chain.Client)
				require.NoError(t, err)
				offRamps, err := previous.GetOffRamps(callOpts)
				require.NoError(t, err)
				require.Equal(t, []common.Address{f.proxyAddr}, offRamps, "only the proxy may remain an offRamp")
				bucket, err := previous.CurrentOffRampRateLimiterState(callOpts, f.proxyAddr)
				require.NoError(t, err)
				require.True(t, bucket.IsEnabled)
				require.Equal(t, 0, big.NewInt(1_000_000).Cmp(bucket.Capacity), "the proxy's limiter must be untouched")
			}
		})
	}
}

// TestMigrateLockReleasePoolLiquidity_ProxyHoldsLiquidityWithPreviousPool migrates a proxy that holds
// the liquidity itself while a previous pool is set but empty: the v1.5.0 rebalancer withdraw runs on
// the proxy and the previous pool is left alone.
func TestMigrateLockReleasePoolLiquidity_ProxyHoldsLiquidityWithPreviousPool(t *testing.T) {
	proxyABI, err := lock_release_token_pool_and_proxy_v150.LockReleaseTokenPoolAndProxyMetaData.GetAbi()
	require.NoError(t, err)

	proxyLiquidity := big.NewInt(10000)
	f := setupProxyFixture(t, oldPoolV1_4_0, big.NewInt(0), proxyLiquidity)
	f.transferToTimelock(t, f.proxy)
	// The previous pool stays deployer-owned: the proxy path must not touch or require it.

	bps5000 := uint16(5000)
	report, ran, err := f.migrate(t, tokens_core.MigrateLockReleasePoolLiquidityInput{BasisPoints: &bps5000})
	require.NoError(t, err)
	require.NotContains(t, ran, v140SourceID, "a proxy holding the liquidity must not delegate")

	expectedAmount := big.NewInt(5000)
	pack := func(method string, args ...any) []byte {
		data, err := proxyABI.Pack(method, args...)
		require.NoError(t, err)
		return data
	}
	require.Len(t, report.BatchOps, 1)
	f.requireSourceThenSink(t, report.BatchOps[0].Transactions, f.proxyAddr, [][]byte{
		pack("setRebalancer", f.timelock),
		pack("withdrawLiquidity", expectedAmount),
		pack("setRebalancer", legacyRebalancer),
	})

	executeBatchOpsAs(t, f.chain, f.timelockKey, report.BatchOps)

	require.Equal(t, 0, new(big.Int).Sub(proxyLiquidity, expectedAmount).Cmp(f.balanceOf(t, f.proxyAddr)), "proxy should keep the rest of its liquidity")
	require.Equal(t, 0, expectedAmount.Cmp(f.balanceOf(t, f.lockBoxAddr)), "lockbox should hold the migrated liquidity")
	require.Equal(t, 0, big.NewInt(0).Cmp(f.balanceOf(t, f.timelock)), "no liquidity may be left in the timelock")
	f.requireProxyUnchanged(t)
}

// TestMigrateLockReleasePoolLiquidity_ProxyPreconditions checks every way a proxy migration must fail
// before emitting anything, leaving the liquidity of both the proxy and the previous pool untouched.
func TestMigrateLockReleasePoolLiquidity_ProxyPreconditions(t *testing.T) {
	bps10000 := uint16(10000)
	full := tokens_core.MigrateLockReleasePoolLiquidityInput{BasisPoints: &bps10000}

	cases := []struct {
		name              string
		previousKind      oldPoolKind
		previousLiquidity int64
		proxyLiquidity    int64
		// prepare runs after the default ownership hand-over; it may repoint the proxy.
		prepare func(t *testing.T, f proxyFixture)
		// keepPreviousOwner leaves the previous pool owned by the deployer.
		keepPreviousOwner bool
		input             tokens_core.MigrateLockReleasePoolLiquidityInput
		expectedErr       string
	}{
		{
			name:         "Split",
			previousKind: oldPoolV1_4_0, previousLiquidity: 10000, proxyLiquidity: 1,
			input:       full,
			expectedErr: "liquidity is split between proxy",
		},
		{
			name:         "NoLiquidity",
			previousKind: oldPoolV1_4_0,
			input:        full,
			expectedErr:  "holds any liquidity to migrate",
		},
		{
			// The proxy is funded, so without the nested check the proxy path would run.
			name:         "NestedLockReleaseProxy",
			previousKind: oldPoolV1_5_0AndProxy, previousLiquidity: 10000, proxyLiquidity: 10000,
			input:       full,
			expectedErr: "is itself a LockReleaseTokenPoolAndProxy; nested proxies are not supported",
		},
		{
			name:         "NestedBurnMintProxy",
			previousKind: oldPoolV1_4_0, proxyLiquidity: 10000,
			prepare: func(t *testing.T, f proxyFixture) {
				addr, tx, _, err := burn_mint_token_pool_and_proxy_v150.DeployBurnMintTokenPoolAndProxy(
					f.chain.DeployerKey, f.chain.Client, f.tokenAddr, []common.Address{}, f.rmnProxyAddr, f.routerAddr)
				require.NoError(t, err)
				confirm(t, f.chain, tx)
				sendTx(t, f.chain)(f.proxy.SetPreviousPool(f.timelockKey, addr))
			},
			input:       full,
			expectedErr: "is itself a BurnMintTokenPoolAndProxy; nested proxies are not supported",
		},
		{
			name:         "PreviousPoolNotOwnedByTimelock",
			previousKind: oldPoolV1_4_0, previousLiquidity: 10000,
			keepPreviousOwner: true,
			input:             full,
			expectedErr:       "is owned by",
		},
		{
			name:         "PreviousAmountExceedsBalance",
			previousKind: oldPoolV1_4_0, previousLiquidity: 10000,
			input:       tokens_core.MigrateLockReleasePoolLiquidityInput{Amount: big.NewInt(10001)},
			expectedErr: "exceeds old pool balance",
		},
		{
			// The delegated v1.2.0 source's own guards still apply to the previous pool.
			name:         "PreviousV1_2_0AcceptsLiquidity",
			previousKind: oldPoolV1_2_0AcceptingLiquidity, previousLiquidity: 10000,
			input:       full,
			expectedErr: "accepts liquidity, so provider credits may exist",
		},
		{
			name:         "ExactSiloMode",
			previousKind: oldPoolV1_4_0, previousLiquidity: 10000,
			input: tokens_core.MigrateLockReleasePoolLiquidityInput{
				SiloExactAmounts: []tokens_core.SiloExactAmount{{ChainSelector: 3379446385462418246, Amount: big.NewInt(100)}},
			},
			expectedErr: "only supported for siloed pools",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			f := setupProxyFixture(t, tc.previousKind, big.NewInt(tc.previousLiquidity), big.NewInt(tc.proxyLiquidity))
			f.transferToTimelock(t, f.proxy)
			if !tc.keepPreviousOwner {
				f.transferToTimelock(t, f.previousOwnable(t, tc.previousKind))
			}
			if tc.prepare != nil {
				tc.prepare(t, f)
			}

			_, err := migrateLockReleasePoolLiquidity(testsetup.BundleWithFreshReporter(f.env.OperationsBundle), f.env.BlockChains, f.input(tc.input))
			require.ErrorContains(t, err, tc.expectedErr)

			require.Equal(t, 0, big.NewInt(tc.previousLiquidity).Cmp(f.balanceOf(t, f.oldPoolAddr)), "previous pool must keep its full balance")
			require.Equal(t, 0, big.NewInt(tc.proxyLiquidity).Cmp(f.balanceOf(t, f.proxyAddr)), "proxy must keep its full balance")
		})
	}
}

// TestMigrateLockReleasePoolLiquidity_ProxyDelegationErrorNamesBothPools checks that an error raised
// by the delegated source names the previous pool and the proxy, since the operator only named the
// proxy.
func TestMigrateLockReleasePoolLiquidity_ProxyDelegationErrorNamesBothPools(t *testing.T) {
	f := setupProxyFixture(t, oldPoolV1_4_0, big.NewInt(10000), big.NewInt(0))
	f.transferToTimelock(t, f.proxy)

	bps10000 := uint16(10000)
	_, err := migrateLockReleasePoolLiquidity(testsetup.BundleWithFreshReporter(f.env.OperationsBundle), f.env.BlockChains,
		f.input(tokens_core.MigrateLockReleasePoolLiquidityInput{BasisPoints: &bps10000}))
	require.ErrorContains(t, err, "previous pool "+f.oldPoolAddr.Hex()+" of proxy "+f.proxyAddr.Hex())
}

func confirm(t *testing.T, chain evm.Chain, tx *types.Transaction) {
	t.Helper()
	_, err := chain.Confirm(tx)
	require.NoError(t, err)
}

// sendTx returns a sink for a binding call's (tx, err) results that requires success and waits for
// the transaction to be mined.
func sendTx(t *testing.T, chain evm.Chain) func(*types.Transaction, error) {
	return func(tx *types.Transaction, err error) {
		t.Helper()
		require.NoError(t, err)
		confirm(t, chain, tx)
	}
}
