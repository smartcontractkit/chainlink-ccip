package deployment

import (
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/stretchr/testify/require"

	evm_datastore_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/datastore"
	bnmERC20ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/burn_mint_erc20"
	bnmOpsV2_0_0 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/burn_mint_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/deployment/finality"
	tokensapi "github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	cciputils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
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

type configureTestEnv struct {
	env          *cldf_deployment.Environment
	selA, selB   uint64
	poolA, poolB common.Address
	clientA      bind.ContractBackend
	clientB      bind.ContractBackend
	tokenSymb    string
	decimalsA    uint8
	decimalsB    uint8
}

// setupV2PoolsForConfigure deploys v2.0.0 chain contracts plus a connected v2 burn-mint
// token/pool pair on two simulated chains, owned by the deployer key (no MCMS).
func setupV2PoolsForConfigure(t *testing.T, tokenSymb string) configureTestEnv {
	t.Helper()
	const decimalsA = 18
	const decimalsB = 6

	selA := chainsel.TEST_90000001.Selector
	selB := chainsel.TEST_90000002.Selector
	e, err := environment.New(t.Context(), environment.WithEVMSimulated(t, []uint64{selA, selB}))
	require.NoError(t, err)

	cumulative := datastore.NewMemoryDataStore()
	DeployChainContractsV2_0_0(t, e, cumulative, selA)
	DeployChainContractsV2_0_0(t, e, cumulative, selB)
	e.DataStore = cumulative.Seal()

	disabledOutbound := tokensapi.RateLimiterConfigFloatInput{IsEnabled: false}
	deployerA := e.BlockChains.EVMChains()[selA].DeployerKey.From
	deployerB := e.BlockChains.EVMChains()[selB].DeployerKey.From

	expansionOut, err := tokensapi.TokenExpansion().Apply(*e, tokensapi.TokenExpansionInput{
		ChainAdapterVersion: cciputils.Version_2_0_0,
		MCMS:                mcms.Input{},
		TokenExpansionInputPerChain: map[uint64]tokensapi.TokenExpansionInputPerChain{
			selA: {
				SkipOwnershipTransfer: true,
				TokenPoolVersion:      bnmOpsV2_0_0.Version,
				DeployTokenInput: &tokensapi.DeployTokenInput{
					Name: tokenSymb, Symbol: tokenSymb, Decimals: decimalsA,
					ExternalAdmin: deployerA.Hex(), CCIPAdmin: deployerA.Hex(),
					Type: bnmERC20ops.ContractType,
				},
				DeployTokenPoolInput: &tokensapi.DeployTokenPoolInput{
					PoolType:              string(bnmOpsV2_0_0.ContractType),
					AllowedFinalityConfig: finality.Config{WaitForFinality: true},
				},
				TokenTransferConfig: &tokensapi.TokenTransferConfig{
					RemoteChains: map[uint64]tokensapi.RemoteChainConfig[*datastore.AddressRef, datastore.AddressRef]{
						selB: {OutboundRateLimiterConfig: &disabledOutbound},
					},
				},
			},
			selB: {
				SkipOwnershipTransfer: true,
				TokenPoolVersion:      bnmOpsV2_0_0.Version,
				DeployTokenInput: &tokensapi.DeployTokenInput{
					Name: tokenSymb, Symbol: tokenSymb, Decimals: decimalsB,
					ExternalAdmin: deployerB.Hex(), CCIPAdmin: deployerB.Hex(),
					Type: bnmERC20ops.ContractType,
				},
				DeployTokenPoolInput: &tokensapi.DeployTokenPoolInput{
					PoolType:              string(bnmOpsV2_0_0.ContractType),
					AllowedFinalityConfig: finality.Config{WaitForFinality: true},
				},
				TokenTransferConfig: &tokensapi.TokenTransferConfig{
					RemoteChains: map[uint64]tokensapi.RemoteChainConfig[*datastore.AddressRef, datastore.AddressRef]{
						selA: {OutboundRateLimiterConfig: &disabledOutbound},
					},
				},
			},
		},
	})
	require.NoError(t, err)
	MergeAddresses(t, e, expansionOut.DataStore)

	fltrA := datastore.AddressRef{ChainSelector: selA, Type: datastore.ContractType(bnmOpsV2_0_0.ContractType), Version: bnmOpsV2_0_0.Version}
	poolA, err := datastore_utils.FindAndFormatRef(e.DataStore, fltrA, selA, evm_datastore_utils.ToEVMAddress)
	require.NoError(t, err)
	fltrB := datastore.AddressRef{ChainSelector: selB, Type: datastore.ContractType(bnmOpsV2_0_0.ContractType), Version: bnmOpsV2_0_0.Version}
	poolB, err := datastore_utils.FindAndFormatRef(e.DataStore, fltrB, selB, evm_datastore_utils.ToEVMAddress)
	require.NoError(t, err)

	return configureTestEnv{
		env: e, selA: selA, selB: selB, poolA: poolA, poolB: poolB,
		clientA: e.BlockChains.EVMChains()[selA].Client, clientB: e.BlockChains.EVMChains()[selB].Client,
		tokenSymb: tokenSymb, decimalsA: decimalsA, decimalsB: decimalsB,
	}
}

func TestConfigureTokenPool_FinalityConfig(t *testing.T) {
	tc := setupV2PoolsForConfigure(t, "CTP_FIN")

	// Sanity: initial finality config comes from deployment.
	validateFinalityConfigV2_0_0(t, tc.poolA, tc.clientA, finality.Config{WaitForFinality: true})

	newCfg := finality.Config{WaitForSafe: true, BlockDepth: 5}
	input := tokensapi.ConfigureTokenPoolInput{
		MCMS: mcms.Input{},
		Chains: []tokensapi.ConfigureTokenPoolPerChain{{
			ChainSelector: tc.selA,
			Pools: []tokensapi.PoolConfigUpdate{{
				TokenPoolRef:   datastore.AddressRef{Address: tc.poolA.Hex()},
				FinalityConfig: &newCfg,
			}},
		}},
	}
	require.NoError(t, tokensapi.ConfigureTokenPool().VerifyPreconditions(*tc.env, input))
	_, err := tokensapi.ConfigureTokenPool().Apply(*tc.env, input)
	require.NoError(t, err)

	// The targeted pool changed; the untouched pool did not.
	validateFinalityConfigV2_0_0(t, tc.poolA, tc.clientA, newCfg)
	validateFinalityConfigV2_0_0(t, tc.poolB, tc.clientB, finality.Config{WaitForFinality: true})

	// Idempotency: identical second apply sends no transactions (no new blocks mined).
	before := currentBlock(t, tc, tc.selA)
	_, err = tokensapi.ConfigureTokenPool().Apply(*tc.env, input)
	require.NoError(t, err)
	after := currentBlock(t, tc, tc.selA)
	require.Equal(t, before, after, "second identical apply must not send transactions")
}

// currentBlock returns the latest block number on the given chain. The simulated backend
// mines exactly one block per transaction, so an unchanged block number across an Apply
// proves no transaction was sent.
func currentBlock(t *testing.T, tc configureTestEnv, sel uint64) uint64 {
	t.Helper()
	header, err := tc.env.BlockChains.EVMChains()[sel].Client.HeaderByNumber(t.Context(), nil)
	require.NoError(t, err)
	return header.Number.Uint64()
}
