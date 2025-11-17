package sequences_test

import (
	"math/big"
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/aws/smithy-go/ptr"
	"github.com/ethereum/go-ethereum/common"
	mcms_types "github.com/smartcontractkit/mcms/types"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	contract_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/adapters"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_4/operations/usdc_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/latest/cctp_message_transmitter_proxy"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/latest/factory_burn_mint_erc20"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/latest/mock_usdc_token_messenger"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/latest/mock_usdc_token_transmitter"
	"github.com/smartcontractkit/chainlink-ccip/deployment/deploy"
	"github.com/smartcontractkit/chainlink-ccip/deployment/testhelpers"

	token_pools_changeset "github.com/smartcontractkit/chainlink-ccip/deployment/token_pools"

	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"

	deploymentutils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"

	chain_selectors "github.com/smartcontractkit/chain-selectors"

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
	// e.DataStore = ds

	evmChain := e.BlockChains.EVMChains()[chainSelector]
	family, err := chain_selectors.GetSelectorFamily(chainSelector)
	if err != nil {
		t.Fatalf("Failed to get selector family for chain selector %d: %v", chainSelector, err)
	}
	evmDeployer := &adapters.EVMDeployer{}
	dReg := deploy.GetRegistry()
	dReg.RegisterDeployer(family, deploy.MCMSVersion, evmDeployer)
	cs := deploy.DeployMCMS(dReg)
	output, err := cs.Apply(*e, deploy.MCMSDeploymentConfig{
		AdapterVersion: semver.MustParse("1.0.0"),
		Chains: map[uint64]deploy.MCMSDeploymentConfigPerChain{
			chainSelector: {
				Canceller:        testhelpers.SingleGroupMCMS(),
				Bypasser:         testhelpers.SingleGroupMCMS(),
				Proposer:         testhelpers.SingleGroupMCMS(),
				TimelockMinDelay: big.NewInt(0),
				Qualifier:        ptr.String("test"),
				TimelockAdmin:    evmChain.DeployerKey.From,
			},
		},
	})
	require.NoError(t, err)

	// store addresses in ds
	allAddrRefs, err := output.DataStore.Addresses().Fetch()
	require.NoError(t, err)
	timelockAddrs := make(map[uint64]string)
	for _, addrRef := range allAddrRefs {
		require.NoError(t, ds.Addresses().Add(addrRef))
		if addrRef.Type == datastore.ContractType(deploymentutils.RBACTimelock) {
			timelockAddrs[chainSelector] = addrRef.Address
			break
		}
	}
	// Update the env datastore from the datastore
	e.DataStore = ds.Seal()

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
		TypeAndVersion: deployment.NewTypeAndVersion(usdc_token_pool.ContractType, *semver.MustParse("1.6.4")),
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
		Version:       semver.MustParse("1.6.4"),
		ChainSelector: chainSelector,
		Address:       usdcTokenPoolRef.Address,
	}))
	e.DataStore = ds.Seal()
	// transfer ownership of USDCTokenPool to respective MCMS
	transferOwnershipInput := deploy.TransferOwnershipInput{
		ChainInputs: []deploy.TransferOwnershipPerChainInput{
			{
				ChainSelector: chainSelector,
				ContractRef: []datastore.AddressRef{
					{
						Type:          datastore.ContractType(usdc_token_pool.ContractType),
						Version:       semver.MustParse("1.6.4"),
						Address:       usdcTokenPoolRef.Address,
						ChainSelector: chainSelector,
					},
				},
				CurrentOwner:  evmChain.DeployerKey.From.String(),
				ProposedOwner: timelockAddrs[chainSelector],
			},
		},
		AdapterVersion: semver.MustParse("1.0.0"),
		MCMS: mcms.Input{
			OverridePreviousRoot: false,
			ValidUntil:           3759765795,
			TimelockDelay:        mcms_types.MustParseDuration("0s"),
			TimelockAction:       mcms_types.TimelockActionSchedule,
			Qualifier:            "test",
			Description:          "Transfer ownership of USDCTokenPool to MCMS",
		},
	}

	// register chain adapter and deploy token, transmitter, messenger, and proxy
	cr := deploy.GetTransferOwnershipRegistry()
	evmAdapter := &adapters.EVMTransferOwnershipAdapter{}
	cr.RegisterAdapter(family, transferOwnershipInput.AdapterVersion, evmAdapter)
	mcmsRegistry := changesets.GetRegistry()
	evmMCMSReader := &adapters.EVMMCMSReader{}
	mcmsRegistry.RegisterMCMSReader(family, evmMCMSReader)
	transferOwnershipChangeset := deploy.TransferOwnershipChangeset(cr, mcmsRegistry)
	output, err = transferOwnershipChangeset.Apply(*e, transferOwnershipInput)
	require.NoError(t, err)
	require.Greater(t, len(output.Reports), 0)
	require.Equal(t, 1, len(output.MCMSTimelockProposals))
	testhelpers.ProcessTimelockProposals(t, *e, output.MCMSTimelockProposals, false)
	t.Logf("Transferred ownership of USDCTokenPool to respective MCMS")

	// Create the input to the changeset
	setDomainsInput := token_pools_changeset.SetDomainsInput{
		ChainInputs: []token_pools_changeset.SetDomainsPerChainInput{
			{
				ChainSelector: chainSelector,
				Address:       common.HexToAddress(usdcTokenPoolRef.Address),
				Domains:       []usdc_token_pool.DomainUpdate{},
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
	setDomainsChangeset := token_pools_changeset.SetDomainsChangeset(mcmsRegistry)
	// Apply the changeset
	output, err = setDomainsChangeset.Apply(*e, setDomainsInput)

	// Check that the changeset applied successfully
	require.NoError(t, err, "SetDomainsChangeset should not error")
	require.Greater(t, len(output.Reports), 0)
	require.Equal(t, 1, len(output.MCMSTimelockProposals))
	testhelpers.ProcessTimelockProposals(t, *e, output.MCMSTimelockProposals, false)
}
