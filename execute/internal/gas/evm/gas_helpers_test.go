package evm

import (
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"

	"github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

const EvmExtraArgsV2Tag = "181dcf10"

func Test_calculateMessageMaxGas(t *testing.T) {
	type args struct {
		dataLen          int
		numTokens        int
		gasLimit         *big.Int
		tokenGasOverhead uint32
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{
			name: "base",
			args: args{dataLen: 5, numTokens: 2, gasLimit: big.NewInt(10), tokenGasOverhead: 10},
			want: 822_294,
		},
		{
			name: "large",
			args: args{dataLen: 1000, numTokens: 1000, gasLimit: big.NewInt(42000), tokenGasOverhead: 1},
			want: 346_520_520,
		},
		{
			name: "overheadGas test 1",
			args: args{dataLen: 0, numTokens: 0, gasLimit: big.NewInt(0), tokenGasOverhead: 100},
			want: 119_920,
		},
		{
			name: "overheadGas test 2",
			args: args{dataLen: len([]byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0}), numTokens: 1, gasLimit: big.NewInt(0),
				tokenGasOverhead: 2},
			want: 475_950,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lggr := logger.Test(t)
			msg := ccipocr3.Message{
				Data:         make([]byte, tt.args.dataLen),
				TokenAmounts: getTokenAmounts(t, tt.args.numTokens, tt.args.tokenGasOverhead),
				ExtraArgs:    getEncodedExtraArgs(t, tt.args.gasLimit),
			}
			ep := NewEstimateProvider(lggr)
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
		gasLimit       *big.Int
		want           uint64
	}{
		{
			name:           "maxGasOverheadGas 1",
			numRequests:    6,
			dataLength:     0,
			numberOfTokens: 0,
			gasLimit:       big.NewInt(250),
			want:           123242,
		},
		{
			name:           "maxGasOverheadGas 2",
			numRequests:    3,
			dataLength:     len([]byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0}),
			numberOfTokens: 1,
			gasLimit:       big.NewInt(250),
			want:           478758,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lggr := logger.Test(t)
			msg := ccipocr3.Message{
				Data:         make([]byte, tt.dataLength),
				TokenAmounts: make([]ccipocr3.RampTokenAmount, tt.numberOfTokens),
				ExtraArgs:    getEncodedExtraArgs(t, tt.gasLimit),
			}
			ep := NewEstimateProvider(lggr)

			gotTree := ep.CalculateMerkleTreeGas(tt.numRequests)
			gotMsg := ep.CalculateMessageMaxGas(msg)
			fmt.Println("want", tt.want, "got", gotTree+gotMsg)
			assert.Equal(t, tt.want, gotTree+gotMsg)
		})
	}
}

func getEncodedExtraArgs(t *testing.T, gasLimit *big.Int) []byte {
	encodedExtraArgs, err := ExtraArgsV2ABI.Pack(gasLimit, false)
	if err != nil {
		log.Fatalf("Failed to pack arguments: %v", err)
	}

	evmExtraArgsV2TagBytes, err := hex.DecodeString(EvmExtraArgsV2Tag)
	assert.NoError(t, err)

	encodedExtraArgs = append(evmExtraArgsV2TagBytes, encodedExtraArgs...)
	return encodedExtraArgs
}

func getTokenAmounts(t *testing.T, numTokens int, tokenGasOverhead uint32) []ccipocr3.RampTokenAmount {
	tokenDestGasOverhead, err := TokenDestGasOverheadABI.Pack(tokenGasOverhead)
	assert.NoError(t, err)

	tokenAmounts := make([]ccipocr3.RampTokenAmount, numTokens)
	for i := 0; i < numTokens; i++ {
		tokenAmounts[i] = ccipocr3.RampTokenAmount{
			DestExecData: tokenDestGasOverhead,
		}
	}
	return tokenAmounts
}

func TestDecodeEVMExtraArgsV2(t *testing.T) {
	abiEncodedEVMExtraArgsV2 := "181dcf100000000000000000000000000000000000000000000000000000000000001248" +
		"0000000000000000000000000000000000000000000000000000000000000001"

	data, err := hex.DecodeString(abiEncodedEVMExtraArgsV2)
	assert.NoError(t, err)

	gasLimit, err := decodeEVMExtraArgsV2GasLimit(data)
	assert.NoError(t, err)
	assert.Equal(t, big.NewInt(4680), gasLimit)
}

func TestDecodeTokenDestGasOverhead(t *testing.T) {
	abiEncodedTokenDestGasOverhead := "000000000000000000000000000000000000000000000000000000000003e6ea"

	data, err := hex.DecodeString(abiEncodedTokenDestGasOverhead)
	assert.NoError(t, err)

	tokenDestGasOverhead, err := decodeTokenDestGasOverhead(data)
	assert.NoError(t, err)
	assert.Equal(t, uint32(255722), tokenDestGasOverhead)
}
