package fastcurse

import (
	"sync"
	"testing"

	"github.com/Masterminds/semver/v3"
	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/stretchr/testify/require"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"

	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
)

// capturingCurseAdapter records the CurseInput received by each of its sequences so
// tests can assert on what the changesets passed down to the family adapter.
type capturingCurseAdapter struct {
	testCurseAdapter

	mu         sync.Mutex
	curseIns   []CurseInput
	uncurseIns []CurseInput
}

func (a *capturingCurseAdapter) Curse() *cldf_ops.Sequence[CurseInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return cldf_ops.NewSequence(
		"test-capturing-curse-sequence",
		semver.MustParse("1.0.0"),
		"test curse sequence",
		func(_ cldf_ops.Bundle, _ cldf_chain.BlockChains, in CurseInput) (sequences.OnChainOutput, error) {
			a.mu.Lock()
			defer a.mu.Unlock()
			a.curseIns = append(a.curseIns, in)
			return sequences.OnChainOutput{}, nil
		},
	)
}

func (a *capturingCurseAdapter) Uncurse() *cldf_ops.Sequence[CurseInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return cldf_ops.NewSequence(
		"test-capturing-uncurse-sequence",
		semver.MustParse("1.0.0"),
		"test uncurse sequence",
		func(_ cldf_ops.Bundle, _ cldf_chain.BlockChains, in CurseInput) (sequences.OnChainOutput, error) {
			a.mu.Lock()
			defer a.mu.Unlock()
			a.uncurseIns = append(a.uncurseIns, in)
			return sequences.OnChainOutput{}, nil
		},
	)
}

// TestCurseChangesetPropagatesMCMSQualifier asserts that both the curse and uncurse
// changesets carry the MCMS qualifier from RMNCurseConfig.MCMS into the CurseInput
// handed to the family adapter. Families whose RMN curse entrypoint takes an explicit
// caller argument (Stellar) need it to resolve the executing timelock; a missed
// construction site would silently strand those families on the fallback route.
func TestCurseChangesetPropagatesMCMSQualifier(t *testing.T) {
	subject := GenericSelectorToSubject(testChain2Selector)
	laneAction := func() CurseActionInput {
		return CurseActionInput{
			ChainSelector: testChain1Selector,
			Version:       semver.MustParse("1.6.0"),
			Subject:       &subject,
		}
	}

	tests := []struct {
		name      string
		qualifier string
	}{
		{name: "qualifier set", qualifier: "RMNMCMS"},
		{name: "qualifier empty on direct runs", qualifier: ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// The adapter reports subjects as uncursed for the curse run and cursed
			// for the uncurse run; both changesets must be exercised.
			env := newTestEnvironment(t)

			curseAdapter := &capturingCurseAdapter{testCurseAdapter: testCurseAdapter{subjectsAreCursed: false}}
			curseRegistry := newTestCurseRegistryWithCaptureAdapter(curseAdapter)
			_, err := CurseChangeset(curseRegistry, changesets.GetRegistry()).Apply(env, RMNCurseConfig{
				CurseActions: []CurseActionInput{laneAction()},
				MCMS:         mcms.Input{Qualifier: tt.qualifier},
			})
			require.NoError(t, err)
			require.Len(t, curseAdapter.curseIns, 1, "curse changeset must invoke the adapter's sequence exactly once")
			require.Equal(t, tt.qualifier, curseAdapter.curseIns[0].MCMSQualifier)
			require.Equal(t, testChain1Selector, curseAdapter.curseIns[0].ChainSelector)

			uncurseAdapter := &capturingCurseAdapter{testCurseAdapter: testCurseAdapter{subjectsAreCursed: true}}
			uncurseRegistry := newTestCurseRegistryWithCaptureAdapter(uncurseAdapter)
			_, err = UncurseChangeset(uncurseRegistry, changesets.GetRegistry()).Apply(env, RMNCurseConfig{
				CurseActions: []CurseActionInput{laneAction()},
				MCMS:         mcms.Input{Qualifier: tt.qualifier},
			})
			require.NoError(t, err)
			require.Len(t, uncurseAdapter.uncurseIns, 1, "uncurse changeset must invoke the adapter's sequence exactly once")
			require.Equal(t, tt.qualifier, uncurseAdapter.uncurseIns[0].MCMSQualifier)
			require.Equal(t, testChain1Selector, uncurseAdapter.uncurseIns[0].ChainSelector)
		})
	}
}

func newTestCurseRegistryWithCaptureAdapter(adapter *capturingCurseAdapter) *CurseRegistry {
	cr := newCurseRegistry()
	cr.RegisterNewCurse(CurseRegistryInput{
		CursingFamily:       chainsel.FamilyEVM,
		CursingVersion:      semver.MustParse("1.6.0"),
		CurseAdapter:        adapter,
		CurseSubjectAdapter: testCurseSubjectAdapter{},
	})
	return cr
}
