package fastcurse

import (
	"fmt"
	"slices"
	"sync"

	"github.com/Masterminds/semver/v3"
	chain_selectors "github.com/smartcontractkit/chain-selectors"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"

	"github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
)

var (
	singletonCurseRegistry *CurseRegistry
	curseRegistryOnce      sync.Once
)

type CurseAdapter interface {
	Initialize(e cldf.Environment, selector uint64) error
	// IsSubjectCursedOnChain has the default RMN behavior.
	// Returns true if subject is cursed or chain is globally cursed. False otherwise.
	IsSubjectCursedOnChain(e cldf.Environment, selector uint64, subject Subject) (bool, error)
	// IsChainConnectedToTargetChain returns true if the chain with selector `selector` is connected to the chain with selector `targetSel`
	// For example, in case of EVM chains, router.isChainSupported(targetSel) should return true when called on the chain with selector `selector`
	IsChainConnectedToTargetChain(e cldf.Environment, selector uint64, targetSel uint64) (bool, error)
	// IsCurseEnabledForChain returns true if the chain with selector `selector` supports cursing subjects
	// For example, in case of EVM 1.6 chains, this could check if the RMNRemote contract is deployed on the chain
	IsCurseEnabledForChain(e cldf.Environment, selector uint64) (bool, error)
	// SubjectToSelector converts a Subject to a chain selector
	SubjectToSelector(subject Subject) (uint64, error)
	// Curse returns the sequence to curse subjects on a chain
	Curse() *cldf_ops.Sequence[CurseInput, sequences.OnChainOutput, cldf_chain.BlockChains]
	// Uncurse returns the sequence to lift the curse on subjects on a chain
	Uncurse() *cldf_ops.Sequence[CurseInput, sequences.OnChainOutput, cldf_chain.BlockChains]
	// ListConnectedChains returns a slice of connected chain selectors
	// It is used to determine which chains needs to curse subjects derived from given selector
	// This is needed to put global curse on a given chain with selector `selector`,
	// so that all chains connected to it can curse the subject derived from `selector`
	ListConnectedChains(e cldf.Environment, selector uint64) ([]uint64, error)
}

type CurseSubjectAdapter interface {
	// SelectorToSubject converts a chain selector to a Subject
	SelectorToSubject(selector uint64) Subject
	// DeriveCurseAdapterVersion derives the curse adapter version to be used to curse the subject on a chain
	// For example, for EVM chains, this could derive the RMN version deployed on the chain with selector `selector`
	DeriveCurseAdapterVersion(e cldf.Environment, selector uint64) (*semver.Version, error)
}

type CurseRegistryInput struct {
	CursingFamily       string
	CursingVersion      *semver.Version
	SubjectFamily       string
	CurseAdapter        CurseAdapter
	CurseSubjectAdapter CurseSubjectAdapter
}

type CurseRegistry struct {
	mu            sync.Mutex
	CurseAdapters map[string]CurseAdapter
	CurseSubjects map[string]CurseSubjectAdapter
}

func newCurseRegistry() *CurseRegistry {
	return &CurseRegistry{
		mu:            sync.Mutex{},
		CurseAdapters: make(map[string]CurseAdapter),
		CurseSubjects: make(map[string]CurseSubjectAdapter),
	}
}

func GetCurseRegistry() *CurseRegistry {
	curseRegistryOnce.Do(func() {
		singletonCurseRegistry = newCurseRegistry()
	})
	return singletonCurseRegistry
}

func (c *CurseRegistry) RegisterNewCurse(
	in CurseRegistryInput,
) {
	c.mu.Lock()
	defer c.mu.Unlock()
	id := utils.NewRegistererID(in.CursingFamily, in.CursingVersion)
	if _, exists := c.CurseAdapters[id]; !exists {
		c.CurseAdapters[id] = in.CurseAdapter
	}

	if _, exists := c.CurseSubjects[in.SubjectFamily]; !exists {
		c.CurseSubjects[in.SubjectFamily] = in.CurseSubjectAdapter
	}
}

func (c *CurseRegistry) GetCurseAdapter(chainFamily string, version *semver.Version) (CurseAdapter, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	id := utils.NewRegistererID(chainFamily, version)
	ca, ok := c.CurseAdapters[id]
	return ca, ok
}

func (c *CurseRegistry) GetCurseSubjectAdapter(family string) (CurseSubjectAdapter, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	csa, ok := c.CurseSubjects[family]
	return csa, ok
}

func (cr *CurseRegistry) groupRMNSubjectBySelector(e cldf.Environment, rmnSubjects []CurseActionInput) (map[uint64]curseActionDetails, error) {
	grouped := make(map[uint64]curseActionDetails)
	for _, s := range rmnSubjects {
		if s.IsGlobalCurse && s.SubjectChainSelector != 0 {
			return nil, fmt.Errorf("invalid curse action input: cannot have both IsGlobalCurse true and SubjectChainSelector set")
		}
		family, err := chain_selectors.GetSelectorFamily(s.ChainSelector)
		if err != nil {
			return nil, err
		}
		adapter, ok := cr.GetCurseAdapter(family, s.Version)
		if !ok {
			return nil, fmt.Errorf("no curse adapter registered for chain family '%s' and RMN version '%s'",
				family, s.Version.String())
		}
		// Skip self-curse
		if s.SubjectChainSelector == s.ChainSelector {
			continue
		}
		err = adapter.Initialize(e, s.ChainSelector)
		if err != nil {
			return nil, fmt.Errorf("failed to initialize curse adapter for chain selector %d: %w", s.ChainSelector, err)
		}
		// Check if curse is enabled on chain
		cursable, err := adapter.IsCurseEnabledForChain(e, s.ChainSelector)
		if err != nil || !cursable {
			return nil, fmt.Errorf("chain %d is not cursable: %w", s.ChainSelector, err)
		}

		// Initialize slice for this chain if needed
		if _, ok := grouped[s.ChainSelector]; !ok {
			grouped[s.ChainSelector] = curseActionDetails{
				curseAdapter: adapter,
				subjects:     make([]Subject, 0),
			}
		}
		if s.IsGlobalCurse {
			grouped[s.ChainSelector] = curseActionDetails{
				curseAdapter: adapter,
				subjects:     []Subject{GlobalCurseSubject()},
			}
			continue
		}
		// find if target subject is connected if not global
		connected, err := adapter.IsChainConnectedToTargetChain(e, s.ChainSelector, s.SubjectChainSelector)
		if err != nil {
			return nil, fmt.Errorf("failed to check if chain selector %d is connected to chain %d: %w", s.ChainSelector, s.SubjectChainSelector, err)
		}
		if !connected {
			continue
		}
		// get the subject from the subject chain selector
		subjectFamily, err := chain_selectors.GetSelectorFamily(s.SubjectChainSelector)
		if err != nil {
			return nil, fmt.Errorf("failed to get chain family for subject selector %d: %w", s.SubjectChainSelector, err)
		}
		subjectAdapter, ok := cr.GetCurseSubjectAdapter(subjectFamily)
		if !ok {
			return nil, fmt.Errorf("no curse subject adapter registered for chain family '%s'", subjectFamily)
		}
		subjectToCurse := subjectAdapter.SelectorToSubject(s.SubjectChainSelector)
		// Ensure uniqueness
		if !slices.Contains(grouped[s.ChainSelector].subjects, subjectToCurse) {
			grouped[s.ChainSelector] = curseActionDetails{
				curseAdapter: adapter,
				subjects:     append(grouped[s.ChainSelector].subjects, subjectToCurse),
			}
		}
	}

	return grouped, nil
}
