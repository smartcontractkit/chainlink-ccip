package v1

import (
	"fmt"
	"testing"

	"github.com/smartcontractkit/chainlink-ccip/commit/committypes"
)

func TestCommitQuery(t *testing.T) {
	jsonCodec := NewCommitCodecJSON()
	protoCodec := NewCommitCodecProto()

	results := make([]resultData, 0, len(dataGenerators))
	for _, gen := range dataGenerators {
		result := runBenchmark(
			t,
			gen.name,
			gen.commitQuery(),
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

	results := make([]resultData, 0, len(dataGenerators))
	for _, gen := range dataGenerators {
		result := runBenchmark(
			t,
			gen.name,
			gen.commitObservation(),
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

	results := make([]resultData, 0, len(dataGenerators))
	for _, gen := range dataGenerators {
		result := runBenchmark(
			t,
			gen.name,
			gen.commitOutcome(),
			func(b []byte) (interface{}, error) { return jsonCodec.DecodeOutcome(b) },
			func(i interface{}) ([]byte, error) { return jsonCodec.EncodeOutcome(i.(committypes.Outcome)) },
			func(b []byte) (interface{}, error) { return protoCodec.DecodeOutcome(b) },
			func(i interface{}) ([]byte, error) { return protoCodec.EncodeOutcome(i.(committypes.Outcome)) },
		)
		results = append(results, result)
	}
	fmt.Println(resultDataArray(results))
}
