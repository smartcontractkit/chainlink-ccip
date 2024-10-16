package evm

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

func Test_calculateMessageMaxGas(t *testing.T) {
	type args struct {
		dataLen   int
		numTokens int
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{
			name: "base",
			args: args{dataLen: 5, numTokens: 2},
			want: 822_264,
		},
		{
			name: "large",
			args: args{dataLen: 1000, numTokens: 1000},
			want: 346_477_520,
		},
		{
			name: "overheadGas test 1",
			args: args{dataLen: 0, numTokens: 0},
			want: 119_920,
		},
		{
			name: "overheadGas test 2",
			args: args{dataLen: len([]byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0}), numTokens: 1},
			want: 475_948,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			msg := ccipocr3.Message{
				Data:         make([]byte, tt.args.dataLen),
				TokenAmounts: make([]ccipocr3.RampTokenAmount, tt.args.numTokens),
			}
			ep := EstimateProvider{}
			got := ep.CalculateMessageMaxGas(msg)
			fmt.Println(got)
			assert.Equalf(t, tt.want, got, "calculateMessageMaxGas(%v, %v)", tt.args.dataLen, tt.args.numTokens)
		})
	}
}

// TestCalculateMaxGas is taken from the ccip repo where the CalculateMerkleTreeGas and CalculateMessageMaxGas values
// are combined to one function.
func TestCalculateMaxGas(t *testing.T) {
	tests := []struct {
		name           string
		numRequests    int
		dataLength     int
		numberOfTokens int
		want           uint64
	}{
		{
			name:           "maxGasOverheadGas 1",
			numRequests:    6,
			dataLength:     0,
			numberOfTokens: 0,
			want:           122992,
		},
		{
			name:           "maxGasOverheadGas 2",
			numRequests:    3,
			dataLength:     len([]byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0}),
			numberOfTokens: 1,
			want:           478508,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			msg := ccipocr3.Message{
				Data:         make([]byte, tt.dataLength),
				TokenAmounts: make([]ccipocr3.RampTokenAmount, tt.numberOfTokens),
			}
			ep := EstimateProvider{}

			gotTree := ep.CalculateMerkleTreeGas(tt.numRequests)
			gotMsg := ep.CalculateMessageMaxGas(msg)
			fmt.Println("want", tt.want, "got", gotTree+gotMsg)
			assert.Equal(t, tt.want, gotTree+gotMsg)
		})
	}
}
