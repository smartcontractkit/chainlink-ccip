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

type testCurseSubjectAdapter struct {
	deriveVersionErr error
}

func (a testCurseSubjectAdapter) SelectorToSubject(selector uint64) Subject {
	return GenericSelectorToSubject(selector)
}

func (a testCurseSubjectAdapter) DeriveCurseAdapterVersion(_ cldf.Environment, _ uint64) (*semver.Version, error) {
	if a.deriveVersionErr != nil {
		return nil, a.deriveVersionErr
	}
	return semver.MustParse("1.6.0"), nil
}

type testCurseAdapter struct {
	subjectsAreCursed bool
	// disconnected, when true, makes IsChainConnectedToTargetChain report false for
	// every target chain, so tests can verify that a literal Subject bypasses the
	// connectivity check entirely.
	disconnected bool
	// connectedChains, when non-nil, is returned by ListConnectedChains to simulate
	// the lane discovery returning a specific set of chains.
	connectedChains []uint64
	// unsupportedTargets, when non-nil, makes IsChainConnectedToTargetChain report
	// false for calls where targetSel matches an entry, simulating a router whose
	// IsChainSupported returns false for that destination chain.
	unsupportedTargets map[uint64]bool
	// connectivityErr, when non-nil, makes IsChainConnectedToTargetChain return
	// an error, simulating an unreachable router.
	connectivityErr error
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

func (a *testCurseAdapter) IsChainConnectedToTargetChain(_ cldf.Environment, _ uint64, targetSel uint64) (bool, error) {
	if a.disconnected {
		return false, nil
	}
	if a.connectivityErr != nil {
		return false, a.connectivityErr
	}
	if a.unsupportedTargets != nil && a.unsupportedTargets[targetSel] {
		return false, nil
	}
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
	return a.connectedChains, nil
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
	return newTestCurseRegistryWithAdapter(&testCurseAdapter{subjectsAreCursed: subjectsAreCursed})
}

func newTestCurseRegistryWithAdapter(adapter *testCurseAdapter) *CurseRegistry {
	return newTestCurseRegistryWithAdapterAndSubject(adapter, testCurseSubjectAdapter{})
}

func newTestCurseRegistryWithAdapterAndSubject(adapter *testCurseAdapter, subjectAdapter CurseSubjectAdapter) *CurseRegistry {
	cr := newCurseRegistry()
	cr.RegisterNewCurse(CurseRegistryInput{
		CursingFamily:       chainsel.FamilyEVM,
		CursingVersion:      semver.MustParse("1.6.0"),
		CurseAdapter:        adapter,
		CurseSubjectAdapter: subjectAdapter,
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
