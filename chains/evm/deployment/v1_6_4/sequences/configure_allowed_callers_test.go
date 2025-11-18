package sequences_test

import (
	"math/big"
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/aws/smithy-go/ptr"
	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	contract_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/adapters"
	erc20_lock_box "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_4/operations/erc20_lock_box"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/latest/burn_mint_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/latest/factory_burn_mint_erc20"
	token_admin_registry_bindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_5_0/token_admin_registry"
	"github.com/smartcontractkit/chainlink-ccip/deployment/deploy"
	erc20_lock_box_changeset "github.com/smartcontractkit/chainlink-ccip/deployment/erc20_lock_box"
	"github.com/smartcontractkit/chainlink-ccip/deployment/testhelpers"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	mcms_types "github.com/smartcontractkit/mcms/types"

	deploymentutils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"

	chain_selectors "github.com/smartcontractkit/chain-selectors"

	"github.com/stretchr/testify/require"
)

func TestConfigureAllowedCallersSequence(t *testing.T) {
	chainSelector := uint64(chain_selectors.TEST_90000001.Selector)
	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{chainSelector}),
	)
	require.NoError(t, err, "Failed to create environment")
	require.NotNil(t, e, "Environment should be created")
	ds := datastore.NewMemoryDataStore()
	e.DataStore = ds.Seal()

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
		"TEST18",
		18,
		big.NewInt(0),             // maxSupply (0 = unlimited)
		big.NewInt(0),             // preMint
		evmChain.DeployerKey.From, // newOwner
	)
	require.NoError(t, err, "Failed to deploy FactoryBurnMintERC20 token")
	_, err = evmChain.Confirm(tx)
	require.NoError(t, err, "Failed to confirm token deployment transaction")

	// Deploy BurnMintTokenPool with tokenAddress as the token
	burnMintTokenPoolAddr, tx, _, err := burn_mint_token_pool.DeployBurnMintTokenPool(
		evmChain.DeployerKey,
		evmChain.Client,
		tokenAddress,
		18,
		[]common.Address{},
		common.Address{1}, // rmnProxy
		common.Address{2}, // router
	)
	var burn_mint_token_pool_type cldf_deployment.ContractType = "BurnMintTokenPool"
	require.NoError(t, err, "Failed to deploy BurnMintTokenPool")
	_, err = evmChain.Confirm(tx)
	require.NoError(t, err, "Failed to confirm BurnMintTokenPool deployment transaction")
	ds.Addresses().Add(datastore.AddressRef{
		Type:          datastore.ContractType(burn_mint_token_pool_type),
		Version:       semver.MustParse("1.6.4"),
		Address:       burnMintTokenPoolAddr.Hex(),
		ChainSelector: chainSelector,
	})

	// Configure the token pool with the TokenAdminRegistry using go bindings directly
	_, tx, tokenAdminRegistry, err := token_admin_registry_bindings.DeployTokenAdminRegistry(evmChain.DeployerKey,
		evmChain.Client,
	)
	require.NoError(t, err, "Failed to deploy TokenAdminRegistry")
	_, err = evmChain.Confirm(tx)
	require.NoError(t, err, "Failed to confirm TokenAdminRegistry deployment transaction")

	// Deploy ERC20LockBox using the deploy operation
	erc20LockBoxRef, err := contract_utils.MaybeDeployContract(e.OperationsBundle, erc20_lock_box.Deploy, evmChain, contract.DeployInput[erc20_lock_box.ConstructorArgs]{
		TypeAndVersion: deployment.NewTypeAndVersion(erc20_lock_box.ContractType, *semver.MustParse("1.6.4")),
		ChainSelector:  chainSelector,
		Args: erc20_lock_box.ConstructorArgs{
			TokenAdminRegistry: burnMintTokenPoolAddr,
		},
	}, nil)
	require.NoError(t, err, "Failed to deploy ERC20LockBox")
	require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
		Type:          datastore.ContractType(erc20_lock_box.ContractType),
		Version:       semver.MustParse("1.6.4"),
		ChainSelector: chainSelector,
		Address:       erc20LockBoxRef.Address,
	}))
	e.DataStore = ds.Seal()

	// Step 1: Propose administrator with the deployer key as the admin
	proposeAdminTx, err := tokenAdminRegistry.ProposeAdministrator(evmChain.DeployerKey, tokenAddress, evmChain.DeployerKey.From)
	require.NoError(t, err, "Failed to propose administrator")
	_, err = evmChain.Confirm(proposeAdminTx)
	require.NoError(t, err, "Failed to confirm propose administrator transaction")

	// Step 2: Accept the administrator role
	acceptAdminTx, err := tokenAdminRegistry.AcceptAdminRole(evmChain.DeployerKey, tokenAddress)
	require.NoError(t, err, "Failed to accept admin role")
	_, err = evmChain.Confirm(acceptAdminTx)
	require.NoError(t, err, "Failed to confirm accept admin role transaction")

	// Step 3: Set the token pool for the token
	setPoolTx, err := tokenAdminRegistry.SetPool(evmChain.DeployerKey, tokenAddress, common.HexToAddress(burnMintTokenPoolAddr.Hex()))
	require.NoError(t, err, "Failed to set token pool")
	_, err = evmChain.Confirm(setPoolTx)
	require.NoError(t, err, "Failed to confirm set token pool transaction")

	// Execute ConfigureAllowedCallers changeset
	mcmsRegistry := changesets.GetRegistry()
	evmMCMSReader := &adapters.EVMMCMSReader{}
	mcmsRegistry.RegisterMCMSReader(family, evmMCMSReader)

	// Call the transferOwnership changeset to transfer the ownership of the burnMintTokenPool to the MCMS
	transferOwnershipInput := deploy.TransferOwnershipInput{
		ChainInputs: []deploy.TransferOwnershipPerChainInput{
			{
				ChainSelector: chainSelector,
				ContractRef: []datastore.AddressRef{
					{
						Type:          datastore.ContractType(burn_mint_token_pool_type),
						Version:       semver.MustParse("1.6.4"),
						Address:       burnMintTokenPoolAddr.Hex(),
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
	cr := deploy.GetTransferOwnershipRegistry()
	evmAdapter := &adapters.EVMTransferOwnershipAdapter{}
	cr.RegisterAdapter(family, transferOwnershipInput.AdapterVersion, evmAdapter)
	transferOwnershipChangeset := deploy.TransferOwnershipChangeset(cr, mcmsRegistry)
	output, err = transferOwnershipChangeset.Apply(*e, transferOwnershipInput)
	require.NoError(t, err)
	require.Greater(t, len(output.Reports), 0)
	require.Equal(t, 1, len(output.MCMSTimelockProposals))
	testhelpers.ProcessTimelockProposals(t, *e, output.MCMSTimelockProposals, false)
	t.Logf("Transferred ownership of burnMintTokenPool to MCMS")

	// Configure the input to the changeset for configure the allowed callers
	configureAllowedCallersInput := erc20_lock_box_changeset.ConfigureAllowedCallersInput{
		ChainInputs: []erc20_lock_box_changeset.ConfigureAllowedCallersPerChainInput{
			{
				ChainSelector: chainSelector,
				Address:       common.HexToAddress(erc20LockBoxRef.Address),
				AllowedCallers: []erc20_lock_box.AllowedCallerConfigArgs{
					{
						Token:   tokenAddress,
						Caller:  evmChain.DeployerKey.From,
						Allowed: true,
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
			Description:          "Configure allowed callers on ERC20LockBox",
		},
	}

	configureAllowedCallersChangeset := erc20_lock_box_changeset.ConfigureAllowedCallersChangeset(mcmsRegistry)
	output, err = configureAllowedCallersChangeset.Apply(*e, configureAllowedCallersInput)
	require.NoError(t, err, "ConfigureAllowedCallersChangeset should not error")
	require.Greater(t, len(output.Reports), 0)
	require.Equal(t, 1, len(output.MCMSTimelockProposals))
	testhelpers.ProcessTimelockProposals(t, *e, output.MCMSTimelockProposals, false)
}
