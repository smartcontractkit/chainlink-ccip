package scripts

import (
	"testing"

	"github.com/smartcontractkit/crib/dependencies/ccip-v2-scripts/config"
)

func TestDeployContracts(t *testing.T) {
	t.Skip()
	t.Parallel()
	devspaceEnv := config.DevspaceEnv{
		Namespace:         "crib-local",
		Provider:          "kind",
		DonBootNodeCount:  1,
		DonNodeCount:      4,
		IngressBaseDomain: "main.stage.cldev.sh",
		TmpDir:            "/Users/radek/code/crib/deployments/ccip-v2/.tmp",
		GethChainsCount:   2,
	}

	configurer := NewRMNConfigurer(devspaceEnv, 3)
	configurer.SetupRMNOnChain()
}
