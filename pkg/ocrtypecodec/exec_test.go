package ocrtypecodec

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"

	"github.com/smartcontractkit/chainlink-ccip/execute/exectypes"
)

func TestExecObservation(t *testing.T) {
	jsonCodec := NewExecCodecJSON()
	protoCodec := NewExecCodecProto()

	queriesVector := []struct {
		name                                                                                              string
		numCommitReports, numMessagesPerChain, numTokenDataPerChain, numNoncesPerChain, numCostlyMessages int
	}{
		{name: "empty observation"},
		{
			name:                 "small observation",
			numCommitReports:     2,
			numMessagesPerChain:  4,
			numTokenDataPerChain: 2,
			numNoncesPerChain:    2,
			numCostlyMessages:    1,
		},
		{
			name:                 "medium observation",
			numCommitReports:     16,
			numMessagesPerChain:  32,
			numTokenDataPerChain: 32,
			numNoncesPerChain:    16,
			numCostlyMessages:    8,
		},
		{
			name:                 "large observation",
			numCommitReports:     128,
			numMessagesPerChain:  200,
			numTokenDataPerChain: 128,
			numNoncesPerChain:    64,
			numCostlyMessages:    32,
		},
	}

	results := make([]resultData, 0, len(queriesVector))

	for _, qv := range queriesVector {
		pbQBytes, err := proto.Marshal(genExecObservation(
			qv.numCommitReports,
			qv.numMessagesPerChain,
			qv.numTokenDataPerChain,
			qv.numNoncesPerChain,
			qv.numCostlyMessages,
		))
		require.NoError(t, err)
		q, err := protoCodec.DecodeObservation(pbQBytes)
		require.NoError(t, err)

		result := runBenchmark(
			t,
			qv.name,
			q,
			func(b []byte) (interface{}, error) { return jsonCodec.DecodeObservation(b) },
			func(i interface{}) ([]byte, error) { return jsonCodec.EncodeObservation(i.(exectypes.Observation)) },
			func(b []byte) (interface{}, error) { return protoCodec.DecodeObservation(b) },
			func(i interface{}) ([]byte, error) { return protoCodec.EncodeObservation(i.(exectypes.Observation)) },
		)
		results = append(results, result)
	}

	fmt.Println(resultDataArray(results))
}
