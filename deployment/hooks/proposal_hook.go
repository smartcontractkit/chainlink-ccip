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
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
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
	SkipSend(env cldf_changeset.ProposalHookEnv) bool
	PreSendValidation(env cldf.Environment, srcSel uint64) error
	SupportedFeeTokens(env cldf.Environment, srcSel uint64, forkContext cldf_changeset.ForkContext) ([]string, error)
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

// GetPostProposalCCIPSendRegistry returns the singleton registry used by
// family-specific post-proposal CCIP send providers.
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
			Timeout:       15 * time.Minute,
		},
		Func: verifyCCIPSend(dom),
	}
}

// verifyCCIPSend builds the post-proposal function that runs CCIP send
// verification across all families represented in the hook environment.
func verifyCCIPSend(dom domain.Domain) cldf_changeset.PostProposalHookFunc {
	return func(ctx context.Context, params cldf_changeset.PostProposalHookParams) error {
		// Hook runs only against chains involved in the proposal execution.
		selectors := params.Env.BlockChains.ListChainSelectors()
		if len(selectors) == 0 {
			return nil
		}
		ds, err := dom.DataStoreByEnv(params.Env.Name)
		if err != nil {
			return fmt.Errorf("verify-ccip-send: datastore for env %q: %w", params.Env.Name, err)
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
			if err := runPostProposalCCIPSends(ctx, params.Env.Logger, params.Env, ds, family, provider, sels); err != nil {
				errs = append(errs, fmt.Errorf("verify-ccip-send: family %s: %w", family, err))
			}
		}
		return errors.Join(errs...)
	}
}

// runPostProposalCCIPSends executes CCIP send verification probes for each
// source selector, destination selector, and supported fee token.
func runPostProposalCCIPSends(
	ctx context.Context,
	lggr logger.Logger,
	hookEnv cldf_changeset.ProposalHookEnv,
	ds datastore.DataStore,
	family string,
	provider PostProposalCCIPSend,
	srcSelectors []uint64,
) error {
	if provider.SkipSend(hookEnv) {
		lggr.Infof("verify-ccip-send: provider for family %s skips CCIP send, skipping send verify for selectors %v",
			family, srcSelectors)
		return nil
	}
	var errs []error
	// Adapters and provider contracts consume cldf.Environment, so rebuild it from hook inputs.
	env := cldf.Environment{
		Name:        hookEnv.Name,
		Logger:      hookEnv.Logger,
		BlockChains: hookEnv.BlockChains,
		DataStore:   ds,
		GetContext: func() context.Context {
			return ctx
		},
	}
	for _, srcSel := range srcSelectors {
		err := provider.PreSendValidation(env, srcSel)
		if err != nil {
			lggr.Warnf("verify-ccip-send: skip CCIP send verify from chain %d: pre-send validation: %v", srcSel, err)
			errs = append(errs, err)
			continue
		}

		dests, err := provider.SupportedDestinations(env, srcSel)
		if err != nil {
			lggr.Warnf("verify-ccip-send: skip CCIP send verify from chain %d: supported destinations: %v", srcSel, err)
			errs = append(errs, err)
			continue
		}
		if len(dests) == 0 {
			lggr.Warnf("verify-ccip-send: no supported destinations from chain %d", srcSel)
			continue
		}

		feeTokens, err := provider.SupportedFeeTokens(env, srcSel, hookEnv.ForkContext)
		if err != nil {
			lggr.Warnf("verify-ccip-send: fee tokens for chain %d: %v", srcSel, err)
			errs = append(errs, err)
			continue
		}
		if len(feeTokens) == 0 {
			// Keep a native-fee send path even when no ERC20 fee token is discoverable.
			feeTokens = []string{""}
		}

		for _, destSel := range dests {
			func(destSel uint64) {
				endDestGroup := beginLogGroup(lggr, "verify-ccip-send: src=%d dest=%d", srcSel, destSel)
				defer endDestGroup()

				adapterVer, err := provider.AdapterVersionForLane(env, srcSel, destSel)
				if err != nil {
					lggr.Warnf("verify-ccip-send: failed to resolve adapter version src=%d dest=%d: %v", srcSel, destSel, err)
					errs = append(errs, fmt.Errorf("verify-ccip-send: adapter version src %d dest %d: %w", srcSel, destSel, err))
					return
				}
				lggr.Infof("verify-ccip-send: adapter version for src=%d dest=%d is %s", srcSel, destSel, adapterVer.String())

				factory, ok := testadapters.GetTestAdapterRegistry().GetTestAdapter(family, adapterVer)
				if !ok {
					env.Logger.Warnf("verify-ccip-send: no test adapter for family %s version %s, skipping src %d dest %d",
						family, adapterVer.String(), srcSel, destSel)
					return
				}

				destFamily, err := chain_selectors.GetSelectorFamily(destSel)
				if err != nil {
					lggr.Warnf("verify-ccip-send: failed to resolve destination family for dest=%d: %v", destSel, err)
					errs = append(errs, fmt.Errorf("verify-ccip-send: dest selector %d: %w", destSel, err))
					return
				}
				lggr.Infof("verify-ccip-send: destination family for dest=%d is %s", destSel, destFamily)

				var destAdapter testadapters.TestAdapterForFamily

				// if dest sel is not present in env we just load the family specific adapter with the selector, otherwise we load the full adapter with env
				if !env.BlockChains.Exists(destSel) {
					destAdapterFactory, ok := testadapters.GetTestAdapterRegistry().GetTestAdapterForFamily(destFamily, adapterVer)
					if !ok {
						env.Logger.Warnf("verify-ccip-send: no test adapter for dest family %s version %s, skipping src %d dest %d",
							destFamily, adapterVer.String(), srcSel, destSel)
						return
					}
					lggr.Infof("verify-ccip-send: destination selector %d not in env; using family-only adapter", destSel)
					destAdapter = destAdapterFactory(env.DataStore, destSel)
				} else {
					if family == destFamily {
						destAdapter = factory(&env, destSel)
					} else {
						// Backward compatibility: fall back to full adapters when available.
						fullDestFactory, hasFullAdapter := testadapters.GetTestAdapterRegistry().GetTestAdapter(destFamily, adapterVer)
						if !hasFullAdapter {
							lggr.Warnf("verify-ccip-send: missing full adapter for dest family %s version %s", destFamily, adapterVer.String())
							errs = append(errs, fmt.Errorf("verify-ccip-send: no test adapter for dest family %s version %s",
								destFamily, adapterVer.String()))
							return
						}
						lggr.Infof("verify-ccip-send: using cross-family full destination adapter for %s", destFamily)
						destAdapter = fullDestFactory(&env, destSel)
					}
				}

				receiver := destAdapter.CCIPReceiver()
				srcAdapter := factory(&env, srcSel)
				extraArgs, err := destAdapter.GetExtraArgs(receiver, family)
				if err != nil {
					lggr.Warnf("verify-ccip-send: failed to build extra args for src=%d dest=%d: %v", srcSel, destSel, err)
					errs = append(errs, fmt.Errorf("verify-ccip-send: extra args for src %d -> dest %d: %w", srcSel, destSel, err))
					return
				}
				lggr.Infof("verify-ccip-send: prepared message components for src=%d dest=%d (receiverLen=%d extraArgsLen=%d)",
					srcSel, destSel, len(receiver), len(extraArgs))

				for _, feeTok := range feeTokens {
					func(feeTok string) {
						endFeeTokenGroup := beginLogGroup(lggr,
							"verify-ccip-send: src=%d dest=%d feeToken=%s",
							srcSel, destSel, formatFeeTokenLogLabel(feeTok),
						)
						defer endFeeTokenGroup()

						// Send one probe per fee token to verify each available fee payment path.
						msg, err := srcAdapter.BuildMessage(testadapters.MessageComponents{
							DestChainSelector: destSel,
							Receiver:          receiver,
							Data:              []byte("hello contract"),
							FeeToken:          feeTok,
							ExtraArgs:         extraArgs,
						})
						if err != nil {
							lggr.Warnf("verify-ccip-send: failed building message src=%d dest=%d fee=%q: %v", srcSel, destSel, feeTok, err)
							errs = append(errs, fmt.Errorf("verify-ccip-send: build message src %d -> dest %d fee %q: %w",
								srcSel, destSel, feeTok, err))
							return
						}

						lggr.Infof("verify-ccip-send: sending CCIP verify message from chain %d to chain %d (feeToken=%q)",
							srcSel, destSel, feeTok)
						_, msgID, err := srcAdapter.SendMessage(ctx, destSel, msg)
						if err != nil {
							lggr.Warnf("verify-ccip-send: :x: failed CCIP send src=%d dest=%d fee=%q: %v",
								srcSel, destSel, feeTok, err)
							errs = append(errs, fmt.Errorf("verify-ccip-send: CCIP send from %d to %d fee %q: %w",
								srcSel, destSel, feeTok, err))
							return
						}
						lggr.Infof("verify-ccip-send: :white_check_mark: successful CCIP send message id %s (src=%d dest=%d fee=%q)",
							msgID, srcSel, destSel, feeTok)
					}(feeTok)
				}
			}(destSel)
		}
	}
	return errors.Join(errs...)
}

func beginLogGroup(lggr logger.Logger, format string, args ...any) func() {
	lggr.Infof("::group::%s", fmt.Sprintf(format, args...))
	return func() {
		lggr.Infof("::endgroup::")
	}
}

func formatFeeTokenLogLabel(feeToken string) string {
	if feeToken == "" {
		return "native"
	}
	return feeToken
}

// groupSelectorsByFamily groups selectors by chain family and drops invalid selectors.
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
