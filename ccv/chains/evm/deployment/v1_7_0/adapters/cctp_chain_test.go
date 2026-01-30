package adapters_test

import (
	"math/big"
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	chain_selectors "github.com/smartcontractkit/chain-selectors"
	evm_adapters "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/adapters"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/cctp_message_transmitter_proxy"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/cctp_through_ccv_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/cctp_verifier"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/create2_factory"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/usdc_token_pool_proxy"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/sequences"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/testsetup"
	cctp_message_transmitter_proxy_bindings "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/gobindings/generated/latest/cctp_message_transmitter_proxy"
	cctp_through_ccv_token_pool_bindings "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/gobindings/generated/latest/cctp_through_ccv_token_pool"
	cctp_verifier_bindings "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/gobindings/generated/latest/cctp_verifier"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/gobindings/generated/latest/mock_usdc_token_messenger"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/gobindings/generated/latest/mock_usdc_token_transmitter"
	usdc_token_pool_proxy_bindings "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/gobindings/generated/latest/usdc_token_pool_proxy"
	versioned_verifier_resolver_bindings "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/gobindings/generated/latest/versioned_verifier_resolver"
	contract_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/rmn_proxy"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/router"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/token_admin_registry"
	v1_6_1_burn_mint_token_pool "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_1/burn_mint_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/adapters"
	v1_7_0_changesets "github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/changesets"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	burn_mint_erc20_bindings "github.com/smartcontractkit/chainlink-evm/gethwrappers/shared/generated/initial/burn_mint_erc20"
	"github.com/stretchr/testify/require"
)

const (
	mechanismCCTPV2WithCCV = "CCTP_V2_WITH_CCV"
)

func convertMechanismToUint8(mechanism string) (uint8, error) {
	switch mechanism {
	case "CCTP_V1":
		return 1, nil
	case "CCTP_V2":
		return 2, nil
	case "LOCK_RELEASE":
		return 3, nil
	case mechanismCCTPV2WithCCV:
		return 4, nil
	default:
		return 0, nil
	}
}

type cctpTestSetup struct {
	Router             common.Address
	RMN                common.Address
	TokenAdminRegistry common.Address
	USDCToken          common.Address
	TokenMessenger     common.Address
	MessageTransmitter common.Address
}

func setupCCTPTestEnvironment(t *testing.T, e *deployment.Environment, chainSelector uint64) cctpTestSetup {
	chain := e.BlockChains.EVMChains()[chainSelector]

	// Deploy chain contracts
	create2FactoryRef, err := contract_utils.MaybeDeployContract(e.OperationsBundle, create2_factory.Deploy, chain, contract_utils.DeployInput[create2_factory.ConstructorArgs]{
		TypeAndVersion: deployment.NewTypeAndVersion(create2_factory.ContractType, *semver.MustParse("1.7.0")),
		ChainSelector:  chainSelector,
		Args: create2_factory.ConstructorArgs{
			AllowList: []common.Address{chain.DeployerKey.From},
		},
	}, nil)
	require.NoError(t, err, "Failed to deploy CREATE2Factory")
	chainReport, err := operations.ExecuteSequence(
		e.OperationsBundle,
		sequences.DeployChainContracts,
		chain,
		sequences.DeployChainContractsInput{
			ChainSelector:  chainSelector,
			ContractParams: testsetup.CreateBasicContractParams(),
			CREATE2Factory: common.HexToAddress(create2FactoryRef.Address),
		},
	)
	require.NoError(t, err, "Failed to deploy chain contracts")

	// Merge chain contract addresses into datastore
	ds := datastore.NewMemoryDataStore()
	err = ds.Merge(e.DataStore)
	require.NoError(t, err)
	for _, addr := range chainReport.Output.Addresses {
		err = ds.Addresses().Add(addr)
		require.NoError(t, err)
	}
	// Also add CREATE2Factory address
	err = ds.Addresses().Add(create2FactoryRef)
	require.NoError(t, err)
	e.DataStore = ds.Seal()

	var routerAddr, rmnAddr, tokenAdminRegistryAddr common.Address
	for _, addr := range chainReport.Output.Addresses {
		if addr.Type == datastore.ContractType(router.ContractType) {
			routerAddr = common.HexToAddress(addr.Address)
		}
		if addr.Type == datastore.ContractType(rmn_proxy.ContractType) {
			rmnAddr = common.HexToAddress(addr.Address)
		}
		if addr.Type == datastore.ContractType(token_admin_registry.ContractType) {
			tokenAdminRegistryAddr = common.HexToAddress(addr.Address)
		}
	}
	require.NotEqual(t, common.Address{}, routerAddr, "Router address should be set")
	require.NotEqual(t, common.Address{}, rmnAddr, "RMN address should be set")
	require.NotEqual(t, common.Address{}, tokenAdminRegistryAddr, "TokenAdminRegistry address should be set")

	// Deploy USDC token (BurnMintERC20)
	usdcTokenAddr, tx, _, err := burn_mint_erc20_bindings.DeployBurnMintERC20(
		chain.DeployerKey,
		chain.Client,
		"USD Coin",
		"USDC",
		6,             // decimals
		big.NewInt(0), // maxSupply
		big.NewInt(0), // pre-mint amount
	)
	require.NoError(t, err, "Failed to deploy USDC token")
	_, err = chain.Confirm(tx)
	require.NoError(t, err, "Failed to confirm USDC token deployment")

	// Deploy MockE2EUSDCTransmitter
	messageTransmitterAddr, tx, _, err := mock_usdc_token_transmitter.DeployMockE2EUSDCTransmitter(
		chain.DeployerKey,
		chain.Client,
		uint32(1),     // version (CCTP V2)
		uint32(1),     // localDomain
		usdcTokenAddr, // token
	)
	require.NoError(t, err, "Failed to deploy MockE2EUSDCTransmitter")
	_, err = chain.Confirm(tx)
	require.NoError(t, err, "Failed to confirm MockE2EUSDCTransmitter deployment")

	// Deploy MockE2EUSDCTokenMessenger
	tokenMessengerAddr, tx, _, err := mock_usdc_token_messenger.DeployMockE2EUSDCTokenMessenger(
		chain.DeployerKey,
		chain.Client,
		uint32(1),              // version (CCTP V2)
		messageTransmitterAddr, // transmitter
	)
	require.NoError(t, err, "Failed to deploy MockE2EUSDCTokenMessenger")
	_, err = chain.Confirm(tx)
	require.NoError(t, err, "Failed to confirm MockE2EUSDCTokenMessenger deployment")

	return cctpTestSetup{
		Router:             routerAddr,
		RMN:                rmnAddr,
		TokenAdminRegistry: tokenAdminRegistryAddr,
		USDCToken:          usdcTokenAddr,
		TokenMessenger:     tokenMessengerAddr,
		MessageTransmitter: messageTransmitterAddr,
	}
}

func TestCCTPChainAdapter_HomeToNonHomeChain(t *testing.T) {
	// Set up home chain (Ethereum Sepolia) and non-home chain
	homeChainSelector := chain_selectors.ETHEREUM_TESTNET_SEPOLIA.Selector
	nonHomeChainSelector := chain_selectors.ETHEREUM_TESTNET_SEPOLIA_ARBITRUM_1.Selector

	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{homeChainSelector, nonHomeChainSelector}),
	)
	require.NoError(t, err, "Failed to create environment")
	require.NotNil(t, e, "Environment should be created")

	// Create a single datastore that will be shared and updated
	ds := datastore.NewMemoryDataStore()
	e.DataStore = ds.Seal()

	// Set up both chains (this will update the shared datastore)
	homeSetup := setupCCTPTestEnvironment(t, e, homeChainSelector)
	nonHomeSetup := setupCCTPTestEnvironment(t, e, nonHomeChainSelector)

	homeChain := e.BlockChains.EVMChains()[homeChainSelector]
	nonHomeChain := e.BlockChains.EVMChains()[nonHomeChainSelector]

	// Deploy CCTP V1 pools on both chains for testing
	homeCCTPV1Pool, tx, _, err := v1_6_1_burn_mint_token_pool.DeployBurnMintTokenPool(homeChain.DeployerKey, homeChain.Client, homeSetup.USDCToken, 6, []common.Address{}, homeSetup.RMN, homeSetup.Router)
	require.NoError(t, err, "Failed to deploy CCTP V1 pool on home chain")
	_, err = homeChain.Confirm(tx)
	require.NoError(t, err, "Failed to confirm CCTP V1 pool deployment on home chain")

	nonHomeCCTPV1Pool, tx, _, err := v1_6_1_burn_mint_token_pool.DeployBurnMintTokenPool(nonHomeChain.DeployerKey, nonHomeChain.Client, nonHomeSetup.USDCToken, 6, []common.Address{}, nonHomeSetup.RMN, nonHomeSetup.Router)
	require.NoError(t, err, "Failed to deploy CCTP V1 pool on non-home chain")
	_, err = nonHomeChain.Confirm(tx)
	require.NoError(t, err, "Failed to confirm CCTP V1 pool deployment on non-home chain")

	// Add pools to datastore
	newDS := datastore.NewMemoryDataStore()

	err = newDS.Addresses().Add(datastore.AddressRef{
		ChainSelector: homeChainSelector,
		Address:       homeCCTPV1Pool.Hex(),
		Type:          datastore.ContractType("USDCTokenPool"),
		Version:       semver.MustParse("1.6.2"),
	})
	require.NoError(t, err, "Failed to add home CCTP V1 pool to datastore")

	err = newDS.Addresses().Add(datastore.AddressRef{
		ChainSelector: nonHomeChainSelector,
		Address:       nonHomeCCTPV1Pool.Hex(),
		Type:          datastore.ContractType("USDCTokenPool"),
		Version:       semver.MustParse("1.6.2"),
	})
	require.NoError(t, err, "Failed to add non-home CCTP V1 pool to datastore")

	err = newDS.Merge(e.DataStore)
	require.NoError(t, err, "Failed to merge datastore")
	e.DataStore = newDS.Seal()

	// Create CCTP chain registry and register adapter
	cctpChainRegistry := adapters.NewCCTPChainRegistry()
	adapter := &evm_adapters.CCTPChainAdapter{}
	cctpChainRegistry.RegisterCCTPChain("evm", adapter)

	// Create MCMS registry
	mcmsRegistry := changesets.GetRegistry()

	// Get CREATE2Factory addresses for both chains
	homeCreate2FactoryRef, err := contract_utils.MaybeDeployContract(e.OperationsBundle, create2_factory.Deploy, homeChain, contract_utils.DeployInput[create2_factory.ConstructorArgs]{
		TypeAndVersion: deployment.NewTypeAndVersion(create2_factory.ContractType, *semver.MustParse("1.7.0")),
		ChainSelector:  homeChainSelector,
		Args: create2_factory.ConstructorArgs{
			AllowList: []common.Address{homeChain.DeployerKey.From},
		},
	}, nil)
	require.NoError(t, err, "Failed to deploy CREATE2Factory on home chain")

	nonHomeCreate2FactoryRef, err := contract_utils.MaybeDeployContract(e.OperationsBundle, create2_factory.Deploy, nonHomeChain, contract_utils.DeployInput[create2_factory.ConstructorArgs]{
		TypeAndVersion: deployment.NewTypeAndVersion(create2_factory.ContractType, *semver.MustParse("1.7.0")),
		ChainSelector:  nonHomeChainSelector,
		Args: create2_factory.ConstructorArgs{
			AllowList: []common.Address{nonHomeChain.DeployerKey.From},
		},
	}, nil)
	require.NoError(t, err, "Failed to deploy CREATE2Factory on non-home chain")

	// Build DeployCCTPChainsConfig
	cfg := v1_7_0_changesets.DeployCCTPChainsConfig{
		Chains: map[uint64]v1_7_0_changesets.CCTPChainConfig{
			homeChainSelector: {
				TokenMessenger:   homeSetup.TokenMessenger.Hex(),
				USDCToken:        homeSetup.USDCToken.Hex(),
				DeployerContract: homeCreate2FactoryRef.Address,
				StorageLocations: []string{"https://test.chain.link.fake"},
				FeeAggregator:    common.HexToAddress("0x04").Hex(),
				FastFinalityBps:  100,
				RemoteChains: map[uint64]adapters.RemoteCCTPChainConfig{
					nonHomeChainSelector: {
						FeeUSDCents:         50,
						GasForVerification:  50_000,
						PayloadSizeBytes:    6*64 + 2*32,
						LockOrBurnMechanism: mechanismCCTPV2WithCCV,
						DomainIdentifier:    1,
						TokenTransferFeeConfig: tokens.TokenTransferFeeConfig{
							IsEnabled:                     true,
							DestGasOverhead:               200_000,
							DestBytesOverhead:             32,
							DefaultFinalityFeeUSDCents:    100,
							CustomFinalityFeeUSDCents:     200,
							DefaultFinalityTransferFeeBps: 100,
							CustomFinalityTransferFeeBps:  100,
						},
					},
				},
			},
			nonHomeChainSelector: {
				TokenMessenger:   nonHomeSetup.TokenMessenger.Hex(),
				USDCToken:        nonHomeSetup.USDCToken.Hex(),
				DeployerContract: nonHomeCreate2FactoryRef.Address,
				StorageLocations: []string{"https://test.chain.link.fake"},
				FeeAggregator:    common.HexToAddress("0x04").Hex(),
				FastFinalityBps:  100,
				RemoteChains: map[uint64]adapters.RemoteCCTPChainConfig{
					homeChainSelector: {
						FeeUSDCents:         50,
						GasForVerification:  50_000,
						PayloadSizeBytes:    6*64 + 2*32,
						LockOrBurnMechanism: mechanismCCTPV2WithCCV,
						DomainIdentifier:    1,
						TokenTransferFeeConfig: tokens.TokenTransferFeeConfig{
							IsEnabled:                     true,
							DestGasOverhead:               200_000,
							DestBytesOverhead:             32,
							DefaultFinalityFeeUSDCents:    100,
							CustomFinalityFeeUSDCents:     200,
							DefaultFinalityTransferFeeBps: 100,
							CustomFinalityTransferFeeBps:  100,
						},
					},
				},
			},
		},
	}

	// Apply the changeset
	changeset := v1_7_0_changesets.DeployCCTPChains(cctpChainRegistry, mcmsRegistry)
	output, err := changeset.Apply(*e, cfg)
	require.NoError(t, err, "Failed to apply DeployCCTPChains changeset")

	// Merge output addresses into datastore for state checks
	ds = datastore.NewMemoryDataStore()
	err = ds.Merge(e.DataStore)
	require.NoError(t, err)
	err = ds.Merge(output.DataStore.Seal())
	require.NoError(t, err)
	e.DataStore = ds.Seal()

	// Extract contract addresses from changeset output datastore
	outputAddresses, err := output.DataStore.Addresses().Fetch()
	require.NoError(t, err, "Failed to fetch addresses from changeset output")

	var homeCCTPTokenPoolAddr, homeCCTPMessageTransmitterProxyAddr, homeCCTPVerifierAddr, homeUSDCTokenPoolProxyAddr, homeCCTPVerifierResolverAddr common.Address
	var nonHomeCCTPTokenPoolAddr, nonHomeCCTPVerifierAddr, nonHomeUSDCTokenPoolProxyAddr common.Address

	for _, addr := range outputAddresses {
		switch addr.ChainSelector {
		case homeChainSelector:
			switch deployment.ContractType(addr.Type) {
			case cctp_through_ccv_token_pool.ContractType:
				homeCCTPTokenPoolAddr = common.HexToAddress(addr.Address)
			case cctp_message_transmitter_proxy.ContractType:
				homeCCTPMessageTransmitterProxyAddr = common.HexToAddress(addr.Address)
			case cctp_verifier.ContractType:
				homeCCTPVerifierAddr = common.HexToAddress(addr.Address)
			case usdc_token_pool_proxy.ContractType:
				homeUSDCTokenPoolProxyAddr = common.HexToAddress(addr.Address)
			case cctp_verifier.ResolverType:
				homeCCTPVerifierResolverAddr = common.HexToAddress(addr.Address)
			}
		case nonHomeChainSelector:
			switch deployment.ContractType(addr.Type) {
			case cctp_through_ccv_token_pool.ContractType:
				nonHomeCCTPTokenPoolAddr = common.HexToAddress(addr.Address)
			case cctp_verifier.ContractType:
				nonHomeCCTPVerifierAddr = common.HexToAddress(addr.Address)
			case usdc_token_pool_proxy.ContractType:
				nonHomeUSDCTokenPoolProxyAddr = common.HexToAddress(addr.Address)
			}
		}
	}

	// ===== State Checks for Home Chain =====

	// Check TokenAdminRegistry points to USDCTokenPoolProxy
	tokenConfigReport, err := operations.ExecuteOperation(
		testsetup.BundleWithFreshReporter(e.OperationsBundle),
		token_admin_registry.GetTokenConfig,
		homeChain,
		contract_utils.FunctionInput[common.Address]{
			ChainSelector: homeChainSelector,
			Address:       homeSetup.TokenAdminRegistry,
			Args:          homeSetup.USDCToken,
		},
	)
	require.NoError(t, err, "Failed to get token config from token admin registry on home chain")
	require.Equal(t, homeUSDCTokenPoolProxyAddr, tokenConfigReport.Output.TokenPool, "Token pool in registry should be the proxy on home chain")

	// Check CCTPTokenPool dynamic config on home chain
	homeCCTPTokenPool, err := cctp_through_ccv_token_pool_bindings.NewCCTPThroughCCVTokenPool(homeCCTPTokenPoolAddr, homeChain.Client)
	require.NoError(t, err, "Failed to instantiate CCTPTokenPool contract on home chain")

	homeDynamicConfig, err := homeCCTPTokenPool.GetDynamicConfig(nil)
	require.NoError(t, err, "Failed to get dynamic config from CCTPTokenPool on home chain")
	require.Equal(t, homeSetup.Router, homeDynamicConfig.Router, "Router should match on home chain")

	// Check CCTPTokenPool authorized callers on home chain
	homeAuthorizedCallers, err := homeCCTPTokenPool.GetAllAuthorizedCallers(nil)
	require.NoError(t, err, "Failed to get authorized callers from CCTPTokenPool on home chain")
	require.Contains(t, homeAuthorizedCallers, homeUSDCTokenPoolProxyAddr, "USDCTokenPoolProxy should be an authorized caller on home chain")

	// Check CCTPTokenPool token on home chain
	homeTokenAddr, err := homeCCTPTokenPool.GetToken(nil)
	require.NoError(t, err, "Failed to get token from CCTPTokenPool on home chain")
	require.Equal(t, homeSetup.USDCToken, homeTokenAddr, "Token address should match on home chain")

	// Check CCTPMessageTransmitterProxy authorized callers on home chain
	homeCCTPMessageTransmitterProxy, err := cctp_message_transmitter_proxy_bindings.NewCCTPMessageTransmitterProxy(homeCCTPMessageTransmitterProxyAddr, homeChain.Client)
	require.NoError(t, err, "Failed to instantiate CCTPMessageTransmitterProxy contract on home chain")
	homeAllowedCallers, err := homeCCTPMessageTransmitterProxy.GetAllAuthorizedCallers(nil)
	require.NoError(t, err, "Failed to get allowed callers from CCTPMessageTransmitterProxy on home chain")
	require.Contains(t, homeAllowedCallers, homeCCTPVerifierAddr, "CCTPVerifier should be an allowed caller on home chain")

	// Check CCTPVerifierResolver inbound implementation on home chain
	homeCCTPVerifierResolver, err := versioned_verifier_resolver_bindings.NewVersionedVerifierResolver(homeCCTPVerifierResolverAddr, homeChain.Client)
	require.NoError(t, err, "Failed to instantiate CCTPVerifierResolver contract on home chain")
	homeAllInboundImpls, err := homeCCTPVerifierResolver.GetAllInboundImplementations(nil)
	require.NoError(t, err, "Failed to get inbound implementations from CCTPVerifierResolver on home chain")
	require.Greater(t, len(homeAllInboundImpls), 0, "Should have at least one inbound implementation on home chain")
	homeFoundInbound := false
	for _, impl := range homeAllInboundImpls {
		if impl.Verifier == homeCCTPVerifierAddr {
			homeFoundInbound = true
			break
		}
	}
	require.True(t, homeFoundInbound, "CCTPVerifier should be registered as inbound implementation on home chain")

	// Check CCTPVerifierResolver outbound implementation on home chain
	homeAllOutboundImpls, err := homeCCTPVerifierResolver.GetAllOutboundImplementations(nil)
	require.NoError(t, err, "Failed to get outbound implementations from CCTPVerifierResolver on home chain")
	homeFoundOutbound := false
	for _, impl := range homeAllOutboundImpls {
		if impl.DestChainSelector == nonHomeChainSelector && impl.Verifier == homeCCTPVerifierAddr {
			homeFoundOutbound = true
			break
		}
	}
	require.True(t, homeFoundOutbound, "CCTPVerifier should be registered as outbound implementation for non-home chain on home chain")

	// Check USDCTokenPoolProxy lock or burn mechanism on home chain
	homeUSDCTokenPoolProxy, err := usdc_token_pool_proxy_bindings.NewUSDCTokenPoolProxy(homeUSDCTokenPoolProxyAddr, homeChain.Client)
	require.NoError(t, err, "Failed to instantiate USDCTokenPoolProxy contract on home chain")
	homeMechanism, err := homeUSDCTokenPoolProxy.GetLockOrBurnMechanism(nil, nonHomeChainSelector)
	require.NoError(t, err, "Failed to get lock or burn mechanism from USDCTokenPoolProxy on home chain")
	expectedMechanism, err := convertMechanismToUint8(mechanismCCTPV2WithCCV)
	require.NoError(t, err, "Failed to convert mechanism to uint8")
	require.Equal(t, expectedMechanism, uint8(homeMechanism), "Lock or burn mechanism should match on home chain")

	// Check USDCTokenPoolProxy pools on home chain
	homePools, err := homeUSDCTokenPoolProxy.GetPools(nil)
	require.NoError(t, err, "Failed to get pools from USDCTokenPoolProxy on home chain")
	require.Equal(t, homeCCTPV1Pool, homePools.CctpV1Pool, "CCTP V1 pool should match on home chain")

	// Check USDCTokenPoolProxy required CCVs on home chain
	homeRequiredCCVs, err := homeUSDCTokenPoolProxy.GetRequiredCCVs(nil, homeSetup.USDCToken, nonHomeChainSelector, big.NewInt(1e18), 1, []byte{}, 0)
	require.NoError(t, err, "Failed to get required CCVs from USDCTokenPoolProxy on home chain")
	require.Equal(t, []common.Address{homeCCTPVerifierResolverAddr}, homeRequiredCCVs, "Required CCVs should match on home chain")

	// Check CCTPVerifier static config on home chain
	homeCCTPVerifier, err := cctp_verifier_bindings.NewCCTPVerifier(homeCCTPVerifierAddr, homeChain.Client)
	require.NoError(t, err, "Failed to instantiate CCTPVerifier contract on home chain")
	homeStaticConfig, err := homeCCTPVerifier.GetStaticConfig(nil)
	require.NoError(t, err, "Failed to get static config from CCTPVerifier on home chain")
	require.Equal(t, homeSetup.TokenMessenger, homeStaticConfig.TokenMessenger, "TokenMessenger should match on home chain")
	require.Equal(t, homeCCTPMessageTransmitterProxyAddr, homeStaticConfig.MessageTransmitterProxy, "MessageTransmitterProxy should match on home chain")
	require.Equal(t, homeSetup.USDCToken, homeStaticConfig.UsdcToken, "USDCToken should match on home chain")

	// Check CCTPVerifier domain on home chain
	homeDomain, err := homeCCTPVerifier.GetDomain(nil, nonHomeChainSelector)
	require.NoError(t, err, "Failed to get domain from CCTPVerifier on home chain")
	require.Equal(t, uint32(1), homeDomain.DomainIdentifier, "Domain identifier should match on home chain")
	require.True(t, homeDomain.Enabled, "Domain should be enabled on home chain")

	// Check adapter methods work correctly
	homePoolAddress, err := adapter.PoolAddress(e.DataStore, e.BlockChains, homeChainSelector)
	require.NoError(t, err, "Failed to get pool address from adapter")
	require.Equal(t, homeUSDCTokenPoolProxyAddr.Bytes(), homePoolAddress, "Pool address should match")

	homeTokenAddress, err := adapter.TokenAddress(e.DataStore, e.BlockChains, homeChainSelector)
	require.NoError(t, err, "Failed to get token address from adapter")
	require.Equal(t, homeSetup.USDCToken.Bytes(), homeTokenAddress, "Token address should match")

	homeAllowedCallerOnDest, err := adapter.AllowedCallerOnDest(e.DataStore, e.BlockChains, homeChainSelector)
	require.NoError(t, err, "Failed to get allowed caller on dest from adapter")
	require.Equal(t, homeCCTPMessageTransmitterProxyAddr.Bytes(), homeAllowedCallerOnDest, "Allowed caller on dest should match")

	homeAllowedCallerOnSource, err := adapter.AllowedCallerOnSource(e.DataStore, e.BlockChains, homeChainSelector)
	require.NoError(t, err, "Failed to get allowed caller on source from adapter")
	require.Equal(t, homeCCTPVerifierAddr.Bytes(), homeAllowedCallerOnSource, "Allowed caller on source should match")

	homeMintRecipient, err := adapter.MintRecipientOnDest(e.DataStore, e.BlockChains, homeChainSelector)
	require.NoError(t, err, "Failed to get mint recipient from adapter")
	require.Equal(t, []byte{}, homeMintRecipient, "Mint recipient should be empty for EVM")

	// ===== State Checks for Non-Home Chain =====

	// Check TokenAdminRegistry points to USDCTokenPoolProxy on non-home chain
	nonHomeTokenConfigReport, err := operations.ExecuteOperation(
		testsetup.BundleWithFreshReporter(e.OperationsBundle),
		token_admin_registry.GetTokenConfig,
		nonHomeChain,
		contract_utils.FunctionInput[common.Address]{
			ChainSelector: nonHomeChainSelector,
			Address:       nonHomeSetup.TokenAdminRegistry,
			Args:          nonHomeSetup.USDCToken,
		},
	)
	require.NoError(t, err, "Failed to get token config from token admin registry on non-home chain")
	require.Equal(t, nonHomeUSDCTokenPoolProxyAddr, nonHomeTokenConfigReport.Output.TokenPool, "Token pool in registry should be the proxy on non-home chain")

	// Check CCTPTokenPool dynamic config on non-home chain
	nonHomeCCTPTokenPool, err := cctp_through_ccv_token_pool_bindings.NewCCTPThroughCCVTokenPool(nonHomeCCTPTokenPoolAddr, nonHomeChain.Client)
	require.NoError(t, err, "Failed to instantiate CCTPTokenPool contract on non-home chain")
	nonHomeDynamicConfig, err := nonHomeCCTPTokenPool.GetDynamicConfig(nil)
	require.NoError(t, err, "Failed to get dynamic config from CCTPTokenPool on non-home chain")
	require.Equal(t, nonHomeSetup.Router, nonHomeDynamicConfig.Router, "Router should match on non-home chain")

	// Check USDCTokenPoolProxy lock or burn mechanism on non-home chain
	nonHomeUSDCTokenPoolProxy, err := usdc_token_pool_proxy_bindings.NewUSDCTokenPoolProxy(nonHomeUSDCTokenPoolProxyAddr, nonHomeChain.Client)
	require.NoError(t, err, "Failed to instantiate USDCTokenPoolProxy contract on non-home chain")
	nonHomeMechanism, err := nonHomeUSDCTokenPoolProxy.GetLockOrBurnMechanism(nil, homeChainSelector)
	require.NoError(t, err, "Failed to get lock or burn mechanism from USDCTokenPoolProxy on non-home chain")
	require.Equal(t, expectedMechanism, uint8(nonHomeMechanism), "Lock or burn mechanism should match on non-home chain")

	// Check CCTPVerifier domain on non-home chain
	nonHomeCCTPVerifier, err := cctp_verifier_bindings.NewCCTPVerifier(nonHomeCCTPVerifierAddr, nonHomeChain.Client)
	require.NoError(t, err, "Failed to instantiate CCTPVerifier contract on non-home chain")
	nonHomeDomain, err := nonHomeCCTPVerifier.GetDomain(nil, homeChainSelector)
	require.NoError(t, err, "Failed to get domain from CCTPVerifier on non-home chain")
	require.Equal(t, uint32(1), nonHomeDomain.DomainIdentifier, "Domain identifier should match on non-home chain")
	require.True(t, nonHomeDomain.Enabled, "Domain should be enabled on non-home chain")

	// Check adapter methods work correctly for non-home chain
	nonHomePoolAddress, err := adapter.PoolAddress(e.DataStore, e.BlockChains, nonHomeChainSelector)
	require.NoError(t, err, "Failed to get pool address from adapter for non-home chain")
	require.Equal(t, nonHomeUSDCTokenPoolProxyAddr.Bytes(), nonHomePoolAddress, "Pool address should match for non-home chain")

	nonHomeTokenAddress, err := adapter.TokenAddress(e.DataStore, e.BlockChains, nonHomeChainSelector)
	require.NoError(t, err, "Failed to get token address from adapter for non-home chain")
	require.Equal(t, nonHomeSetup.USDCToken.Bytes(), nonHomeTokenAddress, "Token address should match for non-home chain")

	// Verify that addresses were created
	require.Greater(t, len(outputAddresses), 0, "Changeset should create addresses")
	require.NotEqual(t, common.Address{}, homeCCTPTokenPoolAddr, "Home chain CCTP token pool should be deployed")
	require.NotEqual(t, common.Address{}, nonHomeCCTPTokenPoolAddr, "Non-home chain CCTP token pool should be deployed")
}
