package reader

import (
	"context"
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/smartcontractkit/chainlink-common/pkg/types"

	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

const (
	execCostLabel = "execCost"
	dataCostLabel = "daCost"
)

var (
	PromChainFeeGauge = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "ccip_chain_fee_components",
			Help: "This metric tracks the chain fee components for a given chain",
		},
		[]string{"chainSelector", "feeType"},
	)
)

type observedCCIPReader struct {
	CCIPReader
	destChain      cciptypes.ChainSelector
	chainFeesGauge *prometheus.GaugeVec
}

func NewObservedCCIPReader(
	reader CCIPReader,
	destChain cciptypes.ChainSelector,
) CCIPReader {
	return &observedCCIPReader{
		CCIPReader:     reader,
		destChain:      destChain,
		chainFeesGauge: PromChainFeeGauge,
	}
}

func (o *observedCCIPReader) GetChainsFeeComponents(
	ctx context.Context,
	chains []cciptypes.ChainSelector,
) map[cciptypes.ChainSelector]types.ChainFeeComponents {
	res := o.CCIPReader.GetChainsFeeComponents(ctx, chains)
	o.trackChainFeeComponents(res)
	return res
}

func (o *observedCCIPReader) GetDestChainFeeComponents(
	ctx context.Context,
) (types.ChainFeeComponents, error) {
	res, err := o.CCIPReader.GetDestChainFeeComponents(ctx)
	if err == nil {
		o.trackChainFeeComponents(
			map[cciptypes.ChainSelector]types.ChainFeeComponents{o.destChain: res},
		)
	}
	return res, err
}

func (o *observedCCIPReader) trackChainFeeComponents(
	components map[cciptypes.ChainSelector]types.ChainFeeComponents,
) {
	for k, v := range components {
		stringSelector := strconv.FormatUint(uint64(k), 10)

		if v.ExecutionFee != nil {
			o.chainFeesGauge.
				WithLabelValues(stringSelector, execCostLabel).
				Set(float64(v.ExecutionFee.Int64()))
		}
		if v.DataAvailabilityFee != nil {
			o.chainFeesGauge.
				WithLabelValues(stringSelector, dataCostLabel).
				Set(float64(v.DataAvailabilityFee.Int64()))
		}
	}
}
