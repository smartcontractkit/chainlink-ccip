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

	_, isMigrator := any(adapter).(tokensapi.TokenPoolMigrator)
	_, isRateLimitReader := any(adapter).(tokensapi.RateLimitReaderAdapter)
	require.True(t, isMigrator, "adapter must implement TokenPoolMigrator")
	require.True(t, isRateLimitReader, "adapter must implement RateLimitReaderAdapter")

	reg := tokensapi.GetTokenAdapterRegistry()
	resolved, ok := reg.GetTokenAdapter(chain_selectors.FamilyEVM, utils.Version_1_4_0)
	require.True(t, ok, "init() should have registered the v1.4.0 adapter")
	require.IsType(t, &TokenAdapter{}, resolved)
}

func TestPoolOpsV140_Version(t *testing.T) {
	t.Parallel()

	require.Equal(t, semver.MustParse("1.4.0").String(), (&poolOpsV140{}).Version().String())
}

func TestPoolOpsV140_SetDynamicPoolConfigs_RejectsFeeAdmin(t *testing.T) {
	t.Parallel()

	feeAdmin := common.HexToAddress("0x00000000000000000000000000000000000000ff")
	_, err := (&poolOpsV140{}).SetDynamicPoolConfigs(
		cldf_ops.Bundle{}, testChain, common.HexToAddress("0x1"), nil, nil, &feeAdmin,
	)
	require.ErrorContains(t, err, "fee admin is not supported on v1.4.0 token pools")
}

func TestPoolOpsV140_SetDynamicPoolConfigs_RejectsRouter(t *testing.T) {
	t.Parallel()

	router := common.HexToAddress("0x00000000000000000000000000000000000000aa")
	_, err := (&poolOpsV140{}).SetDynamicPoolConfigs(
		cldf_ops.Bundle{}, testChain, common.HexToAddress("0x1"), &router, nil, nil,
	)
	require.ErrorContains(t, err, "refusing to set the router on v1.4.0 token pool")
}

func TestTokenAdapter_DoesNotDeployPools(t *testing.T) {
	t.Parallel()

	require.Nil(t, NewTokenAdapter().DeployTokenPoolSeq)
}

func TestPoolOpsV140_SetDynamicPoolConfigs_NoopWhenNothingRequested(t *testing.T) {
	t.Parallel()

	writes, err := (&poolOpsV140{}).SetDynamicPoolConfigs(
		cldf_ops.Bundle{}, testChain, common.HexToAddress("0x1"), nil, nil, nil,
	)
	require.NoError(t, err)
	require.Empty(t, writes)
}

func TestPoolOpsV140_GetCurrentRateLimits_RejectsFastFinality(t *testing.T) {
	t.Parallel()

	_, err := (&poolOpsV140{}).GetCurrentRateLimits(
		cldf_ops.Bundle{}, testChain, common.HexToAddress("0x1"), 1234, true,
	)
	require.ErrorContains(t, err, "fast finality buckets are not supported")
}

func TestPoolOpsV140_RemoveRemotePools_Unsupported(t *testing.T) {
	t.Parallel()

	_, err := (&poolOpsV140{}).RemoveRemotePools(
		cldf_ops.Bundle{}, testChain, common.HexToAddress("0x1"),
		[]tokensapi.RemotePoolToRemove{{Selector: 1234}},
	)
	require.ErrorContains(t, err, "removing individual remote pools is not supported on v1.4.0 token pools")
}
