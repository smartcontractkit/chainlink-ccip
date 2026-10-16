package tokens_test

import (
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"

	"github.com/smartcontractkit/chainlink-ccip/deployment/finality"
	"github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
)

func TestConfigureTokenPool_VerifyPreconditions(t *testing.T) {
	const (
		chainA = uint64(5009297550715157269)
		chainB = uint64(15971525489660198786)
	)

	basePool := func() tokens.PoolConfigUpdate {
		rateLimitAdmin := "0x0000000000000000000000000000000000000001"
		return tokens.PoolConfigUpdate{
			TokenPoolRef: datastore.AddressRef{
				Type:          "TokenPool",
				Version:       semver.MustParse("1.0.0"),
				ChainSelector: chainA,
				Qualifier:     "default",
			},
			RateLimitAdmin: &rateLimitAdmin,
		}
	}

	baseInput := func() tokens.ConfigureTokenPoolInput {
		return tokens.ConfigureTokenPoolInput{
			Chains: []tokens.ConfigureTokenPoolPerChain{
				{ChainSelector: chainA, Pools: []tokens.PoolConfigUpdate{basePool()}},
			},
		}
	}

	tests := []struct {
		name        string
		mutate      func(cfg *tokens.ConfigureTokenPoolInput)
		expectedErr string
	}{
		{
			name:   "Success - valid input",
			mutate: func(cfg *tokens.ConfigureTokenPoolInput) {},
		},
		{
			name: "Failure - no chains",
			mutate: func(cfg *tokens.ConfigureTokenPoolInput) {
				cfg.Chains = nil
			},
			expectedErr: "at least one chain entry",
		},
		{
			name: "Failure - no pools for chain",
			mutate: func(cfg *tokens.ConfigureTokenPoolInput) {
				cfg.Chains[0].Pools = nil
			},
			expectedErr: "no pools provided for chain selector",
		},
		{
			name: "Failure - empty TokenPoolRef",
			mutate: func(cfg *tokens.ConfigureTokenPoolInput) {
				cfg.Chains[0].Pools[0].TokenPoolRef = datastore.AddressRef{}
			},
			expectedErr: "empty tokenPoolRef",
		},
		{
			name: "Failure - TokenPoolRef chain selector mismatch",
			mutate: func(cfg *tokens.ConfigureTokenPoolInput) {
				cfg.Chains[0].Pools[0].TokenPoolRef.ChainSelector = chainB
			},
			expectedErr: "does not match the enclosing chain selector",
		},
		{
			name: "Failure - no fields to update",
			mutate: func(cfg *tokens.ConfigureTokenPoolInput) {
				cfg.Chains[0].Pools[0].RateLimitAdmin = nil
			},
			expectedErr: "has no fields to update",
		},
		{
			name: "Failure - empty rateLimitAdmin",
			mutate: func(cfg *tokens.ConfigureTokenPoolInput) {
				empty := ""
				cfg.Chains[0].Pools[0].RateLimitAdmin = &empty
			},
			expectedErr: "has an empty rateLimitAdmin",
		},
		{
			name: "Failure - empty feeAdmin",
			mutate: func(cfg *tokens.ConfigureTokenPoolInput) {
				empty := ""
				cfg.Chains[0].Pools[0].FeeAdmin = &empty
			},
			expectedErr: "has an empty feeAdmin",
		},
		{
			name: "Failure - empty router",
			mutate: func(cfg *tokens.ConfigureTokenPoolInput) {
				empty := ""
				cfg.Chains[0].Pools[0].Router = &empty
			},
			expectedErr: "has an empty router",
		},
		{
			name: "Failure - invalid finality config",
			mutate: func(cfg *tokens.ConfigureTokenPoolInput) {
				cfg.Chains[0].Pools[0].FinalityConfig = &finality.Config{}
			},
			expectedErr: "finality config",
		},
		{
			name: "Failure - remote chain selector equals own chain selector",
			mutate: func(cfg *tokens.ConfigureTokenPoolInput) {
				cfg.Chains[0].Pools[0].Remotes = []tokens.RemoteConfigUpdate{
					{RemoteChainSelector: chainA},
				}
			},
			expectedErr: "must not equal the pool's own chain selector",
		},
		{
			name: "Failure - duplicate remote chain selector",
			mutate: func(cfg *tokens.ConfigureTokenPoolInput) {
				cfg.Chains[0].Pools[0].Remotes = []tokens.RemoteConfigUpdate{
					{RemoteChainSelector: chainB},
					{RemoteChainSelector: chainB},
				}
			},
			expectedErr: "duplicate remote chain selector",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			cfg := baseInput()
			tc.mutate(&cfg)

			changeset := tokens.ConfigureTokenPool()
			err := changeset.VerifyPreconditions(deployment.Environment{}, cfg)

			if tc.expectedErr == "" {
				require.NoError(t, err)
				return
			}
			require.Error(t, err)
			require.Contains(t, err.Error(), tc.expectedErr)
		})
	}
}
