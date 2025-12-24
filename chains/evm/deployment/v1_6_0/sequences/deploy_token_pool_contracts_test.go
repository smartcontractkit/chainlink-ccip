package sequences

import (
	"math/big"
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	chain_selectors "github.com/smartcontractkit/chain-selectors"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/burn_mint_erc20"
	rmnproxyops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/rmn_proxy"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/router"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_1/operations/burn_from_mint_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_1/operations/burn_mint_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_1/operations/burn_mint_with_lock_release_flag_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_1/operations/burn_to_address_mint_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_1/operations/burn_with_from_mint_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_1/operations/lock_release_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_1/operations/siloed_lock_release_token_pool"
	burn_mint_token_pool_bindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_1/burn_mint_token_pool"
	tokenapi "github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	bnm_bindings "github.com/smartcontractkit/chainlink-evm/gethwrappers/shared/generated/initial/burn_mint_erc20"
)

// TestDeployTokenPool tests the DeployTokenPool sequence for various token pool types.
// It verifies that token pools are deployed correctly with the expected configuration.
func TestDeployTokenPool(t *testing.T) {
	t.Parallel()

	chainSelector := chain_selectors.ETHEREUM_MAINNET.Selector
	tokenSymbol := "TEST"
	tokenDecimals := uint8(18)

	testCases := []struct {
		name                string
		poolType            cldf.ContractType
		poolVersion         *semver.Version
		acceptLiquidity     *bool  // only for LockReleaseTokenPool v1.5.1
		burnAddress         string // only for BurnToAddress pools
		allowlist           []string
		expectedTypeVersion string
	}{
		// v1.6.1 pools
		{
			name:                "BurnMintTokenPool_v1_6_1",
			poolType:            burn_mint_token_pool.ContractType,
			poolVersion:         utils.Version_1_6_1,
			expectedTypeVersion: "BurnMintTokenPool 1.6.1",
		},
		{
			name:                "BurnFromMintTokenPool_v1_6_1",
			poolType:            burn_from_mint_token_pool.ContractType,
			poolVersion:         utils.Version_1_6_1,
			expectedTypeVersion: "BurnFromMintTokenPool 1.6.1",
		},
		{
			name:                "BurnMintWithLockReleaseFlagTokenPool_v1_6_1",
			poolType:            burn_mint_with_lock_release_flag_token_pool.ContractType,
			poolVersion:         utils.Version_1_6_1,
			expectedTypeVersion: "BurnMintTokenPool 1.6.1",
		},
		{
			name:                "BurnToAddressMintTokenPool_v1_6_1",
			poolType:            burn_to_address_mint_token_pool.ContractType,
			poolVersion:         utils.Version_1_6_1,
			burnAddress:         "0x000000000000000000000000000000000000dead",
			expectedTypeVersion: "BurnToAddressTokenPool 1.6.1",
		},
		{
			name:                "BurnWithFromMintTokenPool_v1_6_1",
			poolType:            burn_with_from_mint_token_pool.ContractType,
			poolVersion:         utils.Version_1_6_1,
			expectedTypeVersion: "BurnWithFromMintTokenPool 1.6.1",
		},
		{
			name:                "LockReleaseTokenPool_v1_6_1",
			poolType:            lock_release_token_pool.ContractType,
			poolVersion:         utils.Version_1_6_1,
			expectedTypeVersion: "LockReleaseTokenPool 1.6.1",
		},
		{
			name:                "SiloedLockReleaseTokenPool_v1_6_1",
			poolType:            siloed_lock_release_token_pool.ContractType,
			poolVersion:         utils.Version_1_6_1,
			expectedTypeVersion: "SiloedLockReleaseTokenPool 1.6.1",
		},
		{
			name:                "BurnMintTokenPool_v1_6_1_with_allowlist",
			poolType:            burn_mint_token_pool.ContractType,
			poolVersion:         utils.Version_1_6_1,
			allowlist:           []string{"0x1111111111111111111111111111111111111111", "0x2222222222222222222222222222222222222222"},
			expectedTypeVersion: "BurnMintTokenPool 1.6.1",
		},
		// NOTE: v1.6.0 external minter pools (BurnMintWithExternalMinterTokenPool, HybridWithExternalMinterTokenPool)
		// are excluded from unit tests because they require a real minter contract (token governor) to be deployed.
		// These pools should be tested in integration tests with proper minter contract setup.

		// v1.5.1 pools (using same ContractType constants since they're identical strings)
		{
			name:                "BurnMintTokenPool_v1_5_1",
			poolType:            burn_mint_token_pool.ContractType,
			poolVersion:         utils.Version_1_5_1,
			expectedTypeVersion: "BurnMintTokenPool 1.5.1",
		},
		{
			name:                "BurnFromMintTokenPool_v1_5_1",
			poolType:            burn_from_mint_token_pool.ContractType,
			poolVersion:         utils.Version_1_5_1,
			expectedTypeVersion: "BurnFromMintTokenPool 1.5.1",
		},
		{
			name:                "BurnToAddressMintTokenPool_v1_5_1",
			poolType:            burn_to_address_mint_token_pool.ContractType,
			poolVersion:         utils.Version_1_5_1,
			burnAddress:         "0x000000000000000000000000000000000000dead",
			expectedTypeVersion: "BurnToAddressTokenPool 1.5.1",
		},
		{
			name:                "BurnWithFromMintTokenPool_v1_5_1",
			poolType:            burn_with_from_mint_token_pool.ContractType,
			poolVersion:         utils.Version_1_5_1,
			expectedTypeVersion: "BurnWithFromMintTokenPool 1.5.1",
		},
		{
			name:                "LockReleaseTokenPool_v1_5_1_accept_liquidity",
			poolType:            lock_release_token_pool.ContractType,
			poolVersion:         utils.Version_1_5_1,
			acceptLiquidity:     boolPtr(true),
			expectedTypeVersion: "LockReleaseTokenPool 1.5.1",
		},
		{
			name:                "LockReleaseTokenPool_v1_5_1_no_liquidity",
			poolType:            lock_release_token_pool.ContractType,
			poolVersion:         utils.Version_1_5_1,
			acceptLiquidity:     boolPtr(false),
			expectedTypeVersion: "LockReleaseTokenPool 1.5.1",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			e, err := environment.New(t.Context(),
				environment.WithEVMSimulated(t, []uint64{chainSelector}),
			)
			require.NoError(t, err, "Failed to create test environment")
			require.NotNil(t, e, "Environment should be created")

			chain := e.BlockChains.EVMChains()[chainSelector]
			ds := datastore.NewMemoryDataStore()

			// Deploy and add a token to the datastore
			tokenAddr := deployTestToken(t, chain, tokenSymbol, tokenDecimals)
			err = ds.Addresses().Add(datastore.AddressRef{
				Type:          datastore.ContractType(burn_mint_erc20.ContractType),
				Version:       semver.MustParse("1.0.0"),
				Address:       tokenAddr.Hex(),
				ChainSelector: chainSelector,
				Qualifier:     tokenSymbol,
			})
			require.NoError(t, err, "Failed to add token address to datastore")

			// Add Router address to the datastore (mock address for testing)
			routerAddress := common.HexToAddress("0x1234567890abcdef1234567890abcdef12345678")
			err = ds.Addresses().Add(datastore.AddressRef{
				Type:          datastore.ContractType(router.ContractType),
				Version:       router.Version,
				Address:       routerAddress.Hex(),
				ChainSelector: chainSelector,
			})
			require.NoError(t, err, "Failed to add router address to datastore")

			// Add RMNProxy address to the datastore (mock address for testing)
			rmnProxyAddress := common.HexToAddress("0xabcdef1234567890abcdef1234567890abcdef12")
			err = ds.Addresses().Add(datastore.AddressRef{
				Type:          datastore.ContractType(rmnproxyops.ContractType),
				Version:       semver.MustParse("1.0.0"),
				Address:       rmnProxyAddress.Hex(),
				ChainSelector: chainSelector,
			})
			require.NoError(t, err, "Failed to add RMN proxy address to datastore")

			// Seal and set the datastore
			e.DataStore = ds.Seal()

			// Build input for DeployTokenPool sequence
			input := tokenapi.DeployTokenPoolInput{
				TokenSymbol:       tokenSymbol,
				PoolType:          string(tc.poolType),
				TokenPoolVersion:  tc.poolVersion,
				Allowlist:         tc.allowlist,
				AcceptLiquidity:   tc.acceptLiquidity,
				BurnAddress:       tc.burnAddress,
				ChainSelector:     chainSelector,
				ExistingDataStore: e.DataStore,
			}

			// Execute the DeployTokenPool sequence
			report, err := cldf_ops.ExecuteSequence(e.OperationsBundle, DeployTokenPool, e.BlockChains, input)
			require.NoError(t, err, "Failed to execute DeployTokenPool sequence for %s", tc.name)
			require.NotNil(t, report, "DeployTokenPool report should not be nil for %s", tc.name)

			// Verify output has at least one address
			require.GreaterOrEqual(t, len(report.Output.Addresses), 1, "Should have at least one deployed address")

			// Find the pool address in the output
			var poolRef datastore.AddressRef
			poolFound := false
			for _, addr := range report.Output.Addresses {
				if addr.Type == datastore.ContractType(tc.poolType) &&
					addr.ChainSelector == chainSelector &&
					addr.Qualifier == tokenSymbol {
					poolRef = addr
					poolFound = true
					break
				}
			}
			require.True(t, poolFound, "Pool %s should be found in deployed addresses", tc.name)
			require.NotEmpty(t, poolRef.Address, "Pool address should not be empty")
			t.Logf("Deployed %s at address: %s", tc.name, poolRef.Address)

			// Make on-chain calls to verify the pool was deployed correctly
			poolAddr := common.HexToAddress(poolRef.Address)
			poolContract, err := burn_mint_token_pool_bindings.NewBurnMintTokenPool(poolAddr, chain.Client)
			require.NoError(t, err, "Failed to create pool contract binding")

			// Verify TypeAndVersion
			typeAndVersion, err := poolContract.TypeAndVersion(&bind.CallOpts{})
			require.NoError(t, err, "Failed to get TypeAndVersion from pool")
			require.Equal(t, tc.expectedTypeVersion, typeAndVersion, "TypeAndVersion mismatch for %s", tc.name)
			t.Logf("  TypeAndVersion: %s", typeAndVersion)

			// Verify Token address
			onChainToken, err := poolContract.GetToken(&bind.CallOpts{})
			require.NoError(t, err, "Failed to get token address from pool")
			require.Equal(t, tokenAddr, onChainToken, "Token address mismatch for %s", tc.name)
			t.Logf("  Token: %s", onChainToken.Hex())

			// Verify Token decimals
			onChainDecimals, err := poolContract.GetTokenDecimals(&bind.CallOpts{})
			require.NoError(t, err, "Failed to get token decimals from pool")
			require.Equal(t, tokenDecimals, onChainDecimals, "Token decimals mismatch for %s", tc.name)
			t.Logf("  TokenDecimals: %d", onChainDecimals)

			// Verify Router
			onChainRouter, err := poolContract.GetRouter(&bind.CallOpts{})
			require.NoError(t, err, "Failed to get router address from pool")
			require.Equal(t, routerAddress, onChainRouter, "Router address mismatch for %s", tc.name)
			t.Logf("  Router: %s", onChainRouter.Hex())

			// Verify RMN Proxy
			onChainRmnProxy, err := poolContract.GetRmnProxy(&bind.CallOpts{})
			require.NoError(t, err, "Failed to get RMN proxy address from pool")
			require.Equal(t, rmnProxyAddress, onChainRmnProxy, "RMN proxy address mismatch for %s", tc.name)
			t.Logf("  RmnProxy: %s", onChainRmnProxy.Hex())

			// Verify allowlist if provided
			if len(tc.allowlist) > 0 {
				allowlistEnabled, err := poolContract.GetAllowListEnabled(&bind.CallOpts{})
				require.NoError(t, err, "Failed to get allowlist enabled status from pool")
				require.True(t, allowlistEnabled, "Allowlist should be enabled for %s", tc.name)
				t.Logf("  AllowlistEnabled: %v", allowlistEnabled)

				onChainAllowlist, err := poolContract.GetAllowList(&bind.CallOpts{})
				require.NoError(t, err, "Failed to get allowlist from pool")
				require.Equal(t, len(tc.allowlist), len(onChainAllowlist), "Allowlist length mismatch for %s", tc.name)
				for i, addr := range tc.allowlist {
					expectedAddr := common.HexToAddress(addr)
					require.Equal(t, expectedAddr, onChainAllowlist[i], "Allowlist address mismatch at index %d for %s", i, tc.name)
				}
				t.Logf("  Allowlist: %v", onChainAllowlist)
			}
		})
	}
}

// TestDeployTokenPool_AlreadyDeployed verifies that the sequence skips deployment
// when the token pool is already in the datastore.
func TestDeployTokenPool_AlreadyDeployed(t *testing.T) {
	t.Parallel()

	chainSelector := chain_selectors.ETHEREUM_MAINNET.Selector
	tokenSymbol := "TEST"
	tokenDecimals := uint8(18)
	poolType := burn_mint_token_pool.ContractType

	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{chainSelector}),
	)
	require.NoError(t, err, "Failed to create test environment")
	require.NotNil(t, e, "Environment should be created")

	chain := e.BlockChains.EVMChains()[chainSelector]
	ds := datastore.NewMemoryDataStore()

	// Deploy and add a token to the datastore
	tokenAddr := deployTestToken(t, chain, tokenSymbol, tokenDecimals)
	err = ds.Addresses().Add(datastore.AddressRef{
		Type:          datastore.ContractType(burn_mint_erc20.ContractType),
		Version:       semver.MustParse("1.0.0"),
		Address:       tokenAddr.Hex(),
		ChainSelector: chainSelector,
		Qualifier:     tokenSymbol,
	})
	require.NoError(t, err, "Failed to add token address to datastore")

	// Add Router address to the datastore
	routerAddress := common.HexToAddress("0x1234567890abcdef1234567890abcdef12345678")
	err = ds.Addresses().Add(datastore.AddressRef{
		Type:          datastore.ContractType(router.ContractType),
		Version:       router.Version,
		Address:       routerAddress.Hex(),
		ChainSelector: chainSelector,
	})
	require.NoError(t, err, "Failed to add router address to datastore")

	// Add RMNProxy address to the datastore
	rmnProxyAddress := common.HexToAddress("0xabcdef1234567890abcdef1234567890abcdef12")
	err = ds.Addresses().Add(datastore.AddressRef{
		Type:          datastore.ContractType(rmnproxyops.ContractType),
		Version:       semver.MustParse("1.0.0"),
		Address:       rmnProxyAddress.Hex(),
		ChainSelector: chainSelector,
	})
	require.NoError(t, err, "Failed to add RMN proxy address to datastore")

	// Add an existing pool to the datastore (simulating already deployed)
	existingPoolAddress := common.HexToAddress("0xdeadbeefdeadbeefdeadbeefdeadbeefdeadbeef")
	err = ds.Addresses().Add(datastore.AddressRef{
		Type:          datastore.ContractType(poolType),
		Version:       utils.Version_1_6_1,
		Address:       existingPoolAddress.Hex(),
		ChainSelector: chainSelector,
		Qualifier:     tokenSymbol,
	})
	require.NoError(t, err, "Failed to add existing pool address to datastore")

	e.DataStore = ds.Seal()

	input := tokenapi.DeployTokenPoolInput{
		TokenSymbol:       tokenSymbol,
		PoolType:          string(poolType),
		TokenPoolVersion:  utils.Version_1_6_1,
		ChainSelector:     chainSelector,
		ExistingDataStore: e.DataStore,
	}

	// Execute the sequence - it should return early without deploying
	report, err := cldf_ops.ExecuteSequence(e.OperationsBundle, DeployTokenPool, e.BlockChains, input)
	require.NoError(t, err, "Should not error when pool already exists")
	require.NotNil(t, report, "Report should not be nil")

	// Verify no new addresses were deployed
	require.Equal(t, 0, len(report.Output.Addresses), "Should not deploy any new addresses when pool already exists")
}

// TestDeployTokenPool_MissingTokenPoolVersion verifies that the sequence fails
// when TokenPoolVersion is not provided.
func TestDeployTokenPool_MissingTokenPoolVersion(t *testing.T) {
	t.Parallel()

	chainSelector := chain_selectors.ETHEREUM_MAINNET.Selector

	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{chainSelector}),
	)
	require.NoError(t, err, "Failed to create test environment")

	ds := datastore.NewMemoryDataStore()
	e.DataStore = ds.Seal()

	input := tokenapi.DeployTokenPoolInput{
		TokenSymbol:       "TEST",
		PoolType:          string(burn_mint_token_pool.ContractType),
		TokenPoolVersion:  nil, // Missing version
		ChainSelector:     chainSelector,
		ExistingDataStore: e.DataStore,
	}

	_, err = cldf_ops.ExecuteSequence(e.OperationsBundle, DeployTokenPool, e.BlockChains, input)
	require.Error(t, err, "Should error when TokenPoolVersion is missing")
	require.Contains(t, err.Error(), "TokenPoolVersion is required", "Error message should mention TokenPoolVersion")
}

// TestDeployTokenPool_UnsupportedPoolType verifies that the sequence fails
// when an unsupported pool type is provided.
func TestDeployTokenPool_UnsupportedPoolType(t *testing.T) {
	t.Parallel()

	chainSelector := chain_selectors.ETHEREUM_MAINNET.Selector
	tokenSymbol := "TEST"
	tokenDecimals := uint8(18)

	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{chainSelector}),
	)
	require.NoError(t, err, "Failed to create test environment")

	chain := e.BlockChains.EVMChains()[chainSelector]
	ds := datastore.NewMemoryDataStore()

	// Deploy and add a token to the datastore
	tokenAddr := deployTestToken(t, chain, tokenSymbol, tokenDecimals)
	err = ds.Addresses().Add(datastore.AddressRef{
		Type:          datastore.ContractType(burn_mint_erc20.ContractType),
		Version:       semver.MustParse("1.0.0"),
		Address:       tokenAddr.Hex(),
		ChainSelector: chainSelector,
		Qualifier:     tokenSymbol,
	})
	require.NoError(t, err, "Failed to add token address to datastore")

	// Add Router address
	routerAddress := common.HexToAddress("0x1234567890abcdef1234567890abcdef12345678")
	err = ds.Addresses().Add(datastore.AddressRef{
		Type:          datastore.ContractType(router.ContractType),
		Version:       router.Version,
		Address:       routerAddress.Hex(),
		ChainSelector: chainSelector,
	})
	require.NoError(t, err)

	// Add RMNProxy address
	rmnProxyAddress := common.HexToAddress("0xabcdef1234567890abcdef1234567890abcdef12")
	err = ds.Addresses().Add(datastore.AddressRef{
		Type:          datastore.ContractType(rmnproxyops.ContractType),
		Version:       semver.MustParse("1.0.0"),
		Address:       rmnProxyAddress.Hex(),
		ChainSelector: chainSelector,
	})
	require.NoError(t, err)

	e.DataStore = ds.Seal()

	input := tokenapi.DeployTokenPoolInput{
		TokenSymbol:       tokenSymbol,
		PoolType:          "UnsupportedPoolType",
		TokenPoolVersion:  utils.Version_1_6_1,
		ChainSelector:     chainSelector,
		ExistingDataStore: e.DataStore,
	}

	_, err = cldf_ops.ExecuteSequence(e.OperationsBundle, DeployTokenPool, e.BlockChains, input)
	require.Error(t, err, "Should error when pool type is unsupported")
	require.Contains(t, err.Error(), "unsupported token pool type and version", "Error message should mention unsupported pool type")
}

// TestDeployTokenPool_MissingRouter verifies that the sequence fails
// when the router is not found in the datastore.
func TestDeployTokenPool_MissingRouter(t *testing.T) {
	t.Parallel()

	chainSelector := chain_selectors.ETHEREUM_MAINNET.Selector
	tokenSymbol := "TEST"
	tokenDecimals := uint8(18)

	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{chainSelector}),
	)
	require.NoError(t, err, "Failed to create test environment")

	chain := e.BlockChains.EVMChains()[chainSelector]
	ds := datastore.NewMemoryDataStore()

	// Deploy and add a token to the datastore
	tokenAddr := deployTestToken(t, chain, tokenSymbol, tokenDecimals)
	err = ds.Addresses().Add(datastore.AddressRef{
		Type:          datastore.ContractType(burn_mint_erc20.ContractType),
		Version:       semver.MustParse("1.0.0"),
		Address:       tokenAddr.Hex(),
		ChainSelector: chainSelector,
		Qualifier:     tokenSymbol,
	})
	require.NoError(t, err, "Failed to add token address to datastore")

	// Note: Router is NOT added to the datastore

	// Add RMNProxy address
	rmnProxyAddress := common.HexToAddress("0xabcdef1234567890abcdef1234567890abcdef12")
	err = ds.Addresses().Add(datastore.AddressRef{
		Type:          datastore.ContractType(rmnproxyops.ContractType),
		Version:       semver.MustParse("1.0.0"),
		Address:       rmnProxyAddress.Hex(),
		ChainSelector: chainSelector,
	})
	require.NoError(t, err)

	e.DataStore = ds.Seal()

	input := tokenapi.DeployTokenPoolInput{
		TokenSymbol:       tokenSymbol,
		PoolType:          string(burn_mint_token_pool.ContractType),
		TokenPoolVersion:  utils.Version_1_6_1,
		ChainSelector:     chainSelector,
		ExistingDataStore: e.DataStore,
	}

	_, err = cldf_ops.ExecuteSequence(e.OperationsBundle, DeployTokenPool, e.BlockChains, input)
	require.Error(t, err, "Should error when router is not found in datastore")
	require.Contains(t, err.Error(), "is not found in datastore", "Error message should mention not found in datastore")
}

// TestDeployTokenPool_MissingRMNProxy verifies that the sequence fails
// when the RMN proxy is not found in the datastore.
func TestDeployTokenPool_MissingRMNProxy(t *testing.T) {
	t.Parallel()

	chainSelector := chain_selectors.ETHEREUM_MAINNET.Selector
	tokenSymbol := "TEST"
	tokenDecimals := uint8(18)

	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{chainSelector}),
	)
	require.NoError(t, err, "Failed to create test environment")

	chain := e.BlockChains.EVMChains()[chainSelector]
	ds := datastore.NewMemoryDataStore()

	// Deploy and add a token to the datastore
	tokenAddr := deployTestToken(t, chain, tokenSymbol, tokenDecimals)
	err = ds.Addresses().Add(datastore.AddressRef{
		Type:          datastore.ContractType(burn_mint_erc20.ContractType),
		Version:       semver.MustParse("1.0.0"),
		Address:       tokenAddr.Hex(),
		ChainSelector: chainSelector,
		Qualifier:     tokenSymbol,
	})
	require.NoError(t, err, "Failed to add token address to datastore")

	// Add Router address
	routerAddress := common.HexToAddress("0x1234567890abcdef1234567890abcdef12345678")
	err = ds.Addresses().Add(datastore.AddressRef{
		Type:          datastore.ContractType(router.ContractType),
		Version:       router.Version,
		Address:       routerAddress.Hex(),
		ChainSelector: chainSelector,
	})
	require.NoError(t, err)

	// Note: RMNProxy is NOT added to the datastore

	e.DataStore = ds.Seal()

	input := tokenapi.DeployTokenPoolInput{
		TokenSymbol:       tokenSymbol,
		PoolType:          string(burn_mint_token_pool.ContractType),
		TokenPoolVersion:  utils.Version_1_6_1,
		ChainSelector:     chainSelector,
		ExistingDataStore: e.DataStore,
	}

	_, err = cldf_ops.ExecuteSequence(e.OperationsBundle, DeployTokenPool, e.BlockChains, input)
	require.Error(t, err, "Should error when RMN proxy is not found in datastore")
	require.Contains(t, err.Error(), "is not found in datastore", "Error message should mention not found in datastore")
}

// TestDeployTokenPool_MissingToken verifies that the sequence fails
// when the token is not found in the datastore.
func TestDeployTokenPool_MissingToken(t *testing.T) {
	t.Parallel()

	chainSelector := chain_selectors.ETHEREUM_MAINNET.Selector
	tokenSymbol := "NONEXISTENT"

	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{chainSelector}),
	)
	require.NoError(t, err, "Failed to create test environment")

	ds := datastore.NewMemoryDataStore()

	// Add Router address
	routerAddress := common.HexToAddress("0x1234567890abcdef1234567890abcdef12345678")
	err = ds.Addresses().Add(datastore.AddressRef{
		Type:          datastore.ContractType(router.ContractType),
		Version:       router.Version,
		Address:       routerAddress.Hex(),
		ChainSelector: chainSelector,
	})
	require.NoError(t, err)

	// Add RMNProxy address
	rmnProxyAddress := common.HexToAddress("0xabcdef1234567890abcdef1234567890abcdef12")
	err = ds.Addresses().Add(datastore.AddressRef{
		Type:          datastore.ContractType(rmnproxyops.ContractType),
		Version:       semver.MustParse("1.0.0"),
		Address:       rmnProxyAddress.Hex(),
		ChainSelector: chainSelector,
	})
	require.NoError(t, err)

	e.DataStore = ds.Seal()

	input := tokenapi.DeployTokenPoolInput{
		TokenSymbol:       tokenSymbol, // Token not in datastore
		PoolType:          string(burn_mint_token_pool.ContractType),
		TokenPoolVersion:  utils.Version_1_6_1,
		ChainSelector:     chainSelector,
		ExistingDataStore: e.DataStore,
	}

	_, err = cldf_ops.ExecuteSequence(e.OperationsBundle, DeployTokenPool, e.BlockChains, input)
	require.Error(t, err, "Should error when token is not found in datastore")
	require.Contains(t, err.Error(), "is not found in datastore", "Error message should mention token not found")
}

// deployTestToken deploys a BurnMintERC20 token for testing purposes and returns its address.
func deployTestToken(t *testing.T, chain evm.Chain, symbol string, decimals uint8) common.Address {
	// Use the gobindings to deploy a test token
	tokenAddr, tx, _, err := bnm_bindings.DeployBurnMintERC20(
		chain.DeployerKey,
		chain.Client,
		"Test Token",
		symbol,
		decimals,
		big.NewInt(0).Mul(big.NewInt(1e9), big.NewInt(1e18)), // maxSupply: 1 billion
		big.NewInt(0), // preMint: 0
	)
	require.NoError(t, err, "Failed to deploy test token")

	_, err = deployment.ConfirmIfNoError(chain, tx, err)
	require.NoError(t, err, "Failed to confirm test token deployment")

	return tokenAddr
}

func boolPtr(b bool) *bool {
	return &b
}
