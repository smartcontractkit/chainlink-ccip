package adapters

import (
	"github.com/Masterminds/semver/v3"

	chain_selectors "github.com/smartcontractkit/chain-selectors"

	tokensapi "github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	v1_6_1_adapters "github.com/smartcontractkit/chainlink-ccip/deployment/v1_6_1/adapters"
)

func init() {
	v1_6_1_adapters.GetGasUpdateAdapterRegistry().RegisterGasUpdateAdapter(
		chain_selectors.FamilyEVM,
		&GlamsterdamGasAdapter{},
	)

	adapter := NewTokenAdapter()
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
