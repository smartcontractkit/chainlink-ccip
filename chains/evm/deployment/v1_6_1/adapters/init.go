package adapters

import (
	chain_selectors "github.com/smartcontractkit/chain-selectors"
	v1_6_1_adapters "github.com/smartcontractkit/chainlink-ccip/deployment/v1_6_1/adapters"
)

func init() {
	v1_6_1_adapters.GetGasUpdateAdapterRegistry().RegisterGasUpdateAdapter(
		chain_selectors.FamilyEVM,
		&GlamsterdamGasAdapter{},
	)
}
