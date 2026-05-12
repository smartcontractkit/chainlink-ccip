package deployment

import (
	"fmt"
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	chainsel "github.com/smartcontractkit/chain-selectors"
	evmadapters "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/adapters"
	bnmERC20ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/burn_mint_erc20"
	deployapi "github.com/smartcontractkit/chainlink-ccip/deployment/deploy"
	"github.com/smartcontractkit/chainlink-ccip/deployment/testhelpers"
	tokensapi "github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	cciputils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/onchain"
	bnmERC20gen "github.com/smartcontractkit/chainlink-evm/gethwrappers/shared/generated/initial/burn_mint_erc20"
	"github.com/stretchr/testify/require"

	_ "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_1/adapters"
)

func TestRevokeTokenAdminRoleEVMBurnMint(t *testing.T) {
	v1_6_0 := semver.MustParse("1.6.0")
	v1_6_1 := semver.MustParse("1.6.1")
	chainSelector := chainsel.TEST_90000001.Selector

	env, err := environment.New(t.Context(), environment.WithEVMSimulatedWithConfig(t, []uint64{chainSelector}, onchain.EVMSimLoaderConfig{
		NumAdditionalAccounts: 1,
	}))
	require.NoError(t, err)

	chain, ok := env.BlockChains.EVMChains()[chainSelector]
	require.True(t, ok)

	deployRegistry := deployapi.GetRegistry()
	deployRegistry.RegisterDeployer(chainsel.FamilyEVM, deployapi.MCMSVersion, &evmadapters.EVMDeployer{})
	mcmsRegistry := changesets.GetRegistry()
	mcmsRegistry.RegisterMCMSReader(chainsel.FamilyEVM, &evmadapters.EVMMCMSReader{})

	output, err := deployapi.DeployContracts(deployRegistry).Apply(*env, deployapi.ContractDeploymentConfig{
		Chains: map[uint64]deployapi.ContractDeploymentConfigPerChain{
			chainSelector: NewDefaultDeploymentConfigForEVM(v1_6_0),
		},
		MCMS: mcms.Input{},
	})
	require.NoError(t, err)
	MergeAddresses(t, env, output.DataStore)

	DeployMCMS(t, env, chainSelector, []string{cciputils.CLLQualifier})
	mcmsReader, ok := mcmsRegistry.GetMCMSReader(chainsel.FamilyEVM)
	require.True(t, ok)
	timelockRef, err := mcmsReader.GetTimelockRef(*env, chainSelector, mcms.Input{Qualifier: cciputils.CLLQualifier})
	require.NoError(t, err)
	timelockAddress := common.HexToAddress(timelockRef.Address)

	t.Run("revoke timelock admin while customer admin remains", func(t *testing.T) {
		require.NotEmpty(t, chain.Users)
		customerAdmin := chain.Users[0].From
		token := deployBurnMintTokenForAdminRevocation(t, env, chainSelector, v1_6_1, "REVOKE_POS", "REVOKE_POS_POOL", customerAdmin.Hex())
		tokenContract, defaultAdminRole := loadBurnMintTokenAdminRole(t, chain.Client, token.Address)

		requireBurnMintAdminRole(t, tokenContract, defaultAdminRole, timelockAddress, true)
		requireBurnMintAdminRole(t, tokenContract, defaultAdminRole, customerAdmin, true)

		revokeOutput, err := revokeTokenAdminRoleForTest(t, env, chainSelector, token, "Revoke token admin role")
		require.NoError(t, err)
		require.Len(t, revokeOutput.MCMSTimelockProposals, 1)
		testhelpers.ProcessTimelockProposals(t, *env, revokeOutput.MCMSTimelockProposals, false)

		requireBurnMintAdminRole(t, tokenContract, defaultAdminRole, timelockAddress, false)
		requireBurnMintAdminRole(t, tokenContract, defaultAdminRole, customerAdmin, true)

		revokeOutput, err = revokeTokenAdminRoleForTest(t, env, chainSelector, token, "Revoke token admin role idempotent")
		require.NoError(t, err)
		require.Empty(t, revokeOutput.MCMSTimelockProposals)
	})

	t.Run("reject revoke when timelock is the only known admin", func(t *testing.T) {
		token := deployBurnMintTokenForAdminRevocation(t, env, chainSelector, v1_6_1, "REVOKE_NEG", "REVOKE_NEG_POOL", "")
		tokenContract, defaultAdminRole := loadBurnMintTokenAdminRole(t, chain.Client, token.Address)

		requireBurnMintAdminRole(t, tokenContract, defaultAdminRole, timelockAddress, true)
		requireBurnMintAdminRole(t, tokenContract, defaultAdminRole, chain.DeployerKey.From, false)

		_, err = revokeTokenAdminRoleForTest(t, env, chainSelector, token, "Reject unsafe token admin role revoke")
		require.Error(t, err)
		require.Contains(t, err.Error(), "no remaining admin could be confirmed")
	})
}

func revokeTokenAdminRoleForTest(t *testing.T, env *cldf_deployment.Environment, chainSelector uint64, token datastore.AddressRef, description string) (cldf_deployment.ChangesetOutput, error) {
	t.Helper()

	return tokensapi.RevokeTokenAdminRole().Apply(*env, tokensapi.RevokeTokenAdminRoleInput{
		Revocations: []tokensapi.RevokeTokenAdminRoleConfig{
			{
				ChainSelector: chainSelector,
				TokenRef: datastore.AddressRef{
					Qualifier: token.Qualifier,
				},
			},
		},
		MCMS: NewDefaultInputForMCMS(description),
	})
}

func requireBurnMintAdminRole(t *testing.T, tokenContract *bnmERC20gen.BurnMintERC20, role [32]byte, account common.Address, expected bool) {
	t.Helper()

	hasRole, err := tokenContract.HasRole(&bind.CallOpts{Context: t.Context()}, role, account)
	require.NoError(t, err)
	require.Equal(t, expected, hasRole)
}

func deployBurnMintTokenForAdminRevocation(t *testing.T, env *cldf_deployment.Environment, chainSelector uint64, tokenPoolVersion *semver.Version, tokenSymbol string, tokenPoolQualifier string, externalAdmin string) datastore.AddressRef {
	t.Helper()

	maxSupply := uint64(1e6)
	output, err := tokensapi.TokenExpansion().Apply(*env, tokensapi.TokenExpansionInput{
		TokenExpansionInputPerChain: map[uint64]tokensapi.TokenExpansionInputPerChain{
			chainSelector: {
				TokenPoolVersion: tokenPoolVersion,
				DeployTokenInput: &tokensapi.DeployTokenInput{
					Decimals:      uint8(18),
					Symbol:        tokenSymbol,
					Name:          fmt.Sprintf("%s Token", tokenSymbol),
					Type:          bnmERC20ops.ContractType,
					Supply:        &maxSupply,
					ExternalAdmin: externalAdmin,
				},
				DeployTokenPoolInput: &tokensapi.DeployTokenPoolInput{
					TokenPoolQualifier: tokenPoolQualifier,
					PoolType:           cciputils.BurnMintTokenPool.String(),
				},
			},
		},
		ChainAdapterVersion: tokenPoolVersion,
		MCMS:                NewDefaultInputForMCMS(fmt.Sprintf("Deploy %s", tokenSymbol)),
	})
	require.NoError(t, err)
	MergeAddresses(t, env, output.DataStore)

	tokenRef, err := datastore_utils.FindAndFormatRef(env.DataStore, datastore.AddressRef{
		ChainSelector: chainSelector,
		Qualifier:     tokenSymbol,
	}, chainSelector, datastore_utils.FullRef)
	require.NoError(t, err)
	return tokenRef
}

func loadBurnMintTokenAdminRole(t *testing.T, backend bind.ContractBackend, tokenAddress string) (*bnmERC20gen.BurnMintERC20, [32]byte) {
	t.Helper()

	tokenContract, err := bnmERC20gen.NewBurnMintERC20(common.HexToAddress(tokenAddress), backend)
	require.NoError(t, err)
	defaultAdminRole, err := tokenContract.DEFAULTADMINROLE(&bind.CallOpts{Context: t.Context()})
	require.NoError(t, err)
	return tokenContract, defaultAdminRole
}
