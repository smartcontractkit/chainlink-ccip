package execute

import (
	"math"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_calculateMessageMaxGas(t *testing.T) {
	type args struct {
		gasLimit    *big.Int
		numRequests int
		dataLen     int
		numTokens   int
	}
	tests := []struct {
		name    string
		args    args
		want    uint64
		wantErr bool
	}{
		{
			name:    "base",
			args:    args{gasLimit: big.NewInt(1000), numRequests: 5, dataLen: 5, numTokens: 2},
			want:    826_336,
			wantErr: false,
		},
		{
			name:    "large",
			args:    args{gasLimit: big.NewInt(1000), numRequests: 1000, dataLen: 1000, numTokens: 1000},
			want:    346_485_176,
			wantErr: false,
		},
		{
			name:    "gas limit overflow",
			args:    args{gasLimit: big.NewInt(0).Mul(big.NewInt(math.MaxInt64), big.NewInt(math.MaxInt64))},
			want:    36_391_540,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := calculateMessageMaxGas(tt.args.gasLimit, tt.args.numRequests, tt.args.dataLen, tt.args.numTokens)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equalf(t, tt.want, got, "calculateMessageMaxGas(%v, %v, %v, %v)", tt.args.gasLimit, tt.args.numRequests, tt.args.dataLen, tt.args.numTokens)
		})
	}
}

func TestOverheadGas(t *testing.T) {
	// Only Data and TokenAmounts are used from the messages
	// And only the length is used so the contents doesn't matter.
	tests := []struct {
		dataLength     int
		numberOfTokens int
		want           uint64
	}{
		{
			dataLength:     0,
			numberOfTokens: 0,
			want:           119920,
		},
		{
			dataLength:     len([]byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0}),
			numberOfTokens: 1,
			want:           475948,
		},
	}

	for _, tc := range tests {
		got := overheadGas(tc.dataLength, tc.numberOfTokens)
		assert.Equalf(t, tc.want, got, "overheadGas(%v, %v)", tc.dataLength, tc.numberOfTokens)
	}
}

func TestMaxGasOverHeadGas(t *testing.T) {
	// Only Data and TokenAmounts are used from the messages
	// And only the length is used so the contents doesn't matter.
	tests := []struct {
		numMsgs        int
		dataLength     int
		numberOfTokens int
		want           uint64
	}{
		{
			numMsgs:        6,
			dataLength:     0,
			numberOfTokens: 0,
			want:           122992,
		},
		{
			numMsgs:        3,
			dataLength:     len([]byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0}),
			numberOfTokens: 1,
			want:           478508,
		},
	}

	for _, tc := range tests {
		got := maxGasOverHeadGas(tc.numMsgs, tc.dataLength, tc.numberOfTokens)
		assert.Equalf(t, tc.want, got, "maxGasOverHeadGas(%v, %v, %v)", tc.numMsgs, tc.dataLength, tc.numberOfTokens)
	}
}
