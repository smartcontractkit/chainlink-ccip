package scripts

import "github.com/smartcontractkit/crib/dependencies/ccip-v2-scripts/config"

func TestEnvKindLocal() config.DevspaceEnv {
	env := config.DevspaceEnv{
		Namespace:         "crib-local",
		Provider:          "kind",
		DonBootNodeCount:  1,
		DonNodeCount:      4,
		GethChainsCount:   2,
		IngressBaseDomain: "main.stage.cldev.sh",
		// TmpDir:            "/Users/radek/code/crib/deployments/ccip-v2/.tmp",
	}
	return env
}

func TestEnvFWOG() config.DevspaceEnv {
	env := config.DevspaceEnv{
		Namespace:         "ccip",
		Provider:          "aws",
		DonBootNodeCount:  1,
		DonNodeCount:      4,
		BesuChainsCount:   2,
		IngressBaseDomain: "fwog.sandbox.enterprise.chain.link",
		// TmpDir:            "/Users/radek/code/crib/deployments/cre-enterprise-sandbox/.tmp",
	}
	return env
}
