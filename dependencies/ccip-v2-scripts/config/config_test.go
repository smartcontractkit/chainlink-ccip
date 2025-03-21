package config

import (
	"fmt"
	"testing"

	"github.com/smartcontractkit/chainlink/deployment/environment/crib"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetEnvConfig(t *testing.T) {
	t.Parallel()
	env := GetDevspaceEnvTestData().DevspaceEnv
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
	fmt.Printf("%+v\n", err)
	require.NoError(t, err)
	assert.NotNil(t, config)
	assert.Len(t, config.Chains, 4)
}

func TestGetTransmittedChainConfigs(t *testing.T) {
	t.Parallel()
	env := GetDevspaceEnvTestData().DevspaceEnv
	configs := GetTransmittedChainConfigs(env)
	assert.NotNil(t, configs)
	assert.Len(t, configs, 2)
}

func TestGetTransmittedChainConfigs_AdditionalChains(t *testing.T) {
	t.Parallel()
	env := GetDevspaceEnvTestData().DevspaceEnvWithAdditionalChains
	configs := GetTransmittedChainConfigs(env)
	assert.NotNil(t, configs)
	assert.Len(t, configs, 4)
}

func TestGetTransmittedChainConfigs_MixedChains(t *testing.T) {
	t.Parallel()
	env := DevspaceEnv{
		Namespace:         "crib-test",
		Provider:          "aws",
		DonBootNodeCount:  1,
		DonNodeCount:      4,
		IngressBaseDomain: "test.cl",
		TmpDir:            "./tmp",
		BesuChainsCount:   1,
		GethChainsCount:   1,
		SolanaChainsCount: 1,
	}

	configs := GetTransmittedChainConfigs(env)
	assert.NotNil(t, configs)
	assert.Len(t, configs, 3)
}

func strPtr(s string) *string {
	return &s
}

func TestGetTransmittedChainConfigs_GethAndSolanaChain(t *testing.T) {
	t.Parallel()
	env := GetDevspaceEnvTestData().DevspaceEnvWithGethAndSolanaChains
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

func TestBuildEVMNetworkConfigs(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name        string
		chainsCount int
		expected    []EVMChain
	}{
		{
			name:        "Zero chains",
			chainsCount: 0,
			expected:    []EVMChain{},
		},
		{
			name:        "One chain",
			chainsCount: 1,
			expected: []EVMChain{
				{NetworkId: 1337, FinalityDepth: defaultFinalityDepth},
			},
		},
		{
			name:        "Two chains",
			chainsCount: 2,
			expected: []EVMChain{
				{NetworkId: 1337, FinalityDepth: defaultFinalityDepth},
				{NetworkId: 2337, FinalityDepth: defaultFinalityDepth},
			},
		},
		{
			name:        "Three chains",
			chainsCount: 3,
			expected: []EVMChain{
				{NetworkId: 1337, FinalityDepth: defaultFinalityDepth},
				{NetworkId: 2337, FinalityDepth: defaultFinalityDepth},
				{NetworkId: 90000001, FinalityDepth: defaultFinalityDepth},
			},
		},
		{
			name:        "Four chains",
			chainsCount: 4,
			expected: []EVMChain{
				{NetworkId: 1337, FinalityDepth: defaultFinalityDepth},
				{NetworkId: 2337, FinalityDepth: defaultFinalityDepth},
				{NetworkId: 90000001, FinalityDepth: defaultFinalityDepth},
				{NetworkId: 90000002, FinalityDepth: defaultFinalityDepth},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := BuildEVMNetworkConfigs(tt.chainsCount)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestGetChainName(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name            string
		chainConfigurer ChainConfigurer
		expected        string
	}{
		{
			name: "EVM chain",
			chainConfigurer: ChainConfigurer{
				chainID:   1337,
				chainType: EVM,
			},
			expected: "evm-simulated-1337",
		},
		{
			name: "Solana chain",
			chainConfigurer: ChainConfigurer{
				chainID:   1001,
				chainType: SOLANA,
			},
			expected: "solana-simulated-1001",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := tt.chainConfigurer.GetChainName()
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestChainConfigurer_internalHTTPRPC(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name            string
		chainConfigurer ChainConfigurer
		expected        string
	}{
		{
			name: "EVM Besu chain",
			chainConfigurer: ChainConfigurer{
				chainID:      1337,
				chainType:    EVM,
				chainVariant: "besu",
				chainName:    "alpha",
			},
			expected: "http://besu-node-rpc-1.chain-alpha.svc.cluster.local:8545",
		},
		{
			name: "EVM Geth chain",
			chainConfigurer: ChainConfigurer{
				chainID:      2337,
				chainType:    EVM,
				chainVariant: "geth",
				chainName:    "beta",
			},
			expected: "http://geth-2337:8544",
		},
		{
			name: "Solana chain",
			chainConfigurer: ChainConfigurer{
				chainID:   1001,
				chainType: SOLANA,
				chainName: "solana",
			},
			expected: "http://solana-1001:8544",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := tt.chainConfigurer.InternalHTTPRPC()
			assert.NotNil(t, result)
			assert.Equal(t, tt.expected, *result)
		})
	}
}

func TestChainConfigurer_internalWSRPC(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name            string
		chainConfigurer ChainConfigurer
		expected        string
	}{
		{
			name: "EVM Besu chain",
			chainConfigurer: ChainConfigurer{
				chainID:      1337,
				chainType:    EVM,
				chainVariant: "besu",
				chainName:    "alpha",
			},
			expected: "ws://besu-node-rpc-1.chain-alpha.svc.cluster.local:8546",
		},
		{
			name: "EVM Geth chain",
			chainConfigurer: ChainConfigurer{
				chainID:      2337,
				chainType:    EVM,
				chainVariant: "geth",
				chainName:    "beta",
			},
			expected: "ws://geth-2337-ws:8546",
		},
		{
			name: "Solana chain",
			chainConfigurer: ChainConfigurer{
				chainID:   1001,
				chainType: SOLANA,
				chainName: "solana",
			},
			expected: "ws://solana-1001:8545",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := tt.chainConfigurer.InternalWSRPC()
			assert.NotNil(t, result)
			assert.Equal(t, tt.expected, *result)
		})
	}
}

func TestChainTypeHostNamePart(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name            string
		chainConfigurer ChainConfigurer
		expected        string
	}{
		{
			name: "EVM Besu chain",
			chainConfigurer: ChainConfigurer{
				chainType:    EVM,
				chainVariant: "besu",
			},
			expected: "besu",
		},
		{
			name: "EVM Geth chain",
			chainConfigurer: ChainConfigurer{
				chainType:    EVM,
				chainVariant: "geth",
			},
			expected: "geth",
		},
		{
			name: "Solana chain",
			chainConfigurer: ChainConfigurer{
				chainType: SOLANA,
			},
			expected: "solana",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := tt.chainConfigurer.chainTypeHostNamePart()
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestChainConfigurer_externalHTTPRPC(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name            string
		chainConfigurer ChainConfigurer
		expected        string
	}{
		{
			name: "EVM Besu chain in CI environment",
			chainConfigurer: ChainConfigurer{
				chainID:      1337,
				chainType:    EVM,
				chainVariant: "besu",
				chainName:    "alpha",
				env: DevspaceEnv{
					CIEnv:             true,
					Namespace:         "crib-test",
					IngressBaseDomain: "main.stage.cldev.sh",
				},
			},
			expected: "https://gap-crib-test-besu-1337.public.main.stage.cldev.sh:443",
		},
		{
			name: "EVM Besu chain",
			chainConfigurer: ChainConfigurer{
				chainID:      1337,
				chainType:    EVM,
				chainVariant: "besu",
				chainName:    "alpha",
				env: DevspaceEnv{
					CIEnv:             false,
					Namespace:         "crib-test",
					IngressBaseDomain: "main.stage.cldev.sh",
				},
			},
			expected: "https://chain-alpha-rpc.main.stage.cldev.sh",
		},
		{
			name: "EVM Geth chain",
			chainConfigurer: ChainConfigurer{
				chainID:      2337,
				chainType:    EVM,
				chainVariant: "geth",
				chainName:    "beta",
				env: DevspaceEnv{
					CIEnv:             false,
					Namespace:         "crib-test",
					IngressBaseDomain: "main.stage.cldev.sh",
				},
			},
			expected: "https://crib-test-geth-2337-http.main.stage.cldev.sh:443",
		},
		{
			name: "Solana chain",
			chainConfigurer: ChainConfigurer{
				chainID:   1001,
				chainType: SOLANA,
				chainName: "solana",
				env: DevspaceEnv{
					CIEnv:             false,
					Namespace:         "crib-test",
					IngressBaseDomain: "main.stage.cldev.sh",
				},
			},
			expected: "https://crib-test-solana-1001-http.main.stage.cldev.sh:443",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := tt.chainConfigurer.ExternalHTTPRPC()
			assert.NotNil(t, result)
			assert.Equal(t, tt.expected, *result)
		})
	}
}

func TestChainConfigurer_externalWSRPC(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name            string
		chainConfigurer ChainConfigurer
		expected        string
	}{
		{
			name: "EVM Besu chain in CI environment",
			chainConfigurer: ChainConfigurer{
				chainID:      1337,
				chainType:    EVM,
				chainVariant: "besu",
				chainName:    "alpha",
				env: DevspaceEnv{
					CIEnv:             true,
					Namespace:         "crib-test",
					IngressBaseDomain: "main.stage.cldev.sh",
				},
			},
			expected: "wss://gap-crib-test-alpha-1337-ws.public.main.stage.cldev.sh:443",
		},
		{
			name: "EVM Besu chain",
			chainConfigurer: ChainConfigurer{
				chainID:      1337,
				chainType:    EVM,
				chainVariant: "besu",
				chainName:    "alpha",
				env: DevspaceEnv{
					CIEnv:             false,
					Namespace:         "crib-test",
					IngressBaseDomain: "main.stage.cldev.sh",
				},
			},
			expected: "wss://chain-alpha-rpc.main.stage.cldev.sh/ws/",
		},
		{
			name: "EVM Geth chain",
			chainConfigurer: ChainConfigurer{
				chainID:      2337,
				chainType:    EVM,
				chainVariant: "geth",
				chainName:    "beta",
				env: DevspaceEnv{
					CIEnv:             false,
					Namespace:         "crib-test",
					IngressBaseDomain: "main.stage.cldev.sh",
				},
			},
			expected: "wss://crib-test-geth-2337-ws.main.stage.cldev.sh",
		},
		{
			name: "Solana chain",
			chainConfigurer: ChainConfigurer{
				chainID:   1001,
				chainType: SOLANA,
				chainName: "solana",
				env: DevspaceEnv{
					CIEnv:             false,
					Namespace:         "crib-test",
					IngressBaseDomain: "main.stage.cldev.sh",
				},
			},
			expected: "wss://crib-test-solana-1001-ws.main.stage.cldev.sh",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := tt.chainConfigurer.ExternalWSRPC()
			assert.NotNil(t, result)
			assert.Equal(t, tt.expected, *result)
		})
	}
}
