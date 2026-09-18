package adapters

import (
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	chain_selectors "github.com/smartcontractkit/chain-selectors"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"

	tokensapi "github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils"
)

var testChain = evm.Chain{Selector: chain_selectors.ETHEREUM_MAINNET.Selector}

func TestTokenAdapter_RegistersWithRequiredInterfaces(t *testing.T) {
	t.Parallel()

	adapter := NewTokenAdapter()

	// TokenPoolMigrator without RateLimitReaderAdapter is a hard error inside
	// ConfigureTokensForTransfers' auto-migrate discovery, so assert the pairing explicitly.
	_, isMigrator := any(adapter).(tokensapi.TokenPoolMigrator)
	_, isRateLimitReader := any(adapter).(tokensapi.RateLimitReaderAdapter)
	require.True(t, isMigrator, "adapter must implement TokenPoolMigrator")
	require.True(t, isRateLimitReader, "adapter must implement RateLimitReaderAdapter")

	reg := tokensapi.GetTokenAdapterRegistry()
	resolved, ok := reg.GetTokenAdapter(chain_selectors.FamilyEVM, utils.Version_1_5_0)
	require.True(t, ok, "init() should have registered the v1.5.0 adapter")
	require.IsType(t, &TokenAdapter{}, resolved)
}

func TestPoolOpsV150_Version(t *testing.T) {
	t.Parallel()

	require.Equal(t, semver.MustParse("1.5.0").String(), (&poolOpsV150{}).Version().String())
}

// TestPoolOpsV150_SetDynamicPoolConfigs_RejectsFeeAdmin locks in that the pre-2.0 "no fee admin"
// limitation surfaces as an error rather than being silently dropped.
func TestPoolOpsV150_SetDynamicPoolConfigs_RejectsFeeAdmin(t *testing.T) {
	t.Parallel()

	feeAdmin := common.HexToAddress("0x00000000000000000000000000000000000000ff")
	_, err := (&poolOpsV150{}).SetDynamicPoolConfigs(
		cldf_ops.Bundle{}, testChain, common.HexToAddress("0x1"), nil, nil, &feeAdmin,
	)
	require.ErrorContains(t, err, "fee admin is not supported on v1.5.0 token pools")
}

func TestPoolOpsV150_SetDynamicPoolConfigs_NoopWhenNothingRequested(t *testing.T) {
	t.Parallel()

	writes, err := (&poolOpsV150{}).SetDynamicPoolConfigs(
		cldf_ops.Bundle{}, testChain, common.HexToAddress("0x1"), nil, nil, nil,
	)
	require.NoError(t, err)
	require.Empty(t, writes)
}

// TestPoolOpsV150_GetCurrentRateLimits_RejectsFastFinality guards the fast-finality bucket, which
// does not exist before v2.0. Returning zero-value limits instead would silently misreport them.
func TestPoolOpsV150_GetCurrentRateLimits_RejectsFastFinality(t *testing.T) {
	t.Parallel()

	_, err := (&poolOpsV150{}).GetCurrentRateLimits(
		cldf_ops.Bundle{}, testChain, common.HexToAddress("0x1"), 1234, true,
	)
	require.ErrorContains(t, err, "fast finality buckets are not supported")
}

// TestPoolOpsV150_RemoveRemotePools_Unsupported pins the behaviour chosen for v1.5.0: the contract
// has no removeRemotePool, and applyChainUpdates(allowed=false) would drop the entire remote chain
// config. Erroring is deliberate - a silent no-op would let a "remove remote pools" pipeline
// report success while changing nothing.
func TestPoolOpsV150_RemoveRemotePools_Unsupported(t *testing.T) {
	t.Parallel()

	_, err := (&poolOpsV150{}).RemoveRemotePools(
		cldf_ops.Bundle{}, testChain, common.HexToAddress("0x1"),
		[]tokensapi.RemotePoolToRemove{{Selector: 1234}},
	)
	require.ErrorContains(t, err, "removing individual remote pools is not supported on v1.5.0 token pools")
}
