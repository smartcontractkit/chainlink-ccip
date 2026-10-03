package changesets_test

import (
	"testing"

	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/stretchr/testify/require"

	cs_core "github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v2_0_0/changesets"
)

// TestUpdateGasConfigForGlamsterdamV200_NoAdapterRegistered is a regression test: this test binary
// never imports an EVM (or any other chain family) adapter package, so the GasUpdateAdapterRegistry
// is empty. A missing adapter must fail the changeset explicitly rather than silently returning a
// "successful" empty output — automation must not be able to mistake an unexecuted migration for
// a completed one.
func TestUpdateGasConfigForGlamsterdamV200_NoAdapterRegistered(t *testing.T) {
	registry := cs_core.GetRegistry()
	_, err := changesets.UpdateGasConfigForGlamsterdamV200(registry).Apply(deployment.Environment{}, cs_core.WithMCMS[changesets.GlamsterdamGasUpdateV200Cfg]{
		MCMS: mcms.Input{},
		Cfg: changesets.GlamsterdamGasUpdateV200Cfg{
			TargetChainSelector: 123,
		},
	})
	require.Error(t, err)
	require.Contains(t, err.Error(), "no gas update adapter registered for EVM family")
}

func TestUpdateGasConfigForGlamsterdamV200_ValidateRequiresTarget(t *testing.T) {
	registry := cs_core.GetRegistry()
	_, err := changesets.UpdateGasConfigForGlamsterdamV200(registry).Apply(deployment.Environment{}, cs_core.WithMCMS[changesets.GlamsterdamGasUpdateV200Cfg]{
		MCMS: mcms.Input{},
		Cfg:  changesets.GlamsterdamGasUpdateV200Cfg{},
	})
	require.Error(t, err)
	require.Contains(t, err.Error(), "TargetChainSelector must be set")
}
