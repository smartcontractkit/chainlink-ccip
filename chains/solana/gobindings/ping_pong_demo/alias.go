package ping_pong_demo // revive:disable-line:var-naming

import (
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/example_ccip_receiver"
)

// DO NOT DELETE - imported custom types are are not-automatically resolved by `anchor-go` but are declared in the anchor idl
// this files aliases types from other modules to ensure the go modules compile

type Any2SVMMessage = example_ccip_receiver.Any2SVMMessage
