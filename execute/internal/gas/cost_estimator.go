package gas

import (
	"github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
)

type MessageExecutionCostEstimator interface {
	EstimateMsgCostUSD(msg ccipocr3.Message) (ccipocr3.BigInt, error)
}

type StaticMessageExecutionCostEstimator struct {
	StaticCostUSD ccipocr3.BigInt
}

func NewStaticMessageExecutionCostEstimator(staticCostUSD ccipocr3.BigInt) *StaticMessageExecutionCostEstimator {
	return &StaticMessageExecutionCostEstimator{
		StaticCostUSD: staticCostUSD,
	}
}

func (s StaticMessageExecutionCostEstimator) EstimateMsgCostUSD(msg ccipocr3.Message) (ccipocr3.BigInt, error) {
	return s.StaticCostUSD, nil
}

// Ensure StaticMessageExecutionCostEstimator implements MessageExecutionCostEstimator
var _ MessageExecutionCostEstimator = (*StaticMessageExecutionCostEstimator)(nil)
