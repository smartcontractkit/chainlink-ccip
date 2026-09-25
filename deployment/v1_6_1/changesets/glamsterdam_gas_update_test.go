package changesets_test

import (
	"testing"

	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/stretchr/testify/require"

	cs_core "github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_6_1/changesets"
)

// TestUpdateGasConfigForGlamsterdamV16_NoAdapterRegistered is a regression test: this test binary
// never imports an EVM (or any other chain family) adapter package, so the GasUpdateAdapterRegistry
// is empty. A missing adapter must fail the changeset explicitly rather than silently returning a
// "successful" empty output — automation must not be able to mistake an unexecuted migration for
// a completed one.
func TestUpdateGasConfigForGlamsterdamV16_NoAdapterRegistered(t *testing.T) {
	registry := cs_core.GetRegistry()
	_, err := changesets.UpdateGasConfigForGlamsterdamV16(registry).Apply(deployment.Environment{}, cs_core.WithMCMS[changesets.GlamsterdamGasUpdateV16Cfg]{
		MCMS: mcms.Input{},
		Cfg: changesets.GlamsterdamGasUpdateV16Cfg{
			TargetChainSelector: 123,
		},
	})
	require.Error(t, err)
	require.Contains(t, err.Error(), "no gas update adapter registered for EVM family")
}

func TestUpdateGasConfigForGlamsterdamV16_ValidateRequiresTarget(t *testing.T) {
	registry := cs_core.GetRegistry()
	_, err := changesets.UpdateGasConfigForGlamsterdamV16(registry).Apply(deployment.Environment{}, cs_core.WithMCMS[changesets.GlamsterdamGasUpdateV16Cfg]{
		MCMS: mcms.Input{},
		Cfg:  changesets.GlamsterdamGasUpdateV16Cfg{},
	})
	require.Error(t, err)
	require.Contains(t, err.Error(), "TargetChainSelector must be set")
}
