package changesets

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestResolveRemoveFeeTokensCfg(t *testing.T) {
	input := map[string]interface{}{
		"Cfg": map[string]interface{}{
			"ChainSels": []interface{}{
				"6892437333620424805",
				"14767482510784806043",
				"13264668187771770619",
				"16015286601757825753",
				"3478487238524512106",
			}},
		"MCMS": map[string]interface{}{
			"Description":          "Remove extra fee tokens from FQ2.0",
			"OverridePreviousRoot": true,
			"Qualifier":            "CLLCCIP",
			"TimelockAction":       "schedule",
			"TimelockDelay":        "5m0s",
			"ValidUntil":           "1893456000"},
	}
	cfg, err := resolveRemoveFeeTokensCfg(input)
	require.NoError(t, err)
	expectedCfg := input["Cfg"].(map[string]interface{})
	chainSels := make([]uint64, len(expectedCfg["ChainSels"].([]interface{})))
	for i, chainSel := range expectedCfg["ChainSels"].([]interface{}) {
		chainSels[i], err = strconv.ParseUint(chainSel.(string), 10, 64)
		require.NoError(t, err)
	}
	require.Equal(t, chainSels, cfg.ChainSels)
}
