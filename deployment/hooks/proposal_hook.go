package hooks

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/Masterminds/semver/v3"
	chain_selectors "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_changeset "github.com/smartcontractkit/chainlink-deployments-framework/engine/cld/changeset"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/cld/domain"

	"github.com/smartcontractkit/chainlink-ccip/deployment/testadapters"
)

const (
	postProposalCCIPSendHookName = "verify-ccip-send"
)

// PostProposalCCIPSendHookName is the hook name for changeset wiring and tests.
const PostProposalCCIPSendHookName = postProposalCCIPSendHookName

// PostProposalCCIPSend supplies chain-family metadata used by [verifyCCIPSend] to run CCIP send smoke checks
// after MCMS timelock execution.
type PostProposalCCIPSend interface {
	SkinSend(env cldf.Environment) bool
	PreSendValidation(env cldf.Environment, srcSel uint64) error
	SupportedFeeTokens(env cldf.Environment, srcSel uint64) ([]string, error)
	SupportedDestinations(env cldf.Environment, srcSel uint64) ([]uint64, error)
	AdapterVersionForLane(env cldf.Environment, srcSel, destSel uint64) (*semver.Version, error)
}

// PostProposalCCIPSendRegistry maps chain family to a provider (at most one per family).
type PostProposalCCIPSendRegistry struct {
	providers map[string]PostProposalCCIPSend
	mu        *sync.Mutex
}

func newPostProposalCCIPSendRegistry() *PostProposalCCIPSendRegistry {
	return &PostProposalCCIPSendRegistry{
		providers: make(map[string]PostProposalCCIPSend),
		mu:        &sync.Mutex{},
	}
}

var (
	singletonPostProposalCCIPSendRegistry *PostProposalCCIPSendRegistry
	oncePostProposalCCIPSendRegistry      sync.Once
)

func GetPostProposalCCIPSendRegistry() *PostProposalCCIPSendRegistry {
	oncePostProposalCCIPSendRegistry.Do(func() {
		singletonPostProposalCCIPSendRegistry = newPostProposalCCIPSendRegistry()
	})
	return singletonPostProposalCCIPSendRegistry
}

// Register registers a post-proposal CCIP send provider for a chain family.
// If a provider is already registered for the family, the first registration wins.
func (r *PostProposalCCIPSendRegistry) Register(family string, provider PostProposalCCIPSend) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.providers[family]; !exists {
		r.providers[family] = provider
	}
}

// Get returns the provider for the given chain family.
func (r *PostProposalCCIPSendRegistry) Get(family string) (PostProposalCCIPSend, bool) {
	r.mu.Lock()
	defer r.mu.Unlock()

	provider, ok := r.providers[family]
	return provider, ok
}

// ResetPostProposalCCIPSendRegistryForTest replaces the registry and resets sync.Once for isolated tests.
func ResetPostProposalCCIPSendRegistryForTest() {
	singletonPostProposalCCIPSendRegistry = newPostProposalCCIPSendRegistry()
	oncePostProposalCCIPSendRegistry = sync.Once{}
}

// GlobalPostProposalCCIPSendHook returns a post-proposal hook that verifies CCIP send paths using
// providers registered on [GetPostProposalCCIPSendRegistry]. Chain selectors are taken from successful
// timelock execution reports and grouped by chain family.
func GlobalPostProposalCCIPSendHook(dom domain.Domain) cldf_changeset.PostProposalHook {
	return cldf_changeset.PostProposalHook{
		HookDefinition: cldf_changeset.HookDefinition{
			Name:          postProposalCCIPSendHookName,
			FailurePolicy: cldf_changeset.Abort,
			Timeout:       5 * time.Minute,
		},
		Func: verifyCCIPSend(dom),
	}
}

func verifyCCIPSend(dom domain.Domain) cldf_changeset.PostProposalHookFunc {
	return func(ctx context.Context, params cldf_changeset.PostProposalHookParams) error {
		selectors := chainSelectorsFromSuccessfulTimelockReports(params.Reports)
		if len(selectors) == 0 {
			return nil
		}
		ds, err := dom.DataStoreByEnv(params.Env.Name)
		if err != nil {
			return fmt.Errorf("verify-ccip-send: datastore for env %q: %w", params.Env.Name, err)
		}
		deployEnv := cldf.Environment{
			Name:        params.Env.Name,
			Logger:      params.Env.Logger,
			BlockChains: params.Env.BlockChains,
			DataStore:   ds,
			GetContext: func() context.Context {
				return ctx
			},
		}
		byFamily := groupSelectorsByFamily(params.Env.Logger, selectors)
		var errs []error
		for family, sels := range byFamily {
			provider, ok := GetPostProposalCCIPSendRegistry().Get(family)
			if !ok {
				params.Env.Logger.Warnf("verify-ccip-send: no provider registered for chain family %s, skipping selectors %v",
					family, sels)
				continue
			}
			if err := runPostProposalCCIPSends(ctx, params.Env.Logger, &deployEnv, family, provider, sels); err != nil {
				errs = append(errs, fmt.Errorf("verify-ccip-send: family %s: %w", family, err))
			}
		}
		return errors.Join(errs...)
	}
}

func runPostProposalCCIPSends(
	ctx context.Context,
	lggr logger.Logger,
	env *cldf.Environment,
	family string,
	provider PostProposalCCIPSend,
	srcSelectors []uint64,
) error {
	if provider.SkinSend(*env) {
		lggr.Infof("verify-ccip-send: provider for family %s skips CCIP send, skipping send verify for selectors %v",
			family, srcSelectors)
		return nil
	}
	var errs []error
	for _, srcSel := range srcSelectors {
		err := provider.PreSendValidation(*env, srcSel)
		if err != nil {
			lggr.Warnf("verify-ccip-send: skip CCIP send verify from chain %d: pre-send validation: %v", srcSel, err)
			errs = append(errs, err)
			continue
		}

		dests, err := provider.SupportedDestinations(*env, srcSel)
		if err != nil {
			lggr.Warnf("verify-ccip-send: skip CCIP send verify from chain %d: supported destinations: %v", srcSel, err)
			errs = append(errs, err)
			continue
		}
		if len(dests) == 0 {
			lggr.Warnf("verify-ccip-send: no supported destinations from chain %d", srcSel)
			continue
		}

		feeTokens, err := provider.SupportedFeeTokens(*env, srcSel)
		if err != nil {
			lggr.Warnf("verify-ccip-send: fee tokens for chain %d: %v; using native only", srcSel, err)
			errs = append(errs, err)
			continue
		}
		if len(feeTokens) == 0 {
			feeTokens = []string{""}
		}

		for _, destSel := range dests {
			adapterVer, err := provider.AdapterVersionForLane(*env, srcSel, destSel)
			if err != nil {
				errs = append(errs, fmt.Errorf("verify-ccip-send: adapter version src %d dest %d: %w", srcSel, destSel, err))
				continue
			}
			factory, ok := testadapters.GetTestAdapterRegistry().GetTestAdapter(family, adapterVer)
			if !ok {
				errs = append(errs, fmt.Errorf("verify-ccip-send: no test adapter for family %s version %s",
					family, adapterVer.String()))
				continue
			}

			destFamily, err := chain_selectors.GetSelectorFamily(destSel)
			if err != nil {
				errs = append(errs, fmt.Errorf("verify-ccip-send: dest selector %d: %w", destSel, err))
				continue
			}
			var destAdapter testadapters.TestAdapter
			if destFamily == family {
				destAdapter = factory(env, destSel)
			} else {
				// source and dest selectors should have same version
				// there is no cross version lane supported yet
				destAdapterFactory, ok := testadapters.GetTestAdapterRegistry().GetTestAdapter(destFamily, adapterVer)
				if !ok {
					errs = append(errs, fmt.Errorf("verify-ccip-send: no test adapter for dest family %s version %s",
						destFamily, adapterVer.String()))
					continue
				}
				destAdapter = destAdapterFactory(env, destSel)
			}

			receiver := destAdapter.CCIPReceiver()

			srcAdapter := factory(env, srcSel)
			extraArgs, err := destAdapter.GetExtraArgs(receiver, family)
			if err != nil {
				errs = append(errs, fmt.Errorf("verify-ccip-send: extra args for src %d -> dest %d: %w", srcSel, destSel, err))
				continue
			}

			for _, feeTok := range feeTokens {
				msg, err := srcAdapter.BuildMessage(testadapters.MessageComponents{
					DestChainSelector: destSel,
					Receiver:          receiver,
					Data:              []byte("hello contract"),
					FeeToken:          feeTok,
					ExtraArgs:         extraArgs,
				})
				if err != nil {
					errs = append(errs, fmt.Errorf("verify-ccip-send: build message src %d -> dest %d fee %q: %w",
						srcSel, destSel, feeTok, err))
					continue
				}
				lggr.Infof("verify-ccip-send: sending CCIP verify message from chain %d to chain %d (feeToken=%q)",
					srcSel, destSel, feeTok)
				_, msgID, err := srcAdapter.SendMessage(ctx, destSel, msg)
				if err != nil {
					errs = append(errs, fmt.Errorf("verify-ccip-send: CCIP send from %d to %d fee %q: %w",
						srcSel, destSel, feeTok, err))
					continue
				}
				lggr.Infof("verify-ccip-send: Successful CCIP send message id %s (src=%d dest=%d fee=%q)", msgID, srcSel, destSel, feeTok)
			}
		}
	}
	return errors.Join(errs...)
}

func groupSelectorsByFamily(lggr logger.Logger, selectors []uint64) map[string][]uint64 {
	out := make(map[string][]uint64)
	for _, sel := range selectors {
		family, err := chain_selectors.GetSelectorFamily(sel)
		if err != nil {
			lggr.Warnf("verify-ccip-send: skip invalid chain selector %d: %v", sel, err)
			continue
		}
		out[family] = append(out[family], sel)
	}
	return out
}

func chainSelectorsFromSuccessfulTimelockReports(reports []cldf_changeset.MCMSTimelockExecuteReport) []uint64 {
	seen := make(map[uint64]struct{})
	var out []uint64
	for _, r := range reports {
		if r.Type != cldf_changeset.MCMSTimelockExecuteReportType {
			continue
		}
		if r.Status != "" && r.Status != "SUCCESS" {
			continue
		}
		if r.Error != "" {
			continue
		}
		sel := r.Input.ChainSelector
		if sel == 0 {
			continue
		}
		if _, dup := seen[sel]; dup {
			continue
		}
		seen[sel] = struct{}{}
		out = append(out, sel)
	}
	return out
}
