package fastcurse

import (
	"context"
	"testing"

	"github.com/Masterminds/semver/v3"
	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"

	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
)

var (
	testChain1Selector = chainsel.TEST_90000001.Selector
	testChain2Selector = chainsel.TEST_90000002.Selector
)

type testCurseSubjectAdapter struct{}

func (testCurseSubjectAdapter) SelectorToSubject(selector uint64) Subject {
	return GenericSelectorToSubject(selector)
}

func (testCurseSubjectAdapter) DeriveCurseAdapterVersion(_ cldf.Environment, _ uint64) (*semver.Version, error) {
	return semver.MustParse("1.6.0"), nil
}

type testCurseAdapter struct {
	subjectsAreCursed bool
}

func (a *testCurseAdapter) Initialize(_ cldf.Environment, _ uint64) error {
	return nil
}

func (a *testCurseAdapter) IsSubjectCursedOnChain(_ cldf.Environment, _ uint64, subject Subject) (bool, error) {
	if IfSubjectEqual(subject, GlobalCurseSubject()) {
		return false, nil
	}
	return a.subjectsAreCursed, nil
}

func (a *testCurseAdapter) IsChainConnectedToTargetChain(_ cldf.Environment, _, _ uint64) (bool, error) {
	return true, nil
}

func (a *testCurseAdapter) IsCurseEnabledForChain(_ cldf.Environment, _ uint64) (bool, error) {
	return true, nil
}

func (a *testCurseAdapter) SubjectToSelector(subject Subject) (uint64, error) {
	return GenericSubjectToSelector(subject)
}

func (a *testCurseAdapter) Curse() *cldf_ops.Sequence[CurseInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return testCurseSequence()
}

func (a *testCurseAdapter) Uncurse() *cldf_ops.Sequence[CurseInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return testCurseSequence()
}

func (a *testCurseAdapter) ListConnectedChains(_ cldf.Environment, _ uint64) ([]uint64, error) {
	return nil, nil
}

func testCurseSequence() *cldf_ops.Sequence[CurseInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return cldf_ops.NewSequence(
		"test-curse-sequence",
		semver.MustParse("1.0.0"),
		"test curse sequence",
		func(_ cldf_ops.Bundle, _ cldf_chain.BlockChains, _ CurseInput) (sequences.OnChainOutput, error) {
			return sequences.OnChainOutput{}, nil
		},
	)
}

func newTestCurseRegistry(subjectsAreCursed bool) *CurseRegistry {
	cr := newCurseRegistry()
	cr.RegisterNewCurse(CurseRegistryInput{
		CursingFamily:       chainsel.FamilyEVM,
		CursingVersion:      semver.MustParse("1.6.0"),
		CurseAdapter:        &testCurseAdapter{subjectsAreCursed: subjectsAreCursed},
		CurseSubjectAdapter: testCurseSubjectAdapter{},
	})
	return cr
}

func newTestEnvironment(t *testing.T) cldf.Environment {
	t.Helper()
	lggr := logger.Test(t)
	return cldf.Environment{
		Logger: lggr,
		OperationsBundle: cldf_ops.NewBundle(
			func() context.Context { return t.Context() },
			lggr,
			cldf_ops.NewMemoryReporter(),
		),
		BlockChains: cldf_chain.NewBlockChains(map[uint64]cldf_chain.BlockChain{}),
	}
}
