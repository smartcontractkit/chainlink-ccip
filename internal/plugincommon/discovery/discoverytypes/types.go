package discoverytypes

import "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

// Outcome isn't needed for this processor.
type Outcome struct {
	// TODO: some sort of request flag to avoid including this every time.
	// Request bool
}

// Query isn't needed for this processor.
type Query []byte

// Observation of contract addresses.
type Observation struct {
	FChain           map[ccipocr3.ChainSelector]int `json:"fChain"`
	OnRamp           map[ccipocr3.ChainSelector][]byte
	DestNonceManager []byte
	RMNRemote        []byte

	// TODO: some sort of request flag to avoid including this every time.
	// Request bool
}
