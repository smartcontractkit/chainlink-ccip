package rmn

import "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

type ReportSignatures struct {
	// ReportSignatures are the ECDSA signatures for the lane updates for each node.
	// NOTE: A signature[i] corresponds to the whole LaneUpdates slice and NOT LaneUpdates[i].
	Signatures  []ECDSASignature
	LaneUpdates []FixedDestLaneUpdate
}

type ECDSASignature struct {
	R []byte
	S []byte
}

type FixedDestLaneUpdate struct {
	SourceChain SourceChainInfo
	SeqNumRange ccipocr3.SeqNumRange
	MerkleRoot  []byte
}

type FixedDestLaneUpdateRequest struct {
	SourceChainInfo SourceChainInfo
	SeqNumRange     ccipocr3.SeqNumRange
}

type SourceChainInfo struct {
	Chain         ccipocr3.ChainSelector
	OnRampAddress []byte
}

type DestChainInfo struct {
	Chain          ccipocr3.ChainSelector
	OffRampAddress []byte
}
