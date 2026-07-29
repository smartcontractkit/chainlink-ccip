package fastcurse

import (
	"testing"

	"github.com/stretchr/testify/require"

	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

func TestGroupRMNSubjectBySelector_SubjectBypassesConnectivityCheck(t *testing.T) {
	// The adapter reports every target chain as disconnected; a lane action relying on
	// SubjectChainSelector would fail, but an explicit Subject must bypass that check.
	adapter := &testCurseAdapter{disconnected: true}
	cr := newTestCurseRegistryWithAdapter(adapter)
	env := cldf.Environment{}

	firedrill := FiredrillSubject()
	grouped, err := cr.groupRMNSubjectBySelector(env, []CurseActionInput{
		subjectAction(t, testChain1Selector, firedrill, "1.6.0"),
	})
	require.NoError(t, err)
	require.Len(t, grouped, 1)
	require.Equal(t, []Subject{firedrill}, grouped[testChain1Selector].subjects)
}

func TestGroupRMNSubjectBySelector_SubjectWorksForRealSelectorDerivedValues(t *testing.T) {
	// Subject isn't special-cased to hardcoded/firedrill values - a real,
	// selector-derived subject passed explicitly behaves the same way.
	adapter := &testCurseAdapter{disconnected: true}
	cr := newTestCurseRegistryWithAdapter(adapter)
	env := cldf.Environment{}

	realSubject := GenericSelectorToSubject(testChain2Selector)
	grouped, err := cr.groupRMNSubjectBySelector(env, []CurseActionInput{
		subjectAction(t, testChain1Selector, realSubject, "1.6.0"),
	})
	require.NoError(t, err)
	require.Len(t, grouped, 1)
	require.Equal(t, []Subject{realSubject}, grouped[testChain1Selector].subjects)
}

func TestGroupRMNSubjectBySelector_SubjectMutualExclusivity(t *testing.T) {
	cr := newTestCurseRegistry(false)
	env := cldf.Environment{}
	firedrill := FiredrillSubject()

	t.Run("Subject and SubjectChainSelector both set", func(t *testing.T) {
		_, err := cr.groupRMNSubjectBySelector(env, []CurseActionInput{
			{
				ChainSelector:        testChain1Selector,
				SubjectChainSelector: testChain2Selector,
				Subject:              &firedrill,
				Version:              mustVersion(t, "1.6.0"),
			},
		})
		require.Error(t, err)
		require.Contains(t, err.Error(), "cannot have both Subject and SubjectChainSelector set")
	})

	t.Run("Subject and IsGlobalCurse both set", func(t *testing.T) {
		_, err := cr.groupRMNSubjectBySelector(env, []CurseActionInput{
			{
				ChainSelector: testChain1Selector,
				IsGlobalCurse: true,
				Subject:       &firedrill,
				Version:       mustVersion(t, "1.6.0"),
			},
		})
		require.Error(t, err)
		require.Contains(t, err.Error(), "cannot have both Subject and IsGlobalCurse set")
	})
}

func TestGroupRMNSubjectBySelector_SubjectDeduplicates(t *testing.T) {
	cr := newTestCurseRegistry(false)
	env := cldf.Environment{}
	firedrill := FiredrillSubject()

	grouped, err := cr.groupRMNSubjectBySelector(env, []CurseActionInput{
		subjectAction(t, testChain1Selector, firedrill, "1.6.0"),
		subjectAction(t, testChain1Selector, firedrill, "1.6.0"),
	})
	require.NoError(t, err)
	require.Equal(t, []Subject{firedrill}, grouped[testChain1Selector].subjects)
}
