package datastore_test

import (
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/stretchr/testify/require"
)

func TestToEVMAddress(t *testing.T) {
	tests := []struct {
		desc        string
		ref         datastore.AddressRef
		expectedErr string
	}{
		{
			desc: "happy path",
			ref: datastore.AddressRef{
				Address: "0x0000000000000000000000000000000000000001",
			},
		},
		{
			desc:        "empty address",
			ref:         datastore.AddressRef{},
			expectedErr: "address is empty in ref",
		},
		{
			desc: "invalid address",
			ref: datastore.AddressRef{
				Address: "1234",
			},
			expectedErr: "address is not a valid hex address in ref",
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			addr, err := datastore_utils.ToEVMAddress(test.ref)
			if test.expectedErr != "" {
				require.ErrorContains(t, err, test.expectedErr)
				return
			}
			require.Equal(t, test.ref.Address, addr.Hex())
		})
	}
}

func TestToPaddedEVMAddress(t *testing.T) {
	tests := []struct {
		desc        string
		ref         datastore.AddressRef
		expectedErr string
	}{
		{
			desc: "happy path",
			ref: datastore.AddressRef{
				Address: "0x0000000000000000000000000000000000000001",
			},
		},
		{
			desc:        "empty address",
			ref:         datastore.AddressRef{},
			expectedErr: "address is empty in ref",
		},
		{
			desc: "invalid address",
			ref: datastore.AddressRef{
				Address: "1234",
			},
			expectedErr: "address is not a valid hex address in ref",
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			addr, err := datastore_utils.ToPaddedEVMAddress(test.ref)
			if test.expectedErr != "" {
				require.ErrorContains(t, err, test.expectedErr)
				return
			}
			require.Equal(t, common.LeftPadBytes(common.HexToAddress(test.ref.Address).Bytes(), 32), addr)
		})
	}
}

func TestFindAndFormatEachRef(t *testing.T) {
	tests := []struct {
		desc          string
		makeDatastore func() datastore.DataStore
		expectedErr   string
		refs          []datastore.AddressRef
	}{
		{
			desc: "find one ref",
			makeDatastore: func() datastore.DataStore {
				ds := datastore.NewMemoryDataStore()
				err := ds.Addresses().Add(datastore.AddressRef{
					ChainSelector: 4340886533089894000,
					Address:       common.HexToAddress("0x01").String(),
					Type:          datastore.ContractType("TestContract"),
					Version:       semver.MustParse("1.0.0"),
					Qualifier:     "For testing",
				})
				require.NoError(t, err)
				return ds.Seal()
			},
			refs: []datastore.AddressRef{
				{
					ChainSelector: 4340886533089894000,
					Address:       common.HexToAddress("0x01").String(),
					Type:          datastore.ContractType("TestContract"),
					Version:       semver.MustParse("1.0.0"),
					Qualifier:     "For testing",
				},
			},
		},
		{
			desc: "find two refs",
			makeDatastore: func() datastore.DataStore {
				ds := datastore.NewMemoryDataStore()
				err := ds.Addresses().Add(datastore.AddressRef{
					ChainSelector: 4340886533089894000,
					Address:       common.HexToAddress("0x01").String(),
					Type:          datastore.ContractType("TestContract"),
					Version:       semver.MustParse("1.0.0"),
					Qualifier:     "For testing",
				})
				require.NoError(t, err)
				err = ds.Addresses().Add(datastore.AddressRef{
					ChainSelector: 4340886533089894000,
					Address:       common.HexToAddress("0x02").String(),
					Type:          datastore.ContractType("TestContract"),
					Version:       semver.MustParse("1.0.0"),
					Qualifier:     "For production",
				})
				require.NoError(t, err)
				return ds.Seal()
			},
			refs: []datastore.AddressRef{
				{
					ChainSelector: 4340886533089894000,
					Type:          datastore.ContractType("TestContract"),
					Version:       semver.MustParse("1.0.0"),
				},
			},
			expectedErr: "found 2",
		},
		{
			desc: "find no refs",
			makeDatastore: func() datastore.DataStore {
				return datastore.NewMemoryDataStore().Seal()
			},
			refs: []datastore.AddressRef{
				{
					ChainSelector: 4340886533089894000,
					Type:          datastore.ContractType("TestContract"),
					Version:       semver.MustParse("1.0.0"),
					Qualifier:     "For testing",
				},
			},
			expectedErr: "found 0",
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			ds := test.makeDatastore()
			addrs, err := datastore_utils.FindAndFormatEachRef(ds, test.refs, datastore_utils.ToEVMAddress)
			if test.expectedErr != "" {
				require.ErrorContains(t, err, test.expectedErr)
				return
			}
			require.Len(t, addrs, len(test.refs))
			for i, ref := range test.refs {
				require.Equal(t, ref.Address, addrs[i].Hex())
			}
		})
	}
}
