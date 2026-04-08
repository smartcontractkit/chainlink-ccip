package deployment

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	"github.com/stretchr/testify/require"

	evmadaptersV1_0_0 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/adapters"
	evmadaptersV1_6_0 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/adapters"
	evmonrampops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/onramp"
	evmseqV1_6_0 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/sequences"
	soladaptersV1_6_0 "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_0/adapters"
	solseqV1_6_0 "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_0/sequences"
	"github.com/smartcontractkit/chainlink-ccip/deployment/deploy"
	"github.com/smartcontractkit/chainlink-ccip/deployment/fees"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
)

func TestSetFeeAggregatorEVM_V1_6_0(t *testing.T) {
	evmSel := chainsel.TEST_90000001.Selector

	env, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{evmSel}),
	)
	require.NoError(t, err)

	evmAdapter := evmseqV1_6_0.EVMAdapter{}

	deployRegistry := deploy.GetRegistry()
	deployRegistry.RegisterDeployer(chainsel.FamilyEVM, deploy.MCMSVersion, &evmadaptersV1_0_0.EVMDeployer{})

	mcmsRegistry := changesets.GetRegistry()
	mcmsRegistry.RegisterMCMSReader(chainsel.FamilyEVM, &evmadaptersV1_0_0.EVMMCMSReader{})

	// Adapter is registered via init() in the adapters package import above.
	feeAggAdapter := evmadaptersV1_6_0.NewFeeAggregatorAdapter(&evmAdapter)

	output, err := deploy.DeployContracts(deployRegistry).Apply(*env, deploy.ContractDeploymentConfig{
		MCMS: mcms.Input{},
		Chains: map[uint64]deploy.ContractDeploymentConfigPerChain{
			evmSel: NewDefaultDeploymentConfigForEVM(utils.Version_1_6_0),
		},
	})
	require.NoError(t, err)
	require.NoError(t, output.DataStore.Merge(env.DataStore))
	env.DataStore = output.DataStore.Seal()

	chain := env.BlockChains.EVMChains()[evmSel]
	newFeeAggregator := common.HexToAddress("0x1234567890abcdef1234567890abcdef12345678")

	t.Run("default resolution", func(t *testing.T) {
		initialFeeAgg, err := feeAggAdapter.GetFeeAggregator(*env, evmSel)
		require.NoError(t, err)
		require.Equal(t, chain.DeployerKey.From.Hex(), common.HexToAddress(initialFeeAgg).Hex(),
			"initially the fee aggregator should be the deployer address")

		_, err = fees.SetFeeAggregator().Apply(*env, fees.SetFeeAggregatorInput{
			Version: utils.Version_1_6_0,
			MCMS:    mcms.Input{},
			Args: []fees.FeeAggregatorForChain{
				{
					ChainSelector: evmSel,
					FeeAggregator: newFeeAggregator.Hex(),
				},
			},
		})
		require.NoError(t, err)

		updatedFeeAgg, err := feeAggAdapter.GetFeeAggregator(*env, evmSel)
		require.NoError(t, err)
		require.Equal(t, newFeeAggregator.Hex(), common.HexToAddress(updatedFeeAgg).Hex())
	})

	t.Run("explicit contract ref", func(t *testing.T) {
		anotherFeeAgg := common.HexToAddress("0xdeadbeefdeadbeefdeadbeefdeadbeefdeadbeef")

		_, err = fees.SetFeeAggregator().Apply(*env, fees.SetFeeAggregatorInput{
			Version: utils.Version_1_6_0,
			MCMS:    mcms.Input{},
			Args: []fees.FeeAggregatorForChain{
				{
					ChainSelector: evmSel,
					FeeAggregator: anotherFeeAgg.Hex(),
					Contracts: []datastore.AddressRef{
						{
							Type:    datastore.ContractType(evmonrampops.ContractType),
							Version: utils.Version_1_6_0,
						},
					},
				},
			},
		})
		require.NoError(t, err)

		updatedFeeAgg, err := feeAggAdapter.GetFeeAggregator(*env, evmSel)
		require.NoError(t, err)
		require.Equal(t, anotherFeeAgg.Hex(), common.HexToAddress(updatedFeeAgg).Hex())
	})

	t.Run("idempotency", func(t *testing.T) {
		currentFeeAgg, err := feeAggAdapter.GetFeeAggregator(*env, evmSel)
		require.NoError(t, err)

		_, err = fees.SetFeeAggregator().Apply(*env, fees.SetFeeAggregatorInput{
			Version: utils.Version_1_6_0,
			MCMS:    mcms.Input{},
			Args: []fees.FeeAggregatorForChain{
				{
					ChainSelector: evmSel,
					FeeAggregator: currentFeeAgg,
				},
			},
		})
		require.NoError(t, err)

		afterFeeAgg, err := feeAggAdapter.GetFeeAggregator(*env, evmSel)
		require.NoError(t, err)
		require.Equal(t, currentFeeAgg, afterFeeAgg)
	})
}

func TestSetFeeAggregatorSolana_V1_6_0(t *testing.T) {
	solSel := chainsel.SOLANA_DEVNET.Selector

	programsPath, ds, err := PreloadSolanaEnvironment(t, solSel)
	require.NoError(t, err)

	env, err := environment.New(t.Context(),
		environment.WithSolanaContainer(t, []uint64{solSel}, programsPath, solanaProgramIDs),
	)
	require.NoError(t, err)
	env.DataStore = ds.Seal()

	solAdapter := solseqV1_6_0.SolanaAdapter{}

	deployRegistry := deploy.GetRegistry()
	deployRegistry.RegisterDeployer(chainsel.FamilySolana, deploy.MCMSVersion, &solAdapter)

	// Adapter is registered via init() in the adapters package import above.
	feeAggAdapter := soladaptersV1_6_0.NewFeeAggregatorAdapter(&solAdapter)

	output, err := deploy.DeployContracts(deployRegistry).Apply(*env, deploy.ContractDeploymentConfig{
		MCMS: mcms.Input{},
		Chains: map[uint64]deploy.ContractDeploymentConfigPerChain{
			solSel: NewDefaultDeploymentConfigForSolana(utils.Version_1_6_0),
		},
	})
	require.NoError(t, err)
	require.NoError(t, output.DataStore.Merge(env.DataStore))
	env.DataStore = output.DataStore.Seal()

	initialFeeAgg, err := feeAggAdapter.GetFeeAggregator(*env, solSel)
	require.NoError(t, err)
	require.NotEmpty(t, initialFeeAgg)
	t.Logf("Initial Solana fee aggregator: %s", initialFeeAgg)

	chain := env.BlockChains.SolanaChains()[solSel]
	newFeeAggregator := chain.DeployerKey.PublicKey().String()

	_, err = fees.SetFeeAggregator().Apply(*env, fees.SetFeeAggregatorInput{
		Version: utils.Version_1_6_0,
		MCMS:    mcms.Input{},
		Args: []fees.FeeAggregatorForChain{
			{
				ChainSelector: solSel,
				FeeAggregator: newFeeAggregator,
			},
		},
	})
	require.NoError(t, err)

	updatedFeeAgg, err := feeAggAdapter.GetFeeAggregator(*env, solSel)
	require.NoError(t, err)
	require.Equal(t, newFeeAggregator, updatedFeeAgg,
		"fee aggregator should be updated to the new address")
}

func TestSetFeeAggregator_ValidationErrors(t *testing.T) {
	evmSel := chainsel.TEST_90000001.Selector

	env, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{evmSel}),
	)
	require.NoError(t, err)

	cs := fees.SetFeeAggregator()

	t.Run("nil version", func(t *testing.T) {
		err := cs.VerifyPreconditions(*env, fees.SetFeeAggregatorInput{
			Version: nil,
			Args: []fees.FeeAggregatorForChain{
				{ChainSelector: evmSel, FeeAggregator: "0x1234"},
			},
		})
		require.Error(t, err)
		require.Contains(t, err.Error(), "version is required")
	})

	t.Run("empty args", func(t *testing.T) {
		err := cs.VerifyPreconditions(*env, fees.SetFeeAggregatorInput{
			Version: utils.Version_1_6_0,
			Args:    []fees.FeeAggregatorForChain{},
		})
		require.Error(t, err)
		require.Contains(t, err.Error(), "at least one chain must be specified")
	})

	t.Run("empty fee aggregator address", func(t *testing.T) {
		err := cs.VerifyPreconditions(*env, fees.SetFeeAggregatorInput{
			Version: utils.Version_1_6_0,
			Args: []fees.FeeAggregatorForChain{
				{ChainSelector: evmSel, FeeAggregator: ""},
			},
		})
		require.Error(t, err)
		require.Contains(t, err.Error(), "fee aggregator address is required")
	})

	t.Run("duplicate chain selector", func(t *testing.T) {
		err := cs.VerifyPreconditions(*env, fees.SetFeeAggregatorInput{
			Version: utils.Version_1_6_0,
			Args: []fees.FeeAggregatorForChain{
				{ChainSelector: evmSel, FeeAggregator: "0x1234"},
				{ChainSelector: evmSel, FeeAggregator: "0x5678"},
			},
		})
		require.Error(t, err)
		require.Contains(t, err.Error(), "duplicate chain selector")
	})
}
