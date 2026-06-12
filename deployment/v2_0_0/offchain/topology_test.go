package offchain

import (
	"fmt"
	"strconv"
	"testing"

	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/deployment/v2_0_0/adapters"
)

// stubChainFamily is a minimal ChainFamily test double used to verify how
// ValidateForEnvironment orchestrates the per-chain NOP check (production-only,
// per committee/pool, with error wrapping). It enforces an arbitrary test
// threshold rather than the real production minimum, which is owned and tested
// separately for each chain family. The embedded nil interface
// satisfies the rest of the ChainFamily contract; those methods are never
// called by ValidateForEnvironment.
type stubChainFamily struct {
	adapters.ChainFamily
	minNOPs int
}

func (s stubChainFamily) ValidateNOPsTopology(chainSelector string, nopCount int) error {
	if nopCount < s.minNOPs {
		return fmt.Errorf("chain %q rejected: %d NOPs below test minimum %d", chainSelector, nopCount, s.minNOPs)
	}
	return nil
}

func TestEnvironmentTopologyValidateForEnvironment_ProductionMinimumNOPs(t *testing.T) {
	const testMinNOPs = 3
	chainSelector := strconv.FormatUint(chainsel.TEST_90000001.Selector, 10)
	enough := testNOPAliases(testMinNOPs)
	tooFew := testNOPAliases(testMinNOPs - 1)

	tests := []struct {
		name             string
		envName          string
		committeeAliases []string
		poolAliases      []string
		wantErr          string
	}{
		{
			name:             "production committee below minimum is rejected and wrapped",
			envName:          "prod_mainnet",
			committeeAliases: tooFew,
			poolAliases:      enough,
			wantErr:          fmt.Sprintf(`committee "default" validation failed on chain %q`, chainSelector),
		},
		{
			name:             "production executor pool below minimum is rejected and wrapped",
			envName:          "prod_testnet",
			committeeAliases: enough,
			poolAliases:      tooFew,
			wantErr:          fmt.Sprintf(`executor pool "default" validation failed on chain %q`, chainSelector),
		},
		{
			name:             "production topology meeting the minimum passes",
			envName:          "prod_mainnet",
			committeeAliases: enough,
			poolAliases:      enough,
		},
		{
			name:             "non production environment skips the chain family check",
			envName:          "test",
			committeeAliases: []string{"nop-1"},
			poolAliases:      []string{"nop-1"},
		},
		{
			name:             "duplicate committee aliases do not count toward minimum",
			envName:          "prod_mainnet",
			committeeAliases: append(testNOPAliases(testMinNOPs-1), "nop-1"),
			poolAliases:      enough,
			wantErr:          fmt.Sprintf(`committee "default" validation failed on chain %q`, chainSelector),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			topology := testEnvironmentTopology(chainSelector, tt.committeeAliases, tt.poolAliases)
			registry := adapters.NewChainFamilyRegistry()
			registry.RegisterChainFamily(chainsel.FamilyEVM, stubChainFamily{minNOPs: testMinNOPs})

			err := topology.ValidateForEnvironment(tt.envName, registry)
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

func testEnvironmentTopology(chainSelector string, committeeAliases []string, poolAliases []string) *EnvironmentTopology {
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
						chainSelector: {NOPAliases: committeeAliases, Threshold: 1},
					},
				},
			},
		},
		ExecutorPools: map[string]ExecutorPoolConfig{
			"default": {
				ChainConfigs: map[string]ChainExecutorPoolConfig{
					chainSelector: {NOPAliases: poolAliases},
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
