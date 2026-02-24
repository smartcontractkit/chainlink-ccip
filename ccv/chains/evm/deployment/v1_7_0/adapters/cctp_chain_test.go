package adapters_test

import (
	"math/big"
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	chain_selectors "github.com/smartcontractkit/chain-selectors"
	"github.com/stretchr/testify/require"

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
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/burn_mint_erc20"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/rmn_proxy"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/router"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/token_admin_registry"
	non_canonical_adapters "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_1/adapters"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_1/operations/burn_mint_with_lock_release_flag_token_pool"
	cctp_message_transmitter_proxy_v1_6_2 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_2/operations/cctp_message_transmitter_proxy"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_5/operations/usdc_token_pool_cctp_v2"
	v1_6_1_burn_mint_token_pool "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_1/burn_mint_token_pool"
	burn_mint_with_lock_release_flag_token_pool_bindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_1/burn_mint_with_lock_release_flag_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/adapters"
	v1_7_0_changesets "github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/changesets"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	burn_mint_erc20_bindings "github.com/smartcontractkit/chainlink-evm/gethwrappers/shared/generated/initial/burn_mint_erc20"
)

const (
	mechanismCCTPV2WithCCV = "CCTP_V2_WITH_CCV"
	mechanismLockRelease   = "LOCK_RELEASE"
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
	TokenMessengerV1   common.Address
	TokenMessengerV2   common.Address
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

	// Deploy MockE2EUSDCTransmitter (CCTP V1)
	messageTransmitterV1Addr, tx, _, err := mock_usdc_token_transmitter.DeployMockE2EUSDCTransmitter(
		chain.DeployerKey,
		chain.Client,
		uint32(0),     // version (CCTP V1)
		uint32(1),     // localDomain
		usdcTokenAddr, // token
	)
	require.NoError(t, err, "Failed to deploy MockE2EUSDCTransmitter")
	_, err = chain.Confirm(tx)
	require.NoError(t, err, "Failed to confirm MockE2EUSDCTransmitter deployment")

	// Deploy MockE2EUSDCTokenMessenger (CCTP V1)
	tokenMessengerV1Addr, tx, _, err := mock_usdc_token_messenger.DeployMockE2EUSDCTokenMessenger(
		chain.DeployerKey,
		chain.Client,
		uint32(0),                // version (CCTP V1)
		messageTransmitterV1Addr, // transmitter
	)
	require.NoError(t, err, "Failed to deploy MockE2EUSDCTokenMessenger")
	_, err = chain.Confirm(tx)
	require.NoError(t, err, "Failed to confirm MockE2EUSDCTokenMessenger deployment")

	// Deploy MockE2EUSDCTransmitter (CCTP V2)
	messageTransmitterV2Addr, tx, _, err := mock_usdc_token_transmitter.DeployMockE2EUSDCTransmitter(
		chain.DeployerKey,
		chain.Client,
		uint32(1),     // version (CCTP V2)
		uint32(1),     // localDomain
		usdcTokenAddr, // token
	)
	require.NoError(t, err, "Failed to deploy MockE2EUSDCTransmitter V2")
	_, err = chain.Confirm(tx)
	require.NoError(t, err, "Failed to confirm MockE2EUSDCTransmitter V2 deployment")

	// Deploy MockE2EUSDCTokenMessenger (CCTP V2)
	tokenMessengerV2Addr, tx, _, err := mock_usdc_token_messenger.DeployMockE2EUSDCTokenMessenger(
		chain.DeployerKey,
		chain.Client,
		uint32(1),                // version (CCTP V2)
		messageTransmitterV2Addr, // transmitter
	)
	require.NoError(t, err, "Failed to deploy MockE2EUSDCTokenMessenger V2")
	_, err = chain.Confirm(tx)
	require.NoError(t, err, "Failed to confirm MockE2EUSDCTokenMessenger V2 deployment")

	// Deploy and register required CCTPMessageTransmitterProxy v1.6.2 for CCTP V1 wiring.
	cctpV1ProxyRef, err := contract_utils.MaybeDeployContract(
		e.OperationsBundle,
		cctp_message_transmitter_proxy_v1_6_2.Deploy,
		chain,
		contract_utils.DeployInput[cctp_message_transmitter_proxy_v1_6_2.ConstructorArgs]{
			TypeAndVersion: deployment.NewTypeAndVersion(
				cctp_message_transmitter_proxy_v1_6_2.ContractType,
				*cctp_message_transmitter_proxy_v1_6_2.Version,
			),
			ChainSelector: chainSelector,
			Args: cctp_message_transmitter_proxy_v1_6_2.ConstructorArgs{
				TokenMessenger: tokenMessengerV1Addr,
			},
		},
		nil,
	)
	require.NoError(t, err, "Failed to deploy CCTPMessageTransmitterProxy v1.6.2")

	dsWithCCTPV1Proxy := datastore.NewMemoryDataStore()
	err = dsWithCCTPV1Proxy.Merge(e.DataStore)
	require.NoError(t, err)
	err = dsWithCCTPV1Proxy.Addresses().Add(cctpV1ProxyRef)
	require.NoError(t, err, "Failed to add CCTPMessageTransmitterProxy v1.6.2 to datastore")
	e.DataStore = dsWithCCTPV1Proxy.Seal()

	return cctpTestSetup{
		Router:             routerAddr,
		RMN:                rmnAddr,
		TokenAdminRegistry: tokenAdminRegistryAddr,
		USDCToken:          usdcTokenAddr,
		TokenMessengerV1:   tokenMessengerV1Addr,
		TokenMessengerV2:   tokenMessengerV2Addr,
		MessageTransmitter: messageTransmitterV2Addr,
	}
}

type nonCanonicalTestSetup struct {
	Router             common.Address
	RMN                common.Address
	TokenAdminRegistry common.Address
	USDCToken          common.Address
}

func setupNonCanonicalTestEnvironment(t *testing.T, e *deployment.Environment, chainSelector uint64) nonCanonicalTestSetup {
	chain := e.BlockChains.EVMChains()[chainSelector]

	rmnProxyRef, err := contract_utils.MaybeDeployContract(e.OperationsBundle, rmn_proxy.Deploy, chain, contract_utils.DeployInput[rmn_proxy.ConstructorArgs]{
		TypeAndVersion: deployment.NewTypeAndVersion(rmn_proxy.ContractType, *rmn_proxy.Version),
		ChainSelector:  chainSelector,
		Args:           rmn_proxy.ConstructorArgs{RMN: chain.DeployerKey.From},
	}, nil)
	require.NoError(t, err, "Failed to deploy RMN proxy")
	rmnAddr := common.HexToAddress(rmnProxyRef.Address)

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

	tarRef, err := contract_utils.MaybeDeployContract(e.OperationsBundle, token_admin_registry.Deploy, chain, contract_utils.DeployInput[token_admin_registry.ConstructorArgs]{
		TypeAndVersion: deployment.NewTypeAndVersion(token_admin_registry.ContractType, *token_admin_registry.Version),
		ChainSelector:  chainSelector,
		Args:           token_admin_registry.ConstructorArgs{},
	}, nil)
	require.NoError(t, err, "Failed to deploy TokenAdminRegistry")
	tarAddr := common.HexToAddress(tarRef.Address)

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

	ds := datastore.NewMemoryDataStore()
	if e.DataStore != nil {
		require.NoError(t, ds.Merge(e.DataStore))
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

	for _, sel := range []uint64{homeChainSelector, nonHomeChainSelector} {
		cctpV1ProxyRefs := e.DataStore.Addresses().Filter(
			datastore.AddressRefByChainSelector(sel),
			datastore.AddressRefByType(datastore.ContractType(cctp_message_transmitter_proxy_v1_6_2.ContractType)),
			datastore.AddressRefByVersion(cctp_message_transmitter_proxy_v1_6_2.Version),
		)
		require.Len(t, cctpV1ProxyRefs, 1, "Expected exactly one CCTPMessageTransmitterProxy v1.6.2 in datastore")
	}

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
				USDCType:         adapters.Canonical,
				TokenDecimals:    6,
				TokenMessengerV1: homeSetup.TokenMessengerV1.Hex(),
				TokenMessengerV2: homeSetup.TokenMessengerV2.Hex(),
				USDCToken:        homeSetup.USDCToken.Hex(),
				RegisteredPoolRef: datastore.AddressRef{
					Type:    datastore.ContractType(usdc_token_pool_proxy.ContractType),
					Version: usdc_token_pool_proxy.Version,
				},
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
				USDCType:         adapters.Canonical,
				TokenDecimals:    6,
				TokenMessengerV1: nonHomeSetup.TokenMessengerV1.Hex(),
				TokenMessengerV2: nonHomeSetup.TokenMessengerV2.Hex(),
				USDCToken:        nonHomeSetup.USDCToken.Hex(),
				RegisteredPoolRef: datastore.AddressRef{
					Type:    datastore.ContractType("USDCTokenPool"),
					Version: semver.MustParse("1.6.5"),
				},
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

	var homeCCTPTokenPoolAddr, homeCCTPV2TokenPoolAddr, homeCCTPMessageTransmitterProxyAddr, homeCCTPVerifierAddr, homeUSDCTokenPoolProxyAddr, homeCCTPVerifierResolverAddr common.Address
	var nonHomeCCTPTokenPoolAddr, nonHomeCCTPV2TokenPoolAddr, nonHomeCCTPMessageTransmitterProxyAddr, nonHomeCCTPVerifierAddr, nonHomeUSDCTokenPoolProxyAddr common.Address

	for _, addr := range outputAddresses {
		switch addr.ChainSelector {
		case homeChainSelector:
			switch deployment.ContractType(addr.Type) {
			case deployment.ContractType("USDCTokenPool"):
				if addr.Version != nil && addr.Version.String() == "1.6.5" {
					homeCCTPV1Pool = common.HexToAddress(addr.Address)
				}
			case cctp_through_ccv_token_pool.ContractType:
				homeCCTPTokenPoolAddr = common.HexToAddress(addr.Address)
			case usdc_token_pool_cctp_v2.ContractType:
				homeCCTPV2TokenPoolAddr = common.HexToAddress(addr.Address)
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
			case deployment.ContractType("USDCTokenPool"):
				if addr.Version != nil && addr.Version.String() == "1.6.5" {
					nonHomeCCTPV1Pool = common.HexToAddress(addr.Address)
				}
			case cctp_through_ccv_token_pool.ContractType:
				nonHomeCCTPTokenPoolAddr = common.HexToAddress(addr.Address)
			case usdc_token_pool_cctp_v2.ContractType:
				nonHomeCCTPV2TokenPoolAddr = common.HexToAddress(addr.Address)
			case cctp_message_transmitter_proxy.ContractType:
				nonHomeCCTPMessageTransmitterProxyAddr = common.HexToAddress(addr.Address)
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
	require.Contains(t, homeAllowedCallers, homeCCTPV2TokenPoolAddr, "CCTP V2 token pool should be an allowed caller on CCTPMessageTransmitterProxy on home chain")

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
	require.Equal(t, homeCCTPV2TokenPoolAddr, homePools.CctpV2Pool, "CCTP V2 pool should match on home chain")
	require.Equal(t, homeCCTPTokenPoolAddr, homePools.CctpV2PoolWithCCV, "CCTP V2-with-CCV pool should match on home chain")

	// Check CCTP V2 token pool authorized callers on home chain (proxy must be authorized)
	homeCCTPV2TokenPool, err := usdc_token_pool_cctp_v2.NewUSDCTokenPoolCCTPV2Contract(homeCCTPV2TokenPoolAddr, homeChain.Client)
	require.NoError(t, err, "Failed to instantiate CCTP V2 token pool contract on home chain")
	homeCCTPV2AuthorizedCallers, err := homeCCTPV2TokenPool.GetAllAuthorizedCallers(nil)
	require.NoError(t, err, "Failed to get authorized callers from CCTP V2 token pool on home chain")
	require.Contains(t, homeCCTPV2AuthorizedCallers, homeUSDCTokenPoolProxyAddr, "USDCTokenPoolProxy should be an authorized caller on CCTP V2 token pool on home chain")

	// Check USDCTokenPoolProxy required CCVs on home chain
	homeRequiredCCVs, err := homeUSDCTokenPoolProxy.GetRequiredCCVs(nil, homeSetup.USDCToken, nonHomeChainSelector, big.NewInt(1e18), 1, []byte{}, 0)
	require.NoError(t, err, "Failed to get required CCVs from USDCTokenPoolProxy on home chain")
	require.Equal(t, []common.Address{homeCCTPVerifierResolverAddr}, homeRequiredCCVs, "Required CCVs should match on home chain")

	// Check CCTPVerifier static config on home chain
	homeCCTPVerifier, err := cctp_verifier_bindings.NewCCTPVerifier(homeCCTPVerifierAddr, homeChain.Client)
	require.NoError(t, err, "Failed to instantiate CCTPVerifier contract on home chain")
	homeStaticConfig, err := homeCCTPVerifier.GetStaticConfig(nil)
	require.NoError(t, err, "Failed to get static config from CCTPVerifier on home chain")
	require.Equal(t, homeSetup.TokenMessengerV2, homeStaticConfig.TokenMessenger, "TokenMessenger should match on home chain")
	require.Equal(t, homeCCTPMessageTransmitterProxyAddr, homeStaticConfig.MessageTransmitterProxy, "MessageTransmitterProxy should match on home chain")
	require.Equal(t, homeSetup.USDCToken, homeStaticConfig.UsdcToken, "USDCToken should match on home chain")

	// Check CCTPVerifier domain on home chain
	homeDomain, err := homeCCTPVerifier.GetDomain(nil, nonHomeChainSelector)
	require.NoError(t, err, "Failed to get domain from CCTPVerifier on home chain")
	require.Equal(t, uint32(1), homeDomain.DomainIdentifier, "Domain identifier should match on home chain")
	require.True(t, homeDomain.Enabled, "Domain should be enabled on home chain")
	require.Equal(t, common.LeftPadBytes(nonHomeCCTPMessageTransmitterProxyAddr.Bytes(), 32), homeDomain.AllowedCallerOnDest[:], "AllowedCallerOnDest should be non-home message transmitter proxy on home chain")
	require.Equal(t, common.LeftPadBytes(nonHomeCCTPVerifierAddr.Bytes(), 32), homeDomain.AllowedCallerOnSource[:], "AllowedCallerOnSource should be home CCTPVerifier on home chain")
	require.Equal(t, [32]byte{}, homeDomain.MintRecipientOnDest, "MintRecipientOnDest should be zero for EVM on home chain")

	// Check CCTPVerifier remote chain config on home chain
	homeVerifierRemoteChainConfig, err := homeCCTPVerifier.GetRemoteChainConfig(nil, nonHomeChainSelector)
	require.NoError(t, err, "Failed to get remote chain config from CCTPVerifier on home chain")
	require.Equal(t, homeSetup.Router, homeVerifierRemoteChainConfig.Router, "CCTPVerifier remote chain config Router should match on home chain")

	// Check CCTP V2 token pool domain on home chain
	homeCCTPV2Domain, err := homeCCTPV2TokenPool.GetDomain(nil, nonHomeChainSelector)
	require.NoError(t, err, "Failed to get domain from CCTP V2 token pool on home chain")
	require.Equal(t, uint32(1), homeCCTPV2Domain.DomainIdentifier, "CCTP V2 pool domain identifier should match on home chain")
	require.True(t, homeCCTPV2Domain.Enabled, "CCTP V2 pool domain should be enabled on home chain")
	require.Equal(t, common.LeftPadBytes(nonHomeCCTPMessageTransmitterProxyAddr.Bytes(), 32), homeCCTPV2Domain.AllowedCaller[:], "CCTP V2 pool AllowedCaller should be non-home message transmitter proxy on home chain")
	require.Equal(t, [32]byte{}, homeCCTPV2Domain.MintRecipient, "CCTP V2 pool MintRecipient should be zero for EVM on home chain")

	// Check CCTP V2 token pool remote chain config on home chain (ConfigureTokenPoolForRemoteChain)
	homeCCTPV2SupportedChains, err := homeCCTPV2TokenPool.GetSupportedChains(nil)
	require.NoError(t, err, "Failed to get supported chains from CCTP V2 token pool on home chain")
	require.Contains(t, homeCCTPV2SupportedChains, nonHomeChainSelector, "CCTP V2 pool should support non-home chain on home chain")
	homeCCTPV2RemoteToken, err := homeCCTPV2TokenPool.GetRemoteToken(nil, nonHomeChainSelector)
	require.NoError(t, err, "Failed to get remote token from CCTP V2 token pool on home chain")
	require.Equal(t, common.LeftPadBytes(nonHomeSetup.USDCToken.Bytes(), 32), homeCCTPV2RemoteToken, "CCTP V2 pool remote token should be non-home USDC on home chain")
	homeCCTPV2RemotePools, err := homeCCTPV2TokenPool.GetRemotePools(nil, nonHomeChainSelector)
	require.NoError(t, err, "Failed to get remote pools from CCTP V2 token pool on home chain")
	require.Contains(t, homeCCTPV2RemotePools, common.LeftPadBytes(nonHomeCCTPV1Pool.Bytes(), 32), "CCTP V2 pool should have non-home CCTP V1 pool as remote pool on home chain")

	// Check CCTP V1 token pool remote chain config on home chain.
	// This lane is configured as CCTP_V2_WITH_CCV, so CCTP V1 should not be configured.
	homeCCTPV1PoolBinding, err := v1_6_1_burn_mint_token_pool.NewBurnMintTokenPool(homeCCTPV1Pool, homeChain.Client)
	require.NoError(t, err, "Failed to instantiate CCTP V1 token pool contract on home chain")
	homeCCTPV1SupportedChains, err := homeCCTPV1PoolBinding.GetSupportedChains(nil)
	require.NoError(t, err, "Failed to get supported chains from CCTP V1 token pool on home chain")
	require.NotContains(t, homeCCTPV1SupportedChains, nonHomeChainSelector, "CCTP V1 pool should not be configured for CCTP_V2_WITH_CCV on home chain")

	// Check USDCTokenPoolProxy fee aggregator on home chain
	homeFeeAggregator, err := homeUSDCTokenPoolProxy.GetFeeAggregator(nil)
	require.NoError(t, err, "Failed to get fee aggregator from USDCTokenPoolProxy on home chain")
	require.Equal(t, common.HexToAddress("0x04"), homeFeeAggregator, "Fee aggregator should match on home chain")

	// Check adapter methods work correctly
	homePoolAddress, err := adapter.PoolAddress(e.DataStore, e.BlockChains, homeChainSelector, datastore.AddressRef{
		Type:    datastore.ContractType(usdc_token_pool_proxy.ContractType),
		Version: usdc_token_pool_proxy.Version,
	})
	require.NoError(t, err, "Failed to get pool address from adapter")
	require.Equal(t, homeUSDCTokenPoolProxyAddr.Bytes(), homePoolAddress, "Pool address should match")

	homeTokenAddress, err := adapter.TokenAddress(e.DataStore, e.BlockChains, homeChainSelector)
	require.NoError(t, err, "Failed to get token address from adapter")
	require.Equal(t, homeSetup.USDCToken.Bytes(), homeTokenAddress, "Token address should match")

	homeAllowedCallerOnDest, err := adapter.CCTPV2AllowedCallerOnDest(e.DataStore, e.BlockChains, homeChainSelector)
	require.NoError(t, err, "Failed to get allowed caller on dest from adapter")
	require.Equal(t, homeCCTPMessageTransmitterProxyAddr.Bytes(), homeAllowedCallerOnDest, "Allowed caller on dest should match")

	homeAllowedCallerOnSource, err := adapter.AllowedCallerOnSource(e.DataStore, e.BlockChains, homeChainSelector)
	require.NoError(t, err, "Failed to get allowed caller on source from adapter")
	require.Equal(t, homeCCTPVerifierAddr.Bytes(), homeAllowedCallerOnSource, "Allowed caller on source should match")

	homeMintRecipient, err := adapter.MintRecipientOnDest(e.DataStore, e.BlockChains, homeChainSelector)
	require.NoError(t, err, "Failed to get mint recipient from adapter")
	require.Equal(t, []byte{}, homeMintRecipient, "Mint recipient should be empty for EVM")

	// ===== State Checks for Non-Home Chain =====

	// Check TokenAdminRegistry points to USDCTokenPool on non-home chain
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
	require.Equal(t, nonHomeCCTPV1Pool, nonHomeTokenConfigReport.Output.TokenPool, "Token pool in registry should be the CCTP V1 pool on non-home chain")

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

	// Check USDCTokenPoolProxy pools on non-home chain (CCTP V2 and V2-with-CCV addresses)
	nonHomePools, err := nonHomeUSDCTokenPoolProxy.GetPools(nil)
	require.NoError(t, err, "Failed to get pools from USDCTokenPoolProxy on non-home chain")
	require.Equal(t, nonHomeCCTPV2TokenPoolAddr, nonHomePools.CctpV2Pool, "CCTP V2 pool should match on non-home chain")
	require.Equal(t, nonHomeCCTPTokenPoolAddr, nonHomePools.CctpV2PoolWithCCV, "CCTP V2-with-CCV pool should match on non-home chain")

	// Check CCTPVerifier domain on non-home chain
	nonHomeCCTPVerifier, err := cctp_verifier_bindings.NewCCTPVerifier(nonHomeCCTPVerifierAddr, nonHomeChain.Client)
	require.NoError(t, err, "Failed to instantiate CCTPVerifier contract on non-home chain")
	nonHomeDomain, err := nonHomeCCTPVerifier.GetDomain(nil, homeChainSelector)
	require.NoError(t, err, "Failed to get domain from CCTPVerifier on non-home chain")
	require.Equal(t, uint32(1), nonHomeDomain.DomainIdentifier, "Domain identifier should match on non-home chain")
	require.True(t, nonHomeDomain.Enabled, "Domain should be enabled on non-home chain")
	require.Equal(t, common.LeftPadBytes(homeCCTPMessageTransmitterProxyAddr.Bytes(), 32), nonHomeDomain.AllowedCallerOnDest[:], "AllowedCallerOnDest should be home message transmitter proxy on non-home chain")
	require.Equal(t, common.LeftPadBytes(homeCCTPVerifierAddr.Bytes(), 32), nonHomeDomain.AllowedCallerOnSource[:], "AllowedCallerOnSource should be home CCTPVerifier on non-home chain")
	require.Equal(t, [32]byte{}, nonHomeDomain.MintRecipientOnDest, "MintRecipientOnDest should be zero for EVM on non-home chain")

	// Check CCTPVerifier remote chain config on non-home chain
	nonHomeVerifierRemoteChainConfig, err := nonHomeCCTPVerifier.GetRemoteChainConfig(nil, homeChainSelector)
	require.NoError(t, err, "Failed to get remote chain config from CCTPVerifier on non-home chain")
	require.Equal(t, nonHomeSetup.Router, nonHomeVerifierRemoteChainConfig.Router, "CCTPVerifier remote chain config Router should match on non-home chain")

	// Check CCTP V2 token pool domain on non-home chain
	nonHomeCCTPV2TokenPool, err := usdc_token_pool_cctp_v2.NewUSDCTokenPoolCCTPV2Contract(nonHomeCCTPV2TokenPoolAddr, nonHomeChain.Client)
	require.NoError(t, err, "Failed to instantiate CCTP V2 token pool contract on non-home chain")
	nonHomeCCTPV2Domain, err := nonHomeCCTPV2TokenPool.GetDomain(nil, homeChainSelector)
	require.NoError(t, err, "Failed to get domain from CCTP V2 token pool on non-home chain")
	require.Equal(t, uint32(1), nonHomeCCTPV2Domain.DomainIdentifier, "CCTP V2 pool domain identifier should match on non-home chain")
	require.True(t, nonHomeCCTPV2Domain.Enabled, "CCTP V2 pool domain should be enabled on non-home chain")
	require.Equal(t, common.LeftPadBytes(homeCCTPMessageTransmitterProxyAddr.Bytes(), 32), nonHomeCCTPV2Domain.AllowedCaller[:], "CCTP V2 pool AllowedCaller should be home message transmitter proxy on non-home chain")
	require.Equal(t, [32]byte{}, nonHomeCCTPV2Domain.MintRecipient, "CCTP V2 pool MintRecipient should be zero for EVM on non-home chain")

	// Check CCTP V2 token pool remote chain config on non-home chain
	nonHomeCCTPV2SupportedChains, err := nonHomeCCTPV2TokenPool.GetSupportedChains(nil)
	require.NoError(t, err, "Failed to get supported chains from CCTP V2 token pool on non-home chain")
	require.Contains(t, nonHomeCCTPV2SupportedChains, homeChainSelector, "CCTP V2 pool should support home chain on non-home chain")
	nonHomeCCTPV2RemoteToken, err := nonHomeCCTPV2TokenPool.GetRemoteToken(nil, homeChainSelector)
	require.NoError(t, err, "Failed to get remote token from CCTP V2 token pool on non-home chain")
	require.Equal(t, common.LeftPadBytes(homeSetup.USDCToken.Bytes(), 32), nonHomeCCTPV2RemoteToken, "CCTP V2 pool remote token should be home USDC on non-home chain")
	nonHomeCCTPV2RemotePools, err := nonHomeCCTPV2TokenPool.GetRemotePools(nil, homeChainSelector)
	require.NoError(t, err, "Failed to get remote pools from CCTP V2 token pool on non-home chain")
	require.Contains(t, nonHomeCCTPV2RemotePools, common.LeftPadBytes(homeUSDCTokenPoolProxyAddr.Bytes(), 32), "CCTP V2 pool should have home proxy as remote pool on non-home chain")

	// Check CCTP V1 token pool remote chain config on non-home chain.
	// This lane is configured as CCTP_V2_WITH_CCV, so CCTP V1 should not be configured.
	nonHomeCCTPV1PoolBinding, err := v1_6_1_burn_mint_token_pool.NewBurnMintTokenPool(nonHomeCCTPV1Pool, nonHomeChain.Client)
	require.NoError(t, err, "Failed to instantiate CCTP V1 token pool contract on non-home chain")
	nonHomeCCTPV1SupportedChains, err := nonHomeCCTPV1PoolBinding.GetSupportedChains(nil)
	require.NoError(t, err, "Failed to get supported chains from CCTP V1 token pool on non-home chain")
	require.NotContains(t, nonHomeCCTPV1SupportedChains, homeChainSelector, "CCTP V1 pool should not be configured for CCTP_V2_WITH_CCV on non-home chain")

	// Check USDCTokenPoolProxy fee aggregator on non-home chain
	nonHomeFeeAggregator, err := nonHomeUSDCTokenPoolProxy.GetFeeAggregator(nil)
	require.NoError(t, err, "Failed to get fee aggregator from USDCTokenPoolProxy on non-home chain")
	require.Equal(t, common.HexToAddress("0x04"), nonHomeFeeAggregator, "Fee aggregator should match on non-home chain")

	// Check adapter methods work correctly for non-home chain
	nonHomePoolAddress, err := adapter.PoolAddress(e.DataStore, e.BlockChains, nonHomeChainSelector, datastore.AddressRef{
		Type:    datastore.ContractType("USDCTokenPool"),
		Version: semver.MustParse("1.6.5"),
	})
	require.NoError(t, err, "Failed to get pool address from adapter for non-home chain")
	require.Equal(t, nonHomeCCTPV1Pool.Bytes(), nonHomePoolAddress, "Pool address should match for non-home chain")

	nonHomeTokenAddress, err := adapter.TokenAddress(e.DataStore, e.BlockChains, nonHomeChainSelector)
	require.NoError(t, err, "Failed to get token address from adapter for non-home chain")
	require.Equal(t, nonHomeSetup.USDCToken.Bytes(), nonHomeTokenAddress, "Token address should match for non-home chain")

	// Verify that addresses were created
	require.Greater(t, len(outputAddresses), 0, "Changeset should create addresses")
	require.NotEqual(t, common.Address{}, homeCCTPTokenPoolAddr, "Home chain CCTP token pool should be deployed")
	require.NotEqual(t, common.Address{}, homeCCTPV2TokenPoolAddr, "Home chain CCTP V2 token pool should be deployed")
	require.NotEqual(t, common.Address{}, nonHomeCCTPTokenPoolAddr, "Non-home chain CCTP token pool should be deployed")
	require.NotEqual(t, common.Address{}, nonHomeCCTPV2TokenPoolAddr, "Non-home chain CCTP V2 token pool should be deployed")
}

// remoteChainConfigForNonCanonical returns a remote chain config suitable for lanes to/from non-canonical chains.
// Rate limiter Capacity/Rate must be non-nil for ABI packing in token pool configuration.
func remoteChainConfigForNonCanonical() adapters.RemoteCCTPChainConfig {
	return adapters.RemoteCCTPChainConfig{
		DefaultFinalityInboundRateLimiterConfig:  tokens.RateLimiterConfig{IsEnabled: false, Capacity: big.NewInt(0), Rate: big.NewInt(0)},
		DefaultFinalityOutboundRateLimiterConfig: tokens.RateLimiterConfig{IsEnabled: false, Capacity: big.NewInt(0), Rate: big.NewInt(0)},
		CustomFinalityInboundRateLimiterConfig:   tokens.RateLimiterConfig{IsEnabled: false, Capacity: big.NewInt(0), Rate: big.NewInt(0)},
		CustomFinalityOutboundRateLimiterConfig:  tokens.RateLimiterConfig{IsEnabled: false, Capacity: big.NewInt(0), Rate: big.NewInt(0)},
		TokenTransferFeeConfig: tokens.TokenTransferFeeConfig{
			IsEnabled:                     true,
			DestGasOverhead:               200_000,
			DestBytesOverhead:             32,
			DefaultFinalityFeeUSDCents:    100,
			CustomFinalityFeeUSDCents:     200,
			DefaultFinalityTransferFeeBps: 100,
			CustomFinalityTransferFeeBps:  100,
		},
	}
}

// TestCCTPChainAdapter_CanonicalToNonCanonicalChain connects a canonical 1.7.0 chain with a non-canonical 1.6.1 chain
// via the DeployCCTPChains changeset: canonical chain gets proxy + CCTP verifier/pools; non-canonical gets
// BurnMintWithLockReleaseFlagTokenPool; both sides see each other as remotes (LOCK_RELEASE on canonical side).
func TestCCTPChainAdapter_CanonicalToNonCanonicalChain(t *testing.T) {
	canonicalChainSelector := chain_selectors.ETHEREUM_TESTNET_SEPOLIA.Selector
	nonCanonicalChainSelector := chain_selectors.ETHEREUM_TESTNET_SEPOLIA_ARBITRUM_1.Selector

	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{canonicalChainSelector, nonCanonicalChainSelector}),
	)
	require.NoError(t, err, "Failed to create environment")
	require.NotNil(t, e, "Environment should be created")

	ds := datastore.NewMemoryDataStore()
	e.DataStore = ds.Seal()

	// Setup canonical chain (CCTP mocks, chain contracts, CCTP V1 proxy; CCTP V1 pool added below)
	canonicalSetup := setupCCTPTestEnvironment(t, e, canonicalChainSelector)
	canonicalChain := e.BlockChains.EVMChains()[canonicalChainSelector]

	// Setup non-canonical chain (RMN, Router, TAR, USDC only)
	nonCanonicalSetup := setupNonCanonicalTestEnvironment(t, e, nonCanonicalChainSelector)
	nonCanonicalChain := e.BlockChains.EVMChains()[nonCanonicalChainSelector]

	// Deploy CCTP V1 pool on canonical chain and add to datastore (required for proxy's CCTP V1 pool ref)
	homeCCTPV1Pool, tx, _, err := v1_6_1_burn_mint_token_pool.DeployBurnMintTokenPool(
		canonicalChain.DeployerKey, canonicalChain.Client,
		canonicalSetup.USDCToken, 6, []common.Address{}, canonicalSetup.RMN, canonicalSetup.Router,
	)
	require.NoError(t, err, "Failed to deploy CCTP V1 pool on canonical chain")
	_, err = canonicalChain.Confirm(tx)
	require.NoError(t, err, "Failed to confirm CCTP V1 pool on canonical chain")

	newDS := datastore.NewMemoryDataStore()
	require.NoError(t, newDS.Merge(e.DataStore))
	require.NoError(t, newDS.Addresses().Add(datastore.AddressRef{
		ChainSelector: canonicalChainSelector,
		Address:       homeCCTPV1Pool.Hex(),
		Type:          datastore.ContractType("USDCTokenPool"),
		Version:       semver.MustParse("1.6.2"),
	}))
	e.DataStore = newDS.Seal()

	// CREATE2Factory for canonical chain was already deployed and added by setupCCTPTestEnvironment
	canonicalCreate2FactoryRefs := e.DataStore.Addresses().Filter(
		datastore.AddressRefByChainSelector(canonicalChainSelector),
		datastore.AddressRefByType(datastore.ContractType(create2_factory.ContractType)),
		datastore.AddressRefByVersion(semver.MustParse("1.7.0")),
	)
	require.Len(t, canonicalCreate2FactoryRefs, 1, "Expected exactly one CREATE2Factory for canonical chain")
	canonicalCreate2FactoryRef := canonicalCreate2FactoryRefs[0]

	// Register both adapters: canonical (1.7.0) and non-canonical (1.6.1)
	cctpChainRegistry := adapters.NewCCTPChainRegistry()
	cctpChainRegistry.RegisterCCTPChain("evm", &evm_adapters.CCTPChainAdapter{})
	cctpChainRegistry.RegisterCCTPChain("evm", &non_canonical_adapters.NonCanonicalUSDCChainAdapter{})

	mcmsRegistry := changesets.GetRegistry()

	cfg := v1_7_0_changesets.DeployCCTPChainsConfig{
		Chains: map[uint64]v1_7_0_changesets.CCTPChainConfig{
			canonicalChainSelector: {
				USDCType:         adapters.Canonical,
				TokenDecimals:    6,
				TokenMessengerV1: canonicalSetup.TokenMessengerV1.Hex(),
				TokenMessengerV2: canonicalSetup.TokenMessengerV2.Hex(),
				USDCToken:        canonicalSetup.USDCToken.Hex(),
				RegisteredPoolRef: datastore.AddressRef{
					Type:    datastore.ContractType(usdc_token_pool_proxy.ContractType),
					Version: usdc_token_pool_proxy.Version,
				},
				DeployerContract: canonicalCreate2FactoryRef.Address,
				StorageLocations: []string{"https://test.chain.link.fake"},
				FeeAggregator:    common.HexToAddress("0x04").Hex(),
				FastFinalityBps:  100,
				RemoteChains: map[uint64]adapters.RemoteCCTPChainConfig{
					nonCanonicalChainSelector: {
						LockOrBurnMechanism: mechanismLockRelease,
						TokenTransferFeeConfig: tokens.TokenTransferFeeConfig{
							IsEnabled:                     true,
							DestGasOverhead:               200_000,
							DestBytesOverhead:             32,
							DefaultFinalityFeeUSDCents:    100,
							CustomFinalityFeeUSDCents:     200,
							DefaultFinalityTransferFeeBps: 100,
							CustomFinalityTransferFeeBps:  100,
						},
						DefaultFinalityInboundRateLimiterConfig:  tokens.RateLimiterConfig{IsEnabled: false, Capacity: big.NewInt(0), Rate: big.NewInt(0)},
						DefaultFinalityOutboundRateLimiterConfig: tokens.RateLimiterConfig{IsEnabled: false, Capacity: big.NewInt(0), Rate: big.NewInt(0)},
						CustomFinalityInboundRateLimiterConfig:   tokens.RateLimiterConfig{IsEnabled: false, Capacity: big.NewInt(0), Rate: big.NewInt(0)},
						CustomFinalityOutboundRateLimiterConfig:  tokens.RateLimiterConfig{IsEnabled: false, Capacity: big.NewInt(0), Rate: big.NewInt(0)},
					},
				},
			},
			nonCanonicalChainSelector: {
				USDCType:         adapters.NonCanonical,
				TokenDecimals:    6,
				TokenMessengerV1: "",
				TokenMessengerV2: common.Address{}.Hex(),
				USDCToken:        nonCanonicalSetup.USDCToken.Hex(),
				DeployerContract: nonCanonicalSetup.Router.Hex(),
				StorageLocations: []string{"https://test.chain.link.fake"},
				FeeAggregator:    common.HexToAddress("0x04").Hex(),
				FastFinalityBps:  100,
				RegisteredPoolRef: datastore.AddressRef{
					Type:    datastore.ContractType(burn_mint_with_lock_release_flag_token_pool.ContractType),
					Version: burn_mint_with_lock_release_flag_token_pool.Version,
				},
				RemoteChains: map[uint64]adapters.RemoteCCTPChainConfig{
					canonicalChainSelector: remoteChainConfigForNonCanonical(),
				},
			},
		},
	}

	changeset := v1_7_0_changesets.DeployCCTPChains(cctpChainRegistry, mcmsRegistry)
	output, err := changeset.Apply(*e, cfg)
	require.NoError(t, err, "Failed to apply DeployCCTPChains changeset")

	ds = datastore.NewMemoryDataStore()
	require.NoError(t, ds.Merge(e.DataStore))
	require.NoError(t, ds.Merge(output.DataStore.Seal()))
	e.DataStore = ds.Seal()

	outputAddresses, err := output.DataStore.Addresses().Fetch()
	require.NoError(t, err, "Failed to fetch addresses from changeset output")

	var canonicalProxyAddr, nonCanonicalPoolAddr common.Address
	for _, addr := range outputAddresses {
		switch addr.ChainSelector {
		case canonicalChainSelector:
			if deployment.ContractType(addr.Type) == usdc_token_pool_proxy.ContractType {
				canonicalProxyAddr = common.HexToAddress(addr.Address)
			}
		case nonCanonicalChainSelector:
			if addr.Type == datastore.ContractType(burn_mint_with_lock_release_flag_token_pool.ContractType) {
				nonCanonicalPoolAddr = common.HexToAddress(addr.Address)
			}
		}
	}

	require.NotEqual(t, common.Address{}, canonicalProxyAddr, "Canonical chain should have USDCTokenPoolProxy")
	require.NotEqual(t, common.Address{}, nonCanonicalPoolAddr, "Non-canonical chain should have BurnMintWithLockReleaseFlagTokenPool")

	// Canonical: TokenAdminRegistry points to proxy
	tokenConfigReport, err := operations.ExecuteOperation(
		testsetup.BundleWithFreshReporter(e.OperationsBundle),
		token_admin_registry.GetTokenConfig,
		canonicalChain,
		contract_utils.FunctionInput[common.Address]{
			ChainSelector: canonicalChainSelector,
			Address:       canonicalSetup.TokenAdminRegistry,
			Args:          canonicalSetup.USDCToken,
		},
	)
	require.NoError(t, err, "Failed to get token config on canonical chain")
	require.Equal(t, canonicalProxyAddr, tokenConfigReport.Output.TokenPool, "Canonical TAR should point to proxy")

	// Canonical: proxy reports LOCK_RELEASE for non-canonical remote
	canonicalProxy, err := usdc_token_pool_proxy_bindings.NewUSDCTokenPoolProxy(canonicalProxyAddr, canonicalChain.Client)
	require.NoError(t, err, "Failed to bind canonical proxy")
	mechanism, err := canonicalProxy.GetLockOrBurnMechanism(nil, nonCanonicalChainSelector)
	require.NoError(t, err, "Failed to get lock or burn mechanism for non-canonical remote")
	expectedLockRelease, _ := convertMechanismToUint8(mechanismLockRelease)
	require.Equal(t, expectedLockRelease, uint8(mechanism), "Canonical proxy should use LOCK_RELEASE for non-canonical remote")

	// Non-canonical: TokenAdminRegistry points to BurnMintWithLockReleaseFlagTokenPool
	nonCanonicalTokenConfigReport, err := operations.ExecuteOperation(
		testsetup.BundleWithFreshReporter(e.OperationsBundle),
		token_admin_registry.GetTokenConfig,
		nonCanonicalChain,
		contract_utils.FunctionInput[common.Address]{
			ChainSelector: nonCanonicalChainSelector,
			Address:       nonCanonicalSetup.TokenAdminRegistry,
			Args:          nonCanonicalSetup.USDCToken,
		},
	)
	require.NoError(t, err, "Failed to get token config on non-canonical chain")
	require.Equal(t, nonCanonicalPoolAddr, nonCanonicalTokenConfigReport.Output.TokenPool, "Non-canonical TAR should point to lock-release pool")

	// Non-canonical: pool supports canonical chain; remote token and remote pool match canonical USDC and proxy
	nonCanonicalPool, err := burn_mint_with_lock_release_flag_token_pool_bindings.NewBurnMintWithLockReleaseFlagTokenPool(nonCanonicalPoolAddr, nonCanonicalChain.Client)
	require.NoError(t, err, "Failed to bind non-canonical pool")
	supported, err := nonCanonicalPool.GetSupportedChains(nil)
	require.NoError(t, err, "Failed to get supported chains from non-canonical pool")
	require.Contains(t, supported, canonicalChainSelector, "Non-canonical pool should support canonical chain")
	remoteToken, err := nonCanonicalPool.GetRemoteToken(nil, canonicalChainSelector)
	require.NoError(t, err, "Failed to get remote token from non-canonical pool")
	require.Equal(t, common.LeftPadBytes(canonicalSetup.USDCToken.Bytes(), 32), remoteToken, "Non-canonical pool remote token should be canonical USDC")
	remotePools, err := nonCanonicalPool.GetRemotePools(nil, canonicalChainSelector)
	require.NoError(t, err, "Failed to get remote pools from non-canonical pool")
	require.Contains(t, remotePools, common.LeftPadBytes(canonicalProxyAddr.Bytes(), 32), "Non-canonical pool should have canonical proxy as remote pool")

	// Adapter methods: canonical chain uses 1.7.0 adapter
	canonicalAdapter, _ := cctpChainRegistry.GetCCTPChain("evm", adapters.Canonical)
	poolAddr, err := canonicalAdapter.PoolAddress(e.DataStore, e.BlockChains, canonicalChainSelector, datastore.AddressRef{
		Type:    datastore.ContractType(usdc_token_pool_proxy.ContractType),
		Version: usdc_token_pool_proxy.Version,
	})
	require.NoError(t, err, "Failed to get canonical pool address from adapter")
	require.Equal(t, canonicalProxyAddr.Bytes(), poolAddr, "Canonical adapter pool address should match proxy")

	// Adapter methods: non-canonical chain uses 1.6.1 adapter
	nonCanonicalAdapter, _ := cctpChainRegistry.GetCCTPChain("evm", adapters.NonCanonical)
	poolAddrNC, err := nonCanonicalAdapter.PoolAddress(e.DataStore, e.BlockChains, nonCanonicalChainSelector, datastore.AddressRef{
		Type:    datastore.ContractType(burn_mint_with_lock_release_flag_token_pool.ContractType),
		Version: burn_mint_with_lock_release_flag_token_pool.Version,
	})
	require.NoError(t, err, "Failed to get non-canonical pool address from adapter")
	require.Equal(t, nonCanonicalPoolAddr.Bytes(), poolAddrNC, "Non-canonical adapter pool address should match")
	tokenAddrNC, err := nonCanonicalAdapter.TokenAddress(e.DataStore, e.BlockChains, nonCanonicalChainSelector)
	require.NoError(t, err, "Failed to get non-canonical token address from adapter")
	require.Equal(t, nonCanonicalSetup.USDCToken.Bytes(), tokenAddrNC, "Non-canonical adapter token address should match")
}
