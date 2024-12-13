package usdc

import (
	"context"
	"fmt"
	"testing"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/utils/tests"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/internal"
	"github.com/smartcontractkit/chainlink-ccip/internal/libs/testhelpers/rand"
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

func Test_TimeToAttestation(t *testing.T) {
	ctx := tests.Context(t)
	http := &fake{status: 200}
	client := newObservedHTTPClient(http, logger.Test(t))

	t.Cleanup(func() {
		client.timeToAttestation.Reset()
	})

	message1 := rand.RandomBytes32()
	message2 := rand.RandomBytes32()
	message3 := rand.RandomBytes32()

	t.Run("succesfull response dones't publish anyting", func(t *testing.T) {
		client.timeToAttestation.Reset()

		_, st, _ := client.Get(ctx, message1)
		require.Equal(t, 200, int(st))
		require.Empty(t, client.timeToAttestationCache.Items())
	})

	t.Run("failed and then back to success published duration", func(t *testing.T) {
		client.timeToAttestation.Reset()

		http.status = 400
		_, st, _ := client.Get(ctx, message2)
		require.Equal(t, 400, int(st))

		http.status = 200
		_, st, _ = client.Get(ctx, message2)
		require.Equal(t, 200, int(st))
		require.Empty(t, client.timeToAttestationCache.Items())
		require.Equal(t, 1, internal.CounterFromHistogramByLabels(t, promTimeToAttestation))
	})

	t.Run("forever failed doesn't publish anything", func(t *testing.T) {
		client.timeToAttestation.Reset()

		http.status = 400
		_, st, _ := client.Get(ctx, message3)
		require.Equal(t, 400, int(st))

		http.status = 500
		_, st, _ = client.Get(ctx, message3)
		require.Equal(t, 500, int(st))

		require.Equal(t, 0, internal.CounterFromHistogramByLabels(t, promTimeToAttestation))

		// Different message
		http.status = 200
		_, st, _ = client.Get(ctx, message2)
		require.Equal(t, 200, int(st))

		require.Equal(t, 0, internal.CounterFromHistogramByLabels(t, promTimeToAttestation))
	})
}

type fake struct {
	status int
	err    error
}

func (f fake) Get(_ context.Context, messageHash cciptypes.Bytes32) (cciptypes.Bytes, HTTPStatus, error) {
	return messageHash[:], HTTPStatus(f.status), f.err
}
