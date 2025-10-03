package datastore_test

import (
	"testing"

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
