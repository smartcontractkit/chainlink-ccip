package cctp

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/advanced_pool_hooks"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/cctp_message_transmitter_proxy"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/cctp_through_ccv_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/cctp_verifier"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/usdc_token_pool_proxy"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/sequences"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/testsetup"
	cctp_message_transmitter_proxy_bindings "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/gobindings/generated/latest/cctp_message_transmitter_proxy"
	cctp_through_ccv_token_pool_bindings "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/gobindings/generated/latest/cctp_through_ccv_token_pool"
	cctp_verifier_bindings "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/gobindings/generated/latest/cctp_verifier"
	mock_usdc_token_messenger "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/gobindings/generated/latest/mock_usdc_token_messenger"
	mock_usdc_token_transmitter "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/gobindings/generated/latest/mock_usdc_token_transmitter"
	usdc_token_pool_proxy_bindings "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/gobindings/generated/latest/usdc_token_pool_proxy"
	versioned_verifier_resolver_bindings "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/gobindings/generated/latest/versioned_verifier_resolver"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/rmn_proxy"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/router"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/token_admin_registry"
	"github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/adapters"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	burn_mint_erc20_bindings "github.com/smartcontractkit/chainlink-evm/gethwrappers/shared/generated/initial/burn_mint_erc20"
	"github.com/stretchr/testify/require"
)

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
	chainReport, err := operations.ExecuteSequence(
		e.OperationsBundle,
		sequences.DeployChainContracts,
		chain,
		sequences.DeployChainContractsInput{
			ChainSelector:  chainSelector,
			ContractParams: testsetup.CreateBasicContractParams(),
		},
	)
	require.NoError(t, err, "Failed to deploy chain contracts")

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
		big.NewInt(0), // premintAmount
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

func basicDeployCCTPInput(chainSelector uint64, setup cctpTestSetup, deployerAddr common.Address) adapters.DeployCCTPInput[string, []byte] {
	return adapters.DeployCCTPInput[string, []byte]{
		ChainSelector:                    chainSelector,
		MessageTransmitterProxy:          "",
		TokenAdminRegistry:               setup.TokenAdminRegistry.Hex(),
		TokenMessenger:                   setup.TokenMessenger.Hex(),
		USDCToken:                        setup.USDCToken.Hex(),
		MinFinalityValue:                 1,
		StorageLocations:                 []string{"https://test.chain.link.fake"},
		FeeAggregator:                    common.HexToAddress("0x04").Hex(),
		AllowlistAdmin:                   common.HexToAddress("0x05").Hex(),
		FastFinalityBps:                  100,
		RMN:                              setup.RMN.Hex(),
		Router:                           setup.Router.Hex(),
		DeployerContract:                 "",
		Allowlist:                        []string{common.HexToAddress("0x08").Hex()},
		ThresholdAmountForAdditionalCCVs: big.NewInt(1e18),
		RateLimitAdmin:                   deployerAddr.Hex(),
		RemoteChains:                     make(map[uint64]adapters.RemoteCCTPChainConfig[string, []byte]),
	}
}

func TestDeployCCTPChain(t *testing.T) {
	chainSelector := uint64(3017758115101368649)
	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{chainSelector}),
	)
	require.NoError(t, err, "Failed to create environment")
	require.NotNil(t, e, "Environment should be created")
	e.DataStore = datastore.NewMemoryDataStore().Seal()

	setup := setupCCTPTestEnvironment(t, e, chainSelector)
	chain := e.BlockChains.EVMChains()[chainSelector]
	input := basicDeployCCTPInput(chainSelector, setup, chain.DeployerKey.From)
	// Add a remote chain config for testing (CCTP V2 with CCV)
	remoteChainSelector := uint64(4949039107694359620)
	input.RemoteChains[remoteChainSelector] = adapters.RemoteCCTPChainConfig[string, []byte]{
		FeeUSDCents:         10,
		GasForVerification:  100000,
		PayloadSizeBytes:    1000,
		LockOrBurnMechanism: "CCTP_V2_WITH_CCV",
		RemoteDomain: adapters.RemoteDomain[[]byte]{
			AllowedCallerOnDest:   common.LeftPadBytes(common.HexToAddress("0x0D").Bytes(), 32),
			AllowedCallerOnSource: common.LeftPadBytes(common.HexToAddress("0x0E").Bytes(), 32),
			MintRecipientOnDest:   common.LeftPadBytes(common.HexToAddress("0x0F").Bytes(), 32),
			DomainIdentifier:      1,
		},
		TokenPoolConfig: tokens.RemoteChainConfig[[]byte, string]{
			RemotePool:                               common.LeftPadBytes(common.HexToAddress("0x1E").Bytes(), 32),
			RemoteToken:                              common.LeftPadBytes(setup.USDCToken.Bytes(), 32),
			DefaultFinalityInboundRateLimiterConfig:  testsetup.CreateRateLimiterConfig(0, 0),
			DefaultFinalityOutboundRateLimiterConfig: testsetup.CreateRateLimiterConfig(0, 0),
			CustomFinalityInboundRateLimiterConfig:   testsetup.CreateRateLimiterConfig(0, 0),
			CustomFinalityOutboundRateLimiterConfig:  testsetup.CreateRateLimiterConfig(0, 0),
		},
	}

	report, err := operations.ExecuteSequence(
		e.OperationsBundle,
		DeployCCTPChain,
		e.BlockChains,
		input,
	)
	require.NoError(t, err, "ExecuteSequence should not error")

	exists := map[deployment.ContractType]bool{
		deployment.ContractType(advanced_pool_hooks.ContractType):            false,
		deployment.ContractType(cctp_through_ccv_token_pool.ContractType):    false,
		deployment.ContractType(cctp_message_transmitter_proxy.ContractType): false,
		deployment.ContractType(cctp_verifier.ContractType):                  false,
		deployment.ContractType(usdc_token_pool_proxy.ContractType):          false,
		deployment.ContractType(cctp_verifier.ResolverType):                  false,
	}
	for _, addr := range report.Output.Addresses {
		exists[deployment.ContractType(addr.Type)] = true
	}
	for ctype, found := range exists {
		require.True(t, found, "Expected contract of type %s to be deployed", ctype)
	}

	// Extract contract addresses from report
	var cctpTokenPoolAddr, cctpMessageTransmitterProxyAddr, cctpVerifierAddr, usdcTokenPoolProxyAddr, cctpVerifierResolverAddr common.Address
	for _, addr := range report.Output.Addresses {
		switch deployment.ContractType(addr.Type) {
		case deployment.ContractType(cctp_through_ccv_token_pool.ContractType):
			cctpTokenPoolAddr = common.HexToAddress(addr.Address)
		case deployment.ContractType(cctp_message_transmitter_proxy.ContractType):
			cctpMessageTransmitterProxyAddr = common.HexToAddress(addr.Address)
		case deployment.ContractType(cctp_verifier.ContractType):
			cctpVerifierAddr = common.HexToAddress(addr.Address)
		case deployment.ContractType(usdc_token_pool_proxy.ContractType):
			usdcTokenPoolProxyAddr = common.HexToAddress(addr.Address)
		case deployment.ContractType(cctp_verifier.ResolverType):
			cctpVerifierResolverAddr = common.HexToAddress(addr.Address)
		}
	}

	require.NotEqual(t, common.Address{}, cctpTokenPoolAddr, "CCTPTokenPool address should be set")
	require.NotEqual(t, common.Address{}, cctpMessageTransmitterProxyAddr, "CCTPMessageTransmitterProxy address should be set")
	require.NotEqual(t, common.Address{}, cctpVerifierAddr, "CCTPVerifier address should be set")
	require.NotEqual(t, common.Address{}, usdcTokenPoolProxyAddr, "USDCTokenPoolProxy address should be set")
	require.NotEqual(t, common.Address{}, cctpVerifierResolverAddr, "CCTPVerifierResolver address should be set")

	// Check CCTPTokenPool dynamic config
	cctpTokenPool, err := cctp_through_ccv_token_pool_bindings.NewCCTPThroughCCVTokenPool(cctpTokenPoolAddr, chain.Client)
	require.NoError(t, err, "Failed to instantiate CCTPTokenPool contract")
	dynamicConfig, err := cctpTokenPool.GetDynamicConfig(nil)
	require.NoError(t, err, "Failed to get dynamic config from CCTPTokenPool")
	require.Equal(t, setup.Router, dynamicConfig.Router, "Router should match")
	require.Equal(t, chain.DeployerKey.From, dynamicConfig.RateLimitAdmin, "RateLimitAdmin should match deployer")

	// Check CCTPTokenPool authorized callers
	authorizedCallers, err := cctpTokenPool.GetAllAuthorizedCallers(nil)
	require.NoError(t, err, "Failed to get authorized callers from CCTPTokenPool")
	require.Contains(t, authorizedCallers, usdcTokenPoolProxyAddr, "USDCTokenPoolProxy should be an authorized caller")

	// Check CCTPTokenPool token
	tokenAddr, err := cctpTokenPool.GetToken(nil)
	require.NoError(t, err, "Failed to get token from CCTPTokenPool")
	require.Equal(t, setup.USDCToken, tokenAddr, "Token address should match")

	// Check CCTPMessageTransmitterProxy authorized callers
	cctpMessageTransmitterProxy, err := cctp_message_transmitter_proxy_bindings.NewCCTPMessageTransmitterProxy(cctpMessageTransmitterProxyAddr, chain.Client)
	require.NoError(t, err, "Failed to instantiate CCTPMessageTransmitterProxy contract")
	allowedCallers, err := cctpMessageTransmitterProxy.GetAllAuthorizedCallers(nil)
	require.NoError(t, err, "Failed to get allowed callers from CCTPMessageTransmitterProxy")
	require.Contains(t, allowedCallers, cctpVerifierAddr, "CCTPVerifier should be an allowed caller")

	// Check CCTPVerifierResolver inbound implementation
	cctpVerifierResolver, err := versioned_verifier_resolver_bindings.NewVersionedVerifierResolver(cctpVerifierResolverAddr, chain.Client)
	require.NoError(t, err, "Failed to instantiate CCTPVerifierResolver contract")
	allInboundImpls, err := cctpVerifierResolver.GetAllInboundImplementations(nil)
	require.NoError(t, err, "Failed to get inbound implementations from CCTPVerifierResolver")
	require.Greater(t, len(allInboundImpls), 0, "Should have at least one inbound implementation")
	foundInbound := false
	for _, impl := range allInboundImpls {
		if impl.Verifier == cctpVerifierAddr {
			foundInbound = true
			break
		}
	}
	require.True(t, foundInbound, "CCTPVerifier should be registered as inbound implementation")

	// Check CCTPVerifierResolver outbound implementation
	allOutboundImpls, err := cctpVerifierResolver.GetAllOutboundImplementations(nil)
	require.NoError(t, err, "Failed to get outbound implementations from CCTPVerifierResolver")
	foundOutbound := false
	for _, impl := range allOutboundImpls {
		if impl.DestChainSelector == remoteChainSelector && impl.Verifier == cctpVerifierAddr {
			foundOutbound = true
			break
		}
	}
	require.True(t, foundOutbound, "CCTPVerifier should be registered as outbound implementation for remote chain")

	// Check USDCTokenPoolProxy lock or burn mechanism
	usdcTokenPoolProxy, err := usdc_token_pool_proxy_bindings.NewUSDCTokenPoolProxy(usdcTokenPoolProxyAddr, chain.Client)
	require.NoError(t, err, "Failed to instantiate USDCTokenPoolProxy contract")
	mechanism, err := usdcTokenPoolProxy.GetLockOrBurnMechanism(nil, remoteChainSelector)
	require.NoError(t, err, "Failed to get lock or burn mechanism from USDCTokenPoolProxy")
	expectedMechanism, err := convertMechanismToUint8("CCTP_V2_WITH_CCV")
	require.NoError(t, err, "Failed to convert mechanism to uint8")
	require.Equal(t, expectedMechanism, uint8(mechanism), "Lock or burn mechanism should match")

	// Check USDCTokenPoolProxy required CCVs
	requiredCCVs, err := usdcTokenPoolProxy.GetRequiredCCVs(nil, setup.USDCToken, remoteChainSelector, big.NewInt(1e18), 1, []byte{}, 0)
	require.NoError(t, err, "Failed to get required CCVs from USDCTokenPoolProxy")
	require.Equal(t, []common.Address{cctpVerifierResolverAddr}, requiredCCVs, "Required CCVs should match")

	// Check CCTPVerifier remote chain config
	cctpVerifier, err := cctp_verifier_bindings.NewCCTPVerifier(cctpVerifierAddr, chain.Client)
	require.NoError(t, err, "Failed to instantiate CCTPVerifier contract")
	staticConfig, err := cctpVerifier.GetStaticConfig(nil)
	require.NoError(t, err, "Failed to get static config from CCTPVerifier")
	require.Equal(t, setup.TokenMessenger, staticConfig.TokenMessenger, "TokenMessenger should match")
	require.Equal(t, cctpMessageTransmitterProxyAddr, staticConfig.MessageTransmitterProxy, "MessageTransmitterProxy should match")
	require.Equal(t, setup.USDCToken, staticConfig.UsdcToken, "USDCToken should match")

	// Check CCTPVerifier domains
	domain, err := cctpVerifier.GetDomain(nil, remoteChainSelector)
	require.NoError(t, err, "Failed to get domain from CCTPVerifier")
	require.Equal(t, uint32(1), domain.DomainIdentifier, "Domain identifier should match")
	require.True(t, domain.Enabled, "Domain should be enabled")

	// Convert slices to arrays for comparison
	var expectedAllowedCallerOnDest [32]byte
	copy(expectedAllowedCallerOnDest[:], common.LeftPadBytes(common.HexToAddress("0x0D").Bytes(), 32))
	require.Equal(t, expectedAllowedCallerOnDest, domain.AllowedCallerOnDest, "AllowedCallerOnDest should match")

	var expectedAllowedCallerOnSource [32]byte
	copy(expectedAllowedCallerOnSource[:], common.LeftPadBytes(common.HexToAddress("0x0E").Bytes(), 32))
	require.Equal(t, expectedAllowedCallerOnSource, domain.AllowedCallerOnSource, "AllowedCallerOnSource should match")

	var expectedMintRecipientOnDest [32]byte
	copy(expectedMintRecipientOnDest[:], common.LeftPadBytes(common.HexToAddress("0x0F").Bytes(), 32))
	require.Equal(t, expectedMintRecipientOnDest, domain.MintRecipientOnDest, "MintRecipientOnDest should match")
}
