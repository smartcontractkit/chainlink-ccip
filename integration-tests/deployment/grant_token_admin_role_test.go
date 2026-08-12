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

func TestGrantTokenAdminRoleVerifyPreconditions(t *testing.T) {
	chainSelector := chainsel.TEST_90000001.Selector
	env, err := environment.New(t.Context(), environment.WithEVMSimulated(t, []uint64{chainSelector}))
	require.NoError(t, err)

	tokenRef := datastore.AddressRef{
		ChainSelector: chainSelector,
		Address:       common.HexToAddress("0x00000000000000000000000000000000000000aa").Hex(),
		Type:          datastore.ContractType(bnmERC20ops.ContractType),
		Version:       cciputils.Version_1_0_0,
		Qualifier:     "GRANT_VERIFY",
	}
	ds := datastore.NewMemoryDataStore()
	require.NoError(t, ds.Addresses().Add(tokenRef))
	MergeAddresses(t, env, ds)

	cs := tokensapi.GrantTokenAdminRole()
	tests := []struct {
		input tokensapi.GrantTokenAdminRoleInput
		title string
		error string
	}{
		{
			title: "requires chain adapter version",
			error: "chain adapter version is required",
			input: tokensapi.GrantTokenAdminRoleInput{
				Grants: []tokensapi.GrantTokenAdminRoleConfig{{
					ChainSelector: chainSelector,
					TokenRef:      tokenRef,
					AdminAddress:  common.HexToAddress("0x00000000000000000000000000000000000000bb").Hex(),
				}},
			},
		},
		{
			title: "requires grants",
			error: "at least one token admin role grant is required",
			input: tokensapi.GrantTokenAdminRoleInput{ChainAdapterVersion: cciputils.Version_1_0_0},
		},
		{
			title: "requires chain in environment",
			error: "not found in environment",
			input: tokensapi.GrantTokenAdminRoleInput{
				ChainAdapterVersion: cciputils.Version_1_0_0,
				Grants: []tokensapi.GrantTokenAdminRoleConfig{{
					ChainSelector: chainSelector + 1,
					TokenRef:      datastore.AddressRef{ChainSelector: chainSelector + 1},
					AdminAddress:  common.HexToAddress("0x00000000000000000000000000000000000000bb").Hex(),
				}},
			},
		},
		{
			title: "requires token ref",
			error: "token ref is required",
			input: tokensapi.GrantTokenAdminRoleInput{
				ChainAdapterVersion: cciputils.Version_1_0_0,
				Grants: []tokensapi.GrantTokenAdminRoleConfig{{
					ChainSelector: chainSelector,
					AdminAddress:  common.HexToAddress("0x00000000000000000000000000000000000000bb").Hex(),
				}},
			},
		},
		{
			title: "rejects token ref chain mismatch",
			error: "chain selector mismatch",
			input: tokensapi.GrantTokenAdminRoleInput{
				ChainAdapterVersion: cciputils.Version_1_0_0,
				Grants: []tokensapi.GrantTokenAdminRoleConfig{{
					ChainSelector: chainSelector,
					TokenRef: datastore.AddressRef{
						ChainSelector: chainSelector + 1,
						Address:       tokenRef.Address,
					},
					AdminAddress: common.HexToAddress("0x00000000000000000000000000000000000000bb").Hex(),
				}},
			},
		},
		{
			title: "requires admin address",
			error: "admin address is required",
			input: tokensapi.GrantTokenAdminRoleInput{
				ChainAdapterVersion: cciputils.Version_1_0_0,
				Grants: []tokensapi.GrantTokenAdminRoleConfig{{
					ChainSelector: chainSelector,
					TokenRef:      tokenRef,
				}},
			},
		},
		{
			title: "accepts datastore token ref by address",
			input: tokensapi.GrantTokenAdminRoleInput{
				ChainAdapterVersion: cciputils.Version_1_0_0,
				Grants: []tokensapi.GrantTokenAdminRoleConfig{{
					ChainSelector: chainSelector,
					AdminAddress:  common.HexToAddress("0x00000000000000000000000000000000000000bb").Hex(),
					TokenRef: datastore.AddressRef{
						Address: tokenRef.Address,
					},
				}},
			},
		},
		{
			title: "accepts address and type without datastore version",
			input: tokensapi.GrantTokenAdminRoleInput{
				ChainAdapterVersion: cciputils.Version_1_0_0,
				Grants: []tokensapi.GrantTokenAdminRoleConfig{{
					ChainSelector: chainSelector,
					AdminAddress:  common.HexToAddress("0x00000000000000000000000000000000000000bb").Hex(),
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

func TestGrantTokenAdminRoleTimelock(t *testing.T) {
	selector := chainsel.TEST_90000001.Selector
	env, err := environment.New(t.Context(),
		environment.WithEVMSimulatedWithConfig(t,
			[]uint64{selector},
			onchain.EVMSimLoaderConfig{NumAdditionalAccounts: 1},
		),
	)
	require.NoError(t, err)

	chain, ok := env.BlockChains.EVMChains()[selector]
	require.True(t, ok)
	require.NotEmpty(t, chain.Users)
	customer := chain.Users[0].From

	SeedUltraFastCurseMCMS(t, env)
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

	DeployMCMS(t, env, selector, []string{cciputils.CLLQualifier})
	mcmsReader, ok := changesets.GetRegistry().GetMCMSReader(chainsel.FamilyEVM)
	require.True(t, ok)
	timelockRef, err := mcmsReader.GetTimelockRef(*env, selector, mcms.Input{Qualifier: cciputils.CLLQualifier})
	require.NoError(t, err)
	require.True(t, common.IsHexAddress(timelockRef.Address))
	timelockAddress := common.HexToAddress(timelockRef.Address)

	t.Run("grant admin to customer via timelock", func(t *testing.T) {
		token := DeployBurnMintTokenEVM(t, env, selector, timelockAddress.Hex())
		_ = DeployBurnMintPoolEVM(t, env, selector, cciputils.Version_1_6_1, token.Address())
		defaultAdminRole, err := token.DEFAULTADMINROLE(&bind.CallOpts{Context: t.Context()})
		require.NoError(t, err)

		hasRole, err := token.HasRole(&bind.CallOpts{Context: t.Context()}, defaultAdminRole, customer)
		require.NoError(t, err)
		require.False(t, hasRole)

		grantOutput, err := grantTokenAdminRoleForTest(t, env, selector, token.Address().Hex(), customer.Hex())
		require.NoError(t, err)
		require.Len(t, grantOutput.MCMSTimelockProposals, 1)
		testhelpers.ProcessTimelockProposals(t, *env, grantOutput.MCMSTimelockProposals, false)

		hasRole, err = token.HasRole(&bind.CallOpts{Context: t.Context()}, defaultAdminRole, customer)
		require.NoError(t, err)
		require.True(t, hasRole)

		grantOutput, err = grantTokenAdminRoleForTest(t, env, selector, token.Address().Hex(), customer.Hex())
		require.NoError(t, err)
		require.Empty(t, grantOutput.MCMSTimelockProposals)
	})

	t.Run("skip when caller lacks admin role", func(t *testing.T) {
		token := DeployBurnMintTokenEVM(t, env, selector, "")
		_ = DeployBurnMintPoolEVM(t, env, selector, cciputils.Version_1_6_1, token.Address())
		defaultAdminRole, err := token.DEFAULTADMINROLE(&bind.CallOpts{Context: t.Context()})
		require.NoError(t, err)

		hasRole, err := token.HasRole(&bind.CallOpts{Context: t.Context()}, defaultAdminRole, timelockAddress)
		require.NoError(t, err)
		require.True(t, hasRole)
		hasRole, err = token.HasRole(&bind.CallOpts{Context: t.Context()}, defaultAdminRole, chain.DeployerKey.From)
		require.NoError(t, err)
		require.False(t, hasRole)

		revokeOutput, err := revokeTokenAdminRoleForTest(t, env, selector, token.Address().Hex(), "", "")
		require.NoError(t, err)
		testhelpers.ProcessTimelockProposals(t, *env, revokeOutput.MCMSTimelockProposals, false)

		hasRole, err = token.HasRole(&bind.CallOpts{Context: t.Context()}, defaultAdminRole, timelockAddress)
		require.NoError(t, err)
		require.False(t, hasRole)

		grantOutput, err := grantTokenAdminRoleForTest(t, env, selector, token.Address().Hex(), customer.Hex())
		require.NoError(t, err)
		require.Empty(t, grantOutput.MCMSTimelockProposals)

		hasRole, err = token.HasRole(&bind.CallOpts{Context: t.Context()}, defaultAdminRole, customer)
		require.NoError(t, err)
		require.False(t, hasRole)
	})
}

func TestGrantTokenAdminRoleDeployer(t *testing.T) {
	selector := chainsel.TEST_90000001.Selector
	env, err := environment.New(t.Context(),
		environment.WithEVMSimulatedWithConfig(t,
			[]uint64{selector},
			onchain.EVMSimLoaderConfig{NumAdditionalAccounts: 1},
		),
	)
	require.NoError(t, err)

	chain, ok := env.BlockChains.EVMChains()[selector]
	require.True(t, ok)
	require.NotEmpty(t, chain.Users)
	customer := chain.Users[0].From

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

	t.Run("grant admin to customer via deployer", func(t *testing.T) {
		token := DeployBurnMintTokenEVM(t, env, selector, "")
		_ = DeployBurnMintPoolEVM(t, env, selector, cciputils.Version_1_6_1, token.Address())
		defaultAdminRole, err := token.DEFAULTADMINROLE(&bind.CallOpts{Context: t.Context()})
		require.NoError(t, err)

		hasRole, err := token.HasRole(&bind.CallOpts{Context: t.Context()}, defaultAdminRole, customer)
		require.NoError(t, err)
		require.False(t, hasRole)

		hasRole, err = token.HasRole(&bind.CallOpts{Context: t.Context()}, defaultAdminRole, chain.DeployerKey.From)
		require.NoError(t, err)
		require.True(t, hasRole)

		grantOutput, err := grantTokenAdminRoleForTest(t, env, selector, token.Address().Hex(), customer.Hex())
		require.NoError(t, err)
		require.Empty(t, grantOutput.MCMSTimelockProposals)

		hasRole, err = token.HasRole(&bind.CallOpts{Context: t.Context()}, defaultAdminRole, customer)
		require.NoError(t, err)
		require.True(t, hasRole)

		grantOutput, err = grantTokenAdminRoleForTest(t, env, selector, token.Address().Hex(), customer.Hex())
		require.NoError(t, err)
		require.Empty(t, grantOutput.MCMSTimelockProposals)
	})
}

func grantTokenAdminRoleForTest(t *testing.T,
	env *cldf_deployment.Environment,
	chainSelector uint64,
	tokenAddress string,
	adminAddress string,
) (cldf_deployment.ChangesetOutput, error) {
	t.Helper()

	env.OperationsBundle = cldf_ops.NewBundle(env.GetContext, env.Logger, cldf_ops.NewMemoryReporter())
	return tokensapi.GrantTokenAdminRole().Apply(*env, tokensapi.GrantTokenAdminRoleInput{
		MCMS:                NewDefaultInputForMCMS("Grant admin role on token"),
		ChainAdapterVersion: cciputils.Version_1_0_0,
		Grants: []tokensapi.GrantTokenAdminRoleConfig{
			{
				ChainSelector: chainSelector,
				AdminAddress:  adminAddress,
				TokenRef: datastore.AddressRef{
					Address: tokenAddress,
				},
			},
		},
	})
}
