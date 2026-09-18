package sequences

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	chain_selectors "github.com/smartcontractkit/chain-selectors"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/tokens/tokenimpl"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/burn_mint_erc20"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/burn_mint_erc20_transparent"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/erc20"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/tip20"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/burn_mint_erc20_with_drip"
	bnm_bindings "github.com/smartcontractkit/chainlink-evm/gethwrappers/shared/generated/initial/burn_mint_erc20"
	bnm_transparent_bindings "github.com/smartcontractkit/chainlink-evm/gethwrappers/shared/generated/latest/burn_mint_erc20_transparent"

	tokensapi "github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils"
)

// TestEVMTokenDeployments tests various EVM token deployments using the DeployToken sequence directly.
// This covers all supported EVM token types: ERC20, BurnMintERC20, FactoryBurnMintERC20, and BurnMintERC20WithDrip.
// Note: The full TokenExpansion changeset is not yet implemented for EVM (DeployTokenPoolForToken,
// RegisterToken, SetPool return nil), so we test token deployment directly via the sequence.
func TestEVMTokenDeployments(t *testing.T) {
	t.Parallel()

	evmChains := []uint64{
		chain_selectors.ETHEREUM_MAINNET.Selector,
	}

	maxSupply := uint64(1_000_000_000) // 1 billion tokens
	preMint := uint64(1_000_000)       // 1 million tokens

	testCases := []struct {
		name           string
		tokenType      cldf.ContractType
		tokenName      string
		tokenSymbol    string
		decimals       uint8
		supply         *uint64
		preMint        *uint64
		ccipAdmin      string   // Address to set as CCIP admin
		externalAdmins []string // Addresses to grant admin role
		sender         string   // Address that will receive the pre-mint tokens (if applicable)
		requiresOwner  bool
		requiresSupply bool
	}{
		{
			name:        "ERC20Token",
			tokenType:   erc20.ContractType,
			tokenName:   "Test ERC20",
			tokenSymbol: "TERC20",
			decimals:    18,
		},
		{
			name:           "BurnMintERC20Token",
			tokenType:      burn_mint_erc20.ContractType,
			tokenName:      "Test BurnMint ERC20",
			tokenSymbol:    "TBMERC20",
			decimals:       18,
			ccipAdmin:      "0x1111111111111111111111111111111111111111",
			sender:         "0x1111111111111111111111111111111111111111",
			supply:         &maxSupply,
			preMint:        &preMint,
			requiresSupply: true,
		},
		{
			name:        "BurnMintERC20WithDrip",
			tokenType:   burn_mint_erc20_with_drip.ContractType,
			tokenName:   "Test BurnMint ERC20 With Drip",
			tokenSymbol: "TBMDRIP",
			decimals:    18,
			ccipAdmin:   "0x1111111111111111111111111111111111111111",
		},
		{
			name:        "BurnMintERC20Token_derivesCCIPAdminFromExternal",
			tokenType:   burn_mint_erc20.ContractType,
			tokenName:   "Test BurnMint ERC20 Derived CCIP Admin",
			tokenSymbol: "TBMDERIVE",
			decimals:    18,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			e, err := environment.New(t.Context(),
				environment.WithEVMSimulated(t, evmChains),
			)
			require.NoError(t, err, "Failed to create test environment")
			require.NotNil(t, e, "Environment should be created")

			// Get deployer address for external admin
			chain := e.BlockChains.EVMChains()[chain_selectors.ETHEREUM_MAINNET.Selector]
			deployerAddr := chain.DeployerKey.From

			externalAdmin := "0x2222222222222222222222222222222222222222"
			expectedCCIPAdmin := common.HexToAddress(externalAdmin)
			if tc.ccipAdmin != "" {
				expectedCCIPAdmin = common.HexToAddress(tc.ccipAdmin)
			}

			// Build token input based on test case configuration
			tokenInput := tokensapi.DeployTokenInput{
				Name:          tc.tokenName,
				Symbol:        tc.tokenSymbol,
				Decimals:      tc.decimals,
				Type:          tc.tokenType,
				ExternalAdmin: externalAdmin,
				CCIPAdmin:     tc.ccipAdmin,
				Senders:       []string{},
				ChainSelector: chain_selectors.ETHEREUM_MAINNET.Selector,
			}

			// Add supply and pre-mint for tokens that require it
			if tc.requiresSupply {
				tokenInput.Supply = tc.supply
				if tc.preMint != nil {
					tokenInput.PreMint = tc.preMint
				}
			}

			// Set the expected pre-mint receiver address (defaults to deployer if not specified)
			expectedPreMintReceiver := deployerAddr
			if tc.sender != "" {
				tokenInput.Senders = append(tokenInput.Senders, tc.sender)
				expectedPreMintReceiver = common.HexToAddress(tc.sender)
			}

			tokenInput.ExistingDataStore = e.DataStore
			deployTokenSeq := DeployToken
			require.NotNil(t, deployTokenSeq, "DeployToken sequence should not be nil")

			report, err := cldf_ops.ExecuteSequence(e.OperationsBundle, deployTokenSeq, e.BlockChains, tokenInput)
			require.NoError(t, err, "Failed to execute DeployToken sequence for %s", tc.name)
			require.NotNil(t, report, "DeployToken report should not be nil for %s", tc.name)

			// Verify the token was deployed by checking output addresses
			require.GreaterOrEqual(t, len(report.Output.Addresses), 1, "Token %s should have at least one address deployed", tc.name)

			// Verify the token address has the correct type and make on-chain calls
			tokenFound := false
			for _, addr := range report.Output.Addresses {
				if addr.Type == datastore.ContractType(tc.tokenType) &&
					addr.ChainSelector == chain_selectors.ETHEREUM_MAINNET.Selector &&
					addr.Qualifier == tc.tokenSymbol {
					tokenFound = true
					require.NotEmpty(t, addr.Address, "Token address should not be empty")
					t.Logf("Deployed %s token at address: %s", tc.name, addr.Address)

					// Make on-chain calls to verify token properties
					tokenAddr := common.HexToAddress(addr.Address)
					tokenContract, err := bnm_bindings.NewBurnMintERC20(tokenAddr, chain.Client)
					require.NoError(t, err, "Failed to create token contract binding")

					// Verify name
					onChainName, err := tokenContract.Name(&bind.CallOpts{})
					require.NoError(t, err, "Failed to get token name from chain")
					require.Equal(t, tc.tokenName, onChainName, "Token name mismatch for %s", tc.name)
					t.Logf("  On-chain name: %s", onChainName)

					// Verify symbol
					onChainSymbol, err := tokenContract.Symbol(&bind.CallOpts{})
					require.NoError(t, err, "Failed to get token symbol from chain")
					require.Equal(t, tc.tokenSymbol, onChainSymbol, "Token symbol mismatch for %s", tc.name)
					t.Logf("  On-chain symbol: %s", onChainSymbol)

					// Verify decimals (only for tokens that support custom decimals)
					if tc.requiresSupply {
						onChainDecimals, err := tokenContract.Decimals(&bind.CallOpts{})
						require.NoError(t, err, "Failed to get token decimals from chain")
						require.Equal(t, tc.decimals, onChainDecimals, "Token decimals mismatch for %s", tc.name)
						t.Logf("  On-chain decimals: %d", onChainDecimals)

						// Verify max supply
						expectedMaxSupply := tokensapi.ScaleTokenAmount(new(big.Int).SetUint64(*tc.supply), tc.decimals)
						onChainMaxSupply, err := tokenContract.MaxSupply(&bind.CallOpts{})
						require.NoError(t, err, "Failed to get token max supply from chain")
						require.Equal(t, expectedMaxSupply.String(), onChainMaxSupply.String(), "Token max supply mismatch for %s", tc.name)
						t.Logf("  On-chain maxSupply: %s", onChainMaxSupply.String())

						// Verify total supply (should match preMint if set)
						onChainTotalSupply, err := tokenContract.TotalSupply(&bind.CallOpts{})
						require.NoError(t, err, "Failed to get token total supply from chain")
						if tc.preMint != nil {
							expectedPreMint := tokensapi.ScaleTokenAmount(new(big.Int).SetUint64(*tc.preMint), tc.decimals)
							require.Equal(t, expectedPreMint.String(), onChainTotalSupply.String(), "Token total supply mismatch for %s", tc.name)
							balance, err := tokenContract.BalanceOf(&bind.CallOpts{}, expectedPreMintReceiver)
							require.NoError(t, err, "Failed to get balance of pre-mint receiver from chain")
							require.Equal(t, expectedPreMint.String(), balance.String(), "Pre-mint receiver balance mismatch for address %s and token %s", expectedPreMintReceiver, tc.name)
							t.Logf("  On-chain totalSupply: %s (matches preMint)", onChainTotalSupply.String())
						} else {
							t.Logf("  On-chain totalSupply: %s", onChainTotalSupply.String())
						}
					}

					caps := tokenimpl.Capabilities(tc.tokenType)
					if caps.SupportsCCIPAdmin {
						// Verify CCIP Admin was set correctly
						t.Log("  Verifying CCIP Admin...")
						onChainCCIPAdmin, err := tokenContract.GetCCIPAdmin(&bind.CallOpts{})
						require.NoError(t, err, "Failed to get CCIP admin from chain")
						require.Equal(t, expectedCCIPAdmin, onChainCCIPAdmin, "CCIP admin mismatch")
						t.Logf("  On-chain CCIP admin: %s (expected: %s)", onChainCCIPAdmin.Hex(), expectedCCIPAdmin.Hex())

						// TEST 2: Verify External Admins have the DEFAULT_ADMIN_ROLE
						t.Log("  Verifying External Admin roles...")
						defaultAdminRole, err := tokenContract.DEFAULTADMINROLE(&bind.CallOpts{})
						require.NoError(t, err, "Failed to get DEFAULT_ADMIN_ROLE")
						t.Logf("  DEFAULT_ADMIN_ROLE: 0x%x", defaultAdminRole)

						// Verify externalAdmin has the admin role
						hasRole, err := tokenContract.HasRole(&bind.CallOpts{}, defaultAdminRole, common.HexToAddress(externalAdmin))
						require.NoError(t, err, "Failed to check HasRole for externalAdmin")
						require.True(t, hasRole, "External admin should have DEFAULT_ADMIN_ROLE")
						t.Logf("  External admin (%s) has DEFAULT_ADMIN_ROLE: %v", externalAdmin, hasRole)

						// Verify deployer still has the admin role (original deployer should retain role)
						deployerHasRole, err := tokenContract.HasRole(&bind.CallOpts{}, defaultAdminRole, deployerAddr)
						require.NoError(t, err, "Failed to check HasRole for deployer")
						require.True(t, deployerHasRole, "Deployer should still have DEFAULT_ADMIN_ROLE")
						t.Logf("  Deployer (%s) has DEFAULT_ADMIN_ROLE: %v", deployerAddr.Hex(), deployerHasRole)
					}

					break
				}
			}
			require.True(t, tokenFound, "Token %s should be found in deployed addresses", tc.name)
		})
	}
}

func TestTokenSupportsAdminRole(t *testing.T) {
	t.Parallel()

	tokenTypes := map[cldf.ContractType]bool{
		burn_mint_erc20_with_drip.ContractType:   true,
		burn_mint_erc20.ContractType:             true,
		utils.ERC677TokenHelper:                  false,
		utils.BurnMintToken:                      false,
		tip20.ContractType:                       true,
		erc20.ContractType:                       false,
		burn_mint_erc20_transparent.ContractType: true,
	}

	for tt, supportsAdmin := range tokenTypes {
		require.Equal(t, supportsAdmin, tokenimpl.Capabilities(tt).SupportsAdminRole, "Token type %s admin role support mismatch", tt)
	}
}

func TestTokenUsesAsyncRoleManagement(t *testing.T) {
	t.Parallel()

	tokenTypes := map[cldf.ContractType]bool{
		burn_mint_erc20_with_drip.ContractType:   false,
		burn_mint_erc20.ContractType:             false,
		utils.ERC677TokenHelper:                  false,
		utils.BurnMintToken:                      false,
		tip20.ContractType:                       false,
		erc20.ContractType:                       false,
		burn_mint_erc20_transparent.ContractType: true,
	}

	for tt, usesAsync := range tokenTypes {
		require.Equal(t, usesAsync, tokenimpl.Capabilities(tt).UsesAsyncRoleManagement, "Token type %s async role management mismatch", tt)
	}
}

// TestEVMTokenDeployment_BurnMintERC20TransparentToken exercises the composite deploy (impl +
// TransparentUpgradeableProxy + initialize) via the DeployToken sequence, using its own
// assertions rather than the shared BurnMintERC20 test loop above: this token type has a
// different capability set (no SupportsAdminRole yet - see the adapter's Capabilities doc) and
// a different admin/preMint/ccipAdmin resolution path (all three default to the deployer via
// initialize's defaultAdmin argument, then move to their intended holders exactly like
// BurnMintERC20's constructor-implicit msg.sender pattern).
func TestEVMTokenDeployment_BurnMintERC20TransparentToken(t *testing.T) {
	t.Parallel()

	evmChains := []uint64{chain_selectors.ETHEREUM_MAINNET.Selector}

	e, err := environment.New(t.Context(), environment.WithEVMSimulated(t, evmChains))
	require.NoError(t, err, "Failed to create test environment")

	chain := e.BlockChains.EVMChains()[chain_selectors.ETHEREUM_MAINNET.Selector]
	deployerAddr := chain.DeployerKey.From

	externalAdmin := "0x3333333333333333333333333333333333333333" // proxy admin / upgrade authority
	sender := "0x4444444444444444444444444444444444444444"        // pre-mint recipient
	maxSupply := uint64(1_000_000_000)
	preMint := uint64(1_000_000)

	tokenInput := tokensapi.DeployTokenInput{
		Name:              "Test BurnMint ERC20 Transparent",
		Symbol:            "TBMTRANS",
		Decimals:          18,
		Type:              burn_mint_erc20_transparent.ContractType,
		ExternalAdmin:     externalAdmin,
		Senders:           []string{sender},
		Supply:            &maxSupply,
		PreMint:           &preMint,
		ChainSelector:     chain_selectors.ETHEREUM_MAINNET.Selector,
		ExistingDataStore: e.DataStore,
	}

	report, err := cldf_ops.ExecuteSequence(e.OperationsBundle, DeployToken, e.BlockChains, tokenInput)
	require.NoError(t, err, "Failed to execute DeployToken sequence")
	require.Len(t, report.Output.Addresses, 1, "DeployToken should return exactly one address ref (the proxy)")

	proxyRef := report.Output.Addresses[0]
	require.Equal(t, datastore.ContractType(burn_mint_erc20_transparent.ContractType), proxyRef.Type)
	require.Equal(t, tokenInput.Symbol, proxyRef.Qualifier)
	require.NotEmpty(t, proxyRef.Address)

	proxyAddr := common.HexToAddress(proxyRef.Address)
	token, err := bnm_transparent_bindings.NewBurnMintERC20Transparent(proxyAddr, chain.Client)
	require.NoError(t, err, "Failed to bind to proxy address")

	onChainName, err := token.Name(&bind.CallOpts{})
	require.NoError(t, err)
	require.Equal(t, tokenInput.Name, onChainName)

	onChainSymbol, err := token.Symbol(&bind.CallOpts{})
	require.NoError(t, err)
	require.Equal(t, tokenInput.Symbol, onChainSymbol)

	onChainDecimals, err := token.Decimals(&bind.CallOpts{})
	require.NoError(t, err)
	require.Equal(t, tokenInput.Decimals, onChainDecimals)

	expectedMaxSupply := tokensapi.ScaleTokenAmount(new(big.Int).SetUint64(maxSupply), tokenInput.Decimals)
	onChainMaxSupply, err := token.MaxSupply(&bind.CallOpts{})
	require.NoError(t, err)
	require.Equal(t, expectedMaxSupply.String(), onChainMaxSupply.String())

	expectedPreMint := tokensapi.ScaleTokenAmount(new(big.Int).SetUint64(preMint), tokenInput.Decimals)
	onChainTotalSupply, err := token.TotalSupply(&bind.CallOpts{})
	require.NoError(t, err)
	require.Equal(t, expectedPreMint.String(), onChainTotalSupply.String())

	// preMint was minted to the deployer inside initialize, then moved to Senders[0] by the
	// generic sequence's Transfer step - mirroring BurnMintERC20's flow.
	senderBalance, err := token.BalanceOf(&bind.CallOpts{}, common.HexToAddress(sender))
	require.NoError(t, err)
	require.Equal(t, expectedPreMint.String(), senderBalance.String(), "pre-mint recipient should hold the pre-minted tokens")

	deployerBalance, err := token.BalanceOf(&bind.CallOpts{}, deployerAddr)
	require.NoError(t, err)
	require.Equal(t, "0", deployerBalance.String(), "deployer should have transferred away the pre-mint amount")

	// ccipAdmin defaults to ExternalAdmin (see sequences/token.go), set via the generic SetCCIPAdmin step.
	onChainCCIPAdmin, err := token.GetCCIPAdmin(&bind.CallOpts{})
	require.NoError(t, err)
	require.Equal(t, common.HexToAddress(externalAdmin), onChainCCIPAdmin)

	// DEFAULT_ADMIN_ROLE transfer to ExternalAdmin is only *begun* here (beginDefaultAdminTransfer),
	// not completed: UsesAsyncRoleManagement tokens require a separate accept call signed by the
	// new admin, and this test's ExternalAdmin is a plain address, not the timelock
	// (TimelockAddress is unset), so the sequence does not auto-queue that accept. The deployer
	// key retains the role until ExternalAdmin calls acceptDefaultAdminTransfer itself.
	defaultAdminRole, err := token.DEFAULTADMINROLE(&bind.CallOpts{})
	require.NoError(t, err)
	deployerHasRole, err := token.HasRole(&bind.CallOpts{}, defaultAdminRole, deployerAddr)
	require.NoError(t, err)
	require.True(t, deployerHasRole, "deployer should retain DEFAULT_ADMIN_ROLE until ExternalAdmin accepts")
	externalAdminHasRole, err := token.HasRole(&bind.CallOpts{}, defaultAdminRole, common.HexToAddress(externalAdmin))
	require.NoError(t, err)
	require.False(t, externalAdminHasRole, "external admin should NOT have DEFAULT_ADMIN_ROLE yet (has not accepted)")

	pending, err := token.PendingDefaultAdmin(&bind.CallOpts{})
	require.NoError(t, err)
	require.Equal(t, common.HexToAddress(externalAdmin), pending.NewAdmin, "external admin should be the pending default admin (begin step ran)")
}

// TestEVMTokenDeployment_BurnMintERC20TransparentToken_ExternalAdminIsTimelock verifies the
// deploy-time auto-accept path: when ExternalAdmin matches TimelockAddress (resolved from the
// MCMS config by TokenExpansion), the sequence queues AcceptDefaultAdminTransfer into the batch
// alongside the (synchronously executed) begin-transfer, rather than requiring an out-of-band
// accept as it would for a customer-provided ExternalAdmin.
func TestEVMTokenDeployment_BurnMintERC20TransparentToken_ExternalAdminIsTimelock(t *testing.T) {
	t.Parallel()

	evmChains := []uint64{chain_selectors.ETHEREUM_MAINNET.Selector}
	e, err := environment.New(t.Context(), environment.WithEVMSimulated(t, evmChains))
	require.NoError(t, err)

	chain := e.BlockChains.EVMChains()[chain_selectors.ETHEREUM_MAINNET.Selector]
	timelockStandIn := "0x6666666666666666666666666666666666666666"

	maxSupply := uint64(1_000_000_000)
	tokenInput := tokensapi.DeployTokenInput{
		Name:              "Timelock Admin Token",
		Symbol:            "TLADMIN",
		Decimals:          18,
		Type:              burn_mint_erc20_transparent.ContractType,
		ExternalAdmin:     timelockStandIn,
		TimelockAddress:   timelockStandIn,
		Supply:            &maxSupply,
		ChainSelector:     chain_selectors.ETHEREUM_MAINNET.Selector,
		ExistingDataStore: e.DataStore,
	}
	report, err := cldf_ops.ExecuteSequence(e.OperationsBundle, DeployToken, e.BlockChains, tokenInput)
	require.NoError(t, err)
	proxyAddr := common.HexToAddress(report.Output.Addresses[0].Address)

	token, err := bnm_transparent_bindings.NewBurnMintERC20Transparent(proxyAddr, chain.Client)
	require.NoError(t, err)

	// The begin-transfer step executed synchronously (deployer-signed).
	pending, err := token.PendingDefaultAdmin(&bind.CallOpts{})
	require.NoError(t, err)
	require.Equal(t, common.HexToAddress(timelockStandIn), pending.NewAdmin)

	// The accept step was NOT executed directly (deployer isn't the pending admin) - it must be
	// queued into the batch for the timelock to execute itself.
	defaultAdminRole, err := token.DEFAULTADMINROLE(&bind.CallOpts{})
	require.NoError(t, err)
	deployerHasRole, err := token.HasRole(&bind.CallOpts{}, defaultAdminRole, chain.DeployerKey.From)
	require.NoError(t, err)
	require.True(t, deployerHasRole, "deployer should still hold the role until the queued accept executes")

	require.Len(t, report.Output.BatchOps, 1, "the unexecuted accept call should be queued in a batch")
	require.Len(t, report.Output.BatchOps[0].Transactions, 1, "batch should contain exactly the accept call")
	require.Equal(t, proxyAddr.Hex(), report.Output.BatchOps[0].Transactions[0].To)
}
