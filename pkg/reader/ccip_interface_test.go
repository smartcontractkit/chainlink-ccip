package reader

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

func TestContractAddresses_Append(t *testing.T) {
	type args struct {
		contract string
		chain    ccipocr3.ChainSelector
		address  []byte
	}
	tests := []struct {
		name string
		ca   ContractAddresses
		args args
		want ContractAddresses
	}{
		{
			name: "append on nil",
			ca:   nil,
			args: args{
				contract: "foo",
				chain:    ccipocr3.ChainSelector(1),
				address:  []byte("bar"),
			},
			want: ContractAddresses{
				"foo": {
					ccipocr3.ChainSelector(1): []byte("bar"),
				},
			},
		},
		{
			name: "append on existing",
			ca: ContractAddresses{
				"foo": {
					ccipocr3.ChainSelector(1): []byte("bar"),
				},
			},
			args: args{
				contract: "foo",
				chain:    ccipocr3.ChainSelector(2),
				address:  []byte("baz"),
			},
			want: ContractAddresses{
				"foo": {

					ccipocr3.ChainSelector(1): []byte("bar"),
					ccipocr3.ChainSelector(2): []byte("baz"),
				},
			},
		},
		{
			name: "append on new contract",
			ca: ContractAddresses{
				"foo": {
					ccipocr3.ChainSelector(1): []byte("bar"),
				},
			},
			args: args{
				contract: "newContract",
				chain:    ccipocr3.ChainSelector(1),
				address:  []byte("newAddress"),
			},
			want: ContractAddresses{
				"foo": {
					ccipocr3.ChainSelector(1): []byte("bar"),
				},
				"newContract": {
					ccipocr3.ChainSelector(1): []byte("newAddress"),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t,
				tt.want,
				tt.ca.Append(tt.args.contract, tt.args.chain, tt.args.address),
				"Append(%v, %v, %v)", tt.args.contract, tt.args.chain, tt.args.address)
		})
	}
}
