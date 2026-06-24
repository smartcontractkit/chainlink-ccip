package deployment

import (
	"math/big"
	"testing"

	"github.com/Masterminds/semver/v3"
	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	"github.com/stretchr/testify/require"

	evmadaptersV1_0_0 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/adapters"
	evmadaptersV1_6_0 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/adapters"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/onramp"
	evmseqV1_6_0 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/sequences"
	soladaptersV1_6_0 "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_0/adapters"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_0/operations/router"
	solseqV1_6_0 "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_0/sequences"
	"github.com/smartcontractkit/chainlink-ccip/deployment/deploy"
	"github.com/smartcontractkit/chainlink-ccip/deployment/fees"
	"github.com/smartcontractkit/chainlink-ccip/deployment/lanes"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
)

func TestUpdateFeeQuoterDestsV1_6_0(t *testing.T) {
	t.Run("EVM only", testUpdateFQDestsEVMOnly)
	t.Run("Cross-chain EVM+Solana", testUpdateFQDestsCrossChain)
}

func testUpdateFQDestsEVMOnly(t *testing.T) {
	v1_6_0, err := semver.NewVersion("1.6.0")
	require.NoError(t, err)

	src := chainsel.TEST_90000001.Selector
	dst := chainsel.TEST_90000002.Selector
	chains := []uint64{src, dst}

	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, chains),
	)
	require.NoError(t, err)

	mcmsRegistry := changesets.GetRegistry()

	chainInput := make(map[uint64]deploy.ContractDeploymentConfigPerChain)
	for _, sel := range chains {
		chainInput[sel] = NewDefaultDeploymentConfigForEVM(v1_6_0)
	}
	out, err := deploy.DeployContracts(deploy.GetRegistry()).Apply(*e, deploy.ContractDeploymentConfig{
		MCMS:   mcms.Input{},
		Chains: chainInput,
	})
	require.NoError(t, err)
	MergeAddresses(t, e, out.DataStore)

	_, err = lanes.ConnectChains(lanes.GetLaneAdapterRegistry(), mcmsRegistry).Apply(*e, lanes.ConnectChainsConfig{
		Lanes: []lanes.LaneConfig{
			{
				Version: v1_6_0,
				ChainA:  lanes.ChainDefinition{Selector: src, GasPrice: big.NewInt(1e9)},
				ChainB:  lanes.ChainDefinition{Selector: dst, GasPrice: big.NewInt(1e9)},
			},
		},
	})
	require.NoError(t, err)

	evmFeesAdapter := evmadaptersV1_6_0.NewFeesAdapter(&evmseqV1_6_0.EVMAdapter{})

	srcOnRampRef := datastore.AddressRef{ChainSelector: src, Type: datastore.ContractType(onramp.ContractType), Version: utils.Version_1_6_0}
	srcOnRampRef, err = datastore_utils.FindAndFormatRef(e.DataStore, srcOnRampRef, src, datastore_utils.FullRef)
	require.NoError(t, err)
	srcFQ, err := evmFeesAdapter.GetFeeContractRef(e.OperationsBundle, e.BlockChains, e.DataStore, srcOnRampRef, src, dst)
	require.NoError(t, err)

	initialCfg, err := evmFeesAdapter.GetOnchainDestChainConfig(e.OperationsBundle, e.BlockChains, srcFQ, src, dst)
	require.NoError(t, err)
	require.True(t, initialCfg.IsEnabled)

	originalDestGasOverhead := initialCfg.DestGasOverhead

	newMaxDataBytes := uint32(99_999)
	require.NotEqual(t, initialCfg.MaxDataBytes, newMaxDataBytes)

	override := lanes.FeeQuoterDestChainConfigOverride(func(c *lanes.FeeQuoterDestChainConfig) {
		c.MaxDataBytes = newMaxDataBytes
	})

	_, err = fees.
		UpdateFeeQuoterDests().
		Apply(*e, fees.UpdateFeeQuoterDestsInput{
			Version: v1_6_0,
			MCMS:    mcms.Input{},
			Args: []fees.DestChainConfigForSrc{
				{
					Selector: src,
					Settings: []fees.DestChainConfigForDst{
						{
							Selector: dst,
							Override: &override,
						},
					},
				},
			},
		})
	require.NoError(t, err)

	updatedCfg, err := evmFeesAdapter.GetOnchainDestChainConfig(e.OperationsBundle, e.BlockChains, srcFQ, src, dst)
	require.NoError(t, err)
	require.True(t, updatedCfg.IsEnabled)
	require.Equal(t, newMaxDataBytes, updatedCfg.MaxDataBytes, "MaxDataBytes should be updated")
	require.Equal(t, originalDestGasOverhead, updatedCfg.DestGasOverhead, "DestGasOverhead should be unchanged")
}

func testUpdateFQDestsCrossChain(t *testing.T) {
	v1_6_0, err := semver.NewVersion("1.6.0")
	require.NoError(t, err)

	solSel := chainsel.SOLANA_DEVNET.Selector
	evmSel := chainsel.TEST_90000002.Selector

	programsPath, ds, err := PreloadSolanaEnvironment(t, solSel)
	require.NoError(t, err)

	e, err := environment.New(t.Context(),
		environment.WithSolanaContainer(t, []uint64{solSel}, programsPath, solanaProgramIDs),
		environment.WithEVMSimulated(t, []uint64{evmSel}),
	)
	require.NoError(t, err)
	e.DataStore = ds.Seal()

	solAdapter := solseqV1_6_0.SolanaAdapter{}
	evmAdapter := evmseqV1_6_0.EVMAdapter{}

	deployRegistry := deploy.GetRegistry()
	deployRegistry.RegisterDeployer(chainsel.FamilyEVM, deploy.MCMSVersion, &evmadaptersV1_0_0.EVMDeployer{})
	deployRegistry.RegisterDeployer(chainsel.FamilySolana, deploy.MCMSVersion, &solAdapter)

	mcmsRegistry := changesets.GetRegistry()
	mcmsRegistry.RegisterMCMSReader(chainsel.FamilyEVM, &evmadaptersV1_0_0.EVMMCMSReader{})

	out, err := deploy.DeployContracts(deployRegistry).Apply(*e, deploy.ContractDeploymentConfig{
		MCMS: mcms.Input{},
		Chains: map[uint64]deploy.ContractDeploymentConfigPerChain{
			solSel: NewDefaultDeploymentConfigForSolana(v1_6_0),
			evmSel: NewDefaultDeploymentConfigForEVM(v1_6_0),
		},
	})
	require.NoError(t, err)
	MergeAddresses(t, e, out.DataStore)

	_, err = lanes.ConnectChains(lanes.GetLaneAdapterRegistry(), mcmsRegistry).Apply(*e, lanes.ConnectChainsConfig{
		Lanes: []lanes.LaneConfig{
			{
				Version: v1_6_0,
				ChainA:  lanes.ChainDefinition{Selector: solSel, GasPrice: big.NewInt(1e9)},
				ChainB:  lanes.ChainDefinition{Selector: evmSel, GasPrice: big.NewInt(1e9)},
			},
		},
	})
	require.NoError(t, err)

	evmFeesAdapter := evmadaptersV1_6_0.NewFeesAdapter(&evmAdapter)
	solFeesAdapter := soladaptersV1_6_0.NewFeesAdapter(&solAdapter)

	// Verify EVM src → Solana dst FQ dest update
	t.Run("EVM src to Solana dst", func(t *testing.T) {
		evmOnRampRef := datastore.AddressRef{ChainSelector: evmSel, Type: datastore.ContractType(onramp.ContractType), Version: utils.Version_1_6_0}
		evmOnRampRef, err = datastore_utils.FindAndFormatRef(e.DataStore, evmOnRampRef, evmSel, datastore_utils.FullRef)
		require.NoError(t, err)
		evmFQ, err := evmFeesAdapter.GetFeeContractRef(e.OperationsBundle, e.BlockChains, e.DataStore, evmOnRampRef, evmSel, solSel)
		require.NoError(t, err)

		initialCfg, err := evmFeesAdapter.GetOnchainDestChainConfig(e.OperationsBundle, e.BlockChains, evmFQ, evmSel, solSel)
		require.NoError(t, err)
		require.True(t, initialCfg.IsEnabled)

		originalDestGasOverhead := initialCfg.DestGasOverhead
		newMaxDataBytes := uint32(77_777)
		require.NotEqual(t, initialCfg.MaxDataBytes, newMaxDataBytes)

		override := lanes.FeeQuoterDestChainConfigOverride(func(c *lanes.FeeQuoterDestChainConfig) {
			c.MaxDataBytes = newMaxDataBytes
		})

		_, err = fees.
			UpdateFeeQuoterDests().
			Apply(*e, fees.UpdateFeeQuoterDestsInput{
				Version: v1_6_0,
				MCMS:    mcms.Input{},
				Args: []fees.DestChainConfigForSrc{
					{
						Selector: evmSel,
						Settings: []fees.DestChainConfigForDst{
							{
								Selector: solSel,
								Override: &override,
							},
						},
					},
				},
			})
		require.NoError(t, err)

		updatedCfg, err := evmFeesAdapter.GetOnchainDestChainConfig(e.OperationsBundle, e.BlockChains, evmFQ, evmSel, solSel)
		require.NoError(t, err)
		require.True(t, updatedCfg.IsEnabled)
		require.Equal(t, newMaxDataBytes, updatedCfg.MaxDataBytes, "MaxDataBytes should be updated")
		require.Equal(t, originalDestGasOverhead, updatedCfg.DestGasOverhead, "DestGasOverhead should be unchanged")
	})

	// Verify Solana src → EVM dst FQ dest update
	t.Run("Solana src to EVM dst", func(t *testing.T) {
		solOnRampRef := datastore.AddressRef{ChainSelector: solSel, Type: datastore.ContractType(router.ContractType), Version: utils.Version_1_6_0}
		solOnRampRef, err = datastore_utils.FindAndFormatRef(e.DataStore, solOnRampRef, solSel, datastore_utils.FullRef)
		require.NoError(t, err)
		solFQ, err := solFeesAdapter.GetFeeContractRef(e.OperationsBundle, e.BlockChains, e.DataStore, solOnRampRef, solSel, evmSel)
		require.NoError(t, err)

		initialCfg, err := solFeesAdapter.GetOnchainDestChainConfig(e.OperationsBundle, e.BlockChains, solFQ, solSel, evmSel)
		require.NoError(t, err)
		require.True(t, initialCfg.IsEnabled)

		originalDestGasOverhead := initialCfg.DestGasOverhead
		newMaxDataBytes := uint32(88_888)
		require.NotEqual(t, initialCfg.MaxDataBytes, newMaxDataBytes)

		override := lanes.FeeQuoterDestChainConfigOverride(func(c *lanes.FeeQuoterDestChainConfig) {
			c.MaxDataBytes = newMaxDataBytes
		})

		_, err = fees.
			UpdateFeeQuoterDests().
			Apply(*e, fees.UpdateFeeQuoterDestsInput{
				Version: v1_6_0,
				MCMS:    mcms.Input{},
				Args: []fees.DestChainConfigForSrc{
					{
						Selector: solSel,
						Settings: []fees.DestChainConfigForDst{
							{
								Selector: evmSel,
								Override: &override,
							},
						},
					},
				},
			})
		require.NoError(t, err)

		updatedCfg, err := solFeesAdapter.GetOnchainDestChainConfig(e.OperationsBundle, e.BlockChains, solFQ, solSel, evmSel)
		require.NoError(t, err)
		require.True(t, updatedCfg.IsEnabled)
		require.Equal(t, newMaxDataBytes, updatedCfg.MaxDataBytes, "MaxDataBytes should be updated")
		require.Equal(t, originalDestGasOverhead, updatedCfg.DestGasOverhead, "DestGasOverhead should be unchanged")
	})
}
