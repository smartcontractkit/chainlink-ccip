package changesets

import (
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"

	v1_6_1_adapters "github.com/smartcontractkit/chainlink-ccip/deployment/v1_6_1/adapters"
	generic_changesets "github.com/smartcontractkit/chainlink-ccip/deployment/v1_6_1/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
)

// GlamsterdamGasUpdateV16Cfg is the EVM-specific config for the v1.6.1 Glamsterdam gas update changeset.
type GlamsterdamGasUpdateV16Cfg = generic_changesets.GlamsterdamGasUpdateV16Cfg

// UpdateGasConfigForGlamsterdamV16 returns the v1.6.1 Glamsterdam gas update changeset (EVM adapter).
// This adapts the generic changeset to use EVM-specific naming. adapterRegistry is forwarded
// verbatim to the generic changeset (injected rather than looked up from a global so the
// changeset remains testable with mocks); production callers pass
// v1_6_1_adapters.GetGasUpdateAdapterRegistry().
func UpdateGasConfigForGlamsterdamV16(registry *changesets.MCMSReaderRegistry, adapterRegistry *v1_6_1_adapters.GasUpdateAdapterRegistry) deployment.ChangeSetV2[changesets.WithMCMS[GlamsterdamGasUpdateV16Cfg]] {
	return generic_changesets.UpdateGasConfigForGlamsterdamV16(registry, adapterRegistry)
}
