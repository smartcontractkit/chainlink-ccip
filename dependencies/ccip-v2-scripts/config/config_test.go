package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func getTestDevspaceEnv() DevspaceEnv {
	env := DevspaceEnv{
		Namespace:         "crib-test",
		Provider:          "aws",
		DonBootNodeCount:  1,
		DonNodeCount:      4,
		IngressBaseDomain: "test.cl",
		TmpDir:            "./tmp",
	}
	return env
}

func TestGetEnvConfig(t *testing.T) {
	t.Parallel()
	env := getTestDevspaceEnv()
	config, err := GetEnvConfig(env, "")
	require.NoError(t, err)
	assert.NotNil(t, config)
	assert.NotEmpty(t, config.Chains)
}

func TestGetTransmittedChainConfigs(t *testing.T) {
	t.Parallel()
	env := getTestDevspaceEnv()
	configs := GetTransmittedChainConfigs(env)
	assert.NotNil(t, configs)
	assert.NotEmpty(t, configs)
}
