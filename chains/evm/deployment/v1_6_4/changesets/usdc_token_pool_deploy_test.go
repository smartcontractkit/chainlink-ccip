package changesets_test

import (
	"math/big"
	"testing"

	"github.com/Masterminds/semver/v3"

	"github.com/aws/smithy-go/ptr"
	"github.com/ethereum/go-ethereum/common"
	chain_selectors "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/adapters"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_4/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/deploy"
	"github.com/smartcontractkit/chainlink-ccip/deployment/testhelpers"
	deploymentutils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	mcms_types "github.com/smartcontractkit/mcms/types"
	"github.com/stretchr/testify/require"

	factory_burn_mint_erc20 "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/latest/factory_burn_mint_erc20"
	mock_usdc_token_messenger "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/latest/mock_usdc_token_messenger"
	mock_usdc_token_transmitter "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/latest/mock_usdc_token_transmitter"

	authorized_caller_ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_4/operations/authorized_caller"
	cctp_message_transmitter_proxy_binding "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/latest/cctp_message_transmitter_proxy"
	usdc_token_pool_bindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_4/usdc_token_pool"

	usdc_token_pool_ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_4/operations/usdc_token_pool"

	changesets_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	datastore "github.com/smartcontractkit/chainlink-deployments-framework/datastore"
)

func TestUSDCTokenPoolDeployChangeset(t *testing.T) {
	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{uint64(chain_selectors.TEST_90000001.Selector)}),
	)
	require.NoError(t, err, "Failed to create environment")
	require.NotNil(t, e, "Environment should be created")

	evmChain := e.BlockChains.EVMChains()[uint64(chain_selectors.TEST_90000001.Selector)]
	ds := datastore.NewMemoryDataStore()

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
	err = ds.Addresses().Add(datastore.AddressRef{
		Type:          datastore.ContractType("USDCToken"),
		Version:       semver.MustParse("1.0.0"),
		Address:       tokenAddress.Hex(),
		ChainSelector: uint64(chain_selectors.TEST_90000001.Selector),
	})
	require.NoError(t, err, "Failed to add USDCToken address to datastore")

	rmnProxyAddress := common.Address{4}
	err = ds.Addresses().Add(datastore.AddressRef{
		Type:          datastore.ContractType("RMN"),
		Version:       semver.MustParse("1.5.0"),
		Address:       rmnProxyAddress.Hex(),
		ChainSelector: uint64(chain_selectors.TEST_90000001.Selector),
	})
	require.NoError(t, err, "Failed to add RMN address to datastore")

	routerAddress := common.Address{5}
	err = ds.Addresses().Add(datastore.AddressRef{
		Type:          datastore.ContractType("Router"),
		Version:       semver.MustParse("1.2.0"),
		Address:       routerAddress.Hex(),
		ChainSelector: uint64(chain_selectors.TEST_90000001.Selector),
	})
	require.NoError(t, err, "Failed to add router address to datastore")

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
		mockTokenMessengerAddress,
	)
	require.NoError(t, err, "Failed to deploy CCTPMessageTransmitterProxy")
	_, err = evmChain.Confirm(tx)
	require.NoError(t, err, "Failed to confirm CCTPMessageTransmitterProxy deployment transaction")

	err = ds.Addresses().Add(datastore.AddressRef{
		Type:          datastore.ContractType("CCTPMessageTransmitterProxy"),
		Version:       semver.MustParse("1.6.2"),
		Address:       cctpMessageTransmitterProxyAddress.Hex(),
		ChainSelector: uint64(chain_selectors.TEST_90000001.Selector),
	})
	require.NoError(t, err, "Failed to add CCTPMessageTransmitterProxy address to datastore")
	e.DataStore = ds.Seal()

	deployInput := changesets.USDCTokenPoolDeployInput{
		ChainInputs: []changesets.USDCTokenPoolDeployInputPerChain{
			{
				ChainSelector:  uint64(chain_selectors.TEST_90000001.Selector),
				TokenMessenger: mockTokenMessengerAddress,
				Allowlist:      []common.Address{},
			},
		},
		MCMS: mcms.Input{
			OverridePreviousRoot: false,
			ValidUntil:           3759765795,
			TimelockDelay:        mcms_types.MustParseDuration("0s"),
			TimelockAction:       mcms_types.TimelockActionSchedule,
			Qualifier:            "test",
			Description:          "Deploy USDCTokenPool",
		},
	}

	// deploy mcms
	evmDeployer := &adapters.EVMDeployer{}
	dReg := deploy.GetRegistry()
	dReg.RegisterDeployer(chain_selectors.FamilyEVM, deploy.MCMSVersion, evmDeployer)
	cs := deploy.DeployMCMS(dReg, nil)
	output, err := cs.Apply(*e, deploy.MCMSDeploymentConfig{
		AdapterVersion: semver.MustParse("1.0.0"),
		Chains: map[uint64]deploy.MCMSDeploymentConfigPerChain{
			uint64(chain_selectors.TEST_90000001.Selector): {
				Canceller:        testhelpers.SingleGroupMCMS(),
				Bypasser:         testhelpers.SingleGroupMCMS(),
				Proposer:         testhelpers.SingleGroupMCMS(),
				TimelockMinDelay: big.NewInt(0),
				Qualifier:        ptr.String("test"),
				TimelockAdmin:    evmChain.DeployerKey.From,
			},
		},
	})

	allAddrRefs, err := output.DataStore.Addresses().Fetch()
	require.NoError(t, err)
	timelockAddrs := make(map[uint64]string)
	for _, addrRef := range allAddrRefs {
		require.NoError(t, ds.Addresses().Add(addrRef))
		if addrRef.Type == datastore.ContractType(deploymentutils.RBACTimelock) {
			timelockAddrs[addrRef.ChainSelector] = addrRef.Address
		}
	}
	// update env datastore
	e.DataStore = ds.Seal()

	// Register the MCMS Reader
	mcmsRegistry := changesets_utils.GetRegistry()
	evmMCMSReader := &adapters.EVMMCMSReader{}
	mcmsRegistry.RegisterMCMSReader(chain_selectors.FamilyEVM, evmMCMSReader)

	deployChangeset := changesets.USDCTokenPoolDeployChangeset(mcmsRegistry)
	deployChangesetOutput, err := deployChangeset.Apply(*e, deployInput)
	require.NoError(t, err, "USDCTokenPoolDeployChangeset should not error")
	require.Greater(t, len(deployChangesetOutput.Reports), 0)

	usdcTokenPoolAddress, err := deployChangesetOutput.DataStore.Addresses().Get(
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
	onChain_cctpMessageTransmitterProxyAddress, err := usdcTokenPoolInstance.IMessageTransmitterProxy(nil)
	require.NoError(t, err, "Failed to call IMessageTransmitter on USDCTokenPool")
	require.NotEmpty(t, cctpMessageTransmitterProxyAddress, "CCTPMessageTransmitterProxy address should not be empty")
	require.Equal(t, cctpMessageTransmitterProxyAddress, onChain_cctpMessageTransmitterProxyAddress, "CCTPMessageTransmitterProxy address should match the deployed address")

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
	require.False(t, isAllowed, "Caller should be allowed")

	// Apply authorized caller updates to the USDCTokenPool
	applyAuthorizedCallerUpdatesInput := changesets.ApplyAuthorizedCallerUpdatesInput{
		ChainInputs: []changesets.ApplyAuthorizedCallerUpdatesInputPerChain{
			{
				ChainSelector: uint64(chain_selectors.TEST_90000001.Selector),
				Address:       common.HexToAddress(usdcTokenPoolAddress.Address),
				AuthorizedCallerUpdates: authorized_caller_ops.AuthorizedCallerUpdateArgs{
					AddedCallers: []common.Address{common.HexToAddress(usdcTokenPoolAddress.Address)},
				},
			},
		},
		MCMS: mcms.Input{
			OverridePreviousRoot: false,
			ValidUntil:           3759765795,
			TimelockDelay:        mcms_types.MustParseDuration("0s"),
			TimelockAction:       mcms_types.TimelockActionSchedule,
			Qualifier:            "test",
			Description:          "Apply authorized caller updates",
		},
	}
	applyAuthorizedCallerUpdatesChangeset := changesets.ApplyAuthorizedCallerUpdatesChangeset(mcmsRegistry)
	applyAuthorizedCallerUpdatesChangesetOutput, err := applyAuthorizedCallerUpdatesChangeset.Apply(*e, applyAuthorizedCallerUpdatesInput)
	require.NoError(t, err, "ApplyAuthorizedCallerUpdatesChangeset should not error")
	require.Greater(t, len(applyAuthorizedCallerUpdatesChangesetOutput.Reports), 0)

	// Check that the allowed caller on the USDCTokenPool is the CCTPMessageTransmitterProxy address
	allAuthorizedCallers, err := usdcTokenPoolInstance.GetAllAuthorizedCallers(nil)
	require.NoError(t, err, "Failed to get all authorized callers")

	require.Equal(t, allAuthorizedCallers, []common.Address{common.HexToAddress(usdcTokenPoolAddress.Address)}, "There should only be one authorized caller")
}
