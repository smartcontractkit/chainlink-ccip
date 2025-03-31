package scripts

import (
	"testing"
)

func TestDeployContracts(t *testing.T) {
	t.Skip()
	t.Parallel()
	devspaceEnv := TestEnvKindLocal()

	configurer := NewRMNConfigurer(devspaceEnv, 3)
	configurer.SetupRMNOnChain()
}
