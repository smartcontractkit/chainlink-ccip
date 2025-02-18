package ocrtypecodec

import (
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/smartcontractkit/chainlink-common/pkg/types"
	rmnpb "github.com/smartcontractkit/chainlink-protos/rmn/v1.6/go/serialization"
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
	"github.com/smartcontractkit/chainlink-ccip/pkg/reader"
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

type CommitCodecProto struct{}

func NewCommitCodecProto() *CommitCodecProto {
	return &CommitCodecProto{}
}

func (c *CommitCodecProto) EncodeQuery(query committypes.Query) ([]byte, error) {
	sigs := make([]*ocrtypecodecpb.SignatureEcdsa, 0)
	if query.MerkleRootQuery.RMNSignatures != nil {
		sigs = make([]*ocrtypecodecpb.SignatureEcdsa, len(query.MerkleRootQuery.RMNSignatures.Signatures))
		for i, sig := range query.MerkleRootQuery.RMNSignatures.Signatures {
			sigs[i] = &ocrtypecodecpb.SignatureEcdsa{R: sig.R, S: sig.S}
		}
	}

	laneUpdates := make([]*ocrtypecodecpb.DestChainUpdate, 0)
	if query.MerkleRootQuery.RMNSignatures != nil {
		laneUpdates = make([]*ocrtypecodecpb.DestChainUpdate, len(query.MerkleRootQuery.RMNSignatures.LaneUpdates))
		for i, lu := range query.MerkleRootQuery.RMNSignatures.LaneUpdates {
			laneUpdates[i] = &ocrtypecodecpb.DestChainUpdate{
				LaneSource: &ocrtypecodecpb.SourceChainMeta{
					SourceChainSelector: lu.LaneSource.SourceChainSelector,
					OnrampAddress:       lu.LaneSource.OnrampAddress,
				},
				SeqNumRange: &ocrtypecodecpb.SeqNumRange{
					MinMsgNr: lu.ClosedInterval.MinMsgNr,
					MaxMsgNr: lu.ClosedInterval.MaxMsgNr,
				},
				Root: lu.Root,
			}
		}
	}

	pbQ := &ocrtypecodecpb.Query{
		MerkleRootQuery: &ocrtypecodecpb.MerkleRootQuery{
			RetryRmnSignatures: query.MerkleRootQuery.RetryRMNSignatures,
			RmnSignatures: &ocrtypecodecpb.ReportSignatures{
				Signatures:  sigs,
				LaneUpdates: laneUpdates,
			},
		},
		TokenPriceQuery: &ocrtypecodecpb.TokenPriceQuery{},
		ChainFeeQuery:   &ocrtypecodecpb.ChainFeeQuery{},
	}

	return proto.Marshal(pbQ)
}

func (c *CommitCodecProto) DecodeQuery(data []byte) (committypes.Query, error) {
	pbQ := &ocrtypecodecpb.Query{}
	if err := proto.Unmarshal(data, pbQ); err != nil {
		return committypes.Query{}, fmt.Errorf("decode query: %w", err)
	}

	sigs := make([]*rmnpb.EcdsaSignature, len(pbQ.MerkleRootQuery.RmnSignatures.Signatures))
	for i := range pbQ.MerkleRootQuery.RmnSignatures.Signatures {
		sigs[i] = &rmnpb.EcdsaSignature{
			R: pbQ.MerkleRootQuery.RmnSignatures.Signatures[i].R,
			S: pbQ.MerkleRootQuery.RmnSignatures.Signatures[i].S,
		}
	}

	laneUpdates := make([]*rmnpb.FixedDestLaneUpdate, len(pbQ.MerkleRootQuery.RmnSignatures.LaneUpdates))
	for i := range pbQ.MerkleRootQuery.RmnSignatures.LaneUpdates {
		laneUpdates[i] = &rmnpb.FixedDestLaneUpdate{
			LaneSource: &rmnpb.LaneSource{
				SourceChainSelector: pbQ.MerkleRootQuery.RmnSignatures.LaneUpdates[i].LaneSource.SourceChainSelector,
				OnrampAddress:       pbQ.MerkleRootQuery.RmnSignatures.LaneUpdates[i].LaneSource.OnrampAddress,
			},
			ClosedInterval: &rmnpb.ClosedInterval{
				MinMsgNr: pbQ.MerkleRootQuery.RmnSignatures.LaneUpdates[i].SeqNumRange.MinMsgNr,
				MaxMsgNr: pbQ.MerkleRootQuery.RmnSignatures.LaneUpdates[i].SeqNumRange.MaxMsgNr,
			},
			Root: pbQ.MerkleRootQuery.RmnSignatures.LaneUpdates[i].Root,
		}
	}

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

//nolint:gocyclo
func (c *CommitCodecProto) EncodeObservation(observation committypes.Observation) ([]byte, error) {
	merkleRoots := make([]*ocrtypecodecpb.MerkleRootChain, len(observation.MerkleRootObs.MerkleRoots))
	for i, mr := range observation.MerkleRootObs.MerkleRoots {
		merkleRoots[i] = &ocrtypecodecpb.MerkleRootChain{
			ChainSel:      uint64(mr.ChainSel),
			OnRampAddress: mr.OnRampAddress,
			SeqNumsRange: &ocrtypecodecpb.SeqNumRange{
				MinMsgNr: uint64(mr.SeqNumsRange.Start()),
				MaxMsgNr: uint64(mr.SeqNumsRange.End()),
			},
			MerkleRoot: mr.MerkleRoot[:],
		}
	}

	rmnEnabledChains := make(map[uint64]bool, len(observation.MerkleRootObs.RMNEnabledChains))
	for k, v := range observation.MerkleRootObs.RMNEnabledChains {
		rmnEnabledChains[uint64(k)] = v
	}

	onRampMaxSeqNums := make([]*ocrtypecodecpb.SeqNumChain, len(observation.MerkleRootObs.OnRampMaxSeqNums))
	for i, s := range observation.MerkleRootObs.OnRampMaxSeqNums {
		onRampMaxSeqNums[i] = &ocrtypecodecpb.SeqNumChain{
			ChainSel: uint64(s.ChainSel),
			SeqNum:   uint64(s.SeqNum),
		}
	}

	offRampNextSeqNums := make([]*ocrtypecodecpb.SeqNumChain, len(observation.MerkleRootObs.OffRampNextSeqNums))
	for i, s := range observation.MerkleRootObs.OffRampNextSeqNums {
		offRampNextSeqNums[i] = &ocrtypecodecpb.SeqNumChain{
			ChainSel: uint64(s.ChainSel),
			SeqNum:   uint64(s.SeqNum),
		}
	}

	rmnRemoteConfigSigners := make(
		[]*ocrtypecodecpb.RemoteSignerInfo, len(observation.MerkleRootObs.RMNRemoteConfig.Signers))
	for i, s := range observation.MerkleRootObs.RMNRemoteConfig.Signers {
		rmnRemoteConfigSigners[i] = &ocrtypecodecpb.RemoteSignerInfo{
			OnchainPublicKey: s.OnchainPublicKey,
			NodeIndex:        s.NodeIndex,
		}
	}

	merkleRootsFChain := make(map[uint64]int32, len(observation.FChain))
	for k, v := range observation.MerkleRootObs.FChain {
		merkleRootsFChain[uint64(k)] = int32(v)
	}

	feedTokenPrices := make(map[string][]byte, len(observation.TokenPriceObs.FeedTokenPrices))
	for k, v := range observation.TokenPriceObs.FeedTokenPrices {
		feedTokenPrices[string(k)] = v.Bytes()
	}

	feeQuoterTokenUpdates := make(
		map[string]*ocrtypecodecpb.TimestampedBig, len(observation.TokenPriceObs.FeeQuoterTokenUpdates))
	for k, v := range observation.TokenPriceObs.FeeQuoterTokenUpdates {
		feeQuoterTokenUpdates[string(k)] = &ocrtypecodecpb.TimestampedBig{
			Value:     v.Value.Bytes(),
			Timestamp: timestamppb.New(v.Timestamp),
		}
	}

	tokenPriceFChain := make(map[uint64]int32, len(observation.TokenPriceObs.FChain))
	for k, v := range observation.TokenPriceObs.FChain {
		tokenPriceFChain[uint64(k)] = int32(v)
	}

	feeComponents := make(map[uint64]*ocrtypecodecpb.ChainFeeComponents, len(observation.ChainFeeObs.FeeComponents))
	for k, v := range observation.ChainFeeObs.FeeComponents {
		feeComponents[uint64(k)] = &ocrtypecodecpb.ChainFeeComponents{
			ExecutionFee:        v.ExecutionFee.Bytes(),
			DataAvailabilityFee: v.DataAvailabilityFee.Bytes(),
		}
	}

	nativeTokenPrices := make(map[uint64][]byte, len(observation.ChainFeeObs.NativeTokenPrices))
	for k, v := range observation.ChainFeeObs.NativeTokenPrices {
		nativeTokenPrices[uint64(k)] = v.Bytes()
	}

	chainFeeUpdates := make(map[uint64]*ocrtypecodecpb.ChainFeeUpdate, len(observation.ChainFeeObs.ChainFeeUpdates))
	for k, v := range observation.ChainFeeObs.ChainFeeUpdates {
		chainFeeUpdates[uint64(k)] = &ocrtypecodecpb.ChainFeeUpdate{
			ChainFee: &ocrtypecodecpb.ComponentsUSDPrices{
				ExecutionFeePriceUsd: v.ChainFee.ExecutionFeePriceUSD.Bytes(),
				DataAvFeePriceUsd:    v.ChainFee.DataAvFeePriceUSD.Bytes(),
			},
			Timestamp: timestamppb.New(v.Timestamp),
		}
	}

	chainFeeFChain := make(map[uint64]int32, len(observation.ChainFeeObs.FChain))
	for k, v := range observation.ChainFeeObs.FChain {
		chainFeeFChain[uint64(k)] = int32(v)
	}

	discoveryFChain := make(map[uint64]int32, len(observation.DiscoveryObs.FChain))
	for k, v := range observation.DiscoveryObs.FChain {
		discoveryFChain[uint64(k)] = int32(v)
	}

	discoveryAddrs := make(map[string]*ocrtypecodecpb.ChainAddressMap, len(observation.DiscoveryObs.Addresses))
	for contractName, chains := range observation.DiscoveryObs.Addresses {
		discoveryAddrs[contractName] = &ocrtypecodecpb.ChainAddressMap{
			ChainAddresses: make(map[uint64][]byte, len(chains))}

		for chain, addr := range chains {
			discoveryAddrs[contractName].ChainAddresses[uint64(chain)] = addr
		}
	}

	mainFChain := make(map[uint64]int32, len(observation.FChain))
	for k, v := range observation.FChain {
		mainFChain[uint64(k)] = int32(v)
	}

	pbObs := &ocrtypecodecpb.CommitObservation{
		MerkleRootObs: &ocrtypecodecpb.MerkleRootObservation{
			MerkleRoots:        merkleRoots,
			RmnEnabledChains:   rmnEnabledChains,
			OnRampMaxSeqNums:   onRampMaxSeqNums,
			OffRampNextSeqNums: offRampNextSeqNums,
			RmnRemoteConfig: &ocrtypecodecpb.RmnRemoteConfig{
				ContractAddress:  observation.MerkleRootObs.RMNRemoteConfig.ContractAddress,
				ConfigDigest:     observation.MerkleRootObs.RMNRemoteConfig.ConfigDigest[:],
				Signers:          rmnRemoteConfigSigners,
				FSign:            observation.MerkleRootObs.RMNRemoteConfig.FSign,
				ConfigVersion:    observation.MerkleRootObs.RMNRemoteConfig.ConfigVersion,
				RmnReportVersion: observation.MerkleRootObs.RMNRemoteConfig.RmnReportVersion[:],
			},
			FChain: merkleRootsFChain,
		},
		TokenPriceObs: &ocrtypecodecpb.TokenPriceObservation{
			FeedTokenPrices:       feedTokenPrices,
			FeeQuoterTokenUpdates: feeQuoterTokenUpdates,
			FChain:                tokenPriceFChain,
			Timestamp:             timestamppb.New(observation.TokenPriceObs.Timestamp),
		},
		ChainFeeObs: &ocrtypecodecpb.ChainFeeObservation{
			FeeComponents:     feeComponents,
			NativeTokenPrices: nativeTokenPrices,
			ChainFeeUpdates:   chainFeeUpdates,
			FChain:            chainFeeFChain,
			TimestampNow:      timestamppb.New(observation.ChainFeeObs.TimestampNow),
		},
		DiscoveryObs: &ocrtypecodecpb.DiscoveryObservation{
			FChain: discoveryFChain,
			ContractNames: &ocrtypecodecpb.ContractNameChainAddresses{
				Addresses: discoveryAddrs,
			},
		},
		FChain: mainFChain,
	}

	return proto.Marshal(pbObs)
}

//nolint:gocyclo
func (c *CommitCodecProto) DecodeObservation(data []byte) (committypes.Observation, error) {
	pbObs := &ocrtypecodecpb.CommitObservation{}
	if err := proto.Unmarshal(data, pbObs); err != nil {
		return committypes.Observation{}, err
	}

	merkleRoots := make([]cciptypes.MerkleRootChain, len(pbObs.MerkleRootObs.MerkleRoots))
	for i, mr := range pbObs.MerkleRootObs.MerkleRoots {
		merkleRoots[i] = cciptypes.MerkleRootChain{
			ChainSel:      cciptypes.ChainSelector(mr.ChainSel),
			OnRampAddress: mr.OnRampAddress,
			SeqNumsRange: cciptypes.NewSeqNumRange(
				cciptypes.SeqNum(mr.SeqNumsRange.MinMsgNr),
				cciptypes.SeqNum(mr.SeqNumsRange.MaxMsgNr),
			),
			MerkleRoot: cciptypes.Bytes32(mr.MerkleRoot),
		}
	}

	rmnEnabledChains := make(map[cciptypes.ChainSelector]bool, len(pbObs.MerkleRootObs.RmnEnabledChains))
	for k, v := range pbObs.MerkleRootObs.RmnEnabledChains {
		rmnEnabledChains[cciptypes.ChainSelector(k)] = v
	}

	onRampMaxSeqNums := make([]plugintypes.SeqNumChain, len(pbObs.MerkleRootObs.OnRampMaxSeqNums))
	for i, s := range pbObs.MerkleRootObs.OnRampMaxSeqNums {
		onRampMaxSeqNums[i] = plugintypes.SeqNumChain{
			ChainSel: cciptypes.ChainSelector(s.ChainSel),
			SeqNum:   cciptypes.SeqNum(s.SeqNum),
		}
	}

	offRampNextSeqNums := make([]plugintypes.SeqNumChain, len(pbObs.MerkleRootObs.OffRampNextSeqNums))
	for i, s := range pbObs.MerkleRootObs.OffRampNextSeqNums {
		offRampNextSeqNums[i] = plugintypes.SeqNumChain{
			ChainSel: cciptypes.ChainSelector(s.ChainSel),
			SeqNum:   cciptypes.SeqNum(s.SeqNum),
		}
	}

	rmnSigners := make([]rmntypes.RemoteSignerInfo, len(pbObs.MerkleRootObs.RmnRemoteConfig.Signers))
	for i, s := range pbObs.MerkleRootObs.RmnRemoteConfig.Signers {
		rmnSigners[i] = rmntypes.RemoteSignerInfo{
			OnchainPublicKey: s.OnchainPublicKey,
			NodeIndex:        s.NodeIndex,
		}
	}

	merkleRootsFChain := make(map[cciptypes.ChainSelector]int, len(pbObs.MerkleRootObs.FChain))
	for k, v := range pbObs.MerkleRootObs.FChain {
		merkleRootsFChain[cciptypes.ChainSelector(k)] = int(v)
	}

	feedTokenPrices := make(cciptypes.TokenPriceMap, len(pbObs.TokenPriceObs.FeedTokenPrices))
	for k, v := range pbObs.TokenPriceObs.FeedTokenPrices {
		feedTokenPrices[cciptypes.UnknownEncodedAddress(k)] = cciptypes.NewBigInt(big.NewInt(0).SetBytes(v))
	}

	feeQuoterTokenUpdates := make(
		map[cciptypes.UnknownEncodedAddress]plugintypes.TimestampedBig, len(pbObs.TokenPriceObs.FeeQuoterTokenUpdates))
	for k, v := range pbObs.TokenPriceObs.FeeQuoterTokenUpdates {
		feeQuoterTokenUpdates[cciptypes.UnknownEncodedAddress(k)] = plugintypes.TimestampedBig{
			Value:     cciptypes.NewBigInt(big.NewInt(0).SetBytes(v.Value)),
			Timestamp: v.Timestamp.AsTime(),
		}
	}

	tokenPriceFChain := make(map[cciptypes.ChainSelector]int, len(pbObs.TokenPriceObs.FChain))
	for k, v := range pbObs.TokenPriceObs.FChain {
		tokenPriceFChain[cciptypes.ChainSelector(k)] = int(v)
	}

	feeComponents := make(map[cciptypes.ChainSelector]types.ChainFeeComponents, len(pbObs.ChainFeeObs.FeeComponents))
	for k, v := range pbObs.ChainFeeObs.FeeComponents {
		feeComponents[cciptypes.ChainSelector(k)] = types.ChainFeeComponents{
			ExecutionFee:        big.NewInt(0).SetBytes(v.ExecutionFee),
			DataAvailabilityFee: big.NewInt(0).SetBytes(v.DataAvailabilityFee),
		}
	}

	nativeTokenPrices := make(map[cciptypes.ChainSelector]cciptypes.BigInt, len(pbObs.ChainFeeObs.NativeTokenPrices))
	for k, v := range pbObs.ChainFeeObs.NativeTokenPrices {
		nativeTokenPrices[cciptypes.ChainSelector(k)] = cciptypes.NewBigInt(big.NewInt(0).SetBytes(v))
	}

	chainFeeUpdates := make(map[cciptypes.ChainSelector]chainfee.Update, len(pbObs.ChainFeeObs.ChainFeeUpdates))
	for k, v := range pbObs.ChainFeeObs.ChainFeeUpdates {
		chainFeeUpdates[cciptypes.ChainSelector(k)] = chainfee.Update{
			ChainFee: chainfee.ComponentsUSDPrices{
				ExecutionFeePriceUSD: big.NewInt(0).SetBytes(v.ChainFee.ExecutionFeePriceUsd),
				DataAvFeePriceUSD:    big.NewInt(0).SetBytes(v.ChainFee.DataAvFeePriceUsd),
			},
			Timestamp: v.Timestamp.AsTime(),
		}
	}

	chainFeeFChain := make(map[cciptypes.ChainSelector]int, len(pbObs.ChainFeeObs.FChain))
	for k, v := range pbObs.ChainFeeObs.FChain {
		chainFeeFChain[cciptypes.ChainSelector(k)] = int(v)
	}

	discoveryFChain := make(map[cciptypes.ChainSelector]int, len(pbObs.DiscoveryObs.FChain))
	for k, v := range pbObs.DiscoveryObs.FChain {
		discoveryFChain[cciptypes.ChainSelector(k)] = int(v)
	}

	discoveryAddrs := make(reader.ContractAddresses, len(pbObs.DiscoveryObs.ContractNames.Addresses))
	for contractName, chainMap := range pbObs.DiscoveryObs.ContractNames.Addresses {
		discoveryAddrs[contractName] = make(map[cciptypes.ChainSelector]cciptypes.UnknownAddress)
		for chain, addr := range chainMap.ChainAddresses {
			discoveryAddrs[contractName][cciptypes.ChainSelector(chain)] = addr
		}
	}

	mainFChain := make(map[cciptypes.ChainSelector]int, len(pbObs.FChain))
	for k, v := range pbObs.FChain {
		mainFChain[cciptypes.ChainSelector(k)] = int(v)
	}

	return committypes.Observation{
		MerkleRootObs: merkleroot.Observation{
			MerkleRoots:        merkleRoots,
			RMNEnabledChains:   rmnEnabledChains,
			OnRampMaxSeqNums:   onRampMaxSeqNums,
			OffRampNextSeqNums: offRampNextSeqNums,
			RMNRemoteConfig: rmntypes.RemoteConfig{
				ContractAddress:  pbObs.MerkleRootObs.RmnRemoteConfig.ContractAddress,
				ConfigDigest:     cciptypes.Bytes32(pbObs.MerkleRootObs.RmnRemoteConfig.ConfigDigest),
				Signers:          rmnSigners,
				FSign:            pbObs.MerkleRootObs.RmnRemoteConfig.FSign,
				ConfigVersion:    pbObs.MerkleRootObs.RmnRemoteConfig.ConfigVersion,
				RmnReportVersion: cciptypes.Bytes32(pbObs.MerkleRootObs.RmnRemoteConfig.RmnReportVersion),
			},
			FChain: merkleRootsFChain,
		},
		TokenPriceObs: tokenprice.Observation{
			FeedTokenPrices:       feedTokenPrices,
			FeeQuoterTokenUpdates: feeQuoterTokenUpdates,
			FChain:                tokenPriceFChain,
			Timestamp:             pbObs.TokenPriceObs.Timestamp.AsTime(),
		},
		ChainFeeObs: chainfee.Observation{
			FeeComponents:     feeComponents,
			NativeTokenPrices: nativeTokenPrices,
			ChainFeeUpdates:   chainFeeUpdates,
			FChain:            chainFeeFChain,
			TimestampNow:      pbObs.ChainFeeObs.TimestampNow.AsTime(),
		},
		DiscoveryObs: discoverytypes.Observation{
			FChain:    discoveryFChain,
			Addresses: discoveryAddrs,
		},
		FChain: mainFChain,
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
