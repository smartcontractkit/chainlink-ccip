package deployment

import (
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gagliardetto/solana-go"
	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/stretchr/testify/require"

	evm_datastore_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/datastore"
	evmadapters "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/adapters"
	bnmERC20ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/burn_mint_erc20"
	bnmOpsV2_0_0 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/burn_mint_token_pool"
	evm_testsetup "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/testsetup"
	tokenpoolV1_5_1 "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_5_1/token_pool"
	tokenpoolV1_6_1 "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_1/token_pool"
	tokenpoolV2_0_0 "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v2_0_0/token_pool"
	solanautils "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/utils"
	burnmint_token_pool "github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/v0_1_1/burnmint_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/tokens"
	deployapi "github.com/smartcontractkit/chainlink-ccip/deployment/deploy"
	"github.com/smartcontractkit/chainlink-ccip/deployment/fees"
	"github.com/smartcontractkit/chainlink-ccip/deployment/finality"
	"github.com/smartcontractkit/chainlink-ccip/deployment/testhelpers"
	tokensapi "github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	cciputils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	solchain "github.com/smartcontractkit/chainlink-deployments-framework/chain/solana"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"

	_ "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/adapters"
	_ "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_0/adapters"
	_ "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_0/sequences"
)

func TestConfigureTokenPool_VerifyPreconditions(t *testing.T) {
	sel := chainsel.TEST_90000001.Selector
	dst := chainsel.TEST_90000002.Selector
	env, err := environment.New(t.Context(), environment.WithEVMSimulated(t, []uint64{sel}))
	require.NoError(t, err)

	cs := tokensapi.ConfigureTokenPool()
	poolRef := datastore.AddressRef{Address: "0x1111111111111111111111111111111111111111"}

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
			name:   "rejects_empty_pool_ref",
			input:  singlePoolInput(tokensapi.PoolConfigUpdate{FeeAdmin: new("0x2222222222222222222222222222222222222222")}),
			errors: []string{"empty tokenPoolRef"},
		},
		{
			name: "rejects_mismatched_pool_ref_chain_selector",
			input: singlePoolInput(tokensapi.PoolConfigUpdate{
				TokenPoolRef: datastore.AddressRef{Address: poolRef.Address, ChainSelector: dst},
				FeeAdmin:     new("0x2222222222222222222222222222222222222222"),
			}),
			errors: []string{"does not match the enclosing chain selector"},
		},
		{
			name:   "rejects_empty_pool_update",
			input:  singlePoolInput(tokensapi.PoolConfigUpdate{TokenPoolRef: poolRef}),
			errors: []string{"no fields to update"},
		},
		{
			name: "rejects_all_unset_fee_config",
			input: singlePoolInput(tokensapi.PoolConfigUpdate{
				TokenPoolRef: poolRef,
				Remotes: []tokensapi.RemoteConfigUpdate{{
					RemoteChainSelector:    dst,
					TokenTransferFeeConfig: &tokensapi.PartialTokenTransferFeeConfig{},
				}},
			}),
			errors: []string{"nothing to update"},
		},
		{
			name: "rejects_remote_equal_to_local_chain",
			input: singlePoolInput(tokensapi.PoolConfigUpdate{
				TokenPoolRef: poolRef,
				Remotes: []tokensapi.RemoteConfigUpdate{{
					RemoteChainSelector:    sel,
					TokenTransferFeeConfig: &tokensapi.PartialTokenTransferFeeConfig{DestBytesOverhead: cciputils.NewOptional(uint32(320))},
				}},
			}),
			errors: []string{"must not equal the pool's own chain selector"},
		},
		{
			name: "rejects_duplicate_remote_selectors",
			input: singlePoolInput(tokensapi.PoolConfigUpdate{
				TokenPoolRef: poolRef,
				Remotes: []tokensapi.RemoteConfigUpdate{
					{RemoteChainSelector: dst, TokenTransferFeeConfig: &tokensapi.PartialTokenTransferFeeConfig{IsEnabled: cciputils.NewOptional(true), DestBytesOverhead: cciputils.NewOptional(uint32(320))}},
					{RemoteChainSelector: dst, TokenTransferFeeConfig: &tokensapi.PartialTokenTransferFeeConfig{IsEnabled: cciputils.NewOptional(true), DestBytesOverhead: cciputils.NewOptional(uint32(320))}},
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
						{TokenPoolRef: poolRef, FeeAdmin: new("0x2222222222222222222222222222222222222222")},
						{TokenPoolRef: poolRef, FinalityConfig: &finality.Config{WaitForFinality: true}},
					},
				}},
			},
			errors: []string{"duplicate pool entry"},
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
			name: "rejects_missing_is_enabled",
			input: singlePoolInput(tokensapi.PoolConfigUpdate{
				TokenPoolRef: poolRef,
				Remotes: []tokensapi.RemoteConfigUpdate{{
					RemoteChainSelector:    dst,
					TokenTransferFeeConfig: &tokensapi.PartialTokenTransferFeeConfig{DestBytesOverhead: cciputils.NewOptional(uint32(320))},
				}},
			}),
			errors: []string{"must specify isEnabled"},
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
	tc := setupV2PoolsForConfigureImpl(t, "CTP_FIN", false)

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
	tc.env.OperationsBundle = evm_testsetup.BundleWithFreshReporter(tc.env.OperationsBundle)
	before := CurrentBlockEVM(t, tc.env, tc.selA)
	_, err = tokensapi.ConfigureTokenPool().Apply(*tc.env, input)
	require.NoError(t, err)
	after := CurrentBlockEVM(t, tc.env, tc.selA)
	require.Equal(t, before, after, "second identical apply must not send transactions")
}

func TestConfigureTokenPool_Admins(t *testing.T) {
	tc := setupV2PoolsForConfigureImpl(t, "CTP_ADM", false)

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
				RateLimitAdmin: new(newRateLimitAdmin),
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
		FeeAdmin:     new(newFeeAdmin),
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
		RateLimitAdmin: new(newRateLimitAdmin),
		FeeAdmin:       new(newFeeAdmin),
	}
	tc.env.OperationsBundle = evm_testsetup.BundleWithFreshReporter(tc.env.OperationsBundle)
	before := CurrentBlockEVM(t, tc.env, tc.selA)
	_, err = tokensapi.ConfigureTokenPool().Apply(*tc.env, input)
	require.NoError(t, err)
	after := CurrentBlockEVM(t, tc.env, tc.selA)
	require.Equal(t, before, after, "no-op admin update must not send transactions")

	// Invalid admin address formats are validated by the EVM SetTokenPoolAdmins sequence at
	// apply time; the top-level changeset stays chain-agnostic and no longer checks formats.
	input.Chains[0].Pools[0] = tokensapi.PoolConfigUpdate{
		TokenPoolRef:   datastore.AddressRef{Address: tc.poolA.Hex()},
		RateLimitAdmin: new("not-an-address"),
	}
	require.NoError(t, tokensapi.ConfigureTokenPool().VerifyPreconditions(*tc.env, input))
	tc.env.OperationsBundle = evm_testsetup.BundleWithFreshReporter(tc.env.OperationsBundle)
	_, err = tokensapi.ConfigureTokenPool().Apply(*tc.env, input)
	require.ErrorContains(t, err, "invalid rate limit admin address")

	input.Chains[0].Pools[0] = tokensapi.PoolConfigUpdate{
		TokenPoolRef: datastore.AddressRef{Address: tc.poolA.Hex()},
		FeeAdmin:     new("not-an-address"),
	}
	require.NoError(t, tokensapi.ConfigureTokenPool().VerifyPreconditions(*tc.env, input))
	tc.env.OperationsBundle = evm_testsetup.BundleWithFreshReporter(tc.env.OperationsBundle)
	_, err = tokensapi.ConfigureTokenPool().Apply(*tc.env, input)
	require.ErrorContains(t, err, "invalid fee admin address")
}

func TestConfigureTokenPool_Admins_PreV2(t *testing.T) {
	t.Run("v1_5_1", func(t *testing.T) { testConfigureTokenPoolAdminsPreV2(t, cciputils.Version_1_5_1) })
	t.Run("v1_6_1", func(t *testing.T) { testConfigureTokenPoolAdminsPreV2(t, cciputils.Version_1_6_1) })
}

func testConfigureTokenPoolAdminsPreV2(t *testing.T, version *semver.Version) {
	pair := setupLegacyConnectedBnMPair(t, version)

	chainA := pair.env.BlockChains.EVMChains()[pair.selA]

	newRateLimitAdmin := "0x1111111111111111111111111111111111111111"
	input := tokensapi.ConfigureTokenPoolInput{
		MCMS: mcms.Input{},
		Chains: []tokensapi.ConfigureTokenPoolPerChain{{
			ChainSelector: pair.selA,
			Pools: []tokensapi.PoolConfigUpdate{{
				TokenPoolRef:   datastore.AddressRef{Address: pair.oldPoolAddrA.Hex()},
				RateLimitAdmin: new(newRateLimitAdmin),
			}},
		}},
	}
	require.NoError(t, tokensapi.ConfigureTokenPool().VerifyPreconditions(*pair.env, input))
	_, err := tokensapi.ConfigureTokenPool().Apply(*pair.env, input)
	require.NoError(t, err)

	rlAdmin := getRateLimitAdminPreV2(t, version, pair.oldPoolAddrA, chainA.Client)
	require.Equal(t, common.HexToAddress(newRateLimitAdmin), rlAdmin)

	// Idempotency: re-applying the same value sends no transaction.
	pair.env.OperationsBundle = evm_testsetup.BundleWithFreshReporter(pair.env.OperationsBundle)
	before := CurrentBlockEVM(t, pair.env, pair.selA)
	_, err = tokensapi.ConfigureTokenPool().Apply(*pair.env, input)
	require.NoError(t, err)
	after := CurrentBlockEVM(t, pair.env, pair.selA)
	require.Equal(t, before, after, "no-op rate limit admin update must not send a transaction")

	// FeeAdmin is not supported on pre-2.0 pools: the whole update must fail cleanly, and must
	// not silently change the rate limit admin it was combined with in the same call.
	otherRateLimitAdmin := "0x2222222222222222222222222222222222222222"
	input.Chains[0].Pools[0] = tokensapi.PoolConfigUpdate{
		TokenPoolRef:   datastore.AddressRef{Address: pair.oldPoolAddrA.Hex()},
		RateLimitAdmin: new(otherRateLimitAdmin),
		FeeAdmin:       new("0x3333333333333333333333333333333333333333"),
	}
	require.NoError(t, tokensapi.ConfigureTokenPool().VerifyPreconditions(*pair.env, input))
	pair.env.OperationsBundle = evm_testsetup.BundleWithFreshReporter(pair.env.OperationsBundle)
	_, err = tokensapi.ConfigureTokenPool().Apply(*pair.env, input)
	require.ErrorContains(t, err, "fee admin is not supported")

	rlAdmin = getRateLimitAdminPreV2(t, version, pair.oldPoolAddrA, chainA.Client)
	require.Equal(t, common.HexToAddress(newRateLimitAdmin), rlAdmin, "rate limit admin must be unchanged after a rejected combined update")
}

// getRateLimitAdminPreV2 reads the on-chain rate limit admin for a pre-2.0 EVM token pool.
// getRateLimitAdmin is ABI-identical across 1.5.1 and 1.6.1, but the call was removed from the
// pool ABI in 2.0 (folded into DynamicConfig), so the version-specific binding must be selected
// explicitly here rather than reusing the v2.0 binding the rest of this file relies on.
func getRateLimitAdminPreV2(t *testing.T, version *semver.Version, address common.Address, backend bind.ContractBackend) common.Address {
	t.Helper()
	opts := &bind.CallOpts{Context: t.Context()}
	switch {
	case cciputils.Version_1_5_1.Equal(version):
		tp, err := tokenpoolV1_5_1.NewTokenPool(address, backend)
		require.NoError(t, err)
		rlAdmin, err := tp.GetRateLimitAdmin(opts)
		require.NoError(t, err)
		return rlAdmin
	case cciputils.Version_1_6_1.Equal(version):
		tp, err := tokenpoolV1_6_1.NewTokenPool(address, backend)
		require.NoError(t, err)
		rlAdmin, err := tp.GetRateLimitAdmin(opts)
		require.NoError(t, err)
		return rlAdmin
	default:
		t.Fatalf("unsupported pre-2.0 pool version for fetching rate limit admin: %s", version.String())
		return common.Address{}
	}
}

func TestConfigureTokenPool_FeeConfig_PreV2(t *testing.T) {
	t.Run("v1_5_1", func(t *testing.T) { testConfigureTokenPoolFeeConfigPreV2(t, cciputils.Version_1_5_1) })
	t.Run("v1_6_1", func(t *testing.T) { testConfigureTokenPoolFeeConfigPreV2(t, cciputils.Version_1_6_1) })
}

func testConfigureTokenPoolFeeConfigPreV2(t *testing.T, version *semver.Version) {
	pair := setupLegacyConnectedBnMPair(t, version)

	// ConnectChains + token expansion may have cached pre-lane router reads; refresh before fee I/O
	// (same pattern as runAutoMigrateUpgrade in token_expansion_migration_test.go).
	pair.env.OperationsBundle = evm_testsetup.BundleWithFreshReporter(pair.env.OperationsBundle)

	input := tokensapi.ConfigureTokenPoolInput{
		MCMS: mcms.Input{},
		Chains: []tokensapi.ConfigureTokenPoolPerChain{{
			ChainSelector: pair.selA,
			Pools: []tokensapi.PoolConfigUpdate{{
				TokenPoolRef: datastore.AddressRef{Address: pair.oldPoolAddrA.Hex()},
				Remotes: []tokensapi.RemoteConfigUpdate{{
					RemoteChainSelector: pair.selB,
					TokenTransferFeeConfig: &tokensapi.PartialTokenTransferFeeConfig{
						IsEnabled:         cciputils.NewOptional(true),
						DestBytesOverhead: cciputils.NewOptional(uint32(320)),
						DestGasOverhead:   cciputils.NewOptional(uint32(21_000)),
					},
				}},
			}},
		}},
	}
	require.NoError(t, tokensapi.ConfigureTokenPool().VerifyPreconditions(*pair.env, input))
	_, err := tokensapi.ConfigureTokenPool().Apply(*pair.env, input)
	require.NoError(t, err)

	feeAdapter, fqRef, err := fees.ResolveFeeAdapter(pair.env.OperationsBundle, pair.env.BlockChains, pair.env.DataStore, pair.selA, pair.selB)
	require.NoError(t, err)
	onChainFee, err := feeAdapter.GetOnchainTokenTransferFeeConfig(pair.env.OperationsBundle, pair.env.BlockChains, fqRef, pair.selA, pair.selB, pair.tokAddrA.Hex())
	require.NoError(t, err)
	require.True(t, onChainFee.IsEnabled)
	require.Equal(t, uint32(320), onChainFee.DestBytesOverhead)
	require.Equal(t, uint32(21_000), onChainFee.DestGasOverhead)

	// Idempotency: re-applying the same config sends no transaction.
	pair.env.OperationsBundle = evm_testsetup.BundleWithFreshReporter(pair.env.OperationsBundle)
	before := CurrentBlockEVM(t, pair.env, pair.selA)
	_, err = tokensapi.ConfigureTokenPool().Apply(*pair.env, input)
	require.NoError(t, err)
	after := CurrentBlockEVM(t, pair.env, pair.selA)
	require.Equal(t, before, after, "no-op fee config update must not send a transaction")
}

func TestConfigureTokenPool_UnsupportedFields_PreV2(t *testing.T) {
	t.Run("v1_5_1", func(t *testing.T) { testConfigureTokenPoolUnsupportedFieldsPreV2(t, cciputils.Version_1_5_1) })
	t.Run("v1_6_1", func(t *testing.T) { testConfigureTokenPoolUnsupportedFieldsPreV2(t, cciputils.Version_1_6_1) })
}

func testConfigureTokenPoolUnsupportedFieldsPreV2(t *testing.T, version *semver.Version) {
	pair := setupLegacyConnectedBnMPair(t, version)

	t.Run("finalityConfig", func(t *testing.T) {
		finalityInput := tokensapi.ConfigureTokenPoolInput{
			MCMS: mcms.Input{},
			Chains: []tokensapi.ConfigureTokenPoolPerChain{{
				ChainSelector: pair.selA,
				Pools: []tokensapi.PoolConfigUpdate{{
					TokenPoolRef:   datastore.AddressRef{Address: pair.oldPoolAddrA.Hex()},
					FinalityConfig: &finality.Config{WaitForFinality: true},
				}},
			}},
		}
		require.NoError(t, tokensapi.ConfigureTokenPool().VerifyPreconditions(*pair.env, finalityInput))
		_, err := tokensapi.ConfigureTokenPool().Apply(*pair.env, finalityInput)
		require.ErrorContains(t, err, "does not support finality config updates")
	})

	t.Run("feeAdmin", func(t *testing.T) {
		feeAdminInput := tokensapi.ConfigureTokenPoolInput{
			MCMS: mcms.Input{},
			Chains: []tokensapi.ConfigureTokenPoolPerChain{{
				ChainSelector: pair.selA,
				Pools: []tokensapi.PoolConfigUpdate{{
					TokenPoolRef: datastore.AddressRef{Address: pair.oldPoolAddrA.Hex()},
					FeeAdmin:     new("0x4444444444444444444444444444444444444444"),
				}},
			}},
		}
		require.NoError(t, tokensapi.ConfigureTokenPool().VerifyPreconditions(*pair.env, feeAdminInput))
		_, err := tokensapi.ConfigureTokenPool().Apply(*pair.env, feeAdminInput)
		require.ErrorContains(t, err, "fee admin is not supported")
	})
}

func TestConfigureTokenPool_FeeConfig(t *testing.T) {
	tc := setupV2PoolsForConfigureImpl(t, "CTP_FEE", false)

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
		IsEnabled:       cciputils.NewOptional(true),
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
	tc.env.OperationsBundle = evm_testsetup.BundleWithFreshReporter(tc.env.OperationsBundle)
	before := CurrentBlockEVM(t, tc.env, tc.selA)
	_, err = tokensapi.ConfigureTokenPool().Apply(*tc.env, input)
	require.NoError(t, err)
	after := CurrentBlockEVM(t, tc.env, tc.selA)
	require.Equal(t, before, after, "no-op fee update must not send transactions")
}

func TestConfigureTokenPool_CombinedUpdate(t *testing.T) {
	tc := setupV2PoolsForConfigureImpl(t, "CTP_ALL", false)

	admin := "0x3333333333333333333333333333333333333333"
	newFinality := finality.Config{BlockDepth: 7}
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
					FeeAdmin:       new(admin),
					Remotes: []tokensapi.RemoteConfigUpdate{{
						RemoteChainSelector:    tc.selB,
						TokenTransferFeeConfig: fee,
					}},
				}},
			},
			{
				ChainSelector: tc.selB,
				Pools: []tokensapi.PoolConfigUpdate{{
					TokenPoolRef: datastore.AddressRef{Address: tc.poolB.Hex()},
					Remotes: []tokensapi.RemoteConfigUpdate{{
						RemoteChainSelector:    tc.selA,
						TokenTransferFeeConfig: fee,
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
	require.Equal(t, common.HexToAddress(admin), dynCfg.FeeAdmin)
	feeCfgA, err := pool.GetTokenTransferFeeConfig(&bind.CallOpts{Context: t.Context()}, common.Address{}, tc.selB, [4]byte{}, nil)
	require.NoError(t, err)
	require.True(t, feeCfgA.IsEnabled, "combined fee config poolA A→B must be enabled")
	poolBContract, err := tokenpoolV2_0_0.NewTokenPool(tc.poolB, tc.clientB)
	require.NoError(t, err)
	feeCfgB, err := poolBContract.GetTokenTransferFeeConfig(&bind.CallOpts{Context: t.Context()}, common.Address{}, tc.selA, [4]byte{}, nil)
	require.NoError(t, err)
	require.True(t, feeCfgB.IsEnabled, "combined fee config poolB B→A must be enabled")

	// Full idempotency across every feature at once.
	tc.env.OperationsBundle = evm_testsetup.BundleWithFreshReporter(tc.env.OperationsBundle)
	beforeA := CurrentBlockEVM(t, tc.env, tc.selA)
	beforeB := CurrentBlockEVM(t, tc.env, tc.selB)
	out, err = tokensapi.ConfigureTokenPool().Apply(*tc.env, input)
	require.NoError(t, err)
	require.Empty(t, out.MCMSTimelockProposals, "combined no-op must not emit proposals")
	afterA := CurrentBlockEVM(t, tc.env, tc.selA)
	afterB := CurrentBlockEVM(t, tc.env, tc.selB)
	require.Equal(t, beforeA, afterA, "combined no-op must not send transactions on chain A")
	require.Equal(t, beforeB, afterB, "combined no-op must not send transactions on chain B")
}

func TestConfigureTokenPool_MCMSOwnedPool(t *testing.T) {
	tc := setupV2PoolsForConfigureImpl(t, "CTP_MCMS", true)

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
	tc.env.OperationsBundle = evm_testsetup.BundleWithFreshReporter(tc.env.OperationsBundle)
	out2, err := tokensapi.ConfigureTokenPool().Apply(*tc.env, input)
	require.NoError(t, err)
	require.Empty(t, out2.MCMSTimelockProposals, "no-op re-apply must not emit an MCMS proposal")
}

// solanaPoolFixture identifies a deployed Solana token pool for use by ConfigureTokenPool tests.
// Ref is the fully resolved pool reference (Address is the pool program ID, shared across every
// mint that pool type serves); Mint is the specific token mint the pool was initialized for.
type solanaPoolFixture struct {
	Ref  datastore.AddressRef
	Mint solana.PublicKey
}

// setupSolanaPoolsForConfigure stands up a Solana chain (plus a throwaway EVM chain, matching
// the environment.New pattern used elsewhere in this package) with an initialized BurnMint pool
// and an initialized LockRelease pool, ready for ConfigureTokenPool tests.
func setupSolanaPoolsForConfigure(t *testing.T) (env *cldf_deployment.Environment, bnm solanaPoolFixture, lnr solanaPoolFixture) {
	t.Helper()

	src := chainsel.SOLANA_DEVNET.Selector
	dst := chainsel.TEST_90000002.Selector

	programsPath, ds, err := PreloadSolanaEnvironment(t, src)
	require.NoError(t, err)

	env, err = environment.New(
		t.Context(),
		environment.WithSolanaContainer(t, []uint64{src}, programsPath, solanaProgramIDs),
		environment.WithEVMSimulated(t, []uint64{dst}),
	)
	require.NoError(t, err)
	env.DataStore = ds.Seal()

	solChain, ok := env.BlockChains.SolanaChains()[src]
	require.True(t, ok, "Solana chain not found in environment")

	deployRegistry := deployapi.GetRegistry()
	SeedUltraFastCurseMCMS(t, env)
	deployOut, err := deployapi.DeployContracts(deployRegistry).Apply(*env, deployapi.ContractDeploymentConfig{
		MCMS: mcms.Input{},
		Chains: map[uint64]deployapi.ContractDeploymentConfigPerChain{
			src: NewDefaultDeploymentConfigForSolana(cciputils.Version_1_6_0),
		},
	})
	require.NoError(t, err)
	MergeAddresses(t, env, deployOut.DataStore)

	deployPool := func(symbol, poolType string) solanaPoolFixture {
		expansionOut, err := tokensapi.TokenExpansion().Apply(*env, tokensapi.TokenExpansionInput{
			ChainAdapterVersion: cciputils.Version_1_6_0,
			MCMS:                NewDefaultInputForMCMS("Configure Token Pool test setup"),
			TokenExpansionInputPerChain: map[uint64]tokensapi.TokenExpansionInputPerChain{
				src: {
					TokenPoolVersion: cciputils.Version_1_6_0,
					DeployTokenPoolInput: &tokensapi.DeployTokenPoolInput{
						TokenPoolQualifier: "",
						PoolType:           poolType,
						RateLimitAdmin:     solana.NewWallet().PublicKey().String(),
					},
					DeployTokenInput: &tokensapi.DeployTokenInput{
						Decimals:               9,
						Symbol:                 symbol,
						Name:                   symbol,
						Type:                   solanautils.SPLTokens,
						ExternalAdmin:          solana.NewWallet().PublicKey().String(),
						DisableFreezeAuthority: true,
						Senders:                []string{solChain.DeployerKey.PublicKey().String()},
					},
				},
			},
		})
		require.NoError(t, err)
		MergeAddresses(t, env, expansionOut.DataStore)
		testhelpers.ProcessTimelockProposals(t, *env, expansionOut.MCMSTimelockProposals, false)

		tokenRef, err := datastore_utils.FindAndFormatRef(
			env.DataStore,
			datastore.AddressRef{Qualifier: symbol},
			src,
			datastore_utils.FullRef,
		)
		require.NoError(t, err)

		poolRef, err := datastore_utils.FindAndFormatRef(
			env.DataStore,
			datastore.AddressRef{
				ChainSelector: src,
				Qualifier:     "",
				Type:          datastore.ContractType(poolType),
				Version:       cciputils.Version_1_6_0,
			},
			src,
			datastore_utils.FullRef,
		)
		require.NoError(t, err)

		return solanaPoolFixture{Ref: poolRef, Mint: solana.MustPublicKeyFromBase58(tokenRef.Address)}
	}

	bnm = deployPool("CTP_SOL_BNM", cciputils.BurnMintTokenPool.String())
	lnr = deployPool("CTP_SOL_LNR", cciputils.LockReleaseTokenPool.String())

	return env, bnm, lnr
}

// solanaPoolRateLimitAdmin reads the on-chain rate limit admin for a Solana token pool.
// BnM and LnR pools share the same state layout, so either binding decodes both.
func solanaPoolRateLimitAdmin(t *testing.T, chain solchain.Chain, poolProgramID solana.PublicKey, tokenMint solana.PublicKey) solana.PublicKey {
	t.Helper()
	poolStatePDA, err := tokens.TokenPoolConfigAddress(tokenMint, poolProgramID)
	require.NoError(t, err)
	var poolState burnmint_token_pool.State
	require.NoError(t, chain.GetAccountDataBorshInto(t.Context(), poolStatePDA, &poolState))
	return poolState.Config.RateLimitAdmin
}

func TestConfigureTokenPool_Admins_Solana(t *testing.T) {
	env, bnm, lnr := setupSolanaPoolsForConfigure(t) // helper from Step 6

	for _, tc := range []struct {
		name string
		pool solanaPoolFixture // {Ref datastore.AddressRef, Mint solana.PublicKey}
	}{
		{"burnmint", bnm},
		{"lockrelease", lnr},
	} {
		t.Run(tc.name, func(t *testing.T) {
			newAdmin := solana.NewWallet().PublicKey()
			newAdminStr := newAdmin.String()

			// A Solana pool program ID is shared across every mint that pool type serves, so it
			// cannot alone identify "the pool for this mint" the way an EVM pool address can.
			// The chain-agnostic ConfigureTokenPoolInput has no separate token field, so the
			// mint-specific pool config PDA is what's passed as TokenPoolRef.Address here: the
			// Solana adapter's ResolveTokenPoolRef (see chains/solana/deployment/v1_6_0/sequences/tokens.go)
			// recognizes a PDA, reads its owner to recover the program ID, and tags the resolved
			// ref with the PDA so DeriveTokenAddress can still identify the mint downstream.
			solChain := env.BlockChains.SolanaChains()[tc.pool.Ref.ChainSelector]
			poolProgramID := solana.MustPublicKeyFromBase58(tc.pool.Ref.Address)
			apiPoolRef := solanaApiPoolRef(t, tc.pool)

			out, err := tokensapi.ConfigureTokenPool().Apply(*env, tokensapi.ConfigureTokenPoolInput{
				Chains: []tokensapi.ConfigureTokenPoolPerChain{{
					ChainSelector: tc.pool.Ref.ChainSelector,
					Pools: []tokensapi.PoolConfigUpdate{{
						TokenPoolRef:   apiPoolRef,
						RateLimitAdmin: &newAdminStr,
					}},
				}},
				MCMS: NewDefaultInputForMCMS("Configure Token Pool"),
			})
			require.NoError(t, err)
			testhelpers.ProcessTimelockProposals(t, *env, out.MCMSTimelockProposals, false)

			require.Equal(t, newAdmin, solanaPoolRateLimitAdmin(t, solChain, poolProgramID, tc.pool.Mint),
				"rate limit admin should match the requested value")

			// Re-applying the same value must produce no transactions. Refresh the bundle's
			// reporter first: ExecuteSequence memoizes on (sequence def, input) and would
			// otherwise return the first apply's cached report without re-running the op,
			// making this assertion vacuous. See TestConfigureTokenPool_MCMSOwnedPool above.
			env.OperationsBundle = evm_testsetup.BundleWithFreshReporter(env.OperationsBundle)
			out2, err := tokensapi.ConfigureTokenPool().Apply(*env, tokensapi.ConfigureTokenPoolInput{
				Chains: []tokensapi.ConfigureTokenPoolPerChain{{
					ChainSelector: tc.pool.Ref.ChainSelector,
					Pools: []tokensapi.PoolConfigUpdate{{
						TokenPoolRef:   apiPoolRef,
						RateLimitAdmin: &newAdminStr,
					}},
				}},
				MCMS: NewDefaultInputForMCMS("Configure Token Pool"),
			})
			require.NoError(t, err)
			require.Empty(t, out2.MCMSTimelockProposals, "re-applying an unchanged admin must emit no proposals")
			require.Equal(t, newAdmin, solanaPoolRateLimitAdmin(t, solChain, poolProgramID, tc.pool.Mint),
				"rate limit admin must be unchanged after the no-op apply")
		})
	}
}

// solanaApiPoolRef builds the API-facing TokenPoolRef for a Solana pool fixture: the pool
// config PDA, not the bare pool program ID. Passing the program ID directly into
// ConfigureTokenPool().Apply fails with "token derivation is only possible if a pool PDA is
// provided" — the generic ResolveTokenPoolRef exact-matches the datastore by program ID and
// short-circuits before the Solana PDA-normalizing resolver runs. See
// TestConfigureTokenPool_Admins_Solana for the same pattern.
func solanaApiPoolRef(t *testing.T, pool solanaPoolFixture) datastore.AddressRef {
	t.Helper()
	poolProgramID := solana.MustPublicKeyFromBase58(pool.Ref.Address)
	poolConfigPDA, err := tokens.TokenPoolConfigAddress(pool.Mint, poolProgramID)
	require.NoError(t, err)
	return datastore.AddressRef{ChainSelector: pool.Ref.ChainSelector, Address: poolConfigPDA.String()}
}

func TestConfigureTokenPool_UnsupportedFields_Solana(t *testing.T) {
	env, bnm, _ := setupSolanaPoolsForConfigure(t)
	apiPoolRef := solanaApiPoolRef(t, bnm)
	feeAdmin := solana.NewWallet().PublicKey().String()

	t.Run("feeAdmin", func(t *testing.T) {
		_, err := tokensapi.ConfigureTokenPool().Apply(*env, tokensapi.ConfigureTokenPoolInput{
			Chains: []tokensapi.ConfigureTokenPoolPerChain{{
				ChainSelector: bnm.Ref.ChainSelector,
				Pools: []tokensapi.PoolConfigUpdate{{
					TokenPoolRef: apiPoolRef,
					FeeAdmin:     &feeAdmin,
				}},
			}},
			MCMS: NewDefaultInputForMCMS("Configure Token Pool"),
		})
		require.Error(t, err)
		require.Contains(t, err.Error(), "fee admin is not supported")
	})

	t.Run("finalityConfig", func(t *testing.T) {
		_, err := tokensapi.ConfigureTokenPool().Apply(*env, tokensapi.ConfigureTokenPoolInput{
			Chains: []tokensapi.ConfigureTokenPoolPerChain{{
				ChainSelector: bnm.Ref.ChainSelector,
				Pools: []tokensapi.PoolConfigUpdate{{
					TokenPoolRef:   apiPoolRef,
					FinalityConfig: &finality.Config{},
				}},
			}},
			MCMS: NewDefaultInputForMCMS("Configure Token Pool"),
		})
		require.Error(t, err)
		require.Contains(t, err.Error(), "does not support finality config updates")
	})
}

func TestConfigureTokenPool_FeeConfig_Solana(t *testing.T) {
	env, bnm, _ := setupSolanaPoolsForConfigure(t)
	apiPoolRef := solanaApiPoolRef(t, bnm)
	remote := chainsel.TEST_90000002.Selector // the EVM chain from the helper's env

	out, err := tokensapi.ConfigureTokenPool().Apply(*env, tokensapi.ConfigureTokenPoolInput{
		Chains: []tokensapi.ConfigureTokenPoolPerChain{{
			ChainSelector: bnm.Ref.ChainSelector,
			Pools: []tokensapi.PoolConfigUpdate{{
				TokenPoolRef: apiPoolRef,
				Remotes: []tokensapi.RemoteConfigUpdate{{
					RemoteChainSelector: remote,
					TokenTransferFeeConfig: &tokensapi.PartialTokenTransferFeeConfig{
						IsEnabled:                  cciputils.NewOptional(true),
						DefaultFinalityFeeUSDCents: cciputils.NewOptional(uint32(10)),
						DestBytesOverhead:          cciputils.NewOptional(uint32(200_000)),
					},
				}},
			}},
		}},
		MCMS: NewDefaultInputForMCMS("Configure Token Pool"),
	})
	require.NoError(t, err)
	testhelpers.ProcessTimelockProposals(t, *env, out.MCMSTimelockProposals, false)

	// Solana pools are below v2.0.0, so the fee config lands on the FeeQuoter. Read it back
	// through the same resolution path the changeset used
	// (deployment/tokens/configure_tokens_for_transfers.go:512).
	feeAdapter, fqRef, err := fees.ResolveFeeAdapter(env.OperationsBundle, env.BlockChains, env.DataStore, bnm.Ref.ChainSelector, remote)
	require.NoError(t, err)

	onchain, err := feeAdapter.GetOnchainTokenTransferFeeConfig(
		env.OperationsBundle,
		env.BlockChains,
		fqRef,
		bnm.Ref.ChainSelector,
		remote,
		bnm.Mint.String(),
	)
	require.NoError(t, err)
	require.True(t, onchain.IsEnabled, "fee config should be enabled on-chain")
	require.Equal(t, uint32(200_000), onchain.DestBytesOverhead, "destBytesOverhead should match the requested value")
	require.Equal(t, uint32(10), onchain.MinFeeUSDCents, "min fee USD cents should match the requested default finality fee")
}
