package v1

import (
	"encoding/json"
	"fmt"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-ccip/commit/chainfee"
	"github.com/smartcontractkit/chainlink-ccip/commit/committypes"
	"github.com/smartcontractkit/chainlink-ccip/commit/merkleroot"
	"github.com/smartcontractkit/chainlink-ccip/commit/merkleroot/rmn"
	"github.com/smartcontractkit/chainlink-ccip/commit/tokenprice"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon/discovery/discoverytypes"
	"github.com/smartcontractkit/chainlink-ccip/pkg/ocrtypecodec/v1/ocrtypecodecpb"
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
			RetryRMNSignatures: pbQ.GetMerkleRootQuery().GetRetryRmnSignatures(),
			RMNSignatures: &rmn.ReportSignatures{
				Signatures:  c.tr.rmnSignaturesFromProto(pbQ.GetMerkleRootQuery().GetRmnSignatures().GetSignatures()),
				LaneUpdates: c.tr.laneUpdatesFromProto(pbQ.GetMerkleRootQuery().GetRmnSignatures().GetLaneUpdates()),
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
		FChain:                c.tr.fChainToProto(observation.FChain),
		OnchainPriceOcrSeqNum: observation.OnChainPriceOcrSeqNum,
	}
	return proto.Marshal(pbObs)
}

func (c *CommitCodecProto) DecodeObservation(data []byte) (obs committypes.Observation, err error) {
	if len(data) == 0 {
		return obs, nil
	}

	pbObs := &ocrtypecodecpb.CommitObservation{}
	if err := proto.Unmarshal(data, pbObs); err != nil {
		return obs, fmt.Errorf("proto unmarshal observation: %w", err)
	}

	merkleRoots, err := c.tr.merkleRootsFromProto(pbObs.GetMerkleRootObs().GetMerkleRoots())
	if err != nil {
		return obs, fmt.Errorf("merkle roots from proto: %w", err)
	}
	rmnRemoteCfg, err := c.tr.rmnRemoteConfigFromProto(pbObs.GetMerkleRootObs().GetRmnRemoteConfig())
	if err != nil {
		return obs, fmt.Errorf("rmn remote config from proto: %w", err)
	}
	return committypes.Observation{
		MerkleRootObs: merkleroot.Observation{
			MerkleRoots:        merkleRoots,
			RMNEnabledChains:   c.tr.rmnEnabledChainsFromProto(pbObs.GetMerkleRootObs().GetRmnEnabledChains()),
			OnRampMaxSeqNums:   c.tr.seqNumChainFromProto(pbObs.GetMerkleRootObs().GetOnRampMaxSeqNums()),
			OffRampNextSeqNums: c.tr.seqNumChainFromProto(pbObs.GetMerkleRootObs().GetOffRampNextSeqNums()),
			RMNRemoteConfig:    rmnRemoteCfg,
			FChain:             c.tr.fChainFromProto(pbObs.GetMerkleRootObs().GetFChain()),
		},
		TokenPriceObs: tokenprice.Observation{
			FeedTokenPrices:       c.tr.feedTokenPricesFromProto(pbObs.GetTokenPriceObs().GetFeedTokenPrices()),
			FeeQuoterTokenUpdates: c.tr.feeQuoterTokenUpdatesFromProto(pbObs.GetTokenPriceObs().GetFeeQuoterTokenUpdates()),
			FChain:                c.tr.fChainFromProto(pbObs.GetTokenPriceObs().GetFChain()),
			Timestamp:             pbObs.GetTokenPriceObs().GetTimestamp().AsTime(),
		},
		ChainFeeObs: chainfee.Observation{
			FeeComponents:     c.tr.feeComponentsFromProto(pbObs.GetChainFeeObs().GetFeeComponents()),
			NativeTokenPrices: c.tr.nativeTokenPricesFromProto(pbObs.GetChainFeeObs().GetNativeTokenPrices()),
			ChainFeeUpdates:   c.tr.chainFeeUpdatesFromProto(pbObs.GetChainFeeObs().GetChainFeeUpdates()),
			FChain:            c.tr.fChainFromProto(pbObs.GetChainFeeObs().GetFChain()),
			TimestampNow:      pbObs.GetChainFeeObs().GetTimestampNow().AsTime(),
		},
		DiscoveryObs: discoverytypes.Observation{
			FChain:    c.tr.fChainFromProto(pbObs.GetDiscoveryObs().GetFChain()),
			Addresses: c.tr.discoveryAddressesFromProto(pbObs.GetDiscoveryObs().GetContractNames().GetAddresses()),
		},
		FChain:                c.tr.fChainFromProto(pbObs.GetFChain()),
		OnChainPriceOcrSeqNum: pbObs.GetOnchainPriceOcrSeqNum(),
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

	rootsToReport, err := c.tr.merkleRootsFromProto(pbOutcome.GetMerkleRootOutcome().GetRootsToReport())
	if err != nil {
		return committypes.Outcome{}, fmt.Errorf("merkle roots from proto: %w", err)
	}
	sigs, err := c.tr.ccipRmnSignaturesFromProto(pbOutcome.GetMerkleRootOutcome().GetRmnReportSignatures())
	if err != nil {
		return committypes.Outcome{}, fmt.Errorf("rmn report signatures from proto: %w", err)
	}
	rmnRemoteCfg, err := c.tr.rmnRemoteConfigFromProto(pbOutcome.GetMerkleRootOutcome().GetRmnRemoteCfg())
	if err != nil {
		return committypes.Outcome{}, fmt.Errorf("rmn remote config from proto: %w", err)
	}
	return committypes.Outcome{
		MerkleRootOutcome: merkleroot.Outcome{
			OutcomeType: merkleroot.OutcomeType(pbOutcome.GetMerkleRootOutcome().GetOutcomeType()),
			RangesSelectedForReport: c.tr.chainRangeFromProto(
				pbOutcome.GetMerkleRootOutcome().GetRangesSelectedForReport(),
			),
			RootsToReport: rootsToReport,
			RMNEnabledChains: c.tr.rmnEnabledChainsFromProto(
				pbOutcome.GetMerkleRootOutcome().GetRmnEnabledChains(),
			),
			OffRampNextSeqNums:              c.tr.seqNumChainFromProto(pbOutcome.GetMerkleRootOutcome().GetOffRampNextSeqNums()),
			ReportTransmissionCheckAttempts: uint(pbOutcome.GetMerkleRootOutcome().GetReportTransmissionCheckAttempts()),
			RMNReportSignatures:             sigs,
			RMNRemoteCfg:                    rmnRemoteCfg,
		},
		TokenPriceOutcome: tokenprice.Outcome{
			TokenPrices: c.tr.feedTokenPricesFromProto(pbOutcome.GetTokenPriceOutcome().GetTokenPrices()),
		},
		ChainFeeOutcome: chainfee.Outcome{
			GasPrices: c.tr.gasPriceChainFromProto(pbOutcome.GetChainFeeOutcome().GetGasPrices()),
		},
		MainOutcome: committypes.MainOutcome{
			InflightPriceOcrSequenceNumber: cciptypes.SeqNum(pbOutcome.GetMainOutcome().GetInflightPriceOcrSequenceNumber()),
			RemainingPriceChecks:           int(pbOutcome.GetMainOutcome().GetRemainingPriceChecks()),
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
