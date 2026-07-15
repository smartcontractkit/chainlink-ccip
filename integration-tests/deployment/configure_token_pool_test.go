package deployment

import (
	"testing"

	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/deployment/finality"
	tokensapi "github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	cciputils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"

	_ "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/adapters"
)

func TestConfigureTokenPool_VerifyPreconditions(t *testing.T) {
	sel := chainsel.TEST_90000001.Selector
	dst := chainsel.TEST_90000002.Selector
	env, err := environment.New(t.Context(), environment.WithEVMSimulated(t, []uint64{sel}))
	require.NoError(t, err)

	cs := tokensapi.ConfigureTokenPool()
	poolRef := datastore.AddressRef{Address: "0x1111111111111111111111111111111111111111"}
	enabled := tokensapi.RateLimiterConfigFloatInput{IsEnabled: true, Capacity: 1000, Rate: 100}

	singlePoolInput := func(pool tokensapi.PoolConfigUpdate) tokensapi.ConfigureTokenPoolInput {
		return tokensapi.ConfigureTokenPoolInput{
			MCMS: mcms.Input{},
			Chains: []tokensapi.ConfigureTokenPoolPerChain{
				{ChainSelector: sel, Pools: []tokensapi.PoolConfigUpdate{pool}},
			},
		}
	}

	cases := []struct {
		name   string
		input  tokensapi.ConfigureTokenPoolInput
		errors []string
	}{
		{
			name:   "rejects_empty_input",
			input:  tokensapi.ConfigureTokenPoolInput{MCMS: mcms.Input{}},
			errors: []string{"at least one chain entry"},
		},
		{
			name: "rejects_chain_entry_with_no_pools",
			input: tokensapi.ConfigureTokenPoolInput{
				MCMS:   mcms.Input{},
				Chains: []tokensapi.ConfigureTokenPoolPerChain{{ChainSelector: sel}},
			},
			errors: []string{"no pools provided"},
		},
		{
			name:   "rejects_empty_pool_update",
			input:  singlePoolInput(tokensapi.PoolConfigUpdate{TokenPoolRef: poolRef}),
			errors: []string{"no fields to update"},
		},
		{
			name: "rejects_empty_remote_update",
			input: singlePoolInput(tokensapi.PoolConfigUpdate{
				TokenPoolRef: poolRef,
				Remotes:      []tokensapi.RemoteConfigUpdate{{RemoteChainSelector: dst}},
			}),
			errors: []string{"nothing to update"},
		},
		{
			name: "rejects_remote_equal_to_local_chain",
			input: singlePoolInput(tokensapi.PoolConfigUpdate{
				TokenPoolRef: poolRef,
				Remotes: []tokensapi.RemoteConfigUpdate{{
					RemoteChainSelector: sel,
					RateLimits:          []tokensapi.RateLimitBucketInput{{Outbound: enabled, Inbound: enabled}},
				}},
			}),
			errors: []string{"must not equal the pool's own chain selector"},
		},
		{
			name: "rejects_duplicate_remote_selectors",
			input: singlePoolInput(tokensapi.PoolConfigUpdate{
				TokenPoolRef: poolRef,
				Remotes: []tokensapi.RemoteConfigUpdate{
					{RemoteChainSelector: dst, RateLimits: []tokensapi.RateLimitBucketInput{{Outbound: enabled, Inbound: enabled}}},
					{RemoteChainSelector: dst, RateLimits: []tokensapi.RateLimitBucketInput{{Outbound: enabled, Inbound: enabled}}},
				},
			}),
			errors: []string{"duplicate remote chain selector"},
		},
		{
			name: "rejects_duplicate_pool_entries",
			input: tokensapi.ConfigureTokenPoolInput{
				MCMS: mcms.Input{},
				Chains: []tokensapi.ConfigureTokenPoolPerChain{{
					ChainSelector: sel,
					Pools: []tokensapi.PoolConfigUpdate{
						{TokenPoolRef: poolRef, RateLimitAdmin: ptrTo("0x2222222222222222222222222222222222222222")},
						{TokenPoolRef: poolRef, FeeAdmin: ptrTo("0x2222222222222222222222222222222222222222")},
					},
				}},
			},
			errors: []string{"duplicate pool entry"},
		},
		{
			name: "rejects_three_rate_limit_buckets",
			input: singlePoolInput(tokensapi.PoolConfigUpdate{
				TokenPoolRef: poolRef,
				Remotes: []tokensapi.RemoteConfigUpdate{{
					RemoteChainSelector: dst,
					RateLimits: []tokensapi.RateLimitBucketInput{
						{FastFinality: false, Outbound: enabled, Inbound: enabled},
						{FastFinality: true, Outbound: enabled, Inbound: enabled},
						{FastFinality: false, Outbound: enabled, Inbound: enabled},
					},
				}},
			}),
			errors: []string{"at most two rate limit buckets allowed"},
		},
		{
			name: "rejects_duplicate_default_buckets",
			input: singlePoolInput(tokensapi.PoolConfigUpdate{
				TokenPoolRef: poolRef,
				Remotes: []tokensapi.RemoteConfigUpdate{{
					RemoteChainSelector: dst,
					RateLimits: []tokensapi.RateLimitBucketInput{
						{FastFinality: false, Outbound: enabled, Inbound: enabled},
						{FastFinality: false, Outbound: enabled, Inbound: enabled},
					},
				}},
			}),
			errors: []string{"multiple rate limit buckets with fastFinality=false"},
		},
		{
			name: "rejects_rate_greater_than_capacity",
			input: singlePoolInput(tokensapi.PoolConfigUpdate{
				TokenPoolRef: poolRef,
				Remotes: []tokensapi.RemoteConfigUpdate{{
					RemoteChainSelector: dst,
					RateLimits: []tokensapi.RateLimitBucketInput{{
						Outbound: tokensapi.RateLimiterConfigFloatInput{IsEnabled: true, Capacity: 10, Rate: 100},
						Inbound:  enabled,
					}},
				}},
			}),
			errors: []string{"rate greater than capacity"},
		},
		{
			name: "rejects_disabled_with_nonzero_values",
			input: singlePoolInput(tokensapi.PoolConfigUpdate{
				TokenPoolRef: poolRef,
				Remotes: []tokensapi.RemoteConfigUpdate{{
					RemoteChainSelector: dst,
					RateLimits: []tokensapi.RateLimitBucketInput{{
						Outbound: enabled,
						Inbound:  tokensapi.RateLimiterConfigFloatInput{IsEnabled: false, Capacity: 5},
					}},
				}},
			}),
			errors: []string{"disabled but capacity or rate is non-zero"},
		},
		{
			name: "rejects_zero_finality_config",
			input: singlePoolInput(tokensapi.PoolConfigUpdate{
				TokenPoolRef:   poolRef,
				FinalityConfig: &finality.Config{},
			}),
			errors: []string{"finality config is empty"},
		},
		{
			name: "rejects_invalid_rate_limit_admin_address",
			input: singlePoolInput(tokensapi.PoolConfigUpdate{
				TokenPoolRef:   poolRef,
				RateLimitAdmin: ptrTo("not-an-address"),
			}),
			errors: []string{"invalid rateLimitAdmin address"},
		},
		{
			name: "rejects_low_dest_bytes_overhead",
			input: singlePoolInput(tokensapi.PoolConfigUpdate{
				TokenPoolRef: poolRef,
				Remotes: []tokensapi.RemoteConfigUpdate{{
					RemoteChainSelector:    dst,
					TokenTransferFeeConfig: feeCfgWithDestBytesOverhead(16),
				}},
			}),
			errors: []string{"destBytesOverhead must be at least 32"},
		},
		{
			name: "rejects_unresolvable_pool_ref",
			input: singlePoolInput(tokensapi.PoolConfigUpdate{
				TokenPoolRef:   poolRef,
				RateLimitAdmin: ptrTo("0x2222222222222222222222222222222222222222"),
			}),
			errors: []string{"failed to resolve token pool ref"},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			err := cs.VerifyPreconditions(*env, tt.input)
			require.Error(t, err)
			for _, substr := range tt.errors {
				require.Contains(t, err.Error(), substr)
			}
		})
	}
}

func ptrTo[T any](v T) *T { return &v }

func feeCfgWithDestBytesOverhead(v uint32) *tokensapi.PartialTokenTransferFeeConfig {
	return &tokensapi.PartialTokenTransferFeeConfig{DestBytesOverhead: cciputils.NewOptional(v)}
}
