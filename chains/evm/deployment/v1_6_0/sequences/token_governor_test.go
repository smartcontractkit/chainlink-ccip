package sequences

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	chain_selectors "github.com/smartcontractkit/chain-selectors"
	"github.com/stretchr/testify/require"

	tg_bindings "github.com/smartcontractkit/ccip-contract-examples/chains/evm/gobindings/generated/latest/token_governor"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/token_governor"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/onchain"
	bnm_bindings "github.com/smartcontractkit/chainlink-evm/gethwrappers/shared/generated/initial/burn_mint_erc20"

	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

// deployTestTokenForGovernor deploys a test BurnMintERC20 token and returns its address
func deployTestTokenForGovernor(t *testing.T, chain evm.Chain, symbol string, decimals uint8) common.Address {
	tokenAddr, tx, _, err := bnm_bindings.DeployBurnMintERC20(
		chain.DeployerKey,
		chain.Client,
		"Test Token",
		symbol,
		decimals,
		big.NewInt(0).Mul(big.NewInt(1e9), big.NewInt(1e18)), // maxSupply: 1 billion
		big.NewInt(0), // preMint: 0
	)
	require.NoError(t, err, "Failed to deploy test token")
	_, err = deployment.ConfirmIfNoError(chain, tx, err)
	require.NoError(t, err, "Failed to confirm test token deployment")
	return tokenAddr
}

// TestDeployTokenGovernor tests the DeployTokenGovernor sequence
func TestDeployTokenGovernor(t *testing.T) {
	t.Parallel()

	evmChains := []uint64{
		chain_selectors.ETHEREUM_MAINNET.Selector,
	}

	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, evmChains),
	)
	require.NoError(t, err, "Failed to create test environment")
	require.NotNil(t, e, "Environment should be created")

	chainSelector := chain_selectors.ETHEREUM_MAINNET.Selector
	chain := e.BlockChains.EVMChains()[chainSelector]

	// Deploy a test token first
	tokenSymbol := "TGOV"
	tokenAddr := deployTestTokenForGovernor(t, chain, tokenSymbol, 18)
	t.Logf("Deployed test token at: %s", tokenAddr.Hex())

	// Build input for DeployTokenGovernor sequence
	input := DeployTokenGovernorInput{
		Token:               tokenAddr.Hex(),
		InitialDelay:        big.NewInt(0), // No delay for testing
		InitialDefaultAdmin: chain.DeployerKey.From.Hex(),
		ChainSelector:       chainSelector,
		ExistingDataStore:   e.DataStore,
	}

	// Execute the sequence
	report, err := cldf_ops.ExecuteSequence(e.OperationsBundle, DeployTokenGovernor, e.BlockChains, input)
	require.NoError(t, err, "Failed to execute DeployTokenGovernor sequence")
	require.NotNil(t, report, "Report should not be nil")

	// Verify the token governor was deployed
	require.GreaterOrEqual(t, len(report.Output.Addresses), 1, "Should have at least one address deployed")

	// Find the token governor in the output
	var tgRef datastore.AddressRef
	tgFound := false
	for _, addr := range report.Output.Addresses {
		if addr.Type == datastore.ContractType(token_governor.ContractType) {
			tgRef = addr
			tgFound = true
			break
		}
	}
	require.True(t, tgFound, "TokenGovernor should be found in deployed addresses")
	require.NotEmpty(t, tgRef.Address, "TokenGovernor address should not be empty")
	t.Logf("Deployed TokenGovernor at address: %s", tgRef.Address)

	// Make on-chain calls to verify the token governor
	tgAddr := common.HexToAddress(tgRef.Address)
	tg, err := tg_bindings.NewTokenGovernor(tgAddr, chain.Client)
	require.NoError(t, err, "Failed to create TokenGovernor binding")

	// Verify default admin
	defaultAdmin, err := tg.DefaultAdmin(&bind.CallOpts{})
	require.NoError(t, err, "Failed to get default admin")
	require.Equal(t, chain.DeployerKey.From, defaultAdmin, "Default admin should be deployer")
	t.Logf("  DefaultAdmin: %s", defaultAdmin.Hex())

	// Verify owner
	owner, err := tg.Owner(&bind.CallOpts{})
	require.NoError(t, err, "Failed to get owner")
	require.Equal(t, chain.DeployerKey.From, owner, "Owner should be deployer")
	t.Logf("  Owner: %s", owner.Hex())
}

// TestGrantRevokeRenounceRole tests the GrantRole, RevokeRole, and RenounceRole sequences
func TestGrantRevokeRenounceRole(t *testing.T) {
	t.Parallel()

	evmChains := []uint64{
		chain_selectors.ETHEREUM_MAINNET.Selector,
	}

	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, evmChains),
	)
	require.NoError(t, err, "Failed to create test environment")

	chainSelector := chain_selectors.ETHEREUM_MAINNET.Selector
	chain := e.BlockChains.EVMChains()[chainSelector]
	ds := datastore.NewMemoryDataStore()

	// Deploy a test token
	tokenSymbol := "TROLE"
	tokenAddr := deployTestTokenForGovernor(t, chain, tokenSymbol, 18)
	t.Logf("Deployed test token at: %s", tokenAddr.Hex())

	// Deploy token governor
	deployInput := DeployTokenGovernorInput{
		Token:               tokenAddr.Hex(),
		InitialDelay:        big.NewInt(0),
		InitialDefaultAdmin: chain.DeployerKey.From.Hex(),
		ChainSelector:       chainSelector,
		ExistingDataStore:   e.DataStore,
	}
	deployReport, err := cldf_ops.ExecuteSequence(e.OperationsBundle, DeployTokenGovernor, e.BlockChains, deployInput)
	require.NoError(t, err, "Failed to deploy TokenGovernor")

	// Save the token governor address to datastore
	for _, addr := range deployReport.Output.Addresses {
		err = ds.Addresses().Add(addr)
		require.NoError(t, err, "Failed to save address to datastore")
	}

	// Seal and set the datastore for subsequent sequences
	e.DataStore = ds.Seal()

	// Get the token governor address
	tgAddr, err := GetTokenGovernor(e.DataStore, chainSelector, tokenSymbol)
	require.NoError(t, err, "Failed to get TokenGovernor from datastore")

	// Create bindings
	tg, err := tg_bindings.NewTokenGovernor(tgAddr, chain.Client)
	require.NoError(t, err, "Failed to create TokenGovernor binding")

	// Define a test account to grant/revoke roles
	testAccount := common.HexToAddress("0x1111111111111111111111111111111111111111")

	// Test GrantRole - grant MINTER_ROLE to testAccount
	t.Run("GrantRole", func(t *testing.T) {
		grantInput := TokenGovernorRoleInput{
			Tokens: map[uint64]map[string]TokenGovernorGrantRole{
				chainSelector: {
					tokenSymbol: {
						Role:    RoleMinter,
						Account: testAccount,
					},
				},
			},
			ExistingDataStore: e.DataStore,
		}

		_, err := cldf_ops.ExecuteSequence(e.OperationsBundle, GrantRole, e.BlockChains, grantInput)
		require.NoError(t, err, "Failed to execute GrantRole")

		// Verify role was granted
		minterRole, err := tg.MINTERROLE(&bind.CallOpts{})
		require.NoError(t, err)
		hasRole, err := tg.HasRole(&bind.CallOpts{}, minterRole, testAccount)
		require.NoError(t, err)
		require.True(t, hasRole, "Test account should have MINTER_ROLE after grant")
		t.Logf("Successfully granted MINTER_ROLE to %s", testAccount.Hex())
	})

	// Test RevokeRole - revoke MINTER_ROLE from testAccount
	t.Run("RevokeRole", func(t *testing.T) {
		revokeInput := TokenGovernorRoleInput{
			Tokens: map[uint64]map[string]TokenGovernorGrantRole{
				chainSelector: {
					tokenSymbol: {
						Role:    RoleMinter,
						Account: testAccount,
					},
				},
			},
			ExistingDataStore: e.DataStore,
		}

		_, err := cldf_ops.ExecuteSequence(e.OperationsBundle, RevokeRole, e.BlockChains, revokeInput)
		require.NoError(t, err, "Failed to execute RevokeRole")

		// Verify role was revoked
		minterRole, err := tg.MINTERROLE(&bind.CallOpts{})
		require.NoError(t, err)
		hasRole, err := tg.HasRole(&bind.CallOpts{}, minterRole, testAccount)
		require.NoError(t, err)
		require.False(t, hasRole, "Test account should NOT have MINTER_ROLE after revoke")
		t.Logf("Successfully revoked MINTER_ROLE from %s", testAccount.Hex())
	})

	// Test RenounceRole - deployer renounces MINTER_ROLE
	t.Run("RenounceRole", func(t *testing.T) {
		// First grant MINTER role to deployer so we can renounce it
		grantInput := TokenGovernorRoleInput{
			Tokens: map[uint64]map[string]TokenGovernorGrantRole{
				chainSelector: {
					tokenSymbol: {
						Role:    RoleMinter,
						Account: chain.DeployerKey.From,
					},
				},
			},
			ExistingDataStore: e.DataStore,
		}
		_, err := cldf_ops.ExecuteSequence(e.OperationsBundle, GrantRole, e.BlockChains, grantInput)
		require.NoError(t, err, "Failed to grant role to deployer")

		// Verify role was granted
		minterRole, err := tg.MINTERROLE(&bind.CallOpts{})
		require.NoError(t, err)
		hasRoleBefore, err := tg.HasRole(&bind.CallOpts{}, minterRole, chain.DeployerKey.From)
		require.NoError(t, err)
		require.True(t, hasRoleBefore, "Deployer should have MINTER_ROLE before renounce")

		// Now renounce the role
		renounceInput := TokenGovernorRoleInput{
			Tokens: map[uint64]map[string]TokenGovernorGrantRole{
				chainSelector: {
					tokenSymbol: {
						Role:    RoleMinter,
						Account: chain.DeployerKey.From,
					},
				},
			},
			ExistingDataStore: e.DataStore,
		}
		_, err = cldf_ops.ExecuteSequence(e.OperationsBundle, RenounceRole, e.BlockChains, renounceInput)
		require.NoError(t, err, "Failed to execute RenounceRole")

		// Verify role was renounced
		hasRoleAfter, err := tg.HasRole(&bind.CallOpts{}, minterRole, chain.DeployerKey.From)
		require.NoError(t, err)
		require.False(t, hasRoleAfter, "Deployer should NOT have MINTER_ROLE after renounce")
		t.Logf("Successfully renounced MINTER_ROLE by %s", chain.DeployerKey.From.Hex())
	})
}

// TestTransferAndAcceptOwnership tests the full TransferOwnership and AcceptOwnership flow
// Uses an additional account to properly test the accept step
func TestTransferAndAcceptOwnership(t *testing.T) {
	t.Parallel()

	evmChains := []uint64{
		chain_selectors.ETHEREUM_MAINNET.Selector,
	}

	e, err := environment.New(t.Context(),
		environment.WithEVMSimulatedWithConfig(t, evmChains, onchain.EVMSimLoaderConfig{
			NumAdditionalAccounts: 1,
			BlockTime:             1,
		}))
	require.NoError(t, err, "Failed to create test environment")

	chainSelector := chain_selectors.ETHEREUM_MAINNET.Selector
	chain := e.BlockChains.EVMChains()[chainSelector]
	ds := datastore.NewMemoryDataStore()

	// Deploy a test token
	tokenSymbol := "TOWN"
	tokenAddr := deployTestTokenForGovernor(t, chain, tokenSymbol, 18)
	t.Logf("Deployed test token at: %s", tokenAddr.Hex())

	// Deploy token governor
	deployInput := DeployTokenGovernorInput{
		Token:               tokenAddr.Hex(),
		InitialDelay:        big.NewInt(0),
		InitialDefaultAdmin: chain.DeployerKey.From.Hex(),
		ChainSelector:       chainSelector,
		ExistingDataStore:   e.DataStore,
	}
	deployReport, err := cldf_ops.ExecuteSequence(e.OperationsBundle, DeployTokenGovernor, e.BlockChains, deployInput)
	require.NoError(t, err, "Failed to deploy TokenGovernor")

	// Save to datastore
	for _, addr := range deployReport.Output.Addresses {
		err = ds.Addresses().Add(addr)
		require.NoError(t, err)
	}
	e.DataStore = ds.Seal()

	// Get the token governor
	tgAddr, err := GetTokenGovernor(e.DataStore, chainSelector, tokenSymbol)
	require.NoError(t, err)
	tg, err := tg_bindings.NewTokenGovernor(tgAddr, chain.Client)
	require.NoError(t, err)

	// Verify initial owner
	initialOwner, err := tg.Owner(&bind.CallOpts{})
	require.NoError(t, err)
	require.Equal(t, chain.DeployerKey.From, initialOwner, "Initial owner should be deployer")
	t.Logf("Initial owner: %s", initialOwner.Hex())

	// Get the new owner from additional accounts
	newOwnerAddr := chain.Users[0].From
	t.Logf("New owner address: %s", newOwnerAddr.Hex())

	// Step 1: Transfer ownership to the new owner
	transferInput := TokenGovernorOwnershipInput{
		Tokens: map[uint64]map[string]common.Address{
			chainSelector: {
				tokenSymbol: newOwnerAddr,
			},
		},
		ExistingDataStore: e.DataStore,
	}

	_, err = cldf_ops.ExecuteSequence(e.OperationsBundle, TransferOwnership, e.BlockChains, transferInput)
	require.NoError(t, err, "Failed to execute TransferOwnership")
	t.Logf("Step 1: Ownership transfer initiated to: %s", newOwnerAddr.Hex())

	// Step 2: Accept ownership using the new owner's transact opts
	newOwnerOpts := chain.Users[0]
	newOwnerOpts.Context = t.Context()

	acceptTx, err := tg.AcceptOwnership(newOwnerOpts)
	if err != nil {
		t.Logf("AcceptOwnership error details: %+v", err)
		t.Logf("New owner address: %s", newOwnerAddr.Hex())
		t.Logf("Transaction from: %s", newOwnerOpts.From.Hex())
		require.NoError(t, err, "Failed to call AcceptOwnership")
	}

	_, err = chain.Confirm(acceptTx)
	require.NoError(t, err, "Failed to confirm AcceptOwnership transaction")
	t.Logf("Step 2: AcceptOwnership executed successfully")

	// Verify the owner has changed
	newOwner, err := tg.Owner(&bind.CallOpts{})
	require.NoError(t, err)
	require.Equal(t, newOwnerAddr, newOwner, "Owner should be the new owner")
	t.Logf("  New owner: %s", newOwner.Hex())

	// Verify the old owner is no longer the owner
	require.NotEqual(t, initialOwner, newOwner, "Old owner should no longer be owner")
	t.Logf("  Previous owner (%s) is no longer owner", initialOwner.Hex())
}

// TestTransferOwnership_SameOwner tests that TransferOwnership fails when trying to transfer to the same owner.
func TestTransferOwnership_SameOwner(t *testing.T) {
	t.Parallel()

	evmChains := []uint64{
		chain_selectors.ETHEREUM_MAINNET.Selector,
	}

	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, evmChains),
	)
	require.NoError(t, err, "Failed to create test environment")

	chainSelector := chain_selectors.ETHEREUM_MAINNET.Selector
	chain := e.BlockChains.EVMChains()[chainSelector]
	ds := datastore.NewMemoryDataStore()

	// Deploy a test token
	tokenSymbol := "TOWN2"
	tokenAddr := deployTestTokenForGovernor(t, chain, tokenSymbol, 18)
	t.Logf("Deployed test token at: %s", tokenAddr.Hex())

	// Deploy token governor
	deployInput := DeployTokenGovernorInput{
		Token:               tokenAddr.Hex(),
		InitialDelay:        big.NewInt(0),
		InitialDefaultAdmin: chain.DeployerKey.From.Hex(),
		ChainSelector:       chainSelector,
		ExistingDataStore:   e.DataStore,
	}
	deployReport, err := cldf_ops.ExecuteSequence(e.OperationsBundle, DeployTokenGovernor, e.BlockChains, deployInput)
	require.NoError(t, err, "Failed to deploy TokenGovernor")

	// Save to datastore
	for _, addr := range deployReport.Output.Addresses {
		err = ds.Addresses().Add(addr)
		require.NoError(t, err)
	}
	e.DataStore = ds.Seal()

	// Test validation: Transfer ownership to same owner should fail
	transferInputSameOwner := TokenGovernorOwnershipInput{
		Tokens: map[uint64]map[string]common.Address{
			chainSelector: {
				tokenSymbol: chain.DeployerKey.From, // Same as current owner
			},
		},
		ExistingDataStore: e.DataStore,
	}

	_, err = cldf_ops.ExecuteSequence(e.OperationsBundle, TransferOwnership, e.BlockChains, transferInputSameOwner)
	require.Error(t, err, "Should fail when new owner is already the current owner")
	require.Contains(t, err.Error(), "already the current owner")
	t.Logf("Correctly rejected transfer to current owner: %v", err)
}

// TestBeginAndAcceptDefaultAdminTransfer tests the full BeginDefaultAdminTransfer and AcceptDefaultAdminTransfer flow
// Uses an additional account to properly test the accept step
func TestBeginAndAcceptDefaultAdminTransfer(t *testing.T) {
	t.Parallel()

	evmChains := []uint64{
		chain_selectors.ETHEREUM_MAINNET.Selector,
	}

	e, err := environment.New(t.Context(),
		environment.WithEVMSimulatedWithConfig(t, evmChains, onchain.EVMSimLoaderConfig{
			NumAdditionalAccounts: 1,
			BlockTime:             1,
		}))
	require.NoError(t, err, "Failed to create test environment")

	chainSelector := chain_selectors.ETHEREUM_MAINNET.Selector
	chain := e.BlockChains.EVMChains()[chainSelector]
	ds := datastore.NewMemoryDataStore()

	// Deploy a test token
	tokenSymbol := "TADMIN"
	tokenAddr := deployTestTokenForGovernor(t, chain, tokenSymbol, 18)
	t.Logf("Deployed test token at: %s", tokenAddr.Hex())

	// Deploy token governor with 0 delay for testing
	// Initial admin is the deployer
	deployInput := DeployTokenGovernorInput{
		Token:               tokenAddr.Hex(),
		InitialDelay:        big.NewInt(0), // No delay for testing
		InitialDefaultAdmin: chain.DeployerKey.From.Hex(),
		ChainSelector:       chainSelector,
		ExistingDataStore:   e.DataStore,
	}
	deployReport, err := cldf_ops.ExecuteSequence(e.OperationsBundle, DeployTokenGovernor, e.BlockChains, deployInput)
	require.NoError(t, err, "Failed to deploy TokenGovernor")

	// Save to datastore
	for _, addr := range deployReport.Output.Addresses {
		err = ds.Addresses().Add(addr)
		require.NoError(t, err)
	}
	e.DataStore = ds.Seal()

	// Get the token governor
	tgAddr, err := GetTokenGovernor(e.DataStore, chainSelector, tokenSymbol)
	require.NoError(t, err)
	tg, err := tg_bindings.NewTokenGovernor(tgAddr, chain.Client)
	require.NoError(t, err)

	// Verify initial default admin
	initialAdmin, err := tg.DefaultAdmin(&bind.CallOpts{})
	require.NoError(t, err)
	require.Equal(t, chain.DeployerKey.From, initialAdmin, "Initial default admin should be deployer")
	t.Logf("Initial default admin: %s", initialAdmin.Hex())
	newAdminAddr := chain.Users[0].From

	// Step 1: Begin default admin transfer to the new admin address
	beginInput := TokenGovernorOwnershipInput{
		Tokens: map[uint64]map[string]common.Address{
			chainSelector: {
				tokenSymbol: newAdminAddr,
			},
		},
		ExistingDataStore: e.DataStore,
	}

	_, err = cldf_ops.ExecuteSequence(e.OperationsBundle, BeginDefaultAdminTransfer, e.BlockChains, beginInput)
	require.NoError(t, err, "Failed to execute BeginDefaultAdminTransfer")
	t.Logf("Step 1: Default admin transfer initiated to: %s", newAdminAddr.Hex())

	// Verify pending default admin was set
	pendingAdmin, err := tg.PendingDefaultAdmin(&bind.CallOpts{})
	require.NoError(t, err)
	require.Equal(t, newAdminAddr, pendingAdmin.NewAdmin, "Pending default admin should be set to new admin")
	t.Logf("  Pending default admin: %s", pendingAdmin.NewAdmin.Hex())
	t.Logf("  Pending admin schedule: %d", pendingAdmin.Schedule)

	// Step 2: Accept default admin transfer using the new admin's transact opts
	secondUserOpts := chain.Users[0]
	secondUserOpts.Context = t.Context()

	// Now try the actual transaction
	acceptTx, err := tg.AcceptDefaultAdminTransfer(secondUserOpts)
	if err != nil {
		t.Logf("AcceptDefaultAdminTransfer error details: %+v", err)
		t.Logf("NewAdmin address: %s", newAdminAddr.Hex())
		t.Logf("Transaction from: %s", secondUserOpts.From.Hex())
		t.Logf("Chain Selector used: %d", chain.Selector)
		require.NoError(t, err, "Failed to call AcceptDefaultAdminTransfer")
	}

	_, err = chain.Confirm(acceptTx)
	require.NoError(t, err, "Failed to confirm AcceptDefaultAdminTransfer transaction")
	t.Logf("Step 2: AcceptDefaultAdminTransfer executed successfully")

	// Verify the default admin has changed
	newDefaultAdmin, err := tg.DefaultAdmin(&bind.CallOpts{})
	require.NoError(t, err)
	require.Equal(t, newAdminAddr, newDefaultAdmin, "Default admin should be the new admin")
	t.Logf("  New default admin: %s", newDefaultAdmin.Hex())

	// Verify the old admin is no longer the default admin
	require.NotEqual(t, initialAdmin, newDefaultAdmin, "Old admin should no longer be default admin")
	t.Logf("  Previous admin (%s) is no longer default admin", initialAdmin.Hex())
}

// TestBeginDefaultAdminTransfer_AlreadyAdmin tests that BeginDefaultAdminTransfer fails when new admin is current admin
func TestBeginDefaultAdminTransfer_AlreadyAdmin(t *testing.T) {
	t.Parallel()

	evmChains := []uint64{
		chain_selectors.ETHEREUM_MAINNET.Selector,
	}

	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, evmChains),
	)
	require.NoError(t, err, "Failed to create test environment")

	chainSelector := chain_selectors.ETHEREUM_MAINNET.Selector
	chain := e.BlockChains.EVMChains()[chainSelector]
	ds := datastore.NewMemoryDataStore()

	// Deploy a test token
	tokenSymbol := "TSADM"
	tokenAddr := deployTestTokenForGovernor(t, chain, tokenSymbol, 18)

	// Deploy token governor
	deployInput := DeployTokenGovernorInput{
		Token:               tokenAddr.Hex(),
		InitialDelay:        big.NewInt(0),
		InitialDefaultAdmin: chain.DeployerKey.From.Hex(),
		ChainSelector:       chainSelector,
		ExistingDataStore:   e.DataStore,
	}
	deployReport, err := cldf_ops.ExecuteSequence(e.OperationsBundle, DeployTokenGovernor, e.BlockChains, deployInput)
	require.NoError(t, err, "Failed to deploy TokenGovernor")

	for _, addr := range deployReport.Output.Addresses {
		err = ds.Addresses().Add(addr)
		require.NoError(t, err)
	}
	e.DataStore = ds.Seal()

	// Try to begin admin transfer to the current admin (should fail)
	beginInput := TokenGovernorOwnershipInput{
		Tokens: map[uint64]map[string]common.Address{
			chainSelector: {
				tokenSymbol: chain.DeployerKey.From, // Same as current admin
			},
		},
		ExistingDataStore: e.DataStore,
	}

	_, err = cldf_ops.ExecuteSequence(e.OperationsBundle, BeginDefaultAdminTransfer, e.BlockChains, beginInput)
	require.Error(t, err, "Should fail when new admin is already the current default admin")
	require.Contains(t, err.Error(), "already the current default admin")
	t.Logf("Correctly rejected transfer to current admin: %v", err)
}

// TestGrantRole_AlreadyHasRole tests that GrantRole fails when account already has the role
func TestGrantRole_AlreadyHasRole(t *testing.T) {
	t.Parallel()

	evmChains := []uint64{
		chain_selectors.ETHEREUM_MAINNET.Selector,
	}

	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, evmChains),
	)
	require.NoError(t, err, "Failed to create test environment")

	chainSelector := chain_selectors.ETHEREUM_MAINNET.Selector
	chain := e.BlockChains.EVMChains()[chainSelector]
	ds := datastore.NewMemoryDataStore()

	// Deploy a test token
	tokenSymbol := "TDUP"
	tokenAddr := deployTestTokenForGovernor(t, chain, tokenSymbol, 18)

	// Deploy token governor
	deployInput := DeployTokenGovernorInput{
		Token:               tokenAddr.Hex(),
		InitialDelay:        big.NewInt(0),
		InitialDefaultAdmin: chain.DeployerKey.From.Hex(),
		ChainSelector:       chainSelector,
		ExistingDataStore:   e.DataStore,
	}
	deployReport, err := cldf_ops.ExecuteSequence(e.OperationsBundle, DeployTokenGovernor, e.BlockChains, deployInput)
	require.NoError(t, err, "Failed to deploy TokenGovernor")

	for _, addr := range deployReport.Output.Addresses {
		err = ds.Addresses().Add(addr)
		require.NoError(t, err)
	}
	e.DataStore = ds.Seal()

	testAccount := common.HexToAddress("0x4444444444444444444444444444444444444444")

	// Grant role first time
	grantInput := TokenGovernorRoleInput{
		Tokens: map[uint64]map[string]TokenGovernorGrantRole{
			chainSelector: {
				tokenSymbol: {
					Role:    RoleMinter,
					Account: testAccount,
				},
			},
		},
		ExistingDataStore: e.DataStore,
	}
	_, err = cldf_ops.ExecuteSequence(e.OperationsBundle, GrantRole, e.BlockChains, grantInput)
	require.NoError(t, err, "First grant should succeed")

	// Try to grant the same role again (should fail)
	_, err = cldf_ops.ExecuteSequence(e.OperationsBundle, GrantRole, e.BlockChains, grantInput)
	require.Error(t, err, "Should fail when account already has the role")
	require.Contains(t, err.Error(), "already has role")
	t.Logf("Correctly rejected duplicate grant: %v", err)
}

// TestRevokeRole_DoesNotHaveRole tests that RevokeRole fails when account doesn't have the role
func TestRevokeRole_DoesNotHaveRole(t *testing.T) {
	t.Parallel()

	evmChains := []uint64{
		chain_selectors.ETHEREUM_MAINNET.Selector,
	}

	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, evmChains),
	)
	require.NoError(t, err, "Failed to create test environment")

	chainSelector := chain_selectors.ETHEREUM_MAINNET.Selector
	chain := e.BlockChains.EVMChains()[chainSelector]
	ds := datastore.NewMemoryDataStore()

	// Deploy a test token
	tokenSymbol := "TNOROLE"
	tokenAddr := deployTestTokenForGovernor(t, chain, tokenSymbol, 18)

	// Deploy token governor
	deployInput := DeployTokenGovernorInput{
		Token:               tokenAddr.Hex(),
		InitialDelay:        big.NewInt(0),
		InitialDefaultAdmin: chain.DeployerKey.From.Hex(),
		ChainSelector:       chainSelector,
		ExistingDataStore:   e.DataStore,
	}
	deployReport, err := cldf_ops.ExecuteSequence(e.OperationsBundle, DeployTokenGovernor, e.BlockChains, deployInput)
	require.NoError(t, err, "Failed to deploy TokenGovernor")

	for _, addr := range deployReport.Output.Addresses {
		err = ds.Addresses().Add(addr)
		require.NoError(t, err)
	}
	e.DataStore = ds.Seal()

	testAccount := common.HexToAddress("0x5555555555555555555555555555555555555555")

	// Try to revoke role that account doesn't have (should fail)
	revokeInput := TokenGovernorRoleInput{
		Tokens: map[uint64]map[string]TokenGovernorGrantRole{
			chainSelector: {
				tokenSymbol: {
					Role:    RoleMinter,
					Account: testAccount,
				},
			},
		},
		ExistingDataStore: e.DataStore,
	}
	_, err = cldf_ops.ExecuteSequence(e.OperationsBundle, RevokeRole, e.BlockChains, revokeInput)
	require.Error(t, err, "Should fail when account doesn't have the role")
	require.Contains(t, err.Error(), "doesn't has role")
	t.Logf("Correctly rejected revoke for non-holder: %v", err)
}

// TestGetRoleFromTokenGovernor tests all role mappings
func TestGetRoleFromTokenGovernor(t *testing.T) {
	t.Parallel()

	evmChains := []uint64{
		chain_selectors.ETHEREUM_MAINNET.Selector,
	}

	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, evmChains),
	)
	require.NoError(t, err, "Failed to create test environment")

	chainSelector := chain_selectors.ETHEREUM_MAINNET.Selector
	chain := e.BlockChains.EVMChains()[chainSelector]

	// Deploy a test token
	tokenSymbol := "TROLES"
	tokenAddr := deployTestTokenForGovernor(t, chain, tokenSymbol, 18)

	// Deploy token governor
	deployInput := DeployTokenGovernorInput{
		Token:               tokenAddr.Hex(),
		InitialDelay:        big.NewInt(0),
		InitialDefaultAdmin: chain.DeployerKey.From.Hex(),
		ChainSelector:       chainSelector,
		ExistingDataStore:   e.DataStore,
	}
	deployReport, err := cldf_ops.ExecuteSequence(e.OperationsBundle, DeployTokenGovernor, e.BlockChains, deployInput)
	require.NoError(t, err, "Failed to deploy TokenGovernor")

	// Find the token governor address
	var tgAddr common.Address
	for _, addr := range deployReport.Output.Addresses {
		if addr.Type == datastore.ContractType(token_governor.ContractType) {
			tgAddr = common.HexToAddress(addr.Address)
			break
		}
	}
	require.NotEqual(t, common.Address{}, tgAddr, "TokenGovernor address should be found")

	// Create bindings
	tg, err := tg_bindings.NewTokenGovernor(tgAddr, chain.Client)
	require.NoError(t, err, "Failed to create TokenGovernor binding")

	// Test all roles
	testCases := []struct {
		name string
		role TokenGovernorRole
	}{
		{"RoleMinter", RoleMinter},
		{"RoleBridgerMinterOrBurner", RoleBridgerMinterOrBurner},
		{"RoleBurner", RoleBurner},
		{"RoleFreezer", RoleFreezer},
		{"RoleUnfreezer", RoleUnfreezer},
		{"RolePauser", RolePauser},
		{"RoleUnpauser", RoleUnpauser},
		{"RoleRecovery", RoleRecovery},
		{"RoleCheckerAdmin", RoleCheckerAdmin},
		{"RoleDefaultAdmin", RoleDefaultAdmin},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			roleBytes, err := GetRoleFromTokenGovernor(t.Context(), tg, tc.role)
			require.NoError(t, err, "Failed to get role %s", tc.name)
			// Note: DEFAULT_ADMIN_ROLE is 0x00...00 by design in OpenZeppelin's AccessControl
			if tc.role != RoleDefaultAdmin {
				require.NotEqual(t, [32]byte{}, roleBytes, "Role bytes should not be zero for %s", tc.name)
			}
			t.Logf("  %s: 0x%x", tc.name, roleBytes)
		})
	}
}

// TestTokenGovernorNotInDatastore tests error handling when token governor is not found
func TestTokenGovernorNotInDatastore(t *testing.T) {
	t.Parallel()

	evmChains := []uint64{
		chain_selectors.ETHEREUM_MAINNET.Selector,
	}

	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, evmChains),
	)
	require.NoError(t, err, "Failed to create test environment")

	chainSelector := chain_selectors.ETHEREUM_MAINNET.Selector

	// Try to get token governor that doesn't exist
	_, err = GetTokenGovernor(e.DataStore, chainSelector, "NONEXISTENT")
	require.Error(t, err, "Should fail when token governor is not in datastore")
	require.Contains(t, err.Error(), "not found in datastore")
	t.Logf("Correctly returned error for missing token governor: %v", err)
}
