package deployment

import (
	"bytes"
	"fmt"
	"math/big"
	"slices"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gagliardetto/solana-go"
	chainsel "github.com/smartcontractkit/chain-selectors"
	solchain "github.com/smartcontractkit/chainlink-deployments-framework/chain/solana"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	"github.com/stretchr/testify/require"

	evmadapters "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/adapters"
	bnmERC20ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/burn_mint_erc20"
	erc20ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/erc20"
	evmseqV1_6_0 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/sequences"
	bnmpool "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_5_1/burn_mint_token_pool"
	lrpool "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_5_1/lock_release_token_pool"
	solanautils "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/utils"
	routerops "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_0/operations/router"
	solseqV1_6_0 "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_0/sequences"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/v0_1_1/ccip_common"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/v1_6_0/burnmint_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/v1_6_0/lockrelease_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/state"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/tokens"
	"github.com/smartcontractkit/chainlink-ccip/deployment/testhelpers"
	cciputils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"

	deployapi "github.com/smartcontractkit/chainlink-ccip/deployment/deploy"
	tokensapi "github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"

	bnmERC20gen "github.com/smartcontractkit/chainlink-evm/gethwrappers/shared/generated/initial/burn_mint_erc20"

	_ "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_1/adapters"
	_ "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_1/adapters"

	// Registers SolanaAddressNormalizer (v1_6_0 sequences do not import v1_0_0 adapters).
	_ "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_0_0/adapters"
)

// ---------------------------------------------------------------------------
// Shared helpers
// ---------------------------------------------------------------------------

var (
	v1_6_0_scenarios = cciputils.Version_1_6_0
	v1_5_1_scenarios = cciputils.Version_1_5_1
)

// setupEVMOnlyEnv creates a 2-chain EVM test environment with core contracts,
// MCMS, and ownership transferred. Returns the env + both chain selectors.
func setupEVMOnlyEnv(t *testing.T) (*deployment.Environment, uint64, uint64) {
	t.Helper()

	selA := chainsel.TEST_90000001.Selector
	selB := chainsel.TEST_90000002.Selector

	env, err := environment.New(
		t.Context(),
		environment.WithEVMSimulated(t, []uint64{selA, selB}),
	)
	require.NoError(t, err)

	evmAdapter := evmseqV1_6_0.EVMAdapter{}

	deployRegistry := deployapi.GetRegistry()
	deployRegistry.RegisterDeployer(chainsel.FamilyEVM, deployapi.MCMSVersion, &evmadapters.EVMDeployer{})

	mcmsRegistry := changesets.GetRegistry()
	mcmsRegistry.RegisterMCMSReader(chainsel.FamilyEVM, &evmadapters.EVMMCMSReader{})

	deployInput := map[uint64]deployapi.ContractDeploymentConfigPerChain{
		selA: NewDefaultDeploymentConfigForEVM(v1_6_0_scenarios),
		selB: NewDefaultDeploymentConfigForEVM(v1_6_0_scenarios),
	}
	output, err := deployapi.DeployContracts(deployRegistry).Apply(*env, deployapi.ContractDeploymentConfig{
		Chains: deployInput,
		MCMS:   mcms.Input{},
	})
	require.NoError(t, err)
	MergeAddresses(t, env, output.DataStore)

	for _, sel := range []uint64{selA, selB} {
		DeployMCMS(t, env, sel, []string{cciputils.CLLQualifier})
	}
	for _, sel := range []uint64{selA, selB} {
		EVMTransferOwnership(t, env, sel)
	}

	// Sanity-check TAR deployed
	for _, sel := range []uint64{selA, selB} {
		_, err := evmAdapter.GetTokenAdminRegistryAddress(env.DataStore, sel)
		require.NoError(t, err)
	}

	return env, selA, selB
}

// assertTokenOnlyExistsOnChain verifies a token does not exist in the datastore
// yet, but does exist on-chain with the expected name, symbol, and decimals.
func assertTokenOnlyExistsOnChain(
	t *testing.T,
	out deployment.ChangesetOutput,
	env *deployment.Environment,
	chainSel uint64,
	symbol string,
	expectedName string,
	expectedDecimals uint8,
) common.Address {
	t.Helper()

	evmAdapter := evmseqV1_6_0.EVMAdapter{}
	_, err := evmAdapter.FindOneTokenAddress(env.DataStore, chainSel, &datastore.AddressRef{Qualifier: symbol})
	require.Error(t, err, "token %q should not exist in datastore on chain %d", symbol, chainSel)
	tokAddress, err := evmAdapter.FindOneTokenAddress(out.DataStore.Seal(), chainSel, &datastore.AddressRef{Qualifier: symbol})
	require.NoError(t, err, "token %q should exist in output datastore on chain %d", symbol, chainSel)

	chain, ok := env.BlockChains.EVMChains()[chainSel]
	require.True(t, ok)
	tokn, err := bnmERC20gen.NewBurnMintERC20(tokAddress, chain.Client)
	require.NoError(t, err)
	name, err := tokn.Name(&bind.CallOpts{Context: t.Context()})
	require.NoError(t, err)
	require.Equal(t, expectedName, name)
	symb, err := tokn.Symbol(&bind.CallOpts{Context: t.Context()})
	require.NoError(t, err)
	require.Equal(t, symbol, symb)
	deci, err := tokn.Decimals(&bind.CallOpts{Context: t.Context()})
	require.NoError(t, err)
	require.Equal(t, expectedDecimals, deci)

	return tokAddress
}

// assertTokenExists verifies a token exists in the datastore and matches
// expected name, symbol, and decimals on-chain.
func assertTokenExists(
	t *testing.T,
	env *deployment.Environment,
	chainSel uint64,
	symbol string,
	expectedName string,
	expectedDecimals uint8,
) common.Address {
	t.Helper()

	evmAdapter := evmseqV1_6_0.EVMAdapter{}
	tokAddress, err := evmAdapter.FindOneTokenAddress(env.DataStore, chainSel, &datastore.AddressRef{Qualifier: symbol})
	require.NoError(t, err, "token %q should exist in datastore on chain %d", symbol, chainSel)

	chain, ok := env.BlockChains.EVMChains()[chainSel]
	require.True(t, ok)
	tokn, err := bnmERC20gen.NewBurnMintERC20(tokAddress, chain.Client)
	require.NoError(t, err)
	name, err := tokn.Name(&bind.CallOpts{Context: t.Context()})
	require.NoError(t, err)
	require.Equal(t, expectedName, name)
	symb, err := tokn.Symbol(&bind.CallOpts{Context: t.Context()})
	require.NoError(t, err)
	require.Equal(t, symbol, symb)
	deci, err := tokn.Decimals(&bind.CallOpts{Context: t.Context()})
	require.NoError(t, err)
	require.Equal(t, expectedDecimals, deci)

	return tokAddress
}

type poolContract interface {
	GetToken(opts *bind.CallOpts) (common.Address, error)
	GetTokenDecimals(opts *bind.CallOpts) (uint8, error)
	GetSupportedChains(opts *bind.CallOpts) ([]uint64, error)
	GetRemotePools(opts *bind.CallOpts, remoteChainSelector uint64) ([][]byte, error)
	GetRemoteToken(opts *bind.CallOpts, remoteChainSelector uint64) ([]byte, error)
}

// assertPoolConnected verifies that a pool on chainSel has remoteChainSel as a
// supported chain, and that the remote pool and remote token bytes match.
func assertPoolConnected(
	t *testing.T,
	pool poolContract,
	chainSel uint64,
	remoteChainSel uint64,
	expectedRemotePoolAddr common.Address,
	expectedRemoteTokenAddr common.Address,
) {
	t.Helper()
	ctx := t.Context()

	supportedChains, err := pool.GetSupportedChains(&bind.CallOpts{Context: ctx})
	require.NoError(t, err)
	require.Contains(t, supportedChains, remoteChainSel, "pool on chain %d should support remote chain %d", chainSel, remoteChainSel)

	remotePools, err := pool.GetRemotePools(&bind.CallOpts{Context: ctx}, remoteChainSel)
	require.NoError(t, err)
	require.Len(t, remotePools, 1)
	require.True(t, bytes.Equal(remotePools[0], common.LeftPadBytes(expectedRemotePoolAddr.Bytes(), 32)),
		"remote pool on chain %d should be %s", chainSel, expectedRemotePoolAddr.Hex())

	remoteToken, err := pool.GetRemoteToken(&bind.CallOpts{Context: ctx}, remoteChainSel)
	require.NoError(t, err)
	require.True(t, bytes.Equal(remoteToken, common.LeftPadBytes(expectedRemoteTokenAddr.Bytes(), 32)),
		"remote token on chain %d should be %s", chainSel, expectedRemoteTokenAddr.Hex())
}

// assertBMRateLimitsEnabled verifies rate limits on a BurnMint pool.
func assertBMRateLimitsEnabled(t *testing.T, pool *bnmpool.BurnMintTokenPool, remoteChainSel uint64) {
	t.Helper()
	ctx := t.Context()
	outbound, err := pool.GetCurrentOutboundRateLimiterState(&bind.CallOpts{Context: ctx}, remoteChainSel)
	require.NoError(t, err)
	require.True(t, outbound.IsEnabled, "outbound rate limit should be enabled for remote chain %d", remoteChainSel)
	inbound, err := pool.GetCurrentInboundRateLimiterState(&bind.CallOpts{Context: ctx}, remoteChainSel)
	require.NoError(t, err)
	require.True(t, inbound.IsEnabled, "inbound rate limit should be enabled for remote chain %d", remoteChainSel)
}

// assertLRRateLimitsEnabled verifies rate limits on a LockRelease pool.
func assertLRRateLimitsEnabled(t *testing.T, pool *lrpool.LockReleaseTokenPool, remoteChainSel uint64) {
	t.Helper()
	ctx := t.Context()
	outbound, err := pool.GetCurrentOutboundRateLimiterState(&bind.CallOpts{Context: ctx}, remoteChainSel)
	require.NoError(t, err)
	require.True(t, outbound.IsEnabled, "outbound rate limit should be enabled for remote chain %d", remoteChainSel)
	inbound, err := pool.GetCurrentInboundRateLimiterState(&bind.CallOpts{Context: ctx}, remoteChainSel)
	require.NoError(t, err)
	require.True(t, inbound.IsEnabled, "inbound rate limit should be enabled for remote chain %d", remoteChainSel)
}

// assertPoolNotConnected verifies a pool has no supported remote chains.
func assertPoolNotConnected(t *testing.T, pool poolContract) {
	t.Helper()
	supportedChains, err := pool.GetSupportedChains(&bind.CallOpts{Context: t.Context()})
	require.NoError(t, err)
	require.Empty(t, supportedChains, "pool should have no supported chains yet")
}

// ---------------------------------------------------------------------------
// EVM-only scenarios (1-4)
// ---------------------------------------------------------------------------

func TestTokenExpansionScenariosEVM(t *testing.T) {
	env, selA, selB := setupEVMOnlyEnv(t)
	evmAdapter := evmseqV1_6_0.EVMAdapter{}

	defaultRL := tokensapi.RateLimiterConfigFloatInput{
		Capacity:  100,
		Rate:      10,
		IsEnabled: true,
	}
	defaultMaxSupply := uint64(1e6)
	defaultPreMint := uint64(1e5)

	bmPoolType := cciputils.BurnMintTokenPool
	lrPoolType := cciputils.LockReleaseTokenPool

	// -----------------------------------------------------------------------
	// Scenario 1: Fresh deploy token + pool on two chains, connect in one call
	// -----------------------------------------------------------------------
	t.Run("Scenario1_FreshDeployAndConnect", func(t *testing.T) {
		tokenSymbolA := "S1_TOK_A"
		tokenSymbolB := "S1_TOK_B"
		poolQualA := "S1_POOL_A"
		poolQualB := "S1_POOL_B"

		output, err := tokensapi.TokenExpansion().Apply(*env, tokensapi.TokenExpansionInput{
			ChainAdapterVersion: v1_6_0_scenarios,
			MCMS:                NewDefaultInputForMCMS("Scenario 1"),
			TokenExpansionInputPerChain: map[uint64]tokensapi.TokenExpansionInputPerChain{
				selA: {
					TokenPoolVersion: v1_5_1_scenarios,
					DeployTokenInput: &tokensapi.DeployTokenInput{
						Name:     "Scenario1 Token A",
						Symbol:   tokenSymbolA,
						Decimals: 18,
						Type:     bnmERC20ops.ContractType,
						Supply:   &defaultMaxSupply,
						PreMint:  &defaultPreMint,
					},
					DeployTokenPoolInput: &tokensapi.DeployTokenPoolInput{
						TokenPoolQualifier: poolQualA,
						PoolType:           bmPoolType.String(),
					},
					TokenTransferConfig: &tokensapi.TokenTransferConfig{
						RemoteChains: map[uint64]tokensapi.RemoteChainConfig[*datastore.AddressRef, datastore.AddressRef]{
							selB: {OutboundRateLimiterConfig: &defaultRL},
						},
					},
				},
				selB: {
					TokenPoolVersion: v1_5_1_scenarios,
					DeployTokenInput: &tokensapi.DeployTokenInput{
						Name:     "Scenario1 Token B",
						Symbol:   tokenSymbolB,
						Decimals: 18,
						Type:     bnmERC20ops.ContractType,
						Supply:   &defaultMaxSupply,
						PreMint:  &defaultPreMint,
					},
					DeployTokenPoolInput: &tokensapi.DeployTokenPoolInput{
						TokenPoolQualifier: poolQualB,
						PoolType:           bmPoolType.String(),
					},
					TokenTransferConfig: &tokensapi.TokenTransferConfig{
						RemoteChains: map[uint64]tokensapi.RemoteChainConfig[*datastore.AddressRef, datastore.AddressRef]{
							selA: {OutboundRateLimiterConfig: &defaultRL},
						},
					},
				},
			},
		})
		require.NoError(t, err)
		MergeAddresses(t, env, output.DataStore)
		testhelpers.ProcessTimelockProposals(t, *env, output.MCMSTimelockProposals, false)

		// Assert tokens exist
		tokAddrA := assertTokenExists(t, env, selA, tokenSymbolA, "Scenario1 Token A", 18)
		tokAddrB := assertTokenExists(t, env, selB, tokenSymbolB, "Scenario1 Token B", 18)

		// Get pool instances
		chainA := env.BlockChains.EVMChains()[selA]
		chainB := env.BlockChains.EVMChains()[selB]
		poolAddrA, err := evmAdapter.FindLatestAddressRef(env.DataStore, datastore.AddressRef{ChainSelector: selA, Qualifier: poolQualA, Type: datastore.ContractType(bmPoolType)})
		require.NoError(t, err)
		poolAddrB, err := evmAdapter.FindLatestAddressRef(env.DataStore, datastore.AddressRef{ChainSelector: selB, Qualifier: poolQualB, Type: datastore.ContractType(bmPoolType)})
		require.NoError(t, err)

		poolA, err := bnmpool.NewBurnMintTokenPool(poolAddrA, chainA.Client)
		require.NoError(t, err)
		poolB, err := bnmpool.NewBurnMintTokenPool(poolAddrB, chainB.Client)
		require.NoError(t, err)

		// Verify pool token binding
		gotTokA, err := poolA.GetToken(&bind.CallOpts{Context: t.Context()})
		require.NoError(t, err)
		require.Equal(t, tokAddrA, gotTokA)
		gotTokB, err := poolB.GetToken(&bind.CallOpts{Context: t.Context()})
		require.NoError(t, err)
		require.Equal(t, tokAddrB, gotTokB)

		// Verify bidirectional connectivity
		assertPoolConnected(t, poolA, selA, selB, poolAddrB, tokAddrB)
		assertPoolConnected(t, poolB, selB, selA, poolAddrA, tokAddrA)

		// Verify rate limits
		assertBMRateLimitsEnabled(t, poolA, selB)
		assertBMRateLimitsEnabled(t, poolB, selA)
	})

	// -----------------------------------------------------------------------
	// Scenario 2: Add new LockRelease pools to existing BurnMint setup
	// -----------------------------------------------------------------------
	t.Run("Scenario2_AddLRPoolsToExistingBMSetup", func(t *testing.T) {
		// First, deploy BurnMint tokens + pools and connect them (baseline)
		tokenSymbolA := "S2_TOK_A"
		tokenSymbolB := "S2_TOK_B"
		bmPoolQualA := "S2_BM_POOL_A"
		bmPoolQualB := "S2_BM_POOL_B"

		output, err := tokensapi.TokenExpansion().Apply(*env, tokensapi.TokenExpansionInput{
			ChainAdapterVersion: v1_6_0_scenarios,
			MCMS:                NewDefaultInputForMCMS("Scenario 2 BM setup"),
			TokenExpansionInputPerChain: map[uint64]tokensapi.TokenExpansionInputPerChain{
				selA: {
					TokenPoolVersion: v1_5_1_scenarios,
					DeployTokenInput: &tokensapi.DeployTokenInput{
						Name: "Scenario2 Token A", Symbol: tokenSymbolA, Decimals: 18,
						Type: bnmERC20ops.ContractType, Supply: &defaultMaxSupply,
					},
					DeployTokenPoolInput: &tokensapi.DeployTokenPoolInput{
						TokenPoolQualifier: bmPoolQualA, PoolType: bmPoolType.String(),
					},
					TokenTransferConfig: &tokensapi.TokenTransferConfig{
						RemoteChains: map[uint64]tokensapi.RemoteChainConfig[*datastore.AddressRef, datastore.AddressRef]{
							selB: {OutboundRateLimiterConfig: &defaultRL},
						},
					},
				},
				selB: {
					TokenPoolVersion: v1_5_1_scenarios,
					DeployTokenInput: &tokensapi.DeployTokenInput{
						Name: "Scenario2 Token B", Symbol: tokenSymbolB, Decimals: 18,
						Type: bnmERC20ops.ContractType, Supply: &defaultMaxSupply,
					},
					DeployTokenPoolInput: &tokensapi.DeployTokenPoolInput{
						TokenPoolQualifier: bmPoolQualB, PoolType: bmPoolType.String(),
					},
					TokenTransferConfig: &tokensapi.TokenTransferConfig{
						RemoteChains: map[uint64]tokensapi.RemoteChainConfig[*datastore.AddressRef, datastore.AddressRef]{
							selA: {OutboundRateLimiterConfig: &defaultRL},
						},
					},
				},
			},
		})
		require.NoError(t, err)
		MergeAddresses(t, env, output.DataStore)
		testhelpers.ProcessTimelockProposals(t, *env, output.MCMSTimelockProposals, false)

		// Get BM pool addresses for later comparison
		bmPoolAddrA, err := evmAdapter.FindLatestAddressRef(env.DataStore, datastore.AddressRef{ChainSelector: selA, Qualifier: bmPoolQualA, Type: datastore.ContractType(bmPoolType)})
		require.NoError(t, err)
		bmPoolAddrB, err := evmAdapter.FindLatestAddressRef(env.DataStore, datastore.AddressRef{ChainSelector: selB, Qualifier: bmPoolQualB, Type: datastore.ContractType(bmPoolType)})
		require.NoError(t, err)
		tokAddrA := assertTokenExists(t, env, selA, tokenSymbolA, "Scenario2 Token A", 18)
		tokAddrB := assertTokenExists(t, env, selB, tokenSymbolB, "Scenario2 Token B", 18)

		// Now deploy LockRelease pools backed by the same tokens
		lrPoolQualA := "S2_LR_POOL_A"
		lrPoolQualB := "S2_LR_POOL_B"
		acceptLiquidity := true

		output, err = tokensapi.TokenExpansion().Apply(*env, tokensapi.TokenExpansionInput{
			ChainAdapterVersion: v1_6_0_scenarios,
			MCMS:                NewDefaultInputForMCMS("Scenario 2 LR addition"),
			TokenExpansionInputPerChain: map[uint64]tokensapi.TokenExpansionInputPerChain{
				selA: {
					TokenPoolVersion: v1_5_1_scenarios,
					DeployTokenPoolInput: &tokensapi.DeployTokenPoolInput{
						TokenPoolQualifier: lrPoolQualA,
						PoolType:           lrPoolType.String(),
						AcceptLiquidity:    &acceptLiquidity,
						TokenRef: &datastore.AddressRef{
							Qualifier: tokenSymbolA,
							Type:      datastore.ContractType(bnmERC20ops.ContractType),
						},
					},
					TokenTransferConfig: &tokensapi.TokenTransferConfig{
						RemoteChains: map[uint64]tokensapi.RemoteChainConfig[*datastore.AddressRef, datastore.AddressRef]{
							selB: {OutboundRateLimiterConfig: &defaultRL},
						},
					},
				},
				selB: {
					TokenPoolVersion: v1_5_1_scenarios,
					DeployTokenPoolInput: &tokensapi.DeployTokenPoolInput{
						TokenPoolQualifier: lrPoolQualB,
						PoolType:           lrPoolType.String(),
						AcceptLiquidity:    &acceptLiquidity,
						TokenRef: &datastore.AddressRef{
							Qualifier: tokenSymbolB,
							Type:      datastore.ContractType(bnmERC20ops.ContractType),
						},
					},
					TokenTransferConfig: &tokensapi.TokenTransferConfig{
						RemoteChains: map[uint64]tokensapi.RemoteChainConfig[*datastore.AddressRef, datastore.AddressRef]{
							selA: {OutboundRateLimiterConfig: &defaultRL},
						},
					},
				},
			},
		})
		require.NoError(t, err)
		MergeAddresses(t, env, output.DataStore)
		testhelpers.ProcessTimelockProposals(t, *env, output.MCMSTimelockProposals, false)

		// Verify new LR pools exist and are backed by the same token
		chainA := env.BlockChains.EVMChains()[selA]
		chainB := env.BlockChains.EVMChains()[selB]

		lrPoolAddrA, err := evmAdapter.FindLatestAddressRef(env.DataStore, datastore.AddressRef{ChainSelector: selA, Qualifier: lrPoolQualA, Type: datastore.ContractType(lrPoolType)})
		require.NoError(t, err)
		lrPoolAddrB, err := evmAdapter.FindLatestAddressRef(env.DataStore, datastore.AddressRef{ChainSelector: selB, Qualifier: lrPoolQualB, Type: datastore.ContractType(lrPoolType)})
		require.NoError(t, err)

		lrA, err := lrpool.NewLockReleaseTokenPool(lrPoolAddrA, chainA.Client)
		require.NoError(t, err)
		lrB, err := lrpool.NewLockReleaseTokenPool(lrPoolAddrB, chainB.Client)
		require.NoError(t, err)

		gotTokA, err := lrA.GetToken(&bind.CallOpts{Context: t.Context()})
		require.NoError(t, err)
		require.Equal(t, tokAddrA, gotTokA, "LR pool on A should be backed by the same token")
		gotTokB, err := lrB.GetToken(&bind.CallOpts{Context: t.Context()})
		require.NoError(t, err)
		require.Equal(t, tokAddrB, gotTokB, "LR pool on B should be backed by the same token")

		// LR pools should be connected to each other
		assertPoolConnected(t, lrA, selA, selB, lrPoolAddrB, tokAddrB)
		assertPoolConnected(t, lrB, selB, selA, lrPoolAddrA, tokAddrA)
		assertLRRateLimitsEnabled(t, lrA, selB)
		assertLRRateLimitsEnabled(t, lrB, selA)

		// Original BM pools should still be connected and unaffected
		bmA, err := bnmpool.NewBurnMintTokenPool(bmPoolAddrA, chainA.Client)
		require.NoError(t, err)
		bmB, err := bnmpool.NewBurnMintTokenPool(bmPoolAddrB, chainB.Client)
		require.NoError(t, err)
		assertPoolConnected(t, bmA, selA, selB, bmPoolAddrB, tokAddrB)
		assertPoolConnected(t, bmB, selB, selA, bmPoolAddrA, tokAddrA)
	})

	// -----------------------------------------------------------------------
	// Scenario 3: Deploy pool backed by pre-existing token
	// -----------------------------------------------------------------------
	t.Run("Scenario3_PoolBackedByExistingToken", func(t *testing.T) {
		tokenSymbolA := "S3_TOK_A"
		tokenSymbolB := "S3_TOK_B"

		// First call: deploy token A but do NOT merge it into the environment datastore
			ChainAdapterVersion: v1_6_0_scenarios,
			MCMS:                NewDefaultInputForMCMS("Scenario 3 tokens"),
			TokenExpansionInputPerChain: map[uint64]tokensapi.TokenExpansionInputPerChain{
				selA: {
					TokenPoolVersion: v1_5_1_scenarios,
					DeployTokenInput: &tokensapi.DeployTokenInput{
						Name: "Scenario3 Token A", Symbol: tokenSymbolA, Decimals: 18,
						Type: bnmERC20ops.ContractType, Supply: &defaultMaxSupply,
					},
				},
			},
		})
		require.NoError(t, err)
		testhelpers.ProcessTimelockProposals(t, *env, output1.MCMSTimelockProposals, false)

		// Second call: deploy token B and add it to the datastore
		output2, err := tokensapi.TokenExpansion().Apply(*env, tokensapi.TokenExpansionInput{
			ChainAdapterVersion: v1_6_0_scenarios,
			MCMS:                NewDefaultInputForMCMS("Scenario 3 tokens"),
			TokenExpansionInputPerChain: map[uint64]tokensapi.TokenExpansionInputPerChain{
				selB: {
					TokenPoolVersion: v1_5_1_scenarios,
					DeployTokenInput: &tokensapi.DeployTokenInput{
						Name: "Scenario3 Token B", Symbol: tokenSymbolB, Decimals: 18,
						Type: bnmERC20ops.ContractType, Supply: &defaultMaxSupply,
					},
				},
			},
		})
		require.NoError(t, err)
		MergeAddresses(t, env, output2.DataStore)
		testhelpers.ProcessTimelockProposals(t, *env, output2.MCMSTimelockProposals, false)

		// Tokens should exist, no pools yet
		tokAddrA := assertTokenOnlyExistsOnChain(t, output1, env, selA, tokenSymbolA, "Scenario3 Token A", 18)
		tokAddrB := assertTokenExists(t, env, selB, tokenSymbolB, "Scenario3 Token B", 18)
		_, err = evmAdapter.FindLatestAddressRef(env.DataStore, datastore.AddressRef{ChainSelector: selA, Qualifier: "S3_POOL_A", Type: datastore.ContractType(bmPoolType)})
		require.Error(t, err, "pool should not exist after token-only deploy")

		// Third call: deploy pools referencing existing tokens, and connect
		poolQualA := "S3_POOL_A"
		poolQualB := "S3_POOL_B"

		output, err := tokensapi.TokenExpansion().Apply(*env, tokensapi.TokenExpansionInput{
			ChainAdapterVersion: v1_6_0_scenarios,
			MCMS:                NewDefaultInputForMCMS("Scenario 3 pools"),
			TokenExpansionInputPerChain: map[uint64]tokensapi.TokenExpansionInputPerChain{
				selA: {
					TokenPoolVersion: v1_5_1_scenarios,
					DeployTokenPoolInput: &tokensapi.DeployTokenPoolInput{
						TokenPoolQualifier: poolQualA,
						PoolType:           bmPoolType.String(),
						TokenRef: &datastore.AddressRef{
							// A pool deployment should still be possible even if the token ref
							// is not in the datastore and the input only provides the address.
							Address: tokAddrA.Hex(),
						},
					},
					TokenTransferConfig: &tokensapi.TokenTransferConfig{
						RemoteChains: map[uint64]tokensapi.RemoteChainConfig[*datastore.AddressRef, datastore.AddressRef]{
							selB: {OutboundRateLimiterConfig: &defaultRL},
						},
					},
				},
				selB: {
					TokenPoolVersion: v1_5_1_scenarios,
					DeployTokenPoolInput: &tokensapi.DeployTokenPoolInput{
						TokenPoolQualifier: poolQualB,
						PoolType:           bmPoolType.String(),
						TokenRef: &datastore.AddressRef{
							Qualifier: tokenSymbolB,
							Type:      datastore.ContractType(bnmERC20ops.ContractType),
						},
					},
					TokenTransferConfig: &tokensapi.TokenTransferConfig{
						RemoteChains: map[uint64]tokensapi.RemoteChainConfig[*datastore.AddressRef, datastore.AddressRef]{
							selA: {OutboundRateLimiterConfig: &defaultRL},
						},
					},
				},
			},
		})
		require.NoError(t, err)
		MergeAddresses(t, env, output.DataStore)
		testhelpers.ProcessTimelockProposals(t, *env, output.MCMSTimelockProposals, false)

		chainA := env.BlockChains.EVMChains()[selA]
		chainB := env.BlockChains.EVMChains()[selB]
		poolAddrA, err := evmAdapter.FindLatestAddressRef(env.DataStore, datastore.AddressRef{ChainSelector: selA, Qualifier: poolQualA, Type: datastore.ContractType(bmPoolType)})
		require.NoError(t, err)
		poolAddrB, err := evmAdapter.FindLatestAddressRef(env.DataStore, datastore.AddressRef{ChainSelector: selB, Qualifier: poolQualB, Type: datastore.ContractType(bmPoolType)})
		require.NoError(t, err)

		poolA, err := bnmpool.NewBurnMintTokenPool(poolAddrA, chainA.Client)
		require.NoError(t, err)
		poolB, err := bnmpool.NewBurnMintTokenPool(poolAddrB, chainB.Client)
		require.NoError(t, err)

		// Pools backed by pre-existing tokens
		gotTokA, err := poolA.GetToken(&bind.CallOpts{Context: t.Context()})
		require.NoError(t, err)
		require.Equal(t, tokAddrA, gotTokA)
		gotTokB, err := poolB.GetToken(&bind.CallOpts{Context: t.Context()})
		require.NoError(t, err)
		require.Equal(t, tokAddrB, gotTokB)

		assertPoolConnected(t, poolA, selA, selB, poolAddrB, tokAddrB)
		assertPoolConnected(t, poolB, selB, selA, poolAddrA, tokAddrA)
		assertBMRateLimitsEnabled(t, poolA, selB)
		assertBMRateLimitsEnabled(t, poolB, selA)
	})

	// -----------------------------------------------------------------------
	// Scenario 4: Deploy without connecting, then connect in separate call
	// -----------------------------------------------------------------------
	t.Run("Scenario4_DeferredConnection", func(t *testing.T) {
		tokenSymbolA := "S4_TOK_A"
		tokenSymbolB := "S4_TOK_B"
		poolQualA := "S4_POOL_A"
		poolQualB := "S4_POOL_B"

		// First call: deploy token+pool, no connection
		output, err := tokensapi.TokenExpansion().Apply(*env, tokensapi.TokenExpansionInput{
			ChainAdapterVersion: v1_6_0_scenarios,
			MCMS:                NewDefaultInputForMCMS("Scenario 4 deploy"),
			TokenExpansionInputPerChain: map[uint64]tokensapi.TokenExpansionInputPerChain{
				selA: {
					TokenPoolVersion: v1_5_1_scenarios,
					DeployTokenInput: &tokensapi.DeployTokenInput{
						Name: "Scenario4 Token A", Symbol: tokenSymbolA, Decimals: 18,
						Type: bnmERC20ops.ContractType, Supply: &defaultMaxSupply,
					},
					DeployTokenPoolInput: &tokensapi.DeployTokenPoolInput{
						TokenPoolQualifier: poolQualA,
						PoolType:           bmPoolType.String(),
					},
				},
				selB: {
					TokenPoolVersion: v1_5_1_scenarios,
					DeployTokenInput: &tokensapi.DeployTokenInput{
						Name: "Scenario4 Token B", Symbol: tokenSymbolB, Decimals: 18,
						Type: bnmERC20ops.ContractType, Supply: &defaultMaxSupply,
					},
					DeployTokenPoolInput: &tokensapi.DeployTokenPoolInput{
						TokenPoolQualifier: poolQualB,
						PoolType:           bmPoolType.String(),
					},
				},
			},
		})
		require.NoError(t, err)
		MergeAddresses(t, env, output.DataStore)

		// Pools exist but should have no remote chains configured
		chainA := env.BlockChains.EVMChains()[selA]
		chainB := env.BlockChains.EVMChains()[selB]
		poolAddrA, err := evmAdapter.FindLatestAddressRef(env.DataStore, datastore.AddressRef{ChainSelector: selA, Qualifier: poolQualA, Type: datastore.ContractType(bmPoolType)})
		require.NoError(t, err)
		poolAddrB, err := evmAdapter.FindLatestAddressRef(env.DataStore, datastore.AddressRef{ChainSelector: selB, Qualifier: poolQualB, Type: datastore.ContractType(bmPoolType)})
		require.NoError(t, err)
		poolA, err := bnmpool.NewBurnMintTokenPool(poolAddrA, chainA.Client)
		require.NoError(t, err)
		poolB, err := bnmpool.NewBurnMintTokenPool(poolAddrB, chainB.Client)
		require.NoError(t, err)

		assertPoolNotConnected(t, poolA)
		assertPoolNotConnected(t, poolB)

		tokAddrA := assertTokenExists(t, env, selA, tokenSymbolA, "Scenario4 Token A", 18)
		tokAddrB := assertTokenExists(t, env, selB, tokenSymbolB, "Scenario4 Token B", 18)

		// Second call: connect only (no new deploys)
		output, err = tokensapi.TokenExpansion().Apply(*env, tokensapi.TokenExpansionInput{
			ChainAdapterVersion: v1_6_0_scenarios,
			MCMS:                NewDefaultInputForMCMS("Scenario 4 connect"),
			TokenExpansionInputPerChain: map[uint64]tokensapi.TokenExpansionInputPerChain{
				selA: {
					SkipOwnershipTransfer: true,
					TokenPoolVersion:      v1_5_1_scenarios,
					TokenTransferConfig: &tokensapi.TokenTransferConfig{
						TokenPoolRef: datastore.AddressRef{
							ChainSelector: selA,
							Qualifier:     poolQualA,
							Type:          datastore.ContractType(bmPoolType),
							Version:       v1_5_1_scenarios,
						},
						TokenRef: datastore.AddressRef{
							Qualifier: tokenSymbolA,
							Type:      datastore.ContractType(bnmERC20ops.ContractType),
						},
						RemoteChains: map[uint64]tokensapi.RemoteChainConfig[*datastore.AddressRef, datastore.AddressRef]{
							selB: {
								OutboundRateLimiterConfig: &defaultRL,
								RemoteToken: &datastore.AddressRef{
									ChainSelector: selB,
									Qualifier:     tokenSymbolB,
									Type:          datastore.ContractType(bnmERC20ops.ContractType),
								},
								RemotePool: &datastore.AddressRef{
									ChainSelector: selB,
									Qualifier:     poolQualB,
									Type:          datastore.ContractType(bmPoolType),
									Version:       v1_5_1_scenarios,
								},
							},
						},
					},
				},
				selB: {
					SkipOwnershipTransfer: true,
					TokenPoolVersion:      v1_5_1_scenarios,
					TokenTransferConfig: &tokensapi.TokenTransferConfig{
						TokenPoolRef: datastore.AddressRef{
							ChainSelector: selB,
							Qualifier:     poolQualB,
							Type:          datastore.ContractType(bmPoolType),
							Version:       v1_5_1_scenarios,
						},
						TokenRef: datastore.AddressRef{
							Qualifier: tokenSymbolB,
							Type:      datastore.ContractType(bnmERC20ops.ContractType),
						},
						RemoteChains: map[uint64]tokensapi.RemoteChainConfig[*datastore.AddressRef, datastore.AddressRef]{
							selA: {
								OutboundRateLimiterConfig: &defaultRL,
								RemoteToken: &datastore.AddressRef{
									ChainSelector: selA,
									Qualifier:     tokenSymbolA,
									Type:          datastore.ContractType(bnmERC20ops.ContractType),
								},
								RemotePool: &datastore.AddressRef{
									ChainSelector: selA,
									Qualifier:     poolQualA,
									Type:          datastore.ContractType(bmPoolType),
									Version:       v1_5_1_scenarios,
								},
							},
						},
					},
				},
			},
		})
		require.NoError(t, err)
		MergeAddresses(t, env, output.DataStore)
		testhelpers.ProcessTimelockProposals(t, *env, output.MCMSTimelockProposals, false)

		// Pools should now be connected
		assertPoolConnected(t, poolA, selA, selB, poolAddrB, tokAddrB)
		assertPoolConnected(t, poolB, selB, selA, poolAddrA, tokAddrA)
		assertBMRateLimitsEnabled(t, poolA, selB)
		assertBMRateLimitsEnabled(t, poolB, selA)
	})

	// -----------------------------------------------------------------------
	// Scenario 5: Re-run on pools seeded with unpadded 20-byte remote pool
	// addresses (regression test for the ChainAlreadyExists / padding bug fix)
	// -----------------------------------------------------------------------
	t.Run("Scenario5_RemotePoolPaddingMigration", func(t *testing.T) {
		tokenSymbolA := "S5E_TOK_A"
		tokenSymbolB := "S5E_TOK_B"
		poolQualA := "S5E_POOL_A"
		poolQualB := "S5E_POOL_B"

		// Deploy tokens + pools on both chains without connecting them yet.
		// SkipOwnershipTransfer keeps the deployer as pool owner so we can
		// directly seed the buggy 20-byte address state in the next step.
		output, err := tokensapi.TokenExpansion().Apply(*env, tokensapi.TokenExpansionInput{
			ChainAdapterVersion: v1_6_0_scenarios,
			MCMS:                NewDefaultInputForMCMS("Scenario 5E deploy"),
			TokenExpansionInputPerChain: map[uint64]tokensapi.TokenExpansionInputPerChain{
				selA: {
					SkipOwnershipTransfer: true,
					TokenPoolVersion:      v1_5_1_scenarios,
					DeployTokenInput: &tokensapi.DeployTokenInput{
						Name: "Scenario5E Token A", Symbol: tokenSymbolA, Decimals: 18,
						Type: bnmERC20ops.ContractType, Supply: &defaultMaxSupply,
					},
					DeployTokenPoolInput: &tokensapi.DeployTokenPoolInput{
						TokenPoolQualifier: poolQualA,
						PoolType:           bmPoolType.String(),
					},
				},
				selB: {
					SkipOwnershipTransfer: true,
					TokenPoolVersion:      v1_5_1_scenarios,
					DeployTokenInput: &tokensapi.DeployTokenInput{
						Name: "Scenario5E Token B", Symbol: tokenSymbolB, Decimals: 18,
						Type: bnmERC20ops.ContractType, Supply: &defaultMaxSupply,
					},
					DeployTokenPoolInput: &tokensapi.DeployTokenPoolInput{
						TokenPoolQualifier: poolQualB,
						PoolType:           bmPoolType.String(),
					},
				},
			},
		})
		require.NoError(t, err)
		MergeAddresses(t, env, output.DataStore)

		chainA := env.BlockChains.EVMChains()[selA]
		chainB := env.BlockChains.EVMChains()[selB]

		poolAddrA, err := evmAdapter.FindLatestAddressRef(env.DataStore, datastore.AddressRef{
			ChainSelector: selA, Qualifier: poolQualA, Type: datastore.ContractType(bmPoolType),
		})
		require.NoError(t, err)
		poolAddrB, err := evmAdapter.FindLatestAddressRef(env.DataStore, datastore.AddressRef{
			ChainSelector: selB, Qualifier: poolQualB, Type: datastore.ContractType(bmPoolType),
		})
		require.NoError(t, err)

		tokAddrA := assertTokenExists(t, env, selA, tokenSymbolA, "Scenario5E Token A", 18)
		tokAddrB := assertTokenExists(t, env, selB, tokenSymbolB, "Scenario5E Token B", 18)

		poolA, err := bnmpool.NewBurnMintTokenPool(poolAddrA, chainA.Client)
		require.NoError(t, err)
		poolB, err := bnmpool.NewBurnMintTokenPool(poolAddrB, chainB.Client)
		require.NoError(t, err)

		// Simulate the buggy prior state: call ApplyChainUpdates directly with raw 20-byte
		// (unpadded) pool addresses. CCIP expects 32-byte ABI-encoded (left-padded) addresses,
		// so this mimics what a pre-fix deployment would have stored on-chain.
		disabledRL := bnmpool.RateLimiterConfig{IsEnabled: false, Capacity: big.NewInt(0), Rate: big.NewInt(0)}
		tx, err := poolA.ApplyChainUpdates(chainA.DeployerKey, []uint64{}, []bnmpool.TokenPoolChainUpdate{
			{
				RemoteChainSelector:       selB,
				RemotePoolAddresses:       [][]byte{poolAddrB.Bytes()}, // 20-byte, not left-padded
				RemoteTokenAddress:        common.LeftPadBytes(tokAddrB.Bytes(), 32),
				OutboundRateLimiterConfig: disabledRL,
				InboundRateLimiterConfig:  disabledRL,
			},
		})
		require.NoError(t, err)
		_, err = chainA.Confirm(tx)
		require.NoError(t, err)

		tx, err = poolB.ApplyChainUpdates(chainB.DeployerKey, []uint64{}, []bnmpool.TokenPoolChainUpdate{
			{
				RemoteChainSelector:       selA,
				RemotePoolAddresses:       [][]byte{poolAddrA.Bytes()}, // 20-byte, not left-padded
				RemoteTokenAddress:        common.LeftPadBytes(tokAddrA.Bytes(), 32),
				OutboundRateLimiterConfig: disabledRL,
				InboundRateLimiterConfig:  disabledRL,
			},
		})
		require.NoError(t, err)
		_, err = chainB.Confirm(tx)
		require.NoError(t, err)

		// Verify the seed: only the raw 20-byte entry exists on each pool.
		remotePoolsOnA, err := poolA.GetRemotePools(&bind.CallOpts{Context: t.Context()}, selB)
		require.NoError(t, err)
		require.Len(t, remotePoolsOnA, 1)
		require.Equal(t, poolAddrB.Bytes(), remotePoolsOnA[0], "seed: only the 20-byte entry should exist")

		// Build the connect input once; it is reused for both the fix run and idempotency run.
		disabledOutboundTPRL := tokensapi.RateLimiterConfigFloatInput{IsEnabled: false}
		connectInput := tokensapi.TokenExpansionInput{
			ChainAdapterVersion: v1_6_0_scenarios,
			MCMS:                NewDefaultInputForMCMS("Scenario 5E connect"),
			TokenExpansionInputPerChain: map[uint64]tokensapi.TokenExpansionInputPerChain{
				selA: {
					SkipOwnershipTransfer: true,
					TokenPoolVersion:      v1_5_1_scenarios,
					TokenTransferConfig: &tokensapi.TokenTransferConfig{
						TokenPoolRef: datastore.AddressRef{
							ChainSelector: selA, Qualifier: poolQualA,
							Type: datastore.ContractType(bmPoolType), Version: v1_5_1_scenarios,
						},
						TokenRef: datastore.AddressRef{
							Qualifier: tokenSymbolA,
							Type:      datastore.ContractType(bnmERC20ops.ContractType),
						},
						RemoteChains: map[uint64]tokensapi.RemoteChainConfig[*datastore.AddressRef, datastore.AddressRef]{
							selB: {
								OutboundRateLimiterConfig: &disabledOutboundTPRL,
								RemoteToken: &datastore.AddressRef{
									ChainSelector: selB, Qualifier: tokenSymbolB,
									Type: datastore.ContractType(bnmERC20ops.ContractType),
								},
								RemotePool: &datastore.AddressRef{
									ChainSelector: selB, Qualifier: poolQualB,
									Type: datastore.ContractType(bmPoolType), Version: v1_5_1_scenarios,
								},
							},
						},
					},
				},
				selB: {
					SkipOwnershipTransfer: true,
					TokenPoolVersion:      v1_5_1_scenarios,
					TokenTransferConfig: &tokensapi.TokenTransferConfig{
						TokenPoolRef: datastore.AddressRef{
							ChainSelector: selB, Qualifier: poolQualB,
							Type: datastore.ContractType(bmPoolType), Version: v1_5_1_scenarios,
						},
						TokenRef: datastore.AddressRef{
							Qualifier: tokenSymbolB,
							Type:      datastore.ContractType(bnmERC20ops.ContractType),
						},
						RemoteChains: map[uint64]tokensapi.RemoteChainConfig[*datastore.AddressRef, datastore.AddressRef]{
							selA: {
								OutboundRateLimiterConfig: &disabledOutboundTPRL,
								RemoteToken: &datastore.AddressRef{
									ChainSelector: selA, Qualifier: tokenSymbolA,
									Type: datastore.ContractType(bnmERC20ops.ContractType),
								},
								RemotePool: &datastore.AddressRef{
									ChainSelector: selA, Qualifier: poolQualA,
									Type: datastore.ContractType(bmPoolType), Version: v1_5_1_scenarios,
								},
							},
						},
					},
				},
			},
		}

		// Run TokenExpansion to connect the pools. The exact-comparison fix ensures the sequence
		// recognises that the stored 20-byte entry does not match the required 32-byte padded
		// form and calls AddRemotePool rather than returning early — all without reverting with
		// ChainAlreadyExists.
		output, err = tokensapi.TokenExpansion().Apply(*env, connectInput)
		require.NoError(t, err, "re-configuring pools seeded with 20-byte addresses must not return ChainAlreadyExists")
		MergeAddresses(t, env, output.DataStore)
		testhelpers.ProcessTimelockProposals(t, *env, output.MCMSTimelockProposals, false)

		// The 32-byte padded address must now be registered (alongside the legacy 20-byte entry).
		paddedPoolAddrB := common.LeftPadBytes(poolAddrB.Bytes(), 32)
		remotePoolsOnA, err = poolA.GetRemotePools(&bind.CallOpts{Context: t.Context()}, selB)
		require.NoError(t, err)
		require.True(t, slices.ContainsFunc(remotePoolsOnA, func(rp []byte) bool {
			return bytes.Equal(rp, paddedPoolAddrB)
		}), "32-byte padded pool B address should be registered on pool A after re-configuration")

		paddedPoolAddrA := common.LeftPadBytes(poolAddrA.Bytes(), 32)
		remotePoolsOnB, err := poolB.GetRemotePools(&bind.CallOpts{Context: t.Context()}, selA)
		require.NoError(t, err)
		require.True(t, slices.ContainsFunc(remotePoolsOnB, func(rp []byte) bool {
			return bytes.Equal(rp, paddedPoolAddrA)
		}), "32-byte padded pool A address should be registered on pool B after re-configuration")

		// Idempotency: re-run the same connect config. The 32-byte address is now registered
		// and rate limits are unchanged, so the sequence must return early (no-op). Without the
		// early-return fix this path reverted with ChainAlreadyExists.
		output, err = tokensapi.TokenExpansion().Apply(*env, connectInput)
		require.NoError(t, err, "idempotent re-run must not revert with ChainAlreadyExists")
		testhelpers.ProcessTimelockProposals(t, *env, output.MCMSTimelockProposals, false)
	})
}

// --------------------------------------------------------------------------------------------------
// Scenario 5: Solana LockRelease + EVM BurnMint cross-chain + Rate Limits (with asymmetric decimals)
// --------------------------------------------------------------------------------------------------

func TestTokenExpansionScenariosSolana(t *testing.T) {
	newChainSel := chainsel.TEST_90000002.Selector
	evmChainSel := chainsel.TEST_90000001.Selector
	solChainSel := chainsel.SOLANA_DEVNET.Selector

	programsPath, ds, err := PreloadSolanaEnvironment(t, solChainSel)
	require.NoError(t, err)

	env, err := environment.New(
		t.Context(),
		environment.WithSolanaContainer(t, []uint64{solChainSel}, programsPath, solanaProgramIDs),
		environment.WithEVMSimulated(t, []uint64{evmChainSel, newChainSel}),
	)
	require.NoError(t, err)
	env.DataStore = ds.Seal()

	evmChain, ok := env.BlockChains.EVMChains()[evmChainSel]
	require.True(t, ok)
	solChain, ok := env.BlockChains.SolanaChains()[solChainSel]
	require.True(t, ok)

	solAdapter := solseqV1_6_0.SolanaAdapter{}
	evmAdapter := evmseqV1_6_0.EVMAdapter{}

	deployRegistry := deployapi.GetRegistry()
	deployRegistry.RegisterDeployer(chainsel.FamilyEVM, deployapi.MCMSVersion, &evmadapters.EVMDeployer{})
	deployRegistry.RegisterDeployer(chainsel.FamilySolana, deployapi.MCMSVersion, &solAdapter)

	mcmsRegistry := changesets.GetRegistry()
	mcmsRegistry.RegisterMCMSReader(chainsel.FamilyEVM, &evmadapters.EVMMCMSReader{})

	// Deploy core contracts
	deployInput := map[uint64]deployapi.ContractDeploymentConfigPerChain{
		newChainSel: NewDefaultDeploymentConfigForEVM(v1_6_0_scenarios),
		evmChainSel: NewDefaultDeploymentConfigForEVM(v1_6_0_scenarios),
		solChainSel: NewDefaultDeploymentConfigForSolana(v1_6_0_scenarios),
	}
	output, err := deployapi.DeployContracts(deployRegistry).Apply(*env, deployapi.ContractDeploymentConfig{Chains: deployInput, MCMS: mcms.Input{}})
	require.NoError(t, err)
	MergeAddresses(t, env, output.DataStore)

	DeployMCMS(t, env, newChainSel, []string{cciputils.CLLQualifier})
	DeployMCMS(t, env, evmChainSel, []string{cciputils.CLLQualifier})
	DeployMCMS(t, env, solChainSel, []string{cciputils.CLLQualifier})
	EVMTransferOwnership(t, env, newChainSel)
	EVMTransferOwnership(t, env, evmChainSel)
	SolanaTransferOwnership(t, env, solChainSel)

	t.Run("Scenario5_SolanaLRToEVMBM", func(t *testing.T) {
		defaultMaxSupply := uint64(1e6)
		defaultPreMint := uint64(1e5)
		evmDecimals := uint8(18)
		svmDecimals := uint8(9)

		defaultRL := tokensapi.RateLimiterConfigFloatInput{
			Capacity:  100,
			Rate:      10,
			IsEnabled: true,
		}

		t.Run("BasicDeploymentAndConnection", func(t *testing.T) {
			const evmTokenSymbol = "S5_EVM_TOK"
			const solTokenSymbol = "S5_SOL_TOK"
			const evmPoolQual = "S5_EVM_POOL"

			output, err = tokensapi.TokenExpansion().Apply(*env, tokensapi.TokenExpansionInput{
				ChainAdapterVersion: v1_6_0_scenarios,
				MCMS:                NewDefaultInputForMCMS("Scenario 5"),
				TokenExpansionInputPerChain: map[uint64]tokensapi.TokenExpansionInputPerChain{
					evmChainSel: {
						TokenPoolVersion: v1_5_1_scenarios,
						DeployTokenInput: &tokensapi.DeployTokenInput{
							Name: "Scenario5 EVM Token", Symbol: evmTokenSymbol, Decimals: evmDecimals,
							Type: bnmERC20ops.ContractType, Supply: &defaultMaxSupply, PreMint: &defaultPreMint,
						},
						DeployTokenPoolInput: &tokensapi.DeployTokenPoolInput{
							TokenPoolQualifier: evmPoolQual,
							PoolType:           cciputils.BurnMintTokenPool.String(),
						},
						TokenTransferConfig: &tokensapi.TokenTransferConfig{
							RemoteChains: map[uint64]tokensapi.RemoteChainConfig[*datastore.AddressRef, datastore.AddressRef]{
								solChainSel: {OutboundRateLimiterConfig: &defaultRL},
							},
						},
					},
					solChainSel: {
						TokenPoolVersion: v1_6_0_scenarios,
						DeployTokenInput: &tokensapi.DeployTokenInput{
							Name: "Scenario5 SOL Token", Symbol: solTokenSymbol, Decimals: svmDecimals,
							Type:                   solanautils.SPLTokens,
							ExternalAdmin:          solana.NewWallet().PublicKey().String(),
							DisableFreezeAuthority: true,
							Senders:                []string{solChain.DeployerKey.PublicKey().String()},
							PreMint:                &defaultPreMint,
						},
						DeployTokenPoolInput: &tokensapi.DeployTokenPoolInput{
							TokenPoolQualifier: "",
							PoolType:           cciputils.LockReleaseTokenPool.String(),
						},
						TokenTransferConfig: &tokensapi.TokenTransferConfig{
							RemoteChains: map[uint64]tokensapi.RemoteChainConfig[*datastore.AddressRef, datastore.AddressRef]{
								evmChainSel: {OutboundRateLimiterConfig: &defaultRL},
							},
						},
					},
				},
			})
			require.NoError(t, err)
			MergeAddresses(t, env, output.DataStore)
			testhelpers.ProcessTimelockProposals(t, *env, output.MCMSTimelockProposals, false)

			// ---- EVM assertions ----
			evmTokAddr := assertTokenExists(t, env, evmChainSel, evmTokenSymbol, "Scenario5 EVM Token", 18)
			evmPoolAddr, err := evmAdapter.FindLatestAddressRef(env.DataStore, datastore.AddressRef{
				ChainSelector: evmChainSel,
				Qualifier:     evmPoolQual,
				Type:          datastore.ContractType(cciputils.BurnMintTokenPool),
			})
			require.NoError(t, err)
			evmPool, err := bnmpool.NewBurnMintTokenPool(evmPoolAddr, evmChain.Client)
			require.NoError(t, err)

			gotTok, err := evmPool.GetToken(&bind.CallOpts{Context: t.Context()})
			require.NoError(t, err)
			require.Equal(t, evmTokAddr, gotTok)

			supportedChains, err := evmPool.GetSupportedChains(&bind.CallOpts{Context: t.Context()})
			require.NoError(t, err)
			require.Contains(t, supportedChains, solChainSel, "EVM pool should support Solana as remote chain")

			assertBMRateLimitsEnabled(t, evmPool, solChainSel)

			// ---- Solana assertions ----
			solTokenRef, err := datastore_utils.FindAndFormatRef(
				env.DataStore,
				datastore.AddressRef{Qualifier: solTokenSymbol},
				solChainSel,
				datastore_utils.FullRef,
			)
			require.NoError(t, err)
			require.NotEmpty(t, solTokenRef.Address, "Solana token should exist in datastore")

			// Find the LockRelease pool program for Solana
			solPoolRef, err := datastore_utils.FindAndFormatRef(env.DataStore, datastore.AddressRef{
				ChainSelector: solChainSel,
				Type:          datastore.ContractType(cciputils.LockReleaseTokenPool),
				Version:       v1_6_0_scenarios,
			}, solChainSel, datastore_utils.FullRef)
			require.NoError(t, err)
			solPoolProgramID := solana.MustPublicKeyFromBase58(solPoolRef.Address)
			solTokenMint := solana.MustPublicKeyFromBase58(solTokenRef.Address)

			// Verify Solana pool PDA state
			tokenPoolStatePDA, _ := tokens.TokenPoolConfigAddress(solTokenMint, solPoolProgramID)
			var solPoolState lockrelease_token_pool.State
			err = solChain.GetAccountDataBorshInto(t.Context(), tokenPoolStatePDA, &solPoolState)
			require.NoError(t, err, "Solana LockRelease pool state PDA should be initialized")
			require.Equal(t, solTokenMint, solPoolState.Config.Mint,
				fmt.Sprintf("Solana pool mint should match token %s", solTokenRef.Address))

			timelockSigner := solanautils.GetTimelockSignerPDA(
				env.DataStore.Addresses().Filter(),
				solChainSel,
				cciputils.CLLQualifier,
			)
			require.Equal(t, timelockSigner, solPoolState.Config.RateLimitAdmin,
				"empty DeployTokenPoolInput.RateLimitAdmin should set RL admin to MCMS timelock signer PDA on Solana")

			// Verify Solana router address is set
			routerAddr, err := solAdapter.GetRouterAddress(env.DataStore, solChainSel)
			require.NoError(t, err)
			require.NotEmpty(t, routerAddr, "Solana router should be deployed")

			// Regression: SolanaAdapter.GetOnchainInboundRateLimit used to decode the on-chain
			// ChainConfig PDA into a bare base_token_pool.BaseChain, ignoring the 8-byte Anchor
			// discriminator. The decode would silently fail and the function would return a
			// zero RateLimiterConfig, which in OutboundOnly mode then overwrote the bucket's
			// InboundRateLimiterConfig with zeros — clobbering a previously-enabled inbound
			// rate limit on chain. Solana's inbound is already enabled and non-zero from the
			// symmetric OutboundRateLimiterConfig on both sides in TokenExpansion().Apply above
			// (BasicDeploymentAndConnection). This subtest runs an OutboundOnly apply on Solana
			// and verifies that inbound is preserved by pass-through.
			t.Run("SetTokenPoolRateLimits_OutboundOnly_PreservesSolanaInbound", func(t *testing.T) {
				chainCfgPDA, _, err := tokens.TokenPoolChainConfigPDA(evmChainSel, solTokenMint, solPoolProgramID)
				require.NoError(t, err)

				// Snapshot Solana's current rate limits — both directions should be enabled with
				// non-zero values from the TokenExpansion lane setup above.
				var preCfg lockrelease_token_pool.ChainConfig
				require.NoError(t, solChain.GetAccountDataBorshInto(t.Context(), chainCfgPDA, &preCfg))
				require.True(t, preCfg.Base.InboundRateLimit.Cfg.Enabled, "seed: Solana inbound from EVM should be enabled before OutboundOnly apply")
				require.NotZero(t, preCfg.Base.InboundRateLimit.Cfg.Capacity, "seed: Solana inbound capacity should be non-zero before OutboundOnly apply")
				require.NotZero(t, preCfg.Base.InboundRateLimit.Cfg.Rate, "seed: Solana inbound rate should be non-zero before OutboundOnly apply")

				// Snapshot EVM rate limits too so we can assert they're untouched without
				// depending on which specific values the TokenExpansion seed left on-chain.
				preOutboundEVM, err := evmPool.GetCurrentOutboundRateLimiterState(&bind.CallOpts{Context: t.Context()}, solChainSel)
				require.NoError(t, err)
				preInboundEVM, err := evmPool.GetCurrentInboundRateLimiterState(&bind.CallOpts{Context: t.Context()}, solChainSel)
				require.NoError(t, err)

				// Pick a new outbound that's safely below the counterpart's existing inbound
				// headroom: counterpart inbound (in svmDecimals) >= 1.10 * new_outbound. Use a
				// small value relative to anything the TokenExpansion seed could have set.
				newSVMOutbound := tokensapi.RemoteOutbounds{
					OutboundOnly: true,
					Outbounds: []tokensapi.RateLimitConfig{{
						RateLimit: tokensapi.RateLimiterConfigFloatInput{Capacity: 50, Rate: 5, IsEnabled: true},
					}},
				}

				tprlOut, err := tokensapi.SetTokenPoolRateLimits().Apply(*env, tokensapi.TPRLInput{
					MCMS: NewDefaultInputForMCMS("Scenario 5 TPRL OutboundOnly"),
					Configs: map[uint64]tokensapi.TPRLConfig{
						solChainSel: {
							ChainAdapterVersion: v1_6_0_scenarios,
							TokenPoolRef:        datastore.AddressRef{Address: solPoolProgramID.String()},
							TokenRef:            datastore.AddressRef{Address: solTokenMint.String()},
							RemoteOutbounds: map[uint64]tokensapi.RemoteOutbounds{
								evmChainSel: newSVMOutbound,
							},
						},
						// Counterpart EVM config is refs-only — no RemoteOutbounds entry — which is
						// exactly the shape OutboundOnly is designed for.
						evmChainSel: {
							ChainAdapterVersion: v1_6_0_scenarios,
							TokenPoolRef:        datastore.AddressRef{Address: evmPoolAddr.Hex()},
							TokenRef:            datastore.AddressRef{Address: evmTokAddr.Hex()},
						},
					},
				})
				require.NoError(t, err)
				testhelpers.ProcessTimelockProposals(t, *env, tprlOut.MCMSTimelockProposals, false)

				// Re-read Solana's chain config.
				var postCfg lockrelease_token_pool.ChainConfig
				require.NoError(t, solChain.GetAccountDataBorshInto(t.Context(), chainCfgPDA, &postCfg))

				// Solana inbound must be preserved bit-for-bit (the OutboundOnly pass-through).
				// Before the fix, GetOnchainInboundRateLimit silently returned a zero
				// RateLimiterConfig, the bucket's InboundRateLimiterConfig was overwritten
				// with zeros, and these three assertions failed.
				require.Equal(t, preCfg.Base.InboundRateLimit.Cfg.Enabled, postCfg.Base.InboundRateLimit.Cfg.Enabled, "Solana inbound IsEnabled must be unchanged by OutboundOnly apply")
				require.Equal(t, preCfg.Base.InboundRateLimit.Cfg.Capacity, postCfg.Base.InboundRateLimit.Cfg.Capacity, "Solana inbound capacity must be unchanged by OutboundOnly apply")
				require.Equal(t, preCfg.Base.InboundRateLimit.Cfg.Rate, postCfg.Base.InboundRateLimit.Cfg.Rate, "Solana inbound rate must be unchanged by OutboundOnly apply")

				// Solana outbound was rewritten to the new value.
				expOutCap := tokensapi.ScaleFloatToBigInt(newSVMOutbound.Outbounds[0].RateLimit.Capacity, int(svmDecimals), 0)
				expOutRate := tokensapi.ScaleFloatToBigInt(newSVMOutbound.Outbounds[0].RateLimit.Rate, int(svmDecimals), 0)
				require.True(t, postCfg.Base.OutboundRateLimit.Cfg.Enabled)
				require.Equal(t, expOutCap.Uint64(), postCfg.Base.OutboundRateLimit.Cfg.Capacity, "Solana outbound capacity after OutboundOnly apply")
				require.Equal(t, expOutRate.Uint64(), postCfg.Base.OutboundRateLimit.Cfg.Rate, "Solana outbound rate after OutboundOnly apply")

				// EVM pool was not touched by this OutboundOnly apply (no RemoteOutbounds entry
				// for solChainSel was provided on the EVM side).
				postOutboundEVM, err := evmPool.GetCurrentOutboundRateLimiterState(&bind.CallOpts{Context: t.Context()}, solChainSel)
				require.NoError(t, err)
				postInboundEVM, err := evmPool.GetCurrentInboundRateLimiterState(&bind.CallOpts{Context: t.Context()}, solChainSel)
				require.NoError(t, err)
				require.Equal(t, preOutboundEVM.IsEnabled, postOutboundEVM.IsEnabled, "EVM outbound IsEnabled should be unchanged after OutboundOnly apply on Solana")
				require.Zero(t, preOutboundEVM.Capacity.Cmp(postOutboundEVM.Capacity), "EVM outbound capacity should be unchanged after OutboundOnly apply on Solana")
				require.Zero(t, preOutboundEVM.Rate.Cmp(postOutboundEVM.Rate), "EVM outbound rate should be unchanged after OutboundOnly apply on Solana")
				require.Equal(t, preInboundEVM.IsEnabled, postInboundEVM.IsEnabled, "EVM inbound IsEnabled should be unchanged after OutboundOnly apply on Solana")
				require.Zero(t, preInboundEVM.Capacity.Cmp(postInboundEVM.Capacity), "EVM inbound capacity should be unchanged after OutboundOnly apply on Solana")
				require.Zero(t, preInboundEVM.Rate.Cmp(postInboundEVM.Rate), "EVM inbound rate should be unchanged after OutboundOnly apply on Solana")
			})
		})

		// Test address ref inference
		t.Run("TokenRefResolver", func(t *testing.T) {
			// Helper vars
			refReader := tokensapi.GetTokenAdapterRegistry()
			svmSymbol := "S5_TR_SVM"
			evmSymbol := "S5_TR_EVM"

			// Make sure tokens aren't in the datastore yet
			refs := env.DataStore.Addresses().Filter(datastore.AddressRefByChainSelector(solChainSel), datastore.AddressRefByQualifier(svmSymbol))
			require.Empty(t, refs, "SVM token symbol should not exist in env.DataStore before test")
			refs = env.DataStore.Addresses().Filter(datastore.AddressRefByChainSelector(evmChainSel), datastore.AddressRefByQualifier(evmSymbol))
			require.Empty(t, refs, "EVM token symbol should not exist in env.DataStore before test")

			// Deploy some tokens and pools
			acceptLiquidity := true
			teOut, err := tokensapi.TokenExpansion().Apply(*env, tokensapi.TokenExpansionInput{
				ChainAdapterVersion: v1_6_0_scenarios,
				MCMS:                NewDefaultInputForMCMS("Scenario5 AddressRefInference"),
				TokenExpansionInputPerChain: map[uint64]tokensapi.TokenExpansionInputPerChain{
					evmChainSel: {
						TokenPoolVersion: v1_5_1_scenarios,
						DeployTokenInput: &tokensapi.DeployTokenInput{
							Name: "Scenario5 try-resolve EVM token", Symbol: evmSymbol, Decimals: 18,
							Type: bnmERC20ops.ContractType, Supply: &defaultMaxSupply, PreMint: &defaultPreMint,
						},
						DeployTokenPoolInput: &tokensapi.DeployTokenPoolInput{
							TokenPoolQualifier: "", // use sensible default
							PoolType:           cciputils.LockReleaseTokenPool.String(),
							AcceptLiquidity:    &acceptLiquidity,
						},
					},
					solChainSel: {
						TokenPoolVersion: v1_6_0_scenarios,
						DeployTokenInput: &tokensapi.DeployTokenInput{
							Name: "Scenario5 try-resolve SVM token", Symbol: svmSymbol, Decimals: svmDecimals,
							Type: solanautils.SPLTokens, ExternalAdmin: solana.NewWallet().PublicKey().String(),
							Senders: []string{solChain.DeployerKey.PublicKey().String()},
						},
						DeployTokenPoolInput: &tokensapi.DeployTokenPoolInput{
							TokenPoolQualifier: "",
							PoolType:           cciputils.BurnMintTokenPool.String(),
						},
					},
				},
			})
			require.NoError(t, err)
			testhelpers.ProcessTimelockProposals(t, *env, teOut.MCMSTimelockProposals, false)

			// Get SVM token address from changeset output
			svmTokRefs := teOut.DataStore.Addresses().Filter(datastore.AddressRefByChainSelector(solChainSel), datastore.AddressRefByQualifier(svmSymbol))
			require.Len(t, svmTokRefs, 1, "token expansion output should record the new Solana token")
			svmTokAddr := svmTokRefs[0].Address

			// Get SVM pool address from environment datastore
			svmPoolRefs := env.DataStore.Addresses().Filter(datastore.AddressRefByChainSelector(solChainSel), datastore.AddressRefByType(datastore.ContractType(cciputils.BurnMintTokenPool)))
			require.Len(t, svmPoolRefs, 1, "preloaded BurnMint token pool program should exist in env.DataStore on Solana (TokenExpansion registers tokens under existing programs, not new pool programs)")
			svmPoolAddr := svmPoolRefs[0].Address

			// Get EVM token address from changeset output
			evmTokRefs := teOut.DataStore.Addresses().Filter(datastore.AddressRefByChainSelector(evmChainSel), datastore.AddressRefByQualifier(evmSymbol))
			require.Len(t, evmTokRefs, 1, "token expansion output should record the new EVM token")
			evmTokAddr := evmTokRefs[0].Address

			// Get EVM pool address from changeset output
			evmPoolRefs := teOut.DataStore.Addresses().Filter(datastore.AddressRefByChainSelector(evmChainSel), datastore.AddressRefByQualifier(evmTokAddr))
			require.Len(t, evmPoolRefs, 1, "token expansion output should record the new EVM pool")
			evmPoolAddr := evmPoolRefs[0].Address

			// We did not call MergeAddresses on teOut, so the environment datastore should NOT have the deployed tokens yet
			refs = env.DataStore.Addresses().Filter(datastore.AddressRefByChainSelector(solChainSel), datastore.AddressRefByQualifier(svmSymbol), datastore.AddressRefByAddress(svmTokAddr))
			require.Empty(t, refs, "new SVM mint must not be in env.DataStore before ResolveTokenRef")
			refs = env.DataStore.Addresses().Filter(datastore.AddressRefByChainSelector(evmChainSel), datastore.AddressRefByQualifier(evmSymbol), datastore.AddressRefByAddress(evmTokAddr))
			require.Empty(t, refs, "new EVM token address must not be in env.DataStore before ResolveTokenRef")

			// Get the Solana pool PDA and verify it exists on chain
			svmPoolProgID := solana.MustPublicKeyFromBase58(svmPoolAddr)
			svmMintPubKey := solana.MustPublicKeyFromBase58(svmTokAddr)
			svmPoolCfgPDA, _ := tokens.TokenPoolConfigAddress(svmMintPubKey, svmPoolProgID)
			var solPoolStateBnM burnmint_token_pool.State
			err = solChain.GetAccountDataBorshInto(t.Context(), svmPoolCfgPDA, &solPoolStateBnM)
			require.NoError(t, err, "Solana BnM pool state PDA should be initialized for the new pool")

			// Check that we can resolve the new pool and token refs even though they are not in the environment datastore
			t.Run("ResolveRef", func(t *testing.T) {
				// Token expansion for Solana now supports either the pool program ID or the pool PDA:
				// (1) if the user provides a pool PDA, then the initial DS look up will fail -> code will check if it's a PDA and resolve its program ID from the chain -> datastore lookup is re-attempted
				// (2) if the user provides a prog ID, then the ref must be in the datastore (no onchain resolution is possible)
				resolvedSvmPool, err := tokensapi.ResolveTokenPoolRef(*env, refReader, solChainSel, datastore.AddressRef{Address: svmPoolCfgPDA.String()})
				require.NoError(t, err)
				require.Equal(t, datastore.ContractType(cciputils.BurnMintTokenPool), resolvedSvmPool.Type)
				require.Equal(t, svmPoolProgID.String(), resolvedSvmPool.Address)
				require.Equal(t, "", resolvedSvmPool.Qualifier)
				require.True(t, v1_6_0_scenarios.Equal(resolvedSvmPool.Version))

				// Since the SVM token does not exist in the environment datastore yet, ResolveTokenRef will resolve it from on chain data
				resolvedSvmTok, err := tokensapi.ResolveTokenRef(*env, refReader, solChainSel, datastore.AddressRef{Address: svmTokAddr})
				require.NoError(t, err)
				require.Equal(t, svmTokAddr, resolvedSvmTok.Address)
				require.Equal(t, fmt.Sprintf("%s-%s", svmTokAddr, solanautils.SPLTokens), resolvedSvmTok.Qualifier)
				require.Equal(t, datastore.ContractType(solanautils.SPLTokens), resolvedSvmTok.Type)
				require.True(t, v1_6_0_scenarios.Equal(resolvedSvmTok.Version))

				// Since the EVM pool does not exist in the environment datastore yet, ResolveTokenPoolRef will resolve it from on chain data
				resolvedEvmPool, err := tokensapi.ResolveTokenPoolRef(*env, refReader, evmChainSel, datastore.AddressRef{Address: evmPoolAddr})
				require.NoError(t, err)
				require.Equal(t, datastore.ContractType(cciputils.LockReleaseTokenPool), resolvedEvmPool.Type)
				require.Equal(t, evmPoolAddr, resolvedEvmPool.Address)
				require.Equal(t, evmTokAddr, resolvedEvmPool.Qualifier)
				require.True(t, v1_5_1_scenarios.Equal(resolvedEvmPool.Version))

				// Since the EVM token does not exist in the environment datastore yet, ResolveTokenRef will resolve it from on chain data
				resolvedEvmTok, err := tokensapi.ResolveTokenRef(*env, refReader, evmChainSel, datastore.AddressRef{Address: evmTokAddr})
				require.NoError(t, err)
				require.Equal(t, common.HexToAddress(evmTokAddr).Hex(), resolvedEvmTok.Address)
				require.Equal(t, evmSymbol, resolvedEvmTok.Qualifier)
				require.Equal(t, datastore.ContractType(erc20ops.ContractType), resolvedEvmTok.Type)
				require.True(t, cciputils.Version_1_0_0.Equal(resolvedEvmTok.Version))
			})

			// The deployed tokens are not in the datastore yet so calling TokenExpansion again would normally result in an error.
			// However, the new TokenRefResolver interface allows us to source tokens and pools from the chain so now we no longer
			// need to worry about the datastore containing these token/pool refs - the tooling is smart enough to infer them.
			t.Run("TokenExpansionWithInferredRefs", func(t *testing.T) {
				teOut, err = tokensapi.TokenExpansion().Apply(*env, tokensapi.TokenExpansionInput{
					ChainAdapterVersion: v1_6_0_scenarios,
					MCMS:                NewDefaultInputForMCMS("Scenario5 AddressRefInference"),
					TokenExpansionInputPerChain: map[uint64]tokensapi.TokenExpansionInputPerChain{
						newChainSel: {
							TokenPoolVersion: cciputils.Version_1_6_1,
							DeployTokenInput: &tokensapi.DeployTokenInput{
								Name: "Scenario5 try-resolve EVM token", Symbol: evmSymbol, Decimals: 18,
								Type: bnmERC20ops.ContractType, Supply: &defaultMaxSupply, PreMint: &defaultPreMint,
							},
							DeployTokenPoolInput: &tokensapi.DeployTokenPoolInput{
								TokenPoolQualifier: "", // use sensible default
								PoolType:           cciputils.BurnWithFromMintTokenPool.String(),
							},
							TokenTransferConfig: &tokensapi.TokenTransferConfig{
								RemoteChains: map[uint64]tokensapi.RemoteChainConfig[*datastore.AddressRef, datastore.AddressRef]{
									evmChainSel: {OutboundRateLimiterConfig: &defaultRL},
									solChainSel: {OutboundRateLimiterConfig: &defaultRL},
								},
							},
						},
						evmChainSel: {
							// TokenRef is not needed - DeriveTokenAddress can resolve the whole ref from the chain
							TokenTransferConfig: &tokensapi.TokenTransferConfig{
								TokenPoolRef: datastore.AddressRef{Address: evmPoolAddr},
								RemoteChains: map[uint64]tokensapi.RemoteChainConfig[*datastore.AddressRef, datastore.AddressRef]{
									newChainSel: {OutboundRateLimiterConfig: &defaultRL},
									solChainSel: {OutboundRateLimiterConfig: &defaultRL},
								},
							},
						},
						solChainSel: {
							// TokenRef is not needed - DeriveTokenAddress can resolve the whole ref from the chain since we're using the pool PDA
							TokenTransferConfig: &tokensapi.TokenTransferConfig{
								TokenPoolRef: datastore.AddressRef{Address: svmPoolCfgPDA.String()},
								RemoteChains: map[uint64]tokensapi.RemoteChainConfig[*datastore.AddressRef, datastore.AddressRef]{
									newChainSel: {OutboundRateLimiterConfig: &defaultRL},
									evmChainSel: {OutboundRateLimiterConfig: &defaultRL},
								},
							},
						},
					},
				})
				require.NoError(t, err)
				testhelpers.ProcessTimelockProposals(t, *env, teOut.MCMSTimelockProposals, false)
			})

			// Next, we try updating the rate limits in both directions without providing the token refs (the ref resolver should be able to infer them)
			t.Run("TPRLWithInferredRefs", func(t *testing.T) {
				evmTowardSVM := tokensapi.RemoteOutbounds{RateLimit: &tokensapi.RateLimiterConfigFloatInput{Capacity: 111, Rate: 11, IsEnabled: true}}
				svmTowardEVM := tokensapi.RemoteOutbounds{RateLimit: &tokensapi.RateLimiterConfigFloatInput{Capacity: 222, Rate: 22, IsEnabled: true}}
				tprlOut, err := tokensapi.SetTokenPoolRateLimits().Apply(*env, tokensapi.TPRLInput{
					MCMS: NewDefaultInputForMCMS("Scenario 5 TPRL"),
					Configs: map[uint64]tokensapi.TPRLConfig{
						evmChainSel: {
							TokenPoolRef: datastore.AddressRef{Address: evmPoolAddr},
							RemoteOutbounds: map[uint64]tokensapi.RemoteOutbounds{
								solChainSel: evmTowardSVM,
							},
						},
						solChainSel: {
							TokenPoolRef: datastore.AddressRef{Address: svmPoolCfgPDA.String()},
							RemoteOutbounds: map[uint64]tokensapi.RemoteOutbounds{
								evmChainSel: svmTowardEVM,
							},
						},
					},
				})
				require.NoError(t, err)
				testhelpers.ProcessTimelockProposals(t, *env, tprlOut.MCMSTimelockProposals, false)

				// Read EVM rate limits from the chain
				evmPool, err := lrpool.NewLockReleaseTokenPool(common.HexToAddress(evmPoolAddr), evmChain.Client)
				require.NoError(t, err)
				outboundEVM, err := evmPool.GetCurrentOutboundRateLimiterState(&bind.CallOpts{Context: t.Context()}, solChainSel)
				require.NoError(t, err)
				inboundEVM, err := evmPool.GetCurrentInboundRateLimiterState(&bind.CallOpts{Context: t.Context()}, solChainSel)
				require.NoError(t, err)

				// Verify EVM rate limit state matches expected values (with correct scaling for decimals). The
				// inbound on legacy EVM pools (<v1.6.1 non-external-minter) scales using remote token decimals
				// (see tokensapi.GenerateTPRLConfigs); counterpart outbound is Solana -> use svmDecimals.
				expInboundCapEVM := tokensapi.ScaleFloatToBigInt(svmTowardEVM.RateLimit.Capacity, int(svmDecimals), 0.10)
				expInboundRateEVM := tokensapi.ScaleFloatToBigInt(svmTowardEVM.RateLimit.Rate, int(svmDecimals), 0.10)
				expOutboundCapEVM := tokensapi.ScaleFloatToBigInt(evmTowardSVM.RateLimit.Capacity, int(evmDecimals), 0)
				expOutboundRateEVM := tokensapi.ScaleFloatToBigInt(evmTowardSVM.RateLimit.Rate, int(evmDecimals), 0)
				require.Zero(t, expOutboundCapEVM.Cmp(outboundEVM.Capacity), "EVM outbound capacity toward Solana should match TPRL input after scaling (want %s, got %s)", expOutboundCapEVM.String(), outboundEVM.Capacity.String())
				require.Zero(t, expOutboundRateEVM.Cmp(outboundEVM.Rate), "EVM outbound rate toward Solana should match TPRL input after scaling (want %s, got %s)", expOutboundRateEVM.String(), outboundEVM.Rate.String())
				require.Zero(t, expInboundCapEVM.Cmp(inboundEVM.Capacity), "EVM inbound capacity from Solana should match counterpart outbound TPRL input + inbound scaling (want %s, got %s)", expInboundCapEVM.String(), inboundEVM.Capacity.String())
				require.Zero(t, expInboundRateEVM.Cmp(inboundEVM.Rate), "EVM inbound rate from Solana should match counterpart outbound TPRL input + inbound scaling (want %s, got %s)", expInboundRateEVM.String(), inboundEVM.Rate.String())
				require.True(t, outboundEVM.IsEnabled)
				require.True(t, inboundEVM.IsEnabled)

				// Read Solana rate limits from the chain
				chainCfgPDA, _, err := tokens.TokenPoolChainConfigPDA(evmChainSel, svmMintPubKey, svmPoolProgID)
				require.NoError(t, err)
				var chainCfg lockrelease_token_pool.ChainConfig
				require.NoError(t, solChain.GetAccountDataBorshInto(t.Context(), chainCfgPDA, &chainCfg))
				outboundCapSVM := big.NewInt(0).SetUint64(chainCfg.Base.OutboundRateLimit.Cfg.Capacity)
				outboundRateSVM := big.NewInt(0).SetUint64(chainCfg.Base.OutboundRateLimit.Cfg.Rate)
				inboundCapSVM := big.NewInt(0).SetUint64(chainCfg.Base.InboundRateLimit.Cfg.Capacity)
				inboundRateSVM := big.NewInt(0).SetUint64(chainCfg.Base.InboundRateLimit.Cfg.Rate)

				// Verify Solana rate limit state matches expected values (with correct scaling for decimals)
				expInboundCapSVM := tokensapi.ScaleFloatToBigInt(evmTowardSVM.RateLimit.Capacity, int(svmDecimals), 0.10)
				expInboundRateSVM := tokensapi.ScaleFloatToBigInt(evmTowardSVM.RateLimit.Rate, int(svmDecimals), 0.10)
				expOutboundCapSVM := tokensapi.ScaleFloatToBigInt(svmTowardEVM.RateLimit.Capacity, int(svmDecimals), 0)
				expOutboundRateSVM := tokensapi.ScaleFloatToBigInt(svmTowardEVM.RateLimit.Rate, int(svmDecimals), 0)
				require.Zero(t, expOutboundCapSVM.Cmp(outboundCapSVM), "Solana outbound capacity toward EVM should match TPRL input after scaling (want %s, got %s)", expOutboundCapSVM.String(), outboundCapSVM.String())
				require.Zero(t, expOutboundRateSVM.Cmp(outboundRateSVM), "Solana outbound rate toward EVM should match TPRL input after scaling (want %s, got %s)", expOutboundRateSVM.String(), outboundRateSVM.String())
				require.Zero(t, expInboundCapSVM.Cmp(inboundCapSVM), "Solana inbound capacity from EVM should match counterpart outbound TPRL input + inbound scaling (want %s, got %s)", expInboundCapSVM.String(), inboundCapSVM.String())
				require.Zero(t, expInboundRateSVM.Cmp(inboundRateSVM), "Solana inbound rate from EVM should match counterpart outbound TPRL input + inbound scaling (want %s, got %s)", expInboundRateSVM.String(), inboundRateSVM.String())
				require.True(t, chainCfg.Base.OutboundRateLimit.Cfg.Enabled)
				require.True(t, chainCfg.Base.InboundRateLimit.Cfg.Enabled)
			})
		})
	})
}

// ---------------------------------------------------------------------------
// Scenario 6: MCMS batches do NOT execute immediately, so the changeset should
// not perform validation assuming on-chain state changes have already occurred
// within the same batch until the batch is executed. This is relevant to token
// expansion flows where multiple changesets in the same batch update the Token
// Admin Registry (TAR) state for a mint - such as a RegisterTokenAdminRegistry
// that sets a non-timelock pending admin followed by ConfigureTokenForTransfer
// that overrides pending to timelock, and then Accepts that finalised timelock
// admin slot *all within the same batch*. If the Configure or Accept changeset
// incorrectly assumes the Register's pending admin change is already on-chain,
// it may read stale state and fail validation (e.g., "pending admin  ...  does
// not match timelock signer ...") even though the intended final state after a
// batch execution is correct. The scenario below is based on an actual mainnet
// request that wrongly failed due to this issue.
// ---------------------------------------------------------------------------
func TestSolanaCrossFamilyTokenExpansion_thirdPartyPendingTAR(t *testing.T) {
	// Helper consts
	const (
		RateLimitAdminEVM = "0x8C245711032b426945D04Df60c28DF04d30c15eB"
		RateLimitAdminSVM = "9o4rGhjgughgQxYXibqQfBwSCtXbJ3GJtzNVFhYCcRYg"
		TestTokenSymbol   = "TEST"
	)

	// Helper vars
	evmChainSel := chainsel.TEST_90000001.Selector
	solChainSel := chainsel.SOLANA_DEVNET.Selector
	realWorldRL := tokensapi.RateLimiterConfigFloatInput{
		IsEnabled: true,
		Capacity:  7_000_000,
		Rate:      1944.0,
	}

	// Setup Solana environment
	programsPath, ds, err := PreloadSolanaEnvironment(t, solChainSel)
	require.NoError(t, err)

	// Setup test env
	env, err := environment.New(
		t.Context(),
		environment.WithSolanaContainer(t, []uint64{solChainSel}, programsPath, solanaProgramIDs),
		environment.WithEVMSimulated(t, []uint64{evmChainSel}),
	)
	require.NoError(t, err)
	env.DataStore = ds.Seal()

	// Ensure the chains exist
	solChain, ok := env.BlockChains.SolanaChains()[solChainSel]
	require.True(t, ok)
	_, ok = env.BlockChains.EVMChains()[evmChainSel]
	require.True(t, ok)

	// Define adapters
	solAdapter := solseqV1_6_0.SolanaAdapter{}
	evmDeployer := evmadapters.EVMDeployer{}
	evmReader := evmadapters.EVMMCMSReader{}

	// Setup deployment registry
	deployRegistry := deployapi.GetRegistry()
	deployRegistry.RegisterDeployer(chainsel.FamilySolana, deployapi.MCMSVersion, &solAdapter)
	deployRegistry.RegisterDeployer(chainsel.FamilyEVM, deployapi.MCMSVersion, &evmDeployer)

	// Setup MCMS registry
	mcmsRegistry := changesets.GetRegistry()
	mcmsRegistry.RegisterMCMSReader(chainsel.FamilySolana, &solAdapter)
	mcmsRegistry.RegisterMCMSReader(chainsel.FamilyEVM, &evmReader)

	// Deploy chain contracts
	deployInput := deployapi.ContractDeploymentConfig{
		Chains: map[uint64]deployapi.ContractDeploymentConfigPerChain{
			solChainSel: NewDefaultDeploymentConfigForSolana(v1_6_0_scenarios),
			evmChainSel: NewDefaultDeploymentConfigForEVM(v1_6_0_scenarios),
		},
		MCMS: mcms.Input{},
	}
	deployOut, err := deployapi.DeployContracts(deployRegistry).Apply(*env, deployInput)
	require.NoError(t, err)
	MergeAddresses(t, env, deployOut.DataStore)

	// Setup MCMS
	DeployMCMS(t, env, evmChainSel, []string{cciputils.CLLQualifier})
	DeployMCMS(t, env, solChainSel, []string{cciputils.CLLQualifier})
	EVMTransferOwnership(t, env, evmChainSel)
	SolanaTransferOwnership(t, env, solChainSel)

	// Setup an existing token on Solana
	env.OperationsBundle = operations.NewBundle(env.GetContext, env.Logger, operations.NewMemoryReporter())
	expandTE1, err := tokensapi.TokenExpansion().Apply(*env, tokensapi.TokenExpansionInput{
		ChainAdapterVersion: v1_6_0_scenarios,
		MCMS:                NewDefaultInputForMCMS("deploy Solana token for test"),
		TokenExpansionInputPerChain: map[uint64]tokensapi.TokenExpansionInputPerChain{
			solChainSel: {
				TokenPoolVersion: v1_6_0_scenarios,
				DeployTokenInput: &tokensapi.DeployTokenInput{
					DisableFreezeAuthority: false,
					Decimals:               9,
					Symbol:                 TestTokenSymbol,
					Type:                   solanautils.SPLTokens,
					Name:                   "",
					Senders: []string{
						solChain.DeployerKey.PublicKey().String(),
					},
				},
			},
		},
	})
	require.NoError(t, err)
	MergeAddresses(t, env, expandTE1.DataStore)
	testhelpers.ProcessTimelockProposals(t, *env, expandTE1.MCMSTimelockProposals, false)

	// Fetch the deployed Solana token from the datastore
	solMintPK, err := datastore_utils.FindAndFormatRef(
		env.DataStore,
		datastore.AddressRef{Qualifier: TestTokenSymbol},
		solChainSel,
		solanautils.ToAddress,
	)
	require.NoError(t, err)

	// Fetch the router address for Solana
	routerAddr, err := solAdapter.GetRouterAddress(env.DataStore, solChainSel)
	require.NoError(t, err)
	routerPK := solana.PublicKeyFromBytes(routerAddr)

	// Get the TAR PDA for the deployed token mint and router
	tarPDA, _, err := state.FindTokenAdminRegistryPDA(solMintPK, routerPK)
	require.NoError(t, err)

	// Helper for fetching and decoding the TAR on chain state
	fetchTAR := func(show bool) ccip_common.TokenAdminRegistry {
		t.Helper()
		var tarState ccip_common.TokenAdminRegistry
		decodeErr := solChain.GetAccountDataBorshInto(t.Context(), tarPDA, &tarState)
		require.NoError(t, decodeErr)
		if show {
			t.Logf(
				"TAR: Version=%d Administrator=%s PendingAdministrator=%s Mint=%s LookupTable=%s SupportsAutoDerivation=%v",
				tarState.Version,
				tarState.Administrator,
				tarState.PendingAdministrator,
				tarState.Mint,
				tarState.LookupTable,
				tarState.SupportsAutoDerivation,
			)
		}
		return tarState
	}

	// Configure mint authority on newly deployed Solana token
	externalMintAuth, err := solana.NewRandomPrivateKey()
	require.NoError(t, err)
	err = solanautils.FundFromDeployerKey(solChain, []solana.PublicKey{externalMintAuth.PublicKey()}, 10)
	require.NoError(t, err)
	setAuthIx, err := tokens.SetTokenMintAuthority(solana.TokenProgramID, externalMintAuth.PublicKey(), solMintPK, solChain.DeployerKey.PublicKey())
	require.NoError(t, err)
	err = solChain.Confirm([]solana.Instruction{setAuthIx})
	require.NoError(t, err)

	// Build an MCMS batch that sets the pending administrator for the token on TAR to a 3rd party (not the timelock signer)
	customer := solana.NewWallet().PublicKey()
	reg1Report, err := operations.ExecuteOperation(
		env.OperationsBundle,
		routerops.RegisterTokenAdminRegistry,
		solChain,
		routerops.TokenAdminRegistryParams{
			ExistingAddresses: env.DataStore.Addresses().Filter(),
			TokenMint:         solMintPK,
			Router:            routerPK,
			Admin:             customer,
		},
		operations.WithForceExecute[routerops.TokenAdminRegistryParams, solchain.Chain](),
	)
	require.NoError(t, err, "register should return batch ops when mint authority != deployer")
	require.NotEmpty(t, reg1Report.Output.BatchOps, "first Register should require MCMS")

	// Set the 3rd party pending admin for the token on TAR via timelock
	mcmsProposalInput := NewDefaultInputForMCMS("propose third party as pending TAR admin")
	require.NoError(t, mcmsProposalInput.Validate())
	reg1ProposalOut, err := changesets.NewOutputBuilder(*env, mcmsRegistry).WithBatchOps(reg1Report.Output.BatchOps).Build(mcmsProposalInput)
	require.NoError(t, err)
	require.Len(t, reg1ProposalOut.MCMSTimelockProposals, 1)
	testhelpers.ProcessTimelockProposals(t, *env, reg1ProposalOut.MCMSTimelockProposals, false)

	// Assert that the pending admin is set to the customer after the MCMS batch executes
	tarState := fetchTAR(true)
	require.Equal(t, customer, tarState.PendingAdministrator)

	// Mimic the real-world token expansion scenario:
	//   On Solana: hook up the existing 3rd party token to the CLL self-service LnR pool
	//   On EVM: deploy a new token and pool
	//   Connect them together
	env.OperationsBundle = operations.NewBundle(env.GetContext, env.Logger, operations.NewMemoryReporter())
	expandTE2, err := tokensapi.TokenExpansion().Apply(*env, tokensapi.TokenExpansionInput{
		ChainAdapterVersion: v1_6_0_scenarios,
		MCMS:                NewDefaultInputForMCMS("cross-family TE2 deploy pools + transfer config"),
		TokenExpansionInputPerChain: map[uint64]tokensapi.TokenExpansionInputPerChain{
			solChainSel: {
				TokenPoolVersion:      v1_6_0_scenarios,
				SkipOwnershipTransfer: false,
				DeployTokenPoolInput: &tokensapi.DeployTokenPoolInput{
					TokenPoolQualifier: "",
					RateLimitAdmin:     RateLimitAdminSVM,
					TokenRef:           &datastore.AddressRef{Address: solMintPK.String()},
					PoolType:           cciputils.LockReleaseTokenPool.String(),
				},
				TokenTransferConfig: &tokensapi.TokenTransferConfig{
					TokenRef:     datastore.AddressRef{}, // inferred
					TokenPoolRef: datastore.AddressRef{}, // inferred
					RemoteChains: map[uint64]tokensapi.RemoteChainConfig[*datastore.AddressRef, datastore.AddressRef]{
						evmChainSel: {OutboundRateLimiterConfig: &realWorldRL},
					},
				},
			},
			evmChainSel: {
				SkipOwnershipTransfer: false,
				TokenPoolVersion:      cciputils.Version_1_6_1,
				DeployTokenInput: &tokensapi.DeployTokenInput{
					ExternalAdmin: RateLimitAdminEVM,
					CCIPAdmin:     "",
					Currency:      "",
					Symbol:        TestTokenSymbol,
					Name:          "Test Token",
					Type:          bnmERC20ops.ContractType,
					Decimals:      6,
				},
				DeployTokenPoolInput: &tokensapi.DeployTokenPoolInput{
					TokenPoolQualifier: "", // a sensible default will be used
					RateLimitAdmin:     RateLimitAdminEVM,
					PoolType:           cciputils.BurnMintTokenPool.String(),
				},
				TokenTransferConfig: &tokensapi.TokenTransferConfig{
					TokenRef:     datastore.AddressRef{},
					TokenPoolRef: datastore.AddressRef{},
					RemoteChains: map[uint64]tokensapi.RemoteChainConfig[*datastore.AddressRef, datastore.AddressRef]{
						solChainSel: {OutboundRateLimiterConfig: &realWorldRL},
					},
				},
			},
		},
	})
	require.NoError(t, err)
	MergeAddresses(t, env, expandTE2.DataStore)
	testhelpers.ProcessTimelockProposals(t, *env, expandTE2.MCMSTimelockProposals, false)

	// For Solana, the changeset should have overridden the pending admin in TAR with the timelock
	// signer and then accepted that timelock admin, all within the same batch. We assert that the
	// final TAR state is correct and reflects the intended final state after all changesets in the
	// batch, not stale state partway through the batch execution.
	tarState = fetchTAR(true)

	// Assert final expected Solana TAR state
	timelockSigner := solanautils.GetTimelockSignerPDA(
		env.DataStore.Addresses().Filter(), solChainSel, cciputils.CLLQualifier,
	)
	require.Equal(
		t, solMintPK, tarState.Mint,
		"TAR mint must stay bound to the SPL mint deployed in the first token expansion",
	)
	require.Equal(
		t, timelockSigner, tarState.Administrator,
		"the TAR administrator should be the MCMS timelock signer after override + accept in ConfigureTokenForTransfers",
	)
	require.True(
		t, tarState.PendingAdministrator.IsZero(),
		"no pending administrator should remain after TAR accept completes",
	)
	require.NotEqual(
		t, customer, tarState.Administrator,
		"third-party key from the between-pass Register must not be the final TAR administrator",
	)
}
