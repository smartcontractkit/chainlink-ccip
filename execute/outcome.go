package execute

import (
	"context"
	"fmt"
	"sort"

	"github.com/smartcontractkit/chainlink-ccip/execute/internal"

	mapset "github.com/deckarep/golang-set/v2"

	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/types"

	"github.com/smartcontractkit/chainlink-ccip/execute/exectypes"
	"github.com/smartcontractkit/chainlink-ccip/execute/report"
	"github.com/smartcontractkit/chainlink-ccip/internal/libs/slicelib"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon"
	dt "github.com/smartcontractkit/chainlink-ccip/internal/plugincommon/discovery/discoverytypes"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

// Outcome collects the reports from the two phases and constructs the final outcome. Part of the outcome is a fully
// formed report that will be encoded for final transmission in the reporting phase.
func (p *Plugin) Outcome(
	ctx context.Context, outctx ocr3types.OutcomeContext, query types.Query, aos []types.AttributedObservation,
) (ocr3types.Outcome, error) {
	p.lggr.Debugw("Execute plugin performing outcome",
		"outctx", outctx,
		"query", query,
		"attributedObservations", aos)
	previousOutcome, err := exectypes.DecodeOutcome(outctx.PreviousOutcome)
	if err != nil {
		return nil, fmt.Errorf("unable to decode previous outcome: %w", err)
	}

	decodedAos := decodeAttributedObservations(p.lggr, aos)

	// discovery processor disabled by setting it to nil.
	if p.discovery != nil {
		mapper := func(
			ao plugincommon.AttributedObservation[exectypes.Observation],
		) plugincommon.AttributedObservation[dt.Observation] {
			return plugincommon.AttributedObservation[dt.Observation]{
				OracleID:    ao.OracleID,
				Observation: ao.Observation.Contracts,
			}
		}
		discoveryAos := slicelib.Map(decodedAos, mapper)
		_, err = p.discovery.Outcome(ctx, dt.Outcome{}, dt.Query{}, discoveryAos)
		if err != nil {
			return nil, fmt.Errorf("unable to process outcome of discovery processor: %w", err)
		}
		p.contractsInitialized = true
	}

	fChain, err := p.homeChain.GetFChain()
	if err != nil {
		return ocr3types.Outcome{}, fmt.Errorf("unable to get FChain: %w", err)
	}
	_, ok := fChain[p.destChain] // check if the destination chain is in the FChain.
	if !ok {
		return ocr3types.Outcome{}, fmt.Errorf("destination chain %d is not in FChain", p.destChain)
	}

	observation, err := getConsensusObservation(p.lggr, decodedAos, p.destChain, p.reportingCfg.F, fChain)
	if err != nil {
		return ocr3types.Outcome{}, fmt.Errorf("unable to get consensus observation: %w", err)
	}

	var outcome exectypes.Outcome
	state := previousOutcome.State.Next()
	switch state {
	case exectypes.GetCommitReports:
		outcome = p.getCommitReportsOutcome(observation)
	case exectypes.GetMessages:
		outcome = p.getMessagesOutcome(observation)
	case exectypes.Filter:
		outcome, err = p.getFilterOutcome(ctx, observation, previousOutcome)
	default:
		panic("unknown state")
	}

	if err != nil {
		p.lggr.Warnw(
			fmt.Sprintf("[oracle %d] exec outcome error", p.reportingCfg.OracleID),
			"err", err)
		return nil, fmt.Errorf("unable to get outcome: %w", err)
	}

	// This may happen if there is nothing to observe, or during startup when the contracts have
	// been discovered. In the latter case, getCommitReportsOutcome will return an empty outcome.
	if outcome.IsEmpty() {
		p.lggr.Warnw(
			fmt.Sprintf("[oracle %d] exec outcome: empty outcome", p.reportingCfg.OracleID),
			"execPluginState", state)
		if p.contractsInitialized {
			return exectypes.Outcome{State: exectypes.Initialized}.Encode()
		}
		return nil, nil
	}

	p.observer.TrackOutcome(outcome, state)
	p.lggr.Infow("generated outcome", "execPluginState", state, "outcome", outcome)

	return outcome.Encode()
}

func (p *Plugin) getCommitReportsOutcome(observation exectypes.Observation) exectypes.Outcome {
	// flatten commit reports and sort by timestamp.
	var commitReports []exectypes.CommitData
	for _, report := range observation.CommitReports {
		commitReports = append(commitReports, report...)
	}
	sort.Slice(commitReports, func(i, j int) bool {
		return commitReports[i].Timestamp.Before(commitReports[j].Timestamp)
	})

	// Must use 'NewOutcome' rather than direct struct initialization to ensure the outcome is sorted.
	// TODO: sort in the encoder.
	return exectypes.NewOutcome(exectypes.GetCommitReports, commitReports, cciptypes.ExecutePluginReport{})
}

func (p *Plugin) getMessagesOutcome(
	observation exectypes.Observation,
) exectypes.Outcome {
	commitReports := make([]exectypes.CommitData, 0)
	costlyMessagesSet := mapset.NewSet[cciptypes.Bytes32]()
	for _, msgID := range observation.CostlyMessages {
		costlyMessagesSet.Add(msgID)
	}

	// First ensure that all observed messages has hashes and token data.
	if err := validateHashesExist(observation.Messages, observation.Hashes); err != nil {
		p.lggr.Errorw("validate hashes exist: %w", err)
		return exectypes.Outcome{}
	}
	if err := validateTokenDataObservations(observation.Messages, observation.TokenData); err != nil {
		p.lggr.Errorw("validate token data observations: %w", err)
		return exectypes.Outcome{}
	}

	reports := observation.CommitReports.Flatten()
	// add messages to their commitReports.
	for i, report := range reports {
		report.Messages = nil
		report.CostlyMessages = nil
		for j := report.SequenceNumberRange.Start(); j <= report.SequenceNumberRange.End(); j++ {
			if msg, ok := observation.Messages[report.SourceChain][j]; ok {
				// Always add the message and hash, even if it wont be executed.
				// This slice must have an entry for each message in the commit range.
				report.Messages = append(report.Messages, msg)

				report.Hashes = append(report.Hashes, observation.Hashes[report.SourceChain][j])

				if costlyMessagesSet.Contains(msg.Header.MessageID) {
					report.CostlyMessages = append(report.CostlyMessages, msg.Header.MessageID)
				}
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
	return exectypes.NewOutcome(exectypes.GetMessages, commitReports, cciptypes.ExecutePluginReport{})
}

func (p *Plugin) getFilterOutcome(
	ctx context.Context,
	observation exectypes.Observation,
	previousOutcome exectypes.Outcome,
) (exectypes.Outcome, error) {
	commitReports := previousOutcome.CommitReports

	builder := report.NewBuilder(
		p.lggr,
		p.msgHasher,
		p.reportCodec,
		p.estimateProvider,
		p.destChain,
		report.WithMaxReportSizeBytes(maxReportLength),
		report.WithMaxGas(p.offchainCfg.BatchGasLimit),
		report.WithExtraMessageCheck(report.CheckNonces(observation.Nonces)),
	)

	outcomeReports, selectedReports, err := selectReport(
		ctx,
		p.lggr,
		commitReports,
		builder)
	if err != nil {
		return exectypes.Outcome{}, fmt.Errorf("unable to select report: %w", err)
	}

	execReport := cciptypes.ExecutePluginReport{
		ChainReports: outcomeReports,
	}

	// Must use 'NewOutcome' rather than direct struct initialization to ensure the outcome is sorted.
	// TODO: sort in the encoder.
	return exectypes.NewOutcome(exectypes.Filter, selectedReports, execReport), nil
}
