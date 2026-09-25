package changesets

import (
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"

	generic_changesets "github.com/smartcontractkit/chainlink-ccip/deployment/v2_0_0/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
)

// GlamsterdamGasUpdateCfg is the EVM-specific config for the Glamsterdam gas update changeset.
type GlamsterdamGasUpdateCfg = generic_changesets.GlamsterdamGasUpdateV200Cfg

// UpdateGasConfigForGlamsterdamV2 returns the v2.0.0 Glamsterdam gas update changeset (EVM adapter).
// This adapts the generic changeset to use EVM-specific type names.
func UpdateGasConfigForGlamsterdamV2(registry *changesets.MCMSReaderRegistry) deployment.ChangeSetV2[changesets.WithMCMS[GlamsterdamGasUpdateCfg]] {
	return generic_changesets.UpdateGasConfigForGlamsterdamV200(registry)
}
