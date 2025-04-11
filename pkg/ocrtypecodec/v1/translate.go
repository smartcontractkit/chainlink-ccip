package v1

import (
	"math/big"

	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/smartcontractkit/chainlink-common/pkg/types"
	rmnpb "github.com/smartcontractkit/chainlink-protos/rmn/v1.6/go/serialization"

	"github.com/smartcontractkit/chainlink-ccip/commit/chainfee"
	"github.com/smartcontractkit/chainlink-ccip/commit/merkleroot/rmn"
	"github.com/smartcontractkit/chainlink-ccip/execute/exectypes"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugintypes"
	"github.com/smartcontractkit/chainlink-ccip/pkg/ocrtypecodec/v1/ocrtypecodecpb"
	"github.com/smartcontractkit/chainlink-ccip/pkg/reader"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

type protoTranslator struct{}

func newProtoTranslator() *protoTranslator {
	return &protoTranslator{}
}

func (t *protoTranslator) rmnSignaturesToProto(sigs *rmn.ReportSignatures) []*ocrtypecodecpb.SignatureEcdsa {
	var pbSigs []*ocrtypecodecpb.SignatureEcdsa
	if len(sigs.Signatures) > 0 {
		pbSigs = make([]*ocrtypecodecpb.SignatureEcdsa, len(sigs.Signatures))
	}

	for i, sig := range sigs.Signatures {
		pbSigs[i] = &ocrtypecodecpb.SignatureEcdsa{
			R: sig.R,
			S: sig.S,
		}
	}
	return pbSigs
}

func (t *protoTranslator) rmnSignaturesFromProto(pbSigs []*ocrtypecodecpb.SignatureEcdsa) []*rmnpb.EcdsaSignature {
	sigs := make([]*rmnpb.EcdsaSignature, len(pbSigs))
	for i := range pbSigs {
		sigs[i] = &rmnpb.EcdsaSignature{
			R: pbSigs[i].R,
			S: pbSigs[i].S,
		}
	}
	return sigs
}

func (t *protoTranslator) ccipRmnSignaturesToProto(
	sigs []cciptypes.RMNECDSASignature,
) []*ocrtypecodecpb.SignatureEcdsa {
	pbSigs := make([]*ocrtypecodecpb.SignatureEcdsa, len(sigs))
	for i, sig := range sigs {
		pbSigs[i] = &ocrtypecodecpb.SignatureEcdsa{
			R: sig.R[:],
			S: sig.S[:],
		}
	}
	return pbSigs
}

func (t *protoTranslator) ccipRmnSignaturesFromProto(
	pbSigs []*ocrtypecodecpb.SignatureEcdsa,
) []cciptypes.RMNECDSASignature {
	var sigs []cciptypes.RMNECDSASignature
	if len(pbSigs) > 0 {
		sigs = make([]cciptypes.RMNECDSASignature, len(pbSigs))
	}

	for i := range pbSigs {
		sigs[i] = cciptypes.RMNECDSASignature{
			R: cciptypes.Bytes32(pbSigs[i].R),
			S: cciptypes.Bytes32(pbSigs[i].S),
		}
	}
	return sigs
}

func (t *protoTranslator) laneUpdatesToProto(
	rmnLaneUpdates []*rmnpb.FixedDestLaneUpdate,
) []*ocrtypecodecpb.DestChainUpdate {
	pbLaneUpdates := make([]*ocrtypecodecpb.DestChainUpdate, len(rmnLaneUpdates))
	for i, lu := range rmnLaneUpdates {
		pbLaneUpdates[i] = &ocrtypecodecpb.DestChainUpdate{
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
	return pbLaneUpdates
}

func (t *protoTranslator) laneUpdatesFromProto(
	pbLaneUpdates []*ocrtypecodecpb.DestChainUpdate,
) []*rmnpb.FixedDestLaneUpdate {
	laneUpdates := make([]*rmnpb.FixedDestLaneUpdate, len(pbLaneUpdates))
	for i := range pbLaneUpdates {
		laneUpdates[i] = &rmnpb.FixedDestLaneUpdate{
			LaneSource: &rmnpb.LaneSource{
				SourceChainSelector: pbLaneUpdates[i].LaneSource.SourceChainSelector,
				OnrampAddress:       pbLaneUpdates[i].LaneSource.OnrampAddress,
			},
			ClosedInterval: &rmnpb.ClosedInterval{
				MinMsgNr: pbLaneUpdates[i].SeqNumRange.MinMsgNr,
				MaxMsgNr: pbLaneUpdates[i].SeqNumRange.MaxMsgNr,
			},
			Root: pbLaneUpdates[i].Root,
		}
	}
	return laneUpdates
}

func (t *protoTranslator) merkleRootsToProto(
	merkleRoots []cciptypes.MerkleRootChain,
) []*ocrtypecodecpb.MerkleRootChain {
	var pbMerkleRoots []*ocrtypecodecpb.MerkleRootChain
	if len(merkleRoots) > 0 {
		pbMerkleRoots = make([]*ocrtypecodecpb.MerkleRootChain, len(merkleRoots))
	}

	for i, mr := range merkleRoots {
		pbMerkleRoots[i] = &ocrtypecodecpb.MerkleRootChain{
			ChainSel:      uint64(mr.ChainSel),
			OnRampAddress: mr.OnRampAddress,
			SeqNumsRange: &ocrtypecodecpb.SeqNumRange{
				MinMsgNr: uint64(mr.SeqNumsRange.Start()),
				MaxMsgNr: uint64(mr.SeqNumsRange.End()),
			},
			MerkleRoot: mr.MerkleRoot[:],
		}
	}

	return pbMerkleRoots
}

func (t *protoTranslator) merkleRootsFromProto(
	pbMerkleRoots []*ocrtypecodecpb.MerkleRootChain,
) []cciptypes.MerkleRootChain {
	var merkleRoots []cciptypes.MerkleRootChain
	if len(pbMerkleRoots) > 0 {
		merkleRoots = make([]cciptypes.MerkleRootChain, len(pbMerkleRoots))
	}

	for i, mr := range pbMerkleRoots {
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

	return merkleRoots
}

func (t *protoTranslator) rmnEnabledChainsToProto(rmnEnabled map[cciptypes.ChainSelector]bool) map[uint64]bool {
	var rmnEnabledChains map[uint64]bool
	if len(rmnEnabled) > 0 {
		rmnEnabledChains = make(map[uint64]bool, len(rmnEnabled))
	}

	for k, v := range rmnEnabled {
		rmnEnabledChains[uint64(k)] = v
	}
	return rmnEnabledChains
}

func (t *protoTranslator) rmnEnabledChainsFromProto(rmnEnabledChains map[uint64]bool) map[cciptypes.ChainSelector]bool {
	var rmnEnabled map[cciptypes.ChainSelector]bool
	if len(rmnEnabledChains) > 0 {
		rmnEnabled = make(map[cciptypes.ChainSelector]bool, len(rmnEnabledChains))
	}

	for k, v := range rmnEnabledChains {
		rmnEnabled[cciptypes.ChainSelector(k)] = v
	}
	return rmnEnabled
}

func (t *protoTranslator) seqNumChainToProto(snc []plugintypes.SeqNumChain) []*ocrtypecodecpb.SeqNumChain {
	pbSnc := make([]*ocrtypecodecpb.SeqNumChain, len(snc))
	for i, s := range snc {
		pbSnc[i] = &ocrtypecodecpb.SeqNumChain{
			ChainSel: uint64(s.ChainSel),
			SeqNum:   uint64(s.SeqNum),
		}
	}
	return pbSnc
}

func (t *protoTranslator) seqNumChainFromProto(pbSnc []*ocrtypecodecpb.SeqNumChain) []plugintypes.SeqNumChain {
	var snc []plugintypes.SeqNumChain
	if len(pbSnc) > 0 {
		snc = make([]plugintypes.SeqNumChain, len(pbSnc))
	}

	for i, s := range pbSnc {
		snc[i] = plugintypes.SeqNumChain{
			ChainSel: cciptypes.ChainSelector(s.ChainSel),
			SeqNum:   cciptypes.SeqNum(s.SeqNum),
		}
	}
	return snc
}

func (t *protoTranslator) rmnRemoteConfigToProto(rmnRemoteCfg cciptypes.RemoteConfig) *ocrtypecodecpb.RmnRemoteConfig {
	var rmnRemoteConfigSignersPB []*ocrtypecodecpb.RemoteSignerInfo
	if len(rmnRemoteCfg.Signers) > 0 {
		rmnRemoteConfigSignersPB = make([]*ocrtypecodecpb.RemoteSignerInfo, len(rmnRemoteCfg.Signers))
	}

	for i, s := range rmnRemoteCfg.Signers {
		rmnRemoteConfigSignersPB[i] = &ocrtypecodecpb.RemoteSignerInfo{
			OnchainPublicKey: s.OnchainPublicKey,
			NodeIndex:        s.NodeIndex,
		}
	}

	return &ocrtypecodecpb.RmnRemoteConfig{
		ContractAddress:  rmnRemoteCfg.ContractAddress,
		ConfigDigest:     rmnRemoteCfg.ConfigDigest[:],
		Signers:          rmnRemoteConfigSignersPB,
		FSign:            rmnRemoteCfg.FSign,
		ConfigVersion:    rmnRemoteCfg.ConfigVersion,
		RmnReportVersion: rmnRemoteCfg.RmnReportVersion[:],
	}
}

func (t *protoTranslator) rmnRemoteConfigFromProto(
	pbRmnRemoteCfg *ocrtypecodecpb.RmnRemoteConfig,
) cciptypes.RemoteConfig {
	var rmnSigners []cciptypes.RemoteSignerInfo
	if len(pbRmnRemoteCfg.Signers) > 0 {
		rmnSigners = make([]cciptypes.RemoteSignerInfo, len(pbRmnRemoteCfg.Signers))
	}
	for i, s := range pbRmnRemoteCfg.Signers {
		rmnSigners[i] = cciptypes.RemoteSignerInfo{
			OnchainPublicKey: s.OnchainPublicKey,
			NodeIndex:        s.NodeIndex,
		}
	}

	return cciptypes.RemoteConfig{
		ContractAddress:  pbRmnRemoteCfg.ContractAddress,
		ConfigDigest:     cciptypes.Bytes32(pbRmnRemoteCfg.ConfigDigest),
		Signers:          rmnSigners,
		FSign:            pbRmnRemoteCfg.FSign,
		ConfigVersion:    pbRmnRemoteCfg.ConfigVersion,
		RmnReportVersion: cciptypes.Bytes32(pbRmnRemoteCfg.RmnReportVersion),
	}
}

func (t *protoTranslator) fChainToProto(fChain map[cciptypes.ChainSelector]int) map[uint64]int32 {
	var pbFChain map[uint64]int32
	if len(fChain) > 0 {
		pbFChain = make(map[uint64]int32, len(fChain))
	}

	for k, v := range fChain {
		pbFChain[uint64(k)] = int32(v)
	}
	return pbFChain
}

func (t *protoTranslator) fChainFromProto(pbFChain map[uint64]int32) map[cciptypes.ChainSelector]int {
	var fChain map[cciptypes.ChainSelector]int
	if len(pbFChain) > 0 {
		fChain = make(map[cciptypes.ChainSelector]int, len(pbFChain))
	}

	for k, v := range pbFChain {
		fChain[cciptypes.ChainSelector(k)] = int(v)
	}
	return fChain
}

func (t *protoTranslator) feedTokenPricesToProto(feedPrices cciptypes.TokenPriceMap) map[string][]byte {
	var feedTokenPrices map[string][]byte
	if len(feedPrices) > 0 {
		feedTokenPrices = make(map[string][]byte, len(feedPrices))
	}

	for k, v := range feedPrices {
		feedTokenPrices[string(k)] = v.Bytes()
	}
	return feedTokenPrices
}

func (t *protoTranslator) feedTokenPricesFromProto(pbFeedPrices map[string][]byte) cciptypes.TokenPriceMap {
	var feedTokenPrices cciptypes.TokenPriceMap
	if len(pbFeedPrices) > 0 {
		feedTokenPrices = make(cciptypes.TokenPriceMap, len(pbFeedPrices))
	}

	for k, v := range pbFeedPrices {
		feedTokenPrices[cciptypes.UnknownEncodedAddress(k)] = cciptypes.NewBigInt(big.NewInt(0).SetBytes(v))
	}
	return feedTokenPrices
}

func (t *protoTranslator) feeQuoterTokenUpdatesToProto(
	tokenUpdates map[cciptypes.UnknownEncodedAddress]cciptypes.TimestampedBig,
) map[string]*ocrtypecodecpb.TimestampedBig {
	feeQuoterTokenUpdates := make(map[string]*ocrtypecodecpb.TimestampedBig, len(tokenUpdates))

	for k, v := range tokenUpdates {
		feeQuoterTokenUpdates[string(k)] = &ocrtypecodecpb.TimestampedBig{
			Value:     v.Value.Bytes(),
			Timestamp: timestamppb.New(v.Timestamp),
		}
	}

	return feeQuoterTokenUpdates
}

func (t *protoTranslator) feeQuoterTokenUpdatesFromProto(
	pbTokenUpdates map[string]*ocrtypecodecpb.TimestampedBig,
) map[cciptypes.UnknownEncodedAddress]cciptypes.TimestampedBig {
	feeQuoterTokenUpdates := make(map[cciptypes.UnknownEncodedAddress]cciptypes.TimestampedBig, len(pbTokenUpdates))
	for k, v := range pbTokenUpdates {
		feeQuoterTokenUpdates[cciptypes.UnknownEncodedAddress(k)] = cciptypes.TimestampedBig{
			Value:     cciptypes.NewBigInt(big.NewInt(0).SetBytes(v.Value)),
			Timestamp: v.Timestamp.AsTime(),
		}
	}
	return feeQuoterTokenUpdates
}

func (t *protoTranslator) feeComponentsToProto(
	feeComponents map[cciptypes.ChainSelector]types.ChainFeeComponents,
) map[uint64]*ocrtypecodecpb.ChainFeeComponents {
	pbFeeComponents := make(map[uint64]*ocrtypecodecpb.ChainFeeComponents, len(feeComponents))
	for k, v := range feeComponents {
		pbFeeComponents[uint64(k)] = &ocrtypecodecpb.ChainFeeComponents{
			ExecutionFee:        v.ExecutionFee.Bytes(),
			DataAvailabilityFee: v.DataAvailabilityFee.Bytes(),
		}
	}
	return pbFeeComponents
}

func (t *protoTranslator) feeComponentsFromProto(
	pbFeeComponents map[uint64]*ocrtypecodecpb.ChainFeeComponents,
) map[cciptypes.ChainSelector]types.ChainFeeComponents {
	feeComponents := make(map[cciptypes.ChainSelector]types.ChainFeeComponents, len(pbFeeComponents))
	for k, v := range pbFeeComponents {
		feeComponents[cciptypes.ChainSelector(k)] = types.ChainFeeComponents{
			ExecutionFee:        big.NewInt(0).SetBytes(v.ExecutionFee),
			DataAvailabilityFee: big.NewInt(0).SetBytes(v.DataAvailabilityFee),
		}
	}
	return feeComponents
}

func (t *protoTranslator) nativeTokenPricesToProto(
	prices map[cciptypes.ChainSelector]cciptypes.BigInt,
) map[uint64][]byte {
	pbPrices := make(map[uint64][]byte, len(prices))
	for k, v := range prices {
		pbPrices[uint64(k)] = v.Bytes()
	}
	return pbPrices
}

func (t *protoTranslator) nativeTokenPricesFromProto(
	pbPrices map[uint64][]byte,
) map[cciptypes.ChainSelector]cciptypes.BigInt {
	prices := make(map[cciptypes.ChainSelector]cciptypes.BigInt, len(pbPrices))
	for k, v := range pbPrices {
		prices[cciptypes.ChainSelector(k)] = cciptypes.NewBigInt(big.NewInt(0).SetBytes(v))
	}
	return prices
}

func (t *protoTranslator) chainFeeUpdatesToProto(
	updates map[cciptypes.ChainSelector]chainfee.Update,
) map[uint64]*ocrtypecodecpb.ChainFeeUpdate {
	chainFeeUpdates := make(map[uint64]*ocrtypecodecpb.ChainFeeUpdate, len(updates))

	for k, v := range updates {
		chainFeeUpdates[uint64(k)] = &ocrtypecodecpb.ChainFeeUpdate{
			ChainFee: &ocrtypecodecpb.ComponentsUSDPrices{
				ExecutionFeePriceUsd: v.ChainFee.ExecutionFeePriceUSD.Bytes(),
				DataAvFeePriceUsd:    v.ChainFee.DataAvFeePriceUSD.Bytes(),
			},
			Timestamp: timestamppb.New(v.Timestamp),
		}
	}

	return chainFeeUpdates
}

func (t *protoTranslator) chainFeeUpdatesFromProto(
	pbUpdates map[uint64]*ocrtypecodecpb.ChainFeeUpdate,
) map[cciptypes.ChainSelector]chainfee.Update {
	chainFeeUpdates := make(map[cciptypes.ChainSelector]chainfee.Update, len(pbUpdates))
	for k, v := range pbUpdates {
		chainFeeUpdates[cciptypes.ChainSelector(k)] = chainfee.Update{
			ChainFee: chainfee.ComponentsUSDPrices{
				ExecutionFeePriceUSD: big.NewInt(0).SetBytes(v.ChainFee.ExecutionFeePriceUsd),
				DataAvFeePriceUSD:    big.NewInt(0).SetBytes(v.ChainFee.DataAvFeePriceUsd),
			},
			Timestamp: v.Timestamp.AsTime(),
		}
	}
	return chainFeeUpdates
}

func (t *protoTranslator) discoveryAddressesToProto(
	addresses reader.ContractAddresses,
) map[string]*ocrtypecodecpb.ChainAddressMap {
	var pbAddresses map[string]*ocrtypecodecpb.ChainAddressMap
	if len(addresses) > 0 {
		pbAddresses = make(map[string]*ocrtypecodecpb.ChainAddressMap, len(addresses))
	}

	for contractName, chains := range addresses {
		var chainAddresses map[uint64][]byte
		if len(chains) > 0 {
			chainAddresses = make(map[uint64][]byte, len(chains))
		}
		pbAddresses[contractName] = &ocrtypecodecpb.ChainAddressMap{
			ChainAddresses: chainAddresses,
		}

		for chain, addr := range chains {
			pbAddresses[contractName].ChainAddresses[uint64(chain)] = addr
		}
	}

	return pbAddresses
}

func (t *protoTranslator) discoveryAddressesFromProto(
	pbAddresses map[string]*ocrtypecodecpb.ChainAddressMap,
) reader.ContractAddresses {
	var discoveryAddresses reader.ContractAddresses
	if len(pbAddresses) > 0 {
		discoveryAddresses = make(reader.ContractAddresses, len(pbAddresses))
	}

	for contractName, chainMap := range pbAddresses {
		discoveryAddresses[contractName] = make(map[cciptypes.ChainSelector]cciptypes.UnknownAddress)
		for chain, addr := range chainMap.ChainAddresses {
			discoveryAddresses[contractName][cciptypes.ChainSelector(chain)] = addr
		}
	}

	return discoveryAddresses
}

func (t *protoTranslator) chainRangeToProto(chainRange []plugintypes.ChainRange) []*ocrtypecodecpb.ChainRange {
	var rangesSelectedForReport []*ocrtypecodecpb.ChainRange
	if len(chainRange) > 0 {
		rangesSelectedForReport = make([]*ocrtypecodecpb.ChainRange, len(chainRange))
	}

	for i, r := range chainRange {
		rangesSelectedForReport[i] = &ocrtypecodecpb.ChainRange{
			ChainSel: uint64(r.ChainSel),
			SeqNumRange: &ocrtypecodecpb.SeqNumRange{
				MinMsgNr: uint64(r.SeqNumRange.Start()),
				MaxMsgNr: uint64(r.SeqNumRange.End()),
			},
		}
	}
	return rangesSelectedForReport
}

func (t *protoTranslator) chainRangeFromProto(pbChainRange []*ocrtypecodecpb.ChainRange) []plugintypes.ChainRange {
	var rangesSelectedForReport []plugintypes.ChainRange
	if len(pbChainRange) > 0 {
		rangesSelectedForReport = make([]plugintypes.ChainRange, len(pbChainRange))
	}

	for i, r := range pbChainRange {
		rangesSelectedForReport[i] = plugintypes.ChainRange{
			ChainSel: cciptypes.ChainSelector(r.ChainSel),
			SeqNumRange: cciptypes.NewSeqNumRange(
				cciptypes.SeqNum(r.SeqNumRange.MinMsgNr),
				cciptypes.SeqNum(r.SeqNumRange.MaxMsgNr),
			),
		}
	}
	return rangesSelectedForReport
}

func (t *protoTranslator) gasPriceChainToProto(gpc []cciptypes.GasPriceChain) []*ocrtypecodecpb.GasPriceChain {
	gasPrices := make([]*ocrtypecodecpb.GasPriceChain, len(gpc))
	for i, gp := range gpc {
		gasPrices[i] = &ocrtypecodecpb.GasPriceChain{
			ChainSel: uint64(gp.ChainSel),
			GasPrice: gp.GasPrice.Bytes(),
		}
	}

	return gasPrices
}

func (t *protoTranslator) gasPriceChainFromProto(pbGpc []*ocrtypecodecpb.GasPriceChain) []cciptypes.GasPriceChain {
	var gasPrices []cciptypes.GasPriceChain
	if len(pbGpc) > 0 {
		gasPrices = make([]cciptypes.GasPriceChain, len(pbGpc))
	}

	for i, gp := range pbGpc {
		gasPrices[i] = cciptypes.GasPriceChain{
			ChainSel: cciptypes.ChainSelector(gp.ChainSel),
			GasPrice: cciptypes.NewBigInt(big.NewInt(0).SetBytes(gp.GasPrice)),
		}
	}

	return gasPrices
}

func (t *protoTranslator) commitReportsToProto(
	observations exectypes.CommitObservations,
) map[uint64]*ocrtypecodecpb.CommitObservations {
	var commitReports map[uint64]*ocrtypecodecpb.CommitObservations
	if len(observations) > 0 {
		commitReports = make(map[uint64]*ocrtypecodecpb.CommitObservations, len(observations))
	}

	for chainSel, commits := range observations {
		commitReports[uint64(chainSel)] = &ocrtypecodecpb.CommitObservations{
			CommitData: t.commitDataSliceToProto(commits),
		}
	}

	return commitReports
}

func (t *protoTranslator) commitDataSliceToProto(commits []exectypes.CommitData) []*ocrtypecodecpb.CommitData {
	commitData := make([]*ocrtypecodecpb.CommitData, len(commits))

	for i, commit := range commits {
		commitData[i] = &ocrtypecodecpb.CommitData{
			SourceChain:   uint64(commit.SourceChain),
			OnRampAddress: commit.OnRampAddress,
			Timestamp:     timestamppb.New(commit.Timestamp),
			BlockNum:      commit.BlockNum,
			MerkleRoot:    commit.MerkleRoot[:],
			SequenceNumberRange: &ocrtypecodecpb.SeqNumRange{
				MinMsgNr: uint64(commit.SequenceNumberRange.Start()),
				MaxMsgNr: uint64(commit.SequenceNumberRange.End()),
			},
			ExecutedMessages: t.seqNumsToProto(commit.ExecutedMessages),
			Messages:         t.messagesToProto(commit.Messages),
			Hashes:           t.bytes32SliceToProto(commit.Hashes),
			MessageTokenData: t.messageTokenDataSliceToProto(commit.MessageTokenData),
		}
	}

	return commitData
}

func (t *protoTranslator) commitReportsFromProto(
	pbObservations map[uint64]*ocrtypecodecpb.CommitObservations,
) exectypes.CommitObservations {
	var commitReports exectypes.CommitObservations
	if len(pbObservations) > 0 {
		commitReports = make(exectypes.CommitObservations, len(pbObservations))
	}

	for chainSel, commitObs := range pbObservations {
		commitReports[cciptypes.ChainSelector(chainSel)] = t.commitDataSliceFromProto(commitObs.CommitData)
	}

	return commitReports
}

func (t *protoTranslator) commitDataSliceFromProto(pbCommits []*ocrtypecodecpb.CommitData) []exectypes.CommitData {
	commitData := make([]exectypes.CommitData, len(pbCommits))

	for i, commit := range pbCommits {
		commitData[i] = exectypes.CommitData{
			SourceChain:   cciptypes.ChainSelector(commit.SourceChain),
			OnRampAddress: commit.OnRampAddress,
			Timestamp:     commit.Timestamp.AsTime(),
			BlockNum:      commit.BlockNum,
			MerkleRoot:    cciptypes.Bytes32(commit.MerkleRoot),
			SequenceNumberRange: cciptypes.NewSeqNumRange(
				cciptypes.SeqNum(commit.SequenceNumberRange.MinMsgNr),
				cciptypes.SeqNum(commit.SequenceNumberRange.MaxMsgNr),
			),
			ExecutedMessages: t.decodeSeqNums(commit.ExecutedMessages),
			Messages:         t.decodeMessages(commit.Messages),
			Hashes:           t.bytes32SliceFromProto(commit.Hashes),
			MessageTokenData: t.decodeMessageTokenData(commit.MessageTokenData),
		}
	}

	return commitData
}

func (t *protoTranslator) messagesToProto(messages []cciptypes.Message) []*ocrtypecodecpb.Message {
	var pbMessages []*ocrtypecodecpb.Message
	if len(messages) > 0 {
		pbMessages = make([]*ocrtypecodecpb.Message, len(messages))
	}

	for i, msg := range messages {
		pbMessages[i] = t.encodeMessage(msg)
	}
	return pbMessages
}

func (t *protoTranslator) encodeMessage(msg cciptypes.Message) *ocrtypecodecpb.Message {
	return &ocrtypecodecpb.Message{
		Header: &ocrtypecodecpb.RampMessageHeader{
			MessageId:           msg.Header.MessageID[:],
			SourceChainSelector: uint64(msg.Header.SourceChainSelector),
			DestChainSelector:   uint64(msg.Header.DestChainSelector),
			SequenceNumber:      uint64(msg.Header.SequenceNumber),
			Nonce:               msg.Header.Nonce,
			MsgHash:             msg.Header.MsgHash[:],
			OnRamp:              msg.Header.OnRamp,
		},
		Sender:         msg.Sender,
		Data:           msg.Data,
		Receiver:       msg.Receiver,
		ExtraArgs:      msg.ExtraArgs,
		FeeToken:       msg.FeeToken,
		FeeTokenAmount: msg.FeeTokenAmount.Bytes(),
		FeeValueJuels:  msg.FeeValueJuels.Bytes(),
		TokenAmounts:   t.encodeRampTokenAmounts(msg.TokenAmounts),
	}
}

func (t *protoTranslator) encodeRampTokenAmounts(
	tokenAmounts []cciptypes.RampTokenAmount,
) []*ocrtypecodecpb.RampTokenAmount {
	var result []*ocrtypecodecpb.RampTokenAmount
	if len(tokenAmounts) > 0 {
		result = make([]*ocrtypecodecpb.RampTokenAmount, len(tokenAmounts))
	}

	for i, token := range tokenAmounts {
		result[i] = &ocrtypecodecpb.RampTokenAmount{
			SourcePoolAddress: token.SourcePoolAddress,
			DestTokenAddress:  token.DestTokenAddress,
			ExtraData:         token.ExtraData,
			Amount:            token.Amount.Bytes(),
			DestExecData:      token.DestExecData,
		}
	}

	return result
}

func (t *protoTranslator) seqNumsToProto(seqNums []cciptypes.SeqNum) []uint64 {
	var result []uint64
	if len(seqNums) > 0 {
		result = make([]uint64, len(seqNums))
	}

	for i, num := range seqNums {
		result[i] = uint64(num)
	}
	return result
}

func (t *protoTranslator) bytes32SliceToProto(slice []cciptypes.Bytes32) [][]byte {
	var result [][]byte
	if len(slice) > 0 {
		result = make([][]byte, len(slice))
	}
	for i, val := range slice {
		result[i] = val[:]
	}
	return result
}

func (t *protoTranslator) bytes32SliceFromProto(pbSlice [][]byte) []cciptypes.Bytes32 {
	var result []cciptypes.Bytes32
	if len(pbSlice) > 0 {
		result = make([]cciptypes.Bytes32, len(pbSlice))
	}

	for i, val := range pbSlice {
		result[i] = cciptypes.Bytes32(val)
	}

	return result
}

func (t *protoTranslator) messageTokenDataSliceToProto(
	data []exectypes.MessageTokenData,
) []*ocrtypecodecpb.MessageTokenData {
	var result []*ocrtypecodecpb.MessageTokenData
	if len(data) > 0 {
		result = make([]*ocrtypecodecpb.MessageTokenData, len(data))
	}

	for i, item := range data {
		result[i] = t.messageTokenDataToProto(item)
	}
	return result
}

func (t *protoTranslator) messageTokenDataToProto(data exectypes.MessageTokenData) *ocrtypecodecpb.MessageTokenData {
	tokenData := make([]*ocrtypecodecpb.TokenData, len(data.TokenData))
	for i, td := range data.TokenData {
		tokenData[i] = &ocrtypecodecpb.TokenData{
			Ready: td.Ready,
			Data:  td.Data,
		}
	}
	return &ocrtypecodecpb.MessageTokenData{
		TokenData: tokenData,
	}
}

func (t *protoTranslator) messageObservationsToProto(
	msgs exectypes.MessageObservations,
) map[uint64]*ocrtypecodecpb.SeqNumToMessage {
	var pbMsgs map[uint64]*ocrtypecodecpb.SeqNumToMessage
	if len(msgs) > 0 {
		pbMsgs = make(map[uint64]*ocrtypecodecpb.SeqNumToMessage, len(msgs))
	}

	for chainSel, seqNums := range msgs {
		seqNumToMsg := make(map[uint64]*ocrtypecodecpb.Message, len(seqNums))
		for seqNum, msg := range seqNums {
			seqNumToMsg[uint64(seqNum)] = t.encodeMessage(msg)
		}

		pbMsgs[uint64(chainSel)] = &ocrtypecodecpb.SeqNumToMessage{
			Messages: seqNumToMsg,
		}
	}

	return pbMsgs
}

func (t *protoTranslator) messageObservationsFromProto(
	pbMsgs map[uint64]*ocrtypecodecpb.SeqNumToMessage,
) exectypes.MessageObservations {
	var messages exectypes.MessageObservations
	if len(pbMsgs) > 0 {
		messages = make(exectypes.MessageObservations, len(pbMsgs))
	}

	for chainSel, msgMap := range pbMsgs {
		innerMap := make(map[cciptypes.SeqNum]cciptypes.Message, len(msgMap.Messages))
		for seqNum, msg := range msgMap.Messages {
			innerMap[cciptypes.SeqNum(seqNum)] = t.decodeMessage(msg)
		}
		messages[cciptypes.ChainSelector(chainSel)] = innerMap
	}

	return messages
}

func (t *protoTranslator) messageHashesToProto(
	hashes exectypes.MessageHashes,
) map[uint64]*ocrtypecodecpb.SeqNumToBytes {
	var messageHashes map[uint64]*ocrtypecodecpb.SeqNumToBytes
	if len(hashes) > 0 {
		messageHashes = make(map[uint64]*ocrtypecodecpb.SeqNumToBytes, len(hashes))
	}

	for chainSel, hashMap := range hashes {
		seqNumToBytes := &ocrtypecodecpb.SeqNumToBytes{SeqNumToBytes: make(map[uint64][]byte, len(hashMap))}
		for seqNum, hash := range hashMap {
			seqNumToBytes.SeqNumToBytes[uint64(seqNum)] = hash[:]
		}
		messageHashes[uint64(chainSel)] = seqNumToBytes
	}

	return messageHashes
}

func (t *protoTranslator) messageHashesFromProto(
	pbHashes map[uint64]*ocrtypecodecpb.SeqNumToBytes,
) exectypes.MessageHashes {
	var hashes exectypes.MessageHashes
	if len(pbHashes) > 0 {
		hashes = make(exectypes.MessageHashes, len(pbHashes))
	}

	for chainSel, hashMap := range pbHashes {
		innerMap := make(map[cciptypes.SeqNum]cciptypes.Bytes32, len(hashMap.SeqNumToBytes))
		for seqNum, hash := range hashMap.SeqNumToBytes {
			innerMap[cciptypes.SeqNum(seqNum)] = cciptypes.Bytes32(hash)
		}
		hashes[cciptypes.ChainSelector(chainSel)] = innerMap
	}

	return hashes
}

func (t *protoTranslator) tokenDataObservationsToProto(
	observations exectypes.TokenDataObservations,
) map[uint64]*ocrtypecodecpb.SeqNumToTokenData {
	var tokenDataObservations map[uint64]*ocrtypecodecpb.SeqNumToTokenData
	if len(observations) > 0 {
		tokenDataObservations = make(map[uint64]*ocrtypecodecpb.SeqNumToTokenData, len(observations))
	}

	for chainSel, tokenMap := range observations {
		seqNumToTokenData := &ocrtypecodecpb.SeqNumToTokenData{
			TokenData: make(map[uint64]*ocrtypecodecpb.MessageTokenData),
		}
		for seqNum, tokenData := range tokenMap {
			seqNumToTokenData.TokenData[uint64(seqNum)] = t.messageTokenDataToProto(tokenData)
		}
		tokenDataObservations[uint64(chainSel)] = seqNumToTokenData
	}

	return tokenDataObservations
}

func (t *protoTranslator) tokenDataObservationsFromProto(
	pbObservations map[uint64]*ocrtypecodecpb.SeqNumToTokenData,
) exectypes.TokenDataObservations {
	var tokenDataObservations exectypes.TokenDataObservations
	if len(pbObservations) > 0 {
		tokenDataObservations = make(exectypes.TokenDataObservations, len(pbObservations))
	}

	for chainSel, tokenMap := range pbObservations {
		innerMap := make(map[cciptypes.SeqNum]exectypes.MessageTokenData, len(tokenMap.TokenData))
		for seqNum, tokenData := range tokenMap.TokenData {
			innerMap[cciptypes.SeqNum(seqNum)] = t.decodeMessageTokenDataEntry(tokenData)
		}
		tokenDataObservations[cciptypes.ChainSelector(chainSel)] = innerMap
	}

	return tokenDataObservations
}

func (t *protoTranslator) nonceObservationsToProto(
	observations exectypes.NonceObservations,
) map[uint64]*ocrtypecodecpb.StringAddrToNonce {
	var nonceObservations map[uint64]*ocrtypecodecpb.StringAddrToNonce
	if len(observations) > 0 {
		nonceObservations = make(map[uint64]*ocrtypecodecpb.StringAddrToNonce, len(observations))
	}

	for chainSel, nonceMap := range observations {
		addrToNonce := make(map[string]uint64, len(nonceMap))
		for addr, nonce := range nonceMap {
			addrToNonce[addr] = nonce
		}
		nonceObservations[uint64(chainSel)] = &ocrtypecodecpb.StringAddrToNonce{Nonces: addrToNonce}
	}

	return nonceObservations
}

func (t *protoTranslator) nonceObservationsFromProto(
	pbObservations map[uint64]*ocrtypecodecpb.StringAddrToNonce,
) exectypes.NonceObservations {
	var nonces exectypes.NonceObservations
	if len(pbObservations) > 0 {
		nonces = make(exectypes.NonceObservations, len(pbObservations))
	}

	for chainSel, nonceMap := range pbObservations {
		innerMap := make(map[string]uint64, len(nonceMap.Nonces))
		for addr, nonce := range nonceMap.Nonces {
			innerMap[addr] = nonce
		}
		nonces[cciptypes.ChainSelector(chainSel)] = innerMap
	}

	return nonces
}

func (t *protoTranslator) chainReportsToProto(
	reports []cciptypes.ExecutePluginReportSingleChain,
) []*ocrtypecodecpb.ChainReport {
	pbReports := make([]*ocrtypecodecpb.ChainReport, len(reports))

	for i, r := range reports {
		offchainTokenData := make([]*ocrtypecodecpb.RepeatedBytes, 0)
		for _, data := range r.OffchainTokenData {
			offchainTokenData = append(offchainTokenData, &ocrtypecodecpb.RepeatedBytes{Items: data})
		}

		pbReports[i] = &ocrtypecodecpb.ChainReport{
			SourceChainSelector: uint64(r.SourceChainSelector),
			Messages:            t.messagesToProto(r.Messages),
			OffchainTokenData:   offchainTokenData,
			Proofs:              t.bytes32SliceToProto(r.Proofs),
			ProofFlagBits:       r.ProofFlagBits.Bytes(),
		}
	}
	return pbReports
}

func (t *protoTranslator) chainReportsFromProto(
	pbReports []*ocrtypecodecpb.ChainReport,
) []cciptypes.ExecutePluginReportSingleChain {
	reports := make([]cciptypes.ExecutePluginReportSingleChain, len(pbReports))

	for i, r := range pbReports {
		offchainTokenData := make([][][]byte, 0)
		for _, data := range r.OffchainTokenData {
			offchainTokenData = append(offchainTokenData, data.Items)
		}

		reports[i] = cciptypes.ExecutePluginReportSingleChain{
			SourceChainSelector: cciptypes.ChainSelector(r.SourceChainSelector),
			Messages:            t.decodeMessages(r.Messages),
			OffchainTokenData:   offchainTokenData,
			Proofs:              t.bytes32SliceFromProto(r.Proofs),
			ProofFlagBits:       cciptypes.NewBigInt(big.NewInt(0).SetBytes(r.ProofFlagBits)),
		}
	}

	return reports
}

func (t *protoTranslator) decodeMessageTokenData(data []*ocrtypecodecpb.MessageTokenData) []exectypes.MessageTokenData {
	var result []exectypes.MessageTokenData
	if len(data) > 0 {
		result = make([]exectypes.MessageTokenData, len(data))
	}

	for i, item := range data {
		result[i] = t.decodeMessageTokenDataEntry(item)
	}
	return result
}

func (t *protoTranslator) decodeSeqNums(seqNums []uint64) []cciptypes.SeqNum {
	var result []cciptypes.SeqNum
	if len(seqNums) > 0 {
		result = make([]cciptypes.SeqNum, len(seqNums))
	}

	for i, num := range seqNums {
		result[i] = cciptypes.SeqNum(num)
	}
	return result
}

func (t *protoTranslator) decodeMessages(messages []*ocrtypecodecpb.Message) []cciptypes.Message {
	var result []cciptypes.Message
	if len(messages) > 0 {
		result = make([]cciptypes.Message, len(messages))
	}

	for i, msg := range messages {
		result[i] = t.decodeMessage(msg)
	}
	return result
}

func (t *protoTranslator) decodeMessage(msg *ocrtypecodecpb.Message) cciptypes.Message {
	return cciptypes.Message{
		Header:         t.decodeMessageHeader(msg.Header),
		Sender:         msg.Sender,
		Data:           msg.Data,
		Receiver:       msg.Receiver,
		ExtraArgs:      msg.ExtraArgs,
		FeeToken:       msg.FeeToken,
		FeeTokenAmount: cciptypes.NewBigInt(big.NewInt(0).SetBytes(msg.FeeTokenAmount)),
		FeeValueJuels:  cciptypes.NewBigInt(big.NewInt(0).SetBytes(msg.FeeValueJuels)),
		TokenAmounts:   t.decodeRampTokenAmounts(msg.TokenAmounts),
	}
}

func (t *protoTranslator) decodeMessageHeader(header *ocrtypecodecpb.RampMessageHeader) cciptypes.RampMessageHeader {
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

func (t *protoTranslator) decodeRampTokenAmounts(
	tokenAmounts []*ocrtypecodecpb.RampTokenAmount,
) []cciptypes.RampTokenAmount {
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

func (t *protoTranslator) decodeMessageTokenDataEntry(
	data *ocrtypecodecpb.MessageTokenData,
) exectypes.MessageTokenData {
	tokenData := make([]exectypes.TokenData, len(data.TokenData))
	for i, td := range data.TokenData {
		tokenData[i] = exectypes.TokenData{
			Ready: td.Ready,
			Data:  td.Data,
		}
	}
	return exectypes.MessageTokenData{TokenData: tokenData}
}
