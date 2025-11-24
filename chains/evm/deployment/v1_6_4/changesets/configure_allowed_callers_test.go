package changesets_test

import (
	"math/big"
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	contract_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_4/changesets"
	erc20_lock_box "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_4/operations/erc20_lock_box"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/latest/burn_mint_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/latest/factory_burn_mint_erc20"
	token_admin_registry_bindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_5_0/token_admin_registry"
	changesets_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	mcms_types "github.com/smartcontractkit/mcms/types"

	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"

	chain_selectors "github.com/smartcontractkit/chain-selectors"
	erc20_lock_box_bindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_4/erc20_lock_box"

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
	tokenAdminRegistryAddr, tx, tokenAdminRegistry, err := token_admin_registry_bindings.DeployTokenAdminRegistry(evmChain.DeployerKey,
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
			TokenAdminRegistry: tokenAdminRegistryAddr,
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
	mcmsRegistry := changesets_utils.GetRegistry()

	// Configure the input to the changeset for configure the allowed callers
	configureAllowedCallersInput := changesets.ConfigureAllowedCallersInput{
		ChainInputs: []changesets.ConfigureAllowedCallersPerChainInput{
			{
				ChainSelector: chainSelector,
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

	configureAllowedCallersChangeset := changesets.ConfigureAllowedCallersChangeset(mcmsRegistry)
	output, err := configureAllowedCallersChangeset.Apply(*e, configureAllowedCallersInput)
	require.NoError(t, err, "ConfigureAllowedCallersChangeset should not error")
	require.Greater(t, len(output.Reports), 0)

	// Verify the allowed callers
	erc20LockBox, err := erc20_lock_box_bindings.NewERC20LockBox(common.HexToAddress(erc20LockBoxRef.Address), evmChain.Client)
	require.NoError(t, err, "Failed to create ERC20LockBox")
	isAllowedCaller, err := erc20LockBox.IsAllowedCaller(&bind.CallOpts{Context: t.Context()}, tokenAddress, evmChain.DeployerKey.From)
	require.NoError(t, err, "Failed to check if allowed caller")
	require.True(t, isAllowedCaller, "Caller should be allowed")
}
