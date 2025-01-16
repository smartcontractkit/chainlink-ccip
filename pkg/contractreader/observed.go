package contractreader

import (
	"context"
	"strings"
	"time"

	"golang.org/x/exp/maps"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/types"
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

	CrDirectRequestsDurations = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "contract_reader_direct_request_duration",
			Help:    "The amount of time elapsed during the ChainReader's GetLatestValue execution",
			Buckets: buckets,
		},
		[]string{"chainID", "contract", "method"},
	)
	CrBatchRequestsDurations = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "contract_reader_batch_request_duration",
			Help:    "The amount of time elapsed during the ChainReader's BatchGetLatestValues execution",
			Buckets: buckets,
		},
		[]string{"chainID"},
	)
	CrBatchSizes = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "contract_reader_batch_request_size",
			Help: "The size of the batch request",
		},
		[]string{"chainID"},
	)
	CrErrors = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "contract_reader_errors",
			Help: "The number of errors that occurred during the ChainReader's execution",
		},
		[]string{"chainID", "function", "contract"},
	)
)

type Observed struct {
	ContractReaderFacade
	lggr    logger.Logger
	chainID string

	// Prometheus components for tracking metrics
	directRequestsDurations *prometheus.HistogramVec
	batchRequestsDurations  *prometheus.HistogramVec
	batchSizes              *prometheus.CounterVec
	errors                  *prometheus.CounterVec
}

func NewObserverReader(
	cr ContractReaderFacade,
	lggr logger.Logger,
	chainID string,
) *Observed {
	return &Observed{
		ContractReaderFacade:    cr,
		lggr:                    lggr,
		chainID:                 chainID,
		directRequestsDurations: CrDirectRequestsDurations,
		batchRequestsDurations:  CrBatchRequestsDurations,
		batchSizes:              CrBatchSizes,
		errors:                  CrErrors,
	}
}

func (o *Observed) HealthReport() map[string]error {
	// Health report doesn't seem to be an IO operation, so no need to observe.
	return o.ContractReaderFacade.HealthReport()
}

func (o *Observed) GetLatestValue(
	ctx context.Context,
	readIdentifier string,
	confidenceLevel primitives.ConfidenceLevel,
	params, returnVal any,
) error {
	start := time.Now()
	err := o.ContractReaderFacade.GetLatestValue(ctx, readIdentifier, confidenceLevel, params, returnVal)
	duration := time.Since(start)

	contract, function := unpackReadIdentifier(readIdentifier)
	o.directRequestsDurations.
		WithLabelValues(o.chainID, contract, function).
		Observe(float64(duration))

	o.maybeTrackErrors(err, "GetLatestValue", contract+"-"+function)
	o.lggr.Debugw("Observed GetLatestValue",
		"chainID", o.chainID,
		"contract", contract,
		"function", function,
		"readIdentifier", readIdentifier,
		"millis", duration.Milliseconds(),
	)
	return err
}

func (o *Observed) BatchGetLatestValues(
	ctx context.Context,
	request types.BatchGetLatestValuesRequest,
) (types.BatchGetLatestValuesResult, error) {
	start := time.Now()
	result, err := o.ContractReaderFacade.BatchGetLatestValues(ctx, request)
	duration := time.Since(start)

	o.batchRequestsDurations.
		WithLabelValues(o.chainID).
		Observe(float64(duration))

	o.batchSizes.
		WithLabelValues(o.chainID).
		Add(float64(len(request)))

	o.maybeTrackErrors(err, "BatchGetLatestValues", "")
	o.lggr.Debugw("Observed BatchGetLatestValues",
		"chainID", o.chainID,
		"millis", duration.Milliseconds(),
		"size", len(request),
		"contracts", maps.Keys(request),
	)
	return result, err
}

func (o *Observed) maybeTrackErrors(err error, function string, contract string) {
	if err == nil {
		return
	}
	o.errors.
		WithLabelValues(o.chainID, function, contract).
		Inc()
}

func unpackReadIdentifier(readIdentifier string) (contract string, function string) {
	split := strings.Split(readIdentifier, "-")
	if len(split) >= 3 {
		return split[1], split[2]
	}
	return readIdentifier, ""
}
