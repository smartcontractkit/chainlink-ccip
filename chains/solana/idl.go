package solana

import _ "embed"

//go:embed contracts/target/idl/ccip_router.json
var ccipRouterIdl string

// FetchCCIPRouterIDL returns
func FetchCCIPRouterIDL() string {
	return ccipRouterIdl
}

//go:embed contracts/target/idl/ccip_offramp.json
var ccipOfframpIdl string

// FetchCCIPOfframpIDL returns
func FetchCCIPOfframpIDL() string {
	return ccipOfframpIdl
}

//go:embed contracts/target/idl/fee_quoter.json
var feeQuoterIdl string

// FetchFeeQuoterIDL returns
func FetchFeeQuoterIDL() string {
	return feeQuoterIdl
}

//go:embed contracts/target/idl/rmn_remote.json
var rmnRemoteIdl string

// FetchRMNRemoteIDL returns
func FetchRMNRemoteIDL() string {
	return rmnRemoteIdl
}

//go:embed contracts/target/idl/ccip_common.json
var ccipCommonIdl string

// FetchCommonIDL returns
func FetchCommonIDL() string {
	return ccipCommonIdl
}

//go:embed contracts/target/idl/cctp_token_pool.json
var cctpTokenPoolIdl string

// FetchCommonIDL returns
func FetchCctpTokenPoolIDL() string {
	return cctpTokenPoolIdl
}
