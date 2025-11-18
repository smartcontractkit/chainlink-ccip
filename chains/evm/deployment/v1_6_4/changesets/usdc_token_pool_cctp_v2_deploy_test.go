package changesets_test

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	chain_selectors "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_4/changesets"
	cctp_message_transmitter_proxy_ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_4/operations/cctp_message_transmitter_proxy"
	usdc_token_pool_cctp_v2_ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_4/operations/usdc_token_pool_cctp_v2"

	usdc_token_pool_cctp_v2_binding "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/latest/usdc_token_pool_cctp_v2"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	"github.com/stretchr/testify/require"

	mock_usdc_token_messenger "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/latest/mock_usdc_token_messenger"
	mock_usdc_token_transmitter "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/latest/mock_usdc_token_transmitter"

	cctp_message_transmitter_proxy_binding "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/latest/cctp_message_transmitter_proxy"
	factory_burn_mint_erc20 "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/latest/factory_burn_mint_erc20"
	mcms_types "github.com/smartcontractkit/mcms/types"
)

func TestUSDCTokenPoolCCTPV2DeployChangeset(t *testing.T) {
	chainSelector := uint64(chain_selectors.TEST_90000001.Selector)
	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{chainSelector}),
	)
	require.NoError(t, err, "Failed to create environment")
	require.NotNil(t, e, "Environment should be created")

	evmChain := e.BlockChains.EVMChains()[chainSelector]

	tokenAddress, tx, _, err := factory_burn_mint_erc20.DeployFactoryBurnMintERC20(
		evmChain.DeployerKey,
		evmChain.Client,
		"TestToken",
		"TEST6",
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
		1,            // _version
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
		1,                      // version
		mockTransmitterAddress, // transmitter
	)
	require.NoError(t, err, "Failed to deploy MockUSDCTokenMessenger")
	_, err = evmChain.Confirm(tx)
	require.NoError(t, err, "Failed to confirm MockUSDCTokenMessenger deployment transaction")

	deployInput := changesets.USDCTokenPoolCCTPV2DeployInput{
		ChainInputs: []changesets.USDCTokenPoolCCTPV2DeployInputPerChain{
			{
				ChainSelector:  chainSelector,
				TokenMessenger: mockTokenMessengerAddress,
				Token:          tokenAddress,
				Allowlist:      []common.Address{},
				RMNProxy:       rmnProxyAddress,
				Router:         routerAddress,
			},
		},
		MCMS: mcms.Input{
			OverridePreviousRoot: false,
			ValidUntil:           3759765795,
			TimelockDelay:        mcms_types.MustParseDuration("0s"),
			TimelockAction:       mcms_types.TimelockActionSchedule,
			Qualifier:            "test",
			Description:          "Deploy USDCTokenPoolCCTPV2",
		},
	}

	deployChangeset := changesets.USDCTokenPoolCCTPV2DeployChangeset()
	output, err := deployChangeset.Apply(*e, deployInput)
	require.NoError(t, err, "USDCTokenPoolCCTPV2DeployChangeset should not error")
	require.Greater(t, len(output.Reports), 0)

	// Verify the deployed contract address
	addressRef, err := output.DataStore.Addresses().Get(
		datastore.NewAddressRefKey(
			chainSelector,
			datastore.ContractType(usdc_token_pool_cctp_v2_ops.ContractType),
			usdc_token_pool_cctp_v2_ops.Version,
			"",
		),
	)
	require.NoError(t, err, "Failed to get address from datastore")
	require.NotNil(t, addressRef, "Address should be set in datastore")

	// Verify the CCTPMessageTransmitterProxy address
	cctpMessageTransmitterProxyAddress, err := output.DataStore.Addresses().Get(
		datastore.NewAddressRefKey(
			chainSelector,
			datastore.ContractType(cctp_message_transmitter_proxy_ops.ContractType),
			cctp_message_transmitter_proxy_ops.Version,
			"",
		),
	)
	require.NoError(t, err, "Failed to get address from datastore")
	require.NotNil(t, cctpMessageTransmitterProxyAddress, "Address should be set in datastore")

	usdcTokenPoolCCTPV2, err := usdc_token_pool_cctp_v2_binding.NewUSDCTokenPoolCCTPV2(
		common.HexToAddress(addressRef.Address),
		evmChain.Client,
	)
	require.NoError(t, err, "Failed to bind USDCTokenPoolCCTPV2")

	// Retrieve the CCTPMessageTransmitterProxy address from USDCTokenPoolCCTPV2
	cctpMsgTransmitterProxyFromPool, err := usdcTokenPoolCCTPV2.IMessageTransmitterProxy(nil)
	require.NoError(t, err, "Failed to get CCTPMessageTransmitterProxy address from USDCTokenPoolCCTPV2")

	require.Equal(t,
		common.HexToAddress(cctpMessageTransmitterProxyAddress.Address),
		cctpMsgTransmitterProxyFromPool,
		"CCTPMessageTransmitterProxy address in USDCTokenPoolCCTPV2 should match the deployed proxy address",
	)

	// Instantiate a new CCTPMessageTransmitterProxy contract binding from the deployed address
	cctpMessageTransmitterProxy, err := cctp_message_transmitter_proxy_binding.NewCCTPMessageTransmitterProxy(
		common.HexToAddress(cctpMessageTransmitterProxyAddress.Address),
		evmChain.Client,
	)
	require.NoError(t, err, "Failed to bind CCTPMessageTransmitterProxy")

	// Check that the operation to set the allowed caller on the message transmitter
	// proxy was successful
	isAllowed, err := cctpMessageTransmitterProxy.IsAllowedCaller(nil, common.HexToAddress(addressRef.Address))
	require.NoError(t, err, "Failed to check if the caller is allowed")
	require.True(t, isAllowed, "Caller should be allowed")

	// Call messageTransmitter() on the contract
	messageTransmitterAddr, err := cctpMessageTransmitterProxy.ICctpTransmitter(nil)
	require.NoError(t, err, "Failed to call messageTransmitter on CCTPMessageTransmitterProxy")

	// Check that the returned address matches the mockTokenMessengerAddress used in this test
	require.Equal(t, mockTransmitterAddress, messageTransmitterAddr, "messageTransmitter address should match the test's mockTokenMessengerAddress")
}
