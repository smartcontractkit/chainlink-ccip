package v1

import (
	"encoding/json"
	"fmt"

	"google.golang.org/protobuf/proto"

	"github.com/smartcontractkit/chainlink-ccip/execute/exectypes"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon/discovery/discoverytypes"
	"github.com/smartcontractkit/chainlink-ccip/pkg/ocrtypecodec/v1/ocrtypecodecpb"
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
	tr *protoTranslator
}

func NewExecCodecProto() *ExecCodecProto {
	return &ExecCodecProto{
		tr: newProtoTranslator(),
	}
}

func (e *ExecCodecProto) EncodeObservation(observation exectypes.Observation) ([]byte, error) {
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

	return proto.Marshal(pbObs)
}

func (e *ExecCodecProto) DecodeObservation(data []byte) (exectypes.Observation, error) {
	if len(data) == 0 {
		return exectypes.Observation{}, nil
	}

	pbObs := &ocrtypecodecpb.ExecObservation{}
	if err := proto.Unmarshal(data, pbObs); err != nil {
		return exectypes.Observation{}, fmt.Errorf("proto unmarshal ExecObservation: %w", err)
	}

	return exectypes.Observation{
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
	}, nil
}

func (e *ExecCodecProto) EncodeOutcome(outcome exectypes.Outcome) ([]byte, error) {
	outcome = exectypes.NewOutcome(outcome.State, outcome.CommitReports, outcome.Reports)

	pbObs := &ocrtypecodecpb.ExecOutcome{
		PluginState:          string(outcome.State),
		CommitReports:        e.tr.commitDataSliceToProto(outcome.CommitReports),
		ExecutePluginReports: e.tr.execPluginReportsToProto(outcome.Reports),
	}

	// If there is only one report, use the legacy field. This way new clients can still
	// form consensus with old ones.
	// TODO: Remove backwards compatibility code after a few releases.
	if len(pbObs.ExecutePluginReports) == 1 {
		pbObs.ExecutePluginReport = pbObs.ExecutePluginReports[0]
		pbObs.ExecutePluginReports = nil
	}

	return proto.MarshalOptions{Deterministic: true}.Marshal(pbObs)
}

func (e *ExecCodecProto) DecodeOutcome(data []byte) (exectypes.Outcome, error) {
	if len(data) == 0 {
		return exectypes.Outcome{}, nil
	}

	pbOutc := &ocrtypecodecpb.ExecOutcome{}
	if err := proto.Unmarshal(data, pbOutc); err != nil {
		return exectypes.Outcome{}, fmt.Errorf("proto unmarshal ExecOutcome: %w", err)
	}

	outc := exectypes.Outcome{
		State:         exectypes.PluginState(pbOutc.PluginState),
		CommitReports: e.tr.commitDataSliceFromProto(pbOutc.CommitReports),
		Reports:       e.tr.execPluginReportsFromProto(pbOutc.ExecutePluginReports),
	}

	// Decode the legacy Report field into the new Reports field. This way the plugin layer doesn't
	// need to worry about type migration.
	// TODO: Remove temporary migration code after a few releases.
	if pbOutc.ExecutePluginReport != nil {
		outc.Reports = e.tr.execPluginReportsFromProto([]*ocrtypecodecpb.ExecutePluginReport{pbOutc.ExecutePluginReport})
	}

	return outc, nil
}

// ExecCodecJSON is an implementation of ExecCodec that uses JSON.
// DEPRECATED: Use ExecCodecProto instead.
type ExecCodecJSON struct{}

// DEPRECATED
func NewExecCodecJSON() *ExecCodecJSON {
	return &ExecCodecJSON{}
}

func (*ExecCodecJSON) EncodeObservation(observation exectypes.Observation) ([]byte, error) {
	return json.Marshal(observation)
}

func (*ExecCodecJSON) DecodeObservation(data []byte) (exectypes.Observation, error) {
	if len(data) == 0 {
		return exectypes.Observation{}, nil
	}
	obs := exectypes.Observation{}
	err := json.Unmarshal(data, &obs)
	return obs, err
}

func (*ExecCodecJSON) EncodeOutcome(outcome exectypes.Outcome) ([]byte, error) {
	// We sort again here in case construction is not via the constructor.
	return json.Marshal(exectypes.NewOutcome(outcome.State, outcome.CommitReports, outcome.Reports))
}

func (*ExecCodecJSON) DecodeOutcome(data []byte) (exectypes.Outcome, error) {
	if len(data) == 0 {
		return exectypes.Outcome{}, nil
	}
	o := exectypes.Outcome{}
	err := json.Unmarshal(data, &o)
	return o, err
}
