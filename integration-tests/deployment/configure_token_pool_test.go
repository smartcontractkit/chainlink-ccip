package deployment

import (
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/stretchr/testify/require"
	// sigs.k8s.io/yaml decodes YAML by converting to JSON and using encoding/json, which
	// honors the json `,string` selector tags. This mirrors how the migrations framework
	// ultimately populates changeset inputs (YAML node -> JSON -> struct). Plain gopkg.in/yaml.v3
	// panics on the style-guide-mandated `,string` tag, so it cannot be used here.
	"sigs.k8s.io/yaml"

	evm_datastore_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/datastore"
	evmadapters "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/adapters"
	bnmERC20ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/burn_mint_erc20"
	bnmOpsV2_0_0 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/burn_mint_token_pool"
	evm_testsetup "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/testsetup"
	tokenpoolV2_0_0 "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v2_0_0/token_pool"
	deployapi "github.com/smartcontractkit/chainlink-ccip/deployment/deploy"
	"github.com/smartcontractkit/chainlink-ccip/deployment/finality"
	"github.com/smartcontractkit/chainlink-ccip/deployment/testhelpers"
	tokensapi "github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	cciputils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
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
	return setupV2PoolsForConfigureImpl(t, tokenSymb, false)
}

// setupV2PoolsForConfigureMCMS is like setupV2PoolsForConfigure but deploys MCMS/timelock and
// transfers pool ownership to the timelock, so ConfigureTokenPool produces MCMS proposals.
func setupV2PoolsForConfigureMCMS(t *testing.T, tokenSymb string) configureTestEnv {
	return setupV2PoolsForConfigureImpl(t, tokenSymb, true)
}

func setupV2PoolsForConfigureImpl(t *testing.T, tokenSymb string, useMCMS bool) configureTestEnv {
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

	expansionMCMS := mcms.Input{}
	if useMCMS {
		// Register the EVM MCMS deployer/reader (idempotent) and deploy MCMS + timelock so the
		// pool can be timelock-owned and ConfigureTokenPool can emit timelock proposals.
		deployapi.GetRegistry().RegisterDeployer(chainsel.FamilyEVM, deployapi.MCMSVersion, &evmadapters.EVMDeployer{})
		changesets.GetRegistry().RegisterMCMSReader(chainsel.FamilyEVM, &evmadapters.EVMMCMSReader{})
		for _, sel := range []uint64{selA, selB} {
			DeployMCMS(t, e, sel, []string{cciputils.CLLQualifier})
		}
		expansionMCMS = NewDefaultInputForMCMS("configure token pool test setup")
	}

	disabledOutbound := tokensapi.RateLimiterConfigFloatInput{IsEnabled: false}
	deployerA := e.BlockChains.EVMChains()[selA].DeployerKey.From
	deployerB := e.BlockChains.EVMChains()[selB].DeployerKey.From

	expansionOut, err := tokensapi.TokenExpansion().Apply(*e, tokensapi.TokenExpansionInput{
		ChainAdapterVersion: cciputils.Version_2_0_0,
		MCMS:                expansionMCMS,
		TokenExpansionInputPerChain: map[uint64]tokensapi.TokenExpansionInputPerChain{
			selA: {
				SkipOwnershipTransfer: !useMCMS,
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
				SkipOwnershipTransfer: !useMCMS,
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
	if useMCMS {
		// Execute the ownership-transfer proposals so the pools become timelock-owned.
		testhelpers.ProcessTimelockProposals(t, *e, expansionOut.MCMSTimelockProposals, false)
	}

	fltrA := datastore.AddressRef{ChainSelector: selA, Type: datastore.ContractType(bnmOpsV2_0_0.ContractType), Version: bnmOpsV2_0_0.Version}
	poolA, err := datastore_utils.FindAndFormatRef(e.DataStore, fltrA, selA, evm_datastore_utils.ToEVMAddress)
	require.NoError(t, err)
	fltrB := datastore.AddressRef{ChainSelector: selB, Type: datastore.ContractType(bnmOpsV2_0_0.ContractType), Version: bnmOpsV2_0_0.Version}
	poolB, err := datastore_utils.FindAndFormatRef(e.DataStore, fltrB, selB, evm_datastore_utils.ToEVMAddress)
	require.NoError(t, err)

	// The simulated backend under-estimates gas for setRateLimitConfig updates that rewrite
	// existing bucket storage (SSTORE-refund accounting), causing an out-of-gas revert on the
	// exact estimate. Force a manual gas limit to bypass estimation — matches real-chain
	// behavior (production RPCs return a correct estimate / apply a buffer). Same fix the
	// SetTokenPoolRateLimits tests use via forceSimGasLimit.
	forceSimGasLimit(e, 5_000_000)

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
	refreshBundle(&tc)
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

// refreshBundle swaps in a fresh operations reporter, mirroring production where every migration
// run applies a changeset against a fresh bundle. Idempotency re-applies must use this so they
// exercise the changeset's read-compare-skip against live on-chain state instead of replaying a
// memoized sequence report from the prior apply (operations.ExecuteSequence reuses prior successful
// reports on identical input, which would mask a redundant write/proposal).
func refreshBundle(tc *configureTestEnv) {
	tc.env.OperationsBundle = evm_testsetup.BundleWithFreshReporter(tc.env.OperationsBundle)
}

func TestConfigureTokenPool_Admins(t *testing.T) {
	tc := setupV2PoolsForConfigure(t, "CTP_ADM")

	newRateLimitAdmin := "0x1111111111111111111111111111111111111111"
	newFeeAdmin := "0x2222222222222222222222222222222222222222"

	// Set only the rate limit admin; feeAdmin and router must be untouched.
	pool, err := tokenpoolV2_0_0.NewTokenPool(tc.poolA, tc.clientA)
	require.NoError(t, err)
	preCfg, err := pool.GetDynamicConfig(&bind.CallOpts{Context: t.Context()})
	require.NoError(t, err)

	input := tokensapi.ConfigureTokenPoolInput{
		MCMS: mcms.Input{},
		Chains: []tokensapi.ConfigureTokenPoolPerChain{{
			ChainSelector: tc.selA,
			Pools: []tokensapi.PoolConfigUpdate{{
				TokenPoolRef:   datastore.AddressRef{Address: tc.poolA.Hex()},
				RateLimitAdmin: ptrTo(newRateLimitAdmin),
			}},
		}},
	}
	require.NoError(t, tokensapi.ConfigureTokenPool().VerifyPreconditions(*tc.env, input))
	_, err = tokensapi.ConfigureTokenPool().Apply(*tc.env, input)
	require.NoError(t, err)

	postCfg, err := pool.GetDynamicConfig(&bind.CallOpts{Context: t.Context()})
	require.NoError(t, err)
	require.Equal(t, common.HexToAddress(newRateLimitAdmin), postCfg.RateLimitAdmin)
	require.Equal(t, preCfg.FeeAdmin, postCfg.FeeAdmin, "feeAdmin must be preserved")
	require.Equal(t, preCfg.Router, postCfg.Router, "router must be preserved")

	// Now set only the fee admin; the rate limit admin we just set must survive.
	input.Chains[0].Pools[0] = tokensapi.PoolConfigUpdate{
		TokenPoolRef: datastore.AddressRef{Address: tc.poolA.Hex()},
		FeeAdmin:     ptrTo(newFeeAdmin),
	}
	_, err = tokensapi.ConfigureTokenPool().Apply(*tc.env, input)
	require.NoError(t, err)
	postCfg, err = pool.GetDynamicConfig(&bind.CallOpts{Context: t.Context()})
	require.NoError(t, err)
	require.Equal(t, common.HexToAddress(newFeeAdmin), postCfg.FeeAdmin)
	require.Equal(t, common.HexToAddress(newRateLimitAdmin), postCfg.RateLimitAdmin, "rateLimitAdmin must be preserved")

	// Idempotency: setting both to their current values sends no transactions.
	input.Chains[0].Pools[0] = tokensapi.PoolConfigUpdate{
		TokenPoolRef:   datastore.AddressRef{Address: tc.poolA.Hex()},
		RateLimitAdmin: ptrTo(newRateLimitAdmin),
		FeeAdmin:       ptrTo(newFeeAdmin),
	}
	refreshBundle(&tc)
	before := currentBlock(t, tc, tc.selA)
	_, err = tokensapi.ConfigureTokenPool().Apply(*tc.env, input)
	require.NoError(t, err)
	after := currentBlock(t, tc, tc.selA)
	require.Equal(t, before, after, "no-op admin update must not send transactions")
}

func TestConfigureTokenPool_FeeConfig(t *testing.T) {
	tc := setupV2PoolsForConfigure(t, "CTP_FEE")

	// First apply: enable a fee config on the A→B lane. On-chain config starts disabled
	// (all-zero), so every field we care about must be provided explicitly here.
	initial := &tokensapi.PartialTokenTransferFeeConfig{
		IsEnabled:                  cciputils.NewOptional(true),
		DestBytesOverhead:          cciputils.NewOptional(uint32(320)),
		DestGasOverhead:            cciputils.NewOptional(uint32(21_000)),
		DefaultFinalityFeeUSDCents: cciputils.NewOptional(uint32(50)),
	}
	input := tokensapi.ConfigureTokenPoolInput{
		MCMS: mcms.Input{},
		Chains: []tokensapi.ConfigureTokenPoolPerChain{{
			ChainSelector: tc.selA,
			Pools: []tokensapi.PoolConfigUpdate{{
				TokenPoolRef: datastore.AddressRef{Address: tc.poolA.Hex()},
				Remotes: []tokensapi.RemoteConfigUpdate{{
					RemoteChainSelector:    tc.selB,
					TokenTransferFeeConfig: initial,
				}},
			}},
		}},
	}
	require.NoError(t, tokensapi.ConfigureTokenPool().VerifyPreconditions(*tc.env, input))
	_, err := tokensapi.ConfigureTokenPool().Apply(*tc.env, input)
	require.NoError(t, err)

	pool, err := tokenpoolV2_0_0.NewTokenPool(tc.poolA, tc.clientA)
	require.NoError(t, err)
	opts := &bind.CallOpts{Context: t.Context()}
	cfg, err := pool.GetTokenTransferFeeConfig(opts, common.Address{}, tc.selB, [4]byte{}, nil)
	require.NoError(t, err)
	require.True(t, cfg.IsEnabled)
	require.Equal(t, uint32(320), cfg.DestBytesOverhead)
	require.Equal(t, uint32(21_000), cfg.DestGasOverhead)
	require.Equal(t, uint32(50), cfg.FinalityFeeUSDCents)

	// Second apply: change ONLY destGasOverhead. All other fields must be preserved
	// from on-chain state (merge-with-on-chain, not merge-with-defaults).
	input.Chains[0].Pools[0].Remotes[0].TokenTransferFeeConfig = &tokensapi.PartialTokenTransferFeeConfig{
		DestGasOverhead: cciputils.NewOptional(uint32(31_000)),
	}
	_, err = tokensapi.ConfigureTokenPool().Apply(*tc.env, input)
	require.NoError(t, err)
	cfg, err = pool.GetTokenTransferFeeConfig(opts, common.Address{}, tc.selB, [4]byte{}, nil)
	require.NoError(t, err)
	require.Equal(t, uint32(31_000), cfg.DestGasOverhead, "targeted field must change")
	require.True(t, cfg.IsEnabled, "isEnabled must be preserved from on-chain")
	require.Equal(t, uint32(320), cfg.DestBytesOverhead, "destBytesOverhead must be preserved from on-chain")
	require.Equal(t, uint32(50), cfg.FinalityFeeUSDCents, "minFee must be preserved from on-chain")

	// Idempotency: identical partial re-apply sends no transactions.
	refreshBundle(&tc)
	before := currentBlock(t, tc, tc.selA)
	_, err = tokensapi.ConfigureTokenPool().Apply(*tc.env, input)
	require.NoError(t, err)
	after := currentBlock(t, tc, tc.selA)
	require.Equal(t, before, after, "no-op fee update must not send transactions")
}

func TestConfigureTokenPool_RateLimits(t *testing.T) {
	tc := setupV2PoolsForConfigure(t, "CTP_RL")

	// Unidirectional: configure only pool A's view of the A→B lane. Pool B gets no entry.
	buckets := []tokensapi.RateLimitBucketInput{
		{
			FastFinality: false,
			Outbound:     tokensapi.RateLimiterConfigFloatInput{IsEnabled: true, Capacity: 1000, Rate: 100},
			Inbound:      tokensapi.RateLimiterConfigFloatInput{IsEnabled: true, Capacity: 2000, Rate: 200},
		},
		{
			FastFinality: true,
			Outbound:     tokensapi.RateLimiterConfigFloatInput{IsEnabled: true, Capacity: 500, Rate: 50},
			Inbound:      tokensapi.RateLimiterConfigFloatInput{IsEnabled: true, Capacity: 600, Rate: 60},
		},
	}
	input := tokensapi.ConfigureTokenPoolInput{
		MCMS: mcms.Input{},
		Chains: []tokensapi.ConfigureTokenPoolPerChain{{
			ChainSelector: tc.selA,
			Pools: []tokensapi.PoolConfigUpdate{{
				TokenPoolRef: datastore.AddressRef{Address: tc.poolA.Hex()},
				Remotes: []tokensapi.RemoteConfigUpdate{{
					RemoteChainSelector: tc.selB,
					RateLimits:          buckets,
				}},
			}},
		}},
	}
	require.NoError(t, tokensapi.ConfigureTokenPool().VerifyPreconditions(*tc.env, input))
	_, err := tokensapi.ConfigureTokenPool().Apply(*tc.env, input)
	require.NoError(t, err)

	// Pool A: outbound scaled by local decimals (no bump), inbound scaled with +10% bump.
	defaultA, ffA := getRateLimits(t, cciputils.Version_2_0_0, tc.poolA, tc.clientA, tc.selB)
	validateScaledTPRLBucket(t, "default poolA", tc.decimalsA, defaultA, buckets[0].Outbound, buckets[0].Inbound)
	validateScaledTPRLBucket(t, "fast_finality poolA", tc.decimalsA, ffA, buckets[1].Outbound, buckets[1].Inbound)

	// Pool B: untouched (unidirectional update).
	defaultB, ffB := getRateLimits(t, cciputils.Version_2_0_0, tc.poolB, tc.clientB, tc.selA)
	require.False(t, defaultB.OutboundRateLimiterConfig.IsEnabled, "pool B default outbound must remain disabled")
	require.False(t, ffB.OutboundRateLimiterConfig.IsEnabled, "pool B FF outbound must remain disabled")

	// Idempotency: identical second apply sends no transactions.
	refreshBundle(&tc)
	before := currentBlock(t, tc, tc.selA)
	_, err = tokensapi.ConfigureTokenPool().Apply(*tc.env, input)
	require.NoError(t, err)
	after := currentBlock(t, tc, tc.selA)
	require.Equal(t, before, after, "no-op rate limit update must not send transactions")

	// Partial change: modify only the default bucket; the FF bucket write must be skipped
	// (single-bucket setRateLimitConfig call) and FF state preserved.
	input.Chains[0].Pools[0].Remotes[0].RateLimits = []tokensapi.RateLimitBucketInput{
		{
			FastFinality: false,
			Outbound:     tokensapi.RateLimiterConfigFloatInput{IsEnabled: true, Capacity: 3000, Rate: 300},
			Inbound:      buckets[0].Inbound,
		},
	}
	_, err = tokensapi.ConfigureTokenPool().Apply(*tc.env, input)
	require.NoError(t, err)
	defaultA2, ffA2 := getRateLimits(t, cciputils.Version_2_0_0, tc.poolA, tc.clientA, tc.selB)
	validateScaledTPRLBucket(t, "default poolA after partial", tc.decimalsA, defaultA2,
		tokensapi.RateLimiterConfigFloatInput{IsEnabled: true, Capacity: 3000, Rate: 300}, buckets[0].Inbound)
	validateScaledTPRLBucket(t, "fast_finality poolA preserved", tc.decimalsA, ffA2, buckets[1].Outbound, buckets[1].Inbound)
}

func TestConfigureTokenPool_YAMLRoundTrip(t *testing.T) {
	const payload = `
mcms:
  qualifier: CLL
input:
  - selector: "909606746561742123"
    pools:
      - tokenPoolRef: { address: '0x1111111111111111111111111111111111111111' }
        finalityConfig: { waitForSafe: true, blockDepth: 1, waitForFinality: false }
        rateLimitAdmin: '0x1111111111111111111111111111111111111111'
        feeAdmin: '0x1111111111111111111111111111111111111111'
        remotes:
          - selector: "5548718428018410741"
            tokenTransferFeeConfig:
              destBytesOverhead: 320
              destGasOverhead: 21000
            rateLimits:
              - fastFinality: false
                outbound: { isEnabled: true, capacity: 1000, rate: 100 }
                inbound: { isEnabled: true, capacity: 1000, rate: 100 }
              - fastFinality: true
                outbound: { isEnabled: true, capacity: 1000, rate: 100 }
                inbound: { isEnabled: true, capacity: 1000, rate: 100 }
`
	var input tokensapi.ConfigureTokenPoolInput
	require.NoError(t, yaml.Unmarshal([]byte(payload), &input))

	require.Len(t, input.Chains, 1)
	require.Equal(t, chainsel.TEST_90000001.Selector, input.Chains[0].ChainSelector)
	require.Len(t, input.Chains[0].Pools, 1)
	pool := input.Chains[0].Pools[0]
	require.Equal(t, "0x1111111111111111111111111111111111111111", pool.TokenPoolRef.Address)
	require.NotNil(t, pool.FinalityConfig)
	require.True(t, pool.FinalityConfig.WaitForSafe)
	require.Equal(t, uint16(1), pool.FinalityConfig.BlockDepth)
	require.NotNil(t, pool.RateLimitAdmin)
	require.NotNil(t, pool.FeeAdmin)
	require.Len(t, pool.Remotes, 1)
	remote := pool.Remotes[0]
	require.Equal(t, chainsel.TEST_90000002.Selector, remote.RemoteChainSelector)
	require.NotNil(t, remote.TokenTransferFeeConfig)
	dbo, ok := remote.TokenTransferFeeConfig.DestBytesOverhead.Get()
	require.True(t, ok)
	require.Equal(t, uint32(320), dbo)
	_, ok = remote.TokenTransferFeeConfig.IsEnabled.Get()
	require.False(t, ok, "unset optional fields must stay unset after unmarshal")
	require.Len(t, remote.RateLimits, 2)
	require.False(t, remote.RateLimits[0].FastFinality)
	require.True(t, remote.RateLimits[1].FastFinality)
	require.Equal(t, float64(1000), remote.RateLimits[0].Outbound.Capacity)
}

func TestConfigureTokenPool_RejectsPreV2Pools(t *testing.T) {
	tc := setupV2PoolsForConfigure(t, "CTP_VER")

	// Explicit version override forces adapter selection for a pre-v2 version; verify must
	// reject it in PR#1 regardless of what is registered.
	input := tokensapi.ConfigureTokenPoolInput{
		MCMS: mcms.Input{},
		Chains: []tokensapi.ConfigureTokenPoolPerChain{{
			ChainSelector: tc.selA,
			Pools: []tokensapi.PoolConfigUpdate{{
				TokenPoolRef:   datastore.AddressRef{Address: tc.poolA.Hex()},
				Version:        semver.MustParse("1.6.1"),
				RateLimitAdmin: ptrTo("0x2222222222222222222222222222222222222222"),
			}},
		}},
	}
	err := tokensapi.ConfigureTokenPool().VerifyPreconditions(*tc.env, input)
	require.Error(t, err)
	require.Contains(t, err.Error(), "only v2.0.0+ pools are currently supported")
}

func TestConfigureTokenPool_CombinedUpdate(t *testing.T) {
	tc := setupV2PoolsForConfigure(t, "CTP_ALL")

	admin := "0x3333333333333333333333333333333333333333"
	newFinality := finality.Config{BlockDepth: 7}
	rl := []tokensapi.RateLimitBucketInput{{
		FastFinality: false,
		Outbound:     tokensapi.RateLimiterConfigFloatInput{IsEnabled: true, Capacity: 1000, Rate: 100},
		Inbound:      tokensapi.RateLimiterConfigFloatInput{IsEnabled: true, Capacity: 1000, Rate: 100},
	}}
	fee := &tokensapi.PartialTokenTransferFeeConfig{
		IsEnabled:         cciputils.NewOptional(true),
		DestBytesOverhead: cciputils.NewOptional(uint32(320)),
		DestGasOverhead:   cciputils.NewOptional(uint32(21_000)),
	}

	input := tokensapi.ConfigureTokenPoolInput{
		MCMS: mcms.Input{},
		Chains: []tokensapi.ConfigureTokenPoolPerChain{
			{
				ChainSelector: tc.selA,
				Pools: []tokensapi.PoolConfigUpdate{{
					TokenPoolRef:   datastore.AddressRef{Address: tc.poolA.Hex()},
					FinalityConfig: &newFinality,
					RateLimitAdmin: ptrTo(admin),
					FeeAdmin:       ptrTo(admin),
					Remotes: []tokensapi.RemoteConfigUpdate{{
						RemoteChainSelector:    tc.selB,
						TokenTransferFeeConfig: fee,
						RateLimits:             rl,
					}},
				}},
			},
			{
				ChainSelector: tc.selB,
				Pools: []tokensapi.PoolConfigUpdate{{
					TokenPoolRef: datastore.AddressRef{Address: tc.poolB.Hex()},
					Remotes: []tokensapi.RemoteConfigUpdate{{
						RemoteChainSelector: tc.selA,
						RateLimits:          rl,
					}},
				}},
			},
		},
	}
	require.NoError(t, tokensapi.ConfigureTokenPool().VerifyPreconditions(*tc.env, input))
	out, err := tokensapi.ConfigureTokenPool().Apply(*tc.env, input)
	require.NoError(t, err)
	require.Empty(t, out.MCMSTimelockProposals, "EOA-owned pools must not produce MCMS proposals")

	// Spot-check each feature landed.
	validateFinalityConfigV2_0_0(t, tc.poolA, tc.clientA, newFinality)
	pool, err := tokenpoolV2_0_0.NewTokenPool(tc.poolA, tc.clientA)
	require.NoError(t, err)
	dynCfg, err := pool.GetDynamicConfig(&bind.CallOpts{Context: t.Context()})
	require.NoError(t, err)
	require.Equal(t, common.HexToAddress(admin), dynCfg.RateLimitAdmin)
	require.Equal(t, common.HexToAddress(admin), dynCfg.FeeAdmin)
	defaultA, _ := getRateLimits(t, cciputils.Version_2_0_0, tc.poolA, tc.clientA, tc.selB)
	validateScaledTPRLBucket(t, "combined default poolA", tc.decimalsA, defaultA, rl[0].Outbound, rl[0].Inbound)
	defaultB, _ := getRateLimits(t, cciputils.Version_2_0_0, tc.poolB, tc.clientB, tc.selA)
	validateScaledTPRLBucket(t, "combined default poolB", tc.decimalsB, defaultB, rl[0].Outbound, rl[0].Inbound)

	// Full idempotency across every feature at once.
	refreshBundle(&tc)
	beforeA := currentBlock(t, tc, tc.selA)
	beforeB := currentBlock(t, tc, tc.selB)
	out, err = tokensapi.ConfigureTokenPool().Apply(*tc.env, input)
	require.NoError(t, err)
	require.Empty(t, out.MCMSTimelockProposals, "combined no-op must not emit proposals")
	afterA := currentBlock(t, tc, tc.selA)
	afterB := currentBlock(t, tc, tc.selB)
	require.Equal(t, beforeA, afterA, "combined no-op must not send transactions on chain A")
	require.Equal(t, beforeB, afterB, "combined no-op must not send transactions on chain B")
}

func TestConfigureTokenPool_MCMSOwnedPool(t *testing.T) {
	tc := setupV2PoolsForConfigureMCMS(t, "CTP_MCMS")

	newCfg := finality.Config{BlockDepth: 9}
	input := tokensapi.ConfigureTokenPoolInput{
		MCMS: NewDefaultInputForMCMS("configure token pool finality"),
		Chains: []tokensapi.ConfigureTokenPoolPerChain{{
			ChainSelector: tc.selA,
			Pools: []tokensapi.PoolConfigUpdate{{
				TokenPoolRef:   datastore.AddressRef{Address: tc.poolA.Hex()},
				FinalityConfig: &newCfg,
			}},
		}},
	}
	require.NoError(t, tokensapi.ConfigureTokenPool().VerifyPreconditions(*tc.env, input))
	out, err := tokensapi.ConfigureTokenPool().Apply(*tc.env, input)
	require.NoError(t, err)

	// Timelock-owned pool: the change is proposed, not executed.
	require.Len(t, out.MCMSTimelockProposals, 1, "timelock-owned pool must produce an MCMS proposal")
	validateFinalityConfigV2_0_0(t, tc.poolA, tc.clientA, finality.Config{WaitForFinality: true}) // unchanged until execution

	// Execute the proposal and confirm the change lands.
	testhelpers.ProcessTimelockProposals(t, *tc.env, out.MCMSTimelockProposals, false)
	validateFinalityConfigV2_0_0(t, tc.poolA, tc.clientA, newCfg)

	// Idempotency (the design's core guarantee, asserted on the output not just block height):
	// re-applying the now-satisfied config against a fresh bundle must emit ZERO proposals, so
	// no redundant timelock op / MCMS predecessor conflict is produced.
	refreshBundle(&tc)
	out2, err := tokensapi.ConfigureTokenPool().Apply(*tc.env, input)
	require.NoError(t, err)
	require.Empty(t, out2.MCMSTimelockProposals, "no-op re-apply must not emit an MCMS proposal")
}
