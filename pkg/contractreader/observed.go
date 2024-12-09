package contractreader

import (
	"context"
	"strings"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/smartcontractkit/chainlink-common/pkg/types"
	"github.com/smartcontractkit/chainlink-common/pkg/types/query"
	"github.com/smartcontractkit/chainlink-common/pkg/types/query/primitives"
)

var (
	buckets = []float64{
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
	}

	contractReaderDurations = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "contract_reader_duration",
			Help:    "The amount of time elapsed during the ChainReader's function execution",
			Buckets: buckets,
		},
		[]string{"chainID", "function", "contract"},
	)
)

type Observed struct {
	origin  ContractReaderFacade
	chainID string
}

func NewObserverReader(
	cr ContractReaderFacade,
	chainID string,
) *Observed {
	return &Observed{
		origin:  cr,
		chainID: chainID,
	}
}

func (o Observed) GetLatestValue(ctx context.Context, readIdentifier string, confidenceLevel primitives.ConfidenceLevel, params, returnVal any) error {
	start := time.Now()
	err := o.origin.GetLatestValue(ctx, readIdentifier, confidenceLevel, params, returnVal)
	contractReaderDurations.
		WithLabelValues(o.chainID, "GetLatestValue", unpackReadIdentifier(readIdentifier)).
		Observe(float64(time.Since(start)))
	return err
}

func (o Observed) BatchGetLatestValues(ctx context.Context, request types.BatchGetLatestValuesRequest) (types.BatchGetLatestValuesResult, error) {
	start := time.Now()
	result, err := o.origin.BatchGetLatestValues(ctx, request)
	contractReaderDurations.
		WithLabelValues(o.chainID, "BatchGetLatestValues", "").
		Observe(float64(time.Since(start)))
	return result, err
}

func (o Observed) Bind(ctx context.Context, bindings []types.BoundContract) error {
	return o.origin.Bind(ctx, bindings)
}

func (o Observed) Unbind(ctx context.Context, bindings []types.BoundContract) error {
	return o.origin.Unbind(ctx, bindings)
}

func (o Observed) QueryKey(ctx context.Context, contract types.BoundContract, filter query.KeyFilter, limitAndSort query.LimitAndSort, sequenceDataType any) ([]types.Sequence, error) {
	start := time.Now()
	result, err := o.origin.QueryKey(ctx, contract, filter, limitAndSort, sequenceDataType)
	contractReaderDurations.
		WithLabelValues(o.chainID, "QueryKey", contract.Name+"-"+filter.Key).
		Observe(float64(time.Since(start)))
	return result, err
}

func unpackReadIdentifier(readIdentifier string) string {
	split := strings.Split(readIdentifier, "-")
	if len(split) >= 3 {
		return split[1] + "-" + split[2]
	}
	return readIdentifier
}
