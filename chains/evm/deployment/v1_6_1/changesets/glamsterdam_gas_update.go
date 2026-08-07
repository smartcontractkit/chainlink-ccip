package changesets

import (
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"

	generic_changesets "github.com/smartcontractkit/chainlink-ccip/deployment/v1_6_1/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
)

// GlamsterdamGasUpdateV16Cfg is the EVM-specific config for the v1.6.1 Glamsterdam gas update changeset.
type GlamsterdamGasUpdateV16Cfg = generic_changesets.GlamsterdamGasUpdateV16Cfg

// UpdateGasConfigForGlamsterdamV16 returns the v1.6.1 Glamsterdam gas update changeset (EVM adapter).
// This adapts the generic changeset to use EVM-specific naming.
func UpdateGasConfigForGlamsterdamV16(registry *changesets.MCMSReaderRegistry) deployment.ChangeSetV2[changesets.WithMCMS[GlamsterdamGasUpdateV16Cfg]] {
	return generic_changesets.UpdateGasConfigForGlamsterdamV16(registry)
}
