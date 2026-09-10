package test_verifier

import (
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	"github.com/stretchr/testify/require"

	contract_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/create2_factory"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/sequences"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/testsetup"
)

func setupTestVerifierBaseChain(
	t *testing.T,
	e *deployment.Environment,
	chainSelector uint64,
) datastore.DataStore {
	t.Helper()

	chain := e.BlockChains.EVMChains()[chainSelector]

	create2FactoryRef, err := contract_utils.MaybeDeployContract(
		e.OperationsBundle,
		create2_factory.Deploy,
		chain,
		contract_utils.DeployInput[create2_factory.ConstructorArgs]{
			TypeAndVersion: deployment.NewTypeAndVersion(
				create2_factory.ContractType,
				*semver.MustParse("2.0.0"),
			),
			ChainSelector: chainSelector,
			Args: create2_factory.ConstructorArgs{
				AllowList: []common.Address{chain.DeployerKey.From},
			},
		},
		nil,
	)
	require.NoError(t, err)

	chainReport, err := operations.ExecuteSequence(
		e.OperationsBundle,
		sequences.DeployChainContracts,
		chain,
		sequences.DeployChainContractsInput{
			ChainSelector:     chainSelector,
			ContractParams:    testsetup.CreateBasicContractParams(),
			CREATE2Factory:    common.HexToAddress(create2FactoryRef.Address),
			DeployTestRouter:  true,
			DeployerKeyOwned:  true,
			ExistingAddresses: testsetup.UltraFastCurseMCMSRefs(chainSelector),
		},
	)
	require.NoError(t, err)

	ds := datastore.NewMemoryDataStore()

	for _, ref := range chainReport.Output.Addresses {
		require.NoError(t, ds.Addresses().Add(ref))
	}

	// DeployChainContracts consumes the CREATE2 factory address but does not
	// necessarily return the factory itself in Output.Addresses. The test
	// verifier deployment sequence resolves it from the datastore.
	require.NoError(t, ds.Addresses().Add(create2FactoryRef))

	return ds.Seal()
}

func mergeTestVerifierAddresses(
	t *testing.T,
	base datastore.DataStore,
	addresses []datastore.AddressRef,
) datastore.DataStore {
	t.Helper()

	ds := datastore.NewMemoryDataStore()
	require.NoError(t, ds.Merge(base))

	for _, ref := range addresses {
		// Some sequences return refs they reused from the input datastore.
		// Merge semantics are a better fit than assuming every returned ref
		// is necessarily new.
		existing := ds.Addresses().Filter(
			datastore.AddressRefByChainSelector(ref.ChainSelector),
			datastore.AddressRefByType(ref.Type),
			datastore.AddressRefByQualifier(ref.Qualifier),
		)

		found := false
		for _, current := range existing {
			if current.Address == ref.Address &&
				current.Version.Equal(ref.Version) {
				found = true
				break
			}
		}

		if !found {
			require.NoError(t, ds.Addresses().Add(ref))
		}
	}

	return ds.Seal()
}

func mergeTestVerifierDataStores(
	t *testing.T,
	stores ...datastore.DataStore,
) datastore.DataStore {
	t.Helper()

	ds := datastore.NewMemoryDataStore()
	for _, store := range stores {
		require.NoError(t, ds.Merge(store))
	}

	return ds.Seal()
}

func requireAddressRef(
	t *testing.T,
	ds datastore.DataStore,
	chainSelector uint64,
	contractType datastore.ContractType,
	version *semver.Version,
	qualifier string,
) datastore.AddressRef {
	t.Helper()

	filters := []datastore.FilterFunc[datastore.AddressRefKey, datastore.AddressRef]{
		datastore.AddressRefByChainSelector(chainSelector),
		datastore.AddressRefByType(contractType),
		datastore.AddressRefByQualifier(qualifier),
	}

	if version != nil {
		filters = append(filters, datastore.AddressRefByVersion(version))
	}

	refs := ds.Addresses().Filter(filters...)
	require.Len(
		t,
		refs,
		1,
		"expected exactly one %s ref on chain %d with qualifier %q",
		contractType,
		chainSelector,
		qualifier,
	)

	return refs[0]
}

func containsSelector(selectors []uint64, expected uint64) bool {
	for _, selector := range selectors {
		if selector == expected {
			return true
		}
	}
	return false
}
