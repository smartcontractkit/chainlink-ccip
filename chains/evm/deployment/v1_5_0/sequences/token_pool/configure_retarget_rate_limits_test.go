package token_pool

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"

	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"

	tokenapi "github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils"
)

// Two remote tokens that differ only in the decimals declared for them, which is what makes a
// carried-forward inbound bucket wrong.
var (
	remoteTokenSixDecimals      = common.HexToAddress("0x00000000000000000000000000000000000000a1")
	remoteTokenEighteenDecimals = common.HexToAddress("0x00000000000000000000000000000000000000a2")
)

// configureLane runs the sequence for one remote chain. A nil rl omits the rate limits entirely,
// which is what puts the sequence on the "reuse whatever is on-chain" path.
func configureLane(
	t *testing.T,
	e *cldf.Environment,
	poolAddr, remoteToken common.Address,
	remoteDecimals uint8,
	rl *tokenapi.RateLimiterConfigFloatInput,
) error {
	t.Helper()

	chain := e.BlockChains.EVMChains()[testChainSelector]
	cfg := tokenapi.RemoteChainConfig[[]byte, string]{
		RemoteToken:    remoteToken.Bytes(),
		RemotePool:     common.HexToAddress("0x00000000000000000000000000000000000000bb").Bytes(),
		RemoteDecimals: remoteDecimals,
	}
	if rl != nil {
		cfg.OutboundRateLimiterConfig = rl
		cfg.InboundRateLimiterConfig = rl
	}

	_, err := cldf_ops.ExecuteSequence(e.OperationsBundle, ConfigureTokenPoolForRemoteChains, chain,
		ConfigureTokenPoolForRemoteChainsInput{
			TokenPoolAddress: poolAddr,
			TokenPoolVersion: utils.Version_1_5_0,
			RemoteChains:     map[uint64]tokenapi.RemoteChainConfig[[]byte, string]{remoteChainSelector: cfg},
		})

	return err
}

// TestConfigureTokenPoolForRemoteChains_RejectsRemoteTokenRetargetWithEnabledLimits covers the case
// that made the "reuse the on-chain buckets" shortcut unsafe.
//
// On a pre-1.6.1 pool the inbound bucket is denominated in the REMOTE token's decimals, so its
// stored value only means what it says relative to the remote token configured at the time. Reusing
// it verbatim for a different remote token rescales the limit by 10^(oldDecimals-newDecimals) -
// here 10^12. Going 6 -> 18 decimals that tightens the lane to ~1.1e-10 tokens; going 18 -> 6 it
// loosens it by the same factor, which effectively removes the rate limit. The sequence cannot read
// the previous remote token's decimals to rebase, so it must refuse instead of guessing.
func TestConfigureTokenPoolForRemoteChains_RejectsRemoteTokenRetargetWithEnabledLimits(t *testing.T) {
	t.Parallel()

	e, poolAddr := deployConfigurablePool(t)
	pool := mustPool(t, e, poolAddr)

	// Establish the lane against a 6-decimal remote token with limits enabled.
	require.NoError(t, configureLane(t, e, poolAddr, remoteTokenSixDecimals, 6,
		&tokenapi.RateLimiterConfigFloatInput{IsEnabled: true, Capacity: 100, Rate: 10}))

	before, err := pool.GetCurrentInboundRateLimiterState(&bind.CallOpts{}, remoteChainSelector)
	require.NoError(t, err)
	// Inbound is the counterpart outbound +10%, scaled by the REMOTE token's 6 decimals.
	require.Equal(t, big.NewInt(110_000_000).String(), before.Capacity.String(),
		"inbound bucket should be denominated in the remote token's decimals")

	// Retargeting to an 18-decimal remote token with no limits supplied must fail rather than carry
	// the 6-decimal number over.
	err = configureLane(t, e, poolAddr, remoteTokenEighteenDecimals, 18, nil)
	require.ErrorContains(t, err, "cannot be reinterpreted for the new one")
	require.ErrorContains(t, err, "Specify the rate limits explicitly")

	// Nothing should have been written: the lane still points at the original remote token.
	gotRemoteToken, err := pool.GetRemoteToken(&bind.CallOpts{}, remoteChainSelector)
	require.NoError(t, err)
	require.Equal(t, remoteTokenSixDecimals.Bytes(), gotRemoteToken, "failed run must not retarget the lane")
	after, err := pool.GetCurrentInboundRateLimiterState(&bind.CallOpts{}, remoteChainSelector)
	require.NoError(t, err)
	require.Equal(t, before.Capacity.String(), after.Capacity.String(), "failed run must not change the rate limits")
}

// TestConfigureTokenPoolForRemoteChains_AllowsRemoteTokenRetargetWithDisabledLimits is the other
// half of the rule: a disabled bucket is all zeroes, so it carries no decimal denomination and is
// safe to reuse across a retarget. The check must not block that.
func TestConfigureTokenPoolForRemoteChains_AllowsRemoteTokenRetargetWithDisabledLimits(t *testing.T) {
	t.Parallel()

	e, poolAddr := deployConfigurablePool(t)
	pool := mustPool(t, e, poolAddr)

	require.NoError(t, configureLane(t, e, poolAddr, remoteTokenSixDecimals, 6,
		&tokenapi.RateLimiterConfigFloatInput{IsEnabled: false, Capacity: 0, Rate: 0}))

	require.NoError(t, configureLane(t, e, poolAddr, remoteTokenEighteenDecimals, 18, nil),
		"a disabled bucket is decimal-agnostic, so the retarget should be allowed")

	gotRemoteToken, err := pool.GetRemoteToken(&bind.CallOpts{}, remoteChainSelector)
	require.NoError(t, err)
	require.Equal(t, remoteTokenEighteenDecimals.Bytes(), gotRemoteToken, "lane should have been retargeted")

	inbound, err := pool.GetCurrentInboundRateLimiterState(&bind.CallOpts{}, remoteChainSelector)
	require.NoError(t, err)
	require.False(t, inbound.IsEnabled)
	require.Equal(t, "0", inbound.Capacity.String())
}

// TestConfigureTokenPoolForRemoteChains_RetargetWithExplicitLimitsRescales is the escape hatch the
// error message points operators at: supplying the limits explicitly lets the retarget proceed, and
// the new value is scaled for the NEW remote token rather than carried over.
func TestConfigureTokenPoolForRemoteChains_RetargetWithExplicitLimitsRescales(t *testing.T) {
	t.Parallel()

	e, poolAddr := deployConfigurablePool(t)
	pool := mustPool(t, e, poolAddr)

	require.NoError(t, configureLane(t, e, poolAddr, remoteTokenSixDecimals, 6,
		&tokenapi.RateLimiterConfigFloatInput{IsEnabled: true, Capacity: 100, Rate: 10}))

	require.NoError(t, configureLane(t, e, poolAddr, remoteTokenEighteenDecimals, 18,
		&tokenapi.RateLimiterConfigFloatInput{IsEnabled: true, Capacity: 100, Rate: 10}))

	gotRemoteToken, err := pool.GetRemoteToken(&bind.CallOpts{}, remoteChainSelector)
	require.NoError(t, err)
	require.Equal(t, remoteTokenEighteenDecimals.Bytes(), gotRemoteToken)

	inbound, err := pool.GetCurrentInboundRateLimiterState(&bind.CallOpts{}, remoteChainSelector)
	require.NoError(t, err)
	require.True(t, inbound.IsEnabled)

	// 110 tokens at the NEW remote token's 18 decimals, not the 110e6 the old bucket held. The x1.1
	// inbound premium is computed in float64, so allow a tiny tolerance well below one token.
	want := new(big.Int).Mul(big.NewInt(110), new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil))
	diff := new(big.Int).Abs(new(big.Int).Sub(want, inbound.Capacity))
	require.Negative(t, diff.Cmp(big.NewInt(1e12)),
		"inbound capacity should be rescaled to the new remote token's decimals: want ~%s, got %s", want, inbound.Capacity)
}
