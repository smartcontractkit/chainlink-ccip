package fastcurse

import (
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/stretchr/testify/require"

	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

func mustVersion(t *testing.T, v string) *semver.Version {
	t.Helper()
	ver, err := semver.NewVersion(v)
	require.NoError(t, err)
	return ver
}

func laneAction(t *testing.T, chain, subject uint64, version string) CurseActionInput {
	t.Helper()
	return CurseActionInput{
		ChainSelector:        chain,
		SubjectChainSelector: subject,
		Version:              mustVersion(t, version),
	}
}

func globalAction(t *testing.T, chain uint64, version string) CurseActionInput {
	t.Helper()
	return CurseActionInput{
		IsGlobalCurse: true,
		ChainSelector: chain,
		Version:       mustVersion(t, version),
	}
}

func TestValidateBidirectionalCursing(t *testing.T) {
	tests := []struct {
		name        string
		actions     []CurseActionInput
		expectError bool
	}{
		{
			name: "valid bidirectional v1.6 cursing",
			actions: []CurseActionInput{
				laneAction(t, 1, 2, "1.6.0"),
				laneAction(t, 2, 1, "1.6.0"),
			},
			expectError: false,
		},
		{
			name: "invalid unidirectional v1.6 cursing",
			actions: []CurseActionInput{
				laneAction(t, 1, 2, "1.6.0"),
			},
			expectError: true,
		},
		{
			name: "invalid v1.5 unidirectional cursing",
			actions: []CurseActionInput{
				laneAction(t, 1, 2, "1.5.0"),
			},
			expectError: true,
		},
		{
			name: "valid mixed version bidirectional cursing (v1.6 forward, v1.5 reverse)",
			actions: []CurseActionInput{
				laneAction(t, 1, 2, "1.6.0"),
				laneAction(t, 2, 1, "1.5.0"),
			},
			expectError: false,
		},
		{
			name: "global curse is not subject to bidirectional validation",
			actions: []CurseActionInput{
				globalAction(t, 1, "1.6.0"),
			},
			expectError: false,
		},
		{
			name: "multiple unidirectional lanes",
			actions: []CurseActionInput{
				laneAction(t, 1, 2, "1.6.0"),
				laneAction(t, 3, 4, "1.6.0"),
			},
			expectError: true,
		},
		{
			name: "invalid v1.7+ unidirectional cursing",
			actions: []CurseActionInput{
				laneAction(t, 1, 2, "1.7.0"),
			},
			expectError: true,
		},
		{
			name: "self-lane is ignored",
			actions: []CurseActionInput{
				laneAction(t, 1, 1, "1.6.0"),
			},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := RMNCurseConfig{CurseActions: tt.actions}
			err := validateBidirectionalLaneActions(cfg)
			if tt.expectError {
				require.Error(t, err)
				require.Contains(t, err.Error(), "unidirectional lane")
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestApplyCurse_EnforcesBidirectionalV16Validation(t *testing.T) {
	cfg := RMNCurseConfig{
		CurseActions: []CurseActionInput{
			laneAction(t, 1, 2, "1.6.0"),
		},
	}

	_, err := applyCurse(nil, nil)(cldf.Environment{}, cfg)
	require.Error(t, err)
	require.Contains(t, err.Error(), "curse validation failed")
	require.Contains(t, err.Error(), "unidirectional lane")
}

func TestApplyUncurse_EnforcesBidirectionalV16Validation(t *testing.T) {
	cfg := RMNCurseConfig{
		CurseActions: []CurseActionInput{
			laneAction(t, 1, 2, "1.6.0"),
		},
	}

	_, err := applyUncurse(nil, nil)(cldf.Environment{}, cfg)
	require.Error(t, err)
	require.Contains(t, err.Error(), "uncurse validation failed")
	require.Contains(t, err.Error(), "unidirectional lane")
}
