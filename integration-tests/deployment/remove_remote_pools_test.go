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
	bnmOpsV2_0_0 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/burn_mint_token_pool"
	evm_testsetup "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/testsetup"
	tokenpoolV2_0_0 "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v2_0_0/token_pool"
	"github.com/smartcontractkit/chainlink-ccip/deployment/finality"
	tokensapi "github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	cciputils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
)

func TestRemoveRemotePools_VerifyPreconditions(t *testing.T) {
	sel := chainsel.TEST_90000001.Selector
	dst := chainsel.TEST_90000002.Selector
	env, err := environment.New(t.Context(), environment.WithEVMSimulated(t, []uint64{sel}))
	require.NoError(t, err)

	cs := tokensapi.RemoveRemotePools()
	poolRef := datastore.AddressRef{Address: "0x1111111111111111111111111111111111111111"}
	remoteRef := datastore.AddressRef{Address: "0x2222222222222222222222222222222222222222"}

	singlePoolInput := func(remote tokensapi.RemotePoolToRemove) tokensapi.RemoveRemotePoolsInput {
		return tokensapi.RemoveRemotePoolsInput{
			MCMS: mcms.Input{},
			Pools: []tokensapi.RemoveRemotePoolsPerPool{{
				ChainSelector:       sel,
				Pool:                poolRef,
				RemotePoolsToRemove: []tokensapi.RemotePoolToRemove{remote},
			}},
		}
	}

	cases := []struct {
		name   string
		input  tokensapi.RemoveRemotePoolsInput
		errors []string
	}{
		{
			name:   "rejects_empty_input",
			input:  tokensapi.RemoveRemotePoolsInput{MCMS: mcms.Input{}},
			errors: []string{"at least one pool entry"},
		},
		{
			name: "rejects_empty_pool_ref",
			input: tokensapi.RemoveRemotePoolsInput{
				MCMS: mcms.Input{},
				Pools: []tokensapi.RemoveRemotePoolsPerPool{{
					ChainSelector:       sel,
					RemotePoolsToRemove: []tokensapi.RemotePoolToRemove{{Selector: dst, Remote: remoteRef}},
				}},
			},
			errors: []string{"empty pool ref"},
		},
		{
			name: "rejects_no_remote_pools",
			input: tokensapi.RemoveRemotePoolsInput{
				MCMS: mcms.Input{},
				Pools: []tokensapi.RemoveRemotePoolsPerPool{{
					ChainSelector: sel,
					Pool:          poolRef,
				}},
			},
			errors: []string{"no remote pools to remove"},
		},
		{
			name:   "rejects_remote_equal_to_local_chain",
			input:  singlePoolInput(tokensapi.RemotePoolToRemove{Selector: sel, Remote: remoteRef}),
			errors: []string{"must not equal the pool's own chain selector"},
		},
		{
			name:   "rejects_empty_remote_ref",
			input:  singlePoolInput(tokensapi.RemotePoolToRemove{Selector: dst}),
			errors: []string{"empty remote ref"},
		},
		{
			name: "rejects_duplicate_remote_selectors",
			input: tokensapi.RemoveRemotePoolsInput{
				MCMS: mcms.Input{},
				Pools: []tokensapi.RemoveRemotePoolsPerPool{{
					ChainSelector: sel,
					Pool:          poolRef,
					RemotePoolsToRemove: []tokensapi.RemotePoolToRemove{
						{Selector: dst, Remote: remoteRef},
						{Selector: dst, Remote: remoteRef},
					},
				}},
			},
			errors: []string{"duplicate remote chain selector"},
		},
		{
			name: "rejects_duplicate_pool_entries",
			input: tokensapi.RemoveRemotePoolsInput{
				MCMS: mcms.Input{},
				Pools: []tokensapi.RemoveRemotePoolsPerPool{
					{ChainSelector: sel, Pool: poolRef, RemotePoolsToRemove: []tokensapi.RemotePoolToRemove{{Selector: dst, Remote: remoteRef}}},
					{ChainSelector: sel, Pool: poolRef, RemotePoolsToRemove: []tokensapi.RemotePoolToRemove{{Selector: dst, Remote: remoteRef}}},
				},
			},
			errors: []string{"duplicate pool entry"},
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

// TestRemoveRemotePools_V2 removes a remote pool from a v2.0.0 pool and verifies the on-chain
// state, then confirms that removing an already-removed pool errors clearly.
func TestRemoveRemotePools_V2(t *testing.T) {
	tc := setupV2PoolsForConfigureImpl(t, "RRP_V2", false)

	// Sanity: pool A is connected to B.
	poolA, err := tokenpoolV2_0_0.NewTokenPool(tc.poolA, tc.clientA)
	require.NoError(t, err)
	remotePools, err := poolA.GetRemotePools(&bind.CallOpts{Context: t.Context()}, tc.selB)
	require.NoError(t, err)
	require.NotEmpty(t, remotePools, "pool A should have a remote pool for chain B before removal")

	input := tokensapi.RemoveRemotePoolsInput{
		MCMS: mcms.Input{},
		Pools: []tokensapi.RemoveRemotePoolsPerPool{{
			ChainSelector: tc.selA,
			Pool:          datastore.AddressRef{Address: tc.poolA.Hex()},
			RemotePoolsToRemove: []tokensapi.RemotePoolToRemove{{
				Selector: tc.selB,
				Remote:   datastore.AddressRef{Address: tc.poolB.Hex()},
			}},
		}},
	}
	require.NoError(t, tokensapi.RemoveRemotePools().VerifyPreconditions(*tc.env, input))
	_, err = tokensapi.RemoveRemotePools().Apply(*tc.env, input)
	require.NoError(t, err)

	remotePools, err = poolA.GetRemotePools(&bind.CallOpts{Context: t.Context()}, tc.selB)
	require.NoError(t, err)
	require.Empty(t, remotePools, "pool A should have no remote pool for chain B after removal")

	// Removing an already-removed pool must error clearly (not silently no-op).
	tc.env.OperationsBundle = evm_testsetup.BundleWithFreshReporter(tc.env.OperationsBundle)
	_, err = tokensapi.RemoveRemotePools().Apply(*tc.env, input)
	require.ErrorContains(t, err, "is not configured")
}

// TestRemoveRemotePools_PreV2 exercises remote pool removal on v1.5.1 and v1.6.1 pools.
func TestRemoveRemotePools_PreV2(t *testing.T) {
	t.Run("v1_5_1", func(t *testing.T) { testRemoveRemotePoolsPreV2(t, cciputils.Version_1_5_1) })
	t.Run("v1_6_1", func(t *testing.T) { testRemoveRemotePoolsPreV2(t, cciputils.Version_1_6_1) })
}

func testRemoveRemotePoolsPreV2(t *testing.T, version *semver.Version) {
	pair := setupLegacyConnectedBnMPair(t, version)

	// Sanity: old pool A is connected to B.
	oldPoolA, err := tokenpoolV2_0_0.NewTokenPool(pair.oldPoolAddrA, pair.env.BlockChains.EVMChains()[pair.selA].Client)
	require.NoError(t, err)
	remotePools, err := oldPoolA.GetRemotePools(&bind.CallOpts{Context: t.Context()}, pair.selB)
	require.NoError(t, err)
	require.NotEmpty(t, remotePools, "old pool A should have a remote pool for chain B before removal")

	input := tokensapi.RemoveRemotePoolsInput{
		MCMS: mcms.Input{},
		Pools: []tokensapi.RemoveRemotePoolsPerPool{{
			ChainSelector: pair.selA,
			Pool:          datastore.AddressRef{Address: pair.oldPoolAddrA.Hex()},
			RemotePoolsToRemove: []tokensapi.RemotePoolToRemove{{
				Selector: pair.selB,
				Remote:   datastore.AddressRef{Address: pair.oldPoolAddrB.Hex()},
			}},
		}},
	}
	require.NoError(t, tokensapi.RemoveRemotePools().VerifyPreconditions(*pair.env, input))
	_, err = tokensapi.RemoveRemotePools().Apply(*pair.env, input)
	require.NoError(t, err)

	remotePools, err = oldPoolA.GetRemotePools(&bind.CallOpts{Context: t.Context()}, pair.selB)
	require.NoError(t, err)
	require.Empty(t, remotePools, "old pool A should have no remote pool for chain B after removal")

	// Removing an already-removed pool must error clearly.
	pair.env.OperationsBundle = evm_testsetup.BundleWithFreshReporter(pair.env.OperationsBundle)
	_, err = tokensapi.RemoveRemotePools().Apply(*pair.env, input)
	require.ErrorContains(t, err, "is not configured")
}

// removeRemotePoolsTestEnv is the result of deploying a fully-connected v2.0.0 BurnMint pool
// across three chains (A, B, C). Each pool is connected to the other two, so every pool has two
// remote pool entries to remove.
type removeRemotePoolsTestEnv struct {
	env          *cldf_deployment.Environment
	selA, selB   uint64
	selC         uint64
	poolA, poolB common.Address
	poolC        common.Address
}

// setupV2PoolsForRemoveRemotePools deploys a fully-connected three-chain v2.0.0 BurnMint pool
// (A↔B, A↔C, B↔C) so tests can exercise partial remote-pool removal (removing only a subset of a
// pool's remote entries) and the error path for addresses that are not configured.
func setupV2PoolsForRemoveRemotePools(t *testing.T) removeRemotePoolsTestEnv {
	t.Helper()

	selA := chainsel.TEST_90000001.Selector
	selB := chainsel.TEST_90000002.Selector
	selC := chainsel.TEST_90000003.Selector
	e, err := environment.New(t.Context(), environment.WithEVMSimulated(t, []uint64{selA, selB, selC}))
	require.NoError(t, err)

	cumulative := datastore.NewMemoryDataStore()
	DeployChainContractsV2_0_0(t, e, cumulative, selA)
	DeployChainContractsV2_0_0(t, e, cumulative, selB)
	DeployChainContractsV2_0_0(t, e, cumulative, selC)
	e.DataStore = cumulative.Seal()

	disabledOutbound := tokensapi.RateLimiterConfigFloatInput{IsEnabled: false}
	deployers := map[uint64]common.Address{
		selA: e.BlockChains.EVMChains()[selA].DeployerKey.From,
		selB: e.BlockChains.EVMChains()[selB].DeployerKey.From,
		selC: e.BlockChains.EVMChains()[selC].DeployerKey.From,
	}

	// Build a fully-connected v2.0.0 BurnMint pool: each chain configures the other two as remote
	// chains, so every pool ends up with two remote pool entries.
	perChain := map[uint64]tokensapi.TokenExpansionInputPerChain{}
	for _, src := range []uint64{selA, selB, selC} {
		remoteChains := map[uint64]tokensapi.RemoteChainConfig[*datastore.AddressRef, datastore.AddressRef]{}
		for _, dst := range []uint64{selA, selB, selC} {
			if src != dst {
				remoteChains[dst] = tokensapi.RemoteChainConfig[*datastore.AddressRef, datastore.AddressRef]{
					OutboundRateLimiterConfig: &disabledOutbound,
				}
			}
		}
		deployer := deployers[src]
		perChain[src] = tokensapi.TokenExpansionInputPerChain{
			SkipOwnershipTransfer: true,
			TokenPoolVersion:      bnmOpsV2_0_0.Version,
			DeployTokenInput: &tokensapi.DeployTokenInput{
				Name: fmt.Sprintf("RRP_%d", src), Symbol: fmt.Sprintf("RRP_%d", src), Decimals: 18,
				ExternalAdmin: deployer.Hex(), CCIPAdmin: deployer.Hex(),
				Type: bnmERC20ops.ContractType,
			},
			DeployTokenPoolInput: &tokensapi.DeployTokenPoolInput{
				PoolType:              string(bnmOpsV2_0_0.ContractType),
				AllowedFinalityConfig: finality.Config{WaitForFinality: true},
			},
			TokenTransferConfig: &tokensapi.TokenTransferConfig{
				RemoteChains: remoteChains,
			},
		}
	}

	expansionOut, err := tokensapi.TokenExpansion().Apply(*e, tokensapi.TokenExpansionInput{
		ChainAdapterVersion:         cciputils.Version_2_0_0,
		MCMS:                        mcms.Input{},
		TokenExpansionInputPerChain: perChain,
	})
	require.NoError(t, err)
	MergeAddresses(t, e, expansionOut.DataStore)

	poolRef := func(sel uint64) common.Address {
		fltr := datastore.AddressRef{ChainSelector: sel, Type: datastore.ContractType(bnmOpsV2_0_0.ContractType), Version: bnmOpsV2_0_0.Version}
		addr, err := datastore_utils.FindAndFormatRef(e.DataStore, fltr, sel, evm_datastore_utils.ToEVMAddress)
		require.NoError(t, err)
		return addr
	}

	return removeRemotePoolsTestEnv{
		env: e, selA: selA, selB: selB, selC: selC,
		poolA: poolRef(selA), poolB: poolRef(selB), poolC: poolRef(selC),
	}
}

// TestRemoveRemotePools_PartialRemoval deploys a fully-connected three-chain pool and removes only
// a subset of one pool's remote entries, verifying that the targeted remote pool is removed while
// the untouched remote pool remains configured. It then verifies that removing a remote pool
// address that is not configured (a non-existent address) errors clearly.
func TestRemoveRemotePools_PartialRemoval(t *testing.T) {
	env := setupV2PoolsForRemoveRemotePools(t)

	poolA, err := tokenpoolV2_0_0.NewTokenPool(env.poolA, env.env.BlockChains.EVMChains()[env.selA].Client)
	require.NoError(t, err)

	// Sanity: pool A is connected to both B and C.
	remotePoolsB, err := poolA.GetRemotePools(&bind.CallOpts{Context: t.Context()}, env.selB)
	require.NoError(t, err)
	require.NotEmpty(t, remotePoolsB, "pool A should have a remote pool for chain B before removal")
	remotePoolsC, err := poolA.GetRemotePools(&bind.CallOpts{Context: t.Context()}, env.selC)
	require.NoError(t, err)
	require.NotEmpty(t, remotePoolsC, "pool A should have a remote pool for chain C before removal")

	// Remove only the B remote pool from pool A; the C remote pool must remain untouched.
	input := tokensapi.RemoveRemotePoolsInput{
		MCMS: mcms.Input{},
		Pools: []tokensapi.RemoveRemotePoolsPerPool{{
			ChainSelector: env.selA,
			Pool:          datastore.AddressRef{Address: env.poolA.Hex()},
			RemotePoolsToRemove: []tokensapi.RemotePoolToRemove{{
				Selector: env.selB,
				Remote:   datastore.AddressRef{Address: env.poolB.Hex()},
			}},
		}},
	}
	require.NoError(t, tokensapi.RemoveRemotePools().VerifyPreconditions(*env.env, input))
	_, err = tokensapi.RemoveRemotePools().Apply(*env.env, input)
	require.NoError(t, err)

	remotePoolsB, err = poolA.GetRemotePools(&bind.CallOpts{Context: t.Context()}, env.selB)
	require.NoError(t, err)
	require.Empty(t, remotePoolsB, "pool A should have no remote pool for chain B after removal")

	remotePoolsC, err = poolA.GetRemotePools(&bind.CallOpts{Context: t.Context()}, env.selC)
	require.NoError(t, err)
	require.NotEmpty(t, remotePoolsC, "pool A should still have a remote pool for chain C after partial removal")

	// Removing a remote pool address that is not configured must error clearly.
	env.env.OperationsBundle = evm_testsetup.BundleWithFreshReporter(env.env.OperationsBundle)
	nonexistent := tokensapi.RemoveRemotePoolsInput{
		MCMS: mcms.Input{},
		Pools: []tokensapi.RemoveRemotePoolsPerPool{{
			ChainSelector: env.selA,
			Pool:          datastore.AddressRef{Address: env.poolA.Hex()},
			RemotePoolsToRemove: []tokensapi.RemotePoolToRemove{{
				Selector: env.selC,
				Remote:   datastore.AddressRef{Address: "0x000000000000000000000000000000000000dEaD"},
			}},
		}},
	}
	require.NoError(t, tokensapi.RemoveRemotePools().VerifyPreconditions(*env.env, nonexistent))
	_, err = tokensapi.RemoveRemotePools().Apply(*env.env, nonexistent)
	require.ErrorContains(t, err, "is not configured")
}
