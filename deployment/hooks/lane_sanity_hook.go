package hooks

import (
	"bufio"
	"context"
	"encoding/csv"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"math/big"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
	"unicode/utf8"

	"github.com/Masterminds/semver/v3"
	chain_selectors "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"

	"github.com/smartcontractkit/chainlink-ccip/deployment/testadapters"
)

const (
	laneSanityCheckName      = "lane-sanity-check"
	laneSanityPrivateKeyEnv  = "PRIVATE_KEY"
	laneSanityCSVOutputEnv   = "LANE_SANITY_CSV_OUTPUT"
	laneSanityDefaultCSVName = "lane-sanity-check-results.csv"
)

// PostProposalLaneSanity extends PostProposalCCIPSend with helpers used by the
// CLI lane sanity check flow: sender configuration, transfer-token discovery,
// receiver encoding, and token approval before CCIP sends.
type PostProposalLaneSanity interface {
	PostProposalCCIPSend

	// ApplySenderPrivateKey configures chains in env to send from senderKey.
	// When senderKey is empty the environment is returned unchanged.
	ApplySenderPrivateKey(
		ctx context.Context,
		lggr logger.Logger,
		env *cldf.Environment,
		senderKey string,
	) error

	// AvailableTransferTokens returns selectable transfer tokens for source→dest.
	// Keys are display labels; values are source-chain token addresses.
	AvailableTransferTokens(env cldf.Environment, source, dest uint64) (map[string]string, error)

	// EncodeReceiverAddress encodes receiverAddress for delivery on destSel.
	// Called only when an explicit receiver override is provided.
	EncodeReceiverAddress(env cldf.Environment, destSel uint64, receiverAddress string) ([]byte, error)

	// MockReceiverAddress returns the dest-chain-encoded receiver bytes for the
	// MockReceiver deployed on chainSel. Returns nil, nil when not deployed.
	MockReceiverAddress(env cldf.Environment, chainSel uint64) ([]byte, error)

	// FundAndApproveTransferToken verifies that the sender on srcSel holds at
	// least one whole token unit of tokenAddress and approves the Router to spend
	// it. Returns the approved amount (1 unit = 10^decimals). Returns an error if
	// the sender balance is insufficient — callers must ensure the sender is
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

// LaneSanityChainPair identifies an undirected lane. RunLaneSanityChecks expands
// each pair into both directions (A→B and B→A).
type LaneSanityChainPair struct {
	ChainA uint64
	ChainB uint64
}

type laneSanityCheckRequest struct {
	src  uint64
	dest uint64
}

// transferTokenSelector prompts the user to choose transfer tokens from the
// available set. Tests may replace this variable to avoid interactive stdin.
var transferTokenSelector = promptTransferTokenSelection

// laneSanityWriteResultsCSV writes consolidated results. Tests may override.
var laneSanityWriteResultsCSV = writeLaneSanityResultsCSV

type laneSanityCSVRecord struct {
	SourceChain   string
	DestChain     string
	FeeToken      string
	TransferToken string
	Data          string
	Fee           string
	ExplorerLink  string
}

type laneSanityResultCollector struct {
	mu      sync.Mutex
	records []laneSanityCSVRecord
}

func (c *laneSanityResultCollector) add(record laneSanityCSVRecord) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.records = append(c.records, record)
}

func (c *laneSanityResultCollector) snapshot() []laneSanityCSVRecord {
	c.mu.Lock()
	defer c.mu.Unlock()
	out := make([]laneSanityCSVRecord, len(c.records))
	copy(out, c.records)
	return out
}

// RunLaneSanityChecks is the CLI entry point for lane sanity checks. For each
// chain pair it exercises both directions: fee-token message sends, interactive
// transfer-token selection, basic transfers, and one PTT when a MockReceiver exists.
// Sender key resolution: PRIVATE_KEY env when set, otherwise each chain deployer key.
// Successful sends are consolidated into a CSV (LANE_SANITY_CSV_OUTPUT or default filename).
func RunLaneSanityChecks(
	ctx context.Context,
	lggr logger.Logger,
	env cldf.Environment,
	requests []LaneSanityChainPair,
	receiver string,
) error {
	if len(requests) == 0 {
		return nil
	}

	bidirectionalReq := expandBidirectionalRequests(requests)

	err := applySenderPrivateKeyForRequests(ctx, lggr, &env, bidirectionalReq)
	if err != nil {
		return fmt.Errorf("%s: sender key: %w", laneSanityCheckName, err)
	}

	collector := &laneSanityResultCollector{}
	var errs []error
	for _, req := range bidirectionalReq {
		family, err := chain_selectors.GetSelectorFamily(req.src)
		if err != nil {
			errs = append(errs, fmt.Errorf("%s: src selector %d: %w", laneSanityCheckName, req.src, err))
			continue
		}
		provider, ok := GetPostProposalLaneSanityRegistry().Get(family)
		if !ok {
			lggr.Warnf("%s: no provider for family %s, skipping %s→%s",
				laneSanityCheckName, family, chainLabel(req.src), chainLabel(req.dest))
			continue
		}
		if err := runLaneSanityMessageSendWithAllFeeTokens(ctx, lggr, env, family, provider, req.src, req.dest, receiver, collector); err != nil {
			errs = append(errs, fmt.Errorf("%s: fee-token sends %s→%s: %w",
				laneSanityCheckName, chainLabel(req.src), chainLabel(req.dest), err))
		}
		if err := runLaneSanityTokenChecksForPair(ctx, lggr, env, family, provider, req.src, req.dest, receiver, collector); err != nil {
			errs = append(errs, fmt.Errorf("%s: token transfers %s→%s: %w",
				laneSanityCheckName, chainLabel(req.src), chainLabel(req.dest), err))
		}
	}

	if csvErr := laneSanityWriteResultsCSV(lggr, collector.snapshot()); csvErr != nil {
		errs = append(errs, fmt.Errorf("%s: write csv: %w", laneSanityCheckName, csvErr))
	}
	return errors.Join(errs...)
}

func expandBidirectionalRequests(pairs []LaneSanityChainPair) []laneSanityCheckRequest {
	if len(pairs) == 0 {
		return nil
	}
	out := make([]laneSanityCheckRequest, 0, len(pairs)*2)
	for _, pair := range pairs {
		out = append(out,
			laneSanityCheckRequest{src: pair.ChainA, dest: pair.ChainB},
			laneSanityCheckRequest{src: pair.ChainB, dest: pair.ChainA},
		)
	}
	return out
}

// applySenderPrivateKeyForRequests reads PRIVATE_KEY and delegates to each
// involved chain-family provider once. When PRIVATE_KEY is unset the environment
// is left unchanged so configured deployer keys are used.
func applySenderPrivateKeyForRequests(
	ctx context.Context,
	lggr logger.Logger,
	env *cldf.Environment,
	requests []laneSanityCheckRequest,
) error {
	senderKey := strings.TrimSpace(os.Getenv(laneSanityPrivateKeyEnv))
	if senderKey == "" {
		return nil
	}

	seenFamilies := make(map[string]struct{})
	for _, req := range requests {
		family, err := chain_selectors.GetSelectorFamily(req.src)
		if err != nil {
			return fmt.Errorf("src selector %d: %w", req.src, err)
		}
		if _, ok := seenFamilies[family]; ok {
			continue
		}
		seenFamilies[family] = struct{}{}

		provider, ok := GetPostProposalLaneSanityRegistry().Get(family)
		if !ok {
			return fmt.Errorf("no provider for family %s", family)
		}
		if len(env.BlockChains.ListChainSelectors(cldf_chain.WithFamily(family))) == 0 {
			return fmt.Errorf("no %s chains in environment", family)
		}
		err = provider.ApplySenderPrivateKey(ctx, lggr, env, senderKey)
		if err != nil {
			return err
		}
	}
	return nil
}

// runLaneSanityMessageSendWithAllFeeTokens sends one message per supported fee token on
// the directed srcSel→destSel lane and logs the CCIP Explorer link for each send.
func runLaneSanityMessageSendWithAllFeeTokens(
	ctx context.Context,
	lggr logger.Logger,
	env cldf.Environment,
	family string,
	provider PostProposalLaneSanity,
	srcSel, destSel uint64,
	receiverAddress string,
	collector *laneSanityResultCollector,
) error {
	var errs []error
	failed := make(map[laneFailureKey]map[string]struct{})

	if err := provider.PreSendValidation(env, srcSel); err != nil {
		lggr.Warnf("%s: skip message sends from %s: %v",
			laneSanityCheckName, chainLabel(srcSel), err)
		return err
	}

	// Passing nil fork context requests real-environment discovery (no impersonation).
	feeTokens, err := provider.SupportedFeeTokens(env, srcSel, nil)
	if err != nil {
		lggr.Warnf("%s: fee token discovery from %s: %v — falling back to native",
			laneSanityCheckName, chainLabel(srcSel), err)
		feeTokens = []string{""}
	}
	if len(feeTokens) == 0 {
		feeTokens = []string{""}
	}
	lggr.Infof("%s: fee tokens for %s→%s: %v",
		laneSanityCheckName, chainLabel(srcSel), chainLabel(destSel), feeTokenLabels(feeTokens))

	adapterVer, err := provider.AdapterVersionForLane(env, srcSel, destSel)
	if err != nil {
		lggr.Warnf("%s: adapter version %s→%s: %v",
			laneSanityCheckName, chainLabel(srcSel), chainLabel(destSel), err)
		addLaneFailureSummary(failed, srcSel, destSel, allFeeTokensSummaryLabel)
		return fmt.Errorf("failed message probes: %s", buildLaneFailureSummary(failed))
	}
	factory, ok := testadapters.GetTestAdapterRegistry().GetForkCCIPSendTestAdapter(family, adapterVer)
	if !ok {
		lggr.Warnf("%s: ⏭️ no adapter family=%s version=%s %s→%s",
			laneSanityCheckName, family, adapterVer, chainLabel(srcSel), chainLabel(destSel))
		return nil
	}
	destFamily, err := chain_selectors.GetSelectorFamily(destSel)
	if err != nil {
		lggr.Warnf("%s: bad dest selector %d: %v", laneSanityCheckName, destSel, err)
		addLaneFailureSummary(failed, srcSel, destSel, allFeeTokensSummaryLabel)
		return fmt.Errorf("failed message probes: %s", buildLaneFailureSummary(failed))
	}
	destAdapter, skip := resolveLaneSanityDestAdapter(lggr, env, factory, family, destFamily, adapterVer, destSel)
	if skip {
		return nil
	}

	receiver, err := resolveLaneSanityReceiver(lggr, provider, env, destAdapter, destSel, receiverAddress)
	if err != nil {
		lggr.Warnf("%s: receiver %s→%s: %v",
			laneSanityCheckName, chainLabel(srcSel), chainLabel(destSel), err)
		addLaneFailureSummary(failed, srcSel, destSel, allFeeTokensSummaryLabel)
		return fmt.Errorf("failed message probes: %s", buildLaneFailureSummary(failed))
	}
	extraArgs, err := destAdapter.GetExtraArgs(receiver, family)
	if err != nil {
		lggr.Warnf("%s: extra args %s→%s: %v",
			laneSanityCheckName, chainLabel(srcSel), chainLabel(destSel), err)
		addLaneFailureSummary(failed, srcSel, destSel, allFeeTokensSummaryLabel)
		return fmt.Errorf("failed message probes: %s", buildLaneFailureSummary(failed))
	}

	srcAdapter := factory(&env, srcSel)
	messageData := []byte("lane-sanity-check")

	for _, feeToken := range feeTokens {
		feeLabel := formatFeeTokenLogLabel(feeToken)

		msg, err := srcAdapter.BuildMessage(testadapters.MessageComponents{
			DestChainSelector: destSel,
			Receiver:          receiver,
			Data:              messageData,
			FeeToken:          feeToken,
			ExtraArgs:         extraArgs,
		})
		if err != nil {
			lggr.Warnf("%s: ❌ build message feeToken=%s %s→%s: %v",
				laneSanityCheckName, feeLabel, chainLabel(srcSel), chainLabel(destSel), err)
			addLaneFailureSummary(failed, srcSel, destSel, feeLabel)
			continue
		}

		fee, feeErr := provider.GetMessageFee(ctx, env, srcSel, destSel, msg)
		if feeErr != nil {
			lggr.Warnf("%s: fee query feeToken=%s %s→%s: %v (continuing)",
				laneSanityCheckName, feeLabel, chainLabel(srcSel), chainLabel(destSel), feeErr)
		}

		_, msgID, err := srcAdapter.SendMessage(ctx, destSel, msg)
		if err != nil {
			lggr.Warnf("%s: ❌ send failed feeToken=%s %s→%s: %v",
				laneSanityCheckName, feeLabel, chainLabel(srcSel), chainLabel(destSel), err)
			addLaneFailureSummary(failed, srcSel, destSel, feeLabel)
			continue
		}

		lggr.Infof("%s: ✅ message sent", laneSanityCheckName)
		lggr.Infof("  source:      %s", chainLabel(srcSel))
		lggr.Infof("  destination: %s", chainLabel(destSel))
		lggr.Infof("  fee token:   %s", feeLabel)
		lggr.Infof("  fee charged: %s", fee)
		lggr.Infof("  message ID:  %s", msgID)
		lggr.Infof("  explorer:    %s", ccipExplorerURL(msgID))
		collector.add(laneSanityCSVRecord{
			SourceChain:   chainName(srcSel),
			DestChain:     chainName(destSel),
			FeeToken:      feeLabel,
			TransferToken: "",
			Data:          formatDataForCSV(messageData),
			Fee:           fee,
			ExplorerLink:  ccipExplorerURL(msgID),
		})
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

// runLaneSanityTokenChecksForPair lists available transfer tokens, prompts the
// user to select tokens until they enter "done", then sends a CCIP transfer for
// each selection on the directed srcSel→destSel lane.
func runLaneSanityTokenChecksForPair(
	ctx context.Context,
	lggr logger.Logger,
	env cldf.Environment,
	family string,
	provider PostProposalLaneSanity,
	srcSel, destSel uint64,
	receiverAddress string,
	collector *laneSanityResultCollector,
) error {
	var errs []error
	failed := make(map[laneFailureKey]map[string]struct{})

	if err := provider.PreSendValidation(env, srcSel); err != nil {
		lggr.Warnf("%s: skip token checks from %s: %v",
			laneSanityCheckName, chainLabel(srcSel), err)
		return nil
	}

	adapterVer, err := provider.AdapterVersionForLane(env, srcSel, destSel)
	if err != nil {
		lggr.Warnf("%s: adapter version %s→%s: %v",
			laneSanityCheckName, chainLabel(srcSel), chainLabel(destSel), err)
		addLaneFailureSummary(failed, srcSel, destSel, "n/a")
		return fmt.Errorf("failed token-transfer probes: %s", buildLaneSanityFailureSummary(failed))
	}
	factory, ok := testadapters.GetTestAdapterRegistry().GetForkCCIPSendTestAdapter(family, adapterVer)
	if !ok {
		lggr.Warnf("%s: ⏭️ no adapter family=%s version=%s %s→%s",
			laneSanityCheckName, family, adapterVer, chainLabel(srcSel), chainLabel(destSel))
		return nil
	}
	destFamily, err := chain_selectors.GetSelectorFamily(destSel)
	if err != nil {
		lggr.Warnf("%s: bad dest selector %d: %v", laneSanityCheckName, destSel, err)
		addLaneFailureSummary(failed, srcSel, destSel, "n/a")
		return fmt.Errorf("failed token-transfer probes: %s", buildLaneSanityFailureSummary(failed))
	}
	destAdapter, skip := resolveLaneSanityDestAdapter(lggr, env, factory, family, destFamily, adapterVer, destSel)
	if skip {
		return nil
	}

	srcAdapter := factory(&env, srcSel)
	receiver, err := resolveLaneSanityReceiver(lggr, provider, env, destAdapter, destSel, receiverAddress)
	if err != nil {
		lggr.Warnf("%s: receiver %s→%s: %v",
			laneSanityCheckName, chainLabel(srcSel), chainLabel(destSel), err)
		addLaneFailureSummary(failed, srcSel, destSel, "n/a")
		return fmt.Errorf("failed token-transfer probes: %s", buildLaneSanityFailureSummary(failed))
	}
	extraArgs, err := destAdapter.GetExtraArgs(receiver, family)
	if err != nil {
		lggr.Warnf("%s: extra args %s→%s: %v",
			laneSanityCheckName, chainLabel(srcSel), chainLabel(destSel), err)
		addLaneFailureSummary(failed, srcSel, destSel, "n/a")
		return fmt.Errorf("failed token-transfer probes: %s", buildLaneSanityFailureSummary(failed))
	}

	mockReceiver, err := provider.MockReceiverAddress(env, destSel)
	if err != nil {
		lggr.Warnf("%s: mock receiver dest=%s: %v (PTT skipped)",
			laneSanityCheckName, chainLabel(destSel), err)
	}

	available, err := provider.AvailableTransferTokens(env, srcSel, destSel)
	if err != nil {
		lggr.Warnf("%s: available transfer tokens %s→%s: %v",
			laneSanityCheckName, chainLabel(srcSel), chainLabel(destSel), err)
		addLaneFailureSummary(failed, srcSel, destSel, "native")
		return fmt.Errorf("failed token-transfer probes: %s", buildLaneSanityFailureSummary(failed))
	}

	selectedTokens, err := transferTokenSelector(available, srcSel, destSel)
	if err != nil {
		return fmt.Errorf("token selection %s→%s: %w", chainLabel(srcSel), chainLabel(destSel), err)
	}
	if len(selectedTokens) == 0 {
		lggr.Infof("%s: no transfer tokens selected %s→%s",
			laneSanityCheckName, chainLabel(srcSel), chainLabel(destSel))
		return nil
	}
	pttSent := false
	for _, tok := range selectedTokens {
		if err := runTokenTransferScenario(ctx, lggr, env, srcAdapter, srcSel, destSel,
			tok, receiver, extraArgs, nil, provider, collector); err != nil {
			addLaneFailureSummary(failed, srcSel, destSel, tok)
		}
		if pttSent || len(mockReceiver) == 0 {
			continue
		}
		mockExtraArgs, mockErr := destAdapter.GetExtraArgs(mockReceiver, family)
		if mockErr != nil {
			lggr.Warnf("%s: PTT token=%s: extra args for MockReceiver dest=%s: %v",
				laneSanityCheckName, tok, chainLabel(destSel), mockErr)
			addLaneFailureSummary(failed, srcSel, destSel, "PTT:"+tok)
			continue
		}
		pttData := []byte("lane-sanity-check-payload")
		if err := runTokenTransferScenario(ctx, lggr, env, srcAdapter, srcSel, destSel,
			tok, mockReceiver, mockExtraArgs, pttData, provider, collector); err != nil {
			addLaneFailureSummary(failed, srcSel, destSel, "PTT:"+tok)
		}
		pttSent = true
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
	tok string,
	receiver []byte,
	extraArgs []byte,
	data []byte,
	provider PostProposalLaneSanity,
	collector *laneSanityResultCollector,
) error {
	amount, err := provider.FundAndApproveTransferToken(ctx, env, srcSel, tok)
	if err != nil {
		lggr.Warnf("%s: ❌ fund/approve token=%s src=%s: %v",
			laneSanityCheckName, tok, chainLabel(srcSel), err)
		return err
	}

	msg, err := srcAdapter.BuildMessage(testadapters.MessageComponents{
		DestChainSelector: destSel,
		Receiver:          receiver,
		Data:              data,
		ExtraArgs:         extraArgs,
		TokenAmounts:      []testadapters.TokenAmount{{Token: tok, Amount: amount}},
	})
	if err != nil {
		lggr.Warnf("%s: ❌ build message token=%s %s→%s: %v",
			laneSanityCheckName, tok, chainLabel(srcSel), chainLabel(destSel), err)
		return err
	}

	fee, feeErr := provider.GetMessageFee(ctx, env, srcSel, destSel, msg)
	if feeErr != nil {
		lggr.Warnf("%s: fee query token=%s %s→%s: %v (continuing)",
			laneSanityCheckName, tok, chainLabel(srcSel), chainLabel(destSel), feeErr)
	}

	_, msgID, err := srcAdapter.SendMessage(ctx, destSel, msg)
	if err != nil {
		lggr.Warnf("%s: ❌ send failed token=%s %s→%s: %v",
			laneSanityCheckName, tok, chainLabel(srcSel), chainLabel(destSel), err)
		return err
	}

	lggr.Infof("%s: ✅ token transfer sent", laneSanityCheckName)
	lggr.Infof("  source:      %s", chainLabel(srcSel))
	lggr.Infof("  destination: %s", chainLabel(destSel))
	lggr.Infof("  token:       %s | amount: %s", tok, amount.String())
	if fee != "" {
		lggr.Infof("  fee charged: %s", fee)
	}
	lggr.Infof("  message ID:  %s", msgID)
	lggr.Infof("  explorer:    %s", ccipExplorerURL(msgID))
	collector.add(laneSanityCSVRecord{
		SourceChain:   chainName(srcSel),
		DestChain:     chainName(destSel),
		FeeToken:      formatFeeTokenLogLabel(""),
		TransferToken: tok,
		Data:          formatDataForCSV(data),
		Fee:           fee,
		ExplorerLink:  ccipExplorerURL(msgID),
	})
	return nil
}

// resolveLaneSanityReceiver returns the CCIP receiver bytes for a lane check.
// When receiverAddress is empty the dest adapter's CCIPReceiver is used.
func resolveLaneSanityReceiver(
	lggr logger.Logger,
	provider PostProposalLaneSanity,
	env cldf.Environment,
	destAdapter testadapters.TestAdapterForFamily,
	destSel uint64,
	receiverAddress string,
) ([]byte, error) {
	if receiverAddress == "" {
		receiver := destAdapter.CCIPReceiver()
		lggr.Infof("%s: using default CCIPReceiver on dest=%s", laneSanityCheckName, chainLabel(destSel))
		return receiver, nil
	}
	receiver, err := provider.EncodeReceiverAddress(env, destSel, receiverAddress)
	if err != nil {
		return nil, err
	}
	lggr.Infof("%s: using receiver %q on dest=%s", laneSanityCheckName, receiverAddress, chainLabel(destSel))
	return receiver, nil
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
				laneSanityCheckName, destFamily, adapterVer, chainLabel(destSel))
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
			laneSanityCheckName, destFamily, adapterVer, chainLabel(destSel))
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

// chainName returns the human-readable chain name for a selector.
func chainName(sel uint64) string {
	chain, ok := chain_selectors.ChainBySelector(sel)
	if !ok {
		return fmt.Sprintf("chain-%d", sel)
	}
	return chain.Name
}

// formatDataForCSV renders message data as plain text when printable, else hex.
func formatDataForCSV(data []byte) string {
	if len(data) == 0 {
		return ""
	}
	if utf8.Valid(data) && isPrintableASCII(string(data)) {
		return string(data)
	}
	return "0x" + hex.EncodeToString(data)
}

func isPrintableASCII(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] < 0x20 || s[i] > 0x7e {
			return false
		}
	}
	return true
}

func laneSanityCSVOutputPath() string {
	if path := strings.TrimSpace(os.Getenv(laneSanityCSVOutputEnv)); path != "" {
		return path
	}
	return laneSanityDefaultCSVName
}

func writeLaneSanityResultsCSV(lggr logger.Logger, records []laneSanityCSVRecord) error {
	if len(records) == 0 {
		return nil
	}

	path := laneSanityCSVOutputPath()
	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("create %s: %w", path, err)
	}
	defer f.Close()

	w := csv.NewWriter(f)
	if err := w.Write([]string{
		"source",
		"destination",
		"fee_token",
		"transfer_token",
		"data",
		"fee",
		"explorer_link",
	}); err != nil {
		return fmt.Errorf("write csv header: %w", err)
	}
	for _, r := range records {
		if err := w.Write([]string{
			r.SourceChain,
			r.DestChain,
			r.FeeToken,
			r.TransferToken,
			r.Data,
			r.Fee,
			r.ExplorerLink,
		}); err != nil {
			return fmt.Errorf("write csv row: %w", err)
		}
	}
	w.Flush()
	if err := w.Error(); err != nil {
		return fmt.Errorf("flush csv: %w", err)
	}

	lggr.Infof("%s: wrote %d result(s) to %s", laneSanityCheckName, len(records), path)
	return nil
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

func promptTransferTokenSelection(
	available map[string]string,
	srcSel, destSel uint64,
) ([]string, error) {
	if len(available) == 0 {
		fmt.Fprintf(os.Stdout, "\n%s: no transfer tokens configured for %s→%s\n",
			laneSanityCheckName, chainLabel(srcSel), chainLabel(destSel))
		return nil, nil
	}

	names := make([]string, 0, len(available))
	for name := range available {
		names = append(names, name)
	}
	sort.Strings(names)

	fmt.Fprintf(os.Stdout, "\n%s: available transfer tokens for %s→%s:\n",
		laneSanityCheckName, chainLabel(srcSel), chainLabel(destSel))
	for i, name := range names {
		fmt.Fprintf(os.Stdout, "  %d. %s (%s)\n", i+1, name, available[name])
	}

	reader := bufio.NewReader(os.Stdin)
	var selected []selectedTransferToken
	selectedSet := make(map[string]struct{})

	for {
		printSelectedTransferTokens(selected)
		fmt.Fprintf(os.Stdout, "Select token (number/name) or 'done' to run transfers: ")
		line, err := reader.ReadString('\n')
		if err != nil {
			if errors.Is(err, io.EOF) && len(selected) > 0 {
				break
			}
			return nil, fmt.Errorf("read selection: %w", err)
		}

		input := strings.TrimSpace(line)
		if input == "" {
			continue
		}
		if strings.EqualFold(input, "done") {
			break
		}

		for _, part := range strings.Split(input, ",") {
			part = strings.TrimSpace(part)
			if part == "" {
				continue
			}

			name, addr, ok := resolveTransferTokenChoice(part, names, available)
			if !ok {
				fmt.Fprintf(os.Stdout, "  unknown token %q — enter a number from the list or a token name\n", part)
				continue
			}
			if _, exists := selectedSet[addr]; exists {
				fmt.Fprintf(os.Stdout, "  %s already selected\n", name)
				continue
			}
			selectedSet[addr] = struct{}{}
			selected = append(selected, selectedTransferToken{name: name, addr: addr})
			fmt.Fprintf(os.Stdout, "  added %s (%s)\n", name, addr)
		}
	}

	selectedAddrs := make([]string, len(selected))
	for i, tok := range selected {
		selectedAddrs[i] = tok.addr
	}
	if len(selectedAddrs) > 0 {
		printSelectedTransferTokens(selected)
		fmt.Fprintf(os.Stdout, "Running transfers for %d token(s) on %s→%s\n",
			len(selectedAddrs), chainLabel(srcSel), chainLabel(destSel))
	}
	return selectedAddrs, nil
}

type selectedTransferToken struct {
	name string
	addr string
}

func printSelectedTransferTokens(selected []selectedTransferToken) {
	if len(selected) == 0 {
		fmt.Fprintf(os.Stdout, "Selected tokens: (none)\n")
		return
	}
	fmt.Fprintf(os.Stdout, "Selected tokens:\n")
	for i, tok := range selected {
		fmt.Fprintf(os.Stdout, "  %d. %s (%s)\n", i+1, tok.name, tok.addr)
	}
}

func resolveTransferTokenChoice(
	choice string,
	names []string,
	available map[string]string,
) (name, address string, ok bool) {
	if idx, err := strconv.Atoi(choice); err == nil {
		if idx < 1 || idx > len(names) {
			return "", "", false
		}
		name = names[idx-1]
		return name, available[name], true
	}

	for candidate, addr := range available {
		if strings.EqualFold(candidate, choice) {
			return candidate, addr, true
		}
	}
	return "", "", false
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
