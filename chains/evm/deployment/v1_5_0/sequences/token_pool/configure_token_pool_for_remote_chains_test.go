package token_pool

import (
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	chain_selectors "github.com/smartcontractkit/chain-selectors"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"

	bmtpap "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/burn_mint_token_pool_and_proxy"
	tokenapi "github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
)

var remoteChainSelector = chain_selectors.AVALANCHE_MAINNET.Selector

// TestConfigureTokenPoolForRemoteChains_FreshChain covers the "chain not yet supported" path:
// a single applyChainUpdates entry with allowed=true.
func TestConfigureTokenPoolForRemoteChains_FreshChain(t *testing.T) {
	t.Parallel()

	e, poolAddr := deployConfigurablePool(t)
	chain := e.BlockChains.EVMChains()[testChainSelector]
	pool := mustPool(t, e, poolAddr)

	remoteToken := common.HexToAddress("0x00000000000000000000000000000000000000aa")
	remotePool := common.HexToAddress("0x00000000000000000000000000000000000000bb")

	configure(t, e, chain.Selector, poolAddr, remoteToken, remotePool, 100, 10)

	supported, err := pool.GetSupportedChains(&bind.CallOpts{})
	require.NoError(t, err)
	require.Equal(t, []uint64{remoteChainSelector}, supported)

	onChainRemoteToken, err := pool.GetRemoteToken(&bind.CallOpts{}, remoteChainSelector)
	require.NoError(t, err)
	require.Equal(t, remoteToken.Bytes(), onChainRemoteToken)

	// Remote pools are stored 32-byte left-padded so the value matches what the protocol sends.
	onChainRemotePool, err := pool.GetRemotePool(&bind.CallOpts{}, remoteChainSelector)
	require.NoError(t, err)
	require.Equal(t, common.LeftPadBytes(remotePool.Bytes(), 32), onChainRemotePool)

	outbound, err := pool.GetCurrentOutboundRateLimiterState(&bind.CallOpts{}, remoteChainSelector)
	require.NoError(t, err)
	require.True(t, outbound.IsEnabled)
	require.Positive(t, outbound.Capacity.Sign())
}

// TestConfigureTokenPoolForRemoteChains_Idempotent asserts a second identical run is a no-op.
//
// The assertion is on chain head rather than on the returned BatchOps: when the caller owns the
// pool the writes execute directly, so a BatchOperation with zero transactions comes back whether
// or not anything happened. The re-run uses a fresh operations bundle so the sequence-level report
// cache cannot mask a write that really would have been emitted.
func TestConfigureTokenPoolForRemoteChains_Idempotent(t *testing.T) {
	t.Parallel()

	e, poolAddr := deployConfigurablePool(t)
	chain := e.BlockChains.EVMChains()[testChainSelector]

	remoteToken := common.HexToAddress("0x00000000000000000000000000000000000000aa")
	remotePool := common.HexToAddress("0x00000000000000000000000000000000000000bb")

	configure(t, e, chain.Selector, poolAddr, remoteToken, remotePool, 100, 10)
	before, err := chain.Client.HeaderByNumber(t.Context(), nil)
	require.NoError(t, err)

	e.OperationsBundle = freshBundle(e.OperationsBundle)
	configure(t, e, chain.Selector, poolAddr, remoteToken, remotePool, 100, 10)

	after, err := chain.Client.HeaderByNumber(t.Context(), nil)
	require.NoError(t, err)
	require.Equal(t, before.Number.String(), after.Number.String(),
		"re-running with identical config should not send a transaction")
}

// TestConfigureTokenPoolForRemoteChains_RemoteTokenChanged is the load-bearing case for v1.5.0.
//
// v1.5.1 removes an existing remote chain with a separate `remoteChainSelectorsToRemove`
// argument. v1.5.0 has no such argument: the removal must be expressed as an allowed=false entry
// preceding the allowed=true entry in the SAME ChainUpdate array, relying on applyChainUpdates
// processing the array in order. If that assumption were wrong the call would revert with
// ChainAlreadyExists, so this test is what proves it.
func TestConfigureTokenPoolForRemoteChains_RemoteTokenChanged(t *testing.T) {
	t.Parallel()

	e, poolAddr := deployConfigurablePool(t)
	chain := e.BlockChains.EVMChains()[testChainSelector]
	pool := mustPool(t, e, poolAddr)

	firstToken := common.HexToAddress("0x00000000000000000000000000000000000000aa")
	secondToken := common.HexToAddress("0x00000000000000000000000000000000000000cc")
	remotePool := common.HexToAddress("0x00000000000000000000000000000000000000bb")

	configure(t, e, chain.Selector, poolAddr, firstToken, remotePool, 100, 10)
	configure(t, e, chain.Selector, poolAddr, secondToken, remotePool, 100, 10)

	onChainRemoteToken, err := pool.GetRemoteToken(&bind.CallOpts{}, remoteChainSelector)
	require.NoError(t, err)
	require.Equal(t, secondToken.Bytes(), onChainRemoteToken, "remote token should have been retargeted")

	supported, err := pool.GetSupportedChains(&bind.CallOpts{})
	require.NoError(t, err)
	require.Equal(t, []uint64{remoteChainSelector}, supported, "chain should appear exactly once after remove+add")
}

// TestConfigureTokenPoolForRemoteChains_RemotePoolReplaced covers the setRemotePool path. A
// v1.5.0 pool holds ONE remote pool per lane, so retargeting replaces rather than appends.
func TestConfigureTokenPoolForRemoteChains_RemotePoolReplaced(t *testing.T) {
	t.Parallel()

	e, poolAddr := deployConfigurablePool(t)
	chain := e.BlockChains.EVMChains()[testChainSelector]
	pool := mustPool(t, e, poolAddr)

	remoteToken := common.HexToAddress("0x00000000000000000000000000000000000000aa")
	firstPool := common.HexToAddress("0x00000000000000000000000000000000000000bb")
	secondPool := common.HexToAddress("0x00000000000000000000000000000000000000dd")

	configure(t, e, chain.Selector, poolAddr, remoteToken, firstPool, 100, 10)
	configure(t, e, chain.Selector, poolAddr, remoteToken, secondPool, 100, 10)

	onChainRemotePool, err := pool.GetRemotePool(&bind.CallOpts{}, remoteChainSelector)
	require.NoError(t, err)
	require.Equal(t, common.LeftPadBytes(secondPool.Bytes(), 32), onChainRemotePool,
		"v1.5.0 replaces the single remote pool rather than appending")
}

// TestConfigureTokenPoolForRemoteChains_RateLimitsUpdated covers the setChainRateLimiterConfig
// path taken when only the rate limits differ.
func TestConfigureTokenPoolForRemoteChains_RateLimitsUpdated(t *testing.T) {
	t.Parallel()

	e, poolAddr := deployConfigurablePool(t)
	chain := e.BlockChains.EVMChains()[testChainSelector]
	pool := mustPool(t, e, poolAddr)

	remoteToken := common.HexToAddress("0x00000000000000000000000000000000000000aa")
	remotePool := common.HexToAddress("0x00000000000000000000000000000000000000bb")

	configure(t, e, chain.Selector, poolAddr, remoteToken, remotePool, 100, 10)
	before, err := pool.GetCurrentOutboundRateLimiterState(&bind.CallOpts{}, remoteChainSelector)
	require.NoError(t, err)

	// Advance one block before changing the rate limits.
	//
	// applyChainUpdates above created the buckets with lastUpdated = that block's timestamp. If the
	// next setChainRateLimiterConfig is estimated while the pending block still carries that same
	// timestamp, RateLimiter's `timeDiff != 0` refill branch is skipped during estimation but runs
	// when the tx is actually mined a second later - and the extra SSTOREs push it past the
	// estimate, so the tx runs out of gas (status 0, gasUsed == gasLimit, no revert reason).
	//
	// This is a property of the shared pre-2.0 RateLimiter and the estimate-then-mine pattern, not
	// of this sequence: v1.5.1 and v1.6.x behave identically. It only bites deterministically on
	// the simulated backend, whose clock advances exactly one second per block.
	advanceOneBlock(t, chain)

	configure(t, e, chain.Selector, poolAddr, remoteToken, remotePool, 500, 50)
	after, err := pool.GetCurrentOutboundRateLimiterState(&bind.CallOpts{}, remoteChainSelector)
	require.NoError(t, err)

	require.Positive(t, after.Capacity.Cmp(before.Capacity), "capacity should have been raised")
	require.Positive(t, after.Rate.Cmp(before.Rate), "rate should have been raised")
}

// TestConfigureTokenPoolForRemoteChains_RejectsEnabledZeroRateLimit mirrors the contract-side
// InvalidRateLimitRate guard so the failure surfaces before a transaction is sent.
func TestConfigureTokenPoolForRemoteChains_RejectsEnabledZeroRateLimit(t *testing.T) {
	t.Parallel()

	e, poolAddr := deployConfigurablePool(t)
	chain := e.BlockChains.EVMChains()[testChainSelector]

	rl := &tokenapi.RateLimiterConfigFloatInput{IsEnabled: true, Capacity: 0, Rate: 0}
	_, err := cldf_ops.ExecuteSequence(e.OperationsBundle, ConfigureTokenPoolForRemoteChains, chain,
		ConfigureTokenPoolForRemoteChainsInput{
			ChainSelector:    testChainSelector,
			TokenPoolAddress: poolAddr,
			TokenPoolVersion: utils.Version_1_5_0,
			RemoteChains: map[uint64]tokenapi.RemoteChainConfig[[]byte, string]{
				remoteChainSelector: {
					RemoteToken:               common.HexToAddress("0x00000000000000000000000000000000000000aa").Bytes(),
					RemotePool:                common.HexToAddress("0x00000000000000000000000000000000000000bb").Bytes(),
					RemoteDecimals:            18,
					OutboundRateLimiterConfig: rl,
					InboundRateLimiterConfig:  rl,
				},
			},
		})
	require.ErrorContains(t, err, "rate and capacity are both zero")
}

// configure runs the sequence for a single remote chain and returns its output.
func configure(
	t *testing.T,
	e *cldf.Environment,
	chainSelector uint64,
	poolAddr, remoteToken, remotePool common.Address,
	capacity, rate float64,
) sequences.OnChainOutput {
	t.Helper()

	chain := e.BlockChains.EVMChains()[chainSelector]
	rl := &tokenapi.RateLimiterConfigFloatInput{IsEnabled: true, Capacity: capacity, Rate: rate}

	report, err := cldf_ops.ExecuteSequence(e.OperationsBundle, ConfigureTokenPoolForRemoteChains, chain,
		ConfigureTokenPoolForRemoteChainsInput{
			ChainSelector:    chainSelector,
			TokenPoolAddress: poolAddr,
			TokenPoolVersion: utils.Version_1_5_0,
			RemoteChains: map[uint64]tokenapi.RemoteChainConfig[[]byte, string]{
				remoteChainSelector: {
					RemoteToken:               remoteToken.Bytes(),
					RemotePool:                remotePool.Bytes(),
					RemoteDecimals:            18,
					OutboundRateLimiterConfig: rl,
					InboundRateLimiterConfig:  rl,
				},
			},
		})
	require.NoError(t, err)

	return report.Output
}

func deployConfigurablePool(t *testing.T) (*cldf.Environment, common.Address) {
	t.Helper()

	e, tokenRef, _, _ := setupDeployEnv(t)
	report, err := cldf_ops.ExecuteSequence(e.OperationsBundle, DeployTokenPool, e.BlockChains, tokenapi.DeployTokenPoolInput{
		TokenRef:          &tokenRef,
		PoolType:          string(bmtpap.ContractType),
		TokenPoolVersion:  utils.Version_1_5_0,
		ChainSelector:     testChainSelector,
		ExistingDataStore: e.DataStore,
	})
	require.NoError(t, err)
	require.Len(t, report.Output.Addresses, 1)

	return e, common.HexToAddress(report.Output.Addresses[0].Address)
}

func mustPool(t *testing.T, e *cldf.Environment, poolAddr common.Address) *bmtpap.BurnMintTokenPoolAndProxyContract {
	t.Helper()

	chain := e.BlockChains.EVMChains()[testChainSelector]
	pool, err := bmtpap.NewBurnMintTokenPoolAndProxyContract(poolAddr, chain.Client)
	require.NoError(t, err)

	return pool
}

// freshBundle returns a bundle with an empty report cache, so a repeated sequence call actually
// re-executes instead of replaying the previous report.
func freshBundle(b cldf_ops.Bundle) cldf_ops.Bundle {
	return cldf_ops.NewBundle(b.GetContext, b.Logger, cldf_ops.NewMemoryReporter())
}

// advanceOneBlock mines a block by deploying a throwaway contract, so the next gas estimation sees
// a block timestamp later than any state written by the previous transaction.
func advanceOneBlock(t *testing.T, chain evm.Chain) {
	t.Helper()
	deployTestDripToken(t, chain, "FILLER")
}
