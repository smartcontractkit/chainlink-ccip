package example_ccip_sender // revive:disable-line:var-naming

import (
	feequoter "github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/fee_quoter"
)

// DO NOT DELETE - imported custom types are are not-automatically resolved by `anchor-go` but are declared in the anchor idl
// this files aliases types from other modules to ensure the go modules compile

type SVMTokenAmount = feequoter.SVMTokenAmount
