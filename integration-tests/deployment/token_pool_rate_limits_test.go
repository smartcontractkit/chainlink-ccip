package deployment

import (
	"fmt"
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/stretchr/testify/require"

	evm_datastore_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/datastore"
	bnmERC20ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/burn_mint_erc20"
	bnmOpsV1_6_1 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_1/operations/burn_mint_token_pool"
	bnmOpsV2_0_0 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/burn_mint_token_pool"
	tokenpoolV1_6_1 "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_1/token_pool"
	tokenpoolV2_0_0 "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v2_0_0/token_pool"
	tokensapi "github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	cciputils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"

	_ "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_1/adapters"
	_ "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/adapters"
)

func TestTPRL_VerifyPreconditions_RemoteOutbounds(t *testing.T) {
	sel := chainsel.TEST_90000001.Selector
	dst := chainsel.TEST_90000002.Selector
	env, err := environment.New(t.Context(), environment.WithEVMSimulated(t, []uint64{sel}))
	require.NoError(t, err)

	cs := tokensapi.SetTokenPoolRateLimits()
	baseInput := func(ro tokensapi.RemoteOutbounds) tokensapi.TPRLInput {
		return tokensapi.TPRLInput{
			MCMS: mcms.Input{},
			Configs: map[uint64]tokensapi.TPRLConfig{
				sel: {
					RemoteOutbounds: map[uint64]tokensapi.RemoteOutbounds{
						dst: ro,
					},
				},
				dst: {
					RemoteOutbounds: map[uint64]tokensapi.RemoteOutbounds{
						sel: ro,
					},
				},
			},
		}
	}

	cases := []struct {
		name   string
		remote tokensapi.RemoteOutbounds
		errors []string
	}{
		{
			name:   "rejects_duplicate_fast_finality_buckets",
			errors: []string{"multiple rate limit buckets with fastFinality=true"},
			remote: tokensapi.RemoteOutbounds{
				Outbounds: []tokensapi.RateLimitConfig{
					{RateLimit: tokensapi.RateLimiterConfigFloatInput{IsEnabled: true, Capacity: 100, Rate: 10}, FastFinality: true},
					{RateLimit: tokensapi.RateLimiterConfigFloatInput{IsEnabled: true, Capacity: 200, Rate: 20}, FastFinality: true},
				},
			},
		},
		{
			name:   "rejects_more_than_two_buckets",
			errors: []string{"at most two rate limit buckets allowed"},
			remote: tokensapi.RemoteOutbounds{
				Outbounds: []tokensapi.RateLimitConfig{
					{RateLimit: tokensapi.RateLimiterConfigFloatInput{IsEnabled: true, Capacity: 10, Rate: 1}, FastFinality: false},
					{RateLimit: tokensapi.RateLimiterConfigFloatInput{IsEnabled: true, Capacity: 20, Rate: 2}, FastFinality: true},
					{RateLimit: tokensapi.RateLimiterConfigFloatInput{IsEnabled: true, Capacity: 30, Rate: 3}, FastFinality: false},
				},
			},
		},
		{
			name: "allows_default_plus_fast_finality",
			remote: tokensapi.RemoteOutbounds{
				Outbounds: []tokensapi.RateLimitConfig{
					{RateLimit: tokensapi.RateLimiterConfigFloatInput{IsEnabled: true, Capacity: 100, Rate: 10}, FastFinality: true},
					{RateLimit: tokensapi.RateLimiterConfigFloatInput{IsEnabled: true, Capacity: 50, Rate: 5}, FastFinality: false},
				},
			},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			err := cs.VerifyPreconditions(*env, baseInput(tt.remote))
			if len(tt.errors) > 0 {
				require.Error(t, err)
				for _, substr := range tt.errors {
					require.Contains(t, err.Error(), substr)
				}
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestTPRL_BasicSetTokenPoolRateLimitsV2(t *testing.T) {
	const tokenSymb = "TPRL_V2"
	const decimalsA = 18
	const decimalsB = 6

	cases := []struct {
		name               string
		aTowardB, bTowardA tokensapi.RemoteOutbounds
	}{
		{
			name: "single_default_via_RateLimit_field",
			aTowardB: tokensapi.RemoteOutbounds{
				RateLimit: &tokensapi.RateLimiterConfigFloatInput{IsEnabled: true, Capacity: 1000, Rate: 100},
			},
			bTowardA: tokensapi.RemoteOutbounds{
				RateLimit: &tokensapi.RateLimiterConfigFloatInput{IsEnabled: true, Capacity: 2000, Rate: 200},
			},
		},
		{
			name: "single_default_only_buckets",
			aTowardB: tokensapi.RemoteOutbounds{
				Outbounds: []tokensapi.RateLimitConfig{
					{RateLimit: tokensapi.RateLimiterConfigFloatInput{IsEnabled: true, Capacity: 3030, Rate: 303}, FastFinality: false},
				},
			},
			bTowardA: tokensapi.RemoteOutbounds{
				Outbounds: []tokensapi.RateLimitConfig{
					{RateLimit: tokensapi.RateLimiterConfigFloatInput{IsEnabled: true, Capacity: 4040, Rate: 404}, FastFinality: false},
				},
			},
		},
		{
			name: "single_custom_only_buckets",
			aTowardB: tokensapi.RemoteOutbounds{
				Outbounds: []tokensapi.RateLimitConfig{
					{RateLimit: tokensapi.RateLimiterConfigFloatInput{IsEnabled: true, Capacity: 3030, Rate: 303}, FastFinality: true},
				},
			},
			bTowardA: tokensapi.RemoteOutbounds{
				Outbounds: []tokensapi.RateLimitConfig{
					{RateLimit: tokensapi.RateLimiterConfigFloatInput{IsEnabled: true, Capacity: 4040, Rate: 404}, FastFinality: true},
				},
			},
		},
		{
			name: "default_and_custom_buckets_using_rate_limit_alias",
			aTowardB: tokensapi.RemoteOutbounds{
				RateLimit: &tokensapi.RateLimiterConfigFloatInput{IsEnabled: true, Capacity: 7010, Rate: 701},
				Outbounds: []tokensapi.RateLimitConfig{
					{RateLimit: tokensapi.RateLimiterConfigFloatInput{IsEnabled: true, Capacity: 5050, Rate: 505}, FastFinality: true},
				},
			},
			bTowardA: tokensapi.RemoteOutbounds{
				RateLimit: &tokensapi.RateLimiterConfigFloatInput{IsEnabled: true, Capacity: 7020, Rate: 702},
				Outbounds: []tokensapi.RateLimitConfig{
					{RateLimit: tokensapi.RateLimiterConfigFloatInput{IsEnabled: true, Capacity: 6060, Rate: 606}, FastFinality: true},
				},
			},
		},
		{
			name: "default_and_custom_buckets",
			aTowardB: tokensapi.RemoteOutbounds{
				Outbounds: []tokensapi.RateLimitConfig{
					{RateLimit: tokensapi.RateLimiterConfigFloatInput{IsEnabled: true, Capacity: 1000, Rate: 100}, FastFinality: false},
					{RateLimit: tokensapi.RateLimiterConfigFloatInput{IsEnabled: true, Capacity: 3030, Rate: 303}, FastFinality: true},
				},
			},
			bTowardA: tokensapi.RemoteOutbounds{
				Outbounds: []tokensapi.RateLimitConfig{
					{RateLimit: tokensapi.RateLimiterConfigFloatInput{IsEnabled: true, Capacity: 2000, Rate: 200}, FastFinality: false},
					{RateLimit: tokensapi.RateLimiterConfigFloatInput{IsEnabled: true, Capacity: 4040, Rate: 404}, FastFinality: true},
				},
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			// Setup test env
			selA := chainsel.TEST_90000001.Selector
			selB := chainsel.TEST_90000002.Selector
			e, err := environment.New(t.Context(), environment.WithEVMSimulated(t, []uint64{selA, selB}))
			require.NoError(t, err)

			// Deploy v2 contracts
			cumulative := datastore.NewMemoryDataStore()
			DeployChainContractsV2_0_0(t, e, cumulative, selA)
			DeployChainContractsV2_0_0(t, e, cumulative, selB)
			e.DataStore = cumulative.Seal()

			// We're skipping MCMS setup so we need to use the deployer key
			disabledOutbound := tokensapi.RateLimiterConfigFloatInput{IsEnabled: false}
			deployerA := e.BlockChains.EVMChains()[selA].DeployerKey.From
			deployerB := e.BlockChains.EVMChains()[selB].DeployerKey.From
			clientA := e.BlockChains.EVMChains()[selA].Client
			clientB := e.BlockChains.EVMChains()[selB].Client

			// Deploy and connect v2 token pools
			expansionOut, err := tokensapi.TokenExpansion().Apply(*e, tokensapi.TokenExpansionInput{
				ChainAdapterVersion: cciputils.Version_2_0_0,
				MCMS:                mcms.Input{},
				TokenExpansionInputPerChain: map[uint64]tokensapi.TokenExpansionInputPerChain{
					selA: {
						SkipOwnershipTransfer: true,
						TokenPoolVersion:      bnmOpsV2_0_0.Version,
						DeployTokenInput: &tokensapi.DeployTokenInput{
							Name:          tokenSymb,
							Symbol:        tokenSymb,
							Decimals:      decimalsA,
							ExternalAdmin: deployerA.Hex(),
							CCIPAdmin:     deployerA.Hex(),
							Type:          bnmERC20ops.ContractType,
						},
						DeployTokenPoolInput: &tokensapi.DeployTokenPoolInput{
							PoolType:           string(bnmOpsV2_0_0.ContractType),
							TokenPoolQualifier: "",
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
							Name:          tokenSymb,
							Symbol:        tokenSymb,
							Decimals:      decimalsB,
							ExternalAdmin: deployerB.Hex(),
							CCIPAdmin:     deployerB.Hex(),
							Type:          bnmERC20ops.ContractType,
						},
						DeployTokenPoolInput: &tokensapi.DeployTokenPoolInput{
							PoolType:           string(bnmOpsV2_0_0.ContractType),
							TokenPoolQualifier: "",
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

			// Create token pool instances
			fltrA := datastore.AddressRef{ChainSelector: selA, Type: datastore.ContractType(bnmOpsV2_0_0.ContractType), Version: bnmOpsV2_0_0.Version, Qualifier: ""}
			poolA, err := datastore_utils.FindAndFormatRef(e.DataStore, fltrA, selA, evm_datastore_utils.ToEVMAddress)
			require.NoError(t, err)
			fltrB := datastore.AddressRef{ChainSelector: selB, Type: datastore.ContractType(bnmOpsV2_0_0.ContractType), Version: bnmOpsV2_0_0.Version, Qualifier: ""}
			poolB, err := datastore_utils.FindAndFormatRef(e.DataStore, fltrB, selB, evm_datastore_utils.ToEVMAddress)
			require.NoError(t, err)
			require.NotEqual(t, poolA, poolB)

			// Apply the rate limits
			_, err = tokensapi.SetTokenPoolRateLimits().Apply(*e, tokensapi.TPRLInput{
				MCMS: mcms.Input{},
				Configs: map[uint64]tokensapi.TPRLConfig{
					selA: {
						ChainAdapterVersion: cciputils.Version_2_0_0,
						TokenRef:            datastore.AddressRef{Qualifier: tokenSymb},
						TokenPoolRef:        datastore.AddressRef{Address: poolA.Hex()},
						RemoteOutbounds:     map[uint64]tokensapi.RemoteOutbounds{selB: tc.aTowardB},
					},
					selB: {
						ChainAdapterVersion: cciputils.Version_2_0_0,
						TokenRef:            datastore.AddressRef{Qualifier: tokenSymb},
						TokenPoolRef:        datastore.AddressRef{Address: poolB.Hex()},
						RemoteOutbounds:     map[uint64]tokensapi.RemoteOutbounds{selA: tc.bTowardA},
					},
				},
			})
			require.NoError(t, err)

			// Fetch the latest rate limits
			defaultRateLimitA, ffRateLimitA := getRateLimits(t, cciputils.Version_2_0_0, poolA, clientA, selB)
			defaultRateLimitB, ffRateLimitB := getRateLimits(t, cciputils.Version_2_0_0, poolB, clientB, selA)

			// Validate rate limits were set correctly by reading directly from the contract. We check
			// both pools: each chain's outbound must match its TPRL RemoteOutbounds entry and inbound
			// must match the counterpart's outbound.
			fastFinalityAB, okAB := tc.aTowardB.FastFinalityBucket()
			fastFinalityBA, okBA := tc.bTowardA.FastFinalityBucket()
			require.Equal(t, okAB, okBA, "fast finality bucket presence must match A→B and B→A")
			if okAB {
				validateScaledTPRLBucket(t, "fast_finality selA_pool", decimalsA, ffRateLimitA, fastFinalityAB.RateLimit, fastFinalityBA.RateLimit)
				validateScaledTPRLBucket(t, "fast_finality selB_pool", decimalsB, ffRateLimitB, fastFinalityBA.RateLimit, fastFinalityAB.RateLimit)
			}
			defaultAB, okAB := tc.aTowardB.DefaultBucket()
			defaultBA, okBA := tc.bTowardA.DefaultBucket()
			require.Equal(t, okAB, okBA, "default bucket presence must match A→B and B→A")
			if okAB {
				validateScaledTPRLBucket(t, "default selA_pool", decimalsA, defaultRateLimitA, defaultAB.RateLimit, defaultBA.RateLimit)
				validateScaledTPRLBucket(t, "default selB_pool", decimalsB, defaultRateLimitB, defaultBA.RateLimit, defaultAB.RateLimit)
			}
		})
	}
}

func TestTPRL_NoAccidentalOverwritesV2(t *testing.T) {
	const tokenSymb = "TPRL_V2_FF_NC"
	const decimalsA = 18
	const decimalsB = 6

	// Setup test env
	selA := chainsel.TEST_90000001.Selector
	selB := chainsel.TEST_90000002.Selector
	e, err := environment.New(t.Context(), environment.WithEVMSimulated(t, []uint64{selA, selB}))
	require.NoError(t, err)

	// Deploy v2 contracts
	cumulative := datastore.NewMemoryDataStore()
	DeployChainContractsV2_0_0(t, e, cumulative, selA)
	DeployChainContractsV2_0_0(t, e, cumulative, selB)
	e.DataStore = cumulative.Seal()

	// We're skipping MCMS setup so we need to use the deployer key
	disabledOutbound := tokensapi.RateLimiterConfigFloatInput{IsEnabled: false}
	deployerA := e.BlockChains.EVMChains()[selA].DeployerKey.From
	deployerB := e.BlockChains.EVMChains()[selB].DeployerKey.From
	clientA := e.BlockChains.EVMChains()[selA].Client
	clientB := e.BlockChains.EVMChains()[selB].Client

	// Deploy and connect v2 token pools
	expansionOut, err := tokensapi.TokenExpansion().Apply(*e, tokensapi.TokenExpansionInput{
		ChainAdapterVersion: cciputils.Version_2_0_0,
		MCMS:                mcms.Input{},
		TokenExpansionInputPerChain: map[uint64]tokensapi.TokenExpansionInputPerChain{
			selA: {
				SkipOwnershipTransfer: true,
				TokenPoolVersion:      bnmOpsV2_0_0.Version,
				DeployTokenInput: &tokensapi.DeployTokenInput{
					Name:          tokenSymb,
					Symbol:        tokenSymb,
					Decimals:      decimalsA,
					ExternalAdmin: deployerA.Hex(),
					CCIPAdmin:     deployerA.Hex(),
					Type:          bnmERC20ops.ContractType,
				},
				DeployTokenPoolInput: &tokensapi.DeployTokenPoolInput{
					PoolType:           string(bnmOpsV2_0_0.ContractType),
					TokenPoolQualifier: "",
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
					Name:          tokenSymb,
					Symbol:        tokenSymb,
					Decimals:      decimalsB,
					ExternalAdmin: deployerB.Hex(),
					CCIPAdmin:     deployerB.Hex(),
					Type:          bnmERC20ops.ContractType,
				},
				DeployTokenPoolInput: &tokensapi.DeployTokenPoolInput{
					PoolType:           string(bnmOpsV2_0_0.ContractType),
					TokenPoolQualifier: "",
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

	// Create token pool instances
	fltrA := datastore.AddressRef{ChainSelector: selA, Type: datastore.ContractType(bnmOpsV2_0_0.ContractType), Version: bnmOpsV2_0_0.Version, Qualifier: ""}
	poolA, err := datastore_utils.FindAndFormatRef(e.DataStore, fltrA, selA, evm_datastore_utils.ToEVMAddress)
	require.NoError(t, err)
	fltrB := datastore.AddressRef{ChainSelector: selB, Type: datastore.ContractType(bnmOpsV2_0_0.ContractType), Version: bnmOpsV2_0_0.Version, Qualifier: ""}
	poolB, err := datastore_utils.FindAndFormatRef(e.DataStore, fltrB, selB, evm_datastore_utils.ToEVMAddress)
	require.NoError(t, err)

	// Set initial default rate limits
	initRateLimitAB := tokensapi.RateLimitConfig{RateLimit: tokensapi.RateLimiterConfigFloatInput{IsEnabled: true, Capacity: 1000, Rate: 100}, FastFinality: false}
	initRateLimitBA := tokensapi.RateLimitConfig{RateLimit: tokensapi.RateLimiterConfigFloatInput{IsEnabled: true, Capacity: 2000, Rate: 200}, FastFinality: false}
	initOutboundsAB := tokensapi.RemoteOutbounds{Outbounds: []tokensapi.RateLimitConfig{initRateLimitAB}}
	initOutboundsBA := tokensapi.RemoteOutbounds{Outbounds: []tokensapi.RateLimitConfig{initRateLimitBA}}
	_, err = tokensapi.SetTokenPoolRateLimits().Apply(*e, tokensapi.TPRLInput{
		MCMS: mcms.Input{},
		Configs: map[uint64]tokensapi.TPRLConfig{
			selA: {
				ChainAdapterVersion: cciputils.Version_2_0_0,
				TokenRef:            datastore.AddressRef{Qualifier: tokenSymb},
				TokenPoolRef:        datastore.AddressRef{Address: poolA.Hex()},
				RemoteOutbounds:     map[uint64]tokensapi.RemoteOutbounds{selB: initOutboundsAB},
			},
			selB: {
				ChainAdapterVersion: cciputils.Version_2_0_0,
				TokenRef:            datastore.AddressRef{Qualifier: tokenSymb},
				TokenPoolRef:        datastore.AddressRef{Address: poolB.Hex()},
				RemoteOutbounds:     map[uint64]tokensapi.RemoteOutbounds{selA: initOutboundsBA},
			},
		},
	})
	require.NoError(t, err)

	// Take a snapshot of the initial default rate limits
	preDefaultRateLimitA, _ := getRateLimits(t, cciputils.Version_2_0_0, poolA, clientA, selB)
	preDefaultRateLimitB, _ := getRateLimits(t, cciputils.Version_2_0_0, poolB, clientB, selA)
	require.True(t, preDefaultRateLimitA.OutboundRateLimiterConfig.IsEnabled, "initialize A→B default RL")
	require.True(t, preDefaultRateLimitB.OutboundRateLimiterConfig.IsEnabled, "initialize B→A default RL")

	// Now set the FF rate limits; the initial default rate limits should be preserved
	ffRateLimitAB := tokensapi.RateLimitConfig{RateLimit: tokensapi.RateLimiterConfigFloatInput{IsEnabled: true, Capacity: 3030, Rate: 303}, FastFinality: true}
	ffRateLimitBA := tokensapi.RateLimitConfig{RateLimit: tokensapi.RateLimiterConfigFloatInput{IsEnabled: true, Capacity: 4040, Rate: 404}, FastFinality: true}
	ffOutboundsAB := tokensapi.RemoteOutbounds{Outbounds: []tokensapi.RateLimitConfig{ffRateLimitAB}}
	ffOutboundsBA := tokensapi.RemoteOutbounds{Outbounds: []tokensapi.RateLimitConfig{ffRateLimitBA}}
	_, err = tokensapi.SetTokenPoolRateLimits().Apply(*e, tokensapi.TPRLInput{
		MCMS: mcms.Input{},
		Configs: map[uint64]tokensapi.TPRLConfig{
			selA: {
				ChainAdapterVersion: cciputils.Version_2_0_0,
				TokenRef:            datastore.AddressRef{Qualifier: tokenSymb},
				TokenPoolRef:        datastore.AddressRef{Address: poolA.Hex()},
				RemoteOutbounds:     map[uint64]tokensapi.RemoteOutbounds{selB: ffOutboundsAB},
			},
			selB: {
				ChainAdapterVersion: cciputils.Version_2_0_0,
				TokenRef:            datastore.AddressRef{Qualifier: tokenSymb},
				TokenPoolRef:        datastore.AddressRef{Address: poolB.Hex()},
				RemoteOutbounds:     map[uint64]tokensapi.RemoteOutbounds{selA: ffOutboundsBA},
			},
		},
	})
	require.NoError(t, err)

	// Validate that the default buckets were not modified
	postDefaultRateLimitA, postFastFinalityRateLimitA := getRateLimits(t, cciputils.Version_2_0_0, poolA, clientA, selB)
	postDefaultRateLimitB, postFastFinalityRateLimitB := getRateLimits(t, cciputils.Version_2_0_0, poolB, clientB, selA)
	assertTPRLBucketRateLimiterConfigsUnchanged(t, "check default RL for A→B is unchanged after FF TPRL", preDefaultRateLimitA, postDefaultRateLimitA)
	assertTPRLBucketRateLimiterConfigsUnchanged(t, "check default RL for B→A is unchanged after FF TPRL", preDefaultRateLimitB, postDefaultRateLimitB)

	// Validate that the fast finality buckets were updated
	towardB, towardBOk := ffOutboundsAB.FastFinalityBucket()
	towardA, towardAOk := ffOutboundsBA.FastFinalityBucket()
	require.True(t, towardBOk, "expected to find a fast finality bucket for A→B FF lane")
	require.True(t, towardAOk, "expected to find a fast finality bucket for B→A FF lane")
	validateScaledTPRLBucket(t, "fast_finality selA_pool", decimalsA, postFastFinalityRateLimitA, towardB.RateLimit, towardA.RateLimit)
	validateScaledTPRLBucket(t, "fast_finality selB_pool", decimalsB, postFastFinalityRateLimitB, towardA.RateLimit, towardB.RateLimit)
}

func TestTPRL_AsymmetricPoolVersions(t *testing.T) {
	const tokenSymbl = "TPRL_V161_MIX"
	const decimalsV1 = uint8(18)
	const decimalsV2 = uint8(6)

	selV1 := chainsel.TEST_90000001.Selector
	selV2 := chainsel.TEST_90000002.Selector

	// Expansion: both chains specify symmetric default + fast-finality outbound buckets toward the peer.
	// Inbound slices are synthesized from the counterpart outbound during ConfigureTokensForTransfers.
	outboundTowardV2 := tokensapi.RemoteOutbounds{
		Outbounds: []tokensapi.RateLimitConfig{
			{RateLimit: tokensapi.RateLimiterConfigFloatInput{IsEnabled: true, Capacity: 1001, Rate: 101}, FastFinality: false},
			{RateLimit: tokensapi.RateLimiterConfigFloatInput{IsEnabled: true, Capacity: 3030, Rate: 303}, FastFinality: true},
		},
	}
	outboundTowardV1 := tokensapi.RemoteOutbounds{
		Outbounds: []tokensapi.RateLimitConfig{
			{RateLimit: tokensapi.RateLimiterConfigFloatInput{IsEnabled: true, Capacity: 2002, Rate: 202}, FastFinality: false},
			{RateLimit: tokensapi.RateLimiterConfigFloatInput{IsEnabled: true, Capacity: 4040, Rate: 404}, FastFinality: true},
		},
	}

	// The initial remote chain config input for TokenExpansion
	remoteTowardV2 := tokensapi.RemoteChainConfig[*datastore.AddressRef, datastore.AddressRef]{OutboundRateLimits: outboundTowardV2.Outbounds, RemoteDecimals: decimalsV2}
	remoteTowardV1 := tokensapi.RemoteChainConfig[*datastore.AddressRef, datastore.AddressRef]{OutboundRateLimits: outboundTowardV1.Outbounds, RemoteDecimals: decimalsV1}

	// Define datastore filters to fetch deployed token pools more easily later
	fltrV1 := datastore.AddressRef{ChainSelector: selV1, Type: datastore.ContractType(bnmOpsV1_6_1.ContractType), Version: bnmOpsV1_6_1.Version}
	fltrV2 := datastore.AddressRef{ChainSelector: selV2, Type: datastore.ContractType(bnmOpsV2_0_0.ContractType), Version: bnmOpsV2_0_0.Version}

	// Edge case: if the user specifies both the FF and default rate limits for pre-V2
	// pools, then only the default RL should be applied and FF should be skipped. For
	// the V2 pools, both FF and default should be applied as normal.
	t.Run("tp_applies_default_tprl_and_skips_FF_on_v1", func(t *testing.T) {
		t.Parallel()

		// Setup test env
		ev, err := environment.New(t.Context(), environment.WithEVMSimulated(t, []uint64{selV1, selV2}))
		require.NoError(t, err)
		ds := datastore.NewMemoryDataStore()
		DeployChainContractsV2_0_0(t, ev, ds, selV1)
		DeployChainContractsV2_0_0(t, ev, ds, selV2)
		ev.DataStore = ds.Seal()

		// We're skipping MCMS setup so we need to use the deployer key
		deployerV1 := ev.BlockChains.EVMChains()[selV1].DeployerKey.From
		deployerV2 := ev.BlockChains.EVMChains()[selV2].DeployerKey.From
		clientV1 := ev.BlockChains.EVMChains()[selV1].Client
		clientV2 := ev.BlockChains.EVMChains()[selV2].Client

		// Deploy tokens and pools
		outTE, err := tokensapi.TokenExpansion().Apply(*ev, tokensapi.TokenExpansionInput{
			ChainAdapterVersion: cciputils.Version_2_0_0,
			MCMS:                mcms.Input{},
			TokenExpansionInputPerChain: map[uint64]tokensapi.TokenExpansionInputPerChain{
				selV1: {
					SkipOwnershipTransfer: true,
					TokenPoolVersion:      bnmOpsV1_6_1.Version,
					DeployTokenInput: &tokensapi.DeployTokenInput{
						Name:          tokenSymbl + "_CASE1",
						Symbol:        tokenSymbl + "_CASE1",
						Decimals:      decimalsV1,
						ExternalAdmin: deployerV1.Hex(),
						CCIPAdmin:     deployerV1.Hex(),
						Type:          bnmERC20ops.ContractType,
					},
					DeployTokenPoolInput: &tokensapi.DeployTokenPoolInput{
						PoolType:           string(bnmOpsV1_6_1.ContractType),
						TokenPoolQualifier: "",
					},
					TokenTransferConfig: &tokensapi.TokenTransferConfig{
						RemoteChains: map[uint64]tokensapi.RemoteChainConfig[*datastore.AddressRef, datastore.AddressRef]{selV2: remoteTowardV2},
					},
				},
				selV2: {
					SkipOwnershipTransfer: true,
					TokenPoolVersion:      bnmOpsV2_0_0.Version,
					DeployTokenInput: &tokensapi.DeployTokenInput{
						Name:          tokenSymbl + "_CASE1",
						Symbol:        tokenSymbl + "_CASE1",
						Decimals:      decimalsV2,
						ExternalAdmin: deployerV2.Hex(),
						CCIPAdmin:     deployerV2.Hex(),
						Type:          bnmERC20ops.ContractType,
					},
					DeployTokenPoolInput: &tokensapi.DeployTokenPoolInput{
						PoolType:           string(bnmOpsV2_0_0.ContractType),
						TokenPoolQualifier: "",
					},
					TokenTransferConfig: &tokensapi.TokenTransferConfig{
						RemoteChains: map[uint64]tokensapi.RemoteChainConfig[*datastore.AddressRef, datastore.AddressRef]{selV1: remoteTowardV1},
					},
				},
			},
		})
		require.NoError(t, err)
		MergeAddresses(t, ev, outTE.DataStore)

		// Fetch token pool refs from datastore
		poolV1, err := datastore_utils.FindAndFormatRef(ev.DataStore, fltrV1, selV1, evm_datastore_utils.ToEVMAddress)
		require.NoError(t, err)
		poolV2, err := datastore_utils.FindAndFormatRef(ev.DataStore, fltrV2, selV2, evm_datastore_utils.ToEVMAddress)
		require.NoError(t, err)

		// Setting a FF and default rate limit on a pre-V2 token pool should only change the pool's
		// default rate limits. For the V2 pool, the FF and default rate limits should be updated.
		newRateLimitsTowardV2 := tokensapi.RemoteOutbounds{
			Outbounds: []tokensapi.RateLimitConfig{
				{RateLimit: tokensapi.RateLimiterConfigFloatInput{IsEnabled: true, Capacity: 5005, Rate: 505}, FastFinality: false},
				{RateLimit: tokensapi.RateLimiterConfigFloatInput{IsEnabled: true, Capacity: 1111, Rate: 111}, FastFinality: true},
			},
		}
		newRateLimitsTowardV1 := tokensapi.RemoteOutbounds{
			Outbounds: []tokensapi.RateLimitConfig{
				{RateLimit: tokensapi.RateLimiterConfigFloatInput{IsEnabled: true, Capacity: 6060, Rate: 606}, FastFinality: false},
				{RateLimit: tokensapi.RateLimiterConfigFloatInput{IsEnabled: true, Capacity: 2222, Rate: 222}, FastFinality: true},
			},
		}
		_, err = tokensapi.SetTokenPoolRateLimits().Apply(*ev, tokensapi.TPRLInput{
			MCMS: mcms.Input{},
			Configs: map[uint64]tokensapi.TPRLConfig{
				selV1: {
					TokenRef:        datastore.AddressRef{Qualifier: tokenSymbl + "_CASE1"},
					TokenPoolRef:    datastore.AddressRef{Address: poolV1.Hex()},
					RemoteOutbounds: map[uint64]tokensapi.RemoteOutbounds{selV2: newRateLimitsTowardV2},
				},
				selV2: {
					TokenRef:        datastore.AddressRef{Qualifier: tokenSymbl + "_CASE1"},
					TokenPoolRef:    datastore.AddressRef{Address: poolV2.Hex()},
					RemoteOutbounds: map[uint64]tokensapi.RemoteOutbounds{selV1: newRateLimitsTowardV1},
				},
			},
		})
		require.NoError(t, err)

		// Take a snapshot of the post-update rate limits to compare against the expected values
		postDefaultRateLimitV2, postFastFinalityRateLimitV2 := getRateLimits(t, cciputils.Version_2_0_0, poolV2, clientV2, selV1)
		postDefaultRateLimitV1, _ := getRateLimits(t, cciputils.Version_1_6_1, poolV1, clientV1, selV2)

		// For the V1 pool, the default rate limit should be updated
		v1ExpectedOutboundDefault, ok := newRateLimitsTowardV2.DefaultBucket()
		require.True(t, ok)
		v1ExpectedInboundDefault, ok := newRateLimitsTowardV1.DefaultBucket()
		require.True(t, ok)
		validateScaledTPRLBucket(t, "default v1_pool", decimalsV1, postDefaultRateLimitV1, v1ExpectedOutboundDefault.RateLimit, v1ExpectedInboundDefault.RateLimit)

		// For the V2 pool, the default rate limit should be updated
		v2ExpectedOutboundDefault, ok := newRateLimitsTowardV1.DefaultBucket()
		require.True(t, ok)
		v2ExpectedInboundDefault, ok := newRateLimitsTowardV2.DefaultBucket()
		require.True(t, ok)
		validateScaledTPRLBucket(t, "default v2_pool", decimalsV2, postDefaultRateLimitV2, v2ExpectedOutboundDefault.RateLimit, v2ExpectedInboundDefault.RateLimit)

		// For the V2 pool, the FF rate limit should be updated
		v2ExpectedOutboundFF, ok := newRateLimitsTowardV1.FastFinalityBucket()
		require.True(t, ok)
		v2ExpectedInboundFF, ok := newRateLimitsTowardV2.FastFinalityBucket()
		require.True(t, ok)
		validateScaledTPRLBucket(t, "fast_finality v2_pool", decimalsV2, postFastFinalityRateLimitV2, v2ExpectedOutboundFF.RateLimit, v2ExpectedInboundFF.RateLimit)
	})

	// Edge case: if the user specifies FF rate limits but no default rate limits for
	// a pre-V2 pool, then the FF rate limit should be ignored and the default should
	// remain unchanged. For the V2 pool the FF rate limit should be applied normally
	// and the default lane should remain unchanged.
	t.Run("tp_ff_only_leaves_v1_default_unchanged", func(t *testing.T) {
		t.Parallel()

		// Setup test env
		ev, err := environment.New(t.Context(), environment.WithEVMSimulated(t, []uint64{selV1, selV2}))
		require.NoError(t, err)
		ds := datastore.NewMemoryDataStore()
		DeployChainContractsV2_0_0(t, ev, ds, selV1)
		DeployChainContractsV2_0_0(t, ev, ds, selV2)
		ev.DataStore = ds.Seal()

		// We're skipping MCMS setup so we need to use the deployer key
		deployerV1 := ev.BlockChains.EVMChains()[selV1].DeployerKey.From
		deployerV2 := ev.BlockChains.EVMChains()[selV2].DeployerKey.From
		clientV1 := ev.BlockChains.EVMChains()[selV1].Client
		clientV2 := ev.BlockChains.EVMChains()[selV2].Client

		// Deploy tokens and pools
		outTE, err := tokensapi.TokenExpansion().Apply(*ev, tokensapi.TokenExpansionInput{
			ChainAdapterVersion: cciputils.Version_2_0_0,
			MCMS:                mcms.Input{},
			TokenExpansionInputPerChain: map[uint64]tokensapi.TokenExpansionInputPerChain{
				selV1: {
					SkipOwnershipTransfer: true,
					TokenPoolVersion:      bnmOpsV1_6_1.Version,
					DeployTokenInput: &tokensapi.DeployTokenInput{
						Name:          tokenSymbl + "_CASE2",
						Symbol:        tokenSymbl + "_CASE2",
						Decimals:      decimalsV1,
						ExternalAdmin: deployerV1.Hex(),
						CCIPAdmin:     deployerV1.Hex(),
						Type:          bnmERC20ops.ContractType,
					},
					DeployTokenPoolInput: &tokensapi.DeployTokenPoolInput{
						PoolType:           string(bnmOpsV1_6_1.ContractType),
						TokenPoolQualifier: "",
					},
					TokenTransferConfig: &tokensapi.TokenTransferConfig{
						RemoteChains: map[uint64]tokensapi.RemoteChainConfig[*datastore.AddressRef, datastore.AddressRef]{selV2: remoteTowardV2},
					},
				},
				selV2: {
					SkipOwnershipTransfer: true,
					TokenPoolVersion:      bnmOpsV2_0_0.Version,
					DeployTokenInput: &tokensapi.DeployTokenInput{
						Name:          tokenSymbl + "_CASE2",
						Symbol:        tokenSymbl + "_CASE2",
						Decimals:      decimalsV2,
						ExternalAdmin: deployerV2.Hex(),
						CCIPAdmin:     deployerV2.Hex(),
						Type:          bnmERC20ops.ContractType,
					},
					DeployTokenPoolInput: &tokensapi.DeployTokenPoolInput{
						PoolType:           string(bnmOpsV2_0_0.ContractType),
						TokenPoolQualifier: "",
					},
					TokenTransferConfig: &tokensapi.TokenTransferConfig{
						RemoteChains: map[uint64]tokensapi.RemoteChainConfig[*datastore.AddressRef, datastore.AddressRef]{selV1: remoteTowardV1},
					},
				},
			},
		})
		require.NoError(t, err)
		MergeAddresses(t, ev, outTE.DataStore)

		// Fetch token pool refs from datastore
		poolV1, err := datastore_utils.FindAndFormatRef(ev.DataStore, fltrV1, selV1, evm_datastore_utils.ToEVMAddress)
		require.NoError(t, err)
		poolV2, err := datastore_utils.FindAndFormatRef(ev.DataStore, fltrV2, selV2, evm_datastore_utils.ToEVMAddress)
		require.NoError(t, err)

		// Take a snapshot of the pre-update rate limits to compare against the expected values
		preDefaultRateLimitV1, _ := getRateLimits(t, cciputils.Version_1_6_1, poolV1, clientV1, selV2)
		preDefaultRateLimitV2, _ := getRateLimits(t, cciputils.Version_2_0_0, poolV2, clientV2, selV1)

		// Both pools' default rate limits should be set based on the initial token expansion input
		v1ExpectedOutboundDefault, ok := remoteTowardV2.GetOutboundRateLimitBuckets().DefaultBucket()
		require.True(t, ok)
		v1ExpectedInboundDefault, ok := remoteTowardV1.GetOutboundRateLimitBuckets().DefaultBucket()
		require.True(t, ok)
		validateScaledTPRLBucket(t, "V1 pool default rate limits before FF-only TPRL",
			decimalsV1, preDefaultRateLimitV1,
			v1ExpectedOutboundDefault.RateLimit, v1ExpectedInboundDefault.RateLimit,
		)
		validateScaledTPRLBucket(t, "V2 pool default rate limits before FF-only TPRL",
			decimalsV2, preDefaultRateLimitV2,
			v1ExpectedInboundDefault.RateLimit, v1ExpectedOutboundDefault.RateLimit, // note the flip since it's from the perspective of the V2 pool
		)

		// Setting a FF rate limit on a pre-V2 token pool should not change the pool's existing
		// onchain rate limits. For the V2 pool, only the FF rate limit should be updated.
		newRateLimitsTowardV2 := tokensapi.RemoteOutbounds{
			Outbounds: []tokensapi.RateLimitConfig{
				{RateLimit: tokensapi.RateLimiterConfigFloatInput{IsEnabled: true, Capacity: 1111, Rate: 111}, FastFinality: true},
			},
		}
		newRateLimitsTowardV1 := tokensapi.RemoteOutbounds{
			Outbounds: []tokensapi.RateLimitConfig{
				{RateLimit: tokensapi.RateLimiterConfigFloatInput{IsEnabled: true, Capacity: 2222, Rate: 222}, FastFinality: true},
			},
		}
		_, err = tokensapi.SetTokenPoolRateLimits().Apply(*ev, tokensapi.TPRLInput{
			MCMS: mcms.Input{},
			Configs: map[uint64]tokensapi.TPRLConfig{
				selV1: {
					TokenRef:        datastore.AddressRef{Qualifier: tokenSymbl + "_CASE2"},
					TokenPoolRef:    datastore.AddressRef{Address: poolV1.Hex()},
					RemoteOutbounds: map[uint64]tokensapi.RemoteOutbounds{selV2: newRateLimitsTowardV2},
				},
				selV2: {
					TokenRef:        datastore.AddressRef{Qualifier: tokenSymbl + "_CASE2"},
					TokenPoolRef:    datastore.AddressRef{Address: poolV2.Hex()},
					RemoteOutbounds: map[uint64]tokensapi.RemoteOutbounds{selV1: newRateLimitsTowardV1},
				},
			},
		})
		require.NoError(t, err)

		// The default rate limits for both pools should remain untouched
		postDefaultRateLimitV2, postFastFinalityRateLimitV2 := getRateLimits(t, cciputils.Version_2_0_0, poolV2, clientV2, selV1)
		postDefaultRateLimitV1, _ := getRateLimits(t, cciputils.Version_1_6_1, poolV1, clientV1, selV2)
		assertTPRLBucketRateLimiterConfigsUnchanged(t, "v1 default lane after FF-only TPRL", preDefaultRateLimitV1, postDefaultRateLimitV1)
		assertTPRLBucketRateLimiterConfigsUnchanged(t, "v2 default lane after FF-only TPRL", preDefaultRateLimitV2, postDefaultRateLimitV2)

		// The V2 pool's fast finality rate limits should be updated
		v2ExpectedOutboundFF, ok := newRateLimitsTowardV1.FastFinalityBucket()
		require.True(t, ok)
		v2ExpectedInboundFF, ok := newRateLimitsTowardV2.FastFinalityBucket()
		require.True(t, ok)
		validateScaledTPRLBucket(t, "fast_finality after FF-only TPRL", decimalsV2, postFastFinalityRateLimitV2, v2ExpectedOutboundFF.RateLimit, v2ExpectedInboundFF.RateLimit)
	})
}

func assertTPRLBucketRateLimiterConfigsUnchanged(t *testing.T, label string, before, after *tokensapi.TPRLRateLimitBucket) {
	t.Helper()
	require.NotNil(t, before)
	require.NotNil(t, after)

	// Check outbound side
	outboundLab := label + " outbound"
	beforeOB, afterOB := before.OutboundRateLimiterConfig, after.OutboundRateLimiterConfig
	require.Equal(t, beforeOB.IsEnabled, afterOB.IsEnabled, "%s IsEnabled", outboundLab)
	RequireBigIntsEqual(t, beforeOB.Capacity, afterOB.Capacity, outboundLab+" capacity")
	RequireBigIntsEqual(t, beforeOB.Rate, afterOB.Rate, outboundLab+" rate")

	// Check inbound side
	inboundLabel := label + " inbound"
	beforeIB, afterIB := before.InboundRateLimiterConfig, after.InboundRateLimiterConfig
	require.Equal(t, beforeIB.IsEnabled, afterIB.IsEnabled, "%s IsEnabled", inboundLabel)
	RequireBigIntsEqual(t, beforeIB.Capacity, afterIB.Capacity, inboundLabel+" capacity")
	RequireBigIntsEqual(t, beforeIB.Rate, afterIB.Rate, inboundLabel+" rate")
}

func validateScaledTPRLBucket(t *testing.T, label string, localDecimals uint8, bucket *tokensapi.TPRLRateLimitBucket, expOutbound, expInbound tokensapi.RateLimiterConfigFloatInput) {
	t.Helper()
	require.NotNil(t, bucket)

	// Outbound requires no extra percent
	require.Equal(t, expOutbound.IsEnabled, bucket.OutboundRateLimiterConfig.IsEnabled, "%s outbound enabled", label)
	outCap := tokensapi.ScaleFloatToBigInt(expOutbound.Capacity, int(localDecimals), 0)
	outRate := tokensapi.ScaleFloatToBigInt(expOutbound.Rate, int(localDecimals), 0)
	RequireBigIntsEqual(t, outCap, bucket.OutboundRateLimiterConfig.Capacity, label+" outbound capacity")
	RequireBigIntsEqual(t, outRate, bucket.OutboundRateLimiterConfig.Rate, label+" outbound rate")

	// Inbound should be +10% of counterpart's outbound
	require.Equal(t, expInbound.IsEnabled, bucket.InboundRateLimiterConfig.IsEnabled, "%s inbound enabled", label)
	inCap := tokensapi.ScaleFloatToBigInt(expInbound.Capacity, int(localDecimals), 0.10)
	inRate := tokensapi.ScaleFloatToBigInt(expInbound.Rate, int(localDecimals), 0.10)
	RequireBigIntsEqual(t, inCap, bucket.InboundRateLimiterConfig.Capacity, label+" inbound capacity")
	RequireBigIntsEqual(t, inRate, bucket.InboundRateLimiterConfig.Rate, label+" inbound rate")
}

func getRateLimits(t *testing.T, version *semver.Version, address common.Address, backend bind.ContractBackend, destSel uint64) (*tokensapi.TPRLRateLimitBucket, *tokensapi.TPRLRateLimitBucket) {
	t.Helper()

	opts := &bind.CallOpts{Context: t.Context()}
	switch {
	case cciputils.Version_2_0_0.Equal(version):
		tp, err := tokenpoolV2_0_0.NewTokenPool(address, backend)
		require.NoError(t, err)
		stateDefault, err := tp.GetCurrentRateLimiterState(opts, destSel, false)
		require.NoError(t, err)
		stateFF, err := tp.GetCurrentRateLimiterState(opts, destSel, true)
		require.NoError(t, err)
		defaultBucket := &tokensapi.TPRLRateLimitBucket{
			FastFinality: false,
			OutboundRateLimiterConfig: tokensapi.RateLimiterConfig{
				IsEnabled: stateDefault.OutboundRateLimiterState.IsEnabled,
				Capacity:  stateDefault.OutboundRateLimiterState.Capacity,
				Rate:      stateDefault.OutboundRateLimiterState.Rate,
			},
			InboundRateLimiterConfig: tokensapi.RateLimiterConfig{
				IsEnabled: stateDefault.InboundRateLimiterState.IsEnabled,
				Capacity:  stateDefault.InboundRateLimiterState.Capacity,
				Rate:      stateDefault.InboundRateLimiterState.Rate,
			},
		}
		ffBucket := &tokensapi.TPRLRateLimitBucket{
			FastFinality: true,
			OutboundRateLimiterConfig: tokensapi.RateLimiterConfig{
				IsEnabled: stateFF.OutboundRateLimiterState.IsEnabled,
				Capacity:  stateFF.OutboundRateLimiterState.Capacity,
				Rate:      stateFF.OutboundRateLimiterState.Rate,
			},
			InboundRateLimiterConfig: tokensapi.RateLimiterConfig{
				IsEnabled: stateFF.InboundRateLimiterState.IsEnabled,
				Capacity:  stateFF.InboundRateLimiterState.Capacity,
				Rate:      stateFF.InboundRateLimiterState.Rate,
			},
		}
		return defaultBucket, ffBucket

	case cciputils.Version_1_6_1.Equal(version):
		tp, err := tokenpoolV1_6_1.NewTokenPool(address, backend)
		require.NoError(t, err)
		outbound, err := tp.GetCurrentOutboundRateLimiterState(opts, destSel)
		require.NoError(t, err)
		inbound, err := tp.GetCurrentInboundRateLimiterState(opts, destSel)
		require.NoError(t, err)
		return &tokensapi.TPRLRateLimitBucket{
			FastFinality: false,
			OutboundRateLimiterConfig: tokensapi.RateLimiterConfig{
				IsEnabled: outbound.IsEnabled, Capacity: outbound.Capacity, Rate: outbound.Rate,
			},
			InboundRateLimiterConfig: tokensapi.RateLimiterConfig{
				IsEnabled: inbound.IsEnabled, Capacity: inbound.Capacity, Rate: inbound.Rate,
			},
		}, nil

	default:
		require.FailNow(t, fmt.Sprintf("unsupported token pool version for fetching rate limits: %s", version.String()))
		return nil, nil
	}
}

// forceSimGasLimit bumps the deployer key's GasLimit on every EVM chain in the test environment.
// The simulated chain's gas estimator under-charges SSTORE refunds when a setRateLimitConfig call
// writes the same value back to storage, causing the tx to run out of gas. Setting a manual
// GasLimit bypasses estimation and matches real-chain behavior. The fix is test-only — production
// chains return a correct estimate.
func forceSimGasLimit(env *cldf_deployment.Environment, gasLimit uint64) {
	for _, chain := range env.BlockChains.EVMChains() {
		if chain.DeployerKey != nil {
			chain.DeployerKey.GasLimit = gasLimit
		}
	}
}

// TestTPRL_OutboundOnly_Verify exercises the verify-time rules for the OutboundOnly flag:
// the counterpart's RemoteOutbounds[A] entry is not required, but the counterpart's TPRLConfig
// must still exist, and the local lane must have at least one outbound bucket.
func TestTPRL_OutboundOnly_Verify(t *testing.T) {
	selA := chainsel.TEST_90000001.Selector
	selB := chainsel.TEST_90000002.Selector
	env, err := environment.New(t.Context(), environment.WithEVMSimulated(t, []uint64{selA}))
	require.NoError(t, err)

	cs := tokensapi.SetTokenPoolRateLimits()

	validOutbound := tokensapi.RemoteOutbounds{
		OutboundOnly: true,
		Outbounds: []tokensapi.RateLimitConfig{
			{RateLimit: tokensapi.RateLimiterConfigFloatInput{IsEnabled: true, Capacity: 100, Rate: 10}, FastFinality: false},
		},
	}

	t.Run("passes_without_counterpart_remote_outbounds", func(t *testing.T) {
		err := cs.VerifyPreconditions(*env, tokensapi.TPRLInput{
			MCMS: mcms.Input{},
			Configs: map[uint64]tokensapi.TPRLConfig{
				selA: {RemoteOutbounds: map[uint64]tokensapi.RemoteOutbounds{selB: validOutbound}},
				selB: {}, // refs-only config — RemoteOutbounds[selA] intentionally absent
			},
		})
		require.NoError(t, err)
	})

	t.Run("rejects_when_counterpart_config_missing", func(t *testing.T) {
		err := cs.VerifyPreconditions(*env, tokensapi.TPRLInput{
			MCMS: mcms.Input{},
			Configs: map[uint64]tokensapi.TPRLConfig{
				selA: {RemoteOutbounds: map[uint64]tokensapi.RemoteOutbounds{selB: validOutbound}},
			},
		})
		require.Error(t, err)
		require.Contains(t, err.Error(), "no config provided for remote chain")
	})

	t.Run("rejects_when_no_outbound_buckets", func(t *testing.T) {
		emptyOutboundOnly := tokensapi.RemoteOutbounds{OutboundOnly: true}
		err := cs.VerifyPreconditions(*env, tokensapi.TPRLInput{
			MCMS: mcms.Input{},
			Configs: map[uint64]tokensapi.TPRLConfig{
				selA: {RemoteOutbounds: map[uint64]tokensapi.RemoteOutbounds{selB: emptyOutboundOnly}},
				selB: {},
			},
		})
		require.Error(t, err)
		require.Contains(t, err.Error(), "no outbound buckets")
	})

	t.Run("allows_asymmetric_outbound_only_on_one_side", func(t *testing.T) {
		// A→B is OutboundOnly while B→A is a normal symmetric lane. The symmetry checks only
		// apply when neither side is OutboundOnly; OutboundOnly on A→B should not force B→A
		// to also be OutboundOnly.
		err := cs.VerifyPreconditions(*env, tokensapi.TPRLInput{
			MCMS: mcms.Input{},
			Configs: map[uint64]tokensapi.TPRLConfig{
				selA: {RemoteOutbounds: map[uint64]tokensapi.RemoteOutbounds{selB: validOutbound}},
				selB: {RemoteOutbounds: map[uint64]tokensapi.RemoteOutbounds{selA: {
					Outbounds: []tokensapi.RateLimitConfig{
						{RateLimit: tokensapi.RateLimiterConfigFloatInput{IsEnabled: true, Capacity: 200, Rate: 20}, FastFinality: false},
					},
				}}},
			},
		})
		require.NoError(t, err)
	})
}

// outboundOnlyV2Setup deploys a pair of v2.0.0 pools, performs an initial symmetric TPRL apply
// to seed the on-chain rate limits, and returns the resolved pool refs and decimals. Returning
// these makes the actual OutboundOnly cases in the table-driven tests trivial.
func outboundOnlyV2Setup(t *testing.T, tokenSymb string, decimalsA, decimalsB uint8, initialAB, initialBA tokensapi.RateLimitConfig) (
	env *cldf_deployment.Environment,
	selA, selB uint64,
	poolARef, poolBRef datastore.AddressRef,
	clientA, clientB bind.ContractBackend,
) {
	t.Helper()

	selA = chainsel.TEST_90000001.Selector
	selB = chainsel.TEST_90000002.Selector
	e, err := environment.New(t.Context(), environment.WithEVMSimulated(t, []uint64{selA, selB}))
	require.NoError(t, err)
	env = e

	cumulative := datastore.NewMemoryDataStore()
	DeployChainContractsV2_0_0(t, env, cumulative, selA)
	DeployChainContractsV2_0_0(t, env, cumulative, selB)
	env.DataStore = cumulative.Seal()

	disabledOutbound := tokensapi.RateLimiterConfigFloatInput{IsEnabled: false}
	deployerA := env.BlockChains.EVMChains()[selA].DeployerKey.From
	deployerB := env.BlockChains.EVMChains()[selB].DeployerKey.From
	clientA = env.BlockChains.EVMChains()[selA].Client
	clientB = env.BlockChains.EVMChains()[selB].Client

	expansionOut, err := tokensapi.TokenExpansion().Apply(*env, tokensapi.TokenExpansionInput{
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
					PoolType: string(bnmOpsV2_0_0.ContractType), TokenPoolQualifier: "",
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
					PoolType: string(bnmOpsV2_0_0.ContractType), TokenPoolQualifier: "",
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
	MergeAddresses(t, env, expansionOut.DataStore)

	fltrA := datastore.AddressRef{ChainSelector: selA, Type: datastore.ContractType(bnmOpsV2_0_0.ContractType), Version: bnmOpsV2_0_0.Version, Qualifier: ""}
	poolA, err := datastore_utils.FindAndFormatRef(env.DataStore, fltrA, selA, evm_datastore_utils.ToEVMAddress)
	require.NoError(t, err)
	fltrB := datastore.AddressRef{ChainSelector: selB, Type: datastore.ContractType(bnmOpsV2_0_0.ContractType), Version: bnmOpsV2_0_0.Version, Qualifier: ""}
	poolB, err := datastore_utils.FindAndFormatRef(env.DataStore, fltrB, selB, evm_datastore_utils.ToEVMAddress)
	require.NoError(t, err)

	poolARef = datastore.AddressRef{Address: poolA.Hex()}
	poolBRef = datastore.AddressRef{Address: poolB.Hex()}

	// Seed the on-chain rate limits with a normal (non-OutboundOnly) symmetric apply so that
	// subsequent OutboundOnly tests have something to validate against.
	_, err = tokensapi.SetTokenPoolRateLimits().Apply(*env, tokensapi.TPRLInput{
		MCMS: mcms.Input{},
		Configs: map[uint64]tokensapi.TPRLConfig{
			selA: {
				ChainAdapterVersion: cciputils.Version_2_0_0,
				TokenRef:            datastore.AddressRef{Qualifier: tokenSymb},
				TokenPoolRef:        poolARef,
				RemoteOutbounds: map[uint64]tokensapi.RemoteOutbounds{selB: {
					Outbounds: []tokensapi.RateLimitConfig{initialAB},
				}},
			},
			selB: {
				ChainAdapterVersion: cciputils.Version_2_0_0,
				TokenRef:            datastore.AddressRef{Qualifier: tokenSymb},
				TokenPoolRef:        poolBRef,
				RemoteOutbounds: map[uint64]tokensapi.RemoteOutbounds{selA: {
					Outbounds: []tokensapi.RateLimitConfig{initialBA},
				}},
			},
		},
	})
	require.NoError(t, err)

	return env, selA, selB, poolARef, poolBRef, clientA, clientB
}

// TestTPRL_OutboundOnly_AppliesAndLeavesCounterpartUntouchedV2 verifies that an OutboundOnly
// apply on chain A updates A's outbound, leaves A's inbound and chain B's pool untouched, and
// passes the +110% on-chain validation when the new outbound stays within counterpart headroom.
func TestTPRL_OutboundOnly_AppliesAndLeavesCounterpartUntouchedV2(t *testing.T) {
	const tokenSymb = "TPRL_OBO_OK"
	const decimalsA, decimalsB = uint8(18), uint8(6)

	initialAB := tokensapi.RateLimitConfig{RateLimit: tokensapi.RateLimiterConfigFloatInput{IsEnabled: true, Capacity: 1000, Rate: 100}, FastFinality: false}
	initialBA := tokensapi.RateLimitConfig{RateLimit: tokensapi.RateLimiterConfigFloatInput{IsEnabled: true, Capacity: 2000, Rate: 200}, FastFinality: false}

	env, selA, selB, poolARef, poolBRef, clientA, clientB := outboundOnlyV2Setup(t, tokenSymb, decimalsA, decimalsB, initialAB, initialBA)

	// Snapshot the seeded state. Chain B's inbound (from A) is 1000 * 1.10 in B's decimals,
	// so the new A outbound can be at most 1000 without breaching the +110% rule.
	preA, _ := getRateLimits(t, cciputils.Version_2_0_0, common.HexToAddress(poolARef.Address), clientA, selB)
	preB, _ := getRateLimits(t, cciputils.Version_2_0_0, common.HexToAddress(poolBRef.Address), clientB, selA)
	require.True(t, preA.OutboundRateLimiterConfig.IsEnabled, "seed: A→B outbound should be enabled")
	require.True(t, preB.InboundRateLimiterConfig.IsEnabled, "seed: B's inbound from A should be enabled")

	// Sim-only workaround: passing through the on-chain inbound writes the same uint128 back
	// to storage, which trips the simulated chain's gas estimator. Force a manual gas limit.
	forceSimGasLimit(env, 5_000_000)

	// Set A's outbound to 500 (half the seed): requiredInbound = 500*1.10 = 550 (<= 1100).
	newA := tokensapi.RateLimitConfig{RateLimit: tokensapi.RateLimiterConfigFloatInput{IsEnabled: true, Capacity: 500, Rate: 50}, FastFinality: false}
	_, err := tokensapi.SetTokenPoolRateLimits().Apply(*env, tokensapi.TPRLInput{
		MCMS: mcms.Input{},
		Configs: map[uint64]tokensapi.TPRLConfig{
			selA: {
				ChainAdapterVersion: cciputils.Version_2_0_0,
				TokenRef:            datastore.AddressRef{Qualifier: tokenSymb},
				TokenPoolRef:        poolARef,
				RemoteOutbounds: map[uint64]tokensapi.RemoteOutbounds{selB: {
					OutboundOnly: true,
					Outbounds:    []tokensapi.RateLimitConfig{newA},
				}},
			},
			// chain B has refs-only config; no RemoteOutbounds[selA] entry.
			selB: {
				ChainAdapterVersion: cciputils.Version_2_0_0,
				TokenRef:            datastore.AddressRef{Qualifier: tokenSymb},
				TokenPoolRef:        poolBRef,
			},
		},
	})
	require.NoError(t, err)

	postA, _ := getRateLimits(t, cciputils.Version_2_0_0, common.HexToAddress(poolARef.Address), clientA, selB)
	postB, _ := getRateLimits(t, cciputils.Version_2_0_0, common.HexToAddress(poolBRef.Address), clientB, selA)

	// A's outbound (to B) was rewritten to the new value.
	expOutCap := tokensapi.ScaleFloatToBigInt(newA.RateLimit.Capacity, int(decimalsA), 0)
	expOutRate := tokensapi.ScaleFloatToBigInt(newA.RateLimit.Rate, int(decimalsA), 0)
	require.True(t, postA.OutboundRateLimiterConfig.IsEnabled, "post: A→B outbound should remain enabled")
	RequireBigIntsEqual(t, expOutCap, postA.OutboundRateLimiterConfig.Capacity, "A new outbound capacity")
	RequireBigIntsEqual(t, expOutRate, postA.OutboundRateLimiterConfig.Rate, "A new outbound rate")

	// A's inbound (from B) is unchanged (pass-through).
	require.Equal(t, preA.InboundRateLimiterConfig.IsEnabled, postA.InboundRateLimiterConfig.IsEnabled, "A inbound IsEnabled unchanged")
	RequireBigIntsEqual(t, preA.InboundRateLimiterConfig.Capacity, postA.InboundRateLimiterConfig.Capacity, "A inbound capacity unchanged")
	RequireBigIntsEqual(t, preA.InboundRateLimiterConfig.Rate, postA.InboundRateLimiterConfig.Rate, "A inbound rate unchanged")

	// Chain B was not touched at all — both directions unchanged.
	assertTPRLBucketRateLimiterConfigsUnchanged(t, "B pool unchanged after OutboundOnly on A", preB, postB)
}

// TestTPRL_OutboundOnly_RejectsBelowThresholdV2 verifies the apply path returns an error when
// chain B's on-chain inbound is below the new outbound * 1.10 threshold and that nothing is
// written to chain A as a result.
func TestTPRL_OutboundOnly_RejectsBelowThresholdV2(t *testing.T) {
	const tokenSymb = "TPRL_OBO_REJ"
	const decimalsA, decimalsB = uint8(18), uint8(6)

	initialAB := tokensapi.RateLimitConfig{RateLimit: tokensapi.RateLimiterConfigFloatInput{IsEnabled: true, Capacity: 1000, Rate: 100}, FastFinality: false}
	initialBA := tokensapi.RateLimitConfig{RateLimit: tokensapi.RateLimiterConfigFloatInput{IsEnabled: true, Capacity: 2000, Rate: 200}, FastFinality: false}

	env, selA, selB, poolARef, poolBRef, clientA, _ := outboundOnlyV2Setup(t, tokenSymb, decimalsA, decimalsB, initialAB, initialBA)
	forceSimGasLimit(env, 5_000_000)

	preA, _ := getRateLimits(t, cciputils.Version_2_0_0, common.HexToAddress(poolARef.Address), clientA, selB)

	// Try to bump A's outbound to 1.5x the original. Chain B's existing inbound from A was
	// seeded at 1000 * 1.10 — anything above 1000 should fail the +110% check.
	overLimit := tokensapi.RateLimitConfig{RateLimit: tokensapi.RateLimiterConfigFloatInput{IsEnabled: true, Capacity: 1500, Rate: 150}, FastFinality: false}
	_, err := tokensapi.SetTokenPoolRateLimits().Apply(*env, tokensapi.TPRLInput{
		MCMS: mcms.Input{},
		Configs: map[uint64]tokensapi.TPRLConfig{
			selA: {
				ChainAdapterVersion: cciputils.Version_2_0_0,
				TokenRef:            datastore.AddressRef{Qualifier: tokenSymb},
				TokenPoolRef:        poolARef,
				RemoteOutbounds: map[uint64]tokensapi.RemoteOutbounds{selB: {
					OutboundOnly: true,
					Outbounds:    []tokensapi.RateLimitConfig{overLimit},
				}},
			},
			selB: {
				ChainAdapterVersion: cciputils.Version_2_0_0,
				TokenRef:            datastore.AddressRef{Qualifier: tokenSymb},
				TokenPoolRef:        poolBRef,
			},
		},
	})
	require.Error(t, err)
	require.Contains(t, err.Error(), "below required 110%")

	// A's pool must not have been modified by the failed apply.
	postA, _ := getRateLimits(t, cciputils.Version_2_0_0, common.HexToAddress(poolARef.Address), clientA, selB)
	assertTPRLBucketRateLimiterConfigsUnchanged(t, "A pool unchanged after rejected OutboundOnly", preA, postA)
}

// TestTPRL_OutboundOnly_V161 mirrors the v2.0.0 success-path test but on a v1.6.1 pool. v1.6.1
// pools only support a single default bucket; the test exercises the v1.x adapter path (PoolOps
// GetCurrentInboundRateLimit pass-through and EVMPoolAdapter.GetOnchainInboundRateLimit).
func TestTPRL_OutboundOnly_V161(t *testing.T) {
	const tokenSymb = "TPRL_OBO_V161"
	const decimalsA, decimalsB = uint8(18), uint8(6)

	selA := chainsel.TEST_90000001.Selector
	selB := chainsel.TEST_90000002.Selector
	env, err := environment.New(t.Context(), environment.WithEVMSimulated(t, []uint64{selA, selB}))
	require.NoError(t, err)

	cumulative := datastore.NewMemoryDataStore()
	DeployChainContractsV2_0_0(t, env, cumulative, selA)
	DeployChainContractsV2_0_0(t, env, cumulative, selB)
	env.DataStore = cumulative.Seal()

	disabledOutbound := tokensapi.RateLimiterConfigFloatInput{IsEnabled: false}
	deployerA := env.BlockChains.EVMChains()[selA].DeployerKey.From
	deployerB := env.BlockChains.EVMChains()[selB].DeployerKey.From
	clientA := env.BlockChains.EVMChains()[selA].Client
	clientB := env.BlockChains.EVMChains()[selB].Client

	expansionOut, err := tokensapi.TokenExpansion().Apply(*env, tokensapi.TokenExpansionInput{
		ChainAdapterVersion: cciputils.Version_2_0_0,
		MCMS:                mcms.Input{},
		TokenExpansionInputPerChain: map[uint64]tokensapi.TokenExpansionInputPerChain{
			selA: {
				SkipOwnershipTransfer: true,
				TokenPoolVersion:      bnmOpsV1_6_1.Version,
				DeployTokenInput: &tokensapi.DeployTokenInput{
					Name: tokenSymb, Symbol: tokenSymb, Decimals: decimalsA,
					ExternalAdmin: deployerA.Hex(), CCIPAdmin: deployerA.Hex(),
					Type: bnmERC20ops.ContractType,
				},
				DeployTokenPoolInput: &tokensapi.DeployTokenPoolInput{
					PoolType: string(bnmOpsV1_6_1.ContractType), TokenPoolQualifier: "",
				},
				TokenTransferConfig: &tokensapi.TokenTransferConfig{
					RemoteChains: map[uint64]tokensapi.RemoteChainConfig[*datastore.AddressRef, datastore.AddressRef]{
						selB: {OutboundRateLimiterConfig: &disabledOutbound, RemoteDecimals: decimalsB},
					},
				},
			},
			selB: {
				SkipOwnershipTransfer: true,
				TokenPoolVersion:      bnmOpsV1_6_1.Version,
				DeployTokenInput: &tokensapi.DeployTokenInput{
					Name: tokenSymb, Symbol: tokenSymb, Decimals: decimalsB,
					ExternalAdmin: deployerB.Hex(), CCIPAdmin: deployerB.Hex(),
					Type: bnmERC20ops.ContractType,
				},
				DeployTokenPoolInput: &tokensapi.DeployTokenPoolInput{
					PoolType: string(bnmOpsV1_6_1.ContractType), TokenPoolQualifier: "",
				},
				TokenTransferConfig: &tokensapi.TokenTransferConfig{
					RemoteChains: map[uint64]tokensapi.RemoteChainConfig[*datastore.AddressRef, datastore.AddressRef]{
						selA: {OutboundRateLimiterConfig: &disabledOutbound, RemoteDecimals: decimalsA},
					},
				},
			},
		},
	})
	require.NoError(t, err)
	MergeAddresses(t, env, expansionOut.DataStore)

	fltrA := datastore.AddressRef{ChainSelector: selA, Type: datastore.ContractType(bnmOpsV1_6_1.ContractType), Version: bnmOpsV1_6_1.Version, Qualifier: ""}
	poolA, err := datastore_utils.FindAndFormatRef(env.DataStore, fltrA, selA, evm_datastore_utils.ToEVMAddress)
	require.NoError(t, err)
	fltrB := datastore.AddressRef{ChainSelector: selB, Type: datastore.ContractType(bnmOpsV1_6_1.ContractType), Version: bnmOpsV1_6_1.Version, Qualifier: ""}
	poolB, err := datastore_utils.FindAndFormatRef(env.DataStore, fltrB, selB, evm_datastore_utils.ToEVMAddress)
	require.NoError(t, err)
	poolARef := datastore.AddressRef{Address: poolA.Hex()}
	poolBRef := datastore.AddressRef{Address: poolB.Hex()}

	// Seed symmetric rate limits.
	initialAB := tokensapi.RateLimitConfig{RateLimit: tokensapi.RateLimiterConfigFloatInput{IsEnabled: true, Capacity: 1000, Rate: 100}, FastFinality: false}
	initialBA := tokensapi.RateLimitConfig{RateLimit: tokensapi.RateLimiterConfigFloatInput{IsEnabled: true, Capacity: 2000, Rate: 200}, FastFinality: false}
	_, err = tokensapi.SetTokenPoolRateLimits().Apply(*env, tokensapi.TPRLInput{
		MCMS: mcms.Input{},
		Configs: map[uint64]tokensapi.TPRLConfig{
			selA: {
				TokenRef:        datastore.AddressRef{Qualifier: tokenSymb},
				TokenPoolRef:    poolARef,
				RemoteOutbounds: map[uint64]tokensapi.RemoteOutbounds{selB: {Outbounds: []tokensapi.RateLimitConfig{initialAB}}},
			},
			selB: {
				TokenRef:        datastore.AddressRef{Qualifier: tokenSymb},
				TokenPoolRef:    poolBRef,
				RemoteOutbounds: map[uint64]tokensapi.RemoteOutbounds{selA: {Outbounds: []tokensapi.RateLimitConfig{initialBA}}},
			},
		},
	})
	require.NoError(t, err)

	preA, _ := getRateLimits(t, cciputils.Version_1_6_1, poolA, clientA, selB)
	preB, _ := getRateLimits(t, cciputils.Version_1_6_1, poolB, clientB, selA)

	// Sim-only workaround: passing through the on-chain inbound writes the same uint128 back
	// to storage, which trips the simulated chain's gas estimator. Force a manual gas limit.
	forceSimGasLimit(env, 5_000_000)

	// Within-headroom OutboundOnly update: new outbound is half of original.
	newA := tokensapi.RateLimitConfig{RateLimit: tokensapi.RateLimiterConfigFloatInput{IsEnabled: true, Capacity: 500, Rate: 50}, FastFinality: false}
	_, err = tokensapi.SetTokenPoolRateLimits().Apply(*env, tokensapi.TPRLInput{
		MCMS: mcms.Input{},
		Configs: map[uint64]tokensapi.TPRLConfig{
			selA: {
				TokenRef:     datastore.AddressRef{Qualifier: tokenSymb},
				TokenPoolRef: poolARef,
				RemoteOutbounds: map[uint64]tokensapi.RemoteOutbounds{selB: {
					OutboundOnly: true,
					Outbounds:    []tokensapi.RateLimitConfig{newA},
				}},
			},
			selB: {
				TokenRef:     datastore.AddressRef{Qualifier: tokenSymb},
				TokenPoolRef: poolBRef,
			},
		},
	})
	require.NoError(t, err)

	postA, _ := getRateLimits(t, cciputils.Version_1_6_1, poolA, clientA, selB)
	postB, _ := getRateLimits(t, cciputils.Version_1_6_1, poolB, clientB, selA)

	// A's outbound updated; A's inbound unchanged.
	expOutCap := tokensapi.ScaleFloatToBigInt(newA.RateLimit.Capacity, int(decimalsA), 0)
	expOutRate := tokensapi.ScaleFloatToBigInt(newA.RateLimit.Rate, int(decimalsA), 0)
	RequireBigIntsEqual(t, expOutCap, postA.OutboundRateLimiterConfig.Capacity, "v161 A new outbound capacity")
	RequireBigIntsEqual(t, expOutRate, postA.OutboundRateLimiterConfig.Rate, "v161 A new outbound rate")
	RequireBigIntsEqual(t, preA.InboundRateLimiterConfig.Capacity, postA.InboundRateLimiterConfig.Capacity, "v161 A inbound capacity unchanged")
	RequireBigIntsEqual(t, preA.InboundRateLimiterConfig.Rate, postA.InboundRateLimiterConfig.Rate, "v161 A inbound rate unchanged")

	// Chain B's pool was not touched.
	assertTPRLBucketRateLimiterConfigsUnchanged(t, "v161 B pool unchanged after OutboundOnly on A", preB, postB)
}
