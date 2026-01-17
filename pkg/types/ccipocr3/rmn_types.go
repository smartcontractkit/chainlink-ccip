package ccipocr3

import ccipocr3common "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

// Deprecated: Use ccipocr3common.RMNReport instead.
type RMNReport = ccipocr3common.RMNReport

// Deprecated: Use ccipocr3common.NewRMNReport instead.
func NewRMNReport(
	reportVersionDigest Bytes32,
	destChainID BigInt,
	destChainSelector ChainSelector,
	rmnRemoteContractAddress UnknownAddress,
	offRampAddress UnknownAddress,
	rmnHomeContractConfigDigest Bytes32,
	laneUpdates []RMNLaneUpdate,
) RMNReport {
	commonRMNLaneUpdates := make([]ccipocr3common.RMNLaneUpdate, len(laneUpdates))
	copy(commonRMNLaneUpdates, laneUpdates)
	return ccipocr3common.NewRMNReport(
		reportVersionDigest,
		destChainID,
		destChainSelector,
		rmnRemoteContractAddress,
		offRampAddress,
		rmnHomeContractConfigDigest,
		commonRMNLaneUpdates,
	)

}

// Deprecated: Use ccipocr3common.RMNLaneUpdate instead.
type RMNLaneUpdate = ccipocr3common.RMNLaneUpdate

// Deprecated: Use ccipocr3common.RMNECDSASignature instead.
type RMNECDSASignature = ccipocr3common.RMNECDSASignature
