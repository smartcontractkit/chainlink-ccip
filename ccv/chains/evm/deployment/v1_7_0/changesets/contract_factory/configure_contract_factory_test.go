package contract_factory_test

import (
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/changesets/contract_factory"
	contract_factory_ops "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/contract_factory"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	"github.com/stretchr/testify/require"
)

func TestConfigureContractFactory_Apply(t *testing.T) {
	chainSel := uint64(5009297550715157269)

	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{chainSel}),
	)
	require.NoError(t, err, "Failed to create test environment")
	require.NotNil(t, e, "Environment should be created")

	chain := e.BlockChains.EVMChains()[chainSel]
	initialAllowList := []common.Address{chain.DeployerKey.From}

	// Deploy ContractFactory first
	deployReport, err := cldf_ops.ExecuteOperation(
		e.OperationsBundle,
		contract_factory_ops.Deploy,
		chain,
		contract.DeployInput[contract_factory_ops.ConstructorArgs]{
			ChainSelector:  chainSel,
			TypeAndVersion: deployment.NewTypeAndVersion(contract_factory_ops.ContractType, *semver.MustParse("1.7.0")),
			Args: contract_factory_ops.ConstructorArgs{
				AllowList: initialAllowList,
			},
		},
	)
	require.NoError(t, err, "Failed to deploy ContractFactory")

	// Add the deployed contract to datastore
	ds := datastore.NewMemoryDataStore()
	err = ds.Addresses().Add(deployReport.Output)
	require.NoError(t, err, "Failed to add address to datastore")
	e.DataStore = ds.Seal()

	tests := []struct {
		desc             string
		allowListAdds    []common.Address
		allowListRemoves []common.Address
	}{
		{
			desc: "add single address to allow list",
			allowListAdds: []common.Address{
				common.HexToAddress("0x1234567890123456789012345678901234567890"),
			},
		},
		{
			desc: "add multiple addresses to allow list",
			allowListAdds: []common.Address{
				common.HexToAddress("0x1234567890123456789012345678901234567890"),
				common.HexToAddress("0x0987654321098765432109876543210987654321"),
			},
		},
		{
			desc: "remove single address from allow list",
			allowListRemoves: []common.Address{
				initialAllowList[0],
			},
		},
		{
			desc: "add and remove addresses",
			allowListAdds: []common.Address{
				common.HexToAddress("0x1111111111111111111111111111111111111111"),
			},
			allowListRemoves: []common.Address{
				common.HexToAddress("0x2222222222222222222222222222222222222222"),
			},
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			mcmsRegistry := changesets.NewMCMSReaderRegistry()
			// Apply the configuration changeset
			out, err := contract_factory.ConfigureContractFactory(mcmsRegistry).Apply(*e, changesets.WithMCMS[contract_factory.ConfigureContractFactoryInput[datastore.AddressRef]]{
				Cfg: contract_factory.ConfigureContractFactoryInput[datastore.AddressRef]{
					ChainSel: chainSel,
					ContractFactory: datastore.AddressRef{
						ChainSelector: chainSel,
						Type:          datastore.ContractType(contract_factory_ops.ContractType),
						Version:       semver.MustParse("1.7.0"),
					},
					AllowListAdds:    test.allowListAdds,
					AllowListRemoves: test.allowListRemoves,
				}})
			require.NoError(t, err, "Failed to apply ConfigureContractFactory changeset")
			require.NotNil(t, out.DataStore, "DataStore should be returned")
			require.Len(t, out.Reports, 2, "Should have two reports, one for sequence and one for nested op")
		})
	}
}

func TestConfigureContractFactory_ContractNotFound(t *testing.T) {
	chainSel := uint64(5009297550715157269)

	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{chainSel}),
	)
	require.NoError(t, err, "Failed to create test environment")
	require.NotNil(t, e, "Environment should be created")

	// Try to configure a contract that doesn't exist
	mcmsRegistry := changesets.NewMCMSReaderRegistry()
	_, err = contract_factory.ConfigureContractFactory(mcmsRegistry).Apply(*e, changesets.WithMCMS[contract_factory.ConfigureContractFactoryInput[datastore.AddressRef]]{
		Cfg: contract_factory.ConfigureContractFactoryInput[datastore.AddressRef]{
			ChainSel: chainSel,
			ContractFactory: datastore.AddressRef{
				ChainSelector: chainSel,
				Type:          datastore.ContractType(contract_factory_ops.ContractType),
				Version:       semver.MustParse("1.7.0"),
			},
			AllowListAdds: []common.Address{
				common.HexToAddress("0x1234567890123456789012345678901234567890"),
			},
		}})
	require.Error(t, err, "Should fail when contract is not in datastore")
}

func TestConfigureContractFactory_InvalidChain(t *testing.T) {
	chainSel := uint64(5009297550715157269)

	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{chainSel}),
	)
	require.NoError(t, err, "Failed to create test environment")
	require.NotNil(t, e, "Environment should be created")

	// Try to configure on a chain that doesn't exist
	mcmsRegistry := changesets.NewMCMSReaderRegistry()
	_, err = contract_factory.ConfigureContractFactory(mcmsRegistry).Apply(*e, changesets.WithMCMS[contract_factory.ConfigureContractFactoryInput[datastore.AddressRef]]{
		Cfg: contract_factory.ConfigureContractFactoryInput[datastore.AddressRef]{
			ChainSel: 99999999,
			ContractFactory: datastore.AddressRef{
				ChainSelector: 99999999,
				Type:          datastore.ContractType(contract_factory_ops.ContractType),
				Version:       semver.MustParse("1.7.0"),
			},
			AllowListAdds: []common.Address{
				common.HexToAddress("0x1234567890123456789012345678901234567890"),
			},
		}})
	require.Error(t, err, "Should fail when configuring on non-existent chain")
}

func TestConfigureContractFactory_EmptyUpdates(t *testing.T) {
	chainSel := uint64(5009297550715157269)

	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{chainSel}),
	)
	require.NoError(t, err, "Failed to create test environment")
	require.NotNil(t, e, "Environment should be created")

	chain := e.BlockChains.EVMChains()[chainSel]

	// Deploy ContractFactory first
	deployReport, err := cldf_ops.ExecuteOperation(
		e.OperationsBundle,
		contract_factory_ops.Deploy,
		chain,
		contract.DeployInput[contract_factory_ops.ConstructorArgs]{
			ChainSelector:  chainSel,
			TypeAndVersion: deployment.NewTypeAndVersion(contract_factory_ops.ContractType, *semver.MustParse("1.7.0")),
			Args: contract_factory_ops.ConstructorArgs{
				AllowList: []common.Address{chain.DeployerKey.From},
			},
		},
	)
	require.NoError(t, err, "Failed to deploy ContractFactory")

	// Add the deployed contract to datastore
	ds := datastore.NewMemoryDataStore()
	err = ds.Addresses().Add(deployReport.Output)
	require.NoError(t, err, "Failed to add address to datastore")
	e.DataStore = ds.Seal()

	// Apply configuration with empty adds and removes (should still work)
	mcmsRegistry := changesets.NewMCMSReaderRegistry()
	out, err := contract_factory.ConfigureContractFactory(mcmsRegistry).Apply(*e, changesets.WithMCMS[contract_factory.ConfigureContractFactoryInput[datastore.AddressRef]]{
		Cfg: contract_factory.ConfigureContractFactoryInput[datastore.AddressRef]{
			ChainSel: chainSel,
			ContractFactory: datastore.AddressRef{
				ChainSelector: chainSel,
				Type:          datastore.ContractType(contract_factory_ops.ContractType),
				Version:       semver.MustParse("1.7.0"),
			},
			AllowListAdds:    []common.Address{},
			AllowListRemoves: []common.Address{},
		}})
	require.NoError(t, err, "Should succeed even with empty updates")
	require.NotNil(t, out.DataStore, "DataStore should be returned")
}
