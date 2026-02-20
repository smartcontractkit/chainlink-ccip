package sequences_test

import (
	"math/big"
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	chain_selectors "github.com/smartcontractkit/chain-selectors"
	"github.com/stretchr/testify/require"

	contract_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/burn_mint_erc20"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/rmn_proxy"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/router"
	tar_ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/token_admin_registry"
	evm_adapters "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_1/adapters"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_1/operations/burn_mint_with_lock_release_flag_token_pool"
	tar_bindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_5_0/token_admin_registry"
	pool_bindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_1/burn_mint_with_lock_release_flag_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/adapters"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	burn_mint_erc20_bindings "github.com/smartcontractkit/chainlink-evm/gethwrappers/shared/generated/initial/burn_mint_erc20"
)

type nonCanonicalTestSetup struct {
	Router             common.Address
	RMN                common.Address
	TokenAdminRegistry common.Address
	USDCToken          common.Address
}

// nonCanonicalRemoteChainConfig returns a remote chain config with rate limiters disabled.
// Capacity and Rate must be non-nil for ABI packing in applyChainUpdates (token pool expects *big.Int).
func nonCanonicalRemoteChainConfig() adapters.RemoteCCTPChainConfig {
	return adapters.RemoteCCTPChainConfig{
		DefaultFinalityInboundRateLimiterConfig:  tokens.RateLimiterConfig{IsEnabled: false, Capacity: big.NewInt(0), Rate: big.NewInt(0)},
		DefaultFinalityOutboundRateLimiterConfig: tokens.RateLimiterConfig{IsEnabled: false, Capacity: big.NewInt(0), Rate: big.NewInt(0)},
		CustomFinalityInboundRateLimiterConfig:   tokens.RateLimiterConfig{IsEnabled: false, Capacity: big.NewInt(0), Rate: big.NewInt(0)},
		CustomFinalityOutboundRateLimiterConfig:  tokens.RateLimiterConfig{IsEnabled: false, Capacity: big.NewInt(0), Rate: big.NewInt(0)},
	}
}

func setupNonCanonicalTestEnvironment(t *testing.T, e *deployment.Environment, chainSelector uint64) nonCanonicalTestSetup {
	chain := e.BlockChains.EVMChains()[chainSelector]

	// Deploy RMN Proxy (use deployer as RMN for test)
	rmnProxyRef, err := contract_utils.MaybeDeployContract(e.OperationsBundle, rmn_proxy.Deploy, chain, contract_utils.DeployInput[rmn_proxy.ConstructorArgs]{
		TypeAndVersion: deployment.NewTypeAndVersion(rmn_proxy.ContractType, *rmn_proxy.Version),
		ChainSelector:  chainSelector,
		Args:           rmn_proxy.ConstructorArgs{RMN: chain.DeployerKey.From},
	}, nil)
	require.NoError(t, err, "Failed to deploy RMN proxy")
	rmnAddr := common.HexToAddress(rmnProxyRef.Address)

	// Deploy Router (WrappedNative = zero for test)
	routerRef, err := contract_utils.MaybeDeployContract(e.OperationsBundle, router.Deploy, chain, contract_utils.DeployInput[router.ConstructorArgs]{
		TypeAndVersion: deployment.NewTypeAndVersion(router.ContractType, *router.Version),
		ChainSelector:  chainSelector,
		Args: router.ConstructorArgs{
			WrappedNative: common.Address{},
			RMNProxy:      rmnAddr,
		},
	}, nil)
	require.NoError(t, err, "Failed to deploy Router")
	routerAddr := common.HexToAddress(routerRef.Address)

	// Deploy TokenAdminRegistry
	tarRef, err := contract_utils.MaybeDeployContract(e.OperationsBundle, tar_ops.Deploy, chain, contract_utils.DeployInput[tar_ops.ConstructorArgs]{
		TypeAndVersion: deployment.NewTypeAndVersion(tar_ops.ContractType, *tar_ops.Version),
		ChainSelector:  chainSelector,
		Args:           tar_ops.ConstructorArgs{},
	}, nil)
	require.NoError(t, err, "Failed to deploy TokenAdminRegistry")
	tarAddr := common.HexToAddress(tarRef.Address)

	// Deploy USDC (BurnMintERC20)
	usdcAddr, tx, _, err := burn_mint_erc20_bindings.DeployBurnMintERC20(
		chain.DeployerKey,
		chain.Client,
		"USD Coin",
		"USDC",
		6,
		big.NewInt(0),
		big.NewInt(0),
	)
	require.NoError(t, err, "Failed to deploy USDC token")
	_, err = chain.Confirm(tx)
	require.NoError(t, err, "Failed to confirm USDC deployment")

	// Merge addresses into datastore
	ds := datastore.NewMemoryDataStore()
	if e.DataStore != nil {
		err = ds.Merge(e.DataStore)
		require.NoError(t, err)
	}
	require.NoError(t, ds.Addresses().Add(rmnProxyRef))
	require.NoError(t, ds.Addresses().Add(routerRef))
	require.NoError(t, ds.Addresses().Add(tarRef))
	require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
		ChainSelector: chainSelector,
		Address:       usdcAddr.Hex(),
		Type:          datastore.ContractType(burn_mint_erc20.ContractType),
		Version:       semver.MustParse("1.0.0"),
		Qualifier:     "USDC",
	}))
	e.DataStore = ds.Seal()

	return nonCanonicalTestSetup{
		Router:             routerAddr,
		RMN:                rmnAddr,
		TokenAdminRegistry: tarAddr,
		USDCToken:          usdcAddr,
	}
}

func TestNonCanonicalUSDCChain_TwoChainsConnected(t *testing.T) {
	chainASelector := chain_selectors.ETHEREUM_TESTNET_SEPOLIA.Selector
	chainBSelector := chain_selectors.ETHEREUM_TESTNET_SEPOLIA_ARBITRUM_1.Selector

	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{chainASelector, chainBSelector}),
	)
	require.NoError(t, err)
	require.NotNil(t, e)

	ds := datastore.NewMemoryDataStore()
	e.DataStore = ds.Seal()

	// Set up both chains (shared datastore is updated by each setup)
	setupA := setupNonCanonicalTestEnvironment(t, e, chainASelector)
	setupB := setupNonCanonicalTestEnvironment(t, e, chainBSelector)

	chainA := e.BlockChains.EVMChains()[chainASelector]
	chainB := e.BlockChains.EVMChains()[chainBSelector]
	adapter := &evm_adapters.NonCanonicalUSDCChainAdapter{}

	// Deploy non-canonical USDC on chain A
	deployReportA, err := operations.ExecuteSequence(e.OperationsBundle, adapter.DeployCCTPChain(), adapters.DeployCCTPChainDeps{
		BlockChains: e.BlockChains,
		DataStore:   e.DataStore,
	}, adapters.DeployCCTPInput{
		TokenDecimals: 6,
		ChainSelector: chainASelector,
		USDCToken:     setupA.USDCToken.Hex(),
	})
	require.NoError(t, err, "Failed to deploy non-canonical USDC on chain A")

	ds = datastore.NewMemoryDataStore()
	require.NoError(t, ds.Merge(e.DataStore))
	for _, addr := range deployReportA.Output.Addresses {
		require.NoError(t, ds.Addresses().Add(addr))
	}
	e.DataStore = ds.Seal()

	// Deploy non-canonical USDC on chain B
	deployReportB, err := operations.ExecuteSequence(e.OperationsBundle, adapter.DeployCCTPChain(), adapters.DeployCCTPChainDeps{
		BlockChains: e.BlockChains,
		DataStore:   e.DataStore,
	}, adapters.DeployCCTPInput{
		TokenDecimals: 6,
		ChainSelector: chainBSelector,
		USDCToken:     setupB.USDCToken.Hex(),
	})
	require.NoError(t, err, "Failed to deploy non-canonical USDC on chain B")

	ds = datastore.NewMemoryDataStore()
	require.NoError(t, ds.Merge(e.DataStore))
	for _, addr := range deployReportB.Output.Addresses {
		require.NoError(t, ds.Addresses().Add(addr))
	}
	e.DataStore = ds.Seal()

	// Resolve pool refs for each chain
	poolsA := e.DataStore.Addresses().Filter(
		datastore.AddressRefByChainSelector(chainASelector),
		datastore.AddressRefByType(datastore.ContractType(burn_mint_with_lock_release_flag_token_pool.ContractType)),
	)
	poolsB := e.DataStore.Addresses().Filter(
		datastore.AddressRefByChainSelector(chainBSelector),
		datastore.AddressRefByType(datastore.ContractType(burn_mint_with_lock_release_flag_token_pool.ContractType)),
	)
	require.Len(t, poolsA, 1, "Expected one pool on chain A")
	require.Len(t, poolsB, 1, "Expected one pool on chain B")
	poolRefA := poolsA[0]
	poolRefB := poolsB[0]

	remoteChainsDeps := map[uint64]adapters.RemoteCCTPChain{
		chainASelector: adapter,
		chainBSelector: adapter,
	}

	// Configure chain A for lane to B
	_, err = operations.ExecuteSequence(e.OperationsBundle, adapter.ConfigureCCTPChainForLanes(), adapters.ConfigureCCTPChainForLanesDeps{
		BlockChains:  e.BlockChains,
		DataStore:    e.DataStore,
		RemoteChains: remoteChainsDeps,
	}, adapters.ConfigureCCTPChainForLanesInput{
		ChainSelector: chainASelector,
		USDCToken:     setupA.USDCToken.Hex(),
		RegisteredPoolRef: datastore.AddressRef{
			Type:    datastore.ContractType(burn_mint_with_lock_release_flag_token_pool.ContractType),
			Version: burn_mint_with_lock_release_flag_token_pool.Version,
		},
		RemoteRegisteredPoolRefs: map[uint64]datastore.AddressRef{
			chainBSelector: poolRefB,
		},
		RemoteChains: map[uint64]adapters.RemoteCCTPChainConfig{
			chainBSelector: nonCanonicalRemoteChainConfig(),
		},
	})
	require.NoError(t, err, "Failed to configure chain A for lanes")

	// Configure chain B for lane to A
	_, err = operations.ExecuteSequence(e.OperationsBundle, adapter.ConfigureCCTPChainForLanes(), adapters.ConfigureCCTPChainForLanesDeps{
		BlockChains:  e.BlockChains,
		DataStore:    e.DataStore,
		RemoteChains: remoteChainsDeps,
	}, adapters.ConfigureCCTPChainForLanesInput{
		ChainSelector: chainBSelector,
		USDCToken:     setupB.USDCToken.Hex(),
		RegisteredPoolRef: datastore.AddressRef{
			Type:    datastore.ContractType(burn_mint_with_lock_release_flag_token_pool.ContractType),
			Version: burn_mint_with_lock_release_flag_token_pool.Version,
		},
		RemoteRegisteredPoolRefs: map[uint64]datastore.AddressRef{
			chainASelector: poolRefA,
		},
		RemoteChains: map[uint64]adapters.RemoteCCTPChainConfig{
			chainASelector: nonCanonicalRemoteChainConfig(),
		},
	})
	require.NoError(t, err, "Failed to configure chain B for lanes")

	poolAddrA := common.HexToAddress(poolRefA.Address)
	poolAddrB := common.HexToAddress(poolRefB.Address)

	// Assertions: Chain A
	poolA, err := pool_bindings.NewBurnMintWithLockReleaseFlagTokenPool(poolAddrA, chainA.Client)
	require.NoError(t, err)

	supportedA, err := poolA.GetSupportedChains(nil)
	require.NoError(t, err)
	require.Contains(t, supportedA, chainBSelector, "Chain A pool should support chain B")

	remoteTokenA, err := poolA.GetRemoteToken(nil, chainBSelector)
	require.NoError(t, err)
	require.Equal(t, common.LeftPadBytes(setupB.USDCToken.Bytes(), 32), remoteTokenA, "Chain A pool remote token should be chain B USDC")

	remotePoolsA, err := poolA.GetRemotePools(nil, chainBSelector)
	require.NoError(t, err)
	require.Contains(t, remotePoolsA, common.LeftPadBytes(poolAddrB.Bytes(), 32), "Chain A pool should have chain B pool as remote")

	tokenAddrA, err := poolA.GetToken(nil)
	require.NoError(t, err)
	require.Equal(t, setupA.USDCToken, tokenAddrA, "Chain A pool token should be chain A USDC")

	// TokenAdminRegistry on A: pool should be registered for USDC
	tarA, err := tar_bindings.NewTokenAdminRegistry(setupA.TokenAdminRegistry, chainA.Client)
	require.NoError(t, err)
	tokenConfigA, err := tarA.GetTokenConfig(nil, setupA.USDCToken)
	require.NoError(t, err)
	require.Equal(t, poolAddrA, tokenConfigA.TokenPool, "TokenAdminRegistry on chain A should point to pool")

	// Adapter methods on chain A
	poolBytesA, err := adapter.PoolAddress(e.DataStore, e.BlockChains, chainASelector, poolRefA)
	require.NoError(t, err)
	require.Equal(t, poolAddrA.Bytes(), poolBytesA)

	tokenBytesA, err := adapter.TokenAddress(e.DataStore, e.BlockChains, chainASelector)
	require.NoError(t, err)
	require.Equal(t, setupA.USDCToken.Bytes(), tokenBytesA)

	// Assertions: Chain B
	poolB, err := pool_bindings.NewBurnMintWithLockReleaseFlagTokenPool(poolAddrB, chainB.Client)
	require.NoError(t, err)

	supportedB, err := poolB.GetSupportedChains(nil)
	require.NoError(t, err)
	require.Contains(t, supportedB, chainASelector, "Chain B pool should support chain A")

	remoteTokenB, err := poolB.GetRemoteToken(nil, chainASelector)
	require.NoError(t, err)
	require.Equal(t, common.LeftPadBytes(setupA.USDCToken.Bytes(), 32), remoteTokenB, "Chain B pool remote token should be chain A USDC")

	remotePoolsB, err := poolB.GetRemotePools(nil, chainASelector)
	require.NoError(t, err)
	require.Contains(t, remotePoolsB, common.LeftPadBytes(poolAddrA.Bytes(), 32), "Chain B pool should have chain A pool as remote")

	tokenAddrB, err := poolB.GetToken(nil)
	require.NoError(t, err)
	require.Equal(t, setupB.USDCToken, tokenAddrB, "Chain B pool token should be chain B USDC")

	tarB, err := tar_bindings.NewTokenAdminRegistry(setupB.TokenAdminRegistry, chainB.Client)
	require.NoError(t, err)
	tokenConfigB, err := tarB.GetTokenConfig(nil, setupB.USDCToken)
	require.NoError(t, err)
	require.Equal(t, poolAddrB, tokenConfigB.TokenPool, "TokenAdminRegistry on chain B should point to pool")

	poolBytesB, err := adapter.PoolAddress(e.DataStore, e.BlockChains, chainBSelector, poolRefB)
	require.NoError(t, err)
	require.Equal(t, poolAddrB.Bytes(), poolBytesB)

	tokenBytesB, err := adapter.TokenAddress(e.DataStore, e.BlockChains, chainBSelector)
	require.NoError(t, err)
	require.Equal(t, setupB.USDCToken.Bytes(), tokenBytesB)

	// Non-canonical adapter returns empty for CCTP caller / mint recipient
	cctpV1Dest, _ := adapter.CCTPV1AllowedCallerOnDest(e.DataStore, e.BlockChains, chainASelector)
	require.Empty(t, cctpV1Dest)
	cctpV2Dest, _ := adapter.CCTPV2AllowedCallerOnDest(e.DataStore, e.BlockChains, chainASelector)
	require.Empty(t, cctpV2Dest)
	allowedSource, _ := adapter.AllowedCallerOnSource(e.DataStore, e.BlockChains, chainASelector)
	require.Empty(t, allowedSource)
	mintRecipient, _ := adapter.MintRecipientOnDest(e.DataStore, e.BlockChains, chainASelector)
	require.Empty(t, mintRecipient)
}
