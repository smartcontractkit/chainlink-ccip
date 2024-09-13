package shared

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
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
