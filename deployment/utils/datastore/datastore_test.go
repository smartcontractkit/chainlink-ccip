package datastore_test

import (
	"testing"

	"github.com/Masterminds/semver/v3"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/stretchr/testify/require"
)

func TestFindAndFormatRef(t *testing.T) {
	tests := []struct {
		desc          string
		makeDatastore func() datastore.DataStore
		expectedErr   string
		ref           datastore.AddressRef
	}{
		{
			desc: "find one ref",
			makeDatastore: func() datastore.DataStore {
				ds := datastore.NewMemoryDataStore()
				err := ds.Addresses().Add(datastore.AddressRef{
					ChainSelector: 4340886533089894000,
					Address:       "0x01",
					Type:          datastore.ContractType("TestContract"),
					Version:       semver.MustParse("1.0.0"),
					Qualifier:     "For testing",
				})
				require.NoError(t, err)
				return ds.Seal()
			},
			ref: datastore.AddressRef{
				ChainSelector: 4340886533089894000,
				Address:       "0x01",
				Type:          datastore.ContractType("TestContract"),
				Version:       semver.MustParse("1.0.0"),
				Qualifier:     "For testing",
			},
		},
		{
			desc: "find two refs",
			makeDatastore: func() datastore.DataStore {
				ds := datastore.NewMemoryDataStore()
				err := ds.Addresses().Add(datastore.AddressRef{
					ChainSelector: 4340886533089894000,
					Address:       "0x01",
					Type:          datastore.ContractType("TestContract"),
					Version:       semver.MustParse("1.0.0"),
					Qualifier:     "For testing",
				})
				require.NoError(t, err)
				err = ds.Addresses().Add(datastore.AddressRef{
					ChainSelector: 4340886533089894000,
					Address:       "0x02",
					Type:          datastore.ContractType("TestContract"),
					Version:       semver.MustParse("1.0.0"),
					Qualifier:     "For production",
				})
				require.NoError(t, err)
				return ds.Seal()
			},
			ref: datastore.AddressRef{
				ChainSelector: 4340886533089894000,
				Type:          datastore.ContractType("TestContract"),
				Version:       semver.MustParse("1.0.0"),
			},
			expectedErr: "found 2",
		},
		{
			desc: "find no refs",
			makeDatastore: func() datastore.DataStore {
				return datastore.NewMemoryDataStore().Seal()
			},
			ref: datastore.AddressRef{
				ChainSelector: 4340886533089894000,
				Type:          datastore.ContractType("TestContract"),
				Version:       semver.MustParse("1.0.0"),
				Qualifier:     "For testing",
			},
			expectedErr: "found 0",
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			ds := test.makeDatastore()
			addr, err := datastore_utils.FindAndFormatRef(ds, test.ref, test.ref.ChainSelector, func(ref datastore.AddressRef) (string, error) {
				return ref.Address, nil
			})
			if test.expectedErr != "" {
				require.ErrorContains(t, err, test.expectedErr)
				return
			}
			require.Equal(t, test.ref.Address, addr)
		})
	}
}

func TestFindAndFormatFirstRef(t *testing.T) {
	// Setup datastore with multiple refs
	ds := datastore.NewMemoryDataStore()
	err := ds.Addresses().Add(datastore.AddressRef{
		ChainSelector: 111,
		Address:       "0x01",
		Type:          datastore.ContractType("TestContract1"),
		Qualifier:     "For testing",
	})
	require.NoError(t, err)
	err = ds.Addresses().Add(datastore.AddressRef{
		ChainSelector: 111,
		Address:       "0x02",
		Type:          datastore.ContractType("TestContract2"),
		Qualifier:     "For testing",
	})
	require.NoError(t, err)
	err = ds.Addresses().Add(datastore.AddressRef{
		ChainSelector: 222,
		Address:       "0x03",
		Type:          datastore.ContractType("TestContract3"),
		Qualifier:     "For testing",
	})
	require.NoError(t, err)

	tests := []struct {
		name            string
		chainSelector   uint64
		refs            []datastore.AddressRef
		expectedAddress string
		expectedErr     string
	}{
		{
			name:          "find first ref",
			chainSelector: 111,
			refs: []datastore.AddressRef{
				{
					Type: datastore.ContractType("TestContract1"),
				}, {
					Type: datastore.ContractType("TestContract2"),
				},
			},
			expectedAddress: "0x01",
			expectedErr:     "",
		}, {
			name:          "find second ref",
			chainSelector: 111,
			refs: []datastore.AddressRef{
				{
					Type: datastore.ContractType("TestContract3"),
				}, {
					Type: datastore.ContractType("TestContract2"),
				},
			},
			expectedAddress: "0x02",
			expectedErr:     "",
		}, {
			name:          "find no refs",
			chainSelector: 333,
			refs: []datastore.AddressRef{
				{
					Type: datastore.ContractType("TestContract1"),
				}, {
					Type: datastore.ContractType("TestContract2"),
				}, {
					Type: datastore.ContractType("TestContract3"),
				},
			},
			expectedAddress: "",
			expectedErr:     "any of the provided refs",
		}, {
			name:            "no refs provided",
			chainSelector:   111,
			refs:            []datastore.AddressRef{},
			expectedAddress: "",
			expectedErr:     "at least one address ref must be specified",
		}, {
			name:          "non-unique ref found",
			chainSelector: 111,
			refs: []datastore.AddressRef{
				{
					ChainSelector: 111, // Matches two entries in the datastore, should error out and not continue to the next ref
				}, {
					ChainSelector: 222,
					Type:          datastore.ContractType("TestContract3"),
				},
			},
			expectedAddress: "",
			expectedErr:     "found 2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			addr, err := datastore_utils.FindAndFormatFirstRef(ds.Seal(), tt.chainSelector, func(ref datastore.AddressRef) (string, error) { return ref.Address, nil }, tt.refs...)
			if tt.expectedErr != "" {
				require.ErrorContains(t, err, tt.expectedErr)
				return
			}
			require.Equal(t, tt.expectedAddress, addr)
		})
	}
}
