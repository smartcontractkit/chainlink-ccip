package changesets_test

import (
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/link"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/weth"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/changesets"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/fee_quoter"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/offramp"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/sequences"
	cs_core "github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	"github.com/stretchr/testify/require"
)

func TestDeployChainContracts_VerifyPreconditions(t *testing.T) {
	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{5009297550715157269}),
	)
	require.NoError(t, err, "Failed to create test environment")
	require.NotNil(t, e, "Environment should be created")

	tests := []struct {
		desc        string
		input       cs_core.WithMCMS[changesets.DeployChainContractsCfg]
		expectedErr string
	}{
		{
			desc: "valid input",
			input: cs_core.WithMCMS[changesets.DeployChainContractsCfg]{
				MCMS: mcms.Input{},
				Cfg: changesets.DeployChainContractsCfg{
					ChainSel: 5009297550715157269,
					Params:   sequences.ContractParams{},
				},
			},
		},
		{
			desc: "invalid chain selector",
			input: cs_core.WithMCMS[changesets.DeployChainContractsCfg]{
				MCMS: mcms.Input{},
				Cfg: changesets.DeployChainContractsCfg{
					ChainSel: 12345,
					Params:   sequences.ContractParams{},
				},
			},
			expectedErr: "no EVM chain with selector 12345 found in environment",
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			mcmsRegistry := cs_core.NewMCMSReaderRegistry()
			err := changesets.DeployChainContracts(mcmsRegistry).VerifyPreconditions(*e, test.input)
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
			e, err := environment.New(t.Context(),
				environment.WithEVMSimulated(t, []uint64{5009297550715157269}),
			)
			require.NoError(t, err, "Failed to create test environment")
			require.NotNil(t, e, "Environment should be created")

			ds := test.makeDatastore()
			existingAddrs, err := ds.Addresses().Fetch()
			require.NoError(t, err, "Failed to fetch addresses from datastore")
			e.DataStore = ds.Seal() // Override datastore in environment to include existing addresses

			mcmsRegistry := cs_core.NewMCMSReaderRegistry()
			out, err := changesets.DeployChainContracts(mcmsRegistry).Apply(*e, cs_core.WithMCMS[changesets.DeployChainContractsCfg]{
				MCMS: mcms.Input{},
				Cfg: changesets.DeployChainContractsCfg{
					ChainSel: 5009297550715157269,
					Params: sequences.ContractParams{
						FeeQuoter: fee_quoter.DefaultFeeQuoterParams(),
						OffRamp:   offramp.DefaultOffRampParams(),
					},
				},
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
