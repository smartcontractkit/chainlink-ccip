package execute

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/exp/maps"

	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/types"

	"github.com/smartcontractkit/chainlink-ccip/execute/exectypes"
	typeconv "github.com/smartcontractkit/chainlink-ccip/internal/libs/typeconv"
	dt "github.com/smartcontractkit/chainlink-ccip/internal/plugincommon/discovery/discoverytypes"
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
	var err error
	var previousOutcome exectypes.Outcome

	previousOutcome, err = exectypes.DecodeOutcome(outctx.PreviousOutcome)
	if err != nil {
		return types.Observation{}, fmt.Errorf("unable to decode previous outcome: %w", err)
	}
	p.lggr.Infow("decoded previous outcome", "previousOutcome", previousOutcome)

	var discoveryObs dt.Observation
	// discovery processor disabled by setting it to nil.
	if p.discovery != nil {
		discoveryObs, err = p.discovery.Observation(ctx, dt.Outcome{}, dt.Query{})
		if err != nil {
			p.lggr.Errorw("failed to discover contracts", "err", err)
		}

		if !p.contractsInitialized {
			p.lggr.Infow("contracts not initialized, only making discovery observations",
				"discoveryObs", discoveryObs)
			return exectypes.Observation{Contracts: discoveryObs}.Encode()
		}
	}

	observation := exectypes.Observation{
		Contracts: discoveryObs,
	}

	state := previousOutcome.State.Next()
	p.lggr.Debugw("Execute plugin performing observation", "state", state)
	switch state {
	case exectypes.GetCommitReports:
		observation, err = p.getCommitReportsObservation(ctx, observation)
	case exectypes.GetMessages:
		// Phase 2: Gather messages from the source chains and build the execution report.
		observation, err = p.getMessagesObservation(ctx, previousOutcome, observation)
	case exectypes.Filter:
		// Phase 3: observe nonce for each unique source/sender pair.
		observation, err = p.getFilterObservation(ctx, previousOutcome, observation)
	default:
		err = fmt.Errorf("unknown state")
	}

	if err != nil {
		return nil, err
	}

	p.lggr.Infow("execute plugin got observation", "observation", observation)

	return observation.Encode()
}

// getCommitReportsObservations implements phase1 of the execute plugin state machine. It fetches commit reports from
// the destination chain and determines which messages are ready to be executed. These are added to the provided
// observation object.
func (p *Plugin) getCommitReportsObservation(
	ctx context.Context,
	observation exectypes.Observation,
) (exectypes.Observation, error) {
	fetchFrom := time.Now().Add(-p.offchainCfg.MessageVisibilityInterval.Duration()).UTC()

	// Phase 1: Gather commit reports from the destination chain and determine which messages are required to build
	//          a valid execution report.
	supportsDest, err := p.supportsDestChain()
	if err != nil {
		return exectypes.Observation{}, fmt.Errorf("unable to determine if the destination chain is supported: %w", err)
	}
	if supportsDest {
		groupedCommits, err := getPendingExecutedReports(ctx, p.ccipReader, p.destChain, fetchFrom, p.lggr)
		if err != nil {
			return exectypes.Observation{}, err
		}

		observation.CommitReports = groupedCommits

		// TODO: truncate grouped to a maximum observation size?
		return observation, nil
	}

	// No observation for non-dest readers.
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

func (p *Plugin) getMessagesObservation(
	ctx context.Context,
	previousOutcome exectypes.Outcome,
	observation exectypes.Observation,
) (exectypes.Observation, error) {
	messageObs := make(exectypes.MessageObservations)
	messages := make([]cciptypes.Message, 0)
	messageTimestamps := make(map[cciptypes.Bytes32]time.Time)
	if len(previousOutcome.PendingCommitReports) == 0 {
		p.lggr.Debug("TODO: No reports to execute. This is expected after a cold start.")
		// No reports to execute.
		// This is expected after a cold start.
	} else {
		commitReportCache := make(map[cciptypes.ChainSelector][]exectypes.CommitData)
		for _, report := range previousOutcome.PendingCommitReports {
			commitReportCache[report.SourceChain] = append(commitReportCache[report.SourceChain], report)
		}

		for srcChain, reports := range commitReportCache {
			if len(reports) == 0 {
				continue
			}

			ranges, err := computeRanges(reports)
			if err != nil {
				return exectypes.Observation{}, err
			}

			//executedSeqNrs := executedMessages(reports)

			// Read messages for each range.
			for _, seqRange := range ranges {
				// TODO: check if srcChain is supported.
				msgs, err := p.ccipReader.MsgsBetweenSeqNums(ctx, srcChain, seqRange)
				if err != nil {
					return exectypes.Observation{}, err
				}
				messages = append(messages, msgs...)
				for _, msg := range msgs {
					if _, ok := messageObs[srcChain]; !ok {
						messageObs[srcChain] = make(map[cciptypes.SeqNum]cciptypes.Message)
					}

					//if _, ok := executedSeqNrs[msg.Header.SequenceNumber]; ok {
					//	// already executed, don't add
					//	continue
					//}
					messageObs[srcChain][msg.Header.SequenceNumber] = msg
					// TODO: Check map encoding size and stop when it's more than max observation size.
				}
			}
		}

		var err error
		messageTimestamps, err = getMessageTimestampMap(commitReportCache, messageObs)
		if err != nil {
			return exectypes.Observation{}, err
		}
	}

	tkData, err1 := p.tokenDataObserver.Observe(ctx, messageObs)
	if err1 != nil {
		return exectypes.Observation{}, fmt.Errorf("unable to process token data %w", err1)
	}

	costlyMessages, err := p.costlyMessageObserver.Observe(ctx, messages, messageTimestamps)
	if err != nil {
		return exectypes.Observation{}, fmt.Errorf("unable to observe costly messageObs %w", err)
	}

	observation.CommitReports = regroup(previousOutcome.PendingCommitReports)
	observation.Messages = messageObs
	observation.CostlyMessages = costlyMessages
	observation.TokenData = tkData

	encodedObs, err := observation.Encode()
	if err != nil {
		return exectypes.Observation{}, fmt.Errorf("unable to encode observation: %w", err)
	}

	p.lggr.Infow("encoded observation",
		"size", len(encodedObs),
	)
	if len(encodedObs) >= maxObservationLength {
		return exectypes.Observation{}, fmt.Errorf(
			"observation size exceeds maximum size, current size: %d",
			len(encodedObs),
		)
	}

	return observation, nil
}

func (p *Plugin) getFilterObservation(
	ctx context.Context,
	previousOutcome exectypes.Outcome,
	observation exectypes.Observation,
) (exectypes.Observation, error) {
	// Phase 2: Filter messages and determine which messages are too costly to execute.
	//          This phase also determines which messages are too costly to execute.
	//          These messages will not be executed in the current round, but may be executed in future rounds
	//          (e.g. if gas prices decrease or if these messages' fees are boosted high enough).

	nonceRequestArgs := make(map[cciptypes.ChainSelector]map[string]struct{})

	// Collect unique senders.
	for _, commitReport := range previousOutcome.PendingCommitReports {
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
			return exectypes.Observation{}, fmt.Errorf("unable to get nonces: %w", err)
		}
		nonceObservations[srcChain] = nonces
	}

	observation.Nonces = nonceObservations

	return observation, nil
}
