package burnmint_token_pool // revive:disable-line:var-naming

import (
	base "github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/base_token_pool"
)

// DO NOT DELETE - imported custom types are are not-automatically resolved by `anchor-go` but are declared in the anchor idl
// this files aliases types from other modules to ensure the go modules compile

type RemoteAddress = base.RemoteAddress
type RemoteConfig = base.RemoteConfig
type LockOrBurnInV1 = base.LockOrBurnInV1
type ReleaseOrMintInV1 = base.ReleaseOrMintInV1
type RateLimitConfig = base.RateLimitConfig
type BaseConfig = base.BaseConfig
type BaseChain = base.BaseChain
