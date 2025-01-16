package solana

import _ "embed"

//go:embed contracts/target/idl/ccip_router.json
var ccipRouterIdl string

// FetchCCIPRouterIDL returns
func FetchCCIPRouterIDL() string {
	return ccipRouterIdl
}
