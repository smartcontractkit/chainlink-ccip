package changesets_test

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	chain_selectors "github.com/smartcontractkit/chain-selectors"
	"github.com/stretchr/testify/require"

	tg_bindings "github.com/smartcontractkit/ccip-contract-examples/chains/evm/gobindings/generated/latest/token_governor"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/changesets"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/sequences"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	bnm_bindings "github.com/smartcontractkit/chainlink-evm/gethwrappers/shared/generated/initial/burn_mint_erc20"
)

// deployTestToken deploys a test BurnMintERC20 token and returns its address
func deployTestToken(t *testing.T, chain evm.Chain, symbol string, decimals uint8) common.Address {
	tokenAddr, tx, _, err := bnm_bindings.DeployBurnMintERC20(
		chain.DeployerKey,
		chain.Client,
		"Test Token",
		symbol,
		decimals,
		big.NewInt(0).Mul(big.NewInt(1e9), big.NewInt(1e18)), // maxSupply: 1 billion
		big.NewInt(0),                                        // preMint: 0
	)
	require.NoError(t, err, "Failed to deploy test token")
	_, err = deployment.ConfirmIfNoError(chain, tx, err)
	require.NoError(t, err, "Failed to confirm test token deployment")
	return tokenAddr
}

// TestDeployTokenGovernorChangeset tests the DeployTokenGovernorChangeset
func TestDeployTokenGovernorChangeset(t *testing.T) {
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
	tokenAddr := deployTestToken(t, chain, tokenSymbol, 18)
	t.Logf("Deployed test token at: %s", tokenAddr.Hex())

	// Create the changeset input
	input := sequences.DeployTokenGovernorInput{
		Token:               tokenAddr.Hex(),
		InitialDelay:        big.NewInt(0),
		InitialDefaultAdmin: chain.DeployerKey.From.Hex(),
		ChainSelector:       chainSelector,
	}

	// Get the changeset
	cs := changesets.DeployTokenGovernorChangeset()

	// Verify preconditions
	err = cs.VerifyPreconditions(*e, input)
	require.NoError(t, err, "VerifyPreconditions should pass")

	// Apply the changeset
	output, err := cs.Apply(*e, input)
	require.NoError(t, err, "Apply should succeed")
	require.NotNil(t, output.DataStore, "DataStore should be returned")

	// Verify the token governor was deployed by checking the datastore
	require.NotNil(t, output.DataStore, "Output DataStore should not be nil")

	// Get the sealed datastore and find the token governor
	sealedDS := output.DataStore.Seal()
	tgAddr, err := sequences.GetTokenGovernor(sealedDS, chainSelector, tokenSymbol)
	require.NoError(t, err, "Should find TokenGovernor in datastore")

	// Make on-chain calls to verify the token governor
	tg, err := tg_bindings.NewTokenGovernor(tgAddr, chain.Client)
	require.NoError(t, err, "Failed to create TokenGovernor binding")

	// Verify default admin
	defaultAdmin, err := tg.DefaultAdmin(&bind.CallOpts{Context: t.Context()})
	require.NoError(t, err, "Failed to get default admin")
	require.Equal(t, chain.DeployerKey.From, defaultAdmin, "Default admin should be deployer")
	t.Logf("TokenGovernor deployed at %s with default admin %s", tgAddr.Hex(), defaultAdmin.Hex())
}

// TestDeployTokenGovernorChangeset_Validation tests the validation of DeployTokenGovernorChangeset
func TestDeployTokenGovernorChangeset_Validation(t *testing.T) {
	t.Parallel()

	evmChains := []uint64{
		chain_selectors.ETHEREUM_MAINNET.Selector,
	}

	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, evmChains),
	)
	require.NoError(t, err, "Failed to create test environment")

	cs := changesets.DeployTokenGovernorChangeset()

	t.Run("missing chain selector", func(t *testing.T) {
		input := sequences.DeployTokenGovernorInput{
			Token:         "0x1234567890123456789012345678901234567890",
			ChainSelector: 0,
		}
		err := cs.VerifyPreconditions(*e, input)
		require.Error(t, err)
		require.Contains(t, err.Error(), "chain selector must be provided")
	})

	t.Run("missing token", func(t *testing.T) {
		input := sequences.DeployTokenGovernorInput{
			Token:         "",
			ChainSelector: chain_selectors.ETHEREUM_MAINNET.Selector,
		}
		err := cs.VerifyPreconditions(*e, input)
		require.Error(t, err)
		require.Contains(t, err.Error(), "token address must be provided")
	})

	t.Run("invalid token address", func(t *testing.T) {
		input := sequences.DeployTokenGovernorInput{
			Token:         "not-a-valid-address",
			ChainSelector: chain_selectors.ETHEREUM_MAINNET.Selector,
		}
		err := cs.VerifyPreconditions(*e, input)
		require.Error(t, err)
		require.Contains(t, err.Error(), "token address is not a valid hex address")
	})

	t.Run("invalid initial admin address", func(t *testing.T) {
		input := sequences.DeployTokenGovernorInput{
			Token:               "0x1234567890123456789012345678901234567890",
			ChainSelector:       chain_selectors.ETHEREUM_MAINNET.Selector,
			InitialDefaultAdmin: "not-a-valid-address",
		}
		err := cs.VerifyPreconditions(*e, input)
		require.Error(t, err)
		require.Contains(t, err.Error(), "initial default admin is not a valid hex address")
	})
}

// TestGrantRoleChangeset tests the GrantRoleChangeset
func TestGrantRoleChangeset(t *testing.T) {
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
	tokenSymbol := "TGRANT"
	tokenAddr := deployTestToken(t, chain, tokenSymbol, 18)
	t.Logf("Deployed test token at: %s", tokenAddr.Hex())

	// Deploy token governor first
	deployInput := sequences.DeployTokenGovernorInput{
		Token:               tokenAddr.Hex(),
		InitialDelay:        big.NewInt(0),
		InitialDefaultAdmin: chain.DeployerKey.From.Hex(),
		ChainSelector:       chainSelector,
	}
	deployCS := changesets.DeployTokenGovernorChangeset()
	deployOutput, err := deployCS.Apply(*e, deployInput)
	require.NoError(t, err, "Failed to deploy TokenGovernor")

	// Update the environment with the new datastore
	e.DataStore = deployOutput.DataStore.Seal()

	// Get the token governor address
	tgAddr, err := sequences.GetTokenGovernor(e.DataStore, chainSelector, tokenSymbol)
	require.NoError(t, err)
	tg, err := tg_bindings.NewTokenGovernor(tgAddr, chain.Client)
	require.NoError(t, err)

	// Create a new account to grant a role to
	newAccount := common.HexToAddress("0x1111111111111111111111111111111111111111")

	// Get MINTER_ROLE before granting
	minterRole, err := tg.MINTERROLE(&bind.CallOpts{Context: t.Context()})
	require.NoError(t, err)

	// Verify the account doesn't have the role yet
	hasRole, err := tg.HasRole(&bind.CallOpts{Context: t.Context()}, minterRole, newAccount)
	require.NoError(t, err)
	require.False(t, hasRole, "Account should not have MINTER_ROLE before granting")

	// Create the changeset input
	grantInput := sequences.TokenGovernorRoleInput{
		Tokens: map[uint64]map[string]sequences.TokenGovernorGrantRole{
			chainSelector: {
				tokenSymbol: {
					Role:    "minter",
					Account: newAccount,
				},
			},
		},
	}

	// Apply the changeset
	grantCS := changesets.GrantRoleChangeset()
	err = grantCS.VerifyPreconditions(*e, grantInput)
	require.NoError(t, err, "VerifyPreconditions should pass")

	_, err = grantCS.Apply(*e, grantInput)
	require.NoError(t, err, "Apply should succeed")

	// Verify the role was granted
	hasRole, err = tg.HasRole(&bind.CallOpts{Context: t.Context()}, minterRole, newAccount)
	require.NoError(t, err)
	require.True(t, hasRole, "Account should have MINTER_ROLE after granting")
	t.Logf("Successfully granted MINTER_ROLE to %s", newAccount.Hex())
}

// TestRevokeRoleChangeset tests the RevokeRoleChangeset
func TestRevokeRoleChangeset(t *testing.T) {
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
	tokenSymbol := "TREVOKE"
	tokenAddr := deployTestToken(t, chain, tokenSymbol, 18)

	// Deploy token governor
	deployInput := sequences.DeployTokenGovernorInput{
		Token:               tokenAddr.Hex(),
		InitialDelay:        big.NewInt(0),
		InitialDefaultAdmin: chain.DeployerKey.From.Hex(),
		ChainSelector:       chainSelector,
	}
	deployCS := changesets.DeployTokenGovernorChangeset()
	deployOutput, err := deployCS.Apply(*e, deployInput)
	require.NoError(t, err, "Failed to deploy TokenGovernor")
	e.DataStore = deployOutput.DataStore.Seal()

	// Get the token governor
	tgAddr, err := sequences.GetTokenGovernor(e.DataStore, chainSelector, tokenSymbol)
	require.NoError(t, err)
	tg, err := tg_bindings.NewTokenGovernor(tgAddr, chain.Client)
	require.NoError(t, err)

	// Create a new account and grant MINTER_ROLE first
	newAccount := common.HexToAddress("0x2222222222222222222222222222222222222222")

	grantInput := sequences.TokenGovernorRoleInput{
		Tokens: map[uint64]map[string]sequences.TokenGovernorGrantRole{
			chainSelector: {
				tokenSymbol: {
					Role:    "minter",
					Account: newAccount,
				},
			},
		},
	}
	grantCS := changesets.GrantRoleChangeset()
	_, err = grantCS.Apply(*e, grantInput)
	require.NoError(t, err, "Failed to grant role")

	// Verify the account has the role
	minterRole, err := tg.MINTERROLE(&bind.CallOpts{Context: t.Context()})
	require.NoError(t, err)
	hasRole, err := tg.HasRole(&bind.CallOpts{Context: t.Context()}, minterRole, newAccount)
	require.NoError(t, err)
	require.True(t, hasRole, "Account should have MINTER_ROLE after granting")

	// Now revoke the role
	revokeInput := sequences.TokenGovernorRoleInput{
		Tokens: map[uint64]map[string]sequences.TokenGovernorGrantRole{
			chainSelector: {
				tokenSymbol: {
					Role:    "minter",
					Account: newAccount,
				},
			},
		},
	}
	revokeCS := changesets.RevokeRoleChangeset()
	_, err = revokeCS.Apply(*e, revokeInput)
	require.NoError(t, err, "Apply should succeed")

	// Verify the role was revoked
	hasRole, err = tg.HasRole(&bind.CallOpts{Context: t.Context()}, minterRole, newAccount)
	require.NoError(t, err)
	require.False(t, hasRole, "Account should not have MINTER_ROLE after revoking")
	t.Logf("Successfully revoked MINTER_ROLE from %s", newAccount.Hex())
}

// TestRenounceRoleChangeset tests the RenounceRoleChangeset
func TestRenounceRoleChangeset(t *testing.T) {
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
	tokenSymbol := "TRENOUNCE"
	tokenAddr := deployTestToken(t, chain, tokenSymbol, 18)

	// Deploy token governor
	deployInput := sequences.DeployTokenGovernorInput{
		Token:               tokenAddr.Hex(),
		InitialDelay:        big.NewInt(0),
		InitialDefaultAdmin: chain.DeployerKey.From.Hex(),
		ChainSelector:       chainSelector,
	}
	deployCS := changesets.DeployTokenGovernorChangeset()
	deployOutput, err := deployCS.Apply(*e, deployInput)
	require.NoError(t, err, "Failed to deploy TokenGovernor")
	e.DataStore = deployOutput.DataStore.Seal()

	// Get the token governor
	tgAddr, err := sequences.GetTokenGovernor(e.DataStore, chainSelector, tokenSymbol)
	require.NoError(t, err)
	tg, err := tg_bindings.NewTokenGovernor(tgAddr, chain.Client)
	require.NoError(t, err)

	// Grant MINTER_ROLE to the deployer (so we can renounce it)
	grantInput := sequences.TokenGovernorRoleInput{
		Tokens: map[uint64]map[string]sequences.TokenGovernorGrantRole{
			chainSelector: {
				tokenSymbol: {
					Role:    "minter",
					Account: chain.DeployerKey.From,
				},
			},
		},
	}
	grantCS := changesets.GrantRoleChangeset()
	_, err = grantCS.Apply(*e, grantInput)
	require.NoError(t, err, "Failed to grant role")

	// Verify the deployer has the role
	minterRole, err := tg.MINTERROLE(&bind.CallOpts{Context: t.Context()})
	require.NoError(t, err)
	hasRole, err := tg.HasRole(&bind.CallOpts{Context: t.Context()}, minterRole, chain.DeployerKey.From)
	require.NoError(t, err)
	require.True(t, hasRole, "Deployer should have MINTER_ROLE after granting")

	// Renounce the role
	renounceInput := sequences.TokenGovernorRoleInput{
		Tokens: map[uint64]map[string]sequences.TokenGovernorGrantRole{
			chainSelector: {
				tokenSymbol: {
					Role:    "minter",
					Account: chain.DeployerKey.From,
				},
			},
		},
	}
	renounceCS := changesets.RenounceRoleChangeset()
	_, err = renounceCS.Apply(*e, renounceInput)
	require.NoError(t, err, "Apply should succeed")

	// Verify the role was renounced
	hasRole, err = tg.HasRole(&bind.CallOpts{Context: t.Context()}, minterRole, chain.DeployerKey.From)
	require.NoError(t, err)
	require.False(t, hasRole, "Deployer should not have MINTER_ROLE after renouncing")
	t.Logf("Successfully renounced MINTER_ROLE by %s", chain.DeployerKey.From.Hex())
}

// TestTransferOwnershipChangeset tests the TransferOwnershipChangeset validation
func TestTransferOwnershipChangeset(t *testing.T) {
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
	tokenSymbol := "TOWNER"
	tokenAddr := deployTestToken(t, chain, tokenSymbol, 18)

	// Deploy token governor
	deployInput := sequences.DeployTokenGovernorInput{
		Token:               tokenAddr.Hex(),
		InitialDelay:        big.NewInt(0),
		InitialDefaultAdmin: chain.DeployerKey.From.Hex(),
		ChainSelector:       chainSelector,
	}
	deployCS := changesets.DeployTokenGovernorChangeset()
	deployOutput, err := deployCS.Apply(*e, deployInput)
	require.NoError(t, err, "Failed to deploy TokenGovernor")
	e.DataStore = deployOutput.DataStore.Seal()

	// Get the token governor
	tgAddr, err := sequences.GetTokenGovernor(e.DataStore, chainSelector, tokenSymbol)
	require.NoError(t, err)
	tg, err := tg_bindings.NewTokenGovernor(tgAddr, chain.Client)
	require.NoError(t, err)

	// Verify initial owner
	owner, err := tg.Owner(&bind.CallOpts{Context: t.Context()})
	require.NoError(t, err)
	require.Equal(t, chain.DeployerKey.From, owner, "Initial owner should be deployer")

	// Test validation: transferring to same owner should fail
	transferInput := sequences.TokenGovernorOwnershipInput{
		Tokens: map[uint64]map[string]common.Address{
			chainSelector: {
				tokenSymbol: chain.DeployerKey.From, // Same as current owner
			},
		},
	}

	cs := changesets.TransferOwnershipChangeset()
	err = cs.VerifyPreconditions(*e, transferInput)
	require.NoError(t, err, "Verify preconditions should pass")

	_, err = cs.Apply(*e, transferInput)
	require.Error(t, err, "Should fail when transferring to same owner")
	require.Contains(t, err.Error(), "already the current owner")
	t.Logf("Correctly rejected transfer to same owner")
}

// TestBeginDefaultAdminTransferChangeset tests the BeginDefaultAdminTransferChangeset validation
func TestBeginDefaultAdminTransferChangeset(t *testing.T) {
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
	tokenSymbol := "TADMIN"
	tokenAddr := deployTestToken(t, chain, tokenSymbol, 18)

	// Deploy token governor
	deployInput := sequences.DeployTokenGovernorInput{
		Token:               tokenAddr.Hex(),
		InitialDelay:        big.NewInt(0),
		InitialDefaultAdmin: chain.DeployerKey.From.Hex(),
		ChainSelector:       chainSelector,
	}
	deployCS := changesets.DeployTokenGovernorChangeset()
	deployOutput, err := deployCS.Apply(*e, deployInput)
	require.NoError(t, err, "Failed to deploy TokenGovernor")
	e.DataStore = deployOutput.DataStore.Seal()

	// Get the token governor
	tgAddr, err := sequences.GetTokenGovernor(e.DataStore, chainSelector, tokenSymbol)
	require.NoError(t, err)
	tg, err := tg_bindings.NewTokenGovernor(tgAddr, chain.Client)
	require.NoError(t, err)

	// Verify current default admin
	currentAdmin, err := tg.DefaultAdmin(&bind.CallOpts{Context: t.Context()})
	require.NoError(t, err)
	require.Equal(t, chain.DeployerKey.From, currentAdmin, "Current admin should be deployer")

	// Test validation: transferring to same admin should fail
	adminInput := sequences.TokenGovernorOwnershipInput{
		Tokens: map[uint64]map[string]common.Address{
			chainSelector: {
				tokenSymbol: chain.DeployerKey.From, // Same as current admin
			},
		},
	}

	cs := changesets.BeginDefaultAdminTransferChangeset()
	err = cs.VerifyPreconditions(*e, adminInput)
	require.NoError(t, err, "Verify preconditions should pass")

	_, err = cs.Apply(*e, adminInput)
	require.Error(t, err, "Should fail when transferring to same admin")
	require.Contains(t, err.Error(), "already the current default admin")
	t.Logf("Correctly rejected admin transfer to same admin")
}

// TestRoleChangesetValidation tests the validation for role changesets
func TestRoleChangesetValidation(t *testing.T) {
	t.Parallel()

	evmChains := []uint64{
		chain_selectors.ETHEREUM_MAINNET.Selector,
	}

	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, evmChains),
	)
	require.NoError(t, err, "Failed to create test environment")

	grantCS := changesets.GrantRoleChangeset()

	t.Run("empty tokens map", func(t *testing.T) {
		input := sequences.TokenGovernorRoleInput{
			Tokens: map[uint64]map[string]sequences.TokenGovernorGrantRole{},
		}
		err := grantCS.VerifyPreconditions(*e, input)
		require.Error(t, err)
		require.Contains(t, err.Error(), "tokens map must be provided")
	})

	t.Run("nonexistent chain", func(t *testing.T) {
		input := sequences.TokenGovernorRoleInput{
			Tokens: map[uint64]map[string]sequences.TokenGovernorGrantRole{
				999999: {
					"TEST": {
						Role:    "minter",
						Account: common.HexToAddress("0x1111111111111111111111111111111111111111"),
					},
				},
			},
		}
		err := grantCS.VerifyPreconditions(*e, input)
		require.Error(t, err)
		require.Contains(t, err.Error(), "chain with selector 999999 does not exist")
	})

	t.Run("empty token symbol", func(t *testing.T) {
		input := sequences.TokenGovernorRoleInput{
			Tokens: map[uint64]map[string]sequences.TokenGovernorGrantRole{
				chain_selectors.ETHEREUM_MAINNET.Selector: {
					"": {
						Role:    "minter",
						Account: common.HexToAddress("0x1111111111111111111111111111111111111111"),
					},
				},
			},
		}
		err := grantCS.VerifyPreconditions(*e, input)
		require.Error(t, err)
		require.Contains(t, err.Error(), "token symbol must be provided")
	})

	t.Run("zero account address", func(t *testing.T) {
		input := sequences.TokenGovernorRoleInput{
			Tokens: map[uint64]map[string]sequences.TokenGovernorGrantRole{
				chain_selectors.ETHEREUM_MAINNET.Selector: {
					"TEST": {
						Role:    "minter",
						Account: common.Address{},
					},
				},
			},
		}
		err := grantCS.VerifyPreconditions(*e, input)
		require.Error(t, err)
		require.Contains(t, err.Error(), "account address must be provided")
	})

	t.Run("invalid role string", func(t *testing.T) {
		input := sequences.TokenGovernorRoleInput{
			Tokens: map[uint64]map[string]sequences.TokenGovernorGrantRole{
				chain_selectors.ETHEREUM_MAINNET.Selector: {
					"TEST": {
						Role:    "invalid_role",
						Account: common.HexToAddress("0x1111111111111111111111111111111111111111"),
					},
				},
			},
		}
		err := grantCS.VerifyPreconditions(*e, input)
		require.Error(t, err)
		require.Contains(t, err.Error(), "invalid role")
	})

	t.Run("empty role string", func(t *testing.T) {
		input := sequences.TokenGovernorRoleInput{
			Tokens: map[uint64]map[string]sequences.TokenGovernorGrantRole{
				chain_selectors.ETHEREUM_MAINNET.Selector: {
					"TEST": {
						Role:    "",
						Account: common.HexToAddress("0x1111111111111111111111111111111111111111"),
					},
				},
			},
		}
		err := grantCS.VerifyPreconditions(*e, input)
		require.Error(t, err)
		require.Contains(t, err.Error(), "role must be provided")
	})
}

// TestOwnershipChangesetValidation tests the validation for ownership changesets
func TestOwnershipChangesetValidation(t *testing.T) {
	t.Parallel()

	evmChains := []uint64{
		chain_selectors.ETHEREUM_MAINNET.Selector,
	}

	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, evmChains),
	)
	require.NoError(t, err, "Failed to create test environment")

	cs := changesets.TransferOwnershipChangeset()

	t.Run("empty tokens map", func(t *testing.T) {
		input := sequences.TokenGovernorOwnershipInput{
			Tokens: map[uint64]map[string]common.Address{},
		}
		err := cs.VerifyPreconditions(*e, input)
		require.Error(t, err)
		require.Contains(t, err.Error(), "tokens map must be provided")
	})

	t.Run("nonexistent chain", func(t *testing.T) {
		input := sequences.TokenGovernorOwnershipInput{
			Tokens: map[uint64]map[string]common.Address{
				999999: {
					"TEST": common.HexToAddress("0x1111111111111111111111111111111111111111"),
				},
			},
		}
		err := cs.VerifyPreconditions(*e, input)
		require.Error(t, err)
		require.Contains(t, err.Error(), "chain with selector 999999 does not exist")
	})

	t.Run("empty token symbol", func(t *testing.T) {
		input := sequences.TokenGovernorOwnershipInput{
			Tokens: map[uint64]map[string]common.Address{
				chain_selectors.ETHEREUM_MAINNET.Selector: {
					"": common.HexToAddress("0x1111111111111111111111111111111111111111"),
				},
			},
		}
		err := cs.VerifyPreconditions(*e, input)
		require.Error(t, err)
		require.Contains(t, err.Error(), "token symbol must be provided")
	})

	t.Run("zero new owner address", func(t *testing.T) {
		input := sequences.TokenGovernorOwnershipInput{
			Tokens: map[uint64]map[string]common.Address{
				chain_selectors.ETHEREUM_MAINNET.Selector: {
					"TEST": common.Address{},
				},
			},
		}
		err := cs.VerifyPreconditions(*e, input)
		require.Error(t, err)
		require.Contains(t, err.Error(), "new owner address must be provided")
	})
}

// TestDefaultAdminChangesetValidation tests the validation for default admin changesets
func TestDefaultAdminChangesetValidation(t *testing.T) {
	t.Parallel()

	evmChains := []uint64{
		chain_selectors.ETHEREUM_MAINNET.Selector,
	}

	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, evmChains),
	)
	require.NoError(t, err, "Failed to create test environment")

	cs := changesets.BeginDefaultAdminTransferChangeset()

	t.Run("empty tokens map", func(t *testing.T) {
		input := sequences.TokenGovernorOwnershipInput{
			Tokens: map[uint64]map[string]common.Address{},
		}
		err := cs.VerifyPreconditions(*e, input)
		require.Error(t, err)
		require.Contains(t, err.Error(), "tokens map must be provided")
	})

	t.Run("nonexistent chain", func(t *testing.T) {
		input := sequences.TokenGovernorOwnershipInput{
			Tokens: map[uint64]map[string]common.Address{
				999999: {
					"TEST": common.HexToAddress("0x1111111111111111111111111111111111111111"),
				},
			},
		}
		err := cs.VerifyPreconditions(*e, input)
		require.Error(t, err)
		require.Contains(t, err.Error(), "chain with selector 999999 does not exist")
	})

	t.Run("empty token symbol", func(t *testing.T) {
		input := sequences.TokenGovernorOwnershipInput{
			Tokens: map[uint64]map[string]common.Address{
				chain_selectors.ETHEREUM_MAINNET.Selector: {
					"": common.HexToAddress("0x1111111111111111111111111111111111111111"),
				},
			},
		}
		err := cs.VerifyPreconditions(*e, input)
		require.Error(t, err)
		require.Contains(t, err.Error(), "token symbol must be provided")
	})

	t.Run("zero new admin owner address", func(t *testing.T) {
		input := sequences.TokenGovernorOwnershipInput{
			Tokens: map[uint64]map[string]common.Address{
				chain_selectors.ETHEREUM_MAINNET.Selector: {
					"TEST": common.Address{},
				},
			},
		}
		err := cs.VerifyPreconditions(*e, input)
		require.Error(t, err)
		require.Contains(t, err.Error(), "new admin owner address must be provided")
	})
}

// TestAcceptOwnershipChangesetValidation tests the AcceptOwnershipChangeset validation
func TestAcceptOwnershipChangesetValidation(t *testing.T) {
	t.Parallel()

	evmChains := []uint64{
		chain_selectors.ETHEREUM_MAINNET.Selector,
	}

	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, evmChains),
	)
	require.NoError(t, err, "Failed to create test environment")

	cs := changesets.AcceptOwnershipChangeset()

	t.Run("empty tokens map", func(t *testing.T) {
		input := sequences.TokenGovernorOwnershipInput{
			Tokens: map[uint64]map[string]common.Address{},
		}
		err := cs.VerifyPreconditions(*e, input)
		require.Error(t, err)
		require.Contains(t, err.Error(), "tokens map must be provided")
	})

	t.Run("nonexistent chain", func(t *testing.T) {
		input := sequences.TokenGovernorOwnershipInput{
			Tokens: map[uint64]map[string]common.Address{
				999999: {
					"TEST": common.HexToAddress("0x1111111111111111111111111111111111111111"),
				},
			},
		}
		err := cs.VerifyPreconditions(*e, input)
		require.Error(t, err)
		require.Contains(t, err.Error(), "chain with selector 999999 does not exist")
	})

	t.Run("zero new owner address", func(t *testing.T) {
		input := sequences.TokenGovernorOwnershipInput{
			Tokens: map[uint64]map[string]common.Address{
				chain_selectors.ETHEREUM_MAINNET.Selector: {
					"TEST": common.Address{},
				},
			},
		}
		err := cs.VerifyPreconditions(*e, input)
		require.Error(t, err)
		require.Contains(t, err.Error(), "new owner address must be provided")
	})
}

// TestAcceptDefaultAdminTransferChangesetValidation tests the AcceptDefaultAdminTransferChangeset validation
func TestAcceptDefaultAdminTransferChangesetValidation(t *testing.T) {
	t.Parallel()

	evmChains := []uint64{
		chain_selectors.ETHEREUM_MAINNET.Selector,
	}

	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, evmChains),
	)
	require.NoError(t, err, "Failed to create test environment")

	cs := changesets.AcceptDefaultAdminTransferChangeset()

	t.Run("empty tokens map", func(t *testing.T) {
		input := sequences.TokenGovernorOwnershipInput{
			Tokens: map[uint64]map[string]common.Address{},
		}
		err := cs.VerifyPreconditions(*e, input)
		require.Error(t, err)
		require.Contains(t, err.Error(), "tokens map must be provided")
	})

	t.Run("nonexistent chain", func(t *testing.T) {
		input := sequences.TokenGovernorOwnershipInput{
			Tokens: map[uint64]map[string]common.Address{
				999999: {
					"TEST": common.HexToAddress("0x1111111111111111111111111111111111111111"),
				},
			},
		}
		err := cs.VerifyPreconditions(*e, input)
		require.Error(t, err)
		require.Contains(t, err.Error(), "chain with selector 999999 does not exist")
	})

	t.Run("zero admin address", func(t *testing.T) {
		input := sequences.TokenGovernorOwnershipInput{
			Tokens: map[uint64]map[string]common.Address{
				chain_selectors.ETHEREUM_MAINNET.Selector: {
					"TEST": common.Address{},
				},
			},
		}
		err := cs.VerifyPreconditions(*e, input)
		require.Error(t, err)
		require.Contains(t, err.Error(), "new admin owner address must be provided")
	})
}
