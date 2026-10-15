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

var (
	testChain = evm.Chain{Selector: chain_selectors.ETHEREUM_MAINNET.Selector}
	testPool  = common.HexToAddress("0x1")
)

func TestTokenAdapter_RegistersWithRequiredInterfaces(t *testing.T) {
	t.Parallel()

	adapter := NewTokenAdapter()

	_, isMigrator := any(adapter).(tokensapi.TokenPoolMigrator)
	_, isRateLimitReader := any(adapter).(tokensapi.RateLimitReaderAdapter)
	require.True(t, isMigrator, "adapter must implement TokenPoolMigrator")
	require.True(t, isRateLimitReader, "adapter must implement RateLimitReaderAdapter")

	reg := tokensapi.GetTokenAdapterRegistry()
	resolved, ok := reg.GetTokenAdapter(chain_selectors.FamilyEVM, utils.Version_1_2_0)
	require.True(t, ok, "init() should have registered the v1.2.0 adapter")
	require.IsType(t, &TokenAdapter{}, resolved)
}

func TestPoolOpsV120_Version(t *testing.T) {
	t.Parallel()

	require.Equal(t, semver.MustParse("1.2.0").String(), (&poolOpsV120{}).Version().String())
}

func TestTokenAdapter_DoesNotDeployPools(t *testing.T) {
	t.Parallel()

	require.Nil(t, NewTokenAdapter().DeployTokenPoolSeq)
}

func TestPoolOpsV120_SetRateLimiterConfig_Unsupported(t *testing.T) {
	t.Parallel()

	_, err := (&poolOpsV120{}).SetRateLimiterConfig(cldf_ops.Bundle{}, testChain, testPool, tokensapi.TPRLRemotes{})
	require.ErrorContains(t, err, "setting rate limits is not supported on v1.2.0 token pools")
	require.ErrorContains(t, err, "shared across every lane")
}

func TestPoolOpsV120_SetDynamicPoolConfigs(t *testing.T) {
	t.Parallel()

	addr := common.HexToAddress("0x00000000000000000000000000000000000000aa")

	tests := []struct {
		name                      string
		router, rlAdmin, feeAdmin *common.Address
		wantErr                   string
	}{
		{
			name:    "Failure - router rejected, pool has no setRouter",
			router:  &addr,
			wantErr: "router is not settable on v1.2.0 token pools",
		},
		{
			name:    "Failure - rate limit admin rejected, concept does not exist",
			rlAdmin: &addr,
			wantErr: "rate limit admin is not supported on v1.2.0 token pools",
		},
		{
			name:     "Failure - fee admin rejected, concept does not exist",
			feeAdmin: &addr,
			wantErr:  "fee admin is not supported on v1.2.0 token pools",
		},
		{
			name: "Success - nothing requested is a no-op",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			writes, err := (&poolOpsV120{}).SetDynamicPoolConfigs(
				cldf_ops.Bundle{}, testChain, testPool, tc.router, tc.rlAdmin, tc.feeAdmin,
			)
			if tc.wantErr != "" {
				require.ErrorContains(t, err, tc.wantErr)
				return
			}
			require.NoError(t, err)
			require.Empty(t, writes)
		})
	}
}

func TestPoolOpsV120_RemoveRemotePools_Unsupported(t *testing.T) {
	t.Parallel()

	_, err := (&poolOpsV120{}).RemoveRemotePools(
		cldf_ops.Bundle{}, testChain, testPool, []tokensapi.RemotePoolToRemove{{Selector: 1234}},
	)
	require.ErrorContains(t, err, "removing individual remote pools is not supported on v1.2.0 token pools")
}

func TestPoolOpsV120_GetCurrentRateLimits_Unsupported(t *testing.T) {
	t.Parallel()

	_, err := (&poolOpsV120{}).GetCurrentRateLimits(cldf_ops.Bundle{}, testChain, testPool, 1234, false)
	require.ErrorContains(t, err, "per-lane rate limits cannot be read from a v1.2.0 token pool")
}
