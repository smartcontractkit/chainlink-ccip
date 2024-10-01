package discoverytypes

import (
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
)

// Outcome isn't needed for this processor.
type Outcome struct {
	// TODO: some sort of request flag to avoid including this every time.
	// Request bool
}

// Query isn't needed for this processor.
type Query []byte

// Observation of contract addresses.
type Observation struct {
	//SourceChains     []ccipocr3.ChainSelector
	FChain map[cciptypes.ChainSelector]int
	// See reader.ContractAddresses for more info on this data structure.
	Addresses map[string]map[cciptypes.ChainSelector][]byte

	// TODO: fix circular dependency /w token reader
	//Addresses  reader.ContractAddresses

	// TODO: some sort of request flag to avoid including this every time.
	// Request bool
}
