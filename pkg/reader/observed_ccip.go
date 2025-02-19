package reader

import (
	"context"

	"github.com/prometheus/client_golang/prometheus"
	sel "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
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
		[]string{"chainID", "feeType"},
	)
)

type observedCCIPReader struct {
	CCIPReader
	lggr           logger.Logger
	destChain      cciptypes.ChainSelector
	chainFeesGauge *prometheus.GaugeVec
}

func NewObservedCCIPReader(
	reader CCIPReader,
	lggr logger.Logger,
	destChainSelector cciptypes.ChainSelector,
) CCIPReader {
	return &observedCCIPReader{
		CCIPReader:     reader,
		lggr:           lggr,
		destChain:      destChainSelector,
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
		selector, err := sel.GetChainIDFromSelector(uint64(k))
		if err != nil {
			o.lggr.Error("failed to get chainID from selector", "err", err)
			continue
		}

		if v.ExecutionFee != nil {
			o.chainFeesGauge.
				WithLabelValues(selector, execCostLabel).
				Set(float64(v.ExecutionFee.Int64()))
		}
		if v.DataAvailabilityFee != nil {
			o.chainFeesGauge.
				WithLabelValues(selector, dataCostLabel).
				Set(float64(v.DataAvailabilityFee.Int64()))
		}

		o.lggr.Debugw(
			"observed chain fee components",
			"destChainID", selector,
			"executionFee", v.ExecutionFee,
			"dataAvailabilityFee", v.DataAvailabilityFee,
		)
	}
}
