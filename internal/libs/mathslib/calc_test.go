package mathslib

import (
	"fmt"
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

func TestDeviatesOnCurve(t *testing.T) {
	type args struct {
		xNew  *big.Int
		xOld  *big.Int
		noDev *big.Int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "base case deviates from increase",
			args: args{xNew: big.NewInt(4e14), xOld: big.NewInt(1e13), noDev: big.NewInt(3e13)},
			want: true,
		},
		{
			name: "base case deviates from decrease",
			args: args{xNew: big.NewInt(1e13), xOld: big.NewInt(4e15), noDev: big.NewInt(1)},
			want: true,
		},
		{
			name: "does not deviate when equal",
			args: args{xNew: big.NewInt(3e14), xOld: big.NewInt(3e14), noDev: big.NewInt(3e13)},
			want: false,
		},
		{
			name: "does not deviate with small difference when xNew is bigger",
			args: args{xNew: big.NewInt(3e14 + 1), xOld: big.NewInt(3e14), noDev: big.NewInt(3e13)},
			want: false,
		},
		{
			name: "does not deviate with small difference when xOld is bigger",
			args: args{xNew: big.NewInt(3e14), xOld: big.NewInt(3e14 + 1), noDev: big.NewInt(3e13)},
			want: false,
		},
		{
			name: "does not deviate when xNew is below noDeviationLowerBound",
			args: args{xNew: big.NewInt(2e13), xOld: big.NewInt(1e13), noDev: big.NewInt(3e13)},
			want: false,
		},
		// thresholdPPB = (10e11) / (xNew^0.665) * 1e7
		// diff = (xNew - xOld) / min(xNew, xOld) * 1e9
		// Deviates = abs(diff) > thresholdPPB
		{
			name: "xNew is just below deviation threshold and does deviate",
			args: args{
				xNew:  big.NewInt(3e13),
				xOld:  big.NewInt(2.519478222838e12),
				noDev: big.NewInt(1),
			},
			want: false,
		},
		{
			name: "xNew is just above deviation threshold and does deviate",
			args: args{
				xNew:  big.NewInt(3e13),
				xOld:  big.NewInt(2.519478222838e12 - 30),
				noDev: big.NewInt(1),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(
				t,
				tt.want,
				DeviatesOnCurve(tt.args.xNew, tt.args.xOld, tt.args.noDev),
				"DeviatesOnCurve(%v, %v, %v)", tt.args.xNew, tt.args.xOld, tt.args.noDev)
		})
	}
}

func TestCalculateCurveThresholdPPB(t *testing.T) {
	tests := []struct {
		x             float64
		ppbLowerBound int64
		ppbUpperBound int64
	}{
		{
			x:             500_000,
			ppbLowerBound: 1_000_000_000_000_000, // 100,000,000%
			ppbUpperBound: 3_000_000_000_000_000, // 300,000,000%
		},
		{
			x:             50_000_000,
			ppbLowerBound: 70_000_000_000_000, // 7,000,000%
			ppbUpperBound: 90_000_000_000_000, // 9,000,000%
		},
		{
			x:             350_000_000,
			ppbLowerBound: 20_000_000_000_000, // 2,000,000%
			ppbUpperBound: 30_000_000_000_000, // 3,000,000%
		},
		{
			x:             200_000_000_000,
			ppbLowerBound: 300_000_000_000, // 30,000%
			ppbUpperBound: 400_000_000_000, // 40,000%
		},
		{
			x:             6_000_000_000_000,
			ppbLowerBound: 30_000_000_000, // 3,000%
			ppbUpperBound: 40000000000,    // 4,000%
		},
		{
			x:             50_000_000_000_000,
			ppbLowerBound: 7_000_000_000, // 700%
			ppbUpperBound: 8_000_000_000, // 800%
		},
		{
			x:             225_000_000_000_000,
			ppbLowerBound: 2_000_000_000, // 200%
			ppbUpperBound: 3_000_000_000, // 300%
		},
		{
			x:             5_000_000_000_000_000,
			ppbLowerBound: 300_000_000, // 30%
			ppbUpperBound: 400_000_000, // 40%
		},
		{
			x:             70_000_000_000_000_000,
			ppbLowerBound: 60_000_000, // 6%
			ppbUpperBound: 70_000_000, // 7%
		},
		{
			x:             500_000_000_000_000_000,
			ppbLowerBound: 10_000_000, // 1%
			ppbUpperBound: 20_000_000, // 2%
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%f", tt.x), func(t *testing.T) {
			thresholdPPB := calculateCurveThresholdPPB(new(big.Float).SetFloat64(tt.x))
			assert.GreaterOrEqual(t, thresholdPPB, tt.ppbLowerBound)
			assert.LessOrEqual(t, thresholdPPB, tt.ppbUpperBound)
		})
	}
}
