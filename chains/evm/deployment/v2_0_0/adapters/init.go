package adapters

import (
	chain_selectors "github.com/smartcontractkit/chain-selectors"
	v2_0_0_adapters "github.com/smartcontractkit/chainlink-ccip/deployment/v2_0_0/adapters"
)

func init() {
	v2_0_0_adapters.GetGasUpdateAdapterRegistry().RegisterGasUpdateAdapter(
		chain_selectors.FamilyEVM,
		&GlamsterdamGasAdapter{},
	)
}
