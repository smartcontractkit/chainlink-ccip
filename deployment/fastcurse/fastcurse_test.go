package fastcurse

import (
	"errors"
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
)

func hasLaneAction(actions []CurseActionInput, chainSelector, subjectChainSelector uint64) bool {
	for _, a := range actions {
		if !a.IsGlobalCurse && a.ChainSelector == chainSelector && a.SubjectChainSelector == subjectChainSelector {
			return true
		}
	}
	return false
}

func TestGlobalCurse_SkipsLaneForUnsupportedConnectedChain(t *testing.T) {
	tests := []struct {
		name           string
		adapter        *testCurseAdapter
		subjectAdapter CurseSubjectAdapter
	}{
		{
			name: "IsChainSupported returns false",
			adapter: &testCurseAdapter{
				connectedChains:    []uint64{testChain2Selector},
				unsupportedTargets: map[uint64]bool{testChain1Selector: true},
			},
		},
		{
			name: "connectivity check errors",
			adapter: &testCurseAdapter{
				connectedChains: []uint64{testChain2Selector},
				connectivityErr: errors.New("router unreachable"),
			},
		},
		{
			name: "DeriveCurseAdapterVersion fails (chain not in environment)",
			adapter: &testCurseAdapter{
				connectedChains: []uint64{testChain2Selector},
			},
			subjectAdapter: testCurseSubjectAdapter{
				deriveVersionErr: errors.New("no EVM chain found for selector"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cr := newTestCurseRegistryWithAdapter(tt.adapter)
			if tt.subjectAdapter != nil {
				cr = newTestCurseRegistryWithAdapterAndSubject(tt.adapter, tt.subjectAdapter)
			}
			env := newTestEnvironment(t)

			cfg := GlobalCurseOnNetworkInput{
				ChainSelectors: map[uint64]*semver.Version{
					testChain1Selector: mustVersion(t, "1.6.0"),
				},
			}
			curseCfg, err := formCurseConfigForGlobalCurse(env, cr, cfg)
			require.NoError(t, err)
			require.True(t, containsGlobalCurseAction(curseCfg.CurseActions, testChain1Selector),
				"global curse action must still be present")
			require.False(t, hasLaneAction(curseCfg.CurseActions, testChain2Selector, testChain1Selector),
				"lane action for unsupported chain must be skipped")
			require.False(t, hasLaneAction(curseCfg.CurseActions, testChain1Selector, testChain2Selector),
				"reverse lane action for unsupported chain must be skipped")

			_, err = GloballyCurseChainChangeset(cr, changesets.GetRegistry()).Apply(env, cfg)
			require.NoError(t, err)
		})
	}
}

func TestFormCurseConfigForGlobalCurse_IncludesLaneForFullySupportedConnectedChain(t *testing.T) {
	adapter := &testCurseAdapter{
		connectedChains: []uint64{testChain2Selector},
	}
	cr := newTestCurseRegistryWithAdapter(adapter)
	env := newTestEnvironment(t)

	cfg := GlobalCurseOnNetworkInput{
		ChainSelectors: map[uint64]*semver.Version{
			testChain1Selector: mustVersion(t, "1.6.0"),
		},
	}
	curseCfg, err := formCurseConfigForGlobalCurse(env, cr, cfg)
	require.NoError(t, err)
	require.True(t, containsGlobalCurseAction(curseCfg.CurseActions, testChain1Selector),
		"global curse action must be present")
	require.True(t, hasLaneAction(curseCfg.CurseActions, testChain2Selector, testChain1Selector),
		"lane action for supported chain must be present")
	require.True(t, hasLaneAction(curseCfg.CurseActions, testChain1Selector, testChain2Selector),
		"reverse lane action for supported chain must be present")
}

func containsGlobalCurseAction(actions []CurseActionInput, chainSelector uint64) bool {
	for _, a := range actions {
		if a.IsGlobalCurse && a.ChainSelector == chainSelector {
			return true
		}
	}
	return false
}
