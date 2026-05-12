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
	bnmERC20gen "github.com/smartcontractkit/chainlink-evm/gethwrappers/shared/generated/initial/burn_mint_erc20"
	evmutils "github.com/smartcontractkit/chainlink-evm/pkg/utils"
	"github.com/stretchr/testify/require"

	_ "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_1/adapters"
)

func TestRevokeTokenAdminRoleEVMBurnMint(t *testing.T) {
	v1_6_0 := semver.MustParse("1.6.0")
	v1_6_1 := semver.MustParse("1.6.1")
	chainSelector := chainsel.TEST_90000001.Selector

	env, err := environment.New(t.Context(), environment.WithEVMSimulated(t, []uint64{chainSelector}))
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
		customerAdmin := evmutils.RandomAddress()
		token := deployBurnMintTokenForAdminRevocation(t, env, chainSelector, v1_6_1, "REVOKE_POS", "REVOKE_POS_POOL", customerAdmin.Hex())
		tokenContract, defaultAdminRole := loadBurnMintTokenAdminRole(t, chain.Client, token.Address)

		timelockHasRole, err := tokenContract.HasRole(&bind.CallOpts{Context: t.Context()}, defaultAdminRole, timelockAddress)
		require.NoError(t, err)
		require.True(t, timelockHasRole)
		customerHasRole, err := tokenContract.HasRole(&bind.CallOpts{Context: t.Context()}, defaultAdminRole, customerAdmin)
		require.NoError(t, err)
		require.True(t, customerHasRole)

		revokeOutput, err := tokensapi.RevokeTokenAdminRole().Apply(*env, tokensapi.RevokeTokenAdminRoleInput{
			Revocations: []tokensapi.RevokeTokenAdminRoleConfig{
				{
					ChainSelector: chainSelector,
					TokenRef: datastore.AddressRef{
						Qualifier: token.Qualifier,
					},
				},
			},
			MCMS: NewDefaultInputForMCMS("Revoke token admin role"),
		})
		require.NoError(t, err)
		require.Len(t, revokeOutput.MCMSTimelockProposals, 1)
		testhelpers.ProcessTimelockProposals(t, *env, revokeOutput.MCMSTimelockProposals, false)

		timelockHasRole, err = tokenContract.HasRole(&bind.CallOpts{Context: t.Context()}, defaultAdminRole, timelockAddress)
		require.NoError(t, err)
		require.False(t, timelockHasRole)
		customerHasRole, err = tokenContract.HasRole(&bind.CallOpts{Context: t.Context()}, defaultAdminRole, customerAdmin)
		require.NoError(t, err)
		require.True(t, customerHasRole)

		revokeOutput, err = tokensapi.RevokeTokenAdminRole().Apply(*env, tokensapi.RevokeTokenAdminRoleInput{
			Revocations: []tokensapi.RevokeTokenAdminRoleConfig{
				{
					ChainSelector: chainSelector,
					TokenRef: datastore.AddressRef{
						Qualifier: token.Qualifier,
					},
				},
			},
			MCMS: NewDefaultInputForMCMS("Revoke token admin role idempotent"),
		})
		require.NoError(t, err)
		require.Empty(t, revokeOutput.MCMSTimelockProposals)
	})

	t.Run("reject revoke when timelock is the only known admin", func(t *testing.T) {
		token := deployBurnMintTokenForAdminRevocation(t, env, chainSelector, v1_6_1, "REVOKE_NEG", "REVOKE_NEG_POOL", "")
		tokenContract, defaultAdminRole := loadBurnMintTokenAdminRole(t, chain.Client, token.Address)

		timelockHasRole, err := tokenContract.HasRole(&bind.CallOpts{Context: t.Context()}, defaultAdminRole, timelockAddress)
		require.NoError(t, err)
		require.True(t, timelockHasRole)
		deployerHasRole, err := tokenContract.HasRole(&bind.CallOpts{Context: t.Context()}, defaultAdminRole, chain.DeployerKey.From)
		require.NoError(t, err)
		require.False(t, deployerHasRole)

		_, err = tokensapi.RevokeTokenAdminRole().Apply(*env, tokensapi.RevokeTokenAdminRoleInput{
			Revocations: []tokensapi.RevokeTokenAdminRoleConfig{
				{
					ChainSelector: chainSelector,
					TokenRef: datastore.AddressRef{
						Qualifier: token.Qualifier,
					},
				},
			},
			MCMS: NewDefaultInputForMCMS("Reject unsafe token admin role revoke"),
		})
		require.Error(t, err)
		require.Contains(t, err.Error(), "no remaining admin could be confirmed")
	})
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
