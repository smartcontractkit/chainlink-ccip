package ocrtypecodec

import (
	"encoding/json"
	"fmt"

	"github.com/smartcontractkit/chainlink-ccip/commit/committypes"
)

// CommitCodec is an interface for encoding and decoding OCR related commit plugin types.
type CommitCodec interface {
	EncodeQuery(query committypes.Query) ([]byte, error)
	DecodeQuery(data []byte) (committypes.Query, error)

	EncodeObservation(observation committypes.Observation) ([]byte, error)
	DecodeObservation(data []byte) (committypes.Observation, error)

	EncodeOutcome(outcome committypes.Outcome) ([]byte, error)
	DecodeOutcome(data []byte) (committypes.Outcome, error)
}

// CommitCodecJSON is an implementation of CommitCodec that uses JSON.
type CommitCodecJSON struct{}

// NewCommitCodecJSON returns a new CommitCodecJSON.
func NewCommitCodecJSON() *CommitCodecJSON {
	return &CommitCodecJSON{}
}

func (*CommitCodecJSON) EncodeQuery(query committypes.Query) ([]byte, error) {
	return json.Marshal(query)
}

func (*CommitCodecJSON) DecodeQuery(data []byte) (committypes.Query, error) {
	if len(data) == 0 {
		return committypes.Query{}, nil
	}
	q := committypes.Query{}
	err := json.Unmarshal(data, &q)
	return q, err
}

func (*CommitCodecJSON) EncodeObservation(observation committypes.Observation) ([]byte, error) {
	encodedObservation, err := json.Marshal(observation)
	if err != nil {
		return nil, fmt.Errorf("failed to encode Observation: %w", err)
	}
	return encodedObservation, nil
}

func (*CommitCodecJSON) DecodeObservation(data []byte) (committypes.Observation, error) {
	if len(data) == 0 {
		return committypes.Observation{}, nil
	}
	o := committypes.Observation{}
	err := json.Unmarshal(data, &o)
	return o, err
}

func (*CommitCodecJSON) EncodeOutcome(outcome committypes.Outcome) ([]byte, error) {
	// Sort all lists to ensure deterministic serialization
	outcome.MerkleRootOutcome.Sort()
	encodedOutcome, err := json.Marshal(outcome)
	if err != nil {
		return nil, fmt.Errorf("failed to encode Outcome: %w", err)
	}

	return encodedOutcome, nil
}

func (*CommitCodecJSON) DecodeOutcome(data []byte) (committypes.Outcome, error) {
	if len(data) == 0 {
		return committypes.Outcome{}, nil
	}

	o := committypes.Outcome{}
	if err := json.Unmarshal(data, &o); err != nil {
		return committypes.Outcome{}, fmt.Errorf("decode outcome: %w", err)
	}

	return o, nil
}
