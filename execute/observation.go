package execute

import (
	"context"
	"fmt"
	"time"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"

	"golang.org/x/exp/maps"

	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/types"

	"github.com/smartcontractkit/chainlink-ccip/execute/exectypes"
	typeconv "github.com/smartcontractkit/chainlink-ccip/internal/libs/typeconv"
	dt "github.com/smartcontractkit/chainlink-ccip/internal/plugincommon/discovery/discoverytypes"
	"github.com/smartcontractkit/chainlink-ccip/pkg/logutil"
	"github.com/smartcontractkit/chainlink-ccip/pkg/reader"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

// Observation collects data across two phases which happen in separate rounds.
// These phases happen continuously so that except for the first round, every
// subsequent round can have a new execution report.
//
// Phase 1: Gather commit reports from the destination chain and determine
// which messages are required to build a valid execution report.
//
// Phase 2: Gather messages from the source chains and build the execution
// report.
//
// Phase 3: observe nonce for each unique source/sender pair.
func (p *Plugin) Observation(
	ctx context.Context, outctx ocr3types.OutcomeContext, _ types.Query,
) (types.Observation, error) {
	// Ensure that sequence number is in the context for consumption by all
	// downstream processors and the ccip reader.
	ctx, lggr := logutil.WithOCRInfo(ctx, p.lggr, outctx.SeqNr, logutil.PhaseObservation)

	var err error
	var previousOutcome exectypes.Outcome

	previousOutcome, err = exectypes.DecodeOutcome(outctx.PreviousOutcome)
	if err != nil {
		return types.Observation{}, fmt.Errorf("unable to decode previous outcome: %w", err)
	}
	lggr.Infow("decoded previous outcome", "previousOutcome", previousOutcome)

	var discoveryObs dt.Observation
	// discovery processor disabled by setting it to nil.
	if p.discovery != nil {
		discoveryObs, err = p.discovery.Observation(ctx, dt.Outcome{}, dt.Query{})
		if err != nil {
			lggr.Errorw("failed to discover contracts", "err", err)
		}

		if !p.contractsInitialized {
			lggr.Infow("contracts not initialized, only making discovery observations",
				"discoveryObs", discoveryObs)
			return exectypes.Observation{Contracts: discoveryObs}.Encode()
		}
	}

	observation := exectypes.Observation{
		Contracts: discoveryObs,
	}

	state := previousOutcome.State.Next()
	lggr.Debugw("Execute plugin performing observation", "state", state)
	switch state {
	case exectypes.GetCommitReports:
		observation, err = p.getCommitReportsObservation(ctx, lggr, observation)
	case exectypes.GetMessages:
		// Phase 2: Gather messages from the source chains and build the execution report.
		observation, err = p.getMessagesObservation(ctx, lggr, previousOutcome, observation)
	case exectypes.Filter:
		// Phase 3: observe nonce for each unique source/sender pair.
		observation, err = p.getFilterObservation(ctx, lggr, previousOutcome, observation)
	default:
		err = fmt.Errorf("unknown state")
	}

	if err != nil {
		return nil, err
	}

	p.observer.TrackObservation(observation, state)
	lggr.Infow("execute plugin got observation", "observation", observation)

	return observation.Encode()
}

func (p *Plugin) getCurseInfo(ctx context.Context, lggr logger.Logger) (*reader.CurseInfo, error) {
	allSourceChains, err := p.chainSupport.KnownSourceChainsSlice()
	if err != nil {
		lggr.Warnw("call to KnownSourceChainsSlice failed", "err", err)
		return nil, fmt.Errorf("call to KnownSourceChainsSlice failed: %w", err)
	}

	curseInfo, err := p.ccipReader.GetRmnCurseInfo(ctx, p.chainSupport.DestChain(), allSourceChains)
	if err != nil {
		lggr.Errorw("nothing to observe: rmn read error",
			"err", err,
			"sourceChains", allSourceChains,
		)
		return nil, fmt.Errorf("nothing to observe: rmn read error: %w", err)
	}

	return curseInfo, nil
}

// getCommitReportsObservations implements phase1 of the execute plugin state machine. It fetches commit reports from
// the destination chain and determines which messages are ready to be executed. These are added to the provided
// observation object.
func (p *Plugin) getCommitReportsObservation(
	ctx context.Context,
	lggr logger.Logger,
	observation exectypes.Observation,
) (exectypes.Observation, error) {
	fetchFrom := time.Now().Add(-p.offchainCfg.MessageVisibilityInterval.Duration()).UTC()

	// Phase 1: Gather commit reports from the destination chain and determine which messages are required to build
	//          a valid execution report.
	supportsDest, err := p.supportsDestChain()
	if err != nil {
		return exectypes.Observation{}, fmt.Errorf("unable to determine if the destination chain is supported: %w", err)
	}

	// No observation for non-dest readers.
	if !supportsDest {
		return observation, nil
	}

	// Get curse information from the destination chain.
	ci, err := p.getCurseInfo(ctx, lggr)
	if err != nil {
		// If we can't get curse info, we can't proceed.
		// But we still need to return discovery data.
		// The error is logged by getCurseInfo.
		return observation, nil
	}
	if ci.GlobalCurse || ci.CursedDestination {
		lggr.Warnw("nothing to observe: rmn curse", "curseInfo", ci)
		return observation, nil
	}

	// Get pending exec reports.
	groupedCommits, err := getPendingExecutedReports(ctx, p.ccipReader, p.destChain, fetchFrom, lggr)
	if err != nil {
		return exectypes.Observation{}, err
	}

	// Remove cursed observations.
	for chainSelector, isCursed := range ci.CursedSourceChains {
		if isCursed {
			delete(groupedCommits, chainSelector)
		}
	}

	observation.CommitReports = groupedCommits

	// TODO: truncate grouped to a maximum observation size?
	return observation, nil
}

// regroup converts the previous outcome to the observation format.
// TODO: use same format for Observation and Outcome.
func regroup(commitData []exectypes.CommitData) exectypes.CommitObservations {
	groupedCommits := make(exectypes.CommitObservations)
	for _, report := range commitData {
		if _, ok := groupedCommits[report.SourceChain]; !ok {
			groupedCommits[report.SourceChain] = []exectypes.CommitData{}
		}
		groupedCommits[report.SourceChain] = append(groupedCommits[report.SourceChain], report)
	}
	return groupedCommits
}

func readAllMessages(
	ctx context.Context,
	lggr logger.Logger,
	ccipReader reader.CCIPReader,
	commitData []exectypes.CommitData,
) (exectypes.MessageObservations, exectypes.CommitObservations, map[cciptypes.Bytes32]time.Time) {
	messageObs := make(exectypes.MessageObservations)
	availableReports := make(exectypes.CommitObservations)
	messageTimestamps := make(map[cciptypes.Bytes32]time.Time)

	commitObs := regroup(commitData)

	for srcChain, reports := range commitObs {
		if len(reports) == 0 {
			continue
		}

		messageObs[srcChain] = make(map[cciptypes.SeqNum]cciptypes.Message)
		// Read messages for each range.
		for _, report := range reports {
			msgs, err := ccipReader.MsgsBetweenSeqNums(ctx, srcChain, report.SequenceNumberRange)
			if err != nil || len(msgs) != report.SequenceNumberRange.Length() ||
				gapInSequenceNumbers(msgs, report.SequenceNumberRange) {
				lggr.Errorw("unable to read all messages for report",
					"srcChain", srcChain,
					"seqRange", report.SequenceNumberRange,
					"merkleRoot", report.MerkleRoot,
					"err", err,
				)
				continue
			}
			for _, msg := range msgs {
				messageObs[srcChain][msg.Header.SequenceNumber] = msg
				messageTimestamps[msg.Header.MessageID] = report.Timestamp
			}
			availableReports[srcChain] = append(availableReports[srcChain], report)
		}
		// Remove empty chains.
		if len(messageObs[srcChain]) == 0 {
			delete(messageObs, srcChain)
		}
	}
	return messageObs, availableReports, messageTimestamps
}

// gapInSequenceNumbers returns true if there is a gap messages sequence numbers.
func gapInSequenceNumbers(msgs []cciptypes.Message, numberRange cciptypes.SeqNumRange) bool {
	// Add msgs to a map to check for duplicates.
	msgMap := make(map[cciptypes.SeqNum]struct{})
	for _, msg := range msgs {
		msgMap[msg.Header.SequenceNumber] = struct{}{}
	}
	// Check for gaps in sequence numbers.
	for i := numberRange.Start(); i <= numberRange.End(); i++ {
		if _, ok := msgMap[i]; !ok {
			return true
		}
	}
	return false
}

func (p *Plugin) getMessagesObservation(
	ctx context.Context,
	lggr logger.Logger,
	previousOutcome exectypes.Outcome,
	observation exectypes.Observation,
) (exectypes.Observation, error) {
	// Phase 2: Get messages and determine which messages are too costly to execute.
	//          These messages will not be executed in the current round, but may be executed in future rounds
	//          (e.g. if gas prices decrease or if these messages' fees are boosted high enough).
	if len(previousOutcome.CommitReports) == 0 {
		lggr.Debug("TODO: No reports to execute. This is expected after a cold start.")
		// No reports to execute.
		// This is expected after a cold start.
		return observation, nil
	}

	messageObs, commitReportCache, messageTimestamps := readAllMessages(
		ctx,
		lggr,
		p.ccipReader,
		previousOutcome.CommitReports,
	)

	tkData, err1 := p.tokenDataObserver.Observe(ctx, messageObs)
	if err1 != nil {
		return exectypes.Observation{}, fmt.Errorf("unable to process token data %w", err1)
	}

	// validating before continuing with heavy operations afterwards like truncation and costly messages
	// all messages should have a token data observation even if it's empty
	if validateTokenDataObservations(messageObs, tkData) != nil {
		return exectypes.Observation{}, fmt.Errorf("invalid token data observations")
	}

	costlyMessages, err := p.costlyMessageObserver.Observe(ctx, messageObs.Flatten(), messageTimestamps)
	if err != nil {
		return exectypes.Observation{}, fmt.Errorf("unable to observe costly messages: %w", err)
	}

	hashes, err := exectypes.GetHashes(ctx, messageObs, p.msgHasher)
	if err != nil {
		return exectypes.Observation{}, fmt.Errorf("unable to get message hashes: %w", err)
	}

	observation.CommitReports = commitReportCache
	observation.Messages = messageObs
	observation.Hashes = hashes
	observation.CostlyMessages = costlyMessages
	observation.TokenData = tkData

	// Make sure encoded observation fits within the maximum observation size.
	//observation, err = truncateObservation(observation, maxObservationLength, p.emptyEncodedSizes)
	observation, err = p.observationOptimizer.TruncateObservation(observation)
	if err != nil {
		return exectypes.Observation{}, fmt.Errorf("unable to truncate observation: %w", err)
	}

	return observation, nil
}

func (p *Plugin) getFilterObservation(
	ctx context.Context,
	lggr logger.Logger,
	previousOutcome exectypes.Outcome,
	observation exectypes.Observation,
) (exectypes.Observation, error) {
	supportsDest, err := p.supportsDestChain()
	if err != nil {
		return exectypes.Observation{}, fmt.Errorf("unable to determine if the destination chain is supported: %w", err)
	}
	// No observation for non-dest readers.
	if !supportsDest {
		return observation, nil
	}

	// Collect unique senders.
	nonceRequestArgs := make(map[cciptypes.ChainSelector]map[string]struct{})
	for _, commitReport := range previousOutcome.CommitReports {
		if _, ok := nonceRequestArgs[commitReport.SourceChain]; !ok {
			nonceRequestArgs[commitReport.SourceChain] = make(map[string]struct{})
		}

		for _, msg := range commitReport.Messages {
			sender := typeconv.AddressBytesToString(msg.Sender[:], uint64(p.destChain))
			nonceRequestArgs[commitReport.SourceChain][sender] = struct{}{}
		}
	}

	// Read args from chain.
	nonceObservations := make(exectypes.NonceObservations)
	for srcChain, addrSet := range nonceRequestArgs {
		// TODO: check if srcSelector is supported.
		addrs := maps.Keys(addrSet)
		nonces, err := p.ccipReader.Nonces(ctx, srcChain, p.destChain, addrs)
		if err != nil {
			lggr.Errorw("unable to get nonces", "err", err)
			continue
		}
		nonceObservations[srcChain] = nonces
	}

	// Note: it is technically possible to check for curses at this point. If a curse
	// occurred after the GetMessages observations checking now could possibly recover a report.
	// It would only work for ordered messages.

	observation.Nonces = nonceObservations

	return observation, nil
}
