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

	lrtp "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_4_0/operations/lock_release_token_pool"
	tpbindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_4_0/token_pool"
	tokenapi "github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
)

var remoteChainSelector = chain_selectors.AVALANCHE_MAINNET.Selector

func TestConfigureTokenPoolForRemoteChains_FreshChain(t *testing.T) {
	t.Parallel()

	e, poolAddr := deployConfigurablePool(t)
	chain := e.BlockChains.EVMChains()[testChainSelector]
	pool := mustPool(t, e, poolAddr)

	configure(t, e, chain.Selector, poolAddr, 100, 10)

	supported, err := pool.GetSupportedChains(&bind.CallOpts{})
	require.NoError(t, err)
	require.Equal(t, []uint64{remoteChainSelector}, supported)

	outbound, err := pool.GetCurrentOutboundRateLimiterState(&bind.CallOpts{}, remoteChainSelector)
	require.NoError(t, err)
	require.True(t, outbound.IsEnabled)
	require.Positive(t, outbound.Capacity.Sign())

	inbound, err := pool.GetCurrentInboundRateLimiterState(&bind.CallOpts{}, remoteChainSelector)
	require.NoError(t, err)
	require.True(t, inbound.IsEnabled)
	require.Positive(t, inbound.Capacity.Sign())
}

func TestConfigureTokenPoolForRemoteChains_Idempotent(t *testing.T) {
	t.Parallel()

	e, poolAddr := deployConfigurablePool(t)
	chain := e.BlockChains.EVMChains()[testChainSelector]

	configure(t, e, chain.Selector, poolAddr, 100, 10)
	before, err := chain.Client.HeaderByNumber(t.Context(), nil)
	require.NoError(t, err)

	e.OperationsBundle = freshBundle(e.OperationsBundle)
	configure(t, e, chain.Selector, poolAddr, 100, 10)

	after, err := chain.Client.HeaderByNumber(t.Context(), nil)
	require.NoError(t, err)
	require.Equal(t, before.Number.String(), after.Number.String(),
		"re-running with identical config should not send a transaction")
}

func TestConfigureTokenPoolForRemoteChains_RateLimitUpdate(t *testing.T) {
	t.Parallel()

	e, poolAddr := deployConfigurablePool(t)
	chain := e.BlockChains.EVMChains()[testChainSelector]
	pool := mustPool(t, e, poolAddr)

	configure(t, e, chain.Selector, poolAddr, 100, 10)
	first, err := pool.GetCurrentOutboundRateLimiterState(&bind.CallOpts{}, remoteChainSelector)
	require.NoError(t, err)

	// The first configure writes the rate limiter bucket, which stamps the
	// current block timestamp. If the next setChainRateLimiterConfig is estimated
	// while the pending block still carries that same timestamp, RateLimiter's
	// `timeDiff != 0` refill branch is skipped during estimation but runs when the
	// tx is actually mined a second later - and the extra SSTOREs push it past the
	// estimate, so the tx runs out of gas (status 0, gasUsed == gasLimit, no revert
	// reason).
	//
	// This is a property of the shared pre-2.0 RateLimiter and the
	// estimate-then-mine pattern, not of this sequence: v1.5.0, v1.5.1 and v1.6.x
	// behave identically. It only bites deterministically on the simulated backend,
	// whose clock advances exactly one second per block.
	advanceOneBlock(t, chain)

	e.OperationsBundle = freshBundle(e.OperationsBundle)
	configure(t, e, chain.Selector, poolAddr, 500, 50)

	second, err := pool.GetCurrentOutboundRateLimiterState(&bind.CallOpts{}, remoteChainSelector)
	require.NoError(t, err)
	require.NotEqual(t, first.Capacity.String(), second.Capacity.String(),
		"capacity should have been updated")
	require.Equal(t, 1, second.Capacity.Cmp(first.Capacity), "capacity should have increased")

	// The chain must still be supported exactly once - an update must not add a
	// duplicate entry.
	supported, err := pool.GetSupportedChains(&bind.CallOpts{})
	require.NoError(t, err)
	require.Equal(t, []uint64{remoteChainSelector}, supported)
}

func TestConfigureTokenPoolForRemoteChains_LockReleasePool(t *testing.T) {
	t.Parallel()

	e, poolAddr := deployConfigurablePoolOfType(t, string(lrtp.ContractType))
	chain := e.BlockChains.EVMChains()[testChainSelector]
	pool := mustPool(t, e, poolAddr)

	configure(t, e, chain.Selector, poolAddr, 100, 10)

	supported, err := pool.GetSupportedChains(&bind.CallOpts{})
	require.NoError(t, err)
	require.Equal(t, []uint64{remoteChainSelector}, supported)

	outbound, err := pool.GetCurrentOutboundRateLimiterState(&bind.CallOpts{}, remoteChainSelector)
	require.NoError(t, err)
	require.True(t, outbound.IsEnabled)
}

func TestConfigureTokenPoolForRemoteChains_RejectsEnabledZeroRateLimit(t *testing.T) {
	t.Parallel()

	e, poolAddr := deployConfigurablePool(t)
	chain := e.BlockChains.EVMChains()[testChainSelector]

	rl := &tokenapi.RateLimiterConfigFloatInput{IsEnabled: true, Capacity: 0, Rate: 0}
	_, err := cldf_ops.ExecuteSequence(e.OperationsBundle, ConfigureTokenPoolForRemoteChains, chain,
		ConfigureTokenPoolForRemoteChainsInput{
			ChainSelector:    chain.Selector,
			TokenPoolAddress: poolAddr,
			TokenPoolVersion: utils.Version_1_4_0,
			RemoteChains: map[uint64]tokenapi.RemoteChainConfig[[]byte, string]{
				remoteChainSelector: {
					RemoteDecimals:            18,
					OutboundRateLimiterConfig: rl,
					InboundRateLimiterConfig:  rl,
				},
			},
		})
	require.ErrorContains(t, err, "rate limiter config is enabled but rate and capacity are both zero")
}

func TestConfigureTokenPoolForRemoteChains_RejectsPartialRateLimits(t *testing.T) {
	t.Parallel()

	e, poolAddr := deployConfigurablePool(t)
	chain := e.BlockChains.EVMChains()[testChainSelector]

	rl := &tokenapi.RateLimiterConfigFloatInput{IsEnabled: true, Capacity: 100, Rate: 10}
	_, err := cldf_ops.ExecuteSequence(e.OperationsBundle, ConfigureTokenPoolForRemoteChains, chain,
		ConfigureTokenPoolForRemoteChainsInput{
			ChainSelector:    chain.Selector,
			TokenPoolAddress: poolAddr,
			TokenPoolVersion: utils.Version_1_4_0,
			RemoteChains: map[uint64]tokenapi.RemoteChainConfig[[]byte, string]{
				remoteChainSelector: {
					RemoteDecimals:            18,
					OutboundRateLimiterConfig: rl,
				},
			},
		})
	require.ErrorContains(t, err, "must both be specified together or both omitted")
}

func configure(
	t *testing.T,
	e *cldf.Environment,
	chainSelector uint64,
	poolAddr common.Address,
	capacity, rate float64,
) sequences.OnChainOutput {
	t.Helper()

	chain := e.BlockChains.EVMChains()[chainSelector]
	rl := &tokenapi.RateLimiterConfigFloatInput{IsEnabled: true, Capacity: capacity, Rate: rate}

	report, err := cldf_ops.ExecuteSequence(e.OperationsBundle, ConfigureTokenPoolForRemoteChains, chain,
		ConfigureTokenPoolForRemoteChainsInput{
			ChainSelector:    chainSelector,
			TokenPoolAddress: poolAddr,
			TokenPoolVersion: utils.Version_1_4_0,
			RemoteChains: map[uint64]tokenapi.RemoteChainConfig[[]byte, string]{
				remoteChainSelector: {
					RemoteDecimals:            18,
					OutboundRateLimiterConfig: rl,
					InboundRateLimiterConfig:  rl,
				},
			},
		})
	require.NoError(t, err)

	return report.Output
}

func mustPool(t *testing.T, e *cldf.Environment, poolAddr common.Address) *tpbindings.TokenPool {
	t.Helper()

	chain := e.BlockChains.EVMChains()[testChainSelector]
	pool, err := tpbindings.NewTokenPool(poolAddr, chain.Client)
	require.NoError(t, err)

	return pool
}

// freshBundle returns a bundle with an empty report cache, so a repeated sequence
// call actually re-executes instead of replaying the previous report.
func freshBundle(b cldf_ops.Bundle) cldf_ops.Bundle {
	return cldf_ops.NewBundle(b.GetContext, b.Logger, cldf_ops.NewMemoryReporter())
}

// advanceOneBlock mines a block by deploying a throwaway contract, so the next gas
// estimation sees a block timestamp later than any state written by the previous
// transaction.
func advanceOneBlock(t *testing.T, chain evm.Chain) {
	t.Helper()
	deployTestDripToken(t, chain, "FILLER")
}
