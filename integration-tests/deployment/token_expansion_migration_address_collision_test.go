package deployment

import (
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"

	chainsel "github.com/smartcontractkit/chain-selectors"
	bnmERC20ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/burn_mint_erc20"
	testsetupV2_0_0 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/testsetup"
	tarbindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_5_0/token_admin_registry"
	tokenpoolV2_0_0 "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v2_0_0/token_pool"
	"github.com/smartcontractkit/chainlink-ccip/deployment/testhelpers"
	tokensapi "github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	cciputils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	ccvadapters "github.com/smartcontractkit/chainlink-ccip/deployment/v2_0_0/adapters"
	v2changesets "github.com/smartcontractkit/chainlink-ccip/deployment/v2_0_0/changesets"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/onchain"
	mcms_types "github.com/smartcontractkit/mcms/types"
)

// TestTokenExpansionMigration_AddressCollisionReversePropagation is a regression test for a
// reverse-propagation edge case: a migrating hub A with two counterpart chains B and C whose
// pools live at the SAME address (forced here by giving all simulated chains one shared
// AdminAccount so CREATE addresses coincide).
//
// Root cause: the v1.5.0/v1.5.1 counterpart configure sequences did not carry the chain selector in
// their input, so B's and C's inputs were byte-identical and the operations report cache (keyed on
// input only) served the first chain's cached report for the second - emitting two ops against one
// chain and none for the other. With the chain selector in the input this must be exactly one
// reverse-propagation op for B and one for C.
func TestTokenExpansionMigration_AddressCollisionReversePropagation(t *testing.T) {
	selA := chainsel.TEST_90000001.Selector
	selB := chainsel.TEST_90000002.Selector
	selC := chainsel.TEST_90000003.Selector

	// One shared admin account across all simulated chains makes B and C deploy the same contracts
	// at the same CREATE addresses, which is the precondition for the bug.
	adminKey, err := crypto.GenerateKey()
	require.NoError(t, err)
	e, err := environment.New(t.Context(), environment.WithEVMSimulatedWithConfig(
		t, []uint64{selA, selB, selC}, onchain.EVMSimLoaderConfig{AdminAccount: adminKey},
	))
	require.NoError(t, err)

	SeedUltraFastCurseMCMS(t, e)
	cumulative := datastore.NewMemoryDataStore()
	DeployChainContractsV2_0_0(t, e, cumulative, selA)
	DeployChainContractsV2_0_0(t, e, cumulative, selB)
	DeployChainContractsV2_0_0(t, e, cumulative, selC)
	e.DataStore = cumulative.Seal()

	// Shift chain A's deployer nonce so A's later token pool lands at a DIFFERENT address than B/C's.
	// This matches prod: hub A has a distinct pool, and only the two spokes B and C collide (B == C).
	DeployBurnMintTokenEVM(t, e, selA, e.BlockChains.EVMChains()[selA].DeployerKey.From.Hex())

	DeployMCMS(t, e, selA, []string{cciputils.CLLQualifier})
	DeployMCMS(t, e, selB, []string{cciputils.CLLQualifier})
	DeployMCMS(t, e, selC, []string{cciputils.CLLQualifier})

	deployerA := e.BlockChains.EVMChains()[selA].DeployerKey.From.Hex()
	e.OperationsBundle = testsetupV2_0_0.BundleWithFreshReporter(e.OperationsBundle)
	connectOut, err := v2changesets.ConfigureChainsForLanesFromTopology(
		ccvadapters.GetCommitteeVerifierContractRegistry(),
		ccvadapters.GetChainFamilyRegistry(),
		changesets.GetRegistry(),
	).Apply(*e, v2changesets.ConfigureChainsForLanesFromTopologyConfig{
		Topology: NewLaneTopologyForV2(deployerA, selA, selB, selC),
		BuildLanesCrossFamilyConfig: v2changesets.BuildLanesCrossFamilyConfig{
			MCMS: mcms.Input{},
			Lanes: []v2changesets.CrossFamilyLanePair{
				{ChainA: selA, ChainB: selB, ChainAOverrides: NewLaneOverridesForV2(selA), ChainBOverrides: NewLaneOverridesForV2(selB)},
				{ChainA: selA, ChainB: selC, ChainAOverrides: NewLaneOverridesForV2(selA), ChainBOverrides: NewLaneOverridesForV2(selC)},
				{ChainA: selB, ChainB: selC, ChainAOverrides: NewLaneOverridesForV2(selB), ChainBOverrides: NewLaneOverridesForV2(selC)},
			},
		},
	})
	require.NoError(t, err)
	MergeAddresses(t, e, connectOut.DataStore)

	// Legacy v1.5.1 BurnMint web A<->B, A<->C (B and C pools land at the SAME address). The legacy
	// pools are transferred to the timelock so the OnlyOwner reverse-propagation writes are emitted
	// as MCMS proposals (rather than executed eagerly by the deployer), matching prod.
	legacyWeb := map[uint64]tokensapi.TokenExpansionInputPerChain{}
	for _, src := range []uint64{selA, selB, selC} {
		sym := fmt.Sprintf("COLLIDE_TOKEN_%d", src)
		legacyWeb[src] = tokensapi.TokenExpansionInputPerChain{
			SkipOwnershipTransfer: false,
			TokenPoolVersion:      cciputils.Version_1_5_1,
			DeployTokenInput: &tokensapi.DeployTokenInput{
				Name: sym, Symbol: sym, Decimals: 18, Type: bnmERC20ops.ContractType, Supply: nil,
			},
			DeployTokenPoolInput: &tokensapi.DeployTokenPoolInput{
				TokenPoolQualifier: fmt.Sprintf("COLLIDE_POOL_V1_%d", src),
				PoolType:           cciputils.BurnMintTokenPool.String(),
			},
			TokenTransferConfig: &tokensapi.TokenTransferConfig{
				RemoteChains: map[uint64]tokensapi.RemoteChainConfig[*datastore.AddressRef, datastore.AddressRef]{},
			},
		}
		for _, dst := range []uint64{selA, selB, selC} {
			if src != dst {
				legacyWeb[src].TokenTransferConfig.RemoteChains[dst] = tokensapi.RemoteChainConfig[*datastore.AddressRef, datastore.AddressRef]{}
			}
		}
	}
	e.OperationsBundle = testsetupV2_0_0.BundleWithFreshReporter(e.OperationsBundle)
	webOut, err := tokensapi.TokenExpansion().Apply(*e, tokensapi.TokenExpansionInput{
		TokenExpansionInputPerChain: legacyWeb,
		ChainAdapterVersion:         cciputils.Version_1_6_0,
		MCMS:                        NewDefaultInputForMCMS("legacy web"),
	})
	require.NoError(t, err)
	testhelpers.ProcessTimelockProposals(t, *e, webOut.MCMSTimelockProposals, false)
	MergeAddresses(t, e, webOut.DataStore)

	// Precondition: B and C's legacy pools share the same address, and A's is distinct.
	poolB := webOut.DataStore.Addresses().Filter(
		datastore.AddressRefByChainSelector(selB), datastore.AddressRefByType(datastore.ContractType(cciputils.BurnMintTokenPool)))[0].Address
	poolC := webOut.DataStore.Addresses().Filter(
		datastore.AddressRefByChainSelector(selC), datastore.AddressRefByType(datastore.ContractType(cciputils.BurnMintTokenPool)))[0].Address
	require.Equal(t, poolB, poolC, "B and C counterpart pools must share an address to reproduce the bug")

	tokenRefA, err := datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{Qualifier: legacyWeb[selA].DeployTokenInput.Symbol}, selA, datastore_utils.FullRef)
	require.NoError(t, err)

	legacyPoolA := webOut.DataStore.Addresses().Filter(
		datastore.AddressRefByChainSelector(selA), datastore.AddressRefByType(datastore.ContractType(cciputils.BurnMintTokenPool)))[0]
	require.NotEqual(t, legacyPoolA.Address, poolB, "hub A's pool must be distinct from the colliding B/C pools")

	// Precondition: A's legacy pool is connected to both B and C, and is the active pool in the TAR.
	poolA, err := tokenpoolV2_0_0.NewTokenPool(common.HexToAddress(legacyPoolA.Address), e.BlockChains.EVMChains()[selA].Client)
	require.NoError(t, err)
	sc, err := poolA.GetSupportedChains(&bind.CallOpts{Context: t.Context()})
	require.NoError(t, err)
	require.ElementsMatch(t, []uint64{selB, selC}, sc, "legacy pool A should support B and C before migration")

	// Precondition: A's legacy pool must be active in the TAR
	tarAddr := connectOut.DataStore.Addresses().Filter(
		datastore.AddressRefByChainSelector(selA),
		datastore.AddressRefByType(datastore.ContractType("TokenAdminRegistry")))[0].Address
	tarA, err := tarbindings.NewTokenAdminRegistry(common.HexToAddress(tarAddr), e.BlockChains.EVMChains()[selA].Client)
	require.NoError(t, err)
	tarCfg, err := tarA.GetTokenConfig(&bind.CallOpts{Context: t.Context()}, common.HexToAddress(tokenRefA.Address))
	require.NoError(t, err)
	require.Equal(t, common.HexToAddress(legacyPoolA.Address), tarCfg.TokenPool, "TAR active pool for A must be the legacy pool before migration")

	// Migrate ONLY A -> v2 with autoMigrateRemoteChains (remotes B and C discovered from the active pool).
	e.OperationsBundle = testsetupV2_0_0.BundleWithFreshReporter(e.OperationsBundle)
	migOut, err := tokensapi.TokenExpansion().Apply(*e, tokensapi.TokenExpansionInput{
		ChainAdapterVersion: cciputils.Version_2_0_0,
		MCMS:                NewDefaultInputForMCMS("collision regression: migrate A"),
		TokenExpansionInputPerChain: map[uint64]tokensapi.TokenExpansionInputPerChain{
			selA: {
				SkipOwnershipTransfer: true,
				TokenPoolVersion:      cciputils.Version_2_0_0,
				DeployTokenPoolInput: &tokensapi.DeployTokenPoolInput{
					TokenPoolQualifier: fmt.Sprintf("COLLIDE_POOL_V2_%d", selA),
					PoolType:           cciputils.BurnMintTokenPool.String(),
					TokenRef:           &tokenRefA,
				},
				TokenTransferConfig: &tokensapi.TokenTransferConfig{
					AutoMigrateRemoteChains: true,
					RemoteChains:            nil,
				},
			},
		},
	})
	require.NoError(t, err)

	// The reverse propagation must produce exactly one op per counterpart chain, and each op must
	// target THAT chain's counterpart pool and add A's migrated v2 pool back on the hub's selector.
	// Before the fix the colliding addresses made B and C share a report-cache key, so one chain
	// received two ops and the other none. Note poolB == poolC, so the ops are distinguished by
	// op.ChainSelector, not by the target address.
	matchesForNewPoolA := migOut.DataStore.Addresses().Filter(
		datastore.AddressRefByChainSelector(selA),
		datastore.AddressRefByType(datastore.ContractType(cciputils.BurnMintTokenPool)),
		datastore.AddressRefByQualifier(fmt.Sprintf("COLLIDE_POOL_V2_%d", selA)),
	)
	require.Len(t, matchesForNewPoolA, 1, "expected exactly one migrated v2 pool for A")
	expectedRemotePool := common.LeftPadBytes(common.HexToAddress(matchesForNewPoolA[0].Address).Bytes(), 32)

	poolABI, err := tokenpoolV2_0_0.TokenPoolMetaData.GetAbi()
	require.NoError(t, err)

	expectedPoolPerChain := map[mcms_types.ChainSelector]string{
		mcms_types.ChainSelector(selB): poolB,
		mcms_types.ChainSelector(selC): poolC,
	}

	sentPerChain := map[mcms_types.ChainSelector]int{}
	for _, prop := range migOut.MCMSTimelockProposals {
		for _, op := range prop.Operations {
			wantPool, ok := expectedPoolPerChain[op.ChainSelector]
			if !ok {
				continue
			}

			sentPerChain[op.ChainSelector] += len(op.Transactions)
			require.Len(t, op.Transactions, 1, "expected a single reverse-propagation tx for chain %d", op.ChainSelector)

			tx := op.Transactions[0]
			require.Equal(t, common.HexToAddress(wantPool), common.HexToAddress(tx.To),
				"reverse-propagation op for chain %d must target that chain's counterpart pool", op.ChainSelector)

			method, err := poolABI.MethodById(tx.Data[:4])
			require.NoError(t, err)
			require.Equal(t, "addRemotePool", method.Name, "reverse propagation should add the migrated hub pool")
			args, err := method.Inputs.Unpack(tx.Data[4:])
			require.NoError(t, err)

			sourceSelector, ok := args[0].(uint64)
			require.True(t, ok, "first arg must be a uint64 chain selector")
			require.Equal(t, selA, sourceSelector, "remoteChainSelector must be the migrating hub A")

			remotePoolBytes, ok := args[1].([]byte)
			require.True(t, ok, "second arg must be a bytes32 remote pool address")
			require.Equal(t, expectedRemotePool, remotePoolBytes, "must add A's migrated v2 pool as the remote pool")
		}
	}

	require.Equal(t, 1, sentPerChain[mcms_types.ChainSelector(selB)], "expected exactly one reverse-propagation op for chain B")
	require.Equal(t, 1, sentPerChain[mcms_types.ChainSelector(selC)], "expected exactly one reverse-propagation op for chain C")
}
