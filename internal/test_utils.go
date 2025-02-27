package internal

import (
	"crypto/rand"
	"encoding/hex"
	"strings"
	"testing"

	"github.com/prometheus/client_golang/prometheus"
	io_prometheus_client "github.com/prometheus/client_model/go"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/mocks/pkg/types/ccipocr3"

	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

func MessageWithTokens(t *testing.T, tokenPoolAddr ...string) cciptypes.Message {
	onRampTokens := make([]cciptypes.RampTokenAmount, len(tokenPoolAddr))
	for i, addr := range tokenPoolAddr {
		b, err := cciptypes.NewUnknownAddressFromHex(addr)
		require.NoError(t, err)
		onRampTokens[i] = cciptypes.RampTokenAmount{
			SourcePoolAddress: b,
			Amount:            cciptypes.NewBigIntFromInt64(int64(i + 1)),
		}
	}
	return cciptypes.Message{
		TokenAmounts: onRampTokens,
	}
}

func RandBytes() cciptypes.Bytes {
	var array [32]byte
	_, err := rand.Read(array[:])
	if err != nil {
		panic(err)
	}
	return array[:]
}

func CounterFromHistogramByLabels(t *testing.T, histogramVec *prometheus.HistogramVec, labels ...string) int {
	observer, err := histogramVec.GetMetricWithLabelValues(labels...)
	require.NoError(t, err)

	metricCh := make(chan prometheus.Metric, 1)
	observer.(prometheus.Histogram).Collect(metricCh)
	close(metricCh)

	metric := <-metricCh
	pb := &io_prometheus_client.Metric{}
	err = metric.Write(pb)
	require.NoError(t, err)

	return int(pb.GetHistogram().GetSampleCount())
}

func NewMockAddressCodec(t *testing.T) *ccipocr3.MockAddressCodec {
	mockAddrCodec := ccipocr3.NewMockAddressCodec(t)
	mockAddrCodec.On("AddressBytesToString", mock.Anything, mock.Anything).
		Return(func(addr cciptypes.UnknownAddress, _ cciptypes.ChainSelector) string {
			return "0x" + hex.EncodeToString(addr)
		}, nil).Maybe()
	mockAddrCodec.On("AddressStringToBytes", mock.Anything, mock.Anything).
		Return(func(addr string, _ cciptypes.ChainSelector) (cciptypes.UnknownAddress, error) {
			addrBytes, err := hex.DecodeString(strings.ToLower(strings.TrimPrefix(addr, "0x")))
			if err != nil {
				return nil, err
			}
			return addrBytes, nil
		}).Maybe()
	return mockAddrCodec
}
