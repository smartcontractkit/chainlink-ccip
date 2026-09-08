package changesets

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/deployment/v2_0_0/offchain"
)

// executorPoolWithChains builds an executor pool whose chain_configs cover the given selectors.
func executorPoolWithChains(selectors ...uint64) offchain.ExecutorPoolConfig {
	chainConfigs := make(map[string]offchain.ChainExecutorPoolConfig, len(selectors))
	for _, selector := range selectors {
		chainConfigs[strconv.FormatUint(selector, 10)] = offchain.ChainExecutorPoolConfig{
			NOPAliases: []string{"nop-1"},
		}
	}
	return offchain.ExecutorPoolConfig{ChainConfigs: chainConfigs}
}

// The zkSync topology fix added both a committee chain_configs block and an executor
// pool chain_configs block. Nothing in the lane path reads executor pools, so a lane
// chain absent from every pool went unnoticed.
func TestValidateExecutorPoolCoverage(t *testing.T) {
	chainA := uint64(100)
	chainB := uint64(200)

	tests := []struct {
		name     string
		pools    map[string]offchain.ExecutorPoolConfig
		chains   []uint64
		wantErr  error
		wantChan uint64
	}{
		{
			name:   "Success - single pool covers every lane chain",
			pools:  map[string]offchain.ExecutorPoolConfig{"default": executorPoolWithChains(chainA, chainB)},
			chains: []uint64{chainA, chainB},
		},
		{
			name: "Success - coverage spread across multiple pools",
			pools: map[string]offchain.ExecutorPoolConfig{
				"default": executorPoolWithChains(chainA),
				"premium": executorPoolWithChains(chainB),
			},
			chains: []uint64{chainA, chainB},
		},
		{
			name:   "Success - topology declares no executor pools",
			pools:  nil,
			chains: []uint64{chainA, chainB},
		},
		{
			name:     "Failure - lane chain missing from every pool",
			pools:    map[string]offchain.ExecutorPoolConfig{"default": executorPoolWithChains(chainA)},
			chains:   []uint64{chainA, chainB},
			wantErr:  ErrChainMissingFromExecutorPools,
			wantChan: chainB,
		},
		{
			name:     "Failure - pool exists but covers an unrelated chain",
			pools:    map[string]offchain.ExecutorPoolConfig{"default": executorPoolWithChains(999)},
			chains:   []uint64{chainA},
			wantErr:  ErrChainMissingFromExecutorPools,
			wantChan: chainA,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := validateExecutorPoolCoverage(&offchain.EnvironmentTopology{ExecutorPools: tc.pools}, tc.chains)
			if tc.wantErr != nil {
				require.ErrorIs(t, err, tc.wantErr)
				require.Contains(t, err.Error(), strconv.FormatUint(tc.wantChan, 10))
				return
			}
			require.NoError(t, err)
		})
	}
}
