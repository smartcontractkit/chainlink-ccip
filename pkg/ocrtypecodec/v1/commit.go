package v1

import (
	"encoding/json"
	"fmt"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/smartcontractkit/chainlink-ccip/commit/chainfee"
	"github.com/smartcontractkit/chainlink-ccip/commit/committypes"
	"github.com/smartcontractkit/chainlink-ccip/commit/merkleroot"
	"github.com/smartcontractkit/chainlink-ccip/commit/merkleroot/rmn"
	"github.com/smartcontractkit/chainlink-ccip/commit/tokenprice"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon/discovery/discoverytypes"
	"github.com/smartcontractkit/chainlink-ccip/pkg/ocrtypecodec/v1/ocrtypecodecpb"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

var DefaultCommitCodec CommitCodec = NewCommitCodecProto()

// CommitCodec is an interface for encoding and decoding OCR related commit plugin types.
type CommitCodec interface {
	EncodeQuery(query committypes.Query) ([]byte, error)
	DecodeQuery(data []byte) (committypes.Query, error)

	EncodeObservation(observation committypes.Observation) ([]byte, error)
	DecodeObservation(data []byte) (committypes.Observation, error)

	EncodeOutcome(outcome committypes.Outcome) ([]byte, error)
	DecodeOutcome(data []byte) (committypes.Outcome, error)
}

type CommitCodecProto struct {
	tr *protoTranslator
}

func NewCommitCodecProto() *CommitCodecProto {
	return &CommitCodecProto{
		tr: newProtoTranslator(),
	}
}

func (c *CommitCodecProto) EncodeQuery(query committypes.Query) ([]byte, error) {
	if query.MerkleRootQuery.RMNSignatures == nil {
		query.MerkleRootQuery.RMNSignatures = &rmn.ReportSignatures{}
	}

	pbQ := &ocrtypecodecpb.CommitQuery{
		MerkleRootQuery: &ocrtypecodecpb.MerkleRootQuery{
			RetryRmnSignatures: query.MerkleRootQuery.RetryRMNSignatures,
			RmnSignatures: &ocrtypecodecpb.ReportSignatures{
				Signatures:  c.tr.rmnSignaturesToProto(query.MerkleRootQuery.RMNSignatures),
				LaneUpdates: c.tr.laneUpdatesToProto(query.MerkleRootQuery.RMNSignatures.LaneUpdates),
			},
		},
	}

	return proto.Marshal(pbQ)
}

func (c *CommitCodecProto) DecodeQuery(data []byte) (committypes.Query, error) {
	if len(data) == 0 {
		return committypes.Query{}, nil
	}

	pbQ := &ocrtypecodecpb.CommitQuery{}
	if err := proto.Unmarshal(data, pbQ); err != nil {
		return committypes.Query{}, fmt.Errorf("proto unmarshal query: %w", err)
	}

	q := committypes.Query{
		MerkleRootQuery: merkleroot.Query{
			RetryRMNSignatures: pbQ.MerkleRootQuery.RetryRmnSignatures,
			RMNSignatures: &rmn.ReportSignatures{
				Signatures:  c.tr.rmnSignaturesFromProto(pbQ.MerkleRootQuery.RmnSignatures.Signatures),
				LaneUpdates: c.tr.laneUpdatesFromProto(pbQ.MerkleRootQuery.RmnSignatures.LaneUpdates),
			},
		},
		TokenPriceQuery: tokenprice.Query{},
		ChainFeeQuery:   chainfee.Query{},
	}

	return q, nil
}

func (c *CommitCodecProto) EncodeObservation(observation committypes.Observation) ([]byte, error) {
	pbObs := &ocrtypecodecpb.CommitObservation{
		MerkleRootObs: &ocrtypecodecpb.MerkleRootObservation{
			MerkleRoots:        c.tr.merkleRootsToProto(observation.MerkleRootObs.MerkleRoots),
			RmnEnabledChains:   c.tr.rmnEnabledChainsToProto(observation.MerkleRootObs.RMNEnabledChains),
			OnRampMaxSeqNums:   c.tr.seqNumChainToProto(observation.MerkleRootObs.OnRampMaxSeqNums),
			OffRampNextSeqNums: c.tr.seqNumChainToProto(observation.MerkleRootObs.OffRampNextSeqNums),
			RmnRemoteConfig:    c.tr.rmnRemoteConfigToProto(observation.MerkleRootObs.RMNRemoteConfig),
			FChain:             c.tr.fChainToProto(observation.MerkleRootObs.FChain),
		},
		TokenPriceObs: &ocrtypecodecpb.TokenPriceObservation{
			FeedTokenPrices:       c.tr.feedTokenPricesToProto(observation.TokenPriceObs.FeedTokenPrices),
			FeeQuoterTokenUpdates: c.tr.feeQuoterTokenUpdatesToProto(observation.TokenPriceObs.FeeQuoterTokenUpdates),
			FChain:                c.tr.fChainToProto(observation.TokenPriceObs.FChain),
			Timestamp:             timestamppb.New(observation.TokenPriceObs.Timestamp),
		},
		ChainFeeObs: &ocrtypecodecpb.ChainFeeObservation{
			FeeComponents:     c.tr.feeComponentsToProto(observation.ChainFeeObs.FeeComponents),
			NativeTokenPrices: c.tr.nativeTokenPricesToProto(observation.ChainFeeObs.NativeTokenPrices),
			ChainFeeUpdates:   c.tr.chainFeeUpdatesToProto(observation.ChainFeeObs.ChainFeeUpdates),
			FChain:            c.tr.fChainToProto(observation.ChainFeeObs.FChain),
			TimestampNow:      timestamppb.New(observation.ChainFeeObs.TimestampNow),
		},
		DiscoveryObs: &ocrtypecodecpb.DiscoveryObservation{
			FChain: c.tr.fChainToProto(observation.DiscoveryObs.FChain),
			ContractNames: &ocrtypecodecpb.ContractNameChainAddresses{
				Addresses: c.tr.discoveryAddressesToProto(observation.DiscoveryObs.Addresses),
			},
		},
		FChain: c.tr.fChainToProto(observation.FChain),
	}
	return proto.Marshal(pbObs)
}

func (c *CommitCodecProto) DecodeObservation(data []byte) (committypes.Observation, error) {
	if len(data) == 0 {
		return committypes.Observation{}, nil
	}

	pbObs := &ocrtypecodecpb.CommitObservation{}
	if err := proto.Unmarshal(data, pbObs); err != nil {
		return committypes.Observation{}, fmt.Errorf("proto unmarshal observation: %w", err)
	}

	return committypes.Observation{
		MerkleRootObs: merkleroot.Observation{
			MerkleRoots:        c.tr.merkleRootsFromProto(pbObs.MerkleRootObs.MerkleRoots),
			RMNEnabledChains:   c.tr.rmnEnabledChainsFromProto(pbObs.MerkleRootObs.RmnEnabledChains),
			OnRampMaxSeqNums:   c.tr.seqNumChainFromProto(pbObs.MerkleRootObs.OnRampMaxSeqNums),
			OffRampNextSeqNums: c.tr.seqNumChainFromProto(pbObs.MerkleRootObs.OffRampNextSeqNums),
			RMNRemoteConfig:    c.tr.rmnRemoteConfigFromProto(pbObs.MerkleRootObs.RmnRemoteConfig),
			FChain:             c.tr.fChainFromProto(pbObs.MerkleRootObs.FChain),
		},
		TokenPriceObs: tokenprice.Observation{
			FeedTokenPrices:       c.tr.feedTokenPricesFromProto(pbObs.TokenPriceObs.FeedTokenPrices),
			FeeQuoterTokenUpdates: c.tr.feeQuoterTokenUpdatesFromProto(pbObs.TokenPriceObs.FeeQuoterTokenUpdates),
			FChain:                c.tr.fChainFromProto(pbObs.TokenPriceObs.FChain),
			Timestamp:             pbObs.TokenPriceObs.Timestamp.AsTime(),
		},
		ChainFeeObs: chainfee.Observation{
			FeeComponents:     c.tr.feeComponentsFromProto(pbObs.ChainFeeObs.FeeComponents),
			NativeTokenPrices: c.tr.nativeTokenPricesFromProto(pbObs.ChainFeeObs.NativeTokenPrices),
			ChainFeeUpdates:   c.tr.chainFeeUpdatesFromProto(pbObs.ChainFeeObs.ChainFeeUpdates),
			FChain:            c.tr.fChainFromProto(pbObs.ChainFeeObs.FChain),
			TimestampNow:      pbObs.ChainFeeObs.TimestampNow.AsTime(),
		},
		DiscoveryObs: discoverytypes.Observation{
			FChain:    c.tr.fChainFromProto(pbObs.DiscoveryObs.FChain),
			Addresses: c.tr.discoveryAddressesFromProto(pbObs.DiscoveryObs.ContractNames.Addresses),
		},
		FChain: c.tr.fChainFromProto(pbObs.FChain),
	}, nil
}

func (c *CommitCodecProto) EncodeOutcome(outcome committypes.Outcome) ([]byte, error) {
	pbOutcome := &ocrtypecodecpb.CommitOutcome{
		MerkleRootOutcome: &ocrtypecodecpb.MerkleRootOutcome{
			OutcomeType:                     int32(outcome.MerkleRootOutcome.OutcomeType),
			RangesSelectedForReport:         c.tr.chainRangeToProto(outcome.MerkleRootOutcome.RangesSelectedForReport),
			RootsToReport:                   c.tr.merkleRootsToProto(outcome.MerkleRootOutcome.RootsToReport),
			RmnEnabledChains:                c.tr.rmnEnabledChainsToProto(outcome.MerkleRootOutcome.RMNEnabledChains),
			OffRampNextSeqNums:              c.tr.seqNumChainToProto(outcome.MerkleRootOutcome.OffRampNextSeqNums),
			ReportTransmissionCheckAttempts: uint32(outcome.MerkleRootOutcome.ReportTransmissionCheckAttempts),
			RmnReportSignatures:             c.tr.ccipRmnSignaturesToProto(outcome.MerkleRootOutcome.RMNReportSignatures),
			RmnRemoteCfg:                    c.tr.rmnRemoteConfigToProto(outcome.MerkleRootOutcome.RMNRemoteCfg),
		},
		TokenPriceOutcome: &ocrtypecodecpb.TokenPriceOutcome{
			TokenPrices: c.tr.feedTokenPricesToProto(outcome.TokenPriceOutcome.TokenPrices),
		},
		ChainFeeOutcome: &ocrtypecodecpb.ChainFeeOutcome{
			GasPrices: c.tr.gasPriceChainToProto(outcome.ChainFeeOutcome.GasPrices),
		},
		MainOutcome: &ocrtypecodecpb.MainOutcome{
			InflightPriceOcrSequenceNumber: uint64(outcome.MainOutcome.InflightPriceOcrSequenceNumber),
			RemainingPriceChecks:           int32(outcome.MainOutcome.RemainingPriceChecks),
		},
	}

	return proto.MarshalOptions{Deterministic: true}.Marshal(pbOutcome)
}

func (c *CommitCodecProto) DecodeOutcome(data []byte) (committypes.Outcome, error) {
	if len(data) == 0 {
		return committypes.Outcome{}, nil
	}

	pbOutcome := &ocrtypecodecpb.CommitOutcome{}
	if err := proto.Unmarshal(data, pbOutcome); err != nil {
		return committypes.Outcome{}, fmt.Errorf("proto unmarshal outcome: %w", err)
	}

	return committypes.Outcome{
		MerkleRootOutcome: merkleroot.Outcome{
			OutcomeType:                     merkleroot.OutcomeType(pbOutcome.MerkleRootOutcome.OutcomeType),
			RangesSelectedForReport:         c.tr.chainRangeFromProto(pbOutcome.MerkleRootOutcome.RangesSelectedForReport),
			RootsToReport:                   c.tr.merkleRootsFromProto(pbOutcome.MerkleRootOutcome.RootsToReport),
			RMNEnabledChains:                c.tr.rmnEnabledChainsFromProto(pbOutcome.MerkleRootOutcome.RmnEnabledChains),
			OffRampNextSeqNums:              c.tr.seqNumChainFromProto(pbOutcome.MerkleRootOutcome.OffRampNextSeqNums),
			ReportTransmissionCheckAttempts: uint(pbOutcome.MerkleRootOutcome.ReportTransmissionCheckAttempts),
			RMNReportSignatures:             c.tr.ccipRmnSignaturesFromProto(pbOutcome.MerkleRootOutcome.RmnReportSignatures),
			RMNRemoteCfg:                    c.tr.rmnRemoteConfigFromProto(pbOutcome.MerkleRootOutcome.RmnRemoteCfg),
		},
		TokenPriceOutcome: tokenprice.Outcome{
			TokenPrices: c.tr.feedTokenPricesFromProto(pbOutcome.TokenPriceOutcome.TokenPrices),
		},
		ChainFeeOutcome: chainfee.Outcome{
			GasPrices: c.tr.gasPriceChainFromProto(pbOutcome.ChainFeeOutcome.GasPrices),
		},
		MainOutcome: committypes.MainOutcome{
			InflightPriceOcrSequenceNumber: cciptypes.SeqNum(pbOutcome.MainOutcome.InflightPriceOcrSequenceNumber),
			RemainingPriceChecks:           int(pbOutcome.MainOutcome.RemainingPriceChecks),
		},
	}, nil
}

// CommitCodecJSON is an implementation of CommitCodec that uses JSON.
// DEPRECATED: Use CommitCodecProto instead.
type CommitCodecJSON struct{}

// NewCommitCodecJSON returns a new CommitCodecJSON.
// DEPRECATED
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
