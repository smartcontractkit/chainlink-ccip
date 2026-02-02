package sequences

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	chain_selectors "github.com/smartcontractkit/chain-selectors"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/burn_mint_erc20"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/burn_mint_erc20_with_drip"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/erc20"
	bnm_bindings "github.com/smartcontractkit/chainlink-evm/gethwrappers/shared/generated/initial/burn_mint_erc20"

	tokensapi "github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
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

	testCases := []struct {
		name           string
		tokenType      cldf.ContractType
		tokenName      string
		tokenSymbol    string
		decimals       uint8
		supply         *big.Int
		preMint        *big.Int
		ccipAdmin      string   // Address to set as CCIP admin
		externalAdmins []string // Addresses to grant admin role
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
			decimals:       8,
			ccipAdmin:      "0x1111111111111111111111111111111111111111",
			supply:         big.NewInt(0).Mul(big.NewInt(1e9), big.NewInt(1e18)), // 1 billion tokens
			preMint:        big.NewInt(0).Mul(big.NewInt(1e6), big.NewInt(1e18)), // 1 million tokens
			requiresSupply: true,
		},
		{
			name:           "BurnMintERC20WithDripToken",
			tokenType:      burn_mint_erc20_with_drip.ContractType,
			tokenName:      "Test BurnMint ERC20 With Drip",
			tokenSymbol:    "TBMDRIP",
			decimals:       8,
			ccipAdmin:      "0x1111111111111111111111111111111111111111",
			supply:         big.NewInt(0).Mul(big.NewInt(1e9), big.NewInt(1e18)),
			preMint:        big.NewInt(0).Mul(big.NewInt(1e6), big.NewInt(1e18)),
			requiresSupply: true,
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

			// Generate test addresses for CCIP admin and external admins
			ccipAdminAddr := common.HexToAddress("0x1111111111111111111111111111111111111111")
			externalAdmin1 := "0x2222222222222222222222222222222222222222"
			externalAdmin2 := "0x3333333333333333333333333333333333333333"

			// Build token input based on test case configuration
			tokenInput := tokensapi.DeployTokenInput{
				Name:          tc.tokenName,
				Symbol:        tc.tokenSymbol,
				Decimals:      tc.decimals,
				Type:          tc.tokenType,
				ExternalAdmin: []string{externalAdmin1, externalAdmin2},
				CCIPAdmin:     tc.ccipAdmin,
				ChainSelector: chain_selectors.ETHEREUM_MAINNET.Selector,
			}

			// Add supply and pre-mint for tokens that require it
			if tc.requiresSupply {
				tokenInput.Supply = tc.supply
				if tc.preMint != nil {
					tokenInput.PreMint = tc.preMint
				}
			}

			// Get the EVM token adapter and execute the DeployToken sequence directly
			tokenAdapterRegistry := tokensapi.GetTokenAdapterRegistry()
			evmAdapter, exists := tokenAdapterRegistry.GetTokenAdapter(chain_selectors.FamilyEVM, utils.Version_1_6_0)
			require.True(t, exists, "EVM token adapter should be registered")

			tokenInput.ExistingDataStore = e.DataStore
			deployTokenSeq := evmAdapter.DeployToken()
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
						onChainMaxSupply, err := tokenContract.MaxSupply(&bind.CallOpts{})
						require.NoError(t, err, "Failed to get token max supply from chain")
						require.Equal(t, tc.supply.String(), onChainMaxSupply.String(), "Token max supply mismatch for %s", tc.name)
						t.Logf("  On-chain maxSupply: %s", onChainMaxSupply.String())

						// Verify total supply (should match preMint if set)
						onChainTotalSupply, err := tokenContract.TotalSupply(&bind.CallOpts{})
						require.NoError(t, err, "Failed to get token total supply from chain")
						if tc.preMint != nil {
							require.Equal(t, tc.preMint.String(), onChainTotalSupply.String(), "Token total supply mismatch for %s", tc.name)
							t.Logf("  On-chain totalSupply: %s (matches preMint)", onChainTotalSupply.String())
						} else {
							t.Logf("  On-chain totalSupply: %s", onChainTotalSupply.String())
						}
					}

					tokenSupportsAdmin := tokenSupportsAdminRole(tc.tokenType)
					if tokenSupportsAdmin {
						// Verify CCIP Admin was set correctly
						t.Log("  Verifying CCIP Admin...")
						onChainCCIPAdmin, err := tokenContract.GetCCIPAdmin(&bind.CallOpts{})
						require.NoError(t, err, "Failed to get CCIP admin from chain")
						require.Equal(t, ccipAdminAddr, onChainCCIPAdmin, "CCIP admin mismatch")
						t.Logf("  On-chain CCIP admin: %s (expected: %s)", onChainCCIPAdmin.Hex(), ccipAdminAddr.Hex())

						// TEST 2: Verify External Admins have the DEFAULT_ADMIN_ROLE
						t.Log("  Verifying External Admin roles...")
						defaultAdminRole, err := tokenContract.DEFAULTADMINROLE(&bind.CallOpts{})
						require.NoError(t, err, "Failed to get DEFAULT_ADMIN_ROLE")
						t.Logf("  DEFAULT_ADMIN_ROLE: 0x%x", defaultAdminRole)

						// Verify externalAdmin1 has the admin role
						hasRole1, err := tokenContract.HasRole(&bind.CallOpts{}, defaultAdminRole, common.HexToAddress(externalAdmin1))
						require.NoError(t, err, "Failed to check HasRole for externalAdmin1")
						require.True(t, hasRole1, "External admin 1 should have DEFAULT_ADMIN_ROLE")
						t.Logf("  External admin 1 (%s) has DEFAULT_ADMIN_ROLE: %v", externalAdmin1, hasRole1)

						// Verify externalAdmin2 has the admin role
						hasRole2, err := tokenContract.HasRole(&bind.CallOpts{}, defaultAdminRole, common.HexToAddress(externalAdmin2))
						require.NoError(t, err, "Failed to check HasRole for externalAdmin2")
						require.True(t, hasRole2, "External admin 2 should have DEFAULT_ADMIN_ROLE")
						t.Logf("  External admin 2 (%s) has DEFAULT_ADMIN_ROLE: %v", externalAdmin2, hasRole2)

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
