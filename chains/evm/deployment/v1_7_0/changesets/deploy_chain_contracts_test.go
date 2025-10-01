package changesets_test

import (
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/link"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/weth"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/changesets"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/sequences"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/testsetup"
	cldf_evm_provider "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/provider"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/stretchr/testify/require"
)

func TestDeployChainContracts_VerifyPreconditions(t *testing.T) {
	e, err := testsetup.CreateEnvironment(t, map[uint64]cldf_evm_provider.SimChainProviderConfig{
		5009297550715157269: {NumAdditionalAccounts: 1},
	})
	require.NoError(t, err, "Failed to create test environment")

	tests := []struct {
		desc        string
		input       changesets.DeployChainContractsCfg
		expectedErr string
	}{
		{
			desc: "valid input",
			input: changesets.DeployChainContractsCfg{
				ChainSel: 5009297550715157269,
				Params:   sequences.ContractParams{},
			},
		},
		{
			desc: "invalid chain selector",
			input: changesets.DeployChainContractsCfg{
				ChainSel: 12345,
				Params:   sequences.ContractParams{},
			},
			expectedErr: "no EVM chain with selector 12345 found in environment",
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			err := changesets.DeployChainContracts.VerifyPreconditions(e, test.input)
			if test.expectedErr != "" {
				require.ErrorContains(t, err, test.expectedErr, "Expected error containing %q but got none", test.expectedErr)
			} else {
				require.NoError(t, err, "Did not expect error but got: %v", err)
			}
		})
	}
}

func TestDeployChainContracts_Apply(t *testing.T) {
	tests := []struct {
		desc          string
		makeDatastore func() *datastore.MemoryDataStore
	}{
		{
			desc: "empty datastore",
			makeDatastore: func() *datastore.MemoryDataStore {
				return datastore.NewMemoryDataStore()
			},
		},
		{
			desc: "non-empty datastore",
			makeDatastore: func() *datastore.MemoryDataStore {
				ds := datastore.NewMemoryDataStore()
				_ = ds.Addresses().Add(datastore.AddressRef{
					ChainSelector: 5009297550715157269,
					Type:          datastore.ContractType(link.ContractType),
					Version:       semver.MustParse("1.0.0"),
					Address:       common.HexToAddress("0x01").Hex(),
				})
				_ = ds.Addresses().Add(datastore.AddressRef{
					ChainSelector: 5009297550715157269,
					Type:          datastore.ContractType(weth.ContractType),
					Version:       semver.MustParse("1.0.0"),
					Address:       common.HexToAddress("0x02").Hex(),
				})
				return ds
			},
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			e, err := testsetup.CreateEnvironment(t, map[uint64]cldf_evm_provider.SimChainProviderConfig{
				5009297550715157269: {NumAdditionalAccounts: 1},
			})
			require.NoError(t, err, "Failed to create test environment")

			ds := test.makeDatastore()
			existingAddrs, err := ds.Addresses().Fetch()
			require.NoError(t, err, "Failed to fetch addresses from datastore")
			e.DataStore = ds.Seal() // Override datastore in environment to include existing addresses

			out, err := changesets.DeployChainContracts.Apply(e, changesets.DeployChainContractsCfg{
				ChainSel: 5009297550715157269,
				Params:   testsetup.CreateBasicContractParams(),
			})
			require.NoError(t, err, "Failed to apply DeployChainContracts changeset")

			newAddrs, err := out.DataStore.Addresses().Fetch()
			require.NoError(t, err, "Failed to fetch addresses from datastore")

			for _, addr := range existingAddrs {
				for _, newAddr := range newAddrs {
					if addr.Type == newAddr.Type {
						require.Equal(t, addr.Address, newAddr.Address, "Expected existing address for type %s to remain unchanged", addr.Type)
					}
				}
			}
		})
	}
}
