package ocrtypecodec

import (
	"encoding/json"
	"math/big"

	"google.golang.org/protobuf/proto"

	"github.com/smartcontractkit/chainlink-ccip/execute/exectypes"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon/discovery/discoverytypes"
	"github.com/smartcontractkit/chainlink-ccip/pkg/ocrtypecodec/ocrtypecodecpb"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

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
		CommitReports:     e.tr.commitReportsToProto(observation.CommitReports),
		SeqNumsToMessages: e.tr.messageObservationsToProto(observation.Messages),
		MessageHashes:     e.tr.messageHashesToProto(observation.Hashes),
		TokenDataObservations: &ocrtypecodecpb.TokenDataObservations{
			TokenData: e.tr.tokenDataObservationsToProto(observation.TokenData),
		},
		CostlyMessages: e.tr.bytes32SliceToProto(observation.CostlyMessages),
		Nonces:         e.tr.nonceObservationsToProto(observation.Nonces),
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
	pbObs := &ocrtypecodecpb.ExecObservation{}
	if err := proto.Unmarshal(data, pbObs); err != nil {
		return exectypes.Observation{}, err
	}

	return exectypes.Observation{
		CommitReports:  e.tr.commitReportsFromProto(pbObs.CommitReports),
		Messages:       e.tr.messageObservationsFromProto(pbObs.SeqNumsToMessages),
		Hashes:         e.tr.messageHashesFromProto(pbObs.MessageHashes),
		TokenData:      e.tr.tokenDataObservationsFromProto(pbObs.TokenDataObservations.TokenData),
		CostlyMessages: e.tr.bytes32SliceFromProto(pbObs.CostlyMessages),
		Nonces:         e.tr.nonceObservationsFromProto(pbObs.Nonces),
		Contracts: discoverytypes.Observation{
			FChain:    e.tr.fChainFromProto(pbObs.Contracts.FChain),
			Addresses: e.tr.discoveryAddressesFromProto(pbObs.Contracts.ContractNames.Addresses),
		},
		FChain: e.tr.fChainFromProto(pbObs.FChain),
	}, nil
}

func decodeMessageTokenData(data []*ocrtypecodecpb.MessageTokenData) []exectypes.MessageTokenData {
	result := make([]exectypes.MessageTokenData, len(data))
	for i, item := range data {
		result[i] = decodeMessageTokenDataEntry(item)
	}
	return result
}

func decodeSeqNums(seqNums []uint64) []cciptypes.SeqNum {
	result := make([]cciptypes.SeqNum, len(seqNums))
	for i, num := range seqNums {
		result[i] = cciptypes.SeqNum(num)
	}
	return result
}

func decodeMessages(messages []*ocrtypecodecpb.Message) []cciptypes.Message {
	result := make([]cciptypes.Message, len(messages))
	for i, msg := range messages {
		result[i] = decodeMessage(msg)
	}
	return result
}

func decodeMessage(msg *ocrtypecodecpb.Message) cciptypes.Message {
	return cciptypes.Message{
		Header:         decodeMessageHeader(msg.Header),
		Sender:         msg.Sender,
		Data:           msg.Data,
		Receiver:       msg.Receiver,
		ExtraArgs:      msg.ExtraArgs,
		FeeToken:       msg.FeeToken,
		FeeTokenAmount: cciptypes.NewBigInt(big.NewInt(0).SetBytes(msg.FeeTokenAmount)),
		FeeValueJuels:  cciptypes.NewBigInt(big.NewInt(0).SetBytes(msg.FeeValueJuels)),
		TokenAmounts:   decodeRampTokenAmounts(msg.TokenAmounts),
	}
}

func decodeMessageHeader(header *ocrtypecodecpb.RampMessageHeader) cciptypes.RampMessageHeader {
	return cciptypes.RampMessageHeader{
		MessageID:           cciptypes.Bytes32(header.MessageId),
		SourceChainSelector: cciptypes.ChainSelector(header.SourceChainSelector),
		DestChainSelector:   cciptypes.ChainSelector(header.DestChainSelector),
		SequenceNumber:      cciptypes.SeqNum(header.SequenceNumber),
		Nonce:               header.Nonce,
		MsgHash:             cciptypes.Bytes32(header.MsgHash),
		OnRamp:              header.OnRamp,
	}
}

func decodeRampTokenAmounts(tokenAmounts []*ocrtypecodecpb.RampTokenAmount) []cciptypes.RampTokenAmount {
	result := make([]cciptypes.RampTokenAmount, len(tokenAmounts))
	for i, token := range tokenAmounts {
		result[i] = cciptypes.RampTokenAmount{
			SourcePoolAddress: token.SourcePoolAddress,
			DestTokenAddress:  token.DestTokenAddress,
			ExtraData:         token.ExtraData,
			Amount:            cciptypes.NewBigInt(big.NewInt(0).SetBytes(token.Amount)),
			DestExecData:      token.DestExecData,
		}
	}
	return result
}

func decodeMessageTokenDataEntry(data *ocrtypecodecpb.MessageTokenData) exectypes.MessageTokenData {
	tokenData := make([]exectypes.TokenData, len(data.TokenData))
	for i, td := range data.TokenData {
		tokenData[i] = exectypes.TokenData{
			Ready: td.Ready,
			Data:  td.Data,
		}
	}
	return exectypes.MessageTokenData{TokenData: tokenData}
}

func (e *ExecCodecProto) EncodeOutcome(outcome exectypes.Outcome) ([]byte, error) {
	return nil, nil
}

func (e *ExecCodecProto) DecodeOutcome(data []byte) (exectypes.Outcome, error) {
	return exectypes.Outcome{}, nil
}

type ExecCodecJSON struct{}

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
	return json.Marshal(exectypes.NewSortedOutcome(outcome.State, outcome.CommitReports, outcome.Report))
}

func (*ExecCodecJSON) DecodeOutcome(data []byte) (exectypes.Outcome, error) {
	if len(data) == 0 {
		return exectypes.Outcome{}, nil
	}
	o := exectypes.Outcome{}
	err := json.Unmarshal(data, &o)
	return o, err
}
