package sequences_test

import (
	"testing"

	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/advanced_pool_hooks_extractor"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/sequences"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	"github.com/stretchr/testify/require"
)

func TestDeployAdvancedPoolHooksExtractor(t *testing.T) {
	chainSelector := uint64(5009297550715157269)
	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{chainSelector}),
	)
	require.NoError(t, err)
	require.NotNil(t, e)

	report, err := operations.ExecuteSequence(
		e.OperationsBundle,
		sequences.DeployAdvancedPoolHooksExtractor,
		e.BlockChains.EVMChains()[chainSelector],
		sequences.DeployAdvancedPoolHooksExtractorInput{
			ChainSelector: chainSelector,
		},
	)
	require.NoError(t, err)
	require.Len(t, report.Output.Addresses, 1)

	ref := report.Output.Addresses[0]
	require.Equal(t, datastore.ContractType(advanced_pool_hooks_extractor.ContractType), ref.Type)
	require.Equal(t, advanced_pool_hooks_extractor.Version.String(), ref.Version.String())
	require.Equal(t, chainSelector, ref.ChainSelector)
	require.NotEmpty(t, ref.Address)
}

func TestDeployAdvancedPoolHooksExtractor_Idempotency(t *testing.T) {
	chainSelector := uint64(5009297550715157269)
	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{chainSelector}),
	)
	require.NoError(t, err)
	require.NotNil(t, e)

	report1, err := operations.ExecuteSequence(
		e.OperationsBundle,
		sequences.DeployAdvancedPoolHooksExtractor,
		e.BlockChains.EVMChains()[chainSelector],
		sequences.DeployAdvancedPoolHooksExtractorInput{
			ChainSelector: chainSelector,
		},
	)
	require.NoError(t, err)
	require.Len(t, report1.Output.Addresses, 1)

	// Run again with the deployed address as existing
	report2, err := operations.ExecuteSequence(
		e.OperationsBundle,
		sequences.DeployAdvancedPoolHooksExtractor,
		e.BlockChains.EVMChains()[chainSelector],
		sequences.DeployAdvancedPoolHooksExtractorInput{
			ChainSelector:     chainSelector,
			ExistingAddresses: report1.Output.Addresses,
		},
	)
	require.NoError(t, err)
	require.Len(t, report2.Output.Addresses, 1)

	require.Equal(t, report1.Output.Addresses[0].Address, report2.Output.Addresses[0].Address)
	require.Equal(t, report1.Output.Addresses[0].Type, report2.Output.Addresses[0].Type)
	require.Equal(t, report1.Output.Addresses[0].Version.String(), report2.Output.Addresses[0].Version.String())
}

func TestDeployAdvancedPoolHooksExtractor_MultipleChains(t *testing.T) {
	chainSelectors := []uint64{5009297550715157269, 4949039107694359620}
	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, chainSelectors),
	)
	require.NoError(t, err)
	require.NotNil(t, e)

	addresses := make(map[uint64]string)
	for _, sel := range chainSelectors {
		report, err := operations.ExecuteSequence(
			e.OperationsBundle,
			sequences.DeployAdvancedPoolHooksExtractor,
			e.BlockChains.EVMChains()[sel],
			sequences.DeployAdvancedPoolHooksExtractorInput{
				ChainSelector: sel,
			},
		)
		require.NoError(t, err)
		require.Len(t, report.Output.Addresses, 1)

		ref := report.Output.Addresses[0]
		require.Equal(t, datastore.ContractType(advanced_pool_hooks_extractor.ContractType), ref.Type)
		addresses[sel] = ref.Address
	}

	require.Len(t, addresses, len(chainSelectors))
	require.NotEqual(t, addresses[chainSelectors[0]], addresses[chainSelectors[1]])
}
