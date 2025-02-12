package discoverytypes

import (
	"github.com/smartcontractkit/chainlink-ccip/pkg/reader"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

// Outcome isn't needed for this processor.
type Outcome struct {
	// TODO: some sort of request flag to avoid including this every time.
	// Request bool
}

func (o Outcome) Stats() map[string]int {
	return map[string]int{}
}

// Query isn't needed for this processor.
type Query []byte

// Observation of contract addresses.
type Observation struct {
	//SourceChains     []ccipocr3.ChainSelector
	FChain map[cciptypes.ChainSelector]int
	// See reader.ContractAddresses for more info on this data structure.
	Addresses reader.ContractAddresses

	// TODO: some sort of request flag to avoid including this every time.
	// Request bool
}

func (o Observation) Stats() map[string]int {
	return map[string]int{}
}

func (o Observation) IsEmpty() bool {
	return len(o.Addresses) == 0
}
