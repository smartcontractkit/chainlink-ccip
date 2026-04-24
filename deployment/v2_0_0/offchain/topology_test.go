package offchain

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEnvironmentTopologyValidateForEnvironment_ProductionMinimumNOPs(t *testing.T) {
	sixteenAliases := testNOPAliases(16)
	fifteenAliases := testNOPAliases(15)

	tests := []struct {
		name             string
		envName          string
		committeeAliases []string
		poolAliases      []string
		wantErr          string
	}{
		{
			name:             "prod committee chain with fifteen unique NOPs fails",
			envName:          "prod_mainnet",
			committeeAliases: fifteenAliases,
			poolAliases:      sixteenAliases,
			wantErr:          `committee "default" chain "1" requires at least 16 unique NOPs`,
		},
		{
			name:             "prod committee chain with sixteen unique NOPs passes",
			envName:          "prod_mainnet",
			committeeAliases: sixteenAliases,
			poolAliases:      sixteenAliases,
		},
		{
			name:             "prod executor pool chain with fifteen unique NOPs fails",
			envName:          "prod_testnet",
			committeeAliases: sixteenAliases,
			poolAliases:      fifteenAliases,
			wantErr:          `executor pool "default" chain "1" requires at least 16 unique NOPs`,
		},
		{
			name:             "prod executor pool chain with sixteen unique NOPs passes",
			envName:          "prod_testnet",
			committeeAliases: sixteenAliases,
			poolAliases:      sixteenAliases,
		},
		{
			name:             "non prod chain with fewer than sixteen NOPs passes",
			envName:          "test",
			committeeAliases: []string{"nop-1"},
			poolAliases:      []string{"nop-1"},
		},
		{
			name:             "duplicate committee aliases do not count toward minimum",
			envName:          "prod_mainnet",
			committeeAliases: append(testNOPAliases(15), "nop-1"),
			poolAliases:      sixteenAliases,
			wantErr:          `committee "default" chain "1" requires at least 16 unique NOPs`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			topology := testEnvironmentTopology(tt.committeeAliases, tt.poolAliases)

			err := topology.ValidateForEnvironment(tt.envName)
			if tt.wantErr == "" {
				require.NoError(t, err)
				return
			}
			require.Error(t, err)
			assert.Contains(t, err.Error(), tt.wantErr)
		})
	}
}

func TestEnvironmentTopologyValidate_NilNOPTopology(t *testing.T) {
	topology := &EnvironmentTopology{IndexerAddress: []string{"http://indexer:8080"}}

	err := topology.Validate()

	require.Error(t, err)
	assert.Contains(t, err.Error(), "nop_topology is required")
}

func testEnvironmentTopology(committeeAliases []string, poolAliases []string) *EnvironmentTopology {
	nops := make([]NOPConfig, 0, len(committeeAliases)+len(poolAliases))
	seen := make(map[string]struct{}, len(committeeAliases)+len(poolAliases))
	for _, alias := range append(append([]string{}, committeeAliases...), poolAliases...) {
		if _, ok := seen[alias]; ok {
			continue
		}
		seen[alias] = struct{}{}
		nops = append(nops, NOPConfig{
			Alias: alias,
			Name:  alias + "-name",
		})
	}

	return &EnvironmentTopology{
		IndexerAddress: []string{"http://indexer:8080"},
		NOPTopology: &NOPTopology{
			NOPs: nops,
			Committees: map[string]CommitteeConfig{
				"default": {
					Qualifier: "default",
					Aggregators: []AggregatorConfig{
						{Name: "agg-1", Address: "http://aggregator:8080"},
					},
					ChainConfigs: map[string]ChainCommitteeConfig{
						"1": {NOPAliases: committeeAliases, Threshold: 1},
					},
				},
			},
		},
		ExecutorPools: map[string]ExecutorPoolConfig{
			"default": {
				ChainConfigs: map[string]ChainExecutorPoolConfig{
					"1": {NOPAliases: poolAliases},
				},
			},
		},
	}
}

func testNOPAliases(count int) []string {
	aliases := make([]string, count)
	for i := range aliases {
		aliases[i] = fmt.Sprintf("nop-%d", i+1)
	}
	return aliases
}
