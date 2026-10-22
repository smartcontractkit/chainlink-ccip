package tokens_test

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	evm_contract "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations/contract"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/erc20"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/type_and_version"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/testsetup"
	rmn_latest "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/latest/rmn"
	rmn_proxy_latest "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/latest/rmn_proxy_contract"
	lock_release_token_pool_v120 "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_2_0/lock_release_token_pool"
	tokens_core "github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
)

// laneRamp stands in for the fronting v1.5.0 proxy, which is the only ramp a production v1.2.0 pool
// has. The migration must leave it and its limiter alone.
var laneRamp = common.HexToAddress("0x0000000000000000000000000000000000000B0B")

// v120Fixture is a v1.2.0 lock-release pool holding liquidity credited to no provider, with laneRamp
// registered as its on- and offRamp, ready to be handed to a timelock distinct from the deployer.
type v120Fixture struct {
	migrationTestSetup
	chain       evm.Chain
	pool        *lock_release_token_pool_v120.LockReleaseTokenPool
	timelockKey *bind.TransactOpts
	timelock    common.Address
}

func setupV1_2_0Fixture(t *testing.T, kind oldPoolKind, totalLiquidity *big.Int) v120Fixture {
	t.Helper()

	chainSel := uint64(5009297550715157269)
	s := setupMigrationTestWithOldPool(t, chainSel, totalLiquidity, kind)
	chain := s.env.BlockChains.EVMChains()[chainSel]

	tv, err := operations.ExecuteOperation(testsetup.BundleWithFreshReporter(s.env.OperationsBundle), type_and_version.GetTypeAndVersion, chain,
		evm_contract.FunctionInput[struct{}]{ChainSelector: chainSel, Address: s.oldPoolAddr})
	require.NoError(t, err)
	require.Equal(t, "LockReleaseTokenPool", string(tv.Output.Type))
	require.Equal(t, "1.2.0", tv.Output.Version.String())

	pool, err := lock_release_token_pool_v120.NewLockReleaseTokenPool(s.oldPoolAddr, chain.Client)
	require.NoError(t, err)

	laneLimiter := lock_release_token_pool_v120.RateLimiterConfig{IsEnabled: true, Capacity: big.NewInt(1_000_000), Rate: big.NewInt(1)}
	laneUpdate := []lock_release_token_pool_v120.TokenPoolRampUpdate{{Ramp: laneRamp, Allowed: true, RateLimiterConfig: laneLimiter}}
	tx, err := pool.ApplyRampUpdates(chain.DeployerKey, laneUpdate, laneUpdate)
	require.NoError(t, err)
	_, err = chain.Confirm(tx)
	require.NoError(t, err)

	timelockKey := newFundedKey(t, chain)
	return v120Fixture{migrationTestSetup: s, chain: chain, pool: pool, timelockKey: timelockKey, timelock: timelockKey.From}
}

// registerTimelockOffRamp registers the timelock as an offRamp with the given limiter. Must run while
// the deployer still owns the pool.
func (f v120Fixture) registerTimelockOffRamp(t *testing.T, limiter lock_release_token_pool_v120.RateLimiterConfig) {
	t.Helper()
	tx, err := f.pool.ApplyRampUpdates(f.chain.DeployerKey, nil,
		[]lock_release_token_pool_v120.TokenPoolRampUpdate{{Ramp: f.timelock, Allowed: true, RateLimiterConfig: limiter}})
	require.NoError(t, err)
	_, err = f.chain.Confirm(tx)
	require.NoError(t, err)
}

func (f v120Fixture) transferOwnershipToTimelock(t *testing.T) {
	t.Helper()
	tx, err := f.pool.TransferOwnership(f.chain.DeployerKey, f.timelock)
	require.NoError(t, err)
	_, err = f.chain.Confirm(tx)
	require.NoError(t, err)
	tx, err = f.pool.AcceptOwnership(f.timelockKey)
	require.NoError(t, err)
	_, err = f.chain.Confirm(tx)
	require.NoError(t, err)
}

func (f v120Fixture) balanceOf(t *testing.T, holder common.Address) *big.Int {
	t.Helper()
	report, err := operations.ExecuteOperation(testsetup.BundleWithFreshReporter(f.env.OperationsBundle), erc20.BalanceOf, f.chain,
		evm_contract.FunctionInput[common.Address]{ChainSelector: f.chainSel, Address: f.tokenAddr, Args: holder})
	require.NoError(t, err)
	return report.Output
}

func (f v120Fixture) ramps(t *testing.T) (onRamps, offRamps []common.Address) {
	t.Helper()
	callOpts := &bind.CallOpts{Context: t.Context()}
	onRamps, err := f.pool.GetOnRamps(callOpts)
	require.NoError(t, err)
	offRamps, err = f.pool.GetOffRamps(callOpts)
	require.NoError(t, err)
	return onRamps, offRamps
}

func (f v120Fixture) migrate(input tokens_core.MigrateLockReleasePoolLiquidityInput) error {
	_, err := migrateLockReleasePoolLiquidity(testsetup.BundleWithFreshReporter(f.env.OperationsBundle), f.env.BlockChains, f.input(input))
	return err
}

func (f v120Fixture) input(input tokens_core.MigrateLockReleasePoolLiquidityInput) tokens_core.MigrateLockReleasePoolLiquidityInput {
	input.ChainSelector = f.chainSel
	input.OldPoolAddress = f.oldPoolAddr.Hex()
	input.NewPoolAddress = f.newPoolAddr.Hex()
	if input.TimelockAddress == "" {
		input.TimelockAddress = f.timelock.Hex()
	}
	return input
}

// TestMigrateLockReleasePoolLiquidity_V1_2_0OffRampRescue migrates a timelock-owned v1.2.0 pool, whose
// liquidity is credited to no provider, through the offRamp rescue: register the timelock as an
// offRamp, releaseOrMint the amount to it, deregister it, then the sink funds the lockbox. Proposal
// generation must leave the pool untouched, the batch calldata must match exactly, and the lane ramp
// must survive. A timelock that already is an offRamp is neither registered nor deregistered.
func TestMigrateLockReleasePoolLiquidity_V1_2_0OffRampRescue(t *testing.T) {
	disabled := lock_release_token_pool_v120.RateLimiterConfig{IsEnabled: false, Capacity: big.NewInt(0), Rate: big.NewInt(0)}
	bps5000 := uint16(5000)
	bps10000 := uint16(10000)

	cases := []struct {
		name string
		// preRegistered, when set, registers the timelock as an offRamp with this limiter first.
		preRegistered  *lock_release_token_pool_v120.RateLimiterConfig
		input          tokens_core.MigrateLockReleasePoolLiquidityInput
		expectedAmount *big.Int
	}{
		{
			name:           "BasisPoints",
			input:          tokens_core.MigrateLockReleasePoolLiquidityInput{BasisPoints: &bps5000},
			expectedAmount: big.NewInt(5000),
		},
		{
			name:           "ExactAmount",
			input:          tokens_core.MigrateLockReleasePoolLiquidityInput{Amount: big.NewInt(3000)},
			expectedAmount: big.NewInt(3000),
		},
		{
			name:           "FullBalance",
			input:          tokens_core.MigrateLockReleasePoolLiquidityInput{BasisPoints: &bps10000},
			expectedAmount: big.NewInt(10000),
		},
		{
			name:           "TimelockAlreadyOffRamp/DisabledLimiter",
			preRegistered:  &disabled,
			input:          tokens_core.MigrateLockReleasePoolLiquidityInput{BasisPoints: &bps5000},
			expectedAmount: big.NewInt(5000),
		},
		{
			name: "TimelockAlreadyOffRamp/EnabledLimiterWithCapacity",
			preRegistered: &lock_release_token_pool_v120.RateLimiterConfig{
				IsEnabled: true, Capacity: big.NewInt(10000), Rate: big.NewInt(1),
			},
			input:          tokens_core.MigrateLockReleasePoolLiquidityInput{BasisPoints: &bps10000},
			expectedAmount: big.NewInt(10000),
		},
	}

	poolABI, err := lock_release_token_pool_v120.LockReleaseTokenPoolMetaData.GetAbi()
	require.NoError(t, err)

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			totalLiquidity := big.NewInt(10000)
			f := setupV1_2_0Fixture(t, oldPoolV1_2_0, totalLiquidity)
			if tc.preRegistered != nil {
				f.registerTimelockOffRamp(t, *tc.preRegistered)
			}
			f.transferOwnershipToTimelock(t)
			onRampsBefore, offRampsBefore := f.ramps(t)

			report, err := migrateLockReleasePoolLiquidity(testsetup.BundleWithFreshReporter(f.env.OperationsBundle), f.env.BlockChains, f.input(tc.input))
			require.NoError(t, err)

			// Proposal generation must not touch the old pool.
			onRamps, offRamps := f.ramps(t)
			require.Equal(t, onRampsBefore, onRamps)
			require.Equal(t, offRampsBefore, offRamps)
			require.Equal(t, 0, totalLiquidity.Cmp(f.balanceOf(t, f.oldPoolAddr)), "old pool must keep its liquidity before the proposal executes")

			// One batch: the source writes on the old pool first, then the sink's funding.
			require.Len(t, report.BatchOps, 1)
			txs := report.BatchOps[0].Transactions

			rampUpdate := func(allowed bool) []byte {
				data, err := poolABI.Pack("applyRampUpdates", []lock_release_token_pool_v120.TokenPoolRampUpdate{},
					[]lock_release_token_pool_v120.TokenPoolRampUpdate{{Ramp: f.timelock, Allowed: allowed, RateLimiterConfig: disabled}})
				require.NoError(t, err)
				return data
			}
			release, err := poolABI.Pack("releaseOrMint", []byte{}, f.timelock, tc.expectedAmount, uint64(0), []byte{})
			require.NoError(t, err)
			expectedSource := [][]byte{release}
			if tc.preRegistered == nil {
				expectedSource = [][]byte{rampUpdate(true), release, rampUpdate(false)}
			}
			require.Greater(t, len(txs), len(expectedSource), "the sink's funding must follow the source writes")
			for i, want := range expectedSource {
				require.Equal(t, f.oldPoolAddr, common.HexToAddress(txs[i].To), "tx %d must target the old pool", i)
				require.Equal(t, want, txs[i].Data, "tx %d calldata", i)
			}
			for i, sinkTx := range txs[len(expectedSource):] {
				to := common.HexToAddress(sinkTx.To)
				require.True(t, to == f.tokenAddr || to == f.lockBoxAddr,
					"sink tx %d must target the token or the lockbox, got %s", i, to)
			}
			require.Equal(t, f.lockBoxAddr, common.HexToAddress(txs[len(txs)-1].To), "the batch must end with the lockbox deposit")

			// Execute the proposal as the timelock.
			executeBatchOpsAs(t, f.chain, f.timelockKey, report.BatchOps)

			remaining := new(big.Int).Sub(totalLiquidity, tc.expectedAmount)
			require.Equal(t, 0, remaining.Cmp(f.balanceOf(t, f.oldPoolAddr)), "old pool should keep the rest of its liquidity")
			require.Equal(t, 0, tc.expectedAmount.Cmp(f.balanceOf(t, f.lockBoxAddr)), "lockbox should hold the migrated liquidity")
			require.Equal(t, 0, big.NewInt(0).Cmp(f.balanceOf(t, f.timelock)), "no liquidity may be left in the timelock")

			// The ramp set is back to what it was: the timelock is deregistered unless it already
			// was an offRamp, and the lane ramp and its limiter are untouched.
			onRamps, offRamps = f.ramps(t)
			require.Equal(t, onRampsBefore, onRamps)
			require.ElementsMatch(t, offRampsBefore, offRamps)
			isOffRamp, err := f.pool.IsOffRamp(&bind.CallOpts{Context: t.Context()}, f.timelock)
			require.NoError(t, err)
			require.Equal(t, tc.preRegistered != nil, isOffRamp, "timelock offRamp registration must be as before the migration")
			laneBucket, err := f.pool.CurrentOffRampRateLimiterState(&bind.CallOpts{Context: t.Context()}, laneRamp)
			require.NoError(t, err)
			require.True(t, laneBucket.IsEnabled)
			require.Equal(t, 0, big.NewInt(1_000_000).Cmp(laneBucket.Capacity), "lane limiter must be untouched")
		})
	}
}

// TestMigrateLockReleasePoolLiquidity_V1_2_0Preconditions checks every guard of the v1.2.0 source. Each
// must fail before emitting anything, leaving the pool's liquidity and ramps untouched.
func TestMigrateLockReleasePoolLiquidity_V1_2_0Preconditions(t *testing.T) {
	bps10000 := uint16(10000)

	cases := []struct {
		name string
		kind oldPoolKind
		// prepare runs while the deployer still owns the pool.
		prepare       func(t *testing.T, f v120Fixture)
		skipOwnership bool
		input         tokens_core.MigrateLockReleasePoolLiquidityInput
		expectedErr   string
	}{
		{
			name:        "AcceptsLiquidity",
			kind:        oldPoolV1_2_0AcceptingLiquidity,
			input:       tokens_core.MigrateLockReleasePoolLiquidityInput{BasisPoints: &bps10000},
			expectedErr: "accepts liquidity, so provider credits may exist",
		},
		{
			name: "ARMCursed",
			kind: oldPoolV1_2_0,
			prepare: func(t *testing.T, f v120Fixture) {
				callOpts := &bind.CallOpts{Context: t.Context()}
				armProxy, err := f.pool.GetArmProxy(callOpts)
				require.NoError(t, err)
				proxy, err := rmn_proxy_latest.NewRMNProxy(armProxy, f.chain.Client)
				require.NoError(t, err)
				rmnAddr, err := proxy.GetARM(callOpts)
				require.NoError(t, err)
				rmn, err := rmn_latest.NewRMN(rmnAddr, f.chain.Client)
				require.NoError(t, err)
				globalCurseSubject := [16]byte{0x01, 15: 0x01}
				tx, err := rmn.Curse(f.chain.DeployerKey, globalCurseSubject)
				require.NoError(t, err)
				_, err = f.chain.Confirm(tx)
				require.NoError(t, err)
			},
			input:       tokens_core.MigrateLockReleasePoolLiquidityInput{BasisPoints: &bps10000},
			expectedErr: "reports a curse",
		},
		{
			name:          "NotOwnedByTimelock",
			kind:          oldPoolV1_2_0,
			skipOwnership: true,
			input:         tokens_core.MigrateLockReleasePoolLiquidityInput{BasisPoints: &bps10000},
			expectedErr:   "not the timelock",
		},
		{
			name: "TimelockOffRampLimiterTooLow",
			kind: oldPoolV1_2_0,
			prepare: func(t *testing.T, f v120Fixture) {
				f.registerTimelockOffRamp(t, lock_release_token_pool_v120.RateLimiterConfig{
					IsEnabled: true, Capacity: big.NewInt(100), Rate: big.NewInt(1),
				})
			},
			input:       tokens_core.MigrateLockReleasePoolLiquidityInput{BasisPoints: &bps10000},
			expectedErr: "enabled rate limiter holding 100 tokens, less than the migration amount 10000",
		},
		{
			name: "ExactSiloMode",
			kind: oldPoolV1_2_0,
			input: tokens_core.MigrateLockReleasePoolLiquidityInput{
				SiloExactAmounts: []tokens_core.SiloExactAmount{{ChainSelector: 3379446385462418246, Amount: big.NewInt(100)}},
			},
			expectedErr: "only supported for siloed pools",
		},
		{
			name:        "AmountExceedsBalance",
			kind:        oldPoolV1_2_0,
			input:       tokens_core.MigrateLockReleasePoolLiquidityInput{Amount: big.NewInt(10001)},
			expectedErr: "exceeds old pool balance",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			totalLiquidity := big.NewInt(10000)
			f := setupV1_2_0Fixture(t, tc.kind, totalLiquidity)
			if tc.prepare != nil {
				tc.prepare(t, f)
			}
			if !tc.skipOwnership {
				f.transferOwnershipToTimelock(t)
			}
			onRampsBefore, offRampsBefore := f.ramps(t)

			require.ErrorContains(t, f.migrate(tc.input), tc.expectedErr)

			require.Equal(t, 0, totalLiquidity.Cmp(f.balanceOf(t, f.oldPoolAddr)), "old pool must keep its full balance")
			onRamps, offRamps := f.ramps(t)
			require.Equal(t, onRampsBefore, onRamps)
			require.Equal(t, offRampsBefore, offRamps)
		})
	}
}

// TestMigrateLockReleasePoolLiquidity_V1_2_0ReleaseIsProposalOnly checks that the v1.2.0 releaseOrMint
// cannot be sent by anyone but a registered offRamp, which is why the source forces it into the
// proposal: sent directly by the deployer (the pool owner) it reverts.
func TestMigrateLockReleasePoolLiquidity_V1_2_0ReleaseIsProposalOnly(t *testing.T) {
	f := setupV1_2_0Fixture(t, oldPoolV1_2_0, big.NewInt(10000))

	_, err := f.pool.ReleaseOrMint(f.chain.DeployerKey, nil, f.deployer, big.NewInt(1), 0, nil)
	require.Error(t, err, "the owner is not an offRamp and must not be able to release")
}
