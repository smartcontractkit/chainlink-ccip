package deployment

import (
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	chainsel "github.com/smartcontractkit/chain-selectors"
	bnmERC20ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/burn_mint_erc20"
	deployapi "github.com/smartcontractkit/chainlink-ccip/deployment/deploy"
	"github.com/smartcontractkit/chainlink-ccip/deployment/testhelpers"
	tokensapi "github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	cciputils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/onchain"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	"github.com/stretchr/testify/require"

	_ "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/adapters"
	_ "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_1/adapters"
)

func TestRevokeTokenAdminRoleVerifyPreconditions(t *testing.T) {
	chainSelector := chainsel.TEST_90000001.Selector
	env, err := environment.New(t.Context(), environment.WithEVMSimulated(t, []uint64{chainSelector}))
	require.NoError(t, err)

	tokenRef := datastore.AddressRef{
		ChainSelector: chainSelector,
		Address:       common.HexToAddress("0x00000000000000000000000000000000000000aa").Hex(),
		Type:          datastore.ContractType(bnmERC20ops.ContractType),
		Version:       cciputils.Version_1_0_0,
		Qualifier:     "REVOKE_VERIFY",
	}
	ds := datastore.NewMemoryDataStore()
	require.NoError(t, ds.Addresses().Add(tokenRef))
	MergeAddresses(t, env, ds)

	cs := tokensapi.RevokeTokenAdminRole()
	tests := []struct {
		input tokensapi.RevokeTokenAdminRoleInput
		title string
		error string
	}{
		{
			title: "requires revocations",
			error: "at least one token admin role revocation is required",
			input: tokensapi.RevokeTokenAdminRoleInput{},
		},
		{
			title: "requires chain in environment",
			error: "not found in environment",
			input: tokensapi.RevokeTokenAdminRoleInput{
				Revocations: []tokensapi.RevokeTokenAdminRoleConfig{{
					ChainSelector: chainSelector + 1,
					TokenRef:      datastore.AddressRef{ChainSelector: chainSelector + 1},
				}},
			},
		},
		{
			title: "requires token ref",
			error: "token ref is required",
			input: tokensapi.RevokeTokenAdminRoleInput{
				Revocations: []tokensapi.RevokeTokenAdminRoleConfig{{
					ChainSelector: chainSelector,
				}},
			},
		},
		{
			title: "rejects token ref chain mismatch",
			error: "chain selector mismatch",
			input: tokensapi.RevokeTokenAdminRoleInput{
				Revocations: []tokensapi.RevokeTokenAdminRoleConfig{{
					ChainSelector: chainSelector,
					TokenRef: datastore.AddressRef{
						ChainSelector: chainSelector + 1,
						Address:       tokenRef.Address,
					},
				}},
			},
		},
		{
			title: "accepts datastore token ref by address",
			input: tokensapi.RevokeTokenAdminRoleInput{
				Revocations: []tokensapi.RevokeTokenAdminRoleConfig{{
					ChainSelector: chainSelector,
					TokenRef: datastore.AddressRef{
						Address: tokenRef.Address,
					},
				}},
			},
		},
		{
			title: "accepts address and type without datastore version",
			input: tokensapi.RevokeTokenAdminRoleInput{
				Revocations: []tokensapi.RevokeTokenAdminRoleConfig{{
					ChainSelector: chainSelector,
					TokenRef: datastore.AddressRef{
						Address: "0x00000000000000000000000000000000000000ff",
						Type:    tokenRef.Type,
					},
				}},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			err := cs.VerifyPreconditions(*env, tt.input)
			if tt.error != "" {
				require.ErrorContains(t, err, tt.error)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestRevokeTokenAdminRoleTimelock(t *testing.T) {
	// Setup environment
	selector := chainsel.TEST_90000001.Selector
	env, err := environment.New(t.Context(),
		environment.WithEVMSimulatedWithConfig(t,
			[]uint64{selector},
			onchain.EVMSimLoaderConfig{NumAdditionalAccounts: 1},
		),
	)
	require.NoError(t, err)

	// Get customer address from environment
	chain, ok := env.BlockChains.EVMChains()[selector]
	require.True(t, ok)
	require.NotEmpty(t, chain.Users)
	customer := chain.Users[0].From

	// Deploy contracts
	output, err := deployapi.DeployContracts(deployapi.GetRegistry()).Apply(*env,
		deployapi.ContractDeploymentConfig{
			MCMS: mcms.Input{},
			Chains: map[uint64]deployapi.ContractDeploymentConfigPerChain{
				selector: NewDefaultDeploymentConfigForEVM(cciputils.Version_1_6_0),
			},
		},
	)
	require.NoError(t, err)
	MergeAddresses(t, env, output.DataStore)

	// Setup MCMS
	DeployMCMS(t, env, selector, []string{cciputils.CLLQualifier})
	mcmsReader, ok := changesets.GetRegistry().GetMCMSReader(chainsel.FamilyEVM)
	require.True(t, ok)
	timelockRef, err := mcmsReader.GetTimelockRef(*env, selector, mcms.Input{Qualifier: cciputils.CLLQualifier})
	require.NoError(t, err)
	require.True(t, common.IsHexAddress(timelockRef.Address))
	timelockAddress := common.HexToAddress(timelockRef.Address)

	t.Run("revoke timelock admin while customer admin remains", func(t *testing.T) {
		// Make both the customer and timelock an admin on the token
		token := DeployBurnMintTokenEVM(t, env, selector, customer.Hex())
		_ = DeployBurnMintPoolEVM(t, env, selector, cciputils.Version_1_6_1, token.Address())
		defaultAdminRole, err := token.DEFAULTADMINROLE(&bind.CallOpts{Context: t.Context()})
		require.NoError(t, err)

		// Verify token roles (timelock and customer are both admins)
		hasRole, err := token.HasRole(&bind.CallOpts{Context: t.Context()}, defaultAdminRole, timelockAddress)
		require.NoError(t, err)
		require.True(t, hasRole)
		hasRole, err = token.HasRole(&bind.CallOpts{Context: t.Context()}, defaultAdminRole, customer)
		require.NoError(t, err)
		require.True(t, hasRole)

		// The changeset should default to revoking timelock if no particular account is specified
		revokeOutput, err := revokeTokenAdminRoleForTest(t, env, selector, token.Address().Hex(), "", customer.Hex())
		require.NoError(t, err)
		require.Len(t, revokeOutput.MCMSTimelockProposals, 1)
		testhelpers.ProcessTimelockProposals(t, *env, revokeOutput.MCMSTimelockProposals, false)

		// Verify token roles (customer is still an admin, timelock is not)
		hasRole, err = token.HasRole(&bind.CallOpts{Context: t.Context()}, defaultAdminRole, timelockAddress)
		require.NoError(t, err)
		require.False(t, hasRole)
		hasRole, err = token.HasRole(&bind.CallOpts{Context: t.Context()}, defaultAdminRole, customer)
		require.NoError(t, err)
		require.True(t, hasRole)

		// Revoking again should not generate a proposal (we already revoked timelock)
		revokeOutput, err = revokeTokenAdminRoleForTest(t, env, selector, token.Address().Hex(), "", customer.Hex())
		require.NoError(t, err)
		require.Empty(t, revokeOutput.MCMSTimelockProposals)

		// Even with no fallback, revoking again should not generate a proposal
		revokeOutput, err = revokeTokenAdminRoleForTest(t, env, selector, token.Address().Hex(), "", "")
		require.NoError(t, err)
		require.Empty(t, revokeOutput.MCMSTimelockProposals)
	})

	t.Run("bypass revoke protections", func(t *testing.T) {
		// Only make timelock an admin on the token
		token := DeployBurnMintTokenEVM(t, env, selector, timelockAddress.Hex())
		_ = DeployBurnMintPoolEVM(t, env, selector, cciputils.Version_1_6_1, token.Address())
		defaultAdminRole, err := token.DEFAULTADMINROLE(&bind.CallOpts{Context: t.Context()})
		require.NoError(t, err)

		// Verify token roles (timelock is an admin, customer is not)
		hasRole, err := token.HasRole(&bind.CallOpts{Context: t.Context()}, defaultAdminRole, timelockAddress)
		require.NoError(t, err)
		require.True(t, hasRole)
		hasRole, err = token.HasRole(&bind.CallOpts{Context: t.Context()}, defaultAdminRole, customer)
		require.NoError(t, err)
		require.False(t, hasRole)

		// We should be able to revoke admin if no fallback address is provided
		revokeOutput, err := revokeTokenAdminRoleForTest(t, env, selector, token.Address().Hex(), "", "")
		require.NoError(t, err)
		testhelpers.ProcessTimelockProposals(t, *env, revokeOutput.MCMSTimelockProposals, false)

		// Verify token roles (neither timelock nor customer should be admins)
		hasRole, err = token.HasRole(&bind.CallOpts{Context: t.Context()}, defaultAdminRole, timelockAddress)
		require.NoError(t, err)
		require.False(t, hasRole)
		hasRole, err = token.HasRole(&bind.CallOpts{Context: t.Context()}, defaultAdminRole, customer)
		require.NoError(t, err)
		require.False(t, hasRole)
	})

	t.Run("apply revoke protections", func(t *testing.T) {
		// Only make timelock an admin on the token
		token := DeployBurnMintTokenEVM(t, env, selector, timelockAddress.Hex())
		_ = DeployBurnMintPoolEVM(t, env, selector, cciputils.Version_1_6_1, token.Address())
		defaultAdminRole, err := token.DEFAULTADMINROLE(&bind.CallOpts{Context: t.Context()})
		require.NoError(t, err)

		// Verify token roles (timelock is an admin, customer is not)
		hasRole, err := token.HasRole(&bind.CallOpts{Context: t.Context()}, defaultAdminRole, timelockAddress)
		require.NoError(t, err)
		require.True(t, hasRole)
		hasRole, err = token.HasRole(&bind.CallOpts{Context: t.Context()}, defaultAdminRole, customer)
		require.NoError(t, err)
		require.False(t, hasRole)

		// We should be able to revoke admin if no fallback address is provided
		revokeOutput, err := revokeTokenAdminRoleForTest(t, env, selector, token.Address().Hex(), "", customer.Hex())
		require.NoError(t, err)
		testhelpers.ProcessTimelockProposals(t, *env, revokeOutput.MCMSTimelockProposals, false)

		// Verify token roles (customer is an admin, timelock is not)
		hasRole, err = token.HasRole(&bind.CallOpts{Context: t.Context()}, defaultAdminRole, timelockAddress)
		require.NoError(t, err)
		require.False(t, hasRole)
		hasRole, err = token.HasRole(&bind.CallOpts{Context: t.Context()}, defaultAdminRole, customer)
		require.NoError(t, err)
		require.True(t, hasRole)
	})
}

func TestRevokeTokenAdminRoleDeployer(t *testing.T) {
	// Setup environment
	selector := chainsel.TEST_90000001.Selector
	env, err := environment.New(t.Context(),
		environment.WithEVMSimulatedWithConfig(t,
			[]uint64{selector},
			onchain.EVMSimLoaderConfig{NumAdditionalAccounts: 1},
		),
	)
	require.NoError(t, err)

	// Get customer address from environment
	chain, ok := env.BlockChains.EVMChains()[selector]
	require.True(t, ok)
	require.NotEmpty(t, chain.Users)
	customer := chain.Users[0].From

	// Deploy contracts
	output, err := deployapi.DeployContracts(deployapi.GetRegistry()).Apply(*env,
		deployapi.ContractDeploymentConfig{
			MCMS: mcms.Input{},
			Chains: map[uint64]deployapi.ContractDeploymentConfigPerChain{
				selector: NewDefaultDeploymentConfigForEVM(cciputils.Version_1_6_0),
			},
		},
	)
	require.NoError(t, err)
	MergeAddresses(t, env, output.DataStore)

	t.Run("revoke deployer key while customer admin remains", func(t *testing.T) {
		// Make the deployer and customer admins on the token
		token := DeployBurnMintTokenEVM(t, env, selector, customer.Hex())
		_ = DeployBurnMintPoolEVM(t, env, selector, cciputils.Version_1_6_1, token.Address())
		defaultAdminRole, err := token.DEFAULTADMINROLE(&bind.CallOpts{Context: t.Context()})
		require.NoError(t, err)

		// Verify token roles (deployer and customer are both admins)
		hasRole, err := token.HasRole(&bind.CallOpts{Context: t.Context()}, defaultAdminRole, chain.DeployerKey.From)
		require.NoError(t, err)
		require.True(t, hasRole)
		hasRole, err = token.HasRole(&bind.CallOpts{Context: t.Context()}, defaultAdminRole, customer)
		require.NoError(t, err)
		require.True(t, hasRole)

		// The changeset should default to revoking the deployer key since timelock is not in the datastore
		revokeOutput, err := revokeTokenAdminRoleForTest(t, env, selector, token.Address().Hex(), "", customer.Hex())
		require.NoError(t, err)
		require.Empty(t, revokeOutput.MCMSTimelockProposals)

		// Verify token roles (customer is still an admin, deployer is not)
		hasRole, err = token.HasRole(&bind.CallOpts{Context: t.Context()}, defaultAdminRole, chain.DeployerKey.From)
		require.NoError(t, err)
		require.False(t, hasRole)
		hasRole, err = token.HasRole(&bind.CallOpts{Context: t.Context()}, defaultAdminRole, customer)
		require.NoError(t, err)
		require.True(t, hasRole)
	})
}

func revokeTokenAdminRoleForTest(t *testing.T,
	env *cldf_deployment.Environment,
	chainSelector uint64,
	tokenAddress string,
	adminAddress string,
	fallbackAddress string,
) (cldf_deployment.ChangesetOutput, error) {
	t.Helper()

	env.OperationsBundle = cldf_ops.NewBundle(env.GetContext, env.Logger, cldf_ops.NewMemoryReporter())
	return tokensapi.RevokeTokenAdminRole().Apply(*env, tokensapi.RevokeTokenAdminRoleInput{
		MCMS: NewDefaultInputForMCMS("Revoke admin role on token"),
		Revocations: []tokensapi.RevokeTokenAdminRoleConfig{
			{
				ChainSelector:   chainSelector,
				FallbackAddress: fallbackAddress,
				AdminAddress:    adminAddress,
				TokenRef: datastore.AddressRef{
					Address: tokenAddress,
				},
			},
		},
	})
}
