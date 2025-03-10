package v1

import (
	"encoding/json"
	"fmt"
	"sync"

	"google.golang.org/protobuf/proto"

	"github.com/smartcontractkit/chainlink-ccip/execute/exectypes"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon/discovery/discoverytypes"
	"github.com/smartcontractkit/chainlink-ccip/pkg/ocrtypecodec/v1/ocrtypecodecpb"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

var DefaultExecCodec ExecCodec = NewExecCodecProto()

// ExecCodec is an interface for encoding and decoding OCR related exec plugin types.
type ExecCodec interface {
	EncodeObservation(observation exectypes.Observation) ([]byte, error)
	DecodeObservation(data []byte) (exectypes.Observation, error)

	EncodeOutcome(outcome exectypes.Outcome) ([]byte, error)
	DecodeOutcome(data []byte) (exectypes.Outcome, error)
}

type ExecCodecProto struct {
	encodeObsMu     sync.Mutex
	decodeObsMu     sync.Mutex
	encodeOutcomeMu sync.Mutex
	decodeOutcomeMu sync.Mutex
	tr              *protoTranslator
}

func NewExecCodecProto() *ExecCodecProto {
	return &ExecCodecProto{
		encodeObsMu: sync.Mutex{},
		decodeObsMu: sync.Mutex{},
		tr:          newProtoTranslator(),
	}
}

func (e *ExecCodecProto) EncodeObservation(observation exectypes.Observation) ([]byte, error) {
	e.encodeObsMu.Lock()
	pbObs := &ocrtypecodecpb.ExecObservation{
		CommitReports: e.tr.commitReportsToProto(observation.CommitReports),
		SeqNumsToMsgs: e.tr.messageObservationsToProto(observation.Messages),
		MsgHashes:     e.tr.messageHashesToProto(observation.Hashes),
		TokenDataObservations: &ocrtypecodecpb.TokenDataObservations{
			TokenData: e.tr.tokenDataObservationsToProto(observation.TokenData),
		},
		Nonces: e.tr.nonceObservationsToProto(observation.Nonces),
		Contracts: &ocrtypecodecpb.DiscoveryObservation{
			FChain: e.tr.fChainToProto(observation.Contracts.FChain),
			ContractNames: &ocrtypecodecpb.ContractNameChainAddresses{
				Addresses: e.tr.discoveryAddressesToProto(observation.Contracts.Addresses),
			},
		},
		FChain: e.tr.fChainToProto(observation.FChain),
	}
	encoded, err := proto.Marshal(pbObs)
	e.encodeObsMu.Unlock()

	return encoded, err
}

func (e *ExecCodecProto) DecodeObservation(data []byte) (exectypes.Observation, error) {
	if len(data) == 0 {
		return exectypes.Observation{}, nil
	}

	pbObs := &ocrtypecodecpb.ExecObservation{}
	e.decodeObsMu.Lock()
	if err := proto.Unmarshal(data, pbObs); err != nil {
		return exectypes.Observation{}, fmt.Errorf("proto unmarshal ExecObservation: %w", err)
	}
	decoded := exectypes.Observation{
		CommitReports: e.tr.commitReportsFromProto(pbObs.CommitReports),
		Messages:      e.tr.messageObservationsFromProto(pbObs.SeqNumsToMsgs),
		Hashes:        e.tr.messageHashesFromProto(pbObs.MsgHashes),
		TokenData:     e.tr.tokenDataObservationsFromProto(pbObs.TokenDataObservations.TokenData),
		Nonces:        e.tr.nonceObservationsFromProto(pbObs.Nonces),
		Contracts: discoverytypes.Observation{
			FChain:    e.tr.fChainFromProto(pbObs.Contracts.FChain),
			Addresses: e.tr.discoveryAddressesFromProto(pbObs.Contracts.ContractNames.Addresses),
		},
		FChain: e.tr.fChainFromProto(pbObs.FChain),
	}
	e.decodeObsMu.Unlock()

	return decoded, nil
}

func (e *ExecCodecProto) EncodeOutcome(outcome exectypes.Outcome) ([]byte, error) {
	outcome = exectypes.NewSortedOutcome(outcome.State, outcome.CommitReports, outcome.Report)

	e.encodeOutcomeMu.Lock()
	pbObs := &ocrtypecodecpb.ExecOutcome{
		PluginState:   string(outcome.State),
		CommitReports: e.tr.commitDataSliceToProto(outcome.CommitReports),
		ExecutePluginReport: &ocrtypecodecpb.ExecutePluginReport{
			ChainReports: e.tr.chainReportsToProto(outcome.Report.ChainReports),
		},
	}
	encoded, err := proto.MarshalOptions{Deterministic: true}.Marshal(pbObs)
	e.encodeOutcomeMu.Unlock()

	return encoded, err
}

func (e *ExecCodecProto) DecodeOutcome(data []byte) (exectypes.Outcome, error) {
	if len(data) == 0 {
		return exectypes.Outcome{}, nil
	}

	e.decodeOutcomeMu.Lock()
	pbOutc := &ocrtypecodecpb.ExecOutcome{}
	if err := proto.Unmarshal(data, pbOutc); err != nil {
		return exectypes.Outcome{}, fmt.Errorf("proto unmarshal ExecOutcome: %w", err)
	}

	outc := exectypes.Outcome{
		State:         exectypes.PluginState(pbOutc.PluginState),
		CommitReports: e.tr.commitDataSliceFromProto(pbOutc.CommitReports),
		Report: cciptypes.ExecutePluginReport{
			ChainReports: e.tr.chainReportsFromProto(pbOutc.ExecutePluginReport.ChainReports),
		},
	}
	e.decodeOutcomeMu.Unlock()
	return outc, nil
}

// ExecCodecJSON is an implementation of ExecCodec that uses JSON.
// DEPRECATED: Use ExecCodecProto instead.
type ExecCodecJSON struct {
	mu sync.Mutex
}

// NewExecCodecJSON Used in tests only so far
// DEPRECATED
func NewExecCodecJSON() *ExecCodecJSON {
	return &ExecCodecJSON{
		mu: sync.Mutex{},
	}
}

func (e *ExecCodecJSON) EncodeObservation(observation exectypes.Observation) ([]byte, error) {
	e.mu.Lock()
	encoded, err := json.Marshal(observation)
	e.mu.Unlock()
	return encoded, err
}

func (e *ExecCodecJSON) DecodeObservation(data []byte) (exectypes.Observation, error) {
	if len(data) == 0 {
		return exectypes.Observation{}, nil
	}
	obs := exectypes.Observation{}
	e.mu.Lock()
	err := json.Unmarshal(data, &obs)
	e.mu.Unlock()
	return obs, err
}

func (e *ExecCodecJSON) EncodeOutcome(outcome exectypes.Outcome) ([]byte, error) {
	// We sort again here in case construction is not via the constructor.
	e.mu.Lock()
	encoded, err := json.Marshal(exectypes.NewSortedOutcome(outcome.State, outcome.CommitReports, outcome.Report))
	e.mu.Unlock()
	return encoded, err
}

func (e *ExecCodecJSON) DecodeOutcome(data []byte) (exectypes.Outcome, error) {
	if len(data) == 0 {
		return exectypes.Outcome{}, nil
	}
	o := exectypes.Outcome{}
	e.mu.Lock()
	err := json.Unmarshal(data, &o)
	e.mu.Unlock()
	return o, err
}
