package rmn

import (
	"fmt"

	rmnpb "github.com/smartcontractkit/chainlink-protos/rmn/v1.6/go/serialization"

	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

// NewLaneUpdatesFromPB converts a slice of pb FixedDestLaneUpdate to a slice of RMNLaneUpdate
func NewLaneUpdatesFromPB(pbLaneUpdates []*rmnpb.FixedDestLaneUpdate) ([]cciptypes.RMNLaneUpdate, error) {
	laneUpdates := make([]cciptypes.RMNLaneUpdate, 0, len(pbLaneUpdates))

	for _, lu := range pbLaneUpdates {
		if len(lu.Root) != 32 {
			return nil, fmt.Errorf("invalid merkle root, must be 32 bytes: %v", lu.Root)
		}
		root32b := [32]byte{}
		copy(root32b[:], lu.Root[:32])

		laneUpdates = append(laneUpdates, cciptypes.RMNLaneUpdate{
			SourceChainSelector: cciptypes.ChainSelector(lu.LaneSource.SourceChainSelector),
			OnRampAddress:       lu.LaneSource.OnrampAddress,
			MinSeqNr:            cciptypes.SeqNum(lu.ClosedInterval.MinMsgNr),
			MaxSeqNr:            cciptypes.SeqNum(lu.ClosedInterval.MaxMsgNr),
			MerkleRoot:          root32b,
		})
	}

	return laneUpdates, nil
}

// NewECDSASigsFromPB converts a slice of pb EcdsaSignature to a slice of RMNECDSASignature
func NewECDSASigsFromPB(pbSigs []*rmnpb.EcdsaSignature) ([]cciptypes.RMNECDSASignature, error) {
	sigs := make([]cciptypes.RMNECDSASignature, 0, len(pbSigs))
	for _, pbSig := range pbSigs {
		s, err := NewECDSASigFromPB(pbSig)
		if err != nil {
			return nil, err
		}
		sigs = append(sigs, *s)
	}

	return sigs, nil
}

// NewECDSASigFromPB converts a pb EcdsaSignature to a RMNECDSASignature
func NewECDSASigFromPB(sig *rmnpb.EcdsaSignature) (*cciptypes.RMNECDSASignature, error) {
	if len(sig.R) != 32 || len(sig.S) != 32 {
		return nil, fmt.Errorf("invalid signature, R and S must be 32 bytes: %v", sig)
	}

	r := cciptypes.Bytes32{}
	s := cciptypes.Bytes32{}
	copy(r[:], sig.R[:32])
	copy(s[:], sig.S[:32])

	return &cciptypes.RMNECDSASignature{R: r, S: s}, nil
}
