package test_token_pool // revive:disable-line:var-naming

import "github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/base_token_pool"

// DO NOT DELETE - imported custom types are are not-automatically resolved by `anchor-go` but are declared in the anchor idl
// this files aliases types from other modules to ensure the go modules compile

type BaseChain = base_token_pool.BaseChain
type BaseConfig = base_token_pool.BaseConfig
type LockOrBurnInV1 = base_token_pool.LockOrBurnInV1
type RateLimitConfig = base_token_pool.RateLimitConfig
type ReleaseOrMintInV1 = base_token_pool.ReleaseOrMintInV1
type RemoteAddress = base_token_pool.RemoteAddress
type RemoteConfig = base_token_pool.RemoteConfig
