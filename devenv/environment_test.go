package ccip

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNodeConfigOverrides(t *testing.T) {
	// Create a temporary TOML config file with node_config_overrides
	tmpDir := t.TempDir()
	configContent := `
cl_nodes_funding_eth = 50
cl_nodes_funding_link = 50

node_config_overrides = """
[[EVM]]
ChainID = '43113'
FinalityDepth = 5
LogPollInterval = '3s'
"""

[cldf]

[[blockchains]]
  container_name = "blockchain-src"
  chain_id = "1337"
  type = "anvil"

[[nodesets]]
  name = "don"
  nodes = 2
  override_mode = "each"

  [nodesets.db]
    image = "postgres:15.0"

  [[nodesets.node_specs]]
    [nodesets.node_specs.node]
      image = "public.ecr.aws/chainlink/chainlink:2.30.0"

  [[nodesets.node_specs]]
    [nodesets.node_specs.node]
      image = "public.ecr.aws/chainlink/chainlink:2.30.0"
      test_config_overrides = """
        [[EVM]]
        ChainID = '1337'
        FinalityDepth = 10
      """
`
	configPath := filepath.Join(tmpDir, "test-config.toml")
	err := os.WriteFile(configPath, []byte(configContent), 0644)
	require.NoError(t, err)

	// Change to tmpDir so Load can find the file
	originalDir, err := os.Getwd()
	require.NoError(t, err)
	defer os.Chdir(originalDir)
	err = os.Chdir(tmpDir)
	require.NoError(t, err)

	// Load the config
	cfg, err := Load[Cfg]([]string{"test-config.toml"})
	require.NoError(t, err)

	// Verify the NodeConfigOverrides field is loaded
	assert.Contains(t, cfg.NodeConfigOverrides, "ChainID = '43113'")
	assert.Contains(t, cfg.NodeConfigOverrides, "FinalityDepth = 5")
	assert.Contains(t, cfg.NodeConfigOverrides, "LogPollInterval = '3s'")

	// Verify per-node overrides are preserved
	require.Len(t, cfg.NodeSets, 1)
	require.Len(t, cfg.NodeSets[0].NodeSpecs, 2)

	// First node has no per-node overrides
	assert.Empty(t, cfg.NodeSets[0].NodeSpecs[0].Node.TestConfigOverrides)

	// Second node has per-node overrides
	assert.Contains(t, cfg.NodeSets[0].NodeSpecs[1].Node.TestConfigOverrides, "FinalityDepth = 10")
}

func TestNodeConfigMerging(t *testing.T) {
	// Test the config merging logic
	generatedConfig := `[Log]
Level = 'debug'

[[EVM]]
ChainID = '1337'
FinalityDepth = 1`

	topLevelOverrides := `[[EVM]]
ChainID = '43113'
FinalityDepth = 5`

	perNodeOverrides := `[[EVM]]
ChainID = '1337'
FinalityDepth = 10`

	// Simulate the merging logic from environment.go
	configParts := []string{generatedConfig}
	if topLevelOverrides != "" {
		configParts = append(configParts, topLevelOverrides)
	}
	if perNodeOverrides != "" {
		configParts = append(configParts, perNodeOverrides)
	}

	merged := strings.Join(configParts, "\n")

	// Verify all parts are in the merged config
	assert.Contains(t, merged, "[Log]")
	assert.Contains(t, merged, "Level = 'debug'")
	assert.Contains(t, merged, "ChainID = '1337'")
	assert.Contains(t, merged, "ChainID = '43113'")

	// Verify the order (later overrides should come after)
	generatedIdx := strings.Index(merged, "FinalityDepth = 1")
	topLevelIdx := strings.Index(merged, "FinalityDepth = 5")
	perNodeIdx := strings.Index(merged, "FinalityDepth = 10")

	assert.Less(t, generatedIdx, topLevelIdx, "Generated config should come before top-level overrides")
	assert.Less(t, topLevelIdx, perNodeIdx, "Top-level overrides should come before per-node overrides")
}

func TestLocalChainConfigNoOverrides(t *testing.T) {
	// Test that local chain config (env.toml) still works without any overrides
	tmpDir := t.TempDir()
	configContent := `
cl_nodes_funding_eth = 50
cl_nodes_funding_link = 50

[cldf]

[[blockchains]]
  container_name = "blockchain-src"
  chain_id = "1337"
  type = "anvil"

[[nodesets]]
  name = "don"
  nodes = 2
  override_mode = "each"

  [nodesets.db]
    image = "postgres:15.0"

  [[nodesets.node_specs]]
    [nodesets.node_specs.node]
      image = "public.ecr.aws/chainlink/chainlink:2.30.0"

  [[nodesets.node_specs]]
    [nodesets.node_specs.node]
      image = "public.ecr.aws/chainlink/chainlink:2.30.0"
`
	configPath := filepath.Join(tmpDir, "env.toml")
	err := os.WriteFile(configPath, []byte(configContent), 0644)
	require.NoError(t, err)

	// Change to tmpDir so Load can find the file
	originalDir, err := os.Getwd()
	require.NoError(t, err)
	defer os.Chdir(originalDir)
	err = os.Chdir(tmpDir)
	require.NoError(t, err)

	// Load the config
	cfg, err := Load[Cfg]([]string{"env.toml"})
	require.NoError(t, err)

	// Verify config is loaded correctly
	assert.Equal(t, float64(50), cfg.CLNodesFundingETH)
	assert.Equal(t, float64(50), cfg.CLNodesFundingLink)

	// Verify NodeConfigOverrides is empty (not set)
	assert.Empty(t, cfg.NodeConfigOverrides)

	// Verify blockchains
	require.Len(t, cfg.Blockchains, 1)
	assert.Equal(t, "1337", cfg.Blockchains[0].ChainID)

	// Verify node specs have no overrides
	require.Len(t, cfg.NodeSets, 1)
	require.Len(t, cfg.NodeSets[0].NodeSpecs, 2)
	assert.Empty(t, cfg.NodeSets[0].NodeSpecs[0].Node.TestConfigOverrides)
	assert.Empty(t, cfg.NodeSets[0].NodeSpecs[1].Node.TestConfigOverrides)
}
