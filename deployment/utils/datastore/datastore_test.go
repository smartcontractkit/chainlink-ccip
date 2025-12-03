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
