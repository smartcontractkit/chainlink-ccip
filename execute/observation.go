package execute

import (
	"bytes"
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"slices"
	"sort"
	"time"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"

	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/types"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-ccip/execute/exectypes"
	dt "github.com/smartcontractkit/chainlink-ccip/internal/plugincommon/discovery/discoverytypes"
	"github.com/smartcontractkit/chainlink-ccip/pkg/consts"
	"github.com/smartcontractkit/chainlink-ccip/pkg/logutil"
)

var (
	errOffRampConfigMismatch = errors.New("offramp config digest mismatch")
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
//
//nolint:gocyclo
func (p *Plugin) Observation(
	ctx context.Context, outctx ocr3types.OutcomeContext, _ types.Query,
) (types.Observation, error) {
	// Ensure that sequence number is in the context for consumption by all
	// downstream processors and the ccip reader.
	ctx, lggr := logutil.WithOCRInfo(ctx, p.lggr, outctx.SeqNr, logutil.PhaseObservation)

	var err error
	var previousOutcome exectypes.Outcome

	previousOutcome, err = p.ocrTypeCodec.DecodeOutcome(outctx.PreviousOutcome)
	if err != nil {
		return types.Observation{}, fmt.Errorf("unable to decode previous outcome: %w", err)
	}
	lggr.Infow("decoded previous outcome", "previousOutcome", previousOutcome)

	// If the previous outcome was the filter state, and reports were built, mark the messages as inflight.
	if previousOutcome.State == exectypes.Filter {
		// the lane is invalid due to a config digest mismatch, skip updating
		// the inflight cache so that messages will retry immediately when the digest is valid again.
		if err := p.checkConfigDigest(); err != nil {
			p.lggr.Errorw("skipping marking messages inflight due to config digest mismatch", "err", err)
			return types.Observation{}, fmt.Errorf("skipping observation: %w", err)
		}
		for _, execReport := range previousOutcome.Reports {
			for _, chainReport := range execReport.ChainReports {
				for _, message := range chainReport.Messages {
					p.inflightMessageCache.MarkInflight(chainReport.SourceChainSelector, message.Header.MessageID)
				}
			}
		}
	}

	fChain, err := p.homeChain.GetFChain()
	if err != nil {
		return types.Observation{}, fmt.Errorf("unable to get FChain: %w", err)
	}

	var discoveryObs dt.Observation
	// discovery processor disabled by setting it to nil.
	if p.discovery != nil {
		tStart := time.Now()
		discoveryObs, err = p.discovery.Observation(ctx, dt.Outcome{}, dt.Query{})
		if err != nil {
			lggr.Errorw("failed to discover contracts", "err", err)
		}
		lggr.Debugw("finished exec discovery observation",
			"discoveryObs", discoveryObs, "duration", time.Since(tStart))

		if !p.contractsInitialized {
			lggr.Infow("contracts not initialized, only making discovery observations",
				"discoveryObs", discoveryObs)
			return p.ocrTypeCodec.EncodeObservation(exectypes.Observation{Contracts: discoveryObs, FChain: fChain})
		}
	}

	observation := exectypes.Observation{
		Contracts: discoveryObs,
		FChain:    fChain,
	}

	tStart := time.Now()
	state := previousOutcome.State.Next()
	lggr.Debugw("Execute plugin performing observation", "state", state)
	switch state {
	case exectypes.GetCommitReports:
		observation, err = p.getCommitReportsObservation(ctx, lggr, observation)
		if err != nil {
			return nil, fmt.Errorf("getCommitReportsObservation: %w", err)
		}
	case exectypes.GetMessages:
		// Phase 2: Gather messages from the source chains and build the execution report.
		observation, err = p.getMessagesObservation(ctx, lggr, previousOutcome, observation)
		if err != nil {
			lggr.Errorw("failed to getMessagesObservation", "err", err)
			return nil, nil
		}
	case exectypes.Filter:
		// Phase 3: observe nonce for each unique source/sender pair.
		observation, err = p.getFilterObservation(ctx, lggr, previousOutcome, observation)
		if err != nil {
			lggr.Errorw("failed to getFilterObservation", "err", err)
			return nil, nil
		}
	default:
		return nil, fmt.Errorf("get observation: unknown state")
	}
	p.observer.TrackObservation(
		observation, state, outctx.Round) //nolint:staticcheck // we rely on Round for OTI metrics compatibility
	numCommitReports := 0
	for _, reports := range observation.CommitReports {
		numCommitReports += len(reports)
	}
	lggr.Infow("execute plugin got observation",
		"observationWithoutMsgDataAndDiscoveryObs", observation.ToLogFormat(),
		"duration", time.Since(tStart),
		"state", state,
		"numCommitReports", numCommitReports,
		"numMessages", observation.Messages.Count())

	return p.ocrTypeCodec.EncodeObservation(observation)
}

func (p *Plugin) getCurseInfo(ctx context.Context, lggr logger.Logger) (cciptypes.CurseInfo, error) {
	curseInfo, err := p.ccipReader.GetRmnCurseInfo(ctx)
	if err != nil {
		lggr.Errorw("get rmn curse info: rmn read error", "err", err)
		return cciptypes.CurseInfo{}, fmt.Errorf("get rmn curse info: rmn read error: %w", err)
	}

	return curseInfo, nil
}

// getCommitReportsObservations implements phase1 of the execute plugin state machine. It fetches commit reports from
// the destination chain and determines which messages are ready to be executed. These are added to the provided
// observation object.
//
// This function leverages an optimized caching strategy for commit reports:
// 1. Uses CommitReportCache to determine the optimal timestamp to query from
// 2. Efficiently refreshes the cache with reports containing Merkle roots
// 3. Automatically filters out reports without Merkle roots
// 4. Applies deduplication to prevent redundant processing
//
// Execution and Snoozing Logic:
// 1. For finalized executions:
//   - When messages are executed with finality (on destChain), they are permanently marked as executed
//   - MarkAsExecuted removes the commit root from the cache entirely to prevent reprocessing
//   - These messages will never be considered for execution again
//
// 2. For unfinalized executions:
//   - When messages are executed but not yet finalized, they are temporarily snoozed
//   - Snoozing adds the commit root to a snooze cache with a TTL
//   - If a reorg occurs and invalidates the execution, the messages become available again after the snooze period
//   - This prevents duplicate executions while allowing recovery from reorgs
//
// 3. For cursed chains:
//   - All commit reports from cursed source chains are snoozed
//   - This temporarily prevents execution until the curse is lifted
//   - Messages become available again after the snooze period if the chain is no longer cursed
func (p *Plugin) getCommitReportsObservation(
	ctx context.Context,
	lggr logger.Logger,
	observation exectypes.Observation,
) (exectypes.Observation, error) {
	// Phase 1: Gather commit reports from the destination chain and determine which messages are required to build
	//          a valid execution report.
	supportsDest, err := p.supportsDestChain()
	if err != nil {
		return exectypes.Observation{}, fmt.Errorf("unable to determine if the destination chain is supported: %w", err)
	}

	// Get the timestamp to fetch from.
	fetchFrom := time.Now().Add(-p.offchainCfg.MessageVisibilityInterval.Duration()).UTC()

	// No observation for non-dest readers.
	if !supportsDest {
		lggr.Infow("destination chain not supported, skipping commit reports observation")
		return observation, nil
	}

	// Refresh the commit report cache first
	// TODO: make this refresh async
	if err := p.commitReportCache.RefreshCache(ctx); err != nil {
		// Log error but proceed. If RefreshCache fails, GetReportsToQueryFromTimestamp
		// will likely use a less optimal (wider) window based on its internal state or defaults,
		// which is safer than halting observation entirely.
		lggr.Errorw("failed to refresh commit report cache, proceeding with potentially wider query window", "err", err)
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
	groupedCommits, fullyExecutedFinalized, fullyExecutedUnfinalized,
		err := getPendingReportsForExecution(
		ctx,
		p.ccipReader,
		p.commitReportCache,
		p.commitRootsCache.CanExecute,
		fetchFrom,
		ci.CursedSourceChains,
		int(p.offchainCfg.MaxCommitReportsToFetch),
		lggr,
	)
	if err != nil {
		return exectypes.Observation{}, err
	}

	// TODO: message from fullyExecutedCommits which are in the inflight messages cache could be cleared here.

	// If fully executed reports are detected, mark them in the cache.
	// This cache will be re-initialized on each plugin restart.
	for _, fullyExecutedCommit := range fullyExecutedFinalized {
		p.commitRootsCache.MarkAsExecuted(fullyExecutedCommit.SourceChain, fullyExecutedCommit.MerkleRoot)
	}

	// If fully executed reports are detected, snooze them in the cache.
	// This cache will be re-initialized on each plugin restart.
	for _, fullyExecutedCommit := range fullyExecutedUnfinalized {
		p.commitRootsCache.Snooze(fullyExecutedCommit.SourceChain, fullyExecutedCommit.MerkleRoot)
	}

	observation.CommitReports = groupedCommits

	// TODO: truncate grouped to a maximum observation size?
	return observation, nil
}

// observeTokenDataForMessage observes token data for a given message.
// It uses the token data observer to fetch token data
// and handles errors by initializing the token data with an error.
// It's okay to fetch 1 tokenData at a time without affecting peformance
// Reason for that is we either use background observer that caches data and returns instantly, or even in case
// we're using the sync observers, their implementations calls attestations also one by one.
func (p *Plugin) observeTokenDataForMessage(
	ctx context.Context,
	lggr logger.Logger,
	msg cciptypes.Message,
) exectypes.MessageTokenData {
	srcChain := msg.Header.SourceChainSelector
	seqNum := msg.Header.SequenceNumber
	msgObs := exectypes.MessageObservations{
		srcChain: {
			seqNum: msg,
		},
	}
	msgTkData, err := p.tokenDataObserver.Observe(ctx, msgObs)
	if err != nil {
		// In case of failure, initialize the token data with an error, that will prevent this specific token
		lggr.Errorw("failed to observe token data", "err", err, "messageID", msg.Header.MessageID)
		// from being processed and sent later but won't stop the rest of messages
		msgTkData = make(exectypes.TokenDataObservations)
		msgTkData[srcChain] = make(map[cciptypes.SeqNum]exectypes.MessageTokenData)
		msgTkData[srcChain][seqNum] = exectypes.NewMessageTokenData(exectypes.NewErrorTokenData(err))
	}
	return msgTkData[srcChain][seqNum]
}

// readMessagesForReport reads messages for the given report and validates they conform to the sequence range
func (p *Plugin) readMessagesForReport(
	ctx context.Context,
	lggr logger.Logger,
	srcChain cciptypes.ChainSelector,
	report exectypes.CommitData,
) ([]cciptypes.Message, error) {
	msgs, err := p.ccipReader.MsgsBetweenSeqNums(ctx, srcChain, report.SequenceNumberRange)
	if err != nil {
		return nil, err
	}

	if !msgsConformToSeqRange(msgs, report.SequenceNumberRange) {
		lggr.Errorw("missing messages in range",
			"srcChain", srcChain,
			"seqRange", report.SequenceNumberRange,
		)
		return nil, fmt.Errorf("missing messages in range")
	}

	return msgs, nil
}

// createEmptyMessageWithIDAndSeqNum creates a message with just the sequence number set
func createEmptyMessageWithIDAndSeqNum(msg cciptypes.Message) cciptypes.Message {
	return cciptypes.Message{
		Header: cciptypes.RampMessageHeader{
			MessageID:      msg.Header.MessageID,
			SequenceNumber: msg.Header.SequenceNumber,
		},
	}
}

// getMessagesObservation collects message observations from the provided commit data.
// It processes messages for each commit report, validating them against the sequence number range,
// and builds observation structures including message content and hashes.
//
// The function:
// 1. Sorts commit data for deterministic processing
// 2. Reads messages for each commit report
// 3. Tracks messages by source chain and sequence number
// 4. Computes hashes for each message for merkle tree verification
// 5. Observes token data for each message
// 6. Stops processing if the observation becomes too large to encode - This is a lenient check.
//
// Returns:
//   - Updated observation containing commit reports, messages, and hashes
//
//nolint:gocyclo // TODO: pull out appropriate helpers.
func (p *Plugin) getMessagesObservation(
	ctx context.Context,
	lggr logger.Logger,
	previousOutcome exectypes.Outcome,
	observation exectypes.Observation,
) (exectypes.Observation, error) {
	// Phase 2: Get messages.
	//          These messages will not be executed in the current round, but may be executed in future rounds
	//          (e.g. if gas prices decrease).
	if len(previousOutcome.CommitReports) == 0 {
		lggr.Info("No reports to execute. This is expected after a cold start.")
		// No reports to execute.
		// This is expected after a cold start.
		return observation, nil
	}

	commitData := previousOutcome.CommitReports

	messageObs := make(exectypes.MessageObservations)
	availableReports := make(exectypes.CommitObservations)
	msgHashes := make(exectypes.MessageHashes)
	tkData := make(exectypes.TokenDataObservations)

	sort.Slice(commitData, func(i, j int) bool {
		return exectypes.LessThan(commitData[i], commitData[j])
	})

	supportedChains, err := p.supportedChains(p.reportingCfg.OracleID)
	if err != nil {
		return exectypes.Observation{}, fmt.Errorf("get supported chains: %w", err)
	}

	stop := false

	totalMsgs := 0
	encodedObsSize := 0
	for _, report := range commitData {
		srcChain := report.SourceChain

		if !supportedChains.Contains(srcChain) {
			lggr.Infow("skipping report of unsupported source chain", "srcChain", srcChain)
			continue
		}

		// Read messages for this report's sequence number range
		msgs, err := p.readMessagesForReport(ctx, lggr, srcChain, report)
		if err != nil {
			lggr.Errorw("unable to read all messages for report",
				"srcChain", srcChain,
				"seqRange", report.SequenceNumberRange,
				"merkleRoot", report.MerkleRoot,
				"err", err,
			)
			continue
		}

		// Add the report to available reports
		availableReports[srcChain] = append(availableReports[srcChain], report)
		// Initialize data structures for this source chain if needed
		if _, ok := messageObs[srcChain]; !ok {
			messageObs[srcChain] = make(map[cciptypes.SeqNum]cciptypes.Message)
			msgHashes[srcChain] = make(map[cciptypes.SeqNum]cciptypes.Bytes32)
			tkData[srcChain] = make(map[cciptypes.SeqNum]exectypes.MessageTokenData)
		}

		// Initialize msg hashes with empty messages and empty token data
		// for all messages in the report
		// We need to have it like that to be able to build the merkle tree and proofs in later stages
		//nolint:lll // link
		// https://github.com/smartcontractkit/chainlink-ccip/blob/5d360b7c90126631c51f3e84f26c0417dc3877b6/execute/report/roots.go#L14-L14
		for _, msg := range msgs {
			seqNum := msg.Header.SequenceNumber
			messageObs[srcChain][seqNum] = createEmptyMessageWithIDAndSeqNum(msg)
			tkData[srcChain][seqNum] = exectypes.NewMessageTokenData()
			hash, err := p.msgHasher.Hash(ctx, msg)
			if err != nil {
				return exectypes.Observation{}, fmt.Errorf(
					"unable to hash message srcChain %d, seqNum %d, msgId %v: %w",
					srcChain, seqNum, msg.Header.MessageID, err,
				)
			}
			msgHashes[srcChain][seqNum] = hash
		}

		observation.CommitReports = availableReports
		observation.Hashes = msgHashes
		observation.Messages = messageObs
		observation.TokenData = tkData

		// Process each message in the report and override the empty message and token data if everything fits within
		// the size limits
		for _, msg := range msgs {
			// If a message is inflight or already executed, don't include it fully in the observation
			// because its already been transmitted in a previous report or executed onchain.
			if p.inflightMessageCache.IsInflight(srcChain, msg.Header.MessageID) {
				continue
			}
			if slices.Contains(report.ExecutedMessages, msg.Header.SequenceNumber) {
				continue
			}

			seqNum := msg.Header.SequenceNumber
			messageObs[srcChain][seqNum] = msg
			tkData[srcChain][seqNum] = p.observeTokenDataForMessage(ctx, lggr, msg)
			observation.Messages = messageObs
			observation.TokenData = tkData
			totalMsgs++

			encodedObs, err := p.ocrTypeCodec.EncodeObservation(observation)
			if err != nil {
				return exectypes.Observation{}, fmt.Errorf("unable to encode observation: %w", err)
			}
			encodedObsSize = len(encodedObs)
			if totalMsgs > lenientMaxMsgsPerObs || encodedObsSize > lenientMaxObservationLength {
				// remove the last message and token data
				observation.Messages[srcChain][seqNum] = createEmptyMessageWithIDAndSeqNum(msg)
				observation.TokenData[srcChain][seqNum] = exectypes.NewMessageTokenData()
				stop = true
				break
			}
		}

		// TODO: check if the commit report is all fully executed and inflight messages.
		// Might be able to just not include it in the observation at all? Or is there a
		// risk with that?

		if stop {
			lggr.Infow("Stop processing messages, observation is too large",
				"totalMsgs", totalMsgs, "encodedObsSize", encodedObsSize)
			break
		}
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
		lggr.Infow("destination chain not supported, skipping filter observation")
		return observation, nil
	}

	commitReportSenders := make(map[cciptypes.ChainSelector][]string)
	uniqueSenders := make(map[cciptypes.ChainSelector]map[string]struct{})
	for _, report := range previousOutcome.CommitReports {
		srcChain := report.SourceChain
		if _, ok := commitReportSenders[srcChain]; !ok {
			commitReportSenders[srcChain] = make([]string, 0)
		}

		if _, ok := uniqueSenders[srcChain]; !ok {
			uniqueSenders[srcChain] = make(map[string]struct{})
		}

		for _, msg := range report.Messages {
			// The GetMessages observation could potentially include a pseudo-deleted message,
			// which is either:
			// - a message that is inflight
			// - a message that has already been executed
			// - a message that was not included in the observation due to size limits
			// In all cases, since we don't need to execute these messages, we don't need to fetch their sender's nonce.
			if msg.IsPseudoDeleted() {
				lggr.Debugw("skipping pseudo-deleted message",
					"messageID", msg.Header.MessageID,
				)
				continue
			}

			sender, err := p.addrCodec.AddressBytesToString(msg.Sender[:], srcChain)
			if err != nil {
				lggr.Errorw("unable to convert sender address to string",
					"err", err, "senderAddress", msg.Sender)
				continue
			}

			if _, exists := uniqueSenders[srcChain][sender]; !exists {
				commitReportSenders[report.SourceChain] = append(commitReportSenders[srcChain], sender)
				uniqueSenders[srcChain][sender] = struct{}{}
			}
		}
	}

	// Get nonces of the addresses. If the call fails, we just return other observations
	nonceObservations, err := p.ccipReader.Nonces(ctx, commitReportSenders)
	if err != nil {
		lggr.Errorw("unable to get nonces", "err", err)
	} else {
		// Note: it is technically possible to check for curses at this point. If a curse
		// occurred after the GetMessages observations checking now could possibly recover a report.
		// It would only work for ordered messages.

		observation.Nonces = nonceObservations
	}
	return observation, nil
}

func (p *Plugin) checkConfigDigest() error {
	offRampConfigDigest, err := p.ccipReader.GetOffRampConfigDigest(context.Background(), consts.PluginTypeExecute)
	if err != nil {
		return fmt.Errorf("get offramp config digest: %w", err)
	}

	if !bytes.Equal(offRampConfigDigest[:], p.reportingCfg.ConfigDigest[:]) {
		p.lggr.Warnw("home chain config digest doesn't match offramp's config digest, not starting",
			"homeChainConfigDigest", p.reportingCfg.ConfigDigest,
			"offRampConfigDigest", hex.EncodeToString(offRampConfigDigest[:]),
		)
		return errOffRampConfigMismatch
	}
	return nil
}
