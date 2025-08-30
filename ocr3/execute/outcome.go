package execute

import (
	"context"
	"fmt"
	"sort"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"

	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/types"

	exectypes2 "github.com/smartcontractkit/chainlink-ccip/ocr3/execute/exectypes"
	"github.com/smartcontractkit/chainlink-ccip/ocr3/execute/internal"
	report2 "github.com/smartcontractkit/chainlink-ccip/ocr3/execute/report"
	"github.com/smartcontractkit/chainlink-ccip/ocr3/internal/libs/slicelib"
	"github.com/smartcontractkit/chainlink-ccip/ocr3/internal/plugincommon"
	dt "github.com/smartcontractkit/chainlink-ccip/ocr3/internal/plugincommon/discovery/discoverytypes"
	"github.com/smartcontractkit/chainlink-ccip/ocr3/pkg/logutil"
)

// Outcome collects the reports from the two phases and constructs the final outcome. Part of the outcome is a fully
// formed report that will be encoded for final transmission in the reporting phase.
func (p *Plugin) Outcome(
	ctx context.Context, outctx ocr3types.OutcomeContext, query types.Query, aos []types.AttributedObservation,
) (ocr3types.Outcome, error) {
	// Ensure that sequence number is in the context for consumption by all
	// downstream processors and the ccip reader.
	ctx, lggr := logutil.WithOCRInfo(ctx, p.lggr, outctx.SeqNr, logutil.PhaseOutcome)

	previousOutcome, err := p.ocrTypeCodec.DecodeOutcome(outctx.PreviousOutcome)
	if err != nil {
		return nil, fmt.Errorf("unable to decode previous outcome: %w", err)
	}

	state := previousOutcome.State.Next()
	lggr = logger.With(lggr, "execPluginState", state)
	lggr.Debugw("Execute plugin performing outcome",
		"outctx", outctx,
		"query", query,
		"attributedObservations", aos,
	)

	decodedAos, err := decodeAttributedObservations(aos, p.ocrTypeCodec)
	if err != nil {
		return nil, fmt.Errorf("unable to decode observations: %w", err)
	}

	// discovery processor disabled by setting it to nil.
	if p.discovery != nil {
		mapper := func(
			ao plugincommon.AttributedObservation[exectypes2.Observation],
		) plugincommon.AttributedObservation[dt.Observation] {
			return plugincommon.AttributedObservation[dt.Observation]{
				OracleID:    ao.OracleID,
				Observation: ao.Observation.Contracts,
			}
		}
		discoveryAos := slicelib.Map(decodedAos, mapper)
		_, err = p.discovery.Outcome(ctx, dt.Outcome{}, dt.Query{}, discoveryAos)
		if err != nil {
			lggr.Errorw("discovery processor outcome errored", "err", err)
		} else {
			p.contractsInitialized = true
		}
	}

	observation, err := computeConsensusObservation(lggr, decodedAos, p.destChain, p.reportingCfg.F)
	if err != nil {
		return ocr3types.Outcome{}, fmt.Errorf("unable to get consensus observation: %w", err)
	}

	var outcome exectypes2.Outcome
	switch state {
	case exectypes2.GetCommitReports:
		outcome = p.getCommitReportsOutcome(observation)
	case exectypes2.GetMessages:
		outcome = p.getMessagesOutcome(lggr, observation)
	case exectypes2.Filter:
		outcome, err = p.getFilterOutcome(ctx, lggr, observation, previousOutcome)
		if err != nil {
			// We want to have an empty previousOutcome in the next round. To achieve this we don't return an error.
			lggr.Errorw("get filter outcome", "err", err)
			return nil, nil
		}
	default:
		panic("unknown state")
	}

	// This may happen if there is nothing to observe, or during startup when the contracts have
	// been discovered. In the latter case, getCommitReportsOutcome will return an empty outcome.
	if outcome.IsEmpty() {
		lggr.Warnw("exec outcome: empty outcome")
		if p.contractsInitialized {
			return p.ocrTypeCodec.EncodeOutcome(exectypes2.Outcome{State: exectypes2.Initialized})
		}
		return nil, nil
	}

	p.observer.TrackOutcome(outcome, state)
	lggr.Infow("generated outcome",
		"outcomeWithoutMsgData", outcome.ToLogFormat(),
		"numCommitReports", len(outcome.CommitReports),
		"numExecReports", len(outcome.Reports),
		"numMessages", observation.Messages.Count(),
	)

	return p.ocrTypeCodec.EncodeOutcome(outcome)
}

func (p *Plugin) getCommitReportsOutcome(observation exectypes2.Observation) exectypes2.Outcome {
	// flatten commit reports and sort by timestamp.
	var commitReports []exectypes2.CommitData
	for _, report := range observation.CommitReports {
		commitReports = append(commitReports, report...)
	}
	sort.Slice(commitReports, func(i, j int) bool {
		return commitReports[i].Timestamp.Before(commitReports[j].Timestamp)
	})

	// Must use 'NewOutcome' rather than direct struct initialization to ensure the outcome is sorted.
	// TODO: sort in the encoder.
	return exectypes2.NewOutcomeWithSortedCommitReports(exectypes2.GetCommitReports, commitReports)
}

func (p *Plugin) getMessagesOutcome(
	lggr logger.Logger,
	observation exectypes2.Observation,
) exectypes2.Outcome {
	commitReports := make([]exectypes2.CommitData, 0)

	// First ensure that all observed messages has hashes and token data.
	if err := validateHashesExist(observation.Messages, observation.Hashes); err != nil {
		lggr.Errorw("validate hashes exist", "err", err)
		return exectypes2.Outcome{}
	}
	if err := validateTokenDataObservations(observation.Messages, observation.TokenData); err != nil {
		lggr.Errorw("validate token data observations: %w", err)
		return exectypes2.Outcome{}
	}

	reports := observation.CommitReports.Flatten()
	// add messages to their commitReports.
	for i, report := range reports {
		report.Messages = nil
		for j := report.SequenceNumberRange.Start(); j <= report.SequenceNumberRange.End(); j++ {
			if msg, ok := observation.Messages[report.SourceChain][j]; ok {
				// Always add the message and hash, even if it wont be executed.
				// This slice must have an entry for each message in the commit range.
				report.Messages = append(report.Messages, msg)

				report.Hashes = append(report.Hashes, observation.Hashes[report.SourceChain][j])
				report.MessageTokenData = append(report.MessageTokenData, observation.TokenData[report.SourceChain][j])
			}
		}
		if len(report.Messages) == 0 {
			// If there are no messages, remove the commit report.
			commitReports = internal.RemoveIthElement(commitReports, i)
		}
		commitReports = append(commitReports, report)
	}

	// Must use 'NewOutcome' rather than direct struct initialization to ensure the outcome is sorted.
	// TODO: sort in the encoder.
	return exectypes2.NewOutcomeWithSortedCommitReports(exectypes2.GetMessages, commitReports)
}

// getFilterOutcome is the final phase of the execution plugin. Filter refers to the Nonces
// being passed along with message data to perform a final filtering of messages.
func (p *Plugin) getFilterOutcome(
	ctx context.Context,
	lggr logger.Logger,
	observation exectypes2.Observation,
	previousOutcome exectypes2.Outcome,
) (exectypes2.Outcome, error) {
	commitReports := previousOutcome.CommitReports

	builder := report2.NewBuilder(
		lggr,
		p.msgHasher,
		p.reportCodec,
		p.estimateProvider,
		p.destChain,
		p.addrCodec,
		report2.WithMultipleReports(p.offchainCfg.MultipleReportsEnabled),
		report2.WithMaxReportsCount(maxReportCount),
		report2.WithMaxReportSizeBytes(maxReportLength),
		report2.WithMaxGas(p.offchainCfg.BatchGasLimit),
		report2.WithExtraMessageCheck(report2.CheckNonces(observation.Nonces, p.addrCodec)),
		//TODO: remove as we already check it in GetMessages phase
		report2.WithExtraMessageCheck(report2.CheckIfInflight(p.inflightMessageCache.IsInflight)),
		report2.WithMaxMessages(p.offchainCfg.MaxReportMessages),
		report2.WithMaxSingleChainReports(p.offchainCfg.MaxSingleChainReports),
	)

	execReports, selectedCommitReports, err := selectReports(
		ctx,
		lggr,
		commitReports,
		builder)
	if err != nil {
		return exectypes2.Outcome{}, fmt.Errorf("unable to select report: %w", err)
	}

	// Collapse the commit reports.
	var mergedData []exectypes2.CommitData
	for _, commitData := range selectedCommitReports {
		mergedData = append(mergedData, commitData...)
	}

	// Must use 'NewOutcome' rather than direct struct initialization to ensure the outcome is sorted.
	return exectypes2.NewOutcome(exectypes2.Filter, mergedData, execReports), nil
}
