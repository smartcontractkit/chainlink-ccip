package config

import (
	"testing"

	"github.com/smartcontractkit/chainlink/deployment/environment/crib"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type TestData struct {
	DevspaceEnv                        DevspaceEnv
	DevspaceEnvWithAdditionalChains    DevspaceEnv
	DevspaceEnvWithGethAndSolanaChains DevspaceEnv
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
		Namespace:         "crib-test",
		Provider:          "aws",
		DonBootNodeCount:  1,
		DonNodeCount:      4,
		IngressBaseDomain: "test.cl",
		TmpDir:            "./tmp",
		GethChainsCount:   4,
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

func strPtr(s string) *string {
	return &s
}

func TestGetTransmittedChainConfigs_GethAndSolanaChain(t *testing.T) {
	t.Parallel()
	env := testData().DevspaceEnvWithGethAndSolanaChains
	configs := GetTransmittedChainConfigs(env)

	expected := []crib.ChainConfig{
		{
			ChainID:   1337,
			ChainType: "EVM",
			ChainName: "evm-simulated-1337",
			WSRPCs: []crib.RPC{
				{
					External: strPtr("wss://crib-test-geth-1337-ws.main.stage.cldev.sh"),
					Internal: strPtr("ws://geth-1337-ws:8546"),
				},
			},
			HTTPRPCs: []crib.RPC{
				{
					External: strPtr("https://crib-test-geth-1337-http.main.stage.cldev.sh:443"),
					Internal: strPtr("http://geth-1337:8544"),
				},
			},
		},
		{
			ChainID:   2337,
			ChainType: "EVM",
			ChainName: "evm-simulated-2337",
			WSRPCs: []crib.RPC{
				{
					External: strPtr("wss://crib-test-geth-2337-ws.main.stage.cldev.sh"),
					Internal: strPtr("ws://geth-2337-ws:8546"),
				},
			},
			HTTPRPCs: []crib.RPC{
				{
					External: strPtr("https://crib-test-geth-2337-http.main.stage.cldev.sh:443"),
					Internal: strPtr("http://geth-2337:8544"),
				},
			},
		},
		{
			ChainID:   1001,
			ChainType: "SOLANA",
			ChainName: "solana-simulated-1001",
			WSRPCs: []crib.RPC{
				{
					External: strPtr("wss://crib-test-solana-1001-ws.main.stage.cldev.sh"),
					Internal: strPtr("ws://solana-1001:8545"),
				},
			},
			HTTPRPCs: []crib.RPC{
				{
					External: strPtr("https://crib-test-solana-1001-http.main.stage.cldev.sh:443"),
					Internal: strPtr("http://solana-1001:8544"),
				},
			},
		},
	}

	assert.NotNil(t, configs)
	assert.Equal(t, expected, configs)
}
