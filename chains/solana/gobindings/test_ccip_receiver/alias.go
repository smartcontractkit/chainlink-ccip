package test_ccip_receiver // revive:disable-line:var-naming

import (
	ccipReceiver "github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/example_ccip_receiver"
)

// DO NOT DELETE - imported custom types are are not-automatically resolved by `anchor-go` but are declared in the anchor idl
// this files aliases types from other modules to ensure the go modules compile

type BaseState = ccipReceiver.BaseState
type Any2SVMMessage = ccipReceiver.Any2SVMMessage
