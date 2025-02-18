package ocrtypecodec

import (
	"encoding/json"
	"fmt"
	"math/big"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/smartcontractkit/chainlink-ccip/commit/chainfee"
	"github.com/smartcontractkit/chainlink-ccip/commit/committypes"
	"github.com/smartcontractkit/chainlink-ccip/commit/merkleroot"
	"github.com/smartcontractkit/chainlink-ccip/commit/merkleroot/rmn"
	rmntypes "github.com/smartcontractkit/chainlink-ccip/commit/merkleroot/rmn/types"
	"github.com/smartcontractkit/chainlink-ccip/commit/tokenprice"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon/discovery/discoverytypes"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugintypes"
	"github.com/smartcontractkit/chainlink-ccip/pkg/ocrtypecodec/ocrtypecodecpb"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
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

	sigs := c.tr.rmnSignaturesToProto(query.MerkleRootQuery.RMNSignatures)
	laneUpdates := c.tr.laneUpdatesToProto(query.MerkleRootQuery.RMNSignatures.LaneUpdates)

	pbQ := &ocrtypecodecpb.Query{
		MerkleRootQuery: &ocrtypecodecpb.MerkleRootQuery{
			RetryRmnSignatures: query.MerkleRootQuery.RetryRMNSignatures,
			RmnSignatures: &ocrtypecodecpb.ReportSignatures{
				Signatures:  sigs,
				LaneUpdates: laneUpdates,
			},
		},
		TokenPriceQuery: &ocrtypecodecpb.TokenPriceQuery{}, // always empty
		ChainFeeQuery:   &ocrtypecodecpb.ChainFeeQuery{},   // always empty
	}

	return proto.Marshal(pbQ)
}

func (c *CommitCodecProto) DecodeQuery(data []byte) (committypes.Query, error) {
	pbQ := &ocrtypecodecpb.Query{}
	if err := proto.Unmarshal(data, pbQ); err != nil {
		return committypes.Query{}, fmt.Errorf("decode query: %w", err)
	}

	sigs := c.tr.rmnSignaturesFromProto(pbQ.MerkleRootQuery.RmnSignatures.Signatures)
	laneUpdates := c.tr.laneUpdatesFromProto(pbQ.MerkleRootQuery.RmnSignatures.LaneUpdates)

	q := committypes.Query{
		MerkleRootQuery: merkleroot.Query{
			RetryRMNSignatures: pbQ.MerkleRootQuery.RetryRmnSignatures,
			RMNSignatures: &rmn.ReportSignatures{
				Signatures:  sigs,
				LaneUpdates: laneUpdates,
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
	pbObs := &ocrtypecodecpb.CommitObservation{}
	if err := proto.Unmarshal(data, pbObs); err != nil {
		return committypes.Observation{}, err
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
	rangesSelectedForReport := make([]*ocrtypecodecpb.ChainRange, len(outcome.MerkleRootOutcome.RangesSelectedForReport))
	for i, r := range outcome.MerkleRootOutcome.RangesSelectedForReport {
		rangesSelectedForReport[i] = &ocrtypecodecpb.ChainRange{
			ChainSel: uint64(r.ChainSel),
			SeqNumRange: &ocrtypecodecpb.SeqNumRange{
				MinMsgNr: uint64(r.SeqNumRange.Start()),
				MaxMsgNr: uint64(r.SeqNumRange.End()),
			},
		}
	}

	rootsToReport := make([]*ocrtypecodecpb.MerkleRootChain, len(outcome.MerkleRootOutcome.RootsToReport))
	for i, root := range outcome.MerkleRootOutcome.RootsToReport {
		rootsToReport[i] = &ocrtypecodecpb.MerkleRootChain{
			ChainSel:      uint64(root.ChainSel),
			OnRampAddress: root.OnRampAddress,
			SeqNumsRange: &ocrtypecodecpb.SeqNumRange{
				MinMsgNr: uint64(root.SeqNumsRange.Start()),
				MaxMsgNr: uint64(root.SeqNumsRange.End()),
			},
			MerkleRoot: root.MerkleRoot[:],
		}
	}

	rmnEnabledChains := make(map[uint64]bool, len(outcome.MerkleRootOutcome.RMNEnabledChains))
	for k, v := range outcome.MerkleRootOutcome.RMNEnabledChains {
		rmnEnabledChains[uint64(k)] = v
	}

	offRampNextSeqNums := make([]*ocrtypecodecpb.SeqNumChain, len(outcome.MerkleRootOutcome.OffRampNextSeqNums))
	for i, s := range outcome.MerkleRootOutcome.OffRampNextSeqNums {
		offRampNextSeqNums[i] = &ocrtypecodecpb.SeqNumChain{
			ChainSel: uint64(s.ChainSel),
			SeqNum:   uint64(s.SeqNum),
		}
	}

	rmnReportSignatures := make([]*ocrtypecodecpb.SignatureEcdsa, len(outcome.MerkleRootOutcome.RMNReportSignatures))
	for i, sig := range outcome.MerkleRootOutcome.RMNReportSignatures {
		rmnReportSignatures[i] = &ocrtypecodecpb.SignatureEcdsa{
			R: sig.R[:],
			S: sig.S[:],
		}
	}

	pbMerkleRootOutcome := &ocrtypecodecpb.MerkleRootOutcome{
		OutcomeType:                     int32(outcome.MerkleRootOutcome.OutcomeType),
		RangesSelectedForReport:         rangesSelectedForReport,
		RootsToReport:                   rootsToReport,
		RmnEnabledChains:                rmnEnabledChains,
		OffRampNextSeqNums:              offRampNextSeqNums,
		ReportTransmissionCheckAttempts: uint32(outcome.MerkleRootOutcome.ReportTransmissionCheckAttempts),
		RmnReportSignatures:             rmnReportSignatures,
		RmnRemoteCfg: &ocrtypecodecpb.RmnRemoteConfig{
			ContractAddress:  outcome.MerkleRootOutcome.RMNRemoteCfg.ContractAddress,
			ConfigDigest:     outcome.MerkleRootOutcome.RMNRemoteCfg.ConfigDigest[:],
			Signers:          encodeRemoteSigners(outcome.MerkleRootOutcome.RMNRemoteCfg.Signers),
			FSign:            outcome.MerkleRootOutcome.RMNRemoteCfg.FSign,
			ConfigVersion:    outcome.MerkleRootOutcome.RMNRemoteCfg.ConfigVersion,
			RmnReportVersion: outcome.MerkleRootOutcome.RMNRemoteCfg.RmnReportVersion[:],
		},
	}

	// Encode TokenPriceOutcome
	tokenPrices := make(map[string][]byte, len(outcome.TokenPriceOutcome.TokenPrices))
	for k, v := range outcome.TokenPriceOutcome.TokenPrices {
		tokenPrices[string(k)] = v.Bytes()
	}

	pbTokenPriceOutcome := &ocrtypecodecpb.TokenPriceOutcome{
		TokenPrices: tokenPrices,
	}

	// Encode ChainFeeOutcome
	gasPrices := make([]*ocrtypecodecpb.GasPriceChain, len(outcome.ChainFeeOutcome.GasPrices))
	for i, gp := range outcome.ChainFeeOutcome.GasPrices {
		gasPrices[i] = &ocrtypecodecpb.GasPriceChain{
			ChainSel: uint64(gp.ChainSel),
			GasPrice: gp.GasPrice.Bytes(),
		}
	}

	pbChainFeeOutcome := &ocrtypecodecpb.ChainFeeOutcome{
		GasPrices: gasPrices,
	}

	// Encode MainOutcome
	pbMainOutcome := &ocrtypecodecpb.MainOutcome{
		InflightPriceOcrSequenceNumber: uint64(outcome.MainOutcome.InflightPriceOcrSequenceNumber),
		RemainingPriceChecks:           int32(outcome.MainOutcome.RemainingPriceChecks),
	}

	pbOutcome := &ocrtypecodecpb.CommitOutcome{
		MerkleRootOutcome: pbMerkleRootOutcome,
		TokenPriceOutcome: pbTokenPriceOutcome,
		ChainFeeOutcome:   pbChainFeeOutcome,
		MainOutcome:       pbMainOutcome,
	}

	return proto.Marshal(pbOutcome)
}

// Helper function to encode RemoteSignerInfo
func encodeRemoteSigners(signers []rmntypes.RemoteSignerInfo) []*ocrtypecodecpb.RemoteSignerInfo {
	pbSigners := make([]*ocrtypecodecpb.RemoteSignerInfo, len(signers))
	for i, s := range signers {
		pbSigners[i] = &ocrtypecodecpb.RemoteSignerInfo{
			OnchainPublicKey: s.OnchainPublicKey,
			NodeIndex:        s.NodeIndex,
		}
	}
	return pbSigners
}

func (c *CommitCodecProto) DecodeOutcome(data []byte) (committypes.Outcome, error) {
	pbOutcome := &ocrtypecodecpb.CommitOutcome{}
	if err := proto.Unmarshal(data, pbOutcome); err != nil {
		return committypes.Outcome{}, err
	}

	rangesSelectedForReport := make([]plugintypes.ChainRange, len(pbOutcome.MerkleRootOutcome.RangesSelectedForReport))
	for i, r := range pbOutcome.MerkleRootOutcome.RangesSelectedForReport {
		rangesSelectedForReport[i] = plugintypes.ChainRange{
			ChainSel: cciptypes.ChainSelector(r.ChainSel),
			SeqNumRange: cciptypes.NewSeqNumRange(
				cciptypes.SeqNum(r.SeqNumRange.MinMsgNr),
				cciptypes.SeqNum(r.SeqNumRange.MaxMsgNr),
			),
		}
	}

	rootsToReport := make([]cciptypes.MerkleRootChain, len(pbOutcome.MerkleRootOutcome.RootsToReport))
	for i, root := range pbOutcome.MerkleRootOutcome.RootsToReport {
		rootsToReport[i] = cciptypes.MerkleRootChain{
			ChainSel:      cciptypes.ChainSelector(root.ChainSel),
			OnRampAddress: root.OnRampAddress,
			SeqNumsRange: cciptypes.NewSeqNumRange(
				cciptypes.SeqNum(root.SeqNumsRange.MinMsgNr),
				cciptypes.SeqNum(root.SeqNumsRange.MaxMsgNr),
			),
			MerkleRoot: cciptypes.Bytes32(root.MerkleRoot),
		}
	}

	rmnEnabledChains := make(map[cciptypes.ChainSelector]bool, len(pbOutcome.MerkleRootOutcome.RmnEnabledChains))
	for k, v := range pbOutcome.MerkleRootOutcome.RmnEnabledChains {
		rmnEnabledChains[cciptypes.ChainSelector(k)] = v
	}

	offRampNextSeqNums := make([]plugintypes.SeqNumChain, len(pbOutcome.MerkleRootOutcome.OffRampNextSeqNums))
	for i, s := range pbOutcome.MerkleRootOutcome.OffRampNextSeqNums {
		offRampNextSeqNums[i] = plugintypes.SeqNumChain{
			ChainSel: cciptypes.ChainSelector(s.ChainSel),
			SeqNum:   cciptypes.SeqNum(s.SeqNum),
		}
	}

	rmnReportSignatures := make([]cciptypes.RMNECDSASignature, len(pbOutcome.MerkleRootOutcome.RmnReportSignatures))
	for i, sig := range pbOutcome.MerkleRootOutcome.RmnReportSignatures {
		rmnReportSignatures[i] = cciptypes.RMNECDSASignature{
			R: cciptypes.Bytes32(sig.R),
			S: cciptypes.Bytes32(sig.S),
		}
	}

	merkleRootOutcome := merkleroot.Outcome{
		OutcomeType:                     merkleroot.OutcomeType(pbOutcome.MerkleRootOutcome.OutcomeType),
		RangesSelectedForReport:         rangesSelectedForReport,
		RootsToReport:                   rootsToReport,
		RMNEnabledChains:                rmnEnabledChains,
		OffRampNextSeqNums:              offRampNextSeqNums,
		ReportTransmissionCheckAttempts: uint(pbOutcome.MerkleRootOutcome.ReportTransmissionCheckAttempts),
		RMNReportSignatures:             rmnReportSignatures,
		RMNRemoteCfg: rmntypes.RemoteConfig{
			ContractAddress: pbOutcome.MerkleRootOutcome.RmnRemoteCfg.ContractAddress,
			ConfigDigest:    cciptypes.Bytes32(pbOutcome.MerkleRootOutcome.RmnRemoteCfg.ConfigDigest),
			Signers:         decodeRemoteSigners(pbOutcome.MerkleRootOutcome.RmnRemoteCfg.Signers),
			FSign:           pbOutcome.MerkleRootOutcome.RmnRemoteCfg.FSign,
			ConfigVersion:   pbOutcome.MerkleRootOutcome.RmnRemoteCfg.ConfigVersion,
			RmnReportVersion: cciptypes.Bytes32(
				pbOutcome.MerkleRootOutcome.RmnRemoteCfg.RmnReportVersion,
			),
		},
	}

	tokenPrices := make(cciptypes.TokenPriceMap, len(pbOutcome.TokenPriceOutcome.TokenPrices))
	for k, v := range pbOutcome.TokenPriceOutcome.TokenPrices {
		tokenPrices[cciptypes.UnknownEncodedAddress(k)] = cciptypes.NewBigInt(big.NewInt(0).SetBytes(v))
	}

	gasPrices := make([]cciptypes.GasPriceChain, len(pbOutcome.ChainFeeOutcome.GasPrices))
	for i, gp := range pbOutcome.ChainFeeOutcome.GasPrices {
		gasPrices[i] = cciptypes.GasPriceChain{
			ChainSel: cciptypes.ChainSelector(gp.ChainSel),
			GasPrice: cciptypes.NewBigInt(big.NewInt(0).SetBytes(gp.GasPrice)),
		}
	}

	return committypes.Outcome{
		MerkleRootOutcome: merkleRootOutcome,
		TokenPriceOutcome: tokenprice.Outcome{
			TokenPrices: tokenPrices,
		},
		ChainFeeOutcome: chainfee.Outcome{
			GasPrices: gasPrices,
		},
		MainOutcome: committypes.MainOutcome{
			InflightPriceOcrSequenceNumber: cciptypes.SeqNum(pbOutcome.MainOutcome.InflightPriceOcrSequenceNumber),
			RemainingPriceChecks:           int(pbOutcome.MainOutcome.RemainingPriceChecks),
		},
	}, nil
}

// Helper function to decode RemoteSignerInfo
func decodeRemoteSigners(signers []*ocrtypecodecpb.RemoteSignerInfo) []rmntypes.RemoteSignerInfo {
	decoded := make([]rmntypes.RemoteSignerInfo, len(signers))
	for i, s := range signers {
		decoded[i] = rmntypes.RemoteSignerInfo{
			OnchainPublicKey: s.OnchainPublicKey,
			NodeIndex:        s.NodeIndex,
		}
	}
	return decoded
}

// CommitCodecJSON is an implementation of CommitCodec that uses JSON.
// DEPRECATED: Use CommitCodecProto instead.
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
