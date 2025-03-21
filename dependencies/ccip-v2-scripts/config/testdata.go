package config

type DevspaceEnvTestData struct {
	DevspaceEnv                        DevspaceEnv
	DevspaceEnvWithAdditionalChains    DevspaceEnv
	DevspaceEnvWithGethAndSolanaChains DevspaceEnv
	DevspaceEnvWithBesuAndSolanaChains DevspaceEnv
}

func GetDevspaceEnvTestData() DevspaceEnvTestData {
	return DevspaceEnvTestData{
		DevspaceEnv: DevspaceEnv{
			Namespace:         "crib-test",
			Provider:          "aws",
			DonBootNodeCount:  1,
			DonNodeCount:      4,
			IngressBaseDomain: "test.cl",
			TmpDir:            "./tmp",
			GethChainsCount:   2,
		},
		DevspaceEnvWithAdditionalChains: DevspaceEnv{
			Namespace:         "crib-test",
			Provider:          "aws",
			DonBootNodeCount:  1,
			DonNodeCount:      4,
			IngressBaseDomain: "test.cl",
			TmpDir:            "./tmp",
			GethChainsCount:   4,
		},
		DevspaceEnvWithGethAndSolanaChains: DevspaceEnv{
			Namespace:         "crib-test",
			Provider:          "aws",
			DonBootNodeCount:  1,
			DonNodeCount:      4,
			IngressBaseDomain: "main.stage.cldev.sh",
			TmpDir:            "./tmp",
			GethChainsCount:   2,
			SolanaChainsCount: 1,
		},
		DevspaceEnvWithBesuAndSolanaChains: DevspaceEnv{
			Namespace:         "crib-test",
			Provider:          "aws",
			DonBootNodeCount:  1,
			DonNodeCount:      4,
			IngressBaseDomain: "main.stage.cldev.sh",
			TmpDir:            "./tmp",
			BesuChainsCount:   2,
			GethChainsCount:   0,
			SolanaChainsCount: 1,
		},
	}
}
