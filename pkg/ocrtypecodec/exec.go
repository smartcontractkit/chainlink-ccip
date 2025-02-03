package ocrtypecodec

import (
	"github.com/smartcontractkit/chainlink-ccip/execute/exectypes"
)

// ExecCodec is an interface for encoding and decoding OCR related exec plugin types.
type ExecCodec interface {
	EncodeObservation(observation exectypes.Observation) ([]byte, error)
	DecodeObservation(data []byte) (exectypes.Observation, error)

	EncodeOutcome(outcome exectypes.Outcome) ([]byte, error)
	DecodeOutcome(data []byte) (exectypes.Outcome, error)
}

type ExecCodecJSON struct{}

func NewExecCodecJSON() *ExecCodecJSON {
	return &ExecCodecJSON{}
}

func (*ExecCodecJSON) EncodeObservation(observation exectypes.Observation) ([]byte, error) {
	return observation.Encode()
}

func (*ExecCodecJSON) DecodeObservation(data []byte) (exectypes.Observation, error) {
	return exectypes.DecodeObservation(data)
}

func (*ExecCodecJSON) EncodeOutcome(outcome exectypes.Outcome) ([]byte, error) {
	return outcome.Encode()
}

func (*ExecCodecJSON) DecodeOutcome(data []byte) (exectypes.Outcome, error) {
	return exectypes.DecodeOutcome(data)
}
