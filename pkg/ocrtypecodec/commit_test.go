package ocrtypecodec

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"

	"github.com/smartcontractkit/chainlink-ccip/commit/committypes"
)

type resultData struct {
	name                    string
	jsonEncodingTime        time.Duration
	protoEncodingTime       time.Duration
	jsonDecodingTime        time.Duration
	protoDecodingTime       time.Duration
	jsonEncodingDataLength  int
	protoEncodingDataLength int
}

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

func runBenchmark(
	t *testing.T,
	name string,
	obj interface{},
	decodeJsonFunc func([]byte) (interface{}, error),
	encodeJsonFunc func(interface{}) ([]byte, error),
	decodeProtoFunc func([]byte) (interface{}, error),
	encodeProtoFunc func(interface{}) ([]byte, error),
) resultData {
	result := resultData{name: name}

	tStart := time.Now()
	jsonEnc, err := encodeJsonFunc(obj)
	require.NoError(t, err)
	result.jsonEncodingTime = time.Since(tStart)
	tStart = time.Now()
	jsonDec, err := decodeJsonFunc(jsonEnc)
	result.jsonDecodingTime = time.Since(tStart)
	require.NoError(t, err)
	result.jsonEncodingDataLength = len(jsonEnc)

	tStart = time.Now()
	protoEnc, err := encodeProtoFunc(obj)
	require.NoError(t, err)
	result.protoEncodingTime = time.Since(tStart)
	tStart = time.Now()
	protoDec, err := decodeProtoFunc(protoEnc)
	result.protoDecodingTime = time.Since(tStart)
	require.NoError(t, err)
	result.protoEncodingDataLength = len(protoEnc)

	// sanity check
	require.Equal(t, jsonDec, protoDec)
	return result
}

// Helper functions for pretty-printing results

type resultDataArray []resultData

func (r resultDataArray) String() string {
	if len(r) == 0 {
		return "No results available"
	}

	// Table header
	header := []string{"Name", "JSON Enc", "Proto Enc", "JSON Dec", "Proto Dec", "JSON Size", "Proto Size"}
	columnWidths := []int{0, 20, 20, 20, 20, 12, 12}

	for _, entry := range r {
		if columnWidths[0] < len(entry.name) {
			columnWidths[0] = len(entry.name) + 1
		}
	}

	// Table separator
	separator := strings.Repeat("-", sum(columnWidths)+len(columnWidths)*3)

	// Format header row
	var sb strings.Builder
	sb.WriteString(separator + "\n")
	sb.WriteString(formatRow(header, columnWidths) + "\n")
	sb.WriteString(separator + "\n")

	// Format data rows
	for _, data := range r {
		row := []string{
			data.name,
			data.jsonEncodingTime.String(),
			data.protoEncodingTime.String(),
			data.jsonDecodingTime.String(),
			data.protoDecodingTime.String(),
			fmt.Sprintf("%d", data.jsonEncodingDataLength),
			fmt.Sprintf("%d", data.protoEncodingDataLength),
		}
		sb.WriteString(formatRow(row, columnWidths) + "\n")
	}

	sb.WriteString(separator)
	return sb.String()
}

// formatRow formats a row with padding for each column
func formatRow(fields []string, widths []int) string {
	var parts []string
	for i, field := range fields {
		parts = append(parts, fmt.Sprintf("%-*s", widths[i], field))
	}
	return "| " + strings.Join(parts, " | ") + " |"
}

// sum calculates the total width of all columns
func sum(arr []int) int {
	total := 0
	for _, v := range arr {
		total += v
	}
	return total
}
