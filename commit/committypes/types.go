package committypes

import (
	"github.com/smartcontractkit/chainlink-ccip/commit/chainfee"
	"github.com/smartcontractkit/chainlink-ccip/commit/merkleroot"
	"github.com/smartcontractkit/chainlink-ccip/commit/tokenprice"
	dt "github.com/smartcontractkit/chainlink-ccip/internal/plugincommon/discovery/discoverytypes"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

type Query struct {
	MerkleRootQuery merkleroot.Query `json:"merkleRootQuery"`
	TokenPriceQuery tokenprice.Query `json:"tokenPriceQuery"`
	ChainFeeQuery   chainfee.Query   `json:"chainFeeQuery"`
}

type Observation struct {
	MerkleRootObs merkleroot.Observation          `json:"merkleObs"`
	TokenPriceObs tokenprice.Observation          `json:"tokenObs"`
	ChainFeeObs   chainfee.Observation            `json:"chainFeeObs"`
	DiscoveryObs  dt.Observation                  `json:"discoveryObs"`
	FChain        map[cciptypes.ChainSelector]int `json:"fChain"`
}

type Outcome struct {
	MerkleRootOutcome merkleroot.Outcome `json:"merkleRootOutcome"`
	TokenPriceOutcome tokenprice.Outcome `json:"tokenPriceOutcome"`
	ChainFeeOutcome   chainfee.Outcome   `json:"chainFeeOutcome"`
	MainOutcome       MainOutcome        `json:"mainOutcome"`
}

// MainOutcome contains fields produced by the main commit plugin outcome (not of some sub-processor).
type MainOutcome struct {
	// InflightPriceOcrSequenceNumber is the OCR sequence number of the latest price-related outcome
	// that hasn't been confirmed on the blockchain yet. If it is set to 0, you can assume we don't have to wait.
	InflightPriceOcrSequenceNumber cciptypes.SeqNum `json:"inflightPriceOcrSequenceNumber"`

	// RemainingPriceChecks is how many more times we will check
	// if a previous price report has been recorded on the blockchain.
	// If it is set to 0, you can assume we don't have to wait.
	RemainingPriceChecks int `json:"remainingPriceChecks"`
}
