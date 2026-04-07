package deployment

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gagliardetto/solana-go"
	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/stretchr/testify/require"

	evmadapters "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/adapters"
	bnmERC20ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/burn_mint_erc20"
	evmseqV1_6_0 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/sequences"
	bnmpool "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_5_1/burn_mint_token_pool"
	lrpool "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_5_1/lock_release_token_pool"
	solanautils "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/utils"
	solseqV1_6_0 "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_0/sequences"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/v1_6_0/lockrelease_token_pool"
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

	env, err := environment.New(t.Context(),
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
							selB: {DefaultFinalityOutboundRateLimiterConfig: defaultRL},
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
							selA: {DefaultFinalityOutboundRateLimiterConfig: defaultRL},
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
							selB: {DefaultFinalityOutboundRateLimiterConfig: defaultRL},
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
							selA: {DefaultFinalityOutboundRateLimiterConfig: defaultRL},
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
							selB: {DefaultFinalityOutboundRateLimiterConfig: defaultRL},
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
							selA: {DefaultFinalityOutboundRateLimiterConfig: defaultRL},
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

		// First call: deploy tokens only (no pools, no transfer config)
		output, err := tokensapi.TokenExpansion().Apply(*env, tokensapi.TokenExpansionInput{
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
		MergeAddresses(t, env, output.DataStore)

		// Tokens should exist, no pools yet
		assertTokenExists(t, env, selA, tokenSymbolA, "Scenario3 Token A", 18)
		assertTokenExists(t, env, selB, tokenSymbolB, "Scenario3 Token B", 18)
		_, err = evmAdapter.FindLatestAddressRef(env.DataStore, datastore.AddressRef{ChainSelector: selA, Qualifier: "S3_POOL_A", Type: datastore.ContractType(bmPoolType)})
		require.Error(t, err, "pool should not exist after token-only deploy")

		// Second call: deploy pools referencing existing tokens, and connect
		poolQualA := "S3_POOL_A"
		poolQualB := "S3_POOL_B"

		output, err = tokensapi.TokenExpansion().Apply(*env, tokensapi.TokenExpansionInput{
			ChainAdapterVersion: v1_6_0_scenarios,
			MCMS:                NewDefaultInputForMCMS("Scenario 3 pools"),
			TokenExpansionInputPerChain: map[uint64]tokensapi.TokenExpansionInputPerChain{
				selA: {
					TokenPoolVersion: v1_5_1_scenarios,
					DeployTokenPoolInput: &tokensapi.DeployTokenPoolInput{
						TokenPoolQualifier: poolQualA,
						PoolType:           bmPoolType.String(),
						TokenRef: &datastore.AddressRef{
							Qualifier: tokenSymbolA,
							Type:      datastore.ContractType(bnmERC20ops.ContractType),
						},
					},
					TokenTransferConfig: &tokensapi.TokenTransferConfig{
						RemoteChains: map[uint64]tokensapi.RemoteChainConfig[*datastore.AddressRef, datastore.AddressRef]{
							selB: {DefaultFinalityOutboundRateLimiterConfig: defaultRL},
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
							selA: {DefaultFinalityOutboundRateLimiterConfig: defaultRL},
						},
					},
				},
			},
		})
		require.NoError(t, err)
		MergeAddresses(t, env, output.DataStore)
		testhelpers.ProcessTimelockProposals(t, *env, output.MCMSTimelockProposals, false)

		// Verify
		tokAddrA := assertTokenExists(t, env, selA, tokenSymbolA, "Scenario3 Token A", 18)
		tokAddrB := assertTokenExists(t, env, selB, tokenSymbolB, "Scenario3 Token B", 18)

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
								DefaultFinalityOutboundRateLimiterConfig: defaultRL,
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
								DefaultFinalityOutboundRateLimiterConfig: defaultRL,
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
}

// ---------------------------------------------------------------------------
// Scenario 5: Solana LockRelease + EVM BurnMint cross-chain
// ---------------------------------------------------------------------------

func TestTokenExpansionScenariosSolana(t *testing.T) {
	evmChainSel := chainsel.TEST_90000001.Selector
	solChainSel := chainsel.SOLANA_DEVNET.Selector

	programsPath, ds, err := PreloadSolanaEnvironment(t, solChainSel)
	require.NoError(t, err)

	env, err := environment.New(t.Context(),
		environment.WithSolanaContainer(t, []uint64{solChainSel}, programsPath, solanaProgramIDs),
		environment.WithEVMSimulated(t, []uint64{evmChainSel}),
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
		evmChainSel: NewDefaultDeploymentConfigForEVM(v1_6_0_scenarios),
		solChainSel: NewDefaultDeploymentConfigForSolana(v1_6_0_scenarios),
	}
	output, err := deployapi.DeployContracts(deployRegistry).Apply(*env, deployapi.ContractDeploymentConfig{Chains: deployInput, MCMS: mcms.Input{}})
	require.NoError(t, err)
	MergeAddresses(t, env, output.DataStore)

	DeployMCMS(t, env, evmChainSel, []string{cciputils.CLLQualifier})
	DeployMCMS(t, env, solChainSel, []string{cciputils.CLLQualifier})
	EVMTransferOwnership(t, env, evmChainSel)
	SolanaTransferOwnership(t, env, solChainSel)

	t.Run("Scenario5_SolanaLRToEVMBM", func(t *testing.T) {
		evmTokenSymbol := "S5_EVM_TOK"
		evmPoolQual := "S5_EVM_POOL"
		solTokenSymbol := "S5_SOL_TOK"
		defaultMaxSupply := uint64(1e6)
		defaultPreMint := uint64(1e5)

		defaultRL := tokensapi.RateLimiterConfigFloatInput{
			Capacity:  100,
			Rate:      10,
			IsEnabled: true,
		}

		output, err = tokensapi.TokenExpansion().Apply(*env, tokensapi.TokenExpansionInput{
			ChainAdapterVersion: v1_6_0_scenarios,
			MCMS:                NewDefaultInputForMCMS("Scenario 5"),
			TokenExpansionInputPerChain: map[uint64]tokensapi.TokenExpansionInputPerChain{
				evmChainSel: {
					TokenPoolVersion: v1_5_1_scenarios,
					DeployTokenInput: &tokensapi.DeployTokenInput{
						Name: "Scenario5 EVM Token", Symbol: evmTokenSymbol, Decimals: 18,
						Type: bnmERC20ops.ContractType, Supply: &defaultMaxSupply, PreMint: &defaultPreMint,
					},
					DeployTokenPoolInput: &tokensapi.DeployTokenPoolInput{
						TokenPoolQualifier: evmPoolQual,
						PoolType:           cciputils.BurnMintTokenPool.String(),
					},
					TokenTransferConfig: &tokensapi.TokenTransferConfig{
						RemoteChains: map[uint64]tokensapi.RemoteChainConfig[*datastore.AddressRef, datastore.AddressRef]{
							solChainSel: {DefaultFinalityOutboundRateLimiterConfig: defaultRL},
						},
					},
				},
				solChainSel: {
					TokenPoolVersion: v1_6_0_scenarios,
					DeployTokenInput: &tokensapi.DeployTokenInput{
						Name: "Scenario5 SOL Token", Symbol: solTokenSymbol, Decimals: 9,
						Type:                  solanautils.SPLTokens,
						ExternalAdmin:         solana.NewWallet().PublicKey().String(),
						DisableFreezeAuthority: true,
						Senders:               []string{solChain.DeployerKey.PublicKey().String()},
						PreMint:               &defaultPreMint,
					},
					DeployTokenPoolInput: &tokensapi.DeployTokenPoolInput{
						TokenPoolQualifier: "",
						PoolType:           cciputils.LockReleaseTokenPool.String(),
					},
					TokenTransferConfig: &tokensapi.TokenTransferConfig{
						RemoteChains: map[uint64]tokensapi.RemoteChainConfig[*datastore.AddressRef, datastore.AddressRef]{
							evmChainSel: {DefaultFinalityOutboundRateLimiterConfig: defaultRL},
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

		// Verify Solana router address is set
		routerAddr, err := solAdapter.GetRouterAddress(env.DataStore, solChainSel)
		require.NoError(t, err)
		require.NotEmpty(t, routerAddr, "Solana router should be deployed")
	})
}
