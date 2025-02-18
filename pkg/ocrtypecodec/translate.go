package ocrtypecodec

import (
	rmnpb "github.com/smartcontractkit/chainlink-protos/rmn/v1.6/go/serialization"

	"github.com/smartcontractkit/chainlink-ccip/commit/merkleroot/rmn"
	"github.com/smartcontractkit/chainlink-ccip/pkg/ocrtypecodec/ocrtypecodecpb"
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
