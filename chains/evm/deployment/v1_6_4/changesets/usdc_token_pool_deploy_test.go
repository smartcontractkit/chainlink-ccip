package changesets_test

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	chain_selectors "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_4/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	"github.com/stretchr/testify/require"

	factory_burn_mint_erc20 "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/latest/factory_burn_mint_erc20"
	mock_usdc_token_messenger "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/latest/mock_usdc_token_messenger"
	mock_usdc_token_transmitter "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/latest/mock_usdc_token_transmitter"

	usdc_token_pool_bindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_4/usdc_token_pool"

	cctp_message_transmitter_proxy_binding "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/latest/cctp_message_transmitter_proxy"

	usdc_token_pool_ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_4/operations/usdc_token_pool"

	semver "github.com/Masterminds/semver/v3"
	datastore "github.com/smartcontractkit/chainlink-deployments-framework/datastore"
)

func TestUSDCTokenPoolDeployChangeset_NoExisting_MessageTransmitter_Proxy(t *testing.T) {
	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{uint64(chain_selectors.TEST_90000001.Selector)}),
	)
	require.NoError(t, err, "Failed to create environment")
	require.NotNil(t, e, "Environment should be created")

	evmChain := e.BlockChains.EVMChains()[uint64(chain_selectors.TEST_90000001.Selector)]

	tokenAddress, tx, _, err := factory_burn_mint_erc20.DeployFactoryBurnMintERC20(
		evmChain.DeployerKey,
		evmChain.Client,
		"TestToken",
		"TEST",
		6,
		big.NewInt(0),             // maxSupply (0 = unlimited)
		big.NewInt(0),             // preMint
		evmChain.DeployerKey.From, // newOwner
	)

	require.NoError(t, err, "Failed to deploy FactoryBurnMintERC20 token")
	_, err = evmChain.Confirm(tx)
	require.NoError(t, err, "Failed to confirm FactoryBurnMintERC20 token deployment transaction")

	rmnProxyAddress := common.Address{4}
	routerAddress := common.Address{5}

	// Deploy MockE2EUSDCTransmitter
	mockTransmitterAddress, tx, _, err := mock_usdc_token_transmitter.DeployMockE2EUSDCTransmitter(
		evmChain.DeployerKey,
		evmChain.Client,
		0,            // _version
		2,            // _localDomain
		tokenAddress, // token
	)
	require.NoError(t, err, "Failed to deploy MockE2EUSDCTransmitter")
	_, err = evmChain.Confirm(tx)
	require.NoError(t, err, "Failed to confirm MockE2EUSDCTransmitter deployment transaction")

	// Deploy MockUSDCTokenMessenger
	mockTokenMessengerAddress, tx, _, err := mock_usdc_token_messenger.DeployMockE2EUSDCTokenMessenger(
		evmChain.DeployerKey,
		evmChain.Client,
		0,                      // version
		mockTransmitterAddress, // transmitter
	)
	require.NoError(t, err, "Failed to deploy MockUSDCTokenMessenger")
	_, err = evmChain.Confirm(tx)
	require.NoError(t, err, "Failed to confirm MockUSDCTokenMessenger deployment transaction")

	deployInput := changesets.USDCTokenPoolDeployInput{
		ChainInputs: []changesets.USDCTokenPoolDeployInputPerChain{
			{
				ChainSelector:               uint64(chain_selectors.TEST_90000001.Selector),
				TokenMessenger:              mockTokenMessengerAddress,
				CCTPMessageTransmitterProxy: common.Address{},
				Token:                       tokenAddress,
				Allowlist:                   []common.Address{},
				RMNProxy:                    rmnProxyAddress,
				Router:                      routerAddress,
				SupportedUSDCVersion:        0,
			},
		},
		MCMS: mcms.Input{},
	}

	deployChangeset := changesets.USDCTokenPoolDeployChangeset()
	output, err := deployChangeset.Apply(*e, deployInput)
	require.NoError(t, err, "USDCTokenPoolDeployChangeset should not error")
	require.Greater(t, len(output.Reports), 0)

	usdcTokenPoolAddress, err := output.DataStore.Addresses().Get(
		datastore.NewAddressRefKey(
			uint64(chain_selectors.TEST_90000001.Selector),
			datastore.ContractType(usdc_token_pool_ops.ContractType),
			semver.MustParse("1.6.4"),
			"",
		),
	)

	// Use the go binding to create a new instance of the usdc_token_pool contract
	usdcTokenPoolInstance, err := usdc_token_pool_bindings.NewUSDCTokenPool(
		common.HexToAddress(usdcTokenPoolAddress.Address),
		evmChain.Client,
	)
	require.NoError(t, err, "Failed to bind USDCTokenPool contract")

	// Call isSupportedToken with the token address to check if it is supported
	isSupported, err := usdcTokenPoolInstance.IsSupportedToken(nil, tokenAddress)
	require.NoError(t, err, "Failed to call IsSupportedToken")
	require.True(t, isSupported, "Token should be supported by USDCTokenPool")

	// Get the CCTPMessageTransmitterProxy address from the deployed USDCTokenPool
	cctpMessageTransmitterProxyAddress, err := usdcTokenPoolInstance.IMessageTransmitterProxy(nil)
	require.NoError(t, err, "Failed to call IMessageTransmitter on USDCTokenPool")
	require.NotEmpty(t, cctpMessageTransmitterProxyAddress, "CCTPMessageTransmitterProxy address should not be empty")

	// Use the go binding to create a new instance of the CCTPMessageTransmitterProxy contract
	cctpMessageTransmitterProxyInstance, err := cctp_message_transmitter_proxy_binding.NewCCTPMessageTransmitterProxy(
		cctpMessageTransmitterProxyAddress,
		evmChain.Client,
	)
	require.NoError(t, err, "Failed to bind CCTPMessageTransmitterProxy contract")

	// Call ICctpTransmitter() on the proxy contract to get its message transmitter address
	proxyTransmitterAddress, err := cctpMessageTransmitterProxyInstance.ICctpTransmitter(nil)
	require.NoError(t, err, "Failed to call ICctpTransmitter on CCTPMessageTransmitterProxy")

	// Check that the returned address matches the mockTransmitterAddress created earlier
	require.Equal(t, mockTransmitterAddress, proxyTransmitterAddress, "The proxy's message transmitter address should match the test's mockTransmitterAddress")

	// Check that the allowed caller on the CCTPMessageTransmitterProxy is the USDCTokenPool address
	isAllowed, err := cctpMessageTransmitterProxyInstance.IsAllowedCaller(nil, common.HexToAddress(usdcTokenPoolAddress.Address))
	require.NoError(t, err, "Failed to check if the caller is allowed")
	require.True(t, isAllowed, "Caller should be allowed")
}

func TestUSDCTokenPoolDeployChangeset_Existing_MessageTransmitter_Proxy(t *testing.T) {
	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{uint64(chain_selectors.TEST_90000001.Selector)}),
	)
	require.NoError(t, err, "Failed to create environment")
	require.NotNil(t, e, "Environment should be created")

	evmChain := e.BlockChains.EVMChains()[uint64(chain_selectors.TEST_90000001.Selector)]

	tokenAddress, tx, _, err := factory_burn_mint_erc20.DeployFactoryBurnMintERC20(
		evmChain.DeployerKey,
		evmChain.Client,
		"TestToken",
		"TEST",
		6,
		big.NewInt(0),             // maxSupply (0 = unlimited)
		big.NewInt(0),             // preMint
		evmChain.DeployerKey.From, // newOwner
	)

	require.NoError(t, err, "Failed to deploy FactoryBurnMintERC20 token")
	_, err = evmChain.Confirm(tx)
	require.NoError(t, err, "Failed to confirm FactoryBurnMintERC20 token deployment transaction")

	rmnProxyAddress := common.Address{4}
	routerAddress := common.Address{5}

	// Deploy MockE2EUSDCTransmitter
	mockTransmitterAddress, tx, _, err := mock_usdc_token_transmitter.DeployMockE2EUSDCTransmitter(
		evmChain.DeployerKey,
		evmChain.Client,
		0,            // _version
		2,            // _localDomain
		tokenAddress, // token
	)
	require.NoError(t, err, "Failed to deploy MockE2EUSDCTransmitter")
	_, err = evmChain.Confirm(tx)
	require.NoError(t, err, "Failed to confirm MockE2EUSDCTransmitter deployment transaction")

	// Deploy MockUSDCTokenMessenger
	mockTokenMessengerAddress, tx, _, err := mock_usdc_token_messenger.DeployMockE2EUSDCTokenMessenger(
		evmChain.DeployerKey,
		evmChain.Client,
		0,                      // version
		mockTransmitterAddress, // transmitter
	)
	require.NoError(t, err, "Failed to deploy MockUSDCTokenMessenger")
	_, err = evmChain.Confirm(tx)
	require.NoError(t, err, "Failed to confirm MockUSDCTokenMessenger deployment transaction")

	cctpMessageTransmitterProxyAddress, tx, _, err := cctp_message_transmitter_proxy_binding.DeployCCTPMessageTransmitterProxy(
		evmChain.DeployerKey,
		evmChain.Client,
		mockTokenMessengerAddress, // tokenMessenger
	)
	require.NoError(t, err, "Failed to deploy CCTPMessageTransmitterProxy")
	_, err = evmChain.Confirm(tx)
	require.NoError(t, err, "Failed to confirm CCTPMessageTransmitterProxy deployment transaction")

	deployInput := changesets.USDCTokenPoolDeployInput{
		ChainInputs: []changesets.USDCTokenPoolDeployInputPerChain{
			{
				ChainSelector:               uint64(chain_selectors.TEST_90000001.Selector),
				TokenMessenger:              mockTokenMessengerAddress,
				CCTPMessageTransmitterProxy: cctpMessageTransmitterProxyAddress,
				Token:                       tokenAddress,
				Allowlist:                   []common.Address{},
				RMNProxy:                    rmnProxyAddress,
				Router:                      routerAddress,
				SupportedUSDCVersion:        0,
			},
		},
		MCMS: mcms.Input{},
	}

	deployChangeset := changesets.USDCTokenPoolDeployChangeset()
	output, err := deployChangeset.Apply(*e, deployInput)
	require.NoError(t, err, "USDCTokenPoolDeployChangeset should not error")
	require.Greater(t, len(output.Reports), 0)

	usdcTokenPoolAddress, err := output.DataStore.Addresses().Get(
		datastore.NewAddressRefKey(
			uint64(chain_selectors.TEST_90000001.Selector),
			datastore.ContractType(usdc_token_pool_ops.ContractType),
			semver.MustParse("1.6.4"),
			"",
		),
	)

	// Use the go binding to create a new instance of the usdc_token_pool contract
	usdcTokenPoolInstance, err := usdc_token_pool_bindings.NewUSDCTokenPool(
		common.HexToAddress(usdcTokenPoolAddress.Address),
		evmChain.Client,
	)
	require.NoError(t, err, "Failed to bind USDCTokenPool contract")

	// Call isSupportedToken with the token address to check if it is supported
	isSupported, err := usdcTokenPoolInstance.IsSupportedToken(nil, tokenAddress)
	require.NoError(t, err, "Failed to call IsSupportedToken")
	require.True(t, isSupported, "Token should be supported by USDCTokenPool")

	// Get the CCTPMessageTransmitterProxy address from the deployed USDCTokenPool
	cctpMessageTransmitterProxyAddressFromPool, err := usdcTokenPoolInstance.IMessageTransmitterProxy(nil)
	require.NoError(t, err, "Failed to call IMessageTransmitter on USDCTokenPool")
	require.NotEmpty(t, cctpMessageTransmitterProxyAddressFromPool, "CCTPMessageTransmitterProxy address should not be empty")

	// Check that the CCTPMessageTransmitterProxy address matches the deployed address
	require.Equal(t, cctpMessageTransmitterProxyAddress, cctpMessageTransmitterProxyAddressFromPool, "CCTPMessageTransmitterProxy address should match the deployed address")

	// Use the go binding to create a new instance of the CCTPMessageTransmitterProxy contract
	cctpMessageTransmitterProxyInstance, err := cctp_message_transmitter_proxy_binding.NewCCTPMessageTransmitterProxy(
		cctpMessageTransmitterProxyAddressFromPool,
		evmChain.Client,
	)
	require.NoError(t, err, "Failed to bind CCTPMessageTransmitterProxy contract")

	// Check that the allowed caller on the CCTPMessageTransmitterProxy is the USDCTokenPool address
	isAllowed, err := cctpMessageTransmitterProxyInstance.IsAllowedCaller(nil, common.HexToAddress(usdcTokenPoolAddress.Address))
	require.NoError(t, err, "Failed to check if the caller is allowed")
	require.True(t, isAllowed, "Caller should be allowed")
}
