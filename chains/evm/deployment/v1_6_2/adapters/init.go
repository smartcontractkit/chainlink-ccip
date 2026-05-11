package adapters

import (
	chain_selectors "github.com/smartcontractkit/chain-selectors"

	tokensapi "github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils"
)

// NOTE: v1.6.2 token pool contracts still use the base v1.6.1 token pool bindings

func init() {
	version := utils.Version_1_6_2
	adapter := NewTokenAdapter()
	tokensapi.GetTokenAdapterRegistry().RegisterTokenAdapter(
		chain_selectors.FamilyEVM,
		version,
		adapter,
	)
	tokensapi.GetTokenAdapterRegistry().RegisterTokenAdapter(
		chain_selectors.FamilyEVM,
		version,
		adapter,
	)
}
