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

// FetchRMNRemoteIDL returne
func FetchRMNRemoteIDL() string {
	return rmnRemoteIdl
}

//go:embed contracts/target/idl/ccip_common.json
var ccipCommonIdl string

// FetchCommonIDL returns
func FetchCommonIDL() string {
	return ccipCommonIdl
}
