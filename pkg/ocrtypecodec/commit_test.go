package ocrtypecodec

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"

	"github.com/smartcontractkit/chainlink-ccip/commit/committypes"
)

func TestCommitQuery(t *testing.T) {
	jsonCodec := NewCommitCodecJSON()
	protoCodec := NewCommitCodecProto()

	queriesVector := []struct {
		name         string
		numSigs      int
		numIntervals int
	}{
		{name: "empty query", numSigs: 0, numIntervals: 0},
		{name: "small query", numSigs: 2, numIntervals: 2},
		{name: "medium query", numSigs: 16, numIntervals: 40},
		{name: "large query", numSigs: 32, numIntervals: 500},
	}

	results := make([]resultData, 0, len(queriesVector))

	for _, qv := range queriesVector {
		pbQBytes, err := proto.Marshal(genQuery(qv.numSigs, qv.numIntervals))
		require.NoError(t, err)
		q, err := protoCodec.DecodeQuery(pbQBytes)
		require.NoError(t, err)

		result := runBenchmark(
			t,
			qv.name,
			q,
			func(b []byte) (interface{}, error) { return jsonCodec.DecodeQuery(b) },
			func(i interface{}) ([]byte, error) { return jsonCodec.EncodeQuery(i.(committypes.Query)) },
			func(b []byte) (interface{}, error) { return protoCodec.DecodeQuery(b) },
			func(i interface{}) ([]byte, error) { return protoCodec.EncodeQuery(i.(committypes.Query)) },
		)
		results = append(results, result)
	}

	fmt.Println(resultDataArray(results))
}

func TestCommitObservation(t *testing.T) {
	jsonCodec := NewCommitCodecJSON()
	protoCodec := NewCommitCodecProto()

	queriesVector := []struct {
		name                                        string
		numMerkleRoots, numSeqNumChains, numSigners int
		numTokenPrices, numFeeQuoterUpdates         int
		numFeeComponents                            int
		numContractNames                            int
	}{
		{
			name: "empty observation",
		},
		{
			name:                "small observation",
			numMerkleRoots:      3,
			numSeqNumChains:     3,
			numSigners:          8,
			numTokenPrices:      6,
			numFeeQuoterUpdates: 6,
			numFeeComponents:    6,
			numContractNames:    8,
		},
		{
			name:                "medium observation",
			numMerkleRoots:      16,
			numSeqNumChains:     16,
			numSigners:          16,
			numTokenPrices:      32,
			numFeeQuoterUpdates: 32,
			numFeeComponents:    16,
			numContractNames:    16,
		},
		{
			name:                "large observation",
			numMerkleRoots:      256 * 8,
			numSeqNumChains:     32,
			numSigners:          32,
			numTokenPrices:      64,
			numFeeQuoterUpdates: 32,
			numFeeComponents:    32,
			numContractNames:    32,
		},
	}

	results := make([]resultData, 0, len(queriesVector))

	for _, qv := range queriesVector {
		pbObsBytes, err := proto.Marshal(genObservation(
			qv.numMerkleRoots, qv.numSeqNumChains, qv.numSigners,
			qv.numTokenPrices, qv.numFeeQuoterUpdates, qv.numFeeComponents, qv.numContractNames))
		require.NoError(t, err)
		obs, err := protoCodec.DecodeObservation(pbObsBytes)
		require.NoError(t, err)

		result := runBenchmark(
			t,
			qv.name,
			obs,
			func(b []byte) (interface{}, error) { return jsonCodec.DecodeObservation(b) },
			func(i interface{}) ([]byte, error) { return jsonCodec.EncodeObservation(i.(committypes.Observation)) },
			func(b []byte) (interface{}, error) { return protoCodec.DecodeObservation(b) },
			func(i interface{}) ([]byte, error) { return protoCodec.EncodeObservation(i.(committypes.Observation)) },
		)
		results = append(results, result)
	}

	fmt.Println(resultDataArray(results))
}

func TestCommitOutcome(t *testing.T) {
	jsonCodec := NewCommitCodecJSON()
	protoCodec := NewCommitCodecProto()

	queriesVector := []struct {
		name                                                                           string
		numRanges, numRoots, numSigners, numSeqNumChains, numGasPrices, numTokenPrices int
	}{
		{name: "empty outcome"},
		{
			name:            "small outcome",
			numRanges:       2,
			numRoots:        2,
			numSigners:      4,
			numSeqNumChains: 3,
			numGasPrices:    2,
			numTokenPrices:  2,
		},
		{
			name:            "medium outcome",
			numRanges:       32,
			numRoots:        32,
			numSigners:      16,
			numSeqNumChains: 32,
			numGasPrices:    16,
			numTokenPrices:  32,
		},
		{
			name:            "large outcome",
			numRanges:       256,
			numRoots:        256,
			numSigners:      32,
			numSeqNumChains: 256,
			numGasPrices:    128,
			numTokenPrices:  64,
		},
	}

	results := make([]resultData, 0, len(queriesVector))

	for _, qv := range queriesVector {
		pbBytes, err := proto.Marshal(genOutcome(
			qv.numRanges, qv.numRoots, qv.numSigners, qv.numSeqNumChains, qv.numGasPrices, qv.numTokenPrices))
		require.NoError(t, err)
		q, err := protoCodec.DecodeOutcome(pbBytes)
		require.NoError(t, err)

		result := runBenchmark(
			t,
			qv.name,
			q,
			func(b []byte) (interface{}, error) { return jsonCodec.DecodeOutcome(b) },
			func(i interface{}) ([]byte, error) { return jsonCodec.EncodeOutcome(i.(committypes.Outcome)) },
			func(b []byte) (interface{}, error) { return protoCodec.DecodeOutcome(b) },
			func(i interface{}) ([]byte, error) { return protoCodec.EncodeOutcome(i.(committypes.Outcome)) },
		)
		results = append(results, result)
	}

	fmt.Println(resultDataArray(results))
}
