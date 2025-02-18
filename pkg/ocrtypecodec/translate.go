package ocrtypecodec

import (
	"math/big"

	"github.com/smartcontractkit/chainlink-common/pkg/types"
	rmnpb "github.com/smartcontractkit/chainlink-protos/rmn/v1.6/go/serialization"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/smartcontractkit/chainlink-ccip/commit/chainfee"
	"github.com/smartcontractkit/chainlink-ccip/commit/merkleroot/rmn"
	rmntypes "github.com/smartcontractkit/chainlink-ccip/commit/merkleroot/rmn/types"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugintypes"
	"github.com/smartcontractkit/chainlink-ccip/pkg/ocrtypecodec/ocrtypecodecpb"
	"github.com/smartcontractkit/chainlink-ccip/pkg/reader"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

type protoTranslator struct{}

func newProtoTranslator() *protoTranslator {
	return &protoTranslator{}
}

func (t *protoTranslator) rmnSignaturesToProto(sigs *rmn.ReportSignatures) []*ocrtypecodecpb.SignatureEcdsa {
	pbSigs := make([]*ocrtypecodecpb.SignatureEcdsa, len(sigs.Signatures))
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

func (t *protoTranslator) ccipRmnSignaturesToProto(sigs []cciptypes.RMNECDSASignature) []*ocrtypecodecpb.SignatureEcdsa {
	pbSigs := make([]*ocrtypecodecpb.SignatureEcdsa, len(sigs))
	for i, sig := range sigs {
		pbSigs[i] = &ocrtypecodecpb.SignatureEcdsa{
			R: sig.R[:],
			S: sig.S[:],
		}
	}
	return pbSigs
}

func (t *protoTranslator) ccipRmnSignaturesFromProto(pbSigs []*ocrtypecodecpb.SignatureEcdsa) []cciptypes.RMNECDSASignature {
	sigs := make([]cciptypes.RMNECDSASignature, len(pbSigs))
	for i := range pbSigs {
		sigs[i] = cciptypes.RMNECDSASignature{
			R: cciptypes.Bytes32(pbSigs[i].R),
			S: cciptypes.Bytes32(pbSigs[i].S),
		}
	}
	return sigs
}

func (t *protoTranslator) laneUpdatesToProto(rmnLaneUpdates []*rmnpb.FixedDestLaneUpdate) []*ocrtypecodecpb.DestChainUpdate {
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

func (t *protoTranslator) laneUpdatesFromProto(pbLaneUpdates []*ocrtypecodecpb.DestChainUpdate) []*rmnpb.FixedDestLaneUpdate {
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

func (t *protoTranslator) merkleRootsToProto(merkleRoots []cciptypes.MerkleRootChain) []*ocrtypecodecpb.MerkleRootChain {
	pbMerkleRoots := make([]*ocrtypecodecpb.MerkleRootChain, len(merkleRoots))

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

func (t *protoTranslator) merkleRootsFromProto(pbMerkleRoots []*ocrtypecodecpb.MerkleRootChain) []cciptypes.MerkleRootChain {
	merkleRoots := make([]cciptypes.MerkleRootChain, len(pbMerkleRoots))
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
	rmnEnabledChains := make(map[uint64]bool, len(rmnEnabled))
	for k, v := range rmnEnabled {
		rmnEnabledChains[uint64(k)] = v
	}
	return rmnEnabledChains
}

func (t *protoTranslator) rmnEnabledChainsFromProto(rmnEnabledChains map[uint64]bool) map[cciptypes.ChainSelector]bool {
	rmnEnabled := make(map[cciptypes.ChainSelector]bool, len(rmnEnabledChains))
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
	snc := make([]plugintypes.SeqNumChain, len(pbSnc))
	for i, s := range pbSnc {
		snc[i] = plugintypes.SeqNumChain{
			ChainSel: cciptypes.ChainSelector(s.ChainSel),
			SeqNum:   cciptypes.SeqNum(s.SeqNum),
		}
	}
	return snc
}

func (t *protoTranslator) rmnRemoteConfigToProto(rmnRemoteCfg rmntypes.RemoteConfig) *ocrtypecodecpb.RmnRemoteConfig {
	rmnRemoteConfigSignersPB := make([]*ocrtypecodecpb.RemoteSignerInfo, len(rmnRemoteCfg.Signers))
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

func (t *protoTranslator) rmnRemoteConfigFromProto(pbRmnRemoteCfg *ocrtypecodecpb.RmnRemoteConfig) rmntypes.RemoteConfig {
	rmnSigners := make([]rmntypes.RemoteSignerInfo, len(pbRmnRemoteCfg.Signers))

	for i, s := range pbRmnRemoteCfg.Signers {
		rmnSigners[i] = rmntypes.RemoteSignerInfo{
			OnchainPublicKey: s.OnchainPublicKey,
			NodeIndex:        s.NodeIndex,
		}
	}

	return rmntypes.RemoteConfig{
		ContractAddress:  pbRmnRemoteCfg.ContractAddress,
		ConfigDigest:     cciptypes.Bytes32(pbRmnRemoteCfg.ConfigDigest),
		Signers:          rmnSigners,
		FSign:            pbRmnRemoteCfg.FSign,
		ConfigVersion:    pbRmnRemoteCfg.ConfigVersion,
		RmnReportVersion: cciptypes.Bytes32(pbRmnRemoteCfg.RmnReportVersion),
	}
}

func (t *protoTranslator) fChainToProto(fChain map[cciptypes.ChainSelector]int) map[uint64]int32 {
	pbFChain := make(map[uint64]int32, len(fChain))
	for k, v := range fChain {
		pbFChain[uint64(k)] = int32(v)
	}
	return pbFChain
}

func (t *protoTranslator) fChainFromProto(pbFChain map[uint64]int32) map[cciptypes.ChainSelector]int {
	fChain := make(map[cciptypes.ChainSelector]int, len(pbFChain))
	for k, v := range pbFChain {
		fChain[cciptypes.ChainSelector(k)] = int(v)
	}
	return fChain
}

func (t *protoTranslator) feedTokenPricesToProto(feedPrices cciptypes.TokenPriceMap) map[string][]byte {
	feedTokenPrices := make(map[string][]byte, len(feedPrices))
	for k, v := range feedPrices {
		feedTokenPrices[string(k)] = v.Bytes()
	}
	return feedTokenPrices
}

func (t *protoTranslator) feedTokenPricesFromProto(pbFeedPrices map[string][]byte) cciptypes.TokenPriceMap {
	feedTokenPrices := make(cciptypes.TokenPriceMap, len(pbFeedPrices))
	for k, v := range pbFeedPrices {
		feedTokenPrices[cciptypes.UnknownEncodedAddress(k)] = cciptypes.NewBigInt(big.NewInt(0).SetBytes(v))
	}
	return feedTokenPrices
}

func (t *protoTranslator) feeQuoterTokenUpdatesToProto(
	tokenUpdates map[cciptypes.UnknownEncodedAddress]plugintypes.TimestampedBig,
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

func (t *protoTranslator) feeQuoterTokenUpdatesFromProto(pbTokenUpdates map[string]*ocrtypecodecpb.TimestampedBig) map[cciptypes.UnknownEncodedAddress]plugintypes.TimestampedBig {
	feeQuoterTokenUpdates := make(map[cciptypes.UnknownEncodedAddress]plugintypes.TimestampedBig, len(pbTokenUpdates))
	for k, v := range pbTokenUpdates {
		feeQuoterTokenUpdates[cciptypes.UnknownEncodedAddress(k)] = plugintypes.TimestampedBig{
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
	addrs reader.ContractAddresses,
) map[string]*ocrtypecodecpb.ChainAddressMap {
	pbAddrs := make(map[string]*ocrtypecodecpb.ChainAddressMap, len(addrs))

	for contractName, chains := range addrs {
		pbAddrs[contractName] = &ocrtypecodecpb.ChainAddressMap{
			ChainAddresses: make(map[uint64][]byte, len(chains)),
		}

		for chain, addr := range chains {
			pbAddrs[contractName].ChainAddresses[uint64(chain)] = addr
		}
	}

	return pbAddrs
}

func (t *protoTranslator) discoveryAddressesFromProto(
	pbAddrs map[string]*ocrtypecodecpb.ChainAddressMap,
) reader.ContractAddresses {
	discoveryAddrs := make(reader.ContractAddresses, len(pbAddrs))
	for contractName, chainMap := range pbAddrs {
		discoveryAddrs[contractName] = make(map[cciptypes.ChainSelector]cciptypes.UnknownAddress)
		for chain, addr := range chainMap.ChainAddresses {
			discoveryAddrs[contractName][cciptypes.ChainSelector(chain)] = addr
		}
	}
	return discoveryAddrs
}

func (t *protoTranslator) chainRangeToProto(chainRange []plugintypes.ChainRange) []*ocrtypecodecpb.ChainRange {
	rangesSelectedForReport := make([]*ocrtypecodecpb.ChainRange, len(chainRange))
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
	rangesSelectedForReport := make([]plugintypes.ChainRange, len(pbChainRange))
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
	gasPrices := make([]cciptypes.GasPriceChain, len(pbGpc))
	for i, gp := range pbGpc {
		gasPrices[i] = cciptypes.GasPriceChain{
			ChainSel: cciptypes.ChainSelector(gp.ChainSel),
			GasPrice: cciptypes.NewBigInt(big.NewInt(0).SetBytes(gp.GasPrice)),
		}
	}

	return gasPrices
}
