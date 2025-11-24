package changesets_test

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	mcms_types "github.com/smartcontractkit/mcms/types"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	contract_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_4/operations/usdc_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/latest/cctp_message_transmitter_proxy"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/latest/factory_burn_mint_erc20"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/latest/mock_usdc_token_messenger"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/latest/mock_usdc_token_transmitter"

	changesets "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_4/changesets"
	changesets_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"

	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"

	chain_selectors "github.com/smartcontractkit/chain-selectors"

	usdc_token_pool_bindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_4/usdc_token_pool"

	"github.com/stretchr/testify/require"
)

func TestSetDomainsSequence(t *testing.T) {
	chainSelector := uint64(chain_selectors.TEST_90000001.Selector)
	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{chainSelector}),
	)
	require.NoError(t, err, "Failed to create environment")
	require.NotNil(t, e, "Environment should be created")
	ds := datastore.NewMemoryDataStore()

	evmChain := e.BlockChains.EVMChains()[chainSelector]

	// Deploy a real ERC20 token using factory_burn_mint_erc20
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
	require.NoError(t, err, "Failed to confirm token deployment transaction")

	// Deploy MockE2EUSDCTransmitter
	mockTransmitterAddress, tx, _, err := mock_usdc_token_transmitter.DeployMockE2EUSDCTransmitter(
		evmChain.DeployerKey,
		evmChain.Client,
		0,            // _version
		1,            // _localDomain
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

	// Deploy CCTPMessageTransmitterProxy
	cctpMessageTransmitterProxyAddress, tx, _, err := cctp_message_transmitter_proxy.DeployCCTPMessageTransmitterProxy(
		evmChain.DeployerKey,
		evmChain.Client,
		mockTokenMessengerAddress, // tokenMessenger
	)
	require.NoError(t, err, "Failed to deploy CCTPMessageTransmitterProxy")
	_, err = evmChain.Confirm(tx)
	require.NoError(t, err, "Failed to confirm CCTPMessageTransmitterProxy deployment transaction")

	// Deploy USDC Token Pool with placeholder addresses for dependencies
	usdcTokenPoolRef, err := contract_utils.MaybeDeployContract(e.OperationsBundle, usdc_token_pool.Deploy, evmChain, contract.DeployInput[usdc_token_pool.ConstructorArgs]{
		TypeAndVersion: deployment.NewTypeAndVersion(usdc_token_pool.ContractType, *usdc_token_pool.Version),
		ChainSelector:  chainSelector,
		Args: usdc_token_pool.ConstructorArgs{
			TokenMessenger:              mockTokenMessengerAddress,
			CCTPMessageTransmitterProxy: cctpMessageTransmitterProxyAddress,
			Token:                       tokenAddress,
			Allowlist:                   []common.Address{},
			RMNProxy:                    common.HexToAddress("0x04"),
			Router:                      common.HexToAddress("0x05"),
			SupportedUSDCVersion:        uint32(0),
		},
	}, nil)

	require.NoError(t, err, "Failed to deploy USDCTokenPool")
	require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
		Type:          datastore.ContractType(usdc_token_pool.ContractType),
		Version:       usdc_token_pool.Version,
		ChainSelector: chainSelector,
		Address:       usdcTokenPoolRef.Address,
	}))

	// update env datastore from the ds datastore
	e.DataStore = ds.Seal()

	// Create the input to the changeset
	setDomainsInput := changesets.SetDomainsInput{
		ChainInputs: []changesets.SetDomainsPerChainInput{
			{
				ChainSelector: chainSelector,
				Domains: []usdc_token_pool.DomainUpdate{
					{
						AllowedCaller:                 [32]byte{1},
						MintRecipient:                 [32]byte{2},
						DomainIdentifier:              uint32(3),
						DestChainSelector:             chainSelector,
						Enabled:                       true,
						UseLegacySourcePoolDataFormat: true,
					},
				},
			},
		},
		MCMS: mcms.Input{
			OverridePreviousRoot: false,
			ValidUntil:           3759765795,
			TimelockDelay:        mcms_types.MustParseDuration("0s"),
			TimelockAction:       mcms_types.TimelockActionSchedule,
			Qualifier:            "test",
			Description:          "Set domains on USDCTokenPool",
		},
	}

	// Create the changeset. No adapter is needed here because the changeset is using the MCMSReaderRegistry only
	mcmsRegistry := changesets_utils.GetRegistry()
	setDomainsChangeset := changesets.SetDomainsChangeset(mcmsRegistry)
	output, err := setDomainsChangeset.Apply(*e, setDomainsInput)
	require.NoError(t, err, "SetDomainsChangeset should not error")
	require.Greater(t, len(output.Reports), 0)

	// Verify the domains
	usdcTokenPool, err := usdc_token_pool_bindings.NewUSDCTokenPool(common.HexToAddress(usdcTokenPoolRef.Address), evmChain.Client)
	require.NoError(t, err, "Failed to create USDCTokenPool")
	domain, err := usdcTokenPool.GetDomain(&bind.CallOpts{Context: t.Context()}, chainSelector)
	require.NoError(t, err, "Failed to get domain")
	require.Equal(t, domain.AllowedCaller, [32]byte{1}, "Allowed caller should be 1")
	require.Equal(t, domain.MintRecipient, [32]byte{2}, "Mint recipient should be 2")
	require.Equal(t, domain.DomainIdentifier, uint32(3), "Domain identifier should be 3")
	require.Equal(t, domain.Enabled, true, "Domain should be enabled")
	require.Equal(t, domain.UseLegacySourcePoolDataFormat, true, "Use legacy source pool data format should be true")
}
