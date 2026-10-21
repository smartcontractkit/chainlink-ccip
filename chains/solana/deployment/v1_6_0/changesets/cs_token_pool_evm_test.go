package changesets

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

func TestE2ETokenPoolRejectsEVMStep(t *testing.T) {
	t.Parallel()
	// existing pipeline inputs carry the field as null
	var cfg E2ETokenPoolConfig
	require.NoError(t, json.Unmarshal([]byte(`{"ConfigureTokenPoolContractsChangesets": null}`), &cfg))
	require.Empty(t, cfg.ConfigureTokenPoolContractsChangesets)

	require.NoError(t, json.Unmarshal([]byte(`{"ConfigureTokenPoolContractsChangesets": [{"MCMS": null}]}`), &cfg))
	_, err := E2ETokenPool(cldf.Environment{Logger: logger.Test(t)}, cfg)
	require.ErrorContains(t, err, "ConfigureTokenPoolContractsChangesets is no longer supported")
}

func TestLeftPad32(t *testing.T) {
	t.Parallel()
	require.Equal(t, append(make([]byte, 30), 1, 2), leftPad32([]byte{1, 2}))
	full := make([]byte, 32)
	full[0] = 7
	require.Equal(t, full, leftPad32(full))
}
