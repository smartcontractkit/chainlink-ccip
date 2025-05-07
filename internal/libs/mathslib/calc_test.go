package mathslib

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"

	"github.com/stretchr/testify/assert"

	sel "github.com/smartcontractkit/chain-selectors"
)

var (
	SolChainSelector = ccipocr3.ChainSelector(sel.SOLANA_DEVNET.Selector)
	EvmChainSelector = ccipocr3.ChainSelector(sel.ETHEREUM_TESTNET_SEPOLIA.Selector)
)

func TestDeviates(t *testing.T) {
	type args struct {
		x1  *big.Int
		x2  *big.Int
		ppb int64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "base case",
			args: args{x1: big.NewInt(1e9), x2: big.NewInt(2e9), ppb: 1},
			want: true,
		},
		{
			name: "x1 is zero and x1 neq x2",
			args: args{x1: big.NewInt(0), x2: big.NewInt(1), ppb: 999},
			want: true,
		},
		{
			name: "x2 is zero and x1 neq x2",
			args: args{x1: big.NewInt(1), x2: big.NewInt(0), ppb: 999},
			want: true,
		},
		{
			name: "x1 and x2 are both zero",
			args: args{x1: big.NewInt(0), x2: big.NewInt(0), ppb: 999},
			want: false,
		},
		{
			name: "deviates when ppb is 0",
			args: args{x1: big.NewInt(0), x2: big.NewInt(1), ppb: 0},
			want: true,
		},
		{
			name: "does not deviate when x1 eq x2",
			args: args{x1: big.NewInt(5), x2: big.NewInt(5), ppb: 1},
			want: false,
		},
		{
			name: "does not deviate with high ppb when x2 is greater",
			args: args{x1: big.NewInt(5), x2: big.NewInt(10), ppb: 2e9},
			want: false,
		},
		{
			name: "does not deviate with high ppb when x1 is greater",
			args: args{x1: big.NewInt(10), x2: big.NewInt(5), ppb: 1e9},
			want: false,
		},
		{
			name: "deviates with low ppb when x2 is greater",
			args: args{x1: big.NewInt(5), x2: big.NewInt(10), ppb: 9e8},
			want: true,
		},
		{
			name: "deviates with low ppb when x1 is greater",
			args: args{x1: big.NewInt(10), x2: big.NewInt(5), ppb: 9e8},
			want: true,
		},
		{
			name: "near deviation limit but deviates",
			args: args{x1: big.NewInt(10), x2: big.NewInt(5), ppb: 1e9 - 1},
			want: true,
		},
		{
			name: "near deviation limit but deviates",
			args: args{x1: big.NewInt(9e18), x2: big.NewInt(4e18), ppb: 1e9 - 1},
			want: true,
		},
		{
			name: "at deviation limit but does not deviate",
			args: args{x1: big.NewInt(10), x2: big.NewInt(5), ppb: 1e9},
			want: false,
		},
		{
			name: "near deviation limit but does not deviate",
			args: args{x1: big.NewInt(10), x2: big.NewInt(5), ppb: 1e9 + 1},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t,
				tt.want,
				Deviates(tt.args.x1, tt.args.x2, tt.args.ppb),
				"Deviates(%v, %v, %v)", tt.args.x1, tt.args.x2, tt.args.ppb,
			)
		})
	}
}

func TestCalculateUsdPerUnitGas(t *testing.T) {
	testCases := []struct {
		name           string
		sourceGasPrice *big.Int
		usdPerFeeCoin  *big.Int
		chainSelector  ccipocr3.ChainSelector
		exp            *big.Int
	}{
		{
			name:           "evm base case",
			sourceGasPrice: big.NewInt(1e9),
			usdPerFeeCoin:  new(big.Int).Mul(big.NewInt(2000), big.NewInt(1e18)), // $2000/Eth
			chainSelector:  EvmChainSelector,
			exp:            big.NewInt(2000e9),
		},
		{
			name:           "evm high fee case",
			sourceGasPrice: big.NewInt(2e18),
			usdPerFeeCoin:  new(big.Int).Mul(big.NewInt(4000), big.NewInt(1e18)),
			chainSelector:  EvmChainSelector,
			exp:            new(big.Int).Mul(big.NewInt(8000), big.NewInt(1e18)),
		},
		{
			name:           "evm low fee case",
			sourceGasPrice: big.NewInt(1e3),
			usdPerFeeCoin:  new(big.Int).Mul(big.NewInt(1000), big.NewInt(1e18)),
			chainSelector:  EvmChainSelector,
			exp:            big.NewInt(1000e3),
		},
		{
			name:           "sol base fee case",
			sourceGasPrice: big.NewInt(2000),
			usdPerFeeCoin:  new(big.Int).Mul(big.NewInt(150e9), big.NewInt(1e18)), // $150/Eth
			chainSelector:  SolChainSelector,
			exp:            big.NewInt(3e8),
		},
		{
			name:           "sol high fee case",
			sourceGasPrice: big.NewInt(400000),
			usdPerFeeCoin:  new(big.Int).Mul(big.NewInt(300e9), big.NewInt(1e18)),
			chainSelector:  SolChainSelector,
			exp:            big.NewInt(12e10),
		},
		{
			name:           "sol low fee case",
			sourceGasPrice: big.NewInt(10),
			usdPerFeeCoin:  new(big.Int).Mul(big.NewInt(1e9), big.NewInt(1e18)),
			chainSelector:  SolChainSelector,
			exp:            big.NewInt(10000),
		},
		{
			name:           "sol 0 fee case",
			sourceGasPrice: big.NewInt(0),
			usdPerFeeCoin:  new(big.Int).Mul(big.NewInt(150e9), big.NewInt(1e18)),
			chainSelector:  SolChainSelector,
			exp:            big.NewInt(0),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := CalculateUsdPerUnitGas(tc.chainSelector, tc.sourceGasPrice, tc.usdPerFeeCoin)
			t.Log(res.String())
			require.NoError(t, err)
			require.Zero(t, tc.exp.Cmp(res))
		})
	}
}

func MustBigIntSetString(s string, zeroSuffixSize int) *big.Int {
	// append zeroes to the string
	for i := 0; i < zeroSuffixSize; i++ {
		s += "0"
	}
	bi, ok := new(big.Int).SetString(s, 10)
	if !ok {
		panic("failed to parse big int")
	}
	return bi
}
