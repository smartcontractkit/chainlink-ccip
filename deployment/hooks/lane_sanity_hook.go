package hooks

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"sort"
	"strings"
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
	postProposalLaneSanityHookName = "lane-sanity-check"
)

// TransferToken carries a source-chain token address and a short label for logging.
type TransferToken struct {
	Address string
	Label   string
}

// LaneSanityTokens holds categorized transfer-token sets for a lane sanity check.
type LaneSanityTokens struct {
	// TestTokens are tokens deployed specifically for TestRouter testing,
	// identified by symbols like TEST_TOKEN_<selector>-<family>.
	TestTokens []TransferToken
	// ProdTokens are well-known production tokens: USDC, LBTC, WETH.
	// Populated once the lane has been promoted to the production router.
	ProdTokens []TransferToken
	// FallbackTokens are other tokens that have a TokenTransferFeeConfig set
	// in the FeeQuoter for the destination chain. Used when no ProdTokens are
	// available.
	FallbackTokens []TransferToken
}

// allTransferTokens returns every candidate in priority order:
// TestTokens → ProdTokens → FallbackTokens.
func (t *LaneSanityTokens) allTransferTokens() []TransferToken {
	var out []TransferToken
	out = append(out, t.TestTokens...)
	out = append(out, t.ProdTokens...)
	out = append(out, t.FallbackTokens...)
	return out
}

// PostProposalLaneSanity extends PostProposalCCIPSend with token-transfer
// discovery methods needed to run the full lane sanity check suite.
type PostProposalLaneSanity interface {
	PostProposalCCIPSend

	// SupportedTransferTokens returns categorized token sets for the given lane.
	SupportedTransferTokens(env cldf.Environment, srcSel, destSel uint64) (*LaneSanityTokens, error)

	// MockReceiverAddress returns the dest-chain-encoded receiver bytes for the
	// MockReceiver deployed on chainSel. Returns nil, nil when not deployed.
	MockReceiverAddress(env cldf.Environment, chainSel uint64) ([]byte, error)

	// FundAndApproveTransferToken verifies that the deployer on srcSel holds at
	// least one whole token unit of tokenAddress and approves the Router to spend
	// it. Returns the approved amount (1 unit = 10^decimals). Returns an error if
	// the deployer balance is insufficient — callers must ensure the deployer is
	// funded before running lane sanity checks.
	FundAndApproveTransferToken(
		ctx context.Context,
		env cldf.Environment,
		srcSel uint64,
		tokenAddress string,
	) (*big.Int, error)

	// GetMessageFee returns the fee that will be charged for the given built
	// message (as returned by BuildMessage), formatted as a human-readable string
	// (e.g. "1234567 units"). Returns an empty string when the fee cannot be
	// determined; the caller treats this as informational only.
	GetMessageFee(
		ctx context.Context,
		env cldf.Environment,
		srcSel, destSel uint64,
		msg any,
	) (string, error)
}

// PostProposalLaneSanityRegistry maps a chain-family string to a
// PostProposalLaneSanity provider (at most one entry per family).
type PostProposalLaneSanityRegistry struct {
	providers map[string]PostProposalLaneSanity
	mu        *sync.Mutex
}

func newPostProposalLaneSanityRegistry() *PostProposalLaneSanityRegistry {
	return &PostProposalLaneSanityRegistry{
		providers: make(map[string]PostProposalLaneSanity),
		mu:        &sync.Mutex{},
	}
}

var (
	singletonLaneSanityRegistry *PostProposalLaneSanityRegistry
	onceLaneSanityRegistry      sync.Once
)

// GetPostProposalLaneSanityRegistry returns the global singleton registry.
func GetPostProposalLaneSanityRegistry() *PostProposalLaneSanityRegistry {
	onceLaneSanityRegistry.Do(func() {
		singletonLaneSanityRegistry = newPostProposalLaneSanityRegistry()
	})
	return singletonLaneSanityRegistry
}

// Register registers a provider for the given chain family. First call wins.
func (r *PostProposalLaneSanityRegistry) Register(family string, provider PostProposalLaneSanity) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, exists := r.providers[family]; !exists {
		r.providers[family] = provider
	}
}

// Get returns the provider for the given chain family, if registered.
func (r *PostProposalLaneSanityRegistry) Get(family string) (PostProposalLaneSanity, bool) {
	r.mu.Lock()
	defer r.mu.Unlock()
	p, ok := r.providers[family]
	return p, ok
}

// ResetPostProposalLaneSanityRegistryForTest resets the singleton for isolated tests.
func ResetPostProposalLaneSanityRegistryForTest() {
	singletonLaneSanityRegistry = newPostProposalLaneSanityRegistry()
	onceLaneSanityRegistry = sync.Once{}
}

// GlobalPostProposalLaneSanityHook returns a post-proposal hook that runs the full
// lane sanity check suite across all chains touched by the proposal:
//
//  1. Message sends with each supported fee token.
//  2. Token transfer with the TEST token (TestRouter phase).
//  3. Token transfer with USDC / LBTC / WETH if present (ProdRouter phase),
//     falling back to tokens with a TokenTransferFeeConfig in the FeeQuoter.
//  4. Programmable Token Transfer (PTT) to the deployed MockReceiver.
//  5. Token transfer with an arbitrary data payload.
func GlobalPostProposalLaneSanityHook(dom domain.Domain) cldf_changeset.PostProposalHook {
	return cldf_changeset.PostProposalHook{
		HookDefinition: cldf_changeset.HookDefinition{
			Name:          postProposalLaneSanityHookName,
			FailurePolicy: cldf_changeset.Abort,
			Timeout:       30 * time.Minute,
		},
		Func: laneSanityHookFunc(dom),
	}
}

func laneSanityHookFunc(dom domain.Domain) cldf_changeset.PostProposalHookFunc {
	return func(ctx context.Context, params cldf_changeset.PostProposalHookParams) error {
		selectors := params.Env.BlockChains.ListChainSelectors()
		if len(selectors) == 0 {
			return nil
		}
		ds, err := dom.DataStoreByEnv(params.Env.Name)
		if err != nil {
			return fmt.Errorf("%s: datastore for env %q: %w", postProposalLaneSanityHookName, params.Env.Name, err)
		}

		env := cldf.Environment{
			Name:        params.Env.Name,
			Logger:      params.Env.Logger,
			BlockChains: params.Env.BlockChains,
			DataStore:   ds,
			GetContext:  func() context.Context { return ctx },
		}

		byFamily := groupSelectorsByFamily(params.Env.Logger, selectors)
		var errs []error
		for family, sels := range byFamily {
			provider, ok := GetPostProposalLaneSanityRegistry().Get(family)
			if !ok {
				params.Env.Logger.Warnf("%s: no provider for family %s, skipping %v",
					postProposalLaneSanityHookName, family, sels)
				continue
			}
			if err := runLaneSanityMessageSends(ctx, params.Env.Logger, env, family, provider, sels); err != nil {
				errs = append(errs, fmt.Errorf("%s: fee-token sends family %s: %w",
					postProposalLaneSanityHookName, family, err))
			}
			if err := runLaneSanityTokenChecks(ctx, params.Env.Logger, env, family, provider, sels); err != nil {
				errs = append(errs, fmt.Errorf("%s: token transfers family %s: %w",
					postProposalLaneSanityHookName, family, err))
			}
		}
		return errors.Join(errs...)
	}
}

// RunLaneSanityChecks is the CLI-invocable entry point for lane sanity checks.
// It runs the same scenarios as GlobalPostProposalLaneSanityHook against a real
// environment. The deployer must already hold sufficient token balances.
func RunLaneSanityChecks(ctx context.Context, lggr logger.Logger, env cldf.Environment, selectors []uint64) error {
	byFamily := groupSelectorsByFamily(lggr, selectors)
	var errs []error
	for family, sels := range byFamily {
		provider, ok := GetPostProposalLaneSanityRegistry().Get(family)
		if !ok {
			lggr.Warnf("%s: no provider for family %s, skipping %v",
				postProposalLaneSanityHookName, family, sels)
			continue
		}
		if err := runLaneSanityMessageSends(ctx, lggr, env, family, provider, sels); err != nil {
			errs = append(errs, fmt.Errorf("%s: fee-token sends family %s: %w",
				postProposalLaneSanityHookName, family, err))
		}
		if err := runLaneSanityTokenChecks(ctx, lggr, env, family, provider, sels); err != nil {
			errs = append(errs, fmt.Errorf("%s: token transfers family %s: %w",
				postProposalLaneSanityHookName, family, err))
		}
	}
	return errors.Join(errs...)
}

// runLaneSanityMessageSends sends one message per lane per supported fee token
// and logs the CCIP Explorer link for each successful send.
func runLaneSanityMessageSends(
	ctx context.Context,
	lggr logger.Logger,
	env cldf.Environment,
	family string,
	provider PostProposalLaneSanity,
	srcSelectors []uint64,
) error {
	var errs []error
	failed := make(map[laneFailureKey]map[string]struct{})

	for _, srcSel := range srcSelectors {
		if err := provider.PreSendValidation(env, srcSel); err != nil {
			lggr.Warnf("%s: skip message sends from %s: %v",
				postProposalLaneSanityHookName, chainLabel(srcSel), err)
			errs = append(errs, err)
			continue
		}

		// Discover all fee tokens for this source chain once.
		// Passing nil fork context requests real-environment discovery (no
		// impersonation). Fall back to native-only if discovery fails so that
		// sends are not blocked entirely.
		feeTokens, err := provider.SupportedFeeTokens(env, srcSel, nil)
		if err != nil {
			lggr.Warnf("%s: fee token discovery from %s: %v — falling back to native",
				postProposalLaneSanityHookName, chainLabel(srcSel), err)
			feeTokens = []string{""}
		}
		if len(feeTokens) == 0 {
			feeTokens = []string{""}
		}
		lggr.Infof("%s: fee tokens for %s: %v",
			postProposalLaneSanityHookName, chainLabel(srcSel), feeTokenLabels(feeTokens))

		dests, err := provider.SupportedDestinations(env, srcSel)
		if err != nil {
			lggr.Warnf("%s: supported destinations from %s: %v",
				postProposalLaneSanityHookName, chainLabel(srcSel), err)
			errs = append(errs, err)
			continue
		}

		for _, destSel := range dests {
			adapterVer, err := provider.AdapterVersionForLane(env, srcSel, destSel)
			if err != nil {
				lggr.Warnf("%s: adapter version %s→%s: %v",
					postProposalLaneSanityHookName, chainLabel(srcSel), chainLabel(destSel), err)
				addLaneFailureSummary(failed, srcSel, destSel, allFeeTokensSummaryLabel)
				continue
			}
			factory, ok := testadapters.GetTestAdapterRegistry().GetForkCCIPSendTestAdapter(family, adapterVer)
			if !ok {
				lggr.Warnf("%s: ⏭️ no adapter family=%s version=%s %s→%s",
					postProposalLaneSanityHookName, family, adapterVer, chainLabel(srcSel), chainLabel(destSel))
				continue
			}
			destFamily, err := chain_selectors.GetSelectorFamily(destSel)
			if err != nil {
				lggr.Warnf("%s: bad dest selector %d: %v", postProposalLaneSanityHookName, destSel, err)
				addLaneFailureSummary(failed, srcSel, destSel, allFeeTokensSummaryLabel)
				continue
			}
			destAdapter, skip := resolveLaneSanityDestAdapter(lggr, env, factory, family, destFamily, adapterVer, destSel)
			if skip {
				continue
			}

			receiver := destAdapter.CCIPReceiver()
			extraArgs, err := destAdapter.GetExtraArgs(receiver, family)
			if err != nil {
				lggr.Warnf("%s: extra args %s→%s: %v",
					postProposalLaneSanityHookName, chainLabel(srcSel), chainLabel(destSel), err)
				addLaneFailureSummary(failed, srcSel, destSel, allFeeTokensSummaryLabel)
				continue
			}

			srcAdapter := factory(&env, srcSel)

			// Send one message per fee token for this lane.
			for _, feeToken := range feeTokens {
				feeLabel := formatFeeTokenLogLabel(feeToken)

				msg, err := srcAdapter.BuildMessage(testadapters.MessageComponents{
					DestChainSelector: destSel,
					Receiver:          receiver,
					Data:              []byte("lane-sanity-check"),
					FeeToken:          feeToken,
					ExtraArgs:         extraArgs,
				})
				if err != nil {
					lggr.Warnf("%s: ❌ build message feeToken=%s %s→%s: %v",
						postProposalLaneSanityHookName, feeLabel, chainLabel(srcSel), chainLabel(destSel), err)
					addLaneFailureSummary(failed, srcSel, destSel, feeLabel)
					continue
				}

				fee, feeErr := provider.GetMessageFee(ctx, env, srcSel, destSel, msg)
				if feeErr != nil {
					lggr.Warnf("%s: fee query feeToken=%s %s→%s: %v (continuing)",
						postProposalLaneSanityHookName, feeLabel, chainLabel(srcSel), chainLabel(destSel), feeErr)
				}

				_, msgID, err := srcAdapter.SendMessage(ctx, destSel, msg)
				if err != nil {
					lggr.Warnf("%s: ❌ send failed feeToken=%s %s→%s: %v",
						postProposalLaneSanityHookName, feeLabel, chainLabel(srcSel), chainLabel(destSel), err)
					addLaneFailureSummary(failed, srcSel, destSel, feeLabel)
					continue
				}

				lggr.Infof("%s: ✅ message sent", postProposalLaneSanityHookName)
				lggr.Infof("  source:      %s", chainLabel(srcSel))
				lggr.Infof("  destination: %s", chainLabel(destSel))
				lggr.Infof("  fee token:   %s", feeLabel)
				if fee != "" {
					lggr.Infof("  fee charged: %s", fee)
				}
				lggr.Infof("  message ID:  %s", msgID)
				lggr.Infof("  explorer:    %s", ccipExplorerURL(msgID))
			}
		}
	}
	if len(failed) > 0 {
		errs = append(errs, fmt.Errorf("failed message probes: %s", buildLaneFailureSummary(failed)))
	}
	return errors.Join(errs...)
}

// feeTokenLabels returns display labels for a slice of fee token addresses.
func feeTokenLabels(tokens []string) []string {
	out := make([]string, len(tokens))
	for i, t := range tokens {
		out[i] = formatFeeTokenLogLabel(t)
	}
	return out
}

// runLaneSanityTokenChecks executes the four token-transfer scenarios for every
// lane pair discovered from srcSelectors. The deployer must already hold
// sufficient token balances on srcSel; an error is returned early if not.
func runLaneSanityTokenChecks(
	ctx context.Context,
	lggr logger.Logger,
	env cldf.Environment,
	family string,
	provider PostProposalLaneSanity,
	srcSelectors []uint64,
) error {
	var errs []error
	failed := make(map[laneFailureKey]map[string]struct{})

	for _, srcSel := range srcSelectors {
		if err := provider.PreSendValidation(env, srcSel); err != nil {
			lggr.Warnf("%s: skip token checks from %s: %v",
				postProposalLaneSanityHookName, chainLabel(srcSel), err)
			continue
		}
		dests, err := provider.SupportedDestinations(env, srcSel)
		if err != nil {
			lggr.Warnf("%s: supported destinations from %s: %v",
				postProposalLaneSanityHookName, chainLabel(srcSel), err)
			continue
		}

		for _, destSel := range dests {
			adapterVer, err := provider.AdapterVersionForLane(env, srcSel, destSel)
			if err != nil {
				lggr.Warnf("%s: adapter version %s→%s: %v",
					postProposalLaneSanityHookName, chainLabel(srcSel), chainLabel(destSel), err)
				addLaneFailureSummary(failed, srcSel, destSel, "all-tokens")
				continue
			}
			factory, ok := testadapters.GetTestAdapterRegistry().GetForkCCIPSendTestAdapter(family, adapterVer)
			if !ok {
				lggr.Warnf("%s: ⏭️ no adapter family=%s version=%s %s→%s",
					postProposalLaneSanityHookName, family, adapterVer, chainLabel(srcSel), chainLabel(destSel))
				continue
			}
			destFamily, err := chain_selectors.GetSelectorFamily(destSel)
			if err != nil {
				lggr.Warnf("%s: bad dest selector %d: %v", postProposalLaneSanityHookName, destSel, err)
				addLaneFailureSummary(failed, srcSel, destSel, "all-tokens")
				continue
			}
			destAdapter, skip := resolveLaneSanityDestAdapter(lggr, env, factory, family, destFamily, adapterVer, destSel)
			if skip {
				continue
			}

			srcAdapter := factory(&env, srcSel)
			ccipReceiver := destAdapter.CCIPReceiver()
			extraArgs, err := destAdapter.GetExtraArgs(ccipReceiver, family)
			if err != nil {
				lggr.Warnf("%s: extra args %s→%s: %v",
					postProposalLaneSanityHookName, chainLabel(srcSel), chainLabel(destSel), err)
				addLaneFailureSummary(failed, srcSel, destSel, "all-tokens")
				continue
			}

			tokens, err := provider.SupportedTransferTokens(env, srcSel, destSel)
			if err != nil {
				lggr.Warnf("%s: transfer tokens %s→%s: %v",
					postProposalLaneSanityHookName, chainLabel(srcSel), chainLabel(destSel), err)
				addLaneFailureSummary(failed, srcSel, destSel, "all-tokens")
				continue
			}

			mockReceiver, err := provider.MockReceiverAddress(env, destSel)
			if err != nil {
				lggr.Warnf("%s: mock receiver dest=%s: %v (PTT check skipped)",
					postProposalLaneSanityHookName, chainLabel(destSel), err)
			}

			// --- Scenario A: TEST token (TestRouter phase) ---
			if len(tokens.TestTokens) == 0 {
				lggr.Infof("%s: A-test-token: none configured %s→%s",
					postProposalLaneSanityHookName, chainLabel(srcSel), chainLabel(destSel))
			}
			for _, tok := range tokens.TestTokens {
				if err := runTokenTransferScenario(ctx, lggr, env, srcAdapter, srcSel, destSel,
					tok, ccipReceiver, extraArgs, nil, provider); err != nil {
					addLaneFailureSummary(failed, srcSel, destSel, "A:"+tok.Label)
				}
			}

			// --- Scenario B: USDC/LBTC/WETH, else fallback tokens ---
			prodOrFallback := tokens.ProdTokens
			scenLabel := "B-prod-token"
			if len(prodOrFallback) == 0 {
				prodOrFallback = tokens.FallbackTokens
				scenLabel = "B-fallback-token"
			}
			if len(prodOrFallback) == 0 {
				lggr.Infof("%s: B: no prod or fallback tokens configured %s→%s",
					postProposalLaneSanityHookName, chainLabel(srcSel), chainLabel(destSel))
			}
			for _, tok := range prodOrFallback {
				if err := runTokenTransferScenario(ctx, lggr, env, srcAdapter, srcSel, destSel,
					tok, ccipReceiver, extraArgs, nil, provider); err != nil {
					addLaneFailureSummary(failed, srcSel, destSel, scenLabel+":"+tok.Label)
				}
			}

			// --- Scenario C: PTT to MockReceiver ---
			if len(mockReceiver) == 0 {
				lggr.Infof("%s: C-ptt: no MockReceiver on dest=%s",
					postProposalLaneSanityHookName, chainLabel(destSel))
			} else {
				allToks := tokens.allTransferTokens()
				if len(allToks) == 0 {
					lggr.Infof("%s: C-ptt: no transfer tokens available %s→%s",
						postProposalLaneSanityHookName, chainLabel(srcSel), chainLabel(destSel))
				} else {
					tok := allToks[0]
					mockExtraArgs, mockErr := destAdapter.GetExtraArgs(mockReceiver, family)
					if mockErr != nil {
						lggr.Warnf("%s: C-ptt: extra args for MockReceiver dest=%s: %v",
							postProposalLaneSanityHookName, chainLabel(destSel), mockErr)
					} else {
						if err := runTokenTransferScenario(ctx, lggr, env, srcAdapter, srcSel, destSel,
							tok, mockReceiver, mockExtraArgs, []byte("lane-sanity-check-payload"), provider); err != nil {
							addLaneFailureSummary(failed, srcSel, destSel, "C-ptt:"+tok.Label)
						}
					}
				}
			}
		}
	}
	if len(failed) > 0 {
		errs = append(errs, fmt.Errorf("failed token-transfer probes: %s", buildLaneSanityFailureSummary(failed)))
	}
	return errors.Join(errs...)
}

// runTokenTransferScenario approves the token on srcSel, builds a transfer
// message (with optional data payload), sends it, and logs the result including
// the CCIP Explorer link. Returns a non-nil error on failure.
func runTokenTransferScenario(
	ctx context.Context,
	lggr logger.Logger,
	env cldf.Environment,
	srcAdapter testadapters.ForkCCIPSendTestAdapter,
	srcSel, destSel uint64,
	tok TransferToken,
	receiver []byte,
	extraArgs []byte,
	data []byte,
	provider PostProposalLaneSanity,
) error {
	amount, err := provider.FundAndApproveTransferToken(ctx, env, srcSel, tok.Address)
	if err != nil {
		lggr.Warnf("%s: ❌ fund/approve token=%s src=%s: %v",
			postProposalLaneSanityHookName, tok.Label, chainLabel(srcSel), err)
		return err
	}

	msg, err := srcAdapter.BuildMessage(testadapters.MessageComponents{
		DestChainSelector: destSel,
		Receiver:          receiver,
		Data:              data,
		ExtraArgs:         extraArgs,
		TokenAmounts:      []testadapters.TokenAmount{{Token: tok.Address, Amount: amount}},
	})
	if err != nil {
		lggr.Warnf("%s: ❌ build message token=%s %s→%s: %v",
			postProposalLaneSanityHookName, tok.Label, chainLabel(srcSel), chainLabel(destSel), err)
		return err
	}

	fee, feeErr := provider.GetMessageFee(ctx, env, srcSel, destSel, msg)
	if feeErr != nil {
		lggr.Warnf("%s: fee query token=%s %s→%s: %v (continuing)",
			postProposalLaneSanityHookName, tok.Label, chainLabel(srcSel), chainLabel(destSel), feeErr)
	}

	_, msgID, err := srcAdapter.SendMessage(ctx, destSel, msg)
	if err != nil {
		lggr.Warnf("%s: ❌ send failed token=%s %s→%s: %v",
			postProposalLaneSanityHookName, tok.Label, chainLabel(srcSel), chainLabel(destSel), err)
		return err
	}

	lggr.Infof("%s: ✅ token transfer sent", postProposalLaneSanityHookName)
	lggr.Infof("  source:      %s", chainLabel(srcSel))
	lggr.Infof("  destination: %s", chainLabel(destSel))
	lggr.Infof("  token:       %s | amount: %s", tok.Label, amount.String())
	if fee != "" {
		lggr.Infof("  fee charged: %s", fee)
	}
	lggr.Infof("  message ID:  %s", msgID)
	lggr.Infof("  explorer:    %s", ccipExplorerURL(msgID))
	return nil
}

// resolveLaneSanityDestAdapter returns the TestAdapterForFamily for destSel.
// Returns skip=true when no suitable adapter is registered (non-fatal).
func resolveLaneSanityDestAdapter(
	lggr logger.Logger,
	env cldf.Environment,
	factory testadapters.ForkCCIPSendTestAdapterFactory,
	srcFamily, destFamily string,
	adapterVer *semver.Version,
	destSel uint64,
) (testadapters.TestAdapterForFamily, bool) {
	if !env.BlockChains.Exists(destSel) {
		f, ok := testadapters.GetTestAdapterRegistry().GetTestAdapterForFamily(destFamily, adapterVer)
		if !ok {
			lggr.Warnf("%s: ⏭️ no dest adapter family=%s version=%s (dest=%s not in env)",
				postProposalLaneSanityHookName, destFamily, adapterVer, chainLabel(destSel))
			return nil, true
		}
		return f(env.DataStore, destSel), false
	}
	if srcFamily == destFamily {
		return factory(&env, destSel), false
	}
	f, ok := testadapters.GetTestAdapterRegistry().GetForkCCIPSendTestAdapter(destFamily, adapterVer)
	if !ok {
		lggr.Warnf("%s: ⏭️ no cross-family dest adapter family=%s version=%s (dest=%s)",
			postProposalLaneSanityHookName, destFamily, adapterVer, chainLabel(destSel))
		return nil, true
	}
	return f(&env, destSel), false
}

// ccipExplorerURL formats the CCIP Explorer deep-link for a message ID.
// Normalises the ID to the "0x…" form expected by the explorer regardless of
// whether the adapter returned it with or without the prefix.
func ccipExplorerURL(msgID string) string {
	if !strings.HasPrefix(strings.ToLower(msgID), "0x") {
		msgID = "0x" + msgID
	}
	return "https://ccip.chain.link/msg/" + msgID
}

// chainLabel returns "<name> (<selector>)" for a known chain selector, or
// "chain-<selector>" as a fallback.
func chainLabel(sel uint64) string {
	chain, ok := chain_selectors.ChainBySelector(sel)
	if !ok {
		return fmt.Sprintf("chain-%d", sel)
	}
	return fmt.Sprintf("%s (%d)", chain.Name, sel)
}

func buildLaneSanityFailureSummary(failed map[laneFailureKey]map[string]struct{}) string {
	keys := make([]laneFailureKey, 0, len(failed))
	for k := range failed {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		if keys[i].srcSel != keys[j].srcSel {
			return keys[i].srcSel < keys[j].srcSel
		}
		return keys[i].destSel < keys[j].destSel
	})
	parts := make([]string, 0, len(keys))
	for _, k := range keys {
		scenarios := make([]string, 0, len(failed[k]))
		for s := range failed[k] {
			scenarios = append(scenarios, s)
		}
		sort.Strings(scenarios)
		parts = append(parts, fmt.Sprintf("src=%d dest=%d scenarios=[%s]",
			k.srcSel, k.destSel, strings.Join(scenarios, ",")))
	}
	return strings.Join(parts, "; ")
}
