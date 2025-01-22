package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type TestData struct {
	DevspaceEnv                     DevspaceEnv
	DevspaceEnvWithAdditionalChains DevspaceEnv
}

func testData() TestData {
	return TestData{
		DevspaceEnv: DevspaceEnv{
			Namespace:         "crib-test",
			Provider:          "aws",
			DonBootNodeCount:  1,
			DonNodeCount:      4,
			IngressBaseDomain: "test.cl",
			TmpDir:            "./tmp",
		},
		DevspaceEnvWithAdditionalChains: DevspaceEnv{
			Namespace:             "crib-test",
			Provider:              "aws",
			DonBootNodeCount:      1,
			DonNodeCount:          4,
			IngressBaseDomain:     "test.cl",
			TmpDir:                "./tmp",
			AdditionalChainsCount: 2,
		},
	}
}

func TestGetEnvConfig(t *testing.T) {
	t.Parallel()
	env := testData().DevspaceEnv
	config, err := GetEnvConfig(env)
	require.NoError(t, err)
	assert.NotNil(t, config)
	assert.NotEmpty(t, config.Chains)
}

func TestGetEnvConfig_AdditionalChains(t *testing.T) {
	t.Parallel()
	env := DevspaceEnv{
		Namespace:             "crib-test",
		Provider:              "aws",
		DonBootNodeCount:      1,
		DonNodeCount:          4,
		IngressBaseDomain:     "test.cl",
		TmpDir:                "./tmp",
		AdditionalChainsCount: 2,
	}

	config, err := GetEnvConfig(env)
	require.NoError(t, err)
	assert.NotNil(t, config)
	assert.Len(t, config.Chains, 4)
}

func TestGetTransmittedChainConfigs(t *testing.T) {
	t.Parallel()
	env := testData().DevspaceEnv
	configs := GetTransmittedChainConfigs(env)
	assert.NotNil(t, configs)
	assert.Len(t, configs, 2)
}

func TestGetTransmittedChainConfigs_AdditionalChains(t *testing.T) {
	t.Parallel()
	env := testData().DevspaceEnvWithAdditionalChains
	configs := GetTransmittedChainConfigs(env)
	assert.NotNil(t, configs)
	assert.Len(t, configs, 4)
}
