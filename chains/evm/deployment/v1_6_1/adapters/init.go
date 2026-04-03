package adapters

import (
	"github.com/Masterminds/semver/v3"

	chain_selectors "github.com/smartcontractkit/chain-selectors"

	tokensapi "github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
)

func init() {
	adapter := &TokenAdapter{}
	tokensapi.GetTokenAdapterRegistry().RegisterTokenAdapter(
		chain_selectors.FamilyEVM,
		semver.MustParse("1.6.0"),
		adapter,
	)
	tokensapi.GetTokenAdapterRegistry().RegisterTokenAdapter(
		chain_selectors.FamilyEVM,
		semver.MustParse("1.6.1"),
		adapter,
	)
}
