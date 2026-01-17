package reader

import (
	"context"
	"math/big"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/types"
	"github.com/smartcontractkit/chainlink-common/pkg/types/query/primitives"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-ccip/internal/libs"
)

const (
	execCostLabel = "execCost"
	dataCostLabel = "daCost"
)

var (
	PromChainFeeGauge = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "ccip_reader_chain_fee_components",
			Help: "This metric tracks the chain fee components for a given chain",
		},
		[]string{"chainFamily", "chainID", "feeType"},
	)
	PromQueryHistogram = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "ccip_reader_query_duration",
			Help: "This metric tracks the duration of queries made by the CCIP reader",
			Buckets: []float64{
				float64(10 * time.Millisecond),
				float64(20 * time.Millisecond),
				float64(50 * time.Millisecond),
				float64(70 * time.Millisecond),
				float64(100 * time.Millisecond),
				float64(200 * time.Millisecond),
				float64(500 * time.Millisecond),
				float64(700 * time.Millisecond),
				float64(time.Second),
				float64(2 * time.Second),
				float64(5 * time.Second),
			},
		},
		[]string{"chainFamily", "chainID", "query"},
	)
	PromDataSetSizeGauge = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "ccip_reader_data_set_size",
			Help: "This metric tracks the size of the data returned by the CCIP reader queries",
		},
		[]string{"chainFamily", "chainID", "query"},
	)
)

// observedCCIPReader is a wrapper around CCIPReader that tracks the duration of queries and the
// size of the data returned. It implements the CCIPReader interface and is used to observe the performance
// of the CCIPReader implementation. Every CCIPReader (and Chain Accessor Layer in the future) should be
// decorated with the observedCCIPReader to track the performance of the queries made by the underlying reader.
// Overhead is minimal, as it only measures performance of queries without adding any additional logic or
// complexity to the reader itself.
type observedCCIPReader struct {
	CCIPReader
	lggr              logger.Logger
	destChain         string
	destChainFamily   string
	destChainSelector cciptypes.ChainSelector

	queryHistogram   *prometheus.HistogramVec
	dataSetSizeGauge *prometheus.GaugeVec
	chainFeesGauge   *prometheus.GaugeVec
}

func NewObservedCCIPReader(
	reader CCIPReader,
	lggr logger.Logger,
	destChainSelector cciptypes.ChainSelector,
) CCIPReader {
	chainFamily, chainID, _ := libs.GetChainInfoFromSelector(destChainSelector)

	return &observedCCIPReader{
		CCIPReader: reader,
		lggr:       lggr,

		destChain:         chainID,
		destChainFamily:   chainFamily,
		destChainSelector: destChainSelector,

		queryHistogram:   PromQueryHistogram,
		dataSetSizeGauge: PromDataSetSizeGauge,
		chainFeesGauge:   PromChainFeeGauge,
	}
}

func (o *observedCCIPReader) CommitReportsGTETimestamp(
	ctx context.Context,
	ts time.Time,
	confidence primitives.ConfidenceLevel,
	limit int,
) ([]cciptypes.CommitPluginReportWithMeta, error) {
	return withObservedQueryAndResult(
		o,
		"CommitReportsGTETimestamp",
		func() ([]cciptypes.CommitPluginReportWithMeta, error) {
			return o.CCIPReader.CommitReportsGTETimestamp(ctx, ts, confidence, limit)
		},
		sliceLength,
	)
}

func (o *observedCCIPReader) ExecutedMessages(
	ctx context.Context,
	rangesPerChain map[cciptypes.ChainSelector][]cciptypes.SeqNumRange,
	confidence primitives.ConfidenceLevel,
) (map[cciptypes.ChainSelector][]cciptypes.SeqNum, error) {
	return withObservedQueryAndResult(
		o,
		"ExecutedMessages",
		func() (map[cciptypes.ChainSelector][]cciptypes.SeqNum, error) {
			return o.CCIPReader.ExecutedMessages(ctx, rangesPerChain, confidence)
		},
		mapOfSliceLength,
	)
}

func (o *observedCCIPReader) MsgsBetweenSeqNums(
	ctx context.Context,
	chain cciptypes.ChainSelector,
	seqNumRange cciptypes.SeqNumRange,
) ([]cciptypes.Message, error) {
	return withObservedQueryAndResult(
		o,
		"MsgsBetweenSeqNums",
		func() ([]cciptypes.Message, error) {
			return o.CCIPReader.MsgsBetweenSeqNums(ctx, chain, seqNumRange)
		},
		sliceLength,
	)
}

func (o *observedCCIPReader) LatestMsgSeqNum(
	ctx context.Context,
	chain cciptypes.ChainSelector,
) (cciptypes.SeqNum, error) {
	return withObservedQueryAndResult(
		o,
		"LatestMsgSeqNum",
		func() (cciptypes.SeqNum, error) {
			return o.CCIPReader.LatestMsgSeqNum(ctx, chain)
		},
		nil,
	)
}

func (o *observedCCIPReader) NextSeqNum(
	ctx context.Context,
	chains []cciptypes.ChainSelector,
) (seqNum map[cciptypes.ChainSelector]cciptypes.SeqNum, err error) {
	return withObservedQueryAndResult(
		o,
		"NextSeqNum",
		func() (map[cciptypes.ChainSelector]cciptypes.SeqNum, error) {
			return o.CCIPReader.NextSeqNum(ctx, chains)
		},
		mapLength,
	)
}

func (o *observedCCIPReader) Nonces(
	ctx context.Context,
	addressesByChain map[cciptypes.ChainSelector][]string,
) (map[cciptypes.ChainSelector]map[string]uint64, error) {
	return withObservedQueryAndResult(
		o,
		"Nonces",
		func() (map[cciptypes.ChainSelector]map[string]uint64, error) {
			return o.CCIPReader.Nonces(ctx, addressesByChain)
		},
		mapOfMapLength,
	)
}

func (o *observedCCIPReader) GetChainsFeeComponents(
	ctx context.Context,
	chains []cciptypes.ChainSelector,
) map[cciptypes.ChainSelector]types.ChainFeeComponents {
	res, _ := withObservedQueryAndResult(
		o,
		"GetChainsFeeComponents",
		func() (map[cciptypes.ChainSelector]types.ChainFeeComponents, error) {
			r := o.CCIPReader.GetChainsFeeComponents(ctx, chains)
			return r, nil
		},
		nil,
	)

	o.trackChainFeeComponents(res)
	return res
}

func (o *observedCCIPReader) GetDestChainFeeComponents(
	ctx context.Context,
) (types.ChainFeeComponents, error) {
	res, err := withObservedQueryAndResult(
		o,
		"GetDestChainFeeComponents",
		func() (types.ChainFeeComponents, error) {
			return o.CCIPReader.GetDestChainFeeComponents(ctx)
		},
		nil,
	)

	if err == nil {
		o.trackChainFeeComponents(
			map[cciptypes.ChainSelector]types.ChainFeeComponents{o.destChainSelector: res},
		)
	}
	return res, err
}

func (o *observedCCIPReader) GetWrappedNativeTokenPriceUSD(
	ctx context.Context,
	selectors []cciptypes.ChainSelector,
) map[cciptypes.ChainSelector]cciptypes.BigInt {
	res, _ := withObservedQueryAndResult(
		o,
		"GetWrappedNativeTokenPriceUSD",
		func() (map[cciptypes.ChainSelector]cciptypes.BigInt, error) {
			res := o.CCIPReader.GetWrappedNativeTokenPriceUSD(ctx, selectors)
			return res, nil
		},
		mapLength,
	)
	return res
}

func (o *observedCCIPReader) GetChainFeePriceUpdate(
	ctx context.Context,
	selectors []cciptypes.ChainSelector,
) map[cciptypes.ChainSelector]cciptypes.TimestampedBig {
	res, _ := withObservedQueryAndResult(
		o,
		"GetChainFeePriceUpdate",
		func() (map[cciptypes.ChainSelector]cciptypes.TimestampedBig, error) {
			res := o.CCIPReader.GetChainFeePriceUpdate(ctx, selectors)
			return res, nil
		},
		mapLength,
	)
	return res
}

func (o *observedCCIPReader) GetRMNRemoteConfig(ctx context.Context) (cciptypes.RemoteConfig, error) {
	return withObservedQueryAndResult(
		o,
		"GetRMNRemoteConfig",
		func() (cciptypes.RemoteConfig, error) {
			return o.CCIPReader.GetRMNRemoteConfig(ctx)
		},
		nil,
	)
}

func (o *observedCCIPReader) GetRmnCurseInfo(ctx context.Context) (cciptypes.CurseInfo, error) {
	return withObservedQueryAndResult(
		o,
		"GetRmnCurseInfo",
		func() (cciptypes.CurseInfo, error) {
			return o.CCIPReader.GetRmnCurseInfo(ctx)
		},
		nil,
	)
}

func (o *observedCCIPReader) DiscoverContracts(
	ctx context.Context,
	supportedChains, allChains []cciptypes.ChainSelector,
) (ContractAddresses, error) {
	contractAddressesLength := func(addresses ContractAddresses) float64 {
		return mapOfMapLength(addresses)
	}

	return withObservedQueryAndResult(
		o,
		"DiscoverContracts",
		func() (ContractAddresses, error) {
			return o.CCIPReader.DiscoverContracts(ctx, supportedChains, allChains)
		},
		contractAddressesLength,
	)
}

func (o *observedCCIPReader) GetOffRampSourceChainsConfig(
	ctx context.Context,
	sourceChains []cciptypes.ChainSelector,
) (map[cciptypes.ChainSelector]StaticSourceChainConfig, error) {
	return withObservedQueryAndResult(
		o,
		"GetOffRampSourceChainsConfig",
		func() (map[cciptypes.ChainSelector]StaticSourceChainConfig, error) {
			return o.CCIPReader.GetOffRampSourceChainsConfig(ctx, sourceChains)
		},
		mapLength,
	)
}

func (o *observedCCIPReader) trackChainFeeComponents(
	components map[cciptypes.ChainSelector]types.ChainFeeComponents,
) {
	for k, v := range components {
		chainFamily, chainID, ok := libs.GetChainInfoFromSelector(k)
		if !ok {
			o.lggr.Error("failed to get chainID/chainFamily from selector", "selector", k)
			continue
		}

		if v.ExecutionFee != nil {
			execFeeFloat, _ := new(big.Float).SetInt(v.ExecutionFee).Float64()
			o.chainFeesGauge.
				WithLabelValues(chainFamily, chainID, execCostLabel).
				Set(execFeeFloat)
		}
		if v.DataAvailabilityFee != nil {
			dataFeeFloat, _ := new(big.Float).SetInt(v.DataAvailabilityFee).Float64()
			o.chainFeesGauge.
				WithLabelValues(chainFamily, chainID, dataCostLabel).
				Set(dataFeeFloat)
		}

		o.lggr.Debugw(
			"observed chain fee components",
			"destChainID", chainID,
			"destChainFamily", chainFamily,
			"executionFee", v.ExecutionFee,
			"dataAvailabilityFee", v.DataAvailabilityFee,
		)
	}
}

// withObservedQueryAndResult is a helper function that wraps a query function and tracks its duration and data size
// if dataSizeFn is not provided, it will not track the data size. Data set is also not tracked when error is returned.
// Latency of the query is always tracked, regardless of the error.
func withObservedQueryAndResult[T any](
	o *observedCCIPReader,
	queryName string,
	query func() (T, error),
	dataSizeFn func(T) float64,
) (T, error) {
	queryStarted := time.Now()
	results, err := query()
	duration := time.Since(queryStarted)

	o.queryHistogram.
		WithLabelValues(o.destChainFamily, o.destChain, queryName).
		Observe(float64(duration))

	if err == nil && dataSizeFn != nil {
		o.dataSetSizeGauge.
			WithLabelValues(o.destChainFamily, o.destChain, queryName).
			Set(dataSizeFn(results))
	}

	o.lggr.Debugw("observed CCIP Reader query",
		"duration", duration,
		"query", queryName,
		"error", err,
	)

	return results, err
}

func sliceLength[T any](slice []T) float64 {
	return float64(len(slice))
}

func mapLength[K comparable, V any](m map[K]V) float64 {
	return float64(len(m))
}

func mapOfSliceLength[K comparable, T any](mapOfSlice map[K][]T) float64 {
	totalLength := 0
	for _, slice := range mapOfSlice {
		totalLength += len(slice)
	}
	return float64(totalLength)
}

func mapOfMapLength[K comparable, V comparable, T any](mapOfMap map[K]map[V]T) float64 {
	totalLength := 0
	for _, innerMap := range mapOfMap {
		totalLength += len(innerMap)
	}
	return float64(totalLength)
}
