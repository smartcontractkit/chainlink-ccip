package gas

import (
	"math/big"

	"github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
)

type EstimateProvider interface {
	CalculateMerkleTreeGas(numRequests int) uint64
	CalculateMessageMaxGas(msg ccipocr3.Message) uint64
	CalculateMessageMaxDAGas(
		msg ccipocr3.Message,
		destDAOverheadGas, destGasPerDAByte, destDAMultiplierBps int64) *big.Int
}
