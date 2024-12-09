package ccipocr3

// RMNReport is the payload that is signed by the RMN nodes, transmitted and verified onchain.
type RMNReport struct {
	ReportVersionDigest         Bytes32 // e.g. keccak256("RMN_V1_6_ANY2EVM_REPORT")
	DestChainID                 BigInt  // If applies, a chain specific id, e.g. evm chain id otherwise empty.
	DestChainSelector           ChainSelector
	RmnRemoteContractAddress    UnknownAddress
	OfframpAddress              UnknownAddress
	RmnHomeContractConfigDigest Bytes32
	LaneUpdates                 []RMNLaneUpdate
}

func NewRMNReport(
	reportVersionDigest Bytes32,
	destChainID BigInt,
	destChainSelector ChainSelector,
	rmnRemoteContractAddress UnknownAddress,
	offRampAddress UnknownAddress,
	rmnHomeContractConfigDigest Bytes32,
	laneUpdates []RMNLaneUpdate,
) RMNReport {
	return RMNReport{
		ReportVersionDigest:         reportVersionDigest,
		DestChainID:                 destChainID,
		DestChainSelector:           destChainSelector,
		RmnRemoteContractAddress:    rmnRemoteContractAddress,
		OfframpAddress:              offRampAddress,
		RmnHomeContractConfigDigest: rmnHomeContractConfigDigest,
		LaneUpdates:                 laneUpdates,
	}
}

// RMNLaneUpdate represents an interval that has been observed by an RMN node.
// It is part of the payload that is signed and transmitted onchain.
type RMNLaneUpdate struct {
	SourceChainSelector ChainSelector
	OnRampAddress       UnknownAddress // (for EVM should be abi-encoded)
	MinSeqNr            SeqNum
	MaxSeqNr            SeqNum
	MerkleRoot          Bytes32
}

// RMNECDSASignature represents the signature provided by RMN on the RMNReport structure.
// The V value of the signature is included in the top-level commit report as part of a
// bitmap.
type RMNECDSASignature struct {
	R Bytes32 `json:"r"`
	S Bytes32 `json:"s"`
}
