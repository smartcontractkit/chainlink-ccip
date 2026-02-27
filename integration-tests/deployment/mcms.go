package deployment

import (
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/gagliardetto/solana-go"
	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	mcms_types "github.com/smartcontractkit/mcms/types"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/v0_1_1/mcm"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/state"
	mcmsapi "github.com/smartcontractkit/chainlink-ccip/deployment/deploy"
	"github.com/smartcontractkit/chainlink-ccip/deployment/testhelpers"

	_ "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/adapters"
	deploymentutils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
)

func TestUpdateMCMSConfigSolana(t *testing.T) {
	programsPath, dstr, err := PreloadSolanaEnvironment(t, chainsel.SOLANA_MAINNET.Selector)
	require.NoError(t, err, "Failed to set up Solana environment")
	require.NotNil(t, dstr, "Datastore should be created")
	solanaChains := []uint64{
		chainsel.SOLANA_MAINNET.Selector,
	}
	env, err := environment.New(t.Context(),
		environment.WithSolanaContainer(t, solanaChains, programsPath, solanaProgramIDs),
	)
	require.NoError(t, err)
	require.NotNil(t, env, "Environment should be created")
	env.DataStore = dstr.Seal() // Add preloaded contracts to env datastore
	chain := env.BlockChains.SolanaChains()[chainsel.SOLANA_MAINNET.Selector]

	// deploy MCMS Contracts
	DeployMCMS(t, env, solanaChains[0], []string{deploymentutils.CLLQualifier})

	// get recently deployed MCMS addresses
	mcmsRefs := []datastore.AddressRef{}
	cancellerRef, err := datastore_utils.FindAndFormatRef(env.DataStore, datastore.AddressRef{
		ChainSelector: solanaChains[0],
		Type:          datastore.ContractType(deploymentutils.CancellerManyChainMultisig),
		Qualifier:     deploymentutils.CLLQualifier,
		Version:       semver.MustParse("1.0.0"),
	}, solanaChains[0], datastore_utils.FullRef)
	require.NoError(t, err)
	bypasserRef, err := datastore_utils.FindAndFormatRef(env.DataStore, datastore.AddressRef{
		ChainSelector: solanaChains[0],
		Type:          datastore.ContractType(deploymentutils.CancellerManyChainMultisig),
		Qualifier:     deploymentutils.CLLQualifier,
		Version:       semver.MustParse("1.0.0"),
	}, solanaChains[0], datastore_utils.FullRef)
	require.NoError(t, err)
	proposerRef, err := datastore_utils.FindAndFormatRef(env.DataStore, datastore.AddressRef{
		ChainSelector: solanaChains[0],
		Type:          datastore.ContractType(deploymentutils.CancellerManyChainMultisig),
		Qualifier:     deploymentutils.CLLQualifier,
		Version:       semver.MustParse("1.0.0"),
	}, solanaChains[0], datastore_utils.FullRef)
	require.NoError(t, err)
	mcmsRefs = append(mcmsRefs, cancellerRef, bypasserRef, proposerRef)

	// check that deployed config is correct
	for _, ref := range mcmsRefs {
		var mcmConfig mcm.MultisigConfig
		mcmSeed := state.PDASeed([]byte(ref.Address))
		err := chain.GetAccountDataBorshInto(env.GetContext(), state.GetMCMConfigPDA(solana.MustPublicKeyFromBase58(ref.Address), mcmSeed), &mcmConfig)
		require.NoError(t, err)

		numOfSigners := len(mcmConfig.Signers)
		require.Equal(t, numOfSigners, len(testhelpers.SingleGroupMCMS().Signers)) // should be 1
	}

	// update the config for each MCMS contract
	dReg := mcmsapi.GetRegistry()
	updateMcmsConfigMCMS := mcmsapi.UpdateMCMSConfig(dReg, nil)
	output, err := updateMcmsConfigMCMS.Apply(*env, mcmsapi.UpdateMCMSConfigInput{
		AdapterVersion: semver.MustParse("1.0.0"),
		Chains: map[uint64]mcmsapi.UpdateMCMSConfigInputPerChain{
			solanaChains[0]: {
				MCMConfig:    testhelpers.SingleGroupMCMSTwoSigners(),
				MCMContracts: mcmsRefs,
			},
		},
		MCMS: mcms.Input{
			OverridePreviousRoot: false,
			ValidUntil:           3759765795,
			TimelockDelay:        mcms_types.MustParseDuration("0s"),
			TimelockAction:       mcms_types.TimelockActionSchedule,
			Qualifier:            deploymentutils.CLLQualifier,
			Description:          "update mcms config test",
		},
	})
	require.NoError(t, err)
	require.Greater(t, len(output.Reports), 0)

	// check that MCMS configs are updated correctly
	for _, ref := range mcmsRefs {
		var mcmConfig mcm.MultisigConfig
		mcmSeed := state.PDASeed([]byte(ref.Address))
		err := chain.GetAccountDataBorshInto(env.GetContext(), state.GetMCMConfigPDA(solana.MustPublicKeyFromBase58(ref.Address), mcmSeed), &mcmConfig)
		require.NoError(t, err)

		numOfSigners := len(mcmConfig.Signers)
		require.Equal(t, numOfSigners, len(testhelpers.SingleGroupMCMSTwoSigners().Signers)) // should be 2
	}

}
