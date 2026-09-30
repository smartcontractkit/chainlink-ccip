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
	"github.com/smartcontractkit/chainlink-ccip/deployment/testhelpers"
	tokensapi "github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	cciputils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
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
			name:   "rejects_remote_ref_without_address",
			input:  singlePoolInput(tokensapi.RemotePoolToRemove{Selector: dst, Remote: datastore.AddressRef{Qualifier: "some-pool"}}),
			errors: []string{"must set remote.address"},
		},
		{
			name: "rejects_duplicate_remote_pools",
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
			errors: []string{"duplicate remote pool"},
		},
		{
			name: "rejects_duplicate_remote_pools_in_different_encodings",
			input: tokensapi.RemoveRemotePoolsInput{
				MCMS: mcms.Input{},
				Pools: []tokensapi.RemoveRemotePoolsPerPool{{
					ChainSelector: sel,
					Pool:          poolRef,
					RemotePoolsToRemove: []tokensapi.RemotePoolToRemove{
						{Selector: dst, Remote: datastore.AddressRef{Address: "0x000000000000000000000000000000000000dEaD"}},
						{Selector: dst, Remote: datastore.AddressRef{Address: "0x000000000000000000000000000000000000dead"}},
					},
				}},
			},
			errors: []string{"duplicate remote pool"},
		},
		{
			name: "rejects_skip_unsupported_peers_without_reverse_pass",
			input: tokensapi.RemoveRemotePoolsInput{
				MCMS: mcms.Input{},
				Pools: []tokensapi.RemoveRemotePoolsPerPool{{
					ChainSelector:        sel,
					Pool:                 poolRef,
					RemotePoolsToRemove:  []tokensapi.RemotePoolToRemove{{Selector: dst, Remote: remoteRef}},
					SkipUnsupportedPeers: true,
				}},
			},
			errors: []string{"skipUnsupportedPeers requires bidirectional or deactivate"},
		},
		{
			name:   "rejects_invalid_remote_pool_address",
			input:  singlePoolInput(tokensapi.RemotePoolToRemove{Selector: dst, Remote: datastore.AddressRef{Address: "not-an-address"}}),
			errors: []string{"invalid remote pool address"},
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
		{
			name: "rejects_all_remotes_with_remote_pools",
			input: tokensapi.RemoveRemotePoolsInput{
				MCMS: mcms.Input{},
				Pools: []tokensapi.RemoveRemotePoolsPerPool{{
					ChainSelector:       sel,
					Pool:                poolRef,
					AllRemotes:          true,
					RemotePoolsToRemove: []tokensapi.RemotePoolToRemove{{Selector: dst, Remote: remoteRef}},
				}},
			},
			errors: []string{"allRemotes and remotePoolsToRemove are mutually exclusive"},
		},
		{
			name: "rejects_bidirectional_without_remotes",
			input: tokensapi.RemoveRemotePoolsInput{
				MCMS: mcms.Input{},
				Pools: []tokensapi.RemoveRemotePoolsPerPool{{
					ChainSelector: sel,
					Pool:          poolRef,
					Bidirectional: true,
				}},
			},
			errors: []string{"bidirectional requires allRemotes or remotePoolsToRemove"},
		},
		{
			name: "rejects_deactivate_with_all_remotes",
			input: tokensapi.RemoveRemotePoolsInput{
				MCMS: mcms.Input{},
				Pools: []tokensapi.RemoveRemotePoolsPerPool{{
					ChainSelector: sel,
					Pool:          poolRef,
					Deactivate:    true,
					AllRemotes:    true,
				}},
			},
			errors: []string{"deactivate must be specified alone"},
		},
		{
			name: "rejects_deactivate_with_bidirectional",
			input: tokensapi.RemoveRemotePoolsInput{
				MCMS: mcms.Input{},
				Pools: []tokensapi.RemoveRemotePoolsPerPool{{
					ChainSelector: sel,
					Pool:          poolRef,
					Deactivate:    true,
					Bidirectional: true,
				}},
			},
			errors: []string{"deactivate must be specified alone"},
		},
		{
			name: "rejects_deactivate_with_remote_pools",
			input: tokensapi.RemoveRemotePoolsInput{
				MCMS: mcms.Input{},
				Pools: []tokensapi.RemoveRemotePoolsPerPool{{
					ChainSelector:       sel,
					Pool:                poolRef,
					Deactivate:          true,
					RemotePoolsToRemove: []tokensapi.RemotePoolToRemove{{Selector: dst, Remote: remoteRef}},
				}},
			},
			errors: []string{"deactivate must be specified alone"},
		},
		{
			name: "rejects_deprecated_local_chain",
			input: tokensapi.RemoveRemotePoolsInput{
				MCMS: mcms.Input{},
				Pools: []tokensapi.RemoveRemotePoolsPerPool{{
					ChainSelector:       chainsel.ETHEREUM_TESTNET_SEPOLIA_XLAYER_1.Selector,
					Pool:                poolRef,
					RemotePoolsToRemove: []tokensapi.RemotePoolToRemove{{Selector: dst, Remote: remoteRef}},
				}},
			},
			errors: []string{"deprecated/decommissioned"},
		},
		{
			name: "rejects_unknown_local_chain_selector",
			input: tokensapi.RemoveRemotePoolsInput{
				MCMS: mcms.Input{},
				Pools: []tokensapi.RemoveRemotePoolsPerPool{{
					ChainSelector:       1,
					Pool:                poolRef,
					RemotePoolsToRemove: []tokensapi.RemotePoolToRemove{{Selector: dst, Remote: remoteRef}},
				}},
			},
			errors: []string{"unknown chain selector"},
		},
		{
			name: "rejects_no_mode_selected",
			input: tokensapi.RemoveRemotePoolsInput{
				MCMS: mcms.Input{},
				Pools: []tokensapi.RemoveRemotePoolsPerPool{{
					ChainSelector: sel,
					Pool:          poolRef,
				}},
			},
			errors: []string{"no remote pools to remove"},
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
// state, then confirms that removing an already-removed pool is skipped with a warning.
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
	require.NoError(t, applyRemoveRemotePools(t, tc.env, input))

	remotePools, err = poolA.GetRemotePools(&bind.CallOpts{Context: t.Context()}, tc.selB)
	require.NoError(t, err)
	require.Empty(t, remotePools, "pool A should have no remote pool for chain B after removal")

	// Re-running the removal is idempotent: the pairing is already absent, so it is skipped
	// (with a warn log) rather than erroring.
	require.NoError(t, applyRemoveRemotePools(t, tc.env, input))
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
	require.NoError(t, applyRemoveRemotePools(t, pair.env, input))

	remotePools, err = oldPoolA.GetRemotePools(&bind.CallOpts{Context: t.Context()}, pair.selB)
	require.NoError(t, err)
	require.Empty(t, remotePools, "old pool A should have no remote pool for chain B after removal")

	// Re-running the removal is idempotent: the pairing is already absent, so it is skipped
	// (with a warn log) rather than erroring.
	require.NoError(t, applyRemoveRemotePools(t, pair.env, input))
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
// pool's remote entries) and the warn-and-skip path for addresses that are not configured.
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
// address that is not configured (a non-existent address) is skipped with a warning.
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
	require.NoError(t, applyRemoveRemotePools(t, env.env, input))

	remotePoolsB, err = poolA.GetRemotePools(&bind.CallOpts{Context: t.Context()}, env.selB)
	require.NoError(t, err)
	require.Empty(t, remotePoolsB, "pool A should have no remote pool for chain B after removal")

	remotePoolsC, err = poolA.GetRemotePools(&bind.CallOpts{Context: t.Context()}, env.selC)
	require.NoError(t, err)
	require.NotEmpty(t, remotePoolsC, "pool A should still have a remote pool for chain C after partial removal")

	// Removing a remote pool address that is not configured is an idempotent no-op (warn + skip).
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
	require.NoError(t, applyRemoveRemotePools(t, env.env, nonexistent))
}

func TestRemoveRemotePools_AllRemotes(t *testing.T) {
	env := setupV2PoolsForRemoveRemotePools(t)

	poolA, err := tokenpoolV2_0_0.NewTokenPool(env.poolA, env.env.BlockChains.EVMChains()[env.selA].Client)
	require.NoError(t, err)

	remotePoolsB, err := poolA.GetRemotePools(&bind.CallOpts{Context: t.Context()}, env.selB)
	require.NoError(t, err)
	require.NotEmpty(t, remotePoolsB, "pool A should have a remote pool for chain B before removal")
	remotePoolsC, err := poolA.GetRemotePools(&bind.CallOpts{Context: t.Context()}, env.selC)
	require.NoError(t, err)
	require.NotEmpty(t, remotePoolsC, "pool A should have a remote pool for chain C before removal")

	input := tokensapi.RemoveRemotePoolsInput{
		MCMS: mcms.Input{},
		Pools: []tokensapi.RemoveRemotePoolsPerPool{{
			ChainSelector: env.selA,
			Pool:          datastore.AddressRef{Address: env.poolA.Hex()},
			AllRemotes:    true,
		}},
	}
	require.NoError(t, applyRemoveRemotePools(t, env.env, input))

	remotePoolsB, err = poolA.GetRemotePools(&bind.CallOpts{Context: t.Context()}, env.selB)
	require.NoError(t, err)
	require.Empty(t, remotePoolsB, "pool A should have no remote pool for chain B after allRemotes")
	remotePoolsC, err = poolA.GetRemotePools(&bind.CallOpts{Context: t.Context()}, env.selC)
	require.NoError(t, err)
	require.Empty(t, remotePoolsC, "pool A should have no remote pool for chain C after allRemotes")
}

func TestRemoveRemotePools_Bidirectional(t *testing.T) {
	env := setupV2PoolsForRemoveRemotePools(t)

	poolA, err := tokenpoolV2_0_0.NewTokenPool(env.poolA, env.env.BlockChains.EVMChains()[env.selA].Client)
	require.NoError(t, err)
	poolB, err := tokenpoolV2_0_0.NewTokenPool(env.poolB, env.env.BlockChains.EVMChains()[env.selB].Client)
	require.NoError(t, err)

	remotePoolsB, err := poolA.GetRemotePools(&bind.CallOpts{Context: t.Context()}, env.selB)
	require.NoError(t, err)
	require.NotEmpty(t, remotePoolsB, "pool A should have a remote pool for chain B before removal")
	remotePoolsA, err := poolB.GetRemotePools(&bind.CallOpts{Context: t.Context()}, env.selA)
	require.NoError(t, err)
	require.NotEmpty(t, remotePoolsA, "pool B should have a remote pool for chain A before removal")

	input := tokensapi.RemoveRemotePoolsInput{
		MCMS: mcms.Input{},
		Pools: []tokensapi.RemoveRemotePoolsPerPool{{
			ChainSelector: env.selA,
			Pool:          datastore.AddressRef{Address: env.poolA.Hex()},
			Bidirectional: true,
			RemotePoolsToRemove: []tokensapi.RemotePoolToRemove{{
				Selector: env.selB,
				Remote:   datastore.AddressRef{Address: env.poolB.Hex()},
			}},
		}},
	}
	require.NoError(t, applyRemoveRemotePools(t, env.env, input))

	remotePoolsB, err = poolA.GetRemotePools(&bind.CallOpts{Context: t.Context()}, env.selB)
	require.NoError(t, err)
	require.Empty(t, remotePoolsB, "pool A should have no remote pool for chain B after bidirectional removal")

	remotePoolsA, err = poolB.GetRemotePools(&bind.CallOpts{Context: t.Context()}, env.selA)
	require.NoError(t, err)
	require.Empty(t, remotePoolsA, "pool B should have no remote pool for chain A after bidirectional removal")

	require.NoError(t, applyRemoveRemotePools(t, env.env, input))
}

func TestRemoveRemotePools_Deactivate(t *testing.T) {
	env := setupV2PoolsForRemoveRemotePools(t)

	poolA, err := tokenpoolV2_0_0.NewTokenPool(env.poolA, env.env.BlockChains.EVMChains()[env.selA].Client)
	require.NoError(t, err)
	poolB, err := tokenpoolV2_0_0.NewTokenPool(env.poolB, env.env.BlockChains.EVMChains()[env.selB].Client)
	require.NoError(t, err)

	remotePoolsB, err := poolA.GetRemotePools(&bind.CallOpts{Context: t.Context()}, env.selB)
	require.NoError(t, err)
	require.NotEmpty(t, remotePoolsB, "pool A should have a remote pool for chain B before removal")
	remotePoolsA, err := poolB.GetRemotePools(&bind.CallOpts{Context: t.Context()}, env.selA)
	require.NoError(t, err)
	require.NotEmpty(t, remotePoolsA, "pool B should have a remote pool for chain A before removal")

	input := tokensapi.RemoveRemotePoolsInput{
		MCMS: mcms.Input{},
		Pools: []tokensapi.RemoveRemotePoolsPerPool{{
			ChainSelector: env.selA,
			Pool:          datastore.AddressRef{Address: env.poolA.Hex()},
			Deactivate:    true,
		}},
	}
	require.NoError(t, applyRemoveRemotePools(t, env.env, input))

	remotePoolsB, err = poolA.GetRemotePools(&bind.CallOpts{Context: t.Context()}, env.selB)
	require.NoError(t, err)
	require.Empty(t, remotePoolsB, "pool A should have no remote pool for chain B after deactivate")
	remotePoolsC, err := poolA.GetRemotePools(&bind.CallOpts{Context: t.Context()}, env.selC)
	require.NoError(t, err)
	require.Empty(t, remotePoolsC, "pool A should have no remote pool for chain C after deactivate")

	remotePoolsA, err = poolB.GetRemotePools(&bind.CallOpts{Context: t.Context()}, env.selA)
	require.NoError(t, err)
	require.Empty(t, remotePoolsA, "pool B should have no remote pool for chain A after deactivate")

	tarReader, ok := tokensapi.GetTokenAdapterRegistry().GetTokenAdminRegistryReader(chainsel.FamilyEVM)
	require.True(t, ok, "EVM TAR reader should be registered")
	tokenRef, err := datastore_utils.FindAndFormatRef(env.env.DataStore, datastore.AddressRef{ChainSelector: env.selA, Type: datastore.ContractType(bnmERC20ops.ContractType)}, env.selA, datastore_utils.FullRef)
	require.NoError(t, err)
	activePool, err := tarReader.GetActivePool(*env.env, env.selA, tokenRef)
	require.NoError(t, err)
	require.Empty(t, activePool, "pool A should be unregistered from the TAR after deactivate")
}

func TestRemoveRemotePools_DeactivateSkipsUnregisterWhenActivePoolEmpty(t *testing.T) {
	env := setupV2PoolsForRemoveRemotePools(t)

	input := tokensapi.RemoveRemotePoolsInput{
		MCMS: mcms.Input{},
		Pools: []tokensapi.RemoveRemotePoolsPerPool{{
			ChainSelector: env.selA,
			Pool:          datastore.AddressRef{Address: env.poolA.Hex()},
			Deactivate:    true,
		}},
	}

	require.NoError(t, applyRemoveRemotePools(t, env.env, input))

	tarReader, ok := tokensapi.GetTokenAdapterRegistry().GetTokenAdminRegistryReader(chainsel.FamilyEVM)
	require.True(t, ok, "EVM TAR reader should be registered")
	fullTokenRef, err := datastore_utils.FindAndFormatRef(env.env.DataStore, datastore.AddressRef{ChainSelector: env.selA, Type: datastore.ContractType(bnmERC20ops.ContractType)}, env.selA, datastore_utils.FullRef)
	require.NoError(t, err)
	activePool, err := tarReader.GetActivePool(*env.env, env.selA, fullTokenRef)
	require.NoError(t, err)
	require.Empty(t, activePool, "pool A should be unregistered from the TAR after deactivate")

	require.NoError(t, applyRemoveRemotePools(t, env.env, input))

	activePool, err = tarReader.GetActivePool(*env.env, env.selA, fullTokenRef)
	require.NoError(t, err)
	require.Empty(t, activePool, "the registry entry should remain empty after a repeated deactivate")
}

// TestRemoveRemotePools_DeactivateRetiredPoolAfterPeerUpgrade deactivates a retired pool (A1) in a
// web whose peers were upgraded too (A1->A2, then B1->B2 and C1->C2). Each peer chain then has two
// pools listing A1: the retired peer pool A1 is paired with (B1, C1), and the peer's active pool,
// which copied A1 during the peer's own upgrade (B2, C2). Deactivating A1 must remove it from all
// four, while leaving A2's pairings and the TAR entry (which points to A2) untouched.
func TestRemoveRemotePools_DeactivateRetiredPoolAfterPeerUpgrade(t *testing.T) {
	web := setupIncrementalMigrationWeb(t)
	env := web.e

	selA, selB, selC := web.tokenA.ChainSelector, web.tokenB.ChainSelector, web.tokenC.ChainSelector
	addr := func(ref datastore.AddressRef) common.Address { return common.HexToAddress(ref.Address) }
	poolA1, poolB1, poolC1 := addr(web.v1PoolRefs[selA]), addr(web.v1PoolRefs[selB]), addr(web.v1PoolRefs[selC])
	poolA2, poolB2, poolC2 := addr(web.v2PoolRefs[0]), addr(web.v2PoolRefs[1]), addr(web.v2PoolRefs[2])

	remotePools := func(pool common.Address, sel, remoteSel uint64) []common.Address {
		return BytesToAddressesEVM(ReadRemotePools(t, env, sel, datastore.AddressRef{Address: pool.Hex()}, datastore.AddressRef{}, remoteSel))
	}

	peerPools := []struct {
		name string
		pool common.Address
		sel  uint64
	}{
		{"B1", poolB1, selB},
		{"B2", poolB2, selB},
		{"C1", poolC1, selC},
		{"C2", poolC2, selC},
	}

	// Pre-state: A1 is listed on both pools of each peer chain.
	for _, p := range peerPools {
		require.Contains(t, remotePools(p.pool, p.sel, selA), poolA1, "%s should list A1 before deactivate", p.name)
	}

	input := tokensapi.RemoveRemotePoolsInput{
		MCMS: NewDefaultInputForMCMS("deactivate A1"),
		Pools: []tokensapi.RemoveRemotePoolsPerPool{{
			ChainSelector: selA,
			Pool:          datastore.AddressRef{Address: poolA1.Hex()},
			Deactivate:    true,
		}},
	}
	require.NoError(t, applyRemoveRemotePools(t, env, input))

	// Forward pass: A1 no longer lists any pool on B or C.
	require.Empty(t, remotePools(poolA1, selA, selB), "A1 should have no remote pool for chain B after deactivate")
	require.Empty(t, remotePools(poolA1, selA, selC), "A1 should have no remote pool for chain C after deactivate")

	// Reverse pass: A1 is gone from both the retired and the active pool of each peer, and A2 is
	// still listed everywhere.
	for _, p := range peerPools {
		remotes := remotePools(p.pool, p.sel, selA)
		require.NotContains(t, remotes, poolA1, "%s should no longer list A1 after deactivate", p.name)
		require.Contains(t, remotes, poolA2, "%s should still list A2 after deactivate", p.name)
	}

	// A2's own pairings are untouched.
	require.Contains(t, remotePools(poolA2, selA, selB), poolB2, "A2 should still list B2")
	require.Contains(t, remotePools(poolA2, selA, selC), poolC2, "A2 should still list C2")

	// The TAR still points to A2 (the unregister is skipped because A1 is not the active pool).
	tarReader, ok := tokensapi.GetTokenAdapterRegistry().GetTokenAdminRegistryReader(chainsel.FamilyEVM)
	require.True(t, ok, "EVM TAR reader should be registered")
	activePool, err := tarReader.GetActivePool(*env, selA, web.tokenA)
	require.NoError(t, err)
	require.Equal(t, poolA2, common.BytesToAddress(activePool), "TAR should still point to A2")
}

// TestRemoveRemotePools_DeactivateResumesAfterPartialReverse simulates a run whose reverse pass
// reached peer B but not peer C before failing (so the forward pass never ran). Because the reverse
// pass runs before the forward pass, A's own remote list still records the remaining work, and a
// re-run of deactivate(A) skips B (already clean), cleans C, then runs the forward pass and the TAR
// unregister.
func TestRemoveRemotePools_DeactivateResumesAfterPartialReverse(t *testing.T) {
	env := setupV2PoolsForRemoveRemotePools(t)
	e := env.env
	remotePools := func(pool common.Address, sel, remoteSel uint64) []common.Address {
		return BytesToAddressesEVM(ReadRemotePools(t, e, sel, datastore.AddressRef{Address: pool.Hex()}, datastore.AddressRef{}, remoteSel))
	}

	// Partial first run: only B has dropped A.
	require.NoError(t, applyRemoveRemotePools(t, e, tokensapi.RemoveRemotePoolsInput{
		Pools: []tokensapi.RemoveRemotePoolsPerPool{{
			ChainSelector:       env.selB,
			Pool:                datastore.AddressRef{Address: env.poolB.Hex()},
			RemotePoolsToRemove: []tokensapi.RemotePoolToRemove{{Selector: env.selA, Remote: datastore.AddressRef{Address: env.poolA.Hex()}}},
		}},
	}))
	require.NotContains(t, remotePools(env.poolB, env.selB, env.selA), env.poolA, "B should no longer list A after the partial run")
	require.Contains(t, remotePools(env.poolC, env.selC, env.selA), env.poolA, "C should still list A after the partial run")
	require.Contains(t, remotePools(env.poolA, env.selA, env.selB), env.poolB, "A should still list B (forward pass never ran)")

	// Re-run deactivate(A): it must finish the teardown.
	require.NoError(t, applyRemoveRemotePools(t, e, tokensapi.RemoveRemotePoolsInput{
		Pools: []tokensapi.RemoveRemotePoolsPerPool{{
			ChainSelector: env.selA,
			Pool:          datastore.AddressRef{Address: env.poolA.Hex()},
			Deactivate:    true,
		}},
	}))

	require.Empty(t, remotePools(env.poolA, env.selA, env.selB), "A should have no remote pool for chain B")
	require.Empty(t, remotePools(env.poolA, env.selA, env.selC), "A should have no remote pool for chain C")
	require.NotContains(t, remotePools(env.poolB, env.selB, env.selA), env.poolA, "B should not list A")
	require.NotContains(t, remotePools(env.poolC, env.selC, env.selA), env.poolA, "C should not list A")

	tarReader, ok := tokensapi.GetTokenAdapterRegistry().GetTokenAdminRegistryReader(chainsel.FamilyEVM)
	require.True(t, ok, "EVM TAR reader should be registered")
	activePool, err := tarReader.GetActivePool(*e, env.selA, FindFullRef(t, e, env.selA, datastore.AddressRef{Type: datastore.ContractType(bnmERC20ops.ContractType)}))
	require.NoError(t, err)
	require.Empty(t, activePool, "A should be unregistered from the TAR")
}

// TestRemoveRemotePools_ExplicitModeRecoversDroppedReverse simulates the mixed-ownership caveat
// documented on RemoveRemotePoolsPerPool: A's forward removals landed but the peers' removals did
// not (e.g. their MCMS proposal was dropped). A's own remote list is now empty, so a
// discovery-based re-run can no longer find the peers; the documented recovery is bidirectional with
// an explicit remotePoolsToRemove, which does not depend on A's remote list.
func TestRemoveRemotePools_ExplicitModeRecoversDroppedReverse(t *testing.T) {
	harness := setupV2PoolsForRemoveRemotePools(t)
	env := harness.env

	remotePools := func(pool common.Address, sel, remoteSel uint64) []common.Address {
		return BytesToAddressesEVM(ReadRemotePools(t, env, sel, datastore.AddressRef{Address: pool.Hex()}, datastore.AddressRef{}, remoteSel))
	}

	// Forward side only: A drops B and C, the peers still list A.
	require.NoError(t, applyRemoveRemotePools(t, env, tokensapi.RemoveRemotePoolsInput{
		Pools: []tokensapi.RemoveRemotePoolsPerPool{{
			ChainSelector: harness.selA,
			Pool:          datastore.AddressRef{Address: harness.poolA.Hex()},
			AllRemotes:    true,
		}},
	}))
	require.Empty(t, remotePools(harness.poolA, harness.selA, harness.selB), "A should no longer list B")
	require.Empty(t, remotePools(harness.poolA, harness.selA, harness.selC), "A should no longer list C")
	require.Contains(t, remotePools(harness.poolB, harness.selB, harness.selA), harness.poolA, "B should still list A")
	require.Contains(t, remotePools(harness.poolC, harness.selC, harness.selA), harness.poolA, "C should still list A")

	// Recovery: bidirectional with the peers listed explicitly.
	require.NoError(t, applyRemoveRemotePools(t, env, tokensapi.RemoveRemotePoolsInput{
		Pools: []tokensapi.RemoveRemotePoolsPerPool{{
			ChainSelector: harness.selA,
			Pool:          datastore.AddressRef{Address: harness.poolA.Hex()},
			Bidirectional: true,
			RemotePoolsToRemove: []tokensapi.RemotePoolToRemove{
				{Selector: harness.selB, Remote: datastore.AddressRef{Address: harness.poolB.Hex()}},
				{Selector: harness.selC, Remote: datastore.AddressRef{Address: harness.poolC.Hex()}},
			},
		}},
	}))
	require.NotContains(t, remotePools(harness.poolB, harness.selB, harness.selA), harness.poolA, "B should no longer list A after recovery")
	require.NotContains(t, remotePools(harness.poolC, harness.selC, harness.selA), harness.poolA, "C should no longer list A after recovery")
}

// TestRemoveRemotePools_BidirectionalCleansNamedPoolWhenPeerHasNoActivePool covers a peer whose
// token has no active pool in the TAR. The peer's TAR-active pool is dropped as a reverse target,
// but the pool named by the remote entry is still valid and is cleaned, so the bidirectional
// removal completes: A drops B and B drops A.
func TestRemoveRemotePools_BidirectionalCleansNamedPoolWhenPeerHasNoActivePool(t *testing.T) {
	harness := setupV2PoolsForRemoveRemotePools(t)
	env := harness.env

	remotePools := func(pool common.Address, sel, remoteSel uint64) []common.Address {
		return BytesToAddressesEVM(ReadRemotePools(t, env, sel, datastore.AddressRef{Address: pool.Hex()}, datastore.AddressRef{}, remoteSel))
	}

	// Unregister B's token from its TAR directly, leaving all pool pairings in place.
	tarManager, ok := tokensapi.GetTokenAdapterRegistry().GetTokenAdminRegistryManager(chainsel.FamilyEVM)
	require.True(t, ok, "EVM TAR manager should be registered")
	tokenRefB := FindFullRef(t, env, harness.selB, datastore.AddressRef{Type: datastore.ContractType(bnmERC20ops.ContractType)})
	env.OperationsBundle = evm_testsetup.BundleWithFreshReporter(env.OperationsBundle)
	_, err := cldf_ops.ExecuteSequence(env.OperationsBundle, tarManager.UnregisterToken(), env.BlockChains, tokensapi.UnregisterTokenSequenceInput{
		Selector:          harness.selB,
		TokenRef:          tokenRefB,
		ExistingDataStore: env.DataStore,
	})
	require.NoError(t, err)
	activePoolB, err := tarManager.GetActivePool(*env, harness.selB, tokenRefB)
	require.NoError(t, err)
	require.Empty(t, activePoolB, "B's token should have no active pool")

	require.NoError(t, applyRemoveRemotePools(t, env, tokensapi.RemoveRemotePoolsInput{
		Pools: []tokensapi.RemoveRemotePoolsPerPool{{
			ChainSelector:       harness.selA,
			Pool:                datastore.AddressRef{Address: harness.poolA.Hex()},
			Bidirectional:       true,
			RemotePoolsToRemove: []tokensapi.RemotePoolToRemove{{Selector: harness.selB, Remote: datastore.AddressRef{Address: harness.poolB.Hex()}}},
		}},
	}))

	// B's missing active pool does not block the teardown: A drops B (forward) and B drops A
	// (reverse, matched through the pool the remote entry names).
	require.NotContains(t, remotePools(harness.poolA, harness.selA, harness.selB), harness.poolB, "A should no longer list B")
	require.NotContains(t, remotePools(harness.poolB, harness.selB, harness.selA), harness.poolA, "B should no longer list A")
	// The lanes to C are untouched.
	require.Contains(t, remotePools(harness.poolA, harness.selA, harness.selC), harness.poolC, "A should still list C")
	require.Contains(t, remotePools(harness.poolB, harness.selB, harness.selC), harness.poolC, "B should still list C")
}

// TestRemoveRemotePools_DeactivateSkipsZeroAddressEntry covers a pool that lists the zero address
// as a remote pool (pools only reject empty remote pool bytes, so older tooling could write 32 zero
// bytes). The zero address names no peer pool, so the reverse pass skips it as a target while the
// forward pass still removes it, and deactivate completes.
func TestRemoveRemotePools_DeactivateSkipsZeroAddressEntry(t *testing.T) {
	harness := setupV2PoolsForRemoveRemotePools(t)
	env := harness.env
	remotePools := func(pool common.Address, sel, remoteSel uint64) []common.Address {
		return BytesToAddressesEVM(ReadRemotePools(t, env, sel, datastore.AddressRef{Address: pool.Hex()}, datastore.AddressRef{}, remoteSel))
	}

	addRawRemotePool(t, env, harness.selA, harness.poolA, harness.selC, make([]byte, 32))
	require.Contains(t, remotePools(harness.poolA, harness.selA, harness.selC), common.Address{}, "A should list the zero address for C")

	require.NoError(t, applyRemoveRemotePools(t, env, tokensapi.RemoveRemotePoolsInput{
		Pools: []tokensapi.RemoveRemotePoolsPerPool{{
			ChainSelector: harness.selA,
			Pool:          datastore.AddressRef{Address: harness.poolA.Hex()},
			Deactivate:    true,
		}},
	}))

	require.Empty(t, remotePools(harness.poolA, harness.selA, harness.selB), "A should have no remote pool for chain B")
	require.Empty(t, remotePools(harness.poolA, harness.selA, harness.selC), "A should have no remote pool for chain C, including the zero address")
	require.NotContains(t, remotePools(harness.poolB, harness.selB, harness.selA), harness.poolA, "B should not list A")
	require.NotContains(t, remotePools(harness.poolC, harness.selC, harness.selA), harness.poolA, "C should not list A")
}

// TestRemoveRemotePools_DeactivateUnresolvableEntryWritesNothing covers a pool that lists a remote
// pool the tooling cannot resolve (here an address with no pool behind it). Deactivate must fail
// before writing to any chain, even for peers ordered before the bad entry, and must point the
// operator at a lane-only removal. After that removal, a re-run completes the teardown.
func TestRemoveRemotePools_DeactivateUnresolvableEntryWritesNothing(t *testing.T) {
	harness := setupV2PoolsForRemoveRemotePools(t)
	env := harness.env
	remotePools := func(pool common.Address, sel, remoteSel uint64) []common.Address {
		return BytesToAddressesEVM(ReadRemotePools(t, env, sel, datastore.AddressRef{Address: pool.Hex()}, datastore.AddressRef{}, remoteSel))
	}

	// A lists C's real pool first and the bad entry second, so the bad entry is resolved only after
	// the reverse removals from B and C are already planned.
	badPool := common.HexToAddress("0x000000000000000000000000000000000000dEaD")
	addRawRemotePool(t, env, harness.selA, harness.poolA, harness.selC, common.LeftPadBytes(badPool.Bytes(), 32))
	require.Equal(t, []common.Address{harness.poolC, badPool}, remotePools(harness.poolA, harness.selA, harness.selC))

	deactivateA := tokensapi.RemoveRemotePoolsInput{
		Pools: []tokensapi.RemoveRemotePoolsPerPool{{
			ChainSelector: harness.selA,
			Pool:          datastore.AddressRef{Address: harness.poolA.Hex()},
			Deactivate:    true,
		}},
	}
	err := applyRemoveRemotePools(t, env, deactivateA)
	require.ErrorContains(t, err, "remove it from this pool with a lane-only removal")

	// Nothing was written on any chain.
	require.Equal(t, []common.Address{harness.poolB}, remotePools(harness.poolA, harness.selA, harness.selB), "A should still list B")
	require.Equal(t, []common.Address{harness.poolC, badPool}, remotePools(harness.poolA, harness.selA, harness.selC), "A should still list C and the bad entry")
	require.Contains(t, remotePools(harness.poolB, harness.selB, harness.selA), harness.poolA, "B should still list A")
	require.Contains(t, remotePools(harness.poolC, harness.selC, harness.selA), harness.poolA, "C should still list A")
	tarReader, ok := tokensapi.GetTokenAdapterRegistry().GetTokenAdminRegistryReader(chainsel.FamilyEVM)
	require.True(t, ok, "EVM TAR reader should be registered")
	tokenRefA := FindFullRef(t, env, harness.selA, datastore.AddressRef{Type: datastore.ContractType(bnmERC20ops.ContractType)})
	activePool, err := tarReader.GetActivePool(*env, harness.selA, tokenRefA)
	require.NoError(t, err)
	require.Equal(t, harness.poolA.Bytes(), activePool, "A should still be registered in the TAR")

	// Recovery: remove the bad entry with a lane-only removal, then re-run deactivate.
	require.NoError(t, applyRemoveRemotePools(t, env, tokensapi.RemoveRemotePoolsInput{
		Pools: []tokensapi.RemoveRemotePoolsPerPool{{
			ChainSelector:       harness.selA,
			Pool:                datastore.AddressRef{Address: harness.poolA.Hex()},
			RemotePoolsToRemove: []tokensapi.RemotePoolToRemove{{Selector: harness.selC, Remote: datastore.AddressRef{Address: badPool.Hex()}}},
		}},
	}))
	require.NoError(t, applyRemoveRemotePools(t, env, deactivateA))

	require.Empty(t, remotePools(harness.poolA, harness.selA, harness.selB), "A should have no remote pool for chain B")
	require.Empty(t, remotePools(harness.poolA, harness.selA, harness.selC), "A should have no remote pool for chain C")
	require.NotContains(t, remotePools(harness.poolB, harness.selB, harness.selA), harness.poolA, "B should not list A")
	require.NotContains(t, remotePools(harness.poolC, harness.selC, harness.selA), harness.poolA, "C should not list A")
	activePool, err = tarReader.GetActivePool(*env, harness.selA, tokenRefA)
	require.NoError(t, err)
	require.Empty(t, activePool, "A should be unregistered from the TAR")
}

// TestRemoveRemotePools_SeveralRemotePoolsForOneChain removes two remote pools that A lists for
// the same remote chain with one explicit entry each, in a single run.
func TestRemoveRemotePools_SeveralRemotePoolsForOneChain(t *testing.T) {
	harness := setupV2PoolsForRemoveRemotePools(t)
	env := harness.env
	remotePools := func(pool common.Address, sel, remoteSel uint64) []common.Address {
		return BytesToAddressesEVM(ReadRemotePools(t, env, sel, datastore.AddressRef{Address: pool.Hex()}, datastore.AddressRef{}, remoteSel))
	}

	extraPool := common.HexToAddress("0x000000000000000000000000000000000000dEaD")
	addRawRemotePool(t, env, harness.selA, harness.poolA, harness.selB, common.LeftPadBytes(extraPool.Bytes(), 32))
	require.Equal(t, []common.Address{harness.poolB, extraPool}, remotePools(harness.poolA, harness.selA, harness.selB))

	require.NoError(t, applyRemoveRemotePools(t, env, tokensapi.RemoveRemotePoolsInput{
		Pools: []tokensapi.RemoveRemotePoolsPerPool{{
			ChainSelector: harness.selA,
			Pool:          datastore.AddressRef{Address: harness.poolA.Hex()},
			RemotePoolsToRemove: []tokensapi.RemotePoolToRemove{
				{Selector: harness.selB, Remote: datastore.AddressRef{Address: harness.poolB.Hex()}},
				{Selector: harness.selB, Remote: datastore.AddressRef{Address: extraPool.Hex()}},
			},
		}},
	}))

	require.Empty(t, remotePools(harness.poolA, harness.selA, harness.selB), "A should list no pool for chain B")
	require.Equal(t, []common.Address{harness.poolC}, remotePools(harness.poolA, harness.selA, harness.selC), "A should still list C")
}

// TestRemoveRemotePools_ExplicitRemovesBothEncodings covers a pool that lists the same EVM remote
// pool twice: left-padded to 32 bytes and raw 20 bytes (legacy pools may hold either or both). An
// explicit removal of that address removes both entries.
func TestRemoveRemotePools_ExplicitRemovesBothEncodings(t *testing.T) {
	harness := setupV2PoolsForRemoveRemotePools(t)
	env := harness.env

	addRawRemotePool(t, env, harness.selA, harness.poolA, harness.selB, harness.poolB.Bytes())
	require.ElementsMatch(t,
		[][]byte{common.LeftPadBytes(harness.poolB.Bytes(), 32), harness.poolB.Bytes()},
		ReadRemotePools(t, env, harness.selA, datastore.AddressRef{Address: harness.poolA.Hex()}, datastore.AddressRef{}, harness.selB),
		"A should list B in both encodings")

	require.NoError(t, applyRemoveRemotePools(t, env, tokensapi.RemoveRemotePoolsInput{
		Pools: []tokensapi.RemoveRemotePoolsPerPool{{
			ChainSelector:       harness.selA,
			Pool:                datastore.AddressRef{Address: harness.poolA.Hex()},
			RemotePoolsToRemove: []tokensapi.RemotePoolToRemove{{Selector: harness.selB, Remote: datastore.AddressRef{Address: harness.poolB.Hex()}}},
		}},
	}))

	require.Empty(t, ReadRemotePools(t, env, harness.selA, datastore.AddressRef{Address: harness.poolA.Hex()}, datastore.AddressRef{}, harness.selB), "A should list B in neither encoding")
}

// TestRemoveRemotePools_DeactivateRemovesBothEncodings covers raw 20-byte entries alongside the
// padded ones on both the local pool (forward pass) and a peer (reverse pass): deactivate removes
// every encoding of each pairing.
func TestRemoveRemotePools_DeactivateRemovesBothEncodings(t *testing.T) {
	harness := setupV2PoolsForRemoveRemotePools(t)
	env := harness.env
	rawRemotePools := func(pool common.Address, sel, remoteSel uint64) [][]byte {
		return ReadRemotePools(t, env, sel, datastore.AddressRef{Address: pool.Hex()}, datastore.AddressRef{}, remoteSel)
	}

	addRawRemotePool(t, env, harness.selA, harness.poolA, harness.selB, harness.poolB.Bytes())
	addRawRemotePool(t, env, harness.selA, harness.poolA, harness.selC, harness.poolC.Bytes())
	addRawRemotePool(t, env, harness.selB, harness.poolB, harness.selA, harness.poolA.Bytes())
	require.Len(t, rawRemotePools(harness.poolB, harness.selB, harness.selA), 2, "B should list A in both encodings")

	require.NoError(t, applyRemoveRemotePools(t, env, tokensapi.RemoveRemotePoolsInput{
		Pools: []tokensapi.RemoveRemotePoolsPerPool{{
			ChainSelector: harness.selA,
			Pool:          datastore.AddressRef{Address: harness.poolA.Hex()},
			Deactivate:    true,
		}},
	}))

	require.Empty(t, rawRemotePools(harness.poolA, harness.selA, harness.selB), "A should list nothing for chain B")
	require.Empty(t, rawRemotePools(harness.poolA, harness.selA, harness.selC), "A should list nothing for chain C")
	require.Empty(t, rawRemotePools(harness.poolB, harness.selB, harness.selA), "B should list A in neither encoding")
	require.NotContains(t, BytesToAddressesEVM(rawRemotePools(harness.poolC, harness.selC, harness.selA)), harness.poolA, "C should not list A")
}

// TestRemoveRemotePools_DeactivateWithUnloadedPeer covers a peer chain the environment does not
// load. Deactivate fails before writing anything unless skipUnsupportedPeers is set; with it, the
// unloaded peer is skipped (and keeps listing A) while the rest of the teardown completes.
func TestRemoveRemotePools_DeactivateWithUnloadedPeer(t *testing.T) {
	harness := setupV2PoolsForRemoveRemotePools(t)
	fullEnv := harness.env
	remotePools := func(pool common.Address, sel, remoteSel uint64) []common.Address {
		return BytesToAddressesEVM(ReadRemotePools(t, fullEnv, sel, datastore.AddressRef{Address: pool.Hex()}, datastore.AddressRef{}, remoteSel))
	}

	// An environment that does not load chain C.
	loaded := map[uint64]cldf_chain.BlockChain{}
	for sel, chain := range fullEnv.BlockChains.All() {
		if sel != harness.selC {
			loaded[sel] = chain
		}
	}
	envWithoutC := *fullEnv
	envWithoutC.BlockChains = cldf_chain.NewBlockChains(loaded)

	deactivateA := func(skip bool) tokensapi.RemoveRemotePoolsInput {
		return tokensapi.RemoveRemotePoolsInput{
			Pools: []tokensapi.RemoveRemotePoolsPerPool{{
				ChainSelector:        harness.selA,
				Pool:                 datastore.AddressRef{Address: harness.poolA.Hex()},
				Deactivate:           true,
				SkipUnsupportedPeers: skip,
			}},
		}
	}

	err := applyRemoveRemotePools(t, &envWithoutC, deactivateA(false))
	require.ErrorContains(t, err, "chain is not loaded in the environment")
	require.ErrorContains(t, err, "skipUnsupportedPeers")
	require.Contains(t, remotePools(harness.poolB, harness.selB, harness.selA), harness.poolA, "B should still list A (nothing was written)")
	require.Equal(t, []common.Address{harness.poolC}, remotePools(harness.poolA, harness.selA, harness.selC), "A should still list C (nothing was written)")

	require.NoError(t, applyRemoveRemotePools(t, &envWithoutC, deactivateA(true)))
	require.Empty(t, remotePools(harness.poolA, harness.selA, harness.selB), "A should list nothing for chain B")
	require.Empty(t, remotePools(harness.poolA, harness.selA, harness.selC), "A should list nothing for chain C (the forward pass only needs chain A)")
	require.NotContains(t, remotePools(harness.poolB, harness.selB, harness.selA), harness.poolA, "B should not list A")
	require.Contains(t, remotePools(harness.poolC, harness.selC, harness.selA), harness.poolA, "C was skipped, so it should still list A")
	tarReader, ok := tokensapi.GetTokenAdapterRegistry().GetTokenAdminRegistryReader(chainsel.FamilyEVM)
	require.True(t, ok, "EVM TAR reader should be registered")
	activePool, err := tarReader.GetActivePool(*fullEnv, harness.selA, FindFullRef(t, fullEnv, harness.selA, datastore.AddressRef{Type: datastore.ContractType(bnmERC20ops.ContractType)}))
	require.NoError(t, err)
	require.Empty(t, activePool, "A should be unregistered from the TAR")
}

// addRawRemotePool writes remotePool (raw bytes, unvalidated) to the v2.0.0 pool's remote pool list
// for remoteSel, using the deployer key.
func addRawRemotePool(t *testing.T, e *cldf_deployment.Environment, sel uint64, pool common.Address, remoteSel uint64, remotePool []byte) {
	t.Helper()
	chain := e.BlockChains.EVMChains()[sel]
	tp, err := tokenpoolV2_0_0.NewTokenPool(pool, chain.Client)
	require.NoError(t, err)
	tx, err := tp.AddRemotePool(chain.DeployerKey, remoteSel, remotePool)
	require.NoError(t, err)
	_, err = chain.Confirm(tx)
	require.NoError(t, err)
}

// applyRemoveRemotePools verifies and applies input with a fresh operations reporter (so repeated
// applies re-read on-chain state instead of returning cached results), then executes any MCMS
// timelock proposals the apply produced. It returns the apply error; verification must pass.
func applyRemoveRemotePools(t *testing.T, e *cldf_deployment.Environment, input tokensapi.RemoveRemotePoolsInput) error {
	t.Helper()
	e.OperationsBundle = evm_testsetup.BundleWithFreshReporter(e.OperationsBundle)
	require.NoError(t, tokensapi.RemoveRemotePools().VerifyPreconditions(*e, input))
	out, err := tokensapi.RemoveRemotePools().Apply(*e, input)
	if err != nil {
		return err
	}
	testhelpers.ProcessTimelockProposals(t, *e, out.MCMSTimelockProposals, false)
	return nil
}
