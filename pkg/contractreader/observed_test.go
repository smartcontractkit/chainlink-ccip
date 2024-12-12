package contractreader_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/prometheus/client_golang/prometheus/testutil"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/types"
	"github.com/smartcontractkit/chainlink-common/pkg/types/query/primitives"
	"github.com/smartcontractkit/chainlink-common/pkg/utils/tests"

	"github.com/smartcontractkit/chainlink-ccip/internal"
	mocked "github.com/smartcontractkit/chainlink-ccip/mocks/pkg/contractreader"
	"github.com/smartcontractkit/chainlink-ccip/pkg/contractreader"
)

func Test_GetBatchValue(t *testing.T) {
	t.Cleanup(resetMetrics)

	ctx := tests.Context(t)
	chainID := "solana"
	mockedReader := mocked.NewMockContractReaderFacade(t)

	r := make(types.BatchGetLatestValuesResult)

	request1 := types.BatchGetLatestValuesRequest{
		types.BoundContract{Address: "0xA", Name: "contract1"}: nil,
		types.BoundContract{Address: "0xB", Name: "contract2"}: nil,
	}
	request2 := types.BatchGetLatestValuesRequest{
		types.BoundContract{Address: "0xA", Name: "abc"}: nil,
		types.BoundContract{Address: "0xB", Name: "def"}: nil,
		types.BoundContract{Address: "0xC", Name: "ghi"}: nil,
	}
	request3 := types.BatchGetLatestValuesRequest{}

	mockedReader.EXPECT().BatchGetLatestValues(ctx, request1).
		Return(r, nil)
	mockedReader.EXPECT().BatchGetLatestValues(ctx, request2).
		Return(r, nil)
	mockedReader.EXPECT().BatchGetLatestValues(ctx, request3).
		Return(nil, fmt.Errorf("error"))

	reader := contractreader.NewObserverReader(mockedReader, logger.Test(t), chainID)

	_, err := reader.BatchGetLatestValues(ctx, request1)
	require.NoError(t, err)
	require.Equal(t, float64(2), testutil.ToFloat64(contractreader.CrBatchSizes.WithLabelValues(chainID)))

	_, err = reader.BatchGetLatestValues(ctx, request2)
	require.NoError(t, err)
	require.Equal(t, float64(5), testutil.ToFloat64(contractreader.CrBatchSizes.WithLabelValues(chainID)))

	require.Equal(t,
		float64(0),
		testutil.ToFloat64(contractreader.CrErrors.WithLabelValues(chainID, "BatchGetLatestValues", "")),
	)

	_, err = reader.BatchGetLatestValues(ctx, request3)
	require.Error(t, err)
	require.Equal(t,
		float64(1),
		testutil.ToFloat64(contractreader.CrErrors.WithLabelValues(chainID, "BatchGetLatestValues", "")),
	)
}

func Test_GetLatestValue(t *testing.T) {
	t.Cleanup(resetMetrics)

	ctx := tests.Context(t)
	chainID := "1"
	mockedReader := mocked.NewMockContractReaderFacade(t)

	contractID1 := "0x1-contract-read"
	contractID2 := "0x2-contract-faulty"

	mockedReader.EXPECT().GetLatestValue(ctx, contractID1, mock.Anything, mock.Anything, mock.Anything).
		Return(nil)
	mockedReader.EXPECT().GetLatestValue(ctx, contractID2, mock.Anything, mock.Anything, mock.Anything).
		Return(fmt.Errorf("error"))

	reader := contractreader.NewObserverReader(mockedReader, logger.Test(t), chainID)

	err := reader.GetLatestValue(ctx, contractID1, primitives.Unconfirmed, nil, nil)
	require.NoError(t, err)

	c1 := testutil.ToFloat64(contractreader.CrErrors.WithLabelValues(chainID, "contract", "read"))
	require.Equal(t, 0, int(c1))

	c2 := internal.CounterFromHistogramByLabels(t, contractreader.CrDirectRequestsDurations, chainID, "contract", "read")
	require.Equal(t, 1, c2)

	err = reader.GetLatestValue(ctx, contractID2, primitives.Unconfirmed, nil, nil)
	require.Error(t, err)

	c3 := testutil.ToFloat64(contractreader.CrErrors.WithLabelValues(chainID, "contract", "faulty"))
	require.Equal(t, 0, int(c3))
}

func resetMetrics() {
	contractreader.CrDirectRequestsDurations.Reset()
	contractreader.CrBatchRequestsDurations.Reset()
	contractreader.CrBatchSizes.Reset()
	contractreader.CrErrors.Reset()
}
