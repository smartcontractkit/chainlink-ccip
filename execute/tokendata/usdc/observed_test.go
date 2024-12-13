package usdc

import (
	"context"
	"fmt"
	"testing"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/utils/tests"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/internal"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

func Test_ObservedMetrics(t *testing.T) {
	promAttestationDurations.Reset()

	successCount := 10

	ctx := tests.Context(t)
	client200 := newObservedHTTPClient(fake{status: 200}, logger.Test(t))
	client400 := newObservedHTTPClient(fake{status: 400}, logger.Test(t))
	client500 := newObservedHTTPClient(fake{status: 500, err: fmt.Errorf("error")}, logger.Test(t))

	for i := 0; i < successCount; i++ {
		_, st, _ := client200.Get(ctx, cciptypes.Bytes32{})
		require.Equal(t, 200, int(st))
	}

	_, st, _ := client400.Get(ctx, cciptypes.Bytes32{})
	require.Equal(t, 400, int(st))

	_, st, _ = client500.Get(ctx, cciptypes.Bytes32{})
	require.Equal(t, 500, int(st))

	c1 := internal.CounterFromHistogramByLabels(t, promAttestationDurations, "200")
	require.Equal(t, successCount, c1)

	c2 := internal.CounterFromHistogramByLabels(t, promAttestationDurations, "400")
	require.Equal(t, 1, c2)

	c3 := internal.CounterFromHistogramByLabels(t, promAttestationDurations, "500")
	require.Equal(t, 1, c3)

	c4 := internal.CounterFromHistogramByLabels(t, promAttestationDurations, "404")
	require.Equal(t, 0, c4)
}

type fake struct {
	status int
	err    error
}

func (f fake) Get(_ context.Context, messageHash cciptypes.Bytes32) (cciptypes.Bytes, HTTPStatus, error) {
	return messageHash[:], HTTPStatus(f.status), f.err
}
