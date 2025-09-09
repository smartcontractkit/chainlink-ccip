package v1

import (
	"encoding/json"
	"fmt"

	"google.golang.org/protobuf/proto"

	"github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

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

	pbOtcm := &ocrtypecodecpb.ExecOutcome{
		PluginState:          string(outcome.State),
		CommitReports:        e.tr.commitDataSliceToProto(outcome.CommitReports),
		ExecutePluginReports: e.tr.execPluginReportsToProto(outcome.Reports),
	}

	// If there is only one report, use the legacy field. This way new clients can still
	// form consensus with old ones.
	// TODO: Remove backwards compatibility code after a few releases.
	if len(pbOtcm.ExecutePluginReports) == 1 {
		pbOtcm.ExecutePluginReport = pbOtcm.ExecutePluginReports[0]
		pbOtcm.ExecutePluginReports = nil
	}

	// TODO: Remove after "Reports" is fully supported.
	if len(outcome.Report.ChainReports) != 0 {
		r := e.tr.execPluginReportsToProto([]ccipocr3.ExecutePluginReport{outcome.Report})
		if len(r) > 0 {
			pbOtcm.ExecutePluginReport = r[0]
		}
	}

	return proto.MarshalOptions{Deterministic: true}.Marshal(pbOtcm)
}

func (e *ExecCodecProto) DecodeOutcome(data []byte) (exectypes.Outcome, error) {
	if len(data) == 0 {
		return exectypes.Outcome{}, nil
	}

	pbOtcm := &ocrtypecodecpb.ExecOutcome{}
	if err := proto.Unmarshal(data, pbOtcm); err != nil {
		return exectypes.Outcome{}, fmt.Errorf("proto unmarshal ExecOutcome: %w", err)
	}

	otcm := exectypes.Outcome{
		State:         exectypes.PluginState(pbOtcm.PluginState),
		CommitReports: e.tr.commitDataSliceFromProto(pbOtcm.CommitReports),
		Reports:       e.tr.execPluginReportsFromProto(pbOtcm.ExecutePluginReports),
	}

	// Decode the legacy Report field into the new Reports field. This way the plugin layer doesn't
	// need to worry about type migration.
	// TODO: Remove temporary migration code after a few releases.
	if pbOtcm.ExecutePluginReport != nil {
		otcm.Reports = e.tr.execPluginReportsFromProto([]*ocrtypecodecpb.ExecutePluginReport{pbOtcm.ExecutePluginReport})
	}

	// Decode the new report format into the legacy field as an intermediate step for implementing this feature.
	// TODO: Remove temporary function after the 'Reports' field is fully implemented.
	if len(otcm.Reports) > 0 {
		otcm.Report = otcm.Reports[0]
	}

	return otcm, nil
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
	// Create a new outcome with the same data for consistent representation

	// Backward compatibility handling (similar to Proto version)
	// Ensure Report field is populated if Reports has data
	if len(outcome.Reports) == 1 && len(outcome.Report.ChainReports) == 0 {
		outcome.Report = outcome.Reports[0]
	}

	// Ensure Reports contains Report if Report has data
	if len(outcome.Report.ChainReports) > 0 && (len(outcome.Reports) == 0) {
		outcome.Reports = append([]ccipocr3.ExecutePluginReport{outcome.Report}, outcome.Reports...)
	}

	return json.Marshal(outcome)
}

func (*ExecCodecJSON) DecodeOutcome(data []byte) (exectypes.Outcome, error) {
	if len(data) == 0 {
		return exectypes.Outcome{}, nil
	}
	o := exectypes.Outcome{}
	err := json.Unmarshal(data, &o)
	return o, err
}
